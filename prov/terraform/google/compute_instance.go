// Copyright 2018 Bull S.A.S. Atos Technologies - Bull, Rue Jean Jaures, B.P.68, 78340, Les Clayes-sous-Bois, France.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package google

import (
	"context"
	"fmt"
	"github.com/ystia/yorc/log"
	"path"
	"strings"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/pkg/errors"

	"github.com/ystia/yorc/config"
	"github.com/ystia/yorc/deployments"
	"github.com/ystia/yorc/helper/consulutil"
	"github.com/ystia/yorc/prov/terraform/commons"
)

func (g *googleGenerator) generateComputeInstance(ctx context.Context, kv *api.KV,
	cfg config.Configuration, deploymentID, nodeName, instanceName string, instanceID int,
	infrastructure *commons.Infrastructure,
	outputs map[string]string) error {

	nodeType, err := deployments.GetNodeType(kv, deploymentID, nodeName)
	if err != nil {
		return err
	}
	if nodeType != "yorc.nodes.google.Compute" {
		return errors.Errorf("Unsupported node type for %q: %s", nodeName, nodeType)
	}

	instancesPrefix := path.Join(consulutil.DeploymentKVPrefix, deploymentID,
		"topology", "instances")
	instancesKey := path.Join(instancesPrefix, nodeName)

	instance := ComputeInstance{}

	// Must be a match of regex '(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?)'
	instance.Name = strings.ToLower(cfg.ResourcesPrefix + nodeName + "-" + instanceName)
	instance.Name = strings.Replace(instance.Name, "_", "-", -1)

	// Getting string parameters
	var imageProject, imageFamily, image, serviceAccount string

	stringParams := []struct {
		pAttr        *string
		propertyName string
		mandatory    bool
	}{
		{&instance.MachineType, "machine_type", true},
		{&instance.Zone, "zone", true},
		{&imageProject, "image_project", false},
		{&imageFamily, "image_family", false},
		{&image, "image", false},
		{&instance.Description, "description", false},
		{&serviceAccount, "service_account", false},
	}

	for _, stringParam := range stringParams {
		if *stringParam.pAttr, err = deployments.GetStringNodeProperty(kv, deploymentID, nodeName,
			stringParam.propertyName, stringParam.mandatory); err != nil {
			return err
		}
	}

	// Define the boot disk from image settings
	var bootImage string
	if imageProject != "" {
		bootImage = imageProject
		if image != "" {
			bootImage = bootImage + "/" + image
		} else if imageFamily != "" {
			bootImage = bootImage + "/" + imageFamily
		} else {
			// Unexpected image project without a family or image
			return errors.Errorf("Exepected an image or family for image project %s on %s", imageProject, nodeName)
		}
	} else if image != "" {
		bootImage = image
	} else {
		bootImage = imageFamily
	}

	var bootDisk BootDisk
	if bootImage != "" {
		bootDisk.InitializeParams = InitializeParams{Image: bootImage}
	}
	instance.BootDisk = bootDisk

	// Network definition
	var noAddress bool
	if noAddress, err = deployments.GetBooleanNodeProperty(kv, deploymentID, nodeName, "no_address"); err != nil {
		return err
	}

	networkInterface := NetworkInterface{Network: "default"}
	// Define an external access if there will be an external IP address
	if !noAddress {
		hasStaticAddressReq, addressNode, err := commons.HasAnyRequirement(kv, deploymentID, nodeName, "yorc.capabilities.Assignable", "assignment")
		if err != nil {
			return err
		}
		var externalAddress string
		// External IP address can be static if required
		if hasStaticAddressReq {
			// Address Lookup
			externalAddress, err = addressLookup(ctx, kv, deploymentID, instanceName, addressNode)
			if err != nil {
				return err
			}
		}
		// else externalAddress is empty, which means an ephemeral external IP
		// address will be assigned to the instance
		accessConfig := AccessConfig{NatIP: externalAddress}
		networkInterface.AccessConfigs = []AccessConfig{accessConfig}
	}
	instance.NetworkInterfaces = []NetworkInterface{networkInterface}

	// Scheduling definition
	var preemptible bool
	if preemptible, err = deployments.GetBooleanNodeProperty(kv, deploymentID, nodeName, "preemptible"); err != nil {
		return err
	}

	if preemptible {
		instance.Scheduling = Scheduling{Preemptible: true}
	}

	// Get list of strings parameters
	var scopes []string
	if scopes, err = deployments.GetStringArrayNodeProperty(kv, deploymentID, nodeName, "scopes"); err != nil {
		return err
	}

	if serviceAccount != "" || len(scopes) > 0 {
		// Adding a service account section, where scopes can't be empty
		if len(scopes) == 0 {
			scopes = []string{"cloud-platform"}
		}
		configuredAccount := ServiceAccount{serviceAccount, scopes}
		instance.ServiceAccounts = []ServiceAccount{configuredAccount}
	}

	if instance.Tags, err = deployments.GetStringArrayNodeProperty(kv, deploymentID, nodeName, "tags"); err != nil {
		return err
	}

	// Get list of key/value pairs parameters
	if instance.Labels, err = deployments.GetKeyValuePairsNodeProperty(kv, deploymentID, nodeName, "labels"); err != nil {
		return err
	}

	if instance.Metadata, err = deployments.GetKeyValuePairsNodeProperty(kv, deploymentID, nodeName, "metadata"); err != nil {
		return err
	}

	// Get connection info (user, private key)
	user, privateKeyFilePath, err := commons.GetConnInfoFromEndpointCredentials(kv, deploymentID, nodeName)
	if err != nil {
		return err
	}

	// Add the compute instance
	commons.AddResource(infrastructure, "google_compute_instance", instance.Name, &instance)

	// Provide Consul Keys
	consulKeys := commons.ConsulKeys{Keys: []commons.ConsulKey{}}

	// Define the private IP address using the value exported by Terraform
	privateIP := fmt.Sprintf("${google_compute_instance.%s.network_interface.0.address}",
		instance.Name)

	consulKeyPrivateAddr := commons.ConsulKey{
		Path:  path.Join(instancesKey, instanceName, "/attributes/private_address"),
		Value: privateIP}

	consulKeys.Keys = append(consulKeys.Keys, consulKeyPrivateAddr)

	// Define the public IP using the value exported by Terraform
	// except if it was specified the instance shouldn't have a public address
	var accessIP string
	if noAddress {
		accessIP = privateIP
	} else {
		accessIP = fmt.Sprintf("${google_compute_instance.%s.network_interface.0.access_config.0.assigned_nat_ip}",
			instance.Name)
		consulKeyPublicAddr := commons.ConsulKey{
			Path:  path.Join(instancesKey, instanceName, "/attributes/public_address"),
			Value: accessIP}
		// For backward compatibility...
		consulKeyPublicIPAddr := commons.ConsulKey{
			Path:  path.Join(instancesKey, instanceName, "/attributes/public_ip_address"),
			Value: accessIP}

		consulKeys.Keys = append(consulKeys.Keys, consulKeyPublicAddr,
			consulKeyPublicIPAddr)
	}

	// IP ComputeAddress capability
	capabilityIPAddr := commons.ConsulKey{
		Path:  path.Join(instancesKey, instanceName, "/capabilities/endpoint/attributes/ip_address"),
		Value: accessIP}
	// Default TOSCA Attributes
	consulKeyIPAddr := commons.ConsulKey{
		Path:  path.Join(instancesKey, instanceName, "/attributes/ip_address"),
		Value: accessIP}

	consulKeys.Keys = append(consulKeys.Keys, consulKeyIPAddr, capabilityIPAddr)

	commons.AddResource(infrastructure, "consul_keys", instance.Name, &consulKeys)

	// Check the connection in order to be sure that ansible will be able to log on the instance
	nullResource := commons.Resource{}
	re := commons.RemoteExec{Inline: []string{`echo "connected"`},
		Connection: &commons.Connection{User: user, Host: accessIP,
			PrivateKey: `${file("` + privateKeyFilePath + `")}`}}
	nullResource.Provisioners = make([]map[string]interface{}, 0)
	provMap := make(map[string]interface{})
	provMap["remote-exec"] = re
	nullResource.Provisioners = append(nullResource.Provisioners, provMap)

	commons.AddResource(infrastructure, "null_resource", instance.Name+"-ConnectionCheck", &nullResource)

	return nil
}

func addressLookup(ctx context.Context, kv *api.KV, deploymentID, instanceName, addressNodeName string) (string, error) {
	log.Debugf("Address lookup for deploymentID:%q, address node name:%q, instance:%q", deploymentID, addressNodeName, instanceName)
	var address string
	res := make(chan string, 1)
	go func() {
		for {
			if address, _ := deployments.GetInstanceAttributeValue(kv, deploymentID, addressNodeName, instanceName, "ip_address"); address != nil && address.RawString() != "" {
				res <- address.RawString()
				return
			}

			select {
			case <-time.After(1 * time.Second):
			case <-ctx.Done():
				return
			}
		}
	}()

	select {
	case address = <-res:
		return address, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

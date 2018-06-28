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

// ComputeInstance represents a Google Compute Engine Virtual Machine
// See https://www.terraform.io/docs/providers/google/r/compute_instance.html
// for the latest documentation.
// See https://github.com/terraform-providers/terraform-provider-google/blob/v0.1.1/website/docs/r/compute_instance.html.markdown
// for the version currently supported by the Orchestrator
type ComputeInstance struct {
	Name              string             `json:"name"`
	MachineType       string             `json:"machine_type"`
	Zone              string             `json:"zone"`
	Disks             []Disk             `json:"disk,omitempty"`
	NetworkInterfaces []NetworkInterface `json:"network_interface"`
	Description       string             `json:"description,omitempty"`
	Labels            map[string]string  `json:"labels,omitempty"`
	Metadata          map[string]string  `json:"metadata,omitempty"`
	// NoAdress means no external IP address
	NoAddress   bool `json:"no_address,omitempty"`
	Preemptible bool `json:"preemptible,omitempty"`
	// ServiceAccounts is an array of at most one element
	ServiceAccounts []ServiceAccount `json:"service_account,omitempty"`
	Tags            []string         `json:"tags,omitempty"`
}

// Disk represents a disk attached to the compute instance
type Disk struct {
	AutoDelete bool `json:"auto_delete,omitempty"`
	// InitializeParams is an array of at most one element
	Image string `json:"image,omitempty"`
}

// NetworkInterface represents a network to attach to the instance
type NetworkInterface struct {
	Network           string         `json:"network,omitempty"`
	Subnetwork        string         `json:"subnetwork,omitempty"`
	SubNetworkProject string         `json:"subnetwork_project,omitempty"`
	Address           string         `json:"address,omitempty"`
	AccessConfigs     []AccessConfig `json:"access_config,omitempty"`
}

// AccessConfig provides IPs via which this instance can be accessed via the Internet
type AccessConfig struct {
	NatIP               string `json:"nat_ip,omitempty"`
	NetworkTier         string `json:"network_tier,omitempty"`
	PublicPtrDomainName string `json:"public_ptr_domain_name ,omitempty"`
}

// ServiceAccount to attach to the instance
type ServiceAccount struct {
	Email  string   `json:"email,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
}
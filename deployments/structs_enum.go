// Copyright 2019 Bull S.A.S. Atos Technologies - Bull, Rue Jean Jaures, B.P.68, 78340, Les Clayes-sous-Bois, France.
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

// Code generated by go-enum
// DO NOT EDIT!

package deployments

import (
	"fmt"
)

const (
	// INITIAL is a DeploymentStatus of type INITIAL
	INITIAL DeploymentStatus = iota
	// DEPLOYMENT_IN_PROGRESS is a DeploymentStatus of type DEPLOYMENT_IN_PROGRESS
	DEPLOYMENT_IN_PROGRESS
	// DEPLOYED is a DeploymentStatus of type DEPLOYED
	DEPLOYED
	// UNDEPLOYMENT_IN_PROGRESS is a DeploymentStatus of type UNDEPLOYMENT_IN_PROGRESS
	UNDEPLOYMENT_IN_PROGRESS
	// UNDEPLOYED is a DeploymentStatus of type UNDEPLOYED
	UNDEPLOYED
	// DEPLOYMENT_FAILED is a DeploymentStatus of type DEPLOYMENT_FAILED
	DEPLOYMENT_FAILED
	// UNDEPLOYMENT_FAILED is a DeploymentStatus of type UNDEPLOYMENT_FAILED
	UNDEPLOYMENT_FAILED
	// SCALING_IN_PROGRESS is a DeploymentStatus of type SCALING_IN_PROGRESS
	SCALING_IN_PROGRESS
	// UPDATE_IN_PROGRESS is a DeploymentStatus of type UPDATE_IN_PROGRESS
	UPDATE_IN_PROGRESS
	// UPDATED is a DeploymentStatus of type UPDATED
	UPDATED
	// UPDATE_FAILURE is a DeploymentStatus of type UPDATE_FAILURE
	UPDATE_FAILURE
	// PURGED is a DeploymentStatus of type PURGED
	PURGED
)

const _DeploymentStatusName = "INITIALDEPLOYMENT_IN_PROGRESSDEPLOYEDUNDEPLOYMENT_IN_PROGRESSUNDEPLOYEDDEPLOYMENT_FAILEDUNDEPLOYMENT_FAILEDSCALING_IN_PROGRESSUPDATE_IN_PROGRESSUPDATEDUPDATE_FAILUREPURGED"

var _DeploymentStatusMap = map[DeploymentStatus]string{
	0:  _DeploymentStatusName[0:7],
	1:  _DeploymentStatusName[7:29],
	2:  _DeploymentStatusName[29:37],
	3:  _DeploymentStatusName[37:61],
	4:  _DeploymentStatusName[61:71],
	5:  _DeploymentStatusName[71:88],
	6:  _DeploymentStatusName[88:107],
	7:  _DeploymentStatusName[107:126],
	8:  _DeploymentStatusName[126:144],
	9:  _DeploymentStatusName[144:151],
	10: _DeploymentStatusName[151:165],
	11: _DeploymentStatusName[165:171],
}

func (i DeploymentStatus) String() string {
	if str, ok := _DeploymentStatusMap[i]; ok {
		return str
	}
	return fmt.Sprintf("DeploymentStatus(%d)", i)
}

var _DeploymentStatusValue = map[string]DeploymentStatus{
	_DeploymentStatusName[0:7]:     0,
	_DeploymentStatusName[7:29]:    1,
	_DeploymentStatusName[29:37]:   2,
	_DeploymentStatusName[37:61]:   3,
	_DeploymentStatusName[61:71]:   4,
	_DeploymentStatusName[71:88]:   5,
	_DeploymentStatusName[88:107]:  6,
	_DeploymentStatusName[107:126]: 7,
	_DeploymentStatusName[126:144]: 8,
	_DeploymentStatusName[144:151]: 9,
	_DeploymentStatusName[151:165]: 10,
	_DeploymentStatusName[165:171]: 11,
}

// ParseDeploymentStatus attempts to convert a string to a DeploymentStatus
func ParseDeploymentStatus(name string) (DeploymentStatus, error) {
	if x, ok := _DeploymentStatusValue[name]; ok {
		return DeploymentStatus(x), nil
	}
	return DeploymentStatus(0), fmt.Errorf("%s is not a valid DeploymentStatus", name)
}

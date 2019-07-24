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

package deployments

import "encoding/json"

//go:generate go-enum --noprefix -f=structs.go

// DeploymentStatus x ENUM(
// INITIAL,
// DEPLOYMENT_IN_PROGRESS,
// DEPLOYED,
// UNDEPLOYMENT_IN_PROGRESS,
// UNDEPLOYED,
// DEPLOYMENT_FAILED,
// UNDEPLOYMENT_FAILED,
// SCALING_IN_PROGRESS,
// UPDATE_IN_PROGRESS,
// UPDATED,
// UPDATE_FAILURE,
// PURGED
// )
type DeploymentStatus int

// A TOSCAValue is the result of a resolved property or attribute
type TOSCAValue struct {
	Value    interface{}
	IsSecret bool
}

// String allows to not print secrets in case of involontary usage.
//
// If you want to allways have the actual string representation of the value
// then use RawString instead.
func (v *TOSCAValue) String() string {
	if v.IsSecret {
		return "<secret value redacted>"
	}
	return v.RawString()
}

// RawString returns the native format of the value.
// If value is a literal then it will be a string.
// If value is a complex value it will be its JSON representation.
func (v *TOSCAValue) RawString() string {
	switch t := v.Value.(type) {
	case string:
		return t
	default:
		b, err := json.Marshal(t)
		if err != nil {
			return err.Error()
		}
		return string(b)
	}
}

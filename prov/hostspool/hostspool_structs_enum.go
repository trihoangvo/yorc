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

package hostspool

import (
	"fmt"
	"strings"
)

const (
	// HostStatusFree is a HostStatus of type Free
	HostStatusFree HostStatus = iota
	// HostStatusAllocated is a HostStatus of type Allocated
	HostStatusAllocated
	// HostStatusError is a HostStatus of type Error
	HostStatusError
)

const _HostStatusName = "freeallocatederror"

var _HostStatusMap = map[HostStatus]string{
	0: _HostStatusName[0:4],
	1: _HostStatusName[4:13],
	2: _HostStatusName[13:18],
}

func (i HostStatus) String() string {
	if str, ok := _HostStatusMap[i]; ok {
		return str
	}
	return fmt.Sprintf("HostStatus(%d)", i)
}

var _HostStatusValue = map[string]HostStatus{
	_HostStatusName[0:4]:                    0,
	strings.ToLower(_HostStatusName[0:4]):   0,
	_HostStatusName[4:13]:                   1,
	strings.ToLower(_HostStatusName[4:13]):  1,
	_HostStatusName[13:18]:                  2,
	strings.ToLower(_HostStatusName[13:18]): 2,
}

// ParseHostStatus attempts to convert a string to a HostStatus
func ParseHostStatus(name string) (HostStatus, error) {
	if x, ok := _HostStatusValue[name]; ok {
		return HostStatus(x), nil
	}
	return HostStatus(0), fmt.Errorf("%s is not a valid HostStatus", name)
}

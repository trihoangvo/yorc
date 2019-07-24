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

package monitoring

import (
	"context"
	"testing"

	"github.com/ystia/yorc/v4/config"
	"github.com/ystia/yorc/v4/tasks/workflow/builder"
)

func Test_addMonitoringHook(t *testing.T) {
	type args struct {
		ctx          context.Context
		cfg          config.Configuration
		taskID       string
		deploymentID string
		target       string
		activity     builder.Activity
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addMonitoringHook(tt.args.ctx, tt.args.cfg, tt.args.taskID, tt.args.deploymentID, tt.args.target, tt.args.activity)
		})
	}
}

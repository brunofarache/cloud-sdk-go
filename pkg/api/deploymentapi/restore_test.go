// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package deploymentapi

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestRestore(t *testing.T) {
	type args struct {
		params RestoreParams
	}
	tests := []struct {
		name string
		args args
		want *models.DeploymentRestoreResponse
		err  string
	}{
		{
			name: "fails on parameter validation",
			err: multierror.NewPrefixed("deployment restore",
				apierror.ErrMissingAPI,
				apierror.ErrDeploymentID,
			).Error(),
		},
		{
			name: "fails on API error",
			args: args{params: RestoreParams{
				API:          api.NewMock(mock.New500Response(mock.NewStringBody(`{"error": "some error"}`))),
				DeploymentID: mock.ValidClusterID,
			}},
			err: `{"error": "some error"}`,
		},
		{
			name: "Succeeds",
			args: args{params: RestoreParams{
				API: api.NewMock(mock.New200Response(mock.NewStructBody(models.DeploymentRestoreResponse{
					ID: ec.String(mock.ValidClusterID),
				}))),
				DeploymentID: mock.ValidClusterID,
			}},
			want: &models.DeploymentRestoreResponse{
				ID: ec.String(mock.ValidClusterID),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Restore(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Errorf("Restore() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("Restore() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Copyright 2018 The ksonnet authors
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package actions

import (
	"testing"

	amocks "github.com/ksonnet/ksonnet/metadata/app/mocks"
	"github.com/ksonnet/ksonnet/pkg/registry"
	"github.com/ksonnet/ksonnet/pkg/registry/mocks"
	"github.com/stretchr/testify/require"
)

func TestRegistryAdd(t *testing.T) {
	withApp(t, func(appMock *amocks.App) {
		name := "new"

		cases := []struct {
			name        string
			uri         string
			version     string
			expectedURI string
			protocol    string
		}{
			{
				name:        "github",
				uri:         "github.com/foo/bar",
				expectedURI: "github.com/foo/bar",
				protocol:    registry.ProtocolGitHub,
			},
			{
				name:        "fs",
				uri:         "/path",
				expectedURI: "file:///path",
				protocol:    registry.ProtocolFilesystem,
			},
			{
				name:        "fs with URL",
				uri:         "file:///path",
				expectedURI: "file:///path",
				protocol:    registry.ProtocolFilesystem,
			},
		}

		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				a, err := NewRegistryAdd(appMock, name, tc.uri, tc.version)
				require.NoError(t, err)

				rm := &mocks.Manager{}
				rm.On("Add", appMock, "new", tc.protocol, tc.expectedURI, tc.version).Return(nil, nil)
				a.rm = rm

				err = a.Run()
				require.NoError(t, err)
			})
		}

	})
}

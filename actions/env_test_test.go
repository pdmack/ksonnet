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

	"github.com/ksonnet/ksonnet/env"
	emocks "github.com/ksonnet/ksonnet/env/mocks"
	"github.com/ksonnet/ksonnet/metadata/app"
	amocks "github.com/ksonnet/ksonnet/metadata/app/mocks"
	"github.com/stretchr/testify/require"
)

func TestEnvSet_name(t *testing.T) {
	withApp(t, func(appMock *amocks.App) {
		envName := "default"
		newName := "dev"

		nameOpt := EnvSetName(newName)

		a, err := NewEnvSet(appMock, envName, nameOpt)
		require.NoError(t, err)

		em := &emocks.Manager{}

		config := env.RenameConfig{
			App: appMock,
		}
		em.On("Rename", envName, newName, config).Return(nil)

		a.em = em

		err = a.Run()
		require.NoError(t, err)
	})
}

func TestEnvSet_namespace(t *testing.T) {
	withApp(t, func(appMock *amocks.App) {
		envName := "default"
		nsName := "ns2"

		nsOpt := EnvSetNamespace(nsName)

		a, err := NewEnvSet(appMock, envName, nsOpt)
		require.NoError(t, err)

		spec := &app.EnvironmentSpec{
			Destination: &app.EnvironmentDestinationSpec{
				Namespace: "default",
			},
		}

		updatedSpec := &app.EnvironmentSpec{
			Destination: &app.EnvironmentDestinationSpec{
				Namespace: nsName,
			},
		}

		appMock.On("Environment", envName).Return(spec, nil)

		appMock.On("AddEnvironment", envName, "", updatedSpec).Return(nil)

		err = a.Run()
		require.NoError(t, err)
	})
}

func TestEnvSet_name_and_namespace(t *testing.T) {
	withApp(t, func(appMock *amocks.App) {
		envName := "default"
		newName := "dev"
		nsName := "ns2"

		nameOpt := EnvSetName(newName)
		nsOpt := EnvSetNamespace(nsName)

		a, err := NewEnvSet(appMock, envName, nsOpt, nameOpt)
		require.NoError(t, err)

		spec := &app.EnvironmentSpec{
			Destination: &app.EnvironmentDestinationSpec{
				Namespace: "default",
			},
		}

		updatedSpec := &app.EnvironmentSpec{
			Destination: &app.EnvironmentDestinationSpec{
				Namespace: nsName,
			},
		}

		appMock.On("Environment", newName).Return(spec, nil)
		appMock.On("AddEnvironment", newName, "", updatedSpec).Return(nil)

		em := &emocks.Manager{}

		config := env.RenameConfig{
			App: appMock,
		}
		em.On("Rename", envName, newName, config).Return(nil)

		a.em = em

		err = a.Run()
		require.NoError(t, err)
	})
}

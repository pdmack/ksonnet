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

// Code generated by mockery v1.0.0
package mocks

import app "github.com/ksonnet/ksonnet/metadata/app"
import mock "github.com/stretchr/testify/mock"
import registry "github.com/ksonnet/ksonnet/pkg/registry"

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// Add provides a mock function with given fields: a, name, protoocol, uri, version
func (_m *Manager) Add(a app.App, name string, protoocol string, uri string, version string) (*registry.Spec, error) {
	ret := _m.Called(a, name, protoocol, uri, version)

	var r0 *registry.Spec
	if rf, ok := ret.Get(0).(func(app.App, string, string, string, string) *registry.Spec); ok {
		r0 = rf(a, name, protoocol, uri, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*registry.Spec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(app.App, string, string, string, string) error); ok {
		r1 = rf(a, name, protoocol, uri, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ksApp
func (_m *Manager) List(ksApp app.App) ([]registry.Registry, error) {
	ret := _m.Called(ksApp)

	var r0 []registry.Registry
	if rf, ok := ret.Get(0).(func(app.App) []registry.Registry); ok {
		r0 = rf(ksApp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]registry.Registry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(app.App) error); ok {
		r1 = rf(ksApp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

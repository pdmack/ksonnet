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

package component

import (
	"path/filepath"

	"github.com/ksonnet/ksonnet/metadata/app/mocks"
	"github.com/stretchr/testify/mock"

	"github.com/spf13/afero"
)

func appMock(root string) (*mocks.App, afero.Fs) {
	fs := afero.NewMemMapFs()
	app := &mocks.App{}
	app.On("Fs").Return(fs)
	app.On("Root").Return(root)
	app.On("LibPath", mock.AnythingOfType("string")).Return(filepath.Join(root, "lib", "v1.8.7"), nil)

	return app, fs

}

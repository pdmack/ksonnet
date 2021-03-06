// Copyright 2017 The ksonnet authors
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
package metadata

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"

	str "github.com/ksonnet/ksonnet/strings"
	"github.com/spf13/afero"
)

const (
	componentsPath  = "/componentsPath"
	componentSubdir = "subdir"
	componentFile1  = "component1.jsonnet"
	componentFile2  = "component2.jsonnet"
)

func populateComponentPaths(t *testing.T, fs afero.Fs) *manager {
	specFlag := fmt.Sprintf("file:%s", blankSwagger)

	appPath := componentsPath
	reg := newMockRegistryManager("incubator")
	m, err := initManager("componentPaths", appPath, &specFlag, &mockAPIServer, &mockNamespace, reg, fs)
	if err != nil {
		t.Fatalf("Failed to init cluster spec: %v", err)
	}

	// Create empty app file.
	components := str.AppendToPath(appPath, componentsDir)
	appFile1 := str.AppendToPath(components, componentFile1)
	f1, err := fs.OpenFile(appFile1, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		t.Fatalf("Failed to touch app file '%s'\n%v", appFile1, err)
	}
	f1.Close()

	// Create empty file in a nested directory.
	appSubdir := str.AppendToPath(components, componentSubdir)
	err = fs.MkdirAll(appSubdir, os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create directory '%s'\n%v", appSubdir, err)
	}
	appFile2 := str.AppendToPath(appSubdir, componentFile2)
	f2, err := fs.OpenFile(appFile2, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		t.Fatalf("Failed to touch app file '%s'\n%v", appFile1, err)
	}
	f2.Close()

	// Create a directory that won't be listed in the call to `ComponentPaths`.
	unlistedDir := str.AppendToPath(components, "doNotListMe")
	err = fs.MkdirAll(unlistedDir, os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create directory '%s'\n%v", unlistedDir, err)
	}

	return m
}

func TestComponentPaths(t *testing.T) {
	withFs(func(fs afero.Fs) {
		m := populateComponentPaths(t, fs)

		paths, err := m.ComponentPaths()
		if err != nil {
			t.Fatalf("Failed to find component paths: %v", err)
		}

		sort.Slice(paths, func(i, j int) bool { return paths[i] < paths[j] })

		expectedPath1 := fmt.Sprintf("%s/components/%s", componentsPath, componentFile1)
		expectedPath2 := fmt.Sprintf("%s/components/%s/%s", componentsPath, componentSubdir, componentFile2)

		if len(paths) != 2 || paths[0] != expectedPath1 || paths[1] != expectedPath2 {
			t.Fatalf("m.ComponentPaths failed; expected '%s', got '%s'", []string{expectedPath1, expectedPath2}, paths)
		}
	})
}

// TODO: this logic and tests should be moved to the components namespace.
func TestGetAllComponents(t *testing.T) {
	withFs(func(fs afero.Fs) {
		m := populateComponentPaths(t, fs)

		components, err := m.GetAllComponents()
		if err != nil {
			t.Fatalf("Failed to get all components, %v", err)
		}

		expected1 := strings.TrimSuffix(componentFile1, ".jsonnet")

		if len(components) != 1 {
			t.Fatalf("Expected exactly 1 components, got %d", len(components))
		}

		if components[0].Name(false) != expected1 {
			t.Fatalf("Expected component %s, got %s", expected1, components)
		}
	})
}

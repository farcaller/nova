// Copyright 2020 FairwindsOps Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package output

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileOutput_Send(t *testing.T) {
	path := "/tmp/output.json"
	pathcsv := "/tmp/output.csv"

	out := Output{
		HelmReleases: []ReleaseOutput{
			{
				ReleaseName: "foo",
				Namespace:   "foo",
				Installed: VersionInfo{
					Version:    "1.0",
					AppVersion: "2.0.0",
				},
				Latest: VersionInfo{
					Version:    "1.0",
					AppVersion: "2.0.0",
				},
				HelmVersion: "v3",
				Home:        "https://wiki.example.com",
				Deprecated:  false,
				Description: "Test description for foo chart",
				Icon:        "https://wiki.example.com/logo.png",
				IsOld:       false,
			},
		},
	}

	err := out.ToFile(path)
	assert.Nil(t, err)

	err = out.ToFile(pathcsv)
	assert.Nil(t, err)

	_, existsErr := os.Stat(path)
	assert.Nil(t, existsErr)

	if existsErr == nil {
		os.Remove(path)
	}

	_, existsCSVErr := os.Stat(pathcsv)
	assert.Nil(t, existsCSVErr)

	if existsErr == nil {
		os.Remove(pathcsv)
	}

}

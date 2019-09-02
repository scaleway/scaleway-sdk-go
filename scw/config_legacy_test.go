package scw

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

// TestMigrateLegacyConfig tests legacy config properly migrate to V2 config
func TestMigrateLegacyConfig(t *testing.T) {
	tests := []struct {
		name       string
		env        map[string]string
		files      map[string]string
		isMigrated bool

		expectedFiles map[string]string
	}{
		{
			name:       "No config path",
			isMigrated: false,
			files: map[string]string{
				".scwrc": v1ValidConfigFile,
			},
		},
		{
			name:       "Default config path",
			isMigrated: true,
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".scwrc": v1ValidConfigFile,
			},
			expectedFiles: map[string]string{
				".config/scw/config.yaml": v2FromV1ConfigFile,
			},
		},
		{
			name:       "V2 config already exist",
			isMigrated: false,
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2SimpleValidConfigFile,
				".scwrc":                  v1ValidConfigFile,
			},
			expectedFiles: map[string]string{
				".config/scw/config.yaml": v2SimpleValidConfigFile,
			},
		},
	}
	// create home dir
	dir := initEnv(t)

	// delete home dir and reset env variables
	defer resetEnv(t, os.Environ(), dir)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// set up env and config file(s)
			setEnv(t, test.env, test.files, dir)

			// remove config file(s)
			defer cleanEnv(t, test.files, dir)
			defer cleanEnv(t, test.expectedFiles, dir)

			isMigrated, err := MigrateLegacyConfig()
			testhelpers.AssertNoError(t, err)
			testhelpers.Equals(t, test.isMigrated, isMigrated)

			// test expected files
			for fileName, expectedContent := range test.expectedFiles {
				content, err := ioutil.ReadFile(filepath.Join(dir, fileName))
				testhelpers.AssertNoError(t, err)
				testhelpers.Equals(t, expectedContent, string(content))
			}

		})
	}
}

// v1 config
var (
	v1ValidOrganizationID = "29aa5db6-1d6d-404e-890d-f896913f9ec1"
	v1ValidToken     = "a057b0c1-eb47-4bf8-a589-72c1f2029515"
	v1Version        = "1.19"

	v1ValidConfigFile = `{
"organization":"` + v1ValidOrganizationID + `",
"token":"` + v1ValidToken + `",
"version":"` + v1Version + `"
}`

	v1InvalidConfigFile = `
"organization":"` + v1ValidOrganizationID + `",
"token":"` + v1ValidToken + `",
"version":"` + v1Version + `"
`
)

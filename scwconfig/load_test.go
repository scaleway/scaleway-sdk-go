package scwconfig

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name          string
		env           map[string]string
		files         map[string]string
		expected      *configV2
		expectedErr   string
		expectedFiles map[string]string
	}{
		// valid config
		{
			name:     "No config file",
			expected: &configV2{},
		},
		{
			name: "No config file",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			expected: &configV2{},
		},
		{
			name: "Custom config file is empty", // custom config path
			env: map[string]string{
				scwConfigPathEnv: "{HOME}/valid1/test.conf",
			},
			files: map[string]string{
				"valid1/test.conf": emptyFile,
			},
			expected: &configV2{},
		},
		{
			name: "Custom config file with valid V1",
			env: map[string]string{
				scwConfigPathEnv: "{HOME}/valid2/test.conf",
			},
			files: map[string]string{
				"valid2/test.conf": v1ValidConfigFile,
			},
			expected: v1ValidConfig,
		},
		{
			name: "Custom config file with valid V2",
			env: map[string]string{
				scwConfigPathEnv: "{HOME}/valid3/test.conf",
			},
			files: map[string]string{
				"valid3/test.conf": v2SimpleValidConfigFile,
			},
			expected: v2SimpleValidConfig,
		},
		{
			name: "Default config path with valid V2", // default config path
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2SimpleValidConfigFile,
			},
			expected: v2SimpleValidConfig,
			expectedFiles: map[string]string{
				".config/scw/config.yaml": v2SimpleValidConfigFile,
			},
		},
		{
			name: "Default config path with valid V1",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".scwrc": v1ValidConfigFile,
			},
			expected: v1ValidConfig,
			expectedFiles: map[string]string{
				".config/scw/config.yaml": v2FromV1ConfigFile,
			},
		},
		{
			name: "Default config path with valid V2 and valid V1",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2SimpleValidConfigFile,
				".scwrc":                  v1ValidConfigFile,
			},
			expected: v2SimpleValidConfig,
			expectedFiles: map[string]string{
				".config/scw/config.yaml": v2SimpleValidConfigFile,
			},
		},

		// errors
		{
			name: "Err: custom config file does not exist",
			env: map[string]string{
				scwConfigPathEnv: "{HOME}/fake/test.conf",
			},
			expectedErr: "cannot read $SCW_CONFIG_PATH: open {HOME}/fake/test.conf: no such file or directory",
		},
		{
			name: "Err: custom config file with invalid V2",
			env: map[string]string{
				scwConfigPathEnv: "{HOME}/invalid1/test.conf",
			},
			files: map[string]string{
				"invalid1/test.conf": v2SimpleInValidConfigFile,
			},
			expectedErr: "content of $SCW_CONFIG_PATH ({HOME}/invalid1/test.conf) is invalid: yaml: found unexpected end of stream",
		},
		{
			name: "Err: default config path with invalid V2",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2SimpleInValidConfigFile,
			},
			expectedErr: "content of config file {HOME}/.config/scw/config.yaml is invalid: yaml: found unexpected end of stream",
		},
		{
			name: "Err: default config path with invalid V1",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".scwrc": v1InValidConfigFile,
			},
			expectedErr: "content of config file {HOME}/.scwrc is invalid json: invalid character ':' after top-level value",
		},
	}
	dir, err := ioutil.TempDir("", "home")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	defer resetEnv(t, os.Environ())

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// set up env
			os.Clearenv()
			for key, value := range test.env {
				value = strings.Replace(value, "{HOME}", dir, -1)
				testhelpers.Ok(t, os.Setenv(key, value))
			}

			// set up config file(s)
			for path, content := range test.files {
				targetPath := filepath.Join(dir, path)
				testhelpers.Ok(t, os.MkdirAll(filepath.Dir(targetPath), 0700))
				testhelpers.Ok(t, ioutil.WriteFile(targetPath, []byte(content), defaultConfigPermission))
			}

			// load config
			config, err := Load()

			// test expected outputs
			if test.expectedErr != "" {
				testhelpers.Assert(t, err != nil, "error should not be nil")
				testhelpers.Equals(t, strings.Replace(test.expectedErr, "{HOME}", dir, -1), err.Error())
			} else {
				testhelpers.Ok(t, err)
			}
			testhelpers.Equals(t, test.expected, config)

			// test expected files
			for path, expectedContent := range test.expectedFiles {
				targetPath := filepath.Join(dir, path)
				content, err := ioutil.ReadFile(targetPath)
				testhelpers.Ok(t, err)
				testhelpers.Equals(t, expectedContent, string(content))
				testhelpers.Ok(t, os.RemoveAll(targetPath)) // delete at the end
			}

			// remove config file(s)
			for path := range test.files {
				testhelpers.Ok(t, os.RemoveAll(filepath.Join(dir, path)))
			}
		})
	}

}

// function taken from https://golang.org/src/os/env_test.go
func resetEnv(t *testing.T, origEnv []string) {
	for _, pair := range origEnv {
		// Environment variables on Windows can begin with =
		// https://blogs.msdn.com/b/oldnewthing/archive/2010/05/06/10008132.aspx
		i := strings.Index(pair[1:], "=") + 1
		if err := os.Setenv(pair[:i], pair[i+1:]); err != nil {
			t.Errorf("Setenv(%q, %q) failed during reset: %v", pair[:i], pair[i+1:], err)
		}
	}
}

const emptyFile = ""

// v2 config
var (
	v2ValidAccessKey             = "ACCESS_KEY"
	v2ValidSecretKey             = "539a6564-bf92-4dc9-a0d4-50e3ca827ecb"
	v2ValidDefaultOrganizationID = "d45a075f-18a1-4e9f-824e-43914a3ae8bd"
	v2ValidDefaultRegion         = "fr-par"
)
var v2SimpleValidConfig = &configV2{
	profile: profile{
		AccessKey:             &v2ValidAccessKey,
		SecretKey:             &v2ValidSecretKey,
		DefaultOrganizationId: &v2ValidDefaultOrganizationID,
		DefaultRegion:         &v2ValidDefaultRegion,
	},
}

var v2SimpleValidConfigFile = `
access_key: "` + v2ValidAccessKey + `"
secret_key: "` + v2ValidSecretKey + `"
default_organization_id: "` + v2ValidDefaultOrganizationID + `"
default_region: "` + v2ValidDefaultRegion + `"
`

var v2SimpleInValidConfigFile = `insecure: "bool""`

var v2FromV1ConfigFile = `secret_key: ` + v1ValidToken + `
default_organization_id: ` + v1ValidOrganizationID + `
`

// v1 config
var (
	v1ValidOrganizationID = "29aa5db6-1d6d-404e-890d-f896913f9ec1"
	v1ValidToken          = "a057b0c1-eb47-4bf8-a589-72c1f2029515"
	v1Version             = "1.19"
)

var v1ValidConfig = &configV2{
	profile: profile{
		SecretKey:             &v1ValidToken,
		DefaultOrganizationId: &v1ValidOrganizationID,
	},
}

var v1ValidConfigFile = `{
"organization":"` + v1ValidOrganizationID + `",
"token":"` + v1ValidToken + `",
"version":"` + v1Version + `"
}`

var v1InValidConfigFile = `
"organization":"` + v1ValidOrganizationID + `",
"token":"` + v1ValidToken + `",
"version":"` + v1Version + `"
`

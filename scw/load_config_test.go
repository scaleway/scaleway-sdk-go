package scw

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

// TestLoad tests all valid configuration files:
// - v2 config
// - custom-path with v2 config
// - XDG config path with v2 config
func TestLoad(t *testing.T) {
	tests := []struct {
		name          string
		env           map[string]string
		files         map[string]string
		expected      *Config
		expectedError string
		expectedFiles map[string]string
	}{
		// valid config
		{
			name: "Custom-path config is empty", // custom config path
			env: map[string]string{
				ScwConfigPathEnv: "{HOME}/valid1/test.conf",
			},
			files: map[string]string{
				"valid1/test.conf": emptyFile,
			},
			expected: &Config{},
		},
		{
			name: "Custom-path config with valid V2",
			env: map[string]string{
				ScwConfigPathEnv: "{HOME}/valid3/test.conf",
			},
			files: map[string]string{
				"valid3/test.conf": v2SimpleValidConfigFile,
			},
			expected: v2SimpleValidConfig,
		},
		{
			name: "Default config with valid V2", // default config path
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
			name: "XDG config with valid V2",
			env: map[string]string{
				"HOME":          "{HOME}",
				xdgConfigDirEnv: "{HOME}/plop",
			},
			files: map[string]string{
				"plop/scw/config.yaml": v2SimpleValidConfigFile,
			},
			expected: v2SimpleValidConfig,
		},

		// errors
		{
			name: "Err: custom-path config does not exist",
			env: map[string]string{
				ScwConfigPathEnv: "{HOME}/fake/test.conf",
			},
			expectedError: "scaleway-sdk-go: cannot read config file {HOME}/fake/test.conf: no such file or directory",
		},
		{
			name: "Err: custom-path config with invalid V2",
			env: map[string]string{
				ScwConfigPathEnv: "{HOME}/invalid1/test.conf",
			},
			files: map[string]string{
				"invalid1/test.conf": v2SimpleInvalidConfigFile,
			},
			expectedError: "scaleway-sdk-go: content of config file {HOME}/invalid1/test.conf is invalid: yaml: found unexpected end of stream",
		},
		{
			name: "Err: default config with invalid V2",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2SimpleInvalidConfigFile,
			},
			expectedError: "scaleway-sdk-go: content of config file {HOME}/.config/scw/config.yaml is invalid: yaml: found unexpected end of stream",
		},
		{
			name: "Err: default config with invalid profile",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2SimpleConfigFileWithInvalidProfile,
			},
			expectedError: "scaleway-sdk-go: given profile flantier does not exist",
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
			test.expectedError = strings.Replace(test.expectedError, "{HOME}", dir, -1)

			// remove config file(s)
			defer cleanEnv(t, test.files, dir)

			// load config
			config, err := LoadConfig()

			// test expected outputs
			if test.expectedError != "" {
				if err == nil {
					_, tmpErr := config.GetActiveProfile()
					if tmpErr != nil {
						testhelpers.Equals(t, test.expectedError, tmpErr.Error())
						return
					}
				}
				testhelpers.Assert(t, err != nil, "error should not be nil")
				testhelpers.Equals(t, test.expectedError, err.Error())
			} else {
				testhelpers.AssertNoError(t, err)
				testhelpers.Equals(t, test.expected, config)

				_, err = config.GetActiveProfile()
				testhelpers.AssertNoError(t, err)
			}

			// test expected files
			for path, expectedContent := range test.expectedFiles {
				targetPath := filepath.Join(dir, path)
				content, err := ioutil.ReadFile(targetPath)
				testhelpers.AssertNoError(t, err)
				testhelpers.Equals(t, expectedContent, string(content))
				testhelpers.AssertNoError(t, os.RemoveAll(targetPath)) // delete at the end
			}
		})
	}
}

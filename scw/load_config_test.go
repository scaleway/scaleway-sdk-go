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
// - v1 to v2 config migration
// - custom-path with v2 config
// - custom-path with v1 config
// - XDG config path with v2 config
// - Windows config path with v2 config
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
			name: "Custom-path config is empty", // custom config path
			env: map[string]string{
				scwConfigPathEnv: "{HOME}/valid1/test.conf",
			},
			files: map[string]string{
				"valid1/test.conf": emptyFile,
			},
			expected: &configV2{},
		},
		{
			name: "Custom-path config with valid V1",
			env: map[string]string{
				scwConfigPathEnv: "{HOME}/valid2/test.conf",
			},
			files: map[string]string{
				"valid2/test.conf": v1ValidConfigFile,
			},
			expected: v1ValidConfig,
		},
		{
			name: "Custom-path config with valid V2",
			env: map[string]string{
				scwConfigPathEnv: "{HOME}/valid3/test.conf",
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
			name: "Default config with valid V1",
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
			name: "Default config with valid V2 and valid V1",
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
		{
			name: "Windows config with valid V2",
			env: map[string]string{
				windowsHomeDirEnv: "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2SimpleValidConfigFile,
			},
			expected: v2SimpleValidConfig,
		},

		// errors
		{
			name: "Err: custom-path config does not exist",
			env: map[string]string{
				scwConfigPathEnv: "{HOME}/fake/test.conf",
			},
			expectedErr: "cannot read config file SCW_CONFIG_PATH: open {HOME}/fake/test.conf: no such file or directory",
		},
		{
			name: "Err: custom-path config with invalid V2",
			env: map[string]string{
				scwConfigPathEnv: "{HOME}/invalid1/test.conf",
			},
			files: map[string]string{
				"invalid1/test.conf": v2SimpleInvalidConfigFile,
			},
			expectedErr: "content of config file {HOME}/invalid1/test.conf is invalid: yaml: found unexpected end of stream",
		},
		{
			name: "Err: default config with invalid V2",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2SimpleInvalidConfigFile,
			},
			expectedErr: "content of config file {HOME}/.config/scw/config.yaml is invalid: yaml: found unexpected end of stream",
		},
		{
			name: "Err: default config with invalid V1",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".scwrc": v1InvalidConfigFile,
			},
			expectedErr: "content of config file {HOME}/.scwrc is invalid json: invalid character ':' after top-level value",
		},
		{
			name: "Err: default config with invalid profile",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2SimpleConfigFileWithInvalidProfile,
			},
			expectedErr: "profile flantier does not exist in config file {HOME}/.config/scw/config.yaml",
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

			// load config
			config, err := LoadConfig()

			// test expected outputs
			if test.expectedErr != "" {
				testhelpers.Assert(t, err != nil, "error should not be nil")
				testhelpers.Equals(t, strings.Replace(test.expectedErr, "{HOME}", dir, -1), err.Error())
			} else {
				testhelpers.AssertNoError(t, err)
				testhelpers.Equals(t, test.expected, config)
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

func initEnv(t *testing.T) string {
	dir, err := ioutil.TempDir("", "home")
	if err != nil {
		t.Fatal(err)
	}
	return dir
}

func cleanEnv(t *testing.T, files map[string]string, homeDir string) {
	for path := range files {
		testhelpers.AssertNoError(t, os.RemoveAll(filepath.Join(homeDir, path)))
	}
}

func setEnv(t *testing.T, env, files map[string]string, homeDir string) {
	os.Clearenv()
	for key, value := range env {
		value = strings.Replace(value, "{HOME}", homeDir, -1)
		testhelpers.AssertNoError(t, os.Setenv(key, value))
	}

	for path, content := range files {
		targetPath := filepath.Join(homeDir, path)
		testhelpers.AssertNoError(t, os.MkdirAll(filepath.Dir(targetPath), 0700))
		testhelpers.AssertNoError(t, ioutil.WriteFile(targetPath, []byte(content), defaultConfigPermission))
	}
}

// function taken from https://golang.org/src/os/env_test.go
func resetEnv(t *testing.T, origEnv []string, homeDir string) {
	testhelpers.AssertNoError(t, os.RemoveAll(homeDir))
	for _, pair := range origEnv {
		// Environment variables on Windows can begin with =
		// https://blogs.msdn.com/b/oldnewthing/archive/2010/05/06/10008132.aspx
		i := strings.Index(pair[1:], "=") + 1
		if err := os.Setenv(pair[:i], pair[i+1:]); err != nil {
			t.Errorf("Setenv(%q, %q) failed during reset: %v", pair[:i], pair[i+1:], err)
		}
	}
}

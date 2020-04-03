package scw

import (
	"errors"
	"os"
	"path/filepath"
)

const (
	// XDG wiki: https://wiki.archlinux.org/index.php/XDG_Base_Directory
	xdgConfigDirEnv = "XDG_CONFIG_HOME"
	xdgCacheDirEnv  = "XDG_CACHE_HOME"

	unixHomeDirEnv    = "HOME"
	windowsHomeDirEnv = "USERPROFILE"

	defaultConfigFileName = "config.yaml"
)

var (
	// ErrNoHomeDir errors when no user directory is found
	ErrNoHomeDir = errors.New("user home directory not found")
)

// GetCacheDirectory returns the default cache directory.
// Cache directory is based on the following priority order:
// - $SCW_CACHE_DIR
// - $XDG_CACHE_HOME/scw
// - $HOME/.cache/scw
// - $USERPROFILE/.cache/scw
func GetCacheDirectory() string {
	cacheDir := ""
	switch {
	case os.Getenv(scwCacheDirEnv) != "":
		cacheDir = os.Getenv(scwCacheDirEnv)
	case os.Getenv(xdgCacheDirEnv) != "":
		cacheDir = filepath.Join(os.Getenv(xdgCacheDirEnv), "scw")
	case os.Getenv(unixHomeDirEnv) != "":
		cacheDir = filepath.Join(os.Getenv(unixHomeDirEnv), ".cache", "scw")
	case os.Getenv(windowsHomeDirEnv) != "":
		cacheDir = filepath.Join(os.Getenv(windowsHomeDirEnv), ".cache", "scw")
	default:
		// TODO: fallback on local folder?
	}

	// Clean the cache directory path when exiting the function
	return filepath.Clean(cacheDir)
}

// GetConfigPath returns the default path.
// Default path is based on the following priority order:
// - $SCW_CONFIG_PATH
// - $XDG_CONFIG_HOME/scw/config.yaml
// - $HOME/.config/scw/config.yaml
// - $USERPROFILE/.config/scw/config.yaml
func GetConfigPath() string {
	configPath := os.Getenv(scwConfigPathEnv)
	if configPath == "" {
		configPath, _ = getConfigV2FilePath()
	}
	return filepath.Clean(configPath)
}

// getConfigV2FilePath returns the path to the v2 config file
func getConfigV2FilePath() (string, bool) {
	configDir, err := getScwConfigDir()
	if err != nil {
		return "", false
	}
	return filepath.Clean(filepath.Join(configDir, defaultConfigFileName)), true
}

// getConfigV1FilePath returns the path to the v1 config file
func getConfigV1FilePath() (string, bool) {
	path, err := GetHomeDir()
	if err != nil {
		return "", false
	}
	return filepath.Clean(filepath.Join(path, ".scwrc")), true
}

// getScwConfigDir returns the path to scw config folder
func getScwConfigDir() (string, error) {
	if xdgPath := os.Getenv(xdgConfigDirEnv); xdgPath != "" {
		return filepath.Join(xdgPath, "scw"), nil
	}

	homeDir, err := GetHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".config", "scw"), nil
}

// GetHomeDir returns the path to your home directory
func GetHomeDir() (string, error) {
	switch {
	case os.Getenv(unixHomeDirEnv) != "":
		return os.Getenv(unixHomeDirEnv), nil
	case os.Getenv(windowsHomeDirEnv) != "":
		return os.Getenv(windowsHomeDirEnv), nil
	default:
		return "", ErrNoHomeDir
	}
}

func fileExist(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

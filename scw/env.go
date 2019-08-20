package scw

import (
	"os"

	"github.com/scaleway/scaleway-sdk-go/logger"
)

func LoadEnvProfile() (*profile, error) {
	return &profile{}, nil
}

func getenv(upToDateKey string, deprecatedKeys ...string) (string, string, bool) {
	value, exist := os.LookupEnv(upToDateKey)
	if exist {
		logger.Infof("reading value from %s", upToDateKey)
		return value, upToDateKey, true
	}

	for _, key := range deprecatedKeys {
		value, exist := os.LookupEnv(key)
		if exist {
			logger.Infof("reading value from %s", key)
			logger.Warningf("%s is deprecated, please use %s instead", key, upToDateKey)
			return value, key, true
		}
	}

	return "", "", false
}

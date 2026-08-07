package rdb

import (
	"fmt"
)

func (s *API) FetchLatestEngineVersion(engineName string) (*EngineVersion, error) {
	engines, err := s.ListDatabaseEngines(&ListDatabaseEnginesRequest{})
	if err != nil {
		return nil, err
	}

	var latestEngineVersion *EngineVersion
	for _, engine := range engines.Engines {
		if engine.Name == engineName {
			if len(engine.Versions) > 0 {
				latestEngineVersion = engine.Versions[0]
				break
			}
		}
	}

	if latestEngineVersion == nil {
		return nil, fmt.Errorf("no versions found for engine: %s", engineName)
	}
	return latestEngineVersion, nil
}

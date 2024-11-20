package rdb

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
)

func TestFetchLatestEngineVersion(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("test-latest-engine")

	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop())
	}()

	t.Run("fetch latest version for PostgreSQL", func(t *testing.T) {
		rdbAPI := NewAPI(client)

		version, err := rdbAPI.FetchLatestEngineVersion("PostgreSQL")
		testhelpers.AssertNoError(t, err)

		testhelpers.Equals(t, "PostgreSQL-15", version)
	})

	t.Run("fetch latest version for MySQL", func(t *testing.T) {
		rdbAPI := NewAPI(client)

		version, err := rdbAPI.FetchLatestEngineVersion("MySQL")
		testhelpers.AssertNoError(t, err)

		testhelpers.Equals(t, "MySQL-8", version)
	})

	t.Run("fetch version for unknown engine", func(t *testing.T) {
		rdbAPI := NewAPI(client)

		version, _ := rdbAPI.FetchLatestEngineVersion("UnknownEngine")
		testhelpers.Equals(t, "", version)
	})
}

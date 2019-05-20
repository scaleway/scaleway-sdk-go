package auth

import (
	"net/http"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

func TestToken_Headers(t *testing.T) {
	const (
		accessKey = "ACCESS_KEY"
		secretKey = "SECRET_KEY"
	)
	auth := NewToken(accessKey, secretKey)
	testhelpers.Equals(t, http.Header{
		"X-Auth-Token": []string{secretKey},
	}, auth.Headers())
}

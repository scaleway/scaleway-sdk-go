package vpc

import (
	"errors"
	"net/http"
	"time"

	"github.com/scaleway/scaleway-sdk-go/scw"
)

func (s *API) TryDeletingPrivateNetwork(req *DeletePrivateNetworkRequest, retriesLeft int, opts ...scw.RequestOption) (*PrivateNetwork, error) {
	err := s.DeletePrivateNetwork(&DeletePrivateNetworkRequest{
		PrivateNetworkID: req.PrivateNetworkID,
		Region:           req.Region,
	}, opts...)

	var respErr *scw.ResponseError
	if errors.As(err, &respErr) && respErr.StatusCode == http.StatusInternalServerError {
		time.Sleep(5 * time.Second)
		if retriesLeft > 0 {
			return s.TryDeletingPrivateNetwork(req, retriesLeft-1, opts...)
		}
	}

	return nil, err
}

package vpcgw

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultTimeout       = 5 * time.Minute
	defaultRetryInterval = 15 * time.Second
)

// WaitForDHCPEntriesRequest is used by WaitForDHCPEntries method
type WaitForDHCPEntriesRequest struct {
	GatewayNetworkID *string
	MacAddress       string

	Zone          scw.Zone
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForDHCPEntries waits for at least one dhcp entry with the correct mac address.
// This function can be used to wait for an instance to use dhcp
func (s *API) WaitForDHCPEntries(req *WaitForDHCPEntriesRequest, opts ...scw.RequestOption) (*ListDHCPEntriesResponse, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	dhcpEntries, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			entries, err := s.ListDHCPEntries(&ListDHCPEntriesRequest{
				Zone:             req.Zone,
				GatewayNetworkID: req.GatewayNetworkID,
				MacAddress:       &req.MacAddress,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			containsMacAddress := false
			for _, entry := range entries.DHCPEntries {
				if entry.MacAddress == req.MacAddress {
					containsMacAddress = true
					break
				}
			}

			return entries, containsMacAddress, err
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for gateway network failed")
	}

	return dhcpEntries.(*ListDHCPEntriesResponse), nil
}

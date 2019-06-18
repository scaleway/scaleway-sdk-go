package instance

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

// WaitForServerRequest is used by WaitForServer method
type WaitForServerRequest struct {
	ServerID string
	Zone     utils.Zone

	// Timeout: maximum time to wait before (default: 5 minutes)
	Timeout time.Duration
}

// WaitForServer wait for the server to be in a "terminal state" before returning.
// This function can be used to wait for a server to be started for example.
func (s *API) WaitForServer(req *WaitForServerRequest) (*Server, error) {

	if req.Timeout == 0 {
		req.Timeout = 5 * time.Minute
	}

	terminalStatus := map[ServerState]struct{}{
		ServerStateStopped:        {},
		ServerStateStoppedInPlace: {},
		ServerStateLocked:         {},
		ServerStateRunning:        {},
	}

	server, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, error, bool) {
			res, err := s.GetServer(&GetServerRequest{
				ServerID: req.ServerID,
				Zone:     req.Zone,
			})

			if err != nil {
				return nil, err, false
			}
			_, isTerminal := terminalStatus[res.Server.State]

			return res.Server, err, isTerminal
		},
		Timeout:          req.Timeout,
		IntervalStrategy: async.LinearIntervalStrategy(5 * time.Second),
	})

	return server.(*Server), err
}

// GetServerTypeRequest is used by GetServerType
type GetServerTypeRequest struct {
	Zone utils.Zone
	Name string
}

// GetServerType get server type info by it's name.
func (s *API) GetServerType(req *GetServerTypeRequest) (*ServerType, error) {

	res, err := s.ListServersTypes(&ListServersTypesRequest{
		Zone: req.Zone,
	}, scw.WithAllPages())

	if err != nil {
		return nil, err
	}

	if serverType, exist := res.Servers[req.Name]; exist {
		return serverType, nil
	}

	return nil, errors.New("could not find server type %q", req.Name)
}

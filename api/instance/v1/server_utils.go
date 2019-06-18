package instance

import (
	"fmt"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

// waitForServerRequest is used by waitForServer method
type waitForServerRequest struct {
	ServerID string
	Zone     utils.Zone

	// Timeout: maximum time to wait before (default: 5 minutes)
	Timeout time.Duration
}

// waitForServer wait for the server to be in a "terminal state" before returning.
// This function can be used to wait for a server to be started for example.
func (s *API) waitForServer(req *waitForServerRequest) (*Server, error) {

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

// DoActionAndWaitRequest is used by DoActionAndWait method
type DoActionAndWaitRequest struct {
	ServerID string
	Zone     utils.Zone
	Action   ServerAction
	Timeout  time.Duration
}

// DoActionAndWait start an action and wait for the server to be in the correct "terminal state"
// expected by this action.
func (s *API) DoActionAndWait(req *DoActionAndWaitRequest) error {
	_, err := s.ServerAction(&ServerActionRequest{
		Zone:     req.Zone,
		ServerID: req.ServerID,
		Action:   req.Action,
	})
	if err != nil {
		return err
	}

	finalServer, err := s.waitForServer(&waitForServerRequest{
		Zone:     req.Zone,
		ServerID: req.ServerID,
		Timeout:  req.Timeout,
	})
	if err != nil {
		return err
	}

	// check the action was properly executed
	expectedState := ServerState("")
	switch req.Action {
	case ServerActionPoweron, ServerActionReboot:
		expectedState = ServerStateRunning
	case ServerActionPoweroff:
		expectedState = ServerStateStopped
	case ServerActionStopInPlace:
		expectedState = ServerStateStoppedInPlace
	}

	// backup can be performed from any state
	if req.Action != ServerActionBackup && finalServer.State != expectedState {
		return fmt.Errorf("expected state %s but found %s: %s", expectedState, finalServer.State, finalServer.StateDetail)
	}

	return nil
}

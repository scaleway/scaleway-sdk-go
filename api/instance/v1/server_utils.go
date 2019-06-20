package instance

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

// waitForServerRequest is used by waitForServer method
type waitForServerRequest struct {
	ServerID string
	Zone     utils.Zone
	Timeout  time.Duration
}

// waitForServer wait for the server to be in a "terminal state" before returning.
// This function can be used to wait for a server to be started for example.
func (s *API) waitForServer(req *waitForServerRequest) (*Server, scw.SdkError) {

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
	if err != nil {
		return nil, errors.Wrap(err, "waiting for server failed")
	}
	return server.(*Server), nil
}

// ServerActionAndWaitRequest is used by ServerActionAndWait method
type ServerActionAndWaitRequest struct {
	ServerID string
	Zone     utils.Zone
	Action   ServerAction

	// Timeout: maximum time to wait before (default: 5 minutes)
	Timeout time.Duration
}

// ServerActionAndWait start an action and wait for the server to be in the correct "terminal state"
// expected by this action.
func (s *API) ServerActionAndWait(req *ServerActionAndWaitRequest) error {
	if req.Timeout == 0 {
		req.Timeout = 5 * time.Minute
	}

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
	expectedState := ServerState("unknown")
	switch req.Action {
	case ServerActionPoweron, ServerActionReboot:
		expectedState = ServerStateRunning
	case ServerActionPoweroff:
		expectedState = ServerStateStopped
	case ServerActionStopInPlace:
		expectedState = ServerStateStoppedInPlace
	}

	// backup can be performed from any state
	if expectedState != ServerState("unknown") && finalServer.State != expectedState {
		return errors.New("expected state %s but found %s: %s", expectedState, finalServer.State, finalServer.StateDetail)
	}

	return nil
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

// GetServerUserDataRequest is used by GetServerType method.
type GetServerUserDataRequest struct {
	Zone     utils.Zone `json:"-"`
	ServerID string     `json:"-"`

	// Key define the user data key to get
	Key string `json:"-"`
}

// GetServerUserData get user data
//
// Get the content of a user data with the given key on a server
func (s *API) GetServerUserData(req *GetServerUserDataRequest, opts ...scw.RequestOption) (io.ReadCloser, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	if fmt.Sprint(req.Key) == "" {
		return nil, errors.New("field Key cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/user_data/" + fmt.Sprint(req.Key),
		Headers: http.Header{},
	}

	buffer := &bytes.Buffer{}

	err = s.client.Do(scwReq, buffer, opts...)
	if err != nil {
		return nil, err
	}

	res := ioutil.NopCloser(buffer)
	return res, nil
}

type SetServerUserDataRequest struct {
	Zone     utils.Zone `json:"-"`
	ServerID string     `json:"-"`

	// Key define the user data key to set
	Key string `json:"-"`

	// Content define the data to set
	Content io.Reader
}

// SetServerUserData add/Set user data
//
// Add or update a user data with the given key on a server
func (s *API) SetServerUserData(req *SetServerUserDataRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return errors.New("field ServerID cannot be empty in request")
	}

	if fmt.Sprint(req.Key) == "" {
		return errors.New("field Key cannot be empty in request")
	}

	if req.Content == nil {
		return errors.New("field Content cannot be nil in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/user_data/" + fmt.Sprint(req.Key),
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req.Content)
	if err != nil {
		return err
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}

	return nil
}

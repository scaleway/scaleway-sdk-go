// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package iot_admin provides methods and message types of the iot_admin v1beta1 API.
package iot_admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/scw"
)

// always import dependencies
var (
	_ fmt.Stringer
	_ json.Unmarshaler
	_ url.URL
	_ net.IP
	_ http.Header
	_ bytes.Reader
	_ time.Time
	_ = strings.Join

	_ scw.ScalewayRequest
	_ marshaler.Duration
	_ scw.File
	_ = parameter.AddToQuery
	_ = namegenerator.GetRandomName
)

// ClusterStatusRequest: cluster status request.
type ClusterStatusRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// FindResourcesRequest: find resources request.
type FindResourcesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ID: any kind of UUID known to IoT (organization, hub, device, etc.).
	ID string `json:"-"`
}

// HubForceRedeployRequest: hub force redeploy request.
type HubForceRedeployRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HubID: the ID of the IoT Hub.
	HubID string `json:"-"`
}

// HubLogRequest: hub log request.
type HubLogRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HubID: the ID of the IoT Hub.
	HubID string `json:"-"`
}

// HubStatusRequest: hub status request.
type HubStatusRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HubID: the ID of the IoT Hub.
	HubID string `json:"-"`
}

// HubSyncRequest: hub sync request.
type HubSyncRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HubID: the ID of the IoT Hub.
	HubID string `json:"-"`
}

// HubTestRequest: hub test request.
type HubTestRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HubID: the ID of the IoT Hub.
	HubID string `json:"-"`
}

// StringReply: string reply.
type StringReply struct {
	// Region: region of the API.
	Region scw.Region `json:"region"`

	// Message: the message returned by the API.
	Message string `json:"message"`
}

// This API allows admin console to troubleshot customer problems.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// FindResources: Find customer resources.
func (s *API) FindResources(req *FindResourcesRequest, opts ...scw.RequestOption) (*StringReply, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ID) == "" {
		return nil, errors.New("field ID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/resources/" + fmt.Sprint(req.ID) + "",
	}

	var resp StringReply

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ClusterStatus: Fetch information about the IoT platform cluster.
func (s *API) ClusterStatus(req *ClusterStatusRequest, opts ...scw.RequestOption) (*StringReply, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/cluster-status",
	}

	var resp StringReply

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// HubStatus: Get the status of a hub and its broker.
func (s *API) HubStatus(req *HubStatusRequest, opts ...scw.RequestOption) (*StringReply, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return nil, errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/status",
	}

	var resp StringReply

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// HubLog: Get the logs of a hub's broker.
func (s *API) HubLog(req *HubLogRequest, opts ...scw.RequestOption) (*StringReply, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return nil, errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/log",
	}

	var resp StringReply

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// HubTest: Try to connect directly to the hub's broker (bypassing the proxy).
func (s *API) HubTest(req *HubTestRequest, opts ...scw.RequestOption) (*StringReply, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return nil, errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/test",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp StringReply

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// HubSync: Synchronize the state of a hub's broker with its DB record.
func (s *API) HubSync(req *HubSyncRequest, opts ...scw.RequestOption) (*StringReply, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return nil, errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/sync",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp StringReply

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// HubForceRedeploy: Remove the broker's deployment and let Kubernetes respawn it.
func (s *API) HubForceRedeploy(req *HubForceRedeployRequest, opts ...scw.RequestOption) (*StringReply, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return nil, errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/force-redeploy",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp StringReply

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package test provides methods and message types of the test v1 API.
package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/marshaler"
	"github.com/scaleway/scaleway-sdk-go/internal/parameter"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/scw"
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

// NoAuthAPI: no Auth Service
type NoAuthAPI struct {
	client *scw.Client
}

// NewNoAuthAPI returns a NoAuthAPI object from a Scaleway client.
func NewNoAuthAPI(client *scw.Client) *NoAuthAPI {
	return &NoAuthAPI{
		client: client,
	}
}

// RegionalAPI: regional API
type RegionalAPI struct {
	client *scw.Client
}

// NewRegionalAPI returns a RegionalAPI object from a Scaleway client.
func NewRegionalAPI(client *scw.Client) *RegionalAPI {
	return &RegionalAPI{
		client: client,
	}
}

// StreamAPI: stream Service
type StreamAPI struct {
	client *scw.Client
}

// NewStreamAPI returns a StreamAPI object from a Scaleway client.
func NewStreamAPI(client *scw.Client) *StreamAPI {
	return &StreamAPI{
		client: client,
	}
}

// ZoneAPI: zone Test API
type ZoneAPI struct {
	client *scw.Client
}

// NewZoneAPI returns a ZoneAPI object from a Scaleway client.
func NewZoneAPI(client *scw.Client) *ZoneAPI {
	return &ZoneAPI{
		client: client,
	}
}

// EchoMessage: echo message
type EchoMessage struct {
	Str string `json:"str"`

	Strs []string `json:"strs"`
}

type GetZoneResponse struct {
	Zone scw.Zone `json:"zone"`
}

// MetadataResponse: metadata response
type MetadataResponse struct {
	Metadata map[string]string `json:"metadata"`
}

// getRegionResponse: get region response
type getRegionResponse struct {
	Region scw.Region `json:"region"`
}

// Service NoAuthAPI

// Service RegionalAPI

type RegionalAPIGetServiceInfoRequest struct {
	Region scw.Region `json:"-"`
}

func (s *RegionalAPI) GetServiceInfo(req *RegionalAPIGetServiceInfoRequest, opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/regions/" + fmt.Sprint(req.Region) + "",
		Headers: http.Header{},
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RegionalAPIGetRegionRequest struct {
	Region scw.Region `json:"-"`
}

func (s *RegionalAPI) GetRegion(req *RegionalAPIGetRegionRequest, opts ...scw.RequestOption) (*getRegionResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/regions/" + fmt.Sprint(req.Region) + "/region",
		Headers: http.Header{},
	}

	var resp getRegionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RegionalAPIGetMetadataRequest struct {
	Region scw.Region `json:"-"`
}

func (s *RegionalAPI) GetMetadata(req *RegionalAPIGetMetadataRequest, opts ...scw.RequestOption) (*MetadataResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/regions/" + fmt.Sprint(req.Region) + "/metadata",
		Headers: http.Header{},
	}

	var resp MetadataResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RegionalAPIPostEchoRequest struct {
	Region scw.Region `json:"-"`

	Str *string `json:"str"`

	Strs []string `json:"strs"`
}

func (s *RegionalAPI) PostEcho(req *RegionalAPIPostEchoRequest, opts ...scw.RequestOption) (*EchoMessage, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Str == nil {
		req.Str = scw.StringPtr(namegenerator.GetRandomName("name"))
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/regions/" + fmt.Sprint(req.Region) + "/echo",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp EchoMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Service StreamAPI

// Service ZoneAPI

type ZoneAPIGetZoneRequest struct {
	Zone scw.Zone `json:"-"`
}

// GetZone: get a zone
func (s *ZoneAPI) GetZone(req *ZoneAPIGetZoneRequest, opts ...scw.RequestOption) (*GetZoneResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/zone",
		Headers: http.Header{},
	}

	var resp GetZoneResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ZoneAPIGetMetadataRequest struct {
	Zone scw.Zone `json:"-"`
}

// GetMetadata: get metadata
func (s *ZoneAPI) GetMetadata(req *ZoneAPIGetMetadataRequest, opts ...scw.RequestOption) (*MetadataResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/metadata",
		Headers: http.Header{},
	}

	var resp MetadataResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ZoneAPIPostEchoRequest struct {
	Zone scw.Zone `json:"-"`

	Str *string `json:"str"`

	Strs []string `json:"strs"`
}

func (s *ZoneAPI) PostEcho(req *ZoneAPIPostEchoRequest, opts ...scw.RequestOption) (*EchoMessage, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.Str == nil {
		req.Str = scw.StringPtr(namegenerator.GetRandomName("name"))
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/echo",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp EchoMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

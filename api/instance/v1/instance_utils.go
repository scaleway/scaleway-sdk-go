package instance

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

// AttachIpRequest contains the parameters to attach an IP to a server
type AttachIpRequest struct {
	Zone   utils.Zone `json:"-"`
	IpId   string     `json:"-"`
	Server *string    `json:"server"`
}

// AttachIp attaches an IP to a server.
func (s *Api) AttachIp(req *AttachIpRequest, opts ...scw.RequestOption) (*Ip, error) {
	var err error

	if req.Server == nil || *req.Server == "" {
		return nil, fmt.Errorf("couldn't attach ip: server id was not given")
	}

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IpId) + "",
		Headers: http.Header{},
	}
	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp UpdateIpResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp.Ip, nil
}

// AttachIpRequest contains the parameters to detach an IP from a server
type DetachIpRequest struct {
	Zone utils.Zone `json:"-"`
	IpId string     `json:"-"`
}

// AttachIp detaches an IP from a server.
func (s *Api) DetachIp(req *DetachIpRequest, opts ...scw.RequestOption) (*Ip, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	detachReq := &struct {
		Server *string `json:"server"`
	}{
		Server: nil,
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IpId) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(detachReq)
	if err != nil {
		return nil, err
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp UpdateIpResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp.Ip, nil
}

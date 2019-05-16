package instance

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

// AttachIPRequest contains the parameters to attach an IP to a server
type AttachIPRequest struct {
	Zone     utils.Zone `json:"-"`
	IPID     string     `json:"-"`
	ServerID string     `json:"server"`
}

// AttachIP attaches an IP to a server.
func (s *Api) AttachIP(req *AttachIPRequest, opts ...scw.RequestOption) (*Ip, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IPID) + "",
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

// DetachIPRequest contains the parameters to detach an IP from a server
type DetachIPRequest struct {
	Zone utils.Zone `json:"-"`
	IPID string     `json:"-"`
}

// DetachIP detaches an IP from a server.
func (s *Api) DetachIP(req *DetachIPRequest, opts ...scw.RequestOption) (*Ip, error) {
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
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IPID) + "",
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

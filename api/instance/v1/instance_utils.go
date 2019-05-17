package instance

import (
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

// AttachIPRequest contains the parameters to attach an IP to a server
type AttachIPRequest struct {
	Zone     utils.Zone `json:"-"`
	IPID     string     `json:"-"`
	ServerID string     `json:"server"`
}

// AttachIPResponse contains the updated IP after detaching
type AttachIPResponse struct {
	IP *Ip
}

// AttachIP attaches an IP to a server.
func (s *Api) AttachIP(req *AttachIPRequest, opts ...scw.RequestOption) (*AttachIPResponse, error) {
	var ptrServerID = &req.ServerID
	ipResponse, err := s.updateIp(&updateIpRequest{
		Zone:   req.Zone,
		IpId:   req.IPID,
		Server: &ptrServerID,
	})
	if err != nil {
		return nil, err
	}

	return &AttachIPResponse{IP: ipResponse.Ip}, nil
}

// DetachIPRequest contains the parameters to detach an IP from a server
type DetachIPRequest struct {
	Zone utils.Zone `json:"-"`
	IPID string     `json:"-"`
}

// DetachIPResponse contains the updated IP after detaching
type DetachIPResponse struct {
	IP *Ip
}

// DetachIP detaches an IP from a server.
func (s *Api) DetachIP(req *DetachIPRequest, opts ...scw.RequestOption) (*DetachIPResponse, error) {
	var ptrServerID *string
	ipResponse, err := s.updateIp(&updateIpRequest{
		Zone:   req.Zone,
		IpId:   req.IPID,
		Server: &ptrServerID,
	})
	if err != nil {
		return nil, err
	}

	return &DetachIPResponse{IP: ipResponse.Ip}, nil
}

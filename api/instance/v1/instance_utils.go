package instance

import (
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

// AttachIPRequest contains the parameters to attach an IP to a server
type AttachIPRequest struct {
	Zone     utils.Zone `json:"-"`
	IPID     string     `json:"-"`
	ServerID string     `json:"server_id"`
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

// SetReverseForIPRequest contains the parameters to set the reverse DNS for an IP
type SetReverseForIPRequest struct {
	Zone    utils.Zone `json:"-"`
	IPID    string     `json:"-"`
	Reverse string     `json:"reverse"`
}

// SetReverseForIPResponse contains the updated IP after setting the reverse
type SetReverseForIPResponse struct {
	IP *Ip
}

// SetReverseForIP sets the reverse DNS for an IP.
func (s *Api) SetReverseForIP(req *SetReverseForIPRequest, opts ...scw.RequestOption) (*SetReverseForIPResponse, error) {
	var ptrReverse = &req.Reverse
	ipResponse, err := s.updateIp(&updateIpRequest{
		Zone:    req.Zone,
		IpId:    req.IPID,
		Reverse: &ptrReverse,
	})
	if err != nil {
		return nil, err
	}

	return &SetReverseForIPResponse{IP: ipResponse.Ip}, nil
}

// DeleteReverseForIPRequest contains the parameters to set the reverse DNS for an IP
type DeleteReverseForIPRequest struct {
	Zone utils.Zone `json:"-"`
	IPID string     `json:"-"`
}

// DeleteReverseForIPResponse contains the updated IP after setting the reverse
type DeleteReverseForIPResponse struct {
	IP *Ip
}

// DeleteReverseForIP removes the reverse DNS of an IP.
func (s *Api) DeleteReverseForIP(req *DeleteReverseForIPRequest, opts ...scw.RequestOption) (*DeleteReverseForIPResponse, error) {
	var ptrReverse *string
	ipResponse, err := s.updateIp(&updateIpRequest{
		Zone:    req.Zone,
		IpId:    req.IPID,
		Reverse: &ptrReverse,
	})
	if err != nil {
		return nil, err
	}

	return &DeleteReverseForIPResponse{IP: ipResponse.Ip}, nil
}

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

// AttachIPResponse contains the updated IP after attaching
type AttachIPResponse struct {
	IP *IP
}

// AttachIP attaches an IP to a server.
func (s *API) AttachIP(req *AttachIPRequest, opts ...scw.RequestOption) (*AttachIPResponse, error) {
	var ptrServerID = &req.ServerID
	ipResponse, err := s.updateIP(&updateIPRequest{
		Zone:   req.Zone,
		IPID:   req.IPID,
		Server: &ptrServerID,
	})
	if err != nil {
		return nil, err
	}

	return &AttachIPResponse{IP: ipResponse.IP}, nil
}

// DetachIPRequest contains the parameters to detach an IP from a server
type DetachIPRequest struct {
	Zone utils.Zone `json:"-"`
	IPID string     `json:"-"`
}

// DetachIPResponse contains the updated IP after detaching
type DetachIPResponse struct {
	IP *IP
}

// DetachIP detaches an IP from a server.
func (s *API) DetachIP(req *DetachIPRequest, opts ...scw.RequestOption) (*DetachIPResponse, error) {
	var ptrServerID *string
	ipResponse, err := s.updateIP(&updateIPRequest{
		Zone:   req.Zone,
		IPID:   req.IPID,
		Server: &ptrServerID,
	})
	if err != nil {
		return nil, err
	}

	return &DetachIPResponse{IP: ipResponse.IP}, nil
}

// UpdateIPRequest contains the parameters to update an IP
// if Reverse is an empty string, the reverse will be removed
type UpdateIPRequest struct {
	Zone    utils.Zone `json:"-"`
	IPID    string     `json:"-"`
	Reverse *string    `json:"reverse"`
}

// UpdateIP updates an IP
func (s *API) UpdateIP(req *UpdateIPRequest, opts ...scw.RequestOption) (*UpdateIPResponse, error) {
	var reverse **string
	if req.Reverse != nil {
		if *req.Reverse == "" {
			req.Reverse = nil
		}
		reverse = &req.Reverse
	}
	ipResponse, err := s.updateIP(&updateIPRequest{
		Zone:    req.Zone,
		IPID:    req.IPID,
		Reverse: reverse,
	})
	if err != nil {
		return nil, err
	}

	return &UpdateIPResponse{IP: ipResponse.IP}, nil
}

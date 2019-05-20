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

// SetReverseForIPRequest contains the parameters to set the reverse DNS for an IP
type SetReverseForIPRequest struct {
	Zone    utils.Zone `json:"-"`
	IPID    string     `json:"-"`
	Reverse string     `json:"reverse"`
}

// SetReverseForIPResponse contains the updated IP after setting the reverse
type SetReverseForIPResponse struct {
	IP *IP
}

// SetReverseForIP sets the reverse DNS for an IP.
func (s *API) SetReverseForIP(req *SetReverseForIPRequest, opts ...scw.RequestOption) (*SetReverseForIPResponse, error) {
	var ptrReverse = &req.Reverse
	ipResponse, err := s.updateIP(&updateIPRequest{
		Zone:    req.Zone,
		IPID:    req.IPID,
		Reverse: &ptrReverse,
	})
	if err != nil {
		return nil, err
	}

	return &SetReverseForIPResponse{IP: ipResponse.IP}, nil
}

// DeleteReverseForIPRequest contains the parameters to set the reverse DNS for an IP
type DeleteReverseForIPRequest struct {
	Zone utils.Zone `json:"-"`
	IPID string     `json:"-"`
}

// DeleteReverseForIPResponse contains the updated IP after setting the reverse
type DeleteReverseForIPResponse struct {
	IP *IP
}

// DeleteReverseForIP removes the reverse DNS of an IP.
func (s *API) DeleteReverseForIP(req *DeleteReverseForIPRequest, opts ...scw.RequestOption) (*DeleteReverseForIPResponse, error) {
	var ptrReverse *string
	ipResponse, err := s.updateIP(&updateIPRequest{
		Zone:    req.Zone,
		IPID:    req.IPID,
		Reverse: &ptrReverse,
	})
	if err != nil {
		return nil, err
	}

	return &DeleteReverseForIPResponse{IP: ipResponse.IP}, nil
}

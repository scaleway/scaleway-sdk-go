// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package abuse provides methods and message types of the abuse v1alpha2 API.
package abuse

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

type ComplaintClass string

const (
	// Unknown class.
	ComplaintClassUnknownClass = ComplaintClass("unknown_class")
	// The family of complaint type is activity.
	ComplaintClassActivity = ComplaintClass("activity")
	// The family of complaint type is content.
	ComplaintClassContent = ComplaintClass("content")
)

func (enum ComplaintClass) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_class"
	}
	return string(enum)
}

func (enum ComplaintClass) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ComplaintClass) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ComplaintClass(ComplaintClass(tmp).String())
	return nil
}

type ComplaintStatus string

const (
	// Unknown status.
	ComplaintStatusUnknownStatus = ComplaintStatus("unknown_status")
	// Complaint has just been created.
	ComplaintStatusNew = ComplaintStatus("new")
	// Complaint has been forwarded to the workflow.
	ComplaintStatusForwarded = ComplaintStatus("forwarded")
	// Complaint is being processed by the workflow.
	ComplaintStatusProcessing = ComplaintStatus("processing")
	// This complaint is not supported.
	ComplaintStatusNotSupported = ComplaintStatus("not_supported")
	// Error has been received while processing the complaint.
	ComplaintStatusError = ComplaintStatus("error")
	// The resource doesn't belong to our AS.
	ComplaintStatusNotInAs = ComplaintStatus("not_in_as")
)

func (enum ComplaintStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum ComplaintStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ComplaintStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ComplaintStatus(ComplaintStatus(tmp).String())
	return nil
}

type ComplaintType string

const (
	// Unknown type.
	ComplaintTypeUnknownType = ComplaintType("unknown_type")
	// Complaint is about bruteforce.
	ComplaintTypeBruteforce = ComplaintType("bruteforce")
	// Complaint is about botnet.
	ComplaintTypeBotnet = ComplaintType("botnet")
	// Complaint is about copyright.
	ComplaintTypeCopyright = ComplaintType("copyright")
	// Complaint is about ddos.
	ComplaintTypeDdos = ComplaintType("ddos")
	// Complaint is about illicit.
	ComplaintTypeIllicit = ComplaintType("illicit")
	// Complaint is about malware.
	ComplaintTypeMalware = ComplaintType("malware")
	// Complaint is about phishing.
	ComplaintTypePhishing = ComplaintType("phishing")
	// Complaint is about spam.
	ComplaintTypeSpam = ComplaintType("spam")
	// Complaint is about virus.
	ComplaintTypeVirus = ComplaintType("virus")
	// Complaint is about iptv.
	ComplaintTypeIptv = ComplaintType("iptv")
)

func (enum ComplaintType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum ComplaintType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ComplaintType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ComplaintType(ComplaintType(tmp).String())
	return nil
}

// Complaint: complaint.
type Complaint struct {
	// ID: iD.
	ID string `json:"id"`

	// Evidence: content of the complaint explaining why the complaint is made.
	Evidence *string `json:"evidence"`

	// ObservedAt: datetime at which the infringement as been seen.
	ObservedAt *time.Time `json:"observed_at"`

	// ObserverEmail: email of the person or the entity creating the complaint.
	ObserverEmail string `json:"observer_email"`

	// ObserverName: optional name of the person or the entity creating the complaint (optional in case of CSAM).
	ObserverName *string `json:"observer_name"`

	// OffendingIP: IP reported by the complaint.
	OffendingIP *string `json:"offending_ip"`

	// OffendingURL: URL reported by the complaint.
	OffendingURL *string `json:"offending_url"`

	// Status: status.
	// Default value: unknown_status
	Status ComplaintStatus `json:"status"`

	// Type: type of the complaint.
	// Default value: unknown_type
	Type ComplaintType `json:"type"`

	// Class: class of the complaint.
	// Default value: unknown_class
	Class ComplaintClass `json:"class"`

	// ReportedAt: datetime at which the infringement as been reported.
	ReportedAt *time.Time `json:"reported_at"`

	// RecipientEmail: email on which we received the complaint.
	RecipientEmail string `json:"recipient_email"`

	// ParserResultID: ID of the parser result object.
	ParserResultID string `json:"parser_result_id"`
}

// ComplaintAPICreateComplaintRequest: complaint api create complaint request.
type ComplaintAPICreateComplaintRequest struct {
	// Evidence: content of the complaint explaining why the complaint is made.
	Evidence *string `json:"evidence,omitempty"`

	// ObservedAt: datetime at which the infringement as been seen.
	ObservedAt *time.Time `json:"observed_at,omitempty"`

	// ReportedAt: datetime at which the infringement as been reported.
	ReportedAt *time.Time `json:"reported_at,omitempty"`

	// ObserverEmail: email of the person or the entity creating the complaint.
	ObserverEmail string `json:"observer_email"`

	// ObserverName: optional name of the person or the entity creating the complaint (optional in case of CSAM).
	ObserverName *string `json:"observer_name,omitempty"`

	// OffendingIP: IP reported by the complaint.
	OffendingIP *string `json:"offending_ip,omitempty"`

	// OffendingURL: URL reported by the complaint.
	OffendingURL *string `json:"offending_url,omitempty"`

	// Type: type of the complaint.
	// Default value: unknown_type
	Type ComplaintType `json:"type"`

	// Class: class of the complaint.
	// Default value: unknown_class
	Class ComplaintClass `json:"class"`

	// RecipientEmail: email on which we received the complaint.
	RecipientEmail string `json:"recipient_email"`

	// ParserResultID: ID of the parser result object.
	ParserResultID string `json:"parser_result_id"`
}

// This API is used to create complaints internally.
type ComplaintAPI struct {
	client *scw.Client
}

// NewComplaintAPI returns a ComplaintAPI object from a Scaleway client.
func NewComplaintAPI(client *scw.Client) *ComplaintAPI {
	return &ComplaintAPI{
		client: client,
	}
}

// CreateComplaint: Create complaint.
func (s *ComplaintAPI) CreateComplaint(req *ComplaintAPICreateComplaintRequest, opts ...scw.RequestOption) (*Complaint, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/abuse_internal/v1alpha2/complaints",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Complaint

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

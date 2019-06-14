package metadata

import (
	"encoding/json"
	"net/http"

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

var (
	metadataURL = "http://169.254.42.42"
)

// API metadata API
type API struct {
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI() *API {
	return &API{}
}

// GetMetadata returns the metadata avialable from the server
func (*API) GetMetadata() (*Metadata, scw.SdkError) {
	resp, err := http.Get(metadataURL + "/conf?format=json")
	if err != nil {
		return nil, errors.New(err.Error())
	}

	metadata := &Metadata{}
	err = json.NewDecoder(resp.Body).Decode(metadata)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return metadata, nil
}

// Metadata represents the struct return by the metadata API
type Metadata struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Hostname       string `json:"hostname,omitempty"`
	Organization   string `json:"organization,omitempty"`
	CommercialType string `json:"commercial_type,omitempty"`
	PublicIP       struct {
		Dynamic bool   `json:"dynamic,omitempty"`
		ID      string `json:"id,omitempty"`
		Address string `json:"address,omitempty"`
	} `json:"public_ip,omitempty"`
	PrivateIP string `json:"private_ip,omitempty"`
	IPv6      struct {
		Netmask string `json:"netmask,omitempty"`
		Gateway string `json:"gateway,omitempty"`
		Address string `json:"address,omitempty"`
	} `json:"ipv6,omitempty"`
	Location struct {
		PlatformID   string `json:"platform_id,omitempty"`
		HypervisorID string `json:"hypervisor_id,omitempty"`
		NodeID       string `json:"node_id,omitempty"`
		ClusterID    string `json:"cluster_id,omitempty"`
		ZoneID       string `json:"zone_id,omitempty"`
	} `json:"location,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	StateDetail   string   `json:"state_detail,omitempty"`
	SSHPublicKeys []struct {
		Description      string `json:"description,omitempty"`
		ModificationDate string `json:"modification_date,omitempty"`
		IP               string `json:"ip,omitempty"`
		Email            string `json:"email,omitempty"`
		UserAgent        struct {
			Platform string `json:"platform,omitempty"`
			Version  string `json:"version,omitempty"`
			String   string `json:"string,omitempty"`
			Browser  string `json:"browser,omitempty"`
		} `json:"user_agent,omitempty"`
		Key          string `json:"key,omitempty"`
		Fingerprint  string `json:"fingerprint,omitempty"`
		ID           string `json:"id,omitempty"`
		CreationDate string `json:"creation_date,omitempty"`
		Port         int    `json:"port,omitempty"`
	} `json:"ssh_public_keys,omitempty"`
	Timezone   string `json:"timezone,omitempty"`
	Bootscript struct {
		Kernel       string `json:"kernel,omitempty"`
		Title        string `json:"title,omitempty"`
		Default      bool   `json:"default,omitempty"`
		Dtb          string `json:"dtb,omitempty"`
		Public       bool   `json:"publc,omitempty"`
		Initrd       string `json:"initrd,omitempty"`
		Bootcmdargs  string `json:"bootcmdargs,omitempty"`
		Architecture string `json:"architecture,omitempty"`
		Organization string `json:"organization,omitempty"`
		ID           string `json:"id,omitempty"`
	} `json:"bootscript,omitempty"`
	Volumes map[string]struct {
		Name             string `json:"name,omitempty"`
		ModificationDate string `json:"modification_date,omitempty"`
		ExportURI        string `json:"export_uri,omitempty"`
		VolumeType       string `json:"volume_type,omitempty"`
		CreationDate     string `json:"creation_date,omitempty"`
		State            string `json:"state,omitempty"`
		Organization     string `json:"organization,omitempty"`
		Server           struct {
			ID   string `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"server,omitempty"`
		ID   string `json:"id,omitempty"`
		Size int    `json:"size,omitempty"`
	}
}

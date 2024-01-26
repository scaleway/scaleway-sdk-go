// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package resource provides methods and message types of the resource v1alpha1 API.
package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

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

type ResourceLocality string

const (
	ResourceLocalityUnknownLocality = ResourceLocality("unknown_locality")
	ResourceLocalityGlobal          = ResourceLocality("global")
	ResourceLocalityFrRz            = ResourceLocality("fr_rz")
	ResourceLocalityFrSrr           = ResourceLocality("fr_srr")
	ResourceLocalityFrSrr1          = ResourceLocality("fr_srr_1")
	ResourceLocalityFrPar           = ResourceLocality("fr_par")
	ResourceLocalityFrPar1          = ResourceLocality("fr_par_1")
	ResourceLocalityFrPar2          = ResourceLocality("fr_par_2")
	ResourceLocalityFrPar3          = ResourceLocality("fr_par_3")
	ResourceLocalityNlAms           = ResourceLocality("nl_ams")
	ResourceLocalityNlAms1          = ResourceLocality("nl_ams_1")
	ResourceLocalityNlAms2          = ResourceLocality("nl_ams_2")
	ResourceLocalityNlAms3          = ResourceLocality("nl_ams_3")
	ResourceLocalityPlWaw           = ResourceLocality("pl_waw")
	ResourceLocalityPlWaw1          = ResourceLocality("pl_waw_1")
	ResourceLocalityPlWaw2          = ResourceLocality("pl_waw_2")
	ResourceLocalityPlWaw3          = ResourceLocality("pl_waw_3")
	ResourceLocalityFrInt           = ResourceLocality("fr_int")
	ResourceLocalityFrInt1          = ResourceLocality("fr_int_1")
	ResourceLocalityFrLab           = ResourceLocality("fr_lab")
	ResourceLocalityFrLab1          = ResourceLocality("fr_lab_1")
)

func (enum ResourceLocality) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_locality"
	}
	return string(enum)
}

func (enum ResourceLocality) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ResourceLocality) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ResourceLocality(ResourceLocality(tmp).String())
	return nil
}

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package product_catalog_admin provides methods and message types of the product_catalog_admin v1alpha1 API.
package product_catalog_admin

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

// ExportCatalogCSVRequest: export catalog csv request.
type ExportCatalogCSVRequest struct {
	// LastExportDate: this is an optional date which is the last export date of the catalog.
	LastExportDate *time.Time `json:"-"`
}

// Product Catalog Admin API endpoints.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// ExportCatalogCSV:
func (s *API) ExportCatalogCSV(req *ExportCatalogCSVRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "last_export_date", req.LastExportDate)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/product-catalog-admin/v1alpha1/export-csv",
		Query:  query,
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

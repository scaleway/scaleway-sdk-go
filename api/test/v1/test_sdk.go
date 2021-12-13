// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package test provides methods and message types of the test v1 API.
package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/marshaler"
	"github.com/scaleway/scaleway-sdk-go/internal/parameter"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/scw"
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

// API: this API allows us to test
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// NoAuthAPI: no Auth Service
type NoAuthAPI struct {
	client *scw.Client
}

// NewNoAuthAPI returns a NoAuthAPI object from a Scaleway client.
func NewNoAuthAPI(client *scw.Client) *NoAuthAPI {
	return &NoAuthAPI{
		client: client,
	}
}

// RegionalAPI: regional API
type RegionalAPI struct {
	client *scw.Client
}

// NewRegionalAPI returns a RegionalAPI object from a Scaleway client.
func NewRegionalAPI(client *scw.Client) *RegionalAPI {
	return &RegionalAPI{
		client: client,
	}
}

// StreamAPI: stream Service
type StreamAPI struct {
	client *scw.Client
}

// NewStreamAPI returns a StreamAPI object from a Scaleway client.
func NewStreamAPI(client *scw.Client) *StreamAPI {
	return &StreamAPI{
		client: client,
	}
}

// ZoneAPI: zone Test API
type ZoneAPI struct {
	client *scw.Client
}

// NewZoneAPI returns a ZoneAPI object from a Scaleway client.
func NewZoneAPI(client *scw.Client) *ZoneAPI {
	return &ZoneAPI{
		client: client,
	}
}

type ListCharactersRequestOrderBy string

const (
	// ListCharactersRequestOrderByAsc is [insert doc].
	ListCharactersRequestOrderByAsc = ListCharactersRequestOrderBy("asc")
	// ListCharactersRequestOrderByDesc is [insert doc].
	ListCharactersRequestOrderByDesc = ListCharactersRequestOrderBy("desc")
)

func (enum ListCharactersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "asc"
	}
	return string(enum)
}

func (enum ListCharactersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListCharactersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListCharactersRequestOrderBy(ListCharactersRequestOrderBy(tmp).String())
	return nil
}

type MapEnum string

const (
	// MapEnumMAPENUMFOO is [insert doc].
	MapEnumMAPENUMFOO = MapEnum("MAP_ENUM_FOO")
	// MapEnumMAPENUMBAR is [insert doc].
	MapEnumMAPENUMBAR = MapEnum("MAP_ENUM_BAR")
	// MapEnumMAPENUMBAZ is [insert doc].
	MapEnumMAPENUMBAZ = MapEnum("MAP_ENUM_BAZ")
)

func (enum MapEnum) String() string {
	if enum == "" {
		// return default value if empty
		return "MAP_ENUM_FOO"
	}
	return string(enum)
}

func (enum MapEnum) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MapEnum) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MapEnum(MapEnum(tmp).String())
	return nil
}

type NullValue string

const (
	// NullValueNULLVALUE is [insert doc].
	NullValueNULLVALUE = NullValue("NULL_VALUE")
)

func (enum NullValue) String() string {
	if enum == "" {
		// return default value if empty
		return "NULL_VALUE"
	}
	return string(enum)
}

func (enum NullValue) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NullValue) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NullValue(NullValue(tmp).String())
	return nil
}

type PostAllTypesRequestNestedEnum string

const (
	// PostAllTypesRequestNestedEnumNEG is [insert doc].
	PostAllTypesRequestNestedEnumNEG = PostAllTypesRequestNestedEnum("NEG")
	// PostAllTypesRequestNestedEnumZERO is [insert doc].
	PostAllTypesRequestNestedEnumZERO = PostAllTypesRequestNestedEnum("ZERO")
	// PostAllTypesRequestNestedEnumFOO is [insert doc].
	PostAllTypesRequestNestedEnumFOO = PostAllTypesRequestNestedEnum("FOO")
	// PostAllTypesRequestNestedEnumBAR is [insert doc].
	PostAllTypesRequestNestedEnumBAR = PostAllTypesRequestNestedEnum("BAR")
	// PostAllTypesRequestNestedEnumBAZ is [insert doc].
	PostAllTypesRequestNestedEnumBAZ = PostAllTypesRequestNestedEnum("BAZ")
)

func (enum PostAllTypesRequestNestedEnum) String() string {
	if enum == "" {
		// return default value if empty
		return "ZERO"
	}
	return string(enum)
}

func (enum PostAllTypesRequestNestedEnum) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PostAllTypesRequestNestedEnum) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PostAllTypesRequestNestedEnum(PostAllTypesRequestNestedEnum(tmp).String())
	return nil
}

type PostAllTypesResponseNestedEnum string

const (
	// PostAllTypesResponseNestedEnumNEG is [insert doc].
	PostAllTypesResponseNestedEnumNEG = PostAllTypesResponseNestedEnum("NEG")
	// PostAllTypesResponseNestedEnumZERO is [insert doc].
	PostAllTypesResponseNestedEnumZERO = PostAllTypesResponseNestedEnum("ZERO")
	// PostAllTypesResponseNestedEnumFOO is [insert doc].
	PostAllTypesResponseNestedEnumFOO = PostAllTypesResponseNestedEnum("FOO")
	// PostAllTypesResponseNestedEnumBAR is [insert doc].
	PostAllTypesResponseNestedEnumBAR = PostAllTypesResponseNestedEnum("BAR")
	// PostAllTypesResponseNestedEnumBAZ is [insert doc].
	PostAllTypesResponseNestedEnumBAZ = PostAllTypesResponseNestedEnum("BAZ")
)

func (enum PostAllTypesResponseNestedEnum) String() string {
	if enum == "" {
		// return default value if empty
		return "ZERO"
	}
	return string(enum)
}

func (enum PostAllTypesResponseNestedEnum) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PostAllTypesResponseNestedEnum) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PostAllTypesResponseNestedEnum(PostAllTypesResponseNestedEnum(tmp).String())
	return nil
}

// Type: this is an enum type
type Type string

const (
	// TypeUnknown is [insert doc].
	TypeUnknown = Type("unknown")
	// TypeStart1Xs is [insert doc].
	TypeStart1Xs = Type("start1_xs")
	// TypeStart1S is [insert doc].
	TypeStart1S = Type("start1_s")
	// TypeStart1M is [insert doc].
	TypeStart1M = Type("start1_m")
)

func (enum Type) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum Type) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *Type) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = Type(Type(tmp).String())
	return nil
}

// EchoMessage: echo message
type EchoMessage struct {
	Str string `json:"str"`

	Strs []string `json:"strs"`
}

type GetEnumResponse struct {
	// Type:
	//
	// Default value: unknown
	Type Type `json:"type"`
}

type GetZoneResponse struct {
	Zone scw.Zone `json:"zone"`
}

// ListCharactersResponse: list characters response
type ListCharactersResponse struct {
	TotalCount uint32 `json:"total_count"`

	Characters []string `json:"characters"`
}

// MetadataResponse: metadata response
type MetadataResponse struct {
	Metadata map[string]string `json:"metadata"`
}

type PatchEnumResponse struct {
	// Type:
	//
	// Default value: unknown
	Type Type `json:"type"`
}

type PostAllMapTypesResponse struct {
	MapInt32Int32 map[int32]int32 `json:"map_int32_int32"`

	MapInt64Int64 map[int64]int64 `json:"map_int64_int64"`

	MapUint32Uint32 map[uint32]uint32 `json:"map_uint32_uint32"`

	MapUint64Uint64 map[uint64]uint64 `json:"map_uint64_uint64"`

	MapSint32Sint32 map[int32]int32 `json:"map_sint32_sint32"`

	MapSint64Sint64 map[int64]int64 `json:"map_sint64_sint64"`

	MapFixed32Fixed32 map[uint32]uint32 `json:"map_fixed32_fixed32"`

	MapFixed64Fixed64 map[uint64]uint64 `json:"map_fixed64_fixed64"`

	MapSfixed32Sfixed32 map[int32]int32 `json:"map_sfixed32_sfixed32"`

	MapSfixed64Sfixed64 map[int64]int64 `json:"map_sfixed64_sfixed64"`

	MapInt32Float map[int32]float32 `json:"map_int32_float"`

	MapInt32Double map[int32]float64 `json:"map_int32_double"`

	MapStringString map[string]string `json:"map_string_string"`

	MapInt32Bytes map[int32][]byte `json:"map_int32_bytes"`

	MapInt32Enum map[int32]MapEnum `json:"map_int32_enum"`

	MapInt32AllTypes map[int32]*PostAllTypesRequest `json:"map_int32_all_types"`

	MapInt32IP map[int32]net.IP `json:"map_int32_ip"`

	MapInt32StdDuration map[int32]*time.Duration `json:"map_int32_std_duration"`

	MapInt32StdLongDuration map[int32]*time.Duration `json:"map_int32_std_long_duration"`

	MapInt32Size map[int32]scw.Size `json:"map_int32_size"`

	MapInt32Uint64Size map[int32]scw.Size `json:"map_int32_uint64_size"`

	MapInt32Uint64valueSize map[int32]*scw.Size `json:"map_int32_uint64value_size"`

	MapInt32StringIP map[int32]net.IP `json:"map_int32_string_ip"`

	MapInt32StringValueIP map[int32]*net.IP `json:"map_int32_string_value_ip"`

	MapInt32IPv4 map[int32]net.IP `json:"map_int32_ipv4"`

	MapInt32StringIPv4 map[int32]net.IP `json:"map_int32_string_ipv4"`

	MapInt32StringValueIPv4 map[int32]*net.IP `json:"map_int32_string_value_ipv4"`

	MapInt32IPv6 map[int32]net.IP `json:"map_int32_ipv6"`

	MapInt32StringIPv6 map[int32]net.IP `json:"map_int32_string_ipv6"`

	MapInt32StringValueIPv6 map[int32]*net.IP `json:"map_int32_string_value_ipv6"`

	MapInt32StringsValue map[int32]*[]string `json:"map_int32_strings_value"`

	MapInt32Duration map[int32]*scw.Duration `json:"map_int32_duration"`
}

func (m *PostAllMapTypesResponse) UnmarshalJSON(b []byte) error {
	type tmpType PostAllMapTypesResponse
	tmp := struct {
		tmpType

		TmpMapInt32StdDuration     marshaler.Durationint32Map     `json:"map_int32_std_duration"`
		TmpMapInt32StdLongDuration marshaler.LongDurationint32Map `json:"map_int32_std_long_duration"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = PostAllMapTypesResponse(tmp.tmpType)

	m.MapInt32StdDuration = tmp.TmpMapInt32StdDuration.Standard()
	m.MapInt32StdLongDuration = tmp.TmpMapInt32StdLongDuration.Standard()
	return nil
}

func (m PostAllMapTypesResponse) MarshalJSON() ([]byte, error) {
	type tmpType PostAllMapTypesResponse
	tmp := struct {
		tmpType

		TmpMapInt32StdDuration     marshaler.Durationint32Map     `json:"map_int32_std_duration"`
		TmpMapInt32StdLongDuration marshaler.LongDurationint32Map `json:"map_int32_std_long_duration"`
	}{
		tmpType: tmpType(m),

		TmpMapInt32StdDuration:     marshaler.NewDurationint32Map(m.MapInt32StdDuration),
		TmpMapInt32StdLongDuration: marshaler.NewLongDurationint32Map(m.MapInt32StdLongDuration),
	}
	return json.Marshal(tmp)
}

// PostAllTypesRequest: post all types request
type PostAllTypesRequest struct {
	SingularInt32 int32 `json:"singular_int32"`

	SingularInt64 int64 `json:"singular_int64"`

	SingularUint32 uint32 `json:"singular_uint32"`

	SingularUint64 uint64 `json:"singular_uint64"`

	SingularSint32 int32 `json:"singular_sint32"`

	SingularSint64 int64 `json:"singular_sint64"`

	SingularFixed32 uint32 `json:"singular_fixed32"`

	SingularFixed64 uint64 `json:"singular_fixed64"`

	SingularSfixed32 int32 `json:"singular_sfixed32"`

	SingularSfixed64 int64 `json:"singular_sfixed64"`

	SingularFloat float32 `json:"singular_float"`

	SingularDouble float64 `json:"singular_double"`

	SingularBool bool `json:"singular_bool"`

	SingularString string `json:"singular_string"`

	SingularBytes []byte `json:"singular_bytes"`

	SingularNestedMessage *PostAllTypesRequestNestedMessage `json:"singular_nested_message"`
	// SingularNestedEnum:
	//
	// Default value: ZERO
	SingularNestedEnum PostAllTypesRequestNestedEnum `json:"singular_nested_enum"`

	RepeatedInt32 []int32 `json:"repeated_int32"`

	RepeatedInt64 []int64 `json:"repeated_int64"`

	RepeatedUint32 []uint32 `json:"repeated_uint32"`

	RepeatedUint64 []uint64 `json:"repeated_uint64"`

	RepeatedSint32 []int32 `json:"repeated_sint32"`

	RepeatedSint64 []int64 `json:"repeated_sint64"`

	RepeatedFixed32 []uint32 `json:"repeated_fixed32"`

	RepeatedFixed64 []uint64 `json:"repeated_fixed64"`

	RepeatedSfixed32 []int32 `json:"repeated_sfixed32"`

	RepeatedSfixed64 []int64 `json:"repeated_sfixed64"`

	RepeatedFloat []float32 `json:"repeated_float"`

	RepeatedDouble []float64 `json:"repeated_double"`

	RepeatedBool []bool `json:"repeated_bool"`

	RepeatedString []string `json:"repeated_string"`

	RepeatedBytes [][]byte `json:"repeated_bytes"`

	RepeatedNestedMessage []*PostAllTypesRequestNestedMessage `json:"repeated_nested_message"`

	RepeatedNestedEnum []PostAllTypesRequestNestedEnum `json:"repeated_nested_enum"`

	// Precisely one of OneofBytes, OneofNestedMessage, OneofString, OneofUint32 must be set.
	OneofUint32 *uint32 `json:"oneof_uint32,omitempty"`

	// Precisely one of OneofBytes, OneofNestedMessage, OneofString, OneofUint32 must be set.
	OneofNestedMessage *PostAllTypesRequestNestedMessage `json:"oneof_nested_message,omitempty"`

	// Precisely one of OneofBytes, OneofNestedMessage, OneofString, OneofUint32 must be set.
	OneofString *string `json:"oneof_string,omitempty"`

	// Precisely one of OneofBytes, OneofNestedMessage, OneofString, OneofUint32 must be set.
	OneofBytes *[]byte `json:"oneof_bytes,omitempty"`

	SingularDoubleValue *float64 `json:"singular_double_value"`

	SingularFloatValue *float32 `json:"singular_float_value"`

	SingularInt64Value *int64 `json:"singular_int64_value"`

	SingularUint64Value *uint64 `json:"singular_uint64_value"`

	SingularInt32Value *int32 `json:"singular_int32_value"`

	SingularUint32Value *uint32 `json:"singular_uint32_value"`

	SingularBoolValue *bool `json:"singular_bool_value"`

	SingularStringValue *string `json:"singular_string_value"`

	SingularBytesValue *[]byte `json:"singular_bytes_value"`

	SingularTimestamp *time.Time `json:"singular_timestamp"`

	SingularAny interface{} `json:"singular_any"`

	SingularStruct []byte `json:"singular_struct"`

	SingularMoney *scw.Money `json:"singular_money"`

	SingularStringsValue *[]string `json:"singular_strings_value"`

	SingularDuration *scw.Duration `json:"singular_duration"`

	SingularIP net.IP `json:"singular_ip"`

	SingularStringIP net.IP `json:"singular_string_ip"`

	SingularStringValueIP *net.IP `json:"singular_string_value_ip"`

	SingularIPv4 net.IP `json:"singular_ipv4"`

	SingularStringIPv4 net.IP `json:"singular_string_ipv4"`

	SingularStringValueIPv4 *net.IP `json:"singular_string_value_ipv4"`

	SingularIPv6 net.IP `json:"singular_ipv6"`

	SingularStringIPv6 net.IP `json:"singular_string_ipv6"`

	SingularStringValueIPv6 *net.IP `json:"singular_string_value_ipv6"`

	SingularStdDuration *time.Duration `json:"singular_std_duration"`

	SingularStdLongDuration *time.Duration `json:"singular_std_long_duration"`

	SingularSize scw.Size `json:"singular_size"`

	SingularUint64Size scw.Size `json:"singular_uint64_size"`

	SingularUint64valueSize *scw.Size `json:"singular_uint64value_size"`

	SingularStringIpnet scw.IPNet `json:"singular_string_ipnet"`

	SingularStringValueIpnet *scw.IPNet `json:"singular_string_value_ipnet"`

	RepeatedDoubleValue []*float64 `json:"repeated_double_value"`

	RepeatedFloatValue []*float32 `json:"repeated_float_value"`

	RepeatedInt64Value []*int64 `json:"repeated_int64_value"`

	RepeatedUint64Value []*uint64 `json:"repeated_uint64_value"`

	RepeatedInt32Value []*int32 `json:"repeated_int32_value"`

	RepeatedUint32Value []*uint32 `json:"repeated_uint32_value"`

	RepeatedBoolValue []*bool `json:"repeated_bool_value"`

	RepeatedStringValue []*string `json:"repeated_string_value"`

	RepeatedBytesValue []*[]byte `json:"repeated_bytes_value"`

	RepeatedTimestamp []*time.Time `json:"repeated_timestamp"`

	RepeatedAny []interface{} `json:"repeated_any"`

	RepeatedStruct [][]byte `json:"repeated_struct"`

	RepeatedMoney []*scw.Money `json:"repeated_money"`

	RepeatedStringsValue []*[]string `json:"repeated_strings_value"`

	RepeatedDuration []*scw.Duration `json:"repeated_duration"`

	RepeatedIP []net.IP `json:"repeated_ip"`

	RepeatedStringIP []net.IP `json:"repeated_string_ip"`

	RepeatedStringValueIP []*net.IP `json:"repeated_string_value_ip"`

	RepeatedIPv4 []net.IP `json:"repeated_ipv4"`

	RepeatedStringIPv4 []net.IP `json:"repeated_string_ipv4"`

	RepeatedStringValueIPv4 []*net.IP `json:"repeated_string_value_ipv4"`

	RepeatedIPv6 []net.IP `json:"repeated_ipv6"`

	RepeatedStringIPv6 []net.IP `json:"repeated_string_ipv6"`

	RepeatedStringValueIPv6 []*net.IP `json:"repeated_string_value_ipv6"`

	RepeatedStdDuration []*time.Duration `json:"repeated_std_duration"`

	RepeatedStdLongDuration []*time.Duration `json:"repeated_std_long_duration"`

	RepeatedSize []scw.Size `json:"repeated_size"`

	RepeatedUint64Size []scw.Size `json:"repeated_uint64_size"`

	RepeatedUint64valueSize []*scw.Size `json:"repeated_uint64value_size"`

	RepeatedStringIpnet []scw.IPNet `json:"repeated_string_ipnet"`

	RepeatedStringValueIpnet []*scw.IPNet `json:"repeated_string_value_ipnet"`
}

func (m *PostAllTypesRequest) UnmarshalJSON(b []byte) error {
	type tmpType PostAllTypesRequest
	tmp := struct {
		tmpType

		TmpSingularStdDuration     *marshaler.Duration         `json:"singular_std_duration"`
		TmpSingularStdLongDuration *marshaler.LongDuration     `json:"singular_std_long_duration"`
		TmpRepeatedStdDuration     marshaler.DurationSlice     `json:"repeated_std_duration"`
		TmpRepeatedStdLongDuration marshaler.LongDurationSlice `json:"repeated_std_long_duration"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = PostAllTypesRequest(tmp.tmpType)

	m.SingularStdDuration = tmp.TmpSingularStdDuration.Standard()
	m.SingularStdLongDuration = tmp.TmpSingularStdLongDuration.Standard()
	m.RepeatedStdDuration = tmp.TmpRepeatedStdDuration.Standard()
	m.RepeatedStdLongDuration = tmp.TmpRepeatedStdLongDuration.Standard()
	return nil
}

func (m PostAllTypesRequest) MarshalJSON() ([]byte, error) {
	type tmpType PostAllTypesRequest
	tmp := struct {
		tmpType

		TmpSingularStdDuration     *marshaler.Duration         `json:"singular_std_duration"`
		TmpSingularStdLongDuration *marshaler.LongDuration     `json:"singular_std_long_duration"`
		TmpRepeatedStdDuration     marshaler.DurationSlice     `json:"repeated_std_duration"`
		TmpRepeatedStdLongDuration marshaler.LongDurationSlice `json:"repeated_std_long_duration"`
	}{
		tmpType: tmpType(m),

		TmpSingularStdDuration:     marshaler.NewDuration(m.SingularStdDuration),
		TmpSingularStdLongDuration: marshaler.NewLongDuration(m.SingularStdLongDuration),
		TmpRepeatedStdDuration:     marshaler.NewDurationSlice(m.RepeatedStdDuration),
		TmpRepeatedStdLongDuration: marshaler.NewLongDurationSlice(m.RepeatedStdLongDuration),
	}
	return json.Marshal(tmp)
}

type PostAllTypesRequestNestedMessage struct {
	Bb int32 `json:"bb"`
}

type PostAllTypesResponse struct {
	SingularInt32 int32 `json:"singular_int32"`

	SingularInt64 int64 `json:"singular_int64"`

	SingularUint32 uint32 `json:"singular_uint32"`

	SingularUint64 uint64 `json:"singular_uint64"`

	SingularSint32 int32 `json:"singular_sint32"`

	SingularSint64 int64 `json:"singular_sint64"`

	SingularFixed32 uint32 `json:"singular_fixed32"`

	SingularFixed64 uint64 `json:"singular_fixed64"`

	SingularSfixed32 int32 `json:"singular_sfixed32"`

	SingularSfixed64 int64 `json:"singular_sfixed64"`

	SingularFloat float32 `json:"singular_float"`

	SingularDouble float64 `json:"singular_double"`

	SingularBool bool `json:"singular_bool"`

	SingularString string `json:"singular_string"`

	SingularBytes []byte `json:"singular_bytes"`

	SingularNestedMessage *PostAllTypesResponseNestedMessage `json:"singular_nested_message"`
	// SingularNestedEnum:
	//
	// Default value: ZERO
	SingularNestedEnum PostAllTypesResponseNestedEnum `json:"singular_nested_enum"`

	RepeatedInt32 []int32 `json:"repeated_int32"`

	RepeatedInt64 []int64 `json:"repeated_int64"`

	RepeatedUint32 []uint32 `json:"repeated_uint32"`

	RepeatedUint64 []uint64 `json:"repeated_uint64"`

	RepeatedSint32 []int32 `json:"repeated_sint32"`

	RepeatedSint64 []int64 `json:"repeated_sint64"`

	RepeatedFixed32 []uint32 `json:"repeated_fixed32"`

	RepeatedFixed64 []uint64 `json:"repeated_fixed64"`

	RepeatedSfixed32 []int32 `json:"repeated_sfixed32"`

	RepeatedSfixed64 []int64 `json:"repeated_sfixed64"`

	RepeatedFloat []float32 `json:"repeated_float"`

	RepeatedDouble []float64 `json:"repeated_double"`

	RepeatedBool []bool `json:"repeated_bool"`

	RepeatedString []string `json:"repeated_string"`

	RepeatedBytes [][]byte `json:"repeated_bytes"`

	RepeatedNestedMessage []*PostAllTypesResponseNestedMessage `json:"repeated_nested_message"`

	RepeatedNestedEnum []PostAllTypesResponseNestedEnum `json:"repeated_nested_enum"`

	// Precisely one of OneofBytes, OneofNestedMessage, OneofString, OneofUint32 must be set.
	OneofUint32 *uint32 `json:"oneof_uint32,omitempty"`

	// Precisely one of OneofBytes, OneofNestedMessage, OneofString, OneofUint32 must be set.
	OneofNestedMessage *PostAllTypesResponseNestedMessage `json:"oneof_nested_message,omitempty"`

	// Precisely one of OneofBytes, OneofNestedMessage, OneofString, OneofUint32 must be set.
	OneofString *string `json:"oneof_string,omitempty"`

	// Precisely one of OneofBytes, OneofNestedMessage, OneofString, OneofUint32 must be set.
	OneofBytes *[]byte `json:"oneof_bytes,omitempty"`

	SingularDoubleValue *float64 `json:"singular_double_value"`

	SingularFloatValue *float32 `json:"singular_float_value"`

	SingularInt64Value *int64 `json:"singular_int64_value"`

	SingularUint64Value *uint64 `json:"singular_uint64_value"`

	SingularInt32Value *int32 `json:"singular_int32_value"`

	SingularUint32Value *uint32 `json:"singular_uint32_value"`

	SingularBoolValue *bool `json:"singular_bool_value"`

	SingularStringValue *string `json:"singular_string_value"`

	SingularBytesValue *[]byte `json:"singular_bytes_value"`

	SingularTimestamp *time.Time `json:"singular_timestamp"`

	SingularAny interface{} `json:"singular_any"`

	SingularStruct []byte `json:"singular_struct"`

	SingularMoney *scw.Money `json:"singular_money"`

	SingularStringsValue *[]string `json:"singular_strings_value"`

	SingularIP net.IP `json:"singular_ip"`

	SingularIPv4 net.IP `json:"singular_ipv4"`

	SingularIPv6 net.IP `json:"singular_ipv6"`

	SingularStdDuration *time.Duration `json:"singular_std_duration"`

	SingularStdLongDuration *time.Duration `json:"singular_std_long_duration"`

	SingularSize scw.Size `json:"singular_size"`

	SingularUint64Size scw.Size `json:"singular_uint64_size"`

	SingularUint64valueSize *scw.Size `json:"singular_uint64value_size"`

	SingularStringIpnet scw.IPNet `json:"singular_string_ipnet"`

	SingularStringValueIpnet *scw.IPNet `json:"singular_string_value_ipnet"`

	SingularDuration *scw.Duration `json:"singular_duration"`

	RepeatedDoubleValue []*float64 `json:"repeated_double_value"`

	RepeatedFloatValue []*float32 `json:"repeated_float_value"`

	RepeatedInt64Value []*int64 `json:"repeated_int64_value"`

	RepeatedUint64Value []*uint64 `json:"repeated_uint64_value"`

	RepeatedInt32Value []*int32 `json:"repeated_int32_value"`

	RepeatedUint32Value []*uint32 `json:"repeated_uint32_value"`

	RepeatedBoolValue []*bool `json:"repeated_bool_value"`

	RepeatedStringValue []*string `json:"repeated_string_value"`

	RepeatedBytesValue []*[]byte `json:"repeated_bytes_value"`

	RepeatedTimestamp []*time.Time `json:"repeated_timestamp"`

	RepeatedAny []interface{} `json:"repeated_any"`

	RepeatedStruct [][]byte `json:"repeated_struct"`

	RepeatedMoney []*scw.Money `json:"repeated_money"`

	RepeatedStringsValue []*[]string `json:"repeated_strings_value"`

	RepeatedIP []net.IP `json:"repeated_ip"`

	RepeatedIPv4 []net.IP `json:"repeated_ipv4"`

	RepeatedIPv6 []net.IP `json:"repeated_ipv6"`

	RepeatedStdDuration []*time.Duration `json:"repeated_std_duration"`

	RepeatedStdLongDuration []*time.Duration `json:"repeated_std_long_duration"`

	RepeatedSize []scw.Size `json:"repeated_size"`

	RepeatedUint64Size []scw.Size `json:"repeated_uint64_size"`

	RepeatedUint64valueSize []*scw.Size `json:"repeated_uint64value_size"`

	SingularStringIP net.IP `json:"singular_string_ip"`

	SingularStringValueIP *net.IP `json:"singular_string_value_ip"`

	SingularStringIPv4 net.IP `json:"singular_string_ipv4"`

	SingularStringValueIPv4 *net.IP `json:"singular_string_value_ipv4"`

	SingularStringIPv6 net.IP `json:"singular_string_ipv6"`

	SingularStringValueIPv6 *net.IP `json:"singular_string_value_ipv6"`

	RepeatedStringIP []net.IP `json:"repeated_string_ip"`

	RepeatedStringValueIP []*net.IP `json:"repeated_string_value_ip"`

	RepeatedStringIPv4 []net.IP `json:"repeated_string_ipv4"`

	RepeatedStringValueIPv4 []*net.IP `json:"repeated_string_value_ipv4"`

	RepeatedStringIPv6 []net.IP `json:"repeated_string_ipv6"`

	RepeatedStringValueIPv6 []*net.IP `json:"repeated_string_value_ipv6"`

	RepeatedStringIpnet []scw.IPNet `json:"repeated_string_ipnet"`

	RepeatedStringValueIpnet []*scw.IPNet `json:"repeated_string_value_ipnet"`

	RepeatedDuration []*scw.Duration `json:"repeated_duration"`
}

func (m *PostAllTypesResponse) UnmarshalJSON(b []byte) error {
	type tmpType PostAllTypesResponse
	tmp := struct {
		tmpType

		TmpSingularStdDuration     *marshaler.Duration         `json:"singular_std_duration"`
		TmpSingularStdLongDuration *marshaler.LongDuration     `json:"singular_std_long_duration"`
		TmpRepeatedStdDuration     marshaler.DurationSlice     `json:"repeated_std_duration"`
		TmpRepeatedStdLongDuration marshaler.LongDurationSlice `json:"repeated_std_long_duration"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = PostAllTypesResponse(tmp.tmpType)

	m.SingularStdDuration = tmp.TmpSingularStdDuration.Standard()
	m.SingularStdLongDuration = tmp.TmpSingularStdLongDuration.Standard()
	m.RepeatedStdDuration = tmp.TmpRepeatedStdDuration.Standard()
	m.RepeatedStdLongDuration = tmp.TmpRepeatedStdLongDuration.Standard()
	return nil
}

func (m PostAllTypesResponse) MarshalJSON() ([]byte, error) {
	type tmpType PostAllTypesResponse
	tmp := struct {
		tmpType

		TmpSingularStdDuration     *marshaler.Duration         `json:"singular_std_duration"`
		TmpSingularStdLongDuration *marshaler.LongDuration     `json:"singular_std_long_duration"`
		TmpRepeatedStdDuration     marshaler.DurationSlice     `json:"repeated_std_duration"`
		TmpRepeatedStdLongDuration marshaler.LongDurationSlice `json:"repeated_std_long_duration"`
	}{
		tmpType: tmpType(m),

		TmpSingularStdDuration:     marshaler.NewDuration(m.SingularStdDuration),
		TmpSingularStdLongDuration: marshaler.NewLongDuration(m.SingularStdLongDuration),
		TmpRepeatedStdDuration:     marshaler.NewDurationSlice(m.RepeatedStdDuration),
		TmpRepeatedStdLongDuration: marshaler.NewLongDurationSlice(m.RepeatedStdLongDuration),
	}
	return json.Marshal(tmp)
}

type PostAllTypesResponseNestedMessage struct {
	Bb int32 `json:"bb"`
}

type PostBodyAndPathAndQueryResponse struct {
	Path string `json:"path"`

	Query string `json:"query"`

	Body *PostBodyAndPathSimpleRequest `json:"body"`
}

type PostBodyAndPathComplexResponse struct {
	Path string `json:"path"`

	Body *PostBodyAndPathSimpleRequest `json:"body"`
}

type PostBodyAndPathSimple2Response struct {
	Path string `json:"path"`

	Body string `json:"body"`
}

// PostBodyAndPathSimpleRequest: post body and path simple request
type PostBodyAndPathSimpleRequest struct {
	Path string `json:"path"`

	Body string `json:"body"`
}

type PostBodyAndPathSimpleResponse struct {
	Path string `json:"path"`

	Body string `json:"body"`
}

type PostDeprecatedOrganizationResponse struct {
	// Deprecated
	Organization *string `json:"organization"`
}

type PostDeprecatedProjectResponse struct {
	// Deprecated
	Project *string `json:"project"`
}

type PostEchoTimeSeriesResponse struct {
	Metrics *scw.TimeSeries `json:"metrics"`
}

type PostEnumResponse struct {
	// Type:
	//
	// Default value: unknown
	Type Type `json:"type"`
}

type PostIPResponse struct {
	IPV4 net.IP `json:"ip_v4"`

	IPV6 net.IP `json:"ip_v6"`

	IP net.IP `json:"ip"`
}

type PostOneOfResponse struct {

	// Precisely one of Test, Test2 must be set.
	Test *string `json:"test,omitempty"`

	// Precisely one of Test, Test2 must be set.
	Test2 *int32 `json:"test2,omitempty"`

	// Precisely one of Test3, TestNested must be set.
	TestNested *EchoMessage `json:"test_nested,omitempty"`

	// Precisely one of Test3, TestNested must be set.
	Test3 *int32 `json:"test3,omitempty"`
}

type PostOrganizationIDResponse struct {
	OrganizationID string `json:"organization_id"`
}

type PostProjectIDResponse struct {
	ProjectID string `json:"project_id"`
}

type PostScalarTypesResponse struct {
	DoubleField float64 `json:"double_field"`

	FloatField float32 `json:"float_field"`

	Int32Field int32 `json:"int32_field"`

	Int64Field int64 `json:"int64_field"`

	Uint32Field uint32 `json:"uint32_field"`

	Uint64Field uint64 `json:"uint64_field"`

	BoolField bool `json:"bool_field"`

	StringField string `json:"string_field"`
}

type PostTagsResponse struct {
	Tags *[]string `json:"tags"`
}

// getRegionResponse: get region response
type getRegionResponse struct {
	Region scw.Region `json:"region"`
}

// Service API

type GetServiceInfoRequest struct {
}

func (s *API) GetServiceInfo(req *GetServiceInfoRequest, opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1",
		Headers: http.Header{},
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteNothingRequest struct {
}

// DeleteNothing: this method deletes nothing
func (s *API) DeleteNothing(req *DeleteNothingRequest, opts ...scw.RequestOption) error {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/test-internal/v1",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type GetEchoRequest struct {
	// Str: a string
	Str string `json:"-"`
	// Strs: a slice of strings
	Strs []string `json:"-"`
}

// GetEcho: echo the request message
//
// ### This is a multiline test.
//
func (s *API) GetEcho(req *GetEchoRequest, opts ...scw.RequestOption) (*EchoMessage, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "str", req.Str)
	parameter.AddToQuery(query, "strs", req.Strs)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/echo",
		Query:   query,
		Headers: http.Header{},
	}

	var resp EchoMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostEchoRequest struct {
	Str *string `json:"str"`

	Strs []string `json:"strs"`
}

// postEcho: echo the request message
func (s *API) postEcho(req *PostEchoRequest, opts ...scw.RequestOption) (*EchoMessage, error) {
	var err error

	if req.Str == nil {
		req.Str = scw.StringPtr(namegenerator.GetRandomName("name"))
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/echo",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp EchoMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetEnumRequest struct {
	// Type:
	//
	// Default value: unknown
	Type Type `json:"-"`
}

// GetEnum: get enum
func (s *API) GetEnum(req *GetEnumRequest, opts ...scw.RequestOption) (*GetEnumResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "type", req.Type)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/enum",
		Query:   query,
		Headers: http.Header{},
	}

	var resp GetEnumResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostEnumRequest struct {
	// Type:
	//
	// Default value: unknown
	Type Type `json:"type"`
}

// PostEnum: post enum
func (s *API) PostEnum(req *PostEnumRequest, opts ...scw.RequestOption) (*PostEnumResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/enum",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostEnumResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PatchEnumRequest struct {
	// Type:
	//
	// Default value: unknown
	Type Type `json:"type"`
}

// PatchEnum: patch enum
func (s *API) PatchEnum(req *PatchEnumRequest, opts ...scw.RequestOption) (*PatchEnumResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/test-internal/v1/enum",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PatchEnumResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostTagsRequest struct {
	Tags *[]string `json:"tags"`
}

// PostTags: post tags
func (s *API) PostTags(req *PostTagsRequest, opts ...scw.RequestOption) (*PostTagsResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/test-internal/v1/tags",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostTagsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostFileRequest struct {
	File *scw.File `json:"file"`
}

// PostFile: post file
func (s *API) PostFile(req *PostFileRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/file",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostIPRequest struct {
	IPV4 net.IP `json:"ip_v4"`

	IPV6 net.IP `json:"ip_v6"`

	IP net.IP `json:"ip"`
}

// PostIP: post IP
func (s *API) PostIP(req *PostIPRequest, opts ...scw.RequestOption) (*PostIPResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/ip",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostIPResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostScalarTypesRequest struct {
	DoubleField float64 `json:"double_field"`

	FloatField float32 `json:"float_field"`

	Int32Field int32 `json:"int32_field"`

	Int64Field int64 `json:"int64_field"`

	Uint32Field uint32 `json:"uint32_field"`

	Uint64Field uint64 `json:"uint64_field"`

	BoolField bool `json:"bool_field"`

	StringField string `json:"string_field"`
}

// PostScalarTypes: post scalar types
func (s *API) PostScalarTypes(req *PostScalarTypesRequest, opts ...scw.RequestOption) (*PostScalarTypesResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/scalar-types",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostScalarTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetMetadataRequest struct {
}

func (s *API) GetMetadata(req *GetMetadataRequest, opts ...scw.RequestOption) (*MetadataResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/metadata",
		Headers: http.Header{},
	}

	var resp MetadataResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostWaitRequest struct {
	// Duration: waiting duration
	Duration *time.Duration `json:"duration"`
}

func (m *PostWaitRequest) UnmarshalJSON(b []byte) error {
	type tmpType PostWaitRequest
	tmp := struct {
		tmpType

		TmpDuration *marshaler.Duration `json:"duration"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = PostWaitRequest(tmp.tmpType)

	m.Duration = tmp.TmpDuration.Standard()
	return nil
}

func (m PostWaitRequest) MarshalJSON() ([]byte, error) {
	type tmpType PostWaitRequest
	tmp := struct {
		tmpType

		TmpDuration *marshaler.Duration `json:"duration"`
	}{
		tmpType: tmpType(m),

		TmpDuration: marshaler.NewDuration(m.Duration),
	}
	return json.Marshal(tmp)
}

// PostWait: wait until a given time in second
func (s *API) PostWait(req *PostWaitRequest, opts ...scw.RequestOption) error {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/wait",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return err
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type GetRegionRequest struct {
}

// GetRegion: get a region
func (s *API) GetRegion(req *GetRegionRequest, opts ...scw.RequestOption) (*getRegionResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/region",
		Headers: http.Header{},
	}

	var resp getRegionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostOrganizationIDRequest struct {
	OrganizationID string `json:"organization_id"`
}

// PostOrganizationID: post an organization ID
func (s *API) PostOrganizationID(req *PostOrganizationIDRequest, opts ...scw.RequestOption) (*PostOrganizationIDResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/organization-id",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostOrganizationIDResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostDeprecatedOrganizationRequest struct {
	// Deprecated
	Organization *string `json:"organization"`
}

// Deprecated: PostDeprecatedOrganization: post a deprecated organization ID
func (s *API) PostDeprecatedOrganization(req *PostDeprecatedOrganizationRequest, opts ...scw.RequestOption) (*PostDeprecatedOrganizationResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/deprecated-organization",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostDeprecatedOrganizationResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostProjectIDRequest struct {
	ProjectID string `json:"project_id"`
}

func (s *API) PostProjectID(req *PostProjectIDRequest, opts ...scw.RequestOption) (*PostProjectIDResponse, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/project-id",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostProjectIDResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostDeprecatedProjectRequest struct {
	// Deprecated
	Project *string `json:"project"`
}

// Deprecated
func (s *API) PostDeprecatedProject(req *PostDeprecatedProjectRequest, opts ...scw.RequestOption) (*PostDeprecatedProjectResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/deprecated-project",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostDeprecatedProjectResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostOneOfRequest struct {

	// Precisely one of Test, Test2 must be set.
	Test *string `json:"test,omitempty"`

	// Precisely one of Test, Test2 must be set.
	Test2 *int32 `json:"test2,omitempty"`

	// Precisely one of Test3, TestNested must be set.
	TestNested *EchoMessage `json:"test_nested,omitempty"`

	// Precisely one of Test3, TestNested must be set.
	Test3 *int32 `json:"test3,omitempty"`
}

func (s *API) PostOneOf(req *PostOneOfRequest, opts ...scw.RequestOption) (*PostOneOfResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/oneof",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostOneOfResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostBodyAndPathSimpleRequest struct {
	Path string `json:"-"`

	Body string `json:"body"`
}

func (s *API) PostBodyAndPathSimple(req *PostBodyAndPathSimpleRequest, opts ...scw.RequestOption) (*PostBodyAndPathSimpleResponse, error) {
	var err error

	if fmt.Sprint(req.Path) == "" {
		return nil, errors.New("field Path cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/paths/" + fmt.Sprint(req.Path) + "/simple",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req.Body)
	if err != nil {
		return nil, err
	}

	var resp PostBodyAndPathSimpleResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostBodyAndPathSimple2Request struct {
	Path string `json:"-"`

	Body string `json:"body"`
}

func (s *API) PostBodyAndPathSimple2(req *PostBodyAndPathSimple2Request, opts ...scw.RequestOption) (*PostBodyAndPathSimple2Response, error) {
	var err error

	if fmt.Sprint(req.Path) == "" {
		return nil, errors.New("field Path cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/paths/" + fmt.Sprint(req.Path) + "/simple2",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostBodyAndPathSimple2Response

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostBodyAndPathComplexRequest struct {
	Path string `json:"-"`

	Body *PostBodyAndPathSimpleRequest `json:"body"`
}

func (s *API) PostBodyAndPathComplex(req *PostBodyAndPathComplexRequest, opts ...scw.RequestOption) (*PostBodyAndPathComplexResponse, error) {
	var err error

	if fmt.Sprint(req.Path) == "" {
		return nil, errors.New("field Path cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/paths/" + fmt.Sprint(req.Path) + "/complex",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req.Body)
	if err != nil {
		return nil, err
	}

	var resp PostBodyAndPathComplexResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostBodyAndPathAndQueryRequest struct {
	Path string `json:"-"`

	Query string `json:"-"`

	Body *PostBodyAndPathSimpleRequest `json:"body"`
}

func (s *API) PostBodyAndPathAndQuery(req *PostBodyAndPathAndQueryRequest, opts ...scw.RequestOption) (*PostBodyAndPathAndQueryResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "query", req.Query)

	if fmt.Sprint(req.Path) == "" {
		return nil, errors.New("field Path cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/paths/" + fmt.Sprint(req.Path) + "/bodyquery",
		Query:   query,
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req.Body)
	if err != nil {
		return nil, err
	}

	var resp PostBodyAndPathAndQueryResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostAllMapTypesRequest struct {
	MapInt32Int32 map[int32]int32 `json:"map_int32_int32"`

	MapInt64Int64 map[int64]int64 `json:"map_int64_int64"`

	MapUint32Uint32 map[uint32]uint32 `json:"map_uint32_uint32"`

	MapUint64Uint64 map[uint64]uint64 `json:"map_uint64_uint64"`

	MapSint32Sint32 map[int32]int32 `json:"map_sint32_sint32"`

	MapSint64Sint64 map[int64]int64 `json:"map_sint64_sint64"`

	MapFixed32Fixed32 map[uint32]uint32 `json:"map_fixed32_fixed32"`

	MapFixed64Fixed64 map[uint64]uint64 `json:"map_fixed64_fixed64"`

	MapSfixed32Sfixed32 map[int32]int32 `json:"map_sfixed32_sfixed32"`

	MapSfixed64Sfixed64 map[int64]int64 `json:"map_sfixed64_sfixed64"`

	MapInt32Float map[int32]float32 `json:"map_int32_float"`

	MapInt32Double map[int32]float64 `json:"map_int32_double"`

	MapStringString map[string]string `json:"map_string_string"`

	MapInt32Bytes map[int32][]byte `json:"map_int32_bytes"`

	MapInt32Enum map[int32]MapEnum `json:"map_int32_enum"`

	MapInt32AllTypes map[int32]*PostAllTypesRequest `json:"map_int32_all_types"`

	MapInt32IP map[int32]net.IP `json:"map_int32_ip"`

	MapInt32StdDuration map[int32]*time.Duration `json:"map_int32_std_duration"`

	MapInt32StdLongDuration map[int32]*time.Duration `json:"map_int32_std_long_duration"`

	MapInt32Size map[int32]scw.Size `json:"map_int32_size"`

	MapInt32Uint64Size map[int32]scw.Size `json:"map_int32_uint64_size"`

	MapInt32Uint64valueSize map[int32]*scw.Size `json:"map_int32_uint64value_size"`

	MapInt32StringIP map[int32]net.IP `json:"map_int32_string_ip"`

	MapInt32StringValueIP map[int32]*net.IP `json:"map_int32_string_value_ip"`

	MapInt32IPv4 map[int32]net.IP `json:"map_int32_ipv4"`

	MapInt32StringIPv4 map[int32]net.IP `json:"map_int32_string_ipv4"`

	MapInt32StringValueIPv4 map[int32]*net.IP `json:"map_int32_string_value_ipv4"`

	MapInt32IPv6 map[int32]net.IP `json:"map_int32_ipv6"`

	MapInt32StringIPv6 map[int32]net.IP `json:"map_int32_string_ipv6"`

	MapInt32StringValueIPv6 map[int32]*net.IP `json:"map_int32_string_value_ipv6"`

	MapInt32StringsValue map[int32]*[]string `json:"map_int32_strings_value"`

	MapInt32Duration map[int32]*scw.Duration `json:"map_int32_duration"`
}

func (m *PostAllMapTypesRequest) UnmarshalJSON(b []byte) error {
	type tmpType PostAllMapTypesRequest
	tmp := struct {
		tmpType

		TmpMapInt32StdDuration     marshaler.Durationint32Map     `json:"map_int32_std_duration"`
		TmpMapInt32StdLongDuration marshaler.LongDurationint32Map `json:"map_int32_std_long_duration"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = PostAllMapTypesRequest(tmp.tmpType)

	m.MapInt32StdDuration = tmp.TmpMapInt32StdDuration.Standard()
	m.MapInt32StdLongDuration = tmp.TmpMapInt32StdLongDuration.Standard()
	return nil
}

func (m PostAllMapTypesRequest) MarshalJSON() ([]byte, error) {
	type tmpType PostAllMapTypesRequest
	tmp := struct {
		tmpType

		TmpMapInt32StdDuration     marshaler.Durationint32Map     `json:"map_int32_std_duration"`
		TmpMapInt32StdLongDuration marshaler.LongDurationint32Map `json:"map_int32_std_long_duration"`
	}{
		tmpType: tmpType(m),

		TmpMapInt32StdDuration:     marshaler.NewDurationint32Map(m.MapInt32StdDuration),
		TmpMapInt32StdLongDuration: marshaler.NewLongDurationint32Map(m.MapInt32StdLongDuration),
	}
	return json.Marshal(tmp)
}

func (s *API) PostAllMapTypes(req *PostAllMapTypesRequest, opts ...scw.RequestOption) (*PostAllMapTypesResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/map",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostAllMapTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostAllTypesRequest struct {
	SingularInt32 int32 `json:"singular_int32"`

	SingularInt64 int64 `json:"singular_int64"`

	SingularUint32 uint32 `json:"singular_uint32"`

	SingularUint64 uint64 `json:"singular_uint64"`

	SingularSint32 int32 `json:"singular_sint32"`

	SingularSint64 int64 `json:"singular_sint64"`

	SingularFixed32 uint32 `json:"singular_fixed32"`

	SingularFixed64 uint64 `json:"singular_fixed64"`

	SingularSfixed32 int32 `json:"singular_sfixed32"`

	SingularSfixed64 int64 `json:"singular_sfixed64"`

	SingularFloat float32 `json:"singular_float"`

	SingularDouble float64 `json:"singular_double"`

	SingularBool bool `json:"singular_bool"`

	SingularString string `json:"singular_string"`

	SingularBytes []byte `json:"singular_bytes"`

	SingularNestedMessage *PostAllTypesRequestNestedMessage `json:"singular_nested_message"`
	// SingularNestedEnum:
	//
	// Default value: ZERO
	SingularNestedEnum PostAllTypesRequestNestedEnum `json:"singular_nested_enum"`

	RepeatedInt32 []int32 `json:"repeated_int32"`

	RepeatedInt64 []int64 `json:"repeated_int64"`

	RepeatedUint32 []uint32 `json:"repeated_uint32"`

	RepeatedUint64 []uint64 `json:"repeated_uint64"`

	RepeatedSint32 []int32 `json:"repeated_sint32"`

	RepeatedSint64 []int64 `json:"repeated_sint64"`

	RepeatedFixed32 []uint32 `json:"repeated_fixed32"`

	RepeatedFixed64 []uint64 `json:"repeated_fixed64"`

	RepeatedSfixed32 []int32 `json:"repeated_sfixed32"`

	RepeatedSfixed64 []int64 `json:"repeated_sfixed64"`

	RepeatedFloat []float32 `json:"repeated_float"`

	RepeatedDouble []float64 `json:"repeated_double"`

	RepeatedBool []bool `json:"repeated_bool"`

	RepeatedString []string `json:"repeated_string"`

	RepeatedBytes [][]byte `json:"repeated_bytes"`

	RepeatedNestedMessage []*PostAllTypesRequestNestedMessage `json:"repeated_nested_message"`

	RepeatedNestedEnum []PostAllTypesRequestNestedEnum `json:"repeated_nested_enum"`

	// Precisely one of OneofBytes, OneofNestedMessage, OneofString, OneofUint32 must be set.
	OneofUint32 *uint32 `json:"oneof_uint32,omitempty"`

	// Precisely one of OneofBytes, OneofNestedMessage, OneofString, OneofUint32 must be set.
	OneofNestedMessage *PostAllTypesRequestNestedMessage `json:"oneof_nested_message,omitempty"`

	// Precisely one of OneofBytes, OneofNestedMessage, OneofString, OneofUint32 must be set.
	OneofString *string `json:"oneof_string,omitempty"`

	// Precisely one of OneofBytes, OneofNestedMessage, OneofString, OneofUint32 must be set.
	OneofBytes *[]byte `json:"oneof_bytes,omitempty"`

	SingularDoubleValue *float64 `json:"singular_double_value"`

	SingularFloatValue *float32 `json:"singular_float_value"`

	SingularInt64Value *int64 `json:"singular_int64_value"`

	SingularUint64Value *uint64 `json:"singular_uint64_value"`

	SingularInt32Value *int32 `json:"singular_int32_value"`

	SingularUint32Value *uint32 `json:"singular_uint32_value"`

	SingularBoolValue *bool `json:"singular_bool_value"`

	SingularStringValue *string `json:"singular_string_value"`

	SingularBytesValue *[]byte `json:"singular_bytes_value"`

	SingularTimestamp *time.Time `json:"singular_timestamp"`

	SingularAny interface{} `json:"singular_any"`

	SingularStruct []byte `json:"singular_struct"`

	SingularMoney *scw.Money `json:"singular_money"`

	SingularStringsValue *[]string `json:"singular_strings_value"`

	SingularDuration *scw.Duration `json:"singular_duration"`

	SingularIP net.IP `json:"singular_ip"`

	SingularStringIP net.IP `json:"singular_string_ip"`

	SingularStringValueIP *net.IP `json:"singular_string_value_ip"`

	SingularIPv4 net.IP `json:"singular_ipv4"`

	SingularStringIPv4 net.IP `json:"singular_string_ipv4"`

	SingularStringValueIPv4 *net.IP `json:"singular_string_value_ipv4"`

	SingularIPv6 net.IP `json:"singular_ipv6"`

	SingularStringIPv6 net.IP `json:"singular_string_ipv6"`

	SingularStringValueIPv6 *net.IP `json:"singular_string_value_ipv6"`

	SingularStdDuration *time.Duration `json:"singular_std_duration"`

	SingularStdLongDuration *time.Duration `json:"singular_std_long_duration"`

	SingularSize scw.Size `json:"singular_size"`

	SingularUint64Size scw.Size `json:"singular_uint64_size"`

	SingularUint64valueSize *scw.Size `json:"singular_uint64value_size"`

	SingularStringIpnet scw.IPNet `json:"singular_string_ipnet"`

	SingularStringValueIpnet *scw.IPNet `json:"singular_string_value_ipnet"`

	RepeatedDoubleValue []*float64 `json:"repeated_double_value"`

	RepeatedFloatValue []*float32 `json:"repeated_float_value"`

	RepeatedInt64Value []*int64 `json:"repeated_int64_value"`

	RepeatedUint64Value []*uint64 `json:"repeated_uint64_value"`

	RepeatedInt32Value []*int32 `json:"repeated_int32_value"`

	RepeatedUint32Value []*uint32 `json:"repeated_uint32_value"`

	RepeatedBoolValue []*bool `json:"repeated_bool_value"`

	RepeatedStringValue []*string `json:"repeated_string_value"`

	RepeatedBytesValue []*[]byte `json:"repeated_bytes_value"`

	RepeatedTimestamp []*time.Time `json:"repeated_timestamp"`

	RepeatedAny []interface{} `json:"repeated_any"`

	RepeatedStruct [][]byte `json:"repeated_struct"`

	RepeatedMoney []*scw.Money `json:"repeated_money"`

	RepeatedStringsValue []*[]string `json:"repeated_strings_value"`

	RepeatedDuration []*scw.Duration `json:"repeated_duration"`

	RepeatedIP []net.IP `json:"repeated_ip"`

	RepeatedStringIP []net.IP `json:"repeated_string_ip"`

	RepeatedStringValueIP []*net.IP `json:"repeated_string_value_ip"`

	RepeatedIPv4 []net.IP `json:"repeated_ipv4"`

	RepeatedStringIPv4 []net.IP `json:"repeated_string_ipv4"`

	RepeatedStringValueIPv4 []*net.IP `json:"repeated_string_value_ipv4"`

	RepeatedIPv6 []net.IP `json:"repeated_ipv6"`

	RepeatedStringIPv6 []net.IP `json:"repeated_string_ipv6"`

	RepeatedStringValueIPv6 []*net.IP `json:"repeated_string_value_ipv6"`

	RepeatedStdDuration []*time.Duration `json:"repeated_std_duration"`

	RepeatedStdLongDuration []*time.Duration `json:"repeated_std_long_duration"`

	RepeatedSize []scw.Size `json:"repeated_size"`

	RepeatedUint64Size []scw.Size `json:"repeated_uint64_size"`

	RepeatedUint64valueSize []*scw.Size `json:"repeated_uint64value_size"`

	RepeatedStringIpnet []scw.IPNet `json:"repeated_string_ipnet"`

	RepeatedStringValueIpnet []*scw.IPNet `json:"repeated_string_value_ipnet"`
}

func (m *PostAllTypesRequest) UnmarshalJSON(b []byte) error {
	type tmpType PostAllTypesRequest
	tmp := struct {
		tmpType

		TmpSingularStdDuration     *marshaler.Duration         `json:"singular_std_duration"`
		TmpSingularStdLongDuration *marshaler.LongDuration     `json:"singular_std_long_duration"`
		TmpRepeatedStdDuration     marshaler.DurationSlice     `json:"repeated_std_duration"`
		TmpRepeatedStdLongDuration marshaler.LongDurationSlice `json:"repeated_std_long_duration"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = PostAllTypesRequest(tmp.tmpType)

	m.SingularStdDuration = tmp.TmpSingularStdDuration.Standard()
	m.SingularStdLongDuration = tmp.TmpSingularStdLongDuration.Standard()
	m.RepeatedStdDuration = tmp.TmpRepeatedStdDuration.Standard()
	m.RepeatedStdLongDuration = tmp.TmpRepeatedStdLongDuration.Standard()
	return nil
}

func (m PostAllTypesRequest) MarshalJSON() ([]byte, error) {
	type tmpType PostAllTypesRequest
	tmp := struct {
		tmpType

		TmpSingularStdDuration     *marshaler.Duration         `json:"singular_std_duration"`
		TmpSingularStdLongDuration *marshaler.LongDuration     `json:"singular_std_long_duration"`
		TmpRepeatedStdDuration     marshaler.DurationSlice     `json:"repeated_std_duration"`
		TmpRepeatedStdLongDuration marshaler.LongDurationSlice `json:"repeated_std_long_duration"`
	}{
		tmpType: tmpType(m),

		TmpSingularStdDuration:     marshaler.NewDuration(m.SingularStdDuration),
		TmpSingularStdLongDuration: marshaler.NewLongDuration(m.SingularStdLongDuration),
		TmpRepeatedStdDuration:     marshaler.NewDurationSlice(m.RepeatedStdDuration),
		TmpRepeatedStdLongDuration: marshaler.NewLongDurationSlice(m.RepeatedStdLongDuration),
	}
	return json.Marshal(tmp)
}

// PostAllTypes: post all types
func (s *API) PostAllTypes(req *PostAllTypesRequest, opts ...scw.RequestOption) (*PostAllTypesResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/alltypes",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostAllTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListCharactersRequest struct {
	// Page: a positive integer to choose the page to display
	Page *int32 `json:"-"`
	// PageSize: a positive integer lower or equal to 100 to select the number of items to display
	//
	// Default value: 10
	PageSize *uint32 `json:"-"`
	// OrderBy: order the listing
	//
	// Default value: asc
	OrderBy ListCharactersRequestOrderBy `json:"-"`
	// Name: filter characters by name, this is a "contains" matching
	//
	// Default value: i-am-fake
	Name *string `json:"-"`
	// Tags: dummy tags to check the comma_separated_list argument
	Tags []string `json:"-"`
}

// ListCharacters: list The Lord of the Rings characters
func (s *API) ListCharacters(req *ListCharactersRequest, opts ...scw.RequestOption) (*ListCharactersResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "name", req.Name)
	if len(req.Tags) != 0 {
		parameter.AddToQuery(query, "tags", strings.Join(req.Tags, ","))
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/characters",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListCharactersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostEchoTimeSeriesRequest struct {
	Metrics *scw.TimeSeries `json:"metrics"`
}

// PostEchoTimeSeries: echo metrics
func (s *API) PostEchoTimeSeries(req *PostEchoTimeSeriesRequest, opts ...scw.RequestOption) (*PostEchoTimeSeriesResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/echo-timeseries",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostEchoTimeSeriesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Service NoAuthAPI

// Service RegionalAPI

type RegionalAPIGetServiceInfoRequest struct {
	Region scw.Region `json:"-"`
}

func (s *RegionalAPI) GetServiceInfo(req *RegionalAPIGetServiceInfoRequest, opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/regions/" + fmt.Sprint(req.Region) + "",
		Headers: http.Header{},
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RegionalAPIGetRegionRequest struct {
	Region scw.Region `json:"-"`
}

func (s *RegionalAPI) GetRegion(req *RegionalAPIGetRegionRequest, opts ...scw.RequestOption) (*getRegionResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/regions/" + fmt.Sprint(req.Region) + "/region",
		Headers: http.Header{},
	}

	var resp getRegionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RegionalAPIGetMetadataRequest struct {
	Region scw.Region `json:"-"`
}

func (s *RegionalAPI) GetMetadata(req *RegionalAPIGetMetadataRequest, opts ...scw.RequestOption) (*MetadataResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/regions/" + fmt.Sprint(req.Region) + "/metadata",
		Headers: http.Header{},
	}

	var resp MetadataResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RegionalAPIPostEchoRequest struct {
	Region scw.Region `json:"-"`

	Str *string `json:"str"`

	Strs []string `json:"strs"`
}

func (s *RegionalAPI) PostEcho(req *RegionalAPIPostEchoRequest, opts ...scw.RequestOption) (*EchoMessage, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Str == nil {
		req.Str = scw.StringPtr(namegenerator.GetRandomName("name"))
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/regions/" + fmt.Sprint(req.Region) + "/echo",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp EchoMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Service StreamAPI

// Service ZoneAPI

type ZoneAPIGetZoneRequest struct {
	Zone scw.Zone `json:"-"`
}

// GetZone: get a zone
func (s *ZoneAPI) GetZone(req *ZoneAPIGetZoneRequest, opts ...scw.RequestOption) (*GetZoneResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/zone",
		Headers: http.Header{},
	}

	var resp GetZoneResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ZoneAPIGetMetadataRequest struct {
	Zone scw.Zone `json:"-"`
}

// GetMetadata: get metadata
func (s *ZoneAPI) GetMetadata(req *ZoneAPIGetMetadataRequest, opts ...scw.RequestOption) (*MetadataResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/metadata",
		Headers: http.Header{},
	}

	var resp MetadataResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ZoneAPIPostEchoRequest struct {
	Zone scw.Zone `json:"-"`

	Str *string `json:"str"`

	Strs []string `json:"strs"`
}

func (s *ZoneAPI) PostEcho(req *ZoneAPIPostEchoRequest, opts ...scw.RequestOption) (*EchoMessage, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.Str == nil {
		req.Str = scw.StringPtr(namegenerator.GetRandomName("name"))
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/echo",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp EchoMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCharactersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCharactersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListCharactersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Characters = append(r.Characters, results.Characters...)
	r.TotalCount += uint32(len(results.Characters))
	return uint32(len(results.Characters)), nil
}

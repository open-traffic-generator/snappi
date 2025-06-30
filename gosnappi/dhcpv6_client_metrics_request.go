package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ClientMetricsRequest *****
type dhcpv6ClientMetricsRequest struct {
	validation
	obj          *otg.Dhcpv6ClientMetricsRequest
	marshaller   marshalDhcpv6ClientMetricsRequest
	unMarshaller unMarshalDhcpv6ClientMetricsRequest
}

func NewDhcpv6ClientMetricsRequest() Dhcpv6ClientMetricsRequest {
	obj := dhcpv6ClientMetricsRequest{obj: &otg.Dhcpv6ClientMetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ClientMetricsRequest) msg() *otg.Dhcpv6ClientMetricsRequest {
	return obj.obj
}

func (obj *dhcpv6ClientMetricsRequest) setMsg(msg *otg.Dhcpv6ClientMetricsRequest) Dhcpv6ClientMetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ClientMetricsRequest struct {
	obj *dhcpv6ClientMetricsRequest
}

type marshalDhcpv6ClientMetricsRequest interface {
	// ToProto marshals Dhcpv6ClientMetricsRequest to protobuf object *otg.Dhcpv6ClientMetricsRequest
	ToProto() (*otg.Dhcpv6ClientMetricsRequest, error)
	// ToPbText marshals Dhcpv6ClientMetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ClientMetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ClientMetricsRequest to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv6ClientMetricsRequest struct {
	obj *dhcpv6ClientMetricsRequest
}

type unMarshalDhcpv6ClientMetricsRequest interface {
	// FromProto unmarshals Dhcpv6ClientMetricsRequest from protobuf object *otg.Dhcpv6ClientMetricsRequest
	FromProto(msg *otg.Dhcpv6ClientMetricsRequest) (Dhcpv6ClientMetricsRequest, error)
	// FromPbText unmarshals Dhcpv6ClientMetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ClientMetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ClientMetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ClientMetricsRequest) Marshal() marshalDhcpv6ClientMetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ClientMetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ClientMetricsRequest) Unmarshal() unMarshalDhcpv6ClientMetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ClientMetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ClientMetricsRequest) ToProto() (*otg.Dhcpv6ClientMetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ClientMetricsRequest) FromProto(msg *otg.Dhcpv6ClientMetricsRequest) (Dhcpv6ClientMetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ClientMetricsRequest) ToPbText() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (m *unMarshaldhcpv6ClientMetricsRequest) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshaldhcpv6ClientMetricsRequest) ToYaml() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *unMarshaldhcpv6ClientMetricsRequest) FromYaml(value string) error {
	if value == "" {
		value = "{}"
	}
	data, err := yaml.YAMLToJSON([]byte(value))
	if err != nil {
		return err
	}
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	uError := opts.Unmarshal([]byte(data), m.obj.msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshaldhcpv6ClientMetricsRequest) ToJson() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *unMarshaldhcpv6ClientMetricsRequest) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), m.obj.msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *dhcpv6ClientMetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ClientMetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ClientMetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ClientMetricsRequest) Clone() (Dhcpv6ClientMetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ClientMetricsRequest()
	data, err := proto.Marshal(obj.msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// Dhcpv6ClientMetricsRequest is the request to retrieve DHCPv6 per client metrics/statistics.
type Dhcpv6ClientMetricsRequest interface {
	Validation
	// msg marshals Dhcpv6ClientMetricsRequest to protobuf object *otg.Dhcpv6ClientMetricsRequest
	// and doesn't set defaults
	msg() *otg.Dhcpv6ClientMetricsRequest
	// setMsg unmarshals Dhcpv6ClientMetricsRequest from protobuf object *otg.Dhcpv6ClientMetricsRequest
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ClientMetricsRequest) Dhcpv6ClientMetricsRequest
	// provides marshal interface
	Marshal() marshalDhcpv6ClientMetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ClientMetricsRequest
	// validate validates Dhcpv6ClientMetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ClientMetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ClientNames returns []string, set in Dhcpv6ClientMetricsRequest.
	ClientNames() []string
	// SetClientNames assigns []string provided by user to Dhcpv6ClientMetricsRequest
	SetClientNames(value []string) Dhcpv6ClientMetricsRequest
	// ColumnNames returns []Dhcpv6ClientMetricsRequestColumnNamesEnum, set in Dhcpv6ClientMetricsRequest
	ColumnNames() []Dhcpv6ClientMetricsRequestColumnNamesEnum
	// SetColumnNames assigns []Dhcpv6ClientMetricsRequestColumnNamesEnum provided by user to Dhcpv6ClientMetricsRequest
	SetColumnNames(value []Dhcpv6ClientMetricsRequestColumnNamesEnum) Dhcpv6ClientMetricsRequest
}

// The names of DHCPv6 clients to return results for. An empty list will return results for all DHCPv6 client.
//
// x-constraint:
// - /components/schemas/Device.Dhcpv6client/properties/name
//
// x-constraint:
// - /components/schemas/Device.Dhcpv6client/properties/name
//
// ClientNames returns a []string
func (obj *dhcpv6ClientMetricsRequest) ClientNames() []string {
	if obj.obj.ClientNames == nil {
		obj.obj.ClientNames = make([]string, 0)
	}
	return obj.obj.ClientNames
}

// The names of DHCPv6 clients to return results for. An empty list will return results for all DHCPv6 client.
//
// x-constraint:
// - /components/schemas/Device.Dhcpv6client/properties/name
//
// x-constraint:
// - /components/schemas/Device.Dhcpv6client/properties/name
//
// SetClientNames sets the []string value in the Dhcpv6ClientMetricsRequest object
func (obj *dhcpv6ClientMetricsRequest) SetClientNames(value []string) Dhcpv6ClientMetricsRequest {

	if obj.obj.ClientNames == nil {
		obj.obj.ClientNames = make([]string, 0)
	}
	obj.obj.ClientNames = value

	return obj
}

type Dhcpv6ClientMetricsRequestColumnNamesEnum string

// Enum of ColumnNames on Dhcpv6ClientMetricsRequest
var Dhcpv6ClientMetricsRequestColumnNames = struct {
	SOLICITS_SENT                 Dhcpv6ClientMetricsRequestColumnNamesEnum
	ADVERTISEMENTS_RECEIVED       Dhcpv6ClientMetricsRequestColumnNamesEnum
	ADVERTISEMENTS_IGNORED        Dhcpv6ClientMetricsRequestColumnNamesEnum
	REQUESTS_SENT                 Dhcpv6ClientMetricsRequestColumnNamesEnum
	NACKS_RECEIVED                Dhcpv6ClientMetricsRequestColumnNamesEnum
	REPLIES_RECEIVED              Dhcpv6ClientMetricsRequestColumnNamesEnum
	INFORMATION_REQUESTS_SENT     Dhcpv6ClientMetricsRequestColumnNamesEnum
	RENEWS_SENT                   Dhcpv6ClientMetricsRequestColumnNamesEnum
	REBINDS_SENT                  Dhcpv6ClientMetricsRequestColumnNamesEnum
	RELEASES_SENT                 Dhcpv6ClientMetricsRequestColumnNamesEnum
	RECONFIGURES_RECEIVED         Dhcpv6ClientMetricsRequestColumnNamesEnum
	RAPID_COMMIT_SOLICITS_SENT    Dhcpv6ClientMetricsRequestColumnNamesEnum
	RAPID_COMMIT_REPLIES_RECEIVED Dhcpv6ClientMetricsRequestColumnNamesEnum
}{
	SOLICITS_SENT:                 Dhcpv6ClientMetricsRequestColumnNamesEnum("solicits_sent"),
	ADVERTISEMENTS_RECEIVED:       Dhcpv6ClientMetricsRequestColumnNamesEnum("advertisements_received"),
	ADVERTISEMENTS_IGNORED:        Dhcpv6ClientMetricsRequestColumnNamesEnum("advertisements_ignored"),
	REQUESTS_SENT:                 Dhcpv6ClientMetricsRequestColumnNamesEnum("requests_sent"),
	NACKS_RECEIVED:                Dhcpv6ClientMetricsRequestColumnNamesEnum("nacks_received"),
	REPLIES_RECEIVED:              Dhcpv6ClientMetricsRequestColumnNamesEnum("replies_received"),
	INFORMATION_REQUESTS_SENT:     Dhcpv6ClientMetricsRequestColumnNamesEnum("information_requests_sent"),
	RENEWS_SENT:                   Dhcpv6ClientMetricsRequestColumnNamesEnum("renews_sent"),
	REBINDS_SENT:                  Dhcpv6ClientMetricsRequestColumnNamesEnum("rebinds_sent"),
	RELEASES_SENT:                 Dhcpv6ClientMetricsRequestColumnNamesEnum("releases_sent"),
	RECONFIGURES_RECEIVED:         Dhcpv6ClientMetricsRequestColumnNamesEnum("reconfigures_received"),
	RAPID_COMMIT_SOLICITS_SENT:    Dhcpv6ClientMetricsRequestColumnNamesEnum("rapid_commit_solicits_sent"),
	RAPID_COMMIT_REPLIES_RECEIVED: Dhcpv6ClientMetricsRequestColumnNamesEnum("rapid_commit_replies_received"),
}

func (obj *dhcpv6ClientMetricsRequest) ColumnNames() []Dhcpv6ClientMetricsRequestColumnNamesEnum {
	items := []Dhcpv6ClientMetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, Dhcpv6ClientMetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain. If the list is empty then all columns will be returned except for any result_groups. The name of the DHCPv6 client cannot be excluded.
// SetColumnNames sets the []string value in the Dhcpv6ClientMetricsRequest object
func (obj *dhcpv6ClientMetricsRequest) SetColumnNames(value []Dhcpv6ClientMetricsRequestColumnNamesEnum) Dhcpv6ClientMetricsRequest {

	items := []otg.Dhcpv6ClientMetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.Dhcpv6ClientMetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.Dhcpv6ClientMetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *dhcpv6ClientMetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *dhcpv6ClientMetricsRequest) setDefault() {

}

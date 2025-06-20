package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ServerMetricsRequest *****
type dhcpv6ServerMetricsRequest struct {
	validation
	obj          *otg.Dhcpv6ServerMetricsRequest
	marshaller   marshalDhcpv6ServerMetricsRequest
	unMarshaller unMarshalDhcpv6ServerMetricsRequest
}

func NewDhcpv6ServerMetricsRequest() Dhcpv6ServerMetricsRequest {
	obj := dhcpv6ServerMetricsRequest{obj: &otg.Dhcpv6ServerMetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ServerMetricsRequest) msg() *otg.Dhcpv6ServerMetricsRequest {
	return obj.obj
}

func (obj *dhcpv6ServerMetricsRequest) setMsg(msg *otg.Dhcpv6ServerMetricsRequest) Dhcpv6ServerMetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ServerMetricsRequest struct {
	obj *dhcpv6ServerMetricsRequest
}

type marshalDhcpv6ServerMetricsRequest interface {
	// ToProto marshals Dhcpv6ServerMetricsRequest to protobuf object *otg.Dhcpv6ServerMetricsRequest
	ToProto() (*otg.Dhcpv6ServerMetricsRequest, error)
	// ToPbText marshals Dhcpv6ServerMetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ServerMetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ServerMetricsRequest to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv6ServerMetricsRequest struct {
	obj *dhcpv6ServerMetricsRequest
}

type unMarshalDhcpv6ServerMetricsRequest interface {
	// FromProto unmarshals Dhcpv6ServerMetricsRequest from protobuf object *otg.Dhcpv6ServerMetricsRequest
	FromProto(msg *otg.Dhcpv6ServerMetricsRequest) (Dhcpv6ServerMetricsRequest, error)
	// FromPbText unmarshals Dhcpv6ServerMetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ServerMetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ServerMetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ServerMetricsRequest) Marshal() marshalDhcpv6ServerMetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ServerMetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ServerMetricsRequest) Unmarshal() unMarshalDhcpv6ServerMetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ServerMetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ServerMetricsRequest) ToProto() (*otg.Dhcpv6ServerMetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ServerMetricsRequest) FromProto(msg *otg.Dhcpv6ServerMetricsRequest) (Dhcpv6ServerMetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ServerMetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ServerMetricsRequest) FromPbText(value string) error {
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

func (m *marshaldhcpv6ServerMetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ServerMetricsRequest) FromYaml(value string) error {
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

func (m *marshaldhcpv6ServerMetricsRequest) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ServerMetricsRequest) FromJson(value string) error {
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

func (obj *dhcpv6ServerMetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ServerMetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ServerMetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ServerMetricsRequest) Clone() (Dhcpv6ServerMetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ServerMetricsRequest()
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

// Dhcpv6ServerMetricsRequest is the request to retrieve DHCPv6 per Server metrics/statistics.
type Dhcpv6ServerMetricsRequest interface {
	Validation
	// msg marshals Dhcpv6ServerMetricsRequest to protobuf object *otg.Dhcpv6ServerMetricsRequest
	// and doesn't set defaults
	msg() *otg.Dhcpv6ServerMetricsRequest
	// setMsg unmarshals Dhcpv6ServerMetricsRequest from protobuf object *otg.Dhcpv6ServerMetricsRequest
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ServerMetricsRequest) Dhcpv6ServerMetricsRequest
	// provides marshal interface
	Marshal() marshalDhcpv6ServerMetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ServerMetricsRequest
	// validate validates Dhcpv6ServerMetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ServerMetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ServerNames returns []string, set in Dhcpv6ServerMetricsRequest.
	ServerNames() []string
	// SetServerNames assigns []string provided by user to Dhcpv6ServerMetricsRequest
	SetServerNames(value []string) Dhcpv6ServerMetricsRequest
	// ColumnNames returns []Dhcpv6ServerMetricsRequestColumnNamesEnum, set in Dhcpv6ServerMetricsRequest
	ColumnNames() []Dhcpv6ServerMetricsRequestColumnNamesEnum
	// SetColumnNames assigns []Dhcpv6ServerMetricsRequestColumnNamesEnum provided by user to Dhcpv6ServerMetricsRequest
	SetColumnNames(value []Dhcpv6ServerMetricsRequestColumnNamesEnum) Dhcpv6ServerMetricsRequest
}

// The names of DHCPv6 Servers to return results for. An empty list will return results for all DHCPv6 Server.
//
// x-constraint:
// - /components/schemas/Device.Dhcpv6Server/properties/name
//
// ServerNames returns a []string
func (obj *dhcpv6ServerMetricsRequest) ServerNames() []string {
	if obj.obj.ServerNames == nil {
		obj.obj.ServerNames = make([]string, 0)
	}
	return obj.obj.ServerNames
}

// The names of DHCPv6 Servers to return results for. An empty list will return results for all DHCPv6 Server.
//
// x-constraint:
// - /components/schemas/Device.Dhcpv6Server/properties/name
//
// SetServerNames sets the []string value in the Dhcpv6ServerMetricsRequest object
func (obj *dhcpv6ServerMetricsRequest) SetServerNames(value []string) Dhcpv6ServerMetricsRequest {

	if obj.obj.ServerNames == nil {
		obj.obj.ServerNames = make([]string, 0)
	}
	obj.obj.ServerNames = value

	return obj
}

type Dhcpv6ServerMetricsRequestColumnNamesEnum string

// Enum of ColumnNames on Dhcpv6ServerMetricsRequest
var Dhcpv6ServerMetricsRequestColumnNames = struct {
	SOLICITS_RECEIVED             Dhcpv6ServerMetricsRequestColumnNamesEnum
	SOLICITS_IGNORED              Dhcpv6ServerMetricsRequestColumnNamesEnum
	ADVERTISEMENTS_SENT           Dhcpv6ServerMetricsRequestColumnNamesEnum
	REQUESTS_RECEIVED             Dhcpv6ServerMetricsRequestColumnNamesEnum
	NACKS_SENT                    Dhcpv6ServerMetricsRequestColumnNamesEnum
	CONFIRMS_RECEIVED             Dhcpv6ServerMetricsRequestColumnNamesEnum
	RENEWALS_RECEIVED             Dhcpv6ServerMetricsRequestColumnNamesEnum
	REBINDS_RECEIVED              Dhcpv6ServerMetricsRequestColumnNamesEnum
	REPLIES_SENT                  Dhcpv6ServerMetricsRequestColumnNamesEnum
	RELEASES_RECEIVED             Dhcpv6ServerMetricsRequestColumnNamesEnum
	DECLINES_RECEIVED             Dhcpv6ServerMetricsRequestColumnNamesEnum
	INFORMATION_REQUESTS_RECEIVED Dhcpv6ServerMetricsRequestColumnNamesEnum
	RELAY_FORWARDS_RECEIVED       Dhcpv6ServerMetricsRequestColumnNamesEnum
	RELAY_REPLIES_SENT            Dhcpv6ServerMetricsRequestColumnNamesEnum
	RECONFIGURES_SENT             Dhcpv6ServerMetricsRequestColumnNamesEnum
}{
	SOLICITS_RECEIVED:             Dhcpv6ServerMetricsRequestColumnNamesEnum("solicits_received"),
	SOLICITS_IGNORED:              Dhcpv6ServerMetricsRequestColumnNamesEnum("solicits_ignored"),
	ADVERTISEMENTS_SENT:           Dhcpv6ServerMetricsRequestColumnNamesEnum("advertisements_sent"),
	REQUESTS_RECEIVED:             Dhcpv6ServerMetricsRequestColumnNamesEnum("requests_received"),
	NACKS_SENT:                    Dhcpv6ServerMetricsRequestColumnNamesEnum("nacks_sent"),
	CONFIRMS_RECEIVED:             Dhcpv6ServerMetricsRequestColumnNamesEnum("confirms_received"),
	RENEWALS_RECEIVED:             Dhcpv6ServerMetricsRequestColumnNamesEnum("renewals_received"),
	REBINDS_RECEIVED:              Dhcpv6ServerMetricsRequestColumnNamesEnum("rebinds_received"),
	REPLIES_SENT:                  Dhcpv6ServerMetricsRequestColumnNamesEnum("replies_sent"),
	RELEASES_RECEIVED:             Dhcpv6ServerMetricsRequestColumnNamesEnum("releases_received"),
	DECLINES_RECEIVED:             Dhcpv6ServerMetricsRequestColumnNamesEnum("declines_received"),
	INFORMATION_REQUESTS_RECEIVED: Dhcpv6ServerMetricsRequestColumnNamesEnum("information_requests_received"),
	RELAY_FORWARDS_RECEIVED:       Dhcpv6ServerMetricsRequestColumnNamesEnum("relay_forwards_received"),
	RELAY_REPLIES_SENT:            Dhcpv6ServerMetricsRequestColumnNamesEnum("relay_replies_sent"),
	RECONFIGURES_SENT:             Dhcpv6ServerMetricsRequestColumnNamesEnum("reconfigures_sent"),
}

func (obj *dhcpv6ServerMetricsRequest) ColumnNames() []Dhcpv6ServerMetricsRequestColumnNamesEnum {
	items := []Dhcpv6ServerMetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, Dhcpv6ServerMetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain. If the list is empty then all columns will be returned except for any result_groups. The name of the DHCPv6 server cannot be excluded.
// SetColumnNames sets the []string value in the Dhcpv6ServerMetricsRequest object
func (obj *dhcpv6ServerMetricsRequest) SetColumnNames(value []Dhcpv6ServerMetricsRequestColumnNamesEnum) Dhcpv6ServerMetricsRequest {

	items := []otg.Dhcpv6ServerMetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.Dhcpv6ServerMetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.Dhcpv6ServerMetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *dhcpv6ServerMetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *dhcpv6ServerMetricsRequest) setDefault() {

}

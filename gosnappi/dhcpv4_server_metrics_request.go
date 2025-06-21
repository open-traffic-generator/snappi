package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv4ServerMetricsRequest *****
type dhcpv4ServerMetricsRequest struct {
	validation
	obj          *otg.Dhcpv4ServerMetricsRequest
	marshaller   marshalDhcpv4ServerMetricsRequest
	unMarshaller unMarshalDhcpv4ServerMetricsRequest
}

func NewDhcpv4ServerMetricsRequest() Dhcpv4ServerMetricsRequest {
	obj := dhcpv4ServerMetricsRequest{obj: &otg.Dhcpv4ServerMetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv4ServerMetricsRequest) msg() *otg.Dhcpv4ServerMetricsRequest {
	return obj.obj
}

func (obj *dhcpv4ServerMetricsRequest) setMsg(msg *otg.Dhcpv4ServerMetricsRequest) Dhcpv4ServerMetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv4ServerMetricsRequest struct {
	obj *dhcpv4ServerMetricsRequest
}

type marshalDhcpv4ServerMetricsRequest interface {
	// ToProto marshals Dhcpv4ServerMetricsRequest to protobuf object *otg.Dhcpv4ServerMetricsRequest
	ToProto() (*otg.Dhcpv4ServerMetricsRequest, error)
	// ToPbText marshals Dhcpv4ServerMetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv4ServerMetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv4ServerMetricsRequest to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv4ServerMetricsRequest struct {
	obj *dhcpv4ServerMetricsRequest
}

type unMarshalDhcpv4ServerMetricsRequest interface {
	// FromProto unmarshals Dhcpv4ServerMetricsRequest from protobuf object *otg.Dhcpv4ServerMetricsRequest
	FromProto(msg *otg.Dhcpv4ServerMetricsRequest) (Dhcpv4ServerMetricsRequest, error)
	// FromPbText unmarshals Dhcpv4ServerMetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv4ServerMetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv4ServerMetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *dhcpv4ServerMetricsRequest) Marshal() marshalDhcpv4ServerMetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv4ServerMetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv4ServerMetricsRequest) Unmarshal() unMarshalDhcpv4ServerMetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv4ServerMetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv4ServerMetricsRequest) ToProto() (*otg.Dhcpv4ServerMetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv4ServerMetricsRequest) FromProto(msg *otg.Dhcpv4ServerMetricsRequest) (Dhcpv4ServerMetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv4ServerMetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv4ServerMetricsRequest) FromPbText(value string) error {
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

func (m *marshaldhcpv4ServerMetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv4ServerMetricsRequest) FromYaml(value string) error {
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

func (m *marshaldhcpv4ServerMetricsRequest) ToJson() (string, error) {
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

func (m *unMarshaldhcpv4ServerMetricsRequest) FromJson(value string) error {
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

func (obj *dhcpv4ServerMetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv4ServerMetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv4ServerMetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv4ServerMetricsRequest) Clone() (Dhcpv4ServerMetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv4ServerMetricsRequest()
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

// Dhcpv4ServerMetricsRequest is the request to retrieve DHCPv4 per Server metrics/statistics.
type Dhcpv4ServerMetricsRequest interface {
	Validation
	// msg marshals Dhcpv4ServerMetricsRequest to protobuf object *otg.Dhcpv4ServerMetricsRequest
	// and doesn't set defaults
	msg() *otg.Dhcpv4ServerMetricsRequest
	// setMsg unmarshals Dhcpv4ServerMetricsRequest from protobuf object *otg.Dhcpv4ServerMetricsRequest
	// and doesn't set defaults
	setMsg(*otg.Dhcpv4ServerMetricsRequest) Dhcpv4ServerMetricsRequest
	// provides marshal interface
	Marshal() marshalDhcpv4ServerMetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv4ServerMetricsRequest
	// validate validates Dhcpv4ServerMetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv4ServerMetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ServerNames returns []string, set in Dhcpv4ServerMetricsRequest.
	ServerNames() []string
	// SetServerNames assigns []string provided by user to Dhcpv4ServerMetricsRequest
	SetServerNames(value []string) Dhcpv4ServerMetricsRequest
	// ColumnNames returns []Dhcpv4ServerMetricsRequestColumnNamesEnum, set in Dhcpv4ServerMetricsRequest
	ColumnNames() []Dhcpv4ServerMetricsRequestColumnNamesEnum
	// SetColumnNames assigns []Dhcpv4ServerMetricsRequestColumnNamesEnum provided by user to Dhcpv4ServerMetricsRequest
	SetColumnNames(value []Dhcpv4ServerMetricsRequestColumnNamesEnum) Dhcpv4ServerMetricsRequest
}

// The names of DHCPv4 Servers to return results for. An empty list will return results for all DHCPv4 Server.
//
// x-constraint:
// - /components/schemas/Device.Dhcpv4Server/properties/name
//
// ServerNames returns a []string
func (obj *dhcpv4ServerMetricsRequest) ServerNames() []string {
	if obj.obj.ServerNames == nil {
		obj.obj.ServerNames = make([]string, 0)
	}
	return obj.obj.ServerNames
}

// The names of DHCPv4 Servers to return results for. An empty list will return results for all DHCPv4 Server.
//
// x-constraint:
// - /components/schemas/Device.Dhcpv4Server/properties/name
//
// SetServerNames sets the []string value in the Dhcpv4ServerMetricsRequest object
func (obj *dhcpv4ServerMetricsRequest) SetServerNames(value []string) Dhcpv4ServerMetricsRequest {

	if obj.obj.ServerNames == nil {
		obj.obj.ServerNames = make([]string, 0)
	}
	obj.obj.ServerNames = value

	return obj
}

type Dhcpv4ServerMetricsRequestColumnNamesEnum string

// Enum of ColumnNames on Dhcpv4ServerMetricsRequest
var Dhcpv4ServerMetricsRequestColumnNames = struct {
	DISCOVERS_RECEIVED Dhcpv4ServerMetricsRequestColumnNamesEnum
	OFFERS_SENT        Dhcpv4ServerMetricsRequestColumnNamesEnum
	REQUESTS_RECEIVED  Dhcpv4ServerMetricsRequestColumnNamesEnum
	ACKS_SENT          Dhcpv4ServerMetricsRequestColumnNamesEnum
	NACKS_SENT         Dhcpv4ServerMetricsRequestColumnNamesEnum
	RELEASES_RECEIVED  Dhcpv4ServerMetricsRequestColumnNamesEnum
	DECLINES_RECEIVED  Dhcpv4ServerMetricsRequestColumnNamesEnum
}{
	DISCOVERS_RECEIVED: Dhcpv4ServerMetricsRequestColumnNamesEnum("discovers_received"),
	OFFERS_SENT:        Dhcpv4ServerMetricsRequestColumnNamesEnum("offers_sent"),
	REQUESTS_RECEIVED:  Dhcpv4ServerMetricsRequestColumnNamesEnum("requests_received"),
	ACKS_SENT:          Dhcpv4ServerMetricsRequestColumnNamesEnum("acks_sent"),
	NACKS_SENT:         Dhcpv4ServerMetricsRequestColumnNamesEnum("nacks_sent"),
	RELEASES_RECEIVED:  Dhcpv4ServerMetricsRequestColumnNamesEnum("releases_received"),
	DECLINES_RECEIVED:  Dhcpv4ServerMetricsRequestColumnNamesEnum("declines_received"),
}

func (obj *dhcpv4ServerMetricsRequest) ColumnNames() []Dhcpv4ServerMetricsRequestColumnNamesEnum {
	items := []Dhcpv4ServerMetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, Dhcpv4ServerMetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain. If the list is empty then all columns will be returned.  The name of the DHCPv4 server cannot be excluded.
// SetColumnNames sets the []string value in the Dhcpv4ServerMetricsRequest object
func (obj *dhcpv4ServerMetricsRequest) SetColumnNames(value []Dhcpv4ServerMetricsRequestColumnNamesEnum) Dhcpv4ServerMetricsRequest {

	items := []otg.Dhcpv4ServerMetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.Dhcpv4ServerMetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.Dhcpv4ServerMetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *dhcpv4ServerMetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *dhcpv4ServerMetricsRequest) setDefault() {

}

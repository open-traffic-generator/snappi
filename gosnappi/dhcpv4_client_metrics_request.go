package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv4ClientMetricsRequest *****
type dhcpv4ClientMetricsRequest struct {
	validation
	obj          *otg.Dhcpv4ClientMetricsRequest
	marshaller   marshalDhcpv4ClientMetricsRequest
	unMarshaller unMarshalDhcpv4ClientMetricsRequest
}

func NewDhcpv4ClientMetricsRequest() Dhcpv4ClientMetricsRequest {
	obj := dhcpv4ClientMetricsRequest{obj: &otg.Dhcpv4ClientMetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv4ClientMetricsRequest) msg() *otg.Dhcpv4ClientMetricsRequest {
	return obj.obj
}

func (obj *dhcpv4ClientMetricsRequest) setMsg(msg *otg.Dhcpv4ClientMetricsRequest) Dhcpv4ClientMetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv4ClientMetricsRequest struct {
	obj *dhcpv4ClientMetricsRequest
}

type marshalDhcpv4ClientMetricsRequest interface {
	// ToProto marshals Dhcpv4ClientMetricsRequest to protobuf object *otg.Dhcpv4ClientMetricsRequest
	ToProto() (*otg.Dhcpv4ClientMetricsRequest, error)
	// ToPbText marshals Dhcpv4ClientMetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv4ClientMetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv4ClientMetricsRequest to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv4ClientMetricsRequest struct {
	obj *dhcpv4ClientMetricsRequest
}

type unMarshalDhcpv4ClientMetricsRequest interface {
	// FromProto unmarshals Dhcpv4ClientMetricsRequest from protobuf object *otg.Dhcpv4ClientMetricsRequest
	FromProto(msg *otg.Dhcpv4ClientMetricsRequest) (Dhcpv4ClientMetricsRequest, error)
	// FromPbText unmarshals Dhcpv4ClientMetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv4ClientMetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv4ClientMetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *dhcpv4ClientMetricsRequest) Marshal() marshalDhcpv4ClientMetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv4ClientMetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv4ClientMetricsRequest) Unmarshal() unMarshalDhcpv4ClientMetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv4ClientMetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv4ClientMetricsRequest) ToProto() (*otg.Dhcpv4ClientMetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv4ClientMetricsRequest) FromProto(msg *otg.Dhcpv4ClientMetricsRequest) (Dhcpv4ClientMetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv4ClientMetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv4ClientMetricsRequest) FromPbText(value string) error {
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

func (m *marshaldhcpv4ClientMetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv4ClientMetricsRequest) FromYaml(value string) error {
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

func (m *marshaldhcpv4ClientMetricsRequest) ToJson() (string, error) {
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

func (m *unMarshaldhcpv4ClientMetricsRequest) FromJson(value string) error {
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

func (obj *dhcpv4ClientMetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv4ClientMetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv4ClientMetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv4ClientMetricsRequest) Clone() (Dhcpv4ClientMetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv4ClientMetricsRequest()
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

// Dhcpv4ClientMetricsRequest is the request to retrieve DHCPv4 per client metrics/statistics.
type Dhcpv4ClientMetricsRequest interface {
	Validation
	// msg marshals Dhcpv4ClientMetricsRequest to protobuf object *otg.Dhcpv4ClientMetricsRequest
	// and doesn't set defaults
	msg() *otg.Dhcpv4ClientMetricsRequest
	// setMsg unmarshals Dhcpv4ClientMetricsRequest from protobuf object *otg.Dhcpv4ClientMetricsRequest
	// and doesn't set defaults
	setMsg(*otg.Dhcpv4ClientMetricsRequest) Dhcpv4ClientMetricsRequest
	// provides marshal interface
	Marshal() marshalDhcpv4ClientMetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv4ClientMetricsRequest
	// validate validates Dhcpv4ClientMetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv4ClientMetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ClientNames returns []string, set in Dhcpv4ClientMetricsRequest.
	ClientNames() []string
	// SetClientNames assigns []string provided by user to Dhcpv4ClientMetricsRequest
	SetClientNames(value []string) Dhcpv4ClientMetricsRequest
	// ColumnNames returns []Dhcpv4ClientMetricsRequestColumnNamesEnum, set in Dhcpv4ClientMetricsRequest
	ColumnNames() []Dhcpv4ClientMetricsRequestColumnNamesEnum
	// SetColumnNames assigns []Dhcpv4ClientMetricsRequestColumnNamesEnum provided by user to Dhcpv4ClientMetricsRequest
	SetColumnNames(value []Dhcpv4ClientMetricsRequestColumnNamesEnum) Dhcpv4ClientMetricsRequest
}

// The names of DHCPv4 clients to return results for. An empty list will return results for all DHCPv4 client.
//
// x-constraint:
// - /components/schemas/Device.Dhcpv4client/properties/name
//
// ClientNames returns a []string
func (obj *dhcpv4ClientMetricsRequest) ClientNames() []string {
	if obj.obj.ClientNames == nil {
		obj.obj.ClientNames = make([]string, 0)
	}
	return obj.obj.ClientNames
}

// The names of DHCPv4 clients to return results for. An empty list will return results for all DHCPv4 client.
//
// x-constraint:
// - /components/schemas/Device.Dhcpv4client/properties/name
//
// SetClientNames sets the []string value in the Dhcpv4ClientMetricsRequest object
func (obj *dhcpv4ClientMetricsRequest) SetClientNames(value []string) Dhcpv4ClientMetricsRequest {

	if obj.obj.ClientNames == nil {
		obj.obj.ClientNames = make([]string, 0)
	}
	obj.obj.ClientNames = value

	return obj
}

type Dhcpv4ClientMetricsRequestColumnNamesEnum string

// Enum of ColumnNames on Dhcpv4ClientMetricsRequest
var Dhcpv4ClientMetricsRequestColumnNames = struct {
	DISCOVERS_SENT  Dhcpv4ClientMetricsRequestColumnNamesEnum
	OFFERS_RECEIVED Dhcpv4ClientMetricsRequestColumnNamesEnum
	REQUESTS_SENT   Dhcpv4ClientMetricsRequestColumnNamesEnum
	ACKS_RECEIVED   Dhcpv4ClientMetricsRequestColumnNamesEnum
	NACKS_RECEIVED  Dhcpv4ClientMetricsRequestColumnNamesEnum
	RELEASES_SENT   Dhcpv4ClientMetricsRequestColumnNamesEnum
	DECLINES_SENT   Dhcpv4ClientMetricsRequestColumnNamesEnum
}{
	DISCOVERS_SENT:  Dhcpv4ClientMetricsRequestColumnNamesEnum("discovers_sent"),
	OFFERS_RECEIVED: Dhcpv4ClientMetricsRequestColumnNamesEnum("offers_received"),
	REQUESTS_SENT:   Dhcpv4ClientMetricsRequestColumnNamesEnum("requests_sent"),
	ACKS_RECEIVED:   Dhcpv4ClientMetricsRequestColumnNamesEnum("acks_received"),
	NACKS_RECEIVED:  Dhcpv4ClientMetricsRequestColumnNamesEnum("nacks_received"),
	RELEASES_SENT:   Dhcpv4ClientMetricsRequestColumnNamesEnum("releases_sent"),
	DECLINES_SENT:   Dhcpv4ClientMetricsRequestColumnNamesEnum("declines_sent"),
}

func (obj *dhcpv4ClientMetricsRequest) ColumnNames() []Dhcpv4ClientMetricsRequestColumnNamesEnum {
	items := []Dhcpv4ClientMetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, Dhcpv4ClientMetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain. If the list is empty then all columns will be returned.  The name of the DHCPv4 client cannot be excluded.
// SetColumnNames sets the []string value in the Dhcpv4ClientMetricsRequest object
func (obj *dhcpv4ClientMetricsRequest) SetColumnNames(value []Dhcpv4ClientMetricsRequestColumnNamesEnum) Dhcpv4ClientMetricsRequest {

	items := []otg.Dhcpv4ClientMetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.Dhcpv4ClientMetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.Dhcpv4ClientMetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *dhcpv4ClientMetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *dhcpv4ClientMetricsRequest) setDefault() {

}

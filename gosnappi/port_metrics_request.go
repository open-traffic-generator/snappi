package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PortMetricsRequest *****
type portMetricsRequest struct {
	validation
	obj          *otg.PortMetricsRequest
	marshaller   marshalPortMetricsRequest
	unMarshaller unMarshalPortMetricsRequest
}

func NewPortMetricsRequest() PortMetricsRequest {
	obj := portMetricsRequest{obj: &otg.PortMetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *portMetricsRequest) msg() *otg.PortMetricsRequest {
	return obj.obj
}

func (obj *portMetricsRequest) setMsg(msg *otg.PortMetricsRequest) PortMetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalportMetricsRequest struct {
	obj *portMetricsRequest
}

type marshalPortMetricsRequest interface {
	// ToProto marshals PortMetricsRequest to protobuf object *otg.PortMetricsRequest
	ToProto() (*otg.PortMetricsRequest, error)
	// ToPbText marshals PortMetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PortMetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals PortMetricsRequest to JSON text
	ToJson() (string, error)
}

type unMarshalportMetricsRequest struct {
	obj *portMetricsRequest
}

type unMarshalPortMetricsRequest interface {
	// FromProto unmarshals PortMetricsRequest from protobuf object *otg.PortMetricsRequest
	FromProto(msg *otg.PortMetricsRequest) (PortMetricsRequest, error)
	// FromPbText unmarshals PortMetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PortMetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PortMetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *portMetricsRequest) Marshal() marshalPortMetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalportMetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *portMetricsRequest) Unmarshal() unMarshalPortMetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalportMetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalportMetricsRequest) ToProto() (*otg.PortMetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalportMetricsRequest) FromProto(msg *otg.PortMetricsRequest) (PortMetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalportMetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalportMetricsRequest) FromPbText(value string) error {
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

func (m *marshalportMetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalportMetricsRequest) FromYaml(value string) error {
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

func (m *marshalportMetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalportMetricsRequest) FromJson(value string) error {
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

func (obj *portMetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *portMetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *portMetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *portMetricsRequest) Clone() (PortMetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPortMetricsRequest()
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

// PortMetricsRequest is the port result request to the traffic generator
type PortMetricsRequest interface {
	Validation
	// msg marshals PortMetricsRequest to protobuf object *otg.PortMetricsRequest
	// and doesn't set defaults
	msg() *otg.PortMetricsRequest
	// setMsg unmarshals PortMetricsRequest from protobuf object *otg.PortMetricsRequest
	// and doesn't set defaults
	setMsg(*otg.PortMetricsRequest) PortMetricsRequest
	// provides marshal interface
	Marshal() marshalPortMetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalPortMetricsRequest
	// validate validates PortMetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PortMetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PortNames returns []string, set in PortMetricsRequest.
	PortNames() []string
	// SetPortNames assigns []string provided by user to PortMetricsRequest
	SetPortNames(value []string) PortMetricsRequest
	// ColumnNames returns []PortMetricsRequestColumnNamesEnum, set in PortMetricsRequest
	ColumnNames() []PortMetricsRequestColumnNamesEnum
	// SetColumnNames assigns []PortMetricsRequestColumnNamesEnum provided by user to PortMetricsRequest
	SetColumnNames(value []PortMetricsRequestColumnNamesEnum) PortMetricsRequest
}

// The names of objects to return results for. An empty list will return all port row results.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// PortNames returns a []string
func (obj *portMetricsRequest) PortNames() []string {
	if obj.obj.PortNames == nil {
		obj.obj.PortNames = make([]string, 0)
	}
	return obj.obj.PortNames
}

// The names of objects to return results for. An empty list will return all port row results.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// SetPortNames sets the []string value in the PortMetricsRequest object
func (obj *portMetricsRequest) SetPortNames(value []string) PortMetricsRequest {

	if obj.obj.PortNames == nil {
		obj.obj.PortNames = make([]string, 0)
	}
	obj.obj.PortNames = value

	return obj
}

type PortMetricsRequestColumnNamesEnum string

// Enum of ColumnNames on PortMetricsRequest
var PortMetricsRequestColumnNames = struct {
	TRANSMIT       PortMetricsRequestColumnNamesEnum
	LOCATION       PortMetricsRequestColumnNamesEnum
	LINK           PortMetricsRequestColumnNamesEnum
	CAPTURE        PortMetricsRequestColumnNamesEnum
	FRAMES_TX      PortMetricsRequestColumnNamesEnum
	FRAMES_RX      PortMetricsRequestColumnNamesEnum
	BYTES_TX       PortMetricsRequestColumnNamesEnum
	BYTES_RX       PortMetricsRequestColumnNamesEnum
	FRAMES_TX_RATE PortMetricsRequestColumnNamesEnum
	FRAMES_RX_RATE PortMetricsRequestColumnNamesEnum
	BYTES_TX_RATE  PortMetricsRequestColumnNamesEnum
	BYTES_RX_RATE  PortMetricsRequestColumnNamesEnum
	LAST_CHANGE    PortMetricsRequestColumnNamesEnum
}{
	TRANSMIT:       PortMetricsRequestColumnNamesEnum("transmit"),
	LOCATION:       PortMetricsRequestColumnNamesEnum("location"),
	LINK:           PortMetricsRequestColumnNamesEnum("link"),
	CAPTURE:        PortMetricsRequestColumnNamesEnum("capture"),
	FRAMES_TX:      PortMetricsRequestColumnNamesEnum("frames_tx"),
	FRAMES_RX:      PortMetricsRequestColumnNamesEnum("frames_rx"),
	BYTES_TX:       PortMetricsRequestColumnNamesEnum("bytes_tx"),
	BYTES_RX:       PortMetricsRequestColumnNamesEnum("bytes_rx"),
	FRAMES_TX_RATE: PortMetricsRequestColumnNamesEnum("frames_tx_rate"),
	FRAMES_RX_RATE: PortMetricsRequestColumnNamesEnum("frames_rx_rate"),
	BYTES_TX_RATE:  PortMetricsRequestColumnNamesEnum("bytes_tx_rate"),
	BYTES_RX_RATE:  PortMetricsRequestColumnNamesEnum("bytes_rx_rate"),
	LAST_CHANGE:    PortMetricsRequestColumnNamesEnum("last_change"),
}

func (obj *portMetricsRequest) ColumnNames() []PortMetricsRequestColumnNamesEnum {
	items := []PortMetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, PortMetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain. If the list is empty then all columns will be returned. The name of the port cannot be excluded.
// SetColumnNames sets the []string value in the PortMetricsRequest object
func (obj *portMetricsRequest) SetColumnNames(value []PortMetricsRequestColumnNamesEnum) PortMetricsRequest {

	items := []otg.PortMetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.PortMetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.PortMetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *portMetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *portMetricsRequest) setDefault() {

}

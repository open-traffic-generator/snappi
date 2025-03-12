package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2Flows *****
type rocev2Flows struct {
	validation
	obj              *otg.Rocev2Flows
	marshaller       marshalRocev2Flows
	unMarshaller     unMarshalRocev2Flows
	rocev2VerbHolder Rocev2Verb
}

func NewRocev2Flows() Rocev2Flows {
	obj := rocev2Flows{obj: &otg.Rocev2Flows{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2Flows) msg() *otg.Rocev2Flows {
	return obj.obj
}

func (obj *rocev2Flows) setMsg(msg *otg.Rocev2Flows) Rocev2Flows {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2Flows struct {
	obj *rocev2Flows
}

type marshalRocev2Flows interface {
	// ToProto marshals Rocev2Flows to protobuf object *otg.Rocev2Flows
	ToProto() (*otg.Rocev2Flows, error)
	// ToPbText marshals Rocev2Flows to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2Flows to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2Flows to JSON text
	ToJson() (string, error)
}

type unMarshalrocev2Flows struct {
	obj *rocev2Flows
}

type unMarshalRocev2Flows interface {
	// FromProto unmarshals Rocev2Flows from protobuf object *otg.Rocev2Flows
	FromProto(msg *otg.Rocev2Flows) (Rocev2Flows, error)
	// FromPbText unmarshals Rocev2Flows from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2Flows from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2Flows from JSON text
	FromJson(value string) error
}

func (obj *rocev2Flows) Marshal() marshalRocev2Flows {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2Flows{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2Flows) Unmarshal() unMarshalRocev2Flows {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2Flows{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2Flows) ToProto() (*otg.Rocev2Flows, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2Flows) FromProto(msg *otg.Rocev2Flows) (Rocev2Flows, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2Flows) ToPbText() (string, error) {
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

func (m *unMarshalrocev2Flows) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalrocev2Flows) ToYaml() (string, error) {
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

func (m *unMarshalrocev2Flows) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalrocev2Flows) ToJson() (string, error) {
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

func (m *unMarshalrocev2Flows) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *rocev2Flows) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2Flows) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2Flows) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2Flows) Clone() (Rocev2Flows, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2Flows()
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

func (obj *rocev2Flows) setNil() {
	obj.rocev2VerbHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2Flows is ****Decription here****
type Rocev2Flows interface {
	Validation
	// msg marshals Rocev2Flows to protobuf object *otg.Rocev2Flows
	// and doesn't set defaults
	msg() *otg.Rocev2Flows
	// setMsg unmarshals Rocev2Flows from protobuf object *otg.Rocev2Flows
	// and doesn't set defaults
	setMsg(*otg.Rocev2Flows) Rocev2Flows
	// provides marshal interface
	Marshal() marshalRocev2Flows
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2Flows
	// validate validates Rocev2Flows
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2Flows, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TxEndpoint returns string, set in Rocev2Flows.
	TxEndpoint() string
	// SetTxEndpoint assigns string provided by user to Rocev2Flows
	SetTxEndpoint(value string) Rocev2Flows
	// RxEndpoint returns string, set in Rocev2Flows.
	RxEndpoint() string
	// SetRxEndpoint assigns string provided by user to Rocev2Flows
	SetRxEndpoint(value string) Rocev2Flows
	// HasRxEndpoint checks if RxEndpoint has been set in Rocev2Flows
	HasRxEndpoint() bool
	// Name returns string, set in Rocev2Flows.
	Name() string
	// SetName assigns string provided by user to Rocev2Flows
	SetName(value string) Rocev2Flows
	// Rocev2Verb returns Rocev2Verb, set in Rocev2Flows.
	// Rocev2Verb is rocev2 Verb, Available options are: write, wrtie_with_immediate, send, send_with_immediate and read.
	Rocev2Verb() Rocev2Verb
	// SetRocev2Verb assigns Rocev2Verb provided by user to Rocev2Flows.
	// Rocev2Verb is rocev2 Verb, Available options are: write, wrtie_with_immediate, send, send_with_immediate and read.
	SetRocev2Verb(value Rocev2Verb) Rocev2Flows
	// HasRocev2Verb checks if Rocev2Verb has been set in Rocev2Flows
	HasRocev2Verb() bool
	// MessageSize returns uint32, set in Rocev2Flows.
	MessageSize() uint32
	// SetMessageSize assigns uint32 provided by user to Rocev2Flows
	SetMessageSize(value uint32) Rocev2Flows
	// HasMessageSize checks if MessageSize has been set in Rocev2Flows
	HasMessageSize() bool
	// MessageSizeUnit returns Rocev2FlowsMessageSizeUnitEnum, set in Rocev2Flows
	MessageSizeUnit() Rocev2FlowsMessageSizeUnitEnum
	// SetMessageSizeUnit assigns Rocev2FlowsMessageSizeUnitEnum provided by user to Rocev2Flows
	SetMessageSizeUnit(value Rocev2FlowsMessageSizeUnitEnum) Rocev2Flows
	// HasMessageSizeUnit checks if MessageSizeUnit has been set in Rocev2Flows
	HasMessageSizeUnit() bool
	setNil()
}

// The unique name of an emulated device that will be transmitting.
//
// x-constraint:
// - /components/schemas/Rocev2.QPs/properties/qp_name
//
// TxEndpoint returns a string
func (obj *rocev2Flows) TxEndpoint() string {

	return *obj.obj.TxEndpoint

}

// The unique name of an emulated device that will be transmitting.
//
// x-constraint:
// - /components/schemas/Rocev2.QPs/properties/qp_name
//
// SetTxEndpoint sets the string value in the Rocev2Flows object
func (obj *rocev2Flows) SetTxEndpoint(value string) Rocev2Flows {

	obj.obj.TxEndpoint = &value
	return obj
}

// The unique name of an emulated device that will be recieving.
//
// x-constraint:
// - /components/schemas/Port/properties/name
// - /components/schemas/Rocev2.QPs/properties/qp_name
//
// RxEndpoint returns a string
func (obj *rocev2Flows) RxEndpoint() string {

	return *obj.obj.RxEndpoint

}

// The unique name of an emulated device that will be recieving.
//
// x-constraint:
// - /components/schemas/Port/properties/name
// - /components/schemas/Rocev2.QPs/properties/qp_name
//
// RxEndpoint returns a string
func (obj *rocev2Flows) HasRxEndpoint() bool {
	return obj.obj.RxEndpoint != nil
}

// The unique name of an emulated device that will be recieving.
//
// x-constraint:
// - /components/schemas/Port/properties/name
// - /components/schemas/Rocev2.QPs/properties/qp_name
//
// SetRxEndpoint sets the string value in the Rocev2Flows object
func (obj *rocev2Flows) SetRxEndpoint(value string) Rocev2Flows {

	obj.obj.RxEndpoint = &value
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *rocev2Flows) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the Rocev2Flows object
func (obj *rocev2Flows) SetName(value string) Rocev2Flows {

	obj.obj.Name = &value
	return obj
}

// description is TBD
// Rocev2Verb returns a Rocev2Verb
func (obj *rocev2Flows) Rocev2Verb() Rocev2Verb {
	if obj.obj.Rocev2Verb == nil {
		obj.obj.Rocev2Verb = NewRocev2Verb().msg()
	}
	if obj.rocev2VerbHolder == nil {
		obj.rocev2VerbHolder = &rocev2Verb{obj: obj.obj.Rocev2Verb}
	}
	return obj.rocev2VerbHolder
}

// description is TBD
// Rocev2Verb returns a Rocev2Verb
func (obj *rocev2Flows) HasRocev2Verb() bool {
	return obj.obj.Rocev2Verb != nil
}

// description is TBD
// SetRocev2Verb sets the Rocev2Verb value in the Rocev2Flows object
func (obj *rocev2Flows) SetRocev2Verb(value Rocev2Verb) Rocev2Flows {

	obj.rocev2VerbHolder = nil
	obj.obj.Rocev2Verb = value.msg()

	return obj
}

// The Maximum message size that is allowed to transfer depends on the MTU size and the number of VLANs configured on the interfaces.
// MessageSize returns a uint32
func (obj *rocev2Flows) MessageSize() uint32 {

	return *obj.obj.MessageSize

}

// The Maximum message size that is allowed to transfer depends on the MTU size and the number of VLANs configured on the interfaces.
// MessageSize returns a uint32
func (obj *rocev2Flows) HasMessageSize() bool {
	return obj.obj.MessageSize != nil
}

// The Maximum message size that is allowed to transfer depends on the MTU size and the number of VLANs configured on the interfaces.
// SetMessageSize sets the uint32 value in the Rocev2Flows object
func (obj *rocev2Flows) SetMessageSize(value uint32) Rocev2Flows {

	obj.obj.MessageSize = &value
	return obj
}

type Rocev2FlowsMessageSizeUnitEnum string

// Enum of MessageSizeUnit on Rocev2Flows
var Rocev2FlowsMessageSizeUnit = struct {
	BYTE Rocev2FlowsMessageSizeUnitEnum
	KB   Rocev2FlowsMessageSizeUnitEnum
	MB   Rocev2FlowsMessageSizeUnitEnum
	GB   Rocev2FlowsMessageSizeUnitEnum
}{
	BYTE: Rocev2FlowsMessageSizeUnitEnum("Byte"),
	KB:   Rocev2FlowsMessageSizeUnitEnum("KB"),
	MB:   Rocev2FlowsMessageSizeUnitEnum("MB"),
	GB:   Rocev2FlowsMessageSizeUnitEnum("GB"),
}

func (obj *rocev2Flows) MessageSizeUnit() Rocev2FlowsMessageSizeUnitEnum {
	return Rocev2FlowsMessageSizeUnitEnum(obj.obj.MessageSizeUnit.Enum().String())
}

// Unit of the transfer message size. Available options are Bytes, KB, MB.
// MessageSizeUnit returns a string
func (obj *rocev2Flows) HasMessageSizeUnit() bool {
	return obj.obj.MessageSizeUnit != nil
}

func (obj *rocev2Flows) SetMessageSizeUnit(value Rocev2FlowsMessageSizeUnitEnum) Rocev2Flows {
	intValue, ok := otg.Rocev2Flows_MessageSizeUnit_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Rocev2FlowsMessageSizeUnitEnum", string(value)))
		return obj
	}
	enumValue := otg.Rocev2Flows_MessageSizeUnit_Enum(intValue)
	obj.obj.MessageSizeUnit = &enumValue

	return obj
}

func (obj *rocev2Flows) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// TxEndpoint is required
	if obj.obj.TxEndpoint == nil {
		vObj.validationErrors = append(vObj.validationErrors, "TxEndpoint is required field on interface Rocev2Flows")
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface Rocev2Flows")
	}

	if obj.obj.Rocev2Verb != nil {

		obj.Rocev2Verb().validateObj(vObj, set_default)
	}

	if obj.obj.MessageSize != nil {

		if *obj.obj.MessageSize > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2Flows.MessageSize <= 65535 but Got %d", *obj.obj.MessageSize))
		}

	}

}

func (obj *rocev2Flows) setDefault() {
	if obj.obj.MessageSize == nil {
		obj.SetMessageSize(1)
	}
	if obj.obj.MessageSizeUnit == nil {
		obj.SetMessageSizeUnit(Rocev2FlowsMessageSizeUnit.MB)

	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2Flow *****
type rocev2Flow struct {
	validation
	obj              *otg.Rocev2Flow
	marshaller       marshalRocev2Flow
	unMarshaller     unMarshalRocev2Flow
	rocev2VerbHolder Rocev2Verb
}

func NewRocev2Flow() Rocev2Flow {
	obj := rocev2Flow{obj: &otg.Rocev2Flow{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2Flow) msg() *otg.Rocev2Flow {
	return obj.obj
}

func (obj *rocev2Flow) setMsg(msg *otg.Rocev2Flow) Rocev2Flow {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2Flow struct {
	obj *rocev2Flow
}

type marshalRocev2Flow interface {
	// ToProto marshals Rocev2Flow to protobuf object *otg.Rocev2Flow
	ToProto() (*otg.Rocev2Flow, error)
	// ToPbText marshals Rocev2Flow to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2Flow to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2Flow to JSON text
	ToJson() (string, error)
}

type unMarshalrocev2Flow struct {
	obj *rocev2Flow
}

type unMarshalRocev2Flow interface {
	// FromProto unmarshals Rocev2Flow from protobuf object *otg.Rocev2Flow
	FromProto(msg *otg.Rocev2Flow) (Rocev2Flow, error)
	// FromPbText unmarshals Rocev2Flow from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2Flow from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2Flow from JSON text
	FromJson(value string) error
}

func (obj *rocev2Flow) Marshal() marshalRocev2Flow {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2Flow{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2Flow) Unmarshal() unMarshalRocev2Flow {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2Flow{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2Flow) ToProto() (*otg.Rocev2Flow, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2Flow) FromProto(msg *otg.Rocev2Flow) (Rocev2Flow, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2Flow) ToPbText() (string, error) {
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

func (m *unMarshalrocev2Flow) FromPbText(value string) error {
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

func (m *marshalrocev2Flow) ToYaml() (string, error) {
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

func (m *unMarshalrocev2Flow) FromYaml(value string) error {
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

func (m *marshalrocev2Flow) ToJson() (string, error) {
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

func (m *unMarshalrocev2Flow) FromJson(value string) error {
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

func (obj *rocev2Flow) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2Flow) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2Flow) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2Flow) Clone() (Rocev2Flow, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2Flow()
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

func (obj *rocev2Flow) setNil() {
	obj.rocev2VerbHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2Flow is configure properties for a specific RoCE flow on the Tx port.
type Rocev2Flow interface {
	Validation
	// msg marshals Rocev2Flow to protobuf object *otg.Rocev2Flow
	// and doesn't set defaults
	msg() *otg.Rocev2Flow
	// setMsg unmarshals Rocev2Flow from protobuf object *otg.Rocev2Flow
	// and doesn't set defaults
	setMsg(*otg.Rocev2Flow) Rocev2Flow
	// provides marshal interface
	Marshal() marshalRocev2Flow
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2Flow
	// validate validates Rocev2Flow
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2Flow, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TxEndpoint returns string, set in Rocev2Flow.
	TxEndpoint() string
	// SetTxEndpoint assigns string provided by user to Rocev2Flow
	SetTxEndpoint(value string) Rocev2Flow
	// RxEndpoint returns string, set in Rocev2Flow.
	RxEndpoint() string
	// SetRxEndpoint assigns string provided by user to Rocev2Flow
	SetRxEndpoint(value string) Rocev2Flow
	// HasRxEndpoint checks if RxEndpoint has been set in Rocev2Flow
	HasRxEndpoint() bool
	// Name returns string, set in Rocev2Flow.
	Name() string
	// SetName assigns string provided by user to Rocev2Flow
	SetName(value string) Rocev2Flow
	// Rocev2Verb returns Rocev2Verb, set in Rocev2Flow.
	// Rocev2Verb is roCEv2 Verb. Available options are: WRITE, WRITE_With_Immediate, SEND, SEND_With_Immediate and READ.
	Rocev2Verb() Rocev2Verb
	// SetRocev2Verb assigns Rocev2Verb provided by user to Rocev2Flow.
	// Rocev2Verb is roCEv2 Verb. Available options are: WRITE, WRITE_With_Immediate, SEND, SEND_With_Immediate and READ.
	SetRocev2Verb(value Rocev2Verb) Rocev2Flow
	// HasRocev2Verb checks if Rocev2Verb has been set in Rocev2Flow
	HasRocev2Verb() bool
	// MessageSize returns uint32, set in Rocev2Flow.
	MessageSize() uint32
	// SetMessageSize assigns uint32 provided by user to Rocev2Flow
	SetMessageSize(value uint32) Rocev2Flow
	// HasMessageSize checks if MessageSize has been set in Rocev2Flow
	HasMessageSize() bool
	// MessageSizeUnit returns Rocev2FlowMessageSizeUnitEnum, set in Rocev2Flow
	MessageSizeUnit() Rocev2FlowMessageSizeUnitEnum
	// SetMessageSizeUnit assigns Rocev2FlowMessageSizeUnitEnum provided by user to Rocev2Flow
	SetMessageSizeUnit(value Rocev2FlowMessageSizeUnitEnum) Rocev2Flow
	// HasMessageSizeUnit checks if MessageSizeUnit has been set in Rocev2Flow
	HasMessageSizeUnit() bool
	setNil()
}

// The unique name of an emulated device that will be transmitting the flows.
//
// x-constraint:
// - /components/schemas/Rocev2.QPs/properties/qp_name
//
// TxEndpoint returns a string
func (obj *rocev2Flow) TxEndpoint() string {

	return *obj.obj.TxEndpoint

}

// The unique name of an emulated device that will be transmitting the flows.
//
// x-constraint:
// - /components/schemas/Rocev2.QPs/properties/qp_name
//
// SetTxEndpoint sets the string value in the Rocev2Flow object
func (obj *rocev2Flow) SetTxEndpoint(value string) Rocev2Flow {

	obj.obj.TxEndpoint = &value
	return obj
}

// The unique name of remote QP or port which be receiving the packets for the flow.
//
// x-constraint:
// - /components/schemas/Port/properties/name
// - /components/schemas/Rocev2.QPs/properties/qp_name
//
// RxEndpoint returns a string
func (obj *rocev2Flow) RxEndpoint() string {

	return *obj.obj.RxEndpoint

}

// The unique name of remote QP or port which be receiving the packets for the flow.
//
// x-constraint:
// - /components/schemas/Port/properties/name
// - /components/schemas/Rocev2.QPs/properties/qp_name
//
// RxEndpoint returns a string
func (obj *rocev2Flow) HasRxEndpoint() bool {
	return obj.obj.RxEndpoint != nil
}

// The unique name of remote QP or port which be receiving the packets for the flow.
//
// x-constraint:
// - /components/schemas/Port/properties/name
// - /components/schemas/Rocev2.QPs/properties/qp_name
//
// SetRxEndpoint sets the string value in the Rocev2Flow object
func (obj *rocev2Flow) SetRxEndpoint(value string) Rocev2Flow {

	obj.obj.RxEndpoint = &value
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *rocev2Flow) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the Rocev2Flow object
func (obj *rocev2Flow) SetName(value string) Rocev2Flow {

	obj.obj.Name = &value
	return obj
}

// description is TBD
// Rocev2Verb returns a Rocev2Verb
func (obj *rocev2Flow) Rocev2Verb() Rocev2Verb {
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
func (obj *rocev2Flow) HasRocev2Verb() bool {
	return obj.obj.Rocev2Verb != nil
}

// description is TBD
// SetRocev2Verb sets the Rocev2Verb value in the Rocev2Flow object
func (obj *rocev2Flow) SetRocev2Verb(value Rocev2Verb) Rocev2Flow {

	obj.rocev2VerbHolder = nil
	obj.obj.Rocev2Verb = value.msg()

	return obj
}

// Length of Message that needs to be transmitted to the remote end-point.
// MessageSize returns a uint32
func (obj *rocev2Flow) MessageSize() uint32 {

	return *obj.obj.MessageSize

}

// Length of Message that needs to be transmitted to the remote end-point.
// MessageSize returns a uint32
func (obj *rocev2Flow) HasMessageSize() bool {
	return obj.obj.MessageSize != nil
}

// Length of Message that needs to be transmitted to the remote end-point.
// SetMessageSize sets the uint32 value in the Rocev2Flow object
func (obj *rocev2Flow) SetMessageSize(value uint32) Rocev2Flow {

	obj.obj.MessageSize = &value
	return obj
}

type Rocev2FlowMessageSizeUnitEnum string

// Enum of MessageSizeUnit on Rocev2Flow
var Rocev2FlowMessageSizeUnit = struct {
	BYTES Rocev2FlowMessageSizeUnitEnum
	KB    Rocev2FlowMessageSizeUnitEnum
	MB    Rocev2FlowMessageSizeUnitEnum
	GB    Rocev2FlowMessageSizeUnitEnum
}{
	BYTES: Rocev2FlowMessageSizeUnitEnum("bytes"),
	KB:    Rocev2FlowMessageSizeUnitEnum("kb"),
	MB:    Rocev2FlowMessageSizeUnitEnum("mb"),
	GB:    Rocev2FlowMessageSizeUnitEnum("gb"),
}

func (obj *rocev2Flow) MessageSizeUnit() Rocev2FlowMessageSizeUnitEnum {
	return Rocev2FlowMessageSizeUnitEnum(obj.obj.MessageSizeUnit.Enum().String())
}

// Unit of the transfer message size. Available options are Bytes, KiloBtyes (KB), NegaBytes (MB) and GigaBytes (GB).
// MessageSizeUnit returns a string
func (obj *rocev2Flow) HasMessageSizeUnit() bool {
	return obj.obj.MessageSizeUnit != nil
}

func (obj *rocev2Flow) SetMessageSizeUnit(value Rocev2FlowMessageSizeUnitEnum) Rocev2Flow {
	intValue, ok := otg.Rocev2Flow_MessageSizeUnit_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Rocev2FlowMessageSizeUnitEnum", string(value)))
		return obj
	}
	enumValue := otg.Rocev2Flow_MessageSizeUnit_Enum(intValue)
	obj.obj.MessageSizeUnit = &enumValue

	return obj
}

func (obj *rocev2Flow) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// TxEndpoint is required
	if obj.obj.TxEndpoint == nil {
		vObj.validationErrors = append(vObj.validationErrors, "TxEndpoint is required field on interface Rocev2Flow")
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface Rocev2Flow")
	}

	if obj.obj.Rocev2Verb != nil {

		obj.Rocev2Verb().validateObj(vObj, set_default)
	}

	if obj.obj.MessageSize != nil {

		if *obj.obj.MessageSize > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2Flow.MessageSize <= 65535 but Got %d", *obj.obj.MessageSize))
		}

	}

}

func (obj *rocev2Flow) setDefault() {
	if obj.obj.MessageSize == nil {
		obj.SetMessageSize(1)
	}
	if obj.obj.MessageSizeUnit == nil {
		obj.SetMessageSizeUnit(Rocev2FlowMessageSizeUnit.MB)

	}

}

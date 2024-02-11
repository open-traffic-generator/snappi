package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowDelay *****
type flowDelay struct {
	validation
	obj          *otg.FlowDelay
	marshaller   marshalFlowDelay
	unMarshaller unMarshalFlowDelay
}

func NewFlowDelay() FlowDelay {
	obj := flowDelay{obj: &otg.FlowDelay{}}
	obj.setDefault()
	return &obj
}

func (obj *flowDelay) msg() *otg.FlowDelay {
	return obj.obj
}

func (obj *flowDelay) setMsg(msg *otg.FlowDelay) FlowDelay {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowDelay struct {
	obj *flowDelay
}

type marshalFlowDelay interface {
	// ToProto marshals FlowDelay to protobuf object *otg.FlowDelay
	ToProto() (*otg.FlowDelay, error)
	// ToPbText marshals FlowDelay to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowDelay to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowDelay to JSON text
	ToJson() (string, error)
}

type unMarshalflowDelay struct {
	obj *flowDelay
}

type unMarshalFlowDelay interface {
	// FromProto unmarshals FlowDelay from protobuf object *otg.FlowDelay
	FromProto(msg *otg.FlowDelay) (FlowDelay, error)
	// FromPbText unmarshals FlowDelay from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowDelay from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowDelay from JSON text
	FromJson(value string) error
}

func (obj *flowDelay) Marshal() marshalFlowDelay {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowDelay{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowDelay) Unmarshal() unMarshalFlowDelay {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowDelay{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowDelay) ToProto() (*otg.FlowDelay, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowDelay) FromProto(msg *otg.FlowDelay) (FlowDelay, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowDelay) ToPbText() (string, error) {
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

func (m *unMarshalflowDelay) FromPbText(value string) error {
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

func (m *marshalflowDelay) ToYaml() (string, error) {
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

func (m *unMarshalflowDelay) FromYaml(value string) error {
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

func (m *marshalflowDelay) ToJson() (string, error) {
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

func (m *unMarshalflowDelay) FromJson(value string) error {
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

func (obj *flowDelay) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowDelay) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowDelay) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowDelay) Clone() (FlowDelay, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowDelay()
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

// FlowDelay is the optional container to specify the delay before starting
// transmission of packets.
type FlowDelay interface {
	Validation
	// msg marshals FlowDelay to protobuf object *otg.FlowDelay
	// and doesn't set defaults
	msg() *otg.FlowDelay
	// setMsg unmarshals FlowDelay from protobuf object *otg.FlowDelay
	// and doesn't set defaults
	setMsg(*otg.FlowDelay) FlowDelay
	// provides marshal interface
	Marshal() marshalFlowDelay
	// provides unmarshal interface
	Unmarshal() unMarshalFlowDelay
	// validate validates FlowDelay
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowDelay, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowDelayChoiceEnum, set in FlowDelay
	Choice() FlowDelayChoiceEnum
	// setChoice assigns FlowDelayChoiceEnum provided by user to FlowDelay
	setChoice(value FlowDelayChoiceEnum) FlowDelay
	// HasChoice checks if Choice has been set in FlowDelay
	HasChoice() bool
	// Bytes returns float32, set in FlowDelay.
	Bytes() float32
	// SetBytes assigns float32 provided by user to FlowDelay
	SetBytes(value float32) FlowDelay
	// HasBytes checks if Bytes has been set in FlowDelay
	HasBytes() bool
	// Nanoseconds returns float32, set in FlowDelay.
	Nanoseconds() float32
	// SetNanoseconds assigns float32 provided by user to FlowDelay
	SetNanoseconds(value float32) FlowDelay
	// HasNanoseconds checks if Nanoseconds has been set in FlowDelay
	HasNanoseconds() bool
	// Microseconds returns float32, set in FlowDelay.
	Microseconds() float32
	// SetMicroseconds assigns float32 provided by user to FlowDelay
	SetMicroseconds(value float32) FlowDelay
	// HasMicroseconds checks if Microseconds has been set in FlowDelay
	HasMicroseconds() bool
}

type FlowDelayChoiceEnum string

// Enum of Choice on FlowDelay
var FlowDelayChoice = struct {
	BYTES        FlowDelayChoiceEnum
	NANOSECONDS  FlowDelayChoiceEnum
	MICROSECONDS FlowDelayChoiceEnum
}{
	BYTES:        FlowDelayChoiceEnum("bytes"),
	NANOSECONDS:  FlowDelayChoiceEnum("nanoseconds"),
	MICROSECONDS: FlowDelayChoiceEnum("microseconds"),
}

func (obj *flowDelay) Choice() FlowDelayChoiceEnum {
	return FlowDelayChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowDelay) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowDelay) setChoice(value FlowDelayChoiceEnum) FlowDelay {
	intValue, ok := otg.FlowDelay_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowDelayChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowDelay_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Microseconds = nil
	obj.obj.Nanoseconds = nil
	obj.obj.Bytes = nil

	if value == FlowDelayChoice.BYTES {
		defaultValue := float32(0)
		obj.obj.Bytes = &defaultValue
	}

	if value == FlowDelayChoice.NANOSECONDS {
		defaultValue := float32(0)
		obj.obj.Nanoseconds = &defaultValue
	}

	if value == FlowDelayChoice.MICROSECONDS {
		defaultValue := float32(0)
		obj.obj.Microseconds = &defaultValue
	}

	return obj
}

// The delay before starting transmission of packets.
// A value of 0 indicates no delay.
// Bytes returns a float32
func (obj *flowDelay) Bytes() float32 {

	if obj.obj.Bytes == nil {
		obj.setChoice(FlowDelayChoice.BYTES)
	}

	return *obj.obj.Bytes

}

// The delay before starting transmission of packets.
// A value of 0 indicates no delay.
// Bytes returns a float32
func (obj *flowDelay) HasBytes() bool {
	return obj.obj.Bytes != nil
}

// The delay before starting transmission of packets.
// A value of 0 indicates no delay.
// SetBytes sets the float32 value in the FlowDelay object
func (obj *flowDelay) SetBytes(value float32) FlowDelay {
	obj.setChoice(FlowDelayChoice.BYTES)
	obj.obj.Bytes = &value
	return obj
}

// The delay before starting transmission of packets.
// A value of 0 indicates no delay.
// Nanoseconds returns a float32
func (obj *flowDelay) Nanoseconds() float32 {

	if obj.obj.Nanoseconds == nil {
		obj.setChoice(FlowDelayChoice.NANOSECONDS)
	}

	return *obj.obj.Nanoseconds

}

// The delay before starting transmission of packets.
// A value of 0 indicates no delay.
// Nanoseconds returns a float32
func (obj *flowDelay) HasNanoseconds() bool {
	return obj.obj.Nanoseconds != nil
}

// The delay before starting transmission of packets.
// A value of 0 indicates no delay.
// SetNanoseconds sets the float32 value in the FlowDelay object
func (obj *flowDelay) SetNanoseconds(value float32) FlowDelay {
	obj.setChoice(FlowDelayChoice.NANOSECONDS)
	obj.obj.Nanoseconds = &value
	return obj
}

// The delay before starting transmission of packets.
// A value of 0 indicates no delay.
// Microseconds returns a float32
func (obj *flowDelay) Microseconds() float32 {

	if obj.obj.Microseconds == nil {
		obj.setChoice(FlowDelayChoice.MICROSECONDS)
	}

	return *obj.obj.Microseconds

}

// The delay before starting transmission of packets.
// A value of 0 indicates no delay.
// Microseconds returns a float32
func (obj *flowDelay) HasMicroseconds() bool {
	return obj.obj.Microseconds != nil
}

// The delay before starting transmission of packets.
// A value of 0 indicates no delay.
// SetMicroseconds sets the float32 value in the FlowDelay object
func (obj *flowDelay) SetMicroseconds(value float32) FlowDelay {
	obj.setChoice(FlowDelayChoice.MICROSECONDS)
	obj.obj.Microseconds = &value
	return obj
}

func (obj *flowDelay) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Bytes != nil {

		if *obj.obj.Bytes < 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowDelay.Bytes <= max(float32) but Got %f", *obj.obj.Bytes))
		}

	}

	if obj.obj.Nanoseconds != nil {

		if *obj.obj.Nanoseconds < 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowDelay.Nanoseconds <= max(float32) but Got %f", *obj.obj.Nanoseconds))
		}

	}

	if obj.obj.Microseconds != nil {

		if *obj.obj.Microseconds < 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowDelay.Microseconds <= max(float32) but Got %f", *obj.obj.Microseconds))
		}

	}

}

func (obj *flowDelay) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(FlowDelayChoice.BYTES)

	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowPayload *****
type flowPayload struct {
	validation
	obj          *otg.FlowPayload
	marshaller   marshalFlowPayload
	unMarshaller unMarshalFlowPayload
	fixedHolder  FlowPayloadFixed
}

func NewFlowPayload() FlowPayload {
	obj := flowPayload{obj: &otg.FlowPayload{}}
	obj.setDefault()
	return &obj
}

func (obj *flowPayload) msg() *otg.FlowPayload {
	return obj.obj
}

func (obj *flowPayload) setMsg(msg *otg.FlowPayload) FlowPayload {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowPayload struct {
	obj *flowPayload
}

type marshalFlowPayload interface {
	// ToProto marshals FlowPayload to protobuf object *otg.FlowPayload
	ToProto() (*otg.FlowPayload, error)
	// ToPbText marshals FlowPayload to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowPayload to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowPayload to JSON text
	ToJson() (string, error)
}

type unMarshalflowPayload struct {
	obj *flowPayload
}

type unMarshalFlowPayload interface {
	// FromProto unmarshals FlowPayload from protobuf object *otg.FlowPayload
	FromProto(msg *otg.FlowPayload) (FlowPayload, error)
	// FromPbText unmarshals FlowPayload from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowPayload from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowPayload from JSON text
	FromJson(value string) error
}

func (obj *flowPayload) Marshal() marshalFlowPayload {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowPayload{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowPayload) Unmarshal() unMarshalFlowPayload {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowPayload{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowPayload) ToProto() (*otg.FlowPayload, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowPayload) FromProto(msg *otg.FlowPayload) (FlowPayload, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowPayload) ToPbText() (string, error) {
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

func (m *unMarshalflowPayload) FromPbText(value string) error {
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

func (m *marshalflowPayload) ToYaml() (string, error) {
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

func (m *unMarshalflowPayload) FromYaml(value string) error {
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

func (m *marshalflowPayload) ToJson() (string, error) {
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

func (m *unMarshalflowPayload) FromJson(value string) error {
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

func (obj *flowPayload) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowPayload) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowPayload) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowPayload) Clone() (FlowPayload, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowPayload()
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

func (obj *flowPayload) setNil() {
	obj.fixedHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowPayload is a container for different types of payload, which is the data in the frame after protocol headers.
// Some part of the payload could be overwritten with instrumentation data, contents and placement of  which is implementation specific.
type FlowPayload interface {
	Validation
	// msg marshals FlowPayload to protobuf object *otg.FlowPayload
	// and doesn't set defaults
	msg() *otg.FlowPayload
	// setMsg unmarshals FlowPayload from protobuf object *otg.FlowPayload
	// and doesn't set defaults
	setMsg(*otg.FlowPayload) FlowPayload
	// provides marshal interface
	Marshal() marshalFlowPayload
	// provides unmarshal interface
	Unmarshal() unMarshalFlowPayload
	// validate validates FlowPayload
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowPayload, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowPayloadChoiceEnum, set in FlowPayload
	Choice() FlowPayloadChoiceEnum
	// setChoice assigns FlowPayloadChoiceEnum provided by user to FlowPayload
	setChoice(value FlowPayloadChoiceEnum) FlowPayload
	// HasChoice checks if Choice has been set in FlowPayload
	HasChoice() bool
	// getter for DecrementWord to set choice.
	DecrementWord()
	// getter for DecrementByte to set choice.
	DecrementByte()
	// getter for IncrementWord to set choice.
	IncrementWord()
	// getter for IncrementByte to set choice.
	IncrementByte()
	// Fixed returns FlowPayloadFixed, set in FlowPayload.
	// FlowPayloadFixed is payload with user defined pattern.
	Fixed() FlowPayloadFixed
	// SetFixed assigns FlowPayloadFixed provided by user to FlowPayload.
	// FlowPayloadFixed is payload with user defined pattern.
	SetFixed(value FlowPayloadFixed) FlowPayload
	// HasFixed checks if Fixed has been set in FlowPayload
	HasFixed() bool
	setNil()
}

type FlowPayloadChoiceEnum string

// Enum of Choice on FlowPayload
var FlowPayloadChoice = struct {
	FIXED          FlowPayloadChoiceEnum
	INCREMENT_BYTE FlowPayloadChoiceEnum
	DECREMENT_BYTE FlowPayloadChoiceEnum
	INCREMENT_WORD FlowPayloadChoiceEnum
	DECREMENT_WORD FlowPayloadChoiceEnum
}{
	FIXED:          FlowPayloadChoiceEnum("fixed"),
	INCREMENT_BYTE: FlowPayloadChoiceEnum("increment_byte"),
	DECREMENT_BYTE: FlowPayloadChoiceEnum("decrement_byte"),
	INCREMENT_WORD: FlowPayloadChoiceEnum("increment_word"),
	DECREMENT_WORD: FlowPayloadChoiceEnum("decrement_word"),
}

func (obj *flowPayload) Choice() FlowPayloadChoiceEnum {
	return FlowPayloadChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for DecrementWord to set choice
func (obj *flowPayload) DecrementWord() {
	obj.setChoice(FlowPayloadChoice.DECREMENT_WORD)
}

// getter for DecrementByte to set choice
func (obj *flowPayload) DecrementByte() {
	obj.setChoice(FlowPayloadChoice.DECREMENT_BYTE)
}

// getter for IncrementWord to set choice
func (obj *flowPayload) IncrementWord() {
	obj.setChoice(FlowPayloadChoice.INCREMENT_WORD)
}

// getter for IncrementByte to set choice
func (obj *flowPayload) IncrementByte() {
	obj.setChoice(FlowPayloadChoice.INCREMENT_BYTE)
}

// A choice used to determine the pattern of the bytes in the payload following the protocol headers.
// Example of expected behaviour of each choice is as mentioned below,
// - fixed: this would insert user defined pattern in the payload bytes.
// - increment_byte: this would set the first byte as 00, second as 01 and so on, upto ff and then the pattern would be repeated.
// - decrement_byte: this would set the first byte as ff, second as fe and so on, upto 00 and then the pattern would be repeated.
// - increment_word: this would set the first 2 bytes as 00 00, second as 00 01 and so on, upto ff ff and then the pattern would be repeated.
// - decrement_word: this would set the first 2 bytes as ff ff, second as ff fe and so on, upto 00 00 and then the pattern would be repeated.
// Choice returns a string
func (obj *flowPayload) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowPayload) setChoice(value FlowPayloadChoiceEnum) FlowPayload {
	intValue, ok := otg.FlowPayload_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowPayloadChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowPayload_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Fixed = nil
	obj.fixedHolder = nil

	if value == FlowPayloadChoice.FIXED {
		obj.obj.Fixed = NewFlowPayloadFixed().msg()
	}

	return obj
}

// description is TBD
// Fixed returns a FlowPayloadFixed
func (obj *flowPayload) Fixed() FlowPayloadFixed {
	if obj.obj.Fixed == nil {
		obj.setChoice(FlowPayloadChoice.FIXED)
	}
	if obj.fixedHolder == nil {
		obj.fixedHolder = &flowPayloadFixed{obj: obj.obj.Fixed}
	}
	return obj.fixedHolder
}

// description is TBD
// Fixed returns a FlowPayloadFixed
func (obj *flowPayload) HasFixed() bool {
	return obj.obj.Fixed != nil
}

// description is TBD
// SetFixed sets the FlowPayloadFixed value in the FlowPayload object
func (obj *flowPayload) SetFixed(value FlowPayloadFixed) FlowPayload {
	obj.setChoice(FlowPayloadChoice.FIXED)
	obj.fixedHolder = nil
	obj.obj.Fixed = value.msg()

	return obj
}

func (obj *flowPayload) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Fixed != nil {

		obj.Fixed().validateObj(vObj, set_default)
	}

}

func (obj *flowPayload) setDefault() {
	var choices_set int = 0
	var choice FlowPayloadChoiceEnum

	if obj.obj.Fixed != nil {
		choices_set += 1
		choice = FlowPayloadChoice.FIXED
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowPayloadChoice.FIXED)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowPayload")
			}
		} else {
			intVal := otg.FlowPayload_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowPayload_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

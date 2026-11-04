package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv4OptionsTimestampPointer *****
type flowIpv4OptionsTimestampPointer struct {
	validation
	obj          *otg.FlowIpv4OptionsTimestampPointer
	marshaller   marshalFlowIpv4OptionsTimestampPointer
	unMarshaller unMarshalFlowIpv4OptionsTimestampPointer
}

func NewFlowIpv4OptionsTimestampPointer() FlowIpv4OptionsTimestampPointer {
	obj := flowIpv4OptionsTimestampPointer{obj: &otg.FlowIpv4OptionsTimestampPointer{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv4OptionsTimestampPointer) msg() *otg.FlowIpv4OptionsTimestampPointer {
	return obj.obj
}

func (obj *flowIpv4OptionsTimestampPointer) setMsg(msg *otg.FlowIpv4OptionsTimestampPointer) FlowIpv4OptionsTimestampPointer {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv4OptionsTimestampPointer struct {
	obj *flowIpv4OptionsTimestampPointer
}

type marshalFlowIpv4OptionsTimestampPointer interface {
	// ToProto marshals FlowIpv4OptionsTimestampPointer to protobuf object *otg.FlowIpv4OptionsTimestampPointer
	ToProto() (*otg.FlowIpv4OptionsTimestampPointer, error)
	// ToPbText marshals FlowIpv4OptionsTimestampPointer to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv4OptionsTimestampPointer to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv4OptionsTimestampPointer to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv4OptionsTimestampPointer struct {
	obj *flowIpv4OptionsTimestampPointer
}

type unMarshalFlowIpv4OptionsTimestampPointer interface {
	// FromProto unmarshals FlowIpv4OptionsTimestampPointer from protobuf object *otg.FlowIpv4OptionsTimestampPointer
	FromProto(msg *otg.FlowIpv4OptionsTimestampPointer) (FlowIpv4OptionsTimestampPointer, error)
	// FromPbText unmarshals FlowIpv4OptionsTimestampPointer from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv4OptionsTimestampPointer from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv4OptionsTimestampPointer from JSON text
	FromJson(value string) error
}

func (obj *flowIpv4OptionsTimestampPointer) Marshal() marshalFlowIpv4OptionsTimestampPointer {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv4OptionsTimestampPointer{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv4OptionsTimestampPointer) Unmarshal() unMarshalFlowIpv4OptionsTimestampPointer {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv4OptionsTimestampPointer{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv4OptionsTimestampPointer) ToProto() (*otg.FlowIpv4OptionsTimestampPointer, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv4OptionsTimestampPointer) FromProto(msg *otg.FlowIpv4OptionsTimestampPointer) (FlowIpv4OptionsTimestampPointer, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv4OptionsTimestampPointer) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv4OptionsTimestampPointer) FromPbText(value string) error {
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

func (m *marshalflowIpv4OptionsTimestampPointer) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv4OptionsTimestampPointer) FromYaml(value string) error {
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

func (m *marshalflowIpv4OptionsTimestampPointer) ToJson() (string, error) {
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

func (m *unMarshalflowIpv4OptionsTimestampPointer) FromJson(value string) error {
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

func (obj *flowIpv4OptionsTimestampPointer) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv4OptionsTimestampPointer) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv4OptionsTimestampPointer) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv4OptionsTimestampPointer) Clone() (FlowIpv4OptionsTimestampPointer, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv4OptionsTimestampPointer()
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

// FlowIpv4OptionsTimestampPointer is the attribute pointer indicates the octet offset where the next router should begin recording its data.
// Choices of input are,
// 1. auto : The OTG implementation can provide a system generated value for this property. If the implementation is unable to generate a value, the default value must be used.
// 2. value: User can configure the pointer value.
type FlowIpv4OptionsTimestampPointer interface {
	Validation
	// msg marshals FlowIpv4OptionsTimestampPointer to protobuf object *otg.FlowIpv4OptionsTimestampPointer
	// and doesn't set defaults
	msg() *otg.FlowIpv4OptionsTimestampPointer
	// setMsg unmarshals FlowIpv4OptionsTimestampPointer from protobuf object *otg.FlowIpv4OptionsTimestampPointer
	// and doesn't set defaults
	setMsg(*otg.FlowIpv4OptionsTimestampPointer) FlowIpv4OptionsTimestampPointer
	// provides marshal interface
	Marshal() marshalFlowIpv4OptionsTimestampPointer
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv4OptionsTimestampPointer
	// validate validates FlowIpv4OptionsTimestampPointer
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv4OptionsTimestampPointer, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowIpv4OptionsTimestampPointerChoiceEnum, set in FlowIpv4OptionsTimestampPointer
	Choice() FlowIpv4OptionsTimestampPointerChoiceEnum
	// setChoice assigns FlowIpv4OptionsTimestampPointerChoiceEnum provided by user to FlowIpv4OptionsTimestampPointer
	setChoice(value FlowIpv4OptionsTimestampPointerChoiceEnum) FlowIpv4OptionsTimestampPointer
	// HasChoice checks if Choice has been set in FlowIpv4OptionsTimestampPointer
	HasChoice() bool
	// getter for Auto to set choice.
	Auto()
	// Value returns uint32, set in FlowIpv4OptionsTimestampPointer.
	Value() uint32
	// SetValue assigns uint32 provided by user to FlowIpv4OptionsTimestampPointer
	SetValue(value uint32) FlowIpv4OptionsTimestampPointer
	// HasValue checks if Value has been set in FlowIpv4OptionsTimestampPointer
	HasValue() bool
}

type FlowIpv4OptionsTimestampPointerChoiceEnum string

// Enum of Choice on FlowIpv4OptionsTimestampPointer
var FlowIpv4OptionsTimestampPointerChoice = struct {
	AUTO  FlowIpv4OptionsTimestampPointerChoiceEnum
	VALUE FlowIpv4OptionsTimestampPointerChoiceEnum
}{
	AUTO:  FlowIpv4OptionsTimestampPointerChoiceEnum("auto"),
	VALUE: FlowIpv4OptionsTimestampPointerChoiceEnum("value"),
}

func (obj *flowIpv4OptionsTimestampPointer) Choice() FlowIpv4OptionsTimestampPointerChoiceEnum {
	return FlowIpv4OptionsTimestampPointerChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Auto to set choice
func (obj *flowIpv4OptionsTimestampPointer) Auto() {
	obj.setChoice(FlowIpv4OptionsTimestampPointerChoice.AUTO)
}

// description is TBD
// Choice returns a string
func (obj *flowIpv4OptionsTimestampPointer) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowIpv4OptionsTimestampPointer) setChoice(value FlowIpv4OptionsTimestampPointerChoiceEnum) FlowIpv4OptionsTimestampPointer {
	intValue, ok := otg.FlowIpv4OptionsTimestampPointer_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowIpv4OptionsTimestampPointerChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowIpv4OptionsTimestampPointer_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Value = nil

	if value == FlowIpv4OptionsTimestampPointerChoice.VALUE {
		defaultValue := uint32(5)
		obj.obj.Value = &defaultValue
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *flowIpv4OptionsTimestampPointer) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(FlowIpv4OptionsTimestampPointerChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *flowIpv4OptionsTimestampPointer) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the FlowIpv4OptionsTimestampPointer object
func (obj *flowIpv4OptionsTimestampPointer) SetValue(value uint32) FlowIpv4OptionsTimestampPointer {
	obj.setChoice(FlowIpv4OptionsTimestampPointerChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

func (obj *flowIpv4OptionsTimestampPointer) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value < 5 || *obj.obj.Value > 40 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("5 <= FlowIpv4OptionsTimestampPointer.Value <= 40 but Got %d", *obj.obj.Value))
		}

	}

}

func (obj *flowIpv4OptionsTimestampPointer) setDefault() {
	var choices_set int = 0
	var choice FlowIpv4OptionsTimestampPointerChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = FlowIpv4OptionsTimestampPointerChoice.VALUE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowIpv4OptionsTimestampPointerChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowIpv4OptionsTimestampPointer")
			}
		} else {
			intVal := otg.FlowIpv4OptionsTimestampPointer_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowIpv4OptionsTimestampPointer_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

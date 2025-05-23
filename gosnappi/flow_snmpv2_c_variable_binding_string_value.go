package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowSnmpv2CVariableBindingStringValue *****
type flowSnmpv2CVariableBindingStringValue struct {
	validation
	obj          *otg.FlowSnmpv2CVariableBindingStringValue
	marshaller   marshalFlowSnmpv2CVariableBindingStringValue
	unMarshaller unMarshalFlowSnmpv2CVariableBindingStringValue
}

func NewFlowSnmpv2CVariableBindingStringValue() FlowSnmpv2CVariableBindingStringValue {
	obj := flowSnmpv2CVariableBindingStringValue{obj: &otg.FlowSnmpv2CVariableBindingStringValue{}}
	obj.setDefault()
	return &obj
}

func (obj *flowSnmpv2CVariableBindingStringValue) msg() *otg.FlowSnmpv2CVariableBindingStringValue {
	return obj.obj
}

func (obj *flowSnmpv2CVariableBindingStringValue) setMsg(msg *otg.FlowSnmpv2CVariableBindingStringValue) FlowSnmpv2CVariableBindingStringValue {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowSnmpv2CVariableBindingStringValue struct {
	obj *flowSnmpv2CVariableBindingStringValue
}

type marshalFlowSnmpv2CVariableBindingStringValue interface {
	// ToProto marshals FlowSnmpv2CVariableBindingStringValue to protobuf object *otg.FlowSnmpv2CVariableBindingStringValue
	ToProto() (*otg.FlowSnmpv2CVariableBindingStringValue, error)
	// ToPbText marshals FlowSnmpv2CVariableBindingStringValue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowSnmpv2CVariableBindingStringValue to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowSnmpv2CVariableBindingStringValue to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowSnmpv2CVariableBindingStringValue to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowSnmpv2CVariableBindingStringValue struct {
	obj *flowSnmpv2CVariableBindingStringValue
}

type unMarshalFlowSnmpv2CVariableBindingStringValue interface {
	// FromProto unmarshals FlowSnmpv2CVariableBindingStringValue from protobuf object *otg.FlowSnmpv2CVariableBindingStringValue
	FromProto(msg *otg.FlowSnmpv2CVariableBindingStringValue) (FlowSnmpv2CVariableBindingStringValue, error)
	// FromPbText unmarshals FlowSnmpv2CVariableBindingStringValue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowSnmpv2CVariableBindingStringValue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowSnmpv2CVariableBindingStringValue from JSON text
	FromJson(value string) error
}

func (obj *flowSnmpv2CVariableBindingStringValue) Marshal() marshalFlowSnmpv2CVariableBindingStringValue {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowSnmpv2CVariableBindingStringValue{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowSnmpv2CVariableBindingStringValue) Unmarshal() unMarshalFlowSnmpv2CVariableBindingStringValue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowSnmpv2CVariableBindingStringValue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowSnmpv2CVariableBindingStringValue) ToProto() (*otg.FlowSnmpv2CVariableBindingStringValue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowSnmpv2CVariableBindingStringValue) FromProto(msg *otg.FlowSnmpv2CVariableBindingStringValue) (FlowSnmpv2CVariableBindingStringValue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowSnmpv2CVariableBindingStringValue) ToPbText() (string, error) {
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

func (m *unMarshalflowSnmpv2CVariableBindingStringValue) FromPbText(value string) error {
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

func (m *marshalflowSnmpv2CVariableBindingStringValue) ToYaml() (string, error) {
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

func (m *unMarshalflowSnmpv2CVariableBindingStringValue) FromYaml(value string) error {
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

func (m *marshalflowSnmpv2CVariableBindingStringValue) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalflowSnmpv2CVariableBindingStringValue) ToJson() (string, error) {
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

func (m *unMarshalflowSnmpv2CVariableBindingStringValue) FromJson(value string) error {
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

func (obj *flowSnmpv2CVariableBindingStringValue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowSnmpv2CVariableBindingStringValue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowSnmpv2CVariableBindingStringValue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowSnmpv2CVariableBindingStringValue) Clone() (FlowSnmpv2CVariableBindingStringValue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowSnmpv2CVariableBindingStringValue()
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

// FlowSnmpv2CVariableBindingStringValue is it contains the raw/ascii string value to be sent.
type FlowSnmpv2CVariableBindingStringValue interface {
	Validation
	// msg marshals FlowSnmpv2CVariableBindingStringValue to protobuf object *otg.FlowSnmpv2CVariableBindingStringValue
	// and doesn't set defaults
	msg() *otg.FlowSnmpv2CVariableBindingStringValue
	// setMsg unmarshals FlowSnmpv2CVariableBindingStringValue from protobuf object *otg.FlowSnmpv2CVariableBindingStringValue
	// and doesn't set defaults
	setMsg(*otg.FlowSnmpv2CVariableBindingStringValue) FlowSnmpv2CVariableBindingStringValue
	// provides marshal interface
	Marshal() marshalFlowSnmpv2CVariableBindingStringValue
	// provides unmarshal interface
	Unmarshal() unMarshalFlowSnmpv2CVariableBindingStringValue
	// validate validates FlowSnmpv2CVariableBindingStringValue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowSnmpv2CVariableBindingStringValue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowSnmpv2CVariableBindingStringValueChoiceEnum, set in FlowSnmpv2CVariableBindingStringValue
	Choice() FlowSnmpv2CVariableBindingStringValueChoiceEnum
	// setChoice assigns FlowSnmpv2CVariableBindingStringValueChoiceEnum provided by user to FlowSnmpv2CVariableBindingStringValue
	setChoice(value FlowSnmpv2CVariableBindingStringValueChoiceEnum) FlowSnmpv2CVariableBindingStringValue
	// HasChoice checks if Choice has been set in FlowSnmpv2CVariableBindingStringValue
	HasChoice() bool
	// Ascii returns string, set in FlowSnmpv2CVariableBindingStringValue.
	Ascii() string
	// SetAscii assigns string provided by user to FlowSnmpv2CVariableBindingStringValue
	SetAscii(value string) FlowSnmpv2CVariableBindingStringValue
	// HasAscii checks if Ascii has been set in FlowSnmpv2CVariableBindingStringValue
	HasAscii() bool
	// Raw returns string, set in FlowSnmpv2CVariableBindingStringValue.
	Raw() string
	// SetRaw assigns string provided by user to FlowSnmpv2CVariableBindingStringValue
	SetRaw(value string) FlowSnmpv2CVariableBindingStringValue
	// HasRaw checks if Raw has been set in FlowSnmpv2CVariableBindingStringValue
	HasRaw() bool
}

type FlowSnmpv2CVariableBindingStringValueChoiceEnum string

// Enum of Choice on FlowSnmpv2CVariableBindingStringValue
var FlowSnmpv2CVariableBindingStringValueChoice = struct {
	ASCII FlowSnmpv2CVariableBindingStringValueChoiceEnum
	RAW   FlowSnmpv2CVariableBindingStringValueChoiceEnum
}{
	ASCII: FlowSnmpv2CVariableBindingStringValueChoiceEnum("ascii"),
	RAW:   FlowSnmpv2CVariableBindingStringValueChoiceEnum("raw"),
}

func (obj *flowSnmpv2CVariableBindingStringValue) Choice() FlowSnmpv2CVariableBindingStringValueChoiceEnum {
	return FlowSnmpv2CVariableBindingStringValueChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowSnmpv2CVariableBindingStringValue) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowSnmpv2CVariableBindingStringValue) setChoice(value FlowSnmpv2CVariableBindingStringValueChoiceEnum) FlowSnmpv2CVariableBindingStringValue {
	intValue, ok := otg.FlowSnmpv2CVariableBindingStringValue_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowSnmpv2CVariableBindingStringValueChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowSnmpv2CVariableBindingStringValue_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Raw = nil
	obj.obj.Ascii = nil

	if value == FlowSnmpv2CVariableBindingStringValueChoice.ASCII {
		defaultValue := "ascii"
		obj.obj.Ascii = &defaultValue
	}

	if value == FlowSnmpv2CVariableBindingStringValueChoice.RAW {
		defaultValue := "00"
		obj.obj.Raw = &defaultValue
	}

	return obj
}

// It contains the ASCII string to be sent.  As of now it is restricted to 10000 bytes.
// Ascii returns a string
func (obj *flowSnmpv2CVariableBindingStringValue) Ascii() string {

	if obj.obj.Ascii == nil {
		obj.setChoice(FlowSnmpv2CVariableBindingStringValueChoice.ASCII)
	}

	return *obj.obj.Ascii

}

// It contains the ASCII string to be sent.  As of now it is restricted to 10000 bytes.
// Ascii returns a string
func (obj *flowSnmpv2CVariableBindingStringValue) HasAscii() bool {
	return obj.obj.Ascii != nil
}

// It contains the ASCII string to be sent.  As of now it is restricted to 10000 bytes.
// SetAscii sets the string value in the FlowSnmpv2CVariableBindingStringValue object
func (obj *flowSnmpv2CVariableBindingStringValue) SetAscii(value string) FlowSnmpv2CVariableBindingStringValue {
	obj.setChoice(FlowSnmpv2CVariableBindingStringValueChoice.ASCII)
	obj.obj.Ascii = &value
	return obj
}

// It contains the hex string to be sent.  As of now it is restricted to 10000 bytes.
// Raw returns a string
func (obj *flowSnmpv2CVariableBindingStringValue) Raw() string {

	if obj.obj.Raw == nil {
		obj.setChoice(FlowSnmpv2CVariableBindingStringValueChoice.RAW)
	}

	return *obj.obj.Raw

}

// It contains the hex string to be sent.  As of now it is restricted to 10000 bytes.
// Raw returns a string
func (obj *flowSnmpv2CVariableBindingStringValue) HasRaw() bool {
	return obj.obj.Raw != nil
}

// It contains the hex string to be sent.  As of now it is restricted to 10000 bytes.
// SetRaw sets the string value in the FlowSnmpv2CVariableBindingStringValue object
func (obj *flowSnmpv2CVariableBindingStringValue) SetRaw(value string) FlowSnmpv2CVariableBindingStringValue {
	obj.setChoice(FlowSnmpv2CVariableBindingStringValueChoice.RAW)
	obj.obj.Raw = &value
	return obj
}

func (obj *flowSnmpv2CVariableBindingStringValue) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ascii != nil {

		if len(*obj.obj.Ascii) > 10000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"None <= length of FlowSnmpv2CVariableBindingStringValue.Ascii <= 10000 but Got %d",
					len(*obj.obj.Ascii)))
		}

	}

	if obj.obj.Raw != nil {

		if len(*obj.obj.Raw) > 10000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"None <= length of FlowSnmpv2CVariableBindingStringValue.Raw <= 10000 but Got %d",
					len(*obj.obj.Raw)))
		}

	}

}

func (obj *flowSnmpv2CVariableBindingStringValue) setDefault() {
	var choices_set int = 0
	var choice FlowSnmpv2CVariableBindingStringValueChoiceEnum

	if obj.obj.Ascii != nil {
		choices_set += 1
		choice = FlowSnmpv2CVariableBindingStringValueChoice.ASCII
	}

	if obj.obj.Raw != nil {
		choices_set += 1
		choice = FlowSnmpv2CVariableBindingStringValueChoice.RAW
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowSnmpv2CVariableBindingStringValueChoice.ASCII)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowSnmpv2CVariableBindingStringValue")
			}
		} else {
			intVal := otg.FlowSnmpv2CVariableBindingStringValue_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowSnmpv2CVariableBindingStringValue_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

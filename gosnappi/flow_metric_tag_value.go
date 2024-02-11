package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowMetricTagValue *****
type flowMetricTagValue struct {
	validation
	obj          *otg.FlowMetricTagValue
	marshaller   marshalFlowMetricTagValue
	unMarshaller unMarshalFlowMetricTagValue
}

func NewFlowMetricTagValue() FlowMetricTagValue {
	obj := flowMetricTagValue{obj: &otg.FlowMetricTagValue{}}
	obj.setDefault()
	return &obj
}

func (obj *flowMetricTagValue) msg() *otg.FlowMetricTagValue {
	return obj.obj
}

func (obj *flowMetricTagValue) setMsg(msg *otg.FlowMetricTagValue) FlowMetricTagValue {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowMetricTagValue struct {
	obj *flowMetricTagValue
}

type marshalFlowMetricTagValue interface {
	// ToProto marshals FlowMetricTagValue to protobuf object *otg.FlowMetricTagValue
	ToProto() (*otg.FlowMetricTagValue, error)
	// ToPbText marshals FlowMetricTagValue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowMetricTagValue to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowMetricTagValue to JSON text
	ToJson() (string, error)
}

type unMarshalflowMetricTagValue struct {
	obj *flowMetricTagValue
}

type unMarshalFlowMetricTagValue interface {
	// FromProto unmarshals FlowMetricTagValue from protobuf object *otg.FlowMetricTagValue
	FromProto(msg *otg.FlowMetricTagValue) (FlowMetricTagValue, error)
	// FromPbText unmarshals FlowMetricTagValue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowMetricTagValue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowMetricTagValue from JSON text
	FromJson(value string) error
}

func (obj *flowMetricTagValue) Marshal() marshalFlowMetricTagValue {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowMetricTagValue{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowMetricTagValue) Unmarshal() unMarshalFlowMetricTagValue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowMetricTagValue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowMetricTagValue) ToProto() (*otg.FlowMetricTagValue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowMetricTagValue) FromProto(msg *otg.FlowMetricTagValue) (FlowMetricTagValue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowMetricTagValue) ToPbText() (string, error) {
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

func (m *unMarshalflowMetricTagValue) FromPbText(value string) error {
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

func (m *marshalflowMetricTagValue) ToYaml() (string, error) {
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

func (m *unMarshalflowMetricTagValue) FromYaml(value string) error {
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

func (m *marshalflowMetricTagValue) ToJson() (string, error) {
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

func (m *unMarshalflowMetricTagValue) FromJson(value string) error {
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

func (obj *flowMetricTagValue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowMetricTagValue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowMetricTagValue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowMetricTagValue) Clone() (FlowMetricTagValue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowMetricTagValue()
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

// FlowMetricTagValue is a container for metric tag value
type FlowMetricTagValue interface {
	Validation
	// msg marshals FlowMetricTagValue to protobuf object *otg.FlowMetricTagValue
	// and doesn't set defaults
	msg() *otg.FlowMetricTagValue
	// setMsg unmarshals FlowMetricTagValue from protobuf object *otg.FlowMetricTagValue
	// and doesn't set defaults
	setMsg(*otg.FlowMetricTagValue) FlowMetricTagValue
	// provides marshal interface
	Marshal() marshalFlowMetricTagValue
	// provides unmarshal interface
	Unmarshal() unMarshalFlowMetricTagValue
	// validate validates FlowMetricTagValue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowMetricTagValue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowMetricTagValueChoiceEnum, set in FlowMetricTagValue
	Choice() FlowMetricTagValueChoiceEnum
	// setChoice assigns FlowMetricTagValueChoiceEnum provided by user to FlowMetricTagValue
	setChoice(value FlowMetricTagValueChoiceEnum) FlowMetricTagValue
	// HasChoice checks if Choice has been set in FlowMetricTagValue
	HasChoice() bool
	// Hex returns string, set in FlowMetricTagValue.
	Hex() string
	// SetHex assigns string provided by user to FlowMetricTagValue
	SetHex(value string) FlowMetricTagValue
	// HasHex checks if Hex has been set in FlowMetricTagValue
	HasHex() bool
	// Str returns string, set in FlowMetricTagValue.
	Str() string
	// SetStr assigns string provided by user to FlowMetricTagValue
	SetStr(value string) FlowMetricTagValue
	// HasStr checks if Str has been set in FlowMetricTagValue
	HasStr() bool
}

type FlowMetricTagValueChoiceEnum string

// Enum of Choice on FlowMetricTagValue
var FlowMetricTagValueChoice = struct {
	HEX FlowMetricTagValueChoiceEnum
	STR FlowMetricTagValueChoiceEnum
}{
	HEX: FlowMetricTagValueChoiceEnum("hex"),
	STR: FlowMetricTagValueChoiceEnum("str"),
}

func (obj *flowMetricTagValue) Choice() FlowMetricTagValueChoiceEnum {
	return FlowMetricTagValueChoiceEnum(obj.obj.Choice.Enum().String())
}

// Available formats for metric tag value
// Choice returns a string
func (obj *flowMetricTagValue) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowMetricTagValue) setChoice(value FlowMetricTagValueChoiceEnum) FlowMetricTagValue {
	intValue, ok := otg.FlowMetricTagValue_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowMetricTagValueChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowMetricTagValue_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Str = nil
	obj.obj.Hex = nil
	return obj
}

// Value represented in hexadecimal format
// Hex returns a string
func (obj *flowMetricTagValue) Hex() string {

	if obj.obj.Hex == nil {
		obj.setChoice(FlowMetricTagValueChoice.HEX)
	}

	return *obj.obj.Hex

}

// Value represented in hexadecimal format
// Hex returns a string
func (obj *flowMetricTagValue) HasHex() bool {
	return obj.obj.Hex != nil
}

// Value represented in hexadecimal format
// SetHex sets the string value in the FlowMetricTagValue object
func (obj *flowMetricTagValue) SetHex(value string) FlowMetricTagValue {
	obj.setChoice(FlowMetricTagValueChoice.HEX)
	obj.obj.Hex = &value
	return obj
}

// Value represented in string format
// Str returns a string
func (obj *flowMetricTagValue) Str() string {

	if obj.obj.Str == nil {
		obj.setChoice(FlowMetricTagValueChoice.STR)
	}

	return *obj.obj.Str

}

// Value represented in string format
// Str returns a string
func (obj *flowMetricTagValue) HasStr() bool {
	return obj.obj.Str != nil
}

// Value represented in string format
// SetStr sets the string value in the FlowMetricTagValue object
func (obj *flowMetricTagValue) SetStr(value string) FlowMetricTagValue {
	obj.setChoice(FlowMetricTagValueChoice.STR)
	obj.obj.Str = &value
	return obj
}

func (obj *flowMetricTagValue) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Hex != nil {

		err := obj.validateHex(obj.Hex())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowMetricTagValue.Hex"))
		}

	}

}

func (obj *flowMetricTagValue) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(FlowMetricTagValueChoice.HEX)

	}

}

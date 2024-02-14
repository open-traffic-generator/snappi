package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPExplicitRouteLength *****
type flowRSVPExplicitRouteLength struct {
	validation
	obj          *otg.FlowRSVPExplicitRouteLength
	marshaller   marshalFlowRSVPExplicitRouteLength
	unMarshaller unMarshalFlowRSVPExplicitRouteLength
}

func NewFlowRSVPExplicitRouteLength() FlowRSVPExplicitRouteLength {
	obj := flowRSVPExplicitRouteLength{obj: &otg.FlowRSVPExplicitRouteLength{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPExplicitRouteLength) msg() *otg.FlowRSVPExplicitRouteLength {
	return obj.obj
}

func (obj *flowRSVPExplicitRouteLength) setMsg(msg *otg.FlowRSVPExplicitRouteLength) FlowRSVPExplicitRouteLength {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPExplicitRouteLength struct {
	obj *flowRSVPExplicitRouteLength
}

type marshalFlowRSVPExplicitRouteLength interface {
	// ToProto marshals FlowRSVPExplicitRouteLength to protobuf object *otg.FlowRSVPExplicitRouteLength
	ToProto() (*otg.FlowRSVPExplicitRouteLength, error)
	// ToPbText marshals FlowRSVPExplicitRouteLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPExplicitRouteLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPExplicitRouteLength to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPExplicitRouteLength struct {
	obj *flowRSVPExplicitRouteLength
}

type unMarshalFlowRSVPExplicitRouteLength interface {
	// FromProto unmarshals FlowRSVPExplicitRouteLength from protobuf object *otg.FlowRSVPExplicitRouteLength
	FromProto(msg *otg.FlowRSVPExplicitRouteLength) (FlowRSVPExplicitRouteLength, error)
	// FromPbText unmarshals FlowRSVPExplicitRouteLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPExplicitRouteLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPExplicitRouteLength from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPExplicitRouteLength) Marshal() marshalFlowRSVPExplicitRouteLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPExplicitRouteLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPExplicitRouteLength) Unmarshal() unMarshalFlowRSVPExplicitRouteLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPExplicitRouteLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPExplicitRouteLength) ToProto() (*otg.FlowRSVPExplicitRouteLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPExplicitRouteLength) FromProto(msg *otg.FlowRSVPExplicitRouteLength) (FlowRSVPExplicitRouteLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPExplicitRouteLength) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPExplicitRouteLength) FromPbText(value string) error {
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

func (m *marshalflowRSVPExplicitRouteLength) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPExplicitRouteLength) FromYaml(value string) error {
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

func (m *marshalflowRSVPExplicitRouteLength) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPExplicitRouteLength) FromJson(value string) error {
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

func (obj *flowRSVPExplicitRouteLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPExplicitRouteLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPExplicitRouteLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPExplicitRouteLength) Clone() (FlowRSVPExplicitRouteLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPExplicitRouteLength()
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

// FlowRSVPExplicitRouteLength is description is TBD
type FlowRSVPExplicitRouteLength interface {
	Validation
	// msg marshals FlowRSVPExplicitRouteLength to protobuf object *otg.FlowRSVPExplicitRouteLength
	// and doesn't set defaults
	msg() *otg.FlowRSVPExplicitRouteLength
	// setMsg unmarshals FlowRSVPExplicitRouteLength from protobuf object *otg.FlowRSVPExplicitRouteLength
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPExplicitRouteLength) FlowRSVPExplicitRouteLength
	// provides marshal interface
	Marshal() marshalFlowRSVPExplicitRouteLength
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPExplicitRouteLength
	// validate validates FlowRSVPExplicitRouteLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPExplicitRouteLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPExplicitRouteLengthChoiceEnum, set in FlowRSVPExplicitRouteLength
	Choice() FlowRSVPExplicitRouteLengthChoiceEnum
	// setChoice assigns FlowRSVPExplicitRouteLengthChoiceEnum provided by user to FlowRSVPExplicitRouteLength
	setChoice(value FlowRSVPExplicitRouteLengthChoiceEnum) FlowRSVPExplicitRouteLength
	// HasChoice checks if Choice has been set in FlowRSVPExplicitRouteLength
	HasChoice() bool
	// Auto returns uint32, set in FlowRSVPExplicitRouteLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in FlowRSVPExplicitRouteLength
	HasAuto() bool
	// Value returns uint32, set in FlowRSVPExplicitRouteLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to FlowRSVPExplicitRouteLength
	SetValue(value uint32) FlowRSVPExplicitRouteLength
	// HasValue checks if Value has been set in FlowRSVPExplicitRouteLength
	HasValue() bool
}

type FlowRSVPExplicitRouteLengthChoiceEnum string

// Enum of Choice on FlowRSVPExplicitRouteLength
var FlowRSVPExplicitRouteLengthChoice = struct {
	AUTO  FlowRSVPExplicitRouteLengthChoiceEnum
	VALUE FlowRSVPExplicitRouteLengthChoiceEnum
}{
	AUTO:  FlowRSVPExplicitRouteLengthChoiceEnum("auto"),
	VALUE: FlowRSVPExplicitRouteLengthChoiceEnum("value"),
}

func (obj *flowRSVPExplicitRouteLength) Choice() FlowRSVPExplicitRouteLengthChoiceEnum {
	return FlowRSVPExplicitRouteLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// auto or configured value.
// Choice returns a string
func (obj *flowRSVPExplicitRouteLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPExplicitRouteLength) setChoice(value FlowRSVPExplicitRouteLengthChoiceEnum) FlowRSVPExplicitRouteLength {
	intValue, ok := otg.FlowRSVPExplicitRouteLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPExplicitRouteLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPExplicitRouteLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Value = nil
	obj.obj.Auto = nil

	if value == FlowRSVPExplicitRouteLengthChoice.AUTO {
		defaultValue := uint32(8)
		obj.obj.Auto = &defaultValue
	}

	if value == FlowRSVPExplicitRouteLengthChoice.VALUE {
		defaultValue := uint32(8)
		obj.obj.Value = &defaultValue
	}

	return obj
}

// OTG will provide a system generated value for this property.  If OTG is unable to generate a value the default value must be used.
// Auto returns a uint32
func (obj *flowRSVPExplicitRouteLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(FlowRSVPExplicitRouteLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// OTG will provide a system generated value for this property.  If OTG is unable to generate a value the default value must be used.
// Auto returns a uint32
func (obj *flowRSVPExplicitRouteLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Value returns a uint32
func (obj *flowRSVPExplicitRouteLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(FlowRSVPExplicitRouteLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *flowRSVPExplicitRouteLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the FlowRSVPExplicitRouteLength object
func (obj *flowRSVPExplicitRouteLength) SetValue(value uint32) FlowRSVPExplicitRouteLength {
	obj.setChoice(FlowRSVPExplicitRouteLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

func (obj *flowRSVPExplicitRouteLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowRSVPExplicitRouteLength.Auto <= 256 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowRSVPExplicitRouteLength.Value <= 256 but Got %d", *obj.obj.Value))
		}

	}

}

func (obj *flowRSVPExplicitRouteLength) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(FlowRSVPExplicitRouteLengthChoice.AUTO)

	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathRecordRouteLabel *****
type flowRSVPPathRecordRouteLabel struct {
	validation
	obj          *otg.FlowRSVPPathRecordRouteLabel
	marshaller   marshalFlowRSVPPathRecordRouteLabel
	unMarshaller unMarshalFlowRSVPPathRecordRouteLabel
}

func NewFlowRSVPPathRecordRouteLabel() FlowRSVPPathRecordRouteLabel {
	obj := flowRSVPPathRecordRouteLabel{obj: &otg.FlowRSVPPathRecordRouteLabel{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathRecordRouteLabel) msg() *otg.FlowRSVPPathRecordRouteLabel {
	return obj.obj
}

func (obj *flowRSVPPathRecordRouteLabel) setMsg(msg *otg.FlowRSVPPathRecordRouteLabel) FlowRSVPPathRecordRouteLabel {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathRecordRouteLabel struct {
	obj *flowRSVPPathRecordRouteLabel
}

type marshalFlowRSVPPathRecordRouteLabel interface {
	// ToProto marshals FlowRSVPPathRecordRouteLabel to protobuf object *otg.FlowRSVPPathRecordRouteLabel
	ToProto() (*otg.FlowRSVPPathRecordRouteLabel, error)
	// ToPbText marshals FlowRSVPPathRecordRouteLabel to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathRecordRouteLabel to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathRecordRouteLabel to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowRSVPPathRecordRouteLabel to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowRSVPPathRecordRouteLabel struct {
	obj *flowRSVPPathRecordRouteLabel
}

type unMarshalFlowRSVPPathRecordRouteLabel interface {
	// FromProto unmarshals FlowRSVPPathRecordRouteLabel from protobuf object *otg.FlowRSVPPathRecordRouteLabel
	FromProto(msg *otg.FlowRSVPPathRecordRouteLabel) (FlowRSVPPathRecordRouteLabel, error)
	// FromPbText unmarshals FlowRSVPPathRecordRouteLabel from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathRecordRouteLabel from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathRecordRouteLabel from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathRecordRouteLabel) Marshal() marshalFlowRSVPPathRecordRouteLabel {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathRecordRouteLabel{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathRecordRouteLabel) Unmarshal() unMarshalFlowRSVPPathRecordRouteLabel {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathRecordRouteLabel{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathRecordRouteLabel) ToProto() (*otg.FlowRSVPPathRecordRouteLabel, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathRecordRouteLabel) FromProto(msg *otg.FlowRSVPPathRecordRouteLabel) (FlowRSVPPathRecordRouteLabel, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathRecordRouteLabel) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathRecordRouteLabel) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathRecordRouteLabel) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathRecordRouteLabel) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathRecordRouteLabel) ToJsonRaw() (string, error) {
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

func (m *marshalflowRSVPPathRecordRouteLabel) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathRecordRouteLabel) FromJson(value string) error {
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

func (obj *flowRSVPPathRecordRouteLabel) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathRecordRouteLabel) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathRecordRouteLabel) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathRecordRouteLabel) Clone() (FlowRSVPPathRecordRouteLabel, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathRecordRouteLabel()
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

// FlowRSVPPathRecordRouteLabel is description is TBD
type FlowRSVPPathRecordRouteLabel interface {
	Validation
	// msg marshals FlowRSVPPathRecordRouteLabel to protobuf object *otg.FlowRSVPPathRecordRouteLabel
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathRecordRouteLabel
	// setMsg unmarshals FlowRSVPPathRecordRouteLabel from protobuf object *otg.FlowRSVPPathRecordRouteLabel
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathRecordRouteLabel) FlowRSVPPathRecordRouteLabel
	// provides marshal interface
	Marshal() marshalFlowRSVPPathRecordRouteLabel
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathRecordRouteLabel
	// validate validates FlowRSVPPathRecordRouteLabel
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathRecordRouteLabel, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPPathRecordRouteLabelChoiceEnum, set in FlowRSVPPathRecordRouteLabel
	Choice() FlowRSVPPathRecordRouteLabelChoiceEnum
	// setChoice assigns FlowRSVPPathRecordRouteLabelChoiceEnum provided by user to FlowRSVPPathRecordRouteLabel
	setChoice(value FlowRSVPPathRecordRouteLabelChoiceEnum) FlowRSVPPathRecordRouteLabel
	// HasChoice checks if Choice has been set in FlowRSVPPathRecordRouteLabel
	HasChoice() bool
	// AsInteger returns uint32, set in FlowRSVPPathRecordRouteLabel.
	AsInteger() uint32
	// SetAsInteger assigns uint32 provided by user to FlowRSVPPathRecordRouteLabel
	SetAsInteger(value uint32) FlowRSVPPathRecordRouteLabel
	// HasAsInteger checks if AsInteger has been set in FlowRSVPPathRecordRouteLabel
	HasAsInteger() bool
	// AsHex returns string, set in FlowRSVPPathRecordRouteLabel.
	AsHex() string
	// SetAsHex assigns string provided by user to FlowRSVPPathRecordRouteLabel
	SetAsHex(value string) FlowRSVPPathRecordRouteLabel
	// HasAsHex checks if AsHex has been set in FlowRSVPPathRecordRouteLabel
	HasAsHex() bool
}

type FlowRSVPPathRecordRouteLabelChoiceEnum string

// Enum of Choice on FlowRSVPPathRecordRouteLabel
var FlowRSVPPathRecordRouteLabelChoice = struct {
	AS_INTEGER FlowRSVPPathRecordRouteLabelChoiceEnum
	AS_HEX     FlowRSVPPathRecordRouteLabelChoiceEnum
}{
	AS_INTEGER: FlowRSVPPathRecordRouteLabelChoiceEnum("as_integer"),
	AS_HEX:     FlowRSVPPathRecordRouteLabelChoiceEnum("as_hex"),
}

func (obj *flowRSVPPathRecordRouteLabel) Choice() FlowRSVPPathRecordRouteLabelChoiceEnum {
	return FlowRSVPPathRecordRouteLabelChoiceEnum(obj.obj.Choice.Enum().String())
}

// 32 bit integer or hex value.
// Choice returns a string
func (obj *flowRSVPPathRecordRouteLabel) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPPathRecordRouteLabel) setChoice(value FlowRSVPPathRecordRouteLabelChoiceEnum) FlowRSVPPathRecordRouteLabel {
	intValue, ok := otg.FlowRSVPPathRecordRouteLabel_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPPathRecordRouteLabelChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPPathRecordRouteLabel_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.AsHex = nil
	obj.obj.AsInteger = nil

	if value == FlowRSVPPathRecordRouteLabelChoice.AS_INTEGER {
		defaultValue := uint32(16)
		obj.obj.AsInteger = &defaultValue
	}

	if value == FlowRSVPPathRecordRouteLabelChoice.AS_HEX {
		defaultValue := "10"
		obj.obj.AsHex = &defaultValue
	}

	return obj
}

// description is TBD
// AsInteger returns a uint32
func (obj *flowRSVPPathRecordRouteLabel) AsInteger() uint32 {

	if obj.obj.AsInteger == nil {
		obj.setChoice(FlowRSVPPathRecordRouteLabelChoice.AS_INTEGER)
	}

	return *obj.obj.AsInteger

}

// description is TBD
// AsInteger returns a uint32
func (obj *flowRSVPPathRecordRouteLabel) HasAsInteger() bool {
	return obj.obj.AsInteger != nil
}

// description is TBD
// SetAsInteger sets the uint32 value in the FlowRSVPPathRecordRouteLabel object
func (obj *flowRSVPPathRecordRouteLabel) SetAsInteger(value uint32) FlowRSVPPathRecordRouteLabel {
	obj.setChoice(FlowRSVPPathRecordRouteLabelChoice.AS_INTEGER)
	obj.obj.AsInteger = &value
	return obj
}

// Value of the this field should not excced 4 bytes. Maximum length of this attribute is 8 (4 * 2 hex character per byte).
// AsHex returns a string
func (obj *flowRSVPPathRecordRouteLabel) AsHex() string {

	if obj.obj.AsHex == nil {
		obj.setChoice(FlowRSVPPathRecordRouteLabelChoice.AS_HEX)
	}

	return *obj.obj.AsHex

}

// Value of the this field should not excced 4 bytes. Maximum length of this attribute is 8 (4 * 2 hex character per byte).
// AsHex returns a string
func (obj *flowRSVPPathRecordRouteLabel) HasAsHex() bool {
	return obj.obj.AsHex != nil
}

// Value of the this field should not excced 4 bytes. Maximum length of this attribute is 8 (4 * 2 hex character per byte).
// SetAsHex sets the string value in the FlowRSVPPathRecordRouteLabel object
func (obj *flowRSVPPathRecordRouteLabel) SetAsHex(value string) FlowRSVPPathRecordRouteLabel {
	obj.setChoice(FlowRSVPPathRecordRouteLabelChoice.AS_HEX)
	obj.obj.AsHex = &value
	return obj
}

func (obj *flowRSVPPathRecordRouteLabel) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.AsInteger != nil {

		if *obj.obj.AsInteger > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowRSVPPathRecordRouteLabel.AsInteger <= 1048575 but Got %d", *obj.obj.AsInteger))
		}

	}

	if obj.obj.AsHex != nil {

		if len(*obj.obj.AsHex) > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"None <= length of FlowRSVPPathRecordRouteLabel.AsHex <= 8 but Got %d",
					len(*obj.obj.AsHex)))
		}

	}

}

func (obj *flowRSVPPathRecordRouteLabel) setDefault() {
	var choices_set int = 0
	var choice FlowRSVPPathRecordRouteLabelChoiceEnum

	if obj.obj.AsInteger != nil {
		choices_set += 1
		choice = FlowRSVPPathRecordRouteLabelChoice.AS_INTEGER
	}

	if obj.obj.AsHex != nil {
		choices_set += 1
		choice = FlowRSVPPathRecordRouteLabelChoice.AS_HEX
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowRSVPPathRecordRouteLabelChoice.AS_INTEGER)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowRSVPPathRecordRouteLabel")
			}
		} else {
			intVal := otg.FlowRSVPPathRecordRouteLabel_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowRSVPPathRecordRouteLabel_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

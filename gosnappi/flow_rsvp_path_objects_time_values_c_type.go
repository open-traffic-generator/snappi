package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjectsTimeValuesCType *****
type flowRSVPPathObjectsTimeValuesCType struct {
	validation
	obj          *otg.FlowRSVPPathObjectsTimeValuesCType
	marshaller   marshalFlowRSVPPathObjectsTimeValuesCType
	unMarshaller unMarshalFlowRSVPPathObjectsTimeValuesCType
	type_1Holder FlowRSVPPathTimeValuesType1
}

func NewFlowRSVPPathObjectsTimeValuesCType() FlowRSVPPathObjectsTimeValuesCType {
	obj := flowRSVPPathObjectsTimeValuesCType{obj: &otg.FlowRSVPPathObjectsTimeValuesCType{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjectsTimeValuesCType) msg() *otg.FlowRSVPPathObjectsTimeValuesCType {
	return obj.obj
}

func (obj *flowRSVPPathObjectsTimeValuesCType) setMsg(msg *otg.FlowRSVPPathObjectsTimeValuesCType) FlowRSVPPathObjectsTimeValuesCType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjectsTimeValuesCType struct {
	obj *flowRSVPPathObjectsTimeValuesCType
}

type marshalFlowRSVPPathObjectsTimeValuesCType interface {
	// ToProto marshals FlowRSVPPathObjectsTimeValuesCType to protobuf object *otg.FlowRSVPPathObjectsTimeValuesCType
	ToProto() (*otg.FlowRSVPPathObjectsTimeValuesCType, error)
	// ToPbText marshals FlowRSVPPathObjectsTimeValuesCType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjectsTimeValuesCType to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjectsTimeValuesCType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowRSVPPathObjectsTimeValuesCType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowRSVPPathObjectsTimeValuesCType struct {
	obj *flowRSVPPathObjectsTimeValuesCType
}

type unMarshalFlowRSVPPathObjectsTimeValuesCType interface {
	// FromProto unmarshals FlowRSVPPathObjectsTimeValuesCType from protobuf object *otg.FlowRSVPPathObjectsTimeValuesCType
	FromProto(msg *otg.FlowRSVPPathObjectsTimeValuesCType) (FlowRSVPPathObjectsTimeValuesCType, error)
	// FromPbText unmarshals FlowRSVPPathObjectsTimeValuesCType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjectsTimeValuesCType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjectsTimeValuesCType from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjectsTimeValuesCType) Marshal() marshalFlowRSVPPathObjectsTimeValuesCType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjectsTimeValuesCType{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjectsTimeValuesCType) Unmarshal() unMarshalFlowRSVPPathObjectsTimeValuesCType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjectsTimeValuesCType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjectsTimeValuesCType) ToProto() (*otg.FlowRSVPPathObjectsTimeValuesCType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjectsTimeValuesCType) FromProto(msg *otg.FlowRSVPPathObjectsTimeValuesCType) (FlowRSVPPathObjectsTimeValuesCType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjectsTimeValuesCType) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsTimeValuesCType) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjectsTimeValuesCType) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsTimeValuesCType) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjectsTimeValuesCType) ToJsonRaw() (string, error) {
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

func (m *marshalflowRSVPPathObjectsTimeValuesCType) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsTimeValuesCType) FromJson(value string) error {
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

func (obj *flowRSVPPathObjectsTimeValuesCType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsTimeValuesCType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsTimeValuesCType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjectsTimeValuesCType) Clone() (FlowRSVPPathObjectsTimeValuesCType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjectsTimeValuesCType()
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

func (obj *flowRSVPPathObjectsTimeValuesCType) setNil() {
	obj.type_1Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjectsTimeValuesCType is object for TIME_VALUES class. Currently supported c-type is Type 1 Time Value (1).
type FlowRSVPPathObjectsTimeValuesCType interface {
	Validation
	// msg marshals FlowRSVPPathObjectsTimeValuesCType to protobuf object *otg.FlowRSVPPathObjectsTimeValuesCType
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjectsTimeValuesCType
	// setMsg unmarshals FlowRSVPPathObjectsTimeValuesCType from protobuf object *otg.FlowRSVPPathObjectsTimeValuesCType
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjectsTimeValuesCType) FlowRSVPPathObjectsTimeValuesCType
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjectsTimeValuesCType
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjectsTimeValuesCType
	// validate validates FlowRSVPPathObjectsTimeValuesCType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjectsTimeValuesCType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPPathObjectsTimeValuesCTypeChoiceEnum, set in FlowRSVPPathObjectsTimeValuesCType
	Choice() FlowRSVPPathObjectsTimeValuesCTypeChoiceEnum
	// setChoice assigns FlowRSVPPathObjectsTimeValuesCTypeChoiceEnum provided by user to FlowRSVPPathObjectsTimeValuesCType
	setChoice(value FlowRSVPPathObjectsTimeValuesCTypeChoiceEnum) FlowRSVPPathObjectsTimeValuesCType
	// HasChoice checks if Choice has been set in FlowRSVPPathObjectsTimeValuesCType
	HasChoice() bool
	// Type1 returns FlowRSVPPathTimeValuesType1, set in FlowRSVPPathObjectsTimeValuesCType.
	// FlowRSVPPathTimeValuesType1 is tIME_VALUES Object: Class = 5, C-Type = 1
	Type1() FlowRSVPPathTimeValuesType1
	// SetType1 assigns FlowRSVPPathTimeValuesType1 provided by user to FlowRSVPPathObjectsTimeValuesCType.
	// FlowRSVPPathTimeValuesType1 is tIME_VALUES Object: Class = 5, C-Type = 1
	SetType1(value FlowRSVPPathTimeValuesType1) FlowRSVPPathObjectsTimeValuesCType
	// HasType1 checks if Type1 has been set in FlowRSVPPathObjectsTimeValuesCType
	HasType1() bool
	setNil()
}

type FlowRSVPPathObjectsTimeValuesCTypeChoiceEnum string

// Enum of Choice on FlowRSVPPathObjectsTimeValuesCType
var FlowRSVPPathObjectsTimeValuesCTypeChoice = struct {
	TYPE_1 FlowRSVPPathObjectsTimeValuesCTypeChoiceEnum
}{
	TYPE_1: FlowRSVPPathObjectsTimeValuesCTypeChoiceEnum("type_1"),
}

func (obj *flowRSVPPathObjectsTimeValuesCType) Choice() FlowRSVPPathObjectsTimeValuesCTypeChoiceEnum {
	return FlowRSVPPathObjectsTimeValuesCTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowRSVPPathObjectsTimeValuesCType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPPathObjectsTimeValuesCType) setChoice(value FlowRSVPPathObjectsTimeValuesCTypeChoiceEnum) FlowRSVPPathObjectsTimeValuesCType {
	intValue, ok := otg.FlowRSVPPathObjectsTimeValuesCType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPPathObjectsTimeValuesCTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPPathObjectsTimeValuesCType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Type_1 = nil
	obj.type_1Holder = nil

	if value == FlowRSVPPathObjectsTimeValuesCTypeChoice.TYPE_1 {
		obj.obj.Type_1 = NewFlowRSVPPathTimeValuesType1().msg()
	}

	return obj
}

// description is TBD
// Type1 returns a FlowRSVPPathTimeValuesType1
func (obj *flowRSVPPathObjectsTimeValuesCType) Type1() FlowRSVPPathTimeValuesType1 {
	if obj.obj.Type_1 == nil {
		obj.setChoice(FlowRSVPPathObjectsTimeValuesCTypeChoice.TYPE_1)
	}
	if obj.type_1Holder == nil {
		obj.type_1Holder = &flowRSVPPathTimeValuesType1{obj: obj.obj.Type_1}
	}
	return obj.type_1Holder
}

// description is TBD
// Type1 returns a FlowRSVPPathTimeValuesType1
func (obj *flowRSVPPathObjectsTimeValuesCType) HasType1() bool {
	return obj.obj.Type_1 != nil
}

// description is TBD
// SetType1 sets the FlowRSVPPathTimeValuesType1 value in the FlowRSVPPathObjectsTimeValuesCType object
func (obj *flowRSVPPathObjectsTimeValuesCType) SetType1(value FlowRSVPPathTimeValuesType1) FlowRSVPPathObjectsTimeValuesCType {
	obj.setChoice(FlowRSVPPathObjectsTimeValuesCTypeChoice.TYPE_1)
	obj.type_1Holder = nil
	obj.obj.Type_1 = value.msg()

	return obj
}

func (obj *flowRSVPPathObjectsTimeValuesCType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Type_1 != nil {

		obj.Type1().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathObjectsTimeValuesCType) setDefault() {
	var choices_set int = 0
	var choice FlowRSVPPathObjectsTimeValuesCTypeChoiceEnum

	if obj.obj.Type_1 != nil {
		choices_set += 1
		choice = FlowRSVPPathObjectsTimeValuesCTypeChoice.TYPE_1
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowRSVPPathObjectsTimeValuesCTypeChoice.TYPE_1)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowRSVPPathObjectsTimeValuesCType")
			}
		} else {
			intVal := otg.FlowRSVPPathObjectsTimeValuesCType_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowRSVPPathObjectsTimeValuesCType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

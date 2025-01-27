package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjectsClassExplicitRouteCType *****
type flowRSVPPathObjectsClassExplicitRouteCType struct {
	validation
	obj          *otg.FlowRSVPPathObjectsClassExplicitRouteCType
	marshaller   marshalFlowRSVPPathObjectsClassExplicitRouteCType
	unMarshaller unMarshalFlowRSVPPathObjectsClassExplicitRouteCType
	type_1Holder FlowRSVPPathExplicitRouteType1
}

func NewFlowRSVPPathObjectsClassExplicitRouteCType() FlowRSVPPathObjectsClassExplicitRouteCType {
	obj := flowRSVPPathObjectsClassExplicitRouteCType{obj: &otg.FlowRSVPPathObjectsClassExplicitRouteCType{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjectsClassExplicitRouteCType) msg() *otg.FlowRSVPPathObjectsClassExplicitRouteCType {
	return obj.obj
}

func (obj *flowRSVPPathObjectsClassExplicitRouteCType) setMsg(msg *otg.FlowRSVPPathObjectsClassExplicitRouteCType) FlowRSVPPathObjectsClassExplicitRouteCType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjectsClassExplicitRouteCType struct {
	obj *flowRSVPPathObjectsClassExplicitRouteCType
}

type marshalFlowRSVPPathObjectsClassExplicitRouteCType interface {
	// ToProto marshals FlowRSVPPathObjectsClassExplicitRouteCType to protobuf object *otg.FlowRSVPPathObjectsClassExplicitRouteCType
	ToProto() (*otg.FlowRSVPPathObjectsClassExplicitRouteCType, error)
	// ToPbText marshals FlowRSVPPathObjectsClassExplicitRouteCType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjectsClassExplicitRouteCType to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjectsClassExplicitRouteCType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowRSVPPathObjectsClassExplicitRouteCType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowRSVPPathObjectsClassExplicitRouteCType struct {
	obj *flowRSVPPathObjectsClassExplicitRouteCType
}

type unMarshalFlowRSVPPathObjectsClassExplicitRouteCType interface {
	// FromProto unmarshals FlowRSVPPathObjectsClassExplicitRouteCType from protobuf object *otg.FlowRSVPPathObjectsClassExplicitRouteCType
	FromProto(msg *otg.FlowRSVPPathObjectsClassExplicitRouteCType) (FlowRSVPPathObjectsClassExplicitRouteCType, error)
	// FromPbText unmarshals FlowRSVPPathObjectsClassExplicitRouteCType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjectsClassExplicitRouteCType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjectsClassExplicitRouteCType from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjectsClassExplicitRouteCType) Marshal() marshalFlowRSVPPathObjectsClassExplicitRouteCType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjectsClassExplicitRouteCType{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjectsClassExplicitRouteCType) Unmarshal() unMarshalFlowRSVPPathObjectsClassExplicitRouteCType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjectsClassExplicitRouteCType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjectsClassExplicitRouteCType) ToProto() (*otg.FlowRSVPPathObjectsClassExplicitRouteCType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjectsClassExplicitRouteCType) FromProto(msg *otg.FlowRSVPPathObjectsClassExplicitRouteCType) (FlowRSVPPathObjectsClassExplicitRouteCType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjectsClassExplicitRouteCType) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassExplicitRouteCType) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjectsClassExplicitRouteCType) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassExplicitRouteCType) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjectsClassExplicitRouteCType) ToJsonRaw() (string, error) {
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

func (m *marshalflowRSVPPathObjectsClassExplicitRouteCType) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassExplicitRouteCType) FromJson(value string) error {
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

func (obj *flowRSVPPathObjectsClassExplicitRouteCType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClassExplicitRouteCType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClassExplicitRouteCType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjectsClassExplicitRouteCType) Clone() (FlowRSVPPathObjectsClassExplicitRouteCType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjectsClassExplicitRouteCType()
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

func (obj *flowRSVPPathObjectsClassExplicitRouteCType) setNil() {
	obj.type_1Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjectsClassExplicitRouteCType is object for EXPLICIT_ROUTE class and c-type is Type 1 Explicit Route (1).
type FlowRSVPPathObjectsClassExplicitRouteCType interface {
	Validation
	// msg marshals FlowRSVPPathObjectsClassExplicitRouteCType to protobuf object *otg.FlowRSVPPathObjectsClassExplicitRouteCType
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjectsClassExplicitRouteCType
	// setMsg unmarshals FlowRSVPPathObjectsClassExplicitRouteCType from protobuf object *otg.FlowRSVPPathObjectsClassExplicitRouteCType
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjectsClassExplicitRouteCType) FlowRSVPPathObjectsClassExplicitRouteCType
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjectsClassExplicitRouteCType
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjectsClassExplicitRouteCType
	// validate validates FlowRSVPPathObjectsClassExplicitRouteCType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjectsClassExplicitRouteCType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPPathObjectsClassExplicitRouteCTypeChoiceEnum, set in FlowRSVPPathObjectsClassExplicitRouteCType
	Choice() FlowRSVPPathObjectsClassExplicitRouteCTypeChoiceEnum
	// setChoice assigns FlowRSVPPathObjectsClassExplicitRouteCTypeChoiceEnum provided by user to FlowRSVPPathObjectsClassExplicitRouteCType
	setChoice(value FlowRSVPPathObjectsClassExplicitRouteCTypeChoiceEnum) FlowRSVPPathObjectsClassExplicitRouteCType
	// HasChoice checks if Choice has been set in FlowRSVPPathObjectsClassExplicitRouteCType
	HasChoice() bool
	// Type1 returns FlowRSVPPathExplicitRouteType1, set in FlowRSVPPathObjectsClassExplicitRouteCType.
	// FlowRSVPPathExplicitRouteType1 is type1 Explicit Route has subobjects. Currently supported subobjects are IPv4 prefix and Autonomous system number.
	Type1() FlowRSVPPathExplicitRouteType1
	// SetType1 assigns FlowRSVPPathExplicitRouteType1 provided by user to FlowRSVPPathObjectsClassExplicitRouteCType.
	// FlowRSVPPathExplicitRouteType1 is type1 Explicit Route has subobjects. Currently supported subobjects are IPv4 prefix and Autonomous system number.
	SetType1(value FlowRSVPPathExplicitRouteType1) FlowRSVPPathObjectsClassExplicitRouteCType
	// HasType1 checks if Type1 has been set in FlowRSVPPathObjectsClassExplicitRouteCType
	HasType1() bool
	setNil()
}

type FlowRSVPPathObjectsClassExplicitRouteCTypeChoiceEnum string

// Enum of Choice on FlowRSVPPathObjectsClassExplicitRouteCType
var FlowRSVPPathObjectsClassExplicitRouteCTypeChoice = struct {
	TYPE_1 FlowRSVPPathObjectsClassExplicitRouteCTypeChoiceEnum
}{
	TYPE_1: FlowRSVPPathObjectsClassExplicitRouteCTypeChoiceEnum("type_1"),
}

func (obj *flowRSVPPathObjectsClassExplicitRouteCType) Choice() FlowRSVPPathObjectsClassExplicitRouteCTypeChoiceEnum {
	return FlowRSVPPathObjectsClassExplicitRouteCTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowRSVPPathObjectsClassExplicitRouteCType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPPathObjectsClassExplicitRouteCType) setChoice(value FlowRSVPPathObjectsClassExplicitRouteCTypeChoiceEnum) FlowRSVPPathObjectsClassExplicitRouteCType {
	intValue, ok := otg.FlowRSVPPathObjectsClassExplicitRouteCType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPPathObjectsClassExplicitRouteCTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPPathObjectsClassExplicitRouteCType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Type_1 = nil
	obj.type_1Holder = nil

	if value == FlowRSVPPathObjectsClassExplicitRouteCTypeChoice.TYPE_1 {
		obj.obj.Type_1 = NewFlowRSVPPathExplicitRouteType1().msg()
	}

	return obj
}

// description is TBD
// Type1 returns a FlowRSVPPathExplicitRouteType1
func (obj *flowRSVPPathObjectsClassExplicitRouteCType) Type1() FlowRSVPPathExplicitRouteType1 {
	if obj.obj.Type_1 == nil {
		obj.setChoice(FlowRSVPPathObjectsClassExplicitRouteCTypeChoice.TYPE_1)
	}
	if obj.type_1Holder == nil {
		obj.type_1Holder = &flowRSVPPathExplicitRouteType1{obj: obj.obj.Type_1}
	}
	return obj.type_1Holder
}

// description is TBD
// Type1 returns a FlowRSVPPathExplicitRouteType1
func (obj *flowRSVPPathObjectsClassExplicitRouteCType) HasType1() bool {
	return obj.obj.Type_1 != nil
}

// description is TBD
// SetType1 sets the FlowRSVPPathExplicitRouteType1 value in the FlowRSVPPathObjectsClassExplicitRouteCType object
func (obj *flowRSVPPathObjectsClassExplicitRouteCType) SetType1(value FlowRSVPPathExplicitRouteType1) FlowRSVPPathObjectsClassExplicitRouteCType {
	obj.setChoice(FlowRSVPPathObjectsClassExplicitRouteCTypeChoice.TYPE_1)
	obj.type_1Holder = nil
	obj.obj.Type_1 = value.msg()

	return obj
}

func (obj *flowRSVPPathObjectsClassExplicitRouteCType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Type_1 != nil {

		obj.Type1().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathObjectsClassExplicitRouteCType) setDefault() {
	var choices_set int = 0
	var choice FlowRSVPPathObjectsClassExplicitRouteCTypeChoiceEnum

	if obj.obj.Type_1 != nil {
		choices_set += 1
		choice = FlowRSVPPathObjectsClassExplicitRouteCTypeChoice.TYPE_1
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowRSVPPathObjectsClassExplicitRouteCTypeChoice.TYPE_1)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowRSVPPathObjectsClassExplicitRouteCType")
			}
		} else {
			intVal := otg.FlowRSVPPathObjectsClassExplicitRouteCType_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowRSVPPathObjectsClassExplicitRouteCType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

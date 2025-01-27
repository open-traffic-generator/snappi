package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjectsRecordRouteCType *****
type flowRSVPPathObjectsRecordRouteCType struct {
	validation
	obj          *otg.FlowRSVPPathObjectsRecordRouteCType
	marshaller   marshalFlowRSVPPathObjectsRecordRouteCType
	unMarshaller unMarshalFlowRSVPPathObjectsRecordRouteCType
	type_1Holder FlowRSVPPathRecordRouteType1
}

func NewFlowRSVPPathObjectsRecordRouteCType() FlowRSVPPathObjectsRecordRouteCType {
	obj := flowRSVPPathObjectsRecordRouteCType{obj: &otg.FlowRSVPPathObjectsRecordRouteCType{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjectsRecordRouteCType) msg() *otg.FlowRSVPPathObjectsRecordRouteCType {
	return obj.obj
}

func (obj *flowRSVPPathObjectsRecordRouteCType) setMsg(msg *otg.FlowRSVPPathObjectsRecordRouteCType) FlowRSVPPathObjectsRecordRouteCType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjectsRecordRouteCType struct {
	obj *flowRSVPPathObjectsRecordRouteCType
}

type marshalFlowRSVPPathObjectsRecordRouteCType interface {
	// ToProto marshals FlowRSVPPathObjectsRecordRouteCType to protobuf object *otg.FlowRSVPPathObjectsRecordRouteCType
	ToProto() (*otg.FlowRSVPPathObjectsRecordRouteCType, error)
	// ToPbText marshals FlowRSVPPathObjectsRecordRouteCType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjectsRecordRouteCType to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjectsRecordRouteCType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowRSVPPathObjectsRecordRouteCType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowRSVPPathObjectsRecordRouteCType struct {
	obj *flowRSVPPathObjectsRecordRouteCType
}

type unMarshalFlowRSVPPathObjectsRecordRouteCType interface {
	// FromProto unmarshals FlowRSVPPathObjectsRecordRouteCType from protobuf object *otg.FlowRSVPPathObjectsRecordRouteCType
	FromProto(msg *otg.FlowRSVPPathObjectsRecordRouteCType) (FlowRSVPPathObjectsRecordRouteCType, error)
	// FromPbText unmarshals FlowRSVPPathObjectsRecordRouteCType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjectsRecordRouteCType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjectsRecordRouteCType from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjectsRecordRouteCType) Marshal() marshalFlowRSVPPathObjectsRecordRouteCType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjectsRecordRouteCType{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjectsRecordRouteCType) Unmarshal() unMarshalFlowRSVPPathObjectsRecordRouteCType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjectsRecordRouteCType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjectsRecordRouteCType) ToProto() (*otg.FlowRSVPPathObjectsRecordRouteCType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjectsRecordRouteCType) FromProto(msg *otg.FlowRSVPPathObjectsRecordRouteCType) (FlowRSVPPathObjectsRecordRouteCType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjectsRecordRouteCType) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsRecordRouteCType) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjectsRecordRouteCType) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsRecordRouteCType) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjectsRecordRouteCType) ToJsonRaw() (string, error) {
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

func (m *marshalflowRSVPPathObjectsRecordRouteCType) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsRecordRouteCType) FromJson(value string) error {
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

func (obj *flowRSVPPathObjectsRecordRouteCType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsRecordRouteCType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsRecordRouteCType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjectsRecordRouteCType) Clone() (FlowRSVPPathObjectsRecordRouteCType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjectsRecordRouteCType()
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

func (obj *flowRSVPPathObjectsRecordRouteCType) setNil() {
	obj.type_1Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjectsRecordRouteCType is object for RECORD_ROUTE class. c-type is Type 1 Route Record (1).
type FlowRSVPPathObjectsRecordRouteCType interface {
	Validation
	// msg marshals FlowRSVPPathObjectsRecordRouteCType to protobuf object *otg.FlowRSVPPathObjectsRecordRouteCType
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjectsRecordRouteCType
	// setMsg unmarshals FlowRSVPPathObjectsRecordRouteCType from protobuf object *otg.FlowRSVPPathObjectsRecordRouteCType
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjectsRecordRouteCType) FlowRSVPPathObjectsRecordRouteCType
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjectsRecordRouteCType
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjectsRecordRouteCType
	// validate validates FlowRSVPPathObjectsRecordRouteCType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjectsRecordRouteCType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPPathObjectsRecordRouteCTypeChoiceEnum, set in FlowRSVPPathObjectsRecordRouteCType
	Choice() FlowRSVPPathObjectsRecordRouteCTypeChoiceEnum
	// setChoice assigns FlowRSVPPathObjectsRecordRouteCTypeChoiceEnum provided by user to FlowRSVPPathObjectsRecordRouteCType
	setChoice(value FlowRSVPPathObjectsRecordRouteCTypeChoiceEnum) FlowRSVPPathObjectsRecordRouteCType
	// HasChoice checks if Choice has been set in FlowRSVPPathObjectsRecordRouteCType
	HasChoice() bool
	// Type1 returns FlowRSVPPathRecordRouteType1, set in FlowRSVPPathObjectsRecordRouteCType.
	// FlowRSVPPathRecordRouteType1 is type1 record route has list of subobjects. Currently supported subobjects are IPv4 address(1) and Label(3).
	Type1() FlowRSVPPathRecordRouteType1
	// SetType1 assigns FlowRSVPPathRecordRouteType1 provided by user to FlowRSVPPathObjectsRecordRouteCType.
	// FlowRSVPPathRecordRouteType1 is type1 record route has list of subobjects. Currently supported subobjects are IPv4 address(1) and Label(3).
	SetType1(value FlowRSVPPathRecordRouteType1) FlowRSVPPathObjectsRecordRouteCType
	// HasType1 checks if Type1 has been set in FlowRSVPPathObjectsRecordRouteCType
	HasType1() bool
	setNil()
}

type FlowRSVPPathObjectsRecordRouteCTypeChoiceEnum string

// Enum of Choice on FlowRSVPPathObjectsRecordRouteCType
var FlowRSVPPathObjectsRecordRouteCTypeChoice = struct {
	TYPE_1 FlowRSVPPathObjectsRecordRouteCTypeChoiceEnum
}{
	TYPE_1: FlowRSVPPathObjectsRecordRouteCTypeChoiceEnum("type_1"),
}

func (obj *flowRSVPPathObjectsRecordRouteCType) Choice() FlowRSVPPathObjectsRecordRouteCTypeChoiceEnum {
	return FlowRSVPPathObjectsRecordRouteCTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowRSVPPathObjectsRecordRouteCType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPPathObjectsRecordRouteCType) setChoice(value FlowRSVPPathObjectsRecordRouteCTypeChoiceEnum) FlowRSVPPathObjectsRecordRouteCType {
	intValue, ok := otg.FlowRSVPPathObjectsRecordRouteCType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPPathObjectsRecordRouteCTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPPathObjectsRecordRouteCType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Type_1 = nil
	obj.type_1Holder = nil

	if value == FlowRSVPPathObjectsRecordRouteCTypeChoice.TYPE_1 {
		obj.obj.Type_1 = NewFlowRSVPPathRecordRouteType1().msg()
	}

	return obj
}

// description is TBD
// Type1 returns a FlowRSVPPathRecordRouteType1
func (obj *flowRSVPPathObjectsRecordRouteCType) Type1() FlowRSVPPathRecordRouteType1 {
	if obj.obj.Type_1 == nil {
		obj.setChoice(FlowRSVPPathObjectsRecordRouteCTypeChoice.TYPE_1)
	}
	if obj.type_1Holder == nil {
		obj.type_1Holder = &flowRSVPPathRecordRouteType1{obj: obj.obj.Type_1}
	}
	return obj.type_1Holder
}

// description is TBD
// Type1 returns a FlowRSVPPathRecordRouteType1
func (obj *flowRSVPPathObjectsRecordRouteCType) HasType1() bool {
	return obj.obj.Type_1 != nil
}

// description is TBD
// SetType1 sets the FlowRSVPPathRecordRouteType1 value in the FlowRSVPPathObjectsRecordRouteCType object
func (obj *flowRSVPPathObjectsRecordRouteCType) SetType1(value FlowRSVPPathRecordRouteType1) FlowRSVPPathObjectsRecordRouteCType {
	obj.setChoice(FlowRSVPPathObjectsRecordRouteCTypeChoice.TYPE_1)
	obj.type_1Holder = nil
	obj.obj.Type_1 = value.msg()

	return obj
}

func (obj *flowRSVPPathObjectsRecordRouteCType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Type_1 != nil {

		obj.Type1().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathObjectsRecordRouteCType) setDefault() {
	var choices_set int = 0
	var choice FlowRSVPPathObjectsRecordRouteCTypeChoiceEnum

	if obj.obj.Type_1 != nil {
		choices_set += 1
		choice = FlowRSVPPathObjectsRecordRouteCTypeChoice.TYPE_1
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowRSVPPathObjectsRecordRouteCTypeChoice.TYPE_1)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowRSVPPathObjectsRecordRouteCType")
			}
		} else {
			intVal := otg.FlowRSVPPathObjectsRecordRouteCType_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowRSVPPathObjectsRecordRouteCType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

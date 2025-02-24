package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPType1ExplicitRouteSubobjectsType *****
type flowRSVPType1ExplicitRouteSubobjectsType struct {
	validation
	obj              *otg.FlowRSVPType1ExplicitRouteSubobjectsType
	marshaller       marshalFlowRSVPType1ExplicitRouteSubobjectsType
	unMarshaller     unMarshalFlowRSVPType1ExplicitRouteSubobjectsType
	ipv4PrefixHolder FlowRSVPPathExplicitRouteType1Ipv4Prefix
	asNumberHolder   FlowRSVPPathExplicitRouteType1ASNumber
}

func NewFlowRSVPType1ExplicitRouteSubobjectsType() FlowRSVPType1ExplicitRouteSubobjectsType {
	obj := flowRSVPType1ExplicitRouteSubobjectsType{obj: &otg.FlowRSVPType1ExplicitRouteSubobjectsType{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPType1ExplicitRouteSubobjectsType) msg() *otg.FlowRSVPType1ExplicitRouteSubobjectsType {
	return obj.obj
}

func (obj *flowRSVPType1ExplicitRouteSubobjectsType) setMsg(msg *otg.FlowRSVPType1ExplicitRouteSubobjectsType) FlowRSVPType1ExplicitRouteSubobjectsType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPType1ExplicitRouteSubobjectsType struct {
	obj *flowRSVPType1ExplicitRouteSubobjectsType
}

type marshalFlowRSVPType1ExplicitRouteSubobjectsType interface {
	// ToProto marshals FlowRSVPType1ExplicitRouteSubobjectsType to protobuf object *otg.FlowRSVPType1ExplicitRouteSubobjectsType
	ToProto() (*otg.FlowRSVPType1ExplicitRouteSubobjectsType, error)
	// ToPbText marshals FlowRSVPType1ExplicitRouteSubobjectsType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPType1ExplicitRouteSubobjectsType to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPType1ExplicitRouteSubobjectsType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowRSVPType1ExplicitRouteSubobjectsType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowRSVPType1ExplicitRouteSubobjectsType struct {
	obj *flowRSVPType1ExplicitRouteSubobjectsType
}

type unMarshalFlowRSVPType1ExplicitRouteSubobjectsType interface {
	// FromProto unmarshals FlowRSVPType1ExplicitRouteSubobjectsType from protobuf object *otg.FlowRSVPType1ExplicitRouteSubobjectsType
	FromProto(msg *otg.FlowRSVPType1ExplicitRouteSubobjectsType) (FlowRSVPType1ExplicitRouteSubobjectsType, error)
	// FromPbText unmarshals FlowRSVPType1ExplicitRouteSubobjectsType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPType1ExplicitRouteSubobjectsType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPType1ExplicitRouteSubobjectsType from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPType1ExplicitRouteSubobjectsType) Marshal() marshalFlowRSVPType1ExplicitRouteSubobjectsType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPType1ExplicitRouteSubobjectsType{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPType1ExplicitRouteSubobjectsType) Unmarshal() unMarshalFlowRSVPType1ExplicitRouteSubobjectsType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPType1ExplicitRouteSubobjectsType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPType1ExplicitRouteSubobjectsType) ToProto() (*otg.FlowRSVPType1ExplicitRouteSubobjectsType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPType1ExplicitRouteSubobjectsType) FromProto(msg *otg.FlowRSVPType1ExplicitRouteSubobjectsType) (FlowRSVPType1ExplicitRouteSubobjectsType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPType1ExplicitRouteSubobjectsType) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPType1ExplicitRouteSubobjectsType) FromPbText(value string) error {
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

func (m *marshalflowRSVPType1ExplicitRouteSubobjectsType) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPType1ExplicitRouteSubobjectsType) FromYaml(value string) error {
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

func (m *marshalflowRSVPType1ExplicitRouteSubobjectsType) ToJsonRaw() (string, error) {
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

func (m *marshalflowRSVPType1ExplicitRouteSubobjectsType) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPType1ExplicitRouteSubobjectsType) FromJson(value string) error {
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

func (obj *flowRSVPType1ExplicitRouteSubobjectsType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPType1ExplicitRouteSubobjectsType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPType1ExplicitRouteSubobjectsType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPType1ExplicitRouteSubobjectsType) Clone() (FlowRSVPType1ExplicitRouteSubobjectsType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPType1ExplicitRouteSubobjectsType()
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

func (obj *flowRSVPType1ExplicitRouteSubobjectsType) setNil() {
	obj.ipv4PrefixHolder = nil
	obj.asNumberHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPType1ExplicitRouteSubobjectsType is currently supported subobjects are IPv4 address(1) and Autonomous system number(32).
type FlowRSVPType1ExplicitRouteSubobjectsType interface {
	Validation
	// msg marshals FlowRSVPType1ExplicitRouteSubobjectsType to protobuf object *otg.FlowRSVPType1ExplicitRouteSubobjectsType
	// and doesn't set defaults
	msg() *otg.FlowRSVPType1ExplicitRouteSubobjectsType
	// setMsg unmarshals FlowRSVPType1ExplicitRouteSubobjectsType from protobuf object *otg.FlowRSVPType1ExplicitRouteSubobjectsType
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPType1ExplicitRouteSubobjectsType) FlowRSVPType1ExplicitRouteSubobjectsType
	// provides marshal interface
	Marshal() marshalFlowRSVPType1ExplicitRouteSubobjectsType
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPType1ExplicitRouteSubobjectsType
	// validate validates FlowRSVPType1ExplicitRouteSubobjectsType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPType1ExplicitRouteSubobjectsType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPType1ExplicitRouteSubobjectsTypeChoiceEnum, set in FlowRSVPType1ExplicitRouteSubobjectsType
	Choice() FlowRSVPType1ExplicitRouteSubobjectsTypeChoiceEnum
	// setChoice assigns FlowRSVPType1ExplicitRouteSubobjectsTypeChoiceEnum provided by user to FlowRSVPType1ExplicitRouteSubobjectsType
	setChoice(value FlowRSVPType1ExplicitRouteSubobjectsTypeChoiceEnum) FlowRSVPType1ExplicitRouteSubobjectsType
	// HasChoice checks if Choice has been set in FlowRSVPType1ExplicitRouteSubobjectsType
	HasChoice() bool
	// Ipv4Prefix returns FlowRSVPPathExplicitRouteType1Ipv4Prefix, set in FlowRSVPType1ExplicitRouteSubobjectsType.
	// FlowRSVPPathExplicitRouteType1Ipv4Prefix is class = EXPLICIT_ROUTE, Type1 ROUTE_RECORD C-Type = 1 Subobject: IPv4 Prefix, C-Type: 1
	Ipv4Prefix() FlowRSVPPathExplicitRouteType1Ipv4Prefix
	// SetIpv4Prefix assigns FlowRSVPPathExplicitRouteType1Ipv4Prefix provided by user to FlowRSVPType1ExplicitRouteSubobjectsType.
	// FlowRSVPPathExplicitRouteType1Ipv4Prefix is class = EXPLICIT_ROUTE, Type1 ROUTE_RECORD C-Type = 1 Subobject: IPv4 Prefix, C-Type: 1
	SetIpv4Prefix(value FlowRSVPPathExplicitRouteType1Ipv4Prefix) FlowRSVPType1ExplicitRouteSubobjectsType
	// HasIpv4Prefix checks if Ipv4Prefix has been set in FlowRSVPType1ExplicitRouteSubobjectsType
	HasIpv4Prefix() bool
	// AsNumber returns FlowRSVPPathExplicitRouteType1ASNumber, set in FlowRSVPType1ExplicitRouteSubobjectsType.
	// FlowRSVPPathExplicitRouteType1ASNumber is class = EXPLICIT_ROUTE, Type1 ROUTE_RECORD C-Type = 1 Subobject: Autonomous system number, C-Type: 32
	AsNumber() FlowRSVPPathExplicitRouteType1ASNumber
	// SetAsNumber assigns FlowRSVPPathExplicitRouteType1ASNumber provided by user to FlowRSVPType1ExplicitRouteSubobjectsType.
	// FlowRSVPPathExplicitRouteType1ASNumber is class = EXPLICIT_ROUTE, Type1 ROUTE_RECORD C-Type = 1 Subobject: Autonomous system number, C-Type: 32
	SetAsNumber(value FlowRSVPPathExplicitRouteType1ASNumber) FlowRSVPType1ExplicitRouteSubobjectsType
	// HasAsNumber checks if AsNumber has been set in FlowRSVPType1ExplicitRouteSubobjectsType
	HasAsNumber() bool
	setNil()
}

type FlowRSVPType1ExplicitRouteSubobjectsTypeChoiceEnum string

// Enum of Choice on FlowRSVPType1ExplicitRouteSubobjectsType
var FlowRSVPType1ExplicitRouteSubobjectsTypeChoice = struct {
	IPV4_PREFIX FlowRSVPType1ExplicitRouteSubobjectsTypeChoiceEnum
	AS_NUMBER   FlowRSVPType1ExplicitRouteSubobjectsTypeChoiceEnum
}{
	IPV4_PREFIX: FlowRSVPType1ExplicitRouteSubobjectsTypeChoiceEnum("ipv4_prefix"),
	AS_NUMBER:   FlowRSVPType1ExplicitRouteSubobjectsTypeChoiceEnum("as_number"),
}

func (obj *flowRSVPType1ExplicitRouteSubobjectsType) Choice() FlowRSVPType1ExplicitRouteSubobjectsTypeChoiceEnum {
	return FlowRSVPType1ExplicitRouteSubobjectsTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowRSVPType1ExplicitRouteSubobjectsType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPType1ExplicitRouteSubobjectsType) setChoice(value FlowRSVPType1ExplicitRouteSubobjectsTypeChoiceEnum) FlowRSVPType1ExplicitRouteSubobjectsType {
	intValue, ok := otg.FlowRSVPType1ExplicitRouteSubobjectsType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPType1ExplicitRouteSubobjectsTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPType1ExplicitRouteSubobjectsType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.AsNumber = nil
	obj.asNumberHolder = nil
	obj.obj.Ipv4Prefix = nil
	obj.ipv4PrefixHolder = nil

	if value == FlowRSVPType1ExplicitRouteSubobjectsTypeChoice.IPV4_PREFIX {
		obj.obj.Ipv4Prefix = NewFlowRSVPPathExplicitRouteType1Ipv4Prefix().msg()
	}

	if value == FlowRSVPType1ExplicitRouteSubobjectsTypeChoice.AS_NUMBER {
		obj.obj.AsNumber = NewFlowRSVPPathExplicitRouteType1ASNumber().msg()
	}

	return obj
}

// description is TBD
// Ipv4Prefix returns a FlowRSVPPathExplicitRouteType1Ipv4Prefix
func (obj *flowRSVPType1ExplicitRouteSubobjectsType) Ipv4Prefix() FlowRSVPPathExplicitRouteType1Ipv4Prefix {
	if obj.obj.Ipv4Prefix == nil {
		obj.setChoice(FlowRSVPType1ExplicitRouteSubobjectsTypeChoice.IPV4_PREFIX)
	}
	if obj.ipv4PrefixHolder == nil {
		obj.ipv4PrefixHolder = &flowRSVPPathExplicitRouteType1Ipv4Prefix{obj: obj.obj.Ipv4Prefix}
	}
	return obj.ipv4PrefixHolder
}

// description is TBD
// Ipv4Prefix returns a FlowRSVPPathExplicitRouteType1Ipv4Prefix
func (obj *flowRSVPType1ExplicitRouteSubobjectsType) HasIpv4Prefix() bool {
	return obj.obj.Ipv4Prefix != nil
}

// description is TBD
// SetIpv4Prefix sets the FlowRSVPPathExplicitRouteType1Ipv4Prefix value in the FlowRSVPType1ExplicitRouteSubobjectsType object
func (obj *flowRSVPType1ExplicitRouteSubobjectsType) SetIpv4Prefix(value FlowRSVPPathExplicitRouteType1Ipv4Prefix) FlowRSVPType1ExplicitRouteSubobjectsType {
	obj.setChoice(FlowRSVPType1ExplicitRouteSubobjectsTypeChoice.IPV4_PREFIX)
	obj.ipv4PrefixHolder = nil
	obj.obj.Ipv4Prefix = value.msg()

	return obj
}

// description is TBD
// AsNumber returns a FlowRSVPPathExplicitRouteType1ASNumber
func (obj *flowRSVPType1ExplicitRouteSubobjectsType) AsNumber() FlowRSVPPathExplicitRouteType1ASNumber {
	if obj.obj.AsNumber == nil {
		obj.setChoice(FlowRSVPType1ExplicitRouteSubobjectsTypeChoice.AS_NUMBER)
	}
	if obj.asNumberHolder == nil {
		obj.asNumberHolder = &flowRSVPPathExplicitRouteType1ASNumber{obj: obj.obj.AsNumber}
	}
	return obj.asNumberHolder
}

// description is TBD
// AsNumber returns a FlowRSVPPathExplicitRouteType1ASNumber
func (obj *flowRSVPType1ExplicitRouteSubobjectsType) HasAsNumber() bool {
	return obj.obj.AsNumber != nil
}

// description is TBD
// SetAsNumber sets the FlowRSVPPathExplicitRouteType1ASNumber value in the FlowRSVPType1ExplicitRouteSubobjectsType object
func (obj *flowRSVPType1ExplicitRouteSubobjectsType) SetAsNumber(value FlowRSVPPathExplicitRouteType1ASNumber) FlowRSVPType1ExplicitRouteSubobjectsType {
	obj.setChoice(FlowRSVPType1ExplicitRouteSubobjectsTypeChoice.AS_NUMBER)
	obj.asNumberHolder = nil
	obj.obj.AsNumber = value.msg()

	return obj
}

func (obj *flowRSVPType1ExplicitRouteSubobjectsType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv4Prefix != nil {

		obj.Ipv4Prefix().validateObj(vObj, set_default)
	}

	if obj.obj.AsNumber != nil {

		obj.AsNumber().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPType1ExplicitRouteSubobjectsType) setDefault() {
	var choices_set int = 0
	var choice FlowRSVPType1ExplicitRouteSubobjectsTypeChoiceEnum

	if obj.obj.Ipv4Prefix != nil {
		choices_set += 1
		choice = FlowRSVPType1ExplicitRouteSubobjectsTypeChoice.IPV4_PREFIX
	}

	if obj.obj.AsNumber != nil {
		choices_set += 1
		choice = FlowRSVPType1ExplicitRouteSubobjectsTypeChoice.AS_NUMBER
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowRSVPType1ExplicitRouteSubobjectsTypeChoice.IPV4_PREFIX)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowRSVPType1ExplicitRouteSubobjectsType")
			}
		} else {
			intVal := otg.FlowRSVPType1ExplicitRouteSubobjectsType_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowRSVPType1ExplicitRouteSubobjectsType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

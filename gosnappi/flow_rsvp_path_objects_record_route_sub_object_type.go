package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjectsRecordRouteSubObjectType *****
type flowRSVPPathObjectsRecordRouteSubObjectType struct {
	validation
	obj               *otg.FlowRSVPPathObjectsRecordRouteSubObjectType
	marshaller        marshalFlowRSVPPathObjectsRecordRouteSubObjectType
	unMarshaller      unMarshalFlowRSVPPathObjectsRecordRouteSubObjectType
	ipv4AddressHolder FlowRSVPPathRecordRouteType1Ipv4Address
	labelHolder       FlowRSVPPathRecordRouteType1Label
}

func NewFlowRSVPPathObjectsRecordRouteSubObjectType() FlowRSVPPathObjectsRecordRouteSubObjectType {
	obj := flowRSVPPathObjectsRecordRouteSubObjectType{obj: &otg.FlowRSVPPathObjectsRecordRouteSubObjectType{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjectsRecordRouteSubObjectType) msg() *otg.FlowRSVPPathObjectsRecordRouteSubObjectType {
	return obj.obj
}

func (obj *flowRSVPPathObjectsRecordRouteSubObjectType) setMsg(msg *otg.FlowRSVPPathObjectsRecordRouteSubObjectType) FlowRSVPPathObjectsRecordRouteSubObjectType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjectsRecordRouteSubObjectType struct {
	obj *flowRSVPPathObjectsRecordRouteSubObjectType
}

type marshalFlowRSVPPathObjectsRecordRouteSubObjectType interface {
	// ToProto marshals FlowRSVPPathObjectsRecordRouteSubObjectType to protobuf object *otg.FlowRSVPPathObjectsRecordRouteSubObjectType
	ToProto() (*otg.FlowRSVPPathObjectsRecordRouteSubObjectType, error)
	// ToPbText marshals FlowRSVPPathObjectsRecordRouteSubObjectType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjectsRecordRouteSubObjectType to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjectsRecordRouteSubObjectType to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathObjectsRecordRouteSubObjectType struct {
	obj *flowRSVPPathObjectsRecordRouteSubObjectType
}

type unMarshalFlowRSVPPathObjectsRecordRouteSubObjectType interface {
	// FromProto unmarshals FlowRSVPPathObjectsRecordRouteSubObjectType from protobuf object *otg.FlowRSVPPathObjectsRecordRouteSubObjectType
	FromProto(msg *otg.FlowRSVPPathObjectsRecordRouteSubObjectType) (FlowRSVPPathObjectsRecordRouteSubObjectType, error)
	// FromPbText unmarshals FlowRSVPPathObjectsRecordRouteSubObjectType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjectsRecordRouteSubObjectType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjectsRecordRouteSubObjectType from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjectsRecordRouteSubObjectType) Marshal() marshalFlowRSVPPathObjectsRecordRouteSubObjectType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjectsRecordRouteSubObjectType{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjectsRecordRouteSubObjectType) Unmarshal() unMarshalFlowRSVPPathObjectsRecordRouteSubObjectType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjectsRecordRouteSubObjectType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjectsRecordRouteSubObjectType) ToProto() (*otg.FlowRSVPPathObjectsRecordRouteSubObjectType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjectsRecordRouteSubObjectType) FromProto(msg *otg.FlowRSVPPathObjectsRecordRouteSubObjectType) (FlowRSVPPathObjectsRecordRouteSubObjectType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjectsRecordRouteSubObjectType) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsRecordRouteSubObjectType) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjectsRecordRouteSubObjectType) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsRecordRouteSubObjectType) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjectsRecordRouteSubObjectType) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsRecordRouteSubObjectType) FromJson(value string) error {
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

func (obj *flowRSVPPathObjectsRecordRouteSubObjectType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsRecordRouteSubObjectType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsRecordRouteSubObjectType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjectsRecordRouteSubObjectType) Clone() (FlowRSVPPathObjectsRecordRouteSubObjectType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjectsRecordRouteSubObjectType()
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

func (obj *flowRSVPPathObjectsRecordRouteSubObjectType) setNil() {
	obj.ipv4AddressHolder = nil
	obj.labelHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjectsRecordRouteSubObjectType is currently supported subobjects are IPv4 address(1) and Label(3).
type FlowRSVPPathObjectsRecordRouteSubObjectType interface {
	Validation
	// msg marshals FlowRSVPPathObjectsRecordRouteSubObjectType to protobuf object *otg.FlowRSVPPathObjectsRecordRouteSubObjectType
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjectsRecordRouteSubObjectType
	// setMsg unmarshals FlowRSVPPathObjectsRecordRouteSubObjectType from protobuf object *otg.FlowRSVPPathObjectsRecordRouteSubObjectType
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjectsRecordRouteSubObjectType) FlowRSVPPathObjectsRecordRouteSubObjectType
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjectsRecordRouteSubObjectType
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjectsRecordRouteSubObjectType
	// validate validates FlowRSVPPathObjectsRecordRouteSubObjectType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjectsRecordRouteSubObjectType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPPathObjectsRecordRouteSubObjectTypeChoiceEnum, set in FlowRSVPPathObjectsRecordRouteSubObjectType
	Choice() FlowRSVPPathObjectsRecordRouteSubObjectTypeChoiceEnum
	// setChoice assigns FlowRSVPPathObjectsRecordRouteSubObjectTypeChoiceEnum provided by user to FlowRSVPPathObjectsRecordRouteSubObjectType
	setChoice(value FlowRSVPPathObjectsRecordRouteSubObjectTypeChoiceEnum) FlowRSVPPathObjectsRecordRouteSubObjectType
	// HasChoice checks if Choice has been set in FlowRSVPPathObjectsRecordRouteSubObjectType
	HasChoice() bool
	// Ipv4Address returns FlowRSVPPathRecordRouteType1Ipv4Address, set in FlowRSVPPathObjectsRecordRouteSubObjectType.
	// FlowRSVPPathRecordRouteType1Ipv4Address is class = RECORD_ROUTE, Type1 ROUTE_RECORD C-Type = 1 Subobject: IPv4 Address, C-Type: 1
	Ipv4Address() FlowRSVPPathRecordRouteType1Ipv4Address
	// SetIpv4Address assigns FlowRSVPPathRecordRouteType1Ipv4Address provided by user to FlowRSVPPathObjectsRecordRouteSubObjectType.
	// FlowRSVPPathRecordRouteType1Ipv4Address is class = RECORD_ROUTE, Type1 ROUTE_RECORD C-Type = 1 Subobject: IPv4 Address, C-Type: 1
	SetIpv4Address(value FlowRSVPPathRecordRouteType1Ipv4Address) FlowRSVPPathObjectsRecordRouteSubObjectType
	// HasIpv4Address checks if Ipv4Address has been set in FlowRSVPPathObjectsRecordRouteSubObjectType
	HasIpv4Address() bool
	// Label returns FlowRSVPPathRecordRouteType1Label, set in FlowRSVPPathObjectsRecordRouteSubObjectType.
	// FlowRSVPPathRecordRouteType1Label is class = RECORD_ROUTE, Type1 ROUTE_RECORD C-Type = 1 Subobject: Label, C-Type: 3
	Label() FlowRSVPPathRecordRouteType1Label
	// SetLabel assigns FlowRSVPPathRecordRouteType1Label provided by user to FlowRSVPPathObjectsRecordRouteSubObjectType.
	// FlowRSVPPathRecordRouteType1Label is class = RECORD_ROUTE, Type1 ROUTE_RECORD C-Type = 1 Subobject: Label, C-Type: 3
	SetLabel(value FlowRSVPPathRecordRouteType1Label) FlowRSVPPathObjectsRecordRouteSubObjectType
	// HasLabel checks if Label has been set in FlowRSVPPathObjectsRecordRouteSubObjectType
	HasLabel() bool
	setNil()
}

type FlowRSVPPathObjectsRecordRouteSubObjectTypeChoiceEnum string

// Enum of Choice on FlowRSVPPathObjectsRecordRouteSubObjectType
var FlowRSVPPathObjectsRecordRouteSubObjectTypeChoice = struct {
	IPV4_ADDRESS FlowRSVPPathObjectsRecordRouteSubObjectTypeChoiceEnum
	LABEL        FlowRSVPPathObjectsRecordRouteSubObjectTypeChoiceEnum
}{
	IPV4_ADDRESS: FlowRSVPPathObjectsRecordRouteSubObjectTypeChoiceEnum("ipv4_address"),
	LABEL:        FlowRSVPPathObjectsRecordRouteSubObjectTypeChoiceEnum("label"),
}

func (obj *flowRSVPPathObjectsRecordRouteSubObjectType) Choice() FlowRSVPPathObjectsRecordRouteSubObjectTypeChoiceEnum {
	return FlowRSVPPathObjectsRecordRouteSubObjectTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowRSVPPathObjectsRecordRouteSubObjectType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPPathObjectsRecordRouteSubObjectType) setChoice(value FlowRSVPPathObjectsRecordRouteSubObjectTypeChoiceEnum) FlowRSVPPathObjectsRecordRouteSubObjectType {
	intValue, ok := otg.FlowRSVPPathObjectsRecordRouteSubObjectType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPPathObjectsRecordRouteSubObjectTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPPathObjectsRecordRouteSubObjectType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Label = nil
	obj.labelHolder = nil
	obj.obj.Ipv4Address = nil
	obj.ipv4AddressHolder = nil

	if value == FlowRSVPPathObjectsRecordRouteSubObjectTypeChoice.IPV4_ADDRESS {
		obj.obj.Ipv4Address = NewFlowRSVPPathRecordRouteType1Ipv4Address().msg()
	}

	if value == FlowRSVPPathObjectsRecordRouteSubObjectTypeChoice.LABEL {
		obj.obj.Label = NewFlowRSVPPathRecordRouteType1Label().msg()
	}

	return obj
}

// description is TBD
// Ipv4Address returns a FlowRSVPPathRecordRouteType1Ipv4Address
func (obj *flowRSVPPathObjectsRecordRouteSubObjectType) Ipv4Address() FlowRSVPPathRecordRouteType1Ipv4Address {
	if obj.obj.Ipv4Address == nil {
		obj.setChoice(FlowRSVPPathObjectsRecordRouteSubObjectTypeChoice.IPV4_ADDRESS)
	}
	if obj.ipv4AddressHolder == nil {
		obj.ipv4AddressHolder = &flowRSVPPathRecordRouteType1Ipv4Address{obj: obj.obj.Ipv4Address}
	}
	return obj.ipv4AddressHolder
}

// description is TBD
// Ipv4Address returns a FlowRSVPPathRecordRouteType1Ipv4Address
func (obj *flowRSVPPathObjectsRecordRouteSubObjectType) HasIpv4Address() bool {
	return obj.obj.Ipv4Address != nil
}

// description is TBD
// SetIpv4Address sets the FlowRSVPPathRecordRouteType1Ipv4Address value in the FlowRSVPPathObjectsRecordRouteSubObjectType object
func (obj *flowRSVPPathObjectsRecordRouteSubObjectType) SetIpv4Address(value FlowRSVPPathRecordRouteType1Ipv4Address) FlowRSVPPathObjectsRecordRouteSubObjectType {
	obj.setChoice(FlowRSVPPathObjectsRecordRouteSubObjectTypeChoice.IPV4_ADDRESS)
	obj.ipv4AddressHolder = nil
	obj.obj.Ipv4Address = value.msg()

	return obj
}

// description is TBD
// Label returns a FlowRSVPPathRecordRouteType1Label
func (obj *flowRSVPPathObjectsRecordRouteSubObjectType) Label() FlowRSVPPathRecordRouteType1Label {
	if obj.obj.Label == nil {
		obj.setChoice(FlowRSVPPathObjectsRecordRouteSubObjectTypeChoice.LABEL)
	}
	if obj.labelHolder == nil {
		obj.labelHolder = &flowRSVPPathRecordRouteType1Label{obj: obj.obj.Label}
	}
	return obj.labelHolder
}

// description is TBD
// Label returns a FlowRSVPPathRecordRouteType1Label
func (obj *flowRSVPPathObjectsRecordRouteSubObjectType) HasLabel() bool {
	return obj.obj.Label != nil
}

// description is TBD
// SetLabel sets the FlowRSVPPathRecordRouteType1Label value in the FlowRSVPPathObjectsRecordRouteSubObjectType object
func (obj *flowRSVPPathObjectsRecordRouteSubObjectType) SetLabel(value FlowRSVPPathRecordRouteType1Label) FlowRSVPPathObjectsRecordRouteSubObjectType {
	obj.setChoice(FlowRSVPPathObjectsRecordRouteSubObjectTypeChoice.LABEL)
	obj.labelHolder = nil
	obj.obj.Label = value.msg()

	return obj
}

func (obj *flowRSVPPathObjectsRecordRouteSubObjectType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv4Address != nil {

		obj.Ipv4Address().validateObj(vObj, set_default)
	}

	if obj.obj.Label != nil {

		obj.Label().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathObjectsRecordRouteSubObjectType) setDefault() {
	var choices_set int = 0
	var choice FlowRSVPPathObjectsRecordRouteSubObjectTypeChoiceEnum

	if obj.obj.Ipv4Address != nil {
		choices_set += 1
		choice = FlowRSVPPathObjectsRecordRouteSubObjectTypeChoice.IPV4_ADDRESS
	}

	if obj.obj.Label != nil {
		choices_set += 1
		choice = FlowRSVPPathObjectsRecordRouteSubObjectTypeChoice.LABEL
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowRSVPPathObjectsRecordRouteSubObjectTypeChoice.IPV4_ADDRESS)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowRSVPPathObjectsRecordRouteSubObjectType")
			}
		} else {
			intVal := otg.FlowRSVPPathObjectsRecordRouteSubObjectType_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowRSVPPathObjectsRecordRouteSubObjectType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

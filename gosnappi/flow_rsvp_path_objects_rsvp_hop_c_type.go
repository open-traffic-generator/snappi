package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjectsRsvpHopCType *****
type flowRSVPPathObjectsRsvpHopCType struct {
	validation
	obj          *otg.FlowRSVPPathObjectsRsvpHopCType
	marshaller   marshalFlowRSVPPathObjectsRsvpHopCType
	unMarshaller unMarshalFlowRSVPPathObjectsRsvpHopCType
	ipv4Holder   FlowRSVPPathRsvpHopIpv4
}

func NewFlowRSVPPathObjectsRsvpHopCType() FlowRSVPPathObjectsRsvpHopCType {
	obj := flowRSVPPathObjectsRsvpHopCType{obj: &otg.FlowRSVPPathObjectsRsvpHopCType{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjectsRsvpHopCType) msg() *otg.FlowRSVPPathObjectsRsvpHopCType {
	return obj.obj
}

func (obj *flowRSVPPathObjectsRsvpHopCType) setMsg(msg *otg.FlowRSVPPathObjectsRsvpHopCType) FlowRSVPPathObjectsRsvpHopCType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjectsRsvpHopCType struct {
	obj *flowRSVPPathObjectsRsvpHopCType
}

type marshalFlowRSVPPathObjectsRsvpHopCType interface {
	// ToProto marshals FlowRSVPPathObjectsRsvpHopCType to protobuf object *otg.FlowRSVPPathObjectsRsvpHopCType
	ToProto() (*otg.FlowRSVPPathObjectsRsvpHopCType, error)
	// ToPbText marshals FlowRSVPPathObjectsRsvpHopCType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjectsRsvpHopCType to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjectsRsvpHopCType to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathObjectsRsvpHopCType struct {
	obj *flowRSVPPathObjectsRsvpHopCType
}

type unMarshalFlowRSVPPathObjectsRsvpHopCType interface {
	// FromProto unmarshals FlowRSVPPathObjectsRsvpHopCType from protobuf object *otg.FlowRSVPPathObjectsRsvpHopCType
	FromProto(msg *otg.FlowRSVPPathObjectsRsvpHopCType) (FlowRSVPPathObjectsRsvpHopCType, error)
	// FromPbText unmarshals FlowRSVPPathObjectsRsvpHopCType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjectsRsvpHopCType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjectsRsvpHopCType from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjectsRsvpHopCType) Marshal() marshalFlowRSVPPathObjectsRsvpHopCType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjectsRsvpHopCType{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjectsRsvpHopCType) Unmarshal() unMarshalFlowRSVPPathObjectsRsvpHopCType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjectsRsvpHopCType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjectsRsvpHopCType) ToProto() (*otg.FlowRSVPPathObjectsRsvpHopCType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjectsRsvpHopCType) FromProto(msg *otg.FlowRSVPPathObjectsRsvpHopCType) (FlowRSVPPathObjectsRsvpHopCType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjectsRsvpHopCType) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsRsvpHopCType) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjectsRsvpHopCType) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsRsvpHopCType) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjectsRsvpHopCType) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsRsvpHopCType) FromJson(value string) error {
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

func (obj *flowRSVPPathObjectsRsvpHopCType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsRsvpHopCType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsRsvpHopCType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjectsRsvpHopCType) Clone() (FlowRSVPPathObjectsRsvpHopCType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjectsRsvpHopCType()
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

func (obj *flowRSVPPathObjectsRsvpHopCType) setNil() {
	obj.ipv4Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjectsRsvpHopCType is object for RSVP_HOP class. Currently supported c-type is IPv4 (1).
type FlowRSVPPathObjectsRsvpHopCType interface {
	Validation
	// msg marshals FlowRSVPPathObjectsRsvpHopCType to protobuf object *otg.FlowRSVPPathObjectsRsvpHopCType
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjectsRsvpHopCType
	// setMsg unmarshals FlowRSVPPathObjectsRsvpHopCType from protobuf object *otg.FlowRSVPPathObjectsRsvpHopCType
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjectsRsvpHopCType) FlowRSVPPathObjectsRsvpHopCType
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjectsRsvpHopCType
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjectsRsvpHopCType
	// validate validates FlowRSVPPathObjectsRsvpHopCType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjectsRsvpHopCType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPPathObjectsRsvpHopCTypeChoiceEnum, set in FlowRSVPPathObjectsRsvpHopCType
	Choice() FlowRSVPPathObjectsRsvpHopCTypeChoiceEnum
	// setChoice assigns FlowRSVPPathObjectsRsvpHopCTypeChoiceEnum provided by user to FlowRSVPPathObjectsRsvpHopCType
	setChoice(value FlowRSVPPathObjectsRsvpHopCTypeChoiceEnum) FlowRSVPPathObjectsRsvpHopCType
	// HasChoice checks if Choice has been set in FlowRSVPPathObjectsRsvpHopCType
	HasChoice() bool
	// Ipv4 returns FlowRSVPPathRsvpHopIpv4, set in FlowRSVPPathObjectsRsvpHopCType.
	// FlowRSVPPathRsvpHopIpv4 is iPv4 RSVP_HOP object: Class = 3, C-Type = 1
	Ipv4() FlowRSVPPathRsvpHopIpv4
	// SetIpv4 assigns FlowRSVPPathRsvpHopIpv4 provided by user to FlowRSVPPathObjectsRsvpHopCType.
	// FlowRSVPPathRsvpHopIpv4 is iPv4 RSVP_HOP object: Class = 3, C-Type = 1
	SetIpv4(value FlowRSVPPathRsvpHopIpv4) FlowRSVPPathObjectsRsvpHopCType
	// HasIpv4 checks if Ipv4 has been set in FlowRSVPPathObjectsRsvpHopCType
	HasIpv4() bool
	setNil()
}

type FlowRSVPPathObjectsRsvpHopCTypeChoiceEnum string

// Enum of Choice on FlowRSVPPathObjectsRsvpHopCType
var FlowRSVPPathObjectsRsvpHopCTypeChoice = struct {
	IPV4 FlowRSVPPathObjectsRsvpHopCTypeChoiceEnum
}{
	IPV4: FlowRSVPPathObjectsRsvpHopCTypeChoiceEnum("ipv4"),
}

func (obj *flowRSVPPathObjectsRsvpHopCType) Choice() FlowRSVPPathObjectsRsvpHopCTypeChoiceEnum {
	return FlowRSVPPathObjectsRsvpHopCTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowRSVPPathObjectsRsvpHopCType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPPathObjectsRsvpHopCType) setChoice(value FlowRSVPPathObjectsRsvpHopCTypeChoiceEnum) FlowRSVPPathObjectsRsvpHopCType {
	intValue, ok := otg.FlowRSVPPathObjectsRsvpHopCType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPPathObjectsRsvpHopCTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPPathObjectsRsvpHopCType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Ipv4 = nil
	obj.ipv4Holder = nil

	if value == FlowRSVPPathObjectsRsvpHopCTypeChoice.IPV4 {
		obj.obj.Ipv4 = NewFlowRSVPPathRsvpHopIpv4().msg()
	}

	return obj
}

// description is TBD
// Ipv4 returns a FlowRSVPPathRsvpHopIpv4
func (obj *flowRSVPPathObjectsRsvpHopCType) Ipv4() FlowRSVPPathRsvpHopIpv4 {
	if obj.obj.Ipv4 == nil {
		obj.setChoice(FlowRSVPPathObjectsRsvpHopCTypeChoice.IPV4)
	}
	if obj.ipv4Holder == nil {
		obj.ipv4Holder = &flowRSVPPathRsvpHopIpv4{obj: obj.obj.Ipv4}
	}
	return obj.ipv4Holder
}

// description is TBD
// Ipv4 returns a FlowRSVPPathRsvpHopIpv4
func (obj *flowRSVPPathObjectsRsvpHopCType) HasIpv4() bool {
	return obj.obj.Ipv4 != nil
}

// description is TBD
// SetIpv4 sets the FlowRSVPPathRsvpHopIpv4 value in the FlowRSVPPathObjectsRsvpHopCType object
func (obj *flowRSVPPathObjectsRsvpHopCType) SetIpv4(value FlowRSVPPathRsvpHopIpv4) FlowRSVPPathObjectsRsvpHopCType {
	obj.setChoice(FlowRSVPPathObjectsRsvpHopCTypeChoice.IPV4)
	obj.ipv4Holder = nil
	obj.obj.Ipv4 = value.msg()

	return obj
}

func (obj *flowRSVPPathObjectsRsvpHopCType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv4 != nil {

		obj.Ipv4().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathObjectsRsvpHopCType) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(FlowRSVPPathObjectsRsvpHopCTypeChoice.IPV4)

	}

}

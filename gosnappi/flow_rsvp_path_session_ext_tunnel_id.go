package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathSessionExtTunnelId *****
type flowRSVPPathSessionExtTunnelId struct {
	validation
	obj             *otg.FlowRSVPPathSessionExtTunnelId
	marshaller      marshalFlowRSVPPathSessionExtTunnelId
	unMarshaller    unMarshalFlowRSVPPathSessionExtTunnelId
	asIntegerHolder PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	asIpv4Holder    PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
}

func NewFlowRSVPPathSessionExtTunnelId() FlowRSVPPathSessionExtTunnelId {
	obj := flowRSVPPathSessionExtTunnelId{obj: &otg.FlowRSVPPathSessionExtTunnelId{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathSessionExtTunnelId) msg() *otg.FlowRSVPPathSessionExtTunnelId {
	return obj.obj
}

func (obj *flowRSVPPathSessionExtTunnelId) setMsg(msg *otg.FlowRSVPPathSessionExtTunnelId) FlowRSVPPathSessionExtTunnelId {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathSessionExtTunnelId struct {
	obj *flowRSVPPathSessionExtTunnelId
}

type marshalFlowRSVPPathSessionExtTunnelId interface {
	// ToProto marshals FlowRSVPPathSessionExtTunnelId to protobuf object *otg.FlowRSVPPathSessionExtTunnelId
	ToProto() (*otg.FlowRSVPPathSessionExtTunnelId, error)
	// ToPbText marshals FlowRSVPPathSessionExtTunnelId to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathSessionExtTunnelId to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathSessionExtTunnelId to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowRSVPPathSessionExtTunnelId to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowRSVPPathSessionExtTunnelId struct {
	obj *flowRSVPPathSessionExtTunnelId
}

type unMarshalFlowRSVPPathSessionExtTunnelId interface {
	// FromProto unmarshals FlowRSVPPathSessionExtTunnelId from protobuf object *otg.FlowRSVPPathSessionExtTunnelId
	FromProto(msg *otg.FlowRSVPPathSessionExtTunnelId) (FlowRSVPPathSessionExtTunnelId, error)
	// FromPbText unmarshals FlowRSVPPathSessionExtTunnelId from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathSessionExtTunnelId from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathSessionExtTunnelId from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathSessionExtTunnelId) Marshal() marshalFlowRSVPPathSessionExtTunnelId {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathSessionExtTunnelId{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathSessionExtTunnelId) Unmarshal() unMarshalFlowRSVPPathSessionExtTunnelId {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathSessionExtTunnelId{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathSessionExtTunnelId) ToProto() (*otg.FlowRSVPPathSessionExtTunnelId, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathSessionExtTunnelId) FromProto(msg *otg.FlowRSVPPathSessionExtTunnelId) (FlowRSVPPathSessionExtTunnelId, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathSessionExtTunnelId) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathSessionExtTunnelId) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathSessionExtTunnelId) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathSessionExtTunnelId) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathSessionExtTunnelId) ToJsonRaw() (string, error) {
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

func (m *marshalflowRSVPPathSessionExtTunnelId) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathSessionExtTunnelId) FromJson(value string) error {
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

func (obj *flowRSVPPathSessionExtTunnelId) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathSessionExtTunnelId) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathSessionExtTunnelId) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathSessionExtTunnelId) Clone() (FlowRSVPPathSessionExtTunnelId, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathSessionExtTunnelId()
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

func (obj *flowRSVPPathSessionExtTunnelId) setNil() {
	obj.asIntegerHolder = nil
	obj.asIpv4Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathSessionExtTunnelId is description is TBD
type FlowRSVPPathSessionExtTunnelId interface {
	Validation
	// msg marshals FlowRSVPPathSessionExtTunnelId to protobuf object *otg.FlowRSVPPathSessionExtTunnelId
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathSessionExtTunnelId
	// setMsg unmarshals FlowRSVPPathSessionExtTunnelId from protobuf object *otg.FlowRSVPPathSessionExtTunnelId
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathSessionExtTunnelId) FlowRSVPPathSessionExtTunnelId
	// provides marshal interface
	Marshal() marshalFlowRSVPPathSessionExtTunnelId
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathSessionExtTunnelId
	// validate validates FlowRSVPPathSessionExtTunnelId
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathSessionExtTunnelId, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPPathSessionExtTunnelIdChoiceEnum, set in FlowRSVPPathSessionExtTunnelId
	Choice() FlowRSVPPathSessionExtTunnelIdChoiceEnum
	// setChoice assigns FlowRSVPPathSessionExtTunnelIdChoiceEnum provided by user to FlowRSVPPathSessionExtTunnelId
	setChoice(value FlowRSVPPathSessionExtTunnelIdChoiceEnum) FlowRSVPPathSessionExtTunnelId
	// HasChoice checks if Choice has been set in FlowRSVPPathSessionExtTunnelId
	HasChoice() bool
	// AsInteger returns PatternFlowRSVPPathSessionExtTunnelIdAsInteger, set in FlowRSVPPathSessionExtTunnelId.
	// PatternFlowRSVPPathSessionExtTunnelIdAsInteger is tBD
	AsInteger() PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	// SetAsInteger assigns PatternFlowRSVPPathSessionExtTunnelIdAsInteger provided by user to FlowRSVPPathSessionExtTunnelId.
	// PatternFlowRSVPPathSessionExtTunnelIdAsInteger is tBD
	SetAsInteger(value PatternFlowRSVPPathSessionExtTunnelIdAsInteger) FlowRSVPPathSessionExtTunnelId
	// HasAsInteger checks if AsInteger has been set in FlowRSVPPathSessionExtTunnelId
	HasAsInteger() bool
	// AsIpv4 returns PatternFlowRSVPPathSessionExtTunnelIdAsIpv4, set in FlowRSVPPathSessionExtTunnelId.
	// PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 is iPv4 address of the ingress endpoint for the tunnel.
	AsIpv4() PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	// SetAsIpv4 assigns PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 provided by user to FlowRSVPPathSessionExtTunnelId.
	// PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 is iPv4 address of the ingress endpoint for the tunnel.
	SetAsIpv4(value PatternFlowRSVPPathSessionExtTunnelIdAsIpv4) FlowRSVPPathSessionExtTunnelId
	// HasAsIpv4 checks if AsIpv4 has been set in FlowRSVPPathSessionExtTunnelId
	HasAsIpv4() bool
	setNil()
}

type FlowRSVPPathSessionExtTunnelIdChoiceEnum string

// Enum of Choice on FlowRSVPPathSessionExtTunnelId
var FlowRSVPPathSessionExtTunnelIdChoice = struct {
	AS_INTEGER FlowRSVPPathSessionExtTunnelIdChoiceEnum
	AS_IPV4    FlowRSVPPathSessionExtTunnelIdChoiceEnum
}{
	AS_INTEGER: FlowRSVPPathSessionExtTunnelIdChoiceEnum("as_integer"),
	AS_IPV4:    FlowRSVPPathSessionExtTunnelIdChoiceEnum("as_ipv4"),
}

func (obj *flowRSVPPathSessionExtTunnelId) Choice() FlowRSVPPathSessionExtTunnelIdChoiceEnum {
	return FlowRSVPPathSessionExtTunnelIdChoiceEnum(obj.obj.Choice.Enum().String())
}

// 32 bit integer or IPv4 address.
// Choice returns a string
func (obj *flowRSVPPathSessionExtTunnelId) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPPathSessionExtTunnelId) setChoice(value FlowRSVPPathSessionExtTunnelIdChoiceEnum) FlowRSVPPathSessionExtTunnelId {
	intValue, ok := otg.FlowRSVPPathSessionExtTunnelId_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPPathSessionExtTunnelIdChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPPathSessionExtTunnelId_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.AsIpv4 = nil
	obj.asIpv4Holder = nil
	obj.obj.AsInteger = nil
	obj.asIntegerHolder = nil

	if value == FlowRSVPPathSessionExtTunnelIdChoice.AS_INTEGER {
		obj.obj.AsInteger = NewPatternFlowRSVPPathSessionExtTunnelIdAsInteger().msg()
	}

	if value == FlowRSVPPathSessionExtTunnelIdChoice.AS_IPV4 {
		obj.obj.AsIpv4 = NewPatternFlowRSVPPathSessionExtTunnelIdAsIpv4().msg()
	}

	return obj
}

// description is TBD
// AsInteger returns a PatternFlowRSVPPathSessionExtTunnelIdAsInteger
func (obj *flowRSVPPathSessionExtTunnelId) AsInteger() PatternFlowRSVPPathSessionExtTunnelIdAsInteger {
	if obj.obj.AsInteger == nil {
		obj.setChoice(FlowRSVPPathSessionExtTunnelIdChoice.AS_INTEGER)
	}
	if obj.asIntegerHolder == nil {
		obj.asIntegerHolder = &patternFlowRSVPPathSessionExtTunnelIdAsInteger{obj: obj.obj.AsInteger}
	}
	return obj.asIntegerHolder
}

// description is TBD
// AsInteger returns a PatternFlowRSVPPathSessionExtTunnelIdAsInteger
func (obj *flowRSVPPathSessionExtTunnelId) HasAsInteger() bool {
	return obj.obj.AsInteger != nil
}

// description is TBD
// SetAsInteger sets the PatternFlowRSVPPathSessionExtTunnelIdAsInteger value in the FlowRSVPPathSessionExtTunnelId object
func (obj *flowRSVPPathSessionExtTunnelId) SetAsInteger(value PatternFlowRSVPPathSessionExtTunnelIdAsInteger) FlowRSVPPathSessionExtTunnelId {
	obj.setChoice(FlowRSVPPathSessionExtTunnelIdChoice.AS_INTEGER)
	obj.asIntegerHolder = nil
	obj.obj.AsInteger = value.msg()

	return obj
}

// description is TBD
// AsIpv4 returns a PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
func (obj *flowRSVPPathSessionExtTunnelId) AsIpv4() PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 {
	if obj.obj.AsIpv4 == nil {
		obj.setChoice(FlowRSVPPathSessionExtTunnelIdChoice.AS_IPV4)
	}
	if obj.asIpv4Holder == nil {
		obj.asIpv4Holder = &patternFlowRSVPPathSessionExtTunnelIdAsIpv4{obj: obj.obj.AsIpv4}
	}
	return obj.asIpv4Holder
}

// description is TBD
// AsIpv4 returns a PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
func (obj *flowRSVPPathSessionExtTunnelId) HasAsIpv4() bool {
	return obj.obj.AsIpv4 != nil
}

// description is TBD
// SetAsIpv4 sets the PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 value in the FlowRSVPPathSessionExtTunnelId object
func (obj *flowRSVPPathSessionExtTunnelId) SetAsIpv4(value PatternFlowRSVPPathSessionExtTunnelIdAsIpv4) FlowRSVPPathSessionExtTunnelId {
	obj.setChoice(FlowRSVPPathSessionExtTunnelIdChoice.AS_IPV4)
	obj.asIpv4Holder = nil
	obj.obj.AsIpv4 = value.msg()

	return obj
}

func (obj *flowRSVPPathSessionExtTunnelId) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.AsInteger != nil {

		obj.AsInteger().validateObj(vObj, set_default)
	}

	if obj.obj.AsIpv4 != nil {

		obj.AsIpv4().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathSessionExtTunnelId) setDefault() {
	var choices_set int = 0
	var choice FlowRSVPPathSessionExtTunnelIdChoiceEnum

	if obj.obj.AsInteger != nil {
		choices_set += 1
		choice = FlowRSVPPathSessionExtTunnelIdChoice.AS_INTEGER
	}

	if obj.obj.AsIpv4 != nil {
		choices_set += 1
		choice = FlowRSVPPathSessionExtTunnelIdChoice.AS_IPV4
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowRSVPPathSessionExtTunnelIdChoice.AS_INTEGER)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowRSVPPathSessionExtTunnelId")
			}
		} else {
			intVal := otg.FlowRSVPPathSessionExtTunnelId_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowRSVPPathSessionExtTunnelId_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

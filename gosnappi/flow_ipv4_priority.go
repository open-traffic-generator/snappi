package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv4Priority *****
type flowIpv4Priority struct {
	validation
	obj          *otg.FlowIpv4Priority
	marshaller   marshalFlowIpv4Priority
	unMarshaller unMarshalFlowIpv4Priority
	rawHolder    PatternFlowIpv4PriorityRaw
	tosHolder    FlowIpv4Tos
	dscpHolder   FlowIpv4Dscp
}

func NewFlowIpv4Priority() FlowIpv4Priority {
	obj := flowIpv4Priority{obj: &otg.FlowIpv4Priority{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv4Priority) msg() *otg.FlowIpv4Priority {
	return obj.obj
}

func (obj *flowIpv4Priority) setMsg(msg *otg.FlowIpv4Priority) FlowIpv4Priority {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv4Priority struct {
	obj *flowIpv4Priority
}

type marshalFlowIpv4Priority interface {
	// ToProto marshals FlowIpv4Priority to protobuf object *otg.FlowIpv4Priority
	ToProto() (*otg.FlowIpv4Priority, error)
	// ToPbText marshals FlowIpv4Priority to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv4Priority to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv4Priority to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv4Priority struct {
	obj *flowIpv4Priority
}

type unMarshalFlowIpv4Priority interface {
	// FromProto unmarshals FlowIpv4Priority from protobuf object *otg.FlowIpv4Priority
	FromProto(msg *otg.FlowIpv4Priority) (FlowIpv4Priority, error)
	// FromPbText unmarshals FlowIpv4Priority from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv4Priority from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv4Priority from JSON text
	FromJson(value string) error
}

func (obj *flowIpv4Priority) Marshal() marshalFlowIpv4Priority {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv4Priority{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv4Priority) Unmarshal() unMarshalFlowIpv4Priority {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv4Priority{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv4Priority) ToProto() (*otg.FlowIpv4Priority, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv4Priority) FromProto(msg *otg.FlowIpv4Priority) (FlowIpv4Priority, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv4Priority) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv4Priority) FromPbText(value string) error {
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

func (m *marshalflowIpv4Priority) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv4Priority) FromYaml(value string) error {
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

func (m *marshalflowIpv4Priority) ToJson() (string, error) {
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

func (m *unMarshalflowIpv4Priority) FromJson(value string) error {
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

func (obj *flowIpv4Priority) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv4Priority) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv4Priority) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv4Priority) Clone() (FlowIpv4Priority, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv4Priority()
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

func (obj *flowIpv4Priority) setNil() {
	obj.rawHolder = nil
	obj.tosHolder = nil
	obj.dscpHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv4Priority is a container for ipv4 raw, tos, dscp ip priorities.
type FlowIpv4Priority interface {
	Validation
	// msg marshals FlowIpv4Priority to protobuf object *otg.FlowIpv4Priority
	// and doesn't set defaults
	msg() *otg.FlowIpv4Priority
	// setMsg unmarshals FlowIpv4Priority from protobuf object *otg.FlowIpv4Priority
	// and doesn't set defaults
	setMsg(*otg.FlowIpv4Priority) FlowIpv4Priority
	// provides marshal interface
	Marshal() marshalFlowIpv4Priority
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv4Priority
	// validate validates FlowIpv4Priority
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv4Priority, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowIpv4PriorityChoiceEnum, set in FlowIpv4Priority
	Choice() FlowIpv4PriorityChoiceEnum
	// setChoice assigns FlowIpv4PriorityChoiceEnum provided by user to FlowIpv4Priority
	setChoice(value FlowIpv4PriorityChoiceEnum) FlowIpv4Priority
	// HasChoice checks if Choice has been set in FlowIpv4Priority
	HasChoice() bool
	// Raw returns PatternFlowIpv4PriorityRaw, set in FlowIpv4Priority.
	// PatternFlowIpv4PriorityRaw is raw priority
	Raw() PatternFlowIpv4PriorityRaw
	// SetRaw assigns PatternFlowIpv4PriorityRaw provided by user to FlowIpv4Priority.
	// PatternFlowIpv4PriorityRaw is raw priority
	SetRaw(value PatternFlowIpv4PriorityRaw) FlowIpv4Priority
	// HasRaw checks if Raw has been set in FlowIpv4Priority
	HasRaw() bool
	// Tos returns FlowIpv4Tos, set in FlowIpv4Priority.
	// FlowIpv4Tos is type of service (TOS) packet field.
	Tos() FlowIpv4Tos
	// SetTos assigns FlowIpv4Tos provided by user to FlowIpv4Priority.
	// FlowIpv4Tos is type of service (TOS) packet field.
	SetTos(value FlowIpv4Tos) FlowIpv4Priority
	// HasTos checks if Tos has been set in FlowIpv4Priority
	HasTos() bool
	// Dscp returns FlowIpv4Dscp, set in FlowIpv4Priority.
	// FlowIpv4Dscp is differentiated services code point (DSCP) packet field.
	Dscp() FlowIpv4Dscp
	// SetDscp assigns FlowIpv4Dscp provided by user to FlowIpv4Priority.
	// FlowIpv4Dscp is differentiated services code point (DSCP) packet field.
	SetDscp(value FlowIpv4Dscp) FlowIpv4Priority
	// HasDscp checks if Dscp has been set in FlowIpv4Priority
	HasDscp() bool
	setNil()
}

type FlowIpv4PriorityChoiceEnum string

// Enum of Choice on FlowIpv4Priority
var FlowIpv4PriorityChoice = struct {
	RAW  FlowIpv4PriorityChoiceEnum
	TOS  FlowIpv4PriorityChoiceEnum
	DSCP FlowIpv4PriorityChoiceEnum
}{
	RAW:  FlowIpv4PriorityChoiceEnum("raw"),
	TOS:  FlowIpv4PriorityChoiceEnum("tos"),
	DSCP: FlowIpv4PriorityChoiceEnum("dscp"),
}

func (obj *flowIpv4Priority) Choice() FlowIpv4PriorityChoiceEnum {
	return FlowIpv4PriorityChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowIpv4Priority) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowIpv4Priority) setChoice(value FlowIpv4PriorityChoiceEnum) FlowIpv4Priority {
	intValue, ok := otg.FlowIpv4Priority_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowIpv4PriorityChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowIpv4Priority_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Dscp = nil
	obj.dscpHolder = nil
	obj.obj.Tos = nil
	obj.tosHolder = nil
	obj.obj.Raw = nil
	obj.rawHolder = nil

	if value == FlowIpv4PriorityChoice.RAW {
		obj.obj.Raw = NewPatternFlowIpv4PriorityRaw().msg()
	}

	if value == FlowIpv4PriorityChoice.TOS {
		obj.obj.Tos = NewFlowIpv4Tos().msg()
	}

	if value == FlowIpv4PriorityChoice.DSCP {
		obj.obj.Dscp = NewFlowIpv4Dscp().msg()
	}

	return obj
}

// description is TBD
// Raw returns a PatternFlowIpv4PriorityRaw
func (obj *flowIpv4Priority) Raw() PatternFlowIpv4PriorityRaw {
	if obj.obj.Raw == nil {
		obj.setChoice(FlowIpv4PriorityChoice.RAW)
	}
	if obj.rawHolder == nil {
		obj.rawHolder = &patternFlowIpv4PriorityRaw{obj: obj.obj.Raw}
	}
	return obj.rawHolder
}

// description is TBD
// Raw returns a PatternFlowIpv4PriorityRaw
func (obj *flowIpv4Priority) HasRaw() bool {
	return obj.obj.Raw != nil
}

// description is TBD
// SetRaw sets the PatternFlowIpv4PriorityRaw value in the FlowIpv4Priority object
func (obj *flowIpv4Priority) SetRaw(value PatternFlowIpv4PriorityRaw) FlowIpv4Priority {
	obj.setChoice(FlowIpv4PriorityChoice.RAW)
	obj.rawHolder = nil
	obj.obj.Raw = value.msg()

	return obj
}

// description is TBD
// Tos returns a FlowIpv4Tos
func (obj *flowIpv4Priority) Tos() FlowIpv4Tos {
	if obj.obj.Tos == nil {
		obj.setChoice(FlowIpv4PriorityChoice.TOS)
	}
	if obj.tosHolder == nil {
		obj.tosHolder = &flowIpv4Tos{obj: obj.obj.Tos}
	}
	return obj.tosHolder
}

// description is TBD
// Tos returns a FlowIpv4Tos
func (obj *flowIpv4Priority) HasTos() bool {
	return obj.obj.Tos != nil
}

// description is TBD
// SetTos sets the FlowIpv4Tos value in the FlowIpv4Priority object
func (obj *flowIpv4Priority) SetTos(value FlowIpv4Tos) FlowIpv4Priority {
	obj.setChoice(FlowIpv4PriorityChoice.TOS)
	obj.tosHolder = nil
	obj.obj.Tos = value.msg()

	return obj
}

// description is TBD
// Dscp returns a FlowIpv4Dscp
func (obj *flowIpv4Priority) Dscp() FlowIpv4Dscp {
	if obj.obj.Dscp == nil {
		obj.setChoice(FlowIpv4PriorityChoice.DSCP)
	}
	if obj.dscpHolder == nil {
		obj.dscpHolder = &flowIpv4Dscp{obj: obj.obj.Dscp}
	}
	return obj.dscpHolder
}

// description is TBD
// Dscp returns a FlowIpv4Dscp
func (obj *flowIpv4Priority) HasDscp() bool {
	return obj.obj.Dscp != nil
}

// description is TBD
// SetDscp sets the FlowIpv4Dscp value in the FlowIpv4Priority object
func (obj *flowIpv4Priority) SetDscp(value FlowIpv4Dscp) FlowIpv4Priority {
	obj.setChoice(FlowIpv4PriorityChoice.DSCP)
	obj.dscpHolder = nil
	obj.obj.Dscp = value.msg()

	return obj
}

func (obj *flowIpv4Priority) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Raw != nil {

		obj.Raw().validateObj(vObj, set_default)
	}

	if obj.obj.Tos != nil {

		obj.Tos().validateObj(vObj, set_default)
	}

	if obj.obj.Dscp != nil {

		obj.Dscp().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv4Priority) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(FlowIpv4PriorityChoice.DSCP)

	}

}

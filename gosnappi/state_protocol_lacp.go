package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocolLacp *****
type stateProtocolLacp struct {
	validation
	obj               *otg.StateProtocolLacp
	marshaller        marshalStateProtocolLacp
	unMarshaller      unMarshalStateProtocolLacp
	adminHolder       StateProtocolLacpAdmin
	memberPortsHolder StateProtocolLacpMemberPorts
}

func NewStateProtocolLacp() StateProtocolLacp {
	obj := stateProtocolLacp{obj: &otg.StateProtocolLacp{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocolLacp) msg() *otg.StateProtocolLacp {
	return obj.obj
}

func (obj *stateProtocolLacp) setMsg(msg *otg.StateProtocolLacp) StateProtocolLacp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocolLacp struct {
	obj *stateProtocolLacp
}

type marshalStateProtocolLacp interface {
	// ToProto marshals StateProtocolLacp to protobuf object *otg.StateProtocolLacp
	ToProto() (*otg.StateProtocolLacp, error)
	// ToPbText marshals StateProtocolLacp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocolLacp to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocolLacp to JSON text
	ToJson() (string, error)
}

type unMarshalstateProtocolLacp struct {
	obj *stateProtocolLacp
}

type unMarshalStateProtocolLacp interface {
	// FromProto unmarshals StateProtocolLacp from protobuf object *otg.StateProtocolLacp
	FromProto(msg *otg.StateProtocolLacp) (StateProtocolLacp, error)
	// FromPbText unmarshals StateProtocolLacp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocolLacp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocolLacp from JSON text
	FromJson(value string) error
}

func (obj *stateProtocolLacp) Marshal() marshalStateProtocolLacp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocolLacp{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocolLacp) Unmarshal() unMarshalStateProtocolLacp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocolLacp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocolLacp) ToProto() (*otg.StateProtocolLacp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocolLacp) FromProto(msg *otg.StateProtocolLacp) (StateProtocolLacp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocolLacp) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocolLacp) FromPbText(value string) error {
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

func (m *marshalstateProtocolLacp) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocolLacp) FromYaml(value string) error {
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

func (m *marshalstateProtocolLacp) ToJson() (string, error) {
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

func (m *unMarshalstateProtocolLacp) FromJson(value string) error {
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

func (obj *stateProtocolLacp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocolLacp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocolLacp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocolLacp) Clone() (StateProtocolLacp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocolLacp()
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

func (obj *stateProtocolLacp) setNil() {
	obj.adminHolder = nil
	obj.memberPortsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// StateProtocolLacp is sets state of configured LACP
type StateProtocolLacp interface {
	Validation
	// msg marshals StateProtocolLacp to protobuf object *otg.StateProtocolLacp
	// and doesn't set defaults
	msg() *otg.StateProtocolLacp
	// setMsg unmarshals StateProtocolLacp from protobuf object *otg.StateProtocolLacp
	// and doesn't set defaults
	setMsg(*otg.StateProtocolLacp) StateProtocolLacp
	// provides marshal interface
	Marshal() marshalStateProtocolLacp
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocolLacp
	// validate validates StateProtocolLacp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocolLacp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns StateProtocolLacpChoiceEnum, set in StateProtocolLacp
	Choice() StateProtocolLacpChoiceEnum
	// setChoice assigns StateProtocolLacpChoiceEnum provided by user to StateProtocolLacp
	setChoice(value StateProtocolLacpChoiceEnum) StateProtocolLacp
	// Admin returns StateProtocolLacpAdmin, set in StateProtocolLacp.
	// StateProtocolLacpAdmin is sets admin state of LACP configured on LAG members
	Admin() StateProtocolLacpAdmin
	// SetAdmin assigns StateProtocolLacpAdmin provided by user to StateProtocolLacp.
	// StateProtocolLacpAdmin is sets admin state of LACP configured on LAG members
	SetAdmin(value StateProtocolLacpAdmin) StateProtocolLacp
	// HasAdmin checks if Admin has been set in StateProtocolLacp
	HasAdmin() bool
	// MemberPorts returns StateProtocolLacpMemberPorts, set in StateProtocolLacp.
	// StateProtocolLacpMemberPorts is sets state of LACP member ports configured on LAG.
	MemberPorts() StateProtocolLacpMemberPorts
	// SetMemberPorts assigns StateProtocolLacpMemberPorts provided by user to StateProtocolLacp.
	// StateProtocolLacpMemberPorts is sets state of LACP member ports configured on LAG.
	SetMemberPorts(value StateProtocolLacpMemberPorts) StateProtocolLacp
	// HasMemberPorts checks if MemberPorts has been set in StateProtocolLacp
	HasMemberPorts() bool
	setNil()
}

type StateProtocolLacpChoiceEnum string

// Enum of Choice on StateProtocolLacp
var StateProtocolLacpChoice = struct {
	ADMIN        StateProtocolLacpChoiceEnum
	MEMBER_PORTS StateProtocolLacpChoiceEnum
}{
	ADMIN:        StateProtocolLacpChoiceEnum("admin"),
	MEMBER_PORTS: StateProtocolLacpChoiceEnum("member_ports"),
}

func (obj *stateProtocolLacp) Choice() StateProtocolLacpChoiceEnum {
	return StateProtocolLacpChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *stateProtocolLacp) setChoice(value StateProtocolLacpChoiceEnum) StateProtocolLacp {
	intValue, ok := otg.StateProtocolLacp_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolLacpChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocolLacp_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.MemberPorts = nil
	obj.memberPortsHolder = nil
	obj.obj.Admin = nil
	obj.adminHolder = nil

	if value == StateProtocolLacpChoice.ADMIN {
		obj.obj.Admin = NewStateProtocolLacpAdmin().msg()
	}

	if value == StateProtocolLacpChoice.MEMBER_PORTS {
		obj.obj.MemberPorts = NewStateProtocolLacpMemberPorts().msg()
	}

	return obj
}

// description is TBD
// Admin returns a StateProtocolLacpAdmin
func (obj *stateProtocolLacp) Admin() StateProtocolLacpAdmin {
	if obj.obj.Admin == nil {
		obj.setChoice(StateProtocolLacpChoice.ADMIN)
	}
	if obj.adminHolder == nil {
		obj.adminHolder = &stateProtocolLacpAdmin{obj: obj.obj.Admin}
	}
	return obj.adminHolder
}

// description is TBD
// Admin returns a StateProtocolLacpAdmin
func (obj *stateProtocolLacp) HasAdmin() bool {
	return obj.obj.Admin != nil
}

// description is TBD
// SetAdmin sets the StateProtocolLacpAdmin value in the StateProtocolLacp object
func (obj *stateProtocolLacp) SetAdmin(value StateProtocolLacpAdmin) StateProtocolLacp {
	obj.setChoice(StateProtocolLacpChoice.ADMIN)
	obj.adminHolder = nil
	obj.obj.Admin = value.msg()

	return obj
}

// description is TBD
// MemberPorts returns a StateProtocolLacpMemberPorts
func (obj *stateProtocolLacp) MemberPorts() StateProtocolLacpMemberPorts {
	if obj.obj.MemberPorts == nil {
		obj.setChoice(StateProtocolLacpChoice.MEMBER_PORTS)
	}
	if obj.memberPortsHolder == nil {
		obj.memberPortsHolder = &stateProtocolLacpMemberPorts{obj: obj.obj.MemberPorts}
	}
	return obj.memberPortsHolder
}

// description is TBD
// MemberPorts returns a StateProtocolLacpMemberPorts
func (obj *stateProtocolLacp) HasMemberPorts() bool {
	return obj.obj.MemberPorts != nil
}

// description is TBD
// SetMemberPorts sets the StateProtocolLacpMemberPorts value in the StateProtocolLacp object
func (obj *stateProtocolLacp) SetMemberPorts(value StateProtocolLacpMemberPorts) StateProtocolLacp {
	obj.setChoice(StateProtocolLacpChoice.MEMBER_PORTS)
	obj.memberPortsHolder = nil
	obj.obj.MemberPorts = value.msg()

	return obj
}

func (obj *stateProtocolLacp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface StateProtocolLacp")
	}

	if obj.obj.Admin != nil {

		obj.Admin().validateObj(vObj, set_default)
	}

	if obj.obj.MemberPorts != nil {

		obj.MemberPorts().validateObj(vObj, set_default)
	}

}

func (obj *stateProtocolLacp) setDefault() {

}

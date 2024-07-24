package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocolLacpMemberPorts *****
type stateProtocolLacpMemberPorts struct {
	validation
	obj          *otg.StateProtocolLacpMemberPorts
	marshaller   marshalStateProtocolLacpMemberPorts
	unMarshaller unMarshalStateProtocolLacpMemberPorts
}

func NewStateProtocolLacpMemberPorts() StateProtocolLacpMemberPorts {
	obj := stateProtocolLacpMemberPorts{obj: &otg.StateProtocolLacpMemberPorts{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocolLacpMemberPorts) msg() *otg.StateProtocolLacpMemberPorts {
	return obj.obj
}

func (obj *stateProtocolLacpMemberPorts) setMsg(msg *otg.StateProtocolLacpMemberPorts) StateProtocolLacpMemberPorts {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocolLacpMemberPorts struct {
	obj *stateProtocolLacpMemberPorts
}

type marshalStateProtocolLacpMemberPorts interface {
	// ToProto marshals StateProtocolLacpMemberPorts to protobuf object *otg.StateProtocolLacpMemberPorts
	ToProto() (*otg.StateProtocolLacpMemberPorts, error)
	// ToPbText marshals StateProtocolLacpMemberPorts to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocolLacpMemberPorts to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocolLacpMemberPorts to JSON text
	ToJson() (string, error)
}

type unMarshalstateProtocolLacpMemberPorts struct {
	obj *stateProtocolLacpMemberPorts
}

type unMarshalStateProtocolLacpMemberPorts interface {
	// FromProto unmarshals StateProtocolLacpMemberPorts from protobuf object *otg.StateProtocolLacpMemberPorts
	FromProto(msg *otg.StateProtocolLacpMemberPorts) (StateProtocolLacpMemberPorts, error)
	// FromPbText unmarshals StateProtocolLacpMemberPorts from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocolLacpMemberPorts from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocolLacpMemberPorts from JSON text
	FromJson(value string) error
}

func (obj *stateProtocolLacpMemberPorts) Marshal() marshalStateProtocolLacpMemberPorts {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocolLacpMemberPorts{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocolLacpMemberPorts) Unmarshal() unMarshalStateProtocolLacpMemberPorts {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocolLacpMemberPorts{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocolLacpMemberPorts) ToProto() (*otg.StateProtocolLacpMemberPorts, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocolLacpMemberPorts) FromProto(msg *otg.StateProtocolLacpMemberPorts) (StateProtocolLacpMemberPorts, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocolLacpMemberPorts) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocolLacpMemberPorts) FromPbText(value string) error {
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

func (m *marshalstateProtocolLacpMemberPorts) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocolLacpMemberPorts) FromYaml(value string) error {
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

func (m *marshalstateProtocolLacpMemberPorts) ToJson() (string, error) {
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

func (m *unMarshalstateProtocolLacpMemberPorts) FromJson(value string) error {
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

func (obj *stateProtocolLacpMemberPorts) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocolLacpMemberPorts) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocolLacpMemberPorts) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocolLacpMemberPorts) Clone() (StateProtocolLacpMemberPorts, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocolLacpMemberPorts()
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

// StateProtocolLacpMemberPorts is sets state of LACP member ports configured on LAG.
type StateProtocolLacpMemberPorts interface {
	Validation
	// msg marshals StateProtocolLacpMemberPorts to protobuf object *otg.StateProtocolLacpMemberPorts
	// and doesn't set defaults
	msg() *otg.StateProtocolLacpMemberPorts
	// setMsg unmarshals StateProtocolLacpMemberPorts from protobuf object *otg.StateProtocolLacpMemberPorts
	// and doesn't set defaults
	setMsg(*otg.StateProtocolLacpMemberPorts) StateProtocolLacpMemberPorts
	// provides marshal interface
	Marshal() marshalStateProtocolLacpMemberPorts
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocolLacpMemberPorts
	// validate validates StateProtocolLacpMemberPorts
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocolLacpMemberPorts, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LagMemberNames returns []string, set in StateProtocolLacpMemberPorts.
	LagMemberNames() []string
	// SetLagMemberNames assigns []string provided by user to StateProtocolLacpMemberPorts
	SetLagMemberNames(value []string) StateProtocolLacpMemberPorts
	// State returns StateProtocolLacpMemberPortsStateEnum, set in StateProtocolLacpMemberPorts
	State() StateProtocolLacpMemberPortsStateEnum
	// SetState assigns StateProtocolLacpMemberPortsStateEnum provided by user to StateProtocolLacpMemberPorts
	SetState(value StateProtocolLacpMemberPortsStateEnum) StateProtocolLacpMemberPorts
}

// The names of LAG members (ports) for which the state has to be applied. An empty or null list will control all LAG members.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// LagMemberNames returns a []string
func (obj *stateProtocolLacpMemberPorts) LagMemberNames() []string {
	if obj.obj.LagMemberNames == nil {
		obj.obj.LagMemberNames = make([]string, 0)
	}
	return obj.obj.LagMemberNames
}

// The names of LAG members (ports) for which the state has to be applied. An empty or null list will control all LAG members.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// SetLagMemberNames sets the []string value in the StateProtocolLacpMemberPorts object
func (obj *stateProtocolLacpMemberPorts) SetLagMemberNames(value []string) StateProtocolLacpMemberPorts {

	if obj.obj.LagMemberNames == nil {
		obj.obj.LagMemberNames = make([]string, 0)
	}
	obj.obj.LagMemberNames = value

	return obj
}

type StateProtocolLacpMemberPortsStateEnum string

// Enum of State on StateProtocolLacpMemberPorts
var StateProtocolLacpMemberPortsState = struct {
	UP   StateProtocolLacpMemberPortsStateEnum
	DOWN StateProtocolLacpMemberPortsStateEnum
}{
	UP:   StateProtocolLacpMemberPortsStateEnum("up"),
	DOWN: StateProtocolLacpMemberPortsStateEnum("down"),
}

func (obj *stateProtocolLacpMemberPorts) State() StateProtocolLacpMemberPortsStateEnum {
	return StateProtocolLacpMemberPortsStateEnum(obj.obj.State.Enum().String())
}

func (obj *stateProtocolLacpMemberPorts) SetState(value StateProtocolLacpMemberPortsStateEnum) StateProtocolLacpMemberPorts {
	intValue, ok := otg.StateProtocolLacpMemberPorts_State_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolLacpMemberPortsStateEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocolLacpMemberPorts_State_Enum(intValue)
	obj.obj.State = &enumValue

	return obj
}

func (obj *stateProtocolLacpMemberPorts) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// State is required
	if obj.obj.State == nil {
		vObj.validationErrors = append(vObj.validationErrors, "State is required field on interface StateProtocolLacpMemberPorts")
	}
}

func (obj *stateProtocolLacpMemberPorts) setDefault() {

}

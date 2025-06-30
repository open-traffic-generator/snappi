package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocolLacpAdmin *****
type stateProtocolLacpAdmin struct {
	validation
	obj          *otg.StateProtocolLacpAdmin
	marshaller   marshalStateProtocolLacpAdmin
	unMarshaller unMarshalStateProtocolLacpAdmin
}

func NewStateProtocolLacpAdmin() StateProtocolLacpAdmin {
	obj := stateProtocolLacpAdmin{obj: &otg.StateProtocolLacpAdmin{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocolLacpAdmin) msg() *otg.StateProtocolLacpAdmin {
	return obj.obj
}

func (obj *stateProtocolLacpAdmin) setMsg(msg *otg.StateProtocolLacpAdmin) StateProtocolLacpAdmin {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocolLacpAdmin struct {
	obj *stateProtocolLacpAdmin
}

type marshalStateProtocolLacpAdmin interface {
	// ToProto marshals StateProtocolLacpAdmin to protobuf object *otg.StateProtocolLacpAdmin
	ToProto() (*otg.StateProtocolLacpAdmin, error)
	// ToPbText marshals StateProtocolLacpAdmin to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocolLacpAdmin to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocolLacpAdmin to JSON text
	ToJson() (string, error)
}

type unMarshalstateProtocolLacpAdmin struct {
	obj *stateProtocolLacpAdmin
}

type unMarshalStateProtocolLacpAdmin interface {
	// FromProto unmarshals StateProtocolLacpAdmin from protobuf object *otg.StateProtocolLacpAdmin
	FromProto(msg *otg.StateProtocolLacpAdmin) (StateProtocolLacpAdmin, error)
	// FromPbText unmarshals StateProtocolLacpAdmin from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocolLacpAdmin from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocolLacpAdmin from JSON text
	FromJson(value string) error
}

func (obj *stateProtocolLacpAdmin) Marshal() marshalStateProtocolLacpAdmin {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocolLacpAdmin{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocolLacpAdmin) Unmarshal() unMarshalStateProtocolLacpAdmin {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocolLacpAdmin{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocolLacpAdmin) ToProto() (*otg.StateProtocolLacpAdmin, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocolLacpAdmin) FromProto(msg *otg.StateProtocolLacpAdmin) (StateProtocolLacpAdmin, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocolLacpAdmin) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocolLacpAdmin) FromPbText(value string) error {
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

func (m *marshalstateProtocolLacpAdmin) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocolLacpAdmin) FromYaml(value string) error {
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

func (m *marshalstateProtocolLacpAdmin) ToJson() (string, error) {
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

func (m *unMarshalstateProtocolLacpAdmin) FromJson(value string) error {
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

func (obj *stateProtocolLacpAdmin) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocolLacpAdmin) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocolLacpAdmin) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocolLacpAdmin) Clone() (StateProtocolLacpAdmin, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocolLacpAdmin()
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

// StateProtocolLacpAdmin is sets admin state of LACP configured on LAG members
type StateProtocolLacpAdmin interface {
	Validation
	// msg marshals StateProtocolLacpAdmin to protobuf object *otg.StateProtocolLacpAdmin
	// and doesn't set defaults
	msg() *otg.StateProtocolLacpAdmin
	// setMsg unmarshals StateProtocolLacpAdmin from protobuf object *otg.StateProtocolLacpAdmin
	// and doesn't set defaults
	setMsg(*otg.StateProtocolLacpAdmin) StateProtocolLacpAdmin
	// provides marshal interface
	Marshal() marshalStateProtocolLacpAdmin
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocolLacpAdmin
	// validate validates StateProtocolLacpAdmin
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocolLacpAdmin, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LagMemberNames returns []string, set in StateProtocolLacpAdmin.
	LagMemberNames() []string
	// SetLagMemberNames assigns []string provided by user to StateProtocolLacpAdmin
	SetLagMemberNames(value []string) StateProtocolLacpAdmin
	// State returns StateProtocolLacpAdminStateEnum, set in StateProtocolLacpAdmin
	State() StateProtocolLacpAdminStateEnum
	// SetState assigns StateProtocolLacpAdminStateEnum provided by user to StateProtocolLacpAdmin
	SetState(value StateProtocolLacpAdminStateEnum) StateProtocolLacpAdmin
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
func (obj *stateProtocolLacpAdmin) LagMemberNames() []string {
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
// SetLagMemberNames sets the []string value in the StateProtocolLacpAdmin object
func (obj *stateProtocolLacpAdmin) SetLagMemberNames(value []string) StateProtocolLacpAdmin {

	if obj.obj.LagMemberNames == nil {
		obj.obj.LagMemberNames = make([]string, 0)
	}
	obj.obj.LagMemberNames = value

	return obj
}

type StateProtocolLacpAdminStateEnum string

// Enum of State on StateProtocolLacpAdmin
var StateProtocolLacpAdminState = struct {
	UP   StateProtocolLacpAdminStateEnum
	DOWN StateProtocolLacpAdminStateEnum
}{
	UP:   StateProtocolLacpAdminStateEnum("up"),
	DOWN: StateProtocolLacpAdminStateEnum("down"),
}

func (obj *stateProtocolLacpAdmin) State() StateProtocolLacpAdminStateEnum {
	return StateProtocolLacpAdminStateEnum(obj.obj.State.Enum().String())
}

func (obj *stateProtocolLacpAdmin) SetState(value StateProtocolLacpAdminStateEnum) StateProtocolLacpAdmin {
	intValue, ok := otg.StateProtocolLacpAdmin_State_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolLacpAdminStateEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocolLacpAdmin_State_Enum(intValue)
	obj.obj.State = &enumValue

	return obj
}

func (obj *stateProtocolLacpAdmin) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// State is required
	if obj.obj.State == nil {
		vObj.validationErrors = append(vObj.validationErrors, "State is required field on interface StateProtocolLacpAdmin")
	}
}

func (obj *stateProtocolLacpAdmin) setDefault() {

}

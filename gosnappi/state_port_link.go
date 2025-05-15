package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StatePortLink *****
type statePortLink struct {
	validation
	obj          *otg.StatePortLink
	marshaller   marshalStatePortLink
	unMarshaller unMarshalStatePortLink
}

func NewStatePortLink() StatePortLink {
	obj := statePortLink{obj: &otg.StatePortLink{}}
	obj.setDefault()
	return &obj
}

func (obj *statePortLink) msg() *otg.StatePortLink {
	return obj.obj
}

func (obj *statePortLink) setMsg(msg *otg.StatePortLink) StatePortLink {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstatePortLink struct {
	obj *statePortLink
}

type marshalStatePortLink interface {
	// ToProto marshals StatePortLink to protobuf object *otg.StatePortLink
	ToProto() (*otg.StatePortLink, error)
	// ToPbText marshals StatePortLink to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StatePortLink to YAML text
	ToYaml() (string, error)
	// ToJson marshals StatePortLink to JSON text
	ToJson() (string, error)
}

type unMarshalstatePortLink struct {
	obj *statePortLink
}

type unMarshalStatePortLink interface {
	// FromProto unmarshals StatePortLink from protobuf object *otg.StatePortLink
	FromProto(msg *otg.StatePortLink) (StatePortLink, error)
	// FromPbText unmarshals StatePortLink from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StatePortLink from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StatePortLink from JSON text
	FromJson(value string) error
}

func (obj *statePortLink) Marshal() marshalStatePortLink {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstatePortLink{obj: obj}
	}
	return obj.marshaller
}

func (obj *statePortLink) Unmarshal() unMarshalStatePortLink {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstatePortLink{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstatePortLink) ToProto() (*otg.StatePortLink, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstatePortLink) FromProto(msg *otg.StatePortLink) (StatePortLink, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstatePortLink) ToPbText() (string, error) {
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

func (m *unMarshalstatePortLink) FromPbText(value string) error {
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

func (m *marshalstatePortLink) ToYaml() (string, error) {
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

func (m *unMarshalstatePortLink) FromYaml(value string) error {
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

func (m *marshalstatePortLink) ToJson() (string, error) {
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

func (m *unMarshalstatePortLink) FromJson(value string) error {
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

func (obj *statePortLink) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *statePortLink) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *statePortLink) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *statePortLink) Clone() (StatePortLink, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStatePortLink()
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

// StatePortLink is sets the link of configured ports.
type StatePortLink interface {
	Validation
	// msg marshals StatePortLink to protobuf object *otg.StatePortLink
	// and doesn't set defaults
	msg() *otg.StatePortLink
	// setMsg unmarshals StatePortLink from protobuf object *otg.StatePortLink
	// and doesn't set defaults
	setMsg(*otg.StatePortLink) StatePortLink
	// provides marshal interface
	Marshal() marshalStatePortLink
	// provides unmarshal interface
	Unmarshal() unMarshalStatePortLink
	// validate validates StatePortLink
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StatePortLink, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PortNames returns []string, set in StatePortLink.
	PortNames() []string
	// SetPortNames assigns []string provided by user to StatePortLink
	SetPortNames(value []string) StatePortLink
	// State returns StatePortLinkStateEnum, set in StatePortLink
	State() StatePortLinkStateEnum
	// SetState assigns StatePortLinkStateEnum provided by user to StatePortLink
	SetState(value StatePortLinkStateEnum) StatePortLink
}

// The names of target ports. An empty or null list will target all ports.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// PortNames returns a []string
func (obj *statePortLink) PortNames() []string {
	if obj.obj.PortNames == nil {
		obj.obj.PortNames = make([]string, 0)
	}
	return obj.obj.PortNames
}

// The names of target ports. An empty or null list will target all ports.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// SetPortNames sets the []string value in the StatePortLink object
func (obj *statePortLink) SetPortNames(value []string) StatePortLink {

	if obj.obj.PortNames == nil {
		obj.obj.PortNames = make([]string, 0)
	}
	obj.obj.PortNames = value

	return obj
}

type StatePortLinkStateEnum string

// Enum of State on StatePortLink
var StatePortLinkState = struct {
	UP   StatePortLinkStateEnum
	DOWN StatePortLinkStateEnum
}{
	UP:   StatePortLinkStateEnum("up"),
	DOWN: StatePortLinkStateEnum("down"),
}

func (obj *statePortLink) State() StatePortLinkStateEnum {
	return StatePortLinkStateEnum(obj.obj.State.Enum().String())
}

func (obj *statePortLink) SetState(value StatePortLinkStateEnum) StatePortLink {
	intValue, ok := otg.StatePortLink_State_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StatePortLinkStateEnum", string(value)))
		return obj
	}
	enumValue := otg.StatePortLink_State_Enum(intValue)
	obj.obj.State = &enumValue

	return obj
}

func (obj *statePortLink) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// State is required
	if obj.obj.State == nil {
		vObj.validationErrors = append(vObj.validationErrors, "State is required field on interface StatePortLink")
	}
}

func (obj *statePortLink) setDefault() {

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocolOspfv3Routers *****
type stateProtocolOspfv3Routers struct {
	validation
	obj          *otg.StateProtocolOspfv3Routers
	marshaller   marshalStateProtocolOspfv3Routers
	unMarshaller unMarshalStateProtocolOspfv3Routers
}

func NewStateProtocolOspfv3Routers() StateProtocolOspfv3Routers {
	obj := stateProtocolOspfv3Routers{obj: &otg.StateProtocolOspfv3Routers{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocolOspfv3Routers) msg() *otg.StateProtocolOspfv3Routers {
	return obj.obj
}

func (obj *stateProtocolOspfv3Routers) setMsg(msg *otg.StateProtocolOspfv3Routers) StateProtocolOspfv3Routers {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocolOspfv3Routers struct {
	obj *stateProtocolOspfv3Routers
}

type marshalStateProtocolOspfv3Routers interface {
	// ToProto marshals StateProtocolOspfv3Routers to protobuf object *otg.StateProtocolOspfv3Routers
	ToProto() (*otg.StateProtocolOspfv3Routers, error)
	// ToPbText marshals StateProtocolOspfv3Routers to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocolOspfv3Routers to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocolOspfv3Routers to JSON text
	ToJson() (string, error)
}

type unMarshalstateProtocolOspfv3Routers struct {
	obj *stateProtocolOspfv3Routers
}

type unMarshalStateProtocolOspfv3Routers interface {
	// FromProto unmarshals StateProtocolOspfv3Routers from protobuf object *otg.StateProtocolOspfv3Routers
	FromProto(msg *otg.StateProtocolOspfv3Routers) (StateProtocolOspfv3Routers, error)
	// FromPbText unmarshals StateProtocolOspfv3Routers from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocolOspfv3Routers from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocolOspfv3Routers from JSON text
	FromJson(value string) error
}

func (obj *stateProtocolOspfv3Routers) Marshal() marshalStateProtocolOspfv3Routers {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocolOspfv3Routers{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocolOspfv3Routers) Unmarshal() unMarshalStateProtocolOspfv3Routers {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocolOspfv3Routers{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocolOspfv3Routers) ToProto() (*otg.StateProtocolOspfv3Routers, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocolOspfv3Routers) FromProto(msg *otg.StateProtocolOspfv3Routers) (StateProtocolOspfv3Routers, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocolOspfv3Routers) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocolOspfv3Routers) FromPbText(value string) error {
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

func (m *marshalstateProtocolOspfv3Routers) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocolOspfv3Routers) FromYaml(value string) error {
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

func (m *marshalstateProtocolOspfv3Routers) ToJson() (string, error) {
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

func (m *unMarshalstateProtocolOspfv3Routers) FromJson(value string) error {
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

func (obj *stateProtocolOspfv3Routers) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocolOspfv3Routers) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocolOspfv3Routers) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocolOspfv3Routers) Clone() (StateProtocolOspfv3Routers, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocolOspfv3Routers()
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

// StateProtocolOspfv3Routers is sets state of configured OSPFv3 routers.
type StateProtocolOspfv3Routers interface {
	Validation
	// msg marshals StateProtocolOspfv3Routers to protobuf object *otg.StateProtocolOspfv3Routers
	// and doesn't set defaults
	msg() *otg.StateProtocolOspfv3Routers
	// setMsg unmarshals StateProtocolOspfv3Routers from protobuf object *otg.StateProtocolOspfv3Routers
	// and doesn't set defaults
	setMsg(*otg.StateProtocolOspfv3Routers) StateProtocolOspfv3Routers
	// provides marshal interface
	Marshal() marshalStateProtocolOspfv3Routers
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocolOspfv3Routers
	// validate validates StateProtocolOspfv3Routers
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocolOspfv3Routers, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterNames returns []string, set in StateProtocolOspfv3Routers.
	RouterNames() []string
	// SetRouterNames assigns []string provided by user to StateProtocolOspfv3Routers
	SetRouterNames(value []string) StateProtocolOspfv3Routers
	// State returns StateProtocolOspfv3RoutersStateEnum, set in StateProtocolOspfv3Routers
	State() StateProtocolOspfv3RoutersStateEnum
	// SetState assigns StateProtocolOspfv3RoutersStateEnum provided by user to StateProtocolOspfv3Routers
	SetState(value StateProtocolOspfv3RoutersStateEnum) StateProtocolOspfv3Routers
}

// The names of OSPFv3 routers for which the state has to be applied. An empty or null list will control all OSPFv3 routers.
//
// x-constraint:
// - /components/schemas/Device.Ospfv3/properties/name
//
// RouterNames returns a []string
func (obj *stateProtocolOspfv3Routers) RouterNames() []string {
	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	return obj.obj.RouterNames
}

// The names of OSPFv3 routers for which the state has to be applied. An empty or null list will control all OSPFv3 routers.
//
// x-constraint:
// - /components/schemas/Device.Ospfv3/properties/name
//
// SetRouterNames sets the []string value in the StateProtocolOspfv3Routers object
func (obj *stateProtocolOspfv3Routers) SetRouterNames(value []string) StateProtocolOspfv3Routers {

	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	obj.obj.RouterNames = value

	return obj
}

type StateProtocolOspfv3RoutersStateEnum string

// Enum of State on StateProtocolOspfv3Routers
var StateProtocolOspfv3RoutersState = struct {
	UP   StateProtocolOspfv3RoutersStateEnum
	DOWN StateProtocolOspfv3RoutersStateEnum
}{
	UP:   StateProtocolOspfv3RoutersStateEnum("up"),
	DOWN: StateProtocolOspfv3RoutersStateEnum("down"),
}

func (obj *stateProtocolOspfv3Routers) State() StateProtocolOspfv3RoutersStateEnum {
	return StateProtocolOspfv3RoutersStateEnum(obj.obj.State.Enum().String())
}

func (obj *stateProtocolOspfv3Routers) SetState(value StateProtocolOspfv3RoutersStateEnum) StateProtocolOspfv3Routers {
	intValue, ok := otg.StateProtocolOspfv3Routers_State_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolOspfv3RoutersStateEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocolOspfv3Routers_State_Enum(intValue)
	obj.obj.State = &enumValue

	return obj
}

func (obj *stateProtocolOspfv3Routers) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// State is required
	if obj.obj.State == nil {
		vObj.validationErrors = append(vObj.validationErrors, "State is required field on interface StateProtocolOspfv3Routers")
	}
}

func (obj *stateProtocolOspfv3Routers) setDefault() {

}

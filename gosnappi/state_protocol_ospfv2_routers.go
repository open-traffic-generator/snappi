package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocolOspfv2Routers *****
type stateProtocolOspfv2Routers struct {
	validation
	obj          *otg.StateProtocolOspfv2Routers
	marshaller   marshalStateProtocolOspfv2Routers
	unMarshaller unMarshalStateProtocolOspfv2Routers
}

func NewStateProtocolOspfv2Routers() StateProtocolOspfv2Routers {
	obj := stateProtocolOspfv2Routers{obj: &otg.StateProtocolOspfv2Routers{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocolOspfv2Routers) msg() *otg.StateProtocolOspfv2Routers {
	return obj.obj
}

func (obj *stateProtocolOspfv2Routers) setMsg(msg *otg.StateProtocolOspfv2Routers) StateProtocolOspfv2Routers {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocolOspfv2Routers struct {
	obj *stateProtocolOspfv2Routers
}

type marshalStateProtocolOspfv2Routers interface {
	// ToProto marshals StateProtocolOspfv2Routers to protobuf object *otg.StateProtocolOspfv2Routers
	ToProto() (*otg.StateProtocolOspfv2Routers, error)
	// ToPbText marshals StateProtocolOspfv2Routers to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocolOspfv2Routers to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocolOspfv2Routers to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals StateProtocolOspfv2Routers to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalstateProtocolOspfv2Routers struct {
	obj *stateProtocolOspfv2Routers
}

type unMarshalStateProtocolOspfv2Routers interface {
	// FromProto unmarshals StateProtocolOspfv2Routers from protobuf object *otg.StateProtocolOspfv2Routers
	FromProto(msg *otg.StateProtocolOspfv2Routers) (StateProtocolOspfv2Routers, error)
	// FromPbText unmarshals StateProtocolOspfv2Routers from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocolOspfv2Routers from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocolOspfv2Routers from JSON text
	FromJson(value string) error
}

func (obj *stateProtocolOspfv2Routers) Marshal() marshalStateProtocolOspfv2Routers {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocolOspfv2Routers{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocolOspfv2Routers) Unmarshal() unMarshalStateProtocolOspfv2Routers {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocolOspfv2Routers{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocolOspfv2Routers) ToProto() (*otg.StateProtocolOspfv2Routers, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocolOspfv2Routers) FromProto(msg *otg.StateProtocolOspfv2Routers) (StateProtocolOspfv2Routers, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocolOspfv2Routers) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocolOspfv2Routers) FromPbText(value string) error {
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

func (m *marshalstateProtocolOspfv2Routers) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocolOspfv2Routers) FromYaml(value string) error {
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

func (m *marshalstateProtocolOspfv2Routers) ToJsonRaw() (string, error) {
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

func (m *marshalstateProtocolOspfv2Routers) ToJson() (string, error) {
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

func (m *unMarshalstateProtocolOspfv2Routers) FromJson(value string) error {
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

func (obj *stateProtocolOspfv2Routers) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocolOspfv2Routers) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocolOspfv2Routers) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocolOspfv2Routers) Clone() (StateProtocolOspfv2Routers, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocolOspfv2Routers()
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

// StateProtocolOspfv2Routers is sets state of configured OSPFv2 routers.
type StateProtocolOspfv2Routers interface {
	Validation
	// msg marshals StateProtocolOspfv2Routers to protobuf object *otg.StateProtocolOspfv2Routers
	// and doesn't set defaults
	msg() *otg.StateProtocolOspfv2Routers
	// setMsg unmarshals StateProtocolOspfv2Routers from protobuf object *otg.StateProtocolOspfv2Routers
	// and doesn't set defaults
	setMsg(*otg.StateProtocolOspfv2Routers) StateProtocolOspfv2Routers
	// provides marshal interface
	Marshal() marshalStateProtocolOspfv2Routers
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocolOspfv2Routers
	// validate validates StateProtocolOspfv2Routers
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocolOspfv2Routers, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterNames returns []string, set in StateProtocolOspfv2Routers.
	RouterNames() []string
	// SetRouterNames assigns []string provided by user to StateProtocolOspfv2Routers
	SetRouterNames(value []string) StateProtocolOspfv2Routers
	// State returns StateProtocolOspfv2RoutersStateEnum, set in StateProtocolOspfv2Routers
	State() StateProtocolOspfv2RoutersStateEnum
	// SetState assigns StateProtocolOspfv2RoutersStateEnum provided by user to StateProtocolOspfv2Routers
	SetState(value StateProtocolOspfv2RoutersStateEnum) StateProtocolOspfv2Routers
}

// The names of OSPFv2 routers for which the state has to be applied. An empty or null list will control all OSPFv2 routers.
//
// x-constraint:
// - /components/schemas/Device.Ospfv2/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ospfv2/properties/name
//
// RouterNames returns a []string
func (obj *stateProtocolOspfv2Routers) RouterNames() []string {
	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	return obj.obj.RouterNames
}

// The names of OSPFv2 routers for which the state has to be applied. An empty or null list will control all OSPFv2 routers.
//
// x-constraint:
// - /components/schemas/Device.Ospfv2/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ospfv2/properties/name
//
// SetRouterNames sets the []string value in the StateProtocolOspfv2Routers object
func (obj *stateProtocolOspfv2Routers) SetRouterNames(value []string) StateProtocolOspfv2Routers {

	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	obj.obj.RouterNames = value

	return obj
}

type StateProtocolOspfv2RoutersStateEnum string

// Enum of State on StateProtocolOspfv2Routers
var StateProtocolOspfv2RoutersState = struct {
	UP   StateProtocolOspfv2RoutersStateEnum
	DOWN StateProtocolOspfv2RoutersStateEnum
}{
	UP:   StateProtocolOspfv2RoutersStateEnum("up"),
	DOWN: StateProtocolOspfv2RoutersStateEnum("down"),
}

func (obj *stateProtocolOspfv2Routers) State() StateProtocolOspfv2RoutersStateEnum {
	return StateProtocolOspfv2RoutersStateEnum(obj.obj.State.Enum().String())
}

func (obj *stateProtocolOspfv2Routers) SetState(value StateProtocolOspfv2RoutersStateEnum) StateProtocolOspfv2Routers {
	intValue, ok := otg.StateProtocolOspfv2Routers_State_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolOspfv2RoutersStateEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocolOspfv2Routers_State_Enum(intValue)
	obj.obj.State = &enumValue

	return obj
}

func (obj *stateProtocolOspfv2Routers) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// State is required
	if obj.obj.State == nil {
		vObj.validationErrors = append(vObj.validationErrors, "State is required field on interface StateProtocolOspfv2Routers")
	}
}

func (obj *stateProtocolOspfv2Routers) setDefault() {

}

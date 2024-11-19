package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocolRoute *****
type stateProtocolRoute struct {
	validation
	obj          *otg.StateProtocolRoute
	marshaller   marshalStateProtocolRoute
	unMarshaller unMarshalStateProtocolRoute
}

func NewStateProtocolRoute() StateProtocolRoute {
	obj := stateProtocolRoute{obj: &otg.StateProtocolRoute{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocolRoute) msg() *otg.StateProtocolRoute {
	return obj.obj
}

func (obj *stateProtocolRoute) setMsg(msg *otg.StateProtocolRoute) StateProtocolRoute {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocolRoute struct {
	obj *stateProtocolRoute
}

type marshalStateProtocolRoute interface {
	// ToProto marshals StateProtocolRoute to protobuf object *otg.StateProtocolRoute
	ToProto() (*otg.StateProtocolRoute, error)
	// ToPbText marshals StateProtocolRoute to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocolRoute to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocolRoute to JSON text
	ToJson() (string, error)
}

type unMarshalstateProtocolRoute struct {
	obj *stateProtocolRoute
}

type unMarshalStateProtocolRoute interface {
	// FromProto unmarshals StateProtocolRoute from protobuf object *otg.StateProtocolRoute
	FromProto(msg *otg.StateProtocolRoute) (StateProtocolRoute, error)
	// FromPbText unmarshals StateProtocolRoute from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocolRoute from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocolRoute from JSON text
	FromJson(value string) error
}

func (obj *stateProtocolRoute) Marshal() marshalStateProtocolRoute {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocolRoute{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocolRoute) Unmarshal() unMarshalStateProtocolRoute {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocolRoute{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocolRoute) ToProto() (*otg.StateProtocolRoute, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocolRoute) FromProto(msg *otg.StateProtocolRoute) (StateProtocolRoute, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocolRoute) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocolRoute) FromPbText(value string) error {
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

func (m *marshalstateProtocolRoute) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocolRoute) FromYaml(value string) error {
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

func (m *marshalstateProtocolRoute) ToJson() (string, error) {
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

func (m *unMarshalstateProtocolRoute) FromJson(value string) error {
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

func (obj *stateProtocolRoute) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocolRoute) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocolRoute) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocolRoute) Clone() (StateProtocolRoute, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocolRoute()
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

// StateProtocolRoute is sets the state of configured routes
type StateProtocolRoute interface {
	Validation
	// msg marshals StateProtocolRoute to protobuf object *otg.StateProtocolRoute
	// and doesn't set defaults
	msg() *otg.StateProtocolRoute
	// setMsg unmarshals StateProtocolRoute from protobuf object *otg.StateProtocolRoute
	// and doesn't set defaults
	setMsg(*otg.StateProtocolRoute) StateProtocolRoute
	// provides marshal interface
	Marshal() marshalStateProtocolRoute
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocolRoute
	// validate validates StateProtocolRoute
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocolRoute, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Names returns []string, set in StateProtocolRoute.
	Names() []string
	// SetNames assigns []string provided by user to StateProtocolRoute
	SetNames(value []string) StateProtocolRoute
	// State returns StateProtocolRouteStateEnum, set in StateProtocolRoute
	State() StateProtocolRouteStateEnum
	// SetState assigns StateProtocolRouteStateEnum provided by user to StateProtocolRoute
	SetState(value StateProtocolRouteStateEnum) StateProtocolRoute
}

// The names of device route objects to control. If no names are specified then all route objects that match the x-constraint will be affected.
//
// x-constraint:
// - /components/schemas/Bgp.V4RouteRange/properties/name
// - /components/schemas/Bgp.V6RouteRange/properties/name
// - /components/schemas/Isis.V4RouteRange/properties/name
// - /components/schemas/Isis.V6RouteRange/properties/name
// - /components/schemas/Ospfv2.V4RouteRange/properties/name
//
// Names returns a []string
func (obj *stateProtocolRoute) Names() []string {
	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	return obj.obj.Names
}

// The names of device route objects to control. If no names are specified then all route objects that match the x-constraint will be affected.
//
// x-constraint:
// - /components/schemas/Bgp.V4RouteRange/properties/name
// - /components/schemas/Bgp.V6RouteRange/properties/name
// - /components/schemas/Isis.V4RouteRange/properties/name
// - /components/schemas/Isis.V6RouteRange/properties/name
// - /components/schemas/Ospfv2.V4RouteRange/properties/name
//
// SetNames sets the []string value in the StateProtocolRoute object
func (obj *stateProtocolRoute) SetNames(value []string) StateProtocolRoute {

	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	obj.obj.Names = value

	return obj
}

type StateProtocolRouteStateEnum string

// Enum of State on StateProtocolRoute
var StateProtocolRouteState = struct {
	WITHDRAW  StateProtocolRouteStateEnum
	ADVERTISE StateProtocolRouteStateEnum
}{
	WITHDRAW:  StateProtocolRouteStateEnum("withdraw"),
	ADVERTISE: StateProtocolRouteStateEnum("advertise"),
}

func (obj *stateProtocolRoute) State() StateProtocolRouteStateEnum {
	return StateProtocolRouteStateEnum(obj.obj.State.Enum().String())
}

func (obj *stateProtocolRoute) SetState(value StateProtocolRouteStateEnum) StateProtocolRoute {
	intValue, ok := otg.StateProtocolRoute_State_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolRouteStateEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocolRoute_State_Enum(intValue)
	obj.obj.State = &enumValue

	return obj
}

func (obj *stateProtocolRoute) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// State is required
	if obj.obj.State == nil {
		vObj.validationErrors = append(vObj.validationErrors, "State is required field on interface StateProtocolRoute")
	}
}

func (obj *stateProtocolRoute) setDefault() {

}

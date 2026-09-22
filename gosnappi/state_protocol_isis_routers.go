package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocolIsisRouters *****
type stateProtocolIsisRouters struct {
	validation
	obj          *otg.StateProtocolIsisRouters
	marshaller   marshalStateProtocolIsisRouters
	unMarshaller unMarshalStateProtocolIsisRouters
}

func NewStateProtocolIsisRouters() StateProtocolIsisRouters {
	obj := stateProtocolIsisRouters{obj: &otg.StateProtocolIsisRouters{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocolIsisRouters) msg() *otg.StateProtocolIsisRouters {
	return obj.obj
}

func (obj *stateProtocolIsisRouters) setMsg(msg *otg.StateProtocolIsisRouters) StateProtocolIsisRouters {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocolIsisRouters struct {
	obj *stateProtocolIsisRouters
}

type marshalStateProtocolIsisRouters interface {
	// ToProto marshals StateProtocolIsisRouters to protobuf object *otg.StateProtocolIsisRouters
	ToProto() (*otg.StateProtocolIsisRouters, error)
	// ToPbText marshals StateProtocolIsisRouters to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocolIsisRouters to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocolIsisRouters to JSON text
	ToJson() (string, error)
}

type unMarshalstateProtocolIsisRouters struct {
	obj *stateProtocolIsisRouters
}

type unMarshalStateProtocolIsisRouters interface {
	// FromProto unmarshals StateProtocolIsisRouters from protobuf object *otg.StateProtocolIsisRouters
	FromProto(msg *otg.StateProtocolIsisRouters) (StateProtocolIsisRouters, error)
	// FromPbText unmarshals StateProtocolIsisRouters from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocolIsisRouters from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocolIsisRouters from JSON text
	FromJson(value string) error
}

func (obj *stateProtocolIsisRouters) Marshal() marshalStateProtocolIsisRouters {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocolIsisRouters{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocolIsisRouters) Unmarshal() unMarshalStateProtocolIsisRouters {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocolIsisRouters{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocolIsisRouters) ToProto() (*otg.StateProtocolIsisRouters, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocolIsisRouters) FromProto(msg *otg.StateProtocolIsisRouters) (StateProtocolIsisRouters, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocolIsisRouters) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocolIsisRouters) FromPbText(value string) error {
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

func (m *marshalstateProtocolIsisRouters) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocolIsisRouters) FromYaml(value string) error {
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

func (m *marshalstateProtocolIsisRouters) ToJson() (string, error) {
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

func (m *unMarshalstateProtocolIsisRouters) FromJson(value string) error {
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

func (obj *stateProtocolIsisRouters) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocolIsisRouters) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocolIsisRouters) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocolIsisRouters) Clone() (StateProtocolIsisRouters, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocolIsisRouters()
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

// StateProtocolIsisRouters is sets state of configured ISIS routers.
type StateProtocolIsisRouters interface {
	Validation
	// msg marshals StateProtocolIsisRouters to protobuf object *otg.StateProtocolIsisRouters
	// and doesn't set defaults
	msg() *otg.StateProtocolIsisRouters
	// setMsg unmarshals StateProtocolIsisRouters from protobuf object *otg.StateProtocolIsisRouters
	// and doesn't set defaults
	setMsg(*otg.StateProtocolIsisRouters) StateProtocolIsisRouters
	// provides marshal interface
	Marshal() marshalStateProtocolIsisRouters
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocolIsisRouters
	// validate validates StateProtocolIsisRouters
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocolIsisRouters, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterNames returns []string, set in StateProtocolIsisRouters.
	RouterNames() []string
	// SetRouterNames assigns []string provided by user to StateProtocolIsisRouters
	SetRouterNames(value []string) StateProtocolIsisRouters
	// State returns StateProtocolIsisRoutersStateEnum, set in StateProtocolIsisRouters
	State() StateProtocolIsisRoutersStateEnum
	// SetState assigns StateProtocolIsisRoutersStateEnum provided by user to StateProtocolIsisRouters
	SetState(value StateProtocolIsisRoutersStateEnum) StateProtocolIsisRouters
}

// The names of ISIS routers for which the state has to be applied. An empty or null list will control all ISIS routers.
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// RouterNames returns a []string
func (obj *stateProtocolIsisRouters) RouterNames() []string {
	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	return obj.obj.RouterNames
}

// The names of ISIS routers for which the state has to be applied. An empty or null list will control all ISIS routers.
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// SetRouterNames sets the []string value in the StateProtocolIsisRouters object
func (obj *stateProtocolIsisRouters) SetRouterNames(value []string) StateProtocolIsisRouters {

	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	obj.obj.RouterNames = value

	return obj
}

type StateProtocolIsisRoutersStateEnum string

// Enum of State on StateProtocolIsisRouters
var StateProtocolIsisRoutersState = struct {
	UP   StateProtocolIsisRoutersStateEnum
	DOWN StateProtocolIsisRoutersStateEnum
}{
	UP:   StateProtocolIsisRoutersStateEnum("up"),
	DOWN: StateProtocolIsisRoutersStateEnum("down"),
}

func (obj *stateProtocolIsisRouters) State() StateProtocolIsisRoutersStateEnum {
	return StateProtocolIsisRoutersStateEnum(obj.obj.State.Enum().String())
}

func (obj *stateProtocolIsisRouters) SetState(value StateProtocolIsisRoutersStateEnum) StateProtocolIsisRouters {
	intValue, ok := otg.StateProtocolIsisRouters_State_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolIsisRoutersStateEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocolIsisRouters_State_Enum(intValue)
	obj.obj.State = &enumValue

	return obj
}

func (obj *stateProtocolIsisRouters) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// State is required
	if obj.obj.State == nil {
		vObj.validationErrors = append(vObj.validationErrors, "State is required field on interface StateProtocolIsisRouters")
	}
}

func (obj *stateProtocolIsisRouters) setDefault() {

}

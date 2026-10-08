package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocolLdpRouters *****
type stateProtocolLdpRouters struct {
	validation
	obj          *otg.StateProtocolLdpRouters
	marshaller   marshalStateProtocolLdpRouters
	unMarshaller unMarshalStateProtocolLdpRouters
}

func NewStateProtocolLdpRouters() StateProtocolLdpRouters {
	obj := stateProtocolLdpRouters{obj: &otg.StateProtocolLdpRouters{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocolLdpRouters) msg() *otg.StateProtocolLdpRouters {
	return obj.obj
}

func (obj *stateProtocolLdpRouters) setMsg(msg *otg.StateProtocolLdpRouters) StateProtocolLdpRouters {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocolLdpRouters struct {
	obj *stateProtocolLdpRouters
}

type marshalStateProtocolLdpRouters interface {
	// ToProto marshals StateProtocolLdpRouters to protobuf object *otg.StateProtocolLdpRouters
	ToProto() (*otg.StateProtocolLdpRouters, error)
	// ToPbText marshals StateProtocolLdpRouters to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocolLdpRouters to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocolLdpRouters to JSON text
	ToJson() (string, error)
}

type unMarshalstateProtocolLdpRouters struct {
	obj *stateProtocolLdpRouters
}

type unMarshalStateProtocolLdpRouters interface {
	// FromProto unmarshals StateProtocolLdpRouters from protobuf object *otg.StateProtocolLdpRouters
	FromProto(msg *otg.StateProtocolLdpRouters) (StateProtocolLdpRouters, error)
	// FromPbText unmarshals StateProtocolLdpRouters from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocolLdpRouters from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocolLdpRouters from JSON text
	FromJson(value string) error
}

func (obj *stateProtocolLdpRouters) Marshal() marshalStateProtocolLdpRouters {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocolLdpRouters{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocolLdpRouters) Unmarshal() unMarshalStateProtocolLdpRouters {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocolLdpRouters{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocolLdpRouters) ToProto() (*otg.StateProtocolLdpRouters, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocolLdpRouters) FromProto(msg *otg.StateProtocolLdpRouters) (StateProtocolLdpRouters, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocolLdpRouters) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocolLdpRouters) FromPbText(value string) error {
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

func (m *marshalstateProtocolLdpRouters) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocolLdpRouters) FromYaml(value string) error {
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

func (m *marshalstateProtocolLdpRouters) ToJson() (string, error) {
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

func (m *unMarshalstateProtocolLdpRouters) FromJson(value string) error {
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

func (obj *stateProtocolLdpRouters) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocolLdpRouters) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocolLdpRouters) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocolLdpRouters) Clone() (StateProtocolLdpRouters, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocolLdpRouters()
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

// StateProtocolLdpRouters is sets state of configured LDP routers.
type StateProtocolLdpRouters interface {
	Validation
	// msg marshals StateProtocolLdpRouters to protobuf object *otg.StateProtocolLdpRouters
	// and doesn't set defaults
	msg() *otg.StateProtocolLdpRouters
	// setMsg unmarshals StateProtocolLdpRouters from protobuf object *otg.StateProtocolLdpRouters
	// and doesn't set defaults
	setMsg(*otg.StateProtocolLdpRouters) StateProtocolLdpRouters
	// provides marshal interface
	Marshal() marshalStateProtocolLdpRouters
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocolLdpRouters
	// validate validates StateProtocolLdpRouters
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocolLdpRouters, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterNames returns []string, set in StateProtocolLdpRouters.
	RouterNames() []string
	// SetRouterNames assigns []string provided by user to StateProtocolLdpRouters
	SetRouterNames(value []string) StateProtocolLdpRouters
	// State returns StateProtocolLdpRoutersStateEnum, set in StateProtocolLdpRouters
	State() StateProtocolLdpRoutersStateEnum
	// SetState assigns StateProtocolLdpRoutersStateEnum provided by user to StateProtocolLdpRouters
	SetState(value StateProtocolLdpRoutersStateEnum) StateProtocolLdpRouters
}

// The names of LDP routers for which the state has to be applied. An empty or null list will control all LDP routers.
//
// x-constraint:
// - /components/schemas/Device.LdpRouter/properties/name
//
// RouterNames returns a []string
func (obj *stateProtocolLdpRouters) RouterNames() []string {
	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	return obj.obj.RouterNames
}

// The names of LDP routers for which the state has to be applied. An empty or null list will control all LDP routers.
//
// x-constraint:
// - /components/schemas/Device.LdpRouter/properties/name
//
// SetRouterNames sets the []string value in the StateProtocolLdpRouters object
func (obj *stateProtocolLdpRouters) SetRouterNames(value []string) StateProtocolLdpRouters {

	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	obj.obj.RouterNames = value

	return obj
}

type StateProtocolLdpRoutersStateEnum string

// Enum of State on StateProtocolLdpRouters
var StateProtocolLdpRoutersState = struct {
	UP   StateProtocolLdpRoutersStateEnum
	DOWN StateProtocolLdpRoutersStateEnum
}{
	UP:   StateProtocolLdpRoutersStateEnum("up"),
	DOWN: StateProtocolLdpRoutersStateEnum("down"),
}

func (obj *stateProtocolLdpRouters) State() StateProtocolLdpRoutersStateEnum {
	return StateProtocolLdpRoutersStateEnum(obj.obj.State.Enum().String())
}

func (obj *stateProtocolLdpRouters) SetState(value StateProtocolLdpRoutersStateEnum) StateProtocolLdpRouters {
	intValue, ok := otg.StateProtocolLdpRouters_State_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolLdpRoutersStateEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocolLdpRouters_State_Enum(intValue)
	obj.obj.State = &enumValue

	return obj
}

func (obj *stateProtocolLdpRouters) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// State is required
	if obj.obj.State == nil {
		vObj.validationErrors = append(vObj.validationErrors, "State is required field on interface StateProtocolLdpRouters")
	}
}

func (obj *stateProtocolLdpRouters) setDefault() {

}

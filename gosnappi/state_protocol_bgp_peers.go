package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocolBgpPeers *****
type stateProtocolBgpPeers struct {
	validation
	obj          *otg.StateProtocolBgpPeers
	marshaller   marshalStateProtocolBgpPeers
	unMarshaller unMarshalStateProtocolBgpPeers
}

func NewStateProtocolBgpPeers() StateProtocolBgpPeers {
	obj := stateProtocolBgpPeers{obj: &otg.StateProtocolBgpPeers{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocolBgpPeers) msg() *otg.StateProtocolBgpPeers {
	return obj.obj
}

func (obj *stateProtocolBgpPeers) setMsg(msg *otg.StateProtocolBgpPeers) StateProtocolBgpPeers {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocolBgpPeers struct {
	obj *stateProtocolBgpPeers
}

type marshalStateProtocolBgpPeers interface {
	// ToProto marshals StateProtocolBgpPeers to protobuf object *otg.StateProtocolBgpPeers
	ToProto() (*otg.StateProtocolBgpPeers, error)
	// ToPbText marshals StateProtocolBgpPeers to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocolBgpPeers to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocolBgpPeers to JSON text
	ToJson() (string, error)
}

type unMarshalstateProtocolBgpPeers struct {
	obj *stateProtocolBgpPeers
}

type unMarshalStateProtocolBgpPeers interface {
	// FromProto unmarshals StateProtocolBgpPeers from protobuf object *otg.StateProtocolBgpPeers
	FromProto(msg *otg.StateProtocolBgpPeers) (StateProtocolBgpPeers, error)
	// FromPbText unmarshals StateProtocolBgpPeers from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocolBgpPeers from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocolBgpPeers from JSON text
	FromJson(value string) error
}

func (obj *stateProtocolBgpPeers) Marshal() marshalStateProtocolBgpPeers {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocolBgpPeers{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocolBgpPeers) Unmarshal() unMarshalStateProtocolBgpPeers {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocolBgpPeers{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocolBgpPeers) ToProto() (*otg.StateProtocolBgpPeers, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocolBgpPeers) FromProto(msg *otg.StateProtocolBgpPeers) (StateProtocolBgpPeers, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocolBgpPeers) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocolBgpPeers) FromPbText(value string) error {
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

func (m *marshalstateProtocolBgpPeers) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocolBgpPeers) FromYaml(value string) error {
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

func (m *marshalstateProtocolBgpPeers) ToJson() (string, error) {
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

func (m *unMarshalstateProtocolBgpPeers) FromJson(value string) error {
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

func (obj *stateProtocolBgpPeers) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocolBgpPeers) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocolBgpPeers) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocolBgpPeers) Clone() (StateProtocolBgpPeers, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocolBgpPeers()
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

// StateProtocolBgpPeers is sets state of configured BGP peers.
type StateProtocolBgpPeers interface {
	Validation
	// msg marshals StateProtocolBgpPeers to protobuf object *otg.StateProtocolBgpPeers
	// and doesn't set defaults
	msg() *otg.StateProtocolBgpPeers
	// setMsg unmarshals StateProtocolBgpPeers from protobuf object *otg.StateProtocolBgpPeers
	// and doesn't set defaults
	setMsg(*otg.StateProtocolBgpPeers) StateProtocolBgpPeers
	// provides marshal interface
	Marshal() marshalStateProtocolBgpPeers
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocolBgpPeers
	// validate validates StateProtocolBgpPeers
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocolBgpPeers, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PeerNames returns []string, set in StateProtocolBgpPeers.
	PeerNames() []string
	// SetPeerNames assigns []string provided by user to StateProtocolBgpPeers
	SetPeerNames(value []string) StateProtocolBgpPeers
	// State returns StateProtocolBgpPeersStateEnum, set in StateProtocolBgpPeers
	State() StateProtocolBgpPeersStateEnum
	// SetState assigns StateProtocolBgpPeersStateEnum provided by user to StateProtocolBgpPeers
	SetState(value StateProtocolBgpPeersStateEnum) StateProtocolBgpPeers
}

// The names of BGP peers for which the state has to be applied. An empty or null list will control all BGP peers.
//
// x-constraint:
// - /components/schemas/Bgp.V4Peer/properties/name
// - /components/schemas/Bgp.V6Peer/properties/name
//
// x-constraint:
// - /components/schemas/Bgp.V4Peer/properties/name
// - /components/schemas/Bgp.V6Peer/properties/name
//
// PeerNames returns a []string
func (obj *stateProtocolBgpPeers) PeerNames() []string {
	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	return obj.obj.PeerNames
}

// The names of BGP peers for which the state has to be applied. An empty or null list will control all BGP peers.
//
// x-constraint:
// - /components/schemas/Bgp.V4Peer/properties/name
// - /components/schemas/Bgp.V6Peer/properties/name
//
// x-constraint:
// - /components/schemas/Bgp.V4Peer/properties/name
// - /components/schemas/Bgp.V6Peer/properties/name
//
// SetPeerNames sets the []string value in the StateProtocolBgpPeers object
func (obj *stateProtocolBgpPeers) SetPeerNames(value []string) StateProtocolBgpPeers {

	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	obj.obj.PeerNames = value

	return obj
}

type StateProtocolBgpPeersStateEnum string

// Enum of State on StateProtocolBgpPeers
var StateProtocolBgpPeersState = struct {
	UP   StateProtocolBgpPeersStateEnum
	DOWN StateProtocolBgpPeersStateEnum
}{
	UP:   StateProtocolBgpPeersStateEnum("up"),
	DOWN: StateProtocolBgpPeersStateEnum("down"),
}

func (obj *stateProtocolBgpPeers) State() StateProtocolBgpPeersStateEnum {
	return StateProtocolBgpPeersStateEnum(obj.obj.State.Enum().String())
}

func (obj *stateProtocolBgpPeers) SetState(value StateProtocolBgpPeersStateEnum) StateProtocolBgpPeers {
	intValue, ok := otg.StateProtocolBgpPeers_State_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolBgpPeersStateEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocolBgpPeers_State_Enum(intValue)
	obj.obj.State = &enumValue

	return obj
}

func (obj *stateProtocolBgpPeers) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// State is required
	if obj.obj.State == nil {
		vObj.validationErrors = append(vObj.validationErrors, "State is required field on interface StateProtocolBgpPeers")
	}
}

func (obj *stateProtocolBgpPeers) setDefault() {

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocolRocev2Peers *****
type stateProtocolRocev2Peers struct {
	validation
	obj          *otg.StateProtocolRocev2Peers
	marshaller   marshalStateProtocolRocev2Peers
	unMarshaller unMarshalStateProtocolRocev2Peers
}

func NewStateProtocolRocev2Peers() StateProtocolRocev2Peers {
	obj := stateProtocolRocev2Peers{obj: &otg.StateProtocolRocev2Peers{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocolRocev2Peers) msg() *otg.StateProtocolRocev2Peers {
	return obj.obj
}

func (obj *stateProtocolRocev2Peers) setMsg(msg *otg.StateProtocolRocev2Peers) StateProtocolRocev2Peers {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocolRocev2Peers struct {
	obj *stateProtocolRocev2Peers
}

type marshalStateProtocolRocev2Peers interface {
	// ToProto marshals StateProtocolRocev2Peers to protobuf object *otg.StateProtocolRocev2Peers
	ToProto() (*otg.StateProtocolRocev2Peers, error)
	// ToPbText marshals StateProtocolRocev2Peers to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocolRocev2Peers to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocolRocev2Peers to JSON text
	ToJson() (string, error)
}

type unMarshalstateProtocolRocev2Peers struct {
	obj *stateProtocolRocev2Peers
}

type unMarshalStateProtocolRocev2Peers interface {
	// FromProto unmarshals StateProtocolRocev2Peers from protobuf object *otg.StateProtocolRocev2Peers
	FromProto(msg *otg.StateProtocolRocev2Peers) (StateProtocolRocev2Peers, error)
	// FromPbText unmarshals StateProtocolRocev2Peers from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocolRocev2Peers from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocolRocev2Peers from JSON text
	FromJson(value string) error
}

func (obj *stateProtocolRocev2Peers) Marshal() marshalStateProtocolRocev2Peers {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocolRocev2Peers{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocolRocev2Peers) Unmarshal() unMarshalStateProtocolRocev2Peers {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocolRocev2Peers{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocolRocev2Peers) ToProto() (*otg.StateProtocolRocev2Peers, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocolRocev2Peers) FromProto(msg *otg.StateProtocolRocev2Peers) (StateProtocolRocev2Peers, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocolRocev2Peers) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocolRocev2Peers) FromPbText(value string) error {
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

func (m *marshalstateProtocolRocev2Peers) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocolRocev2Peers) FromYaml(value string) error {
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

func (m *marshalstateProtocolRocev2Peers) ToJson() (string, error) {
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

func (m *unMarshalstateProtocolRocev2Peers) FromJson(value string) error {
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

func (obj *stateProtocolRocev2Peers) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocolRocev2Peers) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocolRocev2Peers) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocolRocev2Peers) Clone() (StateProtocolRocev2Peers, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocolRocev2Peers()
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

// StateProtocolRocev2Peers is sets state of configured RoCEv2 peers.
type StateProtocolRocev2Peers interface {
	Validation
	// msg marshals StateProtocolRocev2Peers to protobuf object *otg.StateProtocolRocev2Peers
	// and doesn't set defaults
	msg() *otg.StateProtocolRocev2Peers
	// setMsg unmarshals StateProtocolRocev2Peers from protobuf object *otg.StateProtocolRocev2Peers
	// and doesn't set defaults
	setMsg(*otg.StateProtocolRocev2Peers) StateProtocolRocev2Peers
	// provides marshal interface
	Marshal() marshalStateProtocolRocev2Peers
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocolRocev2Peers
	// validate validates StateProtocolRocev2Peers
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocolRocev2Peers, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PeerNames returns []string, set in StateProtocolRocev2Peers.
	PeerNames() []string
	// SetPeerNames assigns []string provided by user to StateProtocolRocev2Peers
	SetPeerNames(value []string) StateProtocolRocev2Peers
	// State returns StateProtocolRocev2PeersStateEnum, set in StateProtocolRocev2Peers
	State() StateProtocolRocev2PeersStateEnum
	// SetState assigns StateProtocolRocev2PeersStateEnum provided by user to StateProtocolRocev2Peers
	SetState(value StateProtocolRocev2PeersStateEnum) StateProtocolRocev2Peers
}

// The names of RoCEv2 peers for which the state has to be applied. An empty or null list will control all RoCEv2 peers.
//
// x-constraint:
// - /components/schemas/Rocev2.V4Peer/properties/name
// - /components/schemas/Rocev2.V6Peer/properties/name
//
// PeerNames returns a []string
func (obj *stateProtocolRocev2Peers) PeerNames() []string {
	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	return obj.obj.PeerNames
}

// The names of RoCEv2 peers for which the state has to be applied. An empty or null list will control all RoCEv2 peers.
//
// x-constraint:
// - /components/schemas/Rocev2.V4Peer/properties/name
// - /components/schemas/Rocev2.V6Peer/properties/name
//
// SetPeerNames sets the []string value in the StateProtocolRocev2Peers object
func (obj *stateProtocolRocev2Peers) SetPeerNames(value []string) StateProtocolRocev2Peers {

	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	obj.obj.PeerNames = value

	return obj
}

type StateProtocolRocev2PeersStateEnum string

// Enum of State on StateProtocolRocev2Peers
var StateProtocolRocev2PeersState = struct {
	UP   StateProtocolRocev2PeersStateEnum
	DOWN StateProtocolRocev2PeersStateEnum
}{
	UP:   StateProtocolRocev2PeersStateEnum("up"),
	DOWN: StateProtocolRocev2PeersStateEnum("down"),
}

func (obj *stateProtocolRocev2Peers) State() StateProtocolRocev2PeersStateEnum {
	return StateProtocolRocev2PeersStateEnum(obj.obj.State.Enum().String())
}

func (obj *stateProtocolRocev2Peers) SetState(value StateProtocolRocev2PeersStateEnum) StateProtocolRocev2Peers {
	intValue, ok := otg.StateProtocolRocev2Peers_State_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolRocev2PeersStateEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocolRocev2Peers_State_Enum(intValue)
	obj.obj.State = &enumValue

	return obj
}

func (obj *stateProtocolRocev2Peers) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// State is required
	if obj.obj.State == nil {
		vObj.validationErrors = append(vObj.validationErrors, "State is required field on interface StateProtocolRocev2Peers")
	}
}

func (obj *stateProtocolRocev2Peers) setDefault() {

}

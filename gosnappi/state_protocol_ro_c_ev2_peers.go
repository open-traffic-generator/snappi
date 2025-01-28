package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocolRoCEv2Peers *****
type stateProtocolRoCEv2Peers struct {
	validation
	obj          *otg.StateProtocolRoCEv2Peers
	marshaller   marshalStateProtocolRoCEv2Peers
	unMarshaller unMarshalStateProtocolRoCEv2Peers
}

func NewStateProtocolRoCEv2Peers() StateProtocolRoCEv2Peers {
	obj := stateProtocolRoCEv2Peers{obj: &otg.StateProtocolRoCEv2Peers{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocolRoCEv2Peers) msg() *otg.StateProtocolRoCEv2Peers {
	return obj.obj
}

func (obj *stateProtocolRoCEv2Peers) setMsg(msg *otg.StateProtocolRoCEv2Peers) StateProtocolRoCEv2Peers {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocolRoCEv2Peers struct {
	obj *stateProtocolRoCEv2Peers
}

type marshalStateProtocolRoCEv2Peers interface {
	// ToProto marshals StateProtocolRoCEv2Peers to protobuf object *otg.StateProtocolRoCEv2Peers
	ToProto() (*otg.StateProtocolRoCEv2Peers, error)
	// ToPbText marshals StateProtocolRoCEv2Peers to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocolRoCEv2Peers to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocolRoCEv2Peers to JSON text
	ToJson() (string, error)
}

type unMarshalstateProtocolRoCEv2Peers struct {
	obj *stateProtocolRoCEv2Peers
}

type unMarshalStateProtocolRoCEv2Peers interface {
	// FromProto unmarshals StateProtocolRoCEv2Peers from protobuf object *otg.StateProtocolRoCEv2Peers
	FromProto(msg *otg.StateProtocolRoCEv2Peers) (StateProtocolRoCEv2Peers, error)
	// FromPbText unmarshals StateProtocolRoCEv2Peers from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocolRoCEv2Peers from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocolRoCEv2Peers from JSON text
	FromJson(value string) error
}

func (obj *stateProtocolRoCEv2Peers) Marshal() marshalStateProtocolRoCEv2Peers {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocolRoCEv2Peers{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocolRoCEv2Peers) Unmarshal() unMarshalStateProtocolRoCEv2Peers {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocolRoCEv2Peers{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocolRoCEv2Peers) ToProto() (*otg.StateProtocolRoCEv2Peers, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocolRoCEv2Peers) FromProto(msg *otg.StateProtocolRoCEv2Peers) (StateProtocolRoCEv2Peers, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocolRoCEv2Peers) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocolRoCEv2Peers) FromPbText(value string) error {
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

func (m *marshalstateProtocolRoCEv2Peers) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocolRoCEv2Peers) FromYaml(value string) error {
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

func (m *marshalstateProtocolRoCEv2Peers) ToJson() (string, error) {
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

func (m *unMarshalstateProtocolRoCEv2Peers) FromJson(value string) error {
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

func (obj *stateProtocolRoCEv2Peers) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocolRoCEv2Peers) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocolRoCEv2Peers) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocolRoCEv2Peers) Clone() (StateProtocolRoCEv2Peers, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocolRoCEv2Peers()
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

// StateProtocolRoCEv2Peers is sets state of configured RoCEv2 peers.
type StateProtocolRoCEv2Peers interface {
	Validation
	// msg marshals StateProtocolRoCEv2Peers to protobuf object *otg.StateProtocolRoCEv2Peers
	// and doesn't set defaults
	msg() *otg.StateProtocolRoCEv2Peers
	// setMsg unmarshals StateProtocolRoCEv2Peers from protobuf object *otg.StateProtocolRoCEv2Peers
	// and doesn't set defaults
	setMsg(*otg.StateProtocolRoCEv2Peers) StateProtocolRoCEv2Peers
	// provides marshal interface
	Marshal() marshalStateProtocolRoCEv2Peers
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocolRoCEv2Peers
	// validate validates StateProtocolRoCEv2Peers
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocolRoCEv2Peers, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PeerNames returns []string, set in StateProtocolRoCEv2Peers.
	PeerNames() []string
	// SetPeerNames assigns []string provided by user to StateProtocolRoCEv2Peers
	SetPeerNames(value []string) StateProtocolRoCEv2Peers
	// State returns StateProtocolRoCEv2PeersStateEnum, set in StateProtocolRoCEv2Peers
	State() StateProtocolRoCEv2PeersStateEnum
	// SetState assigns StateProtocolRoCEv2PeersStateEnum provided by user to StateProtocolRoCEv2Peers
	SetState(value StateProtocolRoCEv2PeersStateEnum) StateProtocolRoCEv2Peers
}

// The names of RoCEv2 peers for which the state has to be applied. An empty or null list will control all RoCEv2 peers.
//
// x-constraint:
// - /components/schemas/RoCEv2.V4Peer/properties/name
// - /components/schemas/RoCEv2.V6Peer/properties/name
//
// PeerNames returns a []string
func (obj *stateProtocolRoCEv2Peers) PeerNames() []string {
	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	return obj.obj.PeerNames
}

// The names of RoCEv2 peers for which the state has to be applied. An empty or null list will control all RoCEv2 peers.
//
// x-constraint:
// - /components/schemas/RoCEv2.V4Peer/properties/name
// - /components/schemas/RoCEv2.V6Peer/properties/name
//
// SetPeerNames sets the []string value in the StateProtocolRoCEv2Peers object
func (obj *stateProtocolRoCEv2Peers) SetPeerNames(value []string) StateProtocolRoCEv2Peers {

	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	obj.obj.PeerNames = value

	return obj
}

type StateProtocolRoCEv2PeersStateEnum string

// Enum of State on StateProtocolRoCEv2Peers
var StateProtocolRoCEv2PeersState = struct {
	UP   StateProtocolRoCEv2PeersStateEnum
	DOWN StateProtocolRoCEv2PeersStateEnum
}{
	UP:   StateProtocolRoCEv2PeersStateEnum("up"),
	DOWN: StateProtocolRoCEv2PeersStateEnum("down"),
}

func (obj *stateProtocolRoCEv2Peers) State() StateProtocolRoCEv2PeersStateEnum {
	return StateProtocolRoCEv2PeersStateEnum(obj.obj.State.Enum().String())
}

func (obj *stateProtocolRoCEv2Peers) SetState(value StateProtocolRoCEv2PeersStateEnum) StateProtocolRoCEv2Peers {
	intValue, ok := otg.StateProtocolRoCEv2Peers_State_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolRoCEv2PeersStateEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocolRoCEv2Peers_State_Enum(intValue)
	obj.obj.State = &enumValue

	return obj
}

func (obj *stateProtocolRoCEv2Peers) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// State is required
	if obj.obj.State == nil {
		vObj.validationErrors = append(vObj.validationErrors, "State is required field on interface StateProtocolRoCEv2Peers")
	}
}

func (obj *stateProtocolRoCEv2Peers) setDefault() {

}

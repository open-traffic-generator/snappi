package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocolBgp *****
type stateProtocolBgp struct {
	validation
	obj          *otg.StateProtocolBgp
	marshaller   marshalStateProtocolBgp
	unMarshaller unMarshalStateProtocolBgp
	peersHolder  StateProtocolBgpPeers
}

func NewStateProtocolBgp() StateProtocolBgp {
	obj := stateProtocolBgp{obj: &otg.StateProtocolBgp{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocolBgp) msg() *otg.StateProtocolBgp {
	return obj.obj
}

func (obj *stateProtocolBgp) setMsg(msg *otg.StateProtocolBgp) StateProtocolBgp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocolBgp struct {
	obj *stateProtocolBgp
}

type marshalStateProtocolBgp interface {
	// ToProto marshals StateProtocolBgp to protobuf object *otg.StateProtocolBgp
	ToProto() (*otg.StateProtocolBgp, error)
	// ToPbText marshals StateProtocolBgp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocolBgp to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocolBgp to JSON text
	ToJson() (string, error)
}

type unMarshalstateProtocolBgp struct {
	obj *stateProtocolBgp
}

type unMarshalStateProtocolBgp interface {
	// FromProto unmarshals StateProtocolBgp from protobuf object *otg.StateProtocolBgp
	FromProto(msg *otg.StateProtocolBgp) (StateProtocolBgp, error)
	// FromPbText unmarshals StateProtocolBgp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocolBgp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocolBgp from JSON text
	FromJson(value string) error
}

func (obj *stateProtocolBgp) Marshal() marshalStateProtocolBgp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocolBgp{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocolBgp) Unmarshal() unMarshalStateProtocolBgp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocolBgp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocolBgp) ToProto() (*otg.StateProtocolBgp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocolBgp) FromProto(msg *otg.StateProtocolBgp) (StateProtocolBgp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocolBgp) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocolBgp) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalstateProtocolBgp) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocolBgp) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalstateProtocolBgp) ToJson() (string, error) {
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

func (m *unMarshalstateProtocolBgp) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *stateProtocolBgp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocolBgp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocolBgp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocolBgp) Clone() (StateProtocolBgp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocolBgp()
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

func (obj *stateProtocolBgp) setNil() {
	obj.peersHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// StateProtocolBgp is sets state of configured BGP peers.
type StateProtocolBgp interface {
	Validation
	// msg marshals StateProtocolBgp to protobuf object *otg.StateProtocolBgp
	// and doesn't set defaults
	msg() *otg.StateProtocolBgp
	// setMsg unmarshals StateProtocolBgp from protobuf object *otg.StateProtocolBgp
	// and doesn't set defaults
	setMsg(*otg.StateProtocolBgp) StateProtocolBgp
	// provides marshal interface
	Marshal() marshalStateProtocolBgp
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocolBgp
	// validate validates StateProtocolBgp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocolBgp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns StateProtocolBgpChoiceEnum, set in StateProtocolBgp
	Choice() StateProtocolBgpChoiceEnum
	// setChoice assigns StateProtocolBgpChoiceEnum provided by user to StateProtocolBgp
	setChoice(value StateProtocolBgpChoiceEnum) StateProtocolBgp
	// Peers returns StateProtocolBgpPeers, set in StateProtocolBgp.
	// StateProtocolBgpPeers is sets state of configured BGP peers.
	Peers() StateProtocolBgpPeers
	// SetPeers assigns StateProtocolBgpPeers provided by user to StateProtocolBgp.
	// StateProtocolBgpPeers is sets state of configured BGP peers.
	SetPeers(value StateProtocolBgpPeers) StateProtocolBgp
	// HasPeers checks if Peers has been set in StateProtocolBgp
	HasPeers() bool
	setNil()
}

type StateProtocolBgpChoiceEnum string

// Enum of Choice on StateProtocolBgp
var StateProtocolBgpChoice = struct {
	PEERS StateProtocolBgpChoiceEnum
}{
	PEERS: StateProtocolBgpChoiceEnum("peers"),
}

func (obj *stateProtocolBgp) Choice() StateProtocolBgpChoiceEnum {
	return StateProtocolBgpChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *stateProtocolBgp) setChoice(value StateProtocolBgpChoiceEnum) StateProtocolBgp {
	intValue, ok := otg.StateProtocolBgp_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolBgpChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocolBgp_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Peers = nil
	obj.peersHolder = nil

	if value == StateProtocolBgpChoice.PEERS {
		obj.obj.Peers = NewStateProtocolBgpPeers().msg()
	}

	return obj
}

// description is TBD
// Peers returns a StateProtocolBgpPeers
func (obj *stateProtocolBgp) Peers() StateProtocolBgpPeers {
	if obj.obj.Peers == nil {
		obj.setChoice(StateProtocolBgpChoice.PEERS)
	}
	if obj.peersHolder == nil {
		obj.peersHolder = &stateProtocolBgpPeers{obj: obj.obj.Peers}
	}
	return obj.peersHolder
}

// description is TBD
// Peers returns a StateProtocolBgpPeers
func (obj *stateProtocolBgp) HasPeers() bool {
	return obj.obj.Peers != nil
}

// description is TBD
// SetPeers sets the StateProtocolBgpPeers value in the StateProtocolBgp object
func (obj *stateProtocolBgp) SetPeers(value StateProtocolBgpPeers) StateProtocolBgp {
	obj.setChoice(StateProtocolBgpChoice.PEERS)
	obj.peersHolder = nil
	obj.obj.Peers = value.msg()

	return obj
}

func (obj *stateProtocolBgp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface StateProtocolBgp")
	}

	if obj.obj.Peers != nil {

		obj.Peers().validateObj(vObj, set_default)
	}

}

func (obj *stateProtocolBgp) setDefault() {

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocolRoCEv2 *****
type stateProtocolRoCEv2 struct {
	validation
	obj          *otg.StateProtocolRoCEv2
	marshaller   marshalStateProtocolRoCEv2
	unMarshaller unMarshalStateProtocolRoCEv2
	peersHolder  StateProtocolRoCEv2Peers
}

func NewStateProtocolRoCEv2() StateProtocolRoCEv2 {
	obj := stateProtocolRoCEv2{obj: &otg.StateProtocolRoCEv2{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocolRoCEv2) msg() *otg.StateProtocolRoCEv2 {
	return obj.obj
}

func (obj *stateProtocolRoCEv2) setMsg(msg *otg.StateProtocolRoCEv2) StateProtocolRoCEv2 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocolRoCEv2 struct {
	obj *stateProtocolRoCEv2
}

type marshalStateProtocolRoCEv2 interface {
	// ToProto marshals StateProtocolRoCEv2 to protobuf object *otg.StateProtocolRoCEv2
	ToProto() (*otg.StateProtocolRoCEv2, error)
	// ToPbText marshals StateProtocolRoCEv2 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocolRoCEv2 to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocolRoCEv2 to JSON text
	ToJson() (string, error)
}

type unMarshalstateProtocolRoCEv2 struct {
	obj *stateProtocolRoCEv2
}

type unMarshalStateProtocolRoCEv2 interface {
	// FromProto unmarshals StateProtocolRoCEv2 from protobuf object *otg.StateProtocolRoCEv2
	FromProto(msg *otg.StateProtocolRoCEv2) (StateProtocolRoCEv2, error)
	// FromPbText unmarshals StateProtocolRoCEv2 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocolRoCEv2 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocolRoCEv2 from JSON text
	FromJson(value string) error
}

func (obj *stateProtocolRoCEv2) Marshal() marshalStateProtocolRoCEv2 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocolRoCEv2{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocolRoCEv2) Unmarshal() unMarshalStateProtocolRoCEv2 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocolRoCEv2{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocolRoCEv2) ToProto() (*otg.StateProtocolRoCEv2, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocolRoCEv2) FromProto(msg *otg.StateProtocolRoCEv2) (StateProtocolRoCEv2, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocolRoCEv2) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocolRoCEv2) FromPbText(value string) error {
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

func (m *marshalstateProtocolRoCEv2) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocolRoCEv2) FromYaml(value string) error {
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

func (m *marshalstateProtocolRoCEv2) ToJson() (string, error) {
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

func (m *unMarshalstateProtocolRoCEv2) FromJson(value string) error {
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

func (obj *stateProtocolRoCEv2) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocolRoCEv2) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocolRoCEv2) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocolRoCEv2) Clone() (StateProtocolRoCEv2, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocolRoCEv2()
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

func (obj *stateProtocolRoCEv2) setNil() {
	obj.peersHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// StateProtocolRoCEv2 is sets state of configured RoCEv2 peers.
type StateProtocolRoCEv2 interface {
	Validation
	// msg marshals StateProtocolRoCEv2 to protobuf object *otg.StateProtocolRoCEv2
	// and doesn't set defaults
	msg() *otg.StateProtocolRoCEv2
	// setMsg unmarshals StateProtocolRoCEv2 from protobuf object *otg.StateProtocolRoCEv2
	// and doesn't set defaults
	setMsg(*otg.StateProtocolRoCEv2) StateProtocolRoCEv2
	// provides marshal interface
	Marshal() marshalStateProtocolRoCEv2
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocolRoCEv2
	// validate validates StateProtocolRoCEv2
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocolRoCEv2, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns StateProtocolRoCEv2ChoiceEnum, set in StateProtocolRoCEv2
	Choice() StateProtocolRoCEv2ChoiceEnum
	// setChoice assigns StateProtocolRoCEv2ChoiceEnum provided by user to StateProtocolRoCEv2
	setChoice(value StateProtocolRoCEv2ChoiceEnum) StateProtocolRoCEv2
	// Peers returns StateProtocolRoCEv2Peers, set in StateProtocolRoCEv2.
	// StateProtocolRoCEv2Peers is sets state of configured RoCEv2 peers.
	Peers() StateProtocolRoCEv2Peers
	// SetPeers assigns StateProtocolRoCEv2Peers provided by user to StateProtocolRoCEv2.
	// StateProtocolRoCEv2Peers is sets state of configured RoCEv2 peers.
	SetPeers(value StateProtocolRoCEv2Peers) StateProtocolRoCEv2
	// HasPeers checks if Peers has been set in StateProtocolRoCEv2
	HasPeers() bool
	setNil()
}

type StateProtocolRoCEv2ChoiceEnum string

// Enum of Choice on StateProtocolRoCEv2
var StateProtocolRoCEv2Choice = struct {
	PEERS StateProtocolRoCEv2ChoiceEnum
}{
	PEERS: StateProtocolRoCEv2ChoiceEnum("peers"),
}

func (obj *stateProtocolRoCEv2) Choice() StateProtocolRoCEv2ChoiceEnum {
	return StateProtocolRoCEv2ChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *stateProtocolRoCEv2) setChoice(value StateProtocolRoCEv2ChoiceEnum) StateProtocolRoCEv2 {
	intValue, ok := otg.StateProtocolRoCEv2_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolRoCEv2ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocolRoCEv2_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Peers = nil
	obj.peersHolder = nil

	if value == StateProtocolRoCEv2Choice.PEERS {
		obj.obj.Peers = NewStateProtocolRoCEv2Peers().msg()
	}

	return obj
}

// description is TBD
// Peers returns a StateProtocolRoCEv2Peers
func (obj *stateProtocolRoCEv2) Peers() StateProtocolRoCEv2Peers {
	if obj.obj.Peers == nil {
		obj.setChoice(StateProtocolRoCEv2Choice.PEERS)
	}
	if obj.peersHolder == nil {
		obj.peersHolder = &stateProtocolRoCEv2Peers{obj: obj.obj.Peers}
	}
	return obj.peersHolder
}

// description is TBD
// Peers returns a StateProtocolRoCEv2Peers
func (obj *stateProtocolRoCEv2) HasPeers() bool {
	return obj.obj.Peers != nil
}

// description is TBD
// SetPeers sets the StateProtocolRoCEv2Peers value in the StateProtocolRoCEv2 object
func (obj *stateProtocolRoCEv2) SetPeers(value StateProtocolRoCEv2Peers) StateProtocolRoCEv2 {
	obj.setChoice(StateProtocolRoCEv2Choice.PEERS)
	obj.peersHolder = nil
	obj.obj.Peers = value.msg()

	return obj
}

func (obj *stateProtocolRoCEv2) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface StateProtocolRoCEv2")
	}

	if obj.obj.Peers != nil {

		obj.Peers().validateObj(vObj, set_default)
	}

}

func (obj *stateProtocolRoCEv2) setDefault() {
	var choices_set int = 0
	var choice StateProtocolRoCEv2ChoiceEnum

	if obj.obj.Peers != nil {
		choices_set += 1
		choice = StateProtocolRoCEv2Choice.PEERS
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in StateProtocolRoCEv2")
			}
		} else {
			intVal := otg.StateProtocolRoCEv2_Choice_Enum_value[string(choice)]
			enumValue := otg.StateProtocolRoCEv2_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

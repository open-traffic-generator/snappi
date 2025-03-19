package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocolRocev2 *****
type stateProtocolRocev2 struct {
	validation
	obj          *otg.StateProtocolRocev2
	marshaller   marshalStateProtocolRocev2
	unMarshaller unMarshalStateProtocolRocev2
	peersHolder  StateProtocolRocev2Peers
}

func NewStateProtocolRocev2() StateProtocolRocev2 {
	obj := stateProtocolRocev2{obj: &otg.StateProtocolRocev2{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocolRocev2) msg() *otg.StateProtocolRocev2 {
	return obj.obj
}

func (obj *stateProtocolRocev2) setMsg(msg *otg.StateProtocolRocev2) StateProtocolRocev2 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocolRocev2 struct {
	obj *stateProtocolRocev2
}

type marshalStateProtocolRocev2 interface {
	// ToProto marshals StateProtocolRocev2 to protobuf object *otg.StateProtocolRocev2
	ToProto() (*otg.StateProtocolRocev2, error)
	// ToPbText marshals StateProtocolRocev2 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocolRocev2 to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocolRocev2 to JSON text
	ToJson() (string, error)
}

type unMarshalstateProtocolRocev2 struct {
	obj *stateProtocolRocev2
}

type unMarshalStateProtocolRocev2 interface {
	// FromProto unmarshals StateProtocolRocev2 from protobuf object *otg.StateProtocolRocev2
	FromProto(msg *otg.StateProtocolRocev2) (StateProtocolRocev2, error)
	// FromPbText unmarshals StateProtocolRocev2 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocolRocev2 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocolRocev2 from JSON text
	FromJson(value string) error
}

func (obj *stateProtocolRocev2) Marshal() marshalStateProtocolRocev2 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocolRocev2{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocolRocev2) Unmarshal() unMarshalStateProtocolRocev2 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocolRocev2{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocolRocev2) ToProto() (*otg.StateProtocolRocev2, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocolRocev2) FromProto(msg *otg.StateProtocolRocev2) (StateProtocolRocev2, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocolRocev2) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocolRocev2) FromPbText(value string) error {
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

func (m *marshalstateProtocolRocev2) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocolRocev2) FromYaml(value string) error {
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

func (m *marshalstateProtocolRocev2) ToJson() (string, error) {
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

func (m *unMarshalstateProtocolRocev2) FromJson(value string) error {
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

func (obj *stateProtocolRocev2) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocolRocev2) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocolRocev2) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocolRocev2) Clone() (StateProtocolRocev2, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocolRocev2()
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

func (obj *stateProtocolRocev2) setNil() {
	obj.peersHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// StateProtocolRocev2 is sets state of configured RoCEv2 peers.
type StateProtocolRocev2 interface {
	Validation
	// msg marshals StateProtocolRocev2 to protobuf object *otg.StateProtocolRocev2
	// and doesn't set defaults
	msg() *otg.StateProtocolRocev2
	// setMsg unmarshals StateProtocolRocev2 from protobuf object *otg.StateProtocolRocev2
	// and doesn't set defaults
	setMsg(*otg.StateProtocolRocev2) StateProtocolRocev2
	// provides marshal interface
	Marshal() marshalStateProtocolRocev2
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocolRocev2
	// validate validates StateProtocolRocev2
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocolRocev2, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns StateProtocolRocev2ChoiceEnum, set in StateProtocolRocev2
	Choice() StateProtocolRocev2ChoiceEnum
	// setChoice assigns StateProtocolRocev2ChoiceEnum provided by user to StateProtocolRocev2
	setChoice(value StateProtocolRocev2ChoiceEnum) StateProtocolRocev2
	// Peers returns StateProtocolRocev2Peers, set in StateProtocolRocev2.
	// StateProtocolRocev2Peers is sets state of configured RoCEv2 peers.
	Peers() StateProtocolRocev2Peers
	// SetPeers assigns StateProtocolRocev2Peers provided by user to StateProtocolRocev2.
	// StateProtocolRocev2Peers is sets state of configured RoCEv2 peers.
	SetPeers(value StateProtocolRocev2Peers) StateProtocolRocev2
	// HasPeers checks if Peers has been set in StateProtocolRocev2
	HasPeers() bool
	setNil()
}

type StateProtocolRocev2ChoiceEnum string

// Enum of Choice on StateProtocolRocev2
var StateProtocolRocev2Choice = struct {
	PEERS StateProtocolRocev2ChoiceEnum
}{
	PEERS: StateProtocolRocev2ChoiceEnum("peers"),
}

func (obj *stateProtocolRocev2) Choice() StateProtocolRocev2ChoiceEnum {
	return StateProtocolRocev2ChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *stateProtocolRocev2) setChoice(value StateProtocolRocev2ChoiceEnum) StateProtocolRocev2 {
	intValue, ok := otg.StateProtocolRocev2_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolRocev2ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocolRocev2_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Peers = nil
	obj.peersHolder = nil

	if value == StateProtocolRocev2Choice.PEERS {
		obj.obj.Peers = NewStateProtocolRocev2Peers().msg()
	}

	return obj
}

// description is TBD
// Peers returns a StateProtocolRocev2Peers
func (obj *stateProtocolRocev2) Peers() StateProtocolRocev2Peers {
	if obj.obj.Peers == nil {
		obj.setChoice(StateProtocolRocev2Choice.PEERS)
	}
	if obj.peersHolder == nil {
		obj.peersHolder = &stateProtocolRocev2Peers{obj: obj.obj.Peers}
	}
	return obj.peersHolder
}

// description is TBD
// Peers returns a StateProtocolRocev2Peers
func (obj *stateProtocolRocev2) HasPeers() bool {
	return obj.obj.Peers != nil
}

// description is TBD
// SetPeers sets the StateProtocolRocev2Peers value in the StateProtocolRocev2 object
func (obj *stateProtocolRocev2) SetPeers(value StateProtocolRocev2Peers) StateProtocolRocev2 {
	obj.setChoice(StateProtocolRocev2Choice.PEERS)
	obj.peersHolder = nil
	obj.obj.Peers = value.msg()

	return obj
}

func (obj *stateProtocolRocev2) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface StateProtocolRocev2")
	}

	if obj.obj.Peers != nil {

		obj.Peers().validateObj(vObj, set_default)
	}

}

func (obj *stateProtocolRocev2) setDefault() {
	var choices_set int = 0
	var choice StateProtocolRocev2ChoiceEnum

	if obj.obj.Peers != nil {
		choices_set += 1
		choice = StateProtocolRocev2Choice.PEERS
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in StateProtocolRocev2")
			}
		} else {
			intVal := otg.StateProtocolRocev2_Choice_Enum_value[string(choice)]
			enumValue := otg.StateProtocolRocev2_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

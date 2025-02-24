package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateTrafficFlowTransmit *****
type stateTrafficFlowTransmit struct {
	validation
	obj          *otg.StateTrafficFlowTransmit
	marshaller   marshalStateTrafficFlowTransmit
	unMarshaller unMarshalStateTrafficFlowTransmit
}

func NewStateTrafficFlowTransmit() StateTrafficFlowTransmit {
	obj := stateTrafficFlowTransmit{obj: &otg.StateTrafficFlowTransmit{}}
	obj.setDefault()
	return &obj
}

func (obj *stateTrafficFlowTransmit) msg() *otg.StateTrafficFlowTransmit {
	return obj.obj
}

func (obj *stateTrafficFlowTransmit) setMsg(msg *otg.StateTrafficFlowTransmit) StateTrafficFlowTransmit {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateTrafficFlowTransmit struct {
	obj *stateTrafficFlowTransmit
}

type marshalStateTrafficFlowTransmit interface {
	// ToProto marshals StateTrafficFlowTransmit to protobuf object *otg.StateTrafficFlowTransmit
	ToProto() (*otg.StateTrafficFlowTransmit, error)
	// ToPbText marshals StateTrafficFlowTransmit to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateTrafficFlowTransmit to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateTrafficFlowTransmit to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals StateTrafficFlowTransmit to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalstateTrafficFlowTransmit struct {
	obj *stateTrafficFlowTransmit
}

type unMarshalStateTrafficFlowTransmit interface {
	// FromProto unmarshals StateTrafficFlowTransmit from protobuf object *otg.StateTrafficFlowTransmit
	FromProto(msg *otg.StateTrafficFlowTransmit) (StateTrafficFlowTransmit, error)
	// FromPbText unmarshals StateTrafficFlowTransmit from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateTrafficFlowTransmit from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateTrafficFlowTransmit from JSON text
	FromJson(value string) error
}

func (obj *stateTrafficFlowTransmit) Marshal() marshalStateTrafficFlowTransmit {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateTrafficFlowTransmit{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateTrafficFlowTransmit) Unmarshal() unMarshalStateTrafficFlowTransmit {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateTrafficFlowTransmit{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateTrafficFlowTransmit) ToProto() (*otg.StateTrafficFlowTransmit, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateTrafficFlowTransmit) FromProto(msg *otg.StateTrafficFlowTransmit) (StateTrafficFlowTransmit, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateTrafficFlowTransmit) ToPbText() (string, error) {
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

func (m *unMarshalstateTrafficFlowTransmit) FromPbText(value string) error {
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

func (m *marshalstateTrafficFlowTransmit) ToYaml() (string, error) {
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

func (m *unMarshalstateTrafficFlowTransmit) FromYaml(value string) error {
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

func (m *marshalstateTrafficFlowTransmit) ToJsonRaw() (string, error) {
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

func (m *marshalstateTrafficFlowTransmit) ToJson() (string, error) {
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

func (m *unMarshalstateTrafficFlowTransmit) FromJson(value string) error {
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

func (obj *stateTrafficFlowTransmit) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateTrafficFlowTransmit) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateTrafficFlowTransmit) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateTrafficFlowTransmit) Clone() (StateTrafficFlowTransmit, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateTrafficFlowTransmit()
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

// StateTrafficFlowTransmit is provides state control of flow transmission.
type StateTrafficFlowTransmit interface {
	Validation
	// msg marshals StateTrafficFlowTransmit to protobuf object *otg.StateTrafficFlowTransmit
	// and doesn't set defaults
	msg() *otg.StateTrafficFlowTransmit
	// setMsg unmarshals StateTrafficFlowTransmit from protobuf object *otg.StateTrafficFlowTransmit
	// and doesn't set defaults
	setMsg(*otg.StateTrafficFlowTransmit) StateTrafficFlowTransmit
	// provides marshal interface
	Marshal() marshalStateTrafficFlowTransmit
	// provides unmarshal interface
	Unmarshal() unMarshalStateTrafficFlowTransmit
	// validate validates StateTrafficFlowTransmit
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateTrafficFlowTransmit, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// FlowNames returns []string, set in StateTrafficFlowTransmit.
	FlowNames() []string
	// SetFlowNames assigns []string provided by user to StateTrafficFlowTransmit
	SetFlowNames(value []string) StateTrafficFlowTransmit
	// State returns StateTrafficFlowTransmitStateEnum, set in StateTrafficFlowTransmit
	State() StateTrafficFlowTransmitStateEnum
	// SetState assigns StateTrafficFlowTransmitStateEnum provided by user to StateTrafficFlowTransmit
	SetState(value StateTrafficFlowTransmitStateEnum) StateTrafficFlowTransmit
}

// The names of flows to which the transmit state will be applied to. If the list of flow_names is empty or null the state will be applied to all configured flows.
// If the list is not empty any flow that is not included in the list of flow_names MUST be ignored and not included in the state change.
//
// x-constraint:
// - /components/schemas/Flow/properties/name
//
// x-constraint:
// - /components/schemas/Flow/properties/name
//
// FlowNames returns a []string
func (obj *stateTrafficFlowTransmit) FlowNames() []string {
	if obj.obj.FlowNames == nil {
		obj.obj.FlowNames = make([]string, 0)
	}
	return obj.obj.FlowNames
}

// The names of flows to which the transmit state will be applied to. If the list of flow_names is empty or null the state will be applied to all configured flows.
// If the list is not empty any flow that is not included in the list of flow_names MUST be ignored and not included in the state change.
//
// x-constraint:
// - /components/schemas/Flow/properties/name
//
// x-constraint:
// - /components/schemas/Flow/properties/name
//
// SetFlowNames sets the []string value in the StateTrafficFlowTransmit object
func (obj *stateTrafficFlowTransmit) SetFlowNames(value []string) StateTrafficFlowTransmit {

	if obj.obj.FlowNames == nil {
		obj.obj.FlowNames = make([]string, 0)
	}
	obj.obj.FlowNames = value

	return obj
}

type StateTrafficFlowTransmitStateEnum string

// Enum of State on StateTrafficFlowTransmit
var StateTrafficFlowTransmitState = struct {
	START  StateTrafficFlowTransmitStateEnum
	STOP   StateTrafficFlowTransmitStateEnum
	PAUSE  StateTrafficFlowTransmitStateEnum
	RESUME StateTrafficFlowTransmitStateEnum
}{
	START:  StateTrafficFlowTransmitStateEnum("start"),
	STOP:   StateTrafficFlowTransmitStateEnum("stop"),
	PAUSE:  StateTrafficFlowTransmitStateEnum("pause"),
	RESUME: StateTrafficFlowTransmitStateEnum("resume"),
}

func (obj *stateTrafficFlowTransmit) State() StateTrafficFlowTransmitStateEnum {
	return StateTrafficFlowTransmitStateEnum(obj.obj.State.Enum().String())
}

func (obj *stateTrafficFlowTransmit) SetState(value StateTrafficFlowTransmitStateEnum) StateTrafficFlowTransmit {
	intValue, ok := otg.StateTrafficFlowTransmit_State_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateTrafficFlowTransmitStateEnum", string(value)))
		return obj
	}
	enumValue := otg.StateTrafficFlowTransmit_State_Enum(intValue)
	obj.obj.State = &enumValue

	return obj
}

func (obj *stateTrafficFlowTransmit) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// State is required
	if obj.obj.State == nil {
		vObj.validationErrors = append(vObj.validationErrors, "State is required field on interface StateTrafficFlowTransmit")
	}
}

func (obj *stateTrafficFlowTransmit) setDefault() {

}

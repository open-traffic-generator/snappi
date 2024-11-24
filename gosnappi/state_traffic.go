package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateTraffic *****
type stateTraffic struct {
	validation
	obj                *otg.StateTraffic
	marshaller         marshalStateTraffic
	unMarshaller       unMarshalStateTraffic
	flowTransmitHolder StateTrafficFlowTransmit
}

func NewStateTraffic() StateTraffic {
	obj := stateTraffic{obj: &otg.StateTraffic{}}
	obj.setDefault()
	return &obj
}

func (obj *stateTraffic) msg() *otg.StateTraffic {
	return obj.obj
}

func (obj *stateTraffic) setMsg(msg *otg.StateTraffic) StateTraffic {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateTraffic struct {
	obj *stateTraffic
}

type marshalStateTraffic interface {
	// ToProto marshals StateTraffic to protobuf object *otg.StateTraffic
	ToProto() (*otg.StateTraffic, error)
	// ToPbText marshals StateTraffic to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateTraffic to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateTraffic to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals StateTraffic to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalstateTraffic struct {
	obj *stateTraffic
}

type unMarshalStateTraffic interface {
	// FromProto unmarshals StateTraffic from protobuf object *otg.StateTraffic
	FromProto(msg *otg.StateTraffic) (StateTraffic, error)
	// FromPbText unmarshals StateTraffic from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateTraffic from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateTraffic from JSON text
	FromJson(value string) error
}

func (obj *stateTraffic) Marshal() marshalStateTraffic {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateTraffic{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateTraffic) Unmarshal() unMarshalStateTraffic {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateTraffic{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateTraffic) ToProto() (*otg.StateTraffic, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateTraffic) FromProto(msg *otg.StateTraffic) (StateTraffic, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateTraffic) ToPbText() (string, error) {
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

func (m *unMarshalstateTraffic) FromPbText(value string) error {
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

func (m *marshalstateTraffic) ToYaml() (string, error) {
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

func (m *unMarshalstateTraffic) FromYaml(value string) error {
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

func (m *marshalstateTraffic) ToJsonRaw() (string, error) {
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

func (m *marshalstateTraffic) ToJson() (string, error) {
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

func (m *unMarshalstateTraffic) FromJson(value string) error {
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

func (obj *stateTraffic) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateTraffic) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateTraffic) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateTraffic) Clone() (StateTraffic, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateTraffic()
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

func (obj *stateTraffic) setNil() {
	obj.flowTransmitHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// StateTraffic is states associated with configured flows
type StateTraffic interface {
	Validation
	// msg marshals StateTraffic to protobuf object *otg.StateTraffic
	// and doesn't set defaults
	msg() *otg.StateTraffic
	// setMsg unmarshals StateTraffic from protobuf object *otg.StateTraffic
	// and doesn't set defaults
	setMsg(*otg.StateTraffic) StateTraffic
	// provides marshal interface
	Marshal() marshalStateTraffic
	// provides unmarshal interface
	Unmarshal() unMarshalStateTraffic
	// validate validates StateTraffic
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateTraffic, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns StateTrafficChoiceEnum, set in StateTraffic
	Choice() StateTrafficChoiceEnum
	// setChoice assigns StateTrafficChoiceEnum provided by user to StateTraffic
	setChoice(value StateTrafficChoiceEnum) StateTraffic
	// FlowTransmit returns StateTrafficFlowTransmit, set in StateTraffic.
	// StateTrafficFlowTransmit is provides state control of flow transmission.
	FlowTransmit() StateTrafficFlowTransmit
	// SetFlowTransmit assigns StateTrafficFlowTransmit provided by user to StateTraffic.
	// StateTrafficFlowTransmit is provides state control of flow transmission.
	SetFlowTransmit(value StateTrafficFlowTransmit) StateTraffic
	// HasFlowTransmit checks if FlowTransmit has been set in StateTraffic
	HasFlowTransmit() bool
	setNil()
}

type StateTrafficChoiceEnum string

// Enum of Choice on StateTraffic
var StateTrafficChoice = struct {
	FLOW_TRANSMIT StateTrafficChoiceEnum
}{
	FLOW_TRANSMIT: StateTrafficChoiceEnum("flow_transmit"),
}

func (obj *stateTraffic) Choice() StateTrafficChoiceEnum {
	return StateTrafficChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *stateTraffic) setChoice(value StateTrafficChoiceEnum) StateTraffic {
	intValue, ok := otg.StateTraffic_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateTrafficChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.StateTraffic_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.FlowTransmit = nil
	obj.flowTransmitHolder = nil

	if value == StateTrafficChoice.FLOW_TRANSMIT {
		obj.obj.FlowTransmit = NewStateTrafficFlowTransmit().msg()
	}

	return obj
}

// description is TBD
// FlowTransmit returns a StateTrafficFlowTransmit
func (obj *stateTraffic) FlowTransmit() StateTrafficFlowTransmit {
	if obj.obj.FlowTransmit == nil {
		obj.setChoice(StateTrafficChoice.FLOW_TRANSMIT)
	}
	if obj.flowTransmitHolder == nil {
		obj.flowTransmitHolder = &stateTrafficFlowTransmit{obj: obj.obj.FlowTransmit}
	}
	return obj.flowTransmitHolder
}

// description is TBD
// FlowTransmit returns a StateTrafficFlowTransmit
func (obj *stateTraffic) HasFlowTransmit() bool {
	return obj.obj.FlowTransmit != nil
}

// description is TBD
// SetFlowTransmit sets the StateTrafficFlowTransmit value in the StateTraffic object
func (obj *stateTraffic) SetFlowTransmit(value StateTrafficFlowTransmit) StateTraffic {
	obj.setChoice(StateTrafficChoice.FLOW_TRANSMIT)
	obj.flowTransmitHolder = nil
	obj.obj.FlowTransmit = value.msg()

	return obj
}

func (obj *stateTraffic) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface StateTraffic")
	}

	if obj.obj.FlowTransmit != nil {

		obj.FlowTransmit().validateObj(vObj, set_default)
	}

}

func (obj *stateTraffic) setDefault() {
	var choices_set int = 0
	var choice StateTrafficChoiceEnum

	if obj.obj.FlowTransmit != nil {
		choices_set += 1
		choice = StateTrafficChoice.FLOW_TRANSMIT
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in StateTraffic")
			}
		} else {
			intVal := otg.StateTraffic_Choice_Enum_value[string(choice)]
			enumValue := otg.StateTraffic_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

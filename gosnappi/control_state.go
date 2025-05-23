package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ControlState *****
type controlState struct {
	validation
	obj            *otg.ControlState
	marshaller     marshalControlState
	unMarshaller   unMarshalControlState
	portHolder     StatePort
	protocolHolder StateProtocol
	trafficHolder  StateTraffic
}

func NewControlState() ControlState {
	obj := controlState{obj: &otg.ControlState{}}
	obj.setDefault()
	return &obj
}

func (obj *controlState) msg() *otg.ControlState {
	return obj.obj
}

func (obj *controlState) setMsg(msg *otg.ControlState) ControlState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcontrolState struct {
	obj *controlState
}

type marshalControlState interface {
	// ToProto marshals ControlState to protobuf object *otg.ControlState
	ToProto() (*otg.ControlState, error)
	// ToPbText marshals ControlState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ControlState to YAML text
	ToYaml() (string, error)
	// ToJson marshals ControlState to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals ControlState to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalcontrolState struct {
	obj *controlState
}

type unMarshalControlState interface {
	// FromProto unmarshals ControlState from protobuf object *otg.ControlState
	FromProto(msg *otg.ControlState) (ControlState, error)
	// FromPbText unmarshals ControlState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ControlState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ControlState from JSON text
	FromJson(value string) error
}

func (obj *controlState) Marshal() marshalControlState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcontrolState{obj: obj}
	}
	return obj.marshaller
}

func (obj *controlState) Unmarshal() unMarshalControlState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcontrolState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcontrolState) ToProto() (*otg.ControlState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcontrolState) FromProto(msg *otg.ControlState) (ControlState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcontrolState) ToPbText() (string, error) {
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

func (m *unMarshalcontrolState) FromPbText(value string) error {
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

func (m *marshalcontrolState) ToYaml() (string, error) {
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

func (m *unMarshalcontrolState) FromYaml(value string) error {
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

func (m *marshalcontrolState) ToJsonRaw() (string, error) {
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

func (m *marshalcontrolState) ToJson() (string, error) {
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

func (m *unMarshalcontrolState) FromJson(value string) error {
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

func (obj *controlState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *controlState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *controlState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *controlState) Clone() (ControlState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewControlState()
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

func (obj *controlState) setNil() {
	obj.portHolder = nil
	obj.protocolHolder = nil
	obj.trafficHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ControlState is request for setting operational state of configured resources.
type ControlState interface {
	Validation
	// msg marshals ControlState to protobuf object *otg.ControlState
	// and doesn't set defaults
	msg() *otg.ControlState
	// setMsg unmarshals ControlState from protobuf object *otg.ControlState
	// and doesn't set defaults
	setMsg(*otg.ControlState) ControlState
	// provides marshal interface
	Marshal() marshalControlState
	// provides unmarshal interface
	Unmarshal() unMarshalControlState
	// validate validates ControlState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ControlState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ControlStateChoiceEnum, set in ControlState
	Choice() ControlStateChoiceEnum
	// setChoice assigns ControlStateChoiceEnum provided by user to ControlState
	setChoice(value ControlStateChoiceEnum) ControlState
	// Port returns StatePort, set in ControlState.
	// StatePort is states associated with configured ports.
	Port() StatePort
	// SetPort assigns StatePort provided by user to ControlState.
	// StatePort is states associated with configured ports.
	SetPort(value StatePort) ControlState
	// HasPort checks if Port has been set in ControlState
	HasPort() bool
	// Protocol returns StateProtocol, set in ControlState.
	// StateProtocol is states associated with protocols on configured resources.
	Protocol() StateProtocol
	// SetProtocol assigns StateProtocol provided by user to ControlState.
	// StateProtocol is states associated with protocols on configured resources.
	SetProtocol(value StateProtocol) ControlState
	// HasProtocol checks if Protocol has been set in ControlState
	HasProtocol() bool
	// Traffic returns StateTraffic, set in ControlState.
	// StateTraffic is states associated with configured flows
	Traffic() StateTraffic
	// SetTraffic assigns StateTraffic provided by user to ControlState.
	// StateTraffic is states associated with configured flows
	SetTraffic(value StateTraffic) ControlState
	// HasTraffic checks if Traffic has been set in ControlState
	HasTraffic() bool
	setNil()
}

type ControlStateChoiceEnum string

// Enum of Choice on ControlState
var ControlStateChoice = struct {
	PORT     ControlStateChoiceEnum
	PROTOCOL ControlStateChoiceEnum
	TRAFFIC  ControlStateChoiceEnum
}{
	PORT:     ControlStateChoiceEnum("port"),
	PROTOCOL: ControlStateChoiceEnum("protocol"),
	TRAFFIC:  ControlStateChoiceEnum("traffic"),
}

func (obj *controlState) Choice() ControlStateChoiceEnum {
	return ControlStateChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *controlState) setChoice(value ControlStateChoiceEnum) ControlState {
	intValue, ok := otg.ControlState_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ControlStateChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ControlState_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Traffic = nil
	obj.trafficHolder = nil
	obj.obj.Protocol = nil
	obj.protocolHolder = nil
	obj.obj.Port = nil
	obj.portHolder = nil

	if value == ControlStateChoice.PORT {
		obj.obj.Port = NewStatePort().msg()
	}

	if value == ControlStateChoice.PROTOCOL {
		obj.obj.Protocol = NewStateProtocol().msg()
	}

	if value == ControlStateChoice.TRAFFIC {
		obj.obj.Traffic = NewStateTraffic().msg()
	}

	return obj
}

// description is TBD
// Port returns a StatePort
func (obj *controlState) Port() StatePort {
	if obj.obj.Port == nil {
		obj.setChoice(ControlStateChoice.PORT)
	}
	if obj.portHolder == nil {
		obj.portHolder = &statePort{obj: obj.obj.Port}
	}
	return obj.portHolder
}

// description is TBD
// Port returns a StatePort
func (obj *controlState) HasPort() bool {
	return obj.obj.Port != nil
}

// description is TBD
// SetPort sets the StatePort value in the ControlState object
func (obj *controlState) SetPort(value StatePort) ControlState {
	obj.setChoice(ControlStateChoice.PORT)
	obj.portHolder = nil
	obj.obj.Port = value.msg()

	return obj
}

// description is TBD
// Protocol returns a StateProtocol
func (obj *controlState) Protocol() StateProtocol {
	if obj.obj.Protocol == nil {
		obj.setChoice(ControlStateChoice.PROTOCOL)
	}
	if obj.protocolHolder == nil {
		obj.protocolHolder = &stateProtocol{obj: obj.obj.Protocol}
	}
	return obj.protocolHolder
}

// description is TBD
// Protocol returns a StateProtocol
func (obj *controlState) HasProtocol() bool {
	return obj.obj.Protocol != nil
}

// description is TBD
// SetProtocol sets the StateProtocol value in the ControlState object
func (obj *controlState) SetProtocol(value StateProtocol) ControlState {
	obj.setChoice(ControlStateChoice.PROTOCOL)
	obj.protocolHolder = nil
	obj.obj.Protocol = value.msg()

	return obj
}

// description is TBD
// Traffic returns a StateTraffic
func (obj *controlState) Traffic() StateTraffic {
	if obj.obj.Traffic == nil {
		obj.setChoice(ControlStateChoice.TRAFFIC)
	}
	if obj.trafficHolder == nil {
		obj.trafficHolder = &stateTraffic{obj: obj.obj.Traffic}
	}
	return obj.trafficHolder
}

// description is TBD
// Traffic returns a StateTraffic
func (obj *controlState) HasTraffic() bool {
	return obj.obj.Traffic != nil
}

// description is TBD
// SetTraffic sets the StateTraffic value in the ControlState object
func (obj *controlState) SetTraffic(value StateTraffic) ControlState {
	obj.setChoice(ControlStateChoice.TRAFFIC)
	obj.trafficHolder = nil
	obj.obj.Traffic = value.msg()

	return obj
}

func (obj *controlState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface ControlState")
	}

	if obj.obj.Port != nil {

		obj.Port().validateObj(vObj, set_default)
	}

	if obj.obj.Protocol != nil {

		obj.Protocol().validateObj(vObj, set_default)
	}

	if obj.obj.Traffic != nil {

		obj.Traffic().validateObj(vObj, set_default)
	}

}

func (obj *controlState) setDefault() {
	var choices_set int = 0
	var choice ControlStateChoiceEnum

	if obj.obj.Port != nil {
		choices_set += 1
		choice = ControlStateChoice.PORT
	}

	if obj.obj.Protocol != nil {
		choices_set += 1
		choice = ControlStateChoice.PROTOCOL
	}

	if obj.obj.Traffic != nil {
		choices_set += 1
		choice = ControlStateChoice.TRAFFIC
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ControlState")
			}
		} else {
			intVal := otg.ControlState_Choice_Enum_value[string(choice)]
			enumValue := otg.ControlState_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

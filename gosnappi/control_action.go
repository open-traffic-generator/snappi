package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ControlAction *****
type controlAction struct {
	validation
	obj            *otg.ControlAction
	marshaller     marshalControlAction
	unMarshaller   unMarshalControlAction
	protocolHolder ActionProtocol
}

func NewControlAction() ControlAction {
	obj := controlAction{obj: &otg.ControlAction{}}
	obj.setDefault()
	return &obj
}

func (obj *controlAction) msg() *otg.ControlAction {
	return obj.obj
}

func (obj *controlAction) setMsg(msg *otg.ControlAction) ControlAction {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcontrolAction struct {
	obj *controlAction
}

type marshalControlAction interface {
	// ToProto marshals ControlAction to protobuf object *otg.ControlAction
	ToProto() (*otg.ControlAction, error)
	// ToPbText marshals ControlAction to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ControlAction to YAML text
	ToYaml() (string, error)
	// ToJson marshals ControlAction to JSON text
	ToJson() (string, error)
}

type unMarshalcontrolAction struct {
	obj *controlAction
}

type unMarshalControlAction interface {
	// FromProto unmarshals ControlAction from protobuf object *otg.ControlAction
	FromProto(msg *otg.ControlAction) (ControlAction, error)
	// FromPbText unmarshals ControlAction from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ControlAction from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ControlAction from JSON text
	FromJson(value string) error
}

func (obj *controlAction) Marshal() marshalControlAction {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcontrolAction{obj: obj}
	}
	return obj.marshaller
}

func (obj *controlAction) Unmarshal() unMarshalControlAction {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcontrolAction{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcontrolAction) ToProto() (*otg.ControlAction, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcontrolAction) FromProto(msg *otg.ControlAction) (ControlAction, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcontrolAction) ToPbText() (string, error) {
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

func (m *unMarshalcontrolAction) FromPbText(value string) error {
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

func (m *marshalcontrolAction) ToYaml() (string, error) {
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

func (m *unMarshalcontrolAction) FromYaml(value string) error {
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

func (m *marshalcontrolAction) ToJson() (string, error) {
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

func (m *unMarshalcontrolAction) FromJson(value string) error {
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

func (obj *controlAction) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *controlAction) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *controlAction) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *controlAction) Clone() (ControlAction, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewControlAction()
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

func (obj *controlAction) setNil() {
	obj.protocolHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ControlAction is request for triggering action against configured resources.
type ControlAction interface {
	Validation
	// msg marshals ControlAction to protobuf object *otg.ControlAction
	// and doesn't set defaults
	msg() *otg.ControlAction
	// setMsg unmarshals ControlAction from protobuf object *otg.ControlAction
	// and doesn't set defaults
	setMsg(*otg.ControlAction) ControlAction
	// provides marshal interface
	Marshal() marshalControlAction
	// provides unmarshal interface
	Unmarshal() unMarshalControlAction
	// validate validates ControlAction
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ControlAction, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ControlActionChoiceEnum, set in ControlAction
	Choice() ControlActionChoiceEnum
	// setChoice assigns ControlActionChoiceEnum provided by user to ControlAction
	setChoice(value ControlActionChoiceEnum) ControlAction
	// Protocol returns ActionProtocol, set in ControlAction.
	// ActionProtocol is actions associated with protocols on configured resources.
	Protocol() ActionProtocol
	// SetProtocol assigns ActionProtocol provided by user to ControlAction.
	// ActionProtocol is actions associated with protocols on configured resources.
	SetProtocol(value ActionProtocol) ControlAction
	// HasProtocol checks if Protocol has been set in ControlAction
	HasProtocol() bool
	setNil()
}

type ControlActionChoiceEnum string

// Enum of Choice on ControlAction
var ControlActionChoice = struct {
	PROTOCOL ControlActionChoiceEnum
}{
	PROTOCOL: ControlActionChoiceEnum("protocol"),
}

func (obj *controlAction) Choice() ControlActionChoiceEnum {
	return ControlActionChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *controlAction) setChoice(value ControlActionChoiceEnum) ControlAction {
	intValue, ok := otg.ControlAction_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ControlActionChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ControlAction_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Protocol = nil
	obj.protocolHolder = nil

	if value == ControlActionChoice.PROTOCOL {
		obj.obj.Protocol = NewActionProtocol().msg()
	}

	return obj
}

// description is TBD
// Protocol returns a ActionProtocol
func (obj *controlAction) Protocol() ActionProtocol {
	if obj.obj.Protocol == nil {
		obj.setChoice(ControlActionChoice.PROTOCOL)
	}
	if obj.protocolHolder == nil {
		obj.protocolHolder = &actionProtocol{obj: obj.obj.Protocol}
	}
	return obj.protocolHolder
}

// description is TBD
// Protocol returns a ActionProtocol
func (obj *controlAction) HasProtocol() bool {
	return obj.obj.Protocol != nil
}

// description is TBD
// SetProtocol sets the ActionProtocol value in the ControlAction object
func (obj *controlAction) SetProtocol(value ActionProtocol) ControlAction {
	obj.setChoice(ControlActionChoice.PROTOCOL)
	obj.protocolHolder = nil
	obj.obj.Protocol = value.msg()

	return obj
}

func (obj *controlAction) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface ControlAction")
	}

	if obj.obj.Protocol != nil {

		obj.Protocol().validateObj(vObj, set_default)
	}

}

func (obj *controlAction) setDefault() {

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionPort *****
type actionPort struct {
	validation
	obj          *otg.ActionPort
	marshaller   marshalActionPort
	unMarshaller unMarshalActionPort
	rebootHolder ActionPortReboot
}

func NewActionPort() ActionPort {
	obj := actionPort{obj: &otg.ActionPort{}}
	obj.setDefault()
	return &obj
}

func (obj *actionPort) msg() *otg.ActionPort {
	return obj.obj
}

func (obj *actionPort) setMsg(msg *otg.ActionPort) ActionPort {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionPort struct {
	obj *actionPort
}

type marshalActionPort interface {
	// ToProto marshals ActionPort to protobuf object *otg.ActionPort
	ToProto() (*otg.ActionPort, error)
	// ToPbText marshals ActionPort to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionPort to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionPort to JSON text
	ToJson() (string, error)
}

type unMarshalactionPort struct {
	obj *actionPort
}

type unMarshalActionPort interface {
	// FromProto unmarshals ActionPort from protobuf object *otg.ActionPort
	FromProto(msg *otg.ActionPort) (ActionPort, error)
	// FromPbText unmarshals ActionPort from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionPort from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionPort from JSON text
	FromJson(value string) error
}

func (obj *actionPort) Marshal() marshalActionPort {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionPort{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionPort) Unmarshal() unMarshalActionPort {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionPort{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionPort) ToProto() (*otg.ActionPort, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionPort) FromProto(msg *otg.ActionPort) (ActionPort, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionPort) ToPbText() (string, error) {
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

func (m *unMarshalactionPort) FromPbText(value string) error {
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

func (m *marshalactionPort) ToYaml() (string, error) {
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

func (m *unMarshalactionPort) FromYaml(value string) error {
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

func (m *marshalactionPort) ToJson() (string, error) {
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

func (m *unMarshalactionPort) FromJson(value string) error {
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

func (obj *actionPort) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionPort) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionPort) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionPort) Clone() (ActionPort, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionPort()
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

func (obj *actionPort) setNil() {
	obj.rebootHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionPort is actions associated with ports on configured resources.
type ActionPort interface {
	Validation
	// msg marshals ActionPort to protobuf object *otg.ActionPort
	// and doesn't set defaults
	msg() *otg.ActionPort
	// setMsg unmarshals ActionPort from protobuf object *otg.ActionPort
	// and doesn't set defaults
	setMsg(*otg.ActionPort) ActionPort
	// provides marshal interface
	Marshal() marshalActionPort
	// provides unmarshal interface
	Unmarshal() unMarshalActionPort
	// validate validates ActionPort
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionPort, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ActionPortChoiceEnum, set in ActionPort
	Choice() ActionPortChoiceEnum
	// setChoice assigns ActionPortChoiceEnum provided by user to ActionPort
	setChoice(value ActionPortChoiceEnum) ActionPort
	// Reboot returns ActionPortReboot, set in ActionPort.
	// ActionPortReboot is resets the configured ports to initialized state.
	Reboot() ActionPortReboot
	// SetReboot assigns ActionPortReboot provided by user to ActionPort.
	// ActionPortReboot is resets the configured ports to initialized state.
	SetReboot(value ActionPortReboot) ActionPort
	// HasReboot checks if Reboot has been set in ActionPort
	HasReboot() bool
	setNil()
}

type ActionPortChoiceEnum string

// Enum of Choice on ActionPort
var ActionPortChoice = struct {
	REBOOT ActionPortChoiceEnum
}{
	REBOOT: ActionPortChoiceEnum("reboot"),
}

func (obj *actionPort) Choice() ActionPortChoiceEnum {
	return ActionPortChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *actionPort) setChoice(value ActionPortChoiceEnum) ActionPort {
	intValue, ok := otg.ActionPort_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ActionPortChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ActionPort_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Reboot = nil
	obj.rebootHolder = nil

	if value == ActionPortChoice.REBOOT {
		obj.obj.Reboot = NewActionPortReboot().msg()
	}

	return obj
}

// description is TBD
// Reboot returns a ActionPortReboot
func (obj *actionPort) Reboot() ActionPortReboot {
	if obj.obj.Reboot == nil {
		obj.setChoice(ActionPortChoice.REBOOT)
	}
	if obj.rebootHolder == nil {
		obj.rebootHolder = &actionPortReboot{obj: obj.obj.Reboot}
	}
	return obj.rebootHolder
}

// description is TBD
// Reboot returns a ActionPortReboot
func (obj *actionPort) HasReboot() bool {
	return obj.obj.Reboot != nil
}

// description is TBD
// SetReboot sets the ActionPortReboot value in the ActionPort object
func (obj *actionPort) SetReboot(value ActionPortReboot) ActionPort {
	obj.setChoice(ActionPortChoice.REBOOT)
	obj.rebootHolder = nil
	obj.obj.Reboot = value.msg()

	return obj
}

func (obj *actionPort) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface ActionPort")
	}

	if obj.obj.Reboot != nil {

		obj.Reboot().validateObj(vObj, set_default)
	}

}

func (obj *actionPort) setDefault() {
	var choices_set int = 0
	var choice ActionPortChoiceEnum

	if obj.obj.Reboot != nil {
		choices_set += 1
		choice = ActionPortChoice.REBOOT
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ActionPort")
			}
		} else {
			intVal := otg.ActionPort_Choice_Enum_value[string(choice)]
			enumValue := otg.ActionPort_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

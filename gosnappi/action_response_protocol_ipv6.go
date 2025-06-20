package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionResponseProtocolIpv6 *****
type actionResponseProtocolIpv6 struct {
	validation
	obj          *otg.ActionResponseProtocolIpv6
	marshaller   marshalActionResponseProtocolIpv6
	unMarshaller unMarshalActionResponseProtocolIpv6
	pingHolder   ActionResponseProtocolIpv6Ping
}

func NewActionResponseProtocolIpv6() ActionResponseProtocolIpv6 {
	obj := actionResponseProtocolIpv6{obj: &otg.ActionResponseProtocolIpv6{}}
	obj.setDefault()
	return &obj
}

func (obj *actionResponseProtocolIpv6) msg() *otg.ActionResponseProtocolIpv6 {
	return obj.obj
}

func (obj *actionResponseProtocolIpv6) setMsg(msg *otg.ActionResponseProtocolIpv6) ActionResponseProtocolIpv6 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionResponseProtocolIpv6 struct {
	obj *actionResponseProtocolIpv6
}

type marshalActionResponseProtocolIpv6 interface {
	// ToProto marshals ActionResponseProtocolIpv6 to protobuf object *otg.ActionResponseProtocolIpv6
	ToProto() (*otg.ActionResponseProtocolIpv6, error)
	// ToPbText marshals ActionResponseProtocolIpv6 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionResponseProtocolIpv6 to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionResponseProtocolIpv6 to JSON text
	ToJson() (string, error)
}

type unMarshalactionResponseProtocolIpv6 struct {
	obj *actionResponseProtocolIpv6
}

type unMarshalActionResponseProtocolIpv6 interface {
	// FromProto unmarshals ActionResponseProtocolIpv6 from protobuf object *otg.ActionResponseProtocolIpv6
	FromProto(msg *otg.ActionResponseProtocolIpv6) (ActionResponseProtocolIpv6, error)
	// FromPbText unmarshals ActionResponseProtocolIpv6 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionResponseProtocolIpv6 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionResponseProtocolIpv6 from JSON text
	FromJson(value string) error
}

func (obj *actionResponseProtocolIpv6) Marshal() marshalActionResponseProtocolIpv6 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionResponseProtocolIpv6{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionResponseProtocolIpv6) Unmarshal() unMarshalActionResponseProtocolIpv6 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionResponseProtocolIpv6{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionResponseProtocolIpv6) ToProto() (*otg.ActionResponseProtocolIpv6, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionResponseProtocolIpv6) FromProto(msg *otg.ActionResponseProtocolIpv6) (ActionResponseProtocolIpv6, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionResponseProtocolIpv6) ToPbText() (string, error) {
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

func (m *unMarshalactionResponseProtocolIpv6) FromPbText(value string) error {
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

func (m *marshalactionResponseProtocolIpv6) ToYaml() (string, error) {
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

func (m *unMarshalactionResponseProtocolIpv6) FromYaml(value string) error {
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

func (m *marshalactionResponseProtocolIpv6) ToJson() (string, error) {
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

func (m *unMarshalactionResponseProtocolIpv6) FromJson(value string) error {
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

func (obj *actionResponseProtocolIpv6) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionResponseProtocolIpv6) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionResponseProtocolIpv6) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionResponseProtocolIpv6) Clone() (ActionResponseProtocolIpv6, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionResponseProtocolIpv6()
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

func (obj *actionResponseProtocolIpv6) setNil() {
	obj.pingHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionResponseProtocolIpv6 is response for actions associated with IPv6 on configured resources.
type ActionResponseProtocolIpv6 interface {
	Validation
	// msg marshals ActionResponseProtocolIpv6 to protobuf object *otg.ActionResponseProtocolIpv6
	// and doesn't set defaults
	msg() *otg.ActionResponseProtocolIpv6
	// setMsg unmarshals ActionResponseProtocolIpv6 from protobuf object *otg.ActionResponseProtocolIpv6
	// and doesn't set defaults
	setMsg(*otg.ActionResponseProtocolIpv6) ActionResponseProtocolIpv6
	// provides marshal interface
	Marshal() marshalActionResponseProtocolIpv6
	// provides unmarshal interface
	Unmarshal() unMarshalActionResponseProtocolIpv6
	// validate validates ActionResponseProtocolIpv6
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionResponseProtocolIpv6, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ActionResponseProtocolIpv6ChoiceEnum, set in ActionResponseProtocolIpv6
	Choice() ActionResponseProtocolIpv6ChoiceEnum
	// setChoice assigns ActionResponseProtocolIpv6ChoiceEnum provided by user to ActionResponseProtocolIpv6
	setChoice(value ActionResponseProtocolIpv6ChoiceEnum) ActionResponseProtocolIpv6
	// Ping returns ActionResponseProtocolIpv6Ping, set in ActionResponseProtocolIpv6.
	// ActionResponseProtocolIpv6Ping is response for ping initiated between multiple source and destination pairs.
	Ping() ActionResponseProtocolIpv6Ping
	// SetPing assigns ActionResponseProtocolIpv6Ping provided by user to ActionResponseProtocolIpv6.
	// ActionResponseProtocolIpv6Ping is response for ping initiated between multiple source and destination pairs.
	SetPing(value ActionResponseProtocolIpv6Ping) ActionResponseProtocolIpv6
	// HasPing checks if Ping has been set in ActionResponseProtocolIpv6
	HasPing() bool
	setNil()
}

type ActionResponseProtocolIpv6ChoiceEnum string

// Enum of Choice on ActionResponseProtocolIpv6
var ActionResponseProtocolIpv6Choice = struct {
	PING ActionResponseProtocolIpv6ChoiceEnum
}{
	PING: ActionResponseProtocolIpv6ChoiceEnum("ping"),
}

func (obj *actionResponseProtocolIpv6) Choice() ActionResponseProtocolIpv6ChoiceEnum {
	return ActionResponseProtocolIpv6ChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *actionResponseProtocolIpv6) setChoice(value ActionResponseProtocolIpv6ChoiceEnum) ActionResponseProtocolIpv6 {
	intValue, ok := otg.ActionResponseProtocolIpv6_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ActionResponseProtocolIpv6ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ActionResponseProtocolIpv6_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Ping = nil
	obj.pingHolder = nil

	if value == ActionResponseProtocolIpv6Choice.PING {
		obj.obj.Ping = NewActionResponseProtocolIpv6Ping().msg()
	}

	return obj
}

// description is TBD
// Ping returns a ActionResponseProtocolIpv6Ping
func (obj *actionResponseProtocolIpv6) Ping() ActionResponseProtocolIpv6Ping {
	if obj.obj.Ping == nil {
		obj.setChoice(ActionResponseProtocolIpv6Choice.PING)
	}
	if obj.pingHolder == nil {
		obj.pingHolder = &actionResponseProtocolIpv6Ping{obj: obj.obj.Ping}
	}
	return obj.pingHolder
}

// description is TBD
// Ping returns a ActionResponseProtocolIpv6Ping
func (obj *actionResponseProtocolIpv6) HasPing() bool {
	return obj.obj.Ping != nil
}

// description is TBD
// SetPing sets the ActionResponseProtocolIpv6Ping value in the ActionResponseProtocolIpv6 object
func (obj *actionResponseProtocolIpv6) SetPing(value ActionResponseProtocolIpv6Ping) ActionResponseProtocolIpv6 {
	obj.setChoice(ActionResponseProtocolIpv6Choice.PING)
	obj.pingHolder = nil
	obj.obj.Ping = value.msg()

	return obj
}

func (obj *actionResponseProtocolIpv6) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface ActionResponseProtocolIpv6")
	}

	if obj.obj.Ping != nil {

		obj.Ping().validateObj(vObj, set_default)
	}

}

func (obj *actionResponseProtocolIpv6) setDefault() {
	var choices_set int = 0
	var choice ActionResponseProtocolIpv6ChoiceEnum

	if obj.obj.Ping != nil {
		choices_set += 1
		choice = ActionResponseProtocolIpv6Choice.PING
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ActionResponseProtocolIpv6")
			}
		} else {
			intVal := otg.ActionResponseProtocolIpv6_Choice_Enum_value[string(choice)]
			enumValue := otg.ActionResponseProtocolIpv6_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

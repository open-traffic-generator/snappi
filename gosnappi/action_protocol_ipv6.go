package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolIpv6 *****
type actionProtocolIpv6 struct {
	validation
	obj          *otg.ActionProtocolIpv6
	marshaller   marshalActionProtocolIpv6
	unMarshaller unMarshalActionProtocolIpv6
	pingHolder   ActionProtocolIpv6Ping
}

func NewActionProtocolIpv6() ActionProtocolIpv6 {
	obj := actionProtocolIpv6{obj: &otg.ActionProtocolIpv6{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolIpv6) msg() *otg.ActionProtocolIpv6 {
	return obj.obj
}

func (obj *actionProtocolIpv6) setMsg(msg *otg.ActionProtocolIpv6) ActionProtocolIpv6 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolIpv6 struct {
	obj *actionProtocolIpv6
}

type marshalActionProtocolIpv6 interface {
	// ToProto marshals ActionProtocolIpv6 to protobuf object *otg.ActionProtocolIpv6
	ToProto() (*otg.ActionProtocolIpv6, error)
	// ToPbText marshals ActionProtocolIpv6 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolIpv6 to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolIpv6 to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolIpv6 struct {
	obj *actionProtocolIpv6
}

type unMarshalActionProtocolIpv6 interface {
	// FromProto unmarshals ActionProtocolIpv6 from protobuf object *otg.ActionProtocolIpv6
	FromProto(msg *otg.ActionProtocolIpv6) (ActionProtocolIpv6, error)
	// FromPbText unmarshals ActionProtocolIpv6 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolIpv6 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolIpv6 from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolIpv6) Marshal() marshalActionProtocolIpv6 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolIpv6{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolIpv6) Unmarshal() unMarshalActionProtocolIpv6 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolIpv6{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolIpv6) ToProto() (*otg.ActionProtocolIpv6, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolIpv6) FromProto(msg *otg.ActionProtocolIpv6) (ActionProtocolIpv6, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolIpv6) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolIpv6) FromPbText(value string) error {
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

func (m *marshalactionProtocolIpv6) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolIpv6) FromYaml(value string) error {
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

func (m *marshalactionProtocolIpv6) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolIpv6) FromJson(value string) error {
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

func (obj *actionProtocolIpv6) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolIpv6) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolIpv6) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolIpv6) Clone() (ActionProtocolIpv6, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolIpv6()
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

func (obj *actionProtocolIpv6) setNil() {
	obj.pingHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionProtocolIpv6 is actions associated with IPv6 on configured resources.
type ActionProtocolIpv6 interface {
	Validation
	// msg marshals ActionProtocolIpv6 to protobuf object *otg.ActionProtocolIpv6
	// and doesn't set defaults
	msg() *otg.ActionProtocolIpv6
	// setMsg unmarshals ActionProtocolIpv6 from protobuf object *otg.ActionProtocolIpv6
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolIpv6) ActionProtocolIpv6
	// provides marshal interface
	Marshal() marshalActionProtocolIpv6
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolIpv6
	// validate validates ActionProtocolIpv6
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolIpv6, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ActionProtocolIpv6ChoiceEnum, set in ActionProtocolIpv6
	Choice() ActionProtocolIpv6ChoiceEnum
	// setChoice assigns ActionProtocolIpv6ChoiceEnum provided by user to ActionProtocolIpv6
	setChoice(value ActionProtocolIpv6ChoiceEnum) ActionProtocolIpv6
	// Ping returns ActionProtocolIpv6Ping, set in ActionProtocolIpv6.
	// ActionProtocolIpv6Ping is request for initiating ping between multiple source and destination pairs.
	Ping() ActionProtocolIpv6Ping
	// SetPing assigns ActionProtocolIpv6Ping provided by user to ActionProtocolIpv6.
	// ActionProtocolIpv6Ping is request for initiating ping between multiple source and destination pairs.
	SetPing(value ActionProtocolIpv6Ping) ActionProtocolIpv6
	// HasPing checks if Ping has been set in ActionProtocolIpv6
	HasPing() bool
	setNil()
}

type ActionProtocolIpv6ChoiceEnum string

// Enum of Choice on ActionProtocolIpv6
var ActionProtocolIpv6Choice = struct {
	PING ActionProtocolIpv6ChoiceEnum
}{
	PING: ActionProtocolIpv6ChoiceEnum("ping"),
}

func (obj *actionProtocolIpv6) Choice() ActionProtocolIpv6ChoiceEnum {
	return ActionProtocolIpv6ChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *actionProtocolIpv6) setChoice(value ActionProtocolIpv6ChoiceEnum) ActionProtocolIpv6 {
	intValue, ok := otg.ActionProtocolIpv6_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ActionProtocolIpv6ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ActionProtocolIpv6_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Ping = nil
	obj.pingHolder = nil

	if value == ActionProtocolIpv6Choice.PING {
		obj.obj.Ping = NewActionProtocolIpv6Ping().msg()
	}

	return obj
}

// description is TBD
// Ping returns a ActionProtocolIpv6Ping
func (obj *actionProtocolIpv6) Ping() ActionProtocolIpv6Ping {
	if obj.obj.Ping == nil {
		obj.setChoice(ActionProtocolIpv6Choice.PING)
	}
	if obj.pingHolder == nil {
		obj.pingHolder = &actionProtocolIpv6Ping{obj: obj.obj.Ping}
	}
	return obj.pingHolder
}

// description is TBD
// Ping returns a ActionProtocolIpv6Ping
func (obj *actionProtocolIpv6) HasPing() bool {
	return obj.obj.Ping != nil
}

// description is TBD
// SetPing sets the ActionProtocolIpv6Ping value in the ActionProtocolIpv6 object
func (obj *actionProtocolIpv6) SetPing(value ActionProtocolIpv6Ping) ActionProtocolIpv6 {
	obj.setChoice(ActionProtocolIpv6Choice.PING)
	obj.pingHolder = nil
	obj.obj.Ping = value.msg()

	return obj
}

func (obj *actionProtocolIpv6) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface ActionProtocolIpv6")
	}

	if obj.obj.Ping != nil {

		obj.Ping().validateObj(vObj, set_default)
	}

}

func (obj *actionProtocolIpv6) setDefault() {
	var choices_set int = 0
	var choice ActionProtocolIpv6ChoiceEnum

	if obj.obj.Ping != nil {
		choices_set += 1
		choice = ActionProtocolIpv6Choice.PING
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ActionProtocolIpv6")
			}
		} else {
			intVal := otg.ActionProtocolIpv6_Choice_Enum_value[string(choice)]
			enumValue := otg.ActionProtocolIpv6_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolIpv4 *****
type actionProtocolIpv4 struct {
	validation
	obj          *otg.ActionProtocolIpv4
	marshaller   marshalActionProtocolIpv4
	unMarshaller unMarshalActionProtocolIpv4
	pingHolder   ActionProtocolIpv4Ping
}

func NewActionProtocolIpv4() ActionProtocolIpv4 {
	obj := actionProtocolIpv4{obj: &otg.ActionProtocolIpv4{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolIpv4) msg() *otg.ActionProtocolIpv4 {
	return obj.obj
}

func (obj *actionProtocolIpv4) setMsg(msg *otg.ActionProtocolIpv4) ActionProtocolIpv4 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolIpv4 struct {
	obj *actionProtocolIpv4
}

type marshalActionProtocolIpv4 interface {
	// ToProto marshals ActionProtocolIpv4 to protobuf object *otg.ActionProtocolIpv4
	ToProto() (*otg.ActionProtocolIpv4, error)
	// ToPbText marshals ActionProtocolIpv4 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolIpv4 to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolIpv4 to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals ActionProtocolIpv4 to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalactionProtocolIpv4 struct {
	obj *actionProtocolIpv4
}

type unMarshalActionProtocolIpv4 interface {
	// FromProto unmarshals ActionProtocolIpv4 from protobuf object *otg.ActionProtocolIpv4
	FromProto(msg *otg.ActionProtocolIpv4) (ActionProtocolIpv4, error)
	// FromPbText unmarshals ActionProtocolIpv4 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolIpv4 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolIpv4 from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolIpv4) Marshal() marshalActionProtocolIpv4 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolIpv4{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolIpv4) Unmarshal() unMarshalActionProtocolIpv4 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolIpv4{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolIpv4) ToProto() (*otg.ActionProtocolIpv4, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolIpv4) FromProto(msg *otg.ActionProtocolIpv4) (ActionProtocolIpv4, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolIpv4) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolIpv4) FromPbText(value string) error {
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

func (m *marshalactionProtocolIpv4) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolIpv4) FromYaml(value string) error {
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

func (m *marshalactionProtocolIpv4) ToJsonRaw() (string, error) {
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

func (m *marshalactionProtocolIpv4) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolIpv4) FromJson(value string) error {
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

func (obj *actionProtocolIpv4) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolIpv4) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolIpv4) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolIpv4) Clone() (ActionProtocolIpv4, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolIpv4()
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

func (obj *actionProtocolIpv4) setNil() {
	obj.pingHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionProtocolIpv4 is actions associated with IPv4 on configured resources.
type ActionProtocolIpv4 interface {
	Validation
	// msg marshals ActionProtocolIpv4 to protobuf object *otg.ActionProtocolIpv4
	// and doesn't set defaults
	msg() *otg.ActionProtocolIpv4
	// setMsg unmarshals ActionProtocolIpv4 from protobuf object *otg.ActionProtocolIpv4
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolIpv4) ActionProtocolIpv4
	// provides marshal interface
	Marshal() marshalActionProtocolIpv4
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolIpv4
	// validate validates ActionProtocolIpv4
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolIpv4, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ActionProtocolIpv4ChoiceEnum, set in ActionProtocolIpv4
	Choice() ActionProtocolIpv4ChoiceEnum
	// setChoice assigns ActionProtocolIpv4ChoiceEnum provided by user to ActionProtocolIpv4
	setChoice(value ActionProtocolIpv4ChoiceEnum) ActionProtocolIpv4
	// Ping returns ActionProtocolIpv4Ping, set in ActionProtocolIpv4.
	// ActionProtocolIpv4Ping is request for initiating ping between multiple source and destination pairs.
	Ping() ActionProtocolIpv4Ping
	// SetPing assigns ActionProtocolIpv4Ping provided by user to ActionProtocolIpv4.
	// ActionProtocolIpv4Ping is request for initiating ping between multiple source and destination pairs.
	SetPing(value ActionProtocolIpv4Ping) ActionProtocolIpv4
	// HasPing checks if Ping has been set in ActionProtocolIpv4
	HasPing() bool
	setNil()
}

type ActionProtocolIpv4ChoiceEnum string

// Enum of Choice on ActionProtocolIpv4
var ActionProtocolIpv4Choice = struct {
	PING ActionProtocolIpv4ChoiceEnum
}{
	PING: ActionProtocolIpv4ChoiceEnum("ping"),
}

func (obj *actionProtocolIpv4) Choice() ActionProtocolIpv4ChoiceEnum {
	return ActionProtocolIpv4ChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *actionProtocolIpv4) setChoice(value ActionProtocolIpv4ChoiceEnum) ActionProtocolIpv4 {
	intValue, ok := otg.ActionProtocolIpv4_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ActionProtocolIpv4ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ActionProtocolIpv4_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Ping = nil
	obj.pingHolder = nil

	if value == ActionProtocolIpv4Choice.PING {
		obj.obj.Ping = NewActionProtocolIpv4Ping().msg()
	}

	return obj
}

// description is TBD
// Ping returns a ActionProtocolIpv4Ping
func (obj *actionProtocolIpv4) Ping() ActionProtocolIpv4Ping {
	if obj.obj.Ping == nil {
		obj.setChoice(ActionProtocolIpv4Choice.PING)
	}
	if obj.pingHolder == nil {
		obj.pingHolder = &actionProtocolIpv4Ping{obj: obj.obj.Ping}
	}
	return obj.pingHolder
}

// description is TBD
// Ping returns a ActionProtocolIpv4Ping
func (obj *actionProtocolIpv4) HasPing() bool {
	return obj.obj.Ping != nil
}

// description is TBD
// SetPing sets the ActionProtocolIpv4Ping value in the ActionProtocolIpv4 object
func (obj *actionProtocolIpv4) SetPing(value ActionProtocolIpv4Ping) ActionProtocolIpv4 {
	obj.setChoice(ActionProtocolIpv4Choice.PING)
	obj.pingHolder = nil
	obj.obj.Ping = value.msg()

	return obj
}

func (obj *actionProtocolIpv4) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface ActionProtocolIpv4")
	}

	if obj.obj.Ping != nil {

		obj.Ping().validateObj(vObj, set_default)
	}

}

func (obj *actionProtocolIpv4) setDefault() {
	var choices_set int = 0
	var choice ActionProtocolIpv4ChoiceEnum

	if obj.obj.Ping != nil {
		choices_set += 1
		choice = ActionProtocolIpv4Choice.PING
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ActionProtocolIpv4")
			}
		} else {
			intVal := otg.ActionProtocolIpv4_Choice_Enum_value[string(choice)]
			enumValue := otg.ActionProtocolIpv4_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionResponseProtocolIpv4 *****
type actionResponseProtocolIpv4 struct {
	validation
	obj          *otg.ActionResponseProtocolIpv4
	marshaller   marshalActionResponseProtocolIpv4
	unMarshaller unMarshalActionResponseProtocolIpv4
	pingHolder   ActionResponseProtocolIpv4Ping
}

func NewActionResponseProtocolIpv4() ActionResponseProtocolIpv4 {
	obj := actionResponseProtocolIpv4{obj: &otg.ActionResponseProtocolIpv4{}}
	obj.setDefault()
	return &obj
}

func (obj *actionResponseProtocolIpv4) msg() *otg.ActionResponseProtocolIpv4 {
	return obj.obj
}

func (obj *actionResponseProtocolIpv4) setMsg(msg *otg.ActionResponseProtocolIpv4) ActionResponseProtocolIpv4 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionResponseProtocolIpv4 struct {
	obj *actionResponseProtocolIpv4
}

type marshalActionResponseProtocolIpv4 interface {
	// ToProto marshals ActionResponseProtocolIpv4 to protobuf object *otg.ActionResponseProtocolIpv4
	ToProto() (*otg.ActionResponseProtocolIpv4, error)
	// ToPbText marshals ActionResponseProtocolIpv4 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionResponseProtocolIpv4 to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionResponseProtocolIpv4 to JSON text
	ToJson() (string, error)
}

type unMarshalactionResponseProtocolIpv4 struct {
	obj *actionResponseProtocolIpv4
}

type unMarshalActionResponseProtocolIpv4 interface {
	// FromProto unmarshals ActionResponseProtocolIpv4 from protobuf object *otg.ActionResponseProtocolIpv4
	FromProto(msg *otg.ActionResponseProtocolIpv4) (ActionResponseProtocolIpv4, error)
	// FromPbText unmarshals ActionResponseProtocolIpv4 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionResponseProtocolIpv4 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionResponseProtocolIpv4 from JSON text
	FromJson(value string) error
}

func (obj *actionResponseProtocolIpv4) Marshal() marshalActionResponseProtocolIpv4 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionResponseProtocolIpv4{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionResponseProtocolIpv4) Unmarshal() unMarshalActionResponseProtocolIpv4 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionResponseProtocolIpv4{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionResponseProtocolIpv4) ToProto() (*otg.ActionResponseProtocolIpv4, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionResponseProtocolIpv4) FromProto(msg *otg.ActionResponseProtocolIpv4) (ActionResponseProtocolIpv4, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionResponseProtocolIpv4) ToPbText() (string, error) {
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

func (m *unMarshalactionResponseProtocolIpv4) FromPbText(value string) error {
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

func (m *marshalactionResponseProtocolIpv4) ToYaml() (string, error) {
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

func (m *unMarshalactionResponseProtocolIpv4) FromYaml(value string) error {
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

func (m *marshalactionResponseProtocolIpv4) ToJson() (string, error) {
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

func (m *unMarshalactionResponseProtocolIpv4) FromJson(value string) error {
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

func (obj *actionResponseProtocolIpv4) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionResponseProtocolIpv4) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionResponseProtocolIpv4) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionResponseProtocolIpv4) Clone() (ActionResponseProtocolIpv4, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionResponseProtocolIpv4()
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

func (obj *actionResponseProtocolIpv4) setNil() {
	obj.pingHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionResponseProtocolIpv4 is response for actions associated with IPv4 on configured resources.
type ActionResponseProtocolIpv4 interface {
	Validation
	// msg marshals ActionResponseProtocolIpv4 to protobuf object *otg.ActionResponseProtocolIpv4
	// and doesn't set defaults
	msg() *otg.ActionResponseProtocolIpv4
	// setMsg unmarshals ActionResponseProtocolIpv4 from protobuf object *otg.ActionResponseProtocolIpv4
	// and doesn't set defaults
	setMsg(*otg.ActionResponseProtocolIpv4) ActionResponseProtocolIpv4
	// provides marshal interface
	Marshal() marshalActionResponseProtocolIpv4
	// provides unmarshal interface
	Unmarshal() unMarshalActionResponseProtocolIpv4
	// validate validates ActionResponseProtocolIpv4
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionResponseProtocolIpv4, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ActionResponseProtocolIpv4ChoiceEnum, set in ActionResponseProtocolIpv4
	Choice() ActionResponseProtocolIpv4ChoiceEnum
	// setChoice assigns ActionResponseProtocolIpv4ChoiceEnum provided by user to ActionResponseProtocolIpv4
	setChoice(value ActionResponseProtocolIpv4ChoiceEnum) ActionResponseProtocolIpv4
	// Ping returns ActionResponseProtocolIpv4Ping, set in ActionResponseProtocolIpv4.
	// ActionResponseProtocolIpv4Ping is response for ping initiated between multiple source and destination pairs.
	Ping() ActionResponseProtocolIpv4Ping
	// SetPing assigns ActionResponseProtocolIpv4Ping provided by user to ActionResponseProtocolIpv4.
	// ActionResponseProtocolIpv4Ping is response for ping initiated between multiple source and destination pairs.
	SetPing(value ActionResponseProtocolIpv4Ping) ActionResponseProtocolIpv4
	// HasPing checks if Ping has been set in ActionResponseProtocolIpv4
	HasPing() bool
	setNil()
}

type ActionResponseProtocolIpv4ChoiceEnum string

// Enum of Choice on ActionResponseProtocolIpv4
var ActionResponseProtocolIpv4Choice = struct {
	PING ActionResponseProtocolIpv4ChoiceEnum
}{
	PING: ActionResponseProtocolIpv4ChoiceEnum("ping"),
}

func (obj *actionResponseProtocolIpv4) Choice() ActionResponseProtocolIpv4ChoiceEnum {
	return ActionResponseProtocolIpv4ChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *actionResponseProtocolIpv4) setChoice(value ActionResponseProtocolIpv4ChoiceEnum) ActionResponseProtocolIpv4 {
	intValue, ok := otg.ActionResponseProtocolIpv4_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ActionResponseProtocolIpv4ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ActionResponseProtocolIpv4_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Ping = nil
	obj.pingHolder = nil

	if value == ActionResponseProtocolIpv4Choice.PING {
		obj.obj.Ping = NewActionResponseProtocolIpv4Ping().msg()
	}

	return obj
}

// description is TBD
// Ping returns a ActionResponseProtocolIpv4Ping
func (obj *actionResponseProtocolIpv4) Ping() ActionResponseProtocolIpv4Ping {
	if obj.obj.Ping == nil {
		obj.setChoice(ActionResponseProtocolIpv4Choice.PING)
	}
	if obj.pingHolder == nil {
		obj.pingHolder = &actionResponseProtocolIpv4Ping{obj: obj.obj.Ping}
	}
	return obj.pingHolder
}

// description is TBD
// Ping returns a ActionResponseProtocolIpv4Ping
func (obj *actionResponseProtocolIpv4) HasPing() bool {
	return obj.obj.Ping != nil
}

// description is TBD
// SetPing sets the ActionResponseProtocolIpv4Ping value in the ActionResponseProtocolIpv4 object
func (obj *actionResponseProtocolIpv4) SetPing(value ActionResponseProtocolIpv4Ping) ActionResponseProtocolIpv4 {
	obj.setChoice(ActionResponseProtocolIpv4Choice.PING)
	obj.pingHolder = nil
	obj.obj.Ping = value.msg()

	return obj
}

func (obj *actionResponseProtocolIpv4) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface ActionResponseProtocolIpv4")
	}

	if obj.obj.Ping != nil {

		obj.Ping().validateObj(vObj, set_default)
	}

}

func (obj *actionResponseProtocolIpv4) setDefault() {

}

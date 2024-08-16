package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ClientOptionsOptionsRequestIncludedMessages *****
type dhcpv6ClientOptionsOptionsRequestIncludedMessages struct {
	validation
	obj            *otg.Dhcpv6ClientOptionsOptionsRequestIncludedMessages
	marshaller     marshalDhcpv6ClientOptionsOptionsRequestIncludedMessages
	unMarshaller   unMarshalDhcpv6ClientOptionsOptionsRequestIncludedMessages
	msgTypesHolder Dhcpv6ClientOptionsMessageType
}

func NewDhcpv6ClientOptionsOptionsRequestIncludedMessages() Dhcpv6ClientOptionsOptionsRequestIncludedMessages {
	obj := dhcpv6ClientOptionsOptionsRequestIncludedMessages{obj: &otg.Dhcpv6ClientOptionsOptionsRequestIncludedMessages{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ClientOptionsOptionsRequestIncludedMessages) msg() *otg.Dhcpv6ClientOptionsOptionsRequestIncludedMessages {
	return obj.obj
}

func (obj *dhcpv6ClientOptionsOptionsRequestIncludedMessages) setMsg(msg *otg.Dhcpv6ClientOptionsOptionsRequestIncludedMessages) Dhcpv6ClientOptionsOptionsRequestIncludedMessages {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ClientOptionsOptionsRequestIncludedMessages struct {
	obj *dhcpv6ClientOptionsOptionsRequestIncludedMessages
}

type marshalDhcpv6ClientOptionsOptionsRequestIncludedMessages interface {
	// ToProto marshals Dhcpv6ClientOptionsOptionsRequestIncludedMessages to protobuf object *otg.Dhcpv6ClientOptionsOptionsRequestIncludedMessages
	ToProto() (*otg.Dhcpv6ClientOptionsOptionsRequestIncludedMessages, error)
	// ToPbText marshals Dhcpv6ClientOptionsOptionsRequestIncludedMessages to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ClientOptionsOptionsRequestIncludedMessages to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ClientOptionsOptionsRequestIncludedMessages to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv6ClientOptionsOptionsRequestIncludedMessages struct {
	obj *dhcpv6ClientOptionsOptionsRequestIncludedMessages
}

type unMarshalDhcpv6ClientOptionsOptionsRequestIncludedMessages interface {
	// FromProto unmarshals Dhcpv6ClientOptionsOptionsRequestIncludedMessages from protobuf object *otg.Dhcpv6ClientOptionsOptionsRequestIncludedMessages
	FromProto(msg *otg.Dhcpv6ClientOptionsOptionsRequestIncludedMessages) (Dhcpv6ClientOptionsOptionsRequestIncludedMessages, error)
	// FromPbText unmarshals Dhcpv6ClientOptionsOptionsRequestIncludedMessages from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ClientOptionsOptionsRequestIncludedMessages from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ClientOptionsOptionsRequestIncludedMessages from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ClientOptionsOptionsRequestIncludedMessages) Marshal() marshalDhcpv6ClientOptionsOptionsRequestIncludedMessages {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ClientOptionsOptionsRequestIncludedMessages{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ClientOptionsOptionsRequestIncludedMessages) Unmarshal() unMarshalDhcpv6ClientOptionsOptionsRequestIncludedMessages {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ClientOptionsOptionsRequestIncludedMessages{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ClientOptionsOptionsRequestIncludedMessages) ToProto() (*otg.Dhcpv6ClientOptionsOptionsRequestIncludedMessages, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ClientOptionsOptionsRequestIncludedMessages) FromProto(msg *otg.Dhcpv6ClientOptionsOptionsRequestIncludedMessages) (Dhcpv6ClientOptionsOptionsRequestIncludedMessages, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ClientOptionsOptionsRequestIncludedMessages) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsOptionsRequestIncludedMessages) FromPbText(value string) error {
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

func (m *marshaldhcpv6ClientOptionsOptionsRequestIncludedMessages) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsOptionsRequestIncludedMessages) FromYaml(value string) error {
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

func (m *marshaldhcpv6ClientOptionsOptionsRequestIncludedMessages) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsOptionsRequestIncludedMessages) FromJson(value string) error {
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

func (obj *dhcpv6ClientOptionsOptionsRequestIncludedMessages) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsOptionsRequestIncludedMessages) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsOptionsRequestIncludedMessages) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ClientOptionsOptionsRequestIncludedMessages) Clone() (Dhcpv6ClientOptionsOptionsRequestIncludedMessages, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ClientOptionsOptionsRequestIncludedMessages()
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

func (obj *dhcpv6ClientOptionsOptionsRequestIncludedMessages) setNil() {
	obj.msgTypesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Dhcpv6ClientOptionsOptionsRequestIncludedMessages is the dhcpv6 client messages where the option will be included. If all is selected the selected option will be added in the all the Dhcpv6 client messages, else based on the selection in particular Dhcpv6 client messages the option will be included.
type Dhcpv6ClientOptionsOptionsRequestIncludedMessages interface {
	Validation
	// msg marshals Dhcpv6ClientOptionsOptionsRequestIncludedMessages to protobuf object *otg.Dhcpv6ClientOptionsOptionsRequestIncludedMessages
	// and doesn't set defaults
	msg() *otg.Dhcpv6ClientOptionsOptionsRequestIncludedMessages
	// setMsg unmarshals Dhcpv6ClientOptionsOptionsRequestIncludedMessages from protobuf object *otg.Dhcpv6ClientOptionsOptionsRequestIncludedMessages
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ClientOptionsOptionsRequestIncludedMessages) Dhcpv6ClientOptionsOptionsRequestIncludedMessages
	// provides marshal interface
	Marshal() marshalDhcpv6ClientOptionsOptionsRequestIncludedMessages
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ClientOptionsOptionsRequestIncludedMessages
	// validate validates Dhcpv6ClientOptionsOptionsRequestIncludedMessages
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ClientOptionsOptionsRequestIncludedMessages, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Dhcpv6ClientOptionsOptionsRequestIncludedMessagesChoiceEnum, set in Dhcpv6ClientOptionsOptionsRequestIncludedMessages
	Choice() Dhcpv6ClientOptionsOptionsRequestIncludedMessagesChoiceEnum
	// setChoice assigns Dhcpv6ClientOptionsOptionsRequestIncludedMessagesChoiceEnum provided by user to Dhcpv6ClientOptionsOptionsRequestIncludedMessages
	setChoice(value Dhcpv6ClientOptionsOptionsRequestIncludedMessagesChoiceEnum) Dhcpv6ClientOptionsOptionsRequestIncludedMessages
	// HasChoice checks if Choice has been set in Dhcpv6ClientOptionsOptionsRequestIncludedMessages
	HasChoice() bool
	// getter for All to set choice.
	All()
	// MsgTypes returns Dhcpv6ClientOptionsMessageType, set in Dhcpv6ClientOptionsOptionsRequestIncludedMessages.
	// Dhcpv6ClientOptionsMessageType is the dhcpv6 client messages where the option will be included.
	MsgTypes() Dhcpv6ClientOptionsMessageType
	// SetMsgTypes assigns Dhcpv6ClientOptionsMessageType provided by user to Dhcpv6ClientOptionsOptionsRequestIncludedMessages.
	// Dhcpv6ClientOptionsMessageType is the dhcpv6 client messages where the option will be included.
	SetMsgTypes(value Dhcpv6ClientOptionsMessageType) Dhcpv6ClientOptionsOptionsRequestIncludedMessages
	// HasMsgTypes checks if MsgTypes has been set in Dhcpv6ClientOptionsOptionsRequestIncludedMessages
	HasMsgTypes() bool
	setNil()
}

type Dhcpv6ClientOptionsOptionsRequestIncludedMessagesChoiceEnum string

// Enum of Choice on Dhcpv6ClientOptionsOptionsRequestIncludedMessages
var Dhcpv6ClientOptionsOptionsRequestIncludedMessagesChoice = struct {
	ALL       Dhcpv6ClientOptionsOptionsRequestIncludedMessagesChoiceEnum
	MSG_TYPES Dhcpv6ClientOptionsOptionsRequestIncludedMessagesChoiceEnum
}{
	ALL:       Dhcpv6ClientOptionsOptionsRequestIncludedMessagesChoiceEnum("all"),
	MSG_TYPES: Dhcpv6ClientOptionsOptionsRequestIncludedMessagesChoiceEnum("msg_types"),
}

func (obj *dhcpv6ClientOptionsOptionsRequestIncludedMessages) Choice() Dhcpv6ClientOptionsOptionsRequestIncludedMessagesChoiceEnum {
	return Dhcpv6ClientOptionsOptionsRequestIncludedMessagesChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for All to set choice
func (obj *dhcpv6ClientOptionsOptionsRequestIncludedMessages) All() {
	obj.setChoice(Dhcpv6ClientOptionsOptionsRequestIncludedMessagesChoice.ALL)
}

// The client message name where the option is included, by default it is all.
// Choice returns a string
func (obj *dhcpv6ClientOptionsOptionsRequestIncludedMessages) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *dhcpv6ClientOptionsOptionsRequestIncludedMessages) setChoice(value Dhcpv6ClientOptionsOptionsRequestIncludedMessagesChoiceEnum) Dhcpv6ClientOptionsOptionsRequestIncludedMessages {
	intValue, ok := otg.Dhcpv6ClientOptionsOptionsRequestIncludedMessages_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Dhcpv6ClientOptionsOptionsRequestIncludedMessagesChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Dhcpv6ClientOptionsOptionsRequestIncludedMessages_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.MsgTypes = nil
	obj.msgTypesHolder = nil

	if value == Dhcpv6ClientOptionsOptionsRequestIncludedMessagesChoice.MSG_TYPES {
		obj.obj.MsgTypes = NewDhcpv6ClientOptionsMessageType().msg()
	}

	return obj
}

// User must specify the Dhcpv6 message type.
// MsgTypes returns a Dhcpv6ClientOptionsMessageType
func (obj *dhcpv6ClientOptionsOptionsRequestIncludedMessages) MsgTypes() Dhcpv6ClientOptionsMessageType {
	if obj.obj.MsgTypes == nil {
		obj.setChoice(Dhcpv6ClientOptionsOptionsRequestIncludedMessagesChoice.MSG_TYPES)
	}
	if obj.msgTypesHolder == nil {
		obj.msgTypesHolder = &dhcpv6ClientOptionsMessageType{obj: obj.obj.MsgTypes}
	}
	return obj.msgTypesHolder
}

// User must specify the Dhcpv6 message type.
// MsgTypes returns a Dhcpv6ClientOptionsMessageType
func (obj *dhcpv6ClientOptionsOptionsRequestIncludedMessages) HasMsgTypes() bool {
	return obj.obj.MsgTypes != nil
}

// User must specify the Dhcpv6 message type.
// SetMsgTypes sets the Dhcpv6ClientOptionsMessageType value in the Dhcpv6ClientOptionsOptionsRequestIncludedMessages object
func (obj *dhcpv6ClientOptionsOptionsRequestIncludedMessages) SetMsgTypes(value Dhcpv6ClientOptionsMessageType) Dhcpv6ClientOptionsOptionsRequestIncludedMessages {
	obj.setChoice(Dhcpv6ClientOptionsOptionsRequestIncludedMessagesChoice.MSG_TYPES)
	obj.msgTypesHolder = nil
	obj.obj.MsgTypes = value.msg()

	return obj
}

func (obj *dhcpv6ClientOptionsOptionsRequestIncludedMessages) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MsgTypes != nil {

		obj.MsgTypes().validateObj(vObj, set_default)
	}

}

func (obj *dhcpv6ClientOptionsOptionsRequestIncludedMessages) setDefault() {
	var choices_set int = 0
	var choice Dhcpv6ClientOptionsOptionsRequestIncludedMessagesChoiceEnum

	if obj.obj.MsgTypes != nil {
		choices_set += 1
		choice = Dhcpv6ClientOptionsOptionsRequestIncludedMessagesChoice.MSG_TYPES
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Dhcpv6ClientOptionsOptionsRequestIncludedMessagesChoice.ALL)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Dhcpv6ClientOptionsOptionsRequestIncludedMessages")
			}
		} else {
			intVal := otg.Dhcpv6ClientOptionsOptionsRequestIncludedMessages_Choice_Enum_value[string(choice)]
			enumValue := otg.Dhcpv6ClientOptionsOptionsRequestIncludedMessages_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

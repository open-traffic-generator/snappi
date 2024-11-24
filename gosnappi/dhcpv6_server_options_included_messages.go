package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ServerOptionsIncludedMessages *****
type dhcpv6ServerOptionsIncludedMessages struct {
	validation
	obj            *otg.Dhcpv6ServerOptionsIncludedMessages
	marshaller     marshalDhcpv6ServerOptionsIncludedMessages
	unMarshaller   unMarshalDhcpv6ServerOptionsIncludedMessages
	msgTypesHolder Dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter
}

func NewDhcpv6ServerOptionsIncludedMessages() Dhcpv6ServerOptionsIncludedMessages {
	obj := dhcpv6ServerOptionsIncludedMessages{obj: &otg.Dhcpv6ServerOptionsIncludedMessages{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ServerOptionsIncludedMessages) msg() *otg.Dhcpv6ServerOptionsIncludedMessages {
	return obj.obj
}

func (obj *dhcpv6ServerOptionsIncludedMessages) setMsg(msg *otg.Dhcpv6ServerOptionsIncludedMessages) Dhcpv6ServerOptionsIncludedMessages {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ServerOptionsIncludedMessages struct {
	obj *dhcpv6ServerOptionsIncludedMessages
}

type marshalDhcpv6ServerOptionsIncludedMessages interface {
	// ToProto marshals Dhcpv6ServerOptionsIncludedMessages to protobuf object *otg.Dhcpv6ServerOptionsIncludedMessages
	ToProto() (*otg.Dhcpv6ServerOptionsIncludedMessages, error)
	// ToPbText marshals Dhcpv6ServerOptionsIncludedMessages to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ServerOptionsIncludedMessages to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ServerOptionsIncludedMessages to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv6ServerOptionsIncludedMessages to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv6ServerOptionsIncludedMessages struct {
	obj *dhcpv6ServerOptionsIncludedMessages
}

type unMarshalDhcpv6ServerOptionsIncludedMessages interface {
	// FromProto unmarshals Dhcpv6ServerOptionsIncludedMessages from protobuf object *otg.Dhcpv6ServerOptionsIncludedMessages
	FromProto(msg *otg.Dhcpv6ServerOptionsIncludedMessages) (Dhcpv6ServerOptionsIncludedMessages, error)
	// FromPbText unmarshals Dhcpv6ServerOptionsIncludedMessages from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ServerOptionsIncludedMessages from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ServerOptionsIncludedMessages from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ServerOptionsIncludedMessages) Marshal() marshalDhcpv6ServerOptionsIncludedMessages {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ServerOptionsIncludedMessages{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ServerOptionsIncludedMessages) Unmarshal() unMarshalDhcpv6ServerOptionsIncludedMessages {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ServerOptionsIncludedMessages{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ServerOptionsIncludedMessages) ToProto() (*otg.Dhcpv6ServerOptionsIncludedMessages, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ServerOptionsIncludedMessages) FromProto(msg *otg.Dhcpv6ServerOptionsIncludedMessages) (Dhcpv6ServerOptionsIncludedMessages, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ServerOptionsIncludedMessages) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ServerOptionsIncludedMessages) FromPbText(value string) error {
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

func (m *marshaldhcpv6ServerOptionsIncludedMessages) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ServerOptionsIncludedMessages) FromYaml(value string) error {
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

func (m *marshaldhcpv6ServerOptionsIncludedMessages) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv6ServerOptionsIncludedMessages) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ServerOptionsIncludedMessages) FromJson(value string) error {
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

func (obj *dhcpv6ServerOptionsIncludedMessages) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ServerOptionsIncludedMessages) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ServerOptionsIncludedMessages) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ServerOptionsIncludedMessages) Clone() (Dhcpv6ServerOptionsIncludedMessages, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ServerOptionsIncludedMessages()
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

func (obj *dhcpv6ServerOptionsIncludedMessages) setNil() {
	obj.msgTypesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Dhcpv6ServerOptionsIncludedMessages is the dhcpv6 server messages where the option will be included. If all is selected the selected option will be added  in the all the Dhcpv6 server messages, else based on the selection in particular Dhcpv6 server messages the option  will be included.
type Dhcpv6ServerOptionsIncludedMessages interface {
	Validation
	// msg marshals Dhcpv6ServerOptionsIncludedMessages to protobuf object *otg.Dhcpv6ServerOptionsIncludedMessages
	// and doesn't set defaults
	msg() *otg.Dhcpv6ServerOptionsIncludedMessages
	// setMsg unmarshals Dhcpv6ServerOptionsIncludedMessages from protobuf object *otg.Dhcpv6ServerOptionsIncludedMessages
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ServerOptionsIncludedMessages) Dhcpv6ServerOptionsIncludedMessages
	// provides marshal interface
	Marshal() marshalDhcpv6ServerOptionsIncludedMessages
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ServerOptionsIncludedMessages
	// validate validates Dhcpv6ServerOptionsIncludedMessages
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ServerOptionsIncludedMessages, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Dhcpv6ServerOptionsIncludedMessagesChoiceEnum, set in Dhcpv6ServerOptionsIncludedMessages
	Choice() Dhcpv6ServerOptionsIncludedMessagesChoiceEnum
	// setChoice assigns Dhcpv6ServerOptionsIncludedMessagesChoiceEnum provided by user to Dhcpv6ServerOptionsIncludedMessages
	setChoice(value Dhcpv6ServerOptionsIncludedMessagesChoiceEnum) Dhcpv6ServerOptionsIncludedMessages
	// HasChoice checks if Choice has been set in Dhcpv6ServerOptionsIncludedMessages
	HasChoice() bool
	// getter for All to set choice.
	All()
	// MsgTypes returns Dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIterIter, set in Dhcpv6ServerOptionsIncludedMessages
	MsgTypes() Dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter
	setNil()
}

type Dhcpv6ServerOptionsIncludedMessagesChoiceEnum string

// Enum of Choice on Dhcpv6ServerOptionsIncludedMessages
var Dhcpv6ServerOptionsIncludedMessagesChoice = struct {
	ALL       Dhcpv6ServerOptionsIncludedMessagesChoiceEnum
	MSG_TYPES Dhcpv6ServerOptionsIncludedMessagesChoiceEnum
}{
	ALL:       Dhcpv6ServerOptionsIncludedMessagesChoiceEnum("all"),
	MSG_TYPES: Dhcpv6ServerOptionsIncludedMessagesChoiceEnum("msg_types"),
}

func (obj *dhcpv6ServerOptionsIncludedMessages) Choice() Dhcpv6ServerOptionsIncludedMessagesChoiceEnum {
	return Dhcpv6ServerOptionsIncludedMessagesChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for All to set choice
func (obj *dhcpv6ServerOptionsIncludedMessages) All() {
	obj.setChoice(Dhcpv6ServerOptionsIncludedMessagesChoice.ALL)
}

// The server message name where the option is included, by default it is all.
// Choice returns a string
func (obj *dhcpv6ServerOptionsIncludedMessages) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *dhcpv6ServerOptionsIncludedMessages) setChoice(value Dhcpv6ServerOptionsIncludedMessagesChoiceEnum) Dhcpv6ServerOptionsIncludedMessages {
	intValue, ok := otg.Dhcpv6ServerOptionsIncludedMessages_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Dhcpv6ServerOptionsIncludedMessagesChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Dhcpv6ServerOptionsIncludedMessages_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.MsgTypes = nil
	obj.msgTypesHolder = nil

	if value == Dhcpv6ServerOptionsIncludedMessagesChoice.MSG_TYPES {
		obj.obj.MsgTypes = []*otg.Dhcpv6ServerOptionsMessageType{}
	}

	return obj
}

// User must specify the Dhcpv6 message type.
// MsgTypes returns a []Dhcpv6ServerOptionsMessageType
func (obj *dhcpv6ServerOptionsIncludedMessages) MsgTypes() Dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter {
	if len(obj.obj.MsgTypes) == 0 {
		obj.setChoice(Dhcpv6ServerOptionsIncludedMessagesChoice.MSG_TYPES)
	}
	if obj.msgTypesHolder == nil {
		obj.msgTypesHolder = newDhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter(&obj.obj.MsgTypes).setMsg(obj)
	}
	return obj.msgTypesHolder
}

type dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter struct {
	obj                                 *dhcpv6ServerOptionsIncludedMessages
	dhcpv6ServerOptionsMessageTypeSlice []Dhcpv6ServerOptionsMessageType
	fieldPtr                            *[]*otg.Dhcpv6ServerOptionsMessageType
}

func newDhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter(ptr *[]*otg.Dhcpv6ServerOptionsMessageType) Dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter {
	return &dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter{fieldPtr: ptr}
}

type Dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter interface {
	setMsg(*dhcpv6ServerOptionsIncludedMessages) Dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter
	Items() []Dhcpv6ServerOptionsMessageType
	Add() Dhcpv6ServerOptionsMessageType
	Append(items ...Dhcpv6ServerOptionsMessageType) Dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter
	Set(index int, newObj Dhcpv6ServerOptionsMessageType) Dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter
	Clear() Dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter
	clearHolderSlice() Dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter
	appendHolderSlice(item Dhcpv6ServerOptionsMessageType) Dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter
}

func (obj *dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter) setMsg(msg *dhcpv6ServerOptionsIncludedMessages) Dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpv6ServerOptionsMessageType{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter) Items() []Dhcpv6ServerOptionsMessageType {
	return obj.dhcpv6ServerOptionsMessageTypeSlice
}

func (obj *dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter) Add() Dhcpv6ServerOptionsMessageType {
	newObj := &otg.Dhcpv6ServerOptionsMessageType{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpv6ServerOptionsMessageType{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpv6ServerOptionsMessageTypeSlice = append(obj.dhcpv6ServerOptionsMessageTypeSlice, newLibObj)
	return newLibObj
}

func (obj *dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter) Append(items ...Dhcpv6ServerOptionsMessageType) Dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpv6ServerOptionsMessageTypeSlice = append(obj.dhcpv6ServerOptionsMessageTypeSlice, item)
	}
	return obj
}

func (obj *dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter) Set(index int, newObj Dhcpv6ServerOptionsMessageType) Dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpv6ServerOptionsMessageTypeSlice[index] = newObj
	return obj
}
func (obj *dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter) Clear() Dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Dhcpv6ServerOptionsMessageType{}
		obj.dhcpv6ServerOptionsMessageTypeSlice = []Dhcpv6ServerOptionsMessageType{}
	}
	return obj
}
func (obj *dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter) clearHolderSlice() Dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter {
	if len(obj.dhcpv6ServerOptionsMessageTypeSlice) > 0 {
		obj.dhcpv6ServerOptionsMessageTypeSlice = []Dhcpv6ServerOptionsMessageType{}
	}
	return obj
}
func (obj *dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter) appendHolderSlice(item Dhcpv6ServerOptionsMessageType) Dhcpv6ServerOptionsIncludedMessagesDhcpv6ServerOptionsMessageTypeIter {
	obj.dhcpv6ServerOptionsMessageTypeSlice = append(obj.dhcpv6ServerOptionsMessageTypeSlice, item)
	return obj
}

func (obj *dhcpv6ServerOptionsIncludedMessages) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.MsgTypes) != 0 {

		if set_default {
			obj.MsgTypes().clearHolderSlice()
			for _, item := range obj.obj.MsgTypes {
				obj.MsgTypes().appendHolderSlice(&dhcpv6ServerOptionsMessageType{obj: item})
			}
		}
		for _, item := range obj.MsgTypes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *dhcpv6ServerOptionsIncludedMessages) setDefault() {
	var choices_set int = 0
	var choice Dhcpv6ServerOptionsIncludedMessagesChoiceEnum

	if len(obj.obj.MsgTypes) > 0 {
		choices_set += 1
		choice = Dhcpv6ServerOptionsIncludedMessagesChoice.MSG_TYPES
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Dhcpv6ServerOptionsIncludedMessagesChoice.ALL)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Dhcpv6ServerOptionsIncludedMessages")
			}
		} else {
			intVal := otg.Dhcpv6ServerOptionsIncludedMessages_Choice_Enum_value[string(choice)]
			enumValue := otg.Dhcpv6ServerOptionsIncludedMessages_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

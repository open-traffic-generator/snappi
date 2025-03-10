package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ClientOptionsIncludedMessages *****
type dhcpv6ClientOptionsIncludedMessages struct {
	validation
	obj            *otg.Dhcpv6ClientOptionsIncludedMessages
	marshaller     marshalDhcpv6ClientOptionsIncludedMessages
	unMarshaller   unMarshalDhcpv6ClientOptionsIncludedMessages
	msgTypesHolder Dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter
}

func NewDhcpv6ClientOptionsIncludedMessages() Dhcpv6ClientOptionsIncludedMessages {
	obj := dhcpv6ClientOptionsIncludedMessages{obj: &otg.Dhcpv6ClientOptionsIncludedMessages{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ClientOptionsIncludedMessages) msg() *otg.Dhcpv6ClientOptionsIncludedMessages {
	return obj.obj
}

func (obj *dhcpv6ClientOptionsIncludedMessages) setMsg(msg *otg.Dhcpv6ClientOptionsIncludedMessages) Dhcpv6ClientOptionsIncludedMessages {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ClientOptionsIncludedMessages struct {
	obj *dhcpv6ClientOptionsIncludedMessages
}

type marshalDhcpv6ClientOptionsIncludedMessages interface {
	// ToProto marshals Dhcpv6ClientOptionsIncludedMessages to protobuf object *otg.Dhcpv6ClientOptionsIncludedMessages
	ToProto() (*otg.Dhcpv6ClientOptionsIncludedMessages, error)
	// ToPbText marshals Dhcpv6ClientOptionsIncludedMessages to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ClientOptionsIncludedMessages to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ClientOptionsIncludedMessages to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv6ClientOptionsIncludedMessages to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv6ClientOptionsIncludedMessages struct {
	obj *dhcpv6ClientOptionsIncludedMessages
}

type unMarshalDhcpv6ClientOptionsIncludedMessages interface {
	// FromProto unmarshals Dhcpv6ClientOptionsIncludedMessages from protobuf object *otg.Dhcpv6ClientOptionsIncludedMessages
	FromProto(msg *otg.Dhcpv6ClientOptionsIncludedMessages) (Dhcpv6ClientOptionsIncludedMessages, error)
	// FromPbText unmarshals Dhcpv6ClientOptionsIncludedMessages from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ClientOptionsIncludedMessages from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ClientOptionsIncludedMessages from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ClientOptionsIncludedMessages) Marshal() marshalDhcpv6ClientOptionsIncludedMessages {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ClientOptionsIncludedMessages{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ClientOptionsIncludedMessages) Unmarshal() unMarshalDhcpv6ClientOptionsIncludedMessages {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ClientOptionsIncludedMessages{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ClientOptionsIncludedMessages) ToProto() (*otg.Dhcpv6ClientOptionsIncludedMessages, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ClientOptionsIncludedMessages) FromProto(msg *otg.Dhcpv6ClientOptionsIncludedMessages) (Dhcpv6ClientOptionsIncludedMessages, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ClientOptionsIncludedMessages) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsIncludedMessages) FromPbText(value string) error {
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

func (m *marshaldhcpv6ClientOptionsIncludedMessages) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsIncludedMessages) FromYaml(value string) error {
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

func (m *marshaldhcpv6ClientOptionsIncludedMessages) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv6ClientOptionsIncludedMessages) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsIncludedMessages) FromJson(value string) error {
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

func (obj *dhcpv6ClientOptionsIncludedMessages) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsIncludedMessages) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsIncludedMessages) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ClientOptionsIncludedMessages) Clone() (Dhcpv6ClientOptionsIncludedMessages, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ClientOptionsIncludedMessages()
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

func (obj *dhcpv6ClientOptionsIncludedMessages) setNil() {
	obj.msgTypesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Dhcpv6ClientOptionsIncludedMessages is the dhcpv6 client messages where the option will be included. If all is selected the selected option will be added in the all the Dhcpv6 client messages, else based on the selection in particular Dhcpv6 client messages the option will be included.
type Dhcpv6ClientOptionsIncludedMessages interface {
	Validation
	// msg marshals Dhcpv6ClientOptionsIncludedMessages to protobuf object *otg.Dhcpv6ClientOptionsIncludedMessages
	// and doesn't set defaults
	msg() *otg.Dhcpv6ClientOptionsIncludedMessages
	// setMsg unmarshals Dhcpv6ClientOptionsIncludedMessages from protobuf object *otg.Dhcpv6ClientOptionsIncludedMessages
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ClientOptionsIncludedMessages) Dhcpv6ClientOptionsIncludedMessages
	// provides marshal interface
	Marshal() marshalDhcpv6ClientOptionsIncludedMessages
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ClientOptionsIncludedMessages
	// validate validates Dhcpv6ClientOptionsIncludedMessages
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ClientOptionsIncludedMessages, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Dhcpv6ClientOptionsIncludedMessagesChoiceEnum, set in Dhcpv6ClientOptionsIncludedMessages
	Choice() Dhcpv6ClientOptionsIncludedMessagesChoiceEnum
	// setChoice assigns Dhcpv6ClientOptionsIncludedMessagesChoiceEnum provided by user to Dhcpv6ClientOptionsIncludedMessages
	setChoice(value Dhcpv6ClientOptionsIncludedMessagesChoiceEnum) Dhcpv6ClientOptionsIncludedMessages
	// HasChoice checks if Choice has been set in Dhcpv6ClientOptionsIncludedMessages
	HasChoice() bool
	// getter for All to set choice.
	All()
	// MsgTypes returns Dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIterIter, set in Dhcpv6ClientOptionsIncludedMessages
	MsgTypes() Dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter
	setNil()
}

type Dhcpv6ClientOptionsIncludedMessagesChoiceEnum string

// Enum of Choice on Dhcpv6ClientOptionsIncludedMessages
var Dhcpv6ClientOptionsIncludedMessagesChoice = struct {
	ALL       Dhcpv6ClientOptionsIncludedMessagesChoiceEnum
	MSG_TYPES Dhcpv6ClientOptionsIncludedMessagesChoiceEnum
}{
	ALL:       Dhcpv6ClientOptionsIncludedMessagesChoiceEnum("all"),
	MSG_TYPES: Dhcpv6ClientOptionsIncludedMessagesChoiceEnum("msg_types"),
}

func (obj *dhcpv6ClientOptionsIncludedMessages) Choice() Dhcpv6ClientOptionsIncludedMessagesChoiceEnum {
	return Dhcpv6ClientOptionsIncludedMessagesChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for All to set choice
func (obj *dhcpv6ClientOptionsIncludedMessages) All() {
	obj.setChoice(Dhcpv6ClientOptionsIncludedMessagesChoice.ALL)
}

// The client message name where the option is included, by default it is all.
// Choice returns a string
func (obj *dhcpv6ClientOptionsIncludedMessages) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *dhcpv6ClientOptionsIncludedMessages) setChoice(value Dhcpv6ClientOptionsIncludedMessagesChoiceEnum) Dhcpv6ClientOptionsIncludedMessages {
	intValue, ok := otg.Dhcpv6ClientOptionsIncludedMessages_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Dhcpv6ClientOptionsIncludedMessagesChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Dhcpv6ClientOptionsIncludedMessages_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.MsgTypes = nil
	obj.msgTypesHolder = nil

	if value == Dhcpv6ClientOptionsIncludedMessagesChoice.MSG_TYPES {
		obj.obj.MsgTypes = []*otg.Dhcpv6ClientOptionsMessageType{}
	}

	return obj
}

// User must specify the Dhcpv6 message type.
// MsgTypes returns a []Dhcpv6ClientOptionsMessageType
func (obj *dhcpv6ClientOptionsIncludedMessages) MsgTypes() Dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter {
	if len(obj.obj.MsgTypes) == 0 {
		obj.setChoice(Dhcpv6ClientOptionsIncludedMessagesChoice.MSG_TYPES)
	}
	if obj.msgTypesHolder == nil {
		obj.msgTypesHolder = newDhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter(&obj.obj.MsgTypes).setMsg(obj)
	}
	return obj.msgTypesHolder
}

type dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter struct {
	obj                                 *dhcpv6ClientOptionsIncludedMessages
	dhcpv6ClientOptionsMessageTypeSlice []Dhcpv6ClientOptionsMessageType
	fieldPtr                            *[]*otg.Dhcpv6ClientOptionsMessageType
}

func newDhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter(ptr *[]*otg.Dhcpv6ClientOptionsMessageType) Dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter {
	return &dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter{fieldPtr: ptr}
}

type Dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter interface {
	setMsg(*dhcpv6ClientOptionsIncludedMessages) Dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter
	Items() []Dhcpv6ClientOptionsMessageType
	Add() Dhcpv6ClientOptionsMessageType
	Append(items ...Dhcpv6ClientOptionsMessageType) Dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter
	Set(index int, newObj Dhcpv6ClientOptionsMessageType) Dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter
	Clear() Dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter
	clearHolderSlice() Dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter
	appendHolderSlice(item Dhcpv6ClientOptionsMessageType) Dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter
}

func (obj *dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter) setMsg(msg *dhcpv6ClientOptionsIncludedMessages) Dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpv6ClientOptionsMessageType{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter) Items() []Dhcpv6ClientOptionsMessageType {
	return obj.dhcpv6ClientOptionsMessageTypeSlice
}

func (obj *dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter) Add() Dhcpv6ClientOptionsMessageType {
	newObj := &otg.Dhcpv6ClientOptionsMessageType{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpv6ClientOptionsMessageType{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpv6ClientOptionsMessageTypeSlice = append(obj.dhcpv6ClientOptionsMessageTypeSlice, newLibObj)
	return newLibObj
}

func (obj *dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter) Append(items ...Dhcpv6ClientOptionsMessageType) Dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpv6ClientOptionsMessageTypeSlice = append(obj.dhcpv6ClientOptionsMessageTypeSlice, item)
	}
	return obj
}

func (obj *dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter) Set(index int, newObj Dhcpv6ClientOptionsMessageType) Dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpv6ClientOptionsMessageTypeSlice[index] = newObj
	return obj
}
func (obj *dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter) Clear() Dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Dhcpv6ClientOptionsMessageType{}
		obj.dhcpv6ClientOptionsMessageTypeSlice = []Dhcpv6ClientOptionsMessageType{}
	}
	return obj
}
func (obj *dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter) clearHolderSlice() Dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter {
	if len(obj.dhcpv6ClientOptionsMessageTypeSlice) > 0 {
		obj.dhcpv6ClientOptionsMessageTypeSlice = []Dhcpv6ClientOptionsMessageType{}
	}
	return obj
}
func (obj *dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter) appendHolderSlice(item Dhcpv6ClientOptionsMessageType) Dhcpv6ClientOptionsIncludedMessagesDhcpv6ClientOptionsMessageTypeIter {
	obj.dhcpv6ClientOptionsMessageTypeSlice = append(obj.dhcpv6ClientOptionsMessageTypeSlice, item)
	return obj
}

func (obj *dhcpv6ClientOptionsIncludedMessages) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.MsgTypes) != 0 {

		if set_default {
			obj.MsgTypes().clearHolderSlice()
			for _, item := range obj.obj.MsgTypes {
				obj.MsgTypes().appendHolderSlice(&dhcpv6ClientOptionsMessageType{obj: item})
			}
		}
		for _, item := range obj.MsgTypes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *dhcpv6ClientOptionsIncludedMessages) setDefault() {
	var choices_set int = 0
	var choice Dhcpv6ClientOptionsIncludedMessagesChoiceEnum

	if len(obj.obj.MsgTypes) > 0 {
		choices_set += 1
		choice = Dhcpv6ClientOptionsIncludedMessagesChoice.MSG_TYPES
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Dhcpv6ClientOptionsIncludedMessagesChoice.ALL)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Dhcpv6ClientOptionsIncludedMessages")
			}
		} else {
			intVal := otg.Dhcpv6ClientOptionsIncludedMessages_Choice_Enum_value[string(choice)]
			enumValue := otg.Dhcpv6ClientOptionsIncludedMessages_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

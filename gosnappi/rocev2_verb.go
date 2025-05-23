package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2Verb *****
type rocev2Verb struct {
	validation
	obj                      *otg.Rocev2Verb
	marshaller               marshalRocev2Verb
	unMarshaller             unMarshalRocev2Verb
	writeWithImmediateHolder Rocev2ImmediateData
	sendWithImmediateHolder  Rocev2ImmediateData
}

func NewRocev2Verb() Rocev2Verb {
	obj := rocev2Verb{obj: &otg.Rocev2Verb{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2Verb) msg() *otg.Rocev2Verb {
	return obj.obj
}

func (obj *rocev2Verb) setMsg(msg *otg.Rocev2Verb) Rocev2Verb {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2Verb struct {
	obj *rocev2Verb
}

type marshalRocev2Verb interface {
	// ToProto marshals Rocev2Verb to protobuf object *otg.Rocev2Verb
	ToProto() (*otg.Rocev2Verb, error)
	// ToPbText marshals Rocev2Verb to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2Verb to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2Verb to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Rocev2Verb to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrocev2Verb struct {
	obj *rocev2Verb
}

type unMarshalRocev2Verb interface {
	// FromProto unmarshals Rocev2Verb from protobuf object *otg.Rocev2Verb
	FromProto(msg *otg.Rocev2Verb) (Rocev2Verb, error)
	// FromPbText unmarshals Rocev2Verb from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2Verb from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2Verb from JSON text
	FromJson(value string) error
}

func (obj *rocev2Verb) Marshal() marshalRocev2Verb {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2Verb{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2Verb) Unmarshal() unMarshalRocev2Verb {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2Verb{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2Verb) ToProto() (*otg.Rocev2Verb, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2Verb) FromProto(msg *otg.Rocev2Verb) (Rocev2Verb, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2Verb) ToPbText() (string, error) {
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

func (m *unMarshalrocev2Verb) FromPbText(value string) error {
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

func (m *marshalrocev2Verb) ToYaml() (string, error) {
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

func (m *unMarshalrocev2Verb) FromYaml(value string) error {
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

func (m *marshalrocev2Verb) ToJsonRaw() (string, error) {
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

func (m *marshalrocev2Verb) ToJson() (string, error) {
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

func (m *unMarshalrocev2Verb) FromJson(value string) error {
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

func (obj *rocev2Verb) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2Verb) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2Verb) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2Verb) Clone() (Rocev2Verb, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2Verb()
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

func (obj *rocev2Verb) setNil() {
	obj.writeWithImmediateHolder = nil
	obj.sendWithImmediateHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2Verb is roCEv2 Verb. Available options are: WRITE, WRITE_With_Immediate, SEND, SEND_With_Immediate and READ.
type Rocev2Verb interface {
	Validation
	// msg marshals Rocev2Verb to protobuf object *otg.Rocev2Verb
	// and doesn't set defaults
	msg() *otg.Rocev2Verb
	// setMsg unmarshals Rocev2Verb from protobuf object *otg.Rocev2Verb
	// and doesn't set defaults
	setMsg(*otg.Rocev2Verb) Rocev2Verb
	// provides marshal interface
	Marshal() marshalRocev2Verb
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2Verb
	// validate validates Rocev2Verb
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2Verb, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Rocev2VerbChoiceEnum, set in Rocev2Verb
	Choice() Rocev2VerbChoiceEnum
	// setChoice assigns Rocev2VerbChoiceEnum provided by user to Rocev2Verb
	setChoice(value Rocev2VerbChoiceEnum) Rocev2Verb
	// HasChoice checks if Choice has been set in Rocev2Verb
	HasChoice() bool
	// getter for Write to set choice.
	Write()
	// getter for Read to set choice.
	Read()
	// getter for Send to set choice.
	Send()
	// WriteWithImmediate returns Rocev2ImmediateData, set in Rocev2Verb.
	// Rocev2ImmediateData is four bytes of immediate Data for SEND/WRITE with immediate.
	WriteWithImmediate() Rocev2ImmediateData
	// SetWriteWithImmediate assigns Rocev2ImmediateData provided by user to Rocev2Verb.
	// Rocev2ImmediateData is four bytes of immediate Data for SEND/WRITE with immediate.
	SetWriteWithImmediate(value Rocev2ImmediateData) Rocev2Verb
	// HasWriteWithImmediate checks if WriteWithImmediate has been set in Rocev2Verb
	HasWriteWithImmediate() bool
	// SendWithImmediate returns Rocev2ImmediateData, set in Rocev2Verb.
	// Rocev2ImmediateData is four bytes of immediate Data for SEND/WRITE with immediate.
	SendWithImmediate() Rocev2ImmediateData
	// SetSendWithImmediate assigns Rocev2ImmediateData provided by user to Rocev2Verb.
	// Rocev2ImmediateData is four bytes of immediate Data for SEND/WRITE with immediate.
	SetSendWithImmediate(value Rocev2ImmediateData) Rocev2Verb
	// HasSendWithImmediate checks if SendWithImmediate has been set in Rocev2Verb
	HasSendWithImmediate() bool
	setNil()
}

type Rocev2VerbChoiceEnum string

// Enum of Choice on Rocev2Verb
var Rocev2VerbChoice = struct {
	WRITE                Rocev2VerbChoiceEnum
	WRITE_WITH_IMMEDIATE Rocev2VerbChoiceEnum
	SEND                 Rocev2VerbChoiceEnum
	SEND_WITH_IMMEDIATE  Rocev2VerbChoiceEnum
	READ                 Rocev2VerbChoiceEnum
}{
	WRITE:                Rocev2VerbChoiceEnum("write"),
	WRITE_WITH_IMMEDIATE: Rocev2VerbChoiceEnum("write_with_immediate"),
	SEND:                 Rocev2VerbChoiceEnum("send"),
	SEND_WITH_IMMEDIATE:  Rocev2VerbChoiceEnum("send_with_immediate"),
	READ:                 Rocev2VerbChoiceEnum("read"),
}

func (obj *rocev2Verb) Choice() Rocev2VerbChoiceEnum {
	return Rocev2VerbChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Write to set choice
func (obj *rocev2Verb) Write() {
	obj.setChoice(Rocev2VerbChoice.WRITE)
}

// getter for Read to set choice
func (obj *rocev2Verb) Read() {
	obj.setChoice(Rocev2VerbChoice.READ)
}

// getter for Send to set choice
func (obj *rocev2Verb) Send() {
	obj.setChoice(Rocev2VerbChoice.SEND)
}

// description is TBD
// Choice returns a string
func (obj *rocev2Verb) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *rocev2Verb) setChoice(value Rocev2VerbChoiceEnum) Rocev2Verb {
	intValue, ok := otg.Rocev2Verb_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Rocev2VerbChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Rocev2Verb_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.SendWithImmediate = nil
	obj.sendWithImmediateHolder = nil
	obj.obj.WriteWithImmediate = nil
	obj.writeWithImmediateHolder = nil

	if value == Rocev2VerbChoice.WRITE_WITH_IMMEDIATE {
		obj.obj.WriteWithImmediate = NewRocev2ImmediateData().msg()
	}

	if value == Rocev2VerbChoice.SEND_WITH_IMMEDIATE {
		obj.obj.SendWithImmediate = NewRocev2ImmediateData().msg()
	}

	return obj
}

// description is TBD
// WriteWithImmediate returns a Rocev2ImmediateData
func (obj *rocev2Verb) WriteWithImmediate() Rocev2ImmediateData {
	if obj.obj.WriteWithImmediate == nil {
		obj.setChoice(Rocev2VerbChoice.WRITE_WITH_IMMEDIATE)
	}
	if obj.writeWithImmediateHolder == nil {
		obj.writeWithImmediateHolder = &rocev2ImmediateData{obj: obj.obj.WriteWithImmediate}
	}
	return obj.writeWithImmediateHolder
}

// description is TBD
// WriteWithImmediate returns a Rocev2ImmediateData
func (obj *rocev2Verb) HasWriteWithImmediate() bool {
	return obj.obj.WriteWithImmediate != nil
}

// description is TBD
// SetWriteWithImmediate sets the Rocev2ImmediateData value in the Rocev2Verb object
func (obj *rocev2Verb) SetWriteWithImmediate(value Rocev2ImmediateData) Rocev2Verb {
	obj.setChoice(Rocev2VerbChoice.WRITE_WITH_IMMEDIATE)
	obj.writeWithImmediateHolder = nil
	obj.obj.WriteWithImmediate = value.msg()

	return obj
}

// description is TBD
// SendWithImmediate returns a Rocev2ImmediateData
func (obj *rocev2Verb) SendWithImmediate() Rocev2ImmediateData {
	if obj.obj.SendWithImmediate == nil {
		obj.setChoice(Rocev2VerbChoice.SEND_WITH_IMMEDIATE)
	}
	if obj.sendWithImmediateHolder == nil {
		obj.sendWithImmediateHolder = &rocev2ImmediateData{obj: obj.obj.SendWithImmediate}
	}
	return obj.sendWithImmediateHolder
}

// description is TBD
// SendWithImmediate returns a Rocev2ImmediateData
func (obj *rocev2Verb) HasSendWithImmediate() bool {
	return obj.obj.SendWithImmediate != nil
}

// description is TBD
// SetSendWithImmediate sets the Rocev2ImmediateData value in the Rocev2Verb object
func (obj *rocev2Verb) SetSendWithImmediate(value Rocev2ImmediateData) Rocev2Verb {
	obj.setChoice(Rocev2VerbChoice.SEND_WITH_IMMEDIATE)
	obj.sendWithImmediateHolder = nil
	obj.obj.SendWithImmediate = value.msg()

	return obj
}

func (obj *rocev2Verb) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.WriteWithImmediate != nil {

		obj.WriteWithImmediate().validateObj(vObj, set_default)
	}

	if obj.obj.SendWithImmediate != nil {

		obj.SendWithImmediate().validateObj(vObj, set_default)
	}

}

func (obj *rocev2Verb) setDefault() {
	var choices_set int = 0
	var choice Rocev2VerbChoiceEnum

	if obj.obj.WriteWithImmediate != nil {
		choices_set += 1
		choice = Rocev2VerbChoice.WRITE_WITH_IMMEDIATE
	}

	if obj.obj.SendWithImmediate != nil {
		choices_set += 1
		choice = Rocev2VerbChoice.SEND_WITH_IMMEDIATE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Rocev2VerbChoice.WRITE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Rocev2Verb")
			}
		} else {
			intVal := otg.Rocev2Verb_Choice_Enum_value[string(choice)]
			enumValue := otg.Rocev2Verb_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

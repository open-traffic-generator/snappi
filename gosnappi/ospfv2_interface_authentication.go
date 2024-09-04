package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2InterfaceAuthentication *****
type ospfv2InterfaceAuthentication struct {
	validation
	obj          *otg.Ospfv2InterfaceAuthentication
	marshaller   marshalOspfv2InterfaceAuthentication
	unMarshaller unMarshalOspfv2InterfaceAuthentication
	md5SHolder   Ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter
}

func NewOspfv2InterfaceAuthentication() Ospfv2InterfaceAuthentication {
	obj := ospfv2InterfaceAuthentication{obj: &otg.Ospfv2InterfaceAuthentication{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2InterfaceAuthentication) msg() *otg.Ospfv2InterfaceAuthentication {
	return obj.obj
}

func (obj *ospfv2InterfaceAuthentication) setMsg(msg *otg.Ospfv2InterfaceAuthentication) Ospfv2InterfaceAuthentication {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2InterfaceAuthentication struct {
	obj *ospfv2InterfaceAuthentication
}

type marshalOspfv2InterfaceAuthentication interface {
	// ToProto marshals Ospfv2InterfaceAuthentication to protobuf object *otg.Ospfv2InterfaceAuthentication
	ToProto() (*otg.Ospfv2InterfaceAuthentication, error)
	// ToPbText marshals Ospfv2InterfaceAuthentication to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2InterfaceAuthentication to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2InterfaceAuthentication to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2InterfaceAuthentication struct {
	obj *ospfv2InterfaceAuthentication
}

type unMarshalOspfv2InterfaceAuthentication interface {
	// FromProto unmarshals Ospfv2InterfaceAuthentication from protobuf object *otg.Ospfv2InterfaceAuthentication
	FromProto(msg *otg.Ospfv2InterfaceAuthentication) (Ospfv2InterfaceAuthentication, error)
	// FromPbText unmarshals Ospfv2InterfaceAuthentication from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2InterfaceAuthentication from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2InterfaceAuthentication from JSON text
	FromJson(value string) error
}

func (obj *ospfv2InterfaceAuthentication) Marshal() marshalOspfv2InterfaceAuthentication {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2InterfaceAuthentication{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2InterfaceAuthentication) Unmarshal() unMarshalOspfv2InterfaceAuthentication {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2InterfaceAuthentication{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2InterfaceAuthentication) ToProto() (*otg.Ospfv2InterfaceAuthentication, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2InterfaceAuthentication) FromProto(msg *otg.Ospfv2InterfaceAuthentication) (Ospfv2InterfaceAuthentication, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2InterfaceAuthentication) ToPbText() (string, error) {
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

func (m *unMarshalospfv2InterfaceAuthentication) FromPbText(value string) error {
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

func (m *marshalospfv2InterfaceAuthentication) ToYaml() (string, error) {
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

func (m *unMarshalospfv2InterfaceAuthentication) FromYaml(value string) error {
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

func (m *marshalospfv2InterfaceAuthentication) ToJson() (string, error) {
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

func (m *unMarshalospfv2InterfaceAuthentication) FromJson(value string) error {
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

func (obj *ospfv2InterfaceAuthentication) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2InterfaceAuthentication) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2InterfaceAuthentication) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2InterfaceAuthentication) Clone() (Ospfv2InterfaceAuthentication, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2InterfaceAuthentication()
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

func (obj *ospfv2InterfaceAuthentication) setNil() {
	obj.md5SHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2InterfaceAuthentication is this contains OSPFv2 authentication properties.
// Reference: https://www.rfc-editor.org/rfc/rfc2328#appendix-D
type Ospfv2InterfaceAuthentication interface {
	Validation
	// msg marshals Ospfv2InterfaceAuthentication to protobuf object *otg.Ospfv2InterfaceAuthentication
	// and doesn't set defaults
	msg() *otg.Ospfv2InterfaceAuthentication
	// setMsg unmarshals Ospfv2InterfaceAuthentication from protobuf object *otg.Ospfv2InterfaceAuthentication
	// and doesn't set defaults
	setMsg(*otg.Ospfv2InterfaceAuthentication) Ospfv2InterfaceAuthentication
	// provides marshal interface
	Marshal() marshalOspfv2InterfaceAuthentication
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2InterfaceAuthentication
	// validate validates Ospfv2InterfaceAuthentication
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2InterfaceAuthentication, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Ospfv2InterfaceAuthenticationChoiceEnum, set in Ospfv2InterfaceAuthentication
	Choice() Ospfv2InterfaceAuthenticationChoiceEnum
	// setChoice assigns Ospfv2InterfaceAuthenticationChoiceEnum provided by user to Ospfv2InterfaceAuthentication
	setChoice(value Ospfv2InterfaceAuthenticationChoiceEnum) Ospfv2InterfaceAuthentication
	// HasChoice checks if Choice has been set in Ospfv2InterfaceAuthentication
	HasChoice() bool
	// getter for Md5List to set choice.
	Md5List()
	// Md5S returns Ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5IterIter, set in Ospfv2InterfaceAuthentication
	Md5S() Ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter
	// ClearText returns string, set in Ospfv2InterfaceAuthentication.
	ClearText() string
	// SetClearText assigns string provided by user to Ospfv2InterfaceAuthentication
	SetClearText(value string) Ospfv2InterfaceAuthentication
	// HasClearText checks if ClearText has been set in Ospfv2InterfaceAuthentication
	HasClearText() bool
	setNil()
}

type Ospfv2InterfaceAuthenticationChoiceEnum string

// Enum of Choice on Ospfv2InterfaceAuthentication
var Ospfv2InterfaceAuthenticationChoice = struct {
	MD5_LIST   Ospfv2InterfaceAuthenticationChoiceEnum
	CLEAR_TEXT Ospfv2InterfaceAuthenticationChoiceEnum
}{
	MD5_LIST:   Ospfv2InterfaceAuthenticationChoiceEnum("md5_list"),
	CLEAR_TEXT: Ospfv2InterfaceAuthenticationChoiceEnum("clear_text"),
}

func (obj *ospfv2InterfaceAuthentication) Choice() Ospfv2InterfaceAuthenticationChoiceEnum {
	return Ospfv2InterfaceAuthenticationChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Md5List to set choice
func (obj *ospfv2InterfaceAuthentication) Md5List() {
	obj.setChoice(Ospfv2InterfaceAuthenticationChoice.MD5_LIST)
}

// The authentication method.
// - md5 - Cryptographic authentication.
// - clear_text - Simple password authentication. A 64-bit field is configured on a per-network basis.
// All packets sent on a particular network must have this configured value (in clear text)
// in their OSPF header 64-bit authentication field.
// Choice returns a string
func (obj *ospfv2InterfaceAuthentication) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ospfv2InterfaceAuthentication) setChoice(value Ospfv2InterfaceAuthenticationChoiceEnum) Ospfv2InterfaceAuthentication {
	intValue, ok := otg.Ospfv2InterfaceAuthentication_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2InterfaceAuthenticationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2InterfaceAuthentication_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.ClearText = nil

	if value == Ospfv2InterfaceAuthenticationChoice.CLEAR_TEXT {
		defaultValue := "keysight"
		obj.obj.ClearText = &defaultValue
	}

	return obj
}

// List of MD5 Key IDs and MD5 Keys.
// Md5S returns a []Ospfv2AuthenticationMd5
func (obj *ospfv2InterfaceAuthentication) Md5S() Ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter {
	if len(obj.obj.Md5S) == 0 {
		obj.obj.Md5S = []*otg.Ospfv2AuthenticationMd5{}
	}
	if obj.md5SHolder == nil {
		obj.md5SHolder = newOspfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter(&obj.obj.Md5S).setMsg(obj)
	}
	return obj.md5SHolder
}

type ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter struct {
	obj                          *ospfv2InterfaceAuthentication
	ospfv2AuthenticationMd5Slice []Ospfv2AuthenticationMd5
	fieldPtr                     *[]*otg.Ospfv2AuthenticationMd5
}

func newOspfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter(ptr *[]*otg.Ospfv2AuthenticationMd5) Ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter {
	return &ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter{fieldPtr: ptr}
}

type Ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter interface {
	setMsg(*ospfv2InterfaceAuthentication) Ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter
	Items() []Ospfv2AuthenticationMd5
	Add() Ospfv2AuthenticationMd5
	Append(items ...Ospfv2AuthenticationMd5) Ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter
	Set(index int, newObj Ospfv2AuthenticationMd5) Ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter
	Clear() Ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter
	clearHolderSlice() Ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter
	appendHolderSlice(item Ospfv2AuthenticationMd5) Ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter
}

func (obj *ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter) setMsg(msg *ospfv2InterfaceAuthentication) Ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2AuthenticationMd5{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter) Items() []Ospfv2AuthenticationMd5 {
	return obj.ospfv2AuthenticationMd5Slice
}

func (obj *ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter) Add() Ospfv2AuthenticationMd5 {
	newObj := &otg.Ospfv2AuthenticationMd5{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2AuthenticationMd5{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2AuthenticationMd5Slice = append(obj.ospfv2AuthenticationMd5Slice, newLibObj)
	return newLibObj
}

func (obj *ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter) Append(items ...Ospfv2AuthenticationMd5) Ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2AuthenticationMd5Slice = append(obj.ospfv2AuthenticationMd5Slice, item)
	}
	return obj
}

func (obj *ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter) Set(index int, newObj Ospfv2AuthenticationMd5) Ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2AuthenticationMd5Slice[index] = newObj
	return obj
}
func (obj *ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter) Clear() Ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2AuthenticationMd5{}
		obj.ospfv2AuthenticationMd5Slice = []Ospfv2AuthenticationMd5{}
	}
	return obj
}
func (obj *ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter) clearHolderSlice() Ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter {
	if len(obj.ospfv2AuthenticationMd5Slice) > 0 {
		obj.ospfv2AuthenticationMd5Slice = []Ospfv2AuthenticationMd5{}
	}
	return obj
}
func (obj *ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter) appendHolderSlice(item Ospfv2AuthenticationMd5) Ospfv2InterfaceAuthenticationOspfv2AuthenticationMd5Iter {
	obj.ospfv2AuthenticationMd5Slice = append(obj.ospfv2AuthenticationMd5Slice, item)
	return obj
}

// The 8 Byte authentication field in the OSPF packet.
// ClearText returns a string
func (obj *ospfv2InterfaceAuthentication) ClearText() string {

	if obj.obj.ClearText == nil {
		obj.setChoice(Ospfv2InterfaceAuthenticationChoice.CLEAR_TEXT)
	}

	return *obj.obj.ClearText

}

// The 8 Byte authentication field in the OSPF packet.
// ClearText returns a string
func (obj *ospfv2InterfaceAuthentication) HasClearText() bool {
	return obj.obj.ClearText != nil
}

// The 8 Byte authentication field in the OSPF packet.
// SetClearText sets the string value in the Ospfv2InterfaceAuthentication object
func (obj *ospfv2InterfaceAuthentication) SetClearText(value string) Ospfv2InterfaceAuthentication {
	obj.setChoice(Ospfv2InterfaceAuthenticationChoice.CLEAR_TEXT)
	obj.obj.ClearText = &value
	return obj
}

func (obj *ospfv2InterfaceAuthentication) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Md5S) != 0 {

		if set_default {
			obj.Md5S().clearHolderSlice()
			for _, item := range obj.obj.Md5S {
				obj.Md5S().appendHolderSlice(&ospfv2AuthenticationMd5{obj: item})
			}
		}
		for _, item := range obj.Md5S().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.ClearText != nil {

		if len(*obj.obj.ClearText) < 1 || len(*obj.obj.ClearText) > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of Ospfv2InterfaceAuthentication.ClearText <= 8 but Got %d",
					len(*obj.obj.ClearText)))
		}

	}

}

func (obj *ospfv2InterfaceAuthentication) setDefault() {
	var choices_set int = 0
	var choice Ospfv2InterfaceAuthenticationChoiceEnum

	if obj.obj.ClearText != nil {
		choices_set += 1
		choice = Ospfv2InterfaceAuthenticationChoice.CLEAR_TEXT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Ospfv2InterfaceAuthenticationChoice.CLEAR_TEXT)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Ospfv2InterfaceAuthentication")
			}
		} else {
			intVal := otg.Ospfv2InterfaceAuthentication_Choice_Enum_value[string(choice)]
			enumValue := otg.Ospfv2InterfaceAuthentication_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

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
	// getter for Md5 to set choice.
	Md5()
	// Md5KeyId returns string, set in Ospfv2InterfaceAuthentication.
	Md5KeyId() string
	// SetMd5KeyId assigns string provided by user to Ospfv2InterfaceAuthentication
	SetMd5KeyId(value string) Ospfv2InterfaceAuthentication
	// HasMd5KeyId checks if Md5KeyId has been set in Ospfv2InterfaceAuthentication
	HasMd5KeyId() bool
	// Md5Key returns string, set in Ospfv2InterfaceAuthentication.
	Md5Key() string
	// SetMd5Key assigns string provided by user to Ospfv2InterfaceAuthentication
	SetMd5Key(value string) Ospfv2InterfaceAuthentication
	// HasMd5Key checks if Md5Key has been set in Ospfv2InterfaceAuthentication
	HasMd5Key() bool
	// ClearText returns string, set in Ospfv2InterfaceAuthentication.
	ClearText() string
	// SetClearText assigns string provided by user to Ospfv2InterfaceAuthentication
	SetClearText(value string) Ospfv2InterfaceAuthentication
	// HasClearText checks if ClearText has been set in Ospfv2InterfaceAuthentication
	HasClearText() bool
}

type Ospfv2InterfaceAuthenticationChoiceEnum string

// Enum of Choice on Ospfv2InterfaceAuthentication
var Ospfv2InterfaceAuthenticationChoice = struct {
	MD5        Ospfv2InterfaceAuthenticationChoiceEnum
	CLEAR_TEXT Ospfv2InterfaceAuthenticationChoiceEnum
}{
	MD5:        Ospfv2InterfaceAuthenticationChoiceEnum("md5"),
	CLEAR_TEXT: Ospfv2InterfaceAuthenticationChoiceEnum("clear_text"),
}

func (obj *ospfv2InterfaceAuthentication) Choice() Ospfv2InterfaceAuthenticationChoiceEnum {
	return Ospfv2InterfaceAuthenticationChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Md5 to set choice
func (obj *ospfv2InterfaceAuthentication) Md5() {
	obj.setChoice(Ospfv2InterfaceAuthenticationChoice.MD5)
}

// The authentication method.
// - md5 - Cryptographic authentication. If the authentication type is of 'md5' then 'md5_key_id' and 'md5_key'
// both are to be configured. A shared secret key is configured in all routers attached to a common network/subnet.
// For each OSPF protocol packet, the key is used to generate/verify a "message digest" that is appended to the end
// of the OSPF packet.
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
	return obj
}

// The unique MD5 Key Identifier per-interface.
// Md5KeyId returns a string
func (obj *ospfv2InterfaceAuthentication) Md5KeyId() string {

	return *obj.obj.Md5KeyId

}

// The unique MD5 Key Identifier per-interface.
// Md5KeyId returns a string
func (obj *ospfv2InterfaceAuthentication) HasMd5KeyId() bool {
	return obj.obj.Md5KeyId != nil
}

// The unique MD5 Key Identifier per-interface.
// SetMd5KeyId sets the string value in the Ospfv2InterfaceAuthentication object
func (obj *ospfv2InterfaceAuthentication) SetMd5KeyId(value string) Ospfv2InterfaceAuthentication {

	obj.obj.Md5KeyId = &value
	return obj
}

// An alphanumeric secret used to generate the 16 byte MD5 hash value added
// to the OSPFv2 PDU in the Authentication TLV.
// Md5Key returns a string
func (obj *ospfv2InterfaceAuthentication) Md5Key() string {

	return *obj.obj.Md5Key

}

// An alphanumeric secret used to generate the 16 byte MD5 hash value added
// to the OSPFv2 PDU in the Authentication TLV.
// Md5Key returns a string
func (obj *ospfv2InterfaceAuthentication) HasMd5Key() bool {
	return obj.obj.Md5Key != nil
}

// An alphanumeric secret used to generate the 16 byte MD5 hash value added
// to the OSPFv2 PDU in the Authentication TLV.
// SetMd5Key sets the string value in the Ospfv2InterfaceAuthentication object
func (obj *ospfv2InterfaceAuthentication) SetMd5Key(value string) Ospfv2InterfaceAuthentication {

	obj.obj.Md5Key = &value
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

	if obj.obj.Md5KeyId != nil {

		if len(*obj.obj.Md5KeyId) < 1 || len(*obj.obj.Md5KeyId) > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of Ospfv2InterfaceAuthentication.Md5KeyId <= 255 but Got %d",
					len(*obj.obj.Md5KeyId)))
		}

	}

	if obj.obj.Md5Key != nil {

		if len(*obj.obj.Md5Key) < 1 || len(*obj.obj.Md5Key) > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of Ospfv2InterfaceAuthentication.Md5Key <= 255 but Got %d",
					len(*obj.obj.Md5Key)))
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
	if choices_set == 1 && choice != "" {
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

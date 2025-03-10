package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityKeyGenerationProtocol *****
type secureEntityKeyGenerationProtocol struct {
	validation
	obj             *otg.SecureEntityKeyGenerationProtocol
	marshaller      marshalSecureEntityKeyGenerationProtocol
	unMarshaller    unMarshalSecureEntityKeyGenerationProtocol
	mkaHolder       Mka
	staticKeyHolder SecureEntityStaticKey
}

func NewSecureEntityKeyGenerationProtocol() SecureEntityKeyGenerationProtocol {
	obj := secureEntityKeyGenerationProtocol{obj: &otg.SecureEntityKeyGenerationProtocol{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityKeyGenerationProtocol) msg() *otg.SecureEntityKeyGenerationProtocol {
	return obj.obj
}

func (obj *secureEntityKeyGenerationProtocol) setMsg(msg *otg.SecureEntityKeyGenerationProtocol) SecureEntityKeyGenerationProtocol {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityKeyGenerationProtocol struct {
	obj *secureEntityKeyGenerationProtocol
}

type marshalSecureEntityKeyGenerationProtocol interface {
	// ToProto marshals SecureEntityKeyGenerationProtocol to protobuf object *otg.SecureEntityKeyGenerationProtocol
	ToProto() (*otg.SecureEntityKeyGenerationProtocol, error)
	// ToPbText marshals SecureEntityKeyGenerationProtocol to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityKeyGenerationProtocol to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityKeyGenerationProtocol to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals SecureEntityKeyGenerationProtocol to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalsecureEntityKeyGenerationProtocol struct {
	obj *secureEntityKeyGenerationProtocol
}

type unMarshalSecureEntityKeyGenerationProtocol interface {
	// FromProto unmarshals SecureEntityKeyGenerationProtocol from protobuf object *otg.SecureEntityKeyGenerationProtocol
	FromProto(msg *otg.SecureEntityKeyGenerationProtocol) (SecureEntityKeyGenerationProtocol, error)
	// FromPbText unmarshals SecureEntityKeyGenerationProtocol from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityKeyGenerationProtocol from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityKeyGenerationProtocol from JSON text
	FromJson(value string) error
}

func (obj *secureEntityKeyGenerationProtocol) Marshal() marshalSecureEntityKeyGenerationProtocol {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityKeyGenerationProtocol{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityKeyGenerationProtocol) Unmarshal() unMarshalSecureEntityKeyGenerationProtocol {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityKeyGenerationProtocol{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityKeyGenerationProtocol) ToProto() (*otg.SecureEntityKeyGenerationProtocol, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityKeyGenerationProtocol) FromProto(msg *otg.SecureEntityKeyGenerationProtocol) (SecureEntityKeyGenerationProtocol, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityKeyGenerationProtocol) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityKeyGenerationProtocol) FromPbText(value string) error {
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

func (m *marshalsecureEntityKeyGenerationProtocol) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityKeyGenerationProtocol) FromYaml(value string) error {
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

func (m *marshalsecureEntityKeyGenerationProtocol) ToJsonRaw() (string, error) {
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

func (m *marshalsecureEntityKeyGenerationProtocol) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityKeyGenerationProtocol) FromJson(value string) error {
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

func (obj *secureEntityKeyGenerationProtocol) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityKeyGenerationProtocol) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityKeyGenerationProtocol) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityKeyGenerationProtocol) Clone() (SecureEntityKeyGenerationProtocol, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityKeyGenerationProtocol()
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

func (obj *secureEntityKeyGenerationProtocol) setNil() {
	obj.mkaHolder = nil
	obj.staticKeyHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// SecureEntityKeyGenerationProtocol is container of Key generation protocol configuration.
type SecureEntityKeyGenerationProtocol interface {
	Validation
	// msg marshals SecureEntityKeyGenerationProtocol to protobuf object *otg.SecureEntityKeyGenerationProtocol
	// and doesn't set defaults
	msg() *otg.SecureEntityKeyGenerationProtocol
	// setMsg unmarshals SecureEntityKeyGenerationProtocol from protobuf object *otg.SecureEntityKeyGenerationProtocol
	// and doesn't set defaults
	setMsg(*otg.SecureEntityKeyGenerationProtocol) SecureEntityKeyGenerationProtocol
	// provides marshal interface
	Marshal() marshalSecureEntityKeyGenerationProtocol
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityKeyGenerationProtocol
	// validate validates SecureEntityKeyGenerationProtocol
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityKeyGenerationProtocol, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns SecureEntityKeyGenerationProtocolChoiceEnum, set in SecureEntityKeyGenerationProtocol
	Choice() SecureEntityKeyGenerationProtocolChoiceEnum
	// setChoice assigns SecureEntityKeyGenerationProtocolChoiceEnum provided by user to SecureEntityKeyGenerationProtocol
	setChoice(value SecureEntityKeyGenerationProtocolChoiceEnum) SecureEntityKeyGenerationProtocol
	// HasChoice checks if Choice has been set in SecureEntityKeyGenerationProtocol
	HasChoice() bool
	// Mka returns Mka, set in SecureEntityKeyGenerationProtocol.
	// Mka is configuration of a MKA Key Agreement Entity (KaY). Reference https://1.ieee802.org/security/802-1x/.
	Mka() Mka
	// SetMka assigns Mka provided by user to SecureEntityKeyGenerationProtocol.
	// Mka is configuration of a MKA Key Agreement Entity (KaY). Reference https://1.ieee802.org/security/802-1x/.
	SetMka(value Mka) SecureEntityKeyGenerationProtocol
	// HasMka checks if Mka has been set in SecureEntityKeyGenerationProtocol
	HasMka() bool
	// StaticKey returns SecureEntityStaticKey, set in SecureEntityKeyGenerationProtocol.
	// SecureEntityStaticKey is a container of static key properties for a secure entity(SecY). This configuration is applicable when no dynamic key management protocol i.e. MACsec key agreement(MKA) is configured. If MKA is configured, any static key configuration is not applicable.
	StaticKey() SecureEntityStaticKey
	// SetStaticKey assigns SecureEntityStaticKey provided by user to SecureEntityKeyGenerationProtocol.
	// SecureEntityStaticKey is a container of static key properties for a secure entity(SecY). This configuration is applicable when no dynamic key management protocol i.e. MACsec key agreement(MKA) is configured. If MKA is configured, any static key configuration is not applicable.
	SetStaticKey(value SecureEntityStaticKey) SecureEntityKeyGenerationProtocol
	// HasStaticKey checks if StaticKey has been set in SecureEntityKeyGenerationProtocol
	HasStaticKey() bool
	setNil()
}

type SecureEntityKeyGenerationProtocolChoiceEnum string

// Enum of Choice on SecureEntityKeyGenerationProtocol
var SecureEntityKeyGenerationProtocolChoice = struct {
	MKA        SecureEntityKeyGenerationProtocolChoiceEnum
	STATIC_KEY SecureEntityKeyGenerationProtocolChoiceEnum
}{
	MKA:        SecureEntityKeyGenerationProtocolChoiceEnum("mka"),
	STATIC_KEY: SecureEntityKeyGenerationProtocolChoiceEnum("static_key"),
}

func (obj *secureEntityKeyGenerationProtocol) Choice() SecureEntityKeyGenerationProtocolChoiceEnum {
	return SecureEntityKeyGenerationProtocolChoiceEnum(obj.obj.Choice.Enum().String())
}

// Key generation protocol choices. Choose "mka" for dynamic key distribution using MACsec key agreement(MKA) protocol. Choose "static_key" for static configuration of secure association key(SAK).
// Choice returns a string
func (obj *secureEntityKeyGenerationProtocol) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *secureEntityKeyGenerationProtocol) setChoice(value SecureEntityKeyGenerationProtocolChoiceEnum) SecureEntityKeyGenerationProtocol {
	intValue, ok := otg.SecureEntityKeyGenerationProtocol_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on SecureEntityKeyGenerationProtocolChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.SecureEntityKeyGenerationProtocol_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.StaticKey = nil
	obj.staticKeyHolder = nil
	obj.obj.Mka = nil
	obj.mkaHolder = nil

	if value == SecureEntityKeyGenerationProtocolChoice.MKA {
		obj.obj.Mka = NewMka().msg()
	}

	if value == SecureEntityKeyGenerationProtocolChoice.STATIC_KEY {
		obj.obj.StaticKey = NewSecureEntityStaticKey().msg()
	}

	return obj
}

// This contains the properties of Key Agreement Entity (KaY) in MKA supplicant.
// Mka returns a Mka
func (obj *secureEntityKeyGenerationProtocol) Mka() Mka {
	if obj.obj.Mka == nil {
		obj.setChoice(SecureEntityKeyGenerationProtocolChoice.MKA)
	}
	if obj.mkaHolder == nil {
		obj.mkaHolder = &mka{obj: obj.obj.Mka}
	}
	return obj.mkaHolder
}

// This contains the properties of Key Agreement Entity (KaY) in MKA supplicant.
// Mka returns a Mka
func (obj *secureEntityKeyGenerationProtocol) HasMka() bool {
	return obj.obj.Mka != nil
}

// This contains the properties of Key Agreement Entity (KaY) in MKA supplicant.
// SetMka sets the Mka value in the SecureEntityKeyGenerationProtocol object
func (obj *secureEntityKeyGenerationProtocol) SetMka(value Mka) SecureEntityKeyGenerationProtocol {
	obj.setChoice(SecureEntityKeyGenerationProtocolChoice.MKA)
	obj.mkaHolder = nil
	obj.obj.Mka = value.msg()

	return obj
}

// Static key properties properties of SecY. Static key is used in absence MKA.
// StaticKey returns a SecureEntityStaticKey
func (obj *secureEntityKeyGenerationProtocol) StaticKey() SecureEntityStaticKey {
	if obj.obj.StaticKey == nil {
		obj.setChoice(SecureEntityKeyGenerationProtocolChoice.STATIC_KEY)
	}
	if obj.staticKeyHolder == nil {
		obj.staticKeyHolder = &secureEntityStaticKey{obj: obj.obj.StaticKey}
	}
	return obj.staticKeyHolder
}

// Static key properties properties of SecY. Static key is used in absence MKA.
// StaticKey returns a SecureEntityStaticKey
func (obj *secureEntityKeyGenerationProtocol) HasStaticKey() bool {
	return obj.obj.StaticKey != nil
}

// Static key properties properties of SecY. Static key is used in absence MKA.
// SetStaticKey sets the SecureEntityStaticKey value in the SecureEntityKeyGenerationProtocol object
func (obj *secureEntityKeyGenerationProtocol) SetStaticKey(value SecureEntityStaticKey) SecureEntityKeyGenerationProtocol {
	obj.setChoice(SecureEntityKeyGenerationProtocolChoice.STATIC_KEY)
	obj.staticKeyHolder = nil
	obj.obj.StaticKey = value.msg()

	return obj
}

func (obj *secureEntityKeyGenerationProtocol) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Mka != nil {

		obj.Mka().validateObj(vObj, set_default)
	}

	if obj.obj.StaticKey != nil {

		obj.StaticKey().validateObj(vObj, set_default)
	}

}

func (obj *secureEntityKeyGenerationProtocol) setDefault() {
	var choices_set int = 0
	var choice SecureEntityKeyGenerationProtocolChoiceEnum

	if obj.obj.Mka != nil {
		choices_set += 1
		choice = SecureEntityKeyGenerationProtocolChoice.MKA
	}

	if obj.obj.StaticKey != nil {
		choices_set += 1
		choice = SecureEntityKeyGenerationProtocolChoice.STATIC_KEY
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(SecureEntityKeyGenerationProtocolChoice.MKA)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in SecureEntityKeyGenerationProtocol")
			}
		} else {
			intVal := otg.SecureEntityKeyGenerationProtocol_Choice_Enum_value[string(choice)]
			enumValue := otg.SecureEntityKeyGenerationProtocol_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

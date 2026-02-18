package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityCryptoEngine *****
type secureEntityCryptoEngine struct {
	validation
	obj                  *otg.SecureEntityCryptoEngine
	marshaller           marshalSecureEntityCryptoEngine
	unMarshaller         unMarshalSecureEntityCryptoEngine
	encryptOnlyHolder    SecureEntityCryptoEngineEncryptOnly
	encryptDecryptHolder SecureEntityCryptoEngineEncryptDecrypt
}

func NewSecureEntityCryptoEngine() SecureEntityCryptoEngine {
	obj := secureEntityCryptoEngine{obj: &otg.SecureEntityCryptoEngine{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityCryptoEngine) msg() *otg.SecureEntityCryptoEngine {
	return obj.obj
}

func (obj *secureEntityCryptoEngine) setMsg(msg *otg.SecureEntityCryptoEngine) SecureEntityCryptoEngine {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityCryptoEngine struct {
	obj *secureEntityCryptoEngine
}

type marshalSecureEntityCryptoEngine interface {
	// ToProto marshals SecureEntityCryptoEngine to protobuf object *otg.SecureEntityCryptoEngine
	ToProto() (*otg.SecureEntityCryptoEngine, error)
	// ToPbText marshals SecureEntityCryptoEngine to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityCryptoEngine to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityCryptoEngine to JSON text
	ToJson() (string, error)
}

type unMarshalsecureEntityCryptoEngine struct {
	obj *secureEntityCryptoEngine
}

type unMarshalSecureEntityCryptoEngine interface {
	// FromProto unmarshals SecureEntityCryptoEngine from protobuf object *otg.SecureEntityCryptoEngine
	FromProto(msg *otg.SecureEntityCryptoEngine) (SecureEntityCryptoEngine, error)
	// FromPbText unmarshals SecureEntityCryptoEngine from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityCryptoEngine from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityCryptoEngine from JSON text
	FromJson(value string) error
}

func (obj *secureEntityCryptoEngine) Marshal() marshalSecureEntityCryptoEngine {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityCryptoEngine{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityCryptoEngine) Unmarshal() unMarshalSecureEntityCryptoEngine {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityCryptoEngine{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityCryptoEngine) ToProto() (*otg.SecureEntityCryptoEngine, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityCryptoEngine) FromProto(msg *otg.SecureEntityCryptoEngine) (SecureEntityCryptoEngine, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityCryptoEngine) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngine) FromPbText(value string) error {
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

func (m *marshalsecureEntityCryptoEngine) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngine) FromYaml(value string) error {
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

func (m *marshalsecureEntityCryptoEngine) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngine) FromJson(value string) error {
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

func (obj *secureEntityCryptoEngine) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngine) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngine) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityCryptoEngine) Clone() (SecureEntityCryptoEngine, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityCryptoEngine()
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

func (obj *secureEntityCryptoEngine) setNil() {
	obj.encryptOnlyHolder = nil
	obj.encryptDecryptHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// SecureEntityCryptoEngine is a container of crypto engine properties of a SecY.
type SecureEntityCryptoEngine interface {
	Validation
	// msg marshals SecureEntityCryptoEngine to protobuf object *otg.SecureEntityCryptoEngine
	// and doesn't set defaults
	msg() *otg.SecureEntityCryptoEngine
	// setMsg unmarshals SecureEntityCryptoEngine from protobuf object *otg.SecureEntityCryptoEngine
	// and doesn't set defaults
	setMsg(*otg.SecureEntityCryptoEngine) SecureEntityCryptoEngine
	// provides marshal interface
	Marshal() marshalSecureEntityCryptoEngine
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityCryptoEngine
	// validate validates SecureEntityCryptoEngine
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityCryptoEngine, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns SecureEntityCryptoEngineChoiceEnum, set in SecureEntityCryptoEngine
	Choice() SecureEntityCryptoEngineChoiceEnum
	// setChoice assigns SecureEntityCryptoEngineChoiceEnum provided by user to SecureEntityCryptoEngine
	setChoice(value SecureEntityCryptoEngineChoiceEnum) SecureEntityCryptoEngine
	// HasChoice checks if Choice has been set in SecureEntityCryptoEngine
	HasChoice() bool
	// EncryptOnly returns SecureEntityCryptoEngineEncryptOnly, set in SecureEntityCryptoEngine.
	// SecureEntityCryptoEngineEncryptOnly is the container for encrypt only engine configuration.
	EncryptOnly() SecureEntityCryptoEngineEncryptOnly
	// SetEncryptOnly assigns SecureEntityCryptoEngineEncryptOnly provided by user to SecureEntityCryptoEngine.
	// SecureEntityCryptoEngineEncryptOnly is the container for encrypt only engine configuration.
	SetEncryptOnly(value SecureEntityCryptoEngineEncryptOnly) SecureEntityCryptoEngine
	// HasEncryptOnly checks if EncryptOnly has been set in SecureEntityCryptoEngine
	HasEncryptOnly() bool
	// EncryptDecrypt returns SecureEntityCryptoEngineEncryptDecrypt, set in SecureEntityCryptoEngine.
	// SecureEntityCryptoEngineEncryptDecrypt is the container for encrypt and decrypt engine configuration.
	EncryptDecrypt() SecureEntityCryptoEngineEncryptDecrypt
	// SetEncryptDecrypt assigns SecureEntityCryptoEngineEncryptDecrypt provided by user to SecureEntityCryptoEngine.
	// SecureEntityCryptoEngineEncryptDecrypt is the container for encrypt and decrypt engine configuration.
	SetEncryptDecrypt(value SecureEntityCryptoEngineEncryptDecrypt) SecureEntityCryptoEngine
	// HasEncryptDecrypt checks if EncryptDecrypt has been set in SecureEntityCryptoEngine
	HasEncryptDecrypt() bool
	setNil()
}

type SecureEntityCryptoEngineChoiceEnum string

// Enum of Choice on SecureEntityCryptoEngine
var SecureEntityCryptoEngineChoice = struct {
	ENCRYPT_ONLY    SecureEntityCryptoEngineChoiceEnum
	ENCRYPT_DECRYPT SecureEntityCryptoEngineChoiceEnum
}{
	ENCRYPT_ONLY:    SecureEntityCryptoEngineChoiceEnum("encrypt_only"),
	ENCRYPT_DECRYPT: SecureEntityCryptoEngineChoiceEnum("encrypt_decrypt"),
}

func (obj *secureEntityCryptoEngine) Choice() SecureEntityCryptoEngineChoiceEnum {
	return SecureEntityCryptoEngineChoiceEnum(obj.obj.Choice.Enum().String())
}

// Engine type based on encryption and/ or decryption capability. Supported types: 1) encrypt_only - engine can only encrypt transmitted packets but it cannot decrypt packets upon arrival. As the packets cannot be decrypted on arrival, such packets cannot be delivered to the receiving device. Hence only stateless traffic can be sent. 2) encrypt_decrypt - engine can both encrypt transmitted packets and decrypt packets on arrival. Such engine can have hardware acceleration for faster encryption/ decryption. As both encryption and decryption are possible, stateful (e.g. TCP) traffic can be sent/ received.
// Choice returns a string
func (obj *secureEntityCryptoEngine) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *secureEntityCryptoEngine) setChoice(value SecureEntityCryptoEngineChoiceEnum) SecureEntityCryptoEngine {
	intValue, ok := otg.SecureEntityCryptoEngine_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on SecureEntityCryptoEngineChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.SecureEntityCryptoEngine_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.EncryptDecrypt = nil
	obj.encryptDecryptHolder = nil
	obj.obj.EncryptOnly = nil
	obj.encryptOnlyHolder = nil

	if value == SecureEntityCryptoEngineChoice.ENCRYPT_ONLY {
		obj.obj.EncryptOnly = NewSecureEntityCryptoEngineEncryptOnly().msg()
	}

	if value == SecureEntityCryptoEngineChoice.ENCRYPT_DECRYPT {
		obj.obj.EncryptDecrypt = NewSecureEntityCryptoEngineEncryptDecrypt().msg()
	}

	return obj
}

// description is TBD
// EncryptOnly returns a SecureEntityCryptoEngineEncryptOnly
func (obj *secureEntityCryptoEngine) EncryptOnly() SecureEntityCryptoEngineEncryptOnly {
	if obj.obj.EncryptOnly == nil {
		obj.setChoice(SecureEntityCryptoEngineChoice.ENCRYPT_ONLY)
	}
	if obj.encryptOnlyHolder == nil {
		obj.encryptOnlyHolder = &secureEntityCryptoEngineEncryptOnly{obj: obj.obj.EncryptOnly}
	}
	return obj.encryptOnlyHolder
}

// description is TBD
// EncryptOnly returns a SecureEntityCryptoEngineEncryptOnly
func (obj *secureEntityCryptoEngine) HasEncryptOnly() bool {
	return obj.obj.EncryptOnly != nil
}

// description is TBD
// SetEncryptOnly sets the SecureEntityCryptoEngineEncryptOnly value in the SecureEntityCryptoEngine object
func (obj *secureEntityCryptoEngine) SetEncryptOnly(value SecureEntityCryptoEngineEncryptOnly) SecureEntityCryptoEngine {
	obj.setChoice(SecureEntityCryptoEngineChoice.ENCRYPT_ONLY)
	obj.encryptOnlyHolder = nil
	obj.obj.EncryptOnly = value.msg()

	return obj
}

// description is TBD
// EncryptDecrypt returns a SecureEntityCryptoEngineEncryptDecrypt
func (obj *secureEntityCryptoEngine) EncryptDecrypt() SecureEntityCryptoEngineEncryptDecrypt {
	if obj.obj.EncryptDecrypt == nil {
		obj.setChoice(SecureEntityCryptoEngineChoice.ENCRYPT_DECRYPT)
	}
	if obj.encryptDecryptHolder == nil {
		obj.encryptDecryptHolder = &secureEntityCryptoEngineEncryptDecrypt{obj: obj.obj.EncryptDecrypt}
	}
	return obj.encryptDecryptHolder
}

// description is TBD
// EncryptDecrypt returns a SecureEntityCryptoEngineEncryptDecrypt
func (obj *secureEntityCryptoEngine) HasEncryptDecrypt() bool {
	return obj.obj.EncryptDecrypt != nil
}

// description is TBD
// SetEncryptDecrypt sets the SecureEntityCryptoEngineEncryptDecrypt value in the SecureEntityCryptoEngine object
func (obj *secureEntityCryptoEngine) SetEncryptDecrypt(value SecureEntityCryptoEngineEncryptDecrypt) SecureEntityCryptoEngine {
	obj.setChoice(SecureEntityCryptoEngineChoice.ENCRYPT_DECRYPT)
	obj.encryptDecryptHolder = nil
	obj.obj.EncryptDecrypt = value.msg()

	return obj
}

func (obj *secureEntityCryptoEngine) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.EncryptOnly != nil {

		obj.EncryptOnly().validateObj(vObj, set_default)
	}

	if obj.obj.EncryptDecrypt != nil {

		obj.EncryptDecrypt().validateObj(vObj, set_default)
	}

}

func (obj *secureEntityCryptoEngine) setDefault() {
	var choices_set int = 0
	var choice SecureEntityCryptoEngineChoiceEnum

	if obj.obj.EncryptOnly != nil {
		choices_set += 1
		choice = SecureEntityCryptoEngineChoice.ENCRYPT_ONLY
	}

	if obj.obj.EncryptDecrypt != nil {
		choices_set += 1
		choice = SecureEntityCryptoEngineChoice.ENCRYPT_DECRYPT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(SecureEntityCryptoEngineChoice.ENCRYPT_ONLY)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in SecureEntityCryptoEngine")
			}
		} else {
			intVal := otg.SecureEntityCryptoEngine_Choice_Enum_value[string(choice)]
			enumValue := otg.SecureEntityCryptoEngine_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

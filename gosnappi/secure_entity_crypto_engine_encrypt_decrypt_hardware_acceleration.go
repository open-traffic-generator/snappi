package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration *****
type secureEntityCryptoEngineEncryptDecryptHardwareAcceleration struct {
	validation
	obj          *otg.SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	marshaller   marshalSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	unMarshaller unMarshalSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
}

func NewSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration() SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration {
	obj := secureEntityCryptoEngineEncryptDecryptHardwareAcceleration{obj: &otg.SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAcceleration) msg() *otg.SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration {
	return obj.obj
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAcceleration) setMsg(msg *otg.SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityCryptoEngineEncryptDecryptHardwareAcceleration struct {
	obj *secureEntityCryptoEngineEncryptDecryptHardwareAcceleration
}

type marshalSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration interface {
	// ToProto marshals SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration to protobuf object *otg.SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	ToProto() (*otg.SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration, error)
	// ToPbText marshals SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration to JSON text
	ToJson() (string, error)
}

type unMarshalsecureEntityCryptoEngineEncryptDecryptHardwareAcceleration struct {
	obj *secureEntityCryptoEngineEncryptDecryptHardwareAcceleration
}

type unMarshalSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration interface {
	// FromProto unmarshals SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration from protobuf object *otg.SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	FromProto(msg *otg.SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) (SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration, error)
	// FromPbText unmarshals SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration from JSON text
	FromJson(value string) error
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAcceleration) Marshal() marshalSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityCryptoEngineEncryptDecryptHardwareAcceleration{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAcceleration) Unmarshal() unMarshalSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityCryptoEngineEncryptDecryptHardwareAcceleration{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) ToProto() (*otg.SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) FromProto(msg *otg.SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) (SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) FromPbText(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) FromYaml(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) FromJson(value string) error {
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

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAcceleration) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAcceleration) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAcceleration) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAcceleration) Clone() (SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration()
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

// SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration is hardware acceleration configuration for offloading MACsec processing to hardware.
type SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration interface {
	Validation
	// msg marshals SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration to protobuf object *otg.SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	// and doesn't set defaults
	msg() *otg.SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	// setMsg unmarshals SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration from protobuf object *otg.SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	// and doesn't set defaults
	setMsg(*otg.SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	// provides marshal interface
	Marshal() marshalSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	// validate validates SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum, set in SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	Choice() SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum
	// setChoice assigns SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum provided by user to SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	setChoice(value SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum) SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	// HasChoice checks if Choice has been set in SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	HasChoice() bool
	// getter for None to set choice.
	None()
	// getter for InlineCrypto to set choice.
	InlineCrypto()
}

type SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum string

// Enum of Choice on SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
var SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoice = struct {
	NONE          SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum
	INLINE_CRYPTO SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum
}{
	NONE:          SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum("none"),
	INLINE_CRYPTO: SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum("inline_crypto"),
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAcceleration) Choice() SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum {
	return SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for None to set choice
func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAcceleration) None() {
	obj.setChoice(SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoice.NONE)
}

// getter for InlineCrypto to set choice
func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAcceleration) InlineCrypto() {
	obj.setChoice(SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoice.INLINE_CRYPTO)
}

// Hardware acceleration types. Per port hardware acceleration configuration is availble at options -> per_port_options -> protocols -> macsec -> hardware_accelertation.
// Choice returns a string
func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAcceleration) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAcceleration) setChoice(value SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum) SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration {
	intValue, ok := otg.SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAcceleration) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAcceleration) setDefault() {
	var choices_set int = 0
	var choice SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoice.NONE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration")
			}
		} else {
			intVal := otg.SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration_Choice_Enum_value[string(choice)]
			enumValue := otg.SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

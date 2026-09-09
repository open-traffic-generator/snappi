package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration *****
type lagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration struct {
	validation
	obj          *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	marshaller   marshalLagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	unMarshaller unMarshalLagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
}

func NewLagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration() LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration {
	obj := lagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration{obj: &otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration{}}
	obj.setDefault()
	return &obj
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) msg() *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration {
	return obj.obj
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) setMsg(msg *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration struct {
	obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
}

type marshalLagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration interface {
	// ToProto marshals LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration to protobuf object *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	ToProto() (*otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration, error)
	// ToPbText marshals LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration to JSON text
	ToJson() (string, error)
}

type unMarshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration struct {
	obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
}

type unMarshalLagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration interface {
	// FromProto unmarshals LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration from protobuf object *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	FromProto(msg *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) (LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration, error)
	// FromPbText unmarshals LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration from JSON text
	FromJson(value string) error
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) Marshal() marshalLagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) Unmarshal() unMarshalLagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) ToProto() (*otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) FromProto(msg *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) (LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) ToPbText() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) FromPbText(value string) error {
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

func (m *marshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) ToYaml() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) FromYaml(value string) error {
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

func (m *marshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) ToJson() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) FromJson(value string) error {
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

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) Clone() (LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration()
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

// LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration is hardware acceleration configuration for offloading MACsec processing to hardware.
type LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration interface {
	Validation
	// msg marshals LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration to protobuf object *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	// and doesn't set defaults
	msg() *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	// setMsg unmarshals LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration from protobuf object *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	// and doesn't set defaults
	setMsg(*otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	// provides marshal interface
	Marshal() marshalLagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	// provides unmarshal interface
	Unmarshal() unMarshalLagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	// validate validates LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum, set in LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	Choice() LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum
	// setChoice assigns LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum provided by user to LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	setChoice(value LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum) LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	// HasChoice checks if Choice has been set in LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	HasChoice() bool
	// getter for InlineCrypto to set choice.
	InlineCrypto()
	// getter for None to set choice.
	None()
}

type LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum string

// Enum of Choice on LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
var LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoice = struct {
	NONE          LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum
	INLINE_CRYPTO LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum
}{
	NONE:          LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum("none"),
	INLINE_CRYPTO: LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum("inline_crypto"),
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) Choice() LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum {
	return LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for InlineCrypto to set choice
func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) InlineCrypto() {
	obj.setChoice(LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoice.INLINE_CRYPTO)
}

// getter for None to set choice
func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) None() {
	obj.setChoice(LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoice.NONE)
}

// Hardware acceleration types.
// Choice returns a string
func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) setChoice(value LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum) LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration {
	intValue, ok := otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) setDefault() {
	var choices_set int = 0
	var choice LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationChoice.NONE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration")
			}
		} else {
			intVal := otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration_Choice_Enum_value[string(choice)]
			enumValue := otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

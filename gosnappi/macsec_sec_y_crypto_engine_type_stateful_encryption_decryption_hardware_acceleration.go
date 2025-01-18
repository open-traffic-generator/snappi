package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration *****
type macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration struct {
	validation
	obj                *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	marshaller         marshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	unMarshaller       unMarshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	inlineCryptoHolder MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
}

func NewMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration() MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration {
	obj := macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration{obj: &otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) msg() *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration {
	return obj.obj
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) setMsg(msg *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration struct {
	obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
}

type marshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration interface {
	// ToProto marshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration to protobuf object *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	ToProto() (*otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration, error)
	// ToPbText marshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration struct {
	obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
}

type unMarshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration interface {
	// FromProto unmarshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration from protobuf object *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	FromProto(msg *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) (MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration, error)
	// FromPbText unmarshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) Marshal() marshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) Unmarshal() unMarshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) ToProto() (*otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) FromProto(msg *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) (MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) FromPbText(value string) error {
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

func (m *marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) FromYaml(value string) error {
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

func (m *marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) FromJson(value string) error {
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

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) Clone() (MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration()
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

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) setNil() {
	obj.inlineCryptoHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration is hardware acceleration settings.
type MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration interface {
	Validation
	// msg marshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration to protobuf object *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	// and doesn't set defaults
	msg() *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	// setMsg unmarshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration from protobuf object *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	// provides marshal interface
	Marshal() marshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	// validate validates MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum, set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	Choice() MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum
	// setChoice assigns MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum provided by user to MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	setChoice(value MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	// HasChoice checks if Choice has been set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	HasChoice() bool
	// getter for None to set choice.
	None()
	// InlineCrypto returns MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto, set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration.
	// MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto is inline cryto engine settings.
	InlineCrypto() MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// SetInlineCrypto assigns MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto provided by user to MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration.
	// MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto is inline cryto engine settings.
	SetInlineCrypto(value MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	// HasInlineCrypto checks if InlineCrypto has been set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	HasInlineCrypto() bool
	setNil()
}

type MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum string

// Enum of Choice on MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
var MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoice = struct {
	NONE          MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum
	INLINE_CRYPTO MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum
}{
	NONE:          MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum("none"),
	INLINE_CRYPTO: MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum("inline_crypto"),
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) Choice() MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum {
	return MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for None to set choice
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) None() {
	obj.setChoice(MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoice.NONE)
}

// Hardware acceleration types.
// Choice returns a string
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) setChoice(value MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration {
	intValue, ok := otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.InlineCrypto = nil
	obj.inlineCryptoHolder = nil

	if value == MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoice.INLINE_CRYPTO {
		obj.obj.InlineCrypto = NewMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto().msg()
	}

	return obj
}

// description is TBD
// InlineCrypto returns a MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) InlineCrypto() MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {
	if obj.obj.InlineCrypto == nil {
		obj.setChoice(MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoice.INLINE_CRYPTO)
	}
	if obj.inlineCryptoHolder == nil {
		obj.inlineCryptoHolder = &macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto{obj: obj.obj.InlineCrypto}
	}
	return obj.inlineCryptoHolder
}

// description is TBD
// InlineCrypto returns a MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) HasInlineCrypto() bool {
	return obj.obj.InlineCrypto != nil
}

// description is TBD
// SetInlineCrypto sets the MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto value in the MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration object
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) SetInlineCrypto(value MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration {
	obj.setChoice(MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoice.INLINE_CRYPTO)
	obj.inlineCryptoHolder = nil
	obj.obj.InlineCrypto = value.msg()

	return obj
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.InlineCrypto != nil {

		obj.InlineCrypto().validateObj(vObj, set_default)
	}

}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) setDefault() {
	var choices_set int = 0
	var choice MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum

	if obj.obj.InlineCrypto != nil {
		choices_set += 1
		choice = MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoice.INLINE_CRYPTO
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoice.NONE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration")
			}
		} else {
			intVal := otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration_Choice_Enum_value[string(choice)]
			enumValue := otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

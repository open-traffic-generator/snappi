package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration *****
type macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration struct {
	validation
	obj                *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	marshaller         marshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	unMarshaller       unMarshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	inlineCryptoHolder MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
}

func NewMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration() MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration {
	obj := macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration{obj: &otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) msg() *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration {
	return obj.obj
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) setMsg(msg *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration struct {
	obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
}

type marshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration interface {
	// ToProto marshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration to protobuf object *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	ToProto() (*otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration, error)
	// ToPbText marshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration struct {
	obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
}

type unMarshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration interface {
	// FromProto unmarshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration from protobuf object *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	FromProto(msg *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) (MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration, error)
	// FromPbText unmarshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) Marshal() marshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) Unmarshal() unMarshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) ToProto() (*otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) FromProto(msg *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) (MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) FromJson(value string) error {
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

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) Clone() (MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration()
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

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) setNil() {
	obj.inlineCryptoHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration is hardware acceleration settings.
type MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration interface {
	Validation
	// msg marshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration to protobuf object *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	// setMsg unmarshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration from protobuf object *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	// validate validates MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum, set in MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	Choice() MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum
	// setChoice assigns MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum provided by user to MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	setChoice(value MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum) MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	// HasChoice checks if Choice has been set in MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	HasChoice() bool
	// getter for None to set choice.
	None()
	// InlineCrypto returns MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto, set in MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration.
	// MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto is inline cryto engine settings.
	InlineCrypto() MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// SetInlineCrypto assigns MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto provided by user to MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration.
	// MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto is inline cryto engine settings.
	SetInlineCrypto(value MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	// HasInlineCrypto checks if InlineCrypto has been set in MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	HasInlineCrypto() bool
	setNil()
}

type MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum string

// Enum of Choice on MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
var MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoice = struct {
	NONE          MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum
	INLINE_CRYPTO MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum
}{
	NONE:          MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum("none"),
	INLINE_CRYPTO: MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum("inline_crypto"),
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) Choice() MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum {
	return MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for None to set choice
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) None() {
	obj.setChoice(MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoice.NONE)
}

// Hardware acceleration types.
// Choice returns a string
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) setChoice(value MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum) MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration {
	intValue, ok := otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.InlineCrypto = nil
	obj.inlineCryptoHolder = nil

	if value == MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoice.INLINE_CRYPTO {
		obj.obj.InlineCrypto = NewMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto().msg()
	}

	return obj
}

// description is TBD
// InlineCrypto returns a MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) InlineCrypto() MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {
	if obj.obj.InlineCrypto == nil {
		obj.setChoice(MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoice.INLINE_CRYPTO)
	}
	if obj.inlineCryptoHolder == nil {
		obj.inlineCryptoHolder = &macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto{obj: obj.obj.InlineCrypto}
	}
	return obj.inlineCryptoHolder
}

// description is TBD
// InlineCrypto returns a MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) HasInlineCrypto() bool {
	return obj.obj.InlineCrypto != nil
}

// description is TBD
// SetInlineCrypto sets the MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto value in the MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration object
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) SetInlineCrypto(value MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration {
	obj.setChoice(MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoice.INLINE_CRYPTO)
	obj.inlineCryptoHolder = nil
	obj.obj.InlineCrypto = value.msg()

	return obj
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.InlineCrypto != nil {

		obj.InlineCrypto().validateObj(vObj, set_default)
	}

}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) setDefault() {
	var choices_set int = 0
	var choice MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum

	if obj.obj.InlineCrypto != nil {
		choices_set += 1
		choice = MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoice.INLINE_CRYPTO
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationChoice.NONE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration")
			}
		} else {
			intVal := otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration_Choice_Enum_value[string(choice)]
			enumValue := otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

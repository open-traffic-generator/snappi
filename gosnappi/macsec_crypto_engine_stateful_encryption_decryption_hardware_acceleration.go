package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration *****
type macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration struct {
	validation
	obj                *otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
	marshaller         marshalMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
	unMarshaller       unMarshalMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
	inlineCryptoHolder MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
}

func NewMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration() MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration {
	obj := macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration{obj: &otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) msg() *otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration {
	return obj.obj
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) setMsg(msg *otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration struct {
	obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
}

type marshalMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration interface {
	// ToProto marshals MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration to protobuf object *otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
	ToProto() (*otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration, error)
	// ToPbText marshals MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration struct {
	obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
}

type unMarshalMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration interface {
	// FromProto unmarshals MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration from protobuf object *otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
	FromProto(msg *otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) (MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration, error)
	// FromPbText unmarshals MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) Marshal() marshalMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) Unmarshal() unMarshalMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) ToProto() (*otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) FromProto(msg *otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) (MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) FromJson(value string) error {
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

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) Clone() (MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration()
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

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) setNil() {
	obj.inlineCryptoHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration is hardware acceleration settings.
type MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration interface {
	Validation
	// msg marshals MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration to protobuf object *otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
	// setMsg unmarshals MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration from protobuf object *otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
	// validate validates MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum, set in MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
	Choice() MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum
	// setChoice assigns MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum provided by user to MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
	setChoice(value MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum) MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
	// HasChoice checks if Choice has been set in MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
	HasChoice() bool
	// getter for None to set choice.
	None()
	// InlineCrypto returns MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto, set in MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration.
	// MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto is inline cryto engine settings.
	InlineCrypto() MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// SetInlineCrypto assigns MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto provided by user to MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration.
	// MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto is inline cryto engine settings.
	SetInlineCrypto(value MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
	// HasInlineCrypto checks if InlineCrypto has been set in MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
	HasInlineCrypto() bool
	setNil()
}

type MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum string

// Enum of Choice on MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
var MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationChoice = struct {
	NONE          MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum
	INLINE_CRYPTO MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum
}{
	NONE:          MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum("none"),
	INLINE_CRYPTO: MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum("inline_crypto"),
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) Choice() MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum {
	return MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for None to set choice
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) None() {
	obj.setChoice(MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationChoice.NONE)
}

// Hardware acceleration types.
// Choice returns a string
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) setChoice(value MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum) MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration {
	intValue, ok := otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.InlineCrypto = nil
	obj.inlineCryptoHolder = nil

	if value == MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationChoice.INLINE_CRYPTO {
		obj.obj.InlineCrypto = NewMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto().msg()
	}

	return obj
}

// description is TBD
// InlineCrypto returns a MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) InlineCrypto() MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {
	if obj.obj.InlineCrypto == nil {
		obj.setChoice(MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationChoice.INLINE_CRYPTO)
	}
	if obj.inlineCryptoHolder == nil {
		obj.inlineCryptoHolder = &macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto{obj: obj.obj.InlineCrypto}
	}
	return obj.inlineCryptoHolder
}

// description is TBD
// InlineCrypto returns a MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) HasInlineCrypto() bool {
	return obj.obj.InlineCrypto != nil
}

// description is TBD
// SetInlineCrypto sets the MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto value in the MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration object
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) SetInlineCrypto(value MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration {
	obj.setChoice(MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationChoice.INLINE_CRYPTO)
	obj.inlineCryptoHolder = nil
	obj.obj.InlineCrypto = value.msg()

	return obj
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.InlineCrypto != nil {

		obj.InlineCrypto().validateObj(vObj, set_default)
	}

}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) setDefault() {
	var choices_set int = 0
	var choice MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationChoiceEnum

	if obj.obj.InlineCrypto != nil {
		choices_set += 1
		choice = MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationChoice.INLINE_CRYPTO
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationChoice.NONE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration")
			}
		} else {
			intVal := otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration_Choice_Enum_value[string(choice)]
			enumValue := otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecHardwareAcceleration *****
type macsecHardwareAcceleration struct {
	validation
	obj                *otg.MacsecHardwareAcceleration
	marshaller         marshalMacsecHardwareAcceleration
	unMarshaller       unMarshalMacsecHardwareAcceleration
	inlineCryptoHolder MacsecHardwareAccelerationInlineCrypto
}

func NewMacsecHardwareAcceleration() MacsecHardwareAcceleration {
	obj := macsecHardwareAcceleration{obj: &otg.MacsecHardwareAcceleration{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecHardwareAcceleration) msg() *otg.MacsecHardwareAcceleration {
	return obj.obj
}

func (obj *macsecHardwareAcceleration) setMsg(msg *otg.MacsecHardwareAcceleration) MacsecHardwareAcceleration {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecHardwareAcceleration struct {
	obj *macsecHardwareAcceleration
}

type marshalMacsecHardwareAcceleration interface {
	// ToProto marshals MacsecHardwareAcceleration to protobuf object *otg.MacsecHardwareAcceleration
	ToProto() (*otg.MacsecHardwareAcceleration, error)
	// ToPbText marshals MacsecHardwareAcceleration to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecHardwareAcceleration to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecHardwareAcceleration to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecHardwareAcceleration struct {
	obj *macsecHardwareAcceleration
}

type unMarshalMacsecHardwareAcceleration interface {
	// FromProto unmarshals MacsecHardwareAcceleration from protobuf object *otg.MacsecHardwareAcceleration
	FromProto(msg *otg.MacsecHardwareAcceleration) (MacsecHardwareAcceleration, error)
	// FromPbText unmarshals MacsecHardwareAcceleration from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecHardwareAcceleration from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecHardwareAcceleration from JSON text
	FromJson(value string) error
}

func (obj *macsecHardwareAcceleration) Marshal() marshalMacsecHardwareAcceleration {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecHardwareAcceleration{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecHardwareAcceleration) Unmarshal() unMarshalMacsecHardwareAcceleration {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecHardwareAcceleration{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecHardwareAcceleration) ToProto() (*otg.MacsecHardwareAcceleration, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecHardwareAcceleration) FromProto(msg *otg.MacsecHardwareAcceleration) (MacsecHardwareAcceleration, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecHardwareAcceleration) ToPbText() (string, error) {
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

func (m *unMarshalmacsecHardwareAcceleration) FromPbText(value string) error {
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

func (m *marshalmacsecHardwareAcceleration) ToYaml() (string, error) {
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

func (m *unMarshalmacsecHardwareAcceleration) FromYaml(value string) error {
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

func (m *marshalmacsecHardwareAcceleration) ToJson() (string, error) {
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

func (m *unMarshalmacsecHardwareAcceleration) FromJson(value string) error {
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

func (obj *macsecHardwareAcceleration) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecHardwareAcceleration) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecHardwareAcceleration) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecHardwareAcceleration) Clone() (MacsecHardwareAcceleration, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecHardwareAcceleration()
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

func (obj *macsecHardwareAcceleration) setNil() {
	obj.inlineCryptoHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecHardwareAcceleration is hardware acceleration configuration for offloading MACsec processing to hardware.
type MacsecHardwareAcceleration interface {
	Validation
	// msg marshals MacsecHardwareAcceleration to protobuf object *otg.MacsecHardwareAcceleration
	// and doesn't set defaults
	msg() *otg.MacsecHardwareAcceleration
	// setMsg unmarshals MacsecHardwareAcceleration from protobuf object *otg.MacsecHardwareAcceleration
	// and doesn't set defaults
	setMsg(*otg.MacsecHardwareAcceleration) MacsecHardwareAcceleration
	// provides marshal interface
	Marshal() marshalMacsecHardwareAcceleration
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecHardwareAcceleration
	// validate validates MacsecHardwareAcceleration
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecHardwareAcceleration, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MacsecHardwareAccelerationChoiceEnum, set in MacsecHardwareAcceleration
	Choice() MacsecHardwareAccelerationChoiceEnum
	// setChoice assigns MacsecHardwareAccelerationChoiceEnum provided by user to MacsecHardwareAcceleration
	setChoice(value MacsecHardwareAccelerationChoiceEnum) MacsecHardwareAcceleration
	// HasChoice checks if Choice has been set in MacsecHardwareAcceleration
	HasChoice() bool
	// getter for None to set choice.
	None()
	// InlineCrypto returns MacsecHardwareAccelerationInlineCrypto, set in MacsecHardwareAcceleration.
	// MacsecHardwareAccelerationInlineCrypto is inline cryto engine configuration. Encryption/ decryption are offloaded to hardware. Also dynamic fields e.g. packet number(PN) and integrity check value(ICV) are updated in MACsec header on transmit.
	InlineCrypto() MacsecHardwareAccelerationInlineCrypto
	// SetInlineCrypto assigns MacsecHardwareAccelerationInlineCrypto provided by user to MacsecHardwareAcceleration.
	// MacsecHardwareAccelerationInlineCrypto is inline cryto engine configuration. Encryption/ decryption are offloaded to hardware. Also dynamic fields e.g. packet number(PN) and integrity check value(ICV) are updated in MACsec header on transmit.
	SetInlineCrypto(value MacsecHardwareAccelerationInlineCrypto) MacsecHardwareAcceleration
	// HasInlineCrypto checks if InlineCrypto has been set in MacsecHardwareAcceleration
	HasInlineCrypto() bool
	setNil()
}

type MacsecHardwareAccelerationChoiceEnum string

// Enum of Choice on MacsecHardwareAcceleration
var MacsecHardwareAccelerationChoice = struct {
	NONE          MacsecHardwareAccelerationChoiceEnum
	INLINE_CRYPTO MacsecHardwareAccelerationChoiceEnum
}{
	NONE:          MacsecHardwareAccelerationChoiceEnum("none"),
	INLINE_CRYPTO: MacsecHardwareAccelerationChoiceEnum("inline_crypto"),
}

func (obj *macsecHardwareAcceleration) Choice() MacsecHardwareAccelerationChoiceEnum {
	return MacsecHardwareAccelerationChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for None to set choice
func (obj *macsecHardwareAcceleration) None() {
	obj.setChoice(MacsecHardwareAccelerationChoice.NONE)
}

// Hardware acceleration types.
// Choice returns a string
func (obj *macsecHardwareAcceleration) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *macsecHardwareAcceleration) setChoice(value MacsecHardwareAccelerationChoiceEnum) MacsecHardwareAcceleration {
	intValue, ok := otg.MacsecHardwareAcceleration_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecHardwareAccelerationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecHardwareAcceleration_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.InlineCrypto = nil
	obj.inlineCryptoHolder = nil

	if value == MacsecHardwareAccelerationChoice.INLINE_CRYPTO {
		obj.obj.InlineCrypto = NewMacsecHardwareAccelerationInlineCrypto().msg()
	}

	return obj
}

// description is TBD
// InlineCrypto returns a MacsecHardwareAccelerationInlineCrypto
func (obj *macsecHardwareAcceleration) InlineCrypto() MacsecHardwareAccelerationInlineCrypto {
	if obj.obj.InlineCrypto == nil {
		obj.setChoice(MacsecHardwareAccelerationChoice.INLINE_CRYPTO)
	}
	if obj.inlineCryptoHolder == nil {
		obj.inlineCryptoHolder = &macsecHardwareAccelerationInlineCrypto{obj: obj.obj.InlineCrypto}
	}
	return obj.inlineCryptoHolder
}

// description is TBD
// InlineCrypto returns a MacsecHardwareAccelerationInlineCrypto
func (obj *macsecHardwareAcceleration) HasInlineCrypto() bool {
	return obj.obj.InlineCrypto != nil
}

// description is TBD
// SetInlineCrypto sets the MacsecHardwareAccelerationInlineCrypto value in the MacsecHardwareAcceleration object
func (obj *macsecHardwareAcceleration) SetInlineCrypto(value MacsecHardwareAccelerationInlineCrypto) MacsecHardwareAcceleration {
	obj.setChoice(MacsecHardwareAccelerationChoice.INLINE_CRYPTO)
	obj.inlineCryptoHolder = nil
	obj.obj.InlineCrypto = value.msg()

	return obj
}

func (obj *macsecHardwareAcceleration) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.InlineCrypto != nil {

		obj.InlineCrypto().validateObj(vObj, set_default)
	}

}

func (obj *macsecHardwareAcceleration) setDefault() {
	var choices_set int = 0
	var choice MacsecHardwareAccelerationChoiceEnum

	if obj.obj.InlineCrypto != nil {
		choices_set += 1
		choice = MacsecHardwareAccelerationChoice.INLINE_CRYPTO
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MacsecHardwareAccelerationChoice.NONE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MacsecHardwareAcceleration")
			}
		} else {
			intVal := otg.MacsecHardwareAcceleration_Choice_Enum_value[string(choice)]
			enumValue := otg.MacsecHardwareAcceleration_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

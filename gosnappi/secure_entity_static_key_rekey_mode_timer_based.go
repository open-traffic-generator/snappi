package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityStaticKeyRekeyModeTimerBased *****
type secureEntityStaticKeyRekeyModeTimerBased struct {
	validation
	obj          *otg.SecureEntityStaticKeyRekeyModeTimerBased
	marshaller   marshalSecureEntityStaticKeyRekeyModeTimerBased
	unMarshaller unMarshalSecureEntityStaticKeyRekeyModeTimerBased
}

func NewSecureEntityStaticKeyRekeyModeTimerBased() SecureEntityStaticKeyRekeyModeTimerBased {
	obj := secureEntityStaticKeyRekeyModeTimerBased{obj: &otg.SecureEntityStaticKeyRekeyModeTimerBased{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityStaticKeyRekeyModeTimerBased) msg() *otg.SecureEntityStaticKeyRekeyModeTimerBased {
	return obj.obj
}

func (obj *secureEntityStaticKeyRekeyModeTimerBased) setMsg(msg *otg.SecureEntityStaticKeyRekeyModeTimerBased) SecureEntityStaticKeyRekeyModeTimerBased {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityStaticKeyRekeyModeTimerBased struct {
	obj *secureEntityStaticKeyRekeyModeTimerBased
}

type marshalSecureEntityStaticKeyRekeyModeTimerBased interface {
	// ToProto marshals SecureEntityStaticKeyRekeyModeTimerBased to protobuf object *otg.SecureEntityStaticKeyRekeyModeTimerBased
	ToProto() (*otg.SecureEntityStaticKeyRekeyModeTimerBased, error)
	// ToPbText marshals SecureEntityStaticKeyRekeyModeTimerBased to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityStaticKeyRekeyModeTimerBased to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityStaticKeyRekeyModeTimerBased to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals SecureEntityStaticKeyRekeyModeTimerBased to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalsecureEntityStaticKeyRekeyModeTimerBased struct {
	obj *secureEntityStaticKeyRekeyModeTimerBased
}

type unMarshalSecureEntityStaticKeyRekeyModeTimerBased interface {
	// FromProto unmarshals SecureEntityStaticKeyRekeyModeTimerBased from protobuf object *otg.SecureEntityStaticKeyRekeyModeTimerBased
	FromProto(msg *otg.SecureEntityStaticKeyRekeyModeTimerBased) (SecureEntityStaticKeyRekeyModeTimerBased, error)
	// FromPbText unmarshals SecureEntityStaticKeyRekeyModeTimerBased from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityStaticKeyRekeyModeTimerBased from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityStaticKeyRekeyModeTimerBased from JSON text
	FromJson(value string) error
}

func (obj *secureEntityStaticKeyRekeyModeTimerBased) Marshal() marshalSecureEntityStaticKeyRekeyModeTimerBased {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityStaticKeyRekeyModeTimerBased{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityStaticKeyRekeyModeTimerBased) Unmarshal() unMarshalSecureEntityStaticKeyRekeyModeTimerBased {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityStaticKeyRekeyModeTimerBased{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityStaticKeyRekeyModeTimerBased) ToProto() (*otg.SecureEntityStaticKeyRekeyModeTimerBased, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityStaticKeyRekeyModeTimerBased) FromProto(msg *otg.SecureEntityStaticKeyRekeyModeTimerBased) (SecureEntityStaticKeyRekeyModeTimerBased, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityStaticKeyRekeyModeTimerBased) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityStaticKeyRekeyModeTimerBased) FromPbText(value string) error {
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

func (m *marshalsecureEntityStaticKeyRekeyModeTimerBased) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityStaticKeyRekeyModeTimerBased) FromYaml(value string) error {
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

func (m *marshalsecureEntityStaticKeyRekeyModeTimerBased) ToJsonRaw() (string, error) {
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

func (m *marshalsecureEntityStaticKeyRekeyModeTimerBased) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityStaticKeyRekeyModeTimerBased) FromJson(value string) error {
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

func (obj *secureEntityStaticKeyRekeyModeTimerBased) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityStaticKeyRekeyModeTimerBased) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityStaticKeyRekeyModeTimerBased) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityStaticKeyRekeyModeTimerBased) Clone() (SecureEntityStaticKeyRekeyModeTimerBased, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityStaticKeyRekeyModeTimerBased()
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

// SecureEntityStaticKeyRekeyModeTimerBased is timer based periodic rekey properties.
type SecureEntityStaticKeyRekeyModeTimerBased interface {
	Validation
	// msg marshals SecureEntityStaticKeyRekeyModeTimerBased to protobuf object *otg.SecureEntityStaticKeyRekeyModeTimerBased
	// and doesn't set defaults
	msg() *otg.SecureEntityStaticKeyRekeyModeTimerBased
	// setMsg unmarshals SecureEntityStaticKeyRekeyModeTimerBased from protobuf object *otg.SecureEntityStaticKeyRekeyModeTimerBased
	// and doesn't set defaults
	setMsg(*otg.SecureEntityStaticKeyRekeyModeTimerBased) SecureEntityStaticKeyRekeyModeTimerBased
	// provides marshal interface
	Marshal() marshalSecureEntityStaticKeyRekeyModeTimerBased
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityStaticKeyRekeyModeTimerBased
	// validate validates SecureEntityStaticKeyRekeyModeTimerBased
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityStaticKeyRekeyModeTimerBased, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns SecureEntityStaticKeyRekeyModeTimerBasedChoiceEnum, set in SecureEntityStaticKeyRekeyModeTimerBased
	Choice() SecureEntityStaticKeyRekeyModeTimerBasedChoiceEnum
	// setChoice assigns SecureEntityStaticKeyRekeyModeTimerBasedChoiceEnum provided by user to SecureEntityStaticKeyRekeyModeTimerBased
	setChoice(value SecureEntityStaticKeyRekeyModeTimerBasedChoiceEnum) SecureEntityStaticKeyRekeyModeTimerBased
	// HasChoice checks if Choice has been set in SecureEntityStaticKeyRekeyModeTimerBased
	HasChoice() bool
	// getter for Continuous to set choice.
	Continuous()
	// FixedCount returns uint32, set in SecureEntityStaticKeyRekeyModeTimerBased.
	FixedCount() uint32
	// SetFixedCount assigns uint32 provided by user to SecureEntityStaticKeyRekeyModeTimerBased
	SetFixedCount(value uint32) SecureEntityStaticKeyRekeyModeTimerBased
	// HasFixedCount checks if FixedCount has been set in SecureEntityStaticKeyRekeyModeTimerBased
	HasFixedCount() bool
	// Interval returns uint32, set in SecureEntityStaticKeyRekeyModeTimerBased.
	Interval() uint32
	// SetInterval assigns uint32 provided by user to SecureEntityStaticKeyRekeyModeTimerBased
	SetInterval(value uint32) SecureEntityStaticKeyRekeyModeTimerBased
	// HasInterval checks if Interval has been set in SecureEntityStaticKeyRekeyModeTimerBased
	HasInterval() bool
}

type SecureEntityStaticKeyRekeyModeTimerBasedChoiceEnum string

// Enum of Choice on SecureEntityStaticKeyRekeyModeTimerBased
var SecureEntityStaticKeyRekeyModeTimerBasedChoice = struct {
	CONTINUOUS  SecureEntityStaticKeyRekeyModeTimerBasedChoiceEnum
	FIXED_COUNT SecureEntityStaticKeyRekeyModeTimerBasedChoiceEnum
}{
	CONTINUOUS:  SecureEntityStaticKeyRekeyModeTimerBasedChoiceEnum("continuous"),
	FIXED_COUNT: SecureEntityStaticKeyRekeyModeTimerBasedChoiceEnum("fixed_count"),
}

func (obj *secureEntityStaticKeyRekeyModeTimerBased) Choice() SecureEntityStaticKeyRekeyModeTimerBasedChoiceEnum {
	return SecureEntityStaticKeyRekeyModeTimerBasedChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Continuous to set choice
func (obj *secureEntityStaticKeyRekeyModeTimerBased) Continuous() {
	obj.setChoice(SecureEntityStaticKeyRekeyModeTimerBasedChoice.CONTINUOUS)
}

// Periodic rekey attempt choices.
// Choice returns a string
func (obj *secureEntityStaticKeyRekeyModeTimerBased) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *secureEntityStaticKeyRekeyModeTimerBased) setChoice(value SecureEntityStaticKeyRekeyModeTimerBasedChoiceEnum) SecureEntityStaticKeyRekeyModeTimerBased {
	intValue, ok := otg.SecureEntityStaticKeyRekeyModeTimerBased_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on SecureEntityStaticKeyRekeyModeTimerBasedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.SecureEntityStaticKeyRekeyModeTimerBased_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.FixedCount = nil

	if value == SecureEntityStaticKeyRekeyModeTimerBasedChoice.FIXED_COUNT {
		defaultValue := uint32(10)
		obj.obj.FixedCount = &defaultValue
	}

	return obj
}

// Fixed rekey attempts.
// FixedCount returns a uint32
func (obj *secureEntityStaticKeyRekeyModeTimerBased) FixedCount() uint32 {

	if obj.obj.FixedCount == nil {
		obj.setChoice(SecureEntityStaticKeyRekeyModeTimerBasedChoice.FIXED_COUNT)
	}

	return *obj.obj.FixedCount

}

// Fixed rekey attempts.
// FixedCount returns a uint32
func (obj *secureEntityStaticKeyRekeyModeTimerBased) HasFixedCount() bool {
	return obj.obj.FixedCount != nil
}

// Fixed rekey attempts.
// SetFixedCount sets the uint32 value in the SecureEntityStaticKeyRekeyModeTimerBased object
func (obj *secureEntityStaticKeyRekeyModeTimerBased) SetFixedCount(value uint32) SecureEntityStaticKeyRekeyModeTimerBased {
	obj.setChoice(SecureEntityStaticKeyRekeyModeTimerBasedChoice.FIXED_COUNT)
	obj.obj.FixedCount = &value
	return obj
}

// Periodic rekey interval (sec).
// Interval returns a uint32
func (obj *secureEntityStaticKeyRekeyModeTimerBased) Interval() uint32 {

	return *obj.obj.Interval

}

// Periodic rekey interval (sec).
// Interval returns a uint32
func (obj *secureEntityStaticKeyRekeyModeTimerBased) HasInterval() bool {
	return obj.obj.Interval != nil
}

// Periodic rekey interval (sec).
// SetInterval sets the uint32 value in the SecureEntityStaticKeyRekeyModeTimerBased object
func (obj *secureEntityStaticKeyRekeyModeTimerBased) SetInterval(value uint32) SecureEntityStaticKeyRekeyModeTimerBased {

	obj.obj.Interval = &value
	return obj
}

func (obj *secureEntityStaticKeyRekeyModeTimerBased) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.FixedCount != nil {

		if *obj.obj.FixedCount < 1 || *obj.obj.FixedCount > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= SecureEntityStaticKeyRekeyModeTimerBased.FixedCount <= 65535 but Got %d", *obj.obj.FixedCount))
		}

	}

	if obj.obj.Interval != nil {

		if *obj.obj.Interval < 30 || *obj.obj.Interval > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("30 <= SecureEntityStaticKeyRekeyModeTimerBased.Interval <= 65535 but Got %d", *obj.obj.Interval))
		}

	}

}

func (obj *secureEntityStaticKeyRekeyModeTimerBased) setDefault() {
	var choices_set int = 0
	var choice SecureEntityStaticKeyRekeyModeTimerBasedChoiceEnum

	if obj.obj.FixedCount != nil {
		choices_set += 1
		choice = SecureEntityStaticKeyRekeyModeTimerBasedChoice.FIXED_COUNT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(SecureEntityStaticKeyRekeyModeTimerBasedChoice.CONTINUOUS)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in SecureEntityStaticKeyRekeyModeTimerBased")
			}
		} else {
			intVal := otg.SecureEntityStaticKeyRekeyModeTimerBased_Choice_Enum_value[string(choice)]
			enumValue := otg.SecureEntityStaticKeyRekeyModeTimerBased_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.Interval == nil {
		obj.SetInterval(300)
	}

}

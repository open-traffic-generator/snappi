package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecAdvanceStaticKeyRekeyModeTimerBased *****
type macsecAdvanceStaticKeyRekeyModeTimerBased struct {
	validation
	obj          *otg.MacsecAdvanceStaticKeyRekeyModeTimerBased
	marshaller   marshalMacsecAdvanceStaticKeyRekeyModeTimerBased
	unMarshaller unMarshalMacsecAdvanceStaticKeyRekeyModeTimerBased
}

func NewMacsecAdvanceStaticKeyRekeyModeTimerBased() MacsecAdvanceStaticKeyRekeyModeTimerBased {
	obj := macsecAdvanceStaticKeyRekeyModeTimerBased{obj: &otg.MacsecAdvanceStaticKeyRekeyModeTimerBased{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecAdvanceStaticKeyRekeyModeTimerBased) msg() *otg.MacsecAdvanceStaticKeyRekeyModeTimerBased {
	return obj.obj
}

func (obj *macsecAdvanceStaticKeyRekeyModeTimerBased) setMsg(msg *otg.MacsecAdvanceStaticKeyRekeyModeTimerBased) MacsecAdvanceStaticKeyRekeyModeTimerBased {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecAdvanceStaticKeyRekeyModeTimerBased struct {
	obj *macsecAdvanceStaticKeyRekeyModeTimerBased
}

type marshalMacsecAdvanceStaticKeyRekeyModeTimerBased interface {
	// ToProto marshals MacsecAdvanceStaticKeyRekeyModeTimerBased to protobuf object *otg.MacsecAdvanceStaticKeyRekeyModeTimerBased
	ToProto() (*otg.MacsecAdvanceStaticKeyRekeyModeTimerBased, error)
	// ToPbText marshals MacsecAdvanceStaticKeyRekeyModeTimerBased to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecAdvanceStaticKeyRekeyModeTimerBased to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecAdvanceStaticKeyRekeyModeTimerBased to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecAdvanceStaticKeyRekeyModeTimerBased struct {
	obj *macsecAdvanceStaticKeyRekeyModeTimerBased
}

type unMarshalMacsecAdvanceStaticKeyRekeyModeTimerBased interface {
	// FromProto unmarshals MacsecAdvanceStaticKeyRekeyModeTimerBased from protobuf object *otg.MacsecAdvanceStaticKeyRekeyModeTimerBased
	FromProto(msg *otg.MacsecAdvanceStaticKeyRekeyModeTimerBased) (MacsecAdvanceStaticKeyRekeyModeTimerBased, error)
	// FromPbText unmarshals MacsecAdvanceStaticKeyRekeyModeTimerBased from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecAdvanceStaticKeyRekeyModeTimerBased from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecAdvanceStaticKeyRekeyModeTimerBased from JSON text
	FromJson(value string) error
}

func (obj *macsecAdvanceStaticKeyRekeyModeTimerBased) Marshal() marshalMacsecAdvanceStaticKeyRekeyModeTimerBased {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecAdvanceStaticKeyRekeyModeTimerBased{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecAdvanceStaticKeyRekeyModeTimerBased) Unmarshal() unMarshalMacsecAdvanceStaticKeyRekeyModeTimerBased {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecAdvanceStaticKeyRekeyModeTimerBased{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecAdvanceStaticKeyRekeyModeTimerBased) ToProto() (*otg.MacsecAdvanceStaticKeyRekeyModeTimerBased, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecAdvanceStaticKeyRekeyModeTimerBased) FromProto(msg *otg.MacsecAdvanceStaticKeyRekeyModeTimerBased) (MacsecAdvanceStaticKeyRekeyModeTimerBased, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecAdvanceStaticKeyRekeyModeTimerBased) ToPbText() (string, error) {
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

func (m *unMarshalmacsecAdvanceStaticKeyRekeyModeTimerBased) FromPbText(value string) error {
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

func (m *marshalmacsecAdvanceStaticKeyRekeyModeTimerBased) ToYaml() (string, error) {
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

func (m *unMarshalmacsecAdvanceStaticKeyRekeyModeTimerBased) FromYaml(value string) error {
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

func (m *marshalmacsecAdvanceStaticKeyRekeyModeTimerBased) ToJson() (string, error) {
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

func (m *unMarshalmacsecAdvanceStaticKeyRekeyModeTimerBased) FromJson(value string) error {
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

func (obj *macsecAdvanceStaticKeyRekeyModeTimerBased) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecAdvanceStaticKeyRekeyModeTimerBased) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecAdvanceStaticKeyRekeyModeTimerBased) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecAdvanceStaticKeyRekeyModeTimerBased) Clone() (MacsecAdvanceStaticKeyRekeyModeTimerBased, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecAdvanceStaticKeyRekeyModeTimerBased()
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

// MacsecAdvanceStaticKeyRekeyModeTimerBased is timer based periodic rekey properties.
type MacsecAdvanceStaticKeyRekeyModeTimerBased interface {
	Validation
	// msg marshals MacsecAdvanceStaticKeyRekeyModeTimerBased to protobuf object *otg.MacsecAdvanceStaticKeyRekeyModeTimerBased
	// and doesn't set defaults
	msg() *otg.MacsecAdvanceStaticKeyRekeyModeTimerBased
	// setMsg unmarshals MacsecAdvanceStaticKeyRekeyModeTimerBased from protobuf object *otg.MacsecAdvanceStaticKeyRekeyModeTimerBased
	// and doesn't set defaults
	setMsg(*otg.MacsecAdvanceStaticKeyRekeyModeTimerBased) MacsecAdvanceStaticKeyRekeyModeTimerBased
	// provides marshal interface
	Marshal() marshalMacsecAdvanceStaticKeyRekeyModeTimerBased
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecAdvanceStaticKeyRekeyModeTimerBased
	// validate validates MacsecAdvanceStaticKeyRekeyModeTimerBased
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecAdvanceStaticKeyRekeyModeTimerBased, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MacsecAdvanceStaticKeyRekeyModeTimerBasedChoiceEnum, set in MacsecAdvanceStaticKeyRekeyModeTimerBased
	Choice() MacsecAdvanceStaticKeyRekeyModeTimerBasedChoiceEnum
	// setChoice assigns MacsecAdvanceStaticKeyRekeyModeTimerBasedChoiceEnum provided by user to MacsecAdvanceStaticKeyRekeyModeTimerBased
	setChoice(value MacsecAdvanceStaticKeyRekeyModeTimerBasedChoiceEnum) MacsecAdvanceStaticKeyRekeyModeTimerBased
	// HasChoice checks if Choice has been set in MacsecAdvanceStaticKeyRekeyModeTimerBased
	HasChoice() bool
	// getter for Continuous to set choice.
	Continuous()
	// FixedCount returns uint32, set in MacsecAdvanceStaticKeyRekeyModeTimerBased.
	FixedCount() uint32
	// SetFixedCount assigns uint32 provided by user to MacsecAdvanceStaticKeyRekeyModeTimerBased
	SetFixedCount(value uint32) MacsecAdvanceStaticKeyRekeyModeTimerBased
	// HasFixedCount checks if FixedCount has been set in MacsecAdvanceStaticKeyRekeyModeTimerBased
	HasFixedCount() bool
	// Interval returns uint32, set in MacsecAdvanceStaticKeyRekeyModeTimerBased.
	Interval() uint32
	// SetInterval assigns uint32 provided by user to MacsecAdvanceStaticKeyRekeyModeTimerBased
	SetInterval(value uint32) MacsecAdvanceStaticKeyRekeyModeTimerBased
	// HasInterval checks if Interval has been set in MacsecAdvanceStaticKeyRekeyModeTimerBased
	HasInterval() bool
}

type MacsecAdvanceStaticKeyRekeyModeTimerBasedChoiceEnum string

// Enum of Choice on MacsecAdvanceStaticKeyRekeyModeTimerBased
var MacsecAdvanceStaticKeyRekeyModeTimerBasedChoice = struct {
	CONTINUOUS  MacsecAdvanceStaticKeyRekeyModeTimerBasedChoiceEnum
	FIXED_COUNT MacsecAdvanceStaticKeyRekeyModeTimerBasedChoiceEnum
}{
	CONTINUOUS:  MacsecAdvanceStaticKeyRekeyModeTimerBasedChoiceEnum("continuous"),
	FIXED_COUNT: MacsecAdvanceStaticKeyRekeyModeTimerBasedChoiceEnum("fixed_count"),
}

func (obj *macsecAdvanceStaticKeyRekeyModeTimerBased) Choice() MacsecAdvanceStaticKeyRekeyModeTimerBasedChoiceEnum {
	return MacsecAdvanceStaticKeyRekeyModeTimerBasedChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Continuous to set choice
func (obj *macsecAdvanceStaticKeyRekeyModeTimerBased) Continuous() {
	obj.setChoice(MacsecAdvanceStaticKeyRekeyModeTimerBasedChoice.CONTINUOUS)
}

// Periodic rekey attempt choices.
// Choice returns a string
func (obj *macsecAdvanceStaticKeyRekeyModeTimerBased) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *macsecAdvanceStaticKeyRekeyModeTimerBased) setChoice(value MacsecAdvanceStaticKeyRekeyModeTimerBasedChoiceEnum) MacsecAdvanceStaticKeyRekeyModeTimerBased {
	intValue, ok := otg.MacsecAdvanceStaticKeyRekeyModeTimerBased_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecAdvanceStaticKeyRekeyModeTimerBasedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecAdvanceStaticKeyRekeyModeTimerBased_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.FixedCount = nil

	if value == MacsecAdvanceStaticKeyRekeyModeTimerBasedChoice.FIXED_COUNT {
		defaultValue := uint32(10)
		obj.obj.FixedCount = &defaultValue
	}

	return obj
}

// Fixed rekey attempts.
// FixedCount returns a uint32
func (obj *macsecAdvanceStaticKeyRekeyModeTimerBased) FixedCount() uint32 {

	if obj.obj.FixedCount == nil {
		obj.setChoice(MacsecAdvanceStaticKeyRekeyModeTimerBasedChoice.FIXED_COUNT)
	}

	return *obj.obj.FixedCount

}

// Fixed rekey attempts.
// FixedCount returns a uint32
func (obj *macsecAdvanceStaticKeyRekeyModeTimerBased) HasFixedCount() bool {
	return obj.obj.FixedCount != nil
}

// Fixed rekey attempts.
// SetFixedCount sets the uint32 value in the MacsecAdvanceStaticKeyRekeyModeTimerBased object
func (obj *macsecAdvanceStaticKeyRekeyModeTimerBased) SetFixedCount(value uint32) MacsecAdvanceStaticKeyRekeyModeTimerBased {
	obj.setChoice(MacsecAdvanceStaticKeyRekeyModeTimerBasedChoice.FIXED_COUNT)
	obj.obj.FixedCount = &value
	return obj
}

// Periodic rekey interval (sec).
// Interval returns a uint32
func (obj *macsecAdvanceStaticKeyRekeyModeTimerBased) Interval() uint32 {

	return *obj.obj.Interval

}

// Periodic rekey interval (sec).
// Interval returns a uint32
func (obj *macsecAdvanceStaticKeyRekeyModeTimerBased) HasInterval() bool {
	return obj.obj.Interval != nil
}

// Periodic rekey interval (sec).
// SetInterval sets the uint32 value in the MacsecAdvanceStaticKeyRekeyModeTimerBased object
func (obj *macsecAdvanceStaticKeyRekeyModeTimerBased) SetInterval(value uint32) MacsecAdvanceStaticKeyRekeyModeTimerBased {

	obj.obj.Interval = &value
	return obj
}

func (obj *macsecAdvanceStaticKeyRekeyModeTimerBased) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.FixedCount != nil {

		if *obj.obj.FixedCount < 1 || *obj.obj.FixedCount > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MacsecAdvanceStaticKeyRekeyModeTimerBased.FixedCount <= 65535 but Got %d", *obj.obj.FixedCount))
		}

	}

	if obj.obj.Interval != nil {

		if *obj.obj.Interval < 30 || *obj.obj.Interval > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("30 <= MacsecAdvanceStaticKeyRekeyModeTimerBased.Interval <= 65535 but Got %d", *obj.obj.Interval))
		}

	}

}

func (obj *macsecAdvanceStaticKeyRekeyModeTimerBased) setDefault() {
	var choices_set int = 0
	var choice MacsecAdvanceStaticKeyRekeyModeTimerBasedChoiceEnum

	if obj.obj.FixedCount != nil {
		choices_set += 1
		choice = MacsecAdvanceStaticKeyRekeyModeTimerBasedChoice.FIXED_COUNT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MacsecAdvanceStaticKeyRekeyModeTimerBasedChoice.CONTINUOUS)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MacsecAdvanceStaticKeyRekeyModeTimerBased")
			}
		} else {
			intVal := otg.MacsecAdvanceStaticKeyRekeyModeTimerBased_Choice_Enum_value[string(choice)]
			enumValue := otg.MacsecAdvanceStaticKeyRekeyModeTimerBased_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.Interval == nil {
		obj.SetInterval(300)
	}

}

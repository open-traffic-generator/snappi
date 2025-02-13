package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecStaticKeyRekeyModeTimerBased *****
type macsecStaticKeyRekeyModeTimerBased struct {
	validation
	obj          *otg.MacsecStaticKeyRekeyModeTimerBased
	marshaller   marshalMacsecStaticKeyRekeyModeTimerBased
	unMarshaller unMarshalMacsecStaticKeyRekeyModeTimerBased
}

func NewMacsecStaticKeyRekeyModeTimerBased() MacsecStaticKeyRekeyModeTimerBased {
	obj := macsecStaticKeyRekeyModeTimerBased{obj: &otg.MacsecStaticKeyRekeyModeTimerBased{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecStaticKeyRekeyModeTimerBased) msg() *otg.MacsecStaticKeyRekeyModeTimerBased {
	return obj.obj
}

func (obj *macsecStaticKeyRekeyModeTimerBased) setMsg(msg *otg.MacsecStaticKeyRekeyModeTimerBased) MacsecStaticKeyRekeyModeTimerBased {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecStaticKeyRekeyModeTimerBased struct {
	obj *macsecStaticKeyRekeyModeTimerBased
}

type marshalMacsecStaticKeyRekeyModeTimerBased interface {
	// ToProto marshals MacsecStaticKeyRekeyModeTimerBased to protobuf object *otg.MacsecStaticKeyRekeyModeTimerBased
	ToProto() (*otg.MacsecStaticKeyRekeyModeTimerBased, error)
	// ToPbText marshals MacsecStaticKeyRekeyModeTimerBased to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecStaticKeyRekeyModeTimerBased to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecStaticKeyRekeyModeTimerBased to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecStaticKeyRekeyModeTimerBased struct {
	obj *macsecStaticKeyRekeyModeTimerBased
}

type unMarshalMacsecStaticKeyRekeyModeTimerBased interface {
	// FromProto unmarshals MacsecStaticKeyRekeyModeTimerBased from protobuf object *otg.MacsecStaticKeyRekeyModeTimerBased
	FromProto(msg *otg.MacsecStaticKeyRekeyModeTimerBased) (MacsecStaticKeyRekeyModeTimerBased, error)
	// FromPbText unmarshals MacsecStaticKeyRekeyModeTimerBased from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecStaticKeyRekeyModeTimerBased from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecStaticKeyRekeyModeTimerBased from JSON text
	FromJson(value string) error
}

func (obj *macsecStaticKeyRekeyModeTimerBased) Marshal() marshalMacsecStaticKeyRekeyModeTimerBased {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecStaticKeyRekeyModeTimerBased{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecStaticKeyRekeyModeTimerBased) Unmarshal() unMarshalMacsecStaticKeyRekeyModeTimerBased {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecStaticKeyRekeyModeTimerBased{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecStaticKeyRekeyModeTimerBased) ToProto() (*otg.MacsecStaticKeyRekeyModeTimerBased, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecStaticKeyRekeyModeTimerBased) FromProto(msg *otg.MacsecStaticKeyRekeyModeTimerBased) (MacsecStaticKeyRekeyModeTimerBased, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecStaticKeyRekeyModeTimerBased) ToPbText() (string, error) {
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

func (m *unMarshalmacsecStaticKeyRekeyModeTimerBased) FromPbText(value string) error {
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

func (m *marshalmacsecStaticKeyRekeyModeTimerBased) ToYaml() (string, error) {
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

func (m *unMarshalmacsecStaticKeyRekeyModeTimerBased) FromYaml(value string) error {
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

func (m *marshalmacsecStaticKeyRekeyModeTimerBased) ToJson() (string, error) {
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

func (m *unMarshalmacsecStaticKeyRekeyModeTimerBased) FromJson(value string) error {
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

func (obj *macsecStaticKeyRekeyModeTimerBased) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecStaticKeyRekeyModeTimerBased) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecStaticKeyRekeyModeTimerBased) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecStaticKeyRekeyModeTimerBased) Clone() (MacsecStaticKeyRekeyModeTimerBased, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecStaticKeyRekeyModeTimerBased()
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

// MacsecStaticKeyRekeyModeTimerBased is timer based periodic rekey properties.
type MacsecStaticKeyRekeyModeTimerBased interface {
	Validation
	// msg marshals MacsecStaticKeyRekeyModeTimerBased to protobuf object *otg.MacsecStaticKeyRekeyModeTimerBased
	// and doesn't set defaults
	msg() *otg.MacsecStaticKeyRekeyModeTimerBased
	// setMsg unmarshals MacsecStaticKeyRekeyModeTimerBased from protobuf object *otg.MacsecStaticKeyRekeyModeTimerBased
	// and doesn't set defaults
	setMsg(*otg.MacsecStaticKeyRekeyModeTimerBased) MacsecStaticKeyRekeyModeTimerBased
	// provides marshal interface
	Marshal() marshalMacsecStaticKeyRekeyModeTimerBased
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecStaticKeyRekeyModeTimerBased
	// validate validates MacsecStaticKeyRekeyModeTimerBased
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecStaticKeyRekeyModeTimerBased, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MacsecStaticKeyRekeyModeTimerBasedChoiceEnum, set in MacsecStaticKeyRekeyModeTimerBased
	Choice() MacsecStaticKeyRekeyModeTimerBasedChoiceEnum
	// setChoice assigns MacsecStaticKeyRekeyModeTimerBasedChoiceEnum provided by user to MacsecStaticKeyRekeyModeTimerBased
	setChoice(value MacsecStaticKeyRekeyModeTimerBasedChoiceEnum) MacsecStaticKeyRekeyModeTimerBased
	// HasChoice checks if Choice has been set in MacsecStaticKeyRekeyModeTimerBased
	HasChoice() bool
	// getter for Continuous to set choice.
	Continuous()
	// FixedCount returns uint32, set in MacsecStaticKeyRekeyModeTimerBased.
	FixedCount() uint32
	// SetFixedCount assigns uint32 provided by user to MacsecStaticKeyRekeyModeTimerBased
	SetFixedCount(value uint32) MacsecStaticKeyRekeyModeTimerBased
	// HasFixedCount checks if FixedCount has been set in MacsecStaticKeyRekeyModeTimerBased
	HasFixedCount() bool
	// Interval returns uint32, set in MacsecStaticKeyRekeyModeTimerBased.
	Interval() uint32
	// SetInterval assigns uint32 provided by user to MacsecStaticKeyRekeyModeTimerBased
	SetInterval(value uint32) MacsecStaticKeyRekeyModeTimerBased
	// HasInterval checks if Interval has been set in MacsecStaticKeyRekeyModeTimerBased
	HasInterval() bool
}

type MacsecStaticKeyRekeyModeTimerBasedChoiceEnum string

// Enum of Choice on MacsecStaticKeyRekeyModeTimerBased
var MacsecStaticKeyRekeyModeTimerBasedChoice = struct {
	CONTINUOUS  MacsecStaticKeyRekeyModeTimerBasedChoiceEnum
	FIXED_COUNT MacsecStaticKeyRekeyModeTimerBasedChoiceEnum
}{
	CONTINUOUS:  MacsecStaticKeyRekeyModeTimerBasedChoiceEnum("continuous"),
	FIXED_COUNT: MacsecStaticKeyRekeyModeTimerBasedChoiceEnum("fixed_count"),
}

func (obj *macsecStaticKeyRekeyModeTimerBased) Choice() MacsecStaticKeyRekeyModeTimerBasedChoiceEnum {
	return MacsecStaticKeyRekeyModeTimerBasedChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Continuous to set choice
func (obj *macsecStaticKeyRekeyModeTimerBased) Continuous() {
	obj.setChoice(MacsecStaticKeyRekeyModeTimerBasedChoice.CONTINUOUS)
}

// Periodic rekey attempt choices.
// Choice returns a string
func (obj *macsecStaticKeyRekeyModeTimerBased) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *macsecStaticKeyRekeyModeTimerBased) setChoice(value MacsecStaticKeyRekeyModeTimerBasedChoiceEnum) MacsecStaticKeyRekeyModeTimerBased {
	intValue, ok := otg.MacsecStaticKeyRekeyModeTimerBased_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecStaticKeyRekeyModeTimerBasedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecStaticKeyRekeyModeTimerBased_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.FixedCount = nil

	if value == MacsecStaticKeyRekeyModeTimerBasedChoice.FIXED_COUNT {
		defaultValue := uint32(10)
		obj.obj.FixedCount = &defaultValue
	}

	return obj
}

// Fixed rekey attempts.
// FixedCount returns a uint32
func (obj *macsecStaticKeyRekeyModeTimerBased) FixedCount() uint32 {

	if obj.obj.FixedCount == nil {
		obj.setChoice(MacsecStaticKeyRekeyModeTimerBasedChoice.FIXED_COUNT)
	}

	return *obj.obj.FixedCount

}

// Fixed rekey attempts.
// FixedCount returns a uint32
func (obj *macsecStaticKeyRekeyModeTimerBased) HasFixedCount() bool {
	return obj.obj.FixedCount != nil
}

// Fixed rekey attempts.
// SetFixedCount sets the uint32 value in the MacsecStaticKeyRekeyModeTimerBased object
func (obj *macsecStaticKeyRekeyModeTimerBased) SetFixedCount(value uint32) MacsecStaticKeyRekeyModeTimerBased {
	obj.setChoice(MacsecStaticKeyRekeyModeTimerBasedChoice.FIXED_COUNT)
	obj.obj.FixedCount = &value
	return obj
}

// Periodic rekey interval (sec).
// Interval returns a uint32
func (obj *macsecStaticKeyRekeyModeTimerBased) Interval() uint32 {

	return *obj.obj.Interval

}

// Periodic rekey interval (sec).
// Interval returns a uint32
func (obj *macsecStaticKeyRekeyModeTimerBased) HasInterval() bool {
	return obj.obj.Interval != nil
}

// Periodic rekey interval (sec).
// SetInterval sets the uint32 value in the MacsecStaticKeyRekeyModeTimerBased object
func (obj *macsecStaticKeyRekeyModeTimerBased) SetInterval(value uint32) MacsecStaticKeyRekeyModeTimerBased {

	obj.obj.Interval = &value
	return obj
}

func (obj *macsecStaticKeyRekeyModeTimerBased) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.FixedCount != nil {

		if *obj.obj.FixedCount < 1 || *obj.obj.FixedCount > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MacsecStaticKeyRekeyModeTimerBased.FixedCount <= 65535 but Got %d", *obj.obj.FixedCount))
		}

	}

	if obj.obj.Interval != nil {

		if *obj.obj.Interval < 30 || *obj.obj.Interval > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("30 <= MacsecStaticKeyRekeyModeTimerBased.Interval <= 65535 but Got %d", *obj.obj.Interval))
		}

	}

}

func (obj *macsecStaticKeyRekeyModeTimerBased) setDefault() {
	var choices_set int = 0
	var choice MacsecStaticKeyRekeyModeTimerBasedChoiceEnum

	if obj.obj.FixedCount != nil {
		choices_set += 1
		choice = MacsecStaticKeyRekeyModeTimerBasedChoice.FIXED_COUNT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MacsecStaticKeyRekeyModeTimerBasedChoice.CONTINUOUS)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MacsecStaticKeyRekeyModeTimerBased")
			}
		} else {
			intVal := otg.MacsecStaticKeyRekeyModeTimerBased_Choice_Enum_value[string(choice)]
			enumValue := otg.MacsecStaticKeyRekeyModeTimerBased_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.Interval == nil {
		obj.SetInterval(300)
	}

}

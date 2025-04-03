package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** EgressOnlyTrackingMetricTagValue *****
type egressOnlyTrackingMetricTagValue struct {
	validation
	obj          *otg.EgressOnlyTrackingMetricTagValue
	marshaller   marshalEgressOnlyTrackingMetricTagValue
	unMarshaller unMarshalEgressOnlyTrackingMetricTagValue
}

func NewEgressOnlyTrackingMetricTagValue() EgressOnlyTrackingMetricTagValue {
	obj := egressOnlyTrackingMetricTagValue{obj: &otg.EgressOnlyTrackingMetricTagValue{}}
	obj.setDefault()
	return &obj
}

func (obj *egressOnlyTrackingMetricTagValue) msg() *otg.EgressOnlyTrackingMetricTagValue {
	return obj.obj
}

func (obj *egressOnlyTrackingMetricTagValue) setMsg(msg *otg.EgressOnlyTrackingMetricTagValue) EgressOnlyTrackingMetricTagValue {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalegressOnlyTrackingMetricTagValue struct {
	obj *egressOnlyTrackingMetricTagValue
}

type marshalEgressOnlyTrackingMetricTagValue interface {
	// ToProto marshals EgressOnlyTrackingMetricTagValue to protobuf object *otg.EgressOnlyTrackingMetricTagValue
	ToProto() (*otg.EgressOnlyTrackingMetricTagValue, error)
	// ToPbText marshals EgressOnlyTrackingMetricTagValue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EgressOnlyTrackingMetricTagValue to YAML text
	ToYaml() (string, error)
	// ToJson marshals EgressOnlyTrackingMetricTagValue to JSON text
	ToJson() (string, error)
}

type unMarshalegressOnlyTrackingMetricTagValue struct {
	obj *egressOnlyTrackingMetricTagValue
}

type unMarshalEgressOnlyTrackingMetricTagValue interface {
	// FromProto unmarshals EgressOnlyTrackingMetricTagValue from protobuf object *otg.EgressOnlyTrackingMetricTagValue
	FromProto(msg *otg.EgressOnlyTrackingMetricTagValue) (EgressOnlyTrackingMetricTagValue, error)
	// FromPbText unmarshals EgressOnlyTrackingMetricTagValue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EgressOnlyTrackingMetricTagValue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EgressOnlyTrackingMetricTagValue from JSON text
	FromJson(value string) error
}

func (obj *egressOnlyTrackingMetricTagValue) Marshal() marshalEgressOnlyTrackingMetricTagValue {
	if obj.marshaller == nil {
		obj.marshaller = &marshalegressOnlyTrackingMetricTagValue{obj: obj}
	}
	return obj.marshaller
}

func (obj *egressOnlyTrackingMetricTagValue) Unmarshal() unMarshalEgressOnlyTrackingMetricTagValue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalegressOnlyTrackingMetricTagValue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalegressOnlyTrackingMetricTagValue) ToProto() (*otg.EgressOnlyTrackingMetricTagValue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalegressOnlyTrackingMetricTagValue) FromProto(msg *otg.EgressOnlyTrackingMetricTagValue) (EgressOnlyTrackingMetricTagValue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalegressOnlyTrackingMetricTagValue) ToPbText() (string, error) {
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

func (m *unMarshalegressOnlyTrackingMetricTagValue) FromPbText(value string) error {
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

func (m *marshalegressOnlyTrackingMetricTagValue) ToYaml() (string, error) {
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

func (m *unMarshalegressOnlyTrackingMetricTagValue) FromYaml(value string) error {
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

func (m *marshalegressOnlyTrackingMetricTagValue) ToJson() (string, error) {
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

func (m *unMarshalegressOnlyTrackingMetricTagValue) FromJson(value string) error {
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

func (obj *egressOnlyTrackingMetricTagValue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingMetricTagValue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingMetricTagValue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *egressOnlyTrackingMetricTagValue) Clone() (EgressOnlyTrackingMetricTagValue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEgressOnlyTrackingMetricTagValue()
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

// EgressOnlyTrackingMetricTagValue is a container for metric tag value
type EgressOnlyTrackingMetricTagValue interface {
	Validation
	// msg marshals EgressOnlyTrackingMetricTagValue to protobuf object *otg.EgressOnlyTrackingMetricTagValue
	// and doesn't set defaults
	msg() *otg.EgressOnlyTrackingMetricTagValue
	// setMsg unmarshals EgressOnlyTrackingMetricTagValue from protobuf object *otg.EgressOnlyTrackingMetricTagValue
	// and doesn't set defaults
	setMsg(*otg.EgressOnlyTrackingMetricTagValue) EgressOnlyTrackingMetricTagValue
	// provides marshal interface
	Marshal() marshalEgressOnlyTrackingMetricTagValue
	// provides unmarshal interface
	Unmarshal() unMarshalEgressOnlyTrackingMetricTagValue
	// validate validates EgressOnlyTrackingMetricTagValue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EgressOnlyTrackingMetricTagValue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns EgressOnlyTrackingMetricTagValueChoiceEnum, set in EgressOnlyTrackingMetricTagValue
	Choice() EgressOnlyTrackingMetricTagValueChoiceEnum
	// setChoice assigns EgressOnlyTrackingMetricTagValueChoiceEnum provided by user to EgressOnlyTrackingMetricTagValue
	setChoice(value EgressOnlyTrackingMetricTagValueChoiceEnum) EgressOnlyTrackingMetricTagValue
	// HasChoice checks if Choice has been set in EgressOnlyTrackingMetricTagValue
	HasChoice() bool
	// Hex returns string, set in EgressOnlyTrackingMetricTagValue.
	Hex() string
	// SetHex assigns string provided by user to EgressOnlyTrackingMetricTagValue
	SetHex(value string) EgressOnlyTrackingMetricTagValue
	// HasHex checks if Hex has been set in EgressOnlyTrackingMetricTagValue
	HasHex() bool
	// Str returns string, set in EgressOnlyTrackingMetricTagValue.
	Str() string
	// SetStr assigns string provided by user to EgressOnlyTrackingMetricTagValue
	SetStr(value string) EgressOnlyTrackingMetricTagValue
	// HasStr checks if Str has been set in EgressOnlyTrackingMetricTagValue
	HasStr() bool
}

type EgressOnlyTrackingMetricTagValueChoiceEnum string

// Enum of Choice on EgressOnlyTrackingMetricTagValue
var EgressOnlyTrackingMetricTagValueChoice = struct {
	HEX EgressOnlyTrackingMetricTagValueChoiceEnum
	STR EgressOnlyTrackingMetricTagValueChoiceEnum
}{
	HEX: EgressOnlyTrackingMetricTagValueChoiceEnum("hex"),
	STR: EgressOnlyTrackingMetricTagValueChoiceEnum("str"),
}

func (obj *egressOnlyTrackingMetricTagValue) Choice() EgressOnlyTrackingMetricTagValueChoiceEnum {
	return EgressOnlyTrackingMetricTagValueChoiceEnum(obj.obj.Choice.Enum().String())
}

// Available formats for metric tag value
// Choice returns a string
func (obj *egressOnlyTrackingMetricTagValue) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *egressOnlyTrackingMetricTagValue) setChoice(value EgressOnlyTrackingMetricTagValueChoiceEnum) EgressOnlyTrackingMetricTagValue {
	intValue, ok := otg.EgressOnlyTrackingMetricTagValue_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on EgressOnlyTrackingMetricTagValueChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.EgressOnlyTrackingMetricTagValue_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Str = nil
	obj.obj.Hex = nil
	return obj
}

// Value represented in hexadecimal format
// Hex returns a string
func (obj *egressOnlyTrackingMetricTagValue) Hex() string {

	if obj.obj.Hex == nil {
		obj.setChoice(EgressOnlyTrackingMetricTagValueChoice.HEX)
	}

	return *obj.obj.Hex

}

// Value represented in hexadecimal format
// Hex returns a string
func (obj *egressOnlyTrackingMetricTagValue) HasHex() bool {
	return obj.obj.Hex != nil
}

// Value represented in hexadecimal format
// SetHex sets the string value in the EgressOnlyTrackingMetricTagValue object
func (obj *egressOnlyTrackingMetricTagValue) SetHex(value string) EgressOnlyTrackingMetricTagValue {
	obj.setChoice(EgressOnlyTrackingMetricTagValueChoice.HEX)
	obj.obj.Hex = &value
	return obj
}

// Value represented in string format
// Str returns a string
func (obj *egressOnlyTrackingMetricTagValue) Str() string {

	if obj.obj.Str == nil {
		obj.setChoice(EgressOnlyTrackingMetricTagValueChoice.STR)
	}

	return *obj.obj.Str

}

// Value represented in string format
// Str returns a string
func (obj *egressOnlyTrackingMetricTagValue) HasStr() bool {
	return obj.obj.Str != nil
}

// Value represented in string format
// SetStr sets the string value in the EgressOnlyTrackingMetricTagValue object
func (obj *egressOnlyTrackingMetricTagValue) SetStr(value string) EgressOnlyTrackingMetricTagValue {
	obj.setChoice(EgressOnlyTrackingMetricTagValueChoice.STR)
	obj.obj.Str = &value
	return obj
}

func (obj *egressOnlyTrackingMetricTagValue) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Hex != nil {

		err := obj.validateHex(obj.Hex())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on EgressOnlyTrackingMetricTagValue.Hex"))
		}

	}

}

func (obj *egressOnlyTrackingMetricTagValue) setDefault() {
	var choices_set int = 0
	var choice EgressOnlyTrackingMetricTagValueChoiceEnum

	if obj.obj.Hex != nil {
		choices_set += 1
		choice = EgressOnlyTrackingMetricTagValueChoice.HEX
	}

	if obj.obj.Str != nil {
		choices_set += 1
		choice = EgressOnlyTrackingMetricTagValueChoice.STR
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(EgressOnlyTrackingMetricTagValueChoice.HEX)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in EgressOnlyTrackingMetricTagValue")
			}
		} else {
			intVal := otg.EgressOnlyTrackingMetricTagValue_Choice_Enum_value[string(choice)]
			enumValue := otg.EgressOnlyTrackingMetricTagValue_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

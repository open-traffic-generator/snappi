package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2TransmitType *****
type rocev2TransmitType struct {
	validation
	obj                  *otg.Rocev2TransmitType
	marshaller           marshalRocev2TransmitType
	unMarshaller         unMarshalRocev2TransmitType
	targetLineRateHolder Rocev2TargetLineRate
}

func NewRocev2TransmitType() Rocev2TransmitType {
	obj := rocev2TransmitType{obj: &otg.Rocev2TransmitType{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2TransmitType) msg() *otg.Rocev2TransmitType {
	return obj.obj
}

func (obj *rocev2TransmitType) setMsg(msg *otg.Rocev2TransmitType) Rocev2TransmitType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2TransmitType struct {
	obj *rocev2TransmitType
}

type marshalRocev2TransmitType interface {
	// ToProto marshals Rocev2TransmitType to protobuf object *otg.Rocev2TransmitType
	ToProto() (*otg.Rocev2TransmitType, error)
	// ToPbText marshals Rocev2TransmitType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2TransmitType to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2TransmitType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Rocev2TransmitType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrocev2TransmitType struct {
	obj *rocev2TransmitType
}

type unMarshalRocev2TransmitType interface {
	// FromProto unmarshals Rocev2TransmitType from protobuf object *otg.Rocev2TransmitType
	FromProto(msg *otg.Rocev2TransmitType) (Rocev2TransmitType, error)
	// FromPbText unmarshals Rocev2TransmitType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2TransmitType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2TransmitType from JSON text
	FromJson(value string) error
}

func (obj *rocev2TransmitType) Marshal() marshalRocev2TransmitType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2TransmitType{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2TransmitType) Unmarshal() unMarshalRocev2TransmitType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2TransmitType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2TransmitType) ToProto() (*otg.Rocev2TransmitType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2TransmitType) FromProto(msg *otg.Rocev2TransmitType) (Rocev2TransmitType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2TransmitType) ToPbText() (string, error) {
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

func (m *unMarshalrocev2TransmitType) FromPbText(value string) error {
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

func (m *marshalrocev2TransmitType) ToYaml() (string, error) {
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

func (m *unMarshalrocev2TransmitType) FromYaml(value string) error {
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

func (m *marshalrocev2TransmitType) ToJsonRaw() (string, error) {
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

func (m *marshalrocev2TransmitType) ToJson() (string, error) {
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

func (m *unMarshalrocev2TransmitType) FromJson(value string) error {
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

func (obj *rocev2TransmitType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2TransmitType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2TransmitType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2TransmitType) Clone() (Rocev2TransmitType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2TransmitType()
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

func (obj *rocev2TransmitType) setNil() {
	obj.targetLineRateHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2TransmitType is roCEv2 flows can be configured to run in continuous mode or fixed iteration.
type Rocev2TransmitType interface {
	Validation
	// msg marshals Rocev2TransmitType to protobuf object *otg.Rocev2TransmitType
	// and doesn't set defaults
	msg() *otg.Rocev2TransmitType
	// setMsg unmarshals Rocev2TransmitType from protobuf object *otg.Rocev2TransmitType
	// and doesn't set defaults
	setMsg(*otg.Rocev2TransmitType) Rocev2TransmitType
	// provides marshal interface
	Marshal() marshalRocev2TransmitType
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2TransmitType
	// validate validates Rocev2TransmitType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2TransmitType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Rocev2TransmitTypeChoiceEnum, set in Rocev2TransmitType
	Choice() Rocev2TransmitTypeChoiceEnum
	// setChoice assigns Rocev2TransmitTypeChoiceEnum provided by user to Rocev2TransmitType
	setChoice(value Rocev2TransmitTypeChoiceEnum) Rocev2TransmitType
	// HasChoice checks if Choice has been set in Rocev2TransmitType
	HasChoice() bool
	// TargetLineRate returns Rocev2TargetLineRate, set in Rocev2TransmitType.
	// Rocev2TargetLineRate is configure target line rate of traffic rate on this port as percentage of link speed.
	TargetLineRate() Rocev2TargetLineRate
	// SetTargetLineRate assigns Rocev2TargetLineRate provided by user to Rocev2TransmitType.
	// Rocev2TargetLineRate is configure target line rate of traffic rate on this port as percentage of link speed.
	SetTargetLineRate(value Rocev2TargetLineRate) Rocev2TransmitType
	// HasTargetLineRate checks if TargetLineRate has been set in Rocev2TransmitType
	HasTargetLineRate() bool
	setNil()
}

type Rocev2TransmitTypeChoiceEnum string

// Enum of Choice on Rocev2TransmitType
var Rocev2TransmitTypeChoice = struct {
	TARGET_LINE_RATE Rocev2TransmitTypeChoiceEnum
}{
	TARGET_LINE_RATE: Rocev2TransmitTypeChoiceEnum("target_line_rate"),
}

func (obj *rocev2TransmitType) Choice() Rocev2TransmitTypeChoiceEnum {
	return Rocev2TransmitTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *rocev2TransmitType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *rocev2TransmitType) setChoice(value Rocev2TransmitTypeChoiceEnum) Rocev2TransmitType {
	intValue, ok := otg.Rocev2TransmitType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Rocev2TransmitTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Rocev2TransmitType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.TargetLineRate = nil
	obj.targetLineRateHolder = nil

	if value == Rocev2TransmitTypeChoice.TARGET_LINE_RATE {
		obj.obj.TargetLineRate = NewRocev2TargetLineRate().msg()
	}

	return obj
}

// description is TBD
// TargetLineRate returns a Rocev2TargetLineRate
func (obj *rocev2TransmitType) TargetLineRate() Rocev2TargetLineRate {
	if obj.obj.TargetLineRate == nil {
		obj.setChoice(Rocev2TransmitTypeChoice.TARGET_LINE_RATE)
	}
	if obj.targetLineRateHolder == nil {
		obj.targetLineRateHolder = &rocev2TargetLineRate{obj: obj.obj.TargetLineRate}
	}
	return obj.targetLineRateHolder
}

// description is TBD
// TargetLineRate returns a Rocev2TargetLineRate
func (obj *rocev2TransmitType) HasTargetLineRate() bool {
	return obj.obj.TargetLineRate != nil
}

// description is TBD
// SetTargetLineRate sets the Rocev2TargetLineRate value in the Rocev2TransmitType object
func (obj *rocev2TransmitType) SetTargetLineRate(value Rocev2TargetLineRate) Rocev2TransmitType {
	obj.setChoice(Rocev2TransmitTypeChoice.TARGET_LINE_RATE)
	obj.targetLineRateHolder = nil
	obj.obj.TargetLineRate = value.msg()

	return obj
}

func (obj *rocev2TransmitType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TargetLineRate != nil {

		obj.TargetLineRate().validateObj(vObj, set_default)
	}

}

func (obj *rocev2TransmitType) setDefault() {
	var choices_set int = 0
	var choice Rocev2TransmitTypeChoiceEnum

	if obj.obj.TargetLineRate != nil {
		choices_set += 1
		choice = Rocev2TransmitTypeChoice.TARGET_LINE_RATE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Rocev2TransmitTypeChoice.TARGET_LINE_RATE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Rocev2TransmitType")
			}
		} else {
			intVal := otg.Rocev2TransmitType_Choice_Enum_value[string(choice)]
			enumValue := otg.Rocev2TransmitType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

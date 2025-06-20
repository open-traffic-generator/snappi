package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowDurationInterBurstGap *****
type flowDurationInterBurstGap struct {
	validation
	obj          *otg.FlowDurationInterBurstGap
	marshaller   marshalFlowDurationInterBurstGap
	unMarshaller unMarshalFlowDurationInterBurstGap
}

func NewFlowDurationInterBurstGap() FlowDurationInterBurstGap {
	obj := flowDurationInterBurstGap{obj: &otg.FlowDurationInterBurstGap{}}
	obj.setDefault()
	return &obj
}

func (obj *flowDurationInterBurstGap) msg() *otg.FlowDurationInterBurstGap {
	return obj.obj
}

func (obj *flowDurationInterBurstGap) setMsg(msg *otg.FlowDurationInterBurstGap) FlowDurationInterBurstGap {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowDurationInterBurstGap struct {
	obj *flowDurationInterBurstGap
}

type marshalFlowDurationInterBurstGap interface {
	// ToProto marshals FlowDurationInterBurstGap to protobuf object *otg.FlowDurationInterBurstGap
	ToProto() (*otg.FlowDurationInterBurstGap, error)
	// ToPbText marshals FlowDurationInterBurstGap to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowDurationInterBurstGap to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowDurationInterBurstGap to JSON text
	ToJson() (string, error)
}

type unMarshalflowDurationInterBurstGap struct {
	obj *flowDurationInterBurstGap
}

type unMarshalFlowDurationInterBurstGap interface {
	// FromProto unmarshals FlowDurationInterBurstGap from protobuf object *otg.FlowDurationInterBurstGap
	FromProto(msg *otg.FlowDurationInterBurstGap) (FlowDurationInterBurstGap, error)
	// FromPbText unmarshals FlowDurationInterBurstGap from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowDurationInterBurstGap from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowDurationInterBurstGap from JSON text
	FromJson(value string) error
}

func (obj *flowDurationInterBurstGap) Marshal() marshalFlowDurationInterBurstGap {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowDurationInterBurstGap{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowDurationInterBurstGap) Unmarshal() unMarshalFlowDurationInterBurstGap {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowDurationInterBurstGap{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowDurationInterBurstGap) ToProto() (*otg.FlowDurationInterBurstGap, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowDurationInterBurstGap) FromProto(msg *otg.FlowDurationInterBurstGap) (FlowDurationInterBurstGap, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowDurationInterBurstGap) ToPbText() (string, error) {
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

func (m *unMarshalflowDurationInterBurstGap) FromPbText(value string) error {
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

func (m *marshalflowDurationInterBurstGap) ToYaml() (string, error) {
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

func (m *unMarshalflowDurationInterBurstGap) FromYaml(value string) error {
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

func (m *marshalflowDurationInterBurstGap) ToJson() (string, error) {
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

func (m *unMarshalflowDurationInterBurstGap) FromJson(value string) error {
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

func (obj *flowDurationInterBurstGap) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowDurationInterBurstGap) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowDurationInterBurstGap) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowDurationInterBurstGap) Clone() (FlowDurationInterBurstGap, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowDurationInterBurstGap()
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

// FlowDurationInterBurstGap is the optional container for specifying a gap between bursts.
type FlowDurationInterBurstGap interface {
	Validation
	// msg marshals FlowDurationInterBurstGap to protobuf object *otg.FlowDurationInterBurstGap
	// and doesn't set defaults
	msg() *otg.FlowDurationInterBurstGap
	// setMsg unmarshals FlowDurationInterBurstGap from protobuf object *otg.FlowDurationInterBurstGap
	// and doesn't set defaults
	setMsg(*otg.FlowDurationInterBurstGap) FlowDurationInterBurstGap
	// provides marshal interface
	Marshal() marshalFlowDurationInterBurstGap
	// provides unmarshal interface
	Unmarshal() unMarshalFlowDurationInterBurstGap
	// validate validates FlowDurationInterBurstGap
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowDurationInterBurstGap, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowDurationInterBurstGapChoiceEnum, set in FlowDurationInterBurstGap
	Choice() FlowDurationInterBurstGapChoiceEnum
	// setChoice assigns FlowDurationInterBurstGapChoiceEnum provided by user to FlowDurationInterBurstGap
	setChoice(value FlowDurationInterBurstGapChoiceEnum) FlowDurationInterBurstGap
	// HasChoice checks if Choice has been set in FlowDurationInterBurstGap
	HasChoice() bool
	// Bytes returns float64, set in FlowDurationInterBurstGap.
	Bytes() float64
	// SetBytes assigns float64 provided by user to FlowDurationInterBurstGap
	SetBytes(value float64) FlowDurationInterBurstGap
	// HasBytes checks if Bytes has been set in FlowDurationInterBurstGap
	HasBytes() bool
	// Nanoseconds returns float64, set in FlowDurationInterBurstGap.
	Nanoseconds() float64
	// SetNanoseconds assigns float64 provided by user to FlowDurationInterBurstGap
	SetNanoseconds(value float64) FlowDurationInterBurstGap
	// HasNanoseconds checks if Nanoseconds has been set in FlowDurationInterBurstGap
	HasNanoseconds() bool
	// Microseconds returns float64, set in FlowDurationInterBurstGap.
	Microseconds() float64
	// SetMicroseconds assigns float64 provided by user to FlowDurationInterBurstGap
	SetMicroseconds(value float64) FlowDurationInterBurstGap
	// HasMicroseconds checks if Microseconds has been set in FlowDurationInterBurstGap
	HasMicroseconds() bool
}

type FlowDurationInterBurstGapChoiceEnum string

// Enum of Choice on FlowDurationInterBurstGap
var FlowDurationInterBurstGapChoice = struct {
	BYTES        FlowDurationInterBurstGapChoiceEnum
	NANOSECONDS  FlowDurationInterBurstGapChoiceEnum
	MICROSECONDS FlowDurationInterBurstGapChoiceEnum
}{
	BYTES:        FlowDurationInterBurstGapChoiceEnum("bytes"),
	NANOSECONDS:  FlowDurationInterBurstGapChoiceEnum("nanoseconds"),
	MICROSECONDS: FlowDurationInterBurstGapChoiceEnum("microseconds"),
}

func (obj *flowDurationInterBurstGap) Choice() FlowDurationInterBurstGapChoiceEnum {
	return FlowDurationInterBurstGapChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of inter burst gap units.
// Choice returns a string
func (obj *flowDurationInterBurstGap) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowDurationInterBurstGap) setChoice(value FlowDurationInterBurstGapChoiceEnum) FlowDurationInterBurstGap {
	intValue, ok := otg.FlowDurationInterBurstGap_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowDurationInterBurstGapChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowDurationInterBurstGap_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Microseconds = nil
	obj.obj.Nanoseconds = nil
	obj.obj.Bytes = nil

	if value == FlowDurationInterBurstGapChoice.BYTES {
		defaultValue := float64(12)
		obj.obj.Bytes = &defaultValue
	}

	if value == FlowDurationInterBurstGapChoice.NANOSECONDS {
		defaultValue := float64(96)
		obj.obj.Nanoseconds = &defaultValue
	}

	if value == FlowDurationInterBurstGapChoice.MICROSECONDS {
		defaultValue := float64(0.096)
		obj.obj.Microseconds = &defaultValue
	}

	return obj
}

// The amount of time between bursts expressed in bytes.
// A value of 0 indicates no gap between bursts.
// Bytes returns a float64
func (obj *flowDurationInterBurstGap) Bytes() float64 {

	if obj.obj.Bytes == nil {
		obj.setChoice(FlowDurationInterBurstGapChoice.BYTES)
	}

	return *obj.obj.Bytes

}

// The amount of time between bursts expressed in bytes.
// A value of 0 indicates no gap between bursts.
// Bytes returns a float64
func (obj *flowDurationInterBurstGap) HasBytes() bool {
	return obj.obj.Bytes != nil
}

// The amount of time between bursts expressed in bytes.
// A value of 0 indicates no gap between bursts.
// SetBytes sets the float64 value in the FlowDurationInterBurstGap object
func (obj *flowDurationInterBurstGap) SetBytes(value float64) FlowDurationInterBurstGap {
	obj.setChoice(FlowDurationInterBurstGapChoice.BYTES)
	obj.obj.Bytes = &value
	return obj
}

// The amount of time between bursts expressed in nanoseconds.
// A value of 0 indicates no gap between bursts.
// Nanoseconds returns a float64
func (obj *flowDurationInterBurstGap) Nanoseconds() float64 {

	if obj.obj.Nanoseconds == nil {
		obj.setChoice(FlowDurationInterBurstGapChoice.NANOSECONDS)
	}

	return *obj.obj.Nanoseconds

}

// The amount of time between bursts expressed in nanoseconds.
// A value of 0 indicates no gap between bursts.
// Nanoseconds returns a float64
func (obj *flowDurationInterBurstGap) HasNanoseconds() bool {
	return obj.obj.Nanoseconds != nil
}

// The amount of time between bursts expressed in nanoseconds.
// A value of 0 indicates no gap between bursts.
// SetNanoseconds sets the float64 value in the FlowDurationInterBurstGap object
func (obj *flowDurationInterBurstGap) SetNanoseconds(value float64) FlowDurationInterBurstGap {
	obj.setChoice(FlowDurationInterBurstGapChoice.NANOSECONDS)
	obj.obj.Nanoseconds = &value
	return obj
}

// The amount of time between bursts expressed in microseconds.
// A value of 0 indicates no gap between bursts.
// Microseconds returns a float64
func (obj *flowDurationInterBurstGap) Microseconds() float64 {

	if obj.obj.Microseconds == nil {
		obj.setChoice(FlowDurationInterBurstGapChoice.MICROSECONDS)
	}

	return *obj.obj.Microseconds

}

// The amount of time between bursts expressed in microseconds.
// A value of 0 indicates no gap between bursts.
// Microseconds returns a float64
func (obj *flowDurationInterBurstGap) HasMicroseconds() bool {
	return obj.obj.Microseconds != nil
}

// The amount of time between bursts expressed in microseconds.
// A value of 0 indicates no gap between bursts.
// SetMicroseconds sets the float64 value in the FlowDurationInterBurstGap object
func (obj *flowDurationInterBurstGap) SetMicroseconds(value float64) FlowDurationInterBurstGap {
	obj.setChoice(FlowDurationInterBurstGapChoice.MICROSECONDS)
	obj.obj.Microseconds = &value
	return obj
}

func (obj *flowDurationInterBurstGap) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Bytes != nil {

		if *obj.obj.Bytes < 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowDurationInterBurstGap.Bytes <= max(float64) but Got %f", *obj.obj.Bytes))
		}

	}

	if obj.obj.Nanoseconds != nil {

		if *obj.obj.Nanoseconds < 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowDurationInterBurstGap.Nanoseconds <= max(float64) but Got %f", *obj.obj.Nanoseconds))
		}

	}

	if obj.obj.Microseconds != nil {

		if *obj.obj.Microseconds < 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowDurationInterBurstGap.Microseconds <= max(float64) but Got %f", *obj.obj.Microseconds))
		}

	}

}

func (obj *flowDurationInterBurstGap) setDefault() {
	var choices_set int = 0
	var choice FlowDurationInterBurstGapChoiceEnum

	if obj.obj.Bytes != nil {
		choices_set += 1
		choice = FlowDurationInterBurstGapChoice.BYTES
	}

	if obj.obj.Nanoseconds != nil {
		choices_set += 1
		choice = FlowDurationInterBurstGapChoice.NANOSECONDS
	}

	if obj.obj.Microseconds != nil {
		choices_set += 1
		choice = FlowDurationInterBurstGapChoice.MICROSECONDS
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowDurationInterBurstGapChoice.BYTES)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowDurationInterBurstGap")
			}
		} else {
			intVal := otg.FlowDurationInterBurstGap_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowDurationInterBurstGap_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

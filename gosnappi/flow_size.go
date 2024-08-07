package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowSize *****
type flowSize struct {
	validation
	obj               *otg.FlowSize
	marshaller        marshalFlowSize
	unMarshaller      unMarshalFlowSize
	incrementHolder   FlowSizeIncrement
	randomHolder      FlowSizeRandom
	weightPairsHolder FlowSizeWeightPairs
}

func NewFlowSize() FlowSize {
	obj := flowSize{obj: &otg.FlowSize{}}
	obj.setDefault()
	return &obj
}

func (obj *flowSize) msg() *otg.FlowSize {
	return obj.obj
}

func (obj *flowSize) setMsg(msg *otg.FlowSize) FlowSize {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowSize struct {
	obj *flowSize
}

type marshalFlowSize interface {
	// ToProto marshals FlowSize to protobuf object *otg.FlowSize
	ToProto() (*otg.FlowSize, error)
	// ToPbText marshals FlowSize to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowSize to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowSize to JSON text
	ToJson() (string, error)
}

type unMarshalflowSize struct {
	obj *flowSize
}

type unMarshalFlowSize interface {
	// FromProto unmarshals FlowSize from protobuf object *otg.FlowSize
	FromProto(msg *otg.FlowSize) (FlowSize, error)
	// FromPbText unmarshals FlowSize from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowSize from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowSize from JSON text
	FromJson(value string) error
}

func (obj *flowSize) Marshal() marshalFlowSize {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowSize{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowSize) Unmarshal() unMarshalFlowSize {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowSize{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowSize) ToProto() (*otg.FlowSize, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowSize) FromProto(msg *otg.FlowSize) (FlowSize, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowSize) ToPbText() (string, error) {
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

func (m *unMarshalflowSize) FromPbText(value string) error {
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

func (m *marshalflowSize) ToYaml() (string, error) {
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

func (m *unMarshalflowSize) FromYaml(value string) error {
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

func (m *marshalflowSize) ToJson() (string, error) {
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

func (m *unMarshalflowSize) FromJson(value string) error {
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

func (obj *flowSize) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowSize) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowSize) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowSize) Clone() (FlowSize, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowSize()
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

func (obj *flowSize) setNil() {
	obj.incrementHolder = nil
	obj.randomHolder = nil
	obj.weightPairsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowSize is the frame size which overrides the total length of the packet
type FlowSize interface {
	Validation
	// msg marshals FlowSize to protobuf object *otg.FlowSize
	// and doesn't set defaults
	msg() *otg.FlowSize
	// setMsg unmarshals FlowSize from protobuf object *otg.FlowSize
	// and doesn't set defaults
	setMsg(*otg.FlowSize) FlowSize
	// provides marshal interface
	Marshal() marshalFlowSize
	// provides unmarshal interface
	Unmarshal() unMarshalFlowSize
	// validate validates FlowSize
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowSize, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowSizeChoiceEnum, set in FlowSize
	Choice() FlowSizeChoiceEnum
	// setChoice assigns FlowSizeChoiceEnum provided by user to FlowSize
	setChoice(value FlowSizeChoiceEnum) FlowSize
	// HasChoice checks if Choice has been set in FlowSize
	HasChoice() bool
	// Fixed returns uint32, set in FlowSize.
	Fixed() uint32
	// SetFixed assigns uint32 provided by user to FlowSize
	SetFixed(value uint32) FlowSize
	// HasFixed checks if Fixed has been set in FlowSize
	HasFixed() bool
	// Increment returns FlowSizeIncrement, set in FlowSize.
	// FlowSizeIncrement is frame size that increments from a starting size to
	// an ending size incrementing by a step size.
	Increment() FlowSizeIncrement
	// SetIncrement assigns FlowSizeIncrement provided by user to FlowSize.
	// FlowSizeIncrement is frame size that increments from a starting size to
	// an ending size incrementing by a step size.
	SetIncrement(value FlowSizeIncrement) FlowSize
	// HasIncrement checks if Increment has been set in FlowSize
	HasIncrement() bool
	// Random returns FlowSizeRandom, set in FlowSize.
	// FlowSizeRandom is random frame size from a min value to a max value.
	Random() FlowSizeRandom
	// SetRandom assigns FlowSizeRandom provided by user to FlowSize.
	// FlowSizeRandom is random frame size from a min value to a max value.
	SetRandom(value FlowSizeRandom) FlowSize
	// HasRandom checks if Random has been set in FlowSize
	HasRandom() bool
	// WeightPairs returns FlowSizeWeightPairs, set in FlowSize.
	// FlowSizeWeightPairs is frame size distribution, defined as <size, weight> pairs (including IMIX distribution).
	// Frames are randomly generated such that the proportion of each frame size out of the total number of frames
	// are matching with the weight value of the <size, weight> pair. However, as with any other probability
	// distribution, the sample distribution is close to theoretical value only if the size of the sample is reasonably large.
	// When the number of frames is very low the transmitted frames may not come close to the ratio described in the weight.
	WeightPairs() FlowSizeWeightPairs
	// SetWeightPairs assigns FlowSizeWeightPairs provided by user to FlowSize.
	// FlowSizeWeightPairs is frame size distribution, defined as <size, weight> pairs (including IMIX distribution).
	// Frames are randomly generated such that the proportion of each frame size out of the total number of frames
	// are matching with the weight value of the <size, weight> pair. However, as with any other probability
	// distribution, the sample distribution is close to theoretical value only if the size of the sample is reasonably large.
	// When the number of frames is very low the transmitted frames may not come close to the ratio described in the weight.
	SetWeightPairs(value FlowSizeWeightPairs) FlowSize
	// HasWeightPairs checks if WeightPairs has been set in FlowSize
	HasWeightPairs() bool
	setNil()
}

type FlowSizeChoiceEnum string

// Enum of Choice on FlowSize
var FlowSizeChoice = struct {
	FIXED        FlowSizeChoiceEnum
	INCREMENT    FlowSizeChoiceEnum
	RANDOM       FlowSizeChoiceEnum
	WEIGHT_PAIRS FlowSizeChoiceEnum
}{
	FIXED:        FlowSizeChoiceEnum("fixed"),
	INCREMENT:    FlowSizeChoiceEnum("increment"),
	RANDOM:       FlowSizeChoiceEnum("random"),
	WEIGHT_PAIRS: FlowSizeChoiceEnum("weight_pairs"),
}

func (obj *flowSize) Choice() FlowSizeChoiceEnum {
	return FlowSizeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowSize) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowSize) setChoice(value FlowSizeChoiceEnum) FlowSize {
	intValue, ok := otg.FlowSize_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowSizeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowSize_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.WeightPairs = nil
	obj.weightPairsHolder = nil
	obj.obj.Random = nil
	obj.randomHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Fixed = nil

	if value == FlowSizeChoice.FIXED {
		defaultValue := uint32(64)
		obj.obj.Fixed = &defaultValue
	}

	if value == FlowSizeChoice.INCREMENT {
		obj.obj.Increment = NewFlowSizeIncrement().msg()
	}

	if value == FlowSizeChoice.RANDOM {
		obj.obj.Random = NewFlowSizeRandom().msg()
	}

	if value == FlowSizeChoice.WEIGHT_PAIRS {
		obj.obj.WeightPairs = NewFlowSizeWeightPairs().msg()
	}

	return obj
}

// description is TBD
// Fixed returns a uint32
func (obj *flowSize) Fixed() uint32 {

	if obj.obj.Fixed == nil {
		obj.setChoice(FlowSizeChoice.FIXED)
	}

	return *obj.obj.Fixed

}

// description is TBD
// Fixed returns a uint32
func (obj *flowSize) HasFixed() bool {
	return obj.obj.Fixed != nil
}

// description is TBD
// SetFixed sets the uint32 value in the FlowSize object
func (obj *flowSize) SetFixed(value uint32) FlowSize {
	obj.setChoice(FlowSizeChoice.FIXED)
	obj.obj.Fixed = &value
	return obj
}

// description is TBD
// Increment returns a FlowSizeIncrement
func (obj *flowSize) Increment() FlowSizeIncrement {
	if obj.obj.Increment == nil {
		obj.setChoice(FlowSizeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &flowSizeIncrement{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a FlowSizeIncrement
func (obj *flowSize) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the FlowSizeIncrement value in the FlowSize object
func (obj *flowSize) SetIncrement(value FlowSizeIncrement) FlowSize {
	obj.setChoice(FlowSizeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Random returns a FlowSizeRandom
func (obj *flowSize) Random() FlowSizeRandom {
	if obj.obj.Random == nil {
		obj.setChoice(FlowSizeChoice.RANDOM)
	}
	if obj.randomHolder == nil {
		obj.randomHolder = &flowSizeRandom{obj: obj.obj.Random}
	}
	return obj.randomHolder
}

// description is TBD
// Random returns a FlowSizeRandom
func (obj *flowSize) HasRandom() bool {
	return obj.obj.Random != nil
}

// description is TBD
// SetRandom sets the FlowSizeRandom value in the FlowSize object
func (obj *flowSize) SetRandom(value FlowSizeRandom) FlowSize {
	obj.setChoice(FlowSizeChoice.RANDOM)
	obj.randomHolder = nil
	obj.obj.Random = value.msg()

	return obj
}

// description is TBD
// WeightPairs returns a FlowSizeWeightPairs
func (obj *flowSize) WeightPairs() FlowSizeWeightPairs {
	if obj.obj.WeightPairs == nil {
		obj.setChoice(FlowSizeChoice.WEIGHT_PAIRS)
	}
	if obj.weightPairsHolder == nil {
		obj.weightPairsHolder = &flowSizeWeightPairs{obj: obj.obj.WeightPairs}
	}
	return obj.weightPairsHolder
}

// description is TBD
// WeightPairs returns a FlowSizeWeightPairs
func (obj *flowSize) HasWeightPairs() bool {
	return obj.obj.WeightPairs != nil
}

// description is TBD
// SetWeightPairs sets the FlowSizeWeightPairs value in the FlowSize object
func (obj *flowSize) SetWeightPairs(value FlowSizeWeightPairs) FlowSize {
	obj.setChoice(FlowSizeChoice.WEIGHT_PAIRS)
	obj.weightPairsHolder = nil
	obj.obj.WeightPairs = value.msg()

	return obj
}

func (obj *flowSize) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Random != nil {

		obj.Random().validateObj(vObj, set_default)
	}

	if obj.obj.WeightPairs != nil {

		obj.WeightPairs().validateObj(vObj, set_default)
	}

}

func (obj *flowSize) setDefault() {
	var choices_set int = 0
	var choice FlowSizeChoiceEnum

	if obj.obj.Fixed != nil {
		choices_set += 1
		choice = FlowSizeChoice.FIXED
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = FlowSizeChoice.INCREMENT
	}

	if obj.obj.Random != nil {
		choices_set += 1
		choice = FlowSizeChoice.RANDOM
	}

	if obj.obj.WeightPairs != nil {
		choices_set += 1
		choice = FlowSizeChoice.WEIGHT_PAIRS
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowSizeChoice.FIXED)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowSize")
			}
		} else {
			intVal := otg.FlowSize_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowSize_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

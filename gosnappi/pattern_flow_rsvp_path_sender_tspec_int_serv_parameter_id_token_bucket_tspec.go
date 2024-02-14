package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec *****
type patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec struct {
	validation
	obj             *otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	marshaller      marshalPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	unMarshaller    unMarshalPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	incrementHolder PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	decrementHolder PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
}

func NewPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec() PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec {
	obj := patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec{obj: &otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) msg() *otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec struct {
	obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
}

type marshalPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec struct {
	obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) (PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) (PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) Clone() (PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec()
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

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec is parameter ID, parameter 127 (Token Bucket TSpec)
type PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	// validate validates PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoiceEnum, set in PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	Choice() PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoiceEnum
	// setChoice assigns PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoiceEnum provided by user to PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	setChoice(value PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoiceEnum) PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	// HasValue checks if Value has been set in PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	// Increment returns PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter, set in PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec.
	// PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter is integer counter pattern
	Increment() PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	// SetIncrement assigns PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter provided by user to PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec.
	// PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter, set in PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec.
	// PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	// SetDecrement assigns PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter provided by user to PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec.
	// PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
var PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoice = struct {
	VALUE     PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoiceEnum
	VALUES    PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoiceEnum
	INCREMENT PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoiceEnum
	DECREMENT PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) Choice() PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoiceEnum {
	return PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) setChoice(value PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoiceEnum) PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec {
	intValue, ok := otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoice.VALUE {
		defaultValue := uint32(127)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoice.VALUES {
		defaultValue := []uint32{127}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter().msg()
	}

	if value == PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec object
func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{127})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec object
func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) Increment() PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter value in the PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec object
func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) SetIncrement(value PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) Decrement() PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter value in the PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec object
func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) SetDecrement(value PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecChoice.VALUE)

	}

}

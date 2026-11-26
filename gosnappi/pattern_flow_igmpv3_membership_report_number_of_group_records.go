package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipReportNumberOfGroupRecords *****
type patternFlowIgmpv3MembershipReportNumberOfGroupRecords struct {
	validation
	obj             *otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	marshaller      marshalPatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	unMarshaller    unMarshalPatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	incrementHolder PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	decrementHolder PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
}

func NewPatternFlowIgmpv3MembershipReportNumberOfGroupRecords() PatternFlowIgmpv3MembershipReportNumberOfGroupRecords {
	obj := patternFlowIgmpv3MembershipReportNumberOfGroupRecords{obj: &otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecords{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) msg() *otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecords {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) setMsg(msg *otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecords) PatternFlowIgmpv3MembershipReportNumberOfGroupRecords {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecords struct {
	obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords
}

type marshalPatternFlowIgmpv3MembershipReportNumberOfGroupRecords interface {
	// ToProto marshals PatternFlowIgmpv3MembershipReportNumberOfGroupRecords to protobuf object *otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	ToProto() (*otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecords, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipReportNumberOfGroupRecords to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipReportNumberOfGroupRecords to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipReportNumberOfGroupRecords to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecords struct {
	obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords
}

type unMarshalPatternFlowIgmpv3MembershipReportNumberOfGroupRecords interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipReportNumberOfGroupRecords from protobuf object *otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	FromProto(msg *otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecords) (PatternFlowIgmpv3MembershipReportNumberOfGroupRecords, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipReportNumberOfGroupRecords from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipReportNumberOfGroupRecords from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipReportNumberOfGroupRecords from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) Marshal() marshalPatternFlowIgmpv3MembershipReportNumberOfGroupRecords {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecords{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) Unmarshal() unMarshalPatternFlowIgmpv3MembershipReportNumberOfGroupRecords {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecords{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecords) ToProto() (*otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecords, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecords) FromProto(msg *otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecords) (PatternFlowIgmpv3MembershipReportNumberOfGroupRecords, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecords) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecords) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecords) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecords) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecords) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecords) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) Clone() (PatternFlowIgmpv3MembershipReportNumberOfGroupRecords, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipReportNumberOfGroupRecords()
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

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv3MembershipReportNumberOfGroupRecords is number of group records. When auto, it is set to the number of entries in group_records.
type PatternFlowIgmpv3MembershipReportNumberOfGroupRecords interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipReportNumberOfGroupRecords to protobuf object *otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	// setMsg unmarshals PatternFlowIgmpv3MembershipReportNumberOfGroupRecords from protobuf object *otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecords) PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	// validate validates PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipReportNumberOfGroupRecords, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoiceEnum, set in PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	Choice() PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoiceEnum
	// setChoice assigns PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoiceEnum provided by user to PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	setChoice(value PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoiceEnum) PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	// HasChoice checks if Choice has been set in PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIgmpv3MembershipReportNumberOfGroupRecords.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	SetValue(value uint32) PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	// HasValue checks if Value has been set in PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIgmpv3MembershipReportNumberOfGroupRecords.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	SetValues(value []uint32) PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	// Auto returns uint32, set in PatternFlowIgmpv3MembershipReportNumberOfGroupRecords.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	HasAuto() bool
	// Increment returns PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter, set in PatternFlowIgmpv3MembershipReportNumberOfGroupRecords.
	// PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter is integer counter pattern
	Increment() PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	// SetIncrement assigns PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter provided by user to PatternFlowIgmpv3MembershipReportNumberOfGroupRecords.
	// PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter is integer counter pattern
	SetIncrement(value PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter, set in PatternFlowIgmpv3MembershipReportNumberOfGroupRecords.
	// PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter is integer counter pattern
	Decrement() PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	// SetDecrement assigns PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter provided by user to PatternFlowIgmpv3MembershipReportNumberOfGroupRecords.
	// PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter is integer counter pattern
	SetDecrement(value PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
	HasDecrement() bool
	setNil()
}

type PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoiceEnum string

// Enum of Choice on PatternFlowIgmpv3MembershipReportNumberOfGroupRecords
var PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoice = struct {
	VALUE     PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoiceEnum
	VALUES    PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoiceEnum
	AUTO      PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoiceEnum
	INCREMENT PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoiceEnum
	DECREMENT PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoiceEnum
}{
	VALUE:     PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoiceEnum("values"),
	AUTO:      PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoiceEnum("auto"),
	INCREMENT: PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) Choice() PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoiceEnum {
	return PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) setChoice(value PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoiceEnum) PatternFlowIgmpv3MembershipReportNumberOfGroupRecords {
	intValue, ok := otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecords_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecords_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoice.AUTO {
		defaultValue := uint32(0)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter().msg()
	}

	if value == PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIgmpv3MembershipReportNumberOfGroupRecords object
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) SetValue(value uint32) PatternFlowIgmpv3MembershipReportNumberOfGroupRecords {
	obj.setChoice(PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIgmpv3MembershipReportNumberOfGroupRecords object
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) SetValues(value []uint32) PatternFlowIgmpv3MembershipReportNumberOfGroupRecords {
	obj.setChoice(PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) Increment() PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter value in the PatternFlowIgmpv3MembershipReportNumberOfGroupRecords object
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) SetIncrement(value PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) PatternFlowIgmpv3MembershipReportNumberOfGroupRecords {
	obj.setChoice(PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) Decrement() PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter value in the PatternFlowIgmpv3MembershipReportNumberOfGroupRecords object
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) SetDecrement(value PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) PatternFlowIgmpv3MembershipReportNumberOfGroupRecords {
	obj.setChoice(PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipReportNumberOfGroupRecords.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIgmpv3MembershipReportNumberOfGroupRecords.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipReportNumberOfGroupRecords.Auto <= 65535 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecords) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv3MembershipReportNumberOfGroupRecords")
			}
		} else {
			intVal := otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecords_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecords_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

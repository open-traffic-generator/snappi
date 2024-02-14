package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid *****
type patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid struct {
	validation
	obj             *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	marshaller      marshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	unMarshaller    unMarshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	incrementHolder PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	decrementHolder PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
}

func NewPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid() PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid {
	obj := patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid{obj: &otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) msg() *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid {
	return obj.obj
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) setMsg(msg *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid struct {
	obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
}

type marshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid interface {
	// ToProto marshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid to protobuf object *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	ToProto() (*otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid, error)
	// ToPbText marshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid struct {
	obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
}

type unMarshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid interface {
	// FromProto unmarshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid from protobuf object *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	FromProto(msg *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) (PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid, error)
	// FromPbText unmarshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) Marshal() marshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) Unmarshal() unMarshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) ToProto() (*otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) FromProto(msg *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) (PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) Clone() (PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid()
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

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid is an identifier of the layer 3 protocol using this path.  Standard Ethertype values are used e.g. The default value of 2048 ( 0x0800 ) represents Ethertype for IPv4.
type PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid interface {
	Validation
	// msg marshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid to protobuf object *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	// setMsg unmarshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid from protobuf object *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	// validate validates PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoiceEnum, set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	Choice() PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoiceEnum
	// setChoice assigns PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoiceEnum provided by user to PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	setChoice(value PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoiceEnum) PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	SetValue(value uint32) PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	// HasValue checks if Value has been set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	SetValues(value []uint32) PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	// Increment returns PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter, set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid.
	Increment() PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	// SetIncrement assigns PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter provided by user to PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid.
	SetIncrement(value PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter, set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid.
	Decrement() PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	// SetDecrement assigns PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter provided by user to PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid.
	SetDecrement(value PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
var PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoice = struct {
	VALUE     PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoiceEnum
	VALUES    PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoiceEnum
	INCREMENT PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoiceEnum
	DECREMENT PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) Choice() PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoiceEnum {
	return PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) setChoice(value PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoiceEnum) PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid {
	intValue, ok := otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoice.VALUE {
		defaultValue := uint32(2048)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoice.VALUES {
		defaultValue := []uint32{2048}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter().msg()
	}

	if value == PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid object
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) SetValue(value uint32) PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid {
	obj.setChoice(PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{2048})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid object
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) SetValues(value []uint32) PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid {
	obj.setChoice(PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) Increment() PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter value in the PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid object
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) SetIncrement(value PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid {
	obj.setChoice(PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) Decrement() PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter value in the PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid object
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) SetDecrement(value PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid {
	obj.setChoice(PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidChoice.VALUE)

	}

}

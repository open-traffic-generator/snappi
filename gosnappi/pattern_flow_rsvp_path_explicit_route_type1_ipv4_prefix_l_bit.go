package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit *****
type patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit struct {
	validation
	obj             *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	marshaller      marshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	unMarshaller    unMarshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	incrementHolder PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	decrementHolder PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
}

func NewPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit() PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit {
	obj := patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit{obj: &otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) msg() *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit {
	return obj.obj
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) setMsg(msg *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit struct {
	obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
}

type marshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit interface {
	// ToProto marshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit to protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	ToProto() (*otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit, error)
	// ToPbText marshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit struct {
	obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
}

type unMarshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit interface {
	// FromProto unmarshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit from protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	FromProto(msg *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) (PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit, error)
	// FromPbText unmarshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) Marshal() marshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) Unmarshal() unMarshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) ToProto() (*otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) FromProto(msg *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) (PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) Clone() (PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit()
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

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit is the L bit is an attribute of the subobject. The L bit is set if the subobject represents a loose hop in the explicit route. If the bit is not set, the subobject represents a strict hop in the explicit route.
type PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit interface {
	Validation
	// msg marshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit to protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	// setMsg unmarshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit from protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	// validate validates PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoiceEnum, set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	Choice() PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoiceEnum
	// setChoice assigns PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoiceEnum provided by user to PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	setChoice(value PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoiceEnum) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	SetValue(value uint32) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	// HasValue checks if Value has been set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	SetValues(value []uint32) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	// Increment returns PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter, set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit.
	// PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter is integer counter pattern
	Increment() PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	// SetIncrement assigns PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter provided by user to PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit.
	// PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter, set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit.
	// PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	// SetDecrement assigns PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter provided by user to PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit.
	// PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
var PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoice = struct {
	VALUE     PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoiceEnum
	VALUES    PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoiceEnum
	INCREMENT PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoiceEnum
	DECREMENT PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) Choice() PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoiceEnum {
	return PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) setChoice(value PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoiceEnum) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit {
	intValue, ok := otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter().msg()
	}

	if value == PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit object
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) SetValue(value uint32) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit {
	obj.setChoice(PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit object
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) SetValues(value []uint32) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit {
	obj.setChoice(PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) Increment() PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter value in the PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit object
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) SetIncrement(value PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit {
	obj.setChoice(PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) Decrement() PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter value in the PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit object
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) SetDecrement(value PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit {
	obj.setChoice(PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit.Values <= 1 but Got %d", item))
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

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitChoice.VALUE)

	}

}

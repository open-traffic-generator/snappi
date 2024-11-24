package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CPDURequestId *****
type patternFlowSnmpv2CPDURequestId struct {
	validation
	obj             *otg.PatternFlowSnmpv2CPDURequestId
	marshaller      marshalPatternFlowSnmpv2CPDURequestId
	unMarshaller    unMarshalPatternFlowSnmpv2CPDURequestId
	incrementHolder PatternFlowSnmpv2CPDURequestIdCounter
	decrementHolder PatternFlowSnmpv2CPDURequestIdCounter
}

func NewPatternFlowSnmpv2CPDURequestId() PatternFlowSnmpv2CPDURequestId {
	obj := patternFlowSnmpv2CPDURequestId{obj: &otg.PatternFlowSnmpv2CPDURequestId{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CPDURequestId) msg() *otg.PatternFlowSnmpv2CPDURequestId {
	return obj.obj
}

func (obj *patternFlowSnmpv2CPDURequestId) setMsg(msg *otg.PatternFlowSnmpv2CPDURequestId) PatternFlowSnmpv2CPDURequestId {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CPDURequestId struct {
	obj *patternFlowSnmpv2CPDURequestId
}

type marshalPatternFlowSnmpv2CPDURequestId interface {
	// ToProto marshals PatternFlowSnmpv2CPDURequestId to protobuf object *otg.PatternFlowSnmpv2CPDURequestId
	ToProto() (*otg.PatternFlowSnmpv2CPDURequestId, error)
	// ToPbText marshals PatternFlowSnmpv2CPDURequestId to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CPDURequestId to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CPDURequestId to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowSnmpv2CPDURequestId to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowSnmpv2CPDURequestId struct {
	obj *patternFlowSnmpv2CPDURequestId
}

type unMarshalPatternFlowSnmpv2CPDURequestId interface {
	// FromProto unmarshals PatternFlowSnmpv2CPDURequestId from protobuf object *otg.PatternFlowSnmpv2CPDURequestId
	FromProto(msg *otg.PatternFlowSnmpv2CPDURequestId) (PatternFlowSnmpv2CPDURequestId, error)
	// FromPbText unmarshals PatternFlowSnmpv2CPDURequestId from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CPDURequestId from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CPDURequestId from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CPDURequestId) Marshal() marshalPatternFlowSnmpv2CPDURequestId {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CPDURequestId{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CPDURequestId) Unmarshal() unMarshalPatternFlowSnmpv2CPDURequestId {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CPDURequestId{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CPDURequestId) ToProto() (*otg.PatternFlowSnmpv2CPDURequestId, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CPDURequestId) FromProto(msg *otg.PatternFlowSnmpv2CPDURequestId) (PatternFlowSnmpv2CPDURequestId, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CPDURequestId) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CPDURequestId) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CPDURequestId) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CPDURequestId) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CPDURequestId) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowSnmpv2CPDURequestId) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CPDURequestId) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CPDURequestId) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CPDURequestId) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CPDURequestId) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CPDURequestId) Clone() (PatternFlowSnmpv2CPDURequestId, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CPDURequestId()
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

func (obj *patternFlowSnmpv2CPDURequestId) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowSnmpv2CPDURequestId is identifies a particular SNMP request.
// This index is echoed back in the response from the SNMP agent,
// allowing the SNMP manager to match an incoming response to the appropriate request.
//
// - Encoding of this field follows ASN.1 X.690(section 8.3) specification.
// Refer: http://www.itu.int/ITU-T/asn1
type PatternFlowSnmpv2CPDURequestId interface {
	Validation
	// msg marshals PatternFlowSnmpv2CPDURequestId to protobuf object *otg.PatternFlowSnmpv2CPDURequestId
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CPDURequestId
	// setMsg unmarshals PatternFlowSnmpv2CPDURequestId from protobuf object *otg.PatternFlowSnmpv2CPDURequestId
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CPDURequestId) PatternFlowSnmpv2CPDURequestId
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CPDURequestId
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CPDURequestId
	// validate validates PatternFlowSnmpv2CPDURequestId
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CPDURequestId, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowSnmpv2CPDURequestIdChoiceEnum, set in PatternFlowSnmpv2CPDURequestId
	Choice() PatternFlowSnmpv2CPDURequestIdChoiceEnum
	// setChoice assigns PatternFlowSnmpv2CPDURequestIdChoiceEnum provided by user to PatternFlowSnmpv2CPDURequestId
	setChoice(value PatternFlowSnmpv2CPDURequestIdChoiceEnum) PatternFlowSnmpv2CPDURequestId
	// HasChoice checks if Choice has been set in PatternFlowSnmpv2CPDURequestId
	HasChoice() bool
	// Value returns int32, set in PatternFlowSnmpv2CPDURequestId.
	Value() int32
	// SetValue assigns int32 provided by user to PatternFlowSnmpv2CPDURequestId
	SetValue(value int32) PatternFlowSnmpv2CPDURequestId
	// HasValue checks if Value has been set in PatternFlowSnmpv2CPDURequestId
	HasValue() bool
	// Values returns []int32, set in PatternFlowSnmpv2CPDURequestId.
	Values() []int32
	// SetValues assigns []int32 provided by user to PatternFlowSnmpv2CPDURequestId
	SetValues(value []int32) PatternFlowSnmpv2CPDURequestId
	// Increment returns PatternFlowSnmpv2CPDURequestIdCounter, set in PatternFlowSnmpv2CPDURequestId.
	Increment() PatternFlowSnmpv2CPDURequestIdCounter
	// SetIncrement assigns PatternFlowSnmpv2CPDURequestIdCounter provided by user to PatternFlowSnmpv2CPDURequestId.
	SetIncrement(value PatternFlowSnmpv2CPDURequestIdCounter) PatternFlowSnmpv2CPDURequestId
	// HasIncrement checks if Increment has been set in PatternFlowSnmpv2CPDURequestId
	HasIncrement() bool
	// Decrement returns PatternFlowSnmpv2CPDURequestIdCounter, set in PatternFlowSnmpv2CPDURequestId.
	Decrement() PatternFlowSnmpv2CPDURequestIdCounter
	// SetDecrement assigns PatternFlowSnmpv2CPDURequestIdCounter provided by user to PatternFlowSnmpv2CPDURequestId.
	SetDecrement(value PatternFlowSnmpv2CPDURequestIdCounter) PatternFlowSnmpv2CPDURequestId
	// HasDecrement checks if Decrement has been set in PatternFlowSnmpv2CPDURequestId
	HasDecrement() bool
	setNil()
}

type PatternFlowSnmpv2CPDURequestIdChoiceEnum string

// Enum of Choice on PatternFlowSnmpv2CPDURequestId
var PatternFlowSnmpv2CPDURequestIdChoice = struct {
	VALUE     PatternFlowSnmpv2CPDURequestIdChoiceEnum
	VALUES    PatternFlowSnmpv2CPDURequestIdChoiceEnum
	INCREMENT PatternFlowSnmpv2CPDURequestIdChoiceEnum
	DECREMENT PatternFlowSnmpv2CPDURequestIdChoiceEnum
}{
	VALUE:     PatternFlowSnmpv2CPDURequestIdChoiceEnum("value"),
	VALUES:    PatternFlowSnmpv2CPDURequestIdChoiceEnum("values"),
	INCREMENT: PatternFlowSnmpv2CPDURequestIdChoiceEnum("increment"),
	DECREMENT: PatternFlowSnmpv2CPDURequestIdChoiceEnum("decrement"),
}

func (obj *patternFlowSnmpv2CPDURequestId) Choice() PatternFlowSnmpv2CPDURequestIdChoiceEnum {
	return PatternFlowSnmpv2CPDURequestIdChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowSnmpv2CPDURequestId) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowSnmpv2CPDURequestId) setChoice(value PatternFlowSnmpv2CPDURequestIdChoiceEnum) PatternFlowSnmpv2CPDURequestId {
	intValue, ok := otg.PatternFlowSnmpv2CPDURequestId_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowSnmpv2CPDURequestIdChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowSnmpv2CPDURequestId_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowSnmpv2CPDURequestIdChoice.VALUE {
		defaultValue := int32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowSnmpv2CPDURequestIdChoice.VALUES {
		defaultValue := []int32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowSnmpv2CPDURequestIdChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowSnmpv2CPDURequestIdCounter().msg()
	}

	if value == PatternFlowSnmpv2CPDURequestIdChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowSnmpv2CPDURequestIdCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a int32
func (obj *patternFlowSnmpv2CPDURequestId) Value() int32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowSnmpv2CPDURequestIdChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a int32
func (obj *patternFlowSnmpv2CPDURequestId) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the int32 value in the PatternFlowSnmpv2CPDURequestId object
func (obj *patternFlowSnmpv2CPDURequestId) SetValue(value int32) PatternFlowSnmpv2CPDURequestId {
	obj.setChoice(PatternFlowSnmpv2CPDURequestIdChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []int32
func (obj *patternFlowSnmpv2CPDURequestId) Values() []int32 {
	if obj.obj.Values == nil {
		obj.SetValues([]int32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []int32 value in the PatternFlowSnmpv2CPDURequestId object
func (obj *patternFlowSnmpv2CPDURequestId) SetValues(value []int32) PatternFlowSnmpv2CPDURequestId {
	obj.setChoice(PatternFlowSnmpv2CPDURequestIdChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]int32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CPDURequestIdCounter
func (obj *patternFlowSnmpv2CPDURequestId) Increment() PatternFlowSnmpv2CPDURequestIdCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowSnmpv2CPDURequestIdChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowSnmpv2CPDURequestIdCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CPDURequestIdCounter
func (obj *patternFlowSnmpv2CPDURequestId) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowSnmpv2CPDURequestIdCounter value in the PatternFlowSnmpv2CPDURequestId object
func (obj *patternFlowSnmpv2CPDURequestId) SetIncrement(value PatternFlowSnmpv2CPDURequestIdCounter) PatternFlowSnmpv2CPDURequestId {
	obj.setChoice(PatternFlowSnmpv2CPDURequestIdChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CPDURequestIdCounter
func (obj *patternFlowSnmpv2CPDURequestId) Decrement() PatternFlowSnmpv2CPDURequestIdCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowSnmpv2CPDURequestIdChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowSnmpv2CPDURequestIdCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CPDURequestIdCounter
func (obj *patternFlowSnmpv2CPDURequestId) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowSnmpv2CPDURequestIdCounter value in the PatternFlowSnmpv2CPDURequestId object
func (obj *patternFlowSnmpv2CPDURequestId) SetDecrement(value PatternFlowSnmpv2CPDURequestIdCounter) PatternFlowSnmpv2CPDURequestId {
	obj.setChoice(PatternFlowSnmpv2CPDURequestIdChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowSnmpv2CPDURequestId) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowSnmpv2CPDURequestId) setDefault() {
	var choices_set int = 0
	var choice PatternFlowSnmpv2CPDURequestIdChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CPDURequestIdChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowSnmpv2CPDURequestIdChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CPDURequestIdChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CPDURequestIdChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowSnmpv2CPDURequestIdChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowSnmpv2CPDURequestId")
			}
		} else {
			intVal := otg.PatternFlowSnmpv2CPDURequestId_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowSnmpv2CPDURequestId_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

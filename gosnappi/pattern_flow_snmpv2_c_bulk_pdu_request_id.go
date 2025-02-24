package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CBulkPDURequestId *****
type patternFlowSnmpv2CBulkPDURequestId struct {
	validation
	obj             *otg.PatternFlowSnmpv2CBulkPDURequestId
	marshaller      marshalPatternFlowSnmpv2CBulkPDURequestId
	unMarshaller    unMarshalPatternFlowSnmpv2CBulkPDURequestId
	incrementHolder PatternFlowSnmpv2CBulkPDURequestIdCounter
	decrementHolder PatternFlowSnmpv2CBulkPDURequestIdCounter
}

func NewPatternFlowSnmpv2CBulkPDURequestId() PatternFlowSnmpv2CBulkPDURequestId {
	obj := patternFlowSnmpv2CBulkPDURequestId{obj: &otg.PatternFlowSnmpv2CBulkPDURequestId{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CBulkPDURequestId) msg() *otg.PatternFlowSnmpv2CBulkPDURequestId {
	return obj.obj
}

func (obj *patternFlowSnmpv2CBulkPDURequestId) setMsg(msg *otg.PatternFlowSnmpv2CBulkPDURequestId) PatternFlowSnmpv2CBulkPDURequestId {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CBulkPDURequestId struct {
	obj *patternFlowSnmpv2CBulkPDURequestId
}

type marshalPatternFlowSnmpv2CBulkPDURequestId interface {
	// ToProto marshals PatternFlowSnmpv2CBulkPDURequestId to protobuf object *otg.PatternFlowSnmpv2CBulkPDURequestId
	ToProto() (*otg.PatternFlowSnmpv2CBulkPDURequestId, error)
	// ToPbText marshals PatternFlowSnmpv2CBulkPDURequestId to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CBulkPDURequestId to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CBulkPDURequestId to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowSnmpv2CBulkPDURequestId to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowSnmpv2CBulkPDURequestId struct {
	obj *patternFlowSnmpv2CBulkPDURequestId
}

type unMarshalPatternFlowSnmpv2CBulkPDURequestId interface {
	// FromProto unmarshals PatternFlowSnmpv2CBulkPDURequestId from protobuf object *otg.PatternFlowSnmpv2CBulkPDURequestId
	FromProto(msg *otg.PatternFlowSnmpv2CBulkPDURequestId) (PatternFlowSnmpv2CBulkPDURequestId, error)
	// FromPbText unmarshals PatternFlowSnmpv2CBulkPDURequestId from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CBulkPDURequestId from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CBulkPDURequestId from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CBulkPDURequestId) Marshal() marshalPatternFlowSnmpv2CBulkPDURequestId {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CBulkPDURequestId{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CBulkPDURequestId) Unmarshal() unMarshalPatternFlowSnmpv2CBulkPDURequestId {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CBulkPDURequestId{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CBulkPDURequestId) ToProto() (*otg.PatternFlowSnmpv2CBulkPDURequestId, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CBulkPDURequestId) FromProto(msg *otg.PatternFlowSnmpv2CBulkPDURequestId) (PatternFlowSnmpv2CBulkPDURequestId, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CBulkPDURequestId) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CBulkPDURequestId) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CBulkPDURequestId) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CBulkPDURequestId) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CBulkPDURequestId) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowSnmpv2CBulkPDURequestId) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CBulkPDURequestId) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CBulkPDURequestId) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CBulkPDURequestId) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CBulkPDURequestId) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CBulkPDURequestId) Clone() (PatternFlowSnmpv2CBulkPDURequestId, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CBulkPDURequestId()
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

func (obj *patternFlowSnmpv2CBulkPDURequestId) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowSnmpv2CBulkPDURequestId is identifies a particular SNMP request.
// This index is echoed back in the response from the SNMP agent,
// allowing the SNMP manager to match an incoming response to the appropriate request.
//
// - Encoding of this field follows ASN.1 X.690(section 8.3) specification.
// Refer: http://www.itu.int/ITU-T/asn1
type PatternFlowSnmpv2CBulkPDURequestId interface {
	Validation
	// msg marshals PatternFlowSnmpv2CBulkPDURequestId to protobuf object *otg.PatternFlowSnmpv2CBulkPDURequestId
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CBulkPDURequestId
	// setMsg unmarshals PatternFlowSnmpv2CBulkPDURequestId from protobuf object *otg.PatternFlowSnmpv2CBulkPDURequestId
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CBulkPDURequestId) PatternFlowSnmpv2CBulkPDURequestId
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CBulkPDURequestId
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CBulkPDURequestId
	// validate validates PatternFlowSnmpv2CBulkPDURequestId
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CBulkPDURequestId, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowSnmpv2CBulkPDURequestIdChoiceEnum, set in PatternFlowSnmpv2CBulkPDURequestId
	Choice() PatternFlowSnmpv2CBulkPDURequestIdChoiceEnum
	// setChoice assigns PatternFlowSnmpv2CBulkPDURequestIdChoiceEnum provided by user to PatternFlowSnmpv2CBulkPDURequestId
	setChoice(value PatternFlowSnmpv2CBulkPDURequestIdChoiceEnum) PatternFlowSnmpv2CBulkPDURequestId
	// HasChoice checks if Choice has been set in PatternFlowSnmpv2CBulkPDURequestId
	HasChoice() bool
	// Value returns int32, set in PatternFlowSnmpv2CBulkPDURequestId.
	Value() int32
	// SetValue assigns int32 provided by user to PatternFlowSnmpv2CBulkPDURequestId
	SetValue(value int32) PatternFlowSnmpv2CBulkPDURequestId
	// HasValue checks if Value has been set in PatternFlowSnmpv2CBulkPDURequestId
	HasValue() bool
	// Values returns []int32, set in PatternFlowSnmpv2CBulkPDURequestId.
	Values() []int32
	// SetValues assigns []int32 provided by user to PatternFlowSnmpv2CBulkPDURequestId
	SetValues(value []int32) PatternFlowSnmpv2CBulkPDURequestId
	// Increment returns PatternFlowSnmpv2CBulkPDURequestIdCounter, set in PatternFlowSnmpv2CBulkPDURequestId.
	Increment() PatternFlowSnmpv2CBulkPDURequestIdCounter
	// SetIncrement assigns PatternFlowSnmpv2CBulkPDURequestIdCounter provided by user to PatternFlowSnmpv2CBulkPDURequestId.
	SetIncrement(value PatternFlowSnmpv2CBulkPDURequestIdCounter) PatternFlowSnmpv2CBulkPDURequestId
	// HasIncrement checks if Increment has been set in PatternFlowSnmpv2CBulkPDURequestId
	HasIncrement() bool
	// Decrement returns PatternFlowSnmpv2CBulkPDURequestIdCounter, set in PatternFlowSnmpv2CBulkPDURequestId.
	Decrement() PatternFlowSnmpv2CBulkPDURequestIdCounter
	// SetDecrement assigns PatternFlowSnmpv2CBulkPDURequestIdCounter provided by user to PatternFlowSnmpv2CBulkPDURequestId.
	SetDecrement(value PatternFlowSnmpv2CBulkPDURequestIdCounter) PatternFlowSnmpv2CBulkPDURequestId
	// HasDecrement checks if Decrement has been set in PatternFlowSnmpv2CBulkPDURequestId
	HasDecrement() bool
	setNil()
}

type PatternFlowSnmpv2CBulkPDURequestIdChoiceEnum string

// Enum of Choice on PatternFlowSnmpv2CBulkPDURequestId
var PatternFlowSnmpv2CBulkPDURequestIdChoice = struct {
	VALUE     PatternFlowSnmpv2CBulkPDURequestIdChoiceEnum
	VALUES    PatternFlowSnmpv2CBulkPDURequestIdChoiceEnum
	INCREMENT PatternFlowSnmpv2CBulkPDURequestIdChoiceEnum
	DECREMENT PatternFlowSnmpv2CBulkPDURequestIdChoiceEnum
}{
	VALUE:     PatternFlowSnmpv2CBulkPDURequestIdChoiceEnum("value"),
	VALUES:    PatternFlowSnmpv2CBulkPDURequestIdChoiceEnum("values"),
	INCREMENT: PatternFlowSnmpv2CBulkPDURequestIdChoiceEnum("increment"),
	DECREMENT: PatternFlowSnmpv2CBulkPDURequestIdChoiceEnum("decrement"),
}

func (obj *patternFlowSnmpv2CBulkPDURequestId) Choice() PatternFlowSnmpv2CBulkPDURequestIdChoiceEnum {
	return PatternFlowSnmpv2CBulkPDURequestIdChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowSnmpv2CBulkPDURequestId) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowSnmpv2CBulkPDURequestId) setChoice(value PatternFlowSnmpv2CBulkPDURequestIdChoiceEnum) PatternFlowSnmpv2CBulkPDURequestId {
	intValue, ok := otg.PatternFlowSnmpv2CBulkPDURequestId_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowSnmpv2CBulkPDURequestIdChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowSnmpv2CBulkPDURequestId_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowSnmpv2CBulkPDURequestIdChoice.VALUE {
		defaultValue := int32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowSnmpv2CBulkPDURequestIdChoice.VALUES {
		defaultValue := []int32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowSnmpv2CBulkPDURequestIdChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowSnmpv2CBulkPDURequestIdCounter().msg()
	}

	if value == PatternFlowSnmpv2CBulkPDURequestIdChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowSnmpv2CBulkPDURequestIdCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a int32
func (obj *patternFlowSnmpv2CBulkPDURequestId) Value() int32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowSnmpv2CBulkPDURequestIdChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a int32
func (obj *patternFlowSnmpv2CBulkPDURequestId) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the int32 value in the PatternFlowSnmpv2CBulkPDURequestId object
func (obj *patternFlowSnmpv2CBulkPDURequestId) SetValue(value int32) PatternFlowSnmpv2CBulkPDURequestId {
	obj.setChoice(PatternFlowSnmpv2CBulkPDURequestIdChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []int32
func (obj *patternFlowSnmpv2CBulkPDURequestId) Values() []int32 {
	if obj.obj.Values == nil {
		obj.SetValues([]int32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []int32 value in the PatternFlowSnmpv2CBulkPDURequestId object
func (obj *patternFlowSnmpv2CBulkPDURequestId) SetValues(value []int32) PatternFlowSnmpv2CBulkPDURequestId {
	obj.setChoice(PatternFlowSnmpv2CBulkPDURequestIdChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]int32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CBulkPDURequestIdCounter
func (obj *patternFlowSnmpv2CBulkPDURequestId) Increment() PatternFlowSnmpv2CBulkPDURequestIdCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowSnmpv2CBulkPDURequestIdChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowSnmpv2CBulkPDURequestIdCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CBulkPDURequestIdCounter
func (obj *patternFlowSnmpv2CBulkPDURequestId) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowSnmpv2CBulkPDURequestIdCounter value in the PatternFlowSnmpv2CBulkPDURequestId object
func (obj *patternFlowSnmpv2CBulkPDURequestId) SetIncrement(value PatternFlowSnmpv2CBulkPDURequestIdCounter) PatternFlowSnmpv2CBulkPDURequestId {
	obj.setChoice(PatternFlowSnmpv2CBulkPDURequestIdChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CBulkPDURequestIdCounter
func (obj *patternFlowSnmpv2CBulkPDURequestId) Decrement() PatternFlowSnmpv2CBulkPDURequestIdCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowSnmpv2CBulkPDURequestIdChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowSnmpv2CBulkPDURequestIdCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CBulkPDURequestIdCounter
func (obj *patternFlowSnmpv2CBulkPDURequestId) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowSnmpv2CBulkPDURequestIdCounter value in the PatternFlowSnmpv2CBulkPDURequestId object
func (obj *patternFlowSnmpv2CBulkPDURequestId) SetDecrement(value PatternFlowSnmpv2CBulkPDURequestIdCounter) PatternFlowSnmpv2CBulkPDURequestId {
	obj.setChoice(PatternFlowSnmpv2CBulkPDURequestIdChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowSnmpv2CBulkPDURequestId) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowSnmpv2CBulkPDURequestId) setDefault() {
	var choices_set int = 0
	var choice PatternFlowSnmpv2CBulkPDURequestIdChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CBulkPDURequestIdChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowSnmpv2CBulkPDURequestIdChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CBulkPDURequestIdChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CBulkPDURequestIdChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowSnmpv2CBulkPDURequestIdChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowSnmpv2CBulkPDURequestId")
			}
		} else {
			intVal := otg.PatternFlowSnmpv2CBulkPDURequestId_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowSnmpv2CBulkPDURequestId_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

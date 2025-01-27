package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpEchoChecksum *****
type patternFlowIcmpEchoChecksum struct {
	validation
	obj          *otg.PatternFlowIcmpEchoChecksum
	marshaller   marshalPatternFlowIcmpEchoChecksum
	unMarshaller unMarshalPatternFlowIcmpEchoChecksum
}

func NewPatternFlowIcmpEchoChecksum() PatternFlowIcmpEchoChecksum {
	obj := patternFlowIcmpEchoChecksum{obj: &otg.PatternFlowIcmpEchoChecksum{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpEchoChecksum) msg() *otg.PatternFlowIcmpEchoChecksum {
	return obj.obj
}

func (obj *patternFlowIcmpEchoChecksum) setMsg(msg *otg.PatternFlowIcmpEchoChecksum) PatternFlowIcmpEchoChecksum {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpEchoChecksum struct {
	obj *patternFlowIcmpEchoChecksum
}

type marshalPatternFlowIcmpEchoChecksum interface {
	// ToProto marshals PatternFlowIcmpEchoChecksum to protobuf object *otg.PatternFlowIcmpEchoChecksum
	ToProto() (*otg.PatternFlowIcmpEchoChecksum, error)
	// ToPbText marshals PatternFlowIcmpEchoChecksum to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpEchoChecksum to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpEchoChecksum to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIcmpEchoChecksum to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIcmpEchoChecksum struct {
	obj *patternFlowIcmpEchoChecksum
}

type unMarshalPatternFlowIcmpEchoChecksum interface {
	// FromProto unmarshals PatternFlowIcmpEchoChecksum from protobuf object *otg.PatternFlowIcmpEchoChecksum
	FromProto(msg *otg.PatternFlowIcmpEchoChecksum) (PatternFlowIcmpEchoChecksum, error)
	// FromPbText unmarshals PatternFlowIcmpEchoChecksum from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpEchoChecksum from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpEchoChecksum from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpEchoChecksum) Marshal() marshalPatternFlowIcmpEchoChecksum {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpEchoChecksum{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpEchoChecksum) Unmarshal() unMarshalPatternFlowIcmpEchoChecksum {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpEchoChecksum{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpEchoChecksum) ToProto() (*otg.PatternFlowIcmpEchoChecksum, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpEchoChecksum) FromProto(msg *otg.PatternFlowIcmpEchoChecksum) (PatternFlowIcmpEchoChecksum, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpEchoChecksum) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoChecksum) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpEchoChecksum) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoChecksum) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpEchoChecksum) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIcmpEchoChecksum) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoChecksum) FromJson(value string) error {
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

func (obj *patternFlowIcmpEchoChecksum) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoChecksum) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoChecksum) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpEchoChecksum) Clone() (PatternFlowIcmpEchoChecksum, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpEchoChecksum()
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

// PatternFlowIcmpEchoChecksum is iCMP checksum
type PatternFlowIcmpEchoChecksum interface {
	Validation
	// msg marshals PatternFlowIcmpEchoChecksum to protobuf object *otg.PatternFlowIcmpEchoChecksum
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpEchoChecksum
	// setMsg unmarshals PatternFlowIcmpEchoChecksum from protobuf object *otg.PatternFlowIcmpEchoChecksum
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpEchoChecksum) PatternFlowIcmpEchoChecksum
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpEchoChecksum
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpEchoChecksum
	// validate validates PatternFlowIcmpEchoChecksum
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpEchoChecksum, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpEchoChecksumChoiceEnum, set in PatternFlowIcmpEchoChecksum
	Choice() PatternFlowIcmpEchoChecksumChoiceEnum
	// setChoice assigns PatternFlowIcmpEchoChecksumChoiceEnum provided by user to PatternFlowIcmpEchoChecksum
	setChoice(value PatternFlowIcmpEchoChecksumChoiceEnum) PatternFlowIcmpEchoChecksum
	// HasChoice checks if Choice has been set in PatternFlowIcmpEchoChecksum
	HasChoice() bool
	// Generated returns PatternFlowIcmpEchoChecksumGeneratedEnum, set in PatternFlowIcmpEchoChecksum
	Generated() PatternFlowIcmpEchoChecksumGeneratedEnum
	// SetGenerated assigns PatternFlowIcmpEchoChecksumGeneratedEnum provided by user to PatternFlowIcmpEchoChecksum
	SetGenerated(value PatternFlowIcmpEchoChecksumGeneratedEnum) PatternFlowIcmpEchoChecksum
	// HasGenerated checks if Generated has been set in PatternFlowIcmpEchoChecksum
	HasGenerated() bool
	// Custom returns uint32, set in PatternFlowIcmpEchoChecksum.
	Custom() uint32
	// SetCustom assigns uint32 provided by user to PatternFlowIcmpEchoChecksum
	SetCustom(value uint32) PatternFlowIcmpEchoChecksum
	// HasCustom checks if Custom has been set in PatternFlowIcmpEchoChecksum
	HasCustom() bool
}

type PatternFlowIcmpEchoChecksumChoiceEnum string

// Enum of Choice on PatternFlowIcmpEchoChecksum
var PatternFlowIcmpEchoChecksumChoice = struct {
	GENERATED PatternFlowIcmpEchoChecksumChoiceEnum
	CUSTOM    PatternFlowIcmpEchoChecksumChoiceEnum
}{
	GENERATED: PatternFlowIcmpEchoChecksumChoiceEnum("generated"),
	CUSTOM:    PatternFlowIcmpEchoChecksumChoiceEnum("custom"),
}

func (obj *patternFlowIcmpEchoChecksum) Choice() PatternFlowIcmpEchoChecksumChoiceEnum {
	return PatternFlowIcmpEchoChecksumChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of checksum
// Choice returns a string
func (obj *patternFlowIcmpEchoChecksum) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpEchoChecksum) setChoice(value PatternFlowIcmpEchoChecksumChoiceEnum) PatternFlowIcmpEchoChecksum {
	intValue, ok := otg.PatternFlowIcmpEchoChecksum_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpEchoChecksumChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpEchoChecksum_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.obj.Generated = otg.PatternFlowIcmpEchoChecksum_Generated_unspecified.Enum()
	return obj
}

type PatternFlowIcmpEchoChecksumGeneratedEnum string

// Enum of Generated on PatternFlowIcmpEchoChecksum
var PatternFlowIcmpEchoChecksumGenerated = struct {
	GOOD PatternFlowIcmpEchoChecksumGeneratedEnum
	BAD  PatternFlowIcmpEchoChecksumGeneratedEnum
}{
	GOOD: PatternFlowIcmpEchoChecksumGeneratedEnum("good"),
	BAD:  PatternFlowIcmpEchoChecksumGeneratedEnum("bad"),
}

func (obj *patternFlowIcmpEchoChecksum) Generated() PatternFlowIcmpEchoChecksumGeneratedEnum {
	return PatternFlowIcmpEchoChecksumGeneratedEnum(obj.obj.Generated.Enum().String())
}

// A system generated checksum value
// Generated returns a string
func (obj *patternFlowIcmpEchoChecksum) HasGenerated() bool {
	return obj.obj.Generated != nil
}

func (obj *patternFlowIcmpEchoChecksum) SetGenerated(value PatternFlowIcmpEchoChecksumGeneratedEnum) PatternFlowIcmpEchoChecksum {
	intValue, ok := otg.PatternFlowIcmpEchoChecksum_Generated_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpEchoChecksumGeneratedEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpEchoChecksum_Generated_Enum(intValue)
	obj.obj.Generated = &enumValue

	return obj
}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowIcmpEchoChecksum) Custom() uint32 {

	if obj.obj.Custom == nil {
		obj.setChoice(PatternFlowIcmpEchoChecksumChoice.CUSTOM)
	}

	return *obj.obj.Custom

}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowIcmpEchoChecksum) HasCustom() bool {
	return obj.obj.Custom != nil
}

// A custom checksum value
// SetCustom sets the uint32 value in the PatternFlowIcmpEchoChecksum object
func (obj *patternFlowIcmpEchoChecksum) SetCustom(value uint32) PatternFlowIcmpEchoChecksum {
	obj.setChoice(PatternFlowIcmpEchoChecksumChoice.CUSTOM)
	obj.obj.Custom = &value
	return obj
}

func (obj *patternFlowIcmpEchoChecksum) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		if *obj.obj.Custom > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoChecksum.Custom <= 65535 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternFlowIcmpEchoChecksum) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpEchoChecksumChoiceEnum

	if obj.obj.Generated != nil && obj.obj.Generated.Number() != 0 {
		choices_set += 1
		choice = PatternFlowIcmpEchoChecksumChoice.GENERATED
	}

	if obj.obj.Custom != nil {
		choices_set += 1
		choice = PatternFlowIcmpEchoChecksumChoice.CUSTOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpEchoChecksumChoice.GENERATED)
			if obj.obj.Generated.Number() == 0 {
				obj.SetGenerated(PatternFlowIcmpEchoChecksumGenerated.GOOD)

			}

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpEchoChecksum")
			}
		} else {
			intVal := otg.PatternFlowIcmpEchoChecksum_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpEchoChecksum_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

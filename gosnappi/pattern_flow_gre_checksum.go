package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGreChecksum *****
type patternFlowGreChecksum struct {
	validation
	obj          *otg.PatternFlowGreChecksum
	marshaller   marshalPatternFlowGreChecksum
	unMarshaller unMarshalPatternFlowGreChecksum
}

func NewPatternFlowGreChecksum() PatternFlowGreChecksum {
	obj := patternFlowGreChecksum{obj: &otg.PatternFlowGreChecksum{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGreChecksum) msg() *otg.PatternFlowGreChecksum {
	return obj.obj
}

func (obj *patternFlowGreChecksum) setMsg(msg *otg.PatternFlowGreChecksum) PatternFlowGreChecksum {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGreChecksum struct {
	obj *patternFlowGreChecksum
}

type marshalPatternFlowGreChecksum interface {
	// ToProto marshals PatternFlowGreChecksum to protobuf object *otg.PatternFlowGreChecksum
	ToProto() (*otg.PatternFlowGreChecksum, error)
	// ToPbText marshals PatternFlowGreChecksum to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGreChecksum to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGreChecksum to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGreChecksum struct {
	obj *patternFlowGreChecksum
}

type unMarshalPatternFlowGreChecksum interface {
	// FromProto unmarshals PatternFlowGreChecksum from protobuf object *otg.PatternFlowGreChecksum
	FromProto(msg *otg.PatternFlowGreChecksum) (PatternFlowGreChecksum, error)
	// FromPbText unmarshals PatternFlowGreChecksum from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGreChecksum from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGreChecksum from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGreChecksum) Marshal() marshalPatternFlowGreChecksum {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGreChecksum{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGreChecksum) Unmarshal() unMarshalPatternFlowGreChecksum {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGreChecksum{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGreChecksum) ToProto() (*otg.PatternFlowGreChecksum, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGreChecksum) FromProto(msg *otg.PatternFlowGreChecksum) (PatternFlowGreChecksum, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGreChecksum) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGreChecksum) FromPbText(value string) error {
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

func (m *marshalpatternFlowGreChecksum) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGreChecksum) FromYaml(value string) error {
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

func (m *marshalpatternFlowGreChecksum) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGreChecksum) FromJson(value string) error {
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

func (obj *patternFlowGreChecksum) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGreChecksum) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGreChecksum) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGreChecksum) Clone() (PatternFlowGreChecksum, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGreChecksum()
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

// PatternFlowGreChecksum is optional checksum of GRE header and payload. Only present if the checksum_present bit is set.
type PatternFlowGreChecksum interface {
	Validation
	// msg marshals PatternFlowGreChecksum to protobuf object *otg.PatternFlowGreChecksum
	// and doesn't set defaults
	msg() *otg.PatternFlowGreChecksum
	// setMsg unmarshals PatternFlowGreChecksum from protobuf object *otg.PatternFlowGreChecksum
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGreChecksum) PatternFlowGreChecksum
	// provides marshal interface
	Marshal() marshalPatternFlowGreChecksum
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGreChecksum
	// validate validates PatternFlowGreChecksum
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGreChecksum, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGreChecksumChoiceEnum, set in PatternFlowGreChecksum
	Choice() PatternFlowGreChecksumChoiceEnum
	// setChoice assigns PatternFlowGreChecksumChoiceEnum provided by user to PatternFlowGreChecksum
	setChoice(value PatternFlowGreChecksumChoiceEnum) PatternFlowGreChecksum
	// HasChoice checks if Choice has been set in PatternFlowGreChecksum
	HasChoice() bool
	// Generated returns PatternFlowGreChecksumGeneratedEnum, set in PatternFlowGreChecksum
	Generated() PatternFlowGreChecksumGeneratedEnum
	// SetGenerated assigns PatternFlowGreChecksumGeneratedEnum provided by user to PatternFlowGreChecksum
	SetGenerated(value PatternFlowGreChecksumGeneratedEnum) PatternFlowGreChecksum
	// HasGenerated checks if Generated has been set in PatternFlowGreChecksum
	HasGenerated() bool
	// Custom returns uint32, set in PatternFlowGreChecksum.
	Custom() uint32
	// SetCustom assigns uint32 provided by user to PatternFlowGreChecksum
	SetCustom(value uint32) PatternFlowGreChecksum
	// HasCustom checks if Custom has been set in PatternFlowGreChecksum
	HasCustom() bool
}

type PatternFlowGreChecksumChoiceEnum string

// Enum of Choice on PatternFlowGreChecksum
var PatternFlowGreChecksumChoice = struct {
	GENERATED PatternFlowGreChecksumChoiceEnum
	CUSTOM    PatternFlowGreChecksumChoiceEnum
}{
	GENERATED: PatternFlowGreChecksumChoiceEnum("generated"),
	CUSTOM:    PatternFlowGreChecksumChoiceEnum("custom"),
}

func (obj *patternFlowGreChecksum) Choice() PatternFlowGreChecksumChoiceEnum {
	return PatternFlowGreChecksumChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of checksum
// Choice returns a string
func (obj *patternFlowGreChecksum) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGreChecksum) setChoice(value PatternFlowGreChecksumChoiceEnum) PatternFlowGreChecksum {
	intValue, ok := otg.PatternFlowGreChecksum_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGreChecksumChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGreChecksum_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.obj.Generated = otg.PatternFlowGreChecksum_Generated_unspecified.Enum()
	return obj
}

type PatternFlowGreChecksumGeneratedEnum string

// Enum of Generated on PatternFlowGreChecksum
var PatternFlowGreChecksumGenerated = struct {
	GOOD PatternFlowGreChecksumGeneratedEnum
	BAD  PatternFlowGreChecksumGeneratedEnum
}{
	GOOD: PatternFlowGreChecksumGeneratedEnum("good"),
	BAD:  PatternFlowGreChecksumGeneratedEnum("bad"),
}

func (obj *patternFlowGreChecksum) Generated() PatternFlowGreChecksumGeneratedEnum {
	return PatternFlowGreChecksumGeneratedEnum(obj.obj.Generated.Enum().String())
}

// A system generated checksum value
// Generated returns a string
func (obj *patternFlowGreChecksum) HasGenerated() bool {
	return obj.obj.Generated != nil
}

func (obj *patternFlowGreChecksum) SetGenerated(value PatternFlowGreChecksumGeneratedEnum) PatternFlowGreChecksum {
	intValue, ok := otg.PatternFlowGreChecksum_Generated_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGreChecksumGeneratedEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGreChecksum_Generated_Enum(intValue)
	obj.obj.Generated = &enumValue

	return obj
}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowGreChecksum) Custom() uint32 {

	if obj.obj.Custom == nil {
		obj.setChoice(PatternFlowGreChecksumChoice.CUSTOM)
	}

	return *obj.obj.Custom

}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowGreChecksum) HasCustom() bool {
	return obj.obj.Custom != nil
}

// A custom checksum value
// SetCustom sets the uint32 value in the PatternFlowGreChecksum object
func (obj *patternFlowGreChecksum) SetCustom(value uint32) PatternFlowGreChecksum {
	obj.setChoice(PatternFlowGreChecksumChoice.CUSTOM)
	obj.obj.Custom = &value
	return obj
}

func (obj *patternFlowGreChecksum) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		if *obj.obj.Custom > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreChecksum.Custom <= 65535 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternFlowGreChecksum) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGreChecksumChoiceEnum

	if obj.obj.Generated != nil && obj.obj.Generated.Number() != 0 {
		choices_set += 1
		choice = PatternFlowGreChecksumChoice.GENERATED
	}

	if obj.obj.Custom != nil {
		choices_set += 1
		choice = PatternFlowGreChecksumChoice.CUSTOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGreChecksumChoice.GENERATED)
			if obj.obj.Generated.Number() == 0 {
				obj.SetGenerated(PatternFlowGreChecksumGenerated.GOOD)

			}

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGreChecksum")
			}
		} else {
			intVal := otg.PatternFlowGreChecksum_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGreChecksum_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

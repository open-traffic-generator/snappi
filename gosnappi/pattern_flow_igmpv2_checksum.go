package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv2Checksum *****
type patternFlowIgmpv2Checksum struct {
	validation
	obj          *otg.PatternFlowIgmpv2Checksum
	marshaller   marshalPatternFlowIgmpv2Checksum
	unMarshaller unMarshalPatternFlowIgmpv2Checksum
}

func NewPatternFlowIgmpv2Checksum() PatternFlowIgmpv2Checksum {
	obj := patternFlowIgmpv2Checksum{obj: &otg.PatternFlowIgmpv2Checksum{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv2Checksum) msg() *otg.PatternFlowIgmpv2Checksum {
	return obj.obj
}

func (obj *patternFlowIgmpv2Checksum) setMsg(msg *otg.PatternFlowIgmpv2Checksum) PatternFlowIgmpv2Checksum {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv2Checksum struct {
	obj *patternFlowIgmpv2Checksum
}

type marshalPatternFlowIgmpv2Checksum interface {
	// ToProto marshals PatternFlowIgmpv2Checksum to protobuf object *otg.PatternFlowIgmpv2Checksum
	ToProto() (*otg.PatternFlowIgmpv2Checksum, error)
	// ToPbText marshals PatternFlowIgmpv2Checksum to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv2Checksum to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv2Checksum to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv2Checksum struct {
	obj *patternFlowIgmpv2Checksum
}

type unMarshalPatternFlowIgmpv2Checksum interface {
	// FromProto unmarshals PatternFlowIgmpv2Checksum from protobuf object *otg.PatternFlowIgmpv2Checksum
	FromProto(msg *otg.PatternFlowIgmpv2Checksum) (PatternFlowIgmpv2Checksum, error)
	// FromPbText unmarshals PatternFlowIgmpv2Checksum from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv2Checksum from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv2Checksum from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv2Checksum) Marshal() marshalPatternFlowIgmpv2Checksum {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv2Checksum{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv2Checksum) Unmarshal() unMarshalPatternFlowIgmpv2Checksum {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv2Checksum{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv2Checksum) ToProto() (*otg.PatternFlowIgmpv2Checksum, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv2Checksum) FromProto(msg *otg.PatternFlowIgmpv2Checksum) (PatternFlowIgmpv2Checksum, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv2Checksum) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv2Checksum) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv2Checksum) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv2Checksum) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv2Checksum) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv2Checksum) FromJson(value string) error {
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

func (obj *patternFlowIgmpv2Checksum) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv2Checksum) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv2Checksum) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv2Checksum) Clone() (PatternFlowIgmpv2Checksum, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv2Checksum()
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

// PatternFlowIgmpv2Checksum is checksum
type PatternFlowIgmpv2Checksum interface {
	Validation
	// msg marshals PatternFlowIgmpv2Checksum to protobuf object *otg.PatternFlowIgmpv2Checksum
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv2Checksum
	// setMsg unmarshals PatternFlowIgmpv2Checksum from protobuf object *otg.PatternFlowIgmpv2Checksum
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv2Checksum) PatternFlowIgmpv2Checksum
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv2Checksum
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv2Checksum
	// validate validates PatternFlowIgmpv2Checksum
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv2Checksum, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv2ChecksumChoiceEnum, set in PatternFlowIgmpv2Checksum
	Choice() PatternFlowIgmpv2ChecksumChoiceEnum
	// setChoice assigns PatternFlowIgmpv2ChecksumChoiceEnum provided by user to PatternFlowIgmpv2Checksum
	setChoice(value PatternFlowIgmpv2ChecksumChoiceEnum) PatternFlowIgmpv2Checksum
	// HasChoice checks if Choice has been set in PatternFlowIgmpv2Checksum
	HasChoice() bool
	// Generated returns PatternFlowIgmpv2ChecksumGeneratedEnum, set in PatternFlowIgmpv2Checksum
	Generated() PatternFlowIgmpv2ChecksumGeneratedEnum
	// SetGenerated assigns PatternFlowIgmpv2ChecksumGeneratedEnum provided by user to PatternFlowIgmpv2Checksum
	SetGenerated(value PatternFlowIgmpv2ChecksumGeneratedEnum) PatternFlowIgmpv2Checksum
	// HasGenerated checks if Generated has been set in PatternFlowIgmpv2Checksum
	HasGenerated() bool
	// Custom returns uint32, set in PatternFlowIgmpv2Checksum.
	Custom() uint32
	// SetCustom assigns uint32 provided by user to PatternFlowIgmpv2Checksum
	SetCustom(value uint32) PatternFlowIgmpv2Checksum
	// HasCustom checks if Custom has been set in PatternFlowIgmpv2Checksum
	HasCustom() bool
}

type PatternFlowIgmpv2ChecksumChoiceEnum string

// Enum of Choice on PatternFlowIgmpv2Checksum
var PatternFlowIgmpv2ChecksumChoice = struct {
	GENERATED PatternFlowIgmpv2ChecksumChoiceEnum
	CUSTOM    PatternFlowIgmpv2ChecksumChoiceEnum
}{
	GENERATED: PatternFlowIgmpv2ChecksumChoiceEnum("generated"),
	CUSTOM:    PatternFlowIgmpv2ChecksumChoiceEnum("custom"),
}

func (obj *patternFlowIgmpv2Checksum) Choice() PatternFlowIgmpv2ChecksumChoiceEnum {
	return PatternFlowIgmpv2ChecksumChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of checksum
// Choice returns a string
func (obj *patternFlowIgmpv2Checksum) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv2Checksum) setChoice(value PatternFlowIgmpv2ChecksumChoiceEnum) PatternFlowIgmpv2Checksum {
	intValue, ok := otg.PatternFlowIgmpv2Checksum_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv2ChecksumChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv2Checksum_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.obj.Generated = otg.PatternFlowIgmpv2Checksum_Generated_unspecified.Enum()
	return obj
}

type PatternFlowIgmpv2ChecksumGeneratedEnum string

// Enum of Generated on PatternFlowIgmpv2Checksum
var PatternFlowIgmpv2ChecksumGenerated = struct {
	GOOD PatternFlowIgmpv2ChecksumGeneratedEnum
	BAD  PatternFlowIgmpv2ChecksumGeneratedEnum
}{
	GOOD: PatternFlowIgmpv2ChecksumGeneratedEnum("good"),
	BAD:  PatternFlowIgmpv2ChecksumGeneratedEnum("bad"),
}

func (obj *patternFlowIgmpv2Checksum) Generated() PatternFlowIgmpv2ChecksumGeneratedEnum {
	return PatternFlowIgmpv2ChecksumGeneratedEnum(obj.obj.Generated.Enum().String())
}

// A system generated checksum value
// Generated returns a string
func (obj *patternFlowIgmpv2Checksum) HasGenerated() bool {
	return obj.obj.Generated != nil
}

func (obj *patternFlowIgmpv2Checksum) SetGenerated(value PatternFlowIgmpv2ChecksumGeneratedEnum) PatternFlowIgmpv2Checksum {
	intValue, ok := otg.PatternFlowIgmpv2Checksum_Generated_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv2ChecksumGeneratedEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv2Checksum_Generated_Enum(intValue)
	obj.obj.Generated = &enumValue

	return obj
}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowIgmpv2Checksum) Custom() uint32 {

	if obj.obj.Custom == nil {
		obj.setChoice(PatternFlowIgmpv2ChecksumChoice.CUSTOM)
	}

	return *obj.obj.Custom

}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowIgmpv2Checksum) HasCustom() bool {
	return obj.obj.Custom != nil
}

// A custom checksum value
// SetCustom sets the uint32 value in the PatternFlowIgmpv2Checksum object
func (obj *patternFlowIgmpv2Checksum) SetCustom(value uint32) PatternFlowIgmpv2Checksum {
	obj.setChoice(PatternFlowIgmpv2ChecksumChoice.CUSTOM)
	obj.obj.Custom = &value
	return obj
}

func (obj *patternFlowIgmpv2Checksum) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		if *obj.obj.Custom > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv2Checksum.Custom <= 65535 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternFlowIgmpv2Checksum) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv2ChecksumChoiceEnum

	if obj.obj.Generated != nil && obj.obj.Generated.Number() != 0 {
		choices_set += 1
		choice = PatternFlowIgmpv2ChecksumChoice.GENERATED
	}

	if obj.obj.Custom != nil {
		choices_set += 1
		choice = PatternFlowIgmpv2ChecksumChoice.CUSTOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv2ChecksumChoice.GENERATED)
			if obj.obj.Generated.Number() == 0 {
				obj.SetGenerated(PatternFlowIgmpv2ChecksumGenerated.GOOD)

			}

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv2Checksum")
			}
		} else {
			intVal := otg.PatternFlowIgmpv2Checksum_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv2Checksum_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

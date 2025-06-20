package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv1Checksum *****
type patternFlowIgmpv1Checksum struct {
	validation
	obj          *otg.PatternFlowIgmpv1Checksum
	marshaller   marshalPatternFlowIgmpv1Checksum
	unMarshaller unMarshalPatternFlowIgmpv1Checksum
}

func NewPatternFlowIgmpv1Checksum() PatternFlowIgmpv1Checksum {
	obj := patternFlowIgmpv1Checksum{obj: &otg.PatternFlowIgmpv1Checksum{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv1Checksum) msg() *otg.PatternFlowIgmpv1Checksum {
	return obj.obj
}

func (obj *patternFlowIgmpv1Checksum) setMsg(msg *otg.PatternFlowIgmpv1Checksum) PatternFlowIgmpv1Checksum {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv1Checksum struct {
	obj *patternFlowIgmpv1Checksum
}

type marshalPatternFlowIgmpv1Checksum interface {
	// ToProto marshals PatternFlowIgmpv1Checksum to protobuf object *otg.PatternFlowIgmpv1Checksum
	ToProto() (*otg.PatternFlowIgmpv1Checksum, error)
	// ToPbText marshals PatternFlowIgmpv1Checksum to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv1Checksum to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv1Checksum to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv1Checksum struct {
	obj *patternFlowIgmpv1Checksum
}

type unMarshalPatternFlowIgmpv1Checksum interface {
	// FromProto unmarshals PatternFlowIgmpv1Checksum from protobuf object *otg.PatternFlowIgmpv1Checksum
	FromProto(msg *otg.PatternFlowIgmpv1Checksum) (PatternFlowIgmpv1Checksum, error)
	// FromPbText unmarshals PatternFlowIgmpv1Checksum from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv1Checksum from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv1Checksum from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv1Checksum) Marshal() marshalPatternFlowIgmpv1Checksum {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv1Checksum{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv1Checksum) Unmarshal() unMarshalPatternFlowIgmpv1Checksum {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv1Checksum{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv1Checksum) ToProto() (*otg.PatternFlowIgmpv1Checksum, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv1Checksum) FromProto(msg *otg.PatternFlowIgmpv1Checksum) (PatternFlowIgmpv1Checksum, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv1Checksum) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1Checksum) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv1Checksum) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1Checksum) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv1Checksum) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1Checksum) FromJson(value string) error {
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

func (obj *patternFlowIgmpv1Checksum) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1Checksum) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1Checksum) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv1Checksum) Clone() (PatternFlowIgmpv1Checksum, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv1Checksum()
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

// PatternFlowIgmpv1Checksum is checksum
type PatternFlowIgmpv1Checksum interface {
	Validation
	// msg marshals PatternFlowIgmpv1Checksum to protobuf object *otg.PatternFlowIgmpv1Checksum
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv1Checksum
	// setMsg unmarshals PatternFlowIgmpv1Checksum from protobuf object *otg.PatternFlowIgmpv1Checksum
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv1Checksum) PatternFlowIgmpv1Checksum
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv1Checksum
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv1Checksum
	// validate validates PatternFlowIgmpv1Checksum
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv1Checksum, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv1ChecksumChoiceEnum, set in PatternFlowIgmpv1Checksum
	Choice() PatternFlowIgmpv1ChecksumChoiceEnum
	// setChoice assigns PatternFlowIgmpv1ChecksumChoiceEnum provided by user to PatternFlowIgmpv1Checksum
	setChoice(value PatternFlowIgmpv1ChecksumChoiceEnum) PatternFlowIgmpv1Checksum
	// HasChoice checks if Choice has been set in PatternFlowIgmpv1Checksum
	HasChoice() bool
	// Generated returns PatternFlowIgmpv1ChecksumGeneratedEnum, set in PatternFlowIgmpv1Checksum
	Generated() PatternFlowIgmpv1ChecksumGeneratedEnum
	// SetGenerated assigns PatternFlowIgmpv1ChecksumGeneratedEnum provided by user to PatternFlowIgmpv1Checksum
	SetGenerated(value PatternFlowIgmpv1ChecksumGeneratedEnum) PatternFlowIgmpv1Checksum
	// HasGenerated checks if Generated has been set in PatternFlowIgmpv1Checksum
	HasGenerated() bool
	// Custom returns uint32, set in PatternFlowIgmpv1Checksum.
	Custom() uint32
	// SetCustom assigns uint32 provided by user to PatternFlowIgmpv1Checksum
	SetCustom(value uint32) PatternFlowIgmpv1Checksum
	// HasCustom checks if Custom has been set in PatternFlowIgmpv1Checksum
	HasCustom() bool
}

type PatternFlowIgmpv1ChecksumChoiceEnum string

// Enum of Choice on PatternFlowIgmpv1Checksum
var PatternFlowIgmpv1ChecksumChoice = struct {
	GENERATED PatternFlowIgmpv1ChecksumChoiceEnum
	CUSTOM    PatternFlowIgmpv1ChecksumChoiceEnum
}{
	GENERATED: PatternFlowIgmpv1ChecksumChoiceEnum("generated"),
	CUSTOM:    PatternFlowIgmpv1ChecksumChoiceEnum("custom"),
}

func (obj *patternFlowIgmpv1Checksum) Choice() PatternFlowIgmpv1ChecksumChoiceEnum {
	return PatternFlowIgmpv1ChecksumChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of checksum
// Choice returns a string
func (obj *patternFlowIgmpv1Checksum) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv1Checksum) setChoice(value PatternFlowIgmpv1ChecksumChoiceEnum) PatternFlowIgmpv1Checksum {
	intValue, ok := otg.PatternFlowIgmpv1Checksum_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv1ChecksumChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv1Checksum_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.obj.Generated = otg.PatternFlowIgmpv1Checksum_Generated_unspecified.Enum()
	return obj
}

type PatternFlowIgmpv1ChecksumGeneratedEnum string

// Enum of Generated on PatternFlowIgmpv1Checksum
var PatternFlowIgmpv1ChecksumGenerated = struct {
	GOOD PatternFlowIgmpv1ChecksumGeneratedEnum
	BAD  PatternFlowIgmpv1ChecksumGeneratedEnum
}{
	GOOD: PatternFlowIgmpv1ChecksumGeneratedEnum("good"),
	BAD:  PatternFlowIgmpv1ChecksumGeneratedEnum("bad"),
}

func (obj *patternFlowIgmpv1Checksum) Generated() PatternFlowIgmpv1ChecksumGeneratedEnum {
	return PatternFlowIgmpv1ChecksumGeneratedEnum(obj.obj.Generated.Enum().String())
}

// A system generated checksum value
// Generated returns a string
func (obj *patternFlowIgmpv1Checksum) HasGenerated() bool {
	return obj.obj.Generated != nil
}

func (obj *patternFlowIgmpv1Checksum) SetGenerated(value PatternFlowIgmpv1ChecksumGeneratedEnum) PatternFlowIgmpv1Checksum {
	intValue, ok := otg.PatternFlowIgmpv1Checksum_Generated_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv1ChecksumGeneratedEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv1Checksum_Generated_Enum(intValue)
	obj.obj.Generated = &enumValue

	return obj
}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowIgmpv1Checksum) Custom() uint32 {

	if obj.obj.Custom == nil {
		obj.setChoice(PatternFlowIgmpv1ChecksumChoice.CUSTOM)
	}

	return *obj.obj.Custom

}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowIgmpv1Checksum) HasCustom() bool {
	return obj.obj.Custom != nil
}

// A custom checksum value
// SetCustom sets the uint32 value in the PatternFlowIgmpv1Checksum object
func (obj *patternFlowIgmpv1Checksum) SetCustom(value uint32) PatternFlowIgmpv1Checksum {
	obj.setChoice(PatternFlowIgmpv1ChecksumChoice.CUSTOM)
	obj.obj.Custom = &value
	return obj
}

func (obj *patternFlowIgmpv1Checksum) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		if *obj.obj.Custom > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1Checksum.Custom <= 65535 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternFlowIgmpv1Checksum) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv1ChecksumChoiceEnum

	if obj.obj.Generated != nil && obj.obj.Generated.Number() != 0 {
		choices_set += 1
		choice = PatternFlowIgmpv1ChecksumChoice.GENERATED
	}

	if obj.obj.Custom != nil {
		choices_set += 1
		choice = PatternFlowIgmpv1ChecksumChoice.CUSTOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv1ChecksumChoice.GENERATED)
			if obj.obj.Generated.Number() == 0 {
				obj.SetGenerated(PatternFlowIgmpv1ChecksumGenerated.GOOD)

			}

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv1Checksum")
			}
		} else {
			intVal := otg.PatternFlowIgmpv1Checksum_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv1Checksum_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

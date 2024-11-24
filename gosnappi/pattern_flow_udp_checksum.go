package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowUdpChecksum *****
type patternFlowUdpChecksum struct {
	validation
	obj          *otg.PatternFlowUdpChecksum
	marshaller   marshalPatternFlowUdpChecksum
	unMarshaller unMarshalPatternFlowUdpChecksum
}

func NewPatternFlowUdpChecksum() PatternFlowUdpChecksum {
	obj := patternFlowUdpChecksum{obj: &otg.PatternFlowUdpChecksum{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowUdpChecksum) msg() *otg.PatternFlowUdpChecksum {
	return obj.obj
}

func (obj *patternFlowUdpChecksum) setMsg(msg *otg.PatternFlowUdpChecksum) PatternFlowUdpChecksum {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowUdpChecksum struct {
	obj *patternFlowUdpChecksum
}

type marshalPatternFlowUdpChecksum interface {
	// ToProto marshals PatternFlowUdpChecksum to protobuf object *otg.PatternFlowUdpChecksum
	ToProto() (*otg.PatternFlowUdpChecksum, error)
	// ToPbText marshals PatternFlowUdpChecksum to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowUdpChecksum to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowUdpChecksum to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowUdpChecksum to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowUdpChecksum struct {
	obj *patternFlowUdpChecksum
}

type unMarshalPatternFlowUdpChecksum interface {
	// FromProto unmarshals PatternFlowUdpChecksum from protobuf object *otg.PatternFlowUdpChecksum
	FromProto(msg *otg.PatternFlowUdpChecksum) (PatternFlowUdpChecksum, error)
	// FromPbText unmarshals PatternFlowUdpChecksum from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowUdpChecksum from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowUdpChecksum from JSON text
	FromJson(value string) error
}

func (obj *patternFlowUdpChecksum) Marshal() marshalPatternFlowUdpChecksum {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowUdpChecksum{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowUdpChecksum) Unmarshal() unMarshalPatternFlowUdpChecksum {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowUdpChecksum{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowUdpChecksum) ToProto() (*otg.PatternFlowUdpChecksum, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowUdpChecksum) FromProto(msg *otg.PatternFlowUdpChecksum) (PatternFlowUdpChecksum, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowUdpChecksum) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowUdpChecksum) FromPbText(value string) error {
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

func (m *marshalpatternFlowUdpChecksum) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowUdpChecksum) FromYaml(value string) error {
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

func (m *marshalpatternFlowUdpChecksum) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowUdpChecksum) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowUdpChecksum) FromJson(value string) error {
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

func (obj *patternFlowUdpChecksum) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowUdpChecksum) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowUdpChecksum) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowUdpChecksum) Clone() (PatternFlowUdpChecksum, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowUdpChecksum()
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

// PatternFlowUdpChecksum is uDP checksum
type PatternFlowUdpChecksum interface {
	Validation
	// msg marshals PatternFlowUdpChecksum to protobuf object *otg.PatternFlowUdpChecksum
	// and doesn't set defaults
	msg() *otg.PatternFlowUdpChecksum
	// setMsg unmarshals PatternFlowUdpChecksum from protobuf object *otg.PatternFlowUdpChecksum
	// and doesn't set defaults
	setMsg(*otg.PatternFlowUdpChecksum) PatternFlowUdpChecksum
	// provides marshal interface
	Marshal() marshalPatternFlowUdpChecksum
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowUdpChecksum
	// validate validates PatternFlowUdpChecksum
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowUdpChecksum, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowUdpChecksumChoiceEnum, set in PatternFlowUdpChecksum
	Choice() PatternFlowUdpChecksumChoiceEnum
	// setChoice assigns PatternFlowUdpChecksumChoiceEnum provided by user to PatternFlowUdpChecksum
	setChoice(value PatternFlowUdpChecksumChoiceEnum) PatternFlowUdpChecksum
	// HasChoice checks if Choice has been set in PatternFlowUdpChecksum
	HasChoice() bool
	// Generated returns PatternFlowUdpChecksumGeneratedEnum, set in PatternFlowUdpChecksum
	Generated() PatternFlowUdpChecksumGeneratedEnum
	// SetGenerated assigns PatternFlowUdpChecksumGeneratedEnum provided by user to PatternFlowUdpChecksum
	SetGenerated(value PatternFlowUdpChecksumGeneratedEnum) PatternFlowUdpChecksum
	// HasGenerated checks if Generated has been set in PatternFlowUdpChecksum
	HasGenerated() bool
	// Custom returns uint32, set in PatternFlowUdpChecksum.
	Custom() uint32
	// SetCustom assigns uint32 provided by user to PatternFlowUdpChecksum
	SetCustom(value uint32) PatternFlowUdpChecksum
	// HasCustom checks if Custom has been set in PatternFlowUdpChecksum
	HasCustom() bool
}

type PatternFlowUdpChecksumChoiceEnum string

// Enum of Choice on PatternFlowUdpChecksum
var PatternFlowUdpChecksumChoice = struct {
	GENERATED PatternFlowUdpChecksumChoiceEnum
	CUSTOM    PatternFlowUdpChecksumChoiceEnum
}{
	GENERATED: PatternFlowUdpChecksumChoiceEnum("generated"),
	CUSTOM:    PatternFlowUdpChecksumChoiceEnum("custom"),
}

func (obj *patternFlowUdpChecksum) Choice() PatternFlowUdpChecksumChoiceEnum {
	return PatternFlowUdpChecksumChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of checksum
// Choice returns a string
func (obj *patternFlowUdpChecksum) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowUdpChecksum) setChoice(value PatternFlowUdpChecksumChoiceEnum) PatternFlowUdpChecksum {
	intValue, ok := otg.PatternFlowUdpChecksum_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowUdpChecksumChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowUdpChecksum_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.obj.Generated = otg.PatternFlowUdpChecksum_Generated_unspecified.Enum()
	return obj
}

type PatternFlowUdpChecksumGeneratedEnum string

// Enum of Generated on PatternFlowUdpChecksum
var PatternFlowUdpChecksumGenerated = struct {
	GOOD PatternFlowUdpChecksumGeneratedEnum
	BAD  PatternFlowUdpChecksumGeneratedEnum
}{
	GOOD: PatternFlowUdpChecksumGeneratedEnum("good"),
	BAD:  PatternFlowUdpChecksumGeneratedEnum("bad"),
}

func (obj *patternFlowUdpChecksum) Generated() PatternFlowUdpChecksumGeneratedEnum {
	return PatternFlowUdpChecksumGeneratedEnum(obj.obj.Generated.Enum().String())
}

// A system generated checksum value
// Generated returns a string
func (obj *patternFlowUdpChecksum) HasGenerated() bool {
	return obj.obj.Generated != nil
}

func (obj *patternFlowUdpChecksum) SetGenerated(value PatternFlowUdpChecksumGeneratedEnum) PatternFlowUdpChecksum {
	intValue, ok := otg.PatternFlowUdpChecksum_Generated_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowUdpChecksumGeneratedEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowUdpChecksum_Generated_Enum(intValue)
	obj.obj.Generated = &enumValue

	return obj
}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowUdpChecksum) Custom() uint32 {

	if obj.obj.Custom == nil {
		obj.setChoice(PatternFlowUdpChecksumChoice.CUSTOM)
	}

	return *obj.obj.Custom

}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowUdpChecksum) HasCustom() bool {
	return obj.obj.Custom != nil
}

// A custom checksum value
// SetCustom sets the uint32 value in the PatternFlowUdpChecksum object
func (obj *patternFlowUdpChecksum) SetCustom(value uint32) PatternFlowUdpChecksum {
	obj.setChoice(PatternFlowUdpChecksumChoice.CUSTOM)
	obj.obj.Custom = &value
	return obj
}

func (obj *patternFlowUdpChecksum) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		if *obj.obj.Custom > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowUdpChecksum.Custom <= 65535 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternFlowUdpChecksum) setDefault() {
	var choices_set int = 0
	var choice PatternFlowUdpChecksumChoiceEnum

	if obj.obj.Generated != nil && obj.obj.Generated.Number() != 0 {
		choices_set += 1
		choice = PatternFlowUdpChecksumChoice.GENERATED
	}

	if obj.obj.Custom != nil {
		choices_set += 1
		choice = PatternFlowUdpChecksumChoice.CUSTOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowUdpChecksumChoice.GENERATED)
			if obj.obj.Generated.Number() == 0 {
				obj.SetGenerated(PatternFlowUdpChecksumGenerated.GOOD)

			}

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowUdpChecksum")
			}
		} else {
			intVal := otg.PatternFlowUdpChecksum_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowUdpChecksum_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

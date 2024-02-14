package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRsvpRsvpChecksum *****
type patternFlowRsvpRsvpChecksum struct {
	validation
	obj          *otg.PatternFlowRsvpRsvpChecksum
	marshaller   marshalPatternFlowRsvpRsvpChecksum
	unMarshaller unMarshalPatternFlowRsvpRsvpChecksum
}

func NewPatternFlowRsvpRsvpChecksum() PatternFlowRsvpRsvpChecksum {
	obj := patternFlowRsvpRsvpChecksum{obj: &otg.PatternFlowRsvpRsvpChecksum{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRsvpRsvpChecksum) msg() *otg.PatternFlowRsvpRsvpChecksum {
	return obj.obj
}

func (obj *patternFlowRsvpRsvpChecksum) setMsg(msg *otg.PatternFlowRsvpRsvpChecksum) PatternFlowRsvpRsvpChecksum {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRsvpRsvpChecksum struct {
	obj *patternFlowRsvpRsvpChecksum
}

type marshalPatternFlowRsvpRsvpChecksum interface {
	// ToProto marshals PatternFlowRsvpRsvpChecksum to protobuf object *otg.PatternFlowRsvpRsvpChecksum
	ToProto() (*otg.PatternFlowRsvpRsvpChecksum, error)
	// ToPbText marshals PatternFlowRsvpRsvpChecksum to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRsvpRsvpChecksum to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRsvpRsvpChecksum to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRsvpRsvpChecksum struct {
	obj *patternFlowRsvpRsvpChecksum
}

type unMarshalPatternFlowRsvpRsvpChecksum interface {
	// FromProto unmarshals PatternFlowRsvpRsvpChecksum from protobuf object *otg.PatternFlowRsvpRsvpChecksum
	FromProto(msg *otg.PatternFlowRsvpRsvpChecksum) (PatternFlowRsvpRsvpChecksum, error)
	// FromPbText unmarshals PatternFlowRsvpRsvpChecksum from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRsvpRsvpChecksum from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRsvpRsvpChecksum from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRsvpRsvpChecksum) Marshal() marshalPatternFlowRsvpRsvpChecksum {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRsvpRsvpChecksum{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRsvpRsvpChecksum) Unmarshal() unMarshalPatternFlowRsvpRsvpChecksum {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRsvpRsvpChecksum{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRsvpRsvpChecksum) ToProto() (*otg.PatternFlowRsvpRsvpChecksum, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRsvpRsvpChecksum) FromProto(msg *otg.PatternFlowRsvpRsvpChecksum) (PatternFlowRsvpRsvpChecksum, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRsvpRsvpChecksum) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRsvpRsvpChecksum) FromPbText(value string) error {
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

func (m *marshalpatternFlowRsvpRsvpChecksum) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRsvpRsvpChecksum) FromYaml(value string) error {
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

func (m *marshalpatternFlowRsvpRsvpChecksum) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRsvpRsvpChecksum) FromJson(value string) error {
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

func (obj *patternFlowRsvpRsvpChecksum) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRsvpRsvpChecksum) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRsvpRsvpChecksum) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRsvpRsvpChecksum) Clone() (PatternFlowRsvpRsvpChecksum, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRsvpRsvpChecksum()
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

// PatternFlowRsvpRsvpChecksum is the one's complement of the one's complement sum of the message, with the checksum field replaced by zero for the purpose of computing the checksum.   An all-zero value means that no checksum was transmitted.
type PatternFlowRsvpRsvpChecksum interface {
	Validation
	// msg marshals PatternFlowRsvpRsvpChecksum to protobuf object *otg.PatternFlowRsvpRsvpChecksum
	// and doesn't set defaults
	msg() *otg.PatternFlowRsvpRsvpChecksum
	// setMsg unmarshals PatternFlowRsvpRsvpChecksum from protobuf object *otg.PatternFlowRsvpRsvpChecksum
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRsvpRsvpChecksum) PatternFlowRsvpRsvpChecksum
	// provides marshal interface
	Marshal() marshalPatternFlowRsvpRsvpChecksum
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRsvpRsvpChecksum
	// validate validates PatternFlowRsvpRsvpChecksum
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRsvpRsvpChecksum, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRsvpRsvpChecksumChoiceEnum, set in PatternFlowRsvpRsvpChecksum
	Choice() PatternFlowRsvpRsvpChecksumChoiceEnum
	// setChoice assigns PatternFlowRsvpRsvpChecksumChoiceEnum provided by user to PatternFlowRsvpRsvpChecksum
	setChoice(value PatternFlowRsvpRsvpChecksumChoiceEnum) PatternFlowRsvpRsvpChecksum
	// HasChoice checks if Choice has been set in PatternFlowRsvpRsvpChecksum
	HasChoice() bool
	// Generated returns PatternFlowRsvpRsvpChecksumGeneratedEnum, set in PatternFlowRsvpRsvpChecksum
	Generated() PatternFlowRsvpRsvpChecksumGeneratedEnum
	// SetGenerated assigns PatternFlowRsvpRsvpChecksumGeneratedEnum provided by user to PatternFlowRsvpRsvpChecksum
	SetGenerated(value PatternFlowRsvpRsvpChecksumGeneratedEnum) PatternFlowRsvpRsvpChecksum
	// HasGenerated checks if Generated has been set in PatternFlowRsvpRsvpChecksum
	HasGenerated() bool
	// Custom returns uint32, set in PatternFlowRsvpRsvpChecksum.
	Custom() uint32
	// SetCustom assigns uint32 provided by user to PatternFlowRsvpRsvpChecksum
	SetCustom(value uint32) PatternFlowRsvpRsvpChecksum
	// HasCustom checks if Custom has been set in PatternFlowRsvpRsvpChecksum
	HasCustom() bool
}

type PatternFlowRsvpRsvpChecksumChoiceEnum string

// Enum of Choice on PatternFlowRsvpRsvpChecksum
var PatternFlowRsvpRsvpChecksumChoice = struct {
	GENERATED PatternFlowRsvpRsvpChecksumChoiceEnum
	CUSTOM    PatternFlowRsvpRsvpChecksumChoiceEnum
}{
	GENERATED: PatternFlowRsvpRsvpChecksumChoiceEnum("generated"),
	CUSTOM:    PatternFlowRsvpRsvpChecksumChoiceEnum("custom"),
}

func (obj *patternFlowRsvpRsvpChecksum) Choice() PatternFlowRsvpRsvpChecksumChoiceEnum {
	return PatternFlowRsvpRsvpChecksumChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of checksum
// Choice returns a string
func (obj *patternFlowRsvpRsvpChecksum) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRsvpRsvpChecksum) setChoice(value PatternFlowRsvpRsvpChecksumChoiceEnum) PatternFlowRsvpRsvpChecksum {
	intValue, ok := otg.PatternFlowRsvpRsvpChecksum_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRsvpRsvpChecksumChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRsvpRsvpChecksum_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.obj.Generated = otg.PatternFlowRsvpRsvpChecksum_Generated_unspecified.Enum()
	return obj
}

type PatternFlowRsvpRsvpChecksumGeneratedEnum string

// Enum of Generated on PatternFlowRsvpRsvpChecksum
var PatternFlowRsvpRsvpChecksumGenerated = struct {
	GOOD PatternFlowRsvpRsvpChecksumGeneratedEnum
	BAD  PatternFlowRsvpRsvpChecksumGeneratedEnum
}{
	GOOD: PatternFlowRsvpRsvpChecksumGeneratedEnum("good"),
	BAD:  PatternFlowRsvpRsvpChecksumGeneratedEnum("bad"),
}

func (obj *patternFlowRsvpRsvpChecksum) Generated() PatternFlowRsvpRsvpChecksumGeneratedEnum {
	return PatternFlowRsvpRsvpChecksumGeneratedEnum(obj.obj.Generated.Enum().String())
}

// A system generated checksum value
// Generated returns a string
func (obj *patternFlowRsvpRsvpChecksum) HasGenerated() bool {
	return obj.obj.Generated != nil
}

func (obj *patternFlowRsvpRsvpChecksum) SetGenerated(value PatternFlowRsvpRsvpChecksumGeneratedEnum) PatternFlowRsvpRsvpChecksum {
	intValue, ok := otg.PatternFlowRsvpRsvpChecksum_Generated_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRsvpRsvpChecksumGeneratedEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRsvpRsvpChecksum_Generated_Enum(intValue)
	obj.obj.Generated = &enumValue

	return obj
}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowRsvpRsvpChecksum) Custom() uint32 {

	if obj.obj.Custom == nil {
		obj.setChoice(PatternFlowRsvpRsvpChecksumChoice.CUSTOM)
	}

	return *obj.obj.Custom

}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowRsvpRsvpChecksum) HasCustom() bool {
	return obj.obj.Custom != nil
}

// A custom checksum value
// SetCustom sets the uint32 value in the PatternFlowRsvpRsvpChecksum object
func (obj *patternFlowRsvpRsvpChecksum) SetCustom(value uint32) PatternFlowRsvpRsvpChecksum {
	obj.setChoice(PatternFlowRsvpRsvpChecksumChoice.CUSTOM)
	obj.obj.Custom = &value
	return obj
}

func (obj *patternFlowRsvpRsvpChecksum) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		if *obj.obj.Custom > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRsvpRsvpChecksum.Custom <= 65535 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternFlowRsvpRsvpChecksum) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowRsvpRsvpChecksumChoice.GENERATED)
		if obj.obj.Generated.Number() == 0 {
			obj.SetGenerated(PatternFlowRsvpRsvpChecksumGenerated.GOOD)

		}

	}

}

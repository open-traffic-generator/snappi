package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6FlowLabelRandom *****
type patternFlowIpv6FlowLabelRandom struct {
	validation
	obj          *otg.PatternFlowIpv6FlowLabelRandom
	marshaller   marshalPatternFlowIpv6FlowLabelRandom
	unMarshaller unMarshalPatternFlowIpv6FlowLabelRandom
}

func NewPatternFlowIpv6FlowLabelRandom() PatternFlowIpv6FlowLabelRandom {
	obj := patternFlowIpv6FlowLabelRandom{obj: &otg.PatternFlowIpv6FlowLabelRandom{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6FlowLabelRandom) msg() *otg.PatternFlowIpv6FlowLabelRandom {
	return obj.obj
}

func (obj *patternFlowIpv6FlowLabelRandom) setMsg(msg *otg.PatternFlowIpv6FlowLabelRandom) PatternFlowIpv6FlowLabelRandom {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6FlowLabelRandom struct {
	obj *patternFlowIpv6FlowLabelRandom
}

type marshalPatternFlowIpv6FlowLabelRandom interface {
	// ToProto marshals PatternFlowIpv6FlowLabelRandom to protobuf object *otg.PatternFlowIpv6FlowLabelRandom
	ToProto() (*otg.PatternFlowIpv6FlowLabelRandom, error)
	// ToPbText marshals PatternFlowIpv6FlowLabelRandom to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6FlowLabelRandom to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6FlowLabelRandom to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6FlowLabelRandom struct {
	obj *patternFlowIpv6FlowLabelRandom
}

type unMarshalPatternFlowIpv6FlowLabelRandom interface {
	// FromProto unmarshals PatternFlowIpv6FlowLabelRandom from protobuf object *otg.PatternFlowIpv6FlowLabelRandom
	FromProto(msg *otg.PatternFlowIpv6FlowLabelRandom) (PatternFlowIpv6FlowLabelRandom, error)
	// FromPbText unmarshals PatternFlowIpv6FlowLabelRandom from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6FlowLabelRandom from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6FlowLabelRandom from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6FlowLabelRandom) Marshal() marshalPatternFlowIpv6FlowLabelRandom {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6FlowLabelRandom{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6FlowLabelRandom) Unmarshal() unMarshalPatternFlowIpv6FlowLabelRandom {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6FlowLabelRandom{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6FlowLabelRandom) ToProto() (*otg.PatternFlowIpv6FlowLabelRandom, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6FlowLabelRandom) FromProto(msg *otg.PatternFlowIpv6FlowLabelRandom) (PatternFlowIpv6FlowLabelRandom, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6FlowLabelRandom) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6FlowLabelRandom) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6FlowLabelRandom) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6FlowLabelRandom) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6FlowLabelRandom) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6FlowLabelRandom) FromJson(value string) error {
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

func (obj *patternFlowIpv6FlowLabelRandom) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6FlowLabelRandom) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6FlowLabelRandom) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6FlowLabelRandom) Clone() (PatternFlowIpv6FlowLabelRandom, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6FlowLabelRandom()
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

// PatternFlowIpv6FlowLabelRandom is integer random pattern
type PatternFlowIpv6FlowLabelRandom interface {
	Validation
	// msg marshals PatternFlowIpv6FlowLabelRandom to protobuf object *otg.PatternFlowIpv6FlowLabelRandom
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6FlowLabelRandom
	// setMsg unmarshals PatternFlowIpv6FlowLabelRandom from protobuf object *otg.PatternFlowIpv6FlowLabelRandom
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6FlowLabelRandom) PatternFlowIpv6FlowLabelRandom
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6FlowLabelRandom
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6FlowLabelRandom
	// validate validates PatternFlowIpv6FlowLabelRandom
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6FlowLabelRandom, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Min returns uint32, set in PatternFlowIpv6FlowLabelRandom.
	Min() uint32
	// SetMin assigns uint32 provided by user to PatternFlowIpv6FlowLabelRandom
	SetMin(value uint32) PatternFlowIpv6FlowLabelRandom
	// HasMin checks if Min has been set in PatternFlowIpv6FlowLabelRandom
	HasMin() bool
	// Max returns uint32, set in PatternFlowIpv6FlowLabelRandom.
	Max() uint32
	// SetMax assigns uint32 provided by user to PatternFlowIpv6FlowLabelRandom
	SetMax(value uint32) PatternFlowIpv6FlowLabelRandom
	// HasMax checks if Max has been set in PatternFlowIpv6FlowLabelRandom
	HasMax() bool
	// Seed returns uint32, set in PatternFlowIpv6FlowLabelRandom.
	Seed() uint32
	// SetSeed assigns uint32 provided by user to PatternFlowIpv6FlowLabelRandom
	SetSeed(value uint32) PatternFlowIpv6FlowLabelRandom
	// HasSeed checks if Seed has been set in PatternFlowIpv6FlowLabelRandom
	HasSeed() bool
	// Count returns uint32, set in PatternFlowIpv6FlowLabelRandom.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6FlowLabelRandom
	SetCount(value uint32) PatternFlowIpv6FlowLabelRandom
	// HasCount checks if Count has been set in PatternFlowIpv6FlowLabelRandom
	HasCount() bool
}

// The minimum possible value generated by the random value generator.
// Min returns a uint32
func (obj *patternFlowIpv6FlowLabelRandom) Min() uint32 {

	return *obj.obj.Min

}

// The minimum possible value generated by the random value generator.
// Min returns a uint32
func (obj *patternFlowIpv6FlowLabelRandom) HasMin() bool {
	return obj.obj.Min != nil
}

// The minimum possible value generated by the random value generator.
// SetMin sets the uint32 value in the PatternFlowIpv6FlowLabelRandom object
func (obj *patternFlowIpv6FlowLabelRandom) SetMin(value uint32) PatternFlowIpv6FlowLabelRandom {

	obj.obj.Min = &value
	return obj
}

// The maximum possible value generated by the random value generator.
// Max returns a uint32
func (obj *patternFlowIpv6FlowLabelRandom) Max() uint32 {

	return *obj.obj.Max

}

// The maximum possible value generated by the random value generator.
// Max returns a uint32
func (obj *patternFlowIpv6FlowLabelRandom) HasMax() bool {
	return obj.obj.Max != nil
}

// The maximum possible value generated by the random value generator.
// SetMax sets the uint32 value in the PatternFlowIpv6FlowLabelRandom object
func (obj *patternFlowIpv6FlowLabelRandom) SetMax(value uint32) PatternFlowIpv6FlowLabelRandom {

	obj.obj.Max = &value
	return obj
}

// The seed value is used to initialize the random number generator to a deterministic state. If the user provides a seed value of 0, the implementation will generate a sequence of non-deterministic random values. For any other seed value, the sequence of random numbers will be generated in a deterministic manner (specific to the implementation).
// Seed returns a uint32
func (obj *patternFlowIpv6FlowLabelRandom) Seed() uint32 {

	return *obj.obj.Seed

}

// The seed value is used to initialize the random number generator to a deterministic state. If the user provides a seed value of 0, the implementation will generate a sequence of non-deterministic random values. For any other seed value, the sequence of random numbers will be generated in a deterministic manner (specific to the implementation).
// Seed returns a uint32
func (obj *patternFlowIpv6FlowLabelRandom) HasSeed() bool {
	return obj.obj.Seed != nil
}

// The seed value is used to initialize the random number generator to a deterministic state. If the user provides a seed value of 0, the implementation will generate a sequence of non-deterministic random values. For any other seed value, the sequence of random numbers will be generated in a deterministic manner (specific to the implementation).
// SetSeed sets the uint32 value in the PatternFlowIpv6FlowLabelRandom object
func (obj *patternFlowIpv6FlowLabelRandom) SetSeed(value uint32) PatternFlowIpv6FlowLabelRandom {

	obj.obj.Seed = &value
	return obj
}

// The total number of values to be generated by the random value generator.
// Count returns a uint32
func (obj *patternFlowIpv6FlowLabelRandom) Count() uint32 {

	return *obj.obj.Count

}

// The total number of values to be generated by the random value generator.
// Count returns a uint32
func (obj *patternFlowIpv6FlowLabelRandom) HasCount() bool {
	return obj.obj.Count != nil
}

// The total number of values to be generated by the random value generator.
// SetCount sets the uint32 value in the PatternFlowIpv6FlowLabelRandom object
func (obj *patternFlowIpv6FlowLabelRandom) SetCount(value uint32) PatternFlowIpv6FlowLabelRandom {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6FlowLabelRandom) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Min != nil {

		if *obj.obj.Min > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6FlowLabelRandom.Min <= 1048575 but Got %d", *obj.obj.Min))
		}

	}

	if obj.obj.Max != nil {

		if *obj.obj.Max > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6FlowLabelRandom.Max <= 1048575 but Got %d", *obj.obj.Max))
		}

	}

}

func (obj *patternFlowIpv6FlowLabelRandom) setDefault() {
	if obj.obj.Min == nil {
		obj.SetMin(0)
	}
	if obj.obj.Max == nil {
		obj.SetMax(1048575)
	}
	if obj.obj.Seed == nil {
		obj.SetSeed(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

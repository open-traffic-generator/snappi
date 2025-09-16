package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowUdpSrcPortRandom *****
type patternFlowUdpSrcPortRandom struct {
	validation
	obj          *otg.PatternFlowUdpSrcPortRandom
	marshaller   marshalPatternFlowUdpSrcPortRandom
	unMarshaller unMarshalPatternFlowUdpSrcPortRandom
}

func NewPatternFlowUdpSrcPortRandom() PatternFlowUdpSrcPortRandom {
	obj := patternFlowUdpSrcPortRandom{obj: &otg.PatternFlowUdpSrcPortRandom{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowUdpSrcPortRandom) msg() *otg.PatternFlowUdpSrcPortRandom {
	return obj.obj
}

func (obj *patternFlowUdpSrcPortRandom) setMsg(msg *otg.PatternFlowUdpSrcPortRandom) PatternFlowUdpSrcPortRandom {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowUdpSrcPortRandom struct {
	obj *patternFlowUdpSrcPortRandom
}

type marshalPatternFlowUdpSrcPortRandom interface {
	// ToProto marshals PatternFlowUdpSrcPortRandom to protobuf object *otg.PatternFlowUdpSrcPortRandom
	ToProto() (*otg.PatternFlowUdpSrcPortRandom, error)
	// ToPbText marshals PatternFlowUdpSrcPortRandom to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowUdpSrcPortRandom to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowUdpSrcPortRandom to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowUdpSrcPortRandom struct {
	obj *patternFlowUdpSrcPortRandom
}

type unMarshalPatternFlowUdpSrcPortRandom interface {
	// FromProto unmarshals PatternFlowUdpSrcPortRandom from protobuf object *otg.PatternFlowUdpSrcPortRandom
	FromProto(msg *otg.PatternFlowUdpSrcPortRandom) (PatternFlowUdpSrcPortRandom, error)
	// FromPbText unmarshals PatternFlowUdpSrcPortRandom from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowUdpSrcPortRandom from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowUdpSrcPortRandom from JSON text
	FromJson(value string) error
}

func (obj *patternFlowUdpSrcPortRandom) Marshal() marshalPatternFlowUdpSrcPortRandom {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowUdpSrcPortRandom{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowUdpSrcPortRandom) Unmarshal() unMarshalPatternFlowUdpSrcPortRandom {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowUdpSrcPortRandom{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowUdpSrcPortRandom) ToProto() (*otg.PatternFlowUdpSrcPortRandom, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowUdpSrcPortRandom) FromProto(msg *otg.PatternFlowUdpSrcPortRandom) (PatternFlowUdpSrcPortRandom, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowUdpSrcPortRandom) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowUdpSrcPortRandom) FromPbText(value string) error {
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

func (m *marshalpatternFlowUdpSrcPortRandom) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowUdpSrcPortRandom) FromYaml(value string) error {
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

func (m *marshalpatternFlowUdpSrcPortRandom) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowUdpSrcPortRandom) FromJson(value string) error {
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

func (obj *patternFlowUdpSrcPortRandom) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowUdpSrcPortRandom) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowUdpSrcPortRandom) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowUdpSrcPortRandom) Clone() (PatternFlowUdpSrcPortRandom, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowUdpSrcPortRandom()
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

// PatternFlowUdpSrcPortRandom is integer random pattern
type PatternFlowUdpSrcPortRandom interface {
	Validation
	// msg marshals PatternFlowUdpSrcPortRandom to protobuf object *otg.PatternFlowUdpSrcPortRandom
	// and doesn't set defaults
	msg() *otg.PatternFlowUdpSrcPortRandom
	// setMsg unmarshals PatternFlowUdpSrcPortRandom from protobuf object *otg.PatternFlowUdpSrcPortRandom
	// and doesn't set defaults
	setMsg(*otg.PatternFlowUdpSrcPortRandom) PatternFlowUdpSrcPortRandom
	// provides marshal interface
	Marshal() marshalPatternFlowUdpSrcPortRandom
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowUdpSrcPortRandom
	// validate validates PatternFlowUdpSrcPortRandom
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowUdpSrcPortRandom, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Min returns uint32, set in PatternFlowUdpSrcPortRandom.
	Min() uint32
	// SetMin assigns uint32 provided by user to PatternFlowUdpSrcPortRandom
	SetMin(value uint32) PatternFlowUdpSrcPortRandom
	// HasMin checks if Min has been set in PatternFlowUdpSrcPortRandom
	HasMin() bool
	// Max returns uint32, set in PatternFlowUdpSrcPortRandom.
	Max() uint32
	// SetMax assigns uint32 provided by user to PatternFlowUdpSrcPortRandom
	SetMax(value uint32) PatternFlowUdpSrcPortRandom
	// HasMax checks if Max has been set in PatternFlowUdpSrcPortRandom
	HasMax() bool
	// Seed returns uint32, set in PatternFlowUdpSrcPortRandom.
	Seed() uint32
	// SetSeed assigns uint32 provided by user to PatternFlowUdpSrcPortRandom
	SetSeed(value uint32) PatternFlowUdpSrcPortRandom
	// HasSeed checks if Seed has been set in PatternFlowUdpSrcPortRandom
	HasSeed() bool
	// Count returns uint32, set in PatternFlowUdpSrcPortRandom.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowUdpSrcPortRandom
	SetCount(value uint32) PatternFlowUdpSrcPortRandom
	// HasCount checks if Count has been set in PatternFlowUdpSrcPortRandom
	HasCount() bool
}

// The minimum possible value generated by the random value generator.
// Min returns a uint32
func (obj *patternFlowUdpSrcPortRandom) Min() uint32 {

	return *obj.obj.Min

}

// The minimum possible value generated by the random value generator.
// Min returns a uint32
func (obj *patternFlowUdpSrcPortRandom) HasMin() bool {
	return obj.obj.Min != nil
}

// The minimum possible value generated by the random value generator.
// SetMin sets the uint32 value in the PatternFlowUdpSrcPortRandom object
func (obj *patternFlowUdpSrcPortRandom) SetMin(value uint32) PatternFlowUdpSrcPortRandom {

	obj.obj.Min = &value
	return obj
}

// The maximum possible value generated by the random value generator.
// Max returns a uint32
func (obj *patternFlowUdpSrcPortRandom) Max() uint32 {

	return *obj.obj.Max

}

// The maximum possible value generated by the random value generator.
// Max returns a uint32
func (obj *patternFlowUdpSrcPortRandom) HasMax() bool {
	return obj.obj.Max != nil
}

// The maximum possible value generated by the random value generator.
// SetMax sets the uint32 value in the PatternFlowUdpSrcPortRandom object
func (obj *patternFlowUdpSrcPortRandom) SetMax(value uint32) PatternFlowUdpSrcPortRandom {

	obj.obj.Max = &value
	return obj
}

// The seed value is used to initialize the random number generator to a deterministic state. If the user provides a seed value of 0, the implementation will generate a sequence of non-deterministic random values. For any other seed value, the sequence of random numbers will be generated in a deterministic manner (specific to the implementation).
// Seed returns a uint32
func (obj *patternFlowUdpSrcPortRandom) Seed() uint32 {

	return *obj.obj.Seed

}

// The seed value is used to initialize the random number generator to a deterministic state. If the user provides a seed value of 0, the implementation will generate a sequence of non-deterministic random values. For any other seed value, the sequence of random numbers will be generated in a deterministic manner (specific to the implementation).
// Seed returns a uint32
func (obj *patternFlowUdpSrcPortRandom) HasSeed() bool {
	return obj.obj.Seed != nil
}

// The seed value is used to initialize the random number generator to a deterministic state. If the user provides a seed value of 0, the implementation will generate a sequence of non-deterministic random values. For any other seed value, the sequence of random numbers will be generated in a deterministic manner (specific to the implementation).
// SetSeed sets the uint32 value in the PatternFlowUdpSrcPortRandom object
func (obj *patternFlowUdpSrcPortRandom) SetSeed(value uint32) PatternFlowUdpSrcPortRandom {

	obj.obj.Seed = &value
	return obj
}

// The total number of values to be generated by the random value generator.
// Count returns a uint32
func (obj *patternFlowUdpSrcPortRandom) Count() uint32 {

	return *obj.obj.Count

}

// The total number of values to be generated by the random value generator.
// Count returns a uint32
func (obj *patternFlowUdpSrcPortRandom) HasCount() bool {
	return obj.obj.Count != nil
}

// The total number of values to be generated by the random value generator.
// SetCount sets the uint32 value in the PatternFlowUdpSrcPortRandom object
func (obj *patternFlowUdpSrcPortRandom) SetCount(value uint32) PatternFlowUdpSrcPortRandom {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowUdpSrcPortRandom) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Min != nil {

		if *obj.obj.Min > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowUdpSrcPortRandom.Min <= 65535 but Got %d", *obj.obj.Min))
		}

	}

	if obj.obj.Max != nil {

		if *obj.obj.Max > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowUdpSrcPortRandom.Max <= 65535 but Got %d", *obj.obj.Max))
		}

	}

}

func (obj *patternFlowUdpSrcPortRandom) setDefault() {
	if obj.obj.Min == nil {
		obj.SetMin(0)
	}
	if obj.obj.Max == nil {
		obj.SetMax(65535)
	}
	if obj.obj.Seed == nil {
		obj.SetSeed(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpSrcPortRandom *****
type patternFlowTcpSrcPortRandom struct {
	validation
	obj          *otg.PatternFlowTcpSrcPortRandom
	marshaller   marshalPatternFlowTcpSrcPortRandom
	unMarshaller unMarshalPatternFlowTcpSrcPortRandom
}

func NewPatternFlowTcpSrcPortRandom() PatternFlowTcpSrcPortRandom {
	obj := patternFlowTcpSrcPortRandom{obj: &otg.PatternFlowTcpSrcPortRandom{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpSrcPortRandom) msg() *otg.PatternFlowTcpSrcPortRandom {
	return obj.obj
}

func (obj *patternFlowTcpSrcPortRandom) setMsg(msg *otg.PatternFlowTcpSrcPortRandom) PatternFlowTcpSrcPortRandom {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpSrcPortRandom struct {
	obj *patternFlowTcpSrcPortRandom
}

type marshalPatternFlowTcpSrcPortRandom interface {
	// ToProto marshals PatternFlowTcpSrcPortRandom to protobuf object *otg.PatternFlowTcpSrcPortRandom
	ToProto() (*otg.PatternFlowTcpSrcPortRandom, error)
	// ToPbText marshals PatternFlowTcpSrcPortRandom to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpSrcPortRandom to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpSrcPortRandom to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpSrcPortRandom struct {
	obj *patternFlowTcpSrcPortRandom
}

type unMarshalPatternFlowTcpSrcPortRandom interface {
	// FromProto unmarshals PatternFlowTcpSrcPortRandom from protobuf object *otg.PatternFlowTcpSrcPortRandom
	FromProto(msg *otg.PatternFlowTcpSrcPortRandom) (PatternFlowTcpSrcPortRandom, error)
	// FromPbText unmarshals PatternFlowTcpSrcPortRandom from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpSrcPortRandom from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpSrcPortRandom from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpSrcPortRandom) Marshal() marshalPatternFlowTcpSrcPortRandom {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpSrcPortRandom{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpSrcPortRandom) Unmarshal() unMarshalPatternFlowTcpSrcPortRandom {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpSrcPortRandom{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpSrcPortRandom) ToProto() (*otg.PatternFlowTcpSrcPortRandom, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpSrcPortRandom) FromProto(msg *otg.PatternFlowTcpSrcPortRandom) (PatternFlowTcpSrcPortRandom, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpSrcPortRandom) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpSrcPortRandom) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpSrcPortRandom) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpSrcPortRandom) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpSrcPortRandom) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpSrcPortRandom) FromJson(value string) error {
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

func (obj *patternFlowTcpSrcPortRandom) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpSrcPortRandom) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpSrcPortRandom) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpSrcPortRandom) Clone() (PatternFlowTcpSrcPortRandom, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpSrcPortRandom()
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

// PatternFlowTcpSrcPortRandom is integer random pattern
type PatternFlowTcpSrcPortRandom interface {
	Validation
	// msg marshals PatternFlowTcpSrcPortRandom to protobuf object *otg.PatternFlowTcpSrcPortRandom
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpSrcPortRandom
	// setMsg unmarshals PatternFlowTcpSrcPortRandom from protobuf object *otg.PatternFlowTcpSrcPortRandom
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpSrcPortRandom) PatternFlowTcpSrcPortRandom
	// provides marshal interface
	Marshal() marshalPatternFlowTcpSrcPortRandom
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpSrcPortRandom
	// validate validates PatternFlowTcpSrcPortRandom
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpSrcPortRandom, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Min returns uint32, set in PatternFlowTcpSrcPortRandom.
	Min() uint32
	// SetMin assigns uint32 provided by user to PatternFlowTcpSrcPortRandom
	SetMin(value uint32) PatternFlowTcpSrcPortRandom
	// HasMin checks if Min has been set in PatternFlowTcpSrcPortRandom
	HasMin() bool
	// Max returns uint32, set in PatternFlowTcpSrcPortRandom.
	Max() uint32
	// SetMax assigns uint32 provided by user to PatternFlowTcpSrcPortRandom
	SetMax(value uint32) PatternFlowTcpSrcPortRandom
	// HasMax checks if Max has been set in PatternFlowTcpSrcPortRandom
	HasMax() bool
	// Seed returns uint32, set in PatternFlowTcpSrcPortRandom.
	Seed() uint32
	// SetSeed assigns uint32 provided by user to PatternFlowTcpSrcPortRandom
	SetSeed(value uint32) PatternFlowTcpSrcPortRandom
	// HasSeed checks if Seed has been set in PatternFlowTcpSrcPortRandom
	HasSeed() bool
	// Count returns uint32, set in PatternFlowTcpSrcPortRandom.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowTcpSrcPortRandom
	SetCount(value uint32) PatternFlowTcpSrcPortRandom
	// HasCount checks if Count has been set in PatternFlowTcpSrcPortRandom
	HasCount() bool
}

// The minimum possible value generated by the random value generator.
// Min returns a uint32
func (obj *patternFlowTcpSrcPortRandom) Min() uint32 {

	return *obj.obj.Min

}

// The minimum possible value generated by the random value generator.
// Min returns a uint32
func (obj *patternFlowTcpSrcPortRandom) HasMin() bool {
	return obj.obj.Min != nil
}

// The minimum possible value generated by the random value generator.
// SetMin sets the uint32 value in the PatternFlowTcpSrcPortRandom object
func (obj *patternFlowTcpSrcPortRandom) SetMin(value uint32) PatternFlowTcpSrcPortRandom {

	obj.obj.Min = &value
	return obj
}

// The maximum possible value generated by the random value generator.
// Max returns a uint32
func (obj *patternFlowTcpSrcPortRandom) Max() uint32 {

	return *obj.obj.Max

}

// The maximum possible value generated by the random value generator.
// Max returns a uint32
func (obj *patternFlowTcpSrcPortRandom) HasMax() bool {
	return obj.obj.Max != nil
}

// The maximum possible value generated by the random value generator.
// SetMax sets the uint32 value in the PatternFlowTcpSrcPortRandom object
func (obj *patternFlowTcpSrcPortRandom) SetMax(value uint32) PatternFlowTcpSrcPortRandom {

	obj.obj.Max = &value
	return obj
}

// The seed value is used to initialize the random number generator to a deterministic state. If the user provides a seed value of 0, the implementation will generate a sequence of non-deterministic random values. For any other seed value, the sequence of random numbers will be generated in a deterministic manner (specific to the implementation).
// Seed returns a uint32
func (obj *patternFlowTcpSrcPortRandom) Seed() uint32 {

	return *obj.obj.Seed

}

// The seed value is used to initialize the random number generator to a deterministic state. If the user provides a seed value of 0, the implementation will generate a sequence of non-deterministic random values. For any other seed value, the sequence of random numbers will be generated in a deterministic manner (specific to the implementation).
// Seed returns a uint32
func (obj *patternFlowTcpSrcPortRandom) HasSeed() bool {
	return obj.obj.Seed != nil
}

// The seed value is used to initialize the random number generator to a deterministic state. If the user provides a seed value of 0, the implementation will generate a sequence of non-deterministic random values. For any other seed value, the sequence of random numbers will be generated in a deterministic manner (specific to the implementation).
// SetSeed sets the uint32 value in the PatternFlowTcpSrcPortRandom object
func (obj *patternFlowTcpSrcPortRandom) SetSeed(value uint32) PatternFlowTcpSrcPortRandom {

	obj.obj.Seed = &value
	return obj
}

// The total number of values to be generated by the random value generator.
// Count returns a uint32
func (obj *patternFlowTcpSrcPortRandom) Count() uint32 {

	return *obj.obj.Count

}

// The total number of values to be generated by the random value generator.
// Count returns a uint32
func (obj *patternFlowTcpSrcPortRandom) HasCount() bool {
	return obj.obj.Count != nil
}

// The total number of values to be generated by the random value generator.
// SetCount sets the uint32 value in the PatternFlowTcpSrcPortRandom object
func (obj *patternFlowTcpSrcPortRandom) SetCount(value uint32) PatternFlowTcpSrcPortRandom {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowTcpSrcPortRandom) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Min != nil {

		if *obj.obj.Min > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpSrcPortRandom.Min <= 65535 but Got %d", *obj.obj.Min))
		}

	}

	if obj.obj.Max != nil {

		if *obj.obj.Max > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpSrcPortRandom.Max <= 65535 but Got %d", *obj.obj.Max))
		}

	}

}

func (obj *patternFlowTcpSrcPortRandom) setDefault() {
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

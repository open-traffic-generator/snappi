package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowUdpDstPortRandom *****
type patternFlowUdpDstPortRandom struct {
	validation
	obj          *otg.PatternFlowUdpDstPortRandom
	marshaller   marshalPatternFlowUdpDstPortRandom
	unMarshaller unMarshalPatternFlowUdpDstPortRandom
}

func NewPatternFlowUdpDstPortRandom() PatternFlowUdpDstPortRandom {
	obj := patternFlowUdpDstPortRandom{obj: &otg.PatternFlowUdpDstPortRandom{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowUdpDstPortRandom) msg() *otg.PatternFlowUdpDstPortRandom {
	return obj.obj
}

func (obj *patternFlowUdpDstPortRandom) setMsg(msg *otg.PatternFlowUdpDstPortRandom) PatternFlowUdpDstPortRandom {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowUdpDstPortRandom struct {
	obj *patternFlowUdpDstPortRandom
}

type marshalPatternFlowUdpDstPortRandom interface {
	// ToProto marshals PatternFlowUdpDstPortRandom to protobuf object *otg.PatternFlowUdpDstPortRandom
	ToProto() (*otg.PatternFlowUdpDstPortRandom, error)
	// ToPbText marshals PatternFlowUdpDstPortRandom to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowUdpDstPortRandom to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowUdpDstPortRandom to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowUdpDstPortRandom to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowUdpDstPortRandom struct {
	obj *patternFlowUdpDstPortRandom
}

type unMarshalPatternFlowUdpDstPortRandom interface {
	// FromProto unmarshals PatternFlowUdpDstPortRandom from protobuf object *otg.PatternFlowUdpDstPortRandom
	FromProto(msg *otg.PatternFlowUdpDstPortRandom) (PatternFlowUdpDstPortRandom, error)
	// FromPbText unmarshals PatternFlowUdpDstPortRandom from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowUdpDstPortRandom from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowUdpDstPortRandom from JSON text
	FromJson(value string) error
}

func (obj *patternFlowUdpDstPortRandom) Marshal() marshalPatternFlowUdpDstPortRandom {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowUdpDstPortRandom{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowUdpDstPortRandom) Unmarshal() unMarshalPatternFlowUdpDstPortRandom {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowUdpDstPortRandom{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowUdpDstPortRandom) ToProto() (*otg.PatternFlowUdpDstPortRandom, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowUdpDstPortRandom) FromProto(msg *otg.PatternFlowUdpDstPortRandom) (PatternFlowUdpDstPortRandom, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowUdpDstPortRandom) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowUdpDstPortRandom) FromPbText(value string) error {
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

func (m *marshalpatternFlowUdpDstPortRandom) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowUdpDstPortRandom) FromYaml(value string) error {
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

func (m *marshalpatternFlowUdpDstPortRandom) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowUdpDstPortRandom) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowUdpDstPortRandom) FromJson(value string) error {
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

func (obj *patternFlowUdpDstPortRandom) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowUdpDstPortRandom) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowUdpDstPortRandom) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowUdpDstPortRandom) Clone() (PatternFlowUdpDstPortRandom, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowUdpDstPortRandom()
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

// PatternFlowUdpDstPortRandom is integer random pattern
type PatternFlowUdpDstPortRandom interface {
	Validation
	// msg marshals PatternFlowUdpDstPortRandom to protobuf object *otg.PatternFlowUdpDstPortRandom
	// and doesn't set defaults
	msg() *otg.PatternFlowUdpDstPortRandom
	// setMsg unmarshals PatternFlowUdpDstPortRandom from protobuf object *otg.PatternFlowUdpDstPortRandom
	// and doesn't set defaults
	setMsg(*otg.PatternFlowUdpDstPortRandom) PatternFlowUdpDstPortRandom
	// provides marshal interface
	Marshal() marshalPatternFlowUdpDstPortRandom
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowUdpDstPortRandom
	// validate validates PatternFlowUdpDstPortRandom
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowUdpDstPortRandom, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Min returns uint32, set in PatternFlowUdpDstPortRandom.
	Min() uint32
	// SetMin assigns uint32 provided by user to PatternFlowUdpDstPortRandom
	SetMin(value uint32) PatternFlowUdpDstPortRandom
	// HasMin checks if Min has been set in PatternFlowUdpDstPortRandom
	HasMin() bool
	// Max returns uint32, set in PatternFlowUdpDstPortRandom.
	Max() uint32
	// SetMax assigns uint32 provided by user to PatternFlowUdpDstPortRandom
	SetMax(value uint32) PatternFlowUdpDstPortRandom
	// HasMax checks if Max has been set in PatternFlowUdpDstPortRandom
	HasMax() bool
	// Seed returns uint32, set in PatternFlowUdpDstPortRandom.
	Seed() uint32
	// SetSeed assigns uint32 provided by user to PatternFlowUdpDstPortRandom
	SetSeed(value uint32) PatternFlowUdpDstPortRandom
	// HasSeed checks if Seed has been set in PatternFlowUdpDstPortRandom
	HasSeed() bool
	// Count returns uint32, set in PatternFlowUdpDstPortRandom.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowUdpDstPortRandom
	SetCount(value uint32) PatternFlowUdpDstPortRandom
	// HasCount checks if Count has been set in PatternFlowUdpDstPortRandom
	HasCount() bool
}

// The minimum possible value generated by the random value generator.
// Min returns a uint32
func (obj *patternFlowUdpDstPortRandom) Min() uint32 {

	return *obj.obj.Min

}

// The minimum possible value generated by the random value generator.
// Min returns a uint32
func (obj *patternFlowUdpDstPortRandom) HasMin() bool {
	return obj.obj.Min != nil
}

// The minimum possible value generated by the random value generator.
// SetMin sets the uint32 value in the PatternFlowUdpDstPortRandom object
func (obj *patternFlowUdpDstPortRandom) SetMin(value uint32) PatternFlowUdpDstPortRandom {

	obj.obj.Min = &value
	return obj
}

// The maximum possible value generated by the random value generator.
// Max returns a uint32
func (obj *patternFlowUdpDstPortRandom) Max() uint32 {

	return *obj.obj.Max

}

// The maximum possible value generated by the random value generator.
// Max returns a uint32
func (obj *patternFlowUdpDstPortRandom) HasMax() bool {
	return obj.obj.Max != nil
}

// The maximum possible value generated by the random value generator.
// SetMax sets the uint32 value in the PatternFlowUdpDstPortRandom object
func (obj *patternFlowUdpDstPortRandom) SetMax(value uint32) PatternFlowUdpDstPortRandom {

	obj.obj.Max = &value
	return obj
}

// The seed value is used to initialize the random number generator to a deterministic state. If the user provides a seed value of 0, the implementation will generate a sequence of non-deterministic random values. For any other seed value, the sequence of random numbers will be generated in a deterministic manner (specific to the implementation).
// Seed returns a uint32
func (obj *patternFlowUdpDstPortRandom) Seed() uint32 {

	return *obj.obj.Seed

}

// The seed value is used to initialize the random number generator to a deterministic state. If the user provides a seed value of 0, the implementation will generate a sequence of non-deterministic random values. For any other seed value, the sequence of random numbers will be generated in a deterministic manner (specific to the implementation).
// Seed returns a uint32
func (obj *patternFlowUdpDstPortRandom) HasSeed() bool {
	return obj.obj.Seed != nil
}

// The seed value is used to initialize the random number generator to a deterministic state. If the user provides a seed value of 0, the implementation will generate a sequence of non-deterministic random values. For any other seed value, the sequence of random numbers will be generated in a deterministic manner (specific to the implementation).
// SetSeed sets the uint32 value in the PatternFlowUdpDstPortRandom object
func (obj *patternFlowUdpDstPortRandom) SetSeed(value uint32) PatternFlowUdpDstPortRandom {

	obj.obj.Seed = &value
	return obj
}

// The total number of values to be generated by the random value generator.
// Count returns a uint32
func (obj *patternFlowUdpDstPortRandom) Count() uint32 {

	return *obj.obj.Count

}

// The total number of values to be generated by the random value generator.
// Count returns a uint32
func (obj *patternFlowUdpDstPortRandom) HasCount() bool {
	return obj.obj.Count != nil
}

// The total number of values to be generated by the random value generator.
// SetCount sets the uint32 value in the PatternFlowUdpDstPortRandom object
func (obj *patternFlowUdpDstPortRandom) SetCount(value uint32) PatternFlowUdpDstPortRandom {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowUdpDstPortRandom) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Min != nil {

		if *obj.obj.Min > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowUdpDstPortRandom.Min <= 65535 but Got %d", *obj.obj.Min))
		}

	}

	if obj.obj.Max != nil {

		if *obj.obj.Max > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowUdpDstPortRandom.Max <= 65535 but Got %d", *obj.obj.Max))
		}

	}

}

func (obj *patternFlowUdpDstPortRandom) setDefault() {
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

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpDstPortRandom *****
type patternFlowTcpDstPortRandom struct {
	validation
	obj          *otg.PatternFlowTcpDstPortRandom
	marshaller   marshalPatternFlowTcpDstPortRandom
	unMarshaller unMarshalPatternFlowTcpDstPortRandom
}

func NewPatternFlowTcpDstPortRandom() PatternFlowTcpDstPortRandom {
	obj := patternFlowTcpDstPortRandom{obj: &otg.PatternFlowTcpDstPortRandom{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpDstPortRandom) msg() *otg.PatternFlowTcpDstPortRandom {
	return obj.obj
}

func (obj *patternFlowTcpDstPortRandom) setMsg(msg *otg.PatternFlowTcpDstPortRandom) PatternFlowTcpDstPortRandom {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpDstPortRandom struct {
	obj *patternFlowTcpDstPortRandom
}

type marshalPatternFlowTcpDstPortRandom interface {
	// ToProto marshals PatternFlowTcpDstPortRandom to protobuf object *otg.PatternFlowTcpDstPortRandom
	ToProto() (*otg.PatternFlowTcpDstPortRandom, error)
	// ToPbText marshals PatternFlowTcpDstPortRandom to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpDstPortRandom to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpDstPortRandom to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowTcpDstPortRandom to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowTcpDstPortRandom struct {
	obj *patternFlowTcpDstPortRandom
}

type unMarshalPatternFlowTcpDstPortRandom interface {
	// FromProto unmarshals PatternFlowTcpDstPortRandom from protobuf object *otg.PatternFlowTcpDstPortRandom
	FromProto(msg *otg.PatternFlowTcpDstPortRandom) (PatternFlowTcpDstPortRandom, error)
	// FromPbText unmarshals PatternFlowTcpDstPortRandom from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpDstPortRandom from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpDstPortRandom from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpDstPortRandom) Marshal() marshalPatternFlowTcpDstPortRandom {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpDstPortRandom{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpDstPortRandom) Unmarshal() unMarshalPatternFlowTcpDstPortRandom {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpDstPortRandom{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpDstPortRandom) ToProto() (*otg.PatternFlowTcpDstPortRandom, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpDstPortRandom) FromProto(msg *otg.PatternFlowTcpDstPortRandom) (PatternFlowTcpDstPortRandom, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpDstPortRandom) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpDstPortRandom) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpDstPortRandom) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpDstPortRandom) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpDstPortRandom) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowTcpDstPortRandom) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpDstPortRandom) FromJson(value string) error {
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

func (obj *patternFlowTcpDstPortRandom) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpDstPortRandom) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpDstPortRandom) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpDstPortRandom) Clone() (PatternFlowTcpDstPortRandom, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpDstPortRandom()
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

// PatternFlowTcpDstPortRandom is integer random pattern
type PatternFlowTcpDstPortRandom interface {
	Validation
	// msg marshals PatternFlowTcpDstPortRandom to protobuf object *otg.PatternFlowTcpDstPortRandom
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpDstPortRandom
	// setMsg unmarshals PatternFlowTcpDstPortRandom from protobuf object *otg.PatternFlowTcpDstPortRandom
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpDstPortRandom) PatternFlowTcpDstPortRandom
	// provides marshal interface
	Marshal() marshalPatternFlowTcpDstPortRandom
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpDstPortRandom
	// validate validates PatternFlowTcpDstPortRandom
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpDstPortRandom, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Min returns uint32, set in PatternFlowTcpDstPortRandom.
	Min() uint32
	// SetMin assigns uint32 provided by user to PatternFlowTcpDstPortRandom
	SetMin(value uint32) PatternFlowTcpDstPortRandom
	// HasMin checks if Min has been set in PatternFlowTcpDstPortRandom
	HasMin() bool
	// Max returns uint32, set in PatternFlowTcpDstPortRandom.
	Max() uint32
	// SetMax assigns uint32 provided by user to PatternFlowTcpDstPortRandom
	SetMax(value uint32) PatternFlowTcpDstPortRandom
	// HasMax checks if Max has been set in PatternFlowTcpDstPortRandom
	HasMax() bool
	// Seed returns uint32, set in PatternFlowTcpDstPortRandom.
	Seed() uint32
	// SetSeed assigns uint32 provided by user to PatternFlowTcpDstPortRandom
	SetSeed(value uint32) PatternFlowTcpDstPortRandom
	// HasSeed checks if Seed has been set in PatternFlowTcpDstPortRandom
	HasSeed() bool
	// Count returns uint32, set in PatternFlowTcpDstPortRandom.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowTcpDstPortRandom
	SetCount(value uint32) PatternFlowTcpDstPortRandom
	// HasCount checks if Count has been set in PatternFlowTcpDstPortRandom
	HasCount() bool
}

// The minimum possible value generated by the random value generator.
// Min returns a uint32
func (obj *patternFlowTcpDstPortRandom) Min() uint32 {

	return *obj.obj.Min

}

// The minimum possible value generated by the random value generator.
// Min returns a uint32
func (obj *patternFlowTcpDstPortRandom) HasMin() bool {
	return obj.obj.Min != nil
}

// The minimum possible value generated by the random value generator.
// SetMin sets the uint32 value in the PatternFlowTcpDstPortRandom object
func (obj *patternFlowTcpDstPortRandom) SetMin(value uint32) PatternFlowTcpDstPortRandom {

	obj.obj.Min = &value
	return obj
}

// The maximum possible value generated by the random value generator.
// Max returns a uint32
func (obj *patternFlowTcpDstPortRandom) Max() uint32 {

	return *obj.obj.Max

}

// The maximum possible value generated by the random value generator.
// Max returns a uint32
func (obj *patternFlowTcpDstPortRandom) HasMax() bool {
	return obj.obj.Max != nil
}

// The maximum possible value generated by the random value generator.
// SetMax sets the uint32 value in the PatternFlowTcpDstPortRandom object
func (obj *patternFlowTcpDstPortRandom) SetMax(value uint32) PatternFlowTcpDstPortRandom {

	obj.obj.Max = &value
	return obj
}

// The seed value is used to initialize the random number generator to a deterministic state. If the user provides a seed value of 0, the implementation will generate a sequence of non-deterministic random values. For any other seed value, the sequence of random numbers will be generated in a deterministic manner (specific to the implementation).
// Seed returns a uint32
func (obj *patternFlowTcpDstPortRandom) Seed() uint32 {

	return *obj.obj.Seed

}

// The seed value is used to initialize the random number generator to a deterministic state. If the user provides a seed value of 0, the implementation will generate a sequence of non-deterministic random values. For any other seed value, the sequence of random numbers will be generated in a deterministic manner (specific to the implementation).
// Seed returns a uint32
func (obj *patternFlowTcpDstPortRandom) HasSeed() bool {
	return obj.obj.Seed != nil
}

// The seed value is used to initialize the random number generator to a deterministic state. If the user provides a seed value of 0, the implementation will generate a sequence of non-deterministic random values. For any other seed value, the sequence of random numbers will be generated in a deterministic manner (specific to the implementation).
// SetSeed sets the uint32 value in the PatternFlowTcpDstPortRandom object
func (obj *patternFlowTcpDstPortRandom) SetSeed(value uint32) PatternFlowTcpDstPortRandom {

	obj.obj.Seed = &value
	return obj
}

// The total number of values to be generated by the random value generator.
// Count returns a uint32
func (obj *patternFlowTcpDstPortRandom) Count() uint32 {

	return *obj.obj.Count

}

// The total number of values to be generated by the random value generator.
// Count returns a uint32
func (obj *patternFlowTcpDstPortRandom) HasCount() bool {
	return obj.obj.Count != nil
}

// The total number of values to be generated by the random value generator.
// SetCount sets the uint32 value in the PatternFlowTcpDstPortRandom object
func (obj *patternFlowTcpDstPortRandom) SetCount(value uint32) PatternFlowTcpDstPortRandom {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowTcpDstPortRandom) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Min != nil {

		if *obj.obj.Min > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpDstPortRandom.Min <= 65535 but Got %d", *obj.obj.Min))
		}

	}

	if obj.obj.Max != nil {

		if *obj.obj.Max > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpDstPortRandom.Max <= 65535 but Got %d", *obj.obj.Max))
		}

	}

}

func (obj *patternFlowTcpDstPortRandom) setDefault() {
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

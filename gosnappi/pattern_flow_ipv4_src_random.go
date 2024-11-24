package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4SrcRandom *****
type patternFlowIpv4SrcRandom struct {
	validation
	obj          *otg.PatternFlowIpv4SrcRandom
	marshaller   marshalPatternFlowIpv4SrcRandom
	unMarshaller unMarshalPatternFlowIpv4SrcRandom
}

func NewPatternFlowIpv4SrcRandom() PatternFlowIpv4SrcRandom {
	obj := patternFlowIpv4SrcRandom{obj: &otg.PatternFlowIpv4SrcRandom{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4SrcRandom) msg() *otg.PatternFlowIpv4SrcRandom {
	return obj.obj
}

func (obj *patternFlowIpv4SrcRandom) setMsg(msg *otg.PatternFlowIpv4SrcRandom) PatternFlowIpv4SrcRandom {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4SrcRandom struct {
	obj *patternFlowIpv4SrcRandom
}

type marshalPatternFlowIpv4SrcRandom interface {
	// ToProto marshals PatternFlowIpv4SrcRandom to protobuf object *otg.PatternFlowIpv4SrcRandom
	ToProto() (*otg.PatternFlowIpv4SrcRandom, error)
	// ToPbText marshals PatternFlowIpv4SrcRandom to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4SrcRandom to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4SrcRandom to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4SrcRandom to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4SrcRandom struct {
	obj *patternFlowIpv4SrcRandom
}

type unMarshalPatternFlowIpv4SrcRandom interface {
	// FromProto unmarshals PatternFlowIpv4SrcRandom from protobuf object *otg.PatternFlowIpv4SrcRandom
	FromProto(msg *otg.PatternFlowIpv4SrcRandom) (PatternFlowIpv4SrcRandom, error)
	// FromPbText unmarshals PatternFlowIpv4SrcRandom from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4SrcRandom from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4SrcRandom from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4SrcRandom) Marshal() marshalPatternFlowIpv4SrcRandom {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4SrcRandom{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4SrcRandom) Unmarshal() unMarshalPatternFlowIpv4SrcRandom {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4SrcRandom{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4SrcRandom) ToProto() (*otg.PatternFlowIpv4SrcRandom, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4SrcRandom) FromProto(msg *otg.PatternFlowIpv4SrcRandom) (PatternFlowIpv4SrcRandom, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4SrcRandom) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4SrcRandom) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4SrcRandom) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4SrcRandom) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4SrcRandom) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4SrcRandom) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4SrcRandom) FromJson(value string) error {
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

func (obj *patternFlowIpv4SrcRandom) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4SrcRandom) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4SrcRandom) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4SrcRandom) Clone() (PatternFlowIpv4SrcRandom, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4SrcRandom()
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

// PatternFlowIpv4SrcRandom is ipv4 random pattern
type PatternFlowIpv4SrcRandom interface {
	Validation
	// msg marshals PatternFlowIpv4SrcRandom to protobuf object *otg.PatternFlowIpv4SrcRandom
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4SrcRandom
	// setMsg unmarshals PatternFlowIpv4SrcRandom from protobuf object *otg.PatternFlowIpv4SrcRandom
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4SrcRandom) PatternFlowIpv4SrcRandom
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4SrcRandom
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4SrcRandom
	// validate validates PatternFlowIpv4SrcRandom
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4SrcRandom, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Min returns string, set in PatternFlowIpv4SrcRandom.
	Min() string
	// SetMin assigns string provided by user to PatternFlowIpv4SrcRandom
	SetMin(value string) PatternFlowIpv4SrcRandom
	// HasMin checks if Min has been set in PatternFlowIpv4SrcRandom
	HasMin() bool
	// Max returns string, set in PatternFlowIpv4SrcRandom.
	Max() string
	// SetMax assigns string provided by user to PatternFlowIpv4SrcRandom
	SetMax(value string) PatternFlowIpv4SrcRandom
	// HasMax checks if Max has been set in PatternFlowIpv4SrcRandom
	HasMax() bool
	// Seed returns uint32, set in PatternFlowIpv4SrcRandom.
	Seed() uint32
	// SetSeed assigns uint32 provided by user to PatternFlowIpv4SrcRandom
	SetSeed(value uint32) PatternFlowIpv4SrcRandom
	// HasSeed checks if Seed has been set in PatternFlowIpv4SrcRandom
	HasSeed() bool
	// Count returns uint32, set in PatternFlowIpv4SrcRandom.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4SrcRandom
	SetCount(value uint32) PatternFlowIpv4SrcRandom
	// HasCount checks if Count has been set in PatternFlowIpv4SrcRandom
	HasCount() bool
}

// The minimum possible value generated by the random value generator.
// Min returns a string
func (obj *patternFlowIpv4SrcRandom) Min() string {

	return *obj.obj.Min

}

// The minimum possible value generated by the random value generator.
// Min returns a string
func (obj *patternFlowIpv4SrcRandom) HasMin() bool {
	return obj.obj.Min != nil
}

// The minimum possible value generated by the random value generator.
// SetMin sets the string value in the PatternFlowIpv4SrcRandom object
func (obj *patternFlowIpv4SrcRandom) SetMin(value string) PatternFlowIpv4SrcRandom {

	obj.obj.Min = &value
	return obj
}

// The maximum possible value generated by the random value generator.
// Max returns a string
func (obj *patternFlowIpv4SrcRandom) Max() string {

	return *obj.obj.Max

}

// The maximum possible value generated by the random value generator.
// Max returns a string
func (obj *patternFlowIpv4SrcRandom) HasMax() bool {
	return obj.obj.Max != nil
}

// The maximum possible value generated by the random value generator.
// SetMax sets the string value in the PatternFlowIpv4SrcRandom object
func (obj *patternFlowIpv4SrcRandom) SetMax(value string) PatternFlowIpv4SrcRandom {

	obj.obj.Max = &value
	return obj
}

// The seed value is used to initialize the random number generator to a deterministic state. If the user provides a seed value of 0, the implementation will generate a sequence of non-deterministic random values. For any other seed value, the sequence of random numbers will be generated in a deterministic manner (specific to the implementation).
// Seed returns a uint32
func (obj *patternFlowIpv4SrcRandom) Seed() uint32 {

	return *obj.obj.Seed

}

// The seed value is used to initialize the random number generator to a deterministic state. If the user provides a seed value of 0, the implementation will generate a sequence of non-deterministic random values. For any other seed value, the sequence of random numbers will be generated in a deterministic manner (specific to the implementation).
// Seed returns a uint32
func (obj *patternFlowIpv4SrcRandom) HasSeed() bool {
	return obj.obj.Seed != nil
}

// The seed value is used to initialize the random number generator to a deterministic state. If the user provides a seed value of 0, the implementation will generate a sequence of non-deterministic random values. For any other seed value, the sequence of random numbers will be generated in a deterministic manner (specific to the implementation).
// SetSeed sets the uint32 value in the PatternFlowIpv4SrcRandom object
func (obj *patternFlowIpv4SrcRandom) SetSeed(value uint32) PatternFlowIpv4SrcRandom {

	obj.obj.Seed = &value
	return obj
}

// The total number of values to be generated by the random value generator.
// Count returns a uint32
func (obj *patternFlowIpv4SrcRandom) Count() uint32 {

	return *obj.obj.Count

}

// The total number of values to be generated by the random value generator.
// Count returns a uint32
func (obj *patternFlowIpv4SrcRandom) HasCount() bool {
	return obj.obj.Count != nil
}

// The total number of values to be generated by the random value generator.
// SetCount sets the uint32 value in the PatternFlowIpv4SrcRandom object
func (obj *patternFlowIpv4SrcRandom) SetCount(value uint32) PatternFlowIpv4SrcRandom {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4SrcRandom) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Min != nil {

		err := obj.validateIpv4(obj.Min())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4SrcRandom.Min"))
		}

	}

	if obj.obj.Max != nil {

		err := obj.validateIpv4(obj.Max())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4SrcRandom.Max"))
		}

	}

}

func (obj *patternFlowIpv4SrcRandom) setDefault() {
	if obj.obj.Min == nil {
		obj.SetMin("0.0.0.0")
	}
	if obj.obj.Max == nil {
		obj.SetMax("255.255.255.255")
	}
	if obj.obj.Seed == nil {
		obj.SetSeed(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

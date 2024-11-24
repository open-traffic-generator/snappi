package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4DstRandom *****
type patternFlowIpv4DstRandom struct {
	validation
	obj          *otg.PatternFlowIpv4DstRandom
	marshaller   marshalPatternFlowIpv4DstRandom
	unMarshaller unMarshalPatternFlowIpv4DstRandom
}

func NewPatternFlowIpv4DstRandom() PatternFlowIpv4DstRandom {
	obj := patternFlowIpv4DstRandom{obj: &otg.PatternFlowIpv4DstRandom{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4DstRandom) msg() *otg.PatternFlowIpv4DstRandom {
	return obj.obj
}

func (obj *patternFlowIpv4DstRandom) setMsg(msg *otg.PatternFlowIpv4DstRandom) PatternFlowIpv4DstRandom {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4DstRandom struct {
	obj *patternFlowIpv4DstRandom
}

type marshalPatternFlowIpv4DstRandom interface {
	// ToProto marshals PatternFlowIpv4DstRandom to protobuf object *otg.PatternFlowIpv4DstRandom
	ToProto() (*otg.PatternFlowIpv4DstRandom, error)
	// ToPbText marshals PatternFlowIpv4DstRandom to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4DstRandom to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4DstRandom to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4DstRandom to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4DstRandom struct {
	obj *patternFlowIpv4DstRandom
}

type unMarshalPatternFlowIpv4DstRandom interface {
	// FromProto unmarshals PatternFlowIpv4DstRandom from protobuf object *otg.PatternFlowIpv4DstRandom
	FromProto(msg *otg.PatternFlowIpv4DstRandom) (PatternFlowIpv4DstRandom, error)
	// FromPbText unmarshals PatternFlowIpv4DstRandom from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4DstRandom from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4DstRandom from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4DstRandom) Marshal() marshalPatternFlowIpv4DstRandom {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4DstRandom{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4DstRandom) Unmarshal() unMarshalPatternFlowIpv4DstRandom {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4DstRandom{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4DstRandom) ToProto() (*otg.PatternFlowIpv4DstRandom, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4DstRandom) FromProto(msg *otg.PatternFlowIpv4DstRandom) (PatternFlowIpv4DstRandom, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4DstRandom) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DstRandom) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4DstRandom) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DstRandom) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4DstRandom) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4DstRandom) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DstRandom) FromJson(value string) error {
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

func (obj *patternFlowIpv4DstRandom) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DstRandom) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DstRandom) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4DstRandom) Clone() (PatternFlowIpv4DstRandom, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4DstRandom()
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

// PatternFlowIpv4DstRandom is ipv4 random pattern
type PatternFlowIpv4DstRandom interface {
	Validation
	// msg marshals PatternFlowIpv4DstRandom to protobuf object *otg.PatternFlowIpv4DstRandom
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4DstRandom
	// setMsg unmarshals PatternFlowIpv4DstRandom from protobuf object *otg.PatternFlowIpv4DstRandom
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4DstRandom) PatternFlowIpv4DstRandom
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4DstRandom
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4DstRandom
	// validate validates PatternFlowIpv4DstRandom
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4DstRandom, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Min returns string, set in PatternFlowIpv4DstRandom.
	Min() string
	// SetMin assigns string provided by user to PatternFlowIpv4DstRandom
	SetMin(value string) PatternFlowIpv4DstRandom
	// HasMin checks if Min has been set in PatternFlowIpv4DstRandom
	HasMin() bool
	// Max returns string, set in PatternFlowIpv4DstRandom.
	Max() string
	// SetMax assigns string provided by user to PatternFlowIpv4DstRandom
	SetMax(value string) PatternFlowIpv4DstRandom
	// HasMax checks if Max has been set in PatternFlowIpv4DstRandom
	HasMax() bool
	// Seed returns uint32, set in PatternFlowIpv4DstRandom.
	Seed() uint32
	// SetSeed assigns uint32 provided by user to PatternFlowIpv4DstRandom
	SetSeed(value uint32) PatternFlowIpv4DstRandom
	// HasSeed checks if Seed has been set in PatternFlowIpv4DstRandom
	HasSeed() bool
	// Count returns uint32, set in PatternFlowIpv4DstRandom.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4DstRandom
	SetCount(value uint32) PatternFlowIpv4DstRandom
	// HasCount checks if Count has been set in PatternFlowIpv4DstRandom
	HasCount() bool
}

// The minimum possible value generated by the random value generator.
// Min returns a string
func (obj *patternFlowIpv4DstRandom) Min() string {

	return *obj.obj.Min

}

// The minimum possible value generated by the random value generator.
// Min returns a string
func (obj *patternFlowIpv4DstRandom) HasMin() bool {
	return obj.obj.Min != nil
}

// The minimum possible value generated by the random value generator.
// SetMin sets the string value in the PatternFlowIpv4DstRandom object
func (obj *patternFlowIpv4DstRandom) SetMin(value string) PatternFlowIpv4DstRandom {

	obj.obj.Min = &value
	return obj
}

// The maximum possible value generated by the random value generator.
// Max returns a string
func (obj *patternFlowIpv4DstRandom) Max() string {

	return *obj.obj.Max

}

// The maximum possible value generated by the random value generator.
// Max returns a string
func (obj *patternFlowIpv4DstRandom) HasMax() bool {
	return obj.obj.Max != nil
}

// The maximum possible value generated by the random value generator.
// SetMax sets the string value in the PatternFlowIpv4DstRandom object
func (obj *patternFlowIpv4DstRandom) SetMax(value string) PatternFlowIpv4DstRandom {

	obj.obj.Max = &value
	return obj
}

// The seed value is used to initialize the random number generator to a deterministic state. If the user provides a seed value of 0, the implementation will generate a sequence of non-deterministic random values. For any other seed value, the sequence of random numbers will be generated in a deterministic manner (specific to the implementation).
// Seed returns a uint32
func (obj *patternFlowIpv4DstRandom) Seed() uint32 {

	return *obj.obj.Seed

}

// The seed value is used to initialize the random number generator to a deterministic state. If the user provides a seed value of 0, the implementation will generate a sequence of non-deterministic random values. For any other seed value, the sequence of random numbers will be generated in a deterministic manner (specific to the implementation).
// Seed returns a uint32
func (obj *patternFlowIpv4DstRandom) HasSeed() bool {
	return obj.obj.Seed != nil
}

// The seed value is used to initialize the random number generator to a deterministic state. If the user provides a seed value of 0, the implementation will generate a sequence of non-deterministic random values. For any other seed value, the sequence of random numbers will be generated in a deterministic manner (specific to the implementation).
// SetSeed sets the uint32 value in the PatternFlowIpv4DstRandom object
func (obj *patternFlowIpv4DstRandom) SetSeed(value uint32) PatternFlowIpv4DstRandom {

	obj.obj.Seed = &value
	return obj
}

// The total number of values to be generated by the random value generator.
// Count returns a uint32
func (obj *patternFlowIpv4DstRandom) Count() uint32 {

	return *obj.obj.Count

}

// The total number of values to be generated by the random value generator.
// Count returns a uint32
func (obj *patternFlowIpv4DstRandom) HasCount() bool {
	return obj.obj.Count != nil
}

// The total number of values to be generated by the random value generator.
// SetCount sets the uint32 value in the PatternFlowIpv4DstRandom object
func (obj *patternFlowIpv4DstRandom) SetCount(value uint32) PatternFlowIpv4DstRandom {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4DstRandom) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Min != nil {

		err := obj.validateIpv4(obj.Min())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4DstRandom.Min"))
		}

	}

	if obj.obj.Max != nil {

		err := obj.validateIpv4(obj.Max())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4DstRandom.Max"))
		}

	}

}

func (obj *patternFlowIpv4DstRandom) setDefault() {
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

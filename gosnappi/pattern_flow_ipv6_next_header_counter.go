package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6NextHeaderCounter *****
type patternFlowIpv6NextHeaderCounter struct {
	validation
	obj          *otg.PatternFlowIpv6NextHeaderCounter
	marshaller   marshalPatternFlowIpv6NextHeaderCounter
	unMarshaller unMarshalPatternFlowIpv6NextHeaderCounter
}

func NewPatternFlowIpv6NextHeaderCounter() PatternFlowIpv6NextHeaderCounter {
	obj := patternFlowIpv6NextHeaderCounter{obj: &otg.PatternFlowIpv6NextHeaderCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6NextHeaderCounter) msg() *otg.PatternFlowIpv6NextHeaderCounter {
	return obj.obj
}

func (obj *patternFlowIpv6NextHeaderCounter) setMsg(msg *otg.PatternFlowIpv6NextHeaderCounter) PatternFlowIpv6NextHeaderCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6NextHeaderCounter struct {
	obj *patternFlowIpv6NextHeaderCounter
}

type marshalPatternFlowIpv6NextHeaderCounter interface {
	// ToProto marshals PatternFlowIpv6NextHeaderCounter to protobuf object *otg.PatternFlowIpv6NextHeaderCounter
	ToProto() (*otg.PatternFlowIpv6NextHeaderCounter, error)
	// ToPbText marshals PatternFlowIpv6NextHeaderCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6NextHeaderCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6NextHeaderCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6NextHeaderCounter struct {
	obj *patternFlowIpv6NextHeaderCounter
}

type unMarshalPatternFlowIpv6NextHeaderCounter interface {
	// FromProto unmarshals PatternFlowIpv6NextHeaderCounter from protobuf object *otg.PatternFlowIpv6NextHeaderCounter
	FromProto(msg *otg.PatternFlowIpv6NextHeaderCounter) (PatternFlowIpv6NextHeaderCounter, error)
	// FromPbText unmarshals PatternFlowIpv6NextHeaderCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6NextHeaderCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6NextHeaderCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6NextHeaderCounter) Marshal() marshalPatternFlowIpv6NextHeaderCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6NextHeaderCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6NextHeaderCounter) Unmarshal() unMarshalPatternFlowIpv6NextHeaderCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6NextHeaderCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6NextHeaderCounter) ToProto() (*otg.PatternFlowIpv6NextHeaderCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6NextHeaderCounter) FromProto(msg *otg.PatternFlowIpv6NextHeaderCounter) (PatternFlowIpv6NextHeaderCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6NextHeaderCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6NextHeaderCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6NextHeaderCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6NextHeaderCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6NextHeaderCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6NextHeaderCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6NextHeaderCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6NextHeaderCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6NextHeaderCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6NextHeaderCounter) Clone() (PatternFlowIpv6NextHeaderCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6NextHeaderCounter()
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

// PatternFlowIpv6NextHeaderCounter is integer counter pattern
type PatternFlowIpv6NextHeaderCounter interface {
	Validation
	// msg marshals PatternFlowIpv6NextHeaderCounter to protobuf object *otg.PatternFlowIpv6NextHeaderCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6NextHeaderCounter
	// setMsg unmarshals PatternFlowIpv6NextHeaderCounter from protobuf object *otg.PatternFlowIpv6NextHeaderCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6NextHeaderCounter) PatternFlowIpv6NextHeaderCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6NextHeaderCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6NextHeaderCounter
	// validate validates PatternFlowIpv6NextHeaderCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6NextHeaderCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6NextHeaderCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6NextHeaderCounter
	SetStart(value uint32) PatternFlowIpv6NextHeaderCounter
	// HasStart checks if Start has been set in PatternFlowIpv6NextHeaderCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6NextHeaderCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6NextHeaderCounter
	SetStep(value uint32) PatternFlowIpv6NextHeaderCounter
	// HasStep checks if Step has been set in PatternFlowIpv6NextHeaderCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6NextHeaderCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6NextHeaderCounter
	SetCount(value uint32) PatternFlowIpv6NextHeaderCounter
	// HasCount checks if Count has been set in PatternFlowIpv6NextHeaderCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6NextHeaderCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6NextHeaderCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6NextHeaderCounter object
func (obj *patternFlowIpv6NextHeaderCounter) SetStart(value uint32) PatternFlowIpv6NextHeaderCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6NextHeaderCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6NextHeaderCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6NextHeaderCounter object
func (obj *patternFlowIpv6NextHeaderCounter) SetStep(value uint32) PatternFlowIpv6NextHeaderCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6NextHeaderCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6NextHeaderCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6NextHeaderCounter object
func (obj *patternFlowIpv6NextHeaderCounter) SetCount(value uint32) PatternFlowIpv6NextHeaderCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6NextHeaderCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6NextHeaderCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6NextHeaderCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6NextHeaderCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6NextHeaderCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(59)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

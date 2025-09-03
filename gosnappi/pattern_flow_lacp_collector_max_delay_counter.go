package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpCollectorMaxDelayCounter *****
type patternFlowLacpCollectorMaxDelayCounter struct {
	validation
	obj          *otg.PatternFlowLacpCollectorMaxDelayCounter
	marshaller   marshalPatternFlowLacpCollectorMaxDelayCounter
	unMarshaller unMarshalPatternFlowLacpCollectorMaxDelayCounter
}

func NewPatternFlowLacpCollectorMaxDelayCounter() PatternFlowLacpCollectorMaxDelayCounter {
	obj := patternFlowLacpCollectorMaxDelayCounter{obj: &otg.PatternFlowLacpCollectorMaxDelayCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpCollectorMaxDelayCounter) msg() *otg.PatternFlowLacpCollectorMaxDelayCounter {
	return obj.obj
}

func (obj *patternFlowLacpCollectorMaxDelayCounter) setMsg(msg *otg.PatternFlowLacpCollectorMaxDelayCounter) PatternFlowLacpCollectorMaxDelayCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpCollectorMaxDelayCounter struct {
	obj *patternFlowLacpCollectorMaxDelayCounter
}

type marshalPatternFlowLacpCollectorMaxDelayCounter interface {
	// ToProto marshals PatternFlowLacpCollectorMaxDelayCounter to protobuf object *otg.PatternFlowLacpCollectorMaxDelayCounter
	ToProto() (*otg.PatternFlowLacpCollectorMaxDelayCounter, error)
	// ToPbText marshals PatternFlowLacpCollectorMaxDelayCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpCollectorMaxDelayCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpCollectorMaxDelayCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpCollectorMaxDelayCounter struct {
	obj *patternFlowLacpCollectorMaxDelayCounter
}

type unMarshalPatternFlowLacpCollectorMaxDelayCounter interface {
	// FromProto unmarshals PatternFlowLacpCollectorMaxDelayCounter from protobuf object *otg.PatternFlowLacpCollectorMaxDelayCounter
	FromProto(msg *otg.PatternFlowLacpCollectorMaxDelayCounter) (PatternFlowLacpCollectorMaxDelayCounter, error)
	// FromPbText unmarshals PatternFlowLacpCollectorMaxDelayCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpCollectorMaxDelayCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpCollectorMaxDelayCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpCollectorMaxDelayCounter) Marshal() marshalPatternFlowLacpCollectorMaxDelayCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpCollectorMaxDelayCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpCollectorMaxDelayCounter) Unmarshal() unMarshalPatternFlowLacpCollectorMaxDelayCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpCollectorMaxDelayCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpCollectorMaxDelayCounter) ToProto() (*otg.PatternFlowLacpCollectorMaxDelayCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpCollectorMaxDelayCounter) FromProto(msg *otg.PatternFlowLacpCollectorMaxDelayCounter) (PatternFlowLacpCollectorMaxDelayCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpCollectorMaxDelayCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpCollectorMaxDelayCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpCollectorMaxDelayCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpCollectorMaxDelayCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpCollectorMaxDelayCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpCollectorMaxDelayCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpCollectorMaxDelayCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpCollectorMaxDelayCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpCollectorMaxDelayCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpCollectorMaxDelayCounter) Clone() (PatternFlowLacpCollectorMaxDelayCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpCollectorMaxDelayCounter()
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

// PatternFlowLacpCollectorMaxDelayCounter is integer counter pattern
type PatternFlowLacpCollectorMaxDelayCounter interface {
	Validation
	// msg marshals PatternFlowLacpCollectorMaxDelayCounter to protobuf object *otg.PatternFlowLacpCollectorMaxDelayCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpCollectorMaxDelayCounter
	// setMsg unmarshals PatternFlowLacpCollectorMaxDelayCounter from protobuf object *otg.PatternFlowLacpCollectorMaxDelayCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpCollectorMaxDelayCounter) PatternFlowLacpCollectorMaxDelayCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpCollectorMaxDelayCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpCollectorMaxDelayCounter
	// validate validates PatternFlowLacpCollectorMaxDelayCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpCollectorMaxDelayCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpCollectorMaxDelayCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpCollectorMaxDelayCounter
	SetStart(value uint32) PatternFlowLacpCollectorMaxDelayCounter
	// HasStart checks if Start has been set in PatternFlowLacpCollectorMaxDelayCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpCollectorMaxDelayCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpCollectorMaxDelayCounter
	SetStep(value uint32) PatternFlowLacpCollectorMaxDelayCounter
	// HasStep checks if Step has been set in PatternFlowLacpCollectorMaxDelayCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpCollectorMaxDelayCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpCollectorMaxDelayCounter
	SetCount(value uint32) PatternFlowLacpCollectorMaxDelayCounter
	// HasCount checks if Count has been set in PatternFlowLacpCollectorMaxDelayCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpCollectorMaxDelayCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpCollectorMaxDelayCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpCollectorMaxDelayCounter object
func (obj *patternFlowLacpCollectorMaxDelayCounter) SetStart(value uint32) PatternFlowLacpCollectorMaxDelayCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpCollectorMaxDelayCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpCollectorMaxDelayCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpCollectorMaxDelayCounter object
func (obj *patternFlowLacpCollectorMaxDelayCounter) SetStep(value uint32) PatternFlowLacpCollectorMaxDelayCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpCollectorMaxDelayCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpCollectorMaxDelayCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpCollectorMaxDelayCounter object
func (obj *patternFlowLacpCollectorMaxDelayCounter) SetCount(value uint32) PatternFlowLacpCollectorMaxDelayCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpCollectorMaxDelayCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpCollectorMaxDelayCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpCollectorMaxDelayCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpCollectorMaxDelayCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpCollectorMaxDelayCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(0)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

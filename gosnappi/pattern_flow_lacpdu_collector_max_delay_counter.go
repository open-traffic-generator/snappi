package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduCollectorMaxDelayCounter *****
type patternFlowLacpduCollectorMaxDelayCounter struct {
	validation
	obj          *otg.PatternFlowLacpduCollectorMaxDelayCounter
	marshaller   marshalPatternFlowLacpduCollectorMaxDelayCounter
	unMarshaller unMarshalPatternFlowLacpduCollectorMaxDelayCounter
}

func NewPatternFlowLacpduCollectorMaxDelayCounter() PatternFlowLacpduCollectorMaxDelayCounter {
	obj := patternFlowLacpduCollectorMaxDelayCounter{obj: &otg.PatternFlowLacpduCollectorMaxDelayCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduCollectorMaxDelayCounter) msg() *otg.PatternFlowLacpduCollectorMaxDelayCounter {
	return obj.obj
}

func (obj *patternFlowLacpduCollectorMaxDelayCounter) setMsg(msg *otg.PatternFlowLacpduCollectorMaxDelayCounter) PatternFlowLacpduCollectorMaxDelayCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduCollectorMaxDelayCounter struct {
	obj *patternFlowLacpduCollectorMaxDelayCounter
}

type marshalPatternFlowLacpduCollectorMaxDelayCounter interface {
	// ToProto marshals PatternFlowLacpduCollectorMaxDelayCounter to protobuf object *otg.PatternFlowLacpduCollectorMaxDelayCounter
	ToProto() (*otg.PatternFlowLacpduCollectorMaxDelayCounter, error)
	// ToPbText marshals PatternFlowLacpduCollectorMaxDelayCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduCollectorMaxDelayCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduCollectorMaxDelayCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduCollectorMaxDelayCounter struct {
	obj *patternFlowLacpduCollectorMaxDelayCounter
}

type unMarshalPatternFlowLacpduCollectorMaxDelayCounter interface {
	// FromProto unmarshals PatternFlowLacpduCollectorMaxDelayCounter from protobuf object *otg.PatternFlowLacpduCollectorMaxDelayCounter
	FromProto(msg *otg.PatternFlowLacpduCollectorMaxDelayCounter) (PatternFlowLacpduCollectorMaxDelayCounter, error)
	// FromPbText unmarshals PatternFlowLacpduCollectorMaxDelayCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduCollectorMaxDelayCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduCollectorMaxDelayCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduCollectorMaxDelayCounter) Marshal() marshalPatternFlowLacpduCollectorMaxDelayCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduCollectorMaxDelayCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduCollectorMaxDelayCounter) Unmarshal() unMarshalPatternFlowLacpduCollectorMaxDelayCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduCollectorMaxDelayCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduCollectorMaxDelayCounter) ToProto() (*otg.PatternFlowLacpduCollectorMaxDelayCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduCollectorMaxDelayCounter) FromProto(msg *otg.PatternFlowLacpduCollectorMaxDelayCounter) (PatternFlowLacpduCollectorMaxDelayCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduCollectorMaxDelayCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduCollectorMaxDelayCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduCollectorMaxDelayCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduCollectorMaxDelayCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduCollectorMaxDelayCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduCollectorMaxDelayCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduCollectorMaxDelayCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduCollectorMaxDelayCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduCollectorMaxDelayCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduCollectorMaxDelayCounter) Clone() (PatternFlowLacpduCollectorMaxDelayCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduCollectorMaxDelayCounter()
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

// PatternFlowLacpduCollectorMaxDelayCounter is integer counter pattern
type PatternFlowLacpduCollectorMaxDelayCounter interface {
	Validation
	// msg marshals PatternFlowLacpduCollectorMaxDelayCounter to protobuf object *otg.PatternFlowLacpduCollectorMaxDelayCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduCollectorMaxDelayCounter
	// setMsg unmarshals PatternFlowLacpduCollectorMaxDelayCounter from protobuf object *otg.PatternFlowLacpduCollectorMaxDelayCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduCollectorMaxDelayCounter) PatternFlowLacpduCollectorMaxDelayCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduCollectorMaxDelayCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduCollectorMaxDelayCounter
	// validate validates PatternFlowLacpduCollectorMaxDelayCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduCollectorMaxDelayCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduCollectorMaxDelayCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduCollectorMaxDelayCounter
	SetStart(value uint32) PatternFlowLacpduCollectorMaxDelayCounter
	// HasStart checks if Start has been set in PatternFlowLacpduCollectorMaxDelayCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduCollectorMaxDelayCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduCollectorMaxDelayCounter
	SetStep(value uint32) PatternFlowLacpduCollectorMaxDelayCounter
	// HasStep checks if Step has been set in PatternFlowLacpduCollectorMaxDelayCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduCollectorMaxDelayCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduCollectorMaxDelayCounter
	SetCount(value uint32) PatternFlowLacpduCollectorMaxDelayCounter
	// HasCount checks if Count has been set in PatternFlowLacpduCollectorMaxDelayCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduCollectorMaxDelayCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduCollectorMaxDelayCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduCollectorMaxDelayCounter object
func (obj *patternFlowLacpduCollectorMaxDelayCounter) SetStart(value uint32) PatternFlowLacpduCollectorMaxDelayCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduCollectorMaxDelayCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduCollectorMaxDelayCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduCollectorMaxDelayCounter object
func (obj *patternFlowLacpduCollectorMaxDelayCounter) SetStep(value uint32) PatternFlowLacpduCollectorMaxDelayCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduCollectorMaxDelayCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduCollectorMaxDelayCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduCollectorMaxDelayCounter object
func (obj *patternFlowLacpduCollectorMaxDelayCounter) SetCount(value uint32) PatternFlowLacpduCollectorMaxDelayCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduCollectorMaxDelayCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduCollectorMaxDelayCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduCollectorMaxDelayCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduCollectorMaxDelayCounter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduCollectorMaxDelayCounter) setDefault() {
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

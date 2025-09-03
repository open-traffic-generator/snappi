package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorTypeCounter *****
type patternFlowLacpActorTypeCounter struct {
	validation
	obj          *otg.PatternFlowLacpActorTypeCounter
	marshaller   marshalPatternFlowLacpActorTypeCounter
	unMarshaller unMarshalPatternFlowLacpActorTypeCounter
}

func NewPatternFlowLacpActorTypeCounter() PatternFlowLacpActorTypeCounter {
	obj := patternFlowLacpActorTypeCounter{obj: &otg.PatternFlowLacpActorTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorTypeCounter) msg() *otg.PatternFlowLacpActorTypeCounter {
	return obj.obj
}

func (obj *patternFlowLacpActorTypeCounter) setMsg(msg *otg.PatternFlowLacpActorTypeCounter) PatternFlowLacpActorTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorTypeCounter struct {
	obj *patternFlowLacpActorTypeCounter
}

type marshalPatternFlowLacpActorTypeCounter interface {
	// ToProto marshals PatternFlowLacpActorTypeCounter to protobuf object *otg.PatternFlowLacpActorTypeCounter
	ToProto() (*otg.PatternFlowLacpActorTypeCounter, error)
	// ToPbText marshals PatternFlowLacpActorTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorTypeCounter struct {
	obj *patternFlowLacpActorTypeCounter
}

type unMarshalPatternFlowLacpActorTypeCounter interface {
	// FromProto unmarshals PatternFlowLacpActorTypeCounter from protobuf object *otg.PatternFlowLacpActorTypeCounter
	FromProto(msg *otg.PatternFlowLacpActorTypeCounter) (PatternFlowLacpActorTypeCounter, error)
	// FromPbText unmarshals PatternFlowLacpActorTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorTypeCounter) Marshal() marshalPatternFlowLacpActorTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorTypeCounter) Unmarshal() unMarshalPatternFlowLacpActorTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorTypeCounter) ToProto() (*otg.PatternFlowLacpActorTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorTypeCounter) FromProto(msg *otg.PatternFlowLacpActorTypeCounter) (PatternFlowLacpActorTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpActorTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorTypeCounter) Clone() (PatternFlowLacpActorTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorTypeCounter()
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

// PatternFlowLacpActorTypeCounter is integer counter pattern
type PatternFlowLacpActorTypeCounter interface {
	Validation
	// msg marshals PatternFlowLacpActorTypeCounter to protobuf object *otg.PatternFlowLacpActorTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorTypeCounter
	// setMsg unmarshals PatternFlowLacpActorTypeCounter from protobuf object *otg.PatternFlowLacpActorTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorTypeCounter) PatternFlowLacpActorTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorTypeCounter
	// validate validates PatternFlowLacpActorTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpActorTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpActorTypeCounter
	SetStart(value uint32) PatternFlowLacpActorTypeCounter
	// HasStart checks if Start has been set in PatternFlowLacpActorTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpActorTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpActorTypeCounter
	SetStep(value uint32) PatternFlowLacpActorTypeCounter
	// HasStep checks if Step has been set in PatternFlowLacpActorTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpActorTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpActorTypeCounter
	SetCount(value uint32) PatternFlowLacpActorTypeCounter
	// HasCount checks if Count has been set in PatternFlowLacpActorTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpActorTypeCounter object
func (obj *patternFlowLacpActorTypeCounter) SetStart(value uint32) PatternFlowLacpActorTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpActorTypeCounter object
func (obj *patternFlowLacpActorTypeCounter) SetStep(value uint32) PatternFlowLacpActorTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpActorTypeCounter object
func (obj *patternFlowLacpActorTypeCounter) SetCount(value uint32) PatternFlowLacpActorTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpActorTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorTypeCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpActorTypeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(1)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

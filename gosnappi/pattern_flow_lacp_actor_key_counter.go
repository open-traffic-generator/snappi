package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorKeyCounter *****
type patternFlowLacpActorKeyCounter struct {
	validation
	obj          *otg.PatternFlowLacpActorKeyCounter
	marshaller   marshalPatternFlowLacpActorKeyCounter
	unMarshaller unMarshalPatternFlowLacpActorKeyCounter
}

func NewPatternFlowLacpActorKeyCounter() PatternFlowLacpActorKeyCounter {
	obj := patternFlowLacpActorKeyCounter{obj: &otg.PatternFlowLacpActorKeyCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorKeyCounter) msg() *otg.PatternFlowLacpActorKeyCounter {
	return obj.obj
}

func (obj *patternFlowLacpActorKeyCounter) setMsg(msg *otg.PatternFlowLacpActorKeyCounter) PatternFlowLacpActorKeyCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorKeyCounter struct {
	obj *patternFlowLacpActorKeyCounter
}

type marshalPatternFlowLacpActorKeyCounter interface {
	// ToProto marshals PatternFlowLacpActorKeyCounter to protobuf object *otg.PatternFlowLacpActorKeyCounter
	ToProto() (*otg.PatternFlowLacpActorKeyCounter, error)
	// ToPbText marshals PatternFlowLacpActorKeyCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorKeyCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorKeyCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorKeyCounter struct {
	obj *patternFlowLacpActorKeyCounter
}

type unMarshalPatternFlowLacpActorKeyCounter interface {
	// FromProto unmarshals PatternFlowLacpActorKeyCounter from protobuf object *otg.PatternFlowLacpActorKeyCounter
	FromProto(msg *otg.PatternFlowLacpActorKeyCounter) (PatternFlowLacpActorKeyCounter, error)
	// FromPbText unmarshals PatternFlowLacpActorKeyCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorKeyCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorKeyCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorKeyCounter) Marshal() marshalPatternFlowLacpActorKeyCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorKeyCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorKeyCounter) Unmarshal() unMarshalPatternFlowLacpActorKeyCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorKeyCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorKeyCounter) ToProto() (*otg.PatternFlowLacpActorKeyCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorKeyCounter) FromProto(msg *otg.PatternFlowLacpActorKeyCounter) (PatternFlowLacpActorKeyCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorKeyCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorKeyCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorKeyCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorKeyCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorKeyCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorKeyCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpActorKeyCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorKeyCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorKeyCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorKeyCounter) Clone() (PatternFlowLacpActorKeyCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorKeyCounter()
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

// PatternFlowLacpActorKeyCounter is integer counter pattern
type PatternFlowLacpActorKeyCounter interface {
	Validation
	// msg marshals PatternFlowLacpActorKeyCounter to protobuf object *otg.PatternFlowLacpActorKeyCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorKeyCounter
	// setMsg unmarshals PatternFlowLacpActorKeyCounter from protobuf object *otg.PatternFlowLacpActorKeyCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorKeyCounter) PatternFlowLacpActorKeyCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorKeyCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorKeyCounter
	// validate validates PatternFlowLacpActorKeyCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorKeyCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpActorKeyCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpActorKeyCounter
	SetStart(value uint32) PatternFlowLacpActorKeyCounter
	// HasStart checks if Start has been set in PatternFlowLacpActorKeyCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpActorKeyCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpActorKeyCounter
	SetStep(value uint32) PatternFlowLacpActorKeyCounter
	// HasStep checks if Step has been set in PatternFlowLacpActorKeyCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpActorKeyCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpActorKeyCounter
	SetCount(value uint32) PatternFlowLacpActorKeyCounter
	// HasCount checks if Count has been set in PatternFlowLacpActorKeyCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorKeyCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorKeyCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpActorKeyCounter object
func (obj *patternFlowLacpActorKeyCounter) SetStart(value uint32) PatternFlowLacpActorKeyCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorKeyCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorKeyCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpActorKeyCounter object
func (obj *patternFlowLacpActorKeyCounter) SetStep(value uint32) PatternFlowLacpActorKeyCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorKeyCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorKeyCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpActorKeyCounter object
func (obj *patternFlowLacpActorKeyCounter) SetCount(value uint32) PatternFlowLacpActorKeyCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpActorKeyCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorKeyCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorKeyCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorKeyCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpActorKeyCounter) setDefault() {
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

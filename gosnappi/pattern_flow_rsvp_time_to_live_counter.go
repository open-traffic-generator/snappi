package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRsvpTimeToLiveCounter *****
type patternFlowRsvpTimeToLiveCounter struct {
	validation
	obj          *otg.PatternFlowRsvpTimeToLiveCounter
	marshaller   marshalPatternFlowRsvpTimeToLiveCounter
	unMarshaller unMarshalPatternFlowRsvpTimeToLiveCounter
}

func NewPatternFlowRsvpTimeToLiveCounter() PatternFlowRsvpTimeToLiveCounter {
	obj := patternFlowRsvpTimeToLiveCounter{obj: &otg.PatternFlowRsvpTimeToLiveCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRsvpTimeToLiveCounter) msg() *otg.PatternFlowRsvpTimeToLiveCounter {
	return obj.obj
}

func (obj *patternFlowRsvpTimeToLiveCounter) setMsg(msg *otg.PatternFlowRsvpTimeToLiveCounter) PatternFlowRsvpTimeToLiveCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRsvpTimeToLiveCounter struct {
	obj *patternFlowRsvpTimeToLiveCounter
}

type marshalPatternFlowRsvpTimeToLiveCounter interface {
	// ToProto marshals PatternFlowRsvpTimeToLiveCounter to protobuf object *otg.PatternFlowRsvpTimeToLiveCounter
	ToProto() (*otg.PatternFlowRsvpTimeToLiveCounter, error)
	// ToPbText marshals PatternFlowRsvpTimeToLiveCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRsvpTimeToLiveCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRsvpTimeToLiveCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRsvpTimeToLiveCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRsvpTimeToLiveCounter struct {
	obj *patternFlowRsvpTimeToLiveCounter
}

type unMarshalPatternFlowRsvpTimeToLiveCounter interface {
	// FromProto unmarshals PatternFlowRsvpTimeToLiveCounter from protobuf object *otg.PatternFlowRsvpTimeToLiveCounter
	FromProto(msg *otg.PatternFlowRsvpTimeToLiveCounter) (PatternFlowRsvpTimeToLiveCounter, error)
	// FromPbText unmarshals PatternFlowRsvpTimeToLiveCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRsvpTimeToLiveCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRsvpTimeToLiveCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRsvpTimeToLiveCounter) Marshal() marshalPatternFlowRsvpTimeToLiveCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRsvpTimeToLiveCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRsvpTimeToLiveCounter) Unmarshal() unMarshalPatternFlowRsvpTimeToLiveCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRsvpTimeToLiveCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRsvpTimeToLiveCounter) ToProto() (*otg.PatternFlowRsvpTimeToLiveCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRsvpTimeToLiveCounter) FromProto(msg *otg.PatternFlowRsvpTimeToLiveCounter) (PatternFlowRsvpTimeToLiveCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRsvpTimeToLiveCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRsvpTimeToLiveCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRsvpTimeToLiveCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRsvpTimeToLiveCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRsvpTimeToLiveCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowRsvpTimeToLiveCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRsvpTimeToLiveCounter) FromJson(value string) error {
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

func (obj *patternFlowRsvpTimeToLiveCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRsvpTimeToLiveCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRsvpTimeToLiveCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRsvpTimeToLiveCounter) Clone() (PatternFlowRsvpTimeToLiveCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRsvpTimeToLiveCounter()
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

// PatternFlowRsvpTimeToLiveCounter is integer counter pattern
type PatternFlowRsvpTimeToLiveCounter interface {
	Validation
	// msg marshals PatternFlowRsvpTimeToLiveCounter to protobuf object *otg.PatternFlowRsvpTimeToLiveCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRsvpTimeToLiveCounter
	// setMsg unmarshals PatternFlowRsvpTimeToLiveCounter from protobuf object *otg.PatternFlowRsvpTimeToLiveCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRsvpTimeToLiveCounter) PatternFlowRsvpTimeToLiveCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRsvpTimeToLiveCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRsvpTimeToLiveCounter
	// validate validates PatternFlowRsvpTimeToLiveCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRsvpTimeToLiveCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRsvpTimeToLiveCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRsvpTimeToLiveCounter
	SetStart(value uint32) PatternFlowRsvpTimeToLiveCounter
	// HasStart checks if Start has been set in PatternFlowRsvpTimeToLiveCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRsvpTimeToLiveCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRsvpTimeToLiveCounter
	SetStep(value uint32) PatternFlowRsvpTimeToLiveCounter
	// HasStep checks if Step has been set in PatternFlowRsvpTimeToLiveCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRsvpTimeToLiveCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRsvpTimeToLiveCounter
	SetCount(value uint32) PatternFlowRsvpTimeToLiveCounter
	// HasCount checks if Count has been set in PatternFlowRsvpTimeToLiveCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRsvpTimeToLiveCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRsvpTimeToLiveCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRsvpTimeToLiveCounter object
func (obj *patternFlowRsvpTimeToLiveCounter) SetStart(value uint32) PatternFlowRsvpTimeToLiveCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRsvpTimeToLiveCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRsvpTimeToLiveCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRsvpTimeToLiveCounter object
func (obj *patternFlowRsvpTimeToLiveCounter) SetStep(value uint32) PatternFlowRsvpTimeToLiveCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRsvpTimeToLiveCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRsvpTimeToLiveCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRsvpTimeToLiveCounter object
func (obj *patternFlowRsvpTimeToLiveCounter) SetCount(value uint32) PatternFlowRsvpTimeToLiveCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRsvpTimeToLiveCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRsvpTimeToLiveCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRsvpTimeToLiveCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRsvpTimeToLiveCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRsvpTimeToLiveCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(64)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

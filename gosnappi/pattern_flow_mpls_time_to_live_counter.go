package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowMplsTimeToLiveCounter *****
type patternFlowMplsTimeToLiveCounter struct {
	validation
	obj          *otg.PatternFlowMplsTimeToLiveCounter
	marshaller   marshalPatternFlowMplsTimeToLiveCounter
	unMarshaller unMarshalPatternFlowMplsTimeToLiveCounter
}

func NewPatternFlowMplsTimeToLiveCounter() PatternFlowMplsTimeToLiveCounter {
	obj := patternFlowMplsTimeToLiveCounter{obj: &otg.PatternFlowMplsTimeToLiveCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowMplsTimeToLiveCounter) msg() *otg.PatternFlowMplsTimeToLiveCounter {
	return obj.obj
}

func (obj *patternFlowMplsTimeToLiveCounter) setMsg(msg *otg.PatternFlowMplsTimeToLiveCounter) PatternFlowMplsTimeToLiveCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowMplsTimeToLiveCounter struct {
	obj *patternFlowMplsTimeToLiveCounter
}

type marshalPatternFlowMplsTimeToLiveCounter interface {
	// ToProto marshals PatternFlowMplsTimeToLiveCounter to protobuf object *otg.PatternFlowMplsTimeToLiveCounter
	ToProto() (*otg.PatternFlowMplsTimeToLiveCounter, error)
	// ToPbText marshals PatternFlowMplsTimeToLiveCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowMplsTimeToLiveCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowMplsTimeToLiveCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowMplsTimeToLiveCounter struct {
	obj *patternFlowMplsTimeToLiveCounter
}

type unMarshalPatternFlowMplsTimeToLiveCounter interface {
	// FromProto unmarshals PatternFlowMplsTimeToLiveCounter from protobuf object *otg.PatternFlowMplsTimeToLiveCounter
	FromProto(msg *otg.PatternFlowMplsTimeToLiveCounter) (PatternFlowMplsTimeToLiveCounter, error)
	// FromPbText unmarshals PatternFlowMplsTimeToLiveCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowMplsTimeToLiveCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowMplsTimeToLiveCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowMplsTimeToLiveCounter) Marshal() marshalPatternFlowMplsTimeToLiveCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowMplsTimeToLiveCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowMplsTimeToLiveCounter) Unmarshal() unMarshalPatternFlowMplsTimeToLiveCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowMplsTimeToLiveCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowMplsTimeToLiveCounter) ToProto() (*otg.PatternFlowMplsTimeToLiveCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowMplsTimeToLiveCounter) FromProto(msg *otg.PatternFlowMplsTimeToLiveCounter) (PatternFlowMplsTimeToLiveCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowMplsTimeToLiveCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowMplsTimeToLiveCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowMplsTimeToLiveCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowMplsTimeToLiveCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowMplsTimeToLiveCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowMplsTimeToLiveCounter) FromJson(value string) error {
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

func (obj *patternFlowMplsTimeToLiveCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowMplsTimeToLiveCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowMplsTimeToLiveCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowMplsTimeToLiveCounter) Clone() (PatternFlowMplsTimeToLiveCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowMplsTimeToLiveCounter()
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

// PatternFlowMplsTimeToLiveCounter is integer counter pattern
type PatternFlowMplsTimeToLiveCounter interface {
	Validation
	// msg marshals PatternFlowMplsTimeToLiveCounter to protobuf object *otg.PatternFlowMplsTimeToLiveCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowMplsTimeToLiveCounter
	// setMsg unmarshals PatternFlowMplsTimeToLiveCounter from protobuf object *otg.PatternFlowMplsTimeToLiveCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowMplsTimeToLiveCounter) PatternFlowMplsTimeToLiveCounter
	// provides marshal interface
	Marshal() marshalPatternFlowMplsTimeToLiveCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowMplsTimeToLiveCounter
	// validate validates PatternFlowMplsTimeToLiveCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowMplsTimeToLiveCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowMplsTimeToLiveCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowMplsTimeToLiveCounter
	SetStart(value uint32) PatternFlowMplsTimeToLiveCounter
	// HasStart checks if Start has been set in PatternFlowMplsTimeToLiveCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowMplsTimeToLiveCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowMplsTimeToLiveCounter
	SetStep(value uint32) PatternFlowMplsTimeToLiveCounter
	// HasStep checks if Step has been set in PatternFlowMplsTimeToLiveCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowMplsTimeToLiveCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowMplsTimeToLiveCounter
	SetCount(value uint32) PatternFlowMplsTimeToLiveCounter
	// HasCount checks if Count has been set in PatternFlowMplsTimeToLiveCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowMplsTimeToLiveCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowMplsTimeToLiveCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowMplsTimeToLiveCounter object
func (obj *patternFlowMplsTimeToLiveCounter) SetStart(value uint32) PatternFlowMplsTimeToLiveCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowMplsTimeToLiveCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowMplsTimeToLiveCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowMplsTimeToLiveCounter object
func (obj *patternFlowMplsTimeToLiveCounter) SetStep(value uint32) PatternFlowMplsTimeToLiveCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowMplsTimeToLiveCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowMplsTimeToLiveCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowMplsTimeToLiveCounter object
func (obj *patternFlowMplsTimeToLiveCounter) SetCount(value uint32) PatternFlowMplsTimeToLiveCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowMplsTimeToLiveCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsTimeToLiveCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsTimeToLiveCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsTimeToLiveCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowMplsTimeToLiveCounter) setDefault() {
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

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowMplsLabelCounter *****
type patternFlowMplsLabelCounter struct {
	validation
	obj          *otg.PatternFlowMplsLabelCounter
	marshaller   marshalPatternFlowMplsLabelCounter
	unMarshaller unMarshalPatternFlowMplsLabelCounter
}

func NewPatternFlowMplsLabelCounter() PatternFlowMplsLabelCounter {
	obj := patternFlowMplsLabelCounter{obj: &otg.PatternFlowMplsLabelCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowMplsLabelCounter) msg() *otg.PatternFlowMplsLabelCounter {
	return obj.obj
}

func (obj *patternFlowMplsLabelCounter) setMsg(msg *otg.PatternFlowMplsLabelCounter) PatternFlowMplsLabelCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowMplsLabelCounter struct {
	obj *patternFlowMplsLabelCounter
}

type marshalPatternFlowMplsLabelCounter interface {
	// ToProto marshals PatternFlowMplsLabelCounter to protobuf object *otg.PatternFlowMplsLabelCounter
	ToProto() (*otg.PatternFlowMplsLabelCounter, error)
	// ToPbText marshals PatternFlowMplsLabelCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowMplsLabelCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowMplsLabelCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowMplsLabelCounter struct {
	obj *patternFlowMplsLabelCounter
}

type unMarshalPatternFlowMplsLabelCounter interface {
	// FromProto unmarshals PatternFlowMplsLabelCounter from protobuf object *otg.PatternFlowMplsLabelCounter
	FromProto(msg *otg.PatternFlowMplsLabelCounter) (PatternFlowMplsLabelCounter, error)
	// FromPbText unmarshals PatternFlowMplsLabelCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowMplsLabelCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowMplsLabelCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowMplsLabelCounter) Marshal() marshalPatternFlowMplsLabelCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowMplsLabelCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowMplsLabelCounter) Unmarshal() unMarshalPatternFlowMplsLabelCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowMplsLabelCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowMplsLabelCounter) ToProto() (*otg.PatternFlowMplsLabelCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowMplsLabelCounter) FromProto(msg *otg.PatternFlowMplsLabelCounter) (PatternFlowMplsLabelCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowMplsLabelCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowMplsLabelCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowMplsLabelCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowMplsLabelCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowMplsLabelCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowMplsLabelCounter) FromJson(value string) error {
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

func (obj *patternFlowMplsLabelCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowMplsLabelCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowMplsLabelCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowMplsLabelCounter) Clone() (PatternFlowMplsLabelCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowMplsLabelCounter()
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

// PatternFlowMplsLabelCounter is integer counter pattern
type PatternFlowMplsLabelCounter interface {
	Validation
	// msg marshals PatternFlowMplsLabelCounter to protobuf object *otg.PatternFlowMplsLabelCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowMplsLabelCounter
	// setMsg unmarshals PatternFlowMplsLabelCounter from protobuf object *otg.PatternFlowMplsLabelCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowMplsLabelCounter) PatternFlowMplsLabelCounter
	// provides marshal interface
	Marshal() marshalPatternFlowMplsLabelCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowMplsLabelCounter
	// validate validates PatternFlowMplsLabelCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowMplsLabelCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowMplsLabelCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowMplsLabelCounter
	SetStart(value uint32) PatternFlowMplsLabelCounter
	// HasStart checks if Start has been set in PatternFlowMplsLabelCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowMplsLabelCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowMplsLabelCounter
	SetStep(value uint32) PatternFlowMplsLabelCounter
	// HasStep checks if Step has been set in PatternFlowMplsLabelCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowMplsLabelCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowMplsLabelCounter
	SetCount(value uint32) PatternFlowMplsLabelCounter
	// HasCount checks if Count has been set in PatternFlowMplsLabelCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowMplsLabelCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowMplsLabelCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowMplsLabelCounter object
func (obj *patternFlowMplsLabelCounter) SetStart(value uint32) PatternFlowMplsLabelCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowMplsLabelCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowMplsLabelCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowMplsLabelCounter object
func (obj *patternFlowMplsLabelCounter) SetStep(value uint32) PatternFlowMplsLabelCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowMplsLabelCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowMplsLabelCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowMplsLabelCounter object
func (obj *patternFlowMplsLabelCounter) SetCount(value uint32) PatternFlowMplsLabelCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowMplsLabelCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsLabelCounter.Start <= 1048575 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsLabelCounter.Step <= 1048575 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1048576 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsLabelCounter.Count <= 1048576 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowMplsLabelCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(16)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

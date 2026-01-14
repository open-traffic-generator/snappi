package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpVersionCounter *****
type patternFlowLacpVersionCounter struct {
	validation
	obj          *otg.PatternFlowLacpVersionCounter
	marshaller   marshalPatternFlowLacpVersionCounter
	unMarshaller unMarshalPatternFlowLacpVersionCounter
}

func NewPatternFlowLacpVersionCounter() PatternFlowLacpVersionCounter {
	obj := patternFlowLacpVersionCounter{obj: &otg.PatternFlowLacpVersionCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpVersionCounter) msg() *otg.PatternFlowLacpVersionCounter {
	return obj.obj
}

func (obj *patternFlowLacpVersionCounter) setMsg(msg *otg.PatternFlowLacpVersionCounter) PatternFlowLacpVersionCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpVersionCounter struct {
	obj *patternFlowLacpVersionCounter
}

type marshalPatternFlowLacpVersionCounter interface {
	// ToProto marshals PatternFlowLacpVersionCounter to protobuf object *otg.PatternFlowLacpVersionCounter
	ToProto() (*otg.PatternFlowLacpVersionCounter, error)
	// ToPbText marshals PatternFlowLacpVersionCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpVersionCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpVersionCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpVersionCounter struct {
	obj *patternFlowLacpVersionCounter
}

type unMarshalPatternFlowLacpVersionCounter interface {
	// FromProto unmarshals PatternFlowLacpVersionCounter from protobuf object *otg.PatternFlowLacpVersionCounter
	FromProto(msg *otg.PatternFlowLacpVersionCounter) (PatternFlowLacpVersionCounter, error)
	// FromPbText unmarshals PatternFlowLacpVersionCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpVersionCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpVersionCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpVersionCounter) Marshal() marshalPatternFlowLacpVersionCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpVersionCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpVersionCounter) Unmarshal() unMarshalPatternFlowLacpVersionCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpVersionCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpVersionCounter) ToProto() (*otg.PatternFlowLacpVersionCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpVersionCounter) FromProto(msg *otg.PatternFlowLacpVersionCounter) (PatternFlowLacpVersionCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpVersionCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpVersionCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpVersionCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpVersionCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpVersionCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpVersionCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpVersionCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpVersionCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpVersionCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpVersionCounter) Clone() (PatternFlowLacpVersionCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpVersionCounter()
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

// PatternFlowLacpVersionCounter is integer counter pattern
type PatternFlowLacpVersionCounter interface {
	Validation
	// msg marshals PatternFlowLacpVersionCounter to protobuf object *otg.PatternFlowLacpVersionCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpVersionCounter
	// setMsg unmarshals PatternFlowLacpVersionCounter from protobuf object *otg.PatternFlowLacpVersionCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpVersionCounter) PatternFlowLacpVersionCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpVersionCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpVersionCounter
	// validate validates PatternFlowLacpVersionCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpVersionCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpVersionCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpVersionCounter
	SetStart(value uint32) PatternFlowLacpVersionCounter
	// HasStart checks if Start has been set in PatternFlowLacpVersionCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpVersionCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpVersionCounter
	SetStep(value uint32) PatternFlowLacpVersionCounter
	// HasStep checks if Step has been set in PatternFlowLacpVersionCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpVersionCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpVersionCounter
	SetCount(value uint32) PatternFlowLacpVersionCounter
	// HasCount checks if Count has been set in PatternFlowLacpVersionCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpVersionCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpVersionCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpVersionCounter object
func (obj *patternFlowLacpVersionCounter) SetStart(value uint32) PatternFlowLacpVersionCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpVersionCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpVersionCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpVersionCounter object
func (obj *patternFlowLacpVersionCounter) SetStep(value uint32) PatternFlowLacpVersionCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpVersionCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpVersionCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpVersionCounter object
func (obj *patternFlowLacpVersionCounter) SetCount(value uint32) PatternFlowLacpVersionCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpVersionCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpVersionCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpVersionCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpVersionCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpVersionCounter) setDefault() {
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

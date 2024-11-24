package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGreChecksumPresentCounter *****
type patternFlowGreChecksumPresentCounter struct {
	validation
	obj          *otg.PatternFlowGreChecksumPresentCounter
	marshaller   marshalPatternFlowGreChecksumPresentCounter
	unMarshaller unMarshalPatternFlowGreChecksumPresentCounter
}

func NewPatternFlowGreChecksumPresentCounter() PatternFlowGreChecksumPresentCounter {
	obj := patternFlowGreChecksumPresentCounter{obj: &otg.PatternFlowGreChecksumPresentCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGreChecksumPresentCounter) msg() *otg.PatternFlowGreChecksumPresentCounter {
	return obj.obj
}

func (obj *patternFlowGreChecksumPresentCounter) setMsg(msg *otg.PatternFlowGreChecksumPresentCounter) PatternFlowGreChecksumPresentCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGreChecksumPresentCounter struct {
	obj *patternFlowGreChecksumPresentCounter
}

type marshalPatternFlowGreChecksumPresentCounter interface {
	// ToProto marshals PatternFlowGreChecksumPresentCounter to protobuf object *otg.PatternFlowGreChecksumPresentCounter
	ToProto() (*otg.PatternFlowGreChecksumPresentCounter, error)
	// ToPbText marshals PatternFlowGreChecksumPresentCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGreChecksumPresentCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGreChecksumPresentCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGreChecksumPresentCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGreChecksumPresentCounter struct {
	obj *patternFlowGreChecksumPresentCounter
}

type unMarshalPatternFlowGreChecksumPresentCounter interface {
	// FromProto unmarshals PatternFlowGreChecksumPresentCounter from protobuf object *otg.PatternFlowGreChecksumPresentCounter
	FromProto(msg *otg.PatternFlowGreChecksumPresentCounter) (PatternFlowGreChecksumPresentCounter, error)
	// FromPbText unmarshals PatternFlowGreChecksumPresentCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGreChecksumPresentCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGreChecksumPresentCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGreChecksumPresentCounter) Marshal() marshalPatternFlowGreChecksumPresentCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGreChecksumPresentCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGreChecksumPresentCounter) Unmarshal() unMarshalPatternFlowGreChecksumPresentCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGreChecksumPresentCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGreChecksumPresentCounter) ToProto() (*otg.PatternFlowGreChecksumPresentCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGreChecksumPresentCounter) FromProto(msg *otg.PatternFlowGreChecksumPresentCounter) (PatternFlowGreChecksumPresentCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGreChecksumPresentCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGreChecksumPresentCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGreChecksumPresentCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGreChecksumPresentCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGreChecksumPresentCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGreChecksumPresentCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGreChecksumPresentCounter) FromJson(value string) error {
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

func (obj *patternFlowGreChecksumPresentCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGreChecksumPresentCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGreChecksumPresentCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGreChecksumPresentCounter) Clone() (PatternFlowGreChecksumPresentCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGreChecksumPresentCounter()
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

// PatternFlowGreChecksumPresentCounter is integer counter pattern
type PatternFlowGreChecksumPresentCounter interface {
	Validation
	// msg marshals PatternFlowGreChecksumPresentCounter to protobuf object *otg.PatternFlowGreChecksumPresentCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGreChecksumPresentCounter
	// setMsg unmarshals PatternFlowGreChecksumPresentCounter from protobuf object *otg.PatternFlowGreChecksumPresentCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGreChecksumPresentCounter) PatternFlowGreChecksumPresentCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGreChecksumPresentCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGreChecksumPresentCounter
	// validate validates PatternFlowGreChecksumPresentCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGreChecksumPresentCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGreChecksumPresentCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGreChecksumPresentCounter
	SetStart(value uint32) PatternFlowGreChecksumPresentCounter
	// HasStart checks if Start has been set in PatternFlowGreChecksumPresentCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGreChecksumPresentCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGreChecksumPresentCounter
	SetStep(value uint32) PatternFlowGreChecksumPresentCounter
	// HasStep checks if Step has been set in PatternFlowGreChecksumPresentCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGreChecksumPresentCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGreChecksumPresentCounter
	SetCount(value uint32) PatternFlowGreChecksumPresentCounter
	// HasCount checks if Count has been set in PatternFlowGreChecksumPresentCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGreChecksumPresentCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGreChecksumPresentCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGreChecksumPresentCounter object
func (obj *patternFlowGreChecksumPresentCounter) SetStart(value uint32) PatternFlowGreChecksumPresentCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGreChecksumPresentCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGreChecksumPresentCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGreChecksumPresentCounter object
func (obj *patternFlowGreChecksumPresentCounter) SetStep(value uint32) PatternFlowGreChecksumPresentCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGreChecksumPresentCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGreChecksumPresentCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGreChecksumPresentCounter object
func (obj *patternFlowGreChecksumPresentCounter) SetCount(value uint32) PatternFlowGreChecksumPresentCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGreChecksumPresentCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreChecksumPresentCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreChecksumPresentCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreChecksumPresentCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGreChecksumPresentCounter) setDefault() {
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

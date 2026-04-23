package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmTlvsLengthCounter *****
type patternFlowCfmTlvsLengthCounter struct {
	validation
	obj          *otg.PatternFlowCfmTlvsLengthCounter
	marshaller   marshalPatternFlowCfmTlvsLengthCounter
	unMarshaller unMarshalPatternFlowCfmTlvsLengthCounter
}

func NewPatternFlowCfmTlvsLengthCounter() PatternFlowCfmTlvsLengthCounter {
	obj := patternFlowCfmTlvsLengthCounter{obj: &otg.PatternFlowCfmTlvsLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmTlvsLengthCounter) msg() *otg.PatternFlowCfmTlvsLengthCounter {
	return obj.obj
}

func (obj *patternFlowCfmTlvsLengthCounter) setMsg(msg *otg.PatternFlowCfmTlvsLengthCounter) PatternFlowCfmTlvsLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmTlvsLengthCounter struct {
	obj *patternFlowCfmTlvsLengthCounter
}

type marshalPatternFlowCfmTlvsLengthCounter interface {
	// ToProto marshals PatternFlowCfmTlvsLengthCounter to protobuf object *otg.PatternFlowCfmTlvsLengthCounter
	ToProto() (*otg.PatternFlowCfmTlvsLengthCounter, error)
	// ToPbText marshals PatternFlowCfmTlvsLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmTlvsLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmTlvsLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmTlvsLengthCounter struct {
	obj *patternFlowCfmTlvsLengthCounter
}

type unMarshalPatternFlowCfmTlvsLengthCounter interface {
	// FromProto unmarshals PatternFlowCfmTlvsLengthCounter from protobuf object *otg.PatternFlowCfmTlvsLengthCounter
	FromProto(msg *otg.PatternFlowCfmTlvsLengthCounter) (PatternFlowCfmTlvsLengthCounter, error)
	// FromPbText unmarshals PatternFlowCfmTlvsLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmTlvsLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmTlvsLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmTlvsLengthCounter) Marshal() marshalPatternFlowCfmTlvsLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmTlvsLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmTlvsLengthCounter) Unmarshal() unMarshalPatternFlowCfmTlvsLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmTlvsLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmTlvsLengthCounter) ToProto() (*otg.PatternFlowCfmTlvsLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmTlvsLengthCounter) FromProto(msg *otg.PatternFlowCfmTlvsLengthCounter) (PatternFlowCfmTlvsLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmTlvsLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmTlvsLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmTlvsLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmTlvsLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmTlvsLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmTlvsLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowCfmTlvsLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmTlvsLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmTlvsLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmTlvsLengthCounter) Clone() (PatternFlowCfmTlvsLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmTlvsLengthCounter()
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

// PatternFlowCfmTlvsLengthCounter is integer counter pattern
type PatternFlowCfmTlvsLengthCounter interface {
	Validation
	// msg marshals PatternFlowCfmTlvsLengthCounter to protobuf object *otg.PatternFlowCfmTlvsLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmTlvsLengthCounter
	// setMsg unmarshals PatternFlowCfmTlvsLengthCounter from protobuf object *otg.PatternFlowCfmTlvsLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmTlvsLengthCounter) PatternFlowCfmTlvsLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowCfmTlvsLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmTlvsLengthCounter
	// validate validates PatternFlowCfmTlvsLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmTlvsLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowCfmTlvsLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowCfmTlvsLengthCounter
	SetStart(value uint32) PatternFlowCfmTlvsLengthCounter
	// HasStart checks if Start has been set in PatternFlowCfmTlvsLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowCfmTlvsLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowCfmTlvsLengthCounter
	SetStep(value uint32) PatternFlowCfmTlvsLengthCounter
	// HasStep checks if Step has been set in PatternFlowCfmTlvsLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowCfmTlvsLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowCfmTlvsLengthCounter
	SetCount(value uint32) PatternFlowCfmTlvsLengthCounter
	// HasCount checks if Count has been set in PatternFlowCfmTlvsLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmTlvsLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmTlvsLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowCfmTlvsLengthCounter object
func (obj *patternFlowCfmTlvsLengthCounter) SetStart(value uint32) PatternFlowCfmTlvsLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmTlvsLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmTlvsLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowCfmTlvsLengthCounter object
func (obj *patternFlowCfmTlvsLengthCounter) SetStep(value uint32) PatternFlowCfmTlvsLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmTlvsLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmTlvsLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowCfmTlvsLengthCounter object
func (obj *patternFlowCfmTlvsLengthCounter) SetCount(value uint32) PatternFlowCfmTlvsLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowCfmTlvsLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmTlvsLengthCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmTlvsLengthCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmTlvsLengthCounter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowCfmTlvsLengthCounter) setDefault() {
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

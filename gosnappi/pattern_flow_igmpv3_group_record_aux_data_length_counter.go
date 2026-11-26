package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3GroupRecordAuxDataLengthCounter *****
type patternFlowIgmpv3GroupRecordAuxDataLengthCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	marshaller   marshalPatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	unMarshaller unMarshalPatternFlowIgmpv3GroupRecordAuxDataLengthCounter
}

func NewPatternFlowIgmpv3GroupRecordAuxDataLengthCounter() PatternFlowIgmpv3GroupRecordAuxDataLengthCounter {
	obj := patternFlowIgmpv3GroupRecordAuxDataLengthCounter{obj: &otg.PatternFlowIgmpv3GroupRecordAuxDataLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3GroupRecordAuxDataLengthCounter) msg() *otg.PatternFlowIgmpv3GroupRecordAuxDataLengthCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv3GroupRecordAuxDataLengthCounter) setMsg(msg *otg.PatternFlowIgmpv3GroupRecordAuxDataLengthCounter) PatternFlowIgmpv3GroupRecordAuxDataLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3GroupRecordAuxDataLengthCounter struct {
	obj *patternFlowIgmpv3GroupRecordAuxDataLengthCounter
}

type marshalPatternFlowIgmpv3GroupRecordAuxDataLengthCounter interface {
	// ToProto marshals PatternFlowIgmpv3GroupRecordAuxDataLengthCounter to protobuf object *otg.PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	ToProto() (*otg.PatternFlowIgmpv3GroupRecordAuxDataLengthCounter, error)
	// ToPbText marshals PatternFlowIgmpv3GroupRecordAuxDataLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3GroupRecordAuxDataLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3GroupRecordAuxDataLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3GroupRecordAuxDataLengthCounter struct {
	obj *patternFlowIgmpv3GroupRecordAuxDataLengthCounter
}

type unMarshalPatternFlowIgmpv3GroupRecordAuxDataLengthCounter interface {
	// FromProto unmarshals PatternFlowIgmpv3GroupRecordAuxDataLengthCounter from protobuf object *otg.PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	FromProto(msg *otg.PatternFlowIgmpv3GroupRecordAuxDataLengthCounter) (PatternFlowIgmpv3GroupRecordAuxDataLengthCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv3GroupRecordAuxDataLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3GroupRecordAuxDataLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3GroupRecordAuxDataLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3GroupRecordAuxDataLengthCounter) Marshal() marshalPatternFlowIgmpv3GroupRecordAuxDataLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3GroupRecordAuxDataLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3GroupRecordAuxDataLengthCounter) Unmarshal() unMarshalPatternFlowIgmpv3GroupRecordAuxDataLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3GroupRecordAuxDataLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3GroupRecordAuxDataLengthCounter) ToProto() (*otg.PatternFlowIgmpv3GroupRecordAuxDataLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3GroupRecordAuxDataLengthCounter) FromProto(msg *otg.PatternFlowIgmpv3GroupRecordAuxDataLengthCounter) (PatternFlowIgmpv3GroupRecordAuxDataLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3GroupRecordAuxDataLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordAuxDataLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3GroupRecordAuxDataLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordAuxDataLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3GroupRecordAuxDataLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordAuxDataLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3GroupRecordAuxDataLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3GroupRecordAuxDataLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3GroupRecordAuxDataLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3GroupRecordAuxDataLengthCounter) Clone() (PatternFlowIgmpv3GroupRecordAuxDataLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3GroupRecordAuxDataLengthCounter()
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

// PatternFlowIgmpv3GroupRecordAuxDataLengthCounter is integer counter pattern
type PatternFlowIgmpv3GroupRecordAuxDataLengthCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv3GroupRecordAuxDataLengthCounter to protobuf object *otg.PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	// setMsg unmarshals PatternFlowIgmpv3GroupRecordAuxDataLengthCounter from protobuf object *otg.PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3GroupRecordAuxDataLengthCounter) PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	// validate validates PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3GroupRecordAuxDataLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIgmpv3GroupRecordAuxDataLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	SetStart(value uint32) PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIgmpv3GroupRecordAuxDataLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	SetStep(value uint32) PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv3GroupRecordAuxDataLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	SetCount(value uint32) PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3GroupRecordAuxDataLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3GroupRecordAuxDataLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIgmpv3GroupRecordAuxDataLengthCounter object
func (obj *patternFlowIgmpv3GroupRecordAuxDataLengthCounter) SetStart(value uint32) PatternFlowIgmpv3GroupRecordAuxDataLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3GroupRecordAuxDataLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3GroupRecordAuxDataLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIgmpv3GroupRecordAuxDataLengthCounter object
func (obj *patternFlowIgmpv3GroupRecordAuxDataLengthCounter) SetStep(value uint32) PatternFlowIgmpv3GroupRecordAuxDataLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3GroupRecordAuxDataLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3GroupRecordAuxDataLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv3GroupRecordAuxDataLengthCounter object
func (obj *patternFlowIgmpv3GroupRecordAuxDataLengthCounter) SetCount(value uint32) PatternFlowIgmpv3GroupRecordAuxDataLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv3GroupRecordAuxDataLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3GroupRecordAuxDataLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3GroupRecordAuxDataLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3GroupRecordAuxDataLengthCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIgmpv3GroupRecordAuxDataLengthCounter) setDefault() {
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

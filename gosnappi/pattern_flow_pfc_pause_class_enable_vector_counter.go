package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPauseClassEnableVectorCounter *****
type patternFlowPfcPauseClassEnableVectorCounter struct {
	validation
	obj          *otg.PatternFlowPfcPauseClassEnableVectorCounter
	marshaller   marshalPatternFlowPfcPauseClassEnableVectorCounter
	unMarshaller unMarshalPatternFlowPfcPauseClassEnableVectorCounter
}

func NewPatternFlowPfcPauseClassEnableVectorCounter() PatternFlowPfcPauseClassEnableVectorCounter {
	obj := patternFlowPfcPauseClassEnableVectorCounter{obj: &otg.PatternFlowPfcPauseClassEnableVectorCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPauseClassEnableVectorCounter) msg() *otg.PatternFlowPfcPauseClassEnableVectorCounter {
	return obj.obj
}

func (obj *patternFlowPfcPauseClassEnableVectorCounter) setMsg(msg *otg.PatternFlowPfcPauseClassEnableVectorCounter) PatternFlowPfcPauseClassEnableVectorCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPauseClassEnableVectorCounter struct {
	obj *patternFlowPfcPauseClassEnableVectorCounter
}

type marshalPatternFlowPfcPauseClassEnableVectorCounter interface {
	// ToProto marshals PatternFlowPfcPauseClassEnableVectorCounter to protobuf object *otg.PatternFlowPfcPauseClassEnableVectorCounter
	ToProto() (*otg.PatternFlowPfcPauseClassEnableVectorCounter, error)
	// ToPbText marshals PatternFlowPfcPauseClassEnableVectorCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPauseClassEnableVectorCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPauseClassEnableVectorCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPauseClassEnableVectorCounter struct {
	obj *patternFlowPfcPauseClassEnableVectorCounter
}

type unMarshalPatternFlowPfcPauseClassEnableVectorCounter interface {
	// FromProto unmarshals PatternFlowPfcPauseClassEnableVectorCounter from protobuf object *otg.PatternFlowPfcPauseClassEnableVectorCounter
	FromProto(msg *otg.PatternFlowPfcPauseClassEnableVectorCounter) (PatternFlowPfcPauseClassEnableVectorCounter, error)
	// FromPbText unmarshals PatternFlowPfcPauseClassEnableVectorCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPauseClassEnableVectorCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPauseClassEnableVectorCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPauseClassEnableVectorCounter) Marshal() marshalPatternFlowPfcPauseClassEnableVectorCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPauseClassEnableVectorCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPauseClassEnableVectorCounter) Unmarshal() unMarshalPatternFlowPfcPauseClassEnableVectorCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPauseClassEnableVectorCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPauseClassEnableVectorCounter) ToProto() (*otg.PatternFlowPfcPauseClassEnableVectorCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPauseClassEnableVectorCounter) FromProto(msg *otg.PatternFlowPfcPauseClassEnableVectorCounter) (PatternFlowPfcPauseClassEnableVectorCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPauseClassEnableVectorCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseClassEnableVectorCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPauseClassEnableVectorCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseClassEnableVectorCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPauseClassEnableVectorCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseClassEnableVectorCounter) FromJson(value string) error {
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

func (obj *patternFlowPfcPauseClassEnableVectorCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseClassEnableVectorCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseClassEnableVectorCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPauseClassEnableVectorCounter) Clone() (PatternFlowPfcPauseClassEnableVectorCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPauseClassEnableVectorCounter()
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

// PatternFlowPfcPauseClassEnableVectorCounter is integer counter pattern
type PatternFlowPfcPauseClassEnableVectorCounter interface {
	Validation
	// msg marshals PatternFlowPfcPauseClassEnableVectorCounter to protobuf object *otg.PatternFlowPfcPauseClassEnableVectorCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPauseClassEnableVectorCounter
	// setMsg unmarshals PatternFlowPfcPauseClassEnableVectorCounter from protobuf object *otg.PatternFlowPfcPauseClassEnableVectorCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPauseClassEnableVectorCounter) PatternFlowPfcPauseClassEnableVectorCounter
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPauseClassEnableVectorCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPauseClassEnableVectorCounter
	// validate validates PatternFlowPfcPauseClassEnableVectorCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPauseClassEnableVectorCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowPfcPauseClassEnableVectorCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowPfcPauseClassEnableVectorCounter
	SetStart(value uint32) PatternFlowPfcPauseClassEnableVectorCounter
	// HasStart checks if Start has been set in PatternFlowPfcPauseClassEnableVectorCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowPfcPauseClassEnableVectorCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowPfcPauseClassEnableVectorCounter
	SetStep(value uint32) PatternFlowPfcPauseClassEnableVectorCounter
	// HasStep checks if Step has been set in PatternFlowPfcPauseClassEnableVectorCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowPfcPauseClassEnableVectorCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowPfcPauseClassEnableVectorCounter
	SetCount(value uint32) PatternFlowPfcPauseClassEnableVectorCounter
	// HasCount checks if Count has been set in PatternFlowPfcPauseClassEnableVectorCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPauseClassEnableVectorCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPauseClassEnableVectorCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowPfcPauseClassEnableVectorCounter object
func (obj *patternFlowPfcPauseClassEnableVectorCounter) SetStart(value uint32) PatternFlowPfcPauseClassEnableVectorCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPauseClassEnableVectorCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPauseClassEnableVectorCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowPfcPauseClassEnableVectorCounter object
func (obj *patternFlowPfcPauseClassEnableVectorCounter) SetStep(value uint32) PatternFlowPfcPauseClassEnableVectorCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPauseClassEnableVectorCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPauseClassEnableVectorCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowPfcPauseClassEnableVectorCounter object
func (obj *patternFlowPfcPauseClassEnableVectorCounter) SetCount(value uint32) PatternFlowPfcPauseClassEnableVectorCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowPfcPauseClassEnableVectorCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseClassEnableVectorCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseClassEnableVectorCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseClassEnableVectorCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowPfcPauseClassEnableVectorCounter) setDefault() {
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

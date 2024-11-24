package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPauseControlOpCodeCounter *****
type patternFlowPfcPauseControlOpCodeCounter struct {
	validation
	obj          *otg.PatternFlowPfcPauseControlOpCodeCounter
	marshaller   marshalPatternFlowPfcPauseControlOpCodeCounter
	unMarshaller unMarshalPatternFlowPfcPauseControlOpCodeCounter
}

func NewPatternFlowPfcPauseControlOpCodeCounter() PatternFlowPfcPauseControlOpCodeCounter {
	obj := patternFlowPfcPauseControlOpCodeCounter{obj: &otg.PatternFlowPfcPauseControlOpCodeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPauseControlOpCodeCounter) msg() *otg.PatternFlowPfcPauseControlOpCodeCounter {
	return obj.obj
}

func (obj *patternFlowPfcPauseControlOpCodeCounter) setMsg(msg *otg.PatternFlowPfcPauseControlOpCodeCounter) PatternFlowPfcPauseControlOpCodeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPauseControlOpCodeCounter struct {
	obj *patternFlowPfcPauseControlOpCodeCounter
}

type marshalPatternFlowPfcPauseControlOpCodeCounter interface {
	// ToProto marshals PatternFlowPfcPauseControlOpCodeCounter to protobuf object *otg.PatternFlowPfcPauseControlOpCodeCounter
	ToProto() (*otg.PatternFlowPfcPauseControlOpCodeCounter, error)
	// ToPbText marshals PatternFlowPfcPauseControlOpCodeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPauseControlOpCodeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPauseControlOpCodeCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowPfcPauseControlOpCodeCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowPfcPauseControlOpCodeCounter struct {
	obj *patternFlowPfcPauseControlOpCodeCounter
}

type unMarshalPatternFlowPfcPauseControlOpCodeCounter interface {
	// FromProto unmarshals PatternFlowPfcPauseControlOpCodeCounter from protobuf object *otg.PatternFlowPfcPauseControlOpCodeCounter
	FromProto(msg *otg.PatternFlowPfcPauseControlOpCodeCounter) (PatternFlowPfcPauseControlOpCodeCounter, error)
	// FromPbText unmarshals PatternFlowPfcPauseControlOpCodeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPauseControlOpCodeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPauseControlOpCodeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPauseControlOpCodeCounter) Marshal() marshalPatternFlowPfcPauseControlOpCodeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPauseControlOpCodeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPauseControlOpCodeCounter) Unmarshal() unMarshalPatternFlowPfcPauseControlOpCodeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPauseControlOpCodeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPauseControlOpCodeCounter) ToProto() (*otg.PatternFlowPfcPauseControlOpCodeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPauseControlOpCodeCounter) FromProto(msg *otg.PatternFlowPfcPauseControlOpCodeCounter) (PatternFlowPfcPauseControlOpCodeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPauseControlOpCodeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseControlOpCodeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPauseControlOpCodeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseControlOpCodeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPauseControlOpCodeCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowPfcPauseControlOpCodeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseControlOpCodeCounter) FromJson(value string) error {
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

func (obj *patternFlowPfcPauseControlOpCodeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseControlOpCodeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseControlOpCodeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPauseControlOpCodeCounter) Clone() (PatternFlowPfcPauseControlOpCodeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPauseControlOpCodeCounter()
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

// PatternFlowPfcPauseControlOpCodeCounter is integer counter pattern
type PatternFlowPfcPauseControlOpCodeCounter interface {
	Validation
	// msg marshals PatternFlowPfcPauseControlOpCodeCounter to protobuf object *otg.PatternFlowPfcPauseControlOpCodeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPauseControlOpCodeCounter
	// setMsg unmarshals PatternFlowPfcPauseControlOpCodeCounter from protobuf object *otg.PatternFlowPfcPauseControlOpCodeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPauseControlOpCodeCounter) PatternFlowPfcPauseControlOpCodeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPauseControlOpCodeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPauseControlOpCodeCounter
	// validate validates PatternFlowPfcPauseControlOpCodeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPauseControlOpCodeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowPfcPauseControlOpCodeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowPfcPauseControlOpCodeCounter
	SetStart(value uint32) PatternFlowPfcPauseControlOpCodeCounter
	// HasStart checks if Start has been set in PatternFlowPfcPauseControlOpCodeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowPfcPauseControlOpCodeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowPfcPauseControlOpCodeCounter
	SetStep(value uint32) PatternFlowPfcPauseControlOpCodeCounter
	// HasStep checks if Step has been set in PatternFlowPfcPauseControlOpCodeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowPfcPauseControlOpCodeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowPfcPauseControlOpCodeCounter
	SetCount(value uint32) PatternFlowPfcPauseControlOpCodeCounter
	// HasCount checks if Count has been set in PatternFlowPfcPauseControlOpCodeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPauseControlOpCodeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPauseControlOpCodeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowPfcPauseControlOpCodeCounter object
func (obj *patternFlowPfcPauseControlOpCodeCounter) SetStart(value uint32) PatternFlowPfcPauseControlOpCodeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPauseControlOpCodeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPauseControlOpCodeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowPfcPauseControlOpCodeCounter object
func (obj *patternFlowPfcPauseControlOpCodeCounter) SetStep(value uint32) PatternFlowPfcPauseControlOpCodeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPauseControlOpCodeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPauseControlOpCodeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowPfcPauseControlOpCodeCounter object
func (obj *patternFlowPfcPauseControlOpCodeCounter) SetCount(value uint32) PatternFlowPfcPauseControlOpCodeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowPfcPauseControlOpCodeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseControlOpCodeCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseControlOpCodeCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseControlOpCodeCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowPfcPauseControlOpCodeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(257)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

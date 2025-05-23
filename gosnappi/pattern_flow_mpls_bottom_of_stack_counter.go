package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowMplsBottomOfStackCounter *****
type patternFlowMplsBottomOfStackCounter struct {
	validation
	obj          *otg.PatternFlowMplsBottomOfStackCounter
	marshaller   marshalPatternFlowMplsBottomOfStackCounter
	unMarshaller unMarshalPatternFlowMplsBottomOfStackCounter
}

func NewPatternFlowMplsBottomOfStackCounter() PatternFlowMplsBottomOfStackCounter {
	obj := patternFlowMplsBottomOfStackCounter{obj: &otg.PatternFlowMplsBottomOfStackCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowMplsBottomOfStackCounter) msg() *otg.PatternFlowMplsBottomOfStackCounter {
	return obj.obj
}

func (obj *patternFlowMplsBottomOfStackCounter) setMsg(msg *otg.PatternFlowMplsBottomOfStackCounter) PatternFlowMplsBottomOfStackCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowMplsBottomOfStackCounter struct {
	obj *patternFlowMplsBottomOfStackCounter
}

type marshalPatternFlowMplsBottomOfStackCounter interface {
	// ToProto marshals PatternFlowMplsBottomOfStackCounter to protobuf object *otg.PatternFlowMplsBottomOfStackCounter
	ToProto() (*otg.PatternFlowMplsBottomOfStackCounter, error)
	// ToPbText marshals PatternFlowMplsBottomOfStackCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowMplsBottomOfStackCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowMplsBottomOfStackCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowMplsBottomOfStackCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowMplsBottomOfStackCounter struct {
	obj *patternFlowMplsBottomOfStackCounter
}

type unMarshalPatternFlowMplsBottomOfStackCounter interface {
	// FromProto unmarshals PatternFlowMplsBottomOfStackCounter from protobuf object *otg.PatternFlowMplsBottomOfStackCounter
	FromProto(msg *otg.PatternFlowMplsBottomOfStackCounter) (PatternFlowMplsBottomOfStackCounter, error)
	// FromPbText unmarshals PatternFlowMplsBottomOfStackCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowMplsBottomOfStackCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowMplsBottomOfStackCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowMplsBottomOfStackCounter) Marshal() marshalPatternFlowMplsBottomOfStackCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowMplsBottomOfStackCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowMplsBottomOfStackCounter) Unmarshal() unMarshalPatternFlowMplsBottomOfStackCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowMplsBottomOfStackCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowMplsBottomOfStackCounter) ToProto() (*otg.PatternFlowMplsBottomOfStackCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowMplsBottomOfStackCounter) FromProto(msg *otg.PatternFlowMplsBottomOfStackCounter) (PatternFlowMplsBottomOfStackCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowMplsBottomOfStackCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowMplsBottomOfStackCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowMplsBottomOfStackCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowMplsBottomOfStackCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowMplsBottomOfStackCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowMplsBottomOfStackCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowMplsBottomOfStackCounter) FromJson(value string) error {
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

func (obj *patternFlowMplsBottomOfStackCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowMplsBottomOfStackCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowMplsBottomOfStackCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowMplsBottomOfStackCounter) Clone() (PatternFlowMplsBottomOfStackCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowMplsBottomOfStackCounter()
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

// PatternFlowMplsBottomOfStackCounter is integer counter pattern
type PatternFlowMplsBottomOfStackCounter interface {
	Validation
	// msg marshals PatternFlowMplsBottomOfStackCounter to protobuf object *otg.PatternFlowMplsBottomOfStackCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowMplsBottomOfStackCounter
	// setMsg unmarshals PatternFlowMplsBottomOfStackCounter from protobuf object *otg.PatternFlowMplsBottomOfStackCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowMplsBottomOfStackCounter) PatternFlowMplsBottomOfStackCounter
	// provides marshal interface
	Marshal() marshalPatternFlowMplsBottomOfStackCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowMplsBottomOfStackCounter
	// validate validates PatternFlowMplsBottomOfStackCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowMplsBottomOfStackCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowMplsBottomOfStackCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowMplsBottomOfStackCounter
	SetStart(value uint32) PatternFlowMplsBottomOfStackCounter
	// HasStart checks if Start has been set in PatternFlowMplsBottomOfStackCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowMplsBottomOfStackCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowMplsBottomOfStackCounter
	SetStep(value uint32) PatternFlowMplsBottomOfStackCounter
	// HasStep checks if Step has been set in PatternFlowMplsBottomOfStackCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowMplsBottomOfStackCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowMplsBottomOfStackCounter
	SetCount(value uint32) PatternFlowMplsBottomOfStackCounter
	// HasCount checks if Count has been set in PatternFlowMplsBottomOfStackCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowMplsBottomOfStackCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowMplsBottomOfStackCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowMplsBottomOfStackCounter object
func (obj *patternFlowMplsBottomOfStackCounter) SetStart(value uint32) PatternFlowMplsBottomOfStackCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowMplsBottomOfStackCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowMplsBottomOfStackCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowMplsBottomOfStackCounter object
func (obj *patternFlowMplsBottomOfStackCounter) SetStep(value uint32) PatternFlowMplsBottomOfStackCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowMplsBottomOfStackCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowMplsBottomOfStackCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowMplsBottomOfStackCounter object
func (obj *patternFlowMplsBottomOfStackCounter) SetCount(value uint32) PatternFlowMplsBottomOfStackCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowMplsBottomOfStackCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsBottomOfStackCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsBottomOfStackCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsBottomOfStackCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowMplsBottomOfStackCounter) setDefault() {
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

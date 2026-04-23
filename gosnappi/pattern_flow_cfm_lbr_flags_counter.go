package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmLbrFlagsCounter *****
type patternFlowCfmLbrFlagsCounter struct {
	validation
	obj          *otg.PatternFlowCfmLbrFlagsCounter
	marshaller   marshalPatternFlowCfmLbrFlagsCounter
	unMarshaller unMarshalPatternFlowCfmLbrFlagsCounter
}

func NewPatternFlowCfmLbrFlagsCounter() PatternFlowCfmLbrFlagsCounter {
	obj := patternFlowCfmLbrFlagsCounter{obj: &otg.PatternFlowCfmLbrFlagsCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmLbrFlagsCounter) msg() *otg.PatternFlowCfmLbrFlagsCounter {
	return obj.obj
}

func (obj *patternFlowCfmLbrFlagsCounter) setMsg(msg *otg.PatternFlowCfmLbrFlagsCounter) PatternFlowCfmLbrFlagsCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmLbrFlagsCounter struct {
	obj *patternFlowCfmLbrFlagsCounter
}

type marshalPatternFlowCfmLbrFlagsCounter interface {
	// ToProto marshals PatternFlowCfmLbrFlagsCounter to protobuf object *otg.PatternFlowCfmLbrFlagsCounter
	ToProto() (*otg.PatternFlowCfmLbrFlagsCounter, error)
	// ToPbText marshals PatternFlowCfmLbrFlagsCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmLbrFlagsCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmLbrFlagsCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmLbrFlagsCounter struct {
	obj *patternFlowCfmLbrFlagsCounter
}

type unMarshalPatternFlowCfmLbrFlagsCounter interface {
	// FromProto unmarshals PatternFlowCfmLbrFlagsCounter from protobuf object *otg.PatternFlowCfmLbrFlagsCounter
	FromProto(msg *otg.PatternFlowCfmLbrFlagsCounter) (PatternFlowCfmLbrFlagsCounter, error)
	// FromPbText unmarshals PatternFlowCfmLbrFlagsCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmLbrFlagsCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmLbrFlagsCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmLbrFlagsCounter) Marshal() marshalPatternFlowCfmLbrFlagsCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmLbrFlagsCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmLbrFlagsCounter) Unmarshal() unMarshalPatternFlowCfmLbrFlagsCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmLbrFlagsCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmLbrFlagsCounter) ToProto() (*otg.PatternFlowCfmLbrFlagsCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmLbrFlagsCounter) FromProto(msg *otg.PatternFlowCfmLbrFlagsCounter) (PatternFlowCfmLbrFlagsCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmLbrFlagsCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmLbrFlagsCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmLbrFlagsCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmLbrFlagsCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmLbrFlagsCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmLbrFlagsCounter) FromJson(value string) error {
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

func (obj *patternFlowCfmLbrFlagsCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmLbrFlagsCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmLbrFlagsCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmLbrFlagsCounter) Clone() (PatternFlowCfmLbrFlagsCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmLbrFlagsCounter()
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

// PatternFlowCfmLbrFlagsCounter is integer counter pattern
type PatternFlowCfmLbrFlagsCounter interface {
	Validation
	// msg marshals PatternFlowCfmLbrFlagsCounter to protobuf object *otg.PatternFlowCfmLbrFlagsCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmLbrFlagsCounter
	// setMsg unmarshals PatternFlowCfmLbrFlagsCounter from protobuf object *otg.PatternFlowCfmLbrFlagsCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmLbrFlagsCounter) PatternFlowCfmLbrFlagsCounter
	// provides marshal interface
	Marshal() marshalPatternFlowCfmLbrFlagsCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmLbrFlagsCounter
	// validate validates PatternFlowCfmLbrFlagsCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmLbrFlagsCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowCfmLbrFlagsCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowCfmLbrFlagsCounter
	SetStart(value uint32) PatternFlowCfmLbrFlagsCounter
	// HasStart checks if Start has been set in PatternFlowCfmLbrFlagsCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowCfmLbrFlagsCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowCfmLbrFlagsCounter
	SetStep(value uint32) PatternFlowCfmLbrFlagsCounter
	// HasStep checks if Step has been set in PatternFlowCfmLbrFlagsCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowCfmLbrFlagsCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowCfmLbrFlagsCounter
	SetCount(value uint32) PatternFlowCfmLbrFlagsCounter
	// HasCount checks if Count has been set in PatternFlowCfmLbrFlagsCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmLbrFlagsCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmLbrFlagsCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowCfmLbrFlagsCounter object
func (obj *patternFlowCfmLbrFlagsCounter) SetStart(value uint32) PatternFlowCfmLbrFlagsCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmLbrFlagsCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmLbrFlagsCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowCfmLbrFlagsCounter object
func (obj *patternFlowCfmLbrFlagsCounter) SetStep(value uint32) PatternFlowCfmLbrFlagsCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmLbrFlagsCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmLbrFlagsCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowCfmLbrFlagsCounter object
func (obj *patternFlowCfmLbrFlagsCounter) SetCount(value uint32) PatternFlowCfmLbrFlagsCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowCfmLbrFlagsCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmLbrFlagsCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmLbrFlagsCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmLbrFlagsCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowCfmLbrFlagsCounter) setDefault() {
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

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmVersionCounter *****
type patternFlowCfmVersionCounter struct {
	validation
	obj          *otg.PatternFlowCfmVersionCounter
	marshaller   marshalPatternFlowCfmVersionCounter
	unMarshaller unMarshalPatternFlowCfmVersionCounter
}

func NewPatternFlowCfmVersionCounter() PatternFlowCfmVersionCounter {
	obj := patternFlowCfmVersionCounter{obj: &otg.PatternFlowCfmVersionCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmVersionCounter) msg() *otg.PatternFlowCfmVersionCounter {
	return obj.obj
}

func (obj *patternFlowCfmVersionCounter) setMsg(msg *otg.PatternFlowCfmVersionCounter) PatternFlowCfmVersionCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmVersionCounter struct {
	obj *patternFlowCfmVersionCounter
}

type marshalPatternFlowCfmVersionCounter interface {
	// ToProto marshals PatternFlowCfmVersionCounter to protobuf object *otg.PatternFlowCfmVersionCounter
	ToProto() (*otg.PatternFlowCfmVersionCounter, error)
	// ToPbText marshals PatternFlowCfmVersionCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmVersionCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmVersionCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmVersionCounter struct {
	obj *patternFlowCfmVersionCounter
}

type unMarshalPatternFlowCfmVersionCounter interface {
	// FromProto unmarshals PatternFlowCfmVersionCounter from protobuf object *otg.PatternFlowCfmVersionCounter
	FromProto(msg *otg.PatternFlowCfmVersionCounter) (PatternFlowCfmVersionCounter, error)
	// FromPbText unmarshals PatternFlowCfmVersionCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmVersionCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmVersionCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmVersionCounter) Marshal() marshalPatternFlowCfmVersionCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmVersionCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmVersionCounter) Unmarshal() unMarshalPatternFlowCfmVersionCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmVersionCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmVersionCounter) ToProto() (*otg.PatternFlowCfmVersionCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmVersionCounter) FromProto(msg *otg.PatternFlowCfmVersionCounter) (PatternFlowCfmVersionCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmVersionCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmVersionCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmVersionCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmVersionCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmVersionCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmVersionCounter) FromJson(value string) error {
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

func (obj *patternFlowCfmVersionCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmVersionCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmVersionCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmVersionCounter) Clone() (PatternFlowCfmVersionCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmVersionCounter()
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

// PatternFlowCfmVersionCounter is integer counter pattern
type PatternFlowCfmVersionCounter interface {
	Validation
	// msg marshals PatternFlowCfmVersionCounter to protobuf object *otg.PatternFlowCfmVersionCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmVersionCounter
	// setMsg unmarshals PatternFlowCfmVersionCounter from protobuf object *otg.PatternFlowCfmVersionCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmVersionCounter) PatternFlowCfmVersionCounter
	// provides marshal interface
	Marshal() marshalPatternFlowCfmVersionCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmVersionCounter
	// validate validates PatternFlowCfmVersionCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmVersionCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowCfmVersionCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowCfmVersionCounter
	SetStart(value uint32) PatternFlowCfmVersionCounter
	// HasStart checks if Start has been set in PatternFlowCfmVersionCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowCfmVersionCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowCfmVersionCounter
	SetStep(value uint32) PatternFlowCfmVersionCounter
	// HasStep checks if Step has been set in PatternFlowCfmVersionCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowCfmVersionCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowCfmVersionCounter
	SetCount(value uint32) PatternFlowCfmVersionCounter
	// HasCount checks if Count has been set in PatternFlowCfmVersionCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmVersionCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmVersionCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowCfmVersionCounter object
func (obj *patternFlowCfmVersionCounter) SetStart(value uint32) PatternFlowCfmVersionCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmVersionCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmVersionCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowCfmVersionCounter object
func (obj *patternFlowCfmVersionCounter) SetStep(value uint32) PatternFlowCfmVersionCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmVersionCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmVersionCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowCfmVersionCounter object
func (obj *patternFlowCfmVersionCounter) SetCount(value uint32) PatternFlowCfmVersionCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowCfmVersionCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmVersionCounter.Start <= 31 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmVersionCounter.Step <= 31 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmVersionCounter.Count <= 32 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowCfmVersionCounter) setDefault() {
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

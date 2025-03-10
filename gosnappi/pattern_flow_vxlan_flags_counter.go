package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVxlanFlagsCounter *****
type patternFlowVxlanFlagsCounter struct {
	validation
	obj          *otg.PatternFlowVxlanFlagsCounter
	marshaller   marshalPatternFlowVxlanFlagsCounter
	unMarshaller unMarshalPatternFlowVxlanFlagsCounter
}

func NewPatternFlowVxlanFlagsCounter() PatternFlowVxlanFlagsCounter {
	obj := patternFlowVxlanFlagsCounter{obj: &otg.PatternFlowVxlanFlagsCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVxlanFlagsCounter) msg() *otg.PatternFlowVxlanFlagsCounter {
	return obj.obj
}

func (obj *patternFlowVxlanFlagsCounter) setMsg(msg *otg.PatternFlowVxlanFlagsCounter) PatternFlowVxlanFlagsCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVxlanFlagsCounter struct {
	obj *patternFlowVxlanFlagsCounter
}

type marshalPatternFlowVxlanFlagsCounter interface {
	// ToProto marshals PatternFlowVxlanFlagsCounter to protobuf object *otg.PatternFlowVxlanFlagsCounter
	ToProto() (*otg.PatternFlowVxlanFlagsCounter, error)
	// ToPbText marshals PatternFlowVxlanFlagsCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVxlanFlagsCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVxlanFlagsCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowVxlanFlagsCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowVxlanFlagsCounter struct {
	obj *patternFlowVxlanFlagsCounter
}

type unMarshalPatternFlowVxlanFlagsCounter interface {
	// FromProto unmarshals PatternFlowVxlanFlagsCounter from protobuf object *otg.PatternFlowVxlanFlagsCounter
	FromProto(msg *otg.PatternFlowVxlanFlagsCounter) (PatternFlowVxlanFlagsCounter, error)
	// FromPbText unmarshals PatternFlowVxlanFlagsCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVxlanFlagsCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVxlanFlagsCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVxlanFlagsCounter) Marshal() marshalPatternFlowVxlanFlagsCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVxlanFlagsCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVxlanFlagsCounter) Unmarshal() unMarshalPatternFlowVxlanFlagsCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVxlanFlagsCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVxlanFlagsCounter) ToProto() (*otg.PatternFlowVxlanFlagsCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVxlanFlagsCounter) FromProto(msg *otg.PatternFlowVxlanFlagsCounter) (PatternFlowVxlanFlagsCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVxlanFlagsCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVxlanFlagsCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowVxlanFlagsCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVxlanFlagsCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowVxlanFlagsCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowVxlanFlagsCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVxlanFlagsCounter) FromJson(value string) error {
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

func (obj *patternFlowVxlanFlagsCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVxlanFlagsCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVxlanFlagsCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVxlanFlagsCounter) Clone() (PatternFlowVxlanFlagsCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVxlanFlagsCounter()
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

// PatternFlowVxlanFlagsCounter is integer counter pattern
type PatternFlowVxlanFlagsCounter interface {
	Validation
	// msg marshals PatternFlowVxlanFlagsCounter to protobuf object *otg.PatternFlowVxlanFlagsCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowVxlanFlagsCounter
	// setMsg unmarshals PatternFlowVxlanFlagsCounter from protobuf object *otg.PatternFlowVxlanFlagsCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVxlanFlagsCounter) PatternFlowVxlanFlagsCounter
	// provides marshal interface
	Marshal() marshalPatternFlowVxlanFlagsCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVxlanFlagsCounter
	// validate validates PatternFlowVxlanFlagsCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVxlanFlagsCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowVxlanFlagsCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowVxlanFlagsCounter
	SetStart(value uint32) PatternFlowVxlanFlagsCounter
	// HasStart checks if Start has been set in PatternFlowVxlanFlagsCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowVxlanFlagsCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowVxlanFlagsCounter
	SetStep(value uint32) PatternFlowVxlanFlagsCounter
	// HasStep checks if Step has been set in PatternFlowVxlanFlagsCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowVxlanFlagsCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowVxlanFlagsCounter
	SetCount(value uint32) PatternFlowVxlanFlagsCounter
	// HasCount checks if Count has been set in PatternFlowVxlanFlagsCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowVxlanFlagsCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowVxlanFlagsCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowVxlanFlagsCounter object
func (obj *patternFlowVxlanFlagsCounter) SetStart(value uint32) PatternFlowVxlanFlagsCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowVxlanFlagsCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowVxlanFlagsCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowVxlanFlagsCounter object
func (obj *patternFlowVxlanFlagsCounter) SetStep(value uint32) PatternFlowVxlanFlagsCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowVxlanFlagsCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowVxlanFlagsCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowVxlanFlagsCounter object
func (obj *patternFlowVxlanFlagsCounter) SetCount(value uint32) PatternFlowVxlanFlagsCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowVxlanFlagsCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVxlanFlagsCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVxlanFlagsCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVxlanFlagsCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowVxlanFlagsCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(8)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

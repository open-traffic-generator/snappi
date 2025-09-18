package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduCollectorTypeCounter *****
type patternFlowLacpduCollectorTypeCounter struct {
	validation
	obj          *otg.PatternFlowLacpduCollectorTypeCounter
	marshaller   marshalPatternFlowLacpduCollectorTypeCounter
	unMarshaller unMarshalPatternFlowLacpduCollectorTypeCounter
}

func NewPatternFlowLacpduCollectorTypeCounter() PatternFlowLacpduCollectorTypeCounter {
	obj := patternFlowLacpduCollectorTypeCounter{obj: &otg.PatternFlowLacpduCollectorTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduCollectorTypeCounter) msg() *otg.PatternFlowLacpduCollectorTypeCounter {
	return obj.obj
}

func (obj *patternFlowLacpduCollectorTypeCounter) setMsg(msg *otg.PatternFlowLacpduCollectorTypeCounter) PatternFlowLacpduCollectorTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduCollectorTypeCounter struct {
	obj *patternFlowLacpduCollectorTypeCounter
}

type marshalPatternFlowLacpduCollectorTypeCounter interface {
	// ToProto marshals PatternFlowLacpduCollectorTypeCounter to protobuf object *otg.PatternFlowLacpduCollectorTypeCounter
	ToProto() (*otg.PatternFlowLacpduCollectorTypeCounter, error)
	// ToPbText marshals PatternFlowLacpduCollectorTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduCollectorTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduCollectorTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduCollectorTypeCounter struct {
	obj *patternFlowLacpduCollectorTypeCounter
}

type unMarshalPatternFlowLacpduCollectorTypeCounter interface {
	// FromProto unmarshals PatternFlowLacpduCollectorTypeCounter from protobuf object *otg.PatternFlowLacpduCollectorTypeCounter
	FromProto(msg *otg.PatternFlowLacpduCollectorTypeCounter) (PatternFlowLacpduCollectorTypeCounter, error)
	// FromPbText unmarshals PatternFlowLacpduCollectorTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduCollectorTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduCollectorTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduCollectorTypeCounter) Marshal() marshalPatternFlowLacpduCollectorTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduCollectorTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduCollectorTypeCounter) Unmarshal() unMarshalPatternFlowLacpduCollectorTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduCollectorTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduCollectorTypeCounter) ToProto() (*otg.PatternFlowLacpduCollectorTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduCollectorTypeCounter) FromProto(msg *otg.PatternFlowLacpduCollectorTypeCounter) (PatternFlowLacpduCollectorTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduCollectorTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduCollectorTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduCollectorTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduCollectorTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduCollectorTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduCollectorTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduCollectorTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduCollectorTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduCollectorTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduCollectorTypeCounter) Clone() (PatternFlowLacpduCollectorTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduCollectorTypeCounter()
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

// PatternFlowLacpduCollectorTypeCounter is integer counter pattern
type PatternFlowLacpduCollectorTypeCounter interface {
	Validation
	// msg marshals PatternFlowLacpduCollectorTypeCounter to protobuf object *otg.PatternFlowLacpduCollectorTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduCollectorTypeCounter
	// setMsg unmarshals PatternFlowLacpduCollectorTypeCounter from protobuf object *otg.PatternFlowLacpduCollectorTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduCollectorTypeCounter) PatternFlowLacpduCollectorTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduCollectorTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduCollectorTypeCounter
	// validate validates PatternFlowLacpduCollectorTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduCollectorTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduCollectorTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduCollectorTypeCounter
	SetStart(value uint32) PatternFlowLacpduCollectorTypeCounter
	// HasStart checks if Start has been set in PatternFlowLacpduCollectorTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduCollectorTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduCollectorTypeCounter
	SetStep(value uint32) PatternFlowLacpduCollectorTypeCounter
	// HasStep checks if Step has been set in PatternFlowLacpduCollectorTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduCollectorTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduCollectorTypeCounter
	SetCount(value uint32) PatternFlowLacpduCollectorTypeCounter
	// HasCount checks if Count has been set in PatternFlowLacpduCollectorTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduCollectorTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduCollectorTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduCollectorTypeCounter object
func (obj *patternFlowLacpduCollectorTypeCounter) SetStart(value uint32) PatternFlowLacpduCollectorTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduCollectorTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduCollectorTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduCollectorTypeCounter object
func (obj *patternFlowLacpduCollectorTypeCounter) SetStep(value uint32) PatternFlowLacpduCollectorTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduCollectorTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduCollectorTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduCollectorTypeCounter object
func (obj *patternFlowLacpduCollectorTypeCounter) SetCount(value uint32) PatternFlowLacpduCollectorTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduCollectorTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduCollectorTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduCollectorTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduCollectorTypeCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduCollectorTypeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(3)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

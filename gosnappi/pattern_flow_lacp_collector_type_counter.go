package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpCollectorTypeCounter *****
type patternFlowLacpCollectorTypeCounter struct {
	validation
	obj          *otg.PatternFlowLacpCollectorTypeCounter
	marshaller   marshalPatternFlowLacpCollectorTypeCounter
	unMarshaller unMarshalPatternFlowLacpCollectorTypeCounter
}

func NewPatternFlowLacpCollectorTypeCounter() PatternFlowLacpCollectorTypeCounter {
	obj := patternFlowLacpCollectorTypeCounter{obj: &otg.PatternFlowLacpCollectorTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpCollectorTypeCounter) msg() *otg.PatternFlowLacpCollectorTypeCounter {
	return obj.obj
}

func (obj *patternFlowLacpCollectorTypeCounter) setMsg(msg *otg.PatternFlowLacpCollectorTypeCounter) PatternFlowLacpCollectorTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpCollectorTypeCounter struct {
	obj *patternFlowLacpCollectorTypeCounter
}

type marshalPatternFlowLacpCollectorTypeCounter interface {
	// ToProto marshals PatternFlowLacpCollectorTypeCounter to protobuf object *otg.PatternFlowLacpCollectorTypeCounter
	ToProto() (*otg.PatternFlowLacpCollectorTypeCounter, error)
	// ToPbText marshals PatternFlowLacpCollectorTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpCollectorTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpCollectorTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpCollectorTypeCounter struct {
	obj *patternFlowLacpCollectorTypeCounter
}

type unMarshalPatternFlowLacpCollectorTypeCounter interface {
	// FromProto unmarshals PatternFlowLacpCollectorTypeCounter from protobuf object *otg.PatternFlowLacpCollectorTypeCounter
	FromProto(msg *otg.PatternFlowLacpCollectorTypeCounter) (PatternFlowLacpCollectorTypeCounter, error)
	// FromPbText unmarshals PatternFlowLacpCollectorTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpCollectorTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpCollectorTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpCollectorTypeCounter) Marshal() marshalPatternFlowLacpCollectorTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpCollectorTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpCollectorTypeCounter) Unmarshal() unMarshalPatternFlowLacpCollectorTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpCollectorTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpCollectorTypeCounter) ToProto() (*otg.PatternFlowLacpCollectorTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpCollectorTypeCounter) FromProto(msg *otg.PatternFlowLacpCollectorTypeCounter) (PatternFlowLacpCollectorTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpCollectorTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpCollectorTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpCollectorTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpCollectorTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpCollectorTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpCollectorTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpCollectorTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpCollectorTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpCollectorTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpCollectorTypeCounter) Clone() (PatternFlowLacpCollectorTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpCollectorTypeCounter()
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

// PatternFlowLacpCollectorTypeCounter is integer counter pattern
type PatternFlowLacpCollectorTypeCounter interface {
	Validation
	// msg marshals PatternFlowLacpCollectorTypeCounter to protobuf object *otg.PatternFlowLacpCollectorTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpCollectorTypeCounter
	// setMsg unmarshals PatternFlowLacpCollectorTypeCounter from protobuf object *otg.PatternFlowLacpCollectorTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpCollectorTypeCounter) PatternFlowLacpCollectorTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpCollectorTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpCollectorTypeCounter
	// validate validates PatternFlowLacpCollectorTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpCollectorTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpCollectorTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpCollectorTypeCounter
	SetStart(value uint32) PatternFlowLacpCollectorTypeCounter
	// HasStart checks if Start has been set in PatternFlowLacpCollectorTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpCollectorTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpCollectorTypeCounter
	SetStep(value uint32) PatternFlowLacpCollectorTypeCounter
	// HasStep checks if Step has been set in PatternFlowLacpCollectorTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpCollectorTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpCollectorTypeCounter
	SetCount(value uint32) PatternFlowLacpCollectorTypeCounter
	// HasCount checks if Count has been set in PatternFlowLacpCollectorTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpCollectorTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpCollectorTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpCollectorTypeCounter object
func (obj *patternFlowLacpCollectorTypeCounter) SetStart(value uint32) PatternFlowLacpCollectorTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpCollectorTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpCollectorTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpCollectorTypeCounter object
func (obj *patternFlowLacpCollectorTypeCounter) SetStep(value uint32) PatternFlowLacpCollectorTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpCollectorTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpCollectorTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpCollectorTypeCounter object
func (obj *patternFlowLacpCollectorTypeCounter) SetCount(value uint32) PatternFlowLacpCollectorTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpCollectorTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpCollectorTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpCollectorTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpCollectorTypeCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpCollectorTypeCounter) setDefault() {
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

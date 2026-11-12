package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv2MaxResponseTimeCounter *****
type patternFlowIgmpv2MaxResponseTimeCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv2MaxResponseTimeCounter
	marshaller   marshalPatternFlowIgmpv2MaxResponseTimeCounter
	unMarshaller unMarshalPatternFlowIgmpv2MaxResponseTimeCounter
}

func NewPatternFlowIgmpv2MaxResponseTimeCounter() PatternFlowIgmpv2MaxResponseTimeCounter {
	obj := patternFlowIgmpv2MaxResponseTimeCounter{obj: &otg.PatternFlowIgmpv2MaxResponseTimeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv2MaxResponseTimeCounter) msg() *otg.PatternFlowIgmpv2MaxResponseTimeCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv2MaxResponseTimeCounter) setMsg(msg *otg.PatternFlowIgmpv2MaxResponseTimeCounter) PatternFlowIgmpv2MaxResponseTimeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv2MaxResponseTimeCounter struct {
	obj *patternFlowIgmpv2MaxResponseTimeCounter
}

type marshalPatternFlowIgmpv2MaxResponseTimeCounter interface {
	// ToProto marshals PatternFlowIgmpv2MaxResponseTimeCounter to protobuf object *otg.PatternFlowIgmpv2MaxResponseTimeCounter
	ToProto() (*otg.PatternFlowIgmpv2MaxResponseTimeCounter, error)
	// ToPbText marshals PatternFlowIgmpv2MaxResponseTimeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv2MaxResponseTimeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv2MaxResponseTimeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv2MaxResponseTimeCounter struct {
	obj *patternFlowIgmpv2MaxResponseTimeCounter
}

type unMarshalPatternFlowIgmpv2MaxResponseTimeCounter interface {
	// FromProto unmarshals PatternFlowIgmpv2MaxResponseTimeCounter from protobuf object *otg.PatternFlowIgmpv2MaxResponseTimeCounter
	FromProto(msg *otg.PatternFlowIgmpv2MaxResponseTimeCounter) (PatternFlowIgmpv2MaxResponseTimeCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv2MaxResponseTimeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv2MaxResponseTimeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv2MaxResponseTimeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv2MaxResponseTimeCounter) Marshal() marshalPatternFlowIgmpv2MaxResponseTimeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv2MaxResponseTimeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv2MaxResponseTimeCounter) Unmarshal() unMarshalPatternFlowIgmpv2MaxResponseTimeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv2MaxResponseTimeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv2MaxResponseTimeCounter) ToProto() (*otg.PatternFlowIgmpv2MaxResponseTimeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv2MaxResponseTimeCounter) FromProto(msg *otg.PatternFlowIgmpv2MaxResponseTimeCounter) (PatternFlowIgmpv2MaxResponseTimeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv2MaxResponseTimeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv2MaxResponseTimeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv2MaxResponseTimeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv2MaxResponseTimeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv2MaxResponseTimeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv2MaxResponseTimeCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv2MaxResponseTimeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv2MaxResponseTimeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv2MaxResponseTimeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv2MaxResponseTimeCounter) Clone() (PatternFlowIgmpv2MaxResponseTimeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv2MaxResponseTimeCounter()
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

// PatternFlowIgmpv2MaxResponseTimeCounter is integer counter pattern
type PatternFlowIgmpv2MaxResponseTimeCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv2MaxResponseTimeCounter to protobuf object *otg.PatternFlowIgmpv2MaxResponseTimeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv2MaxResponseTimeCounter
	// setMsg unmarshals PatternFlowIgmpv2MaxResponseTimeCounter from protobuf object *otg.PatternFlowIgmpv2MaxResponseTimeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv2MaxResponseTimeCounter) PatternFlowIgmpv2MaxResponseTimeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv2MaxResponseTimeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv2MaxResponseTimeCounter
	// validate validates PatternFlowIgmpv2MaxResponseTimeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv2MaxResponseTimeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIgmpv2MaxResponseTimeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIgmpv2MaxResponseTimeCounter
	SetStart(value uint32) PatternFlowIgmpv2MaxResponseTimeCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv2MaxResponseTimeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIgmpv2MaxResponseTimeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIgmpv2MaxResponseTimeCounter
	SetStep(value uint32) PatternFlowIgmpv2MaxResponseTimeCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv2MaxResponseTimeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv2MaxResponseTimeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv2MaxResponseTimeCounter
	SetCount(value uint32) PatternFlowIgmpv2MaxResponseTimeCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv2MaxResponseTimeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv2MaxResponseTimeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv2MaxResponseTimeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIgmpv2MaxResponseTimeCounter object
func (obj *patternFlowIgmpv2MaxResponseTimeCounter) SetStart(value uint32) PatternFlowIgmpv2MaxResponseTimeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv2MaxResponseTimeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv2MaxResponseTimeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIgmpv2MaxResponseTimeCounter object
func (obj *patternFlowIgmpv2MaxResponseTimeCounter) SetStep(value uint32) PatternFlowIgmpv2MaxResponseTimeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv2MaxResponseTimeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv2MaxResponseTimeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv2MaxResponseTimeCounter object
func (obj *patternFlowIgmpv2MaxResponseTimeCounter) SetCount(value uint32) PatternFlowIgmpv2MaxResponseTimeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv2MaxResponseTimeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv2MaxResponseTimeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv2MaxResponseTimeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv2MaxResponseTimeCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIgmpv2MaxResponseTimeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(100)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

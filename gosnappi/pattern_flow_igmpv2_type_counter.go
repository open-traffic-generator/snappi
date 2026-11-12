package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv2TypeCounter *****
type patternFlowIgmpv2TypeCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv2TypeCounter
	marshaller   marshalPatternFlowIgmpv2TypeCounter
	unMarshaller unMarshalPatternFlowIgmpv2TypeCounter
}

func NewPatternFlowIgmpv2TypeCounter() PatternFlowIgmpv2TypeCounter {
	obj := patternFlowIgmpv2TypeCounter{obj: &otg.PatternFlowIgmpv2TypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv2TypeCounter) msg() *otg.PatternFlowIgmpv2TypeCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv2TypeCounter) setMsg(msg *otg.PatternFlowIgmpv2TypeCounter) PatternFlowIgmpv2TypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv2TypeCounter struct {
	obj *patternFlowIgmpv2TypeCounter
}

type marshalPatternFlowIgmpv2TypeCounter interface {
	// ToProto marshals PatternFlowIgmpv2TypeCounter to protobuf object *otg.PatternFlowIgmpv2TypeCounter
	ToProto() (*otg.PatternFlowIgmpv2TypeCounter, error)
	// ToPbText marshals PatternFlowIgmpv2TypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv2TypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv2TypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv2TypeCounter struct {
	obj *patternFlowIgmpv2TypeCounter
}

type unMarshalPatternFlowIgmpv2TypeCounter interface {
	// FromProto unmarshals PatternFlowIgmpv2TypeCounter from protobuf object *otg.PatternFlowIgmpv2TypeCounter
	FromProto(msg *otg.PatternFlowIgmpv2TypeCounter) (PatternFlowIgmpv2TypeCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv2TypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv2TypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv2TypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv2TypeCounter) Marshal() marshalPatternFlowIgmpv2TypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv2TypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv2TypeCounter) Unmarshal() unMarshalPatternFlowIgmpv2TypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv2TypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv2TypeCounter) ToProto() (*otg.PatternFlowIgmpv2TypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv2TypeCounter) FromProto(msg *otg.PatternFlowIgmpv2TypeCounter) (PatternFlowIgmpv2TypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv2TypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv2TypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv2TypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv2TypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv2TypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv2TypeCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv2TypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv2TypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv2TypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv2TypeCounter) Clone() (PatternFlowIgmpv2TypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv2TypeCounter()
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

// PatternFlowIgmpv2TypeCounter is integer counter pattern
type PatternFlowIgmpv2TypeCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv2TypeCounter to protobuf object *otg.PatternFlowIgmpv2TypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv2TypeCounter
	// setMsg unmarshals PatternFlowIgmpv2TypeCounter from protobuf object *otg.PatternFlowIgmpv2TypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv2TypeCounter) PatternFlowIgmpv2TypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv2TypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv2TypeCounter
	// validate validates PatternFlowIgmpv2TypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv2TypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIgmpv2TypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIgmpv2TypeCounter
	SetStart(value uint32) PatternFlowIgmpv2TypeCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv2TypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIgmpv2TypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIgmpv2TypeCounter
	SetStep(value uint32) PatternFlowIgmpv2TypeCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv2TypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv2TypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv2TypeCounter
	SetCount(value uint32) PatternFlowIgmpv2TypeCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv2TypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv2TypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv2TypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIgmpv2TypeCounter object
func (obj *patternFlowIgmpv2TypeCounter) SetStart(value uint32) PatternFlowIgmpv2TypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv2TypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv2TypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIgmpv2TypeCounter object
func (obj *patternFlowIgmpv2TypeCounter) SetStep(value uint32) PatternFlowIgmpv2TypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv2TypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv2TypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv2TypeCounter object
func (obj *patternFlowIgmpv2TypeCounter) SetCount(value uint32) PatternFlowIgmpv2TypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv2TypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv2TypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv2TypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv2TypeCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIgmpv2TypeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(17)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

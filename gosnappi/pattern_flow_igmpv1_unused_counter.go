package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv1UnusedCounter *****
type patternFlowIgmpv1UnusedCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv1UnusedCounter
	marshaller   marshalPatternFlowIgmpv1UnusedCounter
	unMarshaller unMarshalPatternFlowIgmpv1UnusedCounter
}

func NewPatternFlowIgmpv1UnusedCounter() PatternFlowIgmpv1UnusedCounter {
	obj := patternFlowIgmpv1UnusedCounter{obj: &otg.PatternFlowIgmpv1UnusedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv1UnusedCounter) msg() *otg.PatternFlowIgmpv1UnusedCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv1UnusedCounter) setMsg(msg *otg.PatternFlowIgmpv1UnusedCounter) PatternFlowIgmpv1UnusedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv1UnusedCounter struct {
	obj *patternFlowIgmpv1UnusedCounter
}

type marshalPatternFlowIgmpv1UnusedCounter interface {
	// ToProto marshals PatternFlowIgmpv1UnusedCounter to protobuf object *otg.PatternFlowIgmpv1UnusedCounter
	ToProto() (*otg.PatternFlowIgmpv1UnusedCounter, error)
	// ToPbText marshals PatternFlowIgmpv1UnusedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv1UnusedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv1UnusedCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIgmpv1UnusedCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIgmpv1UnusedCounter struct {
	obj *patternFlowIgmpv1UnusedCounter
}

type unMarshalPatternFlowIgmpv1UnusedCounter interface {
	// FromProto unmarshals PatternFlowIgmpv1UnusedCounter from protobuf object *otg.PatternFlowIgmpv1UnusedCounter
	FromProto(msg *otg.PatternFlowIgmpv1UnusedCounter) (PatternFlowIgmpv1UnusedCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv1UnusedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv1UnusedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv1UnusedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv1UnusedCounter) Marshal() marshalPatternFlowIgmpv1UnusedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv1UnusedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv1UnusedCounter) Unmarshal() unMarshalPatternFlowIgmpv1UnusedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv1UnusedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv1UnusedCounter) ToProto() (*otg.PatternFlowIgmpv1UnusedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv1UnusedCounter) FromProto(msg *otg.PatternFlowIgmpv1UnusedCounter) (PatternFlowIgmpv1UnusedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv1UnusedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1UnusedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv1UnusedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1UnusedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv1UnusedCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIgmpv1UnusedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1UnusedCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv1UnusedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1UnusedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1UnusedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv1UnusedCounter) Clone() (PatternFlowIgmpv1UnusedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv1UnusedCounter()
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

// PatternFlowIgmpv1UnusedCounter is integer counter pattern
type PatternFlowIgmpv1UnusedCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv1UnusedCounter to protobuf object *otg.PatternFlowIgmpv1UnusedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv1UnusedCounter
	// setMsg unmarshals PatternFlowIgmpv1UnusedCounter from protobuf object *otg.PatternFlowIgmpv1UnusedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv1UnusedCounter) PatternFlowIgmpv1UnusedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv1UnusedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv1UnusedCounter
	// validate validates PatternFlowIgmpv1UnusedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv1UnusedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIgmpv1UnusedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIgmpv1UnusedCounter
	SetStart(value uint32) PatternFlowIgmpv1UnusedCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv1UnusedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIgmpv1UnusedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIgmpv1UnusedCounter
	SetStep(value uint32) PatternFlowIgmpv1UnusedCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv1UnusedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv1UnusedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv1UnusedCounter
	SetCount(value uint32) PatternFlowIgmpv1UnusedCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv1UnusedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv1UnusedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv1UnusedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIgmpv1UnusedCounter object
func (obj *patternFlowIgmpv1UnusedCounter) SetStart(value uint32) PatternFlowIgmpv1UnusedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv1UnusedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv1UnusedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIgmpv1UnusedCounter object
func (obj *patternFlowIgmpv1UnusedCounter) SetStep(value uint32) PatternFlowIgmpv1UnusedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv1UnusedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv1UnusedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv1UnusedCounter object
func (obj *patternFlowIgmpv1UnusedCounter) SetCount(value uint32) PatternFlowIgmpv1UnusedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv1UnusedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1UnusedCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1UnusedCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1UnusedCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIgmpv1UnusedCounter) setDefault() {
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

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv1TypeCounter *****
type patternFlowIgmpv1TypeCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv1TypeCounter
	marshaller   marshalPatternFlowIgmpv1TypeCounter
	unMarshaller unMarshalPatternFlowIgmpv1TypeCounter
}

func NewPatternFlowIgmpv1TypeCounter() PatternFlowIgmpv1TypeCounter {
	obj := patternFlowIgmpv1TypeCounter{obj: &otg.PatternFlowIgmpv1TypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv1TypeCounter) msg() *otg.PatternFlowIgmpv1TypeCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv1TypeCounter) setMsg(msg *otg.PatternFlowIgmpv1TypeCounter) PatternFlowIgmpv1TypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv1TypeCounter struct {
	obj *patternFlowIgmpv1TypeCounter
}

type marshalPatternFlowIgmpv1TypeCounter interface {
	// ToProto marshals PatternFlowIgmpv1TypeCounter to protobuf object *otg.PatternFlowIgmpv1TypeCounter
	ToProto() (*otg.PatternFlowIgmpv1TypeCounter, error)
	// ToPbText marshals PatternFlowIgmpv1TypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv1TypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv1TypeCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIgmpv1TypeCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIgmpv1TypeCounter struct {
	obj *patternFlowIgmpv1TypeCounter
}

type unMarshalPatternFlowIgmpv1TypeCounter interface {
	// FromProto unmarshals PatternFlowIgmpv1TypeCounter from protobuf object *otg.PatternFlowIgmpv1TypeCounter
	FromProto(msg *otg.PatternFlowIgmpv1TypeCounter) (PatternFlowIgmpv1TypeCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv1TypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv1TypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv1TypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv1TypeCounter) Marshal() marshalPatternFlowIgmpv1TypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv1TypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv1TypeCounter) Unmarshal() unMarshalPatternFlowIgmpv1TypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv1TypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv1TypeCounter) ToProto() (*otg.PatternFlowIgmpv1TypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv1TypeCounter) FromProto(msg *otg.PatternFlowIgmpv1TypeCounter) (PatternFlowIgmpv1TypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv1TypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1TypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv1TypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1TypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv1TypeCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIgmpv1TypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1TypeCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv1TypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1TypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1TypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv1TypeCounter) Clone() (PatternFlowIgmpv1TypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv1TypeCounter()
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

// PatternFlowIgmpv1TypeCounter is integer counter pattern
type PatternFlowIgmpv1TypeCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv1TypeCounter to protobuf object *otg.PatternFlowIgmpv1TypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv1TypeCounter
	// setMsg unmarshals PatternFlowIgmpv1TypeCounter from protobuf object *otg.PatternFlowIgmpv1TypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv1TypeCounter) PatternFlowIgmpv1TypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv1TypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv1TypeCounter
	// validate validates PatternFlowIgmpv1TypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv1TypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIgmpv1TypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIgmpv1TypeCounter
	SetStart(value uint32) PatternFlowIgmpv1TypeCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv1TypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIgmpv1TypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIgmpv1TypeCounter
	SetStep(value uint32) PatternFlowIgmpv1TypeCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv1TypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv1TypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv1TypeCounter
	SetCount(value uint32) PatternFlowIgmpv1TypeCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv1TypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv1TypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv1TypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIgmpv1TypeCounter object
func (obj *patternFlowIgmpv1TypeCounter) SetStart(value uint32) PatternFlowIgmpv1TypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv1TypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv1TypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIgmpv1TypeCounter object
func (obj *patternFlowIgmpv1TypeCounter) SetStep(value uint32) PatternFlowIgmpv1TypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv1TypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv1TypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv1TypeCounter object
func (obj *patternFlowIgmpv1TypeCounter) SetCount(value uint32) PatternFlowIgmpv1TypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv1TypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1TypeCounter.Start <= 15 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1TypeCounter.Step <= 15 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1TypeCounter.Count <= 15 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIgmpv1TypeCounter) setDefault() {
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

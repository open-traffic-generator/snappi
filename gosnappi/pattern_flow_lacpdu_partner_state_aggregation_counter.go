package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerStateAggregationCounter *****
type patternFlowLacpduPartnerStateAggregationCounter struct {
	validation
	obj          *otg.PatternFlowLacpduPartnerStateAggregationCounter
	marshaller   marshalPatternFlowLacpduPartnerStateAggregationCounter
	unMarshaller unMarshalPatternFlowLacpduPartnerStateAggregationCounter
}

func NewPatternFlowLacpduPartnerStateAggregationCounter() PatternFlowLacpduPartnerStateAggregationCounter {
	obj := patternFlowLacpduPartnerStateAggregationCounter{obj: &otg.PatternFlowLacpduPartnerStateAggregationCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerStateAggregationCounter) msg() *otg.PatternFlowLacpduPartnerStateAggregationCounter {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerStateAggregationCounter) setMsg(msg *otg.PatternFlowLacpduPartnerStateAggregationCounter) PatternFlowLacpduPartnerStateAggregationCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerStateAggregationCounter struct {
	obj *patternFlowLacpduPartnerStateAggregationCounter
}

type marshalPatternFlowLacpduPartnerStateAggregationCounter interface {
	// ToProto marshals PatternFlowLacpduPartnerStateAggregationCounter to protobuf object *otg.PatternFlowLacpduPartnerStateAggregationCounter
	ToProto() (*otg.PatternFlowLacpduPartnerStateAggregationCounter, error)
	// ToPbText marshals PatternFlowLacpduPartnerStateAggregationCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerStateAggregationCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerStateAggregationCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerStateAggregationCounter struct {
	obj *patternFlowLacpduPartnerStateAggregationCounter
}

type unMarshalPatternFlowLacpduPartnerStateAggregationCounter interface {
	// FromProto unmarshals PatternFlowLacpduPartnerStateAggregationCounter from protobuf object *otg.PatternFlowLacpduPartnerStateAggregationCounter
	FromProto(msg *otg.PatternFlowLacpduPartnerStateAggregationCounter) (PatternFlowLacpduPartnerStateAggregationCounter, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerStateAggregationCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerStateAggregationCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerStateAggregationCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerStateAggregationCounter) Marshal() marshalPatternFlowLacpduPartnerStateAggregationCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerStateAggregationCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerStateAggregationCounter) Unmarshal() unMarshalPatternFlowLacpduPartnerStateAggregationCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerStateAggregationCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerStateAggregationCounter) ToProto() (*otg.PatternFlowLacpduPartnerStateAggregationCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerStateAggregationCounter) FromProto(msg *otg.PatternFlowLacpduPartnerStateAggregationCounter) (PatternFlowLacpduPartnerStateAggregationCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerStateAggregationCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateAggregationCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateAggregationCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateAggregationCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateAggregationCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateAggregationCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerStateAggregationCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateAggregationCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateAggregationCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerStateAggregationCounter) Clone() (PatternFlowLacpduPartnerStateAggregationCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerStateAggregationCounter()
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

// PatternFlowLacpduPartnerStateAggregationCounter is integer counter pattern
type PatternFlowLacpduPartnerStateAggregationCounter interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerStateAggregationCounter to protobuf object *otg.PatternFlowLacpduPartnerStateAggregationCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerStateAggregationCounter
	// setMsg unmarshals PatternFlowLacpduPartnerStateAggregationCounter from protobuf object *otg.PatternFlowLacpduPartnerStateAggregationCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerStateAggregationCounter) PatternFlowLacpduPartnerStateAggregationCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerStateAggregationCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerStateAggregationCounter
	// validate validates PatternFlowLacpduPartnerStateAggregationCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerStateAggregationCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduPartnerStateAggregationCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduPartnerStateAggregationCounter
	SetStart(value uint32) PatternFlowLacpduPartnerStateAggregationCounter
	// HasStart checks if Start has been set in PatternFlowLacpduPartnerStateAggregationCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduPartnerStateAggregationCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduPartnerStateAggregationCounter
	SetStep(value uint32) PatternFlowLacpduPartnerStateAggregationCounter
	// HasStep checks if Step has been set in PatternFlowLacpduPartnerStateAggregationCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduPartnerStateAggregationCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduPartnerStateAggregationCounter
	SetCount(value uint32) PatternFlowLacpduPartnerStateAggregationCounter
	// HasCount checks if Count has been set in PatternFlowLacpduPartnerStateAggregationCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerStateAggregationCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerStateAggregationCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduPartnerStateAggregationCounter object
func (obj *patternFlowLacpduPartnerStateAggregationCounter) SetStart(value uint32) PatternFlowLacpduPartnerStateAggregationCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerStateAggregationCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerStateAggregationCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduPartnerStateAggregationCounter object
func (obj *patternFlowLacpduPartnerStateAggregationCounter) SetStep(value uint32) PatternFlowLacpduPartnerStateAggregationCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerStateAggregationCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerStateAggregationCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduPartnerStateAggregationCounter object
func (obj *patternFlowLacpduPartnerStateAggregationCounter) SetCount(value uint32) PatternFlowLacpduPartnerStateAggregationCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduPartnerStateAggregationCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateAggregationCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateAggregationCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateAggregationCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduPartnerStateAggregationCounter) setDefault() {
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

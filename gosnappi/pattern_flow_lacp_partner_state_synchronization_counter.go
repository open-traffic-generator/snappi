package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerStateSynchronizationCounter *****
type patternFlowLacpPartnerStateSynchronizationCounter struct {
	validation
	obj          *otg.PatternFlowLacpPartnerStateSynchronizationCounter
	marshaller   marshalPatternFlowLacpPartnerStateSynchronizationCounter
	unMarshaller unMarshalPatternFlowLacpPartnerStateSynchronizationCounter
}

func NewPatternFlowLacpPartnerStateSynchronizationCounter() PatternFlowLacpPartnerStateSynchronizationCounter {
	obj := patternFlowLacpPartnerStateSynchronizationCounter{obj: &otg.PatternFlowLacpPartnerStateSynchronizationCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerStateSynchronizationCounter) msg() *otg.PatternFlowLacpPartnerStateSynchronizationCounter {
	return obj.obj
}

func (obj *patternFlowLacpPartnerStateSynchronizationCounter) setMsg(msg *otg.PatternFlowLacpPartnerStateSynchronizationCounter) PatternFlowLacpPartnerStateSynchronizationCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerStateSynchronizationCounter struct {
	obj *patternFlowLacpPartnerStateSynchronizationCounter
}

type marshalPatternFlowLacpPartnerStateSynchronizationCounter interface {
	// ToProto marshals PatternFlowLacpPartnerStateSynchronizationCounter to protobuf object *otg.PatternFlowLacpPartnerStateSynchronizationCounter
	ToProto() (*otg.PatternFlowLacpPartnerStateSynchronizationCounter, error)
	// ToPbText marshals PatternFlowLacpPartnerStateSynchronizationCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerStateSynchronizationCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerStateSynchronizationCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerStateSynchronizationCounter struct {
	obj *patternFlowLacpPartnerStateSynchronizationCounter
}

type unMarshalPatternFlowLacpPartnerStateSynchronizationCounter interface {
	// FromProto unmarshals PatternFlowLacpPartnerStateSynchronizationCounter from protobuf object *otg.PatternFlowLacpPartnerStateSynchronizationCounter
	FromProto(msg *otg.PatternFlowLacpPartnerStateSynchronizationCounter) (PatternFlowLacpPartnerStateSynchronizationCounter, error)
	// FromPbText unmarshals PatternFlowLacpPartnerStateSynchronizationCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerStateSynchronizationCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerStateSynchronizationCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerStateSynchronizationCounter) Marshal() marshalPatternFlowLacpPartnerStateSynchronizationCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerStateSynchronizationCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerStateSynchronizationCounter) Unmarshal() unMarshalPatternFlowLacpPartnerStateSynchronizationCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerStateSynchronizationCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerStateSynchronizationCounter) ToProto() (*otg.PatternFlowLacpPartnerStateSynchronizationCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerStateSynchronizationCounter) FromProto(msg *otg.PatternFlowLacpPartnerStateSynchronizationCounter) (PatternFlowLacpPartnerStateSynchronizationCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerStateSynchronizationCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateSynchronizationCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateSynchronizationCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateSynchronizationCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateSynchronizationCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateSynchronizationCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerStateSynchronizationCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateSynchronizationCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateSynchronizationCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerStateSynchronizationCounter) Clone() (PatternFlowLacpPartnerStateSynchronizationCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerStateSynchronizationCounter()
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

// PatternFlowLacpPartnerStateSynchronizationCounter is integer counter pattern
type PatternFlowLacpPartnerStateSynchronizationCounter interface {
	Validation
	// msg marshals PatternFlowLacpPartnerStateSynchronizationCounter to protobuf object *otg.PatternFlowLacpPartnerStateSynchronizationCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerStateSynchronizationCounter
	// setMsg unmarshals PatternFlowLacpPartnerStateSynchronizationCounter from protobuf object *otg.PatternFlowLacpPartnerStateSynchronizationCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerStateSynchronizationCounter) PatternFlowLacpPartnerStateSynchronizationCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerStateSynchronizationCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerStateSynchronizationCounter
	// validate validates PatternFlowLacpPartnerStateSynchronizationCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerStateSynchronizationCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpPartnerStateSynchronizationCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpPartnerStateSynchronizationCounter
	SetStart(value uint32) PatternFlowLacpPartnerStateSynchronizationCounter
	// HasStart checks if Start has been set in PatternFlowLacpPartnerStateSynchronizationCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpPartnerStateSynchronizationCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpPartnerStateSynchronizationCounter
	SetStep(value uint32) PatternFlowLacpPartnerStateSynchronizationCounter
	// HasStep checks if Step has been set in PatternFlowLacpPartnerStateSynchronizationCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpPartnerStateSynchronizationCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpPartnerStateSynchronizationCounter
	SetCount(value uint32) PatternFlowLacpPartnerStateSynchronizationCounter
	// HasCount checks if Count has been set in PatternFlowLacpPartnerStateSynchronizationCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerStateSynchronizationCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerStateSynchronizationCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpPartnerStateSynchronizationCounter object
func (obj *patternFlowLacpPartnerStateSynchronizationCounter) SetStart(value uint32) PatternFlowLacpPartnerStateSynchronizationCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerStateSynchronizationCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerStateSynchronizationCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpPartnerStateSynchronizationCounter object
func (obj *patternFlowLacpPartnerStateSynchronizationCounter) SetStep(value uint32) PatternFlowLacpPartnerStateSynchronizationCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerStateSynchronizationCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerStateSynchronizationCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpPartnerStateSynchronizationCounter object
func (obj *patternFlowLacpPartnerStateSynchronizationCounter) SetCount(value uint32) PatternFlowLacpPartnerStateSynchronizationCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpPartnerStateSynchronizationCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateSynchronizationCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateSynchronizationCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateSynchronizationCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpPartnerStateSynchronizationCounter) setDefault() {
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

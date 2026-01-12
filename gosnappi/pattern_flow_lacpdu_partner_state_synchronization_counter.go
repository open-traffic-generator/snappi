package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerStateSynchronizationCounter *****
type patternFlowLacpduPartnerStateSynchronizationCounter struct {
	validation
	obj          *otg.PatternFlowLacpduPartnerStateSynchronizationCounter
	marshaller   marshalPatternFlowLacpduPartnerStateSynchronizationCounter
	unMarshaller unMarshalPatternFlowLacpduPartnerStateSynchronizationCounter
}

func NewPatternFlowLacpduPartnerStateSynchronizationCounter() PatternFlowLacpduPartnerStateSynchronizationCounter {
	obj := patternFlowLacpduPartnerStateSynchronizationCounter{obj: &otg.PatternFlowLacpduPartnerStateSynchronizationCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerStateSynchronizationCounter) msg() *otg.PatternFlowLacpduPartnerStateSynchronizationCounter {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerStateSynchronizationCounter) setMsg(msg *otg.PatternFlowLacpduPartnerStateSynchronizationCounter) PatternFlowLacpduPartnerStateSynchronizationCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerStateSynchronizationCounter struct {
	obj *patternFlowLacpduPartnerStateSynchronizationCounter
}

type marshalPatternFlowLacpduPartnerStateSynchronizationCounter interface {
	// ToProto marshals PatternFlowLacpduPartnerStateSynchronizationCounter to protobuf object *otg.PatternFlowLacpduPartnerStateSynchronizationCounter
	ToProto() (*otg.PatternFlowLacpduPartnerStateSynchronizationCounter, error)
	// ToPbText marshals PatternFlowLacpduPartnerStateSynchronizationCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerStateSynchronizationCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerStateSynchronizationCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerStateSynchronizationCounter struct {
	obj *patternFlowLacpduPartnerStateSynchronizationCounter
}

type unMarshalPatternFlowLacpduPartnerStateSynchronizationCounter interface {
	// FromProto unmarshals PatternFlowLacpduPartnerStateSynchronizationCounter from protobuf object *otg.PatternFlowLacpduPartnerStateSynchronizationCounter
	FromProto(msg *otg.PatternFlowLacpduPartnerStateSynchronizationCounter) (PatternFlowLacpduPartnerStateSynchronizationCounter, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerStateSynchronizationCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerStateSynchronizationCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerStateSynchronizationCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerStateSynchronizationCounter) Marshal() marshalPatternFlowLacpduPartnerStateSynchronizationCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerStateSynchronizationCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerStateSynchronizationCounter) Unmarshal() unMarshalPatternFlowLacpduPartnerStateSynchronizationCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerStateSynchronizationCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerStateSynchronizationCounter) ToProto() (*otg.PatternFlowLacpduPartnerStateSynchronizationCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerStateSynchronizationCounter) FromProto(msg *otg.PatternFlowLacpduPartnerStateSynchronizationCounter) (PatternFlowLacpduPartnerStateSynchronizationCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerStateSynchronizationCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateSynchronizationCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateSynchronizationCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateSynchronizationCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateSynchronizationCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateSynchronizationCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerStateSynchronizationCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateSynchronizationCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateSynchronizationCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerStateSynchronizationCounter) Clone() (PatternFlowLacpduPartnerStateSynchronizationCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerStateSynchronizationCounter()
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

// PatternFlowLacpduPartnerStateSynchronizationCounter is integer counter pattern
type PatternFlowLacpduPartnerStateSynchronizationCounter interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerStateSynchronizationCounter to protobuf object *otg.PatternFlowLacpduPartnerStateSynchronizationCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerStateSynchronizationCounter
	// setMsg unmarshals PatternFlowLacpduPartnerStateSynchronizationCounter from protobuf object *otg.PatternFlowLacpduPartnerStateSynchronizationCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerStateSynchronizationCounter) PatternFlowLacpduPartnerStateSynchronizationCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerStateSynchronizationCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerStateSynchronizationCounter
	// validate validates PatternFlowLacpduPartnerStateSynchronizationCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerStateSynchronizationCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduPartnerStateSynchronizationCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduPartnerStateSynchronizationCounter
	SetStart(value uint32) PatternFlowLacpduPartnerStateSynchronizationCounter
	// HasStart checks if Start has been set in PatternFlowLacpduPartnerStateSynchronizationCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduPartnerStateSynchronizationCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduPartnerStateSynchronizationCounter
	SetStep(value uint32) PatternFlowLacpduPartnerStateSynchronizationCounter
	// HasStep checks if Step has been set in PatternFlowLacpduPartnerStateSynchronizationCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduPartnerStateSynchronizationCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduPartnerStateSynchronizationCounter
	SetCount(value uint32) PatternFlowLacpduPartnerStateSynchronizationCounter
	// HasCount checks if Count has been set in PatternFlowLacpduPartnerStateSynchronizationCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerStateSynchronizationCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerStateSynchronizationCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduPartnerStateSynchronizationCounter object
func (obj *patternFlowLacpduPartnerStateSynchronizationCounter) SetStart(value uint32) PatternFlowLacpduPartnerStateSynchronizationCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerStateSynchronizationCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerStateSynchronizationCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduPartnerStateSynchronizationCounter object
func (obj *patternFlowLacpduPartnerStateSynchronizationCounter) SetStep(value uint32) PatternFlowLacpduPartnerStateSynchronizationCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerStateSynchronizationCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerStateSynchronizationCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduPartnerStateSynchronizationCounter object
func (obj *patternFlowLacpduPartnerStateSynchronizationCounter) SetCount(value uint32) PatternFlowLacpduPartnerStateSynchronizationCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduPartnerStateSynchronizationCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateSynchronizationCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateSynchronizationCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateSynchronizationCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduPartnerStateSynchronizationCounter) setDefault() {
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

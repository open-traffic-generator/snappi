package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerStateCollectingCounter *****
type patternFlowLacpPartnerStateCollectingCounter struct {
	validation
	obj          *otg.PatternFlowLacpPartnerStateCollectingCounter
	marshaller   marshalPatternFlowLacpPartnerStateCollectingCounter
	unMarshaller unMarshalPatternFlowLacpPartnerStateCollectingCounter
}

func NewPatternFlowLacpPartnerStateCollectingCounter() PatternFlowLacpPartnerStateCollectingCounter {
	obj := patternFlowLacpPartnerStateCollectingCounter{obj: &otg.PatternFlowLacpPartnerStateCollectingCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerStateCollectingCounter) msg() *otg.PatternFlowLacpPartnerStateCollectingCounter {
	return obj.obj
}

func (obj *patternFlowLacpPartnerStateCollectingCounter) setMsg(msg *otg.PatternFlowLacpPartnerStateCollectingCounter) PatternFlowLacpPartnerStateCollectingCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerStateCollectingCounter struct {
	obj *patternFlowLacpPartnerStateCollectingCounter
}

type marshalPatternFlowLacpPartnerStateCollectingCounter interface {
	// ToProto marshals PatternFlowLacpPartnerStateCollectingCounter to protobuf object *otg.PatternFlowLacpPartnerStateCollectingCounter
	ToProto() (*otg.PatternFlowLacpPartnerStateCollectingCounter, error)
	// ToPbText marshals PatternFlowLacpPartnerStateCollectingCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerStateCollectingCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerStateCollectingCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerStateCollectingCounter struct {
	obj *patternFlowLacpPartnerStateCollectingCounter
}

type unMarshalPatternFlowLacpPartnerStateCollectingCounter interface {
	// FromProto unmarshals PatternFlowLacpPartnerStateCollectingCounter from protobuf object *otg.PatternFlowLacpPartnerStateCollectingCounter
	FromProto(msg *otg.PatternFlowLacpPartnerStateCollectingCounter) (PatternFlowLacpPartnerStateCollectingCounter, error)
	// FromPbText unmarshals PatternFlowLacpPartnerStateCollectingCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerStateCollectingCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerStateCollectingCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerStateCollectingCounter) Marshal() marshalPatternFlowLacpPartnerStateCollectingCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerStateCollectingCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerStateCollectingCounter) Unmarshal() unMarshalPatternFlowLacpPartnerStateCollectingCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerStateCollectingCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerStateCollectingCounter) ToProto() (*otg.PatternFlowLacpPartnerStateCollectingCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerStateCollectingCounter) FromProto(msg *otg.PatternFlowLacpPartnerStateCollectingCounter) (PatternFlowLacpPartnerStateCollectingCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerStateCollectingCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateCollectingCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateCollectingCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateCollectingCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateCollectingCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateCollectingCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerStateCollectingCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateCollectingCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateCollectingCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerStateCollectingCounter) Clone() (PatternFlowLacpPartnerStateCollectingCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerStateCollectingCounter()
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

// PatternFlowLacpPartnerStateCollectingCounter is integer counter pattern
type PatternFlowLacpPartnerStateCollectingCounter interface {
	Validation
	// msg marshals PatternFlowLacpPartnerStateCollectingCounter to protobuf object *otg.PatternFlowLacpPartnerStateCollectingCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerStateCollectingCounter
	// setMsg unmarshals PatternFlowLacpPartnerStateCollectingCounter from protobuf object *otg.PatternFlowLacpPartnerStateCollectingCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerStateCollectingCounter) PatternFlowLacpPartnerStateCollectingCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerStateCollectingCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerStateCollectingCounter
	// validate validates PatternFlowLacpPartnerStateCollectingCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerStateCollectingCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpPartnerStateCollectingCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpPartnerStateCollectingCounter
	SetStart(value uint32) PatternFlowLacpPartnerStateCollectingCounter
	// HasStart checks if Start has been set in PatternFlowLacpPartnerStateCollectingCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpPartnerStateCollectingCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpPartnerStateCollectingCounter
	SetStep(value uint32) PatternFlowLacpPartnerStateCollectingCounter
	// HasStep checks if Step has been set in PatternFlowLacpPartnerStateCollectingCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpPartnerStateCollectingCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpPartnerStateCollectingCounter
	SetCount(value uint32) PatternFlowLacpPartnerStateCollectingCounter
	// HasCount checks if Count has been set in PatternFlowLacpPartnerStateCollectingCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerStateCollectingCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerStateCollectingCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpPartnerStateCollectingCounter object
func (obj *patternFlowLacpPartnerStateCollectingCounter) SetStart(value uint32) PatternFlowLacpPartnerStateCollectingCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerStateCollectingCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerStateCollectingCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpPartnerStateCollectingCounter object
func (obj *patternFlowLacpPartnerStateCollectingCounter) SetStep(value uint32) PatternFlowLacpPartnerStateCollectingCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerStateCollectingCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerStateCollectingCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpPartnerStateCollectingCounter object
func (obj *patternFlowLacpPartnerStateCollectingCounter) SetCount(value uint32) PatternFlowLacpPartnerStateCollectingCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpPartnerStateCollectingCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateCollectingCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateCollectingCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateCollectingCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpPartnerStateCollectingCounter) setDefault() {
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

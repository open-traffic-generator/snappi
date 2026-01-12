package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerStateCollectingCounter *****
type patternFlowLacpduPartnerStateCollectingCounter struct {
	validation
	obj          *otg.PatternFlowLacpduPartnerStateCollectingCounter
	marshaller   marshalPatternFlowLacpduPartnerStateCollectingCounter
	unMarshaller unMarshalPatternFlowLacpduPartnerStateCollectingCounter
}

func NewPatternFlowLacpduPartnerStateCollectingCounter() PatternFlowLacpduPartnerStateCollectingCounter {
	obj := patternFlowLacpduPartnerStateCollectingCounter{obj: &otg.PatternFlowLacpduPartnerStateCollectingCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerStateCollectingCounter) msg() *otg.PatternFlowLacpduPartnerStateCollectingCounter {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerStateCollectingCounter) setMsg(msg *otg.PatternFlowLacpduPartnerStateCollectingCounter) PatternFlowLacpduPartnerStateCollectingCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerStateCollectingCounter struct {
	obj *patternFlowLacpduPartnerStateCollectingCounter
}

type marshalPatternFlowLacpduPartnerStateCollectingCounter interface {
	// ToProto marshals PatternFlowLacpduPartnerStateCollectingCounter to protobuf object *otg.PatternFlowLacpduPartnerStateCollectingCounter
	ToProto() (*otg.PatternFlowLacpduPartnerStateCollectingCounter, error)
	// ToPbText marshals PatternFlowLacpduPartnerStateCollectingCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerStateCollectingCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerStateCollectingCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerStateCollectingCounter struct {
	obj *patternFlowLacpduPartnerStateCollectingCounter
}

type unMarshalPatternFlowLacpduPartnerStateCollectingCounter interface {
	// FromProto unmarshals PatternFlowLacpduPartnerStateCollectingCounter from protobuf object *otg.PatternFlowLacpduPartnerStateCollectingCounter
	FromProto(msg *otg.PatternFlowLacpduPartnerStateCollectingCounter) (PatternFlowLacpduPartnerStateCollectingCounter, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerStateCollectingCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerStateCollectingCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerStateCollectingCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerStateCollectingCounter) Marshal() marshalPatternFlowLacpduPartnerStateCollectingCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerStateCollectingCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerStateCollectingCounter) Unmarshal() unMarshalPatternFlowLacpduPartnerStateCollectingCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerStateCollectingCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerStateCollectingCounter) ToProto() (*otg.PatternFlowLacpduPartnerStateCollectingCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerStateCollectingCounter) FromProto(msg *otg.PatternFlowLacpduPartnerStateCollectingCounter) (PatternFlowLacpduPartnerStateCollectingCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerStateCollectingCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateCollectingCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateCollectingCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateCollectingCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateCollectingCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateCollectingCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerStateCollectingCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateCollectingCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateCollectingCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerStateCollectingCounter) Clone() (PatternFlowLacpduPartnerStateCollectingCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerStateCollectingCounter()
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

// PatternFlowLacpduPartnerStateCollectingCounter is integer counter pattern
type PatternFlowLacpduPartnerStateCollectingCounter interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerStateCollectingCounter to protobuf object *otg.PatternFlowLacpduPartnerStateCollectingCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerStateCollectingCounter
	// setMsg unmarshals PatternFlowLacpduPartnerStateCollectingCounter from protobuf object *otg.PatternFlowLacpduPartnerStateCollectingCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerStateCollectingCounter) PatternFlowLacpduPartnerStateCollectingCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerStateCollectingCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerStateCollectingCounter
	// validate validates PatternFlowLacpduPartnerStateCollectingCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerStateCollectingCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduPartnerStateCollectingCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduPartnerStateCollectingCounter
	SetStart(value uint32) PatternFlowLacpduPartnerStateCollectingCounter
	// HasStart checks if Start has been set in PatternFlowLacpduPartnerStateCollectingCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduPartnerStateCollectingCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduPartnerStateCollectingCounter
	SetStep(value uint32) PatternFlowLacpduPartnerStateCollectingCounter
	// HasStep checks if Step has been set in PatternFlowLacpduPartnerStateCollectingCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduPartnerStateCollectingCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduPartnerStateCollectingCounter
	SetCount(value uint32) PatternFlowLacpduPartnerStateCollectingCounter
	// HasCount checks if Count has been set in PatternFlowLacpduPartnerStateCollectingCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerStateCollectingCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerStateCollectingCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduPartnerStateCollectingCounter object
func (obj *patternFlowLacpduPartnerStateCollectingCounter) SetStart(value uint32) PatternFlowLacpduPartnerStateCollectingCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerStateCollectingCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerStateCollectingCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduPartnerStateCollectingCounter object
func (obj *patternFlowLacpduPartnerStateCollectingCounter) SetStep(value uint32) PatternFlowLacpduPartnerStateCollectingCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerStateCollectingCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerStateCollectingCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduPartnerStateCollectingCounter object
func (obj *patternFlowLacpduPartnerStateCollectingCounter) SetCount(value uint32) PatternFlowLacpduPartnerStateCollectingCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduPartnerStateCollectingCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateCollectingCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateCollectingCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateCollectingCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduPartnerStateCollectingCounter) setDefault() {
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

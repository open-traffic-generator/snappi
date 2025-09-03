package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerStateActivityCounter *****
type patternFlowLacpPartnerStateActivityCounter struct {
	validation
	obj          *otg.PatternFlowLacpPartnerStateActivityCounter
	marshaller   marshalPatternFlowLacpPartnerStateActivityCounter
	unMarshaller unMarshalPatternFlowLacpPartnerStateActivityCounter
}

func NewPatternFlowLacpPartnerStateActivityCounter() PatternFlowLacpPartnerStateActivityCounter {
	obj := patternFlowLacpPartnerStateActivityCounter{obj: &otg.PatternFlowLacpPartnerStateActivityCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerStateActivityCounter) msg() *otg.PatternFlowLacpPartnerStateActivityCounter {
	return obj.obj
}

func (obj *patternFlowLacpPartnerStateActivityCounter) setMsg(msg *otg.PatternFlowLacpPartnerStateActivityCounter) PatternFlowLacpPartnerStateActivityCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerStateActivityCounter struct {
	obj *patternFlowLacpPartnerStateActivityCounter
}

type marshalPatternFlowLacpPartnerStateActivityCounter interface {
	// ToProto marshals PatternFlowLacpPartnerStateActivityCounter to protobuf object *otg.PatternFlowLacpPartnerStateActivityCounter
	ToProto() (*otg.PatternFlowLacpPartnerStateActivityCounter, error)
	// ToPbText marshals PatternFlowLacpPartnerStateActivityCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerStateActivityCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerStateActivityCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerStateActivityCounter struct {
	obj *patternFlowLacpPartnerStateActivityCounter
}

type unMarshalPatternFlowLacpPartnerStateActivityCounter interface {
	// FromProto unmarshals PatternFlowLacpPartnerStateActivityCounter from protobuf object *otg.PatternFlowLacpPartnerStateActivityCounter
	FromProto(msg *otg.PatternFlowLacpPartnerStateActivityCounter) (PatternFlowLacpPartnerStateActivityCounter, error)
	// FromPbText unmarshals PatternFlowLacpPartnerStateActivityCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerStateActivityCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerStateActivityCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerStateActivityCounter) Marshal() marshalPatternFlowLacpPartnerStateActivityCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerStateActivityCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerStateActivityCounter) Unmarshal() unMarshalPatternFlowLacpPartnerStateActivityCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerStateActivityCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerStateActivityCounter) ToProto() (*otg.PatternFlowLacpPartnerStateActivityCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerStateActivityCounter) FromProto(msg *otg.PatternFlowLacpPartnerStateActivityCounter) (PatternFlowLacpPartnerStateActivityCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerStateActivityCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateActivityCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateActivityCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateActivityCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateActivityCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateActivityCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerStateActivityCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateActivityCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateActivityCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerStateActivityCounter) Clone() (PatternFlowLacpPartnerStateActivityCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerStateActivityCounter()
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

// PatternFlowLacpPartnerStateActivityCounter is integer counter pattern
type PatternFlowLacpPartnerStateActivityCounter interface {
	Validation
	// msg marshals PatternFlowLacpPartnerStateActivityCounter to protobuf object *otg.PatternFlowLacpPartnerStateActivityCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerStateActivityCounter
	// setMsg unmarshals PatternFlowLacpPartnerStateActivityCounter from protobuf object *otg.PatternFlowLacpPartnerStateActivityCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerStateActivityCounter) PatternFlowLacpPartnerStateActivityCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerStateActivityCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerStateActivityCounter
	// validate validates PatternFlowLacpPartnerStateActivityCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerStateActivityCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpPartnerStateActivityCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpPartnerStateActivityCounter
	SetStart(value uint32) PatternFlowLacpPartnerStateActivityCounter
	// HasStart checks if Start has been set in PatternFlowLacpPartnerStateActivityCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpPartnerStateActivityCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpPartnerStateActivityCounter
	SetStep(value uint32) PatternFlowLacpPartnerStateActivityCounter
	// HasStep checks if Step has been set in PatternFlowLacpPartnerStateActivityCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpPartnerStateActivityCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpPartnerStateActivityCounter
	SetCount(value uint32) PatternFlowLacpPartnerStateActivityCounter
	// HasCount checks if Count has been set in PatternFlowLacpPartnerStateActivityCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerStateActivityCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerStateActivityCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpPartnerStateActivityCounter object
func (obj *patternFlowLacpPartnerStateActivityCounter) SetStart(value uint32) PatternFlowLacpPartnerStateActivityCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerStateActivityCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerStateActivityCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpPartnerStateActivityCounter object
func (obj *patternFlowLacpPartnerStateActivityCounter) SetStep(value uint32) PatternFlowLacpPartnerStateActivityCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerStateActivityCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerStateActivityCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpPartnerStateActivityCounter object
func (obj *patternFlowLacpPartnerStateActivityCounter) SetCount(value uint32) PatternFlowLacpPartnerStateActivityCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpPartnerStateActivityCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateActivityCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateActivityCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateActivityCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpPartnerStateActivityCounter) setDefault() {
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

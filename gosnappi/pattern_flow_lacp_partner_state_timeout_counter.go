package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerStateTimeoutCounter *****
type patternFlowLacpPartnerStateTimeoutCounter struct {
	validation
	obj          *otg.PatternFlowLacpPartnerStateTimeoutCounter
	marshaller   marshalPatternFlowLacpPartnerStateTimeoutCounter
	unMarshaller unMarshalPatternFlowLacpPartnerStateTimeoutCounter
}

func NewPatternFlowLacpPartnerStateTimeoutCounter() PatternFlowLacpPartnerStateTimeoutCounter {
	obj := patternFlowLacpPartnerStateTimeoutCounter{obj: &otg.PatternFlowLacpPartnerStateTimeoutCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerStateTimeoutCounter) msg() *otg.PatternFlowLacpPartnerStateTimeoutCounter {
	return obj.obj
}

func (obj *patternFlowLacpPartnerStateTimeoutCounter) setMsg(msg *otg.PatternFlowLacpPartnerStateTimeoutCounter) PatternFlowLacpPartnerStateTimeoutCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerStateTimeoutCounter struct {
	obj *patternFlowLacpPartnerStateTimeoutCounter
}

type marshalPatternFlowLacpPartnerStateTimeoutCounter interface {
	// ToProto marshals PatternFlowLacpPartnerStateTimeoutCounter to protobuf object *otg.PatternFlowLacpPartnerStateTimeoutCounter
	ToProto() (*otg.PatternFlowLacpPartnerStateTimeoutCounter, error)
	// ToPbText marshals PatternFlowLacpPartnerStateTimeoutCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerStateTimeoutCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerStateTimeoutCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerStateTimeoutCounter struct {
	obj *patternFlowLacpPartnerStateTimeoutCounter
}

type unMarshalPatternFlowLacpPartnerStateTimeoutCounter interface {
	// FromProto unmarshals PatternFlowLacpPartnerStateTimeoutCounter from protobuf object *otg.PatternFlowLacpPartnerStateTimeoutCounter
	FromProto(msg *otg.PatternFlowLacpPartnerStateTimeoutCounter) (PatternFlowLacpPartnerStateTimeoutCounter, error)
	// FromPbText unmarshals PatternFlowLacpPartnerStateTimeoutCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerStateTimeoutCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerStateTimeoutCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerStateTimeoutCounter) Marshal() marshalPatternFlowLacpPartnerStateTimeoutCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerStateTimeoutCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerStateTimeoutCounter) Unmarshal() unMarshalPatternFlowLacpPartnerStateTimeoutCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerStateTimeoutCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerStateTimeoutCounter) ToProto() (*otg.PatternFlowLacpPartnerStateTimeoutCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerStateTimeoutCounter) FromProto(msg *otg.PatternFlowLacpPartnerStateTimeoutCounter) (PatternFlowLacpPartnerStateTimeoutCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerStateTimeoutCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateTimeoutCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateTimeoutCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateTimeoutCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateTimeoutCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateTimeoutCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerStateTimeoutCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateTimeoutCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateTimeoutCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerStateTimeoutCounter) Clone() (PatternFlowLacpPartnerStateTimeoutCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerStateTimeoutCounter()
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

// PatternFlowLacpPartnerStateTimeoutCounter is integer counter pattern
type PatternFlowLacpPartnerStateTimeoutCounter interface {
	Validation
	// msg marshals PatternFlowLacpPartnerStateTimeoutCounter to protobuf object *otg.PatternFlowLacpPartnerStateTimeoutCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerStateTimeoutCounter
	// setMsg unmarshals PatternFlowLacpPartnerStateTimeoutCounter from protobuf object *otg.PatternFlowLacpPartnerStateTimeoutCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerStateTimeoutCounter) PatternFlowLacpPartnerStateTimeoutCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerStateTimeoutCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerStateTimeoutCounter
	// validate validates PatternFlowLacpPartnerStateTimeoutCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerStateTimeoutCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpPartnerStateTimeoutCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpPartnerStateTimeoutCounter
	SetStart(value uint32) PatternFlowLacpPartnerStateTimeoutCounter
	// HasStart checks if Start has been set in PatternFlowLacpPartnerStateTimeoutCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpPartnerStateTimeoutCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpPartnerStateTimeoutCounter
	SetStep(value uint32) PatternFlowLacpPartnerStateTimeoutCounter
	// HasStep checks if Step has been set in PatternFlowLacpPartnerStateTimeoutCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpPartnerStateTimeoutCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpPartnerStateTimeoutCounter
	SetCount(value uint32) PatternFlowLacpPartnerStateTimeoutCounter
	// HasCount checks if Count has been set in PatternFlowLacpPartnerStateTimeoutCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerStateTimeoutCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerStateTimeoutCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpPartnerStateTimeoutCounter object
func (obj *patternFlowLacpPartnerStateTimeoutCounter) SetStart(value uint32) PatternFlowLacpPartnerStateTimeoutCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerStateTimeoutCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerStateTimeoutCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpPartnerStateTimeoutCounter object
func (obj *patternFlowLacpPartnerStateTimeoutCounter) SetStep(value uint32) PatternFlowLacpPartnerStateTimeoutCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerStateTimeoutCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerStateTimeoutCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpPartnerStateTimeoutCounter object
func (obj *patternFlowLacpPartnerStateTimeoutCounter) SetCount(value uint32) PatternFlowLacpPartnerStateTimeoutCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpPartnerStateTimeoutCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateTimeoutCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateTimeoutCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateTimeoutCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpPartnerStateTimeoutCounter) setDefault() {
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

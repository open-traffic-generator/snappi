package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerReservedCounter *****
type patternFlowLacpPartnerReservedCounter struct {
	validation
	obj          *otg.PatternFlowLacpPartnerReservedCounter
	marshaller   marshalPatternFlowLacpPartnerReservedCounter
	unMarshaller unMarshalPatternFlowLacpPartnerReservedCounter
}

func NewPatternFlowLacpPartnerReservedCounter() PatternFlowLacpPartnerReservedCounter {
	obj := patternFlowLacpPartnerReservedCounter{obj: &otg.PatternFlowLacpPartnerReservedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerReservedCounter) msg() *otg.PatternFlowLacpPartnerReservedCounter {
	return obj.obj
}

func (obj *patternFlowLacpPartnerReservedCounter) setMsg(msg *otg.PatternFlowLacpPartnerReservedCounter) PatternFlowLacpPartnerReservedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerReservedCounter struct {
	obj *patternFlowLacpPartnerReservedCounter
}

type marshalPatternFlowLacpPartnerReservedCounter interface {
	// ToProto marshals PatternFlowLacpPartnerReservedCounter to protobuf object *otg.PatternFlowLacpPartnerReservedCounter
	ToProto() (*otg.PatternFlowLacpPartnerReservedCounter, error)
	// ToPbText marshals PatternFlowLacpPartnerReservedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerReservedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerReservedCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerReservedCounter struct {
	obj *patternFlowLacpPartnerReservedCounter
}

type unMarshalPatternFlowLacpPartnerReservedCounter interface {
	// FromProto unmarshals PatternFlowLacpPartnerReservedCounter from protobuf object *otg.PatternFlowLacpPartnerReservedCounter
	FromProto(msg *otg.PatternFlowLacpPartnerReservedCounter) (PatternFlowLacpPartnerReservedCounter, error)
	// FromPbText unmarshals PatternFlowLacpPartnerReservedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerReservedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerReservedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerReservedCounter) Marshal() marshalPatternFlowLacpPartnerReservedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerReservedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerReservedCounter) Unmarshal() unMarshalPatternFlowLacpPartnerReservedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerReservedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerReservedCounter) ToProto() (*otg.PatternFlowLacpPartnerReservedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerReservedCounter) FromProto(msg *otg.PatternFlowLacpPartnerReservedCounter) (PatternFlowLacpPartnerReservedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerReservedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerReservedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerReservedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerReservedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerReservedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerReservedCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerReservedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerReservedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerReservedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerReservedCounter) Clone() (PatternFlowLacpPartnerReservedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerReservedCounter()
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

// PatternFlowLacpPartnerReservedCounter is integer counter pattern
type PatternFlowLacpPartnerReservedCounter interface {
	Validation
	// msg marshals PatternFlowLacpPartnerReservedCounter to protobuf object *otg.PatternFlowLacpPartnerReservedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerReservedCounter
	// setMsg unmarshals PatternFlowLacpPartnerReservedCounter from protobuf object *otg.PatternFlowLacpPartnerReservedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerReservedCounter) PatternFlowLacpPartnerReservedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerReservedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerReservedCounter
	// validate validates PatternFlowLacpPartnerReservedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerReservedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpPartnerReservedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpPartnerReservedCounter
	SetStart(value uint32) PatternFlowLacpPartnerReservedCounter
	// HasStart checks if Start has been set in PatternFlowLacpPartnerReservedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpPartnerReservedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpPartnerReservedCounter
	SetStep(value uint32) PatternFlowLacpPartnerReservedCounter
	// HasStep checks if Step has been set in PatternFlowLacpPartnerReservedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpPartnerReservedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpPartnerReservedCounter
	SetCount(value uint32) PatternFlowLacpPartnerReservedCounter
	// HasCount checks if Count has been set in PatternFlowLacpPartnerReservedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerReservedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerReservedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpPartnerReservedCounter object
func (obj *patternFlowLacpPartnerReservedCounter) SetStart(value uint32) PatternFlowLacpPartnerReservedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerReservedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerReservedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpPartnerReservedCounter object
func (obj *patternFlowLacpPartnerReservedCounter) SetStep(value uint32) PatternFlowLacpPartnerReservedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerReservedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerReservedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpPartnerReservedCounter object
func (obj *patternFlowLacpPartnerReservedCounter) SetCount(value uint32) PatternFlowLacpPartnerReservedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpPartnerReservedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerReservedCounter.Start <= 16777215 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerReservedCounter.Step <= 16777215 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerReservedCounter.Count <= 16777215 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpPartnerReservedCounter) setDefault() {
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

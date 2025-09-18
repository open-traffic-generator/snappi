package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerReservedCounter *****
type patternFlowLacpduPartnerReservedCounter struct {
	validation
	obj          *otg.PatternFlowLacpduPartnerReservedCounter
	marshaller   marshalPatternFlowLacpduPartnerReservedCounter
	unMarshaller unMarshalPatternFlowLacpduPartnerReservedCounter
}

func NewPatternFlowLacpduPartnerReservedCounter() PatternFlowLacpduPartnerReservedCounter {
	obj := patternFlowLacpduPartnerReservedCounter{obj: &otg.PatternFlowLacpduPartnerReservedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerReservedCounter) msg() *otg.PatternFlowLacpduPartnerReservedCounter {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerReservedCounter) setMsg(msg *otg.PatternFlowLacpduPartnerReservedCounter) PatternFlowLacpduPartnerReservedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerReservedCounter struct {
	obj *patternFlowLacpduPartnerReservedCounter
}

type marshalPatternFlowLacpduPartnerReservedCounter interface {
	// ToProto marshals PatternFlowLacpduPartnerReservedCounter to protobuf object *otg.PatternFlowLacpduPartnerReservedCounter
	ToProto() (*otg.PatternFlowLacpduPartnerReservedCounter, error)
	// ToPbText marshals PatternFlowLacpduPartnerReservedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerReservedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerReservedCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerReservedCounter struct {
	obj *patternFlowLacpduPartnerReservedCounter
}

type unMarshalPatternFlowLacpduPartnerReservedCounter interface {
	// FromProto unmarshals PatternFlowLacpduPartnerReservedCounter from protobuf object *otg.PatternFlowLacpduPartnerReservedCounter
	FromProto(msg *otg.PatternFlowLacpduPartnerReservedCounter) (PatternFlowLacpduPartnerReservedCounter, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerReservedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerReservedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerReservedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerReservedCounter) Marshal() marshalPatternFlowLacpduPartnerReservedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerReservedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerReservedCounter) Unmarshal() unMarshalPatternFlowLacpduPartnerReservedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerReservedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerReservedCounter) ToProto() (*otg.PatternFlowLacpduPartnerReservedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerReservedCounter) FromProto(msg *otg.PatternFlowLacpduPartnerReservedCounter) (PatternFlowLacpduPartnerReservedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerReservedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerReservedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerReservedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerReservedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerReservedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerReservedCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerReservedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerReservedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerReservedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerReservedCounter) Clone() (PatternFlowLacpduPartnerReservedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerReservedCounter()
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

// PatternFlowLacpduPartnerReservedCounter is integer counter pattern
type PatternFlowLacpduPartnerReservedCounter interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerReservedCounter to protobuf object *otg.PatternFlowLacpduPartnerReservedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerReservedCounter
	// setMsg unmarshals PatternFlowLacpduPartnerReservedCounter from protobuf object *otg.PatternFlowLacpduPartnerReservedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerReservedCounter) PatternFlowLacpduPartnerReservedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerReservedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerReservedCounter
	// validate validates PatternFlowLacpduPartnerReservedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerReservedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduPartnerReservedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduPartnerReservedCounter
	SetStart(value uint32) PatternFlowLacpduPartnerReservedCounter
	// HasStart checks if Start has been set in PatternFlowLacpduPartnerReservedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduPartnerReservedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduPartnerReservedCounter
	SetStep(value uint32) PatternFlowLacpduPartnerReservedCounter
	// HasStep checks if Step has been set in PatternFlowLacpduPartnerReservedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduPartnerReservedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduPartnerReservedCounter
	SetCount(value uint32) PatternFlowLacpduPartnerReservedCounter
	// HasCount checks if Count has been set in PatternFlowLacpduPartnerReservedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerReservedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerReservedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduPartnerReservedCounter object
func (obj *patternFlowLacpduPartnerReservedCounter) SetStart(value uint32) PatternFlowLacpduPartnerReservedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerReservedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerReservedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduPartnerReservedCounter object
func (obj *patternFlowLacpduPartnerReservedCounter) SetStep(value uint32) PatternFlowLacpduPartnerReservedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerReservedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerReservedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduPartnerReservedCounter object
func (obj *patternFlowLacpduPartnerReservedCounter) SetCount(value uint32) PatternFlowLacpduPartnerReservedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduPartnerReservedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerReservedCounter.Start <= 16777215 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerReservedCounter.Step <= 16777215 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerReservedCounter.Count <= 16777215 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduPartnerReservedCounter) setDefault() {
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

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerPortNumberCounter *****
type patternFlowLacpPartnerPortNumberCounter struct {
	validation
	obj          *otg.PatternFlowLacpPartnerPortNumberCounter
	marshaller   marshalPatternFlowLacpPartnerPortNumberCounter
	unMarshaller unMarshalPatternFlowLacpPartnerPortNumberCounter
}

func NewPatternFlowLacpPartnerPortNumberCounter() PatternFlowLacpPartnerPortNumberCounter {
	obj := patternFlowLacpPartnerPortNumberCounter{obj: &otg.PatternFlowLacpPartnerPortNumberCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerPortNumberCounter) msg() *otg.PatternFlowLacpPartnerPortNumberCounter {
	return obj.obj
}

func (obj *patternFlowLacpPartnerPortNumberCounter) setMsg(msg *otg.PatternFlowLacpPartnerPortNumberCounter) PatternFlowLacpPartnerPortNumberCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerPortNumberCounter struct {
	obj *patternFlowLacpPartnerPortNumberCounter
}

type marshalPatternFlowLacpPartnerPortNumberCounter interface {
	// ToProto marshals PatternFlowLacpPartnerPortNumberCounter to protobuf object *otg.PatternFlowLacpPartnerPortNumberCounter
	ToProto() (*otg.PatternFlowLacpPartnerPortNumberCounter, error)
	// ToPbText marshals PatternFlowLacpPartnerPortNumberCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerPortNumberCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerPortNumberCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerPortNumberCounter struct {
	obj *patternFlowLacpPartnerPortNumberCounter
}

type unMarshalPatternFlowLacpPartnerPortNumberCounter interface {
	// FromProto unmarshals PatternFlowLacpPartnerPortNumberCounter from protobuf object *otg.PatternFlowLacpPartnerPortNumberCounter
	FromProto(msg *otg.PatternFlowLacpPartnerPortNumberCounter) (PatternFlowLacpPartnerPortNumberCounter, error)
	// FromPbText unmarshals PatternFlowLacpPartnerPortNumberCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerPortNumberCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerPortNumberCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerPortNumberCounter) Marshal() marshalPatternFlowLacpPartnerPortNumberCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerPortNumberCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerPortNumberCounter) Unmarshal() unMarshalPatternFlowLacpPartnerPortNumberCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerPortNumberCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerPortNumberCounter) ToProto() (*otg.PatternFlowLacpPartnerPortNumberCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerPortNumberCounter) FromProto(msg *otg.PatternFlowLacpPartnerPortNumberCounter) (PatternFlowLacpPartnerPortNumberCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerPortNumberCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerPortNumberCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerPortNumberCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerPortNumberCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerPortNumberCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerPortNumberCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerPortNumberCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerPortNumberCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerPortNumberCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerPortNumberCounter) Clone() (PatternFlowLacpPartnerPortNumberCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerPortNumberCounter()
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

// PatternFlowLacpPartnerPortNumberCounter is integer counter pattern
type PatternFlowLacpPartnerPortNumberCounter interface {
	Validation
	// msg marshals PatternFlowLacpPartnerPortNumberCounter to protobuf object *otg.PatternFlowLacpPartnerPortNumberCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerPortNumberCounter
	// setMsg unmarshals PatternFlowLacpPartnerPortNumberCounter from protobuf object *otg.PatternFlowLacpPartnerPortNumberCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerPortNumberCounter) PatternFlowLacpPartnerPortNumberCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerPortNumberCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerPortNumberCounter
	// validate validates PatternFlowLacpPartnerPortNumberCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerPortNumberCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpPartnerPortNumberCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpPartnerPortNumberCounter
	SetStart(value uint32) PatternFlowLacpPartnerPortNumberCounter
	// HasStart checks if Start has been set in PatternFlowLacpPartnerPortNumberCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpPartnerPortNumberCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpPartnerPortNumberCounter
	SetStep(value uint32) PatternFlowLacpPartnerPortNumberCounter
	// HasStep checks if Step has been set in PatternFlowLacpPartnerPortNumberCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpPartnerPortNumberCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpPartnerPortNumberCounter
	SetCount(value uint32) PatternFlowLacpPartnerPortNumberCounter
	// HasCount checks if Count has been set in PatternFlowLacpPartnerPortNumberCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerPortNumberCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerPortNumberCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpPartnerPortNumberCounter object
func (obj *patternFlowLacpPartnerPortNumberCounter) SetStart(value uint32) PatternFlowLacpPartnerPortNumberCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerPortNumberCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerPortNumberCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpPartnerPortNumberCounter object
func (obj *patternFlowLacpPartnerPortNumberCounter) SetStep(value uint32) PatternFlowLacpPartnerPortNumberCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerPortNumberCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerPortNumberCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpPartnerPortNumberCounter object
func (obj *patternFlowLacpPartnerPortNumberCounter) SetCount(value uint32) PatternFlowLacpPartnerPortNumberCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpPartnerPortNumberCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerPortNumberCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerPortNumberCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerPortNumberCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpPartnerPortNumberCounter) setDefault() {
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

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerSystemPriorityCounter *****
type patternFlowLacpPartnerSystemPriorityCounter struct {
	validation
	obj          *otg.PatternFlowLacpPartnerSystemPriorityCounter
	marshaller   marshalPatternFlowLacpPartnerSystemPriorityCounter
	unMarshaller unMarshalPatternFlowLacpPartnerSystemPriorityCounter
}

func NewPatternFlowLacpPartnerSystemPriorityCounter() PatternFlowLacpPartnerSystemPriorityCounter {
	obj := patternFlowLacpPartnerSystemPriorityCounter{obj: &otg.PatternFlowLacpPartnerSystemPriorityCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerSystemPriorityCounter) msg() *otg.PatternFlowLacpPartnerSystemPriorityCounter {
	return obj.obj
}

func (obj *patternFlowLacpPartnerSystemPriorityCounter) setMsg(msg *otg.PatternFlowLacpPartnerSystemPriorityCounter) PatternFlowLacpPartnerSystemPriorityCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerSystemPriorityCounter struct {
	obj *patternFlowLacpPartnerSystemPriorityCounter
}

type marshalPatternFlowLacpPartnerSystemPriorityCounter interface {
	// ToProto marshals PatternFlowLacpPartnerSystemPriorityCounter to protobuf object *otg.PatternFlowLacpPartnerSystemPriorityCounter
	ToProto() (*otg.PatternFlowLacpPartnerSystemPriorityCounter, error)
	// ToPbText marshals PatternFlowLacpPartnerSystemPriorityCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerSystemPriorityCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerSystemPriorityCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerSystemPriorityCounter struct {
	obj *patternFlowLacpPartnerSystemPriorityCounter
}

type unMarshalPatternFlowLacpPartnerSystemPriorityCounter interface {
	// FromProto unmarshals PatternFlowLacpPartnerSystemPriorityCounter from protobuf object *otg.PatternFlowLacpPartnerSystemPriorityCounter
	FromProto(msg *otg.PatternFlowLacpPartnerSystemPriorityCounter) (PatternFlowLacpPartnerSystemPriorityCounter, error)
	// FromPbText unmarshals PatternFlowLacpPartnerSystemPriorityCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerSystemPriorityCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerSystemPriorityCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerSystemPriorityCounter) Marshal() marshalPatternFlowLacpPartnerSystemPriorityCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerSystemPriorityCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerSystemPriorityCounter) Unmarshal() unMarshalPatternFlowLacpPartnerSystemPriorityCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerSystemPriorityCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerSystemPriorityCounter) ToProto() (*otg.PatternFlowLacpPartnerSystemPriorityCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerSystemPriorityCounter) FromProto(msg *otg.PatternFlowLacpPartnerSystemPriorityCounter) (PatternFlowLacpPartnerSystemPriorityCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerSystemPriorityCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerSystemPriorityCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerSystemPriorityCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerSystemPriorityCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerSystemPriorityCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerSystemPriorityCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerSystemPriorityCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerSystemPriorityCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerSystemPriorityCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerSystemPriorityCounter) Clone() (PatternFlowLacpPartnerSystemPriorityCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerSystemPriorityCounter()
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

// PatternFlowLacpPartnerSystemPriorityCounter is integer counter pattern
type PatternFlowLacpPartnerSystemPriorityCounter interface {
	Validation
	// msg marshals PatternFlowLacpPartnerSystemPriorityCounter to protobuf object *otg.PatternFlowLacpPartnerSystemPriorityCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerSystemPriorityCounter
	// setMsg unmarshals PatternFlowLacpPartnerSystemPriorityCounter from protobuf object *otg.PatternFlowLacpPartnerSystemPriorityCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerSystemPriorityCounter) PatternFlowLacpPartnerSystemPriorityCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerSystemPriorityCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerSystemPriorityCounter
	// validate validates PatternFlowLacpPartnerSystemPriorityCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerSystemPriorityCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpPartnerSystemPriorityCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpPartnerSystemPriorityCounter
	SetStart(value uint32) PatternFlowLacpPartnerSystemPriorityCounter
	// HasStart checks if Start has been set in PatternFlowLacpPartnerSystemPriorityCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpPartnerSystemPriorityCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpPartnerSystemPriorityCounter
	SetStep(value uint32) PatternFlowLacpPartnerSystemPriorityCounter
	// HasStep checks if Step has been set in PatternFlowLacpPartnerSystemPriorityCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpPartnerSystemPriorityCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpPartnerSystemPriorityCounter
	SetCount(value uint32) PatternFlowLacpPartnerSystemPriorityCounter
	// HasCount checks if Count has been set in PatternFlowLacpPartnerSystemPriorityCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerSystemPriorityCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerSystemPriorityCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpPartnerSystemPriorityCounter object
func (obj *patternFlowLacpPartnerSystemPriorityCounter) SetStart(value uint32) PatternFlowLacpPartnerSystemPriorityCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerSystemPriorityCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerSystemPriorityCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpPartnerSystemPriorityCounter object
func (obj *patternFlowLacpPartnerSystemPriorityCounter) SetStep(value uint32) PatternFlowLacpPartnerSystemPriorityCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerSystemPriorityCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerSystemPriorityCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpPartnerSystemPriorityCounter object
func (obj *patternFlowLacpPartnerSystemPriorityCounter) SetCount(value uint32) PatternFlowLacpPartnerSystemPriorityCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpPartnerSystemPriorityCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerSystemPriorityCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerSystemPriorityCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerSystemPriorityCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpPartnerSystemPriorityCounter) setDefault() {
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

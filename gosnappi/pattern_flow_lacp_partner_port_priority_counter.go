package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerPortPriorityCounter *****
type patternFlowLacpPartnerPortPriorityCounter struct {
	validation
	obj          *otg.PatternFlowLacpPartnerPortPriorityCounter
	marshaller   marshalPatternFlowLacpPartnerPortPriorityCounter
	unMarshaller unMarshalPatternFlowLacpPartnerPortPriorityCounter
}

func NewPatternFlowLacpPartnerPortPriorityCounter() PatternFlowLacpPartnerPortPriorityCounter {
	obj := patternFlowLacpPartnerPortPriorityCounter{obj: &otg.PatternFlowLacpPartnerPortPriorityCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerPortPriorityCounter) msg() *otg.PatternFlowLacpPartnerPortPriorityCounter {
	return obj.obj
}

func (obj *patternFlowLacpPartnerPortPriorityCounter) setMsg(msg *otg.PatternFlowLacpPartnerPortPriorityCounter) PatternFlowLacpPartnerPortPriorityCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerPortPriorityCounter struct {
	obj *patternFlowLacpPartnerPortPriorityCounter
}

type marshalPatternFlowLacpPartnerPortPriorityCounter interface {
	// ToProto marshals PatternFlowLacpPartnerPortPriorityCounter to protobuf object *otg.PatternFlowLacpPartnerPortPriorityCounter
	ToProto() (*otg.PatternFlowLacpPartnerPortPriorityCounter, error)
	// ToPbText marshals PatternFlowLacpPartnerPortPriorityCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerPortPriorityCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerPortPriorityCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerPortPriorityCounter struct {
	obj *patternFlowLacpPartnerPortPriorityCounter
}

type unMarshalPatternFlowLacpPartnerPortPriorityCounter interface {
	// FromProto unmarshals PatternFlowLacpPartnerPortPriorityCounter from protobuf object *otg.PatternFlowLacpPartnerPortPriorityCounter
	FromProto(msg *otg.PatternFlowLacpPartnerPortPriorityCounter) (PatternFlowLacpPartnerPortPriorityCounter, error)
	// FromPbText unmarshals PatternFlowLacpPartnerPortPriorityCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerPortPriorityCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerPortPriorityCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerPortPriorityCounter) Marshal() marshalPatternFlowLacpPartnerPortPriorityCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerPortPriorityCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerPortPriorityCounter) Unmarshal() unMarshalPatternFlowLacpPartnerPortPriorityCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerPortPriorityCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerPortPriorityCounter) ToProto() (*otg.PatternFlowLacpPartnerPortPriorityCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerPortPriorityCounter) FromProto(msg *otg.PatternFlowLacpPartnerPortPriorityCounter) (PatternFlowLacpPartnerPortPriorityCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerPortPriorityCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerPortPriorityCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerPortPriorityCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerPortPriorityCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerPortPriorityCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerPortPriorityCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerPortPriorityCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerPortPriorityCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerPortPriorityCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerPortPriorityCounter) Clone() (PatternFlowLacpPartnerPortPriorityCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerPortPriorityCounter()
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

// PatternFlowLacpPartnerPortPriorityCounter is integer counter pattern
type PatternFlowLacpPartnerPortPriorityCounter interface {
	Validation
	// msg marshals PatternFlowLacpPartnerPortPriorityCounter to protobuf object *otg.PatternFlowLacpPartnerPortPriorityCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerPortPriorityCounter
	// setMsg unmarshals PatternFlowLacpPartnerPortPriorityCounter from protobuf object *otg.PatternFlowLacpPartnerPortPriorityCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerPortPriorityCounter) PatternFlowLacpPartnerPortPriorityCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerPortPriorityCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerPortPriorityCounter
	// validate validates PatternFlowLacpPartnerPortPriorityCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerPortPriorityCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpPartnerPortPriorityCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpPartnerPortPriorityCounter
	SetStart(value uint32) PatternFlowLacpPartnerPortPriorityCounter
	// HasStart checks if Start has been set in PatternFlowLacpPartnerPortPriorityCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpPartnerPortPriorityCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpPartnerPortPriorityCounter
	SetStep(value uint32) PatternFlowLacpPartnerPortPriorityCounter
	// HasStep checks if Step has been set in PatternFlowLacpPartnerPortPriorityCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpPartnerPortPriorityCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpPartnerPortPriorityCounter
	SetCount(value uint32) PatternFlowLacpPartnerPortPriorityCounter
	// HasCount checks if Count has been set in PatternFlowLacpPartnerPortPriorityCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerPortPriorityCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerPortPriorityCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpPartnerPortPriorityCounter object
func (obj *patternFlowLacpPartnerPortPriorityCounter) SetStart(value uint32) PatternFlowLacpPartnerPortPriorityCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerPortPriorityCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerPortPriorityCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpPartnerPortPriorityCounter object
func (obj *patternFlowLacpPartnerPortPriorityCounter) SetStep(value uint32) PatternFlowLacpPartnerPortPriorityCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerPortPriorityCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerPortPriorityCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpPartnerPortPriorityCounter object
func (obj *patternFlowLacpPartnerPortPriorityCounter) SetCount(value uint32) PatternFlowLacpPartnerPortPriorityCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpPartnerPortPriorityCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerPortPriorityCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerPortPriorityCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerPortPriorityCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpPartnerPortPriorityCounter) setDefault() {
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

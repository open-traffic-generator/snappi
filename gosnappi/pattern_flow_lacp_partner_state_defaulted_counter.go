package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerStateDefaultedCounter *****
type patternFlowLacpPartnerStateDefaultedCounter struct {
	validation
	obj          *otg.PatternFlowLacpPartnerStateDefaultedCounter
	marshaller   marshalPatternFlowLacpPartnerStateDefaultedCounter
	unMarshaller unMarshalPatternFlowLacpPartnerStateDefaultedCounter
}

func NewPatternFlowLacpPartnerStateDefaultedCounter() PatternFlowLacpPartnerStateDefaultedCounter {
	obj := patternFlowLacpPartnerStateDefaultedCounter{obj: &otg.PatternFlowLacpPartnerStateDefaultedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerStateDefaultedCounter) msg() *otg.PatternFlowLacpPartnerStateDefaultedCounter {
	return obj.obj
}

func (obj *patternFlowLacpPartnerStateDefaultedCounter) setMsg(msg *otg.PatternFlowLacpPartnerStateDefaultedCounter) PatternFlowLacpPartnerStateDefaultedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerStateDefaultedCounter struct {
	obj *patternFlowLacpPartnerStateDefaultedCounter
}

type marshalPatternFlowLacpPartnerStateDefaultedCounter interface {
	// ToProto marshals PatternFlowLacpPartnerStateDefaultedCounter to protobuf object *otg.PatternFlowLacpPartnerStateDefaultedCounter
	ToProto() (*otg.PatternFlowLacpPartnerStateDefaultedCounter, error)
	// ToPbText marshals PatternFlowLacpPartnerStateDefaultedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerStateDefaultedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerStateDefaultedCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerStateDefaultedCounter struct {
	obj *patternFlowLacpPartnerStateDefaultedCounter
}

type unMarshalPatternFlowLacpPartnerStateDefaultedCounter interface {
	// FromProto unmarshals PatternFlowLacpPartnerStateDefaultedCounter from protobuf object *otg.PatternFlowLacpPartnerStateDefaultedCounter
	FromProto(msg *otg.PatternFlowLacpPartnerStateDefaultedCounter) (PatternFlowLacpPartnerStateDefaultedCounter, error)
	// FromPbText unmarshals PatternFlowLacpPartnerStateDefaultedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerStateDefaultedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerStateDefaultedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerStateDefaultedCounter) Marshal() marshalPatternFlowLacpPartnerStateDefaultedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerStateDefaultedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerStateDefaultedCounter) Unmarshal() unMarshalPatternFlowLacpPartnerStateDefaultedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerStateDefaultedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerStateDefaultedCounter) ToProto() (*otg.PatternFlowLacpPartnerStateDefaultedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerStateDefaultedCounter) FromProto(msg *otg.PatternFlowLacpPartnerStateDefaultedCounter) (PatternFlowLacpPartnerStateDefaultedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerStateDefaultedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateDefaultedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateDefaultedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateDefaultedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateDefaultedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateDefaultedCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerStateDefaultedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateDefaultedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateDefaultedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerStateDefaultedCounter) Clone() (PatternFlowLacpPartnerStateDefaultedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerStateDefaultedCounter()
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

// PatternFlowLacpPartnerStateDefaultedCounter is integer counter pattern
type PatternFlowLacpPartnerStateDefaultedCounter interface {
	Validation
	// msg marshals PatternFlowLacpPartnerStateDefaultedCounter to protobuf object *otg.PatternFlowLacpPartnerStateDefaultedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerStateDefaultedCounter
	// setMsg unmarshals PatternFlowLacpPartnerStateDefaultedCounter from protobuf object *otg.PatternFlowLacpPartnerStateDefaultedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerStateDefaultedCounter) PatternFlowLacpPartnerStateDefaultedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerStateDefaultedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerStateDefaultedCounter
	// validate validates PatternFlowLacpPartnerStateDefaultedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerStateDefaultedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpPartnerStateDefaultedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpPartnerStateDefaultedCounter
	SetStart(value uint32) PatternFlowLacpPartnerStateDefaultedCounter
	// HasStart checks if Start has been set in PatternFlowLacpPartnerStateDefaultedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpPartnerStateDefaultedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpPartnerStateDefaultedCounter
	SetStep(value uint32) PatternFlowLacpPartnerStateDefaultedCounter
	// HasStep checks if Step has been set in PatternFlowLacpPartnerStateDefaultedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpPartnerStateDefaultedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpPartnerStateDefaultedCounter
	SetCount(value uint32) PatternFlowLacpPartnerStateDefaultedCounter
	// HasCount checks if Count has been set in PatternFlowLacpPartnerStateDefaultedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerStateDefaultedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerStateDefaultedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpPartnerStateDefaultedCounter object
func (obj *patternFlowLacpPartnerStateDefaultedCounter) SetStart(value uint32) PatternFlowLacpPartnerStateDefaultedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerStateDefaultedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerStateDefaultedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpPartnerStateDefaultedCounter object
func (obj *patternFlowLacpPartnerStateDefaultedCounter) SetStep(value uint32) PatternFlowLacpPartnerStateDefaultedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerStateDefaultedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerStateDefaultedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpPartnerStateDefaultedCounter object
func (obj *patternFlowLacpPartnerStateDefaultedCounter) SetCount(value uint32) PatternFlowLacpPartnerStateDefaultedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpPartnerStateDefaultedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateDefaultedCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateDefaultedCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateDefaultedCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpPartnerStateDefaultedCounter) setDefault() {
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

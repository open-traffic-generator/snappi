package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerStateDistributingCounter *****
type patternFlowLacpPartnerStateDistributingCounter struct {
	validation
	obj          *otg.PatternFlowLacpPartnerStateDistributingCounter
	marshaller   marshalPatternFlowLacpPartnerStateDistributingCounter
	unMarshaller unMarshalPatternFlowLacpPartnerStateDistributingCounter
}

func NewPatternFlowLacpPartnerStateDistributingCounter() PatternFlowLacpPartnerStateDistributingCounter {
	obj := patternFlowLacpPartnerStateDistributingCounter{obj: &otg.PatternFlowLacpPartnerStateDistributingCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerStateDistributingCounter) msg() *otg.PatternFlowLacpPartnerStateDistributingCounter {
	return obj.obj
}

func (obj *patternFlowLacpPartnerStateDistributingCounter) setMsg(msg *otg.PatternFlowLacpPartnerStateDistributingCounter) PatternFlowLacpPartnerStateDistributingCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerStateDistributingCounter struct {
	obj *patternFlowLacpPartnerStateDistributingCounter
}

type marshalPatternFlowLacpPartnerStateDistributingCounter interface {
	// ToProto marshals PatternFlowLacpPartnerStateDistributingCounter to protobuf object *otg.PatternFlowLacpPartnerStateDistributingCounter
	ToProto() (*otg.PatternFlowLacpPartnerStateDistributingCounter, error)
	// ToPbText marshals PatternFlowLacpPartnerStateDistributingCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerStateDistributingCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerStateDistributingCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerStateDistributingCounter struct {
	obj *patternFlowLacpPartnerStateDistributingCounter
}

type unMarshalPatternFlowLacpPartnerStateDistributingCounter interface {
	// FromProto unmarshals PatternFlowLacpPartnerStateDistributingCounter from protobuf object *otg.PatternFlowLacpPartnerStateDistributingCounter
	FromProto(msg *otg.PatternFlowLacpPartnerStateDistributingCounter) (PatternFlowLacpPartnerStateDistributingCounter, error)
	// FromPbText unmarshals PatternFlowLacpPartnerStateDistributingCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerStateDistributingCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerStateDistributingCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerStateDistributingCounter) Marshal() marshalPatternFlowLacpPartnerStateDistributingCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerStateDistributingCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerStateDistributingCounter) Unmarshal() unMarshalPatternFlowLacpPartnerStateDistributingCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerStateDistributingCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerStateDistributingCounter) ToProto() (*otg.PatternFlowLacpPartnerStateDistributingCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerStateDistributingCounter) FromProto(msg *otg.PatternFlowLacpPartnerStateDistributingCounter) (PatternFlowLacpPartnerStateDistributingCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerStateDistributingCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateDistributingCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateDistributingCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateDistributingCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateDistributingCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateDistributingCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerStateDistributingCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateDistributingCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateDistributingCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerStateDistributingCounter) Clone() (PatternFlowLacpPartnerStateDistributingCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerStateDistributingCounter()
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

// PatternFlowLacpPartnerStateDistributingCounter is integer counter pattern
type PatternFlowLacpPartnerStateDistributingCounter interface {
	Validation
	// msg marshals PatternFlowLacpPartnerStateDistributingCounter to protobuf object *otg.PatternFlowLacpPartnerStateDistributingCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerStateDistributingCounter
	// setMsg unmarshals PatternFlowLacpPartnerStateDistributingCounter from protobuf object *otg.PatternFlowLacpPartnerStateDistributingCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerStateDistributingCounter) PatternFlowLacpPartnerStateDistributingCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerStateDistributingCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerStateDistributingCounter
	// validate validates PatternFlowLacpPartnerStateDistributingCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerStateDistributingCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpPartnerStateDistributingCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpPartnerStateDistributingCounter
	SetStart(value uint32) PatternFlowLacpPartnerStateDistributingCounter
	// HasStart checks if Start has been set in PatternFlowLacpPartnerStateDistributingCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpPartnerStateDistributingCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpPartnerStateDistributingCounter
	SetStep(value uint32) PatternFlowLacpPartnerStateDistributingCounter
	// HasStep checks if Step has been set in PatternFlowLacpPartnerStateDistributingCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpPartnerStateDistributingCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpPartnerStateDistributingCounter
	SetCount(value uint32) PatternFlowLacpPartnerStateDistributingCounter
	// HasCount checks if Count has been set in PatternFlowLacpPartnerStateDistributingCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerStateDistributingCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerStateDistributingCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpPartnerStateDistributingCounter object
func (obj *patternFlowLacpPartnerStateDistributingCounter) SetStart(value uint32) PatternFlowLacpPartnerStateDistributingCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerStateDistributingCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerStateDistributingCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpPartnerStateDistributingCounter object
func (obj *patternFlowLacpPartnerStateDistributingCounter) SetStep(value uint32) PatternFlowLacpPartnerStateDistributingCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerStateDistributingCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerStateDistributingCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpPartnerStateDistributingCounter object
func (obj *patternFlowLacpPartnerStateDistributingCounter) SetCount(value uint32) PatternFlowLacpPartnerStateDistributingCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpPartnerStateDistributingCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateDistributingCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateDistributingCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateDistributingCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpPartnerStateDistributingCounter) setDefault() {
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

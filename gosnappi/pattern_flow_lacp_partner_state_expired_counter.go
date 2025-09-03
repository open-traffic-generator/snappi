package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerStateExpiredCounter *****
type patternFlowLacpPartnerStateExpiredCounter struct {
	validation
	obj          *otg.PatternFlowLacpPartnerStateExpiredCounter
	marshaller   marshalPatternFlowLacpPartnerStateExpiredCounter
	unMarshaller unMarshalPatternFlowLacpPartnerStateExpiredCounter
}

func NewPatternFlowLacpPartnerStateExpiredCounter() PatternFlowLacpPartnerStateExpiredCounter {
	obj := patternFlowLacpPartnerStateExpiredCounter{obj: &otg.PatternFlowLacpPartnerStateExpiredCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerStateExpiredCounter) msg() *otg.PatternFlowLacpPartnerStateExpiredCounter {
	return obj.obj
}

func (obj *patternFlowLacpPartnerStateExpiredCounter) setMsg(msg *otg.PatternFlowLacpPartnerStateExpiredCounter) PatternFlowLacpPartnerStateExpiredCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerStateExpiredCounter struct {
	obj *patternFlowLacpPartnerStateExpiredCounter
}

type marshalPatternFlowLacpPartnerStateExpiredCounter interface {
	// ToProto marshals PatternFlowLacpPartnerStateExpiredCounter to protobuf object *otg.PatternFlowLacpPartnerStateExpiredCounter
	ToProto() (*otg.PatternFlowLacpPartnerStateExpiredCounter, error)
	// ToPbText marshals PatternFlowLacpPartnerStateExpiredCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerStateExpiredCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerStateExpiredCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerStateExpiredCounter struct {
	obj *patternFlowLacpPartnerStateExpiredCounter
}

type unMarshalPatternFlowLacpPartnerStateExpiredCounter interface {
	// FromProto unmarshals PatternFlowLacpPartnerStateExpiredCounter from protobuf object *otg.PatternFlowLacpPartnerStateExpiredCounter
	FromProto(msg *otg.PatternFlowLacpPartnerStateExpiredCounter) (PatternFlowLacpPartnerStateExpiredCounter, error)
	// FromPbText unmarshals PatternFlowLacpPartnerStateExpiredCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerStateExpiredCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerStateExpiredCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerStateExpiredCounter) Marshal() marshalPatternFlowLacpPartnerStateExpiredCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerStateExpiredCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerStateExpiredCounter) Unmarshal() unMarshalPatternFlowLacpPartnerStateExpiredCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerStateExpiredCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerStateExpiredCounter) ToProto() (*otg.PatternFlowLacpPartnerStateExpiredCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerStateExpiredCounter) FromProto(msg *otg.PatternFlowLacpPartnerStateExpiredCounter) (PatternFlowLacpPartnerStateExpiredCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerStateExpiredCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateExpiredCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateExpiredCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateExpiredCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateExpiredCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateExpiredCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerStateExpiredCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateExpiredCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateExpiredCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerStateExpiredCounter) Clone() (PatternFlowLacpPartnerStateExpiredCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerStateExpiredCounter()
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

// PatternFlowLacpPartnerStateExpiredCounter is integer counter pattern
type PatternFlowLacpPartnerStateExpiredCounter interface {
	Validation
	// msg marshals PatternFlowLacpPartnerStateExpiredCounter to protobuf object *otg.PatternFlowLacpPartnerStateExpiredCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerStateExpiredCounter
	// setMsg unmarshals PatternFlowLacpPartnerStateExpiredCounter from protobuf object *otg.PatternFlowLacpPartnerStateExpiredCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerStateExpiredCounter) PatternFlowLacpPartnerStateExpiredCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerStateExpiredCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerStateExpiredCounter
	// validate validates PatternFlowLacpPartnerStateExpiredCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerStateExpiredCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpPartnerStateExpiredCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpPartnerStateExpiredCounter
	SetStart(value uint32) PatternFlowLacpPartnerStateExpiredCounter
	// HasStart checks if Start has been set in PatternFlowLacpPartnerStateExpiredCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpPartnerStateExpiredCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpPartnerStateExpiredCounter
	SetStep(value uint32) PatternFlowLacpPartnerStateExpiredCounter
	// HasStep checks if Step has been set in PatternFlowLacpPartnerStateExpiredCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpPartnerStateExpiredCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpPartnerStateExpiredCounter
	SetCount(value uint32) PatternFlowLacpPartnerStateExpiredCounter
	// HasCount checks if Count has been set in PatternFlowLacpPartnerStateExpiredCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerStateExpiredCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerStateExpiredCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpPartnerStateExpiredCounter object
func (obj *patternFlowLacpPartnerStateExpiredCounter) SetStart(value uint32) PatternFlowLacpPartnerStateExpiredCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerStateExpiredCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerStateExpiredCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpPartnerStateExpiredCounter object
func (obj *patternFlowLacpPartnerStateExpiredCounter) SetStep(value uint32) PatternFlowLacpPartnerStateExpiredCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerStateExpiredCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerStateExpiredCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpPartnerStateExpiredCounter object
func (obj *patternFlowLacpPartnerStateExpiredCounter) SetCount(value uint32) PatternFlowLacpPartnerStateExpiredCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpPartnerStateExpiredCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateExpiredCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateExpiredCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateExpiredCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpPartnerStateExpiredCounter) setDefault() {
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

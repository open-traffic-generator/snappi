package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerStateTimeoutCounter *****
type patternFlowLacpduPartnerStateTimeoutCounter struct {
	validation
	obj          *otg.PatternFlowLacpduPartnerStateTimeoutCounter
	marshaller   marshalPatternFlowLacpduPartnerStateTimeoutCounter
	unMarshaller unMarshalPatternFlowLacpduPartnerStateTimeoutCounter
}

func NewPatternFlowLacpduPartnerStateTimeoutCounter() PatternFlowLacpduPartnerStateTimeoutCounter {
	obj := patternFlowLacpduPartnerStateTimeoutCounter{obj: &otg.PatternFlowLacpduPartnerStateTimeoutCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerStateTimeoutCounter) msg() *otg.PatternFlowLacpduPartnerStateTimeoutCounter {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerStateTimeoutCounter) setMsg(msg *otg.PatternFlowLacpduPartnerStateTimeoutCounter) PatternFlowLacpduPartnerStateTimeoutCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerStateTimeoutCounter struct {
	obj *patternFlowLacpduPartnerStateTimeoutCounter
}

type marshalPatternFlowLacpduPartnerStateTimeoutCounter interface {
	// ToProto marshals PatternFlowLacpduPartnerStateTimeoutCounter to protobuf object *otg.PatternFlowLacpduPartnerStateTimeoutCounter
	ToProto() (*otg.PatternFlowLacpduPartnerStateTimeoutCounter, error)
	// ToPbText marshals PatternFlowLacpduPartnerStateTimeoutCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerStateTimeoutCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerStateTimeoutCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerStateTimeoutCounter struct {
	obj *patternFlowLacpduPartnerStateTimeoutCounter
}

type unMarshalPatternFlowLacpduPartnerStateTimeoutCounter interface {
	// FromProto unmarshals PatternFlowLacpduPartnerStateTimeoutCounter from protobuf object *otg.PatternFlowLacpduPartnerStateTimeoutCounter
	FromProto(msg *otg.PatternFlowLacpduPartnerStateTimeoutCounter) (PatternFlowLacpduPartnerStateTimeoutCounter, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerStateTimeoutCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerStateTimeoutCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerStateTimeoutCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerStateTimeoutCounter) Marshal() marshalPatternFlowLacpduPartnerStateTimeoutCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerStateTimeoutCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerStateTimeoutCounter) Unmarshal() unMarshalPatternFlowLacpduPartnerStateTimeoutCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerStateTimeoutCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerStateTimeoutCounter) ToProto() (*otg.PatternFlowLacpduPartnerStateTimeoutCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerStateTimeoutCounter) FromProto(msg *otg.PatternFlowLacpduPartnerStateTimeoutCounter) (PatternFlowLacpduPartnerStateTimeoutCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerStateTimeoutCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateTimeoutCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateTimeoutCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateTimeoutCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateTimeoutCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateTimeoutCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerStateTimeoutCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateTimeoutCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateTimeoutCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerStateTimeoutCounter) Clone() (PatternFlowLacpduPartnerStateTimeoutCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerStateTimeoutCounter()
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

// PatternFlowLacpduPartnerStateTimeoutCounter is integer counter pattern
type PatternFlowLacpduPartnerStateTimeoutCounter interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerStateTimeoutCounter to protobuf object *otg.PatternFlowLacpduPartnerStateTimeoutCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerStateTimeoutCounter
	// setMsg unmarshals PatternFlowLacpduPartnerStateTimeoutCounter from protobuf object *otg.PatternFlowLacpduPartnerStateTimeoutCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerStateTimeoutCounter) PatternFlowLacpduPartnerStateTimeoutCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerStateTimeoutCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerStateTimeoutCounter
	// validate validates PatternFlowLacpduPartnerStateTimeoutCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerStateTimeoutCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduPartnerStateTimeoutCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduPartnerStateTimeoutCounter
	SetStart(value uint32) PatternFlowLacpduPartnerStateTimeoutCounter
	// HasStart checks if Start has been set in PatternFlowLacpduPartnerStateTimeoutCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduPartnerStateTimeoutCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduPartnerStateTimeoutCounter
	SetStep(value uint32) PatternFlowLacpduPartnerStateTimeoutCounter
	// HasStep checks if Step has been set in PatternFlowLacpduPartnerStateTimeoutCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduPartnerStateTimeoutCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduPartnerStateTimeoutCounter
	SetCount(value uint32) PatternFlowLacpduPartnerStateTimeoutCounter
	// HasCount checks if Count has been set in PatternFlowLacpduPartnerStateTimeoutCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerStateTimeoutCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerStateTimeoutCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduPartnerStateTimeoutCounter object
func (obj *patternFlowLacpduPartnerStateTimeoutCounter) SetStart(value uint32) PatternFlowLacpduPartnerStateTimeoutCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerStateTimeoutCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerStateTimeoutCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduPartnerStateTimeoutCounter object
func (obj *patternFlowLacpduPartnerStateTimeoutCounter) SetStep(value uint32) PatternFlowLacpduPartnerStateTimeoutCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerStateTimeoutCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerStateTimeoutCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduPartnerStateTimeoutCounter object
func (obj *patternFlowLacpduPartnerStateTimeoutCounter) SetCount(value uint32) PatternFlowLacpduPartnerStateTimeoutCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduPartnerStateTimeoutCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateTimeoutCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateTimeoutCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateTimeoutCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduPartnerStateTimeoutCounter) setDefault() {
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

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerStateActivityCounter *****
type patternFlowLacpduPartnerStateActivityCounter struct {
	validation
	obj          *otg.PatternFlowLacpduPartnerStateActivityCounter
	marshaller   marshalPatternFlowLacpduPartnerStateActivityCounter
	unMarshaller unMarshalPatternFlowLacpduPartnerStateActivityCounter
}

func NewPatternFlowLacpduPartnerStateActivityCounter() PatternFlowLacpduPartnerStateActivityCounter {
	obj := patternFlowLacpduPartnerStateActivityCounter{obj: &otg.PatternFlowLacpduPartnerStateActivityCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerStateActivityCounter) msg() *otg.PatternFlowLacpduPartnerStateActivityCounter {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerStateActivityCounter) setMsg(msg *otg.PatternFlowLacpduPartnerStateActivityCounter) PatternFlowLacpduPartnerStateActivityCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerStateActivityCounter struct {
	obj *patternFlowLacpduPartnerStateActivityCounter
}

type marshalPatternFlowLacpduPartnerStateActivityCounter interface {
	// ToProto marshals PatternFlowLacpduPartnerStateActivityCounter to protobuf object *otg.PatternFlowLacpduPartnerStateActivityCounter
	ToProto() (*otg.PatternFlowLacpduPartnerStateActivityCounter, error)
	// ToPbText marshals PatternFlowLacpduPartnerStateActivityCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerStateActivityCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerStateActivityCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerStateActivityCounter struct {
	obj *patternFlowLacpduPartnerStateActivityCounter
}

type unMarshalPatternFlowLacpduPartnerStateActivityCounter interface {
	// FromProto unmarshals PatternFlowLacpduPartnerStateActivityCounter from protobuf object *otg.PatternFlowLacpduPartnerStateActivityCounter
	FromProto(msg *otg.PatternFlowLacpduPartnerStateActivityCounter) (PatternFlowLacpduPartnerStateActivityCounter, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerStateActivityCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerStateActivityCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerStateActivityCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerStateActivityCounter) Marshal() marshalPatternFlowLacpduPartnerStateActivityCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerStateActivityCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerStateActivityCounter) Unmarshal() unMarshalPatternFlowLacpduPartnerStateActivityCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerStateActivityCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerStateActivityCounter) ToProto() (*otg.PatternFlowLacpduPartnerStateActivityCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerStateActivityCounter) FromProto(msg *otg.PatternFlowLacpduPartnerStateActivityCounter) (PatternFlowLacpduPartnerStateActivityCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerStateActivityCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateActivityCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateActivityCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateActivityCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateActivityCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateActivityCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerStateActivityCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateActivityCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateActivityCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerStateActivityCounter) Clone() (PatternFlowLacpduPartnerStateActivityCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerStateActivityCounter()
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

// PatternFlowLacpduPartnerStateActivityCounter is integer counter pattern
type PatternFlowLacpduPartnerStateActivityCounter interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerStateActivityCounter to protobuf object *otg.PatternFlowLacpduPartnerStateActivityCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerStateActivityCounter
	// setMsg unmarshals PatternFlowLacpduPartnerStateActivityCounter from protobuf object *otg.PatternFlowLacpduPartnerStateActivityCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerStateActivityCounter) PatternFlowLacpduPartnerStateActivityCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerStateActivityCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerStateActivityCounter
	// validate validates PatternFlowLacpduPartnerStateActivityCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerStateActivityCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduPartnerStateActivityCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduPartnerStateActivityCounter
	SetStart(value uint32) PatternFlowLacpduPartnerStateActivityCounter
	// HasStart checks if Start has been set in PatternFlowLacpduPartnerStateActivityCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduPartnerStateActivityCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduPartnerStateActivityCounter
	SetStep(value uint32) PatternFlowLacpduPartnerStateActivityCounter
	// HasStep checks if Step has been set in PatternFlowLacpduPartnerStateActivityCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduPartnerStateActivityCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduPartnerStateActivityCounter
	SetCount(value uint32) PatternFlowLacpduPartnerStateActivityCounter
	// HasCount checks if Count has been set in PatternFlowLacpduPartnerStateActivityCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerStateActivityCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerStateActivityCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduPartnerStateActivityCounter object
func (obj *patternFlowLacpduPartnerStateActivityCounter) SetStart(value uint32) PatternFlowLacpduPartnerStateActivityCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerStateActivityCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerStateActivityCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduPartnerStateActivityCounter object
func (obj *patternFlowLacpduPartnerStateActivityCounter) SetStep(value uint32) PatternFlowLacpduPartnerStateActivityCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerStateActivityCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerStateActivityCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduPartnerStateActivityCounter object
func (obj *patternFlowLacpduPartnerStateActivityCounter) SetCount(value uint32) PatternFlowLacpduPartnerStateActivityCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduPartnerStateActivityCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateActivityCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateActivityCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateActivityCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduPartnerStateActivityCounter) setDefault() {
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

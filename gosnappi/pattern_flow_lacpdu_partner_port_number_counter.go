package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerPortNumberCounter *****
type patternFlowLacpduPartnerPortNumberCounter struct {
	validation
	obj          *otg.PatternFlowLacpduPartnerPortNumberCounter
	marshaller   marshalPatternFlowLacpduPartnerPortNumberCounter
	unMarshaller unMarshalPatternFlowLacpduPartnerPortNumberCounter
}

func NewPatternFlowLacpduPartnerPortNumberCounter() PatternFlowLacpduPartnerPortNumberCounter {
	obj := patternFlowLacpduPartnerPortNumberCounter{obj: &otg.PatternFlowLacpduPartnerPortNumberCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerPortNumberCounter) msg() *otg.PatternFlowLacpduPartnerPortNumberCounter {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerPortNumberCounter) setMsg(msg *otg.PatternFlowLacpduPartnerPortNumberCounter) PatternFlowLacpduPartnerPortNumberCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerPortNumberCounter struct {
	obj *patternFlowLacpduPartnerPortNumberCounter
}

type marshalPatternFlowLacpduPartnerPortNumberCounter interface {
	// ToProto marshals PatternFlowLacpduPartnerPortNumberCounter to protobuf object *otg.PatternFlowLacpduPartnerPortNumberCounter
	ToProto() (*otg.PatternFlowLacpduPartnerPortNumberCounter, error)
	// ToPbText marshals PatternFlowLacpduPartnerPortNumberCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerPortNumberCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerPortNumberCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerPortNumberCounter struct {
	obj *patternFlowLacpduPartnerPortNumberCounter
}

type unMarshalPatternFlowLacpduPartnerPortNumberCounter interface {
	// FromProto unmarshals PatternFlowLacpduPartnerPortNumberCounter from protobuf object *otg.PatternFlowLacpduPartnerPortNumberCounter
	FromProto(msg *otg.PatternFlowLacpduPartnerPortNumberCounter) (PatternFlowLacpduPartnerPortNumberCounter, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerPortNumberCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerPortNumberCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerPortNumberCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerPortNumberCounter) Marshal() marshalPatternFlowLacpduPartnerPortNumberCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerPortNumberCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerPortNumberCounter) Unmarshal() unMarshalPatternFlowLacpduPartnerPortNumberCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerPortNumberCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerPortNumberCounter) ToProto() (*otg.PatternFlowLacpduPartnerPortNumberCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerPortNumberCounter) FromProto(msg *otg.PatternFlowLacpduPartnerPortNumberCounter) (PatternFlowLacpduPartnerPortNumberCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerPortNumberCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerPortNumberCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerPortNumberCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerPortNumberCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerPortNumberCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerPortNumberCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerPortNumberCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerPortNumberCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerPortNumberCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerPortNumberCounter) Clone() (PatternFlowLacpduPartnerPortNumberCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerPortNumberCounter()
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

// PatternFlowLacpduPartnerPortNumberCounter is integer counter pattern
type PatternFlowLacpduPartnerPortNumberCounter interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerPortNumberCounter to protobuf object *otg.PatternFlowLacpduPartnerPortNumberCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerPortNumberCounter
	// setMsg unmarshals PatternFlowLacpduPartnerPortNumberCounter from protobuf object *otg.PatternFlowLacpduPartnerPortNumberCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerPortNumberCounter) PatternFlowLacpduPartnerPortNumberCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerPortNumberCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerPortNumberCounter
	// validate validates PatternFlowLacpduPartnerPortNumberCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerPortNumberCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduPartnerPortNumberCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduPartnerPortNumberCounter
	SetStart(value uint32) PatternFlowLacpduPartnerPortNumberCounter
	// HasStart checks if Start has been set in PatternFlowLacpduPartnerPortNumberCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduPartnerPortNumberCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduPartnerPortNumberCounter
	SetStep(value uint32) PatternFlowLacpduPartnerPortNumberCounter
	// HasStep checks if Step has been set in PatternFlowLacpduPartnerPortNumberCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduPartnerPortNumberCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduPartnerPortNumberCounter
	SetCount(value uint32) PatternFlowLacpduPartnerPortNumberCounter
	// HasCount checks if Count has been set in PatternFlowLacpduPartnerPortNumberCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerPortNumberCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerPortNumberCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduPartnerPortNumberCounter object
func (obj *patternFlowLacpduPartnerPortNumberCounter) SetStart(value uint32) PatternFlowLacpduPartnerPortNumberCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerPortNumberCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerPortNumberCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduPartnerPortNumberCounter object
func (obj *patternFlowLacpduPartnerPortNumberCounter) SetStep(value uint32) PatternFlowLacpduPartnerPortNumberCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerPortNumberCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerPortNumberCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduPartnerPortNumberCounter object
func (obj *patternFlowLacpduPartnerPortNumberCounter) SetCount(value uint32) PatternFlowLacpduPartnerPortNumberCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduPartnerPortNumberCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerPortNumberCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerPortNumberCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerPortNumberCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduPartnerPortNumberCounter) setDefault() {
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

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVlanTpidCounter *****
type patternFlowVlanTpidCounter struct {
	validation
	obj          *otg.PatternFlowVlanTpidCounter
	marshaller   marshalPatternFlowVlanTpidCounter
	unMarshaller unMarshalPatternFlowVlanTpidCounter
}

func NewPatternFlowVlanTpidCounter() PatternFlowVlanTpidCounter {
	obj := patternFlowVlanTpidCounter{obj: &otg.PatternFlowVlanTpidCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVlanTpidCounter) msg() *otg.PatternFlowVlanTpidCounter {
	return obj.obj
}

func (obj *patternFlowVlanTpidCounter) setMsg(msg *otg.PatternFlowVlanTpidCounter) PatternFlowVlanTpidCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVlanTpidCounter struct {
	obj *patternFlowVlanTpidCounter
}

type marshalPatternFlowVlanTpidCounter interface {
	// ToProto marshals PatternFlowVlanTpidCounter to protobuf object *otg.PatternFlowVlanTpidCounter
	ToProto() (*otg.PatternFlowVlanTpidCounter, error)
	// ToPbText marshals PatternFlowVlanTpidCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVlanTpidCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVlanTpidCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowVlanTpidCounter struct {
	obj *patternFlowVlanTpidCounter
}

type unMarshalPatternFlowVlanTpidCounter interface {
	// FromProto unmarshals PatternFlowVlanTpidCounter from protobuf object *otg.PatternFlowVlanTpidCounter
	FromProto(msg *otg.PatternFlowVlanTpidCounter) (PatternFlowVlanTpidCounter, error)
	// FromPbText unmarshals PatternFlowVlanTpidCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVlanTpidCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVlanTpidCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVlanTpidCounter) Marshal() marshalPatternFlowVlanTpidCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVlanTpidCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVlanTpidCounter) Unmarshal() unMarshalPatternFlowVlanTpidCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVlanTpidCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVlanTpidCounter) ToProto() (*otg.PatternFlowVlanTpidCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVlanTpidCounter) FromProto(msg *otg.PatternFlowVlanTpidCounter) (PatternFlowVlanTpidCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVlanTpidCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVlanTpidCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowVlanTpidCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVlanTpidCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowVlanTpidCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVlanTpidCounter) FromJson(value string) error {
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

func (obj *patternFlowVlanTpidCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVlanTpidCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVlanTpidCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVlanTpidCounter) Clone() (PatternFlowVlanTpidCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVlanTpidCounter()
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

// PatternFlowVlanTpidCounter is integer counter pattern
type PatternFlowVlanTpidCounter interface {
	Validation
	// msg marshals PatternFlowVlanTpidCounter to protobuf object *otg.PatternFlowVlanTpidCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowVlanTpidCounter
	// setMsg unmarshals PatternFlowVlanTpidCounter from protobuf object *otg.PatternFlowVlanTpidCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVlanTpidCounter) PatternFlowVlanTpidCounter
	// provides marshal interface
	Marshal() marshalPatternFlowVlanTpidCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVlanTpidCounter
	// validate validates PatternFlowVlanTpidCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVlanTpidCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowVlanTpidCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowVlanTpidCounter
	SetStart(value uint32) PatternFlowVlanTpidCounter
	// HasStart checks if Start has been set in PatternFlowVlanTpidCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowVlanTpidCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowVlanTpidCounter
	SetStep(value uint32) PatternFlowVlanTpidCounter
	// HasStep checks if Step has been set in PatternFlowVlanTpidCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowVlanTpidCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowVlanTpidCounter
	SetCount(value uint32) PatternFlowVlanTpidCounter
	// HasCount checks if Count has been set in PatternFlowVlanTpidCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowVlanTpidCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowVlanTpidCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowVlanTpidCounter object
func (obj *patternFlowVlanTpidCounter) SetStart(value uint32) PatternFlowVlanTpidCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowVlanTpidCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowVlanTpidCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowVlanTpidCounter object
func (obj *patternFlowVlanTpidCounter) SetStep(value uint32) PatternFlowVlanTpidCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowVlanTpidCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowVlanTpidCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowVlanTpidCounter object
func (obj *patternFlowVlanTpidCounter) SetCount(value uint32) PatternFlowVlanTpidCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowVlanTpidCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVlanTpidCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVlanTpidCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVlanTpidCounter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowVlanTpidCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(65535)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

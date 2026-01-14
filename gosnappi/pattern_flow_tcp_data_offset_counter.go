package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpDataOffsetCounter *****
type patternFlowTcpDataOffsetCounter struct {
	validation
	obj          *otg.PatternFlowTcpDataOffsetCounter
	marshaller   marshalPatternFlowTcpDataOffsetCounter
	unMarshaller unMarshalPatternFlowTcpDataOffsetCounter
}

func NewPatternFlowTcpDataOffsetCounter() PatternFlowTcpDataOffsetCounter {
	obj := patternFlowTcpDataOffsetCounter{obj: &otg.PatternFlowTcpDataOffsetCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpDataOffsetCounter) msg() *otg.PatternFlowTcpDataOffsetCounter {
	return obj.obj
}

func (obj *patternFlowTcpDataOffsetCounter) setMsg(msg *otg.PatternFlowTcpDataOffsetCounter) PatternFlowTcpDataOffsetCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpDataOffsetCounter struct {
	obj *patternFlowTcpDataOffsetCounter
}

type marshalPatternFlowTcpDataOffsetCounter interface {
	// ToProto marshals PatternFlowTcpDataOffsetCounter to protobuf object *otg.PatternFlowTcpDataOffsetCounter
	ToProto() (*otg.PatternFlowTcpDataOffsetCounter, error)
	// ToPbText marshals PatternFlowTcpDataOffsetCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpDataOffsetCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpDataOffsetCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpDataOffsetCounter struct {
	obj *patternFlowTcpDataOffsetCounter
}

type unMarshalPatternFlowTcpDataOffsetCounter interface {
	// FromProto unmarshals PatternFlowTcpDataOffsetCounter from protobuf object *otg.PatternFlowTcpDataOffsetCounter
	FromProto(msg *otg.PatternFlowTcpDataOffsetCounter) (PatternFlowTcpDataOffsetCounter, error)
	// FromPbText unmarshals PatternFlowTcpDataOffsetCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpDataOffsetCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpDataOffsetCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpDataOffsetCounter) Marshal() marshalPatternFlowTcpDataOffsetCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpDataOffsetCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpDataOffsetCounter) Unmarshal() unMarshalPatternFlowTcpDataOffsetCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpDataOffsetCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpDataOffsetCounter) ToProto() (*otg.PatternFlowTcpDataOffsetCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpDataOffsetCounter) FromProto(msg *otg.PatternFlowTcpDataOffsetCounter) (PatternFlowTcpDataOffsetCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpDataOffsetCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpDataOffsetCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpDataOffsetCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpDataOffsetCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpDataOffsetCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpDataOffsetCounter) FromJson(value string) error {
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

func (obj *patternFlowTcpDataOffsetCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpDataOffsetCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpDataOffsetCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpDataOffsetCounter) Clone() (PatternFlowTcpDataOffsetCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpDataOffsetCounter()
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

// PatternFlowTcpDataOffsetCounter is integer counter pattern
type PatternFlowTcpDataOffsetCounter interface {
	Validation
	// msg marshals PatternFlowTcpDataOffsetCounter to protobuf object *otg.PatternFlowTcpDataOffsetCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpDataOffsetCounter
	// setMsg unmarshals PatternFlowTcpDataOffsetCounter from protobuf object *otg.PatternFlowTcpDataOffsetCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpDataOffsetCounter) PatternFlowTcpDataOffsetCounter
	// provides marshal interface
	Marshal() marshalPatternFlowTcpDataOffsetCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpDataOffsetCounter
	// validate validates PatternFlowTcpDataOffsetCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpDataOffsetCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowTcpDataOffsetCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowTcpDataOffsetCounter
	SetStart(value uint32) PatternFlowTcpDataOffsetCounter
	// HasStart checks if Start has been set in PatternFlowTcpDataOffsetCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowTcpDataOffsetCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowTcpDataOffsetCounter
	SetStep(value uint32) PatternFlowTcpDataOffsetCounter
	// HasStep checks if Step has been set in PatternFlowTcpDataOffsetCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowTcpDataOffsetCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowTcpDataOffsetCounter
	SetCount(value uint32) PatternFlowTcpDataOffsetCounter
	// HasCount checks if Count has been set in PatternFlowTcpDataOffsetCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpDataOffsetCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpDataOffsetCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowTcpDataOffsetCounter object
func (obj *patternFlowTcpDataOffsetCounter) SetStart(value uint32) PatternFlowTcpDataOffsetCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpDataOffsetCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpDataOffsetCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowTcpDataOffsetCounter object
func (obj *patternFlowTcpDataOffsetCounter) SetStep(value uint32) PatternFlowTcpDataOffsetCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpDataOffsetCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpDataOffsetCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowTcpDataOffsetCounter object
func (obj *patternFlowTcpDataOffsetCounter) SetCount(value uint32) PatternFlowTcpDataOffsetCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowTcpDataOffsetCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpDataOffsetCounter.Start <= 15 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpDataOffsetCounter.Step <= 15 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpDataOffsetCounter.Count <= 16 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowTcpDataOffsetCounter) setDefault() {
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

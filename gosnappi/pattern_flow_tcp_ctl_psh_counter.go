package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpCtlPshCounter *****
type patternFlowTcpCtlPshCounter struct {
	validation
	obj          *otg.PatternFlowTcpCtlPshCounter
	marshaller   marshalPatternFlowTcpCtlPshCounter
	unMarshaller unMarshalPatternFlowTcpCtlPshCounter
}

func NewPatternFlowTcpCtlPshCounter() PatternFlowTcpCtlPshCounter {
	obj := patternFlowTcpCtlPshCounter{obj: &otg.PatternFlowTcpCtlPshCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpCtlPshCounter) msg() *otg.PatternFlowTcpCtlPshCounter {
	return obj.obj
}

func (obj *patternFlowTcpCtlPshCounter) setMsg(msg *otg.PatternFlowTcpCtlPshCounter) PatternFlowTcpCtlPshCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpCtlPshCounter struct {
	obj *patternFlowTcpCtlPshCounter
}

type marshalPatternFlowTcpCtlPshCounter interface {
	// ToProto marshals PatternFlowTcpCtlPshCounter to protobuf object *otg.PatternFlowTcpCtlPshCounter
	ToProto() (*otg.PatternFlowTcpCtlPshCounter, error)
	// ToPbText marshals PatternFlowTcpCtlPshCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpCtlPshCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpCtlPshCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpCtlPshCounter struct {
	obj *patternFlowTcpCtlPshCounter
}

type unMarshalPatternFlowTcpCtlPshCounter interface {
	// FromProto unmarshals PatternFlowTcpCtlPshCounter from protobuf object *otg.PatternFlowTcpCtlPshCounter
	FromProto(msg *otg.PatternFlowTcpCtlPshCounter) (PatternFlowTcpCtlPshCounter, error)
	// FromPbText unmarshals PatternFlowTcpCtlPshCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpCtlPshCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpCtlPshCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpCtlPshCounter) Marshal() marshalPatternFlowTcpCtlPshCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpCtlPshCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpCtlPshCounter) Unmarshal() unMarshalPatternFlowTcpCtlPshCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpCtlPshCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpCtlPshCounter) ToProto() (*otg.PatternFlowTcpCtlPshCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpCtlPshCounter) FromProto(msg *otg.PatternFlowTcpCtlPshCounter) (PatternFlowTcpCtlPshCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpCtlPshCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlPshCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpCtlPshCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlPshCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpCtlPshCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlPshCounter) FromJson(value string) error {
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

func (obj *patternFlowTcpCtlPshCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlPshCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlPshCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpCtlPshCounter) Clone() (PatternFlowTcpCtlPshCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpCtlPshCounter()
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

// PatternFlowTcpCtlPshCounter is integer counter pattern
type PatternFlowTcpCtlPshCounter interface {
	Validation
	// msg marshals PatternFlowTcpCtlPshCounter to protobuf object *otg.PatternFlowTcpCtlPshCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpCtlPshCounter
	// setMsg unmarshals PatternFlowTcpCtlPshCounter from protobuf object *otg.PatternFlowTcpCtlPshCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpCtlPshCounter) PatternFlowTcpCtlPshCounter
	// provides marshal interface
	Marshal() marshalPatternFlowTcpCtlPshCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpCtlPshCounter
	// validate validates PatternFlowTcpCtlPshCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpCtlPshCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowTcpCtlPshCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowTcpCtlPshCounter
	SetStart(value uint32) PatternFlowTcpCtlPshCounter
	// HasStart checks if Start has been set in PatternFlowTcpCtlPshCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowTcpCtlPshCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowTcpCtlPshCounter
	SetStep(value uint32) PatternFlowTcpCtlPshCounter
	// HasStep checks if Step has been set in PatternFlowTcpCtlPshCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowTcpCtlPshCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowTcpCtlPshCounter
	SetCount(value uint32) PatternFlowTcpCtlPshCounter
	// HasCount checks if Count has been set in PatternFlowTcpCtlPshCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpCtlPshCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpCtlPshCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowTcpCtlPshCounter object
func (obj *patternFlowTcpCtlPshCounter) SetStart(value uint32) PatternFlowTcpCtlPshCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpCtlPshCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpCtlPshCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowTcpCtlPshCounter object
func (obj *patternFlowTcpCtlPshCounter) SetStep(value uint32) PatternFlowTcpCtlPshCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpCtlPshCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpCtlPshCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowTcpCtlPshCounter object
func (obj *patternFlowTcpCtlPshCounter) SetCount(value uint32) PatternFlowTcpCtlPshCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowTcpCtlPshCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlPshCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlPshCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlPshCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowTcpCtlPshCounter) setDefault() {
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

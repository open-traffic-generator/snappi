package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpCtlUrgCounter *****
type patternFlowTcpCtlUrgCounter struct {
	validation
	obj          *otg.PatternFlowTcpCtlUrgCounter
	marshaller   marshalPatternFlowTcpCtlUrgCounter
	unMarshaller unMarshalPatternFlowTcpCtlUrgCounter
}

func NewPatternFlowTcpCtlUrgCounter() PatternFlowTcpCtlUrgCounter {
	obj := patternFlowTcpCtlUrgCounter{obj: &otg.PatternFlowTcpCtlUrgCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpCtlUrgCounter) msg() *otg.PatternFlowTcpCtlUrgCounter {
	return obj.obj
}

func (obj *patternFlowTcpCtlUrgCounter) setMsg(msg *otg.PatternFlowTcpCtlUrgCounter) PatternFlowTcpCtlUrgCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpCtlUrgCounter struct {
	obj *patternFlowTcpCtlUrgCounter
}

type marshalPatternFlowTcpCtlUrgCounter interface {
	// ToProto marshals PatternFlowTcpCtlUrgCounter to protobuf object *otg.PatternFlowTcpCtlUrgCounter
	ToProto() (*otg.PatternFlowTcpCtlUrgCounter, error)
	// ToPbText marshals PatternFlowTcpCtlUrgCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpCtlUrgCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpCtlUrgCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpCtlUrgCounter struct {
	obj *patternFlowTcpCtlUrgCounter
}

type unMarshalPatternFlowTcpCtlUrgCounter interface {
	// FromProto unmarshals PatternFlowTcpCtlUrgCounter from protobuf object *otg.PatternFlowTcpCtlUrgCounter
	FromProto(msg *otg.PatternFlowTcpCtlUrgCounter) (PatternFlowTcpCtlUrgCounter, error)
	// FromPbText unmarshals PatternFlowTcpCtlUrgCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpCtlUrgCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpCtlUrgCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpCtlUrgCounter) Marshal() marshalPatternFlowTcpCtlUrgCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpCtlUrgCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpCtlUrgCounter) Unmarshal() unMarshalPatternFlowTcpCtlUrgCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpCtlUrgCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpCtlUrgCounter) ToProto() (*otg.PatternFlowTcpCtlUrgCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpCtlUrgCounter) FromProto(msg *otg.PatternFlowTcpCtlUrgCounter) (PatternFlowTcpCtlUrgCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpCtlUrgCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlUrgCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpCtlUrgCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlUrgCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpCtlUrgCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlUrgCounter) FromJson(value string) error {
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

func (obj *patternFlowTcpCtlUrgCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlUrgCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlUrgCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpCtlUrgCounter) Clone() (PatternFlowTcpCtlUrgCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpCtlUrgCounter()
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

// PatternFlowTcpCtlUrgCounter is integer counter pattern
type PatternFlowTcpCtlUrgCounter interface {
	Validation
	// msg marshals PatternFlowTcpCtlUrgCounter to protobuf object *otg.PatternFlowTcpCtlUrgCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpCtlUrgCounter
	// setMsg unmarshals PatternFlowTcpCtlUrgCounter from protobuf object *otg.PatternFlowTcpCtlUrgCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpCtlUrgCounter) PatternFlowTcpCtlUrgCounter
	// provides marshal interface
	Marshal() marshalPatternFlowTcpCtlUrgCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpCtlUrgCounter
	// validate validates PatternFlowTcpCtlUrgCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpCtlUrgCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowTcpCtlUrgCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowTcpCtlUrgCounter
	SetStart(value uint32) PatternFlowTcpCtlUrgCounter
	// HasStart checks if Start has been set in PatternFlowTcpCtlUrgCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowTcpCtlUrgCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowTcpCtlUrgCounter
	SetStep(value uint32) PatternFlowTcpCtlUrgCounter
	// HasStep checks if Step has been set in PatternFlowTcpCtlUrgCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowTcpCtlUrgCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowTcpCtlUrgCounter
	SetCount(value uint32) PatternFlowTcpCtlUrgCounter
	// HasCount checks if Count has been set in PatternFlowTcpCtlUrgCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpCtlUrgCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpCtlUrgCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowTcpCtlUrgCounter object
func (obj *patternFlowTcpCtlUrgCounter) SetStart(value uint32) PatternFlowTcpCtlUrgCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpCtlUrgCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpCtlUrgCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowTcpCtlUrgCounter object
func (obj *patternFlowTcpCtlUrgCounter) SetStep(value uint32) PatternFlowTcpCtlUrgCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpCtlUrgCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpCtlUrgCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowTcpCtlUrgCounter object
func (obj *patternFlowTcpCtlUrgCounter) SetCount(value uint32) PatternFlowTcpCtlUrgCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowTcpCtlUrgCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlUrgCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlUrgCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlUrgCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowTcpCtlUrgCounter) setDefault() {
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

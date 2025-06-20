package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpCtlRstCounter *****
type patternFlowTcpCtlRstCounter struct {
	validation
	obj          *otg.PatternFlowTcpCtlRstCounter
	marshaller   marshalPatternFlowTcpCtlRstCounter
	unMarshaller unMarshalPatternFlowTcpCtlRstCounter
}

func NewPatternFlowTcpCtlRstCounter() PatternFlowTcpCtlRstCounter {
	obj := patternFlowTcpCtlRstCounter{obj: &otg.PatternFlowTcpCtlRstCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpCtlRstCounter) msg() *otg.PatternFlowTcpCtlRstCounter {
	return obj.obj
}

func (obj *patternFlowTcpCtlRstCounter) setMsg(msg *otg.PatternFlowTcpCtlRstCounter) PatternFlowTcpCtlRstCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpCtlRstCounter struct {
	obj *patternFlowTcpCtlRstCounter
}

type marshalPatternFlowTcpCtlRstCounter interface {
	// ToProto marshals PatternFlowTcpCtlRstCounter to protobuf object *otg.PatternFlowTcpCtlRstCounter
	ToProto() (*otg.PatternFlowTcpCtlRstCounter, error)
	// ToPbText marshals PatternFlowTcpCtlRstCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpCtlRstCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpCtlRstCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpCtlRstCounter struct {
	obj *patternFlowTcpCtlRstCounter
}

type unMarshalPatternFlowTcpCtlRstCounter interface {
	// FromProto unmarshals PatternFlowTcpCtlRstCounter from protobuf object *otg.PatternFlowTcpCtlRstCounter
	FromProto(msg *otg.PatternFlowTcpCtlRstCounter) (PatternFlowTcpCtlRstCounter, error)
	// FromPbText unmarshals PatternFlowTcpCtlRstCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpCtlRstCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpCtlRstCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpCtlRstCounter) Marshal() marshalPatternFlowTcpCtlRstCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpCtlRstCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpCtlRstCounter) Unmarshal() unMarshalPatternFlowTcpCtlRstCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpCtlRstCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpCtlRstCounter) ToProto() (*otg.PatternFlowTcpCtlRstCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpCtlRstCounter) FromProto(msg *otg.PatternFlowTcpCtlRstCounter) (PatternFlowTcpCtlRstCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpCtlRstCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlRstCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpCtlRstCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlRstCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpCtlRstCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlRstCounter) FromJson(value string) error {
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

func (obj *patternFlowTcpCtlRstCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlRstCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlRstCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpCtlRstCounter) Clone() (PatternFlowTcpCtlRstCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpCtlRstCounter()
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

// PatternFlowTcpCtlRstCounter is integer counter pattern
type PatternFlowTcpCtlRstCounter interface {
	Validation
	// msg marshals PatternFlowTcpCtlRstCounter to protobuf object *otg.PatternFlowTcpCtlRstCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpCtlRstCounter
	// setMsg unmarshals PatternFlowTcpCtlRstCounter from protobuf object *otg.PatternFlowTcpCtlRstCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpCtlRstCounter) PatternFlowTcpCtlRstCounter
	// provides marshal interface
	Marshal() marshalPatternFlowTcpCtlRstCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpCtlRstCounter
	// validate validates PatternFlowTcpCtlRstCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpCtlRstCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowTcpCtlRstCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowTcpCtlRstCounter
	SetStart(value uint32) PatternFlowTcpCtlRstCounter
	// HasStart checks if Start has been set in PatternFlowTcpCtlRstCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowTcpCtlRstCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowTcpCtlRstCounter
	SetStep(value uint32) PatternFlowTcpCtlRstCounter
	// HasStep checks if Step has been set in PatternFlowTcpCtlRstCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowTcpCtlRstCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowTcpCtlRstCounter
	SetCount(value uint32) PatternFlowTcpCtlRstCounter
	// HasCount checks if Count has been set in PatternFlowTcpCtlRstCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpCtlRstCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpCtlRstCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowTcpCtlRstCounter object
func (obj *patternFlowTcpCtlRstCounter) SetStart(value uint32) PatternFlowTcpCtlRstCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpCtlRstCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpCtlRstCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowTcpCtlRstCounter object
func (obj *patternFlowTcpCtlRstCounter) SetStep(value uint32) PatternFlowTcpCtlRstCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpCtlRstCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpCtlRstCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowTcpCtlRstCounter object
func (obj *patternFlowTcpCtlRstCounter) SetCount(value uint32) PatternFlowTcpCtlRstCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowTcpCtlRstCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlRstCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlRstCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlRstCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowTcpCtlRstCounter) setDefault() {
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

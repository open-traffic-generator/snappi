package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpCtlAckCounter *****
type patternFlowTcpCtlAckCounter struct {
	validation
	obj          *otg.PatternFlowTcpCtlAckCounter
	marshaller   marshalPatternFlowTcpCtlAckCounter
	unMarshaller unMarshalPatternFlowTcpCtlAckCounter
}

func NewPatternFlowTcpCtlAckCounter() PatternFlowTcpCtlAckCounter {
	obj := patternFlowTcpCtlAckCounter{obj: &otg.PatternFlowTcpCtlAckCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpCtlAckCounter) msg() *otg.PatternFlowTcpCtlAckCounter {
	return obj.obj
}

func (obj *patternFlowTcpCtlAckCounter) setMsg(msg *otg.PatternFlowTcpCtlAckCounter) PatternFlowTcpCtlAckCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpCtlAckCounter struct {
	obj *patternFlowTcpCtlAckCounter
}

type marshalPatternFlowTcpCtlAckCounter interface {
	// ToProto marshals PatternFlowTcpCtlAckCounter to protobuf object *otg.PatternFlowTcpCtlAckCounter
	ToProto() (*otg.PatternFlowTcpCtlAckCounter, error)
	// ToPbText marshals PatternFlowTcpCtlAckCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpCtlAckCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpCtlAckCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpCtlAckCounter struct {
	obj *patternFlowTcpCtlAckCounter
}

type unMarshalPatternFlowTcpCtlAckCounter interface {
	// FromProto unmarshals PatternFlowTcpCtlAckCounter from protobuf object *otg.PatternFlowTcpCtlAckCounter
	FromProto(msg *otg.PatternFlowTcpCtlAckCounter) (PatternFlowTcpCtlAckCounter, error)
	// FromPbText unmarshals PatternFlowTcpCtlAckCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpCtlAckCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpCtlAckCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpCtlAckCounter) Marshal() marshalPatternFlowTcpCtlAckCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpCtlAckCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpCtlAckCounter) Unmarshal() unMarshalPatternFlowTcpCtlAckCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpCtlAckCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpCtlAckCounter) ToProto() (*otg.PatternFlowTcpCtlAckCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpCtlAckCounter) FromProto(msg *otg.PatternFlowTcpCtlAckCounter) (PatternFlowTcpCtlAckCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpCtlAckCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlAckCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpCtlAckCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlAckCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpCtlAckCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlAckCounter) FromJson(value string) error {
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

func (obj *patternFlowTcpCtlAckCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlAckCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlAckCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpCtlAckCounter) Clone() (PatternFlowTcpCtlAckCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpCtlAckCounter()
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

// PatternFlowTcpCtlAckCounter is integer counter pattern
type PatternFlowTcpCtlAckCounter interface {
	Validation
	// msg marshals PatternFlowTcpCtlAckCounter to protobuf object *otg.PatternFlowTcpCtlAckCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpCtlAckCounter
	// setMsg unmarshals PatternFlowTcpCtlAckCounter from protobuf object *otg.PatternFlowTcpCtlAckCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpCtlAckCounter) PatternFlowTcpCtlAckCounter
	// provides marshal interface
	Marshal() marshalPatternFlowTcpCtlAckCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpCtlAckCounter
	// validate validates PatternFlowTcpCtlAckCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpCtlAckCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowTcpCtlAckCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowTcpCtlAckCounter
	SetStart(value uint32) PatternFlowTcpCtlAckCounter
	// HasStart checks if Start has been set in PatternFlowTcpCtlAckCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowTcpCtlAckCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowTcpCtlAckCounter
	SetStep(value uint32) PatternFlowTcpCtlAckCounter
	// HasStep checks if Step has been set in PatternFlowTcpCtlAckCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowTcpCtlAckCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowTcpCtlAckCounter
	SetCount(value uint32) PatternFlowTcpCtlAckCounter
	// HasCount checks if Count has been set in PatternFlowTcpCtlAckCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpCtlAckCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpCtlAckCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowTcpCtlAckCounter object
func (obj *patternFlowTcpCtlAckCounter) SetStart(value uint32) PatternFlowTcpCtlAckCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpCtlAckCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpCtlAckCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowTcpCtlAckCounter object
func (obj *patternFlowTcpCtlAckCounter) SetStep(value uint32) PatternFlowTcpCtlAckCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpCtlAckCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpCtlAckCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowTcpCtlAckCounter object
func (obj *patternFlowTcpCtlAckCounter) SetCount(value uint32) PatternFlowTcpCtlAckCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowTcpCtlAckCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlAckCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlAckCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlAckCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowTcpCtlAckCounter) setDefault() {
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

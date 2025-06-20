package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpDstPortCounter *****
type patternFlowTcpDstPortCounter struct {
	validation
	obj          *otg.PatternFlowTcpDstPortCounter
	marshaller   marshalPatternFlowTcpDstPortCounter
	unMarshaller unMarshalPatternFlowTcpDstPortCounter
}

func NewPatternFlowTcpDstPortCounter() PatternFlowTcpDstPortCounter {
	obj := patternFlowTcpDstPortCounter{obj: &otg.PatternFlowTcpDstPortCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpDstPortCounter) msg() *otg.PatternFlowTcpDstPortCounter {
	return obj.obj
}

func (obj *patternFlowTcpDstPortCounter) setMsg(msg *otg.PatternFlowTcpDstPortCounter) PatternFlowTcpDstPortCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpDstPortCounter struct {
	obj *patternFlowTcpDstPortCounter
}

type marshalPatternFlowTcpDstPortCounter interface {
	// ToProto marshals PatternFlowTcpDstPortCounter to protobuf object *otg.PatternFlowTcpDstPortCounter
	ToProto() (*otg.PatternFlowTcpDstPortCounter, error)
	// ToPbText marshals PatternFlowTcpDstPortCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpDstPortCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpDstPortCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpDstPortCounter struct {
	obj *patternFlowTcpDstPortCounter
}

type unMarshalPatternFlowTcpDstPortCounter interface {
	// FromProto unmarshals PatternFlowTcpDstPortCounter from protobuf object *otg.PatternFlowTcpDstPortCounter
	FromProto(msg *otg.PatternFlowTcpDstPortCounter) (PatternFlowTcpDstPortCounter, error)
	// FromPbText unmarshals PatternFlowTcpDstPortCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpDstPortCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpDstPortCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpDstPortCounter) Marshal() marshalPatternFlowTcpDstPortCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpDstPortCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpDstPortCounter) Unmarshal() unMarshalPatternFlowTcpDstPortCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpDstPortCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpDstPortCounter) ToProto() (*otg.PatternFlowTcpDstPortCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpDstPortCounter) FromProto(msg *otg.PatternFlowTcpDstPortCounter) (PatternFlowTcpDstPortCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpDstPortCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpDstPortCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpDstPortCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpDstPortCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpDstPortCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpDstPortCounter) FromJson(value string) error {
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

func (obj *patternFlowTcpDstPortCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpDstPortCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpDstPortCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpDstPortCounter) Clone() (PatternFlowTcpDstPortCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpDstPortCounter()
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

// PatternFlowTcpDstPortCounter is integer counter pattern
type PatternFlowTcpDstPortCounter interface {
	Validation
	// msg marshals PatternFlowTcpDstPortCounter to protobuf object *otg.PatternFlowTcpDstPortCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpDstPortCounter
	// setMsg unmarshals PatternFlowTcpDstPortCounter from protobuf object *otg.PatternFlowTcpDstPortCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpDstPortCounter) PatternFlowTcpDstPortCounter
	// provides marshal interface
	Marshal() marshalPatternFlowTcpDstPortCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpDstPortCounter
	// validate validates PatternFlowTcpDstPortCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpDstPortCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowTcpDstPortCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowTcpDstPortCounter
	SetStart(value uint32) PatternFlowTcpDstPortCounter
	// HasStart checks if Start has been set in PatternFlowTcpDstPortCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowTcpDstPortCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowTcpDstPortCounter
	SetStep(value uint32) PatternFlowTcpDstPortCounter
	// HasStep checks if Step has been set in PatternFlowTcpDstPortCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowTcpDstPortCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowTcpDstPortCounter
	SetCount(value uint32) PatternFlowTcpDstPortCounter
	// HasCount checks if Count has been set in PatternFlowTcpDstPortCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpDstPortCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpDstPortCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowTcpDstPortCounter object
func (obj *patternFlowTcpDstPortCounter) SetStart(value uint32) PatternFlowTcpDstPortCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpDstPortCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpDstPortCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowTcpDstPortCounter object
func (obj *patternFlowTcpDstPortCounter) SetStep(value uint32) PatternFlowTcpDstPortCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpDstPortCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpDstPortCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowTcpDstPortCounter object
func (obj *patternFlowTcpDstPortCounter) SetCount(value uint32) PatternFlowTcpDstPortCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowTcpDstPortCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpDstPortCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpDstPortCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpDstPortCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowTcpDstPortCounter) setDefault() {
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

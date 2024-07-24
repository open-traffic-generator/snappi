package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpWindowCounter *****
type patternFlowTcpWindowCounter struct {
	validation
	obj          *otg.PatternFlowTcpWindowCounter
	marshaller   marshalPatternFlowTcpWindowCounter
	unMarshaller unMarshalPatternFlowTcpWindowCounter
}

func NewPatternFlowTcpWindowCounter() PatternFlowTcpWindowCounter {
	obj := patternFlowTcpWindowCounter{obj: &otg.PatternFlowTcpWindowCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpWindowCounter) msg() *otg.PatternFlowTcpWindowCounter {
	return obj.obj
}

func (obj *patternFlowTcpWindowCounter) setMsg(msg *otg.PatternFlowTcpWindowCounter) PatternFlowTcpWindowCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpWindowCounter struct {
	obj *patternFlowTcpWindowCounter
}

type marshalPatternFlowTcpWindowCounter interface {
	// ToProto marshals PatternFlowTcpWindowCounter to protobuf object *otg.PatternFlowTcpWindowCounter
	ToProto() (*otg.PatternFlowTcpWindowCounter, error)
	// ToPbText marshals PatternFlowTcpWindowCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpWindowCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpWindowCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpWindowCounter struct {
	obj *patternFlowTcpWindowCounter
}

type unMarshalPatternFlowTcpWindowCounter interface {
	// FromProto unmarshals PatternFlowTcpWindowCounter from protobuf object *otg.PatternFlowTcpWindowCounter
	FromProto(msg *otg.PatternFlowTcpWindowCounter) (PatternFlowTcpWindowCounter, error)
	// FromPbText unmarshals PatternFlowTcpWindowCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpWindowCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpWindowCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpWindowCounter) Marshal() marshalPatternFlowTcpWindowCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpWindowCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpWindowCounter) Unmarshal() unMarshalPatternFlowTcpWindowCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpWindowCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpWindowCounter) ToProto() (*otg.PatternFlowTcpWindowCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpWindowCounter) FromProto(msg *otg.PatternFlowTcpWindowCounter) (PatternFlowTcpWindowCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpWindowCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpWindowCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpWindowCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpWindowCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpWindowCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpWindowCounter) FromJson(value string) error {
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

func (obj *patternFlowTcpWindowCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpWindowCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpWindowCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpWindowCounter) Clone() (PatternFlowTcpWindowCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpWindowCounter()
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

// PatternFlowTcpWindowCounter is integer counter pattern
type PatternFlowTcpWindowCounter interface {
	Validation
	// msg marshals PatternFlowTcpWindowCounter to protobuf object *otg.PatternFlowTcpWindowCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpWindowCounter
	// setMsg unmarshals PatternFlowTcpWindowCounter from protobuf object *otg.PatternFlowTcpWindowCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpWindowCounter) PatternFlowTcpWindowCounter
	// provides marshal interface
	Marshal() marshalPatternFlowTcpWindowCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpWindowCounter
	// validate validates PatternFlowTcpWindowCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpWindowCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowTcpWindowCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowTcpWindowCounter
	SetStart(value uint32) PatternFlowTcpWindowCounter
	// HasStart checks if Start has been set in PatternFlowTcpWindowCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowTcpWindowCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowTcpWindowCounter
	SetStep(value uint32) PatternFlowTcpWindowCounter
	// HasStep checks if Step has been set in PatternFlowTcpWindowCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowTcpWindowCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowTcpWindowCounter
	SetCount(value uint32) PatternFlowTcpWindowCounter
	// HasCount checks if Count has been set in PatternFlowTcpWindowCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpWindowCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpWindowCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowTcpWindowCounter object
func (obj *patternFlowTcpWindowCounter) SetStart(value uint32) PatternFlowTcpWindowCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpWindowCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpWindowCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowTcpWindowCounter object
func (obj *patternFlowTcpWindowCounter) SetStep(value uint32) PatternFlowTcpWindowCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpWindowCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpWindowCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowTcpWindowCounter object
func (obj *patternFlowTcpWindowCounter) SetCount(value uint32) PatternFlowTcpWindowCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowTcpWindowCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpWindowCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpWindowCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpWindowCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowTcpWindowCounter) setDefault() {
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

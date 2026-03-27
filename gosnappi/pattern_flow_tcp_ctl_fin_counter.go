package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpCtlFinCounter *****
type patternFlowTcpCtlFinCounter struct {
	validation
	obj          *otg.PatternFlowTcpCtlFinCounter
	marshaller   marshalPatternFlowTcpCtlFinCounter
	unMarshaller unMarshalPatternFlowTcpCtlFinCounter
}

func NewPatternFlowTcpCtlFinCounter() PatternFlowTcpCtlFinCounter {
	obj := patternFlowTcpCtlFinCounter{obj: &otg.PatternFlowTcpCtlFinCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpCtlFinCounter) msg() *otg.PatternFlowTcpCtlFinCounter {
	return obj.obj
}

func (obj *patternFlowTcpCtlFinCounter) setMsg(msg *otg.PatternFlowTcpCtlFinCounter) PatternFlowTcpCtlFinCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpCtlFinCounter struct {
	obj *patternFlowTcpCtlFinCounter
}

type marshalPatternFlowTcpCtlFinCounter interface {
	// ToProto marshals PatternFlowTcpCtlFinCounter to protobuf object *otg.PatternFlowTcpCtlFinCounter
	ToProto() (*otg.PatternFlowTcpCtlFinCounter, error)
	// ToPbText marshals PatternFlowTcpCtlFinCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpCtlFinCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpCtlFinCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpCtlFinCounter struct {
	obj *patternFlowTcpCtlFinCounter
}

type unMarshalPatternFlowTcpCtlFinCounter interface {
	// FromProto unmarshals PatternFlowTcpCtlFinCounter from protobuf object *otg.PatternFlowTcpCtlFinCounter
	FromProto(msg *otg.PatternFlowTcpCtlFinCounter) (PatternFlowTcpCtlFinCounter, error)
	// FromPbText unmarshals PatternFlowTcpCtlFinCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpCtlFinCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpCtlFinCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpCtlFinCounter) Marshal() marshalPatternFlowTcpCtlFinCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpCtlFinCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpCtlFinCounter) Unmarshal() unMarshalPatternFlowTcpCtlFinCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpCtlFinCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpCtlFinCounter) ToProto() (*otg.PatternFlowTcpCtlFinCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpCtlFinCounter) FromProto(msg *otg.PatternFlowTcpCtlFinCounter) (PatternFlowTcpCtlFinCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpCtlFinCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlFinCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpCtlFinCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlFinCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpCtlFinCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlFinCounter) FromJson(value string) error {
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

func (obj *patternFlowTcpCtlFinCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlFinCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlFinCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpCtlFinCounter) Clone() (PatternFlowTcpCtlFinCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpCtlFinCounter()
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

// PatternFlowTcpCtlFinCounter is integer counter pattern
type PatternFlowTcpCtlFinCounter interface {
	Validation
	// msg marshals PatternFlowTcpCtlFinCounter to protobuf object *otg.PatternFlowTcpCtlFinCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpCtlFinCounter
	// setMsg unmarshals PatternFlowTcpCtlFinCounter from protobuf object *otg.PatternFlowTcpCtlFinCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpCtlFinCounter) PatternFlowTcpCtlFinCounter
	// provides marshal interface
	Marshal() marshalPatternFlowTcpCtlFinCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpCtlFinCounter
	// validate validates PatternFlowTcpCtlFinCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpCtlFinCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowTcpCtlFinCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowTcpCtlFinCounter
	SetStart(value uint32) PatternFlowTcpCtlFinCounter
	// HasStart checks if Start has been set in PatternFlowTcpCtlFinCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowTcpCtlFinCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowTcpCtlFinCounter
	SetStep(value uint32) PatternFlowTcpCtlFinCounter
	// HasStep checks if Step has been set in PatternFlowTcpCtlFinCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowTcpCtlFinCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowTcpCtlFinCounter
	SetCount(value uint32) PatternFlowTcpCtlFinCounter
	// HasCount checks if Count has been set in PatternFlowTcpCtlFinCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpCtlFinCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpCtlFinCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowTcpCtlFinCounter object
func (obj *patternFlowTcpCtlFinCounter) SetStart(value uint32) PatternFlowTcpCtlFinCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpCtlFinCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpCtlFinCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowTcpCtlFinCounter object
func (obj *patternFlowTcpCtlFinCounter) SetStep(value uint32) PatternFlowTcpCtlFinCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpCtlFinCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpCtlFinCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowTcpCtlFinCounter object
func (obj *patternFlowTcpCtlFinCounter) SetCount(value uint32) PatternFlowTcpCtlFinCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowTcpCtlFinCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlFinCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlFinCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlFinCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowTcpCtlFinCounter) setDefault() {
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

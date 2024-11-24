package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpSrcPortCounter *****
type patternFlowTcpSrcPortCounter struct {
	validation
	obj          *otg.PatternFlowTcpSrcPortCounter
	marshaller   marshalPatternFlowTcpSrcPortCounter
	unMarshaller unMarshalPatternFlowTcpSrcPortCounter
}

func NewPatternFlowTcpSrcPortCounter() PatternFlowTcpSrcPortCounter {
	obj := patternFlowTcpSrcPortCounter{obj: &otg.PatternFlowTcpSrcPortCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpSrcPortCounter) msg() *otg.PatternFlowTcpSrcPortCounter {
	return obj.obj
}

func (obj *patternFlowTcpSrcPortCounter) setMsg(msg *otg.PatternFlowTcpSrcPortCounter) PatternFlowTcpSrcPortCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpSrcPortCounter struct {
	obj *patternFlowTcpSrcPortCounter
}

type marshalPatternFlowTcpSrcPortCounter interface {
	// ToProto marshals PatternFlowTcpSrcPortCounter to protobuf object *otg.PatternFlowTcpSrcPortCounter
	ToProto() (*otg.PatternFlowTcpSrcPortCounter, error)
	// ToPbText marshals PatternFlowTcpSrcPortCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpSrcPortCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpSrcPortCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowTcpSrcPortCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowTcpSrcPortCounter struct {
	obj *patternFlowTcpSrcPortCounter
}

type unMarshalPatternFlowTcpSrcPortCounter interface {
	// FromProto unmarshals PatternFlowTcpSrcPortCounter from protobuf object *otg.PatternFlowTcpSrcPortCounter
	FromProto(msg *otg.PatternFlowTcpSrcPortCounter) (PatternFlowTcpSrcPortCounter, error)
	// FromPbText unmarshals PatternFlowTcpSrcPortCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpSrcPortCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpSrcPortCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpSrcPortCounter) Marshal() marshalPatternFlowTcpSrcPortCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpSrcPortCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpSrcPortCounter) Unmarshal() unMarshalPatternFlowTcpSrcPortCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpSrcPortCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpSrcPortCounter) ToProto() (*otg.PatternFlowTcpSrcPortCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpSrcPortCounter) FromProto(msg *otg.PatternFlowTcpSrcPortCounter) (PatternFlowTcpSrcPortCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpSrcPortCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpSrcPortCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpSrcPortCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpSrcPortCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpSrcPortCounter) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalpatternFlowTcpSrcPortCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpSrcPortCounter) FromJson(value string) error {
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

func (obj *patternFlowTcpSrcPortCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpSrcPortCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpSrcPortCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpSrcPortCounter) Clone() (PatternFlowTcpSrcPortCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpSrcPortCounter()
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

// PatternFlowTcpSrcPortCounter is integer counter pattern
type PatternFlowTcpSrcPortCounter interface {
	Validation
	// msg marshals PatternFlowTcpSrcPortCounter to protobuf object *otg.PatternFlowTcpSrcPortCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpSrcPortCounter
	// setMsg unmarshals PatternFlowTcpSrcPortCounter from protobuf object *otg.PatternFlowTcpSrcPortCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpSrcPortCounter) PatternFlowTcpSrcPortCounter
	// provides marshal interface
	Marshal() marshalPatternFlowTcpSrcPortCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpSrcPortCounter
	// validate validates PatternFlowTcpSrcPortCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpSrcPortCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowTcpSrcPortCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowTcpSrcPortCounter
	SetStart(value uint32) PatternFlowTcpSrcPortCounter
	// HasStart checks if Start has been set in PatternFlowTcpSrcPortCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowTcpSrcPortCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowTcpSrcPortCounter
	SetStep(value uint32) PatternFlowTcpSrcPortCounter
	// HasStep checks if Step has been set in PatternFlowTcpSrcPortCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowTcpSrcPortCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowTcpSrcPortCounter
	SetCount(value uint32) PatternFlowTcpSrcPortCounter
	// HasCount checks if Count has been set in PatternFlowTcpSrcPortCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpSrcPortCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpSrcPortCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowTcpSrcPortCounter object
func (obj *patternFlowTcpSrcPortCounter) SetStart(value uint32) PatternFlowTcpSrcPortCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpSrcPortCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpSrcPortCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowTcpSrcPortCounter object
func (obj *patternFlowTcpSrcPortCounter) SetStep(value uint32) PatternFlowTcpSrcPortCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpSrcPortCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpSrcPortCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowTcpSrcPortCounter object
func (obj *patternFlowTcpSrcPortCounter) SetCount(value uint32) PatternFlowTcpSrcPortCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowTcpSrcPortCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpSrcPortCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpSrcPortCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpSrcPortCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowTcpSrcPortCounter) setDefault() {
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

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowUdpDstPortCounter *****
type patternFlowUdpDstPortCounter struct {
	validation
	obj          *otg.PatternFlowUdpDstPortCounter
	marshaller   marshalPatternFlowUdpDstPortCounter
	unMarshaller unMarshalPatternFlowUdpDstPortCounter
}

func NewPatternFlowUdpDstPortCounter() PatternFlowUdpDstPortCounter {
	obj := patternFlowUdpDstPortCounter{obj: &otg.PatternFlowUdpDstPortCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowUdpDstPortCounter) msg() *otg.PatternFlowUdpDstPortCounter {
	return obj.obj
}

func (obj *patternFlowUdpDstPortCounter) setMsg(msg *otg.PatternFlowUdpDstPortCounter) PatternFlowUdpDstPortCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowUdpDstPortCounter struct {
	obj *patternFlowUdpDstPortCounter
}

type marshalPatternFlowUdpDstPortCounter interface {
	// ToProto marshals PatternFlowUdpDstPortCounter to protobuf object *otg.PatternFlowUdpDstPortCounter
	ToProto() (*otg.PatternFlowUdpDstPortCounter, error)
	// ToPbText marshals PatternFlowUdpDstPortCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowUdpDstPortCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowUdpDstPortCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowUdpDstPortCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowUdpDstPortCounter struct {
	obj *patternFlowUdpDstPortCounter
}

type unMarshalPatternFlowUdpDstPortCounter interface {
	// FromProto unmarshals PatternFlowUdpDstPortCounter from protobuf object *otg.PatternFlowUdpDstPortCounter
	FromProto(msg *otg.PatternFlowUdpDstPortCounter) (PatternFlowUdpDstPortCounter, error)
	// FromPbText unmarshals PatternFlowUdpDstPortCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowUdpDstPortCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowUdpDstPortCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowUdpDstPortCounter) Marshal() marshalPatternFlowUdpDstPortCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowUdpDstPortCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowUdpDstPortCounter) Unmarshal() unMarshalPatternFlowUdpDstPortCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowUdpDstPortCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowUdpDstPortCounter) ToProto() (*otg.PatternFlowUdpDstPortCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowUdpDstPortCounter) FromProto(msg *otg.PatternFlowUdpDstPortCounter) (PatternFlowUdpDstPortCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowUdpDstPortCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowUdpDstPortCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowUdpDstPortCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowUdpDstPortCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowUdpDstPortCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowUdpDstPortCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowUdpDstPortCounter) FromJson(value string) error {
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

func (obj *patternFlowUdpDstPortCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowUdpDstPortCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowUdpDstPortCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowUdpDstPortCounter) Clone() (PatternFlowUdpDstPortCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowUdpDstPortCounter()
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

// PatternFlowUdpDstPortCounter is integer counter pattern
type PatternFlowUdpDstPortCounter interface {
	Validation
	// msg marshals PatternFlowUdpDstPortCounter to protobuf object *otg.PatternFlowUdpDstPortCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowUdpDstPortCounter
	// setMsg unmarshals PatternFlowUdpDstPortCounter from protobuf object *otg.PatternFlowUdpDstPortCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowUdpDstPortCounter) PatternFlowUdpDstPortCounter
	// provides marshal interface
	Marshal() marshalPatternFlowUdpDstPortCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowUdpDstPortCounter
	// validate validates PatternFlowUdpDstPortCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowUdpDstPortCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowUdpDstPortCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowUdpDstPortCounter
	SetStart(value uint32) PatternFlowUdpDstPortCounter
	// HasStart checks if Start has been set in PatternFlowUdpDstPortCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowUdpDstPortCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowUdpDstPortCounter
	SetStep(value uint32) PatternFlowUdpDstPortCounter
	// HasStep checks if Step has been set in PatternFlowUdpDstPortCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowUdpDstPortCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowUdpDstPortCounter
	SetCount(value uint32) PatternFlowUdpDstPortCounter
	// HasCount checks if Count has been set in PatternFlowUdpDstPortCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowUdpDstPortCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowUdpDstPortCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowUdpDstPortCounter object
func (obj *patternFlowUdpDstPortCounter) SetStart(value uint32) PatternFlowUdpDstPortCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowUdpDstPortCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowUdpDstPortCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowUdpDstPortCounter object
func (obj *patternFlowUdpDstPortCounter) SetStep(value uint32) PatternFlowUdpDstPortCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowUdpDstPortCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowUdpDstPortCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowUdpDstPortCounter object
func (obj *patternFlowUdpDstPortCounter) SetCount(value uint32) PatternFlowUdpDstPortCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowUdpDstPortCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowUdpDstPortCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowUdpDstPortCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowUdpDstPortCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowUdpDstPortCounter) setDefault() {
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

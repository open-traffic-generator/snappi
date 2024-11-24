package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpProtocolLengthCounter *****
type patternFlowArpProtocolLengthCounter struct {
	validation
	obj          *otg.PatternFlowArpProtocolLengthCounter
	marshaller   marshalPatternFlowArpProtocolLengthCounter
	unMarshaller unMarshalPatternFlowArpProtocolLengthCounter
}

func NewPatternFlowArpProtocolLengthCounter() PatternFlowArpProtocolLengthCounter {
	obj := patternFlowArpProtocolLengthCounter{obj: &otg.PatternFlowArpProtocolLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpProtocolLengthCounter) msg() *otg.PatternFlowArpProtocolLengthCounter {
	return obj.obj
}

func (obj *patternFlowArpProtocolLengthCounter) setMsg(msg *otg.PatternFlowArpProtocolLengthCounter) PatternFlowArpProtocolLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpProtocolLengthCounter struct {
	obj *patternFlowArpProtocolLengthCounter
}

type marshalPatternFlowArpProtocolLengthCounter interface {
	// ToProto marshals PatternFlowArpProtocolLengthCounter to protobuf object *otg.PatternFlowArpProtocolLengthCounter
	ToProto() (*otg.PatternFlowArpProtocolLengthCounter, error)
	// ToPbText marshals PatternFlowArpProtocolLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpProtocolLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpProtocolLengthCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowArpProtocolLengthCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowArpProtocolLengthCounter struct {
	obj *patternFlowArpProtocolLengthCounter
}

type unMarshalPatternFlowArpProtocolLengthCounter interface {
	// FromProto unmarshals PatternFlowArpProtocolLengthCounter from protobuf object *otg.PatternFlowArpProtocolLengthCounter
	FromProto(msg *otg.PatternFlowArpProtocolLengthCounter) (PatternFlowArpProtocolLengthCounter, error)
	// FromPbText unmarshals PatternFlowArpProtocolLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpProtocolLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpProtocolLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpProtocolLengthCounter) Marshal() marshalPatternFlowArpProtocolLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpProtocolLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpProtocolLengthCounter) Unmarshal() unMarshalPatternFlowArpProtocolLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpProtocolLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpProtocolLengthCounter) ToProto() (*otg.PatternFlowArpProtocolLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpProtocolLengthCounter) FromProto(msg *otg.PatternFlowArpProtocolLengthCounter) (PatternFlowArpProtocolLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpProtocolLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpProtocolLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpProtocolLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpProtocolLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpProtocolLengthCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowArpProtocolLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpProtocolLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowArpProtocolLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpProtocolLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpProtocolLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpProtocolLengthCounter) Clone() (PatternFlowArpProtocolLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpProtocolLengthCounter()
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

// PatternFlowArpProtocolLengthCounter is integer counter pattern
type PatternFlowArpProtocolLengthCounter interface {
	Validation
	// msg marshals PatternFlowArpProtocolLengthCounter to protobuf object *otg.PatternFlowArpProtocolLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowArpProtocolLengthCounter
	// setMsg unmarshals PatternFlowArpProtocolLengthCounter from protobuf object *otg.PatternFlowArpProtocolLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpProtocolLengthCounter) PatternFlowArpProtocolLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowArpProtocolLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpProtocolLengthCounter
	// validate validates PatternFlowArpProtocolLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpProtocolLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowArpProtocolLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowArpProtocolLengthCounter
	SetStart(value uint32) PatternFlowArpProtocolLengthCounter
	// HasStart checks if Start has been set in PatternFlowArpProtocolLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowArpProtocolLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowArpProtocolLengthCounter
	SetStep(value uint32) PatternFlowArpProtocolLengthCounter
	// HasStep checks if Step has been set in PatternFlowArpProtocolLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowArpProtocolLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowArpProtocolLengthCounter
	SetCount(value uint32) PatternFlowArpProtocolLengthCounter
	// HasCount checks if Count has been set in PatternFlowArpProtocolLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowArpProtocolLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowArpProtocolLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowArpProtocolLengthCounter object
func (obj *patternFlowArpProtocolLengthCounter) SetStart(value uint32) PatternFlowArpProtocolLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowArpProtocolLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowArpProtocolLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowArpProtocolLengthCounter object
func (obj *patternFlowArpProtocolLengthCounter) SetStep(value uint32) PatternFlowArpProtocolLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowArpProtocolLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowArpProtocolLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowArpProtocolLengthCounter object
func (obj *patternFlowArpProtocolLengthCounter) SetCount(value uint32) PatternFlowArpProtocolLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowArpProtocolLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpProtocolLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpProtocolLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpProtocolLengthCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowArpProtocolLengthCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(4)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

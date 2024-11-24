package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPppProtocolTypeCounter *****
type patternFlowPppProtocolTypeCounter struct {
	validation
	obj          *otg.PatternFlowPppProtocolTypeCounter
	marshaller   marshalPatternFlowPppProtocolTypeCounter
	unMarshaller unMarshalPatternFlowPppProtocolTypeCounter
}

func NewPatternFlowPppProtocolTypeCounter() PatternFlowPppProtocolTypeCounter {
	obj := patternFlowPppProtocolTypeCounter{obj: &otg.PatternFlowPppProtocolTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPppProtocolTypeCounter) msg() *otg.PatternFlowPppProtocolTypeCounter {
	return obj.obj
}

func (obj *patternFlowPppProtocolTypeCounter) setMsg(msg *otg.PatternFlowPppProtocolTypeCounter) PatternFlowPppProtocolTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPppProtocolTypeCounter struct {
	obj *patternFlowPppProtocolTypeCounter
}

type marshalPatternFlowPppProtocolTypeCounter interface {
	// ToProto marshals PatternFlowPppProtocolTypeCounter to protobuf object *otg.PatternFlowPppProtocolTypeCounter
	ToProto() (*otg.PatternFlowPppProtocolTypeCounter, error)
	// ToPbText marshals PatternFlowPppProtocolTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPppProtocolTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPppProtocolTypeCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowPppProtocolTypeCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowPppProtocolTypeCounter struct {
	obj *patternFlowPppProtocolTypeCounter
}

type unMarshalPatternFlowPppProtocolTypeCounter interface {
	// FromProto unmarshals PatternFlowPppProtocolTypeCounter from protobuf object *otg.PatternFlowPppProtocolTypeCounter
	FromProto(msg *otg.PatternFlowPppProtocolTypeCounter) (PatternFlowPppProtocolTypeCounter, error)
	// FromPbText unmarshals PatternFlowPppProtocolTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPppProtocolTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPppProtocolTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPppProtocolTypeCounter) Marshal() marshalPatternFlowPppProtocolTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPppProtocolTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPppProtocolTypeCounter) Unmarshal() unMarshalPatternFlowPppProtocolTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPppProtocolTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPppProtocolTypeCounter) ToProto() (*otg.PatternFlowPppProtocolTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPppProtocolTypeCounter) FromProto(msg *otg.PatternFlowPppProtocolTypeCounter) (PatternFlowPppProtocolTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPppProtocolTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPppProtocolTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowPppProtocolTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPppProtocolTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowPppProtocolTypeCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowPppProtocolTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPppProtocolTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowPppProtocolTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPppProtocolTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPppProtocolTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPppProtocolTypeCounter) Clone() (PatternFlowPppProtocolTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPppProtocolTypeCounter()
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

// PatternFlowPppProtocolTypeCounter is integer counter pattern
type PatternFlowPppProtocolTypeCounter interface {
	Validation
	// msg marshals PatternFlowPppProtocolTypeCounter to protobuf object *otg.PatternFlowPppProtocolTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowPppProtocolTypeCounter
	// setMsg unmarshals PatternFlowPppProtocolTypeCounter from protobuf object *otg.PatternFlowPppProtocolTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPppProtocolTypeCounter) PatternFlowPppProtocolTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowPppProtocolTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPppProtocolTypeCounter
	// validate validates PatternFlowPppProtocolTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPppProtocolTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowPppProtocolTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowPppProtocolTypeCounter
	SetStart(value uint32) PatternFlowPppProtocolTypeCounter
	// HasStart checks if Start has been set in PatternFlowPppProtocolTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowPppProtocolTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowPppProtocolTypeCounter
	SetStep(value uint32) PatternFlowPppProtocolTypeCounter
	// HasStep checks if Step has been set in PatternFlowPppProtocolTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowPppProtocolTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowPppProtocolTypeCounter
	SetCount(value uint32) PatternFlowPppProtocolTypeCounter
	// HasCount checks if Count has been set in PatternFlowPppProtocolTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPppProtocolTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPppProtocolTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowPppProtocolTypeCounter object
func (obj *patternFlowPppProtocolTypeCounter) SetStart(value uint32) PatternFlowPppProtocolTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPppProtocolTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPppProtocolTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowPppProtocolTypeCounter object
func (obj *patternFlowPppProtocolTypeCounter) SetStep(value uint32) PatternFlowPppProtocolTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPppProtocolTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPppProtocolTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowPppProtocolTypeCounter object
func (obj *patternFlowPppProtocolTypeCounter) SetCount(value uint32) PatternFlowPppProtocolTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowPppProtocolTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPppProtocolTypeCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPppProtocolTypeCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPppProtocolTypeCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowPppProtocolTypeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(33)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

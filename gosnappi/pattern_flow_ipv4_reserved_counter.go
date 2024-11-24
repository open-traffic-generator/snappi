package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4ReservedCounter *****
type patternFlowIpv4ReservedCounter struct {
	validation
	obj          *otg.PatternFlowIpv4ReservedCounter
	marshaller   marshalPatternFlowIpv4ReservedCounter
	unMarshaller unMarshalPatternFlowIpv4ReservedCounter
}

func NewPatternFlowIpv4ReservedCounter() PatternFlowIpv4ReservedCounter {
	obj := patternFlowIpv4ReservedCounter{obj: &otg.PatternFlowIpv4ReservedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4ReservedCounter) msg() *otg.PatternFlowIpv4ReservedCounter {
	return obj.obj
}

func (obj *patternFlowIpv4ReservedCounter) setMsg(msg *otg.PatternFlowIpv4ReservedCounter) PatternFlowIpv4ReservedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4ReservedCounter struct {
	obj *patternFlowIpv4ReservedCounter
}

type marshalPatternFlowIpv4ReservedCounter interface {
	// ToProto marshals PatternFlowIpv4ReservedCounter to protobuf object *otg.PatternFlowIpv4ReservedCounter
	ToProto() (*otg.PatternFlowIpv4ReservedCounter, error)
	// ToPbText marshals PatternFlowIpv4ReservedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4ReservedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4ReservedCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4ReservedCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4ReservedCounter struct {
	obj *patternFlowIpv4ReservedCounter
}

type unMarshalPatternFlowIpv4ReservedCounter interface {
	// FromProto unmarshals PatternFlowIpv4ReservedCounter from protobuf object *otg.PatternFlowIpv4ReservedCounter
	FromProto(msg *otg.PatternFlowIpv4ReservedCounter) (PatternFlowIpv4ReservedCounter, error)
	// FromPbText unmarshals PatternFlowIpv4ReservedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4ReservedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4ReservedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4ReservedCounter) Marshal() marshalPatternFlowIpv4ReservedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4ReservedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4ReservedCounter) Unmarshal() unMarshalPatternFlowIpv4ReservedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4ReservedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4ReservedCounter) ToProto() (*otg.PatternFlowIpv4ReservedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4ReservedCounter) FromProto(msg *otg.PatternFlowIpv4ReservedCounter) (PatternFlowIpv4ReservedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4ReservedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4ReservedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4ReservedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4ReservedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4ReservedCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4ReservedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4ReservedCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4ReservedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4ReservedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4ReservedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4ReservedCounter) Clone() (PatternFlowIpv4ReservedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4ReservedCounter()
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

// PatternFlowIpv4ReservedCounter is integer counter pattern
type PatternFlowIpv4ReservedCounter interface {
	Validation
	// msg marshals PatternFlowIpv4ReservedCounter to protobuf object *otg.PatternFlowIpv4ReservedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4ReservedCounter
	// setMsg unmarshals PatternFlowIpv4ReservedCounter from protobuf object *otg.PatternFlowIpv4ReservedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4ReservedCounter) PatternFlowIpv4ReservedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4ReservedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4ReservedCounter
	// validate validates PatternFlowIpv4ReservedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4ReservedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4ReservedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4ReservedCounter
	SetStart(value uint32) PatternFlowIpv4ReservedCounter
	// HasStart checks if Start has been set in PatternFlowIpv4ReservedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4ReservedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4ReservedCounter
	SetStep(value uint32) PatternFlowIpv4ReservedCounter
	// HasStep checks if Step has been set in PatternFlowIpv4ReservedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4ReservedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4ReservedCounter
	SetCount(value uint32) PatternFlowIpv4ReservedCounter
	// HasCount checks if Count has been set in PatternFlowIpv4ReservedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4ReservedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4ReservedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4ReservedCounter object
func (obj *patternFlowIpv4ReservedCounter) SetStart(value uint32) PatternFlowIpv4ReservedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4ReservedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4ReservedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4ReservedCounter object
func (obj *patternFlowIpv4ReservedCounter) SetStep(value uint32) PatternFlowIpv4ReservedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4ReservedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4ReservedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4ReservedCounter object
func (obj *patternFlowIpv4ReservedCounter) SetCount(value uint32) PatternFlowIpv4ReservedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4ReservedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4ReservedCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4ReservedCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4ReservedCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4ReservedCounter) setDefault() {
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

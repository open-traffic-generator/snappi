package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TosPrecedenceCounter *****
type patternFlowIpv4TosPrecedenceCounter struct {
	validation
	obj          *otg.PatternFlowIpv4TosPrecedenceCounter
	marshaller   marshalPatternFlowIpv4TosPrecedenceCounter
	unMarshaller unMarshalPatternFlowIpv4TosPrecedenceCounter
}

func NewPatternFlowIpv4TosPrecedenceCounter() PatternFlowIpv4TosPrecedenceCounter {
	obj := patternFlowIpv4TosPrecedenceCounter{obj: &otg.PatternFlowIpv4TosPrecedenceCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TosPrecedenceCounter) msg() *otg.PatternFlowIpv4TosPrecedenceCounter {
	return obj.obj
}

func (obj *patternFlowIpv4TosPrecedenceCounter) setMsg(msg *otg.PatternFlowIpv4TosPrecedenceCounter) PatternFlowIpv4TosPrecedenceCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TosPrecedenceCounter struct {
	obj *patternFlowIpv4TosPrecedenceCounter
}

type marshalPatternFlowIpv4TosPrecedenceCounter interface {
	// ToProto marshals PatternFlowIpv4TosPrecedenceCounter to protobuf object *otg.PatternFlowIpv4TosPrecedenceCounter
	ToProto() (*otg.PatternFlowIpv4TosPrecedenceCounter, error)
	// ToPbText marshals PatternFlowIpv4TosPrecedenceCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TosPrecedenceCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TosPrecedenceCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4TosPrecedenceCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4TosPrecedenceCounter struct {
	obj *patternFlowIpv4TosPrecedenceCounter
}

type unMarshalPatternFlowIpv4TosPrecedenceCounter interface {
	// FromProto unmarshals PatternFlowIpv4TosPrecedenceCounter from protobuf object *otg.PatternFlowIpv4TosPrecedenceCounter
	FromProto(msg *otg.PatternFlowIpv4TosPrecedenceCounter) (PatternFlowIpv4TosPrecedenceCounter, error)
	// FromPbText unmarshals PatternFlowIpv4TosPrecedenceCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TosPrecedenceCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TosPrecedenceCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TosPrecedenceCounter) Marshal() marshalPatternFlowIpv4TosPrecedenceCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TosPrecedenceCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TosPrecedenceCounter) Unmarshal() unMarshalPatternFlowIpv4TosPrecedenceCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TosPrecedenceCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TosPrecedenceCounter) ToProto() (*otg.PatternFlowIpv4TosPrecedenceCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TosPrecedenceCounter) FromProto(msg *otg.PatternFlowIpv4TosPrecedenceCounter) (PatternFlowIpv4TosPrecedenceCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TosPrecedenceCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosPrecedenceCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TosPrecedenceCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosPrecedenceCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TosPrecedenceCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4TosPrecedenceCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosPrecedenceCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4TosPrecedenceCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosPrecedenceCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosPrecedenceCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TosPrecedenceCounter) Clone() (PatternFlowIpv4TosPrecedenceCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TosPrecedenceCounter()
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

// PatternFlowIpv4TosPrecedenceCounter is integer counter pattern
type PatternFlowIpv4TosPrecedenceCounter interface {
	Validation
	// msg marshals PatternFlowIpv4TosPrecedenceCounter to protobuf object *otg.PatternFlowIpv4TosPrecedenceCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TosPrecedenceCounter
	// setMsg unmarshals PatternFlowIpv4TosPrecedenceCounter from protobuf object *otg.PatternFlowIpv4TosPrecedenceCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TosPrecedenceCounter) PatternFlowIpv4TosPrecedenceCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TosPrecedenceCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TosPrecedenceCounter
	// validate validates PatternFlowIpv4TosPrecedenceCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TosPrecedenceCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4TosPrecedenceCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4TosPrecedenceCounter
	SetStart(value uint32) PatternFlowIpv4TosPrecedenceCounter
	// HasStart checks if Start has been set in PatternFlowIpv4TosPrecedenceCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4TosPrecedenceCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4TosPrecedenceCounter
	SetStep(value uint32) PatternFlowIpv4TosPrecedenceCounter
	// HasStep checks if Step has been set in PatternFlowIpv4TosPrecedenceCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4TosPrecedenceCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4TosPrecedenceCounter
	SetCount(value uint32) PatternFlowIpv4TosPrecedenceCounter
	// HasCount checks if Count has been set in PatternFlowIpv4TosPrecedenceCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4TosPrecedenceCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4TosPrecedenceCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4TosPrecedenceCounter object
func (obj *patternFlowIpv4TosPrecedenceCounter) SetStart(value uint32) PatternFlowIpv4TosPrecedenceCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4TosPrecedenceCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4TosPrecedenceCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4TosPrecedenceCounter object
func (obj *patternFlowIpv4TosPrecedenceCounter) SetStep(value uint32) PatternFlowIpv4TosPrecedenceCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4TosPrecedenceCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4TosPrecedenceCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4TosPrecedenceCounter object
func (obj *patternFlowIpv4TosPrecedenceCounter) SetCount(value uint32) PatternFlowIpv4TosPrecedenceCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4TosPrecedenceCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosPrecedenceCounter.Start <= 7 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosPrecedenceCounter.Step <= 7 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosPrecedenceCounter.Count <= 7 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4TosPrecedenceCounter) setDefault() {
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

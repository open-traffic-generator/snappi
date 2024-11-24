package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4ProtocolCounter *****
type patternFlowIpv4ProtocolCounter struct {
	validation
	obj          *otg.PatternFlowIpv4ProtocolCounter
	marshaller   marshalPatternFlowIpv4ProtocolCounter
	unMarshaller unMarshalPatternFlowIpv4ProtocolCounter
}

func NewPatternFlowIpv4ProtocolCounter() PatternFlowIpv4ProtocolCounter {
	obj := patternFlowIpv4ProtocolCounter{obj: &otg.PatternFlowIpv4ProtocolCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4ProtocolCounter) msg() *otg.PatternFlowIpv4ProtocolCounter {
	return obj.obj
}

func (obj *patternFlowIpv4ProtocolCounter) setMsg(msg *otg.PatternFlowIpv4ProtocolCounter) PatternFlowIpv4ProtocolCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4ProtocolCounter struct {
	obj *patternFlowIpv4ProtocolCounter
}

type marshalPatternFlowIpv4ProtocolCounter interface {
	// ToProto marshals PatternFlowIpv4ProtocolCounter to protobuf object *otg.PatternFlowIpv4ProtocolCounter
	ToProto() (*otg.PatternFlowIpv4ProtocolCounter, error)
	// ToPbText marshals PatternFlowIpv4ProtocolCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4ProtocolCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4ProtocolCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4ProtocolCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4ProtocolCounter struct {
	obj *patternFlowIpv4ProtocolCounter
}

type unMarshalPatternFlowIpv4ProtocolCounter interface {
	// FromProto unmarshals PatternFlowIpv4ProtocolCounter from protobuf object *otg.PatternFlowIpv4ProtocolCounter
	FromProto(msg *otg.PatternFlowIpv4ProtocolCounter) (PatternFlowIpv4ProtocolCounter, error)
	// FromPbText unmarshals PatternFlowIpv4ProtocolCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4ProtocolCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4ProtocolCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4ProtocolCounter) Marshal() marshalPatternFlowIpv4ProtocolCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4ProtocolCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4ProtocolCounter) Unmarshal() unMarshalPatternFlowIpv4ProtocolCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4ProtocolCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4ProtocolCounter) ToProto() (*otg.PatternFlowIpv4ProtocolCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4ProtocolCounter) FromProto(msg *otg.PatternFlowIpv4ProtocolCounter) (PatternFlowIpv4ProtocolCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4ProtocolCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4ProtocolCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4ProtocolCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4ProtocolCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4ProtocolCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4ProtocolCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4ProtocolCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4ProtocolCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4ProtocolCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4ProtocolCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4ProtocolCounter) Clone() (PatternFlowIpv4ProtocolCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4ProtocolCounter()
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

// PatternFlowIpv4ProtocolCounter is integer counter pattern
type PatternFlowIpv4ProtocolCounter interface {
	Validation
	// msg marshals PatternFlowIpv4ProtocolCounter to protobuf object *otg.PatternFlowIpv4ProtocolCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4ProtocolCounter
	// setMsg unmarshals PatternFlowIpv4ProtocolCounter from protobuf object *otg.PatternFlowIpv4ProtocolCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4ProtocolCounter) PatternFlowIpv4ProtocolCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4ProtocolCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4ProtocolCounter
	// validate validates PatternFlowIpv4ProtocolCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4ProtocolCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4ProtocolCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4ProtocolCounter
	SetStart(value uint32) PatternFlowIpv4ProtocolCounter
	// HasStart checks if Start has been set in PatternFlowIpv4ProtocolCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4ProtocolCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4ProtocolCounter
	SetStep(value uint32) PatternFlowIpv4ProtocolCounter
	// HasStep checks if Step has been set in PatternFlowIpv4ProtocolCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4ProtocolCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4ProtocolCounter
	SetCount(value uint32) PatternFlowIpv4ProtocolCounter
	// HasCount checks if Count has been set in PatternFlowIpv4ProtocolCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4ProtocolCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4ProtocolCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4ProtocolCounter object
func (obj *patternFlowIpv4ProtocolCounter) SetStart(value uint32) PatternFlowIpv4ProtocolCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4ProtocolCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4ProtocolCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4ProtocolCounter object
func (obj *patternFlowIpv4ProtocolCounter) SetStep(value uint32) PatternFlowIpv4ProtocolCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4ProtocolCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4ProtocolCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4ProtocolCounter object
func (obj *patternFlowIpv4ProtocolCounter) SetCount(value uint32) PatternFlowIpv4ProtocolCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4ProtocolCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4ProtocolCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4ProtocolCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4ProtocolCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4ProtocolCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(61)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

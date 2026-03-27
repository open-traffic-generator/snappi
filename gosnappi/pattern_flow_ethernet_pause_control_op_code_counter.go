package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetPauseControlOpCodeCounter *****
type patternFlowEthernetPauseControlOpCodeCounter struct {
	validation
	obj          *otg.PatternFlowEthernetPauseControlOpCodeCounter
	marshaller   marshalPatternFlowEthernetPauseControlOpCodeCounter
	unMarshaller unMarshalPatternFlowEthernetPauseControlOpCodeCounter
}

func NewPatternFlowEthernetPauseControlOpCodeCounter() PatternFlowEthernetPauseControlOpCodeCounter {
	obj := patternFlowEthernetPauseControlOpCodeCounter{obj: &otg.PatternFlowEthernetPauseControlOpCodeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetPauseControlOpCodeCounter) msg() *otg.PatternFlowEthernetPauseControlOpCodeCounter {
	return obj.obj
}

func (obj *patternFlowEthernetPauseControlOpCodeCounter) setMsg(msg *otg.PatternFlowEthernetPauseControlOpCodeCounter) PatternFlowEthernetPauseControlOpCodeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetPauseControlOpCodeCounter struct {
	obj *patternFlowEthernetPauseControlOpCodeCounter
}

type marshalPatternFlowEthernetPauseControlOpCodeCounter interface {
	// ToProto marshals PatternFlowEthernetPauseControlOpCodeCounter to protobuf object *otg.PatternFlowEthernetPauseControlOpCodeCounter
	ToProto() (*otg.PatternFlowEthernetPauseControlOpCodeCounter, error)
	// ToPbText marshals PatternFlowEthernetPauseControlOpCodeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetPauseControlOpCodeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetPauseControlOpCodeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowEthernetPauseControlOpCodeCounter struct {
	obj *patternFlowEthernetPauseControlOpCodeCounter
}

type unMarshalPatternFlowEthernetPauseControlOpCodeCounter interface {
	// FromProto unmarshals PatternFlowEthernetPauseControlOpCodeCounter from protobuf object *otg.PatternFlowEthernetPauseControlOpCodeCounter
	FromProto(msg *otg.PatternFlowEthernetPauseControlOpCodeCounter) (PatternFlowEthernetPauseControlOpCodeCounter, error)
	// FromPbText unmarshals PatternFlowEthernetPauseControlOpCodeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetPauseControlOpCodeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetPauseControlOpCodeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetPauseControlOpCodeCounter) Marshal() marshalPatternFlowEthernetPauseControlOpCodeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetPauseControlOpCodeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetPauseControlOpCodeCounter) Unmarshal() unMarshalPatternFlowEthernetPauseControlOpCodeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetPauseControlOpCodeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetPauseControlOpCodeCounter) ToProto() (*otg.PatternFlowEthernetPauseControlOpCodeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetPauseControlOpCodeCounter) FromProto(msg *otg.PatternFlowEthernetPauseControlOpCodeCounter) (PatternFlowEthernetPauseControlOpCodeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetPauseControlOpCodeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseControlOpCodeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetPauseControlOpCodeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseControlOpCodeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetPauseControlOpCodeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseControlOpCodeCounter) FromJson(value string) error {
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

func (obj *patternFlowEthernetPauseControlOpCodeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseControlOpCodeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseControlOpCodeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetPauseControlOpCodeCounter) Clone() (PatternFlowEthernetPauseControlOpCodeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetPauseControlOpCodeCounter()
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

// PatternFlowEthernetPauseControlOpCodeCounter is integer counter pattern
type PatternFlowEthernetPauseControlOpCodeCounter interface {
	Validation
	// msg marshals PatternFlowEthernetPauseControlOpCodeCounter to protobuf object *otg.PatternFlowEthernetPauseControlOpCodeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetPauseControlOpCodeCounter
	// setMsg unmarshals PatternFlowEthernetPauseControlOpCodeCounter from protobuf object *otg.PatternFlowEthernetPauseControlOpCodeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetPauseControlOpCodeCounter) PatternFlowEthernetPauseControlOpCodeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetPauseControlOpCodeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetPauseControlOpCodeCounter
	// validate validates PatternFlowEthernetPauseControlOpCodeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetPauseControlOpCodeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowEthernetPauseControlOpCodeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowEthernetPauseControlOpCodeCounter
	SetStart(value uint32) PatternFlowEthernetPauseControlOpCodeCounter
	// HasStart checks if Start has been set in PatternFlowEthernetPauseControlOpCodeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowEthernetPauseControlOpCodeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowEthernetPauseControlOpCodeCounter
	SetStep(value uint32) PatternFlowEthernetPauseControlOpCodeCounter
	// HasStep checks if Step has been set in PatternFlowEthernetPauseControlOpCodeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowEthernetPauseControlOpCodeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowEthernetPauseControlOpCodeCounter
	SetCount(value uint32) PatternFlowEthernetPauseControlOpCodeCounter
	// HasCount checks if Count has been set in PatternFlowEthernetPauseControlOpCodeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowEthernetPauseControlOpCodeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowEthernetPauseControlOpCodeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowEthernetPauseControlOpCodeCounter object
func (obj *patternFlowEthernetPauseControlOpCodeCounter) SetStart(value uint32) PatternFlowEthernetPauseControlOpCodeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowEthernetPauseControlOpCodeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowEthernetPauseControlOpCodeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowEthernetPauseControlOpCodeCounter object
func (obj *patternFlowEthernetPauseControlOpCodeCounter) SetStep(value uint32) PatternFlowEthernetPauseControlOpCodeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowEthernetPauseControlOpCodeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowEthernetPauseControlOpCodeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowEthernetPauseControlOpCodeCounter object
func (obj *patternFlowEthernetPauseControlOpCodeCounter) SetCount(value uint32) PatternFlowEthernetPauseControlOpCodeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowEthernetPauseControlOpCodeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseControlOpCodeCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseControlOpCodeCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseControlOpCodeCounter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowEthernetPauseControlOpCodeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(1)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

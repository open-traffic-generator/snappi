package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetPauseEtherTypeCounter *****
type patternFlowEthernetPauseEtherTypeCounter struct {
	validation
	obj          *otg.PatternFlowEthernetPauseEtherTypeCounter
	marshaller   marshalPatternFlowEthernetPauseEtherTypeCounter
	unMarshaller unMarshalPatternFlowEthernetPauseEtherTypeCounter
}

func NewPatternFlowEthernetPauseEtherTypeCounter() PatternFlowEthernetPauseEtherTypeCounter {
	obj := patternFlowEthernetPauseEtherTypeCounter{obj: &otg.PatternFlowEthernetPauseEtherTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetPauseEtherTypeCounter) msg() *otg.PatternFlowEthernetPauseEtherTypeCounter {
	return obj.obj
}

func (obj *patternFlowEthernetPauseEtherTypeCounter) setMsg(msg *otg.PatternFlowEthernetPauseEtherTypeCounter) PatternFlowEthernetPauseEtherTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetPauseEtherTypeCounter struct {
	obj *patternFlowEthernetPauseEtherTypeCounter
}

type marshalPatternFlowEthernetPauseEtherTypeCounter interface {
	// ToProto marshals PatternFlowEthernetPauseEtherTypeCounter to protobuf object *otg.PatternFlowEthernetPauseEtherTypeCounter
	ToProto() (*otg.PatternFlowEthernetPauseEtherTypeCounter, error)
	// ToPbText marshals PatternFlowEthernetPauseEtherTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetPauseEtherTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetPauseEtherTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowEthernetPauseEtherTypeCounter struct {
	obj *patternFlowEthernetPauseEtherTypeCounter
}

type unMarshalPatternFlowEthernetPauseEtherTypeCounter interface {
	// FromProto unmarshals PatternFlowEthernetPauseEtherTypeCounter from protobuf object *otg.PatternFlowEthernetPauseEtherTypeCounter
	FromProto(msg *otg.PatternFlowEthernetPauseEtherTypeCounter) (PatternFlowEthernetPauseEtherTypeCounter, error)
	// FromPbText unmarshals PatternFlowEthernetPauseEtherTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetPauseEtherTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetPauseEtherTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetPauseEtherTypeCounter) Marshal() marshalPatternFlowEthernetPauseEtherTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetPauseEtherTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetPauseEtherTypeCounter) Unmarshal() unMarshalPatternFlowEthernetPauseEtherTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetPauseEtherTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetPauseEtherTypeCounter) ToProto() (*otg.PatternFlowEthernetPauseEtherTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetPauseEtherTypeCounter) FromProto(msg *otg.PatternFlowEthernetPauseEtherTypeCounter) (PatternFlowEthernetPauseEtherTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetPauseEtherTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseEtherTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetPauseEtherTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseEtherTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetPauseEtherTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseEtherTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowEthernetPauseEtherTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseEtherTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseEtherTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetPauseEtherTypeCounter) Clone() (PatternFlowEthernetPauseEtherTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetPauseEtherTypeCounter()
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

// PatternFlowEthernetPauseEtherTypeCounter is integer counter pattern
type PatternFlowEthernetPauseEtherTypeCounter interface {
	Validation
	// msg marshals PatternFlowEthernetPauseEtherTypeCounter to protobuf object *otg.PatternFlowEthernetPauseEtherTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetPauseEtherTypeCounter
	// setMsg unmarshals PatternFlowEthernetPauseEtherTypeCounter from protobuf object *otg.PatternFlowEthernetPauseEtherTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetPauseEtherTypeCounter) PatternFlowEthernetPauseEtherTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetPauseEtherTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetPauseEtherTypeCounter
	// validate validates PatternFlowEthernetPauseEtherTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetPauseEtherTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowEthernetPauseEtherTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowEthernetPauseEtherTypeCounter
	SetStart(value uint32) PatternFlowEthernetPauseEtherTypeCounter
	// HasStart checks if Start has been set in PatternFlowEthernetPauseEtherTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowEthernetPauseEtherTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowEthernetPauseEtherTypeCounter
	SetStep(value uint32) PatternFlowEthernetPauseEtherTypeCounter
	// HasStep checks if Step has been set in PatternFlowEthernetPauseEtherTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowEthernetPauseEtherTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowEthernetPauseEtherTypeCounter
	SetCount(value uint32) PatternFlowEthernetPauseEtherTypeCounter
	// HasCount checks if Count has been set in PatternFlowEthernetPauseEtherTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowEthernetPauseEtherTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowEthernetPauseEtherTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowEthernetPauseEtherTypeCounter object
func (obj *patternFlowEthernetPauseEtherTypeCounter) SetStart(value uint32) PatternFlowEthernetPauseEtherTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowEthernetPauseEtherTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowEthernetPauseEtherTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowEthernetPauseEtherTypeCounter object
func (obj *patternFlowEthernetPauseEtherTypeCounter) SetStep(value uint32) PatternFlowEthernetPauseEtherTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowEthernetPauseEtherTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowEthernetPauseEtherTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowEthernetPauseEtherTypeCounter object
func (obj *patternFlowEthernetPauseEtherTypeCounter) SetCount(value uint32) PatternFlowEthernetPauseEtherTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowEthernetPauseEtherTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseEtherTypeCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseEtherTypeCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseEtherTypeCounter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowEthernetPauseEtherTypeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(34824)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

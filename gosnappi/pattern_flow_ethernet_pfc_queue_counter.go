package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetPfcQueueCounter *****
type patternFlowEthernetPfcQueueCounter struct {
	validation
	obj          *otg.PatternFlowEthernetPfcQueueCounter
	marshaller   marshalPatternFlowEthernetPfcQueueCounter
	unMarshaller unMarshalPatternFlowEthernetPfcQueueCounter
}

func NewPatternFlowEthernetPfcQueueCounter() PatternFlowEthernetPfcQueueCounter {
	obj := patternFlowEthernetPfcQueueCounter{obj: &otg.PatternFlowEthernetPfcQueueCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetPfcQueueCounter) msg() *otg.PatternFlowEthernetPfcQueueCounter {
	return obj.obj
}

func (obj *patternFlowEthernetPfcQueueCounter) setMsg(msg *otg.PatternFlowEthernetPfcQueueCounter) PatternFlowEthernetPfcQueueCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetPfcQueueCounter struct {
	obj *patternFlowEthernetPfcQueueCounter
}

type marshalPatternFlowEthernetPfcQueueCounter interface {
	// ToProto marshals PatternFlowEthernetPfcQueueCounter to protobuf object *otg.PatternFlowEthernetPfcQueueCounter
	ToProto() (*otg.PatternFlowEthernetPfcQueueCounter, error)
	// ToPbText marshals PatternFlowEthernetPfcQueueCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetPfcQueueCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetPfcQueueCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowEthernetPfcQueueCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowEthernetPfcQueueCounter struct {
	obj *patternFlowEthernetPfcQueueCounter
}

type unMarshalPatternFlowEthernetPfcQueueCounter interface {
	// FromProto unmarshals PatternFlowEthernetPfcQueueCounter from protobuf object *otg.PatternFlowEthernetPfcQueueCounter
	FromProto(msg *otg.PatternFlowEthernetPfcQueueCounter) (PatternFlowEthernetPfcQueueCounter, error)
	// FromPbText unmarshals PatternFlowEthernetPfcQueueCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetPfcQueueCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetPfcQueueCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetPfcQueueCounter) Marshal() marshalPatternFlowEthernetPfcQueueCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetPfcQueueCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetPfcQueueCounter) Unmarshal() unMarshalPatternFlowEthernetPfcQueueCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetPfcQueueCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetPfcQueueCounter) ToProto() (*otg.PatternFlowEthernetPfcQueueCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetPfcQueueCounter) FromProto(msg *otg.PatternFlowEthernetPfcQueueCounter) (PatternFlowEthernetPfcQueueCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetPfcQueueCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPfcQueueCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetPfcQueueCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPfcQueueCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetPfcQueueCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowEthernetPfcQueueCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPfcQueueCounter) FromJson(value string) error {
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

func (obj *patternFlowEthernetPfcQueueCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPfcQueueCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPfcQueueCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetPfcQueueCounter) Clone() (PatternFlowEthernetPfcQueueCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetPfcQueueCounter()
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

// PatternFlowEthernetPfcQueueCounter is integer counter pattern
type PatternFlowEthernetPfcQueueCounter interface {
	Validation
	// msg marshals PatternFlowEthernetPfcQueueCounter to protobuf object *otg.PatternFlowEthernetPfcQueueCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetPfcQueueCounter
	// setMsg unmarshals PatternFlowEthernetPfcQueueCounter from protobuf object *otg.PatternFlowEthernetPfcQueueCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetPfcQueueCounter) PatternFlowEthernetPfcQueueCounter
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetPfcQueueCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetPfcQueueCounter
	// validate validates PatternFlowEthernetPfcQueueCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetPfcQueueCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowEthernetPfcQueueCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowEthernetPfcQueueCounter
	SetStart(value uint32) PatternFlowEthernetPfcQueueCounter
	// HasStart checks if Start has been set in PatternFlowEthernetPfcQueueCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowEthernetPfcQueueCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowEthernetPfcQueueCounter
	SetStep(value uint32) PatternFlowEthernetPfcQueueCounter
	// HasStep checks if Step has been set in PatternFlowEthernetPfcQueueCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowEthernetPfcQueueCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowEthernetPfcQueueCounter
	SetCount(value uint32) PatternFlowEthernetPfcQueueCounter
	// HasCount checks if Count has been set in PatternFlowEthernetPfcQueueCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowEthernetPfcQueueCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowEthernetPfcQueueCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowEthernetPfcQueueCounter object
func (obj *patternFlowEthernetPfcQueueCounter) SetStart(value uint32) PatternFlowEthernetPfcQueueCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowEthernetPfcQueueCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowEthernetPfcQueueCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowEthernetPfcQueueCounter object
func (obj *patternFlowEthernetPfcQueueCounter) SetStep(value uint32) PatternFlowEthernetPfcQueueCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowEthernetPfcQueueCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowEthernetPfcQueueCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowEthernetPfcQueueCounter object
func (obj *patternFlowEthernetPfcQueueCounter) SetCount(value uint32) PatternFlowEthernetPfcQueueCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowEthernetPfcQueueCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPfcQueueCounter.Start <= 7 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPfcQueueCounter.Step <= 7 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPfcQueueCounter.Count <= 7 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowEthernetPfcQueueCounter) setDefault() {
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

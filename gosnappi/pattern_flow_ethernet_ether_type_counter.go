package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetEtherTypeCounter *****
type patternFlowEthernetEtherTypeCounter struct {
	validation
	obj          *otg.PatternFlowEthernetEtherTypeCounter
	marshaller   marshalPatternFlowEthernetEtherTypeCounter
	unMarshaller unMarshalPatternFlowEthernetEtherTypeCounter
}

func NewPatternFlowEthernetEtherTypeCounter() PatternFlowEthernetEtherTypeCounter {
	obj := patternFlowEthernetEtherTypeCounter{obj: &otg.PatternFlowEthernetEtherTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetEtherTypeCounter) msg() *otg.PatternFlowEthernetEtherTypeCounter {
	return obj.obj
}

func (obj *patternFlowEthernetEtherTypeCounter) setMsg(msg *otg.PatternFlowEthernetEtherTypeCounter) PatternFlowEthernetEtherTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetEtherTypeCounter struct {
	obj *patternFlowEthernetEtherTypeCounter
}

type marshalPatternFlowEthernetEtherTypeCounter interface {
	// ToProto marshals PatternFlowEthernetEtherTypeCounter to protobuf object *otg.PatternFlowEthernetEtherTypeCounter
	ToProto() (*otg.PatternFlowEthernetEtherTypeCounter, error)
	// ToPbText marshals PatternFlowEthernetEtherTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetEtherTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetEtherTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowEthernetEtherTypeCounter struct {
	obj *patternFlowEthernetEtherTypeCounter
}

type unMarshalPatternFlowEthernetEtherTypeCounter interface {
	// FromProto unmarshals PatternFlowEthernetEtherTypeCounter from protobuf object *otg.PatternFlowEthernetEtherTypeCounter
	FromProto(msg *otg.PatternFlowEthernetEtherTypeCounter) (PatternFlowEthernetEtherTypeCounter, error)
	// FromPbText unmarshals PatternFlowEthernetEtherTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetEtherTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetEtherTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetEtherTypeCounter) Marshal() marshalPatternFlowEthernetEtherTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetEtherTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetEtherTypeCounter) Unmarshal() unMarshalPatternFlowEthernetEtherTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetEtherTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetEtherTypeCounter) ToProto() (*otg.PatternFlowEthernetEtherTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetEtherTypeCounter) FromProto(msg *otg.PatternFlowEthernetEtherTypeCounter) (PatternFlowEthernetEtherTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetEtherTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetEtherTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetEtherTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetEtherTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetEtherTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetEtherTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowEthernetEtherTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetEtherTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetEtherTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetEtherTypeCounter) Clone() (PatternFlowEthernetEtherTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetEtherTypeCounter()
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

// PatternFlowEthernetEtherTypeCounter is integer counter pattern
type PatternFlowEthernetEtherTypeCounter interface {
	Validation
	// msg marshals PatternFlowEthernetEtherTypeCounter to protobuf object *otg.PatternFlowEthernetEtherTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetEtherTypeCounter
	// setMsg unmarshals PatternFlowEthernetEtherTypeCounter from protobuf object *otg.PatternFlowEthernetEtherTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetEtherTypeCounter) PatternFlowEthernetEtherTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetEtherTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetEtherTypeCounter
	// validate validates PatternFlowEthernetEtherTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetEtherTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowEthernetEtherTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowEthernetEtherTypeCounter
	SetStart(value uint32) PatternFlowEthernetEtherTypeCounter
	// HasStart checks if Start has been set in PatternFlowEthernetEtherTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowEthernetEtherTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowEthernetEtherTypeCounter
	SetStep(value uint32) PatternFlowEthernetEtherTypeCounter
	// HasStep checks if Step has been set in PatternFlowEthernetEtherTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowEthernetEtherTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowEthernetEtherTypeCounter
	SetCount(value uint32) PatternFlowEthernetEtherTypeCounter
	// HasCount checks if Count has been set in PatternFlowEthernetEtherTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowEthernetEtherTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowEthernetEtherTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowEthernetEtherTypeCounter object
func (obj *patternFlowEthernetEtherTypeCounter) SetStart(value uint32) PatternFlowEthernetEtherTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowEthernetEtherTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowEthernetEtherTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowEthernetEtherTypeCounter object
func (obj *patternFlowEthernetEtherTypeCounter) SetStep(value uint32) PatternFlowEthernetEtherTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowEthernetEtherTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowEthernetEtherTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowEthernetEtherTypeCounter object
func (obj *patternFlowEthernetEtherTypeCounter) SetCount(value uint32) PatternFlowEthernetEtherTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowEthernetEtherTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetEtherTypeCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetEtherTypeCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetEtherTypeCounter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowEthernetEtherTypeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(65535)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetPauseTimeCounter *****
type patternFlowEthernetPauseTimeCounter struct {
	validation
	obj          *otg.PatternFlowEthernetPauseTimeCounter
	marshaller   marshalPatternFlowEthernetPauseTimeCounter
	unMarshaller unMarshalPatternFlowEthernetPauseTimeCounter
}

func NewPatternFlowEthernetPauseTimeCounter() PatternFlowEthernetPauseTimeCounter {
	obj := patternFlowEthernetPauseTimeCounter{obj: &otg.PatternFlowEthernetPauseTimeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetPauseTimeCounter) msg() *otg.PatternFlowEthernetPauseTimeCounter {
	return obj.obj
}

func (obj *patternFlowEthernetPauseTimeCounter) setMsg(msg *otg.PatternFlowEthernetPauseTimeCounter) PatternFlowEthernetPauseTimeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetPauseTimeCounter struct {
	obj *patternFlowEthernetPauseTimeCounter
}

type marshalPatternFlowEthernetPauseTimeCounter interface {
	// ToProto marshals PatternFlowEthernetPauseTimeCounter to protobuf object *otg.PatternFlowEthernetPauseTimeCounter
	ToProto() (*otg.PatternFlowEthernetPauseTimeCounter, error)
	// ToPbText marshals PatternFlowEthernetPauseTimeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetPauseTimeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetPauseTimeCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowEthernetPauseTimeCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowEthernetPauseTimeCounter struct {
	obj *patternFlowEthernetPauseTimeCounter
}

type unMarshalPatternFlowEthernetPauseTimeCounter interface {
	// FromProto unmarshals PatternFlowEthernetPauseTimeCounter from protobuf object *otg.PatternFlowEthernetPauseTimeCounter
	FromProto(msg *otg.PatternFlowEthernetPauseTimeCounter) (PatternFlowEthernetPauseTimeCounter, error)
	// FromPbText unmarshals PatternFlowEthernetPauseTimeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetPauseTimeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetPauseTimeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetPauseTimeCounter) Marshal() marshalPatternFlowEthernetPauseTimeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetPauseTimeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetPauseTimeCounter) Unmarshal() unMarshalPatternFlowEthernetPauseTimeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetPauseTimeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetPauseTimeCounter) ToProto() (*otg.PatternFlowEthernetPauseTimeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetPauseTimeCounter) FromProto(msg *otg.PatternFlowEthernetPauseTimeCounter) (PatternFlowEthernetPauseTimeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetPauseTimeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseTimeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetPauseTimeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseTimeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetPauseTimeCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowEthernetPauseTimeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseTimeCounter) FromJson(value string) error {
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

func (obj *patternFlowEthernetPauseTimeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseTimeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseTimeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetPauseTimeCounter) Clone() (PatternFlowEthernetPauseTimeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetPauseTimeCounter()
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

// PatternFlowEthernetPauseTimeCounter is integer counter pattern
type PatternFlowEthernetPauseTimeCounter interface {
	Validation
	// msg marshals PatternFlowEthernetPauseTimeCounter to protobuf object *otg.PatternFlowEthernetPauseTimeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetPauseTimeCounter
	// setMsg unmarshals PatternFlowEthernetPauseTimeCounter from protobuf object *otg.PatternFlowEthernetPauseTimeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetPauseTimeCounter) PatternFlowEthernetPauseTimeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetPauseTimeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetPauseTimeCounter
	// validate validates PatternFlowEthernetPauseTimeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetPauseTimeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowEthernetPauseTimeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowEthernetPauseTimeCounter
	SetStart(value uint32) PatternFlowEthernetPauseTimeCounter
	// HasStart checks if Start has been set in PatternFlowEthernetPauseTimeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowEthernetPauseTimeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowEthernetPauseTimeCounter
	SetStep(value uint32) PatternFlowEthernetPauseTimeCounter
	// HasStep checks if Step has been set in PatternFlowEthernetPauseTimeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowEthernetPauseTimeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowEthernetPauseTimeCounter
	SetCount(value uint32) PatternFlowEthernetPauseTimeCounter
	// HasCount checks if Count has been set in PatternFlowEthernetPauseTimeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowEthernetPauseTimeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowEthernetPauseTimeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowEthernetPauseTimeCounter object
func (obj *patternFlowEthernetPauseTimeCounter) SetStart(value uint32) PatternFlowEthernetPauseTimeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowEthernetPauseTimeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowEthernetPauseTimeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowEthernetPauseTimeCounter object
func (obj *patternFlowEthernetPauseTimeCounter) SetStep(value uint32) PatternFlowEthernetPauseTimeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowEthernetPauseTimeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowEthernetPauseTimeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowEthernetPauseTimeCounter object
func (obj *patternFlowEthernetPauseTimeCounter) SetCount(value uint32) PatternFlowEthernetPauseTimeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowEthernetPauseTimeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseTimeCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseTimeCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseTimeCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowEthernetPauseTimeCounter) setDefault() {
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

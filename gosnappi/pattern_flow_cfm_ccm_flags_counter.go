package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmCcmFlagsCounter *****
type patternFlowCfmCcmFlagsCounter struct {
	validation
	obj          *otg.PatternFlowCfmCcmFlagsCounter
	marshaller   marshalPatternFlowCfmCcmFlagsCounter
	unMarshaller unMarshalPatternFlowCfmCcmFlagsCounter
}

func NewPatternFlowCfmCcmFlagsCounter() PatternFlowCfmCcmFlagsCounter {
	obj := patternFlowCfmCcmFlagsCounter{obj: &otg.PatternFlowCfmCcmFlagsCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmCcmFlagsCounter) msg() *otg.PatternFlowCfmCcmFlagsCounter {
	return obj.obj
}

func (obj *patternFlowCfmCcmFlagsCounter) setMsg(msg *otg.PatternFlowCfmCcmFlagsCounter) PatternFlowCfmCcmFlagsCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmCcmFlagsCounter struct {
	obj *patternFlowCfmCcmFlagsCounter
}

type marshalPatternFlowCfmCcmFlagsCounter interface {
	// ToProto marshals PatternFlowCfmCcmFlagsCounter to protobuf object *otg.PatternFlowCfmCcmFlagsCounter
	ToProto() (*otg.PatternFlowCfmCcmFlagsCounter, error)
	// ToPbText marshals PatternFlowCfmCcmFlagsCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmCcmFlagsCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmCcmFlagsCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmCcmFlagsCounter struct {
	obj *patternFlowCfmCcmFlagsCounter
}

type unMarshalPatternFlowCfmCcmFlagsCounter interface {
	// FromProto unmarshals PatternFlowCfmCcmFlagsCounter from protobuf object *otg.PatternFlowCfmCcmFlagsCounter
	FromProto(msg *otg.PatternFlowCfmCcmFlagsCounter) (PatternFlowCfmCcmFlagsCounter, error)
	// FromPbText unmarshals PatternFlowCfmCcmFlagsCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmCcmFlagsCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmCcmFlagsCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmCcmFlagsCounter) Marshal() marshalPatternFlowCfmCcmFlagsCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmCcmFlagsCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmCcmFlagsCounter) Unmarshal() unMarshalPatternFlowCfmCcmFlagsCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmCcmFlagsCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmCcmFlagsCounter) ToProto() (*otg.PatternFlowCfmCcmFlagsCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmCcmFlagsCounter) FromProto(msg *otg.PatternFlowCfmCcmFlagsCounter) (PatternFlowCfmCcmFlagsCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmCcmFlagsCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmCcmFlagsCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmCcmFlagsCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmCcmFlagsCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmCcmFlagsCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmCcmFlagsCounter) FromJson(value string) error {
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

func (obj *patternFlowCfmCcmFlagsCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmCcmFlagsCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmCcmFlagsCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmCcmFlagsCounter) Clone() (PatternFlowCfmCcmFlagsCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmCcmFlagsCounter()
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

// PatternFlowCfmCcmFlagsCounter is integer counter pattern
type PatternFlowCfmCcmFlagsCounter interface {
	Validation
	// msg marshals PatternFlowCfmCcmFlagsCounter to protobuf object *otg.PatternFlowCfmCcmFlagsCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmCcmFlagsCounter
	// setMsg unmarshals PatternFlowCfmCcmFlagsCounter from protobuf object *otg.PatternFlowCfmCcmFlagsCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmCcmFlagsCounter) PatternFlowCfmCcmFlagsCounter
	// provides marshal interface
	Marshal() marshalPatternFlowCfmCcmFlagsCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmCcmFlagsCounter
	// validate validates PatternFlowCfmCcmFlagsCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmCcmFlagsCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowCfmCcmFlagsCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowCfmCcmFlagsCounter
	SetStart(value uint32) PatternFlowCfmCcmFlagsCounter
	// HasStart checks if Start has been set in PatternFlowCfmCcmFlagsCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowCfmCcmFlagsCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowCfmCcmFlagsCounter
	SetStep(value uint32) PatternFlowCfmCcmFlagsCounter
	// HasStep checks if Step has been set in PatternFlowCfmCcmFlagsCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowCfmCcmFlagsCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowCfmCcmFlagsCounter
	SetCount(value uint32) PatternFlowCfmCcmFlagsCounter
	// HasCount checks if Count has been set in PatternFlowCfmCcmFlagsCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmCcmFlagsCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmCcmFlagsCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowCfmCcmFlagsCounter object
func (obj *patternFlowCfmCcmFlagsCounter) SetStart(value uint32) PatternFlowCfmCcmFlagsCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmCcmFlagsCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmCcmFlagsCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowCfmCcmFlagsCounter object
func (obj *patternFlowCfmCcmFlagsCounter) SetStep(value uint32) PatternFlowCfmCcmFlagsCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmCcmFlagsCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmCcmFlagsCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowCfmCcmFlagsCounter object
func (obj *patternFlowCfmCcmFlagsCounter) SetCount(value uint32) PatternFlowCfmCcmFlagsCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowCfmCcmFlagsCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmCcmFlagsCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmCcmFlagsCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmCcmFlagsCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowCfmCcmFlagsCounter) setDefault() {
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

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmMDNameSpecifiedMdNameLengthCounter *****
type patternFlowCfmMDNameSpecifiedMdNameLengthCounter struct {
	validation
	obj          *otg.PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	marshaller   marshalPatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	unMarshaller unMarshalPatternFlowCfmMDNameSpecifiedMdNameLengthCounter
}

func NewPatternFlowCfmMDNameSpecifiedMdNameLengthCounter() PatternFlowCfmMDNameSpecifiedMdNameLengthCounter {
	obj := patternFlowCfmMDNameSpecifiedMdNameLengthCounter{obj: &otg.PatternFlowCfmMDNameSpecifiedMdNameLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmMDNameSpecifiedMdNameLengthCounter) msg() *otg.PatternFlowCfmMDNameSpecifiedMdNameLengthCounter {
	return obj.obj
}

func (obj *patternFlowCfmMDNameSpecifiedMdNameLengthCounter) setMsg(msg *otg.PatternFlowCfmMDNameSpecifiedMdNameLengthCounter) PatternFlowCfmMDNameSpecifiedMdNameLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmMDNameSpecifiedMdNameLengthCounter struct {
	obj *patternFlowCfmMDNameSpecifiedMdNameLengthCounter
}

type marshalPatternFlowCfmMDNameSpecifiedMdNameLengthCounter interface {
	// ToProto marshals PatternFlowCfmMDNameSpecifiedMdNameLengthCounter to protobuf object *otg.PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	ToProto() (*otg.PatternFlowCfmMDNameSpecifiedMdNameLengthCounter, error)
	// ToPbText marshals PatternFlowCfmMDNameSpecifiedMdNameLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmMDNameSpecifiedMdNameLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmMDNameSpecifiedMdNameLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmMDNameSpecifiedMdNameLengthCounter struct {
	obj *patternFlowCfmMDNameSpecifiedMdNameLengthCounter
}

type unMarshalPatternFlowCfmMDNameSpecifiedMdNameLengthCounter interface {
	// FromProto unmarshals PatternFlowCfmMDNameSpecifiedMdNameLengthCounter from protobuf object *otg.PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	FromProto(msg *otg.PatternFlowCfmMDNameSpecifiedMdNameLengthCounter) (PatternFlowCfmMDNameSpecifiedMdNameLengthCounter, error)
	// FromPbText unmarshals PatternFlowCfmMDNameSpecifiedMdNameLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmMDNameSpecifiedMdNameLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmMDNameSpecifiedMdNameLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmMDNameSpecifiedMdNameLengthCounter) Marshal() marshalPatternFlowCfmMDNameSpecifiedMdNameLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmMDNameSpecifiedMdNameLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmMDNameSpecifiedMdNameLengthCounter) Unmarshal() unMarshalPatternFlowCfmMDNameSpecifiedMdNameLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmMDNameSpecifiedMdNameLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmMDNameSpecifiedMdNameLengthCounter) ToProto() (*otg.PatternFlowCfmMDNameSpecifiedMdNameLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmMDNameSpecifiedMdNameLengthCounter) FromProto(msg *otg.PatternFlowCfmMDNameSpecifiedMdNameLengthCounter) (PatternFlowCfmMDNameSpecifiedMdNameLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmMDNameSpecifiedMdNameLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmMDNameSpecifiedMdNameLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmMDNameSpecifiedMdNameLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmMDNameSpecifiedMdNameLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmMDNameSpecifiedMdNameLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmMDNameSpecifiedMdNameLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowCfmMDNameSpecifiedMdNameLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmMDNameSpecifiedMdNameLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmMDNameSpecifiedMdNameLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmMDNameSpecifiedMdNameLengthCounter) Clone() (PatternFlowCfmMDNameSpecifiedMdNameLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmMDNameSpecifiedMdNameLengthCounter()
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

// PatternFlowCfmMDNameSpecifiedMdNameLengthCounter is integer counter pattern
type PatternFlowCfmMDNameSpecifiedMdNameLengthCounter interface {
	Validation
	// msg marshals PatternFlowCfmMDNameSpecifiedMdNameLengthCounter to protobuf object *otg.PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	// setMsg unmarshals PatternFlowCfmMDNameSpecifiedMdNameLengthCounter from protobuf object *otg.PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmMDNameSpecifiedMdNameLengthCounter) PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	// validate validates PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmMDNameSpecifiedMdNameLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowCfmMDNameSpecifiedMdNameLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	SetStart(value uint32) PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	// HasStart checks if Start has been set in PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowCfmMDNameSpecifiedMdNameLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	SetStep(value uint32) PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	// HasStep checks if Step has been set in PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowCfmMDNameSpecifiedMdNameLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	SetCount(value uint32) PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	// HasCount checks if Count has been set in PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmMDNameSpecifiedMdNameLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmMDNameSpecifiedMdNameLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowCfmMDNameSpecifiedMdNameLengthCounter object
func (obj *patternFlowCfmMDNameSpecifiedMdNameLengthCounter) SetStart(value uint32) PatternFlowCfmMDNameSpecifiedMdNameLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmMDNameSpecifiedMdNameLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmMDNameSpecifiedMdNameLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowCfmMDNameSpecifiedMdNameLengthCounter object
func (obj *patternFlowCfmMDNameSpecifiedMdNameLengthCounter) SetStep(value uint32) PatternFlowCfmMDNameSpecifiedMdNameLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmMDNameSpecifiedMdNameLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmMDNameSpecifiedMdNameLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowCfmMDNameSpecifiedMdNameLengthCounter object
func (obj *patternFlowCfmMDNameSpecifiedMdNameLengthCounter) SetCount(value uint32) PatternFlowCfmMDNameSpecifiedMdNameLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowCfmMDNameSpecifiedMdNameLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmMDNameSpecifiedMdNameLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmMDNameSpecifiedMdNameLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmMDNameSpecifiedMdNameLengthCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowCfmMDNameSpecifiedMdNameLengthCounter) setDefault() {
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

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmMdLevelCounter *****
type patternFlowCfmMdLevelCounter struct {
	validation
	obj          *otg.PatternFlowCfmMdLevelCounter
	marshaller   marshalPatternFlowCfmMdLevelCounter
	unMarshaller unMarshalPatternFlowCfmMdLevelCounter
}

func NewPatternFlowCfmMdLevelCounter() PatternFlowCfmMdLevelCounter {
	obj := patternFlowCfmMdLevelCounter{obj: &otg.PatternFlowCfmMdLevelCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmMdLevelCounter) msg() *otg.PatternFlowCfmMdLevelCounter {
	return obj.obj
}

func (obj *patternFlowCfmMdLevelCounter) setMsg(msg *otg.PatternFlowCfmMdLevelCounter) PatternFlowCfmMdLevelCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmMdLevelCounter struct {
	obj *patternFlowCfmMdLevelCounter
}

type marshalPatternFlowCfmMdLevelCounter interface {
	// ToProto marshals PatternFlowCfmMdLevelCounter to protobuf object *otg.PatternFlowCfmMdLevelCounter
	ToProto() (*otg.PatternFlowCfmMdLevelCounter, error)
	// ToPbText marshals PatternFlowCfmMdLevelCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmMdLevelCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmMdLevelCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmMdLevelCounter struct {
	obj *patternFlowCfmMdLevelCounter
}

type unMarshalPatternFlowCfmMdLevelCounter interface {
	// FromProto unmarshals PatternFlowCfmMdLevelCounter from protobuf object *otg.PatternFlowCfmMdLevelCounter
	FromProto(msg *otg.PatternFlowCfmMdLevelCounter) (PatternFlowCfmMdLevelCounter, error)
	// FromPbText unmarshals PatternFlowCfmMdLevelCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmMdLevelCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmMdLevelCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmMdLevelCounter) Marshal() marshalPatternFlowCfmMdLevelCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmMdLevelCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmMdLevelCounter) Unmarshal() unMarshalPatternFlowCfmMdLevelCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmMdLevelCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmMdLevelCounter) ToProto() (*otg.PatternFlowCfmMdLevelCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmMdLevelCounter) FromProto(msg *otg.PatternFlowCfmMdLevelCounter) (PatternFlowCfmMdLevelCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmMdLevelCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmMdLevelCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmMdLevelCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmMdLevelCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmMdLevelCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmMdLevelCounter) FromJson(value string) error {
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

func (obj *patternFlowCfmMdLevelCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmMdLevelCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmMdLevelCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmMdLevelCounter) Clone() (PatternFlowCfmMdLevelCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmMdLevelCounter()
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

// PatternFlowCfmMdLevelCounter is integer counter pattern
type PatternFlowCfmMdLevelCounter interface {
	Validation
	// msg marshals PatternFlowCfmMdLevelCounter to protobuf object *otg.PatternFlowCfmMdLevelCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmMdLevelCounter
	// setMsg unmarshals PatternFlowCfmMdLevelCounter from protobuf object *otg.PatternFlowCfmMdLevelCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmMdLevelCounter) PatternFlowCfmMdLevelCounter
	// provides marshal interface
	Marshal() marshalPatternFlowCfmMdLevelCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmMdLevelCounter
	// validate validates PatternFlowCfmMdLevelCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmMdLevelCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowCfmMdLevelCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowCfmMdLevelCounter
	SetStart(value uint32) PatternFlowCfmMdLevelCounter
	// HasStart checks if Start has been set in PatternFlowCfmMdLevelCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowCfmMdLevelCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowCfmMdLevelCounter
	SetStep(value uint32) PatternFlowCfmMdLevelCounter
	// HasStep checks if Step has been set in PatternFlowCfmMdLevelCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowCfmMdLevelCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowCfmMdLevelCounter
	SetCount(value uint32) PatternFlowCfmMdLevelCounter
	// HasCount checks if Count has been set in PatternFlowCfmMdLevelCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmMdLevelCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmMdLevelCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowCfmMdLevelCounter object
func (obj *patternFlowCfmMdLevelCounter) SetStart(value uint32) PatternFlowCfmMdLevelCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmMdLevelCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmMdLevelCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowCfmMdLevelCounter object
func (obj *patternFlowCfmMdLevelCounter) SetStep(value uint32) PatternFlowCfmMdLevelCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmMdLevelCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmMdLevelCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowCfmMdLevelCounter object
func (obj *patternFlowCfmMdLevelCounter) SetCount(value uint32) PatternFlowCfmMdLevelCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowCfmMdLevelCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmMdLevelCounter.Start <= 7 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmMdLevelCounter.Step <= 7 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmMdLevelCounter.Count <= 8 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowCfmMdLevelCounter) setDefault() {
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

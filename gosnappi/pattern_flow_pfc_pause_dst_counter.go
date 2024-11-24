package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPauseDstCounter *****
type patternFlowPfcPauseDstCounter struct {
	validation
	obj          *otg.PatternFlowPfcPauseDstCounter
	marshaller   marshalPatternFlowPfcPauseDstCounter
	unMarshaller unMarshalPatternFlowPfcPauseDstCounter
}

func NewPatternFlowPfcPauseDstCounter() PatternFlowPfcPauseDstCounter {
	obj := patternFlowPfcPauseDstCounter{obj: &otg.PatternFlowPfcPauseDstCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPauseDstCounter) msg() *otg.PatternFlowPfcPauseDstCounter {
	return obj.obj
}

func (obj *patternFlowPfcPauseDstCounter) setMsg(msg *otg.PatternFlowPfcPauseDstCounter) PatternFlowPfcPauseDstCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPauseDstCounter struct {
	obj *patternFlowPfcPauseDstCounter
}

type marshalPatternFlowPfcPauseDstCounter interface {
	// ToProto marshals PatternFlowPfcPauseDstCounter to protobuf object *otg.PatternFlowPfcPauseDstCounter
	ToProto() (*otg.PatternFlowPfcPauseDstCounter, error)
	// ToPbText marshals PatternFlowPfcPauseDstCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPauseDstCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPauseDstCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPauseDstCounter struct {
	obj *patternFlowPfcPauseDstCounter
}

type unMarshalPatternFlowPfcPauseDstCounter interface {
	// FromProto unmarshals PatternFlowPfcPauseDstCounter from protobuf object *otg.PatternFlowPfcPauseDstCounter
	FromProto(msg *otg.PatternFlowPfcPauseDstCounter) (PatternFlowPfcPauseDstCounter, error)
	// FromPbText unmarshals PatternFlowPfcPauseDstCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPauseDstCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPauseDstCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPauseDstCounter) Marshal() marshalPatternFlowPfcPauseDstCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPauseDstCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPauseDstCounter) Unmarshal() unMarshalPatternFlowPfcPauseDstCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPauseDstCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPauseDstCounter) ToProto() (*otg.PatternFlowPfcPauseDstCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPauseDstCounter) FromProto(msg *otg.PatternFlowPfcPauseDstCounter) (PatternFlowPfcPauseDstCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPauseDstCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseDstCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPauseDstCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseDstCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPauseDstCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseDstCounter) FromJson(value string) error {
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

func (obj *patternFlowPfcPauseDstCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseDstCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseDstCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPauseDstCounter) Clone() (PatternFlowPfcPauseDstCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPauseDstCounter()
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

// PatternFlowPfcPauseDstCounter is mac counter pattern
type PatternFlowPfcPauseDstCounter interface {
	Validation
	// msg marshals PatternFlowPfcPauseDstCounter to protobuf object *otg.PatternFlowPfcPauseDstCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPauseDstCounter
	// setMsg unmarshals PatternFlowPfcPauseDstCounter from protobuf object *otg.PatternFlowPfcPauseDstCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPauseDstCounter) PatternFlowPfcPauseDstCounter
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPauseDstCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPauseDstCounter
	// validate validates PatternFlowPfcPauseDstCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPauseDstCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowPfcPauseDstCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowPfcPauseDstCounter
	SetStart(value string) PatternFlowPfcPauseDstCounter
	// HasStart checks if Start has been set in PatternFlowPfcPauseDstCounter
	HasStart() bool
	// Step returns string, set in PatternFlowPfcPauseDstCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowPfcPauseDstCounter
	SetStep(value string) PatternFlowPfcPauseDstCounter
	// HasStep checks if Step has been set in PatternFlowPfcPauseDstCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowPfcPauseDstCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowPfcPauseDstCounter
	SetCount(value uint32) PatternFlowPfcPauseDstCounter
	// HasCount checks if Count has been set in PatternFlowPfcPauseDstCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowPfcPauseDstCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowPfcPauseDstCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowPfcPauseDstCounter object
func (obj *patternFlowPfcPauseDstCounter) SetStart(value string) PatternFlowPfcPauseDstCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowPfcPauseDstCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowPfcPauseDstCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowPfcPauseDstCounter object
func (obj *patternFlowPfcPauseDstCounter) SetStep(value string) PatternFlowPfcPauseDstCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPauseDstCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPauseDstCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowPfcPauseDstCounter object
func (obj *patternFlowPfcPauseDstCounter) SetCount(value uint32) PatternFlowPfcPauseDstCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowPfcPauseDstCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowPfcPauseDstCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowPfcPauseDstCounter.Step"))
		}

	}

}

func (obj *patternFlowPfcPauseDstCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("01:80:c2:00:00:01")
	}
	if obj.obj.Step == nil {
		obj.SetStep("00:00:00:00:00:01")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

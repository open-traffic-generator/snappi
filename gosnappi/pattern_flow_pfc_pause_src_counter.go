package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPauseSrcCounter *****
type patternFlowPfcPauseSrcCounter struct {
	validation
	obj          *otg.PatternFlowPfcPauseSrcCounter
	marshaller   marshalPatternFlowPfcPauseSrcCounter
	unMarshaller unMarshalPatternFlowPfcPauseSrcCounter
}

func NewPatternFlowPfcPauseSrcCounter() PatternFlowPfcPauseSrcCounter {
	obj := patternFlowPfcPauseSrcCounter{obj: &otg.PatternFlowPfcPauseSrcCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPauseSrcCounter) msg() *otg.PatternFlowPfcPauseSrcCounter {
	return obj.obj
}

func (obj *patternFlowPfcPauseSrcCounter) setMsg(msg *otg.PatternFlowPfcPauseSrcCounter) PatternFlowPfcPauseSrcCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPauseSrcCounter struct {
	obj *patternFlowPfcPauseSrcCounter
}

type marshalPatternFlowPfcPauseSrcCounter interface {
	// ToProto marshals PatternFlowPfcPauseSrcCounter to protobuf object *otg.PatternFlowPfcPauseSrcCounter
	ToProto() (*otg.PatternFlowPfcPauseSrcCounter, error)
	// ToPbText marshals PatternFlowPfcPauseSrcCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPauseSrcCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPauseSrcCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPauseSrcCounter struct {
	obj *patternFlowPfcPauseSrcCounter
}

type unMarshalPatternFlowPfcPauseSrcCounter interface {
	// FromProto unmarshals PatternFlowPfcPauseSrcCounter from protobuf object *otg.PatternFlowPfcPauseSrcCounter
	FromProto(msg *otg.PatternFlowPfcPauseSrcCounter) (PatternFlowPfcPauseSrcCounter, error)
	// FromPbText unmarshals PatternFlowPfcPauseSrcCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPauseSrcCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPauseSrcCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPauseSrcCounter) Marshal() marshalPatternFlowPfcPauseSrcCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPauseSrcCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPauseSrcCounter) Unmarshal() unMarshalPatternFlowPfcPauseSrcCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPauseSrcCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPauseSrcCounter) ToProto() (*otg.PatternFlowPfcPauseSrcCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPauseSrcCounter) FromProto(msg *otg.PatternFlowPfcPauseSrcCounter) (PatternFlowPfcPauseSrcCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPauseSrcCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseSrcCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPauseSrcCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseSrcCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPauseSrcCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseSrcCounter) FromJson(value string) error {
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

func (obj *patternFlowPfcPauseSrcCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseSrcCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseSrcCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPauseSrcCounter) Clone() (PatternFlowPfcPauseSrcCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPauseSrcCounter()
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

// PatternFlowPfcPauseSrcCounter is mac counter pattern
type PatternFlowPfcPauseSrcCounter interface {
	Validation
	// msg marshals PatternFlowPfcPauseSrcCounter to protobuf object *otg.PatternFlowPfcPauseSrcCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPauseSrcCounter
	// setMsg unmarshals PatternFlowPfcPauseSrcCounter from protobuf object *otg.PatternFlowPfcPauseSrcCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPauseSrcCounter) PatternFlowPfcPauseSrcCounter
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPauseSrcCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPauseSrcCounter
	// validate validates PatternFlowPfcPauseSrcCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPauseSrcCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowPfcPauseSrcCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowPfcPauseSrcCounter
	SetStart(value string) PatternFlowPfcPauseSrcCounter
	// HasStart checks if Start has been set in PatternFlowPfcPauseSrcCounter
	HasStart() bool
	// Step returns string, set in PatternFlowPfcPauseSrcCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowPfcPauseSrcCounter
	SetStep(value string) PatternFlowPfcPauseSrcCounter
	// HasStep checks if Step has been set in PatternFlowPfcPauseSrcCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowPfcPauseSrcCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowPfcPauseSrcCounter
	SetCount(value uint32) PatternFlowPfcPauseSrcCounter
	// HasCount checks if Count has been set in PatternFlowPfcPauseSrcCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowPfcPauseSrcCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowPfcPauseSrcCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowPfcPauseSrcCounter object
func (obj *patternFlowPfcPauseSrcCounter) SetStart(value string) PatternFlowPfcPauseSrcCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowPfcPauseSrcCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowPfcPauseSrcCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowPfcPauseSrcCounter object
func (obj *patternFlowPfcPauseSrcCounter) SetStep(value string) PatternFlowPfcPauseSrcCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPauseSrcCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPauseSrcCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowPfcPauseSrcCounter object
func (obj *patternFlowPfcPauseSrcCounter) SetCount(value uint32) PatternFlowPfcPauseSrcCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowPfcPauseSrcCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowPfcPauseSrcCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowPfcPauseSrcCounter.Step"))
		}

	}

}

func (obj *patternFlowPfcPauseSrcCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("00:00:00:00:00:00")
	}
	if obj.obj.Step == nil {
		obj.SetStep("00:00:00:00:00:01")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

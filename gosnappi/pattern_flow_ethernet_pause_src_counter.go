package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetPauseSrcCounter *****
type patternFlowEthernetPauseSrcCounter struct {
	validation
	obj          *otg.PatternFlowEthernetPauseSrcCounter
	marshaller   marshalPatternFlowEthernetPauseSrcCounter
	unMarshaller unMarshalPatternFlowEthernetPauseSrcCounter
}

func NewPatternFlowEthernetPauseSrcCounter() PatternFlowEthernetPauseSrcCounter {
	obj := patternFlowEthernetPauseSrcCounter{obj: &otg.PatternFlowEthernetPauseSrcCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetPauseSrcCounter) msg() *otg.PatternFlowEthernetPauseSrcCounter {
	return obj.obj
}

func (obj *patternFlowEthernetPauseSrcCounter) setMsg(msg *otg.PatternFlowEthernetPauseSrcCounter) PatternFlowEthernetPauseSrcCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetPauseSrcCounter struct {
	obj *patternFlowEthernetPauseSrcCounter
}

type marshalPatternFlowEthernetPauseSrcCounter interface {
	// ToProto marshals PatternFlowEthernetPauseSrcCounter to protobuf object *otg.PatternFlowEthernetPauseSrcCounter
	ToProto() (*otg.PatternFlowEthernetPauseSrcCounter, error)
	// ToPbText marshals PatternFlowEthernetPauseSrcCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetPauseSrcCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetPauseSrcCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowEthernetPauseSrcCounter struct {
	obj *patternFlowEthernetPauseSrcCounter
}

type unMarshalPatternFlowEthernetPauseSrcCounter interface {
	// FromProto unmarshals PatternFlowEthernetPauseSrcCounter from protobuf object *otg.PatternFlowEthernetPauseSrcCounter
	FromProto(msg *otg.PatternFlowEthernetPauseSrcCounter) (PatternFlowEthernetPauseSrcCounter, error)
	// FromPbText unmarshals PatternFlowEthernetPauseSrcCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetPauseSrcCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetPauseSrcCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetPauseSrcCounter) Marshal() marshalPatternFlowEthernetPauseSrcCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetPauseSrcCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetPauseSrcCounter) Unmarshal() unMarshalPatternFlowEthernetPauseSrcCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetPauseSrcCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetPauseSrcCounter) ToProto() (*otg.PatternFlowEthernetPauseSrcCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetPauseSrcCounter) FromProto(msg *otg.PatternFlowEthernetPauseSrcCounter) (PatternFlowEthernetPauseSrcCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetPauseSrcCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseSrcCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetPauseSrcCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseSrcCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetPauseSrcCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseSrcCounter) FromJson(value string) error {
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

func (obj *patternFlowEthernetPauseSrcCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseSrcCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseSrcCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetPauseSrcCounter) Clone() (PatternFlowEthernetPauseSrcCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetPauseSrcCounter()
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

// PatternFlowEthernetPauseSrcCounter is mac counter pattern
type PatternFlowEthernetPauseSrcCounter interface {
	Validation
	// msg marshals PatternFlowEthernetPauseSrcCounter to protobuf object *otg.PatternFlowEthernetPauseSrcCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetPauseSrcCounter
	// setMsg unmarshals PatternFlowEthernetPauseSrcCounter from protobuf object *otg.PatternFlowEthernetPauseSrcCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetPauseSrcCounter) PatternFlowEthernetPauseSrcCounter
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetPauseSrcCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetPauseSrcCounter
	// validate validates PatternFlowEthernetPauseSrcCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetPauseSrcCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowEthernetPauseSrcCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowEthernetPauseSrcCounter
	SetStart(value string) PatternFlowEthernetPauseSrcCounter
	// HasStart checks if Start has been set in PatternFlowEthernetPauseSrcCounter
	HasStart() bool
	// Step returns string, set in PatternFlowEthernetPauseSrcCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowEthernetPauseSrcCounter
	SetStep(value string) PatternFlowEthernetPauseSrcCounter
	// HasStep checks if Step has been set in PatternFlowEthernetPauseSrcCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowEthernetPauseSrcCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowEthernetPauseSrcCounter
	SetCount(value uint32) PatternFlowEthernetPauseSrcCounter
	// HasCount checks if Count has been set in PatternFlowEthernetPauseSrcCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowEthernetPauseSrcCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowEthernetPauseSrcCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowEthernetPauseSrcCounter object
func (obj *patternFlowEthernetPauseSrcCounter) SetStart(value string) PatternFlowEthernetPauseSrcCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowEthernetPauseSrcCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowEthernetPauseSrcCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowEthernetPauseSrcCounter object
func (obj *patternFlowEthernetPauseSrcCounter) SetStep(value string) PatternFlowEthernetPauseSrcCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowEthernetPauseSrcCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowEthernetPauseSrcCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowEthernetPauseSrcCounter object
func (obj *patternFlowEthernetPauseSrcCounter) SetCount(value uint32) PatternFlowEthernetPauseSrcCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowEthernetPauseSrcCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetPauseSrcCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetPauseSrcCounter.Step"))
		}

	}

}

func (obj *patternFlowEthernetPauseSrcCounter) setDefault() {
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

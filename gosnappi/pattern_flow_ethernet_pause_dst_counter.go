package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetPauseDstCounter *****
type patternFlowEthernetPauseDstCounter struct {
	validation
	obj          *otg.PatternFlowEthernetPauseDstCounter
	marshaller   marshalPatternFlowEthernetPauseDstCounter
	unMarshaller unMarshalPatternFlowEthernetPauseDstCounter
}

func NewPatternFlowEthernetPauseDstCounter() PatternFlowEthernetPauseDstCounter {
	obj := patternFlowEthernetPauseDstCounter{obj: &otg.PatternFlowEthernetPauseDstCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetPauseDstCounter) msg() *otg.PatternFlowEthernetPauseDstCounter {
	return obj.obj
}

func (obj *patternFlowEthernetPauseDstCounter) setMsg(msg *otg.PatternFlowEthernetPauseDstCounter) PatternFlowEthernetPauseDstCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetPauseDstCounter struct {
	obj *patternFlowEthernetPauseDstCounter
}

type marshalPatternFlowEthernetPauseDstCounter interface {
	// ToProto marshals PatternFlowEthernetPauseDstCounter to protobuf object *otg.PatternFlowEthernetPauseDstCounter
	ToProto() (*otg.PatternFlowEthernetPauseDstCounter, error)
	// ToPbText marshals PatternFlowEthernetPauseDstCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetPauseDstCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetPauseDstCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowEthernetPauseDstCounter struct {
	obj *patternFlowEthernetPauseDstCounter
}

type unMarshalPatternFlowEthernetPauseDstCounter interface {
	// FromProto unmarshals PatternFlowEthernetPauseDstCounter from protobuf object *otg.PatternFlowEthernetPauseDstCounter
	FromProto(msg *otg.PatternFlowEthernetPauseDstCounter) (PatternFlowEthernetPauseDstCounter, error)
	// FromPbText unmarshals PatternFlowEthernetPauseDstCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetPauseDstCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetPauseDstCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetPauseDstCounter) Marshal() marshalPatternFlowEthernetPauseDstCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetPauseDstCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetPauseDstCounter) Unmarshal() unMarshalPatternFlowEthernetPauseDstCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetPauseDstCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetPauseDstCounter) ToProto() (*otg.PatternFlowEthernetPauseDstCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetPauseDstCounter) FromProto(msg *otg.PatternFlowEthernetPauseDstCounter) (PatternFlowEthernetPauseDstCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetPauseDstCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseDstCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetPauseDstCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseDstCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetPauseDstCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseDstCounter) FromJson(value string) error {
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

func (obj *patternFlowEthernetPauseDstCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseDstCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseDstCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetPauseDstCounter) Clone() (PatternFlowEthernetPauseDstCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetPauseDstCounter()
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

// PatternFlowEthernetPauseDstCounter is mac counter pattern
type PatternFlowEthernetPauseDstCounter interface {
	Validation
	// msg marshals PatternFlowEthernetPauseDstCounter to protobuf object *otg.PatternFlowEthernetPauseDstCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetPauseDstCounter
	// setMsg unmarshals PatternFlowEthernetPauseDstCounter from protobuf object *otg.PatternFlowEthernetPauseDstCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetPauseDstCounter) PatternFlowEthernetPauseDstCounter
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetPauseDstCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetPauseDstCounter
	// validate validates PatternFlowEthernetPauseDstCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetPauseDstCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowEthernetPauseDstCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowEthernetPauseDstCounter
	SetStart(value string) PatternFlowEthernetPauseDstCounter
	// HasStart checks if Start has been set in PatternFlowEthernetPauseDstCounter
	HasStart() bool
	// Step returns string, set in PatternFlowEthernetPauseDstCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowEthernetPauseDstCounter
	SetStep(value string) PatternFlowEthernetPauseDstCounter
	// HasStep checks if Step has been set in PatternFlowEthernetPauseDstCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowEthernetPauseDstCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowEthernetPauseDstCounter
	SetCount(value uint32) PatternFlowEthernetPauseDstCounter
	// HasCount checks if Count has been set in PatternFlowEthernetPauseDstCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowEthernetPauseDstCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowEthernetPauseDstCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowEthernetPauseDstCounter object
func (obj *patternFlowEthernetPauseDstCounter) SetStart(value string) PatternFlowEthernetPauseDstCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowEthernetPauseDstCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowEthernetPauseDstCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowEthernetPauseDstCounter object
func (obj *patternFlowEthernetPauseDstCounter) SetStep(value string) PatternFlowEthernetPauseDstCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowEthernetPauseDstCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowEthernetPauseDstCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowEthernetPauseDstCounter object
func (obj *patternFlowEthernetPauseDstCounter) SetCount(value uint32) PatternFlowEthernetPauseDstCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowEthernetPauseDstCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetPauseDstCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetPauseDstCounter.Step"))
		}

	}

}

func (obj *patternFlowEthernetPauseDstCounter) setDefault() {
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

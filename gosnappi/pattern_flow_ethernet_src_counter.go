package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetSrcCounter *****
type patternFlowEthernetSrcCounter struct {
	validation
	obj          *otg.PatternFlowEthernetSrcCounter
	marshaller   marshalPatternFlowEthernetSrcCounter
	unMarshaller unMarshalPatternFlowEthernetSrcCounter
}

func NewPatternFlowEthernetSrcCounter() PatternFlowEthernetSrcCounter {
	obj := patternFlowEthernetSrcCounter{obj: &otg.PatternFlowEthernetSrcCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetSrcCounter) msg() *otg.PatternFlowEthernetSrcCounter {
	return obj.obj
}

func (obj *patternFlowEthernetSrcCounter) setMsg(msg *otg.PatternFlowEthernetSrcCounter) PatternFlowEthernetSrcCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetSrcCounter struct {
	obj *patternFlowEthernetSrcCounter
}

type marshalPatternFlowEthernetSrcCounter interface {
	// ToProto marshals PatternFlowEthernetSrcCounter to protobuf object *otg.PatternFlowEthernetSrcCounter
	ToProto() (*otg.PatternFlowEthernetSrcCounter, error)
	// ToPbText marshals PatternFlowEthernetSrcCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetSrcCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetSrcCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowEthernetSrcCounter struct {
	obj *patternFlowEthernetSrcCounter
}

type unMarshalPatternFlowEthernetSrcCounter interface {
	// FromProto unmarshals PatternFlowEthernetSrcCounter from protobuf object *otg.PatternFlowEthernetSrcCounter
	FromProto(msg *otg.PatternFlowEthernetSrcCounter) (PatternFlowEthernetSrcCounter, error)
	// FromPbText unmarshals PatternFlowEthernetSrcCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetSrcCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetSrcCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetSrcCounter) Marshal() marshalPatternFlowEthernetSrcCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetSrcCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetSrcCounter) Unmarshal() unMarshalPatternFlowEthernetSrcCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetSrcCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetSrcCounter) ToProto() (*otg.PatternFlowEthernetSrcCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetSrcCounter) FromProto(msg *otg.PatternFlowEthernetSrcCounter) (PatternFlowEthernetSrcCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetSrcCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetSrcCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetSrcCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetSrcCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetSrcCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetSrcCounter) FromJson(value string) error {
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

func (obj *patternFlowEthernetSrcCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetSrcCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetSrcCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetSrcCounter) Clone() (PatternFlowEthernetSrcCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetSrcCounter()
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

// PatternFlowEthernetSrcCounter is mac counter pattern
type PatternFlowEthernetSrcCounter interface {
	Validation
	// msg marshals PatternFlowEthernetSrcCounter to protobuf object *otg.PatternFlowEthernetSrcCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetSrcCounter
	// setMsg unmarshals PatternFlowEthernetSrcCounter from protobuf object *otg.PatternFlowEthernetSrcCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetSrcCounter) PatternFlowEthernetSrcCounter
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetSrcCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetSrcCounter
	// validate validates PatternFlowEthernetSrcCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetSrcCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowEthernetSrcCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowEthernetSrcCounter
	SetStart(value string) PatternFlowEthernetSrcCounter
	// HasStart checks if Start has been set in PatternFlowEthernetSrcCounter
	HasStart() bool
	// Step returns string, set in PatternFlowEthernetSrcCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowEthernetSrcCounter
	SetStep(value string) PatternFlowEthernetSrcCounter
	// HasStep checks if Step has been set in PatternFlowEthernetSrcCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowEthernetSrcCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowEthernetSrcCounter
	SetCount(value uint32) PatternFlowEthernetSrcCounter
	// HasCount checks if Count has been set in PatternFlowEthernetSrcCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowEthernetSrcCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowEthernetSrcCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowEthernetSrcCounter object
func (obj *patternFlowEthernetSrcCounter) SetStart(value string) PatternFlowEthernetSrcCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowEthernetSrcCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowEthernetSrcCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowEthernetSrcCounter object
func (obj *patternFlowEthernetSrcCounter) SetStep(value string) PatternFlowEthernetSrcCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowEthernetSrcCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowEthernetSrcCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowEthernetSrcCounter object
func (obj *patternFlowEthernetSrcCounter) SetCount(value uint32) PatternFlowEthernetSrcCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowEthernetSrcCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetSrcCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetSrcCounter.Step"))
		}

	}

}

func (obj *patternFlowEthernetSrcCounter) setDefault() {
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

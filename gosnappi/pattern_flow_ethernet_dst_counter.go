package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetDstCounter *****
type patternFlowEthernetDstCounter struct {
	validation
	obj          *otg.PatternFlowEthernetDstCounter
	marshaller   marshalPatternFlowEthernetDstCounter
	unMarshaller unMarshalPatternFlowEthernetDstCounter
}

func NewPatternFlowEthernetDstCounter() PatternFlowEthernetDstCounter {
	obj := patternFlowEthernetDstCounter{obj: &otg.PatternFlowEthernetDstCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetDstCounter) msg() *otg.PatternFlowEthernetDstCounter {
	return obj.obj
}

func (obj *patternFlowEthernetDstCounter) setMsg(msg *otg.PatternFlowEthernetDstCounter) PatternFlowEthernetDstCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetDstCounter struct {
	obj *patternFlowEthernetDstCounter
}

type marshalPatternFlowEthernetDstCounter interface {
	// ToProto marshals PatternFlowEthernetDstCounter to protobuf object *otg.PatternFlowEthernetDstCounter
	ToProto() (*otg.PatternFlowEthernetDstCounter, error)
	// ToPbText marshals PatternFlowEthernetDstCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetDstCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetDstCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowEthernetDstCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowEthernetDstCounter struct {
	obj *patternFlowEthernetDstCounter
}

type unMarshalPatternFlowEthernetDstCounter interface {
	// FromProto unmarshals PatternFlowEthernetDstCounter from protobuf object *otg.PatternFlowEthernetDstCounter
	FromProto(msg *otg.PatternFlowEthernetDstCounter) (PatternFlowEthernetDstCounter, error)
	// FromPbText unmarshals PatternFlowEthernetDstCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetDstCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetDstCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetDstCounter) Marshal() marshalPatternFlowEthernetDstCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetDstCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetDstCounter) Unmarshal() unMarshalPatternFlowEthernetDstCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetDstCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetDstCounter) ToProto() (*otg.PatternFlowEthernetDstCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetDstCounter) FromProto(msg *otg.PatternFlowEthernetDstCounter) (PatternFlowEthernetDstCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetDstCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetDstCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetDstCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetDstCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetDstCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowEthernetDstCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetDstCounter) FromJson(value string) error {
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

func (obj *patternFlowEthernetDstCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetDstCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetDstCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetDstCounter) Clone() (PatternFlowEthernetDstCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetDstCounter()
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

// PatternFlowEthernetDstCounter is mac counter pattern
type PatternFlowEthernetDstCounter interface {
	Validation
	// msg marshals PatternFlowEthernetDstCounter to protobuf object *otg.PatternFlowEthernetDstCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetDstCounter
	// setMsg unmarshals PatternFlowEthernetDstCounter from protobuf object *otg.PatternFlowEthernetDstCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetDstCounter) PatternFlowEthernetDstCounter
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetDstCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetDstCounter
	// validate validates PatternFlowEthernetDstCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetDstCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowEthernetDstCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowEthernetDstCounter
	SetStart(value string) PatternFlowEthernetDstCounter
	// HasStart checks if Start has been set in PatternFlowEthernetDstCounter
	HasStart() bool
	// Step returns string, set in PatternFlowEthernetDstCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowEthernetDstCounter
	SetStep(value string) PatternFlowEthernetDstCounter
	// HasStep checks if Step has been set in PatternFlowEthernetDstCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowEthernetDstCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowEthernetDstCounter
	SetCount(value uint32) PatternFlowEthernetDstCounter
	// HasCount checks if Count has been set in PatternFlowEthernetDstCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowEthernetDstCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowEthernetDstCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowEthernetDstCounter object
func (obj *patternFlowEthernetDstCounter) SetStart(value string) PatternFlowEthernetDstCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowEthernetDstCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowEthernetDstCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowEthernetDstCounter object
func (obj *patternFlowEthernetDstCounter) SetStep(value string) PatternFlowEthernetDstCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowEthernetDstCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowEthernetDstCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowEthernetDstCounter object
func (obj *patternFlowEthernetDstCounter) SetCount(value uint32) PatternFlowEthernetDstCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowEthernetDstCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetDstCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetDstCounter.Step"))
		}

	}

}

func (obj *patternFlowEthernetDstCounter) setDefault() {
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

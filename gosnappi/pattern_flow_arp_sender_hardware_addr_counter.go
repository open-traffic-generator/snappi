package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpSenderHardwareAddrCounter *****
type patternFlowArpSenderHardwareAddrCounter struct {
	validation
	obj          *otg.PatternFlowArpSenderHardwareAddrCounter
	marshaller   marshalPatternFlowArpSenderHardwareAddrCounter
	unMarshaller unMarshalPatternFlowArpSenderHardwareAddrCounter
}

func NewPatternFlowArpSenderHardwareAddrCounter() PatternFlowArpSenderHardwareAddrCounter {
	obj := patternFlowArpSenderHardwareAddrCounter{obj: &otg.PatternFlowArpSenderHardwareAddrCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpSenderHardwareAddrCounter) msg() *otg.PatternFlowArpSenderHardwareAddrCounter {
	return obj.obj
}

func (obj *patternFlowArpSenderHardwareAddrCounter) setMsg(msg *otg.PatternFlowArpSenderHardwareAddrCounter) PatternFlowArpSenderHardwareAddrCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpSenderHardwareAddrCounter struct {
	obj *patternFlowArpSenderHardwareAddrCounter
}

type marshalPatternFlowArpSenderHardwareAddrCounter interface {
	// ToProto marshals PatternFlowArpSenderHardwareAddrCounter to protobuf object *otg.PatternFlowArpSenderHardwareAddrCounter
	ToProto() (*otg.PatternFlowArpSenderHardwareAddrCounter, error)
	// ToPbText marshals PatternFlowArpSenderHardwareAddrCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpSenderHardwareAddrCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpSenderHardwareAddrCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowArpSenderHardwareAddrCounter struct {
	obj *patternFlowArpSenderHardwareAddrCounter
}

type unMarshalPatternFlowArpSenderHardwareAddrCounter interface {
	// FromProto unmarshals PatternFlowArpSenderHardwareAddrCounter from protobuf object *otg.PatternFlowArpSenderHardwareAddrCounter
	FromProto(msg *otg.PatternFlowArpSenderHardwareAddrCounter) (PatternFlowArpSenderHardwareAddrCounter, error)
	// FromPbText unmarshals PatternFlowArpSenderHardwareAddrCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpSenderHardwareAddrCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpSenderHardwareAddrCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpSenderHardwareAddrCounter) Marshal() marshalPatternFlowArpSenderHardwareAddrCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpSenderHardwareAddrCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpSenderHardwareAddrCounter) Unmarshal() unMarshalPatternFlowArpSenderHardwareAddrCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpSenderHardwareAddrCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpSenderHardwareAddrCounter) ToProto() (*otg.PatternFlowArpSenderHardwareAddrCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpSenderHardwareAddrCounter) FromProto(msg *otg.PatternFlowArpSenderHardwareAddrCounter) (PatternFlowArpSenderHardwareAddrCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpSenderHardwareAddrCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpSenderHardwareAddrCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpSenderHardwareAddrCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpSenderHardwareAddrCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpSenderHardwareAddrCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpSenderHardwareAddrCounter) FromJson(value string) error {
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

func (obj *patternFlowArpSenderHardwareAddrCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpSenderHardwareAddrCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpSenderHardwareAddrCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpSenderHardwareAddrCounter) Clone() (PatternFlowArpSenderHardwareAddrCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpSenderHardwareAddrCounter()
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

// PatternFlowArpSenderHardwareAddrCounter is mac counter pattern
type PatternFlowArpSenderHardwareAddrCounter interface {
	Validation
	// msg marshals PatternFlowArpSenderHardwareAddrCounter to protobuf object *otg.PatternFlowArpSenderHardwareAddrCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowArpSenderHardwareAddrCounter
	// setMsg unmarshals PatternFlowArpSenderHardwareAddrCounter from protobuf object *otg.PatternFlowArpSenderHardwareAddrCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpSenderHardwareAddrCounter) PatternFlowArpSenderHardwareAddrCounter
	// provides marshal interface
	Marshal() marshalPatternFlowArpSenderHardwareAddrCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpSenderHardwareAddrCounter
	// validate validates PatternFlowArpSenderHardwareAddrCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpSenderHardwareAddrCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowArpSenderHardwareAddrCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowArpSenderHardwareAddrCounter
	SetStart(value string) PatternFlowArpSenderHardwareAddrCounter
	// HasStart checks if Start has been set in PatternFlowArpSenderHardwareAddrCounter
	HasStart() bool
	// Step returns string, set in PatternFlowArpSenderHardwareAddrCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowArpSenderHardwareAddrCounter
	SetStep(value string) PatternFlowArpSenderHardwareAddrCounter
	// HasStep checks if Step has been set in PatternFlowArpSenderHardwareAddrCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowArpSenderHardwareAddrCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowArpSenderHardwareAddrCounter
	SetCount(value uint32) PatternFlowArpSenderHardwareAddrCounter
	// HasCount checks if Count has been set in PatternFlowArpSenderHardwareAddrCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowArpSenderHardwareAddrCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowArpSenderHardwareAddrCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowArpSenderHardwareAddrCounter object
func (obj *patternFlowArpSenderHardwareAddrCounter) SetStart(value string) PatternFlowArpSenderHardwareAddrCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowArpSenderHardwareAddrCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowArpSenderHardwareAddrCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowArpSenderHardwareAddrCounter object
func (obj *patternFlowArpSenderHardwareAddrCounter) SetStep(value string) PatternFlowArpSenderHardwareAddrCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowArpSenderHardwareAddrCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowArpSenderHardwareAddrCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowArpSenderHardwareAddrCounter object
func (obj *patternFlowArpSenderHardwareAddrCounter) SetCount(value uint32) PatternFlowArpSenderHardwareAddrCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowArpSenderHardwareAddrCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpSenderHardwareAddrCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpSenderHardwareAddrCounter.Step"))
		}

	}

}

func (obj *patternFlowArpSenderHardwareAddrCounter) setDefault() {
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

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpTargetHardwareAddrCounter *****
type patternFlowArpTargetHardwareAddrCounter struct {
	validation
	obj          *otg.PatternFlowArpTargetHardwareAddrCounter
	marshaller   marshalPatternFlowArpTargetHardwareAddrCounter
	unMarshaller unMarshalPatternFlowArpTargetHardwareAddrCounter
}

func NewPatternFlowArpTargetHardwareAddrCounter() PatternFlowArpTargetHardwareAddrCounter {
	obj := patternFlowArpTargetHardwareAddrCounter{obj: &otg.PatternFlowArpTargetHardwareAddrCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpTargetHardwareAddrCounter) msg() *otg.PatternFlowArpTargetHardwareAddrCounter {
	return obj.obj
}

func (obj *patternFlowArpTargetHardwareAddrCounter) setMsg(msg *otg.PatternFlowArpTargetHardwareAddrCounter) PatternFlowArpTargetHardwareAddrCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpTargetHardwareAddrCounter struct {
	obj *patternFlowArpTargetHardwareAddrCounter
}

type marshalPatternFlowArpTargetHardwareAddrCounter interface {
	// ToProto marshals PatternFlowArpTargetHardwareAddrCounter to protobuf object *otg.PatternFlowArpTargetHardwareAddrCounter
	ToProto() (*otg.PatternFlowArpTargetHardwareAddrCounter, error)
	// ToPbText marshals PatternFlowArpTargetHardwareAddrCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpTargetHardwareAddrCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpTargetHardwareAddrCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowArpTargetHardwareAddrCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowArpTargetHardwareAddrCounter struct {
	obj *patternFlowArpTargetHardwareAddrCounter
}

type unMarshalPatternFlowArpTargetHardwareAddrCounter interface {
	// FromProto unmarshals PatternFlowArpTargetHardwareAddrCounter from protobuf object *otg.PatternFlowArpTargetHardwareAddrCounter
	FromProto(msg *otg.PatternFlowArpTargetHardwareAddrCounter) (PatternFlowArpTargetHardwareAddrCounter, error)
	// FromPbText unmarshals PatternFlowArpTargetHardwareAddrCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpTargetHardwareAddrCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpTargetHardwareAddrCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpTargetHardwareAddrCounter) Marshal() marshalPatternFlowArpTargetHardwareAddrCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpTargetHardwareAddrCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpTargetHardwareAddrCounter) Unmarshal() unMarshalPatternFlowArpTargetHardwareAddrCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpTargetHardwareAddrCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpTargetHardwareAddrCounter) ToProto() (*otg.PatternFlowArpTargetHardwareAddrCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpTargetHardwareAddrCounter) FromProto(msg *otg.PatternFlowArpTargetHardwareAddrCounter) (PatternFlowArpTargetHardwareAddrCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpTargetHardwareAddrCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpTargetHardwareAddrCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpTargetHardwareAddrCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpTargetHardwareAddrCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpTargetHardwareAddrCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowArpTargetHardwareAddrCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpTargetHardwareAddrCounter) FromJson(value string) error {
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

func (obj *patternFlowArpTargetHardwareAddrCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpTargetHardwareAddrCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpTargetHardwareAddrCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpTargetHardwareAddrCounter) Clone() (PatternFlowArpTargetHardwareAddrCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpTargetHardwareAddrCounter()
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

// PatternFlowArpTargetHardwareAddrCounter is mac counter pattern
type PatternFlowArpTargetHardwareAddrCounter interface {
	Validation
	// msg marshals PatternFlowArpTargetHardwareAddrCounter to protobuf object *otg.PatternFlowArpTargetHardwareAddrCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowArpTargetHardwareAddrCounter
	// setMsg unmarshals PatternFlowArpTargetHardwareAddrCounter from protobuf object *otg.PatternFlowArpTargetHardwareAddrCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpTargetHardwareAddrCounter) PatternFlowArpTargetHardwareAddrCounter
	// provides marshal interface
	Marshal() marshalPatternFlowArpTargetHardwareAddrCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpTargetHardwareAddrCounter
	// validate validates PatternFlowArpTargetHardwareAddrCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpTargetHardwareAddrCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowArpTargetHardwareAddrCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowArpTargetHardwareAddrCounter
	SetStart(value string) PatternFlowArpTargetHardwareAddrCounter
	// HasStart checks if Start has been set in PatternFlowArpTargetHardwareAddrCounter
	HasStart() bool
	// Step returns string, set in PatternFlowArpTargetHardwareAddrCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowArpTargetHardwareAddrCounter
	SetStep(value string) PatternFlowArpTargetHardwareAddrCounter
	// HasStep checks if Step has been set in PatternFlowArpTargetHardwareAddrCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowArpTargetHardwareAddrCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowArpTargetHardwareAddrCounter
	SetCount(value uint32) PatternFlowArpTargetHardwareAddrCounter
	// HasCount checks if Count has been set in PatternFlowArpTargetHardwareAddrCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowArpTargetHardwareAddrCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowArpTargetHardwareAddrCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowArpTargetHardwareAddrCounter object
func (obj *patternFlowArpTargetHardwareAddrCounter) SetStart(value string) PatternFlowArpTargetHardwareAddrCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowArpTargetHardwareAddrCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowArpTargetHardwareAddrCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowArpTargetHardwareAddrCounter object
func (obj *patternFlowArpTargetHardwareAddrCounter) SetStep(value string) PatternFlowArpTargetHardwareAddrCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowArpTargetHardwareAddrCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowArpTargetHardwareAddrCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowArpTargetHardwareAddrCounter object
func (obj *patternFlowArpTargetHardwareAddrCounter) SetCount(value uint32) PatternFlowArpTargetHardwareAddrCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowArpTargetHardwareAddrCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpTargetHardwareAddrCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpTargetHardwareAddrCounter.Step"))
		}

	}

}

func (obj *patternFlowArpTargetHardwareAddrCounter) setDefault() {
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

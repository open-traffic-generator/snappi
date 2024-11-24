package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpHardwareTypeCounter *****
type patternFlowArpHardwareTypeCounter struct {
	validation
	obj          *otg.PatternFlowArpHardwareTypeCounter
	marshaller   marshalPatternFlowArpHardwareTypeCounter
	unMarshaller unMarshalPatternFlowArpHardwareTypeCounter
}

func NewPatternFlowArpHardwareTypeCounter() PatternFlowArpHardwareTypeCounter {
	obj := patternFlowArpHardwareTypeCounter{obj: &otg.PatternFlowArpHardwareTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpHardwareTypeCounter) msg() *otg.PatternFlowArpHardwareTypeCounter {
	return obj.obj
}

func (obj *patternFlowArpHardwareTypeCounter) setMsg(msg *otg.PatternFlowArpHardwareTypeCounter) PatternFlowArpHardwareTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpHardwareTypeCounter struct {
	obj *patternFlowArpHardwareTypeCounter
}

type marshalPatternFlowArpHardwareTypeCounter interface {
	// ToProto marshals PatternFlowArpHardwareTypeCounter to protobuf object *otg.PatternFlowArpHardwareTypeCounter
	ToProto() (*otg.PatternFlowArpHardwareTypeCounter, error)
	// ToPbText marshals PatternFlowArpHardwareTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpHardwareTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpHardwareTypeCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowArpHardwareTypeCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowArpHardwareTypeCounter struct {
	obj *patternFlowArpHardwareTypeCounter
}

type unMarshalPatternFlowArpHardwareTypeCounter interface {
	// FromProto unmarshals PatternFlowArpHardwareTypeCounter from protobuf object *otg.PatternFlowArpHardwareTypeCounter
	FromProto(msg *otg.PatternFlowArpHardwareTypeCounter) (PatternFlowArpHardwareTypeCounter, error)
	// FromPbText unmarshals PatternFlowArpHardwareTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpHardwareTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpHardwareTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpHardwareTypeCounter) Marshal() marshalPatternFlowArpHardwareTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpHardwareTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpHardwareTypeCounter) Unmarshal() unMarshalPatternFlowArpHardwareTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpHardwareTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpHardwareTypeCounter) ToProto() (*otg.PatternFlowArpHardwareTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpHardwareTypeCounter) FromProto(msg *otg.PatternFlowArpHardwareTypeCounter) (PatternFlowArpHardwareTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpHardwareTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpHardwareTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpHardwareTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpHardwareTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpHardwareTypeCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowArpHardwareTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpHardwareTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowArpHardwareTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpHardwareTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpHardwareTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpHardwareTypeCounter) Clone() (PatternFlowArpHardwareTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpHardwareTypeCounter()
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

// PatternFlowArpHardwareTypeCounter is integer counter pattern
type PatternFlowArpHardwareTypeCounter interface {
	Validation
	// msg marshals PatternFlowArpHardwareTypeCounter to protobuf object *otg.PatternFlowArpHardwareTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowArpHardwareTypeCounter
	// setMsg unmarshals PatternFlowArpHardwareTypeCounter from protobuf object *otg.PatternFlowArpHardwareTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpHardwareTypeCounter) PatternFlowArpHardwareTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowArpHardwareTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpHardwareTypeCounter
	// validate validates PatternFlowArpHardwareTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpHardwareTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowArpHardwareTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowArpHardwareTypeCounter
	SetStart(value uint32) PatternFlowArpHardwareTypeCounter
	// HasStart checks if Start has been set in PatternFlowArpHardwareTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowArpHardwareTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowArpHardwareTypeCounter
	SetStep(value uint32) PatternFlowArpHardwareTypeCounter
	// HasStep checks if Step has been set in PatternFlowArpHardwareTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowArpHardwareTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowArpHardwareTypeCounter
	SetCount(value uint32) PatternFlowArpHardwareTypeCounter
	// HasCount checks if Count has been set in PatternFlowArpHardwareTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowArpHardwareTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowArpHardwareTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowArpHardwareTypeCounter object
func (obj *patternFlowArpHardwareTypeCounter) SetStart(value uint32) PatternFlowArpHardwareTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowArpHardwareTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowArpHardwareTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowArpHardwareTypeCounter object
func (obj *patternFlowArpHardwareTypeCounter) SetStep(value uint32) PatternFlowArpHardwareTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowArpHardwareTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowArpHardwareTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowArpHardwareTypeCounter object
func (obj *patternFlowArpHardwareTypeCounter) SetCount(value uint32) PatternFlowArpHardwareTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowArpHardwareTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpHardwareTypeCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpHardwareTypeCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpHardwareTypeCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowArpHardwareTypeCounter) setDefault() {
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

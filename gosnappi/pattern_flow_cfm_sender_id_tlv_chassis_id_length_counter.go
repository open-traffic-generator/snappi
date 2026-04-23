package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmSenderIdTlvChassisIdLengthCounter *****
type patternFlowCfmSenderIdTlvChassisIdLengthCounter struct {
	validation
	obj          *otg.PatternFlowCfmSenderIdTlvChassisIdLengthCounter
	marshaller   marshalPatternFlowCfmSenderIdTlvChassisIdLengthCounter
	unMarshaller unMarshalPatternFlowCfmSenderIdTlvChassisIdLengthCounter
}

func NewPatternFlowCfmSenderIdTlvChassisIdLengthCounter() PatternFlowCfmSenderIdTlvChassisIdLengthCounter {
	obj := patternFlowCfmSenderIdTlvChassisIdLengthCounter{obj: &otg.PatternFlowCfmSenderIdTlvChassisIdLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmSenderIdTlvChassisIdLengthCounter) msg() *otg.PatternFlowCfmSenderIdTlvChassisIdLengthCounter {
	return obj.obj
}

func (obj *patternFlowCfmSenderIdTlvChassisIdLengthCounter) setMsg(msg *otg.PatternFlowCfmSenderIdTlvChassisIdLengthCounter) PatternFlowCfmSenderIdTlvChassisIdLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmSenderIdTlvChassisIdLengthCounter struct {
	obj *patternFlowCfmSenderIdTlvChassisIdLengthCounter
}

type marshalPatternFlowCfmSenderIdTlvChassisIdLengthCounter interface {
	// ToProto marshals PatternFlowCfmSenderIdTlvChassisIdLengthCounter to protobuf object *otg.PatternFlowCfmSenderIdTlvChassisIdLengthCounter
	ToProto() (*otg.PatternFlowCfmSenderIdTlvChassisIdLengthCounter, error)
	// ToPbText marshals PatternFlowCfmSenderIdTlvChassisIdLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmSenderIdTlvChassisIdLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmSenderIdTlvChassisIdLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmSenderIdTlvChassisIdLengthCounter struct {
	obj *patternFlowCfmSenderIdTlvChassisIdLengthCounter
}

type unMarshalPatternFlowCfmSenderIdTlvChassisIdLengthCounter interface {
	// FromProto unmarshals PatternFlowCfmSenderIdTlvChassisIdLengthCounter from protobuf object *otg.PatternFlowCfmSenderIdTlvChassisIdLengthCounter
	FromProto(msg *otg.PatternFlowCfmSenderIdTlvChassisIdLengthCounter) (PatternFlowCfmSenderIdTlvChassisIdLengthCounter, error)
	// FromPbText unmarshals PatternFlowCfmSenderIdTlvChassisIdLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmSenderIdTlvChassisIdLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmSenderIdTlvChassisIdLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmSenderIdTlvChassisIdLengthCounter) Marshal() marshalPatternFlowCfmSenderIdTlvChassisIdLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmSenderIdTlvChassisIdLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmSenderIdTlvChassisIdLengthCounter) Unmarshal() unMarshalPatternFlowCfmSenderIdTlvChassisIdLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmSenderIdTlvChassisIdLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmSenderIdTlvChassisIdLengthCounter) ToProto() (*otg.PatternFlowCfmSenderIdTlvChassisIdLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmSenderIdTlvChassisIdLengthCounter) FromProto(msg *otg.PatternFlowCfmSenderIdTlvChassisIdLengthCounter) (PatternFlowCfmSenderIdTlvChassisIdLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmSenderIdTlvChassisIdLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmSenderIdTlvChassisIdLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmSenderIdTlvChassisIdLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmSenderIdTlvChassisIdLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmSenderIdTlvChassisIdLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmSenderIdTlvChassisIdLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowCfmSenderIdTlvChassisIdLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmSenderIdTlvChassisIdLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmSenderIdTlvChassisIdLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmSenderIdTlvChassisIdLengthCounter) Clone() (PatternFlowCfmSenderIdTlvChassisIdLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmSenderIdTlvChassisIdLengthCounter()
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

// PatternFlowCfmSenderIdTlvChassisIdLengthCounter is integer counter pattern
type PatternFlowCfmSenderIdTlvChassisIdLengthCounter interface {
	Validation
	// msg marshals PatternFlowCfmSenderIdTlvChassisIdLengthCounter to protobuf object *otg.PatternFlowCfmSenderIdTlvChassisIdLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmSenderIdTlvChassisIdLengthCounter
	// setMsg unmarshals PatternFlowCfmSenderIdTlvChassisIdLengthCounter from protobuf object *otg.PatternFlowCfmSenderIdTlvChassisIdLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmSenderIdTlvChassisIdLengthCounter) PatternFlowCfmSenderIdTlvChassisIdLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowCfmSenderIdTlvChassisIdLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmSenderIdTlvChassisIdLengthCounter
	// validate validates PatternFlowCfmSenderIdTlvChassisIdLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmSenderIdTlvChassisIdLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowCfmSenderIdTlvChassisIdLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowCfmSenderIdTlvChassisIdLengthCounter
	SetStart(value uint32) PatternFlowCfmSenderIdTlvChassisIdLengthCounter
	// HasStart checks if Start has been set in PatternFlowCfmSenderIdTlvChassisIdLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowCfmSenderIdTlvChassisIdLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowCfmSenderIdTlvChassisIdLengthCounter
	SetStep(value uint32) PatternFlowCfmSenderIdTlvChassisIdLengthCounter
	// HasStep checks if Step has been set in PatternFlowCfmSenderIdTlvChassisIdLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowCfmSenderIdTlvChassisIdLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowCfmSenderIdTlvChassisIdLengthCounter
	SetCount(value uint32) PatternFlowCfmSenderIdTlvChassisIdLengthCounter
	// HasCount checks if Count has been set in PatternFlowCfmSenderIdTlvChassisIdLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmSenderIdTlvChassisIdLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmSenderIdTlvChassisIdLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowCfmSenderIdTlvChassisIdLengthCounter object
func (obj *patternFlowCfmSenderIdTlvChassisIdLengthCounter) SetStart(value uint32) PatternFlowCfmSenderIdTlvChassisIdLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmSenderIdTlvChassisIdLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmSenderIdTlvChassisIdLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowCfmSenderIdTlvChassisIdLengthCounter object
func (obj *patternFlowCfmSenderIdTlvChassisIdLengthCounter) SetStep(value uint32) PatternFlowCfmSenderIdTlvChassisIdLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmSenderIdTlvChassisIdLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmSenderIdTlvChassisIdLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowCfmSenderIdTlvChassisIdLengthCounter object
func (obj *patternFlowCfmSenderIdTlvChassisIdLengthCounter) SetCount(value uint32) PatternFlowCfmSenderIdTlvChassisIdLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowCfmSenderIdTlvChassisIdLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmSenderIdTlvChassisIdLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmSenderIdTlvChassisIdLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmSenderIdTlvChassisIdLengthCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowCfmSenderIdTlvChassisIdLengthCounter) setDefault() {
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

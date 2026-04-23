package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmSenderIdTlvManagementAddressLengthCounter *****
type patternFlowCfmSenderIdTlvManagementAddressLengthCounter struct {
	validation
	obj          *otg.PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	marshaller   marshalPatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	unMarshaller unMarshalPatternFlowCfmSenderIdTlvManagementAddressLengthCounter
}

func NewPatternFlowCfmSenderIdTlvManagementAddressLengthCounter() PatternFlowCfmSenderIdTlvManagementAddressLengthCounter {
	obj := patternFlowCfmSenderIdTlvManagementAddressLengthCounter{obj: &otg.PatternFlowCfmSenderIdTlvManagementAddressLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmSenderIdTlvManagementAddressLengthCounter) msg() *otg.PatternFlowCfmSenderIdTlvManagementAddressLengthCounter {
	return obj.obj
}

func (obj *patternFlowCfmSenderIdTlvManagementAddressLengthCounter) setMsg(msg *otg.PatternFlowCfmSenderIdTlvManagementAddressLengthCounter) PatternFlowCfmSenderIdTlvManagementAddressLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmSenderIdTlvManagementAddressLengthCounter struct {
	obj *patternFlowCfmSenderIdTlvManagementAddressLengthCounter
}

type marshalPatternFlowCfmSenderIdTlvManagementAddressLengthCounter interface {
	// ToProto marshals PatternFlowCfmSenderIdTlvManagementAddressLengthCounter to protobuf object *otg.PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	ToProto() (*otg.PatternFlowCfmSenderIdTlvManagementAddressLengthCounter, error)
	// ToPbText marshals PatternFlowCfmSenderIdTlvManagementAddressLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmSenderIdTlvManagementAddressLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmSenderIdTlvManagementAddressLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmSenderIdTlvManagementAddressLengthCounter struct {
	obj *patternFlowCfmSenderIdTlvManagementAddressLengthCounter
}

type unMarshalPatternFlowCfmSenderIdTlvManagementAddressLengthCounter interface {
	// FromProto unmarshals PatternFlowCfmSenderIdTlvManagementAddressLengthCounter from protobuf object *otg.PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	FromProto(msg *otg.PatternFlowCfmSenderIdTlvManagementAddressLengthCounter) (PatternFlowCfmSenderIdTlvManagementAddressLengthCounter, error)
	// FromPbText unmarshals PatternFlowCfmSenderIdTlvManagementAddressLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmSenderIdTlvManagementAddressLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmSenderIdTlvManagementAddressLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmSenderIdTlvManagementAddressLengthCounter) Marshal() marshalPatternFlowCfmSenderIdTlvManagementAddressLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmSenderIdTlvManagementAddressLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmSenderIdTlvManagementAddressLengthCounter) Unmarshal() unMarshalPatternFlowCfmSenderIdTlvManagementAddressLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmSenderIdTlvManagementAddressLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmSenderIdTlvManagementAddressLengthCounter) ToProto() (*otg.PatternFlowCfmSenderIdTlvManagementAddressLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmSenderIdTlvManagementAddressLengthCounter) FromProto(msg *otg.PatternFlowCfmSenderIdTlvManagementAddressLengthCounter) (PatternFlowCfmSenderIdTlvManagementAddressLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmSenderIdTlvManagementAddressLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmSenderIdTlvManagementAddressLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmSenderIdTlvManagementAddressLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmSenderIdTlvManagementAddressLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmSenderIdTlvManagementAddressLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmSenderIdTlvManagementAddressLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowCfmSenderIdTlvManagementAddressLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmSenderIdTlvManagementAddressLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmSenderIdTlvManagementAddressLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmSenderIdTlvManagementAddressLengthCounter) Clone() (PatternFlowCfmSenderIdTlvManagementAddressLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmSenderIdTlvManagementAddressLengthCounter()
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

// PatternFlowCfmSenderIdTlvManagementAddressLengthCounter is integer counter pattern
type PatternFlowCfmSenderIdTlvManagementAddressLengthCounter interface {
	Validation
	// msg marshals PatternFlowCfmSenderIdTlvManagementAddressLengthCounter to protobuf object *otg.PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	// setMsg unmarshals PatternFlowCfmSenderIdTlvManagementAddressLengthCounter from protobuf object *otg.PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmSenderIdTlvManagementAddressLengthCounter) PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	// validate validates PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmSenderIdTlvManagementAddressLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowCfmSenderIdTlvManagementAddressLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	SetStart(value uint32) PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	// HasStart checks if Start has been set in PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowCfmSenderIdTlvManagementAddressLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	SetStep(value uint32) PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	// HasStep checks if Step has been set in PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowCfmSenderIdTlvManagementAddressLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	SetCount(value uint32) PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	// HasCount checks if Count has been set in PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmSenderIdTlvManagementAddressLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmSenderIdTlvManagementAddressLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowCfmSenderIdTlvManagementAddressLengthCounter object
func (obj *patternFlowCfmSenderIdTlvManagementAddressLengthCounter) SetStart(value uint32) PatternFlowCfmSenderIdTlvManagementAddressLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmSenderIdTlvManagementAddressLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmSenderIdTlvManagementAddressLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowCfmSenderIdTlvManagementAddressLengthCounter object
func (obj *patternFlowCfmSenderIdTlvManagementAddressLengthCounter) SetStep(value uint32) PatternFlowCfmSenderIdTlvManagementAddressLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmSenderIdTlvManagementAddressLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmSenderIdTlvManagementAddressLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowCfmSenderIdTlvManagementAddressLengthCounter object
func (obj *patternFlowCfmSenderIdTlvManagementAddressLengthCounter) SetCount(value uint32) PatternFlowCfmSenderIdTlvManagementAddressLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowCfmSenderIdTlvManagementAddressLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmSenderIdTlvManagementAddressLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmSenderIdTlvManagementAddressLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmSenderIdTlvManagementAddressLengthCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowCfmSenderIdTlvManagementAddressLengthCounter) setDefault() {
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

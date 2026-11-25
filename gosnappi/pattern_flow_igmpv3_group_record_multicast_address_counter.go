package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3GroupRecordMulticastAddressCounter *****
type patternFlowIgmpv3GroupRecordMulticastAddressCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv3GroupRecordMulticastAddressCounter
	marshaller   marshalPatternFlowIgmpv3GroupRecordMulticastAddressCounter
	unMarshaller unMarshalPatternFlowIgmpv3GroupRecordMulticastAddressCounter
}

func NewPatternFlowIgmpv3GroupRecordMulticastAddressCounter() PatternFlowIgmpv3GroupRecordMulticastAddressCounter {
	obj := patternFlowIgmpv3GroupRecordMulticastAddressCounter{obj: &otg.PatternFlowIgmpv3GroupRecordMulticastAddressCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3GroupRecordMulticastAddressCounter) msg() *otg.PatternFlowIgmpv3GroupRecordMulticastAddressCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv3GroupRecordMulticastAddressCounter) setMsg(msg *otg.PatternFlowIgmpv3GroupRecordMulticastAddressCounter) PatternFlowIgmpv3GroupRecordMulticastAddressCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3GroupRecordMulticastAddressCounter struct {
	obj *patternFlowIgmpv3GroupRecordMulticastAddressCounter
}

type marshalPatternFlowIgmpv3GroupRecordMulticastAddressCounter interface {
	// ToProto marshals PatternFlowIgmpv3GroupRecordMulticastAddressCounter to protobuf object *otg.PatternFlowIgmpv3GroupRecordMulticastAddressCounter
	ToProto() (*otg.PatternFlowIgmpv3GroupRecordMulticastAddressCounter, error)
	// ToPbText marshals PatternFlowIgmpv3GroupRecordMulticastAddressCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3GroupRecordMulticastAddressCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3GroupRecordMulticastAddressCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3GroupRecordMulticastAddressCounter struct {
	obj *patternFlowIgmpv3GroupRecordMulticastAddressCounter
}

type unMarshalPatternFlowIgmpv3GroupRecordMulticastAddressCounter interface {
	// FromProto unmarshals PatternFlowIgmpv3GroupRecordMulticastAddressCounter from protobuf object *otg.PatternFlowIgmpv3GroupRecordMulticastAddressCounter
	FromProto(msg *otg.PatternFlowIgmpv3GroupRecordMulticastAddressCounter) (PatternFlowIgmpv3GroupRecordMulticastAddressCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv3GroupRecordMulticastAddressCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3GroupRecordMulticastAddressCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3GroupRecordMulticastAddressCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3GroupRecordMulticastAddressCounter) Marshal() marshalPatternFlowIgmpv3GroupRecordMulticastAddressCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3GroupRecordMulticastAddressCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3GroupRecordMulticastAddressCounter) Unmarshal() unMarshalPatternFlowIgmpv3GroupRecordMulticastAddressCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3GroupRecordMulticastAddressCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3GroupRecordMulticastAddressCounter) ToProto() (*otg.PatternFlowIgmpv3GroupRecordMulticastAddressCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3GroupRecordMulticastAddressCounter) FromProto(msg *otg.PatternFlowIgmpv3GroupRecordMulticastAddressCounter) (PatternFlowIgmpv3GroupRecordMulticastAddressCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3GroupRecordMulticastAddressCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordMulticastAddressCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3GroupRecordMulticastAddressCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordMulticastAddressCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3GroupRecordMulticastAddressCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordMulticastAddressCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3GroupRecordMulticastAddressCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3GroupRecordMulticastAddressCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3GroupRecordMulticastAddressCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3GroupRecordMulticastAddressCounter) Clone() (PatternFlowIgmpv3GroupRecordMulticastAddressCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3GroupRecordMulticastAddressCounter()
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

// PatternFlowIgmpv3GroupRecordMulticastAddressCounter is ipv4 counter pattern
type PatternFlowIgmpv3GroupRecordMulticastAddressCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv3GroupRecordMulticastAddressCounter to protobuf object *otg.PatternFlowIgmpv3GroupRecordMulticastAddressCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3GroupRecordMulticastAddressCounter
	// setMsg unmarshals PatternFlowIgmpv3GroupRecordMulticastAddressCounter from protobuf object *otg.PatternFlowIgmpv3GroupRecordMulticastAddressCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3GroupRecordMulticastAddressCounter) PatternFlowIgmpv3GroupRecordMulticastAddressCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3GroupRecordMulticastAddressCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3GroupRecordMulticastAddressCounter
	// validate validates PatternFlowIgmpv3GroupRecordMulticastAddressCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3GroupRecordMulticastAddressCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIgmpv3GroupRecordMulticastAddressCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIgmpv3GroupRecordMulticastAddressCounter
	SetStart(value string) PatternFlowIgmpv3GroupRecordMulticastAddressCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv3GroupRecordMulticastAddressCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIgmpv3GroupRecordMulticastAddressCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIgmpv3GroupRecordMulticastAddressCounter
	SetStep(value string) PatternFlowIgmpv3GroupRecordMulticastAddressCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv3GroupRecordMulticastAddressCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv3GroupRecordMulticastAddressCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv3GroupRecordMulticastAddressCounter
	SetCount(value uint32) PatternFlowIgmpv3GroupRecordMulticastAddressCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv3GroupRecordMulticastAddressCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIgmpv3GroupRecordMulticastAddressCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIgmpv3GroupRecordMulticastAddressCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIgmpv3GroupRecordMulticastAddressCounter object
func (obj *patternFlowIgmpv3GroupRecordMulticastAddressCounter) SetStart(value string) PatternFlowIgmpv3GroupRecordMulticastAddressCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIgmpv3GroupRecordMulticastAddressCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIgmpv3GroupRecordMulticastAddressCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIgmpv3GroupRecordMulticastAddressCounter object
func (obj *patternFlowIgmpv3GroupRecordMulticastAddressCounter) SetStep(value string) PatternFlowIgmpv3GroupRecordMulticastAddressCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3GroupRecordMulticastAddressCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3GroupRecordMulticastAddressCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv3GroupRecordMulticastAddressCounter object
func (obj *patternFlowIgmpv3GroupRecordMulticastAddressCounter) SetCount(value uint32) PatternFlowIgmpv3GroupRecordMulticastAddressCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv3GroupRecordMulticastAddressCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv3GroupRecordMulticastAddressCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv3GroupRecordMulticastAddressCounter.Step"))
		}

	}

}

func (obj *patternFlowIgmpv3GroupRecordMulticastAddressCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("0.0.0.0")
	}
	if obj.obj.Step == nil {
		obj.SetStep("0.0.0.1")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

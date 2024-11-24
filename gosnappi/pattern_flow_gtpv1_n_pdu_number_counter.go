package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1NPduNumberCounter *****
type patternFlowGtpv1NPduNumberCounter struct {
	validation
	obj          *otg.PatternFlowGtpv1NPduNumberCounter
	marshaller   marshalPatternFlowGtpv1NPduNumberCounter
	unMarshaller unMarshalPatternFlowGtpv1NPduNumberCounter
}

func NewPatternFlowGtpv1NPduNumberCounter() PatternFlowGtpv1NPduNumberCounter {
	obj := patternFlowGtpv1NPduNumberCounter{obj: &otg.PatternFlowGtpv1NPduNumberCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1NPduNumberCounter) msg() *otg.PatternFlowGtpv1NPduNumberCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1NPduNumberCounter) setMsg(msg *otg.PatternFlowGtpv1NPduNumberCounter) PatternFlowGtpv1NPduNumberCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1NPduNumberCounter struct {
	obj *patternFlowGtpv1NPduNumberCounter
}

type marshalPatternFlowGtpv1NPduNumberCounter interface {
	// ToProto marshals PatternFlowGtpv1NPduNumberCounter to protobuf object *otg.PatternFlowGtpv1NPduNumberCounter
	ToProto() (*otg.PatternFlowGtpv1NPduNumberCounter, error)
	// ToPbText marshals PatternFlowGtpv1NPduNumberCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1NPduNumberCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1NPduNumberCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv1NPduNumberCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv1NPduNumberCounter struct {
	obj *patternFlowGtpv1NPduNumberCounter
}

type unMarshalPatternFlowGtpv1NPduNumberCounter interface {
	// FromProto unmarshals PatternFlowGtpv1NPduNumberCounter from protobuf object *otg.PatternFlowGtpv1NPduNumberCounter
	FromProto(msg *otg.PatternFlowGtpv1NPduNumberCounter) (PatternFlowGtpv1NPduNumberCounter, error)
	// FromPbText unmarshals PatternFlowGtpv1NPduNumberCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1NPduNumberCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1NPduNumberCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1NPduNumberCounter) Marshal() marshalPatternFlowGtpv1NPduNumberCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1NPduNumberCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1NPduNumberCounter) Unmarshal() unMarshalPatternFlowGtpv1NPduNumberCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1NPduNumberCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1NPduNumberCounter) ToProto() (*otg.PatternFlowGtpv1NPduNumberCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1NPduNumberCounter) FromProto(msg *otg.PatternFlowGtpv1NPduNumberCounter) (PatternFlowGtpv1NPduNumberCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1NPduNumberCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1NPduNumberCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1NPduNumberCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1NPduNumberCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1NPduNumberCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv1NPduNumberCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1NPduNumberCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpv1NPduNumberCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1NPduNumberCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1NPduNumberCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1NPduNumberCounter) Clone() (PatternFlowGtpv1NPduNumberCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1NPduNumberCounter()
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

// PatternFlowGtpv1NPduNumberCounter is integer counter pattern
type PatternFlowGtpv1NPduNumberCounter interface {
	Validation
	// msg marshals PatternFlowGtpv1NPduNumberCounter to protobuf object *otg.PatternFlowGtpv1NPduNumberCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1NPduNumberCounter
	// setMsg unmarshals PatternFlowGtpv1NPduNumberCounter from protobuf object *otg.PatternFlowGtpv1NPduNumberCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1NPduNumberCounter) PatternFlowGtpv1NPduNumberCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1NPduNumberCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1NPduNumberCounter
	// validate validates PatternFlowGtpv1NPduNumberCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1NPduNumberCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpv1NPduNumberCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpv1NPduNumberCounter
	SetStart(value uint32) PatternFlowGtpv1NPduNumberCounter
	// HasStart checks if Start has been set in PatternFlowGtpv1NPduNumberCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpv1NPduNumberCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpv1NPduNumberCounter
	SetStep(value uint32) PatternFlowGtpv1NPduNumberCounter
	// HasStep checks if Step has been set in PatternFlowGtpv1NPduNumberCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpv1NPduNumberCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpv1NPduNumberCounter
	SetCount(value uint32) PatternFlowGtpv1NPduNumberCounter
	// HasCount checks if Count has been set in PatternFlowGtpv1NPduNumberCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1NPduNumberCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1NPduNumberCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpv1NPduNumberCounter object
func (obj *patternFlowGtpv1NPduNumberCounter) SetStart(value uint32) PatternFlowGtpv1NPduNumberCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1NPduNumberCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1NPduNumberCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpv1NPduNumberCounter object
func (obj *patternFlowGtpv1NPduNumberCounter) SetStep(value uint32) PatternFlowGtpv1NPduNumberCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1NPduNumberCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1NPduNumberCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpv1NPduNumberCounter object
func (obj *patternFlowGtpv1NPduNumberCounter) SetCount(value uint32) PatternFlowGtpv1NPduNumberCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpv1NPduNumberCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1NPduNumberCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1NPduNumberCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1NPduNumberCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGtpv1NPduNumberCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(0)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

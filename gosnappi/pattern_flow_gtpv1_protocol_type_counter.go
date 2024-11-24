package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1ProtocolTypeCounter *****
type patternFlowGtpv1ProtocolTypeCounter struct {
	validation
	obj          *otg.PatternFlowGtpv1ProtocolTypeCounter
	marshaller   marshalPatternFlowGtpv1ProtocolTypeCounter
	unMarshaller unMarshalPatternFlowGtpv1ProtocolTypeCounter
}

func NewPatternFlowGtpv1ProtocolTypeCounter() PatternFlowGtpv1ProtocolTypeCounter {
	obj := patternFlowGtpv1ProtocolTypeCounter{obj: &otg.PatternFlowGtpv1ProtocolTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1ProtocolTypeCounter) msg() *otg.PatternFlowGtpv1ProtocolTypeCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1ProtocolTypeCounter) setMsg(msg *otg.PatternFlowGtpv1ProtocolTypeCounter) PatternFlowGtpv1ProtocolTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1ProtocolTypeCounter struct {
	obj *patternFlowGtpv1ProtocolTypeCounter
}

type marshalPatternFlowGtpv1ProtocolTypeCounter interface {
	// ToProto marshals PatternFlowGtpv1ProtocolTypeCounter to protobuf object *otg.PatternFlowGtpv1ProtocolTypeCounter
	ToProto() (*otg.PatternFlowGtpv1ProtocolTypeCounter, error)
	// ToPbText marshals PatternFlowGtpv1ProtocolTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1ProtocolTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1ProtocolTypeCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv1ProtocolTypeCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv1ProtocolTypeCounter struct {
	obj *patternFlowGtpv1ProtocolTypeCounter
}

type unMarshalPatternFlowGtpv1ProtocolTypeCounter interface {
	// FromProto unmarshals PatternFlowGtpv1ProtocolTypeCounter from protobuf object *otg.PatternFlowGtpv1ProtocolTypeCounter
	FromProto(msg *otg.PatternFlowGtpv1ProtocolTypeCounter) (PatternFlowGtpv1ProtocolTypeCounter, error)
	// FromPbText unmarshals PatternFlowGtpv1ProtocolTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1ProtocolTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1ProtocolTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1ProtocolTypeCounter) Marshal() marshalPatternFlowGtpv1ProtocolTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1ProtocolTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1ProtocolTypeCounter) Unmarshal() unMarshalPatternFlowGtpv1ProtocolTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1ProtocolTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1ProtocolTypeCounter) ToProto() (*otg.PatternFlowGtpv1ProtocolTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1ProtocolTypeCounter) FromProto(msg *otg.PatternFlowGtpv1ProtocolTypeCounter) (PatternFlowGtpv1ProtocolTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1ProtocolTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1ProtocolTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1ProtocolTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1ProtocolTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1ProtocolTypeCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv1ProtocolTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1ProtocolTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpv1ProtocolTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1ProtocolTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1ProtocolTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1ProtocolTypeCounter) Clone() (PatternFlowGtpv1ProtocolTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1ProtocolTypeCounter()
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

// PatternFlowGtpv1ProtocolTypeCounter is integer counter pattern
type PatternFlowGtpv1ProtocolTypeCounter interface {
	Validation
	// msg marshals PatternFlowGtpv1ProtocolTypeCounter to protobuf object *otg.PatternFlowGtpv1ProtocolTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1ProtocolTypeCounter
	// setMsg unmarshals PatternFlowGtpv1ProtocolTypeCounter from protobuf object *otg.PatternFlowGtpv1ProtocolTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1ProtocolTypeCounter) PatternFlowGtpv1ProtocolTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1ProtocolTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1ProtocolTypeCounter
	// validate validates PatternFlowGtpv1ProtocolTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1ProtocolTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpv1ProtocolTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpv1ProtocolTypeCounter
	SetStart(value uint32) PatternFlowGtpv1ProtocolTypeCounter
	// HasStart checks if Start has been set in PatternFlowGtpv1ProtocolTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpv1ProtocolTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpv1ProtocolTypeCounter
	SetStep(value uint32) PatternFlowGtpv1ProtocolTypeCounter
	// HasStep checks if Step has been set in PatternFlowGtpv1ProtocolTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpv1ProtocolTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpv1ProtocolTypeCounter
	SetCount(value uint32) PatternFlowGtpv1ProtocolTypeCounter
	// HasCount checks if Count has been set in PatternFlowGtpv1ProtocolTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1ProtocolTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1ProtocolTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpv1ProtocolTypeCounter object
func (obj *patternFlowGtpv1ProtocolTypeCounter) SetStart(value uint32) PatternFlowGtpv1ProtocolTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1ProtocolTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1ProtocolTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpv1ProtocolTypeCounter object
func (obj *patternFlowGtpv1ProtocolTypeCounter) SetStep(value uint32) PatternFlowGtpv1ProtocolTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1ProtocolTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1ProtocolTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpv1ProtocolTypeCounter object
func (obj *patternFlowGtpv1ProtocolTypeCounter) SetCount(value uint32) PatternFlowGtpv1ProtocolTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpv1ProtocolTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1ProtocolTypeCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1ProtocolTypeCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1ProtocolTypeCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGtpv1ProtocolTypeCounter) setDefault() {
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

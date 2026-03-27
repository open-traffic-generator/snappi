package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter *****
type patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	marshaller   marshalPatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	unMarshaller unMarshalPatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
}

func NewPatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter() PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter {
	obj := patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter{obj: &otg.PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) msg() *otg.PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter struct {
	obj *patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
}

type marshalPatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter struct {
	obj *patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) (PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) (PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) Clone() (PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter()
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

// PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter is integer counter pattern
type PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	// validate validates PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) setDefault() {
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

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpExtensionNextExtensionHeaderCounter *****
type patternFlowGtpExtensionNextExtensionHeaderCounter struct {
	validation
	obj          *otg.PatternFlowGtpExtensionNextExtensionHeaderCounter
	marshaller   marshalPatternFlowGtpExtensionNextExtensionHeaderCounter
	unMarshaller unMarshalPatternFlowGtpExtensionNextExtensionHeaderCounter
}

func NewPatternFlowGtpExtensionNextExtensionHeaderCounter() PatternFlowGtpExtensionNextExtensionHeaderCounter {
	obj := patternFlowGtpExtensionNextExtensionHeaderCounter{obj: &otg.PatternFlowGtpExtensionNextExtensionHeaderCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) msg() *otg.PatternFlowGtpExtensionNextExtensionHeaderCounter {
	return obj.obj
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) setMsg(msg *otg.PatternFlowGtpExtensionNextExtensionHeaderCounter) PatternFlowGtpExtensionNextExtensionHeaderCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpExtensionNextExtensionHeaderCounter struct {
	obj *patternFlowGtpExtensionNextExtensionHeaderCounter
}

type marshalPatternFlowGtpExtensionNextExtensionHeaderCounter interface {
	// ToProto marshals PatternFlowGtpExtensionNextExtensionHeaderCounter to protobuf object *otg.PatternFlowGtpExtensionNextExtensionHeaderCounter
	ToProto() (*otg.PatternFlowGtpExtensionNextExtensionHeaderCounter, error)
	// ToPbText marshals PatternFlowGtpExtensionNextExtensionHeaderCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpExtensionNextExtensionHeaderCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpExtensionNextExtensionHeaderCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpExtensionNextExtensionHeaderCounter struct {
	obj *patternFlowGtpExtensionNextExtensionHeaderCounter
}

type unMarshalPatternFlowGtpExtensionNextExtensionHeaderCounter interface {
	// FromProto unmarshals PatternFlowGtpExtensionNextExtensionHeaderCounter from protobuf object *otg.PatternFlowGtpExtensionNextExtensionHeaderCounter
	FromProto(msg *otg.PatternFlowGtpExtensionNextExtensionHeaderCounter) (PatternFlowGtpExtensionNextExtensionHeaderCounter, error)
	// FromPbText unmarshals PatternFlowGtpExtensionNextExtensionHeaderCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpExtensionNextExtensionHeaderCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpExtensionNextExtensionHeaderCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) Marshal() marshalPatternFlowGtpExtensionNextExtensionHeaderCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpExtensionNextExtensionHeaderCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) Unmarshal() unMarshalPatternFlowGtpExtensionNextExtensionHeaderCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpExtensionNextExtensionHeaderCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpExtensionNextExtensionHeaderCounter) ToProto() (*otg.PatternFlowGtpExtensionNextExtensionHeaderCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpExtensionNextExtensionHeaderCounter) FromProto(msg *otg.PatternFlowGtpExtensionNextExtensionHeaderCounter) (PatternFlowGtpExtensionNextExtensionHeaderCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpExtensionNextExtensionHeaderCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionNextExtensionHeaderCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpExtensionNextExtensionHeaderCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionNextExtensionHeaderCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpExtensionNextExtensionHeaderCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionNextExtensionHeaderCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) Clone() (PatternFlowGtpExtensionNextExtensionHeaderCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpExtensionNextExtensionHeaderCounter()
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

// PatternFlowGtpExtensionNextExtensionHeaderCounter is integer counter pattern
type PatternFlowGtpExtensionNextExtensionHeaderCounter interface {
	Validation
	// msg marshals PatternFlowGtpExtensionNextExtensionHeaderCounter to protobuf object *otg.PatternFlowGtpExtensionNextExtensionHeaderCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpExtensionNextExtensionHeaderCounter
	// setMsg unmarshals PatternFlowGtpExtensionNextExtensionHeaderCounter from protobuf object *otg.PatternFlowGtpExtensionNextExtensionHeaderCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpExtensionNextExtensionHeaderCounter) PatternFlowGtpExtensionNextExtensionHeaderCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpExtensionNextExtensionHeaderCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpExtensionNextExtensionHeaderCounter
	// validate validates PatternFlowGtpExtensionNextExtensionHeaderCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpExtensionNextExtensionHeaderCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpExtensionNextExtensionHeaderCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpExtensionNextExtensionHeaderCounter
	SetStart(value uint32) PatternFlowGtpExtensionNextExtensionHeaderCounter
	// HasStart checks if Start has been set in PatternFlowGtpExtensionNextExtensionHeaderCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpExtensionNextExtensionHeaderCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpExtensionNextExtensionHeaderCounter
	SetStep(value uint32) PatternFlowGtpExtensionNextExtensionHeaderCounter
	// HasStep checks if Step has been set in PatternFlowGtpExtensionNextExtensionHeaderCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpExtensionNextExtensionHeaderCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpExtensionNextExtensionHeaderCounter
	SetCount(value uint32) PatternFlowGtpExtensionNextExtensionHeaderCounter
	// HasCount checks if Count has been set in PatternFlowGtpExtensionNextExtensionHeaderCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpExtensionNextExtensionHeaderCounter object
func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) SetStart(value uint32) PatternFlowGtpExtensionNextExtensionHeaderCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpExtensionNextExtensionHeaderCounter object
func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) SetStep(value uint32) PatternFlowGtpExtensionNextExtensionHeaderCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpExtensionNextExtensionHeaderCounter object
func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) SetCount(value uint32) PatternFlowGtpExtensionNextExtensionHeaderCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionNextExtensionHeaderCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionNextExtensionHeaderCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionNextExtensionHeaderCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) setDefault() {
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

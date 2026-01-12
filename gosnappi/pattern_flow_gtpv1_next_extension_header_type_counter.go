package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1NextExtensionHeaderTypeCounter *****
type patternFlowGtpv1NextExtensionHeaderTypeCounter struct {
	validation
	obj          *otg.PatternFlowGtpv1NextExtensionHeaderTypeCounter
	marshaller   marshalPatternFlowGtpv1NextExtensionHeaderTypeCounter
	unMarshaller unMarshalPatternFlowGtpv1NextExtensionHeaderTypeCounter
}

func NewPatternFlowGtpv1NextExtensionHeaderTypeCounter() PatternFlowGtpv1NextExtensionHeaderTypeCounter {
	obj := patternFlowGtpv1NextExtensionHeaderTypeCounter{obj: &otg.PatternFlowGtpv1NextExtensionHeaderTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) msg() *otg.PatternFlowGtpv1NextExtensionHeaderTypeCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) setMsg(msg *otg.PatternFlowGtpv1NextExtensionHeaderTypeCounter) PatternFlowGtpv1NextExtensionHeaderTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1NextExtensionHeaderTypeCounter struct {
	obj *patternFlowGtpv1NextExtensionHeaderTypeCounter
}

type marshalPatternFlowGtpv1NextExtensionHeaderTypeCounter interface {
	// ToProto marshals PatternFlowGtpv1NextExtensionHeaderTypeCounter to protobuf object *otg.PatternFlowGtpv1NextExtensionHeaderTypeCounter
	ToProto() (*otg.PatternFlowGtpv1NextExtensionHeaderTypeCounter, error)
	// ToPbText marshals PatternFlowGtpv1NextExtensionHeaderTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1NextExtensionHeaderTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1NextExtensionHeaderTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1NextExtensionHeaderTypeCounter struct {
	obj *patternFlowGtpv1NextExtensionHeaderTypeCounter
}

type unMarshalPatternFlowGtpv1NextExtensionHeaderTypeCounter interface {
	// FromProto unmarshals PatternFlowGtpv1NextExtensionHeaderTypeCounter from protobuf object *otg.PatternFlowGtpv1NextExtensionHeaderTypeCounter
	FromProto(msg *otg.PatternFlowGtpv1NextExtensionHeaderTypeCounter) (PatternFlowGtpv1NextExtensionHeaderTypeCounter, error)
	// FromPbText unmarshals PatternFlowGtpv1NextExtensionHeaderTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1NextExtensionHeaderTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1NextExtensionHeaderTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) Marshal() marshalPatternFlowGtpv1NextExtensionHeaderTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1NextExtensionHeaderTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) Unmarshal() unMarshalPatternFlowGtpv1NextExtensionHeaderTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1NextExtensionHeaderTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1NextExtensionHeaderTypeCounter) ToProto() (*otg.PatternFlowGtpv1NextExtensionHeaderTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1NextExtensionHeaderTypeCounter) FromProto(msg *otg.PatternFlowGtpv1NextExtensionHeaderTypeCounter) (PatternFlowGtpv1NextExtensionHeaderTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1NextExtensionHeaderTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1NextExtensionHeaderTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1NextExtensionHeaderTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1NextExtensionHeaderTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1NextExtensionHeaderTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1NextExtensionHeaderTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) Clone() (PatternFlowGtpv1NextExtensionHeaderTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1NextExtensionHeaderTypeCounter()
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

// PatternFlowGtpv1NextExtensionHeaderTypeCounter is integer counter pattern
type PatternFlowGtpv1NextExtensionHeaderTypeCounter interface {
	Validation
	// msg marshals PatternFlowGtpv1NextExtensionHeaderTypeCounter to protobuf object *otg.PatternFlowGtpv1NextExtensionHeaderTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1NextExtensionHeaderTypeCounter
	// setMsg unmarshals PatternFlowGtpv1NextExtensionHeaderTypeCounter from protobuf object *otg.PatternFlowGtpv1NextExtensionHeaderTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1NextExtensionHeaderTypeCounter) PatternFlowGtpv1NextExtensionHeaderTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1NextExtensionHeaderTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1NextExtensionHeaderTypeCounter
	// validate validates PatternFlowGtpv1NextExtensionHeaderTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1NextExtensionHeaderTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpv1NextExtensionHeaderTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpv1NextExtensionHeaderTypeCounter
	SetStart(value uint32) PatternFlowGtpv1NextExtensionHeaderTypeCounter
	// HasStart checks if Start has been set in PatternFlowGtpv1NextExtensionHeaderTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpv1NextExtensionHeaderTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpv1NextExtensionHeaderTypeCounter
	SetStep(value uint32) PatternFlowGtpv1NextExtensionHeaderTypeCounter
	// HasStep checks if Step has been set in PatternFlowGtpv1NextExtensionHeaderTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpv1NextExtensionHeaderTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpv1NextExtensionHeaderTypeCounter
	SetCount(value uint32) PatternFlowGtpv1NextExtensionHeaderTypeCounter
	// HasCount checks if Count has been set in PatternFlowGtpv1NextExtensionHeaderTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpv1NextExtensionHeaderTypeCounter object
func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) SetStart(value uint32) PatternFlowGtpv1NextExtensionHeaderTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpv1NextExtensionHeaderTypeCounter object
func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) SetStep(value uint32) PatternFlowGtpv1NextExtensionHeaderTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpv1NextExtensionHeaderTypeCounter object
func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) SetCount(value uint32) PatternFlowGtpv1NextExtensionHeaderTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1NextExtensionHeaderTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1NextExtensionHeaderTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1NextExtensionHeaderTypeCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) setDefault() {
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

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVxlanReserved1Counter *****
type patternFlowVxlanReserved1Counter struct {
	validation
	obj          *otg.PatternFlowVxlanReserved1Counter
	marshaller   marshalPatternFlowVxlanReserved1Counter
	unMarshaller unMarshalPatternFlowVxlanReserved1Counter
}

func NewPatternFlowVxlanReserved1Counter() PatternFlowVxlanReserved1Counter {
	obj := patternFlowVxlanReserved1Counter{obj: &otg.PatternFlowVxlanReserved1Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVxlanReserved1Counter) msg() *otg.PatternFlowVxlanReserved1Counter {
	return obj.obj
}

func (obj *patternFlowVxlanReserved1Counter) setMsg(msg *otg.PatternFlowVxlanReserved1Counter) PatternFlowVxlanReserved1Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVxlanReserved1Counter struct {
	obj *patternFlowVxlanReserved1Counter
}

type marshalPatternFlowVxlanReserved1Counter interface {
	// ToProto marshals PatternFlowVxlanReserved1Counter to protobuf object *otg.PatternFlowVxlanReserved1Counter
	ToProto() (*otg.PatternFlowVxlanReserved1Counter, error)
	// ToPbText marshals PatternFlowVxlanReserved1Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVxlanReserved1Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVxlanReserved1Counter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowVxlanReserved1Counter struct {
	obj *patternFlowVxlanReserved1Counter
}

type unMarshalPatternFlowVxlanReserved1Counter interface {
	// FromProto unmarshals PatternFlowVxlanReserved1Counter from protobuf object *otg.PatternFlowVxlanReserved1Counter
	FromProto(msg *otg.PatternFlowVxlanReserved1Counter) (PatternFlowVxlanReserved1Counter, error)
	// FromPbText unmarshals PatternFlowVxlanReserved1Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVxlanReserved1Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVxlanReserved1Counter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVxlanReserved1Counter) Marshal() marshalPatternFlowVxlanReserved1Counter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVxlanReserved1Counter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVxlanReserved1Counter) Unmarshal() unMarshalPatternFlowVxlanReserved1Counter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVxlanReserved1Counter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVxlanReserved1Counter) ToProto() (*otg.PatternFlowVxlanReserved1Counter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVxlanReserved1Counter) FromProto(msg *otg.PatternFlowVxlanReserved1Counter) (PatternFlowVxlanReserved1Counter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVxlanReserved1Counter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVxlanReserved1Counter) FromPbText(value string) error {
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

func (m *marshalpatternFlowVxlanReserved1Counter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVxlanReserved1Counter) FromYaml(value string) error {
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

func (m *marshalpatternFlowVxlanReserved1Counter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVxlanReserved1Counter) FromJson(value string) error {
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

func (obj *patternFlowVxlanReserved1Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVxlanReserved1Counter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVxlanReserved1Counter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVxlanReserved1Counter) Clone() (PatternFlowVxlanReserved1Counter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVxlanReserved1Counter()
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

// PatternFlowVxlanReserved1Counter is integer counter pattern
type PatternFlowVxlanReserved1Counter interface {
	Validation
	// msg marshals PatternFlowVxlanReserved1Counter to protobuf object *otg.PatternFlowVxlanReserved1Counter
	// and doesn't set defaults
	msg() *otg.PatternFlowVxlanReserved1Counter
	// setMsg unmarshals PatternFlowVxlanReserved1Counter from protobuf object *otg.PatternFlowVxlanReserved1Counter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVxlanReserved1Counter) PatternFlowVxlanReserved1Counter
	// provides marshal interface
	Marshal() marshalPatternFlowVxlanReserved1Counter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVxlanReserved1Counter
	// validate validates PatternFlowVxlanReserved1Counter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVxlanReserved1Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowVxlanReserved1Counter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowVxlanReserved1Counter
	SetStart(value uint32) PatternFlowVxlanReserved1Counter
	// HasStart checks if Start has been set in PatternFlowVxlanReserved1Counter
	HasStart() bool
	// Step returns uint32, set in PatternFlowVxlanReserved1Counter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowVxlanReserved1Counter
	SetStep(value uint32) PatternFlowVxlanReserved1Counter
	// HasStep checks if Step has been set in PatternFlowVxlanReserved1Counter
	HasStep() bool
	// Count returns uint32, set in PatternFlowVxlanReserved1Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowVxlanReserved1Counter
	SetCount(value uint32) PatternFlowVxlanReserved1Counter
	// HasCount checks if Count has been set in PatternFlowVxlanReserved1Counter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowVxlanReserved1Counter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowVxlanReserved1Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowVxlanReserved1Counter object
func (obj *patternFlowVxlanReserved1Counter) SetStart(value uint32) PatternFlowVxlanReserved1Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowVxlanReserved1Counter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowVxlanReserved1Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowVxlanReserved1Counter object
func (obj *patternFlowVxlanReserved1Counter) SetStep(value uint32) PatternFlowVxlanReserved1Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowVxlanReserved1Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowVxlanReserved1Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowVxlanReserved1Counter object
func (obj *patternFlowVxlanReserved1Counter) SetCount(value uint32) PatternFlowVxlanReserved1Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowVxlanReserved1Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVxlanReserved1Counter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVxlanReserved1Counter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVxlanReserved1Counter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowVxlanReserved1Counter) setDefault() {
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

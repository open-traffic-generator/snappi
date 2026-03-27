package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv2Spare1Counter *****
type patternFlowGtpv2Spare1Counter struct {
	validation
	obj          *otg.PatternFlowGtpv2Spare1Counter
	marshaller   marshalPatternFlowGtpv2Spare1Counter
	unMarshaller unMarshalPatternFlowGtpv2Spare1Counter
}

func NewPatternFlowGtpv2Spare1Counter() PatternFlowGtpv2Spare1Counter {
	obj := patternFlowGtpv2Spare1Counter{obj: &otg.PatternFlowGtpv2Spare1Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2Spare1Counter) msg() *otg.PatternFlowGtpv2Spare1Counter {
	return obj.obj
}

func (obj *patternFlowGtpv2Spare1Counter) setMsg(msg *otg.PatternFlowGtpv2Spare1Counter) PatternFlowGtpv2Spare1Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2Spare1Counter struct {
	obj *patternFlowGtpv2Spare1Counter
}

type marshalPatternFlowGtpv2Spare1Counter interface {
	// ToProto marshals PatternFlowGtpv2Spare1Counter to protobuf object *otg.PatternFlowGtpv2Spare1Counter
	ToProto() (*otg.PatternFlowGtpv2Spare1Counter, error)
	// ToPbText marshals PatternFlowGtpv2Spare1Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2Spare1Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2Spare1Counter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv2Spare1Counter struct {
	obj *patternFlowGtpv2Spare1Counter
}

type unMarshalPatternFlowGtpv2Spare1Counter interface {
	// FromProto unmarshals PatternFlowGtpv2Spare1Counter from protobuf object *otg.PatternFlowGtpv2Spare1Counter
	FromProto(msg *otg.PatternFlowGtpv2Spare1Counter) (PatternFlowGtpv2Spare1Counter, error)
	// FromPbText unmarshals PatternFlowGtpv2Spare1Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2Spare1Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2Spare1Counter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2Spare1Counter) Marshal() marshalPatternFlowGtpv2Spare1Counter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2Spare1Counter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2Spare1Counter) Unmarshal() unMarshalPatternFlowGtpv2Spare1Counter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2Spare1Counter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2Spare1Counter) ToProto() (*otg.PatternFlowGtpv2Spare1Counter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2Spare1Counter) FromProto(msg *otg.PatternFlowGtpv2Spare1Counter) (PatternFlowGtpv2Spare1Counter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2Spare1Counter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2Spare1Counter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2Spare1Counter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2Spare1Counter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2Spare1Counter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2Spare1Counter) FromJson(value string) error {
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

func (obj *patternFlowGtpv2Spare1Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2Spare1Counter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2Spare1Counter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2Spare1Counter) Clone() (PatternFlowGtpv2Spare1Counter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2Spare1Counter()
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

// PatternFlowGtpv2Spare1Counter is integer counter pattern
type PatternFlowGtpv2Spare1Counter interface {
	Validation
	// msg marshals PatternFlowGtpv2Spare1Counter to protobuf object *otg.PatternFlowGtpv2Spare1Counter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2Spare1Counter
	// setMsg unmarshals PatternFlowGtpv2Spare1Counter from protobuf object *otg.PatternFlowGtpv2Spare1Counter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2Spare1Counter) PatternFlowGtpv2Spare1Counter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2Spare1Counter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2Spare1Counter
	// validate validates PatternFlowGtpv2Spare1Counter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2Spare1Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpv2Spare1Counter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpv2Spare1Counter
	SetStart(value uint32) PatternFlowGtpv2Spare1Counter
	// HasStart checks if Start has been set in PatternFlowGtpv2Spare1Counter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpv2Spare1Counter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpv2Spare1Counter
	SetStep(value uint32) PatternFlowGtpv2Spare1Counter
	// HasStep checks if Step has been set in PatternFlowGtpv2Spare1Counter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpv2Spare1Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpv2Spare1Counter
	SetCount(value uint32) PatternFlowGtpv2Spare1Counter
	// HasCount checks if Count has been set in PatternFlowGtpv2Spare1Counter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv2Spare1Counter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv2Spare1Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpv2Spare1Counter object
func (obj *patternFlowGtpv2Spare1Counter) SetStart(value uint32) PatternFlowGtpv2Spare1Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv2Spare1Counter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv2Spare1Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpv2Spare1Counter object
func (obj *patternFlowGtpv2Spare1Counter) SetStep(value uint32) PatternFlowGtpv2Spare1Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv2Spare1Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv2Spare1Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpv2Spare1Counter object
func (obj *patternFlowGtpv2Spare1Counter) SetCount(value uint32) PatternFlowGtpv2Spare1Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpv2Spare1Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2Spare1Counter.Start <= 7 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2Spare1Counter.Step <= 7 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2Spare1Counter.Count <= 8 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGtpv2Spare1Counter) setDefault() {
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

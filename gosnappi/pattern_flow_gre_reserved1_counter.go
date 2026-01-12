package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGreReserved1Counter *****
type patternFlowGreReserved1Counter struct {
	validation
	obj          *otg.PatternFlowGreReserved1Counter
	marshaller   marshalPatternFlowGreReserved1Counter
	unMarshaller unMarshalPatternFlowGreReserved1Counter
}

func NewPatternFlowGreReserved1Counter() PatternFlowGreReserved1Counter {
	obj := patternFlowGreReserved1Counter{obj: &otg.PatternFlowGreReserved1Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGreReserved1Counter) msg() *otg.PatternFlowGreReserved1Counter {
	return obj.obj
}

func (obj *patternFlowGreReserved1Counter) setMsg(msg *otg.PatternFlowGreReserved1Counter) PatternFlowGreReserved1Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGreReserved1Counter struct {
	obj *patternFlowGreReserved1Counter
}

type marshalPatternFlowGreReserved1Counter interface {
	// ToProto marshals PatternFlowGreReserved1Counter to protobuf object *otg.PatternFlowGreReserved1Counter
	ToProto() (*otg.PatternFlowGreReserved1Counter, error)
	// ToPbText marshals PatternFlowGreReserved1Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGreReserved1Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGreReserved1Counter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGreReserved1Counter struct {
	obj *patternFlowGreReserved1Counter
}

type unMarshalPatternFlowGreReserved1Counter interface {
	// FromProto unmarshals PatternFlowGreReserved1Counter from protobuf object *otg.PatternFlowGreReserved1Counter
	FromProto(msg *otg.PatternFlowGreReserved1Counter) (PatternFlowGreReserved1Counter, error)
	// FromPbText unmarshals PatternFlowGreReserved1Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGreReserved1Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGreReserved1Counter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGreReserved1Counter) Marshal() marshalPatternFlowGreReserved1Counter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGreReserved1Counter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGreReserved1Counter) Unmarshal() unMarshalPatternFlowGreReserved1Counter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGreReserved1Counter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGreReserved1Counter) ToProto() (*otg.PatternFlowGreReserved1Counter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGreReserved1Counter) FromProto(msg *otg.PatternFlowGreReserved1Counter) (PatternFlowGreReserved1Counter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGreReserved1Counter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGreReserved1Counter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGreReserved1Counter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGreReserved1Counter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGreReserved1Counter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGreReserved1Counter) FromJson(value string) error {
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

func (obj *patternFlowGreReserved1Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGreReserved1Counter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGreReserved1Counter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGreReserved1Counter) Clone() (PatternFlowGreReserved1Counter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGreReserved1Counter()
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

// PatternFlowGreReserved1Counter is integer counter pattern
type PatternFlowGreReserved1Counter interface {
	Validation
	// msg marshals PatternFlowGreReserved1Counter to protobuf object *otg.PatternFlowGreReserved1Counter
	// and doesn't set defaults
	msg() *otg.PatternFlowGreReserved1Counter
	// setMsg unmarshals PatternFlowGreReserved1Counter from protobuf object *otg.PatternFlowGreReserved1Counter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGreReserved1Counter) PatternFlowGreReserved1Counter
	// provides marshal interface
	Marshal() marshalPatternFlowGreReserved1Counter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGreReserved1Counter
	// validate validates PatternFlowGreReserved1Counter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGreReserved1Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGreReserved1Counter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGreReserved1Counter
	SetStart(value uint32) PatternFlowGreReserved1Counter
	// HasStart checks if Start has been set in PatternFlowGreReserved1Counter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGreReserved1Counter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGreReserved1Counter
	SetStep(value uint32) PatternFlowGreReserved1Counter
	// HasStep checks if Step has been set in PatternFlowGreReserved1Counter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGreReserved1Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGreReserved1Counter
	SetCount(value uint32) PatternFlowGreReserved1Counter
	// HasCount checks if Count has been set in PatternFlowGreReserved1Counter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGreReserved1Counter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGreReserved1Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGreReserved1Counter object
func (obj *patternFlowGreReserved1Counter) SetStart(value uint32) PatternFlowGreReserved1Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGreReserved1Counter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGreReserved1Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGreReserved1Counter object
func (obj *patternFlowGreReserved1Counter) SetStep(value uint32) PatternFlowGreReserved1Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGreReserved1Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGreReserved1Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGreReserved1Counter object
func (obj *patternFlowGreReserved1Counter) SetCount(value uint32) PatternFlowGreReserved1Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGreReserved1Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreReserved1Counter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreReserved1Counter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreReserved1Counter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGreReserved1Counter) setDefault() {
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

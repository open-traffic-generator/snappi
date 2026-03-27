package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGreReserved0Counter *****
type patternFlowGreReserved0Counter struct {
	validation
	obj          *otg.PatternFlowGreReserved0Counter
	marshaller   marshalPatternFlowGreReserved0Counter
	unMarshaller unMarshalPatternFlowGreReserved0Counter
}

func NewPatternFlowGreReserved0Counter() PatternFlowGreReserved0Counter {
	obj := patternFlowGreReserved0Counter{obj: &otg.PatternFlowGreReserved0Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGreReserved0Counter) msg() *otg.PatternFlowGreReserved0Counter {
	return obj.obj
}

func (obj *patternFlowGreReserved0Counter) setMsg(msg *otg.PatternFlowGreReserved0Counter) PatternFlowGreReserved0Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGreReserved0Counter struct {
	obj *patternFlowGreReserved0Counter
}

type marshalPatternFlowGreReserved0Counter interface {
	// ToProto marshals PatternFlowGreReserved0Counter to protobuf object *otg.PatternFlowGreReserved0Counter
	ToProto() (*otg.PatternFlowGreReserved0Counter, error)
	// ToPbText marshals PatternFlowGreReserved0Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGreReserved0Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGreReserved0Counter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGreReserved0Counter struct {
	obj *patternFlowGreReserved0Counter
}

type unMarshalPatternFlowGreReserved0Counter interface {
	// FromProto unmarshals PatternFlowGreReserved0Counter from protobuf object *otg.PatternFlowGreReserved0Counter
	FromProto(msg *otg.PatternFlowGreReserved0Counter) (PatternFlowGreReserved0Counter, error)
	// FromPbText unmarshals PatternFlowGreReserved0Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGreReserved0Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGreReserved0Counter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGreReserved0Counter) Marshal() marshalPatternFlowGreReserved0Counter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGreReserved0Counter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGreReserved0Counter) Unmarshal() unMarshalPatternFlowGreReserved0Counter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGreReserved0Counter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGreReserved0Counter) ToProto() (*otg.PatternFlowGreReserved0Counter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGreReserved0Counter) FromProto(msg *otg.PatternFlowGreReserved0Counter) (PatternFlowGreReserved0Counter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGreReserved0Counter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGreReserved0Counter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGreReserved0Counter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGreReserved0Counter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGreReserved0Counter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGreReserved0Counter) FromJson(value string) error {
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

func (obj *patternFlowGreReserved0Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGreReserved0Counter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGreReserved0Counter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGreReserved0Counter) Clone() (PatternFlowGreReserved0Counter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGreReserved0Counter()
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

// PatternFlowGreReserved0Counter is integer counter pattern
type PatternFlowGreReserved0Counter interface {
	Validation
	// msg marshals PatternFlowGreReserved0Counter to protobuf object *otg.PatternFlowGreReserved0Counter
	// and doesn't set defaults
	msg() *otg.PatternFlowGreReserved0Counter
	// setMsg unmarshals PatternFlowGreReserved0Counter from protobuf object *otg.PatternFlowGreReserved0Counter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGreReserved0Counter) PatternFlowGreReserved0Counter
	// provides marshal interface
	Marshal() marshalPatternFlowGreReserved0Counter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGreReserved0Counter
	// validate validates PatternFlowGreReserved0Counter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGreReserved0Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGreReserved0Counter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGreReserved0Counter
	SetStart(value uint32) PatternFlowGreReserved0Counter
	// HasStart checks if Start has been set in PatternFlowGreReserved0Counter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGreReserved0Counter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGreReserved0Counter
	SetStep(value uint32) PatternFlowGreReserved0Counter
	// HasStep checks if Step has been set in PatternFlowGreReserved0Counter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGreReserved0Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGreReserved0Counter
	SetCount(value uint32) PatternFlowGreReserved0Counter
	// HasCount checks if Count has been set in PatternFlowGreReserved0Counter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGreReserved0Counter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGreReserved0Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGreReserved0Counter object
func (obj *patternFlowGreReserved0Counter) SetStart(value uint32) PatternFlowGreReserved0Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGreReserved0Counter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGreReserved0Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGreReserved0Counter object
func (obj *patternFlowGreReserved0Counter) SetStep(value uint32) PatternFlowGreReserved0Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGreReserved0Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGreReserved0Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGreReserved0Counter object
func (obj *patternFlowGreReserved0Counter) SetCount(value uint32) PatternFlowGreReserved0Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGreReserved0Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 4095 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreReserved0Counter.Start <= 4095 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 4095 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreReserved0Counter.Step <= 4095 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 4096 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreReserved0Counter.Count <= 4096 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGreReserved0Counter) setDefault() {
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

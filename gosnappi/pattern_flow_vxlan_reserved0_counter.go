package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVxlanReserved0Counter *****
type patternFlowVxlanReserved0Counter struct {
	validation
	obj          *otg.PatternFlowVxlanReserved0Counter
	marshaller   marshalPatternFlowVxlanReserved0Counter
	unMarshaller unMarshalPatternFlowVxlanReserved0Counter
}

func NewPatternFlowVxlanReserved0Counter() PatternFlowVxlanReserved0Counter {
	obj := patternFlowVxlanReserved0Counter{obj: &otg.PatternFlowVxlanReserved0Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVxlanReserved0Counter) msg() *otg.PatternFlowVxlanReserved0Counter {
	return obj.obj
}

func (obj *patternFlowVxlanReserved0Counter) setMsg(msg *otg.PatternFlowVxlanReserved0Counter) PatternFlowVxlanReserved0Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVxlanReserved0Counter struct {
	obj *patternFlowVxlanReserved0Counter
}

type marshalPatternFlowVxlanReserved0Counter interface {
	// ToProto marshals PatternFlowVxlanReserved0Counter to protobuf object *otg.PatternFlowVxlanReserved0Counter
	ToProto() (*otg.PatternFlowVxlanReserved0Counter, error)
	// ToPbText marshals PatternFlowVxlanReserved0Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVxlanReserved0Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVxlanReserved0Counter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowVxlanReserved0Counter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowVxlanReserved0Counter struct {
	obj *patternFlowVxlanReserved0Counter
}

type unMarshalPatternFlowVxlanReserved0Counter interface {
	// FromProto unmarshals PatternFlowVxlanReserved0Counter from protobuf object *otg.PatternFlowVxlanReserved0Counter
	FromProto(msg *otg.PatternFlowVxlanReserved0Counter) (PatternFlowVxlanReserved0Counter, error)
	// FromPbText unmarshals PatternFlowVxlanReserved0Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVxlanReserved0Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVxlanReserved0Counter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVxlanReserved0Counter) Marshal() marshalPatternFlowVxlanReserved0Counter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVxlanReserved0Counter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVxlanReserved0Counter) Unmarshal() unMarshalPatternFlowVxlanReserved0Counter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVxlanReserved0Counter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVxlanReserved0Counter) ToProto() (*otg.PatternFlowVxlanReserved0Counter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVxlanReserved0Counter) FromProto(msg *otg.PatternFlowVxlanReserved0Counter) (PatternFlowVxlanReserved0Counter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVxlanReserved0Counter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVxlanReserved0Counter) FromPbText(value string) error {
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

func (m *marshalpatternFlowVxlanReserved0Counter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVxlanReserved0Counter) FromYaml(value string) error {
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

func (m *marshalpatternFlowVxlanReserved0Counter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowVxlanReserved0Counter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVxlanReserved0Counter) FromJson(value string) error {
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

func (obj *patternFlowVxlanReserved0Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVxlanReserved0Counter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVxlanReserved0Counter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVxlanReserved0Counter) Clone() (PatternFlowVxlanReserved0Counter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVxlanReserved0Counter()
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

// PatternFlowVxlanReserved0Counter is integer counter pattern
type PatternFlowVxlanReserved0Counter interface {
	Validation
	// msg marshals PatternFlowVxlanReserved0Counter to protobuf object *otg.PatternFlowVxlanReserved0Counter
	// and doesn't set defaults
	msg() *otg.PatternFlowVxlanReserved0Counter
	// setMsg unmarshals PatternFlowVxlanReserved0Counter from protobuf object *otg.PatternFlowVxlanReserved0Counter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVxlanReserved0Counter) PatternFlowVxlanReserved0Counter
	// provides marshal interface
	Marshal() marshalPatternFlowVxlanReserved0Counter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVxlanReserved0Counter
	// validate validates PatternFlowVxlanReserved0Counter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVxlanReserved0Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowVxlanReserved0Counter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowVxlanReserved0Counter
	SetStart(value uint32) PatternFlowVxlanReserved0Counter
	// HasStart checks if Start has been set in PatternFlowVxlanReserved0Counter
	HasStart() bool
	// Step returns uint32, set in PatternFlowVxlanReserved0Counter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowVxlanReserved0Counter
	SetStep(value uint32) PatternFlowVxlanReserved0Counter
	// HasStep checks if Step has been set in PatternFlowVxlanReserved0Counter
	HasStep() bool
	// Count returns uint32, set in PatternFlowVxlanReserved0Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowVxlanReserved0Counter
	SetCount(value uint32) PatternFlowVxlanReserved0Counter
	// HasCount checks if Count has been set in PatternFlowVxlanReserved0Counter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowVxlanReserved0Counter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowVxlanReserved0Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowVxlanReserved0Counter object
func (obj *patternFlowVxlanReserved0Counter) SetStart(value uint32) PatternFlowVxlanReserved0Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowVxlanReserved0Counter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowVxlanReserved0Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowVxlanReserved0Counter object
func (obj *patternFlowVxlanReserved0Counter) SetStep(value uint32) PatternFlowVxlanReserved0Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowVxlanReserved0Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowVxlanReserved0Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowVxlanReserved0Counter object
func (obj *patternFlowVxlanReserved0Counter) SetCount(value uint32) PatternFlowVxlanReserved0Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowVxlanReserved0Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVxlanReserved0Counter.Start <= 16777215 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVxlanReserved0Counter.Step <= 16777215 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVxlanReserved0Counter.Count <= 16777215 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowVxlanReserved0Counter) setDefault() {
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

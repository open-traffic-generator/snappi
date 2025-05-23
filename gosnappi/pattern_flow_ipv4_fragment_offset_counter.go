package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4FragmentOffsetCounter *****
type patternFlowIpv4FragmentOffsetCounter struct {
	validation
	obj          *otg.PatternFlowIpv4FragmentOffsetCounter
	marshaller   marshalPatternFlowIpv4FragmentOffsetCounter
	unMarshaller unMarshalPatternFlowIpv4FragmentOffsetCounter
}

func NewPatternFlowIpv4FragmentOffsetCounter() PatternFlowIpv4FragmentOffsetCounter {
	obj := patternFlowIpv4FragmentOffsetCounter{obj: &otg.PatternFlowIpv4FragmentOffsetCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4FragmentOffsetCounter) msg() *otg.PatternFlowIpv4FragmentOffsetCounter {
	return obj.obj
}

func (obj *patternFlowIpv4FragmentOffsetCounter) setMsg(msg *otg.PatternFlowIpv4FragmentOffsetCounter) PatternFlowIpv4FragmentOffsetCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4FragmentOffsetCounter struct {
	obj *patternFlowIpv4FragmentOffsetCounter
}

type marshalPatternFlowIpv4FragmentOffsetCounter interface {
	// ToProto marshals PatternFlowIpv4FragmentOffsetCounter to protobuf object *otg.PatternFlowIpv4FragmentOffsetCounter
	ToProto() (*otg.PatternFlowIpv4FragmentOffsetCounter, error)
	// ToPbText marshals PatternFlowIpv4FragmentOffsetCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4FragmentOffsetCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4FragmentOffsetCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4FragmentOffsetCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4FragmentOffsetCounter struct {
	obj *patternFlowIpv4FragmentOffsetCounter
}

type unMarshalPatternFlowIpv4FragmentOffsetCounter interface {
	// FromProto unmarshals PatternFlowIpv4FragmentOffsetCounter from protobuf object *otg.PatternFlowIpv4FragmentOffsetCounter
	FromProto(msg *otg.PatternFlowIpv4FragmentOffsetCounter) (PatternFlowIpv4FragmentOffsetCounter, error)
	// FromPbText unmarshals PatternFlowIpv4FragmentOffsetCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4FragmentOffsetCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4FragmentOffsetCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4FragmentOffsetCounter) Marshal() marshalPatternFlowIpv4FragmentOffsetCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4FragmentOffsetCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4FragmentOffsetCounter) Unmarshal() unMarshalPatternFlowIpv4FragmentOffsetCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4FragmentOffsetCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4FragmentOffsetCounter) ToProto() (*otg.PatternFlowIpv4FragmentOffsetCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4FragmentOffsetCounter) FromProto(msg *otg.PatternFlowIpv4FragmentOffsetCounter) (PatternFlowIpv4FragmentOffsetCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4FragmentOffsetCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4FragmentOffsetCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4FragmentOffsetCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4FragmentOffsetCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4FragmentOffsetCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4FragmentOffsetCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4FragmentOffsetCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4FragmentOffsetCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4FragmentOffsetCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4FragmentOffsetCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4FragmentOffsetCounter) Clone() (PatternFlowIpv4FragmentOffsetCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4FragmentOffsetCounter()
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

// PatternFlowIpv4FragmentOffsetCounter is integer counter pattern
type PatternFlowIpv4FragmentOffsetCounter interface {
	Validation
	// msg marshals PatternFlowIpv4FragmentOffsetCounter to protobuf object *otg.PatternFlowIpv4FragmentOffsetCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4FragmentOffsetCounter
	// setMsg unmarshals PatternFlowIpv4FragmentOffsetCounter from protobuf object *otg.PatternFlowIpv4FragmentOffsetCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4FragmentOffsetCounter) PatternFlowIpv4FragmentOffsetCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4FragmentOffsetCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4FragmentOffsetCounter
	// validate validates PatternFlowIpv4FragmentOffsetCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4FragmentOffsetCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4FragmentOffsetCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4FragmentOffsetCounter
	SetStart(value uint32) PatternFlowIpv4FragmentOffsetCounter
	// HasStart checks if Start has been set in PatternFlowIpv4FragmentOffsetCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4FragmentOffsetCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4FragmentOffsetCounter
	SetStep(value uint32) PatternFlowIpv4FragmentOffsetCounter
	// HasStep checks if Step has been set in PatternFlowIpv4FragmentOffsetCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4FragmentOffsetCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4FragmentOffsetCounter
	SetCount(value uint32) PatternFlowIpv4FragmentOffsetCounter
	// HasCount checks if Count has been set in PatternFlowIpv4FragmentOffsetCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4FragmentOffsetCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4FragmentOffsetCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4FragmentOffsetCounter object
func (obj *patternFlowIpv4FragmentOffsetCounter) SetStart(value uint32) PatternFlowIpv4FragmentOffsetCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4FragmentOffsetCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4FragmentOffsetCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4FragmentOffsetCounter object
func (obj *patternFlowIpv4FragmentOffsetCounter) SetStep(value uint32) PatternFlowIpv4FragmentOffsetCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4FragmentOffsetCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4FragmentOffsetCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4FragmentOffsetCounter object
func (obj *patternFlowIpv4FragmentOffsetCounter) SetCount(value uint32) PatternFlowIpv4FragmentOffsetCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4FragmentOffsetCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4FragmentOffsetCounter.Start <= 31 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4FragmentOffsetCounter.Step <= 31 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4FragmentOffsetCounter.Count <= 31 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4FragmentOffsetCounter) setDefault() {
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

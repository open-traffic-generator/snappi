package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4OptionsCustomTypeOptionClassCounter *****
type patternFlowIpv4OptionsCustomTypeOptionClassCounter struct {
	validation
	obj          *otg.PatternFlowIpv4OptionsCustomTypeOptionClassCounter
	marshaller   marshalPatternFlowIpv4OptionsCustomTypeOptionClassCounter
	unMarshaller unMarshalPatternFlowIpv4OptionsCustomTypeOptionClassCounter
}

func NewPatternFlowIpv4OptionsCustomTypeOptionClassCounter() PatternFlowIpv4OptionsCustomTypeOptionClassCounter {
	obj := patternFlowIpv4OptionsCustomTypeOptionClassCounter{obj: &otg.PatternFlowIpv4OptionsCustomTypeOptionClassCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionClassCounter) msg() *otg.PatternFlowIpv4OptionsCustomTypeOptionClassCounter {
	return obj.obj
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionClassCounter) setMsg(msg *otg.PatternFlowIpv4OptionsCustomTypeOptionClassCounter) PatternFlowIpv4OptionsCustomTypeOptionClassCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4OptionsCustomTypeOptionClassCounter struct {
	obj *patternFlowIpv4OptionsCustomTypeOptionClassCounter
}

type marshalPatternFlowIpv4OptionsCustomTypeOptionClassCounter interface {
	// ToProto marshals PatternFlowIpv4OptionsCustomTypeOptionClassCounter to protobuf object *otg.PatternFlowIpv4OptionsCustomTypeOptionClassCounter
	ToProto() (*otg.PatternFlowIpv4OptionsCustomTypeOptionClassCounter, error)
	// ToPbText marshals PatternFlowIpv4OptionsCustomTypeOptionClassCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4OptionsCustomTypeOptionClassCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4OptionsCustomTypeOptionClassCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4OptionsCustomTypeOptionClassCounter struct {
	obj *patternFlowIpv4OptionsCustomTypeOptionClassCounter
}

type unMarshalPatternFlowIpv4OptionsCustomTypeOptionClassCounter interface {
	// FromProto unmarshals PatternFlowIpv4OptionsCustomTypeOptionClassCounter from protobuf object *otg.PatternFlowIpv4OptionsCustomTypeOptionClassCounter
	FromProto(msg *otg.PatternFlowIpv4OptionsCustomTypeOptionClassCounter) (PatternFlowIpv4OptionsCustomTypeOptionClassCounter, error)
	// FromPbText unmarshals PatternFlowIpv4OptionsCustomTypeOptionClassCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4OptionsCustomTypeOptionClassCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4OptionsCustomTypeOptionClassCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionClassCounter) Marshal() marshalPatternFlowIpv4OptionsCustomTypeOptionClassCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4OptionsCustomTypeOptionClassCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionClassCounter) Unmarshal() unMarshalPatternFlowIpv4OptionsCustomTypeOptionClassCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4OptionsCustomTypeOptionClassCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4OptionsCustomTypeOptionClassCounter) ToProto() (*otg.PatternFlowIpv4OptionsCustomTypeOptionClassCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeOptionClassCounter) FromProto(msg *otg.PatternFlowIpv4OptionsCustomTypeOptionClassCounter) (PatternFlowIpv4OptionsCustomTypeOptionClassCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4OptionsCustomTypeOptionClassCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeOptionClassCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsCustomTypeOptionClassCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeOptionClassCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsCustomTypeOptionClassCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeOptionClassCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4OptionsCustomTypeOptionClassCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionClassCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionClassCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionClassCounter) Clone() (PatternFlowIpv4OptionsCustomTypeOptionClassCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4OptionsCustomTypeOptionClassCounter()
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

// PatternFlowIpv4OptionsCustomTypeOptionClassCounter is integer counter pattern
type PatternFlowIpv4OptionsCustomTypeOptionClassCounter interface {
	Validation
	// msg marshals PatternFlowIpv4OptionsCustomTypeOptionClassCounter to protobuf object *otg.PatternFlowIpv4OptionsCustomTypeOptionClassCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4OptionsCustomTypeOptionClassCounter
	// setMsg unmarshals PatternFlowIpv4OptionsCustomTypeOptionClassCounter from protobuf object *otg.PatternFlowIpv4OptionsCustomTypeOptionClassCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4OptionsCustomTypeOptionClassCounter) PatternFlowIpv4OptionsCustomTypeOptionClassCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4OptionsCustomTypeOptionClassCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4OptionsCustomTypeOptionClassCounter
	// validate validates PatternFlowIpv4OptionsCustomTypeOptionClassCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4OptionsCustomTypeOptionClassCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4OptionsCustomTypeOptionClassCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4OptionsCustomTypeOptionClassCounter
	SetStart(value uint32) PatternFlowIpv4OptionsCustomTypeOptionClassCounter
	// HasStart checks if Start has been set in PatternFlowIpv4OptionsCustomTypeOptionClassCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4OptionsCustomTypeOptionClassCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4OptionsCustomTypeOptionClassCounter
	SetStep(value uint32) PatternFlowIpv4OptionsCustomTypeOptionClassCounter
	// HasStep checks if Step has been set in PatternFlowIpv4OptionsCustomTypeOptionClassCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4OptionsCustomTypeOptionClassCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4OptionsCustomTypeOptionClassCounter
	SetCount(value uint32) PatternFlowIpv4OptionsCustomTypeOptionClassCounter
	// HasCount checks if Count has been set in PatternFlowIpv4OptionsCustomTypeOptionClassCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeOptionClassCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeOptionClassCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4OptionsCustomTypeOptionClassCounter object
func (obj *patternFlowIpv4OptionsCustomTypeOptionClassCounter) SetStart(value uint32) PatternFlowIpv4OptionsCustomTypeOptionClassCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeOptionClassCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeOptionClassCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4OptionsCustomTypeOptionClassCounter object
func (obj *patternFlowIpv4OptionsCustomTypeOptionClassCounter) SetStep(value uint32) PatternFlowIpv4OptionsCustomTypeOptionClassCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeOptionClassCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeOptionClassCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4OptionsCustomTypeOptionClassCounter object
func (obj *patternFlowIpv4OptionsCustomTypeOptionClassCounter) SetCount(value uint32) PatternFlowIpv4OptionsCustomTypeOptionClassCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionClassCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4OptionsCustomTypeOptionClassCounter.Start <= 3 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4OptionsCustomTypeOptionClassCounter.Step <= 3 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 4 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4OptionsCustomTypeOptionClassCounter.Count <= 4 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4OptionsCustomTypeOptionClassCounter) setDefault() {
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

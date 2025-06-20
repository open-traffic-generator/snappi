package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4MoreFragmentsCounter *****
type patternFlowIpv4MoreFragmentsCounter struct {
	validation
	obj          *otg.PatternFlowIpv4MoreFragmentsCounter
	marshaller   marshalPatternFlowIpv4MoreFragmentsCounter
	unMarshaller unMarshalPatternFlowIpv4MoreFragmentsCounter
}

func NewPatternFlowIpv4MoreFragmentsCounter() PatternFlowIpv4MoreFragmentsCounter {
	obj := patternFlowIpv4MoreFragmentsCounter{obj: &otg.PatternFlowIpv4MoreFragmentsCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4MoreFragmentsCounter) msg() *otg.PatternFlowIpv4MoreFragmentsCounter {
	return obj.obj
}

func (obj *patternFlowIpv4MoreFragmentsCounter) setMsg(msg *otg.PatternFlowIpv4MoreFragmentsCounter) PatternFlowIpv4MoreFragmentsCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4MoreFragmentsCounter struct {
	obj *patternFlowIpv4MoreFragmentsCounter
}

type marshalPatternFlowIpv4MoreFragmentsCounter interface {
	// ToProto marshals PatternFlowIpv4MoreFragmentsCounter to protobuf object *otg.PatternFlowIpv4MoreFragmentsCounter
	ToProto() (*otg.PatternFlowIpv4MoreFragmentsCounter, error)
	// ToPbText marshals PatternFlowIpv4MoreFragmentsCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4MoreFragmentsCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4MoreFragmentsCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4MoreFragmentsCounter struct {
	obj *patternFlowIpv4MoreFragmentsCounter
}

type unMarshalPatternFlowIpv4MoreFragmentsCounter interface {
	// FromProto unmarshals PatternFlowIpv4MoreFragmentsCounter from protobuf object *otg.PatternFlowIpv4MoreFragmentsCounter
	FromProto(msg *otg.PatternFlowIpv4MoreFragmentsCounter) (PatternFlowIpv4MoreFragmentsCounter, error)
	// FromPbText unmarshals PatternFlowIpv4MoreFragmentsCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4MoreFragmentsCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4MoreFragmentsCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4MoreFragmentsCounter) Marshal() marshalPatternFlowIpv4MoreFragmentsCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4MoreFragmentsCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4MoreFragmentsCounter) Unmarshal() unMarshalPatternFlowIpv4MoreFragmentsCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4MoreFragmentsCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4MoreFragmentsCounter) ToProto() (*otg.PatternFlowIpv4MoreFragmentsCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4MoreFragmentsCounter) FromProto(msg *otg.PatternFlowIpv4MoreFragmentsCounter) (PatternFlowIpv4MoreFragmentsCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4MoreFragmentsCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4MoreFragmentsCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4MoreFragmentsCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4MoreFragmentsCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4MoreFragmentsCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4MoreFragmentsCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4MoreFragmentsCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4MoreFragmentsCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4MoreFragmentsCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4MoreFragmentsCounter) Clone() (PatternFlowIpv4MoreFragmentsCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4MoreFragmentsCounter()
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

// PatternFlowIpv4MoreFragmentsCounter is integer counter pattern
type PatternFlowIpv4MoreFragmentsCounter interface {
	Validation
	// msg marshals PatternFlowIpv4MoreFragmentsCounter to protobuf object *otg.PatternFlowIpv4MoreFragmentsCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4MoreFragmentsCounter
	// setMsg unmarshals PatternFlowIpv4MoreFragmentsCounter from protobuf object *otg.PatternFlowIpv4MoreFragmentsCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4MoreFragmentsCounter) PatternFlowIpv4MoreFragmentsCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4MoreFragmentsCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4MoreFragmentsCounter
	// validate validates PatternFlowIpv4MoreFragmentsCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4MoreFragmentsCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4MoreFragmentsCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4MoreFragmentsCounter
	SetStart(value uint32) PatternFlowIpv4MoreFragmentsCounter
	// HasStart checks if Start has been set in PatternFlowIpv4MoreFragmentsCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4MoreFragmentsCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4MoreFragmentsCounter
	SetStep(value uint32) PatternFlowIpv4MoreFragmentsCounter
	// HasStep checks if Step has been set in PatternFlowIpv4MoreFragmentsCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4MoreFragmentsCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4MoreFragmentsCounter
	SetCount(value uint32) PatternFlowIpv4MoreFragmentsCounter
	// HasCount checks if Count has been set in PatternFlowIpv4MoreFragmentsCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4MoreFragmentsCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4MoreFragmentsCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4MoreFragmentsCounter object
func (obj *patternFlowIpv4MoreFragmentsCounter) SetStart(value uint32) PatternFlowIpv4MoreFragmentsCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4MoreFragmentsCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4MoreFragmentsCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4MoreFragmentsCounter object
func (obj *patternFlowIpv4MoreFragmentsCounter) SetStep(value uint32) PatternFlowIpv4MoreFragmentsCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4MoreFragmentsCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4MoreFragmentsCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4MoreFragmentsCounter object
func (obj *patternFlowIpv4MoreFragmentsCounter) SetCount(value uint32) PatternFlowIpv4MoreFragmentsCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4MoreFragmentsCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4MoreFragmentsCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4MoreFragmentsCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4MoreFragmentsCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4MoreFragmentsCounter) setDefault() {
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

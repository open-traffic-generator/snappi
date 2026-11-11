package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NeighborSolicitationCodeCounter *****
type patternFlowIcmpv6NeighborSolicitationCodeCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6NeighborSolicitationCodeCounter
	marshaller   marshalPatternFlowIcmpv6NeighborSolicitationCodeCounter
	unMarshaller unMarshalPatternFlowIcmpv6NeighborSolicitationCodeCounter
}

func NewPatternFlowIcmpv6NeighborSolicitationCodeCounter() PatternFlowIcmpv6NeighborSolicitationCodeCounter {
	obj := patternFlowIcmpv6NeighborSolicitationCodeCounter{obj: &otg.PatternFlowIcmpv6NeighborSolicitationCodeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationCodeCounter) msg() *otg.PatternFlowIcmpv6NeighborSolicitationCodeCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationCodeCounter) setMsg(msg *otg.PatternFlowIcmpv6NeighborSolicitationCodeCounter) PatternFlowIcmpv6NeighborSolicitationCodeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NeighborSolicitationCodeCounter struct {
	obj *patternFlowIcmpv6NeighborSolicitationCodeCounter
}

type marshalPatternFlowIcmpv6NeighborSolicitationCodeCounter interface {
	// ToProto marshals PatternFlowIcmpv6NeighborSolicitationCodeCounter to protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationCodeCounter
	ToProto() (*otg.PatternFlowIcmpv6NeighborSolicitationCodeCounter, error)
	// ToPbText marshals PatternFlowIcmpv6NeighborSolicitationCodeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NeighborSolicitationCodeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NeighborSolicitationCodeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NeighborSolicitationCodeCounter struct {
	obj *patternFlowIcmpv6NeighborSolicitationCodeCounter
}

type unMarshalPatternFlowIcmpv6NeighborSolicitationCodeCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6NeighborSolicitationCodeCounter from protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationCodeCounter
	FromProto(msg *otg.PatternFlowIcmpv6NeighborSolicitationCodeCounter) (PatternFlowIcmpv6NeighborSolicitationCodeCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6NeighborSolicitationCodeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NeighborSolicitationCodeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NeighborSolicitationCodeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NeighborSolicitationCodeCounter) Marshal() marshalPatternFlowIcmpv6NeighborSolicitationCodeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NeighborSolicitationCodeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NeighborSolicitationCodeCounter) Unmarshal() unMarshalPatternFlowIcmpv6NeighborSolicitationCodeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NeighborSolicitationCodeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NeighborSolicitationCodeCounter) ToProto() (*otg.PatternFlowIcmpv6NeighborSolicitationCodeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationCodeCounter) FromProto(msg *otg.PatternFlowIcmpv6NeighborSolicitationCodeCounter) (PatternFlowIcmpv6NeighborSolicitationCodeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NeighborSolicitationCodeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationCodeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NeighborSolicitationCodeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationCodeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NeighborSolicitationCodeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationCodeCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NeighborSolicitationCodeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NeighborSolicitationCodeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NeighborSolicitationCodeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NeighborSolicitationCodeCounter) Clone() (PatternFlowIcmpv6NeighborSolicitationCodeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NeighborSolicitationCodeCounter()
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

// PatternFlowIcmpv6NeighborSolicitationCodeCounter is integer counter pattern
type PatternFlowIcmpv6NeighborSolicitationCodeCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6NeighborSolicitationCodeCounter to protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationCodeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NeighborSolicitationCodeCounter
	// setMsg unmarshals PatternFlowIcmpv6NeighborSolicitationCodeCounter from protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationCodeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NeighborSolicitationCodeCounter) PatternFlowIcmpv6NeighborSolicitationCodeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NeighborSolicitationCodeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NeighborSolicitationCodeCounter
	// validate validates PatternFlowIcmpv6NeighborSolicitationCodeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NeighborSolicitationCodeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpv6NeighborSolicitationCodeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpv6NeighborSolicitationCodeCounter
	SetStart(value uint32) PatternFlowIcmpv6NeighborSolicitationCodeCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6NeighborSolicitationCodeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpv6NeighborSolicitationCodeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpv6NeighborSolicitationCodeCounter
	SetStep(value uint32) PatternFlowIcmpv6NeighborSolicitationCodeCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6NeighborSolicitationCodeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6NeighborSolicitationCodeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6NeighborSolicitationCodeCounter
	SetCount(value uint32) PatternFlowIcmpv6NeighborSolicitationCodeCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6NeighborSolicitationCodeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationCodeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationCodeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpv6NeighborSolicitationCodeCounter object
func (obj *patternFlowIcmpv6NeighborSolicitationCodeCounter) SetStart(value uint32) PatternFlowIcmpv6NeighborSolicitationCodeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationCodeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationCodeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpv6NeighborSolicitationCodeCounter object
func (obj *patternFlowIcmpv6NeighborSolicitationCodeCounter) SetStep(value uint32) PatternFlowIcmpv6NeighborSolicitationCodeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationCodeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationCodeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6NeighborSolicitationCodeCounter object
func (obj *patternFlowIcmpv6NeighborSolicitationCodeCounter) SetCount(value uint32) PatternFlowIcmpv6NeighborSolicitationCodeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationCodeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NeighborSolicitationCodeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NeighborSolicitationCodeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NeighborSolicitationCodeCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIcmpv6NeighborSolicitationCodeCounter) setDefault() {
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

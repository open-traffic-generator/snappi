package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NeighborSolicitationTypeCounter *****
type patternFlowIcmpv6NeighborSolicitationTypeCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6NeighborSolicitationTypeCounter
	marshaller   marshalPatternFlowIcmpv6NeighborSolicitationTypeCounter
	unMarshaller unMarshalPatternFlowIcmpv6NeighborSolicitationTypeCounter
}

func NewPatternFlowIcmpv6NeighborSolicitationTypeCounter() PatternFlowIcmpv6NeighborSolicitationTypeCounter {
	obj := patternFlowIcmpv6NeighborSolicitationTypeCounter{obj: &otg.PatternFlowIcmpv6NeighborSolicitationTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationTypeCounter) msg() *otg.PatternFlowIcmpv6NeighborSolicitationTypeCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationTypeCounter) setMsg(msg *otg.PatternFlowIcmpv6NeighborSolicitationTypeCounter) PatternFlowIcmpv6NeighborSolicitationTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NeighborSolicitationTypeCounter struct {
	obj *patternFlowIcmpv6NeighborSolicitationTypeCounter
}

type marshalPatternFlowIcmpv6NeighborSolicitationTypeCounter interface {
	// ToProto marshals PatternFlowIcmpv6NeighborSolicitationTypeCounter to protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationTypeCounter
	ToProto() (*otg.PatternFlowIcmpv6NeighborSolicitationTypeCounter, error)
	// ToPbText marshals PatternFlowIcmpv6NeighborSolicitationTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NeighborSolicitationTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NeighborSolicitationTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NeighborSolicitationTypeCounter struct {
	obj *patternFlowIcmpv6NeighborSolicitationTypeCounter
}

type unMarshalPatternFlowIcmpv6NeighborSolicitationTypeCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6NeighborSolicitationTypeCounter from protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationTypeCounter
	FromProto(msg *otg.PatternFlowIcmpv6NeighborSolicitationTypeCounter) (PatternFlowIcmpv6NeighborSolicitationTypeCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6NeighborSolicitationTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NeighborSolicitationTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NeighborSolicitationTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NeighborSolicitationTypeCounter) Marshal() marshalPatternFlowIcmpv6NeighborSolicitationTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NeighborSolicitationTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NeighborSolicitationTypeCounter) Unmarshal() unMarshalPatternFlowIcmpv6NeighborSolicitationTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NeighborSolicitationTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NeighborSolicitationTypeCounter) ToProto() (*otg.PatternFlowIcmpv6NeighborSolicitationTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationTypeCounter) FromProto(msg *otg.PatternFlowIcmpv6NeighborSolicitationTypeCounter) (PatternFlowIcmpv6NeighborSolicitationTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NeighborSolicitationTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NeighborSolicitationTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NeighborSolicitationTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NeighborSolicitationTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NeighborSolicitationTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NeighborSolicitationTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NeighborSolicitationTypeCounter) Clone() (PatternFlowIcmpv6NeighborSolicitationTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NeighborSolicitationTypeCounter()
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

// PatternFlowIcmpv6NeighborSolicitationTypeCounter is integer counter pattern
type PatternFlowIcmpv6NeighborSolicitationTypeCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6NeighborSolicitationTypeCounter to protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NeighborSolicitationTypeCounter
	// setMsg unmarshals PatternFlowIcmpv6NeighborSolicitationTypeCounter from protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NeighborSolicitationTypeCounter) PatternFlowIcmpv6NeighborSolicitationTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NeighborSolicitationTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NeighborSolicitationTypeCounter
	// validate validates PatternFlowIcmpv6NeighborSolicitationTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NeighborSolicitationTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpv6NeighborSolicitationTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpv6NeighborSolicitationTypeCounter
	SetStart(value uint32) PatternFlowIcmpv6NeighborSolicitationTypeCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6NeighborSolicitationTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpv6NeighborSolicitationTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpv6NeighborSolicitationTypeCounter
	SetStep(value uint32) PatternFlowIcmpv6NeighborSolicitationTypeCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6NeighborSolicitationTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6NeighborSolicitationTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6NeighborSolicitationTypeCounter
	SetCount(value uint32) PatternFlowIcmpv6NeighborSolicitationTypeCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6NeighborSolicitationTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpv6NeighborSolicitationTypeCounter object
func (obj *patternFlowIcmpv6NeighborSolicitationTypeCounter) SetStart(value uint32) PatternFlowIcmpv6NeighborSolicitationTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpv6NeighborSolicitationTypeCounter object
func (obj *patternFlowIcmpv6NeighborSolicitationTypeCounter) SetStep(value uint32) PatternFlowIcmpv6NeighborSolicitationTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6NeighborSolicitationTypeCounter object
func (obj *patternFlowIcmpv6NeighborSolicitationTypeCounter) SetCount(value uint32) PatternFlowIcmpv6NeighborSolicitationTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NeighborSolicitationTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NeighborSolicitationTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NeighborSolicitationTypeCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIcmpv6NeighborSolicitationTypeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(135)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

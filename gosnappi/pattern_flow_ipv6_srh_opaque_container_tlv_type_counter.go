package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter *****
type patternFlowIpv6SRHOpaqueContainerTlvTypeCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	marshaller   marshalPatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	unMarshaller unMarshalPatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
}

func NewPatternFlowIpv6SRHOpaqueContainerTlvTypeCounter() PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter {
	obj := patternFlowIpv6SRHOpaqueContainerTlvTypeCounter{obj: &otg.PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvTypeCounter) msg() *otg.PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvTypeCounter) setMsg(msg *otg.PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter) PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHOpaqueContainerTlvTypeCounter struct {
	obj *patternFlowIpv6SRHOpaqueContainerTlvTypeCounter
}

type marshalPatternFlowIpv6SRHOpaqueContainerTlvTypeCounter interface {
	// ToProto marshals PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter to protobuf object *otg.PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	ToProto() (*otg.PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter, error)
	// ToPbText marshals PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHOpaqueContainerTlvTypeCounter struct {
	obj *patternFlowIpv6SRHOpaqueContainerTlvTypeCounter
}

type unMarshalPatternFlowIpv6SRHOpaqueContainerTlvTypeCounter interface {
	// FromProto unmarshals PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter from protobuf object *otg.PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	FromProto(msg *otg.PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter) (PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvTypeCounter) Marshal() marshalPatternFlowIpv6SRHOpaqueContainerTlvTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHOpaqueContainerTlvTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvTypeCounter) Unmarshal() unMarshalPatternFlowIpv6SRHOpaqueContainerTlvTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHOpaqueContainerTlvTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHOpaqueContainerTlvTypeCounter) ToProto() (*otg.PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHOpaqueContainerTlvTypeCounter) FromProto(msg *otg.PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter) (PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHOpaqueContainerTlvTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHOpaqueContainerTlvTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHOpaqueContainerTlvTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHOpaqueContainerTlvTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHOpaqueContainerTlvTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHOpaqueContainerTlvTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHOpaqueContainerTlvTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvTypeCounter) Clone() (PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHOpaqueContainerTlvTypeCounter()
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

// PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter is integer counter pattern
type PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter to protobuf object *otg.PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	// setMsg unmarshals PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter from protobuf object *otg.PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter) PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	// validate validates PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	SetStart(value uint32) PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	SetStep(value uint32) PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	SetCount(value uint32) PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHOpaqueContainerTlvTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHOpaqueContainerTlvTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter object
func (obj *patternFlowIpv6SRHOpaqueContainerTlvTypeCounter) SetStart(value uint32) PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHOpaqueContainerTlvTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHOpaqueContainerTlvTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter object
func (obj *patternFlowIpv6SRHOpaqueContainerTlvTypeCounter) SetStep(value uint32) PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHOpaqueContainerTlvTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHOpaqueContainerTlvTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter object
func (obj *patternFlowIpv6SRHOpaqueContainerTlvTypeCounter) SetCount(value uint32) PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvTypeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(3)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

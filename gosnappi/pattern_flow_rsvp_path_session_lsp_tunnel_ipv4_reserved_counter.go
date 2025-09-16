package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter *****
type patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	marshaller   marshalPatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	unMarshaller unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
}

func NewPatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter() PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter {
	obj := patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter{obj: &otg.PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) msg() *otg.PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) setMsg(msg *otg.PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter struct {
	obj *patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
}

type marshalPatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter interface {
	// ToProto marshals PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter to protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	ToProto() (*otg.PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter, error)
	// ToPbText marshals PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter struct {
	obj *patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
}

type unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter from protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	FromProto(msg *otg.PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) (PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) Marshal() marshalPatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) Unmarshal() unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) ToProto() (*otg.PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) FromProto(msg *otg.PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) (PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) Clone() (PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter()
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

// PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter is integer counter pattern
type PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter to protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	// setMsg unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter from protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	// validate validates PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	SetStart(value uint32) PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	SetStep(value uint32) PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	SetCount(value uint32) PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter object
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) SetStart(value uint32) PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter object
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) SetStep(value uint32) PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter object
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) SetCount(value uint32) PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) setDefault() {
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

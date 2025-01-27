package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter *****
type patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	marshaller   marshalPatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	unMarshaller unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
}

func NewPatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter() PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter {
	obj := patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter{obj: &otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) msg() *otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) setMsg(msg *otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter struct {
	obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
}

type marshalPatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter interface {
	// ToProto marshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter to protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	ToProto() (*otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter, error)
	// ToPbText marshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter struct {
	obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
}

type unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter from protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	FromProto(msg *otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) (PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) Marshal() marshalPatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) Unmarshal() unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) ToProto() (*otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) FromProto(msg *otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) (PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) Clone() (PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter()
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

// PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter is integer counter pattern
type PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter to protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	// setMsg unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter from protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	// validate validates PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	SetStart(value uint32) PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	SetStep(value uint32) PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	SetCount(value uint32) PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter object
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) SetStart(value uint32) PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter object
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) SetStep(value uint32) PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter object
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) SetCount(value uint32) PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(1)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

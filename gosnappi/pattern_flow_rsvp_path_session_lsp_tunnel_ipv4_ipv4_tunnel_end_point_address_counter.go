package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter *****
type patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	marshaller   marshalPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	unMarshaller unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
}

func NewPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter() PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter {
	obj := patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter{obj: &otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) msg() *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) setMsg(msg *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter struct {
	obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
}

type marshalPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter interface {
	// ToProto marshals PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter to protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	ToProto() (*otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter, error)
	// ToPbText marshals PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter struct {
	obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
}

type unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter from protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	FromProto(msg *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) (PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) Marshal() marshalPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) Unmarshal() unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) ToProto() (*otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) FromProto(msg *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) (PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) Clone() (PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter()
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

// PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter is ipv4 counter pattern
type PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter to protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	// setMsg unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter from protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	// validate validates PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	SetStart(value string) PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	HasStart() bool
	// Step returns string, set in PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	SetStep(value string) PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	SetCount(value uint32) PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter object
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) SetStart(value string) PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter object
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) SetStep(value string) PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter object
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) SetCount(value uint32) PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter.Step"))
		}

	}

}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("0.0.0.0")
	}
	if obj.obj.Step == nil {
		obj.SetStep("0.0.0.1")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

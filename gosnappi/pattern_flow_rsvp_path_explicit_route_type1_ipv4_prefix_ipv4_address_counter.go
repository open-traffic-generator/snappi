package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter *****
type patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	marshaller   marshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	unMarshaller unMarshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
}

func NewPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter() PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter {
	obj := patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter{obj: &otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) msg() *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) setMsg(msg *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter struct {
	obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
}

type marshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter interface {
	// ToProto marshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter to protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	ToProto() (*otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter, error)
	// ToPbText marshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter struct {
	obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
}

type unMarshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter from protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	FromProto(msg *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) (PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) Marshal() marshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) Unmarshal() unMarshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) ToProto() (*otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) FromProto(msg *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) (PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) Clone() (PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter()
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

// PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter is ipv4 counter pattern
type PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter to protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	// setMsg unmarshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter from protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	// validate validates PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	SetStart(value string) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	HasStart() bool
	// Step returns string, set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	SetStep(value string) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	SetCount(value uint32) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter object
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) SetStart(value string) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter object
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) SetStep(value string) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter object
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) SetCount(value uint32) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter.Step"))
		}

	}

}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) setDefault() {
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

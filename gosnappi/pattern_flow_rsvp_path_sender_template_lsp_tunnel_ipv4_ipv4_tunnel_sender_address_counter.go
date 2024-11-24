package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter *****
type patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	marshaller   marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	unMarshaller unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
}

func NewPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter {
	obj := patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter{obj: &otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) msg() *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) setMsg(msg *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter struct {
	obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
}

type marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter interface {
	// ToProto marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter to protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	ToProto() (*otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter struct {
	obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
}

type unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter from protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	FromProto(msg *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) Marshal() marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) Unmarshal() unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) ToProto() (*otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) FromProto(msg *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) Clone() (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter()
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

// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter is ipv4 counter pattern
type PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter to protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	// setMsg unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter from protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	// validate validates PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	SetStart(value string) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	HasStart() bool
	// Step returns string, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	SetStep(value string) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	SetCount(value uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter object
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) SetStart(value string) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter object
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) SetStep(value string) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter object
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) SetCount(value uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter.Step"))
		}

	}

}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) setDefault() {
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

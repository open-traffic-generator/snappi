package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter *****
type patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	marshaller   marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	unMarshaller unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
}

func NewPatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter {
	obj := patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter{obj: &otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) msg() *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) setMsg(msg *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter struct {
	obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
}

type marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter interface {
	// ToProto marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter to protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	ToProto() (*otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter struct {
	obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
}

type unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter from protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	FromProto(msg *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) Marshal() marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) Unmarshal() unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) ToProto() (*otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) FromProto(msg *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) Clone() (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter()
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

// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter is integer counter pattern
type PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter to protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	// setMsg unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter from protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	// validate validates PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	SetStart(value uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	SetStep(value uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	SetCount(value uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter object
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) SetStart(value uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter object
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) SetStep(value uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter object
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) SetCount(value uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) setDefault() {
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

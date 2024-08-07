package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter *****
type patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	marshaller   marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	unMarshaller unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
}

func NewPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter {
	obj := patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter{obj: &otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) msg() *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) setMsg(msg *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter struct {
	obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
}

type marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter interface {
	// ToProto marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter to protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	ToProto() (*otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter struct {
	obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
}

type unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter from protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	FromProto(msg *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) Marshal() marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) Unmarshal() unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) ToProto() (*otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) FromProto(msg *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) Clone() (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter()
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

// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter is integer counter pattern
type PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter to protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	// setMsg unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter from protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	// validate validates PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	SetStart(value uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	SetStep(value uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	SetCount(value uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter object
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) SetStart(value uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter object
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) SetStep(value uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter object
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) SetCount(value uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) setDefault() {
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

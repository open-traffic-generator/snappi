package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHIngressNodeTlvTypeCounter *****
type patternFlowIpv6SRHIngressNodeTlvTypeCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SRHIngressNodeTlvTypeCounter
	marshaller   marshalPatternFlowIpv6SRHIngressNodeTlvTypeCounter
	unMarshaller unMarshalPatternFlowIpv6SRHIngressNodeTlvTypeCounter
}

func NewPatternFlowIpv6SRHIngressNodeTlvTypeCounter() PatternFlowIpv6SRHIngressNodeTlvTypeCounter {
	obj := patternFlowIpv6SRHIngressNodeTlvTypeCounter{obj: &otg.PatternFlowIpv6SRHIngressNodeTlvTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvTypeCounter) msg() *otg.PatternFlowIpv6SRHIngressNodeTlvTypeCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvTypeCounter) setMsg(msg *otg.PatternFlowIpv6SRHIngressNodeTlvTypeCounter) PatternFlowIpv6SRHIngressNodeTlvTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHIngressNodeTlvTypeCounter struct {
	obj *patternFlowIpv6SRHIngressNodeTlvTypeCounter
}

type marshalPatternFlowIpv6SRHIngressNodeTlvTypeCounter interface {
	// ToProto marshals PatternFlowIpv6SRHIngressNodeTlvTypeCounter to protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvTypeCounter
	ToProto() (*otg.PatternFlowIpv6SRHIngressNodeTlvTypeCounter, error)
	// ToPbText marshals PatternFlowIpv6SRHIngressNodeTlvTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHIngressNodeTlvTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHIngressNodeTlvTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHIngressNodeTlvTypeCounter struct {
	obj *patternFlowIpv6SRHIngressNodeTlvTypeCounter
}

type unMarshalPatternFlowIpv6SRHIngressNodeTlvTypeCounter interface {
	// FromProto unmarshals PatternFlowIpv6SRHIngressNodeTlvTypeCounter from protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvTypeCounter
	FromProto(msg *otg.PatternFlowIpv6SRHIngressNodeTlvTypeCounter) (PatternFlowIpv6SRHIngressNodeTlvTypeCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SRHIngressNodeTlvTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHIngressNodeTlvTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHIngressNodeTlvTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHIngressNodeTlvTypeCounter) Marshal() marshalPatternFlowIpv6SRHIngressNodeTlvTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHIngressNodeTlvTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHIngressNodeTlvTypeCounter) Unmarshal() unMarshalPatternFlowIpv6SRHIngressNodeTlvTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHIngressNodeTlvTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvTypeCounter) ToProto() (*otg.PatternFlowIpv6SRHIngressNodeTlvTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvTypeCounter) FromProto(msg *otg.PatternFlowIpv6SRHIngressNodeTlvTypeCounter) (PatternFlowIpv6SRHIngressNodeTlvTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHIngressNodeTlvTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHIngressNodeTlvTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHIngressNodeTlvTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHIngressNodeTlvTypeCounter) Clone() (PatternFlowIpv6SRHIngressNodeTlvTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHIngressNodeTlvTypeCounter()
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

// PatternFlowIpv6SRHIngressNodeTlvTypeCounter is integer counter pattern
type PatternFlowIpv6SRHIngressNodeTlvTypeCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SRHIngressNodeTlvTypeCounter to protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHIngressNodeTlvTypeCounter
	// setMsg unmarshals PatternFlowIpv6SRHIngressNodeTlvTypeCounter from protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHIngressNodeTlvTypeCounter) PatternFlowIpv6SRHIngressNodeTlvTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHIngressNodeTlvTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHIngressNodeTlvTypeCounter
	// validate validates PatternFlowIpv6SRHIngressNodeTlvTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHIngressNodeTlvTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SRHIngressNodeTlvTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SRHIngressNodeTlvTypeCounter
	SetStart(value uint32) PatternFlowIpv6SRHIngressNodeTlvTypeCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SRHIngressNodeTlvTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SRHIngressNodeTlvTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SRHIngressNodeTlvTypeCounter
	SetStep(value uint32) PatternFlowIpv6SRHIngressNodeTlvTypeCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SRHIngressNodeTlvTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SRHIngressNodeTlvTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SRHIngressNodeTlvTypeCounter
	SetCount(value uint32) PatternFlowIpv6SRHIngressNodeTlvTypeCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SRHIngressNodeTlvTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SRHIngressNodeTlvTypeCounter object
func (obj *patternFlowIpv6SRHIngressNodeTlvTypeCounter) SetStart(value uint32) PatternFlowIpv6SRHIngressNodeTlvTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SRHIngressNodeTlvTypeCounter object
func (obj *patternFlowIpv6SRHIngressNodeTlvTypeCounter) SetStep(value uint32) PatternFlowIpv6SRHIngressNodeTlvTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SRHIngressNodeTlvTypeCounter object
func (obj *patternFlowIpv6SRHIngressNodeTlvTypeCounter) SetCount(value uint32) PatternFlowIpv6SRHIngressNodeTlvTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHIngressNodeTlvTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHIngressNodeTlvTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHIngressNodeTlvTypeCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SRHIngressNodeTlvTypeCounter) setDefault() {
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

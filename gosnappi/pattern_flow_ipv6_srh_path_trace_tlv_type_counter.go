package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHPathTraceTlvTypeCounter *****
type patternFlowIpv6SRHPathTraceTlvTypeCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SRHPathTraceTlvTypeCounter
	marshaller   marshalPatternFlowIpv6SRHPathTraceTlvTypeCounter
	unMarshaller unMarshalPatternFlowIpv6SRHPathTraceTlvTypeCounter
}

func NewPatternFlowIpv6SRHPathTraceTlvTypeCounter() PatternFlowIpv6SRHPathTraceTlvTypeCounter {
	obj := patternFlowIpv6SRHPathTraceTlvTypeCounter{obj: &otg.PatternFlowIpv6SRHPathTraceTlvTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvTypeCounter) msg() *otg.PatternFlowIpv6SRHPathTraceTlvTypeCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvTypeCounter) setMsg(msg *otg.PatternFlowIpv6SRHPathTraceTlvTypeCounter) PatternFlowIpv6SRHPathTraceTlvTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHPathTraceTlvTypeCounter struct {
	obj *patternFlowIpv6SRHPathTraceTlvTypeCounter
}

type marshalPatternFlowIpv6SRHPathTraceTlvTypeCounter interface {
	// ToProto marshals PatternFlowIpv6SRHPathTraceTlvTypeCounter to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvTypeCounter
	ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvTypeCounter, error)
	// ToPbText marshals PatternFlowIpv6SRHPathTraceTlvTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHPathTraceTlvTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHPathTraceTlvTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHPathTraceTlvTypeCounter struct {
	obj *patternFlowIpv6SRHPathTraceTlvTypeCounter
}

type unMarshalPatternFlowIpv6SRHPathTraceTlvTypeCounter interface {
	// FromProto unmarshals PatternFlowIpv6SRHPathTraceTlvTypeCounter from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvTypeCounter
	FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvTypeCounter) (PatternFlowIpv6SRHPathTraceTlvTypeCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SRHPathTraceTlvTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHPathTraceTlvTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHPathTraceTlvTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHPathTraceTlvTypeCounter) Marshal() marshalPatternFlowIpv6SRHPathTraceTlvTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHPathTraceTlvTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHPathTraceTlvTypeCounter) Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHPathTraceTlvTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvTypeCounter) ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvTypeCounter) FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvTypeCounter) (PatternFlowIpv6SRHPathTraceTlvTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHPathTraceTlvTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHPathTraceTlvTypeCounter) Clone() (PatternFlowIpv6SRHPathTraceTlvTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHPathTraceTlvTypeCounter()
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

// PatternFlowIpv6SRHPathTraceTlvTypeCounter is integer counter pattern
type PatternFlowIpv6SRHPathTraceTlvTypeCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SRHPathTraceTlvTypeCounter to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHPathTraceTlvTypeCounter
	// setMsg unmarshals PatternFlowIpv6SRHPathTraceTlvTypeCounter from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHPathTraceTlvTypeCounter) PatternFlowIpv6SRHPathTraceTlvTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHPathTraceTlvTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvTypeCounter
	// validate validates PatternFlowIpv6SRHPathTraceTlvTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHPathTraceTlvTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SRHPathTraceTlvTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvTypeCounter
	SetStart(value uint32) PatternFlowIpv6SRHPathTraceTlvTypeCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SRHPathTraceTlvTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SRHPathTraceTlvTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvTypeCounter
	SetStep(value uint32) PatternFlowIpv6SRHPathTraceTlvTypeCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SRHPathTraceTlvTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SRHPathTraceTlvTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvTypeCounter
	SetCount(value uint32) PatternFlowIpv6SRHPathTraceTlvTypeCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SRHPathTraceTlvTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvTypeCounter object
func (obj *patternFlowIpv6SRHPathTraceTlvTypeCounter) SetStart(value uint32) PatternFlowIpv6SRHPathTraceTlvTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvTypeCounter object
func (obj *patternFlowIpv6SRHPathTraceTlvTypeCounter) SetStep(value uint32) PatternFlowIpv6SRHPathTraceTlvTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvTypeCounter object
func (obj *patternFlowIpv6SRHPathTraceTlvTypeCounter) SetCount(value uint32) PatternFlowIpv6SRHPathTraceTlvTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPathTraceTlvTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPathTraceTlvTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPathTraceTlvTypeCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SRHPathTraceTlvTypeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(128)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

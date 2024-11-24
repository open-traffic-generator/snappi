package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpSenderProtocolAddrCounter *****
type patternFlowArpSenderProtocolAddrCounter struct {
	validation
	obj          *otg.PatternFlowArpSenderProtocolAddrCounter
	marshaller   marshalPatternFlowArpSenderProtocolAddrCounter
	unMarshaller unMarshalPatternFlowArpSenderProtocolAddrCounter
}

func NewPatternFlowArpSenderProtocolAddrCounter() PatternFlowArpSenderProtocolAddrCounter {
	obj := patternFlowArpSenderProtocolAddrCounter{obj: &otg.PatternFlowArpSenderProtocolAddrCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpSenderProtocolAddrCounter) msg() *otg.PatternFlowArpSenderProtocolAddrCounter {
	return obj.obj
}

func (obj *patternFlowArpSenderProtocolAddrCounter) setMsg(msg *otg.PatternFlowArpSenderProtocolAddrCounter) PatternFlowArpSenderProtocolAddrCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpSenderProtocolAddrCounter struct {
	obj *patternFlowArpSenderProtocolAddrCounter
}

type marshalPatternFlowArpSenderProtocolAddrCounter interface {
	// ToProto marshals PatternFlowArpSenderProtocolAddrCounter to protobuf object *otg.PatternFlowArpSenderProtocolAddrCounter
	ToProto() (*otg.PatternFlowArpSenderProtocolAddrCounter, error)
	// ToPbText marshals PatternFlowArpSenderProtocolAddrCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpSenderProtocolAddrCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpSenderProtocolAddrCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowArpSenderProtocolAddrCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowArpSenderProtocolAddrCounter struct {
	obj *patternFlowArpSenderProtocolAddrCounter
}

type unMarshalPatternFlowArpSenderProtocolAddrCounter interface {
	// FromProto unmarshals PatternFlowArpSenderProtocolAddrCounter from protobuf object *otg.PatternFlowArpSenderProtocolAddrCounter
	FromProto(msg *otg.PatternFlowArpSenderProtocolAddrCounter) (PatternFlowArpSenderProtocolAddrCounter, error)
	// FromPbText unmarshals PatternFlowArpSenderProtocolAddrCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpSenderProtocolAddrCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpSenderProtocolAddrCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpSenderProtocolAddrCounter) Marshal() marshalPatternFlowArpSenderProtocolAddrCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpSenderProtocolAddrCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpSenderProtocolAddrCounter) Unmarshal() unMarshalPatternFlowArpSenderProtocolAddrCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpSenderProtocolAddrCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpSenderProtocolAddrCounter) ToProto() (*otg.PatternFlowArpSenderProtocolAddrCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpSenderProtocolAddrCounter) FromProto(msg *otg.PatternFlowArpSenderProtocolAddrCounter) (PatternFlowArpSenderProtocolAddrCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpSenderProtocolAddrCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpSenderProtocolAddrCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpSenderProtocolAddrCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpSenderProtocolAddrCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpSenderProtocolAddrCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowArpSenderProtocolAddrCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpSenderProtocolAddrCounter) FromJson(value string) error {
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

func (obj *patternFlowArpSenderProtocolAddrCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpSenderProtocolAddrCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpSenderProtocolAddrCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpSenderProtocolAddrCounter) Clone() (PatternFlowArpSenderProtocolAddrCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpSenderProtocolAddrCounter()
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

// PatternFlowArpSenderProtocolAddrCounter is ipv4 counter pattern
type PatternFlowArpSenderProtocolAddrCounter interface {
	Validation
	// msg marshals PatternFlowArpSenderProtocolAddrCounter to protobuf object *otg.PatternFlowArpSenderProtocolAddrCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowArpSenderProtocolAddrCounter
	// setMsg unmarshals PatternFlowArpSenderProtocolAddrCounter from protobuf object *otg.PatternFlowArpSenderProtocolAddrCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpSenderProtocolAddrCounter) PatternFlowArpSenderProtocolAddrCounter
	// provides marshal interface
	Marshal() marshalPatternFlowArpSenderProtocolAddrCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpSenderProtocolAddrCounter
	// validate validates PatternFlowArpSenderProtocolAddrCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpSenderProtocolAddrCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowArpSenderProtocolAddrCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowArpSenderProtocolAddrCounter
	SetStart(value string) PatternFlowArpSenderProtocolAddrCounter
	// HasStart checks if Start has been set in PatternFlowArpSenderProtocolAddrCounter
	HasStart() bool
	// Step returns string, set in PatternFlowArpSenderProtocolAddrCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowArpSenderProtocolAddrCounter
	SetStep(value string) PatternFlowArpSenderProtocolAddrCounter
	// HasStep checks if Step has been set in PatternFlowArpSenderProtocolAddrCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowArpSenderProtocolAddrCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowArpSenderProtocolAddrCounter
	SetCount(value uint32) PatternFlowArpSenderProtocolAddrCounter
	// HasCount checks if Count has been set in PatternFlowArpSenderProtocolAddrCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowArpSenderProtocolAddrCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowArpSenderProtocolAddrCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowArpSenderProtocolAddrCounter object
func (obj *patternFlowArpSenderProtocolAddrCounter) SetStart(value string) PatternFlowArpSenderProtocolAddrCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowArpSenderProtocolAddrCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowArpSenderProtocolAddrCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowArpSenderProtocolAddrCounter object
func (obj *patternFlowArpSenderProtocolAddrCounter) SetStep(value string) PatternFlowArpSenderProtocolAddrCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowArpSenderProtocolAddrCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowArpSenderProtocolAddrCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowArpSenderProtocolAddrCounter object
func (obj *patternFlowArpSenderProtocolAddrCounter) SetCount(value uint32) PatternFlowArpSenderProtocolAddrCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowArpSenderProtocolAddrCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpSenderProtocolAddrCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpSenderProtocolAddrCounter.Step"))
		}

	}

}

func (obj *patternFlowArpSenderProtocolAddrCounter) setDefault() {
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

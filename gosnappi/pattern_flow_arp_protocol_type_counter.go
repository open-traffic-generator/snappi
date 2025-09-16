package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpProtocolTypeCounter *****
type patternFlowArpProtocolTypeCounter struct {
	validation
	obj          *otg.PatternFlowArpProtocolTypeCounter
	marshaller   marshalPatternFlowArpProtocolTypeCounter
	unMarshaller unMarshalPatternFlowArpProtocolTypeCounter
}

func NewPatternFlowArpProtocolTypeCounter() PatternFlowArpProtocolTypeCounter {
	obj := patternFlowArpProtocolTypeCounter{obj: &otg.PatternFlowArpProtocolTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpProtocolTypeCounter) msg() *otg.PatternFlowArpProtocolTypeCounter {
	return obj.obj
}

func (obj *patternFlowArpProtocolTypeCounter) setMsg(msg *otg.PatternFlowArpProtocolTypeCounter) PatternFlowArpProtocolTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpProtocolTypeCounter struct {
	obj *patternFlowArpProtocolTypeCounter
}

type marshalPatternFlowArpProtocolTypeCounter interface {
	// ToProto marshals PatternFlowArpProtocolTypeCounter to protobuf object *otg.PatternFlowArpProtocolTypeCounter
	ToProto() (*otg.PatternFlowArpProtocolTypeCounter, error)
	// ToPbText marshals PatternFlowArpProtocolTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpProtocolTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpProtocolTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowArpProtocolTypeCounter struct {
	obj *patternFlowArpProtocolTypeCounter
}

type unMarshalPatternFlowArpProtocolTypeCounter interface {
	// FromProto unmarshals PatternFlowArpProtocolTypeCounter from protobuf object *otg.PatternFlowArpProtocolTypeCounter
	FromProto(msg *otg.PatternFlowArpProtocolTypeCounter) (PatternFlowArpProtocolTypeCounter, error)
	// FromPbText unmarshals PatternFlowArpProtocolTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpProtocolTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpProtocolTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpProtocolTypeCounter) Marshal() marshalPatternFlowArpProtocolTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpProtocolTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpProtocolTypeCounter) Unmarshal() unMarshalPatternFlowArpProtocolTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpProtocolTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpProtocolTypeCounter) ToProto() (*otg.PatternFlowArpProtocolTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpProtocolTypeCounter) FromProto(msg *otg.PatternFlowArpProtocolTypeCounter) (PatternFlowArpProtocolTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpProtocolTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpProtocolTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpProtocolTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpProtocolTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpProtocolTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpProtocolTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowArpProtocolTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpProtocolTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpProtocolTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpProtocolTypeCounter) Clone() (PatternFlowArpProtocolTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpProtocolTypeCounter()
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

// PatternFlowArpProtocolTypeCounter is integer counter pattern
type PatternFlowArpProtocolTypeCounter interface {
	Validation
	// msg marshals PatternFlowArpProtocolTypeCounter to protobuf object *otg.PatternFlowArpProtocolTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowArpProtocolTypeCounter
	// setMsg unmarshals PatternFlowArpProtocolTypeCounter from protobuf object *otg.PatternFlowArpProtocolTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpProtocolTypeCounter) PatternFlowArpProtocolTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowArpProtocolTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpProtocolTypeCounter
	// validate validates PatternFlowArpProtocolTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpProtocolTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowArpProtocolTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowArpProtocolTypeCounter
	SetStart(value uint32) PatternFlowArpProtocolTypeCounter
	// HasStart checks if Start has been set in PatternFlowArpProtocolTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowArpProtocolTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowArpProtocolTypeCounter
	SetStep(value uint32) PatternFlowArpProtocolTypeCounter
	// HasStep checks if Step has been set in PatternFlowArpProtocolTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowArpProtocolTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowArpProtocolTypeCounter
	SetCount(value uint32) PatternFlowArpProtocolTypeCounter
	// HasCount checks if Count has been set in PatternFlowArpProtocolTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowArpProtocolTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowArpProtocolTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowArpProtocolTypeCounter object
func (obj *patternFlowArpProtocolTypeCounter) SetStart(value uint32) PatternFlowArpProtocolTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowArpProtocolTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowArpProtocolTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowArpProtocolTypeCounter object
func (obj *patternFlowArpProtocolTypeCounter) SetStep(value uint32) PatternFlowArpProtocolTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowArpProtocolTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowArpProtocolTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowArpProtocolTypeCounter object
func (obj *patternFlowArpProtocolTypeCounter) SetCount(value uint32) PatternFlowArpProtocolTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowArpProtocolTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpProtocolTypeCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpProtocolTypeCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpProtocolTypeCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowArpProtocolTypeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(2048)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

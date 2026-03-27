package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerPortPriorityCounter *****
type patternFlowLacpduPartnerPortPriorityCounter struct {
	validation
	obj          *otg.PatternFlowLacpduPartnerPortPriorityCounter
	marshaller   marshalPatternFlowLacpduPartnerPortPriorityCounter
	unMarshaller unMarshalPatternFlowLacpduPartnerPortPriorityCounter
}

func NewPatternFlowLacpduPartnerPortPriorityCounter() PatternFlowLacpduPartnerPortPriorityCounter {
	obj := patternFlowLacpduPartnerPortPriorityCounter{obj: &otg.PatternFlowLacpduPartnerPortPriorityCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerPortPriorityCounter) msg() *otg.PatternFlowLacpduPartnerPortPriorityCounter {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerPortPriorityCounter) setMsg(msg *otg.PatternFlowLacpduPartnerPortPriorityCounter) PatternFlowLacpduPartnerPortPriorityCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerPortPriorityCounter struct {
	obj *patternFlowLacpduPartnerPortPriorityCounter
}

type marshalPatternFlowLacpduPartnerPortPriorityCounter interface {
	// ToProto marshals PatternFlowLacpduPartnerPortPriorityCounter to protobuf object *otg.PatternFlowLacpduPartnerPortPriorityCounter
	ToProto() (*otg.PatternFlowLacpduPartnerPortPriorityCounter, error)
	// ToPbText marshals PatternFlowLacpduPartnerPortPriorityCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerPortPriorityCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerPortPriorityCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerPortPriorityCounter struct {
	obj *patternFlowLacpduPartnerPortPriorityCounter
}

type unMarshalPatternFlowLacpduPartnerPortPriorityCounter interface {
	// FromProto unmarshals PatternFlowLacpduPartnerPortPriorityCounter from protobuf object *otg.PatternFlowLacpduPartnerPortPriorityCounter
	FromProto(msg *otg.PatternFlowLacpduPartnerPortPriorityCounter) (PatternFlowLacpduPartnerPortPriorityCounter, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerPortPriorityCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerPortPriorityCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerPortPriorityCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerPortPriorityCounter) Marshal() marshalPatternFlowLacpduPartnerPortPriorityCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerPortPriorityCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerPortPriorityCounter) Unmarshal() unMarshalPatternFlowLacpduPartnerPortPriorityCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerPortPriorityCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerPortPriorityCounter) ToProto() (*otg.PatternFlowLacpduPartnerPortPriorityCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerPortPriorityCounter) FromProto(msg *otg.PatternFlowLacpduPartnerPortPriorityCounter) (PatternFlowLacpduPartnerPortPriorityCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerPortPriorityCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerPortPriorityCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerPortPriorityCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerPortPriorityCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerPortPriorityCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerPortPriorityCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerPortPriorityCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerPortPriorityCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerPortPriorityCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerPortPriorityCounter) Clone() (PatternFlowLacpduPartnerPortPriorityCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerPortPriorityCounter()
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

// PatternFlowLacpduPartnerPortPriorityCounter is integer counter pattern
type PatternFlowLacpduPartnerPortPriorityCounter interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerPortPriorityCounter to protobuf object *otg.PatternFlowLacpduPartnerPortPriorityCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerPortPriorityCounter
	// setMsg unmarshals PatternFlowLacpduPartnerPortPriorityCounter from protobuf object *otg.PatternFlowLacpduPartnerPortPriorityCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerPortPriorityCounter) PatternFlowLacpduPartnerPortPriorityCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerPortPriorityCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerPortPriorityCounter
	// validate validates PatternFlowLacpduPartnerPortPriorityCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerPortPriorityCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduPartnerPortPriorityCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduPartnerPortPriorityCounter
	SetStart(value uint32) PatternFlowLacpduPartnerPortPriorityCounter
	// HasStart checks if Start has been set in PatternFlowLacpduPartnerPortPriorityCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduPartnerPortPriorityCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduPartnerPortPriorityCounter
	SetStep(value uint32) PatternFlowLacpduPartnerPortPriorityCounter
	// HasStep checks if Step has been set in PatternFlowLacpduPartnerPortPriorityCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduPartnerPortPriorityCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduPartnerPortPriorityCounter
	SetCount(value uint32) PatternFlowLacpduPartnerPortPriorityCounter
	// HasCount checks if Count has been set in PatternFlowLacpduPartnerPortPriorityCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerPortPriorityCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerPortPriorityCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduPartnerPortPriorityCounter object
func (obj *patternFlowLacpduPartnerPortPriorityCounter) SetStart(value uint32) PatternFlowLacpduPartnerPortPriorityCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerPortPriorityCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerPortPriorityCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduPartnerPortPriorityCounter object
func (obj *patternFlowLacpduPartnerPortPriorityCounter) SetStep(value uint32) PatternFlowLacpduPartnerPortPriorityCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerPortPriorityCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerPortPriorityCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduPartnerPortPriorityCounter object
func (obj *patternFlowLacpduPartnerPortPriorityCounter) SetCount(value uint32) PatternFlowLacpduPartnerPortPriorityCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduPartnerPortPriorityCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerPortPriorityCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerPortPriorityCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerPortPriorityCounter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduPartnerPortPriorityCounter) setDefault() {
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

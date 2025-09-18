package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerSystemPriorityCounter *****
type patternFlowLacpduPartnerSystemPriorityCounter struct {
	validation
	obj          *otg.PatternFlowLacpduPartnerSystemPriorityCounter
	marshaller   marshalPatternFlowLacpduPartnerSystemPriorityCounter
	unMarshaller unMarshalPatternFlowLacpduPartnerSystemPriorityCounter
}

func NewPatternFlowLacpduPartnerSystemPriorityCounter() PatternFlowLacpduPartnerSystemPriorityCounter {
	obj := patternFlowLacpduPartnerSystemPriorityCounter{obj: &otg.PatternFlowLacpduPartnerSystemPriorityCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerSystemPriorityCounter) msg() *otg.PatternFlowLacpduPartnerSystemPriorityCounter {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerSystemPriorityCounter) setMsg(msg *otg.PatternFlowLacpduPartnerSystemPriorityCounter) PatternFlowLacpduPartnerSystemPriorityCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerSystemPriorityCounter struct {
	obj *patternFlowLacpduPartnerSystemPriorityCounter
}

type marshalPatternFlowLacpduPartnerSystemPriorityCounter interface {
	// ToProto marshals PatternFlowLacpduPartnerSystemPriorityCounter to protobuf object *otg.PatternFlowLacpduPartnerSystemPriorityCounter
	ToProto() (*otg.PatternFlowLacpduPartnerSystemPriorityCounter, error)
	// ToPbText marshals PatternFlowLacpduPartnerSystemPriorityCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerSystemPriorityCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerSystemPriorityCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerSystemPriorityCounter struct {
	obj *patternFlowLacpduPartnerSystemPriorityCounter
}

type unMarshalPatternFlowLacpduPartnerSystemPriorityCounter interface {
	// FromProto unmarshals PatternFlowLacpduPartnerSystemPriorityCounter from protobuf object *otg.PatternFlowLacpduPartnerSystemPriorityCounter
	FromProto(msg *otg.PatternFlowLacpduPartnerSystemPriorityCounter) (PatternFlowLacpduPartnerSystemPriorityCounter, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerSystemPriorityCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerSystemPriorityCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerSystemPriorityCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerSystemPriorityCounter) Marshal() marshalPatternFlowLacpduPartnerSystemPriorityCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerSystemPriorityCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerSystemPriorityCounter) Unmarshal() unMarshalPatternFlowLacpduPartnerSystemPriorityCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerSystemPriorityCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerSystemPriorityCounter) ToProto() (*otg.PatternFlowLacpduPartnerSystemPriorityCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerSystemPriorityCounter) FromProto(msg *otg.PatternFlowLacpduPartnerSystemPriorityCounter) (PatternFlowLacpduPartnerSystemPriorityCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerSystemPriorityCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerSystemPriorityCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerSystemPriorityCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerSystemPriorityCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerSystemPriorityCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerSystemPriorityCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerSystemPriorityCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerSystemPriorityCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerSystemPriorityCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerSystemPriorityCounter) Clone() (PatternFlowLacpduPartnerSystemPriorityCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerSystemPriorityCounter()
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

// PatternFlowLacpduPartnerSystemPriorityCounter is integer counter pattern
type PatternFlowLacpduPartnerSystemPriorityCounter interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerSystemPriorityCounter to protobuf object *otg.PatternFlowLacpduPartnerSystemPriorityCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerSystemPriorityCounter
	// setMsg unmarshals PatternFlowLacpduPartnerSystemPriorityCounter from protobuf object *otg.PatternFlowLacpduPartnerSystemPriorityCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerSystemPriorityCounter) PatternFlowLacpduPartnerSystemPriorityCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerSystemPriorityCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerSystemPriorityCounter
	// validate validates PatternFlowLacpduPartnerSystemPriorityCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerSystemPriorityCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduPartnerSystemPriorityCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduPartnerSystemPriorityCounter
	SetStart(value uint32) PatternFlowLacpduPartnerSystemPriorityCounter
	// HasStart checks if Start has been set in PatternFlowLacpduPartnerSystemPriorityCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduPartnerSystemPriorityCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduPartnerSystemPriorityCounter
	SetStep(value uint32) PatternFlowLacpduPartnerSystemPriorityCounter
	// HasStep checks if Step has been set in PatternFlowLacpduPartnerSystemPriorityCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduPartnerSystemPriorityCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduPartnerSystemPriorityCounter
	SetCount(value uint32) PatternFlowLacpduPartnerSystemPriorityCounter
	// HasCount checks if Count has been set in PatternFlowLacpduPartnerSystemPriorityCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerSystemPriorityCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerSystemPriorityCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduPartnerSystemPriorityCounter object
func (obj *patternFlowLacpduPartnerSystemPriorityCounter) SetStart(value uint32) PatternFlowLacpduPartnerSystemPriorityCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerSystemPriorityCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerSystemPriorityCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduPartnerSystemPriorityCounter object
func (obj *patternFlowLacpduPartnerSystemPriorityCounter) SetStep(value uint32) PatternFlowLacpduPartnerSystemPriorityCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerSystemPriorityCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerSystemPriorityCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduPartnerSystemPriorityCounter object
func (obj *patternFlowLacpduPartnerSystemPriorityCounter) SetCount(value uint32) PatternFlowLacpduPartnerSystemPriorityCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduPartnerSystemPriorityCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerSystemPriorityCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerSystemPriorityCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerSystemPriorityCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduPartnerSystemPriorityCounter) setDefault() {
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

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerStateDefaultedCounter *****
type patternFlowLacpduPartnerStateDefaultedCounter struct {
	validation
	obj          *otg.PatternFlowLacpduPartnerStateDefaultedCounter
	marshaller   marshalPatternFlowLacpduPartnerStateDefaultedCounter
	unMarshaller unMarshalPatternFlowLacpduPartnerStateDefaultedCounter
}

func NewPatternFlowLacpduPartnerStateDefaultedCounter() PatternFlowLacpduPartnerStateDefaultedCounter {
	obj := patternFlowLacpduPartnerStateDefaultedCounter{obj: &otg.PatternFlowLacpduPartnerStateDefaultedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerStateDefaultedCounter) msg() *otg.PatternFlowLacpduPartnerStateDefaultedCounter {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerStateDefaultedCounter) setMsg(msg *otg.PatternFlowLacpduPartnerStateDefaultedCounter) PatternFlowLacpduPartnerStateDefaultedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerStateDefaultedCounter struct {
	obj *patternFlowLacpduPartnerStateDefaultedCounter
}

type marshalPatternFlowLacpduPartnerStateDefaultedCounter interface {
	// ToProto marshals PatternFlowLacpduPartnerStateDefaultedCounter to protobuf object *otg.PatternFlowLacpduPartnerStateDefaultedCounter
	ToProto() (*otg.PatternFlowLacpduPartnerStateDefaultedCounter, error)
	// ToPbText marshals PatternFlowLacpduPartnerStateDefaultedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerStateDefaultedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerStateDefaultedCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerStateDefaultedCounter struct {
	obj *patternFlowLacpduPartnerStateDefaultedCounter
}

type unMarshalPatternFlowLacpduPartnerStateDefaultedCounter interface {
	// FromProto unmarshals PatternFlowLacpduPartnerStateDefaultedCounter from protobuf object *otg.PatternFlowLacpduPartnerStateDefaultedCounter
	FromProto(msg *otg.PatternFlowLacpduPartnerStateDefaultedCounter) (PatternFlowLacpduPartnerStateDefaultedCounter, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerStateDefaultedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerStateDefaultedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerStateDefaultedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerStateDefaultedCounter) Marshal() marshalPatternFlowLacpduPartnerStateDefaultedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerStateDefaultedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerStateDefaultedCounter) Unmarshal() unMarshalPatternFlowLacpduPartnerStateDefaultedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerStateDefaultedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerStateDefaultedCounter) ToProto() (*otg.PatternFlowLacpduPartnerStateDefaultedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerStateDefaultedCounter) FromProto(msg *otg.PatternFlowLacpduPartnerStateDefaultedCounter) (PatternFlowLacpduPartnerStateDefaultedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerStateDefaultedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateDefaultedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateDefaultedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateDefaultedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateDefaultedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateDefaultedCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerStateDefaultedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateDefaultedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateDefaultedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerStateDefaultedCounter) Clone() (PatternFlowLacpduPartnerStateDefaultedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerStateDefaultedCounter()
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

// PatternFlowLacpduPartnerStateDefaultedCounter is integer counter pattern
type PatternFlowLacpduPartnerStateDefaultedCounter interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerStateDefaultedCounter to protobuf object *otg.PatternFlowLacpduPartnerStateDefaultedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerStateDefaultedCounter
	// setMsg unmarshals PatternFlowLacpduPartnerStateDefaultedCounter from protobuf object *otg.PatternFlowLacpduPartnerStateDefaultedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerStateDefaultedCounter) PatternFlowLacpduPartnerStateDefaultedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerStateDefaultedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerStateDefaultedCounter
	// validate validates PatternFlowLacpduPartnerStateDefaultedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerStateDefaultedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduPartnerStateDefaultedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduPartnerStateDefaultedCounter
	SetStart(value uint32) PatternFlowLacpduPartnerStateDefaultedCounter
	// HasStart checks if Start has been set in PatternFlowLacpduPartnerStateDefaultedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduPartnerStateDefaultedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduPartnerStateDefaultedCounter
	SetStep(value uint32) PatternFlowLacpduPartnerStateDefaultedCounter
	// HasStep checks if Step has been set in PatternFlowLacpduPartnerStateDefaultedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduPartnerStateDefaultedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduPartnerStateDefaultedCounter
	SetCount(value uint32) PatternFlowLacpduPartnerStateDefaultedCounter
	// HasCount checks if Count has been set in PatternFlowLacpduPartnerStateDefaultedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerStateDefaultedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerStateDefaultedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduPartnerStateDefaultedCounter object
func (obj *patternFlowLacpduPartnerStateDefaultedCounter) SetStart(value uint32) PatternFlowLacpduPartnerStateDefaultedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerStateDefaultedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerStateDefaultedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduPartnerStateDefaultedCounter object
func (obj *patternFlowLacpduPartnerStateDefaultedCounter) SetStep(value uint32) PatternFlowLacpduPartnerStateDefaultedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerStateDefaultedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerStateDefaultedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduPartnerStateDefaultedCounter object
func (obj *patternFlowLacpduPartnerStateDefaultedCounter) SetCount(value uint32) PatternFlowLacpduPartnerStateDefaultedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduPartnerStateDefaultedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateDefaultedCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateDefaultedCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateDefaultedCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduPartnerStateDefaultedCounter) setDefault() {
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

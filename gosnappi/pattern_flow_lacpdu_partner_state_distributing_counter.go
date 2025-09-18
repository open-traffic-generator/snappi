package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerStateDistributingCounter *****
type patternFlowLacpduPartnerStateDistributingCounter struct {
	validation
	obj          *otg.PatternFlowLacpduPartnerStateDistributingCounter
	marshaller   marshalPatternFlowLacpduPartnerStateDistributingCounter
	unMarshaller unMarshalPatternFlowLacpduPartnerStateDistributingCounter
}

func NewPatternFlowLacpduPartnerStateDistributingCounter() PatternFlowLacpduPartnerStateDistributingCounter {
	obj := patternFlowLacpduPartnerStateDistributingCounter{obj: &otg.PatternFlowLacpduPartnerStateDistributingCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerStateDistributingCounter) msg() *otg.PatternFlowLacpduPartnerStateDistributingCounter {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerStateDistributingCounter) setMsg(msg *otg.PatternFlowLacpduPartnerStateDistributingCounter) PatternFlowLacpduPartnerStateDistributingCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerStateDistributingCounter struct {
	obj *patternFlowLacpduPartnerStateDistributingCounter
}

type marshalPatternFlowLacpduPartnerStateDistributingCounter interface {
	// ToProto marshals PatternFlowLacpduPartnerStateDistributingCounter to protobuf object *otg.PatternFlowLacpduPartnerStateDistributingCounter
	ToProto() (*otg.PatternFlowLacpduPartnerStateDistributingCounter, error)
	// ToPbText marshals PatternFlowLacpduPartnerStateDistributingCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerStateDistributingCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerStateDistributingCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerStateDistributingCounter struct {
	obj *patternFlowLacpduPartnerStateDistributingCounter
}

type unMarshalPatternFlowLacpduPartnerStateDistributingCounter interface {
	// FromProto unmarshals PatternFlowLacpduPartnerStateDistributingCounter from protobuf object *otg.PatternFlowLacpduPartnerStateDistributingCounter
	FromProto(msg *otg.PatternFlowLacpduPartnerStateDistributingCounter) (PatternFlowLacpduPartnerStateDistributingCounter, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerStateDistributingCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerStateDistributingCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerStateDistributingCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerStateDistributingCounter) Marshal() marshalPatternFlowLacpduPartnerStateDistributingCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerStateDistributingCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerStateDistributingCounter) Unmarshal() unMarshalPatternFlowLacpduPartnerStateDistributingCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerStateDistributingCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerStateDistributingCounter) ToProto() (*otg.PatternFlowLacpduPartnerStateDistributingCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerStateDistributingCounter) FromProto(msg *otg.PatternFlowLacpduPartnerStateDistributingCounter) (PatternFlowLacpduPartnerStateDistributingCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerStateDistributingCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateDistributingCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateDistributingCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateDistributingCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateDistributingCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateDistributingCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerStateDistributingCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateDistributingCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateDistributingCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerStateDistributingCounter) Clone() (PatternFlowLacpduPartnerStateDistributingCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerStateDistributingCounter()
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

// PatternFlowLacpduPartnerStateDistributingCounter is integer counter pattern
type PatternFlowLacpduPartnerStateDistributingCounter interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerStateDistributingCounter to protobuf object *otg.PatternFlowLacpduPartnerStateDistributingCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerStateDistributingCounter
	// setMsg unmarshals PatternFlowLacpduPartnerStateDistributingCounter from protobuf object *otg.PatternFlowLacpduPartnerStateDistributingCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerStateDistributingCounter) PatternFlowLacpduPartnerStateDistributingCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerStateDistributingCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerStateDistributingCounter
	// validate validates PatternFlowLacpduPartnerStateDistributingCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerStateDistributingCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduPartnerStateDistributingCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduPartnerStateDistributingCounter
	SetStart(value uint32) PatternFlowLacpduPartnerStateDistributingCounter
	// HasStart checks if Start has been set in PatternFlowLacpduPartnerStateDistributingCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduPartnerStateDistributingCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduPartnerStateDistributingCounter
	SetStep(value uint32) PatternFlowLacpduPartnerStateDistributingCounter
	// HasStep checks if Step has been set in PatternFlowLacpduPartnerStateDistributingCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduPartnerStateDistributingCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduPartnerStateDistributingCounter
	SetCount(value uint32) PatternFlowLacpduPartnerStateDistributingCounter
	// HasCount checks if Count has been set in PatternFlowLacpduPartnerStateDistributingCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerStateDistributingCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerStateDistributingCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduPartnerStateDistributingCounter object
func (obj *patternFlowLacpduPartnerStateDistributingCounter) SetStart(value uint32) PatternFlowLacpduPartnerStateDistributingCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerStateDistributingCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerStateDistributingCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduPartnerStateDistributingCounter object
func (obj *patternFlowLacpduPartnerStateDistributingCounter) SetStep(value uint32) PatternFlowLacpduPartnerStateDistributingCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerStateDistributingCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerStateDistributingCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduPartnerStateDistributingCounter object
func (obj *patternFlowLacpduPartnerStateDistributingCounter) SetCount(value uint32) PatternFlowLacpduPartnerStateDistributingCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduPartnerStateDistributingCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateDistributingCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateDistributingCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateDistributingCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduPartnerStateDistributingCounter) setDefault() {
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

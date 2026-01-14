package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerStateExpiredCounter *****
type patternFlowLacpduPartnerStateExpiredCounter struct {
	validation
	obj          *otg.PatternFlowLacpduPartnerStateExpiredCounter
	marshaller   marshalPatternFlowLacpduPartnerStateExpiredCounter
	unMarshaller unMarshalPatternFlowLacpduPartnerStateExpiredCounter
}

func NewPatternFlowLacpduPartnerStateExpiredCounter() PatternFlowLacpduPartnerStateExpiredCounter {
	obj := patternFlowLacpduPartnerStateExpiredCounter{obj: &otg.PatternFlowLacpduPartnerStateExpiredCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerStateExpiredCounter) msg() *otg.PatternFlowLacpduPartnerStateExpiredCounter {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerStateExpiredCounter) setMsg(msg *otg.PatternFlowLacpduPartnerStateExpiredCounter) PatternFlowLacpduPartnerStateExpiredCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerStateExpiredCounter struct {
	obj *patternFlowLacpduPartnerStateExpiredCounter
}

type marshalPatternFlowLacpduPartnerStateExpiredCounter interface {
	// ToProto marshals PatternFlowLacpduPartnerStateExpiredCounter to protobuf object *otg.PatternFlowLacpduPartnerStateExpiredCounter
	ToProto() (*otg.PatternFlowLacpduPartnerStateExpiredCounter, error)
	// ToPbText marshals PatternFlowLacpduPartnerStateExpiredCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerStateExpiredCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerStateExpiredCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerStateExpiredCounter struct {
	obj *patternFlowLacpduPartnerStateExpiredCounter
}

type unMarshalPatternFlowLacpduPartnerStateExpiredCounter interface {
	// FromProto unmarshals PatternFlowLacpduPartnerStateExpiredCounter from protobuf object *otg.PatternFlowLacpduPartnerStateExpiredCounter
	FromProto(msg *otg.PatternFlowLacpduPartnerStateExpiredCounter) (PatternFlowLacpduPartnerStateExpiredCounter, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerStateExpiredCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerStateExpiredCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerStateExpiredCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerStateExpiredCounter) Marshal() marshalPatternFlowLacpduPartnerStateExpiredCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerStateExpiredCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerStateExpiredCounter) Unmarshal() unMarshalPatternFlowLacpduPartnerStateExpiredCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerStateExpiredCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerStateExpiredCounter) ToProto() (*otg.PatternFlowLacpduPartnerStateExpiredCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerStateExpiredCounter) FromProto(msg *otg.PatternFlowLacpduPartnerStateExpiredCounter) (PatternFlowLacpduPartnerStateExpiredCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerStateExpiredCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateExpiredCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateExpiredCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateExpiredCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateExpiredCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateExpiredCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerStateExpiredCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateExpiredCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateExpiredCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerStateExpiredCounter) Clone() (PatternFlowLacpduPartnerStateExpiredCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerStateExpiredCounter()
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

// PatternFlowLacpduPartnerStateExpiredCounter is integer counter pattern
type PatternFlowLacpduPartnerStateExpiredCounter interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerStateExpiredCounter to protobuf object *otg.PatternFlowLacpduPartnerStateExpiredCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerStateExpiredCounter
	// setMsg unmarshals PatternFlowLacpduPartnerStateExpiredCounter from protobuf object *otg.PatternFlowLacpduPartnerStateExpiredCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerStateExpiredCounter) PatternFlowLacpduPartnerStateExpiredCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerStateExpiredCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerStateExpiredCounter
	// validate validates PatternFlowLacpduPartnerStateExpiredCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerStateExpiredCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduPartnerStateExpiredCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduPartnerStateExpiredCounter
	SetStart(value uint32) PatternFlowLacpduPartnerStateExpiredCounter
	// HasStart checks if Start has been set in PatternFlowLacpduPartnerStateExpiredCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduPartnerStateExpiredCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduPartnerStateExpiredCounter
	SetStep(value uint32) PatternFlowLacpduPartnerStateExpiredCounter
	// HasStep checks if Step has been set in PatternFlowLacpduPartnerStateExpiredCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduPartnerStateExpiredCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduPartnerStateExpiredCounter
	SetCount(value uint32) PatternFlowLacpduPartnerStateExpiredCounter
	// HasCount checks if Count has been set in PatternFlowLacpduPartnerStateExpiredCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerStateExpiredCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerStateExpiredCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduPartnerStateExpiredCounter object
func (obj *patternFlowLacpduPartnerStateExpiredCounter) SetStart(value uint32) PatternFlowLacpduPartnerStateExpiredCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerStateExpiredCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerStateExpiredCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduPartnerStateExpiredCounter object
func (obj *patternFlowLacpduPartnerStateExpiredCounter) SetStep(value uint32) PatternFlowLacpduPartnerStateExpiredCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerStateExpiredCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerStateExpiredCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduPartnerStateExpiredCounter object
func (obj *patternFlowLacpduPartnerStateExpiredCounter) SetCount(value uint32) PatternFlowLacpduPartnerStateExpiredCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduPartnerStateExpiredCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateExpiredCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateExpiredCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateExpiredCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduPartnerStateExpiredCounter) setDefault() {
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

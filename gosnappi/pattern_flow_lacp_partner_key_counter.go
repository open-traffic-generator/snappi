package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerKeyCounter *****
type patternFlowLacpPartnerKeyCounter struct {
	validation
	obj          *otg.PatternFlowLacpPartnerKeyCounter
	marshaller   marshalPatternFlowLacpPartnerKeyCounter
	unMarshaller unMarshalPatternFlowLacpPartnerKeyCounter
}

func NewPatternFlowLacpPartnerKeyCounter() PatternFlowLacpPartnerKeyCounter {
	obj := patternFlowLacpPartnerKeyCounter{obj: &otg.PatternFlowLacpPartnerKeyCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerKeyCounter) msg() *otg.PatternFlowLacpPartnerKeyCounter {
	return obj.obj
}

func (obj *patternFlowLacpPartnerKeyCounter) setMsg(msg *otg.PatternFlowLacpPartnerKeyCounter) PatternFlowLacpPartnerKeyCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerKeyCounter struct {
	obj *patternFlowLacpPartnerKeyCounter
}

type marshalPatternFlowLacpPartnerKeyCounter interface {
	// ToProto marshals PatternFlowLacpPartnerKeyCounter to protobuf object *otg.PatternFlowLacpPartnerKeyCounter
	ToProto() (*otg.PatternFlowLacpPartnerKeyCounter, error)
	// ToPbText marshals PatternFlowLacpPartnerKeyCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerKeyCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerKeyCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerKeyCounter struct {
	obj *patternFlowLacpPartnerKeyCounter
}

type unMarshalPatternFlowLacpPartnerKeyCounter interface {
	// FromProto unmarshals PatternFlowLacpPartnerKeyCounter from protobuf object *otg.PatternFlowLacpPartnerKeyCounter
	FromProto(msg *otg.PatternFlowLacpPartnerKeyCounter) (PatternFlowLacpPartnerKeyCounter, error)
	// FromPbText unmarshals PatternFlowLacpPartnerKeyCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerKeyCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerKeyCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerKeyCounter) Marshal() marshalPatternFlowLacpPartnerKeyCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerKeyCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerKeyCounter) Unmarshal() unMarshalPatternFlowLacpPartnerKeyCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerKeyCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerKeyCounter) ToProto() (*otg.PatternFlowLacpPartnerKeyCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerKeyCounter) FromProto(msg *otg.PatternFlowLacpPartnerKeyCounter) (PatternFlowLacpPartnerKeyCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerKeyCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerKeyCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerKeyCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerKeyCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerKeyCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerKeyCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerKeyCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerKeyCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerKeyCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerKeyCounter) Clone() (PatternFlowLacpPartnerKeyCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerKeyCounter()
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

// PatternFlowLacpPartnerKeyCounter is integer counter pattern
type PatternFlowLacpPartnerKeyCounter interface {
	Validation
	// msg marshals PatternFlowLacpPartnerKeyCounter to protobuf object *otg.PatternFlowLacpPartnerKeyCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerKeyCounter
	// setMsg unmarshals PatternFlowLacpPartnerKeyCounter from protobuf object *otg.PatternFlowLacpPartnerKeyCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerKeyCounter) PatternFlowLacpPartnerKeyCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerKeyCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerKeyCounter
	// validate validates PatternFlowLacpPartnerKeyCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerKeyCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpPartnerKeyCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpPartnerKeyCounter
	SetStart(value uint32) PatternFlowLacpPartnerKeyCounter
	// HasStart checks if Start has been set in PatternFlowLacpPartnerKeyCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpPartnerKeyCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpPartnerKeyCounter
	SetStep(value uint32) PatternFlowLacpPartnerKeyCounter
	// HasStep checks if Step has been set in PatternFlowLacpPartnerKeyCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpPartnerKeyCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpPartnerKeyCounter
	SetCount(value uint32) PatternFlowLacpPartnerKeyCounter
	// HasCount checks if Count has been set in PatternFlowLacpPartnerKeyCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerKeyCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerKeyCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpPartnerKeyCounter object
func (obj *patternFlowLacpPartnerKeyCounter) SetStart(value uint32) PatternFlowLacpPartnerKeyCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerKeyCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerKeyCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpPartnerKeyCounter object
func (obj *patternFlowLacpPartnerKeyCounter) SetStep(value uint32) PatternFlowLacpPartnerKeyCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerKeyCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerKeyCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpPartnerKeyCounter object
func (obj *patternFlowLacpPartnerKeyCounter) SetCount(value uint32) PatternFlowLacpPartnerKeyCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpPartnerKeyCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerKeyCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerKeyCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerKeyCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpPartnerKeyCounter) setDefault() {
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

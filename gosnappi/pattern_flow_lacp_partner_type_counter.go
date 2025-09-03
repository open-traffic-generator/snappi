package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerTypeCounter *****
type patternFlowLacpPartnerTypeCounter struct {
	validation
	obj          *otg.PatternFlowLacpPartnerTypeCounter
	marshaller   marshalPatternFlowLacpPartnerTypeCounter
	unMarshaller unMarshalPatternFlowLacpPartnerTypeCounter
}

func NewPatternFlowLacpPartnerTypeCounter() PatternFlowLacpPartnerTypeCounter {
	obj := patternFlowLacpPartnerTypeCounter{obj: &otg.PatternFlowLacpPartnerTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerTypeCounter) msg() *otg.PatternFlowLacpPartnerTypeCounter {
	return obj.obj
}

func (obj *patternFlowLacpPartnerTypeCounter) setMsg(msg *otg.PatternFlowLacpPartnerTypeCounter) PatternFlowLacpPartnerTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerTypeCounter struct {
	obj *patternFlowLacpPartnerTypeCounter
}

type marshalPatternFlowLacpPartnerTypeCounter interface {
	// ToProto marshals PatternFlowLacpPartnerTypeCounter to protobuf object *otg.PatternFlowLacpPartnerTypeCounter
	ToProto() (*otg.PatternFlowLacpPartnerTypeCounter, error)
	// ToPbText marshals PatternFlowLacpPartnerTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerTypeCounter struct {
	obj *patternFlowLacpPartnerTypeCounter
}

type unMarshalPatternFlowLacpPartnerTypeCounter interface {
	// FromProto unmarshals PatternFlowLacpPartnerTypeCounter from protobuf object *otg.PatternFlowLacpPartnerTypeCounter
	FromProto(msg *otg.PatternFlowLacpPartnerTypeCounter) (PatternFlowLacpPartnerTypeCounter, error)
	// FromPbText unmarshals PatternFlowLacpPartnerTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerTypeCounter) Marshal() marshalPatternFlowLacpPartnerTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerTypeCounter) Unmarshal() unMarshalPatternFlowLacpPartnerTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerTypeCounter) ToProto() (*otg.PatternFlowLacpPartnerTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerTypeCounter) FromProto(msg *otg.PatternFlowLacpPartnerTypeCounter) (PatternFlowLacpPartnerTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerTypeCounter) Clone() (PatternFlowLacpPartnerTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerTypeCounter()
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

// PatternFlowLacpPartnerTypeCounter is integer counter pattern
type PatternFlowLacpPartnerTypeCounter interface {
	Validation
	// msg marshals PatternFlowLacpPartnerTypeCounter to protobuf object *otg.PatternFlowLacpPartnerTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerTypeCounter
	// setMsg unmarshals PatternFlowLacpPartnerTypeCounter from protobuf object *otg.PatternFlowLacpPartnerTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerTypeCounter) PatternFlowLacpPartnerTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerTypeCounter
	// validate validates PatternFlowLacpPartnerTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpPartnerTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpPartnerTypeCounter
	SetStart(value uint32) PatternFlowLacpPartnerTypeCounter
	// HasStart checks if Start has been set in PatternFlowLacpPartnerTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpPartnerTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpPartnerTypeCounter
	SetStep(value uint32) PatternFlowLacpPartnerTypeCounter
	// HasStep checks if Step has been set in PatternFlowLacpPartnerTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpPartnerTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpPartnerTypeCounter
	SetCount(value uint32) PatternFlowLacpPartnerTypeCounter
	// HasCount checks if Count has been set in PatternFlowLacpPartnerTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpPartnerTypeCounter object
func (obj *patternFlowLacpPartnerTypeCounter) SetStart(value uint32) PatternFlowLacpPartnerTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpPartnerTypeCounter object
func (obj *patternFlowLacpPartnerTypeCounter) SetStep(value uint32) PatternFlowLacpPartnerTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpPartnerTypeCounter object
func (obj *patternFlowLacpPartnerTypeCounter) SetCount(value uint32) PatternFlowLacpPartnerTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpPartnerTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerTypeCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpPartnerTypeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(2)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

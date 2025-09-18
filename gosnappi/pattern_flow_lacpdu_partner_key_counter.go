package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerKeyCounter *****
type patternFlowLacpduPartnerKeyCounter struct {
	validation
	obj          *otg.PatternFlowLacpduPartnerKeyCounter
	marshaller   marshalPatternFlowLacpduPartnerKeyCounter
	unMarshaller unMarshalPatternFlowLacpduPartnerKeyCounter
}

func NewPatternFlowLacpduPartnerKeyCounter() PatternFlowLacpduPartnerKeyCounter {
	obj := patternFlowLacpduPartnerKeyCounter{obj: &otg.PatternFlowLacpduPartnerKeyCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerKeyCounter) msg() *otg.PatternFlowLacpduPartnerKeyCounter {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerKeyCounter) setMsg(msg *otg.PatternFlowLacpduPartnerKeyCounter) PatternFlowLacpduPartnerKeyCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerKeyCounter struct {
	obj *patternFlowLacpduPartnerKeyCounter
}

type marshalPatternFlowLacpduPartnerKeyCounter interface {
	// ToProto marshals PatternFlowLacpduPartnerKeyCounter to protobuf object *otg.PatternFlowLacpduPartnerKeyCounter
	ToProto() (*otg.PatternFlowLacpduPartnerKeyCounter, error)
	// ToPbText marshals PatternFlowLacpduPartnerKeyCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerKeyCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerKeyCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerKeyCounter struct {
	obj *patternFlowLacpduPartnerKeyCounter
}

type unMarshalPatternFlowLacpduPartnerKeyCounter interface {
	// FromProto unmarshals PatternFlowLacpduPartnerKeyCounter from protobuf object *otg.PatternFlowLacpduPartnerKeyCounter
	FromProto(msg *otg.PatternFlowLacpduPartnerKeyCounter) (PatternFlowLacpduPartnerKeyCounter, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerKeyCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerKeyCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerKeyCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerKeyCounter) Marshal() marshalPatternFlowLacpduPartnerKeyCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerKeyCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerKeyCounter) Unmarshal() unMarshalPatternFlowLacpduPartnerKeyCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerKeyCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerKeyCounter) ToProto() (*otg.PatternFlowLacpduPartnerKeyCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerKeyCounter) FromProto(msg *otg.PatternFlowLacpduPartnerKeyCounter) (PatternFlowLacpduPartnerKeyCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerKeyCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerKeyCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerKeyCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerKeyCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerKeyCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerKeyCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerKeyCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerKeyCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerKeyCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerKeyCounter) Clone() (PatternFlowLacpduPartnerKeyCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerKeyCounter()
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

// PatternFlowLacpduPartnerKeyCounter is integer counter pattern
type PatternFlowLacpduPartnerKeyCounter interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerKeyCounter to protobuf object *otg.PatternFlowLacpduPartnerKeyCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerKeyCounter
	// setMsg unmarshals PatternFlowLacpduPartnerKeyCounter from protobuf object *otg.PatternFlowLacpduPartnerKeyCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerKeyCounter) PatternFlowLacpduPartnerKeyCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerKeyCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerKeyCounter
	// validate validates PatternFlowLacpduPartnerKeyCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerKeyCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduPartnerKeyCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduPartnerKeyCounter
	SetStart(value uint32) PatternFlowLacpduPartnerKeyCounter
	// HasStart checks if Start has been set in PatternFlowLacpduPartnerKeyCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduPartnerKeyCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduPartnerKeyCounter
	SetStep(value uint32) PatternFlowLacpduPartnerKeyCounter
	// HasStep checks if Step has been set in PatternFlowLacpduPartnerKeyCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduPartnerKeyCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduPartnerKeyCounter
	SetCount(value uint32) PatternFlowLacpduPartnerKeyCounter
	// HasCount checks if Count has been set in PatternFlowLacpduPartnerKeyCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerKeyCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerKeyCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduPartnerKeyCounter object
func (obj *patternFlowLacpduPartnerKeyCounter) SetStart(value uint32) PatternFlowLacpduPartnerKeyCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerKeyCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerKeyCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduPartnerKeyCounter object
func (obj *patternFlowLacpduPartnerKeyCounter) SetStep(value uint32) PatternFlowLacpduPartnerKeyCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerKeyCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerKeyCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduPartnerKeyCounter object
func (obj *patternFlowLacpduPartnerKeyCounter) SetCount(value uint32) PatternFlowLacpduPartnerKeyCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduPartnerKeyCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerKeyCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerKeyCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerKeyCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduPartnerKeyCounter) setDefault() {
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

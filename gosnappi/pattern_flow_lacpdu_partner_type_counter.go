package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerTypeCounter *****
type patternFlowLacpduPartnerTypeCounter struct {
	validation
	obj          *otg.PatternFlowLacpduPartnerTypeCounter
	marshaller   marshalPatternFlowLacpduPartnerTypeCounter
	unMarshaller unMarshalPatternFlowLacpduPartnerTypeCounter
}

func NewPatternFlowLacpduPartnerTypeCounter() PatternFlowLacpduPartnerTypeCounter {
	obj := patternFlowLacpduPartnerTypeCounter{obj: &otg.PatternFlowLacpduPartnerTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerTypeCounter) msg() *otg.PatternFlowLacpduPartnerTypeCounter {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerTypeCounter) setMsg(msg *otg.PatternFlowLacpduPartnerTypeCounter) PatternFlowLacpduPartnerTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerTypeCounter struct {
	obj *patternFlowLacpduPartnerTypeCounter
}

type marshalPatternFlowLacpduPartnerTypeCounter interface {
	// ToProto marshals PatternFlowLacpduPartnerTypeCounter to protobuf object *otg.PatternFlowLacpduPartnerTypeCounter
	ToProto() (*otg.PatternFlowLacpduPartnerTypeCounter, error)
	// ToPbText marshals PatternFlowLacpduPartnerTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerTypeCounter struct {
	obj *patternFlowLacpduPartnerTypeCounter
}

type unMarshalPatternFlowLacpduPartnerTypeCounter interface {
	// FromProto unmarshals PatternFlowLacpduPartnerTypeCounter from protobuf object *otg.PatternFlowLacpduPartnerTypeCounter
	FromProto(msg *otg.PatternFlowLacpduPartnerTypeCounter) (PatternFlowLacpduPartnerTypeCounter, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerTypeCounter) Marshal() marshalPatternFlowLacpduPartnerTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerTypeCounter) Unmarshal() unMarshalPatternFlowLacpduPartnerTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerTypeCounter) ToProto() (*otg.PatternFlowLacpduPartnerTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerTypeCounter) FromProto(msg *otg.PatternFlowLacpduPartnerTypeCounter) (PatternFlowLacpduPartnerTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerTypeCounter) Clone() (PatternFlowLacpduPartnerTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerTypeCounter()
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

// PatternFlowLacpduPartnerTypeCounter is integer counter pattern
type PatternFlowLacpduPartnerTypeCounter interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerTypeCounter to protobuf object *otg.PatternFlowLacpduPartnerTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerTypeCounter
	// setMsg unmarshals PatternFlowLacpduPartnerTypeCounter from protobuf object *otg.PatternFlowLacpduPartnerTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerTypeCounter) PatternFlowLacpduPartnerTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerTypeCounter
	// validate validates PatternFlowLacpduPartnerTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduPartnerTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduPartnerTypeCounter
	SetStart(value uint32) PatternFlowLacpduPartnerTypeCounter
	// HasStart checks if Start has been set in PatternFlowLacpduPartnerTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduPartnerTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduPartnerTypeCounter
	SetStep(value uint32) PatternFlowLacpduPartnerTypeCounter
	// HasStep checks if Step has been set in PatternFlowLacpduPartnerTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduPartnerTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduPartnerTypeCounter
	SetCount(value uint32) PatternFlowLacpduPartnerTypeCounter
	// HasCount checks if Count has been set in PatternFlowLacpduPartnerTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduPartnerTypeCounter object
func (obj *patternFlowLacpduPartnerTypeCounter) SetStart(value uint32) PatternFlowLacpduPartnerTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduPartnerTypeCounter object
func (obj *patternFlowLacpduPartnerTypeCounter) SetStep(value uint32) PatternFlowLacpduPartnerTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduPartnerTypeCounter object
func (obj *patternFlowLacpduPartnerTypeCounter) SetCount(value uint32) PatternFlowLacpduPartnerTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduPartnerTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerTypeCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduPartnerTypeCounter) setDefault() {
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

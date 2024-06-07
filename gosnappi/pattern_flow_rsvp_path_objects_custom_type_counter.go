package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathObjectsCustomTypeCounter *****
type patternFlowRSVPPathObjectsCustomTypeCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathObjectsCustomTypeCounter
	marshaller   marshalPatternFlowRSVPPathObjectsCustomTypeCounter
	unMarshaller unMarshalPatternFlowRSVPPathObjectsCustomTypeCounter
}

func NewPatternFlowRSVPPathObjectsCustomTypeCounter() PatternFlowRSVPPathObjectsCustomTypeCounter {
	obj := patternFlowRSVPPathObjectsCustomTypeCounter{obj: &otg.PatternFlowRSVPPathObjectsCustomTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathObjectsCustomTypeCounter) msg() *otg.PatternFlowRSVPPathObjectsCustomTypeCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathObjectsCustomTypeCounter) setMsg(msg *otg.PatternFlowRSVPPathObjectsCustomTypeCounter) PatternFlowRSVPPathObjectsCustomTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathObjectsCustomTypeCounter struct {
	obj *patternFlowRSVPPathObjectsCustomTypeCounter
}

type marshalPatternFlowRSVPPathObjectsCustomTypeCounter interface {
	// ToProto marshals PatternFlowRSVPPathObjectsCustomTypeCounter to protobuf object *otg.PatternFlowRSVPPathObjectsCustomTypeCounter
	ToProto() (*otg.PatternFlowRSVPPathObjectsCustomTypeCounter, error)
	// ToPbText marshals PatternFlowRSVPPathObjectsCustomTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathObjectsCustomTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathObjectsCustomTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathObjectsCustomTypeCounter struct {
	obj *patternFlowRSVPPathObjectsCustomTypeCounter
}

type unMarshalPatternFlowRSVPPathObjectsCustomTypeCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathObjectsCustomTypeCounter from protobuf object *otg.PatternFlowRSVPPathObjectsCustomTypeCounter
	FromProto(msg *otg.PatternFlowRSVPPathObjectsCustomTypeCounter) (PatternFlowRSVPPathObjectsCustomTypeCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathObjectsCustomTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathObjectsCustomTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathObjectsCustomTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathObjectsCustomTypeCounter) Marshal() marshalPatternFlowRSVPPathObjectsCustomTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathObjectsCustomTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathObjectsCustomTypeCounter) Unmarshal() unMarshalPatternFlowRSVPPathObjectsCustomTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathObjectsCustomTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathObjectsCustomTypeCounter) ToProto() (*otg.PatternFlowRSVPPathObjectsCustomTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathObjectsCustomTypeCounter) FromProto(msg *otg.PatternFlowRSVPPathObjectsCustomTypeCounter) (PatternFlowRSVPPathObjectsCustomTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathObjectsCustomTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathObjectsCustomTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathObjectsCustomTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathObjectsCustomTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathObjectsCustomTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathObjectsCustomTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathObjectsCustomTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathObjectsCustomTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathObjectsCustomTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathObjectsCustomTypeCounter) Clone() (PatternFlowRSVPPathObjectsCustomTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathObjectsCustomTypeCounter()
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

// PatternFlowRSVPPathObjectsCustomTypeCounter is integer counter pattern
type PatternFlowRSVPPathObjectsCustomTypeCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathObjectsCustomTypeCounter to protobuf object *otg.PatternFlowRSVPPathObjectsCustomTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathObjectsCustomTypeCounter
	// setMsg unmarshals PatternFlowRSVPPathObjectsCustomTypeCounter from protobuf object *otg.PatternFlowRSVPPathObjectsCustomTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathObjectsCustomTypeCounter) PatternFlowRSVPPathObjectsCustomTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathObjectsCustomTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathObjectsCustomTypeCounter
	// validate validates PatternFlowRSVPPathObjectsCustomTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathObjectsCustomTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathObjectsCustomTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathObjectsCustomTypeCounter
	SetStart(value uint32) PatternFlowRSVPPathObjectsCustomTypeCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathObjectsCustomTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathObjectsCustomTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathObjectsCustomTypeCounter
	SetStep(value uint32) PatternFlowRSVPPathObjectsCustomTypeCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathObjectsCustomTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathObjectsCustomTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathObjectsCustomTypeCounter
	SetCount(value uint32) PatternFlowRSVPPathObjectsCustomTypeCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathObjectsCustomTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathObjectsCustomTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathObjectsCustomTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathObjectsCustomTypeCounter object
func (obj *patternFlowRSVPPathObjectsCustomTypeCounter) SetStart(value uint32) PatternFlowRSVPPathObjectsCustomTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathObjectsCustomTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathObjectsCustomTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathObjectsCustomTypeCounter object
func (obj *patternFlowRSVPPathObjectsCustomTypeCounter) SetStep(value uint32) PatternFlowRSVPPathObjectsCustomTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathObjectsCustomTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathObjectsCustomTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathObjectsCustomTypeCounter object
func (obj *patternFlowRSVPPathObjectsCustomTypeCounter) SetCount(value uint32) PatternFlowRSVPPathObjectsCustomTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathObjectsCustomTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathObjectsCustomTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathObjectsCustomTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathObjectsCustomTypeCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRSVPPathObjectsCustomTypeCounter) setDefault() {
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

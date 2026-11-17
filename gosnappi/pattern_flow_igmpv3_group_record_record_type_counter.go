package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3GroupRecordRecordTypeCounter *****
type patternFlowIgmpv3GroupRecordRecordTypeCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv3GroupRecordRecordTypeCounter
	marshaller   marshalPatternFlowIgmpv3GroupRecordRecordTypeCounter
	unMarshaller unMarshalPatternFlowIgmpv3GroupRecordRecordTypeCounter
}

func NewPatternFlowIgmpv3GroupRecordRecordTypeCounter() PatternFlowIgmpv3GroupRecordRecordTypeCounter {
	obj := patternFlowIgmpv3GroupRecordRecordTypeCounter{obj: &otg.PatternFlowIgmpv3GroupRecordRecordTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3GroupRecordRecordTypeCounter) msg() *otg.PatternFlowIgmpv3GroupRecordRecordTypeCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv3GroupRecordRecordTypeCounter) setMsg(msg *otg.PatternFlowIgmpv3GroupRecordRecordTypeCounter) PatternFlowIgmpv3GroupRecordRecordTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3GroupRecordRecordTypeCounter struct {
	obj *patternFlowIgmpv3GroupRecordRecordTypeCounter
}

type marshalPatternFlowIgmpv3GroupRecordRecordTypeCounter interface {
	// ToProto marshals PatternFlowIgmpv3GroupRecordRecordTypeCounter to protobuf object *otg.PatternFlowIgmpv3GroupRecordRecordTypeCounter
	ToProto() (*otg.PatternFlowIgmpv3GroupRecordRecordTypeCounter, error)
	// ToPbText marshals PatternFlowIgmpv3GroupRecordRecordTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3GroupRecordRecordTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3GroupRecordRecordTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3GroupRecordRecordTypeCounter struct {
	obj *patternFlowIgmpv3GroupRecordRecordTypeCounter
}

type unMarshalPatternFlowIgmpv3GroupRecordRecordTypeCounter interface {
	// FromProto unmarshals PatternFlowIgmpv3GroupRecordRecordTypeCounter from protobuf object *otg.PatternFlowIgmpv3GroupRecordRecordTypeCounter
	FromProto(msg *otg.PatternFlowIgmpv3GroupRecordRecordTypeCounter) (PatternFlowIgmpv3GroupRecordRecordTypeCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv3GroupRecordRecordTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3GroupRecordRecordTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3GroupRecordRecordTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3GroupRecordRecordTypeCounter) Marshal() marshalPatternFlowIgmpv3GroupRecordRecordTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3GroupRecordRecordTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3GroupRecordRecordTypeCounter) Unmarshal() unMarshalPatternFlowIgmpv3GroupRecordRecordTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3GroupRecordRecordTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3GroupRecordRecordTypeCounter) ToProto() (*otg.PatternFlowIgmpv3GroupRecordRecordTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3GroupRecordRecordTypeCounter) FromProto(msg *otg.PatternFlowIgmpv3GroupRecordRecordTypeCounter) (PatternFlowIgmpv3GroupRecordRecordTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3GroupRecordRecordTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordRecordTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3GroupRecordRecordTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordRecordTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3GroupRecordRecordTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordRecordTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3GroupRecordRecordTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3GroupRecordRecordTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3GroupRecordRecordTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3GroupRecordRecordTypeCounter) Clone() (PatternFlowIgmpv3GroupRecordRecordTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3GroupRecordRecordTypeCounter()
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

// PatternFlowIgmpv3GroupRecordRecordTypeCounter is integer counter pattern
type PatternFlowIgmpv3GroupRecordRecordTypeCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv3GroupRecordRecordTypeCounter to protobuf object *otg.PatternFlowIgmpv3GroupRecordRecordTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3GroupRecordRecordTypeCounter
	// setMsg unmarshals PatternFlowIgmpv3GroupRecordRecordTypeCounter from protobuf object *otg.PatternFlowIgmpv3GroupRecordRecordTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3GroupRecordRecordTypeCounter) PatternFlowIgmpv3GroupRecordRecordTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3GroupRecordRecordTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3GroupRecordRecordTypeCounter
	// validate validates PatternFlowIgmpv3GroupRecordRecordTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3GroupRecordRecordTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIgmpv3GroupRecordRecordTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIgmpv3GroupRecordRecordTypeCounter
	SetStart(value uint32) PatternFlowIgmpv3GroupRecordRecordTypeCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv3GroupRecordRecordTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIgmpv3GroupRecordRecordTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIgmpv3GroupRecordRecordTypeCounter
	SetStep(value uint32) PatternFlowIgmpv3GroupRecordRecordTypeCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv3GroupRecordRecordTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv3GroupRecordRecordTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv3GroupRecordRecordTypeCounter
	SetCount(value uint32) PatternFlowIgmpv3GroupRecordRecordTypeCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv3GroupRecordRecordTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3GroupRecordRecordTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3GroupRecordRecordTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIgmpv3GroupRecordRecordTypeCounter object
func (obj *patternFlowIgmpv3GroupRecordRecordTypeCounter) SetStart(value uint32) PatternFlowIgmpv3GroupRecordRecordTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3GroupRecordRecordTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3GroupRecordRecordTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIgmpv3GroupRecordRecordTypeCounter object
func (obj *patternFlowIgmpv3GroupRecordRecordTypeCounter) SetStep(value uint32) PatternFlowIgmpv3GroupRecordRecordTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3GroupRecordRecordTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3GroupRecordRecordTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv3GroupRecordRecordTypeCounter object
func (obj *patternFlowIgmpv3GroupRecordRecordTypeCounter) SetCount(value uint32) PatternFlowIgmpv3GroupRecordRecordTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv3GroupRecordRecordTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3GroupRecordRecordTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3GroupRecordRecordTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3GroupRecordRecordTypeCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIgmpv3GroupRecordRecordTypeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(1)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

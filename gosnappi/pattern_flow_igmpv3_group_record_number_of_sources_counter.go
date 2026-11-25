package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter *****
type patternFlowIgmpv3GroupRecordNumberOfSourcesCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	marshaller   marshalPatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	unMarshaller unMarshalPatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
}

func NewPatternFlowIgmpv3GroupRecordNumberOfSourcesCounter() PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter {
	obj := patternFlowIgmpv3GroupRecordNumberOfSourcesCounter{obj: &otg.PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3GroupRecordNumberOfSourcesCounter) msg() *otg.PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv3GroupRecordNumberOfSourcesCounter) setMsg(msg *otg.PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter) PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3GroupRecordNumberOfSourcesCounter struct {
	obj *patternFlowIgmpv3GroupRecordNumberOfSourcesCounter
}

type marshalPatternFlowIgmpv3GroupRecordNumberOfSourcesCounter interface {
	// ToProto marshals PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter to protobuf object *otg.PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	ToProto() (*otg.PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter, error)
	// ToPbText marshals PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3GroupRecordNumberOfSourcesCounter struct {
	obj *patternFlowIgmpv3GroupRecordNumberOfSourcesCounter
}

type unMarshalPatternFlowIgmpv3GroupRecordNumberOfSourcesCounter interface {
	// FromProto unmarshals PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter from protobuf object *otg.PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	FromProto(msg *otg.PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter) (PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3GroupRecordNumberOfSourcesCounter) Marshal() marshalPatternFlowIgmpv3GroupRecordNumberOfSourcesCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3GroupRecordNumberOfSourcesCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3GroupRecordNumberOfSourcesCounter) Unmarshal() unMarshalPatternFlowIgmpv3GroupRecordNumberOfSourcesCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3GroupRecordNumberOfSourcesCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3GroupRecordNumberOfSourcesCounter) ToProto() (*otg.PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3GroupRecordNumberOfSourcesCounter) FromProto(msg *otg.PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter) (PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3GroupRecordNumberOfSourcesCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordNumberOfSourcesCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3GroupRecordNumberOfSourcesCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordNumberOfSourcesCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3GroupRecordNumberOfSourcesCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordNumberOfSourcesCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3GroupRecordNumberOfSourcesCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3GroupRecordNumberOfSourcesCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3GroupRecordNumberOfSourcesCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3GroupRecordNumberOfSourcesCounter) Clone() (PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3GroupRecordNumberOfSourcesCounter()
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

// PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter is integer counter pattern
type PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter to protobuf object *otg.PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	// setMsg unmarshals PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter from protobuf object *otg.PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter) PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	// validate validates PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	SetStart(value uint32) PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	SetStep(value uint32) PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	SetCount(value uint32) PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3GroupRecordNumberOfSourcesCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3GroupRecordNumberOfSourcesCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter object
func (obj *patternFlowIgmpv3GroupRecordNumberOfSourcesCounter) SetStart(value uint32) PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3GroupRecordNumberOfSourcesCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3GroupRecordNumberOfSourcesCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter object
func (obj *patternFlowIgmpv3GroupRecordNumberOfSourcesCounter) SetStep(value uint32) PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3GroupRecordNumberOfSourcesCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3GroupRecordNumberOfSourcesCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter object
func (obj *patternFlowIgmpv3GroupRecordNumberOfSourcesCounter) SetCount(value uint32) PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv3GroupRecordNumberOfSourcesCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIgmpv3GroupRecordNumberOfSourcesCounter) setDefault() {
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

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPauseEtherTypeCounter *****
type patternFlowPfcPauseEtherTypeCounter struct {
	validation
	obj          *otg.PatternFlowPfcPauseEtherTypeCounter
	marshaller   marshalPatternFlowPfcPauseEtherTypeCounter
	unMarshaller unMarshalPatternFlowPfcPauseEtherTypeCounter
}

func NewPatternFlowPfcPauseEtherTypeCounter() PatternFlowPfcPauseEtherTypeCounter {
	obj := patternFlowPfcPauseEtherTypeCounter{obj: &otg.PatternFlowPfcPauseEtherTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPauseEtherTypeCounter) msg() *otg.PatternFlowPfcPauseEtherTypeCounter {
	return obj.obj
}

func (obj *patternFlowPfcPauseEtherTypeCounter) setMsg(msg *otg.PatternFlowPfcPauseEtherTypeCounter) PatternFlowPfcPauseEtherTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPauseEtherTypeCounter struct {
	obj *patternFlowPfcPauseEtherTypeCounter
}

type marshalPatternFlowPfcPauseEtherTypeCounter interface {
	// ToProto marshals PatternFlowPfcPauseEtherTypeCounter to protobuf object *otg.PatternFlowPfcPauseEtherTypeCounter
	ToProto() (*otg.PatternFlowPfcPauseEtherTypeCounter, error)
	// ToPbText marshals PatternFlowPfcPauseEtherTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPauseEtherTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPauseEtherTypeCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowPfcPauseEtherTypeCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowPfcPauseEtherTypeCounter struct {
	obj *patternFlowPfcPauseEtherTypeCounter
}

type unMarshalPatternFlowPfcPauseEtherTypeCounter interface {
	// FromProto unmarshals PatternFlowPfcPauseEtherTypeCounter from protobuf object *otg.PatternFlowPfcPauseEtherTypeCounter
	FromProto(msg *otg.PatternFlowPfcPauseEtherTypeCounter) (PatternFlowPfcPauseEtherTypeCounter, error)
	// FromPbText unmarshals PatternFlowPfcPauseEtherTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPauseEtherTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPauseEtherTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPauseEtherTypeCounter) Marshal() marshalPatternFlowPfcPauseEtherTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPauseEtherTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPauseEtherTypeCounter) Unmarshal() unMarshalPatternFlowPfcPauseEtherTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPauseEtherTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPauseEtherTypeCounter) ToProto() (*otg.PatternFlowPfcPauseEtherTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPauseEtherTypeCounter) FromProto(msg *otg.PatternFlowPfcPauseEtherTypeCounter) (PatternFlowPfcPauseEtherTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPauseEtherTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseEtherTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPauseEtherTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseEtherTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPauseEtherTypeCounter) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalpatternFlowPfcPauseEtherTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseEtherTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowPfcPauseEtherTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseEtherTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseEtherTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPauseEtherTypeCounter) Clone() (PatternFlowPfcPauseEtherTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPauseEtherTypeCounter()
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

// PatternFlowPfcPauseEtherTypeCounter is integer counter pattern
type PatternFlowPfcPauseEtherTypeCounter interface {
	Validation
	// msg marshals PatternFlowPfcPauseEtherTypeCounter to protobuf object *otg.PatternFlowPfcPauseEtherTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPauseEtherTypeCounter
	// setMsg unmarshals PatternFlowPfcPauseEtherTypeCounter from protobuf object *otg.PatternFlowPfcPauseEtherTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPauseEtherTypeCounter) PatternFlowPfcPauseEtherTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPauseEtherTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPauseEtherTypeCounter
	// validate validates PatternFlowPfcPauseEtherTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPauseEtherTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowPfcPauseEtherTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowPfcPauseEtherTypeCounter
	SetStart(value uint32) PatternFlowPfcPauseEtherTypeCounter
	// HasStart checks if Start has been set in PatternFlowPfcPauseEtherTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowPfcPauseEtherTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowPfcPauseEtherTypeCounter
	SetStep(value uint32) PatternFlowPfcPauseEtherTypeCounter
	// HasStep checks if Step has been set in PatternFlowPfcPauseEtherTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowPfcPauseEtherTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowPfcPauseEtherTypeCounter
	SetCount(value uint32) PatternFlowPfcPauseEtherTypeCounter
	// HasCount checks if Count has been set in PatternFlowPfcPauseEtherTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPauseEtherTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPauseEtherTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowPfcPauseEtherTypeCounter object
func (obj *patternFlowPfcPauseEtherTypeCounter) SetStart(value uint32) PatternFlowPfcPauseEtherTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPauseEtherTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPauseEtherTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowPfcPauseEtherTypeCounter object
func (obj *patternFlowPfcPauseEtherTypeCounter) SetStep(value uint32) PatternFlowPfcPauseEtherTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPauseEtherTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPauseEtherTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowPfcPauseEtherTypeCounter object
func (obj *patternFlowPfcPauseEtherTypeCounter) SetCount(value uint32) PatternFlowPfcPauseEtherTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowPfcPauseEtherTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseEtherTypeCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseEtherTypeCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseEtherTypeCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowPfcPauseEtherTypeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(34824)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

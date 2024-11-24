package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVxlanVniCounter *****
type patternFlowVxlanVniCounter struct {
	validation
	obj          *otg.PatternFlowVxlanVniCounter
	marshaller   marshalPatternFlowVxlanVniCounter
	unMarshaller unMarshalPatternFlowVxlanVniCounter
}

func NewPatternFlowVxlanVniCounter() PatternFlowVxlanVniCounter {
	obj := patternFlowVxlanVniCounter{obj: &otg.PatternFlowVxlanVniCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVxlanVniCounter) msg() *otg.PatternFlowVxlanVniCounter {
	return obj.obj
}

func (obj *patternFlowVxlanVniCounter) setMsg(msg *otg.PatternFlowVxlanVniCounter) PatternFlowVxlanVniCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVxlanVniCounter struct {
	obj *patternFlowVxlanVniCounter
}

type marshalPatternFlowVxlanVniCounter interface {
	// ToProto marshals PatternFlowVxlanVniCounter to protobuf object *otg.PatternFlowVxlanVniCounter
	ToProto() (*otg.PatternFlowVxlanVniCounter, error)
	// ToPbText marshals PatternFlowVxlanVniCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVxlanVniCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVxlanVniCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowVxlanVniCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowVxlanVniCounter struct {
	obj *patternFlowVxlanVniCounter
}

type unMarshalPatternFlowVxlanVniCounter interface {
	// FromProto unmarshals PatternFlowVxlanVniCounter from protobuf object *otg.PatternFlowVxlanVniCounter
	FromProto(msg *otg.PatternFlowVxlanVniCounter) (PatternFlowVxlanVniCounter, error)
	// FromPbText unmarshals PatternFlowVxlanVniCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVxlanVniCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVxlanVniCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVxlanVniCounter) Marshal() marshalPatternFlowVxlanVniCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVxlanVniCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVxlanVniCounter) Unmarshal() unMarshalPatternFlowVxlanVniCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVxlanVniCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVxlanVniCounter) ToProto() (*otg.PatternFlowVxlanVniCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVxlanVniCounter) FromProto(msg *otg.PatternFlowVxlanVniCounter) (PatternFlowVxlanVniCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVxlanVniCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVxlanVniCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowVxlanVniCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVxlanVniCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowVxlanVniCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowVxlanVniCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVxlanVniCounter) FromJson(value string) error {
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

func (obj *patternFlowVxlanVniCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVxlanVniCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVxlanVniCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVxlanVniCounter) Clone() (PatternFlowVxlanVniCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVxlanVniCounter()
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

// PatternFlowVxlanVniCounter is integer counter pattern
type PatternFlowVxlanVniCounter interface {
	Validation
	// msg marshals PatternFlowVxlanVniCounter to protobuf object *otg.PatternFlowVxlanVniCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowVxlanVniCounter
	// setMsg unmarshals PatternFlowVxlanVniCounter from protobuf object *otg.PatternFlowVxlanVniCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVxlanVniCounter) PatternFlowVxlanVniCounter
	// provides marshal interface
	Marshal() marshalPatternFlowVxlanVniCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVxlanVniCounter
	// validate validates PatternFlowVxlanVniCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVxlanVniCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowVxlanVniCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowVxlanVniCounter
	SetStart(value uint32) PatternFlowVxlanVniCounter
	// HasStart checks if Start has been set in PatternFlowVxlanVniCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowVxlanVniCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowVxlanVniCounter
	SetStep(value uint32) PatternFlowVxlanVniCounter
	// HasStep checks if Step has been set in PatternFlowVxlanVniCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowVxlanVniCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowVxlanVniCounter
	SetCount(value uint32) PatternFlowVxlanVniCounter
	// HasCount checks if Count has been set in PatternFlowVxlanVniCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowVxlanVniCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowVxlanVniCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowVxlanVniCounter object
func (obj *patternFlowVxlanVniCounter) SetStart(value uint32) PatternFlowVxlanVniCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowVxlanVniCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowVxlanVniCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowVxlanVniCounter object
func (obj *patternFlowVxlanVniCounter) SetStep(value uint32) PatternFlowVxlanVniCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowVxlanVniCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowVxlanVniCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowVxlanVniCounter object
func (obj *patternFlowVxlanVniCounter) SetCount(value uint32) PatternFlowVxlanVniCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowVxlanVniCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVxlanVniCounter.Start <= 16777215 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVxlanVniCounter.Step <= 16777215 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVxlanVniCounter.Count <= 16777215 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowVxlanVniCounter) setDefault() {
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

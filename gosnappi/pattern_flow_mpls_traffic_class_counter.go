package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowMplsTrafficClassCounter *****
type patternFlowMplsTrafficClassCounter struct {
	validation
	obj          *otg.PatternFlowMplsTrafficClassCounter
	marshaller   marshalPatternFlowMplsTrafficClassCounter
	unMarshaller unMarshalPatternFlowMplsTrafficClassCounter
}

func NewPatternFlowMplsTrafficClassCounter() PatternFlowMplsTrafficClassCounter {
	obj := patternFlowMplsTrafficClassCounter{obj: &otg.PatternFlowMplsTrafficClassCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowMplsTrafficClassCounter) msg() *otg.PatternFlowMplsTrafficClassCounter {
	return obj.obj
}

func (obj *patternFlowMplsTrafficClassCounter) setMsg(msg *otg.PatternFlowMplsTrafficClassCounter) PatternFlowMplsTrafficClassCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowMplsTrafficClassCounter struct {
	obj *patternFlowMplsTrafficClassCounter
}

type marshalPatternFlowMplsTrafficClassCounter interface {
	// ToProto marshals PatternFlowMplsTrafficClassCounter to protobuf object *otg.PatternFlowMplsTrafficClassCounter
	ToProto() (*otg.PatternFlowMplsTrafficClassCounter, error)
	// ToPbText marshals PatternFlowMplsTrafficClassCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowMplsTrafficClassCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowMplsTrafficClassCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowMplsTrafficClassCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowMplsTrafficClassCounter struct {
	obj *patternFlowMplsTrafficClassCounter
}

type unMarshalPatternFlowMplsTrafficClassCounter interface {
	// FromProto unmarshals PatternFlowMplsTrafficClassCounter from protobuf object *otg.PatternFlowMplsTrafficClassCounter
	FromProto(msg *otg.PatternFlowMplsTrafficClassCounter) (PatternFlowMplsTrafficClassCounter, error)
	// FromPbText unmarshals PatternFlowMplsTrafficClassCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowMplsTrafficClassCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowMplsTrafficClassCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowMplsTrafficClassCounter) Marshal() marshalPatternFlowMplsTrafficClassCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowMplsTrafficClassCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowMplsTrafficClassCounter) Unmarshal() unMarshalPatternFlowMplsTrafficClassCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowMplsTrafficClassCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowMplsTrafficClassCounter) ToProto() (*otg.PatternFlowMplsTrafficClassCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowMplsTrafficClassCounter) FromProto(msg *otg.PatternFlowMplsTrafficClassCounter) (PatternFlowMplsTrafficClassCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowMplsTrafficClassCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowMplsTrafficClassCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowMplsTrafficClassCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowMplsTrafficClassCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowMplsTrafficClassCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowMplsTrafficClassCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowMplsTrafficClassCounter) FromJson(value string) error {
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

func (obj *patternFlowMplsTrafficClassCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowMplsTrafficClassCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowMplsTrafficClassCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowMplsTrafficClassCounter) Clone() (PatternFlowMplsTrafficClassCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowMplsTrafficClassCounter()
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

// PatternFlowMplsTrafficClassCounter is integer counter pattern
type PatternFlowMplsTrafficClassCounter interface {
	Validation
	// msg marshals PatternFlowMplsTrafficClassCounter to protobuf object *otg.PatternFlowMplsTrafficClassCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowMplsTrafficClassCounter
	// setMsg unmarshals PatternFlowMplsTrafficClassCounter from protobuf object *otg.PatternFlowMplsTrafficClassCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowMplsTrafficClassCounter) PatternFlowMplsTrafficClassCounter
	// provides marshal interface
	Marshal() marshalPatternFlowMplsTrafficClassCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowMplsTrafficClassCounter
	// validate validates PatternFlowMplsTrafficClassCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowMplsTrafficClassCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowMplsTrafficClassCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowMplsTrafficClassCounter
	SetStart(value uint32) PatternFlowMplsTrafficClassCounter
	// HasStart checks if Start has been set in PatternFlowMplsTrafficClassCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowMplsTrafficClassCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowMplsTrafficClassCounter
	SetStep(value uint32) PatternFlowMplsTrafficClassCounter
	// HasStep checks if Step has been set in PatternFlowMplsTrafficClassCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowMplsTrafficClassCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowMplsTrafficClassCounter
	SetCount(value uint32) PatternFlowMplsTrafficClassCounter
	// HasCount checks if Count has been set in PatternFlowMplsTrafficClassCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowMplsTrafficClassCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowMplsTrafficClassCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowMplsTrafficClassCounter object
func (obj *patternFlowMplsTrafficClassCounter) SetStart(value uint32) PatternFlowMplsTrafficClassCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowMplsTrafficClassCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowMplsTrafficClassCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowMplsTrafficClassCounter object
func (obj *patternFlowMplsTrafficClassCounter) SetStep(value uint32) PatternFlowMplsTrafficClassCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowMplsTrafficClassCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowMplsTrafficClassCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowMplsTrafficClassCounter object
func (obj *patternFlowMplsTrafficClassCounter) SetCount(value uint32) PatternFlowMplsTrafficClassCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowMplsTrafficClassCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsTrafficClassCounter.Start <= 7 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsTrafficClassCounter.Step <= 7 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsTrafficClassCounter.Count <= 7 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowMplsTrafficClassCounter) setDefault() {
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

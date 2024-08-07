package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass5Counter *****
type patternFlowPfcPausePauseClass5Counter struct {
	validation
	obj          *otg.PatternFlowPfcPausePauseClass5Counter
	marshaller   marshalPatternFlowPfcPausePauseClass5Counter
	unMarshaller unMarshalPatternFlowPfcPausePauseClass5Counter
}

func NewPatternFlowPfcPausePauseClass5Counter() PatternFlowPfcPausePauseClass5Counter {
	obj := patternFlowPfcPausePauseClass5Counter{obj: &otg.PatternFlowPfcPausePauseClass5Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass5Counter) msg() *otg.PatternFlowPfcPausePauseClass5Counter {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass5Counter) setMsg(msg *otg.PatternFlowPfcPausePauseClass5Counter) PatternFlowPfcPausePauseClass5Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass5Counter struct {
	obj *patternFlowPfcPausePauseClass5Counter
}

type marshalPatternFlowPfcPausePauseClass5Counter interface {
	// ToProto marshals PatternFlowPfcPausePauseClass5Counter to protobuf object *otg.PatternFlowPfcPausePauseClass5Counter
	ToProto() (*otg.PatternFlowPfcPausePauseClass5Counter, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass5Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass5Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass5Counter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass5Counter struct {
	obj *patternFlowPfcPausePauseClass5Counter
}

type unMarshalPatternFlowPfcPausePauseClass5Counter interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass5Counter from protobuf object *otg.PatternFlowPfcPausePauseClass5Counter
	FromProto(msg *otg.PatternFlowPfcPausePauseClass5Counter) (PatternFlowPfcPausePauseClass5Counter, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass5Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass5Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass5Counter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass5Counter) Marshal() marshalPatternFlowPfcPausePauseClass5Counter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass5Counter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass5Counter) Unmarshal() unMarshalPatternFlowPfcPausePauseClass5Counter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass5Counter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass5Counter) ToProto() (*otg.PatternFlowPfcPausePauseClass5Counter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass5Counter) FromProto(msg *otg.PatternFlowPfcPausePauseClass5Counter) (PatternFlowPfcPausePauseClass5Counter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass5Counter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass5Counter) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass5Counter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass5Counter) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass5Counter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass5Counter) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass5Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass5Counter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass5Counter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass5Counter) Clone() (PatternFlowPfcPausePauseClass5Counter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass5Counter()
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

// PatternFlowPfcPausePauseClass5Counter is integer counter pattern
type PatternFlowPfcPausePauseClass5Counter interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass5Counter to protobuf object *otg.PatternFlowPfcPausePauseClass5Counter
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass5Counter
	// setMsg unmarshals PatternFlowPfcPausePauseClass5Counter from protobuf object *otg.PatternFlowPfcPausePauseClass5Counter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass5Counter) PatternFlowPfcPausePauseClass5Counter
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass5Counter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass5Counter
	// validate validates PatternFlowPfcPausePauseClass5Counter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass5Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowPfcPausePauseClass5Counter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowPfcPausePauseClass5Counter
	SetStart(value uint32) PatternFlowPfcPausePauseClass5Counter
	// HasStart checks if Start has been set in PatternFlowPfcPausePauseClass5Counter
	HasStart() bool
	// Step returns uint32, set in PatternFlowPfcPausePauseClass5Counter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowPfcPausePauseClass5Counter
	SetStep(value uint32) PatternFlowPfcPausePauseClass5Counter
	// HasStep checks if Step has been set in PatternFlowPfcPausePauseClass5Counter
	HasStep() bool
	// Count returns uint32, set in PatternFlowPfcPausePauseClass5Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowPfcPausePauseClass5Counter
	SetCount(value uint32) PatternFlowPfcPausePauseClass5Counter
	// HasCount checks if Count has been set in PatternFlowPfcPausePauseClass5Counter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPausePauseClass5Counter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPausePauseClass5Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowPfcPausePauseClass5Counter object
func (obj *patternFlowPfcPausePauseClass5Counter) SetStart(value uint32) PatternFlowPfcPausePauseClass5Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPausePauseClass5Counter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPausePauseClass5Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowPfcPausePauseClass5Counter object
func (obj *patternFlowPfcPausePauseClass5Counter) SetStep(value uint32) PatternFlowPfcPausePauseClass5Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPausePauseClass5Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPausePauseClass5Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowPfcPausePauseClass5Counter object
func (obj *patternFlowPfcPausePauseClass5Counter) SetCount(value uint32) PatternFlowPfcPausePauseClass5Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass5Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass5Counter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass5Counter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass5Counter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass5Counter) setDefault() {
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

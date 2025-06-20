package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass4Counter *****
type patternFlowPfcPausePauseClass4Counter struct {
	validation
	obj          *otg.PatternFlowPfcPausePauseClass4Counter
	marshaller   marshalPatternFlowPfcPausePauseClass4Counter
	unMarshaller unMarshalPatternFlowPfcPausePauseClass4Counter
}

func NewPatternFlowPfcPausePauseClass4Counter() PatternFlowPfcPausePauseClass4Counter {
	obj := patternFlowPfcPausePauseClass4Counter{obj: &otg.PatternFlowPfcPausePauseClass4Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass4Counter) msg() *otg.PatternFlowPfcPausePauseClass4Counter {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass4Counter) setMsg(msg *otg.PatternFlowPfcPausePauseClass4Counter) PatternFlowPfcPausePauseClass4Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass4Counter struct {
	obj *patternFlowPfcPausePauseClass4Counter
}

type marshalPatternFlowPfcPausePauseClass4Counter interface {
	// ToProto marshals PatternFlowPfcPausePauseClass4Counter to protobuf object *otg.PatternFlowPfcPausePauseClass4Counter
	ToProto() (*otg.PatternFlowPfcPausePauseClass4Counter, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass4Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass4Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass4Counter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass4Counter struct {
	obj *patternFlowPfcPausePauseClass4Counter
}

type unMarshalPatternFlowPfcPausePauseClass4Counter interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass4Counter from protobuf object *otg.PatternFlowPfcPausePauseClass4Counter
	FromProto(msg *otg.PatternFlowPfcPausePauseClass4Counter) (PatternFlowPfcPausePauseClass4Counter, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass4Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass4Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass4Counter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass4Counter) Marshal() marshalPatternFlowPfcPausePauseClass4Counter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass4Counter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass4Counter) Unmarshal() unMarshalPatternFlowPfcPausePauseClass4Counter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass4Counter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass4Counter) ToProto() (*otg.PatternFlowPfcPausePauseClass4Counter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass4Counter) FromProto(msg *otg.PatternFlowPfcPausePauseClass4Counter) (PatternFlowPfcPausePauseClass4Counter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass4Counter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass4Counter) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass4Counter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass4Counter) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass4Counter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass4Counter) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass4Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass4Counter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass4Counter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass4Counter) Clone() (PatternFlowPfcPausePauseClass4Counter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass4Counter()
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

// PatternFlowPfcPausePauseClass4Counter is integer counter pattern
type PatternFlowPfcPausePauseClass4Counter interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass4Counter to protobuf object *otg.PatternFlowPfcPausePauseClass4Counter
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass4Counter
	// setMsg unmarshals PatternFlowPfcPausePauseClass4Counter from protobuf object *otg.PatternFlowPfcPausePauseClass4Counter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass4Counter) PatternFlowPfcPausePauseClass4Counter
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass4Counter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass4Counter
	// validate validates PatternFlowPfcPausePauseClass4Counter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass4Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowPfcPausePauseClass4Counter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowPfcPausePauseClass4Counter
	SetStart(value uint32) PatternFlowPfcPausePauseClass4Counter
	// HasStart checks if Start has been set in PatternFlowPfcPausePauseClass4Counter
	HasStart() bool
	// Step returns uint32, set in PatternFlowPfcPausePauseClass4Counter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowPfcPausePauseClass4Counter
	SetStep(value uint32) PatternFlowPfcPausePauseClass4Counter
	// HasStep checks if Step has been set in PatternFlowPfcPausePauseClass4Counter
	HasStep() bool
	// Count returns uint32, set in PatternFlowPfcPausePauseClass4Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowPfcPausePauseClass4Counter
	SetCount(value uint32) PatternFlowPfcPausePauseClass4Counter
	// HasCount checks if Count has been set in PatternFlowPfcPausePauseClass4Counter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPausePauseClass4Counter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPausePauseClass4Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowPfcPausePauseClass4Counter object
func (obj *patternFlowPfcPausePauseClass4Counter) SetStart(value uint32) PatternFlowPfcPausePauseClass4Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPausePauseClass4Counter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPausePauseClass4Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowPfcPausePauseClass4Counter object
func (obj *patternFlowPfcPausePauseClass4Counter) SetStep(value uint32) PatternFlowPfcPausePauseClass4Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPausePauseClass4Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPausePauseClass4Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowPfcPausePauseClass4Counter object
func (obj *patternFlowPfcPausePauseClass4Counter) SetCount(value uint32) PatternFlowPfcPausePauseClass4Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass4Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass4Counter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass4Counter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass4Counter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass4Counter) setDefault() {
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

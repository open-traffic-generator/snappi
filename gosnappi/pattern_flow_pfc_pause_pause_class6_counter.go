package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass6Counter *****
type patternFlowPfcPausePauseClass6Counter struct {
	validation
	obj          *otg.PatternFlowPfcPausePauseClass6Counter
	marshaller   marshalPatternFlowPfcPausePauseClass6Counter
	unMarshaller unMarshalPatternFlowPfcPausePauseClass6Counter
}

func NewPatternFlowPfcPausePauseClass6Counter() PatternFlowPfcPausePauseClass6Counter {
	obj := patternFlowPfcPausePauseClass6Counter{obj: &otg.PatternFlowPfcPausePauseClass6Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass6Counter) msg() *otg.PatternFlowPfcPausePauseClass6Counter {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass6Counter) setMsg(msg *otg.PatternFlowPfcPausePauseClass6Counter) PatternFlowPfcPausePauseClass6Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass6Counter struct {
	obj *patternFlowPfcPausePauseClass6Counter
}

type marshalPatternFlowPfcPausePauseClass6Counter interface {
	// ToProto marshals PatternFlowPfcPausePauseClass6Counter to protobuf object *otg.PatternFlowPfcPausePauseClass6Counter
	ToProto() (*otg.PatternFlowPfcPausePauseClass6Counter, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass6Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass6Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass6Counter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass6Counter struct {
	obj *patternFlowPfcPausePauseClass6Counter
}

type unMarshalPatternFlowPfcPausePauseClass6Counter interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass6Counter from protobuf object *otg.PatternFlowPfcPausePauseClass6Counter
	FromProto(msg *otg.PatternFlowPfcPausePauseClass6Counter) (PatternFlowPfcPausePauseClass6Counter, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass6Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass6Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass6Counter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass6Counter) Marshal() marshalPatternFlowPfcPausePauseClass6Counter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass6Counter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass6Counter) Unmarshal() unMarshalPatternFlowPfcPausePauseClass6Counter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass6Counter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass6Counter) ToProto() (*otg.PatternFlowPfcPausePauseClass6Counter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass6Counter) FromProto(msg *otg.PatternFlowPfcPausePauseClass6Counter) (PatternFlowPfcPausePauseClass6Counter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass6Counter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass6Counter) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass6Counter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass6Counter) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass6Counter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass6Counter) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass6Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass6Counter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass6Counter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass6Counter) Clone() (PatternFlowPfcPausePauseClass6Counter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass6Counter()
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

// PatternFlowPfcPausePauseClass6Counter is integer counter pattern
type PatternFlowPfcPausePauseClass6Counter interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass6Counter to protobuf object *otg.PatternFlowPfcPausePauseClass6Counter
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass6Counter
	// setMsg unmarshals PatternFlowPfcPausePauseClass6Counter from protobuf object *otg.PatternFlowPfcPausePauseClass6Counter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass6Counter) PatternFlowPfcPausePauseClass6Counter
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass6Counter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass6Counter
	// validate validates PatternFlowPfcPausePauseClass6Counter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass6Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowPfcPausePauseClass6Counter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowPfcPausePauseClass6Counter
	SetStart(value uint32) PatternFlowPfcPausePauseClass6Counter
	// HasStart checks if Start has been set in PatternFlowPfcPausePauseClass6Counter
	HasStart() bool
	// Step returns uint32, set in PatternFlowPfcPausePauseClass6Counter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowPfcPausePauseClass6Counter
	SetStep(value uint32) PatternFlowPfcPausePauseClass6Counter
	// HasStep checks if Step has been set in PatternFlowPfcPausePauseClass6Counter
	HasStep() bool
	// Count returns uint32, set in PatternFlowPfcPausePauseClass6Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowPfcPausePauseClass6Counter
	SetCount(value uint32) PatternFlowPfcPausePauseClass6Counter
	// HasCount checks if Count has been set in PatternFlowPfcPausePauseClass6Counter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPausePauseClass6Counter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPausePauseClass6Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowPfcPausePauseClass6Counter object
func (obj *patternFlowPfcPausePauseClass6Counter) SetStart(value uint32) PatternFlowPfcPausePauseClass6Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPausePauseClass6Counter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPausePauseClass6Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowPfcPausePauseClass6Counter object
func (obj *patternFlowPfcPausePauseClass6Counter) SetStep(value uint32) PatternFlowPfcPausePauseClass6Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPausePauseClass6Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPausePauseClass6Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowPfcPausePauseClass6Counter object
func (obj *patternFlowPfcPausePauseClass6Counter) SetCount(value uint32) PatternFlowPfcPausePauseClass6Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass6Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass6Counter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass6Counter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass6Counter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass6Counter) setDefault() {
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

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass2Counter *****
type patternFlowPfcPausePauseClass2Counter struct {
	validation
	obj          *otg.PatternFlowPfcPausePauseClass2Counter
	marshaller   marshalPatternFlowPfcPausePauseClass2Counter
	unMarshaller unMarshalPatternFlowPfcPausePauseClass2Counter
}

func NewPatternFlowPfcPausePauseClass2Counter() PatternFlowPfcPausePauseClass2Counter {
	obj := patternFlowPfcPausePauseClass2Counter{obj: &otg.PatternFlowPfcPausePauseClass2Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass2Counter) msg() *otg.PatternFlowPfcPausePauseClass2Counter {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass2Counter) setMsg(msg *otg.PatternFlowPfcPausePauseClass2Counter) PatternFlowPfcPausePauseClass2Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass2Counter struct {
	obj *patternFlowPfcPausePauseClass2Counter
}

type marshalPatternFlowPfcPausePauseClass2Counter interface {
	// ToProto marshals PatternFlowPfcPausePauseClass2Counter to protobuf object *otg.PatternFlowPfcPausePauseClass2Counter
	ToProto() (*otg.PatternFlowPfcPausePauseClass2Counter, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass2Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass2Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass2Counter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass2Counter struct {
	obj *patternFlowPfcPausePauseClass2Counter
}

type unMarshalPatternFlowPfcPausePauseClass2Counter interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass2Counter from protobuf object *otg.PatternFlowPfcPausePauseClass2Counter
	FromProto(msg *otg.PatternFlowPfcPausePauseClass2Counter) (PatternFlowPfcPausePauseClass2Counter, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass2Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass2Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass2Counter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass2Counter) Marshal() marshalPatternFlowPfcPausePauseClass2Counter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass2Counter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass2Counter) Unmarshal() unMarshalPatternFlowPfcPausePauseClass2Counter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass2Counter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass2Counter) ToProto() (*otg.PatternFlowPfcPausePauseClass2Counter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass2Counter) FromProto(msg *otg.PatternFlowPfcPausePauseClass2Counter) (PatternFlowPfcPausePauseClass2Counter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass2Counter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass2Counter) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass2Counter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass2Counter) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass2Counter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass2Counter) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass2Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass2Counter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass2Counter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass2Counter) Clone() (PatternFlowPfcPausePauseClass2Counter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass2Counter()
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

// PatternFlowPfcPausePauseClass2Counter is integer counter pattern
type PatternFlowPfcPausePauseClass2Counter interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass2Counter to protobuf object *otg.PatternFlowPfcPausePauseClass2Counter
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass2Counter
	// setMsg unmarshals PatternFlowPfcPausePauseClass2Counter from protobuf object *otg.PatternFlowPfcPausePauseClass2Counter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass2Counter) PatternFlowPfcPausePauseClass2Counter
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass2Counter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass2Counter
	// validate validates PatternFlowPfcPausePauseClass2Counter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass2Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowPfcPausePauseClass2Counter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowPfcPausePauseClass2Counter
	SetStart(value uint32) PatternFlowPfcPausePauseClass2Counter
	// HasStart checks if Start has been set in PatternFlowPfcPausePauseClass2Counter
	HasStart() bool
	// Step returns uint32, set in PatternFlowPfcPausePauseClass2Counter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowPfcPausePauseClass2Counter
	SetStep(value uint32) PatternFlowPfcPausePauseClass2Counter
	// HasStep checks if Step has been set in PatternFlowPfcPausePauseClass2Counter
	HasStep() bool
	// Count returns uint32, set in PatternFlowPfcPausePauseClass2Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowPfcPausePauseClass2Counter
	SetCount(value uint32) PatternFlowPfcPausePauseClass2Counter
	// HasCount checks if Count has been set in PatternFlowPfcPausePauseClass2Counter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPausePauseClass2Counter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPausePauseClass2Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowPfcPausePauseClass2Counter object
func (obj *patternFlowPfcPausePauseClass2Counter) SetStart(value uint32) PatternFlowPfcPausePauseClass2Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPausePauseClass2Counter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPausePauseClass2Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowPfcPausePauseClass2Counter object
func (obj *patternFlowPfcPausePauseClass2Counter) SetStep(value uint32) PatternFlowPfcPausePauseClass2Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPausePauseClass2Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPausePauseClass2Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowPfcPausePauseClass2Counter object
func (obj *patternFlowPfcPausePauseClass2Counter) SetCount(value uint32) PatternFlowPfcPausePauseClass2Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass2Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass2Counter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass2Counter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass2Counter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass2Counter) setDefault() {
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

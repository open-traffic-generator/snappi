package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass7Counter *****
type patternFlowPfcPausePauseClass7Counter struct {
	validation
	obj          *otg.PatternFlowPfcPausePauseClass7Counter
	marshaller   marshalPatternFlowPfcPausePauseClass7Counter
	unMarshaller unMarshalPatternFlowPfcPausePauseClass7Counter
}

func NewPatternFlowPfcPausePauseClass7Counter() PatternFlowPfcPausePauseClass7Counter {
	obj := patternFlowPfcPausePauseClass7Counter{obj: &otg.PatternFlowPfcPausePauseClass7Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass7Counter) msg() *otg.PatternFlowPfcPausePauseClass7Counter {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass7Counter) setMsg(msg *otg.PatternFlowPfcPausePauseClass7Counter) PatternFlowPfcPausePauseClass7Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass7Counter struct {
	obj *patternFlowPfcPausePauseClass7Counter
}

type marshalPatternFlowPfcPausePauseClass7Counter interface {
	// ToProto marshals PatternFlowPfcPausePauseClass7Counter to protobuf object *otg.PatternFlowPfcPausePauseClass7Counter
	ToProto() (*otg.PatternFlowPfcPausePauseClass7Counter, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass7Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass7Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass7Counter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass7Counter struct {
	obj *patternFlowPfcPausePauseClass7Counter
}

type unMarshalPatternFlowPfcPausePauseClass7Counter interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass7Counter from protobuf object *otg.PatternFlowPfcPausePauseClass7Counter
	FromProto(msg *otg.PatternFlowPfcPausePauseClass7Counter) (PatternFlowPfcPausePauseClass7Counter, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass7Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass7Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass7Counter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass7Counter) Marshal() marshalPatternFlowPfcPausePauseClass7Counter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass7Counter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass7Counter) Unmarshal() unMarshalPatternFlowPfcPausePauseClass7Counter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass7Counter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass7Counter) ToProto() (*otg.PatternFlowPfcPausePauseClass7Counter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass7Counter) FromProto(msg *otg.PatternFlowPfcPausePauseClass7Counter) (PatternFlowPfcPausePauseClass7Counter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass7Counter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass7Counter) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass7Counter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass7Counter) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass7Counter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass7Counter) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass7Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass7Counter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass7Counter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass7Counter) Clone() (PatternFlowPfcPausePauseClass7Counter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass7Counter()
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

// PatternFlowPfcPausePauseClass7Counter is integer counter pattern
type PatternFlowPfcPausePauseClass7Counter interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass7Counter to protobuf object *otg.PatternFlowPfcPausePauseClass7Counter
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass7Counter
	// setMsg unmarshals PatternFlowPfcPausePauseClass7Counter from protobuf object *otg.PatternFlowPfcPausePauseClass7Counter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass7Counter) PatternFlowPfcPausePauseClass7Counter
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass7Counter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass7Counter
	// validate validates PatternFlowPfcPausePauseClass7Counter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass7Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowPfcPausePauseClass7Counter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowPfcPausePauseClass7Counter
	SetStart(value uint32) PatternFlowPfcPausePauseClass7Counter
	// HasStart checks if Start has been set in PatternFlowPfcPausePauseClass7Counter
	HasStart() bool
	// Step returns uint32, set in PatternFlowPfcPausePauseClass7Counter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowPfcPausePauseClass7Counter
	SetStep(value uint32) PatternFlowPfcPausePauseClass7Counter
	// HasStep checks if Step has been set in PatternFlowPfcPausePauseClass7Counter
	HasStep() bool
	// Count returns uint32, set in PatternFlowPfcPausePauseClass7Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowPfcPausePauseClass7Counter
	SetCount(value uint32) PatternFlowPfcPausePauseClass7Counter
	// HasCount checks if Count has been set in PatternFlowPfcPausePauseClass7Counter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPausePauseClass7Counter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPausePauseClass7Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowPfcPausePauseClass7Counter object
func (obj *patternFlowPfcPausePauseClass7Counter) SetStart(value uint32) PatternFlowPfcPausePauseClass7Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPausePauseClass7Counter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPausePauseClass7Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowPfcPausePauseClass7Counter object
func (obj *patternFlowPfcPausePauseClass7Counter) SetStep(value uint32) PatternFlowPfcPausePauseClass7Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPausePauseClass7Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPausePauseClass7Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowPfcPausePauseClass7Counter object
func (obj *patternFlowPfcPausePauseClass7Counter) SetCount(value uint32) PatternFlowPfcPausePauseClass7Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass7Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass7Counter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass7Counter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass7Counter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass7Counter) setDefault() {
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

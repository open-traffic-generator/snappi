package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass3Counter *****
type patternFlowPfcPausePauseClass3Counter struct {
	validation
	obj          *otg.PatternFlowPfcPausePauseClass3Counter
	marshaller   marshalPatternFlowPfcPausePauseClass3Counter
	unMarshaller unMarshalPatternFlowPfcPausePauseClass3Counter
}

func NewPatternFlowPfcPausePauseClass3Counter() PatternFlowPfcPausePauseClass3Counter {
	obj := patternFlowPfcPausePauseClass3Counter{obj: &otg.PatternFlowPfcPausePauseClass3Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass3Counter) msg() *otg.PatternFlowPfcPausePauseClass3Counter {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass3Counter) setMsg(msg *otg.PatternFlowPfcPausePauseClass3Counter) PatternFlowPfcPausePauseClass3Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass3Counter struct {
	obj *patternFlowPfcPausePauseClass3Counter
}

type marshalPatternFlowPfcPausePauseClass3Counter interface {
	// ToProto marshals PatternFlowPfcPausePauseClass3Counter to protobuf object *otg.PatternFlowPfcPausePauseClass3Counter
	ToProto() (*otg.PatternFlowPfcPausePauseClass3Counter, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass3Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass3Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass3Counter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass3Counter struct {
	obj *patternFlowPfcPausePauseClass3Counter
}

type unMarshalPatternFlowPfcPausePauseClass3Counter interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass3Counter from protobuf object *otg.PatternFlowPfcPausePauseClass3Counter
	FromProto(msg *otg.PatternFlowPfcPausePauseClass3Counter) (PatternFlowPfcPausePauseClass3Counter, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass3Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass3Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass3Counter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass3Counter) Marshal() marshalPatternFlowPfcPausePauseClass3Counter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass3Counter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass3Counter) Unmarshal() unMarshalPatternFlowPfcPausePauseClass3Counter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass3Counter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass3Counter) ToProto() (*otg.PatternFlowPfcPausePauseClass3Counter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass3Counter) FromProto(msg *otg.PatternFlowPfcPausePauseClass3Counter) (PatternFlowPfcPausePauseClass3Counter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass3Counter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass3Counter) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass3Counter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass3Counter) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass3Counter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass3Counter) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass3Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass3Counter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass3Counter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass3Counter) Clone() (PatternFlowPfcPausePauseClass3Counter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass3Counter()
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

// PatternFlowPfcPausePauseClass3Counter is integer counter pattern
type PatternFlowPfcPausePauseClass3Counter interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass3Counter to protobuf object *otg.PatternFlowPfcPausePauseClass3Counter
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass3Counter
	// setMsg unmarshals PatternFlowPfcPausePauseClass3Counter from protobuf object *otg.PatternFlowPfcPausePauseClass3Counter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass3Counter) PatternFlowPfcPausePauseClass3Counter
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass3Counter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass3Counter
	// validate validates PatternFlowPfcPausePauseClass3Counter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass3Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowPfcPausePauseClass3Counter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowPfcPausePauseClass3Counter
	SetStart(value uint32) PatternFlowPfcPausePauseClass3Counter
	// HasStart checks if Start has been set in PatternFlowPfcPausePauseClass3Counter
	HasStart() bool
	// Step returns uint32, set in PatternFlowPfcPausePauseClass3Counter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowPfcPausePauseClass3Counter
	SetStep(value uint32) PatternFlowPfcPausePauseClass3Counter
	// HasStep checks if Step has been set in PatternFlowPfcPausePauseClass3Counter
	HasStep() bool
	// Count returns uint32, set in PatternFlowPfcPausePauseClass3Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowPfcPausePauseClass3Counter
	SetCount(value uint32) PatternFlowPfcPausePauseClass3Counter
	// HasCount checks if Count has been set in PatternFlowPfcPausePauseClass3Counter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPausePauseClass3Counter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPausePauseClass3Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowPfcPausePauseClass3Counter object
func (obj *patternFlowPfcPausePauseClass3Counter) SetStart(value uint32) PatternFlowPfcPausePauseClass3Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPausePauseClass3Counter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPausePauseClass3Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowPfcPausePauseClass3Counter object
func (obj *patternFlowPfcPausePauseClass3Counter) SetStep(value uint32) PatternFlowPfcPausePauseClass3Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPausePauseClass3Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPausePauseClass3Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowPfcPausePauseClass3Counter object
func (obj *patternFlowPfcPausePauseClass3Counter) SetCount(value uint32) PatternFlowPfcPausePauseClass3Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass3Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass3Counter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass3Counter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass3Counter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass3Counter) setDefault() {
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

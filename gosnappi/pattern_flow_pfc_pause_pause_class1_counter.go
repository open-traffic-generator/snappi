package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass1Counter *****
type patternFlowPfcPausePauseClass1Counter struct {
	validation
	obj          *otg.PatternFlowPfcPausePauseClass1Counter
	marshaller   marshalPatternFlowPfcPausePauseClass1Counter
	unMarshaller unMarshalPatternFlowPfcPausePauseClass1Counter
}

func NewPatternFlowPfcPausePauseClass1Counter() PatternFlowPfcPausePauseClass1Counter {
	obj := patternFlowPfcPausePauseClass1Counter{obj: &otg.PatternFlowPfcPausePauseClass1Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass1Counter) msg() *otg.PatternFlowPfcPausePauseClass1Counter {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass1Counter) setMsg(msg *otg.PatternFlowPfcPausePauseClass1Counter) PatternFlowPfcPausePauseClass1Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass1Counter struct {
	obj *patternFlowPfcPausePauseClass1Counter
}

type marshalPatternFlowPfcPausePauseClass1Counter interface {
	// ToProto marshals PatternFlowPfcPausePauseClass1Counter to protobuf object *otg.PatternFlowPfcPausePauseClass1Counter
	ToProto() (*otg.PatternFlowPfcPausePauseClass1Counter, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass1Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass1Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass1Counter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowPfcPausePauseClass1Counter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass1Counter struct {
	obj *patternFlowPfcPausePauseClass1Counter
}

type unMarshalPatternFlowPfcPausePauseClass1Counter interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass1Counter from protobuf object *otg.PatternFlowPfcPausePauseClass1Counter
	FromProto(msg *otg.PatternFlowPfcPausePauseClass1Counter) (PatternFlowPfcPausePauseClass1Counter, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass1Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass1Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass1Counter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass1Counter) Marshal() marshalPatternFlowPfcPausePauseClass1Counter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass1Counter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass1Counter) Unmarshal() unMarshalPatternFlowPfcPausePauseClass1Counter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass1Counter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass1Counter) ToProto() (*otg.PatternFlowPfcPausePauseClass1Counter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass1Counter) FromProto(msg *otg.PatternFlowPfcPausePauseClass1Counter) (PatternFlowPfcPausePauseClass1Counter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass1Counter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass1Counter) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass1Counter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass1Counter) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass1Counter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowPfcPausePauseClass1Counter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass1Counter) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass1Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass1Counter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass1Counter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass1Counter) Clone() (PatternFlowPfcPausePauseClass1Counter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass1Counter()
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

// PatternFlowPfcPausePauseClass1Counter is integer counter pattern
type PatternFlowPfcPausePauseClass1Counter interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass1Counter to protobuf object *otg.PatternFlowPfcPausePauseClass1Counter
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass1Counter
	// setMsg unmarshals PatternFlowPfcPausePauseClass1Counter from protobuf object *otg.PatternFlowPfcPausePauseClass1Counter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass1Counter) PatternFlowPfcPausePauseClass1Counter
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass1Counter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass1Counter
	// validate validates PatternFlowPfcPausePauseClass1Counter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass1Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowPfcPausePauseClass1Counter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowPfcPausePauseClass1Counter
	SetStart(value uint32) PatternFlowPfcPausePauseClass1Counter
	// HasStart checks if Start has been set in PatternFlowPfcPausePauseClass1Counter
	HasStart() bool
	// Step returns uint32, set in PatternFlowPfcPausePauseClass1Counter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowPfcPausePauseClass1Counter
	SetStep(value uint32) PatternFlowPfcPausePauseClass1Counter
	// HasStep checks if Step has been set in PatternFlowPfcPausePauseClass1Counter
	HasStep() bool
	// Count returns uint32, set in PatternFlowPfcPausePauseClass1Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowPfcPausePauseClass1Counter
	SetCount(value uint32) PatternFlowPfcPausePauseClass1Counter
	// HasCount checks if Count has been set in PatternFlowPfcPausePauseClass1Counter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPausePauseClass1Counter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPausePauseClass1Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowPfcPausePauseClass1Counter object
func (obj *patternFlowPfcPausePauseClass1Counter) SetStart(value uint32) PatternFlowPfcPausePauseClass1Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPausePauseClass1Counter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPausePauseClass1Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowPfcPausePauseClass1Counter object
func (obj *patternFlowPfcPausePauseClass1Counter) SetStep(value uint32) PatternFlowPfcPausePauseClass1Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPausePauseClass1Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPausePauseClass1Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowPfcPausePauseClass1Counter object
func (obj *patternFlowPfcPausePauseClass1Counter) SetCount(value uint32) PatternFlowPfcPausePauseClass1Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass1Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass1Counter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass1Counter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass1Counter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass1Counter) setDefault() {
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

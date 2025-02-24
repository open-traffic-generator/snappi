package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass0Counter *****
type patternFlowPfcPausePauseClass0Counter struct {
	validation
	obj          *otg.PatternFlowPfcPausePauseClass0Counter
	marshaller   marshalPatternFlowPfcPausePauseClass0Counter
	unMarshaller unMarshalPatternFlowPfcPausePauseClass0Counter
}

func NewPatternFlowPfcPausePauseClass0Counter() PatternFlowPfcPausePauseClass0Counter {
	obj := patternFlowPfcPausePauseClass0Counter{obj: &otg.PatternFlowPfcPausePauseClass0Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass0Counter) msg() *otg.PatternFlowPfcPausePauseClass0Counter {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass0Counter) setMsg(msg *otg.PatternFlowPfcPausePauseClass0Counter) PatternFlowPfcPausePauseClass0Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass0Counter struct {
	obj *patternFlowPfcPausePauseClass0Counter
}

type marshalPatternFlowPfcPausePauseClass0Counter interface {
	// ToProto marshals PatternFlowPfcPausePauseClass0Counter to protobuf object *otg.PatternFlowPfcPausePauseClass0Counter
	ToProto() (*otg.PatternFlowPfcPausePauseClass0Counter, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass0Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass0Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass0Counter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowPfcPausePauseClass0Counter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass0Counter struct {
	obj *patternFlowPfcPausePauseClass0Counter
}

type unMarshalPatternFlowPfcPausePauseClass0Counter interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass0Counter from protobuf object *otg.PatternFlowPfcPausePauseClass0Counter
	FromProto(msg *otg.PatternFlowPfcPausePauseClass0Counter) (PatternFlowPfcPausePauseClass0Counter, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass0Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass0Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass0Counter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass0Counter) Marshal() marshalPatternFlowPfcPausePauseClass0Counter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass0Counter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass0Counter) Unmarshal() unMarshalPatternFlowPfcPausePauseClass0Counter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass0Counter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass0Counter) ToProto() (*otg.PatternFlowPfcPausePauseClass0Counter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass0Counter) FromProto(msg *otg.PatternFlowPfcPausePauseClass0Counter) (PatternFlowPfcPausePauseClass0Counter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass0Counter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass0Counter) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass0Counter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass0Counter) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass0Counter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowPfcPausePauseClass0Counter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass0Counter) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass0Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass0Counter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass0Counter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass0Counter) Clone() (PatternFlowPfcPausePauseClass0Counter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass0Counter()
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

// PatternFlowPfcPausePauseClass0Counter is integer counter pattern
type PatternFlowPfcPausePauseClass0Counter interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass0Counter to protobuf object *otg.PatternFlowPfcPausePauseClass0Counter
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass0Counter
	// setMsg unmarshals PatternFlowPfcPausePauseClass0Counter from protobuf object *otg.PatternFlowPfcPausePauseClass0Counter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass0Counter) PatternFlowPfcPausePauseClass0Counter
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass0Counter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass0Counter
	// validate validates PatternFlowPfcPausePauseClass0Counter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass0Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowPfcPausePauseClass0Counter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowPfcPausePauseClass0Counter
	SetStart(value uint32) PatternFlowPfcPausePauseClass0Counter
	// HasStart checks if Start has been set in PatternFlowPfcPausePauseClass0Counter
	HasStart() bool
	// Step returns uint32, set in PatternFlowPfcPausePauseClass0Counter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowPfcPausePauseClass0Counter
	SetStep(value uint32) PatternFlowPfcPausePauseClass0Counter
	// HasStep checks if Step has been set in PatternFlowPfcPausePauseClass0Counter
	HasStep() bool
	// Count returns uint32, set in PatternFlowPfcPausePauseClass0Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowPfcPausePauseClass0Counter
	SetCount(value uint32) PatternFlowPfcPausePauseClass0Counter
	// HasCount checks if Count has been set in PatternFlowPfcPausePauseClass0Counter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPausePauseClass0Counter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPfcPausePauseClass0Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowPfcPausePauseClass0Counter object
func (obj *patternFlowPfcPausePauseClass0Counter) SetStart(value uint32) PatternFlowPfcPausePauseClass0Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPausePauseClass0Counter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPfcPausePauseClass0Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowPfcPausePauseClass0Counter object
func (obj *patternFlowPfcPausePauseClass0Counter) SetStep(value uint32) PatternFlowPfcPausePauseClass0Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPausePauseClass0Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPfcPausePauseClass0Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowPfcPausePauseClass0Counter object
func (obj *patternFlowPfcPausePauseClass0Counter) SetCount(value uint32) PatternFlowPfcPausePauseClass0Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass0Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass0Counter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass0Counter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass0Counter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass0Counter) setDefault() {
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

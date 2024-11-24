package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServReserved1Counter *****
type patternFlowRSVPPathSenderTspecIntServReserved1Counter struct {
	validation
	obj          *otg.PatternFlowRSVPPathSenderTspecIntServReserved1Counter
	marshaller   marshalPatternFlowRSVPPathSenderTspecIntServReserved1Counter
	unMarshaller unMarshalPatternFlowRSVPPathSenderTspecIntServReserved1Counter
}

func NewPatternFlowRSVPPathSenderTspecIntServReserved1Counter() PatternFlowRSVPPathSenderTspecIntServReserved1Counter {
	obj := patternFlowRSVPPathSenderTspecIntServReserved1Counter{obj: &otg.PatternFlowRSVPPathSenderTspecIntServReserved1Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1Counter) msg() *otg.PatternFlowRSVPPathSenderTspecIntServReserved1Counter {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1Counter) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServReserved1Counter) PatternFlowRSVPPathSenderTspecIntServReserved1Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServReserved1Counter struct {
	obj *patternFlowRSVPPathSenderTspecIntServReserved1Counter
}

type marshalPatternFlowRSVPPathSenderTspecIntServReserved1Counter interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServReserved1Counter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServReserved1Counter
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServReserved1Counter, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServReserved1Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServReserved1Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServReserved1Counter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathSenderTspecIntServReserved1Counter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServReserved1Counter struct {
	obj *patternFlowRSVPPathSenderTspecIntServReserved1Counter
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServReserved1Counter interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServReserved1Counter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServReserved1Counter
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServReserved1Counter) (PatternFlowRSVPPathSenderTspecIntServReserved1Counter, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServReserved1Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServReserved1Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServReserved1Counter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1Counter) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServReserved1Counter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServReserved1Counter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1Counter) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServReserved1Counter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServReserved1Counter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServReserved1Counter) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServReserved1Counter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServReserved1Counter) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServReserved1Counter) (PatternFlowRSVPPathSenderTspecIntServReserved1Counter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServReserved1Counter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServReserved1Counter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServReserved1Counter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServReserved1Counter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServReserved1Counter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServReserved1Counter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServReserved1Counter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1Counter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1Counter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1Counter) Clone() (PatternFlowRSVPPathSenderTspecIntServReserved1Counter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServReserved1Counter()
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

// PatternFlowRSVPPathSenderTspecIntServReserved1Counter is integer counter pattern
type PatternFlowRSVPPathSenderTspecIntServReserved1Counter interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServReserved1Counter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServReserved1Counter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServReserved1Counter
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServReserved1Counter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServReserved1Counter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServReserved1Counter) PatternFlowRSVPPathSenderTspecIntServReserved1Counter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServReserved1Counter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServReserved1Counter
	// validate validates PatternFlowRSVPPathSenderTspecIntServReserved1Counter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServReserved1Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathSenderTspecIntServReserved1Counter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServReserved1Counter
	SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServReserved1Counter
	// HasStart checks if Start has been set in PatternFlowRSVPPathSenderTspecIntServReserved1Counter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathSenderTspecIntServReserved1Counter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServReserved1Counter
	SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServReserved1Counter
	// HasStep checks if Step has been set in PatternFlowRSVPPathSenderTspecIntServReserved1Counter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathSenderTspecIntServReserved1Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServReserved1Counter
	SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServReserved1Counter
	// HasCount checks if Count has been set in PatternFlowRSVPPathSenderTspecIntServReserved1Counter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServReserved1Counter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServReserved1Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServReserved1Counter object
func (obj *patternFlowRSVPPathSenderTspecIntServReserved1Counter) SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServReserved1Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServReserved1Counter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServReserved1Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServReserved1Counter object
func (obj *patternFlowRSVPPathSenderTspecIntServReserved1Counter) SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServReserved1Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServReserved1Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServReserved1Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServReserved1Counter object
func (obj *patternFlowRSVPPathSenderTspecIntServReserved1Counter) SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServReserved1Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 4095 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServReserved1Counter.Start <= 4095 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 4095 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServReserved1Counter.Step <= 4095 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 4095 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServReserved1Counter.Count <= 4095 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1Counter) setDefault() {
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

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServReserved2Counter *****
type patternFlowRSVPPathSenderTspecIntServReserved2Counter struct {
	validation
	obj          *otg.PatternFlowRSVPPathSenderTspecIntServReserved2Counter
	marshaller   marshalPatternFlowRSVPPathSenderTspecIntServReserved2Counter
	unMarshaller unMarshalPatternFlowRSVPPathSenderTspecIntServReserved2Counter
}

func NewPatternFlowRSVPPathSenderTspecIntServReserved2Counter() PatternFlowRSVPPathSenderTspecIntServReserved2Counter {
	obj := patternFlowRSVPPathSenderTspecIntServReserved2Counter{obj: &otg.PatternFlowRSVPPathSenderTspecIntServReserved2Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2Counter) msg() *otg.PatternFlowRSVPPathSenderTspecIntServReserved2Counter {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2Counter) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServReserved2Counter) PatternFlowRSVPPathSenderTspecIntServReserved2Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServReserved2Counter struct {
	obj *patternFlowRSVPPathSenderTspecIntServReserved2Counter
}

type marshalPatternFlowRSVPPathSenderTspecIntServReserved2Counter interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServReserved2Counter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServReserved2Counter
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServReserved2Counter, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServReserved2Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServReserved2Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServReserved2Counter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServReserved2Counter struct {
	obj *patternFlowRSVPPathSenderTspecIntServReserved2Counter
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServReserved2Counter interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServReserved2Counter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServReserved2Counter
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServReserved2Counter) (PatternFlowRSVPPathSenderTspecIntServReserved2Counter, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServReserved2Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServReserved2Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServReserved2Counter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2Counter) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServReserved2Counter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServReserved2Counter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2Counter) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServReserved2Counter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServReserved2Counter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServReserved2Counter) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServReserved2Counter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServReserved2Counter) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServReserved2Counter) (PatternFlowRSVPPathSenderTspecIntServReserved2Counter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServReserved2Counter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServReserved2Counter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServReserved2Counter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServReserved2Counter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServReserved2Counter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServReserved2Counter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2Counter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2Counter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2Counter) Clone() (PatternFlowRSVPPathSenderTspecIntServReserved2Counter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServReserved2Counter()
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

// PatternFlowRSVPPathSenderTspecIntServReserved2Counter is integer counter pattern
type PatternFlowRSVPPathSenderTspecIntServReserved2Counter interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServReserved2Counter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServReserved2Counter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServReserved2Counter
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServReserved2Counter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServReserved2Counter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServReserved2Counter) PatternFlowRSVPPathSenderTspecIntServReserved2Counter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServReserved2Counter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServReserved2Counter
	// validate validates PatternFlowRSVPPathSenderTspecIntServReserved2Counter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServReserved2Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathSenderTspecIntServReserved2Counter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServReserved2Counter
	SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServReserved2Counter
	// HasStart checks if Start has been set in PatternFlowRSVPPathSenderTspecIntServReserved2Counter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathSenderTspecIntServReserved2Counter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServReserved2Counter
	SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServReserved2Counter
	// HasStep checks if Step has been set in PatternFlowRSVPPathSenderTspecIntServReserved2Counter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathSenderTspecIntServReserved2Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServReserved2Counter
	SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServReserved2Counter
	// HasCount checks if Count has been set in PatternFlowRSVPPathSenderTspecIntServReserved2Counter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServReserved2Counter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServReserved2Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServReserved2Counter object
func (obj *patternFlowRSVPPathSenderTspecIntServReserved2Counter) SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServReserved2Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServReserved2Counter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServReserved2Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServReserved2Counter object
func (obj *patternFlowRSVPPathSenderTspecIntServReserved2Counter) SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServReserved2Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServReserved2Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServReserved2Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServReserved2Counter object
func (obj *patternFlowRSVPPathSenderTspecIntServReserved2Counter) SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServReserved2Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 127 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServReserved2Counter.Start <= 127 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 127 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServReserved2Counter.Step <= 127 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServReserved2Counter.Count <= 128 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2Counter) setDefault() {
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

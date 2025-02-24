package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter *****
type patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	marshaller   marshalPatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	unMarshaller unMarshalPatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
}

func NewPatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter() PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter {
	obj := patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter{obj: &otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) msg() *otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) setMsg(msg *otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter struct {
	obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
}

type marshalPatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter interface {
	// ToProto marshals PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter to protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	ToProto() (*otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter, error)
	// ToPbText marshals PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter struct {
	obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
}

type unMarshalPatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter from protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	FromProto(msg *otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) (PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) Marshal() marshalPatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) Unmarshal() unMarshalPatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) ToProto() (*otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) FromProto(msg *otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) (PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) Clone() (PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter()
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

// PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter is integer counter pattern
type PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter to protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	// setMsg unmarshals PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter from protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	// validate validates PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	SetStart(value uint32) PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	SetStep(value uint32) PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	SetCount(value uint32) PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter object
func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) SetStart(value uint32) PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter object
func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) SetStep(value uint32) PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter object
func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) SetCount(value uint32) PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) setDefault() {
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

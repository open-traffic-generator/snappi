package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerSystemIdCounter *****
type patternFlowLacpPartnerSystemIdCounter struct {
	validation
	obj          *otg.PatternFlowLacpPartnerSystemIdCounter
	marshaller   marshalPatternFlowLacpPartnerSystemIdCounter
	unMarshaller unMarshalPatternFlowLacpPartnerSystemIdCounter
}

func NewPatternFlowLacpPartnerSystemIdCounter() PatternFlowLacpPartnerSystemIdCounter {
	obj := patternFlowLacpPartnerSystemIdCounter{obj: &otg.PatternFlowLacpPartnerSystemIdCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerSystemIdCounter) msg() *otg.PatternFlowLacpPartnerSystemIdCounter {
	return obj.obj
}

func (obj *patternFlowLacpPartnerSystemIdCounter) setMsg(msg *otg.PatternFlowLacpPartnerSystemIdCounter) PatternFlowLacpPartnerSystemIdCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerSystemIdCounter struct {
	obj *patternFlowLacpPartnerSystemIdCounter
}

type marshalPatternFlowLacpPartnerSystemIdCounter interface {
	// ToProto marshals PatternFlowLacpPartnerSystemIdCounter to protobuf object *otg.PatternFlowLacpPartnerSystemIdCounter
	ToProto() (*otg.PatternFlowLacpPartnerSystemIdCounter, error)
	// ToPbText marshals PatternFlowLacpPartnerSystemIdCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerSystemIdCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerSystemIdCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerSystemIdCounter struct {
	obj *patternFlowLacpPartnerSystemIdCounter
}

type unMarshalPatternFlowLacpPartnerSystemIdCounter interface {
	// FromProto unmarshals PatternFlowLacpPartnerSystemIdCounter from protobuf object *otg.PatternFlowLacpPartnerSystemIdCounter
	FromProto(msg *otg.PatternFlowLacpPartnerSystemIdCounter) (PatternFlowLacpPartnerSystemIdCounter, error)
	// FromPbText unmarshals PatternFlowLacpPartnerSystemIdCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerSystemIdCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerSystemIdCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerSystemIdCounter) Marshal() marshalPatternFlowLacpPartnerSystemIdCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerSystemIdCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerSystemIdCounter) Unmarshal() unMarshalPatternFlowLacpPartnerSystemIdCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerSystemIdCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerSystemIdCounter) ToProto() (*otg.PatternFlowLacpPartnerSystemIdCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerSystemIdCounter) FromProto(msg *otg.PatternFlowLacpPartnerSystemIdCounter) (PatternFlowLacpPartnerSystemIdCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerSystemIdCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerSystemIdCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerSystemIdCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerSystemIdCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerSystemIdCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerSystemIdCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerSystemIdCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerSystemIdCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerSystemIdCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerSystemIdCounter) Clone() (PatternFlowLacpPartnerSystemIdCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerSystemIdCounter()
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

// PatternFlowLacpPartnerSystemIdCounter is mac counter pattern
type PatternFlowLacpPartnerSystemIdCounter interface {
	Validation
	// msg marshals PatternFlowLacpPartnerSystemIdCounter to protobuf object *otg.PatternFlowLacpPartnerSystemIdCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerSystemIdCounter
	// setMsg unmarshals PatternFlowLacpPartnerSystemIdCounter from protobuf object *otg.PatternFlowLacpPartnerSystemIdCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerSystemIdCounter) PatternFlowLacpPartnerSystemIdCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerSystemIdCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerSystemIdCounter
	// validate validates PatternFlowLacpPartnerSystemIdCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerSystemIdCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowLacpPartnerSystemIdCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowLacpPartnerSystemIdCounter
	SetStart(value string) PatternFlowLacpPartnerSystemIdCounter
	// HasStart checks if Start has been set in PatternFlowLacpPartnerSystemIdCounter
	HasStart() bool
	// Step returns string, set in PatternFlowLacpPartnerSystemIdCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowLacpPartnerSystemIdCounter
	SetStep(value string) PatternFlowLacpPartnerSystemIdCounter
	// HasStep checks if Step has been set in PatternFlowLacpPartnerSystemIdCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpPartnerSystemIdCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpPartnerSystemIdCounter
	SetCount(value uint32) PatternFlowLacpPartnerSystemIdCounter
	// HasCount checks if Count has been set in PatternFlowLacpPartnerSystemIdCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowLacpPartnerSystemIdCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowLacpPartnerSystemIdCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowLacpPartnerSystemIdCounter object
func (obj *patternFlowLacpPartnerSystemIdCounter) SetStart(value string) PatternFlowLacpPartnerSystemIdCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowLacpPartnerSystemIdCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowLacpPartnerSystemIdCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowLacpPartnerSystemIdCounter object
func (obj *patternFlowLacpPartnerSystemIdCounter) SetStep(value string) PatternFlowLacpPartnerSystemIdCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerSystemIdCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerSystemIdCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpPartnerSystemIdCounter object
func (obj *patternFlowLacpPartnerSystemIdCounter) SetCount(value uint32) PatternFlowLacpPartnerSystemIdCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpPartnerSystemIdCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowLacpPartnerSystemIdCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowLacpPartnerSystemIdCounter.Step"))
		}

	}

}

func (obj *patternFlowLacpPartnerSystemIdCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("00:00:00:00:00:00")
	}
	if obj.obj.Step == nil {
		obj.SetStep("00:00:00:00:00:01")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

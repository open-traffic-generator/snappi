package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerSystemIdCounter *****
type patternFlowLacpduPartnerSystemIdCounter struct {
	validation
	obj          *otg.PatternFlowLacpduPartnerSystemIdCounter
	marshaller   marshalPatternFlowLacpduPartnerSystemIdCounter
	unMarshaller unMarshalPatternFlowLacpduPartnerSystemIdCounter
}

func NewPatternFlowLacpduPartnerSystemIdCounter() PatternFlowLacpduPartnerSystemIdCounter {
	obj := patternFlowLacpduPartnerSystemIdCounter{obj: &otg.PatternFlowLacpduPartnerSystemIdCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerSystemIdCounter) msg() *otg.PatternFlowLacpduPartnerSystemIdCounter {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerSystemIdCounter) setMsg(msg *otg.PatternFlowLacpduPartnerSystemIdCounter) PatternFlowLacpduPartnerSystemIdCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerSystemIdCounter struct {
	obj *patternFlowLacpduPartnerSystemIdCounter
}

type marshalPatternFlowLacpduPartnerSystemIdCounter interface {
	// ToProto marshals PatternFlowLacpduPartnerSystemIdCounter to protobuf object *otg.PatternFlowLacpduPartnerSystemIdCounter
	ToProto() (*otg.PatternFlowLacpduPartnerSystemIdCounter, error)
	// ToPbText marshals PatternFlowLacpduPartnerSystemIdCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerSystemIdCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerSystemIdCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerSystemIdCounter struct {
	obj *patternFlowLacpduPartnerSystemIdCounter
}

type unMarshalPatternFlowLacpduPartnerSystemIdCounter interface {
	// FromProto unmarshals PatternFlowLacpduPartnerSystemIdCounter from protobuf object *otg.PatternFlowLacpduPartnerSystemIdCounter
	FromProto(msg *otg.PatternFlowLacpduPartnerSystemIdCounter) (PatternFlowLacpduPartnerSystemIdCounter, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerSystemIdCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerSystemIdCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerSystemIdCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerSystemIdCounter) Marshal() marshalPatternFlowLacpduPartnerSystemIdCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerSystemIdCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerSystemIdCounter) Unmarshal() unMarshalPatternFlowLacpduPartnerSystemIdCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerSystemIdCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerSystemIdCounter) ToProto() (*otg.PatternFlowLacpduPartnerSystemIdCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerSystemIdCounter) FromProto(msg *otg.PatternFlowLacpduPartnerSystemIdCounter) (PatternFlowLacpduPartnerSystemIdCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerSystemIdCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerSystemIdCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerSystemIdCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerSystemIdCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerSystemIdCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerSystemIdCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerSystemIdCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerSystemIdCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerSystemIdCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerSystemIdCounter) Clone() (PatternFlowLacpduPartnerSystemIdCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerSystemIdCounter()
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

// PatternFlowLacpduPartnerSystemIdCounter is mac counter pattern
type PatternFlowLacpduPartnerSystemIdCounter interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerSystemIdCounter to protobuf object *otg.PatternFlowLacpduPartnerSystemIdCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerSystemIdCounter
	// setMsg unmarshals PatternFlowLacpduPartnerSystemIdCounter from protobuf object *otg.PatternFlowLacpduPartnerSystemIdCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerSystemIdCounter) PatternFlowLacpduPartnerSystemIdCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerSystemIdCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerSystemIdCounter
	// validate validates PatternFlowLacpduPartnerSystemIdCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerSystemIdCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowLacpduPartnerSystemIdCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowLacpduPartnerSystemIdCounter
	SetStart(value string) PatternFlowLacpduPartnerSystemIdCounter
	// HasStart checks if Start has been set in PatternFlowLacpduPartnerSystemIdCounter
	HasStart() bool
	// Step returns string, set in PatternFlowLacpduPartnerSystemIdCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowLacpduPartnerSystemIdCounter
	SetStep(value string) PatternFlowLacpduPartnerSystemIdCounter
	// HasStep checks if Step has been set in PatternFlowLacpduPartnerSystemIdCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduPartnerSystemIdCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduPartnerSystemIdCounter
	SetCount(value uint32) PatternFlowLacpduPartnerSystemIdCounter
	// HasCount checks if Count has been set in PatternFlowLacpduPartnerSystemIdCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowLacpduPartnerSystemIdCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowLacpduPartnerSystemIdCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowLacpduPartnerSystemIdCounter object
func (obj *patternFlowLacpduPartnerSystemIdCounter) SetStart(value string) PatternFlowLacpduPartnerSystemIdCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowLacpduPartnerSystemIdCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowLacpduPartnerSystemIdCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowLacpduPartnerSystemIdCounter object
func (obj *patternFlowLacpduPartnerSystemIdCounter) SetStep(value string) PatternFlowLacpduPartnerSystemIdCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerSystemIdCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerSystemIdCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduPartnerSystemIdCounter object
func (obj *patternFlowLacpduPartnerSystemIdCounter) SetCount(value uint32) PatternFlowLacpduPartnerSystemIdCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduPartnerSystemIdCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowLacpduPartnerSystemIdCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowLacpduPartnerSystemIdCounter.Step"))
		}

	}

}

func (obj *patternFlowLacpduPartnerSystemIdCounter) setDefault() {
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

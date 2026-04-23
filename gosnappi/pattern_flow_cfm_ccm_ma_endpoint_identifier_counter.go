package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmCcmMaEndpointIdentifierCounter *****
type patternFlowCfmCcmMaEndpointIdentifierCounter struct {
	validation
	obj          *otg.PatternFlowCfmCcmMaEndpointIdentifierCounter
	marshaller   marshalPatternFlowCfmCcmMaEndpointIdentifierCounter
	unMarshaller unMarshalPatternFlowCfmCcmMaEndpointIdentifierCounter
}

func NewPatternFlowCfmCcmMaEndpointIdentifierCounter() PatternFlowCfmCcmMaEndpointIdentifierCounter {
	obj := patternFlowCfmCcmMaEndpointIdentifierCounter{obj: &otg.PatternFlowCfmCcmMaEndpointIdentifierCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmCcmMaEndpointIdentifierCounter) msg() *otg.PatternFlowCfmCcmMaEndpointIdentifierCounter {
	return obj.obj
}

func (obj *patternFlowCfmCcmMaEndpointIdentifierCounter) setMsg(msg *otg.PatternFlowCfmCcmMaEndpointIdentifierCounter) PatternFlowCfmCcmMaEndpointIdentifierCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmCcmMaEndpointIdentifierCounter struct {
	obj *patternFlowCfmCcmMaEndpointIdentifierCounter
}

type marshalPatternFlowCfmCcmMaEndpointIdentifierCounter interface {
	// ToProto marshals PatternFlowCfmCcmMaEndpointIdentifierCounter to protobuf object *otg.PatternFlowCfmCcmMaEndpointIdentifierCounter
	ToProto() (*otg.PatternFlowCfmCcmMaEndpointIdentifierCounter, error)
	// ToPbText marshals PatternFlowCfmCcmMaEndpointIdentifierCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmCcmMaEndpointIdentifierCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmCcmMaEndpointIdentifierCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmCcmMaEndpointIdentifierCounter struct {
	obj *patternFlowCfmCcmMaEndpointIdentifierCounter
}

type unMarshalPatternFlowCfmCcmMaEndpointIdentifierCounter interface {
	// FromProto unmarshals PatternFlowCfmCcmMaEndpointIdentifierCounter from protobuf object *otg.PatternFlowCfmCcmMaEndpointIdentifierCounter
	FromProto(msg *otg.PatternFlowCfmCcmMaEndpointIdentifierCounter) (PatternFlowCfmCcmMaEndpointIdentifierCounter, error)
	// FromPbText unmarshals PatternFlowCfmCcmMaEndpointIdentifierCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmCcmMaEndpointIdentifierCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmCcmMaEndpointIdentifierCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmCcmMaEndpointIdentifierCounter) Marshal() marshalPatternFlowCfmCcmMaEndpointIdentifierCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmCcmMaEndpointIdentifierCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmCcmMaEndpointIdentifierCounter) Unmarshal() unMarshalPatternFlowCfmCcmMaEndpointIdentifierCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmCcmMaEndpointIdentifierCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmCcmMaEndpointIdentifierCounter) ToProto() (*otg.PatternFlowCfmCcmMaEndpointIdentifierCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmCcmMaEndpointIdentifierCounter) FromProto(msg *otg.PatternFlowCfmCcmMaEndpointIdentifierCounter) (PatternFlowCfmCcmMaEndpointIdentifierCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmCcmMaEndpointIdentifierCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmCcmMaEndpointIdentifierCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmCcmMaEndpointIdentifierCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmCcmMaEndpointIdentifierCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmCcmMaEndpointIdentifierCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmCcmMaEndpointIdentifierCounter) FromJson(value string) error {
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

func (obj *patternFlowCfmCcmMaEndpointIdentifierCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmCcmMaEndpointIdentifierCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmCcmMaEndpointIdentifierCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmCcmMaEndpointIdentifierCounter) Clone() (PatternFlowCfmCcmMaEndpointIdentifierCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmCcmMaEndpointIdentifierCounter()
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

// PatternFlowCfmCcmMaEndpointIdentifierCounter is integer counter pattern
type PatternFlowCfmCcmMaEndpointIdentifierCounter interface {
	Validation
	// msg marshals PatternFlowCfmCcmMaEndpointIdentifierCounter to protobuf object *otg.PatternFlowCfmCcmMaEndpointIdentifierCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmCcmMaEndpointIdentifierCounter
	// setMsg unmarshals PatternFlowCfmCcmMaEndpointIdentifierCounter from protobuf object *otg.PatternFlowCfmCcmMaEndpointIdentifierCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmCcmMaEndpointIdentifierCounter) PatternFlowCfmCcmMaEndpointIdentifierCounter
	// provides marshal interface
	Marshal() marshalPatternFlowCfmCcmMaEndpointIdentifierCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmCcmMaEndpointIdentifierCounter
	// validate validates PatternFlowCfmCcmMaEndpointIdentifierCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmCcmMaEndpointIdentifierCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowCfmCcmMaEndpointIdentifierCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowCfmCcmMaEndpointIdentifierCounter
	SetStart(value uint32) PatternFlowCfmCcmMaEndpointIdentifierCounter
	// HasStart checks if Start has been set in PatternFlowCfmCcmMaEndpointIdentifierCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowCfmCcmMaEndpointIdentifierCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowCfmCcmMaEndpointIdentifierCounter
	SetStep(value uint32) PatternFlowCfmCcmMaEndpointIdentifierCounter
	// HasStep checks if Step has been set in PatternFlowCfmCcmMaEndpointIdentifierCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowCfmCcmMaEndpointIdentifierCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowCfmCcmMaEndpointIdentifierCounter
	SetCount(value uint32) PatternFlowCfmCcmMaEndpointIdentifierCounter
	// HasCount checks if Count has been set in PatternFlowCfmCcmMaEndpointIdentifierCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmCcmMaEndpointIdentifierCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmCcmMaEndpointIdentifierCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowCfmCcmMaEndpointIdentifierCounter object
func (obj *patternFlowCfmCcmMaEndpointIdentifierCounter) SetStart(value uint32) PatternFlowCfmCcmMaEndpointIdentifierCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmCcmMaEndpointIdentifierCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmCcmMaEndpointIdentifierCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowCfmCcmMaEndpointIdentifierCounter object
func (obj *patternFlowCfmCcmMaEndpointIdentifierCounter) SetStep(value uint32) PatternFlowCfmCcmMaEndpointIdentifierCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmCcmMaEndpointIdentifierCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmCcmMaEndpointIdentifierCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowCfmCcmMaEndpointIdentifierCounter object
func (obj *patternFlowCfmCcmMaEndpointIdentifierCounter) SetCount(value uint32) PatternFlowCfmCcmMaEndpointIdentifierCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowCfmCcmMaEndpointIdentifierCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmCcmMaEndpointIdentifierCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmCcmMaEndpointIdentifierCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmCcmMaEndpointIdentifierCounter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowCfmCcmMaEndpointIdentifierCounter) setDefault() {
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

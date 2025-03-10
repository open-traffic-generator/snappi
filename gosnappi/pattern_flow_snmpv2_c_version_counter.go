package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CVersionCounter *****
type patternFlowSnmpv2CVersionCounter struct {
	validation
	obj          *otg.PatternFlowSnmpv2CVersionCounter
	marshaller   marshalPatternFlowSnmpv2CVersionCounter
	unMarshaller unMarshalPatternFlowSnmpv2CVersionCounter
}

func NewPatternFlowSnmpv2CVersionCounter() PatternFlowSnmpv2CVersionCounter {
	obj := patternFlowSnmpv2CVersionCounter{obj: &otg.PatternFlowSnmpv2CVersionCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CVersionCounter) msg() *otg.PatternFlowSnmpv2CVersionCounter {
	return obj.obj
}

func (obj *patternFlowSnmpv2CVersionCounter) setMsg(msg *otg.PatternFlowSnmpv2CVersionCounter) PatternFlowSnmpv2CVersionCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CVersionCounter struct {
	obj *patternFlowSnmpv2CVersionCounter
}

type marshalPatternFlowSnmpv2CVersionCounter interface {
	// ToProto marshals PatternFlowSnmpv2CVersionCounter to protobuf object *otg.PatternFlowSnmpv2CVersionCounter
	ToProto() (*otg.PatternFlowSnmpv2CVersionCounter, error)
	// ToPbText marshals PatternFlowSnmpv2CVersionCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CVersionCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CVersionCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowSnmpv2CVersionCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowSnmpv2CVersionCounter struct {
	obj *patternFlowSnmpv2CVersionCounter
}

type unMarshalPatternFlowSnmpv2CVersionCounter interface {
	// FromProto unmarshals PatternFlowSnmpv2CVersionCounter from protobuf object *otg.PatternFlowSnmpv2CVersionCounter
	FromProto(msg *otg.PatternFlowSnmpv2CVersionCounter) (PatternFlowSnmpv2CVersionCounter, error)
	// FromPbText unmarshals PatternFlowSnmpv2CVersionCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CVersionCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CVersionCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CVersionCounter) Marshal() marshalPatternFlowSnmpv2CVersionCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CVersionCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CVersionCounter) Unmarshal() unMarshalPatternFlowSnmpv2CVersionCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CVersionCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CVersionCounter) ToProto() (*otg.PatternFlowSnmpv2CVersionCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CVersionCounter) FromProto(msg *otg.PatternFlowSnmpv2CVersionCounter) (PatternFlowSnmpv2CVersionCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CVersionCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVersionCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVersionCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVersionCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVersionCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowSnmpv2CVersionCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVersionCounter) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CVersionCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVersionCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVersionCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CVersionCounter) Clone() (PatternFlowSnmpv2CVersionCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CVersionCounter()
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

// PatternFlowSnmpv2CVersionCounter is integer counter pattern
type PatternFlowSnmpv2CVersionCounter interface {
	Validation
	// msg marshals PatternFlowSnmpv2CVersionCounter to protobuf object *otg.PatternFlowSnmpv2CVersionCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CVersionCounter
	// setMsg unmarshals PatternFlowSnmpv2CVersionCounter from protobuf object *otg.PatternFlowSnmpv2CVersionCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CVersionCounter) PatternFlowSnmpv2CVersionCounter
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CVersionCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CVersionCounter
	// validate validates PatternFlowSnmpv2CVersionCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CVersionCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowSnmpv2CVersionCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowSnmpv2CVersionCounter
	SetStart(value uint32) PatternFlowSnmpv2CVersionCounter
	// HasStart checks if Start has been set in PatternFlowSnmpv2CVersionCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowSnmpv2CVersionCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowSnmpv2CVersionCounter
	SetStep(value uint32) PatternFlowSnmpv2CVersionCounter
	// HasStep checks if Step has been set in PatternFlowSnmpv2CVersionCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowSnmpv2CVersionCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowSnmpv2CVersionCounter
	SetCount(value uint32) PatternFlowSnmpv2CVersionCounter
	// HasCount checks if Count has been set in PatternFlowSnmpv2CVersionCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowSnmpv2CVersionCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowSnmpv2CVersionCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowSnmpv2CVersionCounter object
func (obj *patternFlowSnmpv2CVersionCounter) SetStart(value uint32) PatternFlowSnmpv2CVersionCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowSnmpv2CVersionCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowSnmpv2CVersionCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowSnmpv2CVersionCounter object
func (obj *patternFlowSnmpv2CVersionCounter) SetStep(value uint32) PatternFlowSnmpv2CVersionCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowSnmpv2CVersionCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowSnmpv2CVersionCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowSnmpv2CVersionCounter object
func (obj *patternFlowSnmpv2CVersionCounter) SetCount(value uint32) PatternFlowSnmpv2CVersionCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowSnmpv2CVersionCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowSnmpv2CVersionCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowSnmpv2CVersionCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowSnmpv2CVersionCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowSnmpv2CVersionCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(1)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

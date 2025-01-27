package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CPDURequestIdCounter *****
type patternFlowSnmpv2CPDURequestIdCounter struct {
	validation
	obj          *otg.PatternFlowSnmpv2CPDURequestIdCounter
	marshaller   marshalPatternFlowSnmpv2CPDURequestIdCounter
	unMarshaller unMarshalPatternFlowSnmpv2CPDURequestIdCounter
}

func NewPatternFlowSnmpv2CPDURequestIdCounter() PatternFlowSnmpv2CPDURequestIdCounter {
	obj := patternFlowSnmpv2CPDURequestIdCounter{obj: &otg.PatternFlowSnmpv2CPDURequestIdCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CPDURequestIdCounter) msg() *otg.PatternFlowSnmpv2CPDURequestIdCounter {
	return obj.obj
}

func (obj *patternFlowSnmpv2CPDURequestIdCounter) setMsg(msg *otg.PatternFlowSnmpv2CPDURequestIdCounter) PatternFlowSnmpv2CPDURequestIdCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CPDURequestIdCounter struct {
	obj *patternFlowSnmpv2CPDURequestIdCounter
}

type marshalPatternFlowSnmpv2CPDURequestIdCounter interface {
	// ToProto marshals PatternFlowSnmpv2CPDURequestIdCounter to protobuf object *otg.PatternFlowSnmpv2CPDURequestIdCounter
	ToProto() (*otg.PatternFlowSnmpv2CPDURequestIdCounter, error)
	// ToPbText marshals PatternFlowSnmpv2CPDURequestIdCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CPDURequestIdCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CPDURequestIdCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowSnmpv2CPDURequestIdCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowSnmpv2CPDURequestIdCounter struct {
	obj *patternFlowSnmpv2CPDURequestIdCounter
}

type unMarshalPatternFlowSnmpv2CPDURequestIdCounter interface {
	// FromProto unmarshals PatternFlowSnmpv2CPDURequestIdCounter from protobuf object *otg.PatternFlowSnmpv2CPDURequestIdCounter
	FromProto(msg *otg.PatternFlowSnmpv2CPDURequestIdCounter) (PatternFlowSnmpv2CPDURequestIdCounter, error)
	// FromPbText unmarshals PatternFlowSnmpv2CPDURequestIdCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CPDURequestIdCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CPDURequestIdCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CPDURequestIdCounter) Marshal() marshalPatternFlowSnmpv2CPDURequestIdCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CPDURequestIdCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CPDURequestIdCounter) Unmarshal() unMarshalPatternFlowSnmpv2CPDURequestIdCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CPDURequestIdCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CPDURequestIdCounter) ToProto() (*otg.PatternFlowSnmpv2CPDURequestIdCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CPDURequestIdCounter) FromProto(msg *otg.PatternFlowSnmpv2CPDURequestIdCounter) (PatternFlowSnmpv2CPDURequestIdCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CPDURequestIdCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CPDURequestIdCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CPDURequestIdCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CPDURequestIdCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CPDURequestIdCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowSnmpv2CPDURequestIdCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CPDURequestIdCounter) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CPDURequestIdCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CPDURequestIdCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CPDURequestIdCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CPDURequestIdCounter) Clone() (PatternFlowSnmpv2CPDURequestIdCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CPDURequestIdCounter()
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

// PatternFlowSnmpv2CPDURequestIdCounter is integer counter pattern
type PatternFlowSnmpv2CPDURequestIdCounter interface {
	Validation
	// msg marshals PatternFlowSnmpv2CPDURequestIdCounter to protobuf object *otg.PatternFlowSnmpv2CPDURequestIdCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CPDURequestIdCounter
	// setMsg unmarshals PatternFlowSnmpv2CPDURequestIdCounter from protobuf object *otg.PatternFlowSnmpv2CPDURequestIdCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CPDURequestIdCounter) PatternFlowSnmpv2CPDURequestIdCounter
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CPDURequestIdCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CPDURequestIdCounter
	// validate validates PatternFlowSnmpv2CPDURequestIdCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CPDURequestIdCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns int32, set in PatternFlowSnmpv2CPDURequestIdCounter.
	Start() int32
	// SetStart assigns int32 provided by user to PatternFlowSnmpv2CPDURequestIdCounter
	SetStart(value int32) PatternFlowSnmpv2CPDURequestIdCounter
	// HasStart checks if Start has been set in PatternFlowSnmpv2CPDURequestIdCounter
	HasStart() bool
	// Step returns int32, set in PatternFlowSnmpv2CPDURequestIdCounter.
	Step() int32
	// SetStep assigns int32 provided by user to PatternFlowSnmpv2CPDURequestIdCounter
	SetStep(value int32) PatternFlowSnmpv2CPDURequestIdCounter
	// HasStep checks if Step has been set in PatternFlowSnmpv2CPDURequestIdCounter
	HasStep() bool
	// Count returns int32, set in PatternFlowSnmpv2CPDURequestIdCounter.
	Count() int32
	// SetCount assigns int32 provided by user to PatternFlowSnmpv2CPDURequestIdCounter
	SetCount(value int32) PatternFlowSnmpv2CPDURequestIdCounter
	// HasCount checks if Count has been set in PatternFlowSnmpv2CPDURequestIdCounter
	HasCount() bool
}

// description is TBD
// Start returns a int32
func (obj *patternFlowSnmpv2CPDURequestIdCounter) Start() int32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a int32
func (obj *patternFlowSnmpv2CPDURequestIdCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the int32 value in the PatternFlowSnmpv2CPDURequestIdCounter object
func (obj *patternFlowSnmpv2CPDURequestIdCounter) SetStart(value int32) PatternFlowSnmpv2CPDURequestIdCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a int32
func (obj *patternFlowSnmpv2CPDURequestIdCounter) Step() int32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a int32
func (obj *patternFlowSnmpv2CPDURequestIdCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the int32 value in the PatternFlowSnmpv2CPDURequestIdCounter object
func (obj *patternFlowSnmpv2CPDURequestIdCounter) SetStep(value int32) PatternFlowSnmpv2CPDURequestIdCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a int32
func (obj *patternFlowSnmpv2CPDURequestIdCounter) Count() int32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a int32
func (obj *patternFlowSnmpv2CPDURequestIdCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the int32 value in the PatternFlowSnmpv2CPDURequestIdCounter object
func (obj *patternFlowSnmpv2CPDURequestIdCounter) SetCount(value int32) PatternFlowSnmpv2CPDURequestIdCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowSnmpv2CPDURequestIdCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowSnmpv2CPDURequestIdCounter) setDefault() {
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

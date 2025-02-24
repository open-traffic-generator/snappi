package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CBulkPDURequestIdCounter *****
type patternFlowSnmpv2CBulkPDURequestIdCounter struct {
	validation
	obj          *otg.PatternFlowSnmpv2CBulkPDURequestIdCounter
	marshaller   marshalPatternFlowSnmpv2CBulkPDURequestIdCounter
	unMarshaller unMarshalPatternFlowSnmpv2CBulkPDURequestIdCounter
}

func NewPatternFlowSnmpv2CBulkPDURequestIdCounter() PatternFlowSnmpv2CBulkPDURequestIdCounter {
	obj := patternFlowSnmpv2CBulkPDURequestIdCounter{obj: &otg.PatternFlowSnmpv2CBulkPDURequestIdCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CBulkPDURequestIdCounter) msg() *otg.PatternFlowSnmpv2CBulkPDURequestIdCounter {
	return obj.obj
}

func (obj *patternFlowSnmpv2CBulkPDURequestIdCounter) setMsg(msg *otg.PatternFlowSnmpv2CBulkPDURequestIdCounter) PatternFlowSnmpv2CBulkPDURequestIdCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CBulkPDURequestIdCounter struct {
	obj *patternFlowSnmpv2CBulkPDURequestIdCounter
}

type marshalPatternFlowSnmpv2CBulkPDURequestIdCounter interface {
	// ToProto marshals PatternFlowSnmpv2CBulkPDURequestIdCounter to protobuf object *otg.PatternFlowSnmpv2CBulkPDURequestIdCounter
	ToProto() (*otg.PatternFlowSnmpv2CBulkPDURequestIdCounter, error)
	// ToPbText marshals PatternFlowSnmpv2CBulkPDURequestIdCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CBulkPDURequestIdCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CBulkPDURequestIdCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowSnmpv2CBulkPDURequestIdCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowSnmpv2CBulkPDURequestIdCounter struct {
	obj *patternFlowSnmpv2CBulkPDURequestIdCounter
}

type unMarshalPatternFlowSnmpv2CBulkPDURequestIdCounter interface {
	// FromProto unmarshals PatternFlowSnmpv2CBulkPDURequestIdCounter from protobuf object *otg.PatternFlowSnmpv2CBulkPDURequestIdCounter
	FromProto(msg *otg.PatternFlowSnmpv2CBulkPDURequestIdCounter) (PatternFlowSnmpv2CBulkPDURequestIdCounter, error)
	// FromPbText unmarshals PatternFlowSnmpv2CBulkPDURequestIdCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CBulkPDURequestIdCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CBulkPDURequestIdCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CBulkPDURequestIdCounter) Marshal() marshalPatternFlowSnmpv2CBulkPDURequestIdCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CBulkPDURequestIdCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CBulkPDURequestIdCounter) Unmarshal() unMarshalPatternFlowSnmpv2CBulkPDURequestIdCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CBulkPDURequestIdCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CBulkPDURequestIdCounter) ToProto() (*otg.PatternFlowSnmpv2CBulkPDURequestIdCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CBulkPDURequestIdCounter) FromProto(msg *otg.PatternFlowSnmpv2CBulkPDURequestIdCounter) (PatternFlowSnmpv2CBulkPDURequestIdCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CBulkPDURequestIdCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CBulkPDURequestIdCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CBulkPDURequestIdCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CBulkPDURequestIdCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CBulkPDURequestIdCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowSnmpv2CBulkPDURequestIdCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CBulkPDURequestIdCounter) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CBulkPDURequestIdCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CBulkPDURequestIdCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CBulkPDURequestIdCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CBulkPDURequestIdCounter) Clone() (PatternFlowSnmpv2CBulkPDURequestIdCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CBulkPDURequestIdCounter()
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

// PatternFlowSnmpv2CBulkPDURequestIdCounter is integer counter pattern
type PatternFlowSnmpv2CBulkPDURequestIdCounter interface {
	Validation
	// msg marshals PatternFlowSnmpv2CBulkPDURequestIdCounter to protobuf object *otg.PatternFlowSnmpv2CBulkPDURequestIdCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CBulkPDURequestIdCounter
	// setMsg unmarshals PatternFlowSnmpv2CBulkPDURequestIdCounter from protobuf object *otg.PatternFlowSnmpv2CBulkPDURequestIdCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CBulkPDURequestIdCounter) PatternFlowSnmpv2CBulkPDURequestIdCounter
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CBulkPDURequestIdCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CBulkPDURequestIdCounter
	// validate validates PatternFlowSnmpv2CBulkPDURequestIdCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CBulkPDURequestIdCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns int32, set in PatternFlowSnmpv2CBulkPDURequestIdCounter.
	Start() int32
	// SetStart assigns int32 provided by user to PatternFlowSnmpv2CBulkPDURequestIdCounter
	SetStart(value int32) PatternFlowSnmpv2CBulkPDURequestIdCounter
	// HasStart checks if Start has been set in PatternFlowSnmpv2CBulkPDURequestIdCounter
	HasStart() bool
	// Step returns int32, set in PatternFlowSnmpv2CBulkPDURequestIdCounter.
	Step() int32
	// SetStep assigns int32 provided by user to PatternFlowSnmpv2CBulkPDURequestIdCounter
	SetStep(value int32) PatternFlowSnmpv2CBulkPDURequestIdCounter
	// HasStep checks if Step has been set in PatternFlowSnmpv2CBulkPDURequestIdCounter
	HasStep() bool
	// Count returns int32, set in PatternFlowSnmpv2CBulkPDURequestIdCounter.
	Count() int32
	// SetCount assigns int32 provided by user to PatternFlowSnmpv2CBulkPDURequestIdCounter
	SetCount(value int32) PatternFlowSnmpv2CBulkPDURequestIdCounter
	// HasCount checks if Count has been set in PatternFlowSnmpv2CBulkPDURequestIdCounter
	HasCount() bool
}

// description is TBD
// Start returns a int32
func (obj *patternFlowSnmpv2CBulkPDURequestIdCounter) Start() int32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a int32
func (obj *patternFlowSnmpv2CBulkPDURequestIdCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the int32 value in the PatternFlowSnmpv2CBulkPDURequestIdCounter object
func (obj *patternFlowSnmpv2CBulkPDURequestIdCounter) SetStart(value int32) PatternFlowSnmpv2CBulkPDURequestIdCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a int32
func (obj *patternFlowSnmpv2CBulkPDURequestIdCounter) Step() int32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a int32
func (obj *patternFlowSnmpv2CBulkPDURequestIdCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the int32 value in the PatternFlowSnmpv2CBulkPDURequestIdCounter object
func (obj *patternFlowSnmpv2CBulkPDURequestIdCounter) SetStep(value int32) PatternFlowSnmpv2CBulkPDURequestIdCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a int32
func (obj *patternFlowSnmpv2CBulkPDURequestIdCounter) Count() int32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a int32
func (obj *patternFlowSnmpv2CBulkPDURequestIdCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the int32 value in the PatternFlowSnmpv2CBulkPDURequestIdCounter object
func (obj *patternFlowSnmpv2CBulkPDURequestIdCounter) SetCount(value int32) PatternFlowSnmpv2CBulkPDURequestIdCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowSnmpv2CBulkPDURequestIdCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowSnmpv2CBulkPDURequestIdCounter) setDefault() {
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

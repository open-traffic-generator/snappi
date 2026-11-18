package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter *****
type patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter struct {
	validation
	obj          *otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
	marshaller   marshalPatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
	unMarshaller unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
}

func NewPatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter() PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter {
	obj := patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter{obj: &otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) setMsg(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
}

type marshalPatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter interface {
	// ToProto marshals PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
	ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter, error)
	// ToPbText marshals PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
}

type unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
	FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) (PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) (PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter()
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

// PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter is integer counter pattern
type PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
	// setMsg unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
	// validate validates PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
	SetStart(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
	// HasStart checks if Start has been set in PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
	SetStep(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
	// HasStep checks if Step has been set in PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
	SetCount(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
	// HasCount checks if Count has been set in PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) SetStart(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) SetStep(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) SetCount(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved2Counter) setDefault() {
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

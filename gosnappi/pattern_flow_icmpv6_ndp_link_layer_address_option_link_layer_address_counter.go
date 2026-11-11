package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter *****
type patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	marshaller   marshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	unMarshaller unMarshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
}

func NewPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter() PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter {
	obj := patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter{obj: &otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) msg() *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) setMsg(msg *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter struct {
	obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
}

type marshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter interface {
	// ToProto marshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter to protobuf object *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	ToProto() (*otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter, error)
	// ToPbText marshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter struct {
	obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
}

type unMarshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter from protobuf object *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	FromProto(msg *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) (PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) Marshal() marshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) Unmarshal() unMarshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) ToProto() (*otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) FromProto(msg *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) (PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) Clone() (PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter()
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

// PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter is mac counter pattern
type PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter to protobuf object *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	// setMsg unmarshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter from protobuf object *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	// validate validates PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	SetStart(value string) PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	SetStep(value string) PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	SetCount(value uint32) PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter object
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) SetStart(value string) PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter object
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) SetStep(value string) PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter object
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) SetCount(value uint32) PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter.Step"))
		}

	}

}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) setDefault() {
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

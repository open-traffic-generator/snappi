package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter *****
type patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter struct {
	validation
	obj          *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
	marshaller   marshalPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
	unMarshaller unMarshalPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
}

func NewPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter() PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter {
	obj := patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter{obj: &otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) msg() *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) setMsg(msg *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter struct {
	obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
}

type marshalPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter interface {
	// ToProto marshals PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter to protobuf object *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
	ToProto() (*otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter, error)
	// ToPbText marshals PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter struct {
	obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
}

type unMarshalPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter from protobuf object *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
	FromProto(msg *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) (PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) Marshal() marshalPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) Unmarshal() unMarshalPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) ToProto() (*otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) FromProto(msg *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) (PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) Clone() (PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter()
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

// PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter is integer counter pattern
type PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter to protobuf object *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
	// setMsg unmarshals PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter from protobuf object *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
	// validate validates PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
	SetStart(value uint32) PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
	// HasStart checks if Start has been set in PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
	SetStep(value uint32) PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
	// HasStep checks if Step has been set in PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
	SetCount(value uint32) PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
	// HasCount checks if Count has been set in PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter object
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) SetStart(value uint32) PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter object
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) SetStep(value uint32) PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter object
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) SetCount(value uint32) PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1Counter) setDefault() {
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

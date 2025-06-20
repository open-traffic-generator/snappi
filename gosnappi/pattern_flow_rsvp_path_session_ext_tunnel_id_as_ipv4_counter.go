package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter *****
type patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter struct {
	validation
	obj          *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	marshaller   marshalPatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	unMarshaller unMarshalPatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
}

func NewPatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter() PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter {
	obj := patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter{obj: &otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) msg() *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter {
	return obj.obj
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) setMsg(msg *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter struct {
	obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
}

type marshalPatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter interface {
	// ToProto marshals PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter to protobuf object *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	ToProto() (*otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter, error)
	// ToPbText marshals PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter struct {
	obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
}

type unMarshalPatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter interface {
	// FromProto unmarshals PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter from protobuf object *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	FromProto(msg *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) (PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter, error)
	// FromPbText unmarshals PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) Marshal() marshalPatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) Unmarshal() unMarshalPatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) ToProto() (*otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) FromProto(msg *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) (PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) Clone() (PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter()
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

// PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter is ipv4 counter pattern
type PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter interface {
	Validation
	// msg marshals PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter to protobuf object *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	// setMsg unmarshals PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter from protobuf object *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	// validate validates PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	SetStart(value string) PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	// HasStart checks if Start has been set in PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	HasStart() bool
	// Step returns string, set in PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	SetStep(value string) PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	// HasStep checks if Step has been set in PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	SetCount(value uint32) PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	// HasCount checks if Count has been set in PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter object
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) SetStart(value string) PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter object
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) SetStep(value string) PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter object
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) SetCount(value uint32) PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter.Step"))
		}

	}

}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("0.0.0.0")
	}
	if obj.obj.Step == nil {
		obj.SetStep("0.0.0.1")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

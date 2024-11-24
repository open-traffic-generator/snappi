package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpCtlSynCounter *****
type patternFlowTcpCtlSynCounter struct {
	validation
	obj          *otg.PatternFlowTcpCtlSynCounter
	marshaller   marshalPatternFlowTcpCtlSynCounter
	unMarshaller unMarshalPatternFlowTcpCtlSynCounter
}

func NewPatternFlowTcpCtlSynCounter() PatternFlowTcpCtlSynCounter {
	obj := patternFlowTcpCtlSynCounter{obj: &otg.PatternFlowTcpCtlSynCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpCtlSynCounter) msg() *otg.PatternFlowTcpCtlSynCounter {
	return obj.obj
}

func (obj *patternFlowTcpCtlSynCounter) setMsg(msg *otg.PatternFlowTcpCtlSynCounter) PatternFlowTcpCtlSynCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpCtlSynCounter struct {
	obj *patternFlowTcpCtlSynCounter
}

type marshalPatternFlowTcpCtlSynCounter interface {
	// ToProto marshals PatternFlowTcpCtlSynCounter to protobuf object *otg.PatternFlowTcpCtlSynCounter
	ToProto() (*otg.PatternFlowTcpCtlSynCounter, error)
	// ToPbText marshals PatternFlowTcpCtlSynCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpCtlSynCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpCtlSynCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowTcpCtlSynCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowTcpCtlSynCounter struct {
	obj *patternFlowTcpCtlSynCounter
}

type unMarshalPatternFlowTcpCtlSynCounter interface {
	// FromProto unmarshals PatternFlowTcpCtlSynCounter from protobuf object *otg.PatternFlowTcpCtlSynCounter
	FromProto(msg *otg.PatternFlowTcpCtlSynCounter) (PatternFlowTcpCtlSynCounter, error)
	// FromPbText unmarshals PatternFlowTcpCtlSynCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpCtlSynCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpCtlSynCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpCtlSynCounter) Marshal() marshalPatternFlowTcpCtlSynCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpCtlSynCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpCtlSynCounter) Unmarshal() unMarshalPatternFlowTcpCtlSynCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpCtlSynCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpCtlSynCounter) ToProto() (*otg.PatternFlowTcpCtlSynCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpCtlSynCounter) FromProto(msg *otg.PatternFlowTcpCtlSynCounter) (PatternFlowTcpCtlSynCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpCtlSynCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlSynCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpCtlSynCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlSynCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpCtlSynCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowTcpCtlSynCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlSynCounter) FromJson(value string) error {
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

func (obj *patternFlowTcpCtlSynCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlSynCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlSynCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpCtlSynCounter) Clone() (PatternFlowTcpCtlSynCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpCtlSynCounter()
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

// PatternFlowTcpCtlSynCounter is integer counter pattern
type PatternFlowTcpCtlSynCounter interface {
	Validation
	// msg marshals PatternFlowTcpCtlSynCounter to protobuf object *otg.PatternFlowTcpCtlSynCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpCtlSynCounter
	// setMsg unmarshals PatternFlowTcpCtlSynCounter from protobuf object *otg.PatternFlowTcpCtlSynCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpCtlSynCounter) PatternFlowTcpCtlSynCounter
	// provides marshal interface
	Marshal() marshalPatternFlowTcpCtlSynCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpCtlSynCounter
	// validate validates PatternFlowTcpCtlSynCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpCtlSynCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowTcpCtlSynCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowTcpCtlSynCounter
	SetStart(value uint32) PatternFlowTcpCtlSynCounter
	// HasStart checks if Start has been set in PatternFlowTcpCtlSynCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowTcpCtlSynCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowTcpCtlSynCounter
	SetStep(value uint32) PatternFlowTcpCtlSynCounter
	// HasStep checks if Step has been set in PatternFlowTcpCtlSynCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowTcpCtlSynCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowTcpCtlSynCounter
	SetCount(value uint32) PatternFlowTcpCtlSynCounter
	// HasCount checks if Count has been set in PatternFlowTcpCtlSynCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpCtlSynCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpCtlSynCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowTcpCtlSynCounter object
func (obj *patternFlowTcpCtlSynCounter) SetStart(value uint32) PatternFlowTcpCtlSynCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpCtlSynCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpCtlSynCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowTcpCtlSynCounter object
func (obj *patternFlowTcpCtlSynCounter) SetStep(value uint32) PatternFlowTcpCtlSynCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpCtlSynCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpCtlSynCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowTcpCtlSynCounter object
func (obj *patternFlowTcpCtlSynCounter) SetCount(value uint32) PatternFlowTcpCtlSynCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowTcpCtlSynCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlSynCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlSynCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlSynCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowTcpCtlSynCounter) setDefault() {
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

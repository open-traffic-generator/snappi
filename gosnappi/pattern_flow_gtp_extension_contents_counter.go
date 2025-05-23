package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpExtensionContentsCounter *****
type patternFlowGtpExtensionContentsCounter struct {
	validation
	obj          *otg.PatternFlowGtpExtensionContentsCounter
	marshaller   marshalPatternFlowGtpExtensionContentsCounter
	unMarshaller unMarshalPatternFlowGtpExtensionContentsCounter
}

func NewPatternFlowGtpExtensionContentsCounter() PatternFlowGtpExtensionContentsCounter {
	obj := patternFlowGtpExtensionContentsCounter{obj: &otg.PatternFlowGtpExtensionContentsCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpExtensionContentsCounter) msg() *otg.PatternFlowGtpExtensionContentsCounter {
	return obj.obj
}

func (obj *patternFlowGtpExtensionContentsCounter) setMsg(msg *otg.PatternFlowGtpExtensionContentsCounter) PatternFlowGtpExtensionContentsCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpExtensionContentsCounter struct {
	obj *patternFlowGtpExtensionContentsCounter
}

type marshalPatternFlowGtpExtensionContentsCounter interface {
	// ToProto marshals PatternFlowGtpExtensionContentsCounter to protobuf object *otg.PatternFlowGtpExtensionContentsCounter
	ToProto() (*otg.PatternFlowGtpExtensionContentsCounter, error)
	// ToPbText marshals PatternFlowGtpExtensionContentsCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpExtensionContentsCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpExtensionContentsCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpExtensionContentsCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpExtensionContentsCounter struct {
	obj *patternFlowGtpExtensionContentsCounter
}

type unMarshalPatternFlowGtpExtensionContentsCounter interface {
	// FromProto unmarshals PatternFlowGtpExtensionContentsCounter from protobuf object *otg.PatternFlowGtpExtensionContentsCounter
	FromProto(msg *otg.PatternFlowGtpExtensionContentsCounter) (PatternFlowGtpExtensionContentsCounter, error)
	// FromPbText unmarshals PatternFlowGtpExtensionContentsCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpExtensionContentsCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpExtensionContentsCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpExtensionContentsCounter) Marshal() marshalPatternFlowGtpExtensionContentsCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpExtensionContentsCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpExtensionContentsCounter) Unmarshal() unMarshalPatternFlowGtpExtensionContentsCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpExtensionContentsCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpExtensionContentsCounter) ToProto() (*otg.PatternFlowGtpExtensionContentsCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpExtensionContentsCounter) FromProto(msg *otg.PatternFlowGtpExtensionContentsCounter) (PatternFlowGtpExtensionContentsCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpExtensionContentsCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionContentsCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpExtensionContentsCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionContentsCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpExtensionContentsCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpExtensionContentsCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionContentsCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpExtensionContentsCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpExtensionContentsCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpExtensionContentsCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpExtensionContentsCounter) Clone() (PatternFlowGtpExtensionContentsCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpExtensionContentsCounter()
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

// PatternFlowGtpExtensionContentsCounter is integer counter pattern
type PatternFlowGtpExtensionContentsCounter interface {
	Validation
	// msg marshals PatternFlowGtpExtensionContentsCounter to protobuf object *otg.PatternFlowGtpExtensionContentsCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpExtensionContentsCounter
	// setMsg unmarshals PatternFlowGtpExtensionContentsCounter from protobuf object *otg.PatternFlowGtpExtensionContentsCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpExtensionContentsCounter) PatternFlowGtpExtensionContentsCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpExtensionContentsCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpExtensionContentsCounter
	// validate validates PatternFlowGtpExtensionContentsCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpExtensionContentsCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint64, set in PatternFlowGtpExtensionContentsCounter.
	Start() uint64
	// SetStart assigns uint64 provided by user to PatternFlowGtpExtensionContentsCounter
	SetStart(value uint64) PatternFlowGtpExtensionContentsCounter
	// HasStart checks if Start has been set in PatternFlowGtpExtensionContentsCounter
	HasStart() bool
	// Step returns uint64, set in PatternFlowGtpExtensionContentsCounter.
	Step() uint64
	// SetStep assigns uint64 provided by user to PatternFlowGtpExtensionContentsCounter
	SetStep(value uint64) PatternFlowGtpExtensionContentsCounter
	// HasStep checks if Step has been set in PatternFlowGtpExtensionContentsCounter
	HasStep() bool
	// Count returns uint64, set in PatternFlowGtpExtensionContentsCounter.
	Count() uint64
	// SetCount assigns uint64 provided by user to PatternFlowGtpExtensionContentsCounter
	SetCount(value uint64) PatternFlowGtpExtensionContentsCounter
	// HasCount checks if Count has been set in PatternFlowGtpExtensionContentsCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint64
func (obj *patternFlowGtpExtensionContentsCounter) Start() uint64 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint64
func (obj *patternFlowGtpExtensionContentsCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint64 value in the PatternFlowGtpExtensionContentsCounter object
func (obj *patternFlowGtpExtensionContentsCounter) SetStart(value uint64) PatternFlowGtpExtensionContentsCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint64
func (obj *patternFlowGtpExtensionContentsCounter) Step() uint64 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint64
func (obj *patternFlowGtpExtensionContentsCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint64 value in the PatternFlowGtpExtensionContentsCounter object
func (obj *patternFlowGtpExtensionContentsCounter) SetStep(value uint64) PatternFlowGtpExtensionContentsCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint64
func (obj *patternFlowGtpExtensionContentsCounter) Count() uint64 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint64
func (obj *patternFlowGtpExtensionContentsCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint64 value in the PatternFlowGtpExtensionContentsCounter object
func (obj *patternFlowGtpExtensionContentsCounter) SetCount(value uint64) PatternFlowGtpExtensionContentsCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpExtensionContentsCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 281474976710655 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionContentsCounter.Start <= 281474976710655 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 281474976710655 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionContentsCounter.Step <= 281474976710655 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 281474976710655 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionContentsCounter.Count <= 281474976710655 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGtpExtensionContentsCounter) setDefault() {
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

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowSizeIncrement *****
type flowSizeIncrement struct {
	validation
	obj          *otg.FlowSizeIncrement
	marshaller   marshalFlowSizeIncrement
	unMarshaller unMarshalFlowSizeIncrement
}

func NewFlowSizeIncrement() FlowSizeIncrement {
	obj := flowSizeIncrement{obj: &otg.FlowSizeIncrement{}}
	obj.setDefault()
	return &obj
}

func (obj *flowSizeIncrement) msg() *otg.FlowSizeIncrement {
	return obj.obj
}

func (obj *flowSizeIncrement) setMsg(msg *otg.FlowSizeIncrement) FlowSizeIncrement {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowSizeIncrement struct {
	obj *flowSizeIncrement
}

type marshalFlowSizeIncrement interface {
	// ToProto marshals FlowSizeIncrement to protobuf object *otg.FlowSizeIncrement
	ToProto() (*otg.FlowSizeIncrement, error)
	// ToPbText marshals FlowSizeIncrement to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowSizeIncrement to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowSizeIncrement to JSON text
	ToJson() (string, error)
}

type unMarshalflowSizeIncrement struct {
	obj *flowSizeIncrement
}

type unMarshalFlowSizeIncrement interface {
	// FromProto unmarshals FlowSizeIncrement from protobuf object *otg.FlowSizeIncrement
	FromProto(msg *otg.FlowSizeIncrement) (FlowSizeIncrement, error)
	// FromPbText unmarshals FlowSizeIncrement from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowSizeIncrement from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowSizeIncrement from JSON text
	FromJson(value string) error
}

func (obj *flowSizeIncrement) Marshal() marshalFlowSizeIncrement {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowSizeIncrement{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowSizeIncrement) Unmarshal() unMarshalFlowSizeIncrement {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowSizeIncrement{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowSizeIncrement) ToProto() (*otg.FlowSizeIncrement, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowSizeIncrement) FromProto(msg *otg.FlowSizeIncrement) (FlowSizeIncrement, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowSizeIncrement) ToPbText() (string, error) {
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

func (m *unMarshalflowSizeIncrement) FromPbText(value string) error {
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

func (m *marshalflowSizeIncrement) ToYaml() (string, error) {
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

func (m *unMarshalflowSizeIncrement) FromYaml(value string) error {
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

func (m *marshalflowSizeIncrement) ToJson() (string, error) {
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

func (m *unMarshalflowSizeIncrement) FromJson(value string) error {
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

func (obj *flowSizeIncrement) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowSizeIncrement) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowSizeIncrement) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowSizeIncrement) Clone() (FlowSizeIncrement, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowSizeIncrement()
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

// FlowSizeIncrement is frame size that increments from a starting size to
// an ending size incrementing by a step size.
type FlowSizeIncrement interface {
	Validation
	// msg marshals FlowSizeIncrement to protobuf object *otg.FlowSizeIncrement
	// and doesn't set defaults
	msg() *otg.FlowSizeIncrement
	// setMsg unmarshals FlowSizeIncrement from protobuf object *otg.FlowSizeIncrement
	// and doesn't set defaults
	setMsg(*otg.FlowSizeIncrement) FlowSizeIncrement
	// provides marshal interface
	Marshal() marshalFlowSizeIncrement
	// provides unmarshal interface
	Unmarshal() unMarshalFlowSizeIncrement
	// validate validates FlowSizeIncrement
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowSizeIncrement, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in FlowSizeIncrement.
	Start() uint32
	// SetStart assigns uint32 provided by user to FlowSizeIncrement
	SetStart(value uint32) FlowSizeIncrement
	// HasStart checks if Start has been set in FlowSizeIncrement
	HasStart() bool
	// End returns uint32, set in FlowSizeIncrement.
	End() uint32
	// SetEnd assigns uint32 provided by user to FlowSizeIncrement
	SetEnd(value uint32) FlowSizeIncrement
	// HasEnd checks if End has been set in FlowSizeIncrement
	HasEnd() bool
	// Step returns uint32, set in FlowSizeIncrement.
	Step() uint32
	// SetStep assigns uint32 provided by user to FlowSizeIncrement
	SetStep(value uint32) FlowSizeIncrement
	// HasStep checks if Step has been set in FlowSizeIncrement
	HasStep() bool
}

// Starting frame size in bytes
// Start returns a uint32
func (obj *flowSizeIncrement) Start() uint32 {

	return *obj.obj.Start

}

// Starting frame size in bytes
// Start returns a uint32
func (obj *flowSizeIncrement) HasStart() bool {
	return obj.obj.Start != nil
}

// Starting frame size in bytes
// SetStart sets the uint32 value in the FlowSizeIncrement object
func (obj *flowSizeIncrement) SetStart(value uint32) FlowSizeIncrement {

	obj.obj.Start = &value
	return obj
}

// Ending frame size in bytes
// End returns a uint32
func (obj *flowSizeIncrement) End() uint32 {

	return *obj.obj.End

}

// Ending frame size in bytes
// End returns a uint32
func (obj *flowSizeIncrement) HasEnd() bool {
	return obj.obj.End != nil
}

// Ending frame size in bytes
// SetEnd sets the uint32 value in the FlowSizeIncrement object
func (obj *flowSizeIncrement) SetEnd(value uint32) FlowSizeIncrement {

	obj.obj.End = &value
	return obj
}

// Step frame size in bytes
// Step returns a uint32
func (obj *flowSizeIncrement) Step() uint32 {

	return *obj.obj.Step

}

// Step frame size in bytes
// Step returns a uint32
func (obj *flowSizeIncrement) HasStep() bool {
	return obj.obj.Step != nil
}

// Step frame size in bytes
// SetStep sets the uint32 value in the FlowSizeIncrement object
func (obj *flowSizeIncrement) SetStep(value uint32) FlowSizeIncrement {

	obj.obj.Step = &value
	return obj
}

func (obj *flowSizeIncrement) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start < 1 || *obj.obj.Start > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= FlowSizeIncrement.Start <= 4294967295 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.End != nil {

		if *obj.obj.End < 64 || *obj.obj.End > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("64 <= FlowSizeIncrement.End <= 4294967295 but Got %d", *obj.obj.End))
		}

	}

}

func (obj *flowSizeIncrement) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(64)
	}
	if obj.obj.End == nil {
		obj.SetEnd(1518)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}

}

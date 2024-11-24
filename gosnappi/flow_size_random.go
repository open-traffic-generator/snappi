package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowSizeRandom *****
type flowSizeRandom struct {
	validation
	obj          *otg.FlowSizeRandom
	marshaller   marshalFlowSizeRandom
	unMarshaller unMarshalFlowSizeRandom
}

func NewFlowSizeRandom() FlowSizeRandom {
	obj := flowSizeRandom{obj: &otg.FlowSizeRandom{}}
	obj.setDefault()
	return &obj
}

func (obj *flowSizeRandom) msg() *otg.FlowSizeRandom {
	return obj.obj
}

func (obj *flowSizeRandom) setMsg(msg *otg.FlowSizeRandom) FlowSizeRandom {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowSizeRandom struct {
	obj *flowSizeRandom
}

type marshalFlowSizeRandom interface {
	// ToProto marshals FlowSizeRandom to protobuf object *otg.FlowSizeRandom
	ToProto() (*otg.FlowSizeRandom, error)
	// ToPbText marshals FlowSizeRandom to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowSizeRandom to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowSizeRandom to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowSizeRandom to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowSizeRandom struct {
	obj *flowSizeRandom
}

type unMarshalFlowSizeRandom interface {
	// FromProto unmarshals FlowSizeRandom from protobuf object *otg.FlowSizeRandom
	FromProto(msg *otg.FlowSizeRandom) (FlowSizeRandom, error)
	// FromPbText unmarshals FlowSizeRandom from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowSizeRandom from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowSizeRandom from JSON text
	FromJson(value string) error
}

func (obj *flowSizeRandom) Marshal() marshalFlowSizeRandom {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowSizeRandom{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowSizeRandom) Unmarshal() unMarshalFlowSizeRandom {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowSizeRandom{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowSizeRandom) ToProto() (*otg.FlowSizeRandom, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowSizeRandom) FromProto(msg *otg.FlowSizeRandom) (FlowSizeRandom, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowSizeRandom) ToPbText() (string, error) {
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

func (m *unMarshalflowSizeRandom) FromPbText(value string) error {
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

func (m *marshalflowSizeRandom) ToYaml() (string, error) {
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

func (m *unMarshalflowSizeRandom) FromYaml(value string) error {
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

func (m *marshalflowSizeRandom) ToJsonRaw() (string, error) {
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

func (m *marshalflowSizeRandom) ToJson() (string, error) {
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

func (m *unMarshalflowSizeRandom) FromJson(value string) error {
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

func (obj *flowSizeRandom) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowSizeRandom) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowSizeRandom) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowSizeRandom) Clone() (FlowSizeRandom, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowSizeRandom()
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

// FlowSizeRandom is random frame size from a min value to a max value.
type FlowSizeRandom interface {
	Validation
	// msg marshals FlowSizeRandom to protobuf object *otg.FlowSizeRandom
	// and doesn't set defaults
	msg() *otg.FlowSizeRandom
	// setMsg unmarshals FlowSizeRandom from protobuf object *otg.FlowSizeRandom
	// and doesn't set defaults
	setMsg(*otg.FlowSizeRandom) FlowSizeRandom
	// provides marshal interface
	Marshal() marshalFlowSizeRandom
	// provides unmarshal interface
	Unmarshal() unMarshalFlowSizeRandom
	// validate validates FlowSizeRandom
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowSizeRandom, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Min returns uint32, set in FlowSizeRandom.
	Min() uint32
	// SetMin assigns uint32 provided by user to FlowSizeRandom
	SetMin(value uint32) FlowSizeRandom
	// HasMin checks if Min has been set in FlowSizeRandom
	HasMin() bool
	// Max returns uint32, set in FlowSizeRandom.
	Max() uint32
	// SetMax assigns uint32 provided by user to FlowSizeRandom
	SetMax(value uint32) FlowSizeRandom
	// HasMax checks if Max has been set in FlowSizeRandom
	HasMax() bool
}

// description is TBD
// Min returns a uint32
func (obj *flowSizeRandom) Min() uint32 {

	return *obj.obj.Min

}

// description is TBD
// Min returns a uint32
func (obj *flowSizeRandom) HasMin() bool {
	return obj.obj.Min != nil
}

// description is TBD
// SetMin sets the uint32 value in the FlowSizeRandom object
func (obj *flowSizeRandom) SetMin(value uint32) FlowSizeRandom {

	obj.obj.Min = &value
	return obj
}

// description is TBD
// Max returns a uint32
func (obj *flowSizeRandom) Max() uint32 {

	return *obj.obj.Max

}

// description is TBD
// Max returns a uint32
func (obj *flowSizeRandom) HasMax() bool {
	return obj.obj.Max != nil
}

// description is TBD
// SetMax sets the uint32 value in the FlowSizeRandom object
func (obj *flowSizeRandom) SetMax(value uint32) FlowSizeRandom {

	obj.obj.Max = &value
	return obj
}

func (obj *flowSizeRandom) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *flowSizeRandom) setDefault() {
	if obj.obj.Min == nil {
		obj.SetMin(64)
	}
	if obj.obj.Max == nil {
		obj.SetMax(1518)
	}

}

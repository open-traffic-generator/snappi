package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowSizeWeightPairsCustom *****
type flowSizeWeightPairsCustom struct {
	validation
	obj          *otg.FlowSizeWeightPairsCustom
	marshaller   marshalFlowSizeWeightPairsCustom
	unMarshaller unMarshalFlowSizeWeightPairsCustom
}

func NewFlowSizeWeightPairsCustom() FlowSizeWeightPairsCustom {
	obj := flowSizeWeightPairsCustom{obj: &otg.FlowSizeWeightPairsCustom{}}
	obj.setDefault()
	return &obj
}

func (obj *flowSizeWeightPairsCustom) msg() *otg.FlowSizeWeightPairsCustom {
	return obj.obj
}

func (obj *flowSizeWeightPairsCustom) setMsg(msg *otg.FlowSizeWeightPairsCustom) FlowSizeWeightPairsCustom {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowSizeWeightPairsCustom struct {
	obj *flowSizeWeightPairsCustom
}

type marshalFlowSizeWeightPairsCustom interface {
	// ToProto marshals FlowSizeWeightPairsCustom to protobuf object *otg.FlowSizeWeightPairsCustom
	ToProto() (*otg.FlowSizeWeightPairsCustom, error)
	// ToPbText marshals FlowSizeWeightPairsCustom to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowSizeWeightPairsCustom to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowSizeWeightPairsCustom to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowSizeWeightPairsCustom to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowSizeWeightPairsCustom struct {
	obj *flowSizeWeightPairsCustom
}

type unMarshalFlowSizeWeightPairsCustom interface {
	// FromProto unmarshals FlowSizeWeightPairsCustom from protobuf object *otg.FlowSizeWeightPairsCustom
	FromProto(msg *otg.FlowSizeWeightPairsCustom) (FlowSizeWeightPairsCustom, error)
	// FromPbText unmarshals FlowSizeWeightPairsCustom from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowSizeWeightPairsCustom from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowSizeWeightPairsCustom from JSON text
	FromJson(value string) error
}

func (obj *flowSizeWeightPairsCustom) Marshal() marshalFlowSizeWeightPairsCustom {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowSizeWeightPairsCustom{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowSizeWeightPairsCustom) Unmarshal() unMarshalFlowSizeWeightPairsCustom {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowSizeWeightPairsCustom{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowSizeWeightPairsCustom) ToProto() (*otg.FlowSizeWeightPairsCustom, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowSizeWeightPairsCustom) FromProto(msg *otg.FlowSizeWeightPairsCustom) (FlowSizeWeightPairsCustom, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowSizeWeightPairsCustom) ToPbText() (string, error) {
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

func (m *unMarshalflowSizeWeightPairsCustom) FromPbText(value string) error {
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

func (m *marshalflowSizeWeightPairsCustom) ToYaml() (string, error) {
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

func (m *unMarshalflowSizeWeightPairsCustom) FromYaml(value string) error {
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

func (m *marshalflowSizeWeightPairsCustom) ToJsonRaw() (string, error) {
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

func (m *marshalflowSizeWeightPairsCustom) ToJson() (string, error) {
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

func (m *unMarshalflowSizeWeightPairsCustom) FromJson(value string) error {
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

func (obj *flowSizeWeightPairsCustom) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowSizeWeightPairsCustom) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowSizeWeightPairsCustom) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowSizeWeightPairsCustom) Clone() (FlowSizeWeightPairsCustom, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowSizeWeightPairsCustom()
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

// FlowSizeWeightPairsCustom is custom frame size distribution <size, weight> pair.
type FlowSizeWeightPairsCustom interface {
	Validation
	// msg marshals FlowSizeWeightPairsCustom to protobuf object *otg.FlowSizeWeightPairsCustom
	// and doesn't set defaults
	msg() *otg.FlowSizeWeightPairsCustom
	// setMsg unmarshals FlowSizeWeightPairsCustom from protobuf object *otg.FlowSizeWeightPairsCustom
	// and doesn't set defaults
	setMsg(*otg.FlowSizeWeightPairsCustom) FlowSizeWeightPairsCustom
	// provides marshal interface
	Marshal() marshalFlowSizeWeightPairsCustom
	// provides unmarshal interface
	Unmarshal() unMarshalFlowSizeWeightPairsCustom
	// validate validates FlowSizeWeightPairsCustom
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowSizeWeightPairsCustom, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Size returns uint32, set in FlowSizeWeightPairsCustom.
	Size() uint32
	// SetSize assigns uint32 provided by user to FlowSizeWeightPairsCustom
	SetSize(value uint32) FlowSizeWeightPairsCustom
	// HasSize checks if Size has been set in FlowSizeWeightPairsCustom
	HasSize() bool
	// Weight returns float32, set in FlowSizeWeightPairsCustom.
	Weight() float32
	// SetWeight assigns float32 provided by user to FlowSizeWeightPairsCustom
	SetWeight(value float32) FlowSizeWeightPairsCustom
	// HasWeight checks if Weight has been set in FlowSizeWeightPairsCustom
	HasWeight() bool
}

// The size of the frame (in bytes) for this weight pair.
// Size returns a uint32
func (obj *flowSizeWeightPairsCustom) Size() uint32 {

	return *obj.obj.Size

}

// The size of the frame (in bytes) for this weight pair.
// Size returns a uint32
func (obj *flowSizeWeightPairsCustom) HasSize() bool {
	return obj.obj.Size != nil
}

// The size of the frame (in bytes) for this weight pair.
// SetSize sets the uint32 value in the FlowSizeWeightPairsCustom object
func (obj *flowSizeWeightPairsCustom) SetSize(value uint32) FlowSizeWeightPairsCustom {

	obj.obj.Size = &value
	return obj
}

// Weight assigned to the corresponding frame size in this weight pair.
// Higher weight means more packets.
// Weight returns a float32
func (obj *flowSizeWeightPairsCustom) Weight() float32 {

	return *obj.obj.Weight

}

// Weight assigned to the corresponding frame size in this weight pair.
// Higher weight means more packets.
// Weight returns a float32
func (obj *flowSizeWeightPairsCustom) HasWeight() bool {
	return obj.obj.Weight != nil
}

// Weight assigned to the corresponding frame size in this weight pair.
// Higher weight means more packets.
// SetWeight sets the float32 value in the FlowSizeWeightPairsCustom object
func (obj *flowSizeWeightPairsCustom) SetWeight(value float32) FlowSizeWeightPairsCustom {

	obj.obj.Weight = &value
	return obj
}

func (obj *flowSizeWeightPairsCustom) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Size != nil {

		if *obj.obj.Size < 12 || *obj.obj.Size > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("12 <= FlowSizeWeightPairsCustom.Size <= 65535 but Got %d", *obj.obj.Size))
		}

	}

}

func (obj *flowSizeWeightPairsCustom) setDefault() {
	if obj.obj.Size == nil {
		obj.SetSize(64)
	}
	if obj.obj.Weight == nil {
		obj.SetWeight(1)
	}

}

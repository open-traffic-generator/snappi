package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSegmentRoutingPolicySegmentListWeight *****
type bgpAttributesSegmentRoutingPolicySegmentListWeight struct {
	validation
	obj          *otg.BgpAttributesSegmentRoutingPolicySegmentListWeight
	marshaller   marshalBgpAttributesSegmentRoutingPolicySegmentListWeight
	unMarshaller unMarshalBgpAttributesSegmentRoutingPolicySegmentListWeight
}

func NewBgpAttributesSegmentRoutingPolicySegmentListWeight() BgpAttributesSegmentRoutingPolicySegmentListWeight {
	obj := bgpAttributesSegmentRoutingPolicySegmentListWeight{obj: &otg.BgpAttributesSegmentRoutingPolicySegmentListWeight{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSegmentRoutingPolicySegmentListWeight) msg() *otg.BgpAttributesSegmentRoutingPolicySegmentListWeight {
	return obj.obj
}

func (obj *bgpAttributesSegmentRoutingPolicySegmentListWeight) setMsg(msg *otg.BgpAttributesSegmentRoutingPolicySegmentListWeight) BgpAttributesSegmentRoutingPolicySegmentListWeight {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSegmentRoutingPolicySegmentListWeight struct {
	obj *bgpAttributesSegmentRoutingPolicySegmentListWeight
}

type marshalBgpAttributesSegmentRoutingPolicySegmentListWeight interface {
	// ToProto marshals BgpAttributesSegmentRoutingPolicySegmentListWeight to protobuf object *otg.BgpAttributesSegmentRoutingPolicySegmentListWeight
	ToProto() (*otg.BgpAttributesSegmentRoutingPolicySegmentListWeight, error)
	// ToPbText marshals BgpAttributesSegmentRoutingPolicySegmentListWeight to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSegmentRoutingPolicySegmentListWeight to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSegmentRoutingPolicySegmentListWeight to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributesSegmentRoutingPolicySegmentListWeight struct {
	obj *bgpAttributesSegmentRoutingPolicySegmentListWeight
}

type unMarshalBgpAttributesSegmentRoutingPolicySegmentListWeight interface {
	// FromProto unmarshals BgpAttributesSegmentRoutingPolicySegmentListWeight from protobuf object *otg.BgpAttributesSegmentRoutingPolicySegmentListWeight
	FromProto(msg *otg.BgpAttributesSegmentRoutingPolicySegmentListWeight) (BgpAttributesSegmentRoutingPolicySegmentListWeight, error)
	// FromPbText unmarshals BgpAttributesSegmentRoutingPolicySegmentListWeight from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSegmentRoutingPolicySegmentListWeight from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSegmentRoutingPolicySegmentListWeight from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSegmentRoutingPolicySegmentListWeight) Marshal() marshalBgpAttributesSegmentRoutingPolicySegmentListWeight {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSegmentRoutingPolicySegmentListWeight{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSegmentRoutingPolicySegmentListWeight) Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicySegmentListWeight {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSegmentRoutingPolicySegmentListWeight{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSegmentRoutingPolicySegmentListWeight) ToProto() (*otg.BgpAttributesSegmentRoutingPolicySegmentListWeight, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSegmentRoutingPolicySegmentListWeight) FromProto(msg *otg.BgpAttributesSegmentRoutingPolicySegmentListWeight) (BgpAttributesSegmentRoutingPolicySegmentListWeight, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSegmentRoutingPolicySegmentListWeight) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicySegmentListWeight) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicySegmentListWeight) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicySegmentListWeight) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicySegmentListWeight) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicySegmentListWeight) FromJson(value string) error {
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

func (obj *bgpAttributesSegmentRoutingPolicySegmentListWeight) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicySegmentListWeight) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicySegmentListWeight) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSegmentRoutingPolicySegmentListWeight) Clone() (BgpAttributesSegmentRoutingPolicySegmentListWeight, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSegmentRoutingPolicySegmentListWeight()
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

// BgpAttributesSegmentRoutingPolicySegmentListWeight is the optional Weight sub-TLV (Type 9) specifies the weight associated with a given segment list. The weight is used for weighted multipath.
type BgpAttributesSegmentRoutingPolicySegmentListWeight interface {
	Validation
	// msg marshals BgpAttributesSegmentRoutingPolicySegmentListWeight to protobuf object *otg.BgpAttributesSegmentRoutingPolicySegmentListWeight
	// and doesn't set defaults
	msg() *otg.BgpAttributesSegmentRoutingPolicySegmentListWeight
	// setMsg unmarshals BgpAttributesSegmentRoutingPolicySegmentListWeight from protobuf object *otg.BgpAttributesSegmentRoutingPolicySegmentListWeight
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSegmentRoutingPolicySegmentListWeight) BgpAttributesSegmentRoutingPolicySegmentListWeight
	// provides marshal interface
	Marshal() marshalBgpAttributesSegmentRoutingPolicySegmentListWeight
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicySegmentListWeight
	// validate validates BgpAttributesSegmentRoutingPolicySegmentListWeight
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSegmentRoutingPolicySegmentListWeight, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Value returns uint32, set in BgpAttributesSegmentRoutingPolicySegmentListWeight.
	Value() uint32
	// SetValue assigns uint32 provided by user to BgpAttributesSegmentRoutingPolicySegmentListWeight
	SetValue(value uint32) BgpAttributesSegmentRoutingPolicySegmentListWeight
	// HasValue checks if Value has been set in BgpAttributesSegmentRoutingPolicySegmentListWeight
	HasValue() bool
}

// Value of the weight.
// Value returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicySegmentListWeight) Value() uint32 {

	return *obj.obj.Value

}

// Value of the weight.
// Value returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicySegmentListWeight) HasValue() bool {
	return obj.obj.Value != nil
}

// Value of the weight.
// SetValue sets the uint32 value in the BgpAttributesSegmentRoutingPolicySegmentListWeight object
func (obj *bgpAttributesSegmentRoutingPolicySegmentListWeight) SetValue(value uint32) BgpAttributesSegmentRoutingPolicySegmentListWeight {

	obj.obj.Value = &value
	return obj
}

func (obj *bgpAttributesSegmentRoutingPolicySegmentListWeight) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpAttributesSegmentRoutingPolicySegmentListWeight) setDefault() {
	if obj.obj.Value == nil {
		obj.SetValue(0)
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesMultiExitDiscriminator *****
type bgpAttributesMultiExitDiscriminator struct {
	validation
	obj          *otg.BgpAttributesMultiExitDiscriminator
	marshaller   marshalBgpAttributesMultiExitDiscriminator
	unMarshaller unMarshalBgpAttributesMultiExitDiscriminator
}

func NewBgpAttributesMultiExitDiscriminator() BgpAttributesMultiExitDiscriminator {
	obj := bgpAttributesMultiExitDiscriminator{obj: &otg.BgpAttributesMultiExitDiscriminator{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesMultiExitDiscriminator) msg() *otg.BgpAttributesMultiExitDiscriminator {
	return obj.obj
}

func (obj *bgpAttributesMultiExitDiscriminator) setMsg(msg *otg.BgpAttributesMultiExitDiscriminator) BgpAttributesMultiExitDiscriminator {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesMultiExitDiscriminator struct {
	obj *bgpAttributesMultiExitDiscriminator
}

type marshalBgpAttributesMultiExitDiscriminator interface {
	// ToProto marshals BgpAttributesMultiExitDiscriminator to protobuf object *otg.BgpAttributesMultiExitDiscriminator
	ToProto() (*otg.BgpAttributesMultiExitDiscriminator, error)
	// ToPbText marshals BgpAttributesMultiExitDiscriminator to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesMultiExitDiscriminator to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesMultiExitDiscriminator to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesMultiExitDiscriminator to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesMultiExitDiscriminator struct {
	obj *bgpAttributesMultiExitDiscriminator
}

type unMarshalBgpAttributesMultiExitDiscriminator interface {
	// FromProto unmarshals BgpAttributesMultiExitDiscriminator from protobuf object *otg.BgpAttributesMultiExitDiscriminator
	FromProto(msg *otg.BgpAttributesMultiExitDiscriminator) (BgpAttributesMultiExitDiscriminator, error)
	// FromPbText unmarshals BgpAttributesMultiExitDiscriminator from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesMultiExitDiscriminator from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesMultiExitDiscriminator from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesMultiExitDiscriminator) Marshal() marshalBgpAttributesMultiExitDiscriminator {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesMultiExitDiscriminator{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesMultiExitDiscriminator) Unmarshal() unMarshalBgpAttributesMultiExitDiscriminator {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesMultiExitDiscriminator{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesMultiExitDiscriminator) ToProto() (*otg.BgpAttributesMultiExitDiscriminator, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesMultiExitDiscriminator) FromProto(msg *otg.BgpAttributesMultiExitDiscriminator) (BgpAttributesMultiExitDiscriminator, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesMultiExitDiscriminator) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesMultiExitDiscriminator) FromPbText(value string) error {
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

func (m *marshalbgpAttributesMultiExitDiscriminator) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesMultiExitDiscriminator) FromYaml(value string) error {
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

func (m *marshalbgpAttributesMultiExitDiscriminator) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesMultiExitDiscriminator) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesMultiExitDiscriminator) FromJson(value string) error {
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

func (obj *bgpAttributesMultiExitDiscriminator) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesMultiExitDiscriminator) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesMultiExitDiscriminator) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesMultiExitDiscriminator) Clone() (BgpAttributesMultiExitDiscriminator, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesMultiExitDiscriminator()
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

// BgpAttributesMultiExitDiscriminator is optional MULTI_EXIT_DISCRIMINATOR attribute sent to the peer to help in the route selection process.
type BgpAttributesMultiExitDiscriminator interface {
	Validation
	// msg marshals BgpAttributesMultiExitDiscriminator to protobuf object *otg.BgpAttributesMultiExitDiscriminator
	// and doesn't set defaults
	msg() *otg.BgpAttributesMultiExitDiscriminator
	// setMsg unmarshals BgpAttributesMultiExitDiscriminator from protobuf object *otg.BgpAttributesMultiExitDiscriminator
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesMultiExitDiscriminator) BgpAttributesMultiExitDiscriminator
	// provides marshal interface
	Marshal() marshalBgpAttributesMultiExitDiscriminator
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesMultiExitDiscriminator
	// validate validates BgpAttributesMultiExitDiscriminator
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesMultiExitDiscriminator, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Value returns uint32, set in BgpAttributesMultiExitDiscriminator.
	Value() uint32
	// SetValue assigns uint32 provided by user to BgpAttributesMultiExitDiscriminator
	SetValue(value uint32) BgpAttributesMultiExitDiscriminator
	// HasValue checks if Value has been set in BgpAttributesMultiExitDiscriminator
	HasValue() bool
}

// The multi exit discriminator (MED) value used for route selection sent to the peer.
// Value returns a uint32
func (obj *bgpAttributesMultiExitDiscriminator) Value() uint32 {

	return *obj.obj.Value

}

// The multi exit discriminator (MED) value used for route selection sent to the peer.
// Value returns a uint32
func (obj *bgpAttributesMultiExitDiscriminator) HasValue() bool {
	return obj.obj.Value != nil
}

// The multi exit discriminator (MED) value used for route selection sent to the peer.
// SetValue sets the uint32 value in the BgpAttributesMultiExitDiscriminator object
func (obj *bgpAttributesMultiExitDiscriminator) SetValue(value uint32) BgpAttributesMultiExitDiscriminator {

	obj.obj.Value = &value
	return obj
}

func (obj *bgpAttributesMultiExitDiscriminator) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpAttributesMultiExitDiscriminator) setDefault() {
	if obj.obj.Value == nil {
		obj.SetValue(0)
	}

}

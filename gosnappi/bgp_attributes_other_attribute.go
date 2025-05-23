package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesOtherAttribute *****
type bgpAttributesOtherAttribute struct {
	validation
	obj          *otg.BgpAttributesOtherAttribute
	marshaller   marshalBgpAttributesOtherAttribute
	unMarshaller unMarshalBgpAttributesOtherAttribute
}

func NewBgpAttributesOtherAttribute() BgpAttributesOtherAttribute {
	obj := bgpAttributesOtherAttribute{obj: &otg.BgpAttributesOtherAttribute{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesOtherAttribute) msg() *otg.BgpAttributesOtherAttribute {
	return obj.obj
}

func (obj *bgpAttributesOtherAttribute) setMsg(msg *otg.BgpAttributesOtherAttribute) BgpAttributesOtherAttribute {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesOtherAttribute struct {
	obj *bgpAttributesOtherAttribute
}

type marshalBgpAttributesOtherAttribute interface {
	// ToProto marshals BgpAttributesOtherAttribute to protobuf object *otg.BgpAttributesOtherAttribute
	ToProto() (*otg.BgpAttributesOtherAttribute, error)
	// ToPbText marshals BgpAttributesOtherAttribute to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesOtherAttribute to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesOtherAttribute to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesOtherAttribute to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesOtherAttribute struct {
	obj *bgpAttributesOtherAttribute
}

type unMarshalBgpAttributesOtherAttribute interface {
	// FromProto unmarshals BgpAttributesOtherAttribute from protobuf object *otg.BgpAttributesOtherAttribute
	FromProto(msg *otg.BgpAttributesOtherAttribute) (BgpAttributesOtherAttribute, error)
	// FromPbText unmarshals BgpAttributesOtherAttribute from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesOtherAttribute from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesOtherAttribute from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesOtherAttribute) Marshal() marshalBgpAttributesOtherAttribute {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesOtherAttribute{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesOtherAttribute) Unmarshal() unMarshalBgpAttributesOtherAttribute {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesOtherAttribute{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesOtherAttribute) ToProto() (*otg.BgpAttributesOtherAttribute, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesOtherAttribute) FromProto(msg *otg.BgpAttributesOtherAttribute) (BgpAttributesOtherAttribute, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesOtherAttribute) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesOtherAttribute) FromPbText(value string) error {
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

func (m *marshalbgpAttributesOtherAttribute) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesOtherAttribute) FromYaml(value string) error {
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

func (m *marshalbgpAttributesOtherAttribute) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesOtherAttribute) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesOtherAttribute) FromJson(value string) error {
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

func (obj *bgpAttributesOtherAttribute) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesOtherAttribute) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesOtherAttribute) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesOtherAttribute) Clone() (BgpAttributesOtherAttribute, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesOtherAttribute()
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

// BgpAttributesOtherAttribute is one unknown attribute stored as hex bytes.
type BgpAttributesOtherAttribute interface {
	Validation
	// msg marshals BgpAttributesOtherAttribute to protobuf object *otg.BgpAttributesOtherAttribute
	// and doesn't set defaults
	msg() *otg.BgpAttributesOtherAttribute
	// setMsg unmarshals BgpAttributesOtherAttribute from protobuf object *otg.BgpAttributesOtherAttribute
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesOtherAttribute) BgpAttributesOtherAttribute
	// provides marshal interface
	Marshal() marshalBgpAttributesOtherAttribute
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesOtherAttribute
	// validate validates BgpAttributesOtherAttribute
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesOtherAttribute, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// FlagOptional returns bool, set in BgpAttributesOtherAttribute.
	FlagOptional() bool
	// SetFlagOptional assigns bool provided by user to BgpAttributesOtherAttribute
	SetFlagOptional(value bool) BgpAttributesOtherAttribute
	// HasFlagOptional checks if FlagOptional has been set in BgpAttributesOtherAttribute
	HasFlagOptional() bool
	// FlagTransitive returns bool, set in BgpAttributesOtherAttribute.
	FlagTransitive() bool
	// SetFlagTransitive assigns bool provided by user to BgpAttributesOtherAttribute
	SetFlagTransitive(value bool) BgpAttributesOtherAttribute
	// HasFlagTransitive checks if FlagTransitive has been set in BgpAttributesOtherAttribute
	HasFlagTransitive() bool
	// FlagPartial returns bool, set in BgpAttributesOtherAttribute.
	FlagPartial() bool
	// SetFlagPartial assigns bool provided by user to BgpAttributesOtherAttribute
	SetFlagPartial(value bool) BgpAttributesOtherAttribute
	// HasFlagPartial checks if FlagPartial has been set in BgpAttributesOtherAttribute
	HasFlagPartial() bool
	// FlagExtendedLength returns bool, set in BgpAttributesOtherAttribute.
	FlagExtendedLength() bool
	// SetFlagExtendedLength assigns bool provided by user to BgpAttributesOtherAttribute
	SetFlagExtendedLength(value bool) BgpAttributesOtherAttribute
	// HasFlagExtendedLength checks if FlagExtendedLength has been set in BgpAttributesOtherAttribute
	HasFlagExtendedLength() bool
	// Type returns uint32, set in BgpAttributesOtherAttribute.
	Type() uint32
	// SetType assigns uint32 provided by user to BgpAttributesOtherAttribute
	SetType(value uint32) BgpAttributesOtherAttribute
	// RawValue returns string, set in BgpAttributesOtherAttribute.
	RawValue() string
	// SetRawValue assigns string provided by user to BgpAttributesOtherAttribute
	SetRawValue(value string) BgpAttributesOtherAttribute
}

// Optional flag in the BGP attribute.
// FlagOptional returns a bool
func (obj *bgpAttributesOtherAttribute) FlagOptional() bool {

	return *obj.obj.FlagOptional

}

// Optional flag in the BGP attribute.
// FlagOptional returns a bool
func (obj *bgpAttributesOtherAttribute) HasFlagOptional() bool {
	return obj.obj.FlagOptional != nil
}

// Optional flag in the BGP attribute.
// SetFlagOptional sets the bool value in the BgpAttributesOtherAttribute object
func (obj *bgpAttributesOtherAttribute) SetFlagOptional(value bool) BgpAttributesOtherAttribute {

	obj.obj.FlagOptional = &value
	return obj
}

// Transitive flag in the BGP attribute.
// FlagTransitive returns a bool
func (obj *bgpAttributesOtherAttribute) FlagTransitive() bool {

	return *obj.obj.FlagTransitive

}

// Transitive flag in the BGP attribute.
// FlagTransitive returns a bool
func (obj *bgpAttributesOtherAttribute) HasFlagTransitive() bool {
	return obj.obj.FlagTransitive != nil
}

// Transitive flag in the BGP attribute.
// SetFlagTransitive sets the bool value in the BgpAttributesOtherAttribute object
func (obj *bgpAttributesOtherAttribute) SetFlagTransitive(value bool) BgpAttributesOtherAttribute {

	obj.obj.FlagTransitive = &value
	return obj
}

// Partial flag in the BGP attribute.
// FlagPartial returns a bool
func (obj *bgpAttributesOtherAttribute) FlagPartial() bool {

	return *obj.obj.FlagPartial

}

// Partial flag in the BGP attribute.
// FlagPartial returns a bool
func (obj *bgpAttributesOtherAttribute) HasFlagPartial() bool {
	return obj.obj.FlagPartial != nil
}

// Partial flag in the BGP attribute.
// SetFlagPartial sets the bool value in the BgpAttributesOtherAttribute object
func (obj *bgpAttributesOtherAttribute) SetFlagPartial(value bool) BgpAttributesOtherAttribute {

	obj.obj.FlagPartial = &value
	return obj
}

// Extended length flag in the BGP attribute.
// FlagExtendedLength returns a bool
func (obj *bgpAttributesOtherAttribute) FlagExtendedLength() bool {

	return *obj.obj.FlagExtendedLength

}

// Extended length flag in the BGP attribute.
// FlagExtendedLength returns a bool
func (obj *bgpAttributesOtherAttribute) HasFlagExtendedLength() bool {
	return obj.obj.FlagExtendedLength != nil
}

// Extended length flag in the BGP attribute.
// SetFlagExtendedLength sets the bool value in the BgpAttributesOtherAttribute object
func (obj *bgpAttributesOtherAttribute) SetFlagExtendedLength(value bool) BgpAttributesOtherAttribute {

	obj.obj.FlagExtendedLength = &value
	return obj
}

// The value of the Type field in the attribute.
// Type returns a uint32
func (obj *bgpAttributesOtherAttribute) Type() uint32 {

	return *obj.obj.Type

}

// The value of the Type field in the attribute.
// SetType sets the uint32 value in the BgpAttributesOtherAttribute object
func (obj *bgpAttributesOtherAttribute) SetType(value uint32) BgpAttributesOtherAttribute {

	obj.obj.Type = &value
	return obj
}

// Contents of the value field ( the contents after the initial two bytes containing the Flags and Type field ) of the attribute in hex bytes.
// It includes the contents of length of the extended length field if included.
// RawValue returns a string
func (obj *bgpAttributesOtherAttribute) RawValue() string {

	return *obj.obj.RawValue

}

// Contents of the value field ( the contents after the initial two bytes containing the Flags and Type field ) of the attribute in hex bytes.
// It includes the contents of length of the extended length field if included.
// SetRawValue sets the string value in the BgpAttributesOtherAttribute object
func (obj *bgpAttributesOtherAttribute) SetRawValue(value string) BgpAttributesOtherAttribute {

	obj.obj.RawValue = &value
	return obj
}

func (obj *bgpAttributesOtherAttribute) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Type is required
	if obj.obj.Type == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Type is required field on interface BgpAttributesOtherAttribute")
	}

	// RawValue is required
	if obj.obj.RawValue == nil {
		vObj.validationErrors = append(vObj.validationErrors, "RawValue is required field on interface BgpAttributesOtherAttribute")
	}
	if obj.obj.RawValue != nil {

		err := obj.validateHex(obj.RawValue())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesOtherAttribute.RawValue"))
		}

	}

}

func (obj *bgpAttributesOtherAttribute) setDefault() {
	if obj.obj.FlagOptional == nil {
		obj.SetFlagOptional(false)
	}
	if obj.obj.FlagTransitive == nil {
		obj.SetFlagTransitive(false)
	}
	if obj.obj.FlagPartial == nil {
		obj.SetFlagPartial(false)
	}
	if obj.obj.FlagExtendedLength == nil {
		obj.SetFlagExtendedLength(false)
	}

}

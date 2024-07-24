package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpExtendedCommunityTransitiveOpaqueTypeColor *****
type bgpExtendedCommunityTransitiveOpaqueTypeColor struct {
	validation
	obj          *otg.BgpExtendedCommunityTransitiveOpaqueTypeColor
	marshaller   marshalBgpExtendedCommunityTransitiveOpaqueTypeColor
	unMarshaller unMarshalBgpExtendedCommunityTransitiveOpaqueTypeColor
}

func NewBgpExtendedCommunityTransitiveOpaqueTypeColor() BgpExtendedCommunityTransitiveOpaqueTypeColor {
	obj := bgpExtendedCommunityTransitiveOpaqueTypeColor{obj: &otg.BgpExtendedCommunityTransitiveOpaqueTypeColor{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpExtendedCommunityTransitiveOpaqueTypeColor) msg() *otg.BgpExtendedCommunityTransitiveOpaqueTypeColor {
	return obj.obj
}

func (obj *bgpExtendedCommunityTransitiveOpaqueTypeColor) setMsg(msg *otg.BgpExtendedCommunityTransitiveOpaqueTypeColor) BgpExtendedCommunityTransitiveOpaqueTypeColor {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpExtendedCommunityTransitiveOpaqueTypeColor struct {
	obj *bgpExtendedCommunityTransitiveOpaqueTypeColor
}

type marshalBgpExtendedCommunityTransitiveOpaqueTypeColor interface {
	// ToProto marshals BgpExtendedCommunityTransitiveOpaqueTypeColor to protobuf object *otg.BgpExtendedCommunityTransitiveOpaqueTypeColor
	ToProto() (*otg.BgpExtendedCommunityTransitiveOpaqueTypeColor, error)
	// ToPbText marshals BgpExtendedCommunityTransitiveOpaqueTypeColor to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpExtendedCommunityTransitiveOpaqueTypeColor to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpExtendedCommunityTransitiveOpaqueTypeColor to JSON text
	ToJson() (string, error)
}

type unMarshalbgpExtendedCommunityTransitiveOpaqueTypeColor struct {
	obj *bgpExtendedCommunityTransitiveOpaqueTypeColor
}

type unMarshalBgpExtendedCommunityTransitiveOpaqueTypeColor interface {
	// FromProto unmarshals BgpExtendedCommunityTransitiveOpaqueTypeColor from protobuf object *otg.BgpExtendedCommunityTransitiveOpaqueTypeColor
	FromProto(msg *otg.BgpExtendedCommunityTransitiveOpaqueTypeColor) (BgpExtendedCommunityTransitiveOpaqueTypeColor, error)
	// FromPbText unmarshals BgpExtendedCommunityTransitiveOpaqueTypeColor from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpExtendedCommunityTransitiveOpaqueTypeColor from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpExtendedCommunityTransitiveOpaqueTypeColor from JSON text
	FromJson(value string) error
}

func (obj *bgpExtendedCommunityTransitiveOpaqueTypeColor) Marshal() marshalBgpExtendedCommunityTransitiveOpaqueTypeColor {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpExtendedCommunityTransitiveOpaqueTypeColor{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpExtendedCommunityTransitiveOpaqueTypeColor) Unmarshal() unMarshalBgpExtendedCommunityTransitiveOpaqueTypeColor {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpExtendedCommunityTransitiveOpaqueTypeColor{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpExtendedCommunityTransitiveOpaqueTypeColor) ToProto() (*otg.BgpExtendedCommunityTransitiveOpaqueTypeColor, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpExtendedCommunityTransitiveOpaqueTypeColor) FromProto(msg *otg.BgpExtendedCommunityTransitiveOpaqueTypeColor) (BgpExtendedCommunityTransitiveOpaqueTypeColor, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpExtendedCommunityTransitiveOpaqueTypeColor) ToPbText() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveOpaqueTypeColor) FromPbText(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitiveOpaqueTypeColor) ToYaml() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveOpaqueTypeColor) FromYaml(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitiveOpaqueTypeColor) ToJson() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveOpaqueTypeColor) FromJson(value string) error {
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

func (obj *bgpExtendedCommunityTransitiveOpaqueTypeColor) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitiveOpaqueTypeColor) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitiveOpaqueTypeColor) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpExtendedCommunityTransitiveOpaqueTypeColor) Clone() (BgpExtendedCommunityTransitiveOpaqueTypeColor, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpExtendedCommunityTransitiveOpaqueTypeColor()
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

// BgpExtendedCommunityTransitiveOpaqueTypeColor is the Color Community contains locally administrator defined 'color' value which is used in conjunction with Encapsulation  attribute to decide whether a data packet can be transmitted on a certain tunnel or not. It is defined in RFC9012 and sent with sub-type as 0x0b.
type BgpExtendedCommunityTransitiveOpaqueTypeColor interface {
	Validation
	// msg marshals BgpExtendedCommunityTransitiveOpaqueTypeColor to protobuf object *otg.BgpExtendedCommunityTransitiveOpaqueTypeColor
	// and doesn't set defaults
	msg() *otg.BgpExtendedCommunityTransitiveOpaqueTypeColor
	// setMsg unmarshals BgpExtendedCommunityTransitiveOpaqueTypeColor from protobuf object *otg.BgpExtendedCommunityTransitiveOpaqueTypeColor
	// and doesn't set defaults
	setMsg(*otg.BgpExtendedCommunityTransitiveOpaqueTypeColor) BgpExtendedCommunityTransitiveOpaqueTypeColor
	// provides marshal interface
	Marshal() marshalBgpExtendedCommunityTransitiveOpaqueTypeColor
	// provides unmarshal interface
	Unmarshal() unMarshalBgpExtendedCommunityTransitiveOpaqueTypeColor
	// validate validates BgpExtendedCommunityTransitiveOpaqueTypeColor
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpExtendedCommunityTransitiveOpaqueTypeColor, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns uint32, set in BgpExtendedCommunityTransitiveOpaqueTypeColor.
	Flags() uint32
	// SetFlags assigns uint32 provided by user to BgpExtendedCommunityTransitiveOpaqueTypeColor
	SetFlags(value uint32) BgpExtendedCommunityTransitiveOpaqueTypeColor
	// HasFlags checks if Flags has been set in BgpExtendedCommunityTransitiveOpaqueTypeColor
	HasFlags() bool
	// Color returns uint32, set in BgpExtendedCommunityTransitiveOpaqueTypeColor.
	Color() uint32
	// SetColor assigns uint32 provided by user to BgpExtendedCommunityTransitiveOpaqueTypeColor
	SetColor(value uint32) BgpExtendedCommunityTransitiveOpaqueTypeColor
	// HasColor checks if Color has been set in BgpExtendedCommunityTransitiveOpaqueTypeColor
	HasColor() bool
}

// Two octet flag values.
// Flags returns a uint32
func (obj *bgpExtendedCommunityTransitiveOpaqueTypeColor) Flags() uint32 {

	return *obj.obj.Flags

}

// Two octet flag values.
// Flags returns a uint32
func (obj *bgpExtendedCommunityTransitiveOpaqueTypeColor) HasFlags() bool {
	return obj.obj.Flags != nil
}

// Two octet flag values.
// SetFlags sets the uint32 value in the BgpExtendedCommunityTransitiveOpaqueTypeColor object
func (obj *bgpExtendedCommunityTransitiveOpaqueTypeColor) SetFlags(value uint32) BgpExtendedCommunityTransitiveOpaqueTypeColor {

	obj.obj.Flags = &value
	return obj
}

// The color value is user defined and configured locally and used to determine whether a data packet can be transmitted on a certain tunnel or not in conjunction with the Encapsulation attribute. It is defined in RFC9012.
// Color returns a uint32
func (obj *bgpExtendedCommunityTransitiveOpaqueTypeColor) Color() uint32 {

	return *obj.obj.Color

}

// The color value is user defined and configured locally and used to determine whether a data packet can be transmitted on a certain tunnel or not in conjunction with the Encapsulation attribute. It is defined in RFC9012.
// Color returns a uint32
func (obj *bgpExtendedCommunityTransitiveOpaqueTypeColor) HasColor() bool {
	return obj.obj.Color != nil
}

// The color value is user defined and configured locally and used to determine whether a data packet can be transmitted on a certain tunnel or not in conjunction with the Encapsulation attribute. It is defined in RFC9012.
// SetColor sets the uint32 value in the BgpExtendedCommunityTransitiveOpaqueTypeColor object
func (obj *bgpExtendedCommunityTransitiveOpaqueTypeColor) SetColor(value uint32) BgpExtendedCommunityTransitiveOpaqueTypeColor {

	obj.obj.Color = &value
	return obj
}

func (obj *bgpExtendedCommunityTransitiveOpaqueTypeColor) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		if *obj.obj.Flags > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpExtendedCommunityTransitiveOpaqueTypeColor.Flags <= 65535 but Got %d", *obj.obj.Flags))
		}

	}

}

func (obj *bgpExtendedCommunityTransitiveOpaqueTypeColor) setDefault() {
	if obj.obj.Flags == nil {
		obj.SetFlags(0)
	}
	if obj.obj.Color == nil {
		obj.SetColor(0)
	}

}

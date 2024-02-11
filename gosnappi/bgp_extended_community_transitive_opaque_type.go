package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpExtendedCommunityTransitiveOpaqueType *****
type bgpExtendedCommunityTransitiveOpaqueType struct {
	validation
	obj                        *otg.BgpExtendedCommunityTransitiveOpaqueType
	marshaller                 marshalBgpExtendedCommunityTransitiveOpaqueType
	unMarshaller               unMarshalBgpExtendedCommunityTransitiveOpaqueType
	colorSubtypeHolder         BgpExtendedCommunityTransitiveOpaqueTypeColor
	encapsulationSubtypeHolder BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
}

func NewBgpExtendedCommunityTransitiveOpaqueType() BgpExtendedCommunityTransitiveOpaqueType {
	obj := bgpExtendedCommunityTransitiveOpaqueType{obj: &otg.BgpExtendedCommunityTransitiveOpaqueType{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpExtendedCommunityTransitiveOpaqueType) msg() *otg.BgpExtendedCommunityTransitiveOpaqueType {
	return obj.obj
}

func (obj *bgpExtendedCommunityTransitiveOpaqueType) setMsg(msg *otg.BgpExtendedCommunityTransitiveOpaqueType) BgpExtendedCommunityTransitiveOpaqueType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpExtendedCommunityTransitiveOpaqueType struct {
	obj *bgpExtendedCommunityTransitiveOpaqueType
}

type marshalBgpExtendedCommunityTransitiveOpaqueType interface {
	// ToProto marshals BgpExtendedCommunityTransitiveOpaqueType to protobuf object *otg.BgpExtendedCommunityTransitiveOpaqueType
	ToProto() (*otg.BgpExtendedCommunityTransitiveOpaqueType, error)
	// ToPbText marshals BgpExtendedCommunityTransitiveOpaqueType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpExtendedCommunityTransitiveOpaqueType to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpExtendedCommunityTransitiveOpaqueType to JSON text
	ToJson() (string, error)
}

type unMarshalbgpExtendedCommunityTransitiveOpaqueType struct {
	obj *bgpExtendedCommunityTransitiveOpaqueType
}

type unMarshalBgpExtendedCommunityTransitiveOpaqueType interface {
	// FromProto unmarshals BgpExtendedCommunityTransitiveOpaqueType from protobuf object *otg.BgpExtendedCommunityTransitiveOpaqueType
	FromProto(msg *otg.BgpExtendedCommunityTransitiveOpaqueType) (BgpExtendedCommunityTransitiveOpaqueType, error)
	// FromPbText unmarshals BgpExtendedCommunityTransitiveOpaqueType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpExtendedCommunityTransitiveOpaqueType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpExtendedCommunityTransitiveOpaqueType from JSON text
	FromJson(value string) error
}

func (obj *bgpExtendedCommunityTransitiveOpaqueType) Marshal() marshalBgpExtendedCommunityTransitiveOpaqueType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpExtendedCommunityTransitiveOpaqueType{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpExtendedCommunityTransitiveOpaqueType) Unmarshal() unMarshalBgpExtendedCommunityTransitiveOpaqueType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpExtendedCommunityTransitiveOpaqueType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpExtendedCommunityTransitiveOpaqueType) ToProto() (*otg.BgpExtendedCommunityTransitiveOpaqueType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpExtendedCommunityTransitiveOpaqueType) FromProto(msg *otg.BgpExtendedCommunityTransitiveOpaqueType) (BgpExtendedCommunityTransitiveOpaqueType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpExtendedCommunityTransitiveOpaqueType) ToPbText() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveOpaqueType) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalbgpExtendedCommunityTransitiveOpaqueType) ToYaml() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveOpaqueType) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalbgpExtendedCommunityTransitiveOpaqueType) ToJson() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveOpaqueType) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *bgpExtendedCommunityTransitiveOpaqueType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitiveOpaqueType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitiveOpaqueType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpExtendedCommunityTransitiveOpaqueType) Clone() (BgpExtendedCommunityTransitiveOpaqueType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpExtendedCommunityTransitiveOpaqueType()
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

func (obj *bgpExtendedCommunityTransitiveOpaqueType) setNil() {
	obj.colorSubtypeHolder = nil
	obj.encapsulationSubtypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpExtendedCommunityTransitiveOpaqueType is the Transitive Opaque Extended Community is sent as type 0x03.
type BgpExtendedCommunityTransitiveOpaqueType interface {
	Validation
	// msg marshals BgpExtendedCommunityTransitiveOpaqueType to protobuf object *otg.BgpExtendedCommunityTransitiveOpaqueType
	// and doesn't set defaults
	msg() *otg.BgpExtendedCommunityTransitiveOpaqueType
	// setMsg unmarshals BgpExtendedCommunityTransitiveOpaqueType from protobuf object *otg.BgpExtendedCommunityTransitiveOpaqueType
	// and doesn't set defaults
	setMsg(*otg.BgpExtendedCommunityTransitiveOpaqueType) BgpExtendedCommunityTransitiveOpaqueType
	// provides marshal interface
	Marshal() marshalBgpExtendedCommunityTransitiveOpaqueType
	// provides unmarshal interface
	Unmarshal() unMarshalBgpExtendedCommunityTransitiveOpaqueType
	// validate validates BgpExtendedCommunityTransitiveOpaqueType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpExtendedCommunityTransitiveOpaqueType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpExtendedCommunityTransitiveOpaqueTypeChoiceEnum, set in BgpExtendedCommunityTransitiveOpaqueType
	Choice() BgpExtendedCommunityTransitiveOpaqueTypeChoiceEnum
	// setChoice assigns BgpExtendedCommunityTransitiveOpaqueTypeChoiceEnum provided by user to BgpExtendedCommunityTransitiveOpaqueType
	setChoice(value BgpExtendedCommunityTransitiveOpaqueTypeChoiceEnum) BgpExtendedCommunityTransitiveOpaqueType
	// HasChoice checks if Choice has been set in BgpExtendedCommunityTransitiveOpaqueType
	HasChoice() bool
	// ColorSubtype returns BgpExtendedCommunityTransitiveOpaqueTypeColor, set in BgpExtendedCommunityTransitiveOpaqueType.
	// BgpExtendedCommunityTransitiveOpaqueTypeColor is the Color Community contains locally administrator defined 'color' value which is used in conjunction with Encapsulation  attribute to decide whether a data packet can be transmitted on a certain tunnel or not. It is defined in RFC9012 and sent with sub-type as 0x0b.
	ColorSubtype() BgpExtendedCommunityTransitiveOpaqueTypeColor
	// SetColorSubtype assigns BgpExtendedCommunityTransitiveOpaqueTypeColor provided by user to BgpExtendedCommunityTransitiveOpaqueType.
	// BgpExtendedCommunityTransitiveOpaqueTypeColor is the Color Community contains locally administrator defined 'color' value which is used in conjunction with Encapsulation  attribute to decide whether a data packet can be transmitted on a certain tunnel or not. It is defined in RFC9012 and sent with sub-type as 0x0b.
	SetColorSubtype(value BgpExtendedCommunityTransitiveOpaqueTypeColor) BgpExtendedCommunityTransitiveOpaqueType
	// HasColorSubtype checks if ColorSubtype has been set in BgpExtendedCommunityTransitiveOpaqueType
	HasColorSubtype() bool
	// EncapsulationSubtype returns BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation, set in BgpExtendedCommunityTransitiveOpaqueType.
	// BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation is this identifies the type of tunneling technology being signalled. It is defined in RFC9012 and sent with sub-type as 0x0c.
	EncapsulationSubtype() BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
	// SetEncapsulationSubtype assigns BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation provided by user to BgpExtendedCommunityTransitiveOpaqueType.
	// BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation is this identifies the type of tunneling technology being signalled. It is defined in RFC9012 and sent with sub-type as 0x0c.
	SetEncapsulationSubtype(value BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) BgpExtendedCommunityTransitiveOpaqueType
	// HasEncapsulationSubtype checks if EncapsulationSubtype has been set in BgpExtendedCommunityTransitiveOpaqueType
	HasEncapsulationSubtype() bool
	setNil()
}

type BgpExtendedCommunityTransitiveOpaqueTypeChoiceEnum string

// Enum of Choice on BgpExtendedCommunityTransitiveOpaqueType
var BgpExtendedCommunityTransitiveOpaqueTypeChoice = struct {
	COLOR_SUBTYPE         BgpExtendedCommunityTransitiveOpaqueTypeChoiceEnum
	ENCAPSULATION_SUBTYPE BgpExtendedCommunityTransitiveOpaqueTypeChoiceEnum
}{
	COLOR_SUBTYPE:         BgpExtendedCommunityTransitiveOpaqueTypeChoiceEnum("color_subtype"),
	ENCAPSULATION_SUBTYPE: BgpExtendedCommunityTransitiveOpaqueTypeChoiceEnum("encapsulation_subtype"),
}

func (obj *bgpExtendedCommunityTransitiveOpaqueType) Choice() BgpExtendedCommunityTransitiveOpaqueTypeChoiceEnum {
	return BgpExtendedCommunityTransitiveOpaqueTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *bgpExtendedCommunityTransitiveOpaqueType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpExtendedCommunityTransitiveOpaqueType) setChoice(value BgpExtendedCommunityTransitiveOpaqueTypeChoiceEnum) BgpExtendedCommunityTransitiveOpaqueType {
	intValue, ok := otg.BgpExtendedCommunityTransitiveOpaqueType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpExtendedCommunityTransitiveOpaqueTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpExtendedCommunityTransitiveOpaqueType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.EncapsulationSubtype = nil
	obj.encapsulationSubtypeHolder = nil
	obj.obj.ColorSubtype = nil
	obj.colorSubtypeHolder = nil

	if value == BgpExtendedCommunityTransitiveOpaqueTypeChoice.COLOR_SUBTYPE {
		obj.obj.ColorSubtype = NewBgpExtendedCommunityTransitiveOpaqueTypeColor().msg()
	}

	if value == BgpExtendedCommunityTransitiveOpaqueTypeChoice.ENCAPSULATION_SUBTYPE {
		obj.obj.EncapsulationSubtype = NewBgpExtendedCommunityTransitiveOpaqueTypeEncapsulation().msg()
	}

	return obj
}

// description is TBD
// ColorSubtype returns a BgpExtendedCommunityTransitiveOpaqueTypeColor
func (obj *bgpExtendedCommunityTransitiveOpaqueType) ColorSubtype() BgpExtendedCommunityTransitiveOpaqueTypeColor {
	if obj.obj.ColorSubtype == nil {
		obj.setChoice(BgpExtendedCommunityTransitiveOpaqueTypeChoice.COLOR_SUBTYPE)
	}
	if obj.colorSubtypeHolder == nil {
		obj.colorSubtypeHolder = &bgpExtendedCommunityTransitiveOpaqueTypeColor{obj: obj.obj.ColorSubtype}
	}
	return obj.colorSubtypeHolder
}

// description is TBD
// ColorSubtype returns a BgpExtendedCommunityTransitiveOpaqueTypeColor
func (obj *bgpExtendedCommunityTransitiveOpaqueType) HasColorSubtype() bool {
	return obj.obj.ColorSubtype != nil
}

// description is TBD
// SetColorSubtype sets the BgpExtendedCommunityTransitiveOpaqueTypeColor value in the BgpExtendedCommunityTransitiveOpaqueType object
func (obj *bgpExtendedCommunityTransitiveOpaqueType) SetColorSubtype(value BgpExtendedCommunityTransitiveOpaqueTypeColor) BgpExtendedCommunityTransitiveOpaqueType {
	obj.setChoice(BgpExtendedCommunityTransitiveOpaqueTypeChoice.COLOR_SUBTYPE)
	obj.colorSubtypeHolder = nil
	obj.obj.ColorSubtype = value.msg()

	return obj
}

// description is TBD
// EncapsulationSubtype returns a BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
func (obj *bgpExtendedCommunityTransitiveOpaqueType) EncapsulationSubtype() BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation {
	if obj.obj.EncapsulationSubtype == nil {
		obj.setChoice(BgpExtendedCommunityTransitiveOpaqueTypeChoice.ENCAPSULATION_SUBTYPE)
	}
	if obj.encapsulationSubtypeHolder == nil {
		obj.encapsulationSubtypeHolder = &bgpExtendedCommunityTransitiveOpaqueTypeEncapsulation{obj: obj.obj.EncapsulationSubtype}
	}
	return obj.encapsulationSubtypeHolder
}

// description is TBD
// EncapsulationSubtype returns a BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
func (obj *bgpExtendedCommunityTransitiveOpaqueType) HasEncapsulationSubtype() bool {
	return obj.obj.EncapsulationSubtype != nil
}

// description is TBD
// SetEncapsulationSubtype sets the BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation value in the BgpExtendedCommunityTransitiveOpaqueType object
func (obj *bgpExtendedCommunityTransitiveOpaqueType) SetEncapsulationSubtype(value BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) BgpExtendedCommunityTransitiveOpaqueType {
	obj.setChoice(BgpExtendedCommunityTransitiveOpaqueTypeChoice.ENCAPSULATION_SUBTYPE)
	obj.encapsulationSubtypeHolder = nil
	obj.obj.EncapsulationSubtype = value.msg()

	return obj
}

func (obj *bgpExtendedCommunityTransitiveOpaqueType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ColorSubtype != nil {

		obj.ColorSubtype().validateObj(vObj, set_default)
	}

	if obj.obj.EncapsulationSubtype != nil {

		obj.EncapsulationSubtype().validateObj(vObj, set_default)
	}

}

func (obj *bgpExtendedCommunityTransitiveOpaqueType) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(BgpExtendedCommunityTransitiveOpaqueTypeChoice.COLOR_SUBTYPE)

	}

}

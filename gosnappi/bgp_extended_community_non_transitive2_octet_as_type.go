package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpExtendedCommunityNonTransitive2OctetAsType *****
type bgpExtendedCommunityNonTransitive2OctetAsType struct {
	validation
	obj                        *otg.BgpExtendedCommunityNonTransitive2OctetAsType
	marshaller                 marshalBgpExtendedCommunityNonTransitive2OctetAsType
	unMarshaller               unMarshalBgpExtendedCommunityNonTransitive2OctetAsType
	linkBandwidthSubtypeHolder BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
}

func NewBgpExtendedCommunityNonTransitive2OctetAsType() BgpExtendedCommunityNonTransitive2OctetAsType {
	obj := bgpExtendedCommunityNonTransitive2OctetAsType{obj: &otg.BgpExtendedCommunityNonTransitive2OctetAsType{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpExtendedCommunityNonTransitive2OctetAsType) msg() *otg.BgpExtendedCommunityNonTransitive2OctetAsType {
	return obj.obj
}

func (obj *bgpExtendedCommunityNonTransitive2OctetAsType) setMsg(msg *otg.BgpExtendedCommunityNonTransitive2OctetAsType) BgpExtendedCommunityNonTransitive2OctetAsType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpExtendedCommunityNonTransitive2OctetAsType struct {
	obj *bgpExtendedCommunityNonTransitive2OctetAsType
}

type marshalBgpExtendedCommunityNonTransitive2OctetAsType interface {
	// ToProto marshals BgpExtendedCommunityNonTransitive2OctetAsType to protobuf object *otg.BgpExtendedCommunityNonTransitive2OctetAsType
	ToProto() (*otg.BgpExtendedCommunityNonTransitive2OctetAsType, error)
	// ToPbText marshals BgpExtendedCommunityNonTransitive2OctetAsType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpExtendedCommunityNonTransitive2OctetAsType to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpExtendedCommunityNonTransitive2OctetAsType to JSON text
	ToJson() (string, error)
}

type unMarshalbgpExtendedCommunityNonTransitive2OctetAsType struct {
	obj *bgpExtendedCommunityNonTransitive2OctetAsType
}

type unMarshalBgpExtendedCommunityNonTransitive2OctetAsType interface {
	// FromProto unmarshals BgpExtendedCommunityNonTransitive2OctetAsType from protobuf object *otg.BgpExtendedCommunityNonTransitive2OctetAsType
	FromProto(msg *otg.BgpExtendedCommunityNonTransitive2OctetAsType) (BgpExtendedCommunityNonTransitive2OctetAsType, error)
	// FromPbText unmarshals BgpExtendedCommunityNonTransitive2OctetAsType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpExtendedCommunityNonTransitive2OctetAsType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpExtendedCommunityNonTransitive2OctetAsType from JSON text
	FromJson(value string) error
}

func (obj *bgpExtendedCommunityNonTransitive2OctetAsType) Marshal() marshalBgpExtendedCommunityNonTransitive2OctetAsType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpExtendedCommunityNonTransitive2OctetAsType{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpExtendedCommunityNonTransitive2OctetAsType) Unmarshal() unMarshalBgpExtendedCommunityNonTransitive2OctetAsType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpExtendedCommunityNonTransitive2OctetAsType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpExtendedCommunityNonTransitive2OctetAsType) ToProto() (*otg.BgpExtendedCommunityNonTransitive2OctetAsType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpExtendedCommunityNonTransitive2OctetAsType) FromProto(msg *otg.BgpExtendedCommunityNonTransitive2OctetAsType) (BgpExtendedCommunityNonTransitive2OctetAsType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpExtendedCommunityNonTransitive2OctetAsType) ToPbText() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityNonTransitive2OctetAsType) FromPbText(value string) error {
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

func (m *marshalbgpExtendedCommunityNonTransitive2OctetAsType) ToYaml() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityNonTransitive2OctetAsType) FromYaml(value string) error {
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

func (m *marshalbgpExtendedCommunityNonTransitive2OctetAsType) ToJson() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityNonTransitive2OctetAsType) FromJson(value string) error {
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

func (obj *bgpExtendedCommunityNonTransitive2OctetAsType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityNonTransitive2OctetAsType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityNonTransitive2OctetAsType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpExtendedCommunityNonTransitive2OctetAsType) Clone() (BgpExtendedCommunityNonTransitive2OctetAsType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpExtendedCommunityNonTransitive2OctetAsType()
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

func (obj *bgpExtendedCommunityNonTransitive2OctetAsType) setNil() {
	obj.linkBandwidthSubtypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpExtendedCommunityNonTransitive2OctetAsType is the Non-Transitive Two-Octet AS-Specific Extended Community is sent as type 0x40.
type BgpExtendedCommunityNonTransitive2OctetAsType interface {
	Validation
	// msg marshals BgpExtendedCommunityNonTransitive2OctetAsType to protobuf object *otg.BgpExtendedCommunityNonTransitive2OctetAsType
	// and doesn't set defaults
	msg() *otg.BgpExtendedCommunityNonTransitive2OctetAsType
	// setMsg unmarshals BgpExtendedCommunityNonTransitive2OctetAsType from protobuf object *otg.BgpExtendedCommunityNonTransitive2OctetAsType
	// and doesn't set defaults
	setMsg(*otg.BgpExtendedCommunityNonTransitive2OctetAsType) BgpExtendedCommunityNonTransitive2OctetAsType
	// provides marshal interface
	Marshal() marshalBgpExtendedCommunityNonTransitive2OctetAsType
	// provides unmarshal interface
	Unmarshal() unMarshalBgpExtendedCommunityNonTransitive2OctetAsType
	// validate validates BgpExtendedCommunityNonTransitive2OctetAsType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpExtendedCommunityNonTransitive2OctetAsType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum, set in BgpExtendedCommunityNonTransitive2OctetAsType
	Choice() BgpExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum
	// setChoice assigns BgpExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum provided by user to BgpExtendedCommunityNonTransitive2OctetAsType
	setChoice(value BgpExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum) BgpExtendedCommunityNonTransitive2OctetAsType
	// HasChoice checks if Choice has been set in BgpExtendedCommunityNonTransitive2OctetAsType
	HasChoice() bool
	// LinkBandwidthSubtype returns BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth, set in BgpExtendedCommunityNonTransitive2OctetAsType.
	// BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth is the Link Bandwidth Extended Community attribute is defined in draft-ietf-idr-link-bandwidth. It is sent with sub-type as 0x04.
	LinkBandwidthSubtype() BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	// SetLinkBandwidthSubtype assigns BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth provided by user to BgpExtendedCommunityNonTransitive2OctetAsType.
	// BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth is the Link Bandwidth Extended Community attribute is defined in draft-ietf-idr-link-bandwidth. It is sent with sub-type as 0x04.
	SetLinkBandwidthSubtype(value BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) BgpExtendedCommunityNonTransitive2OctetAsType
	// HasLinkBandwidthSubtype checks if LinkBandwidthSubtype has been set in BgpExtendedCommunityNonTransitive2OctetAsType
	HasLinkBandwidthSubtype() bool
	setNil()
}

type BgpExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum string

// Enum of Choice on BgpExtendedCommunityNonTransitive2OctetAsType
var BgpExtendedCommunityNonTransitive2OctetAsTypeChoice = struct {
	LINK_BANDWIDTH_SUBTYPE BgpExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum
}{
	LINK_BANDWIDTH_SUBTYPE: BgpExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum("link_bandwidth_subtype"),
}

func (obj *bgpExtendedCommunityNonTransitive2OctetAsType) Choice() BgpExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum {
	return BgpExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *bgpExtendedCommunityNonTransitive2OctetAsType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpExtendedCommunityNonTransitive2OctetAsType) setChoice(value BgpExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum) BgpExtendedCommunityNonTransitive2OctetAsType {
	intValue, ok := otg.BgpExtendedCommunityNonTransitive2OctetAsType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpExtendedCommunityNonTransitive2OctetAsType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.LinkBandwidthSubtype = nil
	obj.linkBandwidthSubtypeHolder = nil

	if value == BgpExtendedCommunityNonTransitive2OctetAsTypeChoice.LINK_BANDWIDTH_SUBTYPE {
		obj.obj.LinkBandwidthSubtype = NewBgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth().msg()
	}

	return obj
}

// description is TBD
// LinkBandwidthSubtype returns a BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
func (obj *bgpExtendedCommunityNonTransitive2OctetAsType) LinkBandwidthSubtype() BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth {
	if obj.obj.LinkBandwidthSubtype == nil {
		obj.setChoice(BgpExtendedCommunityNonTransitive2OctetAsTypeChoice.LINK_BANDWIDTH_SUBTYPE)
	}
	if obj.linkBandwidthSubtypeHolder == nil {
		obj.linkBandwidthSubtypeHolder = &bgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth{obj: obj.obj.LinkBandwidthSubtype}
	}
	return obj.linkBandwidthSubtypeHolder
}

// description is TBD
// LinkBandwidthSubtype returns a BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
func (obj *bgpExtendedCommunityNonTransitive2OctetAsType) HasLinkBandwidthSubtype() bool {
	return obj.obj.LinkBandwidthSubtype != nil
}

// description is TBD
// SetLinkBandwidthSubtype sets the BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth value in the BgpExtendedCommunityNonTransitive2OctetAsType object
func (obj *bgpExtendedCommunityNonTransitive2OctetAsType) SetLinkBandwidthSubtype(value BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) BgpExtendedCommunityNonTransitive2OctetAsType {
	obj.setChoice(BgpExtendedCommunityNonTransitive2OctetAsTypeChoice.LINK_BANDWIDTH_SUBTYPE)
	obj.linkBandwidthSubtypeHolder = nil
	obj.obj.LinkBandwidthSubtype = value.msg()

	return obj
}

func (obj *bgpExtendedCommunityNonTransitive2OctetAsType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LinkBandwidthSubtype != nil {

		obj.LinkBandwidthSubtype().validateObj(vObj, set_default)
	}

}

func (obj *bgpExtendedCommunityNonTransitive2OctetAsType) setDefault() {
	var choices_set int = 0
	var choice BgpExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum

	if obj.obj.LinkBandwidthSubtype != nil {
		choices_set += 1
		choice = BgpExtendedCommunityNonTransitive2OctetAsTypeChoice.LINK_BANDWIDTH_SUBTYPE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(BgpExtendedCommunityNonTransitive2OctetAsTypeChoice.LINK_BANDWIDTH_SUBTYPE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpExtendedCommunityNonTransitive2OctetAsType")
			}
		} else {
			intVal := otg.BgpExtendedCommunityNonTransitive2OctetAsType_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpExtendedCommunityNonTransitive2OctetAsType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

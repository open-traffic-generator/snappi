package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpExtendedCommunityTransitiveEvpnType *****
type bgpExtendedCommunityTransitiveEvpnType struct {
	validation
	obj                    *otg.BgpExtendedCommunityTransitiveEvpnType
	marshaller             marshalBgpExtendedCommunityTransitiveEvpnType
	unMarshaller           unMarshalBgpExtendedCommunityTransitiveEvpnType
	routerMacSubtypeHolder BgpExtendedCommunityTransitiveEvpnTypeRouterMac
}

func NewBgpExtendedCommunityTransitiveEvpnType() BgpExtendedCommunityTransitiveEvpnType {
	obj := bgpExtendedCommunityTransitiveEvpnType{obj: &otg.BgpExtendedCommunityTransitiveEvpnType{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpExtendedCommunityTransitiveEvpnType) msg() *otg.BgpExtendedCommunityTransitiveEvpnType {
	return obj.obj
}

func (obj *bgpExtendedCommunityTransitiveEvpnType) setMsg(msg *otg.BgpExtendedCommunityTransitiveEvpnType) BgpExtendedCommunityTransitiveEvpnType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpExtendedCommunityTransitiveEvpnType struct {
	obj *bgpExtendedCommunityTransitiveEvpnType
}

type marshalBgpExtendedCommunityTransitiveEvpnType interface {
	// ToProto marshals BgpExtendedCommunityTransitiveEvpnType to protobuf object *otg.BgpExtendedCommunityTransitiveEvpnType
	ToProto() (*otg.BgpExtendedCommunityTransitiveEvpnType, error)
	// ToPbText marshals BgpExtendedCommunityTransitiveEvpnType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpExtendedCommunityTransitiveEvpnType to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpExtendedCommunityTransitiveEvpnType to JSON text
	ToJson() (string, error)
}

type unMarshalbgpExtendedCommunityTransitiveEvpnType struct {
	obj *bgpExtendedCommunityTransitiveEvpnType
}

type unMarshalBgpExtendedCommunityTransitiveEvpnType interface {
	// FromProto unmarshals BgpExtendedCommunityTransitiveEvpnType from protobuf object *otg.BgpExtendedCommunityTransitiveEvpnType
	FromProto(msg *otg.BgpExtendedCommunityTransitiveEvpnType) (BgpExtendedCommunityTransitiveEvpnType, error)
	// FromPbText unmarshals BgpExtendedCommunityTransitiveEvpnType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpExtendedCommunityTransitiveEvpnType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpExtendedCommunityTransitiveEvpnType from JSON text
	FromJson(value string) error
}

func (obj *bgpExtendedCommunityTransitiveEvpnType) Marshal() marshalBgpExtendedCommunityTransitiveEvpnType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpExtendedCommunityTransitiveEvpnType{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpExtendedCommunityTransitiveEvpnType) Unmarshal() unMarshalBgpExtendedCommunityTransitiveEvpnType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpExtendedCommunityTransitiveEvpnType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpExtendedCommunityTransitiveEvpnType) ToProto() (*otg.BgpExtendedCommunityTransitiveEvpnType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpExtendedCommunityTransitiveEvpnType) FromProto(msg *otg.BgpExtendedCommunityTransitiveEvpnType) (BgpExtendedCommunityTransitiveEvpnType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpExtendedCommunityTransitiveEvpnType) ToPbText() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveEvpnType) FromPbText(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitiveEvpnType) ToYaml() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveEvpnType) FromYaml(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitiveEvpnType) ToJson() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveEvpnType) FromJson(value string) error {
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

func (obj *bgpExtendedCommunityTransitiveEvpnType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitiveEvpnType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitiveEvpnType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpExtendedCommunityTransitiveEvpnType) Clone() (BgpExtendedCommunityTransitiveEvpnType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpExtendedCommunityTransitiveEvpnType()
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

func (obj *bgpExtendedCommunityTransitiveEvpnType) setNil() {
	obj.routerMacSubtypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpExtendedCommunityTransitiveEvpnType is the Transitive EVPN Extended Community is sent as type 0x06 .
type BgpExtendedCommunityTransitiveEvpnType interface {
	Validation
	// msg marshals BgpExtendedCommunityTransitiveEvpnType to protobuf object *otg.BgpExtendedCommunityTransitiveEvpnType
	// and doesn't set defaults
	msg() *otg.BgpExtendedCommunityTransitiveEvpnType
	// setMsg unmarshals BgpExtendedCommunityTransitiveEvpnType from protobuf object *otg.BgpExtendedCommunityTransitiveEvpnType
	// and doesn't set defaults
	setMsg(*otg.BgpExtendedCommunityTransitiveEvpnType) BgpExtendedCommunityTransitiveEvpnType
	// provides marshal interface
	Marshal() marshalBgpExtendedCommunityTransitiveEvpnType
	// provides unmarshal interface
	Unmarshal() unMarshalBgpExtendedCommunityTransitiveEvpnType
	// validate validates BgpExtendedCommunityTransitiveEvpnType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpExtendedCommunityTransitiveEvpnType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpExtendedCommunityTransitiveEvpnTypeChoiceEnum, set in BgpExtendedCommunityTransitiveEvpnType
	Choice() BgpExtendedCommunityTransitiveEvpnTypeChoiceEnum
	// setChoice assigns BgpExtendedCommunityTransitiveEvpnTypeChoiceEnum provided by user to BgpExtendedCommunityTransitiveEvpnType
	setChoice(value BgpExtendedCommunityTransitiveEvpnTypeChoiceEnum) BgpExtendedCommunityTransitiveEvpnType
	// HasChoice checks if Choice has been set in BgpExtendedCommunityTransitiveEvpnType
	HasChoice() bool
	// RouterMacSubtype returns BgpExtendedCommunityTransitiveEvpnTypeRouterMac, set in BgpExtendedCommunityTransitiveEvpnType.
	// BgpExtendedCommunityTransitiveEvpnTypeRouterMac is the Router MAC EVPN Community is defined in RFC9135 and normally sent only for EVPN Type-2 Routes . It is sent with sub-type 0x03.
	RouterMacSubtype() BgpExtendedCommunityTransitiveEvpnTypeRouterMac
	// SetRouterMacSubtype assigns BgpExtendedCommunityTransitiveEvpnTypeRouterMac provided by user to BgpExtendedCommunityTransitiveEvpnType.
	// BgpExtendedCommunityTransitiveEvpnTypeRouterMac is the Router MAC EVPN Community is defined in RFC9135 and normally sent only for EVPN Type-2 Routes . It is sent with sub-type 0x03.
	SetRouterMacSubtype(value BgpExtendedCommunityTransitiveEvpnTypeRouterMac) BgpExtendedCommunityTransitiveEvpnType
	// HasRouterMacSubtype checks if RouterMacSubtype has been set in BgpExtendedCommunityTransitiveEvpnType
	HasRouterMacSubtype() bool
	setNil()
}

type BgpExtendedCommunityTransitiveEvpnTypeChoiceEnum string

// Enum of Choice on BgpExtendedCommunityTransitiveEvpnType
var BgpExtendedCommunityTransitiveEvpnTypeChoice = struct {
	ROUTER_MAC_SUBTYPE BgpExtendedCommunityTransitiveEvpnTypeChoiceEnum
}{
	ROUTER_MAC_SUBTYPE: BgpExtendedCommunityTransitiveEvpnTypeChoiceEnum("router_mac_subtype"),
}

func (obj *bgpExtendedCommunityTransitiveEvpnType) Choice() BgpExtendedCommunityTransitiveEvpnTypeChoiceEnum {
	return BgpExtendedCommunityTransitiveEvpnTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *bgpExtendedCommunityTransitiveEvpnType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpExtendedCommunityTransitiveEvpnType) setChoice(value BgpExtendedCommunityTransitiveEvpnTypeChoiceEnum) BgpExtendedCommunityTransitiveEvpnType {
	intValue, ok := otg.BgpExtendedCommunityTransitiveEvpnType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpExtendedCommunityTransitiveEvpnTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpExtendedCommunityTransitiveEvpnType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.RouterMacSubtype = nil
	obj.routerMacSubtypeHolder = nil

	if value == BgpExtendedCommunityTransitiveEvpnTypeChoice.ROUTER_MAC_SUBTYPE {
		obj.obj.RouterMacSubtype = NewBgpExtendedCommunityTransitiveEvpnTypeRouterMac().msg()
	}

	return obj
}

// description is TBD
// RouterMacSubtype returns a BgpExtendedCommunityTransitiveEvpnTypeRouterMac
func (obj *bgpExtendedCommunityTransitiveEvpnType) RouterMacSubtype() BgpExtendedCommunityTransitiveEvpnTypeRouterMac {
	if obj.obj.RouterMacSubtype == nil {
		obj.setChoice(BgpExtendedCommunityTransitiveEvpnTypeChoice.ROUTER_MAC_SUBTYPE)
	}
	if obj.routerMacSubtypeHolder == nil {
		obj.routerMacSubtypeHolder = &bgpExtendedCommunityTransitiveEvpnTypeRouterMac{obj: obj.obj.RouterMacSubtype}
	}
	return obj.routerMacSubtypeHolder
}

// description is TBD
// RouterMacSubtype returns a BgpExtendedCommunityTransitiveEvpnTypeRouterMac
func (obj *bgpExtendedCommunityTransitiveEvpnType) HasRouterMacSubtype() bool {
	return obj.obj.RouterMacSubtype != nil
}

// description is TBD
// SetRouterMacSubtype sets the BgpExtendedCommunityTransitiveEvpnTypeRouterMac value in the BgpExtendedCommunityTransitiveEvpnType object
func (obj *bgpExtendedCommunityTransitiveEvpnType) SetRouterMacSubtype(value BgpExtendedCommunityTransitiveEvpnTypeRouterMac) BgpExtendedCommunityTransitiveEvpnType {
	obj.setChoice(BgpExtendedCommunityTransitiveEvpnTypeChoice.ROUTER_MAC_SUBTYPE)
	obj.routerMacSubtypeHolder = nil
	obj.obj.RouterMacSubtype = value.msg()

	return obj
}

func (obj *bgpExtendedCommunityTransitiveEvpnType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RouterMacSubtype != nil {

		obj.RouterMacSubtype().validateObj(vObj, set_default)
	}

}

func (obj *bgpExtendedCommunityTransitiveEvpnType) setDefault() {
	var choices_set int = 0
	var choice BgpExtendedCommunityTransitiveEvpnTypeChoiceEnum

	if obj.obj.RouterMacSubtype != nil {
		choices_set += 1
		choice = BgpExtendedCommunityTransitiveEvpnTypeChoice.ROUTER_MAC_SUBTYPE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(BgpExtendedCommunityTransitiveEvpnTypeChoice.ROUTER_MAC_SUBTYPE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpExtendedCommunityTransitiveEvpnType")
			}
		} else {
			intVal := otg.BgpExtendedCommunityTransitiveEvpnType_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpExtendedCommunityTransitiveEvpnType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

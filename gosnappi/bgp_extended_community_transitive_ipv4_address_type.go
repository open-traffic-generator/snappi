package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpExtendedCommunityTransitiveIpv4AddressType *****
type bgpExtendedCommunityTransitiveIpv4AddressType struct {
	validation
	obj                      *otg.BgpExtendedCommunityTransitiveIpv4AddressType
	marshaller               marshalBgpExtendedCommunityTransitiveIpv4AddressType
	unMarshaller             unMarshalBgpExtendedCommunityTransitiveIpv4AddressType
	routeTargetSubtypeHolder BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	routeOriginSubtypeHolder BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
}

func NewBgpExtendedCommunityTransitiveIpv4AddressType() BgpExtendedCommunityTransitiveIpv4AddressType {
	obj := bgpExtendedCommunityTransitiveIpv4AddressType{obj: &otg.BgpExtendedCommunityTransitiveIpv4AddressType{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressType) msg() *otg.BgpExtendedCommunityTransitiveIpv4AddressType {
	return obj.obj
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressType) setMsg(msg *otg.BgpExtendedCommunityTransitiveIpv4AddressType) BgpExtendedCommunityTransitiveIpv4AddressType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpExtendedCommunityTransitiveIpv4AddressType struct {
	obj *bgpExtendedCommunityTransitiveIpv4AddressType
}

type marshalBgpExtendedCommunityTransitiveIpv4AddressType interface {
	// ToProto marshals BgpExtendedCommunityTransitiveIpv4AddressType to protobuf object *otg.BgpExtendedCommunityTransitiveIpv4AddressType
	ToProto() (*otg.BgpExtendedCommunityTransitiveIpv4AddressType, error)
	// ToPbText marshals BgpExtendedCommunityTransitiveIpv4AddressType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpExtendedCommunityTransitiveIpv4AddressType to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpExtendedCommunityTransitiveIpv4AddressType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpExtendedCommunityTransitiveIpv4AddressType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpExtendedCommunityTransitiveIpv4AddressType struct {
	obj *bgpExtendedCommunityTransitiveIpv4AddressType
}

type unMarshalBgpExtendedCommunityTransitiveIpv4AddressType interface {
	// FromProto unmarshals BgpExtendedCommunityTransitiveIpv4AddressType from protobuf object *otg.BgpExtendedCommunityTransitiveIpv4AddressType
	FromProto(msg *otg.BgpExtendedCommunityTransitiveIpv4AddressType) (BgpExtendedCommunityTransitiveIpv4AddressType, error)
	// FromPbText unmarshals BgpExtendedCommunityTransitiveIpv4AddressType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpExtendedCommunityTransitiveIpv4AddressType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpExtendedCommunityTransitiveIpv4AddressType from JSON text
	FromJson(value string) error
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressType) Marshal() marshalBgpExtendedCommunityTransitiveIpv4AddressType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpExtendedCommunityTransitiveIpv4AddressType{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressType) Unmarshal() unMarshalBgpExtendedCommunityTransitiveIpv4AddressType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpExtendedCommunityTransitiveIpv4AddressType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpExtendedCommunityTransitiveIpv4AddressType) ToProto() (*otg.BgpExtendedCommunityTransitiveIpv4AddressType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpExtendedCommunityTransitiveIpv4AddressType) FromProto(msg *otg.BgpExtendedCommunityTransitiveIpv4AddressType) (BgpExtendedCommunityTransitiveIpv4AddressType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpExtendedCommunityTransitiveIpv4AddressType) ToPbText() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveIpv4AddressType) FromPbText(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitiveIpv4AddressType) ToYaml() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveIpv4AddressType) FromYaml(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitiveIpv4AddressType) ToJsonRaw() (string, error) {
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

func (m *marshalbgpExtendedCommunityTransitiveIpv4AddressType) ToJson() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveIpv4AddressType) FromJson(value string) error {
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

func (obj *bgpExtendedCommunityTransitiveIpv4AddressType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressType) Clone() (BgpExtendedCommunityTransitiveIpv4AddressType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpExtendedCommunityTransitiveIpv4AddressType()
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

func (obj *bgpExtendedCommunityTransitiveIpv4AddressType) setNil() {
	obj.routeTargetSubtypeHolder = nil
	obj.routeOriginSubtypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpExtendedCommunityTransitiveIpv4AddressType is the Transitive IPv4 Address Specific Extended Community is sent as type 0x01.
type BgpExtendedCommunityTransitiveIpv4AddressType interface {
	Validation
	// msg marshals BgpExtendedCommunityTransitiveIpv4AddressType to protobuf object *otg.BgpExtendedCommunityTransitiveIpv4AddressType
	// and doesn't set defaults
	msg() *otg.BgpExtendedCommunityTransitiveIpv4AddressType
	// setMsg unmarshals BgpExtendedCommunityTransitiveIpv4AddressType from protobuf object *otg.BgpExtendedCommunityTransitiveIpv4AddressType
	// and doesn't set defaults
	setMsg(*otg.BgpExtendedCommunityTransitiveIpv4AddressType) BgpExtendedCommunityTransitiveIpv4AddressType
	// provides marshal interface
	Marshal() marshalBgpExtendedCommunityTransitiveIpv4AddressType
	// provides unmarshal interface
	Unmarshal() unMarshalBgpExtendedCommunityTransitiveIpv4AddressType
	// validate validates BgpExtendedCommunityTransitiveIpv4AddressType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpExtendedCommunityTransitiveIpv4AddressType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum, set in BgpExtendedCommunityTransitiveIpv4AddressType
	Choice() BgpExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum
	// setChoice assigns BgpExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum provided by user to BgpExtendedCommunityTransitiveIpv4AddressType
	setChoice(value BgpExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum) BgpExtendedCommunityTransitiveIpv4AddressType
	// HasChoice checks if Choice has been set in BgpExtendedCommunityTransitiveIpv4AddressType
	HasChoice() bool
	// RouteTargetSubtype returns BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget, set in BgpExtendedCommunityTransitiveIpv4AddressType.
	// BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget is the Route Target Community identifies one or more routers that may receive a set of routes (that carry this Community) carried by BGP.  It is sent with sub-type as 0x02.
	RouteTargetSubtype() BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	// SetRouteTargetSubtype assigns BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget provided by user to BgpExtendedCommunityTransitiveIpv4AddressType.
	// BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget is the Route Target Community identifies one or more routers that may receive a set of routes (that carry this Community) carried by BGP.  It is sent with sub-type as 0x02.
	SetRouteTargetSubtype(value BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) BgpExtendedCommunityTransitiveIpv4AddressType
	// HasRouteTargetSubtype checks if RouteTargetSubtype has been set in BgpExtendedCommunityTransitiveIpv4AddressType
	HasRouteTargetSubtype() bool
	// RouteOriginSubtype returns BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin, set in BgpExtendedCommunityTransitiveIpv4AddressType.
	// BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin is the Route Origin Community identifies one or more routers that inject a set of routes (that carry this Community) into BGP It is sent with sub-type as 0x03.
	RouteOriginSubtype() BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	// SetRouteOriginSubtype assigns BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin provided by user to BgpExtendedCommunityTransitiveIpv4AddressType.
	// BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin is the Route Origin Community identifies one or more routers that inject a set of routes (that carry this Community) into BGP It is sent with sub-type as 0x03.
	SetRouteOriginSubtype(value BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) BgpExtendedCommunityTransitiveIpv4AddressType
	// HasRouteOriginSubtype checks if RouteOriginSubtype has been set in BgpExtendedCommunityTransitiveIpv4AddressType
	HasRouteOriginSubtype() bool
	setNil()
}

type BgpExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum string

// Enum of Choice on BgpExtendedCommunityTransitiveIpv4AddressType
var BgpExtendedCommunityTransitiveIpv4AddressTypeChoice = struct {
	ROUTE_TARGET_SUBTYPE BgpExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum
	ROUTE_ORIGIN_SUBTYPE BgpExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum
}{
	ROUTE_TARGET_SUBTYPE: BgpExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum("route_target_subtype"),
	ROUTE_ORIGIN_SUBTYPE: BgpExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum("route_origin_subtype"),
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressType) Choice() BgpExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum {
	return BgpExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *bgpExtendedCommunityTransitiveIpv4AddressType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressType) setChoice(value BgpExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum) BgpExtendedCommunityTransitiveIpv4AddressType {
	intValue, ok := otg.BgpExtendedCommunityTransitiveIpv4AddressType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpExtendedCommunityTransitiveIpv4AddressType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.RouteOriginSubtype = nil
	obj.routeOriginSubtypeHolder = nil
	obj.obj.RouteTargetSubtype = nil
	obj.routeTargetSubtypeHolder = nil

	if value == BgpExtendedCommunityTransitiveIpv4AddressTypeChoice.ROUTE_TARGET_SUBTYPE {
		obj.obj.RouteTargetSubtype = NewBgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget().msg()
	}

	if value == BgpExtendedCommunityTransitiveIpv4AddressTypeChoice.ROUTE_ORIGIN_SUBTYPE {
		obj.obj.RouteOriginSubtype = NewBgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin().msg()
	}

	return obj
}

// description is TBD
// RouteTargetSubtype returns a BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
func (obj *bgpExtendedCommunityTransitiveIpv4AddressType) RouteTargetSubtype() BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget {
	if obj.obj.RouteTargetSubtype == nil {
		obj.setChoice(BgpExtendedCommunityTransitiveIpv4AddressTypeChoice.ROUTE_TARGET_SUBTYPE)
	}
	if obj.routeTargetSubtypeHolder == nil {
		obj.routeTargetSubtypeHolder = &bgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget{obj: obj.obj.RouteTargetSubtype}
	}
	return obj.routeTargetSubtypeHolder
}

// description is TBD
// RouteTargetSubtype returns a BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
func (obj *bgpExtendedCommunityTransitiveIpv4AddressType) HasRouteTargetSubtype() bool {
	return obj.obj.RouteTargetSubtype != nil
}

// description is TBD
// SetRouteTargetSubtype sets the BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget value in the BgpExtendedCommunityTransitiveIpv4AddressType object
func (obj *bgpExtendedCommunityTransitiveIpv4AddressType) SetRouteTargetSubtype(value BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) BgpExtendedCommunityTransitiveIpv4AddressType {
	obj.setChoice(BgpExtendedCommunityTransitiveIpv4AddressTypeChoice.ROUTE_TARGET_SUBTYPE)
	obj.routeTargetSubtypeHolder = nil
	obj.obj.RouteTargetSubtype = value.msg()

	return obj
}

// description is TBD
// RouteOriginSubtype returns a BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
func (obj *bgpExtendedCommunityTransitiveIpv4AddressType) RouteOriginSubtype() BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin {
	if obj.obj.RouteOriginSubtype == nil {
		obj.setChoice(BgpExtendedCommunityTransitiveIpv4AddressTypeChoice.ROUTE_ORIGIN_SUBTYPE)
	}
	if obj.routeOriginSubtypeHolder == nil {
		obj.routeOriginSubtypeHolder = &bgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin{obj: obj.obj.RouteOriginSubtype}
	}
	return obj.routeOriginSubtypeHolder
}

// description is TBD
// RouteOriginSubtype returns a BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
func (obj *bgpExtendedCommunityTransitiveIpv4AddressType) HasRouteOriginSubtype() bool {
	return obj.obj.RouteOriginSubtype != nil
}

// description is TBD
// SetRouteOriginSubtype sets the BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin value in the BgpExtendedCommunityTransitiveIpv4AddressType object
func (obj *bgpExtendedCommunityTransitiveIpv4AddressType) SetRouteOriginSubtype(value BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) BgpExtendedCommunityTransitiveIpv4AddressType {
	obj.setChoice(BgpExtendedCommunityTransitiveIpv4AddressTypeChoice.ROUTE_ORIGIN_SUBTYPE)
	obj.routeOriginSubtypeHolder = nil
	obj.obj.RouteOriginSubtype = value.msg()

	return obj
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RouteTargetSubtype != nil {

		obj.RouteTargetSubtype().validateObj(vObj, set_default)
	}

	if obj.obj.RouteOriginSubtype != nil {

		obj.RouteOriginSubtype().validateObj(vObj, set_default)
	}

}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressType) setDefault() {
	var choices_set int = 0
	var choice BgpExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum

	if obj.obj.RouteTargetSubtype != nil {
		choices_set += 1
		choice = BgpExtendedCommunityTransitiveIpv4AddressTypeChoice.ROUTE_TARGET_SUBTYPE
	}

	if obj.obj.RouteOriginSubtype != nil {
		choices_set += 1
		choice = BgpExtendedCommunityTransitiveIpv4AddressTypeChoice.ROUTE_ORIGIN_SUBTYPE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(BgpExtendedCommunityTransitiveIpv4AddressTypeChoice.ROUTE_TARGET_SUBTYPE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpExtendedCommunityTransitiveIpv4AddressType")
			}
		} else {
			intVal := otg.BgpExtendedCommunityTransitiveIpv4AddressType_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpExtendedCommunityTransitiveIpv4AddressType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

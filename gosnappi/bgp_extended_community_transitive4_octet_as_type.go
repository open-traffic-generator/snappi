package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpExtendedCommunityTransitive4OctetAsType *****
type bgpExtendedCommunityTransitive4OctetAsType struct {
	validation
	obj                      *otg.BgpExtendedCommunityTransitive4OctetAsType
	marshaller               marshalBgpExtendedCommunityTransitive4OctetAsType
	unMarshaller             unMarshalBgpExtendedCommunityTransitive4OctetAsType
	routeTargetSubtypeHolder BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
	routeOriginSubtypeHolder BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
}

func NewBgpExtendedCommunityTransitive4OctetAsType() BgpExtendedCommunityTransitive4OctetAsType {
	obj := bgpExtendedCommunityTransitive4OctetAsType{obj: &otg.BgpExtendedCommunityTransitive4OctetAsType{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpExtendedCommunityTransitive4OctetAsType) msg() *otg.BgpExtendedCommunityTransitive4OctetAsType {
	return obj.obj
}

func (obj *bgpExtendedCommunityTransitive4OctetAsType) setMsg(msg *otg.BgpExtendedCommunityTransitive4OctetAsType) BgpExtendedCommunityTransitive4OctetAsType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpExtendedCommunityTransitive4OctetAsType struct {
	obj *bgpExtendedCommunityTransitive4OctetAsType
}

type marshalBgpExtendedCommunityTransitive4OctetAsType interface {
	// ToProto marshals BgpExtendedCommunityTransitive4OctetAsType to protobuf object *otg.BgpExtendedCommunityTransitive4OctetAsType
	ToProto() (*otg.BgpExtendedCommunityTransitive4OctetAsType, error)
	// ToPbText marshals BgpExtendedCommunityTransitive4OctetAsType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpExtendedCommunityTransitive4OctetAsType to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpExtendedCommunityTransitive4OctetAsType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpExtendedCommunityTransitive4OctetAsType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpExtendedCommunityTransitive4OctetAsType struct {
	obj *bgpExtendedCommunityTransitive4OctetAsType
}

type unMarshalBgpExtendedCommunityTransitive4OctetAsType interface {
	// FromProto unmarshals BgpExtendedCommunityTransitive4OctetAsType from protobuf object *otg.BgpExtendedCommunityTransitive4OctetAsType
	FromProto(msg *otg.BgpExtendedCommunityTransitive4OctetAsType) (BgpExtendedCommunityTransitive4OctetAsType, error)
	// FromPbText unmarshals BgpExtendedCommunityTransitive4OctetAsType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpExtendedCommunityTransitive4OctetAsType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpExtendedCommunityTransitive4OctetAsType from JSON text
	FromJson(value string) error
}

func (obj *bgpExtendedCommunityTransitive4OctetAsType) Marshal() marshalBgpExtendedCommunityTransitive4OctetAsType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpExtendedCommunityTransitive4OctetAsType{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpExtendedCommunityTransitive4OctetAsType) Unmarshal() unMarshalBgpExtendedCommunityTransitive4OctetAsType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpExtendedCommunityTransitive4OctetAsType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpExtendedCommunityTransitive4OctetAsType) ToProto() (*otg.BgpExtendedCommunityTransitive4OctetAsType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpExtendedCommunityTransitive4OctetAsType) FromProto(msg *otg.BgpExtendedCommunityTransitive4OctetAsType) (BgpExtendedCommunityTransitive4OctetAsType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpExtendedCommunityTransitive4OctetAsType) ToPbText() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitive4OctetAsType) FromPbText(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitive4OctetAsType) ToYaml() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitive4OctetAsType) FromYaml(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitive4OctetAsType) ToJsonRaw() (string, error) {
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

func (m *marshalbgpExtendedCommunityTransitive4OctetAsType) ToJson() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitive4OctetAsType) FromJson(value string) error {
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

func (obj *bgpExtendedCommunityTransitive4OctetAsType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitive4OctetAsType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitive4OctetAsType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpExtendedCommunityTransitive4OctetAsType) Clone() (BgpExtendedCommunityTransitive4OctetAsType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpExtendedCommunityTransitive4OctetAsType()
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

func (obj *bgpExtendedCommunityTransitive4OctetAsType) setNil() {
	obj.routeTargetSubtypeHolder = nil
	obj.routeOriginSubtypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpExtendedCommunityTransitive4OctetAsType is the Transitive Four-Octet AS-Specific Extended Community is sent as type 0x02. It is defined in RFC 5668.
type BgpExtendedCommunityTransitive4OctetAsType interface {
	Validation
	// msg marshals BgpExtendedCommunityTransitive4OctetAsType to protobuf object *otg.BgpExtendedCommunityTransitive4OctetAsType
	// and doesn't set defaults
	msg() *otg.BgpExtendedCommunityTransitive4OctetAsType
	// setMsg unmarshals BgpExtendedCommunityTransitive4OctetAsType from protobuf object *otg.BgpExtendedCommunityTransitive4OctetAsType
	// and doesn't set defaults
	setMsg(*otg.BgpExtendedCommunityTransitive4OctetAsType) BgpExtendedCommunityTransitive4OctetAsType
	// provides marshal interface
	Marshal() marshalBgpExtendedCommunityTransitive4OctetAsType
	// provides unmarshal interface
	Unmarshal() unMarshalBgpExtendedCommunityTransitive4OctetAsType
	// validate validates BgpExtendedCommunityTransitive4OctetAsType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpExtendedCommunityTransitive4OctetAsType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpExtendedCommunityTransitive4OctetAsTypeChoiceEnum, set in BgpExtendedCommunityTransitive4OctetAsType
	Choice() BgpExtendedCommunityTransitive4OctetAsTypeChoiceEnum
	// setChoice assigns BgpExtendedCommunityTransitive4OctetAsTypeChoiceEnum provided by user to BgpExtendedCommunityTransitive4OctetAsType
	setChoice(value BgpExtendedCommunityTransitive4OctetAsTypeChoiceEnum) BgpExtendedCommunityTransitive4OctetAsType
	// HasChoice checks if Choice has been set in BgpExtendedCommunityTransitive4OctetAsType
	HasChoice() bool
	// RouteTargetSubtype returns BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget, set in BgpExtendedCommunityTransitive4OctetAsType.
	// BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget is the Route Target Community identifies one or more routers that may receive a set of routes (that carry this Community) carried by BGP.  It is sent with sub-type as 0x02
	RouteTargetSubtype() BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
	// SetRouteTargetSubtype assigns BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget provided by user to BgpExtendedCommunityTransitive4OctetAsType.
	// BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget is the Route Target Community identifies one or more routers that may receive a set of routes (that carry this Community) carried by BGP.  It is sent with sub-type as 0x02
	SetRouteTargetSubtype(value BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) BgpExtendedCommunityTransitive4OctetAsType
	// HasRouteTargetSubtype checks if RouteTargetSubtype has been set in BgpExtendedCommunityTransitive4OctetAsType
	HasRouteTargetSubtype() bool
	// RouteOriginSubtype returns BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin, set in BgpExtendedCommunityTransitive4OctetAsType.
	// BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin is the Route Origin Community identifies one or more routers that inject a set of routes (that carry this Community) into BGP. It is sent with sub-type as 0x03.
	RouteOriginSubtype() BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	// SetRouteOriginSubtype assigns BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin provided by user to BgpExtendedCommunityTransitive4OctetAsType.
	// BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin is the Route Origin Community identifies one or more routers that inject a set of routes (that carry this Community) into BGP. It is sent with sub-type as 0x03.
	SetRouteOriginSubtype(value BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) BgpExtendedCommunityTransitive4OctetAsType
	// HasRouteOriginSubtype checks if RouteOriginSubtype has been set in BgpExtendedCommunityTransitive4OctetAsType
	HasRouteOriginSubtype() bool
	setNil()
}

type BgpExtendedCommunityTransitive4OctetAsTypeChoiceEnum string

// Enum of Choice on BgpExtendedCommunityTransitive4OctetAsType
var BgpExtendedCommunityTransitive4OctetAsTypeChoice = struct {
	ROUTE_TARGET_SUBTYPE BgpExtendedCommunityTransitive4OctetAsTypeChoiceEnum
	ROUTE_ORIGIN_SUBTYPE BgpExtendedCommunityTransitive4OctetAsTypeChoiceEnum
}{
	ROUTE_TARGET_SUBTYPE: BgpExtendedCommunityTransitive4OctetAsTypeChoiceEnum("route_target_subtype"),
	ROUTE_ORIGIN_SUBTYPE: BgpExtendedCommunityTransitive4OctetAsTypeChoiceEnum("route_origin_subtype"),
}

func (obj *bgpExtendedCommunityTransitive4OctetAsType) Choice() BgpExtendedCommunityTransitive4OctetAsTypeChoiceEnum {
	return BgpExtendedCommunityTransitive4OctetAsTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *bgpExtendedCommunityTransitive4OctetAsType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpExtendedCommunityTransitive4OctetAsType) setChoice(value BgpExtendedCommunityTransitive4OctetAsTypeChoiceEnum) BgpExtendedCommunityTransitive4OctetAsType {
	intValue, ok := otg.BgpExtendedCommunityTransitive4OctetAsType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpExtendedCommunityTransitive4OctetAsTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpExtendedCommunityTransitive4OctetAsType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.RouteOriginSubtype = nil
	obj.routeOriginSubtypeHolder = nil
	obj.obj.RouteTargetSubtype = nil
	obj.routeTargetSubtypeHolder = nil

	if value == BgpExtendedCommunityTransitive4OctetAsTypeChoice.ROUTE_TARGET_SUBTYPE {
		obj.obj.RouteTargetSubtype = NewBgpExtendedCommunityTransitive4OctetAsTypeRouteTarget().msg()
	}

	if value == BgpExtendedCommunityTransitive4OctetAsTypeChoice.ROUTE_ORIGIN_SUBTYPE {
		obj.obj.RouteOriginSubtype = NewBgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin().msg()
	}

	return obj
}

// description is TBD
// RouteTargetSubtype returns a BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
func (obj *bgpExtendedCommunityTransitive4OctetAsType) RouteTargetSubtype() BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget {
	if obj.obj.RouteTargetSubtype == nil {
		obj.setChoice(BgpExtendedCommunityTransitive4OctetAsTypeChoice.ROUTE_TARGET_SUBTYPE)
	}
	if obj.routeTargetSubtypeHolder == nil {
		obj.routeTargetSubtypeHolder = &bgpExtendedCommunityTransitive4OctetAsTypeRouteTarget{obj: obj.obj.RouteTargetSubtype}
	}
	return obj.routeTargetSubtypeHolder
}

// description is TBD
// RouteTargetSubtype returns a BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
func (obj *bgpExtendedCommunityTransitive4OctetAsType) HasRouteTargetSubtype() bool {
	return obj.obj.RouteTargetSubtype != nil
}

// description is TBD
// SetRouteTargetSubtype sets the BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget value in the BgpExtendedCommunityTransitive4OctetAsType object
func (obj *bgpExtendedCommunityTransitive4OctetAsType) SetRouteTargetSubtype(value BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) BgpExtendedCommunityTransitive4OctetAsType {
	obj.setChoice(BgpExtendedCommunityTransitive4OctetAsTypeChoice.ROUTE_TARGET_SUBTYPE)
	obj.routeTargetSubtypeHolder = nil
	obj.obj.RouteTargetSubtype = value.msg()

	return obj
}

// description is TBD
// RouteOriginSubtype returns a BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
func (obj *bgpExtendedCommunityTransitive4OctetAsType) RouteOriginSubtype() BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin {
	if obj.obj.RouteOriginSubtype == nil {
		obj.setChoice(BgpExtendedCommunityTransitive4OctetAsTypeChoice.ROUTE_ORIGIN_SUBTYPE)
	}
	if obj.routeOriginSubtypeHolder == nil {
		obj.routeOriginSubtypeHolder = &bgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin{obj: obj.obj.RouteOriginSubtype}
	}
	return obj.routeOriginSubtypeHolder
}

// description is TBD
// RouteOriginSubtype returns a BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
func (obj *bgpExtendedCommunityTransitive4OctetAsType) HasRouteOriginSubtype() bool {
	return obj.obj.RouteOriginSubtype != nil
}

// description is TBD
// SetRouteOriginSubtype sets the BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin value in the BgpExtendedCommunityTransitive4OctetAsType object
func (obj *bgpExtendedCommunityTransitive4OctetAsType) SetRouteOriginSubtype(value BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) BgpExtendedCommunityTransitive4OctetAsType {
	obj.setChoice(BgpExtendedCommunityTransitive4OctetAsTypeChoice.ROUTE_ORIGIN_SUBTYPE)
	obj.routeOriginSubtypeHolder = nil
	obj.obj.RouteOriginSubtype = value.msg()

	return obj
}

func (obj *bgpExtendedCommunityTransitive4OctetAsType) validateObj(vObj *validation, set_default bool) {
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

func (obj *bgpExtendedCommunityTransitive4OctetAsType) setDefault() {
	var choices_set int = 0
	var choice BgpExtendedCommunityTransitive4OctetAsTypeChoiceEnum

	if obj.obj.RouteTargetSubtype != nil {
		choices_set += 1
		choice = BgpExtendedCommunityTransitive4OctetAsTypeChoice.ROUTE_TARGET_SUBTYPE
	}

	if obj.obj.RouteOriginSubtype != nil {
		choices_set += 1
		choice = BgpExtendedCommunityTransitive4OctetAsTypeChoice.ROUTE_ORIGIN_SUBTYPE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(BgpExtendedCommunityTransitive4OctetAsTypeChoice.ROUTE_TARGET_SUBTYPE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpExtendedCommunityTransitive4OctetAsType")
			}
		} else {
			intVal := otg.BgpExtendedCommunityTransitive4OctetAsType_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpExtendedCommunityTransitive4OctetAsType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

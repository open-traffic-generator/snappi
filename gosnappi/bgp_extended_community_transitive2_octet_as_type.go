package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpExtendedCommunityTransitive2OctetAsType *****
type bgpExtendedCommunityTransitive2OctetAsType struct {
	validation
	obj                      *otg.BgpExtendedCommunityTransitive2OctetAsType
	marshaller               marshalBgpExtendedCommunityTransitive2OctetAsType
	unMarshaller             unMarshalBgpExtendedCommunityTransitive2OctetAsType
	routeTargetSubtypeHolder BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
	routeOriginSubtypeHolder BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
}

func NewBgpExtendedCommunityTransitive2OctetAsType() BgpExtendedCommunityTransitive2OctetAsType {
	obj := bgpExtendedCommunityTransitive2OctetAsType{obj: &otg.BgpExtendedCommunityTransitive2OctetAsType{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpExtendedCommunityTransitive2OctetAsType) msg() *otg.BgpExtendedCommunityTransitive2OctetAsType {
	return obj.obj
}

func (obj *bgpExtendedCommunityTransitive2OctetAsType) setMsg(msg *otg.BgpExtendedCommunityTransitive2OctetAsType) BgpExtendedCommunityTransitive2OctetAsType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpExtendedCommunityTransitive2OctetAsType struct {
	obj *bgpExtendedCommunityTransitive2OctetAsType
}

type marshalBgpExtendedCommunityTransitive2OctetAsType interface {
	// ToProto marshals BgpExtendedCommunityTransitive2OctetAsType to protobuf object *otg.BgpExtendedCommunityTransitive2OctetAsType
	ToProto() (*otg.BgpExtendedCommunityTransitive2OctetAsType, error)
	// ToPbText marshals BgpExtendedCommunityTransitive2OctetAsType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpExtendedCommunityTransitive2OctetAsType to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpExtendedCommunityTransitive2OctetAsType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpExtendedCommunityTransitive2OctetAsType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpExtendedCommunityTransitive2OctetAsType struct {
	obj *bgpExtendedCommunityTransitive2OctetAsType
}

type unMarshalBgpExtendedCommunityTransitive2OctetAsType interface {
	// FromProto unmarshals BgpExtendedCommunityTransitive2OctetAsType from protobuf object *otg.BgpExtendedCommunityTransitive2OctetAsType
	FromProto(msg *otg.BgpExtendedCommunityTransitive2OctetAsType) (BgpExtendedCommunityTransitive2OctetAsType, error)
	// FromPbText unmarshals BgpExtendedCommunityTransitive2OctetAsType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpExtendedCommunityTransitive2OctetAsType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpExtendedCommunityTransitive2OctetAsType from JSON text
	FromJson(value string) error
}

func (obj *bgpExtendedCommunityTransitive2OctetAsType) Marshal() marshalBgpExtendedCommunityTransitive2OctetAsType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpExtendedCommunityTransitive2OctetAsType{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpExtendedCommunityTransitive2OctetAsType) Unmarshal() unMarshalBgpExtendedCommunityTransitive2OctetAsType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpExtendedCommunityTransitive2OctetAsType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpExtendedCommunityTransitive2OctetAsType) ToProto() (*otg.BgpExtendedCommunityTransitive2OctetAsType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpExtendedCommunityTransitive2OctetAsType) FromProto(msg *otg.BgpExtendedCommunityTransitive2OctetAsType) (BgpExtendedCommunityTransitive2OctetAsType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpExtendedCommunityTransitive2OctetAsType) ToPbText() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitive2OctetAsType) FromPbText(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitive2OctetAsType) ToYaml() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitive2OctetAsType) FromYaml(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitive2OctetAsType) ToJsonRaw() (string, error) {
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

func (m *marshalbgpExtendedCommunityTransitive2OctetAsType) ToJson() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitive2OctetAsType) FromJson(value string) error {
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

func (obj *bgpExtendedCommunityTransitive2OctetAsType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitive2OctetAsType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitive2OctetAsType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpExtendedCommunityTransitive2OctetAsType) Clone() (BgpExtendedCommunityTransitive2OctetAsType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpExtendedCommunityTransitive2OctetAsType()
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

func (obj *bgpExtendedCommunityTransitive2OctetAsType) setNil() {
	obj.routeTargetSubtypeHolder = nil
	obj.routeOriginSubtypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpExtendedCommunityTransitive2OctetAsType is the Transitive Two-Octet AS-Specific Extended Community is sent as type 0x00 .
type BgpExtendedCommunityTransitive2OctetAsType interface {
	Validation
	// msg marshals BgpExtendedCommunityTransitive2OctetAsType to protobuf object *otg.BgpExtendedCommunityTransitive2OctetAsType
	// and doesn't set defaults
	msg() *otg.BgpExtendedCommunityTransitive2OctetAsType
	// setMsg unmarshals BgpExtendedCommunityTransitive2OctetAsType from protobuf object *otg.BgpExtendedCommunityTransitive2OctetAsType
	// and doesn't set defaults
	setMsg(*otg.BgpExtendedCommunityTransitive2OctetAsType) BgpExtendedCommunityTransitive2OctetAsType
	// provides marshal interface
	Marshal() marshalBgpExtendedCommunityTransitive2OctetAsType
	// provides unmarshal interface
	Unmarshal() unMarshalBgpExtendedCommunityTransitive2OctetAsType
	// validate validates BgpExtendedCommunityTransitive2OctetAsType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpExtendedCommunityTransitive2OctetAsType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpExtendedCommunityTransitive2OctetAsTypeChoiceEnum, set in BgpExtendedCommunityTransitive2OctetAsType
	Choice() BgpExtendedCommunityTransitive2OctetAsTypeChoiceEnum
	// setChoice assigns BgpExtendedCommunityTransitive2OctetAsTypeChoiceEnum provided by user to BgpExtendedCommunityTransitive2OctetAsType
	setChoice(value BgpExtendedCommunityTransitive2OctetAsTypeChoiceEnum) BgpExtendedCommunityTransitive2OctetAsType
	// HasChoice checks if Choice has been set in BgpExtendedCommunityTransitive2OctetAsType
	HasChoice() bool
	// RouteTargetSubtype returns BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget, set in BgpExtendedCommunityTransitive2OctetAsType.
	// BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget is the Route Target Community identifies one or more routers that may receive a set of routes (that carry this Community) carried by BGP.  It is sent with sub-type as 0x02.
	RouteTargetSubtype() BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
	// SetRouteTargetSubtype assigns BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget provided by user to BgpExtendedCommunityTransitive2OctetAsType.
	// BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget is the Route Target Community identifies one or more routers that may receive a set of routes (that carry this Community) carried by BGP.  It is sent with sub-type as 0x02.
	SetRouteTargetSubtype(value BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) BgpExtendedCommunityTransitive2OctetAsType
	// HasRouteTargetSubtype checks if RouteTargetSubtype has been set in BgpExtendedCommunityTransitive2OctetAsType
	HasRouteTargetSubtype() bool
	// RouteOriginSubtype returns BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin, set in BgpExtendedCommunityTransitive2OctetAsType.
	// BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin is the Route Origin Community identifies one or more routers that inject a set of routes (that carry this Community) into BGP. It is sent with sub-type as 0x03 .
	RouteOriginSubtype() BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	// SetRouteOriginSubtype assigns BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin provided by user to BgpExtendedCommunityTransitive2OctetAsType.
	// BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin is the Route Origin Community identifies one or more routers that inject a set of routes (that carry this Community) into BGP. It is sent with sub-type as 0x03 .
	SetRouteOriginSubtype(value BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) BgpExtendedCommunityTransitive2OctetAsType
	// HasRouteOriginSubtype checks if RouteOriginSubtype has been set in BgpExtendedCommunityTransitive2OctetAsType
	HasRouteOriginSubtype() bool
	setNil()
}

type BgpExtendedCommunityTransitive2OctetAsTypeChoiceEnum string

// Enum of Choice on BgpExtendedCommunityTransitive2OctetAsType
var BgpExtendedCommunityTransitive2OctetAsTypeChoice = struct {
	ROUTE_TARGET_SUBTYPE BgpExtendedCommunityTransitive2OctetAsTypeChoiceEnum
	ROUTE_ORIGIN_SUBTYPE BgpExtendedCommunityTransitive2OctetAsTypeChoiceEnum
}{
	ROUTE_TARGET_SUBTYPE: BgpExtendedCommunityTransitive2OctetAsTypeChoiceEnum("route_target_subtype"),
	ROUTE_ORIGIN_SUBTYPE: BgpExtendedCommunityTransitive2OctetAsTypeChoiceEnum("route_origin_subtype"),
}

func (obj *bgpExtendedCommunityTransitive2OctetAsType) Choice() BgpExtendedCommunityTransitive2OctetAsTypeChoiceEnum {
	return BgpExtendedCommunityTransitive2OctetAsTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *bgpExtendedCommunityTransitive2OctetAsType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpExtendedCommunityTransitive2OctetAsType) setChoice(value BgpExtendedCommunityTransitive2OctetAsTypeChoiceEnum) BgpExtendedCommunityTransitive2OctetAsType {
	intValue, ok := otg.BgpExtendedCommunityTransitive2OctetAsType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpExtendedCommunityTransitive2OctetAsTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpExtendedCommunityTransitive2OctetAsType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.RouteOriginSubtype = nil
	obj.routeOriginSubtypeHolder = nil
	obj.obj.RouteTargetSubtype = nil
	obj.routeTargetSubtypeHolder = nil

	if value == BgpExtendedCommunityTransitive2OctetAsTypeChoice.ROUTE_TARGET_SUBTYPE {
		obj.obj.RouteTargetSubtype = NewBgpExtendedCommunityTransitive2OctetAsTypeRouteTarget().msg()
	}

	if value == BgpExtendedCommunityTransitive2OctetAsTypeChoice.ROUTE_ORIGIN_SUBTYPE {
		obj.obj.RouteOriginSubtype = NewBgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin().msg()
	}

	return obj
}

// description is TBD
// RouteTargetSubtype returns a BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
func (obj *bgpExtendedCommunityTransitive2OctetAsType) RouteTargetSubtype() BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget {
	if obj.obj.RouteTargetSubtype == nil {
		obj.setChoice(BgpExtendedCommunityTransitive2OctetAsTypeChoice.ROUTE_TARGET_SUBTYPE)
	}
	if obj.routeTargetSubtypeHolder == nil {
		obj.routeTargetSubtypeHolder = &bgpExtendedCommunityTransitive2OctetAsTypeRouteTarget{obj: obj.obj.RouteTargetSubtype}
	}
	return obj.routeTargetSubtypeHolder
}

// description is TBD
// RouteTargetSubtype returns a BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
func (obj *bgpExtendedCommunityTransitive2OctetAsType) HasRouteTargetSubtype() bool {
	return obj.obj.RouteTargetSubtype != nil
}

// description is TBD
// SetRouteTargetSubtype sets the BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget value in the BgpExtendedCommunityTransitive2OctetAsType object
func (obj *bgpExtendedCommunityTransitive2OctetAsType) SetRouteTargetSubtype(value BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) BgpExtendedCommunityTransitive2OctetAsType {
	obj.setChoice(BgpExtendedCommunityTransitive2OctetAsTypeChoice.ROUTE_TARGET_SUBTYPE)
	obj.routeTargetSubtypeHolder = nil
	obj.obj.RouteTargetSubtype = value.msg()

	return obj
}

// description is TBD
// RouteOriginSubtype returns a BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
func (obj *bgpExtendedCommunityTransitive2OctetAsType) RouteOriginSubtype() BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin {
	if obj.obj.RouteOriginSubtype == nil {
		obj.setChoice(BgpExtendedCommunityTransitive2OctetAsTypeChoice.ROUTE_ORIGIN_SUBTYPE)
	}
	if obj.routeOriginSubtypeHolder == nil {
		obj.routeOriginSubtypeHolder = &bgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin{obj: obj.obj.RouteOriginSubtype}
	}
	return obj.routeOriginSubtypeHolder
}

// description is TBD
// RouteOriginSubtype returns a BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
func (obj *bgpExtendedCommunityTransitive2OctetAsType) HasRouteOriginSubtype() bool {
	return obj.obj.RouteOriginSubtype != nil
}

// description is TBD
// SetRouteOriginSubtype sets the BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin value in the BgpExtendedCommunityTransitive2OctetAsType object
func (obj *bgpExtendedCommunityTransitive2OctetAsType) SetRouteOriginSubtype(value BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) BgpExtendedCommunityTransitive2OctetAsType {
	obj.setChoice(BgpExtendedCommunityTransitive2OctetAsTypeChoice.ROUTE_ORIGIN_SUBTYPE)
	obj.routeOriginSubtypeHolder = nil
	obj.obj.RouteOriginSubtype = value.msg()

	return obj
}

func (obj *bgpExtendedCommunityTransitive2OctetAsType) validateObj(vObj *validation, set_default bool) {
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

func (obj *bgpExtendedCommunityTransitive2OctetAsType) setDefault() {
	var choices_set int = 0
	var choice BgpExtendedCommunityTransitive2OctetAsTypeChoiceEnum

	if obj.obj.RouteTargetSubtype != nil {
		choices_set += 1
		choice = BgpExtendedCommunityTransitive2OctetAsTypeChoice.ROUTE_TARGET_SUBTYPE
	}

	if obj.obj.RouteOriginSubtype != nil {
		choices_set += 1
		choice = BgpExtendedCommunityTransitive2OctetAsTypeChoice.ROUTE_ORIGIN_SUBTYPE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(BgpExtendedCommunityTransitive2OctetAsTypeChoice.ROUTE_TARGET_SUBTYPE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpExtendedCommunityTransitive2OctetAsType")
			}
		} else {
			intVal := otg.BgpExtendedCommunityTransitive2OctetAsType_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpExtendedCommunityTransitive2OctetAsType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

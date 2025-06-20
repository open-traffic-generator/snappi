package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ResultExtendedCommunityTransitiveIpv4AddressType *****
type resultExtendedCommunityTransitiveIpv4AddressType struct {
	validation
	obj                      *otg.ResultExtendedCommunityTransitiveIpv4AddressType
	marshaller               marshalResultExtendedCommunityTransitiveIpv4AddressType
	unMarshaller             unMarshalResultExtendedCommunityTransitiveIpv4AddressType
	routeTargetSubtypeHolder ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	routeOriginSubtypeHolder ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
}

func NewResultExtendedCommunityTransitiveIpv4AddressType() ResultExtendedCommunityTransitiveIpv4AddressType {
	obj := resultExtendedCommunityTransitiveIpv4AddressType{obj: &otg.ResultExtendedCommunityTransitiveIpv4AddressType{}}
	obj.setDefault()
	return &obj
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressType) msg() *otg.ResultExtendedCommunityTransitiveIpv4AddressType {
	return obj.obj
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressType) setMsg(msg *otg.ResultExtendedCommunityTransitiveIpv4AddressType) ResultExtendedCommunityTransitiveIpv4AddressType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalresultExtendedCommunityTransitiveIpv4AddressType struct {
	obj *resultExtendedCommunityTransitiveIpv4AddressType
}

type marshalResultExtendedCommunityTransitiveIpv4AddressType interface {
	// ToProto marshals ResultExtendedCommunityTransitiveIpv4AddressType to protobuf object *otg.ResultExtendedCommunityTransitiveIpv4AddressType
	ToProto() (*otg.ResultExtendedCommunityTransitiveIpv4AddressType, error)
	// ToPbText marshals ResultExtendedCommunityTransitiveIpv4AddressType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ResultExtendedCommunityTransitiveIpv4AddressType to YAML text
	ToYaml() (string, error)
	// ToJson marshals ResultExtendedCommunityTransitiveIpv4AddressType to JSON text
	ToJson() (string, error)
}

type unMarshalresultExtendedCommunityTransitiveIpv4AddressType struct {
	obj *resultExtendedCommunityTransitiveIpv4AddressType
}

type unMarshalResultExtendedCommunityTransitiveIpv4AddressType interface {
	// FromProto unmarshals ResultExtendedCommunityTransitiveIpv4AddressType from protobuf object *otg.ResultExtendedCommunityTransitiveIpv4AddressType
	FromProto(msg *otg.ResultExtendedCommunityTransitiveIpv4AddressType) (ResultExtendedCommunityTransitiveIpv4AddressType, error)
	// FromPbText unmarshals ResultExtendedCommunityTransitiveIpv4AddressType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ResultExtendedCommunityTransitiveIpv4AddressType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ResultExtendedCommunityTransitiveIpv4AddressType from JSON text
	FromJson(value string) error
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressType) Marshal() marshalResultExtendedCommunityTransitiveIpv4AddressType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalresultExtendedCommunityTransitiveIpv4AddressType{obj: obj}
	}
	return obj.marshaller
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressType) Unmarshal() unMarshalResultExtendedCommunityTransitiveIpv4AddressType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalresultExtendedCommunityTransitiveIpv4AddressType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalresultExtendedCommunityTransitiveIpv4AddressType) ToProto() (*otg.ResultExtendedCommunityTransitiveIpv4AddressType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalresultExtendedCommunityTransitiveIpv4AddressType) FromProto(msg *otg.ResultExtendedCommunityTransitiveIpv4AddressType) (ResultExtendedCommunityTransitiveIpv4AddressType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalresultExtendedCommunityTransitiveIpv4AddressType) ToPbText() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitiveIpv4AddressType) FromPbText(value string) error {
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

func (m *marshalresultExtendedCommunityTransitiveIpv4AddressType) ToYaml() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitiveIpv4AddressType) FromYaml(value string) error {
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

func (m *marshalresultExtendedCommunityTransitiveIpv4AddressType) ToJson() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitiveIpv4AddressType) FromJson(value string) error {
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

func (obj *resultExtendedCommunityTransitiveIpv4AddressType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressType) Clone() (ResultExtendedCommunityTransitiveIpv4AddressType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewResultExtendedCommunityTransitiveIpv4AddressType()
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

func (obj *resultExtendedCommunityTransitiveIpv4AddressType) setNil() {
	obj.routeTargetSubtypeHolder = nil
	obj.routeOriginSubtypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ResultExtendedCommunityTransitiveIpv4AddressType is the Transitive IPv4 Address Specific Extended Community is sent as type 0x01.
type ResultExtendedCommunityTransitiveIpv4AddressType interface {
	Validation
	// msg marshals ResultExtendedCommunityTransitiveIpv4AddressType to protobuf object *otg.ResultExtendedCommunityTransitiveIpv4AddressType
	// and doesn't set defaults
	msg() *otg.ResultExtendedCommunityTransitiveIpv4AddressType
	// setMsg unmarshals ResultExtendedCommunityTransitiveIpv4AddressType from protobuf object *otg.ResultExtendedCommunityTransitiveIpv4AddressType
	// and doesn't set defaults
	setMsg(*otg.ResultExtendedCommunityTransitiveIpv4AddressType) ResultExtendedCommunityTransitiveIpv4AddressType
	// provides marshal interface
	Marshal() marshalResultExtendedCommunityTransitiveIpv4AddressType
	// provides unmarshal interface
	Unmarshal() unMarshalResultExtendedCommunityTransitiveIpv4AddressType
	// validate validates ResultExtendedCommunityTransitiveIpv4AddressType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ResultExtendedCommunityTransitiveIpv4AddressType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ResultExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum, set in ResultExtendedCommunityTransitiveIpv4AddressType
	Choice() ResultExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum
	// setChoice assigns ResultExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum provided by user to ResultExtendedCommunityTransitiveIpv4AddressType
	setChoice(value ResultExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum) ResultExtendedCommunityTransitiveIpv4AddressType
	// HasChoice checks if Choice has been set in ResultExtendedCommunityTransitiveIpv4AddressType
	HasChoice() bool
	// RouteTargetSubtype returns ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget, set in ResultExtendedCommunityTransitiveIpv4AddressType.
	// ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget is the Route Target Community identifies one or more routers that may receive a set of routes (that carry this Community) carried by BGP.  It is sent with sub-type as 0x02.
	RouteTargetSubtype() ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	// SetRouteTargetSubtype assigns ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget provided by user to ResultExtendedCommunityTransitiveIpv4AddressType.
	// ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget is the Route Target Community identifies one or more routers that may receive a set of routes (that carry this Community) carried by BGP.  It is sent with sub-type as 0x02.
	SetRouteTargetSubtype(value ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) ResultExtendedCommunityTransitiveIpv4AddressType
	// HasRouteTargetSubtype checks if RouteTargetSubtype has been set in ResultExtendedCommunityTransitiveIpv4AddressType
	HasRouteTargetSubtype() bool
	// RouteOriginSubtype returns ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin, set in ResultExtendedCommunityTransitiveIpv4AddressType.
	// ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin is the Route Origin Community identifies one or more routers that inject a set of routes (that carry this Community) into BGP It is sent with sub-type as 0x03.
	RouteOriginSubtype() ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	// SetRouteOriginSubtype assigns ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin provided by user to ResultExtendedCommunityTransitiveIpv4AddressType.
	// ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin is the Route Origin Community identifies one or more routers that inject a set of routes (that carry this Community) into BGP It is sent with sub-type as 0x03.
	SetRouteOriginSubtype(value ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) ResultExtendedCommunityTransitiveIpv4AddressType
	// HasRouteOriginSubtype checks if RouteOriginSubtype has been set in ResultExtendedCommunityTransitiveIpv4AddressType
	HasRouteOriginSubtype() bool
	setNil()
}

type ResultExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum string

// Enum of Choice on ResultExtendedCommunityTransitiveIpv4AddressType
var ResultExtendedCommunityTransitiveIpv4AddressTypeChoice = struct {
	ROUTE_TARGET_SUBTYPE ResultExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum
	ROUTE_ORIGIN_SUBTYPE ResultExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum
}{
	ROUTE_TARGET_SUBTYPE: ResultExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum("route_target_subtype"),
	ROUTE_ORIGIN_SUBTYPE: ResultExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum("route_origin_subtype"),
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressType) Choice() ResultExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum {
	return ResultExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *resultExtendedCommunityTransitiveIpv4AddressType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressType) setChoice(value ResultExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum) ResultExtendedCommunityTransitiveIpv4AddressType {
	intValue, ok := otg.ResultExtendedCommunityTransitiveIpv4AddressType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ResultExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ResultExtendedCommunityTransitiveIpv4AddressType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.RouteOriginSubtype = nil
	obj.routeOriginSubtypeHolder = nil
	obj.obj.RouteTargetSubtype = nil
	obj.routeTargetSubtypeHolder = nil

	if value == ResultExtendedCommunityTransitiveIpv4AddressTypeChoice.ROUTE_TARGET_SUBTYPE {
		obj.obj.RouteTargetSubtype = NewResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget().msg()
	}

	if value == ResultExtendedCommunityTransitiveIpv4AddressTypeChoice.ROUTE_ORIGIN_SUBTYPE {
		obj.obj.RouteOriginSubtype = NewResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin().msg()
	}

	return obj
}

// description is TBD
// RouteTargetSubtype returns a ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
func (obj *resultExtendedCommunityTransitiveIpv4AddressType) RouteTargetSubtype() ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget {
	if obj.obj.RouteTargetSubtype == nil {
		obj.setChoice(ResultExtendedCommunityTransitiveIpv4AddressTypeChoice.ROUTE_TARGET_SUBTYPE)
	}
	if obj.routeTargetSubtypeHolder == nil {
		obj.routeTargetSubtypeHolder = &resultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget{obj: obj.obj.RouteTargetSubtype}
	}
	return obj.routeTargetSubtypeHolder
}

// description is TBD
// RouteTargetSubtype returns a ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
func (obj *resultExtendedCommunityTransitiveIpv4AddressType) HasRouteTargetSubtype() bool {
	return obj.obj.RouteTargetSubtype != nil
}

// description is TBD
// SetRouteTargetSubtype sets the ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget value in the ResultExtendedCommunityTransitiveIpv4AddressType object
func (obj *resultExtendedCommunityTransitiveIpv4AddressType) SetRouteTargetSubtype(value ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) ResultExtendedCommunityTransitiveIpv4AddressType {
	obj.setChoice(ResultExtendedCommunityTransitiveIpv4AddressTypeChoice.ROUTE_TARGET_SUBTYPE)
	obj.routeTargetSubtypeHolder = nil
	obj.obj.RouteTargetSubtype = value.msg()

	return obj
}

// description is TBD
// RouteOriginSubtype returns a ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
func (obj *resultExtendedCommunityTransitiveIpv4AddressType) RouteOriginSubtype() ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin {
	if obj.obj.RouteOriginSubtype == nil {
		obj.setChoice(ResultExtendedCommunityTransitiveIpv4AddressTypeChoice.ROUTE_ORIGIN_SUBTYPE)
	}
	if obj.routeOriginSubtypeHolder == nil {
		obj.routeOriginSubtypeHolder = &resultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin{obj: obj.obj.RouteOriginSubtype}
	}
	return obj.routeOriginSubtypeHolder
}

// description is TBD
// RouteOriginSubtype returns a ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
func (obj *resultExtendedCommunityTransitiveIpv4AddressType) HasRouteOriginSubtype() bool {
	return obj.obj.RouteOriginSubtype != nil
}

// description is TBD
// SetRouteOriginSubtype sets the ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin value in the ResultExtendedCommunityTransitiveIpv4AddressType object
func (obj *resultExtendedCommunityTransitiveIpv4AddressType) SetRouteOriginSubtype(value ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) ResultExtendedCommunityTransitiveIpv4AddressType {
	obj.setChoice(ResultExtendedCommunityTransitiveIpv4AddressTypeChoice.ROUTE_ORIGIN_SUBTYPE)
	obj.routeOriginSubtypeHolder = nil
	obj.obj.RouteOriginSubtype = value.msg()

	return obj
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressType) validateObj(vObj *validation, set_default bool) {
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

func (obj *resultExtendedCommunityTransitiveIpv4AddressType) setDefault() {
	var choices_set int = 0
	var choice ResultExtendedCommunityTransitiveIpv4AddressTypeChoiceEnum

	if obj.obj.RouteTargetSubtype != nil {
		choices_set += 1
		choice = ResultExtendedCommunityTransitiveIpv4AddressTypeChoice.ROUTE_TARGET_SUBTYPE
	}

	if obj.obj.RouteOriginSubtype != nil {
		choices_set += 1
		choice = ResultExtendedCommunityTransitiveIpv4AddressTypeChoice.ROUTE_ORIGIN_SUBTYPE
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ResultExtendedCommunityTransitiveIpv4AddressType")
			}
		} else {
			intVal := otg.ResultExtendedCommunityTransitiveIpv4AddressType_Choice_Enum_value[string(choice)]
			enumValue := otg.ResultExtendedCommunityTransitiveIpv4AddressType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ResultExtendedCommunityTransitive4OctetAsType *****
type resultExtendedCommunityTransitive4OctetAsType struct {
	validation
	obj                      *otg.ResultExtendedCommunityTransitive4OctetAsType
	marshaller               marshalResultExtendedCommunityTransitive4OctetAsType
	unMarshaller             unMarshalResultExtendedCommunityTransitive4OctetAsType
	routeTargetSubtypeHolder ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
	routeOriginSubtypeHolder ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
}

func NewResultExtendedCommunityTransitive4OctetAsType() ResultExtendedCommunityTransitive4OctetAsType {
	obj := resultExtendedCommunityTransitive4OctetAsType{obj: &otg.ResultExtendedCommunityTransitive4OctetAsType{}}
	obj.setDefault()
	return &obj
}

func (obj *resultExtendedCommunityTransitive4OctetAsType) msg() *otg.ResultExtendedCommunityTransitive4OctetAsType {
	return obj.obj
}

func (obj *resultExtendedCommunityTransitive4OctetAsType) setMsg(msg *otg.ResultExtendedCommunityTransitive4OctetAsType) ResultExtendedCommunityTransitive4OctetAsType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalresultExtendedCommunityTransitive4OctetAsType struct {
	obj *resultExtendedCommunityTransitive4OctetAsType
}

type marshalResultExtendedCommunityTransitive4OctetAsType interface {
	// ToProto marshals ResultExtendedCommunityTransitive4OctetAsType to protobuf object *otg.ResultExtendedCommunityTransitive4OctetAsType
	ToProto() (*otg.ResultExtendedCommunityTransitive4OctetAsType, error)
	// ToPbText marshals ResultExtendedCommunityTransitive4OctetAsType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ResultExtendedCommunityTransitive4OctetAsType to YAML text
	ToYaml() (string, error)
	// ToJson marshals ResultExtendedCommunityTransitive4OctetAsType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals ResultExtendedCommunityTransitive4OctetAsType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalresultExtendedCommunityTransitive4OctetAsType struct {
	obj *resultExtendedCommunityTransitive4OctetAsType
}

type unMarshalResultExtendedCommunityTransitive4OctetAsType interface {
	// FromProto unmarshals ResultExtendedCommunityTransitive4OctetAsType from protobuf object *otg.ResultExtendedCommunityTransitive4OctetAsType
	FromProto(msg *otg.ResultExtendedCommunityTransitive4OctetAsType) (ResultExtendedCommunityTransitive4OctetAsType, error)
	// FromPbText unmarshals ResultExtendedCommunityTransitive4OctetAsType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ResultExtendedCommunityTransitive4OctetAsType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ResultExtendedCommunityTransitive4OctetAsType from JSON text
	FromJson(value string) error
}

func (obj *resultExtendedCommunityTransitive4OctetAsType) Marshal() marshalResultExtendedCommunityTransitive4OctetAsType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalresultExtendedCommunityTransitive4OctetAsType{obj: obj}
	}
	return obj.marshaller
}

func (obj *resultExtendedCommunityTransitive4OctetAsType) Unmarshal() unMarshalResultExtendedCommunityTransitive4OctetAsType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalresultExtendedCommunityTransitive4OctetAsType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalresultExtendedCommunityTransitive4OctetAsType) ToProto() (*otg.ResultExtendedCommunityTransitive4OctetAsType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalresultExtendedCommunityTransitive4OctetAsType) FromProto(msg *otg.ResultExtendedCommunityTransitive4OctetAsType) (ResultExtendedCommunityTransitive4OctetAsType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalresultExtendedCommunityTransitive4OctetAsType) ToPbText() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitive4OctetAsType) FromPbText(value string) error {
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

func (m *marshalresultExtendedCommunityTransitive4OctetAsType) ToYaml() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitive4OctetAsType) FromYaml(value string) error {
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

func (m *marshalresultExtendedCommunityTransitive4OctetAsType) ToJsonRaw() (string, error) {
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

func (m *marshalresultExtendedCommunityTransitive4OctetAsType) ToJson() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitive4OctetAsType) FromJson(value string) error {
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

func (obj *resultExtendedCommunityTransitive4OctetAsType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitive4OctetAsType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitive4OctetAsType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *resultExtendedCommunityTransitive4OctetAsType) Clone() (ResultExtendedCommunityTransitive4OctetAsType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewResultExtendedCommunityTransitive4OctetAsType()
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

func (obj *resultExtendedCommunityTransitive4OctetAsType) setNil() {
	obj.routeTargetSubtypeHolder = nil
	obj.routeOriginSubtypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ResultExtendedCommunityTransitive4OctetAsType is the Transitive Four-Octet AS-Specific Extended Community is sent as type 0x02. It is defined in RFC 5668.
type ResultExtendedCommunityTransitive4OctetAsType interface {
	Validation
	// msg marshals ResultExtendedCommunityTransitive4OctetAsType to protobuf object *otg.ResultExtendedCommunityTransitive4OctetAsType
	// and doesn't set defaults
	msg() *otg.ResultExtendedCommunityTransitive4OctetAsType
	// setMsg unmarshals ResultExtendedCommunityTransitive4OctetAsType from protobuf object *otg.ResultExtendedCommunityTransitive4OctetAsType
	// and doesn't set defaults
	setMsg(*otg.ResultExtendedCommunityTransitive4OctetAsType) ResultExtendedCommunityTransitive4OctetAsType
	// provides marshal interface
	Marshal() marshalResultExtendedCommunityTransitive4OctetAsType
	// provides unmarshal interface
	Unmarshal() unMarshalResultExtendedCommunityTransitive4OctetAsType
	// validate validates ResultExtendedCommunityTransitive4OctetAsType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ResultExtendedCommunityTransitive4OctetAsType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ResultExtendedCommunityTransitive4OctetAsTypeChoiceEnum, set in ResultExtendedCommunityTransitive4OctetAsType
	Choice() ResultExtendedCommunityTransitive4OctetAsTypeChoiceEnum
	// setChoice assigns ResultExtendedCommunityTransitive4OctetAsTypeChoiceEnum provided by user to ResultExtendedCommunityTransitive4OctetAsType
	setChoice(value ResultExtendedCommunityTransitive4OctetAsTypeChoiceEnum) ResultExtendedCommunityTransitive4OctetAsType
	// HasChoice checks if Choice has been set in ResultExtendedCommunityTransitive4OctetAsType
	HasChoice() bool
	// RouteTargetSubtype returns ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget, set in ResultExtendedCommunityTransitive4OctetAsType.
	// ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget is the Route Target Community identifies one or more routers that may receive a set of routes (that carry this Community) carried by BGP.  It is sent with sub-type as 0x02
	RouteTargetSubtype() ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
	// SetRouteTargetSubtype assigns ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget provided by user to ResultExtendedCommunityTransitive4OctetAsType.
	// ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget is the Route Target Community identifies one or more routers that may receive a set of routes (that carry this Community) carried by BGP.  It is sent with sub-type as 0x02
	SetRouteTargetSubtype(value ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget) ResultExtendedCommunityTransitive4OctetAsType
	// HasRouteTargetSubtype checks if RouteTargetSubtype has been set in ResultExtendedCommunityTransitive4OctetAsType
	HasRouteTargetSubtype() bool
	// RouteOriginSubtype returns ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin, set in ResultExtendedCommunityTransitive4OctetAsType.
	// ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin is the Route Origin Community identifies one or more routers that inject a set of routes (that carry this Community) into BGP. It is sent with sub-type as 0x03.
	RouteOriginSubtype() ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	// SetRouteOriginSubtype assigns ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin provided by user to ResultExtendedCommunityTransitive4OctetAsType.
	// ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin is the Route Origin Community identifies one or more routers that inject a set of routes (that carry this Community) into BGP. It is sent with sub-type as 0x03.
	SetRouteOriginSubtype(value ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) ResultExtendedCommunityTransitive4OctetAsType
	// HasRouteOriginSubtype checks if RouteOriginSubtype has been set in ResultExtendedCommunityTransitive4OctetAsType
	HasRouteOriginSubtype() bool
	setNil()
}

type ResultExtendedCommunityTransitive4OctetAsTypeChoiceEnum string

// Enum of Choice on ResultExtendedCommunityTransitive4OctetAsType
var ResultExtendedCommunityTransitive4OctetAsTypeChoice = struct {
	ROUTE_TARGET_SUBTYPE ResultExtendedCommunityTransitive4OctetAsTypeChoiceEnum
	ROUTE_ORIGIN_SUBTYPE ResultExtendedCommunityTransitive4OctetAsTypeChoiceEnum
}{
	ROUTE_TARGET_SUBTYPE: ResultExtendedCommunityTransitive4OctetAsTypeChoiceEnum("route_target_subtype"),
	ROUTE_ORIGIN_SUBTYPE: ResultExtendedCommunityTransitive4OctetAsTypeChoiceEnum("route_origin_subtype"),
}

func (obj *resultExtendedCommunityTransitive4OctetAsType) Choice() ResultExtendedCommunityTransitive4OctetAsTypeChoiceEnum {
	return ResultExtendedCommunityTransitive4OctetAsTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *resultExtendedCommunityTransitive4OctetAsType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *resultExtendedCommunityTransitive4OctetAsType) setChoice(value ResultExtendedCommunityTransitive4OctetAsTypeChoiceEnum) ResultExtendedCommunityTransitive4OctetAsType {
	intValue, ok := otg.ResultExtendedCommunityTransitive4OctetAsType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ResultExtendedCommunityTransitive4OctetAsTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ResultExtendedCommunityTransitive4OctetAsType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.RouteOriginSubtype = nil
	obj.routeOriginSubtypeHolder = nil
	obj.obj.RouteTargetSubtype = nil
	obj.routeTargetSubtypeHolder = nil

	if value == ResultExtendedCommunityTransitive4OctetAsTypeChoice.ROUTE_TARGET_SUBTYPE {
		obj.obj.RouteTargetSubtype = NewResultExtendedCommunityTransitive4OctetAsTypeRouteTarget().msg()
	}

	if value == ResultExtendedCommunityTransitive4OctetAsTypeChoice.ROUTE_ORIGIN_SUBTYPE {
		obj.obj.RouteOriginSubtype = NewResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin().msg()
	}

	return obj
}

// description is TBD
// RouteTargetSubtype returns a ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
func (obj *resultExtendedCommunityTransitive4OctetAsType) RouteTargetSubtype() ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget {
	if obj.obj.RouteTargetSubtype == nil {
		obj.setChoice(ResultExtendedCommunityTransitive4OctetAsTypeChoice.ROUTE_TARGET_SUBTYPE)
	}
	if obj.routeTargetSubtypeHolder == nil {
		obj.routeTargetSubtypeHolder = &resultExtendedCommunityTransitive4OctetAsTypeRouteTarget{obj: obj.obj.RouteTargetSubtype}
	}
	return obj.routeTargetSubtypeHolder
}

// description is TBD
// RouteTargetSubtype returns a ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget
func (obj *resultExtendedCommunityTransitive4OctetAsType) HasRouteTargetSubtype() bool {
	return obj.obj.RouteTargetSubtype != nil
}

// description is TBD
// SetRouteTargetSubtype sets the ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget value in the ResultExtendedCommunityTransitive4OctetAsType object
func (obj *resultExtendedCommunityTransitive4OctetAsType) SetRouteTargetSubtype(value ResultExtendedCommunityTransitive4OctetAsTypeRouteTarget) ResultExtendedCommunityTransitive4OctetAsType {
	obj.setChoice(ResultExtendedCommunityTransitive4OctetAsTypeChoice.ROUTE_TARGET_SUBTYPE)
	obj.routeTargetSubtypeHolder = nil
	obj.obj.RouteTargetSubtype = value.msg()

	return obj
}

// description is TBD
// RouteOriginSubtype returns a ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
func (obj *resultExtendedCommunityTransitive4OctetAsType) RouteOriginSubtype() ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin {
	if obj.obj.RouteOriginSubtype == nil {
		obj.setChoice(ResultExtendedCommunityTransitive4OctetAsTypeChoice.ROUTE_ORIGIN_SUBTYPE)
	}
	if obj.routeOriginSubtypeHolder == nil {
		obj.routeOriginSubtypeHolder = &resultExtendedCommunityTransitive4OctetAsTypeRouteOrigin{obj: obj.obj.RouteOriginSubtype}
	}
	return obj.routeOriginSubtypeHolder
}

// description is TBD
// RouteOriginSubtype returns a ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin
func (obj *resultExtendedCommunityTransitive4OctetAsType) HasRouteOriginSubtype() bool {
	return obj.obj.RouteOriginSubtype != nil
}

// description is TBD
// SetRouteOriginSubtype sets the ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin value in the ResultExtendedCommunityTransitive4OctetAsType object
func (obj *resultExtendedCommunityTransitive4OctetAsType) SetRouteOriginSubtype(value ResultExtendedCommunityTransitive4OctetAsTypeRouteOrigin) ResultExtendedCommunityTransitive4OctetAsType {
	obj.setChoice(ResultExtendedCommunityTransitive4OctetAsTypeChoice.ROUTE_ORIGIN_SUBTYPE)
	obj.routeOriginSubtypeHolder = nil
	obj.obj.RouteOriginSubtype = value.msg()

	return obj
}

func (obj *resultExtendedCommunityTransitive4OctetAsType) validateObj(vObj *validation, set_default bool) {
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

func (obj *resultExtendedCommunityTransitive4OctetAsType) setDefault() {
	var choices_set int = 0
	var choice ResultExtendedCommunityTransitive4OctetAsTypeChoiceEnum

	if obj.obj.RouteTargetSubtype != nil {
		choices_set += 1
		choice = ResultExtendedCommunityTransitive4OctetAsTypeChoice.ROUTE_TARGET_SUBTYPE
	}

	if obj.obj.RouteOriginSubtype != nil {
		choices_set += 1
		choice = ResultExtendedCommunityTransitive4OctetAsTypeChoice.ROUTE_ORIGIN_SUBTYPE
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ResultExtendedCommunityTransitive4OctetAsType")
			}
		} else {
			intVal := otg.ResultExtendedCommunityTransitive4OctetAsType_Choice_Enum_value[string(choice)]
			enumValue := otg.ResultExtendedCommunityTransitive4OctetAsType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

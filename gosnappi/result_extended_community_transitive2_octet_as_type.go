package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ResultExtendedCommunityTransitive2OctetAsType *****
type resultExtendedCommunityTransitive2OctetAsType struct {
	validation
	obj                      *otg.ResultExtendedCommunityTransitive2OctetAsType
	marshaller               marshalResultExtendedCommunityTransitive2OctetAsType
	unMarshaller             unMarshalResultExtendedCommunityTransitive2OctetAsType
	routeTargetSubtypeHolder ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
	routeOriginSubtypeHolder ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
}

func NewResultExtendedCommunityTransitive2OctetAsType() ResultExtendedCommunityTransitive2OctetAsType {
	obj := resultExtendedCommunityTransitive2OctetAsType{obj: &otg.ResultExtendedCommunityTransitive2OctetAsType{}}
	obj.setDefault()
	return &obj
}

func (obj *resultExtendedCommunityTransitive2OctetAsType) msg() *otg.ResultExtendedCommunityTransitive2OctetAsType {
	return obj.obj
}

func (obj *resultExtendedCommunityTransitive2OctetAsType) setMsg(msg *otg.ResultExtendedCommunityTransitive2OctetAsType) ResultExtendedCommunityTransitive2OctetAsType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalresultExtendedCommunityTransitive2OctetAsType struct {
	obj *resultExtendedCommunityTransitive2OctetAsType
}

type marshalResultExtendedCommunityTransitive2OctetAsType interface {
	// ToProto marshals ResultExtendedCommunityTransitive2OctetAsType to protobuf object *otg.ResultExtendedCommunityTransitive2OctetAsType
	ToProto() (*otg.ResultExtendedCommunityTransitive2OctetAsType, error)
	// ToPbText marshals ResultExtendedCommunityTransitive2OctetAsType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ResultExtendedCommunityTransitive2OctetAsType to YAML text
	ToYaml() (string, error)
	// ToJson marshals ResultExtendedCommunityTransitive2OctetAsType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals ResultExtendedCommunityTransitive2OctetAsType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalresultExtendedCommunityTransitive2OctetAsType struct {
	obj *resultExtendedCommunityTransitive2OctetAsType
}

type unMarshalResultExtendedCommunityTransitive2OctetAsType interface {
	// FromProto unmarshals ResultExtendedCommunityTransitive2OctetAsType from protobuf object *otg.ResultExtendedCommunityTransitive2OctetAsType
	FromProto(msg *otg.ResultExtendedCommunityTransitive2OctetAsType) (ResultExtendedCommunityTransitive2OctetAsType, error)
	// FromPbText unmarshals ResultExtendedCommunityTransitive2OctetAsType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ResultExtendedCommunityTransitive2OctetAsType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ResultExtendedCommunityTransitive2OctetAsType from JSON text
	FromJson(value string) error
}

func (obj *resultExtendedCommunityTransitive2OctetAsType) Marshal() marshalResultExtendedCommunityTransitive2OctetAsType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalresultExtendedCommunityTransitive2OctetAsType{obj: obj}
	}
	return obj.marshaller
}

func (obj *resultExtendedCommunityTransitive2OctetAsType) Unmarshal() unMarshalResultExtendedCommunityTransitive2OctetAsType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalresultExtendedCommunityTransitive2OctetAsType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalresultExtendedCommunityTransitive2OctetAsType) ToProto() (*otg.ResultExtendedCommunityTransitive2OctetAsType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalresultExtendedCommunityTransitive2OctetAsType) FromProto(msg *otg.ResultExtendedCommunityTransitive2OctetAsType) (ResultExtendedCommunityTransitive2OctetAsType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalresultExtendedCommunityTransitive2OctetAsType) ToPbText() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitive2OctetAsType) FromPbText(value string) error {
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

func (m *marshalresultExtendedCommunityTransitive2OctetAsType) ToYaml() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitive2OctetAsType) FromYaml(value string) error {
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

func (m *marshalresultExtendedCommunityTransitive2OctetAsType) ToJsonRaw() (string, error) {
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

func (m *marshalresultExtendedCommunityTransitive2OctetAsType) ToJson() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitive2OctetAsType) FromJson(value string) error {
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

func (obj *resultExtendedCommunityTransitive2OctetAsType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitive2OctetAsType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitive2OctetAsType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *resultExtendedCommunityTransitive2OctetAsType) Clone() (ResultExtendedCommunityTransitive2OctetAsType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewResultExtendedCommunityTransitive2OctetAsType()
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

func (obj *resultExtendedCommunityTransitive2OctetAsType) setNil() {
	obj.routeTargetSubtypeHolder = nil
	obj.routeOriginSubtypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ResultExtendedCommunityTransitive2OctetAsType is the Transitive Two-Octet AS-Specific Extended Community is sent as type 0x00 .
type ResultExtendedCommunityTransitive2OctetAsType interface {
	Validation
	// msg marshals ResultExtendedCommunityTransitive2OctetAsType to protobuf object *otg.ResultExtendedCommunityTransitive2OctetAsType
	// and doesn't set defaults
	msg() *otg.ResultExtendedCommunityTransitive2OctetAsType
	// setMsg unmarshals ResultExtendedCommunityTransitive2OctetAsType from protobuf object *otg.ResultExtendedCommunityTransitive2OctetAsType
	// and doesn't set defaults
	setMsg(*otg.ResultExtendedCommunityTransitive2OctetAsType) ResultExtendedCommunityTransitive2OctetAsType
	// provides marshal interface
	Marshal() marshalResultExtendedCommunityTransitive2OctetAsType
	// provides unmarshal interface
	Unmarshal() unMarshalResultExtendedCommunityTransitive2OctetAsType
	// validate validates ResultExtendedCommunityTransitive2OctetAsType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ResultExtendedCommunityTransitive2OctetAsType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ResultExtendedCommunityTransitive2OctetAsTypeChoiceEnum, set in ResultExtendedCommunityTransitive2OctetAsType
	Choice() ResultExtendedCommunityTransitive2OctetAsTypeChoiceEnum
	// setChoice assigns ResultExtendedCommunityTransitive2OctetAsTypeChoiceEnum provided by user to ResultExtendedCommunityTransitive2OctetAsType
	setChoice(value ResultExtendedCommunityTransitive2OctetAsTypeChoiceEnum) ResultExtendedCommunityTransitive2OctetAsType
	// HasChoice checks if Choice has been set in ResultExtendedCommunityTransitive2OctetAsType
	HasChoice() bool
	// RouteTargetSubtype returns ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget, set in ResultExtendedCommunityTransitive2OctetAsType.
	// ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget is the Route Target Community identifies one or more routers that may receive a set of routes (that carry this Community) carried by BGP Update message.  It is sent with sub-type as 0x02.
	RouteTargetSubtype() ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
	// SetRouteTargetSubtype assigns ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget provided by user to ResultExtendedCommunityTransitive2OctetAsType.
	// ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget is the Route Target Community identifies one or more routers that may receive a set of routes (that carry this Community) carried by BGP Update message.  It is sent with sub-type as 0x02.
	SetRouteTargetSubtype(value ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget) ResultExtendedCommunityTransitive2OctetAsType
	// HasRouteTargetSubtype checks if RouteTargetSubtype has been set in ResultExtendedCommunityTransitive2OctetAsType
	HasRouteTargetSubtype() bool
	// RouteOriginSubtype returns ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin, set in ResultExtendedCommunityTransitive2OctetAsType.
	// ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin is the Route Origin Community identifies one or more routers that inject a set of routes (that carry this Community) into BGP. It is sent with sub-type as 0x03 .
	RouteOriginSubtype() ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	// SetRouteOriginSubtype assigns ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin provided by user to ResultExtendedCommunityTransitive2OctetAsType.
	// ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin is the Route Origin Community identifies one or more routers that inject a set of routes (that carry this Community) into BGP. It is sent with sub-type as 0x03 .
	SetRouteOriginSubtype(value ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) ResultExtendedCommunityTransitive2OctetAsType
	// HasRouteOriginSubtype checks if RouteOriginSubtype has been set in ResultExtendedCommunityTransitive2OctetAsType
	HasRouteOriginSubtype() bool
	setNil()
}

type ResultExtendedCommunityTransitive2OctetAsTypeChoiceEnum string

// Enum of Choice on ResultExtendedCommunityTransitive2OctetAsType
var ResultExtendedCommunityTransitive2OctetAsTypeChoice = struct {
	ROUTE_TARGET_SUBTYPE ResultExtendedCommunityTransitive2OctetAsTypeChoiceEnum
	ROUTE_ORIGIN_SUBTYPE ResultExtendedCommunityTransitive2OctetAsTypeChoiceEnum
}{
	ROUTE_TARGET_SUBTYPE: ResultExtendedCommunityTransitive2OctetAsTypeChoiceEnum("route_target_subtype"),
	ROUTE_ORIGIN_SUBTYPE: ResultExtendedCommunityTransitive2OctetAsTypeChoiceEnum("route_origin_subtype"),
}

func (obj *resultExtendedCommunityTransitive2OctetAsType) Choice() ResultExtendedCommunityTransitive2OctetAsTypeChoiceEnum {
	return ResultExtendedCommunityTransitive2OctetAsTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *resultExtendedCommunityTransitive2OctetAsType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *resultExtendedCommunityTransitive2OctetAsType) setChoice(value ResultExtendedCommunityTransitive2OctetAsTypeChoiceEnum) ResultExtendedCommunityTransitive2OctetAsType {
	intValue, ok := otg.ResultExtendedCommunityTransitive2OctetAsType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ResultExtendedCommunityTransitive2OctetAsTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ResultExtendedCommunityTransitive2OctetAsType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.RouteOriginSubtype = nil
	obj.routeOriginSubtypeHolder = nil
	obj.obj.RouteTargetSubtype = nil
	obj.routeTargetSubtypeHolder = nil

	if value == ResultExtendedCommunityTransitive2OctetAsTypeChoice.ROUTE_TARGET_SUBTYPE {
		obj.obj.RouteTargetSubtype = NewResultExtendedCommunityTransitive2OctetAsTypeRouteTarget().msg()
	}

	if value == ResultExtendedCommunityTransitive2OctetAsTypeChoice.ROUTE_ORIGIN_SUBTYPE {
		obj.obj.RouteOriginSubtype = NewResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin().msg()
	}

	return obj
}

// description is TBD
// RouteTargetSubtype returns a ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
func (obj *resultExtendedCommunityTransitive2OctetAsType) RouteTargetSubtype() ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget {
	if obj.obj.RouteTargetSubtype == nil {
		obj.setChoice(ResultExtendedCommunityTransitive2OctetAsTypeChoice.ROUTE_TARGET_SUBTYPE)
	}
	if obj.routeTargetSubtypeHolder == nil {
		obj.routeTargetSubtypeHolder = &resultExtendedCommunityTransitive2OctetAsTypeRouteTarget{obj: obj.obj.RouteTargetSubtype}
	}
	return obj.routeTargetSubtypeHolder
}

// description is TBD
// RouteTargetSubtype returns a ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget
func (obj *resultExtendedCommunityTransitive2OctetAsType) HasRouteTargetSubtype() bool {
	return obj.obj.RouteTargetSubtype != nil
}

// description is TBD
// SetRouteTargetSubtype sets the ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget value in the ResultExtendedCommunityTransitive2OctetAsType object
func (obj *resultExtendedCommunityTransitive2OctetAsType) SetRouteTargetSubtype(value ResultExtendedCommunityTransitive2OctetAsTypeRouteTarget) ResultExtendedCommunityTransitive2OctetAsType {
	obj.setChoice(ResultExtendedCommunityTransitive2OctetAsTypeChoice.ROUTE_TARGET_SUBTYPE)
	obj.routeTargetSubtypeHolder = nil
	obj.obj.RouteTargetSubtype = value.msg()

	return obj
}

// description is TBD
// RouteOriginSubtype returns a ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
func (obj *resultExtendedCommunityTransitive2OctetAsType) RouteOriginSubtype() ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin {
	if obj.obj.RouteOriginSubtype == nil {
		obj.setChoice(ResultExtendedCommunityTransitive2OctetAsTypeChoice.ROUTE_ORIGIN_SUBTYPE)
	}
	if obj.routeOriginSubtypeHolder == nil {
		obj.routeOriginSubtypeHolder = &resultExtendedCommunityTransitive2OctetAsTypeRouteOrigin{obj: obj.obj.RouteOriginSubtype}
	}
	return obj.routeOriginSubtypeHolder
}

// description is TBD
// RouteOriginSubtype returns a ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
func (obj *resultExtendedCommunityTransitive2OctetAsType) HasRouteOriginSubtype() bool {
	return obj.obj.RouteOriginSubtype != nil
}

// description is TBD
// SetRouteOriginSubtype sets the ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin value in the ResultExtendedCommunityTransitive2OctetAsType object
func (obj *resultExtendedCommunityTransitive2OctetAsType) SetRouteOriginSubtype(value ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) ResultExtendedCommunityTransitive2OctetAsType {
	obj.setChoice(ResultExtendedCommunityTransitive2OctetAsTypeChoice.ROUTE_ORIGIN_SUBTYPE)
	obj.routeOriginSubtypeHolder = nil
	obj.obj.RouteOriginSubtype = value.msg()

	return obj
}

func (obj *resultExtendedCommunityTransitive2OctetAsType) validateObj(vObj *validation, set_default bool) {
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

func (obj *resultExtendedCommunityTransitive2OctetAsType) setDefault() {
	var choices_set int = 0
	var choice ResultExtendedCommunityTransitive2OctetAsTypeChoiceEnum

	if obj.obj.RouteTargetSubtype != nil {
		choices_set += 1
		choice = ResultExtendedCommunityTransitive2OctetAsTypeChoice.ROUTE_TARGET_SUBTYPE
	}

	if obj.obj.RouteOriginSubtype != nil {
		choices_set += 1
		choice = ResultExtendedCommunityTransitive2OctetAsTypeChoice.ROUTE_ORIGIN_SUBTYPE
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ResultExtendedCommunityTransitive2OctetAsType")
			}
		} else {
			intVal := otg.ResultExtendedCommunityTransitive2OctetAsType_Choice_Enum_value[string(choice)]
			enumValue := otg.ResultExtendedCommunityTransitive2OctetAsType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

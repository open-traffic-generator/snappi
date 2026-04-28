package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRv6AdvertiseLocatorAsPrefix *****
type isisSRv6AdvertiseLocatorAsPrefix struct {
	validation
	obj                    *otg.IsisSRv6AdvertiseLocatorAsPrefix
	marshaller             marshalIsisSRv6AdvertiseLocatorAsPrefix
	unMarshaller           unMarshalIsisSRv6AdvertiseLocatorAsPrefix
	prefixAttributesHolder IsisSRv6PrefixAttributes
}

func NewIsisSRv6AdvertiseLocatorAsPrefix() IsisSRv6AdvertiseLocatorAsPrefix {
	obj := isisSRv6AdvertiseLocatorAsPrefix{obj: &otg.IsisSRv6AdvertiseLocatorAsPrefix{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRv6AdvertiseLocatorAsPrefix) msg() *otg.IsisSRv6AdvertiseLocatorAsPrefix {
	return obj.obj
}

func (obj *isisSRv6AdvertiseLocatorAsPrefix) setMsg(msg *otg.IsisSRv6AdvertiseLocatorAsPrefix) IsisSRv6AdvertiseLocatorAsPrefix {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRv6AdvertiseLocatorAsPrefix struct {
	obj *isisSRv6AdvertiseLocatorAsPrefix
}

type marshalIsisSRv6AdvertiseLocatorAsPrefix interface {
	// ToProto marshals IsisSRv6AdvertiseLocatorAsPrefix to protobuf object *otg.IsisSRv6AdvertiseLocatorAsPrefix
	ToProto() (*otg.IsisSRv6AdvertiseLocatorAsPrefix, error)
	// ToPbText marshals IsisSRv6AdvertiseLocatorAsPrefix to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRv6AdvertiseLocatorAsPrefix to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRv6AdvertiseLocatorAsPrefix to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRv6AdvertiseLocatorAsPrefix struct {
	obj *isisSRv6AdvertiseLocatorAsPrefix
}

type unMarshalIsisSRv6AdvertiseLocatorAsPrefix interface {
	// FromProto unmarshals IsisSRv6AdvertiseLocatorAsPrefix from protobuf object *otg.IsisSRv6AdvertiseLocatorAsPrefix
	FromProto(msg *otg.IsisSRv6AdvertiseLocatorAsPrefix) (IsisSRv6AdvertiseLocatorAsPrefix, error)
	// FromPbText unmarshals IsisSRv6AdvertiseLocatorAsPrefix from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRv6AdvertiseLocatorAsPrefix from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRv6AdvertiseLocatorAsPrefix from JSON text
	FromJson(value string) error
}

func (obj *isisSRv6AdvertiseLocatorAsPrefix) Marshal() marshalIsisSRv6AdvertiseLocatorAsPrefix {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRv6AdvertiseLocatorAsPrefix{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRv6AdvertiseLocatorAsPrefix) Unmarshal() unMarshalIsisSRv6AdvertiseLocatorAsPrefix {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRv6AdvertiseLocatorAsPrefix{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRv6AdvertiseLocatorAsPrefix) ToProto() (*otg.IsisSRv6AdvertiseLocatorAsPrefix, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRv6AdvertiseLocatorAsPrefix) FromProto(msg *otg.IsisSRv6AdvertiseLocatorAsPrefix) (IsisSRv6AdvertiseLocatorAsPrefix, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRv6AdvertiseLocatorAsPrefix) ToPbText() (string, error) {
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

func (m *unMarshalisisSRv6AdvertiseLocatorAsPrefix) FromPbText(value string) error {
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

func (m *marshalisisSRv6AdvertiseLocatorAsPrefix) ToYaml() (string, error) {
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

func (m *unMarshalisisSRv6AdvertiseLocatorAsPrefix) FromYaml(value string) error {
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

func (m *marshalisisSRv6AdvertiseLocatorAsPrefix) ToJson() (string, error) {
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

func (m *unMarshalisisSRv6AdvertiseLocatorAsPrefix) FromJson(value string) error {
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

func (obj *isisSRv6AdvertiseLocatorAsPrefix) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRv6AdvertiseLocatorAsPrefix) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRv6AdvertiseLocatorAsPrefix) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRv6AdvertiseLocatorAsPrefix) Clone() (IsisSRv6AdvertiseLocatorAsPrefix, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRv6AdvertiseLocatorAsPrefix()
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

func (obj *isisSRv6AdvertiseLocatorAsPrefix) setNil() {
	obj.prefixAttributesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisSRv6AdvertiseLocatorAsPrefix is controls advertisement of the SRv6 locator prefix as an IS-IS IPv6 Reachability prefix (TLV 236/237) alongside the SRv6 Locator TLV (type 27). When this object is present the secondary prefix advertisement is enabled; when absent it is suppressed. Reference: RFC 9352 Section 7.1.
type IsisSRv6AdvertiseLocatorAsPrefix interface {
	Validation
	// msg marshals IsisSRv6AdvertiseLocatorAsPrefix to protobuf object *otg.IsisSRv6AdvertiseLocatorAsPrefix
	// and doesn't set defaults
	msg() *otg.IsisSRv6AdvertiseLocatorAsPrefix
	// setMsg unmarshals IsisSRv6AdvertiseLocatorAsPrefix from protobuf object *otg.IsisSRv6AdvertiseLocatorAsPrefix
	// and doesn't set defaults
	setMsg(*otg.IsisSRv6AdvertiseLocatorAsPrefix) IsisSRv6AdvertiseLocatorAsPrefix
	// provides marshal interface
	Marshal() marshalIsisSRv6AdvertiseLocatorAsPrefix
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRv6AdvertiseLocatorAsPrefix
	// validate validates IsisSRv6AdvertiseLocatorAsPrefix
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRv6AdvertiseLocatorAsPrefix, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouteMetric returns uint32, set in IsisSRv6AdvertiseLocatorAsPrefix.
	RouteMetric() uint32
	// SetRouteMetric assigns uint32 provided by user to IsisSRv6AdvertiseLocatorAsPrefix
	SetRouteMetric(value uint32) IsisSRv6AdvertiseLocatorAsPrefix
	// HasRouteMetric checks if RouteMetric has been set in IsisSRv6AdvertiseLocatorAsPrefix
	HasRouteMetric() bool
	// RedistributionType returns IsisSRv6AdvertiseLocatorAsPrefixRedistributionTypeEnum, set in IsisSRv6AdvertiseLocatorAsPrefix
	RedistributionType() IsisSRv6AdvertiseLocatorAsPrefixRedistributionTypeEnum
	// SetRedistributionType assigns IsisSRv6AdvertiseLocatorAsPrefixRedistributionTypeEnum provided by user to IsisSRv6AdvertiseLocatorAsPrefix
	SetRedistributionType(value IsisSRv6AdvertiseLocatorAsPrefixRedistributionTypeEnum) IsisSRv6AdvertiseLocatorAsPrefix
	// HasRedistributionType checks if RedistributionType has been set in IsisSRv6AdvertiseLocatorAsPrefix
	HasRedistributionType() bool
	// RouteOrigin returns IsisSRv6AdvertiseLocatorAsPrefixRouteOriginEnum, set in IsisSRv6AdvertiseLocatorAsPrefix
	RouteOrigin() IsisSRv6AdvertiseLocatorAsPrefixRouteOriginEnum
	// SetRouteOrigin assigns IsisSRv6AdvertiseLocatorAsPrefixRouteOriginEnum provided by user to IsisSRv6AdvertiseLocatorAsPrefix
	SetRouteOrigin(value IsisSRv6AdvertiseLocatorAsPrefixRouteOriginEnum) IsisSRv6AdvertiseLocatorAsPrefix
	// HasRouteOrigin checks if RouteOrigin has been set in IsisSRv6AdvertiseLocatorAsPrefix
	HasRouteOrigin() bool
	// PrefixAttributes returns IsisSRv6PrefixAttributes, set in IsisSRv6AdvertiseLocatorAsPrefix.
	// IsisSRv6PrefixAttributes is prefix Attribute Flags Sub-TLV (sub-TLV type 4, RFC 7794) for an SRv6 locator prefix advertisement. Carried within the IS-IS IPv6 Reachability TLV (TLV 236/237) when the locator is also advertised as a prefix. Presence of this object causes the sub-TLV to be included; absence suppresses it. Reference: RFC 7794.
	PrefixAttributes() IsisSRv6PrefixAttributes
	// SetPrefixAttributes assigns IsisSRv6PrefixAttributes provided by user to IsisSRv6AdvertiseLocatorAsPrefix.
	// IsisSRv6PrefixAttributes is prefix Attribute Flags Sub-TLV (sub-TLV type 4, RFC 7794) for an SRv6 locator prefix advertisement. Carried within the IS-IS IPv6 Reachability TLV (TLV 236/237) when the locator is also advertised as a prefix. Presence of this object causes the sub-TLV to be included; absence suppresses it. Reference: RFC 7794.
	SetPrefixAttributes(value IsisSRv6PrefixAttributes) IsisSRv6AdvertiseLocatorAsPrefix
	// HasPrefixAttributes checks if PrefixAttributes has been set in IsisSRv6AdvertiseLocatorAsPrefix
	HasPrefixAttributes() bool
	setNil()
}

// Metric value used when advertising the locator prefix in the IPv6 Reachability TLV (TLV 236/237).
// RouteMetric returns a uint32
func (obj *isisSRv6AdvertiseLocatorAsPrefix) RouteMetric() uint32 {

	return *obj.obj.RouteMetric

}

// Metric value used when advertising the locator prefix in the IPv6 Reachability TLV (TLV 236/237).
// RouteMetric returns a uint32
func (obj *isisSRv6AdvertiseLocatorAsPrefix) HasRouteMetric() bool {
	return obj.obj.RouteMetric != nil
}

// Metric value used when advertising the locator prefix in the IPv6 Reachability TLV (TLV 236/237).
// SetRouteMetric sets the uint32 value in the IsisSRv6AdvertiseLocatorAsPrefix object
func (obj *isisSRv6AdvertiseLocatorAsPrefix) SetRouteMetric(value uint32) IsisSRv6AdvertiseLocatorAsPrefix {

	obj.obj.RouteMetric = &value
	return obj
}

type IsisSRv6AdvertiseLocatorAsPrefixRedistributionTypeEnum string

// Enum of RedistributionType on IsisSRv6AdvertiseLocatorAsPrefix
var IsisSRv6AdvertiseLocatorAsPrefixRedistributionType = struct {
	UP   IsisSRv6AdvertiseLocatorAsPrefixRedistributionTypeEnum
	DOWN IsisSRv6AdvertiseLocatorAsPrefixRedistributionTypeEnum
}{
	UP:   IsisSRv6AdvertiseLocatorAsPrefixRedistributionTypeEnum("up"),
	DOWN: IsisSRv6AdvertiseLocatorAsPrefixRedistributionTypeEnum("down"),
}

func (obj *isisSRv6AdvertiseLocatorAsPrefix) RedistributionType() IsisSRv6AdvertiseLocatorAsPrefixRedistributionTypeEnum {
	return IsisSRv6AdvertiseLocatorAsPrefixRedistributionTypeEnum(obj.obj.RedistributionType.Enum().String())
}

// Controls the Up/Down redistribution bit for this locator prefix when advertised in the IPv6 Reachability TLV (TLV 236/237). 'up' = normal advertisement (default); 'down' = the locator has been leaked from Level 2 to Level 1 (RFC 5305).
// RedistributionType returns a string
func (obj *isisSRv6AdvertiseLocatorAsPrefix) HasRedistributionType() bool {
	return obj.obj.RedistributionType != nil
}

func (obj *isisSRv6AdvertiseLocatorAsPrefix) SetRedistributionType(value IsisSRv6AdvertiseLocatorAsPrefixRedistributionTypeEnum) IsisSRv6AdvertiseLocatorAsPrefix {
	intValue, ok := otg.IsisSRv6AdvertiseLocatorAsPrefix_RedistributionType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisSRv6AdvertiseLocatorAsPrefixRedistributionTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisSRv6AdvertiseLocatorAsPrefix_RedistributionType_Enum(intValue)
	obj.obj.RedistributionType = &enumValue

	return obj
}

type IsisSRv6AdvertiseLocatorAsPrefixRouteOriginEnum string

// Enum of RouteOrigin on IsisSRv6AdvertiseLocatorAsPrefix
var IsisSRv6AdvertiseLocatorAsPrefixRouteOrigin = struct {
	INTERNAL IsisSRv6AdvertiseLocatorAsPrefixRouteOriginEnum
	EXTERNAL IsisSRv6AdvertiseLocatorAsPrefixRouteOriginEnum
}{
	INTERNAL: IsisSRv6AdvertiseLocatorAsPrefixRouteOriginEnum("internal"),
	EXTERNAL: IsisSRv6AdvertiseLocatorAsPrefixRouteOriginEnum("external"),
}

func (obj *isisSRv6AdvertiseLocatorAsPrefix) RouteOrigin() IsisSRv6AdvertiseLocatorAsPrefixRouteOriginEnum {
	return IsisSRv6AdvertiseLocatorAsPrefixRouteOriginEnum(obj.obj.RouteOrigin.Enum().String())
}

// Origin type for the locator prefix in the IPv6 Reachability TLV (TLV 236/237). 'internal' = intra-area prefix (default); 'external' = redistributed from another protocol.
// RouteOrigin returns a string
func (obj *isisSRv6AdvertiseLocatorAsPrefix) HasRouteOrigin() bool {
	return obj.obj.RouteOrigin != nil
}

func (obj *isisSRv6AdvertiseLocatorAsPrefix) SetRouteOrigin(value IsisSRv6AdvertiseLocatorAsPrefixRouteOriginEnum) IsisSRv6AdvertiseLocatorAsPrefix {
	intValue, ok := otg.IsisSRv6AdvertiseLocatorAsPrefix_RouteOrigin_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisSRv6AdvertiseLocatorAsPrefixRouteOriginEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisSRv6AdvertiseLocatorAsPrefix_RouteOrigin_Enum(intValue)
	obj.obj.RouteOrigin = &enumValue

	return obj
}

// When present, the Prefix Attribute Flags Sub-TLV (sub-TLV type 4, RFC 7794) is included in this prefix advertisement, carrying the X, R, N, and A flags. Absence suppresses the sub-TLV.
// PrefixAttributes returns a IsisSRv6PrefixAttributes
func (obj *isisSRv6AdvertiseLocatorAsPrefix) PrefixAttributes() IsisSRv6PrefixAttributes {
	if obj.obj.PrefixAttributes == nil {
		obj.obj.PrefixAttributes = NewIsisSRv6PrefixAttributes().msg()
	}
	if obj.prefixAttributesHolder == nil {
		obj.prefixAttributesHolder = &isisSRv6PrefixAttributes{obj: obj.obj.PrefixAttributes}
	}
	return obj.prefixAttributesHolder
}

// When present, the Prefix Attribute Flags Sub-TLV (sub-TLV type 4, RFC 7794) is included in this prefix advertisement, carrying the X, R, N, and A flags. Absence suppresses the sub-TLV.
// PrefixAttributes returns a IsisSRv6PrefixAttributes
func (obj *isisSRv6AdvertiseLocatorAsPrefix) HasPrefixAttributes() bool {
	return obj.obj.PrefixAttributes != nil
}

// When present, the Prefix Attribute Flags Sub-TLV (sub-TLV type 4, RFC 7794) is included in this prefix advertisement, carrying the X, R, N, and A flags. Absence suppresses the sub-TLV.
// SetPrefixAttributes sets the IsisSRv6PrefixAttributes value in the IsisSRv6AdvertiseLocatorAsPrefix object
func (obj *isisSRv6AdvertiseLocatorAsPrefix) SetPrefixAttributes(value IsisSRv6PrefixAttributes) IsisSRv6AdvertiseLocatorAsPrefix {

	obj.prefixAttributesHolder = nil
	obj.obj.PrefixAttributes = value.msg()

	return obj
}

func (obj *isisSRv6AdvertiseLocatorAsPrefix) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RouteMetric != nil {

		if *obj.obj.RouteMetric > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6AdvertiseLocatorAsPrefix.RouteMetric <= 16777215 but Got %d", *obj.obj.RouteMetric))
		}

	}

	if obj.obj.PrefixAttributes != nil {

		obj.PrefixAttributes().validateObj(vObj, set_default)
	}

}

func (obj *isisSRv6AdvertiseLocatorAsPrefix) setDefault() {
	if obj.obj.RouteMetric == nil {
		obj.SetRouteMetric(0)
	}
	if obj.obj.RedistributionType == nil {
		obj.SetRedistributionType(IsisSRv6AdvertiseLocatorAsPrefixRedistributionType.UP)

	}
	if obj.obj.RouteOrigin == nil {
		obj.SetRouteOrigin(IsisSRv6AdvertiseLocatorAsPrefixRouteOrigin.INTERNAL)

	}

}

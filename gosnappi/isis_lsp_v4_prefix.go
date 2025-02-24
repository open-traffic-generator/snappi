package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspV4Prefix *****
type isisLspV4Prefix struct {
	validation
	obj          *otg.IsisLspV4Prefix
	marshaller   marshalIsisLspV4Prefix
	unMarshaller unMarshalIsisLspV4Prefix
}

func NewIsisLspV4Prefix() IsisLspV4Prefix {
	obj := isisLspV4Prefix{obj: &otg.IsisLspV4Prefix{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspV4Prefix) msg() *otg.IsisLspV4Prefix {
	return obj.obj
}

func (obj *isisLspV4Prefix) setMsg(msg *otg.IsisLspV4Prefix) IsisLspV4Prefix {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspV4Prefix struct {
	obj *isisLspV4Prefix
}

type marshalIsisLspV4Prefix interface {
	// ToProto marshals IsisLspV4Prefix to protobuf object *otg.IsisLspV4Prefix
	ToProto() (*otg.IsisLspV4Prefix, error)
	// ToPbText marshals IsisLspV4Prefix to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspV4Prefix to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspV4Prefix to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisLspV4Prefix to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisLspV4Prefix struct {
	obj *isisLspV4Prefix
}

type unMarshalIsisLspV4Prefix interface {
	// FromProto unmarshals IsisLspV4Prefix from protobuf object *otg.IsisLspV4Prefix
	FromProto(msg *otg.IsisLspV4Prefix) (IsisLspV4Prefix, error)
	// FromPbText unmarshals IsisLspV4Prefix from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspV4Prefix from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspV4Prefix from JSON text
	FromJson(value string) error
}

func (obj *isisLspV4Prefix) Marshal() marshalIsisLspV4Prefix {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspV4Prefix{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspV4Prefix) Unmarshal() unMarshalIsisLspV4Prefix {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspV4Prefix{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspV4Prefix) ToProto() (*otg.IsisLspV4Prefix, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspV4Prefix) FromProto(msg *otg.IsisLspV4Prefix) (IsisLspV4Prefix, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspV4Prefix) ToPbText() (string, error) {
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

func (m *unMarshalisisLspV4Prefix) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalisisLspV4Prefix) ToYaml() (string, error) {
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

func (m *unMarshalisisLspV4Prefix) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalisisLspV4Prefix) ToJsonRaw() (string, error) {
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

func (m *marshalisisLspV4Prefix) ToJson() (string, error) {
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

func (m *unMarshalisisLspV4Prefix) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *isisLspV4Prefix) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspV4Prefix) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspV4Prefix) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspV4Prefix) Clone() (IsisLspV4Prefix, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspV4Prefix()
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

// IsisLspV4Prefix is this group defines attributes of an IPv4 standard prefix.
type IsisLspV4Prefix interface {
	Validation
	// msg marshals IsisLspV4Prefix to protobuf object *otg.IsisLspV4Prefix
	// and doesn't set defaults
	msg() *otg.IsisLspV4Prefix
	// setMsg unmarshals IsisLspV4Prefix from protobuf object *otg.IsisLspV4Prefix
	// and doesn't set defaults
	setMsg(*otg.IsisLspV4Prefix) IsisLspV4Prefix
	// provides marshal interface
	Marshal() marshalIsisLspV4Prefix
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspV4Prefix
	// validate validates IsisLspV4Prefix
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspV4Prefix, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Address returns string, set in IsisLspV4Prefix.
	Ipv4Address() string
	// SetIpv4Address assigns string provided by user to IsisLspV4Prefix
	SetIpv4Address(value string) IsisLspV4Prefix
	// HasIpv4Address checks if Ipv4Address has been set in IsisLspV4Prefix
	HasIpv4Address() bool
	// PrefixLength returns uint32, set in IsisLspV4Prefix.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to IsisLspV4Prefix
	SetPrefixLength(value uint32) IsisLspV4Prefix
	// HasPrefixLength checks if PrefixLength has been set in IsisLspV4Prefix
	HasPrefixLength() bool
	// RedistributionType returns IsisLspV4PrefixRedistributionTypeEnum, set in IsisLspV4Prefix
	RedistributionType() IsisLspV4PrefixRedistributionTypeEnum
	// SetRedistributionType assigns IsisLspV4PrefixRedistributionTypeEnum provided by user to IsisLspV4Prefix
	SetRedistributionType(value IsisLspV4PrefixRedistributionTypeEnum) IsisLspV4Prefix
	// HasRedistributionType checks if RedistributionType has been set in IsisLspV4Prefix
	HasRedistributionType() bool
	// DefaultMetric returns uint32, set in IsisLspV4Prefix.
	DefaultMetric() uint32
	// SetDefaultMetric assigns uint32 provided by user to IsisLspV4Prefix
	SetDefaultMetric(value uint32) IsisLspV4Prefix
	// HasDefaultMetric checks if DefaultMetric has been set in IsisLspV4Prefix
	HasDefaultMetric() bool
	// OriginType returns IsisLspV4PrefixOriginTypeEnum, set in IsisLspV4Prefix
	OriginType() IsisLspV4PrefixOriginTypeEnum
	// SetOriginType assigns IsisLspV4PrefixOriginTypeEnum provided by user to IsisLspV4Prefix
	SetOriginType(value IsisLspV4PrefixOriginTypeEnum) IsisLspV4Prefix
	// HasOriginType checks if OriginType has been set in IsisLspV4Prefix
	HasOriginType() bool
}

// An IPv4 unicast prefix reachable via the originator of this LSP.
// Ipv4Address returns a string
func (obj *isisLspV4Prefix) Ipv4Address() string {

	return *obj.obj.Ipv4Address

}

// An IPv4 unicast prefix reachable via the originator of this LSP.
// Ipv4Address returns a string
func (obj *isisLspV4Prefix) HasIpv4Address() bool {
	return obj.obj.Ipv4Address != nil
}

// An IPv4 unicast prefix reachable via the originator of this LSP.
// SetIpv4Address sets the string value in the IsisLspV4Prefix object
func (obj *isisLspV4Prefix) SetIpv4Address(value string) IsisLspV4Prefix {

	obj.obj.Ipv4Address = &value
	return obj
}

// The length of the IPv4 prefix.
// PrefixLength returns a uint32
func (obj *isisLspV4Prefix) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The length of the IPv4 prefix.
// PrefixLength returns a uint32
func (obj *isisLspV4Prefix) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The length of the IPv4 prefix.
// SetPrefixLength sets the uint32 value in the IsisLspV4Prefix object
func (obj *isisLspV4Prefix) SetPrefixLength(value uint32) IsisLspV4Prefix {

	obj.obj.PrefixLength = &value
	return obj
}

type IsisLspV4PrefixRedistributionTypeEnum string

// Enum of RedistributionType on IsisLspV4Prefix
var IsisLspV4PrefixRedistributionType = struct {
	UP   IsisLspV4PrefixRedistributionTypeEnum
	DOWN IsisLspV4PrefixRedistributionTypeEnum
}{
	UP:   IsisLspV4PrefixRedistributionTypeEnum("up"),
	DOWN: IsisLspV4PrefixRedistributionTypeEnum("down"),
}

func (obj *isisLspV4Prefix) RedistributionType() IsisLspV4PrefixRedistributionTypeEnum {
	return IsisLspV4PrefixRedistributionTypeEnum(obj.obj.RedistributionType.Enum().String())
}

// Up (0)-used when a prefix is initially advertised within the ISIS L3 hierarchy,
// and for all other prefixes in L1 and L2 LSPs. (default)
// Down (1)-used when an L1/L2 router advertises L2 prefixes in L1 LSPs.
// The prefixes are being advertised from a higher level (L2) down to a lower level (L1).
// RedistributionType returns a string
func (obj *isisLspV4Prefix) HasRedistributionType() bool {
	return obj.obj.RedistributionType != nil
}

func (obj *isisLspV4Prefix) SetRedistributionType(value IsisLspV4PrefixRedistributionTypeEnum) IsisLspV4Prefix {
	intValue, ok := otg.IsisLspV4Prefix_RedistributionType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisLspV4PrefixRedistributionTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisLspV4Prefix_RedistributionType_Enum(intValue)
	obj.obj.RedistributionType = &enumValue

	return obj
}

// ISIS default metric value.
// DefaultMetric returns a uint32
func (obj *isisLspV4Prefix) DefaultMetric() uint32 {

	return *obj.obj.DefaultMetric

}

// ISIS default metric value.
// DefaultMetric returns a uint32
func (obj *isisLspV4Prefix) HasDefaultMetric() bool {
	return obj.obj.DefaultMetric != nil
}

// ISIS default metric value.
// SetDefaultMetric sets the uint32 value in the IsisLspV4Prefix object
func (obj *isisLspV4Prefix) SetDefaultMetric(value uint32) IsisLspV4Prefix {

	obj.obj.DefaultMetric = &value
	return obj
}

type IsisLspV4PrefixOriginTypeEnum string

// Enum of OriginType on IsisLspV4Prefix
var IsisLspV4PrefixOriginType = struct {
	INTERNAL IsisLspV4PrefixOriginTypeEnum
	EXTERNAL IsisLspV4PrefixOriginTypeEnum
}{
	INTERNAL: IsisLspV4PrefixOriginTypeEnum("internal"),
	EXTERNAL: IsisLspV4PrefixOriginTypeEnum("external"),
}

func (obj *isisLspV4Prefix) OriginType() IsisLspV4PrefixOriginTypeEnum {
	return IsisLspV4PrefixOriginTypeEnum(obj.obj.OriginType.Enum().String())
}

// The origin of the advertised route-internal or external to the ISIS area. Options include the following:
// Internal-for intra-area routes, through Level 1 LSPs.
// External-for inter-area routes redistributed within L1, through Level
// 1 LSPs.
// OriginType returns a string
func (obj *isisLspV4Prefix) HasOriginType() bool {
	return obj.obj.OriginType != nil
}

func (obj *isisLspV4Prefix) SetOriginType(value IsisLspV4PrefixOriginTypeEnum) IsisLspV4Prefix {
	intValue, ok := otg.IsisLspV4Prefix_OriginType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisLspV4PrefixOriginTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisLspV4Prefix_OriginType_Enum(intValue)
	obj.obj.OriginType = &enumValue

	return obj
}

func (obj *isisLspV4Prefix) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisLspV4Prefix.PrefixLength <= 32 but Got %d", *obj.obj.PrefixLength))
		}

	}

}

func (obj *isisLspV4Prefix) setDefault() {

}

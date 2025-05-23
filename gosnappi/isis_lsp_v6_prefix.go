package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspV6Prefix *****
type isisLspV6Prefix struct {
	validation
	obj                    *otg.IsisLspV6Prefix
	marshaller             marshalIsisLspV6Prefix
	unMarshaller           unMarshalIsisLspV6Prefix
	prefixAttributesHolder IsisLspPrefixAttributes
	prefixSidsHolder       IsisLspV6PrefixIsisLspPrefixSidIter
}

func NewIsisLspV6Prefix() IsisLspV6Prefix {
	obj := isisLspV6Prefix{obj: &otg.IsisLspV6Prefix{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspV6Prefix) msg() *otg.IsisLspV6Prefix {
	return obj.obj
}

func (obj *isisLspV6Prefix) setMsg(msg *otg.IsisLspV6Prefix) IsisLspV6Prefix {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspV6Prefix struct {
	obj *isisLspV6Prefix
}

type marshalIsisLspV6Prefix interface {
	// ToProto marshals IsisLspV6Prefix to protobuf object *otg.IsisLspV6Prefix
	ToProto() (*otg.IsisLspV6Prefix, error)
	// ToPbText marshals IsisLspV6Prefix to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspV6Prefix to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspV6Prefix to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisLspV6Prefix to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisLspV6Prefix struct {
	obj *isisLspV6Prefix
}

type unMarshalIsisLspV6Prefix interface {
	// FromProto unmarshals IsisLspV6Prefix from protobuf object *otg.IsisLspV6Prefix
	FromProto(msg *otg.IsisLspV6Prefix) (IsisLspV6Prefix, error)
	// FromPbText unmarshals IsisLspV6Prefix from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspV6Prefix from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspV6Prefix from JSON text
	FromJson(value string) error
}

func (obj *isisLspV6Prefix) Marshal() marshalIsisLspV6Prefix {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspV6Prefix{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspV6Prefix) Unmarshal() unMarshalIsisLspV6Prefix {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspV6Prefix{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspV6Prefix) ToProto() (*otg.IsisLspV6Prefix, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspV6Prefix) FromProto(msg *otg.IsisLspV6Prefix) (IsisLspV6Prefix, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspV6Prefix) ToPbText() (string, error) {
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

func (m *unMarshalisisLspV6Prefix) FromPbText(value string) error {
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

func (m *marshalisisLspV6Prefix) ToYaml() (string, error) {
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

func (m *unMarshalisisLspV6Prefix) FromYaml(value string) error {
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

func (m *marshalisisLspV6Prefix) ToJsonRaw() (string, error) {
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

func (m *marshalisisLspV6Prefix) ToJson() (string, error) {
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

func (m *unMarshalisisLspV6Prefix) FromJson(value string) error {
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

func (obj *isisLspV6Prefix) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspV6Prefix) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspV6Prefix) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspV6Prefix) Clone() (IsisLspV6Prefix, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspV6Prefix()
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

func (obj *isisLspV6Prefix) setNil() {
	obj.prefixAttributesHolder = nil
	obj.prefixSidsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspV6Prefix is it defines attributes of an IPv6 standard prefix.
type IsisLspV6Prefix interface {
	Validation
	// msg marshals IsisLspV6Prefix to protobuf object *otg.IsisLspV6Prefix
	// and doesn't set defaults
	msg() *otg.IsisLspV6Prefix
	// setMsg unmarshals IsisLspV6Prefix from protobuf object *otg.IsisLspV6Prefix
	// and doesn't set defaults
	setMsg(*otg.IsisLspV6Prefix) IsisLspV6Prefix
	// provides marshal interface
	Marshal() marshalIsisLspV6Prefix
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspV6Prefix
	// validate validates IsisLspV6Prefix
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspV6Prefix, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv6Address returns string, set in IsisLspV6Prefix.
	Ipv6Address() string
	// SetIpv6Address assigns string provided by user to IsisLspV6Prefix
	SetIpv6Address(value string) IsisLspV6Prefix
	// HasIpv6Address checks if Ipv6Address has been set in IsisLspV6Prefix
	HasIpv6Address() bool
	// PrefixLength returns uint32, set in IsisLspV6Prefix.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to IsisLspV6Prefix
	SetPrefixLength(value uint32) IsisLspV6Prefix
	// HasPrefixLength checks if PrefixLength has been set in IsisLspV6Prefix
	HasPrefixLength() bool
	// Metric returns uint32, set in IsisLspV6Prefix.
	Metric() uint32
	// SetMetric assigns uint32 provided by user to IsisLspV6Prefix
	SetMetric(value uint32) IsisLspV6Prefix
	// HasMetric checks if Metric has been set in IsisLspV6Prefix
	HasMetric() bool
	// RedistributionType returns IsisLspV6PrefixRedistributionTypeEnum, set in IsisLspV6Prefix
	RedistributionType() IsisLspV6PrefixRedistributionTypeEnum
	// SetRedistributionType assigns IsisLspV6PrefixRedistributionTypeEnum provided by user to IsisLspV6Prefix
	SetRedistributionType(value IsisLspV6PrefixRedistributionTypeEnum) IsisLspV6Prefix
	// HasRedistributionType checks if RedistributionType has been set in IsisLspV6Prefix
	HasRedistributionType() bool
	// OriginType returns IsisLspV6PrefixOriginTypeEnum, set in IsisLspV6Prefix
	OriginType() IsisLspV6PrefixOriginTypeEnum
	// SetOriginType assigns IsisLspV6PrefixOriginTypeEnum provided by user to IsisLspV6Prefix
	SetOriginType(value IsisLspV6PrefixOriginTypeEnum) IsisLspV6Prefix
	// HasOriginType checks if OriginType has been set in IsisLspV6Prefix
	HasOriginType() bool
	// PrefixAttributes returns IsisLspPrefixAttributes, set in IsisLspV6Prefix.
	// IsisLspPrefixAttributes is one bit value of ISIS Prefix attributes for the extended IPv4 and IPv6 reachability. https://www.rfc-editor.org/rfc/rfc7794.html.
	PrefixAttributes() IsisLspPrefixAttributes
	// SetPrefixAttributes assigns IsisLspPrefixAttributes provided by user to IsisLspV6Prefix.
	// IsisLspPrefixAttributes is one bit value of ISIS Prefix attributes for the extended IPv4 and IPv6 reachability. https://www.rfc-editor.org/rfc/rfc7794.html.
	SetPrefixAttributes(value IsisLspPrefixAttributes) IsisLspV6Prefix
	// HasPrefixAttributes checks if PrefixAttributes has been set in IsisLspV6Prefix
	HasPrefixAttributes() bool
	// PrefixSids returns IsisLspV6PrefixIsisLspPrefixSidIterIter, set in IsisLspV6Prefix
	PrefixSids() IsisLspV6PrefixIsisLspPrefixSidIter
	setNil()
}

// An IPv6 unicast prefix reachable via the originator of this LSP.
// Ipv6Address returns a string
func (obj *isisLspV6Prefix) Ipv6Address() string {

	return *obj.obj.Ipv6Address

}

// An IPv6 unicast prefix reachable via the originator of this LSP.
// Ipv6Address returns a string
func (obj *isisLspV6Prefix) HasIpv6Address() bool {
	return obj.obj.Ipv6Address != nil
}

// An IPv6 unicast prefix reachable via the originator of this LSP.
// SetIpv6Address sets the string value in the IsisLspV6Prefix object
func (obj *isisLspV6Prefix) SetIpv6Address(value string) IsisLspV6Prefix {

	obj.obj.Ipv6Address = &value
	return obj
}

// The length of the IPv6 prefix.
// PrefixLength returns a uint32
func (obj *isisLspV6Prefix) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The length of the IPv6 prefix.
// PrefixLength returns a uint32
func (obj *isisLspV6Prefix) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The length of the IPv6 prefix.
// SetPrefixLength sets the uint32 value in the IsisLspV6Prefix object
func (obj *isisLspV6Prefix) SetPrefixLength(value uint32) IsisLspV6Prefix {

	obj.obj.PrefixLength = &value
	return obj
}

// ISIS wide metric.
// Metric returns a uint32
func (obj *isisLspV6Prefix) Metric() uint32 {

	return *obj.obj.Metric

}

// ISIS wide metric.
// Metric returns a uint32
func (obj *isisLspV6Prefix) HasMetric() bool {
	return obj.obj.Metric != nil
}

// ISIS wide metric.
// SetMetric sets the uint32 value in the IsisLspV6Prefix object
func (obj *isisLspV6Prefix) SetMetric(value uint32) IsisLspV6Prefix {

	obj.obj.Metric = &value
	return obj
}

type IsisLspV6PrefixRedistributionTypeEnum string

// Enum of RedistributionType on IsisLspV6Prefix
var IsisLspV6PrefixRedistributionType = struct {
	UP   IsisLspV6PrefixRedistributionTypeEnum
	DOWN IsisLspV6PrefixRedistributionTypeEnum
}{
	UP:   IsisLspV6PrefixRedistributionTypeEnum("up"),
	DOWN: IsisLspV6PrefixRedistributionTypeEnum("down"),
}

func (obj *isisLspV6Prefix) RedistributionType() IsisLspV6PrefixRedistributionTypeEnum {
	return IsisLspV6PrefixRedistributionTypeEnum(obj.obj.RedistributionType.Enum().String())
}

// Up (0)-used when a prefix is initially advertised within the ISIS L3 hierarchy,
// and for all other prefixes in L1 and L2 LSPs. (default)
// Down (1)-used when an L1/L2 router advertises L2 prefixes in L1 LSPs.
// The prefixes are being advertised from a higher level (L2) down to a lower level (L1).
// RedistributionType returns a string
func (obj *isisLspV6Prefix) HasRedistributionType() bool {
	return obj.obj.RedistributionType != nil
}

func (obj *isisLspV6Prefix) SetRedistributionType(value IsisLspV6PrefixRedistributionTypeEnum) IsisLspV6Prefix {
	intValue, ok := otg.IsisLspV6Prefix_RedistributionType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisLspV6PrefixRedistributionTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisLspV6Prefix_RedistributionType_Enum(intValue)
	obj.obj.RedistributionType = &enumValue

	return obj
}

type IsisLspV6PrefixOriginTypeEnum string

// Enum of OriginType on IsisLspV6Prefix
var IsisLspV6PrefixOriginType = struct {
	INTERNAL IsisLspV6PrefixOriginTypeEnum
	EXTERNAL IsisLspV6PrefixOriginTypeEnum
}{
	INTERNAL: IsisLspV6PrefixOriginTypeEnum("internal"),
	EXTERNAL: IsisLspV6PrefixOriginTypeEnum("external"),
}

func (obj *isisLspV6Prefix) OriginType() IsisLspV6PrefixOriginTypeEnum {
	return IsisLspV6PrefixOriginTypeEnum(obj.obj.OriginType.Enum().String())
}

// The origin of the advertised route-internal or external to the ISIS area. Options include the following:
// Internal-for intra-area routes, through Level 1 LSPs.
// External-for inter-area routes redistributed within L1, through Level
// 1 LSPs.
// OriginType returns a string
func (obj *isisLspV6Prefix) HasOriginType() bool {
	return obj.obj.OriginType != nil
}

func (obj *isisLspV6Prefix) SetOriginType(value IsisLspV6PrefixOriginTypeEnum) IsisLspV6Prefix {
	intValue, ok := otg.IsisLspV6Prefix_OriginType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisLspV6PrefixOriginTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisLspV6Prefix_OriginType_Enum(intValue)
	obj.obj.OriginType = &enumValue

	return obj
}

// Extended Prefix Attribute flags container sub-TLV is type 4.
// PrefixAttributes returns a IsisLspPrefixAttributes
func (obj *isisLspV6Prefix) PrefixAttributes() IsisLspPrefixAttributes {
	if obj.obj.PrefixAttributes == nil {
		obj.obj.PrefixAttributes = NewIsisLspPrefixAttributes().msg()
	}
	if obj.prefixAttributesHolder == nil {
		obj.prefixAttributesHolder = &isisLspPrefixAttributes{obj: obj.obj.PrefixAttributes}
	}
	return obj.prefixAttributesHolder
}

// Extended Prefix Attribute flags container sub-TLV is type 4.
// PrefixAttributes returns a IsisLspPrefixAttributes
func (obj *isisLspV6Prefix) HasPrefixAttributes() bool {
	return obj.obj.PrefixAttributes != nil
}

// Extended Prefix Attribute flags container sub-TLV is type 4.
// SetPrefixAttributes sets the IsisLspPrefixAttributes value in the IsisLspV6Prefix object
func (obj *isisLspV6Prefix) SetPrefixAttributes(value IsisLspPrefixAttributes) IsisLspV6Prefix {

	obj.prefixAttributesHolder = nil
	obj.obj.PrefixAttributes = value.msg()

	return obj
}

// Prefix Segment-ID list. IGP-Prefix Segment is an IGP segment attached to an IGP prefix. An IGP-Prefix Segment is global
// (unless explicitly advertised otherwise) within the SR/IGP domain.
// PrefixSids returns a []IsisLspPrefixSid
func (obj *isisLspV6Prefix) PrefixSids() IsisLspV6PrefixIsisLspPrefixSidIter {
	if len(obj.obj.PrefixSids) == 0 {
		obj.obj.PrefixSids = []*otg.IsisLspPrefixSid{}
	}
	if obj.prefixSidsHolder == nil {
		obj.prefixSidsHolder = newIsisLspV6PrefixIsisLspPrefixSidIter(&obj.obj.PrefixSids).setMsg(obj)
	}
	return obj.prefixSidsHolder
}

type isisLspV6PrefixIsisLspPrefixSidIter struct {
	obj                   *isisLspV6Prefix
	isisLspPrefixSidSlice []IsisLspPrefixSid
	fieldPtr              *[]*otg.IsisLspPrefixSid
}

func newIsisLspV6PrefixIsisLspPrefixSidIter(ptr *[]*otg.IsisLspPrefixSid) IsisLspV6PrefixIsisLspPrefixSidIter {
	return &isisLspV6PrefixIsisLspPrefixSidIter{fieldPtr: ptr}
}

type IsisLspV6PrefixIsisLspPrefixSidIter interface {
	setMsg(*isisLspV6Prefix) IsisLspV6PrefixIsisLspPrefixSidIter
	Items() []IsisLspPrefixSid
	Add() IsisLspPrefixSid
	Append(items ...IsisLspPrefixSid) IsisLspV6PrefixIsisLspPrefixSidIter
	Set(index int, newObj IsisLspPrefixSid) IsisLspV6PrefixIsisLspPrefixSidIter
	Clear() IsisLspV6PrefixIsisLspPrefixSidIter
	clearHolderSlice() IsisLspV6PrefixIsisLspPrefixSidIter
	appendHolderSlice(item IsisLspPrefixSid) IsisLspV6PrefixIsisLspPrefixSidIter
}

func (obj *isisLspV6PrefixIsisLspPrefixSidIter) setMsg(msg *isisLspV6Prefix) IsisLspV6PrefixIsisLspPrefixSidIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspPrefixSid{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspV6PrefixIsisLspPrefixSidIter) Items() []IsisLspPrefixSid {
	return obj.isisLspPrefixSidSlice
}

func (obj *isisLspV6PrefixIsisLspPrefixSidIter) Add() IsisLspPrefixSid {
	newObj := &otg.IsisLspPrefixSid{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspPrefixSid{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspPrefixSidSlice = append(obj.isisLspPrefixSidSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspV6PrefixIsisLspPrefixSidIter) Append(items ...IsisLspPrefixSid) IsisLspV6PrefixIsisLspPrefixSidIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspPrefixSidSlice = append(obj.isisLspPrefixSidSlice, item)
	}
	return obj
}

func (obj *isisLspV6PrefixIsisLspPrefixSidIter) Set(index int, newObj IsisLspPrefixSid) IsisLspV6PrefixIsisLspPrefixSidIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspPrefixSidSlice[index] = newObj
	return obj
}
func (obj *isisLspV6PrefixIsisLspPrefixSidIter) Clear() IsisLspV6PrefixIsisLspPrefixSidIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspPrefixSid{}
		obj.isisLspPrefixSidSlice = []IsisLspPrefixSid{}
	}
	return obj
}
func (obj *isisLspV6PrefixIsisLspPrefixSidIter) clearHolderSlice() IsisLspV6PrefixIsisLspPrefixSidIter {
	if len(obj.isisLspPrefixSidSlice) > 0 {
		obj.isisLspPrefixSidSlice = []IsisLspPrefixSid{}
	}
	return obj
}
func (obj *isisLspV6PrefixIsisLspPrefixSidIter) appendHolderSlice(item IsisLspPrefixSid) IsisLspV6PrefixIsisLspPrefixSidIter {
	obj.isisLspPrefixSidSlice = append(obj.isisLspPrefixSidSlice, item)
	return obj
}

func (obj *isisLspV6Prefix) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv6Address != nil {

		err := obj.validateIpv6(obj.Ipv6Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on IsisLspV6Prefix.Ipv6Address"))
		}

	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisLspV6Prefix.PrefixLength <= 128 but Got %d", *obj.obj.PrefixLength))
		}

	}

	if obj.obj.PrefixAttributes != nil {

		obj.PrefixAttributes().validateObj(vObj, set_default)
	}

	if len(obj.obj.PrefixSids) != 0 {

		if set_default {
			obj.PrefixSids().clearHolderSlice()
			for _, item := range obj.obj.PrefixSids {
				obj.PrefixSids().appendHolderSlice(&isisLspPrefixSid{obj: item})
			}
		}
		for _, item := range obj.PrefixSids().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *isisLspV6Prefix) setDefault() {

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisV4RouteRange *****
type isisV4RouteRange struct {
	validation
	obj              *otg.IsisV4RouteRange
	marshaller       marshalIsisV4RouteRange
	unMarshaller     unMarshalIsisV4RouteRange
	addressesHolder  IsisV4RouteRangeV4RouteAddressIter
	prefixSidsHolder IsisV4RouteRangeIsisSRPrefixSidIter
}

func NewIsisV4RouteRange() IsisV4RouteRange {
	obj := isisV4RouteRange{obj: &otg.IsisV4RouteRange{}}
	obj.setDefault()
	return &obj
}

func (obj *isisV4RouteRange) msg() *otg.IsisV4RouteRange {
	return obj.obj
}

func (obj *isisV4RouteRange) setMsg(msg *otg.IsisV4RouteRange) IsisV4RouteRange {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisV4RouteRange struct {
	obj *isisV4RouteRange
}

type marshalIsisV4RouteRange interface {
	// ToProto marshals IsisV4RouteRange to protobuf object *otg.IsisV4RouteRange
	ToProto() (*otg.IsisV4RouteRange, error)
	// ToPbText marshals IsisV4RouteRange to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisV4RouteRange to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisV4RouteRange to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisV4RouteRange to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisV4RouteRange struct {
	obj *isisV4RouteRange
}

type unMarshalIsisV4RouteRange interface {
	// FromProto unmarshals IsisV4RouteRange from protobuf object *otg.IsisV4RouteRange
	FromProto(msg *otg.IsisV4RouteRange) (IsisV4RouteRange, error)
	// FromPbText unmarshals IsisV4RouteRange from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisV4RouteRange from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisV4RouteRange from JSON text
	FromJson(value string) error
}

func (obj *isisV4RouteRange) Marshal() marshalIsisV4RouteRange {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisV4RouteRange{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisV4RouteRange) Unmarshal() unMarshalIsisV4RouteRange {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisV4RouteRange{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisV4RouteRange) ToProto() (*otg.IsisV4RouteRange, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisV4RouteRange) FromProto(msg *otg.IsisV4RouteRange) (IsisV4RouteRange, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisV4RouteRange) ToPbText() (string, error) {
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

func (m *unMarshalisisV4RouteRange) FromPbText(value string) error {
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

func (m *marshalisisV4RouteRange) ToYaml() (string, error) {
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

func (m *unMarshalisisV4RouteRange) FromYaml(value string) error {
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

func (m *marshalisisV4RouteRange) ToJsonRaw() (string, error) {
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

func (m *marshalisisV4RouteRange) ToJson() (string, error) {
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

func (m *unMarshalisisV4RouteRange) FromJson(value string) error {
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

func (obj *isisV4RouteRange) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisV4RouteRange) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisV4RouteRange) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisV4RouteRange) Clone() (IsisV4RouteRange, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisV4RouteRange()
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

func (obj *isisV4RouteRange) setNil() {
	obj.addressesHolder = nil
	obj.prefixSidsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisV4RouteRange is emulated ISIS IPv4 routes.
type IsisV4RouteRange interface {
	Validation
	// msg marshals IsisV4RouteRange to protobuf object *otg.IsisV4RouteRange
	// and doesn't set defaults
	msg() *otg.IsisV4RouteRange
	// setMsg unmarshals IsisV4RouteRange from protobuf object *otg.IsisV4RouteRange
	// and doesn't set defaults
	setMsg(*otg.IsisV4RouteRange) IsisV4RouteRange
	// provides marshal interface
	Marshal() marshalIsisV4RouteRange
	// provides unmarshal interface
	Unmarshal() unMarshalIsisV4RouteRange
	// validate validates IsisV4RouteRange
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisV4RouteRange, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Addresses returns IsisV4RouteRangeV4RouteAddressIterIter, set in IsisV4RouteRange
	Addresses() IsisV4RouteRangeV4RouteAddressIter
	// LinkMetric returns uint32, set in IsisV4RouteRange.
	LinkMetric() uint32
	// SetLinkMetric assigns uint32 provided by user to IsisV4RouteRange
	SetLinkMetric(value uint32) IsisV4RouteRange
	// HasLinkMetric checks if LinkMetric has been set in IsisV4RouteRange
	HasLinkMetric() bool
	// OriginType returns IsisV4RouteRangeOriginTypeEnum, set in IsisV4RouteRange
	OriginType() IsisV4RouteRangeOriginTypeEnum
	// SetOriginType assigns IsisV4RouteRangeOriginTypeEnum provided by user to IsisV4RouteRange
	SetOriginType(value IsisV4RouteRangeOriginTypeEnum) IsisV4RouteRange
	// HasOriginType checks if OriginType has been set in IsisV4RouteRange
	HasOriginType() bool
	// RedistributionType returns IsisV4RouteRangeRedistributionTypeEnum, set in IsisV4RouteRange
	RedistributionType() IsisV4RouteRangeRedistributionTypeEnum
	// SetRedistributionType assigns IsisV4RouteRangeRedistributionTypeEnum provided by user to IsisV4RouteRange
	SetRedistributionType(value IsisV4RouteRangeRedistributionTypeEnum) IsisV4RouteRange
	// HasRedistributionType checks if RedistributionType has been set in IsisV4RouteRange
	HasRedistributionType() bool
	// Name returns string, set in IsisV4RouteRange.
	Name() string
	// SetName assigns string provided by user to IsisV4RouteRange
	SetName(value string) IsisV4RouteRange
	// PrefixAttrEnabled returns bool, set in IsisV4RouteRange.
	PrefixAttrEnabled() bool
	// SetPrefixAttrEnabled assigns bool provided by user to IsisV4RouteRange
	SetPrefixAttrEnabled(value bool) IsisV4RouteRange
	// HasPrefixAttrEnabled checks if PrefixAttrEnabled has been set in IsisV4RouteRange
	HasPrefixAttrEnabled() bool
	// XFlag returns bool, set in IsisV4RouteRange.
	XFlag() bool
	// SetXFlag assigns bool provided by user to IsisV4RouteRange
	SetXFlag(value bool) IsisV4RouteRange
	// HasXFlag checks if XFlag has been set in IsisV4RouteRange
	HasXFlag() bool
	// RFlag returns bool, set in IsisV4RouteRange.
	RFlag() bool
	// SetRFlag assigns bool provided by user to IsisV4RouteRange
	SetRFlag(value bool) IsisV4RouteRange
	// HasRFlag checks if RFlag has been set in IsisV4RouteRange
	HasRFlag() bool
	// NFlag returns bool, set in IsisV4RouteRange.
	NFlag() bool
	// SetNFlag assigns bool provided by user to IsisV4RouteRange
	SetNFlag(value bool) IsisV4RouteRange
	// HasNFlag checks if NFlag has been set in IsisV4RouteRange
	HasNFlag() bool
	// PrefixSids returns IsisV4RouteRangeIsisSRPrefixSidIterIter, set in IsisV4RouteRange
	PrefixSids() IsisV4RouteRangeIsisSRPrefixSidIter
	setNil()
}

// A list of group of IPv4 route addresses.
// Addresses returns a []V4RouteAddress
func (obj *isisV4RouteRange) Addresses() IsisV4RouteRangeV4RouteAddressIter {
	if len(obj.obj.Addresses) == 0 {
		obj.obj.Addresses = []*otg.V4RouteAddress{}
	}
	if obj.addressesHolder == nil {
		obj.addressesHolder = newIsisV4RouteRangeV4RouteAddressIter(&obj.obj.Addresses).setMsg(obj)
	}
	return obj.addressesHolder
}

type isisV4RouteRangeV4RouteAddressIter struct {
	obj                 *isisV4RouteRange
	v4RouteAddressSlice []V4RouteAddress
	fieldPtr            *[]*otg.V4RouteAddress
}

func newIsisV4RouteRangeV4RouteAddressIter(ptr *[]*otg.V4RouteAddress) IsisV4RouteRangeV4RouteAddressIter {
	return &isisV4RouteRangeV4RouteAddressIter{fieldPtr: ptr}
}

type IsisV4RouteRangeV4RouteAddressIter interface {
	setMsg(*isisV4RouteRange) IsisV4RouteRangeV4RouteAddressIter
	Items() []V4RouteAddress
	Add() V4RouteAddress
	Append(items ...V4RouteAddress) IsisV4RouteRangeV4RouteAddressIter
	Set(index int, newObj V4RouteAddress) IsisV4RouteRangeV4RouteAddressIter
	Clear() IsisV4RouteRangeV4RouteAddressIter
	clearHolderSlice() IsisV4RouteRangeV4RouteAddressIter
	appendHolderSlice(item V4RouteAddress) IsisV4RouteRangeV4RouteAddressIter
}

func (obj *isisV4RouteRangeV4RouteAddressIter) setMsg(msg *isisV4RouteRange) IsisV4RouteRangeV4RouteAddressIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&v4RouteAddress{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisV4RouteRangeV4RouteAddressIter) Items() []V4RouteAddress {
	return obj.v4RouteAddressSlice
}

func (obj *isisV4RouteRangeV4RouteAddressIter) Add() V4RouteAddress {
	newObj := &otg.V4RouteAddress{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &v4RouteAddress{obj: newObj}
	newLibObj.setDefault()
	obj.v4RouteAddressSlice = append(obj.v4RouteAddressSlice, newLibObj)
	return newLibObj
}

func (obj *isisV4RouteRangeV4RouteAddressIter) Append(items ...V4RouteAddress) IsisV4RouteRangeV4RouteAddressIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.v4RouteAddressSlice = append(obj.v4RouteAddressSlice, item)
	}
	return obj
}

func (obj *isisV4RouteRangeV4RouteAddressIter) Set(index int, newObj V4RouteAddress) IsisV4RouteRangeV4RouteAddressIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.v4RouteAddressSlice[index] = newObj
	return obj
}
func (obj *isisV4RouteRangeV4RouteAddressIter) Clear() IsisV4RouteRangeV4RouteAddressIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.V4RouteAddress{}
		obj.v4RouteAddressSlice = []V4RouteAddress{}
	}
	return obj
}
func (obj *isisV4RouteRangeV4RouteAddressIter) clearHolderSlice() IsisV4RouteRangeV4RouteAddressIter {
	if len(obj.v4RouteAddressSlice) > 0 {
		obj.v4RouteAddressSlice = []V4RouteAddress{}
	}
	return obj
}
func (obj *isisV4RouteRangeV4RouteAddressIter) appendHolderSlice(item V4RouteAddress) IsisV4RouteRangeV4RouteAddressIter {
	obj.v4RouteAddressSlice = append(obj.v4RouteAddressSlice, item)
	return obj
}

// The user-defined metric associated with this route range.
// LinkMetric returns a uint32
func (obj *isisV4RouteRange) LinkMetric() uint32 {

	return *obj.obj.LinkMetric

}

// The user-defined metric associated with this route range.
// LinkMetric returns a uint32
func (obj *isisV4RouteRange) HasLinkMetric() bool {
	return obj.obj.LinkMetric != nil
}

// The user-defined metric associated with this route range.
// SetLinkMetric sets the uint32 value in the IsisV4RouteRange object
func (obj *isisV4RouteRange) SetLinkMetric(value uint32) IsisV4RouteRange {

	obj.obj.LinkMetric = &value
	return obj
}

type IsisV4RouteRangeOriginTypeEnum string

// Enum of OriginType on IsisV4RouteRange
var IsisV4RouteRangeOriginType = struct {
	INTERNAL IsisV4RouteRangeOriginTypeEnum
	EXTERNAL IsisV4RouteRangeOriginTypeEnum
}{
	INTERNAL: IsisV4RouteRangeOriginTypeEnum("internal"),
	EXTERNAL: IsisV4RouteRangeOriginTypeEnum("external"),
}

func (obj *isisV4RouteRange) OriginType() IsisV4RouteRangeOriginTypeEnum {
	return IsisV4RouteRangeOriginTypeEnum(obj.obj.OriginType.Enum().String())
}

// The origin of the advertised route-internal or external to the ISIS area. Options include the following:
// Internal-for intra-area routes, through Level 1 LSPs.
// External-for inter-area routes redistributed within L1, through Level
// 1 LSPs.
// OriginType returns a string
func (obj *isisV4RouteRange) HasOriginType() bool {
	return obj.obj.OriginType != nil
}

func (obj *isisV4RouteRange) SetOriginType(value IsisV4RouteRangeOriginTypeEnum) IsisV4RouteRange {
	intValue, ok := otg.IsisV4RouteRange_OriginType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisV4RouteRangeOriginTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisV4RouteRange_OriginType_Enum(intValue)
	obj.obj.OriginType = &enumValue

	return obj
}

type IsisV4RouteRangeRedistributionTypeEnum string

// Enum of RedistributionType on IsisV4RouteRange
var IsisV4RouteRangeRedistributionType = struct {
	UP   IsisV4RouteRangeRedistributionTypeEnum
	DOWN IsisV4RouteRangeRedistributionTypeEnum
}{
	UP:   IsisV4RouteRangeRedistributionTypeEnum("up"),
	DOWN: IsisV4RouteRangeRedistributionTypeEnum("down"),
}

func (obj *isisV4RouteRange) RedistributionType() IsisV4RouteRangeRedistributionTypeEnum {
	return IsisV4RouteRangeRedistributionTypeEnum(obj.obj.RedistributionType.Enum().String())
}

// Defines the Up/Down (Redistribution) bit defined for TLVs 128 and 130 by RFC 2966.  It is used for domain-wide advertisement of prefix information.
//
// Up (0)-used when a prefix is initially advertised within the ISIS L3
// hierarchy,
// and for all other prefixes in L1 and L2 LSPs. (default)
// Down (1)-used when an L1/L2 router advertises L2 prefixes in L1 LSPs.
//
// The prefixes are being advertised from a higher level (L2) down to a lower level (L1).
// RedistributionType returns a string
func (obj *isisV4RouteRange) HasRedistributionType() bool {
	return obj.obj.RedistributionType != nil
}

func (obj *isisV4RouteRange) SetRedistributionType(value IsisV4RouteRangeRedistributionTypeEnum) IsisV4RouteRange {
	intValue, ok := otg.IsisV4RouteRange_RedistributionType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisV4RouteRangeRedistributionTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisV4RouteRange_RedistributionType_Enum(intValue)
	obj.obj.RedistributionType = &enumValue

	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *isisV4RouteRange) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the IsisV4RouteRange object
func (obj *isisV4RouteRange) SetName(value string) IsisV4RouteRange {

	obj.obj.Name = &value
	return obj
}

// Specifies whether the sub-TLV for IPv4/IPv6 Extended Reachability Attribute Flags
// will be advertised or not.
// PrefixAttrEnabled returns a bool
func (obj *isisV4RouteRange) PrefixAttrEnabled() bool {

	return *obj.obj.PrefixAttrEnabled

}

// Specifies whether the sub-TLV for IPv4/IPv6 Extended Reachability Attribute Flags
// will be advertised or not.
// PrefixAttrEnabled returns a bool
func (obj *isisV4RouteRange) HasPrefixAttrEnabled() bool {
	return obj.obj.PrefixAttrEnabled != nil
}

// Specifies whether the sub-TLV for IPv4/IPv6 Extended Reachability Attribute Flags
// will be advertised or not.
// SetPrefixAttrEnabled sets the bool value in the IsisV4RouteRange object
func (obj *isisV4RouteRange) SetPrefixAttrEnabled(value bool) IsisV4RouteRange {

	obj.obj.PrefixAttrEnabled = &value
	return obj
}

// External Prefix Flag (Bit 0)
// XFlag returns a bool
func (obj *isisV4RouteRange) XFlag() bool {

	return *obj.obj.XFlag

}

// External Prefix Flag (Bit 0)
// XFlag returns a bool
func (obj *isisV4RouteRange) HasXFlag() bool {
	return obj.obj.XFlag != nil
}

// External Prefix Flag (Bit 0)
// SetXFlag sets the bool value in the IsisV4RouteRange object
func (obj *isisV4RouteRange) SetXFlag(value bool) IsisV4RouteRange {

	obj.obj.XFlag = &value
	return obj
}

// Re-advertisement Flag (Bit 1)
// RFlag returns a bool
func (obj *isisV4RouteRange) RFlag() bool {

	return *obj.obj.RFlag

}

// Re-advertisement Flag (Bit 1)
// RFlag returns a bool
func (obj *isisV4RouteRange) HasRFlag() bool {
	return obj.obj.RFlag != nil
}

// Re-advertisement Flag (Bit 1)
// SetRFlag sets the bool value in the IsisV4RouteRange object
func (obj *isisV4RouteRange) SetRFlag(value bool) IsisV4RouteRange {

	obj.obj.RFlag = &value
	return obj
}

// Node Flag (Bit 2)
// NFlag returns a bool
func (obj *isisV4RouteRange) NFlag() bool {

	return *obj.obj.NFlag

}

// Node Flag (Bit 2)
// NFlag returns a bool
func (obj *isisV4RouteRange) HasNFlag() bool {
	return obj.obj.NFlag != nil
}

// Node Flag (Bit 2)
// SetNFlag sets the bool value in the IsisV4RouteRange object
func (obj *isisV4RouteRange) SetNFlag(value bool) IsisV4RouteRange {

	obj.obj.NFlag = &value
	return obj
}

// A list of SID paramters for a group of IPv4 route addresses.
// PrefixSids returns a []IsisSRPrefixSid
func (obj *isisV4RouteRange) PrefixSids() IsisV4RouteRangeIsisSRPrefixSidIter {
	if len(obj.obj.PrefixSids) == 0 {
		obj.obj.PrefixSids = []*otg.IsisSRPrefixSid{}
	}
	if obj.prefixSidsHolder == nil {
		obj.prefixSidsHolder = newIsisV4RouteRangeIsisSRPrefixSidIter(&obj.obj.PrefixSids).setMsg(obj)
	}
	return obj.prefixSidsHolder
}

type isisV4RouteRangeIsisSRPrefixSidIter struct {
	obj                  *isisV4RouteRange
	isisSRPrefixSidSlice []IsisSRPrefixSid
	fieldPtr             *[]*otg.IsisSRPrefixSid
}

func newIsisV4RouteRangeIsisSRPrefixSidIter(ptr *[]*otg.IsisSRPrefixSid) IsisV4RouteRangeIsisSRPrefixSidIter {
	return &isisV4RouteRangeIsisSRPrefixSidIter{fieldPtr: ptr}
}

type IsisV4RouteRangeIsisSRPrefixSidIter interface {
	setMsg(*isisV4RouteRange) IsisV4RouteRangeIsisSRPrefixSidIter
	Items() []IsisSRPrefixSid
	Add() IsisSRPrefixSid
	Append(items ...IsisSRPrefixSid) IsisV4RouteRangeIsisSRPrefixSidIter
	Set(index int, newObj IsisSRPrefixSid) IsisV4RouteRangeIsisSRPrefixSidIter
	Clear() IsisV4RouteRangeIsisSRPrefixSidIter
	clearHolderSlice() IsisV4RouteRangeIsisSRPrefixSidIter
	appendHolderSlice(item IsisSRPrefixSid) IsisV4RouteRangeIsisSRPrefixSidIter
}

func (obj *isisV4RouteRangeIsisSRPrefixSidIter) setMsg(msg *isisV4RouteRange) IsisV4RouteRangeIsisSRPrefixSidIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisSRPrefixSid{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisV4RouteRangeIsisSRPrefixSidIter) Items() []IsisSRPrefixSid {
	return obj.isisSRPrefixSidSlice
}

func (obj *isisV4RouteRangeIsisSRPrefixSidIter) Add() IsisSRPrefixSid {
	newObj := &otg.IsisSRPrefixSid{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisSRPrefixSid{obj: newObj}
	newLibObj.setDefault()
	obj.isisSRPrefixSidSlice = append(obj.isisSRPrefixSidSlice, newLibObj)
	return newLibObj
}

func (obj *isisV4RouteRangeIsisSRPrefixSidIter) Append(items ...IsisSRPrefixSid) IsisV4RouteRangeIsisSRPrefixSidIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisSRPrefixSidSlice = append(obj.isisSRPrefixSidSlice, item)
	}
	return obj
}

func (obj *isisV4RouteRangeIsisSRPrefixSidIter) Set(index int, newObj IsisSRPrefixSid) IsisV4RouteRangeIsisSRPrefixSidIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisSRPrefixSidSlice[index] = newObj
	return obj
}
func (obj *isisV4RouteRangeIsisSRPrefixSidIter) Clear() IsisV4RouteRangeIsisSRPrefixSidIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisSRPrefixSid{}
		obj.isisSRPrefixSidSlice = []IsisSRPrefixSid{}
	}
	return obj
}
func (obj *isisV4RouteRangeIsisSRPrefixSidIter) clearHolderSlice() IsisV4RouteRangeIsisSRPrefixSidIter {
	if len(obj.isisSRPrefixSidSlice) > 0 {
		obj.isisSRPrefixSidSlice = []IsisSRPrefixSid{}
	}
	return obj
}
func (obj *isisV4RouteRangeIsisSRPrefixSidIter) appendHolderSlice(item IsisSRPrefixSid) IsisV4RouteRangeIsisSRPrefixSidIter {
	obj.isisSRPrefixSidSlice = append(obj.isisSRPrefixSidSlice, item)
	return obj
}

func (obj *isisV4RouteRange) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Addresses) != 0 {

		if set_default {
			obj.Addresses().clearHolderSlice()
			for _, item := range obj.obj.Addresses {
				obj.Addresses().appendHolderSlice(&v4RouteAddress{obj: item})
			}
		}
		for _, item := range obj.Addresses().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.LinkMetric != nil {

		if *obj.obj.LinkMetric > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisV4RouteRange.LinkMetric <= 16777215 but Got %d", *obj.obj.LinkMetric))
		}

	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface IsisV4RouteRange")
	}

	if len(obj.obj.PrefixSids) != 0 {

		if set_default {
			obj.PrefixSids().clearHolderSlice()
			for _, item := range obj.obj.PrefixSids {
				obj.PrefixSids().appendHolderSlice(&isisSRPrefixSid{obj: item})
			}
		}
		for _, item := range obj.PrefixSids().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *isisV4RouteRange) setDefault() {
	if obj.obj.LinkMetric == nil {
		obj.SetLinkMetric(0)
	}
	if obj.obj.OriginType == nil {
		obj.SetOriginType(IsisV4RouteRangeOriginType.INTERNAL)

	}
	if obj.obj.RedistributionType == nil {
		obj.SetRedistributionType(IsisV4RouteRangeRedistributionType.UP)

	}
	if obj.obj.PrefixAttrEnabled == nil {
		obj.SetPrefixAttrEnabled(false)
	}
	if obj.obj.XFlag == nil {
		obj.SetXFlag(false)
	}
	if obj.obj.RFlag == nil {
		obj.SetRFlag(false)
	}
	if obj.obj.NFlag == nil {
		obj.SetNFlag(false)
	}

}

package gosnappi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisV6RouteRange *****
type isisV6RouteRange struct {
	validation
	obj              *otg.IsisV6RouteRange
	marshaller       marshalIsisV6RouteRange
	unMarshaller     unMarshalIsisV6RouteRange
	addressesHolder  IsisV6RouteRangeV6RouteAddressIter
	prefixSidsHolder IsisV6RouteRangeIsisSRPrefixSidIter
}

func NewIsisV6RouteRange() IsisV6RouteRange {
	obj := isisV6RouteRange{obj: &otg.IsisV6RouteRange{}}
	obj.setDefault()
	return &obj
}

func (obj *isisV6RouteRange) msg() *otg.IsisV6RouteRange {
	return obj.obj
}

func (obj *isisV6RouteRange) setMsg(msg *otg.IsisV6RouteRange) IsisV6RouteRange {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisV6RouteRange struct {
	obj *isisV6RouteRange
}

type marshalIsisV6RouteRange interface {
	// ToProto marshals IsisV6RouteRange to protobuf object *otg.IsisV6RouteRange
	ToProto() (*otg.IsisV6RouteRange, error)
	// ToPbText marshals IsisV6RouteRange to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisV6RouteRange to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisV6RouteRange to JSON text
	ToJson() (string, error)
}

type unMarshalisisV6RouteRange struct {
	obj *isisV6RouteRange
}

type unMarshalIsisV6RouteRange interface {
	// FromProto unmarshals IsisV6RouteRange from protobuf object *otg.IsisV6RouteRange
	FromProto(msg *otg.IsisV6RouteRange) (IsisV6RouteRange, error)
	// FromPbText unmarshals IsisV6RouteRange from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisV6RouteRange from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisV6RouteRange from JSON text
	FromJson(value string) error
}

func (obj *isisV6RouteRange) Marshal() marshalIsisV6RouteRange {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisV6RouteRange{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisV6RouteRange) Unmarshal() unMarshalIsisV6RouteRange {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisV6RouteRange{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisV6RouteRange) ToProto() (*otg.IsisV6RouteRange, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisV6RouteRange) FromProto(msg *otg.IsisV6RouteRange) (IsisV6RouteRange, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisV6RouteRange) ToPbText() (string, error) {
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

func (m *unMarshalisisV6RouteRange) FromPbText(value string) error {
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

func (m *marshalisisV6RouteRange) ToYaml() (string, error) {
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

func (m *unMarshalisisV6RouteRange) FromYaml(value string) error {
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

func (m *marshalisisV6RouteRange) ToJson() (string, error) {
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

func (m *unMarshalisisV6RouteRange) FromJson(value string) error {
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

func (obj *isisV6RouteRange) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisV6RouteRange) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisV6RouteRange) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisV6RouteRange) Clone() (IsisV6RouteRange, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisV6RouteRange()
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

func (obj *isisV6RouteRange) setNil() {
	obj.addressesHolder = nil
	obj.prefixSidsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisV6RouteRange is emulated ISIS IPv6 routes.
type IsisV6RouteRange interface {
	Validation
	// msg marshals IsisV6RouteRange to protobuf object *otg.IsisV6RouteRange
	// and doesn't set defaults
	msg() *otg.IsisV6RouteRange
	// setMsg unmarshals IsisV6RouteRange from protobuf object *otg.IsisV6RouteRange
	// and doesn't set defaults
	setMsg(*otg.IsisV6RouteRange) IsisV6RouteRange
	// provides marshal interface
	Marshal() marshalIsisV6RouteRange
	// provides unmarshal interface
	Unmarshal() unMarshalIsisV6RouteRange
	// validate validates IsisV6RouteRange
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisV6RouteRange, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Addresses returns IsisV6RouteRangeV6RouteAddressIterIter, set in IsisV6RouteRange
	Addresses() IsisV6RouteRangeV6RouteAddressIter
	// LinkMetric returns uint32, set in IsisV6RouteRange.
	LinkMetric() uint32
	// SetLinkMetric assigns uint32 provided by user to IsisV6RouteRange
	SetLinkMetric(value uint32) IsisV6RouteRange
	// HasLinkMetric checks if LinkMetric has been set in IsisV6RouteRange
	HasLinkMetric() bool
	// OriginType returns IsisV6RouteRangeOriginTypeEnum, set in IsisV6RouteRange
	OriginType() IsisV6RouteRangeOriginTypeEnum
	// SetOriginType assigns IsisV6RouteRangeOriginTypeEnum provided by user to IsisV6RouteRange
	SetOriginType(value IsisV6RouteRangeOriginTypeEnum) IsisV6RouteRange
	// HasOriginType checks if OriginType has been set in IsisV6RouteRange
	HasOriginType() bool
	// RedistributionType returns IsisV6RouteRangeRedistributionTypeEnum, set in IsisV6RouteRange
	RedistributionType() IsisV6RouteRangeRedistributionTypeEnum
	// SetRedistributionType assigns IsisV6RouteRangeRedistributionTypeEnum provided by user to IsisV6RouteRange
	SetRedistributionType(value IsisV6RouteRangeRedistributionTypeEnum) IsisV6RouteRange
	// HasRedistributionType checks if RedistributionType has been set in IsisV6RouteRange
	HasRedistributionType() bool
	// Name returns string, set in IsisV6RouteRange.
	Name() string
	// SetName assigns string provided by user to IsisV6RouteRange
	SetName(value string) IsisV6RouteRange
	// PrefixAttrEnabled returns bool, set in IsisV6RouteRange.
	PrefixAttrEnabled() bool
	// SetPrefixAttrEnabled assigns bool provided by user to IsisV6RouteRange
	SetPrefixAttrEnabled(value bool) IsisV6RouteRange
	// HasPrefixAttrEnabled checks if PrefixAttrEnabled has been set in IsisV6RouteRange
	HasPrefixAttrEnabled() bool
	// XFlag returns bool, set in IsisV6RouteRange.
	XFlag() bool
	// SetXFlag assigns bool provided by user to IsisV6RouteRange
	SetXFlag(value bool) IsisV6RouteRange
	// HasXFlag checks if XFlag has been set in IsisV6RouteRange
	HasXFlag() bool
	// RFlag returns bool, set in IsisV6RouteRange.
	RFlag() bool
	// SetRFlag assigns bool provided by user to IsisV6RouteRange
	SetRFlag(value bool) IsisV6RouteRange
	// HasRFlag checks if RFlag has been set in IsisV6RouteRange
	HasRFlag() bool
	// NFlag returns bool, set in IsisV6RouteRange.
	NFlag() bool
	// SetNFlag assigns bool provided by user to IsisV6RouteRange
	SetNFlag(value bool) IsisV6RouteRange
	// HasNFlag checks if NFlag has been set in IsisV6RouteRange
	HasNFlag() bool
	// PrefixSids returns IsisV6RouteRangeIsisSRPrefixSidIterIter, set in IsisV6RouteRange
	PrefixSids() IsisV6RouteRangeIsisSRPrefixSidIter
	setNil()
}

// A list of group of IPv6 route addresses.
// Addresses returns a []V6RouteAddress
func (obj *isisV6RouteRange) Addresses() IsisV6RouteRangeV6RouteAddressIter {
	if len(obj.obj.Addresses) == 0 {
		obj.obj.Addresses = []*otg.V6RouteAddress{}
	}
	if obj.addressesHolder == nil {
		obj.addressesHolder = newIsisV6RouteRangeV6RouteAddressIter(&obj.obj.Addresses).setMsg(obj)
	}
	return obj.addressesHolder
}

type isisV6RouteRangeV6RouteAddressIter struct {
	obj                 *isisV6RouteRange
	v6RouteAddressSlice []V6RouteAddress
	fieldPtr            *[]*otg.V6RouteAddress
}

func newIsisV6RouteRangeV6RouteAddressIter(ptr *[]*otg.V6RouteAddress) IsisV6RouteRangeV6RouteAddressIter {
	return &isisV6RouteRangeV6RouteAddressIter{fieldPtr: ptr}
}

type IsisV6RouteRangeV6RouteAddressIter interface {
	setMsg(*isisV6RouteRange) IsisV6RouteRangeV6RouteAddressIter
	Items() []V6RouteAddress
	Add() V6RouteAddress
	Append(items ...V6RouteAddress) IsisV6RouteRangeV6RouteAddressIter
	Set(index int, newObj V6RouteAddress) IsisV6RouteRangeV6RouteAddressIter
	Clear() IsisV6RouteRangeV6RouteAddressIter
	clearHolderSlice() IsisV6RouteRangeV6RouteAddressIter
	appendHolderSlice(item V6RouteAddress) IsisV6RouteRangeV6RouteAddressIter
}

func (obj *isisV6RouteRangeV6RouteAddressIter) setMsg(msg *isisV6RouteRange) IsisV6RouteRangeV6RouteAddressIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&v6RouteAddress{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisV6RouteRangeV6RouteAddressIter) Items() []V6RouteAddress {
	return obj.v6RouteAddressSlice
}

func (obj *isisV6RouteRangeV6RouteAddressIter) Add() V6RouteAddress {
	newObj := &otg.V6RouteAddress{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &v6RouteAddress{obj: newObj}
	newLibObj.setDefault()
	obj.v6RouteAddressSlice = append(obj.v6RouteAddressSlice, newLibObj)
	return newLibObj
}

func (obj *isisV6RouteRangeV6RouteAddressIter) Append(items ...V6RouteAddress) IsisV6RouteRangeV6RouteAddressIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.v6RouteAddressSlice = append(obj.v6RouteAddressSlice, item)
	}
	return obj
}

func (obj *isisV6RouteRangeV6RouteAddressIter) Set(index int, newObj V6RouteAddress) IsisV6RouteRangeV6RouteAddressIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.v6RouteAddressSlice[index] = newObj
	return obj
}
func (obj *isisV6RouteRangeV6RouteAddressIter) Clear() IsisV6RouteRangeV6RouteAddressIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.V6RouteAddress{}
		obj.v6RouteAddressSlice = []V6RouteAddress{}
	}
	return obj
}
func (obj *isisV6RouteRangeV6RouteAddressIter) clearHolderSlice() IsisV6RouteRangeV6RouteAddressIter {
	if len(obj.v6RouteAddressSlice) > 0 {
		obj.v6RouteAddressSlice = []V6RouteAddress{}
	}
	return obj
}
func (obj *isisV6RouteRangeV6RouteAddressIter) appendHolderSlice(item V6RouteAddress) IsisV6RouteRangeV6RouteAddressIter {
	obj.v6RouteAddressSlice = append(obj.v6RouteAddressSlice, item)
	return obj
}

// The user-defined metric associated with this route range.
// LinkMetric returns a uint32
func (obj *isisV6RouteRange) LinkMetric() uint32 {

	return *obj.obj.LinkMetric

}

// The user-defined metric associated with this route range.
// LinkMetric returns a uint32
func (obj *isisV6RouteRange) HasLinkMetric() bool {
	return obj.obj.LinkMetric != nil
}

// The user-defined metric associated with this route range.
// SetLinkMetric sets the uint32 value in the IsisV6RouteRange object
func (obj *isisV6RouteRange) SetLinkMetric(value uint32) IsisV6RouteRange {

	obj.obj.LinkMetric = &value
	return obj
}

type IsisV6RouteRangeOriginTypeEnum string

// Enum of OriginType on IsisV6RouteRange
var IsisV6RouteRangeOriginType = struct {
	INTERNAL IsisV6RouteRangeOriginTypeEnum
	EXTERNAL IsisV6RouteRangeOriginTypeEnum
}{
	INTERNAL: IsisV6RouteRangeOriginTypeEnum("internal"),
	EXTERNAL: IsisV6RouteRangeOriginTypeEnum("external"),
}

func (obj *isisV6RouteRange) OriginType() IsisV6RouteRangeOriginTypeEnum {
	return IsisV6RouteRangeOriginTypeEnum(obj.obj.OriginType.Enum().String())
}

// The origin of the advertised route-internal or external to the ISIS area. Options include the following:
// Internal-for intra-area routes, through Level 1 LSPs.
// External-for inter-area routes redistributed within L1, through Level
// 1 LSPs.
// OriginType returns a string
func (obj *isisV6RouteRange) HasOriginType() bool {
	return obj.obj.OriginType != nil
}

func (obj *isisV6RouteRange) SetOriginType(value IsisV6RouteRangeOriginTypeEnum) IsisV6RouteRange {
	intValue, ok := otg.IsisV6RouteRange_OriginType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisV6RouteRangeOriginTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisV6RouteRange_OriginType_Enum(intValue)
	obj.obj.OriginType = &enumValue

	return obj
}

type IsisV6RouteRangeRedistributionTypeEnum string

// Enum of RedistributionType on IsisV6RouteRange
var IsisV6RouteRangeRedistributionType = struct {
	UP   IsisV6RouteRangeRedistributionTypeEnum
	DOWN IsisV6RouteRangeRedistributionTypeEnum
}{
	UP:   IsisV6RouteRangeRedistributionTypeEnum("up"),
	DOWN: IsisV6RouteRangeRedistributionTypeEnum("down"),
}

func (obj *isisV6RouteRange) RedistributionType() IsisV6RouteRangeRedistributionTypeEnum {
	return IsisV6RouteRangeRedistributionTypeEnum(obj.obj.RedistributionType.Enum().String())
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
func (obj *isisV6RouteRange) HasRedistributionType() bool {
	return obj.obj.RedistributionType != nil
}

func (obj *isisV6RouteRange) SetRedistributionType(value IsisV6RouteRangeRedistributionTypeEnum) IsisV6RouteRange {
	intValue, ok := otg.IsisV6RouteRange_RedistributionType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisV6RouteRangeRedistributionTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisV6RouteRange_RedistributionType_Enum(intValue)
	obj.obj.RedistributionType = &enumValue

	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *isisV6RouteRange) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the IsisV6RouteRange object
func (obj *isisV6RouteRange) SetName(value string) IsisV6RouteRange {

	obj.obj.Name = &value
	return obj
}

// Specifies whether the sub-TLV for IPv4/IPv6 Extended Reachability Attribute Flags
// will be advertised or not.
// PrefixAttrEnabled returns a bool
func (obj *isisV6RouteRange) PrefixAttrEnabled() bool {

	return *obj.obj.PrefixAttrEnabled

}

// Specifies whether the sub-TLV for IPv4/IPv6 Extended Reachability Attribute Flags
// will be advertised or not.
// PrefixAttrEnabled returns a bool
func (obj *isisV6RouteRange) HasPrefixAttrEnabled() bool {
	return obj.obj.PrefixAttrEnabled != nil
}

// Specifies whether the sub-TLV for IPv4/IPv6 Extended Reachability Attribute Flags
// will be advertised or not.
// SetPrefixAttrEnabled sets the bool value in the IsisV6RouteRange object
func (obj *isisV6RouteRange) SetPrefixAttrEnabled(value bool) IsisV6RouteRange {

	obj.obj.PrefixAttrEnabled = &value
	return obj
}

// External Prefix Flag (Bit 0)
// XFlag returns a bool
func (obj *isisV6RouteRange) XFlag() bool {

	return *obj.obj.XFlag

}

// External Prefix Flag (Bit 0)
// XFlag returns a bool
func (obj *isisV6RouteRange) HasXFlag() bool {
	return obj.obj.XFlag != nil
}

// External Prefix Flag (Bit 0)
// SetXFlag sets the bool value in the IsisV6RouteRange object
func (obj *isisV6RouteRange) SetXFlag(value bool) IsisV6RouteRange {

	obj.obj.XFlag = &value
	return obj
}

// Re-advertisement Flag (Bit 1)
// RFlag returns a bool
func (obj *isisV6RouteRange) RFlag() bool {

	return *obj.obj.RFlag

}

// Re-advertisement Flag (Bit 1)
// RFlag returns a bool
func (obj *isisV6RouteRange) HasRFlag() bool {
	return obj.obj.RFlag != nil
}

// Re-advertisement Flag (Bit 1)
// SetRFlag sets the bool value in the IsisV6RouteRange object
func (obj *isisV6RouteRange) SetRFlag(value bool) IsisV6RouteRange {

	obj.obj.RFlag = &value
	return obj
}

// Node Flag (Bit 2)
// NFlag returns a bool
func (obj *isisV6RouteRange) NFlag() bool {

	return *obj.obj.NFlag

}

// Node Flag (Bit 2)
// NFlag returns a bool
func (obj *isisV6RouteRange) HasNFlag() bool {
	return obj.obj.NFlag != nil
}

// Node Flag (Bit 2)
// SetNFlag sets the bool value in the IsisV6RouteRange object
func (obj *isisV6RouteRange) SetNFlag(value bool) IsisV6RouteRange {

	obj.obj.NFlag = &value
	return obj
}

// A list of SID parameters for a group of IPv6 route addresses.
// PrefixSids returns a []IsisSRPrefixSid
func (obj *isisV6RouteRange) PrefixSids() IsisV6RouteRangeIsisSRPrefixSidIter {
	if len(obj.obj.PrefixSids) == 0 {
		obj.obj.PrefixSids = []*otg.IsisSRPrefixSid{}
	}
	if obj.prefixSidsHolder == nil {
		obj.prefixSidsHolder = newIsisV6RouteRangeIsisSRPrefixSidIter(&obj.obj.PrefixSids).setMsg(obj)
	}
	return obj.prefixSidsHolder
}

type isisV6RouteRangeIsisSRPrefixSidIter struct {
	obj                  *isisV6RouteRange
	isisSRPrefixSidSlice []IsisSRPrefixSid
	fieldPtr             *[]*otg.IsisSRPrefixSid
}

func newIsisV6RouteRangeIsisSRPrefixSidIter(ptr *[]*otg.IsisSRPrefixSid) IsisV6RouteRangeIsisSRPrefixSidIter {
	return &isisV6RouteRangeIsisSRPrefixSidIter{fieldPtr: ptr}
}

type IsisV6RouteRangeIsisSRPrefixSidIter interface {
	setMsg(*isisV6RouteRange) IsisV6RouteRangeIsisSRPrefixSidIter
	Items() []IsisSRPrefixSid
	Add() IsisSRPrefixSid
	Append(items ...IsisSRPrefixSid) IsisV6RouteRangeIsisSRPrefixSidIter
	Set(index int, newObj IsisSRPrefixSid) IsisV6RouteRangeIsisSRPrefixSidIter
	Clear() IsisV6RouteRangeIsisSRPrefixSidIter
	clearHolderSlice() IsisV6RouteRangeIsisSRPrefixSidIter
	appendHolderSlice(item IsisSRPrefixSid) IsisV6RouteRangeIsisSRPrefixSidIter
}

func (obj *isisV6RouteRangeIsisSRPrefixSidIter) setMsg(msg *isisV6RouteRange) IsisV6RouteRangeIsisSRPrefixSidIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisSRPrefixSid{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisV6RouteRangeIsisSRPrefixSidIter) Items() []IsisSRPrefixSid {
	return obj.isisSRPrefixSidSlice
}

func (obj *isisV6RouteRangeIsisSRPrefixSidIter) Add() IsisSRPrefixSid {
	newObj := &otg.IsisSRPrefixSid{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisSRPrefixSid{obj: newObj}
	newLibObj.setDefault()
	obj.isisSRPrefixSidSlice = append(obj.isisSRPrefixSidSlice, newLibObj)
	return newLibObj
}

func (obj *isisV6RouteRangeIsisSRPrefixSidIter) Append(items ...IsisSRPrefixSid) IsisV6RouteRangeIsisSRPrefixSidIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisSRPrefixSidSlice = append(obj.isisSRPrefixSidSlice, item)
	}
	return obj
}

func (obj *isisV6RouteRangeIsisSRPrefixSidIter) Set(index int, newObj IsisSRPrefixSid) IsisV6RouteRangeIsisSRPrefixSidIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisSRPrefixSidSlice[index] = newObj
	return obj
}
func (obj *isisV6RouteRangeIsisSRPrefixSidIter) Clear() IsisV6RouteRangeIsisSRPrefixSidIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisSRPrefixSid{}
		obj.isisSRPrefixSidSlice = []IsisSRPrefixSid{}
	}
	return obj
}
func (obj *isisV6RouteRangeIsisSRPrefixSidIter) clearHolderSlice() IsisV6RouteRangeIsisSRPrefixSidIter {
	if len(obj.isisSRPrefixSidSlice) > 0 {
		obj.isisSRPrefixSidSlice = []IsisSRPrefixSid{}
	}
	return obj
}
func (obj *isisV6RouteRangeIsisSRPrefixSidIter) appendHolderSlice(item IsisSRPrefixSid) IsisV6RouteRangeIsisSRPrefixSidIter {
	obj.isisSRPrefixSidSlice = append(obj.isisSRPrefixSidSlice, item)
	return obj
}

func (obj *isisV6RouteRange) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Addresses) != 0 {

		if set_default {
			obj.Addresses().clearHolderSlice()
			for _, item := range obj.obj.Addresses {
				obj.Addresses().appendHolderSlice(&v6RouteAddress{obj: item})
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
				fmt.Sprintf("0 <= IsisV6RouteRange.LinkMetric <= 16777215 but Got %d", *obj.obj.LinkMetric))
		}

	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface IsisV6RouteRange")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"IsisV6RouteRange.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

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

func (obj *isisV6RouteRange) setDefault() {
	if obj.obj.LinkMetric == nil {
		obj.SetLinkMetric(0)
	}
	if obj.obj.OriginType == nil {
		obj.SetOriginType(IsisV6RouteRangeOriginType.INTERNAL)

	}
	if obj.obj.RedistributionType == nil {
		obj.SetRedistributionType(IsisV6RouteRangeRedistributionType.UP)

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

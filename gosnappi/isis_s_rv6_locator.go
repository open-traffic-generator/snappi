package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRv6Locator *****
type isisSRv6Locator struct {
	validation
	obj                            *otg.IsisSRv6Locator
	marshaller                     marshalIsisSRv6Locator
	unMarshaller                   unMarshalIsisSRv6Locator
	endSidsHolder                  IsisSRv6LocatorIsisSRv6EndSidIter
	advertiseLocatorAsPrefixHolder IsisSRv6AdvertiseLocatorAsPrefix
}

func NewIsisSRv6Locator() IsisSRv6Locator {
	obj := isisSRv6Locator{obj: &otg.IsisSRv6Locator{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRv6Locator) msg() *otg.IsisSRv6Locator {
	return obj.obj
}

func (obj *isisSRv6Locator) setMsg(msg *otg.IsisSRv6Locator) IsisSRv6Locator {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRv6Locator struct {
	obj *isisSRv6Locator
}

type marshalIsisSRv6Locator interface {
	// ToProto marshals IsisSRv6Locator to protobuf object *otg.IsisSRv6Locator
	ToProto() (*otg.IsisSRv6Locator, error)
	// ToPbText marshals IsisSRv6Locator to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRv6Locator to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRv6Locator to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRv6Locator struct {
	obj *isisSRv6Locator
}

type unMarshalIsisSRv6Locator interface {
	// FromProto unmarshals IsisSRv6Locator from protobuf object *otg.IsisSRv6Locator
	FromProto(msg *otg.IsisSRv6Locator) (IsisSRv6Locator, error)
	// FromPbText unmarshals IsisSRv6Locator from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRv6Locator from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRv6Locator from JSON text
	FromJson(value string) error
}

func (obj *isisSRv6Locator) Marshal() marshalIsisSRv6Locator {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRv6Locator{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRv6Locator) Unmarshal() unMarshalIsisSRv6Locator {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRv6Locator{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRv6Locator) ToProto() (*otg.IsisSRv6Locator, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRv6Locator) FromProto(msg *otg.IsisSRv6Locator) (IsisSRv6Locator, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRv6Locator) ToPbText() (string, error) {
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

func (m *unMarshalisisSRv6Locator) FromPbText(value string) error {
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

func (m *marshalisisSRv6Locator) ToYaml() (string, error) {
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

func (m *unMarshalisisSRv6Locator) FromYaml(value string) error {
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

func (m *marshalisisSRv6Locator) ToJson() (string, error) {
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

func (m *unMarshalisisSRv6Locator) FromJson(value string) error {
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

func (obj *isisSRv6Locator) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRv6Locator) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRv6Locator) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRv6Locator) Clone() (IsisSRv6Locator, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRv6Locator()
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

func (obj *isisSRv6Locator) setNil() {
	obj.endSidsHolder = nil
	obj.advertiseLocatorAsPrefixHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisSRv6Locator is iS-IS SRv6 Locator TLV (TLV type 27), as defined in RFC 9352 Section 7.1. Advertises an SRv6 locator prefix that identifies this SRv6-capable router and binds it to a specific IGP algorithm (standard SPF or Flex-Algo per RFC 9350). One Locator TLV is required per topology/algorithm combination. The SRv6 SID space is structured as: Locator prefix | Function | Argument, where the locator prefix is routable and the function identifies the specific endpoint behavior (RFC 8986 Section 3.1).
type IsisSRv6Locator interface {
	Validation
	// msg marshals IsisSRv6Locator to protobuf object *otg.IsisSRv6Locator
	// and doesn't set defaults
	msg() *otg.IsisSRv6Locator
	// setMsg unmarshals IsisSRv6Locator from protobuf object *otg.IsisSRv6Locator
	// and doesn't set defaults
	setMsg(*otg.IsisSRv6Locator) IsisSRv6Locator
	// provides marshal interface
	Marshal() marshalIsisSRv6Locator
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRv6Locator
	// validate validates IsisSRv6Locator
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRv6Locator, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in IsisSRv6Locator.
	Name() string
	// SetName assigns string provided by user to IsisSRv6Locator
	SetName(value string) IsisSRv6Locator
	// Locator returns string, set in IsisSRv6Locator.
	Locator() string
	// SetLocator assigns string provided by user to IsisSRv6Locator
	SetLocator(value string) IsisSRv6Locator
	// PrefixLength returns uint32, set in IsisSRv6Locator.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to IsisSRv6Locator
	SetPrefixLength(value uint32) IsisSRv6Locator
	// HasPrefixLength checks if PrefixLength has been set in IsisSRv6Locator
	HasPrefixLength() bool
	// Algorithm returns uint32, set in IsisSRv6Locator.
	Algorithm() uint32
	// SetAlgorithm assigns uint32 provided by user to IsisSRv6Locator
	SetAlgorithm(value uint32) IsisSRv6Locator
	// HasAlgorithm checks if Algorithm has been set in IsisSRv6Locator
	HasAlgorithm() bool
	// Metric returns uint32, set in IsisSRv6Locator.
	Metric() uint32
	// SetMetric assigns uint32 provided by user to IsisSRv6Locator
	SetMetric(value uint32) IsisSRv6Locator
	// HasMetric checks if Metric has been set in IsisSRv6Locator
	HasMetric() bool
	// RedistributionType returns IsisSRv6LocatorRedistributionTypeEnum, set in IsisSRv6Locator
	RedistributionType() IsisSRv6LocatorRedistributionTypeEnum
	// SetRedistributionType assigns IsisSRv6LocatorRedistributionTypeEnum provided by user to IsisSRv6Locator
	SetRedistributionType(value IsisSRv6LocatorRedistributionTypeEnum) IsisSRv6Locator
	// HasRedistributionType checks if RedistributionType has been set in IsisSRv6Locator
	HasRedistributionType() bool
	// DFlag returns bool, set in IsisSRv6Locator.
	DFlag() bool
	// SetDFlag assigns bool provided by user to IsisSRv6Locator
	SetDFlag(value bool) IsisSRv6Locator
	// HasDFlag checks if DFlag has been set in IsisSRv6Locator
	HasDFlag() bool
	// MtId returns uint32, set in IsisSRv6Locator.
	MtId() uint32
	// SetMtId assigns uint32 provided by user to IsisSRv6Locator
	SetMtId(value uint32) IsisSRv6Locator
	// HasMtId checks if MtId has been set in IsisSRv6Locator
	HasMtId() bool
	// EndSids returns IsisSRv6LocatorIsisSRv6EndSidIterIter, set in IsisSRv6Locator
	EndSids() IsisSRv6LocatorIsisSRv6EndSidIter
	// AdvertiseLocatorAsPrefix returns IsisSRv6AdvertiseLocatorAsPrefix, set in IsisSRv6Locator.
	// IsisSRv6AdvertiseLocatorAsPrefix is controls advertisement of the SRv6 locator prefix as an IS-IS IPv6 Reachability prefix (TLV 236/237) alongside the SRv6 Locator TLV (type 27). When this object is present the secondary prefix advertisement is enabled; when absent it is suppressed. Reference: RFC 9352 Section 7.1.
	AdvertiseLocatorAsPrefix() IsisSRv6AdvertiseLocatorAsPrefix
	// SetAdvertiseLocatorAsPrefix assigns IsisSRv6AdvertiseLocatorAsPrefix provided by user to IsisSRv6Locator.
	// IsisSRv6AdvertiseLocatorAsPrefix is controls advertisement of the SRv6 locator prefix as an IS-IS IPv6 Reachability prefix (TLV 236/237) alongside the SRv6 Locator TLV (type 27). When this object is present the secondary prefix advertisement is enabled; when absent it is suppressed. Reference: RFC 9352 Section 7.1.
	SetAdvertiseLocatorAsPrefix(value IsisSRv6AdvertiseLocatorAsPrefix) IsisSRv6Locator
	// HasAdvertiseLocatorAsPrefix checks if AdvertiseLocatorAsPrefix has been set in IsisSRv6Locator
	HasAdvertiseLocatorAsPrefix() bool
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *isisSRv6Locator) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the IsisSRv6Locator object
func (obj *isisSRv6Locator) SetName(value string) IsisSRv6Locator {

	obj.obj.Name = &value
	return obj
}

// The SRv6 locator IPv6 prefix for this TLV. This routable IPv6 prefix identifies the advertising node within the SRv6 domain. All SRv6 SIDs advertised under this locator must be allocated from this prefix.
// Locator returns a string
func (obj *isisSRv6Locator) Locator() string {

	return *obj.obj.Locator

}

// The SRv6 locator IPv6 prefix for this TLV. This routable IPv6 prefix identifies the advertising node within the SRv6 domain. All SRv6 SIDs advertised under this locator must be allocated from this prefix.
// SetLocator sets the string value in the IsisSRv6Locator object
func (obj *isisSRv6Locator) SetLocator(value string) IsisSRv6Locator {

	obj.obj.Locator = &value
	return obj
}

// Number of significant bits in the locator field (1-128). Defines the length of the SRv6 locator prefix. The combined bit budget of the locator, function, and argument fields must not exceed 128 bits (RFC 9352 Section 3.1).
// PrefixLength returns a uint32
func (obj *isisSRv6Locator) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// Number of significant bits in the locator field (1-128). Defines the length of the SRv6 locator prefix. The combined bit budget of the locator, function, and argument fields must not exceed 128 bits (RFC 9352 Section 3.1).
// PrefixLength returns a uint32
func (obj *isisSRv6Locator) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// Number of significant bits in the locator field (1-128). Defines the length of the SRv6 locator prefix. The combined bit budget of the locator, function, and argument fields must not exceed 128 bits (RFC 9352 Section 3.1).
// SetPrefixLength sets the uint32 value in the IsisSRv6Locator object
func (obj *isisSRv6Locator) SetPrefixLength(value uint32) IsisSRv6Locator {

	obj.obj.PrefixLength = &value
	return obj
}

// IGP Algorithm identifier for this locator (RFC 8665). 0 = standard SPF, 1 = Strict SPF. Values 128-255 represent Flexible Algorithms (Flex-Algo) as defined in RFC 9350. The algorithm determines the path computation method used to reach this locator. One Locator TLV is required per algorithm binding.
// Algorithm returns a uint32
func (obj *isisSRv6Locator) Algorithm() uint32 {

	return *obj.obj.Algorithm

}

// IGP Algorithm identifier for this locator (RFC 8665). 0 = standard SPF, 1 = Strict SPF. Values 128-255 represent Flexible Algorithms (Flex-Algo) as defined in RFC 9350. The algorithm determines the path computation method used to reach this locator. One Locator TLV is required per algorithm binding.
// Algorithm returns a uint32
func (obj *isisSRv6Locator) HasAlgorithm() bool {
	return obj.obj.Algorithm != nil
}

// IGP Algorithm identifier for this locator (RFC 8665). 0 = standard SPF, 1 = Strict SPF. Values 128-255 represent Flexible Algorithms (Flex-Algo) as defined in RFC 9350. The algorithm determines the path computation method used to reach this locator. One Locator TLV is required per algorithm binding.
// SetAlgorithm sets the uint32 value in the IsisSRv6Locator object
func (obj *isisSRv6Locator) SetAlgorithm(value uint32) IsisSRv6Locator {

	obj.obj.Algorithm = &value
	return obj
}

// The metric (cost) associated with this SRv6 locator prefix. Uses the same semantics as IS-IS Extended IP reachability metric (RFC 9352 Section 7.1).
// Metric returns a uint32
func (obj *isisSRv6Locator) Metric() uint32 {

	return *obj.obj.Metric

}

// The metric (cost) associated with this SRv6 locator prefix. Uses the same semantics as IS-IS Extended IP reachability metric (RFC 9352 Section 7.1).
// Metric returns a uint32
func (obj *isisSRv6Locator) HasMetric() bool {
	return obj.obj.Metric != nil
}

// The metric (cost) associated with this SRv6 locator prefix. Uses the same semantics as IS-IS Extended IP reachability metric (RFC 9352 Section 7.1).
// SetMetric sets the uint32 value in the IsisSRv6Locator object
func (obj *isisSRv6Locator) SetMetric(value uint32) IsisSRv6Locator {

	obj.obj.Metric = &value
	return obj
}

type IsisSRv6LocatorRedistributionTypeEnum string

// Enum of RedistributionType on IsisSRv6Locator
var IsisSRv6LocatorRedistributionType = struct {
	UP   IsisSRv6LocatorRedistributionTypeEnum
	DOWN IsisSRv6LocatorRedistributionTypeEnum
}{
	UP:   IsisSRv6LocatorRedistributionTypeEnum("up"),
	DOWN: IsisSRv6LocatorRedistributionTypeEnum("down"),
}

func (obj *isisSRv6Locator) RedistributionType() IsisSRv6LocatorRedistributionTypeEnum {
	return IsisSRv6LocatorRedistributionTypeEnum(obj.obj.RedistributionType.Enum().String())
}

// Controls the Up/Down redistribution bit for this locator. 'up' = normal advertisement (default); 'down' = the locator has been leaked from Level 2 to Level 1 (RFC 5305). Must be consistent with the d_flag setting.
// RedistributionType returns a string
func (obj *isisSRv6Locator) HasRedistributionType() bool {
	return obj.obj.RedistributionType != nil
}

func (obj *isisSRv6Locator) SetRedistributionType(value IsisSRv6LocatorRedistributionTypeEnum) IsisSRv6Locator {
	intValue, ok := otg.IsisSRv6Locator_RedistributionType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisSRv6LocatorRedistributionTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisSRv6Locator_RedistributionType_Enum(intValue)
	obj.obj.RedistributionType = &enumValue

	return obj
}

// Down bit (D-flag, bit 0 of the Flags field in the SRv6 Locator TLV). MUST be set when the locator is leaked from Level 2 to Level 1 to prevent routing loops. Locator TLVs with the D-flag set MUST NOT be re-advertised from L1 to L2 (RFC 9352 Section 7.1).
// DFlag returns a bool
func (obj *isisSRv6Locator) DFlag() bool {

	return *obj.obj.DFlag

}

// Down bit (D-flag, bit 0 of the Flags field in the SRv6 Locator TLV). MUST be set when the locator is leaked from Level 2 to Level 1 to prevent routing loops. Locator TLVs with the D-flag set MUST NOT be re-advertised from L1 to L2 (RFC 9352 Section 7.1).
// DFlag returns a bool
func (obj *isisSRv6Locator) HasDFlag() bool {
	return obj.obj.DFlag != nil
}

// Down bit (D-flag, bit 0 of the Flags field in the SRv6 Locator TLV). MUST be set when the locator is leaked from Level 2 to Level 1 to prevent routing loops. Locator TLVs with the D-flag set MUST NOT be re-advertised from L1 to L2 (RFC 9352 Section 7.1).
// SetDFlag sets the bool value in the IsisSRv6Locator object
func (obj *isisSRv6Locator) SetDFlag(value bool) IsisSRv6Locator {

	obj.obj.DFlag = &value
	return obj
}

// Multi-Topology Identifier (MT-ID) for this locator advertisement. Specifies which IS-IS topology this locator belongs to. 0 = default topology. Valid range 0-4095 (RFC 5120).
// MtId returns a uint32
func (obj *isisSRv6Locator) MtId() uint32 {

	return *obj.obj.MtId

}

// Multi-Topology Identifier (MT-ID) for this locator advertisement. Specifies which IS-IS topology this locator belongs to. 0 = default topology. Valid range 0-4095 (RFC 5120).
// MtId returns a uint32
func (obj *isisSRv6Locator) HasMtId() bool {
	return obj.obj.MtId != nil
}

// Multi-Topology Identifier (MT-ID) for this locator advertisement. Specifies which IS-IS topology this locator belongs to. 0 = default topology. Valid range 0-4095 (RFC 5120).
// SetMtId sets the uint32 value in the IsisSRv6Locator object
func (obj *isisSRv6Locator) SetMtId(value uint32) IsisSRv6Locator {

	obj.obj.MtId = &value
	return obj
}

// List of SRv6 End SID Sub-TLVs (sub-TLV type 5) associated with this locator. Each End SID advertises a locally instantiated node-local SRv6 SID, its endpoint behavior (End, End.DT4, End.DT6, End.DT46 and their flavor variants), and optional SID structure information (RFC 9352 Section 7.2).
// EndSids returns a []IsisSRv6EndSid
func (obj *isisSRv6Locator) EndSids() IsisSRv6LocatorIsisSRv6EndSidIter {
	if len(obj.obj.EndSids) == 0 {
		obj.obj.EndSids = []*otg.IsisSRv6EndSid{}
	}
	if obj.endSidsHolder == nil {
		obj.endSidsHolder = newIsisSRv6LocatorIsisSRv6EndSidIter(&obj.obj.EndSids).setMsg(obj)
	}
	return obj.endSidsHolder
}

type isisSRv6LocatorIsisSRv6EndSidIter struct {
	obj                 *isisSRv6Locator
	isisSRv6EndSidSlice []IsisSRv6EndSid
	fieldPtr            *[]*otg.IsisSRv6EndSid
}

func newIsisSRv6LocatorIsisSRv6EndSidIter(ptr *[]*otg.IsisSRv6EndSid) IsisSRv6LocatorIsisSRv6EndSidIter {
	return &isisSRv6LocatorIsisSRv6EndSidIter{fieldPtr: ptr}
}

type IsisSRv6LocatorIsisSRv6EndSidIter interface {
	setMsg(*isisSRv6Locator) IsisSRv6LocatorIsisSRv6EndSidIter
	Items() []IsisSRv6EndSid
	Add() IsisSRv6EndSid
	Append(items ...IsisSRv6EndSid) IsisSRv6LocatorIsisSRv6EndSidIter
	Set(index int, newObj IsisSRv6EndSid) IsisSRv6LocatorIsisSRv6EndSidIter
	Clear() IsisSRv6LocatorIsisSRv6EndSidIter
	clearHolderSlice() IsisSRv6LocatorIsisSRv6EndSidIter
	appendHolderSlice(item IsisSRv6EndSid) IsisSRv6LocatorIsisSRv6EndSidIter
}

func (obj *isisSRv6LocatorIsisSRv6EndSidIter) setMsg(msg *isisSRv6Locator) IsisSRv6LocatorIsisSRv6EndSidIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisSRv6EndSid{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisSRv6LocatorIsisSRv6EndSidIter) Items() []IsisSRv6EndSid {
	return obj.isisSRv6EndSidSlice
}

func (obj *isisSRv6LocatorIsisSRv6EndSidIter) Add() IsisSRv6EndSid {
	newObj := &otg.IsisSRv6EndSid{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisSRv6EndSid{obj: newObj}
	newLibObj.setDefault()
	obj.isisSRv6EndSidSlice = append(obj.isisSRv6EndSidSlice, newLibObj)
	return newLibObj
}

func (obj *isisSRv6LocatorIsisSRv6EndSidIter) Append(items ...IsisSRv6EndSid) IsisSRv6LocatorIsisSRv6EndSidIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisSRv6EndSidSlice = append(obj.isisSRv6EndSidSlice, item)
	}
	return obj
}

func (obj *isisSRv6LocatorIsisSRv6EndSidIter) Set(index int, newObj IsisSRv6EndSid) IsisSRv6LocatorIsisSRv6EndSidIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisSRv6EndSidSlice[index] = newObj
	return obj
}
func (obj *isisSRv6LocatorIsisSRv6EndSidIter) Clear() IsisSRv6LocatorIsisSRv6EndSidIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisSRv6EndSid{}
		obj.isisSRv6EndSidSlice = []IsisSRv6EndSid{}
	}
	return obj
}
func (obj *isisSRv6LocatorIsisSRv6EndSidIter) clearHolderSlice() IsisSRv6LocatorIsisSRv6EndSidIter {
	if len(obj.isisSRv6EndSidSlice) > 0 {
		obj.isisSRv6EndSidSlice = []IsisSRv6EndSid{}
	}
	return obj
}
func (obj *isisSRv6LocatorIsisSRv6EndSidIter) appendHolderSlice(item IsisSRv6EndSid) IsisSRv6LocatorIsisSRv6EndSidIter {
	obj.isisSRv6EndSidSlice = append(obj.isisSRv6EndSidSlice, item)
	return obj
}

// When present, the locator prefix is also advertised as an IS-IS IPv6 Reachability prefix (TLV 236/237) in addition to the SRv6 Locator TLV (type 27). This enables non-SRv6-capable routers to route toward the locator as an ordinary IPv6 prefix. Absence of this object suppresses the secondary prefix advertisement.
// AdvertiseLocatorAsPrefix returns a IsisSRv6AdvertiseLocatorAsPrefix
func (obj *isisSRv6Locator) AdvertiseLocatorAsPrefix() IsisSRv6AdvertiseLocatorAsPrefix {
	if obj.obj.AdvertiseLocatorAsPrefix == nil {
		obj.obj.AdvertiseLocatorAsPrefix = NewIsisSRv6AdvertiseLocatorAsPrefix().msg()
	}
	if obj.advertiseLocatorAsPrefixHolder == nil {
		obj.advertiseLocatorAsPrefixHolder = &isisSRv6AdvertiseLocatorAsPrefix{obj: obj.obj.AdvertiseLocatorAsPrefix}
	}
	return obj.advertiseLocatorAsPrefixHolder
}

// When present, the locator prefix is also advertised as an IS-IS IPv6 Reachability prefix (TLV 236/237) in addition to the SRv6 Locator TLV (type 27). This enables non-SRv6-capable routers to route toward the locator as an ordinary IPv6 prefix. Absence of this object suppresses the secondary prefix advertisement.
// AdvertiseLocatorAsPrefix returns a IsisSRv6AdvertiseLocatorAsPrefix
func (obj *isisSRv6Locator) HasAdvertiseLocatorAsPrefix() bool {
	return obj.obj.AdvertiseLocatorAsPrefix != nil
}

// When present, the locator prefix is also advertised as an IS-IS IPv6 Reachability prefix (TLV 236/237) in addition to the SRv6 Locator TLV (type 27). This enables non-SRv6-capable routers to route toward the locator as an ordinary IPv6 prefix. Absence of this object suppresses the secondary prefix advertisement.
// SetAdvertiseLocatorAsPrefix sets the IsisSRv6AdvertiseLocatorAsPrefix value in the IsisSRv6Locator object
func (obj *isisSRv6Locator) SetAdvertiseLocatorAsPrefix(value IsisSRv6AdvertiseLocatorAsPrefix) IsisSRv6Locator {

	obj.advertiseLocatorAsPrefixHolder = nil
	obj.obj.AdvertiseLocatorAsPrefix = value.msg()

	return obj
}

func (obj *isisSRv6Locator) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface IsisSRv6Locator")
	}

	// Locator is required
	if obj.obj.Locator == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Locator is required field on interface IsisSRv6Locator")
	}
	if obj.obj.Locator != nil {

		err := obj.validateIpv6(obj.Locator())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on IsisSRv6Locator.Locator"))
		}

	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength < 1 || *obj.obj.PrefixLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= IsisSRv6Locator.PrefixLength <= 128 but Got %d", *obj.obj.PrefixLength))
		}

	}

	if obj.obj.Algorithm != nil {

		if *obj.obj.Algorithm > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6Locator.Algorithm <= 255 but Got %d", *obj.obj.Algorithm))
		}

	}

	if obj.obj.Metric != nil {

		if *obj.obj.Metric > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6Locator.Metric <= 16777215 but Got %d", *obj.obj.Metric))
		}

	}

	if obj.obj.MtId != nil {

		if *obj.obj.MtId > 4095 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6Locator.MtId <= 4095 but Got %d", *obj.obj.MtId))
		}

	}

	if len(obj.obj.EndSids) != 0 {

		if set_default {
			obj.EndSids().clearHolderSlice()
			for _, item := range obj.obj.EndSids {
				obj.EndSids().appendHolderSlice(&isisSRv6EndSid{obj: item})
			}
		}
		for _, item := range obj.EndSids().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.AdvertiseLocatorAsPrefix != nil {

		obj.AdvertiseLocatorAsPrefix().validateObj(vObj, set_default)
	}

}

func (obj *isisSRv6Locator) setDefault() {
	if obj.obj.PrefixLength == nil {
		obj.SetPrefixLength(64)
	}
	if obj.obj.Algorithm == nil {
		obj.SetAlgorithm(0)
	}
	if obj.obj.Metric == nil {
		obj.SetMetric(0)
	}
	if obj.obj.RedistributionType == nil {
		obj.SetRedistributionType(IsisSRv6LocatorRedistributionType.UP)

	}
	if obj.obj.DFlag == nil {
		obj.SetDFlag(false)
	}
	if obj.obj.MtId == nil {
		obj.SetMtId(0)
	}

}

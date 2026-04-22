package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspSRv6LocatorTlv *****
type isisLspSRv6LocatorTlv struct {
	validation
	obj                            *otg.IsisLspSRv6LocatorTlv
	marshaller                     marshalIsisLspSRv6LocatorTlv
	unMarshaller                   unMarshalIsisLspSRv6LocatorTlv
	endSidsHolder                  IsisLspSRv6LocatorTlvIsisLspSRv6EndSidIter
	advertiseLocatorAsPrefixHolder IsisLspSRv6AdvertisedLocatorAsPrefix
}

func NewIsisLspSRv6LocatorTlv() IsisLspSRv6LocatorTlv {
	obj := isisLspSRv6LocatorTlv{obj: &otg.IsisLspSRv6LocatorTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspSRv6LocatorTlv) msg() *otg.IsisLspSRv6LocatorTlv {
	return obj.obj
}

func (obj *isisLspSRv6LocatorTlv) setMsg(msg *otg.IsisLspSRv6LocatorTlv) IsisLspSRv6LocatorTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspSRv6LocatorTlv struct {
	obj *isisLspSRv6LocatorTlv
}

type marshalIsisLspSRv6LocatorTlv interface {
	// ToProto marshals IsisLspSRv6LocatorTlv to protobuf object *otg.IsisLspSRv6LocatorTlv
	ToProto() (*otg.IsisLspSRv6LocatorTlv, error)
	// ToPbText marshals IsisLspSRv6LocatorTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspSRv6LocatorTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspSRv6LocatorTlv to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspSRv6LocatorTlv struct {
	obj *isisLspSRv6LocatorTlv
}

type unMarshalIsisLspSRv6LocatorTlv interface {
	// FromProto unmarshals IsisLspSRv6LocatorTlv from protobuf object *otg.IsisLspSRv6LocatorTlv
	FromProto(msg *otg.IsisLspSRv6LocatorTlv) (IsisLspSRv6LocatorTlv, error)
	// FromPbText unmarshals IsisLspSRv6LocatorTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspSRv6LocatorTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspSRv6LocatorTlv from JSON text
	FromJson(value string) error
}

func (obj *isisLspSRv6LocatorTlv) Marshal() marshalIsisLspSRv6LocatorTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspSRv6LocatorTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspSRv6LocatorTlv) Unmarshal() unMarshalIsisLspSRv6LocatorTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspSRv6LocatorTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspSRv6LocatorTlv) ToProto() (*otg.IsisLspSRv6LocatorTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspSRv6LocatorTlv) FromProto(msg *otg.IsisLspSRv6LocatorTlv) (IsisLspSRv6LocatorTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspSRv6LocatorTlv) ToPbText() (string, error) {
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

func (m *unMarshalisisLspSRv6LocatorTlv) FromPbText(value string) error {
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

func (m *marshalisisLspSRv6LocatorTlv) ToYaml() (string, error) {
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

func (m *unMarshalisisLspSRv6LocatorTlv) FromYaml(value string) error {
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

func (m *marshalisisLspSRv6LocatorTlv) ToJson() (string, error) {
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

func (m *unMarshalisisLspSRv6LocatorTlv) FromJson(value string) error {
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

func (obj *isisLspSRv6LocatorTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspSRv6LocatorTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspSRv6LocatorTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspSRv6LocatorTlv) Clone() (IsisLspSRv6LocatorTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspSRv6LocatorTlv()
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

func (obj *isisLspSRv6LocatorTlv) setNil() {
	obj.endSidsHolder = nil
	obj.advertiseLocatorAsPrefixHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspSRv6LocatorTlv is sRv6 Locator TLV (TLV type 27) learned from a received IS-IS LSP. Carries an SRv6 locator prefix that identifies the originating node, the associated IGP algorithm, metric, and a list of locally instantiated End SID Sub-TLVs. Reference: RFC 9352 Section 7.1.
type IsisLspSRv6LocatorTlv interface {
	Validation
	// msg marshals IsisLspSRv6LocatorTlv to protobuf object *otg.IsisLspSRv6LocatorTlv
	// and doesn't set defaults
	msg() *otg.IsisLspSRv6LocatorTlv
	// setMsg unmarshals IsisLspSRv6LocatorTlv from protobuf object *otg.IsisLspSRv6LocatorTlv
	// and doesn't set defaults
	setMsg(*otg.IsisLspSRv6LocatorTlv) IsisLspSRv6LocatorTlv
	// provides marshal interface
	Marshal() marshalIsisLspSRv6LocatorTlv
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspSRv6LocatorTlv
	// validate validates IsisLspSRv6LocatorTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspSRv6LocatorTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Locator returns string, set in IsisLspSRv6LocatorTlv.
	Locator() string
	// SetLocator assigns string provided by user to IsisLspSRv6LocatorTlv
	SetLocator(value string) IsisLspSRv6LocatorTlv
	// HasLocator checks if Locator has been set in IsisLspSRv6LocatorTlv
	HasLocator() bool
	// PrefixLength returns uint32, set in IsisLspSRv6LocatorTlv.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to IsisLspSRv6LocatorTlv
	SetPrefixLength(value uint32) IsisLspSRv6LocatorTlv
	// HasPrefixLength checks if PrefixLength has been set in IsisLspSRv6LocatorTlv
	HasPrefixLength() bool
	// Algorithm returns uint32, set in IsisLspSRv6LocatorTlv.
	Algorithm() uint32
	// SetAlgorithm assigns uint32 provided by user to IsisLspSRv6LocatorTlv
	SetAlgorithm(value uint32) IsisLspSRv6LocatorTlv
	// HasAlgorithm checks if Algorithm has been set in IsisLspSRv6LocatorTlv
	HasAlgorithm() bool
	// Metric returns uint32, set in IsisLspSRv6LocatorTlv.
	Metric() uint32
	// SetMetric assigns uint32 provided by user to IsisLspSRv6LocatorTlv
	SetMetric(value uint32) IsisLspSRv6LocatorTlv
	// HasMetric checks if Metric has been set in IsisLspSRv6LocatorTlv
	HasMetric() bool
	// RedistributionType returns IsisLspSRv6LocatorTlvRedistributionTypeEnum, set in IsisLspSRv6LocatorTlv
	RedistributionType() IsisLspSRv6LocatorTlvRedistributionTypeEnum
	// SetRedistributionType assigns IsisLspSRv6LocatorTlvRedistributionTypeEnum provided by user to IsisLspSRv6LocatorTlv
	SetRedistributionType(value IsisLspSRv6LocatorTlvRedistributionTypeEnum) IsisLspSRv6LocatorTlv
	// HasRedistributionType checks if RedistributionType has been set in IsisLspSRv6LocatorTlv
	HasRedistributionType() bool
	// DFlag returns bool, set in IsisLspSRv6LocatorTlv.
	DFlag() bool
	// SetDFlag assigns bool provided by user to IsisLspSRv6LocatorTlv
	SetDFlag(value bool) IsisLspSRv6LocatorTlv
	// HasDFlag checks if DFlag has been set in IsisLspSRv6LocatorTlv
	HasDFlag() bool
	// MtId returns uint32, set in IsisLspSRv6LocatorTlv.
	MtId() uint32
	// SetMtId assigns uint32 provided by user to IsisLspSRv6LocatorTlv
	SetMtId(value uint32) IsisLspSRv6LocatorTlv
	// HasMtId checks if MtId has been set in IsisLspSRv6LocatorTlv
	HasMtId() bool
	// EndSids returns IsisLspSRv6LocatorTlvIsisLspSRv6EndSidIterIter, set in IsisLspSRv6LocatorTlv
	EndSids() IsisLspSRv6LocatorTlvIsisLspSRv6EndSidIter
	// AdvertiseLocatorAsPrefix returns IsisLspSRv6AdvertisedLocatorAsPrefix, set in IsisLspSRv6LocatorTlv.
	// IsisLspSRv6AdvertisedLocatorAsPrefix is learned state for the secondary IPv6 Reachability prefix advertisement (TLV 236/237) of an SRv6 locator. Present only when the originating router advertised the locator prefix alongside the SRv6 Locator TLV. Reference: RFC 9352 Section 7.1.
	AdvertiseLocatorAsPrefix() IsisLspSRv6AdvertisedLocatorAsPrefix
	// SetAdvertiseLocatorAsPrefix assigns IsisLspSRv6AdvertisedLocatorAsPrefix provided by user to IsisLspSRv6LocatorTlv.
	// IsisLspSRv6AdvertisedLocatorAsPrefix is learned state for the secondary IPv6 Reachability prefix advertisement (TLV 236/237) of an SRv6 locator. Present only when the originating router advertised the locator prefix alongside the SRv6 Locator TLV. Reference: RFC 9352 Section 7.1.
	SetAdvertiseLocatorAsPrefix(value IsisLspSRv6AdvertisedLocatorAsPrefix) IsisLspSRv6LocatorTlv
	// HasAdvertiseLocatorAsPrefix checks if AdvertiseLocatorAsPrefix has been set in IsisLspSRv6LocatorTlv
	HasAdvertiseLocatorAsPrefix() bool
	setNil()
}

// The SRv6 locator IPv6 prefix advertised in this TLV.
// Locator returns a string
func (obj *isisLspSRv6LocatorTlv) Locator() string {

	return *obj.obj.Locator

}

// The SRv6 locator IPv6 prefix advertised in this TLV.
// Locator returns a string
func (obj *isisLspSRv6LocatorTlv) HasLocator() bool {
	return obj.obj.Locator != nil
}

// The SRv6 locator IPv6 prefix advertised in this TLV.
// SetLocator sets the string value in the IsisLspSRv6LocatorTlv object
func (obj *isisLspSRv6LocatorTlv) SetLocator(value string) IsisLspSRv6LocatorTlv {

	obj.obj.Locator = &value
	return obj
}

// Number of significant bits in the locator field (1-128).
// PrefixLength returns a uint32
func (obj *isisLspSRv6LocatorTlv) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// Number of significant bits in the locator field (1-128).
// PrefixLength returns a uint32
func (obj *isisLspSRv6LocatorTlv) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// Number of significant bits in the locator field (1-128).
// SetPrefixLength sets the uint32 value in the IsisLspSRv6LocatorTlv object
func (obj *isisLspSRv6LocatorTlv) SetPrefixLength(value uint32) IsisLspSRv6LocatorTlv {

	obj.obj.PrefixLength = &value
	return obj
}

// IGP algorithm identifier for this locator. 0 = SPF, 1 = Strict SPF, 128-255 = Flex-Algo (RFC 9350, RFC 8665).
// Algorithm returns a uint32
func (obj *isisLspSRv6LocatorTlv) Algorithm() uint32 {

	return *obj.obj.Algorithm

}

// IGP algorithm identifier for this locator. 0 = SPF, 1 = Strict SPF, 128-255 = Flex-Algo (RFC 9350, RFC 8665).
// Algorithm returns a uint32
func (obj *isisLspSRv6LocatorTlv) HasAlgorithm() bool {
	return obj.obj.Algorithm != nil
}

// IGP algorithm identifier for this locator. 0 = SPF, 1 = Strict SPF, 128-255 = Flex-Algo (RFC 9350, RFC 8665).
// SetAlgorithm sets the uint32 value in the IsisLspSRv6LocatorTlv object
func (obj *isisLspSRv6LocatorTlv) SetAlgorithm(value uint32) IsisLspSRv6LocatorTlv {

	obj.obj.Algorithm = &value
	return obj
}

// The metric (cost) associated with this SRv6 locator prefix.
// Metric returns a uint32
func (obj *isisLspSRv6LocatorTlv) Metric() uint32 {

	return *obj.obj.Metric

}

// The metric (cost) associated with this SRv6 locator prefix.
// Metric returns a uint32
func (obj *isisLspSRv6LocatorTlv) HasMetric() bool {
	return obj.obj.Metric != nil
}

// The metric (cost) associated with this SRv6 locator prefix.
// SetMetric sets the uint32 value in the IsisLspSRv6LocatorTlv object
func (obj *isisLspSRv6LocatorTlv) SetMetric(value uint32) IsisLspSRv6LocatorTlv {

	obj.obj.Metric = &value
	return obj
}

type IsisLspSRv6LocatorTlvRedistributionTypeEnum string

// Enum of RedistributionType on IsisLspSRv6LocatorTlv
var IsisLspSRv6LocatorTlvRedistributionType = struct {
	UP   IsisLspSRv6LocatorTlvRedistributionTypeEnum
	DOWN IsisLspSRv6LocatorTlvRedistributionTypeEnum
}{
	UP:   IsisLspSRv6LocatorTlvRedistributionTypeEnum("up"),
	DOWN: IsisLspSRv6LocatorTlvRedistributionTypeEnum("down"),
}

func (obj *isisLspSRv6LocatorTlv) RedistributionType() IsisLspSRv6LocatorTlvRedistributionTypeEnum {
	return IsisLspSRv6LocatorTlvRedistributionTypeEnum(obj.obj.RedistributionType.Enum().String())
}

// Up/Down redistribution bit for this locator advertisement. 'up' = normal advertisement; 'down' = leaked from L2 to L1.
// RedistributionType returns a string
func (obj *isisLspSRv6LocatorTlv) HasRedistributionType() bool {
	return obj.obj.RedistributionType != nil
}

func (obj *isisLspSRv6LocatorTlv) SetRedistributionType(value IsisLspSRv6LocatorTlvRedistributionTypeEnum) IsisLspSRv6LocatorTlv {
	intValue, ok := otg.IsisLspSRv6LocatorTlv_RedistributionType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisLspSRv6LocatorTlvRedistributionTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisLspSRv6LocatorTlv_RedistributionType_Enum(intValue)
	obj.obj.RedistributionType = &enumValue

	return obj
}

// Down bit (D-flag). Set when this locator was leaked from Level 2 to Level 1 (RFC 9352 Section 7.1).
// DFlag returns a bool
func (obj *isisLspSRv6LocatorTlv) DFlag() bool {

	return *obj.obj.DFlag

}

// Down bit (D-flag). Set when this locator was leaked from Level 2 to Level 1 (RFC 9352 Section 7.1).
// DFlag returns a bool
func (obj *isisLspSRv6LocatorTlv) HasDFlag() bool {
	return obj.obj.DFlag != nil
}

// Down bit (D-flag). Set when this locator was leaked from Level 2 to Level 1 (RFC 9352 Section 7.1).
// SetDFlag sets the bool value in the IsisLspSRv6LocatorTlv object
func (obj *isisLspSRv6LocatorTlv) SetDFlag(value bool) IsisLspSRv6LocatorTlv {

	obj.obj.DFlag = &value
	return obj
}

// Multi-Topology ID for this locator. 0 = default topology.
// MtId returns a uint32
func (obj *isisLspSRv6LocatorTlv) MtId() uint32 {

	return *obj.obj.MtId

}

// Multi-Topology ID for this locator. 0 = default topology.
// MtId returns a uint32
func (obj *isisLspSRv6LocatorTlv) HasMtId() bool {
	return obj.obj.MtId != nil
}

// Multi-Topology ID for this locator. 0 = default topology.
// SetMtId sets the uint32 value in the IsisLspSRv6LocatorTlv object
func (obj *isisLspSRv6LocatorTlv) SetMtId(value uint32) IsisLspSRv6LocatorTlv {

	obj.obj.MtId = &value
	return obj
}

// List of SRv6 End SID Sub-TLVs (sub-TLV type 5) carried within this Locator TLV. Each entry describes a node-local SRv6 SID and its endpoint behavior. Reference: RFC 9352 Section 7.2.
// EndSids returns a []IsisLspSRv6EndSid
func (obj *isisLspSRv6LocatorTlv) EndSids() IsisLspSRv6LocatorTlvIsisLspSRv6EndSidIter {
	if len(obj.obj.EndSids) == 0 {
		obj.obj.EndSids = []*otg.IsisLspSRv6EndSid{}
	}
	if obj.endSidsHolder == nil {
		obj.endSidsHolder = newIsisLspSRv6LocatorTlvIsisLspSRv6EndSidIter(&obj.obj.EndSids).setMsg(obj)
	}
	return obj.endSidsHolder
}

type isisLspSRv6LocatorTlvIsisLspSRv6EndSidIter struct {
	obj                    *isisLspSRv6LocatorTlv
	isisLspSRv6EndSidSlice []IsisLspSRv6EndSid
	fieldPtr               *[]*otg.IsisLspSRv6EndSid
}

func newIsisLspSRv6LocatorTlvIsisLspSRv6EndSidIter(ptr *[]*otg.IsisLspSRv6EndSid) IsisLspSRv6LocatorTlvIsisLspSRv6EndSidIter {
	return &isisLspSRv6LocatorTlvIsisLspSRv6EndSidIter{fieldPtr: ptr}
}

type IsisLspSRv6LocatorTlvIsisLspSRv6EndSidIter interface {
	setMsg(*isisLspSRv6LocatorTlv) IsisLspSRv6LocatorTlvIsisLspSRv6EndSidIter
	Items() []IsisLspSRv6EndSid
	Add() IsisLspSRv6EndSid
	Append(items ...IsisLspSRv6EndSid) IsisLspSRv6LocatorTlvIsisLspSRv6EndSidIter
	Set(index int, newObj IsisLspSRv6EndSid) IsisLspSRv6LocatorTlvIsisLspSRv6EndSidIter
	Clear() IsisLspSRv6LocatorTlvIsisLspSRv6EndSidIter
	clearHolderSlice() IsisLspSRv6LocatorTlvIsisLspSRv6EndSidIter
	appendHolderSlice(item IsisLspSRv6EndSid) IsisLspSRv6LocatorTlvIsisLspSRv6EndSidIter
}

func (obj *isisLspSRv6LocatorTlvIsisLspSRv6EndSidIter) setMsg(msg *isisLspSRv6LocatorTlv) IsisLspSRv6LocatorTlvIsisLspSRv6EndSidIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspSRv6EndSid{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspSRv6LocatorTlvIsisLspSRv6EndSidIter) Items() []IsisLspSRv6EndSid {
	return obj.isisLspSRv6EndSidSlice
}

func (obj *isisLspSRv6LocatorTlvIsisLspSRv6EndSidIter) Add() IsisLspSRv6EndSid {
	newObj := &otg.IsisLspSRv6EndSid{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspSRv6EndSid{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspSRv6EndSidSlice = append(obj.isisLspSRv6EndSidSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspSRv6LocatorTlvIsisLspSRv6EndSidIter) Append(items ...IsisLspSRv6EndSid) IsisLspSRv6LocatorTlvIsisLspSRv6EndSidIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspSRv6EndSidSlice = append(obj.isisLspSRv6EndSidSlice, item)
	}
	return obj
}

func (obj *isisLspSRv6LocatorTlvIsisLspSRv6EndSidIter) Set(index int, newObj IsisLspSRv6EndSid) IsisLspSRv6LocatorTlvIsisLspSRv6EndSidIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspSRv6EndSidSlice[index] = newObj
	return obj
}
func (obj *isisLspSRv6LocatorTlvIsisLspSRv6EndSidIter) Clear() IsisLspSRv6LocatorTlvIsisLspSRv6EndSidIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspSRv6EndSid{}
		obj.isisLspSRv6EndSidSlice = []IsisLspSRv6EndSid{}
	}
	return obj
}
func (obj *isisLspSRv6LocatorTlvIsisLspSRv6EndSidIter) clearHolderSlice() IsisLspSRv6LocatorTlvIsisLspSRv6EndSidIter {
	if len(obj.isisLspSRv6EndSidSlice) > 0 {
		obj.isisLspSRv6EndSidSlice = []IsisLspSRv6EndSid{}
	}
	return obj
}
func (obj *isisLspSRv6LocatorTlvIsisLspSRv6EndSidIter) appendHolderSlice(item IsisLspSRv6EndSid) IsisLspSRv6LocatorTlvIsisLspSRv6EndSidIter {
	obj.isisLspSRv6EndSidSlice = append(obj.isisLspSRv6EndSidSlice, item)
	return obj
}

// Present when the originating router also advertised this locator as an IS-IS IPv6 Reachability prefix (TLV 236/237). Carries the prefix attribute flags received in that advertisement.
// AdvertiseLocatorAsPrefix returns a IsisLspSRv6AdvertisedLocatorAsPrefix
func (obj *isisLspSRv6LocatorTlv) AdvertiseLocatorAsPrefix() IsisLspSRv6AdvertisedLocatorAsPrefix {
	if obj.obj.AdvertiseLocatorAsPrefix == nil {
		obj.obj.AdvertiseLocatorAsPrefix = NewIsisLspSRv6AdvertisedLocatorAsPrefix().msg()
	}
	if obj.advertiseLocatorAsPrefixHolder == nil {
		obj.advertiseLocatorAsPrefixHolder = &isisLspSRv6AdvertisedLocatorAsPrefix{obj: obj.obj.AdvertiseLocatorAsPrefix}
	}
	return obj.advertiseLocatorAsPrefixHolder
}

// Present when the originating router also advertised this locator as an IS-IS IPv6 Reachability prefix (TLV 236/237). Carries the prefix attribute flags received in that advertisement.
// AdvertiseLocatorAsPrefix returns a IsisLspSRv6AdvertisedLocatorAsPrefix
func (obj *isisLspSRv6LocatorTlv) HasAdvertiseLocatorAsPrefix() bool {
	return obj.obj.AdvertiseLocatorAsPrefix != nil
}

// Present when the originating router also advertised this locator as an IS-IS IPv6 Reachability prefix (TLV 236/237). Carries the prefix attribute flags received in that advertisement.
// SetAdvertiseLocatorAsPrefix sets the IsisLspSRv6AdvertisedLocatorAsPrefix value in the IsisLspSRv6LocatorTlv object
func (obj *isisLspSRv6LocatorTlv) SetAdvertiseLocatorAsPrefix(value IsisLspSRv6AdvertisedLocatorAsPrefix) IsisLspSRv6LocatorTlv {

	obj.advertiseLocatorAsPrefixHolder = nil
	obj.obj.AdvertiseLocatorAsPrefix = value.msg()

	return obj
}

func (obj *isisLspSRv6LocatorTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Locator != nil {

		err := obj.validateIpv6(obj.Locator())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on IsisLspSRv6LocatorTlv.Locator"))
		}

	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength < 1 || *obj.obj.PrefixLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= IsisLspSRv6LocatorTlv.PrefixLength <= 128 but Got %d", *obj.obj.PrefixLength))
		}

	}

	if obj.obj.Algorithm != nil {

		if *obj.obj.Algorithm > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisLspSRv6LocatorTlv.Algorithm <= 255 but Got %d", *obj.obj.Algorithm))
		}

	}

	if obj.obj.MtId != nil {

		if *obj.obj.MtId > 4095 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisLspSRv6LocatorTlv.MtId <= 4095 but Got %d", *obj.obj.MtId))
		}

	}

	if len(obj.obj.EndSids) != 0 {

		if set_default {
			obj.EndSids().clearHolderSlice()
			for _, item := range obj.obj.EndSids {
				obj.EndSids().appendHolderSlice(&isisLspSRv6EndSid{obj: item})
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

func (obj *isisLspSRv6LocatorTlv) setDefault() {

}

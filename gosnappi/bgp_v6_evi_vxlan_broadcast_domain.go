package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpV6EviVxlanBroadcastDomain *****
type bgpV6EviVxlanBroadcastDomain struct {
	validation
	obj               *otg.BgpV6EviVxlanBroadcastDomain
	marshaller        marshalBgpV6EviVxlanBroadcastDomain
	unMarshaller      unMarshalBgpV6EviVxlanBroadcastDomain
	cmacIpRangeHolder BgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter
}

func NewBgpV6EviVxlanBroadcastDomain() BgpV6EviVxlanBroadcastDomain {
	obj := bgpV6EviVxlanBroadcastDomain{obj: &otg.BgpV6EviVxlanBroadcastDomain{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpV6EviVxlanBroadcastDomain) msg() *otg.BgpV6EviVxlanBroadcastDomain {
	return obj.obj
}

func (obj *bgpV6EviVxlanBroadcastDomain) setMsg(msg *otg.BgpV6EviVxlanBroadcastDomain) BgpV6EviVxlanBroadcastDomain {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpV6EviVxlanBroadcastDomain struct {
	obj *bgpV6EviVxlanBroadcastDomain
}

type marshalBgpV6EviVxlanBroadcastDomain interface {
	// ToProto marshals BgpV6EviVxlanBroadcastDomain to protobuf object *otg.BgpV6EviVxlanBroadcastDomain
	ToProto() (*otg.BgpV6EviVxlanBroadcastDomain, error)
	// ToPbText marshals BgpV6EviVxlanBroadcastDomain to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpV6EviVxlanBroadcastDomain to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpV6EviVxlanBroadcastDomain to JSON text
	ToJson() (string, error)
}

type unMarshalbgpV6EviVxlanBroadcastDomain struct {
	obj *bgpV6EviVxlanBroadcastDomain
}

type unMarshalBgpV6EviVxlanBroadcastDomain interface {
	// FromProto unmarshals BgpV6EviVxlanBroadcastDomain from protobuf object *otg.BgpV6EviVxlanBroadcastDomain
	FromProto(msg *otg.BgpV6EviVxlanBroadcastDomain) (BgpV6EviVxlanBroadcastDomain, error)
	// FromPbText unmarshals BgpV6EviVxlanBroadcastDomain from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpV6EviVxlanBroadcastDomain from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpV6EviVxlanBroadcastDomain from JSON text
	FromJson(value string) error
}

func (obj *bgpV6EviVxlanBroadcastDomain) Marshal() marshalBgpV6EviVxlanBroadcastDomain {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpV6EviVxlanBroadcastDomain{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpV6EviVxlanBroadcastDomain) Unmarshal() unMarshalBgpV6EviVxlanBroadcastDomain {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpV6EviVxlanBroadcastDomain{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpV6EviVxlanBroadcastDomain) ToProto() (*otg.BgpV6EviVxlanBroadcastDomain, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpV6EviVxlanBroadcastDomain) FromProto(msg *otg.BgpV6EviVxlanBroadcastDomain) (BgpV6EviVxlanBroadcastDomain, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpV6EviVxlanBroadcastDomain) ToPbText() (string, error) {
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

func (m *unMarshalbgpV6EviVxlanBroadcastDomain) FromPbText(value string) error {
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

func (m *marshalbgpV6EviVxlanBroadcastDomain) ToYaml() (string, error) {
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

func (m *unMarshalbgpV6EviVxlanBroadcastDomain) FromYaml(value string) error {
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

func (m *marshalbgpV6EviVxlanBroadcastDomain) ToJson() (string, error) {
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

func (m *unMarshalbgpV6EviVxlanBroadcastDomain) FromJson(value string) error {
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

func (obj *bgpV6EviVxlanBroadcastDomain) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpV6EviVxlanBroadcastDomain) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpV6EviVxlanBroadcastDomain) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpV6EviVxlanBroadcastDomain) Clone() (BgpV6EviVxlanBroadcastDomain, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpV6EviVxlanBroadcastDomain()
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

func (obj *bgpV6EviVxlanBroadcastDomain) setNil() {
	obj.cmacIpRangeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpV6EviVxlanBroadcastDomain is configuration for Broadcast Domains per EVI.
type BgpV6EviVxlanBroadcastDomain interface {
	Validation
	// msg marshals BgpV6EviVxlanBroadcastDomain to protobuf object *otg.BgpV6EviVxlanBroadcastDomain
	// and doesn't set defaults
	msg() *otg.BgpV6EviVxlanBroadcastDomain
	// setMsg unmarshals BgpV6EviVxlanBroadcastDomain from protobuf object *otg.BgpV6EviVxlanBroadcastDomain
	// and doesn't set defaults
	setMsg(*otg.BgpV6EviVxlanBroadcastDomain) BgpV6EviVxlanBroadcastDomain
	// provides marshal interface
	Marshal() marshalBgpV6EviVxlanBroadcastDomain
	// provides unmarshal interface
	Unmarshal() unMarshalBgpV6EviVxlanBroadcastDomain
	// validate validates BgpV6EviVxlanBroadcastDomain
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpV6EviVxlanBroadcastDomain, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CmacIpRange returns BgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIterIter, set in BgpV6EviVxlanBroadcastDomain
	CmacIpRange() BgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter
	// EthernetTagId returns uint32, set in BgpV6EviVxlanBroadcastDomain.
	EthernetTagId() uint32
	// SetEthernetTagId assigns uint32 provided by user to BgpV6EviVxlanBroadcastDomain
	SetEthernetTagId(value uint32) BgpV6EviVxlanBroadcastDomain
	// HasEthernetTagId checks if EthernetTagId has been set in BgpV6EviVxlanBroadcastDomain
	HasEthernetTagId() bool
	// VlanAwareService returns bool, set in BgpV6EviVxlanBroadcastDomain.
	VlanAwareService() bool
	// SetVlanAwareService assigns bool provided by user to BgpV6EviVxlanBroadcastDomain
	SetVlanAwareService(value bool) BgpV6EviVxlanBroadcastDomain
	// HasVlanAwareService checks if VlanAwareService has been set in BgpV6EviVxlanBroadcastDomain
	HasVlanAwareService() bool
	setNil()
}

// This contains the list of Customer MAC/IP Ranges to be configured per Broadcast Domain.
//
// Advertises following route -
// Type 2 - MAC/IP Advertisement Route.
// CmacIpRange returns a []BgpCMacIpRange
func (obj *bgpV6EviVxlanBroadcastDomain) CmacIpRange() BgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter {
	if len(obj.obj.CmacIpRange) == 0 {
		obj.obj.CmacIpRange = []*otg.BgpCMacIpRange{}
	}
	if obj.cmacIpRangeHolder == nil {
		obj.cmacIpRangeHolder = newBgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter(&obj.obj.CmacIpRange).setMsg(obj)
	}
	return obj.cmacIpRangeHolder
}

type bgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter struct {
	obj                 *bgpV6EviVxlanBroadcastDomain
	bgpCMacIpRangeSlice []BgpCMacIpRange
	fieldPtr            *[]*otg.BgpCMacIpRange
}

func newBgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter(ptr *[]*otg.BgpCMacIpRange) BgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter {
	return &bgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter{fieldPtr: ptr}
}

type BgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter interface {
	setMsg(*bgpV6EviVxlanBroadcastDomain) BgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter
	Items() []BgpCMacIpRange
	Add() BgpCMacIpRange
	Append(items ...BgpCMacIpRange) BgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter
	Set(index int, newObj BgpCMacIpRange) BgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter
	Clear() BgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter
	clearHolderSlice() BgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter
	appendHolderSlice(item BgpCMacIpRange) BgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter
}

func (obj *bgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter) setMsg(msg *bgpV6EviVxlanBroadcastDomain) BgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpCMacIpRange{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter) Items() []BgpCMacIpRange {
	return obj.bgpCMacIpRangeSlice
}

func (obj *bgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter) Add() BgpCMacIpRange {
	newObj := &otg.BgpCMacIpRange{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpCMacIpRange{obj: newObj}
	newLibObj.setDefault()
	obj.bgpCMacIpRangeSlice = append(obj.bgpCMacIpRangeSlice, newLibObj)
	return newLibObj
}

func (obj *bgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter) Append(items ...BgpCMacIpRange) BgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpCMacIpRangeSlice = append(obj.bgpCMacIpRangeSlice, item)
	}
	return obj
}

func (obj *bgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter) Set(index int, newObj BgpCMacIpRange) BgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpCMacIpRangeSlice[index] = newObj
	return obj
}
func (obj *bgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter) Clear() BgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpCMacIpRange{}
		obj.bgpCMacIpRangeSlice = []BgpCMacIpRange{}
	}
	return obj
}
func (obj *bgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter) clearHolderSlice() BgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter {
	if len(obj.bgpCMacIpRangeSlice) > 0 {
		obj.bgpCMacIpRangeSlice = []BgpCMacIpRange{}
	}
	return obj
}
func (obj *bgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter) appendHolderSlice(item BgpCMacIpRange) BgpV6EviVxlanBroadcastDomainBgpCMacIpRangeIter {
	obj.bgpCMacIpRangeSlice = append(obj.bgpCMacIpRangeSlice, item)
	return obj
}

// The Ethernet Tag ID of the Broadcast Domain.
// EthernetTagId returns a uint32
func (obj *bgpV6EviVxlanBroadcastDomain) EthernetTagId() uint32 {

	return *obj.obj.EthernetTagId

}

// The Ethernet Tag ID of the Broadcast Domain.
// EthernetTagId returns a uint32
func (obj *bgpV6EviVxlanBroadcastDomain) HasEthernetTagId() bool {
	return obj.obj.EthernetTagId != nil
}

// The Ethernet Tag ID of the Broadcast Domain.
// SetEthernetTagId sets the uint32 value in the BgpV6EviVxlanBroadcastDomain object
func (obj *bgpV6EviVxlanBroadcastDomain) SetEthernetTagId(value uint32) BgpV6EviVxlanBroadcastDomain {

	obj.obj.EthernetTagId = &value
	return obj
}

// VLAN-Aware service to be enabled or disabled.
// VlanAwareService returns a bool
func (obj *bgpV6EviVxlanBroadcastDomain) VlanAwareService() bool {

	return *obj.obj.VlanAwareService

}

// VLAN-Aware service to be enabled or disabled.
// VlanAwareService returns a bool
func (obj *bgpV6EviVxlanBroadcastDomain) HasVlanAwareService() bool {
	return obj.obj.VlanAwareService != nil
}

// VLAN-Aware service to be enabled or disabled.
// SetVlanAwareService sets the bool value in the BgpV6EviVxlanBroadcastDomain object
func (obj *bgpV6EviVxlanBroadcastDomain) SetVlanAwareService(value bool) BgpV6EviVxlanBroadcastDomain {

	obj.obj.VlanAwareService = &value
	return obj
}

func (obj *bgpV6EviVxlanBroadcastDomain) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.CmacIpRange) != 0 {

		if set_default {
			obj.CmacIpRange().clearHolderSlice()
			for _, item := range obj.obj.CmacIpRange {
				obj.CmacIpRange().appendHolderSlice(&bgpCMacIpRange{obj: item})
			}
		}
		for _, item := range obj.CmacIpRange().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bgpV6EviVxlanBroadcastDomain) setDefault() {
	if obj.obj.EthernetTagId == nil {
		obj.SetEthernetTagId(0)
	}
	if obj.obj.VlanAwareService == nil {
		obj.SetVlanAwareService(false)
	}

}

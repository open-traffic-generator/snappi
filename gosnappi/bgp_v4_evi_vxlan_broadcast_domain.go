package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpV4EviVxlanBroadcastDomain *****
type bgpV4EviVxlanBroadcastDomain struct {
	validation
	obj               *otg.BgpV4EviVxlanBroadcastDomain
	marshaller        marshalBgpV4EviVxlanBroadcastDomain
	unMarshaller      unMarshalBgpV4EviVxlanBroadcastDomain
	cmacIpRangeHolder BgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter
}

func NewBgpV4EviVxlanBroadcastDomain() BgpV4EviVxlanBroadcastDomain {
	obj := bgpV4EviVxlanBroadcastDomain{obj: &otg.BgpV4EviVxlanBroadcastDomain{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpV4EviVxlanBroadcastDomain) msg() *otg.BgpV4EviVxlanBroadcastDomain {
	return obj.obj
}

func (obj *bgpV4EviVxlanBroadcastDomain) setMsg(msg *otg.BgpV4EviVxlanBroadcastDomain) BgpV4EviVxlanBroadcastDomain {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpV4EviVxlanBroadcastDomain struct {
	obj *bgpV4EviVxlanBroadcastDomain
}

type marshalBgpV4EviVxlanBroadcastDomain interface {
	// ToProto marshals BgpV4EviVxlanBroadcastDomain to protobuf object *otg.BgpV4EviVxlanBroadcastDomain
	ToProto() (*otg.BgpV4EviVxlanBroadcastDomain, error)
	// ToPbText marshals BgpV4EviVxlanBroadcastDomain to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpV4EviVxlanBroadcastDomain to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpV4EviVxlanBroadcastDomain to JSON text
	ToJson() (string, error)
}

type unMarshalbgpV4EviVxlanBroadcastDomain struct {
	obj *bgpV4EviVxlanBroadcastDomain
}

type unMarshalBgpV4EviVxlanBroadcastDomain interface {
	// FromProto unmarshals BgpV4EviVxlanBroadcastDomain from protobuf object *otg.BgpV4EviVxlanBroadcastDomain
	FromProto(msg *otg.BgpV4EviVxlanBroadcastDomain) (BgpV4EviVxlanBroadcastDomain, error)
	// FromPbText unmarshals BgpV4EviVxlanBroadcastDomain from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpV4EviVxlanBroadcastDomain from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpV4EviVxlanBroadcastDomain from JSON text
	FromJson(value string) error
}

func (obj *bgpV4EviVxlanBroadcastDomain) Marshal() marshalBgpV4EviVxlanBroadcastDomain {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpV4EviVxlanBroadcastDomain{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpV4EviVxlanBroadcastDomain) Unmarshal() unMarshalBgpV4EviVxlanBroadcastDomain {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpV4EviVxlanBroadcastDomain{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpV4EviVxlanBroadcastDomain) ToProto() (*otg.BgpV4EviVxlanBroadcastDomain, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpV4EviVxlanBroadcastDomain) FromProto(msg *otg.BgpV4EviVxlanBroadcastDomain) (BgpV4EviVxlanBroadcastDomain, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpV4EviVxlanBroadcastDomain) ToPbText() (string, error) {
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

func (m *unMarshalbgpV4EviVxlanBroadcastDomain) FromPbText(value string) error {
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

func (m *marshalbgpV4EviVxlanBroadcastDomain) ToYaml() (string, error) {
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

func (m *unMarshalbgpV4EviVxlanBroadcastDomain) FromYaml(value string) error {
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

func (m *marshalbgpV4EviVxlanBroadcastDomain) ToJson() (string, error) {
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

func (m *unMarshalbgpV4EviVxlanBroadcastDomain) FromJson(value string) error {
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

func (obj *bgpV4EviVxlanBroadcastDomain) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpV4EviVxlanBroadcastDomain) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpV4EviVxlanBroadcastDomain) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpV4EviVxlanBroadcastDomain) Clone() (BgpV4EviVxlanBroadcastDomain, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpV4EviVxlanBroadcastDomain()
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

func (obj *bgpV4EviVxlanBroadcastDomain) setNil() {
	obj.cmacIpRangeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpV4EviVxlanBroadcastDomain is configuration for Broadcast Domains per EVI.
type BgpV4EviVxlanBroadcastDomain interface {
	Validation
	// msg marshals BgpV4EviVxlanBroadcastDomain to protobuf object *otg.BgpV4EviVxlanBroadcastDomain
	// and doesn't set defaults
	msg() *otg.BgpV4EviVxlanBroadcastDomain
	// setMsg unmarshals BgpV4EviVxlanBroadcastDomain from protobuf object *otg.BgpV4EviVxlanBroadcastDomain
	// and doesn't set defaults
	setMsg(*otg.BgpV4EviVxlanBroadcastDomain) BgpV4EviVxlanBroadcastDomain
	// provides marshal interface
	Marshal() marshalBgpV4EviVxlanBroadcastDomain
	// provides unmarshal interface
	Unmarshal() unMarshalBgpV4EviVxlanBroadcastDomain
	// validate validates BgpV4EviVxlanBroadcastDomain
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpV4EviVxlanBroadcastDomain, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CmacIpRange returns BgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIterIter, set in BgpV4EviVxlanBroadcastDomain
	CmacIpRange() BgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter
	// EthernetTagId returns uint32, set in BgpV4EviVxlanBroadcastDomain.
	EthernetTagId() uint32
	// SetEthernetTagId assigns uint32 provided by user to BgpV4EviVxlanBroadcastDomain
	SetEthernetTagId(value uint32) BgpV4EviVxlanBroadcastDomain
	// HasEthernetTagId checks if EthernetTagId has been set in BgpV4EviVxlanBroadcastDomain
	HasEthernetTagId() bool
	// VlanAwareService returns bool, set in BgpV4EviVxlanBroadcastDomain.
	VlanAwareService() bool
	// SetVlanAwareService assigns bool provided by user to BgpV4EviVxlanBroadcastDomain
	SetVlanAwareService(value bool) BgpV4EviVxlanBroadcastDomain
	// HasVlanAwareService checks if VlanAwareService has been set in BgpV4EviVxlanBroadcastDomain
	HasVlanAwareService() bool
	setNil()
}

// This contains the list of Customer MAC/IP Ranges to be configured per Broadcast Domain.
//
// Advertises following route -
// Type 2 - MAC/IP Advertisement Route.
// CmacIpRange returns a []BgpCMacIpRange
func (obj *bgpV4EviVxlanBroadcastDomain) CmacIpRange() BgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter {
	if len(obj.obj.CmacIpRange) == 0 {
		obj.obj.CmacIpRange = []*otg.BgpCMacIpRange{}
	}
	if obj.cmacIpRangeHolder == nil {
		obj.cmacIpRangeHolder = newBgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter(&obj.obj.CmacIpRange).setMsg(obj)
	}
	return obj.cmacIpRangeHolder
}

type bgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter struct {
	obj                 *bgpV4EviVxlanBroadcastDomain
	bgpCMacIpRangeSlice []BgpCMacIpRange
	fieldPtr            *[]*otg.BgpCMacIpRange
}

func newBgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter(ptr *[]*otg.BgpCMacIpRange) BgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter {
	return &bgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter{fieldPtr: ptr}
}

type BgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter interface {
	setMsg(*bgpV4EviVxlanBroadcastDomain) BgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter
	Items() []BgpCMacIpRange
	Add() BgpCMacIpRange
	Append(items ...BgpCMacIpRange) BgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter
	Set(index int, newObj BgpCMacIpRange) BgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter
	Clear() BgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter
	clearHolderSlice() BgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter
	appendHolderSlice(item BgpCMacIpRange) BgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter
}

func (obj *bgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter) setMsg(msg *bgpV4EviVxlanBroadcastDomain) BgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpCMacIpRange{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter) Items() []BgpCMacIpRange {
	return obj.bgpCMacIpRangeSlice
}

func (obj *bgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter) Add() BgpCMacIpRange {
	newObj := &otg.BgpCMacIpRange{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpCMacIpRange{obj: newObj}
	newLibObj.setDefault()
	obj.bgpCMacIpRangeSlice = append(obj.bgpCMacIpRangeSlice, newLibObj)
	return newLibObj
}

func (obj *bgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter) Append(items ...BgpCMacIpRange) BgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpCMacIpRangeSlice = append(obj.bgpCMacIpRangeSlice, item)
	}
	return obj
}

func (obj *bgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter) Set(index int, newObj BgpCMacIpRange) BgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpCMacIpRangeSlice[index] = newObj
	return obj
}
func (obj *bgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter) Clear() BgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpCMacIpRange{}
		obj.bgpCMacIpRangeSlice = []BgpCMacIpRange{}
	}
	return obj
}
func (obj *bgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter) clearHolderSlice() BgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter {
	if len(obj.bgpCMacIpRangeSlice) > 0 {
		obj.bgpCMacIpRangeSlice = []BgpCMacIpRange{}
	}
	return obj
}
func (obj *bgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter) appendHolderSlice(item BgpCMacIpRange) BgpV4EviVxlanBroadcastDomainBgpCMacIpRangeIter {
	obj.bgpCMacIpRangeSlice = append(obj.bgpCMacIpRangeSlice, item)
	return obj
}

// The Ethernet Tag ID of the Broadcast Domain.
// EthernetTagId returns a uint32
func (obj *bgpV4EviVxlanBroadcastDomain) EthernetTagId() uint32 {

	return *obj.obj.EthernetTagId

}

// The Ethernet Tag ID of the Broadcast Domain.
// EthernetTagId returns a uint32
func (obj *bgpV4EviVxlanBroadcastDomain) HasEthernetTagId() bool {
	return obj.obj.EthernetTagId != nil
}

// The Ethernet Tag ID of the Broadcast Domain.
// SetEthernetTagId sets the uint32 value in the BgpV4EviVxlanBroadcastDomain object
func (obj *bgpV4EviVxlanBroadcastDomain) SetEthernetTagId(value uint32) BgpV4EviVxlanBroadcastDomain {

	obj.obj.EthernetTagId = &value
	return obj
}

// VLAN-Aware service to be enabled or disabled.
// VlanAwareService returns a bool
func (obj *bgpV4EviVxlanBroadcastDomain) VlanAwareService() bool {

	return *obj.obj.VlanAwareService

}

// VLAN-Aware service to be enabled or disabled.
// VlanAwareService returns a bool
func (obj *bgpV4EviVxlanBroadcastDomain) HasVlanAwareService() bool {
	return obj.obj.VlanAwareService != nil
}

// VLAN-Aware service to be enabled or disabled.
// SetVlanAwareService sets the bool value in the BgpV4EviVxlanBroadcastDomain object
func (obj *bgpV4EviVxlanBroadcastDomain) SetVlanAwareService(value bool) BgpV4EviVxlanBroadcastDomain {

	obj.obj.VlanAwareService = &value
	return obj
}

func (obj *bgpV4EviVxlanBroadcastDomain) validateObj(vObj *validation, set_default bool) {
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

func (obj *bgpV4EviVxlanBroadcastDomain) setDefault() {
	if obj.obj.EthernetTagId == nil {
		obj.SetEthernetTagId(0)
	}
	if obj.obj.VlanAwareService == nil {
		obj.SetVlanAwareService(false)
	}

}

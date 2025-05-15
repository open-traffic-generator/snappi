package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixStateRequest *****
type bgpPrefixStateRequest struct {
	validation
	obj                      *otg.BgpPrefixStateRequest
	marshaller               marshalBgpPrefixStateRequest
	unMarshaller             unMarshalBgpPrefixStateRequest
	ipv4UnicastFiltersHolder BgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter
	ipv6UnicastFiltersHolder BgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter
}

func NewBgpPrefixStateRequest() BgpPrefixStateRequest {
	obj := bgpPrefixStateRequest{obj: &otg.BgpPrefixStateRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixStateRequest) msg() *otg.BgpPrefixStateRequest {
	return obj.obj
}

func (obj *bgpPrefixStateRequest) setMsg(msg *otg.BgpPrefixStateRequest) BgpPrefixStateRequest {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixStateRequest struct {
	obj *bgpPrefixStateRequest
}

type marshalBgpPrefixStateRequest interface {
	// ToProto marshals BgpPrefixStateRequest to protobuf object *otg.BgpPrefixStateRequest
	ToProto() (*otg.BgpPrefixStateRequest, error)
	// ToPbText marshals BgpPrefixStateRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixStateRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixStateRequest to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixStateRequest struct {
	obj *bgpPrefixStateRequest
}

type unMarshalBgpPrefixStateRequest interface {
	// FromProto unmarshals BgpPrefixStateRequest from protobuf object *otg.BgpPrefixStateRequest
	FromProto(msg *otg.BgpPrefixStateRequest) (BgpPrefixStateRequest, error)
	// FromPbText unmarshals BgpPrefixStateRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixStateRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixStateRequest from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixStateRequest) Marshal() marshalBgpPrefixStateRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixStateRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixStateRequest) Unmarshal() unMarshalBgpPrefixStateRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixStateRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixStateRequest) ToProto() (*otg.BgpPrefixStateRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixStateRequest) FromProto(msg *otg.BgpPrefixStateRequest) (BgpPrefixStateRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixStateRequest) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixStateRequest) FromPbText(value string) error {
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

func (m *marshalbgpPrefixStateRequest) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixStateRequest) FromYaml(value string) error {
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

func (m *marshalbgpPrefixStateRequest) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixStateRequest) FromJson(value string) error {
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

func (obj *bgpPrefixStateRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixStateRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixStateRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixStateRequest) Clone() (BgpPrefixStateRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixStateRequest()
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

func (obj *bgpPrefixStateRequest) setNil() {
	obj.ipv4UnicastFiltersHolder = nil
	obj.ipv6UnicastFiltersHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpPrefixStateRequest is the request to retrieve BGP peer prefix information.
type BgpPrefixStateRequest interface {
	Validation
	// msg marshals BgpPrefixStateRequest to protobuf object *otg.BgpPrefixStateRequest
	// and doesn't set defaults
	msg() *otg.BgpPrefixStateRequest
	// setMsg unmarshals BgpPrefixStateRequest from protobuf object *otg.BgpPrefixStateRequest
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixStateRequest) BgpPrefixStateRequest
	// provides marshal interface
	Marshal() marshalBgpPrefixStateRequest
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixStateRequest
	// validate validates BgpPrefixStateRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixStateRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// BgpPeerNames returns []string, set in BgpPrefixStateRequest.
	BgpPeerNames() []string
	// SetBgpPeerNames assigns []string provided by user to BgpPrefixStateRequest
	SetBgpPeerNames(value []string) BgpPrefixStateRequest
	// PrefixFilters returns []BgpPrefixStateRequestPrefixFiltersEnum, set in BgpPrefixStateRequest
	PrefixFilters() []BgpPrefixStateRequestPrefixFiltersEnum
	// SetPrefixFilters assigns []BgpPrefixStateRequestPrefixFiltersEnum provided by user to BgpPrefixStateRequest
	SetPrefixFilters(value []BgpPrefixStateRequestPrefixFiltersEnum) BgpPrefixStateRequest
	// Ipv4UnicastFilters returns BgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIterIter, set in BgpPrefixStateRequest
	Ipv4UnicastFilters() BgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter
	// Ipv6UnicastFilters returns BgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIterIter, set in BgpPrefixStateRequest
	Ipv6UnicastFilters() BgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter
	setNil()
}

// The names of BGP peers for which prefix information will be retrieved. If no names are specified then the results will contain prefix information for all configured BGP peers.
//
// x-constraint:
// - /components/schemas/Bgp.V4Peer/properties/name
// - /components/schemas/Bgp.V6Peer/properties/name
//
// BgpPeerNames returns a []string
func (obj *bgpPrefixStateRequest) BgpPeerNames() []string {
	if obj.obj.BgpPeerNames == nil {
		obj.obj.BgpPeerNames = make([]string, 0)
	}
	return obj.obj.BgpPeerNames
}

// The names of BGP peers for which prefix information will be retrieved. If no names are specified then the results will contain prefix information for all configured BGP peers.
//
// x-constraint:
// - /components/schemas/Bgp.V4Peer/properties/name
// - /components/schemas/Bgp.V6Peer/properties/name
//
// SetBgpPeerNames sets the []string value in the BgpPrefixStateRequest object
func (obj *bgpPrefixStateRequest) SetBgpPeerNames(value []string) BgpPrefixStateRequest {

	if obj.obj.BgpPeerNames == nil {
		obj.obj.BgpPeerNames = make([]string, 0)
	}
	obj.obj.BgpPeerNames = value

	return obj
}

type BgpPrefixStateRequestPrefixFiltersEnum string

// Enum of PrefixFilters on BgpPrefixStateRequest
var BgpPrefixStateRequestPrefixFilters = struct {
	IPV4_UNICAST BgpPrefixStateRequestPrefixFiltersEnum
	IPV6_UNICAST BgpPrefixStateRequestPrefixFiltersEnum
}{
	IPV4_UNICAST: BgpPrefixStateRequestPrefixFiltersEnum("ipv4_unicast"),
	IPV6_UNICAST: BgpPrefixStateRequestPrefixFiltersEnum("ipv6_unicast"),
}

func (obj *bgpPrefixStateRequest) PrefixFilters() []BgpPrefixStateRequestPrefixFiltersEnum {
	items := []BgpPrefixStateRequestPrefixFiltersEnum{}
	for _, item := range obj.obj.PrefixFilters {
		items = append(items, BgpPrefixStateRequestPrefixFiltersEnum(item.String()))
	}
	return items
}

// Specify which prefixes to return. If the list is empty or missing then all prefixes will be returned.
// SetPrefixFilters sets the []string value in the BgpPrefixStateRequest object
func (obj *bgpPrefixStateRequest) SetPrefixFilters(value []BgpPrefixStateRequestPrefixFiltersEnum) BgpPrefixStateRequest {

	items := []otg.BgpPrefixStateRequest_PrefixFilters_Enum{}
	for _, item := range value {
		intValue := otg.BgpPrefixStateRequest_PrefixFilters_Enum_value[string(item)]
		items = append(items, otg.BgpPrefixStateRequest_PrefixFilters_Enum(intValue))
	}
	obj.obj.PrefixFilters = items
	return obj
}

// The IPv4 unicast results can be filtered by specifying additional prefix search criteria. If the ipv4_unicast_filters property is missing or empty then all IPv4 prefixes will be returned.
// Ipv4UnicastFilters returns a []BgpPrefixIpv4UnicastFilter
func (obj *bgpPrefixStateRequest) Ipv4UnicastFilters() BgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter {
	if len(obj.obj.Ipv4UnicastFilters) == 0 {
		obj.obj.Ipv4UnicastFilters = []*otg.BgpPrefixIpv4UnicastFilter{}
	}
	if obj.ipv4UnicastFiltersHolder == nil {
		obj.ipv4UnicastFiltersHolder = newBgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter(&obj.obj.Ipv4UnicastFilters).setMsg(obj)
	}
	return obj.ipv4UnicastFiltersHolder
}

type bgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter struct {
	obj                             *bgpPrefixStateRequest
	bgpPrefixIpv4UnicastFilterSlice []BgpPrefixIpv4UnicastFilter
	fieldPtr                        *[]*otg.BgpPrefixIpv4UnicastFilter
}

func newBgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter(ptr *[]*otg.BgpPrefixIpv4UnicastFilter) BgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter {
	return &bgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter{fieldPtr: ptr}
}

type BgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter interface {
	setMsg(*bgpPrefixStateRequest) BgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter
	Items() []BgpPrefixIpv4UnicastFilter
	Add() BgpPrefixIpv4UnicastFilter
	Append(items ...BgpPrefixIpv4UnicastFilter) BgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter
	Set(index int, newObj BgpPrefixIpv4UnicastFilter) BgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter
	Clear() BgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter
	clearHolderSlice() BgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter
	appendHolderSlice(item BgpPrefixIpv4UnicastFilter) BgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter
}

func (obj *bgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter) setMsg(msg *bgpPrefixStateRequest) BgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpPrefixIpv4UnicastFilter{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter) Items() []BgpPrefixIpv4UnicastFilter {
	return obj.bgpPrefixIpv4UnicastFilterSlice
}

func (obj *bgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter) Add() BgpPrefixIpv4UnicastFilter {
	newObj := &otg.BgpPrefixIpv4UnicastFilter{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpPrefixIpv4UnicastFilter{obj: newObj}
	newLibObj.setDefault()
	obj.bgpPrefixIpv4UnicastFilterSlice = append(obj.bgpPrefixIpv4UnicastFilterSlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter) Append(items ...BgpPrefixIpv4UnicastFilter) BgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpPrefixIpv4UnicastFilterSlice = append(obj.bgpPrefixIpv4UnicastFilterSlice, item)
	}
	return obj
}

func (obj *bgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter) Set(index int, newObj BgpPrefixIpv4UnicastFilter) BgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpPrefixIpv4UnicastFilterSlice[index] = newObj
	return obj
}
func (obj *bgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter) Clear() BgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpPrefixIpv4UnicastFilter{}
		obj.bgpPrefixIpv4UnicastFilterSlice = []BgpPrefixIpv4UnicastFilter{}
	}
	return obj
}
func (obj *bgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter) clearHolderSlice() BgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter {
	if len(obj.bgpPrefixIpv4UnicastFilterSlice) > 0 {
		obj.bgpPrefixIpv4UnicastFilterSlice = []BgpPrefixIpv4UnicastFilter{}
	}
	return obj
}
func (obj *bgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter) appendHolderSlice(item BgpPrefixIpv4UnicastFilter) BgpPrefixStateRequestBgpPrefixIpv4UnicastFilterIter {
	obj.bgpPrefixIpv4UnicastFilterSlice = append(obj.bgpPrefixIpv4UnicastFilterSlice, item)
	return obj
}

// The IPv6 unicast results can be filtered by specifying additional prefix search criteria. If the ipv6_unicast_filters property is missing or empty then all IPv6 prefixes will be returned.
// Ipv6UnicastFilters returns a []BgpPrefixIpv6UnicastFilter
func (obj *bgpPrefixStateRequest) Ipv6UnicastFilters() BgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter {
	if len(obj.obj.Ipv6UnicastFilters) == 0 {
		obj.obj.Ipv6UnicastFilters = []*otg.BgpPrefixIpv6UnicastFilter{}
	}
	if obj.ipv6UnicastFiltersHolder == nil {
		obj.ipv6UnicastFiltersHolder = newBgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter(&obj.obj.Ipv6UnicastFilters).setMsg(obj)
	}
	return obj.ipv6UnicastFiltersHolder
}

type bgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter struct {
	obj                             *bgpPrefixStateRequest
	bgpPrefixIpv6UnicastFilterSlice []BgpPrefixIpv6UnicastFilter
	fieldPtr                        *[]*otg.BgpPrefixIpv6UnicastFilter
}

func newBgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter(ptr *[]*otg.BgpPrefixIpv6UnicastFilter) BgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter {
	return &bgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter{fieldPtr: ptr}
}

type BgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter interface {
	setMsg(*bgpPrefixStateRequest) BgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter
	Items() []BgpPrefixIpv6UnicastFilter
	Add() BgpPrefixIpv6UnicastFilter
	Append(items ...BgpPrefixIpv6UnicastFilter) BgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter
	Set(index int, newObj BgpPrefixIpv6UnicastFilter) BgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter
	Clear() BgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter
	clearHolderSlice() BgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter
	appendHolderSlice(item BgpPrefixIpv6UnicastFilter) BgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter
}

func (obj *bgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter) setMsg(msg *bgpPrefixStateRequest) BgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpPrefixIpv6UnicastFilter{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter) Items() []BgpPrefixIpv6UnicastFilter {
	return obj.bgpPrefixIpv6UnicastFilterSlice
}

func (obj *bgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter) Add() BgpPrefixIpv6UnicastFilter {
	newObj := &otg.BgpPrefixIpv6UnicastFilter{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpPrefixIpv6UnicastFilter{obj: newObj}
	newLibObj.setDefault()
	obj.bgpPrefixIpv6UnicastFilterSlice = append(obj.bgpPrefixIpv6UnicastFilterSlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter) Append(items ...BgpPrefixIpv6UnicastFilter) BgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpPrefixIpv6UnicastFilterSlice = append(obj.bgpPrefixIpv6UnicastFilterSlice, item)
	}
	return obj
}

func (obj *bgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter) Set(index int, newObj BgpPrefixIpv6UnicastFilter) BgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpPrefixIpv6UnicastFilterSlice[index] = newObj
	return obj
}
func (obj *bgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter) Clear() BgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpPrefixIpv6UnicastFilter{}
		obj.bgpPrefixIpv6UnicastFilterSlice = []BgpPrefixIpv6UnicastFilter{}
	}
	return obj
}
func (obj *bgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter) clearHolderSlice() BgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter {
	if len(obj.bgpPrefixIpv6UnicastFilterSlice) > 0 {
		obj.bgpPrefixIpv6UnicastFilterSlice = []BgpPrefixIpv6UnicastFilter{}
	}
	return obj
}
func (obj *bgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter) appendHolderSlice(item BgpPrefixIpv6UnicastFilter) BgpPrefixStateRequestBgpPrefixIpv6UnicastFilterIter {
	obj.bgpPrefixIpv6UnicastFilterSlice = append(obj.bgpPrefixIpv6UnicastFilterSlice, item)
	return obj
}

func (obj *bgpPrefixStateRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Ipv4UnicastFilters) != 0 {

		if set_default {
			obj.Ipv4UnicastFilters().clearHolderSlice()
			for _, item := range obj.obj.Ipv4UnicastFilters {
				obj.Ipv4UnicastFilters().appendHolderSlice(&bgpPrefixIpv4UnicastFilter{obj: item})
			}
		}
		for _, item := range obj.Ipv4UnicastFilters().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Ipv6UnicastFilters) != 0 {

		if set_default {
			obj.Ipv6UnicastFilters().clearHolderSlice()
			for _, item := range obj.obj.Ipv6UnicastFilters {
				obj.Ipv6UnicastFilters().appendHolderSlice(&bgpPrefixIpv6UnicastFilter{obj: item})
			}
		}
		for _, item := range obj.Ipv6UnicastFilters().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bgpPrefixStateRequest) setDefault() {

}

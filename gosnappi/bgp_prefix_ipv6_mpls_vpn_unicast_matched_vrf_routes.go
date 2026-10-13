package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes *****
type bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes struct {
	validation
	obj          *otg.BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes
	marshaller   marshalBgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes
	unMarshaller unMarshalBgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes
	vrfsHolder   BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter
}

func NewBgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes() BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes {
	obj := bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes{obj: &otg.BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) msg() *otg.BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes {
	return obj.obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) setMsg(msg *otg.BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes struct {
	obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes
}

type marshalBgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes interface {
	// ToProto marshals BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes to protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes
	ToProto() (*otg.BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes, error)
	// ToPbText marshals BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes struct {
	obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes
}

type unMarshalBgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes interface {
	// FromProto unmarshals BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes from protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes
	FromProto(msg *otg.BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) (BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes, error)
	// FromPbText unmarshals BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) Marshal() marshalBgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) Unmarshal() unMarshalBgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) ToProto() (*otg.BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) FromProto(msg *otg.BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) (BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) FromPbText(value string) error {
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

func (m *marshalbgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) FromYaml(value string) error {
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

func (m *marshalbgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) FromJson(value string) error {
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

func (obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) Clone() (BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes()
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

func (obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) setNil() {
	obj.vrfsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes is the local VRFs that this received VPN-IPv6 NLRI was imported into via Route Target matching (RFC 4364 Section 4.3.1), each with its own copy of the received path attributes.
type BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes interface {
	Validation
	// msg marshals BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes to protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes
	// and doesn't set defaults
	msg() *otg.BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes
	// setMsg unmarshals BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes from protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes
	// provides marshal interface
	Marshal() marshalBgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes
	// validate validates BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Vrfs returns BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIterIter, set in BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes
	Vrfs() BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter
	setNil()
}

// One entry per local VRF (Bgp.L3vpn.Vrf) that imported this prefix. The same VPN-IPv6 NLRI can be imported into more than one VRF when multiple VRFs' route_target_import values match one of the route's Route Target Extended Communities.
// Vrfs returns a []BgpPrefixIpv6MplsVpnUnicastVrfRoute
func (obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) Vrfs() BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter {
	if len(obj.obj.Vrfs) == 0 {
		obj.obj.Vrfs = []*otg.BgpPrefixIpv6MplsVpnUnicastVrfRoute{}
	}
	if obj.vrfsHolder == nil {
		obj.vrfsHolder = newBgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter(&obj.obj.Vrfs).setMsg(obj)
	}
	return obj.vrfsHolder
}

type bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter struct {
	obj                                      *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes
	bgpPrefixIpv6MplsVpnUnicastVrfRouteSlice []BgpPrefixIpv6MplsVpnUnicastVrfRoute
	fieldPtr                                 *[]*otg.BgpPrefixIpv6MplsVpnUnicastVrfRoute
}

func newBgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter(ptr *[]*otg.BgpPrefixIpv6MplsVpnUnicastVrfRoute) BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter {
	return &bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter{fieldPtr: ptr}
}

type BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter interface {
	setMsg(*bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter
	Items() []BgpPrefixIpv6MplsVpnUnicastVrfRoute
	Add() BgpPrefixIpv6MplsVpnUnicastVrfRoute
	Append(items ...BgpPrefixIpv6MplsVpnUnicastVrfRoute) BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter
	Set(index int, newObj BgpPrefixIpv6MplsVpnUnicastVrfRoute) BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter
	Clear() BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter
	clearHolderSlice() BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter
	appendHolderSlice(item BgpPrefixIpv6MplsVpnUnicastVrfRoute) BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter
}

func (obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter) setMsg(msg *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpPrefixIpv6MplsVpnUnicastVrfRoute{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter) Items() []BgpPrefixIpv6MplsVpnUnicastVrfRoute {
	return obj.bgpPrefixIpv6MplsVpnUnicastVrfRouteSlice
}

func (obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter) Add() BgpPrefixIpv6MplsVpnUnicastVrfRoute {
	newObj := &otg.BgpPrefixIpv6MplsVpnUnicastVrfRoute{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpPrefixIpv6MplsVpnUnicastVrfRoute{obj: newObj}
	newLibObj.setDefault()
	obj.bgpPrefixIpv6MplsVpnUnicastVrfRouteSlice = append(obj.bgpPrefixIpv6MplsVpnUnicastVrfRouteSlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter) Append(items ...BgpPrefixIpv6MplsVpnUnicastVrfRoute) BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpPrefixIpv6MplsVpnUnicastVrfRouteSlice = append(obj.bgpPrefixIpv6MplsVpnUnicastVrfRouteSlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter) Set(index int, newObj BgpPrefixIpv6MplsVpnUnicastVrfRoute) BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpPrefixIpv6MplsVpnUnicastVrfRouteSlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter) Clear() BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpPrefixIpv6MplsVpnUnicastVrfRoute{}
		obj.bgpPrefixIpv6MplsVpnUnicastVrfRouteSlice = []BgpPrefixIpv6MplsVpnUnicastVrfRoute{}
	}
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter) clearHolderSlice() BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter {
	if len(obj.bgpPrefixIpv6MplsVpnUnicastVrfRouteSlice) > 0 {
		obj.bgpPrefixIpv6MplsVpnUnicastVrfRouteSlice = []BgpPrefixIpv6MplsVpnUnicastVrfRoute{}
	}
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter) appendHolderSlice(item BgpPrefixIpv6MplsVpnUnicastVrfRoute) BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv6MplsVpnUnicastVrfRouteIter {
	obj.bgpPrefixIpv6MplsVpnUnicastVrfRouteSlice = append(obj.bgpPrefixIpv6MplsVpnUnicastVrfRouteSlice, item)
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Vrfs) != 0 {

		if set_default {
			obj.Vrfs().clearHolderSlice()
			for _, item := range obj.obj.Vrfs {
				obj.Vrfs().appendHolderSlice(&bgpPrefixIpv6MplsVpnUnicastVrfRoute{obj: item})
			}
		}
		for _, item := range obj.Vrfs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) setDefault() {

}

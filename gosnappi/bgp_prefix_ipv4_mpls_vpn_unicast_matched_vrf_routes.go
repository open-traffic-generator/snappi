package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes *****
type bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes struct {
	validation
	obj          *otg.BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes
	marshaller   marshalBgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes
	unMarshaller unMarshalBgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes
	vrfsHolder   BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter
}

func NewBgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes() BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes {
	obj := bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes{obj: &otg.BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) msg() *otg.BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes {
	return obj.obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) setMsg(msg *otg.BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes struct {
	obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes
}

type marshalBgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes interface {
	// ToProto marshals BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes to protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes
	ToProto() (*otg.BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes, error)
	// ToPbText marshals BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes struct {
	obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes
}

type unMarshalBgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes interface {
	// FromProto unmarshals BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes from protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes
	FromProto(msg *otg.BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) (BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes, error)
	// FromPbText unmarshals BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) Marshal() marshalBgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) Unmarshal() unMarshalBgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) ToProto() (*otg.BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) FromProto(msg *otg.BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) (BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) FromPbText(value string) error {
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

func (m *marshalbgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) FromYaml(value string) error {
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

func (m *marshalbgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) FromJson(value string) error {
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

func (obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) Clone() (BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes()
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

func (obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) setNil() {
	obj.vrfsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes is the local VRFs that received VPN-IPv4 NLRIs were imported into via Route Target matching (RFC 4364 Section 4.3.1), each with its own list of matched routes.
type BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes interface {
	Validation
	// msg marshals BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes to protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes
	// and doesn't set defaults
	msg() *otg.BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes
	// setMsg unmarshals BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes from protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes
	// provides marshal interface
	Marshal() marshalBgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes
	// validate validates BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Vrfs returns BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIterIter, set in BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes
	Vrfs() BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter
	setNil()
}

// One entry per local VRF (Bgp.L3vpn.Vrf) that imported one or more VPN-IPv4 routes. The same VPN-IPv4 NLRI can be imported into more than one VRF when multiple VRFs' route_target_import values match one of the route's Route Target Extended Communities, in which case it appears under each matching VRF's routes.
// Vrfs returns a []BgpPrefixIpv4MplsVpnUnicastVrf
func (obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) Vrfs() BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter {
	if len(obj.obj.Vrfs) == 0 {
		obj.obj.Vrfs = []*otg.BgpPrefixIpv4MplsVpnUnicastVrf{}
	}
	if obj.vrfsHolder == nil {
		obj.vrfsHolder = newBgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter(&obj.obj.Vrfs).setMsg(obj)
	}
	return obj.vrfsHolder
}

type bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter struct {
	obj                                 *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes
	bgpPrefixIpv4MplsVpnUnicastVrfSlice []BgpPrefixIpv4MplsVpnUnicastVrf
	fieldPtr                            *[]*otg.BgpPrefixIpv4MplsVpnUnicastVrf
}

func newBgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter(ptr *[]*otg.BgpPrefixIpv4MplsVpnUnicastVrf) BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter {
	return &bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter{fieldPtr: ptr}
}

type BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter interface {
	setMsg(*bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter
	Items() []BgpPrefixIpv4MplsVpnUnicastVrf
	Add() BgpPrefixIpv4MplsVpnUnicastVrf
	Append(items ...BgpPrefixIpv4MplsVpnUnicastVrf) BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter
	Set(index int, newObj BgpPrefixIpv4MplsVpnUnicastVrf) BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter
	Clear() BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter
	clearHolderSlice() BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter
	appendHolderSlice(item BgpPrefixIpv4MplsVpnUnicastVrf) BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter
}

func (obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter) setMsg(msg *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpPrefixIpv4MplsVpnUnicastVrf{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter) Items() []BgpPrefixIpv4MplsVpnUnicastVrf {
	return obj.bgpPrefixIpv4MplsVpnUnicastVrfSlice
}

func (obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter) Add() BgpPrefixIpv4MplsVpnUnicastVrf {
	newObj := &otg.BgpPrefixIpv4MplsVpnUnicastVrf{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpPrefixIpv4MplsVpnUnicastVrf{obj: newObj}
	newLibObj.setDefault()
	obj.bgpPrefixIpv4MplsVpnUnicastVrfSlice = append(obj.bgpPrefixIpv4MplsVpnUnicastVrfSlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter) Append(items ...BgpPrefixIpv4MplsVpnUnicastVrf) BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpPrefixIpv4MplsVpnUnicastVrfSlice = append(obj.bgpPrefixIpv4MplsVpnUnicastVrfSlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter) Set(index int, newObj BgpPrefixIpv4MplsVpnUnicastVrf) BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpPrefixIpv4MplsVpnUnicastVrfSlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter) Clear() BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpPrefixIpv4MplsVpnUnicastVrf{}
		obj.bgpPrefixIpv4MplsVpnUnicastVrfSlice = []BgpPrefixIpv4MplsVpnUnicastVrf{}
	}
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter) clearHolderSlice() BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter {
	if len(obj.bgpPrefixIpv4MplsVpnUnicastVrfSlice) > 0 {
		obj.bgpPrefixIpv4MplsVpnUnicastVrfSlice = []BgpPrefixIpv4MplsVpnUnicastVrf{}
	}
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter) appendHolderSlice(item BgpPrefixIpv4MplsVpnUnicastVrf) BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutesBgpPrefixIpv4MplsVpnUnicastVrfIter {
	obj.bgpPrefixIpv4MplsVpnUnicastVrfSlice = append(obj.bgpPrefixIpv4MplsVpnUnicastVrfSlice, item)
	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Vrfs) != 0 {

		if set_default {
			obj.Vrfs().clearHolderSlice()
			for _, item := range obj.obj.Vrfs {
				obj.Vrfs().appendHolderSlice(&bgpPrefixIpv4MplsVpnUnicastVrf{obj: item})
			}
		}
		for _, item := range obj.Vrfs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) setDefault() {

}

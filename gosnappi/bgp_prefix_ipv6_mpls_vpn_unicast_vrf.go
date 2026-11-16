package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixIpv6MplsVpnUnicastVrf *****
type bgpPrefixIpv6MplsVpnUnicastVrf struct {
	validation
	obj          *otg.BgpPrefixIpv6MplsVpnUnicastVrf
	marshaller   marshalBgpPrefixIpv6MplsVpnUnicastVrf
	unMarshaller unMarshalBgpPrefixIpv6MplsVpnUnicastVrf
	routesHolder BgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter
}

func NewBgpPrefixIpv6MplsVpnUnicastVrf() BgpPrefixIpv6MplsVpnUnicastVrf {
	obj := bgpPrefixIpv6MplsVpnUnicastVrf{obj: &otg.BgpPrefixIpv6MplsVpnUnicastVrf{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrf) msg() *otg.BgpPrefixIpv6MplsVpnUnicastVrf {
	return obj.obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrf) setMsg(msg *otg.BgpPrefixIpv6MplsVpnUnicastVrf) BgpPrefixIpv6MplsVpnUnicastVrf {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixIpv6MplsVpnUnicastVrf struct {
	obj *bgpPrefixIpv6MplsVpnUnicastVrf
}

type marshalBgpPrefixIpv6MplsVpnUnicastVrf interface {
	// ToProto marshals BgpPrefixIpv6MplsVpnUnicastVrf to protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastVrf
	ToProto() (*otg.BgpPrefixIpv6MplsVpnUnicastVrf, error)
	// ToPbText marshals BgpPrefixIpv6MplsVpnUnicastVrf to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixIpv6MplsVpnUnicastVrf to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixIpv6MplsVpnUnicastVrf to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixIpv6MplsVpnUnicastVrf struct {
	obj *bgpPrefixIpv6MplsVpnUnicastVrf
}

type unMarshalBgpPrefixIpv6MplsVpnUnicastVrf interface {
	// FromProto unmarshals BgpPrefixIpv6MplsVpnUnicastVrf from protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastVrf
	FromProto(msg *otg.BgpPrefixIpv6MplsVpnUnicastVrf) (BgpPrefixIpv6MplsVpnUnicastVrf, error)
	// FromPbText unmarshals BgpPrefixIpv6MplsVpnUnicastVrf from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixIpv6MplsVpnUnicastVrf from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixIpv6MplsVpnUnicastVrf from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrf) Marshal() marshalBgpPrefixIpv6MplsVpnUnicastVrf {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixIpv6MplsVpnUnicastVrf{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrf) Unmarshal() unMarshalBgpPrefixIpv6MplsVpnUnicastVrf {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixIpv6MplsVpnUnicastVrf{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixIpv6MplsVpnUnicastVrf) ToProto() (*otg.BgpPrefixIpv6MplsVpnUnicastVrf, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastVrf) FromProto(msg *otg.BgpPrefixIpv6MplsVpnUnicastVrf) (BgpPrefixIpv6MplsVpnUnicastVrf, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixIpv6MplsVpnUnicastVrf) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastVrf) FromPbText(value string) error {
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

func (m *marshalbgpPrefixIpv6MplsVpnUnicastVrf) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastVrf) FromYaml(value string) error {
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

func (m *marshalbgpPrefixIpv6MplsVpnUnicastVrf) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastVrf) FromJson(value string) error {
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

func (obj *bgpPrefixIpv6MplsVpnUnicastVrf) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrf) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrf) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrf) Clone() (BgpPrefixIpv6MplsVpnUnicastVrf, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixIpv6MplsVpnUnicastVrf()
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

func (obj *bgpPrefixIpv6MplsVpnUnicastVrf) setNil() {
	obj.routesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpPrefixIpv6MplsVpnUnicastVrf is a single local VRF (Bgp.L3vpn.Vrf) and the VPN-IPv6 routes that were imported into it.
type BgpPrefixIpv6MplsVpnUnicastVrf interface {
	Validation
	// msg marshals BgpPrefixIpv6MplsVpnUnicastVrf to protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastVrf
	// and doesn't set defaults
	msg() *otg.BgpPrefixIpv6MplsVpnUnicastVrf
	// setMsg unmarshals BgpPrefixIpv6MplsVpnUnicastVrf from protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastVrf
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixIpv6MplsVpnUnicastVrf) BgpPrefixIpv6MplsVpnUnicastVrf
	// provides marshal interface
	Marshal() marshalBgpPrefixIpv6MplsVpnUnicastVrf
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixIpv6MplsVpnUnicastVrf
	// validate validates BgpPrefixIpv6MplsVpnUnicastVrf
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixIpv6MplsVpnUnicastVrf, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in BgpPrefixIpv6MplsVpnUnicastVrf.
	Name() string
	// SetName assigns string provided by user to BgpPrefixIpv6MplsVpnUnicastVrf
	SetName(value string) BgpPrefixIpv6MplsVpnUnicastVrf
	// HasName checks if Name has been set in BgpPrefixIpv6MplsVpnUnicastVrf
	HasName() bool
	// Routes returns BgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIterIter, set in BgpPrefixIpv6MplsVpnUnicastVrf
	Routes() BgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter
	setNil()
}

// The name of the local VRF (Bgp.L3vpn.Vrf) that imported these routes.
// Name returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastVrf) Name() string {

	return *obj.obj.Name

}

// The name of the local VRF (Bgp.L3vpn.Vrf) that imported these routes.
// Name returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastVrf) HasName() bool {
	return obj.obj.Name != nil
}

// The name of the local VRF (Bgp.L3vpn.Vrf) that imported these routes.
// SetName sets the string value in the BgpPrefixIpv6MplsVpnUnicastVrf object
func (obj *bgpPrefixIpv6MplsVpnUnicastVrf) SetName(value string) BgpPrefixIpv6MplsVpnUnicastVrf {

	obj.obj.Name = &value
	return obj
}

// The VPN-IPv6 routes imported into this VRF.
// Routes returns a []BgpPrefixIpv6MplsVpnUnicastRoute
func (obj *bgpPrefixIpv6MplsVpnUnicastVrf) Routes() BgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter {
	if len(obj.obj.Routes) == 0 {
		obj.obj.Routes = []*otg.BgpPrefixIpv6MplsVpnUnicastRoute{}
	}
	if obj.routesHolder == nil {
		obj.routesHolder = newBgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter(&obj.obj.Routes).setMsg(obj)
	}
	return obj.routesHolder
}

type bgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter struct {
	obj                                   *bgpPrefixIpv6MplsVpnUnicastVrf
	bgpPrefixIpv6MplsVpnUnicastRouteSlice []BgpPrefixIpv6MplsVpnUnicastRoute
	fieldPtr                              *[]*otg.BgpPrefixIpv6MplsVpnUnicastRoute
}

func newBgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter(ptr *[]*otg.BgpPrefixIpv6MplsVpnUnicastRoute) BgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter {
	return &bgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter{fieldPtr: ptr}
}

type BgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter interface {
	setMsg(*bgpPrefixIpv6MplsVpnUnicastVrf) BgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter
	Items() []BgpPrefixIpv6MplsVpnUnicastRoute
	Add() BgpPrefixIpv6MplsVpnUnicastRoute
	Append(items ...BgpPrefixIpv6MplsVpnUnicastRoute) BgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter
	Set(index int, newObj BgpPrefixIpv6MplsVpnUnicastRoute) BgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter
	Clear() BgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter
	clearHolderSlice() BgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter
	appendHolderSlice(item BgpPrefixIpv6MplsVpnUnicastRoute) BgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter) setMsg(msg *bgpPrefixIpv6MplsVpnUnicastVrf) BgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpPrefixIpv6MplsVpnUnicastRoute{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter) Items() []BgpPrefixIpv6MplsVpnUnicastRoute {
	return obj.bgpPrefixIpv6MplsVpnUnicastRouteSlice
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter) Add() BgpPrefixIpv6MplsVpnUnicastRoute {
	newObj := &otg.BgpPrefixIpv6MplsVpnUnicastRoute{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpPrefixIpv6MplsVpnUnicastRoute{obj: newObj}
	newLibObj.setDefault()
	obj.bgpPrefixIpv6MplsVpnUnicastRouteSlice = append(obj.bgpPrefixIpv6MplsVpnUnicastRouteSlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter) Append(items ...BgpPrefixIpv6MplsVpnUnicastRoute) BgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpPrefixIpv6MplsVpnUnicastRouteSlice = append(obj.bgpPrefixIpv6MplsVpnUnicastRouteSlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter) Set(index int, newObj BgpPrefixIpv6MplsVpnUnicastRoute) BgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpPrefixIpv6MplsVpnUnicastRouteSlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter) Clear() BgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpPrefixIpv6MplsVpnUnicastRoute{}
		obj.bgpPrefixIpv6MplsVpnUnicastRouteSlice = []BgpPrefixIpv6MplsVpnUnicastRoute{}
	}
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter) clearHolderSlice() BgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter {
	if len(obj.bgpPrefixIpv6MplsVpnUnicastRouteSlice) > 0 {
		obj.bgpPrefixIpv6MplsVpnUnicastRouteSlice = []BgpPrefixIpv6MplsVpnUnicastRoute{}
	}
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter) appendHolderSlice(item BgpPrefixIpv6MplsVpnUnicastRoute) BgpPrefixIpv6MplsVpnUnicastVrfBgpPrefixIpv6MplsVpnUnicastRouteIter {
	obj.bgpPrefixIpv6MplsVpnUnicastRouteSlice = append(obj.bgpPrefixIpv6MplsVpnUnicastRouteSlice, item)
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrf) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Routes) != 0 {

		if set_default {
			obj.Routes().clearHolderSlice()
			for _, item := range obj.obj.Routes {
				obj.Routes().appendHolderSlice(&bgpPrefixIpv6MplsVpnUnicastRoute{obj: item})
			}
		}
		for _, item := range obj.Routes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrf) setDefault() {

}

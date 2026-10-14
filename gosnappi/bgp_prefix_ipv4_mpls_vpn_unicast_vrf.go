package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixIpv4MplsVpnUnicastVrf *****
type bgpPrefixIpv4MplsVpnUnicastVrf struct {
	validation
	obj          *otg.BgpPrefixIpv4MplsVpnUnicastVrf
	marshaller   marshalBgpPrefixIpv4MplsVpnUnicastVrf
	unMarshaller unMarshalBgpPrefixIpv4MplsVpnUnicastVrf
	routesHolder BgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter
}

func NewBgpPrefixIpv4MplsVpnUnicastVrf() BgpPrefixIpv4MplsVpnUnicastVrf {
	obj := bgpPrefixIpv4MplsVpnUnicastVrf{obj: &otg.BgpPrefixIpv4MplsVpnUnicastVrf{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrf) msg() *otg.BgpPrefixIpv4MplsVpnUnicastVrf {
	return obj.obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrf) setMsg(msg *otg.BgpPrefixIpv4MplsVpnUnicastVrf) BgpPrefixIpv4MplsVpnUnicastVrf {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixIpv4MplsVpnUnicastVrf struct {
	obj *bgpPrefixIpv4MplsVpnUnicastVrf
}

type marshalBgpPrefixIpv4MplsVpnUnicastVrf interface {
	// ToProto marshals BgpPrefixIpv4MplsVpnUnicastVrf to protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastVrf
	ToProto() (*otg.BgpPrefixIpv4MplsVpnUnicastVrf, error)
	// ToPbText marshals BgpPrefixIpv4MplsVpnUnicastVrf to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixIpv4MplsVpnUnicastVrf to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixIpv4MplsVpnUnicastVrf to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixIpv4MplsVpnUnicastVrf struct {
	obj *bgpPrefixIpv4MplsVpnUnicastVrf
}

type unMarshalBgpPrefixIpv4MplsVpnUnicastVrf interface {
	// FromProto unmarshals BgpPrefixIpv4MplsVpnUnicastVrf from protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastVrf
	FromProto(msg *otg.BgpPrefixIpv4MplsVpnUnicastVrf) (BgpPrefixIpv4MplsVpnUnicastVrf, error)
	// FromPbText unmarshals BgpPrefixIpv4MplsVpnUnicastVrf from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixIpv4MplsVpnUnicastVrf from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixIpv4MplsVpnUnicastVrf from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrf) Marshal() marshalBgpPrefixIpv4MplsVpnUnicastVrf {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixIpv4MplsVpnUnicastVrf{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrf) Unmarshal() unMarshalBgpPrefixIpv4MplsVpnUnicastVrf {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixIpv4MplsVpnUnicastVrf{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixIpv4MplsVpnUnicastVrf) ToProto() (*otg.BgpPrefixIpv4MplsVpnUnicastVrf, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastVrf) FromProto(msg *otg.BgpPrefixIpv4MplsVpnUnicastVrf) (BgpPrefixIpv4MplsVpnUnicastVrf, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixIpv4MplsVpnUnicastVrf) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastVrf) FromPbText(value string) error {
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

func (m *marshalbgpPrefixIpv4MplsVpnUnicastVrf) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastVrf) FromYaml(value string) error {
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

func (m *marshalbgpPrefixIpv4MplsVpnUnicastVrf) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastVrf) FromJson(value string) error {
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

func (obj *bgpPrefixIpv4MplsVpnUnicastVrf) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrf) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrf) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrf) Clone() (BgpPrefixIpv4MplsVpnUnicastVrf, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixIpv4MplsVpnUnicastVrf()
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

func (obj *bgpPrefixIpv4MplsVpnUnicastVrf) setNil() {
	obj.routesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpPrefixIpv4MplsVpnUnicastVrf is a single local VRF (Bgp.L3vpn.Vrf) and the VPN-IPv4 routes that were imported into it.
type BgpPrefixIpv4MplsVpnUnicastVrf interface {
	Validation
	// msg marshals BgpPrefixIpv4MplsVpnUnicastVrf to protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastVrf
	// and doesn't set defaults
	msg() *otg.BgpPrefixIpv4MplsVpnUnicastVrf
	// setMsg unmarshals BgpPrefixIpv4MplsVpnUnicastVrf from protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastVrf
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixIpv4MplsVpnUnicastVrf) BgpPrefixIpv4MplsVpnUnicastVrf
	// provides marshal interface
	Marshal() marshalBgpPrefixIpv4MplsVpnUnicastVrf
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixIpv4MplsVpnUnicastVrf
	// validate validates BgpPrefixIpv4MplsVpnUnicastVrf
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixIpv4MplsVpnUnicastVrf, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in BgpPrefixIpv4MplsVpnUnicastVrf.
	Name() string
	// SetName assigns string provided by user to BgpPrefixIpv4MplsVpnUnicastVrf
	SetName(value string) BgpPrefixIpv4MplsVpnUnicastVrf
	// HasName checks if Name has been set in BgpPrefixIpv4MplsVpnUnicastVrf
	HasName() bool
	// Routes returns BgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIterIter, set in BgpPrefixIpv4MplsVpnUnicastVrf
	Routes() BgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter
	setNil()
}

// The name of the local VRF (Bgp.L3vpn.Vrf) that imported these routes.
// Name returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastVrf) Name() string {

	return *obj.obj.Name

}

// The name of the local VRF (Bgp.L3vpn.Vrf) that imported these routes.
// Name returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastVrf) HasName() bool {
	return obj.obj.Name != nil
}

// The name of the local VRF (Bgp.L3vpn.Vrf) that imported these routes.
// SetName sets the string value in the BgpPrefixIpv4MplsVpnUnicastVrf object
func (obj *bgpPrefixIpv4MplsVpnUnicastVrf) SetName(value string) BgpPrefixIpv4MplsVpnUnicastVrf {

	obj.obj.Name = &value
	return obj
}

// The VPN-IPv4 routes imported into this VRF.
// Routes returns a []BgpPrefixIpv4MplsVpnUnicastRoute
func (obj *bgpPrefixIpv4MplsVpnUnicastVrf) Routes() BgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter {
	if len(obj.obj.Routes) == 0 {
		obj.obj.Routes = []*otg.BgpPrefixIpv4MplsVpnUnicastRoute{}
	}
	if obj.routesHolder == nil {
		obj.routesHolder = newBgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter(&obj.obj.Routes).setMsg(obj)
	}
	return obj.routesHolder
}

type bgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter struct {
	obj                                   *bgpPrefixIpv4MplsVpnUnicastVrf
	bgpPrefixIpv4MplsVpnUnicastRouteSlice []BgpPrefixIpv4MplsVpnUnicastRoute
	fieldPtr                              *[]*otg.BgpPrefixIpv4MplsVpnUnicastRoute
}

func newBgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter(ptr *[]*otg.BgpPrefixIpv4MplsVpnUnicastRoute) BgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter {
	return &bgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter{fieldPtr: ptr}
}

type BgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter interface {
	setMsg(*bgpPrefixIpv4MplsVpnUnicastVrf) BgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter
	Items() []BgpPrefixIpv4MplsVpnUnicastRoute
	Add() BgpPrefixIpv4MplsVpnUnicastRoute
	Append(items ...BgpPrefixIpv4MplsVpnUnicastRoute) BgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter
	Set(index int, newObj BgpPrefixIpv4MplsVpnUnicastRoute) BgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter
	Clear() BgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter
	clearHolderSlice() BgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter
	appendHolderSlice(item BgpPrefixIpv4MplsVpnUnicastRoute) BgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter) setMsg(msg *bgpPrefixIpv4MplsVpnUnicastVrf) BgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpPrefixIpv4MplsVpnUnicastRoute{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter) Items() []BgpPrefixIpv4MplsVpnUnicastRoute {
	return obj.bgpPrefixIpv4MplsVpnUnicastRouteSlice
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter) Add() BgpPrefixIpv4MplsVpnUnicastRoute {
	newObj := &otg.BgpPrefixIpv4MplsVpnUnicastRoute{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpPrefixIpv4MplsVpnUnicastRoute{obj: newObj}
	newLibObj.setDefault()
	obj.bgpPrefixIpv4MplsVpnUnicastRouteSlice = append(obj.bgpPrefixIpv4MplsVpnUnicastRouteSlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter) Append(items ...BgpPrefixIpv4MplsVpnUnicastRoute) BgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpPrefixIpv4MplsVpnUnicastRouteSlice = append(obj.bgpPrefixIpv4MplsVpnUnicastRouteSlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter) Set(index int, newObj BgpPrefixIpv4MplsVpnUnicastRoute) BgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpPrefixIpv4MplsVpnUnicastRouteSlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter) Clear() BgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpPrefixIpv4MplsVpnUnicastRoute{}
		obj.bgpPrefixIpv4MplsVpnUnicastRouteSlice = []BgpPrefixIpv4MplsVpnUnicastRoute{}
	}
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter) clearHolderSlice() BgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter {
	if len(obj.bgpPrefixIpv4MplsVpnUnicastRouteSlice) > 0 {
		obj.bgpPrefixIpv4MplsVpnUnicastRouteSlice = []BgpPrefixIpv4MplsVpnUnicastRoute{}
	}
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter) appendHolderSlice(item BgpPrefixIpv4MplsVpnUnicastRoute) BgpPrefixIpv4MplsVpnUnicastVrfBgpPrefixIpv4MplsVpnUnicastRouteIter {
	obj.bgpPrefixIpv4MplsVpnUnicastRouteSlice = append(obj.bgpPrefixIpv4MplsVpnUnicastRouteSlice, item)
	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrf) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Routes) != 0 {

		if set_default {
			obj.Routes().clearHolderSlice()
			for _, item := range obj.obj.Routes {
				obj.Routes().appendHolderSlice(&bgpPrefixIpv4MplsVpnUnicastRoute{obj: item})
			}
		}
		for _, item := range obj.Routes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrf) setDefault() {

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpL3VpnVrf *****
type bgpL3VpnVrf struct {
	validation
	obj                      *otg.BgpL3VpnVrf
	marshaller               marshalBgpL3VpnVrf
	unMarshaller             unMarshalBgpL3VpnVrf
	routeDistinguisherHolder BgpRouteDistinguisher
	routeTargetExportHolder  BgpL3VpnVrfBgpRouteTargetIter
	routeTargetImportHolder  BgpL3VpnVrfBgpRouteTargetIter
	v4RoutesHolder           BgpL3VpnVrfBgpV4RouteRangeIter
}

func NewBgpL3VpnVrf() BgpL3VpnVrf {
	obj := bgpL3VpnVrf{obj: &otg.BgpL3VpnVrf{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpL3VpnVrf) msg() *otg.BgpL3VpnVrf {
	return obj.obj
}

func (obj *bgpL3VpnVrf) setMsg(msg *otg.BgpL3VpnVrf) BgpL3VpnVrf {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpL3VpnVrf struct {
	obj *bgpL3VpnVrf
}

type marshalBgpL3VpnVrf interface {
	// ToProto marshals BgpL3VpnVrf to protobuf object *otg.BgpL3VpnVrf
	ToProto() (*otg.BgpL3VpnVrf, error)
	// ToPbText marshals BgpL3VpnVrf to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpL3VpnVrf to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpL3VpnVrf to JSON text
	ToJson() (string, error)
}

type unMarshalbgpL3VpnVrf struct {
	obj *bgpL3VpnVrf
}

type unMarshalBgpL3VpnVrf interface {
	// FromProto unmarshals BgpL3VpnVrf from protobuf object *otg.BgpL3VpnVrf
	FromProto(msg *otg.BgpL3VpnVrf) (BgpL3VpnVrf, error)
	// FromPbText unmarshals BgpL3VpnVrf from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpL3VpnVrf from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpL3VpnVrf from JSON text
	FromJson(value string) error
}

func (obj *bgpL3VpnVrf) Marshal() marshalBgpL3VpnVrf {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpL3VpnVrf{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpL3VpnVrf) Unmarshal() unMarshalBgpL3VpnVrf {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpL3VpnVrf{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpL3VpnVrf) ToProto() (*otg.BgpL3VpnVrf, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpL3VpnVrf) FromProto(msg *otg.BgpL3VpnVrf) (BgpL3VpnVrf, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpL3VpnVrf) ToPbText() (string, error) {
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

func (m *unMarshalbgpL3VpnVrf) FromPbText(value string) error {
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

func (m *marshalbgpL3VpnVrf) ToYaml() (string, error) {
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

func (m *unMarshalbgpL3VpnVrf) FromYaml(value string) error {
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

func (m *marshalbgpL3VpnVrf) ToJson() (string, error) {
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

func (m *unMarshalbgpL3VpnVrf) FromJson(value string) error {
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

func (obj *bgpL3VpnVrf) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpL3VpnVrf) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpL3VpnVrf) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpL3VpnVrf) Clone() (BgpL3VpnVrf, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpL3VpnVrf()
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

func (obj *bgpL3VpnVrf) setNil() {
	obj.routeDistinguisherHolder = nil
	obj.routeTargetExportHolder = nil
	obj.routeTargetImportHolder = nil
	obj.v4RoutesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpL3VpnVrf is a BGP/MPLS Layer 3 VPN VRF (RFC 4364). Binds a Route Distinguisher and
// Route Target import/export policy to a set of customer (PE-CE learned
// or locally originated) IPv4 route ranges, so that they are advertised
// as VPN-IPv4 NLRI (AFI 1, SAFI 128) instead of plain IPv4 unicast NLRI.
//
// The BGP capability device.bgp.capability.ipv4_mpls_vpn must be enabled
// on the peer for VPN-IPv4 NLRI to be negotiated and advertised.
type BgpL3VpnVrf interface {
	Validation
	// msg marshals BgpL3VpnVrf to protobuf object *otg.BgpL3VpnVrf
	// and doesn't set defaults
	msg() *otg.BgpL3VpnVrf
	// setMsg unmarshals BgpL3VpnVrf from protobuf object *otg.BgpL3VpnVrf
	// and doesn't set defaults
	setMsg(*otg.BgpL3VpnVrf) BgpL3VpnVrf
	// provides marshal interface
	Marshal() marshalBgpL3VpnVrf
	// provides unmarshal interface
	Unmarshal() unMarshalBgpL3VpnVrf
	// validate validates BgpL3VpnVrf
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpL3VpnVrf, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in BgpL3VpnVrf.
	Name() string
	// SetName assigns string provided by user to BgpL3VpnVrf
	SetName(value string) BgpL3VpnVrf
	// RouteDistinguisher returns BgpRouteDistinguisher, set in BgpL3VpnVrf.
	// BgpRouteDistinguisher is bGP Route Distinguisher.
	RouteDistinguisher() BgpRouteDistinguisher
	// SetRouteDistinguisher assigns BgpRouteDistinguisher provided by user to BgpL3VpnVrf.
	// BgpRouteDistinguisher is bGP Route Distinguisher.
	SetRouteDistinguisher(value BgpRouteDistinguisher) BgpL3VpnVrf
	// RouteTargetExport returns BgpL3VpnVrfBgpRouteTargetIterIter, set in BgpL3VpnVrf
	RouteTargetExport() BgpL3VpnVrfBgpRouteTargetIter
	// RouteTargetImport returns BgpL3VpnVrfBgpRouteTargetIterIter, set in BgpL3VpnVrf
	RouteTargetImport() BgpL3VpnVrfBgpRouteTargetIter
	// V4Routes returns BgpL3VpnVrfBgpV4RouteRangeIterIter, set in BgpL3VpnVrf
	V4Routes() BgpL3VpnVrfBgpV4RouteRangeIter
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *bgpL3VpnVrf) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the BgpL3VpnVrf object
func (obj *bgpL3VpnVrf) SetName(value string) BgpL3VpnVrf {

	obj.obj.Name = &value
	return obj
}

// The Route Distinguisher (RFC 4364 Section 4.1) prepended to every IPv4 prefix advertised from this VRF's v4_routes, forming the VPN-IPv4 NLRI. All routes in this VRF share the same RD.
// RouteDistinguisher returns a BgpRouteDistinguisher
func (obj *bgpL3VpnVrf) RouteDistinguisher() BgpRouteDistinguisher {
	if obj.obj.RouteDistinguisher == nil {
		obj.obj.RouteDistinguisher = NewBgpRouteDistinguisher().msg()
	}
	if obj.routeDistinguisherHolder == nil {
		obj.routeDistinguisherHolder = &bgpRouteDistinguisher{obj: obj.obj.RouteDistinguisher}
	}
	return obj.routeDistinguisherHolder
}

// The Route Distinguisher (RFC 4364 Section 4.1) prepended to every IPv4 prefix advertised from this VRF's v4_routes, forming the VPN-IPv4 NLRI. All routes in this VRF share the same RD.
// SetRouteDistinguisher sets the BgpRouteDistinguisher value in the BgpL3VpnVrf object
func (obj *bgpL3VpnVrf) SetRouteDistinguisher(value BgpRouteDistinguisher) BgpL3VpnVrf {

	obj.routeDistinguisherHolder = nil
	obj.obj.RouteDistinguisher = value.msg()

	return obj
}

// List of Route Targets (RFC 4364 Section 4.3.1) attached as BGP Extended Communities to every route advertised from this VRF. A remote VRF imports the route only if one of its route_target_import values matches one of these.
// RouteTargetExport returns a []BgpRouteTarget
func (obj *bgpL3VpnVrf) RouteTargetExport() BgpL3VpnVrfBgpRouteTargetIter {
	if len(obj.obj.RouteTargetExport) == 0 {
		obj.obj.RouteTargetExport = []*otg.BgpRouteTarget{}
	}
	if obj.routeTargetExportHolder == nil {
		obj.routeTargetExportHolder = newBgpL3VpnVrfBgpRouteTargetIter(&obj.obj.RouteTargetExport).setMsg(obj)
	}
	return obj.routeTargetExportHolder
}

type bgpL3VpnVrfBgpRouteTargetIter struct {
	obj                 *bgpL3VpnVrf
	bgpRouteTargetSlice []BgpRouteTarget
	fieldPtr            *[]*otg.BgpRouteTarget
}

func newBgpL3VpnVrfBgpRouteTargetIter(ptr *[]*otg.BgpRouteTarget) BgpL3VpnVrfBgpRouteTargetIter {
	return &bgpL3VpnVrfBgpRouteTargetIter{fieldPtr: ptr}
}

type BgpL3VpnVrfBgpRouteTargetIter interface {
	setMsg(*bgpL3VpnVrf) BgpL3VpnVrfBgpRouteTargetIter
	Items() []BgpRouteTarget
	Add() BgpRouteTarget
	Append(items ...BgpRouteTarget) BgpL3VpnVrfBgpRouteTargetIter
	Set(index int, newObj BgpRouteTarget) BgpL3VpnVrfBgpRouteTargetIter
	Clear() BgpL3VpnVrfBgpRouteTargetIter
	clearHolderSlice() BgpL3VpnVrfBgpRouteTargetIter
	appendHolderSlice(item BgpRouteTarget) BgpL3VpnVrfBgpRouteTargetIter
}

func (obj *bgpL3VpnVrfBgpRouteTargetIter) setMsg(msg *bgpL3VpnVrf) BgpL3VpnVrfBgpRouteTargetIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpRouteTarget{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpL3VpnVrfBgpRouteTargetIter) Items() []BgpRouteTarget {
	return obj.bgpRouteTargetSlice
}

func (obj *bgpL3VpnVrfBgpRouteTargetIter) Add() BgpRouteTarget {
	newObj := &otg.BgpRouteTarget{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpRouteTarget{obj: newObj}
	newLibObj.setDefault()
	obj.bgpRouteTargetSlice = append(obj.bgpRouteTargetSlice, newLibObj)
	return newLibObj
}

func (obj *bgpL3VpnVrfBgpRouteTargetIter) Append(items ...BgpRouteTarget) BgpL3VpnVrfBgpRouteTargetIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpRouteTargetSlice = append(obj.bgpRouteTargetSlice, item)
	}
	return obj
}

func (obj *bgpL3VpnVrfBgpRouteTargetIter) Set(index int, newObj BgpRouteTarget) BgpL3VpnVrfBgpRouteTargetIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpRouteTargetSlice[index] = newObj
	return obj
}
func (obj *bgpL3VpnVrfBgpRouteTargetIter) Clear() BgpL3VpnVrfBgpRouteTargetIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpRouteTarget{}
		obj.bgpRouteTargetSlice = []BgpRouteTarget{}
	}
	return obj
}
func (obj *bgpL3VpnVrfBgpRouteTargetIter) clearHolderSlice() BgpL3VpnVrfBgpRouteTargetIter {
	if len(obj.bgpRouteTargetSlice) > 0 {
		obj.bgpRouteTargetSlice = []BgpRouteTarget{}
	}
	return obj
}
func (obj *bgpL3VpnVrfBgpRouteTargetIter) appendHolderSlice(item BgpRouteTarget) BgpL3VpnVrfBgpRouteTargetIter {
	obj.bgpRouteTargetSlice = append(obj.bgpRouteTargetSlice, item)
	return obj
}

// List of Route Targets (RFC 4364 Section 4.3.1) used to filter which received VPN-IPv4 routes are imported into this VRF. A received route is imported if any of its Route Target Extended Communities matches one of these.
// RouteTargetImport returns a []BgpRouteTarget
func (obj *bgpL3VpnVrf) RouteTargetImport() BgpL3VpnVrfBgpRouteTargetIter {
	if len(obj.obj.RouteTargetImport) == 0 {
		obj.obj.RouteTargetImport = []*otg.BgpRouteTarget{}
	}
	if obj.routeTargetImportHolder == nil {
		obj.routeTargetImportHolder = newBgpL3VpnVrfBgpRouteTargetIter(&obj.obj.RouteTargetImport).setMsg(obj)
	}
	return obj.routeTargetImportHolder
}

// Emulated IPv4 customer route ranges belonging to this VRF. Each is advertised as VPN-IPv4 NLRI using this VRF's route_distinguisher and route_target_export.
// V4Routes returns a []BgpV4RouteRange
func (obj *bgpL3VpnVrf) V4Routes() BgpL3VpnVrfBgpV4RouteRangeIter {
	if len(obj.obj.V4Routes) == 0 {
		obj.obj.V4Routes = []*otg.BgpV4RouteRange{}
	}
	if obj.v4RoutesHolder == nil {
		obj.v4RoutesHolder = newBgpL3VpnVrfBgpV4RouteRangeIter(&obj.obj.V4Routes).setMsg(obj)
	}
	return obj.v4RoutesHolder
}

type bgpL3VpnVrfBgpV4RouteRangeIter struct {
	obj                  *bgpL3VpnVrf
	bgpV4RouteRangeSlice []BgpV4RouteRange
	fieldPtr             *[]*otg.BgpV4RouteRange
}

func newBgpL3VpnVrfBgpV4RouteRangeIter(ptr *[]*otg.BgpV4RouteRange) BgpL3VpnVrfBgpV4RouteRangeIter {
	return &bgpL3VpnVrfBgpV4RouteRangeIter{fieldPtr: ptr}
}

type BgpL3VpnVrfBgpV4RouteRangeIter interface {
	setMsg(*bgpL3VpnVrf) BgpL3VpnVrfBgpV4RouteRangeIter
	Items() []BgpV4RouteRange
	Add() BgpV4RouteRange
	Append(items ...BgpV4RouteRange) BgpL3VpnVrfBgpV4RouteRangeIter
	Set(index int, newObj BgpV4RouteRange) BgpL3VpnVrfBgpV4RouteRangeIter
	Clear() BgpL3VpnVrfBgpV4RouteRangeIter
	clearHolderSlice() BgpL3VpnVrfBgpV4RouteRangeIter
	appendHolderSlice(item BgpV4RouteRange) BgpL3VpnVrfBgpV4RouteRangeIter
}

func (obj *bgpL3VpnVrfBgpV4RouteRangeIter) setMsg(msg *bgpL3VpnVrf) BgpL3VpnVrfBgpV4RouteRangeIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpV4RouteRange{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpL3VpnVrfBgpV4RouteRangeIter) Items() []BgpV4RouteRange {
	return obj.bgpV4RouteRangeSlice
}

func (obj *bgpL3VpnVrfBgpV4RouteRangeIter) Add() BgpV4RouteRange {
	newObj := &otg.BgpV4RouteRange{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpV4RouteRange{obj: newObj}
	newLibObj.setDefault()
	obj.bgpV4RouteRangeSlice = append(obj.bgpV4RouteRangeSlice, newLibObj)
	return newLibObj
}

func (obj *bgpL3VpnVrfBgpV4RouteRangeIter) Append(items ...BgpV4RouteRange) BgpL3VpnVrfBgpV4RouteRangeIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpV4RouteRangeSlice = append(obj.bgpV4RouteRangeSlice, item)
	}
	return obj
}

func (obj *bgpL3VpnVrfBgpV4RouteRangeIter) Set(index int, newObj BgpV4RouteRange) BgpL3VpnVrfBgpV4RouteRangeIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpV4RouteRangeSlice[index] = newObj
	return obj
}
func (obj *bgpL3VpnVrfBgpV4RouteRangeIter) Clear() BgpL3VpnVrfBgpV4RouteRangeIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpV4RouteRange{}
		obj.bgpV4RouteRangeSlice = []BgpV4RouteRange{}
	}
	return obj
}
func (obj *bgpL3VpnVrfBgpV4RouteRangeIter) clearHolderSlice() BgpL3VpnVrfBgpV4RouteRangeIter {
	if len(obj.bgpV4RouteRangeSlice) > 0 {
		obj.bgpV4RouteRangeSlice = []BgpV4RouteRange{}
	}
	return obj
}
func (obj *bgpL3VpnVrfBgpV4RouteRangeIter) appendHolderSlice(item BgpV4RouteRange) BgpL3VpnVrfBgpV4RouteRangeIter {
	obj.bgpV4RouteRangeSlice = append(obj.bgpV4RouteRangeSlice, item)
	return obj
}

func (obj *bgpL3VpnVrf) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface BgpL3VpnVrf")
	}

	// RouteDistinguisher is required
	if obj.obj.RouteDistinguisher == nil {
		vObj.validationErrors = append(vObj.validationErrors, "RouteDistinguisher is required field on interface BgpL3VpnVrf")
	}

	if obj.obj.RouteDistinguisher != nil {

		obj.RouteDistinguisher().validateObj(vObj, set_default)
	}

	if len(obj.obj.RouteTargetExport) != 0 {

		if set_default {
			obj.RouteTargetExport().clearHolderSlice()
			for _, item := range obj.obj.RouteTargetExport {
				obj.RouteTargetExport().appendHolderSlice(&bgpRouteTarget{obj: item})
			}
		}
		for _, item := range obj.RouteTargetExport().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.RouteTargetImport) != 0 {

		if set_default {
			obj.RouteTargetImport().clearHolderSlice()
			for _, item := range obj.obj.RouteTargetImport {
				obj.RouteTargetImport().appendHolderSlice(&bgpRouteTarget{obj: item})
			}
		}
		for _, item := range obj.RouteTargetImport().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.V4Routes) != 0 {

		if set_default {
			obj.V4Routes().clearHolderSlice()
			for _, item := range obj.obj.V4Routes {
				obj.V4Routes().appendHolderSlice(&bgpV4RouteRange{obj: item})
			}
		}
		for _, item := range obj.V4Routes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bgpL3VpnVrf) setDefault() {

}

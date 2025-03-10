package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2Interface *****
type ospfv2Interface struct {
	validation
	obj                      *otg.Ospfv2Interface
	marshaller               marshalOspfv2Interface
	unMarshaller             unMarshalOspfv2Interface
	areaHolder               Ospfv2InterfaceArea
	networkTypeHolder        Ospfv2InterfaceNetworkType
	trafficEngineeringHolder Ospfv2InterfaceLinkStateTEIter
	authenticationHolder     Ospfv2InterfaceAuthentication
	advancedHolder           Ospfv2InterfaceAdvanced
	linkProtectionHolder     Ospfv2InterfaceLinkProtection
}

func NewOspfv2Interface() Ospfv2Interface {
	obj := ospfv2Interface{obj: &otg.Ospfv2Interface{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2Interface) msg() *otg.Ospfv2Interface {
	return obj.obj
}

func (obj *ospfv2Interface) setMsg(msg *otg.Ospfv2Interface) Ospfv2Interface {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2Interface struct {
	obj *ospfv2Interface
}

type marshalOspfv2Interface interface {
	// ToProto marshals Ospfv2Interface to protobuf object *otg.Ospfv2Interface
	ToProto() (*otg.Ospfv2Interface, error)
	// ToPbText marshals Ospfv2Interface to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2Interface to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2Interface to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv2Interface to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv2Interface struct {
	obj *ospfv2Interface
}

type unMarshalOspfv2Interface interface {
	// FromProto unmarshals Ospfv2Interface from protobuf object *otg.Ospfv2Interface
	FromProto(msg *otg.Ospfv2Interface) (Ospfv2Interface, error)
	// FromPbText unmarshals Ospfv2Interface from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2Interface from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2Interface from JSON text
	FromJson(value string) error
}

func (obj *ospfv2Interface) Marshal() marshalOspfv2Interface {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2Interface{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2Interface) Unmarshal() unMarshalOspfv2Interface {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2Interface{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2Interface) ToProto() (*otg.Ospfv2Interface, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2Interface) FromProto(msg *otg.Ospfv2Interface) (Ospfv2Interface, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2Interface) ToPbText() (string, error) {
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

func (m *unMarshalospfv2Interface) FromPbText(value string) error {
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

func (m *marshalospfv2Interface) ToYaml() (string, error) {
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

func (m *unMarshalospfv2Interface) FromYaml(value string) error {
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

func (m *marshalospfv2Interface) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalospfv2Interface) ToJson() (string, error) {
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

func (m *unMarshalospfv2Interface) FromJson(value string) error {
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

func (obj *ospfv2Interface) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2Interface) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2Interface) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2Interface) Clone() (Ospfv2Interface, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2Interface()
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

func (obj *ospfv2Interface) setNil() {
	obj.areaHolder = nil
	obj.networkTypeHolder = nil
	obj.trafficEngineeringHolder = nil
	obj.authenticationHolder = nil
	obj.advancedHolder = nil
	obj.linkProtectionHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2Interface is configuration for single OSPFv2 interface.
type Ospfv2Interface interface {
	Validation
	// msg marshals Ospfv2Interface to protobuf object *otg.Ospfv2Interface
	// and doesn't set defaults
	msg() *otg.Ospfv2Interface
	// setMsg unmarshals Ospfv2Interface from protobuf object *otg.Ospfv2Interface
	// and doesn't set defaults
	setMsg(*otg.Ospfv2Interface) Ospfv2Interface
	// provides marshal interface
	Marshal() marshalOspfv2Interface
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2Interface
	// validate validates Ospfv2Interface
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2Interface, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in Ospfv2Interface.
	Name() string
	// SetName assigns string provided by user to Ospfv2Interface
	SetName(value string) Ospfv2Interface
	// Ipv4Name returns string, set in Ospfv2Interface.
	Ipv4Name() string
	// SetIpv4Name assigns string provided by user to Ospfv2Interface
	SetIpv4Name(value string) Ospfv2Interface
	// Area returns Ospfv2InterfaceArea, set in Ospfv2Interface.
	// Ospfv2InterfaceArea is container for OSPF Area ID identifies the routing area to which the host belongs..
	Area() Ospfv2InterfaceArea
	// SetArea assigns Ospfv2InterfaceArea provided by user to Ospfv2Interface.
	// Ospfv2InterfaceArea is container for OSPF Area ID identifies the routing area to which the host belongs..
	SetArea(value Ospfv2InterfaceArea) Ospfv2Interface
	// HasArea checks if Area has been set in Ospfv2Interface
	HasArea() bool
	// NetworkType returns Ospfv2InterfaceNetworkType, set in Ospfv2Interface.
	// Ospfv2InterfaceNetworkType is the OSPF network link type options.
	// - Point to Point:
	// - Broadcast:
	// - Point to Multipoint: In this case, at least a neigbor to be configured.
	NetworkType() Ospfv2InterfaceNetworkType
	// SetNetworkType assigns Ospfv2InterfaceNetworkType provided by user to Ospfv2Interface.
	// Ospfv2InterfaceNetworkType is the OSPF network link type options.
	// - Point to Point:
	// - Broadcast:
	// - Point to Multipoint: In this case, at least a neigbor to be configured.
	SetNetworkType(value Ospfv2InterfaceNetworkType) Ospfv2Interface
	// HasNetworkType checks if NetworkType has been set in Ospfv2Interface
	HasNetworkType() bool
	// TrafficEngineering returns Ospfv2InterfaceLinkStateTEIterIter, set in Ospfv2Interface
	TrafficEngineering() Ospfv2InterfaceLinkStateTEIter
	// Authentication returns Ospfv2InterfaceAuthentication, set in Ospfv2Interface.
	// Ospfv2InterfaceAuthentication is this contains OSPFv2 authentication properties.
	// Reference: https://www.rfc-editor.org/rfc/rfc2328#appendix-D
	Authentication() Ospfv2InterfaceAuthentication
	// SetAuthentication assigns Ospfv2InterfaceAuthentication provided by user to Ospfv2Interface.
	// Ospfv2InterfaceAuthentication is this contains OSPFv2 authentication properties.
	// Reference: https://www.rfc-editor.org/rfc/rfc2328#appendix-D
	SetAuthentication(value Ospfv2InterfaceAuthentication) Ospfv2Interface
	// HasAuthentication checks if Authentication has been set in Ospfv2Interface
	HasAuthentication() bool
	// Advanced returns Ospfv2InterfaceAdvanced, set in Ospfv2Interface.
	// Ospfv2InterfaceAdvanced is contains OSPFv2 advanced properties.
	Advanced() Ospfv2InterfaceAdvanced
	// SetAdvanced assigns Ospfv2InterfaceAdvanced provided by user to Ospfv2Interface.
	// Ospfv2InterfaceAdvanced is contains OSPFv2 advanced properties.
	SetAdvanced(value Ospfv2InterfaceAdvanced) Ospfv2Interface
	// HasAdvanced checks if Advanced has been set in Ospfv2Interface
	HasAdvanced() bool
	// LinkProtection returns Ospfv2InterfaceLinkProtection, set in Ospfv2Interface.
	// Ospfv2InterfaceLinkProtection is optional container for the link protection sub TLV (type 20).
	LinkProtection() Ospfv2InterfaceLinkProtection
	// SetLinkProtection assigns Ospfv2InterfaceLinkProtection provided by user to Ospfv2Interface.
	// Ospfv2InterfaceLinkProtection is optional container for the link protection sub TLV (type 20).
	SetLinkProtection(value Ospfv2InterfaceLinkProtection) Ospfv2Interface
	// HasLinkProtection checks if LinkProtection has been set in Ospfv2Interface
	HasLinkProtection() bool
	// SrlgValues returns []uint32, set in Ospfv2Interface.
	SrlgValues() []uint32
	// SetSrlgValues assigns []uint32 provided by user to Ospfv2Interface
	SetSrlgValues(value []uint32) Ospfv2Interface
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *ospfv2Interface) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the Ospfv2Interface object
func (obj *ospfv2Interface) SetName(value string) Ospfv2Interface {

	obj.obj.Name = &value
	return obj
}

// The globally unique name of the IPv4 interface connected to the DUT.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// Ipv4Name returns a string
func (obj *ospfv2Interface) Ipv4Name() string {

	return *obj.obj.Ipv4Name

}

// The globally unique name of the IPv4 interface connected to the DUT.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// SetIpv4Name sets the string value in the Ospfv2Interface object
func (obj *ospfv2Interface) SetIpv4Name(value string) Ospfv2Interface {

	obj.obj.Ipv4Name = &value
	return obj
}

// The Area ID of the area to which the attached network belongs.
// All routing protocol packets originating from the interface are
// labelled with this Area ID.
// Area returns a Ospfv2InterfaceArea
func (obj *ospfv2Interface) Area() Ospfv2InterfaceArea {
	if obj.obj.Area == nil {
		obj.obj.Area = NewOspfv2InterfaceArea().msg()
	}
	if obj.areaHolder == nil {
		obj.areaHolder = &ospfv2InterfaceArea{obj: obj.obj.Area}
	}
	return obj.areaHolder
}

// The Area ID of the area to which the attached network belongs.
// All routing protocol packets originating from the interface are
// labelled with this Area ID.
// Area returns a Ospfv2InterfaceArea
func (obj *ospfv2Interface) HasArea() bool {
	return obj.obj.Area != nil
}

// The Area ID of the area to which the attached network belongs.
// All routing protocol packets originating from the interface are
// labelled with this Area ID.
// SetArea sets the Ospfv2InterfaceArea value in the Ospfv2Interface object
func (obj *ospfv2Interface) SetArea(value Ospfv2InterfaceArea) Ospfv2Interface {

	obj.areaHolder = nil
	obj.obj.Area = value.msg()

	return obj
}

// The OSPF network link type.
// NetworkType returns a Ospfv2InterfaceNetworkType
func (obj *ospfv2Interface) NetworkType() Ospfv2InterfaceNetworkType {
	if obj.obj.NetworkType == nil {
		obj.obj.NetworkType = NewOspfv2InterfaceNetworkType().msg()
	}
	if obj.networkTypeHolder == nil {
		obj.networkTypeHolder = &ospfv2InterfaceNetworkType{obj: obj.obj.NetworkType}
	}
	return obj.networkTypeHolder
}

// The OSPF network link type.
// NetworkType returns a Ospfv2InterfaceNetworkType
func (obj *ospfv2Interface) HasNetworkType() bool {
	return obj.obj.NetworkType != nil
}

// The OSPF network link type.
// SetNetworkType sets the Ospfv2InterfaceNetworkType value in the Ospfv2Interface object
func (obj *ospfv2Interface) SetNetworkType(value Ospfv2InterfaceNetworkType) Ospfv2Interface {

	obj.networkTypeHolder = nil
	obj.obj.NetworkType = value.msg()

	return obj
}

// Contains a list of Traffic Engineering attributes.
// TrafficEngineering returns a []LinkStateTE
func (obj *ospfv2Interface) TrafficEngineering() Ospfv2InterfaceLinkStateTEIter {
	if len(obj.obj.TrafficEngineering) == 0 {
		obj.obj.TrafficEngineering = []*otg.LinkStateTE{}
	}
	if obj.trafficEngineeringHolder == nil {
		obj.trafficEngineeringHolder = newOspfv2InterfaceLinkStateTEIter(&obj.obj.TrafficEngineering).setMsg(obj)
	}
	return obj.trafficEngineeringHolder
}

type ospfv2InterfaceLinkStateTEIter struct {
	obj              *ospfv2Interface
	linkStateTESlice []LinkStateTE
	fieldPtr         *[]*otg.LinkStateTE
}

func newOspfv2InterfaceLinkStateTEIter(ptr *[]*otg.LinkStateTE) Ospfv2InterfaceLinkStateTEIter {
	return &ospfv2InterfaceLinkStateTEIter{fieldPtr: ptr}
}

type Ospfv2InterfaceLinkStateTEIter interface {
	setMsg(*ospfv2Interface) Ospfv2InterfaceLinkStateTEIter
	Items() []LinkStateTE
	Add() LinkStateTE
	Append(items ...LinkStateTE) Ospfv2InterfaceLinkStateTEIter
	Set(index int, newObj LinkStateTE) Ospfv2InterfaceLinkStateTEIter
	Clear() Ospfv2InterfaceLinkStateTEIter
	clearHolderSlice() Ospfv2InterfaceLinkStateTEIter
	appendHolderSlice(item LinkStateTE) Ospfv2InterfaceLinkStateTEIter
}

func (obj *ospfv2InterfaceLinkStateTEIter) setMsg(msg *ospfv2Interface) Ospfv2InterfaceLinkStateTEIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&linkStateTE{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2InterfaceLinkStateTEIter) Items() []LinkStateTE {
	return obj.linkStateTESlice
}

func (obj *ospfv2InterfaceLinkStateTEIter) Add() LinkStateTE {
	newObj := &otg.LinkStateTE{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &linkStateTE{obj: newObj}
	newLibObj.setDefault()
	obj.linkStateTESlice = append(obj.linkStateTESlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2InterfaceLinkStateTEIter) Append(items ...LinkStateTE) Ospfv2InterfaceLinkStateTEIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.linkStateTESlice = append(obj.linkStateTESlice, item)
	}
	return obj
}

func (obj *ospfv2InterfaceLinkStateTEIter) Set(index int, newObj LinkStateTE) Ospfv2InterfaceLinkStateTEIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.linkStateTESlice[index] = newObj
	return obj
}
func (obj *ospfv2InterfaceLinkStateTEIter) Clear() Ospfv2InterfaceLinkStateTEIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.LinkStateTE{}
		obj.linkStateTESlice = []LinkStateTE{}
	}
	return obj
}
func (obj *ospfv2InterfaceLinkStateTEIter) clearHolderSlice() Ospfv2InterfaceLinkStateTEIter {
	if len(obj.linkStateTESlice) > 0 {
		obj.linkStateTESlice = []LinkStateTE{}
	}
	return obj
}
func (obj *ospfv2InterfaceLinkStateTEIter) appendHolderSlice(item LinkStateTE) Ospfv2InterfaceLinkStateTEIter {
	obj.linkStateTESlice = append(obj.linkStateTESlice, item)
	return obj
}

// OSPFv2 authentication properties.
// If the authentication is not configured, none OSPF packet exchange is authenticated.
// Authentication returns a Ospfv2InterfaceAuthentication
func (obj *ospfv2Interface) Authentication() Ospfv2InterfaceAuthentication {
	if obj.obj.Authentication == nil {
		obj.obj.Authentication = NewOspfv2InterfaceAuthentication().msg()
	}
	if obj.authenticationHolder == nil {
		obj.authenticationHolder = &ospfv2InterfaceAuthentication{obj: obj.obj.Authentication}
	}
	return obj.authenticationHolder
}

// OSPFv2 authentication properties.
// If the authentication is not configured, none OSPF packet exchange is authenticated.
// Authentication returns a Ospfv2InterfaceAuthentication
func (obj *ospfv2Interface) HasAuthentication() bool {
	return obj.obj.Authentication != nil
}

// OSPFv2 authentication properties.
// If the authentication is not configured, none OSPF packet exchange is authenticated.
// SetAuthentication sets the Ospfv2InterfaceAuthentication value in the Ospfv2Interface object
func (obj *ospfv2Interface) SetAuthentication(value Ospfv2InterfaceAuthentication) Ospfv2Interface {

	obj.authenticationHolder = nil
	obj.obj.Authentication = value.msg()

	return obj
}

// Optional container for advanced interface properties.
// Advanced returns a Ospfv2InterfaceAdvanced
func (obj *ospfv2Interface) Advanced() Ospfv2InterfaceAdvanced {
	if obj.obj.Advanced == nil {
		obj.obj.Advanced = NewOspfv2InterfaceAdvanced().msg()
	}
	if obj.advancedHolder == nil {
		obj.advancedHolder = &ospfv2InterfaceAdvanced{obj: obj.obj.Advanced}
	}
	return obj.advancedHolder
}

// Optional container for advanced interface properties.
// Advanced returns a Ospfv2InterfaceAdvanced
func (obj *ospfv2Interface) HasAdvanced() bool {
	return obj.obj.Advanced != nil
}

// Optional container for advanced interface properties.
// SetAdvanced sets the Ospfv2InterfaceAdvanced value in the Ospfv2Interface object
func (obj *ospfv2Interface) SetAdvanced(value Ospfv2InterfaceAdvanced) Ospfv2Interface {

	obj.advancedHolder = nil
	obj.obj.Advanced = value.msg()

	return obj
}

// Link protection on the OSPFv2 link between two interfaces.
// LinkProtection returns a Ospfv2InterfaceLinkProtection
func (obj *ospfv2Interface) LinkProtection() Ospfv2InterfaceLinkProtection {
	if obj.obj.LinkProtection == nil {
		obj.obj.LinkProtection = NewOspfv2InterfaceLinkProtection().msg()
	}
	if obj.linkProtectionHolder == nil {
		obj.linkProtectionHolder = &ospfv2InterfaceLinkProtection{obj: obj.obj.LinkProtection}
	}
	return obj.linkProtectionHolder
}

// Link protection on the OSPFv2 link between two interfaces.
// LinkProtection returns a Ospfv2InterfaceLinkProtection
func (obj *ospfv2Interface) HasLinkProtection() bool {
	return obj.obj.LinkProtection != nil
}

// Link protection on the OSPFv2 link between two interfaces.
// SetLinkProtection sets the Ospfv2InterfaceLinkProtection value in the Ospfv2Interface object
func (obj *ospfv2Interface) SetLinkProtection(value Ospfv2InterfaceLinkProtection) Ospfv2Interface {

	obj.linkProtectionHolder = nil
	obj.obj.LinkProtection = value.msg()

	return obj
}

// A Shared Risk Link Group (SRLG) is represented by a 32-bit number unique within an IGP (OSPFv2 and IS-IS) domain.
// An SRLG is a set of links sharing a common resource, which affects all links in the set if the common resource fails.
// Links share the same risk of failure and are therefore considered to belong to the same SRLG.
// SrlgValues returns a []uint32
func (obj *ospfv2Interface) SrlgValues() []uint32 {
	if obj.obj.SrlgValues == nil {
		obj.obj.SrlgValues = make([]uint32, 0)
	}
	return obj.obj.SrlgValues
}

// A Shared Risk Link Group (SRLG) is represented by a 32-bit number unique within an IGP (OSPFv2 and IS-IS) domain.
// An SRLG is a set of links sharing a common resource, which affects all links in the set if the common resource fails.
// Links share the same risk of failure and are therefore considered to belong to the same SRLG.
// SetSrlgValues sets the []uint32 value in the Ospfv2Interface object
func (obj *ospfv2Interface) SetSrlgValues(value []uint32) Ospfv2Interface {

	if obj.obj.SrlgValues == nil {
		obj.obj.SrlgValues = make([]uint32, 0)
	}
	obj.obj.SrlgValues = value

	return obj
}

func (obj *ospfv2Interface) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface Ospfv2Interface")
	}

	// Ipv4Name is required
	if obj.obj.Ipv4Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv4Name is required field on interface Ospfv2Interface")
	}

	if obj.obj.Area != nil {

		obj.Area().validateObj(vObj, set_default)
	}

	if obj.obj.NetworkType != nil {

		obj.NetworkType().validateObj(vObj, set_default)
	}

	if len(obj.obj.TrafficEngineering) != 0 {

		if set_default {
			obj.TrafficEngineering().clearHolderSlice()
			for _, item := range obj.obj.TrafficEngineering {
				obj.TrafficEngineering().appendHolderSlice(&linkStateTE{obj: item})
			}
		}
		for _, item := range obj.TrafficEngineering().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Authentication != nil {

		obj.Authentication().validateObj(vObj, set_default)
	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(vObj, set_default)
	}

	if obj.obj.LinkProtection != nil {

		obj.LinkProtection().validateObj(vObj, set_default)
	}

	if obj.obj.SrlgValues != nil {

		for _, item := range obj.obj.SrlgValues {
			if item > 16777215 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= Ospfv2Interface.SrlgValues <= 16777215 but Got %d", item))
			}

		}

	}

}

func (obj *ospfv2Interface) setDefault() {

}

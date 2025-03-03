package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3Interface *****
type ospfv3Interface struct {
	validation
	obj                *otg.Ospfv3Interface
	marshaller         marshalOspfv3Interface
	unMarshaller       unMarshalOspfv3Interface
	areaHolder         Ospfv3InterfaceArea
	networkTypeHolder  Ospfv3InterfaceNetworkType
	capabilitiesHolder Ospfv3InterfaceCapabilities
	optionsHolder      Ospfv3InterfaceOptions
}

func NewOspfv3Interface() Ospfv3Interface {
	obj := ospfv3Interface{obj: &otg.Ospfv3Interface{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3Interface) msg() *otg.Ospfv3Interface {
	return obj.obj
}

func (obj *ospfv3Interface) setMsg(msg *otg.Ospfv3Interface) Ospfv3Interface {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3Interface struct {
	obj *ospfv3Interface
}

type marshalOspfv3Interface interface {
	// ToProto marshals Ospfv3Interface to protobuf object *otg.Ospfv3Interface
	ToProto() (*otg.Ospfv3Interface, error)
	// ToPbText marshals Ospfv3Interface to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3Interface to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3Interface to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3Interface struct {
	obj *ospfv3Interface
}

type unMarshalOspfv3Interface interface {
	// FromProto unmarshals Ospfv3Interface from protobuf object *otg.Ospfv3Interface
	FromProto(msg *otg.Ospfv3Interface) (Ospfv3Interface, error)
	// FromPbText unmarshals Ospfv3Interface from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3Interface from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3Interface from JSON text
	FromJson(value string) error
}

func (obj *ospfv3Interface) Marshal() marshalOspfv3Interface {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3Interface{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3Interface) Unmarshal() unMarshalOspfv3Interface {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3Interface{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3Interface) ToProto() (*otg.Ospfv3Interface, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3Interface) FromProto(msg *otg.Ospfv3Interface) (Ospfv3Interface, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3Interface) ToPbText() (string, error) {
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

func (m *unMarshalospfv3Interface) FromPbText(value string) error {
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

func (m *marshalospfv3Interface) ToYaml() (string, error) {
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

func (m *unMarshalospfv3Interface) FromYaml(value string) error {
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

func (m *marshalospfv3Interface) ToJson() (string, error) {
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

func (m *unMarshalospfv3Interface) FromJson(value string) error {
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

func (obj *ospfv3Interface) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3Interface) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3Interface) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3Interface) Clone() (Ospfv3Interface, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3Interface()
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

func (obj *ospfv3Interface) setNil() {
	obj.areaHolder = nil
	obj.networkTypeHolder = nil
	obj.capabilitiesHolder = nil
	obj.optionsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv3Interface is configuration for single OSPFv3 interface.
type Ospfv3Interface interface {
	Validation
	// msg marshals Ospfv3Interface to protobuf object *otg.Ospfv3Interface
	// and doesn't set defaults
	msg() *otg.Ospfv3Interface
	// setMsg unmarshals Ospfv3Interface from protobuf object *otg.Ospfv3Interface
	// and doesn't set defaults
	setMsg(*otg.Ospfv3Interface) Ospfv3Interface
	// provides marshal interface
	Marshal() marshalOspfv3Interface
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3Interface
	// validate validates Ospfv3Interface
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3Interface, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in Ospfv3Interface.
	Name() string
	// SetName assigns string provided by user to Ospfv3Interface
	SetName(value string) Ospfv3Interface
	// Ipv6Name returns string, set in Ospfv3Interface.
	Ipv6Name() string
	// SetIpv6Name assigns string provided by user to Ospfv3Interface
	SetIpv6Name(value string) Ospfv3Interface
	// Area returns Ospfv3InterfaceArea, set in Ospfv3Interface.
	// Ospfv3InterfaceArea is container for OSPFv3 Area ID identifies the routing area to which the host belongs.
	Area() Ospfv3InterfaceArea
	// SetArea assigns Ospfv3InterfaceArea provided by user to Ospfv3Interface.
	// Ospfv3InterfaceArea is container for OSPFv3 Area ID identifies the routing area to which the host belongs.
	SetArea(value Ospfv3InterfaceArea) Ospfv3Interface
	// HasArea checks if Area has been set in Ospfv3Interface
	HasArea() bool
	// NetworkType returns Ospfv3InterfaceNetworkType, set in Ospfv3Interface.
	// Ospfv3InterfaceNetworkType is the OSPFv3 network link type options.
	// - Broadcast
	// - Point to Point
	NetworkType() Ospfv3InterfaceNetworkType
	// SetNetworkType assigns Ospfv3InterfaceNetworkType provided by user to Ospfv3Interface.
	// Ospfv3InterfaceNetworkType is the OSPFv3 network link type options.
	// - Broadcast
	// - Point to Point
	SetNetworkType(value Ospfv3InterfaceNetworkType) Ospfv3Interface
	// HasNetworkType checks if NetworkType has been set in Ospfv3Interface
	HasNetworkType() bool
	// Capabilities returns Ospfv3InterfaceCapabilities, set in Ospfv3Interface.
	// Ospfv3InterfaceCapabilities is contains OSPFv3 interface capabilities.
	Capabilities() Ospfv3InterfaceCapabilities
	// SetCapabilities assigns Ospfv3InterfaceCapabilities provided by user to Ospfv3Interface.
	// Ospfv3InterfaceCapabilities is contains OSPFv3 interface capabilities.
	SetCapabilities(value Ospfv3InterfaceCapabilities) Ospfv3Interface
	// HasCapabilities checks if Capabilities has been set in Ospfv3Interface
	HasCapabilities() bool
	// Options returns Ospfv3InterfaceOptions, set in Ospfv3Interface.
	// Ospfv3InterfaceOptions is the Options field is present in OSPFv3 Hello packets, Database Description packets and all LSAs.
	// The Options field enables OSPF routers to support (or not support) optional capabilities, and to
	// communicate their capability level to other OSPF routers (https://datatracker.ietf.org/doc/html/rfc2740#appendix-A.2).
	// When used in Hello packets, the Options field allows a router to reject a neighbor because of a capability mismatch.
	Options() Ospfv3InterfaceOptions
	// SetOptions assigns Ospfv3InterfaceOptions provided by user to Ospfv3Interface.
	// Ospfv3InterfaceOptions is the Options field is present in OSPFv3 Hello packets, Database Description packets and all LSAs.
	// The Options field enables OSPF routers to support (or not support) optional capabilities, and to
	// communicate their capability level to other OSPF routers (https://datatracker.ietf.org/doc/html/rfc2740#appendix-A.2).
	// When used in Hello packets, the Options field allows a router to reject a neighbor because of a capability mismatch.
	SetOptions(value Ospfv3InterfaceOptions) Ospfv3Interface
	// HasOptions checks if Options has been set in Ospfv3Interface
	HasOptions() bool
	// InstanceId returns uint32, set in Ospfv3Interface.
	InstanceId() uint32
	// SetInstanceId assigns uint32 provided by user to Ospfv3Interface
	SetInstanceId(value uint32) Ospfv3Interface
	// HasInstanceId checks if InstanceId has been set in Ospfv3Interface
	HasInstanceId() bool
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *ospfv3Interface) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the Ospfv3Interface object
func (obj *ospfv3Interface) SetName(value string) Ospfv3Interface {

	obj.obj.Name = &value
	return obj
}

// The globally unique name of the IPv6 interface connected to the DUT.
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// Ipv6Name returns a string
func (obj *ospfv3Interface) Ipv6Name() string {

	return *obj.obj.Ipv6Name

}

// The globally unique name of the IPv6 interface connected to the DUT.
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// SetIpv6Name sets the string value in the Ospfv3Interface object
func (obj *ospfv3Interface) SetIpv6Name(value string) Ospfv3Interface {

	obj.obj.Ipv6Name = &value
	return obj
}

// The Area ID of the area to which the attached network belongs.
// All routing protocol packets originating from the interface are
// labelled with this Area ID.
// Area returns a Ospfv3InterfaceArea
func (obj *ospfv3Interface) Area() Ospfv3InterfaceArea {
	if obj.obj.Area == nil {
		obj.obj.Area = NewOspfv3InterfaceArea().msg()
	}
	if obj.areaHolder == nil {
		obj.areaHolder = &ospfv3InterfaceArea{obj: obj.obj.Area}
	}
	return obj.areaHolder
}

// The Area ID of the area to which the attached network belongs.
// All routing protocol packets originating from the interface are
// labelled with this Area ID.
// Area returns a Ospfv3InterfaceArea
func (obj *ospfv3Interface) HasArea() bool {
	return obj.obj.Area != nil
}

// The Area ID of the area to which the attached network belongs.
// All routing protocol packets originating from the interface are
// labelled with this Area ID.
// SetArea sets the Ospfv3InterfaceArea value in the Ospfv3Interface object
func (obj *ospfv3Interface) SetArea(value Ospfv3InterfaceArea) Ospfv3Interface {

	obj.areaHolder = nil
	obj.obj.Area = value.msg()

	return obj
}

// The OSPF network link type.
// NetworkType returns a Ospfv3InterfaceNetworkType
func (obj *ospfv3Interface) NetworkType() Ospfv3InterfaceNetworkType {
	if obj.obj.NetworkType == nil {
		obj.obj.NetworkType = NewOspfv3InterfaceNetworkType().msg()
	}
	if obj.networkTypeHolder == nil {
		obj.networkTypeHolder = &ospfv3InterfaceNetworkType{obj: obj.obj.NetworkType}
	}
	return obj.networkTypeHolder
}

// The OSPF network link type.
// NetworkType returns a Ospfv3InterfaceNetworkType
func (obj *ospfv3Interface) HasNetworkType() bool {
	return obj.obj.NetworkType != nil
}

// The OSPF network link type.
// SetNetworkType sets the Ospfv3InterfaceNetworkType value in the Ospfv3Interface object
func (obj *ospfv3Interface) SetNetworkType(value Ospfv3InterfaceNetworkType) Ospfv3Interface {

	obj.networkTypeHolder = nil
	obj.obj.NetworkType = value.msg()

	return obj
}

// Container for OSPFv3 interface capabilities.
// Capabilities returns a Ospfv3InterfaceCapabilities
func (obj *ospfv3Interface) Capabilities() Ospfv3InterfaceCapabilities {
	if obj.obj.Capabilities == nil {
		obj.obj.Capabilities = NewOspfv3InterfaceCapabilities().msg()
	}
	if obj.capabilitiesHolder == nil {
		obj.capabilitiesHolder = &ospfv3InterfaceCapabilities{obj: obj.obj.Capabilities}
	}
	return obj.capabilitiesHolder
}

// Container for OSPFv3 interface capabilities.
// Capabilities returns a Ospfv3InterfaceCapabilities
func (obj *ospfv3Interface) HasCapabilities() bool {
	return obj.obj.Capabilities != nil
}

// Container for OSPFv3 interface capabilities.
// SetCapabilities sets the Ospfv3InterfaceCapabilities value in the Ospfv3Interface object
func (obj *ospfv3Interface) SetCapabilities(value Ospfv3InterfaceCapabilities) Ospfv3Interface {

	obj.capabilitiesHolder = nil
	obj.obj.Capabilities = value.msg()

	return obj
}

// Container for OSPFv3 optional interface properties.
// Options returns a Ospfv3InterfaceOptions
func (obj *ospfv3Interface) Options() Ospfv3InterfaceOptions {
	if obj.obj.Options == nil {
		obj.obj.Options = NewOspfv3InterfaceOptions().msg()
	}
	if obj.optionsHolder == nil {
		obj.optionsHolder = &ospfv3InterfaceOptions{obj: obj.obj.Options}
	}
	return obj.optionsHolder
}

// Container for OSPFv3 optional interface properties.
// Options returns a Ospfv3InterfaceOptions
func (obj *ospfv3Interface) HasOptions() bool {
	return obj.obj.Options != nil
}

// Container for OSPFv3 optional interface properties.
// SetOptions sets the Ospfv3InterfaceOptions value in the Ospfv3Interface object
func (obj *ospfv3Interface) SetOptions(value Ospfv3InterfaceOptions) Ospfv3Interface {

	obj.optionsHolder = nil
	obj.obj.Options = value.msg()

	return obj
}

// Enables multiple instances of OSPF to be run over a single link.
// Each protocol instance should be assigned a separate Instance ID;
// the Instance ID has link-local significance only.
// InstanceId returns a uint32
func (obj *ospfv3Interface) InstanceId() uint32 {

	return *obj.obj.InstanceId

}

// Enables multiple instances of OSPF to be run over a single link.
// Each protocol instance should be assigned a separate Instance ID;
// the Instance ID has link-local significance only.
// InstanceId returns a uint32
func (obj *ospfv3Interface) HasInstanceId() bool {
	return obj.obj.InstanceId != nil
}

// Enables multiple instances of OSPF to be run over a single link.
// Each protocol instance should be assigned a separate Instance ID;
// the Instance ID has link-local significance only.
// SetInstanceId sets the uint32 value in the Ospfv3Interface object
func (obj *ospfv3Interface) SetInstanceId(value uint32) Ospfv3Interface {

	obj.obj.InstanceId = &value
	return obj
}

func (obj *ospfv3Interface) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface Ospfv3Interface")
	}

	// Ipv6Name is required
	if obj.obj.Ipv6Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv6Name is required field on interface Ospfv3Interface")
	}

	if obj.obj.Area != nil {

		obj.Area().validateObj(vObj, set_default)
	}

	if obj.obj.NetworkType != nil {

		obj.NetworkType().validateObj(vObj, set_default)
	}

	if obj.obj.Capabilities != nil {

		obj.Capabilities().validateObj(vObj, set_default)
	}

	if obj.obj.Options != nil {

		obj.Options().validateObj(vObj, set_default)
	}

	if obj.obj.InstanceId != nil {

		if *obj.obj.InstanceId > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv3Interface.InstanceId <= 255 but Got %d", *obj.obj.InstanceId))
		}

	}

}

func (obj *ospfv3Interface) setDefault() {
	if obj.obj.InstanceId == nil {
		obj.SetInstanceId(0)
	}

}

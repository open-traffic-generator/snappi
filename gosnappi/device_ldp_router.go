package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceLdpRouter *****
type deviceLdpRouter struct {
	validation
	obj                     *otg.DeviceLdpRouter
	marshaller              marshalDeviceLdpRouter
	unMarshaller            unMarshalDeviceLdpRouter
	gracefulRestartHolder   LdpGracefulRestart
	ipv4InterfacesHolder    DeviceLdpRouterLdpIpv4InterfaceIter
	ipv4TargetedPeersHolder DeviceLdpRouterLdpIpv4TargetedPeerIter
	ipv4FecRangesHolder     DeviceLdpRouterLdpIpv4FecRangeIter
}

func NewDeviceLdpRouter() DeviceLdpRouter {
	obj := deviceLdpRouter{obj: &otg.DeviceLdpRouter{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceLdpRouter) msg() *otg.DeviceLdpRouter {
	return obj.obj
}

func (obj *deviceLdpRouter) setMsg(msg *otg.DeviceLdpRouter) DeviceLdpRouter {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceLdpRouter struct {
	obj *deviceLdpRouter
}

type marshalDeviceLdpRouter interface {
	// ToProto marshals DeviceLdpRouter to protobuf object *otg.DeviceLdpRouter
	ToProto() (*otg.DeviceLdpRouter, error)
	// ToPbText marshals DeviceLdpRouter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceLdpRouter to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceLdpRouter to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceLdpRouter struct {
	obj *deviceLdpRouter
}

type unMarshalDeviceLdpRouter interface {
	// FromProto unmarshals DeviceLdpRouter from protobuf object *otg.DeviceLdpRouter
	FromProto(msg *otg.DeviceLdpRouter) (DeviceLdpRouter, error)
	// FromPbText unmarshals DeviceLdpRouter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceLdpRouter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceLdpRouter from JSON text
	FromJson(value string) error
}

func (obj *deviceLdpRouter) Marshal() marshalDeviceLdpRouter {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceLdpRouter{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceLdpRouter) Unmarshal() unMarshalDeviceLdpRouter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceLdpRouter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceLdpRouter) ToProto() (*otg.DeviceLdpRouter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceLdpRouter) FromProto(msg *otg.DeviceLdpRouter) (DeviceLdpRouter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceLdpRouter) ToPbText() (string, error) {
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

func (m *unMarshaldeviceLdpRouter) FromPbText(value string) error {
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

func (m *marshaldeviceLdpRouter) ToYaml() (string, error) {
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

func (m *unMarshaldeviceLdpRouter) FromYaml(value string) error {
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

func (m *marshaldeviceLdpRouter) ToJson() (string, error) {
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

func (m *unMarshaldeviceLdpRouter) FromJson(value string) error {
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

func (obj *deviceLdpRouter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceLdpRouter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceLdpRouter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceLdpRouter) Clone() (DeviceLdpRouter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceLdpRouter()
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

func (obj *deviceLdpRouter) setNil() {
	obj.gracefulRestartHolder = nil
	obj.ipv4InterfacesHolder = nil
	obj.ipv4TargetedPeersHolder = nil
	obj.ipv4FecRangesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceLdpRouter is configuration for an emulated LDP Label Switching Router (LSR) as per RFC 5036. In this model, LDP over IPv4 is supported, with Basic (link) discovery on IPv4 interfaces (RFC 5036 Section 2.4.1) and Extended (targeted) discovery to configured peers (RFC 5036 Section 2.4.2). The LDP Identifier carried in every LDP PDU is lsr_id:label_space_id (RFC 5036 Section 2.2.2). At least one entry in ipv4_interfaces or ipv4_targeted_peers is needed to form an LDP session. All sessions of this router use the session parameters configured here (label space, label advertisement mode, keepalive and graceful restart).
type DeviceLdpRouter interface {
	Validation
	// msg marshals DeviceLdpRouter to protobuf object *otg.DeviceLdpRouter
	// and doesn't set defaults
	msg() *otg.DeviceLdpRouter
	// setMsg unmarshals DeviceLdpRouter from protobuf object *otg.DeviceLdpRouter
	// and doesn't set defaults
	setMsg(*otg.DeviceLdpRouter) DeviceLdpRouter
	// provides marshal interface
	Marshal() marshalDeviceLdpRouter
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceLdpRouter
	// validate validates DeviceLdpRouter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceLdpRouter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in DeviceLdpRouter.
	Name() string
	// SetName assigns string provided by user to DeviceLdpRouter
	SetName(value string) DeviceLdpRouter
	// LsrId returns string, set in DeviceLdpRouter.
	LsrId() string
	// SetLsrId assigns string provided by user to DeviceLdpRouter
	SetLsrId(value string) DeviceLdpRouter
	// LabelSpaceId returns uint32, set in DeviceLdpRouter.
	LabelSpaceId() uint32
	// SetLabelSpaceId assigns uint32 provided by user to DeviceLdpRouter
	SetLabelSpaceId(value uint32) DeviceLdpRouter
	// HasLabelSpaceId checks if LabelSpaceId has been set in DeviceLdpRouter
	HasLabelSpaceId() bool
	// LabelAdvertisementMode returns DeviceLdpRouterLabelAdvertisementModeEnum, set in DeviceLdpRouter
	LabelAdvertisementMode() DeviceLdpRouterLabelAdvertisementModeEnum
	// SetLabelAdvertisementMode assigns DeviceLdpRouterLabelAdvertisementModeEnum provided by user to DeviceLdpRouter
	SetLabelAdvertisementMode(value DeviceLdpRouterLabelAdvertisementModeEnum) DeviceLdpRouter
	// HasLabelAdvertisementMode checks if LabelAdvertisementMode has been set in DeviceLdpRouter
	HasLabelAdvertisementMode() bool
	// KeepaliveHoldTime returns uint32, set in DeviceLdpRouter.
	KeepaliveHoldTime() uint32
	// SetKeepaliveHoldTime assigns uint32 provided by user to DeviceLdpRouter
	SetKeepaliveHoldTime(value uint32) DeviceLdpRouter
	// HasKeepaliveHoldTime checks if KeepaliveHoldTime has been set in DeviceLdpRouter
	HasKeepaliveHoldTime() bool
	// KeepaliveInterval returns uint32, set in DeviceLdpRouter.
	KeepaliveInterval() uint32
	// SetKeepaliveInterval assigns uint32 provided by user to DeviceLdpRouter
	SetKeepaliveInterval(value uint32) DeviceLdpRouter
	// HasKeepaliveInterval checks if KeepaliveInterval has been set in DeviceLdpRouter
	HasKeepaliveInterval() bool
	// GracefulRestart returns LdpGracefulRestart, set in DeviceLdpRouter.
	// LdpGracefulRestart is lDP graceful restart parameters (RFC 3478). When this object is present, the emulated LSR includes the Fault Tolerant (FT) Session TLV (type 0x0503) with the L flag set in its Initialization message (RFC 3478 Section 2). Both timers are configured in seconds; they are carried in milliseconds on the wire, and the implementation converts them.
	GracefulRestart() LdpGracefulRestart
	// SetGracefulRestart assigns LdpGracefulRestart provided by user to DeviceLdpRouter.
	// LdpGracefulRestart is lDP graceful restart parameters (RFC 3478). When this object is present, the emulated LSR includes the Fault Tolerant (FT) Session TLV (type 0x0503) with the L flag set in its Initialization message (RFC 3478 Section 2). Both timers are configured in seconds; they are carried in milliseconds on the wire, and the implementation converts them.
	SetGracefulRestart(value LdpGracefulRestart) DeviceLdpRouter
	// HasGracefulRestart checks if GracefulRestart has been set in DeviceLdpRouter
	HasGracefulRestart() bool
	// Ipv4Interfaces returns DeviceLdpRouterLdpIpv4InterfaceIterIter, set in DeviceLdpRouter
	Ipv4Interfaces() DeviceLdpRouterLdpIpv4InterfaceIter
	// Ipv4TargetedPeers returns DeviceLdpRouterLdpIpv4TargetedPeerIterIter, set in DeviceLdpRouter
	Ipv4TargetedPeers() DeviceLdpRouterLdpIpv4TargetedPeerIter
	// Ipv4FecRanges returns DeviceLdpRouterLdpIpv4FecRangeIterIter, set in DeviceLdpRouter
	Ipv4FecRanges() DeviceLdpRouterLdpIpv4FecRangeIter
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *deviceLdpRouter) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the DeviceLdpRouter object
func (obj *deviceLdpRouter) SetName(value string) DeviceLdpRouter {

	obj.obj.Name = &value
	return obj
}

// The 4-octet LSR ID, the first part of the 6-octet LDP Identifier (RFC 5036 Section 2.2.2). It must be globally unique in the network and is normally the address of an IPv4 loopback of this device. An implementation that uses one router ID per emulated device requires lsr_id to be equal to any other router ID configured on the same device, such as Device.BgpRouter.router_id or a custom OSPFv2 / OSPFv3 router ID, and rejects a mismatch.
// LsrId returns a string
func (obj *deviceLdpRouter) LsrId() string {

	return *obj.obj.LsrId

}

// The 4-octet LSR ID, the first part of the 6-octet LDP Identifier (RFC 5036 Section 2.2.2). It must be globally unique in the network and is normally the address of an IPv4 loopback of this device. An implementation that uses one router ID per emulated device requires lsr_id to be equal to any other router ID configured on the same device, such as Device.BgpRouter.router_id or a custom OSPFv2 / OSPFv3 router ID, and rejects a mismatch.
// SetLsrId sets the string value in the DeviceLdpRouter object
func (obj *deviceLdpRouter) SetLsrId(value string) DeviceLdpRouter {

	obj.obj.LsrId = &value
	return obj
}

// The 2-octet label space identifier, the second part of the LDP Identifier (RFC 5036 Section 2.2.2). The value 0 identifies the platform-wide label space (RFC 5036 Section 2.2.1), which is the normal value for IPv4 LDP over Ethernet.
// LabelSpaceId returns a uint32
func (obj *deviceLdpRouter) LabelSpaceId() uint32 {

	return *obj.obj.LabelSpaceId

}

// The 2-octet label space identifier, the second part of the LDP Identifier (RFC 5036 Section 2.2.2). The value 0 identifies the platform-wide label space (RFC 5036 Section 2.2.1), which is the normal value for IPv4 LDP over Ethernet.
// LabelSpaceId returns a uint32
func (obj *deviceLdpRouter) HasLabelSpaceId() bool {
	return obj.obj.LabelSpaceId != nil
}

// The 2-octet label space identifier, the second part of the LDP Identifier (RFC 5036 Section 2.2.2). The value 0 identifies the platform-wide label space (RFC 5036 Section 2.2.1), which is the normal value for IPv4 LDP over Ethernet.
// SetLabelSpaceId sets the uint32 value in the DeviceLdpRouter object
func (obj *deviceLdpRouter) SetLabelSpaceId(value uint32) DeviceLdpRouter {

	obj.obj.LabelSpaceId = &value
	return obj
}

type DeviceLdpRouterLabelAdvertisementModeEnum string

// Enum of LabelAdvertisementMode on DeviceLdpRouter
var DeviceLdpRouterLabelAdvertisementMode = struct {
	DOWNSTREAM_UNSOLICITED DeviceLdpRouterLabelAdvertisementModeEnum
	DOWNSTREAM_ON_DEMAND   DeviceLdpRouterLabelAdvertisementModeEnum
}{
	DOWNSTREAM_UNSOLICITED: DeviceLdpRouterLabelAdvertisementModeEnum("downstream_unsolicited"),
	DOWNSTREAM_ON_DEMAND:   DeviceLdpRouterLabelAdvertisementModeEnum("downstream_on_demand"),
}

func (obj *deviceLdpRouter) LabelAdvertisementMode() DeviceLdpRouterLabelAdvertisementModeEnum {
	return DeviceLdpRouterLabelAdvertisementModeEnum(obj.obj.LabelAdvertisementMode.Enum().String())
}

// The label advertisement discipline proposed by this LSR in the A bit of the Common Session Parameters TLV of the Initialization message (RFC 5036 Sections 2.6.3 and 3.5.3). It applies to every session of this router. downstream_unsolicited - labels for the FECs in ipv4_fec_ranges are advertised to every peer without a request (A bit = 0). downstream_on_demand - a label for a FEC is advertised to a peer only after that peer sends a Label Request for it (A bit = 1).
// LabelAdvertisementMode returns a string
func (obj *deviceLdpRouter) HasLabelAdvertisementMode() bool {
	return obj.obj.LabelAdvertisementMode != nil
}

func (obj *deviceLdpRouter) SetLabelAdvertisementMode(value DeviceLdpRouterLabelAdvertisementModeEnum) DeviceLdpRouter {
	intValue, ok := otg.DeviceLdpRouter_LabelAdvertisementMode_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on DeviceLdpRouterLabelAdvertisementModeEnum", string(value)))
		return obj
	}
	enumValue := otg.DeviceLdpRouter_LabelAdvertisementMode_Enum(intValue)
	obj.obj.LabelAdvertisementMode = &enumValue

	return obj
}

// The KeepAlive Time in seconds advertised in the Common Session Parameters TLV of the Initialization message (RFC 5036 Section 3.5.3). A session is closed when no LDP PDU is received from the peer within the negotiated KeepAlive Time, which is the smaller of the two advertised values. The value must be non-zero.
// KeepaliveHoldTime returns a uint32
func (obj *deviceLdpRouter) KeepaliveHoldTime() uint32 {

	return *obj.obj.KeepaliveHoldTime

}

// The KeepAlive Time in seconds advertised in the Common Session Parameters TLV of the Initialization message (RFC 5036 Section 3.5.3). A session is closed when no LDP PDU is received from the peer within the negotiated KeepAlive Time, which is the smaller of the two advertised values. The value must be non-zero.
// KeepaliveHoldTime returns a uint32
func (obj *deviceLdpRouter) HasKeepaliveHoldTime() bool {
	return obj.obj.KeepaliveHoldTime != nil
}

// The KeepAlive Time in seconds advertised in the Common Session Parameters TLV of the Initialization message (RFC 5036 Section 3.5.3). A session is closed when no LDP PDU is received from the peer within the negotiated KeepAlive Time, which is the smaller of the two advertised values. The value must be non-zero.
// SetKeepaliveHoldTime sets the uint32 value in the DeviceLdpRouter object
func (obj *deviceLdpRouter) SetKeepaliveHoldTime(value uint32) DeviceLdpRouter {

	obj.obj.KeepaliveHoldTime = &value
	return obj
}

// The interval in seconds at which KeepAlive messages are sent when no other LDP PDU has been sent (RFC 5036 Sections 2.5.6 and 3.5.4). It should be smaller than keepalive_hold_time. One third of keepalive_hold_time is common.
// KeepaliveInterval returns a uint32
func (obj *deviceLdpRouter) KeepaliveInterval() uint32 {

	return *obj.obj.KeepaliveInterval

}

// The interval in seconds at which KeepAlive messages are sent when no other LDP PDU has been sent (RFC 5036 Sections 2.5.6 and 3.5.4). It should be smaller than keepalive_hold_time. One third of keepalive_hold_time is common.
// KeepaliveInterval returns a uint32
func (obj *deviceLdpRouter) HasKeepaliveInterval() bool {
	return obj.obj.KeepaliveInterval != nil
}

// The interval in seconds at which KeepAlive messages are sent when no other LDP PDU has been sent (RFC 5036 Sections 2.5.6 and 3.5.4). It should be smaller than keepalive_hold_time. One third of keepalive_hold_time is common.
// SetKeepaliveInterval sets the uint32 value in the DeviceLdpRouter object
func (obj *deviceLdpRouter) SetKeepaliveInterval(value uint32) DeviceLdpRouter {

	obj.obj.KeepaliveInterval = &value
	return obj
}

// When present, LDP graceful restart (RFC 3478) is enabled and the Fault Tolerant Session TLV is included in the Initialization message. Omit to disable graceful restart.
// GracefulRestart returns a LdpGracefulRestart
func (obj *deviceLdpRouter) GracefulRestart() LdpGracefulRestart {
	if obj.obj.GracefulRestart == nil {
		obj.obj.GracefulRestart = NewLdpGracefulRestart().msg()
	}
	if obj.gracefulRestartHolder == nil {
		obj.gracefulRestartHolder = &ldpGracefulRestart{obj: obj.obj.GracefulRestart}
	}
	return obj.gracefulRestartHolder
}

// When present, LDP graceful restart (RFC 3478) is enabled and the Fault Tolerant Session TLV is included in the Initialization message. Omit to disable graceful restart.
// GracefulRestart returns a LdpGracefulRestart
func (obj *deviceLdpRouter) HasGracefulRestart() bool {
	return obj.obj.GracefulRestart != nil
}

// When present, LDP graceful restart (RFC 3478) is enabled and the Fault Tolerant Session TLV is included in the Initialization message. Omit to disable graceful restart.
// SetGracefulRestart sets the LdpGracefulRestart value in the DeviceLdpRouter object
func (obj *deviceLdpRouter) SetGracefulRestart(value LdpGracefulRestart) DeviceLdpRouter {

	obj.gracefulRestartHolder = nil
	obj.obj.GracefulRestart = value.msg()

	return obj
}

// IPv4 interfaces on which LDP Basic Discovery runs (RFC 5036 Section 2.4.1). A session is formed with each LSR discovered on the link.
// Ipv4Interfaces returns a []LdpIpv4Interface
func (obj *deviceLdpRouter) Ipv4Interfaces() DeviceLdpRouterLdpIpv4InterfaceIter {
	if len(obj.obj.Ipv4Interfaces) == 0 {
		obj.obj.Ipv4Interfaces = []*otg.LdpIpv4Interface{}
	}
	if obj.ipv4InterfacesHolder == nil {
		obj.ipv4InterfacesHolder = newDeviceLdpRouterLdpIpv4InterfaceIter(&obj.obj.Ipv4Interfaces).setMsg(obj)
	}
	return obj.ipv4InterfacesHolder
}

type deviceLdpRouterLdpIpv4InterfaceIter struct {
	obj                   *deviceLdpRouter
	ldpIpv4InterfaceSlice []LdpIpv4Interface
	fieldPtr              *[]*otg.LdpIpv4Interface
}

func newDeviceLdpRouterLdpIpv4InterfaceIter(ptr *[]*otg.LdpIpv4Interface) DeviceLdpRouterLdpIpv4InterfaceIter {
	return &deviceLdpRouterLdpIpv4InterfaceIter{fieldPtr: ptr}
}

type DeviceLdpRouterLdpIpv4InterfaceIter interface {
	setMsg(*deviceLdpRouter) DeviceLdpRouterLdpIpv4InterfaceIter
	Items() []LdpIpv4Interface
	Add() LdpIpv4Interface
	Append(items ...LdpIpv4Interface) DeviceLdpRouterLdpIpv4InterfaceIter
	Set(index int, newObj LdpIpv4Interface) DeviceLdpRouterLdpIpv4InterfaceIter
	Clear() DeviceLdpRouterLdpIpv4InterfaceIter
	clearHolderSlice() DeviceLdpRouterLdpIpv4InterfaceIter
	appendHolderSlice(item LdpIpv4Interface) DeviceLdpRouterLdpIpv4InterfaceIter
}

func (obj *deviceLdpRouterLdpIpv4InterfaceIter) setMsg(msg *deviceLdpRouter) DeviceLdpRouterLdpIpv4InterfaceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ldpIpv4Interface{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceLdpRouterLdpIpv4InterfaceIter) Items() []LdpIpv4Interface {
	return obj.ldpIpv4InterfaceSlice
}

func (obj *deviceLdpRouterLdpIpv4InterfaceIter) Add() LdpIpv4Interface {
	newObj := &otg.LdpIpv4Interface{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ldpIpv4Interface{obj: newObj}
	newLibObj.setDefault()
	obj.ldpIpv4InterfaceSlice = append(obj.ldpIpv4InterfaceSlice, newLibObj)
	return newLibObj
}

func (obj *deviceLdpRouterLdpIpv4InterfaceIter) Append(items ...LdpIpv4Interface) DeviceLdpRouterLdpIpv4InterfaceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ldpIpv4InterfaceSlice = append(obj.ldpIpv4InterfaceSlice, item)
	}
	return obj
}

func (obj *deviceLdpRouterLdpIpv4InterfaceIter) Set(index int, newObj LdpIpv4Interface) DeviceLdpRouterLdpIpv4InterfaceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ldpIpv4InterfaceSlice[index] = newObj
	return obj
}
func (obj *deviceLdpRouterLdpIpv4InterfaceIter) Clear() DeviceLdpRouterLdpIpv4InterfaceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.LdpIpv4Interface{}
		obj.ldpIpv4InterfaceSlice = []LdpIpv4Interface{}
	}
	return obj
}
func (obj *deviceLdpRouterLdpIpv4InterfaceIter) clearHolderSlice() DeviceLdpRouterLdpIpv4InterfaceIter {
	if len(obj.ldpIpv4InterfaceSlice) > 0 {
		obj.ldpIpv4InterfaceSlice = []LdpIpv4Interface{}
	}
	return obj
}
func (obj *deviceLdpRouterLdpIpv4InterfaceIter) appendHolderSlice(item LdpIpv4Interface) DeviceLdpRouterLdpIpv4InterfaceIter {
	obj.ldpIpv4InterfaceSlice = append(obj.ldpIpv4InterfaceSlice, item)
	return obj
}

// IPv4 peers to which LDP Extended Discovery runs (RFC 5036 Section 2.4.2). A session is formed with each configured peer.
// Ipv4TargetedPeers returns a []LdpIpv4TargetedPeer
func (obj *deviceLdpRouter) Ipv4TargetedPeers() DeviceLdpRouterLdpIpv4TargetedPeerIter {
	if len(obj.obj.Ipv4TargetedPeers) == 0 {
		obj.obj.Ipv4TargetedPeers = []*otg.LdpIpv4TargetedPeer{}
	}
	if obj.ipv4TargetedPeersHolder == nil {
		obj.ipv4TargetedPeersHolder = newDeviceLdpRouterLdpIpv4TargetedPeerIter(&obj.obj.Ipv4TargetedPeers).setMsg(obj)
	}
	return obj.ipv4TargetedPeersHolder
}

type deviceLdpRouterLdpIpv4TargetedPeerIter struct {
	obj                      *deviceLdpRouter
	ldpIpv4TargetedPeerSlice []LdpIpv4TargetedPeer
	fieldPtr                 *[]*otg.LdpIpv4TargetedPeer
}

func newDeviceLdpRouterLdpIpv4TargetedPeerIter(ptr *[]*otg.LdpIpv4TargetedPeer) DeviceLdpRouterLdpIpv4TargetedPeerIter {
	return &deviceLdpRouterLdpIpv4TargetedPeerIter{fieldPtr: ptr}
}

type DeviceLdpRouterLdpIpv4TargetedPeerIter interface {
	setMsg(*deviceLdpRouter) DeviceLdpRouterLdpIpv4TargetedPeerIter
	Items() []LdpIpv4TargetedPeer
	Add() LdpIpv4TargetedPeer
	Append(items ...LdpIpv4TargetedPeer) DeviceLdpRouterLdpIpv4TargetedPeerIter
	Set(index int, newObj LdpIpv4TargetedPeer) DeviceLdpRouterLdpIpv4TargetedPeerIter
	Clear() DeviceLdpRouterLdpIpv4TargetedPeerIter
	clearHolderSlice() DeviceLdpRouterLdpIpv4TargetedPeerIter
	appendHolderSlice(item LdpIpv4TargetedPeer) DeviceLdpRouterLdpIpv4TargetedPeerIter
}

func (obj *deviceLdpRouterLdpIpv4TargetedPeerIter) setMsg(msg *deviceLdpRouter) DeviceLdpRouterLdpIpv4TargetedPeerIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ldpIpv4TargetedPeer{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceLdpRouterLdpIpv4TargetedPeerIter) Items() []LdpIpv4TargetedPeer {
	return obj.ldpIpv4TargetedPeerSlice
}

func (obj *deviceLdpRouterLdpIpv4TargetedPeerIter) Add() LdpIpv4TargetedPeer {
	newObj := &otg.LdpIpv4TargetedPeer{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ldpIpv4TargetedPeer{obj: newObj}
	newLibObj.setDefault()
	obj.ldpIpv4TargetedPeerSlice = append(obj.ldpIpv4TargetedPeerSlice, newLibObj)
	return newLibObj
}

func (obj *deviceLdpRouterLdpIpv4TargetedPeerIter) Append(items ...LdpIpv4TargetedPeer) DeviceLdpRouterLdpIpv4TargetedPeerIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ldpIpv4TargetedPeerSlice = append(obj.ldpIpv4TargetedPeerSlice, item)
	}
	return obj
}

func (obj *deviceLdpRouterLdpIpv4TargetedPeerIter) Set(index int, newObj LdpIpv4TargetedPeer) DeviceLdpRouterLdpIpv4TargetedPeerIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ldpIpv4TargetedPeerSlice[index] = newObj
	return obj
}
func (obj *deviceLdpRouterLdpIpv4TargetedPeerIter) Clear() DeviceLdpRouterLdpIpv4TargetedPeerIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.LdpIpv4TargetedPeer{}
		obj.ldpIpv4TargetedPeerSlice = []LdpIpv4TargetedPeer{}
	}
	return obj
}
func (obj *deviceLdpRouterLdpIpv4TargetedPeerIter) clearHolderSlice() DeviceLdpRouterLdpIpv4TargetedPeerIter {
	if len(obj.ldpIpv4TargetedPeerSlice) > 0 {
		obj.ldpIpv4TargetedPeerSlice = []LdpIpv4TargetedPeer{}
	}
	return obj
}
func (obj *deviceLdpRouterLdpIpv4TargetedPeerIter) appendHolderSlice(item LdpIpv4TargetedPeer) DeviceLdpRouterLdpIpv4TargetedPeerIter {
	obj.ldpIpv4TargetedPeerSlice = append(obj.ldpIpv4TargetedPeerSlice, item)
	return obj
}

// IPv4 Prefix FECs (RFC 5036 Section 3.4.1) that this LSR advertises to all its peers in Label Mapping messages.
// Ipv4FecRanges returns a []LdpIpv4FecRange
func (obj *deviceLdpRouter) Ipv4FecRanges() DeviceLdpRouterLdpIpv4FecRangeIter {
	if len(obj.obj.Ipv4FecRanges) == 0 {
		obj.obj.Ipv4FecRanges = []*otg.LdpIpv4FecRange{}
	}
	if obj.ipv4FecRangesHolder == nil {
		obj.ipv4FecRangesHolder = newDeviceLdpRouterLdpIpv4FecRangeIter(&obj.obj.Ipv4FecRanges).setMsg(obj)
	}
	return obj.ipv4FecRangesHolder
}

type deviceLdpRouterLdpIpv4FecRangeIter struct {
	obj                  *deviceLdpRouter
	ldpIpv4FecRangeSlice []LdpIpv4FecRange
	fieldPtr             *[]*otg.LdpIpv4FecRange
}

func newDeviceLdpRouterLdpIpv4FecRangeIter(ptr *[]*otg.LdpIpv4FecRange) DeviceLdpRouterLdpIpv4FecRangeIter {
	return &deviceLdpRouterLdpIpv4FecRangeIter{fieldPtr: ptr}
}

type DeviceLdpRouterLdpIpv4FecRangeIter interface {
	setMsg(*deviceLdpRouter) DeviceLdpRouterLdpIpv4FecRangeIter
	Items() []LdpIpv4FecRange
	Add() LdpIpv4FecRange
	Append(items ...LdpIpv4FecRange) DeviceLdpRouterLdpIpv4FecRangeIter
	Set(index int, newObj LdpIpv4FecRange) DeviceLdpRouterLdpIpv4FecRangeIter
	Clear() DeviceLdpRouterLdpIpv4FecRangeIter
	clearHolderSlice() DeviceLdpRouterLdpIpv4FecRangeIter
	appendHolderSlice(item LdpIpv4FecRange) DeviceLdpRouterLdpIpv4FecRangeIter
}

func (obj *deviceLdpRouterLdpIpv4FecRangeIter) setMsg(msg *deviceLdpRouter) DeviceLdpRouterLdpIpv4FecRangeIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ldpIpv4FecRange{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceLdpRouterLdpIpv4FecRangeIter) Items() []LdpIpv4FecRange {
	return obj.ldpIpv4FecRangeSlice
}

func (obj *deviceLdpRouterLdpIpv4FecRangeIter) Add() LdpIpv4FecRange {
	newObj := &otg.LdpIpv4FecRange{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ldpIpv4FecRange{obj: newObj}
	newLibObj.setDefault()
	obj.ldpIpv4FecRangeSlice = append(obj.ldpIpv4FecRangeSlice, newLibObj)
	return newLibObj
}

func (obj *deviceLdpRouterLdpIpv4FecRangeIter) Append(items ...LdpIpv4FecRange) DeviceLdpRouterLdpIpv4FecRangeIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ldpIpv4FecRangeSlice = append(obj.ldpIpv4FecRangeSlice, item)
	}
	return obj
}

func (obj *deviceLdpRouterLdpIpv4FecRangeIter) Set(index int, newObj LdpIpv4FecRange) DeviceLdpRouterLdpIpv4FecRangeIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ldpIpv4FecRangeSlice[index] = newObj
	return obj
}
func (obj *deviceLdpRouterLdpIpv4FecRangeIter) Clear() DeviceLdpRouterLdpIpv4FecRangeIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.LdpIpv4FecRange{}
		obj.ldpIpv4FecRangeSlice = []LdpIpv4FecRange{}
	}
	return obj
}
func (obj *deviceLdpRouterLdpIpv4FecRangeIter) clearHolderSlice() DeviceLdpRouterLdpIpv4FecRangeIter {
	if len(obj.ldpIpv4FecRangeSlice) > 0 {
		obj.ldpIpv4FecRangeSlice = []LdpIpv4FecRange{}
	}
	return obj
}
func (obj *deviceLdpRouterLdpIpv4FecRangeIter) appendHolderSlice(item LdpIpv4FecRange) DeviceLdpRouterLdpIpv4FecRangeIter {
	obj.ldpIpv4FecRangeSlice = append(obj.ldpIpv4FecRangeSlice, item)
	return obj
}

func (obj *deviceLdpRouter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface DeviceLdpRouter")
	}

	// LsrId is required
	if obj.obj.LsrId == nil {
		vObj.validationErrors = append(vObj.validationErrors, "LsrId is required field on interface DeviceLdpRouter")
	}
	if obj.obj.LsrId != nil {

		err := obj.validateIpv4(obj.LsrId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DeviceLdpRouter.LsrId"))
		}

	}

	if obj.obj.LabelSpaceId != nil {

		if *obj.obj.LabelSpaceId > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= DeviceLdpRouter.LabelSpaceId <= 65535 but Got %d", *obj.obj.LabelSpaceId))
		}

	}

	if obj.obj.KeepaliveHoldTime != nil {

		if *obj.obj.KeepaliveHoldTime < 1 || *obj.obj.KeepaliveHoldTime > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= DeviceLdpRouter.KeepaliveHoldTime <= 65535 but Got %d", *obj.obj.KeepaliveHoldTime))
		}

	}

	if obj.obj.KeepaliveInterval != nil {

		if *obj.obj.KeepaliveInterval < 1 || *obj.obj.KeepaliveInterval > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= DeviceLdpRouter.KeepaliveInterval <= 65535 but Got %d", *obj.obj.KeepaliveInterval))
		}

	}

	if obj.obj.GracefulRestart != nil {

		obj.GracefulRestart().validateObj(vObj, set_default)
	}

	if len(obj.obj.Ipv4Interfaces) != 0 {

		if set_default {
			obj.Ipv4Interfaces().clearHolderSlice()
			for _, item := range obj.obj.Ipv4Interfaces {
				obj.Ipv4Interfaces().appendHolderSlice(&ldpIpv4Interface{obj: item})
			}
		}
		for _, item := range obj.Ipv4Interfaces().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Ipv4TargetedPeers) != 0 {

		if set_default {
			obj.Ipv4TargetedPeers().clearHolderSlice()
			for _, item := range obj.obj.Ipv4TargetedPeers {
				obj.Ipv4TargetedPeers().appendHolderSlice(&ldpIpv4TargetedPeer{obj: item})
			}
		}
		for _, item := range obj.Ipv4TargetedPeers().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Ipv4FecRanges) != 0 {

		if set_default {
			obj.Ipv4FecRanges().clearHolderSlice()
			for _, item := range obj.obj.Ipv4FecRanges {
				obj.Ipv4FecRanges().appendHolderSlice(&ldpIpv4FecRange{obj: item})
			}
		}
		for _, item := range obj.Ipv4FecRanges().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *deviceLdpRouter) setDefault() {
	if obj.obj.LabelSpaceId == nil {
		obj.SetLabelSpaceId(0)
	}
	if obj.obj.LabelAdvertisementMode == nil {
		obj.SetLabelAdvertisementMode(DeviceLdpRouterLabelAdvertisementMode.DOWNSTREAM_UNSOLICITED)

	}
	if obj.obj.KeepaliveHoldTime == nil {
		obj.SetKeepaliveHoldTime(180)
	}
	if obj.obj.KeepaliveInterval == nil {
		obj.SetKeepaliveInterval(60)
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceOspfv2Router *****
type deviceOspfv2Router struct {
	validation
	obj                   *otg.DeviceOspfv2Router
	marshaller            marshalDeviceOspfv2Router
	unMarshaller          unMarshalDeviceOspfv2Router
	routerIdHolder        Ospfv2RouterId
	gracefulRestartHolder Ospfv2GracefulRestart
	capabilitiesHolder    Ospfv2Options
	interfacesHolder      DeviceOspfv2RouterOspfv2InterfaceIter
	v4RoutesHolder        DeviceOspfv2RouterOspfv2V4RouteRangeIter
}

func NewDeviceOspfv2Router() DeviceOspfv2Router {
	obj := deviceOspfv2Router{obj: &otg.DeviceOspfv2Router{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceOspfv2Router) msg() *otg.DeviceOspfv2Router {
	return obj.obj
}

func (obj *deviceOspfv2Router) setMsg(msg *otg.DeviceOspfv2Router) DeviceOspfv2Router {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceOspfv2Router struct {
	obj *deviceOspfv2Router
}

type marshalDeviceOspfv2Router interface {
	// ToProto marshals DeviceOspfv2Router to protobuf object *otg.DeviceOspfv2Router
	ToProto() (*otg.DeviceOspfv2Router, error)
	// ToPbText marshals DeviceOspfv2Router to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceOspfv2Router to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceOspfv2Router to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceOspfv2Router struct {
	obj *deviceOspfv2Router
}

type unMarshalDeviceOspfv2Router interface {
	// FromProto unmarshals DeviceOspfv2Router from protobuf object *otg.DeviceOspfv2Router
	FromProto(msg *otg.DeviceOspfv2Router) (DeviceOspfv2Router, error)
	// FromPbText unmarshals DeviceOspfv2Router from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceOspfv2Router from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceOspfv2Router from JSON text
	FromJson(value string) error
}

func (obj *deviceOspfv2Router) Marshal() marshalDeviceOspfv2Router {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceOspfv2Router{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceOspfv2Router) Unmarshal() unMarshalDeviceOspfv2Router {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceOspfv2Router{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceOspfv2Router) ToProto() (*otg.DeviceOspfv2Router, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceOspfv2Router) FromProto(msg *otg.DeviceOspfv2Router) (DeviceOspfv2Router, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceOspfv2Router) ToPbText() (string, error) {
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

func (m *unMarshaldeviceOspfv2Router) FromPbText(value string) error {
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

func (m *marshaldeviceOspfv2Router) ToYaml() (string, error) {
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

func (m *unMarshaldeviceOspfv2Router) FromYaml(value string) error {
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

func (m *marshaldeviceOspfv2Router) ToJson() (string, error) {
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

func (m *unMarshaldeviceOspfv2Router) FromJson(value string) error {
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

func (obj *deviceOspfv2Router) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceOspfv2Router) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceOspfv2Router) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceOspfv2Router) Clone() (DeviceOspfv2Router, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceOspfv2Router()
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

func (obj *deviceOspfv2Router) setNil() {
	obj.routerIdHolder = nil
	obj.gracefulRestartHolder = nil
	obj.capabilitiesHolder = nil
	obj.interfacesHolder = nil
	obj.v4RoutesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceOspfv2Router is under Review: OSPFv2 is currently under review for pending exploration on use cases.
//
// A container of properties for an OSPFv2 router and its interfaces & Route Ranges.
type DeviceOspfv2Router interface {
	Validation
	// msg marshals DeviceOspfv2Router to protobuf object *otg.DeviceOspfv2Router
	// and doesn't set defaults
	msg() *otg.DeviceOspfv2Router
	// setMsg unmarshals DeviceOspfv2Router from protobuf object *otg.DeviceOspfv2Router
	// and doesn't set defaults
	setMsg(*otg.DeviceOspfv2Router) DeviceOspfv2Router
	// provides marshal interface
	Marshal() marshalDeviceOspfv2Router
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceOspfv2Router
	// validate validates DeviceOspfv2Router
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceOspfv2Router, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in DeviceOspfv2Router.
	Name() string
	// SetName assigns string provided by user to DeviceOspfv2Router
	SetName(value string) DeviceOspfv2Router
	// RouterId returns Ospfv2RouterId, set in DeviceOspfv2Router.
	// Ospfv2RouterId is container for OSPFv2 Router ID configuration.
	RouterId() Ospfv2RouterId
	// SetRouterId assigns Ospfv2RouterId provided by user to DeviceOspfv2Router.
	// Ospfv2RouterId is container for OSPFv2 Router ID configuration.
	SetRouterId(value Ospfv2RouterId) DeviceOspfv2Router
	// HasRouterId checks if RouterId has been set in DeviceOspfv2Router
	HasRouterId() bool
	// LsaRetransmitTime returns uint32, set in DeviceOspfv2Router.
	LsaRetransmitTime() uint32
	// SetLsaRetransmitTime assigns uint32 provided by user to DeviceOspfv2Router
	SetLsaRetransmitTime(value uint32) DeviceOspfv2Router
	// HasLsaRetransmitTime checks if LsaRetransmitTime has been set in DeviceOspfv2Router
	HasLsaRetransmitTime() bool
	// LsaRefreshTime returns uint32, set in DeviceOspfv2Router.
	LsaRefreshTime() uint32
	// SetLsaRefreshTime assigns uint32 provided by user to DeviceOspfv2Router
	SetLsaRefreshTime(value uint32) DeviceOspfv2Router
	// HasLsaRefreshTime checks if LsaRefreshTime has been set in DeviceOspfv2Router
	HasLsaRefreshTime() bool
	// InterBurstLsuInterval returns uint32, set in DeviceOspfv2Router.
	InterBurstLsuInterval() uint32
	// SetInterBurstLsuInterval assigns uint32 provided by user to DeviceOspfv2Router
	SetInterBurstLsuInterval(value uint32) DeviceOspfv2Router
	// HasInterBurstLsuInterval checks if InterBurstLsuInterval has been set in DeviceOspfv2Router
	HasInterBurstLsuInterval() bool
	// MaxFloodLsuPerBurst returns uint32, set in DeviceOspfv2Router.
	MaxFloodLsuPerBurst() uint32
	// SetMaxFloodLsuPerBurst assigns uint32 provided by user to DeviceOspfv2Router
	SetMaxFloodLsuPerBurst(value uint32) DeviceOspfv2Router
	// HasMaxFloodLsuPerBurst checks if MaxFloodLsuPerBurst has been set in DeviceOspfv2Router
	HasMaxFloodLsuPerBurst() bool
	// GracefulRestart returns Ospfv2GracefulRestart, set in DeviceOspfv2Router.
	// Ospfv2GracefulRestart is container of properties of OSPFv2 Graceful Retstart.
	GracefulRestart() Ospfv2GracefulRestart
	// SetGracefulRestart assigns Ospfv2GracefulRestart provided by user to DeviceOspfv2Router.
	// Ospfv2GracefulRestart is container of properties of OSPFv2 Graceful Retstart.
	SetGracefulRestart(value Ospfv2GracefulRestart) DeviceOspfv2Router
	// HasGracefulRestart checks if GracefulRestart has been set in DeviceOspfv2Router
	HasGracefulRestart() bool
	// StoreLsa returns bool, set in DeviceOspfv2Router.
	StoreLsa() bool
	// SetStoreLsa assigns bool provided by user to DeviceOspfv2Router
	SetStoreLsa(value bool) DeviceOspfv2Router
	// HasStoreLsa checks if StoreLsa has been set in DeviceOspfv2Router
	HasStoreLsa() bool
	// Capabilities returns Ospfv2Options, set in DeviceOspfv2Router.
	// Ospfv2Options is the OSPFv2 Options field is present Database Description packets and all LSAs.
	// This enables OSPF routers to support (or not support) optional capabilities,
	// and to communicate their capability level to other OSPF routers.
	// When capabilities are exchanged in Database Description packets a
	// router can choose not to forward certain LSAs to a neighbor because
	// of its reduced functionality.
	// Reference: A.2 The Options field: https://www.rfc-editor.org/rfc/rfc2328#page-46.
	Capabilities() Ospfv2Options
	// SetCapabilities assigns Ospfv2Options provided by user to DeviceOspfv2Router.
	// Ospfv2Options is the OSPFv2 Options field is present Database Description packets and all LSAs.
	// This enables OSPF routers to support (or not support) optional capabilities,
	// and to communicate their capability level to other OSPF routers.
	// When capabilities are exchanged in Database Description packets a
	// router can choose not to forward certain LSAs to a neighbor because
	// of its reduced functionality.
	// Reference: A.2 The Options field: https://www.rfc-editor.org/rfc/rfc2328#page-46.
	SetCapabilities(value Ospfv2Options) DeviceOspfv2Router
	// HasCapabilities checks if Capabilities has been set in DeviceOspfv2Router
	HasCapabilities() bool
	// Interfaces returns DeviceOspfv2RouterOspfv2InterfaceIterIter, set in DeviceOspfv2Router
	Interfaces() DeviceOspfv2RouterOspfv2InterfaceIter
	// V4Routes returns DeviceOspfv2RouterOspfv2V4RouteRangeIterIter, set in DeviceOspfv2Router
	V4Routes() DeviceOspfv2RouterOspfv2V4RouteRangeIter
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *deviceOspfv2Router) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the DeviceOspfv2Router object
func (obj *deviceOspfv2Router) SetName(value string) DeviceOspfv2Router {

	obj.obj.Name = &value
	return obj
}

// OSPFv2 Router Id.
// RouterId returns a Ospfv2RouterId
func (obj *deviceOspfv2Router) RouterId() Ospfv2RouterId {
	if obj.obj.RouterId == nil {
		obj.obj.RouterId = NewOspfv2RouterId().msg()
	}
	if obj.routerIdHolder == nil {
		obj.routerIdHolder = &ospfv2RouterId{obj: obj.obj.RouterId}
	}
	return obj.routerIdHolder
}

// OSPFv2 Router Id.
// RouterId returns a Ospfv2RouterId
func (obj *deviceOspfv2Router) HasRouterId() bool {
	return obj.obj.RouterId != nil
}

// OSPFv2 Router Id.
// SetRouterId sets the Ospfv2RouterId value in the DeviceOspfv2Router object
func (obj *deviceOspfv2Router) SetRouterId(value Ospfv2RouterId) DeviceOspfv2Router {

	obj.routerIdHolder = nil
	obj.obj.RouterId = value.msg()

	return obj
}

// The time in seconds for LSA retransmission.
// LsaRetransmitTime returns a uint32
func (obj *deviceOspfv2Router) LsaRetransmitTime() uint32 {

	return *obj.obj.LsaRetransmitTime

}

// The time in seconds for LSA retransmission.
// LsaRetransmitTime returns a uint32
func (obj *deviceOspfv2Router) HasLsaRetransmitTime() bool {
	return obj.obj.LsaRetransmitTime != nil
}

// The time in seconds for LSA retransmission.
// SetLsaRetransmitTime sets the uint32 value in the DeviceOspfv2Router object
func (obj *deviceOspfv2Router) SetLsaRetransmitTime(value uint32) DeviceOspfv2Router {

	obj.obj.LsaRetransmitTime = &value
	return obj
}

// The time in seconds required for LSA refresh.
// LsaRefreshTime returns a uint32
func (obj *deviceOspfv2Router) LsaRefreshTime() uint32 {

	return *obj.obj.LsaRefreshTime

}

// The time in seconds required for LSA refresh.
// LsaRefreshTime returns a uint32
func (obj *deviceOspfv2Router) HasLsaRefreshTime() bool {
	return obj.obj.LsaRefreshTime != nil
}

// The time in seconds required for LSA refresh.
// SetLsaRefreshTime sets the uint32 value in the DeviceOspfv2Router object
func (obj *deviceOspfv2Router) SetLsaRefreshTime(value uint32) DeviceOspfv2Router {

	obj.obj.LsaRefreshTime = &value
	return obj
}

// The gap in miliseconds between each Flood Link State Update Burst
// InterBurstLsuInterval returns a uint32
func (obj *deviceOspfv2Router) InterBurstLsuInterval() uint32 {

	return *obj.obj.InterBurstLsuInterval

}

// The gap in miliseconds between each Flood Link State Update Burst
// InterBurstLsuInterval returns a uint32
func (obj *deviceOspfv2Router) HasInterBurstLsuInterval() bool {
	return obj.obj.InterBurstLsuInterval != nil
}

// The gap in miliseconds between each Flood Link State Update Burst
// SetInterBurstLsuInterval sets the uint32 value in the DeviceOspfv2Router object
func (obj *deviceOspfv2Router) SetInterBurstLsuInterval(value uint32) DeviceOspfv2Router {

	obj.obj.InterBurstLsuInterval = &value
	return obj
}

// The maximum number of Flood LSUpdates for each burst
// MaxFloodLsuPerBurst returns a uint32
func (obj *deviceOspfv2Router) MaxFloodLsuPerBurst() uint32 {

	return *obj.obj.MaxFloodLsuPerBurst

}

// The maximum number of Flood LSUpdates for each burst
// MaxFloodLsuPerBurst returns a uint32
func (obj *deviceOspfv2Router) HasMaxFloodLsuPerBurst() bool {
	return obj.obj.MaxFloodLsuPerBurst != nil
}

// The maximum number of Flood LSUpdates for each burst
// SetMaxFloodLsuPerBurst sets the uint32 value in the DeviceOspfv2Router object
func (obj *deviceOspfv2Router) SetMaxFloodLsuPerBurst(value uint32) DeviceOspfv2Router {

	obj.obj.MaxFloodLsuPerBurst = &value
	return obj
}

// description is TBD
// GracefulRestart returns a Ospfv2GracefulRestart
func (obj *deviceOspfv2Router) GracefulRestart() Ospfv2GracefulRestart {
	if obj.obj.GracefulRestart == nil {
		obj.obj.GracefulRestart = NewOspfv2GracefulRestart().msg()
	}
	if obj.gracefulRestartHolder == nil {
		obj.gracefulRestartHolder = &ospfv2GracefulRestart{obj: obj.obj.GracefulRestart}
	}
	return obj.gracefulRestartHolder
}

// description is TBD
// GracefulRestart returns a Ospfv2GracefulRestart
func (obj *deviceOspfv2Router) HasGracefulRestart() bool {
	return obj.obj.GracefulRestart != nil
}

// description is TBD
// SetGracefulRestart sets the Ospfv2GracefulRestart value in the DeviceOspfv2Router object
func (obj *deviceOspfv2Router) SetGracefulRestart(value Ospfv2GracefulRestart) DeviceOspfv2Router {

	obj.gracefulRestartHolder = nil
	obj.obj.GracefulRestart = value.msg()

	return obj
}

// Configuration for controlling storage of OSPFv2 learned LSAs received from the neighbors.
// StoreLsa returns a bool
func (obj *deviceOspfv2Router) StoreLsa() bool {

	return *obj.obj.StoreLsa

}

// Configuration for controlling storage of OSPFv2 learned LSAs received from the neighbors.
// StoreLsa returns a bool
func (obj *deviceOspfv2Router) HasStoreLsa() bool {
	return obj.obj.StoreLsa != nil
}

// Configuration for controlling storage of OSPFv2 learned LSAs received from the neighbors.
// SetStoreLsa sets the bool value in the DeviceOspfv2Router object
func (obj *deviceOspfv2Router) SetStoreLsa(value bool) DeviceOspfv2Router {

	obj.obj.StoreLsa = &value
	return obj
}

// A router indicates the optional capabilities that it supports in its OSPF Hello packets, Database Description packets and in its LSAs.
// Capabilities returns a Ospfv2Options
func (obj *deviceOspfv2Router) Capabilities() Ospfv2Options {
	if obj.obj.Capabilities == nil {
		obj.obj.Capabilities = NewOspfv2Options().msg()
	}
	if obj.capabilitiesHolder == nil {
		obj.capabilitiesHolder = &ospfv2Options{obj: obj.obj.Capabilities}
	}
	return obj.capabilitiesHolder
}

// A router indicates the optional capabilities that it supports in its OSPF Hello packets, Database Description packets and in its LSAs.
// Capabilities returns a Ospfv2Options
func (obj *deviceOspfv2Router) HasCapabilities() bool {
	return obj.obj.Capabilities != nil
}

// A router indicates the optional capabilities that it supports in its OSPF Hello packets, Database Description packets and in its LSAs.
// SetCapabilities sets the Ospfv2Options value in the DeviceOspfv2Router object
func (obj *deviceOspfv2Router) SetCapabilities(value Ospfv2Options) DeviceOspfv2Router {

	obj.capabilitiesHolder = nil
	obj.obj.Capabilities = value.msg()

	return obj
}

// List of OSPFv2 interfaces for this router.
// Interfaces returns a []Ospfv2Interface
func (obj *deviceOspfv2Router) Interfaces() DeviceOspfv2RouterOspfv2InterfaceIter {
	if len(obj.obj.Interfaces) == 0 {
		obj.obj.Interfaces = []*otg.Ospfv2Interface{}
	}
	if obj.interfacesHolder == nil {
		obj.interfacesHolder = newDeviceOspfv2RouterOspfv2InterfaceIter(&obj.obj.Interfaces).setMsg(obj)
	}
	return obj.interfacesHolder
}

type deviceOspfv2RouterOspfv2InterfaceIter struct {
	obj                  *deviceOspfv2Router
	ospfv2InterfaceSlice []Ospfv2Interface
	fieldPtr             *[]*otg.Ospfv2Interface
}

func newDeviceOspfv2RouterOspfv2InterfaceIter(ptr *[]*otg.Ospfv2Interface) DeviceOspfv2RouterOspfv2InterfaceIter {
	return &deviceOspfv2RouterOspfv2InterfaceIter{fieldPtr: ptr}
}

type DeviceOspfv2RouterOspfv2InterfaceIter interface {
	setMsg(*deviceOspfv2Router) DeviceOspfv2RouterOspfv2InterfaceIter
	Items() []Ospfv2Interface
	Add() Ospfv2Interface
	Append(items ...Ospfv2Interface) DeviceOspfv2RouterOspfv2InterfaceIter
	Set(index int, newObj Ospfv2Interface) DeviceOspfv2RouterOspfv2InterfaceIter
	Clear() DeviceOspfv2RouterOspfv2InterfaceIter
	clearHolderSlice() DeviceOspfv2RouterOspfv2InterfaceIter
	appendHolderSlice(item Ospfv2Interface) DeviceOspfv2RouterOspfv2InterfaceIter
}

func (obj *deviceOspfv2RouterOspfv2InterfaceIter) setMsg(msg *deviceOspfv2Router) DeviceOspfv2RouterOspfv2InterfaceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2Interface{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceOspfv2RouterOspfv2InterfaceIter) Items() []Ospfv2Interface {
	return obj.ospfv2InterfaceSlice
}

func (obj *deviceOspfv2RouterOspfv2InterfaceIter) Add() Ospfv2Interface {
	newObj := &otg.Ospfv2Interface{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2Interface{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2InterfaceSlice = append(obj.ospfv2InterfaceSlice, newLibObj)
	return newLibObj
}

func (obj *deviceOspfv2RouterOspfv2InterfaceIter) Append(items ...Ospfv2Interface) DeviceOspfv2RouterOspfv2InterfaceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2InterfaceSlice = append(obj.ospfv2InterfaceSlice, item)
	}
	return obj
}

func (obj *deviceOspfv2RouterOspfv2InterfaceIter) Set(index int, newObj Ospfv2Interface) DeviceOspfv2RouterOspfv2InterfaceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2InterfaceSlice[index] = newObj
	return obj
}
func (obj *deviceOspfv2RouterOspfv2InterfaceIter) Clear() DeviceOspfv2RouterOspfv2InterfaceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2Interface{}
		obj.ospfv2InterfaceSlice = []Ospfv2Interface{}
	}
	return obj
}
func (obj *deviceOspfv2RouterOspfv2InterfaceIter) clearHolderSlice() DeviceOspfv2RouterOspfv2InterfaceIter {
	if len(obj.ospfv2InterfaceSlice) > 0 {
		obj.ospfv2InterfaceSlice = []Ospfv2Interface{}
	}
	return obj
}
func (obj *deviceOspfv2RouterOspfv2InterfaceIter) appendHolderSlice(item Ospfv2Interface) DeviceOspfv2RouterOspfv2InterfaceIter {
	obj.ospfv2InterfaceSlice = append(obj.ospfv2InterfaceSlice, item)
	return obj
}

// Emulated OSPFv4 IPv4 routes.
// V4Routes returns a []Ospfv2V4RouteRange
func (obj *deviceOspfv2Router) V4Routes() DeviceOspfv2RouterOspfv2V4RouteRangeIter {
	if len(obj.obj.V4Routes) == 0 {
		obj.obj.V4Routes = []*otg.Ospfv2V4RouteRange{}
	}
	if obj.v4RoutesHolder == nil {
		obj.v4RoutesHolder = newDeviceOspfv2RouterOspfv2V4RouteRangeIter(&obj.obj.V4Routes).setMsg(obj)
	}
	return obj.v4RoutesHolder
}

type deviceOspfv2RouterOspfv2V4RouteRangeIter struct {
	obj                     *deviceOspfv2Router
	ospfv2V4RouteRangeSlice []Ospfv2V4RouteRange
	fieldPtr                *[]*otg.Ospfv2V4RouteRange
}

func newDeviceOspfv2RouterOspfv2V4RouteRangeIter(ptr *[]*otg.Ospfv2V4RouteRange) DeviceOspfv2RouterOspfv2V4RouteRangeIter {
	return &deviceOspfv2RouterOspfv2V4RouteRangeIter{fieldPtr: ptr}
}

type DeviceOspfv2RouterOspfv2V4RouteRangeIter interface {
	setMsg(*deviceOspfv2Router) DeviceOspfv2RouterOspfv2V4RouteRangeIter
	Items() []Ospfv2V4RouteRange
	Add() Ospfv2V4RouteRange
	Append(items ...Ospfv2V4RouteRange) DeviceOspfv2RouterOspfv2V4RouteRangeIter
	Set(index int, newObj Ospfv2V4RouteRange) DeviceOspfv2RouterOspfv2V4RouteRangeIter
	Clear() DeviceOspfv2RouterOspfv2V4RouteRangeIter
	clearHolderSlice() DeviceOspfv2RouterOspfv2V4RouteRangeIter
	appendHolderSlice(item Ospfv2V4RouteRange) DeviceOspfv2RouterOspfv2V4RouteRangeIter
}

func (obj *deviceOspfv2RouterOspfv2V4RouteRangeIter) setMsg(msg *deviceOspfv2Router) DeviceOspfv2RouterOspfv2V4RouteRangeIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2V4RouteRange{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceOspfv2RouterOspfv2V4RouteRangeIter) Items() []Ospfv2V4RouteRange {
	return obj.ospfv2V4RouteRangeSlice
}

func (obj *deviceOspfv2RouterOspfv2V4RouteRangeIter) Add() Ospfv2V4RouteRange {
	newObj := &otg.Ospfv2V4RouteRange{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2V4RouteRange{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2V4RouteRangeSlice = append(obj.ospfv2V4RouteRangeSlice, newLibObj)
	return newLibObj
}

func (obj *deviceOspfv2RouterOspfv2V4RouteRangeIter) Append(items ...Ospfv2V4RouteRange) DeviceOspfv2RouterOspfv2V4RouteRangeIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2V4RouteRangeSlice = append(obj.ospfv2V4RouteRangeSlice, item)
	}
	return obj
}

func (obj *deviceOspfv2RouterOspfv2V4RouteRangeIter) Set(index int, newObj Ospfv2V4RouteRange) DeviceOspfv2RouterOspfv2V4RouteRangeIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2V4RouteRangeSlice[index] = newObj
	return obj
}
func (obj *deviceOspfv2RouterOspfv2V4RouteRangeIter) Clear() DeviceOspfv2RouterOspfv2V4RouteRangeIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2V4RouteRange{}
		obj.ospfv2V4RouteRangeSlice = []Ospfv2V4RouteRange{}
	}
	return obj
}
func (obj *deviceOspfv2RouterOspfv2V4RouteRangeIter) clearHolderSlice() DeviceOspfv2RouterOspfv2V4RouteRangeIter {
	if len(obj.ospfv2V4RouteRangeSlice) > 0 {
		obj.ospfv2V4RouteRangeSlice = []Ospfv2V4RouteRange{}
	}
	return obj
}
func (obj *deviceOspfv2RouterOspfv2V4RouteRangeIter) appendHolderSlice(item Ospfv2V4RouteRange) DeviceOspfv2RouterOspfv2V4RouteRangeIter {
	obj.ospfv2V4RouteRangeSlice = append(obj.ospfv2V4RouteRangeSlice, item)
	return obj
}

func (obj *deviceOspfv2Router) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	obj.addWarnings("DeviceOspfv2Router is under review, OSPFv2 is currently under review for pending exploration on use cases.")

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface DeviceOspfv2Router")
	}

	if obj.obj.RouterId != nil {

		obj.RouterId().validateObj(vObj, set_default)
	}

	if obj.obj.LsaRetransmitTime != nil {

		if *obj.obj.LsaRetransmitTime < 1 || *obj.obj.LsaRetransmitTime > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= DeviceOspfv2Router.LsaRetransmitTime <= 4294967295 but Got %d", *obj.obj.LsaRetransmitTime))
		}

	}

	if obj.obj.LsaRefreshTime != nil {

		if *obj.obj.LsaRefreshTime < 5 || *obj.obj.LsaRefreshTime > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("5 <= DeviceOspfv2Router.LsaRefreshTime <= 4294967295 but Got %d", *obj.obj.LsaRefreshTime))
		}

	}

	if obj.obj.MaxFloodLsuPerBurst != nil {

		if *obj.obj.MaxFloodLsuPerBurst < 1 || *obj.obj.MaxFloodLsuPerBurst > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= DeviceOspfv2Router.MaxFloodLsuPerBurst <= 4294967295 but Got %d", *obj.obj.MaxFloodLsuPerBurst))
		}

	}

	if obj.obj.GracefulRestart != nil {

		obj.GracefulRestart().validateObj(vObj, set_default)
	}

	if obj.obj.Capabilities != nil {

		obj.Capabilities().validateObj(vObj, set_default)
	}

	if len(obj.obj.Interfaces) != 0 {

		if set_default {
			obj.Interfaces().clearHolderSlice()
			for _, item := range obj.obj.Interfaces {
				obj.Interfaces().appendHolderSlice(&ospfv2Interface{obj: item})
			}
		}
		for _, item := range obj.Interfaces().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.V4Routes) != 0 {

		if set_default {
			obj.V4Routes().clearHolderSlice()
			for _, item := range obj.obj.V4Routes {
				obj.V4Routes().appendHolderSlice(&ospfv2V4RouteRange{obj: item})
			}
		}
		for _, item := range obj.V4Routes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *deviceOspfv2Router) setDefault() {
	if obj.obj.LsaRetransmitTime == nil {
		obj.SetLsaRetransmitTime(5)
	}
	if obj.obj.LsaRefreshTime == nil {
		obj.SetLsaRefreshTime(1800)
	}
	if obj.obj.InterBurstLsuInterval == nil {
		obj.SetInterBurstLsuInterval(33)
	}
	if obj.obj.MaxFloodLsuPerBurst == nil {
		obj.SetMaxFloodLsuPerBurst(1)
	}
	if obj.obj.StoreLsa == nil {
		obj.SetStoreLsa(false)
	}

}

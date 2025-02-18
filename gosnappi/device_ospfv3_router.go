package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceOspfv3Router *****
type deviceOspfv3Router struct {
	validation
	obj                   *otg.DeviceOspfv3Router
	marshaller            marshalDeviceOspfv3Router
	unMarshaller          unMarshalDeviceOspfv3Router
	routerIdHolder        Ospfv3RouterId
	gracefulRestartHolder Ospfv3GracefulRestart
	capabilitiesHolder    Ospfv3Capabilities
	interfacesHolder      DeviceOspfv3RouterOspfv3InterfaceIter
	v6RoutesHolder        DeviceOspfv3RouterOspfv3V6RouteRangeIter
}

func NewDeviceOspfv3Router() DeviceOspfv3Router {
	obj := deviceOspfv3Router{obj: &otg.DeviceOspfv3Router{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceOspfv3Router) msg() *otg.DeviceOspfv3Router {
	return obj.obj
}

func (obj *deviceOspfv3Router) setMsg(msg *otg.DeviceOspfv3Router) DeviceOspfv3Router {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceOspfv3Router struct {
	obj *deviceOspfv3Router
}

type marshalDeviceOspfv3Router interface {
	// ToProto marshals DeviceOspfv3Router to protobuf object *otg.DeviceOspfv3Router
	ToProto() (*otg.DeviceOspfv3Router, error)
	// ToPbText marshals DeviceOspfv3Router to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceOspfv3Router to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceOspfv3Router to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceOspfv3Router struct {
	obj *deviceOspfv3Router
}

type unMarshalDeviceOspfv3Router interface {
	// FromProto unmarshals DeviceOspfv3Router from protobuf object *otg.DeviceOspfv3Router
	FromProto(msg *otg.DeviceOspfv3Router) (DeviceOspfv3Router, error)
	// FromPbText unmarshals DeviceOspfv3Router from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceOspfv3Router from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceOspfv3Router from JSON text
	FromJson(value string) error
}

func (obj *deviceOspfv3Router) Marshal() marshalDeviceOspfv3Router {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceOspfv3Router{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceOspfv3Router) Unmarshal() unMarshalDeviceOspfv3Router {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceOspfv3Router{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceOspfv3Router) ToProto() (*otg.DeviceOspfv3Router, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceOspfv3Router) FromProto(msg *otg.DeviceOspfv3Router) (DeviceOspfv3Router, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceOspfv3Router) ToPbText() (string, error) {
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

func (m *unMarshaldeviceOspfv3Router) FromPbText(value string) error {
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

func (m *marshaldeviceOspfv3Router) ToYaml() (string, error) {
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

func (m *unMarshaldeviceOspfv3Router) FromYaml(value string) error {
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

func (m *marshaldeviceOspfv3Router) ToJson() (string, error) {
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

func (m *unMarshaldeviceOspfv3Router) FromJson(value string) error {
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

func (obj *deviceOspfv3Router) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceOspfv3Router) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceOspfv3Router) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceOspfv3Router) Clone() (DeviceOspfv3Router, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceOspfv3Router()
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

func (obj *deviceOspfv3Router) setNil() {
	obj.routerIdHolder = nil
	obj.gracefulRestartHolder = nil
	obj.capabilitiesHolder = nil
	obj.interfacesHolder = nil
	obj.v6RoutesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceOspfv3Router is under Review: OSPFv3 is currently under review for pending exploration on use cases.
//
// A container of properties for an OSPFv3 router and its interfaces & Route Ranges.
type DeviceOspfv3Router interface {
	Validation
	// msg marshals DeviceOspfv3Router to protobuf object *otg.DeviceOspfv3Router
	// and doesn't set defaults
	msg() *otg.DeviceOspfv3Router
	// setMsg unmarshals DeviceOspfv3Router from protobuf object *otg.DeviceOspfv3Router
	// and doesn't set defaults
	setMsg(*otg.DeviceOspfv3Router) DeviceOspfv3Router
	// provides marshal interface
	Marshal() marshalDeviceOspfv3Router
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceOspfv3Router
	// validate validates DeviceOspfv3Router
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceOspfv3Router, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in DeviceOspfv3Router.
	Name() string
	// SetName assigns string provided by user to DeviceOspfv3Router
	SetName(value string) DeviceOspfv3Router
	// RouterId returns Ospfv3RouterId, set in DeviceOspfv3Router.
	// Ospfv3RouterId is container for OSPFv3 Router ID configuration.
	RouterId() Ospfv3RouterId
	// SetRouterId assigns Ospfv3RouterId provided by user to DeviceOspfv3Router.
	// Ospfv3RouterId is container for OSPFv3 Router ID configuration.
	SetRouterId(value Ospfv3RouterId) DeviceOspfv3Router
	// HasRouterId checks if RouterId has been set in DeviceOspfv3Router
	HasRouterId() bool
	// LsaRetransmitTime returns uint32, set in DeviceOspfv3Router.
	LsaRetransmitTime() uint32
	// SetLsaRetransmitTime assigns uint32 provided by user to DeviceOspfv3Router
	SetLsaRetransmitTime(value uint32) DeviceOspfv3Router
	// HasLsaRetransmitTime checks if LsaRetransmitTime has been set in DeviceOspfv3Router
	HasLsaRetransmitTime() bool
	// LsaRefreshTime returns uint32, set in DeviceOspfv3Router.
	LsaRefreshTime() uint32
	// SetLsaRefreshTime assigns uint32 provided by user to DeviceOspfv3Router
	SetLsaRefreshTime(value uint32) DeviceOspfv3Router
	// HasLsaRefreshTime checks if LsaRefreshTime has been set in DeviceOspfv3Router
	HasLsaRefreshTime() bool
	// InterBurstLsuInterval returns uint32, set in DeviceOspfv3Router.
	InterBurstLsuInterval() uint32
	// SetInterBurstLsuInterval assigns uint32 provided by user to DeviceOspfv3Router
	SetInterBurstLsuInterval(value uint32) DeviceOspfv3Router
	// HasInterBurstLsuInterval checks if InterBurstLsuInterval has been set in DeviceOspfv3Router
	HasInterBurstLsuInterval() bool
	// GracefulRestart returns Ospfv3GracefulRestart, set in DeviceOspfv3Router.
	// Ospfv3GracefulRestart is container of properties of OSPFv3 Graceful Retstart.
	GracefulRestart() Ospfv3GracefulRestart
	// SetGracefulRestart assigns Ospfv3GracefulRestart provided by user to DeviceOspfv3Router.
	// Ospfv3GracefulRestart is container of properties of OSPFv3 Graceful Retstart.
	SetGracefulRestart(value Ospfv3GracefulRestart) DeviceOspfv3Router
	// HasGracefulRestart checks if GracefulRestart has been set in DeviceOspfv3Router
	HasGracefulRestart() bool
	// StoreLsa returns bool, set in DeviceOspfv3Router.
	StoreLsa() bool
	// SetStoreLsa assigns bool provided by user to DeviceOspfv3Router
	SetStoreLsa(value bool) DeviceOspfv3Router
	// HasStoreLsa checks if StoreLsa has been set in DeviceOspfv3Router
	HasStoreLsa() bool
	// Capabilities returns Ospfv3Capabilities, set in DeviceOspfv3Router.
	// Ospfv3Capabilities is the OSPFv3 router optional attributes.
	Capabilities() Ospfv3Capabilities
	// SetCapabilities assigns Ospfv3Capabilities provided by user to DeviceOspfv3Router.
	// Ospfv3Capabilities is the OSPFv3 router optional attributes.
	SetCapabilities(value Ospfv3Capabilities) DeviceOspfv3Router
	// HasCapabilities checks if Capabilities has been set in DeviceOspfv3Router
	HasCapabilities() bool
	// Interfaces returns DeviceOspfv3RouterOspfv3InterfaceIterIter, set in DeviceOspfv3Router
	Interfaces() DeviceOspfv3RouterOspfv3InterfaceIter
	// V6Routes returns DeviceOspfv3RouterOspfv3V6RouteRangeIterIter, set in DeviceOspfv3Router
	V6Routes() DeviceOspfv3RouterOspfv3V6RouteRangeIter
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *deviceOspfv3Router) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the DeviceOspfv3Router object
func (obj *deviceOspfv3Router) SetName(value string) DeviceOspfv3Router {

	obj.obj.Name = &value
	return obj
}

// OSPFv3 Router Id.
// RouterId returns a Ospfv3RouterId
func (obj *deviceOspfv3Router) RouterId() Ospfv3RouterId {
	if obj.obj.RouterId == nil {
		obj.obj.RouterId = NewOspfv3RouterId().msg()
	}
	if obj.routerIdHolder == nil {
		obj.routerIdHolder = &ospfv3RouterId{obj: obj.obj.RouterId}
	}
	return obj.routerIdHolder
}

// OSPFv3 Router Id.
// RouterId returns a Ospfv3RouterId
func (obj *deviceOspfv3Router) HasRouterId() bool {
	return obj.obj.RouterId != nil
}

// OSPFv3 Router Id.
// SetRouterId sets the Ospfv3RouterId value in the DeviceOspfv3Router object
func (obj *deviceOspfv3Router) SetRouterId(value Ospfv3RouterId) DeviceOspfv3Router {

	obj.routerIdHolder = nil
	obj.obj.RouterId = value.msg()

	return obj
}

// The time in seconds for LSA retransmission.
// LsaRetransmitTime returns a uint32
func (obj *deviceOspfv3Router) LsaRetransmitTime() uint32 {

	return *obj.obj.LsaRetransmitTime

}

// The time in seconds for LSA retransmission.
// LsaRetransmitTime returns a uint32
func (obj *deviceOspfv3Router) HasLsaRetransmitTime() bool {
	return obj.obj.LsaRetransmitTime != nil
}

// The time in seconds for LSA retransmission.
// SetLsaRetransmitTime sets the uint32 value in the DeviceOspfv3Router object
func (obj *deviceOspfv3Router) SetLsaRetransmitTime(value uint32) DeviceOspfv3Router {

	obj.obj.LsaRetransmitTime = &value
	return obj
}

// The time in seconds required for LSA refresh.
// LsaRefreshTime returns a uint32
func (obj *deviceOspfv3Router) LsaRefreshTime() uint32 {

	return *obj.obj.LsaRefreshTime

}

// The time in seconds required for LSA refresh.
// LsaRefreshTime returns a uint32
func (obj *deviceOspfv3Router) HasLsaRefreshTime() bool {
	return obj.obj.LsaRefreshTime != nil
}

// The time in seconds required for LSA refresh.
// SetLsaRefreshTime sets the uint32 value in the DeviceOspfv3Router object
func (obj *deviceOspfv3Router) SetLsaRefreshTime(value uint32) DeviceOspfv3Router {

	obj.obj.LsaRefreshTime = &value
	return obj
}

// The gap in miliseconds between each Flood Link State Update burst.
// InterBurstLsuInterval returns a uint32
func (obj *deviceOspfv3Router) InterBurstLsuInterval() uint32 {

	return *obj.obj.InterBurstLsuInterval

}

// The gap in miliseconds between each Flood Link State Update burst.
// InterBurstLsuInterval returns a uint32
func (obj *deviceOspfv3Router) HasInterBurstLsuInterval() bool {
	return obj.obj.InterBurstLsuInterval != nil
}

// The gap in miliseconds between each Flood Link State Update burst.
// SetInterBurstLsuInterval sets the uint32 value in the DeviceOspfv3Router object
func (obj *deviceOspfv3Router) SetInterBurstLsuInterval(value uint32) DeviceOspfv3Router {

	obj.obj.InterBurstLsuInterval = &value
	return obj
}

// description is TBD
// GracefulRestart returns a Ospfv3GracefulRestart
func (obj *deviceOspfv3Router) GracefulRestart() Ospfv3GracefulRestart {
	if obj.obj.GracefulRestart == nil {
		obj.obj.GracefulRestart = NewOspfv3GracefulRestart().msg()
	}
	if obj.gracefulRestartHolder == nil {
		obj.gracefulRestartHolder = &ospfv3GracefulRestart{obj: obj.obj.GracefulRestart}
	}
	return obj.gracefulRestartHolder
}

// description is TBD
// GracefulRestart returns a Ospfv3GracefulRestart
func (obj *deviceOspfv3Router) HasGracefulRestart() bool {
	return obj.obj.GracefulRestart != nil
}

// description is TBD
// SetGracefulRestart sets the Ospfv3GracefulRestart value in the DeviceOspfv3Router object
func (obj *deviceOspfv3Router) SetGracefulRestart(value Ospfv3GracefulRestart) DeviceOspfv3Router {

	obj.gracefulRestartHolder = nil
	obj.obj.GracefulRestart = value.msg()

	return obj
}

// Configuration for controlling storage of OSPFv3 learned LSAs received from the neighbors.
// StoreLsa returns a bool
func (obj *deviceOspfv3Router) StoreLsa() bool {

	return *obj.obj.StoreLsa

}

// Configuration for controlling storage of OSPFv3 learned LSAs received from the neighbors.
// StoreLsa returns a bool
func (obj *deviceOspfv3Router) HasStoreLsa() bool {
	return obj.obj.StoreLsa != nil
}

// Configuration for controlling storage of OSPFv3 learned LSAs received from the neighbors.
// SetStoreLsa sets the bool value in the DeviceOspfv3Router object
func (obj *deviceOspfv3Router) SetStoreLsa(value bool) DeviceOspfv3Router {

	obj.obj.StoreLsa = &value
	return obj
}

// Optional container for OSPFv3 router capabilities.
// Capabilities returns a Ospfv3Capabilities
func (obj *deviceOspfv3Router) Capabilities() Ospfv3Capabilities {
	if obj.obj.Capabilities == nil {
		obj.obj.Capabilities = NewOspfv3Capabilities().msg()
	}
	if obj.capabilitiesHolder == nil {
		obj.capabilitiesHolder = &ospfv3Capabilities{obj: obj.obj.Capabilities}
	}
	return obj.capabilitiesHolder
}

// Optional container for OSPFv3 router capabilities.
// Capabilities returns a Ospfv3Capabilities
func (obj *deviceOspfv3Router) HasCapabilities() bool {
	return obj.obj.Capabilities != nil
}

// Optional container for OSPFv3 router capabilities.
// SetCapabilities sets the Ospfv3Capabilities value in the DeviceOspfv3Router object
func (obj *deviceOspfv3Router) SetCapabilities(value Ospfv3Capabilities) DeviceOspfv3Router {

	obj.capabilitiesHolder = nil
	obj.obj.Capabilities = value.msg()

	return obj
}

// List of OSPFv3 interfaces for this router.
// Interfaces returns a []Ospfv3Interface
func (obj *deviceOspfv3Router) Interfaces() DeviceOspfv3RouterOspfv3InterfaceIter {
	if len(obj.obj.Interfaces) == 0 {
		obj.obj.Interfaces = []*otg.Ospfv3Interface{}
	}
	if obj.interfacesHolder == nil {
		obj.interfacesHolder = newDeviceOspfv3RouterOspfv3InterfaceIter(&obj.obj.Interfaces).setMsg(obj)
	}
	return obj.interfacesHolder
}

type deviceOspfv3RouterOspfv3InterfaceIter struct {
	obj                  *deviceOspfv3Router
	ospfv3InterfaceSlice []Ospfv3Interface
	fieldPtr             *[]*otg.Ospfv3Interface
}

func newDeviceOspfv3RouterOspfv3InterfaceIter(ptr *[]*otg.Ospfv3Interface) DeviceOspfv3RouterOspfv3InterfaceIter {
	return &deviceOspfv3RouterOspfv3InterfaceIter{fieldPtr: ptr}
}

type DeviceOspfv3RouterOspfv3InterfaceIter interface {
	setMsg(*deviceOspfv3Router) DeviceOspfv3RouterOspfv3InterfaceIter
	Items() []Ospfv3Interface
	Add() Ospfv3Interface
	Append(items ...Ospfv3Interface) DeviceOspfv3RouterOspfv3InterfaceIter
	Set(index int, newObj Ospfv3Interface) DeviceOspfv3RouterOspfv3InterfaceIter
	Clear() DeviceOspfv3RouterOspfv3InterfaceIter
	clearHolderSlice() DeviceOspfv3RouterOspfv3InterfaceIter
	appendHolderSlice(item Ospfv3Interface) DeviceOspfv3RouterOspfv3InterfaceIter
}

func (obj *deviceOspfv3RouterOspfv3InterfaceIter) setMsg(msg *deviceOspfv3Router) DeviceOspfv3RouterOspfv3InterfaceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv3Interface{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceOspfv3RouterOspfv3InterfaceIter) Items() []Ospfv3Interface {
	return obj.ospfv3InterfaceSlice
}

func (obj *deviceOspfv3RouterOspfv3InterfaceIter) Add() Ospfv3Interface {
	newObj := &otg.Ospfv3Interface{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv3Interface{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv3InterfaceSlice = append(obj.ospfv3InterfaceSlice, newLibObj)
	return newLibObj
}

func (obj *deviceOspfv3RouterOspfv3InterfaceIter) Append(items ...Ospfv3Interface) DeviceOspfv3RouterOspfv3InterfaceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv3InterfaceSlice = append(obj.ospfv3InterfaceSlice, item)
	}
	return obj
}

func (obj *deviceOspfv3RouterOspfv3InterfaceIter) Set(index int, newObj Ospfv3Interface) DeviceOspfv3RouterOspfv3InterfaceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv3InterfaceSlice[index] = newObj
	return obj
}
func (obj *deviceOspfv3RouterOspfv3InterfaceIter) Clear() DeviceOspfv3RouterOspfv3InterfaceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv3Interface{}
		obj.ospfv3InterfaceSlice = []Ospfv3Interface{}
	}
	return obj
}
func (obj *deviceOspfv3RouterOspfv3InterfaceIter) clearHolderSlice() DeviceOspfv3RouterOspfv3InterfaceIter {
	if len(obj.ospfv3InterfaceSlice) > 0 {
		obj.ospfv3InterfaceSlice = []Ospfv3Interface{}
	}
	return obj
}
func (obj *deviceOspfv3RouterOspfv3InterfaceIter) appendHolderSlice(item Ospfv3Interface) DeviceOspfv3RouterOspfv3InterfaceIter {
	obj.ospfv3InterfaceSlice = append(obj.ospfv3InterfaceSlice, item)
	return obj
}

// Emulated OSPFv4 IPv6 routes.
// V6Routes returns a []Ospfv3V6RouteRange
func (obj *deviceOspfv3Router) V6Routes() DeviceOspfv3RouterOspfv3V6RouteRangeIter {
	if len(obj.obj.V6Routes) == 0 {
		obj.obj.V6Routes = []*otg.Ospfv3V6RouteRange{}
	}
	if obj.v6RoutesHolder == nil {
		obj.v6RoutesHolder = newDeviceOspfv3RouterOspfv3V6RouteRangeIter(&obj.obj.V6Routes).setMsg(obj)
	}
	return obj.v6RoutesHolder
}

type deviceOspfv3RouterOspfv3V6RouteRangeIter struct {
	obj                     *deviceOspfv3Router
	ospfv3V6RouteRangeSlice []Ospfv3V6RouteRange
	fieldPtr                *[]*otg.Ospfv3V6RouteRange
}

func newDeviceOspfv3RouterOspfv3V6RouteRangeIter(ptr *[]*otg.Ospfv3V6RouteRange) DeviceOspfv3RouterOspfv3V6RouteRangeIter {
	return &deviceOspfv3RouterOspfv3V6RouteRangeIter{fieldPtr: ptr}
}

type DeviceOspfv3RouterOspfv3V6RouteRangeIter interface {
	setMsg(*deviceOspfv3Router) DeviceOspfv3RouterOspfv3V6RouteRangeIter
	Items() []Ospfv3V6RouteRange
	Add() Ospfv3V6RouteRange
	Append(items ...Ospfv3V6RouteRange) DeviceOspfv3RouterOspfv3V6RouteRangeIter
	Set(index int, newObj Ospfv3V6RouteRange) DeviceOspfv3RouterOspfv3V6RouteRangeIter
	Clear() DeviceOspfv3RouterOspfv3V6RouteRangeIter
	clearHolderSlice() DeviceOspfv3RouterOspfv3V6RouteRangeIter
	appendHolderSlice(item Ospfv3V6RouteRange) DeviceOspfv3RouterOspfv3V6RouteRangeIter
}

func (obj *deviceOspfv3RouterOspfv3V6RouteRangeIter) setMsg(msg *deviceOspfv3Router) DeviceOspfv3RouterOspfv3V6RouteRangeIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv3V6RouteRange{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceOspfv3RouterOspfv3V6RouteRangeIter) Items() []Ospfv3V6RouteRange {
	return obj.ospfv3V6RouteRangeSlice
}

func (obj *deviceOspfv3RouterOspfv3V6RouteRangeIter) Add() Ospfv3V6RouteRange {
	newObj := &otg.Ospfv3V6RouteRange{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv3V6RouteRange{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv3V6RouteRangeSlice = append(obj.ospfv3V6RouteRangeSlice, newLibObj)
	return newLibObj
}

func (obj *deviceOspfv3RouterOspfv3V6RouteRangeIter) Append(items ...Ospfv3V6RouteRange) DeviceOspfv3RouterOspfv3V6RouteRangeIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv3V6RouteRangeSlice = append(obj.ospfv3V6RouteRangeSlice, item)
	}
	return obj
}

func (obj *deviceOspfv3RouterOspfv3V6RouteRangeIter) Set(index int, newObj Ospfv3V6RouteRange) DeviceOspfv3RouterOspfv3V6RouteRangeIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv3V6RouteRangeSlice[index] = newObj
	return obj
}
func (obj *deviceOspfv3RouterOspfv3V6RouteRangeIter) Clear() DeviceOspfv3RouterOspfv3V6RouteRangeIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv3V6RouteRange{}
		obj.ospfv3V6RouteRangeSlice = []Ospfv3V6RouteRange{}
	}
	return obj
}
func (obj *deviceOspfv3RouterOspfv3V6RouteRangeIter) clearHolderSlice() DeviceOspfv3RouterOspfv3V6RouteRangeIter {
	if len(obj.ospfv3V6RouteRangeSlice) > 0 {
		obj.ospfv3V6RouteRangeSlice = []Ospfv3V6RouteRange{}
	}
	return obj
}
func (obj *deviceOspfv3RouterOspfv3V6RouteRangeIter) appendHolderSlice(item Ospfv3V6RouteRange) DeviceOspfv3RouterOspfv3V6RouteRangeIter {
	obj.ospfv3V6RouteRangeSlice = append(obj.ospfv3V6RouteRangeSlice, item)
	return obj
}

func (obj *deviceOspfv3Router) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	obj.addWarnings("DeviceOspfv3Router is under review, OSPFv3 is currently under review for pending exploration on use cases.")

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface DeviceOspfv3Router")
	}

	if obj.obj.RouterId != nil {

		obj.RouterId().validateObj(vObj, set_default)
	}

	if obj.obj.LsaRetransmitTime != nil {

		if *obj.obj.LsaRetransmitTime < 1 || *obj.obj.LsaRetransmitTime > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= DeviceOspfv3Router.LsaRetransmitTime <= 4294967295 but Got %d", *obj.obj.LsaRetransmitTime))
		}

	}

	if obj.obj.LsaRefreshTime != nil {

		if *obj.obj.LsaRefreshTime < 5 || *obj.obj.LsaRefreshTime > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("5 <= DeviceOspfv3Router.LsaRefreshTime <= 4294967295 but Got %d", *obj.obj.LsaRefreshTime))
		}

	}

	if obj.obj.InterBurstLsuInterval != nil {

		if *obj.obj.InterBurstLsuInterval < 1 || *obj.obj.InterBurstLsuInterval > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= DeviceOspfv3Router.InterBurstLsuInterval <= 4294967295 but Got %d", *obj.obj.InterBurstLsuInterval))
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
				obj.Interfaces().appendHolderSlice(&ospfv3Interface{obj: item})
			}
		}
		for _, item := range obj.Interfaces().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.V6Routes) != 0 {

		if set_default {
			obj.V6Routes().clearHolderSlice()
			for _, item := range obj.obj.V6Routes {
				obj.V6Routes().appendHolderSlice(&ospfv3V6RouteRange{obj: item})
			}
		}
		for _, item := range obj.V6Routes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *deviceOspfv3Router) setDefault() {
	if obj.obj.LsaRetransmitTime == nil {
		obj.SetLsaRetransmitTime(5)
	}
	if obj.obj.LsaRefreshTime == nil {
		obj.SetLsaRefreshTime(1800)
	}
	if obj.obj.InterBurstLsuInterval == nil {
		obj.SetInterBurstLsuInterval(1000)
	}
	if obj.obj.StoreLsa == nil {
		obj.SetStoreLsa(false)
	}

}

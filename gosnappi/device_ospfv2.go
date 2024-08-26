package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceOspfv2 *****
type deviceOspfv2 struct {
	validation
	obj                   *otg.DeviceOspfv2
	marshaller            marshalDeviceOspfv2
	unMarshaller          unMarshalDeviceOspfv2
	gracefulRestartHolder Ospfv2GracefulRestart
	capabilitiesHolder    Ospfv2Options
	interfacesHolder      DeviceOspfv2Ospfv2InterfaceIter
	v4RoutesHolder        DeviceOspfv2Ospfv2V4RouteRangeIter
}

func NewDeviceOspfv2() DeviceOspfv2 {
	obj := deviceOspfv2{obj: &otg.DeviceOspfv2{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceOspfv2) msg() *otg.DeviceOspfv2 {
	return obj.obj
}

func (obj *deviceOspfv2) setMsg(msg *otg.DeviceOspfv2) DeviceOspfv2 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceOspfv2 struct {
	obj *deviceOspfv2
}

type marshalDeviceOspfv2 interface {
	// ToProto marshals DeviceOspfv2 to protobuf object *otg.DeviceOspfv2
	ToProto() (*otg.DeviceOspfv2, error)
	// ToPbText marshals DeviceOspfv2 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceOspfv2 to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceOspfv2 to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceOspfv2 struct {
	obj *deviceOspfv2
}

type unMarshalDeviceOspfv2 interface {
	// FromProto unmarshals DeviceOspfv2 from protobuf object *otg.DeviceOspfv2
	FromProto(msg *otg.DeviceOspfv2) (DeviceOspfv2, error)
	// FromPbText unmarshals DeviceOspfv2 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceOspfv2 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceOspfv2 from JSON text
	FromJson(value string) error
}

func (obj *deviceOspfv2) Marshal() marshalDeviceOspfv2 {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceOspfv2{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceOspfv2) Unmarshal() unMarshalDeviceOspfv2 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceOspfv2{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceOspfv2) ToProto() (*otg.DeviceOspfv2, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceOspfv2) FromProto(msg *otg.DeviceOspfv2) (DeviceOspfv2, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceOspfv2) ToPbText() (string, error) {
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

func (m *unMarshaldeviceOspfv2) FromPbText(value string) error {
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

func (m *marshaldeviceOspfv2) ToYaml() (string, error) {
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

func (m *unMarshaldeviceOspfv2) FromYaml(value string) error {
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

func (m *marshaldeviceOspfv2) ToJson() (string, error) {
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

func (m *unMarshaldeviceOspfv2) FromJson(value string) error {
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

func (obj *deviceOspfv2) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceOspfv2) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceOspfv2) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceOspfv2) Clone() (DeviceOspfv2, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceOspfv2()
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

func (obj *deviceOspfv2) setNil() {
	obj.gracefulRestartHolder = nil
	obj.capabilitiesHolder = nil
	obj.interfacesHolder = nil
	obj.v4RoutesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceOspfv2 is a container of properties for an OSPFv2 router and its interfaces & Route Ranges.
type DeviceOspfv2 interface {
	Validation
	// msg marshals DeviceOspfv2 to protobuf object *otg.DeviceOspfv2
	// and doesn't set defaults
	msg() *otg.DeviceOspfv2
	// setMsg unmarshals DeviceOspfv2 from protobuf object *otg.DeviceOspfv2
	// and doesn't set defaults
	setMsg(*otg.DeviceOspfv2) DeviceOspfv2
	// provides marshal interface
	Marshal() marshalDeviceOspfv2
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceOspfv2
	// validate validates DeviceOspfv2
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceOspfv2, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in DeviceOspfv2.
	Name() string
	// SetName assigns string provided by user to DeviceOspfv2
	SetName(value string) DeviceOspfv2
	// Choice returns DeviceOspfv2ChoiceEnum, set in DeviceOspfv2
	Choice() DeviceOspfv2ChoiceEnum
	// setChoice assigns DeviceOspfv2ChoiceEnum provided by user to DeviceOspfv2
	setChoice(value DeviceOspfv2ChoiceEnum) DeviceOspfv2
	// HasChoice checks if Choice has been set in DeviceOspfv2
	HasChoice() bool
	// getter for InterfaceIp to set choice.
	InterfaceIp()
	// CustomRouterId returns string, set in DeviceOspfv2.
	CustomRouterId() string
	// SetCustomRouterId assigns string provided by user to DeviceOspfv2
	SetCustomRouterId(value string) DeviceOspfv2
	// HasCustomRouterId checks if CustomRouterId has been set in DeviceOspfv2
	HasCustomRouterId() bool
	// LsaRetransmitTimer returns uint32, set in DeviceOspfv2.
	LsaRetransmitTimer() uint32
	// SetLsaRetransmitTimer assigns uint32 provided by user to DeviceOspfv2
	SetLsaRetransmitTimer(value uint32) DeviceOspfv2
	// HasLsaRetransmitTimer checks if LsaRetransmitTimer has been set in DeviceOspfv2
	HasLsaRetransmitTimer() bool
	// LsaRefreshTime returns uint32, set in DeviceOspfv2.
	LsaRefreshTime() uint32
	// SetLsaRefreshTime assigns uint32 provided by user to DeviceOspfv2
	SetLsaRefreshTime(value uint32) DeviceOspfv2
	// HasLsaRefreshTime checks if LsaRefreshTime has been set in DeviceOspfv2
	HasLsaRefreshTime() bool
	// InterBurstLsuInterval returns uint32, set in DeviceOspfv2.
	InterBurstLsuInterval() uint32
	// SetInterBurstLsuInterval assigns uint32 provided by user to DeviceOspfv2
	SetInterBurstLsuInterval(value uint32) DeviceOspfv2
	// HasInterBurstLsuInterval checks if InterBurstLsuInterval has been set in DeviceOspfv2
	HasInterBurstLsuInterval() bool
	// MaxFloodLsuPerBurst returns uint32, set in DeviceOspfv2.
	MaxFloodLsuPerBurst() uint32
	// SetMaxFloodLsuPerBurst assigns uint32 provided by user to DeviceOspfv2
	SetMaxFloodLsuPerBurst(value uint32) DeviceOspfv2
	// HasMaxFloodLsuPerBurst checks if MaxFloodLsuPerBurst has been set in DeviceOspfv2
	HasMaxFloodLsuPerBurst() bool
	// GracefulRestart returns Ospfv2GracefulRestart, set in DeviceOspfv2.
	// Ospfv2GracefulRestart is containter of properties of OSPFv2 Graceful Retstart.
	GracefulRestart() Ospfv2GracefulRestart
	// SetGracefulRestart assigns Ospfv2GracefulRestart provided by user to DeviceOspfv2.
	// Ospfv2GracefulRestart is containter of properties of OSPFv2 Graceful Retstart.
	SetGracefulRestart(value Ospfv2GracefulRestart) DeviceOspfv2
	// HasGracefulRestart checks if GracefulRestart has been set in DeviceOspfv2
	HasGracefulRestart() bool
	// LearnedLspFilter returns bool, set in DeviceOspfv2.
	LearnedLspFilter() bool
	// SetLearnedLspFilter assigns bool provided by user to DeviceOspfv2
	SetLearnedLspFilter(value bool) DeviceOspfv2
	// HasLearnedLspFilter checks if LearnedLspFilter has been set in DeviceOspfv2
	HasLearnedLspFilter() bool
	// Capabilities returns Ospfv2Options, set in DeviceOspfv2.
	// Ospfv2Options is the OSPFv2 Options field is present Database Description packets and all LSAs.
	// This enables OSPF routers to support (or not support) optional capabilities,
	// and to communicate their capability level to other OSPF routers.
	// When capabilities are exchanged in Database Description packets a
	// router can choose not to forward certain LSAs to a neighbor because
	// of its reduced functionality.
	// Reference: A.2 The Options field: https://www.rfc-editor.org/rfc/rfc2328#page-46.
	Capabilities() Ospfv2Options
	// SetCapabilities assigns Ospfv2Options provided by user to DeviceOspfv2.
	// Ospfv2Options is the OSPFv2 Options field is present Database Description packets and all LSAs.
	// This enables OSPF routers to support (or not support) optional capabilities,
	// and to communicate their capability level to other OSPF routers.
	// When capabilities are exchanged in Database Description packets a
	// router can choose not to forward certain LSAs to a neighbor because
	// of its reduced functionality.
	// Reference: A.2 The Options field: https://www.rfc-editor.org/rfc/rfc2328#page-46.
	SetCapabilities(value Ospfv2Options) DeviceOspfv2
	// HasCapabilities checks if Capabilities has been set in DeviceOspfv2
	HasCapabilities() bool
	// Interfaces returns DeviceOspfv2Ospfv2InterfaceIterIter, set in DeviceOspfv2
	Interfaces() DeviceOspfv2Ospfv2InterfaceIter
	// V4Routes returns DeviceOspfv2Ospfv2V4RouteRangeIterIter, set in DeviceOspfv2
	V4Routes() DeviceOspfv2Ospfv2V4RouteRangeIter
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *deviceOspfv2) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the DeviceOspfv2 object
func (obj *deviceOspfv2) SetName(value string) DeviceOspfv2 {

	obj.obj.Name = &value
	return obj
}

type DeviceOspfv2ChoiceEnum string

// Enum of Choice on DeviceOspfv2
var DeviceOspfv2Choice = struct {
	INTERFACE_IP     DeviceOspfv2ChoiceEnum
	CUSTOM_ROUTER_ID DeviceOspfv2ChoiceEnum
}{
	INTERFACE_IP:     DeviceOspfv2ChoiceEnum("interface_ip"),
	CUSTOM_ROUTER_ID: DeviceOspfv2ChoiceEnum("custom_router_id"),
}

func (obj *deviceOspfv2) Choice() DeviceOspfv2ChoiceEnum {
	return DeviceOspfv2ChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for InterfaceIp to set choice
func (obj *deviceOspfv2) InterfaceIp() {
	obj.setChoice(DeviceOspfv2Choice.INTERFACE_IP)
}

// IP address of Router ID for this emulated OSPFv2 router.
// - interface_ip: When IPv4 interface address to be assigned as Router ID.
// - custom_router_id: When, Router ID needs to be configured different from Interface IPv4 address.
// Choice returns a string
func (obj *deviceOspfv2) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *deviceOspfv2) setChoice(value DeviceOspfv2ChoiceEnum) DeviceOspfv2 {
	intValue, ok := otg.DeviceOspfv2_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on DeviceOspfv2ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.DeviceOspfv2_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.CustomRouterId = nil
	return obj
}

// Router ID in IPv4 address format.
// CustomRouterId returns a string
func (obj *deviceOspfv2) CustomRouterId() string {

	if obj.obj.CustomRouterId == nil {
		obj.setChoice(DeviceOspfv2Choice.CUSTOM_ROUTER_ID)
	}

	return *obj.obj.CustomRouterId

}

// Router ID in IPv4 address format.
// CustomRouterId returns a string
func (obj *deviceOspfv2) HasCustomRouterId() bool {
	return obj.obj.CustomRouterId != nil
}

// Router ID in IPv4 address format.
// SetCustomRouterId sets the string value in the DeviceOspfv2 object
func (obj *deviceOspfv2) SetCustomRouterId(value string) DeviceOspfv2 {
	obj.setChoice(DeviceOspfv2Choice.CUSTOM_ROUTER_ID)
	obj.obj.CustomRouterId = &value
	return obj
}

// The time in seconds for LSA retransmission.
// LsaRetransmitTimer returns a uint32
func (obj *deviceOspfv2) LsaRetransmitTimer() uint32 {

	return *obj.obj.LsaRetransmitTimer

}

// The time in seconds for LSA retransmission.
// LsaRetransmitTimer returns a uint32
func (obj *deviceOspfv2) HasLsaRetransmitTimer() bool {
	return obj.obj.LsaRetransmitTimer != nil
}

// The time in seconds for LSA retransmission.
// SetLsaRetransmitTimer sets the uint32 value in the DeviceOspfv2 object
func (obj *deviceOspfv2) SetLsaRetransmitTimer(value uint32) DeviceOspfv2 {

	obj.obj.LsaRetransmitTimer = &value
	return obj
}

// The time in seconds required for LSA refresh.
// LsaRefreshTime returns a uint32
func (obj *deviceOspfv2) LsaRefreshTime() uint32 {

	return *obj.obj.LsaRefreshTime

}

// The time in seconds required for LSA refresh.
// LsaRefreshTime returns a uint32
func (obj *deviceOspfv2) HasLsaRefreshTime() bool {
	return obj.obj.LsaRefreshTime != nil
}

// The time in seconds required for LSA refresh.
// SetLsaRefreshTime sets the uint32 value in the DeviceOspfv2 object
func (obj *deviceOspfv2) SetLsaRefreshTime(value uint32) DeviceOspfv2 {

	obj.obj.LsaRefreshTime = &value
	return obj
}

// The gap in miliseconds between each Flood Link State Update Burst
// InterBurstLsuInterval returns a uint32
func (obj *deviceOspfv2) InterBurstLsuInterval() uint32 {

	return *obj.obj.InterBurstLsuInterval

}

// The gap in miliseconds between each Flood Link State Update Burst
// InterBurstLsuInterval returns a uint32
func (obj *deviceOspfv2) HasInterBurstLsuInterval() bool {
	return obj.obj.InterBurstLsuInterval != nil
}

// The gap in miliseconds between each Flood Link State Update Burst
// SetInterBurstLsuInterval sets the uint32 value in the DeviceOspfv2 object
func (obj *deviceOspfv2) SetInterBurstLsuInterval(value uint32) DeviceOspfv2 {

	obj.obj.InterBurstLsuInterval = &value
	return obj
}

// The maximum number of Flood LSUpdates for each burst
// MaxFloodLsuPerBurst returns a uint32
func (obj *deviceOspfv2) MaxFloodLsuPerBurst() uint32 {

	return *obj.obj.MaxFloodLsuPerBurst

}

// The maximum number of Flood LSUpdates for each burst
// MaxFloodLsuPerBurst returns a uint32
func (obj *deviceOspfv2) HasMaxFloodLsuPerBurst() bool {
	return obj.obj.MaxFloodLsuPerBurst != nil
}

// The maximum number of Flood LSUpdates for each burst
// SetMaxFloodLsuPerBurst sets the uint32 value in the DeviceOspfv2 object
func (obj *deviceOspfv2) SetMaxFloodLsuPerBurst(value uint32) DeviceOspfv2 {

	obj.obj.MaxFloodLsuPerBurst = &value
	return obj
}

// description is TBD
// GracefulRestart returns a Ospfv2GracefulRestart
func (obj *deviceOspfv2) GracefulRestart() Ospfv2GracefulRestart {
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
func (obj *deviceOspfv2) HasGracefulRestart() bool {
	return obj.obj.GracefulRestart != nil
}

// description is TBD
// SetGracefulRestart sets the Ospfv2GracefulRestart value in the DeviceOspfv2 object
func (obj *deviceOspfv2) SetGracefulRestart(value Ospfv2GracefulRestart) DeviceOspfv2 {

	obj.gracefulRestartHolder = nil
	obj.obj.GracefulRestart = value.msg()

	return obj
}

// Configuration for controlling storage of OSPFv2 learned LSAs are received from the neighbors.
// LearnedLspFilter returns a bool
func (obj *deviceOspfv2) LearnedLspFilter() bool {

	return *obj.obj.LearnedLspFilter

}

// Configuration for controlling storage of OSPFv2 learned LSAs are received from the neighbors.
// LearnedLspFilter returns a bool
func (obj *deviceOspfv2) HasLearnedLspFilter() bool {
	return obj.obj.LearnedLspFilter != nil
}

// Configuration for controlling storage of OSPFv2 learned LSAs are received from the neighbors.
// SetLearnedLspFilter sets the bool value in the DeviceOspfv2 object
func (obj *deviceOspfv2) SetLearnedLspFilter(value bool) DeviceOspfv2 {

	obj.obj.LearnedLspFilter = &value
	return obj
}

// A router indicates the optional capabilities that it supports in its OSPF Hello packets, Database Description packets and in its LSAs.
// Capabilities returns a Ospfv2Options
func (obj *deviceOspfv2) Capabilities() Ospfv2Options {
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
func (obj *deviceOspfv2) HasCapabilities() bool {
	return obj.obj.Capabilities != nil
}

// A router indicates the optional capabilities that it supports in its OSPF Hello packets, Database Description packets and in its LSAs.
// SetCapabilities sets the Ospfv2Options value in the DeviceOspfv2 object
func (obj *deviceOspfv2) SetCapabilities(value Ospfv2Options) DeviceOspfv2 {

	obj.capabilitiesHolder = nil
	obj.obj.Capabilities = value.msg()

	return obj
}

// List of OSPFv2 interfaces for this router.
// Interfaces returns a []Ospfv2Interface
func (obj *deviceOspfv2) Interfaces() DeviceOspfv2Ospfv2InterfaceIter {
	if len(obj.obj.Interfaces) == 0 {
		obj.obj.Interfaces = []*otg.Ospfv2Interface{}
	}
	if obj.interfacesHolder == nil {
		obj.interfacesHolder = newDeviceOspfv2Ospfv2InterfaceIter(&obj.obj.Interfaces).setMsg(obj)
	}
	return obj.interfacesHolder
}

type deviceOspfv2Ospfv2InterfaceIter struct {
	obj                  *deviceOspfv2
	ospfv2InterfaceSlice []Ospfv2Interface
	fieldPtr             *[]*otg.Ospfv2Interface
}

func newDeviceOspfv2Ospfv2InterfaceIter(ptr *[]*otg.Ospfv2Interface) DeviceOspfv2Ospfv2InterfaceIter {
	return &deviceOspfv2Ospfv2InterfaceIter{fieldPtr: ptr}
}

type DeviceOspfv2Ospfv2InterfaceIter interface {
	setMsg(*deviceOspfv2) DeviceOspfv2Ospfv2InterfaceIter
	Items() []Ospfv2Interface
	Add() Ospfv2Interface
	Append(items ...Ospfv2Interface) DeviceOspfv2Ospfv2InterfaceIter
	Set(index int, newObj Ospfv2Interface) DeviceOspfv2Ospfv2InterfaceIter
	Clear() DeviceOspfv2Ospfv2InterfaceIter
	clearHolderSlice() DeviceOspfv2Ospfv2InterfaceIter
	appendHolderSlice(item Ospfv2Interface) DeviceOspfv2Ospfv2InterfaceIter
}

func (obj *deviceOspfv2Ospfv2InterfaceIter) setMsg(msg *deviceOspfv2) DeviceOspfv2Ospfv2InterfaceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2Interface{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceOspfv2Ospfv2InterfaceIter) Items() []Ospfv2Interface {
	return obj.ospfv2InterfaceSlice
}

func (obj *deviceOspfv2Ospfv2InterfaceIter) Add() Ospfv2Interface {
	newObj := &otg.Ospfv2Interface{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2Interface{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2InterfaceSlice = append(obj.ospfv2InterfaceSlice, newLibObj)
	return newLibObj
}

func (obj *deviceOspfv2Ospfv2InterfaceIter) Append(items ...Ospfv2Interface) DeviceOspfv2Ospfv2InterfaceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2InterfaceSlice = append(obj.ospfv2InterfaceSlice, item)
	}
	return obj
}

func (obj *deviceOspfv2Ospfv2InterfaceIter) Set(index int, newObj Ospfv2Interface) DeviceOspfv2Ospfv2InterfaceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2InterfaceSlice[index] = newObj
	return obj
}
func (obj *deviceOspfv2Ospfv2InterfaceIter) Clear() DeviceOspfv2Ospfv2InterfaceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2Interface{}
		obj.ospfv2InterfaceSlice = []Ospfv2Interface{}
	}
	return obj
}
func (obj *deviceOspfv2Ospfv2InterfaceIter) clearHolderSlice() DeviceOspfv2Ospfv2InterfaceIter {
	if len(obj.ospfv2InterfaceSlice) > 0 {
		obj.ospfv2InterfaceSlice = []Ospfv2Interface{}
	}
	return obj
}
func (obj *deviceOspfv2Ospfv2InterfaceIter) appendHolderSlice(item Ospfv2Interface) DeviceOspfv2Ospfv2InterfaceIter {
	obj.ospfv2InterfaceSlice = append(obj.ospfv2InterfaceSlice, item)
	return obj
}

// Emulated OSPFv4 IPv4 routes.
// V4Routes returns a []Ospfv2V4RouteRange
func (obj *deviceOspfv2) V4Routes() DeviceOspfv2Ospfv2V4RouteRangeIter {
	if len(obj.obj.V4Routes) == 0 {
		obj.obj.V4Routes = []*otg.Ospfv2V4RouteRange{}
	}
	if obj.v4RoutesHolder == nil {
		obj.v4RoutesHolder = newDeviceOspfv2Ospfv2V4RouteRangeIter(&obj.obj.V4Routes).setMsg(obj)
	}
	return obj.v4RoutesHolder
}

type deviceOspfv2Ospfv2V4RouteRangeIter struct {
	obj                     *deviceOspfv2
	ospfv2V4RouteRangeSlice []Ospfv2V4RouteRange
	fieldPtr                *[]*otg.Ospfv2V4RouteRange
}

func newDeviceOspfv2Ospfv2V4RouteRangeIter(ptr *[]*otg.Ospfv2V4RouteRange) DeviceOspfv2Ospfv2V4RouteRangeIter {
	return &deviceOspfv2Ospfv2V4RouteRangeIter{fieldPtr: ptr}
}

type DeviceOspfv2Ospfv2V4RouteRangeIter interface {
	setMsg(*deviceOspfv2) DeviceOspfv2Ospfv2V4RouteRangeIter
	Items() []Ospfv2V4RouteRange
	Add() Ospfv2V4RouteRange
	Append(items ...Ospfv2V4RouteRange) DeviceOspfv2Ospfv2V4RouteRangeIter
	Set(index int, newObj Ospfv2V4RouteRange) DeviceOspfv2Ospfv2V4RouteRangeIter
	Clear() DeviceOspfv2Ospfv2V4RouteRangeIter
	clearHolderSlice() DeviceOspfv2Ospfv2V4RouteRangeIter
	appendHolderSlice(item Ospfv2V4RouteRange) DeviceOspfv2Ospfv2V4RouteRangeIter
}

func (obj *deviceOspfv2Ospfv2V4RouteRangeIter) setMsg(msg *deviceOspfv2) DeviceOspfv2Ospfv2V4RouteRangeIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2V4RouteRange{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceOspfv2Ospfv2V4RouteRangeIter) Items() []Ospfv2V4RouteRange {
	return obj.ospfv2V4RouteRangeSlice
}

func (obj *deviceOspfv2Ospfv2V4RouteRangeIter) Add() Ospfv2V4RouteRange {
	newObj := &otg.Ospfv2V4RouteRange{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2V4RouteRange{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2V4RouteRangeSlice = append(obj.ospfv2V4RouteRangeSlice, newLibObj)
	return newLibObj
}

func (obj *deviceOspfv2Ospfv2V4RouteRangeIter) Append(items ...Ospfv2V4RouteRange) DeviceOspfv2Ospfv2V4RouteRangeIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2V4RouteRangeSlice = append(obj.ospfv2V4RouteRangeSlice, item)
	}
	return obj
}

func (obj *deviceOspfv2Ospfv2V4RouteRangeIter) Set(index int, newObj Ospfv2V4RouteRange) DeviceOspfv2Ospfv2V4RouteRangeIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2V4RouteRangeSlice[index] = newObj
	return obj
}
func (obj *deviceOspfv2Ospfv2V4RouteRangeIter) Clear() DeviceOspfv2Ospfv2V4RouteRangeIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2V4RouteRange{}
		obj.ospfv2V4RouteRangeSlice = []Ospfv2V4RouteRange{}
	}
	return obj
}
func (obj *deviceOspfv2Ospfv2V4RouteRangeIter) clearHolderSlice() DeviceOspfv2Ospfv2V4RouteRangeIter {
	if len(obj.ospfv2V4RouteRangeSlice) > 0 {
		obj.ospfv2V4RouteRangeSlice = []Ospfv2V4RouteRange{}
	}
	return obj
}
func (obj *deviceOspfv2Ospfv2V4RouteRangeIter) appendHolderSlice(item Ospfv2V4RouteRange) DeviceOspfv2Ospfv2V4RouteRangeIter {
	obj.ospfv2V4RouteRangeSlice = append(obj.ospfv2V4RouteRangeSlice, item)
	return obj
}

func (obj *deviceOspfv2) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface DeviceOspfv2")
	}

	if obj.obj.CustomRouterId != nil {

		err := obj.validateIpv4(obj.CustomRouterId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DeviceOspfv2.CustomRouterId"))
		}

	}

	if obj.obj.LsaRetransmitTimer != nil {

		if *obj.obj.LsaRetransmitTimer < 1 || *obj.obj.LsaRetransmitTimer > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= DeviceOspfv2.LsaRetransmitTimer <= 4294967295 but Got %d", *obj.obj.LsaRetransmitTimer))
		}

	}

	if obj.obj.LsaRefreshTime != nil {

		if *obj.obj.LsaRefreshTime < 5 || *obj.obj.LsaRefreshTime > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("5 <= DeviceOspfv2.LsaRefreshTime <= 4294967295 but Got %d", *obj.obj.LsaRefreshTime))
		}

	}

	if obj.obj.MaxFloodLsuPerBurst != nil {

		if *obj.obj.MaxFloodLsuPerBurst < 1 || *obj.obj.MaxFloodLsuPerBurst > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= DeviceOspfv2.MaxFloodLsuPerBurst <= 4294967295 but Got %d", *obj.obj.MaxFloodLsuPerBurst))
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

func (obj *deviceOspfv2) setDefault() {
	var choices_set int = 0
	var choice DeviceOspfv2ChoiceEnum

	if obj.obj.CustomRouterId != nil {
		choices_set += 1
		choice = DeviceOspfv2Choice.CUSTOM_ROUTER_ID
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(DeviceOspfv2Choice.INTERFACE_IP)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in DeviceOspfv2")
			}
		} else {
			intVal := otg.DeviceOspfv2_Choice_Enum_value[string(choice)]
			enumValue := otg.DeviceOspfv2_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.LsaRetransmitTimer == nil {
		obj.SetLsaRetransmitTimer(5)
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
	if obj.obj.LearnedLspFilter == nil {
		obj.SetLearnedLspFilter(false)
	}

}

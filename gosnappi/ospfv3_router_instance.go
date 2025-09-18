package gosnappi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3RouterInstance *****
type ospfv3RouterInstance struct {
	validation
	obj                   *otg.Ospfv3RouterInstance
	marshaller            marshalOspfv3RouterInstance
	unMarshaller          unMarshalOspfv3RouterInstance
	gracefulRestartHolder Ospfv3GracefulRestart
	capabilitiesHolder    Ospfv3Capabilities
	interfacesHolder      Ospfv3RouterInstanceOspfv3InterfaceIter
	v6RoutesHolder        Ospfv3RouterInstanceOspfv3V6RouteRangeIter
}

func NewOspfv3RouterInstance() Ospfv3RouterInstance {
	obj := ospfv3RouterInstance{obj: &otg.Ospfv3RouterInstance{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3RouterInstance) msg() *otg.Ospfv3RouterInstance {
	return obj.obj
}

func (obj *ospfv3RouterInstance) setMsg(msg *otg.Ospfv3RouterInstance) Ospfv3RouterInstance {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3RouterInstance struct {
	obj *ospfv3RouterInstance
}

type marshalOspfv3RouterInstance interface {
	// ToProto marshals Ospfv3RouterInstance to protobuf object *otg.Ospfv3RouterInstance
	ToProto() (*otg.Ospfv3RouterInstance, error)
	// ToPbText marshals Ospfv3RouterInstance to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3RouterInstance to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3RouterInstance to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3RouterInstance struct {
	obj *ospfv3RouterInstance
}

type unMarshalOspfv3RouterInstance interface {
	// FromProto unmarshals Ospfv3RouterInstance from protobuf object *otg.Ospfv3RouterInstance
	FromProto(msg *otg.Ospfv3RouterInstance) (Ospfv3RouterInstance, error)
	// FromPbText unmarshals Ospfv3RouterInstance from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3RouterInstance from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3RouterInstance from JSON text
	FromJson(value string) error
}

func (obj *ospfv3RouterInstance) Marshal() marshalOspfv3RouterInstance {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3RouterInstance{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3RouterInstance) Unmarshal() unMarshalOspfv3RouterInstance {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3RouterInstance{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3RouterInstance) ToProto() (*otg.Ospfv3RouterInstance, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3RouterInstance) FromProto(msg *otg.Ospfv3RouterInstance) (Ospfv3RouterInstance, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3RouterInstance) ToPbText() (string, error) {
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

func (m *unMarshalospfv3RouterInstance) FromPbText(value string) error {
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

func (m *marshalospfv3RouterInstance) ToYaml() (string, error) {
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

func (m *unMarshalospfv3RouterInstance) FromYaml(value string) error {
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

func (m *marshalospfv3RouterInstance) ToJson() (string, error) {
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

func (m *unMarshalospfv3RouterInstance) FromJson(value string) error {
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

func (obj *ospfv3RouterInstance) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3RouterInstance) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3RouterInstance) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3RouterInstance) Clone() (Ospfv3RouterInstance, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3RouterInstance()
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

func (obj *ospfv3RouterInstance) setNil() {
	obj.gracefulRestartHolder = nil
	obj.capabilitiesHolder = nil
	obj.interfacesHolder = nil
	obj.v6RoutesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv3RouterInstance is a container of properties for an OSPFv3 router instance and its interfaces and route ranges.
type Ospfv3RouterInstance interface {
	Validation
	// msg marshals Ospfv3RouterInstance to protobuf object *otg.Ospfv3RouterInstance
	// and doesn't set defaults
	msg() *otg.Ospfv3RouterInstance
	// setMsg unmarshals Ospfv3RouterInstance from protobuf object *otg.Ospfv3RouterInstance
	// and doesn't set defaults
	setMsg(*otg.Ospfv3RouterInstance) Ospfv3RouterInstance
	// provides marshal interface
	Marshal() marshalOspfv3RouterInstance
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3RouterInstance
	// validate validates Ospfv3RouterInstance
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3RouterInstance, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in Ospfv3RouterInstance.
	Name() string
	// SetName assigns string provided by user to Ospfv3RouterInstance
	SetName(value string) Ospfv3RouterInstance
	// LsaRetransmitTime returns uint32, set in Ospfv3RouterInstance.
	LsaRetransmitTime() uint32
	// SetLsaRetransmitTime assigns uint32 provided by user to Ospfv3RouterInstance
	SetLsaRetransmitTime(value uint32) Ospfv3RouterInstance
	// HasLsaRetransmitTime checks if LsaRetransmitTime has been set in Ospfv3RouterInstance
	HasLsaRetransmitTime() bool
	// LsaRefreshTime returns uint32, set in Ospfv3RouterInstance.
	LsaRefreshTime() uint32
	// SetLsaRefreshTime assigns uint32 provided by user to Ospfv3RouterInstance
	SetLsaRefreshTime(value uint32) Ospfv3RouterInstance
	// HasLsaRefreshTime checks if LsaRefreshTime has been set in Ospfv3RouterInstance
	HasLsaRefreshTime() bool
	// InterBurstLsuInterval returns uint32, set in Ospfv3RouterInstance.
	InterBurstLsuInterval() uint32
	// SetInterBurstLsuInterval assigns uint32 provided by user to Ospfv3RouterInstance
	SetInterBurstLsuInterval(value uint32) Ospfv3RouterInstance
	// HasInterBurstLsuInterval checks if InterBurstLsuInterval has been set in Ospfv3RouterInstance
	HasInterBurstLsuInterval() bool
	// GracefulRestart returns Ospfv3GracefulRestart, set in Ospfv3RouterInstance.
	// Ospfv3GracefulRestart is container of properties of OSPFv3 Graceful Retstart.
	GracefulRestart() Ospfv3GracefulRestart
	// SetGracefulRestart assigns Ospfv3GracefulRestart provided by user to Ospfv3RouterInstance.
	// Ospfv3GracefulRestart is container of properties of OSPFv3 Graceful Retstart.
	SetGracefulRestart(value Ospfv3GracefulRestart) Ospfv3RouterInstance
	// HasGracefulRestart checks if GracefulRestart has been set in Ospfv3RouterInstance
	HasGracefulRestart() bool
	// StoreLsa returns bool, set in Ospfv3RouterInstance.
	StoreLsa() bool
	// SetStoreLsa assigns bool provided by user to Ospfv3RouterInstance
	SetStoreLsa(value bool) Ospfv3RouterInstance
	// HasStoreLsa checks if StoreLsa has been set in Ospfv3RouterInstance
	HasStoreLsa() bool
	// Capabilities returns Ospfv3Capabilities, set in Ospfv3RouterInstance.
	// Ospfv3Capabilities is the OSPFv3 router optional attributes.
	Capabilities() Ospfv3Capabilities
	// SetCapabilities assigns Ospfv3Capabilities provided by user to Ospfv3RouterInstance.
	// Ospfv3Capabilities is the OSPFv3 router optional attributes.
	SetCapabilities(value Ospfv3Capabilities) Ospfv3RouterInstance
	// HasCapabilities checks if Capabilities has been set in Ospfv3RouterInstance
	HasCapabilities() bool
	// Interfaces returns Ospfv3RouterInstanceOspfv3InterfaceIterIter, set in Ospfv3RouterInstance
	Interfaces() Ospfv3RouterInstanceOspfv3InterfaceIter
	// V6Routes returns Ospfv3RouterInstanceOspfv3V6RouteRangeIterIter, set in Ospfv3RouterInstance
	V6Routes() Ospfv3RouterInstanceOspfv3V6RouteRangeIter
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *ospfv3RouterInstance) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the Ospfv3RouterInstance object
func (obj *ospfv3RouterInstance) SetName(value string) Ospfv3RouterInstance {

	obj.obj.Name = &value
	return obj
}

// The time in seconds for LSA retransmission.
// LsaRetransmitTime returns a uint32
func (obj *ospfv3RouterInstance) LsaRetransmitTime() uint32 {

	return *obj.obj.LsaRetransmitTime

}

// The time in seconds for LSA retransmission.
// LsaRetransmitTime returns a uint32
func (obj *ospfv3RouterInstance) HasLsaRetransmitTime() bool {
	return obj.obj.LsaRetransmitTime != nil
}

// The time in seconds for LSA retransmission.
// SetLsaRetransmitTime sets the uint32 value in the Ospfv3RouterInstance object
func (obj *ospfv3RouterInstance) SetLsaRetransmitTime(value uint32) Ospfv3RouterInstance {

	obj.obj.LsaRetransmitTime = &value
	return obj
}

// The time in seconds required for LSA refresh.
// LsaRefreshTime returns a uint32
func (obj *ospfv3RouterInstance) LsaRefreshTime() uint32 {

	return *obj.obj.LsaRefreshTime

}

// The time in seconds required for LSA refresh.
// LsaRefreshTime returns a uint32
func (obj *ospfv3RouterInstance) HasLsaRefreshTime() bool {
	return obj.obj.LsaRefreshTime != nil
}

// The time in seconds required for LSA refresh.
// SetLsaRefreshTime sets the uint32 value in the Ospfv3RouterInstance object
func (obj *ospfv3RouterInstance) SetLsaRefreshTime(value uint32) Ospfv3RouterInstance {

	obj.obj.LsaRefreshTime = &value
	return obj
}

// The gap in miliseconds between each Flood Link State Update burst.
// InterBurstLsuInterval returns a uint32
func (obj *ospfv3RouterInstance) InterBurstLsuInterval() uint32 {

	return *obj.obj.InterBurstLsuInterval

}

// The gap in miliseconds between each Flood Link State Update burst.
// InterBurstLsuInterval returns a uint32
func (obj *ospfv3RouterInstance) HasInterBurstLsuInterval() bool {
	return obj.obj.InterBurstLsuInterval != nil
}

// The gap in miliseconds between each Flood Link State Update burst.
// SetInterBurstLsuInterval sets the uint32 value in the Ospfv3RouterInstance object
func (obj *ospfv3RouterInstance) SetInterBurstLsuInterval(value uint32) Ospfv3RouterInstance {

	obj.obj.InterBurstLsuInterval = &value
	return obj
}

// description is TBD
// GracefulRestart returns a Ospfv3GracefulRestart
func (obj *ospfv3RouterInstance) GracefulRestart() Ospfv3GracefulRestart {
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
func (obj *ospfv3RouterInstance) HasGracefulRestart() bool {
	return obj.obj.GracefulRestart != nil
}

// description is TBD
// SetGracefulRestart sets the Ospfv3GracefulRestart value in the Ospfv3RouterInstance object
func (obj *ospfv3RouterInstance) SetGracefulRestart(value Ospfv3GracefulRestart) Ospfv3RouterInstance {

	obj.gracefulRestartHolder = nil
	obj.obj.GracefulRestart = value.msg()

	return obj
}

// Configuration for controlling storage of OSPFv3 learned LSAs received from the neighbors.
// StoreLsa returns a bool
func (obj *ospfv3RouterInstance) StoreLsa() bool {

	return *obj.obj.StoreLsa

}

// Configuration for controlling storage of OSPFv3 learned LSAs received from the neighbors.
// StoreLsa returns a bool
func (obj *ospfv3RouterInstance) HasStoreLsa() bool {
	return obj.obj.StoreLsa != nil
}

// Configuration for controlling storage of OSPFv3 learned LSAs received from the neighbors.
// SetStoreLsa sets the bool value in the Ospfv3RouterInstance object
func (obj *ospfv3RouterInstance) SetStoreLsa(value bool) Ospfv3RouterInstance {

	obj.obj.StoreLsa = &value
	return obj
}

// Optional container for OSPFv3 router capabilities.
// Capabilities returns a Ospfv3Capabilities
func (obj *ospfv3RouterInstance) Capabilities() Ospfv3Capabilities {
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
func (obj *ospfv3RouterInstance) HasCapabilities() bool {
	return obj.obj.Capabilities != nil
}

// Optional container for OSPFv3 router capabilities.
// SetCapabilities sets the Ospfv3Capabilities value in the Ospfv3RouterInstance object
func (obj *ospfv3RouterInstance) SetCapabilities(value Ospfv3Capabilities) Ospfv3RouterInstance {

	obj.capabilitiesHolder = nil
	obj.obj.Capabilities = value.msg()

	return obj
}

// List of OSPFv3 interfaces for this router.
// Interfaces returns a []Ospfv3Interface
func (obj *ospfv3RouterInstance) Interfaces() Ospfv3RouterInstanceOspfv3InterfaceIter {
	if len(obj.obj.Interfaces) == 0 {
		obj.obj.Interfaces = []*otg.Ospfv3Interface{}
	}
	if obj.interfacesHolder == nil {
		obj.interfacesHolder = newOspfv3RouterInstanceOspfv3InterfaceIter(&obj.obj.Interfaces).setMsg(obj)
	}
	return obj.interfacesHolder
}

type ospfv3RouterInstanceOspfv3InterfaceIter struct {
	obj                  *ospfv3RouterInstance
	ospfv3InterfaceSlice []Ospfv3Interface
	fieldPtr             *[]*otg.Ospfv3Interface
}

func newOspfv3RouterInstanceOspfv3InterfaceIter(ptr *[]*otg.Ospfv3Interface) Ospfv3RouterInstanceOspfv3InterfaceIter {
	return &ospfv3RouterInstanceOspfv3InterfaceIter{fieldPtr: ptr}
}

type Ospfv3RouterInstanceOspfv3InterfaceIter interface {
	setMsg(*ospfv3RouterInstance) Ospfv3RouterInstanceOspfv3InterfaceIter
	Items() []Ospfv3Interface
	Add() Ospfv3Interface
	Append(items ...Ospfv3Interface) Ospfv3RouterInstanceOspfv3InterfaceIter
	Set(index int, newObj Ospfv3Interface) Ospfv3RouterInstanceOspfv3InterfaceIter
	Clear() Ospfv3RouterInstanceOspfv3InterfaceIter
	clearHolderSlice() Ospfv3RouterInstanceOspfv3InterfaceIter
	appendHolderSlice(item Ospfv3Interface) Ospfv3RouterInstanceOspfv3InterfaceIter
}

func (obj *ospfv3RouterInstanceOspfv3InterfaceIter) setMsg(msg *ospfv3RouterInstance) Ospfv3RouterInstanceOspfv3InterfaceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv3Interface{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv3RouterInstanceOspfv3InterfaceIter) Items() []Ospfv3Interface {
	return obj.ospfv3InterfaceSlice
}

func (obj *ospfv3RouterInstanceOspfv3InterfaceIter) Add() Ospfv3Interface {
	newObj := &otg.Ospfv3Interface{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv3Interface{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv3InterfaceSlice = append(obj.ospfv3InterfaceSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv3RouterInstanceOspfv3InterfaceIter) Append(items ...Ospfv3Interface) Ospfv3RouterInstanceOspfv3InterfaceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv3InterfaceSlice = append(obj.ospfv3InterfaceSlice, item)
	}
	return obj
}

func (obj *ospfv3RouterInstanceOspfv3InterfaceIter) Set(index int, newObj Ospfv3Interface) Ospfv3RouterInstanceOspfv3InterfaceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv3InterfaceSlice[index] = newObj
	return obj
}
func (obj *ospfv3RouterInstanceOspfv3InterfaceIter) Clear() Ospfv3RouterInstanceOspfv3InterfaceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv3Interface{}
		obj.ospfv3InterfaceSlice = []Ospfv3Interface{}
	}
	return obj
}
func (obj *ospfv3RouterInstanceOspfv3InterfaceIter) clearHolderSlice() Ospfv3RouterInstanceOspfv3InterfaceIter {
	if len(obj.ospfv3InterfaceSlice) > 0 {
		obj.ospfv3InterfaceSlice = []Ospfv3Interface{}
	}
	return obj
}
func (obj *ospfv3RouterInstanceOspfv3InterfaceIter) appendHolderSlice(item Ospfv3Interface) Ospfv3RouterInstanceOspfv3InterfaceIter {
	obj.ospfv3InterfaceSlice = append(obj.ospfv3InterfaceSlice, item)
	return obj
}

// Emulated OSPFv4 IPv6 routes.
// V6Routes returns a []Ospfv3V6RouteRange
func (obj *ospfv3RouterInstance) V6Routes() Ospfv3RouterInstanceOspfv3V6RouteRangeIter {
	if len(obj.obj.V6Routes) == 0 {
		obj.obj.V6Routes = []*otg.Ospfv3V6RouteRange{}
	}
	if obj.v6RoutesHolder == nil {
		obj.v6RoutesHolder = newOspfv3RouterInstanceOspfv3V6RouteRangeIter(&obj.obj.V6Routes).setMsg(obj)
	}
	return obj.v6RoutesHolder
}

type ospfv3RouterInstanceOspfv3V6RouteRangeIter struct {
	obj                     *ospfv3RouterInstance
	ospfv3V6RouteRangeSlice []Ospfv3V6RouteRange
	fieldPtr                *[]*otg.Ospfv3V6RouteRange
}

func newOspfv3RouterInstanceOspfv3V6RouteRangeIter(ptr *[]*otg.Ospfv3V6RouteRange) Ospfv3RouterInstanceOspfv3V6RouteRangeIter {
	return &ospfv3RouterInstanceOspfv3V6RouteRangeIter{fieldPtr: ptr}
}

type Ospfv3RouterInstanceOspfv3V6RouteRangeIter interface {
	setMsg(*ospfv3RouterInstance) Ospfv3RouterInstanceOspfv3V6RouteRangeIter
	Items() []Ospfv3V6RouteRange
	Add() Ospfv3V6RouteRange
	Append(items ...Ospfv3V6RouteRange) Ospfv3RouterInstanceOspfv3V6RouteRangeIter
	Set(index int, newObj Ospfv3V6RouteRange) Ospfv3RouterInstanceOspfv3V6RouteRangeIter
	Clear() Ospfv3RouterInstanceOspfv3V6RouteRangeIter
	clearHolderSlice() Ospfv3RouterInstanceOspfv3V6RouteRangeIter
	appendHolderSlice(item Ospfv3V6RouteRange) Ospfv3RouterInstanceOspfv3V6RouteRangeIter
}

func (obj *ospfv3RouterInstanceOspfv3V6RouteRangeIter) setMsg(msg *ospfv3RouterInstance) Ospfv3RouterInstanceOspfv3V6RouteRangeIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv3V6RouteRange{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv3RouterInstanceOspfv3V6RouteRangeIter) Items() []Ospfv3V6RouteRange {
	return obj.ospfv3V6RouteRangeSlice
}

func (obj *ospfv3RouterInstanceOspfv3V6RouteRangeIter) Add() Ospfv3V6RouteRange {
	newObj := &otg.Ospfv3V6RouteRange{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv3V6RouteRange{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv3V6RouteRangeSlice = append(obj.ospfv3V6RouteRangeSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv3RouterInstanceOspfv3V6RouteRangeIter) Append(items ...Ospfv3V6RouteRange) Ospfv3RouterInstanceOspfv3V6RouteRangeIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv3V6RouteRangeSlice = append(obj.ospfv3V6RouteRangeSlice, item)
	}
	return obj
}

func (obj *ospfv3RouterInstanceOspfv3V6RouteRangeIter) Set(index int, newObj Ospfv3V6RouteRange) Ospfv3RouterInstanceOspfv3V6RouteRangeIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv3V6RouteRangeSlice[index] = newObj
	return obj
}
func (obj *ospfv3RouterInstanceOspfv3V6RouteRangeIter) Clear() Ospfv3RouterInstanceOspfv3V6RouteRangeIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv3V6RouteRange{}
		obj.ospfv3V6RouteRangeSlice = []Ospfv3V6RouteRange{}
	}
	return obj
}
func (obj *ospfv3RouterInstanceOspfv3V6RouteRangeIter) clearHolderSlice() Ospfv3RouterInstanceOspfv3V6RouteRangeIter {
	if len(obj.ospfv3V6RouteRangeSlice) > 0 {
		obj.ospfv3V6RouteRangeSlice = []Ospfv3V6RouteRange{}
	}
	return obj
}
func (obj *ospfv3RouterInstanceOspfv3V6RouteRangeIter) appendHolderSlice(item Ospfv3V6RouteRange) Ospfv3RouterInstanceOspfv3V6RouteRangeIter {
	obj.ospfv3V6RouteRangeSlice = append(obj.ospfv3V6RouteRangeSlice, item)
	return obj
}

func (obj *ospfv3RouterInstance) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface Ospfv3RouterInstance")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"Ospfv3RouterInstance.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.LsaRetransmitTime != nil {

		if *obj.obj.LsaRetransmitTime < 1 || *obj.obj.LsaRetransmitTime > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= Ospfv3RouterInstance.LsaRetransmitTime <= 4294967295 but Got %d", *obj.obj.LsaRetransmitTime))
		}

	}

	if obj.obj.LsaRefreshTime != nil {

		if *obj.obj.LsaRefreshTime < 5 || *obj.obj.LsaRefreshTime > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("5 <= Ospfv3RouterInstance.LsaRefreshTime <= 4294967295 but Got %d", *obj.obj.LsaRefreshTime))
		}

	}

	if obj.obj.InterBurstLsuInterval != nil {

		if *obj.obj.InterBurstLsuInterval < 1 || *obj.obj.InterBurstLsuInterval > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= Ospfv3RouterInstance.InterBurstLsuInterval <= 4294967295 but Got %d", *obj.obj.InterBurstLsuInterval))
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

func (obj *ospfv3RouterInstance) setDefault() {
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

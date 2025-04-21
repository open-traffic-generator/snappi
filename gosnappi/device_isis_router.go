package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceIsisRouter *****
type deviceIsisRouter struct {
	validation
	obj                   *otg.DeviceIsisRouter
	marshaller            marshalDeviceIsisRouter
	unMarshaller          unMarshalDeviceIsisRouter
	instanceHolder        DeviceIsisMultiInstance
	interfacesHolder      DeviceIsisRouterIsisInterfaceIter
	basicHolder           IsisBasic
	advancedHolder        IsisAdvanced
	routerAuthHolder      IsisAuthentication
	v4RoutesHolder        DeviceIsisRouterIsisV4RouteRangeIter
	v6RoutesHolder        DeviceIsisRouterIsisV6RouteRangeIter
	segmentRoutingHolder  IsisSegmentRouting
	gracefulRestartHolder IsisGracefulRestart
}

func NewDeviceIsisRouter() DeviceIsisRouter {
	obj := deviceIsisRouter{obj: &otg.DeviceIsisRouter{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceIsisRouter) msg() *otg.DeviceIsisRouter {
	return obj.obj
}

func (obj *deviceIsisRouter) setMsg(msg *otg.DeviceIsisRouter) DeviceIsisRouter {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceIsisRouter struct {
	obj *deviceIsisRouter
}

type marshalDeviceIsisRouter interface {
	// ToProto marshals DeviceIsisRouter to protobuf object *otg.DeviceIsisRouter
	ToProto() (*otg.DeviceIsisRouter, error)
	// ToPbText marshals DeviceIsisRouter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceIsisRouter to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceIsisRouter to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceIsisRouter struct {
	obj *deviceIsisRouter
}

type unMarshalDeviceIsisRouter interface {
	// FromProto unmarshals DeviceIsisRouter from protobuf object *otg.DeviceIsisRouter
	FromProto(msg *otg.DeviceIsisRouter) (DeviceIsisRouter, error)
	// FromPbText unmarshals DeviceIsisRouter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceIsisRouter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceIsisRouter from JSON text
	FromJson(value string) error
}

func (obj *deviceIsisRouter) Marshal() marshalDeviceIsisRouter {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceIsisRouter{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceIsisRouter) Unmarshal() unMarshalDeviceIsisRouter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceIsisRouter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceIsisRouter) ToProto() (*otg.DeviceIsisRouter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceIsisRouter) FromProto(msg *otg.DeviceIsisRouter) (DeviceIsisRouter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceIsisRouter) ToPbText() (string, error) {
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

func (m *unMarshaldeviceIsisRouter) FromPbText(value string) error {
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

func (m *marshaldeviceIsisRouter) ToYaml() (string, error) {
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

func (m *unMarshaldeviceIsisRouter) FromYaml(value string) error {
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

func (m *marshaldeviceIsisRouter) ToJson() (string, error) {
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

func (m *unMarshaldeviceIsisRouter) FromJson(value string) error {
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

func (obj *deviceIsisRouter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceIsisRouter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceIsisRouter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceIsisRouter) Clone() (DeviceIsisRouter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceIsisRouter()
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

func (obj *deviceIsisRouter) setNil() {
	obj.instanceHolder = nil
	obj.interfacesHolder = nil
	obj.basicHolder = nil
	obj.advancedHolder = nil
	obj.routerAuthHolder = nil
	obj.v4RoutesHolder = nil
	obj.v6RoutesHolder = nil
	obj.segmentRoutingHolder = nil
	obj.gracefulRestartHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceIsisRouter is a container of properties for an ISIS router and its interfaces.
type DeviceIsisRouter interface {
	Validation
	// msg marshals DeviceIsisRouter to protobuf object *otg.DeviceIsisRouter
	// and doesn't set defaults
	msg() *otg.DeviceIsisRouter
	// setMsg unmarshals DeviceIsisRouter from protobuf object *otg.DeviceIsisRouter
	// and doesn't set defaults
	setMsg(*otg.DeviceIsisRouter) DeviceIsisRouter
	// provides marshal interface
	Marshal() marshalDeviceIsisRouter
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceIsisRouter
	// validate validates DeviceIsisRouter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceIsisRouter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Instance returns DeviceIsisMultiInstance, set in DeviceIsisRouter.
	// DeviceIsisMultiInstance is this container properties of an Multi-Instance-capable router (MI-RTR).
	Instance() DeviceIsisMultiInstance
	// SetInstance assigns DeviceIsisMultiInstance provided by user to DeviceIsisRouter.
	// DeviceIsisMultiInstance is this container properties of an Multi-Instance-capable router (MI-RTR).
	SetInstance(value DeviceIsisMultiInstance) DeviceIsisRouter
	// HasInstance checks if Instance has been set in DeviceIsisRouter
	HasInstance() bool
	// SystemId returns string, set in DeviceIsisRouter.
	SystemId() string
	// SetSystemId assigns string provided by user to DeviceIsisRouter
	SetSystemId(value string) DeviceIsisRouter
	// Interfaces returns DeviceIsisRouterIsisInterfaceIterIter, set in DeviceIsisRouter
	Interfaces() DeviceIsisRouterIsisInterfaceIter
	// Basic returns IsisBasic, set in DeviceIsisRouter.
	// IsisBasic is this contains ISIS router basic properties.
	Basic() IsisBasic
	// SetBasic assigns IsisBasic provided by user to DeviceIsisRouter.
	// IsisBasic is this contains ISIS router basic properties.
	SetBasic(value IsisBasic) DeviceIsisRouter
	// HasBasic checks if Basic has been set in DeviceIsisRouter
	HasBasic() bool
	// Advanced returns IsisAdvanced, set in DeviceIsisRouter.
	// IsisAdvanced is contains ISIS router advanced properties.
	Advanced() IsisAdvanced
	// SetAdvanced assigns IsisAdvanced provided by user to DeviceIsisRouter.
	// IsisAdvanced is contains ISIS router advanced properties.
	SetAdvanced(value IsisAdvanced) DeviceIsisRouter
	// HasAdvanced checks if Advanced has been set in DeviceIsisRouter
	HasAdvanced() bool
	// RouterAuth returns IsisAuthentication, set in DeviceIsisRouter.
	// IsisAuthentication is this contains ISIS Area/Domain authentication properties.
	RouterAuth() IsisAuthentication
	// SetRouterAuth assigns IsisAuthentication provided by user to DeviceIsisRouter.
	// IsisAuthentication is this contains ISIS Area/Domain authentication properties.
	SetRouterAuth(value IsisAuthentication) DeviceIsisRouter
	// HasRouterAuth checks if RouterAuth has been set in DeviceIsisRouter
	HasRouterAuth() bool
	// V4Routes returns DeviceIsisRouterIsisV4RouteRangeIterIter, set in DeviceIsisRouter
	V4Routes() DeviceIsisRouterIsisV4RouteRangeIter
	// V6Routes returns DeviceIsisRouterIsisV6RouteRangeIterIter, set in DeviceIsisRouter
	V6Routes() DeviceIsisRouterIsisV6RouteRangeIter
	// Name returns string, set in DeviceIsisRouter.
	Name() string
	// SetName assigns string provided by user to DeviceIsisRouter
	SetName(value string) DeviceIsisRouter
	// SegmentRouting returns IsisSegmentRouting, set in DeviceIsisRouter.
	// IsisSegmentRouting is segment Routing (SR) allows for a flexible definition of end-to-end paths within IGP topologies by encoding paths as sequences of topological sub-paths,
	// called "segments". These segments are advertised by the link-state routing protocols (IS-IS and OSPF).
	// Prefix segments represent an ECMP-aware shortest path to a prefix (or a node), as per the state of the IGP topology.
	// Adjacency segments represent a hop over a specific adjacency between two nodes in the IGP.
	// A prefix segment is typically a multi-hop path while an adjacency segment, in most of the cases, is a one-hop path.
	// These segments act as topological sub-paths that can be combined together to form the required path.
	// Reference: https://datatracker.ietf.org/doc/html/rfc8667.:w
	// An implementation may advertise Router Capability with default values if a user does not even set the properties
	// of Router Capability and Segment Routing Capability.
	SegmentRouting() IsisSegmentRouting
	// SetSegmentRouting assigns IsisSegmentRouting provided by user to DeviceIsisRouter.
	// IsisSegmentRouting is segment Routing (SR) allows for a flexible definition of end-to-end paths within IGP topologies by encoding paths as sequences of topological sub-paths,
	// called "segments". These segments are advertised by the link-state routing protocols (IS-IS and OSPF).
	// Prefix segments represent an ECMP-aware shortest path to a prefix (or a node), as per the state of the IGP topology.
	// Adjacency segments represent a hop over a specific adjacency between two nodes in the IGP.
	// A prefix segment is typically a multi-hop path while an adjacency segment, in most of the cases, is a one-hop path.
	// These segments act as topological sub-paths that can be combined together to form the required path.
	// Reference: https://datatracker.ietf.org/doc/html/rfc8667.:w
	// An implementation may advertise Router Capability with default values if a user does not even set the properties
	// of Router Capability and Segment Routing Capability.
	SetSegmentRouting(value IsisSegmentRouting) DeviceIsisRouter
	// HasSegmentRouting checks if SegmentRouting has been set in DeviceIsisRouter
	HasSegmentRouting() bool
	// GracefulRestart returns IsisGracefulRestart, set in DeviceIsisRouter.
	// IsisGracefulRestart is contains IS-IS Graceful configuration parameters. https://datatracker.ietf.org/doc/html/rfc5306
	GracefulRestart() IsisGracefulRestart
	// SetGracefulRestart assigns IsisGracefulRestart provided by user to DeviceIsisRouter.
	// IsisGracefulRestart is contains IS-IS Graceful configuration parameters. https://datatracker.ietf.org/doc/html/rfc5306
	SetGracefulRestart(value IsisGracefulRestart) DeviceIsisRouter
	// HasGracefulRestart checks if GracefulRestart has been set in DeviceIsisRouter
	HasGracefulRestart() bool
	setNil()
}

// This contains the properties of a Multi-Instance-capable routers or MI-RTR. Each router can emulate one ISIS instance at a time.
// Instance returns a DeviceIsisMultiInstance
func (obj *deviceIsisRouter) Instance() DeviceIsisMultiInstance {
	if obj.obj.Instance == nil {
		obj.obj.Instance = NewDeviceIsisMultiInstance().msg()
	}
	if obj.instanceHolder == nil {
		obj.instanceHolder = &deviceIsisMultiInstance{obj: obj.obj.Instance}
	}
	return obj.instanceHolder
}

// This contains the properties of a Multi-Instance-capable routers or MI-RTR. Each router can emulate one ISIS instance at a time.
// Instance returns a DeviceIsisMultiInstance
func (obj *deviceIsisRouter) HasInstance() bool {
	return obj.obj.Instance != nil
}

// This contains the properties of a Multi-Instance-capable routers or MI-RTR. Each router can emulate one ISIS instance at a time.
// SetInstance sets the DeviceIsisMultiInstance value in the DeviceIsisRouter object
func (obj *deviceIsisRouter) SetInstance(value DeviceIsisMultiInstance) DeviceIsisRouter {

	obj.instanceHolder = nil
	obj.obj.Instance = value.msg()

	return obj
}

// The System ID for this emulated ISIS router, e.g. "640100010000".
// SystemId returns a string
func (obj *deviceIsisRouter) SystemId() string {

	return *obj.obj.SystemId

}

// The System ID for this emulated ISIS router, e.g. "640100010000".
// SetSystemId sets the string value in the DeviceIsisRouter object
func (obj *deviceIsisRouter) SetSystemId(value string) DeviceIsisRouter {

	obj.obj.SystemId = &value
	return obj
}

// List of ISIS interfaces for this router.
// Interfaces returns a []IsisInterface
func (obj *deviceIsisRouter) Interfaces() DeviceIsisRouterIsisInterfaceIter {
	if len(obj.obj.Interfaces) == 0 {
		obj.obj.Interfaces = []*otg.IsisInterface{}
	}
	if obj.interfacesHolder == nil {
		obj.interfacesHolder = newDeviceIsisRouterIsisInterfaceIter(&obj.obj.Interfaces).setMsg(obj)
	}
	return obj.interfacesHolder
}

type deviceIsisRouterIsisInterfaceIter struct {
	obj                *deviceIsisRouter
	isisInterfaceSlice []IsisInterface
	fieldPtr           *[]*otg.IsisInterface
}

func newDeviceIsisRouterIsisInterfaceIter(ptr *[]*otg.IsisInterface) DeviceIsisRouterIsisInterfaceIter {
	return &deviceIsisRouterIsisInterfaceIter{fieldPtr: ptr}
}

type DeviceIsisRouterIsisInterfaceIter interface {
	setMsg(*deviceIsisRouter) DeviceIsisRouterIsisInterfaceIter
	Items() []IsisInterface
	Add() IsisInterface
	Append(items ...IsisInterface) DeviceIsisRouterIsisInterfaceIter
	Set(index int, newObj IsisInterface) DeviceIsisRouterIsisInterfaceIter
	Clear() DeviceIsisRouterIsisInterfaceIter
	clearHolderSlice() DeviceIsisRouterIsisInterfaceIter
	appendHolderSlice(item IsisInterface) DeviceIsisRouterIsisInterfaceIter
}

func (obj *deviceIsisRouterIsisInterfaceIter) setMsg(msg *deviceIsisRouter) DeviceIsisRouterIsisInterfaceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisInterface{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceIsisRouterIsisInterfaceIter) Items() []IsisInterface {
	return obj.isisInterfaceSlice
}

func (obj *deviceIsisRouterIsisInterfaceIter) Add() IsisInterface {
	newObj := &otg.IsisInterface{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisInterface{obj: newObj}
	newLibObj.setDefault()
	obj.isisInterfaceSlice = append(obj.isisInterfaceSlice, newLibObj)
	return newLibObj
}

func (obj *deviceIsisRouterIsisInterfaceIter) Append(items ...IsisInterface) DeviceIsisRouterIsisInterfaceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisInterfaceSlice = append(obj.isisInterfaceSlice, item)
	}
	return obj
}

func (obj *deviceIsisRouterIsisInterfaceIter) Set(index int, newObj IsisInterface) DeviceIsisRouterIsisInterfaceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisInterfaceSlice[index] = newObj
	return obj
}
func (obj *deviceIsisRouterIsisInterfaceIter) Clear() DeviceIsisRouterIsisInterfaceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisInterface{}
		obj.isisInterfaceSlice = []IsisInterface{}
	}
	return obj
}
func (obj *deviceIsisRouterIsisInterfaceIter) clearHolderSlice() DeviceIsisRouterIsisInterfaceIter {
	if len(obj.isisInterfaceSlice) > 0 {
		obj.isisInterfaceSlice = []IsisInterface{}
	}
	return obj
}
func (obj *deviceIsisRouterIsisInterfaceIter) appendHolderSlice(item IsisInterface) DeviceIsisRouterIsisInterfaceIter {
	obj.isisInterfaceSlice = append(obj.isisInterfaceSlice, item)
	return obj
}

// Contains basic properties of an ISIS Router.
// Basic returns a IsisBasic
func (obj *deviceIsisRouter) Basic() IsisBasic {
	if obj.obj.Basic == nil {
		obj.obj.Basic = NewIsisBasic().msg()
	}
	if obj.basicHolder == nil {
		obj.basicHolder = &isisBasic{obj: obj.obj.Basic}
	}
	return obj.basicHolder
}

// Contains basic properties of an ISIS Router.
// Basic returns a IsisBasic
func (obj *deviceIsisRouter) HasBasic() bool {
	return obj.obj.Basic != nil
}

// Contains basic properties of an ISIS Router.
// SetBasic sets the IsisBasic value in the DeviceIsisRouter object
func (obj *deviceIsisRouter) SetBasic(value IsisBasic) DeviceIsisRouter {

	obj.basicHolder = nil
	obj.obj.Basic = value.msg()

	return obj
}

// Contains advance properties of an ISIS Router..
// Advanced returns a IsisAdvanced
func (obj *deviceIsisRouter) Advanced() IsisAdvanced {
	if obj.obj.Advanced == nil {
		obj.obj.Advanced = NewIsisAdvanced().msg()
	}
	if obj.advancedHolder == nil {
		obj.advancedHolder = &isisAdvanced{obj: obj.obj.Advanced}
	}
	return obj.advancedHolder
}

// Contains advance properties of an ISIS Router..
// Advanced returns a IsisAdvanced
func (obj *deviceIsisRouter) HasAdvanced() bool {
	return obj.obj.Advanced != nil
}

// Contains advance properties of an ISIS Router..
// SetAdvanced sets the IsisAdvanced value in the DeviceIsisRouter object
func (obj *deviceIsisRouter) SetAdvanced(value IsisAdvanced) DeviceIsisRouter {

	obj.advancedHolder = nil
	obj.obj.Advanced = value.msg()

	return obj
}

// ISIS Router authentication properties.
// RouterAuth returns a IsisAuthentication
func (obj *deviceIsisRouter) RouterAuth() IsisAuthentication {
	if obj.obj.RouterAuth == nil {
		obj.obj.RouterAuth = NewIsisAuthentication().msg()
	}
	if obj.routerAuthHolder == nil {
		obj.routerAuthHolder = &isisAuthentication{obj: obj.obj.RouterAuth}
	}
	return obj.routerAuthHolder
}

// ISIS Router authentication properties.
// RouterAuth returns a IsisAuthentication
func (obj *deviceIsisRouter) HasRouterAuth() bool {
	return obj.obj.RouterAuth != nil
}

// ISIS Router authentication properties.
// SetRouterAuth sets the IsisAuthentication value in the DeviceIsisRouter object
func (obj *deviceIsisRouter) SetRouterAuth(value IsisAuthentication) DeviceIsisRouter {

	obj.routerAuthHolder = nil
	obj.obj.RouterAuth = value.msg()

	return obj
}

// Emulated ISIS IPv4 routes.
// V4Routes returns a []IsisV4RouteRange
func (obj *deviceIsisRouter) V4Routes() DeviceIsisRouterIsisV4RouteRangeIter {
	if len(obj.obj.V4Routes) == 0 {
		obj.obj.V4Routes = []*otg.IsisV4RouteRange{}
	}
	if obj.v4RoutesHolder == nil {
		obj.v4RoutesHolder = newDeviceIsisRouterIsisV4RouteRangeIter(&obj.obj.V4Routes).setMsg(obj)
	}
	return obj.v4RoutesHolder
}

type deviceIsisRouterIsisV4RouteRangeIter struct {
	obj                   *deviceIsisRouter
	isisV4RouteRangeSlice []IsisV4RouteRange
	fieldPtr              *[]*otg.IsisV4RouteRange
}

func newDeviceIsisRouterIsisV4RouteRangeIter(ptr *[]*otg.IsisV4RouteRange) DeviceIsisRouterIsisV4RouteRangeIter {
	return &deviceIsisRouterIsisV4RouteRangeIter{fieldPtr: ptr}
}

type DeviceIsisRouterIsisV4RouteRangeIter interface {
	setMsg(*deviceIsisRouter) DeviceIsisRouterIsisV4RouteRangeIter
	Items() []IsisV4RouteRange
	Add() IsisV4RouteRange
	Append(items ...IsisV4RouteRange) DeviceIsisRouterIsisV4RouteRangeIter
	Set(index int, newObj IsisV4RouteRange) DeviceIsisRouterIsisV4RouteRangeIter
	Clear() DeviceIsisRouterIsisV4RouteRangeIter
	clearHolderSlice() DeviceIsisRouterIsisV4RouteRangeIter
	appendHolderSlice(item IsisV4RouteRange) DeviceIsisRouterIsisV4RouteRangeIter
}

func (obj *deviceIsisRouterIsisV4RouteRangeIter) setMsg(msg *deviceIsisRouter) DeviceIsisRouterIsisV4RouteRangeIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisV4RouteRange{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceIsisRouterIsisV4RouteRangeIter) Items() []IsisV4RouteRange {
	return obj.isisV4RouteRangeSlice
}

func (obj *deviceIsisRouterIsisV4RouteRangeIter) Add() IsisV4RouteRange {
	newObj := &otg.IsisV4RouteRange{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisV4RouteRange{obj: newObj}
	newLibObj.setDefault()
	obj.isisV4RouteRangeSlice = append(obj.isisV4RouteRangeSlice, newLibObj)
	return newLibObj
}

func (obj *deviceIsisRouterIsisV4RouteRangeIter) Append(items ...IsisV4RouteRange) DeviceIsisRouterIsisV4RouteRangeIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisV4RouteRangeSlice = append(obj.isisV4RouteRangeSlice, item)
	}
	return obj
}

func (obj *deviceIsisRouterIsisV4RouteRangeIter) Set(index int, newObj IsisV4RouteRange) DeviceIsisRouterIsisV4RouteRangeIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisV4RouteRangeSlice[index] = newObj
	return obj
}
func (obj *deviceIsisRouterIsisV4RouteRangeIter) Clear() DeviceIsisRouterIsisV4RouteRangeIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisV4RouteRange{}
		obj.isisV4RouteRangeSlice = []IsisV4RouteRange{}
	}
	return obj
}
func (obj *deviceIsisRouterIsisV4RouteRangeIter) clearHolderSlice() DeviceIsisRouterIsisV4RouteRangeIter {
	if len(obj.isisV4RouteRangeSlice) > 0 {
		obj.isisV4RouteRangeSlice = []IsisV4RouteRange{}
	}
	return obj
}
func (obj *deviceIsisRouterIsisV4RouteRangeIter) appendHolderSlice(item IsisV4RouteRange) DeviceIsisRouterIsisV4RouteRangeIter {
	obj.isisV4RouteRangeSlice = append(obj.isisV4RouteRangeSlice, item)
	return obj
}

// Emulated ISIS IPv6 routes.
// V6Routes returns a []IsisV6RouteRange
func (obj *deviceIsisRouter) V6Routes() DeviceIsisRouterIsisV6RouteRangeIter {
	if len(obj.obj.V6Routes) == 0 {
		obj.obj.V6Routes = []*otg.IsisV6RouteRange{}
	}
	if obj.v6RoutesHolder == nil {
		obj.v6RoutesHolder = newDeviceIsisRouterIsisV6RouteRangeIter(&obj.obj.V6Routes).setMsg(obj)
	}
	return obj.v6RoutesHolder
}

type deviceIsisRouterIsisV6RouteRangeIter struct {
	obj                   *deviceIsisRouter
	isisV6RouteRangeSlice []IsisV6RouteRange
	fieldPtr              *[]*otg.IsisV6RouteRange
}

func newDeviceIsisRouterIsisV6RouteRangeIter(ptr *[]*otg.IsisV6RouteRange) DeviceIsisRouterIsisV6RouteRangeIter {
	return &deviceIsisRouterIsisV6RouteRangeIter{fieldPtr: ptr}
}

type DeviceIsisRouterIsisV6RouteRangeIter interface {
	setMsg(*deviceIsisRouter) DeviceIsisRouterIsisV6RouteRangeIter
	Items() []IsisV6RouteRange
	Add() IsisV6RouteRange
	Append(items ...IsisV6RouteRange) DeviceIsisRouterIsisV6RouteRangeIter
	Set(index int, newObj IsisV6RouteRange) DeviceIsisRouterIsisV6RouteRangeIter
	Clear() DeviceIsisRouterIsisV6RouteRangeIter
	clearHolderSlice() DeviceIsisRouterIsisV6RouteRangeIter
	appendHolderSlice(item IsisV6RouteRange) DeviceIsisRouterIsisV6RouteRangeIter
}

func (obj *deviceIsisRouterIsisV6RouteRangeIter) setMsg(msg *deviceIsisRouter) DeviceIsisRouterIsisV6RouteRangeIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisV6RouteRange{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceIsisRouterIsisV6RouteRangeIter) Items() []IsisV6RouteRange {
	return obj.isisV6RouteRangeSlice
}

func (obj *deviceIsisRouterIsisV6RouteRangeIter) Add() IsisV6RouteRange {
	newObj := &otg.IsisV6RouteRange{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisV6RouteRange{obj: newObj}
	newLibObj.setDefault()
	obj.isisV6RouteRangeSlice = append(obj.isisV6RouteRangeSlice, newLibObj)
	return newLibObj
}

func (obj *deviceIsisRouterIsisV6RouteRangeIter) Append(items ...IsisV6RouteRange) DeviceIsisRouterIsisV6RouteRangeIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisV6RouteRangeSlice = append(obj.isisV6RouteRangeSlice, item)
	}
	return obj
}

func (obj *deviceIsisRouterIsisV6RouteRangeIter) Set(index int, newObj IsisV6RouteRange) DeviceIsisRouterIsisV6RouteRangeIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisV6RouteRangeSlice[index] = newObj
	return obj
}
func (obj *deviceIsisRouterIsisV6RouteRangeIter) Clear() DeviceIsisRouterIsisV6RouteRangeIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisV6RouteRange{}
		obj.isisV6RouteRangeSlice = []IsisV6RouteRange{}
	}
	return obj
}
func (obj *deviceIsisRouterIsisV6RouteRangeIter) clearHolderSlice() DeviceIsisRouterIsisV6RouteRangeIter {
	if len(obj.isisV6RouteRangeSlice) > 0 {
		obj.isisV6RouteRangeSlice = []IsisV6RouteRange{}
	}
	return obj
}
func (obj *deviceIsisRouterIsisV6RouteRangeIter) appendHolderSlice(item IsisV6RouteRange) DeviceIsisRouterIsisV6RouteRangeIter {
	obj.isisV6RouteRangeSlice = append(obj.isisV6RouteRangeSlice, item)
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *deviceIsisRouter) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the DeviceIsisRouter object
func (obj *deviceIsisRouter) SetName(value string) DeviceIsisRouter {

	obj.obj.Name = &value
	return obj
}

// Optional Segment Routing (SR).
// SegmentRouting returns a IsisSegmentRouting
func (obj *deviceIsisRouter) SegmentRouting() IsisSegmentRouting {
	if obj.obj.SegmentRouting == nil {
		obj.obj.SegmentRouting = NewIsisSegmentRouting().msg()
	}
	if obj.segmentRoutingHolder == nil {
		obj.segmentRoutingHolder = &isisSegmentRouting{obj: obj.obj.SegmentRouting}
	}
	return obj.segmentRoutingHolder
}

// Optional Segment Routing (SR).
// SegmentRouting returns a IsisSegmentRouting
func (obj *deviceIsisRouter) HasSegmentRouting() bool {
	return obj.obj.SegmentRouting != nil
}

// Optional Segment Routing (SR).
// SetSegmentRouting sets the IsisSegmentRouting value in the DeviceIsisRouter object
func (obj *deviceIsisRouter) SetSegmentRouting(value IsisSegmentRouting) DeviceIsisRouter {

	obj.segmentRoutingHolder = nil
	obj.obj.SegmentRouting = value.msg()

	return obj
}

// Optional IS-IS Graceful Restart Configuration.
// GracefulRestart returns a IsisGracefulRestart
func (obj *deviceIsisRouter) GracefulRestart() IsisGracefulRestart {
	if obj.obj.GracefulRestart == nil {
		obj.obj.GracefulRestart = NewIsisGracefulRestart().msg()
	}
	if obj.gracefulRestartHolder == nil {
		obj.gracefulRestartHolder = &isisGracefulRestart{obj: obj.obj.GracefulRestart}
	}
	return obj.gracefulRestartHolder
}

// Optional IS-IS Graceful Restart Configuration.
// GracefulRestart returns a IsisGracefulRestart
func (obj *deviceIsisRouter) HasGracefulRestart() bool {
	return obj.obj.GracefulRestart != nil
}

// Optional IS-IS Graceful Restart Configuration.
// SetGracefulRestart sets the IsisGracefulRestart value in the DeviceIsisRouter object
func (obj *deviceIsisRouter) SetGracefulRestart(value IsisGracefulRestart) DeviceIsisRouter {

	obj.gracefulRestartHolder = nil
	obj.obj.GracefulRestart = value.msg()

	return obj
}

func (obj *deviceIsisRouter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Instance != nil {

		obj.Instance().validateObj(vObj, set_default)
	}

	// SystemId is required
	if obj.obj.SystemId == nil {
		vObj.validationErrors = append(vObj.validationErrors, "SystemId is required field on interface DeviceIsisRouter")
	}
	if obj.obj.SystemId != nil {

		err := obj.validateHex(obj.SystemId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIsisRouter.SystemId"))
		}

	}

	if len(obj.obj.Interfaces) != 0 {

		if set_default {
			obj.Interfaces().clearHolderSlice()
			for _, item := range obj.obj.Interfaces {
				obj.Interfaces().appendHolderSlice(&isisInterface{obj: item})
			}
		}
		for _, item := range obj.Interfaces().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Basic != nil {

		obj.Basic().validateObj(vObj, set_default)
	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(vObj, set_default)
	}

	if obj.obj.RouterAuth != nil {

		obj.RouterAuth().validateObj(vObj, set_default)
	}

	if len(obj.obj.V4Routes) != 0 {

		if set_default {
			obj.V4Routes().clearHolderSlice()
			for _, item := range obj.obj.V4Routes {
				obj.V4Routes().appendHolderSlice(&isisV4RouteRange{obj: item})
			}
		}
		for _, item := range obj.V4Routes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.V6Routes) != 0 {

		if set_default {
			obj.V6Routes().clearHolderSlice()
			for _, item := range obj.obj.V6Routes {
				obj.V6Routes().appendHolderSlice(&isisV6RouteRange{obj: item})
			}
		}
		for _, item := range obj.V6Routes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface DeviceIsisRouter")
	}

	if obj.obj.SegmentRouting != nil {

		obj.SegmentRouting().validateObj(vObj, set_default)
	}

	if obj.obj.GracefulRestart != nil {

		obj.GracefulRestart().validateObj(vObj, set_default)
	}

}

func (obj *deviceIsisRouter) setDefault() {

}

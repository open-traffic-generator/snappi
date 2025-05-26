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
	obj             *otg.DeviceOspfv3Router
	marshaller      marshalDeviceOspfv3Router
	unMarshaller    unMarshalDeviceOspfv3Router
	routerIdHolder  Ospfv3RouterId
	instancesHolder DeviceOspfv3RouterOspfv3RouterInstanceIter
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
	obj.instancesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceOspfv3Router is under Review: OSPFv3 is currently under review for pending exploration on use cases.
//
// Under Review: OSPFv3 is currently under review for pending exploration on use cases.
//
// A container of properties for an OSPFv3 router.
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
	// RouterId returns Ospfv3RouterId, set in DeviceOspfv3Router.
	// Ospfv3RouterId is container for OSPFv3 Router ID configuration.
	RouterId() Ospfv3RouterId
	// SetRouterId assigns Ospfv3RouterId provided by user to DeviceOspfv3Router.
	// Ospfv3RouterId is container for OSPFv3 Router ID configuration.
	SetRouterId(value Ospfv3RouterId) DeviceOspfv3Router
	// HasRouterId checks if RouterId has been set in DeviceOspfv3Router
	HasRouterId() bool
	// Instances returns DeviceOspfv3RouterOspfv3RouterInstanceIterIter, set in DeviceOspfv3Router
	Instances() DeviceOspfv3RouterOspfv3RouterInstanceIter
	setNil()
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

// List of OSPFv3 router instances for this router.
// Instances returns a []Ospfv3RouterInstance
func (obj *deviceOspfv3Router) Instances() DeviceOspfv3RouterOspfv3RouterInstanceIter {
	if len(obj.obj.Instances) == 0 {
		obj.obj.Instances = []*otg.Ospfv3RouterInstance{}
	}
	if obj.instancesHolder == nil {
		obj.instancesHolder = newDeviceOspfv3RouterOspfv3RouterInstanceIter(&obj.obj.Instances).setMsg(obj)
	}
	return obj.instancesHolder
}

type deviceOspfv3RouterOspfv3RouterInstanceIter struct {
	obj                       *deviceOspfv3Router
	ospfv3RouterInstanceSlice []Ospfv3RouterInstance
	fieldPtr                  *[]*otg.Ospfv3RouterInstance
}

func newDeviceOspfv3RouterOspfv3RouterInstanceIter(ptr *[]*otg.Ospfv3RouterInstance) DeviceOspfv3RouterOspfv3RouterInstanceIter {
	return &deviceOspfv3RouterOspfv3RouterInstanceIter{fieldPtr: ptr}
}

type DeviceOspfv3RouterOspfv3RouterInstanceIter interface {
	setMsg(*deviceOspfv3Router) DeviceOspfv3RouterOspfv3RouterInstanceIter
	Items() []Ospfv3RouterInstance
	Add() Ospfv3RouterInstance
	Append(items ...Ospfv3RouterInstance) DeviceOspfv3RouterOspfv3RouterInstanceIter
	Set(index int, newObj Ospfv3RouterInstance) DeviceOspfv3RouterOspfv3RouterInstanceIter
	Clear() DeviceOspfv3RouterOspfv3RouterInstanceIter
	clearHolderSlice() DeviceOspfv3RouterOspfv3RouterInstanceIter
	appendHolderSlice(item Ospfv3RouterInstance) DeviceOspfv3RouterOspfv3RouterInstanceIter
}

func (obj *deviceOspfv3RouterOspfv3RouterInstanceIter) setMsg(msg *deviceOspfv3Router) DeviceOspfv3RouterOspfv3RouterInstanceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv3RouterInstance{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceOspfv3RouterOspfv3RouterInstanceIter) Items() []Ospfv3RouterInstance {
	return obj.ospfv3RouterInstanceSlice
}

func (obj *deviceOspfv3RouterOspfv3RouterInstanceIter) Add() Ospfv3RouterInstance {
	newObj := &otg.Ospfv3RouterInstance{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv3RouterInstance{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv3RouterInstanceSlice = append(obj.ospfv3RouterInstanceSlice, newLibObj)
	return newLibObj
}

func (obj *deviceOspfv3RouterOspfv3RouterInstanceIter) Append(items ...Ospfv3RouterInstance) DeviceOspfv3RouterOspfv3RouterInstanceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv3RouterInstanceSlice = append(obj.ospfv3RouterInstanceSlice, item)
	}
	return obj
}

func (obj *deviceOspfv3RouterOspfv3RouterInstanceIter) Set(index int, newObj Ospfv3RouterInstance) DeviceOspfv3RouterOspfv3RouterInstanceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv3RouterInstanceSlice[index] = newObj
	return obj
}
func (obj *deviceOspfv3RouterOspfv3RouterInstanceIter) Clear() DeviceOspfv3RouterOspfv3RouterInstanceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv3RouterInstance{}
		obj.ospfv3RouterInstanceSlice = []Ospfv3RouterInstance{}
	}
	return obj
}
func (obj *deviceOspfv3RouterOspfv3RouterInstanceIter) clearHolderSlice() DeviceOspfv3RouterOspfv3RouterInstanceIter {
	if len(obj.ospfv3RouterInstanceSlice) > 0 {
		obj.ospfv3RouterInstanceSlice = []Ospfv3RouterInstance{}
	}
	return obj
}
func (obj *deviceOspfv3RouterOspfv3RouterInstanceIter) appendHolderSlice(item Ospfv3RouterInstance) DeviceOspfv3RouterOspfv3RouterInstanceIter {
	obj.ospfv3RouterInstanceSlice = append(obj.ospfv3RouterInstanceSlice, item)
	return obj
}

func (obj *deviceOspfv3Router) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	obj.addWarnings("DeviceOspfv3Router is under review, OSPFv3 is currently under review for pending exploration on use cases.")

	if obj.obj.RouterId != nil {

		obj.RouterId().validateObj(vObj, set_default)
	}

	if len(obj.obj.Instances) != 0 {

		if set_default {
			obj.Instances().clearHolderSlice()
			for _, item := range obj.obj.Instances {
				obj.Instances().appendHolderSlice(&ospfv3RouterInstance{obj: item})
			}
		}
		for _, item := range obj.Instances().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *deviceOspfv3Router) setDefault() {

}

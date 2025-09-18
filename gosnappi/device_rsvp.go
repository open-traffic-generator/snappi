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

// ***** DeviceRsvp *****
type deviceRsvp struct {
	validation
	obj                     *otg.DeviceRsvp
	marshaller              marshalDeviceRsvp
	unMarshaller            unMarshalDeviceRsvp
	ipv4InterfacesHolder    DeviceRsvpRsvpIpv4InterfaceIter
	lspIpv4InterfacesHolder DeviceRsvpRsvpLspIpv4InterfaceIter
}

func NewDeviceRsvp() DeviceRsvp {
	obj := deviceRsvp{obj: &otg.DeviceRsvp{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceRsvp) msg() *otg.DeviceRsvp {
	return obj.obj
}

func (obj *deviceRsvp) setMsg(msg *otg.DeviceRsvp) DeviceRsvp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceRsvp struct {
	obj *deviceRsvp
}

type marshalDeviceRsvp interface {
	// ToProto marshals DeviceRsvp to protobuf object *otg.DeviceRsvp
	ToProto() (*otg.DeviceRsvp, error)
	// ToPbText marshals DeviceRsvp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceRsvp to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceRsvp to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceRsvp struct {
	obj *deviceRsvp
}

type unMarshalDeviceRsvp interface {
	// FromProto unmarshals DeviceRsvp from protobuf object *otg.DeviceRsvp
	FromProto(msg *otg.DeviceRsvp) (DeviceRsvp, error)
	// FromPbText unmarshals DeviceRsvp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceRsvp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceRsvp from JSON text
	FromJson(value string) error
}

func (obj *deviceRsvp) Marshal() marshalDeviceRsvp {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceRsvp{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceRsvp) Unmarshal() unMarshalDeviceRsvp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceRsvp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceRsvp) ToProto() (*otg.DeviceRsvp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceRsvp) FromProto(msg *otg.DeviceRsvp) (DeviceRsvp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceRsvp) ToPbText() (string, error) {
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

func (m *unMarshaldeviceRsvp) FromPbText(value string) error {
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

func (m *marshaldeviceRsvp) ToYaml() (string, error) {
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

func (m *unMarshaldeviceRsvp) FromYaml(value string) error {
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

func (m *marshaldeviceRsvp) ToJson() (string, error) {
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

func (m *unMarshaldeviceRsvp) FromJson(value string) error {
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

func (obj *deviceRsvp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceRsvp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceRsvp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceRsvp) Clone() (DeviceRsvp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceRsvp()
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

func (obj *deviceRsvp) setNil() {
	obj.ipv4InterfacesHolder = nil
	obj.lspIpv4InterfacesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceRsvp is configuration for one or more RSVP interfaces, ingress and egress LSPs. In this model, currently IPv4 RSVP and point-to-point LSPs are supported as per RFC3209 and related specifications.
type DeviceRsvp interface {
	Validation
	// msg marshals DeviceRsvp to protobuf object *otg.DeviceRsvp
	// and doesn't set defaults
	msg() *otg.DeviceRsvp
	// setMsg unmarshals DeviceRsvp from protobuf object *otg.DeviceRsvp
	// and doesn't set defaults
	setMsg(*otg.DeviceRsvp) DeviceRsvp
	// provides marshal interface
	Marshal() marshalDeviceRsvp
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceRsvp
	// validate validates DeviceRsvp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceRsvp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Interfaces returns DeviceRsvpRsvpIpv4InterfaceIterIter, set in DeviceRsvp
	Ipv4Interfaces() DeviceRsvpRsvpIpv4InterfaceIter
	// LspIpv4Interfaces returns DeviceRsvpRsvpLspIpv4InterfaceIterIter, set in DeviceRsvp
	LspIpv4Interfaces() DeviceRsvpRsvpLspIpv4InterfaceIter
	// Name returns string, set in DeviceRsvp.
	Name() string
	// SetName assigns string provided by user to DeviceRsvp
	SetName(value string) DeviceRsvp
	// HasName checks if Name has been set in DeviceRsvp
	HasName() bool
	setNil()
}

// List of IPv4 RSVP connected interfaces. At least one interface should be present for device connected to the DUT. For unconnected devices, this array must be empty.
// Ipv4Interfaces returns a []RsvpIpv4Interface
func (obj *deviceRsvp) Ipv4Interfaces() DeviceRsvpRsvpIpv4InterfaceIter {
	if len(obj.obj.Ipv4Interfaces) == 0 {
		obj.obj.Ipv4Interfaces = []*otg.RsvpIpv4Interface{}
	}
	if obj.ipv4InterfacesHolder == nil {
		obj.ipv4InterfacesHolder = newDeviceRsvpRsvpIpv4InterfaceIter(&obj.obj.Ipv4Interfaces).setMsg(obj)
	}
	return obj.ipv4InterfacesHolder
}

type deviceRsvpRsvpIpv4InterfaceIter struct {
	obj                    *deviceRsvp
	rsvpIpv4InterfaceSlice []RsvpIpv4Interface
	fieldPtr               *[]*otg.RsvpIpv4Interface
}

func newDeviceRsvpRsvpIpv4InterfaceIter(ptr *[]*otg.RsvpIpv4Interface) DeviceRsvpRsvpIpv4InterfaceIter {
	return &deviceRsvpRsvpIpv4InterfaceIter{fieldPtr: ptr}
}

type DeviceRsvpRsvpIpv4InterfaceIter interface {
	setMsg(*deviceRsvp) DeviceRsvpRsvpIpv4InterfaceIter
	Items() []RsvpIpv4Interface
	Add() RsvpIpv4Interface
	Append(items ...RsvpIpv4Interface) DeviceRsvpRsvpIpv4InterfaceIter
	Set(index int, newObj RsvpIpv4Interface) DeviceRsvpRsvpIpv4InterfaceIter
	Clear() DeviceRsvpRsvpIpv4InterfaceIter
	clearHolderSlice() DeviceRsvpRsvpIpv4InterfaceIter
	appendHolderSlice(item RsvpIpv4Interface) DeviceRsvpRsvpIpv4InterfaceIter
}

func (obj *deviceRsvpRsvpIpv4InterfaceIter) setMsg(msg *deviceRsvp) DeviceRsvpRsvpIpv4InterfaceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&rsvpIpv4Interface{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceRsvpRsvpIpv4InterfaceIter) Items() []RsvpIpv4Interface {
	return obj.rsvpIpv4InterfaceSlice
}

func (obj *deviceRsvpRsvpIpv4InterfaceIter) Add() RsvpIpv4Interface {
	newObj := &otg.RsvpIpv4Interface{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &rsvpIpv4Interface{obj: newObj}
	newLibObj.setDefault()
	obj.rsvpIpv4InterfaceSlice = append(obj.rsvpIpv4InterfaceSlice, newLibObj)
	return newLibObj
}

func (obj *deviceRsvpRsvpIpv4InterfaceIter) Append(items ...RsvpIpv4Interface) DeviceRsvpRsvpIpv4InterfaceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.rsvpIpv4InterfaceSlice = append(obj.rsvpIpv4InterfaceSlice, item)
	}
	return obj
}

func (obj *deviceRsvpRsvpIpv4InterfaceIter) Set(index int, newObj RsvpIpv4Interface) DeviceRsvpRsvpIpv4InterfaceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.rsvpIpv4InterfaceSlice[index] = newObj
	return obj
}
func (obj *deviceRsvpRsvpIpv4InterfaceIter) Clear() DeviceRsvpRsvpIpv4InterfaceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.RsvpIpv4Interface{}
		obj.rsvpIpv4InterfaceSlice = []RsvpIpv4Interface{}
	}
	return obj
}
func (obj *deviceRsvpRsvpIpv4InterfaceIter) clearHolderSlice() DeviceRsvpRsvpIpv4InterfaceIter {
	if len(obj.rsvpIpv4InterfaceSlice) > 0 {
		obj.rsvpIpv4InterfaceSlice = []RsvpIpv4Interface{}
	}
	return obj
}
func (obj *deviceRsvpRsvpIpv4InterfaceIter) appendHolderSlice(item RsvpIpv4Interface) DeviceRsvpRsvpIpv4InterfaceIter {
	obj.rsvpIpv4InterfaceSlice = append(obj.rsvpIpv4InterfaceSlice, item)
	return obj
}

// List of IPv4 Loopback or IPv4 connected interfaces acting as RSVP ingress and egress endpoints.
// LspIpv4Interfaces returns a []RsvpLspIpv4Interface
func (obj *deviceRsvp) LspIpv4Interfaces() DeviceRsvpRsvpLspIpv4InterfaceIter {
	if len(obj.obj.LspIpv4Interfaces) == 0 {
		obj.obj.LspIpv4Interfaces = []*otg.RsvpLspIpv4Interface{}
	}
	if obj.lspIpv4InterfacesHolder == nil {
		obj.lspIpv4InterfacesHolder = newDeviceRsvpRsvpLspIpv4InterfaceIter(&obj.obj.LspIpv4Interfaces).setMsg(obj)
	}
	return obj.lspIpv4InterfacesHolder
}

type deviceRsvpRsvpLspIpv4InterfaceIter struct {
	obj                       *deviceRsvp
	rsvpLspIpv4InterfaceSlice []RsvpLspIpv4Interface
	fieldPtr                  *[]*otg.RsvpLspIpv4Interface
}

func newDeviceRsvpRsvpLspIpv4InterfaceIter(ptr *[]*otg.RsvpLspIpv4Interface) DeviceRsvpRsvpLspIpv4InterfaceIter {
	return &deviceRsvpRsvpLspIpv4InterfaceIter{fieldPtr: ptr}
}

type DeviceRsvpRsvpLspIpv4InterfaceIter interface {
	setMsg(*deviceRsvp) DeviceRsvpRsvpLspIpv4InterfaceIter
	Items() []RsvpLspIpv4Interface
	Add() RsvpLspIpv4Interface
	Append(items ...RsvpLspIpv4Interface) DeviceRsvpRsvpLspIpv4InterfaceIter
	Set(index int, newObj RsvpLspIpv4Interface) DeviceRsvpRsvpLspIpv4InterfaceIter
	Clear() DeviceRsvpRsvpLspIpv4InterfaceIter
	clearHolderSlice() DeviceRsvpRsvpLspIpv4InterfaceIter
	appendHolderSlice(item RsvpLspIpv4Interface) DeviceRsvpRsvpLspIpv4InterfaceIter
}

func (obj *deviceRsvpRsvpLspIpv4InterfaceIter) setMsg(msg *deviceRsvp) DeviceRsvpRsvpLspIpv4InterfaceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&rsvpLspIpv4Interface{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceRsvpRsvpLspIpv4InterfaceIter) Items() []RsvpLspIpv4Interface {
	return obj.rsvpLspIpv4InterfaceSlice
}

func (obj *deviceRsvpRsvpLspIpv4InterfaceIter) Add() RsvpLspIpv4Interface {
	newObj := &otg.RsvpLspIpv4Interface{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &rsvpLspIpv4Interface{obj: newObj}
	newLibObj.setDefault()
	obj.rsvpLspIpv4InterfaceSlice = append(obj.rsvpLspIpv4InterfaceSlice, newLibObj)
	return newLibObj
}

func (obj *deviceRsvpRsvpLspIpv4InterfaceIter) Append(items ...RsvpLspIpv4Interface) DeviceRsvpRsvpLspIpv4InterfaceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.rsvpLspIpv4InterfaceSlice = append(obj.rsvpLspIpv4InterfaceSlice, item)
	}
	return obj
}

func (obj *deviceRsvpRsvpLspIpv4InterfaceIter) Set(index int, newObj RsvpLspIpv4Interface) DeviceRsvpRsvpLspIpv4InterfaceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.rsvpLspIpv4InterfaceSlice[index] = newObj
	return obj
}
func (obj *deviceRsvpRsvpLspIpv4InterfaceIter) Clear() DeviceRsvpRsvpLspIpv4InterfaceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.RsvpLspIpv4Interface{}
		obj.rsvpLspIpv4InterfaceSlice = []RsvpLspIpv4Interface{}
	}
	return obj
}
func (obj *deviceRsvpRsvpLspIpv4InterfaceIter) clearHolderSlice() DeviceRsvpRsvpLspIpv4InterfaceIter {
	if len(obj.rsvpLspIpv4InterfaceSlice) > 0 {
		obj.rsvpLspIpv4InterfaceSlice = []RsvpLspIpv4Interface{}
	}
	return obj
}
func (obj *deviceRsvpRsvpLspIpv4InterfaceIter) appendHolderSlice(item RsvpLspIpv4Interface) DeviceRsvpRsvpLspIpv4InterfaceIter {
	obj.rsvpLspIpv4InterfaceSlice = append(obj.rsvpLspIpv4InterfaceSlice, item)
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *deviceRsvp) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *deviceRsvp) HasName() bool {
	return obj.obj.Name != nil
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the DeviceRsvp object
func (obj *deviceRsvp) SetName(value string) DeviceRsvp {

	obj.obj.Name = &value
	return obj
}

func (obj *deviceRsvp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Ipv4Interfaces) != 0 {

		if set_default {
			obj.Ipv4Interfaces().clearHolderSlice()
			for _, item := range obj.obj.Ipv4Interfaces {
				obj.Ipv4Interfaces().appendHolderSlice(&rsvpIpv4Interface{obj: item})
			}
		}
		for _, item := range obj.Ipv4Interfaces().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.LspIpv4Interfaces) != 0 {

		if set_default {
			obj.LspIpv4Interfaces().clearHolderSlice()
			for _, item := range obj.obj.LspIpv4Interfaces {
				obj.LspIpv4Interfaces().appendHolderSlice(&rsvpLspIpv4Interface{obj: item})
			}
		}
		for _, item := range obj.LspIpv4Interfaces().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"DeviceRsvp.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

}

func (obj *deviceRsvp) setDefault() {

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6InterfaceState *****
type dhcpv6InterfaceState struct {
	validation
	obj                 *otg.Dhcpv6InterfaceState
	marshaller          marshalDhcpv6InterfaceState
	unMarshaller        unMarshalDhcpv6InterfaceState
	iapdAddressesHolder Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter
	iaAddressesHolder   Dhcpv6InterfaceStateDhcpv6InterfaceIaIter
}

func NewDhcpv6InterfaceState() Dhcpv6InterfaceState {
	obj := dhcpv6InterfaceState{obj: &otg.Dhcpv6InterfaceState{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6InterfaceState) msg() *otg.Dhcpv6InterfaceState {
	return obj.obj
}

func (obj *dhcpv6InterfaceState) setMsg(msg *otg.Dhcpv6InterfaceState) Dhcpv6InterfaceState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6InterfaceState struct {
	obj *dhcpv6InterfaceState
}

type marshalDhcpv6InterfaceState interface {
	// ToProto marshals Dhcpv6InterfaceState to protobuf object *otg.Dhcpv6InterfaceState
	ToProto() (*otg.Dhcpv6InterfaceState, error)
	// ToPbText marshals Dhcpv6InterfaceState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6InterfaceState to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6InterfaceState to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv6InterfaceState struct {
	obj *dhcpv6InterfaceState
}

type unMarshalDhcpv6InterfaceState interface {
	// FromProto unmarshals Dhcpv6InterfaceState from protobuf object *otg.Dhcpv6InterfaceState
	FromProto(msg *otg.Dhcpv6InterfaceState) (Dhcpv6InterfaceState, error)
	// FromPbText unmarshals Dhcpv6InterfaceState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6InterfaceState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6InterfaceState from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6InterfaceState) Marshal() marshalDhcpv6InterfaceState {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6InterfaceState{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6InterfaceState) Unmarshal() unMarshalDhcpv6InterfaceState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6InterfaceState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6InterfaceState) ToProto() (*otg.Dhcpv6InterfaceState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6InterfaceState) FromProto(msg *otg.Dhcpv6InterfaceState) (Dhcpv6InterfaceState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6InterfaceState) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6InterfaceState) FromPbText(value string) error {
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

func (m *marshaldhcpv6InterfaceState) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6InterfaceState) FromYaml(value string) error {
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

func (m *marshaldhcpv6InterfaceState) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6InterfaceState) FromJson(value string) error {
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

func (obj *dhcpv6InterfaceState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6InterfaceState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6InterfaceState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6InterfaceState) Clone() (Dhcpv6InterfaceState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6InterfaceState()
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

func (obj *dhcpv6InterfaceState) setNil() {
	obj.iapdAddressesHolder = nil
	obj.iaAddressesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Dhcpv6InterfaceState is the IPv6 address associated with this DHCP Client session.
type Dhcpv6InterfaceState interface {
	Validation
	// msg marshals Dhcpv6InterfaceState to protobuf object *otg.Dhcpv6InterfaceState
	// and doesn't set defaults
	msg() *otg.Dhcpv6InterfaceState
	// setMsg unmarshals Dhcpv6InterfaceState from protobuf object *otg.Dhcpv6InterfaceState
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6InterfaceState) Dhcpv6InterfaceState
	// provides marshal interface
	Marshal() marshalDhcpv6InterfaceState
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6InterfaceState
	// validate validates Dhcpv6InterfaceState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6InterfaceState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// DhcpClientName returns string, set in Dhcpv6InterfaceState.
	DhcpClientName() string
	// SetDhcpClientName assigns string provided by user to Dhcpv6InterfaceState
	SetDhcpClientName(value string) Dhcpv6InterfaceState
	// HasDhcpClientName checks if DhcpClientName has been set in Dhcpv6InterfaceState
	HasDhcpClientName() bool
	// IapdAddresses returns Dhcpv6InterfaceStateDhcpv6InterfaceIapdIterIter, set in Dhcpv6InterfaceState
	IapdAddresses() Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter
	// IaAddresses returns Dhcpv6InterfaceStateDhcpv6InterfaceIaIterIter, set in Dhcpv6InterfaceState
	IaAddresses() Dhcpv6InterfaceStateDhcpv6InterfaceIaIter
	setNil()
}

// The name of a DHCPv6 Client.
// DhcpClientName returns a string
func (obj *dhcpv6InterfaceState) DhcpClientName() string {

	return *obj.obj.DhcpClientName

}

// The name of a DHCPv6 Client.
// DhcpClientName returns a string
func (obj *dhcpv6InterfaceState) HasDhcpClientName() bool {
	return obj.obj.DhcpClientName != nil
}

// The name of a DHCPv6 Client.
// SetDhcpClientName sets the string value in the Dhcpv6InterfaceState object
func (obj *dhcpv6InterfaceState) SetDhcpClientName(value string) Dhcpv6InterfaceState {

	obj.obj.DhcpClientName = &value
	return obj
}

// The IPv6 IAPD addresses and prefixes associated with this DHCP Client session.
// IapdAddresses returns a []Dhcpv6InterfaceIapd
func (obj *dhcpv6InterfaceState) IapdAddresses() Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter {
	if len(obj.obj.IapdAddresses) == 0 {
		obj.obj.IapdAddresses = []*otg.Dhcpv6InterfaceIapd{}
	}
	if obj.iapdAddressesHolder == nil {
		obj.iapdAddressesHolder = newDhcpv6InterfaceStateDhcpv6InterfaceIapdIter(&obj.obj.IapdAddresses).setMsg(obj)
	}
	return obj.iapdAddressesHolder
}

type dhcpv6InterfaceStateDhcpv6InterfaceIapdIter struct {
	obj                      *dhcpv6InterfaceState
	dhcpv6InterfaceIapdSlice []Dhcpv6InterfaceIapd
	fieldPtr                 *[]*otg.Dhcpv6InterfaceIapd
}

func newDhcpv6InterfaceStateDhcpv6InterfaceIapdIter(ptr *[]*otg.Dhcpv6InterfaceIapd) Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter {
	return &dhcpv6InterfaceStateDhcpv6InterfaceIapdIter{fieldPtr: ptr}
}

type Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter interface {
	setMsg(*dhcpv6InterfaceState) Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter
	Items() []Dhcpv6InterfaceIapd
	Add() Dhcpv6InterfaceIapd
	Append(items ...Dhcpv6InterfaceIapd) Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter
	Set(index int, newObj Dhcpv6InterfaceIapd) Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter
	Clear() Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter
	clearHolderSlice() Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter
	appendHolderSlice(item Dhcpv6InterfaceIapd) Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter
}

func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIapdIter) setMsg(msg *dhcpv6InterfaceState) Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpv6InterfaceIapd{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIapdIter) Items() []Dhcpv6InterfaceIapd {
	return obj.dhcpv6InterfaceIapdSlice
}

func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIapdIter) Add() Dhcpv6InterfaceIapd {
	newObj := &otg.Dhcpv6InterfaceIapd{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpv6InterfaceIapd{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpv6InterfaceIapdSlice = append(obj.dhcpv6InterfaceIapdSlice, newLibObj)
	return newLibObj
}

func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIapdIter) Append(items ...Dhcpv6InterfaceIapd) Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpv6InterfaceIapdSlice = append(obj.dhcpv6InterfaceIapdSlice, item)
	}
	return obj
}

func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIapdIter) Set(index int, newObj Dhcpv6InterfaceIapd) Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpv6InterfaceIapdSlice[index] = newObj
	return obj
}
func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIapdIter) Clear() Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Dhcpv6InterfaceIapd{}
		obj.dhcpv6InterfaceIapdSlice = []Dhcpv6InterfaceIapd{}
	}
	return obj
}
func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIapdIter) clearHolderSlice() Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter {
	if len(obj.dhcpv6InterfaceIapdSlice) > 0 {
		obj.dhcpv6InterfaceIapdSlice = []Dhcpv6InterfaceIapd{}
	}
	return obj
}
func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIapdIter) appendHolderSlice(item Dhcpv6InterfaceIapd) Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter {
	obj.dhcpv6InterfaceIapdSlice = append(obj.dhcpv6InterfaceIapdSlice, item)
	return obj
}

// The IPv6 IATA/IANA addresses and gateways associated with this DHCP Client session.
// IaAddresses returns a []Dhcpv6InterfaceIa
func (obj *dhcpv6InterfaceState) IaAddresses() Dhcpv6InterfaceStateDhcpv6InterfaceIaIter {
	if len(obj.obj.IaAddresses) == 0 {
		obj.obj.IaAddresses = []*otg.Dhcpv6InterfaceIa{}
	}
	if obj.iaAddressesHolder == nil {
		obj.iaAddressesHolder = newDhcpv6InterfaceStateDhcpv6InterfaceIaIter(&obj.obj.IaAddresses).setMsg(obj)
	}
	return obj.iaAddressesHolder
}

type dhcpv6InterfaceStateDhcpv6InterfaceIaIter struct {
	obj                    *dhcpv6InterfaceState
	dhcpv6InterfaceIaSlice []Dhcpv6InterfaceIa
	fieldPtr               *[]*otg.Dhcpv6InterfaceIa
}

func newDhcpv6InterfaceStateDhcpv6InterfaceIaIter(ptr *[]*otg.Dhcpv6InterfaceIa) Dhcpv6InterfaceStateDhcpv6InterfaceIaIter {
	return &dhcpv6InterfaceStateDhcpv6InterfaceIaIter{fieldPtr: ptr}
}

type Dhcpv6InterfaceStateDhcpv6InterfaceIaIter interface {
	setMsg(*dhcpv6InterfaceState) Dhcpv6InterfaceStateDhcpv6InterfaceIaIter
	Items() []Dhcpv6InterfaceIa
	Add() Dhcpv6InterfaceIa
	Append(items ...Dhcpv6InterfaceIa) Dhcpv6InterfaceStateDhcpv6InterfaceIaIter
	Set(index int, newObj Dhcpv6InterfaceIa) Dhcpv6InterfaceStateDhcpv6InterfaceIaIter
	Clear() Dhcpv6InterfaceStateDhcpv6InterfaceIaIter
	clearHolderSlice() Dhcpv6InterfaceStateDhcpv6InterfaceIaIter
	appendHolderSlice(item Dhcpv6InterfaceIa) Dhcpv6InterfaceStateDhcpv6InterfaceIaIter
}

func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIaIter) setMsg(msg *dhcpv6InterfaceState) Dhcpv6InterfaceStateDhcpv6InterfaceIaIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpv6InterfaceIa{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIaIter) Items() []Dhcpv6InterfaceIa {
	return obj.dhcpv6InterfaceIaSlice
}

func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIaIter) Add() Dhcpv6InterfaceIa {
	newObj := &otg.Dhcpv6InterfaceIa{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpv6InterfaceIa{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpv6InterfaceIaSlice = append(obj.dhcpv6InterfaceIaSlice, newLibObj)
	return newLibObj
}

func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIaIter) Append(items ...Dhcpv6InterfaceIa) Dhcpv6InterfaceStateDhcpv6InterfaceIaIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpv6InterfaceIaSlice = append(obj.dhcpv6InterfaceIaSlice, item)
	}
	return obj
}

func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIaIter) Set(index int, newObj Dhcpv6InterfaceIa) Dhcpv6InterfaceStateDhcpv6InterfaceIaIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpv6InterfaceIaSlice[index] = newObj
	return obj
}
func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIaIter) Clear() Dhcpv6InterfaceStateDhcpv6InterfaceIaIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Dhcpv6InterfaceIa{}
		obj.dhcpv6InterfaceIaSlice = []Dhcpv6InterfaceIa{}
	}
	return obj
}
func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIaIter) clearHolderSlice() Dhcpv6InterfaceStateDhcpv6InterfaceIaIter {
	if len(obj.dhcpv6InterfaceIaSlice) > 0 {
		obj.dhcpv6InterfaceIaSlice = []Dhcpv6InterfaceIa{}
	}
	return obj
}
func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIaIter) appendHolderSlice(item Dhcpv6InterfaceIa) Dhcpv6InterfaceStateDhcpv6InterfaceIaIter {
	obj.dhcpv6InterfaceIaSlice = append(obj.dhcpv6InterfaceIaSlice, item)
	return obj
}

func (obj *dhcpv6InterfaceState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.IapdAddresses) != 0 {

		if set_default {
			obj.IapdAddresses().clearHolderSlice()
			for _, item := range obj.obj.IapdAddresses {
				obj.IapdAddresses().appendHolderSlice(&dhcpv6InterfaceIapd{obj: item})
			}
		}
		for _, item := range obj.IapdAddresses().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.IaAddresses) != 0 {

		if set_default {
			obj.IaAddresses().clearHolderSlice()
			for _, item := range obj.obj.IaAddresses {
				obj.IaAddresses().appendHolderSlice(&dhcpv6InterfaceIa{obj: item})
			}
		}
		for _, item := range obj.IaAddresses().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *dhcpv6InterfaceState) setDefault() {

}

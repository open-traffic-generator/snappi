package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6LeasesState *****
type dhcpv6LeasesState struct {
	validation
	obj          *otg.Dhcpv6LeasesState
	marshaller   marshalDhcpv6LeasesState
	unMarshaller unMarshalDhcpv6LeasesState
	leasesHolder Dhcpv6LeasesStateDhcpv6ServerLeaseStateIter
}

func NewDhcpv6LeasesState() Dhcpv6LeasesState {
	obj := dhcpv6LeasesState{obj: &otg.Dhcpv6LeasesState{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6LeasesState) msg() *otg.Dhcpv6LeasesState {
	return obj.obj
}

func (obj *dhcpv6LeasesState) setMsg(msg *otg.Dhcpv6LeasesState) Dhcpv6LeasesState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6LeasesState struct {
	obj *dhcpv6LeasesState
}

type marshalDhcpv6LeasesState interface {
	// ToProto marshals Dhcpv6LeasesState to protobuf object *otg.Dhcpv6LeasesState
	ToProto() (*otg.Dhcpv6LeasesState, error)
	// ToPbText marshals Dhcpv6LeasesState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6LeasesState to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6LeasesState to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv6LeasesState struct {
	obj *dhcpv6LeasesState
}

type unMarshalDhcpv6LeasesState interface {
	// FromProto unmarshals Dhcpv6LeasesState from protobuf object *otg.Dhcpv6LeasesState
	FromProto(msg *otg.Dhcpv6LeasesState) (Dhcpv6LeasesState, error)
	// FromPbText unmarshals Dhcpv6LeasesState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6LeasesState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6LeasesState from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6LeasesState) Marshal() marshalDhcpv6LeasesState {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6LeasesState{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6LeasesState) Unmarshal() unMarshalDhcpv6LeasesState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6LeasesState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6LeasesState) ToProto() (*otg.Dhcpv6LeasesState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6LeasesState) FromProto(msg *otg.Dhcpv6LeasesState) (Dhcpv6LeasesState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6LeasesState) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6LeasesState) FromPbText(value string) error {
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

func (m *marshaldhcpv6LeasesState) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6LeasesState) FromYaml(value string) error {
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

func (m *marshaldhcpv6LeasesState) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6LeasesState) FromJson(value string) error {
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

func (obj *dhcpv6LeasesState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6LeasesState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6LeasesState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6LeasesState) Clone() (Dhcpv6LeasesState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6LeasesState()
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

func (obj *dhcpv6LeasesState) setNil() {
	obj.leasesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Dhcpv6LeasesState is lease information of DHCP Server.
type Dhcpv6LeasesState interface {
	Validation
	// msg marshals Dhcpv6LeasesState to protobuf object *otg.Dhcpv6LeasesState
	// and doesn't set defaults
	msg() *otg.Dhcpv6LeasesState
	// setMsg unmarshals Dhcpv6LeasesState from protobuf object *otg.Dhcpv6LeasesState
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6LeasesState) Dhcpv6LeasesState
	// provides marshal interface
	Marshal() marshalDhcpv6LeasesState
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6LeasesState
	// validate validates Dhcpv6LeasesState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6LeasesState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// DhcpServerName returns string, set in Dhcpv6LeasesState.
	DhcpServerName() string
	// SetDhcpServerName assigns string provided by user to Dhcpv6LeasesState
	SetDhcpServerName(value string) Dhcpv6LeasesState
	// HasDhcpServerName checks if DhcpServerName has been set in Dhcpv6LeasesState
	HasDhcpServerName() bool
	// Leases returns Dhcpv6LeasesStateDhcpv6ServerLeaseStateIterIter, set in Dhcpv6LeasesState
	Leases() Dhcpv6LeasesStateDhcpv6ServerLeaseStateIter
	setNil()
}

// The name of a DHCP Server.
// DhcpServerName returns a string
func (obj *dhcpv6LeasesState) DhcpServerName() string {

	return *obj.obj.DhcpServerName

}

// The name of a DHCP Server.
// DhcpServerName returns a string
func (obj *dhcpv6LeasesState) HasDhcpServerName() bool {
	return obj.obj.DhcpServerName != nil
}

// The name of a DHCP Server.
// SetDhcpServerName sets the string value in the Dhcpv6LeasesState object
func (obj *dhcpv6LeasesState) SetDhcpServerName(value string) Dhcpv6LeasesState {

	obj.obj.DhcpServerName = &value
	return obj
}

// description is TBD
// Leases returns a []Dhcpv6ServerLeaseState
func (obj *dhcpv6LeasesState) Leases() Dhcpv6LeasesStateDhcpv6ServerLeaseStateIter {
	if len(obj.obj.Leases) == 0 {
		obj.obj.Leases = []*otg.Dhcpv6ServerLeaseState{}
	}
	if obj.leasesHolder == nil {
		obj.leasesHolder = newDhcpv6LeasesStateDhcpv6ServerLeaseStateIter(&obj.obj.Leases).setMsg(obj)
	}
	return obj.leasesHolder
}

type dhcpv6LeasesStateDhcpv6ServerLeaseStateIter struct {
	obj                         *dhcpv6LeasesState
	dhcpv6ServerLeaseStateSlice []Dhcpv6ServerLeaseState
	fieldPtr                    *[]*otg.Dhcpv6ServerLeaseState
}

func newDhcpv6LeasesStateDhcpv6ServerLeaseStateIter(ptr *[]*otg.Dhcpv6ServerLeaseState) Dhcpv6LeasesStateDhcpv6ServerLeaseStateIter {
	return &dhcpv6LeasesStateDhcpv6ServerLeaseStateIter{fieldPtr: ptr}
}

type Dhcpv6LeasesStateDhcpv6ServerLeaseStateIter interface {
	setMsg(*dhcpv6LeasesState) Dhcpv6LeasesStateDhcpv6ServerLeaseStateIter
	Items() []Dhcpv6ServerLeaseState
	Add() Dhcpv6ServerLeaseState
	Append(items ...Dhcpv6ServerLeaseState) Dhcpv6LeasesStateDhcpv6ServerLeaseStateIter
	Set(index int, newObj Dhcpv6ServerLeaseState) Dhcpv6LeasesStateDhcpv6ServerLeaseStateIter
	Clear() Dhcpv6LeasesStateDhcpv6ServerLeaseStateIter
	clearHolderSlice() Dhcpv6LeasesStateDhcpv6ServerLeaseStateIter
	appendHolderSlice(item Dhcpv6ServerLeaseState) Dhcpv6LeasesStateDhcpv6ServerLeaseStateIter
}

func (obj *dhcpv6LeasesStateDhcpv6ServerLeaseStateIter) setMsg(msg *dhcpv6LeasesState) Dhcpv6LeasesStateDhcpv6ServerLeaseStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpv6ServerLeaseState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *dhcpv6LeasesStateDhcpv6ServerLeaseStateIter) Items() []Dhcpv6ServerLeaseState {
	return obj.dhcpv6ServerLeaseStateSlice
}

func (obj *dhcpv6LeasesStateDhcpv6ServerLeaseStateIter) Add() Dhcpv6ServerLeaseState {
	newObj := &otg.Dhcpv6ServerLeaseState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpv6ServerLeaseState{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpv6ServerLeaseStateSlice = append(obj.dhcpv6ServerLeaseStateSlice, newLibObj)
	return newLibObj
}

func (obj *dhcpv6LeasesStateDhcpv6ServerLeaseStateIter) Append(items ...Dhcpv6ServerLeaseState) Dhcpv6LeasesStateDhcpv6ServerLeaseStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpv6ServerLeaseStateSlice = append(obj.dhcpv6ServerLeaseStateSlice, item)
	}
	return obj
}

func (obj *dhcpv6LeasesStateDhcpv6ServerLeaseStateIter) Set(index int, newObj Dhcpv6ServerLeaseState) Dhcpv6LeasesStateDhcpv6ServerLeaseStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpv6ServerLeaseStateSlice[index] = newObj
	return obj
}
func (obj *dhcpv6LeasesStateDhcpv6ServerLeaseStateIter) Clear() Dhcpv6LeasesStateDhcpv6ServerLeaseStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Dhcpv6ServerLeaseState{}
		obj.dhcpv6ServerLeaseStateSlice = []Dhcpv6ServerLeaseState{}
	}
	return obj
}
func (obj *dhcpv6LeasesStateDhcpv6ServerLeaseStateIter) clearHolderSlice() Dhcpv6LeasesStateDhcpv6ServerLeaseStateIter {
	if len(obj.dhcpv6ServerLeaseStateSlice) > 0 {
		obj.dhcpv6ServerLeaseStateSlice = []Dhcpv6ServerLeaseState{}
	}
	return obj
}
func (obj *dhcpv6LeasesStateDhcpv6ServerLeaseStateIter) appendHolderSlice(item Dhcpv6ServerLeaseState) Dhcpv6LeasesStateDhcpv6ServerLeaseStateIter {
	obj.dhcpv6ServerLeaseStateSlice = append(obj.dhcpv6ServerLeaseStateSlice, item)
	return obj
}

func (obj *dhcpv6LeasesState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Leases) != 0 {

		if set_default {
			obj.Leases().clearHolderSlice()
			for _, item := range obj.obj.Leases {
				obj.Leases().appendHolderSlice(&dhcpv6ServerLeaseState{obj: item})
			}
		}
		for _, item := range obj.Leases().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *dhcpv6LeasesState) setDefault() {

}

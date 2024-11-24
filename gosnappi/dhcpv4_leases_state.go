package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv4LeasesState *****
type dhcpv4LeasesState struct {
	validation
	obj          *otg.Dhcpv4LeasesState
	marshaller   marshalDhcpv4LeasesState
	unMarshaller unMarshalDhcpv4LeasesState
	leasesHolder Dhcpv4LeasesStateDhcpv4LeaseStateIter
}

func NewDhcpv4LeasesState() Dhcpv4LeasesState {
	obj := dhcpv4LeasesState{obj: &otg.Dhcpv4LeasesState{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv4LeasesState) msg() *otg.Dhcpv4LeasesState {
	return obj.obj
}

func (obj *dhcpv4LeasesState) setMsg(msg *otg.Dhcpv4LeasesState) Dhcpv4LeasesState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv4LeasesState struct {
	obj *dhcpv4LeasesState
}

type marshalDhcpv4LeasesState interface {
	// ToProto marshals Dhcpv4LeasesState to protobuf object *otg.Dhcpv4LeasesState
	ToProto() (*otg.Dhcpv4LeasesState, error)
	// ToPbText marshals Dhcpv4LeasesState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv4LeasesState to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv4LeasesState to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv4LeasesState to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv4LeasesState struct {
	obj *dhcpv4LeasesState
}

type unMarshalDhcpv4LeasesState interface {
	// FromProto unmarshals Dhcpv4LeasesState from protobuf object *otg.Dhcpv4LeasesState
	FromProto(msg *otg.Dhcpv4LeasesState) (Dhcpv4LeasesState, error)
	// FromPbText unmarshals Dhcpv4LeasesState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv4LeasesState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv4LeasesState from JSON text
	FromJson(value string) error
}

func (obj *dhcpv4LeasesState) Marshal() marshalDhcpv4LeasesState {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv4LeasesState{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv4LeasesState) Unmarshal() unMarshalDhcpv4LeasesState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv4LeasesState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv4LeasesState) ToProto() (*otg.Dhcpv4LeasesState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv4LeasesState) FromProto(msg *otg.Dhcpv4LeasesState) (Dhcpv4LeasesState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv4LeasesState) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv4LeasesState) FromPbText(value string) error {
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

func (m *marshaldhcpv4LeasesState) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv4LeasesState) FromYaml(value string) error {
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

func (m *marshaldhcpv4LeasesState) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshaldhcpv4LeasesState) ToJson() (string, error) {
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

func (m *unMarshaldhcpv4LeasesState) FromJson(value string) error {
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

func (obj *dhcpv4LeasesState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv4LeasesState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv4LeasesState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv4LeasesState) Clone() (Dhcpv4LeasesState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv4LeasesState()
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

func (obj *dhcpv4LeasesState) setNil() {
	obj.leasesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Dhcpv4LeasesState is lease information of DHCP Server.
type Dhcpv4LeasesState interface {
	Validation
	// msg marshals Dhcpv4LeasesState to protobuf object *otg.Dhcpv4LeasesState
	// and doesn't set defaults
	msg() *otg.Dhcpv4LeasesState
	// setMsg unmarshals Dhcpv4LeasesState from protobuf object *otg.Dhcpv4LeasesState
	// and doesn't set defaults
	setMsg(*otg.Dhcpv4LeasesState) Dhcpv4LeasesState
	// provides marshal interface
	Marshal() marshalDhcpv4LeasesState
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv4LeasesState
	// validate validates Dhcpv4LeasesState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv4LeasesState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// DhcpServerName returns string, set in Dhcpv4LeasesState.
	DhcpServerName() string
	// SetDhcpServerName assigns string provided by user to Dhcpv4LeasesState
	SetDhcpServerName(value string) Dhcpv4LeasesState
	// HasDhcpServerName checks if DhcpServerName has been set in Dhcpv4LeasesState
	HasDhcpServerName() bool
	// Leases returns Dhcpv4LeasesStateDhcpv4LeaseStateIterIter, set in Dhcpv4LeasesState
	Leases() Dhcpv4LeasesStateDhcpv4LeaseStateIter
	setNil()
}

// The name of a DHCP Server.
// DhcpServerName returns a string
func (obj *dhcpv4LeasesState) DhcpServerName() string {

	return *obj.obj.DhcpServerName

}

// The name of a DHCP Server.
// DhcpServerName returns a string
func (obj *dhcpv4LeasesState) HasDhcpServerName() bool {
	return obj.obj.DhcpServerName != nil
}

// The name of a DHCP Server.
// SetDhcpServerName sets the string value in the Dhcpv4LeasesState object
func (obj *dhcpv4LeasesState) SetDhcpServerName(value string) Dhcpv4LeasesState {

	obj.obj.DhcpServerName = &value
	return obj
}

// description is TBD
// Leases returns a []Dhcpv4LeaseState
func (obj *dhcpv4LeasesState) Leases() Dhcpv4LeasesStateDhcpv4LeaseStateIter {
	if len(obj.obj.Leases) == 0 {
		obj.obj.Leases = []*otg.Dhcpv4LeaseState{}
	}
	if obj.leasesHolder == nil {
		obj.leasesHolder = newDhcpv4LeasesStateDhcpv4LeaseStateIter(&obj.obj.Leases).setMsg(obj)
	}
	return obj.leasesHolder
}

type dhcpv4LeasesStateDhcpv4LeaseStateIter struct {
	obj                   *dhcpv4LeasesState
	dhcpv4LeaseStateSlice []Dhcpv4LeaseState
	fieldPtr              *[]*otg.Dhcpv4LeaseState
}

func newDhcpv4LeasesStateDhcpv4LeaseStateIter(ptr *[]*otg.Dhcpv4LeaseState) Dhcpv4LeasesStateDhcpv4LeaseStateIter {
	return &dhcpv4LeasesStateDhcpv4LeaseStateIter{fieldPtr: ptr}
}

type Dhcpv4LeasesStateDhcpv4LeaseStateIter interface {
	setMsg(*dhcpv4LeasesState) Dhcpv4LeasesStateDhcpv4LeaseStateIter
	Items() []Dhcpv4LeaseState
	Add() Dhcpv4LeaseState
	Append(items ...Dhcpv4LeaseState) Dhcpv4LeasesStateDhcpv4LeaseStateIter
	Set(index int, newObj Dhcpv4LeaseState) Dhcpv4LeasesStateDhcpv4LeaseStateIter
	Clear() Dhcpv4LeasesStateDhcpv4LeaseStateIter
	clearHolderSlice() Dhcpv4LeasesStateDhcpv4LeaseStateIter
	appendHolderSlice(item Dhcpv4LeaseState) Dhcpv4LeasesStateDhcpv4LeaseStateIter
}

func (obj *dhcpv4LeasesStateDhcpv4LeaseStateIter) setMsg(msg *dhcpv4LeasesState) Dhcpv4LeasesStateDhcpv4LeaseStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpv4LeaseState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *dhcpv4LeasesStateDhcpv4LeaseStateIter) Items() []Dhcpv4LeaseState {
	return obj.dhcpv4LeaseStateSlice
}

func (obj *dhcpv4LeasesStateDhcpv4LeaseStateIter) Add() Dhcpv4LeaseState {
	newObj := &otg.Dhcpv4LeaseState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpv4LeaseState{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpv4LeaseStateSlice = append(obj.dhcpv4LeaseStateSlice, newLibObj)
	return newLibObj
}

func (obj *dhcpv4LeasesStateDhcpv4LeaseStateIter) Append(items ...Dhcpv4LeaseState) Dhcpv4LeasesStateDhcpv4LeaseStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpv4LeaseStateSlice = append(obj.dhcpv4LeaseStateSlice, item)
	}
	return obj
}

func (obj *dhcpv4LeasesStateDhcpv4LeaseStateIter) Set(index int, newObj Dhcpv4LeaseState) Dhcpv4LeasesStateDhcpv4LeaseStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpv4LeaseStateSlice[index] = newObj
	return obj
}
func (obj *dhcpv4LeasesStateDhcpv4LeaseStateIter) Clear() Dhcpv4LeasesStateDhcpv4LeaseStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Dhcpv4LeaseState{}
		obj.dhcpv4LeaseStateSlice = []Dhcpv4LeaseState{}
	}
	return obj
}
func (obj *dhcpv4LeasesStateDhcpv4LeaseStateIter) clearHolderSlice() Dhcpv4LeasesStateDhcpv4LeaseStateIter {
	if len(obj.dhcpv4LeaseStateSlice) > 0 {
		obj.dhcpv4LeaseStateSlice = []Dhcpv4LeaseState{}
	}
	return obj
}
func (obj *dhcpv4LeasesStateDhcpv4LeaseStateIter) appendHolderSlice(item Dhcpv4LeaseState) Dhcpv4LeasesStateDhcpv4LeaseStateIter {
	obj.dhcpv4LeaseStateSlice = append(obj.dhcpv4LeaseStateSlice, item)
	return obj
}

func (obj *dhcpv4LeasesState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Leases) != 0 {

		if set_default {
			obj.Leases().clearHolderSlice()
			for _, item := range obj.obj.Leases {
				obj.Leases().appendHolderSlice(&dhcpv4LeaseState{obj: item})
			}
		}
		for _, item := range obj.Leases().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *dhcpv4LeasesState) setDefault() {

}

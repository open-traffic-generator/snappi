package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LdpBindingsState *****
type ldpBindingsState struct {
	validation
	obj                *otg.LdpBindingsState
	marshaller         marshalLdpBindingsState
	unMarshaller       unMarshalLdpBindingsState
	ipv4BindingsHolder LdpBindingsStateLdpIpv4BindingStateIter
}

func NewLdpBindingsState() LdpBindingsState {
	obj := ldpBindingsState{obj: &otg.LdpBindingsState{}}
	obj.setDefault()
	return &obj
}

func (obj *ldpBindingsState) msg() *otg.LdpBindingsState {
	return obj.obj
}

func (obj *ldpBindingsState) setMsg(msg *otg.LdpBindingsState) LdpBindingsState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalldpBindingsState struct {
	obj *ldpBindingsState
}

type marshalLdpBindingsState interface {
	// ToProto marshals LdpBindingsState to protobuf object *otg.LdpBindingsState
	ToProto() (*otg.LdpBindingsState, error)
	// ToPbText marshals LdpBindingsState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LdpBindingsState to YAML text
	ToYaml() (string, error)
	// ToJson marshals LdpBindingsState to JSON text
	ToJson() (string, error)
}

type unMarshalldpBindingsState struct {
	obj *ldpBindingsState
}

type unMarshalLdpBindingsState interface {
	// FromProto unmarshals LdpBindingsState from protobuf object *otg.LdpBindingsState
	FromProto(msg *otg.LdpBindingsState) (LdpBindingsState, error)
	// FromPbText unmarshals LdpBindingsState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LdpBindingsState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LdpBindingsState from JSON text
	FromJson(value string) error
}

func (obj *ldpBindingsState) Marshal() marshalLdpBindingsState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalldpBindingsState{obj: obj}
	}
	return obj.marshaller
}

func (obj *ldpBindingsState) Unmarshal() unMarshalLdpBindingsState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalldpBindingsState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalldpBindingsState) ToProto() (*otg.LdpBindingsState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalldpBindingsState) FromProto(msg *otg.LdpBindingsState) (LdpBindingsState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalldpBindingsState) ToPbText() (string, error) {
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

func (m *unMarshalldpBindingsState) FromPbText(value string) error {
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

func (m *marshalldpBindingsState) ToYaml() (string, error) {
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

func (m *unMarshalldpBindingsState) FromYaml(value string) error {
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

func (m *marshalldpBindingsState) ToJson() (string, error) {
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

func (m *unMarshalldpBindingsState) FromJson(value string) error {
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

func (obj *ldpBindingsState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ldpBindingsState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ldpBindingsState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ldpBindingsState) Clone() (LdpBindingsState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLdpBindingsState()
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

func (obj *ldpBindingsState) setNil() {
	obj.ipv4BindingsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LdpBindingsState is the FEC-label bindings received by one LDP router from its peers.
type LdpBindingsState interface {
	Validation
	// msg marshals LdpBindingsState to protobuf object *otg.LdpBindingsState
	// and doesn't set defaults
	msg() *otg.LdpBindingsState
	// setMsg unmarshals LdpBindingsState from protobuf object *otg.LdpBindingsState
	// and doesn't set defaults
	setMsg(*otg.LdpBindingsState) LdpBindingsState
	// provides marshal interface
	Marshal() marshalLdpBindingsState
	// provides unmarshal interface
	Unmarshal() unMarshalLdpBindingsState
	// validate validates LdpBindingsState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LdpBindingsState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LdpRouterName returns string, set in LdpBindingsState.
	LdpRouterName() string
	// SetLdpRouterName assigns string provided by user to LdpBindingsState
	SetLdpRouterName(value string) LdpBindingsState
	// HasLdpRouterName checks if LdpRouterName has been set in LdpBindingsState
	HasLdpRouterName() bool
	// Ipv4Bindings returns LdpBindingsStateLdpIpv4BindingStateIterIter, set in LdpBindingsState
	Ipv4Bindings() LdpBindingsStateLdpIpv4BindingStateIter
	setNil()
}

// The name of the LDP router.
// LdpRouterName returns a string
func (obj *ldpBindingsState) LdpRouterName() string {

	return *obj.obj.LdpRouterName

}

// The name of the LDP router.
// LdpRouterName returns a string
func (obj *ldpBindingsState) HasLdpRouterName() bool {
	return obj.obj.LdpRouterName != nil
}

// The name of the LDP router.
// SetLdpRouterName sets the string value in the LdpBindingsState object
func (obj *ldpBindingsState) SetLdpRouterName(value string) LdpBindingsState {

	obj.obj.LdpRouterName = &value
	return obj
}

// The IPv4 Prefix FEC-label bindings received in Label Mapping messages (RFC 5036 Section 3.5.7).
// Ipv4Bindings returns a []LdpIpv4BindingState
func (obj *ldpBindingsState) Ipv4Bindings() LdpBindingsStateLdpIpv4BindingStateIter {
	if len(obj.obj.Ipv4Bindings) == 0 {
		obj.obj.Ipv4Bindings = []*otg.LdpIpv4BindingState{}
	}
	if obj.ipv4BindingsHolder == nil {
		obj.ipv4BindingsHolder = newLdpBindingsStateLdpIpv4BindingStateIter(&obj.obj.Ipv4Bindings).setMsg(obj)
	}
	return obj.ipv4BindingsHolder
}

type ldpBindingsStateLdpIpv4BindingStateIter struct {
	obj                      *ldpBindingsState
	ldpIpv4BindingStateSlice []LdpIpv4BindingState
	fieldPtr                 *[]*otg.LdpIpv4BindingState
}

func newLdpBindingsStateLdpIpv4BindingStateIter(ptr *[]*otg.LdpIpv4BindingState) LdpBindingsStateLdpIpv4BindingStateIter {
	return &ldpBindingsStateLdpIpv4BindingStateIter{fieldPtr: ptr}
}

type LdpBindingsStateLdpIpv4BindingStateIter interface {
	setMsg(*ldpBindingsState) LdpBindingsStateLdpIpv4BindingStateIter
	Items() []LdpIpv4BindingState
	Add() LdpIpv4BindingState
	Append(items ...LdpIpv4BindingState) LdpBindingsStateLdpIpv4BindingStateIter
	Set(index int, newObj LdpIpv4BindingState) LdpBindingsStateLdpIpv4BindingStateIter
	Clear() LdpBindingsStateLdpIpv4BindingStateIter
	clearHolderSlice() LdpBindingsStateLdpIpv4BindingStateIter
	appendHolderSlice(item LdpIpv4BindingState) LdpBindingsStateLdpIpv4BindingStateIter
}

func (obj *ldpBindingsStateLdpIpv4BindingStateIter) setMsg(msg *ldpBindingsState) LdpBindingsStateLdpIpv4BindingStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ldpIpv4BindingState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ldpBindingsStateLdpIpv4BindingStateIter) Items() []LdpIpv4BindingState {
	return obj.ldpIpv4BindingStateSlice
}

func (obj *ldpBindingsStateLdpIpv4BindingStateIter) Add() LdpIpv4BindingState {
	newObj := &otg.LdpIpv4BindingState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ldpIpv4BindingState{obj: newObj}
	newLibObj.setDefault()
	obj.ldpIpv4BindingStateSlice = append(obj.ldpIpv4BindingStateSlice, newLibObj)
	return newLibObj
}

func (obj *ldpBindingsStateLdpIpv4BindingStateIter) Append(items ...LdpIpv4BindingState) LdpBindingsStateLdpIpv4BindingStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ldpIpv4BindingStateSlice = append(obj.ldpIpv4BindingStateSlice, item)
	}
	return obj
}

func (obj *ldpBindingsStateLdpIpv4BindingStateIter) Set(index int, newObj LdpIpv4BindingState) LdpBindingsStateLdpIpv4BindingStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ldpIpv4BindingStateSlice[index] = newObj
	return obj
}
func (obj *ldpBindingsStateLdpIpv4BindingStateIter) Clear() LdpBindingsStateLdpIpv4BindingStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.LdpIpv4BindingState{}
		obj.ldpIpv4BindingStateSlice = []LdpIpv4BindingState{}
	}
	return obj
}
func (obj *ldpBindingsStateLdpIpv4BindingStateIter) clearHolderSlice() LdpBindingsStateLdpIpv4BindingStateIter {
	if len(obj.ldpIpv4BindingStateSlice) > 0 {
		obj.ldpIpv4BindingStateSlice = []LdpIpv4BindingState{}
	}
	return obj
}
func (obj *ldpBindingsStateLdpIpv4BindingStateIter) appendHolderSlice(item LdpIpv4BindingState) LdpBindingsStateLdpIpv4BindingStateIter {
	obj.ldpIpv4BindingStateSlice = append(obj.ldpIpv4BindingStateSlice, item)
	return obj
}

func (obj *ldpBindingsState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Ipv4Bindings) != 0 {

		if set_default {
			obj.Ipv4Bindings().clearHolderSlice()
			for _, item := range obj.obj.Ipv4Bindings {
				obj.Ipv4Bindings().appendHolderSlice(&ldpIpv4BindingState{obj: item})
			}
		}
		for _, item := range obj.Ipv4Bindings().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *ldpBindingsState) setDefault() {

}

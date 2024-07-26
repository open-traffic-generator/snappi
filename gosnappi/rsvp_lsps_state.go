package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RsvpLspsState *****
type rsvpLspsState struct {
	validation
	obj            *otg.RsvpLspsState
	marshaller     marshalRsvpLspsState
	unMarshaller   unMarshalRsvpLspsState
	ipv4LspsHolder RsvpLspsStateRsvpIPv4LspStateIter
}

func NewRsvpLspsState() RsvpLspsState {
	obj := rsvpLspsState{obj: &otg.RsvpLspsState{}}
	obj.setDefault()
	return &obj
}

func (obj *rsvpLspsState) msg() *otg.RsvpLspsState {
	return obj.obj
}

func (obj *rsvpLspsState) setMsg(msg *otg.RsvpLspsState) RsvpLspsState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrsvpLspsState struct {
	obj *rsvpLspsState
}

type marshalRsvpLspsState interface {
	// ToProto marshals RsvpLspsState to protobuf object *otg.RsvpLspsState
	ToProto() (*otg.RsvpLspsState, error)
	// ToPbText marshals RsvpLspsState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RsvpLspsState to YAML text
	ToYaml() (string, error)
	// ToJson marshals RsvpLspsState to JSON text
	ToJson() (string, error)
}

type unMarshalrsvpLspsState struct {
	obj *rsvpLspsState
}

type unMarshalRsvpLspsState interface {
	// FromProto unmarshals RsvpLspsState from protobuf object *otg.RsvpLspsState
	FromProto(msg *otg.RsvpLspsState) (RsvpLspsState, error)
	// FromPbText unmarshals RsvpLspsState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RsvpLspsState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RsvpLspsState from JSON text
	FromJson(value string) error
}

func (obj *rsvpLspsState) Marshal() marshalRsvpLspsState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrsvpLspsState{obj: obj}
	}
	return obj.marshaller
}

func (obj *rsvpLspsState) Unmarshal() unMarshalRsvpLspsState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrsvpLspsState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrsvpLspsState) ToProto() (*otg.RsvpLspsState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrsvpLspsState) FromProto(msg *otg.RsvpLspsState) (RsvpLspsState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrsvpLspsState) ToPbText() (string, error) {
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

func (m *unMarshalrsvpLspsState) FromPbText(value string) error {
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

func (m *marshalrsvpLspsState) ToYaml() (string, error) {
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

func (m *unMarshalrsvpLspsState) FromYaml(value string) error {
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

func (m *marshalrsvpLspsState) ToJson() (string, error) {
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

func (m *unMarshalrsvpLspsState) FromJson(value string) error {
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

func (obj *rsvpLspsState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rsvpLspsState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rsvpLspsState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rsvpLspsState) Clone() (RsvpLspsState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRsvpLspsState()
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

func (obj *rsvpLspsState) setNil() {
	obj.ipv4LspsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// RsvpLspsState is discovered IPv4 Point-to-Point LSPs of a RSVP-TE router.
type RsvpLspsState interface {
	Validation
	// msg marshals RsvpLspsState to protobuf object *otg.RsvpLspsState
	// and doesn't set defaults
	msg() *otg.RsvpLspsState
	// setMsg unmarshals RsvpLspsState from protobuf object *otg.RsvpLspsState
	// and doesn't set defaults
	setMsg(*otg.RsvpLspsState) RsvpLspsState
	// provides marshal interface
	Marshal() marshalRsvpLspsState
	// provides unmarshal interface
	Unmarshal() unMarshalRsvpLspsState
	// validate validates RsvpLspsState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RsvpLspsState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RsvpRouterName returns string, set in RsvpLspsState.
	RsvpRouterName() string
	// SetRsvpRouterName assigns string provided by user to RsvpLspsState
	SetRsvpRouterName(value string) RsvpLspsState
	// HasRsvpRouterName checks if RsvpRouterName has been set in RsvpLspsState
	HasRsvpRouterName() bool
	// Ipv4Lsps returns RsvpLspsStateRsvpIPv4LspStateIterIter, set in RsvpLspsState
	Ipv4Lsps() RsvpLspsStateRsvpIPv4LspStateIter
	setNil()
}

// The name of the RSVP-TE Router.
// RsvpRouterName returns a string
func (obj *rsvpLspsState) RsvpRouterName() string {

	return *obj.obj.RsvpRouterName

}

// The name of the RSVP-TE Router.
// RsvpRouterName returns a string
func (obj *rsvpLspsState) HasRsvpRouterName() bool {
	return obj.obj.RsvpRouterName != nil
}

// The name of the RSVP-TE Router.
// SetRsvpRouterName sets the string value in the RsvpLspsState object
func (obj *rsvpLspsState) SetRsvpRouterName(value string) RsvpLspsState {

	obj.obj.RsvpRouterName = &value
	return obj
}

// IPv4 Point-to-Point RSVP-TE Discovered LSPs.
// Ipv4Lsps returns a []RsvpIPv4LspState
func (obj *rsvpLspsState) Ipv4Lsps() RsvpLspsStateRsvpIPv4LspStateIter {
	if len(obj.obj.Ipv4Lsps) == 0 {
		obj.obj.Ipv4Lsps = []*otg.RsvpIPv4LspState{}
	}
	if obj.ipv4LspsHolder == nil {
		obj.ipv4LspsHolder = newRsvpLspsStateRsvpIPv4LspStateIter(&obj.obj.Ipv4Lsps).setMsg(obj)
	}
	return obj.ipv4LspsHolder
}

type rsvpLspsStateRsvpIPv4LspStateIter struct {
	obj                   *rsvpLspsState
	rsvpIPv4LspStateSlice []RsvpIPv4LspState
	fieldPtr              *[]*otg.RsvpIPv4LspState
}

func newRsvpLspsStateRsvpIPv4LspStateIter(ptr *[]*otg.RsvpIPv4LspState) RsvpLspsStateRsvpIPv4LspStateIter {
	return &rsvpLspsStateRsvpIPv4LspStateIter{fieldPtr: ptr}
}

type RsvpLspsStateRsvpIPv4LspStateIter interface {
	setMsg(*rsvpLspsState) RsvpLspsStateRsvpIPv4LspStateIter
	Items() []RsvpIPv4LspState
	Add() RsvpIPv4LspState
	Append(items ...RsvpIPv4LspState) RsvpLspsStateRsvpIPv4LspStateIter
	Set(index int, newObj RsvpIPv4LspState) RsvpLspsStateRsvpIPv4LspStateIter
	Clear() RsvpLspsStateRsvpIPv4LspStateIter
	clearHolderSlice() RsvpLspsStateRsvpIPv4LspStateIter
	appendHolderSlice(item RsvpIPv4LspState) RsvpLspsStateRsvpIPv4LspStateIter
}

func (obj *rsvpLspsStateRsvpIPv4LspStateIter) setMsg(msg *rsvpLspsState) RsvpLspsStateRsvpIPv4LspStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&rsvpIPv4LspState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *rsvpLspsStateRsvpIPv4LspStateIter) Items() []RsvpIPv4LspState {
	return obj.rsvpIPv4LspStateSlice
}

func (obj *rsvpLspsStateRsvpIPv4LspStateIter) Add() RsvpIPv4LspState {
	newObj := &otg.RsvpIPv4LspState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &rsvpIPv4LspState{obj: newObj}
	newLibObj.setDefault()
	obj.rsvpIPv4LspStateSlice = append(obj.rsvpIPv4LspStateSlice, newLibObj)
	return newLibObj
}

func (obj *rsvpLspsStateRsvpIPv4LspStateIter) Append(items ...RsvpIPv4LspState) RsvpLspsStateRsvpIPv4LspStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.rsvpIPv4LspStateSlice = append(obj.rsvpIPv4LspStateSlice, item)
	}
	return obj
}

func (obj *rsvpLspsStateRsvpIPv4LspStateIter) Set(index int, newObj RsvpIPv4LspState) RsvpLspsStateRsvpIPv4LspStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.rsvpIPv4LspStateSlice[index] = newObj
	return obj
}
func (obj *rsvpLspsStateRsvpIPv4LspStateIter) Clear() RsvpLspsStateRsvpIPv4LspStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.RsvpIPv4LspState{}
		obj.rsvpIPv4LspStateSlice = []RsvpIPv4LspState{}
	}
	return obj
}
func (obj *rsvpLspsStateRsvpIPv4LspStateIter) clearHolderSlice() RsvpLspsStateRsvpIPv4LspStateIter {
	if len(obj.rsvpIPv4LspStateSlice) > 0 {
		obj.rsvpIPv4LspStateSlice = []RsvpIPv4LspState{}
	}
	return obj
}
func (obj *rsvpLspsStateRsvpIPv4LspStateIter) appendHolderSlice(item RsvpIPv4LspState) RsvpLspsStateRsvpIPv4LspStateIter {
	obj.rsvpIPv4LspStateSlice = append(obj.rsvpIPv4LspStateSlice, item)
	return obj
}

func (obj *rsvpLspsState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Ipv4Lsps) != 0 {

		if set_default {
			obj.Ipv4Lsps().clearHolderSlice()
			for _, item := range obj.obj.Ipv4Lsps {
				obj.Ipv4Lsps().appendHolderSlice(&rsvpIPv4LspState{obj: item})
			}
		}
		for _, item := range obj.Ipv4Lsps().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *rsvpLspsState) setDefault() {

}

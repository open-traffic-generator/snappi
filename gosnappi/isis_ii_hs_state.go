package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisIIHsState *****
type isisIIHsState struct {
	validation
	obj                   *otg.IsisIIHsState
	marshaller            marshalIsisIIHsState
	unMarshaller          unMarshalIsisIIHsState
	statesHolder          IsisIIHsStateIsisLocalIIHStateIter
	neighborsStatesHolder IsisIIHsStateIsisNeighborIIHStateIter
}

func NewIsisIIHsState() IsisIIHsState {
	obj := isisIIHsState{obj: &otg.IsisIIHsState{}}
	obj.setDefault()
	return &obj
}

func (obj *isisIIHsState) msg() *otg.IsisIIHsState {
	return obj.obj
}

func (obj *isisIIHsState) setMsg(msg *otg.IsisIIHsState) IsisIIHsState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisIIHsState struct {
	obj *isisIIHsState
}

type marshalIsisIIHsState interface {
	// ToProto marshals IsisIIHsState to protobuf object *otg.IsisIIHsState
	ToProto() (*otg.IsisIIHsState, error)
	// ToPbText marshals IsisIIHsState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisIIHsState to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisIIHsState to JSON text
	ToJson() (string, error)
}

type unMarshalisisIIHsState struct {
	obj *isisIIHsState
}

type unMarshalIsisIIHsState interface {
	// FromProto unmarshals IsisIIHsState from protobuf object *otg.IsisIIHsState
	FromProto(msg *otg.IsisIIHsState) (IsisIIHsState, error)
	// FromPbText unmarshals IsisIIHsState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisIIHsState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisIIHsState from JSON text
	FromJson(value string) error
}

func (obj *isisIIHsState) Marshal() marshalIsisIIHsState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisIIHsState{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisIIHsState) Unmarshal() unMarshalIsisIIHsState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisIIHsState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisIIHsState) ToProto() (*otg.IsisIIHsState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisIIHsState) FromProto(msg *otg.IsisIIHsState) (IsisIIHsState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisIIHsState) ToPbText() (string, error) {
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

func (m *unMarshalisisIIHsState) FromPbText(value string) error {
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

func (m *marshalisisIIHsState) ToYaml() (string, error) {
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

func (m *unMarshalisisIIHsState) FromYaml(value string) error {
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

func (m *marshalisisIIHsState) ToJson() (string, error) {
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

func (m *unMarshalisisIIHsState) FromJson(value string) error {
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

func (obj *isisIIHsState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisIIHsState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisIIHsState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisIIHsState) Clone() (IsisIIHsState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisIIHsState()
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

func (obj *isisIIHsState) setNil() {
	obj.statesHolder = nil
	obj.neighborsStatesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisIIHsState is the result of ISIS IIH information that are exchanged.
type IsisIIHsState interface {
	Validation
	// msg marshals IsisIIHsState to protobuf object *otg.IsisIIHsState
	// and doesn't set defaults
	msg() *otg.IsisIIHsState
	// setMsg unmarshals IsisIIHsState from protobuf object *otg.IsisIIHsState
	// and doesn't set defaults
	setMsg(*otg.IsisIIHsState) IsisIIHsState
	// provides marshal interface
	Marshal() marshalIsisIIHsState
	// provides unmarshal interface
	Unmarshal() unMarshalIsisIIHsState
	// validate validates IsisIIHsState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisIIHsState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// IsisRouterName returns string, set in IsisIIHsState.
	IsisRouterName() string
	// SetIsisRouterName assigns string provided by user to IsisIIHsState
	SetIsisRouterName(value string) IsisIIHsState
	// HasIsisRouterName checks if IsisRouterName has been set in IsisIIHsState
	HasIsisRouterName() bool
	// States returns IsisIIHsStateIsisLocalIIHStateIterIter, set in IsisIIHsState
	States() IsisIIHsStateIsisLocalIIHStateIter
	// NeighborsStates returns IsisIIHsStateIsisNeighborIIHStateIterIter, set in IsisIIHsState
	NeighborsStates() IsisIIHsStateIsisNeighborIIHStateIter
	setNil()
}

// The name of the ISIS Router.
// IsisRouterName returns a string
func (obj *isisIIHsState) IsisRouterName() string {

	return *obj.obj.IsisRouterName

}

// The name of the ISIS Router.
// IsisRouterName returns a string
func (obj *isisIIHsState) HasIsisRouterName() bool {
	return obj.obj.IsisRouterName != nil
}

// The name of the ISIS Router.
// SetIsisRouterName sets the string value in the IsisIIHsState object
func (obj *isisIIHsState) SetIsisRouterName(value string) IsisIIHsState {

	obj.obj.IsisRouterName = &value
	return obj
}

// State of this router.
// States returns a []IsisLocalIIHState
func (obj *isisIIHsState) States() IsisIIHsStateIsisLocalIIHStateIter {
	if len(obj.obj.States) == 0 {
		obj.obj.States = []*otg.IsisLocalIIHState{}
	}
	if obj.statesHolder == nil {
		obj.statesHolder = newIsisIIHsStateIsisLocalIIHStateIter(&obj.obj.States).setMsg(obj)
	}
	return obj.statesHolder
}

type isisIIHsStateIsisLocalIIHStateIter struct {
	obj                    *isisIIHsState
	isisLocalIIHStateSlice []IsisLocalIIHState
	fieldPtr               *[]*otg.IsisLocalIIHState
}

func newIsisIIHsStateIsisLocalIIHStateIter(ptr *[]*otg.IsisLocalIIHState) IsisIIHsStateIsisLocalIIHStateIter {
	return &isisIIHsStateIsisLocalIIHStateIter{fieldPtr: ptr}
}

type IsisIIHsStateIsisLocalIIHStateIter interface {
	setMsg(*isisIIHsState) IsisIIHsStateIsisLocalIIHStateIter
	Items() []IsisLocalIIHState
	Add() IsisLocalIIHState
	Append(items ...IsisLocalIIHState) IsisIIHsStateIsisLocalIIHStateIter
	Set(index int, newObj IsisLocalIIHState) IsisIIHsStateIsisLocalIIHStateIter
	Clear() IsisIIHsStateIsisLocalIIHStateIter
	clearHolderSlice() IsisIIHsStateIsisLocalIIHStateIter
	appendHolderSlice(item IsisLocalIIHState) IsisIIHsStateIsisLocalIIHStateIter
}

func (obj *isisIIHsStateIsisLocalIIHStateIter) setMsg(msg *isisIIHsState) IsisIIHsStateIsisLocalIIHStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLocalIIHState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisIIHsStateIsisLocalIIHStateIter) Items() []IsisLocalIIHState {
	return obj.isisLocalIIHStateSlice
}

func (obj *isisIIHsStateIsisLocalIIHStateIter) Add() IsisLocalIIHState {
	newObj := &otg.IsisLocalIIHState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLocalIIHState{obj: newObj}
	newLibObj.setDefault()
	obj.isisLocalIIHStateSlice = append(obj.isisLocalIIHStateSlice, newLibObj)
	return newLibObj
}

func (obj *isisIIHsStateIsisLocalIIHStateIter) Append(items ...IsisLocalIIHState) IsisIIHsStateIsisLocalIIHStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLocalIIHStateSlice = append(obj.isisLocalIIHStateSlice, item)
	}
	return obj
}

func (obj *isisIIHsStateIsisLocalIIHStateIter) Set(index int, newObj IsisLocalIIHState) IsisIIHsStateIsisLocalIIHStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLocalIIHStateSlice[index] = newObj
	return obj
}
func (obj *isisIIHsStateIsisLocalIIHStateIter) Clear() IsisIIHsStateIsisLocalIIHStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLocalIIHState{}
		obj.isisLocalIIHStateSlice = []IsisLocalIIHState{}
	}
	return obj
}
func (obj *isisIIHsStateIsisLocalIIHStateIter) clearHolderSlice() IsisIIHsStateIsisLocalIIHStateIter {
	if len(obj.isisLocalIIHStateSlice) > 0 {
		obj.isisLocalIIHStateSlice = []IsisLocalIIHState{}
	}
	return obj
}
func (obj *isisIIHsStateIsisLocalIIHStateIter) appendHolderSlice(item IsisLocalIIHState) IsisIIHsStateIsisLocalIIHStateIter {
	obj.isisLocalIIHStateSlice = append(obj.isisLocalIIHStateSlice, item)
	return obj
}

// One or more Neigbors that are learned by this ISIS router.
// NeighborsStates returns a []IsisNeighborIIHState
func (obj *isisIIHsState) NeighborsStates() IsisIIHsStateIsisNeighborIIHStateIter {
	if len(obj.obj.NeighborsStates) == 0 {
		obj.obj.NeighborsStates = []*otg.IsisNeighborIIHState{}
	}
	if obj.neighborsStatesHolder == nil {
		obj.neighborsStatesHolder = newIsisIIHsStateIsisNeighborIIHStateIter(&obj.obj.NeighborsStates).setMsg(obj)
	}
	return obj.neighborsStatesHolder
}

type isisIIHsStateIsisNeighborIIHStateIter struct {
	obj                       *isisIIHsState
	isisNeighborIIHStateSlice []IsisNeighborIIHState
	fieldPtr                  *[]*otg.IsisNeighborIIHState
}

func newIsisIIHsStateIsisNeighborIIHStateIter(ptr *[]*otg.IsisNeighborIIHState) IsisIIHsStateIsisNeighborIIHStateIter {
	return &isisIIHsStateIsisNeighborIIHStateIter{fieldPtr: ptr}
}

type IsisIIHsStateIsisNeighborIIHStateIter interface {
	setMsg(*isisIIHsState) IsisIIHsStateIsisNeighborIIHStateIter
	Items() []IsisNeighborIIHState
	Add() IsisNeighborIIHState
	Append(items ...IsisNeighborIIHState) IsisIIHsStateIsisNeighborIIHStateIter
	Set(index int, newObj IsisNeighborIIHState) IsisIIHsStateIsisNeighborIIHStateIter
	Clear() IsisIIHsStateIsisNeighborIIHStateIter
	clearHolderSlice() IsisIIHsStateIsisNeighborIIHStateIter
	appendHolderSlice(item IsisNeighborIIHState) IsisIIHsStateIsisNeighborIIHStateIter
}

func (obj *isisIIHsStateIsisNeighborIIHStateIter) setMsg(msg *isisIIHsState) IsisIIHsStateIsisNeighborIIHStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisNeighborIIHState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisIIHsStateIsisNeighborIIHStateIter) Items() []IsisNeighborIIHState {
	return obj.isisNeighborIIHStateSlice
}

func (obj *isisIIHsStateIsisNeighborIIHStateIter) Add() IsisNeighborIIHState {
	newObj := &otg.IsisNeighborIIHState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisNeighborIIHState{obj: newObj}
	newLibObj.setDefault()
	obj.isisNeighborIIHStateSlice = append(obj.isisNeighborIIHStateSlice, newLibObj)
	return newLibObj
}

func (obj *isisIIHsStateIsisNeighborIIHStateIter) Append(items ...IsisNeighborIIHState) IsisIIHsStateIsisNeighborIIHStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisNeighborIIHStateSlice = append(obj.isisNeighborIIHStateSlice, item)
	}
	return obj
}

func (obj *isisIIHsStateIsisNeighborIIHStateIter) Set(index int, newObj IsisNeighborIIHState) IsisIIHsStateIsisNeighborIIHStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisNeighborIIHStateSlice[index] = newObj
	return obj
}
func (obj *isisIIHsStateIsisNeighborIIHStateIter) Clear() IsisIIHsStateIsisNeighborIIHStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisNeighborIIHState{}
		obj.isisNeighborIIHStateSlice = []IsisNeighborIIHState{}
	}
	return obj
}
func (obj *isisIIHsStateIsisNeighborIIHStateIter) clearHolderSlice() IsisIIHsStateIsisNeighborIIHStateIter {
	if len(obj.isisNeighborIIHStateSlice) > 0 {
		obj.isisNeighborIIHStateSlice = []IsisNeighborIIHState{}
	}
	return obj
}
func (obj *isisIIHsStateIsisNeighborIIHStateIter) appendHolderSlice(item IsisNeighborIIHState) IsisIIHsStateIsisNeighborIIHStateIter {
	obj.isisNeighborIIHStateSlice = append(obj.isisNeighborIIHStateSlice, item)
	return obj
}

func (obj *isisIIHsState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.States) != 0 {

		if set_default {
			obj.States().clearHolderSlice()
			for _, item := range obj.obj.States {
				obj.States().appendHolderSlice(&isisLocalIIHState{obj: item})
			}
		}
		for _, item := range obj.States().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.NeighborsStates) != 0 {

		if set_default {
			obj.NeighborsStates().clearHolderSlice()
			for _, item := range obj.obj.NeighborsStates {
				obj.NeighborsStates().appendHolderSlice(&isisNeighborIIHState{obj: item})
			}
		}
		for _, item := range obj.NeighborsStates().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *isisIIHsState) setDefault() {

}

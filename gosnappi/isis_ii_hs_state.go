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
	adjacencyStatesHolder IsisIIHsStateIsisLocalIIHAdjacencyStatesIter
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
	obj.adjacencyStatesHolder = nil
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
	// AdjacencyStates returns IsisIIHsStateIsisLocalIIHAdjacencyStatesIterIter, set in IsisIIHsState
	AdjacencyStates() IsisIIHsStateIsisLocalIIHAdjacencyStatesIter
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

// Current state of adjacencies.
// AdjacencyStates returns a []IsisLocalIIHAdjacencyStates
func (obj *isisIIHsState) AdjacencyStates() IsisIIHsStateIsisLocalIIHAdjacencyStatesIter {
	if len(obj.obj.AdjacencyStates) == 0 {
		obj.obj.AdjacencyStates = []*otg.IsisLocalIIHAdjacencyStates{}
	}
	if obj.adjacencyStatesHolder == nil {
		obj.adjacencyStatesHolder = newIsisIIHsStateIsisLocalIIHAdjacencyStatesIter(&obj.obj.AdjacencyStates).setMsg(obj)
	}
	return obj.adjacencyStatesHolder
}

type isisIIHsStateIsisLocalIIHAdjacencyStatesIter struct {
	obj                              *isisIIHsState
	isisLocalIIHAdjacencyStatesSlice []IsisLocalIIHAdjacencyStates
	fieldPtr                         *[]*otg.IsisLocalIIHAdjacencyStates
}

func newIsisIIHsStateIsisLocalIIHAdjacencyStatesIter(ptr *[]*otg.IsisLocalIIHAdjacencyStates) IsisIIHsStateIsisLocalIIHAdjacencyStatesIter {
	return &isisIIHsStateIsisLocalIIHAdjacencyStatesIter{fieldPtr: ptr}
}

type IsisIIHsStateIsisLocalIIHAdjacencyStatesIter interface {
	setMsg(*isisIIHsState) IsisIIHsStateIsisLocalIIHAdjacencyStatesIter
	Items() []IsisLocalIIHAdjacencyStates
	Add() IsisLocalIIHAdjacencyStates
	Append(items ...IsisLocalIIHAdjacencyStates) IsisIIHsStateIsisLocalIIHAdjacencyStatesIter
	Set(index int, newObj IsisLocalIIHAdjacencyStates) IsisIIHsStateIsisLocalIIHAdjacencyStatesIter
	Clear() IsisIIHsStateIsisLocalIIHAdjacencyStatesIter
	clearHolderSlice() IsisIIHsStateIsisLocalIIHAdjacencyStatesIter
	appendHolderSlice(item IsisLocalIIHAdjacencyStates) IsisIIHsStateIsisLocalIIHAdjacencyStatesIter
}

func (obj *isisIIHsStateIsisLocalIIHAdjacencyStatesIter) setMsg(msg *isisIIHsState) IsisIIHsStateIsisLocalIIHAdjacencyStatesIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLocalIIHAdjacencyStates{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisIIHsStateIsisLocalIIHAdjacencyStatesIter) Items() []IsisLocalIIHAdjacencyStates {
	return obj.isisLocalIIHAdjacencyStatesSlice
}

func (obj *isisIIHsStateIsisLocalIIHAdjacencyStatesIter) Add() IsisLocalIIHAdjacencyStates {
	newObj := &otg.IsisLocalIIHAdjacencyStates{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLocalIIHAdjacencyStates{obj: newObj}
	newLibObj.setDefault()
	obj.isisLocalIIHAdjacencyStatesSlice = append(obj.isisLocalIIHAdjacencyStatesSlice, newLibObj)
	return newLibObj
}

func (obj *isisIIHsStateIsisLocalIIHAdjacencyStatesIter) Append(items ...IsisLocalIIHAdjacencyStates) IsisIIHsStateIsisLocalIIHAdjacencyStatesIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLocalIIHAdjacencyStatesSlice = append(obj.isisLocalIIHAdjacencyStatesSlice, item)
	}
	return obj
}

func (obj *isisIIHsStateIsisLocalIIHAdjacencyStatesIter) Set(index int, newObj IsisLocalIIHAdjacencyStates) IsisIIHsStateIsisLocalIIHAdjacencyStatesIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLocalIIHAdjacencyStatesSlice[index] = newObj
	return obj
}
func (obj *isisIIHsStateIsisLocalIIHAdjacencyStatesIter) Clear() IsisIIHsStateIsisLocalIIHAdjacencyStatesIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLocalIIHAdjacencyStates{}
		obj.isisLocalIIHAdjacencyStatesSlice = []IsisLocalIIHAdjacencyStates{}
	}
	return obj
}
func (obj *isisIIHsStateIsisLocalIIHAdjacencyStatesIter) clearHolderSlice() IsisIIHsStateIsisLocalIIHAdjacencyStatesIter {
	if len(obj.isisLocalIIHAdjacencyStatesSlice) > 0 {
		obj.isisLocalIIHAdjacencyStatesSlice = []IsisLocalIIHAdjacencyStates{}
	}
	return obj
}
func (obj *isisIIHsStateIsisLocalIIHAdjacencyStatesIter) appendHolderSlice(item IsisLocalIIHAdjacencyStates) IsisIIHsStateIsisLocalIIHAdjacencyStatesIter {
	obj.isisLocalIIHAdjacencyStatesSlice = append(obj.isisLocalIIHAdjacencyStatesSlice, item)
	return obj
}

func (obj *isisIIHsState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.AdjacencyStates) != 0 {

		if set_default {
			obj.AdjacencyStates().clearHolderSlice()
			for _, item := range obj.obj.AdjacencyStates {
				obj.AdjacencyStates().appendHolderSlice(&isisLocalIIHAdjacencyStates{obj: item})
			}
		}
		for _, item := range obj.AdjacencyStates().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *isisIIHsState) setDefault() {

}

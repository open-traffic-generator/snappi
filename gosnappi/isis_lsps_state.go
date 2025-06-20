package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspsState *****
type isisLspsState struct {
	validation
	obj          *otg.IsisLspsState
	marshaller   marshalIsisLspsState
	unMarshaller unMarshalIsisLspsState
	lspsHolder   IsisLspsStateIsisLspStateIter
}

func NewIsisLspsState() IsisLspsState {
	obj := isisLspsState{obj: &otg.IsisLspsState{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspsState) msg() *otg.IsisLspsState {
	return obj.obj
}

func (obj *isisLspsState) setMsg(msg *otg.IsisLspsState) IsisLspsState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspsState struct {
	obj *isisLspsState
}

type marshalIsisLspsState interface {
	// ToProto marshals IsisLspsState to protobuf object *otg.IsisLspsState
	ToProto() (*otg.IsisLspsState, error)
	// ToPbText marshals IsisLspsState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspsState to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspsState to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspsState struct {
	obj *isisLspsState
}

type unMarshalIsisLspsState interface {
	// FromProto unmarshals IsisLspsState from protobuf object *otg.IsisLspsState
	FromProto(msg *otg.IsisLspsState) (IsisLspsState, error)
	// FromPbText unmarshals IsisLspsState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspsState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspsState from JSON text
	FromJson(value string) error
}

func (obj *isisLspsState) Marshal() marshalIsisLspsState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspsState{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspsState) Unmarshal() unMarshalIsisLspsState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspsState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspsState) ToProto() (*otg.IsisLspsState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspsState) FromProto(msg *otg.IsisLspsState) (IsisLspsState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspsState) ToPbText() (string, error) {
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

func (m *unMarshalisisLspsState) FromPbText(value string) error {
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

func (m *marshalisisLspsState) ToYaml() (string, error) {
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

func (m *unMarshalisisLspsState) FromYaml(value string) error {
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

func (m *marshalisisLspsState) ToJson() (string, error) {
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

func (m *unMarshalisisLspsState) FromJson(value string) error {
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

func (obj *isisLspsState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspsState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspsState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspsState) Clone() (IsisLspsState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspsState()
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

func (obj *isisLspsState) setNil() {
	obj.lspsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspsState is the result of ISIS LSP information that are retrieved.
type IsisLspsState interface {
	Validation
	// msg marshals IsisLspsState to protobuf object *otg.IsisLspsState
	// and doesn't set defaults
	msg() *otg.IsisLspsState
	// setMsg unmarshals IsisLspsState from protobuf object *otg.IsisLspsState
	// and doesn't set defaults
	setMsg(*otg.IsisLspsState) IsisLspsState
	// provides marshal interface
	Marshal() marshalIsisLspsState
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspsState
	// validate validates IsisLspsState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspsState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// IsisRouterName returns string, set in IsisLspsState.
	IsisRouterName() string
	// SetIsisRouterName assigns string provided by user to IsisLspsState
	SetIsisRouterName(value string) IsisLspsState
	// HasIsisRouterName checks if IsisRouterName has been set in IsisLspsState
	HasIsisRouterName() bool
	// Lsps returns IsisLspsStateIsisLspStateIterIter, set in IsisLspsState
	Lsps() IsisLspsStateIsisLspStateIter
	setNil()
}

// The name of the ISIS Router.
// IsisRouterName returns a string
func (obj *isisLspsState) IsisRouterName() string {

	return *obj.obj.IsisRouterName

}

// The name of the ISIS Router.
// IsisRouterName returns a string
func (obj *isisLspsState) HasIsisRouterName() bool {
	return obj.obj.IsisRouterName != nil
}

// The name of the ISIS Router.
// SetIsisRouterName sets the string value in the IsisLspsState object
func (obj *isisLspsState) SetIsisRouterName(value string) IsisLspsState {

	obj.obj.IsisRouterName = &value
	return obj
}

// One or more LSPs that are learned by this ISIS router.
// Lsps returns a []IsisLspState
func (obj *isisLspsState) Lsps() IsisLspsStateIsisLspStateIter {
	if len(obj.obj.Lsps) == 0 {
		obj.obj.Lsps = []*otg.IsisLspState{}
	}
	if obj.lspsHolder == nil {
		obj.lspsHolder = newIsisLspsStateIsisLspStateIter(&obj.obj.Lsps).setMsg(obj)
	}
	return obj.lspsHolder
}

type isisLspsStateIsisLspStateIter struct {
	obj               *isisLspsState
	isisLspStateSlice []IsisLspState
	fieldPtr          *[]*otg.IsisLspState
}

func newIsisLspsStateIsisLspStateIter(ptr *[]*otg.IsisLspState) IsisLspsStateIsisLspStateIter {
	return &isisLspsStateIsisLspStateIter{fieldPtr: ptr}
}

type IsisLspsStateIsisLspStateIter interface {
	setMsg(*isisLspsState) IsisLspsStateIsisLspStateIter
	Items() []IsisLspState
	Add() IsisLspState
	Append(items ...IsisLspState) IsisLspsStateIsisLspStateIter
	Set(index int, newObj IsisLspState) IsisLspsStateIsisLspStateIter
	Clear() IsisLspsStateIsisLspStateIter
	clearHolderSlice() IsisLspsStateIsisLspStateIter
	appendHolderSlice(item IsisLspState) IsisLspsStateIsisLspStateIter
}

func (obj *isisLspsStateIsisLspStateIter) setMsg(msg *isisLspsState) IsisLspsStateIsisLspStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspsStateIsisLspStateIter) Items() []IsisLspState {
	return obj.isisLspStateSlice
}

func (obj *isisLspsStateIsisLspStateIter) Add() IsisLspState {
	newObj := &otg.IsisLspState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspState{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspStateSlice = append(obj.isisLspStateSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspsStateIsisLspStateIter) Append(items ...IsisLspState) IsisLspsStateIsisLspStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspStateSlice = append(obj.isisLspStateSlice, item)
	}
	return obj
}

func (obj *isisLspsStateIsisLspStateIter) Set(index int, newObj IsisLspState) IsisLspsStateIsisLspStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspStateSlice[index] = newObj
	return obj
}
func (obj *isisLspsStateIsisLspStateIter) Clear() IsisLspsStateIsisLspStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspState{}
		obj.isisLspStateSlice = []IsisLspState{}
	}
	return obj
}
func (obj *isisLspsStateIsisLspStateIter) clearHolderSlice() IsisLspsStateIsisLspStateIter {
	if len(obj.isisLspStateSlice) > 0 {
		obj.isisLspStateSlice = []IsisLspState{}
	}
	return obj
}
func (obj *isisLspsStateIsisLspStateIter) appendHolderSlice(item IsisLspState) IsisLspsStateIsisLspStateIter {
	obj.isisLspStateSlice = append(obj.isisLspStateSlice, item)
	return obj
}

func (obj *isisLspsState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Lsps) != 0 {

		if set_default {
			obj.Lsps().clearHolderSlice()
			for _, item := range obj.obj.Lsps {
				obj.Lsps().appendHolderSlice(&isisLspState{obj: item})
			}
		}
		for _, item := range obj.Lsps().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *isisLspsState) setDefault() {

}

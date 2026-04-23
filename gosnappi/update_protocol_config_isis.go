package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateProtocolConfigIsis *****
type updateProtocolConfigIsis struct {
	validation
	obj              *otg.UpdateProtocolConfigIsis
	marshaller       marshalUpdateProtocolConfigIsis
	unMarshaller     unMarshalUpdateProtocolConfigIsis
	routersHolder    UpdateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter
	interfacesHolder UpdateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter
}

func NewUpdateProtocolConfigIsis() UpdateProtocolConfigIsis {
	obj := updateProtocolConfigIsis{obj: &otg.UpdateProtocolConfigIsis{}}
	obj.setDefault()
	return &obj
}

func (obj *updateProtocolConfigIsis) msg() *otg.UpdateProtocolConfigIsis {
	return obj.obj
}

func (obj *updateProtocolConfigIsis) setMsg(msg *otg.UpdateProtocolConfigIsis) UpdateProtocolConfigIsis {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateProtocolConfigIsis struct {
	obj *updateProtocolConfigIsis
}

type marshalUpdateProtocolConfigIsis interface {
	// ToProto marshals UpdateProtocolConfigIsis to protobuf object *otg.UpdateProtocolConfigIsis
	ToProto() (*otg.UpdateProtocolConfigIsis, error)
	// ToPbText marshals UpdateProtocolConfigIsis to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateProtocolConfigIsis to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateProtocolConfigIsis to JSON text
	ToJson() (string, error)
}

type unMarshalupdateProtocolConfigIsis struct {
	obj *updateProtocolConfigIsis
}

type unMarshalUpdateProtocolConfigIsis interface {
	// FromProto unmarshals UpdateProtocolConfigIsis from protobuf object *otg.UpdateProtocolConfigIsis
	FromProto(msg *otg.UpdateProtocolConfigIsis) (UpdateProtocolConfigIsis, error)
	// FromPbText unmarshals UpdateProtocolConfigIsis from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateProtocolConfigIsis from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateProtocolConfigIsis from JSON text
	FromJson(value string) error
}

func (obj *updateProtocolConfigIsis) Marshal() marshalUpdateProtocolConfigIsis {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateProtocolConfigIsis{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateProtocolConfigIsis) Unmarshal() unMarshalUpdateProtocolConfigIsis {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateProtocolConfigIsis{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateProtocolConfigIsis) ToProto() (*otg.UpdateProtocolConfigIsis, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateProtocolConfigIsis) FromProto(msg *otg.UpdateProtocolConfigIsis) (UpdateProtocolConfigIsis, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateProtocolConfigIsis) ToPbText() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsis) FromPbText(value string) error {
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

func (m *marshalupdateProtocolConfigIsis) ToYaml() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsis) FromYaml(value string) error {
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

func (m *marshalupdateProtocolConfigIsis) ToJson() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsis) FromJson(value string) error {
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

func (obj *updateProtocolConfigIsis) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsis) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsis) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateProtocolConfigIsis) Clone() (UpdateProtocolConfigIsis, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateProtocolConfigIsis()
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

func (obj *updateProtocolConfigIsis) setNil() {
	obj.routersHolder = nil
	obj.interfacesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UpdateProtocolConfigIsis is a container for IS-IS properties to be updated. Presence of this object indicates that one or more IS-IS properties require updating. The choice field selects the category of attributes to be updated. Router-level and interface-level updates are controlled independently.
type UpdateProtocolConfigIsis interface {
	Validation
	// msg marshals UpdateProtocolConfigIsis to protobuf object *otg.UpdateProtocolConfigIsis
	// and doesn't set defaults
	msg() *otg.UpdateProtocolConfigIsis
	// setMsg unmarshals UpdateProtocolConfigIsis from protobuf object *otg.UpdateProtocolConfigIsis
	// and doesn't set defaults
	setMsg(*otg.UpdateProtocolConfigIsis) UpdateProtocolConfigIsis
	// provides marshal interface
	Marshal() marshalUpdateProtocolConfigIsis
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateProtocolConfigIsis
	// validate validates UpdateProtocolConfigIsis
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateProtocolConfigIsis, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Routers returns UpdateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIterIter, set in UpdateProtocolConfigIsis
	Routers() UpdateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter
	// Interfaces returns UpdateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIterIter, set in UpdateProtocolConfigIsis
	Interfaces() UpdateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter
	setNil()
}

// List of IS-IS router-level update entries. Populate this field to update router-level attributes independently of interface-level attributes.
// Routers returns a []UpdateProtocolConfigIsisRouterUpdateGroup
func (obj *updateProtocolConfigIsis) Routers() UpdateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter {
	if len(obj.obj.Routers) == 0 {
		obj.obj.Routers = []*otg.UpdateProtocolConfigIsisRouterUpdateGroup{}
	}
	if obj.routersHolder == nil {
		obj.routersHolder = newUpdateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter(&obj.obj.Routers).setMsg(obj)
	}
	return obj.routersHolder
}

type updateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter struct {
	obj                                            *updateProtocolConfigIsis
	updateProtocolConfigIsisRouterUpdateGroupSlice []UpdateProtocolConfigIsisRouterUpdateGroup
	fieldPtr                                       *[]*otg.UpdateProtocolConfigIsisRouterUpdateGroup
}

func newUpdateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter(ptr *[]*otg.UpdateProtocolConfigIsisRouterUpdateGroup) UpdateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter {
	return &updateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter{fieldPtr: ptr}
}

type UpdateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter interface {
	setMsg(*updateProtocolConfigIsis) UpdateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter
	Items() []UpdateProtocolConfigIsisRouterUpdateGroup
	Add() UpdateProtocolConfigIsisRouterUpdateGroup
	Append(items ...UpdateProtocolConfigIsisRouterUpdateGroup) UpdateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter
	Set(index int, newObj UpdateProtocolConfigIsisRouterUpdateGroup) UpdateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter
	Clear() UpdateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter
	clearHolderSlice() UpdateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter
	appendHolderSlice(item UpdateProtocolConfigIsisRouterUpdateGroup) UpdateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter
}

func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter) setMsg(msg *updateProtocolConfigIsis) UpdateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&updateProtocolConfigIsisRouterUpdateGroup{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter) Items() []UpdateProtocolConfigIsisRouterUpdateGroup {
	return obj.updateProtocolConfigIsisRouterUpdateGroupSlice
}

func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter) Add() UpdateProtocolConfigIsisRouterUpdateGroup {
	newObj := &otg.UpdateProtocolConfigIsisRouterUpdateGroup{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &updateProtocolConfigIsisRouterUpdateGroup{obj: newObj}
	newLibObj.setDefault()
	obj.updateProtocolConfigIsisRouterUpdateGroupSlice = append(obj.updateProtocolConfigIsisRouterUpdateGroupSlice, newLibObj)
	return newLibObj
}

func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter) Append(items ...UpdateProtocolConfigIsisRouterUpdateGroup) UpdateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.updateProtocolConfigIsisRouterUpdateGroupSlice = append(obj.updateProtocolConfigIsisRouterUpdateGroupSlice, item)
	}
	return obj
}

func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter) Set(index int, newObj UpdateProtocolConfigIsisRouterUpdateGroup) UpdateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.updateProtocolConfigIsisRouterUpdateGroupSlice[index] = newObj
	return obj
}
func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter) Clear() UpdateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.UpdateProtocolConfigIsisRouterUpdateGroup{}
		obj.updateProtocolConfigIsisRouterUpdateGroupSlice = []UpdateProtocolConfigIsisRouterUpdateGroup{}
	}
	return obj
}
func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter) clearHolderSlice() UpdateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter {
	if len(obj.updateProtocolConfigIsisRouterUpdateGroupSlice) > 0 {
		obj.updateProtocolConfigIsisRouterUpdateGroupSlice = []UpdateProtocolConfigIsisRouterUpdateGroup{}
	}
	return obj
}
func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter) appendHolderSlice(item UpdateProtocolConfigIsisRouterUpdateGroup) UpdateProtocolConfigIsisUpdateProtocolConfigIsisRouterUpdateGroupIter {
	obj.updateProtocolConfigIsisRouterUpdateGroupSlice = append(obj.updateProtocolConfigIsisRouterUpdateGroupSlice, item)
	return obj
}

// List of IS-IS interface update groups. IS-IS interface names are globally unique across all routers so no parent router name is required. Each update group applies a common set of property updates to one or more interfaces simultaneously. Use multiple update groups to handle asymmetric changes across different subsets of interfaces.
// Interfaces returns a []UpdateProtocolConfigIsisInterfaceUpdateGroup
func (obj *updateProtocolConfigIsis) Interfaces() UpdateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter {
	if len(obj.obj.Interfaces) == 0 {
		obj.obj.Interfaces = []*otg.UpdateProtocolConfigIsisInterfaceUpdateGroup{}
	}
	if obj.interfacesHolder == nil {
		obj.interfacesHolder = newUpdateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter(&obj.obj.Interfaces).setMsg(obj)
	}
	return obj.interfacesHolder
}

type updateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter struct {
	obj                                               *updateProtocolConfigIsis
	updateProtocolConfigIsisInterfaceUpdateGroupSlice []UpdateProtocolConfigIsisInterfaceUpdateGroup
	fieldPtr                                          *[]*otg.UpdateProtocolConfigIsisInterfaceUpdateGroup
}

func newUpdateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter(ptr *[]*otg.UpdateProtocolConfigIsisInterfaceUpdateGroup) UpdateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter {
	return &updateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter{fieldPtr: ptr}
}

type UpdateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter interface {
	setMsg(*updateProtocolConfigIsis) UpdateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter
	Items() []UpdateProtocolConfigIsisInterfaceUpdateGroup
	Add() UpdateProtocolConfigIsisInterfaceUpdateGroup
	Append(items ...UpdateProtocolConfigIsisInterfaceUpdateGroup) UpdateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter
	Set(index int, newObj UpdateProtocolConfigIsisInterfaceUpdateGroup) UpdateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter
	Clear() UpdateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter
	clearHolderSlice() UpdateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter
	appendHolderSlice(item UpdateProtocolConfigIsisInterfaceUpdateGroup) UpdateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter
}

func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter) setMsg(msg *updateProtocolConfigIsis) UpdateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&updateProtocolConfigIsisInterfaceUpdateGroup{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter) Items() []UpdateProtocolConfigIsisInterfaceUpdateGroup {
	return obj.updateProtocolConfigIsisInterfaceUpdateGroupSlice
}

func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter) Add() UpdateProtocolConfigIsisInterfaceUpdateGroup {
	newObj := &otg.UpdateProtocolConfigIsisInterfaceUpdateGroup{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &updateProtocolConfigIsisInterfaceUpdateGroup{obj: newObj}
	newLibObj.setDefault()
	obj.updateProtocolConfigIsisInterfaceUpdateGroupSlice = append(obj.updateProtocolConfigIsisInterfaceUpdateGroupSlice, newLibObj)
	return newLibObj
}

func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter) Append(items ...UpdateProtocolConfigIsisInterfaceUpdateGroup) UpdateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.updateProtocolConfigIsisInterfaceUpdateGroupSlice = append(obj.updateProtocolConfigIsisInterfaceUpdateGroupSlice, item)
	}
	return obj
}

func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter) Set(index int, newObj UpdateProtocolConfigIsisInterfaceUpdateGroup) UpdateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.updateProtocolConfigIsisInterfaceUpdateGroupSlice[index] = newObj
	return obj
}
func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter) Clear() UpdateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.UpdateProtocolConfigIsisInterfaceUpdateGroup{}
		obj.updateProtocolConfigIsisInterfaceUpdateGroupSlice = []UpdateProtocolConfigIsisInterfaceUpdateGroup{}
	}
	return obj
}
func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter) clearHolderSlice() UpdateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter {
	if len(obj.updateProtocolConfigIsisInterfaceUpdateGroupSlice) > 0 {
		obj.updateProtocolConfigIsisInterfaceUpdateGroupSlice = []UpdateProtocolConfigIsisInterfaceUpdateGroup{}
	}
	return obj
}
func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter) appendHolderSlice(item UpdateProtocolConfigIsisInterfaceUpdateGroup) UpdateProtocolConfigIsisUpdateProtocolConfigIsisInterfaceUpdateGroupIter {
	obj.updateProtocolConfigIsisInterfaceUpdateGroupSlice = append(obj.updateProtocolConfigIsisInterfaceUpdateGroupSlice, item)
	return obj
}

func (obj *updateProtocolConfigIsis) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Routers) != 0 {

		if set_default {
			obj.Routers().clearHolderSlice()
			for _, item := range obj.obj.Routers {
				obj.Routers().appendHolderSlice(&updateProtocolConfigIsisRouterUpdateGroup{obj: item})
			}
		}
		for _, item := range obj.Routers().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Interfaces) != 0 {

		if set_default {
			obj.Interfaces().clearHolderSlice()
			for _, item := range obj.obj.Interfaces {
				obj.Interfaces().appendHolderSlice(&updateProtocolConfigIsisInterfaceUpdateGroup{obj: item})
			}
		}
		for _, item := range obj.Interfaces().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *updateProtocolConfigIsis) setDefault() {

}

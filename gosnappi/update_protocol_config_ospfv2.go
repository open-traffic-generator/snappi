package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateProtocolConfigOspfv2 *****
type updateProtocolConfigOspfv2 struct {
	validation
	obj              *otg.UpdateProtocolConfigOspfv2
	marshaller       marshalUpdateProtocolConfigOspfv2
	unMarshaller     unMarshalUpdateProtocolConfigOspfv2
	interfacesHolder UpdateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter
}

func NewUpdateProtocolConfigOspfv2() UpdateProtocolConfigOspfv2 {
	obj := updateProtocolConfigOspfv2{obj: &otg.UpdateProtocolConfigOspfv2{}}
	obj.setDefault()
	return &obj
}

func (obj *updateProtocolConfigOspfv2) msg() *otg.UpdateProtocolConfigOspfv2 {
	return obj.obj
}

func (obj *updateProtocolConfigOspfv2) setMsg(msg *otg.UpdateProtocolConfigOspfv2) UpdateProtocolConfigOspfv2 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateProtocolConfigOspfv2 struct {
	obj *updateProtocolConfigOspfv2
}

type marshalUpdateProtocolConfigOspfv2 interface {
	// ToProto marshals UpdateProtocolConfigOspfv2 to protobuf object *otg.UpdateProtocolConfigOspfv2
	ToProto() (*otg.UpdateProtocolConfigOspfv2, error)
	// ToPbText marshals UpdateProtocolConfigOspfv2 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateProtocolConfigOspfv2 to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateProtocolConfigOspfv2 to JSON text
	ToJson() (string, error)
}

type unMarshalupdateProtocolConfigOspfv2 struct {
	obj *updateProtocolConfigOspfv2
}

type unMarshalUpdateProtocolConfigOspfv2 interface {
	// FromProto unmarshals UpdateProtocolConfigOspfv2 from protobuf object *otg.UpdateProtocolConfigOspfv2
	FromProto(msg *otg.UpdateProtocolConfigOspfv2) (UpdateProtocolConfigOspfv2, error)
	// FromPbText unmarshals UpdateProtocolConfigOspfv2 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateProtocolConfigOspfv2 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateProtocolConfigOspfv2 from JSON text
	FromJson(value string) error
}

func (obj *updateProtocolConfigOspfv2) Marshal() marshalUpdateProtocolConfigOspfv2 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateProtocolConfigOspfv2{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateProtocolConfigOspfv2) Unmarshal() unMarshalUpdateProtocolConfigOspfv2 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateProtocolConfigOspfv2{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateProtocolConfigOspfv2) ToProto() (*otg.UpdateProtocolConfigOspfv2, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateProtocolConfigOspfv2) FromProto(msg *otg.UpdateProtocolConfigOspfv2) (UpdateProtocolConfigOspfv2, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateProtocolConfigOspfv2) ToPbText() (string, error) {
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

func (m *unMarshalupdateProtocolConfigOspfv2) FromPbText(value string) error {
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

func (m *marshalupdateProtocolConfigOspfv2) ToYaml() (string, error) {
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

func (m *unMarshalupdateProtocolConfigOspfv2) FromYaml(value string) error {
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

func (m *marshalupdateProtocolConfigOspfv2) ToJson() (string, error) {
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

func (m *unMarshalupdateProtocolConfigOspfv2) FromJson(value string) error {
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

func (obj *updateProtocolConfigOspfv2) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateProtocolConfigOspfv2) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateProtocolConfigOspfv2) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateProtocolConfigOspfv2) Clone() (UpdateProtocolConfigOspfv2, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateProtocolConfigOspfv2()
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

func (obj *updateProtocolConfigOspfv2) setNil() {
	obj.interfacesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UpdateProtocolConfigOspfv2 is a container for OSPFv2 properties to be updated. Presence of this object indicates that one or more OSPFv2 properties require updating.
type UpdateProtocolConfigOspfv2 interface {
	Validation
	// msg marshals UpdateProtocolConfigOspfv2 to protobuf object *otg.UpdateProtocolConfigOspfv2
	// and doesn't set defaults
	msg() *otg.UpdateProtocolConfigOspfv2
	// setMsg unmarshals UpdateProtocolConfigOspfv2 from protobuf object *otg.UpdateProtocolConfigOspfv2
	// and doesn't set defaults
	setMsg(*otg.UpdateProtocolConfigOspfv2) UpdateProtocolConfigOspfv2
	// provides marshal interface
	Marshal() marshalUpdateProtocolConfigOspfv2
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateProtocolConfigOspfv2
	// validate validates UpdateProtocolConfigOspfv2
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateProtocolConfigOspfv2, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Interfaces returns UpdateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIterIter, set in UpdateProtocolConfigOspfv2
	Interfaces() UpdateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter
	setNil()
}

// List of OSPFv2 interface update groups. One update group targets one or more OSPFv2 interfaces.
// Use multiple update groups to handle asymmetric changes across different subsets of interfaces.
//
// All interfaces listed in names in each update group receive every attribute update specified in the attributes list.
//
// If the session is up but true on-the-fly update is not supported for an attribute (e.g. metric change on an emulated interface),
// a warning is returned indicating that the session will be disabled and re-enabled, and the updated attribute will be reflected
// once the session comes back up.
//
// An error should be returned if any provided name is not found in the current configuration's list of OSPFv2 interface names.
// Interfaces returns a []UpdateProtocolConfigOspfv2InterfaceUpdateGroup
func (obj *updateProtocolConfigOspfv2) Interfaces() UpdateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter {
	if len(obj.obj.Interfaces) == 0 {
		obj.obj.Interfaces = []*otg.UpdateProtocolConfigOspfv2InterfaceUpdateGroup{}
	}
	if obj.interfacesHolder == nil {
		obj.interfacesHolder = newUpdateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter(&obj.obj.Interfaces).setMsg(obj)
	}
	return obj.interfacesHolder
}

type updateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter struct {
	obj                                                 *updateProtocolConfigOspfv2
	updateProtocolConfigOspfv2InterfaceUpdateGroupSlice []UpdateProtocolConfigOspfv2InterfaceUpdateGroup
	fieldPtr                                            *[]*otg.UpdateProtocolConfigOspfv2InterfaceUpdateGroup
}

func newUpdateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter(ptr *[]*otg.UpdateProtocolConfigOspfv2InterfaceUpdateGroup) UpdateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter {
	return &updateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter{fieldPtr: ptr}
}

type UpdateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter interface {
	setMsg(*updateProtocolConfigOspfv2) UpdateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter
	Items() []UpdateProtocolConfigOspfv2InterfaceUpdateGroup
	Add() UpdateProtocolConfigOspfv2InterfaceUpdateGroup
	Append(items ...UpdateProtocolConfigOspfv2InterfaceUpdateGroup) UpdateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter
	Set(index int, newObj UpdateProtocolConfigOspfv2InterfaceUpdateGroup) UpdateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter
	Clear() UpdateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter
	clearHolderSlice() UpdateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter
	appendHolderSlice(item UpdateProtocolConfigOspfv2InterfaceUpdateGroup) UpdateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter
}

func (obj *updateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter) setMsg(msg *updateProtocolConfigOspfv2) UpdateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&updateProtocolConfigOspfv2InterfaceUpdateGroup{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *updateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter) Items() []UpdateProtocolConfigOspfv2InterfaceUpdateGroup {
	return obj.updateProtocolConfigOspfv2InterfaceUpdateGroupSlice
}

func (obj *updateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter) Add() UpdateProtocolConfigOspfv2InterfaceUpdateGroup {
	newObj := &otg.UpdateProtocolConfigOspfv2InterfaceUpdateGroup{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &updateProtocolConfigOspfv2InterfaceUpdateGroup{obj: newObj}
	newLibObj.setDefault()
	obj.updateProtocolConfigOspfv2InterfaceUpdateGroupSlice = append(obj.updateProtocolConfigOspfv2InterfaceUpdateGroupSlice, newLibObj)
	return newLibObj
}

func (obj *updateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter) Append(items ...UpdateProtocolConfigOspfv2InterfaceUpdateGroup) UpdateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.updateProtocolConfigOspfv2InterfaceUpdateGroupSlice = append(obj.updateProtocolConfigOspfv2InterfaceUpdateGroupSlice, item)
	}
	return obj
}

func (obj *updateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter) Set(index int, newObj UpdateProtocolConfigOspfv2InterfaceUpdateGroup) UpdateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.updateProtocolConfigOspfv2InterfaceUpdateGroupSlice[index] = newObj
	return obj
}
func (obj *updateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter) Clear() UpdateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.UpdateProtocolConfigOspfv2InterfaceUpdateGroup{}
		obj.updateProtocolConfigOspfv2InterfaceUpdateGroupSlice = []UpdateProtocolConfigOspfv2InterfaceUpdateGroup{}
	}
	return obj
}
func (obj *updateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter) clearHolderSlice() UpdateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter {
	if len(obj.updateProtocolConfigOspfv2InterfaceUpdateGroupSlice) > 0 {
		obj.updateProtocolConfigOspfv2InterfaceUpdateGroupSlice = []UpdateProtocolConfigOspfv2InterfaceUpdateGroup{}
	}
	return obj
}
func (obj *updateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter) appendHolderSlice(item UpdateProtocolConfigOspfv2InterfaceUpdateGroup) UpdateProtocolConfigOspfv2UpdateProtocolConfigOspfv2InterfaceUpdateGroupIter {
	obj.updateProtocolConfigOspfv2InterfaceUpdateGroupSlice = append(obj.updateProtocolConfigOspfv2InterfaceUpdateGroupSlice, item)
	return obj
}

func (obj *updateProtocolConfigOspfv2) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Interfaces) != 0 {

		if set_default {
			obj.Interfaces().clearHolderSlice()
			for _, item := range obj.obj.Interfaces {
				obj.Interfaces().appendHolderSlice(&updateProtocolConfigOspfv2InterfaceUpdateGroup{obj: item})
			}
		}
		for _, item := range obj.Interfaces().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *updateProtocolConfigOspfv2) setDefault() {

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateProtocolConfigOspfv2InterfaceUpdateGroup *****
type updateProtocolConfigOspfv2InterfaceUpdateGroup struct {
	validation
	obj              *otg.UpdateProtocolConfigOspfv2InterfaceUpdateGroup
	marshaller       marshalUpdateProtocolConfigOspfv2InterfaceUpdateGroup
	unMarshaller     unMarshalUpdateProtocolConfigOspfv2InterfaceUpdateGroup
	attributesHolder UpdateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter
}

func NewUpdateProtocolConfigOspfv2InterfaceUpdateGroup() UpdateProtocolConfigOspfv2InterfaceUpdateGroup {
	obj := updateProtocolConfigOspfv2InterfaceUpdateGroup{obj: &otg.UpdateProtocolConfigOspfv2InterfaceUpdateGroup{}}
	obj.setDefault()
	return &obj
}

func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroup) msg() *otg.UpdateProtocolConfigOspfv2InterfaceUpdateGroup {
	return obj.obj
}

func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroup) setMsg(msg *otg.UpdateProtocolConfigOspfv2InterfaceUpdateGroup) UpdateProtocolConfigOspfv2InterfaceUpdateGroup {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateProtocolConfigOspfv2InterfaceUpdateGroup struct {
	obj *updateProtocolConfigOspfv2InterfaceUpdateGroup
}

type marshalUpdateProtocolConfigOspfv2InterfaceUpdateGroup interface {
	// ToProto marshals UpdateProtocolConfigOspfv2InterfaceUpdateGroup to protobuf object *otg.UpdateProtocolConfigOspfv2InterfaceUpdateGroup
	ToProto() (*otg.UpdateProtocolConfigOspfv2InterfaceUpdateGroup, error)
	// ToPbText marshals UpdateProtocolConfigOspfv2InterfaceUpdateGroup to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateProtocolConfigOspfv2InterfaceUpdateGroup to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateProtocolConfigOspfv2InterfaceUpdateGroup to JSON text
	ToJson() (string, error)
}

type unMarshalupdateProtocolConfigOspfv2InterfaceUpdateGroup struct {
	obj *updateProtocolConfigOspfv2InterfaceUpdateGroup
}

type unMarshalUpdateProtocolConfigOspfv2InterfaceUpdateGroup interface {
	// FromProto unmarshals UpdateProtocolConfigOspfv2InterfaceUpdateGroup from protobuf object *otg.UpdateProtocolConfigOspfv2InterfaceUpdateGroup
	FromProto(msg *otg.UpdateProtocolConfigOspfv2InterfaceUpdateGroup) (UpdateProtocolConfigOspfv2InterfaceUpdateGroup, error)
	// FromPbText unmarshals UpdateProtocolConfigOspfv2InterfaceUpdateGroup from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateProtocolConfigOspfv2InterfaceUpdateGroup from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateProtocolConfigOspfv2InterfaceUpdateGroup from JSON text
	FromJson(value string) error
}

func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroup) Marshal() marshalUpdateProtocolConfigOspfv2InterfaceUpdateGroup {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateProtocolConfigOspfv2InterfaceUpdateGroup{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroup) Unmarshal() unMarshalUpdateProtocolConfigOspfv2InterfaceUpdateGroup {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateProtocolConfigOspfv2InterfaceUpdateGroup{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateProtocolConfigOspfv2InterfaceUpdateGroup) ToProto() (*otg.UpdateProtocolConfigOspfv2InterfaceUpdateGroup, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateProtocolConfigOspfv2InterfaceUpdateGroup) FromProto(msg *otg.UpdateProtocolConfigOspfv2InterfaceUpdateGroup) (UpdateProtocolConfigOspfv2InterfaceUpdateGroup, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateProtocolConfigOspfv2InterfaceUpdateGroup) ToPbText() (string, error) {
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

func (m *unMarshalupdateProtocolConfigOspfv2InterfaceUpdateGroup) FromPbText(value string) error {
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

func (m *marshalupdateProtocolConfigOspfv2InterfaceUpdateGroup) ToYaml() (string, error) {
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

func (m *unMarshalupdateProtocolConfigOspfv2InterfaceUpdateGroup) FromYaml(value string) error {
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

func (m *marshalupdateProtocolConfigOspfv2InterfaceUpdateGroup) ToJson() (string, error) {
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

func (m *unMarshalupdateProtocolConfigOspfv2InterfaceUpdateGroup) FromJson(value string) error {
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

func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroup) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroup) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroup) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroup) Clone() (UpdateProtocolConfigOspfv2InterfaceUpdateGroup, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateProtocolConfigOspfv2InterfaceUpdateGroup()
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

func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroup) setNil() {
	obj.attributesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UpdateProtocolConfigOspfv2InterfaceUpdateGroup is an update group targeting one or more OSPFv2 interfaces. All interfaces listed in names receive every attribute update specified in the attributes list. Use multiple update groups to apply different updates to different subsets of interfaces (asymmetric updates).
// If the session is up but true on-the-fly update is not supported for an attribute (e.g. metric change on an emulated interface): a warning is returned indicating that the session will be disabled and re-enabled,
// and the updated attribute will be reflected once the session comes back up. An error should be returned if any provided name is not found in the current configuration's list of OSPFv2 interface names.
type UpdateProtocolConfigOspfv2InterfaceUpdateGroup interface {
	Validation
	// msg marshals UpdateProtocolConfigOspfv2InterfaceUpdateGroup to protobuf object *otg.UpdateProtocolConfigOspfv2InterfaceUpdateGroup
	// and doesn't set defaults
	msg() *otg.UpdateProtocolConfigOspfv2InterfaceUpdateGroup
	// setMsg unmarshals UpdateProtocolConfigOspfv2InterfaceUpdateGroup from protobuf object *otg.UpdateProtocolConfigOspfv2InterfaceUpdateGroup
	// and doesn't set defaults
	setMsg(*otg.UpdateProtocolConfigOspfv2InterfaceUpdateGroup) UpdateProtocolConfigOspfv2InterfaceUpdateGroup
	// provides marshal interface
	Marshal() marshalUpdateProtocolConfigOspfv2InterfaceUpdateGroup
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateProtocolConfigOspfv2InterfaceUpdateGroup
	// validate validates UpdateProtocolConfigOspfv2InterfaceUpdateGroup
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateProtocolConfigOspfv2InterfaceUpdateGroup, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Names returns []string, set in UpdateProtocolConfigOspfv2InterfaceUpdateGroup.
	Names() []string
	// SetNames assigns []string provided by user to UpdateProtocolConfigOspfv2InterfaceUpdateGroup
	SetNames(value []string) UpdateProtocolConfigOspfv2InterfaceUpdateGroup
	// Attributes returns UpdateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIterIter, set in UpdateProtocolConfigOspfv2InterfaceUpdateGroup
	Attributes() UpdateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter
	setNil()
}

// The names of the OSPFv2 interfaces to which all attribute updates in this group will be applied.
//
// x-constraint:
// - /components/schemas/Ospfv2.Interface/properties/name
//
// Names returns a []string
func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroup) Names() []string {
	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	return obj.obj.Names
}

// The names of the OSPFv2 interfaces to which all attribute updates in this group will be applied.
//
// x-constraint:
// - /components/schemas/Ospfv2.Interface/properties/name
//
// SetNames sets the []string value in the UpdateProtocolConfigOspfv2InterfaceUpdateGroup object
func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroup) SetNames(value []string) UpdateProtocolConfigOspfv2InterfaceUpdateGroup {

	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	obj.obj.Names = value

	return obj
}

// The list of interface attribute updates to apply. Each entry selects one attribute via the choice discriminator. Multiple attributes can be updated in a single group without repeating the names list.
// Attributes returns a []UpdateProtocolConfigOspfv2InterfaceAttribute
func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroup) Attributes() UpdateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter {
	if len(obj.obj.Attributes) == 0 {
		obj.obj.Attributes = []*otg.UpdateProtocolConfigOspfv2InterfaceAttribute{}
	}
	if obj.attributesHolder == nil {
		obj.attributesHolder = newUpdateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter(&obj.obj.Attributes).setMsg(obj)
	}
	return obj.attributesHolder
}

type updateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter struct {
	obj                                               *updateProtocolConfigOspfv2InterfaceUpdateGroup
	updateProtocolConfigOspfv2InterfaceAttributeSlice []UpdateProtocolConfigOspfv2InterfaceAttribute
	fieldPtr                                          *[]*otg.UpdateProtocolConfigOspfv2InterfaceAttribute
}

func newUpdateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter(ptr *[]*otg.UpdateProtocolConfigOspfv2InterfaceAttribute) UpdateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter {
	return &updateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter{fieldPtr: ptr}
}

type UpdateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter interface {
	setMsg(*updateProtocolConfigOspfv2InterfaceUpdateGroup) UpdateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter
	Items() []UpdateProtocolConfigOspfv2InterfaceAttribute
	Add() UpdateProtocolConfigOspfv2InterfaceAttribute
	Append(items ...UpdateProtocolConfigOspfv2InterfaceAttribute) UpdateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter
	Set(index int, newObj UpdateProtocolConfigOspfv2InterfaceAttribute) UpdateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter
	Clear() UpdateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter
	clearHolderSlice() UpdateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter
	appendHolderSlice(item UpdateProtocolConfigOspfv2InterfaceAttribute) UpdateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter
}

func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter) setMsg(msg *updateProtocolConfigOspfv2InterfaceUpdateGroup) UpdateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&updateProtocolConfigOspfv2InterfaceAttribute{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter) Items() []UpdateProtocolConfigOspfv2InterfaceAttribute {
	return obj.updateProtocolConfigOspfv2InterfaceAttributeSlice
}

func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter) Add() UpdateProtocolConfigOspfv2InterfaceAttribute {
	newObj := &otg.UpdateProtocolConfigOspfv2InterfaceAttribute{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &updateProtocolConfigOspfv2InterfaceAttribute{obj: newObj}
	newLibObj.setDefault()
	obj.updateProtocolConfigOspfv2InterfaceAttributeSlice = append(obj.updateProtocolConfigOspfv2InterfaceAttributeSlice, newLibObj)
	return newLibObj
}

func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter) Append(items ...UpdateProtocolConfigOspfv2InterfaceAttribute) UpdateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.updateProtocolConfigOspfv2InterfaceAttributeSlice = append(obj.updateProtocolConfigOspfv2InterfaceAttributeSlice, item)
	}
	return obj
}

func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter) Set(index int, newObj UpdateProtocolConfigOspfv2InterfaceAttribute) UpdateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.updateProtocolConfigOspfv2InterfaceAttributeSlice[index] = newObj
	return obj
}
func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter) Clear() UpdateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.UpdateProtocolConfigOspfv2InterfaceAttribute{}
		obj.updateProtocolConfigOspfv2InterfaceAttributeSlice = []UpdateProtocolConfigOspfv2InterfaceAttribute{}
	}
	return obj
}
func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter) clearHolderSlice() UpdateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter {
	if len(obj.updateProtocolConfigOspfv2InterfaceAttributeSlice) > 0 {
		obj.updateProtocolConfigOspfv2InterfaceAttributeSlice = []UpdateProtocolConfigOspfv2InterfaceAttribute{}
	}
	return obj
}
func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter) appendHolderSlice(item UpdateProtocolConfigOspfv2InterfaceAttribute) UpdateProtocolConfigOspfv2InterfaceUpdateGroupUpdateProtocolConfigOspfv2InterfaceAttributeIter {
	obj.updateProtocolConfigOspfv2InterfaceAttributeSlice = append(obj.updateProtocolConfigOspfv2InterfaceAttributeSlice, item)
	return obj
}

func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroup) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Attributes) != 0 {

		if set_default {
			obj.Attributes().clearHolderSlice()
			for _, item := range obj.obj.Attributes {
				obj.Attributes().appendHolderSlice(&updateProtocolConfigOspfv2InterfaceAttribute{obj: item})
			}
		}
		for _, item := range obj.Attributes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *updateProtocolConfigOspfv2InterfaceUpdateGroup) setDefault() {

}

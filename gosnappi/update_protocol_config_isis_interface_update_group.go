package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateProtocolConfigIsisInterfaceUpdateGroup *****
type updateProtocolConfigIsisInterfaceUpdateGroup struct {
	validation
	obj              *otg.UpdateProtocolConfigIsisInterfaceUpdateGroup
	marshaller       marshalUpdateProtocolConfigIsisInterfaceUpdateGroup
	unMarshaller     unMarshalUpdateProtocolConfigIsisInterfaceUpdateGroup
	attributesHolder UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter
}

func NewUpdateProtocolConfigIsisInterfaceUpdateGroup() UpdateProtocolConfigIsisInterfaceUpdateGroup {
	obj := updateProtocolConfigIsisInterfaceUpdateGroup{obj: &otg.UpdateProtocolConfigIsisInterfaceUpdateGroup{}}
	obj.setDefault()
	return &obj
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) msg() *otg.UpdateProtocolConfigIsisInterfaceUpdateGroup {
	return obj.obj
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) setMsg(msg *otg.UpdateProtocolConfigIsisInterfaceUpdateGroup) UpdateProtocolConfigIsisInterfaceUpdateGroup {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateProtocolConfigIsisInterfaceUpdateGroup struct {
	obj *updateProtocolConfigIsisInterfaceUpdateGroup
}

type marshalUpdateProtocolConfigIsisInterfaceUpdateGroup interface {
	// ToProto marshals UpdateProtocolConfigIsisInterfaceUpdateGroup to protobuf object *otg.UpdateProtocolConfigIsisInterfaceUpdateGroup
	ToProto() (*otg.UpdateProtocolConfigIsisInterfaceUpdateGroup, error)
	// ToPbText marshals UpdateProtocolConfigIsisInterfaceUpdateGroup to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateProtocolConfigIsisInterfaceUpdateGroup to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateProtocolConfigIsisInterfaceUpdateGroup to JSON text
	ToJson() (string, error)
}

type unMarshalupdateProtocolConfigIsisInterfaceUpdateGroup struct {
	obj *updateProtocolConfigIsisInterfaceUpdateGroup
}

type unMarshalUpdateProtocolConfigIsisInterfaceUpdateGroup interface {
	// FromProto unmarshals UpdateProtocolConfigIsisInterfaceUpdateGroup from protobuf object *otg.UpdateProtocolConfigIsisInterfaceUpdateGroup
	FromProto(msg *otg.UpdateProtocolConfigIsisInterfaceUpdateGroup) (UpdateProtocolConfigIsisInterfaceUpdateGroup, error)
	// FromPbText unmarshals UpdateProtocolConfigIsisInterfaceUpdateGroup from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateProtocolConfigIsisInterfaceUpdateGroup from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateProtocolConfigIsisInterfaceUpdateGroup from JSON text
	FromJson(value string) error
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) Marshal() marshalUpdateProtocolConfigIsisInterfaceUpdateGroup {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateProtocolConfigIsisInterfaceUpdateGroup{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) Unmarshal() unMarshalUpdateProtocolConfigIsisInterfaceUpdateGroup {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateProtocolConfigIsisInterfaceUpdateGroup{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateProtocolConfigIsisInterfaceUpdateGroup) ToProto() (*otg.UpdateProtocolConfigIsisInterfaceUpdateGroup, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateProtocolConfigIsisInterfaceUpdateGroup) FromProto(msg *otg.UpdateProtocolConfigIsisInterfaceUpdateGroup) (UpdateProtocolConfigIsisInterfaceUpdateGroup, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateProtocolConfigIsisInterfaceUpdateGroup) ToPbText() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisInterfaceUpdateGroup) FromPbText(value string) error {
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

func (m *marshalupdateProtocolConfigIsisInterfaceUpdateGroup) ToYaml() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisInterfaceUpdateGroup) FromYaml(value string) error {
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

func (m *marshalupdateProtocolConfigIsisInterfaceUpdateGroup) ToJson() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisInterfaceUpdateGroup) FromJson(value string) error {
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

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) Clone() (UpdateProtocolConfigIsisInterfaceUpdateGroup, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateProtocolConfigIsisInterfaceUpdateGroup()
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

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) setNil() {
	obj.attributesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UpdateProtocolConfigIsisInterfaceUpdateGroup is an update group targeting one or more IS-IS interfaces. All interfaces listed in names receive every attribute update specified in the attributes list. Use multiple update groups to apply different updates to different subsets of interfaces (asymmetric updates).
// If the session is up but true on-the-fly update is not supported for an attribute (e.g. metric change on an emulated interface): a warning is returned indicating that the session will be disabled and re-enabled, and the updated attribute will be reflected once the session comes back up.
type UpdateProtocolConfigIsisInterfaceUpdateGroup interface {
	Validation
	// msg marshals UpdateProtocolConfigIsisInterfaceUpdateGroup to protobuf object *otg.UpdateProtocolConfigIsisInterfaceUpdateGroup
	// and doesn't set defaults
	msg() *otg.UpdateProtocolConfigIsisInterfaceUpdateGroup
	// setMsg unmarshals UpdateProtocolConfigIsisInterfaceUpdateGroup from protobuf object *otg.UpdateProtocolConfigIsisInterfaceUpdateGroup
	// and doesn't set defaults
	setMsg(*otg.UpdateProtocolConfigIsisInterfaceUpdateGroup) UpdateProtocolConfigIsisInterfaceUpdateGroup
	// provides marshal interface
	Marshal() marshalUpdateProtocolConfigIsisInterfaceUpdateGroup
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateProtocolConfigIsisInterfaceUpdateGroup
	// validate validates UpdateProtocolConfigIsisInterfaceUpdateGroup
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateProtocolConfigIsisInterfaceUpdateGroup, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Names returns []string, set in UpdateProtocolConfigIsisInterfaceUpdateGroup.
	Names() []string
	// SetNames assigns []string provided by user to UpdateProtocolConfigIsisInterfaceUpdateGroup
	SetNames(value []string) UpdateProtocolConfigIsisInterfaceUpdateGroup
	// Attributes returns UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIterIter, set in UpdateProtocolConfigIsisInterfaceUpdateGroup
	Attributes() UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter
	setNil()
}

// The names of the IS-IS interfaces to which all attribute updates in this group will be applied.
// Names returns a []string
func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) Names() []string {
	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	return obj.obj.Names
}

// The names of the IS-IS interfaces to which all attribute updates in this group will be applied.
// SetNames sets the []string value in the UpdateProtocolConfigIsisInterfaceUpdateGroup object
func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) SetNames(value []string) UpdateProtocolConfigIsisInterfaceUpdateGroup {

	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	obj.obj.Names = value

	return obj
}

// The list of interface attribute updates to apply. Each entry selects one attribute via the choice discriminator. Multiple attributes can be updated in a single group without repeating the names list.
// Attributes returns a []UpdateProtocolConfigIsisInterfaceAttribute
func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) Attributes() UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter {
	if len(obj.obj.Attributes) == 0 {
		obj.obj.Attributes = []*otg.UpdateProtocolConfigIsisInterfaceAttribute{}
	}
	if obj.attributesHolder == nil {
		obj.attributesHolder = newUpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter(&obj.obj.Attributes).setMsg(obj)
	}
	return obj.attributesHolder
}

type updateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter struct {
	obj                                             *updateProtocolConfigIsisInterfaceUpdateGroup
	updateProtocolConfigIsisInterfaceAttributeSlice []UpdateProtocolConfigIsisInterfaceAttribute
	fieldPtr                                        *[]*otg.UpdateProtocolConfigIsisInterfaceAttribute
}

func newUpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter(ptr *[]*otg.UpdateProtocolConfigIsisInterfaceAttribute) UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter {
	return &updateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter{fieldPtr: ptr}
}

type UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter interface {
	setMsg(*updateProtocolConfigIsisInterfaceUpdateGroup) UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter
	Items() []UpdateProtocolConfigIsisInterfaceAttribute
	Add() UpdateProtocolConfigIsisInterfaceAttribute
	Append(items ...UpdateProtocolConfigIsisInterfaceAttribute) UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter
	Set(index int, newObj UpdateProtocolConfigIsisInterfaceAttribute) UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter
	Clear() UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter
	clearHolderSlice() UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter
	appendHolderSlice(item UpdateProtocolConfigIsisInterfaceAttribute) UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter) setMsg(msg *updateProtocolConfigIsisInterfaceUpdateGroup) UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&updateProtocolConfigIsisInterfaceAttribute{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter) Items() []UpdateProtocolConfigIsisInterfaceAttribute {
	return obj.updateProtocolConfigIsisInterfaceAttributeSlice
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter) Add() UpdateProtocolConfigIsisInterfaceAttribute {
	newObj := &otg.UpdateProtocolConfigIsisInterfaceAttribute{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &updateProtocolConfigIsisInterfaceAttribute{obj: newObj}
	newLibObj.setDefault()
	obj.updateProtocolConfigIsisInterfaceAttributeSlice = append(obj.updateProtocolConfigIsisInterfaceAttributeSlice, newLibObj)
	return newLibObj
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter) Append(items ...UpdateProtocolConfigIsisInterfaceAttribute) UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.updateProtocolConfigIsisInterfaceAttributeSlice = append(obj.updateProtocolConfigIsisInterfaceAttributeSlice, item)
	}
	return obj
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter) Set(index int, newObj UpdateProtocolConfigIsisInterfaceAttribute) UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.updateProtocolConfigIsisInterfaceAttributeSlice[index] = newObj
	return obj
}
func (obj *updateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter) Clear() UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.UpdateProtocolConfigIsisInterfaceAttribute{}
		obj.updateProtocolConfigIsisInterfaceAttributeSlice = []UpdateProtocolConfigIsisInterfaceAttribute{}
	}
	return obj
}
func (obj *updateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter) clearHolderSlice() UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter {
	if len(obj.updateProtocolConfigIsisInterfaceAttributeSlice) > 0 {
		obj.updateProtocolConfigIsisInterfaceAttributeSlice = []UpdateProtocolConfigIsisInterfaceAttribute{}
	}
	return obj
}
func (obj *updateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter) appendHolderSlice(item UpdateProtocolConfigIsisInterfaceAttribute) UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfaceAttributeIter {
	obj.updateProtocolConfigIsisInterfaceAttributeSlice = append(obj.updateProtocolConfigIsisInterfaceAttributeSlice, item)
	return obj
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Attributes) != 0 {

		if set_default {
			obj.Attributes().clearHolderSlice()
			for _, item := range obj.obj.Attributes {
				obj.Attributes().appendHolderSlice(&updateProtocolConfigIsisInterfaceAttribute{obj: item})
			}
		}
		for _, item := range obj.Attributes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) setDefault() {

}

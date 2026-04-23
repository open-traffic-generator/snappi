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
	propertiesHolder UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter
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
	obj.propertiesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UpdateProtocolConfigIsisInterfaceUpdateGroup is an update group that applies a common set of property updates to one or more IS-IS interfaces. All interfaces listed in names receive the same property updates defined in properties. Use multiple update groups in a single request to apply different property values to different subsets of interfaces (asymmetric updates).
// A property must not appear more than once within the properties list of a single update group.
// The warnings field in the response is populated in the following cases:
// a) Session started but not yet up: the updates are accepted and will be
// applied when the session comes up.
//
// b) Session is up but true on-the-fly update is not supported for the
// attribute (e.g. metric change on an emulated interface): a warning is
// returned indicating that the session will be disabled and re-enabled,
// and the updated attribute will be reflected once the session comes back up.
//
// An error is returned in the following cases:
// a) The protocol or session is not in a started state.
// b) A referenced interface name does not exist in the configuration.
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
	// Properties returns UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIterIter, set in UpdateProtocolConfigIsisInterfaceUpdateGroup
	Properties() UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter
	setNil()
}

// The names of the IS-IS interfaces to which all property updates in this group will be applied.
// Names returns a []string
func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) Names() []string {
	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	return obj.obj.Names
}

// The names of the IS-IS interfaces to which all property updates in this group will be applied.
// SetNames sets the []string value in the UpdateProtocolConfigIsisInterfaceUpdateGroup object
func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) SetNames(value []string) UpdateProtocolConfigIsisInterfaceUpdateGroup {

	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	obj.obj.Names = value

	return obj
}

// The list of property updates to apply to all named interfaces. Each property may appear at most once within this list.
// Properties returns a []UpdateProtocolConfigIsisInterfaceProperty
func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) Properties() UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter {
	if len(obj.obj.Properties) == 0 {
		obj.obj.Properties = []*otg.UpdateProtocolConfigIsisInterfaceProperty{}
	}
	if obj.propertiesHolder == nil {
		obj.propertiesHolder = newUpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter(&obj.obj.Properties).setMsg(obj)
	}
	return obj.propertiesHolder
}

type updateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter struct {
	obj                                            *updateProtocolConfigIsisInterfaceUpdateGroup
	updateProtocolConfigIsisInterfacePropertySlice []UpdateProtocolConfigIsisInterfaceProperty
	fieldPtr                                       *[]*otg.UpdateProtocolConfigIsisInterfaceProperty
}

func newUpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter(ptr *[]*otg.UpdateProtocolConfigIsisInterfaceProperty) UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter {
	return &updateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter{fieldPtr: ptr}
}

type UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter interface {
	setMsg(*updateProtocolConfigIsisInterfaceUpdateGroup) UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter
	Items() []UpdateProtocolConfigIsisInterfaceProperty
	Add() UpdateProtocolConfigIsisInterfaceProperty
	Append(items ...UpdateProtocolConfigIsisInterfaceProperty) UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter
	Set(index int, newObj UpdateProtocolConfigIsisInterfaceProperty) UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter
	Clear() UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter
	clearHolderSlice() UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter
	appendHolderSlice(item UpdateProtocolConfigIsisInterfaceProperty) UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter) setMsg(msg *updateProtocolConfigIsisInterfaceUpdateGroup) UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&updateProtocolConfigIsisInterfaceProperty{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter) Items() []UpdateProtocolConfigIsisInterfaceProperty {
	return obj.updateProtocolConfigIsisInterfacePropertySlice
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter) Add() UpdateProtocolConfigIsisInterfaceProperty {
	newObj := &otg.UpdateProtocolConfigIsisInterfaceProperty{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &updateProtocolConfigIsisInterfaceProperty{obj: newObj}
	newLibObj.setDefault()
	obj.updateProtocolConfigIsisInterfacePropertySlice = append(obj.updateProtocolConfigIsisInterfacePropertySlice, newLibObj)
	return newLibObj
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter) Append(items ...UpdateProtocolConfigIsisInterfaceProperty) UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.updateProtocolConfigIsisInterfacePropertySlice = append(obj.updateProtocolConfigIsisInterfacePropertySlice, item)
	}
	return obj
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter) Set(index int, newObj UpdateProtocolConfigIsisInterfaceProperty) UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.updateProtocolConfigIsisInterfacePropertySlice[index] = newObj
	return obj
}
func (obj *updateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter) Clear() UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.UpdateProtocolConfigIsisInterfaceProperty{}
		obj.updateProtocolConfigIsisInterfacePropertySlice = []UpdateProtocolConfigIsisInterfaceProperty{}
	}
	return obj
}
func (obj *updateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter) clearHolderSlice() UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter {
	if len(obj.updateProtocolConfigIsisInterfacePropertySlice) > 0 {
		obj.updateProtocolConfigIsisInterfacePropertySlice = []UpdateProtocolConfigIsisInterfaceProperty{}
	}
	return obj
}
func (obj *updateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter) appendHolderSlice(item UpdateProtocolConfigIsisInterfaceProperty) UpdateProtocolConfigIsisInterfaceUpdateGroupUpdateProtocolConfigIsisInterfacePropertyIter {
	obj.updateProtocolConfigIsisInterfacePropertySlice = append(obj.updateProtocolConfigIsisInterfacePropertySlice, item)
	return obj
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Properties) != 0 {

		if set_default {
			obj.Properties().clearHolderSlice()
			for _, item := range obj.obj.Properties {
				obj.Properties().appendHolderSlice(&updateProtocolConfigIsisInterfaceProperty{obj: item})
			}
		}
		for _, item := range obj.Properties().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) setDefault() {

}

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
	obj          *otg.UpdateProtocolConfigIsis
	marshaller   marshalUpdateProtocolConfigIsis
	unMarshaller unMarshalUpdateProtocolConfigIsis
	linksHolder  UpdateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter
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
	obj.linksHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UpdateProtocolConfigIsis is a container of IS-IS with associated properties to be updated that may or may not able to maintain the current IS-IS Up session state always.
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
	// Links returns UpdateProtocolConfigIsisUpdateProtocolConfigIsisLinkIterIter, set in UpdateProtocolConfigIsis
	Links() UpdateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter
	setNil()
}

// The flows that will be added to existing configuration on the traffic generator.
// The flow name must not already be used in existing configuration.
// Links returns a []UpdateProtocolConfigIsisLink
func (obj *updateProtocolConfigIsis) Links() UpdateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter {
	if len(obj.obj.Links) == 0 {
		obj.obj.Links = []*otg.UpdateProtocolConfigIsisLink{}
	}
	if obj.linksHolder == nil {
		obj.linksHolder = newUpdateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter(&obj.obj.Links).setMsg(obj)
	}
	return obj.linksHolder
}

type updateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter struct {
	obj                               *updateProtocolConfigIsis
	updateProtocolConfigIsisLinkSlice []UpdateProtocolConfigIsisLink
	fieldPtr                          *[]*otg.UpdateProtocolConfigIsisLink
}

func newUpdateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter(ptr *[]*otg.UpdateProtocolConfigIsisLink) UpdateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter {
	return &updateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter{fieldPtr: ptr}
}

type UpdateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter interface {
	setMsg(*updateProtocolConfigIsis) UpdateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter
	Items() []UpdateProtocolConfigIsisLink
	Add() UpdateProtocolConfigIsisLink
	Append(items ...UpdateProtocolConfigIsisLink) UpdateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter
	Set(index int, newObj UpdateProtocolConfigIsisLink) UpdateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter
	Clear() UpdateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter
	clearHolderSlice() UpdateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter
	appendHolderSlice(item UpdateProtocolConfigIsisLink) UpdateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter
}

func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter) setMsg(msg *updateProtocolConfigIsis) UpdateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&updateProtocolConfigIsisLink{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter) Items() []UpdateProtocolConfigIsisLink {
	return obj.updateProtocolConfigIsisLinkSlice
}

func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter) Add() UpdateProtocolConfigIsisLink {
	newObj := &otg.UpdateProtocolConfigIsisLink{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &updateProtocolConfigIsisLink{obj: newObj}
	newLibObj.setDefault()
	obj.updateProtocolConfigIsisLinkSlice = append(obj.updateProtocolConfigIsisLinkSlice, newLibObj)
	return newLibObj
}

func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter) Append(items ...UpdateProtocolConfigIsisLink) UpdateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.updateProtocolConfigIsisLinkSlice = append(obj.updateProtocolConfigIsisLinkSlice, item)
	}
	return obj
}

func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter) Set(index int, newObj UpdateProtocolConfigIsisLink) UpdateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.updateProtocolConfigIsisLinkSlice[index] = newObj
	return obj
}
func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter) Clear() UpdateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.UpdateProtocolConfigIsisLink{}
		obj.updateProtocolConfigIsisLinkSlice = []UpdateProtocolConfigIsisLink{}
	}
	return obj
}
func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter) clearHolderSlice() UpdateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter {
	if len(obj.updateProtocolConfigIsisLinkSlice) > 0 {
		obj.updateProtocolConfigIsisLinkSlice = []UpdateProtocolConfigIsisLink{}
	}
	return obj
}
func (obj *updateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter) appendHolderSlice(item UpdateProtocolConfigIsisLink) UpdateProtocolConfigIsisUpdateProtocolConfigIsisLinkIter {
	obj.updateProtocolConfigIsisLinkSlice = append(obj.updateProtocolConfigIsisLinkSlice, item)
	return obj
}

func (obj *updateProtocolConfigIsis) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Links) != 0 {

		if set_default {
			obj.Links().clearHolderSlice()
			for _, item := range obj.obj.Links {
				obj.Links().appendHolderSlice(&updateProtocolConfigIsisLink{obj: item})
			}
		}
		for _, item := range obj.Links().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *updateProtocolConfigIsis) setDefault() {

}

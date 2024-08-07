package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** VxlanV4TunnelDestinationIPModeUnicast *****
type vxlanV4TunnelDestinationIPModeUnicast struct {
	validation
	obj          *otg.VxlanV4TunnelDestinationIPModeUnicast
	marshaller   marshalVxlanV4TunnelDestinationIPModeUnicast
	unMarshaller unMarshalVxlanV4TunnelDestinationIPModeUnicast
	vtepsHolder  VxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter
}

func NewVxlanV4TunnelDestinationIPModeUnicast() VxlanV4TunnelDestinationIPModeUnicast {
	obj := vxlanV4TunnelDestinationIPModeUnicast{obj: &otg.VxlanV4TunnelDestinationIPModeUnicast{}}
	obj.setDefault()
	return &obj
}

func (obj *vxlanV4TunnelDestinationIPModeUnicast) msg() *otg.VxlanV4TunnelDestinationIPModeUnicast {
	return obj.obj
}

func (obj *vxlanV4TunnelDestinationIPModeUnicast) setMsg(msg *otg.VxlanV4TunnelDestinationIPModeUnicast) VxlanV4TunnelDestinationIPModeUnicast {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalvxlanV4TunnelDestinationIPModeUnicast struct {
	obj *vxlanV4TunnelDestinationIPModeUnicast
}

type marshalVxlanV4TunnelDestinationIPModeUnicast interface {
	// ToProto marshals VxlanV4TunnelDestinationIPModeUnicast to protobuf object *otg.VxlanV4TunnelDestinationIPModeUnicast
	ToProto() (*otg.VxlanV4TunnelDestinationIPModeUnicast, error)
	// ToPbText marshals VxlanV4TunnelDestinationIPModeUnicast to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals VxlanV4TunnelDestinationIPModeUnicast to YAML text
	ToYaml() (string, error)
	// ToJson marshals VxlanV4TunnelDestinationIPModeUnicast to JSON text
	ToJson() (string, error)
}

type unMarshalvxlanV4TunnelDestinationIPModeUnicast struct {
	obj *vxlanV4TunnelDestinationIPModeUnicast
}

type unMarshalVxlanV4TunnelDestinationIPModeUnicast interface {
	// FromProto unmarshals VxlanV4TunnelDestinationIPModeUnicast from protobuf object *otg.VxlanV4TunnelDestinationIPModeUnicast
	FromProto(msg *otg.VxlanV4TunnelDestinationIPModeUnicast) (VxlanV4TunnelDestinationIPModeUnicast, error)
	// FromPbText unmarshals VxlanV4TunnelDestinationIPModeUnicast from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals VxlanV4TunnelDestinationIPModeUnicast from YAML text
	FromYaml(value string) error
	// FromJson unmarshals VxlanV4TunnelDestinationIPModeUnicast from JSON text
	FromJson(value string) error
}

func (obj *vxlanV4TunnelDestinationIPModeUnicast) Marshal() marshalVxlanV4TunnelDestinationIPModeUnicast {
	if obj.marshaller == nil {
		obj.marshaller = &marshalvxlanV4TunnelDestinationIPModeUnicast{obj: obj}
	}
	return obj.marshaller
}

func (obj *vxlanV4TunnelDestinationIPModeUnicast) Unmarshal() unMarshalVxlanV4TunnelDestinationIPModeUnicast {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalvxlanV4TunnelDestinationIPModeUnicast{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalvxlanV4TunnelDestinationIPModeUnicast) ToProto() (*otg.VxlanV4TunnelDestinationIPModeUnicast, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalvxlanV4TunnelDestinationIPModeUnicast) FromProto(msg *otg.VxlanV4TunnelDestinationIPModeUnicast) (VxlanV4TunnelDestinationIPModeUnicast, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalvxlanV4TunnelDestinationIPModeUnicast) ToPbText() (string, error) {
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

func (m *unMarshalvxlanV4TunnelDestinationIPModeUnicast) FromPbText(value string) error {
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

func (m *marshalvxlanV4TunnelDestinationIPModeUnicast) ToYaml() (string, error) {
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

func (m *unMarshalvxlanV4TunnelDestinationIPModeUnicast) FromYaml(value string) error {
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

func (m *marshalvxlanV4TunnelDestinationIPModeUnicast) ToJson() (string, error) {
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

func (m *unMarshalvxlanV4TunnelDestinationIPModeUnicast) FromJson(value string) error {
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

func (obj *vxlanV4TunnelDestinationIPModeUnicast) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *vxlanV4TunnelDestinationIPModeUnicast) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *vxlanV4TunnelDestinationIPModeUnicast) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *vxlanV4TunnelDestinationIPModeUnicast) Clone() (VxlanV4TunnelDestinationIPModeUnicast, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewVxlanV4TunnelDestinationIPModeUnicast()
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

func (obj *vxlanV4TunnelDestinationIPModeUnicast) setNil() {
	obj.vtepsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// VxlanV4TunnelDestinationIPModeUnicast is description is TBD
type VxlanV4TunnelDestinationIPModeUnicast interface {
	Validation
	// msg marshals VxlanV4TunnelDestinationIPModeUnicast to protobuf object *otg.VxlanV4TunnelDestinationIPModeUnicast
	// and doesn't set defaults
	msg() *otg.VxlanV4TunnelDestinationIPModeUnicast
	// setMsg unmarshals VxlanV4TunnelDestinationIPModeUnicast from protobuf object *otg.VxlanV4TunnelDestinationIPModeUnicast
	// and doesn't set defaults
	setMsg(*otg.VxlanV4TunnelDestinationIPModeUnicast) VxlanV4TunnelDestinationIPModeUnicast
	// provides marshal interface
	Marshal() marshalVxlanV4TunnelDestinationIPModeUnicast
	// provides unmarshal interface
	Unmarshal() unMarshalVxlanV4TunnelDestinationIPModeUnicast
	// validate validates VxlanV4TunnelDestinationIPModeUnicast
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (VxlanV4TunnelDestinationIPModeUnicast, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Vteps returns VxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIterIter, set in VxlanV4TunnelDestinationIPModeUnicast
	Vteps() VxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter
	setNil()
}

// List of VTEPs for member VNI(VXLAN Network Identifier)
// Vteps returns a []VxlanV4TunnelDestinationIPModeUnicastVtep
func (obj *vxlanV4TunnelDestinationIPModeUnicast) Vteps() VxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter {
	if len(obj.obj.Vteps) == 0 {
		obj.obj.Vteps = []*otg.VxlanV4TunnelDestinationIPModeUnicastVtep{}
	}
	if obj.vtepsHolder == nil {
		obj.vtepsHolder = newVxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter(&obj.obj.Vteps).setMsg(obj)
	}
	return obj.vtepsHolder
}

type vxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter struct {
	obj                                            *vxlanV4TunnelDestinationIPModeUnicast
	vxlanV4TunnelDestinationIPModeUnicastVtepSlice []VxlanV4TunnelDestinationIPModeUnicastVtep
	fieldPtr                                       *[]*otg.VxlanV4TunnelDestinationIPModeUnicastVtep
}

func newVxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter(ptr *[]*otg.VxlanV4TunnelDestinationIPModeUnicastVtep) VxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter {
	return &vxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter{fieldPtr: ptr}
}

type VxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter interface {
	setMsg(*vxlanV4TunnelDestinationIPModeUnicast) VxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter
	Items() []VxlanV4TunnelDestinationIPModeUnicastVtep
	Add() VxlanV4TunnelDestinationIPModeUnicastVtep
	Append(items ...VxlanV4TunnelDestinationIPModeUnicastVtep) VxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter
	Set(index int, newObj VxlanV4TunnelDestinationIPModeUnicastVtep) VxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter
	Clear() VxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter
	clearHolderSlice() VxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter
	appendHolderSlice(item VxlanV4TunnelDestinationIPModeUnicastVtep) VxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter
}

func (obj *vxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter) setMsg(msg *vxlanV4TunnelDestinationIPModeUnicast) VxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&vxlanV4TunnelDestinationIPModeUnicastVtep{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *vxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter) Items() []VxlanV4TunnelDestinationIPModeUnicastVtep {
	return obj.vxlanV4TunnelDestinationIPModeUnicastVtepSlice
}

func (obj *vxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter) Add() VxlanV4TunnelDestinationIPModeUnicastVtep {
	newObj := &otg.VxlanV4TunnelDestinationIPModeUnicastVtep{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &vxlanV4TunnelDestinationIPModeUnicastVtep{obj: newObj}
	newLibObj.setDefault()
	obj.vxlanV4TunnelDestinationIPModeUnicastVtepSlice = append(obj.vxlanV4TunnelDestinationIPModeUnicastVtepSlice, newLibObj)
	return newLibObj
}

func (obj *vxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter) Append(items ...VxlanV4TunnelDestinationIPModeUnicastVtep) VxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.vxlanV4TunnelDestinationIPModeUnicastVtepSlice = append(obj.vxlanV4TunnelDestinationIPModeUnicastVtepSlice, item)
	}
	return obj
}

func (obj *vxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter) Set(index int, newObj VxlanV4TunnelDestinationIPModeUnicastVtep) VxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.vxlanV4TunnelDestinationIPModeUnicastVtepSlice[index] = newObj
	return obj
}
func (obj *vxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter) Clear() VxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.VxlanV4TunnelDestinationIPModeUnicastVtep{}
		obj.vxlanV4TunnelDestinationIPModeUnicastVtepSlice = []VxlanV4TunnelDestinationIPModeUnicastVtep{}
	}
	return obj
}
func (obj *vxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter) clearHolderSlice() VxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter {
	if len(obj.vxlanV4TunnelDestinationIPModeUnicastVtepSlice) > 0 {
		obj.vxlanV4TunnelDestinationIPModeUnicastVtepSlice = []VxlanV4TunnelDestinationIPModeUnicastVtep{}
	}
	return obj
}
func (obj *vxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter) appendHolderSlice(item VxlanV4TunnelDestinationIPModeUnicastVtep) VxlanV4TunnelDestinationIPModeUnicastVxlanV4TunnelDestinationIPModeUnicastVtepIter {
	obj.vxlanV4TunnelDestinationIPModeUnicastVtepSlice = append(obj.vxlanV4TunnelDestinationIPModeUnicastVtepSlice, item)
	return obj
}

func (obj *vxlanV4TunnelDestinationIPModeUnicast) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Vteps) != 0 {

		if set_default {
			obj.Vteps().clearHolderSlice()
			for _, item := range obj.obj.Vteps {
				obj.Vteps().appendHolderSlice(&vxlanV4TunnelDestinationIPModeUnicastVtep{obj: item})
			}
		}
		for _, item := range obj.Vteps().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *vxlanV4TunnelDestinationIPModeUnicast) setDefault() {

}

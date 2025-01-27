package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** VxlanV6TunnelDestinationIPModeUnicast *****
type vxlanV6TunnelDestinationIPModeUnicast struct {
	validation
	obj          *otg.VxlanV6TunnelDestinationIPModeUnicast
	marshaller   marshalVxlanV6TunnelDestinationIPModeUnicast
	unMarshaller unMarshalVxlanV6TunnelDestinationIPModeUnicast
	vtepsHolder  VxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter
}

func NewVxlanV6TunnelDestinationIPModeUnicast() VxlanV6TunnelDestinationIPModeUnicast {
	obj := vxlanV6TunnelDestinationIPModeUnicast{obj: &otg.VxlanV6TunnelDestinationIPModeUnicast{}}
	obj.setDefault()
	return &obj
}

func (obj *vxlanV6TunnelDestinationIPModeUnicast) msg() *otg.VxlanV6TunnelDestinationIPModeUnicast {
	return obj.obj
}

func (obj *vxlanV6TunnelDestinationIPModeUnicast) setMsg(msg *otg.VxlanV6TunnelDestinationIPModeUnicast) VxlanV6TunnelDestinationIPModeUnicast {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalvxlanV6TunnelDestinationIPModeUnicast struct {
	obj *vxlanV6TunnelDestinationIPModeUnicast
}

type marshalVxlanV6TunnelDestinationIPModeUnicast interface {
	// ToProto marshals VxlanV6TunnelDestinationIPModeUnicast to protobuf object *otg.VxlanV6TunnelDestinationIPModeUnicast
	ToProto() (*otg.VxlanV6TunnelDestinationIPModeUnicast, error)
	// ToPbText marshals VxlanV6TunnelDestinationIPModeUnicast to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals VxlanV6TunnelDestinationIPModeUnicast to YAML text
	ToYaml() (string, error)
	// ToJson marshals VxlanV6TunnelDestinationIPModeUnicast to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals VxlanV6TunnelDestinationIPModeUnicast to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalvxlanV6TunnelDestinationIPModeUnicast struct {
	obj *vxlanV6TunnelDestinationIPModeUnicast
}

type unMarshalVxlanV6TunnelDestinationIPModeUnicast interface {
	// FromProto unmarshals VxlanV6TunnelDestinationIPModeUnicast from protobuf object *otg.VxlanV6TunnelDestinationIPModeUnicast
	FromProto(msg *otg.VxlanV6TunnelDestinationIPModeUnicast) (VxlanV6TunnelDestinationIPModeUnicast, error)
	// FromPbText unmarshals VxlanV6TunnelDestinationIPModeUnicast from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals VxlanV6TunnelDestinationIPModeUnicast from YAML text
	FromYaml(value string) error
	// FromJson unmarshals VxlanV6TunnelDestinationIPModeUnicast from JSON text
	FromJson(value string) error
}

func (obj *vxlanV6TunnelDestinationIPModeUnicast) Marshal() marshalVxlanV6TunnelDestinationIPModeUnicast {
	if obj.marshaller == nil {
		obj.marshaller = &marshalvxlanV6TunnelDestinationIPModeUnicast{obj: obj}
	}
	return obj.marshaller
}

func (obj *vxlanV6TunnelDestinationIPModeUnicast) Unmarshal() unMarshalVxlanV6TunnelDestinationIPModeUnicast {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalvxlanV6TunnelDestinationIPModeUnicast{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalvxlanV6TunnelDestinationIPModeUnicast) ToProto() (*otg.VxlanV6TunnelDestinationIPModeUnicast, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalvxlanV6TunnelDestinationIPModeUnicast) FromProto(msg *otg.VxlanV6TunnelDestinationIPModeUnicast) (VxlanV6TunnelDestinationIPModeUnicast, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalvxlanV6TunnelDestinationIPModeUnicast) ToPbText() (string, error) {
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

func (m *unMarshalvxlanV6TunnelDestinationIPModeUnicast) FromPbText(value string) error {
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

func (m *marshalvxlanV6TunnelDestinationIPModeUnicast) ToYaml() (string, error) {
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

func (m *unMarshalvxlanV6TunnelDestinationIPModeUnicast) FromYaml(value string) error {
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

func (m *marshalvxlanV6TunnelDestinationIPModeUnicast) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalvxlanV6TunnelDestinationIPModeUnicast) ToJson() (string, error) {
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

func (m *unMarshalvxlanV6TunnelDestinationIPModeUnicast) FromJson(value string) error {
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

func (obj *vxlanV6TunnelDestinationIPModeUnicast) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *vxlanV6TunnelDestinationIPModeUnicast) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *vxlanV6TunnelDestinationIPModeUnicast) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *vxlanV6TunnelDestinationIPModeUnicast) Clone() (VxlanV6TunnelDestinationIPModeUnicast, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewVxlanV6TunnelDestinationIPModeUnicast()
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

func (obj *vxlanV6TunnelDestinationIPModeUnicast) setNil() {
	obj.vtepsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// VxlanV6TunnelDestinationIPModeUnicast is description is TBD
type VxlanV6TunnelDestinationIPModeUnicast interface {
	Validation
	// msg marshals VxlanV6TunnelDestinationIPModeUnicast to protobuf object *otg.VxlanV6TunnelDestinationIPModeUnicast
	// and doesn't set defaults
	msg() *otg.VxlanV6TunnelDestinationIPModeUnicast
	// setMsg unmarshals VxlanV6TunnelDestinationIPModeUnicast from protobuf object *otg.VxlanV6TunnelDestinationIPModeUnicast
	// and doesn't set defaults
	setMsg(*otg.VxlanV6TunnelDestinationIPModeUnicast) VxlanV6TunnelDestinationIPModeUnicast
	// provides marshal interface
	Marshal() marshalVxlanV6TunnelDestinationIPModeUnicast
	// provides unmarshal interface
	Unmarshal() unMarshalVxlanV6TunnelDestinationIPModeUnicast
	// validate validates VxlanV6TunnelDestinationIPModeUnicast
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (VxlanV6TunnelDestinationIPModeUnicast, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Vteps returns VxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIterIter, set in VxlanV6TunnelDestinationIPModeUnicast
	Vteps() VxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter
	setNil()
}

// List of VTEPs for member VNI(VXLAN Network Identifier)
// Vteps returns a []VxlanV6TunnelDestinationIPModeUnicastVtep
func (obj *vxlanV6TunnelDestinationIPModeUnicast) Vteps() VxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter {
	if len(obj.obj.Vteps) == 0 {
		obj.obj.Vteps = []*otg.VxlanV6TunnelDestinationIPModeUnicastVtep{}
	}
	if obj.vtepsHolder == nil {
		obj.vtepsHolder = newVxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter(&obj.obj.Vteps).setMsg(obj)
	}
	return obj.vtepsHolder
}

type vxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter struct {
	obj                                            *vxlanV6TunnelDestinationIPModeUnicast
	vxlanV6TunnelDestinationIPModeUnicastVtepSlice []VxlanV6TunnelDestinationIPModeUnicastVtep
	fieldPtr                                       *[]*otg.VxlanV6TunnelDestinationIPModeUnicastVtep
}

func newVxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter(ptr *[]*otg.VxlanV6TunnelDestinationIPModeUnicastVtep) VxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter {
	return &vxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter{fieldPtr: ptr}
}

type VxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter interface {
	setMsg(*vxlanV6TunnelDestinationIPModeUnicast) VxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter
	Items() []VxlanV6TunnelDestinationIPModeUnicastVtep
	Add() VxlanV6TunnelDestinationIPModeUnicastVtep
	Append(items ...VxlanV6TunnelDestinationIPModeUnicastVtep) VxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter
	Set(index int, newObj VxlanV6TunnelDestinationIPModeUnicastVtep) VxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter
	Clear() VxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter
	clearHolderSlice() VxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter
	appendHolderSlice(item VxlanV6TunnelDestinationIPModeUnicastVtep) VxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter
}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter) setMsg(msg *vxlanV6TunnelDestinationIPModeUnicast) VxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&vxlanV6TunnelDestinationIPModeUnicastVtep{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter) Items() []VxlanV6TunnelDestinationIPModeUnicastVtep {
	return obj.vxlanV6TunnelDestinationIPModeUnicastVtepSlice
}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter) Add() VxlanV6TunnelDestinationIPModeUnicastVtep {
	newObj := &otg.VxlanV6TunnelDestinationIPModeUnicastVtep{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &vxlanV6TunnelDestinationIPModeUnicastVtep{obj: newObj}
	newLibObj.setDefault()
	obj.vxlanV6TunnelDestinationIPModeUnicastVtepSlice = append(obj.vxlanV6TunnelDestinationIPModeUnicastVtepSlice, newLibObj)
	return newLibObj
}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter) Append(items ...VxlanV6TunnelDestinationIPModeUnicastVtep) VxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.vxlanV6TunnelDestinationIPModeUnicastVtepSlice = append(obj.vxlanV6TunnelDestinationIPModeUnicastVtepSlice, item)
	}
	return obj
}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter) Set(index int, newObj VxlanV6TunnelDestinationIPModeUnicastVtep) VxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.vxlanV6TunnelDestinationIPModeUnicastVtepSlice[index] = newObj
	return obj
}
func (obj *vxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter) Clear() VxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.VxlanV6TunnelDestinationIPModeUnicastVtep{}
		obj.vxlanV6TunnelDestinationIPModeUnicastVtepSlice = []VxlanV6TunnelDestinationIPModeUnicastVtep{}
	}
	return obj
}
func (obj *vxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter) clearHolderSlice() VxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter {
	if len(obj.vxlanV6TunnelDestinationIPModeUnicastVtepSlice) > 0 {
		obj.vxlanV6TunnelDestinationIPModeUnicastVtepSlice = []VxlanV6TunnelDestinationIPModeUnicastVtep{}
	}
	return obj
}
func (obj *vxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter) appendHolderSlice(item VxlanV6TunnelDestinationIPModeUnicastVtep) VxlanV6TunnelDestinationIPModeUnicastVxlanV6TunnelDestinationIPModeUnicastVtepIter {
	obj.vxlanV6TunnelDestinationIPModeUnicastVtepSlice = append(obj.vxlanV6TunnelDestinationIPModeUnicastVtepSlice, item)
	return obj
}

func (obj *vxlanV6TunnelDestinationIPModeUnicast) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Vteps) != 0 {

		if set_default {
			obj.Vteps().clearHolderSlice()
			for _, item := range obj.obj.Vteps {
				obj.Vteps().appendHolderSlice(&vxlanV6TunnelDestinationIPModeUnicastVtep{obj: item})
			}
		}
		for _, item := range obj.Vteps().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *vxlanV6TunnelDestinationIPModeUnicast) setDefault() {

}

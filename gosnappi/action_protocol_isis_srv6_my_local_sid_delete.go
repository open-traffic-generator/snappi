package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolIsisSrv6MyLocalSidDelete *****
type actionProtocolIsisSrv6MyLocalSidDelete struct {
	validation
	obj           *otg.ActionProtocolIsisSrv6MyLocalSidDelete
	marshaller    marshalActionProtocolIsisSrv6MyLocalSidDelete
	unMarshaller  unMarshalActionProtocolIsisSrv6MyLocalSidDelete
	sidRefsHolder ActionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter
}

func NewActionProtocolIsisSrv6MyLocalSidDelete() ActionProtocolIsisSrv6MyLocalSidDelete {
	obj := actionProtocolIsisSrv6MyLocalSidDelete{obj: &otg.ActionProtocolIsisSrv6MyLocalSidDelete{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolIsisSrv6MyLocalSidDelete) msg() *otg.ActionProtocolIsisSrv6MyLocalSidDelete {
	return obj.obj
}

func (obj *actionProtocolIsisSrv6MyLocalSidDelete) setMsg(msg *otg.ActionProtocolIsisSrv6MyLocalSidDelete) ActionProtocolIsisSrv6MyLocalSidDelete {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolIsisSrv6MyLocalSidDelete struct {
	obj *actionProtocolIsisSrv6MyLocalSidDelete
}

type marshalActionProtocolIsisSrv6MyLocalSidDelete interface {
	// ToProto marshals ActionProtocolIsisSrv6MyLocalSidDelete to protobuf object *otg.ActionProtocolIsisSrv6MyLocalSidDelete
	ToProto() (*otg.ActionProtocolIsisSrv6MyLocalSidDelete, error)
	// ToPbText marshals ActionProtocolIsisSrv6MyLocalSidDelete to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolIsisSrv6MyLocalSidDelete to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolIsisSrv6MyLocalSidDelete to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolIsisSrv6MyLocalSidDelete struct {
	obj *actionProtocolIsisSrv6MyLocalSidDelete
}

type unMarshalActionProtocolIsisSrv6MyLocalSidDelete interface {
	// FromProto unmarshals ActionProtocolIsisSrv6MyLocalSidDelete from protobuf object *otg.ActionProtocolIsisSrv6MyLocalSidDelete
	FromProto(msg *otg.ActionProtocolIsisSrv6MyLocalSidDelete) (ActionProtocolIsisSrv6MyLocalSidDelete, error)
	// FromPbText unmarshals ActionProtocolIsisSrv6MyLocalSidDelete from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolIsisSrv6MyLocalSidDelete from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolIsisSrv6MyLocalSidDelete from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolIsisSrv6MyLocalSidDelete) Marshal() marshalActionProtocolIsisSrv6MyLocalSidDelete {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolIsisSrv6MyLocalSidDelete{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolIsisSrv6MyLocalSidDelete) Unmarshal() unMarshalActionProtocolIsisSrv6MyLocalSidDelete {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolIsisSrv6MyLocalSidDelete{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolIsisSrv6MyLocalSidDelete) ToProto() (*otg.ActionProtocolIsisSrv6MyLocalSidDelete, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolIsisSrv6MyLocalSidDelete) FromProto(msg *otg.ActionProtocolIsisSrv6MyLocalSidDelete) (ActionProtocolIsisSrv6MyLocalSidDelete, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolIsisSrv6MyLocalSidDelete) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolIsisSrv6MyLocalSidDelete) FromPbText(value string) error {
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

func (m *marshalactionProtocolIsisSrv6MyLocalSidDelete) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolIsisSrv6MyLocalSidDelete) FromYaml(value string) error {
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

func (m *marshalactionProtocolIsisSrv6MyLocalSidDelete) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolIsisSrv6MyLocalSidDelete) FromJson(value string) error {
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

func (obj *actionProtocolIsisSrv6MyLocalSidDelete) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolIsisSrv6MyLocalSidDelete) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolIsisSrv6MyLocalSidDelete) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolIsisSrv6MyLocalSidDelete) Clone() (ActionProtocolIsisSrv6MyLocalSidDelete, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolIsisSrv6MyLocalSidDelete()
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

func (obj *actionProtocolIsisSrv6MyLocalSidDelete) setNil() {
	obj.sidRefsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionProtocolIsisSrv6MyLocalSidDelete is remove SRv6 My Local SID entries from the routers specified by router_names. After deletion, packets whose DA matches these SID prefixes are no longer forwarded (RFC 8986 Section 4 Un-instantiation). Session impact: implementation-specific. In IxNetwork IS-IS withdraws the affected SIDs from its Locator TLV.
type ActionProtocolIsisSrv6MyLocalSidDelete interface {
	Validation
	// msg marshals ActionProtocolIsisSrv6MyLocalSidDelete to protobuf object *otg.ActionProtocolIsisSrv6MyLocalSidDelete
	// and doesn't set defaults
	msg() *otg.ActionProtocolIsisSrv6MyLocalSidDelete
	// setMsg unmarshals ActionProtocolIsisSrv6MyLocalSidDelete from protobuf object *otg.ActionProtocolIsisSrv6MyLocalSidDelete
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolIsisSrv6MyLocalSidDelete) ActionProtocolIsisSrv6MyLocalSidDelete
	// provides marshal interface
	Marshal() marshalActionProtocolIsisSrv6MyLocalSidDelete
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolIsisSrv6MyLocalSidDelete
	// validate validates ActionProtocolIsisSrv6MyLocalSidDelete
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolIsisSrv6MyLocalSidDelete, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SidRefs returns ActionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIterIter, set in ActionProtocolIsisSrv6MyLocalSidDelete
	SidRefs() ActionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter
	setNil()
}

// List of My Local SID entries to delete, identified by prefix and prefix length.
// SidRefs returns a []ActionProtocolIsisSrv6MyLocalSidSidRef
func (obj *actionProtocolIsisSrv6MyLocalSidDelete) SidRefs() ActionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter {
	if len(obj.obj.SidRefs) == 0 {
		obj.obj.SidRefs = []*otg.ActionProtocolIsisSrv6MyLocalSidSidRef{}
	}
	if obj.sidRefsHolder == nil {
		obj.sidRefsHolder = newActionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter(&obj.obj.SidRefs).setMsg(obj)
	}
	return obj.sidRefsHolder
}

type actionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter struct {
	obj                                         *actionProtocolIsisSrv6MyLocalSidDelete
	actionProtocolIsisSrv6MyLocalSidSidRefSlice []ActionProtocolIsisSrv6MyLocalSidSidRef
	fieldPtr                                    *[]*otg.ActionProtocolIsisSrv6MyLocalSidSidRef
}

func newActionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter(ptr *[]*otg.ActionProtocolIsisSrv6MyLocalSidSidRef) ActionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter {
	return &actionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter{fieldPtr: ptr}
}

type ActionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter interface {
	setMsg(*actionProtocolIsisSrv6MyLocalSidDelete) ActionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter
	Items() []ActionProtocolIsisSrv6MyLocalSidSidRef
	Add() ActionProtocolIsisSrv6MyLocalSidSidRef
	Append(items ...ActionProtocolIsisSrv6MyLocalSidSidRef) ActionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter
	Set(index int, newObj ActionProtocolIsisSrv6MyLocalSidSidRef) ActionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter
	Clear() ActionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter
	clearHolderSlice() ActionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter
	appendHolderSlice(item ActionProtocolIsisSrv6MyLocalSidSidRef) ActionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter
}

func (obj *actionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter) setMsg(msg *actionProtocolIsisSrv6MyLocalSidDelete) ActionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&actionProtocolIsisSrv6MyLocalSidSidRef{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *actionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter) Items() []ActionProtocolIsisSrv6MyLocalSidSidRef {
	return obj.actionProtocolIsisSrv6MyLocalSidSidRefSlice
}

func (obj *actionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter) Add() ActionProtocolIsisSrv6MyLocalSidSidRef {
	newObj := &otg.ActionProtocolIsisSrv6MyLocalSidSidRef{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &actionProtocolIsisSrv6MyLocalSidSidRef{obj: newObj}
	newLibObj.setDefault()
	obj.actionProtocolIsisSrv6MyLocalSidSidRefSlice = append(obj.actionProtocolIsisSrv6MyLocalSidSidRefSlice, newLibObj)
	return newLibObj
}

func (obj *actionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter) Append(items ...ActionProtocolIsisSrv6MyLocalSidSidRef) ActionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.actionProtocolIsisSrv6MyLocalSidSidRefSlice = append(obj.actionProtocolIsisSrv6MyLocalSidSidRefSlice, item)
	}
	return obj
}

func (obj *actionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter) Set(index int, newObj ActionProtocolIsisSrv6MyLocalSidSidRef) ActionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.actionProtocolIsisSrv6MyLocalSidSidRefSlice[index] = newObj
	return obj
}
func (obj *actionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter) Clear() ActionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ActionProtocolIsisSrv6MyLocalSidSidRef{}
		obj.actionProtocolIsisSrv6MyLocalSidSidRefSlice = []ActionProtocolIsisSrv6MyLocalSidSidRef{}
	}
	return obj
}
func (obj *actionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter) clearHolderSlice() ActionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter {
	if len(obj.actionProtocolIsisSrv6MyLocalSidSidRefSlice) > 0 {
		obj.actionProtocolIsisSrv6MyLocalSidSidRefSlice = []ActionProtocolIsisSrv6MyLocalSidSidRef{}
	}
	return obj
}
func (obj *actionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter) appendHolderSlice(item ActionProtocolIsisSrv6MyLocalSidSidRef) ActionProtocolIsisSrv6MyLocalSidDeleteActionProtocolIsisSrv6MyLocalSidSidRefIter {
	obj.actionProtocolIsisSrv6MyLocalSidSidRefSlice = append(obj.actionProtocolIsisSrv6MyLocalSidSidRefSlice, item)
	return obj
}

func (obj *actionProtocolIsisSrv6MyLocalSidDelete) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.SidRefs) != 0 {

		if set_default {
			obj.SidRefs().clearHolderSlice()
			for _, item := range obj.obj.SidRefs {
				obj.SidRefs().appendHolderSlice(&actionProtocolIsisSrv6MyLocalSidSidRef{obj: item})
			}
		}
		for _, item := range obj.SidRefs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *actionProtocolIsisSrv6MyLocalSidDelete) setDefault() {

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolIsisSrv6MyLocalSidModify *****
type actionProtocolIsisSrv6MyLocalSidModify struct {
	validation
	obj           *otg.ActionProtocolIsisSrv6MyLocalSidModify
	marshaller    marshalActionProtocolIsisSrv6MyLocalSidModify
	unMarshaller  unMarshalActionProtocolIsisSrv6MyLocalSidModify
	entriesHolder ActionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter
}

func NewActionProtocolIsisSrv6MyLocalSidModify() ActionProtocolIsisSrv6MyLocalSidModify {
	obj := actionProtocolIsisSrv6MyLocalSidModify{obj: &otg.ActionProtocolIsisSrv6MyLocalSidModify{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolIsisSrv6MyLocalSidModify) msg() *otg.ActionProtocolIsisSrv6MyLocalSidModify {
	return obj.obj
}

func (obj *actionProtocolIsisSrv6MyLocalSidModify) setMsg(msg *otg.ActionProtocolIsisSrv6MyLocalSidModify) ActionProtocolIsisSrv6MyLocalSidModify {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolIsisSrv6MyLocalSidModify struct {
	obj *actionProtocolIsisSrv6MyLocalSidModify
}

type marshalActionProtocolIsisSrv6MyLocalSidModify interface {
	// ToProto marshals ActionProtocolIsisSrv6MyLocalSidModify to protobuf object *otg.ActionProtocolIsisSrv6MyLocalSidModify
	ToProto() (*otg.ActionProtocolIsisSrv6MyLocalSidModify, error)
	// ToPbText marshals ActionProtocolIsisSrv6MyLocalSidModify to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolIsisSrv6MyLocalSidModify to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolIsisSrv6MyLocalSidModify to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolIsisSrv6MyLocalSidModify struct {
	obj *actionProtocolIsisSrv6MyLocalSidModify
}

type unMarshalActionProtocolIsisSrv6MyLocalSidModify interface {
	// FromProto unmarshals ActionProtocolIsisSrv6MyLocalSidModify from protobuf object *otg.ActionProtocolIsisSrv6MyLocalSidModify
	FromProto(msg *otg.ActionProtocolIsisSrv6MyLocalSidModify) (ActionProtocolIsisSrv6MyLocalSidModify, error)
	// FromPbText unmarshals ActionProtocolIsisSrv6MyLocalSidModify from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolIsisSrv6MyLocalSidModify from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolIsisSrv6MyLocalSidModify from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolIsisSrv6MyLocalSidModify) Marshal() marshalActionProtocolIsisSrv6MyLocalSidModify {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolIsisSrv6MyLocalSidModify{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolIsisSrv6MyLocalSidModify) Unmarshal() unMarshalActionProtocolIsisSrv6MyLocalSidModify {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolIsisSrv6MyLocalSidModify{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolIsisSrv6MyLocalSidModify) ToProto() (*otg.ActionProtocolIsisSrv6MyLocalSidModify, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolIsisSrv6MyLocalSidModify) FromProto(msg *otg.ActionProtocolIsisSrv6MyLocalSidModify) (ActionProtocolIsisSrv6MyLocalSidModify, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolIsisSrv6MyLocalSidModify) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolIsisSrv6MyLocalSidModify) FromPbText(value string) error {
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

func (m *marshalactionProtocolIsisSrv6MyLocalSidModify) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolIsisSrv6MyLocalSidModify) FromYaml(value string) error {
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

func (m *marshalactionProtocolIsisSrv6MyLocalSidModify) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolIsisSrv6MyLocalSidModify) FromJson(value string) error {
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

func (obj *actionProtocolIsisSrv6MyLocalSidModify) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolIsisSrv6MyLocalSidModify) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolIsisSrv6MyLocalSidModify) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolIsisSrv6MyLocalSidModify) Clone() (ActionProtocolIsisSrv6MyLocalSidModify, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolIsisSrv6MyLocalSidModify()
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

func (obj *actionProtocolIsisSrv6MyLocalSidModify) setNil() {
	obj.entriesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionProtocolIsisSrv6MyLocalSidModify is modify existing SRv6 My Local SID entries on the routers specified by router_names. Each entry is matched by sid_prefix + prefix_length; the remaining fields replace the existing values without a delete-and-re-add cycle. Session impact: implementation-specific. In IxNetwork IS-IS re-sends its LSP with the updated SID set.
type ActionProtocolIsisSrv6MyLocalSidModify interface {
	Validation
	// msg marshals ActionProtocolIsisSrv6MyLocalSidModify to protobuf object *otg.ActionProtocolIsisSrv6MyLocalSidModify
	// and doesn't set defaults
	msg() *otg.ActionProtocolIsisSrv6MyLocalSidModify
	// setMsg unmarshals ActionProtocolIsisSrv6MyLocalSidModify from protobuf object *otg.ActionProtocolIsisSrv6MyLocalSidModify
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolIsisSrv6MyLocalSidModify) ActionProtocolIsisSrv6MyLocalSidModify
	// provides marshal interface
	Marshal() marshalActionProtocolIsisSrv6MyLocalSidModify
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolIsisSrv6MyLocalSidModify
	// validate validates ActionProtocolIsisSrv6MyLocalSidModify
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolIsisSrv6MyLocalSidModify, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Entries returns ActionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIterIter, set in ActionProtocolIsisSrv6MyLocalSidModify
	Entries() ActionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter
	setNil()
}

// List of My Local SID entries to modify. Each entry is matched by sid_prefix + prefix_length; other fields are updated.
// Entries returns a []ActionProtocolIsisSrv6MyLocalSidEntry
func (obj *actionProtocolIsisSrv6MyLocalSidModify) Entries() ActionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter {
	if len(obj.obj.Entries) == 0 {
		obj.obj.Entries = []*otg.ActionProtocolIsisSrv6MyLocalSidEntry{}
	}
	if obj.entriesHolder == nil {
		obj.entriesHolder = newActionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter(&obj.obj.Entries).setMsg(obj)
	}
	return obj.entriesHolder
}

type actionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter struct {
	obj                                        *actionProtocolIsisSrv6MyLocalSidModify
	actionProtocolIsisSrv6MyLocalSidEntrySlice []ActionProtocolIsisSrv6MyLocalSidEntry
	fieldPtr                                   *[]*otg.ActionProtocolIsisSrv6MyLocalSidEntry
}

func newActionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter(ptr *[]*otg.ActionProtocolIsisSrv6MyLocalSidEntry) ActionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter {
	return &actionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter{fieldPtr: ptr}
}

type ActionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter interface {
	setMsg(*actionProtocolIsisSrv6MyLocalSidModify) ActionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter
	Items() []ActionProtocolIsisSrv6MyLocalSidEntry
	Add() ActionProtocolIsisSrv6MyLocalSidEntry
	Append(items ...ActionProtocolIsisSrv6MyLocalSidEntry) ActionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter
	Set(index int, newObj ActionProtocolIsisSrv6MyLocalSidEntry) ActionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter
	Clear() ActionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter
	clearHolderSlice() ActionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter
	appendHolderSlice(item ActionProtocolIsisSrv6MyLocalSidEntry) ActionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter
}

func (obj *actionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter) setMsg(msg *actionProtocolIsisSrv6MyLocalSidModify) ActionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&actionProtocolIsisSrv6MyLocalSidEntry{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *actionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter) Items() []ActionProtocolIsisSrv6MyLocalSidEntry {
	return obj.actionProtocolIsisSrv6MyLocalSidEntrySlice
}

func (obj *actionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter) Add() ActionProtocolIsisSrv6MyLocalSidEntry {
	newObj := &otg.ActionProtocolIsisSrv6MyLocalSidEntry{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &actionProtocolIsisSrv6MyLocalSidEntry{obj: newObj}
	newLibObj.setDefault()
	obj.actionProtocolIsisSrv6MyLocalSidEntrySlice = append(obj.actionProtocolIsisSrv6MyLocalSidEntrySlice, newLibObj)
	return newLibObj
}

func (obj *actionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter) Append(items ...ActionProtocolIsisSrv6MyLocalSidEntry) ActionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.actionProtocolIsisSrv6MyLocalSidEntrySlice = append(obj.actionProtocolIsisSrv6MyLocalSidEntrySlice, item)
	}
	return obj
}

func (obj *actionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter) Set(index int, newObj ActionProtocolIsisSrv6MyLocalSidEntry) ActionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.actionProtocolIsisSrv6MyLocalSidEntrySlice[index] = newObj
	return obj
}
func (obj *actionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter) Clear() ActionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ActionProtocolIsisSrv6MyLocalSidEntry{}
		obj.actionProtocolIsisSrv6MyLocalSidEntrySlice = []ActionProtocolIsisSrv6MyLocalSidEntry{}
	}
	return obj
}
func (obj *actionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter) clearHolderSlice() ActionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter {
	if len(obj.actionProtocolIsisSrv6MyLocalSidEntrySlice) > 0 {
		obj.actionProtocolIsisSrv6MyLocalSidEntrySlice = []ActionProtocolIsisSrv6MyLocalSidEntry{}
	}
	return obj
}
func (obj *actionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter) appendHolderSlice(item ActionProtocolIsisSrv6MyLocalSidEntry) ActionProtocolIsisSrv6MyLocalSidModifyActionProtocolIsisSrv6MyLocalSidEntryIter {
	obj.actionProtocolIsisSrv6MyLocalSidEntrySlice = append(obj.actionProtocolIsisSrv6MyLocalSidEntrySlice, item)
	return obj
}

func (obj *actionProtocolIsisSrv6MyLocalSidModify) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Entries) != 0 {

		if set_default {
			obj.Entries().clearHolderSlice()
			for _, item := range obj.obj.Entries {
				obj.Entries().appendHolderSlice(&actionProtocolIsisSrv6MyLocalSidEntry{obj: item})
			}
		}
		for _, item := range obj.Entries().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *actionProtocolIsisSrv6MyLocalSidModify) setDefault() {

}

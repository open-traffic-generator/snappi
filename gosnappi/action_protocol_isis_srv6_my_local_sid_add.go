package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolIsisSrv6MyLocalSidAdd *****
type actionProtocolIsisSrv6MyLocalSidAdd struct {
	validation
	obj           *otg.ActionProtocolIsisSrv6MyLocalSidAdd
	marshaller    marshalActionProtocolIsisSrv6MyLocalSidAdd
	unMarshaller  unMarshalActionProtocolIsisSrv6MyLocalSidAdd
	entriesHolder ActionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter
}

func NewActionProtocolIsisSrv6MyLocalSidAdd() ActionProtocolIsisSrv6MyLocalSidAdd {
	obj := actionProtocolIsisSrv6MyLocalSidAdd{obj: &otg.ActionProtocolIsisSrv6MyLocalSidAdd{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolIsisSrv6MyLocalSidAdd) msg() *otg.ActionProtocolIsisSrv6MyLocalSidAdd {
	return obj.obj
}

func (obj *actionProtocolIsisSrv6MyLocalSidAdd) setMsg(msg *otg.ActionProtocolIsisSrv6MyLocalSidAdd) ActionProtocolIsisSrv6MyLocalSidAdd {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolIsisSrv6MyLocalSidAdd struct {
	obj *actionProtocolIsisSrv6MyLocalSidAdd
}

type marshalActionProtocolIsisSrv6MyLocalSidAdd interface {
	// ToProto marshals ActionProtocolIsisSrv6MyLocalSidAdd to protobuf object *otg.ActionProtocolIsisSrv6MyLocalSidAdd
	ToProto() (*otg.ActionProtocolIsisSrv6MyLocalSidAdd, error)
	// ToPbText marshals ActionProtocolIsisSrv6MyLocalSidAdd to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolIsisSrv6MyLocalSidAdd to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolIsisSrv6MyLocalSidAdd to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolIsisSrv6MyLocalSidAdd struct {
	obj *actionProtocolIsisSrv6MyLocalSidAdd
}

type unMarshalActionProtocolIsisSrv6MyLocalSidAdd interface {
	// FromProto unmarshals ActionProtocolIsisSrv6MyLocalSidAdd from protobuf object *otg.ActionProtocolIsisSrv6MyLocalSidAdd
	FromProto(msg *otg.ActionProtocolIsisSrv6MyLocalSidAdd) (ActionProtocolIsisSrv6MyLocalSidAdd, error)
	// FromPbText unmarshals ActionProtocolIsisSrv6MyLocalSidAdd from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolIsisSrv6MyLocalSidAdd from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolIsisSrv6MyLocalSidAdd from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolIsisSrv6MyLocalSidAdd) Marshal() marshalActionProtocolIsisSrv6MyLocalSidAdd {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolIsisSrv6MyLocalSidAdd{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolIsisSrv6MyLocalSidAdd) Unmarshal() unMarshalActionProtocolIsisSrv6MyLocalSidAdd {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolIsisSrv6MyLocalSidAdd{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolIsisSrv6MyLocalSidAdd) ToProto() (*otg.ActionProtocolIsisSrv6MyLocalSidAdd, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolIsisSrv6MyLocalSidAdd) FromProto(msg *otg.ActionProtocolIsisSrv6MyLocalSidAdd) (ActionProtocolIsisSrv6MyLocalSidAdd, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolIsisSrv6MyLocalSidAdd) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolIsisSrv6MyLocalSidAdd) FromPbText(value string) error {
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

func (m *marshalactionProtocolIsisSrv6MyLocalSidAdd) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolIsisSrv6MyLocalSidAdd) FromYaml(value string) error {
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

func (m *marshalactionProtocolIsisSrv6MyLocalSidAdd) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolIsisSrv6MyLocalSidAdd) FromJson(value string) error {
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

func (obj *actionProtocolIsisSrv6MyLocalSidAdd) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolIsisSrv6MyLocalSidAdd) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolIsisSrv6MyLocalSidAdd) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolIsisSrv6MyLocalSidAdd) Clone() (ActionProtocolIsisSrv6MyLocalSidAdd, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolIsisSrv6MyLocalSidAdd()
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

func (obj *actionProtocolIsisSrv6MyLocalSidAdd) setNil() {
	obj.entriesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionProtocolIsisSrv6MyLocalSidAdd is program new SRv6 My Local SID entries on the routers specified by router_names. After this operation the router matches packets whose DA falls within each SID prefix and applies the configured behavior. Session impact: implementation-specific. In IxNetwork the IS-IS router re-advertises its Locator TLV after the add.
type ActionProtocolIsisSrv6MyLocalSidAdd interface {
	Validation
	// msg marshals ActionProtocolIsisSrv6MyLocalSidAdd to protobuf object *otg.ActionProtocolIsisSrv6MyLocalSidAdd
	// and doesn't set defaults
	msg() *otg.ActionProtocolIsisSrv6MyLocalSidAdd
	// setMsg unmarshals ActionProtocolIsisSrv6MyLocalSidAdd from protobuf object *otg.ActionProtocolIsisSrv6MyLocalSidAdd
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolIsisSrv6MyLocalSidAdd) ActionProtocolIsisSrv6MyLocalSidAdd
	// provides marshal interface
	Marshal() marshalActionProtocolIsisSrv6MyLocalSidAdd
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolIsisSrv6MyLocalSidAdd
	// validate validates ActionProtocolIsisSrv6MyLocalSidAdd
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolIsisSrv6MyLocalSidAdd, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Entries returns ActionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIterIter, set in ActionProtocolIsisSrv6MyLocalSidAdd
	Entries() ActionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter
	setNil()
}

// List of My Local SID entries to add.
// Entries returns a []ActionProtocolIsisSrv6MyLocalSidEntry
func (obj *actionProtocolIsisSrv6MyLocalSidAdd) Entries() ActionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter {
	if len(obj.obj.Entries) == 0 {
		obj.obj.Entries = []*otg.ActionProtocolIsisSrv6MyLocalSidEntry{}
	}
	if obj.entriesHolder == nil {
		obj.entriesHolder = newActionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter(&obj.obj.Entries).setMsg(obj)
	}
	return obj.entriesHolder
}

type actionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter struct {
	obj                                        *actionProtocolIsisSrv6MyLocalSidAdd
	actionProtocolIsisSrv6MyLocalSidEntrySlice []ActionProtocolIsisSrv6MyLocalSidEntry
	fieldPtr                                   *[]*otg.ActionProtocolIsisSrv6MyLocalSidEntry
}

func newActionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter(ptr *[]*otg.ActionProtocolIsisSrv6MyLocalSidEntry) ActionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter {
	return &actionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter{fieldPtr: ptr}
}

type ActionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter interface {
	setMsg(*actionProtocolIsisSrv6MyLocalSidAdd) ActionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter
	Items() []ActionProtocolIsisSrv6MyLocalSidEntry
	Add() ActionProtocolIsisSrv6MyLocalSidEntry
	Append(items ...ActionProtocolIsisSrv6MyLocalSidEntry) ActionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter
	Set(index int, newObj ActionProtocolIsisSrv6MyLocalSidEntry) ActionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter
	Clear() ActionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter
	clearHolderSlice() ActionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter
	appendHolderSlice(item ActionProtocolIsisSrv6MyLocalSidEntry) ActionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter
}

func (obj *actionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter) setMsg(msg *actionProtocolIsisSrv6MyLocalSidAdd) ActionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&actionProtocolIsisSrv6MyLocalSidEntry{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *actionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter) Items() []ActionProtocolIsisSrv6MyLocalSidEntry {
	return obj.actionProtocolIsisSrv6MyLocalSidEntrySlice
}

func (obj *actionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter) Add() ActionProtocolIsisSrv6MyLocalSidEntry {
	newObj := &otg.ActionProtocolIsisSrv6MyLocalSidEntry{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &actionProtocolIsisSrv6MyLocalSidEntry{obj: newObj}
	newLibObj.setDefault()
	obj.actionProtocolIsisSrv6MyLocalSidEntrySlice = append(obj.actionProtocolIsisSrv6MyLocalSidEntrySlice, newLibObj)
	return newLibObj
}

func (obj *actionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter) Append(items ...ActionProtocolIsisSrv6MyLocalSidEntry) ActionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.actionProtocolIsisSrv6MyLocalSidEntrySlice = append(obj.actionProtocolIsisSrv6MyLocalSidEntrySlice, item)
	}
	return obj
}

func (obj *actionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter) Set(index int, newObj ActionProtocolIsisSrv6MyLocalSidEntry) ActionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.actionProtocolIsisSrv6MyLocalSidEntrySlice[index] = newObj
	return obj
}
func (obj *actionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter) Clear() ActionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ActionProtocolIsisSrv6MyLocalSidEntry{}
		obj.actionProtocolIsisSrv6MyLocalSidEntrySlice = []ActionProtocolIsisSrv6MyLocalSidEntry{}
	}
	return obj
}
func (obj *actionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter) clearHolderSlice() ActionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter {
	if len(obj.actionProtocolIsisSrv6MyLocalSidEntrySlice) > 0 {
		obj.actionProtocolIsisSrv6MyLocalSidEntrySlice = []ActionProtocolIsisSrv6MyLocalSidEntry{}
	}
	return obj
}
func (obj *actionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter) appendHolderSlice(item ActionProtocolIsisSrv6MyLocalSidEntry) ActionProtocolIsisSrv6MyLocalSidAddActionProtocolIsisSrv6MyLocalSidEntryIter {
	obj.actionProtocolIsisSrv6MyLocalSidEntrySlice = append(obj.actionProtocolIsisSrv6MyLocalSidEntrySlice, item)
	return obj
}

func (obj *actionProtocolIsisSrv6MyLocalSidAdd) validateObj(vObj *validation, set_default bool) {
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

func (obj *actionProtocolIsisSrv6MyLocalSidAdd) setDefault() {

}

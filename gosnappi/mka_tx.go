package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MkaTx *****
type mkaTx struct {
	validation
	obj                  *otg.MkaTx
	marshaller           marshalMkaTx
	unMarshaller         unMarshalMkaTx
	secureChannelsHolder MkaTxMkaTxScIter
}

func NewMkaTx() MkaTx {
	obj := mkaTx{obj: &otg.MkaTx{}}
	obj.setDefault()
	return &obj
}

func (obj *mkaTx) msg() *otg.MkaTx {
	return obj.obj
}

func (obj *mkaTx) setMsg(msg *otg.MkaTx) MkaTx {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmkaTx struct {
	obj *mkaTx
}

type marshalMkaTx interface {
	// ToProto marshals MkaTx to protobuf object *otg.MkaTx
	ToProto() (*otg.MkaTx, error)
	// ToPbText marshals MkaTx to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MkaTx to YAML text
	ToYaml() (string, error)
	// ToJson marshals MkaTx to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals MkaTx to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalmkaTx struct {
	obj *mkaTx
}

type unMarshalMkaTx interface {
	// FromProto unmarshals MkaTx from protobuf object *otg.MkaTx
	FromProto(msg *otg.MkaTx) (MkaTx, error)
	// FromPbText unmarshals MkaTx from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MkaTx from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MkaTx from JSON text
	FromJson(value string) error
}

func (obj *mkaTx) Marshal() marshalMkaTx {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmkaTx{obj: obj}
	}
	return obj.marshaller
}

func (obj *mkaTx) Unmarshal() unMarshalMkaTx {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmkaTx{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmkaTx) ToProto() (*otg.MkaTx, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmkaTx) FromProto(msg *otg.MkaTx) (MkaTx, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmkaTx) ToPbText() (string, error) {
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

func (m *unMarshalmkaTx) FromPbText(value string) error {
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

func (m *marshalmkaTx) ToYaml() (string, error) {
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

func (m *unMarshalmkaTx) FromYaml(value string) error {
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

func (m *marshalmkaTx) ToJsonRaw() (string, error) {
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

func (m *marshalmkaTx) ToJson() (string, error) {
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

func (m *unMarshalmkaTx) FromJson(value string) error {
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

func (obj *mkaTx) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mkaTx) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mkaTx) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mkaTx) Clone() (MkaTx, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMkaTx()
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

func (obj *mkaTx) setNil() {
	obj.secureChannelsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MkaTx is a container of Tx properties.
type MkaTx interface {
	Validation
	// msg marshals MkaTx to protobuf object *otg.MkaTx
	// and doesn't set defaults
	msg() *otg.MkaTx
	// setMsg unmarshals MkaTx from protobuf object *otg.MkaTx
	// and doesn't set defaults
	setMsg(*otg.MkaTx) MkaTx
	// provides marshal interface
	Marshal() marshalMkaTx
	// provides unmarshal interface
	Unmarshal() unMarshalMkaTx
	// validate validates MkaTx
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MkaTx, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SecureChannels returns MkaTxMkaTxScIterIter, set in MkaTx
	SecureChannels() MkaTxMkaTxScIter
	setNil()
}

// Tx secure channels.
// SecureChannels returns a []MkaTxSc
func (obj *mkaTx) SecureChannels() MkaTxMkaTxScIter {
	if len(obj.obj.SecureChannels) == 0 {
		obj.obj.SecureChannels = []*otg.MkaTxSc{}
	}
	if obj.secureChannelsHolder == nil {
		obj.secureChannelsHolder = newMkaTxMkaTxScIter(&obj.obj.SecureChannels).setMsg(obj)
	}
	return obj.secureChannelsHolder
}

type mkaTxMkaTxScIter struct {
	obj          *mkaTx
	mkaTxScSlice []MkaTxSc
	fieldPtr     *[]*otg.MkaTxSc
}

func newMkaTxMkaTxScIter(ptr *[]*otg.MkaTxSc) MkaTxMkaTxScIter {
	return &mkaTxMkaTxScIter{fieldPtr: ptr}
}

type MkaTxMkaTxScIter interface {
	setMsg(*mkaTx) MkaTxMkaTxScIter
	Items() []MkaTxSc
	Add() MkaTxSc
	Append(items ...MkaTxSc) MkaTxMkaTxScIter
	Set(index int, newObj MkaTxSc) MkaTxMkaTxScIter
	Clear() MkaTxMkaTxScIter
	clearHolderSlice() MkaTxMkaTxScIter
	appendHolderSlice(item MkaTxSc) MkaTxMkaTxScIter
}

func (obj *mkaTxMkaTxScIter) setMsg(msg *mkaTx) MkaTxMkaTxScIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&mkaTxSc{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *mkaTxMkaTxScIter) Items() []MkaTxSc {
	return obj.mkaTxScSlice
}

func (obj *mkaTxMkaTxScIter) Add() MkaTxSc {
	newObj := &otg.MkaTxSc{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &mkaTxSc{obj: newObj}
	newLibObj.setDefault()
	obj.mkaTxScSlice = append(obj.mkaTxScSlice, newLibObj)
	return newLibObj
}

func (obj *mkaTxMkaTxScIter) Append(items ...MkaTxSc) MkaTxMkaTxScIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.mkaTxScSlice = append(obj.mkaTxScSlice, item)
	}
	return obj
}

func (obj *mkaTxMkaTxScIter) Set(index int, newObj MkaTxSc) MkaTxMkaTxScIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.mkaTxScSlice[index] = newObj
	return obj
}
func (obj *mkaTxMkaTxScIter) Clear() MkaTxMkaTxScIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.MkaTxSc{}
		obj.mkaTxScSlice = []MkaTxSc{}
	}
	return obj
}
func (obj *mkaTxMkaTxScIter) clearHolderSlice() MkaTxMkaTxScIter {
	if len(obj.mkaTxScSlice) > 0 {
		obj.mkaTxScSlice = []MkaTxSc{}
	}
	return obj
}
func (obj *mkaTxMkaTxScIter) appendHolderSlice(item MkaTxSc) MkaTxMkaTxScIter {
	obj.mkaTxScSlice = append(obj.mkaTxScSlice, item)
	return obj
}

func (obj *mkaTx) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.SecureChannels) != 0 {

		if set_default {
			obj.SecureChannels().clearHolderSlice()
			for _, item := range obj.obj.SecureChannels {
				obj.SecureChannels().appendHolderSlice(&mkaTxSc{obj: item})
			}
		}
		for _, item := range obj.SecureChannels().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *mkaTx) setDefault() {

}

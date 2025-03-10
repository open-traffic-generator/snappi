package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityStaticKeyRx *****
type secureEntityStaticKeyRx struct {
	validation
	obj                  *otg.SecureEntityStaticKeyRx
	marshaller           marshalSecureEntityStaticKeyRx
	unMarshaller         unMarshalSecureEntityStaticKeyRx
	secureChannelsHolder SecureEntityStaticKeyRxSecureEntityStaticKeyRxScIter
}

func NewSecureEntityStaticKeyRx() SecureEntityStaticKeyRx {
	obj := secureEntityStaticKeyRx{obj: &otg.SecureEntityStaticKeyRx{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityStaticKeyRx) msg() *otg.SecureEntityStaticKeyRx {
	return obj.obj
}

func (obj *secureEntityStaticKeyRx) setMsg(msg *otg.SecureEntityStaticKeyRx) SecureEntityStaticKeyRx {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityStaticKeyRx struct {
	obj *secureEntityStaticKeyRx
}

type marshalSecureEntityStaticKeyRx interface {
	// ToProto marshals SecureEntityStaticKeyRx to protobuf object *otg.SecureEntityStaticKeyRx
	ToProto() (*otg.SecureEntityStaticKeyRx, error)
	// ToPbText marshals SecureEntityStaticKeyRx to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityStaticKeyRx to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityStaticKeyRx to JSON text
	ToJson() (string, error)
}

type unMarshalsecureEntityStaticKeyRx struct {
	obj *secureEntityStaticKeyRx
}

type unMarshalSecureEntityStaticKeyRx interface {
	// FromProto unmarshals SecureEntityStaticKeyRx from protobuf object *otg.SecureEntityStaticKeyRx
	FromProto(msg *otg.SecureEntityStaticKeyRx) (SecureEntityStaticKeyRx, error)
	// FromPbText unmarshals SecureEntityStaticKeyRx from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityStaticKeyRx from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityStaticKeyRx from JSON text
	FromJson(value string) error
}

func (obj *secureEntityStaticKeyRx) Marshal() marshalSecureEntityStaticKeyRx {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityStaticKeyRx{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityStaticKeyRx) Unmarshal() unMarshalSecureEntityStaticKeyRx {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityStaticKeyRx{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityStaticKeyRx) ToProto() (*otg.SecureEntityStaticKeyRx, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityStaticKeyRx) FromProto(msg *otg.SecureEntityStaticKeyRx) (SecureEntityStaticKeyRx, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityStaticKeyRx) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityStaticKeyRx) FromPbText(value string) error {
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

func (m *marshalsecureEntityStaticKeyRx) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityStaticKeyRx) FromYaml(value string) error {
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

func (m *marshalsecureEntityStaticKeyRx) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityStaticKeyRx) FromJson(value string) error {
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

func (obj *secureEntityStaticKeyRx) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityStaticKeyRx) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityStaticKeyRx) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityStaticKeyRx) Clone() (SecureEntityStaticKeyRx, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityStaticKeyRx()
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

func (obj *secureEntityStaticKeyRx) setNil() {
	obj.secureChannelsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// SecureEntityStaticKeyRx is a container of static key Rx properties.
type SecureEntityStaticKeyRx interface {
	Validation
	// msg marshals SecureEntityStaticKeyRx to protobuf object *otg.SecureEntityStaticKeyRx
	// and doesn't set defaults
	msg() *otg.SecureEntityStaticKeyRx
	// setMsg unmarshals SecureEntityStaticKeyRx from protobuf object *otg.SecureEntityStaticKeyRx
	// and doesn't set defaults
	setMsg(*otg.SecureEntityStaticKeyRx) SecureEntityStaticKeyRx
	// provides marshal interface
	Marshal() marshalSecureEntityStaticKeyRx
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityStaticKeyRx
	// validate validates SecureEntityStaticKeyRx
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityStaticKeyRx, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SecureChannels returns SecureEntityStaticKeyRxSecureEntityStaticKeyRxScIterIter, set in SecureEntityStaticKeyRx
	SecureChannels() SecureEntityStaticKeyRxSecureEntityStaticKeyRxScIter
	setNil()
}

// Rx secure channels.
// SecureChannels returns a []SecureEntityStaticKeyRxSc
func (obj *secureEntityStaticKeyRx) SecureChannels() SecureEntityStaticKeyRxSecureEntityStaticKeyRxScIter {
	if len(obj.obj.SecureChannels) == 0 {
		obj.obj.SecureChannels = []*otg.SecureEntityStaticKeyRxSc{}
	}
	if obj.secureChannelsHolder == nil {
		obj.secureChannelsHolder = newSecureEntityStaticKeyRxSecureEntityStaticKeyRxScIter(&obj.obj.SecureChannels).setMsg(obj)
	}
	return obj.secureChannelsHolder
}

type secureEntityStaticKeyRxSecureEntityStaticKeyRxScIter struct {
	obj                            *secureEntityStaticKeyRx
	secureEntityStaticKeyRxScSlice []SecureEntityStaticKeyRxSc
	fieldPtr                       *[]*otg.SecureEntityStaticKeyRxSc
}

func newSecureEntityStaticKeyRxSecureEntityStaticKeyRxScIter(ptr *[]*otg.SecureEntityStaticKeyRxSc) SecureEntityStaticKeyRxSecureEntityStaticKeyRxScIter {
	return &secureEntityStaticKeyRxSecureEntityStaticKeyRxScIter{fieldPtr: ptr}
}

type SecureEntityStaticKeyRxSecureEntityStaticKeyRxScIter interface {
	setMsg(*secureEntityStaticKeyRx) SecureEntityStaticKeyRxSecureEntityStaticKeyRxScIter
	Items() []SecureEntityStaticKeyRxSc
	Add() SecureEntityStaticKeyRxSc
	Append(items ...SecureEntityStaticKeyRxSc) SecureEntityStaticKeyRxSecureEntityStaticKeyRxScIter
	Set(index int, newObj SecureEntityStaticKeyRxSc) SecureEntityStaticKeyRxSecureEntityStaticKeyRxScIter
	Clear() SecureEntityStaticKeyRxSecureEntityStaticKeyRxScIter
	clearHolderSlice() SecureEntityStaticKeyRxSecureEntityStaticKeyRxScIter
	appendHolderSlice(item SecureEntityStaticKeyRxSc) SecureEntityStaticKeyRxSecureEntityStaticKeyRxScIter
}

func (obj *secureEntityStaticKeyRxSecureEntityStaticKeyRxScIter) setMsg(msg *secureEntityStaticKeyRx) SecureEntityStaticKeyRxSecureEntityStaticKeyRxScIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&secureEntityStaticKeyRxSc{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *secureEntityStaticKeyRxSecureEntityStaticKeyRxScIter) Items() []SecureEntityStaticKeyRxSc {
	return obj.secureEntityStaticKeyRxScSlice
}

func (obj *secureEntityStaticKeyRxSecureEntityStaticKeyRxScIter) Add() SecureEntityStaticKeyRxSc {
	newObj := &otg.SecureEntityStaticKeyRxSc{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &secureEntityStaticKeyRxSc{obj: newObj}
	newLibObj.setDefault()
	obj.secureEntityStaticKeyRxScSlice = append(obj.secureEntityStaticKeyRxScSlice, newLibObj)
	return newLibObj
}

func (obj *secureEntityStaticKeyRxSecureEntityStaticKeyRxScIter) Append(items ...SecureEntityStaticKeyRxSc) SecureEntityStaticKeyRxSecureEntityStaticKeyRxScIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.secureEntityStaticKeyRxScSlice = append(obj.secureEntityStaticKeyRxScSlice, item)
	}
	return obj
}

func (obj *secureEntityStaticKeyRxSecureEntityStaticKeyRxScIter) Set(index int, newObj SecureEntityStaticKeyRxSc) SecureEntityStaticKeyRxSecureEntityStaticKeyRxScIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.secureEntityStaticKeyRxScSlice[index] = newObj
	return obj
}
func (obj *secureEntityStaticKeyRxSecureEntityStaticKeyRxScIter) Clear() SecureEntityStaticKeyRxSecureEntityStaticKeyRxScIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.SecureEntityStaticKeyRxSc{}
		obj.secureEntityStaticKeyRxScSlice = []SecureEntityStaticKeyRxSc{}
	}
	return obj
}
func (obj *secureEntityStaticKeyRxSecureEntityStaticKeyRxScIter) clearHolderSlice() SecureEntityStaticKeyRxSecureEntityStaticKeyRxScIter {
	if len(obj.secureEntityStaticKeyRxScSlice) > 0 {
		obj.secureEntityStaticKeyRxScSlice = []SecureEntityStaticKeyRxSc{}
	}
	return obj
}
func (obj *secureEntityStaticKeyRxSecureEntityStaticKeyRxScIter) appendHolderSlice(item SecureEntityStaticKeyRxSc) SecureEntityStaticKeyRxSecureEntityStaticKeyRxScIter {
	obj.secureEntityStaticKeyRxScSlice = append(obj.secureEntityStaticKeyRxScSlice, item)
	return obj
}

func (obj *secureEntityStaticKeyRx) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.SecureChannels) != 0 {

		if set_default {
			obj.SecureChannels().clearHolderSlice()
			for _, item := range obj.obj.SecureChannels {
				obj.SecureChannels().appendHolderSlice(&secureEntityStaticKeyRxSc{obj: item})
			}
		}
		for _, item := range obj.SecureChannels().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *secureEntityStaticKeyRx) setDefault() {

}

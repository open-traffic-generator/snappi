package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityStaticKeyTx *****
type secureEntityStaticKeyTx struct {
	validation
	obj                  *otg.SecureEntityStaticKeyTx
	marshaller           marshalSecureEntityStaticKeyTx
	unMarshaller         unMarshalSecureEntityStaticKeyTx
	secureChannelsHolder SecureEntityStaticKeyTxSecureEntityStaticKeyTxScIter
	rekeyModeHolder      SecureEntityStaticKeyRekeyMode
}

func NewSecureEntityStaticKeyTx() SecureEntityStaticKeyTx {
	obj := secureEntityStaticKeyTx{obj: &otg.SecureEntityStaticKeyTx{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityStaticKeyTx) msg() *otg.SecureEntityStaticKeyTx {
	return obj.obj
}

func (obj *secureEntityStaticKeyTx) setMsg(msg *otg.SecureEntityStaticKeyTx) SecureEntityStaticKeyTx {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityStaticKeyTx struct {
	obj *secureEntityStaticKeyTx
}

type marshalSecureEntityStaticKeyTx interface {
	// ToProto marshals SecureEntityStaticKeyTx to protobuf object *otg.SecureEntityStaticKeyTx
	ToProto() (*otg.SecureEntityStaticKeyTx, error)
	// ToPbText marshals SecureEntityStaticKeyTx to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityStaticKeyTx to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityStaticKeyTx to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals SecureEntityStaticKeyTx to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalsecureEntityStaticKeyTx struct {
	obj *secureEntityStaticKeyTx
}

type unMarshalSecureEntityStaticKeyTx interface {
	// FromProto unmarshals SecureEntityStaticKeyTx from protobuf object *otg.SecureEntityStaticKeyTx
	FromProto(msg *otg.SecureEntityStaticKeyTx) (SecureEntityStaticKeyTx, error)
	// FromPbText unmarshals SecureEntityStaticKeyTx from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityStaticKeyTx from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityStaticKeyTx from JSON text
	FromJson(value string) error
}

func (obj *secureEntityStaticKeyTx) Marshal() marshalSecureEntityStaticKeyTx {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityStaticKeyTx{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityStaticKeyTx) Unmarshal() unMarshalSecureEntityStaticKeyTx {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityStaticKeyTx{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityStaticKeyTx) ToProto() (*otg.SecureEntityStaticKeyTx, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityStaticKeyTx) FromProto(msg *otg.SecureEntityStaticKeyTx) (SecureEntityStaticKeyTx, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityStaticKeyTx) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityStaticKeyTx) FromPbText(value string) error {
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

func (m *marshalsecureEntityStaticKeyTx) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityStaticKeyTx) FromYaml(value string) error {
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

func (m *marshalsecureEntityStaticKeyTx) ToJsonRaw() (string, error) {
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

func (m *marshalsecureEntityStaticKeyTx) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityStaticKeyTx) FromJson(value string) error {
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

func (obj *secureEntityStaticKeyTx) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityStaticKeyTx) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityStaticKeyTx) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityStaticKeyTx) Clone() (SecureEntityStaticKeyTx, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityStaticKeyTx()
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

func (obj *secureEntityStaticKeyTx) setNil() {
	obj.secureChannelsHolder = nil
	obj.rekeyModeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// SecureEntityStaticKeyTx is a container of static key Tx properties.
type SecureEntityStaticKeyTx interface {
	Validation
	// msg marshals SecureEntityStaticKeyTx to protobuf object *otg.SecureEntityStaticKeyTx
	// and doesn't set defaults
	msg() *otg.SecureEntityStaticKeyTx
	// setMsg unmarshals SecureEntityStaticKeyTx from protobuf object *otg.SecureEntityStaticKeyTx
	// and doesn't set defaults
	setMsg(*otg.SecureEntityStaticKeyTx) SecureEntityStaticKeyTx
	// provides marshal interface
	Marshal() marshalSecureEntityStaticKeyTx
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityStaticKeyTx
	// validate validates SecureEntityStaticKeyTx
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityStaticKeyTx, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SecureChannels returns SecureEntityStaticKeyTxSecureEntityStaticKeyTxScIterIter, set in SecureEntityStaticKeyTx
	SecureChannels() SecureEntityStaticKeyTxSecureEntityStaticKeyTxScIter
	// RekeyMode returns SecureEntityStaticKeyRekeyMode, set in SecureEntityStaticKeyTx.
	// SecureEntityStaticKeyRekeyMode is rekey mode.
	RekeyMode() SecureEntityStaticKeyRekeyMode
	// SetRekeyMode assigns SecureEntityStaticKeyRekeyMode provided by user to SecureEntityStaticKeyTx.
	// SecureEntityStaticKeyRekeyMode is rekey mode.
	SetRekeyMode(value SecureEntityStaticKeyRekeyMode) SecureEntityStaticKeyTx
	// HasRekeyMode checks if RekeyMode has been set in SecureEntityStaticKeyTx
	HasRekeyMode() bool
	setNil()
}

// Tx secure channels.
// SecureChannels returns a []SecureEntityStaticKeyTxSc
func (obj *secureEntityStaticKeyTx) SecureChannels() SecureEntityStaticKeyTxSecureEntityStaticKeyTxScIter {
	if len(obj.obj.SecureChannels) == 0 {
		obj.obj.SecureChannels = []*otg.SecureEntityStaticKeyTxSc{}
	}
	if obj.secureChannelsHolder == nil {
		obj.secureChannelsHolder = newSecureEntityStaticKeyTxSecureEntityStaticKeyTxScIter(&obj.obj.SecureChannels).setMsg(obj)
	}
	return obj.secureChannelsHolder
}

type secureEntityStaticKeyTxSecureEntityStaticKeyTxScIter struct {
	obj                            *secureEntityStaticKeyTx
	secureEntityStaticKeyTxScSlice []SecureEntityStaticKeyTxSc
	fieldPtr                       *[]*otg.SecureEntityStaticKeyTxSc
}

func newSecureEntityStaticKeyTxSecureEntityStaticKeyTxScIter(ptr *[]*otg.SecureEntityStaticKeyTxSc) SecureEntityStaticKeyTxSecureEntityStaticKeyTxScIter {
	return &secureEntityStaticKeyTxSecureEntityStaticKeyTxScIter{fieldPtr: ptr}
}

type SecureEntityStaticKeyTxSecureEntityStaticKeyTxScIter interface {
	setMsg(*secureEntityStaticKeyTx) SecureEntityStaticKeyTxSecureEntityStaticKeyTxScIter
	Items() []SecureEntityStaticKeyTxSc
	Add() SecureEntityStaticKeyTxSc
	Append(items ...SecureEntityStaticKeyTxSc) SecureEntityStaticKeyTxSecureEntityStaticKeyTxScIter
	Set(index int, newObj SecureEntityStaticKeyTxSc) SecureEntityStaticKeyTxSecureEntityStaticKeyTxScIter
	Clear() SecureEntityStaticKeyTxSecureEntityStaticKeyTxScIter
	clearHolderSlice() SecureEntityStaticKeyTxSecureEntityStaticKeyTxScIter
	appendHolderSlice(item SecureEntityStaticKeyTxSc) SecureEntityStaticKeyTxSecureEntityStaticKeyTxScIter
}

func (obj *secureEntityStaticKeyTxSecureEntityStaticKeyTxScIter) setMsg(msg *secureEntityStaticKeyTx) SecureEntityStaticKeyTxSecureEntityStaticKeyTxScIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&secureEntityStaticKeyTxSc{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *secureEntityStaticKeyTxSecureEntityStaticKeyTxScIter) Items() []SecureEntityStaticKeyTxSc {
	return obj.secureEntityStaticKeyTxScSlice
}

func (obj *secureEntityStaticKeyTxSecureEntityStaticKeyTxScIter) Add() SecureEntityStaticKeyTxSc {
	newObj := &otg.SecureEntityStaticKeyTxSc{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &secureEntityStaticKeyTxSc{obj: newObj}
	newLibObj.setDefault()
	obj.secureEntityStaticKeyTxScSlice = append(obj.secureEntityStaticKeyTxScSlice, newLibObj)
	return newLibObj
}

func (obj *secureEntityStaticKeyTxSecureEntityStaticKeyTxScIter) Append(items ...SecureEntityStaticKeyTxSc) SecureEntityStaticKeyTxSecureEntityStaticKeyTxScIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.secureEntityStaticKeyTxScSlice = append(obj.secureEntityStaticKeyTxScSlice, item)
	}
	return obj
}

func (obj *secureEntityStaticKeyTxSecureEntityStaticKeyTxScIter) Set(index int, newObj SecureEntityStaticKeyTxSc) SecureEntityStaticKeyTxSecureEntityStaticKeyTxScIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.secureEntityStaticKeyTxScSlice[index] = newObj
	return obj
}
func (obj *secureEntityStaticKeyTxSecureEntityStaticKeyTxScIter) Clear() SecureEntityStaticKeyTxSecureEntityStaticKeyTxScIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.SecureEntityStaticKeyTxSc{}
		obj.secureEntityStaticKeyTxScSlice = []SecureEntityStaticKeyTxSc{}
	}
	return obj
}
func (obj *secureEntityStaticKeyTxSecureEntityStaticKeyTxScIter) clearHolderSlice() SecureEntityStaticKeyTxSecureEntityStaticKeyTxScIter {
	if len(obj.secureEntityStaticKeyTxScSlice) > 0 {
		obj.secureEntityStaticKeyTxScSlice = []SecureEntityStaticKeyTxSc{}
	}
	return obj
}
func (obj *secureEntityStaticKeyTxSecureEntityStaticKeyTxScIter) appendHolderSlice(item SecureEntityStaticKeyTxSc) SecureEntityStaticKeyTxSecureEntityStaticKeyTxScIter {
	obj.secureEntityStaticKeyTxScSlice = append(obj.secureEntityStaticKeyTxScSlice, item)
	return obj
}

// description is TBD
// RekeyMode returns a SecureEntityStaticKeyRekeyMode
func (obj *secureEntityStaticKeyTx) RekeyMode() SecureEntityStaticKeyRekeyMode {
	if obj.obj.RekeyMode == nil {
		obj.obj.RekeyMode = NewSecureEntityStaticKeyRekeyMode().msg()
	}
	if obj.rekeyModeHolder == nil {
		obj.rekeyModeHolder = &secureEntityStaticKeyRekeyMode{obj: obj.obj.RekeyMode}
	}
	return obj.rekeyModeHolder
}

// description is TBD
// RekeyMode returns a SecureEntityStaticKeyRekeyMode
func (obj *secureEntityStaticKeyTx) HasRekeyMode() bool {
	return obj.obj.RekeyMode != nil
}

// description is TBD
// SetRekeyMode sets the SecureEntityStaticKeyRekeyMode value in the SecureEntityStaticKeyTx object
func (obj *secureEntityStaticKeyTx) SetRekeyMode(value SecureEntityStaticKeyRekeyMode) SecureEntityStaticKeyTx {

	obj.rekeyModeHolder = nil
	obj.obj.RekeyMode = value.msg()

	return obj
}

func (obj *secureEntityStaticKeyTx) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.SecureChannels) != 0 {

		if set_default {
			obj.SecureChannels().clearHolderSlice()
			for _, item := range obj.obj.SecureChannels {
				obj.SecureChannels().appendHolderSlice(&secureEntityStaticKeyTxSc{obj: item})
			}
		}
		for _, item := range obj.SecureChannels().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.RekeyMode != nil {

		obj.RekeyMode().validateObj(vObj, set_default)
	}

}

func (obj *secureEntityStaticKeyTx) setDefault() {

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityStaticKeyTxSc *****
type secureEntityStaticKeyTxSc struct {
	validation
	obj          *otg.SecureEntityStaticKeyTxSc
	marshaller   marshalSecureEntityStaticKeyTxSc
	unMarshaller unMarshalSecureEntityStaticKeyTxSc
	saksHolder   SecureEntityStaticKeyTxScSecureEntityStaticKeySakIter
}

func NewSecureEntityStaticKeyTxSc() SecureEntityStaticKeyTxSc {
	obj := secureEntityStaticKeyTxSc{obj: &otg.SecureEntityStaticKeyTxSc{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityStaticKeyTxSc) msg() *otg.SecureEntityStaticKeyTxSc {
	return obj.obj
}

func (obj *secureEntityStaticKeyTxSc) setMsg(msg *otg.SecureEntityStaticKeyTxSc) SecureEntityStaticKeyTxSc {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityStaticKeyTxSc struct {
	obj *secureEntityStaticKeyTxSc
}

type marshalSecureEntityStaticKeyTxSc interface {
	// ToProto marshals SecureEntityStaticKeyTxSc to protobuf object *otg.SecureEntityStaticKeyTxSc
	ToProto() (*otg.SecureEntityStaticKeyTxSc, error)
	// ToPbText marshals SecureEntityStaticKeyTxSc to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityStaticKeyTxSc to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityStaticKeyTxSc to JSON text
	ToJson() (string, error)
}

type unMarshalsecureEntityStaticKeyTxSc struct {
	obj *secureEntityStaticKeyTxSc
}

type unMarshalSecureEntityStaticKeyTxSc interface {
	// FromProto unmarshals SecureEntityStaticKeyTxSc from protobuf object *otg.SecureEntityStaticKeyTxSc
	FromProto(msg *otg.SecureEntityStaticKeyTxSc) (SecureEntityStaticKeyTxSc, error)
	// FromPbText unmarshals SecureEntityStaticKeyTxSc from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityStaticKeyTxSc from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityStaticKeyTxSc from JSON text
	FromJson(value string) error
}

func (obj *secureEntityStaticKeyTxSc) Marshal() marshalSecureEntityStaticKeyTxSc {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityStaticKeyTxSc{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityStaticKeyTxSc) Unmarshal() unMarshalSecureEntityStaticKeyTxSc {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityStaticKeyTxSc{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityStaticKeyTxSc) ToProto() (*otg.SecureEntityStaticKeyTxSc, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityStaticKeyTxSc) FromProto(msg *otg.SecureEntityStaticKeyTxSc) (SecureEntityStaticKeyTxSc, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityStaticKeyTxSc) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityStaticKeyTxSc) FromPbText(value string) error {
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

func (m *marshalsecureEntityStaticKeyTxSc) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityStaticKeyTxSc) FromYaml(value string) error {
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

func (m *marshalsecureEntityStaticKeyTxSc) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityStaticKeyTxSc) FromJson(value string) error {
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

func (obj *secureEntityStaticKeyTxSc) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityStaticKeyTxSc) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityStaticKeyTxSc) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityStaticKeyTxSc) Clone() (SecureEntityStaticKeyTxSc, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityStaticKeyTxSc()
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

func (obj *secureEntityStaticKeyTxSc) setNil() {
	obj.saksHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// SecureEntityStaticKeyTxSc is tx SC setting for static key.
type SecureEntityStaticKeyTxSc interface {
	Validation
	// msg marshals SecureEntityStaticKeyTxSc to protobuf object *otg.SecureEntityStaticKeyTxSc
	// and doesn't set defaults
	msg() *otg.SecureEntityStaticKeyTxSc
	// setMsg unmarshals SecureEntityStaticKeyTxSc from protobuf object *otg.SecureEntityStaticKeyTxSc
	// and doesn't set defaults
	setMsg(*otg.SecureEntityStaticKeyTxSc) SecureEntityStaticKeyTxSc
	// provides marshal interface
	Marshal() marshalSecureEntityStaticKeyTxSc
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityStaticKeyTxSc
	// validate validates SecureEntityStaticKeyTxSc
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityStaticKeyTxSc, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SystemId returns string, set in SecureEntityStaticKeyTxSc.
	SystemId() string
	// SetSystemId assigns string provided by user to SecureEntityStaticKeyTxSc
	SetSystemId(value string) SecureEntityStaticKeyTxSc
	// HasSystemId checks if SystemId has been set in SecureEntityStaticKeyTxSc
	HasSystemId() bool
	// PortId returns uint32, set in SecureEntityStaticKeyTxSc.
	PortId() uint32
	// SetPortId assigns uint32 provided by user to SecureEntityStaticKeyTxSc
	SetPortId(value uint32) SecureEntityStaticKeyTxSc
	// HasPortId checks if PortId has been set in SecureEntityStaticKeyTxSc
	HasPortId() bool
	// Saks returns SecureEntityStaticKeyTxScSecureEntityStaticKeySakIterIter, set in SecureEntityStaticKeyTxSc
	Saks() SecureEntityStaticKeyTxScSecureEntityStaticKeySakIter
	setNil()
}

// System ID.
// SystemId returns a string
func (obj *secureEntityStaticKeyTxSc) SystemId() string {

	return *obj.obj.SystemId

}

// System ID.
// SystemId returns a string
func (obj *secureEntityStaticKeyTxSc) HasSystemId() bool {
	return obj.obj.SystemId != nil
}

// System ID.
// SetSystemId sets the string value in the SecureEntityStaticKeyTxSc object
func (obj *secureEntityStaticKeyTxSc) SetSystemId(value string) SecureEntityStaticKeyTxSc {

	obj.obj.SystemId = &value
	return obj
}

// Port ID.
// PortId returns a uint32
func (obj *secureEntityStaticKeyTxSc) PortId() uint32 {

	return *obj.obj.PortId

}

// Port ID.
// PortId returns a uint32
func (obj *secureEntityStaticKeyTxSc) HasPortId() bool {
	return obj.obj.PortId != nil
}

// Port ID.
// SetPortId sets the uint32 value in the SecureEntityStaticKeyTxSc object
func (obj *secureEntityStaticKeyTxSc) SetPortId(value uint32) SecureEntityStaticKeyTxSc {

	obj.obj.PortId = &value
	return obj
}

// Tx SAK pool.
// Saks returns a []SecureEntityStaticKeySak
func (obj *secureEntityStaticKeyTxSc) Saks() SecureEntityStaticKeyTxScSecureEntityStaticKeySakIter {
	if len(obj.obj.Saks) == 0 {
		obj.obj.Saks = []*otg.SecureEntityStaticKeySak{}
	}
	if obj.saksHolder == nil {
		obj.saksHolder = newSecureEntityStaticKeyTxScSecureEntityStaticKeySakIter(&obj.obj.Saks).setMsg(obj)
	}
	return obj.saksHolder
}

type secureEntityStaticKeyTxScSecureEntityStaticKeySakIter struct {
	obj                           *secureEntityStaticKeyTxSc
	secureEntityStaticKeySakSlice []SecureEntityStaticKeySak
	fieldPtr                      *[]*otg.SecureEntityStaticKeySak
}

func newSecureEntityStaticKeyTxScSecureEntityStaticKeySakIter(ptr *[]*otg.SecureEntityStaticKeySak) SecureEntityStaticKeyTxScSecureEntityStaticKeySakIter {
	return &secureEntityStaticKeyTxScSecureEntityStaticKeySakIter{fieldPtr: ptr}
}

type SecureEntityStaticKeyTxScSecureEntityStaticKeySakIter interface {
	setMsg(*secureEntityStaticKeyTxSc) SecureEntityStaticKeyTxScSecureEntityStaticKeySakIter
	Items() []SecureEntityStaticKeySak
	Add() SecureEntityStaticKeySak
	Append(items ...SecureEntityStaticKeySak) SecureEntityStaticKeyTxScSecureEntityStaticKeySakIter
	Set(index int, newObj SecureEntityStaticKeySak) SecureEntityStaticKeyTxScSecureEntityStaticKeySakIter
	Clear() SecureEntityStaticKeyTxScSecureEntityStaticKeySakIter
	clearHolderSlice() SecureEntityStaticKeyTxScSecureEntityStaticKeySakIter
	appendHolderSlice(item SecureEntityStaticKeySak) SecureEntityStaticKeyTxScSecureEntityStaticKeySakIter
}

func (obj *secureEntityStaticKeyTxScSecureEntityStaticKeySakIter) setMsg(msg *secureEntityStaticKeyTxSc) SecureEntityStaticKeyTxScSecureEntityStaticKeySakIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&secureEntityStaticKeySak{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *secureEntityStaticKeyTxScSecureEntityStaticKeySakIter) Items() []SecureEntityStaticKeySak {
	return obj.secureEntityStaticKeySakSlice
}

func (obj *secureEntityStaticKeyTxScSecureEntityStaticKeySakIter) Add() SecureEntityStaticKeySak {
	newObj := &otg.SecureEntityStaticKeySak{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &secureEntityStaticKeySak{obj: newObj}
	newLibObj.setDefault()
	obj.secureEntityStaticKeySakSlice = append(obj.secureEntityStaticKeySakSlice, newLibObj)
	return newLibObj
}

func (obj *secureEntityStaticKeyTxScSecureEntityStaticKeySakIter) Append(items ...SecureEntityStaticKeySak) SecureEntityStaticKeyTxScSecureEntityStaticKeySakIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.secureEntityStaticKeySakSlice = append(obj.secureEntityStaticKeySakSlice, item)
	}
	return obj
}

func (obj *secureEntityStaticKeyTxScSecureEntityStaticKeySakIter) Set(index int, newObj SecureEntityStaticKeySak) SecureEntityStaticKeyTxScSecureEntityStaticKeySakIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.secureEntityStaticKeySakSlice[index] = newObj
	return obj
}
func (obj *secureEntityStaticKeyTxScSecureEntityStaticKeySakIter) Clear() SecureEntityStaticKeyTxScSecureEntityStaticKeySakIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.SecureEntityStaticKeySak{}
		obj.secureEntityStaticKeySakSlice = []SecureEntityStaticKeySak{}
	}
	return obj
}
func (obj *secureEntityStaticKeyTxScSecureEntityStaticKeySakIter) clearHolderSlice() SecureEntityStaticKeyTxScSecureEntityStaticKeySakIter {
	if len(obj.secureEntityStaticKeySakSlice) > 0 {
		obj.secureEntityStaticKeySakSlice = []SecureEntityStaticKeySak{}
	}
	return obj
}
func (obj *secureEntityStaticKeyTxScSecureEntityStaticKeySakIter) appendHolderSlice(item SecureEntityStaticKeySak) SecureEntityStaticKeyTxScSecureEntityStaticKeySakIter {
	obj.secureEntityStaticKeySakSlice = append(obj.secureEntityStaticKeySakSlice, item)
	return obj
}

func (obj *secureEntityStaticKeyTxSc) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SystemId != nil {

		err := obj.validateMac(obj.SystemId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on SecureEntityStaticKeyTxSc.SystemId"))
		}

	}

	if obj.obj.PortId != nil {

		if *obj.obj.PortId < 1 || *obj.obj.PortId > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= SecureEntityStaticKeyTxSc.PortId <= 65535 but Got %d", *obj.obj.PortId))
		}

	}

	if len(obj.obj.Saks) != 0 {

		if set_default {
			obj.Saks().clearHolderSlice()
			for _, item := range obj.obj.Saks {
				obj.Saks().appendHolderSlice(&secureEntityStaticKeySak{obj: item})
			}
		}
		for _, item := range obj.Saks().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *secureEntityStaticKeyTxSc) setDefault() {
	if obj.obj.PortId == nil {
		obj.SetPortId(1)
	}

}

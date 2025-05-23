package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityStaticKeyRxSc *****
type secureEntityStaticKeyRxSc struct {
	validation
	obj          *otg.SecureEntityStaticKeyRxSc
	marshaller   marshalSecureEntityStaticKeyRxSc
	unMarshaller unMarshalSecureEntityStaticKeyRxSc
	saksHolder   SecureEntityStaticKeyRxScSecureEntityStaticKeySakIter
}

func NewSecureEntityStaticKeyRxSc() SecureEntityStaticKeyRxSc {
	obj := secureEntityStaticKeyRxSc{obj: &otg.SecureEntityStaticKeyRxSc{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityStaticKeyRxSc) msg() *otg.SecureEntityStaticKeyRxSc {
	return obj.obj
}

func (obj *secureEntityStaticKeyRxSc) setMsg(msg *otg.SecureEntityStaticKeyRxSc) SecureEntityStaticKeyRxSc {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityStaticKeyRxSc struct {
	obj *secureEntityStaticKeyRxSc
}

type marshalSecureEntityStaticKeyRxSc interface {
	// ToProto marshals SecureEntityStaticKeyRxSc to protobuf object *otg.SecureEntityStaticKeyRxSc
	ToProto() (*otg.SecureEntityStaticKeyRxSc, error)
	// ToPbText marshals SecureEntityStaticKeyRxSc to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityStaticKeyRxSc to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityStaticKeyRxSc to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals SecureEntityStaticKeyRxSc to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalsecureEntityStaticKeyRxSc struct {
	obj *secureEntityStaticKeyRxSc
}

type unMarshalSecureEntityStaticKeyRxSc interface {
	// FromProto unmarshals SecureEntityStaticKeyRxSc from protobuf object *otg.SecureEntityStaticKeyRxSc
	FromProto(msg *otg.SecureEntityStaticKeyRxSc) (SecureEntityStaticKeyRxSc, error)
	// FromPbText unmarshals SecureEntityStaticKeyRxSc from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityStaticKeyRxSc from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityStaticKeyRxSc from JSON text
	FromJson(value string) error
}

func (obj *secureEntityStaticKeyRxSc) Marshal() marshalSecureEntityStaticKeyRxSc {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityStaticKeyRxSc{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityStaticKeyRxSc) Unmarshal() unMarshalSecureEntityStaticKeyRxSc {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityStaticKeyRxSc{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityStaticKeyRxSc) ToProto() (*otg.SecureEntityStaticKeyRxSc, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityStaticKeyRxSc) FromProto(msg *otg.SecureEntityStaticKeyRxSc) (SecureEntityStaticKeyRxSc, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityStaticKeyRxSc) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityStaticKeyRxSc) FromPbText(value string) error {
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

func (m *marshalsecureEntityStaticKeyRxSc) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityStaticKeyRxSc) FromYaml(value string) error {
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

func (m *marshalsecureEntityStaticKeyRxSc) ToJsonRaw() (string, error) {
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

func (m *marshalsecureEntityStaticKeyRxSc) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityStaticKeyRxSc) FromJson(value string) error {
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

func (obj *secureEntityStaticKeyRxSc) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityStaticKeyRxSc) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityStaticKeyRxSc) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityStaticKeyRxSc) Clone() (SecureEntityStaticKeyRxSc, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityStaticKeyRxSc()
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

func (obj *secureEntityStaticKeyRxSc) setNil() {
	obj.saksHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// SecureEntityStaticKeyRxSc is rx SC settings.
type SecureEntityStaticKeyRxSc interface {
	Validation
	// msg marshals SecureEntityStaticKeyRxSc to protobuf object *otg.SecureEntityStaticKeyRxSc
	// and doesn't set defaults
	msg() *otg.SecureEntityStaticKeyRxSc
	// setMsg unmarshals SecureEntityStaticKeyRxSc from protobuf object *otg.SecureEntityStaticKeyRxSc
	// and doesn't set defaults
	setMsg(*otg.SecureEntityStaticKeyRxSc) SecureEntityStaticKeyRxSc
	// provides marshal interface
	Marshal() marshalSecureEntityStaticKeyRxSc
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityStaticKeyRxSc
	// validate validates SecureEntityStaticKeyRxSc
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityStaticKeyRxSc, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// DutSciSystemId returns string, set in SecureEntityStaticKeyRxSc.
	DutSciSystemId() string
	// SetDutSciSystemId assigns string provided by user to SecureEntityStaticKeyRxSc
	SetDutSciSystemId(value string) SecureEntityStaticKeyRxSc
	// HasDutSciSystemId checks if DutSciSystemId has been set in SecureEntityStaticKeyRxSc
	HasDutSciSystemId() bool
	// DutSciPortId returns uint32, set in SecureEntityStaticKeyRxSc.
	DutSciPortId() uint32
	// SetDutSciPortId assigns uint32 provided by user to SecureEntityStaticKeyRxSc
	SetDutSciPortId(value uint32) SecureEntityStaticKeyRxSc
	// HasDutSciPortId checks if DutSciPortId has been set in SecureEntityStaticKeyRxSc
	HasDutSciPortId() bool
	// DutMsbXpn returns uint32, set in SecureEntityStaticKeyRxSc.
	DutMsbXpn() uint32
	// SetDutMsbXpn assigns uint32 provided by user to SecureEntityStaticKeyRxSc
	SetDutMsbXpn(value uint32) SecureEntityStaticKeyRxSc
	// HasDutMsbXpn checks if DutMsbXpn has been set in SecureEntityStaticKeyRxSc
	HasDutMsbXpn() bool
	// Saks returns SecureEntityStaticKeyRxScSecureEntityStaticKeySakIterIter, set in SecureEntityStaticKeyRxSc
	Saks() SecureEntityStaticKeyRxScSecureEntityStaticKeySakIter
	setNil()
}

// System ID in DUT SCI.
// DutSciSystemId returns a string
func (obj *secureEntityStaticKeyRxSc) DutSciSystemId() string {

	return *obj.obj.DutSciSystemId

}

// System ID in DUT SCI.
// DutSciSystemId returns a string
func (obj *secureEntityStaticKeyRxSc) HasDutSciSystemId() bool {
	return obj.obj.DutSciSystemId != nil
}

// System ID in DUT SCI.
// SetDutSciSystemId sets the string value in the SecureEntityStaticKeyRxSc object
func (obj *secureEntityStaticKeyRxSc) SetDutSciSystemId(value string) SecureEntityStaticKeyRxSc {

	obj.obj.DutSciSystemId = &value
	return obj
}

// Port ID in DUT SCI.
// DutSciPortId returns a uint32
func (obj *secureEntityStaticKeyRxSc) DutSciPortId() uint32 {

	return *obj.obj.DutSciPortId

}

// Port ID in DUT SCI.
// DutSciPortId returns a uint32
func (obj *secureEntityStaticKeyRxSc) HasDutSciPortId() bool {
	return obj.obj.DutSciPortId != nil
}

// Port ID in DUT SCI.
// SetDutSciPortId sets the uint32 value in the SecureEntityStaticKeyRxSc object
func (obj *secureEntityStaticKeyRxSc) SetDutSciPortId(value uint32) SecureEntityStaticKeyRxSc {

	obj.obj.DutSciPortId = &value
	return obj
}

// DUT MSB of XPN. The 32 most significant bits of the XPN that DUT will be using to construct the 64 bits XPN value when test starts.
// DutMsbXpn returns a uint32
func (obj *secureEntityStaticKeyRxSc) DutMsbXpn() uint32 {

	return *obj.obj.DutMsbXpn

}

// DUT MSB of XPN. The 32 most significant bits of the XPN that DUT will be using to construct the 64 bits XPN value when test starts.
// DutMsbXpn returns a uint32
func (obj *secureEntityStaticKeyRxSc) HasDutMsbXpn() bool {
	return obj.obj.DutMsbXpn != nil
}

// DUT MSB of XPN. The 32 most significant bits of the XPN that DUT will be using to construct the 64 bits XPN value when test starts.
// SetDutMsbXpn sets the uint32 value in the SecureEntityStaticKeyRxSc object
func (obj *secureEntityStaticKeyRxSc) SetDutMsbXpn(value uint32) SecureEntityStaticKeyRxSc {

	obj.obj.DutMsbXpn = &value
	return obj
}

// Rx SAK pool.
// Saks returns a []SecureEntityStaticKeySak
func (obj *secureEntityStaticKeyRxSc) Saks() SecureEntityStaticKeyRxScSecureEntityStaticKeySakIter {
	if len(obj.obj.Saks) == 0 {
		obj.obj.Saks = []*otg.SecureEntityStaticKeySak{}
	}
	if obj.saksHolder == nil {
		obj.saksHolder = newSecureEntityStaticKeyRxScSecureEntityStaticKeySakIter(&obj.obj.Saks).setMsg(obj)
	}
	return obj.saksHolder
}

type secureEntityStaticKeyRxScSecureEntityStaticKeySakIter struct {
	obj                           *secureEntityStaticKeyRxSc
	secureEntityStaticKeySakSlice []SecureEntityStaticKeySak
	fieldPtr                      *[]*otg.SecureEntityStaticKeySak
}

func newSecureEntityStaticKeyRxScSecureEntityStaticKeySakIter(ptr *[]*otg.SecureEntityStaticKeySak) SecureEntityStaticKeyRxScSecureEntityStaticKeySakIter {
	return &secureEntityStaticKeyRxScSecureEntityStaticKeySakIter{fieldPtr: ptr}
}

type SecureEntityStaticKeyRxScSecureEntityStaticKeySakIter interface {
	setMsg(*secureEntityStaticKeyRxSc) SecureEntityStaticKeyRxScSecureEntityStaticKeySakIter
	Items() []SecureEntityStaticKeySak
	Add() SecureEntityStaticKeySak
	Append(items ...SecureEntityStaticKeySak) SecureEntityStaticKeyRxScSecureEntityStaticKeySakIter
	Set(index int, newObj SecureEntityStaticKeySak) SecureEntityStaticKeyRxScSecureEntityStaticKeySakIter
	Clear() SecureEntityStaticKeyRxScSecureEntityStaticKeySakIter
	clearHolderSlice() SecureEntityStaticKeyRxScSecureEntityStaticKeySakIter
	appendHolderSlice(item SecureEntityStaticKeySak) SecureEntityStaticKeyRxScSecureEntityStaticKeySakIter
}

func (obj *secureEntityStaticKeyRxScSecureEntityStaticKeySakIter) setMsg(msg *secureEntityStaticKeyRxSc) SecureEntityStaticKeyRxScSecureEntityStaticKeySakIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&secureEntityStaticKeySak{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *secureEntityStaticKeyRxScSecureEntityStaticKeySakIter) Items() []SecureEntityStaticKeySak {
	return obj.secureEntityStaticKeySakSlice
}

func (obj *secureEntityStaticKeyRxScSecureEntityStaticKeySakIter) Add() SecureEntityStaticKeySak {
	newObj := &otg.SecureEntityStaticKeySak{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &secureEntityStaticKeySak{obj: newObj}
	newLibObj.setDefault()
	obj.secureEntityStaticKeySakSlice = append(obj.secureEntityStaticKeySakSlice, newLibObj)
	return newLibObj
}

func (obj *secureEntityStaticKeyRxScSecureEntityStaticKeySakIter) Append(items ...SecureEntityStaticKeySak) SecureEntityStaticKeyRxScSecureEntityStaticKeySakIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.secureEntityStaticKeySakSlice = append(obj.secureEntityStaticKeySakSlice, item)
	}
	return obj
}

func (obj *secureEntityStaticKeyRxScSecureEntityStaticKeySakIter) Set(index int, newObj SecureEntityStaticKeySak) SecureEntityStaticKeyRxScSecureEntityStaticKeySakIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.secureEntityStaticKeySakSlice[index] = newObj
	return obj
}
func (obj *secureEntityStaticKeyRxScSecureEntityStaticKeySakIter) Clear() SecureEntityStaticKeyRxScSecureEntityStaticKeySakIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.SecureEntityStaticKeySak{}
		obj.secureEntityStaticKeySakSlice = []SecureEntityStaticKeySak{}
	}
	return obj
}
func (obj *secureEntityStaticKeyRxScSecureEntityStaticKeySakIter) clearHolderSlice() SecureEntityStaticKeyRxScSecureEntityStaticKeySakIter {
	if len(obj.secureEntityStaticKeySakSlice) > 0 {
		obj.secureEntityStaticKeySakSlice = []SecureEntityStaticKeySak{}
	}
	return obj
}
func (obj *secureEntityStaticKeyRxScSecureEntityStaticKeySakIter) appendHolderSlice(item SecureEntityStaticKeySak) SecureEntityStaticKeyRxScSecureEntityStaticKeySakIter {
	obj.secureEntityStaticKeySakSlice = append(obj.secureEntityStaticKeySakSlice, item)
	return obj
}

func (obj *secureEntityStaticKeyRxSc) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.DutSciSystemId != nil {

		err := obj.validateMac(obj.DutSciSystemId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on SecureEntityStaticKeyRxSc.DutSciSystemId"))
		}

	}

	if obj.obj.DutSciPortId != nil {

		if *obj.obj.DutSciPortId < 1 || *obj.obj.DutSciPortId > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= SecureEntityStaticKeyRxSc.DutSciPortId <= 65535 but Got %d", *obj.obj.DutSciPortId))
		}

	}

	if obj.obj.DutMsbXpn != nil {

		if *obj.obj.DutMsbXpn > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= SecureEntityStaticKeyRxSc.DutMsbXpn <= 4294967295 but Got %d", *obj.obj.DutMsbXpn))
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

func (obj *secureEntityStaticKeyRxSc) setDefault() {
	if obj.obj.DutSciPortId == nil {
		obj.SetDutSciPortId(1)
	}
	if obj.obj.DutMsbXpn == nil {
		obj.SetDutMsbXpn(0)
	}

}

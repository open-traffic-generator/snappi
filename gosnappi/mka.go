package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Mka *****
type mka struct {
	validation
	obj             *otg.Mka
	marshaller      marshalMka
	unMarshaller    unMarshalMka
	basicHolder     MkaBasic
	keyServerHolder MkaKeyServer
	txscsHolder     MkaMkaTxScIter
}

func NewMka() Mka {
	obj := mka{obj: &otg.Mka{}}
	obj.setDefault()
	return &obj
}

func (obj *mka) msg() *otg.Mka {
	return obj.obj
}

func (obj *mka) setMsg(msg *otg.Mka) Mka {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmka struct {
	obj *mka
}

type marshalMka interface {
	// ToProto marshals Mka to protobuf object *otg.Mka
	ToProto() (*otg.Mka, error)
	// ToPbText marshals Mka to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Mka to YAML text
	ToYaml() (string, error)
	// ToJson marshals Mka to JSON text
	ToJson() (string, error)
}

type unMarshalmka struct {
	obj *mka
}

type unMarshalMka interface {
	// FromProto unmarshals Mka from protobuf object *otg.Mka
	FromProto(msg *otg.Mka) (Mka, error)
	// FromPbText unmarshals Mka from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Mka from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Mka from JSON text
	FromJson(value string) error
}

func (obj *mka) Marshal() marshalMka {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmka{obj: obj}
	}
	return obj.marshaller
}

func (obj *mka) Unmarshal() unMarshalMka {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmka{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmka) ToProto() (*otg.Mka, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmka) FromProto(msg *otg.Mka) (Mka, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmka) ToPbText() (string, error) {
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

func (m *unMarshalmka) FromPbText(value string) error {
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

func (m *marshalmka) ToYaml() (string, error) {
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

func (m *unMarshalmka) FromYaml(value string) error {
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

func (m *marshalmka) ToJson() (string, error) {
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

func (m *unMarshalmka) FromJson(value string) error {
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

func (obj *mka) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mka) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mka) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mka) Clone() (Mka, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMka()
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

func (obj *mka) setNil() {
	obj.basicHolder = nil
	obj.keyServerHolder = nil
	obj.txscsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Mka is configuration of a Key Agreement Entity (KaY).
type Mka interface {
	Validation
	// msg marshals Mka to protobuf object *otg.Mka
	// and doesn't set defaults
	msg() *otg.Mka
	// setMsg unmarshals Mka from protobuf object *otg.Mka
	// and doesn't set defaults
	setMsg(*otg.Mka) Mka
	// provides marshal interface
	Marshal() marshalMka
	// provides unmarshal interface
	Unmarshal() unMarshalMka
	// validate validates Mka
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Mka, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in Mka.
	Name() string
	// SetName assigns string provided by user to Mka
	SetName(value string) Mka
	// Basic returns MkaBasic, set in Mka.
	// MkaBasic is a container of basic properties for a KaY.
	Basic() MkaBasic
	// SetBasic assigns MkaBasic provided by user to Mka.
	// MkaBasic is a container of basic properties for a KaY.
	SetBasic(value MkaBasic) Mka
	// KeyServer returns MkaKeyServer, set in Mka.
	// MkaKeyServer is key server attributes of a KaY.
	KeyServer() MkaKeyServer
	// SetKeyServer assigns MkaKeyServer provided by user to Mka.
	// MkaKeyServer is key server attributes of a KaY.
	SetKeyServer(value MkaKeyServer) Mka
	// HasKeyServer checks if KeyServer has been set in Mka
	HasKeyServer() bool
	// Txscs returns MkaMkaTxScIterIter, set in Mka
	Txscs() MkaMkaTxScIter
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *mka) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the Mka object
func (obj *mka) SetName(value string) Mka {

	obj.obj.Name = &value
	return obj
}

// This contains the basic properties of KaY.
// Basic returns a MkaBasic
func (obj *mka) Basic() MkaBasic {
	if obj.obj.Basic == nil {
		obj.obj.Basic = NewMkaBasic().msg()
	}
	if obj.basicHolder == nil {
		obj.basicHolder = &mkaBasic{obj: obj.obj.Basic}
	}
	return obj.basicHolder
}

// This contains the basic properties of KaY.
// SetBasic sets the MkaBasic value in the Mka object
func (obj *mka) SetBasic(value MkaBasic) Mka {

	obj.basicHolder = nil
	obj.obj.Basic = value.msg()

	return obj
}

// Key server attributes.
// KeyServer returns a MkaKeyServer
func (obj *mka) KeyServer() MkaKeyServer {
	if obj.obj.KeyServer == nil {
		obj.obj.KeyServer = NewMkaKeyServer().msg()
	}
	if obj.keyServerHolder == nil {
		obj.keyServerHolder = &mkaKeyServer{obj: obj.obj.KeyServer}
	}
	return obj.keyServerHolder
}

// Key server attributes.
// KeyServer returns a MkaKeyServer
func (obj *mka) HasKeyServer() bool {
	return obj.obj.KeyServer != nil
}

// Key server attributes.
// SetKeyServer sets the MkaKeyServer value in the Mka object
func (obj *mka) SetKeyServer(value MkaKeyServer) Mka {

	obj.keyServerHolder = nil
	obj.obj.KeyServer = value.msg()

	return obj
}

// Properties for Tx secure channels.
// Txscs returns a []MkaTxSc
func (obj *mka) Txscs() MkaMkaTxScIter {
	if len(obj.obj.Txscs) == 0 {
		obj.obj.Txscs = []*otg.MkaTxSc{}
	}
	if obj.txscsHolder == nil {
		obj.txscsHolder = newMkaMkaTxScIter(&obj.obj.Txscs).setMsg(obj)
	}
	return obj.txscsHolder
}

type mkaMkaTxScIter struct {
	obj          *mka
	mkaTxScSlice []MkaTxSc
	fieldPtr     *[]*otg.MkaTxSc
}

func newMkaMkaTxScIter(ptr *[]*otg.MkaTxSc) MkaMkaTxScIter {
	return &mkaMkaTxScIter{fieldPtr: ptr}
}

type MkaMkaTxScIter interface {
	setMsg(*mka) MkaMkaTxScIter
	Items() []MkaTxSc
	Add() MkaTxSc
	Append(items ...MkaTxSc) MkaMkaTxScIter
	Set(index int, newObj MkaTxSc) MkaMkaTxScIter
	Clear() MkaMkaTxScIter
	clearHolderSlice() MkaMkaTxScIter
	appendHolderSlice(item MkaTxSc) MkaMkaTxScIter
}

func (obj *mkaMkaTxScIter) setMsg(msg *mka) MkaMkaTxScIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&mkaTxSc{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *mkaMkaTxScIter) Items() []MkaTxSc {
	return obj.mkaTxScSlice
}

func (obj *mkaMkaTxScIter) Add() MkaTxSc {
	newObj := &otg.MkaTxSc{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &mkaTxSc{obj: newObj}
	newLibObj.setDefault()
	obj.mkaTxScSlice = append(obj.mkaTxScSlice, newLibObj)
	return newLibObj
}

func (obj *mkaMkaTxScIter) Append(items ...MkaTxSc) MkaMkaTxScIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.mkaTxScSlice = append(obj.mkaTxScSlice, item)
	}
	return obj
}

func (obj *mkaMkaTxScIter) Set(index int, newObj MkaTxSc) MkaMkaTxScIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.mkaTxScSlice[index] = newObj
	return obj
}
func (obj *mkaMkaTxScIter) Clear() MkaMkaTxScIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.MkaTxSc{}
		obj.mkaTxScSlice = []MkaTxSc{}
	}
	return obj
}
func (obj *mkaMkaTxScIter) clearHolderSlice() MkaMkaTxScIter {
	if len(obj.mkaTxScSlice) > 0 {
		obj.mkaTxScSlice = []MkaTxSc{}
	}
	return obj
}
func (obj *mkaMkaTxScIter) appendHolderSlice(item MkaTxSc) MkaMkaTxScIter {
	obj.mkaTxScSlice = append(obj.mkaTxScSlice, item)
	return obj
}

func (obj *mka) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface Mka")
	}

	// Basic is required
	if obj.obj.Basic == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Basic is required field on interface Mka")
	}

	if obj.obj.Basic != nil {

		obj.Basic().validateObj(vObj, set_default)
	}

	if obj.obj.KeyServer != nil {

		obj.KeyServer().validateObj(vObj, set_default)
	}

	if len(obj.obj.Txscs) != 0 {

		if set_default {
			obj.Txscs().clearHolderSlice()
			for _, item := range obj.obj.Txscs {
				obj.Txscs().appendHolderSlice(&mkaTxSc{obj: item})
			}
		}
		for _, item := range obj.Txscs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *mka) setDefault() {

}

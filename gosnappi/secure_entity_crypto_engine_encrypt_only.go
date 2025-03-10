package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityCryptoEngineEncryptOnly *****
type secureEntityCryptoEngineEncryptOnly struct {
	validation
	obj                  *otg.SecureEntityCryptoEngineEncryptOnly
	marshaller           marshalSecureEntityCryptoEngineEncryptOnly
	unMarshaller         unMarshalSecureEntityCryptoEngineEncryptOnly
	secureChannelsHolder SecureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter
	trafficOptionsHolder SecureEntityCryptoEngineEncryptOnlyTrafficOptions
}

func NewSecureEntityCryptoEngineEncryptOnly() SecureEntityCryptoEngineEncryptOnly {
	obj := secureEntityCryptoEngineEncryptOnly{obj: &otg.SecureEntityCryptoEngineEncryptOnly{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityCryptoEngineEncryptOnly) msg() *otg.SecureEntityCryptoEngineEncryptOnly {
	return obj.obj
}

func (obj *secureEntityCryptoEngineEncryptOnly) setMsg(msg *otg.SecureEntityCryptoEngineEncryptOnly) SecureEntityCryptoEngineEncryptOnly {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityCryptoEngineEncryptOnly struct {
	obj *secureEntityCryptoEngineEncryptOnly
}

type marshalSecureEntityCryptoEngineEncryptOnly interface {
	// ToProto marshals SecureEntityCryptoEngineEncryptOnly to protobuf object *otg.SecureEntityCryptoEngineEncryptOnly
	ToProto() (*otg.SecureEntityCryptoEngineEncryptOnly, error)
	// ToPbText marshals SecureEntityCryptoEngineEncryptOnly to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityCryptoEngineEncryptOnly to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityCryptoEngineEncryptOnly to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals SecureEntityCryptoEngineEncryptOnly to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalsecureEntityCryptoEngineEncryptOnly struct {
	obj *secureEntityCryptoEngineEncryptOnly
}

type unMarshalSecureEntityCryptoEngineEncryptOnly interface {
	// FromProto unmarshals SecureEntityCryptoEngineEncryptOnly from protobuf object *otg.SecureEntityCryptoEngineEncryptOnly
	FromProto(msg *otg.SecureEntityCryptoEngineEncryptOnly) (SecureEntityCryptoEngineEncryptOnly, error)
	// FromPbText unmarshals SecureEntityCryptoEngineEncryptOnly from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityCryptoEngineEncryptOnly from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityCryptoEngineEncryptOnly from JSON text
	FromJson(value string) error
}

func (obj *secureEntityCryptoEngineEncryptOnly) Marshal() marshalSecureEntityCryptoEngineEncryptOnly {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityCryptoEngineEncryptOnly{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityCryptoEngineEncryptOnly) Unmarshal() unMarshalSecureEntityCryptoEngineEncryptOnly {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityCryptoEngineEncryptOnly{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityCryptoEngineEncryptOnly) ToProto() (*otg.SecureEntityCryptoEngineEncryptOnly, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityCryptoEngineEncryptOnly) FromProto(msg *otg.SecureEntityCryptoEngineEncryptOnly) (SecureEntityCryptoEngineEncryptOnly, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityCryptoEngineEncryptOnly) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptOnly) FromPbText(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptOnly) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptOnly) FromYaml(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptOnly) ToJsonRaw() (string, error) {
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

func (m *marshalsecureEntityCryptoEngineEncryptOnly) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptOnly) FromJson(value string) error {
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

func (obj *secureEntityCryptoEngineEncryptOnly) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptOnly) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptOnly) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityCryptoEngineEncryptOnly) Clone() (SecureEntityCryptoEngineEncryptOnly, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityCryptoEngineEncryptOnly()
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

func (obj *secureEntityCryptoEngineEncryptOnly) setNil() {
	obj.secureChannelsHolder = nil
	obj.trafficOptionsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// SecureEntityCryptoEngineEncryptOnly is the container for encrypt only engine configuration.
type SecureEntityCryptoEngineEncryptOnly interface {
	Validation
	// msg marshals SecureEntityCryptoEngineEncryptOnly to protobuf object *otg.SecureEntityCryptoEngineEncryptOnly
	// and doesn't set defaults
	msg() *otg.SecureEntityCryptoEngineEncryptOnly
	// setMsg unmarshals SecureEntityCryptoEngineEncryptOnly from protobuf object *otg.SecureEntityCryptoEngineEncryptOnly
	// and doesn't set defaults
	setMsg(*otg.SecureEntityCryptoEngineEncryptOnly) SecureEntityCryptoEngineEncryptOnly
	// provides marshal interface
	Marshal() marshalSecureEntityCryptoEngineEncryptOnly
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityCryptoEngineEncryptOnly
	// validate validates SecureEntityCryptoEngineEncryptOnly
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityCryptoEngineEncryptOnly, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SecureChannels returns SecureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIterIter, set in SecureEntityCryptoEngineEncryptOnly
	SecureChannels() SecureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter
	// TrafficOptions returns SecureEntityCryptoEngineEncryptOnlyTrafficOptions, set in SecureEntityCryptoEngineEncryptOnly.
	// SecureEntityCryptoEngineEncryptOnlyTrafficOptions is encrypt only traffic options.
	TrafficOptions() SecureEntityCryptoEngineEncryptOnlyTrafficOptions
	// SetTrafficOptions assigns SecureEntityCryptoEngineEncryptOnlyTrafficOptions provided by user to SecureEntityCryptoEngineEncryptOnly.
	// SecureEntityCryptoEngineEncryptOnlyTrafficOptions is encrypt only traffic options.
	SetTrafficOptions(value SecureEntityCryptoEngineEncryptOnlyTrafficOptions) SecureEntityCryptoEngineEncryptOnly
	// HasTrafficOptions checks if TrafficOptions has been set in SecureEntityCryptoEngineEncryptOnly
	HasTrafficOptions() bool
	setNil()
}

// description is TBD
// SecureChannels returns a []SecureEntityCryptoEngineEncryptOnlyTxSc
func (obj *secureEntityCryptoEngineEncryptOnly) SecureChannels() SecureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter {
	if len(obj.obj.SecureChannels) == 0 {
		obj.obj.SecureChannels = []*otg.SecureEntityCryptoEngineEncryptOnlyTxSc{}
	}
	if obj.secureChannelsHolder == nil {
		obj.secureChannelsHolder = newSecureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter(&obj.obj.SecureChannels).setMsg(obj)
	}
	return obj.secureChannelsHolder
}

type secureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter struct {
	obj                                          *secureEntityCryptoEngineEncryptOnly
	secureEntityCryptoEngineEncryptOnlyTxScSlice []SecureEntityCryptoEngineEncryptOnlyTxSc
	fieldPtr                                     *[]*otg.SecureEntityCryptoEngineEncryptOnlyTxSc
}

func newSecureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter(ptr *[]*otg.SecureEntityCryptoEngineEncryptOnlyTxSc) SecureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter {
	return &secureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter{fieldPtr: ptr}
}

type SecureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter interface {
	setMsg(*secureEntityCryptoEngineEncryptOnly) SecureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter
	Items() []SecureEntityCryptoEngineEncryptOnlyTxSc
	Add() SecureEntityCryptoEngineEncryptOnlyTxSc
	Append(items ...SecureEntityCryptoEngineEncryptOnlyTxSc) SecureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter
	Set(index int, newObj SecureEntityCryptoEngineEncryptOnlyTxSc) SecureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter
	Clear() SecureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter
	clearHolderSlice() SecureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter
	appendHolderSlice(item SecureEntityCryptoEngineEncryptOnlyTxSc) SecureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter
}

func (obj *secureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter) setMsg(msg *secureEntityCryptoEngineEncryptOnly) SecureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&secureEntityCryptoEngineEncryptOnlyTxSc{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *secureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter) Items() []SecureEntityCryptoEngineEncryptOnlyTxSc {
	return obj.secureEntityCryptoEngineEncryptOnlyTxScSlice
}

func (obj *secureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter) Add() SecureEntityCryptoEngineEncryptOnlyTxSc {
	newObj := &otg.SecureEntityCryptoEngineEncryptOnlyTxSc{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &secureEntityCryptoEngineEncryptOnlyTxSc{obj: newObj}
	newLibObj.setDefault()
	obj.secureEntityCryptoEngineEncryptOnlyTxScSlice = append(obj.secureEntityCryptoEngineEncryptOnlyTxScSlice, newLibObj)
	return newLibObj
}

func (obj *secureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter) Append(items ...SecureEntityCryptoEngineEncryptOnlyTxSc) SecureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.secureEntityCryptoEngineEncryptOnlyTxScSlice = append(obj.secureEntityCryptoEngineEncryptOnlyTxScSlice, item)
	}
	return obj
}

func (obj *secureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter) Set(index int, newObj SecureEntityCryptoEngineEncryptOnlyTxSc) SecureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.secureEntityCryptoEngineEncryptOnlyTxScSlice[index] = newObj
	return obj
}
func (obj *secureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter) Clear() SecureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.SecureEntityCryptoEngineEncryptOnlyTxSc{}
		obj.secureEntityCryptoEngineEncryptOnlyTxScSlice = []SecureEntityCryptoEngineEncryptOnlyTxSc{}
	}
	return obj
}
func (obj *secureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter) clearHolderSlice() SecureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter {
	if len(obj.secureEntityCryptoEngineEncryptOnlyTxScSlice) > 0 {
		obj.secureEntityCryptoEngineEncryptOnlyTxScSlice = []SecureEntityCryptoEngineEncryptOnlyTxSc{}
	}
	return obj
}
func (obj *secureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter) appendHolderSlice(item SecureEntityCryptoEngineEncryptOnlyTxSc) SecureEntityCryptoEngineEncryptOnlySecureEntityCryptoEngineEncryptOnlyTxScIter {
	obj.secureEntityCryptoEngineEncryptOnlyTxScSlice = append(obj.secureEntityCryptoEngineEncryptOnlyTxScSlice, item)
	return obj
}

// description is TBD
// TrafficOptions returns a SecureEntityCryptoEngineEncryptOnlyTrafficOptions
func (obj *secureEntityCryptoEngineEncryptOnly) TrafficOptions() SecureEntityCryptoEngineEncryptOnlyTrafficOptions {
	if obj.obj.TrafficOptions == nil {
		obj.obj.TrafficOptions = NewSecureEntityCryptoEngineEncryptOnlyTrafficOptions().msg()
	}
	if obj.trafficOptionsHolder == nil {
		obj.trafficOptionsHolder = &secureEntityCryptoEngineEncryptOnlyTrafficOptions{obj: obj.obj.TrafficOptions}
	}
	return obj.trafficOptionsHolder
}

// description is TBD
// TrafficOptions returns a SecureEntityCryptoEngineEncryptOnlyTrafficOptions
func (obj *secureEntityCryptoEngineEncryptOnly) HasTrafficOptions() bool {
	return obj.obj.TrafficOptions != nil
}

// description is TBD
// SetTrafficOptions sets the SecureEntityCryptoEngineEncryptOnlyTrafficOptions value in the SecureEntityCryptoEngineEncryptOnly object
func (obj *secureEntityCryptoEngineEncryptOnly) SetTrafficOptions(value SecureEntityCryptoEngineEncryptOnlyTrafficOptions) SecureEntityCryptoEngineEncryptOnly {

	obj.trafficOptionsHolder = nil
	obj.obj.TrafficOptions = value.msg()

	return obj
}

func (obj *secureEntityCryptoEngineEncryptOnly) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.SecureChannels) != 0 {

		if set_default {
			obj.SecureChannels().clearHolderSlice()
			for _, item := range obj.obj.SecureChannels {
				obj.SecureChannels().appendHolderSlice(&secureEntityCryptoEngineEncryptOnlyTxSc{obj: item})
			}
		}
		for _, item := range obj.SecureChannels().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.TrafficOptions != nil {

		obj.TrafficOptions().validateObj(vObj, set_default)
	}

}

func (obj *secureEntityCryptoEngineEncryptOnly) setDefault() {

}

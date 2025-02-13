package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecRxStaticKey *****
type macsecRxStaticKey struct {
	validation
	obj          *otg.MacsecRxStaticKey
	marshaller   marshalMacsecRxStaticKey
	unMarshaller unMarshalMacsecRxStaticKey
	scsHolder    MacsecRxStaticKeyMacsecRxScIter
}

func NewMacsecRxStaticKey() MacsecRxStaticKey {
	obj := macsecRxStaticKey{obj: &otg.MacsecRxStaticKey{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecRxStaticKey) msg() *otg.MacsecRxStaticKey {
	return obj.obj
}

func (obj *macsecRxStaticKey) setMsg(msg *otg.MacsecRxStaticKey) MacsecRxStaticKey {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecRxStaticKey struct {
	obj *macsecRxStaticKey
}

type marshalMacsecRxStaticKey interface {
	// ToProto marshals MacsecRxStaticKey to protobuf object *otg.MacsecRxStaticKey
	ToProto() (*otg.MacsecRxStaticKey, error)
	// ToPbText marshals MacsecRxStaticKey to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecRxStaticKey to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecRxStaticKey to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecRxStaticKey struct {
	obj *macsecRxStaticKey
}

type unMarshalMacsecRxStaticKey interface {
	// FromProto unmarshals MacsecRxStaticKey from protobuf object *otg.MacsecRxStaticKey
	FromProto(msg *otg.MacsecRxStaticKey) (MacsecRxStaticKey, error)
	// FromPbText unmarshals MacsecRxStaticKey from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecRxStaticKey from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecRxStaticKey from JSON text
	FromJson(value string) error
}

func (obj *macsecRxStaticKey) Marshal() marshalMacsecRxStaticKey {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecRxStaticKey{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecRxStaticKey) Unmarshal() unMarshalMacsecRxStaticKey {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecRxStaticKey{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecRxStaticKey) ToProto() (*otg.MacsecRxStaticKey, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecRxStaticKey) FromProto(msg *otg.MacsecRxStaticKey) (MacsecRxStaticKey, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecRxStaticKey) ToPbText() (string, error) {
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

func (m *unMarshalmacsecRxStaticKey) FromPbText(value string) error {
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

func (m *marshalmacsecRxStaticKey) ToYaml() (string, error) {
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

func (m *unMarshalmacsecRxStaticKey) FromYaml(value string) error {
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

func (m *marshalmacsecRxStaticKey) ToJson() (string, error) {
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

func (m *unMarshalmacsecRxStaticKey) FromJson(value string) error {
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

func (obj *macsecRxStaticKey) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecRxStaticKey) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecRxStaticKey) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecRxStaticKey) Clone() (MacsecRxStaticKey, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecRxStaticKey()
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

func (obj *macsecRxStaticKey) setNil() {
	obj.scsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecRxStaticKey is container for Rx setting for static key.
type MacsecRxStaticKey interface {
	Validation
	// msg marshals MacsecRxStaticKey to protobuf object *otg.MacsecRxStaticKey
	// and doesn't set defaults
	msg() *otg.MacsecRxStaticKey
	// setMsg unmarshals MacsecRxStaticKey from protobuf object *otg.MacsecRxStaticKey
	// and doesn't set defaults
	setMsg(*otg.MacsecRxStaticKey) MacsecRxStaticKey
	// provides marshal interface
	Marshal() marshalMacsecRxStaticKey
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecRxStaticKey
	// validate validates MacsecRxStaticKey
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecRxStaticKey, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Scs returns MacsecRxStaticKeyMacsecRxScIterIter, set in MacsecRxStaticKey
	Scs() MacsecRxStaticKeyMacsecRxScIter
	setNil()
}

// Rx secure channels.
// Scs returns a []MacsecRxSc
func (obj *macsecRxStaticKey) Scs() MacsecRxStaticKeyMacsecRxScIter {
	if len(obj.obj.Scs) == 0 {
		obj.obj.Scs = []*otg.MacsecRxSc{}
	}
	if obj.scsHolder == nil {
		obj.scsHolder = newMacsecRxStaticKeyMacsecRxScIter(&obj.obj.Scs).setMsg(obj)
	}
	return obj.scsHolder
}

type macsecRxStaticKeyMacsecRxScIter struct {
	obj             *macsecRxStaticKey
	macsecRxScSlice []MacsecRxSc
	fieldPtr        *[]*otg.MacsecRxSc
}

func newMacsecRxStaticKeyMacsecRxScIter(ptr *[]*otg.MacsecRxSc) MacsecRxStaticKeyMacsecRxScIter {
	return &macsecRxStaticKeyMacsecRxScIter{fieldPtr: ptr}
}

type MacsecRxStaticKeyMacsecRxScIter interface {
	setMsg(*macsecRxStaticKey) MacsecRxStaticKeyMacsecRxScIter
	Items() []MacsecRxSc
	Add() MacsecRxSc
	Append(items ...MacsecRxSc) MacsecRxStaticKeyMacsecRxScIter
	Set(index int, newObj MacsecRxSc) MacsecRxStaticKeyMacsecRxScIter
	Clear() MacsecRxStaticKeyMacsecRxScIter
	clearHolderSlice() MacsecRxStaticKeyMacsecRxScIter
	appendHolderSlice(item MacsecRxSc) MacsecRxStaticKeyMacsecRxScIter
}

func (obj *macsecRxStaticKeyMacsecRxScIter) setMsg(msg *macsecRxStaticKey) MacsecRxStaticKeyMacsecRxScIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&macsecRxSc{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *macsecRxStaticKeyMacsecRxScIter) Items() []MacsecRxSc {
	return obj.macsecRxScSlice
}

func (obj *macsecRxStaticKeyMacsecRxScIter) Add() MacsecRxSc {
	newObj := &otg.MacsecRxSc{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &macsecRxSc{obj: newObj}
	newLibObj.setDefault()
	obj.macsecRxScSlice = append(obj.macsecRxScSlice, newLibObj)
	return newLibObj
}

func (obj *macsecRxStaticKeyMacsecRxScIter) Append(items ...MacsecRxSc) MacsecRxStaticKeyMacsecRxScIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.macsecRxScSlice = append(obj.macsecRxScSlice, item)
	}
	return obj
}

func (obj *macsecRxStaticKeyMacsecRxScIter) Set(index int, newObj MacsecRxSc) MacsecRxStaticKeyMacsecRxScIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.macsecRxScSlice[index] = newObj
	return obj
}
func (obj *macsecRxStaticKeyMacsecRxScIter) Clear() MacsecRxStaticKeyMacsecRxScIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.MacsecRxSc{}
		obj.macsecRxScSlice = []MacsecRxSc{}
	}
	return obj
}
func (obj *macsecRxStaticKeyMacsecRxScIter) clearHolderSlice() MacsecRxStaticKeyMacsecRxScIter {
	if len(obj.macsecRxScSlice) > 0 {
		obj.macsecRxScSlice = []MacsecRxSc{}
	}
	return obj
}
func (obj *macsecRxStaticKeyMacsecRxScIter) appendHolderSlice(item MacsecRxSc) MacsecRxStaticKeyMacsecRxScIter {
	obj.macsecRxScSlice = append(obj.macsecRxScSlice, item)
	return obj
}

func (obj *macsecRxStaticKey) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Scs) != 0 {

		if set_default {
			obj.Scs().clearHolderSlice()
			for _, item := range obj.obj.Scs {
				obj.Scs().appendHolderSlice(&macsecRxSc{obj: item})
			}
		}
		for _, item := range obj.Scs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *macsecRxStaticKey) setDefault() {

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecTx *****
type macsecTx struct {
	validation
	obj             *otg.MacsecTx
	marshaller      marshalMacsecTx
	unMarshaller    unMarshalMacsecTx
	scsHolder       MacsecTxMacsecTxScIter
	staticKeyHolder MacsecTxStaticKey
}

func NewMacsecTx() MacsecTx {
	obj := macsecTx{obj: &otg.MacsecTx{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecTx) msg() *otg.MacsecTx {
	return obj.obj
}

func (obj *macsecTx) setMsg(msg *otg.MacsecTx) MacsecTx {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecTx struct {
	obj *macsecTx
}

type marshalMacsecTx interface {
	// ToProto marshals MacsecTx to protobuf object *otg.MacsecTx
	ToProto() (*otg.MacsecTx, error)
	// ToPbText marshals MacsecTx to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecTx to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecTx to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecTx struct {
	obj *macsecTx
}

type unMarshalMacsecTx interface {
	// FromProto unmarshals MacsecTx from protobuf object *otg.MacsecTx
	FromProto(msg *otg.MacsecTx) (MacsecTx, error)
	// FromPbText unmarshals MacsecTx from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecTx from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecTx from JSON text
	FromJson(value string) error
}

func (obj *macsecTx) Marshal() marshalMacsecTx {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecTx{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecTx) Unmarshal() unMarshalMacsecTx {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecTx{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecTx) ToProto() (*otg.MacsecTx, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecTx) FromProto(msg *otg.MacsecTx) (MacsecTx, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecTx) ToPbText() (string, error) {
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

func (m *unMarshalmacsecTx) FromPbText(value string) error {
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

func (m *marshalmacsecTx) ToYaml() (string, error) {
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

func (m *unMarshalmacsecTx) FromYaml(value string) error {
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

func (m *marshalmacsecTx) ToJson() (string, error) {
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

func (m *unMarshalmacsecTx) FromJson(value string) error {
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

func (obj *macsecTx) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecTx) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecTx) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecTx) Clone() (MacsecTx, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecTx()
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

func (obj *macsecTx) setNil() {
	obj.scsHolder = nil
	obj.staticKeyHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecTx is a container of Tx properties of SecY.
type MacsecTx interface {
	Validation
	// msg marshals MacsecTx to protobuf object *otg.MacsecTx
	// and doesn't set defaults
	msg() *otg.MacsecTx
	// setMsg unmarshals MacsecTx from protobuf object *otg.MacsecTx
	// and doesn't set defaults
	setMsg(*otg.MacsecTx) MacsecTx
	// provides marshal interface
	Marshal() marshalMacsecTx
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecTx
	// validate validates MacsecTx
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecTx, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Scs returns MacsecTxMacsecTxScIterIter, set in MacsecTx
	Scs() MacsecTxMacsecTxScIter
	// StaticKey returns MacsecTxStaticKey, set in MacsecTx.
	// MacsecTxStaticKey is tx setting for static key.
	StaticKey() MacsecTxStaticKey
	// SetStaticKey assigns MacsecTxStaticKey provided by user to MacsecTx.
	// MacsecTxStaticKey is tx setting for static key.
	SetStaticKey(value MacsecTxStaticKey) MacsecTx
	// HasStaticKey checks if StaticKey has been set in MacsecTx
	HasStaticKey() bool
	setNil()
}

// Tx secure channels.
// Scs returns a []MacsecTxSc
func (obj *macsecTx) Scs() MacsecTxMacsecTxScIter {
	if len(obj.obj.Scs) == 0 {
		obj.obj.Scs = []*otg.MacsecTxSc{}
	}
	if obj.scsHolder == nil {
		obj.scsHolder = newMacsecTxMacsecTxScIter(&obj.obj.Scs).setMsg(obj)
	}
	return obj.scsHolder
}

type macsecTxMacsecTxScIter struct {
	obj             *macsecTx
	macsecTxScSlice []MacsecTxSc
	fieldPtr        *[]*otg.MacsecTxSc
}

func newMacsecTxMacsecTxScIter(ptr *[]*otg.MacsecTxSc) MacsecTxMacsecTxScIter {
	return &macsecTxMacsecTxScIter{fieldPtr: ptr}
}

type MacsecTxMacsecTxScIter interface {
	setMsg(*macsecTx) MacsecTxMacsecTxScIter
	Items() []MacsecTxSc
	Add() MacsecTxSc
	Append(items ...MacsecTxSc) MacsecTxMacsecTxScIter
	Set(index int, newObj MacsecTxSc) MacsecTxMacsecTxScIter
	Clear() MacsecTxMacsecTxScIter
	clearHolderSlice() MacsecTxMacsecTxScIter
	appendHolderSlice(item MacsecTxSc) MacsecTxMacsecTxScIter
}

func (obj *macsecTxMacsecTxScIter) setMsg(msg *macsecTx) MacsecTxMacsecTxScIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&macsecTxSc{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *macsecTxMacsecTxScIter) Items() []MacsecTxSc {
	return obj.macsecTxScSlice
}

func (obj *macsecTxMacsecTxScIter) Add() MacsecTxSc {
	newObj := &otg.MacsecTxSc{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &macsecTxSc{obj: newObj}
	newLibObj.setDefault()
	obj.macsecTxScSlice = append(obj.macsecTxScSlice, newLibObj)
	return newLibObj
}

func (obj *macsecTxMacsecTxScIter) Append(items ...MacsecTxSc) MacsecTxMacsecTxScIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.macsecTxScSlice = append(obj.macsecTxScSlice, item)
	}
	return obj
}

func (obj *macsecTxMacsecTxScIter) Set(index int, newObj MacsecTxSc) MacsecTxMacsecTxScIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.macsecTxScSlice[index] = newObj
	return obj
}
func (obj *macsecTxMacsecTxScIter) Clear() MacsecTxMacsecTxScIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.MacsecTxSc{}
		obj.macsecTxScSlice = []MacsecTxSc{}
	}
	return obj
}
func (obj *macsecTxMacsecTxScIter) clearHolderSlice() MacsecTxMacsecTxScIter {
	if len(obj.macsecTxScSlice) > 0 {
		obj.macsecTxScSlice = []MacsecTxSc{}
	}
	return obj
}
func (obj *macsecTxMacsecTxScIter) appendHolderSlice(item MacsecTxSc) MacsecTxMacsecTxScIter {
	obj.macsecTxScSlice = append(obj.macsecTxScSlice, item)
	return obj
}

// description is TBD
// StaticKey returns a MacsecTxStaticKey
func (obj *macsecTx) StaticKey() MacsecTxStaticKey {
	if obj.obj.StaticKey == nil {
		obj.obj.StaticKey = NewMacsecTxStaticKey().msg()
	}
	if obj.staticKeyHolder == nil {
		obj.staticKeyHolder = &macsecTxStaticKey{obj: obj.obj.StaticKey}
	}
	return obj.staticKeyHolder
}

// description is TBD
// StaticKey returns a MacsecTxStaticKey
func (obj *macsecTx) HasStaticKey() bool {
	return obj.obj.StaticKey != nil
}

// description is TBD
// SetStaticKey sets the MacsecTxStaticKey value in the MacsecTx object
func (obj *macsecTx) SetStaticKey(value MacsecTxStaticKey) MacsecTx {

	obj.staticKeyHolder = nil
	obj.obj.StaticKey = value.msg()

	return obj
}

func (obj *macsecTx) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Scs) != 0 {

		if set_default {
			obj.Scs().clearHolderSlice()
			for _, item := range obj.obj.Scs {
				obj.Scs().appendHolderSlice(&macsecTxSc{obj: item})
			}
		}
		for _, item := range obj.Scs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.StaticKey != nil {

		obj.StaticKey().validateObj(vObj, set_default)
	}

}

func (obj *macsecTx) setDefault() {

}

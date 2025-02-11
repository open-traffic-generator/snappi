package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MkaBasicKeySourcePskChain *****
type mkaBasicKeySourcePskChain struct {
	validation
	obj          *otg.MkaBasicKeySourcePskChain
	marshaller   marshalMkaBasicKeySourcePskChain
	unMarshaller unMarshalMkaBasicKeySourcePskChain
	psksHolder   MkaBasicKeySourcePskChainMkaBasicKeySourcePskIter
}

func NewMkaBasicKeySourcePskChain() MkaBasicKeySourcePskChain {
	obj := mkaBasicKeySourcePskChain{obj: &otg.MkaBasicKeySourcePskChain{}}
	obj.setDefault()
	return &obj
}

func (obj *mkaBasicKeySourcePskChain) msg() *otg.MkaBasicKeySourcePskChain {
	return obj.obj
}

func (obj *mkaBasicKeySourcePskChain) setMsg(msg *otg.MkaBasicKeySourcePskChain) MkaBasicKeySourcePskChain {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmkaBasicKeySourcePskChain struct {
	obj *mkaBasicKeySourcePskChain
}

type marshalMkaBasicKeySourcePskChain interface {
	// ToProto marshals MkaBasicKeySourcePskChain to protobuf object *otg.MkaBasicKeySourcePskChain
	ToProto() (*otg.MkaBasicKeySourcePskChain, error)
	// ToPbText marshals MkaBasicKeySourcePskChain to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MkaBasicKeySourcePskChain to YAML text
	ToYaml() (string, error)
	// ToJson marshals MkaBasicKeySourcePskChain to JSON text
	ToJson() (string, error)
}

type unMarshalmkaBasicKeySourcePskChain struct {
	obj *mkaBasicKeySourcePskChain
}

type unMarshalMkaBasicKeySourcePskChain interface {
	// FromProto unmarshals MkaBasicKeySourcePskChain from protobuf object *otg.MkaBasicKeySourcePskChain
	FromProto(msg *otg.MkaBasicKeySourcePskChain) (MkaBasicKeySourcePskChain, error)
	// FromPbText unmarshals MkaBasicKeySourcePskChain from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MkaBasicKeySourcePskChain from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MkaBasicKeySourcePskChain from JSON text
	FromJson(value string) error
}

func (obj *mkaBasicKeySourcePskChain) Marshal() marshalMkaBasicKeySourcePskChain {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmkaBasicKeySourcePskChain{obj: obj}
	}
	return obj.marshaller
}

func (obj *mkaBasicKeySourcePskChain) Unmarshal() unMarshalMkaBasicKeySourcePskChain {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmkaBasicKeySourcePskChain{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmkaBasicKeySourcePskChain) ToProto() (*otg.MkaBasicKeySourcePskChain, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmkaBasicKeySourcePskChain) FromProto(msg *otg.MkaBasicKeySourcePskChain) (MkaBasicKeySourcePskChain, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmkaBasicKeySourcePskChain) ToPbText() (string, error) {
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

func (m *unMarshalmkaBasicKeySourcePskChain) FromPbText(value string) error {
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

func (m *marshalmkaBasicKeySourcePskChain) ToYaml() (string, error) {
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

func (m *unMarshalmkaBasicKeySourcePskChain) FromYaml(value string) error {
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

func (m *marshalmkaBasicKeySourcePskChain) ToJson() (string, error) {
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

func (m *unMarshalmkaBasicKeySourcePskChain) FromJson(value string) error {
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

func (obj *mkaBasicKeySourcePskChain) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mkaBasicKeySourcePskChain) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mkaBasicKeySourcePskChain) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mkaBasicKeySourcePskChain) Clone() (MkaBasicKeySourcePskChain, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMkaBasicKeySourcePskChain()
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

func (obj *mkaBasicKeySourcePskChain) setNil() {
	obj.psksHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MkaBasicKeySourcePskChain is the container for PSK chain settings.
type MkaBasicKeySourcePskChain interface {
	Validation
	// msg marshals MkaBasicKeySourcePskChain to protobuf object *otg.MkaBasicKeySourcePskChain
	// and doesn't set defaults
	msg() *otg.MkaBasicKeySourcePskChain
	// setMsg unmarshals MkaBasicKeySourcePskChain from protobuf object *otg.MkaBasicKeySourcePskChain
	// and doesn't set defaults
	setMsg(*otg.MkaBasicKeySourcePskChain) MkaBasicKeySourcePskChain
	// provides marshal interface
	Marshal() marshalMkaBasicKeySourcePskChain
	// provides unmarshal interface
	Unmarshal() unMarshalMkaBasicKeySourcePskChain
	// validate validates MkaBasicKeySourcePskChain
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MkaBasicKeySourcePskChain, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in MkaBasicKeySourcePskChain.
	Name() string
	// SetName assigns string provided by user to MkaBasicKeySourcePskChain
	SetName(value string) MkaBasicKeySourcePskChain
	// HasName checks if Name has been set in MkaBasicKeySourcePskChain
	HasName() bool
	// Psks returns MkaBasicKeySourcePskChainMkaBasicKeySourcePskIterIter, set in MkaBasicKeySourcePskChain
	Psks() MkaBasicKeySourcePskChainMkaBasicKeySourcePskIter
	setNil()
}

// PSK chain name.
// Name returns a string
func (obj *mkaBasicKeySourcePskChain) Name() string {

	return *obj.obj.Name

}

// PSK chain name.
// Name returns a string
func (obj *mkaBasicKeySourcePskChain) HasName() bool {
	return obj.obj.Name != nil
}

// PSK chain name.
// SetName sets the string value in the MkaBasicKeySourcePskChain object
func (obj *mkaBasicKeySourcePskChain) SetName(value string) MkaBasicKeySourcePskChain {

	obj.obj.Name = &value
	return obj
}

// Pre-shared Keys.
// Psks returns a []MkaBasicKeySourcePsk
func (obj *mkaBasicKeySourcePskChain) Psks() MkaBasicKeySourcePskChainMkaBasicKeySourcePskIter {
	if len(obj.obj.Psks) == 0 {
		obj.obj.Psks = []*otg.MkaBasicKeySourcePsk{}
	}
	if obj.psksHolder == nil {
		obj.psksHolder = newMkaBasicKeySourcePskChainMkaBasicKeySourcePskIter(&obj.obj.Psks).setMsg(obj)
	}
	return obj.psksHolder
}

type mkaBasicKeySourcePskChainMkaBasicKeySourcePskIter struct {
	obj                       *mkaBasicKeySourcePskChain
	mkaBasicKeySourcePskSlice []MkaBasicKeySourcePsk
	fieldPtr                  *[]*otg.MkaBasicKeySourcePsk
}

func newMkaBasicKeySourcePskChainMkaBasicKeySourcePskIter(ptr *[]*otg.MkaBasicKeySourcePsk) MkaBasicKeySourcePskChainMkaBasicKeySourcePskIter {
	return &mkaBasicKeySourcePskChainMkaBasicKeySourcePskIter{fieldPtr: ptr}
}

type MkaBasicKeySourcePskChainMkaBasicKeySourcePskIter interface {
	setMsg(*mkaBasicKeySourcePskChain) MkaBasicKeySourcePskChainMkaBasicKeySourcePskIter
	Items() []MkaBasicKeySourcePsk
	Add() MkaBasicKeySourcePsk
	Append(items ...MkaBasicKeySourcePsk) MkaBasicKeySourcePskChainMkaBasicKeySourcePskIter
	Set(index int, newObj MkaBasicKeySourcePsk) MkaBasicKeySourcePskChainMkaBasicKeySourcePskIter
	Clear() MkaBasicKeySourcePskChainMkaBasicKeySourcePskIter
	clearHolderSlice() MkaBasicKeySourcePskChainMkaBasicKeySourcePskIter
	appendHolderSlice(item MkaBasicKeySourcePsk) MkaBasicKeySourcePskChainMkaBasicKeySourcePskIter
}

func (obj *mkaBasicKeySourcePskChainMkaBasicKeySourcePskIter) setMsg(msg *mkaBasicKeySourcePskChain) MkaBasicKeySourcePskChainMkaBasicKeySourcePskIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&mkaBasicKeySourcePsk{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *mkaBasicKeySourcePskChainMkaBasicKeySourcePskIter) Items() []MkaBasicKeySourcePsk {
	return obj.mkaBasicKeySourcePskSlice
}

func (obj *mkaBasicKeySourcePskChainMkaBasicKeySourcePskIter) Add() MkaBasicKeySourcePsk {
	newObj := &otg.MkaBasicKeySourcePsk{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &mkaBasicKeySourcePsk{obj: newObj}
	newLibObj.setDefault()
	obj.mkaBasicKeySourcePskSlice = append(obj.mkaBasicKeySourcePskSlice, newLibObj)
	return newLibObj
}

func (obj *mkaBasicKeySourcePskChainMkaBasicKeySourcePskIter) Append(items ...MkaBasicKeySourcePsk) MkaBasicKeySourcePskChainMkaBasicKeySourcePskIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.mkaBasicKeySourcePskSlice = append(obj.mkaBasicKeySourcePskSlice, item)
	}
	return obj
}

func (obj *mkaBasicKeySourcePskChainMkaBasicKeySourcePskIter) Set(index int, newObj MkaBasicKeySourcePsk) MkaBasicKeySourcePskChainMkaBasicKeySourcePskIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.mkaBasicKeySourcePskSlice[index] = newObj
	return obj
}
func (obj *mkaBasicKeySourcePskChainMkaBasicKeySourcePskIter) Clear() MkaBasicKeySourcePskChainMkaBasicKeySourcePskIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.MkaBasicKeySourcePsk{}
		obj.mkaBasicKeySourcePskSlice = []MkaBasicKeySourcePsk{}
	}
	return obj
}
func (obj *mkaBasicKeySourcePskChainMkaBasicKeySourcePskIter) clearHolderSlice() MkaBasicKeySourcePskChainMkaBasicKeySourcePskIter {
	if len(obj.mkaBasicKeySourcePskSlice) > 0 {
		obj.mkaBasicKeySourcePskSlice = []MkaBasicKeySourcePsk{}
	}
	return obj
}
func (obj *mkaBasicKeySourcePskChainMkaBasicKeySourcePskIter) appendHolderSlice(item MkaBasicKeySourcePsk) MkaBasicKeySourcePskChainMkaBasicKeySourcePskIter {
	obj.mkaBasicKeySourcePskSlice = append(obj.mkaBasicKeySourcePskSlice, item)
	return obj
}

func (obj *mkaBasicKeySourcePskChain) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Psks) != 0 {

		if set_default {
			obj.Psks().clearHolderSlice()
			for _, item := range obj.obj.Psks {
				obj.Psks().appendHolderSlice(&mkaBasicKeySourcePsk{obj: item})
			}
		}
		for _, item := range obj.Psks().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *mkaBasicKeySourcePskChain) setDefault() {

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecStaticKeySakPool *****
type macsecStaticKeySakPool struct {
	validation
	obj          *otg.MacsecStaticKeySakPool
	marshaller   marshalMacsecStaticKeySakPool
	unMarshaller unMarshalMacsecStaticKeySakPool
	saksHolder   MacsecStaticKeySakPoolMacsecStaticKeySakIter
}

func NewMacsecStaticKeySakPool() MacsecStaticKeySakPool {
	obj := macsecStaticKeySakPool{obj: &otg.MacsecStaticKeySakPool{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecStaticKeySakPool) msg() *otg.MacsecStaticKeySakPool {
	return obj.obj
}

func (obj *macsecStaticKeySakPool) setMsg(msg *otg.MacsecStaticKeySakPool) MacsecStaticKeySakPool {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecStaticKeySakPool struct {
	obj *macsecStaticKeySakPool
}

type marshalMacsecStaticKeySakPool interface {
	// ToProto marshals MacsecStaticKeySakPool to protobuf object *otg.MacsecStaticKeySakPool
	ToProto() (*otg.MacsecStaticKeySakPool, error)
	// ToPbText marshals MacsecStaticKeySakPool to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecStaticKeySakPool to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecStaticKeySakPool to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecStaticKeySakPool struct {
	obj *macsecStaticKeySakPool
}

type unMarshalMacsecStaticKeySakPool interface {
	// FromProto unmarshals MacsecStaticKeySakPool from protobuf object *otg.MacsecStaticKeySakPool
	FromProto(msg *otg.MacsecStaticKeySakPool) (MacsecStaticKeySakPool, error)
	// FromPbText unmarshals MacsecStaticKeySakPool from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecStaticKeySakPool from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecStaticKeySakPool from JSON text
	FromJson(value string) error
}

func (obj *macsecStaticKeySakPool) Marshal() marshalMacsecStaticKeySakPool {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecStaticKeySakPool{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecStaticKeySakPool) Unmarshal() unMarshalMacsecStaticKeySakPool {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecStaticKeySakPool{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecStaticKeySakPool) ToProto() (*otg.MacsecStaticKeySakPool, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecStaticKeySakPool) FromProto(msg *otg.MacsecStaticKeySakPool) (MacsecStaticKeySakPool, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecStaticKeySakPool) ToPbText() (string, error) {
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

func (m *unMarshalmacsecStaticKeySakPool) FromPbText(value string) error {
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

func (m *marshalmacsecStaticKeySakPool) ToYaml() (string, error) {
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

func (m *unMarshalmacsecStaticKeySakPool) FromYaml(value string) error {
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

func (m *marshalmacsecStaticKeySakPool) ToJson() (string, error) {
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

func (m *unMarshalmacsecStaticKeySakPool) FromJson(value string) error {
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

func (obj *macsecStaticKeySakPool) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecStaticKeySakPool) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecStaticKeySakPool) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecStaticKeySakPool) Clone() (MacsecStaticKeySakPool, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecStaticKeySakPool()
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

func (obj *macsecStaticKeySakPool) setNil() {
	obj.saksHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecStaticKeySakPool is the container for Secure Association Key(SAK) Pool.
type MacsecStaticKeySakPool interface {
	Validation
	// msg marshals MacsecStaticKeySakPool to protobuf object *otg.MacsecStaticKeySakPool
	// and doesn't set defaults
	msg() *otg.MacsecStaticKeySakPool
	// setMsg unmarshals MacsecStaticKeySakPool from protobuf object *otg.MacsecStaticKeySakPool
	// and doesn't set defaults
	setMsg(*otg.MacsecStaticKeySakPool) MacsecStaticKeySakPool
	// provides marshal interface
	Marshal() marshalMacsecStaticKeySakPool
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecStaticKeySakPool
	// validate validates MacsecStaticKeySakPool
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecStaticKeySakPool, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in MacsecStaticKeySakPool.
	Name() string
	// SetName assigns string provided by user to MacsecStaticKeySakPool
	SetName(value string) MacsecStaticKeySakPool
	// HasName checks if Name has been set in MacsecStaticKeySakPool
	HasName() bool
	// Saks returns MacsecStaticKeySakPoolMacsecStaticKeySakIterIter, set in MacsecStaticKeySakPool
	Saks() MacsecStaticKeySakPoolMacsecStaticKeySakIter
	setNil()
}

// SAK pool name.
// Name returns a string
func (obj *macsecStaticKeySakPool) Name() string {

	return *obj.obj.Name

}

// SAK pool name.
// Name returns a string
func (obj *macsecStaticKeySakPool) HasName() bool {
	return obj.obj.Name != nil
}

// SAK pool name.
// SetName sets the string value in the MacsecStaticKeySakPool object
func (obj *macsecStaticKeySakPool) SetName(value string) MacsecStaticKeySakPool {

	obj.obj.Name = &value
	return obj
}

// Secure Association Keys.
// Saks returns a []MacsecStaticKeySak
func (obj *macsecStaticKeySakPool) Saks() MacsecStaticKeySakPoolMacsecStaticKeySakIter {
	if len(obj.obj.Saks) == 0 {
		obj.obj.Saks = []*otg.MacsecStaticKeySak{}
	}
	if obj.saksHolder == nil {
		obj.saksHolder = newMacsecStaticKeySakPoolMacsecStaticKeySakIter(&obj.obj.Saks).setMsg(obj)
	}
	return obj.saksHolder
}

type macsecStaticKeySakPoolMacsecStaticKeySakIter struct {
	obj                     *macsecStaticKeySakPool
	macsecStaticKeySakSlice []MacsecStaticKeySak
	fieldPtr                *[]*otg.MacsecStaticKeySak
}

func newMacsecStaticKeySakPoolMacsecStaticKeySakIter(ptr *[]*otg.MacsecStaticKeySak) MacsecStaticKeySakPoolMacsecStaticKeySakIter {
	return &macsecStaticKeySakPoolMacsecStaticKeySakIter{fieldPtr: ptr}
}

type MacsecStaticKeySakPoolMacsecStaticKeySakIter interface {
	setMsg(*macsecStaticKeySakPool) MacsecStaticKeySakPoolMacsecStaticKeySakIter
	Items() []MacsecStaticKeySak
	Add() MacsecStaticKeySak
	Append(items ...MacsecStaticKeySak) MacsecStaticKeySakPoolMacsecStaticKeySakIter
	Set(index int, newObj MacsecStaticKeySak) MacsecStaticKeySakPoolMacsecStaticKeySakIter
	Clear() MacsecStaticKeySakPoolMacsecStaticKeySakIter
	clearHolderSlice() MacsecStaticKeySakPoolMacsecStaticKeySakIter
	appendHolderSlice(item MacsecStaticKeySak) MacsecStaticKeySakPoolMacsecStaticKeySakIter
}

func (obj *macsecStaticKeySakPoolMacsecStaticKeySakIter) setMsg(msg *macsecStaticKeySakPool) MacsecStaticKeySakPoolMacsecStaticKeySakIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&macsecStaticKeySak{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *macsecStaticKeySakPoolMacsecStaticKeySakIter) Items() []MacsecStaticKeySak {
	return obj.macsecStaticKeySakSlice
}

func (obj *macsecStaticKeySakPoolMacsecStaticKeySakIter) Add() MacsecStaticKeySak {
	newObj := &otg.MacsecStaticKeySak{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &macsecStaticKeySak{obj: newObj}
	newLibObj.setDefault()
	obj.macsecStaticKeySakSlice = append(obj.macsecStaticKeySakSlice, newLibObj)
	return newLibObj
}

func (obj *macsecStaticKeySakPoolMacsecStaticKeySakIter) Append(items ...MacsecStaticKeySak) MacsecStaticKeySakPoolMacsecStaticKeySakIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.macsecStaticKeySakSlice = append(obj.macsecStaticKeySakSlice, item)
	}
	return obj
}

func (obj *macsecStaticKeySakPoolMacsecStaticKeySakIter) Set(index int, newObj MacsecStaticKeySak) MacsecStaticKeySakPoolMacsecStaticKeySakIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.macsecStaticKeySakSlice[index] = newObj
	return obj
}
func (obj *macsecStaticKeySakPoolMacsecStaticKeySakIter) Clear() MacsecStaticKeySakPoolMacsecStaticKeySakIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.MacsecStaticKeySak{}
		obj.macsecStaticKeySakSlice = []MacsecStaticKeySak{}
	}
	return obj
}
func (obj *macsecStaticKeySakPoolMacsecStaticKeySakIter) clearHolderSlice() MacsecStaticKeySakPoolMacsecStaticKeySakIter {
	if len(obj.macsecStaticKeySakSlice) > 0 {
		obj.macsecStaticKeySakSlice = []MacsecStaticKeySak{}
	}
	return obj
}
func (obj *macsecStaticKeySakPoolMacsecStaticKeySakIter) appendHolderSlice(item MacsecStaticKeySak) MacsecStaticKeySakPoolMacsecStaticKeySakIter {
	obj.macsecStaticKeySakSlice = append(obj.macsecStaticKeySakSlice, item)
	return obj
}

func (obj *macsecStaticKeySakPool) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Saks) != 0 {

		if set_default {
			obj.Saks().clearHolderSlice()
			for _, item := range obj.obj.Saks {
				obj.Saks().appendHolderSlice(&macsecStaticKeySak{obj: item})
			}
		}
		for _, item := range obj.Saks().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *macsecStaticKeySakPool) setDefault() {

}

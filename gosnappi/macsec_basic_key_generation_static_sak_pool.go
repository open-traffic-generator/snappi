package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecBasicKeyGenerationStaticSakPool *****
type macsecBasicKeyGenerationStaticSakPool struct {
	validation
	obj          *otg.MacsecBasicKeyGenerationStaticSakPool
	marshaller   marshalMacsecBasicKeyGenerationStaticSakPool
	unMarshaller unMarshalMacsecBasicKeyGenerationStaticSakPool
	sakHolder    MacsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter
}

func NewMacsecBasicKeyGenerationStaticSakPool() MacsecBasicKeyGenerationStaticSakPool {
	obj := macsecBasicKeyGenerationStaticSakPool{obj: &otg.MacsecBasicKeyGenerationStaticSakPool{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecBasicKeyGenerationStaticSakPool) msg() *otg.MacsecBasicKeyGenerationStaticSakPool {
	return obj.obj
}

func (obj *macsecBasicKeyGenerationStaticSakPool) setMsg(msg *otg.MacsecBasicKeyGenerationStaticSakPool) MacsecBasicKeyGenerationStaticSakPool {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecBasicKeyGenerationStaticSakPool struct {
	obj *macsecBasicKeyGenerationStaticSakPool
}

type marshalMacsecBasicKeyGenerationStaticSakPool interface {
	// ToProto marshals MacsecBasicKeyGenerationStaticSakPool to protobuf object *otg.MacsecBasicKeyGenerationStaticSakPool
	ToProto() (*otg.MacsecBasicKeyGenerationStaticSakPool, error)
	// ToPbText marshals MacsecBasicKeyGenerationStaticSakPool to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecBasicKeyGenerationStaticSakPool to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecBasicKeyGenerationStaticSakPool to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecBasicKeyGenerationStaticSakPool struct {
	obj *macsecBasicKeyGenerationStaticSakPool
}

type unMarshalMacsecBasicKeyGenerationStaticSakPool interface {
	// FromProto unmarshals MacsecBasicKeyGenerationStaticSakPool from protobuf object *otg.MacsecBasicKeyGenerationStaticSakPool
	FromProto(msg *otg.MacsecBasicKeyGenerationStaticSakPool) (MacsecBasicKeyGenerationStaticSakPool, error)
	// FromPbText unmarshals MacsecBasicKeyGenerationStaticSakPool from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecBasicKeyGenerationStaticSakPool from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecBasicKeyGenerationStaticSakPool from JSON text
	FromJson(value string) error
}

func (obj *macsecBasicKeyGenerationStaticSakPool) Marshal() marshalMacsecBasicKeyGenerationStaticSakPool {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecBasicKeyGenerationStaticSakPool{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecBasicKeyGenerationStaticSakPool) Unmarshal() unMarshalMacsecBasicKeyGenerationStaticSakPool {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecBasicKeyGenerationStaticSakPool{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecBasicKeyGenerationStaticSakPool) ToProto() (*otg.MacsecBasicKeyGenerationStaticSakPool, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecBasicKeyGenerationStaticSakPool) FromProto(msg *otg.MacsecBasicKeyGenerationStaticSakPool) (MacsecBasicKeyGenerationStaticSakPool, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecBasicKeyGenerationStaticSakPool) ToPbText() (string, error) {
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

func (m *unMarshalmacsecBasicKeyGenerationStaticSakPool) FromPbText(value string) error {
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

func (m *marshalmacsecBasicKeyGenerationStaticSakPool) ToYaml() (string, error) {
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

func (m *unMarshalmacsecBasicKeyGenerationStaticSakPool) FromYaml(value string) error {
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

func (m *marshalmacsecBasicKeyGenerationStaticSakPool) ToJson() (string, error) {
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

func (m *unMarshalmacsecBasicKeyGenerationStaticSakPool) FromJson(value string) error {
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

func (obj *macsecBasicKeyGenerationStaticSakPool) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecBasicKeyGenerationStaticSakPool) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecBasicKeyGenerationStaticSakPool) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecBasicKeyGenerationStaticSakPool) Clone() (MacsecBasicKeyGenerationStaticSakPool, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecBasicKeyGenerationStaticSakPool()
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

func (obj *macsecBasicKeyGenerationStaticSakPool) setNil() {
	obj.sakHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecBasicKeyGenerationStaticSakPool is the container for SAK Pool.
type MacsecBasicKeyGenerationStaticSakPool interface {
	Validation
	// msg marshals MacsecBasicKeyGenerationStaticSakPool to protobuf object *otg.MacsecBasicKeyGenerationStaticSakPool
	// and doesn't set defaults
	msg() *otg.MacsecBasicKeyGenerationStaticSakPool
	// setMsg unmarshals MacsecBasicKeyGenerationStaticSakPool from protobuf object *otg.MacsecBasicKeyGenerationStaticSakPool
	// and doesn't set defaults
	setMsg(*otg.MacsecBasicKeyGenerationStaticSakPool) MacsecBasicKeyGenerationStaticSakPool
	// provides marshal interface
	Marshal() marshalMacsecBasicKeyGenerationStaticSakPool
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecBasicKeyGenerationStaticSakPool
	// validate validates MacsecBasicKeyGenerationStaticSakPool
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecBasicKeyGenerationStaticSakPool, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in MacsecBasicKeyGenerationStaticSakPool.
	Name() string
	// SetName assigns string provided by user to MacsecBasicKeyGenerationStaticSakPool
	SetName(value string) MacsecBasicKeyGenerationStaticSakPool
	// HasName checks if Name has been set in MacsecBasicKeyGenerationStaticSakPool
	HasName() bool
	// Sak returns MacsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIterIter, set in MacsecBasicKeyGenerationStaticSakPool
	Sak() MacsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter
	setNil()
}

// SAK pool name.
// Name returns a string
func (obj *macsecBasicKeyGenerationStaticSakPool) Name() string {

	return *obj.obj.Name

}

// SAK pool name.
// Name returns a string
func (obj *macsecBasicKeyGenerationStaticSakPool) HasName() bool {
	return obj.obj.Name != nil
}

// SAK pool name.
// SetName sets the string value in the MacsecBasicKeyGenerationStaticSakPool object
func (obj *macsecBasicKeyGenerationStaticSakPool) SetName(value string) MacsecBasicKeyGenerationStaticSakPool {

	obj.obj.Name = &value
	return obj
}

// SAK.
// Sak returns a []MacsecBasicKeyGenerationStaticSak
func (obj *macsecBasicKeyGenerationStaticSakPool) Sak() MacsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter {
	if len(obj.obj.Sak) == 0 {
		obj.obj.Sak = []*otg.MacsecBasicKeyGenerationStaticSak{}
	}
	if obj.sakHolder == nil {
		obj.sakHolder = newMacsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter(&obj.obj.Sak).setMsg(obj)
	}
	return obj.sakHolder
}

type macsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter struct {
	obj                                    *macsecBasicKeyGenerationStaticSakPool
	macsecBasicKeyGenerationStaticSakSlice []MacsecBasicKeyGenerationStaticSak
	fieldPtr                               *[]*otg.MacsecBasicKeyGenerationStaticSak
}

func newMacsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter(ptr *[]*otg.MacsecBasicKeyGenerationStaticSak) MacsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter {
	return &macsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter{fieldPtr: ptr}
}

type MacsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter interface {
	setMsg(*macsecBasicKeyGenerationStaticSakPool) MacsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter
	Items() []MacsecBasicKeyGenerationStaticSak
	Add() MacsecBasicKeyGenerationStaticSak
	Append(items ...MacsecBasicKeyGenerationStaticSak) MacsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter
	Set(index int, newObj MacsecBasicKeyGenerationStaticSak) MacsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter
	Clear() MacsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter
	clearHolderSlice() MacsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter
	appendHolderSlice(item MacsecBasicKeyGenerationStaticSak) MacsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter
}

func (obj *macsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter) setMsg(msg *macsecBasicKeyGenerationStaticSakPool) MacsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&macsecBasicKeyGenerationStaticSak{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *macsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter) Items() []MacsecBasicKeyGenerationStaticSak {
	return obj.macsecBasicKeyGenerationStaticSakSlice
}

func (obj *macsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter) Add() MacsecBasicKeyGenerationStaticSak {
	newObj := &otg.MacsecBasicKeyGenerationStaticSak{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &macsecBasicKeyGenerationStaticSak{obj: newObj}
	newLibObj.setDefault()
	obj.macsecBasicKeyGenerationStaticSakSlice = append(obj.macsecBasicKeyGenerationStaticSakSlice, newLibObj)
	return newLibObj
}

func (obj *macsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter) Append(items ...MacsecBasicKeyGenerationStaticSak) MacsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.macsecBasicKeyGenerationStaticSakSlice = append(obj.macsecBasicKeyGenerationStaticSakSlice, item)
	}
	return obj
}

func (obj *macsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter) Set(index int, newObj MacsecBasicKeyGenerationStaticSak) MacsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.macsecBasicKeyGenerationStaticSakSlice[index] = newObj
	return obj
}
func (obj *macsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter) Clear() MacsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.MacsecBasicKeyGenerationStaticSak{}
		obj.macsecBasicKeyGenerationStaticSakSlice = []MacsecBasicKeyGenerationStaticSak{}
	}
	return obj
}
func (obj *macsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter) clearHolderSlice() MacsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter {
	if len(obj.macsecBasicKeyGenerationStaticSakSlice) > 0 {
		obj.macsecBasicKeyGenerationStaticSakSlice = []MacsecBasicKeyGenerationStaticSak{}
	}
	return obj
}
func (obj *macsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter) appendHolderSlice(item MacsecBasicKeyGenerationStaticSak) MacsecBasicKeyGenerationStaticSakPoolMacsecBasicKeyGenerationStaticSakIter {
	obj.macsecBasicKeyGenerationStaticSakSlice = append(obj.macsecBasicKeyGenerationStaticSakSlice, item)
	return obj
}

func (obj *macsecBasicKeyGenerationStaticSakPool) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Sak) != 0 {

		if set_default {
			obj.Sak().clearHolderSlice()
			for _, item := range obj.obj.Sak {
				obj.Sak().appendHolderSlice(&macsecBasicKeyGenerationStaticSak{obj: item})
			}
		}
		for _, item := range obj.Sak().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *macsecBasicKeyGenerationStaticSakPool) setDefault() {

}

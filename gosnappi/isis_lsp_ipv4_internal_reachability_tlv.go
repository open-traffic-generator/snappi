package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspIpv4InternalReachabilityTlv *****
type isisLspIpv4InternalReachabilityTlv struct {
	validation
	obj            *otg.IsisLspIpv4InternalReachabilityTlv
	marshaller     marshalIsisLspIpv4InternalReachabilityTlv
	unMarshaller   unMarshalIsisLspIpv4InternalReachabilityTlv
	prefixesHolder IsisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter
}

func NewIsisLspIpv4InternalReachabilityTlv() IsisLspIpv4InternalReachabilityTlv {
	obj := isisLspIpv4InternalReachabilityTlv{obj: &otg.IsisLspIpv4InternalReachabilityTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspIpv4InternalReachabilityTlv) msg() *otg.IsisLspIpv4InternalReachabilityTlv {
	return obj.obj
}

func (obj *isisLspIpv4InternalReachabilityTlv) setMsg(msg *otg.IsisLspIpv4InternalReachabilityTlv) IsisLspIpv4InternalReachabilityTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspIpv4InternalReachabilityTlv struct {
	obj *isisLspIpv4InternalReachabilityTlv
}

type marshalIsisLspIpv4InternalReachabilityTlv interface {
	// ToProto marshals IsisLspIpv4InternalReachabilityTlv to protobuf object *otg.IsisLspIpv4InternalReachabilityTlv
	ToProto() (*otg.IsisLspIpv4InternalReachabilityTlv, error)
	// ToPbText marshals IsisLspIpv4InternalReachabilityTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspIpv4InternalReachabilityTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspIpv4InternalReachabilityTlv to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspIpv4InternalReachabilityTlv struct {
	obj *isisLspIpv4InternalReachabilityTlv
}

type unMarshalIsisLspIpv4InternalReachabilityTlv interface {
	// FromProto unmarshals IsisLspIpv4InternalReachabilityTlv from protobuf object *otg.IsisLspIpv4InternalReachabilityTlv
	FromProto(msg *otg.IsisLspIpv4InternalReachabilityTlv) (IsisLspIpv4InternalReachabilityTlv, error)
	// FromPbText unmarshals IsisLspIpv4InternalReachabilityTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspIpv4InternalReachabilityTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspIpv4InternalReachabilityTlv from JSON text
	FromJson(value string) error
}

func (obj *isisLspIpv4InternalReachabilityTlv) Marshal() marshalIsisLspIpv4InternalReachabilityTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspIpv4InternalReachabilityTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspIpv4InternalReachabilityTlv) Unmarshal() unMarshalIsisLspIpv4InternalReachabilityTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspIpv4InternalReachabilityTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspIpv4InternalReachabilityTlv) ToProto() (*otg.IsisLspIpv4InternalReachabilityTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspIpv4InternalReachabilityTlv) FromProto(msg *otg.IsisLspIpv4InternalReachabilityTlv) (IsisLspIpv4InternalReachabilityTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspIpv4InternalReachabilityTlv) ToPbText() (string, error) {
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

func (m *unMarshalisisLspIpv4InternalReachabilityTlv) FromPbText(value string) error {
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

func (m *marshalisisLspIpv4InternalReachabilityTlv) ToYaml() (string, error) {
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

func (m *unMarshalisisLspIpv4InternalReachabilityTlv) FromYaml(value string) error {
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

func (m *marshalisisLspIpv4InternalReachabilityTlv) ToJson() (string, error) {
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

func (m *unMarshalisisLspIpv4InternalReachabilityTlv) FromJson(value string) error {
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

func (obj *isisLspIpv4InternalReachabilityTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspIpv4InternalReachabilityTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspIpv4InternalReachabilityTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspIpv4InternalReachabilityTlv) Clone() (IsisLspIpv4InternalReachabilityTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspIpv4InternalReachabilityTlv()
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

func (obj *isisLspIpv4InternalReachabilityTlv) setNil() {
	obj.prefixesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspIpv4InternalReachabilityTlv is this container defines list of IPv4 internal reachability information in one IPv4 internal reachability TLV.
// This is advertised when the origin-type is set 'internal' in route range configurations.
type IsisLspIpv4InternalReachabilityTlv interface {
	Validation
	// msg marshals IsisLspIpv4InternalReachabilityTlv to protobuf object *otg.IsisLspIpv4InternalReachabilityTlv
	// and doesn't set defaults
	msg() *otg.IsisLspIpv4InternalReachabilityTlv
	// setMsg unmarshals IsisLspIpv4InternalReachabilityTlv from protobuf object *otg.IsisLspIpv4InternalReachabilityTlv
	// and doesn't set defaults
	setMsg(*otg.IsisLspIpv4InternalReachabilityTlv) IsisLspIpv4InternalReachabilityTlv
	// provides marshal interface
	Marshal() marshalIsisLspIpv4InternalReachabilityTlv
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspIpv4InternalReachabilityTlv
	// validate validates IsisLspIpv4InternalReachabilityTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspIpv4InternalReachabilityTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Prefixes returns IsisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIterIter, set in IsisLspIpv4InternalReachabilityTlv
	Prefixes() IsisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter
	setNil()
}

// Describes list of IPv4 prefixes in this TLV.
// Prefixes returns a []IsisLspV4Prefix
func (obj *isisLspIpv4InternalReachabilityTlv) Prefixes() IsisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter {
	if len(obj.obj.Prefixes) == 0 {
		obj.obj.Prefixes = []*otg.IsisLspV4Prefix{}
	}
	if obj.prefixesHolder == nil {
		obj.prefixesHolder = newIsisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter(&obj.obj.Prefixes).setMsg(obj)
	}
	return obj.prefixesHolder
}

type isisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter struct {
	obj                  *isisLspIpv4InternalReachabilityTlv
	isisLspV4PrefixSlice []IsisLspV4Prefix
	fieldPtr             *[]*otg.IsisLspV4Prefix
}

func newIsisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter(ptr *[]*otg.IsisLspV4Prefix) IsisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter {
	return &isisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter{fieldPtr: ptr}
}

type IsisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter interface {
	setMsg(*isisLspIpv4InternalReachabilityTlv) IsisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter
	Items() []IsisLspV4Prefix
	Add() IsisLspV4Prefix
	Append(items ...IsisLspV4Prefix) IsisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter
	Set(index int, newObj IsisLspV4Prefix) IsisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter
	Clear() IsisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter
	clearHolderSlice() IsisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter
	appendHolderSlice(item IsisLspV4Prefix) IsisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter
}

func (obj *isisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter) setMsg(msg *isisLspIpv4InternalReachabilityTlv) IsisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspV4Prefix{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter) Items() []IsisLspV4Prefix {
	return obj.isisLspV4PrefixSlice
}

func (obj *isisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter) Add() IsisLspV4Prefix {
	newObj := &otg.IsisLspV4Prefix{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspV4Prefix{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspV4PrefixSlice = append(obj.isisLspV4PrefixSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter) Append(items ...IsisLspV4Prefix) IsisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspV4PrefixSlice = append(obj.isisLspV4PrefixSlice, item)
	}
	return obj
}

func (obj *isisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter) Set(index int, newObj IsisLspV4Prefix) IsisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspV4PrefixSlice[index] = newObj
	return obj
}
func (obj *isisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter) Clear() IsisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspV4Prefix{}
		obj.isisLspV4PrefixSlice = []IsisLspV4Prefix{}
	}
	return obj
}
func (obj *isisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter) clearHolderSlice() IsisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter {
	if len(obj.isisLspV4PrefixSlice) > 0 {
		obj.isisLspV4PrefixSlice = []IsisLspV4Prefix{}
	}
	return obj
}
func (obj *isisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter) appendHolderSlice(item IsisLspV4Prefix) IsisLspIpv4InternalReachabilityTlvIsisLspV4PrefixIter {
	obj.isisLspV4PrefixSlice = append(obj.isisLspV4PrefixSlice, item)
	return obj
}

func (obj *isisLspIpv4InternalReachabilityTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Prefixes) != 0 {

		if set_default {
			obj.Prefixes().clearHolderSlice()
			for _, item := range obj.obj.Prefixes {
				obj.Prefixes().appendHolderSlice(&isisLspV4Prefix{obj: item})
			}
		}
		for _, item := range obj.Prefixes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *isisLspIpv4InternalReachabilityTlv) setDefault() {

}

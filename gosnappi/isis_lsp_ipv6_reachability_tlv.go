package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspIpv6ReachabilityTlv *****
type isisLspIpv6ReachabilityTlv struct {
	validation
	obj            *otg.IsisLspIpv6ReachabilityTlv
	marshaller     marshalIsisLspIpv6ReachabilityTlv
	unMarshaller   unMarshalIsisLspIpv6ReachabilityTlv
	prefixesHolder IsisLspIpv6ReachabilityTlvIsisLspV6PrefixIter
}

func NewIsisLspIpv6ReachabilityTlv() IsisLspIpv6ReachabilityTlv {
	obj := isisLspIpv6ReachabilityTlv{obj: &otg.IsisLspIpv6ReachabilityTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspIpv6ReachabilityTlv) msg() *otg.IsisLspIpv6ReachabilityTlv {
	return obj.obj
}

func (obj *isisLspIpv6ReachabilityTlv) setMsg(msg *otg.IsisLspIpv6ReachabilityTlv) IsisLspIpv6ReachabilityTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspIpv6ReachabilityTlv struct {
	obj *isisLspIpv6ReachabilityTlv
}

type marshalIsisLspIpv6ReachabilityTlv interface {
	// ToProto marshals IsisLspIpv6ReachabilityTlv to protobuf object *otg.IsisLspIpv6ReachabilityTlv
	ToProto() (*otg.IsisLspIpv6ReachabilityTlv, error)
	// ToPbText marshals IsisLspIpv6ReachabilityTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspIpv6ReachabilityTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspIpv6ReachabilityTlv to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisLspIpv6ReachabilityTlv to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisLspIpv6ReachabilityTlv struct {
	obj *isisLspIpv6ReachabilityTlv
}

type unMarshalIsisLspIpv6ReachabilityTlv interface {
	// FromProto unmarshals IsisLspIpv6ReachabilityTlv from protobuf object *otg.IsisLspIpv6ReachabilityTlv
	FromProto(msg *otg.IsisLspIpv6ReachabilityTlv) (IsisLspIpv6ReachabilityTlv, error)
	// FromPbText unmarshals IsisLspIpv6ReachabilityTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspIpv6ReachabilityTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspIpv6ReachabilityTlv from JSON text
	FromJson(value string) error
}

func (obj *isisLspIpv6ReachabilityTlv) Marshal() marshalIsisLspIpv6ReachabilityTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspIpv6ReachabilityTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspIpv6ReachabilityTlv) Unmarshal() unMarshalIsisLspIpv6ReachabilityTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspIpv6ReachabilityTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspIpv6ReachabilityTlv) ToProto() (*otg.IsisLspIpv6ReachabilityTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspIpv6ReachabilityTlv) FromProto(msg *otg.IsisLspIpv6ReachabilityTlv) (IsisLspIpv6ReachabilityTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspIpv6ReachabilityTlv) ToPbText() (string, error) {
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

func (m *unMarshalisisLspIpv6ReachabilityTlv) FromPbText(value string) error {
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

func (m *marshalisisLspIpv6ReachabilityTlv) ToYaml() (string, error) {
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

func (m *unMarshalisisLspIpv6ReachabilityTlv) FromYaml(value string) error {
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

func (m *marshalisisLspIpv6ReachabilityTlv) ToJsonRaw() (string, error) {
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

func (m *marshalisisLspIpv6ReachabilityTlv) ToJson() (string, error) {
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

func (m *unMarshalisisLspIpv6ReachabilityTlv) FromJson(value string) error {
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

func (obj *isisLspIpv6ReachabilityTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspIpv6ReachabilityTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspIpv6ReachabilityTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspIpv6ReachabilityTlv) Clone() (IsisLspIpv6ReachabilityTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspIpv6ReachabilityTlv()
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

func (obj *isisLspIpv6ReachabilityTlv) setNil() {
	obj.prefixesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspIpv6ReachabilityTlv is it defines list of IPv6 extended reachability information in one IPv6 Reachability TLV.
type IsisLspIpv6ReachabilityTlv interface {
	Validation
	// msg marshals IsisLspIpv6ReachabilityTlv to protobuf object *otg.IsisLspIpv6ReachabilityTlv
	// and doesn't set defaults
	msg() *otg.IsisLspIpv6ReachabilityTlv
	// setMsg unmarshals IsisLspIpv6ReachabilityTlv from protobuf object *otg.IsisLspIpv6ReachabilityTlv
	// and doesn't set defaults
	setMsg(*otg.IsisLspIpv6ReachabilityTlv) IsisLspIpv6ReachabilityTlv
	// provides marshal interface
	Marshal() marshalIsisLspIpv6ReachabilityTlv
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspIpv6ReachabilityTlv
	// validate validates IsisLspIpv6ReachabilityTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspIpv6ReachabilityTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Prefixes returns IsisLspIpv6ReachabilityTlvIsisLspV6PrefixIterIter, set in IsisLspIpv6ReachabilityTlv
	Prefixes() IsisLspIpv6ReachabilityTlvIsisLspV6PrefixIter
	setNil()
}

// IPv6 prefix contained within reachability TLVs.
// Prefixes returns a []IsisLspV6Prefix
func (obj *isisLspIpv6ReachabilityTlv) Prefixes() IsisLspIpv6ReachabilityTlvIsisLspV6PrefixIter {
	if len(obj.obj.Prefixes) == 0 {
		obj.obj.Prefixes = []*otg.IsisLspV6Prefix{}
	}
	if obj.prefixesHolder == nil {
		obj.prefixesHolder = newIsisLspIpv6ReachabilityTlvIsisLspV6PrefixIter(&obj.obj.Prefixes).setMsg(obj)
	}
	return obj.prefixesHolder
}

type isisLspIpv6ReachabilityTlvIsisLspV6PrefixIter struct {
	obj                  *isisLspIpv6ReachabilityTlv
	isisLspV6PrefixSlice []IsisLspV6Prefix
	fieldPtr             *[]*otg.IsisLspV6Prefix
}

func newIsisLspIpv6ReachabilityTlvIsisLspV6PrefixIter(ptr *[]*otg.IsisLspV6Prefix) IsisLspIpv6ReachabilityTlvIsisLspV6PrefixIter {
	return &isisLspIpv6ReachabilityTlvIsisLspV6PrefixIter{fieldPtr: ptr}
}

type IsisLspIpv6ReachabilityTlvIsisLspV6PrefixIter interface {
	setMsg(*isisLspIpv6ReachabilityTlv) IsisLspIpv6ReachabilityTlvIsisLspV6PrefixIter
	Items() []IsisLspV6Prefix
	Add() IsisLspV6Prefix
	Append(items ...IsisLspV6Prefix) IsisLspIpv6ReachabilityTlvIsisLspV6PrefixIter
	Set(index int, newObj IsisLspV6Prefix) IsisLspIpv6ReachabilityTlvIsisLspV6PrefixIter
	Clear() IsisLspIpv6ReachabilityTlvIsisLspV6PrefixIter
	clearHolderSlice() IsisLspIpv6ReachabilityTlvIsisLspV6PrefixIter
	appendHolderSlice(item IsisLspV6Prefix) IsisLspIpv6ReachabilityTlvIsisLspV6PrefixIter
}

func (obj *isisLspIpv6ReachabilityTlvIsisLspV6PrefixIter) setMsg(msg *isisLspIpv6ReachabilityTlv) IsisLspIpv6ReachabilityTlvIsisLspV6PrefixIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspV6Prefix{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspIpv6ReachabilityTlvIsisLspV6PrefixIter) Items() []IsisLspV6Prefix {
	return obj.isisLspV6PrefixSlice
}

func (obj *isisLspIpv6ReachabilityTlvIsisLspV6PrefixIter) Add() IsisLspV6Prefix {
	newObj := &otg.IsisLspV6Prefix{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspV6Prefix{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspV6PrefixSlice = append(obj.isisLspV6PrefixSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspIpv6ReachabilityTlvIsisLspV6PrefixIter) Append(items ...IsisLspV6Prefix) IsisLspIpv6ReachabilityTlvIsisLspV6PrefixIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspV6PrefixSlice = append(obj.isisLspV6PrefixSlice, item)
	}
	return obj
}

func (obj *isisLspIpv6ReachabilityTlvIsisLspV6PrefixIter) Set(index int, newObj IsisLspV6Prefix) IsisLspIpv6ReachabilityTlvIsisLspV6PrefixIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspV6PrefixSlice[index] = newObj
	return obj
}
func (obj *isisLspIpv6ReachabilityTlvIsisLspV6PrefixIter) Clear() IsisLspIpv6ReachabilityTlvIsisLspV6PrefixIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspV6Prefix{}
		obj.isisLspV6PrefixSlice = []IsisLspV6Prefix{}
	}
	return obj
}
func (obj *isisLspIpv6ReachabilityTlvIsisLspV6PrefixIter) clearHolderSlice() IsisLspIpv6ReachabilityTlvIsisLspV6PrefixIter {
	if len(obj.isisLspV6PrefixSlice) > 0 {
		obj.isisLspV6PrefixSlice = []IsisLspV6Prefix{}
	}
	return obj
}
func (obj *isisLspIpv6ReachabilityTlvIsisLspV6PrefixIter) appendHolderSlice(item IsisLspV6Prefix) IsisLspIpv6ReachabilityTlvIsisLspV6PrefixIter {
	obj.isisLspV6PrefixSlice = append(obj.isisLspV6PrefixSlice, item)
	return obj
}

func (obj *isisLspIpv6ReachabilityTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Prefixes) != 0 {

		if set_default {
			obj.Prefixes().clearHolderSlice()
			for _, item := range obj.obj.Prefixes {
				obj.Prefixes().appendHolderSlice(&isisLspV6Prefix{obj: item})
			}
		}
		for _, item := range obj.Prefixes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *isisLspIpv6ReachabilityTlv) setDefault() {

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspIpv4ExternalReachabilityTlv *****
type isisLspIpv4ExternalReachabilityTlv struct {
	validation
	obj            *otg.IsisLspIpv4ExternalReachabilityTlv
	marshaller     marshalIsisLspIpv4ExternalReachabilityTlv
	unMarshaller   unMarshalIsisLspIpv4ExternalReachabilityTlv
	prefixesHolder IsisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter
}

func NewIsisLspIpv4ExternalReachabilityTlv() IsisLspIpv4ExternalReachabilityTlv {
	obj := isisLspIpv4ExternalReachabilityTlv{obj: &otg.IsisLspIpv4ExternalReachabilityTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspIpv4ExternalReachabilityTlv) msg() *otg.IsisLspIpv4ExternalReachabilityTlv {
	return obj.obj
}

func (obj *isisLspIpv4ExternalReachabilityTlv) setMsg(msg *otg.IsisLspIpv4ExternalReachabilityTlv) IsisLspIpv4ExternalReachabilityTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspIpv4ExternalReachabilityTlv struct {
	obj *isisLspIpv4ExternalReachabilityTlv
}

type marshalIsisLspIpv4ExternalReachabilityTlv interface {
	// ToProto marshals IsisLspIpv4ExternalReachabilityTlv to protobuf object *otg.IsisLspIpv4ExternalReachabilityTlv
	ToProto() (*otg.IsisLspIpv4ExternalReachabilityTlv, error)
	// ToPbText marshals IsisLspIpv4ExternalReachabilityTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspIpv4ExternalReachabilityTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspIpv4ExternalReachabilityTlv to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisLspIpv4ExternalReachabilityTlv to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisLspIpv4ExternalReachabilityTlv struct {
	obj *isisLspIpv4ExternalReachabilityTlv
}

type unMarshalIsisLspIpv4ExternalReachabilityTlv interface {
	// FromProto unmarshals IsisLspIpv4ExternalReachabilityTlv from protobuf object *otg.IsisLspIpv4ExternalReachabilityTlv
	FromProto(msg *otg.IsisLspIpv4ExternalReachabilityTlv) (IsisLspIpv4ExternalReachabilityTlv, error)
	// FromPbText unmarshals IsisLspIpv4ExternalReachabilityTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspIpv4ExternalReachabilityTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspIpv4ExternalReachabilityTlv from JSON text
	FromJson(value string) error
}

func (obj *isisLspIpv4ExternalReachabilityTlv) Marshal() marshalIsisLspIpv4ExternalReachabilityTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspIpv4ExternalReachabilityTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspIpv4ExternalReachabilityTlv) Unmarshal() unMarshalIsisLspIpv4ExternalReachabilityTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspIpv4ExternalReachabilityTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspIpv4ExternalReachabilityTlv) ToProto() (*otg.IsisLspIpv4ExternalReachabilityTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspIpv4ExternalReachabilityTlv) FromProto(msg *otg.IsisLspIpv4ExternalReachabilityTlv) (IsisLspIpv4ExternalReachabilityTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspIpv4ExternalReachabilityTlv) ToPbText() (string, error) {
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

func (m *unMarshalisisLspIpv4ExternalReachabilityTlv) FromPbText(value string) error {
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

func (m *marshalisisLspIpv4ExternalReachabilityTlv) ToYaml() (string, error) {
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

func (m *unMarshalisisLspIpv4ExternalReachabilityTlv) FromYaml(value string) error {
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

func (m *marshalisisLspIpv4ExternalReachabilityTlv) ToJsonRaw() (string, error) {
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

func (m *marshalisisLspIpv4ExternalReachabilityTlv) ToJson() (string, error) {
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

func (m *unMarshalisisLspIpv4ExternalReachabilityTlv) FromJson(value string) error {
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

func (obj *isisLspIpv4ExternalReachabilityTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspIpv4ExternalReachabilityTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspIpv4ExternalReachabilityTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspIpv4ExternalReachabilityTlv) Clone() (IsisLspIpv4ExternalReachabilityTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspIpv4ExternalReachabilityTlv()
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

func (obj *isisLspIpv4ExternalReachabilityTlv) setNil() {
	obj.prefixesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspIpv4ExternalReachabilityTlv is this container defines list of IPv4 external reachability information in one IPv4 external reachability TLV.
// This is advertised when the origin-type is set 'external' in route range configurations.
type IsisLspIpv4ExternalReachabilityTlv interface {
	Validation
	// msg marshals IsisLspIpv4ExternalReachabilityTlv to protobuf object *otg.IsisLspIpv4ExternalReachabilityTlv
	// and doesn't set defaults
	msg() *otg.IsisLspIpv4ExternalReachabilityTlv
	// setMsg unmarshals IsisLspIpv4ExternalReachabilityTlv from protobuf object *otg.IsisLspIpv4ExternalReachabilityTlv
	// and doesn't set defaults
	setMsg(*otg.IsisLspIpv4ExternalReachabilityTlv) IsisLspIpv4ExternalReachabilityTlv
	// provides marshal interface
	Marshal() marshalIsisLspIpv4ExternalReachabilityTlv
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspIpv4ExternalReachabilityTlv
	// validate validates IsisLspIpv4ExternalReachabilityTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspIpv4ExternalReachabilityTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Prefixes returns IsisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIterIter, set in IsisLspIpv4ExternalReachabilityTlv
	Prefixes() IsisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter
	setNil()
}

// Describes list of IPv4 prefixes in this TLV..
// Prefixes returns a []IsisLspV4Prefix
func (obj *isisLspIpv4ExternalReachabilityTlv) Prefixes() IsisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter {
	if len(obj.obj.Prefixes) == 0 {
		obj.obj.Prefixes = []*otg.IsisLspV4Prefix{}
	}
	if obj.prefixesHolder == nil {
		obj.prefixesHolder = newIsisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter(&obj.obj.Prefixes).setMsg(obj)
	}
	return obj.prefixesHolder
}

type isisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter struct {
	obj                  *isisLspIpv4ExternalReachabilityTlv
	isisLspV4PrefixSlice []IsisLspV4Prefix
	fieldPtr             *[]*otg.IsisLspV4Prefix
}

func newIsisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter(ptr *[]*otg.IsisLspV4Prefix) IsisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter {
	return &isisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter{fieldPtr: ptr}
}

type IsisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter interface {
	setMsg(*isisLspIpv4ExternalReachabilityTlv) IsisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter
	Items() []IsisLspV4Prefix
	Add() IsisLspV4Prefix
	Append(items ...IsisLspV4Prefix) IsisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter
	Set(index int, newObj IsisLspV4Prefix) IsisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter
	Clear() IsisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter
	clearHolderSlice() IsisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter
	appendHolderSlice(item IsisLspV4Prefix) IsisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter
}

func (obj *isisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter) setMsg(msg *isisLspIpv4ExternalReachabilityTlv) IsisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspV4Prefix{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter) Items() []IsisLspV4Prefix {
	return obj.isisLspV4PrefixSlice
}

func (obj *isisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter) Add() IsisLspV4Prefix {
	newObj := &otg.IsisLspV4Prefix{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspV4Prefix{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspV4PrefixSlice = append(obj.isisLspV4PrefixSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter) Append(items ...IsisLspV4Prefix) IsisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspV4PrefixSlice = append(obj.isisLspV4PrefixSlice, item)
	}
	return obj
}

func (obj *isisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter) Set(index int, newObj IsisLspV4Prefix) IsisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspV4PrefixSlice[index] = newObj
	return obj
}
func (obj *isisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter) Clear() IsisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspV4Prefix{}
		obj.isisLspV4PrefixSlice = []IsisLspV4Prefix{}
	}
	return obj
}
func (obj *isisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter) clearHolderSlice() IsisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter {
	if len(obj.isisLspV4PrefixSlice) > 0 {
		obj.isisLspV4PrefixSlice = []IsisLspV4Prefix{}
	}
	return obj
}
func (obj *isisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter) appendHolderSlice(item IsisLspV4Prefix) IsisLspIpv4ExternalReachabilityTlvIsisLspV4PrefixIter {
	obj.isisLspV4PrefixSlice = append(obj.isisLspV4PrefixSlice, item)
	return obj
}

func (obj *isisLspIpv4ExternalReachabilityTlv) validateObj(vObj *validation, set_default bool) {
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

func (obj *isisLspIpv4ExternalReachabilityTlv) setDefault() {

}

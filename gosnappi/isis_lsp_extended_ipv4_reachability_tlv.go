package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspExtendedIpv4ReachabilityTlv *****
type isisLspExtendedIpv4ReachabilityTlv struct {
	validation
	obj            *otg.IsisLspExtendedIpv4ReachabilityTlv
	marshaller     marshalIsisLspExtendedIpv4ReachabilityTlv
	unMarshaller   unMarshalIsisLspExtendedIpv4ReachabilityTlv
	prefixesHolder IsisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter
}

func NewIsisLspExtendedIpv4ReachabilityTlv() IsisLspExtendedIpv4ReachabilityTlv {
	obj := isisLspExtendedIpv4ReachabilityTlv{obj: &otg.IsisLspExtendedIpv4ReachabilityTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspExtendedIpv4ReachabilityTlv) msg() *otg.IsisLspExtendedIpv4ReachabilityTlv {
	return obj.obj
}

func (obj *isisLspExtendedIpv4ReachabilityTlv) setMsg(msg *otg.IsisLspExtendedIpv4ReachabilityTlv) IsisLspExtendedIpv4ReachabilityTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspExtendedIpv4ReachabilityTlv struct {
	obj *isisLspExtendedIpv4ReachabilityTlv
}

type marshalIsisLspExtendedIpv4ReachabilityTlv interface {
	// ToProto marshals IsisLspExtendedIpv4ReachabilityTlv to protobuf object *otg.IsisLspExtendedIpv4ReachabilityTlv
	ToProto() (*otg.IsisLspExtendedIpv4ReachabilityTlv, error)
	// ToPbText marshals IsisLspExtendedIpv4ReachabilityTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspExtendedIpv4ReachabilityTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspExtendedIpv4ReachabilityTlv to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisLspExtendedIpv4ReachabilityTlv to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisLspExtendedIpv4ReachabilityTlv struct {
	obj *isisLspExtendedIpv4ReachabilityTlv
}

type unMarshalIsisLspExtendedIpv4ReachabilityTlv interface {
	// FromProto unmarshals IsisLspExtendedIpv4ReachabilityTlv from protobuf object *otg.IsisLspExtendedIpv4ReachabilityTlv
	FromProto(msg *otg.IsisLspExtendedIpv4ReachabilityTlv) (IsisLspExtendedIpv4ReachabilityTlv, error)
	// FromPbText unmarshals IsisLspExtendedIpv4ReachabilityTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspExtendedIpv4ReachabilityTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspExtendedIpv4ReachabilityTlv from JSON text
	FromJson(value string) error
}

func (obj *isisLspExtendedIpv4ReachabilityTlv) Marshal() marshalIsisLspExtendedIpv4ReachabilityTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspExtendedIpv4ReachabilityTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspExtendedIpv4ReachabilityTlv) Unmarshal() unMarshalIsisLspExtendedIpv4ReachabilityTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspExtendedIpv4ReachabilityTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspExtendedIpv4ReachabilityTlv) ToProto() (*otg.IsisLspExtendedIpv4ReachabilityTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspExtendedIpv4ReachabilityTlv) FromProto(msg *otg.IsisLspExtendedIpv4ReachabilityTlv) (IsisLspExtendedIpv4ReachabilityTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspExtendedIpv4ReachabilityTlv) ToPbText() (string, error) {
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

func (m *unMarshalisisLspExtendedIpv4ReachabilityTlv) FromPbText(value string) error {
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

func (m *marshalisisLspExtendedIpv4ReachabilityTlv) ToYaml() (string, error) {
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

func (m *unMarshalisisLspExtendedIpv4ReachabilityTlv) FromYaml(value string) error {
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

func (m *marshalisisLspExtendedIpv4ReachabilityTlv) ToJsonRaw() (string, error) {
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

func (m *marshalisisLspExtendedIpv4ReachabilityTlv) ToJson() (string, error) {
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

func (m *unMarshalisisLspExtendedIpv4ReachabilityTlv) FromJson(value string) error {
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

func (obj *isisLspExtendedIpv4ReachabilityTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspExtendedIpv4ReachabilityTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspExtendedIpv4ReachabilityTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspExtendedIpv4ReachabilityTlv) Clone() (IsisLspExtendedIpv4ReachabilityTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspExtendedIpv4ReachabilityTlv()
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

func (obj *isisLspExtendedIpv4ReachabilityTlv) setNil() {
	obj.prefixesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspExtendedIpv4ReachabilityTlv is this container defines list of IPv4 extended reachability information in one Extended IPv4 External Reachability TLV.
// It is advertised when the 'wide metric' is enabled.
type IsisLspExtendedIpv4ReachabilityTlv interface {
	Validation
	// msg marshals IsisLspExtendedIpv4ReachabilityTlv to protobuf object *otg.IsisLspExtendedIpv4ReachabilityTlv
	// and doesn't set defaults
	msg() *otg.IsisLspExtendedIpv4ReachabilityTlv
	// setMsg unmarshals IsisLspExtendedIpv4ReachabilityTlv from protobuf object *otg.IsisLspExtendedIpv4ReachabilityTlv
	// and doesn't set defaults
	setMsg(*otg.IsisLspExtendedIpv4ReachabilityTlv) IsisLspExtendedIpv4ReachabilityTlv
	// provides marshal interface
	Marshal() marshalIsisLspExtendedIpv4ReachabilityTlv
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspExtendedIpv4ReachabilityTlv
	// validate validates IsisLspExtendedIpv4ReachabilityTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspExtendedIpv4ReachabilityTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Prefixes returns IsisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIterIter, set in IsisLspExtendedIpv4ReachabilityTlv
	Prefixes() IsisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter
	setNil()
}

// IPv4 prefix contained within extended reachability TLVs.
// Prefixes returns a []IsisLspExtendedV4Prefix
func (obj *isisLspExtendedIpv4ReachabilityTlv) Prefixes() IsisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter {
	if len(obj.obj.Prefixes) == 0 {
		obj.obj.Prefixes = []*otg.IsisLspExtendedV4Prefix{}
	}
	if obj.prefixesHolder == nil {
		obj.prefixesHolder = newIsisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter(&obj.obj.Prefixes).setMsg(obj)
	}
	return obj.prefixesHolder
}

type isisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter struct {
	obj                          *isisLspExtendedIpv4ReachabilityTlv
	isisLspExtendedV4PrefixSlice []IsisLspExtendedV4Prefix
	fieldPtr                     *[]*otg.IsisLspExtendedV4Prefix
}

func newIsisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter(ptr *[]*otg.IsisLspExtendedV4Prefix) IsisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter {
	return &isisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter{fieldPtr: ptr}
}

type IsisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter interface {
	setMsg(*isisLspExtendedIpv4ReachabilityTlv) IsisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter
	Items() []IsisLspExtendedV4Prefix
	Add() IsisLspExtendedV4Prefix
	Append(items ...IsisLspExtendedV4Prefix) IsisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter
	Set(index int, newObj IsisLspExtendedV4Prefix) IsisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter
	Clear() IsisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter
	clearHolderSlice() IsisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter
	appendHolderSlice(item IsisLspExtendedV4Prefix) IsisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter
}

func (obj *isisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter) setMsg(msg *isisLspExtendedIpv4ReachabilityTlv) IsisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspExtendedV4Prefix{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter) Items() []IsisLspExtendedV4Prefix {
	return obj.isisLspExtendedV4PrefixSlice
}

func (obj *isisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter) Add() IsisLspExtendedV4Prefix {
	newObj := &otg.IsisLspExtendedV4Prefix{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspExtendedV4Prefix{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspExtendedV4PrefixSlice = append(obj.isisLspExtendedV4PrefixSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter) Append(items ...IsisLspExtendedV4Prefix) IsisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspExtendedV4PrefixSlice = append(obj.isisLspExtendedV4PrefixSlice, item)
	}
	return obj
}

func (obj *isisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter) Set(index int, newObj IsisLspExtendedV4Prefix) IsisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspExtendedV4PrefixSlice[index] = newObj
	return obj
}
func (obj *isisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter) Clear() IsisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspExtendedV4Prefix{}
		obj.isisLspExtendedV4PrefixSlice = []IsisLspExtendedV4Prefix{}
	}
	return obj
}
func (obj *isisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter) clearHolderSlice() IsisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter {
	if len(obj.isisLspExtendedV4PrefixSlice) > 0 {
		obj.isisLspExtendedV4PrefixSlice = []IsisLspExtendedV4Prefix{}
	}
	return obj
}
func (obj *isisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter) appendHolderSlice(item IsisLspExtendedV4Prefix) IsisLspExtendedIpv4ReachabilityTlvIsisLspExtendedV4PrefixIter {
	obj.isisLspExtendedV4PrefixSlice = append(obj.isisLspExtendedV4PrefixSlice, item)
	return obj
}

func (obj *isisLspExtendedIpv4ReachabilityTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Prefixes) != 0 {

		if set_default {
			obj.Prefixes().clearHolderSlice()
			for _, item := range obj.obj.Prefixes {
				obj.Prefixes().appendHolderSlice(&isisLspExtendedV4Prefix{obj: item})
			}
		}
		for _, item := range obj.Prefixes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *isisLspExtendedIpv4ReachabilityTlv) setDefault() {

}

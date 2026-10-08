package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LdpIpv4FecRange *****
type ldpIpv4FecRange struct {
	validation
	obj             *otg.LdpIpv4FecRange
	marshaller      marshalLdpIpv4FecRange
	unMarshaller    unMarshalLdpIpv4FecRange
	addressesHolder LdpIpv4FecRangeV4RouteAddressIter
	labelHolder     LdpFecLabel
}

func NewLdpIpv4FecRange() LdpIpv4FecRange {
	obj := ldpIpv4FecRange{obj: &otg.LdpIpv4FecRange{}}
	obj.setDefault()
	return &obj
}

func (obj *ldpIpv4FecRange) msg() *otg.LdpIpv4FecRange {
	return obj.obj
}

func (obj *ldpIpv4FecRange) setMsg(msg *otg.LdpIpv4FecRange) LdpIpv4FecRange {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalldpIpv4FecRange struct {
	obj *ldpIpv4FecRange
}

type marshalLdpIpv4FecRange interface {
	// ToProto marshals LdpIpv4FecRange to protobuf object *otg.LdpIpv4FecRange
	ToProto() (*otg.LdpIpv4FecRange, error)
	// ToPbText marshals LdpIpv4FecRange to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LdpIpv4FecRange to YAML text
	ToYaml() (string, error)
	// ToJson marshals LdpIpv4FecRange to JSON text
	ToJson() (string, error)
}

type unMarshalldpIpv4FecRange struct {
	obj *ldpIpv4FecRange
}

type unMarshalLdpIpv4FecRange interface {
	// FromProto unmarshals LdpIpv4FecRange from protobuf object *otg.LdpIpv4FecRange
	FromProto(msg *otg.LdpIpv4FecRange) (LdpIpv4FecRange, error)
	// FromPbText unmarshals LdpIpv4FecRange from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LdpIpv4FecRange from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LdpIpv4FecRange from JSON text
	FromJson(value string) error
}

func (obj *ldpIpv4FecRange) Marshal() marshalLdpIpv4FecRange {
	if obj.marshaller == nil {
		obj.marshaller = &marshalldpIpv4FecRange{obj: obj}
	}
	return obj.marshaller
}

func (obj *ldpIpv4FecRange) Unmarshal() unMarshalLdpIpv4FecRange {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalldpIpv4FecRange{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalldpIpv4FecRange) ToProto() (*otg.LdpIpv4FecRange, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalldpIpv4FecRange) FromProto(msg *otg.LdpIpv4FecRange) (LdpIpv4FecRange, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalldpIpv4FecRange) ToPbText() (string, error) {
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

func (m *unMarshalldpIpv4FecRange) FromPbText(value string) error {
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

func (m *marshalldpIpv4FecRange) ToYaml() (string, error) {
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

func (m *unMarshalldpIpv4FecRange) FromYaml(value string) error {
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

func (m *marshalldpIpv4FecRange) ToJson() (string, error) {
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

func (m *unMarshalldpIpv4FecRange) FromJson(value string) error {
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

func (obj *ldpIpv4FecRange) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ldpIpv4FecRange) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ldpIpv4FecRange) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ldpIpv4FecRange) Clone() (LdpIpv4FecRange, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLdpIpv4FecRange()
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

func (obj *ldpIpv4FecRange) setNil() {
	obj.addressesHolder = nil
	obj.labelHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LdpIpv4FecRange is a range of IPv4 Prefix FEC elements (FEC element type 0x02, RFC 5036 Section 3.4.1) that the emulated LSR advertises to all its peers in Label Mapping messages (RFC 5036 Section 3.5.7), each bound to a Generic Label (RFC 5036 Section 3.4.2.1). A DUT normally installs a label for a FEC only when it has a matching route (RFC 5036 Section 2.6.1), so the same prefixes are usually also advertised through an IGP, for example an Isis.V4RouteRange on the same device.
type LdpIpv4FecRange interface {
	Validation
	// msg marshals LdpIpv4FecRange to protobuf object *otg.LdpIpv4FecRange
	// and doesn't set defaults
	msg() *otg.LdpIpv4FecRange
	// setMsg unmarshals LdpIpv4FecRange from protobuf object *otg.LdpIpv4FecRange
	// and doesn't set defaults
	setMsg(*otg.LdpIpv4FecRange) LdpIpv4FecRange
	// provides marshal interface
	Marshal() marshalLdpIpv4FecRange
	// provides unmarshal interface
	Unmarshal() unMarshalLdpIpv4FecRange
	// validate validates LdpIpv4FecRange
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LdpIpv4FecRange, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in LdpIpv4FecRange.
	Name() string
	// SetName assigns string provided by user to LdpIpv4FecRange
	SetName(value string) LdpIpv4FecRange
	// Addresses returns LdpIpv4FecRangeV4RouteAddressIterIter, set in LdpIpv4FecRange
	Addresses() LdpIpv4FecRangeV4RouteAddressIter
	// Label returns LdpFecLabel, set in LdpIpv4FecRange.
	// LdpFecLabel is the Generic Label (RFC 5036 Section 3.4.2.1) bound to each Prefix FEC of a Ldp.Ipv4FecRange. increment - the n-th prefix generated by addresses (counting across all entries, in order, from 0) gets label start + n. fixed - every prefix of the range gets the same label. Use 3 (Implicit NULL) or 0 (IPv4 Explicit NULL) to emulate an egress LSR (RFC 3032 Section 2.1).
	Label() LdpFecLabel
	// SetLabel assigns LdpFecLabel provided by user to LdpIpv4FecRange.
	// LdpFecLabel is the Generic Label (RFC 5036 Section 3.4.2.1) bound to each Prefix FEC of a Ldp.Ipv4FecRange. increment - the n-th prefix generated by addresses (counting across all entries, in order, from 0) gets label start + n. fixed - every prefix of the range gets the same label. Use 3 (Implicit NULL) or 0 (IPv4 Explicit NULL) to emulate an egress LSR (RFC 3032 Section 2.1).
	SetLabel(value LdpFecLabel) LdpIpv4FecRange
	// HasLabel checks if Label has been set in LdpIpv4FecRange
	HasLabel() bool
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *ldpIpv4FecRange) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the LdpIpv4FecRange object
func (obj *ldpIpv4FecRange) SetName(value string) LdpIpv4FecRange {

	obj.obj.Name = &value
	return obj
}

// A list of groups of IPv4 prefixes. Each generated prefix is one Prefix FEC element.
// Addresses returns a []V4RouteAddress
func (obj *ldpIpv4FecRange) Addresses() LdpIpv4FecRangeV4RouteAddressIter {
	if len(obj.obj.Addresses) == 0 {
		obj.obj.Addresses = []*otg.V4RouteAddress{}
	}
	if obj.addressesHolder == nil {
		obj.addressesHolder = newLdpIpv4FecRangeV4RouteAddressIter(&obj.obj.Addresses).setMsg(obj)
	}
	return obj.addressesHolder
}

type ldpIpv4FecRangeV4RouteAddressIter struct {
	obj                 *ldpIpv4FecRange
	v4RouteAddressSlice []V4RouteAddress
	fieldPtr            *[]*otg.V4RouteAddress
}

func newLdpIpv4FecRangeV4RouteAddressIter(ptr *[]*otg.V4RouteAddress) LdpIpv4FecRangeV4RouteAddressIter {
	return &ldpIpv4FecRangeV4RouteAddressIter{fieldPtr: ptr}
}

type LdpIpv4FecRangeV4RouteAddressIter interface {
	setMsg(*ldpIpv4FecRange) LdpIpv4FecRangeV4RouteAddressIter
	Items() []V4RouteAddress
	Add() V4RouteAddress
	Append(items ...V4RouteAddress) LdpIpv4FecRangeV4RouteAddressIter
	Set(index int, newObj V4RouteAddress) LdpIpv4FecRangeV4RouteAddressIter
	Clear() LdpIpv4FecRangeV4RouteAddressIter
	clearHolderSlice() LdpIpv4FecRangeV4RouteAddressIter
	appendHolderSlice(item V4RouteAddress) LdpIpv4FecRangeV4RouteAddressIter
}

func (obj *ldpIpv4FecRangeV4RouteAddressIter) setMsg(msg *ldpIpv4FecRange) LdpIpv4FecRangeV4RouteAddressIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&v4RouteAddress{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ldpIpv4FecRangeV4RouteAddressIter) Items() []V4RouteAddress {
	return obj.v4RouteAddressSlice
}

func (obj *ldpIpv4FecRangeV4RouteAddressIter) Add() V4RouteAddress {
	newObj := &otg.V4RouteAddress{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &v4RouteAddress{obj: newObj}
	newLibObj.setDefault()
	obj.v4RouteAddressSlice = append(obj.v4RouteAddressSlice, newLibObj)
	return newLibObj
}

func (obj *ldpIpv4FecRangeV4RouteAddressIter) Append(items ...V4RouteAddress) LdpIpv4FecRangeV4RouteAddressIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.v4RouteAddressSlice = append(obj.v4RouteAddressSlice, item)
	}
	return obj
}

func (obj *ldpIpv4FecRangeV4RouteAddressIter) Set(index int, newObj V4RouteAddress) LdpIpv4FecRangeV4RouteAddressIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.v4RouteAddressSlice[index] = newObj
	return obj
}
func (obj *ldpIpv4FecRangeV4RouteAddressIter) Clear() LdpIpv4FecRangeV4RouteAddressIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.V4RouteAddress{}
		obj.v4RouteAddressSlice = []V4RouteAddress{}
	}
	return obj
}
func (obj *ldpIpv4FecRangeV4RouteAddressIter) clearHolderSlice() LdpIpv4FecRangeV4RouteAddressIter {
	if len(obj.v4RouteAddressSlice) > 0 {
		obj.v4RouteAddressSlice = []V4RouteAddress{}
	}
	return obj
}
func (obj *ldpIpv4FecRangeV4RouteAddressIter) appendHolderSlice(item V4RouteAddress) LdpIpv4FecRangeV4RouteAddressIter {
	obj.v4RouteAddressSlice = append(obj.v4RouteAddressSlice, item)
	return obj
}

// The labels bound to the prefixes of this range.
// Label returns a LdpFecLabel
func (obj *ldpIpv4FecRange) Label() LdpFecLabel {
	if obj.obj.Label == nil {
		obj.obj.Label = NewLdpFecLabel().msg()
	}
	if obj.labelHolder == nil {
		obj.labelHolder = &ldpFecLabel{obj: obj.obj.Label}
	}
	return obj.labelHolder
}

// The labels bound to the prefixes of this range.
// Label returns a LdpFecLabel
func (obj *ldpIpv4FecRange) HasLabel() bool {
	return obj.obj.Label != nil
}

// The labels bound to the prefixes of this range.
// SetLabel sets the LdpFecLabel value in the LdpIpv4FecRange object
func (obj *ldpIpv4FecRange) SetLabel(value LdpFecLabel) LdpIpv4FecRange {

	obj.labelHolder = nil
	obj.obj.Label = value.msg()

	return obj
}

func (obj *ldpIpv4FecRange) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface LdpIpv4FecRange")
	}

	if len(obj.obj.Addresses) != 0 {

		if set_default {
			obj.Addresses().clearHolderSlice()
			for _, item := range obj.obj.Addresses {
				obj.Addresses().appendHolderSlice(&v4RouteAddress{obj: item})
			}
		}
		for _, item := range obj.Addresses().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Label != nil {

		obj.Label().validateObj(vObj, set_default)
	}

}

func (obj *ldpIpv4FecRange) setDefault() {

}

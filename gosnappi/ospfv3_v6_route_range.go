package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3V6RouteRange *****
type ospfv3V6RouteRange struct {
	validation
	obj               *otg.Ospfv3V6RouteRange
	marshaller        marshalOspfv3V6RouteRange
	unMarshaller      unMarshalOspfv3V6RouteRange
	addressesHolder   Ospfv3V6RouteRangeV6RouteAddressIter
	routeOriginHolder Ospfv3V6RRRouteOrigin
}

func NewOspfv3V6RouteRange() Ospfv3V6RouteRange {
	obj := ospfv3V6RouteRange{obj: &otg.Ospfv3V6RouteRange{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3V6RouteRange) msg() *otg.Ospfv3V6RouteRange {
	return obj.obj
}

func (obj *ospfv3V6RouteRange) setMsg(msg *otg.Ospfv3V6RouteRange) Ospfv3V6RouteRange {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3V6RouteRange struct {
	obj *ospfv3V6RouteRange
}

type marshalOspfv3V6RouteRange interface {
	// ToProto marshals Ospfv3V6RouteRange to protobuf object *otg.Ospfv3V6RouteRange
	ToProto() (*otg.Ospfv3V6RouteRange, error)
	// ToPbText marshals Ospfv3V6RouteRange to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3V6RouteRange to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3V6RouteRange to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3V6RouteRange struct {
	obj *ospfv3V6RouteRange
}

type unMarshalOspfv3V6RouteRange interface {
	// FromProto unmarshals Ospfv3V6RouteRange from protobuf object *otg.Ospfv3V6RouteRange
	FromProto(msg *otg.Ospfv3V6RouteRange) (Ospfv3V6RouteRange, error)
	// FromPbText unmarshals Ospfv3V6RouteRange from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3V6RouteRange from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3V6RouteRange from JSON text
	FromJson(value string) error
}

func (obj *ospfv3V6RouteRange) Marshal() marshalOspfv3V6RouteRange {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3V6RouteRange{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3V6RouteRange) Unmarshal() unMarshalOspfv3V6RouteRange {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3V6RouteRange{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3V6RouteRange) ToProto() (*otg.Ospfv3V6RouteRange, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3V6RouteRange) FromProto(msg *otg.Ospfv3V6RouteRange) (Ospfv3V6RouteRange, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3V6RouteRange) ToPbText() (string, error) {
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

func (m *unMarshalospfv3V6RouteRange) FromPbText(value string) error {
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

func (m *marshalospfv3V6RouteRange) ToYaml() (string, error) {
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

func (m *unMarshalospfv3V6RouteRange) FromYaml(value string) error {
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

func (m *marshalospfv3V6RouteRange) ToJson() (string, error) {
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

func (m *unMarshalospfv3V6RouteRange) FromJson(value string) error {
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

func (obj *ospfv3V6RouteRange) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3V6RouteRange) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3V6RouteRange) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3V6RouteRange) Clone() (Ospfv3V6RouteRange, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3V6RouteRange()
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

func (obj *ospfv3V6RouteRange) setNil() {
	obj.addressesHolder = nil
	obj.routeOriginHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv3V6RouteRange is emulated OSPFv3 IPv6 routes.
type Ospfv3V6RouteRange interface {
	Validation
	// msg marshals Ospfv3V6RouteRange to protobuf object *otg.Ospfv3V6RouteRange
	// and doesn't set defaults
	msg() *otg.Ospfv3V6RouteRange
	// setMsg unmarshals Ospfv3V6RouteRange from protobuf object *otg.Ospfv3V6RouteRange
	// and doesn't set defaults
	setMsg(*otg.Ospfv3V6RouteRange) Ospfv3V6RouteRange
	// provides marshal interface
	Marshal() marshalOspfv3V6RouteRange
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3V6RouteRange
	// validate validates Ospfv3V6RouteRange
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3V6RouteRange, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in Ospfv3V6RouteRange.
	Name() string
	// SetName assigns string provided by user to Ospfv3V6RouteRange
	SetName(value string) Ospfv3V6RouteRange
	// Addresses returns Ospfv3V6RouteRangeV6RouteAddressIterIter, set in Ospfv3V6RouteRange
	Addresses() Ospfv3V6RouteRangeV6RouteAddressIter
	// Metric returns uint32, set in Ospfv3V6RouteRange.
	Metric() uint32
	// SetMetric assigns uint32 provided by user to Ospfv3V6RouteRange
	SetMetric(value uint32) Ospfv3V6RouteRange
	// HasMetric checks if Metric has been set in Ospfv3V6RouteRange
	HasMetric() bool
	// RouteOrigin returns Ospfv3V6RRRouteOrigin, set in Ospfv3V6RouteRange.
	// Ospfv3V6RRRouteOrigin is container of type of the OSPFv3 types correspond directly to the OSPFv3 LSAs types.
	RouteOrigin() Ospfv3V6RRRouteOrigin
	// SetRouteOrigin assigns Ospfv3V6RRRouteOrigin provided by user to Ospfv3V6RouteRange.
	// Ospfv3V6RRRouteOrigin is container of type of the OSPFv3 types correspond directly to the OSPFv3 LSAs types.
	SetRouteOrigin(value Ospfv3V6RRRouteOrigin) Ospfv3V6RouteRange
	// HasRouteOrigin checks if RouteOrigin has been set in Ospfv3V6RouteRange
	HasRouteOrigin() bool
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *ospfv3V6RouteRange) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the Ospfv3V6RouteRange object
func (obj *ospfv3V6RouteRange) SetName(value string) Ospfv3V6RouteRange {

	obj.obj.Name = &value
	return obj
}

// A list of group of IPv6 route addresses.
// Addresses returns a []V6RouteAddress
func (obj *ospfv3V6RouteRange) Addresses() Ospfv3V6RouteRangeV6RouteAddressIter {
	if len(obj.obj.Addresses) == 0 {
		obj.obj.Addresses = []*otg.V6RouteAddress{}
	}
	if obj.addressesHolder == nil {
		obj.addressesHolder = newOspfv3V6RouteRangeV6RouteAddressIter(&obj.obj.Addresses).setMsg(obj)
	}
	return obj.addressesHolder
}

type ospfv3V6RouteRangeV6RouteAddressIter struct {
	obj                 *ospfv3V6RouteRange
	v6RouteAddressSlice []V6RouteAddress
	fieldPtr            *[]*otg.V6RouteAddress
}

func newOspfv3V6RouteRangeV6RouteAddressIter(ptr *[]*otg.V6RouteAddress) Ospfv3V6RouteRangeV6RouteAddressIter {
	return &ospfv3V6RouteRangeV6RouteAddressIter{fieldPtr: ptr}
}

type Ospfv3V6RouteRangeV6RouteAddressIter interface {
	setMsg(*ospfv3V6RouteRange) Ospfv3V6RouteRangeV6RouteAddressIter
	Items() []V6RouteAddress
	Add() V6RouteAddress
	Append(items ...V6RouteAddress) Ospfv3V6RouteRangeV6RouteAddressIter
	Set(index int, newObj V6RouteAddress) Ospfv3V6RouteRangeV6RouteAddressIter
	Clear() Ospfv3V6RouteRangeV6RouteAddressIter
	clearHolderSlice() Ospfv3V6RouteRangeV6RouteAddressIter
	appendHolderSlice(item V6RouteAddress) Ospfv3V6RouteRangeV6RouteAddressIter
}

func (obj *ospfv3V6RouteRangeV6RouteAddressIter) setMsg(msg *ospfv3V6RouteRange) Ospfv3V6RouteRangeV6RouteAddressIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&v6RouteAddress{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv3V6RouteRangeV6RouteAddressIter) Items() []V6RouteAddress {
	return obj.v6RouteAddressSlice
}

func (obj *ospfv3V6RouteRangeV6RouteAddressIter) Add() V6RouteAddress {
	newObj := &otg.V6RouteAddress{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &v6RouteAddress{obj: newObj}
	newLibObj.setDefault()
	obj.v6RouteAddressSlice = append(obj.v6RouteAddressSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv3V6RouteRangeV6RouteAddressIter) Append(items ...V6RouteAddress) Ospfv3V6RouteRangeV6RouteAddressIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.v6RouteAddressSlice = append(obj.v6RouteAddressSlice, item)
	}
	return obj
}

func (obj *ospfv3V6RouteRangeV6RouteAddressIter) Set(index int, newObj V6RouteAddress) Ospfv3V6RouteRangeV6RouteAddressIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.v6RouteAddressSlice[index] = newObj
	return obj
}
func (obj *ospfv3V6RouteRangeV6RouteAddressIter) Clear() Ospfv3V6RouteRangeV6RouteAddressIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.V6RouteAddress{}
		obj.v6RouteAddressSlice = []V6RouteAddress{}
	}
	return obj
}
func (obj *ospfv3V6RouteRangeV6RouteAddressIter) clearHolderSlice() Ospfv3V6RouteRangeV6RouteAddressIter {
	if len(obj.v6RouteAddressSlice) > 0 {
		obj.v6RouteAddressSlice = []V6RouteAddress{}
	}
	return obj
}
func (obj *ospfv3V6RouteRangeV6RouteAddressIter) appendHolderSlice(item V6RouteAddress) Ospfv3V6RouteRangeV6RouteAddressIter {
	obj.v6RouteAddressSlice = append(obj.v6RouteAddressSlice, item)
	return obj
}

// The user-defined metric associated with this route range.
// Metric returns a uint32
func (obj *ospfv3V6RouteRange) Metric() uint32 {

	return *obj.obj.Metric

}

// The user-defined metric associated with this route range.
// Metric returns a uint32
func (obj *ospfv3V6RouteRange) HasMetric() bool {
	return obj.obj.Metric != nil
}

// The user-defined metric associated with this route range.
// SetMetric sets the uint32 value in the Ospfv3V6RouteRange object
func (obj *ospfv3V6RouteRange) SetMetric(value uint32) Ospfv3V6RouteRange {

	obj.obj.Metric = &value
	return obj
}

// The type of the OSPFv3 routes.
// RouteOrigin returns a Ospfv3V6RRRouteOrigin
func (obj *ospfv3V6RouteRange) RouteOrigin() Ospfv3V6RRRouteOrigin {
	if obj.obj.RouteOrigin == nil {
		obj.obj.RouteOrigin = NewOspfv3V6RRRouteOrigin().msg()
	}
	if obj.routeOriginHolder == nil {
		obj.routeOriginHolder = &ospfv3V6RRRouteOrigin{obj: obj.obj.RouteOrigin}
	}
	return obj.routeOriginHolder
}

// The type of the OSPFv3 routes.
// RouteOrigin returns a Ospfv3V6RRRouteOrigin
func (obj *ospfv3V6RouteRange) HasRouteOrigin() bool {
	return obj.obj.RouteOrigin != nil
}

// The type of the OSPFv3 routes.
// SetRouteOrigin sets the Ospfv3V6RRRouteOrigin value in the Ospfv3V6RouteRange object
func (obj *ospfv3V6RouteRange) SetRouteOrigin(value Ospfv3V6RRRouteOrigin) Ospfv3V6RouteRange {

	obj.routeOriginHolder = nil
	obj.obj.RouteOrigin = value.msg()

	return obj
}

func (obj *ospfv3V6RouteRange) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface Ospfv3V6RouteRange")
	}

	if len(obj.obj.Addresses) != 0 {

		if set_default {
			obj.Addresses().clearHolderSlice()
			for _, item := range obj.obj.Addresses {
				obj.Addresses().appendHolderSlice(&v6RouteAddress{obj: item})
			}
		}
		for _, item := range obj.Addresses().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Metric != nil {

		if *obj.obj.Metric > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv3V6RouteRange.Metric <= 16777215 but Got %d", *obj.obj.Metric))
		}

	}

	if obj.obj.RouteOrigin != nil {

		obj.RouteOrigin().validateObj(vObj, set_default)
	}

}

func (obj *ospfv3V6RouteRange) setDefault() {
	if obj.obj.Metric == nil {
		obj.SetMetric(0)
	}

}

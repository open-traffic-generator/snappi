package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2V4RouteRange *****
type ospfv2V4RouteRange struct {
	validation
	obj             *otg.Ospfv2V4RouteRange
	marshaller      marshalOspfv2V4RouteRange
	unMarshaller    unMarshalOspfv2V4RouteRange
	addressesHolder Ospfv2V4RouteRangeV4RouteAddressIter
}

func NewOspfv2V4RouteRange() Ospfv2V4RouteRange {
	obj := ospfv2V4RouteRange{obj: &otg.Ospfv2V4RouteRange{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2V4RouteRange) msg() *otg.Ospfv2V4RouteRange {
	return obj.obj
}

func (obj *ospfv2V4RouteRange) setMsg(msg *otg.Ospfv2V4RouteRange) Ospfv2V4RouteRange {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2V4RouteRange struct {
	obj *ospfv2V4RouteRange
}

type marshalOspfv2V4RouteRange interface {
	// ToProto marshals Ospfv2V4RouteRange to protobuf object *otg.Ospfv2V4RouteRange
	ToProto() (*otg.Ospfv2V4RouteRange, error)
	// ToPbText marshals Ospfv2V4RouteRange to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2V4RouteRange to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2V4RouteRange to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2V4RouteRange struct {
	obj *ospfv2V4RouteRange
}

type unMarshalOspfv2V4RouteRange interface {
	// FromProto unmarshals Ospfv2V4RouteRange from protobuf object *otg.Ospfv2V4RouteRange
	FromProto(msg *otg.Ospfv2V4RouteRange) (Ospfv2V4RouteRange, error)
	// FromPbText unmarshals Ospfv2V4RouteRange from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2V4RouteRange from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2V4RouteRange from JSON text
	FromJson(value string) error
}

func (obj *ospfv2V4RouteRange) Marshal() marshalOspfv2V4RouteRange {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2V4RouteRange{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2V4RouteRange) Unmarshal() unMarshalOspfv2V4RouteRange {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2V4RouteRange{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2V4RouteRange) ToProto() (*otg.Ospfv2V4RouteRange, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2V4RouteRange) FromProto(msg *otg.Ospfv2V4RouteRange) (Ospfv2V4RouteRange, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2V4RouteRange) ToPbText() (string, error) {
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

func (m *unMarshalospfv2V4RouteRange) FromPbText(value string) error {
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

func (m *marshalospfv2V4RouteRange) ToYaml() (string, error) {
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

func (m *unMarshalospfv2V4RouteRange) FromYaml(value string) error {
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

func (m *marshalospfv2V4RouteRange) ToJson() (string, error) {
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

func (m *unMarshalospfv2V4RouteRange) FromJson(value string) error {
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

func (obj *ospfv2V4RouteRange) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2V4RouteRange) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2V4RouteRange) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2V4RouteRange) Clone() (Ospfv2V4RouteRange, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2V4RouteRange()
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

func (obj *ospfv2V4RouteRange) setNil() {
	obj.addressesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2V4RouteRange is emulated OSPFv2 IPv4 routes.
type Ospfv2V4RouteRange interface {
	Validation
	// msg marshals Ospfv2V4RouteRange to protobuf object *otg.Ospfv2V4RouteRange
	// and doesn't set defaults
	msg() *otg.Ospfv2V4RouteRange
	// setMsg unmarshals Ospfv2V4RouteRange from protobuf object *otg.Ospfv2V4RouteRange
	// and doesn't set defaults
	setMsg(*otg.Ospfv2V4RouteRange) Ospfv2V4RouteRange
	// provides marshal interface
	Marshal() marshalOspfv2V4RouteRange
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2V4RouteRange
	// validate validates Ospfv2V4RouteRange
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2V4RouteRange, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in Ospfv2V4RouteRange.
	Name() string
	// SetName assigns string provided by user to Ospfv2V4RouteRange
	SetName(value string) Ospfv2V4RouteRange
	// Addresses returns Ospfv2V4RouteRangeV4RouteAddressIterIter, set in Ospfv2V4RouteRange
	Addresses() Ospfv2V4RouteRangeV4RouteAddressIter
	// RouterMetric returns uint32, set in Ospfv2V4RouteRange.
	RouterMetric() uint32
	// SetRouterMetric assigns uint32 provided by user to Ospfv2V4RouteRange
	SetRouterMetric(value uint32) Ospfv2V4RouteRange
	// HasRouterMetric checks if RouterMetric has been set in Ospfv2V4RouteRange
	HasRouterMetric() bool
	// Choice returns Ospfv2V4RouteRangeChoiceEnum, set in Ospfv2V4RouteRange
	Choice() Ospfv2V4RouteRangeChoiceEnum
	// setChoice assigns Ospfv2V4RouteRangeChoiceEnum provided by user to Ospfv2V4RouteRange
	setChoice(value Ospfv2V4RouteRangeChoiceEnum) Ospfv2V4RouteRange
	// HasChoice checks if Choice has been set in Ospfv2V4RouteRange
	HasChoice() bool
	// getter for ExternalType_1 to set choice.
	ExternalType_1()
	// getter for InterArea to set choice.
	InterArea()
	// getter for IntraArea to set choice.
	IntraArea()
	// getter for NssaExternal to set choice.
	NssaExternal()
	// getter for ExternalType_2 to set choice.
	ExternalType_2()
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *ospfv2V4RouteRange) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the Ospfv2V4RouteRange object
func (obj *ospfv2V4RouteRange) SetName(value string) Ospfv2V4RouteRange {

	obj.obj.Name = &value
	return obj
}

// A list of group of IPv4 route addresses.
// Addresses returns a []V4RouteAddress
func (obj *ospfv2V4RouteRange) Addresses() Ospfv2V4RouteRangeV4RouteAddressIter {
	if len(obj.obj.Addresses) == 0 {
		obj.obj.Addresses = []*otg.V4RouteAddress{}
	}
	if obj.addressesHolder == nil {
		obj.addressesHolder = newOspfv2V4RouteRangeV4RouteAddressIter(&obj.obj.Addresses).setMsg(obj)
	}
	return obj.addressesHolder
}

type ospfv2V4RouteRangeV4RouteAddressIter struct {
	obj                 *ospfv2V4RouteRange
	v4RouteAddressSlice []V4RouteAddress
	fieldPtr            *[]*otg.V4RouteAddress
}

func newOspfv2V4RouteRangeV4RouteAddressIter(ptr *[]*otg.V4RouteAddress) Ospfv2V4RouteRangeV4RouteAddressIter {
	return &ospfv2V4RouteRangeV4RouteAddressIter{fieldPtr: ptr}
}

type Ospfv2V4RouteRangeV4RouteAddressIter interface {
	setMsg(*ospfv2V4RouteRange) Ospfv2V4RouteRangeV4RouteAddressIter
	Items() []V4RouteAddress
	Add() V4RouteAddress
	Append(items ...V4RouteAddress) Ospfv2V4RouteRangeV4RouteAddressIter
	Set(index int, newObj V4RouteAddress) Ospfv2V4RouteRangeV4RouteAddressIter
	Clear() Ospfv2V4RouteRangeV4RouteAddressIter
	clearHolderSlice() Ospfv2V4RouteRangeV4RouteAddressIter
	appendHolderSlice(item V4RouteAddress) Ospfv2V4RouteRangeV4RouteAddressIter
}

func (obj *ospfv2V4RouteRangeV4RouteAddressIter) setMsg(msg *ospfv2V4RouteRange) Ospfv2V4RouteRangeV4RouteAddressIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&v4RouteAddress{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2V4RouteRangeV4RouteAddressIter) Items() []V4RouteAddress {
	return obj.v4RouteAddressSlice
}

func (obj *ospfv2V4RouteRangeV4RouteAddressIter) Add() V4RouteAddress {
	newObj := &otg.V4RouteAddress{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &v4RouteAddress{obj: newObj}
	newLibObj.setDefault()
	obj.v4RouteAddressSlice = append(obj.v4RouteAddressSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2V4RouteRangeV4RouteAddressIter) Append(items ...V4RouteAddress) Ospfv2V4RouteRangeV4RouteAddressIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.v4RouteAddressSlice = append(obj.v4RouteAddressSlice, item)
	}
	return obj
}

func (obj *ospfv2V4RouteRangeV4RouteAddressIter) Set(index int, newObj V4RouteAddress) Ospfv2V4RouteRangeV4RouteAddressIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.v4RouteAddressSlice[index] = newObj
	return obj
}
func (obj *ospfv2V4RouteRangeV4RouteAddressIter) Clear() Ospfv2V4RouteRangeV4RouteAddressIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.V4RouteAddress{}
		obj.v4RouteAddressSlice = []V4RouteAddress{}
	}
	return obj
}
func (obj *ospfv2V4RouteRangeV4RouteAddressIter) clearHolderSlice() Ospfv2V4RouteRangeV4RouteAddressIter {
	if len(obj.v4RouteAddressSlice) > 0 {
		obj.v4RouteAddressSlice = []V4RouteAddress{}
	}
	return obj
}
func (obj *ospfv2V4RouteRangeV4RouteAddressIter) appendHolderSlice(item V4RouteAddress) Ospfv2V4RouteRangeV4RouteAddressIter {
	obj.v4RouteAddressSlice = append(obj.v4RouteAddressSlice, item)
	return obj
}

// The user-defined metric associated with this route range.
// RouterMetric returns a uint32
func (obj *ospfv2V4RouteRange) RouterMetric() uint32 {

	return *obj.obj.RouterMetric

}

// The user-defined metric associated with this route range.
// RouterMetric returns a uint32
func (obj *ospfv2V4RouteRange) HasRouterMetric() bool {
	return obj.obj.RouterMetric != nil
}

// The user-defined metric associated with this route range.
// SetRouterMetric sets the uint32 value in the Ospfv2V4RouteRange object
func (obj *ospfv2V4RouteRange) SetRouterMetric(value uint32) Ospfv2V4RouteRange {

	obj.obj.RouterMetric = &value
	return obj
}

type Ospfv2V4RouteRangeChoiceEnum string

// Enum of Choice on Ospfv2V4RouteRange
var Ospfv2V4RouteRangeChoice = struct {
	INTRA_AREA      Ospfv2V4RouteRangeChoiceEnum
	INTER_AREA      Ospfv2V4RouteRangeChoiceEnum
	EXTERNAL_TYPE_1 Ospfv2V4RouteRangeChoiceEnum
	EXTERNAL_TYPE_2 Ospfv2V4RouteRangeChoiceEnum
	NSSA_EXTERNAL   Ospfv2V4RouteRangeChoiceEnum
}{
	INTRA_AREA:      Ospfv2V4RouteRangeChoiceEnum("intra_area"),
	INTER_AREA:      Ospfv2V4RouteRangeChoiceEnum("inter_area"),
	EXTERNAL_TYPE_1: Ospfv2V4RouteRangeChoiceEnum("external_type_1"),
	EXTERNAL_TYPE_2: Ospfv2V4RouteRangeChoiceEnum("external_type_2"),
	NSSA_EXTERNAL:   Ospfv2V4RouteRangeChoiceEnum("nssa_external"),
}

func (obj *ospfv2V4RouteRange) Choice() Ospfv2V4RouteRangeChoiceEnum {
	return Ospfv2V4RouteRangeChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for ExternalType_1 to set choice
func (obj *ospfv2V4RouteRange) ExternalType_1() {
	obj.setChoice(Ospfv2V4RouteRangeChoice.EXTERNAL_TYPE_1)
}

// getter for InterArea to set choice
func (obj *ospfv2V4RouteRange) InterArea() {
	obj.setChoice(Ospfv2V4RouteRangeChoice.INTER_AREA)
}

// getter for IntraArea to set choice
func (obj *ospfv2V4RouteRange) IntraArea() {
	obj.setChoice(Ospfv2V4RouteRangeChoice.INTRA_AREA)
}

// getter for NssaExternal to set choice
func (obj *ospfv2V4RouteRange) NssaExternal() {
	obj.setChoice(Ospfv2V4RouteRangeChoice.NSSA_EXTERNAL)
}

// getter for ExternalType_2 to set choice
func (obj *ospfv2V4RouteRange) ExternalType_2() {
	obj.setChoice(Ospfv2V4RouteRangeChoice.EXTERNAL_TYPE_2)
}

// The type of the OSPFv2 routes. route types correspond directly to the OSPFv2 LSAs types as
// defined in the "OSPFv2 Link State (LS) Type - <http://www.iana.org/assignments/ospfv2-parameters.
// Supported types are:
// - intra_area: for 1 Intra-Area.
// - inter_area: for 3 Inter Area.
// - external_type_1: for 5 - Autonomous System (AS) External with internal AS meteric.
// - external_type_2: for 5 - Autonomous System (AS) External with internal and external AS meteric.
// - nssa_external: for 7 - Not-So-Stubby Area (NSSA) External.
// Choice returns a string
func (obj *ospfv2V4RouteRange) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ospfv2V4RouteRange) setChoice(value Ospfv2V4RouteRangeChoiceEnum) Ospfv2V4RouteRange {
	intValue, ok := otg.Ospfv2V4RouteRange_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2V4RouteRangeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2V4RouteRange_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *ospfv2V4RouteRange) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface Ospfv2V4RouteRange")
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

	if obj.obj.RouterMetric != nil {

		if *obj.obj.RouterMetric > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv2V4RouteRange.RouterMetric <= 16777215 but Got %d", *obj.obj.RouterMetric))
		}

	}

}

func (obj *ospfv2V4RouteRange) setDefault() {
	var choices_set int = 0
	var choice Ospfv2V4RouteRangeChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Ospfv2V4RouteRangeChoice.INTER_AREA)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Ospfv2V4RouteRange")
			}
		} else {
			intVal := otg.Ospfv2V4RouteRange_Choice_Enum_value[string(choice)]
			enumValue := otg.Ospfv2V4RouteRange_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.RouterMetric == nil {
		obj.SetRouterMetric(0)
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2InterfaceNetworkType *****
type ospfv2InterfaceNetworkType struct {
	validation
	obj                     *otg.Ospfv2InterfaceNetworkType
	marshaller              marshalOspfv2InterfaceNetworkType
	unMarshaller            unMarshalOspfv2InterfaceNetworkType
	pointToMultipointHolder Ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter
}

func NewOspfv2InterfaceNetworkType() Ospfv2InterfaceNetworkType {
	obj := ospfv2InterfaceNetworkType{obj: &otg.Ospfv2InterfaceNetworkType{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2InterfaceNetworkType) msg() *otg.Ospfv2InterfaceNetworkType {
	return obj.obj
}

func (obj *ospfv2InterfaceNetworkType) setMsg(msg *otg.Ospfv2InterfaceNetworkType) Ospfv2InterfaceNetworkType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2InterfaceNetworkType struct {
	obj *ospfv2InterfaceNetworkType
}

type marshalOspfv2InterfaceNetworkType interface {
	// ToProto marshals Ospfv2InterfaceNetworkType to protobuf object *otg.Ospfv2InterfaceNetworkType
	ToProto() (*otg.Ospfv2InterfaceNetworkType, error)
	// ToPbText marshals Ospfv2InterfaceNetworkType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2InterfaceNetworkType to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2InterfaceNetworkType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv2InterfaceNetworkType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv2InterfaceNetworkType struct {
	obj *ospfv2InterfaceNetworkType
}

type unMarshalOspfv2InterfaceNetworkType interface {
	// FromProto unmarshals Ospfv2InterfaceNetworkType from protobuf object *otg.Ospfv2InterfaceNetworkType
	FromProto(msg *otg.Ospfv2InterfaceNetworkType) (Ospfv2InterfaceNetworkType, error)
	// FromPbText unmarshals Ospfv2InterfaceNetworkType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2InterfaceNetworkType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2InterfaceNetworkType from JSON text
	FromJson(value string) error
}

func (obj *ospfv2InterfaceNetworkType) Marshal() marshalOspfv2InterfaceNetworkType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2InterfaceNetworkType{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2InterfaceNetworkType) Unmarshal() unMarshalOspfv2InterfaceNetworkType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2InterfaceNetworkType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2InterfaceNetworkType) ToProto() (*otg.Ospfv2InterfaceNetworkType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2InterfaceNetworkType) FromProto(msg *otg.Ospfv2InterfaceNetworkType) (Ospfv2InterfaceNetworkType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2InterfaceNetworkType) ToPbText() (string, error) {
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

func (m *unMarshalospfv2InterfaceNetworkType) FromPbText(value string) error {
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

func (m *marshalospfv2InterfaceNetworkType) ToYaml() (string, error) {
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

func (m *unMarshalospfv2InterfaceNetworkType) FromYaml(value string) error {
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

func (m *marshalospfv2InterfaceNetworkType) ToJsonRaw() (string, error) {
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

func (m *marshalospfv2InterfaceNetworkType) ToJson() (string, error) {
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

func (m *unMarshalospfv2InterfaceNetworkType) FromJson(value string) error {
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

func (obj *ospfv2InterfaceNetworkType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2InterfaceNetworkType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2InterfaceNetworkType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2InterfaceNetworkType) Clone() (Ospfv2InterfaceNetworkType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2InterfaceNetworkType()
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

func (obj *ospfv2InterfaceNetworkType) setNil() {
	obj.pointToMultipointHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2InterfaceNetworkType is the OSPF network link type options.
// - Point to Point:
// - Broadcast:
// - Point to Multipoint: In this case, at least a neigbor to be configured.
type Ospfv2InterfaceNetworkType interface {
	Validation
	// msg marshals Ospfv2InterfaceNetworkType to protobuf object *otg.Ospfv2InterfaceNetworkType
	// and doesn't set defaults
	msg() *otg.Ospfv2InterfaceNetworkType
	// setMsg unmarshals Ospfv2InterfaceNetworkType from protobuf object *otg.Ospfv2InterfaceNetworkType
	// and doesn't set defaults
	setMsg(*otg.Ospfv2InterfaceNetworkType) Ospfv2InterfaceNetworkType
	// provides marshal interface
	Marshal() marshalOspfv2InterfaceNetworkType
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2InterfaceNetworkType
	// validate validates Ospfv2InterfaceNetworkType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2InterfaceNetworkType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Ospfv2InterfaceNetworkTypeChoiceEnum, set in Ospfv2InterfaceNetworkType
	Choice() Ospfv2InterfaceNetworkTypeChoiceEnum
	// setChoice assigns Ospfv2InterfaceNetworkTypeChoiceEnum provided by user to Ospfv2InterfaceNetworkType
	setChoice(value Ospfv2InterfaceNetworkTypeChoiceEnum) Ospfv2InterfaceNetworkType
	// HasChoice checks if Choice has been set in Ospfv2InterfaceNetworkType
	HasChoice() bool
	// getter for PointToPoint to set choice.
	PointToPoint()
	// getter for Broadcast to set choice.
	Broadcast()
	// PointToMultipoint returns Ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIterIter, set in Ospfv2InterfaceNetworkType
	PointToMultipoint() Ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter
	setNil()
}

type Ospfv2InterfaceNetworkTypeChoiceEnum string

// Enum of Choice on Ospfv2InterfaceNetworkType
var Ospfv2InterfaceNetworkTypeChoice = struct {
	BROADCAST           Ospfv2InterfaceNetworkTypeChoiceEnum
	POINT_TO_POINT      Ospfv2InterfaceNetworkTypeChoiceEnum
	POINT_TO_MULTIPOINT Ospfv2InterfaceNetworkTypeChoiceEnum
}{
	BROADCAST:           Ospfv2InterfaceNetworkTypeChoiceEnum("broadcast"),
	POINT_TO_POINT:      Ospfv2InterfaceNetworkTypeChoiceEnum("point_to_point"),
	POINT_TO_MULTIPOINT: Ospfv2InterfaceNetworkTypeChoiceEnum("point_to_multipoint"),
}

func (obj *ospfv2InterfaceNetworkType) Choice() Ospfv2InterfaceNetworkTypeChoiceEnum {
	return Ospfv2InterfaceNetworkTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for PointToPoint to set choice
func (obj *ospfv2InterfaceNetworkType) PointToPoint() {
	obj.setChoice(Ospfv2InterfaceNetworkTypeChoice.POINT_TO_POINT)
}

// getter for Broadcast to set choice
func (obj *ospfv2InterfaceNetworkType) Broadcast() {
	obj.setChoice(Ospfv2InterfaceNetworkTypeChoice.BROADCAST)
}

// description is TBD
// Choice returns a string
func (obj *ospfv2InterfaceNetworkType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ospfv2InterfaceNetworkType) setChoice(value Ospfv2InterfaceNetworkTypeChoiceEnum) Ospfv2InterfaceNetworkType {
	intValue, ok := otg.Ospfv2InterfaceNetworkType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2InterfaceNetworkTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2InterfaceNetworkType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.PointToMultipoint = nil
	obj.pointToMultipointHolder = nil

	if value == Ospfv2InterfaceNetworkTypeChoice.POINT_TO_MULTIPOINT {
		obj.obj.PointToMultipoint = []*otg.Ospfv2InterfaceNeighbor{}
	}

	return obj
}

// List of Neigbhors.
// PointToMultipoint returns a []Ospfv2InterfaceNeighbor
func (obj *ospfv2InterfaceNetworkType) PointToMultipoint() Ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter {
	if len(obj.obj.PointToMultipoint) == 0 {
		obj.setChoice(Ospfv2InterfaceNetworkTypeChoice.POINT_TO_MULTIPOINT)
	}
	if obj.pointToMultipointHolder == nil {
		obj.pointToMultipointHolder = newOspfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter(&obj.obj.PointToMultipoint).setMsg(obj)
	}
	return obj.pointToMultipointHolder
}

type ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter struct {
	obj                          *ospfv2InterfaceNetworkType
	ospfv2InterfaceNeighborSlice []Ospfv2InterfaceNeighbor
	fieldPtr                     *[]*otg.Ospfv2InterfaceNeighbor
}

func newOspfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter(ptr *[]*otg.Ospfv2InterfaceNeighbor) Ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter {
	return &ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter{fieldPtr: ptr}
}

type Ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter interface {
	setMsg(*ospfv2InterfaceNetworkType) Ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter
	Items() []Ospfv2InterfaceNeighbor
	Add() Ospfv2InterfaceNeighbor
	Append(items ...Ospfv2InterfaceNeighbor) Ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter
	Set(index int, newObj Ospfv2InterfaceNeighbor) Ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter
	Clear() Ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter
	clearHolderSlice() Ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter
	appendHolderSlice(item Ospfv2InterfaceNeighbor) Ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter
}

func (obj *ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter) setMsg(msg *ospfv2InterfaceNetworkType) Ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2InterfaceNeighbor{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter) Items() []Ospfv2InterfaceNeighbor {
	return obj.ospfv2InterfaceNeighborSlice
}

func (obj *ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter) Add() Ospfv2InterfaceNeighbor {
	newObj := &otg.Ospfv2InterfaceNeighbor{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2InterfaceNeighbor{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2InterfaceNeighborSlice = append(obj.ospfv2InterfaceNeighborSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter) Append(items ...Ospfv2InterfaceNeighbor) Ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2InterfaceNeighborSlice = append(obj.ospfv2InterfaceNeighborSlice, item)
	}
	return obj
}

func (obj *ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter) Set(index int, newObj Ospfv2InterfaceNeighbor) Ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2InterfaceNeighborSlice[index] = newObj
	return obj
}
func (obj *ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter) Clear() Ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2InterfaceNeighbor{}
		obj.ospfv2InterfaceNeighborSlice = []Ospfv2InterfaceNeighbor{}
	}
	return obj
}
func (obj *ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter) clearHolderSlice() Ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter {
	if len(obj.ospfv2InterfaceNeighborSlice) > 0 {
		obj.ospfv2InterfaceNeighborSlice = []Ospfv2InterfaceNeighbor{}
	}
	return obj
}
func (obj *ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter) appendHolderSlice(item Ospfv2InterfaceNeighbor) Ospfv2InterfaceNetworkTypeOspfv2InterfaceNeighborIter {
	obj.ospfv2InterfaceNeighborSlice = append(obj.ospfv2InterfaceNeighborSlice, item)
	return obj
}

func (obj *ospfv2InterfaceNetworkType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.PointToMultipoint) != 0 {

		if set_default {
			obj.PointToMultipoint().clearHolderSlice()
			for _, item := range obj.obj.PointToMultipoint {
				obj.PointToMultipoint().appendHolderSlice(&ospfv2InterfaceNeighbor{obj: item})
			}
		}
		for _, item := range obj.PointToMultipoint().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *ospfv2InterfaceNetworkType) setDefault() {
	var choices_set int = 0
	var choice Ospfv2InterfaceNetworkTypeChoiceEnum

	if len(obj.obj.PointToMultipoint) > 0 {
		choices_set += 1
		choice = Ospfv2InterfaceNetworkTypeChoice.POINT_TO_MULTIPOINT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Ospfv2InterfaceNetworkTypeChoice.BROADCAST)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Ospfv2InterfaceNetworkType")
			}
		} else {
			intVal := otg.Ospfv2InterfaceNetworkType_Choice_Enum_value[string(choice)]
			enumValue := otg.Ospfv2InterfaceNetworkType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

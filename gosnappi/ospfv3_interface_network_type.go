package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3InterfaceNetworkType *****
type ospfv3InterfaceNetworkType struct {
	validation
	obj             *otg.Ospfv3InterfaceNetworkType
	marshaller      marshalOspfv3InterfaceNetworkType
	unMarshaller    unMarshalOspfv3InterfaceNetworkType
	broadcastHolder Ospfv3InterfaceBroadcast
}

func NewOspfv3InterfaceNetworkType() Ospfv3InterfaceNetworkType {
	obj := ospfv3InterfaceNetworkType{obj: &otg.Ospfv3InterfaceNetworkType{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3InterfaceNetworkType) msg() *otg.Ospfv3InterfaceNetworkType {
	return obj.obj
}

func (obj *ospfv3InterfaceNetworkType) setMsg(msg *otg.Ospfv3InterfaceNetworkType) Ospfv3InterfaceNetworkType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3InterfaceNetworkType struct {
	obj *ospfv3InterfaceNetworkType
}

type marshalOspfv3InterfaceNetworkType interface {
	// ToProto marshals Ospfv3InterfaceNetworkType to protobuf object *otg.Ospfv3InterfaceNetworkType
	ToProto() (*otg.Ospfv3InterfaceNetworkType, error)
	// ToPbText marshals Ospfv3InterfaceNetworkType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3InterfaceNetworkType to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3InterfaceNetworkType to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3InterfaceNetworkType struct {
	obj *ospfv3InterfaceNetworkType
}

type unMarshalOspfv3InterfaceNetworkType interface {
	// FromProto unmarshals Ospfv3InterfaceNetworkType from protobuf object *otg.Ospfv3InterfaceNetworkType
	FromProto(msg *otg.Ospfv3InterfaceNetworkType) (Ospfv3InterfaceNetworkType, error)
	// FromPbText unmarshals Ospfv3InterfaceNetworkType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3InterfaceNetworkType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3InterfaceNetworkType from JSON text
	FromJson(value string) error
}

func (obj *ospfv3InterfaceNetworkType) Marshal() marshalOspfv3InterfaceNetworkType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3InterfaceNetworkType{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3InterfaceNetworkType) Unmarshal() unMarshalOspfv3InterfaceNetworkType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3InterfaceNetworkType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3InterfaceNetworkType) ToProto() (*otg.Ospfv3InterfaceNetworkType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3InterfaceNetworkType) FromProto(msg *otg.Ospfv3InterfaceNetworkType) (Ospfv3InterfaceNetworkType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3InterfaceNetworkType) ToPbText() (string, error) {
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

func (m *unMarshalospfv3InterfaceNetworkType) FromPbText(value string) error {
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

func (m *marshalospfv3InterfaceNetworkType) ToYaml() (string, error) {
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

func (m *unMarshalospfv3InterfaceNetworkType) FromYaml(value string) error {
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

func (m *marshalospfv3InterfaceNetworkType) ToJson() (string, error) {
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

func (m *unMarshalospfv3InterfaceNetworkType) FromJson(value string) error {
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

func (obj *ospfv3InterfaceNetworkType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3InterfaceNetworkType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3InterfaceNetworkType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3InterfaceNetworkType) Clone() (Ospfv3InterfaceNetworkType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3InterfaceNetworkType()
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

func (obj *ospfv3InterfaceNetworkType) setNil() {
	obj.broadcastHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv3InterfaceNetworkType is the OSPFv3 network link type options.
// - Broadcast
// - Point to Point
type Ospfv3InterfaceNetworkType interface {
	Validation
	// msg marshals Ospfv3InterfaceNetworkType to protobuf object *otg.Ospfv3InterfaceNetworkType
	// and doesn't set defaults
	msg() *otg.Ospfv3InterfaceNetworkType
	// setMsg unmarshals Ospfv3InterfaceNetworkType from protobuf object *otg.Ospfv3InterfaceNetworkType
	// and doesn't set defaults
	setMsg(*otg.Ospfv3InterfaceNetworkType) Ospfv3InterfaceNetworkType
	// provides marshal interface
	Marshal() marshalOspfv3InterfaceNetworkType
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3InterfaceNetworkType
	// validate validates Ospfv3InterfaceNetworkType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3InterfaceNetworkType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Ospfv3InterfaceNetworkTypeChoiceEnum, set in Ospfv3InterfaceNetworkType
	Choice() Ospfv3InterfaceNetworkTypeChoiceEnum
	// setChoice assigns Ospfv3InterfaceNetworkTypeChoiceEnum provided by user to Ospfv3InterfaceNetworkType
	setChoice(value Ospfv3InterfaceNetworkTypeChoiceEnum) Ospfv3InterfaceNetworkType
	// HasChoice checks if Choice has been set in Ospfv3InterfaceNetworkType
	HasChoice() bool
	// getter for PointToPoint to set choice.
	PointToPoint()
	// Broadcast returns Ospfv3InterfaceBroadcast, set in Ospfv3InterfaceNetworkType.
	// Ospfv3InterfaceBroadcast is container for capabilities associated with network type broadcast.
	Broadcast() Ospfv3InterfaceBroadcast
	// SetBroadcast assigns Ospfv3InterfaceBroadcast provided by user to Ospfv3InterfaceNetworkType.
	// Ospfv3InterfaceBroadcast is container for capabilities associated with network type broadcast.
	SetBroadcast(value Ospfv3InterfaceBroadcast) Ospfv3InterfaceNetworkType
	// HasBroadcast checks if Broadcast has been set in Ospfv3InterfaceNetworkType
	HasBroadcast() bool
	setNil()
}

type Ospfv3InterfaceNetworkTypeChoiceEnum string

// Enum of Choice on Ospfv3InterfaceNetworkType
var Ospfv3InterfaceNetworkTypeChoice = struct {
	BROADCAST      Ospfv3InterfaceNetworkTypeChoiceEnum
	POINT_TO_POINT Ospfv3InterfaceNetworkTypeChoiceEnum
}{
	BROADCAST:      Ospfv3InterfaceNetworkTypeChoiceEnum("broadcast"),
	POINT_TO_POINT: Ospfv3InterfaceNetworkTypeChoiceEnum("point_to_point"),
}

func (obj *ospfv3InterfaceNetworkType) Choice() Ospfv3InterfaceNetworkTypeChoiceEnum {
	return Ospfv3InterfaceNetworkTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for PointToPoint to set choice
func (obj *ospfv3InterfaceNetworkType) PointToPoint() {
	obj.setChoice(Ospfv3InterfaceNetworkTypeChoice.POINT_TO_POINT)
}

// description is TBD
// Choice returns a string
func (obj *ospfv3InterfaceNetworkType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ospfv3InterfaceNetworkType) setChoice(value Ospfv3InterfaceNetworkTypeChoiceEnum) Ospfv3InterfaceNetworkType {
	intValue, ok := otg.Ospfv3InterfaceNetworkType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv3InterfaceNetworkTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv3InterfaceNetworkType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Broadcast = nil
	obj.broadcastHolder = nil

	if value == Ospfv3InterfaceNetworkTypeChoice.BROADCAST {
		obj.obj.Broadcast = NewOspfv3InterfaceBroadcast().msg()
	}

	return obj
}

// Capabilities associated with network type broadcast.
// Broadcast returns a Ospfv3InterfaceBroadcast
func (obj *ospfv3InterfaceNetworkType) Broadcast() Ospfv3InterfaceBroadcast {
	if obj.obj.Broadcast == nil {
		obj.setChoice(Ospfv3InterfaceNetworkTypeChoice.BROADCAST)
	}
	if obj.broadcastHolder == nil {
		obj.broadcastHolder = &ospfv3InterfaceBroadcast{obj: obj.obj.Broadcast}
	}
	return obj.broadcastHolder
}

// Capabilities associated with network type broadcast.
// Broadcast returns a Ospfv3InterfaceBroadcast
func (obj *ospfv3InterfaceNetworkType) HasBroadcast() bool {
	return obj.obj.Broadcast != nil
}

// Capabilities associated with network type broadcast.
// SetBroadcast sets the Ospfv3InterfaceBroadcast value in the Ospfv3InterfaceNetworkType object
func (obj *ospfv3InterfaceNetworkType) SetBroadcast(value Ospfv3InterfaceBroadcast) Ospfv3InterfaceNetworkType {
	obj.setChoice(Ospfv3InterfaceNetworkTypeChoice.BROADCAST)
	obj.broadcastHolder = nil
	obj.obj.Broadcast = value.msg()

	return obj
}

func (obj *ospfv3InterfaceNetworkType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Broadcast != nil {

		obj.Broadcast().validateObj(vObj, set_default)
	}

}

func (obj *ospfv3InterfaceNetworkType) setDefault() {
	var choices_set int = 0
	var choice Ospfv3InterfaceNetworkTypeChoiceEnum

	if obj.obj.Broadcast != nil {
		choices_set += 1
		choice = Ospfv3InterfaceNetworkTypeChoice.BROADCAST
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Ospfv3InterfaceNetworkTypeChoice.BROADCAST)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Ospfv3InterfaceNetworkType")
			}
		} else {
			intVal := otg.Ospfv3InterfaceNetworkType_Choice_Enum_value[string(choice)]
			enumValue := otg.Ospfv3InterfaceNetworkType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

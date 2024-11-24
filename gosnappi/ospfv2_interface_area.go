package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2InterfaceArea *****
type ospfv2InterfaceArea struct {
	validation
	obj          *otg.Ospfv2InterfaceArea
	marshaller   marshalOspfv2InterfaceArea
	unMarshaller unMarshalOspfv2InterfaceArea
}

func NewOspfv2InterfaceArea() Ospfv2InterfaceArea {
	obj := ospfv2InterfaceArea{obj: &otg.Ospfv2InterfaceArea{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2InterfaceArea) msg() *otg.Ospfv2InterfaceArea {
	return obj.obj
}

func (obj *ospfv2InterfaceArea) setMsg(msg *otg.Ospfv2InterfaceArea) Ospfv2InterfaceArea {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2InterfaceArea struct {
	obj *ospfv2InterfaceArea
}

type marshalOspfv2InterfaceArea interface {
	// ToProto marshals Ospfv2InterfaceArea to protobuf object *otg.Ospfv2InterfaceArea
	ToProto() (*otg.Ospfv2InterfaceArea, error)
	// ToPbText marshals Ospfv2InterfaceArea to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2InterfaceArea to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2InterfaceArea to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv2InterfaceArea to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv2InterfaceArea struct {
	obj *ospfv2InterfaceArea
}

type unMarshalOspfv2InterfaceArea interface {
	// FromProto unmarshals Ospfv2InterfaceArea from protobuf object *otg.Ospfv2InterfaceArea
	FromProto(msg *otg.Ospfv2InterfaceArea) (Ospfv2InterfaceArea, error)
	// FromPbText unmarshals Ospfv2InterfaceArea from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2InterfaceArea from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2InterfaceArea from JSON text
	FromJson(value string) error
}

func (obj *ospfv2InterfaceArea) Marshal() marshalOspfv2InterfaceArea {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2InterfaceArea{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2InterfaceArea) Unmarshal() unMarshalOspfv2InterfaceArea {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2InterfaceArea{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2InterfaceArea) ToProto() (*otg.Ospfv2InterfaceArea, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2InterfaceArea) FromProto(msg *otg.Ospfv2InterfaceArea) (Ospfv2InterfaceArea, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2InterfaceArea) ToPbText() (string, error) {
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

func (m *unMarshalospfv2InterfaceArea) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalospfv2InterfaceArea) ToYaml() (string, error) {
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

func (m *unMarshalospfv2InterfaceArea) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalospfv2InterfaceArea) ToJsonRaw() (string, error) {
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

func (m *marshalospfv2InterfaceArea) ToJson() (string, error) {
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

func (m *unMarshalospfv2InterfaceArea) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *ospfv2InterfaceArea) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2InterfaceArea) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2InterfaceArea) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2InterfaceArea) Clone() (Ospfv2InterfaceArea, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2InterfaceArea()
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

// Ospfv2InterfaceArea is container for OSPF Area ID identifies the routing area to which the host belongs..
type Ospfv2InterfaceArea interface {
	Validation
	// msg marshals Ospfv2InterfaceArea to protobuf object *otg.Ospfv2InterfaceArea
	// and doesn't set defaults
	msg() *otg.Ospfv2InterfaceArea
	// setMsg unmarshals Ospfv2InterfaceArea from protobuf object *otg.Ospfv2InterfaceArea
	// and doesn't set defaults
	setMsg(*otg.Ospfv2InterfaceArea) Ospfv2InterfaceArea
	// provides marshal interface
	Marshal() marshalOspfv2InterfaceArea
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2InterfaceArea
	// validate validates Ospfv2InterfaceArea
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2InterfaceArea, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Ospfv2InterfaceAreaChoiceEnum, set in Ospfv2InterfaceArea
	Choice() Ospfv2InterfaceAreaChoiceEnum
	// setChoice assigns Ospfv2InterfaceAreaChoiceEnum provided by user to Ospfv2InterfaceArea
	setChoice(value Ospfv2InterfaceAreaChoiceEnum) Ospfv2InterfaceArea
	// HasChoice checks if Choice has been set in Ospfv2InterfaceArea
	HasChoice() bool
	// Id returns uint32, set in Ospfv2InterfaceArea.
	Id() uint32
	// SetId assigns uint32 provided by user to Ospfv2InterfaceArea
	SetId(value uint32) Ospfv2InterfaceArea
	// HasId checks if Id has been set in Ospfv2InterfaceArea
	HasId() bool
	// Ip returns string, set in Ospfv2InterfaceArea.
	Ip() string
	// SetIp assigns string provided by user to Ospfv2InterfaceArea
	SetIp(value string) Ospfv2InterfaceArea
	// HasIp checks if Ip has been set in Ospfv2InterfaceArea
	HasIp() bool
}

type Ospfv2InterfaceAreaChoiceEnum string

// Enum of Choice on Ospfv2InterfaceArea
var Ospfv2InterfaceAreaChoice = struct {
	ID Ospfv2InterfaceAreaChoiceEnum
	IP Ospfv2InterfaceAreaChoiceEnum
}{
	ID: Ospfv2InterfaceAreaChoiceEnum("id"),
	IP: Ospfv2InterfaceAreaChoiceEnum("ip"),
}

func (obj *ospfv2InterfaceArea) Choice() Ospfv2InterfaceAreaChoiceEnum {
	return Ospfv2InterfaceAreaChoiceEnum(obj.obj.Choice.Enum().String())
}

// The OSPF Area ID identifies the routing area to which the host belongs. Area ID type can be following format.
// - id: A 32-bit number identifying the area.
// - ip:     The Area ID in IPv4 address format.
// Choice returns a string
func (obj *ospfv2InterfaceArea) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ospfv2InterfaceArea) setChoice(value Ospfv2InterfaceAreaChoiceEnum) Ospfv2InterfaceArea {
	intValue, ok := otg.Ospfv2InterfaceArea_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2InterfaceAreaChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2InterfaceArea_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Ip = nil
	obj.obj.Id = nil

	if value == Ospfv2InterfaceAreaChoice.ID {
		defaultValue := uint32(0)
		obj.obj.Id = &defaultValue
	}

	return obj
}

// The Area ID.
// Id returns a uint32
func (obj *ospfv2InterfaceArea) Id() uint32 {

	if obj.obj.Id == nil {
		obj.setChoice(Ospfv2InterfaceAreaChoice.ID)
	}

	return *obj.obj.Id

}

// The Area ID.
// Id returns a uint32
func (obj *ospfv2InterfaceArea) HasId() bool {
	return obj.obj.Id != nil
}

// The Area ID.
// SetId sets the uint32 value in the Ospfv2InterfaceArea object
func (obj *ospfv2InterfaceArea) SetId(value uint32) Ospfv2InterfaceArea {
	obj.setChoice(Ospfv2InterfaceAreaChoice.ID)
	obj.obj.Id = &value
	return obj
}

// The Area ID in IPv4 address format.
// Ip returns a string
func (obj *ospfv2InterfaceArea) Ip() string {

	if obj.obj.Ip == nil {
		obj.setChoice(Ospfv2InterfaceAreaChoice.IP)
	}

	return *obj.obj.Ip

}

// The Area ID in IPv4 address format.
// Ip returns a string
func (obj *ospfv2InterfaceArea) HasIp() bool {
	return obj.obj.Ip != nil
}

// The Area ID in IPv4 address format.
// SetIp sets the string value in the Ospfv2InterfaceArea object
func (obj *ospfv2InterfaceArea) SetIp(value string) Ospfv2InterfaceArea {
	obj.setChoice(Ospfv2InterfaceAreaChoice.IP)
	obj.obj.Ip = &value
	return obj
}

func (obj *ospfv2InterfaceArea) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ip != nil {

		err := obj.validateIpv4(obj.Ip())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2InterfaceArea.Ip"))
		}

	}

}

func (obj *ospfv2InterfaceArea) setDefault() {
	var choices_set int = 0
	var choice Ospfv2InterfaceAreaChoiceEnum

	if obj.obj.Id != nil {
		choices_set += 1
		choice = Ospfv2InterfaceAreaChoice.ID
	}

	if obj.obj.Ip != nil {
		choices_set += 1
		choice = Ospfv2InterfaceAreaChoice.IP
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Ospfv2InterfaceAreaChoice.ID)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Ospfv2InterfaceArea")
			}
		} else {
			intVal := otg.Ospfv2InterfaceArea_Choice_Enum_value[string(choice)]
			enumValue := otg.Ospfv2InterfaceArea_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

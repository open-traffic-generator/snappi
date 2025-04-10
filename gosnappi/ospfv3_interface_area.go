package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3InterfaceArea *****
type ospfv3InterfaceArea struct {
	validation
	obj          *otg.Ospfv3InterfaceArea
	marshaller   marshalOspfv3InterfaceArea
	unMarshaller unMarshalOspfv3InterfaceArea
}

func NewOspfv3InterfaceArea() Ospfv3InterfaceArea {
	obj := ospfv3InterfaceArea{obj: &otg.Ospfv3InterfaceArea{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3InterfaceArea) msg() *otg.Ospfv3InterfaceArea {
	return obj.obj
}

func (obj *ospfv3InterfaceArea) setMsg(msg *otg.Ospfv3InterfaceArea) Ospfv3InterfaceArea {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3InterfaceArea struct {
	obj *ospfv3InterfaceArea
}

type marshalOspfv3InterfaceArea interface {
	// ToProto marshals Ospfv3InterfaceArea to protobuf object *otg.Ospfv3InterfaceArea
	ToProto() (*otg.Ospfv3InterfaceArea, error)
	// ToPbText marshals Ospfv3InterfaceArea to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3InterfaceArea to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3InterfaceArea to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3InterfaceArea struct {
	obj *ospfv3InterfaceArea
}

type unMarshalOspfv3InterfaceArea interface {
	// FromProto unmarshals Ospfv3InterfaceArea from protobuf object *otg.Ospfv3InterfaceArea
	FromProto(msg *otg.Ospfv3InterfaceArea) (Ospfv3InterfaceArea, error)
	// FromPbText unmarshals Ospfv3InterfaceArea from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3InterfaceArea from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3InterfaceArea from JSON text
	FromJson(value string) error
}

func (obj *ospfv3InterfaceArea) Marshal() marshalOspfv3InterfaceArea {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3InterfaceArea{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3InterfaceArea) Unmarshal() unMarshalOspfv3InterfaceArea {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3InterfaceArea{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3InterfaceArea) ToProto() (*otg.Ospfv3InterfaceArea, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3InterfaceArea) FromProto(msg *otg.Ospfv3InterfaceArea) (Ospfv3InterfaceArea, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3InterfaceArea) ToPbText() (string, error) {
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

func (m *unMarshalospfv3InterfaceArea) FromPbText(value string) error {
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

func (m *marshalospfv3InterfaceArea) ToYaml() (string, error) {
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

func (m *unMarshalospfv3InterfaceArea) FromYaml(value string) error {
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

func (m *marshalospfv3InterfaceArea) ToJson() (string, error) {
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

func (m *unMarshalospfv3InterfaceArea) FromJson(value string) error {
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

func (obj *ospfv3InterfaceArea) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3InterfaceArea) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3InterfaceArea) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3InterfaceArea) Clone() (Ospfv3InterfaceArea, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3InterfaceArea()
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

// Ospfv3InterfaceArea is container for OSPFv3 Area ID identifies the routing area to which the host belongs.
type Ospfv3InterfaceArea interface {
	Validation
	// msg marshals Ospfv3InterfaceArea to protobuf object *otg.Ospfv3InterfaceArea
	// and doesn't set defaults
	msg() *otg.Ospfv3InterfaceArea
	// setMsg unmarshals Ospfv3InterfaceArea from protobuf object *otg.Ospfv3InterfaceArea
	// and doesn't set defaults
	setMsg(*otg.Ospfv3InterfaceArea) Ospfv3InterfaceArea
	// provides marshal interface
	Marshal() marshalOspfv3InterfaceArea
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3InterfaceArea
	// validate validates Ospfv3InterfaceArea
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3InterfaceArea, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Ospfv3InterfaceAreaChoiceEnum, set in Ospfv3InterfaceArea
	Choice() Ospfv3InterfaceAreaChoiceEnum
	// setChoice assigns Ospfv3InterfaceAreaChoiceEnum provided by user to Ospfv3InterfaceArea
	setChoice(value Ospfv3InterfaceAreaChoiceEnum) Ospfv3InterfaceArea
	// HasChoice checks if Choice has been set in Ospfv3InterfaceArea
	HasChoice() bool
	// Id returns uint32, set in Ospfv3InterfaceArea.
	Id() uint32
	// SetId assigns uint32 provided by user to Ospfv3InterfaceArea
	SetId(value uint32) Ospfv3InterfaceArea
	// HasId checks if Id has been set in Ospfv3InterfaceArea
	HasId() bool
	// Ip returns string, set in Ospfv3InterfaceArea.
	Ip() string
	// SetIp assigns string provided by user to Ospfv3InterfaceArea
	SetIp(value string) Ospfv3InterfaceArea
	// HasIp checks if Ip has been set in Ospfv3InterfaceArea
	HasIp() bool
}

type Ospfv3InterfaceAreaChoiceEnum string

// Enum of Choice on Ospfv3InterfaceArea
var Ospfv3InterfaceAreaChoice = struct {
	ID Ospfv3InterfaceAreaChoiceEnum
	IP Ospfv3InterfaceAreaChoiceEnum
}{
	ID: Ospfv3InterfaceAreaChoiceEnum("id"),
	IP: Ospfv3InterfaceAreaChoiceEnum("ip"),
}

func (obj *ospfv3InterfaceArea) Choice() Ospfv3InterfaceAreaChoiceEnum {
	return Ospfv3InterfaceAreaChoiceEnum(obj.obj.Choice.Enum().String())
}

// The OSPFv3 Area ID identifies the routing area to which the host belongs. Area ID type can be following format.
// - id: A 32-bit number identifying the area.
// - ip: The Area ID in IPv4 address format.
// Choice returns a string
func (obj *ospfv3InterfaceArea) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ospfv3InterfaceArea) setChoice(value Ospfv3InterfaceAreaChoiceEnum) Ospfv3InterfaceArea {
	intValue, ok := otg.Ospfv3InterfaceArea_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv3InterfaceAreaChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv3InterfaceArea_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Ip = nil
	obj.obj.Id = nil

	if value == Ospfv3InterfaceAreaChoice.ID {
		defaultValue := uint32(0)
		obj.obj.Id = &defaultValue
	}

	return obj
}

// The Area ID.
// Id returns a uint32
func (obj *ospfv3InterfaceArea) Id() uint32 {

	if obj.obj.Id == nil {
		obj.setChoice(Ospfv3InterfaceAreaChoice.ID)
	}

	return *obj.obj.Id

}

// The Area ID.
// Id returns a uint32
func (obj *ospfv3InterfaceArea) HasId() bool {
	return obj.obj.Id != nil
}

// The Area ID.
// SetId sets the uint32 value in the Ospfv3InterfaceArea object
func (obj *ospfv3InterfaceArea) SetId(value uint32) Ospfv3InterfaceArea {
	obj.setChoice(Ospfv3InterfaceAreaChoice.ID)
	obj.obj.Id = &value
	return obj
}

// The Area ID in IPv4 address format.
// Ip returns a string
func (obj *ospfv3InterfaceArea) Ip() string {

	if obj.obj.Ip == nil {
		obj.setChoice(Ospfv3InterfaceAreaChoice.IP)
	}

	return *obj.obj.Ip

}

// The Area ID in IPv4 address format.
// Ip returns a string
func (obj *ospfv3InterfaceArea) HasIp() bool {
	return obj.obj.Ip != nil
}

// The Area ID in IPv4 address format.
// SetIp sets the string value in the Ospfv3InterfaceArea object
func (obj *ospfv3InterfaceArea) SetIp(value string) Ospfv3InterfaceArea {
	obj.setChoice(Ospfv3InterfaceAreaChoice.IP)
	obj.obj.Ip = &value
	return obj
}

func (obj *ospfv3InterfaceArea) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ip != nil {

		err := obj.validateIpv4(obj.Ip())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv3InterfaceArea.Ip"))
		}

	}

}

func (obj *ospfv3InterfaceArea) setDefault() {
	var choices_set int = 0
	var choice Ospfv3InterfaceAreaChoiceEnum

	if obj.obj.Id != nil {
		choices_set += 1
		choice = Ospfv3InterfaceAreaChoice.ID
	}

	if obj.obj.Ip != nil {
		choices_set += 1
		choice = Ospfv3InterfaceAreaChoice.IP
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Ospfv3InterfaceAreaChoice.ID)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Ospfv3InterfaceArea")
			}
		} else {
			intVal := otg.Ospfv3InterfaceArea_Choice_Enum_value[string(choice)]
			enumValue := otg.Ospfv3InterfaceArea_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

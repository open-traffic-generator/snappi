package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LldpChassisId *****
type lldpChassisId struct {
	validation
	obj                     *otg.LldpChassisId
	marshaller              marshalLldpChassisId
	unMarshaller            unMarshalLldpChassisId
	macAddressSubtypeHolder LldpChassisMacSubType
}

func NewLldpChassisId() LldpChassisId {
	obj := lldpChassisId{obj: &otg.LldpChassisId{}}
	obj.setDefault()
	return &obj
}

func (obj *lldpChassisId) msg() *otg.LldpChassisId {
	return obj.obj
}

func (obj *lldpChassisId) setMsg(msg *otg.LldpChassisId) LldpChassisId {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallldpChassisId struct {
	obj *lldpChassisId
}

type marshalLldpChassisId interface {
	// ToProto marshals LldpChassisId to protobuf object *otg.LldpChassisId
	ToProto() (*otg.LldpChassisId, error)
	// ToPbText marshals LldpChassisId to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LldpChassisId to YAML text
	ToYaml() (string, error)
	// ToJson marshals LldpChassisId to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals LldpChassisId to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshallldpChassisId struct {
	obj *lldpChassisId
}

type unMarshalLldpChassisId interface {
	// FromProto unmarshals LldpChassisId from protobuf object *otg.LldpChassisId
	FromProto(msg *otg.LldpChassisId) (LldpChassisId, error)
	// FromPbText unmarshals LldpChassisId from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LldpChassisId from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LldpChassisId from JSON text
	FromJson(value string) error
}

func (obj *lldpChassisId) Marshal() marshalLldpChassisId {
	if obj.marshaller == nil {
		obj.marshaller = &marshallldpChassisId{obj: obj}
	}
	return obj.marshaller
}

func (obj *lldpChassisId) Unmarshal() unMarshalLldpChassisId {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallldpChassisId{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallldpChassisId) ToProto() (*otg.LldpChassisId, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallldpChassisId) FromProto(msg *otg.LldpChassisId) (LldpChassisId, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallldpChassisId) ToPbText() (string, error) {
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

func (m *unMarshallldpChassisId) FromPbText(value string) error {
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

func (m *marshallldpChassisId) ToYaml() (string, error) {
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

func (m *unMarshallldpChassisId) FromYaml(value string) error {
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

func (m *marshallldpChassisId) ToJsonRaw() (string, error) {
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

func (m *marshallldpChassisId) ToJson() (string, error) {
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

func (m *unMarshallldpChassisId) FromJson(value string) error {
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

func (obj *lldpChassisId) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lldpChassisId) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lldpChassisId) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lldpChassisId) Clone() (LldpChassisId, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLldpChassisId()
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

func (obj *lldpChassisId) setNil() {
	obj.macAddressSubtypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LldpChassisId is the Chassis ID is a mandatory TLV which identifies the chassis component of the endpoint identifier associated  with the transmitting LLDP agent. This field identifies the format and source of the chassis identifier string.  It is based on the enumerator defined by the LldpChassisIdSubtype object from IEEE 802.1AB MIB.
type LldpChassisId interface {
	Validation
	// msg marshals LldpChassisId to protobuf object *otg.LldpChassisId
	// and doesn't set defaults
	msg() *otg.LldpChassisId
	// setMsg unmarshals LldpChassisId from protobuf object *otg.LldpChassisId
	// and doesn't set defaults
	setMsg(*otg.LldpChassisId) LldpChassisId
	// provides marshal interface
	Marshal() marshalLldpChassisId
	// provides unmarshal interface
	Unmarshal() unMarshalLldpChassisId
	// validate validates LldpChassisId
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LldpChassisId, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns LldpChassisIdChoiceEnum, set in LldpChassisId
	Choice() LldpChassisIdChoiceEnum
	// setChoice assigns LldpChassisIdChoiceEnum provided by user to LldpChassisId
	setChoice(value LldpChassisIdChoiceEnum) LldpChassisId
	// HasChoice checks if Choice has been set in LldpChassisId
	HasChoice() bool
	// MacAddressSubtype returns LldpChassisMacSubType, set in LldpChassisId.
	// LldpChassisMacSubType is the MAC address configured in the Chassis ID TLV.
	MacAddressSubtype() LldpChassisMacSubType
	// SetMacAddressSubtype assigns LldpChassisMacSubType provided by user to LldpChassisId.
	// LldpChassisMacSubType is the MAC address configured in the Chassis ID TLV.
	SetMacAddressSubtype(value LldpChassisMacSubType) LldpChassisId
	// HasMacAddressSubtype checks if MacAddressSubtype has been set in LldpChassisId
	HasMacAddressSubtype() bool
	// InterfaceNameSubtype returns string, set in LldpChassisId.
	InterfaceNameSubtype() string
	// SetInterfaceNameSubtype assigns string provided by user to LldpChassisId
	SetInterfaceNameSubtype(value string) LldpChassisId
	// HasInterfaceNameSubtype checks if InterfaceNameSubtype has been set in LldpChassisId
	HasInterfaceNameSubtype() bool
	// LocalSubtype returns string, set in LldpChassisId.
	LocalSubtype() string
	// SetLocalSubtype assigns string provided by user to LldpChassisId
	SetLocalSubtype(value string) LldpChassisId
	// HasLocalSubtype checks if LocalSubtype has been set in LldpChassisId
	HasLocalSubtype() bool
	setNil()
}

type LldpChassisIdChoiceEnum string

// Enum of Choice on LldpChassisId
var LldpChassisIdChoice = struct {
	MAC_ADDRESS_SUBTYPE    LldpChassisIdChoiceEnum
	INTERFACE_NAME_SUBTYPE LldpChassisIdChoiceEnum
	LOCAL_SUBTYPE          LldpChassisIdChoiceEnum
}{
	MAC_ADDRESS_SUBTYPE:    LldpChassisIdChoiceEnum("mac_address_subtype"),
	INTERFACE_NAME_SUBTYPE: LldpChassisIdChoiceEnum("interface_name_subtype"),
	LOCAL_SUBTYPE:          LldpChassisIdChoiceEnum("local_subtype"),
}

func (obj *lldpChassisId) Choice() LldpChassisIdChoiceEnum {
	return LldpChassisIdChoiceEnum(obj.obj.Choice.Enum().String())
}

// Chassis ID subtype to be used in Chassis ID TLV.
// Choice returns a string
func (obj *lldpChassisId) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *lldpChassisId) setChoice(value LldpChassisIdChoiceEnum) LldpChassisId {
	intValue, ok := otg.LldpChassisId_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LldpChassisIdChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.LldpChassisId_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.LocalSubtype = nil
	obj.obj.InterfaceNameSubtype = nil
	obj.obj.MacAddressSubtype = nil
	obj.macAddressSubtypeHolder = nil

	if value == LldpChassisIdChoice.MAC_ADDRESS_SUBTYPE {
		obj.obj.MacAddressSubtype = NewLldpChassisMacSubType().msg()
	}

	return obj
}

// description is TBD
// MacAddressSubtype returns a LldpChassisMacSubType
func (obj *lldpChassisId) MacAddressSubtype() LldpChassisMacSubType {
	if obj.obj.MacAddressSubtype == nil {
		obj.setChoice(LldpChassisIdChoice.MAC_ADDRESS_SUBTYPE)
	}
	if obj.macAddressSubtypeHolder == nil {
		obj.macAddressSubtypeHolder = &lldpChassisMacSubType{obj: obj.obj.MacAddressSubtype}
	}
	return obj.macAddressSubtypeHolder
}

// description is TBD
// MacAddressSubtype returns a LldpChassisMacSubType
func (obj *lldpChassisId) HasMacAddressSubtype() bool {
	return obj.obj.MacAddressSubtype != nil
}

// description is TBD
// SetMacAddressSubtype sets the LldpChassisMacSubType value in the LldpChassisId object
func (obj *lldpChassisId) SetMacAddressSubtype(value LldpChassisMacSubType) LldpChassisId {
	obj.setChoice(LldpChassisIdChoice.MAC_ADDRESS_SUBTYPE)
	obj.macAddressSubtypeHolder = nil
	obj.obj.MacAddressSubtype = value.msg()

	return obj
}

// Name of an interface of the chassis that uniquely identifies the chassis.
// InterfaceNameSubtype returns a string
func (obj *lldpChassisId) InterfaceNameSubtype() string {

	if obj.obj.InterfaceNameSubtype == nil {
		obj.setChoice(LldpChassisIdChoice.INTERFACE_NAME_SUBTYPE)
	}

	return *obj.obj.InterfaceNameSubtype

}

// Name of an interface of the chassis that uniquely identifies the chassis.
// InterfaceNameSubtype returns a string
func (obj *lldpChassisId) HasInterfaceNameSubtype() bool {
	return obj.obj.InterfaceNameSubtype != nil
}

// Name of an interface of the chassis that uniquely identifies the chassis.
// SetInterfaceNameSubtype sets the string value in the LldpChassisId object
func (obj *lldpChassisId) SetInterfaceNameSubtype(value string) LldpChassisId {
	obj.setChoice(LldpChassisIdChoice.INTERFACE_NAME_SUBTYPE)
	obj.obj.InterfaceNameSubtype = &value
	return obj
}

// Locally assigned name of the chassis.
// LocalSubtype returns a string
func (obj *lldpChassisId) LocalSubtype() string {

	if obj.obj.LocalSubtype == nil {
		obj.setChoice(LldpChassisIdChoice.LOCAL_SUBTYPE)
	}

	return *obj.obj.LocalSubtype

}

// Locally assigned name of the chassis.
// LocalSubtype returns a string
func (obj *lldpChassisId) HasLocalSubtype() bool {
	return obj.obj.LocalSubtype != nil
}

// Locally assigned name of the chassis.
// SetLocalSubtype sets the string value in the LldpChassisId object
func (obj *lldpChassisId) SetLocalSubtype(value string) LldpChassisId {
	obj.setChoice(LldpChassisIdChoice.LOCAL_SUBTYPE)
	obj.obj.LocalSubtype = &value
	return obj
}

func (obj *lldpChassisId) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MacAddressSubtype != nil {

		obj.MacAddressSubtype().validateObj(vObj, set_default)
	}

}

func (obj *lldpChassisId) setDefault() {
	var choices_set int = 0
	var choice LldpChassisIdChoiceEnum

	if obj.obj.MacAddressSubtype != nil {
		choices_set += 1
		choice = LldpChassisIdChoice.MAC_ADDRESS_SUBTYPE
	}

	if obj.obj.InterfaceNameSubtype != nil {
		choices_set += 1
		choice = LldpChassisIdChoice.INTERFACE_NAME_SUBTYPE
	}

	if obj.obj.LocalSubtype != nil {
		choices_set += 1
		choice = LldpChassisIdChoice.LOCAL_SUBTYPE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(LldpChassisIdChoice.MAC_ADDRESS_SUBTYPE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in LldpChassisId")
			}
		} else {
			intVal := otg.LldpChassisId_Choice_Enum_value[string(choice)]
			enumValue := otg.LldpChassisId_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

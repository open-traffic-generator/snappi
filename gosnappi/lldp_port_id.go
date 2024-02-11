package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LldpPortId *****
type lldpPortId struct {
	validation
	obj                        *otg.LldpPortId
	marshaller                 marshalLldpPortId
	unMarshaller               unMarshalLldpPortId
	interfaceNameSubtypeHolder LldpPortInterfaceNameSubType
}

func NewLldpPortId() LldpPortId {
	obj := lldpPortId{obj: &otg.LldpPortId{}}
	obj.setDefault()
	return &obj
}

func (obj *lldpPortId) msg() *otg.LldpPortId {
	return obj.obj
}

func (obj *lldpPortId) setMsg(msg *otg.LldpPortId) LldpPortId {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallldpPortId struct {
	obj *lldpPortId
}

type marshalLldpPortId interface {
	// ToProto marshals LldpPortId to protobuf object *otg.LldpPortId
	ToProto() (*otg.LldpPortId, error)
	// ToPbText marshals LldpPortId to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LldpPortId to YAML text
	ToYaml() (string, error)
	// ToJson marshals LldpPortId to JSON text
	ToJson() (string, error)
}

type unMarshallldpPortId struct {
	obj *lldpPortId
}

type unMarshalLldpPortId interface {
	// FromProto unmarshals LldpPortId from protobuf object *otg.LldpPortId
	FromProto(msg *otg.LldpPortId) (LldpPortId, error)
	// FromPbText unmarshals LldpPortId from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LldpPortId from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LldpPortId from JSON text
	FromJson(value string) error
}

func (obj *lldpPortId) Marshal() marshalLldpPortId {
	if obj.marshaller == nil {
		obj.marshaller = &marshallldpPortId{obj: obj}
	}
	return obj.marshaller
}

func (obj *lldpPortId) Unmarshal() unMarshalLldpPortId {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallldpPortId{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallldpPortId) ToProto() (*otg.LldpPortId, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallldpPortId) FromProto(msg *otg.LldpPortId) (LldpPortId, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallldpPortId) ToPbText() (string, error) {
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

func (m *unMarshallldpPortId) FromPbText(value string) error {
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

func (m *marshallldpPortId) ToYaml() (string, error) {
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

func (m *unMarshallldpPortId) FromYaml(value string) error {
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

func (m *marshallldpPortId) ToJson() (string, error) {
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

func (m *unMarshallldpPortId) FromJson(value string) error {
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

func (obj *lldpPortId) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lldpPortId) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lldpPortId) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lldpPortId) Clone() (LldpPortId, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLldpPortId()
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

func (obj *lldpPortId) setNil() {
	obj.interfaceNameSubtypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LldpPortId is the Port ID is a mandatory TLV which identifies the port component of the endpoint identifier associated with  the transmitting LLDP agent.This field identifies the format and source of the port identifier string. It is  based on the enumerator defined by the PtopoPortIdType object from RFC2922.
type LldpPortId interface {
	Validation
	// msg marshals LldpPortId to protobuf object *otg.LldpPortId
	// and doesn't set defaults
	msg() *otg.LldpPortId
	// setMsg unmarshals LldpPortId from protobuf object *otg.LldpPortId
	// and doesn't set defaults
	setMsg(*otg.LldpPortId) LldpPortId
	// provides marshal interface
	Marshal() marshalLldpPortId
	// provides unmarshal interface
	Unmarshal() unMarshalLldpPortId
	// validate validates LldpPortId
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LldpPortId, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns LldpPortIdChoiceEnum, set in LldpPortId
	Choice() LldpPortIdChoiceEnum
	// setChoice assigns LldpPortIdChoiceEnum provided by user to LldpPortId
	setChoice(value LldpPortIdChoiceEnum) LldpPortId
	// HasChoice checks if Choice has been set in LldpPortId
	HasChoice() bool
	// MacAddressSubtype returns string, set in LldpPortId.
	MacAddressSubtype() string
	// SetMacAddressSubtype assigns string provided by user to LldpPortId
	SetMacAddressSubtype(value string) LldpPortId
	// HasMacAddressSubtype checks if MacAddressSubtype has been set in LldpPortId
	HasMacAddressSubtype() bool
	// InterfaceNameSubtype returns LldpPortInterfaceNameSubType, set in LldpPortId.
	// LldpPortInterfaceNameSubType is the interface name configured in the Port ID TLV.
	InterfaceNameSubtype() LldpPortInterfaceNameSubType
	// SetInterfaceNameSubtype assigns LldpPortInterfaceNameSubType provided by user to LldpPortId.
	// LldpPortInterfaceNameSubType is the interface name configured in the Port ID TLV.
	SetInterfaceNameSubtype(value LldpPortInterfaceNameSubType) LldpPortId
	// HasInterfaceNameSubtype checks if InterfaceNameSubtype has been set in LldpPortId
	HasInterfaceNameSubtype() bool
	// LocalSubtype returns string, set in LldpPortId.
	LocalSubtype() string
	// SetLocalSubtype assigns string provided by user to LldpPortId
	SetLocalSubtype(value string) LldpPortId
	// HasLocalSubtype checks if LocalSubtype has been set in LldpPortId
	HasLocalSubtype() bool
	setNil()
}

type LldpPortIdChoiceEnum string

// Enum of Choice on LldpPortId
var LldpPortIdChoice = struct {
	MAC_ADDRESS_SUBTYPE    LldpPortIdChoiceEnum
	INTERFACE_NAME_SUBTYPE LldpPortIdChoiceEnum
	LOCAL_SUBTYPE          LldpPortIdChoiceEnum
}{
	MAC_ADDRESS_SUBTYPE:    LldpPortIdChoiceEnum("mac_address_subtype"),
	INTERFACE_NAME_SUBTYPE: LldpPortIdChoiceEnum("interface_name_subtype"),
	LOCAL_SUBTYPE:          LldpPortIdChoiceEnum("local_subtype"),
}

func (obj *lldpPortId) Choice() LldpPortIdChoiceEnum {
	return LldpPortIdChoiceEnum(obj.obj.Choice.Enum().String())
}

// Port ID subtype to be used in Port ID TLV.
// Choice returns a string
func (obj *lldpPortId) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *lldpPortId) setChoice(value LldpPortIdChoiceEnum) LldpPortId {
	intValue, ok := otg.LldpPortId_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LldpPortIdChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.LldpPortId_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.LocalSubtype = nil
	obj.obj.InterfaceNameSubtype = nil
	obj.interfaceNameSubtypeHolder = nil
	obj.obj.MacAddressSubtype = nil

	if value == LldpPortIdChoice.INTERFACE_NAME_SUBTYPE {
		obj.obj.InterfaceNameSubtype = NewLldpPortInterfaceNameSubType().msg()
	}

	return obj
}

// The MAC Address configured in the Port ID TLV.
// MacAddressSubtype returns a string
func (obj *lldpPortId) MacAddressSubtype() string {

	if obj.obj.MacAddressSubtype == nil {
		obj.setChoice(LldpPortIdChoice.MAC_ADDRESS_SUBTYPE)
	}

	return *obj.obj.MacAddressSubtype

}

// The MAC Address configured in the Port ID TLV.
// MacAddressSubtype returns a string
func (obj *lldpPortId) HasMacAddressSubtype() bool {
	return obj.obj.MacAddressSubtype != nil
}

// The MAC Address configured in the Port ID TLV.
// SetMacAddressSubtype sets the string value in the LldpPortId object
func (obj *lldpPortId) SetMacAddressSubtype(value string) LldpPortId {
	obj.setChoice(LldpPortIdChoice.MAC_ADDRESS_SUBTYPE)
	obj.obj.MacAddressSubtype = &value
	return obj
}

// description is TBD
// InterfaceNameSubtype returns a LldpPortInterfaceNameSubType
func (obj *lldpPortId) InterfaceNameSubtype() LldpPortInterfaceNameSubType {
	if obj.obj.InterfaceNameSubtype == nil {
		obj.setChoice(LldpPortIdChoice.INTERFACE_NAME_SUBTYPE)
	}
	if obj.interfaceNameSubtypeHolder == nil {
		obj.interfaceNameSubtypeHolder = &lldpPortInterfaceNameSubType{obj: obj.obj.InterfaceNameSubtype}
	}
	return obj.interfaceNameSubtypeHolder
}

// description is TBD
// InterfaceNameSubtype returns a LldpPortInterfaceNameSubType
func (obj *lldpPortId) HasInterfaceNameSubtype() bool {
	return obj.obj.InterfaceNameSubtype != nil
}

// description is TBD
// SetInterfaceNameSubtype sets the LldpPortInterfaceNameSubType value in the LldpPortId object
func (obj *lldpPortId) SetInterfaceNameSubtype(value LldpPortInterfaceNameSubType) LldpPortId {
	obj.setChoice(LldpPortIdChoice.INTERFACE_NAME_SUBTYPE)
	obj.interfaceNameSubtypeHolder = nil
	obj.obj.InterfaceNameSubtype = value.msg()

	return obj
}

// The Locally assigned name configured in the Port ID TLV.
// LocalSubtype returns a string
func (obj *lldpPortId) LocalSubtype() string {

	if obj.obj.LocalSubtype == nil {
		obj.setChoice(LldpPortIdChoice.LOCAL_SUBTYPE)
	}

	return *obj.obj.LocalSubtype

}

// The Locally assigned name configured in the Port ID TLV.
// LocalSubtype returns a string
func (obj *lldpPortId) HasLocalSubtype() bool {
	return obj.obj.LocalSubtype != nil
}

// The Locally assigned name configured in the Port ID TLV.
// SetLocalSubtype sets the string value in the LldpPortId object
func (obj *lldpPortId) SetLocalSubtype(value string) LldpPortId {
	obj.setChoice(LldpPortIdChoice.LOCAL_SUBTYPE)
	obj.obj.LocalSubtype = &value
	return obj
}

func (obj *lldpPortId) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.InterfaceNameSubtype != nil {

		obj.InterfaceNameSubtype().validateObj(vObj, set_default)
	}

}

func (obj *lldpPortId) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(LldpPortIdChoice.INTERFACE_NAME_SUBTYPE)

	}

}

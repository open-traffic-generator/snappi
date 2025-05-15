package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ServerIaType *****
type dhcpv6ServerIaType struct {
	validation
	obj          *otg.Dhcpv6ServerIaType
	marshaller   marshalDhcpv6ServerIaType
	unMarshaller unMarshalDhcpv6ServerIaType
	ianaHolder   Dhcpv6ServerPoolInfo
	iataHolder   Dhcpv6ServerPoolInfo
	iapdHolder   Dhcpv6ServerIapdPoolInfo
	ianapdHolder Dhcpv6ServerIanapdPoolInfo
}

func NewDhcpv6ServerIaType() Dhcpv6ServerIaType {
	obj := dhcpv6ServerIaType{obj: &otg.Dhcpv6ServerIaType{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ServerIaType) msg() *otg.Dhcpv6ServerIaType {
	return obj.obj
}

func (obj *dhcpv6ServerIaType) setMsg(msg *otg.Dhcpv6ServerIaType) Dhcpv6ServerIaType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ServerIaType struct {
	obj *dhcpv6ServerIaType
}

type marshalDhcpv6ServerIaType interface {
	// ToProto marshals Dhcpv6ServerIaType to protobuf object *otg.Dhcpv6ServerIaType
	ToProto() (*otg.Dhcpv6ServerIaType, error)
	// ToPbText marshals Dhcpv6ServerIaType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ServerIaType to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ServerIaType to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv6ServerIaType struct {
	obj *dhcpv6ServerIaType
}

type unMarshalDhcpv6ServerIaType interface {
	// FromProto unmarshals Dhcpv6ServerIaType from protobuf object *otg.Dhcpv6ServerIaType
	FromProto(msg *otg.Dhcpv6ServerIaType) (Dhcpv6ServerIaType, error)
	// FromPbText unmarshals Dhcpv6ServerIaType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ServerIaType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ServerIaType from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ServerIaType) Marshal() marshalDhcpv6ServerIaType {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ServerIaType{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ServerIaType) Unmarshal() unMarshalDhcpv6ServerIaType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ServerIaType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ServerIaType) ToProto() (*otg.Dhcpv6ServerIaType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ServerIaType) FromProto(msg *otg.Dhcpv6ServerIaType) (Dhcpv6ServerIaType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ServerIaType) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ServerIaType) FromPbText(value string) error {
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

func (m *marshaldhcpv6ServerIaType) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ServerIaType) FromYaml(value string) error {
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

func (m *marshaldhcpv6ServerIaType) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ServerIaType) FromJson(value string) error {
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

func (obj *dhcpv6ServerIaType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ServerIaType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ServerIaType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ServerIaType) Clone() (Dhcpv6ServerIaType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ServerIaType()
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

func (obj *dhcpv6ServerIaType) setNil() {
	obj.ianaHolder = nil
	obj.iataHolder = nil
	obj.iapdHolder = nil
	obj.ianapdHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Dhcpv6ServerIaType is description is TBD
type Dhcpv6ServerIaType interface {
	Validation
	// msg marshals Dhcpv6ServerIaType to protobuf object *otg.Dhcpv6ServerIaType
	// and doesn't set defaults
	msg() *otg.Dhcpv6ServerIaType
	// setMsg unmarshals Dhcpv6ServerIaType from protobuf object *otg.Dhcpv6ServerIaType
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ServerIaType) Dhcpv6ServerIaType
	// provides marshal interface
	Marshal() marshalDhcpv6ServerIaType
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ServerIaType
	// validate validates Dhcpv6ServerIaType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ServerIaType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Dhcpv6ServerIaTypeChoiceEnum, set in Dhcpv6ServerIaType
	Choice() Dhcpv6ServerIaTypeChoiceEnum
	// setChoice assigns Dhcpv6ServerIaTypeChoiceEnum provided by user to Dhcpv6ServerIaType
	setChoice(value Dhcpv6ServerIaTypeChoiceEnum) Dhcpv6ServerIaType
	// HasChoice checks if Choice has been set in Dhcpv6ServerIaType
	HasChoice() bool
	// Iana returns Dhcpv6ServerPoolInfo, set in Dhcpv6ServerIaType.
	// Dhcpv6ServerPoolInfo is the container for pool configurations for IA types iana and iata.
	Iana() Dhcpv6ServerPoolInfo
	// SetIana assigns Dhcpv6ServerPoolInfo provided by user to Dhcpv6ServerIaType.
	// Dhcpv6ServerPoolInfo is the container for pool configurations for IA types iana and iata.
	SetIana(value Dhcpv6ServerPoolInfo) Dhcpv6ServerIaType
	// HasIana checks if Iana has been set in Dhcpv6ServerIaType
	HasIana() bool
	// Iata returns Dhcpv6ServerPoolInfo, set in Dhcpv6ServerIaType.
	// Dhcpv6ServerPoolInfo is the container for pool configurations for IA types iana and iata.
	Iata() Dhcpv6ServerPoolInfo
	// SetIata assigns Dhcpv6ServerPoolInfo provided by user to Dhcpv6ServerIaType.
	// Dhcpv6ServerPoolInfo is the container for pool configurations for IA types iana and iata.
	SetIata(value Dhcpv6ServerPoolInfo) Dhcpv6ServerIaType
	// HasIata checks if Iata has been set in Dhcpv6ServerIaType
	HasIata() bool
	// Iapd returns Dhcpv6ServerIapdPoolInfo, set in Dhcpv6ServerIaType.
	// Dhcpv6ServerIapdPoolInfo is the container for prefix pool configurations for IA type iapd.
	Iapd() Dhcpv6ServerIapdPoolInfo
	// SetIapd assigns Dhcpv6ServerIapdPoolInfo provided by user to Dhcpv6ServerIaType.
	// Dhcpv6ServerIapdPoolInfo is the container for prefix pool configurations for IA type iapd.
	SetIapd(value Dhcpv6ServerIapdPoolInfo) Dhcpv6ServerIaType
	// HasIapd checks if Iapd has been set in Dhcpv6ServerIaType
	HasIapd() bool
	// Ianapd returns Dhcpv6ServerIanapdPoolInfo, set in Dhcpv6ServerIaType.
	// Dhcpv6ServerIanapdPoolInfo is the container for pool configurations for IA type ianapd.
	Ianapd() Dhcpv6ServerIanapdPoolInfo
	// SetIanapd assigns Dhcpv6ServerIanapdPoolInfo provided by user to Dhcpv6ServerIaType.
	// Dhcpv6ServerIanapdPoolInfo is the container for pool configurations for IA type ianapd.
	SetIanapd(value Dhcpv6ServerIanapdPoolInfo) Dhcpv6ServerIaType
	// HasIanapd checks if Ianapd has been set in Dhcpv6ServerIaType
	HasIanapd() bool
	setNil()
}

type Dhcpv6ServerIaTypeChoiceEnum string

// Enum of Choice on Dhcpv6ServerIaType
var Dhcpv6ServerIaTypeChoice = struct {
	IANA   Dhcpv6ServerIaTypeChoiceEnum
	IATA   Dhcpv6ServerIaTypeChoiceEnum
	IAPD   Dhcpv6ServerIaTypeChoiceEnum
	IANAPD Dhcpv6ServerIaTypeChoiceEnum
}{
	IANA:   Dhcpv6ServerIaTypeChoiceEnum("iana"),
	IATA:   Dhcpv6ServerIaTypeChoiceEnum("iata"),
	IAPD:   Dhcpv6ServerIaTypeChoiceEnum("iapd"),
	IANAPD: Dhcpv6ServerIaTypeChoiceEnum("ianapd"),
}

func (obj *dhcpv6ServerIaType) Choice() Dhcpv6ServerIaTypeChoiceEnum {
	return Dhcpv6ServerIaTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// Identity Association: a collection of leases assigned to a client. Each IA has an associated IAID.  Each IA holds one type of lease, like an identity association for temporary addresses (IA_TA) holds  temporary addresses, and an identity association for prefix delegation (IA_PD).
// Choice returns a string
func (obj *dhcpv6ServerIaType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *dhcpv6ServerIaType) setChoice(value Dhcpv6ServerIaTypeChoiceEnum) Dhcpv6ServerIaType {
	intValue, ok := otg.Dhcpv6ServerIaType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Dhcpv6ServerIaTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Dhcpv6ServerIaType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Ianapd = nil
	obj.ianapdHolder = nil
	obj.obj.Iapd = nil
	obj.iapdHolder = nil
	obj.obj.Iata = nil
	obj.iataHolder = nil
	obj.obj.Iana = nil
	obj.ianaHolder = nil

	if value == Dhcpv6ServerIaTypeChoice.IANA {
		obj.obj.Iana = NewDhcpv6ServerPoolInfo().msg()
	}

	if value == Dhcpv6ServerIaTypeChoice.IATA {
		obj.obj.Iata = NewDhcpv6ServerPoolInfo().msg()
	}

	if value == Dhcpv6ServerIaTypeChoice.IAPD {
		obj.obj.Iapd = NewDhcpv6ServerIapdPoolInfo().msg()
	}

	if value == Dhcpv6ServerIaTypeChoice.IANAPD {
		obj.obj.Ianapd = NewDhcpv6ServerIanapdPoolInfo().msg()
	}

	return obj
}

// description is TBD
// Iana returns a Dhcpv6ServerPoolInfo
func (obj *dhcpv6ServerIaType) Iana() Dhcpv6ServerPoolInfo {
	if obj.obj.Iana == nil {
		obj.setChoice(Dhcpv6ServerIaTypeChoice.IANA)
	}
	if obj.ianaHolder == nil {
		obj.ianaHolder = &dhcpv6ServerPoolInfo{obj: obj.obj.Iana}
	}
	return obj.ianaHolder
}

// description is TBD
// Iana returns a Dhcpv6ServerPoolInfo
func (obj *dhcpv6ServerIaType) HasIana() bool {
	return obj.obj.Iana != nil
}

// description is TBD
// SetIana sets the Dhcpv6ServerPoolInfo value in the Dhcpv6ServerIaType object
func (obj *dhcpv6ServerIaType) SetIana(value Dhcpv6ServerPoolInfo) Dhcpv6ServerIaType {
	obj.setChoice(Dhcpv6ServerIaTypeChoice.IANA)
	obj.ianaHolder = nil
	obj.obj.Iana = value.msg()

	return obj
}

// description is TBD
// Iata returns a Dhcpv6ServerPoolInfo
func (obj *dhcpv6ServerIaType) Iata() Dhcpv6ServerPoolInfo {
	if obj.obj.Iata == nil {
		obj.setChoice(Dhcpv6ServerIaTypeChoice.IATA)
	}
	if obj.iataHolder == nil {
		obj.iataHolder = &dhcpv6ServerPoolInfo{obj: obj.obj.Iata}
	}
	return obj.iataHolder
}

// description is TBD
// Iata returns a Dhcpv6ServerPoolInfo
func (obj *dhcpv6ServerIaType) HasIata() bool {
	return obj.obj.Iata != nil
}

// description is TBD
// SetIata sets the Dhcpv6ServerPoolInfo value in the Dhcpv6ServerIaType object
func (obj *dhcpv6ServerIaType) SetIata(value Dhcpv6ServerPoolInfo) Dhcpv6ServerIaType {
	obj.setChoice(Dhcpv6ServerIaTypeChoice.IATA)
	obj.iataHolder = nil
	obj.obj.Iata = value.msg()

	return obj
}

// description is TBD
// Iapd returns a Dhcpv6ServerIapdPoolInfo
func (obj *dhcpv6ServerIaType) Iapd() Dhcpv6ServerIapdPoolInfo {
	if obj.obj.Iapd == nil {
		obj.setChoice(Dhcpv6ServerIaTypeChoice.IAPD)
	}
	if obj.iapdHolder == nil {
		obj.iapdHolder = &dhcpv6ServerIapdPoolInfo{obj: obj.obj.Iapd}
	}
	return obj.iapdHolder
}

// description is TBD
// Iapd returns a Dhcpv6ServerIapdPoolInfo
func (obj *dhcpv6ServerIaType) HasIapd() bool {
	return obj.obj.Iapd != nil
}

// description is TBD
// SetIapd sets the Dhcpv6ServerIapdPoolInfo value in the Dhcpv6ServerIaType object
func (obj *dhcpv6ServerIaType) SetIapd(value Dhcpv6ServerIapdPoolInfo) Dhcpv6ServerIaType {
	obj.setChoice(Dhcpv6ServerIaTypeChoice.IAPD)
	obj.iapdHolder = nil
	obj.obj.Iapd = value.msg()

	return obj
}

// description is TBD
// Ianapd returns a Dhcpv6ServerIanapdPoolInfo
func (obj *dhcpv6ServerIaType) Ianapd() Dhcpv6ServerIanapdPoolInfo {
	if obj.obj.Ianapd == nil {
		obj.setChoice(Dhcpv6ServerIaTypeChoice.IANAPD)
	}
	if obj.ianapdHolder == nil {
		obj.ianapdHolder = &dhcpv6ServerIanapdPoolInfo{obj: obj.obj.Ianapd}
	}
	return obj.ianapdHolder
}

// description is TBD
// Ianapd returns a Dhcpv6ServerIanapdPoolInfo
func (obj *dhcpv6ServerIaType) HasIanapd() bool {
	return obj.obj.Ianapd != nil
}

// description is TBD
// SetIanapd sets the Dhcpv6ServerIanapdPoolInfo value in the Dhcpv6ServerIaType object
func (obj *dhcpv6ServerIaType) SetIanapd(value Dhcpv6ServerIanapdPoolInfo) Dhcpv6ServerIaType {
	obj.setChoice(Dhcpv6ServerIaTypeChoice.IANAPD)
	obj.ianapdHolder = nil
	obj.obj.Ianapd = value.msg()

	return obj
}

func (obj *dhcpv6ServerIaType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Iana != nil {

		obj.Iana().validateObj(vObj, set_default)
	}

	if obj.obj.Iata != nil {

		obj.Iata().validateObj(vObj, set_default)
	}

	if obj.obj.Iapd != nil {

		obj.Iapd().validateObj(vObj, set_default)
	}

	if obj.obj.Ianapd != nil {

		obj.Ianapd().validateObj(vObj, set_default)
	}

}

func (obj *dhcpv6ServerIaType) setDefault() {
	var choices_set int = 0
	var choice Dhcpv6ServerIaTypeChoiceEnum

	if obj.obj.Iana != nil {
		choices_set += 1
		choice = Dhcpv6ServerIaTypeChoice.IANA
	}

	if obj.obj.Iata != nil {
		choices_set += 1
		choice = Dhcpv6ServerIaTypeChoice.IATA
	}

	if obj.obj.Iapd != nil {
		choices_set += 1
		choice = Dhcpv6ServerIaTypeChoice.IAPD
	}

	if obj.obj.Ianapd != nil {
		choices_set += 1
		choice = Dhcpv6ServerIaTypeChoice.IANAPD
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Dhcpv6ServerIaTypeChoice.IANA)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Dhcpv6ServerIaType")
			}
		} else {
			intVal := otg.Dhcpv6ServerIaType_Choice_Enum_value[string(choice)]
			enumValue := otg.Dhcpv6ServerIaType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

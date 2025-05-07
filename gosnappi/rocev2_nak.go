package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2NAK *****
type rocev2NAK struct {
	validation
	obj          *otg.Rocev2NAK
	marshaller   marshalRocev2NAK
	unMarshaller unMarshalRocev2NAK
	ipDscpHolder Rocev2PriorityValue
}

func NewRocev2NAK() Rocev2NAK {
	obj := rocev2NAK{obj: &otg.Rocev2NAK{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2NAK) msg() *otg.Rocev2NAK {
	return obj.obj
}

func (obj *rocev2NAK) setMsg(msg *otg.Rocev2NAK) Rocev2NAK {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2NAK struct {
	obj *rocev2NAK
}

type marshalRocev2NAK interface {
	// ToProto marshals Rocev2NAK to protobuf object *otg.Rocev2NAK
	ToProto() (*otg.Rocev2NAK, error)
	// ToPbText marshals Rocev2NAK to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2NAK to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2NAK to JSON text
	ToJson() (string, error)
}

type unMarshalrocev2NAK struct {
	obj *rocev2NAK
}

type unMarshalRocev2NAK interface {
	// FromProto unmarshals Rocev2NAK from protobuf object *otg.Rocev2NAK
	FromProto(msg *otg.Rocev2NAK) (Rocev2NAK, error)
	// FromPbText unmarshals Rocev2NAK from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2NAK from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2NAK from JSON text
	FromJson(value string) error
}

func (obj *rocev2NAK) Marshal() marshalRocev2NAK {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2NAK{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2NAK) Unmarshal() unMarshalRocev2NAK {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2NAK{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2NAK) ToProto() (*otg.Rocev2NAK, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2NAK) FromProto(msg *otg.Rocev2NAK) (Rocev2NAK, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2NAK) ToPbText() (string, error) {
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

func (m *unMarshalrocev2NAK) FromPbText(value string) error {
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

func (m *marshalrocev2NAK) ToYaml() (string, error) {
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

func (m *unMarshalrocev2NAK) FromYaml(value string) error {
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

func (m *marshalrocev2NAK) ToJson() (string, error) {
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

func (m *unMarshalrocev2NAK) FromJson(value string) error {
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

func (obj *rocev2NAK) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2NAK) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2NAK) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2NAK) Clone() (Rocev2NAK, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2NAK()
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

func (obj *rocev2NAK) setNil() {
	obj.ipDscpHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2NAK is nAK parameters.
type Rocev2NAK interface {
	Validation
	// msg marshals Rocev2NAK to protobuf object *otg.Rocev2NAK
	// and doesn't set defaults
	msg() *otg.Rocev2NAK
	// setMsg unmarshals Rocev2NAK from protobuf object *otg.Rocev2NAK
	// and doesn't set defaults
	setMsg(*otg.Rocev2NAK) Rocev2NAK
	// provides marshal interface
	Marshal() marshalRocev2NAK
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2NAK
	// validate validates Rocev2NAK
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2NAK, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Rocev2NAKChoiceEnum, set in Rocev2NAK
	Choice() Rocev2NAKChoiceEnum
	// setChoice assigns Rocev2NAKChoiceEnum provided by user to Rocev2NAK
	setChoice(value Rocev2NAKChoiceEnum) Rocev2NAK
	// HasChoice checks if Choice has been set in Rocev2NAK
	HasChoice() bool
	// IpDscp returns Rocev2PriorityValue, set in Rocev2NAK.
	// Rocev2PriorityValue is priority value for CNP, ACK or NAK packets.
	IpDscp() Rocev2PriorityValue
	// SetIpDscp assigns Rocev2PriorityValue provided by user to Rocev2NAK.
	// Rocev2PriorityValue is priority value for CNP, ACK or NAK packets.
	SetIpDscp(value Rocev2PriorityValue) Rocev2NAK
	// HasIpDscp checks if IpDscp has been set in Rocev2NAK
	HasIpDscp() bool
	// EcnValue returns Rocev2NAKEcnValueEnum, set in Rocev2NAK
	EcnValue() Rocev2NAKEcnValueEnum
	// SetEcnValue assigns Rocev2NAKEcnValueEnum provided by user to Rocev2NAK
	SetEcnValue(value Rocev2NAKEcnValueEnum) Rocev2NAK
	// HasEcnValue checks if EcnValue has been set in Rocev2NAK
	HasEcnValue() bool
	setNil()
}

type Rocev2NAKChoiceEnum string

// Enum of Choice on Rocev2NAK
var Rocev2NAKChoice = struct {
	IP_DSCP Rocev2NAKChoiceEnum
}{
	IP_DSCP: Rocev2NAKChoiceEnum("ip_dscp"),
}

func (obj *rocev2NAK) Choice() Rocev2NAKChoiceEnum {
	return Rocev2NAKChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *rocev2NAK) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *rocev2NAK) setChoice(value Rocev2NAKChoiceEnum) Rocev2NAK {
	intValue, ok := otg.Rocev2NAK_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Rocev2NAKChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Rocev2NAK_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.IpDscp = nil
	obj.ipDscpHolder = nil

	if value == Rocev2NAKChoice.IP_DSCP {
		obj.obj.IpDscp = NewRocev2PriorityValue().msg()
	}

	return obj
}

// IP DSCP value for the NAK packets.
// IpDscp returns a Rocev2PriorityValue
func (obj *rocev2NAK) IpDscp() Rocev2PriorityValue {
	if obj.obj.IpDscp == nil {
		obj.setChoice(Rocev2NAKChoice.IP_DSCP)
	}
	if obj.ipDscpHolder == nil {
		obj.ipDscpHolder = &rocev2PriorityValue{obj: obj.obj.IpDscp}
	}
	return obj.ipDscpHolder
}

// IP DSCP value for the NAK packets.
// IpDscp returns a Rocev2PriorityValue
func (obj *rocev2NAK) HasIpDscp() bool {
	return obj.obj.IpDscp != nil
}

// IP DSCP value for the NAK packets.
// SetIpDscp sets the Rocev2PriorityValue value in the Rocev2NAK object
func (obj *rocev2NAK) SetIpDscp(value Rocev2PriorityValue) Rocev2NAK {
	obj.setChoice(Rocev2NAKChoice.IP_DSCP)
	obj.ipDscpHolder = nil
	obj.obj.IpDscp = value.msg()

	return obj
}

type Rocev2NAKEcnValueEnum string

// Enum of EcnValue on Rocev2NAK
var Rocev2NAKEcnValue = struct {
	NON_ECT Rocev2NAKEcnValueEnum
	ECT_1   Rocev2NAKEcnValueEnum
	ECT_0   Rocev2NAKEcnValueEnum
	CE      Rocev2NAKEcnValueEnum
}{
	NON_ECT: Rocev2NAKEcnValueEnum("non_ect"),
	ECT_1:   Rocev2NAKEcnValueEnum("ect_1"),
	ECT_0:   Rocev2NAKEcnValueEnum("ect_0"),
	CE:      Rocev2NAKEcnValueEnum("ce"),
}

func (obj *rocev2NAK) EcnValue() Rocev2NAKEcnValueEnum {
	return Rocev2NAKEcnValueEnum(obj.obj.EcnValue.Enum().String())
}

// ECN Code point to add in NAK packets.
// EcnValue returns a string
func (obj *rocev2NAK) HasEcnValue() bool {
	return obj.obj.EcnValue != nil
}

func (obj *rocev2NAK) SetEcnValue(value Rocev2NAKEcnValueEnum) Rocev2NAK {
	intValue, ok := otg.Rocev2NAK_EcnValue_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Rocev2NAKEcnValueEnum", string(value)))
		return obj
	}
	enumValue := otg.Rocev2NAK_EcnValue_Enum(intValue)
	obj.obj.EcnValue = &enumValue

	return obj
}

func (obj *rocev2NAK) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.IpDscp != nil {

		obj.IpDscp().validateObj(vObj, set_default)
	}

}

func (obj *rocev2NAK) setDefault() {
	var choices_set int = 0
	var choice Rocev2NAKChoiceEnum

	if obj.obj.IpDscp != nil {
		choices_set += 1
		choice = Rocev2NAKChoice.IP_DSCP
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Rocev2NAKChoice.IP_DSCP)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Rocev2NAK")
			}
		} else {
			intVal := otg.Rocev2NAK_Choice_Enum_value[string(choice)]
			enumValue := otg.Rocev2NAK_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.EcnValue == nil {
		obj.SetEcnValue(Rocev2NAKEcnValue.ECT_1)

	}

}

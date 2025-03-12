package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2NACK *****
type rocev2NACK struct {
	validation
	obj          *otg.Rocev2NACK
	marshaller   marshalRocev2NACK
	unMarshaller unMarshalRocev2NACK
	ipDscpHolder Rocev2PriorityValue
}

func NewRocev2NACK() Rocev2NACK {
	obj := rocev2NACK{obj: &otg.Rocev2NACK{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2NACK) msg() *otg.Rocev2NACK {
	return obj.obj
}

func (obj *rocev2NACK) setMsg(msg *otg.Rocev2NACK) Rocev2NACK {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2NACK struct {
	obj *rocev2NACK
}

type marshalRocev2NACK interface {
	// ToProto marshals Rocev2NACK to protobuf object *otg.Rocev2NACK
	ToProto() (*otg.Rocev2NACK, error)
	// ToPbText marshals Rocev2NACK to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2NACK to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2NACK to JSON text
	ToJson() (string, error)
}

type unMarshalrocev2NACK struct {
	obj *rocev2NACK
}

type unMarshalRocev2NACK interface {
	// FromProto unmarshals Rocev2NACK from protobuf object *otg.Rocev2NACK
	FromProto(msg *otg.Rocev2NACK) (Rocev2NACK, error)
	// FromPbText unmarshals Rocev2NACK from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2NACK from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2NACK from JSON text
	FromJson(value string) error
}

func (obj *rocev2NACK) Marshal() marshalRocev2NACK {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2NACK{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2NACK) Unmarshal() unMarshalRocev2NACK {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2NACK{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2NACK) ToProto() (*otg.Rocev2NACK, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2NACK) FromProto(msg *otg.Rocev2NACK) (Rocev2NACK, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2NACK) ToPbText() (string, error) {
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

func (m *unMarshalrocev2NACK) FromPbText(value string) error {
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

func (m *marshalrocev2NACK) ToYaml() (string, error) {
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

func (m *unMarshalrocev2NACK) FromYaml(value string) error {
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

func (m *marshalrocev2NACK) ToJson() (string, error) {
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

func (m *unMarshalrocev2NACK) FromJson(value string) error {
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

func (obj *rocev2NACK) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2NACK) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2NACK) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2NACK) Clone() (Rocev2NACK, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2NACK()
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

func (obj *rocev2NACK) setNil() {
	obj.ipDscpHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2NACK is nACK parameters.
type Rocev2NACK interface {
	Validation
	// msg marshals Rocev2NACK to protobuf object *otg.Rocev2NACK
	// and doesn't set defaults
	msg() *otg.Rocev2NACK
	// setMsg unmarshals Rocev2NACK from protobuf object *otg.Rocev2NACK
	// and doesn't set defaults
	setMsg(*otg.Rocev2NACK) Rocev2NACK
	// provides marshal interface
	Marshal() marshalRocev2NACK
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2NACK
	// validate validates Rocev2NACK
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2NACK, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Rocev2NACKChoiceEnum, set in Rocev2NACK
	Choice() Rocev2NACKChoiceEnum
	// setChoice assigns Rocev2NACKChoiceEnum provided by user to Rocev2NACK
	setChoice(value Rocev2NACKChoiceEnum) Rocev2NACK
	// HasChoice checks if Choice has been set in Rocev2NACK
	HasChoice() bool
	// IpDscp returns Rocev2PriorityValue, set in Rocev2NACK.
	// Rocev2PriorityValue is value.
	IpDscp() Rocev2PriorityValue
	// SetIpDscp assigns Rocev2PriorityValue provided by user to Rocev2NACK.
	// Rocev2PriorityValue is value.
	SetIpDscp(value Rocev2PriorityValue) Rocev2NACK
	// HasIpDscp checks if IpDscp has been set in Rocev2NACK
	HasIpDscp() bool
	// EcnValue returns Rocev2NACKEcnValueEnum, set in Rocev2NACK
	EcnValue() Rocev2NACKEcnValueEnum
	// SetEcnValue assigns Rocev2NACKEcnValueEnum provided by user to Rocev2NACK
	SetEcnValue(value Rocev2NACKEcnValueEnum) Rocev2NACK
	// HasEcnValue checks if EcnValue has been set in Rocev2NACK
	HasEcnValue() bool
	setNil()
}

type Rocev2NACKChoiceEnum string

// Enum of Choice on Rocev2NACK
var Rocev2NACKChoice = struct {
	IP_DSCP Rocev2NACKChoiceEnum
}{
	IP_DSCP: Rocev2NACKChoiceEnum("ip_dscp"),
}

func (obj *rocev2NACK) Choice() Rocev2NACKChoiceEnum {
	return Rocev2NACKChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *rocev2NACK) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *rocev2NACK) setChoice(value Rocev2NACKChoiceEnum) Rocev2NACK {
	intValue, ok := otg.Rocev2NACK_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Rocev2NACKChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Rocev2NACK_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.IpDscp = nil
	obj.ipDscpHolder = nil

	if value == Rocev2NACKChoice.IP_DSCP {
		obj.obj.IpDscp = NewRocev2PriorityValue().msg()
	}

	return obj
}

// The value that would be carried in NACK packet. Available option is IP DSCP. The given value is carried in IP Header DSCP value of the NACK packet.
// IpDscp returns a Rocev2PriorityValue
func (obj *rocev2NACK) IpDscp() Rocev2PriorityValue {
	if obj.obj.IpDscp == nil {
		obj.setChoice(Rocev2NACKChoice.IP_DSCP)
	}
	if obj.ipDscpHolder == nil {
		obj.ipDscpHolder = &rocev2PriorityValue{obj: obj.obj.IpDscp}
	}
	return obj.ipDscpHolder
}

// The value that would be carried in NACK packet. Available option is IP DSCP. The given value is carried in IP Header DSCP value of the NACK packet.
// IpDscp returns a Rocev2PriorityValue
func (obj *rocev2NACK) HasIpDscp() bool {
	return obj.obj.IpDscp != nil
}

// The value that would be carried in NACK packet. Available option is IP DSCP. The given value is carried in IP Header DSCP value of the NACK packet.
// SetIpDscp sets the Rocev2PriorityValue value in the Rocev2NACK object
func (obj *rocev2NACK) SetIpDscp(value Rocev2PriorityValue) Rocev2NACK {
	obj.setChoice(Rocev2NACKChoice.IP_DSCP)
	obj.ipDscpHolder = nil
	obj.obj.IpDscp = value.msg()

	return obj
}

type Rocev2NACKEcnValueEnum string

// Enum of EcnValue on Rocev2NACK
var Rocev2NACKEcnValue = struct {
	NON_ECT Rocev2NACKEcnValueEnum
	ECT_1   Rocev2NACKEcnValueEnum
	ECT_0   Rocev2NACKEcnValueEnum
	CE      Rocev2NACKEcnValueEnum
}{
	NON_ECT: Rocev2NACKEcnValueEnum("non_ect"),
	ECT_1:   Rocev2NACKEcnValueEnum("ect_1"),
	ECT_0:   Rocev2NACKEcnValueEnum("ect_0"),
	CE:      Rocev2NACKEcnValueEnum("ce"),
}

func (obj *rocev2NACK) EcnValue() Rocev2NACKEcnValueEnum {
	return Rocev2NACKEcnValueEnum(obj.obj.EcnValue.Enum().String())
}

// NACK ECN Value.
// EcnValue returns a string
func (obj *rocev2NACK) HasEcnValue() bool {
	return obj.obj.EcnValue != nil
}

func (obj *rocev2NACK) SetEcnValue(value Rocev2NACKEcnValueEnum) Rocev2NACK {
	intValue, ok := otg.Rocev2NACK_EcnValue_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Rocev2NACKEcnValueEnum", string(value)))
		return obj
	}
	enumValue := otg.Rocev2NACK_EcnValue_Enum(intValue)
	obj.obj.EcnValue = &enumValue

	return obj
}

func (obj *rocev2NACK) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.IpDscp != nil {

		obj.IpDscp().validateObj(vObj, set_default)
	}

}

func (obj *rocev2NACK) setDefault() {
	var choices_set int = 0
	var choice Rocev2NACKChoiceEnum

	if obj.obj.IpDscp != nil {
		choices_set += 1
		choice = Rocev2NACKChoice.IP_DSCP
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Rocev2NACKChoice.IP_DSCP)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Rocev2NACK")
			}
		} else {
			intVal := otg.Rocev2NACK_Choice_Enum_value[string(choice)]
			enumValue := otg.Rocev2NACK_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.EcnValue == nil {
		obj.SetEcnValue(Rocev2NACKEcnValue.ECT_1)

	}

}

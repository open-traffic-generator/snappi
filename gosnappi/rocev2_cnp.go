package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2CNP *****
type rocev2CNP struct {
	validation
	obj          *otg.Rocev2CNP
	marshaller   marshalRocev2CNP
	unMarshaller unMarshalRocev2CNP
	ipDscpHolder Rocev2PriorityValue
}

func NewRocev2CNP() Rocev2CNP {
	obj := rocev2CNP{obj: &otg.Rocev2CNP{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2CNP) msg() *otg.Rocev2CNP {
	return obj.obj
}

func (obj *rocev2CNP) setMsg(msg *otg.Rocev2CNP) Rocev2CNP {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2CNP struct {
	obj *rocev2CNP
}

type marshalRocev2CNP interface {
	// ToProto marshals Rocev2CNP to protobuf object *otg.Rocev2CNP
	ToProto() (*otg.Rocev2CNP, error)
	// ToPbText marshals Rocev2CNP to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2CNP to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2CNP to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Rocev2CNP to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrocev2CNP struct {
	obj *rocev2CNP
}

type unMarshalRocev2CNP interface {
	// FromProto unmarshals Rocev2CNP from protobuf object *otg.Rocev2CNP
	FromProto(msg *otg.Rocev2CNP) (Rocev2CNP, error)
	// FromPbText unmarshals Rocev2CNP from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2CNP from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2CNP from JSON text
	FromJson(value string) error
}

func (obj *rocev2CNP) Marshal() marshalRocev2CNP {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2CNP{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2CNP) Unmarshal() unMarshalRocev2CNP {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2CNP{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2CNP) ToProto() (*otg.Rocev2CNP, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2CNP) FromProto(msg *otg.Rocev2CNP) (Rocev2CNP, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2CNP) ToPbText() (string, error) {
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

func (m *unMarshalrocev2CNP) FromPbText(value string) error {
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

func (m *marshalrocev2CNP) ToYaml() (string, error) {
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

func (m *unMarshalrocev2CNP) FromYaml(value string) error {
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

func (m *marshalrocev2CNP) ToJsonRaw() (string, error) {
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

func (m *marshalrocev2CNP) ToJson() (string, error) {
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

func (m *unMarshalrocev2CNP) FromJson(value string) error {
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

func (obj *rocev2CNP) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2CNP) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2CNP) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2CNP) Clone() (Rocev2CNP, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2CNP()
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

func (obj *rocev2CNP) setNil() {
	obj.ipDscpHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2CNP is cNP parameters.
type Rocev2CNP interface {
	Validation
	// msg marshals Rocev2CNP to protobuf object *otg.Rocev2CNP
	// and doesn't set defaults
	msg() *otg.Rocev2CNP
	// setMsg unmarshals Rocev2CNP from protobuf object *otg.Rocev2CNP
	// and doesn't set defaults
	setMsg(*otg.Rocev2CNP) Rocev2CNP
	// provides marshal interface
	Marshal() marshalRocev2CNP
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2CNP
	// validate validates Rocev2CNP
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2CNP, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Rocev2CNPChoiceEnum, set in Rocev2CNP
	Choice() Rocev2CNPChoiceEnum
	// setChoice assigns Rocev2CNPChoiceEnum provided by user to Rocev2CNP
	setChoice(value Rocev2CNPChoiceEnum) Rocev2CNP
	// HasChoice checks if Choice has been set in Rocev2CNP
	HasChoice() bool
	// IpDscp returns Rocev2PriorityValue, set in Rocev2CNP.
	// Rocev2PriorityValue is priority value for CNP, ACK or NAK packets.
	IpDscp() Rocev2PriorityValue
	// SetIpDscp assigns Rocev2PriorityValue provided by user to Rocev2CNP.
	// Rocev2PriorityValue is priority value for CNP, ACK or NAK packets.
	SetIpDscp(value Rocev2PriorityValue) Rocev2CNP
	// HasIpDscp checks if IpDscp has been set in Rocev2CNP
	HasIpDscp() bool
	// EcnValue returns Rocev2CNPEcnValueEnum, set in Rocev2CNP
	EcnValue() Rocev2CNPEcnValueEnum
	// SetEcnValue assigns Rocev2CNPEcnValueEnum provided by user to Rocev2CNP
	SetEcnValue(value Rocev2CNPEcnValueEnum) Rocev2CNP
	// HasEcnValue checks if EcnValue has been set in Rocev2CNP
	HasEcnValue() bool
	// CnpDelayTimer returns uint32, set in Rocev2CNP.
	CnpDelayTimer() uint32
	// SetCnpDelayTimer assigns uint32 provided by user to Rocev2CNP
	SetCnpDelayTimer(value uint32) Rocev2CNP
	// HasCnpDelayTimer checks if CnpDelayTimer has been set in Rocev2CNP
	HasCnpDelayTimer() bool
	setNil()
}

type Rocev2CNPChoiceEnum string

// Enum of Choice on Rocev2CNP
var Rocev2CNPChoice = struct {
	IP_DSCP Rocev2CNPChoiceEnum
}{
	IP_DSCP: Rocev2CNPChoiceEnum("ip_dscp"),
}

func (obj *rocev2CNP) Choice() Rocev2CNPChoiceEnum {
	return Rocev2CNPChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *rocev2CNP) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *rocev2CNP) setChoice(value Rocev2CNPChoiceEnum) Rocev2CNP {
	intValue, ok := otg.Rocev2CNP_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Rocev2CNPChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Rocev2CNP_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.IpDscp = nil
	obj.ipDscpHolder = nil

	if value == Rocev2CNPChoice.IP_DSCP {
		obj.obj.IpDscp = NewRocev2PriorityValue().msg()
	}

	return obj
}

// IP DSCP value for the CNP packets.
// IpDscp returns a Rocev2PriorityValue
func (obj *rocev2CNP) IpDscp() Rocev2PriorityValue {
	if obj.obj.IpDscp == nil {
		obj.setChoice(Rocev2CNPChoice.IP_DSCP)
	}
	if obj.ipDscpHolder == nil {
		obj.ipDscpHolder = &rocev2PriorityValue{obj: obj.obj.IpDscp}
	}
	return obj.ipDscpHolder
}

// IP DSCP value for the CNP packets.
// IpDscp returns a Rocev2PriorityValue
func (obj *rocev2CNP) HasIpDscp() bool {
	return obj.obj.IpDscp != nil
}

// IP DSCP value for the CNP packets.
// SetIpDscp sets the Rocev2PriorityValue value in the Rocev2CNP object
func (obj *rocev2CNP) SetIpDscp(value Rocev2PriorityValue) Rocev2CNP {
	obj.setChoice(Rocev2CNPChoice.IP_DSCP)
	obj.ipDscpHolder = nil
	obj.obj.IpDscp = value.msg()

	return obj
}

type Rocev2CNPEcnValueEnum string

// Enum of EcnValue on Rocev2CNP
var Rocev2CNPEcnValue = struct {
	NON_ECT Rocev2CNPEcnValueEnum
	ECT_1   Rocev2CNPEcnValueEnum
	ECT_0   Rocev2CNPEcnValueEnum
	CE      Rocev2CNPEcnValueEnum
}{
	NON_ECT: Rocev2CNPEcnValueEnum("non_ect"),
	ECT_1:   Rocev2CNPEcnValueEnum("ect_1"),
	ECT_0:   Rocev2CNPEcnValueEnum("ect_0"),
	CE:      Rocev2CNPEcnValueEnum("ce"),
}

func (obj *rocev2CNP) EcnValue() Rocev2CNPEcnValueEnum {
	return Rocev2CNPEcnValueEnum(obj.obj.EcnValue.Enum().String())
}

// CNP ECN Value.
// EcnValue returns a string
func (obj *rocev2CNP) HasEcnValue() bool {
	return obj.obj.EcnValue != nil
}

func (obj *rocev2CNP) SetEcnValue(value Rocev2CNPEcnValueEnum) Rocev2CNP {
	intValue, ok := otg.Rocev2CNP_EcnValue_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Rocev2CNPEcnValueEnum", string(value)))
		return obj
	}
	enumValue := otg.Rocev2CNP_EcnValue_Enum(intValue)
	obj.obj.EcnValue = &enumValue

	return obj
}

// The interval duration between the generation of successive CNP packets should be specified in microseconds.
// CnpDelayTimer returns a uint32
func (obj *rocev2CNP) CnpDelayTimer() uint32 {

	return *obj.obj.CnpDelayTimer

}

// The interval duration between the generation of successive CNP packets should be specified in microseconds.
// CnpDelayTimer returns a uint32
func (obj *rocev2CNP) HasCnpDelayTimer() bool {
	return obj.obj.CnpDelayTimer != nil
}

// The interval duration between the generation of successive CNP packets should be specified in microseconds.
// SetCnpDelayTimer sets the uint32 value in the Rocev2CNP object
func (obj *rocev2CNP) SetCnpDelayTimer(value uint32) Rocev2CNP {

	obj.obj.CnpDelayTimer = &value
	return obj
}

func (obj *rocev2CNP) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.IpDscp != nil {

		obj.IpDscp().validateObj(vObj, set_default)
	}

	if obj.obj.CnpDelayTimer != nil {

		if *obj.obj.CnpDelayTimer > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2CNP.CnpDelayTimer <= 255 but Got %d", *obj.obj.CnpDelayTimer))
		}

	}

}

func (obj *rocev2CNP) setDefault() {
	var choices_set int = 0
	var choice Rocev2CNPChoiceEnum

	if obj.obj.IpDscp != nil {
		choices_set += 1
		choice = Rocev2CNPChoice.IP_DSCP
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Rocev2CNPChoice.IP_DSCP)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Rocev2CNP")
			}
		} else {
			intVal := otg.Rocev2CNP_Choice_Enum_value[string(choice)]
			enumValue := otg.Rocev2CNP_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.EcnValue == nil {
		obj.SetEcnValue(Rocev2CNPEcnValue.ECT_1)

	}
	if obj.obj.CnpDelayTimer == nil {
		obj.SetCnpDelayTimer(55)
	}

}

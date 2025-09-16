package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2ACK *****
type rocev2ACK struct {
	validation
	obj          *otg.Rocev2ACK
	marshaller   marshalRocev2ACK
	unMarshaller unMarshalRocev2ACK
	ipDscpHolder Rocev2PriorityValue
}

func NewRocev2ACK() Rocev2ACK {
	obj := rocev2ACK{obj: &otg.Rocev2ACK{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2ACK) msg() *otg.Rocev2ACK {
	return obj.obj
}

func (obj *rocev2ACK) setMsg(msg *otg.Rocev2ACK) Rocev2ACK {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2ACK struct {
	obj *rocev2ACK
}

type marshalRocev2ACK interface {
	// ToProto marshals Rocev2ACK to protobuf object *otg.Rocev2ACK
	ToProto() (*otg.Rocev2ACK, error)
	// ToPbText marshals Rocev2ACK to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2ACK to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2ACK to JSON text
	ToJson() (string, error)
}

type unMarshalrocev2ACK struct {
	obj *rocev2ACK
}

type unMarshalRocev2ACK interface {
	// FromProto unmarshals Rocev2ACK from protobuf object *otg.Rocev2ACK
	FromProto(msg *otg.Rocev2ACK) (Rocev2ACK, error)
	// FromPbText unmarshals Rocev2ACK from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2ACK from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2ACK from JSON text
	FromJson(value string) error
}

func (obj *rocev2ACK) Marshal() marshalRocev2ACK {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2ACK{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2ACK) Unmarshal() unMarshalRocev2ACK {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2ACK{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2ACK) ToProto() (*otg.Rocev2ACK, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2ACK) FromProto(msg *otg.Rocev2ACK) (Rocev2ACK, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2ACK) ToPbText() (string, error) {
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

func (m *unMarshalrocev2ACK) FromPbText(value string) error {
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

func (m *marshalrocev2ACK) ToYaml() (string, error) {
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

func (m *unMarshalrocev2ACK) FromYaml(value string) error {
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

func (m *marshalrocev2ACK) ToJson() (string, error) {
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

func (m *unMarshalrocev2ACK) FromJson(value string) error {
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

func (obj *rocev2ACK) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2ACK) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2ACK) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2ACK) Clone() (Rocev2ACK, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2ACK()
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

func (obj *rocev2ACK) setNil() {
	obj.ipDscpHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2ACK is aCK parameters.
type Rocev2ACK interface {
	Validation
	// msg marshals Rocev2ACK to protobuf object *otg.Rocev2ACK
	// and doesn't set defaults
	msg() *otg.Rocev2ACK
	// setMsg unmarshals Rocev2ACK from protobuf object *otg.Rocev2ACK
	// and doesn't set defaults
	setMsg(*otg.Rocev2ACK) Rocev2ACK
	// provides marshal interface
	Marshal() marshalRocev2ACK
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2ACK
	// validate validates Rocev2ACK
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2ACK, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Rocev2ACKChoiceEnum, set in Rocev2ACK
	Choice() Rocev2ACKChoiceEnum
	// setChoice assigns Rocev2ACKChoiceEnum provided by user to Rocev2ACK
	setChoice(value Rocev2ACKChoiceEnum) Rocev2ACK
	// HasChoice checks if Choice has been set in Rocev2ACK
	HasChoice() bool
	// IpDscp returns Rocev2PriorityValue, set in Rocev2ACK.
	// Rocev2PriorityValue is priority value for CNP, ACK or NAK packets.
	IpDscp() Rocev2PriorityValue
	// SetIpDscp assigns Rocev2PriorityValue provided by user to Rocev2ACK.
	// Rocev2PriorityValue is priority value for CNP, ACK or NAK packets.
	SetIpDscp(value Rocev2PriorityValue) Rocev2ACK
	// HasIpDscp checks if IpDscp has been set in Rocev2ACK
	HasIpDscp() bool
	// EcnValue returns Rocev2ACKEcnValueEnum, set in Rocev2ACK
	EcnValue() Rocev2ACKEcnValueEnum
	// SetEcnValue assigns Rocev2ACKEcnValueEnum provided by user to Rocev2ACK
	SetEcnValue(value Rocev2ACKEcnValueEnum) Rocev2ACK
	// HasEcnValue checks if EcnValue has been set in Rocev2ACK
	HasEcnValue() bool
	setNil()
}

type Rocev2ACKChoiceEnum string

// Enum of Choice on Rocev2ACK
var Rocev2ACKChoice = struct {
	IP_DSCP Rocev2ACKChoiceEnum
}{
	IP_DSCP: Rocev2ACKChoiceEnum("ip_dscp"),
}

func (obj *rocev2ACK) Choice() Rocev2ACKChoiceEnum {
	return Rocev2ACKChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *rocev2ACK) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *rocev2ACK) setChoice(value Rocev2ACKChoiceEnum) Rocev2ACK {
	intValue, ok := otg.Rocev2ACK_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Rocev2ACKChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Rocev2ACK_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.IpDscp = nil
	obj.ipDscpHolder = nil

	if value == Rocev2ACKChoice.IP_DSCP {
		obj.obj.IpDscp = NewRocev2PriorityValue().msg()
	}

	return obj
}

// IP DSCP value for the ACK packets.
// IpDscp returns a Rocev2PriorityValue
func (obj *rocev2ACK) IpDscp() Rocev2PriorityValue {
	if obj.obj.IpDscp == nil {
		obj.setChoice(Rocev2ACKChoice.IP_DSCP)
	}
	if obj.ipDscpHolder == nil {
		obj.ipDscpHolder = &rocev2PriorityValue{obj: obj.obj.IpDscp}
	}
	return obj.ipDscpHolder
}

// IP DSCP value for the ACK packets.
// IpDscp returns a Rocev2PriorityValue
func (obj *rocev2ACK) HasIpDscp() bool {
	return obj.obj.IpDscp != nil
}

// IP DSCP value for the ACK packets.
// SetIpDscp sets the Rocev2PriorityValue value in the Rocev2ACK object
func (obj *rocev2ACK) SetIpDscp(value Rocev2PriorityValue) Rocev2ACK {
	obj.setChoice(Rocev2ACKChoice.IP_DSCP)
	obj.ipDscpHolder = nil
	obj.obj.IpDscp = value.msg()

	return obj
}

type Rocev2ACKEcnValueEnum string

// Enum of EcnValue on Rocev2ACK
var Rocev2ACKEcnValue = struct {
	NON_ECT Rocev2ACKEcnValueEnum
	ECT_1   Rocev2ACKEcnValueEnum
	ECT_0   Rocev2ACKEcnValueEnum
	CE      Rocev2ACKEcnValueEnum
}{
	NON_ECT: Rocev2ACKEcnValueEnum("non_ect"),
	ECT_1:   Rocev2ACKEcnValueEnum("ect_1"),
	ECT_0:   Rocev2ACKEcnValueEnum("ect_0"),
	CE:      Rocev2ACKEcnValueEnum("ce"),
}

func (obj *rocev2ACK) EcnValue() Rocev2ACKEcnValueEnum {
	return Rocev2ACKEcnValueEnum(obj.obj.EcnValue.Enum().String())
}

// ACK ECN Value.
// EcnValue returns a string
func (obj *rocev2ACK) HasEcnValue() bool {
	return obj.obj.EcnValue != nil
}

func (obj *rocev2ACK) SetEcnValue(value Rocev2ACKEcnValueEnum) Rocev2ACK {
	intValue, ok := otg.Rocev2ACK_EcnValue_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Rocev2ACKEcnValueEnum", string(value)))
		return obj
	}
	enumValue := otg.Rocev2ACK_EcnValue_Enum(intValue)
	obj.obj.EcnValue = &enumValue

	return obj
}

func (obj *rocev2ACK) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.IpDscp != nil {

		obj.IpDscp().validateObj(vObj, set_default)
	}

}

func (obj *rocev2ACK) setDefault() {
	var choices_set int = 0
	var choice Rocev2ACKChoiceEnum

	if obj.obj.IpDscp != nil {
		choices_set += 1
		choice = Rocev2ACKChoice.IP_DSCP
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Rocev2ACKChoice.IP_DSCP)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Rocev2ACK")
			}
		} else {
			intVal := otg.Rocev2ACK_Choice_Enum_value[string(choice)]
			enumValue := otg.Rocev2ACK_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.EcnValue == nil {
		obj.SetEcnValue(Rocev2ACKEcnValue.ECT_1)

	}

}

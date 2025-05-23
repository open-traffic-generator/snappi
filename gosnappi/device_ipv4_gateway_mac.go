package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceIpv4GatewayMAC *****
type deviceIpv4GatewayMAC struct {
	validation
	obj          *otg.DeviceIpv4GatewayMAC
	marshaller   marshalDeviceIpv4GatewayMAC
	unMarshaller unMarshalDeviceIpv4GatewayMAC
}

func NewDeviceIpv4GatewayMAC() DeviceIpv4GatewayMAC {
	obj := deviceIpv4GatewayMAC{obj: &otg.DeviceIpv4GatewayMAC{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceIpv4GatewayMAC) msg() *otg.DeviceIpv4GatewayMAC {
	return obj.obj
}

func (obj *deviceIpv4GatewayMAC) setMsg(msg *otg.DeviceIpv4GatewayMAC) DeviceIpv4GatewayMAC {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceIpv4GatewayMAC struct {
	obj *deviceIpv4GatewayMAC
}

type marshalDeviceIpv4GatewayMAC interface {
	// ToProto marshals DeviceIpv4GatewayMAC to protobuf object *otg.DeviceIpv4GatewayMAC
	ToProto() (*otg.DeviceIpv4GatewayMAC, error)
	// ToPbText marshals DeviceIpv4GatewayMAC to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceIpv4GatewayMAC to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceIpv4GatewayMAC to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceIpv4GatewayMAC struct {
	obj *deviceIpv4GatewayMAC
}

type unMarshalDeviceIpv4GatewayMAC interface {
	// FromProto unmarshals DeviceIpv4GatewayMAC from protobuf object *otg.DeviceIpv4GatewayMAC
	FromProto(msg *otg.DeviceIpv4GatewayMAC) (DeviceIpv4GatewayMAC, error)
	// FromPbText unmarshals DeviceIpv4GatewayMAC from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceIpv4GatewayMAC from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceIpv4GatewayMAC from JSON text
	FromJson(value string) error
}

func (obj *deviceIpv4GatewayMAC) Marshal() marshalDeviceIpv4GatewayMAC {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceIpv4GatewayMAC{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceIpv4GatewayMAC) Unmarshal() unMarshalDeviceIpv4GatewayMAC {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceIpv4GatewayMAC{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceIpv4GatewayMAC) ToProto() (*otg.DeviceIpv4GatewayMAC, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceIpv4GatewayMAC) FromProto(msg *otg.DeviceIpv4GatewayMAC) (DeviceIpv4GatewayMAC, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceIpv4GatewayMAC) ToPbText() (string, error) {
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

func (m *unMarshaldeviceIpv4GatewayMAC) FromPbText(value string) error {
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

func (m *marshaldeviceIpv4GatewayMAC) ToYaml() (string, error) {
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

func (m *unMarshaldeviceIpv4GatewayMAC) FromYaml(value string) error {
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

func (m *marshaldeviceIpv4GatewayMAC) ToJson() (string, error) {
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

func (m *unMarshaldeviceIpv4GatewayMAC) FromJson(value string) error {
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

func (obj *deviceIpv4GatewayMAC) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceIpv4GatewayMAC) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceIpv4GatewayMAC) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceIpv4GatewayMAC) Clone() (DeviceIpv4GatewayMAC, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceIpv4GatewayMAC()
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

// DeviceIpv4GatewayMAC is by default auto(resolved gateway mac) is set.  Setting a value would mean that ARP will not be used for learning MAC of connected device. The user-configured MAC address will be used for auto-filling the destination
// MAC address in the control and data packets sent from this IPv4 endpoint
// whenever applicable.
type DeviceIpv4GatewayMAC interface {
	Validation
	// msg marshals DeviceIpv4GatewayMAC to protobuf object *otg.DeviceIpv4GatewayMAC
	// and doesn't set defaults
	msg() *otg.DeviceIpv4GatewayMAC
	// setMsg unmarshals DeviceIpv4GatewayMAC from protobuf object *otg.DeviceIpv4GatewayMAC
	// and doesn't set defaults
	setMsg(*otg.DeviceIpv4GatewayMAC) DeviceIpv4GatewayMAC
	// provides marshal interface
	Marshal() marshalDeviceIpv4GatewayMAC
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceIpv4GatewayMAC
	// validate validates DeviceIpv4GatewayMAC
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceIpv4GatewayMAC, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns DeviceIpv4GatewayMACChoiceEnum, set in DeviceIpv4GatewayMAC
	Choice() DeviceIpv4GatewayMACChoiceEnum
	// setChoice assigns DeviceIpv4GatewayMACChoiceEnum provided by user to DeviceIpv4GatewayMAC
	setChoice(value DeviceIpv4GatewayMACChoiceEnum) DeviceIpv4GatewayMAC
	// HasChoice checks if Choice has been set in DeviceIpv4GatewayMAC
	HasChoice() bool
	// Auto returns string, set in DeviceIpv4GatewayMAC.
	Auto() string
	// HasAuto checks if Auto has been set in DeviceIpv4GatewayMAC
	HasAuto() bool
	// Value returns string, set in DeviceIpv4GatewayMAC.
	Value() string
	// SetValue assigns string provided by user to DeviceIpv4GatewayMAC
	SetValue(value string) DeviceIpv4GatewayMAC
	// HasValue checks if Value has been set in DeviceIpv4GatewayMAC
	HasValue() bool
}

type DeviceIpv4GatewayMACChoiceEnum string

// Enum of Choice on DeviceIpv4GatewayMAC
var DeviceIpv4GatewayMACChoice = struct {
	AUTO  DeviceIpv4GatewayMACChoiceEnum
	VALUE DeviceIpv4GatewayMACChoiceEnum
}{
	AUTO:  DeviceIpv4GatewayMACChoiceEnum("auto"),
	VALUE: DeviceIpv4GatewayMACChoiceEnum("value"),
}

func (obj *deviceIpv4GatewayMAC) Choice() DeviceIpv4GatewayMACChoiceEnum {
	return DeviceIpv4GatewayMACChoiceEnum(obj.obj.Choice.Enum().String())
}

// auto or configured value.
// Choice returns a string
func (obj *deviceIpv4GatewayMAC) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *deviceIpv4GatewayMAC) setChoice(value DeviceIpv4GatewayMACChoiceEnum) DeviceIpv4GatewayMAC {
	intValue, ok := otg.DeviceIpv4GatewayMAC_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on DeviceIpv4GatewayMACChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.DeviceIpv4GatewayMAC_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Value = nil
	obj.obj.Auto = nil

	if value == DeviceIpv4GatewayMACChoice.AUTO {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Auto = &defaultValue
	}

	if value == DeviceIpv4GatewayMACChoice.VALUE {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Value = &defaultValue
	}

	return obj
}

// The OTG implementation can provide a system generated value for this property. If the OTG is unable to generate a value the default value must be used.
// Auto returns a string
func (obj *deviceIpv4GatewayMAC) Auto() string {

	if obj.obj.Auto == nil {
		obj.setChoice(DeviceIpv4GatewayMACChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated value for this property. If the OTG is unable to generate a value the default value must be used.
// Auto returns a string
func (obj *deviceIpv4GatewayMAC) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Value returns a string
func (obj *deviceIpv4GatewayMAC) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(DeviceIpv4GatewayMACChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *deviceIpv4GatewayMAC) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the DeviceIpv4GatewayMAC object
func (obj *deviceIpv4GatewayMAC) SetValue(value string) DeviceIpv4GatewayMAC {
	obj.setChoice(DeviceIpv4GatewayMACChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

func (obj *deviceIpv4GatewayMAC) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Auto != nil {

		err := obj.validateMac(obj.Auto())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIpv4GatewayMAC.Auto"))
		}

	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIpv4GatewayMAC.Value"))
		}

	}

}

func (obj *deviceIpv4GatewayMAC) setDefault() {
	var choices_set int = 0
	var choice DeviceIpv4GatewayMACChoiceEnum

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = DeviceIpv4GatewayMACChoice.AUTO
	}

	if obj.obj.Value != nil {
		choices_set += 1
		choice = DeviceIpv4GatewayMACChoice.VALUE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(DeviceIpv4GatewayMACChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in DeviceIpv4GatewayMAC")
			}
		} else {
			intVal := otg.DeviceIpv4GatewayMAC_Choice_Enum_value[string(choice)]
			enumValue := otg.DeviceIpv4GatewayMAC_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

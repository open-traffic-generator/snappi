package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceIpv6GatewayMAC *****
type deviceIpv6GatewayMAC struct {
	validation
	obj          *otg.DeviceIpv6GatewayMAC
	marshaller   marshalDeviceIpv6GatewayMAC
	unMarshaller unMarshalDeviceIpv6GatewayMAC
}

func NewDeviceIpv6GatewayMAC() DeviceIpv6GatewayMAC {
	obj := deviceIpv6GatewayMAC{obj: &otg.DeviceIpv6GatewayMAC{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceIpv6GatewayMAC) msg() *otg.DeviceIpv6GatewayMAC {
	return obj.obj
}

func (obj *deviceIpv6GatewayMAC) setMsg(msg *otg.DeviceIpv6GatewayMAC) DeviceIpv6GatewayMAC {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceIpv6GatewayMAC struct {
	obj *deviceIpv6GatewayMAC
}

type marshalDeviceIpv6GatewayMAC interface {
	// ToProto marshals DeviceIpv6GatewayMAC to protobuf object *otg.DeviceIpv6GatewayMAC
	ToProto() (*otg.DeviceIpv6GatewayMAC, error)
	// ToPbText marshals DeviceIpv6GatewayMAC to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceIpv6GatewayMAC to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceIpv6GatewayMAC to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceIpv6GatewayMAC struct {
	obj *deviceIpv6GatewayMAC
}

type unMarshalDeviceIpv6GatewayMAC interface {
	// FromProto unmarshals DeviceIpv6GatewayMAC from protobuf object *otg.DeviceIpv6GatewayMAC
	FromProto(msg *otg.DeviceIpv6GatewayMAC) (DeviceIpv6GatewayMAC, error)
	// FromPbText unmarshals DeviceIpv6GatewayMAC from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceIpv6GatewayMAC from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceIpv6GatewayMAC from JSON text
	FromJson(value string) error
}

func (obj *deviceIpv6GatewayMAC) Marshal() marshalDeviceIpv6GatewayMAC {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceIpv6GatewayMAC{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceIpv6GatewayMAC) Unmarshal() unMarshalDeviceIpv6GatewayMAC {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceIpv6GatewayMAC{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceIpv6GatewayMAC) ToProto() (*otg.DeviceIpv6GatewayMAC, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceIpv6GatewayMAC) FromProto(msg *otg.DeviceIpv6GatewayMAC) (DeviceIpv6GatewayMAC, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceIpv6GatewayMAC) ToPbText() (string, error) {
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

func (m *unMarshaldeviceIpv6GatewayMAC) FromPbText(value string) error {
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

func (m *marshaldeviceIpv6GatewayMAC) ToYaml() (string, error) {
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

func (m *unMarshaldeviceIpv6GatewayMAC) FromYaml(value string) error {
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

func (m *marshaldeviceIpv6GatewayMAC) ToJson() (string, error) {
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

func (m *unMarshaldeviceIpv6GatewayMAC) FromJson(value string) error {
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

func (obj *deviceIpv6GatewayMAC) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceIpv6GatewayMAC) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceIpv6GatewayMAC) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceIpv6GatewayMAC) Clone() (DeviceIpv6GatewayMAC, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceIpv6GatewayMAC()
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

// DeviceIpv6GatewayMAC is by default auto(resolved gateway mac) is set. Setting a value would mean that ND will not be used for learning MAC of connected device. The user-configured MAC address will be used for auto-filling the destination
// MAC address in the control and data packets sent from this IPv6 endpoint
// whenever applicable.
type DeviceIpv6GatewayMAC interface {
	Validation
	// msg marshals DeviceIpv6GatewayMAC to protobuf object *otg.DeviceIpv6GatewayMAC
	// and doesn't set defaults
	msg() *otg.DeviceIpv6GatewayMAC
	// setMsg unmarshals DeviceIpv6GatewayMAC from protobuf object *otg.DeviceIpv6GatewayMAC
	// and doesn't set defaults
	setMsg(*otg.DeviceIpv6GatewayMAC) DeviceIpv6GatewayMAC
	// provides marshal interface
	Marshal() marshalDeviceIpv6GatewayMAC
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceIpv6GatewayMAC
	// validate validates DeviceIpv6GatewayMAC
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceIpv6GatewayMAC, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns DeviceIpv6GatewayMACChoiceEnum, set in DeviceIpv6GatewayMAC
	Choice() DeviceIpv6GatewayMACChoiceEnum
	// setChoice assigns DeviceIpv6GatewayMACChoiceEnum provided by user to DeviceIpv6GatewayMAC
	setChoice(value DeviceIpv6GatewayMACChoiceEnum) DeviceIpv6GatewayMAC
	// HasChoice checks if Choice has been set in DeviceIpv6GatewayMAC
	HasChoice() bool
	// Auto returns string, set in DeviceIpv6GatewayMAC.
	Auto() string
	// HasAuto checks if Auto has been set in DeviceIpv6GatewayMAC
	HasAuto() bool
	// Value returns string, set in DeviceIpv6GatewayMAC.
	Value() string
	// SetValue assigns string provided by user to DeviceIpv6GatewayMAC
	SetValue(value string) DeviceIpv6GatewayMAC
	// HasValue checks if Value has been set in DeviceIpv6GatewayMAC
	HasValue() bool
}

type DeviceIpv6GatewayMACChoiceEnum string

// Enum of Choice on DeviceIpv6GatewayMAC
var DeviceIpv6GatewayMACChoice = struct {
	AUTO  DeviceIpv6GatewayMACChoiceEnum
	VALUE DeviceIpv6GatewayMACChoiceEnum
}{
	AUTO:  DeviceIpv6GatewayMACChoiceEnum("auto"),
	VALUE: DeviceIpv6GatewayMACChoiceEnum("value"),
}

func (obj *deviceIpv6GatewayMAC) Choice() DeviceIpv6GatewayMACChoiceEnum {
	return DeviceIpv6GatewayMACChoiceEnum(obj.obj.Choice.Enum().String())
}

// auto or configured value.
// Choice returns a string
func (obj *deviceIpv6GatewayMAC) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *deviceIpv6GatewayMAC) setChoice(value DeviceIpv6GatewayMACChoiceEnum) DeviceIpv6GatewayMAC {
	intValue, ok := otg.DeviceIpv6GatewayMAC_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on DeviceIpv6GatewayMACChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.DeviceIpv6GatewayMAC_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Value = nil
	obj.obj.Auto = nil

	if value == DeviceIpv6GatewayMACChoice.AUTO {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Auto = &defaultValue
	}

	if value == DeviceIpv6GatewayMACChoice.VALUE {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Value = &defaultValue
	}

	return obj
}

// The OTG implementation can provide a system generated value for this property. If the OTG is unable to generate a value the default value must be used.
// Auto returns a string
func (obj *deviceIpv6GatewayMAC) Auto() string {

	if obj.obj.Auto == nil {
		obj.setChoice(DeviceIpv6GatewayMACChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated value for this property. If the OTG is unable to generate a value the default value must be used.
// Auto returns a string
func (obj *deviceIpv6GatewayMAC) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Value returns a string
func (obj *deviceIpv6GatewayMAC) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(DeviceIpv6GatewayMACChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *deviceIpv6GatewayMAC) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the DeviceIpv6GatewayMAC object
func (obj *deviceIpv6GatewayMAC) SetValue(value string) DeviceIpv6GatewayMAC {
	obj.setChoice(DeviceIpv6GatewayMACChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

func (obj *deviceIpv6GatewayMAC) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Auto != nil {

		err := obj.validateMac(obj.Auto())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIpv6GatewayMAC.Auto"))
		}

	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIpv6GatewayMAC.Value"))
		}

	}

}

func (obj *deviceIpv6GatewayMAC) setDefault() {
	var choices_set int = 0
	var choice DeviceIpv6GatewayMACChoiceEnum

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = DeviceIpv6GatewayMACChoice.AUTO
	}

	if obj.obj.Value != nil {
		choices_set += 1
		choice = DeviceIpv6GatewayMACChoice.VALUE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(DeviceIpv6GatewayMACChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in DeviceIpv6GatewayMAC")
			}
		} else {
			intVal := otg.DeviceIpv6GatewayMAC_Choice_Enum_value[string(choice)]
			enumValue := otg.DeviceIpv6GatewayMAC_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ServerOptionsMessageType *****
type dhcpv6ServerOptionsMessageType struct {
	validation
	obj          *otg.Dhcpv6ServerOptionsMessageType
	marshaller   marshalDhcpv6ServerOptionsMessageType
	unMarshaller unMarshalDhcpv6ServerOptionsMessageType
}

func NewDhcpv6ServerOptionsMessageType() Dhcpv6ServerOptionsMessageType {
	obj := dhcpv6ServerOptionsMessageType{obj: &otg.Dhcpv6ServerOptionsMessageType{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ServerOptionsMessageType) msg() *otg.Dhcpv6ServerOptionsMessageType {
	return obj.obj
}

func (obj *dhcpv6ServerOptionsMessageType) setMsg(msg *otg.Dhcpv6ServerOptionsMessageType) Dhcpv6ServerOptionsMessageType {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ServerOptionsMessageType struct {
	obj *dhcpv6ServerOptionsMessageType
}

type marshalDhcpv6ServerOptionsMessageType interface {
	// ToProto marshals Dhcpv6ServerOptionsMessageType to protobuf object *otg.Dhcpv6ServerOptionsMessageType
	ToProto() (*otg.Dhcpv6ServerOptionsMessageType, error)
	// ToPbText marshals Dhcpv6ServerOptionsMessageType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ServerOptionsMessageType to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ServerOptionsMessageType to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv6ServerOptionsMessageType struct {
	obj *dhcpv6ServerOptionsMessageType
}

type unMarshalDhcpv6ServerOptionsMessageType interface {
	// FromProto unmarshals Dhcpv6ServerOptionsMessageType from protobuf object *otg.Dhcpv6ServerOptionsMessageType
	FromProto(msg *otg.Dhcpv6ServerOptionsMessageType) (Dhcpv6ServerOptionsMessageType, error)
	// FromPbText unmarshals Dhcpv6ServerOptionsMessageType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ServerOptionsMessageType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ServerOptionsMessageType from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ServerOptionsMessageType) Marshal() marshalDhcpv6ServerOptionsMessageType {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ServerOptionsMessageType{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ServerOptionsMessageType) Unmarshal() unMarshalDhcpv6ServerOptionsMessageType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ServerOptionsMessageType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ServerOptionsMessageType) ToProto() (*otg.Dhcpv6ServerOptionsMessageType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ServerOptionsMessageType) FromProto(msg *otg.Dhcpv6ServerOptionsMessageType) (Dhcpv6ServerOptionsMessageType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ServerOptionsMessageType) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ServerOptionsMessageType) FromPbText(value string) error {
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

func (m *marshaldhcpv6ServerOptionsMessageType) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ServerOptionsMessageType) FromYaml(value string) error {
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

func (m *marshaldhcpv6ServerOptionsMessageType) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ServerOptionsMessageType) FromJson(value string) error {
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

func (obj *dhcpv6ServerOptionsMessageType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ServerOptionsMessageType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ServerOptionsMessageType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ServerOptionsMessageType) Clone() (Dhcpv6ServerOptionsMessageType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ServerOptionsMessageType()
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

// Dhcpv6ServerOptionsMessageType is the dhcpv6 server messages where the option will be included.
type Dhcpv6ServerOptionsMessageType interface {
	Validation
	// msg marshals Dhcpv6ServerOptionsMessageType to protobuf object *otg.Dhcpv6ServerOptionsMessageType
	// and doesn't set defaults
	msg() *otg.Dhcpv6ServerOptionsMessageType
	// setMsg unmarshals Dhcpv6ServerOptionsMessageType from protobuf object *otg.Dhcpv6ServerOptionsMessageType
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ServerOptionsMessageType) Dhcpv6ServerOptionsMessageType
	// provides marshal interface
	Marshal() marshalDhcpv6ServerOptionsMessageType
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ServerOptionsMessageType
	// validate validates Dhcpv6ServerOptionsMessageType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ServerOptionsMessageType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Dhcpv6ServerOptionsMessageTypeChoiceEnum, set in Dhcpv6ServerOptionsMessageType
	Choice() Dhcpv6ServerOptionsMessageTypeChoiceEnum
	// setChoice assigns Dhcpv6ServerOptionsMessageTypeChoiceEnum provided by user to Dhcpv6ServerOptionsMessageType
	setChoice(value Dhcpv6ServerOptionsMessageTypeChoiceEnum) Dhcpv6ServerOptionsMessageType
	// HasChoice checks if Choice has been set in Dhcpv6ServerOptionsMessageType
	HasChoice() bool
	// getter for Reply to set choice.
	Reply()
	// getter for ReConfigure to set choice.
	ReConfigure()
	// getter for Advertise to set choice.
	Advertise()
	// getter for ReConfigure to set choice.
	ReConfigure()
}

type Dhcpv6ServerOptionsMessageTypeChoiceEnum string

// Enum of Choice on Dhcpv6ServerOptionsMessageType
var Dhcpv6ServerOptionsMessageTypeChoice = struct {
	ADVERTISE    Dhcpv6ServerOptionsMessageTypeChoiceEnum
	REPLY        Dhcpv6ServerOptionsMessageTypeChoiceEnum
	RE_CONFIGURE Dhcpv6ServerOptionsMessageTypeChoiceEnum
}{
	ADVERTISE:    Dhcpv6ServerOptionsMessageTypeChoiceEnum("advertise"),
	REPLY:        Dhcpv6ServerOptionsMessageTypeChoiceEnum("reply"),
	RE_CONFIGURE: Dhcpv6ServerOptionsMessageTypeChoiceEnum("re_configure"),
}

func (obj *dhcpv6ServerOptionsMessageType) Choice() Dhcpv6ServerOptionsMessageTypeChoiceEnum {
	return Dhcpv6ServerOptionsMessageTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Reply to set choice
func (obj *dhcpv6ServerOptionsMessageType) Reply() {
	obj.setChoice(Dhcpv6ServerOptionsMessageTypeChoice.REPLY)
}

// getter for ReConfigure to set choice
func (obj *dhcpv6ServerOptionsMessageType) ReConfigure() {
	obj.setChoice(Dhcpv6ServerOptionsMessageTypeChoice.RE_CONFIGURE)
}

// getter for Advertise to set choice
func (obj *dhcpv6ServerOptionsMessageType) Advertise() {
	obj.setChoice(Dhcpv6ServerOptionsMessageTypeChoice.ADVERTISE)
}

// getter for ReConfigure to set choice
func (obj *dhcpv6ServerOptionsMessageType) ReConfigure() {
	obj.setChoice(Dhcpv6ServerOptionsMessageTypeChoice.RE_CONFIGURE)
}

// The server message name where the option is included, by default it is all.
// Choice returns a string
func (obj *dhcpv6ServerOptionsMessageType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *dhcpv6ServerOptionsMessageType) setChoice(value Dhcpv6ServerOptionsMessageTypeChoiceEnum) Dhcpv6ServerOptionsMessageType {
	intValue, ok := otg.Dhcpv6ServerOptionsMessageType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Dhcpv6ServerOptionsMessageTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Dhcpv6ServerOptionsMessageType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *dhcpv6ServerOptionsMessageType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *dhcpv6ServerOptionsMessageType) setDefault() {
	var choices_set int = 0
	var choice Dhcpv6ServerOptionsMessageTypeChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Dhcpv6ServerOptionsMessageTypeChoice.ADVERTISE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Dhcpv6ServerOptionsMessageType")
			}
		} else {
			intVal := otg.Dhcpv6ServerOptionsMessageType_Choice_Enum_value[string(choice)]
			enumValue := otg.Dhcpv6ServerOptionsMessageType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

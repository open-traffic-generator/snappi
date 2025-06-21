package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ClientOptionsMessageType *****
type dhcpv6ClientOptionsMessageType struct {
	validation
	obj          *otg.Dhcpv6ClientOptionsMessageType
	marshaller   marshalDhcpv6ClientOptionsMessageType
	unMarshaller unMarshalDhcpv6ClientOptionsMessageType
}

func NewDhcpv6ClientOptionsMessageType() Dhcpv6ClientOptionsMessageType {
	obj := dhcpv6ClientOptionsMessageType{obj: &otg.Dhcpv6ClientOptionsMessageType{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ClientOptionsMessageType) msg() *otg.Dhcpv6ClientOptionsMessageType {
	return obj.obj
}

func (obj *dhcpv6ClientOptionsMessageType) setMsg(msg *otg.Dhcpv6ClientOptionsMessageType) Dhcpv6ClientOptionsMessageType {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ClientOptionsMessageType struct {
	obj *dhcpv6ClientOptionsMessageType
}

type marshalDhcpv6ClientOptionsMessageType interface {
	// ToProto marshals Dhcpv6ClientOptionsMessageType to protobuf object *otg.Dhcpv6ClientOptionsMessageType
	ToProto() (*otg.Dhcpv6ClientOptionsMessageType, error)
	// ToPbText marshals Dhcpv6ClientOptionsMessageType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ClientOptionsMessageType to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ClientOptionsMessageType to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv6ClientOptionsMessageType struct {
	obj *dhcpv6ClientOptionsMessageType
}

type unMarshalDhcpv6ClientOptionsMessageType interface {
	// FromProto unmarshals Dhcpv6ClientOptionsMessageType from protobuf object *otg.Dhcpv6ClientOptionsMessageType
	FromProto(msg *otg.Dhcpv6ClientOptionsMessageType) (Dhcpv6ClientOptionsMessageType, error)
	// FromPbText unmarshals Dhcpv6ClientOptionsMessageType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ClientOptionsMessageType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ClientOptionsMessageType from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ClientOptionsMessageType) Marshal() marshalDhcpv6ClientOptionsMessageType {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ClientOptionsMessageType{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ClientOptionsMessageType) Unmarshal() unMarshalDhcpv6ClientOptionsMessageType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ClientOptionsMessageType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ClientOptionsMessageType) ToProto() (*otg.Dhcpv6ClientOptionsMessageType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ClientOptionsMessageType) FromProto(msg *otg.Dhcpv6ClientOptionsMessageType) (Dhcpv6ClientOptionsMessageType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ClientOptionsMessageType) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsMessageType) FromPbText(value string) error {
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

func (m *marshaldhcpv6ClientOptionsMessageType) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsMessageType) FromYaml(value string) error {
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

func (m *marshaldhcpv6ClientOptionsMessageType) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsMessageType) FromJson(value string) error {
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

func (obj *dhcpv6ClientOptionsMessageType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsMessageType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsMessageType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ClientOptionsMessageType) Clone() (Dhcpv6ClientOptionsMessageType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ClientOptionsMessageType()
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

// Dhcpv6ClientOptionsMessageType is the dhcpv6 client messages where the option will be included.
type Dhcpv6ClientOptionsMessageType interface {
	Validation
	// msg marshals Dhcpv6ClientOptionsMessageType to protobuf object *otg.Dhcpv6ClientOptionsMessageType
	// and doesn't set defaults
	msg() *otg.Dhcpv6ClientOptionsMessageType
	// setMsg unmarshals Dhcpv6ClientOptionsMessageType from protobuf object *otg.Dhcpv6ClientOptionsMessageType
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ClientOptionsMessageType) Dhcpv6ClientOptionsMessageType
	// provides marshal interface
	Marshal() marshalDhcpv6ClientOptionsMessageType
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ClientOptionsMessageType
	// validate validates Dhcpv6ClientOptionsMessageType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ClientOptionsMessageType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Dhcpv6ClientOptionsMessageTypeChoiceEnum, set in Dhcpv6ClientOptionsMessageType
	Choice() Dhcpv6ClientOptionsMessageTypeChoiceEnum
	// setChoice assigns Dhcpv6ClientOptionsMessageTypeChoiceEnum provided by user to Dhcpv6ClientOptionsMessageType
	setChoice(value Dhcpv6ClientOptionsMessageTypeChoiceEnum) Dhcpv6ClientOptionsMessageType
	// HasChoice checks if Choice has been set in Dhcpv6ClientOptionsMessageType
	HasChoice() bool
	// getter for InformRequest to set choice.
	InformRequest()
	// getter for Release to set choice.
	Release()
	// getter for Request to set choice.
	Request()
	// getter for Renew to set choice.
	Renew()
	// getter for Solicit to set choice.
	Solicit()
	// getter for Rebind to set choice.
	Rebind()
}

type Dhcpv6ClientOptionsMessageTypeChoiceEnum string

// Enum of Choice on Dhcpv6ClientOptionsMessageType
var Dhcpv6ClientOptionsMessageTypeChoice = struct {
	SOLICIT        Dhcpv6ClientOptionsMessageTypeChoiceEnum
	REQUEST        Dhcpv6ClientOptionsMessageTypeChoiceEnum
	INFORM_REQUEST Dhcpv6ClientOptionsMessageTypeChoiceEnum
	RELEASE        Dhcpv6ClientOptionsMessageTypeChoiceEnum
	RENEW          Dhcpv6ClientOptionsMessageTypeChoiceEnum
	REBIND         Dhcpv6ClientOptionsMessageTypeChoiceEnum
}{
	SOLICIT:        Dhcpv6ClientOptionsMessageTypeChoiceEnum("solicit"),
	REQUEST:        Dhcpv6ClientOptionsMessageTypeChoiceEnum("request"),
	INFORM_REQUEST: Dhcpv6ClientOptionsMessageTypeChoiceEnum("inform_request"),
	RELEASE:        Dhcpv6ClientOptionsMessageTypeChoiceEnum("release"),
	RENEW:          Dhcpv6ClientOptionsMessageTypeChoiceEnum("renew"),
	REBIND:         Dhcpv6ClientOptionsMessageTypeChoiceEnum("rebind"),
}

func (obj *dhcpv6ClientOptionsMessageType) Choice() Dhcpv6ClientOptionsMessageTypeChoiceEnum {
	return Dhcpv6ClientOptionsMessageTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for InformRequest to set choice
func (obj *dhcpv6ClientOptionsMessageType) InformRequest() {
	obj.setChoice(Dhcpv6ClientOptionsMessageTypeChoice.INFORM_REQUEST)
}

// getter for Release to set choice
func (obj *dhcpv6ClientOptionsMessageType) Release() {
	obj.setChoice(Dhcpv6ClientOptionsMessageTypeChoice.RELEASE)
}

// getter for Request to set choice
func (obj *dhcpv6ClientOptionsMessageType) Request() {
	obj.setChoice(Dhcpv6ClientOptionsMessageTypeChoice.REQUEST)
}

// getter for Renew to set choice
func (obj *dhcpv6ClientOptionsMessageType) Renew() {
	obj.setChoice(Dhcpv6ClientOptionsMessageTypeChoice.RENEW)
}

// getter for Solicit to set choice
func (obj *dhcpv6ClientOptionsMessageType) Solicit() {
	obj.setChoice(Dhcpv6ClientOptionsMessageTypeChoice.SOLICIT)
}

// getter for Rebind to set choice
func (obj *dhcpv6ClientOptionsMessageType) Rebind() {
	obj.setChoice(Dhcpv6ClientOptionsMessageTypeChoice.REBIND)
}

// The client message name where the option is included, by default it is all.
// Choice returns a string
func (obj *dhcpv6ClientOptionsMessageType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *dhcpv6ClientOptionsMessageType) setChoice(value Dhcpv6ClientOptionsMessageTypeChoiceEnum) Dhcpv6ClientOptionsMessageType {
	intValue, ok := otg.Dhcpv6ClientOptionsMessageType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Dhcpv6ClientOptionsMessageTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Dhcpv6ClientOptionsMessageType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *dhcpv6ClientOptionsMessageType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *dhcpv6ClientOptionsMessageType) setDefault() {
	var choices_set int = 0
	var choice Dhcpv6ClientOptionsMessageTypeChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Dhcpv6ClientOptionsMessageTypeChoice.SOLICIT)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Dhcpv6ClientOptionsMessageType")
			}
		} else {
			intVal := otg.Dhcpv6ClientOptionsMessageType_Choice_Enum_value[string(choice)]
			enumValue := otg.Dhcpv6ClientOptionsMessageType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ClientOptionsDuidUuidVersion *****
type dhcpv6ClientOptionsDuidUuidVersion struct {
	validation
	obj          *otg.Dhcpv6ClientOptionsDuidUuidVersion
	marshaller   marshalDhcpv6ClientOptionsDuidUuidVersion
	unMarshaller unMarshalDhcpv6ClientOptionsDuidUuidVersion
}

func NewDhcpv6ClientOptionsDuidUuidVersion() Dhcpv6ClientOptionsDuidUuidVersion {
	obj := dhcpv6ClientOptionsDuidUuidVersion{obj: &otg.Dhcpv6ClientOptionsDuidUuidVersion{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ClientOptionsDuidUuidVersion) msg() *otg.Dhcpv6ClientOptionsDuidUuidVersion {
	return obj.obj
}

func (obj *dhcpv6ClientOptionsDuidUuidVersion) setMsg(msg *otg.Dhcpv6ClientOptionsDuidUuidVersion) Dhcpv6ClientOptionsDuidUuidVersion {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ClientOptionsDuidUuidVersion struct {
	obj *dhcpv6ClientOptionsDuidUuidVersion
}

type marshalDhcpv6ClientOptionsDuidUuidVersion interface {
	// ToProto marshals Dhcpv6ClientOptionsDuidUuidVersion to protobuf object *otg.Dhcpv6ClientOptionsDuidUuidVersion
	ToProto() (*otg.Dhcpv6ClientOptionsDuidUuidVersion, error)
	// ToPbText marshals Dhcpv6ClientOptionsDuidUuidVersion to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ClientOptionsDuidUuidVersion to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ClientOptionsDuidUuidVersion to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv6ClientOptionsDuidUuidVersion struct {
	obj *dhcpv6ClientOptionsDuidUuidVersion
}

type unMarshalDhcpv6ClientOptionsDuidUuidVersion interface {
	// FromProto unmarshals Dhcpv6ClientOptionsDuidUuidVersion from protobuf object *otg.Dhcpv6ClientOptionsDuidUuidVersion
	FromProto(msg *otg.Dhcpv6ClientOptionsDuidUuidVersion) (Dhcpv6ClientOptionsDuidUuidVersion, error)
	// FromPbText unmarshals Dhcpv6ClientOptionsDuidUuidVersion from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ClientOptionsDuidUuidVersion from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ClientOptionsDuidUuidVersion from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ClientOptionsDuidUuidVersion) Marshal() marshalDhcpv6ClientOptionsDuidUuidVersion {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ClientOptionsDuidUuidVersion{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ClientOptionsDuidUuidVersion) Unmarshal() unMarshalDhcpv6ClientOptionsDuidUuidVersion {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ClientOptionsDuidUuidVersion{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ClientOptionsDuidUuidVersion) ToProto() (*otg.Dhcpv6ClientOptionsDuidUuidVersion, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ClientOptionsDuidUuidVersion) FromProto(msg *otg.Dhcpv6ClientOptionsDuidUuidVersion) (Dhcpv6ClientOptionsDuidUuidVersion, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ClientOptionsDuidUuidVersion) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsDuidUuidVersion) FromPbText(value string) error {
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

func (m *marshaldhcpv6ClientOptionsDuidUuidVersion) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsDuidUuidVersion) FromYaml(value string) error {
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

func (m *marshaldhcpv6ClientOptionsDuidUuidVersion) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsDuidUuidVersion) FromJson(value string) error {
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

func (obj *dhcpv6ClientOptionsDuidUuidVersion) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsDuidUuidVersion) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsDuidUuidVersion) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ClientOptionsDuidUuidVersion) Clone() (Dhcpv6ClientOptionsDuidUuidVersion, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ClientOptionsDuidUuidVersion()
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

// Dhcpv6ClientOptionsDuidUuidVersion is the version number is in the most significant 4 bits of the timestamp (bits 4 through 7 of the time_hi_and_version field).
type Dhcpv6ClientOptionsDuidUuidVersion interface {
	Validation
	// msg marshals Dhcpv6ClientOptionsDuidUuidVersion to protobuf object *otg.Dhcpv6ClientOptionsDuidUuidVersion
	// and doesn't set defaults
	msg() *otg.Dhcpv6ClientOptionsDuidUuidVersion
	// setMsg unmarshals Dhcpv6ClientOptionsDuidUuidVersion from protobuf object *otg.Dhcpv6ClientOptionsDuidUuidVersion
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ClientOptionsDuidUuidVersion) Dhcpv6ClientOptionsDuidUuidVersion
	// provides marshal interface
	Marshal() marshalDhcpv6ClientOptionsDuidUuidVersion
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ClientOptionsDuidUuidVersion
	// validate validates Dhcpv6ClientOptionsDuidUuidVersion
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ClientOptionsDuidUuidVersion, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Dhcpv6ClientOptionsDuidUuidVersionChoiceEnum, set in Dhcpv6ClientOptionsDuidUuidVersion
	Choice() Dhcpv6ClientOptionsDuidUuidVersionChoiceEnum
	// setChoice assigns Dhcpv6ClientOptionsDuidUuidVersionChoiceEnum provided by user to Dhcpv6ClientOptionsDuidUuidVersion
	setChoice(value Dhcpv6ClientOptionsDuidUuidVersionChoiceEnum) Dhcpv6ClientOptionsDuidUuidVersion
	// HasChoice checks if Choice has been set in Dhcpv6ClientOptionsDuidUuidVersion
	HasChoice() bool
	// getter for V_2 to set choice.
	V_2()
	// getter for V_4 to set choice.
	V_4()
	// getter for V_3 to set choice.
	V_3()
	// getter for V_1 to set choice.
	V_1()
	// getter for V_5 to set choice.
	V_5()
}

type Dhcpv6ClientOptionsDuidUuidVersionChoiceEnum string

// Enum of Choice on Dhcpv6ClientOptionsDuidUuidVersion
var Dhcpv6ClientOptionsDuidUuidVersionChoice = struct {
	V_1 Dhcpv6ClientOptionsDuidUuidVersionChoiceEnum
	V_2 Dhcpv6ClientOptionsDuidUuidVersionChoiceEnum
	V_3 Dhcpv6ClientOptionsDuidUuidVersionChoiceEnum
	V_4 Dhcpv6ClientOptionsDuidUuidVersionChoiceEnum
	V_5 Dhcpv6ClientOptionsDuidUuidVersionChoiceEnum
}{
	V_1: Dhcpv6ClientOptionsDuidUuidVersionChoiceEnum("v_1"),
	V_2: Dhcpv6ClientOptionsDuidUuidVersionChoiceEnum("v_2"),
	V_3: Dhcpv6ClientOptionsDuidUuidVersionChoiceEnum("v_3"),
	V_4: Dhcpv6ClientOptionsDuidUuidVersionChoiceEnum("v_4"),
	V_5: Dhcpv6ClientOptionsDuidUuidVersionChoiceEnum("v_5"),
}

func (obj *dhcpv6ClientOptionsDuidUuidVersion) Choice() Dhcpv6ClientOptionsDuidUuidVersionChoiceEnum {
	return Dhcpv6ClientOptionsDuidUuidVersionChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for V_2 to set choice
func (obj *dhcpv6ClientOptionsDuidUuidVersion) V_2() {
	obj.setChoice(Dhcpv6ClientOptionsDuidUuidVersionChoice.V_2)
}

// getter for V_4 to set choice
func (obj *dhcpv6ClientOptionsDuidUuidVersion) V_4() {
	obj.setChoice(Dhcpv6ClientOptionsDuidUuidVersionChoice.V_4)
}

// getter for V_3 to set choice
func (obj *dhcpv6ClientOptionsDuidUuidVersion) V_3() {
	obj.setChoice(Dhcpv6ClientOptionsDuidUuidVersionChoice.V_3)
}

// getter for V_1 to set choice
func (obj *dhcpv6ClientOptionsDuidUuidVersion) V_1() {
	obj.setChoice(Dhcpv6ClientOptionsDuidUuidVersionChoice.V_1)
}

// getter for V_5 to set choice
func (obj *dhcpv6ClientOptionsDuidUuidVersion) V_5() {
	obj.setChoice(Dhcpv6ClientOptionsDuidUuidVersionChoice.V_5)
}

// The version values are from 1 to 5 in the most significant 4 bits of the timestamp (bits 4 through 7 of the  time_hi_and_version field).
// Choice returns a string
func (obj *dhcpv6ClientOptionsDuidUuidVersion) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *dhcpv6ClientOptionsDuidUuidVersion) setChoice(value Dhcpv6ClientOptionsDuidUuidVersionChoiceEnum) Dhcpv6ClientOptionsDuidUuidVersion {
	intValue, ok := otg.Dhcpv6ClientOptionsDuidUuidVersion_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Dhcpv6ClientOptionsDuidUuidVersionChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Dhcpv6ClientOptionsDuidUuidVersion_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *dhcpv6ClientOptionsDuidUuidVersion) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *dhcpv6ClientOptionsDuidUuidVersion) setDefault() {
	var choices_set int = 0
	var choice Dhcpv6ClientOptionsDuidUuidVersionChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Dhcpv6ClientOptionsDuidUuidVersionChoice.V_1)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Dhcpv6ClientOptionsDuidUuidVersion")
			}
		} else {
			intVal := otg.Dhcpv6ClientOptionsDuidUuidVersion_Choice_Enum_value[string(choice)]
			enumValue := otg.Dhcpv6ClientOptionsDuidUuidVersion_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ClientOptionsDuidUuidVariant *****
type dhcpv6ClientOptionsDuidUuidVariant struct {
	validation
	obj          *otg.Dhcpv6ClientOptionsDuidUuidVariant
	marshaller   marshalDhcpv6ClientOptionsDuidUuidVariant
	unMarshaller unMarshalDhcpv6ClientOptionsDuidUuidVariant
}

func NewDhcpv6ClientOptionsDuidUuidVariant() Dhcpv6ClientOptionsDuidUuidVariant {
	obj := dhcpv6ClientOptionsDuidUuidVariant{obj: &otg.Dhcpv6ClientOptionsDuidUuidVariant{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ClientOptionsDuidUuidVariant) msg() *otg.Dhcpv6ClientOptionsDuidUuidVariant {
	return obj.obj
}

func (obj *dhcpv6ClientOptionsDuidUuidVariant) setMsg(msg *otg.Dhcpv6ClientOptionsDuidUuidVariant) Dhcpv6ClientOptionsDuidUuidVariant {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ClientOptionsDuidUuidVariant struct {
	obj *dhcpv6ClientOptionsDuidUuidVariant
}

type marshalDhcpv6ClientOptionsDuidUuidVariant interface {
	// ToProto marshals Dhcpv6ClientOptionsDuidUuidVariant to protobuf object *otg.Dhcpv6ClientOptionsDuidUuidVariant
	ToProto() (*otg.Dhcpv6ClientOptionsDuidUuidVariant, error)
	// ToPbText marshals Dhcpv6ClientOptionsDuidUuidVariant to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ClientOptionsDuidUuidVariant to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ClientOptionsDuidUuidVariant to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv6ClientOptionsDuidUuidVariant to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv6ClientOptionsDuidUuidVariant struct {
	obj *dhcpv6ClientOptionsDuidUuidVariant
}

type unMarshalDhcpv6ClientOptionsDuidUuidVariant interface {
	// FromProto unmarshals Dhcpv6ClientOptionsDuidUuidVariant from protobuf object *otg.Dhcpv6ClientOptionsDuidUuidVariant
	FromProto(msg *otg.Dhcpv6ClientOptionsDuidUuidVariant) (Dhcpv6ClientOptionsDuidUuidVariant, error)
	// FromPbText unmarshals Dhcpv6ClientOptionsDuidUuidVariant from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ClientOptionsDuidUuidVariant from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ClientOptionsDuidUuidVariant from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ClientOptionsDuidUuidVariant) Marshal() marshalDhcpv6ClientOptionsDuidUuidVariant {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ClientOptionsDuidUuidVariant{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ClientOptionsDuidUuidVariant) Unmarshal() unMarshalDhcpv6ClientOptionsDuidUuidVariant {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ClientOptionsDuidUuidVariant{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ClientOptionsDuidUuidVariant) ToProto() (*otg.Dhcpv6ClientOptionsDuidUuidVariant, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ClientOptionsDuidUuidVariant) FromProto(msg *otg.Dhcpv6ClientOptionsDuidUuidVariant) (Dhcpv6ClientOptionsDuidUuidVariant, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ClientOptionsDuidUuidVariant) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsDuidUuidVariant) FromPbText(value string) error {
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

func (m *marshaldhcpv6ClientOptionsDuidUuidVariant) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsDuidUuidVariant) FromYaml(value string) error {
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

func (m *marshaldhcpv6ClientOptionsDuidUuidVariant) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv6ClientOptionsDuidUuidVariant) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsDuidUuidVariant) FromJson(value string) error {
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

func (obj *dhcpv6ClientOptionsDuidUuidVariant) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsDuidUuidVariant) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsDuidUuidVariant) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ClientOptionsDuidUuidVariant) Clone() (Dhcpv6ClientOptionsDuidUuidVariant, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ClientOptionsDuidUuidVariant()
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

// Dhcpv6ClientOptionsDuidUuidVariant is the variant field determines the layout of the UUID.  That is, the interpretation of all other bits in the  UUID depends on the setting of the bits in the variant field).
type Dhcpv6ClientOptionsDuidUuidVariant interface {
	Validation
	// msg marshals Dhcpv6ClientOptionsDuidUuidVariant to protobuf object *otg.Dhcpv6ClientOptionsDuidUuidVariant
	// and doesn't set defaults
	msg() *otg.Dhcpv6ClientOptionsDuidUuidVariant
	// setMsg unmarshals Dhcpv6ClientOptionsDuidUuidVariant from protobuf object *otg.Dhcpv6ClientOptionsDuidUuidVariant
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ClientOptionsDuidUuidVariant) Dhcpv6ClientOptionsDuidUuidVariant
	// provides marshal interface
	Marshal() marshalDhcpv6ClientOptionsDuidUuidVariant
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ClientOptionsDuidUuidVariant
	// validate validates Dhcpv6ClientOptionsDuidUuidVariant
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ClientOptionsDuidUuidVariant, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Dhcpv6ClientOptionsDuidUuidVariantChoiceEnum, set in Dhcpv6ClientOptionsDuidUuidVariant
	Choice() Dhcpv6ClientOptionsDuidUuidVariantChoiceEnum
	// setChoice assigns Dhcpv6ClientOptionsDuidUuidVariantChoiceEnum provided by user to Dhcpv6ClientOptionsDuidUuidVariant
	setChoice(value Dhcpv6ClientOptionsDuidUuidVariantChoiceEnum) Dhcpv6ClientOptionsDuidUuidVariant
	// HasChoice checks if Choice has been set in Dhcpv6ClientOptionsDuidUuidVariant
	HasChoice() bool
	// getter for Ncs to set choice.
	Ncs()
	// getter for Dce to set choice.
	Dce()
	// getter for VarReserved to set choice.
	VarReserved()
	// getter for Guid to set choice.
	Guid()
}

type Dhcpv6ClientOptionsDuidUuidVariantChoiceEnum string

// Enum of Choice on Dhcpv6ClientOptionsDuidUuidVariant
var Dhcpv6ClientOptionsDuidUuidVariantChoice = struct {
	NCS          Dhcpv6ClientOptionsDuidUuidVariantChoiceEnum
	DCE          Dhcpv6ClientOptionsDuidUuidVariantChoiceEnum
	GUID         Dhcpv6ClientOptionsDuidUuidVariantChoiceEnum
	VAR_RESERVED Dhcpv6ClientOptionsDuidUuidVariantChoiceEnum
}{
	NCS:          Dhcpv6ClientOptionsDuidUuidVariantChoiceEnum("ncs"),
	DCE:          Dhcpv6ClientOptionsDuidUuidVariantChoiceEnum("dce"),
	GUID:         Dhcpv6ClientOptionsDuidUuidVariantChoiceEnum("guid"),
	VAR_RESERVED: Dhcpv6ClientOptionsDuidUuidVariantChoiceEnum("var_reserved"),
}

func (obj *dhcpv6ClientOptionsDuidUuidVariant) Choice() Dhcpv6ClientOptionsDuidUuidVariantChoiceEnum {
	return Dhcpv6ClientOptionsDuidUuidVariantChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Ncs to set choice
func (obj *dhcpv6ClientOptionsDuidUuidVariant) Ncs() {
	obj.setChoice(Dhcpv6ClientOptionsDuidUuidVariantChoice.NCS)
}

// getter for Dce to set choice
func (obj *dhcpv6ClientOptionsDuidUuidVariant) Dce() {
	obj.setChoice(Dhcpv6ClientOptionsDuidUuidVariantChoice.DCE)
}

// getter for VarReserved to set choice
func (obj *dhcpv6ClientOptionsDuidUuidVariant) VarReserved() {
	obj.setChoice(Dhcpv6ClientOptionsDuidUuidVariantChoice.VAR_RESERVED)
}

// getter for Guid to set choice
func (obj *dhcpv6ClientOptionsDuidUuidVariant) Guid() {
	obj.setChoice(Dhcpv6ClientOptionsDuidUuidVariantChoice.GUID)
}

// The current variants are ncs, dce,microsoft guid and reserved.
// Choice returns a string
func (obj *dhcpv6ClientOptionsDuidUuidVariant) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *dhcpv6ClientOptionsDuidUuidVariant) setChoice(value Dhcpv6ClientOptionsDuidUuidVariantChoiceEnum) Dhcpv6ClientOptionsDuidUuidVariant {
	intValue, ok := otg.Dhcpv6ClientOptionsDuidUuidVariant_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Dhcpv6ClientOptionsDuidUuidVariantChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Dhcpv6ClientOptionsDuidUuidVariant_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *dhcpv6ClientOptionsDuidUuidVariant) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *dhcpv6ClientOptionsDuidUuidVariant) setDefault() {
	var choices_set int = 0
	var choice Dhcpv6ClientOptionsDuidUuidVariantChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Dhcpv6ClientOptionsDuidUuidVariantChoice.NCS)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Dhcpv6ClientOptionsDuidUuidVariant")
			}
		} else {
			intVal := otg.Dhcpv6ClientOptionsDuidUuidVariant_Choice_Enum_value[string(choice)]
			enumValue := otg.Dhcpv6ClientOptionsDuidUuidVariant_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

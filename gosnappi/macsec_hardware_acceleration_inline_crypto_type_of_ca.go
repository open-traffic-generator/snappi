package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecHardwareAccelerationInlineCryptoTypeOfCa *****
type macsecHardwareAccelerationInlineCryptoTypeOfCa struct {
	validation
	obj           *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCa
	marshaller    marshalMacsecHardwareAccelerationInlineCryptoTypeOfCa
	unMarshaller  unMarshalMacsecHardwareAccelerationInlineCryptoTypeOfCa
	groupCaHolder MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
}

func NewMacsecHardwareAccelerationInlineCryptoTypeOfCa() MacsecHardwareAccelerationInlineCryptoTypeOfCa {
	obj := macsecHardwareAccelerationInlineCryptoTypeOfCa{obj: &otg.MacsecHardwareAccelerationInlineCryptoTypeOfCa{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCa) msg() *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCa {
	return obj.obj
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCa) setMsg(msg *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCa) MacsecHardwareAccelerationInlineCryptoTypeOfCa {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecHardwareAccelerationInlineCryptoTypeOfCa struct {
	obj *macsecHardwareAccelerationInlineCryptoTypeOfCa
}

type marshalMacsecHardwareAccelerationInlineCryptoTypeOfCa interface {
	// ToProto marshals MacsecHardwareAccelerationInlineCryptoTypeOfCa to protobuf object *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCa
	ToProto() (*otg.MacsecHardwareAccelerationInlineCryptoTypeOfCa, error)
	// ToPbText marshals MacsecHardwareAccelerationInlineCryptoTypeOfCa to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecHardwareAccelerationInlineCryptoTypeOfCa to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecHardwareAccelerationInlineCryptoTypeOfCa to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecHardwareAccelerationInlineCryptoTypeOfCa struct {
	obj *macsecHardwareAccelerationInlineCryptoTypeOfCa
}

type unMarshalMacsecHardwareAccelerationInlineCryptoTypeOfCa interface {
	// FromProto unmarshals MacsecHardwareAccelerationInlineCryptoTypeOfCa from protobuf object *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCa
	FromProto(msg *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCa) (MacsecHardwareAccelerationInlineCryptoTypeOfCa, error)
	// FromPbText unmarshals MacsecHardwareAccelerationInlineCryptoTypeOfCa from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecHardwareAccelerationInlineCryptoTypeOfCa from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecHardwareAccelerationInlineCryptoTypeOfCa from JSON text
	FromJson(value string) error
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCa) Marshal() marshalMacsecHardwareAccelerationInlineCryptoTypeOfCa {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecHardwareAccelerationInlineCryptoTypeOfCa{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCa) Unmarshal() unMarshalMacsecHardwareAccelerationInlineCryptoTypeOfCa {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecHardwareAccelerationInlineCryptoTypeOfCa{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecHardwareAccelerationInlineCryptoTypeOfCa) ToProto() (*otg.MacsecHardwareAccelerationInlineCryptoTypeOfCa, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecHardwareAccelerationInlineCryptoTypeOfCa) FromProto(msg *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCa) (MacsecHardwareAccelerationInlineCryptoTypeOfCa, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecHardwareAccelerationInlineCryptoTypeOfCa) ToPbText() (string, error) {
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

func (m *unMarshalmacsecHardwareAccelerationInlineCryptoTypeOfCa) FromPbText(value string) error {
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

func (m *marshalmacsecHardwareAccelerationInlineCryptoTypeOfCa) ToYaml() (string, error) {
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

func (m *unMarshalmacsecHardwareAccelerationInlineCryptoTypeOfCa) FromYaml(value string) error {
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

func (m *marshalmacsecHardwareAccelerationInlineCryptoTypeOfCa) ToJson() (string, error) {
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

func (m *unMarshalmacsecHardwareAccelerationInlineCryptoTypeOfCa) FromJson(value string) error {
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

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCa) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCa) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCa) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCa) Clone() (MacsecHardwareAccelerationInlineCryptoTypeOfCa, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecHardwareAccelerationInlineCryptoTypeOfCa()
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

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCa) setNil() {
	obj.groupCaHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecHardwareAccelerationInlineCryptoTypeOfCa is type of connectivity association(CA).
type MacsecHardwareAccelerationInlineCryptoTypeOfCa interface {
	Validation
	// msg marshals MacsecHardwareAccelerationInlineCryptoTypeOfCa to protobuf object *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCa
	// and doesn't set defaults
	msg() *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCa
	// setMsg unmarshals MacsecHardwareAccelerationInlineCryptoTypeOfCa from protobuf object *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCa
	// and doesn't set defaults
	setMsg(*otg.MacsecHardwareAccelerationInlineCryptoTypeOfCa) MacsecHardwareAccelerationInlineCryptoTypeOfCa
	// provides marshal interface
	Marshal() marshalMacsecHardwareAccelerationInlineCryptoTypeOfCa
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecHardwareAccelerationInlineCryptoTypeOfCa
	// validate validates MacsecHardwareAccelerationInlineCryptoTypeOfCa
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecHardwareAccelerationInlineCryptoTypeOfCa, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MacsecHardwareAccelerationInlineCryptoTypeOfCaChoiceEnum, set in MacsecHardwareAccelerationInlineCryptoTypeOfCa
	Choice() MacsecHardwareAccelerationInlineCryptoTypeOfCaChoiceEnum
	// setChoice assigns MacsecHardwareAccelerationInlineCryptoTypeOfCaChoiceEnum provided by user to MacsecHardwareAccelerationInlineCryptoTypeOfCa
	setChoice(value MacsecHardwareAccelerationInlineCryptoTypeOfCaChoiceEnum) MacsecHardwareAccelerationInlineCryptoTypeOfCa
	// HasChoice checks if Choice has been set in MacsecHardwareAccelerationInlineCryptoTypeOfCa
	HasChoice() bool
	// getter for PairwiseCa to set choice.
	PairwiseCa()
	// GroupCa returns MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa, set in MacsecHardwareAccelerationInlineCryptoTypeOfCa.
	// MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa is group onnectivity association(CA).
	GroupCa() MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
	// SetGroupCa assigns MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa provided by user to MacsecHardwareAccelerationInlineCryptoTypeOfCa.
	// MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa is group onnectivity association(CA).
	SetGroupCa(value MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) MacsecHardwareAccelerationInlineCryptoTypeOfCa
	// HasGroupCa checks if GroupCa has been set in MacsecHardwareAccelerationInlineCryptoTypeOfCa
	HasGroupCa() bool
	setNil()
}

type MacsecHardwareAccelerationInlineCryptoTypeOfCaChoiceEnum string

// Enum of Choice on MacsecHardwareAccelerationInlineCryptoTypeOfCa
var MacsecHardwareAccelerationInlineCryptoTypeOfCaChoice = struct {
	PAIRWISE_CA MacsecHardwareAccelerationInlineCryptoTypeOfCaChoiceEnum
	GROUP_CA    MacsecHardwareAccelerationInlineCryptoTypeOfCaChoiceEnum
}{
	PAIRWISE_CA: MacsecHardwareAccelerationInlineCryptoTypeOfCaChoiceEnum("pairwise_ca"),
	GROUP_CA:    MacsecHardwareAccelerationInlineCryptoTypeOfCaChoiceEnum("group_ca"),
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCa) Choice() MacsecHardwareAccelerationInlineCryptoTypeOfCaChoiceEnum {
	return MacsecHardwareAccelerationInlineCryptoTypeOfCaChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for PairwiseCa to set choice
func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCa) PairwiseCa() {
	obj.setChoice(MacsecHardwareAccelerationInlineCryptoTypeOfCaChoice.PAIRWISE_CA)
}

// Type of connectivity association(CA).
// Choice returns a string
func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCa) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCa) setChoice(value MacsecHardwareAccelerationInlineCryptoTypeOfCaChoiceEnum) MacsecHardwareAccelerationInlineCryptoTypeOfCa {
	intValue, ok := otg.MacsecHardwareAccelerationInlineCryptoTypeOfCa_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecHardwareAccelerationInlineCryptoTypeOfCaChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecHardwareAccelerationInlineCryptoTypeOfCa_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.GroupCa = nil
	obj.groupCaHolder = nil

	if value == MacsecHardwareAccelerationInlineCryptoTypeOfCaChoice.GROUP_CA {
		obj.obj.GroupCa = NewMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa().msg()
	}

	return obj
}

// description is TBD
// GroupCa returns a MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCa) GroupCa() MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa {
	if obj.obj.GroupCa == nil {
		obj.setChoice(MacsecHardwareAccelerationInlineCryptoTypeOfCaChoice.GROUP_CA)
	}
	if obj.groupCaHolder == nil {
		obj.groupCaHolder = &macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa{obj: obj.obj.GroupCa}
	}
	return obj.groupCaHolder
}

// description is TBD
// GroupCa returns a MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCa) HasGroupCa() bool {
	return obj.obj.GroupCa != nil
}

// description is TBD
// SetGroupCa sets the MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa value in the MacsecHardwareAccelerationInlineCryptoTypeOfCa object
func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCa) SetGroupCa(value MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) MacsecHardwareAccelerationInlineCryptoTypeOfCa {
	obj.setChoice(MacsecHardwareAccelerationInlineCryptoTypeOfCaChoice.GROUP_CA)
	obj.groupCaHolder = nil
	obj.obj.GroupCa = value.msg()

	return obj
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.GroupCa != nil {

		obj.GroupCa().validateObj(vObj, set_default)
	}

}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCa) setDefault() {
	var choices_set int = 0
	var choice MacsecHardwareAccelerationInlineCryptoTypeOfCaChoiceEnum

	if obj.obj.GroupCa != nil {
		choices_set += 1
		choice = MacsecHardwareAccelerationInlineCryptoTypeOfCaChoice.GROUP_CA
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MacsecHardwareAccelerationInlineCryptoTypeOfCaChoice.PAIRWISE_CA)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MacsecHardwareAccelerationInlineCryptoTypeOfCa")
			}
		} else {
			intVal := otg.MacsecHardwareAccelerationInlineCryptoTypeOfCa_Choice_Enum_value[string(choice)]
			enumValue := otg.MacsecHardwareAccelerationInlineCryptoTypeOfCa_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

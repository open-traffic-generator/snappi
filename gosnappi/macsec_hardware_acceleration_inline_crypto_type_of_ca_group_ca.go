package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa *****
type macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa struct {
	validation
	obj                   *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
	marshaller            marshalMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
	unMarshaller          unMarshalMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
	multipleDutTxScHolder MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc
}

func NewMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa() MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa {
	obj := macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa{obj: &otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) msg() *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa {
	return obj.obj
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) setMsg(msg *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa struct {
	obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
}

type marshalMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa interface {
	// ToProto marshals MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa to protobuf object *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
	ToProto() (*otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa, error)
	// ToPbText marshals MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa struct {
	obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
}

type unMarshalMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa interface {
	// FromProto unmarshals MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa from protobuf object *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
	FromProto(msg *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) (MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa, error)
	// FromPbText unmarshals MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa from JSON text
	FromJson(value string) error
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) Marshal() marshalMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) Unmarshal() unMarshalMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) ToProto() (*otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) FromProto(msg *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) (MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) ToPbText() (string, error) {
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

func (m *unMarshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) FromPbText(value string) error {
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

func (m *marshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) ToYaml() (string, error) {
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

func (m *unMarshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) FromYaml(value string) error {
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

func (m *marshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) ToJson() (string, error) {
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

func (m *unMarshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) FromJson(value string) error {
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

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) Clone() (MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa()
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

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) setNil() {
	obj.multipleDutTxScHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa is group onnectivity association(CA).
type MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa interface {
	Validation
	// msg marshals MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa to protobuf object *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
	// and doesn't set defaults
	msg() *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
	// setMsg unmarshals MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa from protobuf object *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
	// and doesn't set defaults
	setMsg(*otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
	// provides marshal interface
	Marshal() marshalMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
	// validate validates MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaChoiceEnum, set in MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
	Choice() MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaChoiceEnum
	// setChoice assigns MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaChoiceEnum provided by user to MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
	setChoice(value MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaChoiceEnum) MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
	// HasChoice checks if Choice has been set in MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
	HasChoice() bool
	// getter for SingleDutTxSc to set choice.
	SingleDutTxSc()
	// MultipleDutTxSc returns MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc, set in MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa.
	// MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc is maximum possible DUT Tx secure channel(SC) count per group connectivity association(CA).
	MultipleDutTxSc() MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc
	// SetMultipleDutTxSc assigns MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc provided by user to MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa.
	// MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc is maximum possible DUT Tx secure channel(SC) count per group connectivity association(CA).
	SetMultipleDutTxSc(value MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
	// HasMultipleDutTxSc checks if MultipleDutTxSc has been set in MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
	HasMultipleDutTxSc() bool
	setNil()
}

type MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaChoiceEnum string

// Enum of Choice on MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa
var MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaChoice = struct {
	SINGLE_DUT_TX_SC   MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaChoiceEnum
	MULTIPLE_DUT_TX_SC MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaChoiceEnum
}{
	SINGLE_DUT_TX_SC:   MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaChoiceEnum("single_dut_tx_sc"),
	MULTIPLE_DUT_TX_SC: MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaChoiceEnum("multiple_dut_tx_sc"),
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) Choice() MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaChoiceEnum {
	return MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for SingleDutTxSc to set choice
func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) SingleDutTxSc() {
	obj.setChoice(MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaChoice.SINGLE_DUT_TX_SC)
}

// DUT Tx secure channel(SC) count in group connectivity association(CA).
// Choice returns a string
func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) setChoice(value MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaChoiceEnum) MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa {
	intValue, ok := otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.MultipleDutTxSc = nil
	obj.multipleDutTxScHolder = nil

	if value == MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaChoice.MULTIPLE_DUT_TX_SC {
		obj.obj.MultipleDutTxSc = NewMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc().msg()
	}

	return obj
}

// description is TBD
// MultipleDutTxSc returns a MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc
func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) MultipleDutTxSc() MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc {
	if obj.obj.MultipleDutTxSc == nil {
		obj.setChoice(MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaChoice.MULTIPLE_DUT_TX_SC)
	}
	if obj.multipleDutTxScHolder == nil {
		obj.multipleDutTxScHolder = &macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc{obj: obj.obj.MultipleDutTxSc}
	}
	return obj.multipleDutTxScHolder
}

// description is TBD
// MultipleDutTxSc returns a MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc
func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) HasMultipleDutTxSc() bool {
	return obj.obj.MultipleDutTxSc != nil
}

// description is TBD
// SetMultipleDutTxSc sets the MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc value in the MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa object
func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) SetMultipleDutTxSc(value MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa {
	obj.setChoice(MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaChoice.MULTIPLE_DUT_TX_SC)
	obj.multipleDutTxScHolder = nil
	obj.obj.MultipleDutTxSc = value.msg()

	return obj
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MultipleDutTxSc != nil {

		obj.MultipleDutTxSc().validateObj(vObj, set_default)
	}

}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa) setDefault() {
	var choices_set int = 0
	var choice MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaChoiceEnum

	if obj.obj.MultipleDutTxSc != nil {
		choices_set += 1
		choice = MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaChoice.MULTIPLE_DUT_TX_SC
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaChoice.MULTIPLE_DUT_TX_SC)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa")
			}
		} else {
			intVal := otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa_Choice_Enum_value[string(choice)]
			enumValue := otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCa_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

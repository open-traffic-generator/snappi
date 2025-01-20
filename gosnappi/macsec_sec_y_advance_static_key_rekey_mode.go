package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYAdvanceStaticKeyRekeyMode *****
type macsecSecYAdvanceStaticKeyRekeyMode struct {
	validation
	obj          *otg.MacsecSecYAdvanceStaticKeyRekeyMode
	marshaller   marshalMacsecSecYAdvanceStaticKeyRekeyMode
	unMarshaller unMarshalMacsecSecYAdvanceStaticKeyRekeyMode
}

func NewMacsecSecYAdvanceStaticKeyRekeyMode() MacsecSecYAdvanceStaticKeyRekeyMode {
	obj := macsecSecYAdvanceStaticKeyRekeyMode{obj: &otg.MacsecSecYAdvanceStaticKeyRekeyMode{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYAdvanceStaticKeyRekeyMode) msg() *otg.MacsecSecYAdvanceStaticKeyRekeyMode {
	return obj.obj
}

func (obj *macsecSecYAdvanceStaticKeyRekeyMode) setMsg(msg *otg.MacsecSecYAdvanceStaticKeyRekeyMode) MacsecSecYAdvanceStaticKeyRekeyMode {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYAdvanceStaticKeyRekeyMode struct {
	obj *macsecSecYAdvanceStaticKeyRekeyMode
}

type marshalMacsecSecYAdvanceStaticKeyRekeyMode interface {
	// ToProto marshals MacsecSecYAdvanceStaticKeyRekeyMode to protobuf object *otg.MacsecSecYAdvanceStaticKeyRekeyMode
	ToProto() (*otg.MacsecSecYAdvanceStaticKeyRekeyMode, error)
	// ToPbText marshals MacsecSecYAdvanceStaticKeyRekeyMode to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYAdvanceStaticKeyRekeyMode to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYAdvanceStaticKeyRekeyMode to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYAdvanceStaticKeyRekeyMode struct {
	obj *macsecSecYAdvanceStaticKeyRekeyMode
}

type unMarshalMacsecSecYAdvanceStaticKeyRekeyMode interface {
	// FromProto unmarshals MacsecSecYAdvanceStaticKeyRekeyMode from protobuf object *otg.MacsecSecYAdvanceStaticKeyRekeyMode
	FromProto(msg *otg.MacsecSecYAdvanceStaticKeyRekeyMode) (MacsecSecYAdvanceStaticKeyRekeyMode, error)
	// FromPbText unmarshals MacsecSecYAdvanceStaticKeyRekeyMode from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYAdvanceStaticKeyRekeyMode from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYAdvanceStaticKeyRekeyMode from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYAdvanceStaticKeyRekeyMode) Marshal() marshalMacsecSecYAdvanceStaticKeyRekeyMode {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYAdvanceStaticKeyRekeyMode{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYAdvanceStaticKeyRekeyMode) Unmarshal() unMarshalMacsecSecYAdvanceStaticKeyRekeyMode {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYAdvanceStaticKeyRekeyMode{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYAdvanceStaticKeyRekeyMode) ToProto() (*otg.MacsecSecYAdvanceStaticKeyRekeyMode, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYAdvanceStaticKeyRekeyMode) FromProto(msg *otg.MacsecSecYAdvanceStaticKeyRekeyMode) (MacsecSecYAdvanceStaticKeyRekeyMode, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYAdvanceStaticKeyRekeyMode) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYAdvanceStaticKeyRekeyMode) FromPbText(value string) error {
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

func (m *marshalmacsecSecYAdvanceStaticKeyRekeyMode) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYAdvanceStaticKeyRekeyMode) FromYaml(value string) error {
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

func (m *marshalmacsecSecYAdvanceStaticKeyRekeyMode) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYAdvanceStaticKeyRekeyMode) FromJson(value string) error {
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

func (obj *macsecSecYAdvanceStaticKeyRekeyMode) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYAdvanceStaticKeyRekeyMode) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYAdvanceStaticKeyRekeyMode) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYAdvanceStaticKeyRekeyMode) Clone() (MacsecSecYAdvanceStaticKeyRekeyMode, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYAdvanceStaticKeyRekeyMode()
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

// MacsecSecYAdvanceStaticKeyRekeyMode is rekey mode.
type MacsecSecYAdvanceStaticKeyRekeyMode interface {
	Validation
	// msg marshals MacsecSecYAdvanceStaticKeyRekeyMode to protobuf object *otg.MacsecSecYAdvanceStaticKeyRekeyMode
	// and doesn't set defaults
	msg() *otg.MacsecSecYAdvanceStaticKeyRekeyMode
	// setMsg unmarshals MacsecSecYAdvanceStaticKeyRekeyMode from protobuf object *otg.MacsecSecYAdvanceStaticKeyRekeyMode
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYAdvanceStaticKeyRekeyMode) MacsecSecYAdvanceStaticKeyRekeyMode
	// provides marshal interface
	Marshal() marshalMacsecSecYAdvanceStaticKeyRekeyMode
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYAdvanceStaticKeyRekeyMode
	// validate validates MacsecSecYAdvanceStaticKeyRekeyMode
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYAdvanceStaticKeyRekeyMode, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MacsecSecYAdvanceStaticKeyRekeyModeChoiceEnum, set in MacsecSecYAdvanceStaticKeyRekeyMode
	Choice() MacsecSecYAdvanceStaticKeyRekeyModeChoiceEnum
	// setChoice assigns MacsecSecYAdvanceStaticKeyRekeyModeChoiceEnum provided by user to MacsecSecYAdvanceStaticKeyRekeyMode
	setChoice(value MacsecSecYAdvanceStaticKeyRekeyModeChoiceEnum) MacsecSecYAdvanceStaticKeyRekeyMode
	// HasChoice checks if Choice has been set in MacsecSecYAdvanceStaticKeyRekeyMode
	HasChoice() bool
	// getter for DontRekey to set choice.
	DontRekey()
	// getter for TimerBased to set choice.
	TimerBased()
	// getter for PnBased to set choice.
	PnBased()
}

type MacsecSecYAdvanceStaticKeyRekeyModeChoiceEnum string

// Enum of Choice on MacsecSecYAdvanceStaticKeyRekeyMode
var MacsecSecYAdvanceStaticKeyRekeyModeChoice = struct {
	DONT_REKEY  MacsecSecYAdvanceStaticKeyRekeyModeChoiceEnum
	TIMER_BASED MacsecSecYAdvanceStaticKeyRekeyModeChoiceEnum
	PN_BASED    MacsecSecYAdvanceStaticKeyRekeyModeChoiceEnum
}{
	DONT_REKEY:  MacsecSecYAdvanceStaticKeyRekeyModeChoiceEnum("dont_rekey"),
	TIMER_BASED: MacsecSecYAdvanceStaticKeyRekeyModeChoiceEnum("timer_based"),
	PN_BASED:    MacsecSecYAdvanceStaticKeyRekeyModeChoiceEnum("pn_based"),
}

func (obj *macsecSecYAdvanceStaticKeyRekeyMode) Choice() MacsecSecYAdvanceStaticKeyRekeyModeChoiceEnum {
	return MacsecSecYAdvanceStaticKeyRekeyModeChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for DontRekey to set choice
func (obj *macsecSecYAdvanceStaticKeyRekeyMode) DontRekey() {
	obj.setChoice(MacsecSecYAdvanceStaticKeyRekeyModeChoice.DONT_REKEY)
}

// getter for TimerBased to set choice
func (obj *macsecSecYAdvanceStaticKeyRekeyMode) TimerBased() {
	obj.setChoice(MacsecSecYAdvanceStaticKeyRekeyModeChoice.TIMER_BASED)
}

// getter for PnBased to set choice
func (obj *macsecSecYAdvanceStaticKeyRekeyMode) PnBased() {
	obj.setChoice(MacsecSecYAdvanceStaticKeyRekeyModeChoice.PN_BASED)
}

// Rekey mode choices.
// Choice returns a string
func (obj *macsecSecYAdvanceStaticKeyRekeyMode) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *macsecSecYAdvanceStaticKeyRekeyMode) setChoice(value MacsecSecYAdvanceStaticKeyRekeyModeChoiceEnum) MacsecSecYAdvanceStaticKeyRekeyMode {
	intValue, ok := otg.MacsecSecYAdvanceStaticKeyRekeyMode_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecSecYAdvanceStaticKeyRekeyModeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecSecYAdvanceStaticKeyRekeyMode_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *macsecSecYAdvanceStaticKeyRekeyMode) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecSecYAdvanceStaticKeyRekeyMode) setDefault() {
	var choices_set int = 0
	var choice MacsecSecYAdvanceStaticKeyRekeyModeChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MacsecSecYAdvanceStaticKeyRekeyModeChoice.DONT_REKEY)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MacsecSecYAdvanceStaticKeyRekeyMode")
			}
		} else {
			intVal := otg.MacsecSecYAdvanceStaticKeyRekeyMode_Choice_Enum_value[string(choice)]
			enumValue := otg.MacsecSecYAdvanceStaticKeyRekeyMode_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

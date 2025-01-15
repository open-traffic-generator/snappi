package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecAdvanceStaticKeyRekeyMode *****
type macsecAdvanceStaticKeyRekeyMode struct {
	validation
	obj          *otg.MacsecAdvanceStaticKeyRekeyMode
	marshaller   marshalMacsecAdvanceStaticKeyRekeyMode
	unMarshaller unMarshalMacsecAdvanceStaticKeyRekeyMode
}

func NewMacsecAdvanceStaticKeyRekeyMode() MacsecAdvanceStaticKeyRekeyMode {
	obj := macsecAdvanceStaticKeyRekeyMode{obj: &otg.MacsecAdvanceStaticKeyRekeyMode{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecAdvanceStaticKeyRekeyMode) msg() *otg.MacsecAdvanceStaticKeyRekeyMode {
	return obj.obj
}

func (obj *macsecAdvanceStaticKeyRekeyMode) setMsg(msg *otg.MacsecAdvanceStaticKeyRekeyMode) MacsecAdvanceStaticKeyRekeyMode {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecAdvanceStaticKeyRekeyMode struct {
	obj *macsecAdvanceStaticKeyRekeyMode
}

type marshalMacsecAdvanceStaticKeyRekeyMode interface {
	// ToProto marshals MacsecAdvanceStaticKeyRekeyMode to protobuf object *otg.MacsecAdvanceStaticKeyRekeyMode
	ToProto() (*otg.MacsecAdvanceStaticKeyRekeyMode, error)
	// ToPbText marshals MacsecAdvanceStaticKeyRekeyMode to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecAdvanceStaticKeyRekeyMode to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecAdvanceStaticKeyRekeyMode to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecAdvanceStaticKeyRekeyMode struct {
	obj *macsecAdvanceStaticKeyRekeyMode
}

type unMarshalMacsecAdvanceStaticKeyRekeyMode interface {
	// FromProto unmarshals MacsecAdvanceStaticKeyRekeyMode from protobuf object *otg.MacsecAdvanceStaticKeyRekeyMode
	FromProto(msg *otg.MacsecAdvanceStaticKeyRekeyMode) (MacsecAdvanceStaticKeyRekeyMode, error)
	// FromPbText unmarshals MacsecAdvanceStaticKeyRekeyMode from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecAdvanceStaticKeyRekeyMode from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecAdvanceStaticKeyRekeyMode from JSON text
	FromJson(value string) error
}

func (obj *macsecAdvanceStaticKeyRekeyMode) Marshal() marshalMacsecAdvanceStaticKeyRekeyMode {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecAdvanceStaticKeyRekeyMode{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecAdvanceStaticKeyRekeyMode) Unmarshal() unMarshalMacsecAdvanceStaticKeyRekeyMode {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecAdvanceStaticKeyRekeyMode{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecAdvanceStaticKeyRekeyMode) ToProto() (*otg.MacsecAdvanceStaticKeyRekeyMode, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecAdvanceStaticKeyRekeyMode) FromProto(msg *otg.MacsecAdvanceStaticKeyRekeyMode) (MacsecAdvanceStaticKeyRekeyMode, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecAdvanceStaticKeyRekeyMode) ToPbText() (string, error) {
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

func (m *unMarshalmacsecAdvanceStaticKeyRekeyMode) FromPbText(value string) error {
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

func (m *marshalmacsecAdvanceStaticKeyRekeyMode) ToYaml() (string, error) {
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

func (m *unMarshalmacsecAdvanceStaticKeyRekeyMode) FromYaml(value string) error {
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

func (m *marshalmacsecAdvanceStaticKeyRekeyMode) ToJson() (string, error) {
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

func (m *unMarshalmacsecAdvanceStaticKeyRekeyMode) FromJson(value string) error {
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

func (obj *macsecAdvanceStaticKeyRekeyMode) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecAdvanceStaticKeyRekeyMode) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecAdvanceStaticKeyRekeyMode) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecAdvanceStaticKeyRekeyMode) Clone() (MacsecAdvanceStaticKeyRekeyMode, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecAdvanceStaticKeyRekeyMode()
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

// MacsecAdvanceStaticKeyRekeyMode is rekey mode.
type MacsecAdvanceStaticKeyRekeyMode interface {
	Validation
	// msg marshals MacsecAdvanceStaticKeyRekeyMode to protobuf object *otg.MacsecAdvanceStaticKeyRekeyMode
	// and doesn't set defaults
	msg() *otg.MacsecAdvanceStaticKeyRekeyMode
	// setMsg unmarshals MacsecAdvanceStaticKeyRekeyMode from protobuf object *otg.MacsecAdvanceStaticKeyRekeyMode
	// and doesn't set defaults
	setMsg(*otg.MacsecAdvanceStaticKeyRekeyMode) MacsecAdvanceStaticKeyRekeyMode
	// provides marshal interface
	Marshal() marshalMacsecAdvanceStaticKeyRekeyMode
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecAdvanceStaticKeyRekeyMode
	// validate validates MacsecAdvanceStaticKeyRekeyMode
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecAdvanceStaticKeyRekeyMode, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MacsecAdvanceStaticKeyRekeyModeChoiceEnum, set in MacsecAdvanceStaticKeyRekeyMode
	Choice() MacsecAdvanceStaticKeyRekeyModeChoiceEnum
	// setChoice assigns MacsecAdvanceStaticKeyRekeyModeChoiceEnum provided by user to MacsecAdvanceStaticKeyRekeyMode
	setChoice(value MacsecAdvanceStaticKeyRekeyModeChoiceEnum) MacsecAdvanceStaticKeyRekeyMode
	// HasChoice checks if Choice has been set in MacsecAdvanceStaticKeyRekeyMode
	HasChoice() bool
	// getter for DontRekey to set choice.
	DontRekey()
	// getter for PnBased to set choice.
	PnBased()
	// getter for TimerBased to set choice.
	TimerBased()
}

type MacsecAdvanceStaticKeyRekeyModeChoiceEnum string

// Enum of Choice on MacsecAdvanceStaticKeyRekeyMode
var MacsecAdvanceStaticKeyRekeyModeChoice = struct {
	DONT_REKEY  MacsecAdvanceStaticKeyRekeyModeChoiceEnum
	TIMER_BASED MacsecAdvanceStaticKeyRekeyModeChoiceEnum
	PN_BASED    MacsecAdvanceStaticKeyRekeyModeChoiceEnum
}{
	DONT_REKEY:  MacsecAdvanceStaticKeyRekeyModeChoiceEnum("dont_rekey"),
	TIMER_BASED: MacsecAdvanceStaticKeyRekeyModeChoiceEnum("timer_based"),
	PN_BASED:    MacsecAdvanceStaticKeyRekeyModeChoiceEnum("pn_based"),
}

func (obj *macsecAdvanceStaticKeyRekeyMode) Choice() MacsecAdvanceStaticKeyRekeyModeChoiceEnum {
	return MacsecAdvanceStaticKeyRekeyModeChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for DontRekey to set choice
func (obj *macsecAdvanceStaticKeyRekeyMode) DontRekey() {
	obj.setChoice(MacsecAdvanceStaticKeyRekeyModeChoice.DONT_REKEY)
}

// getter for PnBased to set choice
func (obj *macsecAdvanceStaticKeyRekeyMode) PnBased() {
	obj.setChoice(MacsecAdvanceStaticKeyRekeyModeChoice.PN_BASED)
}

// getter for TimerBased to set choice
func (obj *macsecAdvanceStaticKeyRekeyMode) TimerBased() {
	obj.setChoice(MacsecAdvanceStaticKeyRekeyModeChoice.TIMER_BASED)
}

// Rekey mode choices.
// Choice returns a string
func (obj *macsecAdvanceStaticKeyRekeyMode) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *macsecAdvanceStaticKeyRekeyMode) setChoice(value MacsecAdvanceStaticKeyRekeyModeChoiceEnum) MacsecAdvanceStaticKeyRekeyMode {
	intValue, ok := otg.MacsecAdvanceStaticKeyRekeyMode_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecAdvanceStaticKeyRekeyModeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecAdvanceStaticKeyRekeyMode_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *macsecAdvanceStaticKeyRekeyMode) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecAdvanceStaticKeyRekeyMode) setDefault() {
	var choices_set int = 0
	var choice MacsecAdvanceStaticKeyRekeyModeChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MacsecAdvanceStaticKeyRekeyModeChoice.DONT_REKEY)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MacsecAdvanceStaticKeyRekeyMode")
			}
		} else {
			intVal := otg.MacsecAdvanceStaticKeyRekeyMode_Choice_Enum_value[string(choice)]
			enumValue := otg.MacsecAdvanceStaticKeyRekeyMode_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityStaticKeyRekeyMode *****
type secureEntityStaticKeyRekeyMode struct {
	validation
	obj              *otg.SecureEntityStaticKeyRekeyMode
	marshaller       marshalSecureEntityStaticKeyRekeyMode
	unMarshaller     unMarshalSecureEntityStaticKeyRekeyMode
	timerBasedHolder SecureEntityStaticKeyRekeyModeTimerBased
}

func NewSecureEntityStaticKeyRekeyMode() SecureEntityStaticKeyRekeyMode {
	obj := secureEntityStaticKeyRekeyMode{obj: &otg.SecureEntityStaticKeyRekeyMode{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityStaticKeyRekeyMode) msg() *otg.SecureEntityStaticKeyRekeyMode {
	return obj.obj
}

func (obj *secureEntityStaticKeyRekeyMode) setMsg(msg *otg.SecureEntityStaticKeyRekeyMode) SecureEntityStaticKeyRekeyMode {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityStaticKeyRekeyMode struct {
	obj *secureEntityStaticKeyRekeyMode
}

type marshalSecureEntityStaticKeyRekeyMode interface {
	// ToProto marshals SecureEntityStaticKeyRekeyMode to protobuf object *otg.SecureEntityStaticKeyRekeyMode
	ToProto() (*otg.SecureEntityStaticKeyRekeyMode, error)
	// ToPbText marshals SecureEntityStaticKeyRekeyMode to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityStaticKeyRekeyMode to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityStaticKeyRekeyMode to JSON text
	ToJson() (string, error)
}

type unMarshalsecureEntityStaticKeyRekeyMode struct {
	obj *secureEntityStaticKeyRekeyMode
}

type unMarshalSecureEntityStaticKeyRekeyMode interface {
	// FromProto unmarshals SecureEntityStaticKeyRekeyMode from protobuf object *otg.SecureEntityStaticKeyRekeyMode
	FromProto(msg *otg.SecureEntityStaticKeyRekeyMode) (SecureEntityStaticKeyRekeyMode, error)
	// FromPbText unmarshals SecureEntityStaticKeyRekeyMode from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityStaticKeyRekeyMode from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityStaticKeyRekeyMode from JSON text
	FromJson(value string) error
}

func (obj *secureEntityStaticKeyRekeyMode) Marshal() marshalSecureEntityStaticKeyRekeyMode {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityStaticKeyRekeyMode{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityStaticKeyRekeyMode) Unmarshal() unMarshalSecureEntityStaticKeyRekeyMode {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityStaticKeyRekeyMode{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityStaticKeyRekeyMode) ToProto() (*otg.SecureEntityStaticKeyRekeyMode, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityStaticKeyRekeyMode) FromProto(msg *otg.SecureEntityStaticKeyRekeyMode) (SecureEntityStaticKeyRekeyMode, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityStaticKeyRekeyMode) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityStaticKeyRekeyMode) FromPbText(value string) error {
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

func (m *marshalsecureEntityStaticKeyRekeyMode) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityStaticKeyRekeyMode) FromYaml(value string) error {
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

func (m *marshalsecureEntityStaticKeyRekeyMode) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityStaticKeyRekeyMode) FromJson(value string) error {
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

func (obj *secureEntityStaticKeyRekeyMode) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityStaticKeyRekeyMode) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityStaticKeyRekeyMode) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityStaticKeyRekeyMode) Clone() (SecureEntityStaticKeyRekeyMode, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityStaticKeyRekeyMode()
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

func (obj *secureEntityStaticKeyRekeyMode) setNil() {
	obj.timerBasedHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// SecureEntityStaticKeyRekeyMode is rekey mode.
type SecureEntityStaticKeyRekeyMode interface {
	Validation
	// msg marshals SecureEntityStaticKeyRekeyMode to protobuf object *otg.SecureEntityStaticKeyRekeyMode
	// and doesn't set defaults
	msg() *otg.SecureEntityStaticKeyRekeyMode
	// setMsg unmarshals SecureEntityStaticKeyRekeyMode from protobuf object *otg.SecureEntityStaticKeyRekeyMode
	// and doesn't set defaults
	setMsg(*otg.SecureEntityStaticKeyRekeyMode) SecureEntityStaticKeyRekeyMode
	// provides marshal interface
	Marshal() marshalSecureEntityStaticKeyRekeyMode
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityStaticKeyRekeyMode
	// validate validates SecureEntityStaticKeyRekeyMode
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityStaticKeyRekeyMode, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns SecureEntityStaticKeyRekeyModeChoiceEnum, set in SecureEntityStaticKeyRekeyMode
	Choice() SecureEntityStaticKeyRekeyModeChoiceEnum
	// setChoice assigns SecureEntityStaticKeyRekeyModeChoiceEnum provided by user to SecureEntityStaticKeyRekeyMode
	setChoice(value SecureEntityStaticKeyRekeyModeChoiceEnum) SecureEntityStaticKeyRekeyMode
	// HasChoice checks if Choice has been set in SecureEntityStaticKeyRekeyMode
	HasChoice() bool
	// getter for DontRekey to set choice.
	DontRekey()
	// getter for PnBased to set choice.
	PnBased()
	// TimerBased returns SecureEntityStaticKeyRekeyModeTimerBased, set in SecureEntityStaticKeyRekeyMode.
	// SecureEntityStaticKeyRekeyModeTimerBased is timer based periodic rekey properties.
	TimerBased() SecureEntityStaticKeyRekeyModeTimerBased
	// SetTimerBased assigns SecureEntityStaticKeyRekeyModeTimerBased provided by user to SecureEntityStaticKeyRekeyMode.
	// SecureEntityStaticKeyRekeyModeTimerBased is timer based periodic rekey properties.
	SetTimerBased(value SecureEntityStaticKeyRekeyModeTimerBased) SecureEntityStaticKeyRekeyMode
	// HasTimerBased checks if TimerBased has been set in SecureEntityStaticKeyRekeyMode
	HasTimerBased() bool
	setNil()
}

type SecureEntityStaticKeyRekeyModeChoiceEnum string

// Enum of Choice on SecureEntityStaticKeyRekeyMode
var SecureEntityStaticKeyRekeyModeChoice = struct {
	DONT_REKEY  SecureEntityStaticKeyRekeyModeChoiceEnum
	TIMER_BASED SecureEntityStaticKeyRekeyModeChoiceEnum
	PN_BASED    SecureEntityStaticKeyRekeyModeChoiceEnum
}{
	DONT_REKEY:  SecureEntityStaticKeyRekeyModeChoiceEnum("dont_rekey"),
	TIMER_BASED: SecureEntityStaticKeyRekeyModeChoiceEnum("timer_based"),
	PN_BASED:    SecureEntityStaticKeyRekeyModeChoiceEnum("pn_based"),
}

func (obj *secureEntityStaticKeyRekeyMode) Choice() SecureEntityStaticKeyRekeyModeChoiceEnum {
	return SecureEntityStaticKeyRekeyModeChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for DontRekey to set choice
func (obj *secureEntityStaticKeyRekeyMode) DontRekey() {
	obj.setChoice(SecureEntityStaticKeyRekeyModeChoice.DONT_REKEY)
}

// getter for PnBased to set choice
func (obj *secureEntityStaticKeyRekeyMode) PnBased() {
	obj.setChoice(SecureEntityStaticKeyRekeyModeChoice.PN_BASED)
}

// Rekey mode choices.
// Choice returns a string
func (obj *secureEntityStaticKeyRekeyMode) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *secureEntityStaticKeyRekeyMode) setChoice(value SecureEntityStaticKeyRekeyModeChoiceEnum) SecureEntityStaticKeyRekeyMode {
	intValue, ok := otg.SecureEntityStaticKeyRekeyMode_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on SecureEntityStaticKeyRekeyModeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.SecureEntityStaticKeyRekeyMode_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.TimerBased = nil
	obj.timerBasedHolder = nil

	if value == SecureEntityStaticKeyRekeyModeChoice.TIMER_BASED {
		obj.obj.TimerBased = NewSecureEntityStaticKeyRekeyModeTimerBased().msg()
	}

	return obj
}

// Container for timer based periodic rekey properties.
// TimerBased returns a SecureEntityStaticKeyRekeyModeTimerBased
func (obj *secureEntityStaticKeyRekeyMode) TimerBased() SecureEntityStaticKeyRekeyModeTimerBased {
	if obj.obj.TimerBased == nil {
		obj.setChoice(SecureEntityStaticKeyRekeyModeChoice.TIMER_BASED)
	}
	if obj.timerBasedHolder == nil {
		obj.timerBasedHolder = &secureEntityStaticKeyRekeyModeTimerBased{obj: obj.obj.TimerBased}
	}
	return obj.timerBasedHolder
}

// Container for timer based periodic rekey properties.
// TimerBased returns a SecureEntityStaticKeyRekeyModeTimerBased
func (obj *secureEntityStaticKeyRekeyMode) HasTimerBased() bool {
	return obj.obj.TimerBased != nil
}

// Container for timer based periodic rekey properties.
// SetTimerBased sets the SecureEntityStaticKeyRekeyModeTimerBased value in the SecureEntityStaticKeyRekeyMode object
func (obj *secureEntityStaticKeyRekeyMode) SetTimerBased(value SecureEntityStaticKeyRekeyModeTimerBased) SecureEntityStaticKeyRekeyMode {
	obj.setChoice(SecureEntityStaticKeyRekeyModeChoice.TIMER_BASED)
	obj.timerBasedHolder = nil
	obj.obj.TimerBased = value.msg()

	return obj
}

func (obj *secureEntityStaticKeyRekeyMode) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TimerBased != nil {

		obj.TimerBased().validateObj(vObj, set_default)
	}

}

func (obj *secureEntityStaticKeyRekeyMode) setDefault() {
	var choices_set int = 0
	var choice SecureEntityStaticKeyRekeyModeChoiceEnum

	if obj.obj.TimerBased != nil {
		choices_set += 1
		choice = SecureEntityStaticKeyRekeyModeChoice.TIMER_BASED
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(SecureEntityStaticKeyRekeyModeChoice.DONT_REKEY)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in SecureEntityStaticKeyRekeyMode")
			}
		} else {
			intVal := otg.SecureEntityStaticKeyRekeyMode_Choice_Enum_value[string(choice)]
			enumValue := otg.SecureEntityStaticKeyRekeyMode_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

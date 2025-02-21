package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecStaticKeyRekeyMode *****
type macsecStaticKeyRekeyMode struct {
	validation
	obj              *otg.MacsecStaticKeyRekeyMode
	marshaller       marshalMacsecStaticKeyRekeyMode
	unMarshaller     unMarshalMacsecStaticKeyRekeyMode
	timerBasedHolder MacsecStaticKeyRekeyModeTimerBased
}

func NewMacsecStaticKeyRekeyMode() MacsecStaticKeyRekeyMode {
	obj := macsecStaticKeyRekeyMode{obj: &otg.MacsecStaticKeyRekeyMode{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecStaticKeyRekeyMode) msg() *otg.MacsecStaticKeyRekeyMode {
	return obj.obj
}

func (obj *macsecStaticKeyRekeyMode) setMsg(msg *otg.MacsecStaticKeyRekeyMode) MacsecStaticKeyRekeyMode {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecStaticKeyRekeyMode struct {
	obj *macsecStaticKeyRekeyMode
}

type marshalMacsecStaticKeyRekeyMode interface {
	// ToProto marshals MacsecStaticKeyRekeyMode to protobuf object *otg.MacsecStaticKeyRekeyMode
	ToProto() (*otg.MacsecStaticKeyRekeyMode, error)
	// ToPbText marshals MacsecStaticKeyRekeyMode to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecStaticKeyRekeyMode to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecStaticKeyRekeyMode to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecStaticKeyRekeyMode struct {
	obj *macsecStaticKeyRekeyMode
}

type unMarshalMacsecStaticKeyRekeyMode interface {
	// FromProto unmarshals MacsecStaticKeyRekeyMode from protobuf object *otg.MacsecStaticKeyRekeyMode
	FromProto(msg *otg.MacsecStaticKeyRekeyMode) (MacsecStaticKeyRekeyMode, error)
	// FromPbText unmarshals MacsecStaticKeyRekeyMode from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecStaticKeyRekeyMode from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecStaticKeyRekeyMode from JSON text
	FromJson(value string) error
}

func (obj *macsecStaticKeyRekeyMode) Marshal() marshalMacsecStaticKeyRekeyMode {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecStaticKeyRekeyMode{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecStaticKeyRekeyMode) Unmarshal() unMarshalMacsecStaticKeyRekeyMode {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecStaticKeyRekeyMode{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecStaticKeyRekeyMode) ToProto() (*otg.MacsecStaticKeyRekeyMode, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecStaticKeyRekeyMode) FromProto(msg *otg.MacsecStaticKeyRekeyMode) (MacsecStaticKeyRekeyMode, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecStaticKeyRekeyMode) ToPbText() (string, error) {
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

func (m *unMarshalmacsecStaticKeyRekeyMode) FromPbText(value string) error {
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

func (m *marshalmacsecStaticKeyRekeyMode) ToYaml() (string, error) {
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

func (m *unMarshalmacsecStaticKeyRekeyMode) FromYaml(value string) error {
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

func (m *marshalmacsecStaticKeyRekeyMode) ToJson() (string, error) {
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

func (m *unMarshalmacsecStaticKeyRekeyMode) FromJson(value string) error {
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

func (obj *macsecStaticKeyRekeyMode) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecStaticKeyRekeyMode) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecStaticKeyRekeyMode) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecStaticKeyRekeyMode) Clone() (MacsecStaticKeyRekeyMode, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecStaticKeyRekeyMode()
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

func (obj *macsecStaticKeyRekeyMode) setNil() {
	obj.timerBasedHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecStaticKeyRekeyMode is rekey mode.
type MacsecStaticKeyRekeyMode interface {
	Validation
	// msg marshals MacsecStaticKeyRekeyMode to protobuf object *otg.MacsecStaticKeyRekeyMode
	// and doesn't set defaults
	msg() *otg.MacsecStaticKeyRekeyMode
	// setMsg unmarshals MacsecStaticKeyRekeyMode from protobuf object *otg.MacsecStaticKeyRekeyMode
	// and doesn't set defaults
	setMsg(*otg.MacsecStaticKeyRekeyMode) MacsecStaticKeyRekeyMode
	// provides marshal interface
	Marshal() marshalMacsecStaticKeyRekeyMode
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecStaticKeyRekeyMode
	// validate validates MacsecStaticKeyRekeyMode
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecStaticKeyRekeyMode, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MacsecStaticKeyRekeyModeChoiceEnum, set in MacsecStaticKeyRekeyMode
	Choice() MacsecStaticKeyRekeyModeChoiceEnum
	// setChoice assigns MacsecStaticKeyRekeyModeChoiceEnum provided by user to MacsecStaticKeyRekeyMode
	setChoice(value MacsecStaticKeyRekeyModeChoiceEnum) MacsecStaticKeyRekeyMode
	// HasChoice checks if Choice has been set in MacsecStaticKeyRekeyMode
	HasChoice() bool
	// getter for PnBased to set choice.
	PnBased()
	// getter for DontRekey to set choice.
	DontRekey()
	// TimerBased returns MacsecStaticKeyRekeyModeTimerBased, set in MacsecStaticKeyRekeyMode.
	// MacsecStaticKeyRekeyModeTimerBased is timer based periodic rekey properties.
	TimerBased() MacsecStaticKeyRekeyModeTimerBased
	// SetTimerBased assigns MacsecStaticKeyRekeyModeTimerBased provided by user to MacsecStaticKeyRekeyMode.
	// MacsecStaticKeyRekeyModeTimerBased is timer based periodic rekey properties.
	SetTimerBased(value MacsecStaticKeyRekeyModeTimerBased) MacsecStaticKeyRekeyMode
	// HasTimerBased checks if TimerBased has been set in MacsecStaticKeyRekeyMode
	HasTimerBased() bool
	setNil()
}

type MacsecStaticKeyRekeyModeChoiceEnum string

// Enum of Choice on MacsecStaticKeyRekeyMode
var MacsecStaticKeyRekeyModeChoice = struct {
	DONT_REKEY  MacsecStaticKeyRekeyModeChoiceEnum
	TIMER_BASED MacsecStaticKeyRekeyModeChoiceEnum
	PN_BASED    MacsecStaticKeyRekeyModeChoiceEnum
}{
	DONT_REKEY:  MacsecStaticKeyRekeyModeChoiceEnum("dont_rekey"),
	TIMER_BASED: MacsecStaticKeyRekeyModeChoiceEnum("timer_based"),
	PN_BASED:    MacsecStaticKeyRekeyModeChoiceEnum("pn_based"),
}

func (obj *macsecStaticKeyRekeyMode) Choice() MacsecStaticKeyRekeyModeChoiceEnum {
	return MacsecStaticKeyRekeyModeChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for PnBased to set choice
func (obj *macsecStaticKeyRekeyMode) PnBased() {
	obj.setChoice(MacsecStaticKeyRekeyModeChoice.PN_BASED)
}

// getter for DontRekey to set choice
func (obj *macsecStaticKeyRekeyMode) DontRekey() {
	obj.setChoice(MacsecStaticKeyRekeyModeChoice.DONT_REKEY)
}

// Rekey mode choices.
// Choice returns a string
func (obj *macsecStaticKeyRekeyMode) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *macsecStaticKeyRekeyMode) setChoice(value MacsecStaticKeyRekeyModeChoiceEnum) MacsecStaticKeyRekeyMode {
	intValue, ok := otg.MacsecStaticKeyRekeyMode_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecStaticKeyRekeyModeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecStaticKeyRekeyMode_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.TimerBased = nil
	obj.timerBasedHolder = nil

	if value == MacsecStaticKeyRekeyModeChoice.TIMER_BASED {
		obj.obj.TimerBased = NewMacsecStaticKeyRekeyModeTimerBased().msg()
	}

	return obj
}

// Container for timer based periodic rekey properties.
// TimerBased returns a MacsecStaticKeyRekeyModeTimerBased
func (obj *macsecStaticKeyRekeyMode) TimerBased() MacsecStaticKeyRekeyModeTimerBased {
	if obj.obj.TimerBased == nil {
		obj.setChoice(MacsecStaticKeyRekeyModeChoice.TIMER_BASED)
	}
	if obj.timerBasedHolder == nil {
		obj.timerBasedHolder = &macsecStaticKeyRekeyModeTimerBased{obj: obj.obj.TimerBased}
	}
	return obj.timerBasedHolder
}

// Container for timer based periodic rekey properties.
// TimerBased returns a MacsecStaticKeyRekeyModeTimerBased
func (obj *macsecStaticKeyRekeyMode) HasTimerBased() bool {
	return obj.obj.TimerBased != nil
}

// Container for timer based periodic rekey properties.
// SetTimerBased sets the MacsecStaticKeyRekeyModeTimerBased value in the MacsecStaticKeyRekeyMode object
func (obj *macsecStaticKeyRekeyMode) SetTimerBased(value MacsecStaticKeyRekeyModeTimerBased) MacsecStaticKeyRekeyMode {
	obj.setChoice(MacsecStaticKeyRekeyModeChoice.TIMER_BASED)
	obj.timerBasedHolder = nil
	obj.obj.TimerBased = value.msg()

	return obj
}

func (obj *macsecStaticKeyRekeyMode) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TimerBased != nil {

		obj.TimerBased().validateObj(vObj, set_default)
	}

}

func (obj *macsecStaticKeyRekeyMode) setDefault() {
	var choices_set int = 0
	var choice MacsecStaticKeyRekeyModeChoiceEnum

	if obj.obj.TimerBased != nil {
		choices_set += 1
		choice = MacsecStaticKeyRekeyModeChoice.TIMER_BASED
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MacsecStaticKeyRekeyModeChoice.DONT_REKEY)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MacsecStaticKeyRekeyMode")
			}
		} else {
			intVal := otg.MacsecStaticKeyRekeyMode_Choice_Enum_value[string(choice)]
			enumValue := otg.MacsecStaticKeyRekeyMode_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

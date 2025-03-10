package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MkaBasicRekeyMode *****
type mkaBasicRekeyMode struct {
	validation
	obj              *otg.MkaBasicRekeyMode
	marshaller       marshalMkaBasicRekeyMode
	unMarshaller     unMarshalMkaBasicRekeyMode
	timerBasedHolder MkaBasicRekeyModeTimerBased
}

func NewMkaBasicRekeyMode() MkaBasicRekeyMode {
	obj := mkaBasicRekeyMode{obj: &otg.MkaBasicRekeyMode{}}
	obj.setDefault()
	return &obj
}

func (obj *mkaBasicRekeyMode) msg() *otg.MkaBasicRekeyMode {
	return obj.obj
}

func (obj *mkaBasicRekeyMode) setMsg(msg *otg.MkaBasicRekeyMode) MkaBasicRekeyMode {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmkaBasicRekeyMode struct {
	obj *mkaBasicRekeyMode
}

type marshalMkaBasicRekeyMode interface {
	// ToProto marshals MkaBasicRekeyMode to protobuf object *otg.MkaBasicRekeyMode
	ToProto() (*otg.MkaBasicRekeyMode, error)
	// ToPbText marshals MkaBasicRekeyMode to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MkaBasicRekeyMode to YAML text
	ToYaml() (string, error)
	// ToJson marshals MkaBasicRekeyMode to JSON text
	ToJson() (string, error)
}

type unMarshalmkaBasicRekeyMode struct {
	obj *mkaBasicRekeyMode
}

type unMarshalMkaBasicRekeyMode interface {
	// FromProto unmarshals MkaBasicRekeyMode from protobuf object *otg.MkaBasicRekeyMode
	FromProto(msg *otg.MkaBasicRekeyMode) (MkaBasicRekeyMode, error)
	// FromPbText unmarshals MkaBasicRekeyMode from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MkaBasicRekeyMode from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MkaBasicRekeyMode from JSON text
	FromJson(value string) error
}

func (obj *mkaBasicRekeyMode) Marshal() marshalMkaBasicRekeyMode {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmkaBasicRekeyMode{obj: obj}
	}
	return obj.marshaller
}

func (obj *mkaBasicRekeyMode) Unmarshal() unMarshalMkaBasicRekeyMode {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmkaBasicRekeyMode{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmkaBasicRekeyMode) ToProto() (*otg.MkaBasicRekeyMode, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmkaBasicRekeyMode) FromProto(msg *otg.MkaBasicRekeyMode) (MkaBasicRekeyMode, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmkaBasicRekeyMode) ToPbText() (string, error) {
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

func (m *unMarshalmkaBasicRekeyMode) FromPbText(value string) error {
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

func (m *marshalmkaBasicRekeyMode) ToYaml() (string, error) {
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

func (m *unMarshalmkaBasicRekeyMode) FromYaml(value string) error {
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

func (m *marshalmkaBasicRekeyMode) ToJson() (string, error) {
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

func (m *unMarshalmkaBasicRekeyMode) FromJson(value string) error {
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

func (obj *mkaBasicRekeyMode) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mkaBasicRekeyMode) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mkaBasicRekeyMode) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mkaBasicRekeyMode) Clone() (MkaBasicRekeyMode, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMkaBasicRekeyMode()
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

func (obj *mkaBasicRekeyMode) setNil() {
	obj.timerBasedHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MkaBasicRekeyMode is rekey mode.
type MkaBasicRekeyMode interface {
	Validation
	// msg marshals MkaBasicRekeyMode to protobuf object *otg.MkaBasicRekeyMode
	// and doesn't set defaults
	msg() *otg.MkaBasicRekeyMode
	// setMsg unmarshals MkaBasicRekeyMode from protobuf object *otg.MkaBasicRekeyMode
	// and doesn't set defaults
	setMsg(*otg.MkaBasicRekeyMode) MkaBasicRekeyMode
	// provides marshal interface
	Marshal() marshalMkaBasicRekeyMode
	// provides unmarshal interface
	Unmarshal() unMarshalMkaBasicRekeyMode
	// validate validates MkaBasicRekeyMode
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MkaBasicRekeyMode, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MkaBasicRekeyModeChoiceEnum, set in MkaBasicRekeyMode
	Choice() MkaBasicRekeyModeChoiceEnum
	// setChoice assigns MkaBasicRekeyModeChoiceEnum provided by user to MkaBasicRekeyMode
	setChoice(value MkaBasicRekeyModeChoiceEnum) MkaBasicRekeyMode
	// HasChoice checks if Choice has been set in MkaBasicRekeyMode
	HasChoice() bool
	// getter for PnBased to set choice.
	PnBased()
	// getter for DontRekey to set choice.
	DontRekey()
	// TimerBased returns MkaBasicRekeyModeTimerBased, set in MkaBasicRekeyMode.
	// MkaBasicRekeyModeTimerBased is timer based periodic rekey properties.
	TimerBased() MkaBasicRekeyModeTimerBased
	// SetTimerBased assigns MkaBasicRekeyModeTimerBased provided by user to MkaBasicRekeyMode.
	// MkaBasicRekeyModeTimerBased is timer based periodic rekey properties.
	SetTimerBased(value MkaBasicRekeyModeTimerBased) MkaBasicRekeyMode
	// HasTimerBased checks if TimerBased has been set in MkaBasicRekeyMode
	HasTimerBased() bool
	setNil()
}

type MkaBasicRekeyModeChoiceEnum string

// Enum of Choice on MkaBasicRekeyMode
var MkaBasicRekeyModeChoice = struct {
	DONT_REKEY  MkaBasicRekeyModeChoiceEnum
	TIMER_BASED MkaBasicRekeyModeChoiceEnum
	PN_BASED    MkaBasicRekeyModeChoiceEnum
}{
	DONT_REKEY:  MkaBasicRekeyModeChoiceEnum("dont_rekey"),
	TIMER_BASED: MkaBasicRekeyModeChoiceEnum("timer_based"),
	PN_BASED:    MkaBasicRekeyModeChoiceEnum("pn_based"),
}

func (obj *mkaBasicRekeyMode) Choice() MkaBasicRekeyModeChoiceEnum {
	return MkaBasicRekeyModeChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for PnBased to set choice
func (obj *mkaBasicRekeyMode) PnBased() {
	obj.setChoice(MkaBasicRekeyModeChoice.PN_BASED)
}

// getter for DontRekey to set choice
func (obj *mkaBasicRekeyMode) DontRekey() {
	obj.setChoice(MkaBasicRekeyModeChoice.DONT_REKEY)
}

// Mode choices.
// Choice returns a string
func (obj *mkaBasicRekeyMode) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *mkaBasicRekeyMode) setChoice(value MkaBasicRekeyModeChoiceEnum) MkaBasicRekeyMode {
	intValue, ok := otg.MkaBasicRekeyMode_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MkaBasicRekeyModeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MkaBasicRekeyMode_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.TimerBased = nil
	obj.timerBasedHolder = nil

	if value == MkaBasicRekeyModeChoice.TIMER_BASED {
		obj.obj.TimerBased = NewMkaBasicRekeyModeTimerBased().msg()
	}

	return obj
}

// Container for timer based periodic rekey properties.
// TimerBased returns a MkaBasicRekeyModeTimerBased
func (obj *mkaBasicRekeyMode) TimerBased() MkaBasicRekeyModeTimerBased {
	if obj.obj.TimerBased == nil {
		obj.setChoice(MkaBasicRekeyModeChoice.TIMER_BASED)
	}
	if obj.timerBasedHolder == nil {
		obj.timerBasedHolder = &mkaBasicRekeyModeTimerBased{obj: obj.obj.TimerBased}
	}
	return obj.timerBasedHolder
}

// Container for timer based periodic rekey properties.
// TimerBased returns a MkaBasicRekeyModeTimerBased
func (obj *mkaBasicRekeyMode) HasTimerBased() bool {
	return obj.obj.TimerBased != nil
}

// Container for timer based periodic rekey properties.
// SetTimerBased sets the MkaBasicRekeyModeTimerBased value in the MkaBasicRekeyMode object
func (obj *mkaBasicRekeyMode) SetTimerBased(value MkaBasicRekeyModeTimerBased) MkaBasicRekeyMode {
	obj.setChoice(MkaBasicRekeyModeChoice.TIMER_BASED)
	obj.timerBasedHolder = nil
	obj.obj.TimerBased = value.msg()

	return obj
}

func (obj *mkaBasicRekeyMode) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TimerBased != nil {

		obj.TimerBased().validateObj(vObj, set_default)
	}

}

func (obj *mkaBasicRekeyMode) setDefault() {
	var choices_set int = 0
	var choice MkaBasicRekeyModeChoiceEnum

	if obj.obj.TimerBased != nil {
		choices_set += 1
		choice = MkaBasicRekeyModeChoice.TIMER_BASED
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MkaBasicRekeyModeChoice.DONT_REKEY)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MkaBasicRekeyMode")
			}
		} else {
			intVal := otg.MkaBasicRekeyMode_Choice_Enum_value[string(choice)]
			enumValue := otg.MkaBasicRekeyMode_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

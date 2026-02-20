package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecValidateFrames *****
type macsecValidateFrames struct {
	validation
	obj          *otg.MacsecValidateFrames
	marshaller   marshalMacsecValidateFrames
	unMarshaller unMarshalMacsecValidateFrames
}

func NewMacsecValidateFrames() MacsecValidateFrames {
	obj := macsecValidateFrames{obj: &otg.MacsecValidateFrames{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecValidateFrames) msg() *otg.MacsecValidateFrames {
	return obj.obj
}

func (obj *macsecValidateFrames) setMsg(msg *otg.MacsecValidateFrames) MacsecValidateFrames {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecValidateFrames struct {
	obj *macsecValidateFrames
}

type marshalMacsecValidateFrames interface {
	// ToProto marshals MacsecValidateFrames to protobuf object *otg.MacsecValidateFrames
	ToProto() (*otg.MacsecValidateFrames, error)
	// ToPbText marshals MacsecValidateFrames to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecValidateFrames to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecValidateFrames to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecValidateFrames struct {
	obj *macsecValidateFrames
}

type unMarshalMacsecValidateFrames interface {
	// FromProto unmarshals MacsecValidateFrames from protobuf object *otg.MacsecValidateFrames
	FromProto(msg *otg.MacsecValidateFrames) (MacsecValidateFrames, error)
	// FromPbText unmarshals MacsecValidateFrames from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecValidateFrames from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecValidateFrames from JSON text
	FromJson(value string) error
}

func (obj *macsecValidateFrames) Marshal() marshalMacsecValidateFrames {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecValidateFrames{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecValidateFrames) Unmarshal() unMarshalMacsecValidateFrames {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecValidateFrames{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecValidateFrames) ToProto() (*otg.MacsecValidateFrames, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecValidateFrames) FromProto(msg *otg.MacsecValidateFrames) (MacsecValidateFrames, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecValidateFrames) ToPbText() (string, error) {
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

func (m *unMarshalmacsecValidateFrames) FromPbText(value string) error {
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

func (m *marshalmacsecValidateFrames) ToYaml() (string, error) {
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

func (m *unMarshalmacsecValidateFrames) FromYaml(value string) error {
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

func (m *marshalmacsecValidateFrames) ToJson() (string, error) {
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

func (m *unMarshalmacsecValidateFrames) FromJson(value string) error {
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

func (obj *macsecValidateFrames) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecValidateFrames) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecValidateFrames) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecValidateFrames) Clone() (MacsecValidateFrames, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecValidateFrames()
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

// MacsecValidateFrames is controls validation of received frames.
type MacsecValidateFrames interface {
	Validation
	// msg marshals MacsecValidateFrames to protobuf object *otg.MacsecValidateFrames
	// and doesn't set defaults
	msg() *otg.MacsecValidateFrames
	// setMsg unmarshals MacsecValidateFrames from protobuf object *otg.MacsecValidateFrames
	// and doesn't set defaults
	setMsg(*otg.MacsecValidateFrames) MacsecValidateFrames
	// provides marshal interface
	Marshal() marshalMacsecValidateFrames
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecValidateFrames
	// validate validates MacsecValidateFrames
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecValidateFrames, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MacsecValidateFramesChoiceEnum, set in MacsecValidateFrames
	Choice() MacsecValidateFramesChoiceEnum
	// setChoice assigns MacsecValidateFramesChoiceEnum provided by user to MacsecValidateFrames
	setChoice(value MacsecValidateFramesChoiceEnum) MacsecValidateFrames
	// HasChoice checks if Choice has been set in MacsecValidateFrames
	HasChoice() bool
	// getter for None to set choice.
	None()
	// getter for Disabled to set choice.
	Disabled()
	// getter for Strict to set choice.
	Strict()
	// getter for Check to set choice.
	Check()
}

type MacsecValidateFramesChoiceEnum string

// Enum of Choice on MacsecValidateFrames
var MacsecValidateFramesChoice = struct {
	DISABLED MacsecValidateFramesChoiceEnum
	CHECK    MacsecValidateFramesChoiceEnum
	STRICT   MacsecValidateFramesChoiceEnum
	NONE     MacsecValidateFramesChoiceEnum
}{
	DISABLED: MacsecValidateFramesChoiceEnum("disabled"),
	CHECK:    MacsecValidateFramesChoiceEnum("check"),
	STRICT:   MacsecValidateFramesChoiceEnum("strict"),
	NONE:     MacsecValidateFramesChoiceEnum("none"),
}

func (obj *macsecValidateFrames) Choice() MacsecValidateFramesChoiceEnum {
	return MacsecValidateFramesChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for None to set choice
func (obj *macsecValidateFrames) None() {
	obj.setChoice(MacsecValidateFramesChoice.NONE)
}

// getter for Disabled to set choice
func (obj *macsecValidateFrames) Disabled() {
	obj.setChoice(MacsecValidateFramesChoice.DISABLED)
}

// getter for Strict to set choice
func (obj *macsecValidateFrames) Strict() {
	obj.setChoice(MacsecValidateFramesChoice.STRICT)
}

// getter for Check to set choice
func (obj *macsecValidateFrames) Check() {
	obj.setChoice(MacsecValidateFramesChoice.CHECK)
}

// Controls validation of received frames. disabled - disable validation, remove SecTAGs and ICVs (if present) from received frames. check - enable validation, do not discard invalid frames. strict - enable validation and discard invalid frames. none - no processing, do not remove SecTAGs or ICVs.
// Choice returns a string
func (obj *macsecValidateFrames) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *macsecValidateFrames) setChoice(value MacsecValidateFramesChoiceEnum) MacsecValidateFrames {
	intValue, ok := otg.MacsecValidateFrames_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecValidateFramesChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecValidateFrames_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *macsecValidateFrames) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecValidateFrames) setDefault() {
	var choices_set int = 0
	var choice MacsecValidateFramesChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MacsecValidateFramesChoice.CHECK)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MacsecValidateFrames")
			}
		} else {
			intVal := otg.MacsecValidateFrames_Choice_Enum_value[string(choice)]
			enumValue := otg.MacsecValidateFrames_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

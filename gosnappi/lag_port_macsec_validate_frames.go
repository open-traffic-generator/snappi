package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagPortMacsecValidateFrames *****
type lagPortMacsecValidateFrames struct {
	validation
	obj          *otg.LagPortMacsecValidateFrames
	marshaller   marshalLagPortMacsecValidateFrames
	unMarshaller unMarshalLagPortMacsecValidateFrames
}

func NewLagPortMacsecValidateFrames() LagPortMacsecValidateFrames {
	obj := lagPortMacsecValidateFrames{obj: &otg.LagPortMacsecValidateFrames{}}
	obj.setDefault()
	return &obj
}

func (obj *lagPortMacsecValidateFrames) msg() *otg.LagPortMacsecValidateFrames {
	return obj.obj
}

func (obj *lagPortMacsecValidateFrames) setMsg(msg *otg.LagPortMacsecValidateFrames) LagPortMacsecValidateFrames {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagPortMacsecValidateFrames struct {
	obj *lagPortMacsecValidateFrames
}

type marshalLagPortMacsecValidateFrames interface {
	// ToProto marshals LagPortMacsecValidateFrames to protobuf object *otg.LagPortMacsecValidateFrames
	ToProto() (*otg.LagPortMacsecValidateFrames, error)
	// ToPbText marshals LagPortMacsecValidateFrames to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagPortMacsecValidateFrames to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagPortMacsecValidateFrames to JSON text
	ToJson() (string, error)
}

type unMarshallagPortMacsecValidateFrames struct {
	obj *lagPortMacsecValidateFrames
}

type unMarshalLagPortMacsecValidateFrames interface {
	// FromProto unmarshals LagPortMacsecValidateFrames from protobuf object *otg.LagPortMacsecValidateFrames
	FromProto(msg *otg.LagPortMacsecValidateFrames) (LagPortMacsecValidateFrames, error)
	// FromPbText unmarshals LagPortMacsecValidateFrames from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagPortMacsecValidateFrames from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagPortMacsecValidateFrames from JSON text
	FromJson(value string) error
}

func (obj *lagPortMacsecValidateFrames) Marshal() marshalLagPortMacsecValidateFrames {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagPortMacsecValidateFrames{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagPortMacsecValidateFrames) Unmarshal() unMarshalLagPortMacsecValidateFrames {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagPortMacsecValidateFrames{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagPortMacsecValidateFrames) ToProto() (*otg.LagPortMacsecValidateFrames, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagPortMacsecValidateFrames) FromProto(msg *otg.LagPortMacsecValidateFrames) (LagPortMacsecValidateFrames, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagPortMacsecValidateFrames) ToPbText() (string, error) {
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

func (m *unMarshallagPortMacsecValidateFrames) FromPbText(value string) error {
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

func (m *marshallagPortMacsecValidateFrames) ToYaml() (string, error) {
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

func (m *unMarshallagPortMacsecValidateFrames) FromYaml(value string) error {
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

func (m *marshallagPortMacsecValidateFrames) ToJson() (string, error) {
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

func (m *unMarshallagPortMacsecValidateFrames) FromJson(value string) error {
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

func (obj *lagPortMacsecValidateFrames) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagPortMacsecValidateFrames) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagPortMacsecValidateFrames) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagPortMacsecValidateFrames) Clone() (LagPortMacsecValidateFrames, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagPortMacsecValidateFrames()
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

// LagPortMacsecValidateFrames is controls validation of received frames.
type LagPortMacsecValidateFrames interface {
	Validation
	// msg marshals LagPortMacsecValidateFrames to protobuf object *otg.LagPortMacsecValidateFrames
	// and doesn't set defaults
	msg() *otg.LagPortMacsecValidateFrames
	// setMsg unmarshals LagPortMacsecValidateFrames from protobuf object *otg.LagPortMacsecValidateFrames
	// and doesn't set defaults
	setMsg(*otg.LagPortMacsecValidateFrames) LagPortMacsecValidateFrames
	// provides marshal interface
	Marshal() marshalLagPortMacsecValidateFrames
	// provides unmarshal interface
	Unmarshal() unMarshalLagPortMacsecValidateFrames
	// validate validates LagPortMacsecValidateFrames
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagPortMacsecValidateFrames, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns LagPortMacsecValidateFramesChoiceEnum, set in LagPortMacsecValidateFrames
	Choice() LagPortMacsecValidateFramesChoiceEnum
	// setChoice assigns LagPortMacsecValidateFramesChoiceEnum provided by user to LagPortMacsecValidateFrames
	setChoice(value LagPortMacsecValidateFramesChoiceEnum) LagPortMacsecValidateFrames
	// HasChoice checks if Choice has been set in LagPortMacsecValidateFrames
	HasChoice() bool
	// getter for Check to set choice.
	Check()
	// getter for Strict to set choice.
	Strict()
}

type LagPortMacsecValidateFramesChoiceEnum string

// Enum of Choice on LagPortMacsecValidateFrames
var LagPortMacsecValidateFramesChoice = struct {
	CHECK  LagPortMacsecValidateFramesChoiceEnum
	STRICT LagPortMacsecValidateFramesChoiceEnum
}{
	CHECK:  LagPortMacsecValidateFramesChoiceEnum("check"),
	STRICT: LagPortMacsecValidateFramesChoiceEnum("strict"),
}

func (obj *lagPortMacsecValidateFrames) Choice() LagPortMacsecValidateFramesChoiceEnum {
	return LagPortMacsecValidateFramesChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Check to set choice
func (obj *lagPortMacsecValidateFrames) Check() {
	obj.setChoice(LagPortMacsecValidateFramesChoice.CHECK)
}

// getter for Strict to set choice
func (obj *lagPortMacsecValidateFrames) Strict() {
	obj.setChoice(LagPortMacsecValidateFramesChoice.STRICT)
}

// Controls validation of received frames. check - enable validation, do not discard invalid frames and increment in_pkts_invalid stats. strict - enable validation and discard invalid frames and increment in_pkts_not_valid stats.
// Choice returns a string
func (obj *lagPortMacsecValidateFrames) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *lagPortMacsecValidateFrames) setChoice(value LagPortMacsecValidateFramesChoiceEnum) LagPortMacsecValidateFrames {
	intValue, ok := otg.LagPortMacsecValidateFrames_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LagPortMacsecValidateFramesChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.LagPortMacsecValidateFrames_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *lagPortMacsecValidateFrames) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *lagPortMacsecValidateFrames) setDefault() {
	var choices_set int = 0
	var choice LagPortMacsecValidateFramesChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(LagPortMacsecValidateFramesChoice.CHECK)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in LagPortMacsecValidateFrames")
			}
		} else {
			intVal := otg.LagPortMacsecValidateFrames_Choice_Enum_value[string(choice)]
			enumValue := otg.LagPortMacsecValidateFrames_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

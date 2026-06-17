package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagPortMacsecSecureEntityDataPlaneRxValidateFrames *****
type lagPortMacsecSecureEntityDataPlaneRxValidateFrames struct {
	validation
	obj          *otg.LagPortMacsecSecureEntityDataPlaneRxValidateFrames
	marshaller   marshalLagPortMacsecSecureEntityDataPlaneRxValidateFrames
	unMarshaller unMarshalLagPortMacsecSecureEntityDataPlaneRxValidateFrames
}

func NewLagPortMacsecSecureEntityDataPlaneRxValidateFrames() LagPortMacsecSecureEntityDataPlaneRxValidateFrames {
	obj := lagPortMacsecSecureEntityDataPlaneRxValidateFrames{obj: &otg.LagPortMacsecSecureEntityDataPlaneRxValidateFrames{}}
	obj.setDefault()
	return &obj
}

func (obj *lagPortMacsecSecureEntityDataPlaneRxValidateFrames) msg() *otg.LagPortMacsecSecureEntityDataPlaneRxValidateFrames {
	return obj.obj
}

func (obj *lagPortMacsecSecureEntityDataPlaneRxValidateFrames) setMsg(msg *otg.LagPortMacsecSecureEntityDataPlaneRxValidateFrames) LagPortMacsecSecureEntityDataPlaneRxValidateFrames {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagPortMacsecSecureEntityDataPlaneRxValidateFrames struct {
	obj *lagPortMacsecSecureEntityDataPlaneRxValidateFrames
}

type marshalLagPortMacsecSecureEntityDataPlaneRxValidateFrames interface {
	// ToProto marshals LagPortMacsecSecureEntityDataPlaneRxValidateFrames to protobuf object *otg.LagPortMacsecSecureEntityDataPlaneRxValidateFrames
	ToProto() (*otg.LagPortMacsecSecureEntityDataPlaneRxValidateFrames, error)
	// ToPbText marshals LagPortMacsecSecureEntityDataPlaneRxValidateFrames to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagPortMacsecSecureEntityDataPlaneRxValidateFrames to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagPortMacsecSecureEntityDataPlaneRxValidateFrames to JSON text
	ToJson() (string, error)
}

type unMarshallagPortMacsecSecureEntityDataPlaneRxValidateFrames struct {
	obj *lagPortMacsecSecureEntityDataPlaneRxValidateFrames
}

type unMarshalLagPortMacsecSecureEntityDataPlaneRxValidateFrames interface {
	// FromProto unmarshals LagPortMacsecSecureEntityDataPlaneRxValidateFrames from protobuf object *otg.LagPortMacsecSecureEntityDataPlaneRxValidateFrames
	FromProto(msg *otg.LagPortMacsecSecureEntityDataPlaneRxValidateFrames) (LagPortMacsecSecureEntityDataPlaneRxValidateFrames, error)
	// FromPbText unmarshals LagPortMacsecSecureEntityDataPlaneRxValidateFrames from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagPortMacsecSecureEntityDataPlaneRxValidateFrames from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagPortMacsecSecureEntityDataPlaneRxValidateFrames from JSON text
	FromJson(value string) error
}

func (obj *lagPortMacsecSecureEntityDataPlaneRxValidateFrames) Marshal() marshalLagPortMacsecSecureEntityDataPlaneRxValidateFrames {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagPortMacsecSecureEntityDataPlaneRxValidateFrames{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagPortMacsecSecureEntityDataPlaneRxValidateFrames) Unmarshal() unMarshalLagPortMacsecSecureEntityDataPlaneRxValidateFrames {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagPortMacsecSecureEntityDataPlaneRxValidateFrames{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagPortMacsecSecureEntityDataPlaneRxValidateFrames) ToProto() (*otg.LagPortMacsecSecureEntityDataPlaneRxValidateFrames, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagPortMacsecSecureEntityDataPlaneRxValidateFrames) FromProto(msg *otg.LagPortMacsecSecureEntityDataPlaneRxValidateFrames) (LagPortMacsecSecureEntityDataPlaneRxValidateFrames, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagPortMacsecSecureEntityDataPlaneRxValidateFrames) ToPbText() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityDataPlaneRxValidateFrames) FromPbText(value string) error {
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

func (m *marshallagPortMacsecSecureEntityDataPlaneRxValidateFrames) ToYaml() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityDataPlaneRxValidateFrames) FromYaml(value string) error {
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

func (m *marshallagPortMacsecSecureEntityDataPlaneRxValidateFrames) ToJson() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityDataPlaneRxValidateFrames) FromJson(value string) error {
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

func (obj *lagPortMacsecSecureEntityDataPlaneRxValidateFrames) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagPortMacsecSecureEntityDataPlaneRxValidateFrames) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagPortMacsecSecureEntityDataPlaneRxValidateFrames) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagPortMacsecSecureEntityDataPlaneRxValidateFrames) Clone() (LagPortMacsecSecureEntityDataPlaneRxValidateFrames, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagPortMacsecSecureEntityDataPlaneRxValidateFrames()
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

// LagPortMacsecSecureEntityDataPlaneRxValidateFrames is controls validation of received frames.
type LagPortMacsecSecureEntityDataPlaneRxValidateFrames interface {
	Validation
	// msg marshals LagPortMacsecSecureEntityDataPlaneRxValidateFrames to protobuf object *otg.LagPortMacsecSecureEntityDataPlaneRxValidateFrames
	// and doesn't set defaults
	msg() *otg.LagPortMacsecSecureEntityDataPlaneRxValidateFrames
	// setMsg unmarshals LagPortMacsecSecureEntityDataPlaneRxValidateFrames from protobuf object *otg.LagPortMacsecSecureEntityDataPlaneRxValidateFrames
	// and doesn't set defaults
	setMsg(*otg.LagPortMacsecSecureEntityDataPlaneRxValidateFrames) LagPortMacsecSecureEntityDataPlaneRxValidateFrames
	// provides marshal interface
	Marshal() marshalLagPortMacsecSecureEntityDataPlaneRxValidateFrames
	// provides unmarshal interface
	Unmarshal() unMarshalLagPortMacsecSecureEntityDataPlaneRxValidateFrames
	// validate validates LagPortMacsecSecureEntityDataPlaneRxValidateFrames
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagPortMacsecSecureEntityDataPlaneRxValidateFrames, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns LagPortMacsecSecureEntityDataPlaneRxValidateFramesChoiceEnum, set in LagPortMacsecSecureEntityDataPlaneRxValidateFrames
	Choice() LagPortMacsecSecureEntityDataPlaneRxValidateFramesChoiceEnum
	// setChoice assigns LagPortMacsecSecureEntityDataPlaneRxValidateFramesChoiceEnum provided by user to LagPortMacsecSecureEntityDataPlaneRxValidateFrames
	setChoice(value LagPortMacsecSecureEntityDataPlaneRxValidateFramesChoiceEnum) LagPortMacsecSecureEntityDataPlaneRxValidateFrames
	// HasChoice checks if Choice has been set in LagPortMacsecSecureEntityDataPlaneRxValidateFrames
	HasChoice() bool
	// getter for Check to set choice.
	Check()
	// getter for Strict to set choice.
	Strict()
}

type LagPortMacsecSecureEntityDataPlaneRxValidateFramesChoiceEnum string

// Enum of Choice on LagPortMacsecSecureEntityDataPlaneRxValidateFrames
var LagPortMacsecSecureEntityDataPlaneRxValidateFramesChoice = struct {
	CHECK  LagPortMacsecSecureEntityDataPlaneRxValidateFramesChoiceEnum
	STRICT LagPortMacsecSecureEntityDataPlaneRxValidateFramesChoiceEnum
}{
	CHECK:  LagPortMacsecSecureEntityDataPlaneRxValidateFramesChoiceEnum("check"),
	STRICT: LagPortMacsecSecureEntityDataPlaneRxValidateFramesChoiceEnum("strict"),
}

func (obj *lagPortMacsecSecureEntityDataPlaneRxValidateFrames) Choice() LagPortMacsecSecureEntityDataPlaneRxValidateFramesChoiceEnum {
	return LagPortMacsecSecureEntityDataPlaneRxValidateFramesChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Check to set choice
func (obj *lagPortMacsecSecureEntityDataPlaneRxValidateFrames) Check() {
	obj.setChoice(LagPortMacsecSecureEntityDataPlaneRxValidateFramesChoice.CHECK)
}

// getter for Strict to set choice
func (obj *lagPortMacsecSecureEntityDataPlaneRxValidateFrames) Strict() {
	obj.setChoice(LagPortMacsecSecureEntityDataPlaneRxValidateFramesChoice.STRICT)
}

// Controls validation of received frames. check - enable validation, do not discard invalid frames and increment in_pkts_invalid stats. strict - enable validation and discard invalid frames and increment in_pkts_not_valid stats.
// Choice returns a string
func (obj *lagPortMacsecSecureEntityDataPlaneRxValidateFrames) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *lagPortMacsecSecureEntityDataPlaneRxValidateFrames) setChoice(value LagPortMacsecSecureEntityDataPlaneRxValidateFramesChoiceEnum) LagPortMacsecSecureEntityDataPlaneRxValidateFrames {
	intValue, ok := otg.LagPortMacsecSecureEntityDataPlaneRxValidateFrames_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LagPortMacsecSecureEntityDataPlaneRxValidateFramesChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.LagPortMacsecSecureEntityDataPlaneRxValidateFrames_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *lagPortMacsecSecureEntityDataPlaneRxValidateFrames) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *lagPortMacsecSecureEntityDataPlaneRxValidateFrames) setDefault() {
	var choices_set int = 0
	var choice LagPortMacsecSecureEntityDataPlaneRxValidateFramesChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(LagPortMacsecSecureEntityDataPlaneRxValidateFramesChoice.CHECK)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in LagPortMacsecSecureEntityDataPlaneRxValidateFrames")
			}
		} else {
			intVal := otg.LagPortMacsecSecureEntityDataPlaneRxValidateFrames_Choice_Enum_value[string(choice)]
			enumValue := otg.LagPortMacsecSecureEntityDataPlaneRxValidateFrames_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

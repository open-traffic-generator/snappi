package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmMDName *****
type flowCfmMDName struct {
	validation
	obj             *otg.FlowCfmMDName
	marshaller      marshalFlowCfmMDName
	unMarshaller    unMarshalFlowCfmMDName
	specifiedHolder FlowCfmMDNameSpecified
}

func NewFlowCfmMDName() FlowCfmMDName {
	obj := flowCfmMDName{obj: &otg.FlowCfmMDName{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmMDName) msg() *otg.FlowCfmMDName {
	return obj.obj
}

func (obj *flowCfmMDName) setMsg(msg *otg.FlowCfmMDName) FlowCfmMDName {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmMDName struct {
	obj *flowCfmMDName
}

type marshalFlowCfmMDName interface {
	// ToProto marshals FlowCfmMDName to protobuf object *otg.FlowCfmMDName
	ToProto() (*otg.FlowCfmMDName, error)
	// ToPbText marshals FlowCfmMDName to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmMDName to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmMDName to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmMDName struct {
	obj *flowCfmMDName
}

type unMarshalFlowCfmMDName interface {
	// FromProto unmarshals FlowCfmMDName from protobuf object *otg.FlowCfmMDName
	FromProto(msg *otg.FlowCfmMDName) (FlowCfmMDName, error)
	// FromPbText unmarshals FlowCfmMDName from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmMDName from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmMDName from JSON text
	FromJson(value string) error
}

func (obj *flowCfmMDName) Marshal() marshalFlowCfmMDName {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmMDName{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmMDName) Unmarshal() unMarshalFlowCfmMDName {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmMDName{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmMDName) ToProto() (*otg.FlowCfmMDName, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmMDName) FromProto(msg *otg.FlowCfmMDName) (FlowCfmMDName, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmMDName) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmMDName) FromPbText(value string) error {
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

func (m *marshalflowCfmMDName) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmMDName) FromYaml(value string) error {
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

func (m *marshalflowCfmMDName) ToJson() (string, error) {
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

func (m *unMarshalflowCfmMDName) FromJson(value string) error {
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

func (obj *flowCfmMDName) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmMDName) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmMDName) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmMDName) Clone() (FlowCfmMDName, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmMDName()
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

func (obj *flowCfmMDName) setNil() {
	obj.specifiedHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCfmMDName is maintenance domain name, can be auto or specified
type FlowCfmMDName interface {
	Validation
	// msg marshals FlowCfmMDName to protobuf object *otg.FlowCfmMDName
	// and doesn't set defaults
	msg() *otg.FlowCfmMDName
	// setMsg unmarshals FlowCfmMDName from protobuf object *otg.FlowCfmMDName
	// and doesn't set defaults
	setMsg(*otg.FlowCfmMDName) FlowCfmMDName
	// provides marshal interface
	Marshal() marshalFlowCfmMDName
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmMDName
	// validate validates FlowCfmMDName
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmMDName, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowCfmMDNameChoiceEnum, set in FlowCfmMDName
	Choice() FlowCfmMDNameChoiceEnum
	// setChoice assigns FlowCfmMDNameChoiceEnum provided by user to FlowCfmMDName
	setChoice(value FlowCfmMDNameChoiceEnum) FlowCfmMDName
	// HasChoice checks if Choice has been set in FlowCfmMDName
	HasChoice() bool
	// getter for Auto to set choice.
	Auto()
	// Specified returns FlowCfmMDNameSpecified, set in FlowCfmMDName.
	// FlowCfmMDNameSpecified is maintenance domain name
	Specified() FlowCfmMDNameSpecified
	// SetSpecified assigns FlowCfmMDNameSpecified provided by user to FlowCfmMDName.
	// FlowCfmMDNameSpecified is maintenance domain name
	SetSpecified(value FlowCfmMDNameSpecified) FlowCfmMDName
	// HasSpecified checks if Specified has been set in FlowCfmMDName
	HasSpecified() bool
	setNil()
}

type FlowCfmMDNameChoiceEnum string

// Enum of Choice on FlowCfmMDName
var FlowCfmMDNameChoice = struct {
	AUTO      FlowCfmMDNameChoiceEnum
	SPECIFIED FlowCfmMDNameChoiceEnum
}{
	AUTO:      FlowCfmMDNameChoiceEnum("auto"),
	SPECIFIED: FlowCfmMDNameChoiceEnum("specified"),
}

func (obj *flowCfmMDName) Choice() FlowCfmMDNameChoiceEnum {
	return FlowCfmMDNameChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Auto to set choice
func (obj *flowCfmMDName) Auto() {
	obj.setChoice(FlowCfmMDNameChoice.AUTO)
}

// description is TBD
// Choice returns a string
func (obj *flowCfmMDName) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowCfmMDName) setChoice(value FlowCfmMDNameChoiceEnum) FlowCfmMDName {
	intValue, ok := otg.FlowCfmMDName_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowCfmMDNameChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowCfmMDName_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Specified = nil
	obj.specifiedHolder = nil

	if value == FlowCfmMDNameChoice.SPECIFIED {
		obj.obj.Specified = NewFlowCfmMDNameSpecified().msg()
	}

	return obj
}

// description is TBD
// Specified returns a FlowCfmMDNameSpecified
func (obj *flowCfmMDName) Specified() FlowCfmMDNameSpecified {
	if obj.obj.Specified == nil {
		obj.setChoice(FlowCfmMDNameChoice.SPECIFIED)
	}
	if obj.specifiedHolder == nil {
		obj.specifiedHolder = &flowCfmMDNameSpecified{obj: obj.obj.Specified}
	}
	return obj.specifiedHolder
}

// description is TBD
// Specified returns a FlowCfmMDNameSpecified
func (obj *flowCfmMDName) HasSpecified() bool {
	return obj.obj.Specified != nil
}

// description is TBD
// SetSpecified sets the FlowCfmMDNameSpecified value in the FlowCfmMDName object
func (obj *flowCfmMDName) SetSpecified(value FlowCfmMDNameSpecified) FlowCfmMDName {
	obj.setChoice(FlowCfmMDNameChoice.SPECIFIED)
	obj.specifiedHolder = nil
	obj.obj.Specified = value.msg()

	return obj
}

func (obj *flowCfmMDName) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Specified != nil {

		obj.Specified().validateObj(vObj, set_default)
	}

}

func (obj *flowCfmMDName) setDefault() {
	var choices_set int = 0
	var choice FlowCfmMDNameChoiceEnum

	if obj.obj.Specified != nil {
		choices_set += 1
		choice = FlowCfmMDNameChoice.SPECIFIED
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowCfmMDNameChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowCfmMDName")
			}
		} else {
			intVal := otg.FlowCfmMDName_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowCfmMDName_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

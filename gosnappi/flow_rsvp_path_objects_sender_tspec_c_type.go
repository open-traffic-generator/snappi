package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjectsSenderTspecCType *****
type flowRSVPPathObjectsSenderTspecCType struct {
	validation
	obj           *otg.FlowRSVPPathObjectsSenderTspecCType
	marshaller    marshalFlowRSVPPathObjectsSenderTspecCType
	unMarshaller  unMarshalFlowRSVPPathObjectsSenderTspecCType
	intServHolder FlowRSVPPathSenderTspecIntServ
}

func NewFlowRSVPPathObjectsSenderTspecCType() FlowRSVPPathObjectsSenderTspecCType {
	obj := flowRSVPPathObjectsSenderTspecCType{obj: &otg.FlowRSVPPathObjectsSenderTspecCType{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjectsSenderTspecCType) msg() *otg.FlowRSVPPathObjectsSenderTspecCType {
	return obj.obj
}

func (obj *flowRSVPPathObjectsSenderTspecCType) setMsg(msg *otg.FlowRSVPPathObjectsSenderTspecCType) FlowRSVPPathObjectsSenderTspecCType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjectsSenderTspecCType struct {
	obj *flowRSVPPathObjectsSenderTspecCType
}

type marshalFlowRSVPPathObjectsSenderTspecCType interface {
	// ToProto marshals FlowRSVPPathObjectsSenderTspecCType to protobuf object *otg.FlowRSVPPathObjectsSenderTspecCType
	ToProto() (*otg.FlowRSVPPathObjectsSenderTspecCType, error)
	// ToPbText marshals FlowRSVPPathObjectsSenderTspecCType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjectsSenderTspecCType to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjectsSenderTspecCType to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathObjectsSenderTspecCType struct {
	obj *flowRSVPPathObjectsSenderTspecCType
}

type unMarshalFlowRSVPPathObjectsSenderTspecCType interface {
	// FromProto unmarshals FlowRSVPPathObjectsSenderTspecCType from protobuf object *otg.FlowRSVPPathObjectsSenderTspecCType
	FromProto(msg *otg.FlowRSVPPathObjectsSenderTspecCType) (FlowRSVPPathObjectsSenderTspecCType, error)
	// FromPbText unmarshals FlowRSVPPathObjectsSenderTspecCType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjectsSenderTspecCType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjectsSenderTspecCType from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjectsSenderTspecCType) Marshal() marshalFlowRSVPPathObjectsSenderTspecCType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjectsSenderTspecCType{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjectsSenderTspecCType) Unmarshal() unMarshalFlowRSVPPathObjectsSenderTspecCType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjectsSenderTspecCType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjectsSenderTspecCType) ToProto() (*otg.FlowRSVPPathObjectsSenderTspecCType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjectsSenderTspecCType) FromProto(msg *otg.FlowRSVPPathObjectsSenderTspecCType) (FlowRSVPPathObjectsSenderTspecCType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjectsSenderTspecCType) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsSenderTspecCType) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjectsSenderTspecCType) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsSenderTspecCType) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjectsSenderTspecCType) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsSenderTspecCType) FromJson(value string) error {
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

func (obj *flowRSVPPathObjectsSenderTspecCType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsSenderTspecCType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsSenderTspecCType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjectsSenderTspecCType) Clone() (FlowRSVPPathObjectsSenderTspecCType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjectsSenderTspecCType()
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

func (obj *flowRSVPPathObjectsSenderTspecCType) setNil() {
	obj.intServHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjectsSenderTspecCType is object for SENDER_TSPEC class. Currently supported c-type is int-serv (2).
type FlowRSVPPathObjectsSenderTspecCType interface {
	Validation
	// msg marshals FlowRSVPPathObjectsSenderTspecCType to protobuf object *otg.FlowRSVPPathObjectsSenderTspecCType
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjectsSenderTspecCType
	// setMsg unmarshals FlowRSVPPathObjectsSenderTspecCType from protobuf object *otg.FlowRSVPPathObjectsSenderTspecCType
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjectsSenderTspecCType) FlowRSVPPathObjectsSenderTspecCType
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjectsSenderTspecCType
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjectsSenderTspecCType
	// validate validates FlowRSVPPathObjectsSenderTspecCType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjectsSenderTspecCType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPPathObjectsSenderTspecCTypeChoiceEnum, set in FlowRSVPPathObjectsSenderTspecCType
	Choice() FlowRSVPPathObjectsSenderTspecCTypeChoiceEnum
	// setChoice assigns FlowRSVPPathObjectsSenderTspecCTypeChoiceEnum provided by user to FlowRSVPPathObjectsSenderTspecCType
	setChoice(value FlowRSVPPathObjectsSenderTspecCTypeChoiceEnum) FlowRSVPPathObjectsSenderTspecCType
	// HasChoice checks if Choice has been set in FlowRSVPPathObjectsSenderTspecCType
	HasChoice() bool
	// IntServ returns FlowRSVPPathSenderTspecIntServ, set in FlowRSVPPathObjectsSenderTspecCType.
	// FlowRSVPPathSenderTspecIntServ is int-serv SENDER_TSPEC object: Class = 12, C-Type = 2
	IntServ() FlowRSVPPathSenderTspecIntServ
	// SetIntServ assigns FlowRSVPPathSenderTspecIntServ provided by user to FlowRSVPPathObjectsSenderTspecCType.
	// FlowRSVPPathSenderTspecIntServ is int-serv SENDER_TSPEC object: Class = 12, C-Type = 2
	SetIntServ(value FlowRSVPPathSenderTspecIntServ) FlowRSVPPathObjectsSenderTspecCType
	// HasIntServ checks if IntServ has been set in FlowRSVPPathObjectsSenderTspecCType
	HasIntServ() bool
	setNil()
}

type FlowRSVPPathObjectsSenderTspecCTypeChoiceEnum string

// Enum of Choice on FlowRSVPPathObjectsSenderTspecCType
var FlowRSVPPathObjectsSenderTspecCTypeChoice = struct {
	INT_SERV FlowRSVPPathObjectsSenderTspecCTypeChoiceEnum
}{
	INT_SERV: FlowRSVPPathObjectsSenderTspecCTypeChoiceEnum("int_serv"),
}

func (obj *flowRSVPPathObjectsSenderTspecCType) Choice() FlowRSVPPathObjectsSenderTspecCTypeChoiceEnum {
	return FlowRSVPPathObjectsSenderTspecCTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowRSVPPathObjectsSenderTspecCType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPPathObjectsSenderTspecCType) setChoice(value FlowRSVPPathObjectsSenderTspecCTypeChoiceEnum) FlowRSVPPathObjectsSenderTspecCType {
	intValue, ok := otg.FlowRSVPPathObjectsSenderTspecCType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPPathObjectsSenderTspecCTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPPathObjectsSenderTspecCType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.IntServ = nil
	obj.intServHolder = nil

	if value == FlowRSVPPathObjectsSenderTspecCTypeChoice.INT_SERV {
		obj.obj.IntServ = NewFlowRSVPPathSenderTspecIntServ().msg()
	}

	return obj
}

// description is TBD
// IntServ returns a FlowRSVPPathSenderTspecIntServ
func (obj *flowRSVPPathObjectsSenderTspecCType) IntServ() FlowRSVPPathSenderTspecIntServ {
	if obj.obj.IntServ == nil {
		obj.setChoice(FlowRSVPPathObjectsSenderTspecCTypeChoice.INT_SERV)
	}
	if obj.intServHolder == nil {
		obj.intServHolder = &flowRSVPPathSenderTspecIntServ{obj: obj.obj.IntServ}
	}
	return obj.intServHolder
}

// description is TBD
// IntServ returns a FlowRSVPPathSenderTspecIntServ
func (obj *flowRSVPPathObjectsSenderTspecCType) HasIntServ() bool {
	return obj.obj.IntServ != nil
}

// description is TBD
// SetIntServ sets the FlowRSVPPathSenderTspecIntServ value in the FlowRSVPPathObjectsSenderTspecCType object
func (obj *flowRSVPPathObjectsSenderTspecCType) SetIntServ(value FlowRSVPPathSenderTspecIntServ) FlowRSVPPathObjectsSenderTspecCType {
	obj.setChoice(FlowRSVPPathObjectsSenderTspecCTypeChoice.INT_SERV)
	obj.intServHolder = nil
	obj.obj.IntServ = value.msg()

	return obj
}

func (obj *flowRSVPPathObjectsSenderTspecCType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.IntServ != nil {

		obj.IntServ().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathObjectsSenderTspecCType) setDefault() {
	var choices_set int = 0
	var choice FlowRSVPPathObjectsSenderTspecCTypeChoiceEnum

	if obj.obj.IntServ != nil {
		choices_set += 1
		choice = FlowRSVPPathObjectsSenderTspecCTypeChoice.INT_SERV
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowRSVPPathObjectsSenderTspecCTypeChoice.INT_SERV)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowRSVPPathObjectsSenderTspecCType")
			}
		} else {
			intVal := otg.FlowRSVPPathObjectsSenderTspecCType_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowRSVPPathObjectsSenderTspecCType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmOamProtocol *****
type flowCfmOamProtocol struct {
	validation
	obj          *otg.FlowCfmOamProtocol
	marshaller   marshalFlowCfmOamProtocol
	unMarshaller unMarshalFlowCfmOamProtocol
	cfmHolder    FlowCfmCfm
}

func NewFlowCfmOamProtocol() FlowCfmOamProtocol {
	obj := flowCfmOamProtocol{obj: &otg.FlowCfmOamProtocol{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmOamProtocol) msg() *otg.FlowCfmOamProtocol {
	return obj.obj
}

func (obj *flowCfmOamProtocol) setMsg(msg *otg.FlowCfmOamProtocol) FlowCfmOamProtocol {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmOamProtocol struct {
	obj *flowCfmOamProtocol
}

type marshalFlowCfmOamProtocol interface {
	// ToProto marshals FlowCfmOamProtocol to protobuf object *otg.FlowCfmOamProtocol
	ToProto() (*otg.FlowCfmOamProtocol, error)
	// ToPbText marshals FlowCfmOamProtocol to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmOamProtocol to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmOamProtocol to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmOamProtocol struct {
	obj *flowCfmOamProtocol
}

type unMarshalFlowCfmOamProtocol interface {
	// FromProto unmarshals FlowCfmOamProtocol from protobuf object *otg.FlowCfmOamProtocol
	FromProto(msg *otg.FlowCfmOamProtocol) (FlowCfmOamProtocol, error)
	// FromPbText unmarshals FlowCfmOamProtocol from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmOamProtocol from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmOamProtocol from JSON text
	FromJson(value string) error
}

func (obj *flowCfmOamProtocol) Marshal() marshalFlowCfmOamProtocol {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmOamProtocol{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmOamProtocol) Unmarshal() unMarshalFlowCfmOamProtocol {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmOamProtocol{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmOamProtocol) ToProto() (*otg.FlowCfmOamProtocol, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmOamProtocol) FromProto(msg *otg.FlowCfmOamProtocol) (FlowCfmOamProtocol, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmOamProtocol) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmOamProtocol) FromPbText(value string) error {
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

func (m *marshalflowCfmOamProtocol) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmOamProtocol) FromYaml(value string) error {
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

func (m *marshalflowCfmOamProtocol) ToJson() (string, error) {
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

func (m *unMarshalflowCfmOamProtocol) FromJson(value string) error {
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

func (obj *flowCfmOamProtocol) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmOamProtocol) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmOamProtocol) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmOamProtocol) Clone() (FlowCfmOamProtocol, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmOamProtocol()
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

func (obj *flowCfmOamProtocol) setNil() {
	obj.cfmHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCfmOamProtocol is ethernet OAM standard. Currently only Connectivity Fault Management (CFM) is supported.
type FlowCfmOamProtocol interface {
	Validation
	// msg marshals FlowCfmOamProtocol to protobuf object *otg.FlowCfmOamProtocol
	// and doesn't set defaults
	msg() *otg.FlowCfmOamProtocol
	// setMsg unmarshals FlowCfmOamProtocol from protobuf object *otg.FlowCfmOamProtocol
	// and doesn't set defaults
	setMsg(*otg.FlowCfmOamProtocol) FlowCfmOamProtocol
	// provides marshal interface
	Marshal() marshalFlowCfmOamProtocol
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmOamProtocol
	// validate validates FlowCfmOamProtocol
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmOamProtocol, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowCfmOamProtocolChoiceEnum, set in FlowCfmOamProtocol
	Choice() FlowCfmOamProtocolChoiceEnum
	// setChoice assigns FlowCfmOamProtocolChoiceEnum provided by user to FlowCfmOamProtocol
	setChoice(value FlowCfmOamProtocolChoiceEnum) FlowCfmOamProtocol
	// HasChoice checks if Choice has been set in FlowCfmOamProtocol
	HasChoice() bool
	// Cfm returns FlowCfmCfm, set in FlowCfmOamProtocol.
	// FlowCfmCfm is connectivity Fault Management
	Cfm() FlowCfmCfm
	// SetCfm assigns FlowCfmCfm provided by user to FlowCfmOamProtocol.
	// FlowCfmCfm is connectivity Fault Management
	SetCfm(value FlowCfmCfm) FlowCfmOamProtocol
	// HasCfm checks if Cfm has been set in FlowCfmOamProtocol
	HasCfm() bool
	setNil()
}

type FlowCfmOamProtocolChoiceEnum string

// Enum of Choice on FlowCfmOamProtocol
var FlowCfmOamProtocolChoice = struct {
	CFM FlowCfmOamProtocolChoiceEnum
}{
	CFM: FlowCfmOamProtocolChoiceEnum("cfm"),
}

func (obj *flowCfmOamProtocol) Choice() FlowCfmOamProtocolChoiceEnum {
	return FlowCfmOamProtocolChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowCfmOamProtocol) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowCfmOamProtocol) setChoice(value FlowCfmOamProtocolChoiceEnum) FlowCfmOamProtocol {
	intValue, ok := otg.FlowCfmOamProtocol_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowCfmOamProtocolChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowCfmOamProtocol_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Cfm = nil
	obj.cfmHolder = nil

	if value == FlowCfmOamProtocolChoice.CFM {
		obj.obj.Cfm = NewFlowCfmCfm().msg()
	}

	return obj
}

// description is TBD
// Cfm returns a FlowCfmCfm
func (obj *flowCfmOamProtocol) Cfm() FlowCfmCfm {
	if obj.obj.Cfm == nil {
		obj.setChoice(FlowCfmOamProtocolChoice.CFM)
	}
	if obj.cfmHolder == nil {
		obj.cfmHolder = &flowCfmCfm{obj: obj.obj.Cfm}
	}
	return obj.cfmHolder
}

// description is TBD
// Cfm returns a FlowCfmCfm
func (obj *flowCfmOamProtocol) HasCfm() bool {
	return obj.obj.Cfm != nil
}

// description is TBD
// SetCfm sets the FlowCfmCfm value in the FlowCfmOamProtocol object
func (obj *flowCfmOamProtocol) SetCfm(value FlowCfmCfm) FlowCfmOamProtocol {
	obj.setChoice(FlowCfmOamProtocolChoice.CFM)
	obj.cfmHolder = nil
	obj.obj.Cfm = value.msg()

	return obj
}

func (obj *flowCfmOamProtocol) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Cfm != nil {

		obj.Cfm().validateObj(vObj, set_default)
	}

}

func (obj *flowCfmOamProtocol) setDefault() {
	var choices_set int = 0
	var choice FlowCfmOamProtocolChoiceEnum

	if obj.obj.Cfm != nil {
		choices_set += 1
		choice = FlowCfmOamProtocolChoice.CFM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowCfmOamProtocolChoice.CFM)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowCfmOamProtocol")
			}
		} else {
			intVal := otg.FlowCfmOamProtocol_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowCfmOamProtocol_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ConfigUpdate *****
type configUpdate struct {
	validation
	obj          *otg.ConfigUpdate
	marshaller   marshalConfigUpdate
	unMarshaller unMarshalConfigUpdate
	flowsHolder  FlowsUpdate
}

func NewConfigUpdate() ConfigUpdate {
	obj := configUpdate{obj: &otg.ConfigUpdate{}}
	obj.setDefault()
	return &obj
}

func (obj *configUpdate) msg() *otg.ConfigUpdate {
	return obj.obj
}

func (obj *configUpdate) setMsg(msg *otg.ConfigUpdate) ConfigUpdate {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalconfigUpdate struct {
	obj *configUpdate
}

type marshalConfigUpdate interface {
	// ToProto marshals ConfigUpdate to protobuf object *otg.ConfigUpdate
	ToProto() (*otg.ConfigUpdate, error)
	// ToPbText marshals ConfigUpdate to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ConfigUpdate to YAML text
	ToYaml() (string, error)
	// ToJson marshals ConfigUpdate to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals ConfigUpdate to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalconfigUpdate struct {
	obj *configUpdate
}

type unMarshalConfigUpdate interface {
	// FromProto unmarshals ConfigUpdate from protobuf object *otg.ConfigUpdate
	FromProto(msg *otg.ConfigUpdate) (ConfigUpdate, error)
	// FromPbText unmarshals ConfigUpdate from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ConfigUpdate from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ConfigUpdate from JSON text
	FromJson(value string) error
}

func (obj *configUpdate) Marshal() marshalConfigUpdate {
	if obj.marshaller == nil {
		obj.marshaller = &marshalconfigUpdate{obj: obj}
	}
	return obj.marshaller
}

func (obj *configUpdate) Unmarshal() unMarshalConfigUpdate {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalconfigUpdate{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalconfigUpdate) ToProto() (*otg.ConfigUpdate, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalconfigUpdate) FromProto(msg *otg.ConfigUpdate) (ConfigUpdate, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalconfigUpdate) ToPbText() (string, error) {
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

func (m *unMarshalconfigUpdate) FromPbText(value string) error {
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

func (m *marshalconfigUpdate) ToYaml() (string, error) {
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

func (m *unMarshalconfigUpdate) FromYaml(value string) error {
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

func (m *marshalconfigUpdate) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalconfigUpdate) ToJson() (string, error) {
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

func (m *unMarshalconfigUpdate) FromJson(value string) error {
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

func (obj *configUpdate) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *configUpdate) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *configUpdate) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *configUpdate) Clone() (ConfigUpdate, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewConfigUpdate()
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

func (obj *configUpdate) setNil() {
	obj.flowsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ConfigUpdate is request for updating specific attributes of resources in traffic generator
type ConfigUpdate interface {
	Validation
	// msg marshals ConfigUpdate to protobuf object *otg.ConfigUpdate
	// and doesn't set defaults
	msg() *otg.ConfigUpdate
	// setMsg unmarshals ConfigUpdate from protobuf object *otg.ConfigUpdate
	// and doesn't set defaults
	setMsg(*otg.ConfigUpdate) ConfigUpdate
	// provides marshal interface
	Marshal() marshalConfigUpdate
	// provides unmarshal interface
	Unmarshal() unMarshalConfigUpdate
	// validate validates ConfigUpdate
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ConfigUpdate, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ConfigUpdateChoiceEnum, set in ConfigUpdate
	Choice() ConfigUpdateChoiceEnum
	// setChoice assigns ConfigUpdateChoiceEnum provided by user to ConfigUpdate
	setChoice(value ConfigUpdateChoiceEnum) ConfigUpdate
	// HasChoice checks if Choice has been set in ConfigUpdate
	HasChoice() bool
	// Flows returns FlowsUpdate, set in ConfigUpdate.
	// FlowsUpdate is a container of flows with associated properties to be updated without affecting the flows current transmit state.
	Flows() FlowsUpdate
	// SetFlows assigns FlowsUpdate provided by user to ConfigUpdate.
	// FlowsUpdate is a container of flows with associated properties to be updated without affecting the flows current transmit state.
	SetFlows(value FlowsUpdate) ConfigUpdate
	// HasFlows checks if Flows has been set in ConfigUpdate
	HasFlows() bool
	setNil()
}

type ConfigUpdateChoiceEnum string

// Enum of Choice on ConfigUpdate
var ConfigUpdateChoice = struct {
	FLOWS ConfigUpdateChoiceEnum
}{
	FLOWS: ConfigUpdateChoiceEnum("flows"),
}

func (obj *configUpdate) Choice() ConfigUpdateChoiceEnum {
	return ConfigUpdateChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *configUpdate) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *configUpdate) setChoice(value ConfigUpdateChoiceEnum) ConfigUpdate {
	intValue, ok := otg.ConfigUpdate_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ConfigUpdateChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ConfigUpdate_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Flows = nil
	obj.flowsHolder = nil

	if value == ConfigUpdateChoice.FLOWS {
		obj.obj.Flows = NewFlowsUpdate().msg()
	}

	return obj
}

// description is TBD
// Flows returns a FlowsUpdate
func (obj *configUpdate) Flows() FlowsUpdate {
	if obj.obj.Flows == nil {
		obj.setChoice(ConfigUpdateChoice.FLOWS)
	}
	if obj.flowsHolder == nil {
		obj.flowsHolder = &flowsUpdate{obj: obj.obj.Flows}
	}
	return obj.flowsHolder
}

// description is TBD
// Flows returns a FlowsUpdate
func (obj *configUpdate) HasFlows() bool {
	return obj.obj.Flows != nil
}

// description is TBD
// SetFlows sets the FlowsUpdate value in the ConfigUpdate object
func (obj *configUpdate) SetFlows(value FlowsUpdate) ConfigUpdate {
	obj.setChoice(ConfigUpdateChoice.FLOWS)
	obj.flowsHolder = nil
	obj.obj.Flows = value.msg()

	return obj
}

func (obj *configUpdate) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flows != nil {

		obj.Flows().validateObj(vObj, set_default)
	}

}

func (obj *configUpdate) setDefault() {
	var choices_set int = 0
	var choice ConfigUpdateChoiceEnum

	if obj.obj.Flows != nil {
		choices_set += 1
		choice = ConfigUpdateChoice.FLOWS
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ConfigUpdate")
			}
		} else {
			intVal := otg.ConfigUpdate_Choice_Enum_value[string(choice)]
			enumValue := otg.ConfigUpdate_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

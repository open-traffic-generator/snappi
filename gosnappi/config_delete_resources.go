package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ConfigDeleteResources *****
type configDeleteResources struct {
	validation
	obj          *otg.ConfigDeleteResources
	marshaller   marshalConfigDeleteResources
	unMarshaller unMarshalConfigDeleteResources
}

func NewConfigDeleteResources() ConfigDeleteResources {
	obj := configDeleteResources{obj: &otg.ConfigDeleteResources{}}
	obj.setDefault()
	return &obj
}

func (obj *configDeleteResources) msg() *otg.ConfigDeleteResources {
	return obj.obj
}

func (obj *configDeleteResources) setMsg(msg *otg.ConfigDeleteResources) ConfigDeleteResources {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalconfigDeleteResources struct {
	obj *configDeleteResources
}

type marshalConfigDeleteResources interface {
	// ToProto marshals ConfigDeleteResources to protobuf object *otg.ConfigDeleteResources
	ToProto() (*otg.ConfigDeleteResources, error)
	// ToPbText marshals ConfigDeleteResources to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ConfigDeleteResources to YAML text
	ToYaml() (string, error)
	// ToJson marshals ConfigDeleteResources to JSON text
	ToJson() (string, error)
}

type unMarshalconfigDeleteResources struct {
	obj *configDeleteResources
}

type unMarshalConfigDeleteResources interface {
	// FromProto unmarshals ConfigDeleteResources from protobuf object *otg.ConfigDeleteResources
	FromProto(msg *otg.ConfigDeleteResources) (ConfigDeleteResources, error)
	// FromPbText unmarshals ConfigDeleteResources from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ConfigDeleteResources from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ConfigDeleteResources from JSON text
	FromJson(value string) error
}

func (obj *configDeleteResources) Marshal() marshalConfigDeleteResources {
	if obj.marshaller == nil {
		obj.marshaller = &marshalconfigDeleteResources{obj: obj}
	}
	return obj.marshaller
}

func (obj *configDeleteResources) Unmarshal() unMarshalConfigDeleteResources {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalconfigDeleteResources{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalconfigDeleteResources) ToProto() (*otg.ConfigDeleteResources, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalconfigDeleteResources) FromProto(msg *otg.ConfigDeleteResources) (ConfigDeleteResources, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalconfigDeleteResources) ToPbText() (string, error) {
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

func (m *unMarshalconfigDeleteResources) FromPbText(value string) error {
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

func (m *marshalconfigDeleteResources) ToYaml() (string, error) {
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

func (m *unMarshalconfigDeleteResources) FromYaml(value string) error {
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

func (m *marshalconfigDeleteResources) ToJson() (string, error) {
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

func (m *unMarshalconfigDeleteResources) FromJson(value string) error {
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

func (obj *configDeleteResources) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *configDeleteResources) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *configDeleteResources) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *configDeleteResources) Clone() (ConfigDeleteResources, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewConfigDeleteResources()
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

// ConfigDeleteResources is a container for an existing resource to be deleted.
type ConfigDeleteResources interface {
	Validation
	// msg marshals ConfigDeleteResources to protobuf object *otg.ConfigDeleteResources
	// and doesn't set defaults
	msg() *otg.ConfigDeleteResources
	// setMsg unmarshals ConfigDeleteResources from protobuf object *otg.ConfigDeleteResources
	// and doesn't set defaults
	setMsg(*otg.ConfigDeleteResources) ConfigDeleteResources
	// provides marshal interface
	Marshal() marshalConfigDeleteResources
	// provides unmarshal interface
	Unmarshal() unMarshalConfigDeleteResources
	// validate validates ConfigDeleteResources
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ConfigDeleteResources, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ConfigDeleteResourcesChoiceEnum, set in ConfigDeleteResources
	Choice() ConfigDeleteResourcesChoiceEnum
	// setChoice assigns ConfigDeleteResourcesChoiceEnum provided by user to ConfigDeleteResources
	setChoice(value ConfigDeleteResourcesChoiceEnum) ConfigDeleteResources
	// Flows returns []string, set in ConfigDeleteResources.
	Flows() []string
	// SetFlows assigns []string provided by user to ConfigDeleteResources
	SetFlows(value []string) ConfigDeleteResources
}

type ConfigDeleteResourcesChoiceEnum string

// Enum of Choice on ConfigDeleteResources
var ConfigDeleteResourcesChoice = struct {
	FLOWS ConfigDeleteResourcesChoiceEnum
}{
	FLOWS: ConfigDeleteResourcesChoiceEnum("flows"),
}

func (obj *configDeleteResources) Choice() ConfigDeleteResourcesChoiceEnum {
	return ConfigDeleteResourcesChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *configDeleteResources) setChoice(value ConfigDeleteResourcesChoiceEnum) ConfigDeleteResources {
	intValue, ok := otg.ConfigDeleteResources_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ConfigDeleteResourcesChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ConfigDeleteResources_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Flows = nil
	return obj
}

// List of flows that will be deleted from existing configuration on the traffic generator.
//
// x-constraint:
// - /components/schemas/Flow/properties/name
//
// x-constraint:
// - /components/schemas/Flow/properties/name
//
// Flows returns a []string
func (obj *configDeleteResources) Flows() []string {
	if obj.obj.Flows == nil {

		obj.setChoice(ConfigDeleteResourcesChoice.FLOWS)

	}
	return obj.obj.Flows
}

// List of flows that will be deleted from existing configuration on the traffic generator.
//
// x-constraint:
// - /components/schemas/Flow/properties/name
//
// x-constraint:
// - /components/schemas/Flow/properties/name
//
// SetFlows sets the []string value in the ConfigDeleteResources object
func (obj *configDeleteResources) SetFlows(value []string) ConfigDeleteResources {
	obj.setChoice(ConfigDeleteResourcesChoice.FLOWS)
	if obj.obj.Flows == nil {
		obj.obj.Flows = make([]string, 0)
	}
	obj.obj.Flows = value

	return obj
}

func (obj *configDeleteResources) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface ConfigDeleteResources")
	}
}

func (obj *configDeleteResources) setDefault() {
	var choices_set int = 0
	var choice ConfigDeleteResourcesChoiceEnum

	if len(obj.obj.Flows) > 0 {
		choices_set += 1
		choice = ConfigDeleteResourcesChoice.FLOWS
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ConfigDeleteResources")
			}
		} else {
			intVal := otg.ConfigDeleteResources_Choice_Enum_value[string(choice)]
			enumValue := otg.ConfigDeleteResources_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

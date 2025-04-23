package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ConfigDeleteResource *****
type configDeleteResource struct {
	validation
	obj          *otg.ConfigDeleteResource
	marshaller   marshalConfigDeleteResource
	unMarshaller unMarshalConfigDeleteResource
}

func NewConfigDeleteResource() ConfigDeleteResource {
	obj := configDeleteResource{obj: &otg.ConfigDeleteResource{}}
	obj.setDefault()
	return &obj
}

func (obj *configDeleteResource) msg() *otg.ConfigDeleteResource {
	return obj.obj
}

func (obj *configDeleteResource) setMsg(msg *otg.ConfigDeleteResource) ConfigDeleteResource {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalconfigDeleteResource struct {
	obj *configDeleteResource
}

type marshalConfigDeleteResource interface {
	// ToProto marshals ConfigDeleteResource to protobuf object *otg.ConfigDeleteResource
	ToProto() (*otg.ConfigDeleteResource, error)
	// ToPbText marshals ConfigDeleteResource to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ConfigDeleteResource to YAML text
	ToYaml() (string, error)
	// ToJson marshals ConfigDeleteResource to JSON text
	ToJson() (string, error)
}

type unMarshalconfigDeleteResource struct {
	obj *configDeleteResource
}

type unMarshalConfigDeleteResource interface {
	// FromProto unmarshals ConfigDeleteResource from protobuf object *otg.ConfigDeleteResource
	FromProto(msg *otg.ConfigDeleteResource) (ConfigDeleteResource, error)
	// FromPbText unmarshals ConfigDeleteResource from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ConfigDeleteResource from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ConfigDeleteResource from JSON text
	FromJson(value string) error
}

func (obj *configDeleteResource) Marshal() marshalConfigDeleteResource {
	if obj.marshaller == nil {
		obj.marshaller = &marshalconfigDeleteResource{obj: obj}
	}
	return obj.marshaller
}

func (obj *configDeleteResource) Unmarshal() unMarshalConfigDeleteResource {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalconfigDeleteResource{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalconfigDeleteResource) ToProto() (*otg.ConfigDeleteResource, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalconfigDeleteResource) FromProto(msg *otg.ConfigDeleteResource) (ConfigDeleteResource, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalconfigDeleteResource) ToPbText() (string, error) {
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

func (m *unMarshalconfigDeleteResource) FromPbText(value string) error {
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

func (m *marshalconfigDeleteResource) ToYaml() (string, error) {
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

func (m *unMarshalconfigDeleteResource) FromYaml(value string) error {
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

func (m *marshalconfigDeleteResource) ToJson() (string, error) {
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

func (m *unMarshalconfigDeleteResource) FromJson(value string) error {
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

func (obj *configDeleteResource) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *configDeleteResource) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *configDeleteResource) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *configDeleteResource) Clone() (ConfigDeleteResource, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewConfigDeleteResource()
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

// ConfigDeleteResource is a container for an existing resource to be deleted.
type ConfigDeleteResource interface {
	Validation
	// msg marshals ConfigDeleteResource to protobuf object *otg.ConfigDeleteResource
	// and doesn't set defaults
	msg() *otg.ConfigDeleteResource
	// setMsg unmarshals ConfigDeleteResource from protobuf object *otg.ConfigDeleteResource
	// and doesn't set defaults
	setMsg(*otg.ConfigDeleteResource) ConfigDeleteResource
	// provides marshal interface
	Marshal() marshalConfigDeleteResource
	// provides unmarshal interface
	Unmarshal() unMarshalConfigDeleteResource
	// validate validates ConfigDeleteResource
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ConfigDeleteResource, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ConfigDeleteResourceChoiceEnum, set in ConfigDeleteResource
	Choice() ConfigDeleteResourceChoiceEnum
	// setChoice assigns ConfigDeleteResourceChoiceEnum provided by user to ConfigDeleteResource
	setChoice(value ConfigDeleteResourceChoiceEnum) ConfigDeleteResource
	// Flows returns []string, set in ConfigDeleteResource.
	Flows() []string
	// SetFlows assigns []string provided by user to ConfigDeleteResource
	SetFlows(value []string) ConfigDeleteResource
}

type ConfigDeleteResourceChoiceEnum string

// Enum of Choice on ConfigDeleteResource
var ConfigDeleteResourceChoice = struct {
	FLOWS ConfigDeleteResourceChoiceEnum
}{
	FLOWS: ConfigDeleteResourceChoiceEnum("flows"),
}

func (obj *configDeleteResource) Choice() ConfigDeleteResourceChoiceEnum {
	return ConfigDeleteResourceChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *configDeleteResource) setChoice(value ConfigDeleteResourceChoiceEnum) ConfigDeleteResource {
	intValue, ok := otg.ConfigDeleteResource_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ConfigDeleteResourceChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ConfigDeleteResource_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Flows = nil
	return obj
}

// List of flows that will be deleted from existing configuration on the traffic generator.
//
// x-constraint:
// - /components/schemas/Flow/properties/name
//
// Flows returns a []string
func (obj *configDeleteResource) Flows() []string {
	if obj.obj.Flows == nil {

		obj.setChoice(ConfigDeleteResourceChoice.FLOWS)

	}
	return obj.obj.Flows
}

// List of flows that will be deleted from existing configuration on the traffic generator.
//
// x-constraint:
// - /components/schemas/Flow/properties/name
//
// SetFlows sets the []string value in the ConfigDeleteResource object
func (obj *configDeleteResource) SetFlows(value []string) ConfigDeleteResource {
	obj.setChoice(ConfigDeleteResourceChoice.FLOWS)
	if obj.obj.Flows == nil {
		obj.obj.Flows = make([]string, 0)
	}
	obj.obj.Flows = value

	return obj
}

func (obj *configDeleteResource) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface ConfigDeleteResource")
	}
}

func (obj *configDeleteResource) setDefault() {
	var choices_set int = 0
	var choice ConfigDeleteResourceChoiceEnum

	if len(obj.obj.Flows) > 0 {
		choices_set += 1
		choice = ConfigDeleteResourceChoice.FLOWS
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ConfigDeleteResource")
			}
		} else {
			intVal := otg.ConfigDeleteResource_Choice_Enum_value[string(choice)]
			enumValue := otg.ConfigDeleteResource_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

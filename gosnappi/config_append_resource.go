package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ConfigAppendResource *****
type configAppendResource struct {
	validation
	obj          *otg.ConfigAppendResource
	marshaller   marshalConfigAppendResource
	unMarshaller unMarshalConfigAppendResource
	flowsHolder  ConfigAppendResourceFlowIter
}

func NewConfigAppendResource() ConfigAppendResource {
	obj := configAppendResource{obj: &otg.ConfigAppendResource{}}
	obj.setDefault()
	return &obj
}

func (obj *configAppendResource) msg() *otg.ConfigAppendResource {
	return obj.obj
}

func (obj *configAppendResource) setMsg(msg *otg.ConfigAppendResource) ConfigAppendResource {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalconfigAppendResource struct {
	obj *configAppendResource
}

type marshalConfigAppendResource interface {
	// ToProto marshals ConfigAppendResource to protobuf object *otg.ConfigAppendResource
	ToProto() (*otg.ConfigAppendResource, error)
	// ToPbText marshals ConfigAppendResource to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ConfigAppendResource to YAML text
	ToYaml() (string, error)
	// ToJson marshals ConfigAppendResource to JSON text
	ToJson() (string, error)
}

type unMarshalconfigAppendResource struct {
	obj *configAppendResource
}

type unMarshalConfigAppendResource interface {
	// FromProto unmarshals ConfigAppendResource from protobuf object *otg.ConfigAppendResource
	FromProto(msg *otg.ConfigAppendResource) (ConfigAppendResource, error)
	// FromPbText unmarshals ConfigAppendResource from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ConfigAppendResource from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ConfigAppendResource from JSON text
	FromJson(value string) error
}

func (obj *configAppendResource) Marshal() marshalConfigAppendResource {
	if obj.marshaller == nil {
		obj.marshaller = &marshalconfigAppendResource{obj: obj}
	}
	return obj.marshaller
}

func (obj *configAppendResource) Unmarshal() unMarshalConfigAppendResource {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalconfigAppendResource{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalconfigAppendResource) ToProto() (*otg.ConfigAppendResource, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalconfigAppendResource) FromProto(msg *otg.ConfigAppendResource) (ConfigAppendResource, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalconfigAppendResource) ToPbText() (string, error) {
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

func (m *unMarshalconfigAppendResource) FromPbText(value string) error {
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

func (m *marshalconfigAppendResource) ToYaml() (string, error) {
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

func (m *unMarshalconfigAppendResource) FromYaml(value string) error {
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

func (m *marshalconfigAppendResource) ToJson() (string, error) {
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

func (m *unMarshalconfigAppendResource) FromJson(value string) error {
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

func (obj *configAppendResource) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *configAppendResource) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *configAppendResource) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *configAppendResource) Clone() (ConfigAppendResource, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewConfigAppendResource()
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

func (obj *configAppendResource) setNil() {
	obj.flowsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ConfigAppendResource is a container for an existing resource to be appended.
type ConfigAppendResource interface {
	Validation
	// msg marshals ConfigAppendResource to protobuf object *otg.ConfigAppendResource
	// and doesn't set defaults
	msg() *otg.ConfigAppendResource
	// setMsg unmarshals ConfigAppendResource from protobuf object *otg.ConfigAppendResource
	// and doesn't set defaults
	setMsg(*otg.ConfigAppendResource) ConfigAppendResource
	// provides marshal interface
	Marshal() marshalConfigAppendResource
	// provides unmarshal interface
	Unmarshal() unMarshalConfigAppendResource
	// validate validates ConfigAppendResource
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ConfigAppendResource, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ConfigAppendResourceChoiceEnum, set in ConfigAppendResource
	Choice() ConfigAppendResourceChoiceEnum
	// setChoice assigns ConfigAppendResourceChoiceEnum provided by user to ConfigAppendResource
	setChoice(value ConfigAppendResourceChoiceEnum) ConfigAppendResource
	// Flows returns ConfigAppendResourceFlowIterIter, set in ConfigAppendResource
	Flows() ConfigAppendResourceFlowIter
	setNil()
}

type ConfigAppendResourceChoiceEnum string

// Enum of Choice on ConfigAppendResource
var ConfigAppendResourceChoice = struct {
	FLOWS ConfigAppendResourceChoiceEnum
}{
	FLOWS: ConfigAppendResourceChoiceEnum("flows"),
}

func (obj *configAppendResource) Choice() ConfigAppendResourceChoiceEnum {
	return ConfigAppendResourceChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *configAppendResource) setChoice(value ConfigAppendResourceChoiceEnum) ConfigAppendResource {
	intValue, ok := otg.ConfigAppendResource_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ConfigAppendResourceChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ConfigAppendResource_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Flows = nil
	obj.flowsHolder = nil

	if value == ConfigAppendResourceChoice.FLOWS {
		obj.obj.Flows = []*otg.Flow{}
	}

	return obj
}

// The flows that will be added to existing configuration on the traffic generator.
// The flow name must not already be used in existing configuration.
// Flows returns a []Flow
func (obj *configAppendResource) Flows() ConfigAppendResourceFlowIter {
	if len(obj.obj.Flows) == 0 {
		obj.setChoice(ConfigAppendResourceChoice.FLOWS)
	}
	if obj.flowsHolder == nil {
		obj.flowsHolder = newConfigAppendResourceFlowIter(&obj.obj.Flows).setMsg(obj)
	}
	return obj.flowsHolder
}

type configAppendResourceFlowIter struct {
	obj       *configAppendResource
	flowSlice []Flow
	fieldPtr  *[]*otg.Flow
}

func newConfigAppendResourceFlowIter(ptr *[]*otg.Flow) ConfigAppendResourceFlowIter {
	return &configAppendResourceFlowIter{fieldPtr: ptr}
}

type ConfigAppendResourceFlowIter interface {
	setMsg(*configAppendResource) ConfigAppendResourceFlowIter
	Items() []Flow
	Add() Flow
	Append(items ...Flow) ConfigAppendResourceFlowIter
	Set(index int, newObj Flow) ConfigAppendResourceFlowIter
	Clear() ConfigAppendResourceFlowIter
	clearHolderSlice() ConfigAppendResourceFlowIter
	appendHolderSlice(item Flow) ConfigAppendResourceFlowIter
}

func (obj *configAppendResourceFlowIter) setMsg(msg *configAppendResource) ConfigAppendResourceFlowIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flow{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *configAppendResourceFlowIter) Items() []Flow {
	return obj.flowSlice
}

func (obj *configAppendResourceFlowIter) Add() Flow {
	newObj := &otg.Flow{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flow{obj: newObj}
	newLibObj.setDefault()
	obj.flowSlice = append(obj.flowSlice, newLibObj)
	return newLibObj
}

func (obj *configAppendResourceFlowIter) Append(items ...Flow) ConfigAppendResourceFlowIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowSlice = append(obj.flowSlice, item)
	}
	return obj
}

func (obj *configAppendResourceFlowIter) Set(index int, newObj Flow) ConfigAppendResourceFlowIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowSlice[index] = newObj
	return obj
}
func (obj *configAppendResourceFlowIter) Clear() ConfigAppendResourceFlowIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Flow{}
		obj.flowSlice = []Flow{}
	}
	return obj
}
func (obj *configAppendResourceFlowIter) clearHolderSlice() ConfigAppendResourceFlowIter {
	if len(obj.flowSlice) > 0 {
		obj.flowSlice = []Flow{}
	}
	return obj
}
func (obj *configAppendResourceFlowIter) appendHolderSlice(item Flow) ConfigAppendResourceFlowIter {
	obj.flowSlice = append(obj.flowSlice, item)
	return obj
}

func (obj *configAppendResource) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface ConfigAppendResource")
	}

	if len(obj.obj.Flows) != 0 {

		if set_default {
			obj.Flows().clearHolderSlice()
			for _, item := range obj.obj.Flows {
				obj.Flows().appendHolderSlice(&flow{obj: item})
			}
		}
		for _, item := range obj.Flows().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *configAppendResource) setDefault() {
	var choices_set int = 0
	var choice ConfigAppendResourceChoiceEnum

	if len(obj.obj.Flows) > 0 {
		choices_set += 1
		choice = ConfigAppendResourceChoice.FLOWS
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ConfigAppendResource")
			}
		} else {
			intVal := otg.ConfigAppendResource_Choice_Enum_value[string(choice)]
			enumValue := otg.ConfigAppendResource_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

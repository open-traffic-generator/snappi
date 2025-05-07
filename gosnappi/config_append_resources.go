package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ConfigAppendResources *****
type configAppendResources struct {
	validation
	obj          *otg.ConfigAppendResources
	marshaller   marshalConfigAppendResources
	unMarshaller unMarshalConfigAppendResources
	flowsHolder  ConfigAppendResourcesFlowIter
}

func NewConfigAppendResources() ConfigAppendResources {
	obj := configAppendResources{obj: &otg.ConfigAppendResources{}}
	obj.setDefault()
	return &obj
}

func (obj *configAppendResources) msg() *otg.ConfigAppendResources {
	return obj.obj
}

func (obj *configAppendResources) setMsg(msg *otg.ConfigAppendResources) ConfigAppendResources {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalconfigAppendResources struct {
	obj *configAppendResources
}

type marshalConfigAppendResources interface {
	// ToProto marshals ConfigAppendResources to protobuf object *otg.ConfigAppendResources
	ToProto() (*otg.ConfigAppendResources, error)
	// ToPbText marshals ConfigAppendResources to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ConfigAppendResources to YAML text
	ToYaml() (string, error)
	// ToJson marshals ConfigAppendResources to JSON text
	ToJson() (string, error)
}

type unMarshalconfigAppendResources struct {
	obj *configAppendResources
}

type unMarshalConfigAppendResources interface {
	// FromProto unmarshals ConfigAppendResources from protobuf object *otg.ConfigAppendResources
	FromProto(msg *otg.ConfigAppendResources) (ConfigAppendResources, error)
	// FromPbText unmarshals ConfigAppendResources from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ConfigAppendResources from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ConfigAppendResources from JSON text
	FromJson(value string) error
}

func (obj *configAppendResources) Marshal() marshalConfigAppendResources {
	if obj.marshaller == nil {
		obj.marshaller = &marshalconfigAppendResources{obj: obj}
	}
	return obj.marshaller
}

func (obj *configAppendResources) Unmarshal() unMarshalConfigAppendResources {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalconfigAppendResources{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalconfigAppendResources) ToProto() (*otg.ConfigAppendResources, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalconfigAppendResources) FromProto(msg *otg.ConfigAppendResources) (ConfigAppendResources, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalconfigAppendResources) ToPbText() (string, error) {
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

func (m *unMarshalconfigAppendResources) FromPbText(value string) error {
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

func (m *marshalconfigAppendResources) ToYaml() (string, error) {
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

func (m *unMarshalconfigAppendResources) FromYaml(value string) error {
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

func (m *marshalconfigAppendResources) ToJson() (string, error) {
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

func (m *unMarshalconfigAppendResources) FromJson(value string) error {
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

func (obj *configAppendResources) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *configAppendResources) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *configAppendResources) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *configAppendResources) Clone() (ConfigAppendResources, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewConfigAppendResources()
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

func (obj *configAppendResources) setNil() {
	obj.flowsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ConfigAppendResources is a container for an existing resource to be appended.
type ConfigAppendResources interface {
	Validation
	// msg marshals ConfigAppendResources to protobuf object *otg.ConfigAppendResources
	// and doesn't set defaults
	msg() *otg.ConfigAppendResources
	// setMsg unmarshals ConfigAppendResources from protobuf object *otg.ConfigAppendResources
	// and doesn't set defaults
	setMsg(*otg.ConfigAppendResources) ConfigAppendResources
	// provides marshal interface
	Marshal() marshalConfigAppendResources
	// provides unmarshal interface
	Unmarshal() unMarshalConfigAppendResources
	// validate validates ConfigAppendResources
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ConfigAppendResources, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ConfigAppendResourcesChoiceEnum, set in ConfigAppendResources
	Choice() ConfigAppendResourcesChoiceEnum
	// setChoice assigns ConfigAppendResourcesChoiceEnum provided by user to ConfigAppendResources
	setChoice(value ConfigAppendResourcesChoiceEnum) ConfigAppendResources
	// Flows returns ConfigAppendResourcesFlowIterIter, set in ConfigAppendResources
	Flows() ConfigAppendResourcesFlowIter
	setNil()
}

type ConfigAppendResourcesChoiceEnum string

// Enum of Choice on ConfigAppendResources
var ConfigAppendResourcesChoice = struct {
	FLOWS ConfigAppendResourcesChoiceEnum
}{
	FLOWS: ConfigAppendResourcesChoiceEnum("flows"),
}

func (obj *configAppendResources) Choice() ConfigAppendResourcesChoiceEnum {
	return ConfigAppendResourcesChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *configAppendResources) setChoice(value ConfigAppendResourcesChoiceEnum) ConfigAppendResources {
	intValue, ok := otg.ConfigAppendResources_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ConfigAppendResourcesChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ConfigAppendResources_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Flows = nil
	obj.flowsHolder = nil

	if value == ConfigAppendResourcesChoice.FLOWS {
		obj.obj.Flows = []*otg.Flow{}
	}

	return obj
}

// The flows that will be added to existing configuration on the traffic generator.
// The flow name must not already be used in existing configuration.
// Flows returns a []Flow
func (obj *configAppendResources) Flows() ConfigAppendResourcesFlowIter {
	if len(obj.obj.Flows) == 0 {
		obj.setChoice(ConfigAppendResourcesChoice.FLOWS)
	}
	if obj.flowsHolder == nil {
		obj.flowsHolder = newConfigAppendResourcesFlowIter(&obj.obj.Flows).setMsg(obj)
	}
	return obj.flowsHolder
}

type configAppendResourcesFlowIter struct {
	obj       *configAppendResources
	flowSlice []Flow
	fieldPtr  *[]*otg.Flow
}

func newConfigAppendResourcesFlowIter(ptr *[]*otg.Flow) ConfigAppendResourcesFlowIter {
	return &configAppendResourcesFlowIter{fieldPtr: ptr}
}

type ConfigAppendResourcesFlowIter interface {
	setMsg(*configAppendResources) ConfigAppendResourcesFlowIter
	Items() []Flow
	Add() Flow
	Append(items ...Flow) ConfigAppendResourcesFlowIter
	Set(index int, newObj Flow) ConfigAppendResourcesFlowIter
	Clear() ConfigAppendResourcesFlowIter
	clearHolderSlice() ConfigAppendResourcesFlowIter
	appendHolderSlice(item Flow) ConfigAppendResourcesFlowIter
}

func (obj *configAppendResourcesFlowIter) setMsg(msg *configAppendResources) ConfigAppendResourcesFlowIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flow{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *configAppendResourcesFlowIter) Items() []Flow {
	return obj.flowSlice
}

func (obj *configAppendResourcesFlowIter) Add() Flow {
	newObj := &otg.Flow{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flow{obj: newObj}
	newLibObj.setDefault()
	obj.flowSlice = append(obj.flowSlice, newLibObj)
	return newLibObj
}

func (obj *configAppendResourcesFlowIter) Append(items ...Flow) ConfigAppendResourcesFlowIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowSlice = append(obj.flowSlice, item)
	}
	return obj
}

func (obj *configAppendResourcesFlowIter) Set(index int, newObj Flow) ConfigAppendResourcesFlowIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowSlice[index] = newObj
	return obj
}
func (obj *configAppendResourcesFlowIter) Clear() ConfigAppendResourcesFlowIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Flow{}
		obj.flowSlice = []Flow{}
	}
	return obj
}
func (obj *configAppendResourcesFlowIter) clearHolderSlice() ConfigAppendResourcesFlowIter {
	if len(obj.flowSlice) > 0 {
		obj.flowSlice = []Flow{}
	}
	return obj
}
func (obj *configAppendResourcesFlowIter) appendHolderSlice(item Flow) ConfigAppendResourcesFlowIter {
	obj.flowSlice = append(obj.flowSlice, item)
	return obj
}

func (obj *configAppendResources) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface ConfigAppendResources")
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

func (obj *configAppendResources) setDefault() {
	var choices_set int = 0
	var choice ConfigAppendResourcesChoiceEnum

	if len(obj.obj.Flows) > 0 {
		choices_set += 1
		choice = ConfigAppendResourcesChoice.FLOWS
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ConfigAppendResources")
			}
		} else {
			intVal := otg.ConfigAppendResources_Choice_Enum_value[string(choice)]
			enumValue := otg.ConfigAppendResources_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

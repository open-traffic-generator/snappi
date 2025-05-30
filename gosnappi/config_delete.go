package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ConfigDelete *****
type configDelete struct {
	validation
	obj                    *otg.ConfigDelete
	marshaller             marshalConfigDelete
	unMarshaller           unMarshalConfigDelete
	configDeleteListHolder ConfigDeleteConfigDeleteResourcesIter
}

func NewConfigDelete() ConfigDelete {
	obj := configDelete{obj: &otg.ConfigDelete{}}
	obj.setDefault()
	return &obj
}

func (obj *configDelete) msg() *otg.ConfigDelete {
	return obj.obj
}

func (obj *configDelete) setMsg(msg *otg.ConfigDelete) ConfigDelete {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalconfigDelete struct {
	obj *configDelete
}

type marshalConfigDelete interface {
	// ToProto marshals ConfigDelete to protobuf object *otg.ConfigDelete
	ToProto() (*otg.ConfigDelete, error)
	// ToPbText marshals ConfigDelete to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ConfigDelete to YAML text
	ToYaml() (string, error)
	// ToJson marshals ConfigDelete to JSON text
	ToJson() (string, error)
}

type unMarshalconfigDelete struct {
	obj *configDelete
}

type unMarshalConfigDelete interface {
	// FromProto unmarshals ConfigDelete from protobuf object *otg.ConfigDelete
	FromProto(msg *otg.ConfigDelete) (ConfigDelete, error)
	// FromPbText unmarshals ConfigDelete from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ConfigDelete from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ConfigDelete from JSON text
	FromJson(value string) error
}

func (obj *configDelete) Marshal() marshalConfigDelete {
	if obj.marshaller == nil {
		obj.marshaller = &marshalconfigDelete{obj: obj}
	}
	return obj.marshaller
}

func (obj *configDelete) Unmarshal() unMarshalConfigDelete {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalconfigDelete{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalconfigDelete) ToProto() (*otg.ConfigDelete, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalconfigDelete) FromProto(msg *otg.ConfigDelete) (ConfigDelete, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalconfigDelete) ToPbText() (string, error) {
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

func (m *unMarshalconfigDelete) FromPbText(value string) error {
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

func (m *marshalconfigDelete) ToYaml() (string, error) {
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

func (m *unMarshalconfigDelete) FromYaml(value string) error {
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

func (m *marshalconfigDelete) ToJson() (string, error) {
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

func (m *unMarshalconfigDelete) FromJson(value string) error {
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

func (obj *configDelete) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *configDelete) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *configDelete) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *configDelete) Clone() (ConfigDelete, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewConfigDelete()
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

func (obj *configDelete) setNil() {
	obj.configDeleteListHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ConfigDelete is a container for all configuration resources of various types to be deleted.
type ConfigDelete interface {
	Validation
	// msg marshals ConfigDelete to protobuf object *otg.ConfigDelete
	// and doesn't set defaults
	msg() *otg.ConfigDelete
	// setMsg unmarshals ConfigDelete from protobuf object *otg.ConfigDelete
	// and doesn't set defaults
	setMsg(*otg.ConfigDelete) ConfigDelete
	// provides marshal interface
	Marshal() marshalConfigDelete
	// provides unmarshal interface
	Unmarshal() unMarshalConfigDelete
	// validate validates ConfigDelete
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ConfigDelete, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ConfigDeleteList returns ConfigDeleteConfigDeleteResourcesIterIter, set in ConfigDelete
	ConfigDeleteList() ConfigDeleteConfigDeleteResourcesIter
	setNil()
}

// description is TBD
// ConfigDeleteList returns a []ConfigDeleteResources
func (obj *configDelete) ConfigDeleteList() ConfigDeleteConfigDeleteResourcesIter {
	if len(obj.obj.ConfigDeleteList) == 0 {
		obj.obj.ConfigDeleteList = []*otg.ConfigDeleteResources{}
	}
	if obj.configDeleteListHolder == nil {
		obj.configDeleteListHolder = newConfigDeleteConfigDeleteResourcesIter(&obj.obj.ConfigDeleteList).setMsg(obj)
	}
	return obj.configDeleteListHolder
}

type configDeleteConfigDeleteResourcesIter struct {
	obj                        *configDelete
	configDeleteResourcesSlice []ConfigDeleteResources
	fieldPtr                   *[]*otg.ConfigDeleteResources
}

func newConfigDeleteConfigDeleteResourcesIter(ptr *[]*otg.ConfigDeleteResources) ConfigDeleteConfigDeleteResourcesIter {
	return &configDeleteConfigDeleteResourcesIter{fieldPtr: ptr}
}

type ConfigDeleteConfigDeleteResourcesIter interface {
	setMsg(*configDelete) ConfigDeleteConfigDeleteResourcesIter
	Items() []ConfigDeleteResources
	Add() ConfigDeleteResources
	Append(items ...ConfigDeleteResources) ConfigDeleteConfigDeleteResourcesIter
	Set(index int, newObj ConfigDeleteResources) ConfigDeleteConfigDeleteResourcesIter
	Clear() ConfigDeleteConfigDeleteResourcesIter
	clearHolderSlice() ConfigDeleteConfigDeleteResourcesIter
	appendHolderSlice(item ConfigDeleteResources) ConfigDeleteConfigDeleteResourcesIter
}

func (obj *configDeleteConfigDeleteResourcesIter) setMsg(msg *configDelete) ConfigDeleteConfigDeleteResourcesIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&configDeleteResources{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *configDeleteConfigDeleteResourcesIter) Items() []ConfigDeleteResources {
	return obj.configDeleteResourcesSlice
}

func (obj *configDeleteConfigDeleteResourcesIter) Add() ConfigDeleteResources {
	newObj := &otg.ConfigDeleteResources{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &configDeleteResources{obj: newObj}
	newLibObj.setDefault()
	obj.configDeleteResourcesSlice = append(obj.configDeleteResourcesSlice, newLibObj)
	return newLibObj
}

func (obj *configDeleteConfigDeleteResourcesIter) Append(items ...ConfigDeleteResources) ConfigDeleteConfigDeleteResourcesIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.configDeleteResourcesSlice = append(obj.configDeleteResourcesSlice, item)
	}
	return obj
}

func (obj *configDeleteConfigDeleteResourcesIter) Set(index int, newObj ConfigDeleteResources) ConfigDeleteConfigDeleteResourcesIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.configDeleteResourcesSlice[index] = newObj
	return obj
}
func (obj *configDeleteConfigDeleteResourcesIter) Clear() ConfigDeleteConfigDeleteResourcesIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ConfigDeleteResources{}
		obj.configDeleteResourcesSlice = []ConfigDeleteResources{}
	}
	return obj
}
func (obj *configDeleteConfigDeleteResourcesIter) clearHolderSlice() ConfigDeleteConfigDeleteResourcesIter {
	if len(obj.configDeleteResourcesSlice) > 0 {
		obj.configDeleteResourcesSlice = []ConfigDeleteResources{}
	}
	return obj
}
func (obj *configDeleteConfigDeleteResourcesIter) appendHolderSlice(item ConfigDeleteResources) ConfigDeleteConfigDeleteResourcesIter {
	obj.configDeleteResourcesSlice = append(obj.configDeleteResourcesSlice, item)
	return obj
}

func (obj *configDelete) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.ConfigDeleteList) != 0 {

		if set_default {
			obj.ConfigDeleteList().clearHolderSlice()
			for _, item := range obj.obj.ConfigDeleteList {
				obj.ConfigDeleteList().appendHolderSlice(&configDeleteResources{obj: item})
			}
		}
		for _, item := range obj.ConfigDeleteList().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *configDelete) setDefault() {

}

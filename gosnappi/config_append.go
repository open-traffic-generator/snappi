package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ConfigAppend *****
type configAppend struct {
	validation
	obj                    *otg.ConfigAppend
	marshaller             marshalConfigAppend
	unMarshaller           unMarshalConfigAppend
	configAppendListHolder ConfigAppendConfigAppendResourcesIter
}

func NewConfigAppend() ConfigAppend {
	obj := configAppend{obj: &otg.ConfigAppend{}}
	obj.setDefault()
	return &obj
}

func (obj *configAppend) msg() *otg.ConfigAppend {
	return obj.obj
}

func (obj *configAppend) setMsg(msg *otg.ConfigAppend) ConfigAppend {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalconfigAppend struct {
	obj *configAppend
}

type marshalConfigAppend interface {
	// ToProto marshals ConfigAppend to protobuf object *otg.ConfigAppend
	ToProto() (*otg.ConfigAppend, error)
	// ToPbText marshals ConfigAppend to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ConfigAppend to YAML text
	ToYaml() (string, error)
	// ToJson marshals ConfigAppend to JSON text
	ToJson() (string, error)
}

type unMarshalconfigAppend struct {
	obj *configAppend
}

type unMarshalConfigAppend interface {
	// FromProto unmarshals ConfigAppend from protobuf object *otg.ConfigAppend
	FromProto(msg *otg.ConfigAppend) (ConfigAppend, error)
	// FromPbText unmarshals ConfigAppend from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ConfigAppend from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ConfigAppend from JSON text
	FromJson(value string) error
}

func (obj *configAppend) Marshal() marshalConfigAppend {
	if obj.marshaller == nil {
		obj.marshaller = &marshalconfigAppend{obj: obj}
	}
	return obj.marshaller
}

func (obj *configAppend) Unmarshal() unMarshalConfigAppend {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalconfigAppend{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalconfigAppend) ToProto() (*otg.ConfigAppend, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalconfigAppend) FromProto(msg *otg.ConfigAppend) (ConfigAppend, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalconfigAppend) ToPbText() (string, error) {
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

func (m *unMarshalconfigAppend) FromPbText(value string) error {
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

func (m *marshalconfigAppend) ToYaml() (string, error) {
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

func (m *unMarshalconfigAppend) FromYaml(value string) error {
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

func (m *marshalconfigAppend) ToJson() (string, error) {
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

func (m *unMarshalconfigAppend) FromJson(value string) error {
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

func (obj *configAppend) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *configAppend) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *configAppend) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *configAppend) Clone() (ConfigAppend, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewConfigAppend()
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

func (obj *configAppend) setNil() {
	obj.configAppendListHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ConfigAppend is a container for all configuration resources of various types to be appended.
type ConfigAppend interface {
	Validation
	// msg marshals ConfigAppend to protobuf object *otg.ConfigAppend
	// and doesn't set defaults
	msg() *otg.ConfigAppend
	// setMsg unmarshals ConfigAppend from protobuf object *otg.ConfigAppend
	// and doesn't set defaults
	setMsg(*otg.ConfigAppend) ConfigAppend
	// provides marshal interface
	Marshal() marshalConfigAppend
	// provides unmarshal interface
	Unmarshal() unMarshalConfigAppend
	// validate validates ConfigAppend
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ConfigAppend, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ConfigAppendList returns ConfigAppendConfigAppendResourcesIterIter, set in ConfigAppend
	ConfigAppendList() ConfigAppendConfigAppendResourcesIter
	setNil()
}

// description is TBD
// ConfigAppendList returns a []ConfigAppendResources
func (obj *configAppend) ConfigAppendList() ConfigAppendConfigAppendResourcesIter {
	if len(obj.obj.ConfigAppendList) == 0 {
		obj.obj.ConfigAppendList = []*otg.ConfigAppendResources{}
	}
	if obj.configAppendListHolder == nil {
		obj.configAppendListHolder = newConfigAppendConfigAppendResourcesIter(&obj.obj.ConfigAppendList).setMsg(obj)
	}
	return obj.configAppendListHolder
}

type configAppendConfigAppendResourcesIter struct {
	obj                        *configAppend
	configAppendResourcesSlice []ConfigAppendResources
	fieldPtr                   *[]*otg.ConfigAppendResources
}

func newConfigAppendConfigAppendResourcesIter(ptr *[]*otg.ConfigAppendResources) ConfigAppendConfigAppendResourcesIter {
	return &configAppendConfigAppendResourcesIter{fieldPtr: ptr}
}

type ConfigAppendConfigAppendResourcesIter interface {
	setMsg(*configAppend) ConfigAppendConfigAppendResourcesIter
	Items() []ConfigAppendResources
	Add() ConfigAppendResources
	Append(items ...ConfigAppendResources) ConfigAppendConfigAppendResourcesIter
	Set(index int, newObj ConfigAppendResources) ConfigAppendConfigAppendResourcesIter
	Clear() ConfigAppendConfigAppendResourcesIter
	clearHolderSlice() ConfigAppendConfigAppendResourcesIter
	appendHolderSlice(item ConfigAppendResources) ConfigAppendConfigAppendResourcesIter
}

func (obj *configAppendConfigAppendResourcesIter) setMsg(msg *configAppend) ConfigAppendConfigAppendResourcesIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&configAppendResources{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *configAppendConfigAppendResourcesIter) Items() []ConfigAppendResources {
	return obj.configAppendResourcesSlice
}

func (obj *configAppendConfigAppendResourcesIter) Add() ConfigAppendResources {
	newObj := &otg.ConfigAppendResources{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &configAppendResources{obj: newObj}
	newLibObj.setDefault()
	obj.configAppendResourcesSlice = append(obj.configAppendResourcesSlice, newLibObj)
	return newLibObj
}

func (obj *configAppendConfigAppendResourcesIter) Append(items ...ConfigAppendResources) ConfigAppendConfigAppendResourcesIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.configAppendResourcesSlice = append(obj.configAppendResourcesSlice, item)
	}
	return obj
}

func (obj *configAppendConfigAppendResourcesIter) Set(index int, newObj ConfigAppendResources) ConfigAppendConfigAppendResourcesIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.configAppendResourcesSlice[index] = newObj
	return obj
}
func (obj *configAppendConfigAppendResourcesIter) Clear() ConfigAppendConfigAppendResourcesIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ConfigAppendResources{}
		obj.configAppendResourcesSlice = []ConfigAppendResources{}
	}
	return obj
}
func (obj *configAppendConfigAppendResourcesIter) clearHolderSlice() ConfigAppendConfigAppendResourcesIter {
	if len(obj.configAppendResourcesSlice) > 0 {
		obj.configAppendResourcesSlice = []ConfigAppendResources{}
	}
	return obj
}
func (obj *configAppendConfigAppendResourcesIter) appendHolderSlice(item ConfigAppendResources) ConfigAppendConfigAppendResourcesIter {
	obj.configAppendResourcesSlice = append(obj.configAppendResourcesSlice, item)
	return obj
}

func (obj *configAppend) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.ConfigAppendList) != 0 {

		if set_default {
			obj.ConfigAppendList().clearHolderSlice()
			for _, item := range obj.obj.ConfigAppendList {
				obj.ConfigAppendList().appendHolderSlice(&configAppendResources{obj: item})
			}
		}
		for _, item := range obj.ConfigAppendList().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *configAppend) setDefault() {

}

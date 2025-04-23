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
	configAppendListHolder ConfigAppendConfigAppendResourceIter
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
	// ConfigAppendList returns ConfigAppendConfigAppendResourceIterIter, set in ConfigAppend
	ConfigAppendList() ConfigAppendConfigAppendResourceIter
	setNil()
}

// description is TBD
// ConfigAppendList returns a []ConfigAppendResource
func (obj *configAppend) ConfigAppendList() ConfigAppendConfigAppendResourceIter {
	if len(obj.obj.ConfigAppendList) == 0 {
		obj.obj.ConfigAppendList = []*otg.ConfigAppendResource{}
	}
	if obj.configAppendListHolder == nil {
		obj.configAppendListHolder = newConfigAppendConfigAppendResourceIter(&obj.obj.ConfigAppendList).setMsg(obj)
	}
	return obj.configAppendListHolder
}

type configAppendConfigAppendResourceIter struct {
	obj                       *configAppend
	configAppendResourceSlice []ConfigAppendResource
	fieldPtr                  *[]*otg.ConfigAppendResource
}

func newConfigAppendConfigAppendResourceIter(ptr *[]*otg.ConfigAppendResource) ConfigAppendConfigAppendResourceIter {
	return &configAppendConfigAppendResourceIter{fieldPtr: ptr}
}

type ConfigAppendConfigAppendResourceIter interface {
	setMsg(*configAppend) ConfigAppendConfigAppendResourceIter
	Items() []ConfigAppendResource
	Add() ConfigAppendResource
	Append(items ...ConfigAppendResource) ConfigAppendConfigAppendResourceIter
	Set(index int, newObj ConfigAppendResource) ConfigAppendConfigAppendResourceIter
	Clear() ConfigAppendConfigAppendResourceIter
	clearHolderSlice() ConfigAppendConfigAppendResourceIter
	appendHolderSlice(item ConfigAppendResource) ConfigAppendConfigAppendResourceIter
}

func (obj *configAppendConfigAppendResourceIter) setMsg(msg *configAppend) ConfigAppendConfigAppendResourceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&configAppendResource{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *configAppendConfigAppendResourceIter) Items() []ConfigAppendResource {
	return obj.configAppendResourceSlice
}

func (obj *configAppendConfigAppendResourceIter) Add() ConfigAppendResource {
	newObj := &otg.ConfigAppendResource{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &configAppendResource{obj: newObj}
	newLibObj.setDefault()
	obj.configAppendResourceSlice = append(obj.configAppendResourceSlice, newLibObj)
	return newLibObj
}

func (obj *configAppendConfigAppendResourceIter) Append(items ...ConfigAppendResource) ConfigAppendConfigAppendResourceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.configAppendResourceSlice = append(obj.configAppendResourceSlice, item)
	}
	return obj
}

func (obj *configAppendConfigAppendResourceIter) Set(index int, newObj ConfigAppendResource) ConfigAppendConfigAppendResourceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.configAppendResourceSlice[index] = newObj
	return obj
}
func (obj *configAppendConfigAppendResourceIter) Clear() ConfigAppendConfigAppendResourceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ConfigAppendResource{}
		obj.configAppendResourceSlice = []ConfigAppendResource{}
	}
	return obj
}
func (obj *configAppendConfigAppendResourceIter) clearHolderSlice() ConfigAppendConfigAppendResourceIter {
	if len(obj.configAppendResourceSlice) > 0 {
		obj.configAppendResourceSlice = []ConfigAppendResource{}
	}
	return obj
}
func (obj *configAppendConfigAppendResourceIter) appendHolderSlice(item ConfigAppendResource) ConfigAppendConfigAppendResourceIter {
	obj.configAppendResourceSlice = append(obj.configAppendResourceSlice, item)
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
				obj.ConfigAppendList().appendHolderSlice(&configAppendResource{obj: item})
			}
		}
		for _, item := range obj.ConfigAppendList().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *configAppend) setDefault() {

}

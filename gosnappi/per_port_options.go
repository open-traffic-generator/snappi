package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PerPortOptions *****
type perPortOptions struct {
	validation
	obj                *otg.PerPortOptions
	marshaller         marshalPerPortOptions
	unMarshaller       unMarshalPerPortOptions
	portSettingsHolder PerPortOptionsOptionsPortSettingsIter
}

func NewPerPortOptions() PerPortOptions {
	obj := perPortOptions{obj: &otg.PerPortOptions{}}
	obj.setDefault()
	return &obj
}

func (obj *perPortOptions) msg() *otg.PerPortOptions {
	return obj.obj
}

func (obj *perPortOptions) setMsg(msg *otg.PerPortOptions) PerPortOptions {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalperPortOptions struct {
	obj *perPortOptions
}

type marshalPerPortOptions interface {
	// ToProto marshals PerPortOptions to protobuf object *otg.PerPortOptions
	ToProto() (*otg.PerPortOptions, error)
	// ToPbText marshals PerPortOptions to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PerPortOptions to YAML text
	ToYaml() (string, error)
	// ToJson marshals PerPortOptions to JSON text
	ToJson() (string, error)
}

type unMarshalperPortOptions struct {
	obj *perPortOptions
}

type unMarshalPerPortOptions interface {
	// FromProto unmarshals PerPortOptions from protobuf object *otg.PerPortOptions
	FromProto(msg *otg.PerPortOptions) (PerPortOptions, error)
	// FromPbText unmarshals PerPortOptions from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PerPortOptions from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PerPortOptions from JSON text
	FromJson(value string) error
}

func (obj *perPortOptions) Marshal() marshalPerPortOptions {
	if obj.marshaller == nil {
		obj.marshaller = &marshalperPortOptions{obj: obj}
	}
	return obj.marshaller
}

func (obj *perPortOptions) Unmarshal() unMarshalPerPortOptions {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalperPortOptions{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalperPortOptions) ToProto() (*otg.PerPortOptions, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalperPortOptions) FromProto(msg *otg.PerPortOptions) (PerPortOptions, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalperPortOptions) ToPbText() (string, error) {
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

func (m *unMarshalperPortOptions) FromPbText(value string) error {
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

func (m *marshalperPortOptions) ToYaml() (string, error) {
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

func (m *unMarshalperPortOptions) FromYaml(value string) error {
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

func (m *marshalperPortOptions) ToJson() (string, error) {
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

func (m *unMarshalperPortOptions) FromJson(value string) error {
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

func (obj *perPortOptions) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *perPortOptions) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *perPortOptions) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *perPortOptions) Clone() (PerPortOptions, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPerPortOptions()
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

func (obj *perPortOptions) setNil() {
	obj.portSettingsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PerPortOptions is ****Add proper description here****
type PerPortOptions interface {
	Validation
	// msg marshals PerPortOptions to protobuf object *otg.PerPortOptions
	// and doesn't set defaults
	msg() *otg.PerPortOptions
	// setMsg unmarshals PerPortOptions from protobuf object *otg.PerPortOptions
	// and doesn't set defaults
	setMsg(*otg.PerPortOptions) PerPortOptions
	// provides marshal interface
	Marshal() marshalPerPortOptions
	// provides unmarshal interface
	Unmarshal() unMarshalPerPortOptions
	// validate validates PerPortOptions
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PerPortOptions, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PortSettings returns PerPortOptionsOptionsPortSettingsIterIter, set in PerPortOptions
	PortSettings() PerPortOptionsOptionsPortSettingsIter
	setNil()
}

// This contains all the per port settings for each protocol.
// PortSettings returns a []OptionsPortSettings
func (obj *perPortOptions) PortSettings() PerPortOptionsOptionsPortSettingsIter {
	if len(obj.obj.PortSettings) == 0 {
		obj.obj.PortSettings = []*otg.OptionsPortSettings{}
	}
	if obj.portSettingsHolder == nil {
		obj.portSettingsHolder = newPerPortOptionsOptionsPortSettingsIter(&obj.obj.PortSettings).setMsg(obj)
	}
	return obj.portSettingsHolder
}

type perPortOptionsOptionsPortSettingsIter struct {
	obj                      *perPortOptions
	optionsPortSettingsSlice []OptionsPortSettings
	fieldPtr                 *[]*otg.OptionsPortSettings
}

func newPerPortOptionsOptionsPortSettingsIter(ptr *[]*otg.OptionsPortSettings) PerPortOptionsOptionsPortSettingsIter {
	return &perPortOptionsOptionsPortSettingsIter{fieldPtr: ptr}
}

type PerPortOptionsOptionsPortSettingsIter interface {
	setMsg(*perPortOptions) PerPortOptionsOptionsPortSettingsIter
	Items() []OptionsPortSettings
	Add() OptionsPortSettings
	Append(items ...OptionsPortSettings) PerPortOptionsOptionsPortSettingsIter
	Set(index int, newObj OptionsPortSettings) PerPortOptionsOptionsPortSettingsIter
	Clear() PerPortOptionsOptionsPortSettingsIter
	clearHolderSlice() PerPortOptionsOptionsPortSettingsIter
	appendHolderSlice(item OptionsPortSettings) PerPortOptionsOptionsPortSettingsIter
}

func (obj *perPortOptionsOptionsPortSettingsIter) setMsg(msg *perPortOptions) PerPortOptionsOptionsPortSettingsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&optionsPortSettings{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *perPortOptionsOptionsPortSettingsIter) Items() []OptionsPortSettings {
	return obj.optionsPortSettingsSlice
}

func (obj *perPortOptionsOptionsPortSettingsIter) Add() OptionsPortSettings {
	newObj := &otg.OptionsPortSettings{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &optionsPortSettings{obj: newObj}
	newLibObj.setDefault()
	obj.optionsPortSettingsSlice = append(obj.optionsPortSettingsSlice, newLibObj)
	return newLibObj
}

func (obj *perPortOptionsOptionsPortSettingsIter) Append(items ...OptionsPortSettings) PerPortOptionsOptionsPortSettingsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.optionsPortSettingsSlice = append(obj.optionsPortSettingsSlice, item)
	}
	return obj
}

func (obj *perPortOptionsOptionsPortSettingsIter) Set(index int, newObj OptionsPortSettings) PerPortOptionsOptionsPortSettingsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.optionsPortSettingsSlice[index] = newObj
	return obj
}
func (obj *perPortOptionsOptionsPortSettingsIter) Clear() PerPortOptionsOptionsPortSettingsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.OptionsPortSettings{}
		obj.optionsPortSettingsSlice = []OptionsPortSettings{}
	}
	return obj
}
func (obj *perPortOptionsOptionsPortSettingsIter) clearHolderSlice() PerPortOptionsOptionsPortSettingsIter {
	if len(obj.optionsPortSettingsSlice) > 0 {
		obj.optionsPortSettingsSlice = []OptionsPortSettings{}
	}
	return obj
}
func (obj *perPortOptionsOptionsPortSettingsIter) appendHolderSlice(item OptionsPortSettings) PerPortOptionsOptionsPortSettingsIter {
	obj.optionsPortSettingsSlice = append(obj.optionsPortSettingsSlice, item)
	return obj
}

func (obj *perPortOptions) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.PortSettings) != 0 {

		if set_default {
			obj.PortSettings().clearHolderSlice()
			for _, item := range obj.obj.PortSettings {
				obj.PortSettings().appendHolderSlice(&optionsPortSettings{obj: item})
			}
		}
		for _, item := range obj.PortSettings().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *perPortOptions) setDefault() {

}

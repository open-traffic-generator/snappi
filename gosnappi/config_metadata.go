package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ConfigMetadata *****
type configMetadata struct {
	validation
	obj                 *otg.ConfigMetadata
	marshaller          marshalConfigMetadata
	unMarshaller        unMarshalConfigMetadata
	tagValuePairsHolder ConfigMetadataConfigTagValuePairIter
}

func NewConfigMetadata() ConfigMetadata {
	obj := configMetadata{obj: &otg.ConfigMetadata{}}
	obj.setDefault()
	return &obj
}

func (obj *configMetadata) msg() *otg.ConfigMetadata {
	return obj.obj
}

func (obj *configMetadata) setMsg(msg *otg.ConfigMetadata) ConfigMetadata {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalconfigMetadata struct {
	obj *configMetadata
}

type marshalConfigMetadata interface {
	// ToProto marshals ConfigMetadata to protobuf object *otg.ConfigMetadata
	ToProto() (*otg.ConfigMetadata, error)
	// ToPbText marshals ConfigMetadata to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ConfigMetadata to YAML text
	ToYaml() (string, error)
	// ToJson marshals ConfigMetadata to JSON text
	ToJson() (string, error)
}

type unMarshalconfigMetadata struct {
	obj *configMetadata
}

type unMarshalConfigMetadata interface {
	// FromProto unmarshals ConfigMetadata from protobuf object *otg.ConfigMetadata
	FromProto(msg *otg.ConfigMetadata) (ConfigMetadata, error)
	// FromPbText unmarshals ConfigMetadata from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ConfigMetadata from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ConfigMetadata from JSON text
	FromJson(value string) error
}

func (obj *configMetadata) Marshal() marshalConfigMetadata {
	if obj.marshaller == nil {
		obj.marshaller = &marshalconfigMetadata{obj: obj}
	}
	return obj.marshaller
}

func (obj *configMetadata) Unmarshal() unMarshalConfigMetadata {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalconfigMetadata{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalconfigMetadata) ToProto() (*otg.ConfigMetadata, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalconfigMetadata) FromProto(msg *otg.ConfigMetadata) (ConfigMetadata, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalconfigMetadata) ToPbText() (string, error) {
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

func (m *unMarshalconfigMetadata) FromPbText(value string) error {
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

func (m *marshalconfigMetadata) ToYaml() (string, error) {
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

func (m *unMarshalconfigMetadata) FromYaml(value string) error {
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

func (m *marshalconfigMetadata) ToJson() (string, error) {
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

func (m *unMarshalconfigMetadata) FromJson(value string) error {
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

func (obj *configMetadata) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *configMetadata) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *configMetadata) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *configMetadata) Clone() (ConfigMetadata, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewConfigMetadata()
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

func (obj *configMetadata) setNil() {
	obj.tagValuePairsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ConfigMetadata is this optional container can be used to specify additional information about the test and specific
// configuration being applied via the set_config api.
// It can be used to correlate logs generated by the test tool to particular test being run by being logged
// as part of the information recorded in the test tool.
// e.g. In a test harness, for every test and sub-test the test name and sub-test name can be provided
// for every set_config done as part of the test . Or the timestamp of the set_config api call on the
// client can be recorded as a tag-value pair.
type ConfigMetadata interface {
	Validation
	// msg marshals ConfigMetadata to protobuf object *otg.ConfigMetadata
	// and doesn't set defaults
	msg() *otg.ConfigMetadata
	// setMsg unmarshals ConfigMetadata from protobuf object *otg.ConfigMetadata
	// and doesn't set defaults
	setMsg(*otg.ConfigMetadata) ConfigMetadata
	// provides marshal interface
	Marshal() marshalConfigMetadata
	// provides unmarshal interface
	Unmarshal() unMarshalConfigMetadata
	// validate validates ConfigMetadata
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ConfigMetadata, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Description returns string, set in ConfigMetadata.
	Description() string
	// SetDescription assigns string provided by user to ConfigMetadata
	SetDescription(value string) ConfigMetadata
	// HasDescription checks if Description has been set in ConfigMetadata
	HasDescription() bool
	// TagValuePairs returns ConfigMetadataConfigTagValuePairIterIter, set in ConfigMetadata
	TagValuePairs() ConfigMetadataConfigTagValuePairIter
	setNil()
}

// A generic field which can be used to provide a human readable string describing the test being run.
// e.g. the test or sub-test name
//
// Description returns a string
func (obj *configMetadata) Description() string {

	return *obj.obj.Description

}

// A generic field which can be used to provide a human readable string describing the test being run.
// e.g. the test or sub-test name
//
// Description returns a string
func (obj *configMetadata) HasDescription() bool {
	return obj.obj.Description != nil
}

// A generic field which can be used to provide a human readable string describing the test being run.
// e.g. the test or sub-test name
//
// SetDescription sets the string value in the ConfigMetadata object
func (obj *configMetadata) SetDescription(value string) ConfigMetadata {

	obj.obj.Description = &value
	return obj
}

// The list of generic tag value pairs which can be used to add generic metadata about the set_config or test being run.
// e.g. {"sub-test","sub-test-name"} or {}"apply-time","The time in UTC at which the set_config was applied on the client"}
//
// TagValuePairs returns a []ConfigTagValuePair
func (obj *configMetadata) TagValuePairs() ConfigMetadataConfigTagValuePairIter {
	if len(obj.obj.TagValuePairs) == 0 {
		obj.obj.TagValuePairs = []*otg.ConfigTagValuePair{}
	}
	if obj.tagValuePairsHolder == nil {
		obj.tagValuePairsHolder = newConfigMetadataConfigTagValuePairIter(&obj.obj.TagValuePairs).setMsg(obj)
	}
	return obj.tagValuePairsHolder
}

type configMetadataConfigTagValuePairIter struct {
	obj                     *configMetadata
	configTagValuePairSlice []ConfigTagValuePair
	fieldPtr                *[]*otg.ConfigTagValuePair
}

func newConfigMetadataConfigTagValuePairIter(ptr *[]*otg.ConfigTagValuePair) ConfigMetadataConfigTagValuePairIter {
	return &configMetadataConfigTagValuePairIter{fieldPtr: ptr}
}

type ConfigMetadataConfigTagValuePairIter interface {
	setMsg(*configMetadata) ConfigMetadataConfigTagValuePairIter
	Items() []ConfigTagValuePair
	Add() ConfigTagValuePair
	Append(items ...ConfigTagValuePair) ConfigMetadataConfigTagValuePairIter
	Set(index int, newObj ConfigTagValuePair) ConfigMetadataConfigTagValuePairIter
	Clear() ConfigMetadataConfigTagValuePairIter
	clearHolderSlice() ConfigMetadataConfigTagValuePairIter
	appendHolderSlice(item ConfigTagValuePair) ConfigMetadataConfigTagValuePairIter
}

func (obj *configMetadataConfigTagValuePairIter) setMsg(msg *configMetadata) ConfigMetadataConfigTagValuePairIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&configTagValuePair{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *configMetadataConfigTagValuePairIter) Items() []ConfigTagValuePair {
	return obj.configTagValuePairSlice
}

func (obj *configMetadataConfigTagValuePairIter) Add() ConfigTagValuePair {
	newObj := &otg.ConfigTagValuePair{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &configTagValuePair{obj: newObj}
	newLibObj.setDefault()
	obj.configTagValuePairSlice = append(obj.configTagValuePairSlice, newLibObj)
	return newLibObj
}

func (obj *configMetadataConfigTagValuePairIter) Append(items ...ConfigTagValuePair) ConfigMetadataConfigTagValuePairIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.configTagValuePairSlice = append(obj.configTagValuePairSlice, item)
	}
	return obj
}

func (obj *configMetadataConfigTagValuePairIter) Set(index int, newObj ConfigTagValuePair) ConfigMetadataConfigTagValuePairIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.configTagValuePairSlice[index] = newObj
	return obj
}
func (obj *configMetadataConfigTagValuePairIter) Clear() ConfigMetadataConfigTagValuePairIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ConfigTagValuePair{}
		obj.configTagValuePairSlice = []ConfigTagValuePair{}
	}
	return obj
}
func (obj *configMetadataConfigTagValuePairIter) clearHolderSlice() ConfigMetadataConfigTagValuePairIter {
	if len(obj.configTagValuePairSlice) > 0 {
		obj.configTagValuePairSlice = []ConfigTagValuePair{}
	}
	return obj
}
func (obj *configMetadataConfigTagValuePairIter) appendHolderSlice(item ConfigTagValuePair) ConfigMetadataConfigTagValuePairIter {
	obj.configTagValuePairSlice = append(obj.configTagValuePairSlice, item)
	return obj
}

func (obj *configMetadata) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.TagValuePairs) != 0 {

		if set_default {
			obj.TagValuePairs().clearHolderSlice()
			for _, item := range obj.obj.TagValuePairs {
				obj.TagValuePairs().appendHolderSlice(&configTagValuePair{obj: item})
			}
		}
		for _, item := range obj.TagValuePairs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *configMetadata) setDefault() {

}

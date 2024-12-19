package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ConfigOptions *****
type configOptions struct {
	validation
	obj                   *otg.ConfigOptions
	marshaller            marshalConfigOptions
	unMarshaller          unMarshalConfigOptions
	portOptionsHolder     PortOptions
	protocolOptionsHolder ProtocolOptions
	vendorOptionsHolder   ConfigOptionsVendorOptionsIter
}

func NewConfigOptions() ConfigOptions {
	obj := configOptions{obj: &otg.ConfigOptions{}}
	obj.setDefault()
	return &obj
}

func (obj *configOptions) msg() *otg.ConfigOptions {
	return obj.obj
}

func (obj *configOptions) setMsg(msg *otg.ConfigOptions) ConfigOptions {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalconfigOptions struct {
	obj *configOptions
}

type marshalConfigOptions interface {
	// ToProto marshals ConfigOptions to protobuf object *otg.ConfigOptions
	ToProto() (*otg.ConfigOptions, error)
	// ToPbText marshals ConfigOptions to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ConfigOptions to YAML text
	ToYaml() (string, error)
	// ToJson marshals ConfigOptions to JSON text
	ToJson() (string, error)
}

type unMarshalconfigOptions struct {
	obj *configOptions
}

type unMarshalConfigOptions interface {
	// FromProto unmarshals ConfigOptions from protobuf object *otg.ConfigOptions
	FromProto(msg *otg.ConfigOptions) (ConfigOptions, error)
	// FromPbText unmarshals ConfigOptions from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ConfigOptions from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ConfigOptions from JSON text
	FromJson(value string) error
}

func (obj *configOptions) Marshal() marshalConfigOptions {
	if obj.marshaller == nil {
		obj.marshaller = &marshalconfigOptions{obj: obj}
	}
	return obj.marshaller
}

func (obj *configOptions) Unmarshal() unMarshalConfigOptions {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalconfigOptions{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalconfigOptions) ToProto() (*otg.ConfigOptions, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalconfigOptions) FromProto(msg *otg.ConfigOptions) (ConfigOptions, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalconfigOptions) ToPbText() (string, error) {
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

func (m *unMarshalconfigOptions) FromPbText(value string) error {
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

func (m *marshalconfigOptions) ToYaml() (string, error) {
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

func (m *unMarshalconfigOptions) FromYaml(value string) error {
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

func (m *marshalconfigOptions) ToJson() (string, error) {
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

func (m *unMarshalconfigOptions) FromJson(value string) error {
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

func (obj *configOptions) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *configOptions) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *configOptions) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *configOptions) Clone() (ConfigOptions, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewConfigOptions()
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

func (obj *configOptions) setNil() {
	obj.portOptionsHolder = nil
	obj.protocolOptionsHolder = nil
	obj.vendorOptionsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ConfigOptions is global configuration options.
type ConfigOptions interface {
	Validation
	// msg marshals ConfigOptions to protobuf object *otg.ConfigOptions
	// and doesn't set defaults
	msg() *otg.ConfigOptions
	// setMsg unmarshals ConfigOptions from protobuf object *otg.ConfigOptions
	// and doesn't set defaults
	setMsg(*otg.ConfigOptions) ConfigOptions
	// provides marshal interface
	Marshal() marshalConfigOptions
	// provides unmarshal interface
	Unmarshal() unMarshalConfigOptions
	// validate validates ConfigOptions
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ConfigOptions, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PortOptions returns PortOptions, set in ConfigOptions.
	// PortOptions is common port options that apply to all configured Port objects.
	PortOptions() PortOptions
	// SetPortOptions assigns PortOptions provided by user to ConfigOptions.
	// PortOptions is common port options that apply to all configured Port objects.
	SetPortOptions(value PortOptions) ConfigOptions
	// HasPortOptions checks if PortOptions has been set in ConfigOptions
	HasPortOptions() bool
	// ProtocolOptions returns ProtocolOptions, set in ConfigOptions.
	// ProtocolOptions is common options that apply to all configured protocols and interfaces.
	ProtocolOptions() ProtocolOptions
	// SetProtocolOptions assigns ProtocolOptions provided by user to ConfigOptions.
	// ProtocolOptions is common options that apply to all configured protocols and interfaces.
	SetProtocolOptions(value ProtocolOptions) ConfigOptions
	// HasProtocolOptions checks if ProtocolOptions has been set in ConfigOptions
	HasProtocolOptions() bool
	// VendorOptions returns ConfigOptionsVendorOptionsIterIter, set in ConfigOptions
	VendorOptions() ConfigOptionsVendorOptionsIter
	setNil()
}

// description is TBD
// PortOptions returns a PortOptions
func (obj *configOptions) PortOptions() PortOptions {
	if obj.obj.PortOptions == nil {
		obj.obj.PortOptions = NewPortOptions().msg()
	}
	if obj.portOptionsHolder == nil {
		obj.portOptionsHolder = &portOptions{obj: obj.obj.PortOptions}
	}
	return obj.portOptionsHolder
}

// description is TBD
// PortOptions returns a PortOptions
func (obj *configOptions) HasPortOptions() bool {
	return obj.obj.PortOptions != nil
}

// description is TBD
// SetPortOptions sets the PortOptions value in the ConfigOptions object
func (obj *configOptions) SetPortOptions(value PortOptions) ConfigOptions {

	obj.portOptionsHolder = nil
	obj.obj.PortOptions = value.msg()

	return obj
}

// description is TBD
// ProtocolOptions returns a ProtocolOptions
func (obj *configOptions) ProtocolOptions() ProtocolOptions {
	if obj.obj.ProtocolOptions == nil {
		obj.obj.ProtocolOptions = NewProtocolOptions().msg()
	}
	if obj.protocolOptionsHolder == nil {
		obj.protocolOptionsHolder = &protocolOptions{obj: obj.obj.ProtocolOptions}
	}
	return obj.protocolOptionsHolder
}

// description is TBD
// ProtocolOptions returns a ProtocolOptions
func (obj *configOptions) HasProtocolOptions() bool {
	return obj.obj.ProtocolOptions != nil
}

// description is TBD
// SetProtocolOptions sets the ProtocolOptions value in the ConfigOptions object
func (obj *configOptions) SetProtocolOptions(value ProtocolOptions) ConfigOptions {

	obj.protocolOptionsHolder = nil
	obj.obj.ProtocolOptions = value.msg()

	return obj
}

// This is a backdoor access to underlying OTG application.
// VendorOptions returns a []VendorOptions
func (obj *configOptions) VendorOptions() ConfigOptionsVendorOptionsIter {
	if len(obj.obj.VendorOptions) == 0 {
		obj.obj.VendorOptions = []*otg.VendorOptions{}
	}
	if obj.vendorOptionsHolder == nil {
		obj.vendorOptionsHolder = newConfigOptionsVendorOptionsIter(&obj.obj.VendorOptions).setMsg(obj)
	}
	return obj.vendorOptionsHolder
}

type configOptionsVendorOptionsIter struct {
	obj                *configOptions
	vendorOptionsSlice []VendorOptions
	fieldPtr           *[]*otg.VendorOptions
}

func newConfigOptionsVendorOptionsIter(ptr *[]*otg.VendorOptions) ConfigOptionsVendorOptionsIter {
	return &configOptionsVendorOptionsIter{fieldPtr: ptr}
}

type ConfigOptionsVendorOptionsIter interface {
	setMsg(*configOptions) ConfigOptionsVendorOptionsIter
	Items() []VendorOptions
	Add() VendorOptions
	Append(items ...VendorOptions) ConfigOptionsVendorOptionsIter
	Set(index int, newObj VendorOptions) ConfigOptionsVendorOptionsIter
	Clear() ConfigOptionsVendorOptionsIter
	clearHolderSlice() ConfigOptionsVendorOptionsIter
	appendHolderSlice(item VendorOptions) ConfigOptionsVendorOptionsIter
}

func (obj *configOptionsVendorOptionsIter) setMsg(msg *configOptions) ConfigOptionsVendorOptionsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&vendorOptions{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *configOptionsVendorOptionsIter) Items() []VendorOptions {
	return obj.vendorOptionsSlice
}

func (obj *configOptionsVendorOptionsIter) Add() VendorOptions {
	newObj := &otg.VendorOptions{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &vendorOptions{obj: newObj}
	newLibObj.setDefault()
	obj.vendorOptionsSlice = append(obj.vendorOptionsSlice, newLibObj)
	return newLibObj
}

func (obj *configOptionsVendorOptionsIter) Append(items ...VendorOptions) ConfigOptionsVendorOptionsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.vendorOptionsSlice = append(obj.vendorOptionsSlice, item)
	}
	return obj
}

func (obj *configOptionsVendorOptionsIter) Set(index int, newObj VendorOptions) ConfigOptionsVendorOptionsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.vendorOptionsSlice[index] = newObj
	return obj
}
func (obj *configOptionsVendorOptionsIter) Clear() ConfigOptionsVendorOptionsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.VendorOptions{}
		obj.vendorOptionsSlice = []VendorOptions{}
	}
	return obj
}
func (obj *configOptionsVendorOptionsIter) clearHolderSlice() ConfigOptionsVendorOptionsIter {
	if len(obj.vendorOptionsSlice) > 0 {
		obj.vendorOptionsSlice = []VendorOptions{}
	}
	return obj
}
func (obj *configOptionsVendorOptionsIter) appendHolderSlice(item VendorOptions) ConfigOptionsVendorOptionsIter {
	obj.vendorOptionsSlice = append(obj.vendorOptionsSlice, item)
	return obj
}

func (obj *configOptions) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PortOptions != nil {

		obj.PortOptions().validateObj(vObj, set_default)
	}

	if obj.obj.ProtocolOptions != nil {

		obj.ProtocolOptions().validateObj(vObj, set_default)
	}

	if len(obj.obj.VendorOptions) != 0 {

		if set_default {
			obj.VendorOptions().clearHolderSlice()
			for _, item := range obj.obj.VendorOptions {
				obj.VendorOptions().appendHolderSlice(&vendorOptions{obj: item})
			}
		}
		for _, item := range obj.VendorOptions().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *configOptions) setDefault() {

}

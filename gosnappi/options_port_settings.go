package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** OptionsPortSettings *****
type optionsPortSettings struct {
	validation
	obj             *otg.OptionsPortSettings
	marshaller      marshalOptionsPortSettings
	unMarshaller    unMarshalOptionsPortSettings
	protocolsHolder OptionsPortSettingsPortProtocolsIter
}

func NewOptionsPortSettings() OptionsPortSettings {
	obj := optionsPortSettings{obj: &otg.OptionsPortSettings{}}
	obj.setDefault()
	return &obj
}

func (obj *optionsPortSettings) msg() *otg.OptionsPortSettings {
	return obj.obj
}

func (obj *optionsPortSettings) setMsg(msg *otg.OptionsPortSettings) OptionsPortSettings {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaloptionsPortSettings struct {
	obj *optionsPortSettings
}

type marshalOptionsPortSettings interface {
	// ToProto marshals OptionsPortSettings to protobuf object *otg.OptionsPortSettings
	ToProto() (*otg.OptionsPortSettings, error)
	// ToPbText marshals OptionsPortSettings to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals OptionsPortSettings to YAML text
	ToYaml() (string, error)
	// ToJson marshals OptionsPortSettings to JSON text
	ToJson() (string, error)
}

type unMarshaloptionsPortSettings struct {
	obj *optionsPortSettings
}

type unMarshalOptionsPortSettings interface {
	// FromProto unmarshals OptionsPortSettings from protobuf object *otg.OptionsPortSettings
	FromProto(msg *otg.OptionsPortSettings) (OptionsPortSettings, error)
	// FromPbText unmarshals OptionsPortSettings from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals OptionsPortSettings from YAML text
	FromYaml(value string) error
	// FromJson unmarshals OptionsPortSettings from JSON text
	FromJson(value string) error
}

func (obj *optionsPortSettings) Marshal() marshalOptionsPortSettings {
	if obj.marshaller == nil {
		obj.marshaller = &marshaloptionsPortSettings{obj: obj}
	}
	return obj.marshaller
}

func (obj *optionsPortSettings) Unmarshal() unMarshalOptionsPortSettings {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaloptionsPortSettings{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaloptionsPortSettings) ToProto() (*otg.OptionsPortSettings, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaloptionsPortSettings) FromProto(msg *otg.OptionsPortSettings) (OptionsPortSettings, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaloptionsPortSettings) ToPbText() (string, error) {
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

func (m *unMarshaloptionsPortSettings) FromPbText(value string) error {
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

func (m *marshaloptionsPortSettings) ToYaml() (string, error) {
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

func (m *unMarshaloptionsPortSettings) FromYaml(value string) error {
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

func (m *marshaloptionsPortSettings) ToJson() (string, error) {
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

func (m *unMarshaloptionsPortSettings) FromJson(value string) error {
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

func (obj *optionsPortSettings) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *optionsPortSettings) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *optionsPortSettings) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *optionsPortSettings) Clone() (OptionsPortSettings, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOptionsPortSettings()
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

func (obj *optionsPortSettings) setNil() {
	obj.protocolsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// OptionsPortSettings is ****Add proper descriptio****
type OptionsPortSettings interface {
	Validation
	// msg marshals OptionsPortSettings to protobuf object *otg.OptionsPortSettings
	// and doesn't set defaults
	msg() *otg.OptionsPortSettings
	// setMsg unmarshals OptionsPortSettings from protobuf object *otg.OptionsPortSettings
	// and doesn't set defaults
	setMsg(*otg.OptionsPortSettings) OptionsPortSettings
	// provides marshal interface
	Marshal() marshalOptionsPortSettings
	// provides unmarshal interface
	Unmarshal() unMarshalOptionsPortSettings
	// validate validates OptionsPortSettings
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (OptionsPortSettings, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PortName returns string, set in OptionsPortSettings.
	PortName() string
	// SetPortName assigns string provided by user to OptionsPortSettings
	SetPortName(value string) OptionsPortSettings
	// HasPortName checks if PortName has been set in OptionsPortSettings
	HasPortName() bool
	// Protocols returns OptionsPortSettingsPortProtocolsIterIter, set in OptionsPortSettings
	Protocols() OptionsPortSettingsPortProtocolsIter
	setNil()
}

// The name of port for which this settings will be applied to.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// PortName returns a string
func (obj *optionsPortSettings) PortName() string {

	return *obj.obj.PortName

}

// The name of port for which this settings will be applied to.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// PortName returns a string
func (obj *optionsPortSettings) HasPortName() bool {
	return obj.obj.PortName != nil
}

// The name of port for which this settings will be applied to.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// SetPortName sets the string value in the OptionsPortSettings object
func (obj *optionsPortSettings) SetPortName(value string) OptionsPortSettings {

	obj.obj.PortName = &value
	return obj
}

// description is TBD
// Protocols returns a []PortProtocols
func (obj *optionsPortSettings) Protocols() OptionsPortSettingsPortProtocolsIter {
	if len(obj.obj.Protocols) == 0 {
		obj.obj.Protocols = []*otg.PortProtocols{}
	}
	if obj.protocolsHolder == nil {
		obj.protocolsHolder = newOptionsPortSettingsPortProtocolsIter(&obj.obj.Protocols).setMsg(obj)
	}
	return obj.protocolsHolder
}

type optionsPortSettingsPortProtocolsIter struct {
	obj                *optionsPortSettings
	portProtocolsSlice []PortProtocols
	fieldPtr           *[]*otg.PortProtocols
}

func newOptionsPortSettingsPortProtocolsIter(ptr *[]*otg.PortProtocols) OptionsPortSettingsPortProtocolsIter {
	return &optionsPortSettingsPortProtocolsIter{fieldPtr: ptr}
}

type OptionsPortSettingsPortProtocolsIter interface {
	setMsg(*optionsPortSettings) OptionsPortSettingsPortProtocolsIter
	Items() []PortProtocols
	Add() PortProtocols
	Append(items ...PortProtocols) OptionsPortSettingsPortProtocolsIter
	Set(index int, newObj PortProtocols) OptionsPortSettingsPortProtocolsIter
	Clear() OptionsPortSettingsPortProtocolsIter
	clearHolderSlice() OptionsPortSettingsPortProtocolsIter
	appendHolderSlice(item PortProtocols) OptionsPortSettingsPortProtocolsIter
}

func (obj *optionsPortSettingsPortProtocolsIter) setMsg(msg *optionsPortSettings) OptionsPortSettingsPortProtocolsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&portProtocols{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *optionsPortSettingsPortProtocolsIter) Items() []PortProtocols {
	return obj.portProtocolsSlice
}

func (obj *optionsPortSettingsPortProtocolsIter) Add() PortProtocols {
	newObj := &otg.PortProtocols{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &portProtocols{obj: newObj}
	newLibObj.setDefault()
	obj.portProtocolsSlice = append(obj.portProtocolsSlice, newLibObj)
	return newLibObj
}

func (obj *optionsPortSettingsPortProtocolsIter) Append(items ...PortProtocols) OptionsPortSettingsPortProtocolsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.portProtocolsSlice = append(obj.portProtocolsSlice, item)
	}
	return obj
}

func (obj *optionsPortSettingsPortProtocolsIter) Set(index int, newObj PortProtocols) OptionsPortSettingsPortProtocolsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.portProtocolsSlice[index] = newObj
	return obj
}
func (obj *optionsPortSettingsPortProtocolsIter) Clear() OptionsPortSettingsPortProtocolsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PortProtocols{}
		obj.portProtocolsSlice = []PortProtocols{}
	}
	return obj
}
func (obj *optionsPortSettingsPortProtocolsIter) clearHolderSlice() OptionsPortSettingsPortProtocolsIter {
	if len(obj.portProtocolsSlice) > 0 {
		obj.portProtocolsSlice = []PortProtocols{}
	}
	return obj
}
func (obj *optionsPortSettingsPortProtocolsIter) appendHolderSlice(item PortProtocols) OptionsPortSettingsPortProtocolsIter {
	obj.portProtocolsSlice = append(obj.portProtocolsSlice, item)
	return obj
}

func (obj *optionsPortSettings) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Protocols) != 0 {

		if set_default {
			obj.Protocols().clearHolderSlice()
			for _, item := range obj.obj.Protocols {
				obj.Protocols().appendHolderSlice(&portProtocols{obj: item})
			}
		}
		for _, item := range obj.Protocols().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *optionsPortSettings) setDefault() {

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PortProtocols *****
type portProtocols struct {
	validation
	obj          *otg.PortProtocols
	marshaller   marshalPortProtocols
	unMarshaller unMarshalPortProtocols
	rocev2Holder PortProtocolsRocev2PerPortSettingsIter
}

func NewPortProtocols() PortProtocols {
	obj := portProtocols{obj: &otg.PortProtocols{}}
	obj.setDefault()
	return &obj
}

func (obj *portProtocols) msg() *otg.PortProtocols {
	return obj.obj
}

func (obj *portProtocols) setMsg(msg *otg.PortProtocols) PortProtocols {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalportProtocols struct {
	obj *portProtocols
}

type marshalPortProtocols interface {
	// ToProto marshals PortProtocols to protobuf object *otg.PortProtocols
	ToProto() (*otg.PortProtocols, error)
	// ToPbText marshals PortProtocols to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PortProtocols to YAML text
	ToYaml() (string, error)
	// ToJson marshals PortProtocols to JSON text
	ToJson() (string, error)
}

type unMarshalportProtocols struct {
	obj *portProtocols
}

type unMarshalPortProtocols interface {
	// FromProto unmarshals PortProtocols from protobuf object *otg.PortProtocols
	FromProto(msg *otg.PortProtocols) (PortProtocols, error)
	// FromPbText unmarshals PortProtocols from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PortProtocols from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PortProtocols from JSON text
	FromJson(value string) error
}

func (obj *portProtocols) Marshal() marshalPortProtocols {
	if obj.marshaller == nil {
		obj.marshaller = &marshalportProtocols{obj: obj}
	}
	return obj.marshaller
}

func (obj *portProtocols) Unmarshal() unMarshalPortProtocols {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalportProtocols{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalportProtocols) ToProto() (*otg.PortProtocols, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalportProtocols) FromProto(msg *otg.PortProtocols) (PortProtocols, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalportProtocols) ToPbText() (string, error) {
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

func (m *unMarshalportProtocols) FromPbText(value string) error {
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

func (m *marshalportProtocols) ToYaml() (string, error) {
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

func (m *unMarshalportProtocols) FromYaml(value string) error {
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

func (m *marshalportProtocols) ToJson() (string, error) {
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

func (m *unMarshalportProtocols) FromJson(value string) error {
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

func (obj *portProtocols) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *portProtocols) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *portProtocols) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *portProtocols) Clone() (PortProtocols, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPortProtocols()
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

func (obj *portProtocols) setNil() {
	obj.rocev2Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PortProtocols is supprted protocols.
type PortProtocols interface {
	Validation
	// msg marshals PortProtocols to protobuf object *otg.PortProtocols
	// and doesn't set defaults
	msg() *otg.PortProtocols
	// setMsg unmarshals PortProtocols from protobuf object *otg.PortProtocols
	// and doesn't set defaults
	setMsg(*otg.PortProtocols) PortProtocols
	// provides marshal interface
	Marshal() marshalPortProtocols
	// provides unmarshal interface
	Unmarshal() unMarshalPortProtocols
	// validate validates PortProtocols
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PortProtocols, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// DummyField returns uint32, set in PortProtocols.
	DummyField() uint32
	// SetDummyField assigns uint32 provided by user to PortProtocols
	SetDummyField(value uint32) PortProtocols
	// HasDummyField checks if DummyField has been set in PortProtocols
	HasDummyField() bool
	// Choice returns PortProtocolsChoiceEnum, set in PortProtocols
	Choice() PortProtocolsChoiceEnum
	// setChoice assigns PortProtocolsChoiceEnum provided by user to PortProtocols
	setChoice(value PortProtocolsChoiceEnum) PortProtocols
	// HasChoice checks if Choice has been set in PortProtocols
	HasChoice() bool
	// Rocev2 returns PortProtocolsRocev2PerPortSettingsIterIter, set in PortProtocols
	Rocev2() PortProtocolsRocev2PerPortSettingsIter
	setNil()
}

// Dummy
// DummyField returns a uint32
func (obj *portProtocols) DummyField() uint32 {

	return *obj.obj.DummyField

}

// Dummy
// DummyField returns a uint32
func (obj *portProtocols) HasDummyField() bool {
	return obj.obj.DummyField != nil
}

// Dummy
// SetDummyField sets the uint32 value in the PortProtocols object
func (obj *portProtocols) SetDummyField(value uint32) PortProtocols {

	obj.obj.DummyField = &value
	return obj
}

type PortProtocolsChoiceEnum string

// Enum of Choice on PortProtocols
var PortProtocolsChoice = struct {
	ROCEV2 PortProtocolsChoiceEnum
}{
	ROCEV2: PortProtocolsChoiceEnum("rocev2"),
}

func (obj *portProtocols) Choice() PortProtocolsChoiceEnum {
	return PortProtocolsChoiceEnum(obj.obj.Choice.Enum().String())
}

// list of protocols that have per port settings
// Choice returns a string
func (obj *portProtocols) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *portProtocols) setChoice(value PortProtocolsChoiceEnum) PortProtocols {
	intValue, ok := otg.PortProtocols_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PortProtocolsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PortProtocols_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Rocev2 = nil
	obj.rocev2Holder = nil

	if value == PortProtocolsChoice.ROCEV2 {
		obj.obj.Rocev2 = []*otg.Rocev2PerPortSettings{}
	}

	return obj
}

// Rocev2 Port Settings.
// Rocev2 returns a []Rocev2PerPortSettings
func (obj *portProtocols) Rocev2() PortProtocolsRocev2PerPortSettingsIter {
	if len(obj.obj.Rocev2) == 0 {
		obj.setChoice(PortProtocolsChoice.ROCEV2)
	}
	if obj.rocev2Holder == nil {
		obj.rocev2Holder = newPortProtocolsRocev2PerPortSettingsIter(&obj.obj.Rocev2).setMsg(obj)
	}
	return obj.rocev2Holder
}

type portProtocolsRocev2PerPortSettingsIter struct {
	obj                        *portProtocols
	rocev2PerPortSettingsSlice []Rocev2PerPortSettings
	fieldPtr                   *[]*otg.Rocev2PerPortSettings
}

func newPortProtocolsRocev2PerPortSettingsIter(ptr *[]*otg.Rocev2PerPortSettings) PortProtocolsRocev2PerPortSettingsIter {
	return &portProtocolsRocev2PerPortSettingsIter{fieldPtr: ptr}
}

type PortProtocolsRocev2PerPortSettingsIter interface {
	setMsg(*portProtocols) PortProtocolsRocev2PerPortSettingsIter
	Items() []Rocev2PerPortSettings
	Add() Rocev2PerPortSettings
	Append(items ...Rocev2PerPortSettings) PortProtocolsRocev2PerPortSettingsIter
	Set(index int, newObj Rocev2PerPortSettings) PortProtocolsRocev2PerPortSettingsIter
	Clear() PortProtocolsRocev2PerPortSettingsIter
	clearHolderSlice() PortProtocolsRocev2PerPortSettingsIter
	appendHolderSlice(item Rocev2PerPortSettings) PortProtocolsRocev2PerPortSettingsIter
}

func (obj *portProtocolsRocev2PerPortSettingsIter) setMsg(msg *portProtocols) PortProtocolsRocev2PerPortSettingsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&rocev2PerPortSettings{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *portProtocolsRocev2PerPortSettingsIter) Items() []Rocev2PerPortSettings {
	return obj.rocev2PerPortSettingsSlice
}

func (obj *portProtocolsRocev2PerPortSettingsIter) Add() Rocev2PerPortSettings {
	newObj := &otg.Rocev2PerPortSettings{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &rocev2PerPortSettings{obj: newObj}
	newLibObj.setDefault()
	obj.rocev2PerPortSettingsSlice = append(obj.rocev2PerPortSettingsSlice, newLibObj)
	return newLibObj
}

func (obj *portProtocolsRocev2PerPortSettingsIter) Append(items ...Rocev2PerPortSettings) PortProtocolsRocev2PerPortSettingsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.rocev2PerPortSettingsSlice = append(obj.rocev2PerPortSettingsSlice, item)
	}
	return obj
}

func (obj *portProtocolsRocev2PerPortSettingsIter) Set(index int, newObj Rocev2PerPortSettings) PortProtocolsRocev2PerPortSettingsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.rocev2PerPortSettingsSlice[index] = newObj
	return obj
}
func (obj *portProtocolsRocev2PerPortSettingsIter) Clear() PortProtocolsRocev2PerPortSettingsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Rocev2PerPortSettings{}
		obj.rocev2PerPortSettingsSlice = []Rocev2PerPortSettings{}
	}
	return obj
}
func (obj *portProtocolsRocev2PerPortSettingsIter) clearHolderSlice() PortProtocolsRocev2PerPortSettingsIter {
	if len(obj.rocev2PerPortSettingsSlice) > 0 {
		obj.rocev2PerPortSettingsSlice = []Rocev2PerPortSettings{}
	}
	return obj
}
func (obj *portProtocolsRocev2PerPortSettingsIter) appendHolderSlice(item Rocev2PerPortSettings) PortProtocolsRocev2PerPortSettingsIter {
	obj.rocev2PerPortSettingsSlice = append(obj.rocev2PerPortSettingsSlice, item)
	return obj
}

func (obj *portProtocols) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.DummyField != nil {

		if *obj.obj.DummyField > 14000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PortProtocols.DummyField <= 14000 but Got %d", *obj.obj.DummyField))
		}

	}

	if len(obj.obj.Rocev2) != 0 {

		if set_default {
			obj.Rocev2().clearHolderSlice()
			for _, item := range obj.obj.Rocev2 {
				obj.Rocev2().appendHolderSlice(&rocev2PerPortSettings{obj: item})
			}
		}
		for _, item := range obj.Rocev2().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *portProtocols) setDefault() {
	var choices_set int = 0
	var choice PortProtocolsChoiceEnum

	if len(obj.obj.Rocev2) > 0 {
		choices_set += 1
		choice = PortProtocolsChoice.ROCEV2
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PortProtocolsChoice.ROCEV2)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PortProtocols")
			}
		} else {
			intVal := otg.PortProtocols_Choice_Enum_value[string(choice)]
			enumValue := otg.PortProtocols_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.DummyField == nil {
		obj.SetDummyField(1024)
	}

}

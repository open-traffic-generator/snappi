package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MkaBasicKeySource *****
type mkaBasicKeySource struct {
	validation
	obj          *otg.MkaBasicKeySource
	marshaller   marshalMkaBasicKeySource
	unMarshaller unMarshalMkaBasicKeySource
	psksHolder   MkaBasicKeySourceMkaBasicKeySourcePskIter
}

func NewMkaBasicKeySource() MkaBasicKeySource {
	obj := mkaBasicKeySource{obj: &otg.MkaBasicKeySource{}}
	obj.setDefault()
	return &obj
}

func (obj *mkaBasicKeySource) msg() *otg.MkaBasicKeySource {
	return obj.obj
}

func (obj *mkaBasicKeySource) setMsg(msg *otg.MkaBasicKeySource) MkaBasicKeySource {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmkaBasicKeySource struct {
	obj *mkaBasicKeySource
}

type marshalMkaBasicKeySource interface {
	// ToProto marshals MkaBasicKeySource to protobuf object *otg.MkaBasicKeySource
	ToProto() (*otg.MkaBasicKeySource, error)
	// ToPbText marshals MkaBasicKeySource to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MkaBasicKeySource to YAML text
	ToYaml() (string, error)
	// ToJson marshals MkaBasicKeySource to JSON text
	ToJson() (string, error)
}

type unMarshalmkaBasicKeySource struct {
	obj *mkaBasicKeySource
}

type unMarshalMkaBasicKeySource interface {
	// FromProto unmarshals MkaBasicKeySource from protobuf object *otg.MkaBasicKeySource
	FromProto(msg *otg.MkaBasicKeySource) (MkaBasicKeySource, error)
	// FromPbText unmarshals MkaBasicKeySource from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MkaBasicKeySource from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MkaBasicKeySource from JSON text
	FromJson(value string) error
}

func (obj *mkaBasicKeySource) Marshal() marshalMkaBasicKeySource {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmkaBasicKeySource{obj: obj}
	}
	return obj.marshaller
}

func (obj *mkaBasicKeySource) Unmarshal() unMarshalMkaBasicKeySource {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmkaBasicKeySource{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmkaBasicKeySource) ToProto() (*otg.MkaBasicKeySource, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmkaBasicKeySource) FromProto(msg *otg.MkaBasicKeySource) (MkaBasicKeySource, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmkaBasicKeySource) ToPbText() (string, error) {
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

func (m *unMarshalmkaBasicKeySource) FromPbText(value string) error {
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

func (m *marshalmkaBasicKeySource) ToYaml() (string, error) {
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

func (m *unMarshalmkaBasicKeySource) FromYaml(value string) error {
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

func (m *marshalmkaBasicKeySource) ToJson() (string, error) {
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

func (m *unMarshalmkaBasicKeySource) FromJson(value string) error {
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

func (obj *mkaBasicKeySource) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mkaBasicKeySource) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mkaBasicKeySource) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mkaBasicKeySource) Clone() (MkaBasicKeySource, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMkaBasicKeySource()
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

func (obj *mkaBasicKeySource) setNil() {
	obj.psksHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MkaBasicKeySource is the container for static key settings.
type MkaBasicKeySource interface {
	Validation
	// msg marshals MkaBasicKeySource to protobuf object *otg.MkaBasicKeySource
	// and doesn't set defaults
	msg() *otg.MkaBasicKeySource
	// setMsg unmarshals MkaBasicKeySource from protobuf object *otg.MkaBasicKeySource
	// and doesn't set defaults
	setMsg(*otg.MkaBasicKeySource) MkaBasicKeySource
	// provides marshal interface
	Marshal() marshalMkaBasicKeySource
	// provides unmarshal interface
	Unmarshal() unMarshalMkaBasicKeySource
	// validate validates MkaBasicKeySource
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MkaBasicKeySource, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MkaBasicKeySourceChoiceEnum, set in MkaBasicKeySource
	Choice() MkaBasicKeySourceChoiceEnum
	// setChoice assigns MkaBasicKeySourceChoiceEnum provided by user to MkaBasicKeySource
	setChoice(value MkaBasicKeySourceChoiceEnum) MkaBasicKeySource
	// HasChoice checks if Choice has been set in MkaBasicKeySource
	HasChoice() bool
	// getter for Msk to set choice.
	Msk()
	// getter for Psk to set choice.
	Psk()
	// Psks returns MkaBasicKeySourceMkaBasicKeySourcePskIterIter, set in MkaBasicKeySource
	Psks() MkaBasicKeySourceMkaBasicKeySourcePskIter
	setNil()
}

type MkaBasicKeySourceChoiceEnum string

// Enum of Choice on MkaBasicKeySource
var MkaBasicKeySourceChoice = struct {
	PSK MkaBasicKeySourceChoiceEnum
	MSK MkaBasicKeySourceChoiceEnum
}{
	PSK: MkaBasicKeySourceChoiceEnum("psk"),
	MSK: MkaBasicKeySourceChoiceEnum("msk"),
}

func (obj *mkaBasicKeySource) Choice() MkaBasicKeySourceChoiceEnum {
	return MkaBasicKeySourceChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Msk to set choice
func (obj *mkaBasicKeySource) Msk() {
	obj.setChoice(MkaBasicKeySourceChoice.MSK)
}

// getter for Psk to set choice
func (obj *mkaBasicKeySource) Psk() {
	obj.setChoice(MkaBasicKeySourceChoice.PSK)
}

// Key source. Choose one from PSK or MSK.
// Choice returns a string
func (obj *mkaBasicKeySource) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *mkaBasicKeySource) setChoice(value MkaBasicKeySourceChoiceEnum) MkaBasicKeySource {
	intValue, ok := otg.MkaBasicKeySource_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MkaBasicKeySourceChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MkaBasicKeySource_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

// Pre-shared Keys.
// Psks returns a []MkaBasicKeySourcePsk
func (obj *mkaBasicKeySource) Psks() MkaBasicKeySourceMkaBasicKeySourcePskIter {
	if len(obj.obj.Psks) == 0 {
		obj.obj.Psks = []*otg.MkaBasicKeySourcePsk{}
	}
	if obj.psksHolder == nil {
		obj.psksHolder = newMkaBasicKeySourceMkaBasicKeySourcePskIter(&obj.obj.Psks).setMsg(obj)
	}
	return obj.psksHolder
}

type mkaBasicKeySourceMkaBasicKeySourcePskIter struct {
	obj                       *mkaBasicKeySource
	mkaBasicKeySourcePskSlice []MkaBasicKeySourcePsk
	fieldPtr                  *[]*otg.MkaBasicKeySourcePsk
}

func newMkaBasicKeySourceMkaBasicKeySourcePskIter(ptr *[]*otg.MkaBasicKeySourcePsk) MkaBasicKeySourceMkaBasicKeySourcePskIter {
	return &mkaBasicKeySourceMkaBasicKeySourcePskIter{fieldPtr: ptr}
}

type MkaBasicKeySourceMkaBasicKeySourcePskIter interface {
	setMsg(*mkaBasicKeySource) MkaBasicKeySourceMkaBasicKeySourcePskIter
	Items() []MkaBasicKeySourcePsk
	Add() MkaBasicKeySourcePsk
	Append(items ...MkaBasicKeySourcePsk) MkaBasicKeySourceMkaBasicKeySourcePskIter
	Set(index int, newObj MkaBasicKeySourcePsk) MkaBasicKeySourceMkaBasicKeySourcePskIter
	Clear() MkaBasicKeySourceMkaBasicKeySourcePskIter
	clearHolderSlice() MkaBasicKeySourceMkaBasicKeySourcePskIter
	appendHolderSlice(item MkaBasicKeySourcePsk) MkaBasicKeySourceMkaBasicKeySourcePskIter
}

func (obj *mkaBasicKeySourceMkaBasicKeySourcePskIter) setMsg(msg *mkaBasicKeySource) MkaBasicKeySourceMkaBasicKeySourcePskIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&mkaBasicKeySourcePsk{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *mkaBasicKeySourceMkaBasicKeySourcePskIter) Items() []MkaBasicKeySourcePsk {
	return obj.mkaBasicKeySourcePskSlice
}

func (obj *mkaBasicKeySourceMkaBasicKeySourcePskIter) Add() MkaBasicKeySourcePsk {
	newObj := &otg.MkaBasicKeySourcePsk{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &mkaBasicKeySourcePsk{obj: newObj}
	newLibObj.setDefault()
	obj.mkaBasicKeySourcePskSlice = append(obj.mkaBasicKeySourcePskSlice, newLibObj)
	return newLibObj
}

func (obj *mkaBasicKeySourceMkaBasicKeySourcePskIter) Append(items ...MkaBasicKeySourcePsk) MkaBasicKeySourceMkaBasicKeySourcePskIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.mkaBasicKeySourcePskSlice = append(obj.mkaBasicKeySourcePskSlice, item)
	}
	return obj
}

func (obj *mkaBasicKeySourceMkaBasicKeySourcePskIter) Set(index int, newObj MkaBasicKeySourcePsk) MkaBasicKeySourceMkaBasicKeySourcePskIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.mkaBasicKeySourcePskSlice[index] = newObj
	return obj
}
func (obj *mkaBasicKeySourceMkaBasicKeySourcePskIter) Clear() MkaBasicKeySourceMkaBasicKeySourcePskIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.MkaBasicKeySourcePsk{}
		obj.mkaBasicKeySourcePskSlice = []MkaBasicKeySourcePsk{}
	}
	return obj
}
func (obj *mkaBasicKeySourceMkaBasicKeySourcePskIter) clearHolderSlice() MkaBasicKeySourceMkaBasicKeySourcePskIter {
	if len(obj.mkaBasicKeySourcePskSlice) > 0 {
		obj.mkaBasicKeySourcePskSlice = []MkaBasicKeySourcePsk{}
	}
	return obj
}
func (obj *mkaBasicKeySourceMkaBasicKeySourcePskIter) appendHolderSlice(item MkaBasicKeySourcePsk) MkaBasicKeySourceMkaBasicKeySourcePskIter {
	obj.mkaBasicKeySourcePskSlice = append(obj.mkaBasicKeySourcePskSlice, item)
	return obj
}

func (obj *mkaBasicKeySource) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Psks) != 0 {

		if set_default {
			obj.Psks().clearHolderSlice()
			for _, item := range obj.obj.Psks {
				obj.Psks().appendHolderSlice(&mkaBasicKeySourcePsk{obj: item})
			}
		}
		for _, item := range obj.Psks().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *mkaBasicKeySource) setDefault() {
	var choices_set int = 0
	var choice MkaBasicKeySourceChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MkaBasicKeySourceChoice.PSK)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MkaBasicKeySource")
			}
		} else {
			intVal := otg.MkaBasicKeySource_Choice_Enum_value[string(choice)]
			enumValue := otg.MkaBasicKeySource_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

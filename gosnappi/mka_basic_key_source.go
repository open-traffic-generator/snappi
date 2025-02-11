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
	obj            *otg.MkaBasicKeySource
	marshaller     marshalMkaBasicKeySource
	unMarshaller   unMarshalMkaBasicKeySource
	pskChainHolder MkaBasicKeySourcePskChain
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
	obj.pskChainHolder = nil
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
	// PskChain returns MkaBasicKeySourcePskChain, set in MkaBasicKeySource.
	// MkaBasicKeySourcePskChain is the container for PSK chain settings.
	PskChain() MkaBasicKeySourcePskChain
	// SetPskChain assigns MkaBasicKeySourcePskChain provided by user to MkaBasicKeySource.
	// MkaBasicKeySourcePskChain is the container for PSK chain settings.
	SetPskChain(value MkaBasicKeySourcePskChain) MkaBasicKeySource
	// HasPskChain checks if PskChain has been set in MkaBasicKeySource
	HasPskChain() bool
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

// PSK chain.
// PskChain returns a MkaBasicKeySourcePskChain
func (obj *mkaBasicKeySource) PskChain() MkaBasicKeySourcePskChain {
	if obj.obj.PskChain == nil {
		obj.obj.PskChain = NewMkaBasicKeySourcePskChain().msg()
	}
	if obj.pskChainHolder == nil {
		obj.pskChainHolder = &mkaBasicKeySourcePskChain{obj: obj.obj.PskChain}
	}
	return obj.pskChainHolder
}

// PSK chain.
// PskChain returns a MkaBasicKeySourcePskChain
func (obj *mkaBasicKeySource) HasPskChain() bool {
	return obj.obj.PskChain != nil
}

// PSK chain.
// SetPskChain sets the MkaBasicKeySourcePskChain value in the MkaBasicKeySource object
func (obj *mkaBasicKeySource) SetPskChain(value MkaBasicKeySourcePskChain) MkaBasicKeySource {

	obj.pskChainHolder = nil
	obj.obj.PskChain = value.msg()

	return obj
}

func (obj *mkaBasicKeySource) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PskChain != nil {

		obj.PskChain().validateObj(vObj, set_default)
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

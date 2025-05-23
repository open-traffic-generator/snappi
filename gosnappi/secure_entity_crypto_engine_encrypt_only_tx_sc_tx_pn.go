package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityCryptoEngineEncryptOnlyTxScTxPn *****
type secureEntityCryptoEngineEncryptOnlyTxScTxPn struct {
	validation
	obj                *otg.SecureEntityCryptoEngineEncryptOnlyTxScTxPn
	marshaller         marshalSecureEntityCryptoEngineEncryptOnlyTxScTxPn
	unMarshaller       unMarshalSecureEntityCryptoEngineEncryptOnlyTxScTxPn
	fixedHolder        SecureEntityCryptoEngineEncryptOnlyFixedPn
	incrementingHolder SecureEntityCryptoEngineEncryptOnlyIncrementingPn
}

func NewSecureEntityCryptoEngineEncryptOnlyTxScTxPn() SecureEntityCryptoEngineEncryptOnlyTxScTxPn {
	obj := secureEntityCryptoEngineEncryptOnlyTxScTxPn{obj: &otg.SecureEntityCryptoEngineEncryptOnlyTxScTxPn{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) msg() *otg.SecureEntityCryptoEngineEncryptOnlyTxScTxPn {
	return obj.obj
}

func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) setMsg(msg *otg.SecureEntityCryptoEngineEncryptOnlyTxScTxPn) SecureEntityCryptoEngineEncryptOnlyTxScTxPn {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityCryptoEngineEncryptOnlyTxScTxPn struct {
	obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn
}

type marshalSecureEntityCryptoEngineEncryptOnlyTxScTxPn interface {
	// ToProto marshals SecureEntityCryptoEngineEncryptOnlyTxScTxPn to protobuf object *otg.SecureEntityCryptoEngineEncryptOnlyTxScTxPn
	ToProto() (*otg.SecureEntityCryptoEngineEncryptOnlyTxScTxPn, error)
	// ToPbText marshals SecureEntityCryptoEngineEncryptOnlyTxScTxPn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityCryptoEngineEncryptOnlyTxScTxPn to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityCryptoEngineEncryptOnlyTxScTxPn to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals SecureEntityCryptoEngineEncryptOnlyTxScTxPn to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalsecureEntityCryptoEngineEncryptOnlyTxScTxPn struct {
	obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn
}

type unMarshalSecureEntityCryptoEngineEncryptOnlyTxScTxPn interface {
	// FromProto unmarshals SecureEntityCryptoEngineEncryptOnlyTxScTxPn from protobuf object *otg.SecureEntityCryptoEngineEncryptOnlyTxScTxPn
	FromProto(msg *otg.SecureEntityCryptoEngineEncryptOnlyTxScTxPn) (SecureEntityCryptoEngineEncryptOnlyTxScTxPn, error)
	// FromPbText unmarshals SecureEntityCryptoEngineEncryptOnlyTxScTxPn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityCryptoEngineEncryptOnlyTxScTxPn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityCryptoEngineEncryptOnlyTxScTxPn from JSON text
	FromJson(value string) error
}

func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) Marshal() marshalSecureEntityCryptoEngineEncryptOnlyTxScTxPn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityCryptoEngineEncryptOnlyTxScTxPn{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) Unmarshal() unMarshalSecureEntityCryptoEngineEncryptOnlyTxScTxPn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityCryptoEngineEncryptOnlyTxScTxPn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityCryptoEngineEncryptOnlyTxScTxPn) ToProto() (*otg.SecureEntityCryptoEngineEncryptOnlyTxScTxPn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityCryptoEngineEncryptOnlyTxScTxPn) FromProto(msg *otg.SecureEntityCryptoEngineEncryptOnlyTxScTxPn) (SecureEntityCryptoEngineEncryptOnlyTxScTxPn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityCryptoEngineEncryptOnlyTxScTxPn) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptOnlyTxScTxPn) FromPbText(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptOnlyTxScTxPn) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptOnlyTxScTxPn) FromYaml(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptOnlyTxScTxPn) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalsecureEntityCryptoEngineEncryptOnlyTxScTxPn) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptOnlyTxScTxPn) FromJson(value string) error {
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

func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) Clone() (SecureEntityCryptoEngineEncryptOnlyTxScTxPn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityCryptoEngineEncryptOnlyTxScTxPn()
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

func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) setNil() {
	obj.fixedHolder = nil
	obj.incrementingHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// SecureEntityCryptoEngineEncryptOnlyTxScTxPn is tx packet number(PN) configuration.
type SecureEntityCryptoEngineEncryptOnlyTxScTxPn interface {
	Validation
	// msg marshals SecureEntityCryptoEngineEncryptOnlyTxScTxPn to protobuf object *otg.SecureEntityCryptoEngineEncryptOnlyTxScTxPn
	// and doesn't set defaults
	msg() *otg.SecureEntityCryptoEngineEncryptOnlyTxScTxPn
	// setMsg unmarshals SecureEntityCryptoEngineEncryptOnlyTxScTxPn from protobuf object *otg.SecureEntityCryptoEngineEncryptOnlyTxScTxPn
	// and doesn't set defaults
	setMsg(*otg.SecureEntityCryptoEngineEncryptOnlyTxScTxPn) SecureEntityCryptoEngineEncryptOnlyTxScTxPn
	// provides marshal interface
	Marshal() marshalSecureEntityCryptoEngineEncryptOnlyTxScTxPn
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityCryptoEngineEncryptOnlyTxScTxPn
	// validate validates SecureEntityCryptoEngineEncryptOnlyTxScTxPn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityCryptoEngineEncryptOnlyTxScTxPn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns SecureEntityCryptoEngineEncryptOnlyTxScTxPnChoiceEnum, set in SecureEntityCryptoEngineEncryptOnlyTxScTxPn
	Choice() SecureEntityCryptoEngineEncryptOnlyTxScTxPnChoiceEnum
	// setChoice assigns SecureEntityCryptoEngineEncryptOnlyTxScTxPnChoiceEnum provided by user to SecureEntityCryptoEngineEncryptOnlyTxScTxPn
	setChoice(value SecureEntityCryptoEngineEncryptOnlyTxScTxPnChoiceEnum) SecureEntityCryptoEngineEncryptOnlyTxScTxPn
	// HasChoice checks if Choice has been set in SecureEntityCryptoEngineEncryptOnlyTxScTxPn
	HasChoice() bool
	// getter for IncrementingPn to set choice.
	IncrementingPn()
	// getter for FixedPn to set choice.
	FixedPn()
	// Fixed returns SecureEntityCryptoEngineEncryptOnlyFixedPn, set in SecureEntityCryptoEngineEncryptOnlyTxScTxPn.
	// SecureEntityCryptoEngineEncryptOnlyFixedPn is fixed packet number(PN) configuration.
	Fixed() SecureEntityCryptoEngineEncryptOnlyFixedPn
	// SetFixed assigns SecureEntityCryptoEngineEncryptOnlyFixedPn provided by user to SecureEntityCryptoEngineEncryptOnlyTxScTxPn.
	// SecureEntityCryptoEngineEncryptOnlyFixedPn is fixed packet number(PN) configuration.
	SetFixed(value SecureEntityCryptoEngineEncryptOnlyFixedPn) SecureEntityCryptoEngineEncryptOnlyTxScTxPn
	// HasFixed checks if Fixed has been set in SecureEntityCryptoEngineEncryptOnlyTxScTxPn
	HasFixed() bool
	// Incrementing returns SecureEntityCryptoEngineEncryptOnlyIncrementingPn, set in SecureEntityCryptoEngineEncryptOnlyTxScTxPn.
	// SecureEntityCryptoEngineEncryptOnlyIncrementingPn is incrementing packet number(PN) configuration.
	Incrementing() SecureEntityCryptoEngineEncryptOnlyIncrementingPn
	// SetIncrementing assigns SecureEntityCryptoEngineEncryptOnlyIncrementingPn provided by user to SecureEntityCryptoEngineEncryptOnlyTxScTxPn.
	// SecureEntityCryptoEngineEncryptOnlyIncrementingPn is incrementing packet number(PN) configuration.
	SetIncrementing(value SecureEntityCryptoEngineEncryptOnlyIncrementingPn) SecureEntityCryptoEngineEncryptOnlyTxScTxPn
	// HasIncrementing checks if Incrementing has been set in SecureEntityCryptoEngineEncryptOnlyTxScTxPn
	HasIncrementing() bool
	setNil()
}

type SecureEntityCryptoEngineEncryptOnlyTxScTxPnChoiceEnum string

// Enum of Choice on SecureEntityCryptoEngineEncryptOnlyTxScTxPn
var SecureEntityCryptoEngineEncryptOnlyTxScTxPnChoice = struct {
	FIXED_PN        SecureEntityCryptoEngineEncryptOnlyTxScTxPnChoiceEnum
	INCREMENTING_PN SecureEntityCryptoEngineEncryptOnlyTxScTxPnChoiceEnum
}{
	FIXED_PN:        SecureEntityCryptoEngineEncryptOnlyTxScTxPnChoiceEnum("fixed_pn"),
	INCREMENTING_PN: SecureEntityCryptoEngineEncryptOnlyTxScTxPnChoiceEnum("incrementing_pn"),
}

func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) Choice() SecureEntityCryptoEngineEncryptOnlyTxScTxPnChoiceEnum {
	return SecureEntityCryptoEngineEncryptOnlyTxScTxPnChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for IncrementingPn to set choice
func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) IncrementingPn() {
	obj.setChoice(SecureEntityCryptoEngineEncryptOnlyTxScTxPnChoice.INCREMENTING_PN)
}

// getter for FixedPn to set choice
func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) FixedPn() {
	obj.setChoice(SecureEntityCryptoEngineEncryptOnlyTxScTxPnChoice.FIXED_PN)
}

// Types of Tx packet number(PN) series. Supported choices: 1) fixed PN - MACsec packets will be sent out with the configured fixed PN or lower half of configured fixed XPN. 2) incrementing PN - MACsec packets will be sent out by single device with an incrementing PN or XPN.
// Choice returns a string
func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) setChoice(value SecureEntityCryptoEngineEncryptOnlyTxScTxPnChoiceEnum) SecureEntityCryptoEngineEncryptOnlyTxScTxPn {
	intValue, ok := otg.SecureEntityCryptoEngineEncryptOnlyTxScTxPn_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on SecureEntityCryptoEngineEncryptOnlyTxScTxPnChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.SecureEntityCryptoEngineEncryptOnlyTxScTxPn_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

// description is TBD
// Fixed returns a SecureEntityCryptoEngineEncryptOnlyFixedPn
func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) Fixed() SecureEntityCryptoEngineEncryptOnlyFixedPn {
	if obj.obj.Fixed == nil {
		obj.obj.Fixed = NewSecureEntityCryptoEngineEncryptOnlyFixedPn().msg()
	}
	if obj.fixedHolder == nil {
		obj.fixedHolder = &secureEntityCryptoEngineEncryptOnlyFixedPn{obj: obj.obj.Fixed}
	}
	return obj.fixedHolder
}

// description is TBD
// Fixed returns a SecureEntityCryptoEngineEncryptOnlyFixedPn
func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) HasFixed() bool {
	return obj.obj.Fixed != nil
}

// description is TBD
// SetFixed sets the SecureEntityCryptoEngineEncryptOnlyFixedPn value in the SecureEntityCryptoEngineEncryptOnlyTxScTxPn object
func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) SetFixed(value SecureEntityCryptoEngineEncryptOnlyFixedPn) SecureEntityCryptoEngineEncryptOnlyTxScTxPn {

	obj.fixedHolder = nil
	obj.obj.Fixed = value.msg()

	return obj
}

// description is TBD
// Incrementing returns a SecureEntityCryptoEngineEncryptOnlyIncrementingPn
func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) Incrementing() SecureEntityCryptoEngineEncryptOnlyIncrementingPn {
	if obj.obj.Incrementing == nil {
		obj.obj.Incrementing = NewSecureEntityCryptoEngineEncryptOnlyIncrementingPn().msg()
	}
	if obj.incrementingHolder == nil {
		obj.incrementingHolder = &secureEntityCryptoEngineEncryptOnlyIncrementingPn{obj: obj.obj.Incrementing}
	}
	return obj.incrementingHolder
}

// description is TBD
// Incrementing returns a SecureEntityCryptoEngineEncryptOnlyIncrementingPn
func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) HasIncrementing() bool {
	return obj.obj.Incrementing != nil
}

// description is TBD
// SetIncrementing sets the SecureEntityCryptoEngineEncryptOnlyIncrementingPn value in the SecureEntityCryptoEngineEncryptOnlyTxScTxPn object
func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) SetIncrementing(value SecureEntityCryptoEngineEncryptOnlyIncrementingPn) SecureEntityCryptoEngineEncryptOnlyTxScTxPn {

	obj.incrementingHolder = nil
	obj.obj.Incrementing = value.msg()

	return obj
}

func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Fixed != nil {

		obj.Fixed().validateObj(vObj, set_default)
	}

	if obj.obj.Incrementing != nil {

		obj.Incrementing().validateObj(vObj, set_default)
	}

}

func (obj *secureEntityCryptoEngineEncryptOnlyTxScTxPn) setDefault() {
	var choices_set int = 0
	var choice SecureEntityCryptoEngineEncryptOnlyTxScTxPnChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(SecureEntityCryptoEngineEncryptOnlyTxScTxPnChoice.FIXED_PN)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in SecureEntityCryptoEngineEncryptOnlyTxScTxPn")
			}
		} else {
			intVal := otg.SecureEntityCryptoEngineEncryptOnlyTxScTxPn_Choice_Enum_value[string(choice)]
			enumValue := otg.SecureEntityCryptoEngineEncryptOnlyTxScTxPn_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineStatelessEncryptionOnlyTxPn *****
type macsecCryptoEngineStatelessEncryptionOnlyTxPn struct {
	validation
	obj                *otg.MacsecCryptoEngineStatelessEncryptionOnlyTxPn
	marshaller         marshalMacsecCryptoEngineStatelessEncryptionOnlyTxPn
	unMarshaller       unMarshalMacsecCryptoEngineStatelessEncryptionOnlyTxPn
	fixedHolder        MacsecCryptoEngineStatelessEncryptionOnlyFixedPn
	incrementingHolder MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
}

func NewMacsecCryptoEngineStatelessEncryptionOnlyTxPn() MacsecCryptoEngineStatelessEncryptionOnlyTxPn {
	obj := macsecCryptoEngineStatelessEncryptionOnlyTxPn{obj: &otg.MacsecCryptoEngineStatelessEncryptionOnlyTxPn{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) msg() *otg.MacsecCryptoEngineStatelessEncryptionOnlyTxPn {
	return obj.obj
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) setMsg(msg *otg.MacsecCryptoEngineStatelessEncryptionOnlyTxPn) MacsecCryptoEngineStatelessEncryptionOnlyTxPn {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineStatelessEncryptionOnlyTxPn struct {
	obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn
}

type marshalMacsecCryptoEngineStatelessEncryptionOnlyTxPn interface {
	// ToProto marshals MacsecCryptoEngineStatelessEncryptionOnlyTxPn to protobuf object *otg.MacsecCryptoEngineStatelessEncryptionOnlyTxPn
	ToProto() (*otg.MacsecCryptoEngineStatelessEncryptionOnlyTxPn, error)
	// ToPbText marshals MacsecCryptoEngineStatelessEncryptionOnlyTxPn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineStatelessEncryptionOnlyTxPn to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineStatelessEncryptionOnlyTxPn to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineStatelessEncryptionOnlyTxPn struct {
	obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn
}

type unMarshalMacsecCryptoEngineStatelessEncryptionOnlyTxPn interface {
	// FromProto unmarshals MacsecCryptoEngineStatelessEncryptionOnlyTxPn from protobuf object *otg.MacsecCryptoEngineStatelessEncryptionOnlyTxPn
	FromProto(msg *otg.MacsecCryptoEngineStatelessEncryptionOnlyTxPn) (MacsecCryptoEngineStatelessEncryptionOnlyTxPn, error)
	// FromPbText unmarshals MacsecCryptoEngineStatelessEncryptionOnlyTxPn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineStatelessEncryptionOnlyTxPn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineStatelessEncryptionOnlyTxPn from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) Marshal() marshalMacsecCryptoEngineStatelessEncryptionOnlyTxPn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineStatelessEncryptionOnlyTxPn{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) Unmarshal() unMarshalMacsecCryptoEngineStatelessEncryptionOnlyTxPn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineStatelessEncryptionOnlyTxPn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineStatelessEncryptionOnlyTxPn) ToProto() (*otg.MacsecCryptoEngineStatelessEncryptionOnlyTxPn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineStatelessEncryptionOnlyTxPn) FromProto(msg *otg.MacsecCryptoEngineStatelessEncryptionOnlyTxPn) (MacsecCryptoEngineStatelessEncryptionOnlyTxPn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineStatelessEncryptionOnlyTxPn) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatelessEncryptionOnlyTxPn) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineStatelessEncryptionOnlyTxPn) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatelessEncryptionOnlyTxPn) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineStatelessEncryptionOnlyTxPn) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatelessEncryptionOnlyTxPn) FromJson(value string) error {
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

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) Clone() (MacsecCryptoEngineStatelessEncryptionOnlyTxPn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineStatelessEncryptionOnlyTxPn()
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

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) setNil() {
	obj.fixedHolder = nil
	obj.incrementingHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecCryptoEngineStatelessEncryptionOnlyTxPn is tx packet number(PN) configuration.
type MacsecCryptoEngineStatelessEncryptionOnlyTxPn interface {
	Validation
	// msg marshals MacsecCryptoEngineStatelessEncryptionOnlyTxPn to protobuf object *otg.MacsecCryptoEngineStatelessEncryptionOnlyTxPn
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineStatelessEncryptionOnlyTxPn
	// setMsg unmarshals MacsecCryptoEngineStatelessEncryptionOnlyTxPn from protobuf object *otg.MacsecCryptoEngineStatelessEncryptionOnlyTxPn
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineStatelessEncryptionOnlyTxPn) MacsecCryptoEngineStatelessEncryptionOnlyTxPn
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineStatelessEncryptionOnlyTxPn
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineStatelessEncryptionOnlyTxPn
	// validate validates MacsecCryptoEngineStatelessEncryptionOnlyTxPn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineStatelessEncryptionOnlyTxPn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MacsecCryptoEngineStatelessEncryptionOnlyTxPnChoiceEnum, set in MacsecCryptoEngineStatelessEncryptionOnlyTxPn
	Choice() MacsecCryptoEngineStatelessEncryptionOnlyTxPnChoiceEnum
	// setChoice assigns MacsecCryptoEngineStatelessEncryptionOnlyTxPnChoiceEnum provided by user to MacsecCryptoEngineStatelessEncryptionOnlyTxPn
	setChoice(value MacsecCryptoEngineStatelessEncryptionOnlyTxPnChoiceEnum) MacsecCryptoEngineStatelessEncryptionOnlyTxPn
	// HasChoice checks if Choice has been set in MacsecCryptoEngineStatelessEncryptionOnlyTxPn
	HasChoice() bool
	// getter for FixedPn to set choice.
	FixedPn()
	// getter for IncrementingPn to set choice.
	IncrementingPn()
	// Fixed returns MacsecCryptoEngineStatelessEncryptionOnlyFixedPn, set in MacsecCryptoEngineStatelessEncryptionOnlyTxPn.
	// MacsecCryptoEngineStatelessEncryptionOnlyFixedPn is fixed packet number(PN) configuration.
	Fixed() MacsecCryptoEngineStatelessEncryptionOnlyFixedPn
	// SetFixed assigns MacsecCryptoEngineStatelessEncryptionOnlyFixedPn provided by user to MacsecCryptoEngineStatelessEncryptionOnlyTxPn.
	// MacsecCryptoEngineStatelessEncryptionOnlyFixedPn is fixed packet number(PN) configuration.
	SetFixed(value MacsecCryptoEngineStatelessEncryptionOnlyFixedPn) MacsecCryptoEngineStatelessEncryptionOnlyTxPn
	// HasFixed checks if Fixed has been set in MacsecCryptoEngineStatelessEncryptionOnlyTxPn
	HasFixed() bool
	// Incrementing returns MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn, set in MacsecCryptoEngineStatelessEncryptionOnlyTxPn.
	// MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn is incrementing packet number(PN) configuration.
	Incrementing() MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
	// SetIncrementing assigns MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn provided by user to MacsecCryptoEngineStatelessEncryptionOnlyTxPn.
	// MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn is incrementing packet number(PN) configuration.
	SetIncrementing(value MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) MacsecCryptoEngineStatelessEncryptionOnlyTxPn
	// HasIncrementing checks if Incrementing has been set in MacsecCryptoEngineStatelessEncryptionOnlyTxPn
	HasIncrementing() bool
	setNil()
}

type MacsecCryptoEngineStatelessEncryptionOnlyTxPnChoiceEnum string

// Enum of Choice on MacsecCryptoEngineStatelessEncryptionOnlyTxPn
var MacsecCryptoEngineStatelessEncryptionOnlyTxPnChoice = struct {
	FIXED_PN        MacsecCryptoEngineStatelessEncryptionOnlyTxPnChoiceEnum
	INCREMENTING_PN MacsecCryptoEngineStatelessEncryptionOnlyTxPnChoiceEnum
}{
	FIXED_PN:        MacsecCryptoEngineStatelessEncryptionOnlyTxPnChoiceEnum("fixed_pn"),
	INCREMENTING_PN: MacsecCryptoEngineStatelessEncryptionOnlyTxPnChoiceEnum("incrementing_pn"),
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) Choice() MacsecCryptoEngineStatelessEncryptionOnlyTxPnChoiceEnum {
	return MacsecCryptoEngineStatelessEncryptionOnlyTxPnChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for FixedPn to set choice
func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) FixedPn() {
	obj.setChoice(MacsecCryptoEngineStatelessEncryptionOnlyTxPnChoice.FIXED_PN)
}

// getter for IncrementingPn to set choice
func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) IncrementingPn() {
	obj.setChoice(MacsecCryptoEngineStatelessEncryptionOnlyTxPnChoice.INCREMENTING_PN)
}

// Types of Tx packet number(PN) series. Supported choices: 1) fixed PN - MACsec packets will be sent out with the configured fixed PN or lower half of configured fixed XPN. 2) incrementing PN - MACsec packets will be sent out by single device with an incrementing PN or XPN.
// Choice returns a string
func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) setChoice(value MacsecCryptoEngineStatelessEncryptionOnlyTxPnChoiceEnum) MacsecCryptoEngineStatelessEncryptionOnlyTxPn {
	intValue, ok := otg.MacsecCryptoEngineStatelessEncryptionOnlyTxPn_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecCryptoEngineStatelessEncryptionOnlyTxPnChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecCryptoEngineStatelessEncryptionOnlyTxPn_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

// description is TBD
// Fixed returns a MacsecCryptoEngineStatelessEncryptionOnlyFixedPn
func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) Fixed() MacsecCryptoEngineStatelessEncryptionOnlyFixedPn {
	if obj.obj.Fixed == nil {
		obj.obj.Fixed = NewMacsecCryptoEngineStatelessEncryptionOnlyFixedPn().msg()
	}
	if obj.fixedHolder == nil {
		obj.fixedHolder = &macsecCryptoEngineStatelessEncryptionOnlyFixedPn{obj: obj.obj.Fixed}
	}
	return obj.fixedHolder
}

// description is TBD
// Fixed returns a MacsecCryptoEngineStatelessEncryptionOnlyFixedPn
func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) HasFixed() bool {
	return obj.obj.Fixed != nil
}

// description is TBD
// SetFixed sets the MacsecCryptoEngineStatelessEncryptionOnlyFixedPn value in the MacsecCryptoEngineStatelessEncryptionOnlyTxPn object
func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) SetFixed(value MacsecCryptoEngineStatelessEncryptionOnlyFixedPn) MacsecCryptoEngineStatelessEncryptionOnlyTxPn {

	obj.fixedHolder = nil
	obj.obj.Fixed = value.msg()

	return obj
}

// description is TBD
// Incrementing returns a MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) Incrementing() MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn {
	if obj.obj.Incrementing == nil {
		obj.obj.Incrementing = NewMacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn().msg()
	}
	if obj.incrementingHolder == nil {
		obj.incrementingHolder = &macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn{obj: obj.obj.Incrementing}
	}
	return obj.incrementingHolder
}

// description is TBD
// Incrementing returns a MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) HasIncrementing() bool {
	return obj.obj.Incrementing != nil
}

// description is TBD
// SetIncrementing sets the MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn value in the MacsecCryptoEngineStatelessEncryptionOnlyTxPn object
func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) SetIncrementing(value MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) MacsecCryptoEngineStatelessEncryptionOnlyTxPn {

	obj.incrementingHolder = nil
	obj.obj.Incrementing = value.msg()

	return obj
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) validateObj(vObj *validation, set_default bool) {
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

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTxPn) setDefault() {
	var choices_set int = 0
	var choice MacsecCryptoEngineStatelessEncryptionOnlyTxPnChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MacsecCryptoEngineStatelessEncryptionOnlyTxPnChoice.FIXED_PN)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MacsecCryptoEngineStatelessEncryptionOnlyTxPn")
			}
		} else {
			intVal := otg.MacsecCryptoEngineStatelessEncryptionOnlyTxPn_Choice_Enum_value[string(choice)]
			enumValue := otg.MacsecCryptoEngineStatelessEncryptionOnlyTxPn_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

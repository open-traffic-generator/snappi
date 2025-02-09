package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn *****
type macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn struct {
	validation
	obj                *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
	marshaller         marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
	unMarshaller       unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
	fixedHolder        MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn
	incrementingHolder MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
}

func NewMacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn() MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn {
	obj := macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn{obj: &otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) msg() *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn {
	return obj.obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) setMsg(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn struct {
	obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
}

type marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn interface {
	// ToProto marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn to protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
	ToProto() (*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn, error)
	// ToPbText marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn struct {
	obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
}

type unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn interface {
	// FromProto unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn from protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
	FromProto(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) (MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn, error)
	// FromPbText unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) Marshal() marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) Unmarshal() unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) ToProto() (*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) FromProto(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) (MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) FromJson(value string) error {
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

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) Clone() (MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn()
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

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) setNil() {
	obj.fixedHolder = nil
	obj.incrementingHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn is tx XPN settings.
type MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn interface {
	Validation
	// msg marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn to protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
	// setMsg unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn from protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
	// validate validates MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpnChoiceEnum, set in MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
	Choice() MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpnChoiceEnum
	// setChoice assigns MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpnChoiceEnum provided by user to MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
	setChoice(value MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpnChoiceEnum) MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
	// HasChoice checks if Choice has been set in MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
	HasChoice() bool
	// getter for FixedPn to set choice.
	FixedPn()
	// getter for IncrementingPn to set choice.
	IncrementingPn()
	// Fixed returns MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn, set in MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn.
	// MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn is fixed XPN settings.
	Fixed() MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn
	// SetFixed assigns MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn provided by user to MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn.
	// MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn is fixed XPN settings.
	SetFixed(value MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
	// HasFixed checks if Fixed has been set in MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
	HasFixed() bool
	// Incrementing returns MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn, set in MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn.
	// MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn is incrementing XPN settings.
	Incrementing() MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
	// SetIncrementing assigns MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn provided by user to MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn.
	// MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn is incrementing XPN settings.
	SetIncrementing(value MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
	// HasIncrementing checks if Incrementing has been set in MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
	HasIncrementing() bool
	setNil()
}

type MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpnChoiceEnum string

// Enum of Choice on MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn
var MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpnChoice = struct {
	FIXED_PN        MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpnChoiceEnum
	INCREMENTING_PN MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpnChoiceEnum
}{
	FIXED_PN:        MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpnChoiceEnum("fixed_pn"),
	INCREMENTING_PN: MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpnChoiceEnum("incrementing_pn"),
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) Choice() MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpnChoiceEnum {
	return MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpnChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for FixedPn to set choice
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) FixedPn() {
	obj.setChoice(MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpnChoice.FIXED_PN)
}

// getter for IncrementingPn to set choice
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) IncrementingPn() {
	obj.setChoice(MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpnChoice.INCREMENTING_PN)
}

// Types of Tx XPN series.
// Choice returns a string
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) setChoice(value MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpnChoiceEnum) MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn {
	intValue, ok := otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpnChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

// description is TBD
// Fixed returns a MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) Fixed() MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn {
	if obj.obj.Fixed == nil {
		obj.obj.Fixed = NewMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn().msg()
	}
	if obj.fixedHolder == nil {
		obj.fixedHolder = &macsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn{obj: obj.obj.Fixed}
	}
	return obj.fixedHolder
}

// description is TBD
// Fixed returns a MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) HasFixed() bool {
	return obj.obj.Fixed != nil
}

// description is TBD
// SetFixed sets the MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn value in the MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn object
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) SetFixed(value MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn {

	obj.fixedHolder = nil
	obj.obj.Fixed = value.msg()

	return obj
}

// description is TBD
// Incrementing returns a MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) Incrementing() MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn {
	if obj.obj.Incrementing == nil {
		obj.obj.Incrementing = NewMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn().msg()
	}
	if obj.incrementingHolder == nil {
		obj.incrementingHolder = &macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn{obj: obj.obj.Incrementing}
	}
	return obj.incrementingHolder
}

// description is TBD
// Incrementing returns a MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) HasIncrementing() bool {
	return obj.obj.Incrementing != nil
}

// description is TBD
// SetIncrementing sets the MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn value in the MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn object
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) SetIncrementing(value MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn {

	obj.incrementingHolder = nil
	obj.obj.Incrementing = value.msg()

	return obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) validateObj(vObj *validation, set_default bool) {
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

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn) setDefault() {
	var choices_set int = 0
	var choice MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpnChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpnChoice.FIXED_PN)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn")
			}
		} else {
			intVal := otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn_Choice_Enum_value[string(choice)]
			enumValue := otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxXpn_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

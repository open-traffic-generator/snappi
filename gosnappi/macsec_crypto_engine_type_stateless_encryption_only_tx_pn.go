package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn *****
type macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn struct {
	validation
	obj                *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	marshaller         marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	unMarshaller       unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	fixedHolder        MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	incrementingHolder MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
}

func NewMacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn() MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn {
	obj := macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn{obj: &otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) msg() *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn {
	return obj.obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) setMsg(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn struct {
	obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
}

type marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn interface {
	// ToProto marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn to protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	ToProto() (*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn, error)
	// ToPbText marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn struct {
	obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
}

type unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn interface {
	// FromProto unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn from protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	FromProto(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) (MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn, error)
	// FromPbText unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) Marshal() marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) Unmarshal() unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) ToProto() (*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) FromProto(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) (MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) FromJson(value string) error {
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

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) Clone() (MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn()
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

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) setNil() {
	obj.fixedHolder = nil
	obj.incrementingHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn is tx PN settings.
type MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn interface {
	Validation
	// msg marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn to protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// setMsg unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn from protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// validate validates MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum, set in MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	Choice() MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum
	// setChoice assigns MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum provided by user to MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	setChoice(value MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum) MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// HasChoice checks if Choice has been set in MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	HasChoice() bool
	// getter for IncrementingPn to set choice.
	IncrementingPn()
	// getter for FixedPn to set choice.
	FixedPn()
	// Fixed returns MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn, set in MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn.
	// MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn is fixed PN settings.
	Fixed() MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	// SetFixed assigns MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn provided by user to MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn.
	// MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn is fixed PN settings.
	SetFixed(value MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// HasFixed checks if Fixed has been set in MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	HasFixed() bool
	// Incrementing returns MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn, set in MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn.
	// MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn is incrementing PN settings.
	Incrementing() MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	// SetIncrementing assigns MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn provided by user to MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn.
	// MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn is incrementing PN settings.
	SetIncrementing(value MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// HasIncrementing checks if Incrementing has been set in MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	HasIncrementing() bool
	setNil()
}

type MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum string

// Enum of Choice on MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
var MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPnChoice = struct {
	FIXED_PN        MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum
	INCREMENTING_PN MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum
}{
	FIXED_PN:        MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum("fixed_pn"),
	INCREMENTING_PN: MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum("incrementing_pn"),
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) Choice() MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum {
	return MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for IncrementingPn to set choice
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) IncrementingPn() {
	obj.setChoice(MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPnChoice.INCREMENTING_PN)
}

// getter for FixedPn to set choice
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) FixedPn() {
	obj.setChoice(MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPnChoice.FIXED_PN)
}

// Types of Tx PN series.
// Choice returns a string
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) setChoice(value MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum) MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn {
	intValue, ok := otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

// description is TBD
// Fixed returns a MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) Fixed() MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn {
	if obj.obj.Fixed == nil {
		obj.obj.Fixed = NewMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn().msg()
	}
	if obj.fixedHolder == nil {
		obj.fixedHolder = &macsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn{obj: obj.obj.Fixed}
	}
	return obj.fixedHolder
}

// description is TBD
// Fixed returns a MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) HasFixed() bool {
	return obj.obj.Fixed != nil
}

// description is TBD
// SetFixed sets the MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn value in the MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn object
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) SetFixed(value MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn {

	obj.fixedHolder = nil
	obj.obj.Fixed = value.msg()

	return obj
}

// description is TBD
// Incrementing returns a MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) Incrementing() MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn {
	if obj.obj.Incrementing == nil {
		obj.obj.Incrementing = NewMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn().msg()
	}
	if obj.incrementingHolder == nil {
		obj.incrementingHolder = &macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn{obj: obj.obj.Incrementing}
	}
	return obj.incrementingHolder
}

// description is TBD
// Incrementing returns a MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) HasIncrementing() bool {
	return obj.obj.Incrementing != nil
}

// description is TBD
// SetIncrementing sets the MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn value in the MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn object
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) SetIncrementing(value MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn {

	obj.incrementingHolder = nil
	obj.obj.Incrementing = value.msg()

	return obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) validateObj(vObj *validation, set_default bool) {
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

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) setDefault() {
	var choices_set int = 0
	var choice MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPnChoice.FIXED_PN)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn")
			}
		} else {
			intVal := otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn_Choice_Enum_value[string(choice)]
			enumValue := otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn *****
type macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn struct {
	validation
	obj                *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	marshaller         marshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	unMarshaller       unMarshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	fixedHolder        MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	incrementingHolder MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
}

func NewMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn() MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn {
	obj := macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn{obj: &otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) msg() *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn {
	return obj.obj
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) setMsg(msg *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn struct {
	obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
}

type marshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn interface {
	// ToProto marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn to protobuf object *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	ToProto() (*otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn, error)
	// ToPbText marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn struct {
	obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
}

type unMarshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn interface {
	// FromProto unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn from protobuf object *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	FromProto(msg *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) (MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn, error)
	// FromPbText unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) Marshal() marshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) Unmarshal() unMarshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) ToProto() (*otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) FromProto(msg *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) (MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) FromPbText(value string) error {
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

func (m *marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) FromYaml(value string) error {
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

func (m *marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) FromJson(value string) error {
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

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) Clone() (MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn()
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

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) setNil() {
	obj.fixedHolder = nil
	obj.incrementingHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn is tx PN settings.
type MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn interface {
	Validation
	// msg marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn to protobuf object *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// and doesn't set defaults
	msg() *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// setMsg unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn from protobuf object *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// provides marshal interface
	Marshal() marshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// validate validates MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum, set in MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	Choice() MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum
	// setChoice assigns MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum provided by user to MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	setChoice(value MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// HasChoice checks if Choice has been set in MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	HasChoice() bool
	// getter for FixedPn to set choice.
	FixedPn()
	// getter for IncrementingPn to set choice.
	IncrementingPn()
	// Fixed returns MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn, set in MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn.
	// MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn is fixed PN settings.
	Fixed() MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	// SetFixed assigns MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn provided by user to MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn.
	// MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn is fixed PN settings.
	SetFixed(value MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// HasFixed checks if Fixed has been set in MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	HasFixed() bool
	// Incrementing returns MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn, set in MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn.
	// MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn is incrementing PN settings.
	Incrementing() MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	// SetIncrementing assigns MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn provided by user to MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn.
	// MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn is incrementing PN settings.
	SetIncrementing(value MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// HasIncrementing checks if Incrementing has been set in MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	HasIncrementing() bool
	setNil()
}

type MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum string

// Enum of Choice on MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
var MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPnChoice = struct {
	FIXED_PN        MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum
	INCREMENTING_PN MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum
}{
	FIXED_PN:        MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum("fixed_pn"),
	INCREMENTING_PN: MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum("incrementing_pn"),
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) Choice() MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum {
	return MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for FixedPn to set choice
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) FixedPn() {
	obj.setChoice(MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPnChoice.FIXED_PN)
}

// getter for IncrementingPn to set choice
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) IncrementingPn() {
	obj.setChoice(MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPnChoice.INCREMENTING_PN)
}

// Types of Tx PN series.
// Choice returns a string
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) setChoice(value MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn {
	intValue, ok := otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

// description is TBD
// Fixed returns a MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) Fixed() MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn {
	if obj.obj.Fixed == nil {
		obj.obj.Fixed = NewMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn().msg()
	}
	if obj.fixedHolder == nil {
		obj.fixedHolder = &macsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn{obj: obj.obj.Fixed}
	}
	return obj.fixedHolder
}

// description is TBD
// Fixed returns a MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) HasFixed() bool {
	return obj.obj.Fixed != nil
}

// description is TBD
// SetFixed sets the MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn value in the MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn object
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) SetFixed(value MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn {

	obj.fixedHolder = nil
	obj.obj.Fixed = value.msg()

	return obj
}

// description is TBD
// Incrementing returns a MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) Incrementing() MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn {
	if obj.obj.Incrementing == nil {
		obj.obj.Incrementing = NewMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn().msg()
	}
	if obj.incrementingHolder == nil {
		obj.incrementingHolder = &macsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn{obj: obj.obj.Incrementing}
	}
	return obj.incrementingHolder
}

// description is TBD
// Incrementing returns a MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) HasIncrementing() bool {
	return obj.obj.Incrementing != nil
}

// description is TBD
// SetIncrementing sets the MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn value in the MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn object
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) SetIncrementing(value MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn {

	obj.incrementingHolder = nil
	obj.obj.Incrementing = value.msg()

	return obj
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) validateObj(vObj *validation, set_default bool) {
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

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) setDefault() {
	var choices_set int = 0
	var choice MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPnChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPnChoice.FIXED_PN)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn")
			}
		} else {
			intVal := otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn_Choice_Enum_value[string(choice)]
			enumValue := otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

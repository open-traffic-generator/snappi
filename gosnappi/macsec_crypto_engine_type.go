package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineType *****
type macsecCryptoEngineType struct {
	validation
	obj                                *otg.MacsecCryptoEngineType
	marshaller                         marshalMacsecCryptoEngineType
	unMarshaller                       unMarshalMacsecCryptoEngineType
	statelessEncryptionOnlyHolder      MacsecCryptoEngineTypeStatelessEncryptionOnly
	statefulEncryptionDecryptionHolder MacsecCryptoEngineTypeStatefulEncryptionDecryption
}

func NewMacsecCryptoEngineType() MacsecCryptoEngineType {
	obj := macsecCryptoEngineType{obj: &otg.MacsecCryptoEngineType{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineType) msg() *otg.MacsecCryptoEngineType {
	return obj.obj
}

func (obj *macsecCryptoEngineType) setMsg(msg *otg.MacsecCryptoEngineType) MacsecCryptoEngineType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineType struct {
	obj *macsecCryptoEngineType
}

type marshalMacsecCryptoEngineType interface {
	// ToProto marshals MacsecCryptoEngineType to protobuf object *otg.MacsecCryptoEngineType
	ToProto() (*otg.MacsecCryptoEngineType, error)
	// ToPbText marshals MacsecCryptoEngineType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineType to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineType to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineType struct {
	obj *macsecCryptoEngineType
}

type unMarshalMacsecCryptoEngineType interface {
	// FromProto unmarshals MacsecCryptoEngineType from protobuf object *otg.MacsecCryptoEngineType
	FromProto(msg *otg.MacsecCryptoEngineType) (MacsecCryptoEngineType, error)
	// FromPbText unmarshals MacsecCryptoEngineType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineType from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineType) Marshal() marshalMacsecCryptoEngineType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineType{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineType) Unmarshal() unMarshalMacsecCryptoEngineType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineType) ToProto() (*otg.MacsecCryptoEngineType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineType) FromProto(msg *otg.MacsecCryptoEngineType) (MacsecCryptoEngineType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineType) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineType) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineType) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineType) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineType) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineType) FromJson(value string) error {
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

func (obj *macsecCryptoEngineType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineType) Clone() (MacsecCryptoEngineType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineType()
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

func (obj *macsecCryptoEngineType) setNil() {
	obj.statelessEncryptionOnlyHolder = nil
	obj.statefulEncryptionDecryptionHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecCryptoEngineType is crypto engine type.
type MacsecCryptoEngineType interface {
	Validation
	// msg marshals MacsecCryptoEngineType to protobuf object *otg.MacsecCryptoEngineType
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineType
	// setMsg unmarshals MacsecCryptoEngineType from protobuf object *otg.MacsecCryptoEngineType
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineType) MacsecCryptoEngineType
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineType
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineType
	// validate validates MacsecCryptoEngineType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MacsecCryptoEngineTypeChoiceEnum, set in MacsecCryptoEngineType
	Choice() MacsecCryptoEngineTypeChoiceEnum
	// setChoice assigns MacsecCryptoEngineTypeChoiceEnum provided by user to MacsecCryptoEngineType
	setChoice(value MacsecCryptoEngineTypeChoiceEnum) MacsecCryptoEngineType
	// HasChoice checks if Choice has been set in MacsecCryptoEngineType
	HasChoice() bool
	// StatelessEncryptionOnly returns MacsecCryptoEngineTypeStatelessEncryptionOnly, set in MacsecCryptoEngineType.
	// MacsecCryptoEngineTypeStatelessEncryptionOnly is the container for stateless encryption only engine settings.
	StatelessEncryptionOnly() MacsecCryptoEngineTypeStatelessEncryptionOnly
	// SetStatelessEncryptionOnly assigns MacsecCryptoEngineTypeStatelessEncryptionOnly provided by user to MacsecCryptoEngineType.
	// MacsecCryptoEngineTypeStatelessEncryptionOnly is the container for stateless encryption only engine settings.
	SetStatelessEncryptionOnly(value MacsecCryptoEngineTypeStatelessEncryptionOnly) MacsecCryptoEngineType
	// HasStatelessEncryptionOnly checks if StatelessEncryptionOnly has been set in MacsecCryptoEngineType
	HasStatelessEncryptionOnly() bool
	// StatefulEncryptionDecryption returns MacsecCryptoEngineTypeStatefulEncryptionDecryption, set in MacsecCryptoEngineType.
	// MacsecCryptoEngineTypeStatefulEncryptionDecryption is the container for stateful encryption and decryption engine settings.
	StatefulEncryptionDecryption() MacsecCryptoEngineTypeStatefulEncryptionDecryption
	// SetStatefulEncryptionDecryption assigns MacsecCryptoEngineTypeStatefulEncryptionDecryption provided by user to MacsecCryptoEngineType.
	// MacsecCryptoEngineTypeStatefulEncryptionDecryption is the container for stateful encryption and decryption engine settings.
	SetStatefulEncryptionDecryption(value MacsecCryptoEngineTypeStatefulEncryptionDecryption) MacsecCryptoEngineType
	// HasStatefulEncryptionDecryption checks if StatefulEncryptionDecryption has been set in MacsecCryptoEngineType
	HasStatefulEncryptionDecryption() bool
	setNil()
}

type MacsecCryptoEngineTypeChoiceEnum string

// Enum of Choice on MacsecCryptoEngineType
var MacsecCryptoEngineTypeChoice = struct {
	STATELESS_ENCRYPTION_ONLY      MacsecCryptoEngineTypeChoiceEnum
	STATEFUL_ENCRYPTION_DECRYPTION MacsecCryptoEngineTypeChoiceEnum
}{
	STATELESS_ENCRYPTION_ONLY:      MacsecCryptoEngineTypeChoiceEnum("stateless_encryption_only"),
	STATEFUL_ENCRYPTION_DECRYPTION: MacsecCryptoEngineTypeChoiceEnum("stateful_encryption_decryption"),
}

func (obj *macsecCryptoEngineType) Choice() MacsecCryptoEngineTypeChoiceEnum {
	return MacsecCryptoEngineTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// Engine type based on encryption and/ or decryption capability..
// Choice returns a string
func (obj *macsecCryptoEngineType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *macsecCryptoEngineType) setChoice(value MacsecCryptoEngineTypeChoiceEnum) MacsecCryptoEngineType {
	intValue, ok := otg.MacsecCryptoEngineType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecCryptoEngineTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecCryptoEngineType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.StatefulEncryptionDecryption = nil
	obj.statefulEncryptionDecryptionHolder = nil
	obj.obj.StatelessEncryptionOnly = nil
	obj.statelessEncryptionOnlyHolder = nil

	if value == MacsecCryptoEngineTypeChoice.STATELESS_ENCRYPTION_ONLY {
		obj.obj.StatelessEncryptionOnly = NewMacsecCryptoEngineTypeStatelessEncryptionOnly().msg()
	}

	if value == MacsecCryptoEngineTypeChoice.STATEFUL_ENCRYPTION_DECRYPTION {
		obj.obj.StatefulEncryptionDecryption = NewMacsecCryptoEngineTypeStatefulEncryptionDecryption().msg()
	}

	return obj
}

// description is TBD
// StatelessEncryptionOnly returns a MacsecCryptoEngineTypeStatelessEncryptionOnly
func (obj *macsecCryptoEngineType) StatelessEncryptionOnly() MacsecCryptoEngineTypeStatelessEncryptionOnly {
	if obj.obj.StatelessEncryptionOnly == nil {
		obj.setChoice(MacsecCryptoEngineTypeChoice.STATELESS_ENCRYPTION_ONLY)
	}
	if obj.statelessEncryptionOnlyHolder == nil {
		obj.statelessEncryptionOnlyHolder = &macsecCryptoEngineTypeStatelessEncryptionOnly{obj: obj.obj.StatelessEncryptionOnly}
	}
	return obj.statelessEncryptionOnlyHolder
}

// description is TBD
// StatelessEncryptionOnly returns a MacsecCryptoEngineTypeStatelessEncryptionOnly
func (obj *macsecCryptoEngineType) HasStatelessEncryptionOnly() bool {
	return obj.obj.StatelessEncryptionOnly != nil
}

// description is TBD
// SetStatelessEncryptionOnly sets the MacsecCryptoEngineTypeStatelessEncryptionOnly value in the MacsecCryptoEngineType object
func (obj *macsecCryptoEngineType) SetStatelessEncryptionOnly(value MacsecCryptoEngineTypeStatelessEncryptionOnly) MacsecCryptoEngineType {
	obj.setChoice(MacsecCryptoEngineTypeChoice.STATELESS_ENCRYPTION_ONLY)
	obj.statelessEncryptionOnlyHolder = nil
	obj.obj.StatelessEncryptionOnly = value.msg()

	return obj
}

// description is TBD
// StatefulEncryptionDecryption returns a MacsecCryptoEngineTypeStatefulEncryptionDecryption
func (obj *macsecCryptoEngineType) StatefulEncryptionDecryption() MacsecCryptoEngineTypeStatefulEncryptionDecryption {
	if obj.obj.StatefulEncryptionDecryption == nil {
		obj.setChoice(MacsecCryptoEngineTypeChoice.STATEFUL_ENCRYPTION_DECRYPTION)
	}
	if obj.statefulEncryptionDecryptionHolder == nil {
		obj.statefulEncryptionDecryptionHolder = &macsecCryptoEngineTypeStatefulEncryptionDecryption{obj: obj.obj.StatefulEncryptionDecryption}
	}
	return obj.statefulEncryptionDecryptionHolder
}

// description is TBD
// StatefulEncryptionDecryption returns a MacsecCryptoEngineTypeStatefulEncryptionDecryption
func (obj *macsecCryptoEngineType) HasStatefulEncryptionDecryption() bool {
	return obj.obj.StatefulEncryptionDecryption != nil
}

// description is TBD
// SetStatefulEncryptionDecryption sets the MacsecCryptoEngineTypeStatefulEncryptionDecryption value in the MacsecCryptoEngineType object
func (obj *macsecCryptoEngineType) SetStatefulEncryptionDecryption(value MacsecCryptoEngineTypeStatefulEncryptionDecryption) MacsecCryptoEngineType {
	obj.setChoice(MacsecCryptoEngineTypeChoice.STATEFUL_ENCRYPTION_DECRYPTION)
	obj.statefulEncryptionDecryptionHolder = nil
	obj.obj.StatefulEncryptionDecryption = value.msg()

	return obj
}

func (obj *macsecCryptoEngineType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.StatelessEncryptionOnly != nil {

		obj.StatelessEncryptionOnly().validateObj(vObj, set_default)
	}

	if obj.obj.StatefulEncryptionDecryption != nil {

		obj.StatefulEncryptionDecryption().validateObj(vObj, set_default)
	}

}

func (obj *macsecCryptoEngineType) setDefault() {
	var choices_set int = 0
	var choice MacsecCryptoEngineTypeChoiceEnum

	if obj.obj.StatelessEncryptionOnly != nil {
		choices_set += 1
		choice = MacsecCryptoEngineTypeChoice.STATELESS_ENCRYPTION_ONLY
	}

	if obj.obj.StatefulEncryptionDecryption != nil {
		choices_set += 1
		choice = MacsecCryptoEngineTypeChoice.STATEFUL_ENCRYPTION_DECRYPTION
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MacsecCryptoEngineTypeChoice.STATELESS_ENCRYPTION_ONLY)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MacsecCryptoEngineType")
			}
		} else {
			intVal := otg.MacsecCryptoEngineType_Choice_Enum_value[string(choice)]
			enumValue := otg.MacsecCryptoEngineType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

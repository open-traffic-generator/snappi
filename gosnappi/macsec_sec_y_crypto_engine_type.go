package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYCryptoEngineType *****
type macsecSecYCryptoEngineType struct {
	validation
	obj                                *otg.MacsecSecYCryptoEngineType
	marshaller                         marshalMacsecSecYCryptoEngineType
	unMarshaller                       unMarshalMacsecSecYCryptoEngineType
	statelessEncryptionOnlyHolder      MacsecSecYCryptoEngineTypeStatelessEncryptionOnly
	statefulEncryptionDecryptionHolder MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption
}

func NewMacsecSecYCryptoEngineType() MacsecSecYCryptoEngineType {
	obj := macsecSecYCryptoEngineType{obj: &otg.MacsecSecYCryptoEngineType{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYCryptoEngineType) msg() *otg.MacsecSecYCryptoEngineType {
	return obj.obj
}

func (obj *macsecSecYCryptoEngineType) setMsg(msg *otg.MacsecSecYCryptoEngineType) MacsecSecYCryptoEngineType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYCryptoEngineType struct {
	obj *macsecSecYCryptoEngineType
}

type marshalMacsecSecYCryptoEngineType interface {
	// ToProto marshals MacsecSecYCryptoEngineType to protobuf object *otg.MacsecSecYCryptoEngineType
	ToProto() (*otg.MacsecSecYCryptoEngineType, error)
	// ToPbText marshals MacsecSecYCryptoEngineType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYCryptoEngineType to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYCryptoEngineType to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYCryptoEngineType struct {
	obj *macsecSecYCryptoEngineType
}

type unMarshalMacsecSecYCryptoEngineType interface {
	// FromProto unmarshals MacsecSecYCryptoEngineType from protobuf object *otg.MacsecSecYCryptoEngineType
	FromProto(msg *otg.MacsecSecYCryptoEngineType) (MacsecSecYCryptoEngineType, error)
	// FromPbText unmarshals MacsecSecYCryptoEngineType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYCryptoEngineType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYCryptoEngineType from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYCryptoEngineType) Marshal() marshalMacsecSecYCryptoEngineType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYCryptoEngineType{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYCryptoEngineType) Unmarshal() unMarshalMacsecSecYCryptoEngineType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYCryptoEngineType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYCryptoEngineType) ToProto() (*otg.MacsecSecYCryptoEngineType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYCryptoEngineType) FromProto(msg *otg.MacsecSecYCryptoEngineType) (MacsecSecYCryptoEngineType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYCryptoEngineType) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineType) FromPbText(value string) error {
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

func (m *marshalmacsecSecYCryptoEngineType) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineType) FromYaml(value string) error {
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

func (m *marshalmacsecSecYCryptoEngineType) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineType) FromJson(value string) error {
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

func (obj *macsecSecYCryptoEngineType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngineType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngineType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYCryptoEngineType) Clone() (MacsecSecYCryptoEngineType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYCryptoEngineType()
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

func (obj *macsecSecYCryptoEngineType) setNil() {
	obj.statelessEncryptionOnlyHolder = nil
	obj.statefulEncryptionDecryptionHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecSecYCryptoEngineType is crypto engine type.
type MacsecSecYCryptoEngineType interface {
	Validation
	// msg marshals MacsecSecYCryptoEngineType to protobuf object *otg.MacsecSecYCryptoEngineType
	// and doesn't set defaults
	msg() *otg.MacsecSecYCryptoEngineType
	// setMsg unmarshals MacsecSecYCryptoEngineType from protobuf object *otg.MacsecSecYCryptoEngineType
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYCryptoEngineType) MacsecSecYCryptoEngineType
	// provides marshal interface
	Marshal() marshalMacsecSecYCryptoEngineType
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYCryptoEngineType
	// validate validates MacsecSecYCryptoEngineType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYCryptoEngineType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MacsecSecYCryptoEngineTypeChoiceEnum, set in MacsecSecYCryptoEngineType
	Choice() MacsecSecYCryptoEngineTypeChoiceEnum
	// setChoice assigns MacsecSecYCryptoEngineTypeChoiceEnum provided by user to MacsecSecYCryptoEngineType
	setChoice(value MacsecSecYCryptoEngineTypeChoiceEnum) MacsecSecYCryptoEngineType
	// HasChoice checks if Choice has been set in MacsecSecYCryptoEngineType
	HasChoice() bool
	// StatelessEncryptionOnly returns MacsecSecYCryptoEngineTypeStatelessEncryptionOnly, set in MacsecSecYCryptoEngineType.
	// MacsecSecYCryptoEngineTypeStatelessEncryptionOnly is the container for stateless encryption only engine settings.
	StatelessEncryptionOnly() MacsecSecYCryptoEngineTypeStatelessEncryptionOnly
	// SetStatelessEncryptionOnly assigns MacsecSecYCryptoEngineTypeStatelessEncryptionOnly provided by user to MacsecSecYCryptoEngineType.
	// MacsecSecYCryptoEngineTypeStatelessEncryptionOnly is the container for stateless encryption only engine settings.
	SetStatelessEncryptionOnly(value MacsecSecYCryptoEngineTypeStatelessEncryptionOnly) MacsecSecYCryptoEngineType
	// HasStatelessEncryptionOnly checks if StatelessEncryptionOnly has been set in MacsecSecYCryptoEngineType
	HasStatelessEncryptionOnly() bool
	// StatefulEncryptionDecryption returns MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption, set in MacsecSecYCryptoEngineType.
	// MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption is the container for stateful encryption and decryption engine settings.
	StatefulEncryptionDecryption() MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption
	// SetStatefulEncryptionDecryption assigns MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption provided by user to MacsecSecYCryptoEngineType.
	// MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption is the container for stateful encryption and decryption engine settings.
	SetStatefulEncryptionDecryption(value MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption) MacsecSecYCryptoEngineType
	// HasStatefulEncryptionDecryption checks if StatefulEncryptionDecryption has been set in MacsecSecYCryptoEngineType
	HasStatefulEncryptionDecryption() bool
	setNil()
}

type MacsecSecYCryptoEngineTypeChoiceEnum string

// Enum of Choice on MacsecSecYCryptoEngineType
var MacsecSecYCryptoEngineTypeChoice = struct {
	STATELESS_ENCRYPTION_ONLY      MacsecSecYCryptoEngineTypeChoiceEnum
	STATEFUL_ENCRYPTION_DECRYPTION MacsecSecYCryptoEngineTypeChoiceEnum
}{
	STATELESS_ENCRYPTION_ONLY:      MacsecSecYCryptoEngineTypeChoiceEnum("stateless_encryption_only"),
	STATEFUL_ENCRYPTION_DECRYPTION: MacsecSecYCryptoEngineTypeChoiceEnum("stateful_encryption_decryption"),
}

func (obj *macsecSecYCryptoEngineType) Choice() MacsecSecYCryptoEngineTypeChoiceEnum {
	return MacsecSecYCryptoEngineTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// Engine type based on encryption and/ or decryption capability..
// Choice returns a string
func (obj *macsecSecYCryptoEngineType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *macsecSecYCryptoEngineType) setChoice(value MacsecSecYCryptoEngineTypeChoiceEnum) MacsecSecYCryptoEngineType {
	intValue, ok := otg.MacsecSecYCryptoEngineType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecSecYCryptoEngineTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecSecYCryptoEngineType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.StatefulEncryptionDecryption = nil
	obj.statefulEncryptionDecryptionHolder = nil
	obj.obj.StatelessEncryptionOnly = nil
	obj.statelessEncryptionOnlyHolder = nil

	if value == MacsecSecYCryptoEngineTypeChoice.STATELESS_ENCRYPTION_ONLY {
		obj.obj.StatelessEncryptionOnly = NewMacsecSecYCryptoEngineTypeStatelessEncryptionOnly().msg()
	}

	if value == MacsecSecYCryptoEngineTypeChoice.STATEFUL_ENCRYPTION_DECRYPTION {
		obj.obj.StatefulEncryptionDecryption = NewMacsecSecYCryptoEngineTypeStatefulEncryptionDecryption().msg()
	}

	return obj
}

// description is TBD
// StatelessEncryptionOnly returns a MacsecSecYCryptoEngineTypeStatelessEncryptionOnly
func (obj *macsecSecYCryptoEngineType) StatelessEncryptionOnly() MacsecSecYCryptoEngineTypeStatelessEncryptionOnly {
	if obj.obj.StatelessEncryptionOnly == nil {
		obj.setChoice(MacsecSecYCryptoEngineTypeChoice.STATELESS_ENCRYPTION_ONLY)
	}
	if obj.statelessEncryptionOnlyHolder == nil {
		obj.statelessEncryptionOnlyHolder = &macsecSecYCryptoEngineTypeStatelessEncryptionOnly{obj: obj.obj.StatelessEncryptionOnly}
	}
	return obj.statelessEncryptionOnlyHolder
}

// description is TBD
// StatelessEncryptionOnly returns a MacsecSecYCryptoEngineTypeStatelessEncryptionOnly
func (obj *macsecSecYCryptoEngineType) HasStatelessEncryptionOnly() bool {
	return obj.obj.StatelessEncryptionOnly != nil
}

// description is TBD
// SetStatelessEncryptionOnly sets the MacsecSecYCryptoEngineTypeStatelessEncryptionOnly value in the MacsecSecYCryptoEngineType object
func (obj *macsecSecYCryptoEngineType) SetStatelessEncryptionOnly(value MacsecSecYCryptoEngineTypeStatelessEncryptionOnly) MacsecSecYCryptoEngineType {
	obj.setChoice(MacsecSecYCryptoEngineTypeChoice.STATELESS_ENCRYPTION_ONLY)
	obj.statelessEncryptionOnlyHolder = nil
	obj.obj.StatelessEncryptionOnly = value.msg()

	return obj
}

// description is TBD
// StatefulEncryptionDecryption returns a MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption
func (obj *macsecSecYCryptoEngineType) StatefulEncryptionDecryption() MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption {
	if obj.obj.StatefulEncryptionDecryption == nil {
		obj.setChoice(MacsecSecYCryptoEngineTypeChoice.STATEFUL_ENCRYPTION_DECRYPTION)
	}
	if obj.statefulEncryptionDecryptionHolder == nil {
		obj.statefulEncryptionDecryptionHolder = &macsecSecYCryptoEngineTypeStatefulEncryptionDecryption{obj: obj.obj.StatefulEncryptionDecryption}
	}
	return obj.statefulEncryptionDecryptionHolder
}

// description is TBD
// StatefulEncryptionDecryption returns a MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption
func (obj *macsecSecYCryptoEngineType) HasStatefulEncryptionDecryption() bool {
	return obj.obj.StatefulEncryptionDecryption != nil
}

// description is TBD
// SetStatefulEncryptionDecryption sets the MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption value in the MacsecSecYCryptoEngineType object
func (obj *macsecSecYCryptoEngineType) SetStatefulEncryptionDecryption(value MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption) MacsecSecYCryptoEngineType {
	obj.setChoice(MacsecSecYCryptoEngineTypeChoice.STATEFUL_ENCRYPTION_DECRYPTION)
	obj.statefulEncryptionDecryptionHolder = nil
	obj.obj.StatefulEncryptionDecryption = value.msg()

	return obj
}

func (obj *macsecSecYCryptoEngineType) validateObj(vObj *validation, set_default bool) {
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

func (obj *macsecSecYCryptoEngineType) setDefault() {
	var choices_set int = 0
	var choice MacsecSecYCryptoEngineTypeChoiceEnum

	if obj.obj.StatelessEncryptionOnly != nil {
		choices_set += 1
		choice = MacsecSecYCryptoEngineTypeChoice.STATELESS_ENCRYPTION_ONLY
	}

	if obj.obj.StatefulEncryptionDecryption != nil {
		choices_set += 1
		choice = MacsecSecYCryptoEngineTypeChoice.STATEFUL_ENCRYPTION_DECRYPTION
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MacsecSecYCryptoEngineTypeChoice.STATELESS_ENCRYPTION_ONLY)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MacsecSecYCryptoEngineType")
			}
		} else {
			intVal := otg.MacsecSecYCryptoEngineType_Choice_Enum_value[string(choice)]
			enumValue := otg.MacsecSecYCryptoEngineType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

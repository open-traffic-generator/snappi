package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngine *****
type macsecCryptoEngine struct {
	validation
	obj                                *otg.MacsecCryptoEngine
	marshaller                         marshalMacsecCryptoEngine
	unMarshaller                       unMarshalMacsecCryptoEngine
	statelessEncryptionOnlyHolder      MacsecCryptoEngineStatelessEncryptionOnly
	statefulEncryptionDecryptionHolder MacsecCryptoEngineStatefulEncryptionDecryption
}

func NewMacsecCryptoEngine() MacsecCryptoEngine {
	obj := macsecCryptoEngine{obj: &otg.MacsecCryptoEngine{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngine) msg() *otg.MacsecCryptoEngine {
	return obj.obj
}

func (obj *macsecCryptoEngine) setMsg(msg *otg.MacsecCryptoEngine) MacsecCryptoEngine {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngine struct {
	obj *macsecCryptoEngine
}

type marshalMacsecCryptoEngine interface {
	// ToProto marshals MacsecCryptoEngine to protobuf object *otg.MacsecCryptoEngine
	ToProto() (*otg.MacsecCryptoEngine, error)
	// ToPbText marshals MacsecCryptoEngine to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngine to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngine to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngine struct {
	obj *macsecCryptoEngine
}

type unMarshalMacsecCryptoEngine interface {
	// FromProto unmarshals MacsecCryptoEngine from protobuf object *otg.MacsecCryptoEngine
	FromProto(msg *otg.MacsecCryptoEngine) (MacsecCryptoEngine, error)
	// FromPbText unmarshals MacsecCryptoEngine from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngine from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngine from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngine) Marshal() marshalMacsecCryptoEngine {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngine{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngine) Unmarshal() unMarshalMacsecCryptoEngine {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngine{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngine) ToProto() (*otg.MacsecCryptoEngine, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngine) FromProto(msg *otg.MacsecCryptoEngine) (MacsecCryptoEngine, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngine) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngine) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngine) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngine) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngine) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngine) FromJson(value string) error {
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

func (obj *macsecCryptoEngine) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngine) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngine) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngine) Clone() (MacsecCryptoEngine, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngine()
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

func (obj *macsecCryptoEngine) setNil() {
	obj.statelessEncryptionOnlyHolder = nil
	obj.statefulEncryptionDecryptionHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecCryptoEngine is a container of crypto engine properties of a SecY.
type MacsecCryptoEngine interface {
	Validation
	// msg marshals MacsecCryptoEngine to protobuf object *otg.MacsecCryptoEngine
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngine
	// setMsg unmarshals MacsecCryptoEngine from protobuf object *otg.MacsecCryptoEngine
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngine) MacsecCryptoEngine
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngine
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngine
	// validate validates MacsecCryptoEngine
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngine, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MacsecCryptoEngineChoiceEnum, set in MacsecCryptoEngine
	Choice() MacsecCryptoEngineChoiceEnum
	// setChoice assigns MacsecCryptoEngineChoiceEnum provided by user to MacsecCryptoEngine
	setChoice(value MacsecCryptoEngineChoiceEnum) MacsecCryptoEngine
	// HasChoice checks if Choice has been set in MacsecCryptoEngine
	HasChoice() bool
	// StatelessEncryptionOnly returns MacsecCryptoEngineStatelessEncryptionOnly, set in MacsecCryptoEngine.
	// MacsecCryptoEngineStatelessEncryptionOnly is the container for stateless encryption only engine configuration.
	StatelessEncryptionOnly() MacsecCryptoEngineStatelessEncryptionOnly
	// SetStatelessEncryptionOnly assigns MacsecCryptoEngineStatelessEncryptionOnly provided by user to MacsecCryptoEngine.
	// MacsecCryptoEngineStatelessEncryptionOnly is the container for stateless encryption only engine configuration.
	SetStatelessEncryptionOnly(value MacsecCryptoEngineStatelessEncryptionOnly) MacsecCryptoEngine
	// HasStatelessEncryptionOnly checks if StatelessEncryptionOnly has been set in MacsecCryptoEngine
	HasStatelessEncryptionOnly() bool
	// StatefulEncryptionDecryption returns MacsecCryptoEngineStatefulEncryptionDecryption, set in MacsecCryptoEngine.
	// MacsecCryptoEngineStatefulEncryptionDecryption is the container for stateful encryption and decryption engine configuration.
	StatefulEncryptionDecryption() MacsecCryptoEngineStatefulEncryptionDecryption
	// SetStatefulEncryptionDecryption assigns MacsecCryptoEngineStatefulEncryptionDecryption provided by user to MacsecCryptoEngine.
	// MacsecCryptoEngineStatefulEncryptionDecryption is the container for stateful encryption and decryption engine configuration.
	SetStatefulEncryptionDecryption(value MacsecCryptoEngineStatefulEncryptionDecryption) MacsecCryptoEngine
	// HasStatefulEncryptionDecryption checks if StatefulEncryptionDecryption has been set in MacsecCryptoEngine
	HasStatefulEncryptionDecryption() bool
	setNil()
}

type MacsecCryptoEngineChoiceEnum string

// Enum of Choice on MacsecCryptoEngine
var MacsecCryptoEngineChoice = struct {
	STATELESS_ENCRYPTION_ONLY      MacsecCryptoEngineChoiceEnum
	STATEFUL_ENCRYPTION_DECRYPTION MacsecCryptoEngineChoiceEnum
}{
	STATELESS_ENCRYPTION_ONLY:      MacsecCryptoEngineChoiceEnum("stateless_encryption_only"),
	STATEFUL_ENCRYPTION_DECRYPTION: MacsecCryptoEngineChoiceEnum("stateful_encryption_decryption"),
}

func (obj *macsecCryptoEngine) Choice() MacsecCryptoEngineChoiceEnum {
	return MacsecCryptoEngineChoiceEnum(obj.obj.Choice.Enum().String())
}

// Engine type based on encryption and/ or decryption capability. Supported types: 1) stateless_encryption_only - engine can only encrypt transmitted packets but such engine cannot decrypt packets upon arrival. As the packets cannot be decrypted on arrival, such packets cannot be delivered to the receiving device. Hence only stateless traffic can be sent. 2) stateful_encryption_decryption - engine can both encrypt transmitted packets and decrypt packets on arrival. Such engine can have hardware acceleration for faster encryption/ ddecryption. As both encryption and decryption are possible, stateful (e.g. TCP) traffic can be sent/ received.
// Choice returns a string
func (obj *macsecCryptoEngine) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *macsecCryptoEngine) setChoice(value MacsecCryptoEngineChoiceEnum) MacsecCryptoEngine {
	intValue, ok := otg.MacsecCryptoEngine_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecCryptoEngineChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecCryptoEngine_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.StatefulEncryptionDecryption = nil
	obj.statefulEncryptionDecryptionHolder = nil
	obj.obj.StatelessEncryptionOnly = nil
	obj.statelessEncryptionOnlyHolder = nil

	if value == MacsecCryptoEngineChoice.STATELESS_ENCRYPTION_ONLY {
		obj.obj.StatelessEncryptionOnly = NewMacsecCryptoEngineStatelessEncryptionOnly().msg()
	}

	if value == MacsecCryptoEngineChoice.STATEFUL_ENCRYPTION_DECRYPTION {
		obj.obj.StatefulEncryptionDecryption = NewMacsecCryptoEngineStatefulEncryptionDecryption().msg()
	}

	return obj
}

// description is TBD
// StatelessEncryptionOnly returns a MacsecCryptoEngineStatelessEncryptionOnly
func (obj *macsecCryptoEngine) StatelessEncryptionOnly() MacsecCryptoEngineStatelessEncryptionOnly {
	if obj.obj.StatelessEncryptionOnly == nil {
		obj.setChoice(MacsecCryptoEngineChoice.STATELESS_ENCRYPTION_ONLY)
	}
	if obj.statelessEncryptionOnlyHolder == nil {
		obj.statelessEncryptionOnlyHolder = &macsecCryptoEngineStatelessEncryptionOnly{obj: obj.obj.StatelessEncryptionOnly}
	}
	return obj.statelessEncryptionOnlyHolder
}

// description is TBD
// StatelessEncryptionOnly returns a MacsecCryptoEngineStatelessEncryptionOnly
func (obj *macsecCryptoEngine) HasStatelessEncryptionOnly() bool {
	return obj.obj.StatelessEncryptionOnly != nil
}

// description is TBD
// SetStatelessEncryptionOnly sets the MacsecCryptoEngineStatelessEncryptionOnly value in the MacsecCryptoEngine object
func (obj *macsecCryptoEngine) SetStatelessEncryptionOnly(value MacsecCryptoEngineStatelessEncryptionOnly) MacsecCryptoEngine {
	obj.setChoice(MacsecCryptoEngineChoice.STATELESS_ENCRYPTION_ONLY)
	obj.statelessEncryptionOnlyHolder = nil
	obj.obj.StatelessEncryptionOnly = value.msg()

	return obj
}

// description is TBD
// StatefulEncryptionDecryption returns a MacsecCryptoEngineStatefulEncryptionDecryption
func (obj *macsecCryptoEngine) StatefulEncryptionDecryption() MacsecCryptoEngineStatefulEncryptionDecryption {
	if obj.obj.StatefulEncryptionDecryption == nil {
		obj.setChoice(MacsecCryptoEngineChoice.STATEFUL_ENCRYPTION_DECRYPTION)
	}
	if obj.statefulEncryptionDecryptionHolder == nil {
		obj.statefulEncryptionDecryptionHolder = &macsecCryptoEngineStatefulEncryptionDecryption{obj: obj.obj.StatefulEncryptionDecryption}
	}
	return obj.statefulEncryptionDecryptionHolder
}

// description is TBD
// StatefulEncryptionDecryption returns a MacsecCryptoEngineStatefulEncryptionDecryption
func (obj *macsecCryptoEngine) HasStatefulEncryptionDecryption() bool {
	return obj.obj.StatefulEncryptionDecryption != nil
}

// description is TBD
// SetStatefulEncryptionDecryption sets the MacsecCryptoEngineStatefulEncryptionDecryption value in the MacsecCryptoEngine object
func (obj *macsecCryptoEngine) SetStatefulEncryptionDecryption(value MacsecCryptoEngineStatefulEncryptionDecryption) MacsecCryptoEngine {
	obj.setChoice(MacsecCryptoEngineChoice.STATEFUL_ENCRYPTION_DECRYPTION)
	obj.statefulEncryptionDecryptionHolder = nil
	obj.obj.StatefulEncryptionDecryption = value.msg()

	return obj
}

func (obj *macsecCryptoEngine) validateObj(vObj *validation, set_default bool) {
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

func (obj *macsecCryptoEngine) setDefault() {
	var choices_set int = 0
	var choice MacsecCryptoEngineChoiceEnum

	if obj.obj.StatelessEncryptionOnly != nil {
		choices_set += 1
		choice = MacsecCryptoEngineChoice.STATELESS_ENCRYPTION_ONLY
	}

	if obj.obj.StatefulEncryptionDecryption != nil {
		choices_set += 1
		choice = MacsecCryptoEngineChoice.STATEFUL_ENCRYPTION_DECRYPTION
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MacsecCryptoEngineChoice.STATELESS_ENCRYPTION_ONLY)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MacsecCryptoEngine")
			}
		} else {
			intVal := otg.MacsecCryptoEngine_Choice_Enum_value[string(choice)]
			enumValue := otg.MacsecCryptoEngine_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

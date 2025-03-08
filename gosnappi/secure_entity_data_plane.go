package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityDataPlane *****
type secureEntityDataPlane struct {
	validation
	obj                 *otg.SecureEntityDataPlane
	marshaller          marshalSecureEntityDataPlane
	unMarshaller        unMarshalSecureEntityDataPlane
	encapsulationHolder SecureEntityDataPlaneEncapsulation
}

func NewSecureEntityDataPlane() SecureEntityDataPlane {
	obj := secureEntityDataPlane{obj: &otg.SecureEntityDataPlane{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityDataPlane) msg() *otg.SecureEntityDataPlane {
	return obj.obj
}

func (obj *secureEntityDataPlane) setMsg(msg *otg.SecureEntityDataPlane) SecureEntityDataPlane {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityDataPlane struct {
	obj *secureEntityDataPlane
}

type marshalSecureEntityDataPlane interface {
	// ToProto marshals SecureEntityDataPlane to protobuf object *otg.SecureEntityDataPlane
	ToProto() (*otg.SecureEntityDataPlane, error)
	// ToPbText marshals SecureEntityDataPlane to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityDataPlane to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityDataPlane to JSON text
	ToJson() (string, error)
}

type unMarshalsecureEntityDataPlane struct {
	obj *secureEntityDataPlane
}

type unMarshalSecureEntityDataPlane interface {
	// FromProto unmarshals SecureEntityDataPlane from protobuf object *otg.SecureEntityDataPlane
	FromProto(msg *otg.SecureEntityDataPlane) (SecureEntityDataPlane, error)
	// FromPbText unmarshals SecureEntityDataPlane from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityDataPlane from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityDataPlane from JSON text
	FromJson(value string) error
}

func (obj *secureEntityDataPlane) Marshal() marshalSecureEntityDataPlane {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityDataPlane{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityDataPlane) Unmarshal() unMarshalSecureEntityDataPlane {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityDataPlane{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityDataPlane) ToProto() (*otg.SecureEntityDataPlane, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityDataPlane) FromProto(msg *otg.SecureEntityDataPlane) (SecureEntityDataPlane, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityDataPlane) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityDataPlane) FromPbText(value string) error {
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

func (m *marshalsecureEntityDataPlane) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityDataPlane) FromYaml(value string) error {
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

func (m *marshalsecureEntityDataPlane) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityDataPlane) FromJson(value string) error {
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

func (obj *secureEntityDataPlane) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityDataPlane) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityDataPlane) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityDataPlane) Clone() (SecureEntityDataPlane, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityDataPlane()
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

func (obj *secureEntityDataPlane) setNil() {
	obj.encapsulationHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// SecureEntityDataPlane is a container of data plane properties.
type SecureEntityDataPlane interface {
	Validation
	// msg marshals SecureEntityDataPlane to protobuf object *otg.SecureEntityDataPlane
	// and doesn't set defaults
	msg() *otg.SecureEntityDataPlane
	// setMsg unmarshals SecureEntityDataPlane from protobuf object *otg.SecureEntityDataPlane
	// and doesn't set defaults
	setMsg(*otg.SecureEntityDataPlane) SecureEntityDataPlane
	// provides marshal interface
	Marshal() marshalSecureEntityDataPlane
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityDataPlane
	// validate validates SecureEntityDataPlane
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityDataPlane, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns SecureEntityDataPlaneChoiceEnum, set in SecureEntityDataPlane
	Choice() SecureEntityDataPlaneChoiceEnum
	// setChoice assigns SecureEntityDataPlaneChoiceEnum provided by user to SecureEntityDataPlane
	setChoice(value SecureEntityDataPlaneChoiceEnum) SecureEntityDataPlane
	// HasChoice checks if Choice has been set in SecureEntityDataPlane
	HasChoice() bool
	// getter for NoEncapsulation to set choice.
	NoEncapsulation()
	// Encapsulation returns SecureEntityDataPlaneEncapsulation, set in SecureEntityDataPlane.
	// SecureEntityDataPlaneEncapsulation is a container of encapsulation properties for a secure entity(SecY).
	Encapsulation() SecureEntityDataPlaneEncapsulation
	// SetEncapsulation assigns SecureEntityDataPlaneEncapsulation provided by user to SecureEntityDataPlane.
	// SecureEntityDataPlaneEncapsulation is a container of encapsulation properties for a secure entity(SecY).
	SetEncapsulation(value SecureEntityDataPlaneEncapsulation) SecureEntityDataPlane
	// HasEncapsulation checks if Encapsulation has been set in SecureEntityDataPlane
	HasEncapsulation() bool
	setNil()
}

type SecureEntityDataPlaneChoiceEnum string

// Enum of Choice on SecureEntityDataPlane
var SecureEntityDataPlaneChoice = struct {
	ENCAPSULATION    SecureEntityDataPlaneChoiceEnum
	NO_ENCAPSULATION SecureEntityDataPlaneChoiceEnum
}{
	ENCAPSULATION:    SecureEntityDataPlaneChoiceEnum("encapsulation"),
	NO_ENCAPSULATION: SecureEntityDataPlaneChoiceEnum("no_encapsulation"),
}

func (obj *secureEntityDataPlane) Choice() SecureEntityDataPlaneChoiceEnum {
	return SecureEntityDataPlaneChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for NoEncapsulation to set choice
func (obj *secureEntityDataPlane) NoEncapsulation() {
	obj.setChoice(SecureEntityDataPlaneChoice.NO_ENCAPSULATION)
}

// Choose "encapsulation" so that data packets are sent with MACsec encapsulation. Choose "no_encapsulation" so that data packets are sent without MACsec encapsulation.
// Choice returns a string
func (obj *secureEntityDataPlane) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *secureEntityDataPlane) setChoice(value SecureEntityDataPlaneChoiceEnum) SecureEntityDataPlane {
	intValue, ok := otg.SecureEntityDataPlane_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on SecureEntityDataPlaneChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.SecureEntityDataPlane_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Encapsulation = nil
	obj.encapsulationHolder = nil

	if value == SecureEntityDataPlaneChoice.ENCAPSULATION {
		obj.obj.Encapsulation = NewSecureEntityDataPlaneEncapsulation().msg()
	}

	return obj
}

// A container of encapsulation properties for a secure entity(SecY).
// Encapsulation returns a SecureEntityDataPlaneEncapsulation
func (obj *secureEntityDataPlane) Encapsulation() SecureEntityDataPlaneEncapsulation {
	if obj.obj.Encapsulation == nil {
		obj.setChoice(SecureEntityDataPlaneChoice.ENCAPSULATION)
	}
	if obj.encapsulationHolder == nil {
		obj.encapsulationHolder = &secureEntityDataPlaneEncapsulation{obj: obj.obj.Encapsulation}
	}
	return obj.encapsulationHolder
}

// A container of encapsulation properties for a secure entity(SecY).
// Encapsulation returns a SecureEntityDataPlaneEncapsulation
func (obj *secureEntityDataPlane) HasEncapsulation() bool {
	return obj.obj.Encapsulation != nil
}

// A container of encapsulation properties for a secure entity(SecY).
// SetEncapsulation sets the SecureEntityDataPlaneEncapsulation value in the SecureEntityDataPlane object
func (obj *secureEntityDataPlane) SetEncapsulation(value SecureEntityDataPlaneEncapsulation) SecureEntityDataPlane {
	obj.setChoice(SecureEntityDataPlaneChoice.ENCAPSULATION)
	obj.encapsulationHolder = nil
	obj.obj.Encapsulation = value.msg()

	return obj
}

func (obj *secureEntityDataPlane) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Encapsulation != nil {

		obj.Encapsulation().validateObj(vObj, set_default)
	}

}

func (obj *secureEntityDataPlane) setDefault() {
	var choices_set int = 0
	var choice SecureEntityDataPlaneChoiceEnum

	if obj.obj.Encapsulation != nil {
		choices_set += 1
		choice = SecureEntityDataPlaneChoice.ENCAPSULATION
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(SecureEntityDataPlaneChoice.ENCAPSULATION)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in SecureEntityDataPlane")
			}
		} else {
			intVal := otg.SecureEntityDataPlane_Choice_Enum_value[string(choice)]
			enumValue := otg.SecureEntityDataPlane_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

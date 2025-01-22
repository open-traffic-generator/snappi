package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecBasicKeyGeneration *****
type macsecBasicKeyGeneration struct {
	validation
	obj          *otg.MacsecBasicKeyGeneration
	marshaller   marshalMacsecBasicKeyGeneration
	unMarshaller unMarshalMacsecBasicKeyGeneration
	staticHolder MacsecBasicKeyGenerationStatic
}

func NewMacsecBasicKeyGeneration() MacsecBasicKeyGeneration {
	obj := macsecBasicKeyGeneration{obj: &otg.MacsecBasicKeyGeneration{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecBasicKeyGeneration) msg() *otg.MacsecBasicKeyGeneration {
	return obj.obj
}

func (obj *macsecBasicKeyGeneration) setMsg(msg *otg.MacsecBasicKeyGeneration) MacsecBasicKeyGeneration {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecBasicKeyGeneration struct {
	obj *macsecBasicKeyGeneration
}

type marshalMacsecBasicKeyGeneration interface {
	// ToProto marshals MacsecBasicKeyGeneration to protobuf object *otg.MacsecBasicKeyGeneration
	ToProto() (*otg.MacsecBasicKeyGeneration, error)
	// ToPbText marshals MacsecBasicKeyGeneration to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecBasicKeyGeneration to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecBasicKeyGeneration to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecBasicKeyGeneration struct {
	obj *macsecBasicKeyGeneration
}

type unMarshalMacsecBasicKeyGeneration interface {
	// FromProto unmarshals MacsecBasicKeyGeneration from protobuf object *otg.MacsecBasicKeyGeneration
	FromProto(msg *otg.MacsecBasicKeyGeneration) (MacsecBasicKeyGeneration, error)
	// FromPbText unmarshals MacsecBasicKeyGeneration from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecBasicKeyGeneration from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecBasicKeyGeneration from JSON text
	FromJson(value string) error
}

func (obj *macsecBasicKeyGeneration) Marshal() marshalMacsecBasicKeyGeneration {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecBasicKeyGeneration{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecBasicKeyGeneration) Unmarshal() unMarshalMacsecBasicKeyGeneration {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecBasicKeyGeneration{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecBasicKeyGeneration) ToProto() (*otg.MacsecBasicKeyGeneration, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecBasicKeyGeneration) FromProto(msg *otg.MacsecBasicKeyGeneration) (MacsecBasicKeyGeneration, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecBasicKeyGeneration) ToPbText() (string, error) {
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

func (m *unMarshalmacsecBasicKeyGeneration) FromPbText(value string) error {
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

func (m *marshalmacsecBasicKeyGeneration) ToYaml() (string, error) {
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

func (m *unMarshalmacsecBasicKeyGeneration) FromYaml(value string) error {
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

func (m *marshalmacsecBasicKeyGeneration) ToJson() (string, error) {
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

func (m *unMarshalmacsecBasicKeyGeneration) FromJson(value string) error {
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

func (obj *macsecBasicKeyGeneration) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecBasicKeyGeneration) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecBasicKeyGeneration) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecBasicKeyGeneration) Clone() (MacsecBasicKeyGeneration, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecBasicKeyGeneration()
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

func (obj *macsecBasicKeyGeneration) setNil() {
	obj.staticHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecBasicKeyGeneration is key Generation.
type MacsecBasicKeyGeneration interface {
	Validation
	// msg marshals MacsecBasicKeyGeneration to protobuf object *otg.MacsecBasicKeyGeneration
	// and doesn't set defaults
	msg() *otg.MacsecBasicKeyGeneration
	// setMsg unmarshals MacsecBasicKeyGeneration from protobuf object *otg.MacsecBasicKeyGeneration
	// and doesn't set defaults
	setMsg(*otg.MacsecBasicKeyGeneration) MacsecBasicKeyGeneration
	// provides marshal interface
	Marshal() marshalMacsecBasicKeyGeneration
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecBasicKeyGeneration
	// validate validates MacsecBasicKeyGeneration
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecBasicKeyGeneration, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MacsecBasicKeyGenerationChoiceEnum, set in MacsecBasicKeyGeneration
	Choice() MacsecBasicKeyGenerationChoiceEnum
	// setChoice assigns MacsecBasicKeyGenerationChoiceEnum provided by user to MacsecBasicKeyGeneration
	setChoice(value MacsecBasicKeyGenerationChoiceEnum) MacsecBasicKeyGeneration
	// HasChoice checks if Choice has been set in MacsecBasicKeyGeneration
	HasChoice() bool
	// getter for Mka to set choice.
	Mka()
	// Static returns MacsecBasicKeyGenerationStatic, set in MacsecBasicKeyGeneration.
	// MacsecBasicKeyGenerationStatic is the container for static key settings.
	Static() MacsecBasicKeyGenerationStatic
	// SetStatic assigns MacsecBasicKeyGenerationStatic provided by user to MacsecBasicKeyGeneration.
	// MacsecBasicKeyGenerationStatic is the container for static key settings.
	SetStatic(value MacsecBasicKeyGenerationStatic) MacsecBasicKeyGeneration
	// HasStatic checks if Static has been set in MacsecBasicKeyGeneration
	HasStatic() bool
	setNil()
}

type MacsecBasicKeyGenerationChoiceEnum string

// Enum of Choice on MacsecBasicKeyGeneration
var MacsecBasicKeyGenerationChoice = struct {
	MKA    MacsecBasicKeyGenerationChoiceEnum
	STATIC MacsecBasicKeyGenerationChoiceEnum
}{
	MKA:    MacsecBasicKeyGenerationChoiceEnum("mka"),
	STATIC: MacsecBasicKeyGenerationChoiceEnum("static"),
}

func (obj *macsecBasicKeyGeneration) Choice() MacsecBasicKeyGenerationChoiceEnum {
	return MacsecBasicKeyGenerationChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Mka to set choice
func (obj *macsecBasicKeyGeneration) Mka() {
	obj.setChoice(MacsecBasicKeyGenerationChoice.MKA)
}

// The type of control protocol of SecY.
// Choice returns a string
func (obj *macsecBasicKeyGeneration) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *macsecBasicKeyGeneration) setChoice(value MacsecBasicKeyGenerationChoiceEnum) MacsecBasicKeyGeneration {
	intValue, ok := otg.MacsecBasicKeyGeneration_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecBasicKeyGenerationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecBasicKeyGeneration_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Static = nil
	obj.staticHolder = nil

	if value == MacsecBasicKeyGenerationChoice.STATIC {
		obj.obj.Static = NewMacsecBasicKeyGenerationStatic().msg()
	}

	return obj
}

// description is TBD
// Static returns a MacsecBasicKeyGenerationStatic
func (obj *macsecBasicKeyGeneration) Static() MacsecBasicKeyGenerationStatic {
	if obj.obj.Static == nil {
		obj.setChoice(MacsecBasicKeyGenerationChoice.STATIC)
	}
	if obj.staticHolder == nil {
		obj.staticHolder = &macsecBasicKeyGenerationStatic{obj: obj.obj.Static}
	}
	return obj.staticHolder
}

// description is TBD
// Static returns a MacsecBasicKeyGenerationStatic
func (obj *macsecBasicKeyGeneration) HasStatic() bool {
	return obj.obj.Static != nil
}

// description is TBD
// SetStatic sets the MacsecBasicKeyGenerationStatic value in the MacsecBasicKeyGeneration object
func (obj *macsecBasicKeyGeneration) SetStatic(value MacsecBasicKeyGenerationStatic) MacsecBasicKeyGeneration {
	obj.setChoice(MacsecBasicKeyGenerationChoice.STATIC)
	obj.staticHolder = nil
	obj.obj.Static = value.msg()

	return obj
}

func (obj *macsecBasicKeyGeneration) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Static != nil {

		obj.Static().validateObj(vObj, set_default)
	}

}

func (obj *macsecBasicKeyGeneration) setDefault() {
	var choices_set int = 0
	var choice MacsecBasicKeyGenerationChoiceEnum

	if obj.obj.Static != nil {
		choices_set += 1
		choice = MacsecBasicKeyGenerationChoice.STATIC
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MacsecBasicKeyGenerationChoice.MKA)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MacsecBasicKeyGeneration")
			}
		} else {
			intVal := otg.MacsecBasicKeyGeneration_Choice_Enum_value[string(choice)]
			enumValue := otg.MacsecBasicKeyGeneration_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYBasicKeyGeneration *****
type macsecSecYBasicKeyGeneration struct {
	validation
	obj          *otg.MacsecSecYBasicKeyGeneration
	marshaller   marshalMacsecSecYBasicKeyGeneration
	unMarshaller unMarshalMacsecSecYBasicKeyGeneration
	staticHolder MacsecSecYBasicKeyGenerationStatic
}

func NewMacsecSecYBasicKeyGeneration() MacsecSecYBasicKeyGeneration {
	obj := macsecSecYBasicKeyGeneration{obj: &otg.MacsecSecYBasicKeyGeneration{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYBasicKeyGeneration) msg() *otg.MacsecSecYBasicKeyGeneration {
	return obj.obj
}

func (obj *macsecSecYBasicKeyGeneration) setMsg(msg *otg.MacsecSecYBasicKeyGeneration) MacsecSecYBasicKeyGeneration {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYBasicKeyGeneration struct {
	obj *macsecSecYBasicKeyGeneration
}

type marshalMacsecSecYBasicKeyGeneration interface {
	// ToProto marshals MacsecSecYBasicKeyGeneration to protobuf object *otg.MacsecSecYBasicKeyGeneration
	ToProto() (*otg.MacsecSecYBasicKeyGeneration, error)
	// ToPbText marshals MacsecSecYBasicKeyGeneration to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYBasicKeyGeneration to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYBasicKeyGeneration to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYBasicKeyGeneration struct {
	obj *macsecSecYBasicKeyGeneration
}

type unMarshalMacsecSecYBasicKeyGeneration interface {
	// FromProto unmarshals MacsecSecYBasicKeyGeneration from protobuf object *otg.MacsecSecYBasicKeyGeneration
	FromProto(msg *otg.MacsecSecYBasicKeyGeneration) (MacsecSecYBasicKeyGeneration, error)
	// FromPbText unmarshals MacsecSecYBasicKeyGeneration from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYBasicKeyGeneration from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYBasicKeyGeneration from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYBasicKeyGeneration) Marshal() marshalMacsecSecYBasicKeyGeneration {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYBasicKeyGeneration{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYBasicKeyGeneration) Unmarshal() unMarshalMacsecSecYBasicKeyGeneration {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYBasicKeyGeneration{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYBasicKeyGeneration) ToProto() (*otg.MacsecSecYBasicKeyGeneration, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYBasicKeyGeneration) FromProto(msg *otg.MacsecSecYBasicKeyGeneration) (MacsecSecYBasicKeyGeneration, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYBasicKeyGeneration) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYBasicKeyGeneration) FromPbText(value string) error {
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

func (m *marshalmacsecSecYBasicKeyGeneration) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYBasicKeyGeneration) FromYaml(value string) error {
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

func (m *marshalmacsecSecYBasicKeyGeneration) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYBasicKeyGeneration) FromJson(value string) error {
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

func (obj *macsecSecYBasicKeyGeneration) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYBasicKeyGeneration) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYBasicKeyGeneration) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYBasicKeyGeneration) Clone() (MacsecSecYBasicKeyGeneration, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYBasicKeyGeneration()
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

func (obj *macsecSecYBasicKeyGeneration) setNil() {
	obj.staticHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecSecYBasicKeyGeneration is key Generation.
type MacsecSecYBasicKeyGeneration interface {
	Validation
	// msg marshals MacsecSecYBasicKeyGeneration to protobuf object *otg.MacsecSecYBasicKeyGeneration
	// and doesn't set defaults
	msg() *otg.MacsecSecYBasicKeyGeneration
	// setMsg unmarshals MacsecSecYBasicKeyGeneration from protobuf object *otg.MacsecSecYBasicKeyGeneration
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYBasicKeyGeneration) MacsecSecYBasicKeyGeneration
	// provides marshal interface
	Marshal() marshalMacsecSecYBasicKeyGeneration
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYBasicKeyGeneration
	// validate validates MacsecSecYBasicKeyGeneration
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYBasicKeyGeneration, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MacsecSecYBasicKeyGenerationChoiceEnum, set in MacsecSecYBasicKeyGeneration
	Choice() MacsecSecYBasicKeyGenerationChoiceEnum
	// setChoice assigns MacsecSecYBasicKeyGenerationChoiceEnum provided by user to MacsecSecYBasicKeyGeneration
	setChoice(value MacsecSecYBasicKeyGenerationChoiceEnum) MacsecSecYBasicKeyGeneration
	// HasChoice checks if Choice has been set in MacsecSecYBasicKeyGeneration
	HasChoice() bool
	// getter for Mka to set choice.
	Mka()
	// Static returns MacsecSecYBasicKeyGenerationStatic, set in MacsecSecYBasicKeyGeneration.
	// MacsecSecYBasicKeyGenerationStatic is the container for static key settings.
	Static() MacsecSecYBasicKeyGenerationStatic
	// SetStatic assigns MacsecSecYBasicKeyGenerationStatic provided by user to MacsecSecYBasicKeyGeneration.
	// MacsecSecYBasicKeyGenerationStatic is the container for static key settings.
	SetStatic(value MacsecSecYBasicKeyGenerationStatic) MacsecSecYBasicKeyGeneration
	// HasStatic checks if Static has been set in MacsecSecYBasicKeyGeneration
	HasStatic() bool
	setNil()
}

type MacsecSecYBasicKeyGenerationChoiceEnum string

// Enum of Choice on MacsecSecYBasicKeyGeneration
var MacsecSecYBasicKeyGenerationChoice = struct {
	MKA    MacsecSecYBasicKeyGenerationChoiceEnum
	STATIC MacsecSecYBasicKeyGenerationChoiceEnum
}{
	MKA:    MacsecSecYBasicKeyGenerationChoiceEnum("mka"),
	STATIC: MacsecSecYBasicKeyGenerationChoiceEnum("static"),
}

func (obj *macsecSecYBasicKeyGeneration) Choice() MacsecSecYBasicKeyGenerationChoiceEnum {
	return MacsecSecYBasicKeyGenerationChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Mka to set choice
func (obj *macsecSecYBasicKeyGeneration) Mka() {
	obj.setChoice(MacsecSecYBasicKeyGenerationChoice.MKA)
}

// The type of control protocol of SecY.
// Choice returns a string
func (obj *macsecSecYBasicKeyGeneration) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *macsecSecYBasicKeyGeneration) setChoice(value MacsecSecYBasicKeyGenerationChoiceEnum) MacsecSecYBasicKeyGeneration {
	intValue, ok := otg.MacsecSecYBasicKeyGeneration_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecSecYBasicKeyGenerationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecSecYBasicKeyGeneration_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Static = nil
	obj.staticHolder = nil

	if value == MacsecSecYBasicKeyGenerationChoice.STATIC {
		obj.obj.Static = NewMacsecSecYBasicKeyGenerationStatic().msg()
	}

	return obj
}

// description is TBD
// Static returns a MacsecSecYBasicKeyGenerationStatic
func (obj *macsecSecYBasicKeyGeneration) Static() MacsecSecYBasicKeyGenerationStatic {
	if obj.obj.Static == nil {
		obj.setChoice(MacsecSecYBasicKeyGenerationChoice.STATIC)
	}
	if obj.staticHolder == nil {
		obj.staticHolder = &macsecSecYBasicKeyGenerationStatic{obj: obj.obj.Static}
	}
	return obj.staticHolder
}

// description is TBD
// Static returns a MacsecSecYBasicKeyGenerationStatic
func (obj *macsecSecYBasicKeyGeneration) HasStatic() bool {
	return obj.obj.Static != nil
}

// description is TBD
// SetStatic sets the MacsecSecYBasicKeyGenerationStatic value in the MacsecSecYBasicKeyGeneration object
func (obj *macsecSecYBasicKeyGeneration) SetStatic(value MacsecSecYBasicKeyGenerationStatic) MacsecSecYBasicKeyGeneration {
	obj.setChoice(MacsecSecYBasicKeyGenerationChoice.STATIC)
	obj.staticHolder = nil
	obj.obj.Static = value.msg()

	return obj
}

func (obj *macsecSecYBasicKeyGeneration) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Static != nil {

		obj.Static().validateObj(vObj, set_default)
	}

}

func (obj *macsecSecYBasicKeyGeneration) setDefault() {
	var choices_set int = 0
	var choice MacsecSecYBasicKeyGenerationChoiceEnum

	if obj.obj.Static != nil {
		choices_set += 1
		choice = MacsecSecYBasicKeyGenerationChoice.STATIC
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MacsecSecYBasicKeyGenerationChoice.MKA)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MacsecSecYBasicKeyGeneration")
			}
		} else {
			intVal := otg.MacsecSecYBasicKeyGeneration_Choice_Enum_value[string(choice)]
			enumValue := otg.MacsecSecYBasicKeyGeneration_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

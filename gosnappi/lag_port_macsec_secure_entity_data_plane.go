package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagPortMacsecSecureEntityDataPlane *****
type lagPortMacsecSecureEntityDataPlane struct {
	validation
	obj                 *otg.LagPortMacsecSecureEntityDataPlane
	marshaller          marshalLagPortMacsecSecureEntityDataPlane
	unMarshaller        unMarshalLagPortMacsecSecureEntityDataPlane
	encapsulationHolder LagPortMacsecSecureEntityDataPlaneEncapsulation
}

func NewLagPortMacsecSecureEntityDataPlane() LagPortMacsecSecureEntityDataPlane {
	obj := lagPortMacsecSecureEntityDataPlane{obj: &otg.LagPortMacsecSecureEntityDataPlane{}}
	obj.setDefault()
	return &obj
}

func (obj *lagPortMacsecSecureEntityDataPlane) msg() *otg.LagPortMacsecSecureEntityDataPlane {
	return obj.obj
}

func (obj *lagPortMacsecSecureEntityDataPlane) setMsg(msg *otg.LagPortMacsecSecureEntityDataPlane) LagPortMacsecSecureEntityDataPlane {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagPortMacsecSecureEntityDataPlane struct {
	obj *lagPortMacsecSecureEntityDataPlane
}

type marshalLagPortMacsecSecureEntityDataPlane interface {
	// ToProto marshals LagPortMacsecSecureEntityDataPlane to protobuf object *otg.LagPortMacsecSecureEntityDataPlane
	ToProto() (*otg.LagPortMacsecSecureEntityDataPlane, error)
	// ToPbText marshals LagPortMacsecSecureEntityDataPlane to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagPortMacsecSecureEntityDataPlane to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagPortMacsecSecureEntityDataPlane to JSON text
	ToJson() (string, error)
}

type unMarshallagPortMacsecSecureEntityDataPlane struct {
	obj *lagPortMacsecSecureEntityDataPlane
}

type unMarshalLagPortMacsecSecureEntityDataPlane interface {
	// FromProto unmarshals LagPortMacsecSecureEntityDataPlane from protobuf object *otg.LagPortMacsecSecureEntityDataPlane
	FromProto(msg *otg.LagPortMacsecSecureEntityDataPlane) (LagPortMacsecSecureEntityDataPlane, error)
	// FromPbText unmarshals LagPortMacsecSecureEntityDataPlane from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagPortMacsecSecureEntityDataPlane from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagPortMacsecSecureEntityDataPlane from JSON text
	FromJson(value string) error
}

func (obj *lagPortMacsecSecureEntityDataPlane) Marshal() marshalLagPortMacsecSecureEntityDataPlane {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagPortMacsecSecureEntityDataPlane{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagPortMacsecSecureEntityDataPlane) Unmarshal() unMarshalLagPortMacsecSecureEntityDataPlane {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagPortMacsecSecureEntityDataPlane{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagPortMacsecSecureEntityDataPlane) ToProto() (*otg.LagPortMacsecSecureEntityDataPlane, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagPortMacsecSecureEntityDataPlane) FromProto(msg *otg.LagPortMacsecSecureEntityDataPlane) (LagPortMacsecSecureEntityDataPlane, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagPortMacsecSecureEntityDataPlane) ToPbText() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityDataPlane) FromPbText(value string) error {
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

func (m *marshallagPortMacsecSecureEntityDataPlane) ToYaml() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityDataPlane) FromYaml(value string) error {
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

func (m *marshallagPortMacsecSecureEntityDataPlane) ToJson() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityDataPlane) FromJson(value string) error {
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

func (obj *lagPortMacsecSecureEntityDataPlane) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagPortMacsecSecureEntityDataPlane) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagPortMacsecSecureEntityDataPlane) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagPortMacsecSecureEntityDataPlane) Clone() (LagPortMacsecSecureEntityDataPlane, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagPortMacsecSecureEntityDataPlane()
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

func (obj *lagPortMacsecSecureEntityDataPlane) setNil() {
	obj.encapsulationHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LagPortMacsecSecureEntityDataPlane is a container of data plane properties.
type LagPortMacsecSecureEntityDataPlane interface {
	Validation
	// msg marshals LagPortMacsecSecureEntityDataPlane to protobuf object *otg.LagPortMacsecSecureEntityDataPlane
	// and doesn't set defaults
	msg() *otg.LagPortMacsecSecureEntityDataPlane
	// setMsg unmarshals LagPortMacsecSecureEntityDataPlane from protobuf object *otg.LagPortMacsecSecureEntityDataPlane
	// and doesn't set defaults
	setMsg(*otg.LagPortMacsecSecureEntityDataPlane) LagPortMacsecSecureEntityDataPlane
	// provides marshal interface
	Marshal() marshalLagPortMacsecSecureEntityDataPlane
	// provides unmarshal interface
	Unmarshal() unMarshalLagPortMacsecSecureEntityDataPlane
	// validate validates LagPortMacsecSecureEntityDataPlane
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagPortMacsecSecureEntityDataPlane, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns LagPortMacsecSecureEntityDataPlaneChoiceEnum, set in LagPortMacsecSecureEntityDataPlane
	Choice() LagPortMacsecSecureEntityDataPlaneChoiceEnum
	// setChoice assigns LagPortMacsecSecureEntityDataPlaneChoiceEnum provided by user to LagPortMacsecSecureEntityDataPlane
	setChoice(value LagPortMacsecSecureEntityDataPlaneChoiceEnum) LagPortMacsecSecureEntityDataPlane
	// HasChoice checks if Choice has been set in LagPortMacsecSecureEntityDataPlane
	HasChoice() bool
	// getter for NoEncapsulation to set choice.
	NoEncapsulation()
	// Encapsulation returns LagPortMacsecSecureEntityDataPlaneEncapsulation, set in LagPortMacsecSecureEntityDataPlane.
	// LagPortMacsecSecureEntityDataPlaneEncapsulation is a container of encapsulation properties for a secure entity(SecY).
	Encapsulation() LagPortMacsecSecureEntityDataPlaneEncapsulation
	// SetEncapsulation assigns LagPortMacsecSecureEntityDataPlaneEncapsulation provided by user to LagPortMacsecSecureEntityDataPlane.
	// LagPortMacsecSecureEntityDataPlaneEncapsulation is a container of encapsulation properties for a secure entity(SecY).
	SetEncapsulation(value LagPortMacsecSecureEntityDataPlaneEncapsulation) LagPortMacsecSecureEntityDataPlane
	// HasEncapsulation checks if Encapsulation has been set in LagPortMacsecSecureEntityDataPlane
	HasEncapsulation() bool
	setNil()
}

type LagPortMacsecSecureEntityDataPlaneChoiceEnum string

// Enum of Choice on LagPortMacsecSecureEntityDataPlane
var LagPortMacsecSecureEntityDataPlaneChoice = struct {
	ENCAPSULATION    LagPortMacsecSecureEntityDataPlaneChoiceEnum
	NO_ENCAPSULATION LagPortMacsecSecureEntityDataPlaneChoiceEnum
}{
	ENCAPSULATION:    LagPortMacsecSecureEntityDataPlaneChoiceEnum("encapsulation"),
	NO_ENCAPSULATION: LagPortMacsecSecureEntityDataPlaneChoiceEnum("no_encapsulation"),
}

func (obj *lagPortMacsecSecureEntityDataPlane) Choice() LagPortMacsecSecureEntityDataPlaneChoiceEnum {
	return LagPortMacsecSecureEntityDataPlaneChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for NoEncapsulation to set choice
func (obj *lagPortMacsecSecureEntityDataPlane) NoEncapsulation() {
	obj.setChoice(LagPortMacsecSecureEntityDataPlaneChoice.NO_ENCAPSULATION)
}

// Choose "encapsulation" so that data packets are sent with MACsec encapsulation. Choose "no_encapsulation" so that data packets are sent without MACsec encapsulation.
// Choice returns a string
func (obj *lagPortMacsecSecureEntityDataPlane) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *lagPortMacsecSecureEntityDataPlane) setChoice(value LagPortMacsecSecureEntityDataPlaneChoiceEnum) LagPortMacsecSecureEntityDataPlane {
	intValue, ok := otg.LagPortMacsecSecureEntityDataPlane_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LagPortMacsecSecureEntityDataPlaneChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.LagPortMacsecSecureEntityDataPlane_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Encapsulation = nil
	obj.encapsulationHolder = nil

	if value == LagPortMacsecSecureEntityDataPlaneChoice.ENCAPSULATION {
		obj.obj.Encapsulation = NewLagPortMacsecSecureEntityDataPlaneEncapsulation().msg()
	}

	return obj
}

// A container of encapsulation properties for a secure entity(SecY).
// Encapsulation returns a LagPortMacsecSecureEntityDataPlaneEncapsulation
func (obj *lagPortMacsecSecureEntityDataPlane) Encapsulation() LagPortMacsecSecureEntityDataPlaneEncapsulation {
	if obj.obj.Encapsulation == nil {
		obj.setChoice(LagPortMacsecSecureEntityDataPlaneChoice.ENCAPSULATION)
	}
	if obj.encapsulationHolder == nil {
		obj.encapsulationHolder = &lagPortMacsecSecureEntityDataPlaneEncapsulation{obj: obj.obj.Encapsulation}
	}
	return obj.encapsulationHolder
}

// A container of encapsulation properties for a secure entity(SecY).
// Encapsulation returns a LagPortMacsecSecureEntityDataPlaneEncapsulation
func (obj *lagPortMacsecSecureEntityDataPlane) HasEncapsulation() bool {
	return obj.obj.Encapsulation != nil
}

// A container of encapsulation properties for a secure entity(SecY).
// SetEncapsulation sets the LagPortMacsecSecureEntityDataPlaneEncapsulation value in the LagPortMacsecSecureEntityDataPlane object
func (obj *lagPortMacsecSecureEntityDataPlane) SetEncapsulation(value LagPortMacsecSecureEntityDataPlaneEncapsulation) LagPortMacsecSecureEntityDataPlane {
	obj.setChoice(LagPortMacsecSecureEntityDataPlaneChoice.ENCAPSULATION)
	obj.encapsulationHolder = nil
	obj.obj.Encapsulation = value.msg()

	return obj
}

func (obj *lagPortMacsecSecureEntityDataPlane) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Encapsulation != nil {

		obj.Encapsulation().validateObj(vObj, set_default)
	}

}

func (obj *lagPortMacsecSecureEntityDataPlane) setDefault() {
	var choices_set int = 0
	var choice LagPortMacsecSecureEntityDataPlaneChoiceEnum

	if obj.obj.Encapsulation != nil {
		choices_set += 1
		choice = LagPortMacsecSecureEntityDataPlaneChoice.ENCAPSULATION
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(LagPortMacsecSecureEntityDataPlaneChoice.ENCAPSULATION)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in LagPortMacsecSecureEntityDataPlane")
			}
		} else {
			intVal := otg.LagPortMacsecSecureEntityDataPlane_Choice_Enum_value[string(choice)]
			enumValue := otg.LagPortMacsecSecureEntityDataPlane_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

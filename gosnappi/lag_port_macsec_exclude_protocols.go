package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagPortMacsecExcludeProtocols *****
type lagPortMacsecExcludeProtocols struct {
	validation
	obj               *otg.LagPortMacsecExcludeProtocols
	marshaller        marshalLagPortMacsecExcludeProtocols
	unMarshaller      unMarshalLagPortMacsecExcludeProtocols
	perProtocolHolder LagPortMacsecExcludeProtocolsPerProtocol
}

func NewLagPortMacsecExcludeProtocols() LagPortMacsecExcludeProtocols {
	obj := lagPortMacsecExcludeProtocols{obj: &otg.LagPortMacsecExcludeProtocols{}}
	obj.setDefault()
	return &obj
}

func (obj *lagPortMacsecExcludeProtocols) msg() *otg.LagPortMacsecExcludeProtocols {
	return obj.obj
}

func (obj *lagPortMacsecExcludeProtocols) setMsg(msg *otg.LagPortMacsecExcludeProtocols) LagPortMacsecExcludeProtocols {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagPortMacsecExcludeProtocols struct {
	obj *lagPortMacsecExcludeProtocols
}

type marshalLagPortMacsecExcludeProtocols interface {
	// ToProto marshals LagPortMacsecExcludeProtocols to protobuf object *otg.LagPortMacsecExcludeProtocols
	ToProto() (*otg.LagPortMacsecExcludeProtocols, error)
	// ToPbText marshals LagPortMacsecExcludeProtocols to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagPortMacsecExcludeProtocols to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagPortMacsecExcludeProtocols to JSON text
	ToJson() (string, error)
}

type unMarshallagPortMacsecExcludeProtocols struct {
	obj *lagPortMacsecExcludeProtocols
}

type unMarshalLagPortMacsecExcludeProtocols interface {
	// FromProto unmarshals LagPortMacsecExcludeProtocols from protobuf object *otg.LagPortMacsecExcludeProtocols
	FromProto(msg *otg.LagPortMacsecExcludeProtocols) (LagPortMacsecExcludeProtocols, error)
	// FromPbText unmarshals LagPortMacsecExcludeProtocols from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagPortMacsecExcludeProtocols from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagPortMacsecExcludeProtocols from JSON text
	FromJson(value string) error
}

func (obj *lagPortMacsecExcludeProtocols) Marshal() marshalLagPortMacsecExcludeProtocols {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagPortMacsecExcludeProtocols{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagPortMacsecExcludeProtocols) Unmarshal() unMarshalLagPortMacsecExcludeProtocols {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagPortMacsecExcludeProtocols{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagPortMacsecExcludeProtocols) ToProto() (*otg.LagPortMacsecExcludeProtocols, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagPortMacsecExcludeProtocols) FromProto(msg *otg.LagPortMacsecExcludeProtocols) (LagPortMacsecExcludeProtocols, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagPortMacsecExcludeProtocols) ToPbText() (string, error) {
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

func (m *unMarshallagPortMacsecExcludeProtocols) FromPbText(value string) error {
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

func (m *marshallagPortMacsecExcludeProtocols) ToYaml() (string, error) {
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

func (m *unMarshallagPortMacsecExcludeProtocols) FromYaml(value string) error {
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

func (m *marshallagPortMacsecExcludeProtocols) ToJson() (string, error) {
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

func (m *unMarshallagPortMacsecExcludeProtocols) FromJson(value string) error {
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

func (obj *lagPortMacsecExcludeProtocols) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagPortMacsecExcludeProtocols) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagPortMacsecExcludeProtocols) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagPortMacsecExcludeProtocols) Clone() (LagPortMacsecExcludeProtocols, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagPortMacsecExcludeProtocols()
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

func (obj *lagPortMacsecExcludeProtocols) setNil() {
	obj.perProtocolHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LagPortMacsecExcludeProtocols is protocols excluded from MACsec encapsulation at Tx.
type LagPortMacsecExcludeProtocols interface {
	Validation
	// msg marshals LagPortMacsecExcludeProtocols to protobuf object *otg.LagPortMacsecExcludeProtocols
	// and doesn't set defaults
	msg() *otg.LagPortMacsecExcludeProtocols
	// setMsg unmarshals LagPortMacsecExcludeProtocols from protobuf object *otg.LagPortMacsecExcludeProtocols
	// and doesn't set defaults
	setMsg(*otg.LagPortMacsecExcludeProtocols) LagPortMacsecExcludeProtocols
	// provides marshal interface
	Marshal() marshalLagPortMacsecExcludeProtocols
	// provides unmarshal interface
	Unmarshal() unMarshalLagPortMacsecExcludeProtocols
	// validate validates LagPortMacsecExcludeProtocols
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagPortMacsecExcludeProtocols, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns LagPortMacsecExcludeProtocolsChoiceEnum, set in LagPortMacsecExcludeProtocols
	Choice() LagPortMacsecExcludeProtocolsChoiceEnum
	// setChoice assigns LagPortMacsecExcludeProtocolsChoiceEnum provided by user to LagPortMacsecExcludeProtocols
	setChoice(value LagPortMacsecExcludeProtocolsChoiceEnum) LagPortMacsecExcludeProtocols
	// HasChoice checks if Choice has been set in LagPortMacsecExcludeProtocols
	HasChoice() bool
	// getter for None to set choice.
	None()
	// PerProtocol returns LagPortMacsecExcludeProtocolsPerProtocol, set in LagPortMacsecExcludeProtocols.
	// LagPortMacsecExcludeProtocolsPerProtocol is per protocol choice to exclude from MACsec encapsulation at Tx.
	PerProtocol() LagPortMacsecExcludeProtocolsPerProtocol
	// SetPerProtocol assigns LagPortMacsecExcludeProtocolsPerProtocol provided by user to LagPortMacsecExcludeProtocols.
	// LagPortMacsecExcludeProtocolsPerProtocol is per protocol choice to exclude from MACsec encapsulation at Tx.
	SetPerProtocol(value LagPortMacsecExcludeProtocolsPerProtocol) LagPortMacsecExcludeProtocols
	// HasPerProtocol checks if PerProtocol has been set in LagPortMacsecExcludeProtocols
	HasPerProtocol() bool
	setNil()
}

type LagPortMacsecExcludeProtocolsChoiceEnum string

// Enum of Choice on LagPortMacsecExcludeProtocols
var LagPortMacsecExcludeProtocolsChoice = struct {
	NONE         LagPortMacsecExcludeProtocolsChoiceEnum
	PER_PROTOCOL LagPortMacsecExcludeProtocolsChoiceEnum
}{
	NONE:         LagPortMacsecExcludeProtocolsChoiceEnum("none"),
	PER_PROTOCOL: LagPortMacsecExcludeProtocolsChoiceEnum("per_protocol"),
}

func (obj *lagPortMacsecExcludeProtocols) Choice() LagPortMacsecExcludeProtocolsChoiceEnum {
	return LagPortMacsecExcludeProtocolsChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for None to set choice
func (obj *lagPortMacsecExcludeProtocols) None() {
	obj.setChoice(LagPortMacsecExcludeProtocolsChoice.NONE)
}

// Choose none to exclude no protocols i.e. all protocols are MACsec encapsulated at Tx. Choose per_protocol to opt per protocol.
// Choice returns a string
func (obj *lagPortMacsecExcludeProtocols) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *lagPortMacsecExcludeProtocols) setChoice(value LagPortMacsecExcludeProtocolsChoiceEnum) LagPortMacsecExcludeProtocols {
	intValue, ok := otg.LagPortMacsecExcludeProtocols_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LagPortMacsecExcludeProtocolsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.LagPortMacsecExcludeProtocols_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.PerProtocol = nil
	obj.perProtocolHolder = nil

	if value == LagPortMacsecExcludeProtocolsChoice.PER_PROTOCOL {
		obj.obj.PerProtocol = NewLagPortMacsecExcludeProtocolsPerProtocol().msg()
	}

	return obj
}

// description is TBD
// PerProtocol returns a LagPortMacsecExcludeProtocolsPerProtocol
func (obj *lagPortMacsecExcludeProtocols) PerProtocol() LagPortMacsecExcludeProtocolsPerProtocol {
	if obj.obj.PerProtocol == nil {
		obj.setChoice(LagPortMacsecExcludeProtocolsChoice.PER_PROTOCOL)
	}
	if obj.perProtocolHolder == nil {
		obj.perProtocolHolder = &lagPortMacsecExcludeProtocolsPerProtocol{obj: obj.obj.PerProtocol}
	}
	return obj.perProtocolHolder
}

// description is TBD
// PerProtocol returns a LagPortMacsecExcludeProtocolsPerProtocol
func (obj *lagPortMacsecExcludeProtocols) HasPerProtocol() bool {
	return obj.obj.PerProtocol != nil
}

// description is TBD
// SetPerProtocol sets the LagPortMacsecExcludeProtocolsPerProtocol value in the LagPortMacsecExcludeProtocols object
func (obj *lagPortMacsecExcludeProtocols) SetPerProtocol(value LagPortMacsecExcludeProtocolsPerProtocol) LagPortMacsecExcludeProtocols {
	obj.setChoice(LagPortMacsecExcludeProtocolsChoice.PER_PROTOCOL)
	obj.perProtocolHolder = nil
	obj.obj.PerProtocol = value.msg()

	return obj
}

func (obj *lagPortMacsecExcludeProtocols) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PerProtocol != nil {

		obj.PerProtocol().validateObj(vObj, set_default)
	}

}

func (obj *lagPortMacsecExcludeProtocols) setDefault() {
	var choices_set int = 0
	var choice LagPortMacsecExcludeProtocolsChoiceEnum

	if obj.obj.PerProtocol != nil {
		choices_set += 1
		choice = LagPortMacsecExcludeProtocolsChoice.PER_PROTOCOL
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(LagPortMacsecExcludeProtocolsChoice.NONE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in LagPortMacsecExcludeProtocols")
			}
		} else {
			intVal := otg.LagPortMacsecExcludeProtocols_Choice_Enum_value[string(choice)]
			enumValue := otg.LagPortMacsecExcludeProtocols_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

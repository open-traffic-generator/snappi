package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIcmp *****
type flowIcmp struct {
	validation
	obj          *otg.FlowIcmp
	marshaller   marshalFlowIcmp
	unMarshaller unMarshalFlowIcmp
	echoHolder   FlowIcmpEcho
}

func NewFlowIcmp() FlowIcmp {
	obj := flowIcmp{obj: &otg.FlowIcmp{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIcmp) msg() *otg.FlowIcmp {
	return obj.obj
}

func (obj *flowIcmp) setMsg(msg *otg.FlowIcmp) FlowIcmp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIcmp struct {
	obj *flowIcmp
}

type marshalFlowIcmp interface {
	// ToProto marshals FlowIcmp to protobuf object *otg.FlowIcmp
	ToProto() (*otg.FlowIcmp, error)
	// ToPbText marshals FlowIcmp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIcmp to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIcmp to JSON text
	ToJson() (string, error)
}

type unMarshalflowIcmp struct {
	obj *flowIcmp
}

type unMarshalFlowIcmp interface {
	// FromProto unmarshals FlowIcmp from protobuf object *otg.FlowIcmp
	FromProto(msg *otg.FlowIcmp) (FlowIcmp, error)
	// FromPbText unmarshals FlowIcmp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIcmp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIcmp from JSON text
	FromJson(value string) error
}

func (obj *flowIcmp) Marshal() marshalFlowIcmp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIcmp{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIcmp) Unmarshal() unMarshalFlowIcmp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIcmp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIcmp) ToProto() (*otg.FlowIcmp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIcmp) FromProto(msg *otg.FlowIcmp) (FlowIcmp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIcmp) ToPbText() (string, error) {
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

func (m *unMarshalflowIcmp) FromPbText(value string) error {
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

func (m *marshalflowIcmp) ToYaml() (string, error) {
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

func (m *unMarshalflowIcmp) FromYaml(value string) error {
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

func (m *marshalflowIcmp) ToJson() (string, error) {
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

func (m *unMarshalflowIcmp) FromJson(value string) error {
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

func (obj *flowIcmp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIcmp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIcmp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIcmp) Clone() (FlowIcmp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIcmp()
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

func (obj *flowIcmp) setNil() {
	obj.echoHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIcmp is iCMP packet header
type FlowIcmp interface {
	Validation
	// msg marshals FlowIcmp to protobuf object *otg.FlowIcmp
	// and doesn't set defaults
	msg() *otg.FlowIcmp
	// setMsg unmarshals FlowIcmp from protobuf object *otg.FlowIcmp
	// and doesn't set defaults
	setMsg(*otg.FlowIcmp) FlowIcmp
	// provides marshal interface
	Marshal() marshalFlowIcmp
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIcmp
	// validate validates FlowIcmp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIcmp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowIcmpChoiceEnum, set in FlowIcmp
	Choice() FlowIcmpChoiceEnum
	// setChoice assigns FlowIcmpChoiceEnum provided by user to FlowIcmp
	setChoice(value FlowIcmpChoiceEnum) FlowIcmp
	// HasChoice checks if Choice has been set in FlowIcmp
	HasChoice() bool
	// Echo returns FlowIcmpEcho, set in FlowIcmp.
	// FlowIcmpEcho is packet Header for ICMP echo request
	Echo() FlowIcmpEcho
	// SetEcho assigns FlowIcmpEcho provided by user to FlowIcmp.
	// FlowIcmpEcho is packet Header for ICMP echo request
	SetEcho(value FlowIcmpEcho) FlowIcmp
	// HasEcho checks if Echo has been set in FlowIcmp
	HasEcho() bool
	setNil()
}

type FlowIcmpChoiceEnum string

// Enum of Choice on FlowIcmp
var FlowIcmpChoice = struct {
	ECHO FlowIcmpChoiceEnum
}{
	ECHO: FlowIcmpChoiceEnum("echo"),
}

func (obj *flowIcmp) Choice() FlowIcmpChoiceEnum {
	return FlowIcmpChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowIcmp) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowIcmp) setChoice(value FlowIcmpChoiceEnum) FlowIcmp {
	intValue, ok := otg.FlowIcmp_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowIcmpChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowIcmp_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Echo = nil
	obj.echoHolder = nil

	if value == FlowIcmpChoice.ECHO {
		obj.obj.Echo = NewFlowIcmpEcho().msg()
	}

	return obj
}

// description is TBD
// Echo returns a FlowIcmpEcho
func (obj *flowIcmp) Echo() FlowIcmpEcho {
	if obj.obj.Echo == nil {
		obj.setChoice(FlowIcmpChoice.ECHO)
	}
	if obj.echoHolder == nil {
		obj.echoHolder = &flowIcmpEcho{obj: obj.obj.Echo}
	}
	return obj.echoHolder
}

// description is TBD
// Echo returns a FlowIcmpEcho
func (obj *flowIcmp) HasEcho() bool {
	return obj.obj.Echo != nil
}

// description is TBD
// SetEcho sets the FlowIcmpEcho value in the FlowIcmp object
func (obj *flowIcmp) SetEcho(value FlowIcmpEcho) FlowIcmp {
	obj.setChoice(FlowIcmpChoice.ECHO)
	obj.echoHolder = nil
	obj.obj.Echo = value.msg()

	return obj
}

func (obj *flowIcmp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Echo != nil {

		obj.Echo().validateObj(vObj, set_default)
	}

}

func (obj *flowIcmp) setDefault() {
	var choices_set int = 0
	var choice FlowIcmpChoiceEnum

	if obj.obj.Echo != nil {
		choices_set += 1
		choice = FlowIcmpChoice.ECHO
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowIcmpChoice.ECHO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowIcmp")
			}
		} else {
			intVal := otg.FlowIcmp_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowIcmp_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

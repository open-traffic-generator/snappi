package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIcmpv6 *****
type flowIcmpv6 struct {
	validation
	obj          *otg.FlowIcmpv6
	marshaller   marshalFlowIcmpv6
	unMarshaller unMarshalFlowIcmpv6
	echoHolder   FlowIcmpv6Echo
}

func NewFlowIcmpv6() FlowIcmpv6 {
	obj := flowIcmpv6{obj: &otg.FlowIcmpv6{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIcmpv6) msg() *otg.FlowIcmpv6 {
	return obj.obj
}

func (obj *flowIcmpv6) setMsg(msg *otg.FlowIcmpv6) FlowIcmpv6 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIcmpv6 struct {
	obj *flowIcmpv6
}

type marshalFlowIcmpv6 interface {
	// ToProto marshals FlowIcmpv6 to protobuf object *otg.FlowIcmpv6
	ToProto() (*otg.FlowIcmpv6, error)
	// ToPbText marshals FlowIcmpv6 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIcmpv6 to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIcmpv6 to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowIcmpv6 to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowIcmpv6 struct {
	obj *flowIcmpv6
}

type unMarshalFlowIcmpv6 interface {
	// FromProto unmarshals FlowIcmpv6 from protobuf object *otg.FlowIcmpv6
	FromProto(msg *otg.FlowIcmpv6) (FlowIcmpv6, error)
	// FromPbText unmarshals FlowIcmpv6 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIcmpv6 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIcmpv6 from JSON text
	FromJson(value string) error
}

func (obj *flowIcmpv6) Marshal() marshalFlowIcmpv6 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIcmpv6{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIcmpv6) Unmarshal() unMarshalFlowIcmpv6 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIcmpv6{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIcmpv6) ToProto() (*otg.FlowIcmpv6, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIcmpv6) FromProto(msg *otg.FlowIcmpv6) (FlowIcmpv6, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIcmpv6) ToPbText() (string, error) {
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

func (m *unMarshalflowIcmpv6) FromPbText(value string) error {
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

func (m *marshalflowIcmpv6) ToYaml() (string, error) {
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

func (m *unMarshalflowIcmpv6) FromYaml(value string) error {
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

func (m *marshalflowIcmpv6) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalflowIcmpv6) ToJson() (string, error) {
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

func (m *unMarshalflowIcmpv6) FromJson(value string) error {
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

func (obj *flowIcmpv6) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIcmpv6) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIcmpv6) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIcmpv6) Clone() (FlowIcmpv6, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIcmpv6()
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

func (obj *flowIcmpv6) setNil() {
	obj.echoHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIcmpv6 is iCMPv6 packet header
type FlowIcmpv6 interface {
	Validation
	// msg marshals FlowIcmpv6 to protobuf object *otg.FlowIcmpv6
	// and doesn't set defaults
	msg() *otg.FlowIcmpv6
	// setMsg unmarshals FlowIcmpv6 from protobuf object *otg.FlowIcmpv6
	// and doesn't set defaults
	setMsg(*otg.FlowIcmpv6) FlowIcmpv6
	// provides marshal interface
	Marshal() marshalFlowIcmpv6
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIcmpv6
	// validate validates FlowIcmpv6
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIcmpv6, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowIcmpv6ChoiceEnum, set in FlowIcmpv6
	Choice() FlowIcmpv6ChoiceEnum
	// setChoice assigns FlowIcmpv6ChoiceEnum provided by user to FlowIcmpv6
	setChoice(value FlowIcmpv6ChoiceEnum) FlowIcmpv6
	// HasChoice checks if Choice has been set in FlowIcmpv6
	HasChoice() bool
	// Echo returns FlowIcmpv6Echo, set in FlowIcmpv6.
	// FlowIcmpv6Echo is packet Header for ICMPv6 Echo
	Echo() FlowIcmpv6Echo
	// SetEcho assigns FlowIcmpv6Echo provided by user to FlowIcmpv6.
	// FlowIcmpv6Echo is packet Header for ICMPv6 Echo
	SetEcho(value FlowIcmpv6Echo) FlowIcmpv6
	// HasEcho checks if Echo has been set in FlowIcmpv6
	HasEcho() bool
	setNil()
}

type FlowIcmpv6ChoiceEnum string

// Enum of Choice on FlowIcmpv6
var FlowIcmpv6Choice = struct {
	ECHO FlowIcmpv6ChoiceEnum
}{
	ECHO: FlowIcmpv6ChoiceEnum("echo"),
}

func (obj *flowIcmpv6) Choice() FlowIcmpv6ChoiceEnum {
	return FlowIcmpv6ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowIcmpv6) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowIcmpv6) setChoice(value FlowIcmpv6ChoiceEnum) FlowIcmpv6 {
	intValue, ok := otg.FlowIcmpv6_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowIcmpv6ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowIcmpv6_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Echo = nil
	obj.echoHolder = nil

	if value == FlowIcmpv6Choice.ECHO {
		obj.obj.Echo = NewFlowIcmpv6Echo().msg()
	}

	return obj
}

// description is TBD
// Echo returns a FlowIcmpv6Echo
func (obj *flowIcmpv6) Echo() FlowIcmpv6Echo {
	if obj.obj.Echo == nil {
		obj.setChoice(FlowIcmpv6Choice.ECHO)
	}
	if obj.echoHolder == nil {
		obj.echoHolder = &flowIcmpv6Echo{obj: obj.obj.Echo}
	}
	return obj.echoHolder
}

// description is TBD
// Echo returns a FlowIcmpv6Echo
func (obj *flowIcmpv6) HasEcho() bool {
	return obj.obj.Echo != nil
}

// description is TBD
// SetEcho sets the FlowIcmpv6Echo value in the FlowIcmpv6 object
func (obj *flowIcmpv6) SetEcho(value FlowIcmpv6Echo) FlowIcmpv6 {
	obj.setChoice(FlowIcmpv6Choice.ECHO)
	obj.echoHolder = nil
	obj.obj.Echo = value.msg()

	return obj
}

func (obj *flowIcmpv6) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Echo != nil {

		obj.Echo().validateObj(vObj, set_default)
	}

}

func (obj *flowIcmpv6) setDefault() {
	var choices_set int = 0
	var choice FlowIcmpv6ChoiceEnum

	if obj.obj.Echo != nil {
		choices_set += 1
		choice = FlowIcmpv6Choice.ECHO
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowIcmpv6Choice.ECHO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowIcmpv6")
			}
		} else {
			intVal := otg.FlowIcmpv6_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowIcmpv6_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

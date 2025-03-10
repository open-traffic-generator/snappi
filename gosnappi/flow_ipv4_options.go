package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv4Options *****
type flowIpv4Options struct {
	validation
	obj          *otg.FlowIpv4Options
	marshaller   marshalFlowIpv4Options
	unMarshaller unMarshalFlowIpv4Options
	customHolder FlowIpv4OptionsCustom
}

func NewFlowIpv4Options() FlowIpv4Options {
	obj := flowIpv4Options{obj: &otg.FlowIpv4Options{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv4Options) msg() *otg.FlowIpv4Options {
	return obj.obj
}

func (obj *flowIpv4Options) setMsg(msg *otg.FlowIpv4Options) FlowIpv4Options {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv4Options struct {
	obj *flowIpv4Options
}

type marshalFlowIpv4Options interface {
	// ToProto marshals FlowIpv4Options to protobuf object *otg.FlowIpv4Options
	ToProto() (*otg.FlowIpv4Options, error)
	// ToPbText marshals FlowIpv4Options to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv4Options to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv4Options to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowIpv4Options to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowIpv4Options struct {
	obj *flowIpv4Options
}

type unMarshalFlowIpv4Options interface {
	// FromProto unmarshals FlowIpv4Options from protobuf object *otg.FlowIpv4Options
	FromProto(msg *otg.FlowIpv4Options) (FlowIpv4Options, error)
	// FromPbText unmarshals FlowIpv4Options from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv4Options from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv4Options from JSON text
	FromJson(value string) error
}

func (obj *flowIpv4Options) Marshal() marshalFlowIpv4Options {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv4Options{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv4Options) Unmarshal() unMarshalFlowIpv4Options {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv4Options{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv4Options) ToProto() (*otg.FlowIpv4Options, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv4Options) FromProto(msg *otg.FlowIpv4Options) (FlowIpv4Options, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv4Options) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv4Options) FromPbText(value string) error {
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

func (m *marshalflowIpv4Options) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv4Options) FromYaml(value string) error {
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

func (m *marshalflowIpv4Options) ToJsonRaw() (string, error) {
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

func (m *marshalflowIpv4Options) ToJson() (string, error) {
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

func (m *unMarshalflowIpv4Options) FromJson(value string) error {
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

func (obj *flowIpv4Options) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv4Options) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv4Options) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv4Options) Clone() (FlowIpv4Options, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv4Options()
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

func (obj *flowIpv4Options) setNil() {
	obj.customHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv4Options is iPv4 options are optional extensions for the IPv4 header that can be utilised to provide additional information about the IPv4 datagram.  It is encoded as a series of type, length and value attributes.  The IP header length MUST be increased to accommodate the extra bytes needed to encode the IP options. The length of the all options included to a IPv4 header should not exceed 40 bytes since IPv4 Header length (4 bits) can at max specify 15 4-word octets for a total of 60 bytes which includes 20 bytes needed for mandatory attributes of the IPv4 header. If the user adds multiples IPv4 options that exceeds 40 bytes and specify header length as "auto", implementation should throw error. Currently IP options supported are: 1. router_alert option allows devices to intercept packets not addressed to them directly as defined in RFC2113. 2. custom option is provided to configure user defined IP options as needed.
type FlowIpv4Options interface {
	Validation
	// msg marshals FlowIpv4Options to protobuf object *otg.FlowIpv4Options
	// and doesn't set defaults
	msg() *otg.FlowIpv4Options
	// setMsg unmarshals FlowIpv4Options from protobuf object *otg.FlowIpv4Options
	// and doesn't set defaults
	setMsg(*otg.FlowIpv4Options) FlowIpv4Options
	// provides marshal interface
	Marshal() marshalFlowIpv4Options
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv4Options
	// validate validates FlowIpv4Options
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv4Options, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowIpv4OptionsChoiceEnum, set in FlowIpv4Options
	Choice() FlowIpv4OptionsChoiceEnum
	// setChoice assigns FlowIpv4OptionsChoiceEnum provided by user to FlowIpv4Options
	setChoice(value FlowIpv4OptionsChoiceEnum) FlowIpv4Options
	// HasChoice checks if Choice has been set in FlowIpv4Options
	HasChoice() bool
	// getter for RouterAlert to set choice.
	RouterAlert()
	// Custom returns FlowIpv4OptionsCustom, set in FlowIpv4Options.
	// FlowIpv4OptionsCustom is user defined IP options to be appended to the IPv4 header.
	Custom() FlowIpv4OptionsCustom
	// SetCustom assigns FlowIpv4OptionsCustom provided by user to FlowIpv4Options.
	// FlowIpv4OptionsCustom is user defined IP options to be appended to the IPv4 header.
	SetCustom(value FlowIpv4OptionsCustom) FlowIpv4Options
	// HasCustom checks if Custom has been set in FlowIpv4Options
	HasCustom() bool
	setNil()
}

type FlowIpv4OptionsChoiceEnum string

// Enum of Choice on FlowIpv4Options
var FlowIpv4OptionsChoice = struct {
	ROUTER_ALERT FlowIpv4OptionsChoiceEnum
	CUSTOM       FlowIpv4OptionsChoiceEnum
}{
	ROUTER_ALERT: FlowIpv4OptionsChoiceEnum("router_alert"),
	CUSTOM:       FlowIpv4OptionsChoiceEnum("custom"),
}

func (obj *flowIpv4Options) Choice() FlowIpv4OptionsChoiceEnum {
	return FlowIpv4OptionsChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for RouterAlert to set choice
func (obj *flowIpv4Options) RouterAlert() {
	obj.setChoice(FlowIpv4OptionsChoice.ROUTER_ALERT)
}

// description is TBD
// Choice returns a string
func (obj *flowIpv4Options) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowIpv4Options) setChoice(value FlowIpv4OptionsChoiceEnum) FlowIpv4Options {
	intValue, ok := otg.FlowIpv4Options_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowIpv4OptionsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowIpv4Options_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.customHolder = nil

	if value == FlowIpv4OptionsChoice.CUSTOM {
		obj.obj.Custom = NewFlowIpv4OptionsCustom().msg()
	}

	return obj
}

// description is TBD
// Custom returns a FlowIpv4OptionsCustom
func (obj *flowIpv4Options) Custom() FlowIpv4OptionsCustom {
	if obj.obj.Custom == nil {
		obj.setChoice(FlowIpv4OptionsChoice.CUSTOM)
	}
	if obj.customHolder == nil {
		obj.customHolder = &flowIpv4OptionsCustom{obj: obj.obj.Custom}
	}
	return obj.customHolder
}

// description is TBD
// Custom returns a FlowIpv4OptionsCustom
func (obj *flowIpv4Options) HasCustom() bool {
	return obj.obj.Custom != nil
}

// description is TBD
// SetCustom sets the FlowIpv4OptionsCustom value in the FlowIpv4Options object
func (obj *flowIpv4Options) SetCustom(value FlowIpv4OptionsCustom) FlowIpv4Options {
	obj.setChoice(FlowIpv4OptionsChoice.CUSTOM)
	obj.customHolder = nil
	obj.obj.Custom = value.msg()

	return obj
}

func (obj *flowIpv4Options) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		obj.Custom().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv4Options) setDefault() {
	var choices_set int = 0
	var choice FlowIpv4OptionsChoiceEnum

	if obj.obj.Custom != nil {
		choices_set += 1
		choice = FlowIpv4OptionsChoice.CUSTOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowIpv4OptionsChoice.ROUTER_ALERT)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowIpv4Options")
			}
		} else {
			intVal := otg.FlowIpv4Options_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowIpv4Options_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionResponse *****
type actionResponse struct {
	validation
	obj            *otg.ActionResponse
	marshaller     marshalActionResponse
	unMarshaller   unMarshalActionResponse
	protocolHolder ActionResponseProtocol
}

func NewActionResponse() ActionResponse {
	obj := actionResponse{obj: &otg.ActionResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *actionResponse) msg() *otg.ActionResponse {
	return obj.obj
}

func (obj *actionResponse) setMsg(msg *otg.ActionResponse) ActionResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionResponse struct {
	obj *actionResponse
}

type marshalActionResponse interface {
	// ToProto marshals ActionResponse to protobuf object *otg.ActionResponse
	ToProto() (*otg.ActionResponse, error)
	// ToPbText marshals ActionResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionResponse to JSON text
	ToJson() (string, error)
}

type unMarshalactionResponse struct {
	obj *actionResponse
}

type unMarshalActionResponse interface {
	// FromProto unmarshals ActionResponse from protobuf object *otg.ActionResponse
	FromProto(msg *otg.ActionResponse) (ActionResponse, error)
	// FromPbText unmarshals ActionResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionResponse from JSON text
	FromJson(value string) error
}

func (obj *actionResponse) Marshal() marshalActionResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionResponse) Unmarshal() unMarshalActionResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionResponse) ToProto() (*otg.ActionResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionResponse) FromProto(msg *otg.ActionResponse) (ActionResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionResponse) ToPbText() (string, error) {
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

func (m *unMarshalactionResponse) FromPbText(value string) error {
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

func (m *marshalactionResponse) ToYaml() (string, error) {
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

func (m *unMarshalactionResponse) FromYaml(value string) error {
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

func (m *marshalactionResponse) ToJson() (string, error) {
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

func (m *unMarshalactionResponse) FromJson(value string) error {
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

func (obj *actionResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionResponse) Clone() (ActionResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionResponse()
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

func (obj *actionResponse) setNil() {
	obj.protocolHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionResponse is response for action triggered against configured resources.
type ActionResponse interface {
	Validation
	// msg marshals ActionResponse to protobuf object *otg.ActionResponse
	// and doesn't set defaults
	msg() *otg.ActionResponse
	// setMsg unmarshals ActionResponse from protobuf object *otg.ActionResponse
	// and doesn't set defaults
	setMsg(*otg.ActionResponse) ActionResponse
	// provides marshal interface
	Marshal() marshalActionResponse
	// provides unmarshal interface
	Unmarshal() unMarshalActionResponse
	// validate validates ActionResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ActionResponseChoiceEnum, set in ActionResponse
	Choice() ActionResponseChoiceEnum
	// setChoice assigns ActionResponseChoiceEnum provided by user to ActionResponse
	setChoice(value ActionResponseChoiceEnum) ActionResponse
	// Protocol returns ActionResponseProtocol, set in ActionResponse.
	// ActionResponseProtocol is response for actions associated with protocols on configured resources.
	Protocol() ActionResponseProtocol
	// SetProtocol assigns ActionResponseProtocol provided by user to ActionResponse.
	// ActionResponseProtocol is response for actions associated with protocols on configured resources.
	SetProtocol(value ActionResponseProtocol) ActionResponse
	// HasProtocol checks if Protocol has been set in ActionResponse
	HasProtocol() bool
	setNil()
}

type ActionResponseChoiceEnum string

// Enum of Choice on ActionResponse
var ActionResponseChoice = struct {
	PROTOCOL ActionResponseChoiceEnum
}{
	PROTOCOL: ActionResponseChoiceEnum("protocol"),
}

func (obj *actionResponse) Choice() ActionResponseChoiceEnum {
	return ActionResponseChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *actionResponse) setChoice(value ActionResponseChoiceEnum) ActionResponse {
	intValue, ok := otg.ActionResponse_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ActionResponseChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ActionResponse_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Protocol = nil
	obj.protocolHolder = nil

	if value == ActionResponseChoice.PROTOCOL {
		obj.obj.Protocol = NewActionResponseProtocol().msg()
	}

	return obj
}

// description is TBD
// Protocol returns a ActionResponseProtocol
func (obj *actionResponse) Protocol() ActionResponseProtocol {
	if obj.obj.Protocol == nil {
		obj.setChoice(ActionResponseChoice.PROTOCOL)
	}
	if obj.protocolHolder == nil {
		obj.protocolHolder = &actionResponseProtocol{obj: obj.obj.Protocol}
	}
	return obj.protocolHolder
}

// description is TBD
// Protocol returns a ActionResponseProtocol
func (obj *actionResponse) HasProtocol() bool {
	return obj.obj.Protocol != nil
}

// description is TBD
// SetProtocol sets the ActionResponseProtocol value in the ActionResponse object
func (obj *actionResponse) SetProtocol(value ActionResponseProtocol) ActionResponse {
	obj.setChoice(ActionResponseChoice.PROTOCOL)
	obj.protocolHolder = nil
	obj.obj.Protocol = value.msg()

	return obj
}

func (obj *actionResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface ActionResponse")
	}

	if obj.obj.Protocol != nil {

		obj.Protocol().validateObj(vObj, set_default)
	}

}

func (obj *actionResponse) setDefault() {
	var choices_set int = 0
	var choice ActionResponseChoiceEnum

	if obj.obj.Protocol != nil {
		choices_set += 1
		choice = ActionResponseChoice.PROTOCOL
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ActionResponse")
			}
		} else {
			intVal := otg.ActionResponse_Choice_Enum_value[string(choice)]
			enumValue := otg.ActionResponse_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

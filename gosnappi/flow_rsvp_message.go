package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPMessage *****
type flowRSVPMessage struct {
	validation
	obj          *otg.FlowRSVPMessage
	marshaller   marshalFlowRSVPMessage
	unMarshaller unMarshalFlowRSVPMessage
	pathHolder   FlowRSVPPathMessage
}

func NewFlowRSVPMessage() FlowRSVPMessage {
	obj := flowRSVPMessage{obj: &otg.FlowRSVPMessage{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPMessage) msg() *otg.FlowRSVPMessage {
	return obj.obj
}

func (obj *flowRSVPMessage) setMsg(msg *otg.FlowRSVPMessage) FlowRSVPMessage {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPMessage struct {
	obj *flowRSVPMessage
}

type marshalFlowRSVPMessage interface {
	// ToProto marshals FlowRSVPMessage to protobuf object *otg.FlowRSVPMessage
	ToProto() (*otg.FlowRSVPMessage, error)
	// ToPbText marshals FlowRSVPMessage to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPMessage to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPMessage to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPMessage struct {
	obj *flowRSVPMessage
}

type unMarshalFlowRSVPMessage interface {
	// FromProto unmarshals FlowRSVPMessage from protobuf object *otg.FlowRSVPMessage
	FromProto(msg *otg.FlowRSVPMessage) (FlowRSVPMessage, error)
	// FromPbText unmarshals FlowRSVPMessage from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPMessage from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPMessage from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPMessage) Marshal() marshalFlowRSVPMessage {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPMessage{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPMessage) Unmarshal() unMarshalFlowRSVPMessage {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPMessage{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPMessage) ToProto() (*otg.FlowRSVPMessage, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPMessage) FromProto(msg *otg.FlowRSVPMessage) (FlowRSVPMessage, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPMessage) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPMessage) FromPbText(value string) error {
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

func (m *marshalflowRSVPMessage) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPMessage) FromYaml(value string) error {
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

func (m *marshalflowRSVPMessage) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPMessage) FromJson(value string) error {
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

func (obj *flowRSVPMessage) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPMessage) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPMessage) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPMessage) Clone() (FlowRSVPMessage, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPMessage()
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

func (obj *flowRSVPMessage) setNil() {
	obj.pathHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPMessage is description is TBD
type FlowRSVPMessage interface {
	Validation
	// msg marshals FlowRSVPMessage to protobuf object *otg.FlowRSVPMessage
	// and doesn't set defaults
	msg() *otg.FlowRSVPMessage
	// setMsg unmarshals FlowRSVPMessage from protobuf object *otg.FlowRSVPMessage
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPMessage) FlowRSVPMessage
	// provides marshal interface
	Marshal() marshalFlowRSVPMessage
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPMessage
	// validate validates FlowRSVPMessage
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPMessage, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPMessageChoiceEnum, set in FlowRSVPMessage
	Choice() FlowRSVPMessageChoiceEnum
	// setChoice assigns FlowRSVPMessageChoiceEnum provided by user to FlowRSVPMessage
	setChoice(value FlowRSVPMessageChoiceEnum) FlowRSVPMessage
	// HasChoice checks if Choice has been set in FlowRSVPMessage
	HasChoice() bool
	// Path returns FlowRSVPPathMessage, set in FlowRSVPMessage.
	// FlowRSVPPathMessage is "Path" message requires the following list of objects in order as defined in https://www.rfc-editor.org/rfc/rfc3209.html#page-15: 1. SESSION 2. RSVP_HOP 3. TIME_VALUES 4. EXPLICIT_ROUTE [optional] 5. LABEL_REQUEST 6. SESSION_ATTRIBUTE [optional] 7. SENDER_TEMPLATE 8. SENDER_TSPEC 9. RECORD_ROUTE [optional]
	Path() FlowRSVPPathMessage
	// SetPath assigns FlowRSVPPathMessage provided by user to FlowRSVPMessage.
	// FlowRSVPPathMessage is "Path" message requires the following list of objects in order as defined in https://www.rfc-editor.org/rfc/rfc3209.html#page-15: 1. SESSION 2. RSVP_HOP 3. TIME_VALUES 4. EXPLICIT_ROUTE [optional] 5. LABEL_REQUEST 6. SESSION_ATTRIBUTE [optional] 7. SENDER_TEMPLATE 8. SENDER_TSPEC 9. RECORD_ROUTE [optional]
	SetPath(value FlowRSVPPathMessage) FlowRSVPMessage
	// HasPath checks if Path has been set in FlowRSVPMessage
	HasPath() bool
	setNil()
}

type FlowRSVPMessageChoiceEnum string

// Enum of Choice on FlowRSVPMessage
var FlowRSVPMessageChoice = struct {
	PATH FlowRSVPMessageChoiceEnum
}{
	PATH: FlowRSVPMessageChoiceEnum("path"),
}

func (obj *flowRSVPMessage) Choice() FlowRSVPMessageChoiceEnum {
	return FlowRSVPMessageChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowRSVPMessage) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPMessage) setChoice(value FlowRSVPMessageChoiceEnum) FlowRSVPMessage {
	intValue, ok := otg.FlowRSVPMessage_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPMessageChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPMessage_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Path = nil
	obj.pathHolder = nil

	if value == FlowRSVPMessageChoice.PATH {
		obj.obj.Path = NewFlowRSVPPathMessage().msg()
	}

	return obj
}

// description is TBD
// Path returns a FlowRSVPPathMessage
func (obj *flowRSVPMessage) Path() FlowRSVPPathMessage {
	if obj.obj.Path == nil {
		obj.setChoice(FlowRSVPMessageChoice.PATH)
	}
	if obj.pathHolder == nil {
		obj.pathHolder = &flowRSVPPathMessage{obj: obj.obj.Path}
	}
	return obj.pathHolder
}

// description is TBD
// Path returns a FlowRSVPPathMessage
func (obj *flowRSVPMessage) HasPath() bool {
	return obj.obj.Path != nil
}

// description is TBD
// SetPath sets the FlowRSVPPathMessage value in the FlowRSVPMessage object
func (obj *flowRSVPMessage) SetPath(value FlowRSVPPathMessage) FlowRSVPMessage {
	obj.setChoice(FlowRSVPMessageChoice.PATH)
	obj.pathHolder = nil
	obj.obj.Path = value.msg()

	return obj
}

func (obj *flowRSVPMessage) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Path != nil {

		obj.Path().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPMessage) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(FlowRSVPMessageChoice.PATH)

	}

}

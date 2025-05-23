package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolBgpGracefulRestartNotification *****
type actionProtocolBgpGracefulRestartNotification struct {
	validation
	obj                           *otg.ActionProtocolBgpGracefulRestartNotification
	marshaller                    marshalActionProtocolBgpGracefulRestartNotification
	unMarshaller                  unMarshalActionProtocolBgpGracefulRestartNotification
	ceaseHolder                   DeviceBgpCeaseError
	messageHeaderErrorHolder      DeviceBgpMessageHeaderError
	openMessageErrorHolder        DeviceBgpOpenMessageError
	updateMessageErrorHolder      DeviceBgpUpdateMessageError
	holdTimerExpiredHolder        DeviceBgpHoldTimerExpired
	finiteStateMachineErrorHolder DeviceBgpFiniteStateMachineError
	customHolder                  DeviceBgpCustomError
}

func NewActionProtocolBgpGracefulRestartNotification() ActionProtocolBgpGracefulRestartNotification {
	obj := actionProtocolBgpGracefulRestartNotification{obj: &otg.ActionProtocolBgpGracefulRestartNotification{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolBgpGracefulRestartNotification) msg() *otg.ActionProtocolBgpGracefulRestartNotification {
	return obj.obj
}

func (obj *actionProtocolBgpGracefulRestartNotification) setMsg(msg *otg.ActionProtocolBgpGracefulRestartNotification) ActionProtocolBgpGracefulRestartNotification {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolBgpGracefulRestartNotification struct {
	obj *actionProtocolBgpGracefulRestartNotification
}

type marshalActionProtocolBgpGracefulRestartNotification interface {
	// ToProto marshals ActionProtocolBgpGracefulRestartNotification to protobuf object *otg.ActionProtocolBgpGracefulRestartNotification
	ToProto() (*otg.ActionProtocolBgpGracefulRestartNotification, error)
	// ToPbText marshals ActionProtocolBgpGracefulRestartNotification to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolBgpGracefulRestartNotification to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolBgpGracefulRestartNotification to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals ActionProtocolBgpGracefulRestartNotification to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalactionProtocolBgpGracefulRestartNotification struct {
	obj *actionProtocolBgpGracefulRestartNotification
}

type unMarshalActionProtocolBgpGracefulRestartNotification interface {
	// FromProto unmarshals ActionProtocolBgpGracefulRestartNotification from protobuf object *otg.ActionProtocolBgpGracefulRestartNotification
	FromProto(msg *otg.ActionProtocolBgpGracefulRestartNotification) (ActionProtocolBgpGracefulRestartNotification, error)
	// FromPbText unmarshals ActionProtocolBgpGracefulRestartNotification from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolBgpGracefulRestartNotification from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolBgpGracefulRestartNotification from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolBgpGracefulRestartNotification) Marshal() marshalActionProtocolBgpGracefulRestartNotification {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolBgpGracefulRestartNotification{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolBgpGracefulRestartNotification) Unmarshal() unMarshalActionProtocolBgpGracefulRestartNotification {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolBgpGracefulRestartNotification{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolBgpGracefulRestartNotification) ToProto() (*otg.ActionProtocolBgpGracefulRestartNotification, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolBgpGracefulRestartNotification) FromProto(msg *otg.ActionProtocolBgpGracefulRestartNotification) (ActionProtocolBgpGracefulRestartNotification, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolBgpGracefulRestartNotification) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolBgpGracefulRestartNotification) FromPbText(value string) error {
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

func (m *marshalactionProtocolBgpGracefulRestartNotification) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolBgpGracefulRestartNotification) FromYaml(value string) error {
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

func (m *marshalactionProtocolBgpGracefulRestartNotification) ToJsonRaw() (string, error) {
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

func (m *marshalactionProtocolBgpGracefulRestartNotification) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolBgpGracefulRestartNotification) FromJson(value string) error {
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

func (obj *actionProtocolBgpGracefulRestartNotification) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolBgpGracefulRestartNotification) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolBgpGracefulRestartNotification) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolBgpGracefulRestartNotification) Clone() (ActionProtocolBgpGracefulRestartNotification, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolBgpGracefulRestartNotification()
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

func (obj *actionProtocolBgpGracefulRestartNotification) setNil() {
	obj.ceaseHolder = nil
	obj.messageHeaderErrorHolder = nil
	obj.openMessageErrorHolder = nil
	obj.updateMessageErrorHolder = nil
	obj.holdTimerExpiredHolder = nil
	obj.finiteStateMachineErrorHolder = nil
	obj.customHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionProtocolBgpGracefulRestartNotification is defines the explicit contents of the NOTIFICATION message to be sent when executing InitiateGracefulRestart trigger.  This causes the BGP connection to close.If a user wants to send custom Error Code and Error Subcode the custom object should be configured. A user can send IANA defined BGP NOTIFICATIONs according to https://www.iana.org/assignments/bgp-parameters/bgp-parameters.xhtml.
type ActionProtocolBgpGracefulRestartNotification interface {
	Validation
	// msg marshals ActionProtocolBgpGracefulRestartNotification to protobuf object *otg.ActionProtocolBgpGracefulRestartNotification
	// and doesn't set defaults
	msg() *otg.ActionProtocolBgpGracefulRestartNotification
	// setMsg unmarshals ActionProtocolBgpGracefulRestartNotification from protobuf object *otg.ActionProtocolBgpGracefulRestartNotification
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolBgpGracefulRestartNotification) ActionProtocolBgpGracefulRestartNotification
	// provides marshal interface
	Marshal() marshalActionProtocolBgpGracefulRestartNotification
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolBgpGracefulRestartNotification
	// validate validates ActionProtocolBgpGracefulRestartNotification
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolBgpGracefulRestartNotification, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ActionProtocolBgpGracefulRestartNotificationChoiceEnum, set in ActionProtocolBgpGracefulRestartNotification
	Choice() ActionProtocolBgpGracefulRestartNotificationChoiceEnum
	// setChoice assigns ActionProtocolBgpGracefulRestartNotificationChoiceEnum provided by user to ActionProtocolBgpGracefulRestartNotification
	setChoice(value ActionProtocolBgpGracefulRestartNotificationChoiceEnum) ActionProtocolBgpGracefulRestartNotification
	// HasChoice checks if Choice has been set in ActionProtocolBgpGracefulRestartNotification
	HasChoice() bool
	// Cease returns DeviceBgpCeaseError, set in ActionProtocolBgpGracefulRestartNotification.
	// DeviceBgpCeaseError is in the absence of any fatal errors, a BGP peer can close its BGP connection by sending the NOTIFICATION message with the  Error Code Cease. The 'hard_reset_code6_subcode9' subcode for Cease Notification can be used to signal a hard reset that will indicate that  Graceful Restart cannot be performed, even when Notification extensions to Graceful Restart procedure is supported.
	Cease() DeviceBgpCeaseError
	// SetCease assigns DeviceBgpCeaseError provided by user to ActionProtocolBgpGracefulRestartNotification.
	// DeviceBgpCeaseError is in the absence of any fatal errors, a BGP peer can close its BGP connection by sending the NOTIFICATION message with the  Error Code Cease. The 'hard_reset_code6_subcode9' subcode for Cease Notification can be used to signal a hard reset that will indicate that  Graceful Restart cannot be performed, even when Notification extensions to Graceful Restart procedure is supported.
	SetCease(value DeviceBgpCeaseError) ActionProtocolBgpGracefulRestartNotification
	// HasCease checks if Cease has been set in ActionProtocolBgpGracefulRestartNotification
	HasCease() bool
	// MessageHeaderError returns DeviceBgpMessageHeaderError, set in ActionProtocolBgpGracefulRestartNotification.
	// DeviceBgpMessageHeaderError is all errors detected while processing the Message Header are indicated by sending the NOTIFICATION message  with the Error Code-Message Header Error. The Error Subcode elaborates on the specific nature of the error.
	MessageHeaderError() DeviceBgpMessageHeaderError
	// SetMessageHeaderError assigns DeviceBgpMessageHeaderError provided by user to ActionProtocolBgpGracefulRestartNotification.
	// DeviceBgpMessageHeaderError is all errors detected while processing the Message Header are indicated by sending the NOTIFICATION message  with the Error Code-Message Header Error. The Error Subcode elaborates on the specific nature of the error.
	SetMessageHeaderError(value DeviceBgpMessageHeaderError) ActionProtocolBgpGracefulRestartNotification
	// HasMessageHeaderError checks if MessageHeaderError has been set in ActionProtocolBgpGracefulRestartNotification
	HasMessageHeaderError() bool
	// OpenMessageError returns DeviceBgpOpenMessageError, set in ActionProtocolBgpGracefulRestartNotification.
	// DeviceBgpOpenMessageError is all errors detected while processing the OPEN message are indicated by sending the NOTIFICATION message  with the Error Code-Open Message Error. The Error Subcode elaborates on the specific nature of the error.
	OpenMessageError() DeviceBgpOpenMessageError
	// SetOpenMessageError assigns DeviceBgpOpenMessageError provided by user to ActionProtocolBgpGracefulRestartNotification.
	// DeviceBgpOpenMessageError is all errors detected while processing the OPEN message are indicated by sending the NOTIFICATION message  with the Error Code-Open Message Error. The Error Subcode elaborates on the specific nature of the error.
	SetOpenMessageError(value DeviceBgpOpenMessageError) ActionProtocolBgpGracefulRestartNotification
	// HasOpenMessageError checks if OpenMessageError has been set in ActionProtocolBgpGracefulRestartNotification
	HasOpenMessageError() bool
	// UpdateMessageError returns DeviceBgpUpdateMessageError, set in ActionProtocolBgpGracefulRestartNotification.
	// DeviceBgpUpdateMessageError is all errors detected while processing the UPDATE message are indicated by sending the NOTIFICATION message  with the Error Code-Update Message Error. The Error Subcode elaborates on the specific nature of the error.
	UpdateMessageError() DeviceBgpUpdateMessageError
	// SetUpdateMessageError assigns DeviceBgpUpdateMessageError provided by user to ActionProtocolBgpGracefulRestartNotification.
	// DeviceBgpUpdateMessageError is all errors detected while processing the UPDATE message are indicated by sending the NOTIFICATION message  with the Error Code-Update Message Error. The Error Subcode elaborates on the specific nature of the error.
	SetUpdateMessageError(value DeviceBgpUpdateMessageError) ActionProtocolBgpGracefulRestartNotification
	// HasUpdateMessageError checks if UpdateMessageError has been set in ActionProtocolBgpGracefulRestartNotification
	HasUpdateMessageError() bool
	// HoldTimerExpired returns DeviceBgpHoldTimerExpired, set in ActionProtocolBgpGracefulRestartNotification.
	// DeviceBgpHoldTimerExpired is if a system does not receive successive KEEPALIVE, UPDATE, and/or NOTIFICATION messages within the period specified  in the Hold Time field of the OPEN message, then the NOTIFICATION message with the Hold Timer Expired Error Code(Error Code 4) is  sent and the BGP connection is closed. The Sub Code used is 0. If a user wants to use non zero Sub Code then CustomError can be used.
	HoldTimerExpired() DeviceBgpHoldTimerExpired
	// SetHoldTimerExpired assigns DeviceBgpHoldTimerExpired provided by user to ActionProtocolBgpGracefulRestartNotification.
	// DeviceBgpHoldTimerExpired is if a system does not receive successive KEEPALIVE, UPDATE, and/or NOTIFICATION messages within the period specified  in the Hold Time field of the OPEN message, then the NOTIFICATION message with the Hold Timer Expired Error Code(Error Code 4) is  sent and the BGP connection is closed. The Sub Code used is 0. If a user wants to use non zero Sub Code then CustomError can be used.
	SetHoldTimerExpired(value DeviceBgpHoldTimerExpired) ActionProtocolBgpGracefulRestartNotification
	// HasHoldTimerExpired checks if HoldTimerExpired has been set in ActionProtocolBgpGracefulRestartNotification
	HasHoldTimerExpired() bool
	// FiniteStateMachineError returns DeviceBgpFiniteStateMachineError, set in ActionProtocolBgpGracefulRestartNotification.
	// DeviceBgpFiniteStateMachineError is any error detected by the BGP Finite State Machine (e.g., receipt of an unexpected event) is indicated by  sending the NOTIFICATION message with the Error Code-Finite State Machine Error(Error Code 5). The Sub Code used is 0.  If a user wants to use non zero Sub Code then CustomError can be used.
	FiniteStateMachineError() DeviceBgpFiniteStateMachineError
	// SetFiniteStateMachineError assigns DeviceBgpFiniteStateMachineError provided by user to ActionProtocolBgpGracefulRestartNotification.
	// DeviceBgpFiniteStateMachineError is any error detected by the BGP Finite State Machine (e.g., receipt of an unexpected event) is indicated by  sending the NOTIFICATION message with the Error Code-Finite State Machine Error(Error Code 5). The Sub Code used is 0.  If a user wants to use non zero Sub Code then CustomError can be used.
	SetFiniteStateMachineError(value DeviceBgpFiniteStateMachineError) ActionProtocolBgpGracefulRestartNotification
	// HasFiniteStateMachineError checks if FiniteStateMachineError has been set in ActionProtocolBgpGracefulRestartNotification
	HasFiniteStateMachineError() bool
	// Custom returns DeviceBgpCustomError, set in ActionProtocolBgpGracefulRestartNotification.
	// DeviceBgpCustomError is a BGP peer can send NOTIFICATION message with user defined Error Code and Error Subcode.
	Custom() DeviceBgpCustomError
	// SetCustom assigns DeviceBgpCustomError provided by user to ActionProtocolBgpGracefulRestartNotification.
	// DeviceBgpCustomError is a BGP peer can send NOTIFICATION message with user defined Error Code and Error Subcode.
	SetCustom(value DeviceBgpCustomError) ActionProtocolBgpGracefulRestartNotification
	// HasCustom checks if Custom has been set in ActionProtocolBgpGracefulRestartNotification
	HasCustom() bool
	setNil()
}

type ActionProtocolBgpGracefulRestartNotificationChoiceEnum string

// Enum of Choice on ActionProtocolBgpGracefulRestartNotification
var ActionProtocolBgpGracefulRestartNotificationChoice = struct {
	CEASE                      ActionProtocolBgpGracefulRestartNotificationChoiceEnum
	MESSAGE_HEADER_ERROR       ActionProtocolBgpGracefulRestartNotificationChoiceEnum
	OPEN_MESSAGE_ERROR         ActionProtocolBgpGracefulRestartNotificationChoiceEnum
	UPDATE_MESSAGE_ERROR       ActionProtocolBgpGracefulRestartNotificationChoiceEnum
	HOLD_TIMER_EXPIRED         ActionProtocolBgpGracefulRestartNotificationChoiceEnum
	FINITE_STATE_MACHINE_ERROR ActionProtocolBgpGracefulRestartNotificationChoiceEnum
	CUSTOM                     ActionProtocolBgpGracefulRestartNotificationChoiceEnum
}{
	CEASE:                      ActionProtocolBgpGracefulRestartNotificationChoiceEnum("cease"),
	MESSAGE_HEADER_ERROR:       ActionProtocolBgpGracefulRestartNotificationChoiceEnum("message_header_error"),
	OPEN_MESSAGE_ERROR:         ActionProtocolBgpGracefulRestartNotificationChoiceEnum("open_message_error"),
	UPDATE_MESSAGE_ERROR:       ActionProtocolBgpGracefulRestartNotificationChoiceEnum("update_message_error"),
	HOLD_TIMER_EXPIRED:         ActionProtocolBgpGracefulRestartNotificationChoiceEnum("hold_timer_expired"),
	FINITE_STATE_MACHINE_ERROR: ActionProtocolBgpGracefulRestartNotificationChoiceEnum("finite_state_machine_error"),
	CUSTOM:                     ActionProtocolBgpGracefulRestartNotificationChoiceEnum("custom"),
}

func (obj *actionProtocolBgpGracefulRestartNotification) Choice() ActionProtocolBgpGracefulRestartNotificationChoiceEnum {
	return ActionProtocolBgpGracefulRestartNotificationChoiceEnum(obj.obj.Choice.Enum().String())
}

// Each BGP NOTIFICATION message includes an Error Code field indicating what type of problem occurred. For certain Error Codes, an Error  Subcode field provides additional details about the specific nature of the problem.  The choice value will provide the Error Code used in NOTIFICATION message.  The Subcode can be set for each of the corresponding errors except for Hold Timer Expired error and BGP Finite State Machine error.  In both of these cases Subcode 0 will be sent. If a user wants to use non zero Sub Code then custom choice can be used.
// Choice returns a string
func (obj *actionProtocolBgpGracefulRestartNotification) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *actionProtocolBgpGracefulRestartNotification) setChoice(value ActionProtocolBgpGracefulRestartNotificationChoiceEnum) ActionProtocolBgpGracefulRestartNotification {
	intValue, ok := otg.ActionProtocolBgpGracefulRestartNotification_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ActionProtocolBgpGracefulRestartNotificationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ActionProtocolBgpGracefulRestartNotification_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.customHolder = nil
	obj.obj.FiniteStateMachineError = nil
	obj.finiteStateMachineErrorHolder = nil
	obj.obj.HoldTimerExpired = nil
	obj.holdTimerExpiredHolder = nil
	obj.obj.UpdateMessageError = nil
	obj.updateMessageErrorHolder = nil
	obj.obj.OpenMessageError = nil
	obj.openMessageErrorHolder = nil
	obj.obj.MessageHeaderError = nil
	obj.messageHeaderErrorHolder = nil
	obj.obj.Cease = nil
	obj.ceaseHolder = nil

	if value == ActionProtocolBgpGracefulRestartNotificationChoice.CEASE {
		obj.obj.Cease = NewDeviceBgpCeaseError().msg()
	}

	if value == ActionProtocolBgpGracefulRestartNotificationChoice.MESSAGE_HEADER_ERROR {
		obj.obj.MessageHeaderError = NewDeviceBgpMessageHeaderError().msg()
	}

	if value == ActionProtocolBgpGracefulRestartNotificationChoice.OPEN_MESSAGE_ERROR {
		obj.obj.OpenMessageError = NewDeviceBgpOpenMessageError().msg()
	}

	if value == ActionProtocolBgpGracefulRestartNotificationChoice.UPDATE_MESSAGE_ERROR {
		obj.obj.UpdateMessageError = NewDeviceBgpUpdateMessageError().msg()
	}

	if value == ActionProtocolBgpGracefulRestartNotificationChoice.HOLD_TIMER_EXPIRED {
		obj.obj.HoldTimerExpired = NewDeviceBgpHoldTimerExpired().msg()
	}

	if value == ActionProtocolBgpGracefulRestartNotificationChoice.FINITE_STATE_MACHINE_ERROR {
		obj.obj.FiniteStateMachineError = NewDeviceBgpFiniteStateMachineError().msg()
	}

	if value == ActionProtocolBgpGracefulRestartNotificationChoice.CUSTOM {
		obj.obj.Custom = NewDeviceBgpCustomError().msg()
	}

	return obj
}

// description is TBD
// Cease returns a DeviceBgpCeaseError
func (obj *actionProtocolBgpGracefulRestartNotification) Cease() DeviceBgpCeaseError {
	if obj.obj.Cease == nil {
		obj.setChoice(ActionProtocolBgpGracefulRestartNotificationChoice.CEASE)
	}
	if obj.ceaseHolder == nil {
		obj.ceaseHolder = &deviceBgpCeaseError{obj: obj.obj.Cease}
	}
	return obj.ceaseHolder
}

// description is TBD
// Cease returns a DeviceBgpCeaseError
func (obj *actionProtocolBgpGracefulRestartNotification) HasCease() bool {
	return obj.obj.Cease != nil
}

// description is TBD
// SetCease sets the DeviceBgpCeaseError value in the ActionProtocolBgpGracefulRestartNotification object
func (obj *actionProtocolBgpGracefulRestartNotification) SetCease(value DeviceBgpCeaseError) ActionProtocolBgpGracefulRestartNotification {
	obj.setChoice(ActionProtocolBgpGracefulRestartNotificationChoice.CEASE)
	obj.ceaseHolder = nil
	obj.obj.Cease = value.msg()

	return obj
}

// description is TBD
// MessageHeaderError returns a DeviceBgpMessageHeaderError
func (obj *actionProtocolBgpGracefulRestartNotification) MessageHeaderError() DeviceBgpMessageHeaderError {
	if obj.obj.MessageHeaderError == nil {
		obj.setChoice(ActionProtocolBgpGracefulRestartNotificationChoice.MESSAGE_HEADER_ERROR)
	}
	if obj.messageHeaderErrorHolder == nil {
		obj.messageHeaderErrorHolder = &deviceBgpMessageHeaderError{obj: obj.obj.MessageHeaderError}
	}
	return obj.messageHeaderErrorHolder
}

// description is TBD
// MessageHeaderError returns a DeviceBgpMessageHeaderError
func (obj *actionProtocolBgpGracefulRestartNotification) HasMessageHeaderError() bool {
	return obj.obj.MessageHeaderError != nil
}

// description is TBD
// SetMessageHeaderError sets the DeviceBgpMessageHeaderError value in the ActionProtocolBgpGracefulRestartNotification object
func (obj *actionProtocolBgpGracefulRestartNotification) SetMessageHeaderError(value DeviceBgpMessageHeaderError) ActionProtocolBgpGracefulRestartNotification {
	obj.setChoice(ActionProtocolBgpGracefulRestartNotificationChoice.MESSAGE_HEADER_ERROR)
	obj.messageHeaderErrorHolder = nil
	obj.obj.MessageHeaderError = value.msg()

	return obj
}

// description is TBD
// OpenMessageError returns a DeviceBgpOpenMessageError
func (obj *actionProtocolBgpGracefulRestartNotification) OpenMessageError() DeviceBgpOpenMessageError {
	if obj.obj.OpenMessageError == nil {
		obj.setChoice(ActionProtocolBgpGracefulRestartNotificationChoice.OPEN_MESSAGE_ERROR)
	}
	if obj.openMessageErrorHolder == nil {
		obj.openMessageErrorHolder = &deviceBgpOpenMessageError{obj: obj.obj.OpenMessageError}
	}
	return obj.openMessageErrorHolder
}

// description is TBD
// OpenMessageError returns a DeviceBgpOpenMessageError
func (obj *actionProtocolBgpGracefulRestartNotification) HasOpenMessageError() bool {
	return obj.obj.OpenMessageError != nil
}

// description is TBD
// SetOpenMessageError sets the DeviceBgpOpenMessageError value in the ActionProtocolBgpGracefulRestartNotification object
func (obj *actionProtocolBgpGracefulRestartNotification) SetOpenMessageError(value DeviceBgpOpenMessageError) ActionProtocolBgpGracefulRestartNotification {
	obj.setChoice(ActionProtocolBgpGracefulRestartNotificationChoice.OPEN_MESSAGE_ERROR)
	obj.openMessageErrorHolder = nil
	obj.obj.OpenMessageError = value.msg()

	return obj
}

// description is TBD
// UpdateMessageError returns a DeviceBgpUpdateMessageError
func (obj *actionProtocolBgpGracefulRestartNotification) UpdateMessageError() DeviceBgpUpdateMessageError {
	if obj.obj.UpdateMessageError == nil {
		obj.setChoice(ActionProtocolBgpGracefulRestartNotificationChoice.UPDATE_MESSAGE_ERROR)
	}
	if obj.updateMessageErrorHolder == nil {
		obj.updateMessageErrorHolder = &deviceBgpUpdateMessageError{obj: obj.obj.UpdateMessageError}
	}
	return obj.updateMessageErrorHolder
}

// description is TBD
// UpdateMessageError returns a DeviceBgpUpdateMessageError
func (obj *actionProtocolBgpGracefulRestartNotification) HasUpdateMessageError() bool {
	return obj.obj.UpdateMessageError != nil
}

// description is TBD
// SetUpdateMessageError sets the DeviceBgpUpdateMessageError value in the ActionProtocolBgpGracefulRestartNotification object
func (obj *actionProtocolBgpGracefulRestartNotification) SetUpdateMessageError(value DeviceBgpUpdateMessageError) ActionProtocolBgpGracefulRestartNotification {
	obj.setChoice(ActionProtocolBgpGracefulRestartNotificationChoice.UPDATE_MESSAGE_ERROR)
	obj.updateMessageErrorHolder = nil
	obj.obj.UpdateMessageError = value.msg()

	return obj
}

// description is TBD
// HoldTimerExpired returns a DeviceBgpHoldTimerExpired
func (obj *actionProtocolBgpGracefulRestartNotification) HoldTimerExpired() DeviceBgpHoldTimerExpired {
	if obj.obj.HoldTimerExpired == nil {
		obj.setChoice(ActionProtocolBgpGracefulRestartNotificationChoice.HOLD_TIMER_EXPIRED)
	}
	if obj.holdTimerExpiredHolder == nil {
		obj.holdTimerExpiredHolder = &deviceBgpHoldTimerExpired{obj: obj.obj.HoldTimerExpired}
	}
	return obj.holdTimerExpiredHolder
}

// description is TBD
// HoldTimerExpired returns a DeviceBgpHoldTimerExpired
func (obj *actionProtocolBgpGracefulRestartNotification) HasHoldTimerExpired() bool {
	return obj.obj.HoldTimerExpired != nil
}

// description is TBD
// SetHoldTimerExpired sets the DeviceBgpHoldTimerExpired value in the ActionProtocolBgpGracefulRestartNotification object
func (obj *actionProtocolBgpGracefulRestartNotification) SetHoldTimerExpired(value DeviceBgpHoldTimerExpired) ActionProtocolBgpGracefulRestartNotification {
	obj.setChoice(ActionProtocolBgpGracefulRestartNotificationChoice.HOLD_TIMER_EXPIRED)
	obj.holdTimerExpiredHolder = nil
	obj.obj.HoldTimerExpired = value.msg()

	return obj
}

// description is TBD
// FiniteStateMachineError returns a DeviceBgpFiniteStateMachineError
func (obj *actionProtocolBgpGracefulRestartNotification) FiniteStateMachineError() DeviceBgpFiniteStateMachineError {
	if obj.obj.FiniteStateMachineError == nil {
		obj.setChoice(ActionProtocolBgpGracefulRestartNotificationChoice.FINITE_STATE_MACHINE_ERROR)
	}
	if obj.finiteStateMachineErrorHolder == nil {
		obj.finiteStateMachineErrorHolder = &deviceBgpFiniteStateMachineError{obj: obj.obj.FiniteStateMachineError}
	}
	return obj.finiteStateMachineErrorHolder
}

// description is TBD
// FiniteStateMachineError returns a DeviceBgpFiniteStateMachineError
func (obj *actionProtocolBgpGracefulRestartNotification) HasFiniteStateMachineError() bool {
	return obj.obj.FiniteStateMachineError != nil
}

// description is TBD
// SetFiniteStateMachineError sets the DeviceBgpFiniteStateMachineError value in the ActionProtocolBgpGracefulRestartNotification object
func (obj *actionProtocolBgpGracefulRestartNotification) SetFiniteStateMachineError(value DeviceBgpFiniteStateMachineError) ActionProtocolBgpGracefulRestartNotification {
	obj.setChoice(ActionProtocolBgpGracefulRestartNotificationChoice.FINITE_STATE_MACHINE_ERROR)
	obj.finiteStateMachineErrorHolder = nil
	obj.obj.FiniteStateMachineError = value.msg()

	return obj
}

// description is TBD
// Custom returns a DeviceBgpCustomError
func (obj *actionProtocolBgpGracefulRestartNotification) Custom() DeviceBgpCustomError {
	if obj.obj.Custom == nil {
		obj.setChoice(ActionProtocolBgpGracefulRestartNotificationChoice.CUSTOM)
	}
	if obj.customHolder == nil {
		obj.customHolder = &deviceBgpCustomError{obj: obj.obj.Custom}
	}
	return obj.customHolder
}

// description is TBD
// Custom returns a DeviceBgpCustomError
func (obj *actionProtocolBgpGracefulRestartNotification) HasCustom() bool {
	return obj.obj.Custom != nil
}

// description is TBD
// SetCustom sets the DeviceBgpCustomError value in the ActionProtocolBgpGracefulRestartNotification object
func (obj *actionProtocolBgpGracefulRestartNotification) SetCustom(value DeviceBgpCustomError) ActionProtocolBgpGracefulRestartNotification {
	obj.setChoice(ActionProtocolBgpGracefulRestartNotificationChoice.CUSTOM)
	obj.customHolder = nil
	obj.obj.Custom = value.msg()

	return obj
}

func (obj *actionProtocolBgpGracefulRestartNotification) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Cease != nil {

		obj.Cease().validateObj(vObj, set_default)
	}

	if obj.obj.MessageHeaderError != nil {

		obj.MessageHeaderError().validateObj(vObj, set_default)
	}

	if obj.obj.OpenMessageError != nil {

		obj.OpenMessageError().validateObj(vObj, set_default)
	}

	if obj.obj.UpdateMessageError != nil {

		obj.UpdateMessageError().validateObj(vObj, set_default)
	}

	if obj.obj.HoldTimerExpired != nil {

		obj.HoldTimerExpired().validateObj(vObj, set_default)
	}

	if obj.obj.FiniteStateMachineError != nil {

		obj.FiniteStateMachineError().validateObj(vObj, set_default)
	}

	if obj.obj.Custom != nil {

		obj.Custom().validateObj(vObj, set_default)
	}

}

func (obj *actionProtocolBgpGracefulRestartNotification) setDefault() {
	var choices_set int = 0
	var choice ActionProtocolBgpGracefulRestartNotificationChoiceEnum

	if obj.obj.Cease != nil {
		choices_set += 1
		choice = ActionProtocolBgpGracefulRestartNotificationChoice.CEASE
	}

	if obj.obj.MessageHeaderError != nil {
		choices_set += 1
		choice = ActionProtocolBgpGracefulRestartNotificationChoice.MESSAGE_HEADER_ERROR
	}

	if obj.obj.OpenMessageError != nil {
		choices_set += 1
		choice = ActionProtocolBgpGracefulRestartNotificationChoice.OPEN_MESSAGE_ERROR
	}

	if obj.obj.UpdateMessageError != nil {
		choices_set += 1
		choice = ActionProtocolBgpGracefulRestartNotificationChoice.UPDATE_MESSAGE_ERROR
	}

	if obj.obj.HoldTimerExpired != nil {
		choices_set += 1
		choice = ActionProtocolBgpGracefulRestartNotificationChoice.HOLD_TIMER_EXPIRED
	}

	if obj.obj.FiniteStateMachineError != nil {
		choices_set += 1
		choice = ActionProtocolBgpGracefulRestartNotificationChoice.FINITE_STATE_MACHINE_ERROR
	}

	if obj.obj.Custom != nil {
		choices_set += 1
		choice = ActionProtocolBgpGracefulRestartNotificationChoice.CUSTOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(ActionProtocolBgpGracefulRestartNotificationChoice.CEASE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ActionProtocolBgpGracefulRestartNotification")
			}
		} else {
			intVal := otg.ActionProtocolBgpGracefulRestartNotification_Choice_Enum_value[string(choice)]
			enumValue := otg.ActionProtocolBgpGracefulRestartNotification_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

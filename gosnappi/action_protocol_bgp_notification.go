package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolBgpNotification *****
type actionProtocolBgpNotification struct {
	validation
	obj                           *otg.ActionProtocolBgpNotification
	marshaller                    marshalActionProtocolBgpNotification
	unMarshaller                  unMarshalActionProtocolBgpNotification
	ceaseHolder                   DeviceBgpCeaseError
	messageHeaderErrorHolder      DeviceBgpMessageHeaderError
	openMessageErrorHolder        DeviceBgpOpenMessageError
	updateMessageErrorHolder      DeviceBgpUpdateMessageError
	holdTimerExpiredHolder        DeviceBgpHoldTimerExpired
	finiteStateMachineErrorHolder DeviceBgpFiniteStateMachineError
	customHolder                  DeviceBgpCustomError
}

func NewActionProtocolBgpNotification() ActionProtocolBgpNotification {
	obj := actionProtocolBgpNotification{obj: &otg.ActionProtocolBgpNotification{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolBgpNotification) msg() *otg.ActionProtocolBgpNotification {
	return obj.obj
}

func (obj *actionProtocolBgpNotification) setMsg(msg *otg.ActionProtocolBgpNotification) ActionProtocolBgpNotification {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolBgpNotification struct {
	obj *actionProtocolBgpNotification
}

type marshalActionProtocolBgpNotification interface {
	// ToProto marshals ActionProtocolBgpNotification to protobuf object *otg.ActionProtocolBgpNotification
	ToProto() (*otg.ActionProtocolBgpNotification, error)
	// ToPbText marshals ActionProtocolBgpNotification to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolBgpNotification to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolBgpNotification to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolBgpNotification struct {
	obj *actionProtocolBgpNotification
}

type unMarshalActionProtocolBgpNotification interface {
	// FromProto unmarshals ActionProtocolBgpNotification from protobuf object *otg.ActionProtocolBgpNotification
	FromProto(msg *otg.ActionProtocolBgpNotification) (ActionProtocolBgpNotification, error)
	// FromPbText unmarshals ActionProtocolBgpNotification from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolBgpNotification from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolBgpNotification from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolBgpNotification) Marshal() marshalActionProtocolBgpNotification {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolBgpNotification{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolBgpNotification) Unmarshal() unMarshalActionProtocolBgpNotification {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolBgpNotification{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolBgpNotification) ToProto() (*otg.ActionProtocolBgpNotification, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolBgpNotification) FromProto(msg *otg.ActionProtocolBgpNotification) (ActionProtocolBgpNotification, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolBgpNotification) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolBgpNotification) FromPbText(value string) error {
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

func (m *marshalactionProtocolBgpNotification) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolBgpNotification) FromYaml(value string) error {
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

func (m *marshalactionProtocolBgpNotification) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolBgpNotification) FromJson(value string) error {
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

func (obj *actionProtocolBgpNotification) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolBgpNotification) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolBgpNotification) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolBgpNotification) Clone() (ActionProtocolBgpNotification, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolBgpNotification()
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

func (obj *actionProtocolBgpNotification) setNil() {
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

// ActionProtocolBgpNotification is a NOTIFICATION message is sent when an error is detected with the BGP session, such as hold timer expiring, misconfigured AS number  or a BGP session reset is requested. This causes the BGP connection to close. Send explicit NOTIFICATIONs for list of specified  BGP peers. If a user wants to send custom Error Code and Error Subcode the custom object should be configured. A user can send IANA defined BGP NOTIFICATIONs according to https://www.iana.org/assignments/bgp-parameters/bgp-parameters.xhtml.
type ActionProtocolBgpNotification interface {
	Validation
	// msg marshals ActionProtocolBgpNotification to protobuf object *otg.ActionProtocolBgpNotification
	// and doesn't set defaults
	msg() *otg.ActionProtocolBgpNotification
	// setMsg unmarshals ActionProtocolBgpNotification from protobuf object *otg.ActionProtocolBgpNotification
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolBgpNotification) ActionProtocolBgpNotification
	// provides marshal interface
	Marshal() marshalActionProtocolBgpNotification
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolBgpNotification
	// validate validates ActionProtocolBgpNotification
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolBgpNotification, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Names returns []string, set in ActionProtocolBgpNotification.
	Names() []string
	// SetNames assigns []string provided by user to ActionProtocolBgpNotification
	SetNames(value []string) ActionProtocolBgpNotification
	// Choice returns ActionProtocolBgpNotificationChoiceEnum, set in ActionProtocolBgpNotification
	Choice() ActionProtocolBgpNotificationChoiceEnum
	// setChoice assigns ActionProtocolBgpNotificationChoiceEnum provided by user to ActionProtocolBgpNotification
	setChoice(value ActionProtocolBgpNotificationChoiceEnum) ActionProtocolBgpNotification
	// HasChoice checks if Choice has been set in ActionProtocolBgpNotification
	HasChoice() bool
	// Cease returns DeviceBgpCeaseError, set in ActionProtocolBgpNotification.
	// DeviceBgpCeaseError is in the absence of any fatal errors, a BGP peer can close its BGP connection by sending the NOTIFICATION message with the  Error Code Cease. The 'hard_reset_code6_subcode9' subcode for Cease Notification can be used to signal a hard reset that will indicate that  Graceful Restart cannot be performed, even when Notification extensions to Graceful Restart procedure is supported.
	Cease() DeviceBgpCeaseError
	// SetCease assigns DeviceBgpCeaseError provided by user to ActionProtocolBgpNotification.
	// DeviceBgpCeaseError is in the absence of any fatal errors, a BGP peer can close its BGP connection by sending the NOTIFICATION message with the  Error Code Cease. The 'hard_reset_code6_subcode9' subcode for Cease Notification can be used to signal a hard reset that will indicate that  Graceful Restart cannot be performed, even when Notification extensions to Graceful Restart procedure is supported.
	SetCease(value DeviceBgpCeaseError) ActionProtocolBgpNotification
	// HasCease checks if Cease has been set in ActionProtocolBgpNotification
	HasCease() bool
	// MessageHeaderError returns DeviceBgpMessageHeaderError, set in ActionProtocolBgpNotification.
	// DeviceBgpMessageHeaderError is all errors detected while processing the Message Header are indicated by sending the NOTIFICATION message  with the Error Code-Message Header Error. The Error Subcode elaborates on the specific nature of the error.
	MessageHeaderError() DeviceBgpMessageHeaderError
	// SetMessageHeaderError assigns DeviceBgpMessageHeaderError provided by user to ActionProtocolBgpNotification.
	// DeviceBgpMessageHeaderError is all errors detected while processing the Message Header are indicated by sending the NOTIFICATION message  with the Error Code-Message Header Error. The Error Subcode elaborates on the specific nature of the error.
	SetMessageHeaderError(value DeviceBgpMessageHeaderError) ActionProtocolBgpNotification
	// HasMessageHeaderError checks if MessageHeaderError has been set in ActionProtocolBgpNotification
	HasMessageHeaderError() bool
	// OpenMessageError returns DeviceBgpOpenMessageError, set in ActionProtocolBgpNotification.
	// DeviceBgpOpenMessageError is all errors detected while processing the OPEN message are indicated by sending the NOTIFICATION message  with the Error Code-Open Message Error. The Error Subcode elaborates on the specific nature of the error.
	OpenMessageError() DeviceBgpOpenMessageError
	// SetOpenMessageError assigns DeviceBgpOpenMessageError provided by user to ActionProtocolBgpNotification.
	// DeviceBgpOpenMessageError is all errors detected while processing the OPEN message are indicated by sending the NOTIFICATION message  with the Error Code-Open Message Error. The Error Subcode elaborates on the specific nature of the error.
	SetOpenMessageError(value DeviceBgpOpenMessageError) ActionProtocolBgpNotification
	// HasOpenMessageError checks if OpenMessageError has been set in ActionProtocolBgpNotification
	HasOpenMessageError() bool
	// UpdateMessageError returns DeviceBgpUpdateMessageError, set in ActionProtocolBgpNotification.
	// DeviceBgpUpdateMessageError is all errors detected while processing the UPDATE message are indicated by sending the NOTIFICATION message  with the Error Code-Update Message Error. The Error Subcode elaborates on the specific nature of the error.
	UpdateMessageError() DeviceBgpUpdateMessageError
	// SetUpdateMessageError assigns DeviceBgpUpdateMessageError provided by user to ActionProtocolBgpNotification.
	// DeviceBgpUpdateMessageError is all errors detected while processing the UPDATE message are indicated by sending the NOTIFICATION message  with the Error Code-Update Message Error. The Error Subcode elaborates on the specific nature of the error.
	SetUpdateMessageError(value DeviceBgpUpdateMessageError) ActionProtocolBgpNotification
	// HasUpdateMessageError checks if UpdateMessageError has been set in ActionProtocolBgpNotification
	HasUpdateMessageError() bool
	// HoldTimerExpired returns DeviceBgpHoldTimerExpired, set in ActionProtocolBgpNotification.
	// DeviceBgpHoldTimerExpired is if a system does not receive successive KEEPALIVE, UPDATE, and/or NOTIFICATION messages within the period specified  in the Hold Time field of the OPEN message, then the NOTIFICATION message with the Hold Timer Expired Error Code(Error Code 4) is  sent and the BGP connection is closed. The Sub Code used is 0. If a user wants to use non zero Sub Code then CustomError can be used.
	HoldTimerExpired() DeviceBgpHoldTimerExpired
	// SetHoldTimerExpired assigns DeviceBgpHoldTimerExpired provided by user to ActionProtocolBgpNotification.
	// DeviceBgpHoldTimerExpired is if a system does not receive successive KEEPALIVE, UPDATE, and/or NOTIFICATION messages within the period specified  in the Hold Time field of the OPEN message, then the NOTIFICATION message with the Hold Timer Expired Error Code(Error Code 4) is  sent and the BGP connection is closed. The Sub Code used is 0. If a user wants to use non zero Sub Code then CustomError can be used.
	SetHoldTimerExpired(value DeviceBgpHoldTimerExpired) ActionProtocolBgpNotification
	// HasHoldTimerExpired checks if HoldTimerExpired has been set in ActionProtocolBgpNotification
	HasHoldTimerExpired() bool
	// FiniteStateMachineError returns DeviceBgpFiniteStateMachineError, set in ActionProtocolBgpNotification.
	// DeviceBgpFiniteStateMachineError is any error detected by the BGP Finite State Machine (e.g., receipt of an unexpected event) is indicated by  sending the NOTIFICATION message with the Error Code-Finite State Machine Error(Error Code 5). The Sub Code used is 0.  If a user wants to use non zero Sub Code then CustomError can be used.
	FiniteStateMachineError() DeviceBgpFiniteStateMachineError
	// SetFiniteStateMachineError assigns DeviceBgpFiniteStateMachineError provided by user to ActionProtocolBgpNotification.
	// DeviceBgpFiniteStateMachineError is any error detected by the BGP Finite State Machine (e.g., receipt of an unexpected event) is indicated by  sending the NOTIFICATION message with the Error Code-Finite State Machine Error(Error Code 5). The Sub Code used is 0.  If a user wants to use non zero Sub Code then CustomError can be used.
	SetFiniteStateMachineError(value DeviceBgpFiniteStateMachineError) ActionProtocolBgpNotification
	// HasFiniteStateMachineError checks if FiniteStateMachineError has been set in ActionProtocolBgpNotification
	HasFiniteStateMachineError() bool
	// Custom returns DeviceBgpCustomError, set in ActionProtocolBgpNotification.
	// DeviceBgpCustomError is a BGP peer can send NOTIFICATION message with user defined Error Code and Error Subcode.
	Custom() DeviceBgpCustomError
	// SetCustom assigns DeviceBgpCustomError provided by user to ActionProtocolBgpNotification.
	// DeviceBgpCustomError is a BGP peer can send NOTIFICATION message with user defined Error Code and Error Subcode.
	SetCustom(value DeviceBgpCustomError) ActionProtocolBgpNotification
	// HasCustom checks if Custom has been set in ActionProtocolBgpNotification
	HasCustom() bool
	setNil()
}

// The names of BGP Peers to send NOTIFICATION to. If no name is specified then NOTIFICATION will be sent to all configured BGP peers.
//
// x-constraint:
// - /components/schemas/Bgp.V4Peer/properties/name
// - /components/schemas/Bgp.V6Peer/properties/name
//
// Names returns a []string
func (obj *actionProtocolBgpNotification) Names() []string {
	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	return obj.obj.Names
}

// The names of BGP Peers to send NOTIFICATION to. If no name is specified then NOTIFICATION will be sent to all configured BGP peers.
//
// x-constraint:
// - /components/schemas/Bgp.V4Peer/properties/name
// - /components/schemas/Bgp.V6Peer/properties/name
//
// SetNames sets the []string value in the ActionProtocolBgpNotification object
func (obj *actionProtocolBgpNotification) SetNames(value []string) ActionProtocolBgpNotification {

	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	obj.obj.Names = value

	return obj
}

type ActionProtocolBgpNotificationChoiceEnum string

// Enum of Choice on ActionProtocolBgpNotification
var ActionProtocolBgpNotificationChoice = struct {
	CEASE                      ActionProtocolBgpNotificationChoiceEnum
	MESSAGE_HEADER_ERROR       ActionProtocolBgpNotificationChoiceEnum
	OPEN_MESSAGE_ERROR         ActionProtocolBgpNotificationChoiceEnum
	UPDATE_MESSAGE_ERROR       ActionProtocolBgpNotificationChoiceEnum
	HOLD_TIMER_EXPIRED         ActionProtocolBgpNotificationChoiceEnum
	FINITE_STATE_MACHINE_ERROR ActionProtocolBgpNotificationChoiceEnum
	CUSTOM                     ActionProtocolBgpNotificationChoiceEnum
}{
	CEASE:                      ActionProtocolBgpNotificationChoiceEnum("cease"),
	MESSAGE_HEADER_ERROR:       ActionProtocolBgpNotificationChoiceEnum("message_header_error"),
	OPEN_MESSAGE_ERROR:         ActionProtocolBgpNotificationChoiceEnum("open_message_error"),
	UPDATE_MESSAGE_ERROR:       ActionProtocolBgpNotificationChoiceEnum("update_message_error"),
	HOLD_TIMER_EXPIRED:         ActionProtocolBgpNotificationChoiceEnum("hold_timer_expired"),
	FINITE_STATE_MACHINE_ERROR: ActionProtocolBgpNotificationChoiceEnum("finite_state_machine_error"),
	CUSTOM:                     ActionProtocolBgpNotificationChoiceEnum("custom"),
}

func (obj *actionProtocolBgpNotification) Choice() ActionProtocolBgpNotificationChoiceEnum {
	return ActionProtocolBgpNotificationChoiceEnum(obj.obj.Choice.Enum().String())
}

// Each BGP NOTIFICATION message includes an Error Code field indicating what type of problem occurred. For certain Error Codes, an Error  Subcode field provides additional details about the specific nature of the problem.  The choice value will provide the Error Code used in NOTIFICATION message.  The Subcode can be set for each of the corresponding errors except for Hold Timer Expired error and BGP Finite State Machine error.  In both of these cases Subcode 0 will be sent. If a user wants to use non zero Sub Code then custom choice can be used.
// Choice returns a string
func (obj *actionProtocolBgpNotification) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *actionProtocolBgpNotification) setChoice(value ActionProtocolBgpNotificationChoiceEnum) ActionProtocolBgpNotification {
	intValue, ok := otg.ActionProtocolBgpNotification_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ActionProtocolBgpNotificationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ActionProtocolBgpNotification_Choice_Enum(intValue)
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

	if value == ActionProtocolBgpNotificationChoice.CEASE {
		obj.obj.Cease = NewDeviceBgpCeaseError().msg()
	}

	if value == ActionProtocolBgpNotificationChoice.MESSAGE_HEADER_ERROR {
		obj.obj.MessageHeaderError = NewDeviceBgpMessageHeaderError().msg()
	}

	if value == ActionProtocolBgpNotificationChoice.OPEN_MESSAGE_ERROR {
		obj.obj.OpenMessageError = NewDeviceBgpOpenMessageError().msg()
	}

	if value == ActionProtocolBgpNotificationChoice.UPDATE_MESSAGE_ERROR {
		obj.obj.UpdateMessageError = NewDeviceBgpUpdateMessageError().msg()
	}

	if value == ActionProtocolBgpNotificationChoice.HOLD_TIMER_EXPIRED {
		obj.obj.HoldTimerExpired = NewDeviceBgpHoldTimerExpired().msg()
	}

	if value == ActionProtocolBgpNotificationChoice.FINITE_STATE_MACHINE_ERROR {
		obj.obj.FiniteStateMachineError = NewDeviceBgpFiniteStateMachineError().msg()
	}

	if value == ActionProtocolBgpNotificationChoice.CUSTOM {
		obj.obj.Custom = NewDeviceBgpCustomError().msg()
	}

	return obj
}

// description is TBD
// Cease returns a DeviceBgpCeaseError
func (obj *actionProtocolBgpNotification) Cease() DeviceBgpCeaseError {
	if obj.obj.Cease == nil {
		obj.setChoice(ActionProtocolBgpNotificationChoice.CEASE)
	}
	if obj.ceaseHolder == nil {
		obj.ceaseHolder = &deviceBgpCeaseError{obj: obj.obj.Cease}
	}
	return obj.ceaseHolder
}

// description is TBD
// Cease returns a DeviceBgpCeaseError
func (obj *actionProtocolBgpNotification) HasCease() bool {
	return obj.obj.Cease != nil
}

// description is TBD
// SetCease sets the DeviceBgpCeaseError value in the ActionProtocolBgpNotification object
func (obj *actionProtocolBgpNotification) SetCease(value DeviceBgpCeaseError) ActionProtocolBgpNotification {
	obj.setChoice(ActionProtocolBgpNotificationChoice.CEASE)
	obj.ceaseHolder = nil
	obj.obj.Cease = value.msg()

	return obj
}

// description is TBD
// MessageHeaderError returns a DeviceBgpMessageHeaderError
func (obj *actionProtocolBgpNotification) MessageHeaderError() DeviceBgpMessageHeaderError {
	if obj.obj.MessageHeaderError == nil {
		obj.setChoice(ActionProtocolBgpNotificationChoice.MESSAGE_HEADER_ERROR)
	}
	if obj.messageHeaderErrorHolder == nil {
		obj.messageHeaderErrorHolder = &deviceBgpMessageHeaderError{obj: obj.obj.MessageHeaderError}
	}
	return obj.messageHeaderErrorHolder
}

// description is TBD
// MessageHeaderError returns a DeviceBgpMessageHeaderError
func (obj *actionProtocolBgpNotification) HasMessageHeaderError() bool {
	return obj.obj.MessageHeaderError != nil
}

// description is TBD
// SetMessageHeaderError sets the DeviceBgpMessageHeaderError value in the ActionProtocolBgpNotification object
func (obj *actionProtocolBgpNotification) SetMessageHeaderError(value DeviceBgpMessageHeaderError) ActionProtocolBgpNotification {
	obj.setChoice(ActionProtocolBgpNotificationChoice.MESSAGE_HEADER_ERROR)
	obj.messageHeaderErrorHolder = nil
	obj.obj.MessageHeaderError = value.msg()

	return obj
}

// description is TBD
// OpenMessageError returns a DeviceBgpOpenMessageError
func (obj *actionProtocolBgpNotification) OpenMessageError() DeviceBgpOpenMessageError {
	if obj.obj.OpenMessageError == nil {
		obj.setChoice(ActionProtocolBgpNotificationChoice.OPEN_MESSAGE_ERROR)
	}
	if obj.openMessageErrorHolder == nil {
		obj.openMessageErrorHolder = &deviceBgpOpenMessageError{obj: obj.obj.OpenMessageError}
	}
	return obj.openMessageErrorHolder
}

// description is TBD
// OpenMessageError returns a DeviceBgpOpenMessageError
func (obj *actionProtocolBgpNotification) HasOpenMessageError() bool {
	return obj.obj.OpenMessageError != nil
}

// description is TBD
// SetOpenMessageError sets the DeviceBgpOpenMessageError value in the ActionProtocolBgpNotification object
func (obj *actionProtocolBgpNotification) SetOpenMessageError(value DeviceBgpOpenMessageError) ActionProtocolBgpNotification {
	obj.setChoice(ActionProtocolBgpNotificationChoice.OPEN_MESSAGE_ERROR)
	obj.openMessageErrorHolder = nil
	obj.obj.OpenMessageError = value.msg()

	return obj
}

// description is TBD
// UpdateMessageError returns a DeviceBgpUpdateMessageError
func (obj *actionProtocolBgpNotification) UpdateMessageError() DeviceBgpUpdateMessageError {
	if obj.obj.UpdateMessageError == nil {
		obj.setChoice(ActionProtocolBgpNotificationChoice.UPDATE_MESSAGE_ERROR)
	}
	if obj.updateMessageErrorHolder == nil {
		obj.updateMessageErrorHolder = &deviceBgpUpdateMessageError{obj: obj.obj.UpdateMessageError}
	}
	return obj.updateMessageErrorHolder
}

// description is TBD
// UpdateMessageError returns a DeviceBgpUpdateMessageError
func (obj *actionProtocolBgpNotification) HasUpdateMessageError() bool {
	return obj.obj.UpdateMessageError != nil
}

// description is TBD
// SetUpdateMessageError sets the DeviceBgpUpdateMessageError value in the ActionProtocolBgpNotification object
func (obj *actionProtocolBgpNotification) SetUpdateMessageError(value DeviceBgpUpdateMessageError) ActionProtocolBgpNotification {
	obj.setChoice(ActionProtocolBgpNotificationChoice.UPDATE_MESSAGE_ERROR)
	obj.updateMessageErrorHolder = nil
	obj.obj.UpdateMessageError = value.msg()

	return obj
}

// description is TBD
// HoldTimerExpired returns a DeviceBgpHoldTimerExpired
func (obj *actionProtocolBgpNotification) HoldTimerExpired() DeviceBgpHoldTimerExpired {
	if obj.obj.HoldTimerExpired == nil {
		obj.setChoice(ActionProtocolBgpNotificationChoice.HOLD_TIMER_EXPIRED)
	}
	if obj.holdTimerExpiredHolder == nil {
		obj.holdTimerExpiredHolder = &deviceBgpHoldTimerExpired{obj: obj.obj.HoldTimerExpired}
	}
	return obj.holdTimerExpiredHolder
}

// description is TBD
// HoldTimerExpired returns a DeviceBgpHoldTimerExpired
func (obj *actionProtocolBgpNotification) HasHoldTimerExpired() bool {
	return obj.obj.HoldTimerExpired != nil
}

// description is TBD
// SetHoldTimerExpired sets the DeviceBgpHoldTimerExpired value in the ActionProtocolBgpNotification object
func (obj *actionProtocolBgpNotification) SetHoldTimerExpired(value DeviceBgpHoldTimerExpired) ActionProtocolBgpNotification {
	obj.setChoice(ActionProtocolBgpNotificationChoice.HOLD_TIMER_EXPIRED)
	obj.holdTimerExpiredHolder = nil
	obj.obj.HoldTimerExpired = value.msg()

	return obj
}

// description is TBD
// FiniteStateMachineError returns a DeviceBgpFiniteStateMachineError
func (obj *actionProtocolBgpNotification) FiniteStateMachineError() DeviceBgpFiniteStateMachineError {
	if obj.obj.FiniteStateMachineError == nil {
		obj.setChoice(ActionProtocolBgpNotificationChoice.FINITE_STATE_MACHINE_ERROR)
	}
	if obj.finiteStateMachineErrorHolder == nil {
		obj.finiteStateMachineErrorHolder = &deviceBgpFiniteStateMachineError{obj: obj.obj.FiniteStateMachineError}
	}
	return obj.finiteStateMachineErrorHolder
}

// description is TBD
// FiniteStateMachineError returns a DeviceBgpFiniteStateMachineError
func (obj *actionProtocolBgpNotification) HasFiniteStateMachineError() bool {
	return obj.obj.FiniteStateMachineError != nil
}

// description is TBD
// SetFiniteStateMachineError sets the DeviceBgpFiniteStateMachineError value in the ActionProtocolBgpNotification object
func (obj *actionProtocolBgpNotification) SetFiniteStateMachineError(value DeviceBgpFiniteStateMachineError) ActionProtocolBgpNotification {
	obj.setChoice(ActionProtocolBgpNotificationChoice.FINITE_STATE_MACHINE_ERROR)
	obj.finiteStateMachineErrorHolder = nil
	obj.obj.FiniteStateMachineError = value.msg()

	return obj
}

// description is TBD
// Custom returns a DeviceBgpCustomError
func (obj *actionProtocolBgpNotification) Custom() DeviceBgpCustomError {
	if obj.obj.Custom == nil {
		obj.setChoice(ActionProtocolBgpNotificationChoice.CUSTOM)
	}
	if obj.customHolder == nil {
		obj.customHolder = &deviceBgpCustomError{obj: obj.obj.Custom}
	}
	return obj.customHolder
}

// description is TBD
// Custom returns a DeviceBgpCustomError
func (obj *actionProtocolBgpNotification) HasCustom() bool {
	return obj.obj.Custom != nil
}

// description is TBD
// SetCustom sets the DeviceBgpCustomError value in the ActionProtocolBgpNotification object
func (obj *actionProtocolBgpNotification) SetCustom(value DeviceBgpCustomError) ActionProtocolBgpNotification {
	obj.setChoice(ActionProtocolBgpNotificationChoice.CUSTOM)
	obj.customHolder = nil
	obj.obj.Custom = value.msg()

	return obj
}

func (obj *actionProtocolBgpNotification) validateObj(vObj *validation, set_default bool) {
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

func (obj *actionProtocolBgpNotification) setDefault() {
	var choices_set int = 0
	var choice ActionProtocolBgpNotificationChoiceEnum

	if obj.obj.Cease != nil {
		choices_set += 1
		choice = ActionProtocolBgpNotificationChoice.CEASE
	}

	if obj.obj.MessageHeaderError != nil {
		choices_set += 1
		choice = ActionProtocolBgpNotificationChoice.MESSAGE_HEADER_ERROR
	}

	if obj.obj.OpenMessageError != nil {
		choices_set += 1
		choice = ActionProtocolBgpNotificationChoice.OPEN_MESSAGE_ERROR
	}

	if obj.obj.UpdateMessageError != nil {
		choices_set += 1
		choice = ActionProtocolBgpNotificationChoice.UPDATE_MESSAGE_ERROR
	}

	if obj.obj.HoldTimerExpired != nil {
		choices_set += 1
		choice = ActionProtocolBgpNotificationChoice.HOLD_TIMER_EXPIRED
	}

	if obj.obj.FiniteStateMachineError != nil {
		choices_set += 1
		choice = ActionProtocolBgpNotificationChoice.FINITE_STATE_MACHINE_ERROR
	}

	if obj.obj.Custom != nil {
		choices_set += 1
		choice = ActionProtocolBgpNotificationChoice.CUSTOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(ActionProtocolBgpNotificationChoice.CEASE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ActionProtocolBgpNotification")
			}
		} else {
			intVal := otg.ActionProtocolBgpNotification_Choice_Enum_value[string(choice)]
			enumValue := otg.ActionProtocolBgpNotification_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

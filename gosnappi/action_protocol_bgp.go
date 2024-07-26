package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolBgp *****
type actionProtocolBgp struct {
	validation
	obj                           *otg.ActionProtocolBgp
	marshaller                    marshalActionProtocolBgp
	unMarshaller                  unMarshalActionProtocolBgp
	notificationHolder            ActionProtocolBgpNotification
	initiateGracefulRestartHolder ActionProtocolBgpInitiateGracefulRestart
}

func NewActionProtocolBgp() ActionProtocolBgp {
	obj := actionProtocolBgp{obj: &otg.ActionProtocolBgp{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolBgp) msg() *otg.ActionProtocolBgp {
	return obj.obj
}

func (obj *actionProtocolBgp) setMsg(msg *otg.ActionProtocolBgp) ActionProtocolBgp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolBgp struct {
	obj *actionProtocolBgp
}

type marshalActionProtocolBgp interface {
	// ToProto marshals ActionProtocolBgp to protobuf object *otg.ActionProtocolBgp
	ToProto() (*otg.ActionProtocolBgp, error)
	// ToPbText marshals ActionProtocolBgp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolBgp to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolBgp to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolBgp struct {
	obj *actionProtocolBgp
}

type unMarshalActionProtocolBgp interface {
	// FromProto unmarshals ActionProtocolBgp from protobuf object *otg.ActionProtocolBgp
	FromProto(msg *otg.ActionProtocolBgp) (ActionProtocolBgp, error)
	// FromPbText unmarshals ActionProtocolBgp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolBgp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolBgp from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolBgp) Marshal() marshalActionProtocolBgp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolBgp{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolBgp) Unmarshal() unMarshalActionProtocolBgp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolBgp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolBgp) ToProto() (*otg.ActionProtocolBgp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolBgp) FromProto(msg *otg.ActionProtocolBgp) (ActionProtocolBgp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolBgp) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolBgp) FromPbText(value string) error {
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

func (m *marshalactionProtocolBgp) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolBgp) FromYaml(value string) error {
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

func (m *marshalactionProtocolBgp) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolBgp) FromJson(value string) error {
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

func (obj *actionProtocolBgp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolBgp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolBgp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolBgp) Clone() (ActionProtocolBgp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolBgp()
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

func (obj *actionProtocolBgp) setNil() {
	obj.notificationHolder = nil
	obj.initiateGracefulRestartHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionProtocolBgp is actions associated with BGP on configured resources.
type ActionProtocolBgp interface {
	Validation
	// msg marshals ActionProtocolBgp to protobuf object *otg.ActionProtocolBgp
	// and doesn't set defaults
	msg() *otg.ActionProtocolBgp
	// setMsg unmarshals ActionProtocolBgp from protobuf object *otg.ActionProtocolBgp
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolBgp) ActionProtocolBgp
	// provides marshal interface
	Marshal() marshalActionProtocolBgp
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolBgp
	// validate validates ActionProtocolBgp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolBgp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ActionProtocolBgpChoiceEnum, set in ActionProtocolBgp
	Choice() ActionProtocolBgpChoiceEnum
	// setChoice assigns ActionProtocolBgpChoiceEnum provided by user to ActionProtocolBgp
	setChoice(value ActionProtocolBgpChoiceEnum) ActionProtocolBgp
	// Notification returns ActionProtocolBgpNotification, set in ActionProtocolBgp.
	// ActionProtocolBgpNotification is a NOTIFICATION message is sent when an error is detected with the BGP session, such as hold timer expiring, misconfigured AS number  or a BGP session reset is requested. This causes the BGP connection to close. Send explicit NOTIFICATIONs for list of specified  BGP peers. If a user wants to send custom Error Code and Error Subcode the custom object should be configured. A user can send IANA defined BGP NOTIFICATIONs according to https://www.iana.org/assignments/bgp-parameters/bgp-parameters.xhtml.
	Notification() ActionProtocolBgpNotification
	// SetNotification assigns ActionProtocolBgpNotification provided by user to ActionProtocolBgp.
	// ActionProtocolBgpNotification is a NOTIFICATION message is sent when an error is detected with the BGP session, such as hold timer expiring, misconfigured AS number  or a BGP session reset is requested. This causes the BGP connection to close. Send explicit NOTIFICATIONs for list of specified  BGP peers. If a user wants to send custom Error Code and Error Subcode the custom object should be configured. A user can send IANA defined BGP NOTIFICATIONs according to https://www.iana.org/assignments/bgp-parameters/bgp-parameters.xhtml.
	SetNotification(value ActionProtocolBgpNotification) ActionProtocolBgp
	// HasNotification checks if Notification has been set in ActionProtocolBgp
	HasNotification() bool
	// InitiateGracefulRestart returns ActionProtocolBgpInitiateGracefulRestart, set in ActionProtocolBgp.
	// ActionProtocolBgpInitiateGracefulRestart is initiates BGP Graceful Restart process for the selected BGP peers. If no name is specified then Graceful Restart will be sent to all configured BGP peers.
	InitiateGracefulRestart() ActionProtocolBgpInitiateGracefulRestart
	// SetInitiateGracefulRestart assigns ActionProtocolBgpInitiateGracefulRestart provided by user to ActionProtocolBgp.
	// ActionProtocolBgpInitiateGracefulRestart is initiates BGP Graceful Restart process for the selected BGP peers. If no name is specified then Graceful Restart will be sent to all configured BGP peers.
	SetInitiateGracefulRestart(value ActionProtocolBgpInitiateGracefulRestart) ActionProtocolBgp
	// HasInitiateGracefulRestart checks if InitiateGracefulRestart has been set in ActionProtocolBgp
	HasInitiateGracefulRestart() bool
	setNil()
}

type ActionProtocolBgpChoiceEnum string

// Enum of Choice on ActionProtocolBgp
var ActionProtocolBgpChoice = struct {
	NOTIFICATION              ActionProtocolBgpChoiceEnum
	INITIATE_GRACEFUL_RESTART ActionProtocolBgpChoiceEnum
}{
	NOTIFICATION:              ActionProtocolBgpChoiceEnum("notification"),
	INITIATE_GRACEFUL_RESTART: ActionProtocolBgpChoiceEnum("initiate_graceful_restart"),
}

func (obj *actionProtocolBgp) Choice() ActionProtocolBgpChoiceEnum {
	return ActionProtocolBgpChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *actionProtocolBgp) setChoice(value ActionProtocolBgpChoiceEnum) ActionProtocolBgp {
	intValue, ok := otg.ActionProtocolBgp_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ActionProtocolBgpChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ActionProtocolBgp_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.InitiateGracefulRestart = nil
	obj.initiateGracefulRestartHolder = nil
	obj.obj.Notification = nil
	obj.notificationHolder = nil

	if value == ActionProtocolBgpChoice.NOTIFICATION {
		obj.obj.Notification = NewActionProtocolBgpNotification().msg()
	}

	if value == ActionProtocolBgpChoice.INITIATE_GRACEFUL_RESTART {
		obj.obj.InitiateGracefulRestart = NewActionProtocolBgpInitiateGracefulRestart().msg()
	}

	return obj
}

// description is TBD
// Notification returns a ActionProtocolBgpNotification
func (obj *actionProtocolBgp) Notification() ActionProtocolBgpNotification {
	if obj.obj.Notification == nil {
		obj.setChoice(ActionProtocolBgpChoice.NOTIFICATION)
	}
	if obj.notificationHolder == nil {
		obj.notificationHolder = &actionProtocolBgpNotification{obj: obj.obj.Notification}
	}
	return obj.notificationHolder
}

// description is TBD
// Notification returns a ActionProtocolBgpNotification
func (obj *actionProtocolBgp) HasNotification() bool {
	return obj.obj.Notification != nil
}

// description is TBD
// SetNotification sets the ActionProtocolBgpNotification value in the ActionProtocolBgp object
func (obj *actionProtocolBgp) SetNotification(value ActionProtocolBgpNotification) ActionProtocolBgp {
	obj.setChoice(ActionProtocolBgpChoice.NOTIFICATION)
	obj.notificationHolder = nil
	obj.obj.Notification = value.msg()

	return obj
}

// description is TBD
// InitiateGracefulRestart returns a ActionProtocolBgpInitiateGracefulRestart
func (obj *actionProtocolBgp) InitiateGracefulRestart() ActionProtocolBgpInitiateGracefulRestart {
	if obj.obj.InitiateGracefulRestart == nil {
		obj.setChoice(ActionProtocolBgpChoice.INITIATE_GRACEFUL_RESTART)
	}
	if obj.initiateGracefulRestartHolder == nil {
		obj.initiateGracefulRestartHolder = &actionProtocolBgpInitiateGracefulRestart{obj: obj.obj.InitiateGracefulRestart}
	}
	return obj.initiateGracefulRestartHolder
}

// description is TBD
// InitiateGracefulRestart returns a ActionProtocolBgpInitiateGracefulRestart
func (obj *actionProtocolBgp) HasInitiateGracefulRestart() bool {
	return obj.obj.InitiateGracefulRestart != nil
}

// description is TBD
// SetInitiateGracefulRestart sets the ActionProtocolBgpInitiateGracefulRestart value in the ActionProtocolBgp object
func (obj *actionProtocolBgp) SetInitiateGracefulRestart(value ActionProtocolBgpInitiateGracefulRestart) ActionProtocolBgp {
	obj.setChoice(ActionProtocolBgpChoice.INITIATE_GRACEFUL_RESTART)
	obj.initiateGracefulRestartHolder = nil
	obj.obj.InitiateGracefulRestart = value.msg()

	return obj
}

func (obj *actionProtocolBgp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface ActionProtocolBgp")
	}

	if obj.obj.Notification != nil {

		obj.Notification().validateObj(vObj, set_default)
	}

	if obj.obj.InitiateGracefulRestart != nil {

		obj.InitiateGracefulRestart().validateObj(vObj, set_default)
	}

}

func (obj *actionProtocolBgp) setDefault() {
	var choices_set int = 0
	var choice ActionProtocolBgpChoiceEnum

	if obj.obj.Notification != nil {
		choices_set += 1
		choice = ActionProtocolBgpChoice.NOTIFICATION
	}

	if obj.obj.InitiateGracefulRestart != nil {
		choices_set += 1
		choice = ActionProtocolBgpChoice.INITIATE_GRACEFUL_RESTART
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ActionProtocolBgp")
			}
		} else {
			intVal := otg.ActionProtocolBgp_Choice_Enum_value[string(choice)]
			enumValue := otg.ActionProtocolBgp_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

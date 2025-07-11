package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolIsis *****
type actionProtocolIsis struct {
	validation
	obj                           *otg.ActionProtocolIsis
	marshaller                    marshalActionProtocolIsis
	unMarshaller                  unMarshalActionProtocolIsis
	initiateGracefulRestartHolder ActionProtocolIsisInitiateRestart
}

func NewActionProtocolIsis() ActionProtocolIsis {
	obj := actionProtocolIsis{obj: &otg.ActionProtocolIsis{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolIsis) msg() *otg.ActionProtocolIsis {
	return obj.obj
}

func (obj *actionProtocolIsis) setMsg(msg *otg.ActionProtocolIsis) ActionProtocolIsis {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolIsis struct {
	obj *actionProtocolIsis
}

type marshalActionProtocolIsis interface {
	// ToProto marshals ActionProtocolIsis to protobuf object *otg.ActionProtocolIsis
	ToProto() (*otg.ActionProtocolIsis, error)
	// ToPbText marshals ActionProtocolIsis to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolIsis to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolIsis to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolIsis struct {
	obj *actionProtocolIsis
}

type unMarshalActionProtocolIsis interface {
	// FromProto unmarshals ActionProtocolIsis from protobuf object *otg.ActionProtocolIsis
	FromProto(msg *otg.ActionProtocolIsis) (ActionProtocolIsis, error)
	// FromPbText unmarshals ActionProtocolIsis from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolIsis from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolIsis from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolIsis) Marshal() marshalActionProtocolIsis {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolIsis{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolIsis) Unmarshal() unMarshalActionProtocolIsis {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolIsis{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolIsis) ToProto() (*otg.ActionProtocolIsis, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolIsis) FromProto(msg *otg.ActionProtocolIsis) (ActionProtocolIsis, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolIsis) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolIsis) FromPbText(value string) error {
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

func (m *marshalactionProtocolIsis) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolIsis) FromYaml(value string) error {
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

func (m *marshalactionProtocolIsis) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolIsis) FromJson(value string) error {
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

func (obj *actionProtocolIsis) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolIsis) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolIsis) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolIsis) Clone() (ActionProtocolIsis, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolIsis()
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

func (obj *actionProtocolIsis) setNil() {
	obj.initiateGracefulRestartHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionProtocolIsis is actions associated with IS-IS on configured resources.
type ActionProtocolIsis interface {
	Validation
	// msg marshals ActionProtocolIsis to protobuf object *otg.ActionProtocolIsis
	// and doesn't set defaults
	msg() *otg.ActionProtocolIsis
	// setMsg unmarshals ActionProtocolIsis from protobuf object *otg.ActionProtocolIsis
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolIsis) ActionProtocolIsis
	// provides marshal interface
	Marshal() marshalActionProtocolIsis
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolIsis
	// validate validates ActionProtocolIsis
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolIsis, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ActionProtocolIsisChoiceEnum, set in ActionProtocolIsis
	Choice() ActionProtocolIsisChoiceEnum
	// setChoice assigns ActionProtocolIsisChoiceEnum provided by user to ActionProtocolIsis
	setChoice(value ActionProtocolIsisChoiceEnum) ActionProtocolIsis
	// InitiateGracefulRestart returns ActionProtocolIsisInitiateRestart, set in ActionProtocolIsis.
	// ActionProtocolIsisInitiateRestart is timers T1 and T2 are used both by a restarting router and a starting router. Timer T3 is used only by a restarting router.
	// - Timer T1 is maintained per interface and indicates the time after which an unacknowledged (re)start attempt will be repeated. Its value is 3 seconds.
	// - Timer T2 is maintained for each LSP database (LSDB) for Level 1 and Level 2. Default value is 90 seconds.
	// When the timer T2 expires or is canceled, indicating that synchronization of that level is complete and SPF for that level is run.
	// - Timer T3 is maintained for the entire system after which the router will declare that it has failed to achieve database synchronization
	// (by setting the overload bit in its own LSP). Its initial value is 65535 seconds and is set to minimum of the remaining times of received IIHs
	// containing a Restart TLV with the RA set.
	InitiateGracefulRestart() ActionProtocolIsisInitiateRestart
	// SetInitiateGracefulRestart assigns ActionProtocolIsisInitiateRestart provided by user to ActionProtocolIsis.
	// ActionProtocolIsisInitiateRestart is timers T1 and T2 are used both by a restarting router and a starting router. Timer T3 is used only by a restarting router.
	// - Timer T1 is maintained per interface and indicates the time after which an unacknowledged (re)start attempt will be repeated. Its value is 3 seconds.
	// - Timer T2 is maintained for each LSP database (LSDB) for Level 1 and Level 2. Default value is 90 seconds.
	// When the timer T2 expires or is canceled, indicating that synchronization of that level is complete and SPF for that level is run.
	// - Timer T3 is maintained for the entire system after which the router will declare that it has failed to achieve database synchronization
	// (by setting the overload bit in its own LSP). Its initial value is 65535 seconds and is set to minimum of the remaining times of received IIHs
	// containing a Restart TLV with the RA set.
	SetInitiateGracefulRestart(value ActionProtocolIsisInitiateRestart) ActionProtocolIsis
	// HasInitiateGracefulRestart checks if InitiateGracefulRestart has been set in ActionProtocolIsis
	HasInitiateGracefulRestart() bool
	setNil()
}

type ActionProtocolIsisChoiceEnum string

// Enum of Choice on ActionProtocolIsis
var ActionProtocolIsisChoice = struct {
	INITIATE_GRACEFUL_RESTART ActionProtocolIsisChoiceEnum
}{
	INITIATE_GRACEFUL_RESTART: ActionProtocolIsisChoiceEnum("initiate_graceful_restart"),
}

func (obj *actionProtocolIsis) Choice() ActionProtocolIsisChoiceEnum {
	return ActionProtocolIsisChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *actionProtocolIsis) setChoice(value ActionProtocolIsisChoiceEnum) ActionProtocolIsis {
	intValue, ok := otg.ActionProtocolIsis_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ActionProtocolIsisChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ActionProtocolIsis_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.InitiateGracefulRestart = nil
	obj.initiateGracefulRestartHolder = nil

	if value == ActionProtocolIsisChoice.INITIATE_GRACEFUL_RESTART {
		obj.obj.InitiateGracefulRestart = NewActionProtocolIsisInitiateRestart().msg()
	}

	return obj
}

// Configuration for the initiation of the IS-IS Graceful Restart.
// InitiateGracefulRestart returns a ActionProtocolIsisInitiateRestart
func (obj *actionProtocolIsis) InitiateGracefulRestart() ActionProtocolIsisInitiateRestart {
	if obj.obj.InitiateGracefulRestart == nil {
		obj.setChoice(ActionProtocolIsisChoice.INITIATE_GRACEFUL_RESTART)
	}
	if obj.initiateGracefulRestartHolder == nil {
		obj.initiateGracefulRestartHolder = &actionProtocolIsisInitiateRestart{obj: obj.obj.InitiateGracefulRestart}
	}
	return obj.initiateGracefulRestartHolder
}

// Configuration for the initiation of the IS-IS Graceful Restart.
// InitiateGracefulRestart returns a ActionProtocolIsisInitiateRestart
func (obj *actionProtocolIsis) HasInitiateGracefulRestart() bool {
	return obj.obj.InitiateGracefulRestart != nil
}

// Configuration for the initiation of the IS-IS Graceful Restart.
// SetInitiateGracefulRestart sets the ActionProtocolIsisInitiateRestart value in the ActionProtocolIsis object
func (obj *actionProtocolIsis) SetInitiateGracefulRestart(value ActionProtocolIsisInitiateRestart) ActionProtocolIsis {
	obj.setChoice(ActionProtocolIsisChoice.INITIATE_GRACEFUL_RESTART)
	obj.initiateGracefulRestartHolder = nil
	obj.obj.InitiateGracefulRestart = value.msg()

	return obj
}

func (obj *actionProtocolIsis) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface ActionProtocolIsis")
	}

	if obj.obj.InitiateGracefulRestart != nil {

		obj.InitiateGracefulRestart().validateObj(vObj, set_default)
	}

}

func (obj *actionProtocolIsis) setDefault() {
	var choices_set int = 0
	var choice ActionProtocolIsisChoiceEnum

	if obj.obj.InitiateGracefulRestart != nil {
		choices_set += 1
		choice = ActionProtocolIsisChoice.INITIATE_GRACEFUL_RESTART
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ActionProtocolIsis")
			}
		} else {
			intVal := otg.ActionProtocolIsis_Choice_Enum_value[string(choice)]
			enumValue := otg.ActionProtocolIsis_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

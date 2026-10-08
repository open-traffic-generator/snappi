package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolLdp *****
type actionProtocolLdp struct {
	validation
	obj                           *otg.ActionProtocolLdp
	marshaller                    marshalActionProtocolLdp
	unMarshaller                  unMarshalActionProtocolLdp
	initiateGracefulRestartHolder ActionProtocolLdpInitiateGracefulRestart
}

func NewActionProtocolLdp() ActionProtocolLdp {
	obj := actionProtocolLdp{obj: &otg.ActionProtocolLdp{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolLdp) msg() *otg.ActionProtocolLdp {
	return obj.obj
}

func (obj *actionProtocolLdp) setMsg(msg *otg.ActionProtocolLdp) ActionProtocolLdp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolLdp struct {
	obj *actionProtocolLdp
}

type marshalActionProtocolLdp interface {
	// ToProto marshals ActionProtocolLdp to protobuf object *otg.ActionProtocolLdp
	ToProto() (*otg.ActionProtocolLdp, error)
	// ToPbText marshals ActionProtocolLdp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolLdp to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolLdp to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolLdp struct {
	obj *actionProtocolLdp
}

type unMarshalActionProtocolLdp interface {
	// FromProto unmarshals ActionProtocolLdp from protobuf object *otg.ActionProtocolLdp
	FromProto(msg *otg.ActionProtocolLdp) (ActionProtocolLdp, error)
	// FromPbText unmarshals ActionProtocolLdp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolLdp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolLdp from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolLdp) Marshal() marshalActionProtocolLdp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolLdp{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolLdp) Unmarshal() unMarshalActionProtocolLdp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolLdp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolLdp) ToProto() (*otg.ActionProtocolLdp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolLdp) FromProto(msg *otg.ActionProtocolLdp) (ActionProtocolLdp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolLdp) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolLdp) FromPbText(value string) error {
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

func (m *marshalactionProtocolLdp) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolLdp) FromYaml(value string) error {
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

func (m *marshalactionProtocolLdp) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolLdp) FromJson(value string) error {
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

func (obj *actionProtocolLdp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolLdp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolLdp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolLdp) Clone() (ActionProtocolLdp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolLdp()
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

func (obj *actionProtocolLdp) setNil() {
	obj.initiateGracefulRestartHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionProtocolLdp is actions associated with LDP on configured resources.
type ActionProtocolLdp interface {
	Validation
	// msg marshals ActionProtocolLdp to protobuf object *otg.ActionProtocolLdp
	// and doesn't set defaults
	msg() *otg.ActionProtocolLdp
	// setMsg unmarshals ActionProtocolLdp from protobuf object *otg.ActionProtocolLdp
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolLdp) ActionProtocolLdp
	// provides marshal interface
	Marshal() marshalActionProtocolLdp
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolLdp
	// validate validates ActionProtocolLdp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolLdp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ActionProtocolLdpChoiceEnum, set in ActionProtocolLdp
	Choice() ActionProtocolLdpChoiceEnum
	// setChoice assigns ActionProtocolLdpChoiceEnum provided by user to ActionProtocolLdp
	setChoice(value ActionProtocolLdpChoiceEnum) ActionProtocolLdp
	// InitiateGracefulRestart returns ActionProtocolLdpInitiateGracefulRestart, set in ActionProtocolLdp.
	// ActionProtocolLdpInitiateGracefulRestart is initiates an LDP graceful restart (RFC 3478) on the selected LDP routers. Each selected router closes its LDP sessions without withdrawing its labels, and re-establishes them after restart_delay seconds. The Initialization messages sent after the restart carry the FT Session TLV with the configured reconnect_time and recovery_time (RFC 3478 Section 2). A selected router must have graceful_restart configured; otherwise the implementation returns an error. If no name is specified, all configured LDP routers are restarted.
	InitiateGracefulRestart() ActionProtocolLdpInitiateGracefulRestart
	// SetInitiateGracefulRestart assigns ActionProtocolLdpInitiateGracefulRestart provided by user to ActionProtocolLdp.
	// ActionProtocolLdpInitiateGracefulRestart is initiates an LDP graceful restart (RFC 3478) on the selected LDP routers. Each selected router closes its LDP sessions without withdrawing its labels, and re-establishes them after restart_delay seconds. The Initialization messages sent after the restart carry the FT Session TLV with the configured reconnect_time and recovery_time (RFC 3478 Section 2). A selected router must have graceful_restart configured; otherwise the implementation returns an error. If no name is specified, all configured LDP routers are restarted.
	SetInitiateGracefulRestart(value ActionProtocolLdpInitiateGracefulRestart) ActionProtocolLdp
	// HasInitiateGracefulRestart checks if InitiateGracefulRestart has been set in ActionProtocolLdp
	HasInitiateGracefulRestart() bool
	setNil()
}

type ActionProtocolLdpChoiceEnum string

// Enum of Choice on ActionProtocolLdp
var ActionProtocolLdpChoice = struct {
	INITIATE_GRACEFUL_RESTART ActionProtocolLdpChoiceEnum
}{
	INITIATE_GRACEFUL_RESTART: ActionProtocolLdpChoiceEnum("initiate_graceful_restart"),
}

func (obj *actionProtocolLdp) Choice() ActionProtocolLdpChoiceEnum {
	return ActionProtocolLdpChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *actionProtocolLdp) setChoice(value ActionProtocolLdpChoiceEnum) ActionProtocolLdp {
	intValue, ok := otg.ActionProtocolLdp_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ActionProtocolLdpChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ActionProtocolLdp_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.InitiateGracefulRestart = nil
	obj.initiateGracefulRestartHolder = nil

	if value == ActionProtocolLdpChoice.INITIATE_GRACEFUL_RESTART {
		obj.obj.InitiateGracefulRestart = NewActionProtocolLdpInitiateGracefulRestart().msg()
	}

	return obj
}

// Configuration for the initiation of an LDP graceful restart.
// InitiateGracefulRestart returns a ActionProtocolLdpInitiateGracefulRestart
func (obj *actionProtocolLdp) InitiateGracefulRestart() ActionProtocolLdpInitiateGracefulRestart {
	if obj.obj.InitiateGracefulRestart == nil {
		obj.setChoice(ActionProtocolLdpChoice.INITIATE_GRACEFUL_RESTART)
	}
	if obj.initiateGracefulRestartHolder == nil {
		obj.initiateGracefulRestartHolder = &actionProtocolLdpInitiateGracefulRestart{obj: obj.obj.InitiateGracefulRestart}
	}
	return obj.initiateGracefulRestartHolder
}

// Configuration for the initiation of an LDP graceful restart.
// InitiateGracefulRestart returns a ActionProtocolLdpInitiateGracefulRestart
func (obj *actionProtocolLdp) HasInitiateGracefulRestart() bool {
	return obj.obj.InitiateGracefulRestart != nil
}

// Configuration for the initiation of an LDP graceful restart.
// SetInitiateGracefulRestart sets the ActionProtocolLdpInitiateGracefulRestart value in the ActionProtocolLdp object
func (obj *actionProtocolLdp) SetInitiateGracefulRestart(value ActionProtocolLdpInitiateGracefulRestart) ActionProtocolLdp {
	obj.setChoice(ActionProtocolLdpChoice.INITIATE_GRACEFUL_RESTART)
	obj.initiateGracefulRestartHolder = nil
	obj.obj.InitiateGracefulRestart = value.msg()

	return obj
}

func (obj *actionProtocolLdp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface ActionProtocolLdp")
	}

	if obj.obj.InitiateGracefulRestart != nil {

		obj.InitiateGracefulRestart().validateObj(vObj, set_default)
	}

}

func (obj *actionProtocolLdp) setDefault() {
	var choices_set int = 0
	var choice ActionProtocolLdpChoiceEnum

	if obj.obj.InitiateGracefulRestart != nil {
		choices_set += 1
		choice = ActionProtocolLdpChoice.INITIATE_GRACEFUL_RESTART
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ActionProtocolLdp")
			}
		} else {
			intVal := otg.ActionProtocolLdp_Choice_Enum_value[string(choice)]
			enumValue := otg.ActionProtocolLdp_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

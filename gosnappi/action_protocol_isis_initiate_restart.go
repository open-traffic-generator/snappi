package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolIsisInitiateRestart *****
type actionProtocolIsisInitiateRestart struct {
	validation
	obj             *otg.ActionProtocolIsisInitiateRestart
	marshaller      marshalActionProtocolIsisInitiateRestart
	unMarshaller    unMarshalActionProtocolIsisInitiateRestart
	unplannedHolder ActionProtocolIsisUnplannedRestart
	plannedHolder   ActionProtocolIsisPlannedRestart
}

func NewActionProtocolIsisInitiateRestart() ActionProtocolIsisInitiateRestart {
	obj := actionProtocolIsisInitiateRestart{obj: &otg.ActionProtocolIsisInitiateRestart{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolIsisInitiateRestart) msg() *otg.ActionProtocolIsisInitiateRestart {
	return obj.obj
}

func (obj *actionProtocolIsisInitiateRestart) setMsg(msg *otg.ActionProtocolIsisInitiateRestart) ActionProtocolIsisInitiateRestart {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolIsisInitiateRestart struct {
	obj *actionProtocolIsisInitiateRestart
}

type marshalActionProtocolIsisInitiateRestart interface {
	// ToProto marshals ActionProtocolIsisInitiateRestart to protobuf object *otg.ActionProtocolIsisInitiateRestart
	ToProto() (*otg.ActionProtocolIsisInitiateRestart, error)
	// ToPbText marshals ActionProtocolIsisInitiateRestart to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolIsisInitiateRestart to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolIsisInitiateRestart to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolIsisInitiateRestart struct {
	obj *actionProtocolIsisInitiateRestart
}

type unMarshalActionProtocolIsisInitiateRestart interface {
	// FromProto unmarshals ActionProtocolIsisInitiateRestart from protobuf object *otg.ActionProtocolIsisInitiateRestart
	FromProto(msg *otg.ActionProtocolIsisInitiateRestart) (ActionProtocolIsisInitiateRestart, error)
	// FromPbText unmarshals ActionProtocolIsisInitiateRestart from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolIsisInitiateRestart from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolIsisInitiateRestart from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolIsisInitiateRestart) Marshal() marshalActionProtocolIsisInitiateRestart {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolIsisInitiateRestart{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolIsisInitiateRestart) Unmarshal() unMarshalActionProtocolIsisInitiateRestart {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolIsisInitiateRestart{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolIsisInitiateRestart) ToProto() (*otg.ActionProtocolIsisInitiateRestart, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolIsisInitiateRestart) FromProto(msg *otg.ActionProtocolIsisInitiateRestart) (ActionProtocolIsisInitiateRestart, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolIsisInitiateRestart) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolIsisInitiateRestart) FromPbText(value string) error {
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

func (m *marshalactionProtocolIsisInitiateRestart) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolIsisInitiateRestart) FromYaml(value string) error {
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

func (m *marshalactionProtocolIsisInitiateRestart) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolIsisInitiateRestart) FromJson(value string) error {
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

func (obj *actionProtocolIsisInitiateRestart) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolIsisInitiateRestart) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolIsisInitiateRestart) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolIsisInitiateRestart) Clone() (ActionProtocolIsisInitiateRestart, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolIsisInitiateRestart()
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

func (obj *actionProtocolIsisInitiateRestart) setNil() {
	obj.unplannedHolder = nil
	obj.plannedHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionProtocolIsisInitiateRestart is timers T1 and T2 are used both by a restarting router and a starting router. Timer T3 is used only by a restarting router.
// - Timer T1 is maintained per interface and indicates the time after which an unacknowledged (re)start attempt will be repeated. Its value is 3 seconds.
// - Timer T2 is maintained for each LSP database (LSDB) for Level 1 and Level 2. Default value is 90 seconds.
// When the timer T2 expires or is canceled, indicating that synchronization of that level is complete and SPF for that level is run.
// - Timer T3 is maintained for the entire system after which the router will declare that it has failed to achieve database synchronization
// (by setting the overload bit in its own LSP). Its initial value is 65535 seconds and is set to minimum of the remaining times of received IIHs
// containing a Restart TLV with the RA set.
type ActionProtocolIsisInitiateRestart interface {
	Validation
	// msg marshals ActionProtocolIsisInitiateRestart to protobuf object *otg.ActionProtocolIsisInitiateRestart
	// and doesn't set defaults
	msg() *otg.ActionProtocolIsisInitiateRestart
	// setMsg unmarshals ActionProtocolIsisInitiateRestart from protobuf object *otg.ActionProtocolIsisInitiateRestart
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolIsisInitiateRestart) ActionProtocolIsisInitiateRestart
	// provides marshal interface
	Marshal() marshalActionProtocolIsisInitiateRestart
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolIsisInitiateRestart
	// validate validates ActionProtocolIsisInitiateRestart
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolIsisInitiateRestart, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterNames returns []string, set in ActionProtocolIsisInitiateRestart.
	RouterNames() []string
	// SetRouterNames assigns []string provided by user to ActionProtocolIsisInitiateRestart
	SetRouterNames(value []string) ActionProtocolIsisInitiateRestart
	// Choice returns ActionProtocolIsisInitiateRestartChoiceEnum, set in ActionProtocolIsisInitiateRestart
	Choice() ActionProtocolIsisInitiateRestartChoiceEnum
	// setChoice assigns ActionProtocolIsisInitiateRestartChoiceEnum provided by user to ActionProtocolIsisInitiateRestart
	setChoice(value ActionProtocolIsisInitiateRestartChoiceEnum) ActionProtocolIsisInitiateRestart
	// HasChoice checks if Choice has been set in ActionProtocolIsisInitiateRestart
	HasChoice() bool
	// Unplanned returns ActionProtocolIsisUnplannedRestart, set in ActionProtocolIsisInitiateRestart.
	// ActionProtocolIsisUnplannedRestart is initiates IS-IS Unplanned Graceful Restart process for the selected IS-IS routers. If no name is specified then Graceful Restart will be sent to all configured IS-IS routers. When an emulated IS-IS router is in the unplanned "Restarting" mode, it sends an IIH PDU containing a Restart TLV with the RR (Restart Request) bit set and holding_time updated to as specified by user to indicate the maximum time within which this router  or routers will complete the graceful restart.  It waits for RA (Restart Acknowledge) in an IIH PDU from Neigbhor(s).  The timer T1 is maintained per interface and indicates the time after which an unacknowledged (re)start attempt will be repeated.
	Unplanned() ActionProtocolIsisUnplannedRestart
	// SetUnplanned assigns ActionProtocolIsisUnplannedRestart provided by user to ActionProtocolIsisInitiateRestart.
	// ActionProtocolIsisUnplannedRestart is initiates IS-IS Unplanned Graceful Restart process for the selected IS-IS routers. If no name is specified then Graceful Restart will be sent to all configured IS-IS routers. When an emulated IS-IS router is in the unplanned "Restarting" mode, it sends an IIH PDU containing a Restart TLV with the RR (Restart Request) bit set and holding_time updated to as specified by user to indicate the maximum time within which this router  or routers will complete the graceful restart.  It waits for RA (Restart Acknowledge) in an IIH PDU from Neigbhor(s).  The timer T1 is maintained per interface and indicates the time after which an unacknowledged (re)start attempt will be repeated.
	SetUnplanned(value ActionProtocolIsisUnplannedRestart) ActionProtocolIsisInitiateRestart
	// HasUnplanned checks if Unplanned has been set in ActionProtocolIsisInitiateRestart
	HasUnplanned() bool
	// Planned returns ActionProtocolIsisPlannedRestart, set in ActionProtocolIsisInitiateRestart.
	// ActionProtocolIsisPlannedRestart is initiates IS-IS Planned Graceful Restart process for the selected IS-IS routers. If no name is specified then Graceful Restart will be sent to all configured IS-IS routers. When an emulated IS-IS router is in the planned "Restarting" mode, it sends an IIH PDU containing a Restart TLV with the PR (Planned Restart Request) bit set and sets the Remaining Time with the restart_time greater than the expected control-plane restart time that is the maximum time within which this router or routers will complete the graceful restart. It waits for PA (Planned Restart Acknowledge) in an IIH PDU from Neighbor(s). The use of the PR bit provides a means to safely support restart periods that are significantly longer than standard Holding Times. The PR bit SHOULD remain set in IIHs until the restart is initiated. Reference: https://datatracker.ietf.org/doc/html/rfc8706#section-3.2.3. Once the Restarting Router receives the Restart Tlv with PA bit is set, it intiates the Restart Request with the RR bit is set. The holding_time is set as the Remaining Time as received in Restart Tlv or by the remaining time of restart_time that was sent in the Planned Restart Request Tlv. This is left to the choice of the implementation.
	Planned() ActionProtocolIsisPlannedRestart
	// SetPlanned assigns ActionProtocolIsisPlannedRestart provided by user to ActionProtocolIsisInitiateRestart.
	// ActionProtocolIsisPlannedRestart is initiates IS-IS Planned Graceful Restart process for the selected IS-IS routers. If no name is specified then Graceful Restart will be sent to all configured IS-IS routers. When an emulated IS-IS router is in the planned "Restarting" mode, it sends an IIH PDU containing a Restart TLV with the PR (Planned Restart Request) bit set and sets the Remaining Time with the restart_time greater than the expected control-plane restart time that is the maximum time within which this router or routers will complete the graceful restart. It waits for PA (Planned Restart Acknowledge) in an IIH PDU from Neighbor(s). The use of the PR bit provides a means to safely support restart periods that are significantly longer than standard Holding Times. The PR bit SHOULD remain set in IIHs until the restart is initiated. Reference: https://datatracker.ietf.org/doc/html/rfc8706#section-3.2.3. Once the Restarting Router receives the Restart Tlv with PA bit is set, it intiates the Restart Request with the RR bit is set. The holding_time is set as the Remaining Time as received in Restart Tlv or by the remaining time of restart_time that was sent in the Planned Restart Request Tlv. This is left to the choice of the implementation.
	SetPlanned(value ActionProtocolIsisPlannedRestart) ActionProtocolIsisInitiateRestart
	// HasPlanned checks if Planned has been set in ActionProtocolIsisInitiateRestart
	HasPlanned() bool
	setNil()
}

// The names of device objects to control.
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// RouterNames returns a []string
func (obj *actionProtocolIsisInitiateRestart) RouterNames() []string {
	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	return obj.obj.RouterNames
}

// The names of device objects to control.
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// SetRouterNames sets the []string value in the ActionProtocolIsisInitiateRestart object
func (obj *actionProtocolIsisInitiateRestart) SetRouterNames(value []string) ActionProtocolIsisInitiateRestart {

	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	obj.obj.RouterNames = value

	return obj
}

type ActionProtocolIsisInitiateRestartChoiceEnum string

// Enum of Choice on ActionProtocolIsisInitiateRestart
var ActionProtocolIsisInitiateRestartChoice = struct {
	UNPLANNED ActionProtocolIsisInitiateRestartChoiceEnum
	PLANNED   ActionProtocolIsisInitiateRestartChoiceEnum
}{
	UNPLANNED: ActionProtocolIsisInitiateRestartChoiceEnum("unplanned"),
	PLANNED:   ActionProtocolIsisInitiateRestartChoiceEnum("planned"),
}

func (obj *actionProtocolIsisInitiateRestart) Choice() ActionProtocolIsisInitiateRestartChoiceEnum {
	return ActionProtocolIsisInitiateRestartChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *actionProtocolIsisInitiateRestart) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *actionProtocolIsisInitiateRestart) setChoice(value ActionProtocolIsisInitiateRestartChoiceEnum) ActionProtocolIsisInitiateRestart {
	intValue, ok := otg.ActionProtocolIsisInitiateRestart_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ActionProtocolIsisInitiateRestartChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ActionProtocolIsisInitiateRestart_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Planned = nil
	obj.plannedHolder = nil
	obj.obj.Unplanned = nil
	obj.unplannedHolder = nil

	if value == ActionProtocolIsisInitiateRestartChoice.UNPLANNED {
		obj.obj.Unplanned = NewActionProtocolIsisUnplannedRestart().msg()
	}

	if value == ActionProtocolIsisInitiateRestartChoice.PLANNED {
		obj.obj.Planned = NewActionProtocolIsisPlannedRestart().msg()
	}

	return obj
}

// description is TBD
// Unplanned returns a ActionProtocolIsisUnplannedRestart
func (obj *actionProtocolIsisInitiateRestart) Unplanned() ActionProtocolIsisUnplannedRestart {
	if obj.obj.Unplanned == nil {
		obj.setChoice(ActionProtocolIsisInitiateRestartChoice.UNPLANNED)
	}
	if obj.unplannedHolder == nil {
		obj.unplannedHolder = &actionProtocolIsisUnplannedRestart{obj: obj.obj.Unplanned}
	}
	return obj.unplannedHolder
}

// description is TBD
// Unplanned returns a ActionProtocolIsisUnplannedRestart
func (obj *actionProtocolIsisInitiateRestart) HasUnplanned() bool {
	return obj.obj.Unplanned != nil
}

// description is TBD
// SetUnplanned sets the ActionProtocolIsisUnplannedRestart value in the ActionProtocolIsisInitiateRestart object
func (obj *actionProtocolIsisInitiateRestart) SetUnplanned(value ActionProtocolIsisUnplannedRestart) ActionProtocolIsisInitiateRestart {
	obj.setChoice(ActionProtocolIsisInitiateRestartChoice.UNPLANNED)
	obj.unplannedHolder = nil
	obj.obj.Unplanned = value.msg()

	return obj
}

// description is TBD
// Planned returns a ActionProtocolIsisPlannedRestart
func (obj *actionProtocolIsisInitiateRestart) Planned() ActionProtocolIsisPlannedRestart {
	if obj.obj.Planned == nil {
		obj.setChoice(ActionProtocolIsisInitiateRestartChoice.PLANNED)
	}
	if obj.plannedHolder == nil {
		obj.plannedHolder = &actionProtocolIsisPlannedRestart{obj: obj.obj.Planned}
	}
	return obj.plannedHolder
}

// description is TBD
// Planned returns a ActionProtocolIsisPlannedRestart
func (obj *actionProtocolIsisInitiateRestart) HasPlanned() bool {
	return obj.obj.Planned != nil
}

// description is TBD
// SetPlanned sets the ActionProtocolIsisPlannedRestart value in the ActionProtocolIsisInitiateRestart object
func (obj *actionProtocolIsisInitiateRestart) SetPlanned(value ActionProtocolIsisPlannedRestart) ActionProtocolIsisInitiateRestart {
	obj.setChoice(ActionProtocolIsisInitiateRestartChoice.PLANNED)
	obj.plannedHolder = nil
	obj.obj.Planned = value.msg()

	return obj
}

func (obj *actionProtocolIsisInitiateRestart) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Unplanned != nil {

		obj.Unplanned().validateObj(vObj, set_default)
	}

	if obj.obj.Planned != nil {

		obj.Planned().validateObj(vObj, set_default)
	}

}

func (obj *actionProtocolIsisInitiateRestart) setDefault() {
	var choices_set int = 0
	var choice ActionProtocolIsisInitiateRestartChoiceEnum

	if obj.obj.Unplanned != nil {
		choices_set += 1
		choice = ActionProtocolIsisInitiateRestartChoice.UNPLANNED
	}

	if obj.obj.Planned != nil {
		choices_set += 1
		choice = ActionProtocolIsisInitiateRestartChoice.PLANNED
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(ActionProtocolIsisInitiateRestartChoice.UNPLANNED)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ActionProtocolIsisInitiateRestart")
			}
		} else {
			intVal := otg.ActionProtocolIsisInitiateRestart_Choice_Enum_value[string(choice)]
			enumValue := otg.ActionProtocolIsisInitiateRestart_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

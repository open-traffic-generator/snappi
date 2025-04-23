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
	obj                            *otg.ActionProtocolIsis
	marshaller                     marshalActionProtocolIsis
	unMarshaller                   unMarshalActionProtocolIsis
	initiateUnplannedRestartHolder IsisActionProtocolIsisStartParams
	initiateStartHolder            IsisActionProtocolIsisStartParams
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
	obj.initiateUnplannedRestartHolder = nil
	obj.initiateStartHolder = nil
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
	// getter for InitiateStarting to set choice.
	InitiateStarting()
	// InitiateUnplannedRestart returns IsisActionProtocolIsisStartParams, set in ActionProtocolIsis.
	// IsisActionProtocolIsisStartParams is configuration for IS-IS Graceful Restart
	InitiateUnplannedRestart() IsisActionProtocolIsisStartParams
	// SetInitiateUnplannedRestart assigns IsisActionProtocolIsisStartParams provided by user to ActionProtocolIsis.
	// IsisActionProtocolIsisStartParams is configuration for IS-IS Graceful Restart
	SetInitiateUnplannedRestart(value IsisActionProtocolIsisStartParams) ActionProtocolIsis
	// HasInitiateUnplannedRestart checks if InitiateUnplannedRestart has been set in ActionProtocolIsis
	HasInitiateUnplannedRestart() bool
	// InitiateStart returns IsisActionProtocolIsisStartParams, set in ActionProtocolIsis.
	// IsisActionProtocolIsisStartParams is configuration for IS-IS Graceful Restart
	InitiateStart() IsisActionProtocolIsisStartParams
	// SetInitiateStart assigns IsisActionProtocolIsisStartParams provided by user to ActionProtocolIsis.
	// IsisActionProtocolIsisStartParams is configuration for IS-IS Graceful Restart
	SetInitiateStart(value IsisActionProtocolIsisStartParams) ActionProtocolIsis
	// HasInitiateStart checks if InitiateStart has been set in ActionProtocolIsis
	HasInitiateStart() bool
	setNil()
}

type ActionProtocolIsisChoiceEnum string

// Enum of Choice on ActionProtocolIsis
var ActionProtocolIsisChoice = struct {
	INITIATE_UNPLANNED_RESTART ActionProtocolIsisChoiceEnum
	INITIATE_STARTING          ActionProtocolIsisChoiceEnum
}{
	INITIATE_UNPLANNED_RESTART: ActionProtocolIsisChoiceEnum("initiate_unplanned_restart"),
	INITIATE_STARTING:          ActionProtocolIsisChoiceEnum("initiate_starting"),
}

func (obj *actionProtocolIsis) Choice() ActionProtocolIsisChoiceEnum {
	return ActionProtocolIsisChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for InitiateStarting to set choice
func (obj *actionProtocolIsis) InitiateStarting() {
	obj.setChoice(ActionProtocolIsisChoice.INITIATE_STARTING)
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
	obj.obj.InitiateUnplannedRestart = nil
	obj.initiateUnplannedRestartHolder = nil

	if value == ActionProtocolIsisChoice.INITIATE_UNPLANNED_RESTART {
		obj.obj.InitiateUnplannedRestart = NewIsisActionProtocolIsisStartParams().msg()
	}

	return obj
}

// Initiates IS-IS Unplanned Graceful Restart process for the selected IS-IS routers. If no name is specified then Graceful Restart will be sent to all configured IS-IS routers. When an emulated IS-IS router is in the unplanned "Restarting" mode, it sends to the neighbor router (DUT) an IIH containing a Restart TLV with the RR (Restart Request) bit set. To emulate scenarios where a router sends Graceful Restart TlV this will result in  Unplanned Graceful Restart scenario to be triggered as per Reference: https://datatracker.ietf.org/doc/html/rfc5306.
// InitiateUnplannedRestart returns a IsisActionProtocolIsisStartParams
func (obj *actionProtocolIsis) InitiateUnplannedRestart() IsisActionProtocolIsisStartParams {
	if obj.obj.InitiateUnplannedRestart == nil {
		obj.setChoice(ActionProtocolIsisChoice.INITIATE_UNPLANNED_RESTART)
	}
	if obj.initiateUnplannedRestartHolder == nil {
		obj.initiateUnplannedRestartHolder = &isisActionProtocolIsisStartParams{obj: obj.obj.InitiateUnplannedRestart}
	}
	return obj.initiateUnplannedRestartHolder
}

// Initiates IS-IS Unplanned Graceful Restart process for the selected IS-IS routers. If no name is specified then Graceful Restart will be sent to all configured IS-IS routers. When an emulated IS-IS router is in the unplanned "Restarting" mode, it sends to the neighbor router (DUT) an IIH containing a Restart TLV with the RR (Restart Request) bit set. To emulate scenarios where a router sends Graceful Restart TlV this will result in  Unplanned Graceful Restart scenario to be triggered as per Reference: https://datatracker.ietf.org/doc/html/rfc5306.
// InitiateUnplannedRestart returns a IsisActionProtocolIsisStartParams
func (obj *actionProtocolIsis) HasInitiateUnplannedRestart() bool {
	return obj.obj.InitiateUnplannedRestart != nil
}

// Initiates IS-IS Unplanned Graceful Restart process for the selected IS-IS routers. If no name is specified then Graceful Restart will be sent to all configured IS-IS routers. When an emulated IS-IS router is in the unplanned "Restarting" mode, it sends to the neighbor router (DUT) an IIH containing a Restart TLV with the RR (Restart Request) bit set. To emulate scenarios where a router sends Graceful Restart TlV this will result in  Unplanned Graceful Restart scenario to be triggered as per Reference: https://datatracker.ietf.org/doc/html/rfc5306.
// SetInitiateUnplannedRestart sets the IsisActionProtocolIsisStartParams value in the ActionProtocolIsis object
func (obj *actionProtocolIsis) SetInitiateUnplannedRestart(value IsisActionProtocolIsisStartParams) ActionProtocolIsis {
	obj.setChoice(ActionProtocolIsisChoice.INITIATE_UNPLANNED_RESTART)
	obj.initiateUnplannedRestartHolder = nil
	obj.obj.InitiateUnplannedRestart = value.msg()

	return obj
}

// Initiates IS-IS Start Process for the selected IS-IS routers.  It sends to the neighbor router (DUT) an IIH containing a Restart TLV with the SA (Suppress Adjacency Advertisement) bit set.
// InitiateStart returns a IsisActionProtocolIsisStartParams
func (obj *actionProtocolIsis) InitiateStart() IsisActionProtocolIsisStartParams {
	if obj.obj.InitiateStart == nil {
		obj.obj.InitiateStart = NewIsisActionProtocolIsisStartParams().msg()
	}
	if obj.initiateStartHolder == nil {
		obj.initiateStartHolder = &isisActionProtocolIsisStartParams{obj: obj.obj.InitiateStart}
	}
	return obj.initiateStartHolder
}

// Initiates IS-IS Start Process for the selected IS-IS routers.  It sends to the neighbor router (DUT) an IIH containing a Restart TLV with the SA (Suppress Adjacency Advertisement) bit set.
// InitiateStart returns a IsisActionProtocolIsisStartParams
func (obj *actionProtocolIsis) HasInitiateStart() bool {
	return obj.obj.InitiateStart != nil
}

// Initiates IS-IS Start Process for the selected IS-IS routers.  It sends to the neighbor router (DUT) an IIH containing a Restart TLV with the SA (Suppress Adjacency Advertisement) bit set.
// SetInitiateStart sets the IsisActionProtocolIsisStartParams value in the ActionProtocolIsis object
func (obj *actionProtocolIsis) SetInitiateStart(value IsisActionProtocolIsisStartParams) ActionProtocolIsis {

	obj.initiateStartHolder = nil
	obj.obj.InitiateStart = value.msg()

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

	if obj.obj.InitiateUnplannedRestart != nil {

		obj.InitiateUnplannedRestart().validateObj(vObj, set_default)
	}

	if obj.obj.InitiateStart != nil {

		obj.InitiateStart().validateObj(vObj, set_default)
	}

}

func (obj *actionProtocolIsis) setDefault() {
	var choices_set int = 0
	var choice ActionProtocolIsisChoiceEnum

	if obj.obj.InitiateUnplannedRestart != nil {
		choices_set += 1
		choice = ActionProtocolIsisChoice.INITIATE_UNPLANNED_RESTART
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

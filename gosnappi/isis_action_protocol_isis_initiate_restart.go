package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisActionProtocolIsisInitiateRestart *****
type isisActionProtocolIsisInitiateRestart struct {
	validation
	obj                     *otg.IsisActionProtocolIsisInitiateRestart
	marshaller              marshalIsisActionProtocolIsisInitiateRestart
	unMarshaller            unMarshalIsisActionProtocolIsisInitiateRestart
	unplannedHolder         IsisActionProtocolIsisUnplannedRestart
	subpressAdjacencyHolder IsisActionProtocolIsisSubpressAdjacency
}

func NewIsisActionProtocolIsisInitiateRestart() IsisActionProtocolIsisInitiateRestart {
	obj := isisActionProtocolIsisInitiateRestart{obj: &otg.IsisActionProtocolIsisInitiateRestart{}}
	obj.setDefault()
	return &obj
}

func (obj *isisActionProtocolIsisInitiateRestart) msg() *otg.IsisActionProtocolIsisInitiateRestart {
	return obj.obj
}

func (obj *isisActionProtocolIsisInitiateRestart) setMsg(msg *otg.IsisActionProtocolIsisInitiateRestart) IsisActionProtocolIsisInitiateRestart {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisActionProtocolIsisInitiateRestart struct {
	obj *isisActionProtocolIsisInitiateRestart
}

type marshalIsisActionProtocolIsisInitiateRestart interface {
	// ToProto marshals IsisActionProtocolIsisInitiateRestart to protobuf object *otg.IsisActionProtocolIsisInitiateRestart
	ToProto() (*otg.IsisActionProtocolIsisInitiateRestart, error)
	// ToPbText marshals IsisActionProtocolIsisInitiateRestart to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisActionProtocolIsisInitiateRestart to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisActionProtocolIsisInitiateRestart to JSON text
	ToJson() (string, error)
}

type unMarshalisisActionProtocolIsisInitiateRestart struct {
	obj *isisActionProtocolIsisInitiateRestart
}

type unMarshalIsisActionProtocolIsisInitiateRestart interface {
	// FromProto unmarshals IsisActionProtocolIsisInitiateRestart from protobuf object *otg.IsisActionProtocolIsisInitiateRestart
	FromProto(msg *otg.IsisActionProtocolIsisInitiateRestart) (IsisActionProtocolIsisInitiateRestart, error)
	// FromPbText unmarshals IsisActionProtocolIsisInitiateRestart from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisActionProtocolIsisInitiateRestart from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisActionProtocolIsisInitiateRestart from JSON text
	FromJson(value string) error
}

func (obj *isisActionProtocolIsisInitiateRestart) Marshal() marshalIsisActionProtocolIsisInitiateRestart {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisActionProtocolIsisInitiateRestart{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisActionProtocolIsisInitiateRestart) Unmarshal() unMarshalIsisActionProtocolIsisInitiateRestart {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisActionProtocolIsisInitiateRestart{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisActionProtocolIsisInitiateRestart) ToProto() (*otg.IsisActionProtocolIsisInitiateRestart, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisActionProtocolIsisInitiateRestart) FromProto(msg *otg.IsisActionProtocolIsisInitiateRestart) (IsisActionProtocolIsisInitiateRestart, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisActionProtocolIsisInitiateRestart) ToPbText() (string, error) {
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

func (m *unMarshalisisActionProtocolIsisInitiateRestart) FromPbText(value string) error {
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

func (m *marshalisisActionProtocolIsisInitiateRestart) ToYaml() (string, error) {
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

func (m *unMarshalisisActionProtocolIsisInitiateRestart) FromYaml(value string) error {
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

func (m *marshalisisActionProtocolIsisInitiateRestart) ToJson() (string, error) {
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

func (m *unMarshalisisActionProtocolIsisInitiateRestart) FromJson(value string) error {
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

func (obj *isisActionProtocolIsisInitiateRestart) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisActionProtocolIsisInitiateRestart) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisActionProtocolIsisInitiateRestart) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisActionProtocolIsisInitiateRestart) Clone() (IsisActionProtocolIsisInitiateRestart, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisActionProtocolIsisInitiateRestart()
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

func (obj *isisActionProtocolIsisInitiateRestart) setNil() {
	obj.unplannedHolder = nil
	obj.subpressAdjacencyHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisActionProtocolIsisInitiateRestart is container of the configiuration for the initiation of IS-IS Graceful Restart.
// 3 Timers T1 and T2 are used both by a restarting router and a starting router. Timer T3 is used only by a restarting router.
// - Timer T1 is maintained per interface and indicates the time after which an unacknowledged (re)start attempt will be repeated. Its value is 3 seconds.
// - Timer T2 is maintained for each LSP database (LSPDB) for Level 1 and Level 2. Default value is 60 seconds.
// - Timer T3 is maintained for the entire system fter which the router will declare that it has failed to achieve database synchronization
// (by setting the overload bit in its own LSP). Its initial value is 65535 seconds and is set to minimum of the remaining times of received IIHs
// containing a Restart TlV with the RA set.
type IsisActionProtocolIsisInitiateRestart interface {
	Validation
	// msg marshals IsisActionProtocolIsisInitiateRestart to protobuf object *otg.IsisActionProtocolIsisInitiateRestart
	// and doesn't set defaults
	msg() *otg.IsisActionProtocolIsisInitiateRestart
	// setMsg unmarshals IsisActionProtocolIsisInitiateRestart from protobuf object *otg.IsisActionProtocolIsisInitiateRestart
	// and doesn't set defaults
	setMsg(*otg.IsisActionProtocolIsisInitiateRestart) IsisActionProtocolIsisInitiateRestart
	// provides marshal interface
	Marshal() marshalIsisActionProtocolIsisInitiateRestart
	// provides unmarshal interface
	Unmarshal() unMarshalIsisActionProtocolIsisInitiateRestart
	// validate validates IsisActionProtocolIsisInitiateRestart
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisActionProtocolIsisInitiateRestart, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterNames returns []string, set in IsisActionProtocolIsisInitiateRestart.
	RouterNames() []string
	// SetRouterNames assigns []string provided by user to IsisActionProtocolIsisInitiateRestart
	SetRouterNames(value []string) IsisActionProtocolIsisInitiateRestart
	// Choice returns IsisActionProtocolIsisInitiateRestartChoiceEnum, set in IsisActionProtocolIsisInitiateRestart
	Choice() IsisActionProtocolIsisInitiateRestartChoiceEnum
	// setChoice assigns IsisActionProtocolIsisInitiateRestartChoiceEnum provided by user to IsisActionProtocolIsisInitiateRestart
	setChoice(value IsisActionProtocolIsisInitiateRestartChoiceEnum) IsisActionProtocolIsisInitiateRestart
	// HasChoice checks if Choice has been set in IsisActionProtocolIsisInitiateRestart
	HasChoice() bool
	// Unplanned returns IsisActionProtocolIsisUnplannedRestart, set in IsisActionProtocolIsisInitiateRestart.
	// IsisActionProtocolIsisUnplannedRestart is initiates IS-IS Unplanned Graceful Restart process for the selected IS-IS routers. If no name is specified then Graceful Restart will be sent to all configured IS-IS routers. When an emulated IS-IS router is in the unplanned "Restarting" mode, it sends to the neighbor router (DUT) an IIH containing a Restart TLV with the RR (Restart Request) bit set.
	Unplanned() IsisActionProtocolIsisUnplannedRestart
	// SetUnplanned assigns IsisActionProtocolIsisUnplannedRestart provided by user to IsisActionProtocolIsisInitiateRestart.
	// IsisActionProtocolIsisUnplannedRestart is initiates IS-IS Unplanned Graceful Restart process for the selected IS-IS routers. If no name is specified then Graceful Restart will be sent to all configured IS-IS routers. When an emulated IS-IS router is in the unplanned "Restarting" mode, it sends to the neighbor router (DUT) an IIH containing a Restart TLV with the RR (Restart Request) bit set.
	SetUnplanned(value IsisActionProtocolIsisUnplannedRestart) IsisActionProtocolIsisInitiateRestart
	// HasUnplanned checks if Unplanned has been set in IsisActionProtocolIsisInitiateRestart
	HasUnplanned() bool
	// SubpressAdjacency returns IsisActionProtocolIsisSubpressAdjacency, set in IsisActionProtocolIsisInitiateRestart.
	// IsisActionProtocolIsisSubpressAdjacency is initiates IS-IS Suppress adjacency advertisement process for the selected IS-IS routers. If no name is specified then Start will be sent to all configured IS-IS routers. When an emulated IS-IS router is in the unplanned "Restarting" mode, it sends to the neighbor router (DUT) an IIH containing a Restart TLV with the SA (Suppress adjacency advertisement) bit set.
	SubpressAdjacency() IsisActionProtocolIsisSubpressAdjacency
	// SetSubpressAdjacency assigns IsisActionProtocolIsisSubpressAdjacency provided by user to IsisActionProtocolIsisInitiateRestart.
	// IsisActionProtocolIsisSubpressAdjacency is initiates IS-IS Suppress adjacency advertisement process for the selected IS-IS routers. If no name is specified then Start will be sent to all configured IS-IS routers. When an emulated IS-IS router is in the unplanned "Restarting" mode, it sends to the neighbor router (DUT) an IIH containing a Restart TLV with the SA (Suppress adjacency advertisement) bit set.
	SetSubpressAdjacency(value IsisActionProtocolIsisSubpressAdjacency) IsisActionProtocolIsisInitiateRestart
	// HasSubpressAdjacency checks if SubpressAdjacency has been set in IsisActionProtocolIsisInitiateRestart
	HasSubpressAdjacency() bool
	setNil()
}

// The names of device objects to control.
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// RouterNames returns a []string
func (obj *isisActionProtocolIsisInitiateRestart) RouterNames() []string {
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
// SetRouterNames sets the []string value in the IsisActionProtocolIsisInitiateRestart object
func (obj *isisActionProtocolIsisInitiateRestart) SetRouterNames(value []string) IsisActionProtocolIsisInitiateRestart {

	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	obj.obj.RouterNames = value

	return obj
}

type IsisActionProtocolIsisInitiateRestartChoiceEnum string

// Enum of Choice on IsisActionProtocolIsisInitiateRestart
var IsisActionProtocolIsisInitiateRestartChoice = struct {
	UNPLANNED          IsisActionProtocolIsisInitiateRestartChoiceEnum
	SUBPRESS_ADJACENCY IsisActionProtocolIsisInitiateRestartChoiceEnum
}{
	UNPLANNED:          IsisActionProtocolIsisInitiateRestartChoiceEnum("unplanned"),
	SUBPRESS_ADJACENCY: IsisActionProtocolIsisInitiateRestartChoiceEnum("subpress_adjacency"),
}

func (obj *isisActionProtocolIsisInitiateRestart) Choice() IsisActionProtocolIsisInitiateRestartChoiceEnum {
	return IsisActionProtocolIsisInitiateRestartChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *isisActionProtocolIsisInitiateRestart) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *isisActionProtocolIsisInitiateRestart) setChoice(value IsisActionProtocolIsisInitiateRestartChoiceEnum) IsisActionProtocolIsisInitiateRestart {
	intValue, ok := otg.IsisActionProtocolIsisInitiateRestart_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisActionProtocolIsisInitiateRestartChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisActionProtocolIsisInitiateRestart_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.SubpressAdjacency = nil
	obj.subpressAdjacencyHolder = nil
	obj.obj.Unplanned = nil
	obj.unplannedHolder = nil

	if value == IsisActionProtocolIsisInitiateRestartChoice.UNPLANNED {
		obj.obj.Unplanned = NewIsisActionProtocolIsisUnplannedRestart().msg()
	}

	if value == IsisActionProtocolIsisInitiateRestartChoice.SUBPRESS_ADJACENCY {
		obj.obj.SubpressAdjacency = NewIsisActionProtocolIsisSubpressAdjacency().msg()
	}

	return obj
}

// description is TBD
// Unplanned returns a IsisActionProtocolIsisUnplannedRestart
func (obj *isisActionProtocolIsisInitiateRestart) Unplanned() IsisActionProtocolIsisUnplannedRestart {
	if obj.obj.Unplanned == nil {
		obj.setChoice(IsisActionProtocolIsisInitiateRestartChoice.UNPLANNED)
	}
	if obj.unplannedHolder == nil {
		obj.unplannedHolder = &isisActionProtocolIsisUnplannedRestart{obj: obj.obj.Unplanned}
	}
	return obj.unplannedHolder
}

// description is TBD
// Unplanned returns a IsisActionProtocolIsisUnplannedRestart
func (obj *isisActionProtocolIsisInitiateRestart) HasUnplanned() bool {
	return obj.obj.Unplanned != nil
}

// description is TBD
// SetUnplanned sets the IsisActionProtocolIsisUnplannedRestart value in the IsisActionProtocolIsisInitiateRestart object
func (obj *isisActionProtocolIsisInitiateRestart) SetUnplanned(value IsisActionProtocolIsisUnplannedRestart) IsisActionProtocolIsisInitiateRestart {
	obj.setChoice(IsisActionProtocolIsisInitiateRestartChoice.UNPLANNED)
	obj.unplannedHolder = nil
	obj.obj.Unplanned = value.msg()

	return obj
}

// description is TBD
// SubpressAdjacency returns a IsisActionProtocolIsisSubpressAdjacency
func (obj *isisActionProtocolIsisInitiateRestart) SubpressAdjacency() IsisActionProtocolIsisSubpressAdjacency {
	if obj.obj.SubpressAdjacency == nil {
		obj.setChoice(IsisActionProtocolIsisInitiateRestartChoice.SUBPRESS_ADJACENCY)
	}
	if obj.subpressAdjacencyHolder == nil {
		obj.subpressAdjacencyHolder = &isisActionProtocolIsisSubpressAdjacency{obj: obj.obj.SubpressAdjacency}
	}
	return obj.subpressAdjacencyHolder
}

// description is TBD
// SubpressAdjacency returns a IsisActionProtocolIsisSubpressAdjacency
func (obj *isisActionProtocolIsisInitiateRestart) HasSubpressAdjacency() bool {
	return obj.obj.SubpressAdjacency != nil
}

// description is TBD
// SetSubpressAdjacency sets the IsisActionProtocolIsisSubpressAdjacency value in the IsisActionProtocolIsisInitiateRestart object
func (obj *isisActionProtocolIsisInitiateRestart) SetSubpressAdjacency(value IsisActionProtocolIsisSubpressAdjacency) IsisActionProtocolIsisInitiateRestart {
	obj.setChoice(IsisActionProtocolIsisInitiateRestartChoice.SUBPRESS_ADJACENCY)
	obj.subpressAdjacencyHolder = nil
	obj.obj.SubpressAdjacency = value.msg()

	return obj
}

func (obj *isisActionProtocolIsisInitiateRestart) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Unplanned != nil {

		obj.Unplanned().validateObj(vObj, set_default)
	}

	if obj.obj.SubpressAdjacency != nil {

		obj.SubpressAdjacency().validateObj(vObj, set_default)
	}

}

func (obj *isisActionProtocolIsisInitiateRestart) setDefault() {
	var choices_set int = 0
	var choice IsisActionProtocolIsisInitiateRestartChoiceEnum

	if obj.obj.Unplanned != nil {
		choices_set += 1
		choice = IsisActionProtocolIsisInitiateRestartChoice.UNPLANNED
	}

	if obj.obj.SubpressAdjacency != nil {
		choices_set += 1
		choice = IsisActionProtocolIsisInitiateRestartChoice.SUBPRESS_ADJACENCY
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(IsisActionProtocolIsisInitiateRestartChoice.UNPLANNED)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in IsisActionProtocolIsisInitiateRestart")
			}
		} else {
			intVal := otg.IsisActionProtocolIsisInitiateRestart_Choice_Enum_value[string(choice)]
			enumValue := otg.IsisActionProtocolIsisInitiateRestart_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

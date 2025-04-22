package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocol *****
type stateProtocol struct {
	validation
	obj          *otg.StateProtocol
	marshaller   marshalStateProtocol
	unMarshaller unMarshalStateProtocol
	allHolder    StateProtocolAll
	routeHolder  StateProtocolRoute
	lacpHolder   StateProtocolLacp
	bgpHolder    StateProtocolBgp
	isisHolder   StateProtocolIsis
	ospfv2Holder StateProtocolOspfv2
	ospfv3Holder StateProtocolOspfv3
	rocev2Holder StateProtocolRocev2
}

func NewStateProtocol() StateProtocol {
	obj := stateProtocol{obj: &otg.StateProtocol{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocol) msg() *otg.StateProtocol {
	return obj.obj
}

func (obj *stateProtocol) setMsg(msg *otg.StateProtocol) StateProtocol {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocol struct {
	obj *stateProtocol
}

type marshalStateProtocol interface {
	// ToProto marshals StateProtocol to protobuf object *otg.StateProtocol
	ToProto() (*otg.StateProtocol, error)
	// ToPbText marshals StateProtocol to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocol to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocol to JSON text
	ToJson() (string, error)
}

type unMarshalstateProtocol struct {
	obj *stateProtocol
}

type unMarshalStateProtocol interface {
	// FromProto unmarshals StateProtocol from protobuf object *otg.StateProtocol
	FromProto(msg *otg.StateProtocol) (StateProtocol, error)
	// FromPbText unmarshals StateProtocol from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocol from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocol from JSON text
	FromJson(value string) error
}

func (obj *stateProtocol) Marshal() marshalStateProtocol {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocol{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocol) Unmarshal() unMarshalStateProtocol {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocol{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocol) ToProto() (*otg.StateProtocol, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocol) FromProto(msg *otg.StateProtocol) (StateProtocol, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocol) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocol) FromPbText(value string) error {
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

func (m *marshalstateProtocol) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocol) FromYaml(value string) error {
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

func (m *marshalstateProtocol) ToJson() (string, error) {
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

func (m *unMarshalstateProtocol) FromJson(value string) error {
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

func (obj *stateProtocol) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocol) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocol) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocol) Clone() (StateProtocol, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocol()
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

func (obj *stateProtocol) setNil() {
	obj.allHolder = nil
	obj.routeHolder = nil
	obj.lacpHolder = nil
	obj.bgpHolder = nil
	obj.isisHolder = nil
	obj.ospfv2Holder = nil
	obj.ospfv3Holder = nil
	obj.rocev2Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// StateProtocol is states associated with protocols on configured resources.
type StateProtocol interface {
	Validation
	// msg marshals StateProtocol to protobuf object *otg.StateProtocol
	// and doesn't set defaults
	msg() *otg.StateProtocol
	// setMsg unmarshals StateProtocol from protobuf object *otg.StateProtocol
	// and doesn't set defaults
	setMsg(*otg.StateProtocol) StateProtocol
	// provides marshal interface
	Marshal() marshalStateProtocol
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocol
	// validate validates StateProtocol
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocol, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns StateProtocolChoiceEnum, set in StateProtocol
	Choice() StateProtocolChoiceEnum
	// setChoice assigns StateProtocolChoiceEnum provided by user to StateProtocol
	setChoice(value StateProtocolChoiceEnum) StateProtocol
	// All returns StateProtocolAll, set in StateProtocol.
	// StateProtocolAll is sets all configured protocols to `start` or `stop` state.
	// Setting protocol state to `start` shall be a no-op if preceding `set_config` API call was made with `config.options.protocol_options.auto_start_all` set to `true` or if all the configured protocols are already started.
	All() StateProtocolAll
	// SetAll assigns StateProtocolAll provided by user to StateProtocol.
	// StateProtocolAll is sets all configured protocols to `start` or `stop` state.
	// Setting protocol state to `start` shall be a no-op if preceding `set_config` API call was made with `config.options.protocol_options.auto_start_all` set to `true` or if all the configured protocols are already started.
	SetAll(value StateProtocolAll) StateProtocol
	// HasAll checks if All has been set in StateProtocol
	HasAll() bool
	// Route returns StateProtocolRoute, set in StateProtocol.
	// StateProtocolRoute is sets the state of configured routes
	Route() StateProtocolRoute
	// SetRoute assigns StateProtocolRoute provided by user to StateProtocol.
	// StateProtocolRoute is sets the state of configured routes
	SetRoute(value StateProtocolRoute) StateProtocol
	// HasRoute checks if Route has been set in StateProtocol
	HasRoute() bool
	// Lacp returns StateProtocolLacp, set in StateProtocol.
	// StateProtocolLacp is sets state of configured LACP
	Lacp() StateProtocolLacp
	// SetLacp assigns StateProtocolLacp provided by user to StateProtocol.
	// StateProtocolLacp is sets state of configured LACP
	SetLacp(value StateProtocolLacp) StateProtocol
	// HasLacp checks if Lacp has been set in StateProtocol
	HasLacp() bool
	// Bgp returns StateProtocolBgp, set in StateProtocol.
	// StateProtocolBgp is sets state of configured BGP peers.
	Bgp() StateProtocolBgp
	// SetBgp assigns StateProtocolBgp provided by user to StateProtocol.
	// StateProtocolBgp is sets state of configured BGP peers.
	SetBgp(value StateProtocolBgp) StateProtocol
	// HasBgp checks if Bgp has been set in StateProtocol
	HasBgp() bool
	// Isis returns StateProtocolIsis, set in StateProtocol.
	// StateProtocolIsis is sets state of configured ISIS routers.
	Isis() StateProtocolIsis
	// SetIsis assigns StateProtocolIsis provided by user to StateProtocol.
	// StateProtocolIsis is sets state of configured ISIS routers.
	SetIsis(value StateProtocolIsis) StateProtocol
	// HasIsis checks if Isis has been set in StateProtocol
	HasIsis() bool
	// Ospfv2 returns StateProtocolOspfv2, set in StateProtocol.
	// StateProtocolOspfv2 is sets state of configured OSPFv2 routers.
	Ospfv2() StateProtocolOspfv2
	// SetOspfv2 assigns StateProtocolOspfv2 provided by user to StateProtocol.
	// StateProtocolOspfv2 is sets state of configured OSPFv2 routers.
	SetOspfv2(value StateProtocolOspfv2) StateProtocol
	// HasOspfv2 checks if Ospfv2 has been set in StateProtocol
	HasOspfv2() bool
	// Ospfv3 returns StateProtocolOspfv3, set in StateProtocol.
	// StateProtocolOspfv3 is sets state of configured OSPFv3 routers.
	Ospfv3() StateProtocolOspfv3
	// SetOspfv3 assigns StateProtocolOspfv3 provided by user to StateProtocol.
	// StateProtocolOspfv3 is sets state of configured OSPFv3 routers.
	SetOspfv3(value StateProtocolOspfv3) StateProtocol
	// HasOspfv3 checks if Ospfv3 has been set in StateProtocol
	HasOspfv3() bool
	// Rocev2 returns StateProtocolRocev2, set in StateProtocol.
	// StateProtocolRocev2 is sets state of configured RoCEv2 peers.
	Rocev2() StateProtocolRocev2
	// SetRocev2 assigns StateProtocolRocev2 provided by user to StateProtocol.
	// StateProtocolRocev2 is sets state of configured RoCEv2 peers.
	SetRocev2(value StateProtocolRocev2) StateProtocol
	// HasRocev2 checks if Rocev2 has been set in StateProtocol
	HasRocev2() bool
	setNil()
}

type StateProtocolChoiceEnum string

// Enum of Choice on StateProtocol
var StateProtocolChoice = struct {
	ALL    StateProtocolChoiceEnum
	ROUTE  StateProtocolChoiceEnum
	LACP   StateProtocolChoiceEnum
	BGP    StateProtocolChoiceEnum
	ISIS   StateProtocolChoiceEnum
	OSPFV2 StateProtocolChoiceEnum
	OSPFV3 StateProtocolChoiceEnum
}{
	ALL:    StateProtocolChoiceEnum("all"),
	ROUTE:  StateProtocolChoiceEnum("route"),
	LACP:   StateProtocolChoiceEnum("lacp"),
	BGP:    StateProtocolChoiceEnum("bgp"),
	ISIS:   StateProtocolChoiceEnum("isis"),
	OSPFV2: StateProtocolChoiceEnum("ospfv2"),
	OSPFV3: StateProtocolChoiceEnum("ospfv3"),
}

func (obj *stateProtocol) Choice() StateProtocolChoiceEnum {
	return StateProtocolChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *stateProtocol) setChoice(value StateProtocolChoiceEnum) StateProtocol {
	intValue, ok := otg.StateProtocol_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocol_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Ospfv3 = nil
	obj.ospfv3Holder = nil
	obj.obj.Ospfv2 = nil
	obj.ospfv2Holder = nil
	obj.obj.Isis = nil
	obj.isisHolder = nil
	obj.obj.Bgp = nil
	obj.bgpHolder = nil
	obj.obj.Lacp = nil
	obj.lacpHolder = nil
	obj.obj.Route = nil
	obj.routeHolder = nil
	obj.obj.All = nil
	obj.allHolder = nil

	if value == StateProtocolChoice.ALL {
		obj.obj.All = NewStateProtocolAll().msg()
	}

	if value == StateProtocolChoice.ROUTE {
		obj.obj.Route = NewStateProtocolRoute().msg()
	}

	if value == StateProtocolChoice.LACP {
		obj.obj.Lacp = NewStateProtocolLacp().msg()
	}

	if value == StateProtocolChoice.BGP {
		obj.obj.Bgp = NewStateProtocolBgp().msg()
	}

	if value == StateProtocolChoice.ISIS {
		obj.obj.Isis = NewStateProtocolIsis().msg()
	}

	if value == StateProtocolChoice.OSPFV2 {
		obj.obj.Ospfv2 = NewStateProtocolOspfv2().msg()
	}

	if value == StateProtocolChoice.OSPFV3 {
		obj.obj.Ospfv3 = NewStateProtocolOspfv3().msg()
	}

	return obj
}

// description is TBD
// All returns a StateProtocolAll
func (obj *stateProtocol) All() StateProtocolAll {
	if obj.obj.All == nil {
		obj.setChoice(StateProtocolChoice.ALL)
	}
	if obj.allHolder == nil {
		obj.allHolder = &stateProtocolAll{obj: obj.obj.All}
	}
	return obj.allHolder
}

// description is TBD
// All returns a StateProtocolAll
func (obj *stateProtocol) HasAll() bool {
	return obj.obj.All != nil
}

// description is TBD
// SetAll sets the StateProtocolAll value in the StateProtocol object
func (obj *stateProtocol) SetAll(value StateProtocolAll) StateProtocol {
	obj.setChoice(StateProtocolChoice.ALL)
	obj.allHolder = nil
	obj.obj.All = value.msg()

	return obj
}

// description is TBD
// Route returns a StateProtocolRoute
func (obj *stateProtocol) Route() StateProtocolRoute {
	if obj.obj.Route == nil {
		obj.setChoice(StateProtocolChoice.ROUTE)
	}
	if obj.routeHolder == nil {
		obj.routeHolder = &stateProtocolRoute{obj: obj.obj.Route}
	}
	return obj.routeHolder
}

// description is TBD
// Route returns a StateProtocolRoute
func (obj *stateProtocol) HasRoute() bool {
	return obj.obj.Route != nil
}

// description is TBD
// SetRoute sets the StateProtocolRoute value in the StateProtocol object
func (obj *stateProtocol) SetRoute(value StateProtocolRoute) StateProtocol {
	obj.setChoice(StateProtocolChoice.ROUTE)
	obj.routeHolder = nil
	obj.obj.Route = value.msg()

	return obj
}

// description is TBD
// Lacp returns a StateProtocolLacp
func (obj *stateProtocol) Lacp() StateProtocolLacp {
	if obj.obj.Lacp == nil {
		obj.setChoice(StateProtocolChoice.LACP)
	}
	if obj.lacpHolder == nil {
		obj.lacpHolder = &stateProtocolLacp{obj: obj.obj.Lacp}
	}
	return obj.lacpHolder
}

// description is TBD
// Lacp returns a StateProtocolLacp
func (obj *stateProtocol) HasLacp() bool {
	return obj.obj.Lacp != nil
}

// description is TBD
// SetLacp sets the StateProtocolLacp value in the StateProtocol object
func (obj *stateProtocol) SetLacp(value StateProtocolLacp) StateProtocol {
	obj.setChoice(StateProtocolChoice.LACP)
	obj.lacpHolder = nil
	obj.obj.Lacp = value.msg()

	return obj
}

// description is TBD
// Bgp returns a StateProtocolBgp
func (obj *stateProtocol) Bgp() StateProtocolBgp {
	if obj.obj.Bgp == nil {
		obj.setChoice(StateProtocolChoice.BGP)
	}
	if obj.bgpHolder == nil {
		obj.bgpHolder = &stateProtocolBgp{obj: obj.obj.Bgp}
	}
	return obj.bgpHolder
}

// description is TBD
// Bgp returns a StateProtocolBgp
func (obj *stateProtocol) HasBgp() bool {
	return obj.obj.Bgp != nil
}

// description is TBD
// SetBgp sets the StateProtocolBgp value in the StateProtocol object
func (obj *stateProtocol) SetBgp(value StateProtocolBgp) StateProtocol {
	obj.setChoice(StateProtocolChoice.BGP)
	obj.bgpHolder = nil
	obj.obj.Bgp = value.msg()

	return obj
}

// description is TBD
// Isis returns a StateProtocolIsis
func (obj *stateProtocol) Isis() StateProtocolIsis {
	if obj.obj.Isis == nil {
		obj.setChoice(StateProtocolChoice.ISIS)
	}
	if obj.isisHolder == nil {
		obj.isisHolder = &stateProtocolIsis{obj: obj.obj.Isis}
	}
	return obj.isisHolder
}

// description is TBD
// Isis returns a StateProtocolIsis
func (obj *stateProtocol) HasIsis() bool {
	return obj.obj.Isis != nil
}

// description is TBD
// SetIsis sets the StateProtocolIsis value in the StateProtocol object
func (obj *stateProtocol) SetIsis(value StateProtocolIsis) StateProtocol {
	obj.setChoice(StateProtocolChoice.ISIS)
	obj.isisHolder = nil
	obj.obj.Isis = value.msg()

	return obj
}

// description is TBD
// Ospfv2 returns a StateProtocolOspfv2
func (obj *stateProtocol) Ospfv2() StateProtocolOspfv2 {
	if obj.obj.Ospfv2 == nil {
		obj.setChoice(StateProtocolChoice.OSPFV2)
	}
	if obj.ospfv2Holder == nil {
		obj.ospfv2Holder = &stateProtocolOspfv2{obj: obj.obj.Ospfv2}
	}
	return obj.ospfv2Holder
}

// description is TBD
// Ospfv2 returns a StateProtocolOspfv2
func (obj *stateProtocol) HasOspfv2() bool {
	return obj.obj.Ospfv2 != nil
}

// description is TBD
// SetOspfv2 sets the StateProtocolOspfv2 value in the StateProtocol object
func (obj *stateProtocol) SetOspfv2(value StateProtocolOspfv2) StateProtocol {
	obj.setChoice(StateProtocolChoice.OSPFV2)
	obj.ospfv2Holder = nil
	obj.obj.Ospfv2 = value.msg()

	return obj
}

// description is TBD
// Ospfv3 returns a StateProtocolOspfv3
func (obj *stateProtocol) Ospfv3() StateProtocolOspfv3 {
	if obj.obj.Ospfv3 == nil {
		obj.setChoice(StateProtocolChoice.OSPFV3)
	}
	if obj.ospfv3Holder == nil {
		obj.ospfv3Holder = &stateProtocolOspfv3{obj: obj.obj.Ospfv3}
	}
	return obj.ospfv3Holder
}

// description is TBD
// Ospfv3 returns a StateProtocolOspfv3
func (obj *stateProtocol) HasOspfv3() bool {
	return obj.obj.Ospfv3 != nil
}

// description is TBD
// SetOspfv3 sets the StateProtocolOspfv3 value in the StateProtocol object
func (obj *stateProtocol) SetOspfv3(value StateProtocolOspfv3) StateProtocol {
	obj.setChoice(StateProtocolChoice.OSPFV3)
	obj.ospfv3Holder = nil
	obj.obj.Ospfv3 = value.msg()

	return obj
}

// description is TBD
// Rocev2 returns a StateProtocolRocev2
func (obj *stateProtocol) Rocev2() StateProtocolRocev2 {
	if obj.obj.Rocev2 == nil {
		obj.obj.Rocev2 = NewStateProtocolRocev2().msg()
	}
	if obj.rocev2Holder == nil {
		obj.rocev2Holder = &stateProtocolRocev2{obj: obj.obj.Rocev2}
	}
	return obj.rocev2Holder
}

// description is TBD
// Rocev2 returns a StateProtocolRocev2
func (obj *stateProtocol) HasRocev2() bool {
	return obj.obj.Rocev2 != nil
}

// description is TBD
// SetRocev2 sets the StateProtocolRocev2 value in the StateProtocol object
func (obj *stateProtocol) SetRocev2(value StateProtocolRocev2) StateProtocol {

	obj.rocev2Holder = nil
	obj.obj.Rocev2 = value.msg()

	return obj
}

func (obj *stateProtocol) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface StateProtocol")
	}

	if obj.obj.All != nil {

		obj.All().validateObj(vObj, set_default)
	}

	if obj.obj.Route != nil {

		obj.Route().validateObj(vObj, set_default)
	}

	if obj.obj.Lacp != nil {

		obj.Lacp().validateObj(vObj, set_default)
	}

	if obj.obj.Bgp != nil {

		obj.Bgp().validateObj(vObj, set_default)
	}

	if obj.obj.Isis != nil {

		obj.Isis().validateObj(vObj, set_default)
	}

	if obj.obj.Ospfv2 != nil {

		obj.Ospfv2().validateObj(vObj, set_default)
	}

	if obj.obj.Ospfv3 != nil {

		obj.Ospfv3().validateObj(vObj, set_default)
	}

	if obj.obj.Rocev2 != nil {

		obj.Rocev2().validateObj(vObj, set_default)
	}

}

func (obj *stateProtocol) setDefault() {
	var choices_set int = 0
	var choice StateProtocolChoiceEnum

	if obj.obj.All != nil {
		choices_set += 1
		choice = StateProtocolChoice.ALL
	}

	if obj.obj.Route != nil {
		choices_set += 1
		choice = StateProtocolChoice.ROUTE
	}

	if obj.obj.Lacp != nil {
		choices_set += 1
		choice = StateProtocolChoice.LACP
	}

	if obj.obj.Bgp != nil {
		choices_set += 1
		choice = StateProtocolChoice.BGP
	}

	if obj.obj.Isis != nil {
		choices_set += 1
		choice = StateProtocolChoice.ISIS
	}

	if obj.obj.Ospfv2 != nil {
		choices_set += 1
		choice = StateProtocolChoice.OSPFV2
	}

	if obj.obj.Ospfv3 != nil {
		choices_set += 1
		choice = StateProtocolChoice.OSPFV3
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in StateProtocol")
			}
		} else {
			intVal := otg.StateProtocol_Choice_Enum_value[string(choice)]
			enumValue := otg.StateProtocol_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

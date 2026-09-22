package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocolIsisSimLinks *****
type stateProtocolIsisSimLinks struct {
	validation
	obj          *otg.StateProtocolIsisSimLinks
	marshaller   marshalStateProtocolIsisSimLinks
	unMarshaller unMarshalStateProtocolIsisSimLinks
}

func NewStateProtocolIsisSimLinks() StateProtocolIsisSimLinks {
	obj := stateProtocolIsisSimLinks{obj: &otg.StateProtocolIsisSimLinks{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocolIsisSimLinks) msg() *otg.StateProtocolIsisSimLinks {
	return obj.obj
}

func (obj *stateProtocolIsisSimLinks) setMsg(msg *otg.StateProtocolIsisSimLinks) StateProtocolIsisSimLinks {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocolIsisSimLinks struct {
	obj *stateProtocolIsisSimLinks
}

type marshalStateProtocolIsisSimLinks interface {
	// ToProto marshals StateProtocolIsisSimLinks to protobuf object *otg.StateProtocolIsisSimLinks
	ToProto() (*otg.StateProtocolIsisSimLinks, error)
	// ToPbText marshals StateProtocolIsisSimLinks to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocolIsisSimLinks to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocolIsisSimLinks to JSON text
	ToJson() (string, error)
}

type unMarshalstateProtocolIsisSimLinks struct {
	obj *stateProtocolIsisSimLinks
}

type unMarshalStateProtocolIsisSimLinks interface {
	// FromProto unmarshals StateProtocolIsisSimLinks from protobuf object *otg.StateProtocolIsisSimLinks
	FromProto(msg *otg.StateProtocolIsisSimLinks) (StateProtocolIsisSimLinks, error)
	// FromPbText unmarshals StateProtocolIsisSimLinks from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocolIsisSimLinks from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocolIsisSimLinks from JSON text
	FromJson(value string) error
}

func (obj *stateProtocolIsisSimLinks) Marshal() marshalStateProtocolIsisSimLinks {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocolIsisSimLinks{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocolIsisSimLinks) Unmarshal() unMarshalStateProtocolIsisSimLinks {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocolIsisSimLinks{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocolIsisSimLinks) ToProto() (*otg.StateProtocolIsisSimLinks, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocolIsisSimLinks) FromProto(msg *otg.StateProtocolIsisSimLinks) (StateProtocolIsisSimLinks, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocolIsisSimLinks) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocolIsisSimLinks) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalstateProtocolIsisSimLinks) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocolIsisSimLinks) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalstateProtocolIsisSimLinks) ToJson() (string, error) {
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

func (m *unMarshalstateProtocolIsisSimLinks) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *stateProtocolIsisSimLinks) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocolIsisSimLinks) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocolIsisSimLinks) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocolIsisSimLinks) Clone() (StateProtocolIsisSimLinks, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocolIsisSimLinks()
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

// StateProtocolIsisSimLinks is sets the state of configured one or more Simulated Links (Interfaces)
type StateProtocolIsisSimLinks interface {
	Validation
	// msg marshals StateProtocolIsisSimLinks to protobuf object *otg.StateProtocolIsisSimLinks
	// and doesn't set defaults
	msg() *otg.StateProtocolIsisSimLinks
	// setMsg unmarshals StateProtocolIsisSimLinks from protobuf object *otg.StateProtocolIsisSimLinks
	// and doesn't set defaults
	setMsg(*otg.StateProtocolIsisSimLinks) StateProtocolIsisSimLinks
	// provides marshal interface
	Marshal() marshalStateProtocolIsisSimLinks
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocolIsisSimLinks
	// validate validates StateProtocolIsisSimLinks
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocolIsisSimLinks, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Names returns []string, set in StateProtocolIsisSimLinks.
	Names() []string
	// SetNames assigns []string provided by user to StateProtocolIsisSimLinks
	SetNames(value []string) StateProtocolIsisSimLinks
	// State returns StateProtocolIsisSimLinksStateEnum, set in StateProtocolIsisSimLinks
	State() StateProtocolIsisSimLinksStateEnum
	// SetState assigns StateProtocolIsisSimLinksStateEnum provided by user to StateProtocolIsisSimLinks
	SetState(value StateProtocolIsisSimLinksStateEnum) StateProtocolIsisSimLinks
	// HasState checks if State has been set in StateProtocolIsisSimLinks
	HasState() bool
}

// The names of ISIS Simulated Links to control. If no names are specified then all ISIS Simulated Links in the configuration will be affected..
//
// x-constraint:
// - /components/schemas/Isis.Interface/properties/name
//
// Names returns a []string
func (obj *stateProtocolIsisSimLinks) Names() []string {
	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	return obj.obj.Names
}

// The names of ISIS Simulated Links to control. If no names are specified then all ISIS Simulated Links in the configuration will be affected..
//
// x-constraint:
// - /components/schemas/Isis.Interface/properties/name
//
// SetNames sets the []string value in the StateProtocolIsisSimLinks object
func (obj *stateProtocolIsisSimLinks) SetNames(value []string) StateProtocolIsisSimLinks {

	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	obj.obj.Names = value

	return obj
}

type StateProtocolIsisSimLinksStateEnum string

// Enum of State on StateProtocolIsisSimLinks
var StateProtocolIsisSimLinksState = struct {
	UP   StateProtocolIsisSimLinksStateEnum
	DOWN StateProtocolIsisSimLinksStateEnum
}{
	UP:   StateProtocolIsisSimLinksStateEnum("up"),
	DOWN: StateProtocolIsisSimLinksStateEnum("down"),
}

func (obj *stateProtocolIsisSimLinks) State() StateProtocolIsisSimLinksStateEnum {
	return StateProtocolIsisSimLinksStateEnum(obj.obj.State.Enum().String())
}

// Sets the Control State of one or more Simulated Links to UP or DOWN.
// The state change is applied bidirectionally — a link between IS-IS Routers A and B is affected
// in both directions simultaneously.
// Setting Control State to DOWN transitions the selected Simulated Links to a disconnected state.
// Both the Simulated/Emulated Router hosting the link and the neighboring router at the far end
// will remove the link from the Extended IS Reachability TLV in their next LSP transmission (with an incremented Sequence Number).
// Setting Control State to UP reconnects the selected Simulated Links.
// Both routers will re-advertise the neighbor relationship in their next LSP Update.
//
// Example:
// Suppose Emulated Router ER (680000000003) is connected to Simulated Routers:
// ST1('770000000001'), ST2('770000000002') and ST3('770000000003') in a ring topology.
//
// ER <--> ST1 <--> ST2(A) <--> ST3(B) <--> ER
//
// Before the AB Link Down operation between ST2 & ST3  the neighbors of ST2 & ST3 will be seen
// as in ISIS Link State PDU (LSP) information in Get State
// LSP-ID of ST2: 770000000002-00-00: IS-Reachability Neighbors : ['770000000001', '770000000003']
// LSP-ID of ST3: 770000000003-00-00: IS-Reachability Neighbors : ['770000000002', '770000000001']
//
// After the AB Link Down operation between ST2 & ST3  the neighbors of ST2 & ST3 will be seen
// as in ISIS Link State PDU (LSP) information in Get State
// LSP-ID of ST2: 770000000002-00-00: IS-Reachability Neighbors : ['770000000001']
// LSP-ID of ST3: 770000000003-00-00: IS-Reachability Neighbors : ['770000000001']
// State returns a string
func (obj *stateProtocolIsisSimLinks) HasState() bool {
	return obj.obj.State != nil
}

func (obj *stateProtocolIsisSimLinks) SetState(value StateProtocolIsisSimLinksStateEnum) StateProtocolIsisSimLinks {
	intValue, ok := otg.StateProtocolIsisSimLinks_State_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolIsisSimLinksStateEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocolIsisSimLinks_State_Enum(intValue)
	obj.obj.State = &enumValue

	return obj
}

func (obj *stateProtocolIsisSimLinks) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *stateProtocolIsisSimLinks) setDefault() {

}

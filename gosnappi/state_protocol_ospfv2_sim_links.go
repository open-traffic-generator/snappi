package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocolOspfv2SimLinks *****
type stateProtocolOspfv2SimLinks struct {
	validation
	obj          *otg.StateProtocolOspfv2SimLinks
	marshaller   marshalStateProtocolOspfv2SimLinks
	unMarshaller unMarshalStateProtocolOspfv2SimLinks
}

func NewStateProtocolOspfv2SimLinks() StateProtocolOspfv2SimLinks {
	obj := stateProtocolOspfv2SimLinks{obj: &otg.StateProtocolOspfv2SimLinks{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocolOspfv2SimLinks) msg() *otg.StateProtocolOspfv2SimLinks {
	return obj.obj
}

func (obj *stateProtocolOspfv2SimLinks) setMsg(msg *otg.StateProtocolOspfv2SimLinks) StateProtocolOspfv2SimLinks {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocolOspfv2SimLinks struct {
	obj *stateProtocolOspfv2SimLinks
}

type marshalStateProtocolOspfv2SimLinks interface {
	// ToProto marshals StateProtocolOspfv2SimLinks to protobuf object *otg.StateProtocolOspfv2SimLinks
	ToProto() (*otg.StateProtocolOspfv2SimLinks, error)
	// ToPbText marshals StateProtocolOspfv2SimLinks to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocolOspfv2SimLinks to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocolOspfv2SimLinks to JSON text
	ToJson() (string, error)
}

type unMarshalstateProtocolOspfv2SimLinks struct {
	obj *stateProtocolOspfv2SimLinks
}

type unMarshalStateProtocolOspfv2SimLinks interface {
	// FromProto unmarshals StateProtocolOspfv2SimLinks from protobuf object *otg.StateProtocolOspfv2SimLinks
	FromProto(msg *otg.StateProtocolOspfv2SimLinks) (StateProtocolOspfv2SimLinks, error)
	// FromPbText unmarshals StateProtocolOspfv2SimLinks from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocolOspfv2SimLinks from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocolOspfv2SimLinks from JSON text
	FromJson(value string) error
}

func (obj *stateProtocolOspfv2SimLinks) Marshal() marshalStateProtocolOspfv2SimLinks {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocolOspfv2SimLinks{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocolOspfv2SimLinks) Unmarshal() unMarshalStateProtocolOspfv2SimLinks {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocolOspfv2SimLinks{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocolOspfv2SimLinks) ToProto() (*otg.StateProtocolOspfv2SimLinks, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocolOspfv2SimLinks) FromProto(msg *otg.StateProtocolOspfv2SimLinks) (StateProtocolOspfv2SimLinks, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocolOspfv2SimLinks) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocolOspfv2SimLinks) FromPbText(value string) error {
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

func (m *marshalstateProtocolOspfv2SimLinks) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocolOspfv2SimLinks) FromYaml(value string) error {
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

func (m *marshalstateProtocolOspfv2SimLinks) ToJson() (string, error) {
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

func (m *unMarshalstateProtocolOspfv2SimLinks) FromJson(value string) error {
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

func (obj *stateProtocolOspfv2SimLinks) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocolOspfv2SimLinks) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocolOspfv2SimLinks) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocolOspfv2SimLinks) Clone() (StateProtocolOspfv2SimLinks, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocolOspfv2SimLinks()
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

// StateProtocolOspfv2SimLinks is sets the state of one or more configured OSPFv2 Simulated Links (Interfaces).
type StateProtocolOspfv2SimLinks interface {
	Validation
	// msg marshals StateProtocolOspfv2SimLinks to protobuf object *otg.StateProtocolOspfv2SimLinks
	// and doesn't set defaults
	msg() *otg.StateProtocolOspfv2SimLinks
	// setMsg unmarshals StateProtocolOspfv2SimLinks from protobuf object *otg.StateProtocolOspfv2SimLinks
	// and doesn't set defaults
	setMsg(*otg.StateProtocolOspfv2SimLinks) StateProtocolOspfv2SimLinks
	// provides marshal interface
	Marshal() marshalStateProtocolOspfv2SimLinks
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocolOspfv2SimLinks
	// validate validates StateProtocolOspfv2SimLinks
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocolOspfv2SimLinks, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Names returns []string, set in StateProtocolOspfv2SimLinks.
	Names() []string
	// SetNames assigns []string provided by user to StateProtocolOspfv2SimLinks
	SetNames(value []string) StateProtocolOspfv2SimLinks
	// State returns StateProtocolOspfv2SimLinksStateEnum, set in StateProtocolOspfv2SimLinks
	State() StateProtocolOspfv2SimLinksStateEnum
	// SetState assigns StateProtocolOspfv2SimLinksStateEnum provided by user to StateProtocolOspfv2SimLinks
	SetState(value StateProtocolOspfv2SimLinksStateEnum) StateProtocolOspfv2SimLinks
	// HasState checks if State has been set in StateProtocolOspfv2SimLinks
	HasState() bool
}

// The names of OSPFv2 Simulated Links to control. If no names are specified then all OSPFv2 Simulated Links in the configuration will be affected.
//
// x-constraint:
// - /components/schemas/Ospfv2.Interface/properties/name
//
// Names returns a []string
func (obj *stateProtocolOspfv2SimLinks) Names() []string {
	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	return obj.obj.Names
}

// The names of OSPFv2 Simulated Links to control. If no names are specified then all OSPFv2 Simulated Links in the configuration will be affected.
//
// x-constraint:
// - /components/schemas/Ospfv2.Interface/properties/name
//
// SetNames sets the []string value in the StateProtocolOspfv2SimLinks object
func (obj *stateProtocolOspfv2SimLinks) SetNames(value []string) StateProtocolOspfv2SimLinks {

	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	obj.obj.Names = value

	return obj
}

type StateProtocolOspfv2SimLinksStateEnum string

// Enum of State on StateProtocolOspfv2SimLinks
var StateProtocolOspfv2SimLinksState = struct {
	UP   StateProtocolOspfv2SimLinksStateEnum
	DOWN StateProtocolOspfv2SimLinksStateEnum
}{
	UP:   StateProtocolOspfv2SimLinksStateEnum("up"),
	DOWN: StateProtocolOspfv2SimLinksStateEnum("down"),
}

func (obj *stateProtocolOspfv2SimLinks) State() StateProtocolOspfv2SimLinksStateEnum {
	return StateProtocolOspfv2SimLinksStateEnum(obj.obj.State.Enum().String())
}

// Sets the Control State of one or more Simulated Links to UP or DOWN.
// The state change is applied bidirectionally - a link between OSPFv2 Routers A and B is affected
// in both directions simultaneously.
// Setting Control State to DOWN transitions the selected Simulated Links to a disconnected state.
// Both the Simulated/Emulated Router hosting the link and the neighboring router at the far end
// will remove the link from their Router-LSA (Type 1 LSA, RFC 2328 Section 12.4.1) in their next
// LSA re-origination (with an incremented LS Sequence Number), and flood the updated Router-LSA so
// all routers in the area recompute the SPF tree.
// Setting Control State to UP reconnects the selected Simulated Links.
// Both routers will re-advertise the adjacency as a link in their next Router-LSA update.
//
// Example:
// Suppose Emulated Router ER is connected to Simulated Routers:
// ST1, ST2 and ST3 in a ring topology.
//
// ER <--> ST1 <--> ST2(A) <--> ST3(B) <--> ER
//
// Before the AB Link Down operation between ST2 & ST3 the neighbors of ST2 & ST3 will be seen
// as link entries in their Router-LSA (Type 1 LSA) information in Get State:
// Router-LSA of ST2: links to ST1 and ST3
// Router-LSA of ST3: links to ST2 and ER
//
// After the AB Link Down operation between ST2 & ST3 the neighbors of ST2 & ST3 will be seen
// as link entries in their Router-LSA (Type 1 LSA) information in Get State:
// Router-LSA of ST2: links to ST1
// Router-LSA of ST3: links to ER
// State returns a string
func (obj *stateProtocolOspfv2SimLinks) HasState() bool {
	return obj.obj.State != nil
}

func (obj *stateProtocolOspfv2SimLinks) SetState(value StateProtocolOspfv2SimLinksStateEnum) StateProtocolOspfv2SimLinks {
	intValue, ok := otg.StateProtocolOspfv2SimLinks_State_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolOspfv2SimLinksStateEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocolOspfv2SimLinks_State_Enum(intValue)
	obj.obj.State = &enumValue

	return obj
}

func (obj *stateProtocolOspfv2SimLinks) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *stateProtocolOspfv2SimLinks) setDefault() {

}

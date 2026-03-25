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
	// Up returns bool, set in StateProtocolIsisSimLinks.
	Up() bool
	// SetUp assigns bool provided by user to StateProtocolIsisSimLinks
	SetUp(value bool) StateProtocolIsisSimLinks
	// HasUp checks if Up has been set in StateProtocolIsisSimLinks
	HasUp() bool
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

// Set up to true or false to interface state in UP or DOWN state respectively in both direction. Thats is the UP/DOWN state is applied for a simulated link,
// connecting, say, a Node A to Node B and Node B to Node A in both directions.
// Setting to false will result in the selected set of Simulated Links to be set in disconnected state. This should result in the link to be removed from the
// Extended IS Reachability TLV in subsequent LSP transmitted ( with incremented Sequence Number ) from both the Simulated or Emulated Router on which this link is
// located as well the Simulated or Emulated Router at the other end of the link i.e. this will bring down the Simulated Links on both ends.
// Similarly, setting to true should result in the neighbor information to be re-introduced in Extended IS Reachability sent in subsequent LSP Update by the
// Router containing the Simulated Link and the neighboring Router at the other end of the link. Note that if the link is already in 'up' state and the value is set
// to 'true' or it is in 'down' state and the value is set to 'false', those links in the request should be ignored.
//
// Example:
// Suppose Emulated Router ER (680000000003) is connected to Simulated Routers:  ST1('770000000001'), ST2('770000000002') and ST3('770000000003') in a ring topology.
// ER <--> ST1 <--> ST2
// ^	            (A) ^
// |		              |
// v 	          (B) v
// -------ST3-------
//
// Before the AB Link Down operation between ST2 & ST3  the neighbors of ST2 & ST3 will be seen as in ISIS Link State PDU (LSP) information in Get State
// LSP-ID of ST2: 770000000002-00-00: IS-Reachability Neighbors : ['770000000001', '770000000003']
// LSP-ID of ST3: 770000000003-00-00: IS-Reachability Neighbors : ['770000000002', '770000000001']
//
// After the AB Link Down operation between ST2 & ST3  the neighbors of ST2 & ST3 will be seen as in ISIS Link State PDU (LSP) information in Get State
// LSP-ID of ST2: 770000000002-00-00: IS-Reachability Neighbors : ['770000000001']
// LSP-ID of ST3: 770000000003-00-00: IS-Reachability Neighbors : ['770000000001']
// Up returns a bool
func (obj *stateProtocolIsisSimLinks) Up() bool {

	return *obj.obj.Up

}

// Set up to true or false to interface state in UP or DOWN state respectively in both direction. Thats is the UP/DOWN state is applied for a simulated link,
// connecting, say, a Node A to Node B and Node B to Node A in both directions.
// Setting to false will result in the selected set of Simulated Links to be set in disconnected state. This should result in the link to be removed from the
// Extended IS Reachability TLV in subsequent LSP transmitted ( with incremented Sequence Number ) from both the Simulated or Emulated Router on which this link is
// located as well the Simulated or Emulated Router at the other end of the link i.e. this will bring down the Simulated Links on both ends.
// Similarly, setting to true should result in the neighbor information to be re-introduced in Extended IS Reachability sent in subsequent LSP Update by the
// Router containing the Simulated Link and the neighboring Router at the other end of the link. Note that if the link is already in 'up' state and the value is set
// to 'true' or it is in 'down' state and the value is set to 'false', those links in the request should be ignored.
//
// Example:
// Suppose Emulated Router ER (680000000003) is connected to Simulated Routers:  ST1('770000000001'), ST2('770000000002') and ST3('770000000003') in a ring topology.
// ER <--> ST1 <--> ST2
// ^	            (A) ^
// |		              |
// v 	          (B) v
// -------ST3-------
//
// Before the AB Link Down operation between ST2 & ST3  the neighbors of ST2 & ST3 will be seen as in ISIS Link State PDU (LSP) information in Get State
// LSP-ID of ST2: 770000000002-00-00: IS-Reachability Neighbors : ['770000000001', '770000000003']
// LSP-ID of ST3: 770000000003-00-00: IS-Reachability Neighbors : ['770000000002', '770000000001']
//
// After the AB Link Down operation between ST2 & ST3  the neighbors of ST2 & ST3 will be seen as in ISIS Link State PDU (LSP) information in Get State
// LSP-ID of ST2: 770000000002-00-00: IS-Reachability Neighbors : ['770000000001']
// LSP-ID of ST3: 770000000003-00-00: IS-Reachability Neighbors : ['770000000001']
// Up returns a bool
func (obj *stateProtocolIsisSimLinks) HasUp() bool {
	return obj.obj.Up != nil
}

// Set up to true or false to interface state in UP or DOWN state respectively in both direction. Thats is the UP/DOWN state is applied for a simulated link,
// connecting, say, a Node A to Node B and Node B to Node A in both directions.
// Setting to false will result in the selected set of Simulated Links to be set in disconnected state. This should result in the link to be removed from the
// Extended IS Reachability TLV in subsequent LSP transmitted ( with incremented Sequence Number ) from both the Simulated or Emulated Router on which this link is
// located as well the Simulated or Emulated Router at the other end of the link i.e. this will bring down the Simulated Links on both ends.
// Similarly, setting to true should result in the neighbor information to be re-introduced in Extended IS Reachability sent in subsequent LSP Update by the
// Router containing the Simulated Link and the neighboring Router at the other end of the link. Note that if the link is already in 'up' state and the value is set
// to 'true' or it is in 'down' state and the value is set to 'false', those links in the request should be ignored.
//
// Example:
// Suppose Emulated Router ER (680000000003) is connected to Simulated Routers:  ST1('770000000001'), ST2('770000000002') and ST3('770000000003') in a ring topology.
// ER <--> ST1 <--> ST2
// ^	            (A) ^
// |		              |
// v 	          (B) v
// -------ST3-------
//
// Before the AB Link Down operation between ST2 & ST3  the neighbors of ST2 & ST3 will be seen as in ISIS Link State PDU (LSP) information in Get State
// LSP-ID of ST2: 770000000002-00-00: IS-Reachability Neighbors : ['770000000001', '770000000003']
// LSP-ID of ST3: 770000000003-00-00: IS-Reachability Neighbors : ['770000000002', '770000000001']
//
// After the AB Link Down operation between ST2 & ST3  the neighbors of ST2 & ST3 will be seen as in ISIS Link State PDU (LSP) information in Get State
// LSP-ID of ST2: 770000000002-00-00: IS-Reachability Neighbors : ['770000000001']
// LSP-ID of ST3: 770000000003-00-00: IS-Reachability Neighbors : ['770000000001']
// SetUp sets the bool value in the StateProtocolIsisSimLinks object
func (obj *stateProtocolIsisSimLinks) SetUp(value bool) StateProtocolIsisSimLinks {

	obj.obj.Up = &value
	return obj
}

func (obj *stateProtocolIsisSimLinks) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *stateProtocolIsisSimLinks) setDefault() {
	if obj.obj.Up == nil {
		obj.SetUp(false)
	}

}

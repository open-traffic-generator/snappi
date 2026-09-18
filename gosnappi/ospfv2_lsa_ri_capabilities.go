package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2LsaRiCapabilities *****
type ospfv2LsaRiCapabilities struct {
	validation
	obj          *otg.Ospfv2LsaRiCapabilities
	marshaller   marshalOspfv2LsaRiCapabilities
	unMarshaller unMarshalOspfv2LsaRiCapabilities
}

func NewOspfv2LsaRiCapabilities() Ospfv2LsaRiCapabilities {
	obj := ospfv2LsaRiCapabilities{obj: &otg.Ospfv2LsaRiCapabilities{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2LsaRiCapabilities) msg() *otg.Ospfv2LsaRiCapabilities {
	return obj.obj
}

func (obj *ospfv2LsaRiCapabilities) setMsg(msg *otg.Ospfv2LsaRiCapabilities) Ospfv2LsaRiCapabilities {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2LsaRiCapabilities struct {
	obj *ospfv2LsaRiCapabilities
}

type marshalOspfv2LsaRiCapabilities interface {
	// ToProto marshals Ospfv2LsaRiCapabilities to protobuf object *otg.Ospfv2LsaRiCapabilities
	ToProto() (*otg.Ospfv2LsaRiCapabilities, error)
	// ToPbText marshals Ospfv2LsaRiCapabilities to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2LsaRiCapabilities to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2LsaRiCapabilities to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2LsaRiCapabilities struct {
	obj *ospfv2LsaRiCapabilities
}

type unMarshalOspfv2LsaRiCapabilities interface {
	// FromProto unmarshals Ospfv2LsaRiCapabilities from protobuf object *otg.Ospfv2LsaRiCapabilities
	FromProto(msg *otg.Ospfv2LsaRiCapabilities) (Ospfv2LsaRiCapabilities, error)
	// FromPbText unmarshals Ospfv2LsaRiCapabilities from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2LsaRiCapabilities from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2LsaRiCapabilities from JSON text
	FromJson(value string) error
}

func (obj *ospfv2LsaRiCapabilities) Marshal() marshalOspfv2LsaRiCapabilities {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2LsaRiCapabilities{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2LsaRiCapabilities) Unmarshal() unMarshalOspfv2LsaRiCapabilities {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2LsaRiCapabilities{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2LsaRiCapabilities) ToProto() (*otg.Ospfv2LsaRiCapabilities, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2LsaRiCapabilities) FromProto(msg *otg.Ospfv2LsaRiCapabilities) (Ospfv2LsaRiCapabilities, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2LsaRiCapabilities) ToPbText() (string, error) {
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

func (m *unMarshalospfv2LsaRiCapabilities) FromPbText(value string) error {
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

func (m *marshalospfv2LsaRiCapabilities) ToYaml() (string, error) {
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

func (m *unMarshalospfv2LsaRiCapabilities) FromYaml(value string) error {
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

func (m *marshalospfv2LsaRiCapabilities) ToJson() (string, error) {
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

func (m *unMarshalospfv2LsaRiCapabilities) FromJson(value string) error {
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

func (obj *ospfv2LsaRiCapabilities) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2LsaRiCapabilities) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2LsaRiCapabilities) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2LsaRiCapabilities) Clone() (Ospfv2LsaRiCapabilities, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2LsaRiCapabilities()
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

// Ospfv2LsaRiCapabilities is the Router Informational Capabilities (TLV type 1) and Router Functional Capabilities
// (TLV type 2) learned from the Router Information (RI) Opaque LSA (RFC 7770 Sections 2.5,
// 2.6).
type Ospfv2LsaRiCapabilities interface {
	Validation
	// msg marshals Ospfv2LsaRiCapabilities to protobuf object *otg.Ospfv2LsaRiCapabilities
	// and doesn't set defaults
	msg() *otg.Ospfv2LsaRiCapabilities
	// setMsg unmarshals Ospfv2LsaRiCapabilities from protobuf object *otg.Ospfv2LsaRiCapabilities
	// and doesn't set defaults
	setMsg(*otg.Ospfv2LsaRiCapabilities) Ospfv2LsaRiCapabilities
	// provides marshal interface
	Marshal() marshalOspfv2LsaRiCapabilities
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2LsaRiCapabilities
	// validate validates Ospfv2LsaRiCapabilities
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2LsaRiCapabilities, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// GracefulRestart returns bool, set in Ospfv2LsaRiCapabilities.
	GracefulRestart() bool
	// SetGracefulRestart assigns bool provided by user to Ospfv2LsaRiCapabilities
	SetGracefulRestart(value bool) Ospfv2LsaRiCapabilities
	// HasGracefulRestart checks if GracefulRestart has been set in Ospfv2LsaRiCapabilities
	HasGracefulRestart() bool
	// GracefulRestartHelper returns bool, set in Ospfv2LsaRiCapabilities.
	GracefulRestartHelper() bool
	// SetGracefulRestartHelper assigns bool provided by user to Ospfv2LsaRiCapabilities
	SetGracefulRestartHelper(value bool) Ospfv2LsaRiCapabilities
	// HasGracefulRestartHelper checks if GracefulRestartHelper has been set in Ospfv2LsaRiCapabilities
	HasGracefulRestartHelper() bool
	// StubRouterSupport returns bool, set in Ospfv2LsaRiCapabilities.
	StubRouterSupport() bool
	// SetStubRouterSupport assigns bool provided by user to Ospfv2LsaRiCapabilities
	SetStubRouterSupport(value bool) Ospfv2LsaRiCapabilities
	// HasStubRouterSupport checks if StubRouterSupport has been set in Ospfv2LsaRiCapabilities
	HasStubRouterSupport() bool
	// TrafficEngineeringSupport returns bool, set in Ospfv2LsaRiCapabilities.
	TrafficEngineeringSupport() bool
	// SetTrafficEngineeringSupport assigns bool provided by user to Ospfv2LsaRiCapabilities
	SetTrafficEngineeringSupport(value bool) Ospfv2LsaRiCapabilities
	// HasTrafficEngineeringSupport checks if TrafficEngineeringSupport has been set in Ospfv2LsaRiCapabilities
	HasTrafficEngineeringSupport() bool
	// PointToPointOverLan returns bool, set in Ospfv2LsaRiCapabilities.
	PointToPointOverLan() bool
	// SetPointToPointOverLan assigns bool provided by user to Ospfv2LsaRiCapabilities
	SetPointToPointOverLan(value bool) Ospfv2LsaRiCapabilities
	// HasPointToPointOverLan checks if PointToPointOverLan has been set in Ospfv2LsaRiCapabilities
	HasPointToPointOverLan() bool
	// ExperimentalTe returns bool, set in Ospfv2LsaRiCapabilities.
	ExperimentalTe() bool
	// SetExperimentalTe assigns bool provided by user to Ospfv2LsaRiCapabilities
	SetExperimentalTe(value bool) Ospfv2LsaRiCapabilities
	// HasExperimentalTe checks if ExperimentalTe has been set in Ospfv2LsaRiCapabilities
	HasExperimentalTe() bool
	// FunctionalCapabilities returns uint32, set in Ospfv2LsaRiCapabilities.
	FunctionalCapabilities() uint32
	// SetFunctionalCapabilities assigns uint32 provided by user to Ospfv2LsaRiCapabilities
	SetFunctionalCapabilities(value uint32) Ospfv2LsaRiCapabilities
	// HasFunctionalCapabilities checks if FunctionalCapabilities has been set in Ospfv2LsaRiCapabilities
	HasFunctionalCapabilities() bool
}

// Informational Capabilities bit 0: the router is Graceful Restart capable (RFC 7770 Section 2.5).
// GracefulRestart returns a bool
func (obj *ospfv2LsaRiCapabilities) GracefulRestart() bool {

	return *obj.obj.GracefulRestart

}

// Informational Capabilities bit 0: the router is Graceful Restart capable (RFC 7770 Section 2.5).
// GracefulRestart returns a bool
func (obj *ospfv2LsaRiCapabilities) HasGracefulRestart() bool {
	return obj.obj.GracefulRestart != nil
}

// Informational Capabilities bit 0: the router is Graceful Restart capable (RFC 7770 Section 2.5).
// SetGracefulRestart sets the bool value in the Ospfv2LsaRiCapabilities object
func (obj *ospfv2LsaRiCapabilities) SetGracefulRestart(value bool) Ospfv2LsaRiCapabilities {

	obj.obj.GracefulRestart = &value
	return obj
}

// Informational Capabilities bit 1: the router is Graceful Restart Helper capable (RFC 7770 Section 2.5).
// GracefulRestartHelper returns a bool
func (obj *ospfv2LsaRiCapabilities) GracefulRestartHelper() bool {

	return *obj.obj.GracefulRestartHelper

}

// Informational Capabilities bit 1: the router is Graceful Restart Helper capable (RFC 7770 Section 2.5).
// GracefulRestartHelper returns a bool
func (obj *ospfv2LsaRiCapabilities) HasGracefulRestartHelper() bool {
	return obj.obj.GracefulRestartHelper != nil
}

// Informational Capabilities bit 1: the router is Graceful Restart Helper capable (RFC 7770 Section 2.5).
// SetGracefulRestartHelper sets the bool value in the Ospfv2LsaRiCapabilities object
func (obj *ospfv2LsaRiCapabilities) SetGracefulRestartHelper(value bool) Ospfv2LsaRiCapabilities {

	obj.obj.GracefulRestartHelper = &value
	return obj
}

// Informational Capabilities bit 2: the router supports OSPF Stub Router functionality (RFC 7770 Section 2.5).
// StubRouterSupport returns a bool
func (obj *ospfv2LsaRiCapabilities) StubRouterSupport() bool {

	return *obj.obj.StubRouterSupport

}

// Informational Capabilities bit 2: the router supports OSPF Stub Router functionality (RFC 7770 Section 2.5).
// StubRouterSupport returns a bool
func (obj *ospfv2LsaRiCapabilities) HasStubRouterSupport() bool {
	return obj.obj.StubRouterSupport != nil
}

// Informational Capabilities bit 2: the router supports OSPF Stub Router functionality (RFC 7770 Section 2.5).
// SetStubRouterSupport sets the bool value in the Ospfv2LsaRiCapabilities object
func (obj *ospfv2LsaRiCapabilities) SetStubRouterSupport(value bool) Ospfv2LsaRiCapabilities {

	obj.obj.StubRouterSupport = &value
	return obj
}

// Informational Capabilities bit 3: the router supports OSPF Traffic Engineering (RFC 7770 Section 2.5).
// TrafficEngineeringSupport returns a bool
func (obj *ospfv2LsaRiCapabilities) TrafficEngineeringSupport() bool {

	return *obj.obj.TrafficEngineeringSupport

}

// Informational Capabilities bit 3: the router supports OSPF Traffic Engineering (RFC 7770 Section 2.5).
// TrafficEngineeringSupport returns a bool
func (obj *ospfv2LsaRiCapabilities) HasTrafficEngineeringSupport() bool {
	return obj.obj.TrafficEngineeringSupport != nil
}

// Informational Capabilities bit 3: the router supports OSPF Traffic Engineering (RFC 7770 Section 2.5).
// SetTrafficEngineeringSupport sets the bool value in the Ospfv2LsaRiCapabilities object
func (obj *ospfv2LsaRiCapabilities) SetTrafficEngineeringSupport(value bool) Ospfv2LsaRiCapabilities {

	obj.obj.TrafficEngineeringSupport = &value
	return obj
}

// Informational Capabilities bit 4: the router supports point-to-point operation over LAN (RFC 7770 Section 2.5).
// PointToPointOverLan returns a bool
func (obj *ospfv2LsaRiCapabilities) PointToPointOverLan() bool {

	return *obj.obj.PointToPointOverLan

}

// Informational Capabilities bit 4: the router supports point-to-point operation over LAN (RFC 7770 Section 2.5).
// PointToPointOverLan returns a bool
func (obj *ospfv2LsaRiCapabilities) HasPointToPointOverLan() bool {
	return obj.obj.PointToPointOverLan != nil
}

// Informational Capabilities bit 4: the router supports point-to-point operation over LAN (RFC 7770 Section 2.5).
// SetPointToPointOverLan sets the bool value in the Ospfv2LsaRiCapabilities object
func (obj *ospfv2LsaRiCapabilities) SetPointToPointOverLan(value bool) Ospfv2LsaRiCapabilities {

	obj.obj.PointToPointOverLan = &value
	return obj
}

// Informational Capabilities bit 5: the router supports Experimental Traffic Engineering (RFC 7770 Section 2.5).
// ExperimentalTe returns a bool
func (obj *ospfv2LsaRiCapabilities) ExperimentalTe() bool {

	return *obj.obj.ExperimentalTe

}

// Informational Capabilities bit 5: the router supports Experimental Traffic Engineering (RFC 7770 Section 2.5).
// ExperimentalTe returns a bool
func (obj *ospfv2LsaRiCapabilities) HasExperimentalTe() bool {
	return obj.obj.ExperimentalTe != nil
}

// Informational Capabilities bit 5: the router supports Experimental Traffic Engineering (RFC 7770 Section 2.5).
// SetExperimentalTe sets the bool value in the Ospfv2LsaRiCapabilities object
func (obj *ospfv2LsaRiCapabilities) SetExperimentalTe(value bool) Ospfv2LsaRiCapabilities {

	obj.obj.ExperimentalTe = &value
	return obj
}

// The raw 32-bit Router Functional Capabilities bitmask (RFC 7770 Section 2.6). No
// capability bits are assigned in this registry as of RFC 7770; returned raw since there
// are no named bits to decode.
// FunctionalCapabilities returns a uint32
func (obj *ospfv2LsaRiCapabilities) FunctionalCapabilities() uint32 {

	return *obj.obj.FunctionalCapabilities

}

// The raw 32-bit Router Functional Capabilities bitmask (RFC 7770 Section 2.6). No
// capability bits are assigned in this registry as of RFC 7770; returned raw since there
// are no named bits to decode.
// FunctionalCapabilities returns a uint32
func (obj *ospfv2LsaRiCapabilities) HasFunctionalCapabilities() bool {
	return obj.obj.FunctionalCapabilities != nil
}

// The raw 32-bit Router Functional Capabilities bitmask (RFC 7770 Section 2.6). No
// capability bits are assigned in this registry as of RFC 7770; returned raw since there
// are no named bits to decode.
// SetFunctionalCapabilities sets the uint32 value in the Ospfv2LsaRiCapabilities object
func (obj *ospfv2LsaRiCapabilities) SetFunctionalCapabilities(value uint32) Ospfv2LsaRiCapabilities {

	obj.obj.FunctionalCapabilities = &value
	return obj
}

func (obj *ospfv2LsaRiCapabilities) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv2LsaRiCapabilities) setDefault() {

}

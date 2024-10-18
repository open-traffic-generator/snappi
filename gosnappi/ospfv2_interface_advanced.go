package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2InterfaceAdvanced *****
type ospfv2InterfaceAdvanced struct {
	validation
	obj          *otg.Ospfv2InterfaceAdvanced
	marshaller   marshalOspfv2InterfaceAdvanced
	unMarshaller unMarshalOspfv2InterfaceAdvanced
}

func NewOspfv2InterfaceAdvanced() Ospfv2InterfaceAdvanced {
	obj := ospfv2InterfaceAdvanced{obj: &otg.Ospfv2InterfaceAdvanced{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2InterfaceAdvanced) msg() *otg.Ospfv2InterfaceAdvanced {
	return obj.obj
}

func (obj *ospfv2InterfaceAdvanced) setMsg(msg *otg.Ospfv2InterfaceAdvanced) Ospfv2InterfaceAdvanced {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2InterfaceAdvanced struct {
	obj *ospfv2InterfaceAdvanced
}

type marshalOspfv2InterfaceAdvanced interface {
	// ToProto marshals Ospfv2InterfaceAdvanced to protobuf object *otg.Ospfv2InterfaceAdvanced
	ToProto() (*otg.Ospfv2InterfaceAdvanced, error)
	// ToPbText marshals Ospfv2InterfaceAdvanced to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2InterfaceAdvanced to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2InterfaceAdvanced to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2InterfaceAdvanced struct {
	obj *ospfv2InterfaceAdvanced
}

type unMarshalOspfv2InterfaceAdvanced interface {
	// FromProto unmarshals Ospfv2InterfaceAdvanced from protobuf object *otg.Ospfv2InterfaceAdvanced
	FromProto(msg *otg.Ospfv2InterfaceAdvanced) (Ospfv2InterfaceAdvanced, error)
	// FromPbText unmarshals Ospfv2InterfaceAdvanced from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2InterfaceAdvanced from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2InterfaceAdvanced from JSON text
	FromJson(value string) error
}

func (obj *ospfv2InterfaceAdvanced) Marshal() marshalOspfv2InterfaceAdvanced {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2InterfaceAdvanced{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2InterfaceAdvanced) Unmarshal() unMarshalOspfv2InterfaceAdvanced {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2InterfaceAdvanced{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2InterfaceAdvanced) ToProto() (*otg.Ospfv2InterfaceAdvanced, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2InterfaceAdvanced) FromProto(msg *otg.Ospfv2InterfaceAdvanced) (Ospfv2InterfaceAdvanced, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2InterfaceAdvanced) ToPbText() (string, error) {
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

func (m *unMarshalospfv2InterfaceAdvanced) FromPbText(value string) error {
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

func (m *marshalospfv2InterfaceAdvanced) ToYaml() (string, error) {
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

func (m *unMarshalospfv2InterfaceAdvanced) FromYaml(value string) error {
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

func (m *marshalospfv2InterfaceAdvanced) ToJson() (string, error) {
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

func (m *unMarshalospfv2InterfaceAdvanced) FromJson(value string) error {
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

func (obj *ospfv2InterfaceAdvanced) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2InterfaceAdvanced) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2InterfaceAdvanced) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2InterfaceAdvanced) Clone() (Ospfv2InterfaceAdvanced, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2InterfaceAdvanced()
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

// Ospfv2InterfaceAdvanced is contains OSPFv2 advanced properties.
type Ospfv2InterfaceAdvanced interface {
	Validation
	// msg marshals Ospfv2InterfaceAdvanced to protobuf object *otg.Ospfv2InterfaceAdvanced
	// and doesn't set defaults
	msg() *otg.Ospfv2InterfaceAdvanced
	// setMsg unmarshals Ospfv2InterfaceAdvanced from protobuf object *otg.Ospfv2InterfaceAdvanced
	// and doesn't set defaults
	setMsg(*otg.Ospfv2InterfaceAdvanced) Ospfv2InterfaceAdvanced
	// provides marshal interface
	Marshal() marshalOspfv2InterfaceAdvanced
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2InterfaceAdvanced
	// validate validates Ospfv2InterfaceAdvanced
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2InterfaceAdvanced, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// HelloInterval returns uint32, set in Ospfv2InterfaceAdvanced.
	HelloInterval() uint32
	// SetHelloInterval assigns uint32 provided by user to Ospfv2InterfaceAdvanced
	SetHelloInterval(value uint32) Ospfv2InterfaceAdvanced
	// HasHelloInterval checks if HelloInterval has been set in Ospfv2InterfaceAdvanced
	HasHelloInterval() bool
	// DeadInterval returns uint32, set in Ospfv2InterfaceAdvanced.
	DeadInterval() uint32
	// SetDeadInterval assigns uint32 provided by user to Ospfv2InterfaceAdvanced
	SetDeadInterval(value uint32) Ospfv2InterfaceAdvanced
	// HasDeadInterval checks if DeadInterval has been set in Ospfv2InterfaceAdvanced
	HasDeadInterval() bool
	// RoutingMetric returns uint32, set in Ospfv2InterfaceAdvanced.
	RoutingMetric() uint32
	// SetRoutingMetric assigns uint32 provided by user to Ospfv2InterfaceAdvanced
	SetRoutingMetric(value uint32) Ospfv2InterfaceAdvanced
	// HasRoutingMetric checks if RoutingMetric has been set in Ospfv2InterfaceAdvanced
	HasRoutingMetric() bool
	// Priority returns uint32, set in Ospfv2InterfaceAdvanced.
	Priority() uint32
	// SetPriority assigns uint32 provided by user to Ospfv2InterfaceAdvanced
	SetPriority(value uint32) Ospfv2InterfaceAdvanced
	// HasPriority checks if Priority has been set in Ospfv2InterfaceAdvanced
	HasPriority() bool
	// ValidateReceivedMtu returns bool, set in Ospfv2InterfaceAdvanced.
	ValidateReceivedMtu() bool
	// SetValidateReceivedMtu assigns bool provided by user to Ospfv2InterfaceAdvanced
	SetValidateReceivedMtu(value bool) Ospfv2InterfaceAdvanced
	// HasValidateReceivedMtu checks if ValidateReceivedMtu has been set in Ospfv2InterfaceAdvanced
	HasValidateReceivedMtu() bool
}

// The time interval, in seconds, between the Hello packets that
// the router sends on the interface. Advertised in Hello packets
// sent out this interface.
// HelloInterval returns a uint32
func (obj *ospfv2InterfaceAdvanced) HelloInterval() uint32 {

	return *obj.obj.HelloInterval

}

// The time interval, in seconds, between the Hello packets that
// the router sends on the interface. Advertised in Hello packets
// sent out this interface.
// HelloInterval returns a uint32
func (obj *ospfv2InterfaceAdvanced) HasHelloInterval() bool {
	return obj.obj.HelloInterval != nil
}

// The time interval, in seconds, between the Hello packets that
// the router sends on the interface. Advertised in Hello packets
// sent out this interface.
// SetHelloInterval sets the uint32 value in the Ospfv2InterfaceAdvanced object
func (obj *ospfv2InterfaceAdvanced) SetHelloInterval(value uint32) Ospfv2InterfaceAdvanced {

	obj.obj.HelloInterval = &value
	return obj
}

// The time interval in seconds before the router's neighbors will declare
// it down, when they stop hearing the router's Hello Packets.
// Advertised in Hello packets sent out this interface.
// DeadInterval returns a uint32
func (obj *ospfv2InterfaceAdvanced) DeadInterval() uint32 {

	return *obj.obj.DeadInterval

}

// The time interval in seconds before the router's neighbors will declare
// it down, when they stop hearing the router's Hello Packets.
// Advertised in Hello packets sent out this interface.
// DeadInterval returns a uint32
func (obj *ospfv2InterfaceAdvanced) HasDeadInterval() bool {
	return obj.obj.DeadInterval != nil
}

// The time interval in seconds before the router's neighbors will declare
// it down, when they stop hearing the router's Hello Packets.
// Advertised in Hello packets sent out this interface.
// SetDeadInterval sets the uint32 value in the Ospfv2InterfaceAdvanced object
func (obj *ospfv2InterfaceAdvanced) SetDeadInterval(value uint32) Ospfv2InterfaceAdvanced {

	obj.obj.DeadInterval = &value
	return obj
}

// Routing metric associated with the interface..
// RoutingMetric returns a uint32
func (obj *ospfv2InterfaceAdvanced) RoutingMetric() uint32 {

	return *obj.obj.RoutingMetric

}

// Routing metric associated with the interface..
// RoutingMetric returns a uint32
func (obj *ospfv2InterfaceAdvanced) HasRoutingMetric() bool {
	return obj.obj.RoutingMetric != nil
}

// Routing metric associated with the interface..
// SetRoutingMetric sets the uint32 value in the Ospfv2InterfaceAdvanced object
func (obj *ospfv2InterfaceAdvanced) SetRoutingMetric(value uint32) Ospfv2InterfaceAdvanced {

	obj.obj.RoutingMetric = &value
	return obj
}

// The Priority for (Backup) Designated Router election.
// This value is used in Hello packets for the Designated Router (DR) election process.
// The default is 0, which indicates that the router will not participate in the DR election process.
// Priority returns a uint32
func (obj *ospfv2InterfaceAdvanced) Priority() uint32 {

	return *obj.obj.Priority

}

// The Priority for (Backup) Designated Router election.
// This value is used in Hello packets for the Designated Router (DR) election process.
// The default is 0, which indicates that the router will not participate in the DR election process.
// Priority returns a uint32
func (obj *ospfv2InterfaceAdvanced) HasPriority() bool {
	return obj.obj.Priority != nil
}

// The Priority for (Backup) Designated Router election.
// This value is used in Hello packets for the Designated Router (DR) election process.
// The default is 0, which indicates that the router will not participate in the DR election process.
// SetPriority sets the uint32 value in the Ospfv2InterfaceAdvanced object
func (obj *ospfv2InterfaceAdvanced) SetPriority(value uint32) Ospfv2InterfaceAdvanced {

	obj.obj.Priority = &value
	return obj
}

// This is for to verify the MTU during the Database (DB) exchange.
// If this is set to true, then the MTU received from the neighbor during Database (DB) Exchange
// will be validated, otherwise it will be ignored.
//
// ValidateReceivedMtu returns a bool
func (obj *ospfv2InterfaceAdvanced) ValidateReceivedMtu() bool {

	return *obj.obj.ValidateReceivedMtu

}

// This is for to verify the MTU during the Database (DB) exchange.
// If this is set to true, then the MTU received from the neighbor during Database (DB) Exchange
// will be validated, otherwise it will be ignored.
//
// ValidateReceivedMtu returns a bool
func (obj *ospfv2InterfaceAdvanced) HasValidateReceivedMtu() bool {
	return obj.obj.ValidateReceivedMtu != nil
}

// This is for to verify the MTU during the Database (DB) exchange.
// If this is set to true, then the MTU received from the neighbor during Database (DB) Exchange
// will be validated, otherwise it will be ignored.
//
// SetValidateReceivedMtu sets the bool value in the Ospfv2InterfaceAdvanced object
func (obj *ospfv2InterfaceAdvanced) SetValidateReceivedMtu(value bool) Ospfv2InterfaceAdvanced {

	obj.obj.ValidateReceivedMtu = &value
	return obj
}

func (obj *ospfv2InterfaceAdvanced) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv2InterfaceAdvanced) setDefault() {
	if obj.obj.HelloInterval == nil {
		obj.SetHelloInterval(10)
	}
	if obj.obj.DeadInterval == nil {
		obj.SetDeadInterval(40)
	}
	if obj.obj.RoutingMetric == nil {
		obj.SetRoutingMetric(0)
	}
	if obj.obj.Priority == nil {
		obj.SetPriority(0)
	}
	if obj.obj.ValidateReceivedMtu == nil {
		obj.SetValidateReceivedMtu(true)
	}

}

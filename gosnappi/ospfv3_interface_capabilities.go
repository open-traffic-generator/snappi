package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3InterfaceCapabilities *****
type ospfv3InterfaceCapabilities struct {
	validation
	obj          *otg.Ospfv3InterfaceCapabilities
	marshaller   marshalOspfv3InterfaceCapabilities
	unMarshaller unMarshalOspfv3InterfaceCapabilities
}

func NewOspfv3InterfaceCapabilities() Ospfv3InterfaceCapabilities {
	obj := ospfv3InterfaceCapabilities{obj: &otg.Ospfv3InterfaceCapabilities{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3InterfaceCapabilities) msg() *otg.Ospfv3InterfaceCapabilities {
	return obj.obj
}

func (obj *ospfv3InterfaceCapabilities) setMsg(msg *otg.Ospfv3InterfaceCapabilities) Ospfv3InterfaceCapabilities {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3InterfaceCapabilities struct {
	obj *ospfv3InterfaceCapabilities
}

type marshalOspfv3InterfaceCapabilities interface {
	// ToProto marshals Ospfv3InterfaceCapabilities to protobuf object *otg.Ospfv3InterfaceCapabilities
	ToProto() (*otg.Ospfv3InterfaceCapabilities, error)
	// ToPbText marshals Ospfv3InterfaceCapabilities to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3InterfaceCapabilities to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3InterfaceCapabilities to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3InterfaceCapabilities struct {
	obj *ospfv3InterfaceCapabilities
}

type unMarshalOspfv3InterfaceCapabilities interface {
	// FromProto unmarshals Ospfv3InterfaceCapabilities from protobuf object *otg.Ospfv3InterfaceCapabilities
	FromProto(msg *otg.Ospfv3InterfaceCapabilities) (Ospfv3InterfaceCapabilities, error)
	// FromPbText unmarshals Ospfv3InterfaceCapabilities from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3InterfaceCapabilities from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3InterfaceCapabilities from JSON text
	FromJson(value string) error
}

func (obj *ospfv3InterfaceCapabilities) Marshal() marshalOspfv3InterfaceCapabilities {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3InterfaceCapabilities{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3InterfaceCapabilities) Unmarshal() unMarshalOspfv3InterfaceCapabilities {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3InterfaceCapabilities{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3InterfaceCapabilities) ToProto() (*otg.Ospfv3InterfaceCapabilities, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3InterfaceCapabilities) FromProto(msg *otg.Ospfv3InterfaceCapabilities) (Ospfv3InterfaceCapabilities, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3InterfaceCapabilities) ToPbText() (string, error) {
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

func (m *unMarshalospfv3InterfaceCapabilities) FromPbText(value string) error {
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

func (m *marshalospfv3InterfaceCapabilities) ToYaml() (string, error) {
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

func (m *unMarshalospfv3InterfaceCapabilities) FromYaml(value string) error {
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

func (m *marshalospfv3InterfaceCapabilities) ToJson() (string, error) {
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

func (m *unMarshalospfv3InterfaceCapabilities) FromJson(value string) error {
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

func (obj *ospfv3InterfaceCapabilities) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3InterfaceCapabilities) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3InterfaceCapabilities) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3InterfaceCapabilities) Clone() (Ospfv3InterfaceCapabilities, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3InterfaceCapabilities()
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

// Ospfv3InterfaceCapabilities is contains OSPFv3 interface capabilities.
type Ospfv3InterfaceCapabilities interface {
	Validation
	// msg marshals Ospfv3InterfaceCapabilities to protobuf object *otg.Ospfv3InterfaceCapabilities
	// and doesn't set defaults
	msg() *otg.Ospfv3InterfaceCapabilities
	// setMsg unmarshals Ospfv3InterfaceCapabilities from protobuf object *otg.Ospfv3InterfaceCapabilities
	// and doesn't set defaults
	setMsg(*otg.Ospfv3InterfaceCapabilities) Ospfv3InterfaceCapabilities
	// provides marshal interface
	Marshal() marshalOspfv3InterfaceCapabilities
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3InterfaceCapabilities
	// validate validates Ospfv3InterfaceCapabilities
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3InterfaceCapabilities, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// HelloInterval returns uint32, set in Ospfv3InterfaceCapabilities.
	HelloInterval() uint32
	// SetHelloInterval assigns uint32 provided by user to Ospfv3InterfaceCapabilities
	SetHelloInterval(value uint32) Ospfv3InterfaceCapabilities
	// HasHelloInterval checks if HelloInterval has been set in Ospfv3InterfaceCapabilities
	HasHelloInterval() bool
	// DeadInterval returns uint32, set in Ospfv3InterfaceCapabilities.
	DeadInterval() uint32
	// SetDeadInterval assigns uint32 provided by user to Ospfv3InterfaceCapabilities
	SetDeadInterval(value uint32) Ospfv3InterfaceCapabilities
	// HasDeadInterval checks if DeadInterval has been set in Ospfv3InterfaceCapabilities
	HasDeadInterval() bool
	// LinkMetric returns uint32, set in Ospfv3InterfaceCapabilities.
	LinkMetric() uint32
	// SetLinkMetric assigns uint32 provided by user to Ospfv3InterfaceCapabilities
	SetLinkMetric(value uint32) Ospfv3InterfaceCapabilities
	// HasLinkMetric checks if LinkMetric has been set in Ospfv3InterfaceCapabilities
	HasLinkMetric() bool
	// ValidateReceivedMtu returns bool, set in Ospfv3InterfaceCapabilities.
	ValidateReceivedMtu() bool
	// SetValidateReceivedMtu assigns bool provided by user to Ospfv3InterfaceCapabilities
	SetValidateReceivedMtu(value bool) Ospfv3InterfaceCapabilities
	// HasValidateReceivedMtu checks if ValidateReceivedMtu has been set in Ospfv3InterfaceCapabilities
	HasValidateReceivedMtu() bool
}

// The time interval, in seconds, between the Hello packets that
// the router sends on the interface. Advertised in Hello packets
// sent out this interface.
// HelloInterval returns a uint32
func (obj *ospfv3InterfaceCapabilities) HelloInterval() uint32 {

	return *obj.obj.HelloInterval

}

// The time interval, in seconds, between the Hello packets that
// the router sends on the interface. Advertised in Hello packets
// sent out this interface.
// HelloInterval returns a uint32
func (obj *ospfv3InterfaceCapabilities) HasHelloInterval() bool {
	return obj.obj.HelloInterval != nil
}

// The time interval, in seconds, between the Hello packets that
// the router sends on the interface. Advertised in Hello packets
// sent out this interface.
// SetHelloInterval sets the uint32 value in the Ospfv3InterfaceCapabilities object
func (obj *ospfv3InterfaceCapabilities) SetHelloInterval(value uint32) Ospfv3InterfaceCapabilities {

	obj.obj.HelloInterval = &value
	return obj
}

// The time interval in seconds before the router's neighbors will declare
// it down, when they stop hearing the router's Hello Packets.
// Advertised in Hello packets sent out this interface.
// DeadInterval returns a uint32
func (obj *ospfv3InterfaceCapabilities) DeadInterval() uint32 {

	return *obj.obj.DeadInterval

}

// The time interval in seconds before the router's neighbors will declare
// it down, when they stop hearing the router's Hello Packets.
// Advertised in Hello packets sent out this interface.
// DeadInterval returns a uint32
func (obj *ospfv3InterfaceCapabilities) HasDeadInterval() bool {
	return obj.obj.DeadInterval != nil
}

// The time interval in seconds before the router's neighbors will declare
// it down, when they stop hearing the router's Hello Packets.
// Advertised in Hello packets sent out this interface.
// SetDeadInterval sets the uint32 value in the Ospfv3InterfaceCapabilities object
func (obj *ospfv3InterfaceCapabilities) SetDeadInterval(value uint32) Ospfv3InterfaceCapabilities {

	obj.obj.DeadInterval = &value
	return obj
}

// The cost of transmitting data on this link.
// LinkMetric returns a uint32
func (obj *ospfv3InterfaceCapabilities) LinkMetric() uint32 {

	return *obj.obj.LinkMetric

}

// The cost of transmitting data on this link.
// LinkMetric returns a uint32
func (obj *ospfv3InterfaceCapabilities) HasLinkMetric() bool {
	return obj.obj.LinkMetric != nil
}

// The cost of transmitting data on this link.
// SetLinkMetric sets the uint32 value in the Ospfv3InterfaceCapabilities object
func (obj *ospfv3InterfaceCapabilities) SetLinkMetric(value uint32) Ospfv3InterfaceCapabilities {

	obj.obj.LinkMetric = &value
	return obj
}

// If this is set to true, then the MTU received from the neighbor during Database (DB) Exchange
// will be validated, otherwise it will be ignored.
//
// ValidateReceivedMtu returns a bool
func (obj *ospfv3InterfaceCapabilities) ValidateReceivedMtu() bool {

	return *obj.obj.ValidateReceivedMtu

}

// If this is set to true, then the MTU received from the neighbor during Database (DB) Exchange
// will be validated, otherwise it will be ignored.
//
// ValidateReceivedMtu returns a bool
func (obj *ospfv3InterfaceCapabilities) HasValidateReceivedMtu() bool {
	return obj.obj.ValidateReceivedMtu != nil
}

// If this is set to true, then the MTU received from the neighbor during Database (DB) Exchange
// will be validated, otherwise it will be ignored.
//
// SetValidateReceivedMtu sets the bool value in the Ospfv3InterfaceCapabilities object
func (obj *ospfv3InterfaceCapabilities) SetValidateReceivedMtu(value bool) Ospfv3InterfaceCapabilities {

	obj.obj.ValidateReceivedMtu = &value
	return obj
}

func (obj *ospfv3InterfaceCapabilities) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.HelloInterval != nil {

		if *obj.obj.HelloInterval > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv3InterfaceCapabilities.HelloInterval <= 65535 but Got %d", *obj.obj.HelloInterval))
		}

	}

	if obj.obj.DeadInterval != nil {

		if *obj.obj.DeadInterval > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv3InterfaceCapabilities.DeadInterval <= 65535 but Got %d", *obj.obj.DeadInterval))
		}

	}

}

func (obj *ospfv3InterfaceCapabilities) setDefault() {
	if obj.obj.HelloInterval == nil {
		obj.SetHelloInterval(10)
	}
	if obj.obj.DeadInterval == nil {
		obj.SetDeadInterval(40)
	}
	if obj.obj.LinkMetric == nil {
		obj.SetLinkMetric(10)
	}
	if obj.obj.ValidateReceivedMtu == nil {
		obj.SetValidateReceivedMtu(true)
	}

}

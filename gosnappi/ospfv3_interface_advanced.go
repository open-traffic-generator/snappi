package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3InterfaceAdvanced *****
type ospfv3InterfaceAdvanced struct {
	validation
	obj          *otg.Ospfv3InterfaceAdvanced
	marshaller   marshalOspfv3InterfaceAdvanced
	unMarshaller unMarshalOspfv3InterfaceAdvanced
}

func NewOspfv3InterfaceAdvanced() Ospfv3InterfaceAdvanced {
	obj := ospfv3InterfaceAdvanced{obj: &otg.Ospfv3InterfaceAdvanced{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3InterfaceAdvanced) msg() *otg.Ospfv3InterfaceAdvanced {
	return obj.obj
}

func (obj *ospfv3InterfaceAdvanced) setMsg(msg *otg.Ospfv3InterfaceAdvanced) Ospfv3InterfaceAdvanced {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3InterfaceAdvanced struct {
	obj *ospfv3InterfaceAdvanced
}

type marshalOspfv3InterfaceAdvanced interface {
	// ToProto marshals Ospfv3InterfaceAdvanced to protobuf object *otg.Ospfv3InterfaceAdvanced
	ToProto() (*otg.Ospfv3InterfaceAdvanced, error)
	// ToPbText marshals Ospfv3InterfaceAdvanced to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3InterfaceAdvanced to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3InterfaceAdvanced to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3InterfaceAdvanced struct {
	obj *ospfv3InterfaceAdvanced
}

type unMarshalOspfv3InterfaceAdvanced interface {
	// FromProto unmarshals Ospfv3InterfaceAdvanced from protobuf object *otg.Ospfv3InterfaceAdvanced
	FromProto(msg *otg.Ospfv3InterfaceAdvanced) (Ospfv3InterfaceAdvanced, error)
	// FromPbText unmarshals Ospfv3InterfaceAdvanced from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3InterfaceAdvanced from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3InterfaceAdvanced from JSON text
	FromJson(value string) error
}

func (obj *ospfv3InterfaceAdvanced) Marshal() marshalOspfv3InterfaceAdvanced {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3InterfaceAdvanced{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3InterfaceAdvanced) Unmarshal() unMarshalOspfv3InterfaceAdvanced {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3InterfaceAdvanced{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3InterfaceAdvanced) ToProto() (*otg.Ospfv3InterfaceAdvanced, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3InterfaceAdvanced) FromProto(msg *otg.Ospfv3InterfaceAdvanced) (Ospfv3InterfaceAdvanced, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3InterfaceAdvanced) ToPbText() (string, error) {
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

func (m *unMarshalospfv3InterfaceAdvanced) FromPbText(value string) error {
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

func (m *marshalospfv3InterfaceAdvanced) ToYaml() (string, error) {
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

func (m *unMarshalospfv3InterfaceAdvanced) FromYaml(value string) error {
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

func (m *marshalospfv3InterfaceAdvanced) ToJson() (string, error) {
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

func (m *unMarshalospfv3InterfaceAdvanced) FromJson(value string) error {
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

func (obj *ospfv3InterfaceAdvanced) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3InterfaceAdvanced) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3InterfaceAdvanced) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3InterfaceAdvanced) Clone() (Ospfv3InterfaceAdvanced, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3InterfaceAdvanced()
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

// Ospfv3InterfaceAdvanced is contains OSPFv3 advanced properties.
type Ospfv3InterfaceAdvanced interface {
	Validation
	// msg marshals Ospfv3InterfaceAdvanced to protobuf object *otg.Ospfv3InterfaceAdvanced
	// and doesn't set defaults
	msg() *otg.Ospfv3InterfaceAdvanced
	// setMsg unmarshals Ospfv3InterfaceAdvanced from protobuf object *otg.Ospfv3InterfaceAdvanced
	// and doesn't set defaults
	setMsg(*otg.Ospfv3InterfaceAdvanced) Ospfv3InterfaceAdvanced
	// provides marshal interface
	Marshal() marshalOspfv3InterfaceAdvanced
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3InterfaceAdvanced
	// validate validates Ospfv3InterfaceAdvanced
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3InterfaceAdvanced, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// HelloInterval returns uint32, set in Ospfv3InterfaceAdvanced.
	HelloInterval() uint32
	// SetHelloInterval assigns uint32 provided by user to Ospfv3InterfaceAdvanced
	SetHelloInterval(value uint32) Ospfv3InterfaceAdvanced
	// HasHelloInterval checks if HelloInterval has been set in Ospfv3InterfaceAdvanced
	HasHelloInterval() bool
	// DeadInterval returns uint32, set in Ospfv3InterfaceAdvanced.
	DeadInterval() uint32
	// SetDeadInterval assigns uint32 provided by user to Ospfv3InterfaceAdvanced
	SetDeadInterval(value uint32) Ospfv3InterfaceAdvanced
	// HasDeadInterval checks if DeadInterval has been set in Ospfv3InterfaceAdvanced
	HasDeadInterval() bool
	// LinkMetric returns uint32, set in Ospfv3InterfaceAdvanced.
	LinkMetric() uint32
	// SetLinkMetric assigns uint32 provided by user to Ospfv3InterfaceAdvanced
	SetLinkMetric(value uint32) Ospfv3InterfaceAdvanced
	// HasLinkMetric checks if LinkMetric has been set in Ospfv3InterfaceAdvanced
	HasLinkMetric() bool
	// ValidateReceivedMtu returns bool, set in Ospfv3InterfaceAdvanced.
	ValidateReceivedMtu() bool
	// SetValidateReceivedMtu assigns bool provided by user to Ospfv3InterfaceAdvanced
	SetValidateReceivedMtu(value bool) Ospfv3InterfaceAdvanced
	// HasValidateReceivedMtu checks if ValidateReceivedMtu has been set in Ospfv3InterfaceAdvanced
	HasValidateReceivedMtu() bool
}

// The time interval, in seconds, between the Hello packets that
// the router sends on the interface. Advertised in Hello packets
// sent out this interface.
// HelloInterval returns a uint32
func (obj *ospfv3InterfaceAdvanced) HelloInterval() uint32 {

	return *obj.obj.HelloInterval

}

// The time interval, in seconds, between the Hello packets that
// the router sends on the interface. Advertised in Hello packets
// sent out this interface.
// HelloInterval returns a uint32
func (obj *ospfv3InterfaceAdvanced) HasHelloInterval() bool {
	return obj.obj.HelloInterval != nil
}

// The time interval, in seconds, between the Hello packets that
// the router sends on the interface. Advertised in Hello packets
// sent out this interface.
// SetHelloInterval sets the uint32 value in the Ospfv3InterfaceAdvanced object
func (obj *ospfv3InterfaceAdvanced) SetHelloInterval(value uint32) Ospfv3InterfaceAdvanced {

	obj.obj.HelloInterval = &value
	return obj
}

// The time interval in seconds before the router's neighbors will declare
// it down, when they stop hearing the router's Hello Packets.
// Advertised in Hello packets sent out this interface.
// DeadInterval returns a uint32
func (obj *ospfv3InterfaceAdvanced) DeadInterval() uint32 {

	return *obj.obj.DeadInterval

}

// The time interval in seconds before the router's neighbors will declare
// it down, when they stop hearing the router's Hello Packets.
// Advertised in Hello packets sent out this interface.
// DeadInterval returns a uint32
func (obj *ospfv3InterfaceAdvanced) HasDeadInterval() bool {
	return obj.obj.DeadInterval != nil
}

// The time interval in seconds before the router's neighbors will declare
// it down, when they stop hearing the router's Hello Packets.
// Advertised in Hello packets sent out this interface.
// SetDeadInterval sets the uint32 value in the Ospfv3InterfaceAdvanced object
func (obj *ospfv3InterfaceAdvanced) SetDeadInterval(value uint32) Ospfv3InterfaceAdvanced {

	obj.obj.DeadInterval = &value
	return obj
}

// The cost of transmitting data on this link.
// LinkMetric returns a uint32
func (obj *ospfv3InterfaceAdvanced) LinkMetric() uint32 {

	return *obj.obj.LinkMetric

}

// The cost of transmitting data on this link.
// LinkMetric returns a uint32
func (obj *ospfv3InterfaceAdvanced) HasLinkMetric() bool {
	return obj.obj.LinkMetric != nil
}

// The cost of transmitting data on this link.
// SetLinkMetric sets the uint32 value in the Ospfv3InterfaceAdvanced object
func (obj *ospfv3InterfaceAdvanced) SetLinkMetric(value uint32) Ospfv3InterfaceAdvanced {

	obj.obj.LinkMetric = &value
	return obj
}

// If this is set to true, then the MTU received from the neighbor during Database (DB) Exchange
// will be validated, otherwise it will be ignored.
//
// ValidateReceivedMtu returns a bool
func (obj *ospfv3InterfaceAdvanced) ValidateReceivedMtu() bool {

	return *obj.obj.ValidateReceivedMtu

}

// If this is set to true, then the MTU received from the neighbor during Database (DB) Exchange
// will be validated, otherwise it will be ignored.
//
// ValidateReceivedMtu returns a bool
func (obj *ospfv3InterfaceAdvanced) HasValidateReceivedMtu() bool {
	return obj.obj.ValidateReceivedMtu != nil
}

// If this is set to true, then the MTU received from the neighbor during Database (DB) Exchange
// will be validated, otherwise it will be ignored.
//
// SetValidateReceivedMtu sets the bool value in the Ospfv3InterfaceAdvanced object
func (obj *ospfv3InterfaceAdvanced) SetValidateReceivedMtu(value bool) Ospfv3InterfaceAdvanced {

	obj.obj.ValidateReceivedMtu = &value
	return obj
}

func (obj *ospfv3InterfaceAdvanced) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.HelloInterval != nil {

		if *obj.obj.HelloInterval > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv3InterfaceAdvanced.HelloInterval <= 65535 but Got %d", *obj.obj.HelloInterval))
		}

	}

	if obj.obj.DeadInterval != nil {

		if *obj.obj.DeadInterval > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv3InterfaceAdvanced.DeadInterval <= 65535 but Got %d", *obj.obj.DeadInterval))
		}

	}

	if obj.obj.LinkMetric != nil {

		if *obj.obj.LinkMetric > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv3InterfaceAdvanced.LinkMetric <= 65535 but Got %d", *obj.obj.LinkMetric))
		}

	}

}

func (obj *ospfv3InterfaceAdvanced) setDefault() {
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

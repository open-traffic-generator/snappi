package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Neighborsv4State *****
type neighborsv4State struct {
	validation
	obj          *otg.Neighborsv4State
	marshaller   marshalNeighborsv4State
	unMarshaller unMarshalNeighborsv4State
}

func NewNeighborsv4State() Neighborsv4State {
	obj := neighborsv4State{obj: &otg.Neighborsv4State{}}
	obj.setDefault()
	return &obj
}

func (obj *neighborsv4State) msg() *otg.Neighborsv4State {
	return obj.obj
}

func (obj *neighborsv4State) setMsg(msg *otg.Neighborsv4State) Neighborsv4State {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalneighborsv4State struct {
	obj *neighborsv4State
}

type marshalNeighborsv4State interface {
	// ToProto marshals Neighborsv4State to protobuf object *otg.Neighborsv4State
	ToProto() (*otg.Neighborsv4State, error)
	// ToPbText marshals Neighborsv4State to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Neighborsv4State to YAML text
	ToYaml() (string, error)
	// ToJson marshals Neighborsv4State to JSON text
	ToJson() (string, error)
}

type unMarshalneighborsv4State struct {
	obj *neighborsv4State
}

type unMarshalNeighborsv4State interface {
	// FromProto unmarshals Neighborsv4State from protobuf object *otg.Neighborsv4State
	FromProto(msg *otg.Neighborsv4State) (Neighborsv4State, error)
	// FromPbText unmarshals Neighborsv4State from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Neighborsv4State from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Neighborsv4State from JSON text
	FromJson(value string) error
}

func (obj *neighborsv4State) Marshal() marshalNeighborsv4State {
	if obj.marshaller == nil {
		obj.marshaller = &marshalneighborsv4State{obj: obj}
	}
	return obj.marshaller
}

func (obj *neighborsv4State) Unmarshal() unMarshalNeighborsv4State {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalneighborsv4State{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalneighborsv4State) ToProto() (*otg.Neighborsv4State, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalneighborsv4State) FromProto(msg *otg.Neighborsv4State) (Neighborsv4State, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalneighborsv4State) ToPbText() (string, error) {
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

func (m *unMarshalneighborsv4State) FromPbText(value string) error {
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

func (m *marshalneighborsv4State) ToYaml() (string, error) {
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

func (m *unMarshalneighborsv4State) FromYaml(value string) error {
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

func (m *marshalneighborsv4State) ToJson() (string, error) {
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

func (m *unMarshalneighborsv4State) FromJson(value string) error {
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

func (obj *neighborsv4State) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *neighborsv4State) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *neighborsv4State) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *neighborsv4State) Clone() (Neighborsv4State, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewNeighborsv4State()
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

// Neighborsv4State is iPv4 Neighbor state (ARP cache entry).
type Neighborsv4State interface {
	Validation
	// msg marshals Neighborsv4State to protobuf object *otg.Neighborsv4State
	// and doesn't set defaults
	msg() *otg.Neighborsv4State
	// setMsg unmarshals Neighborsv4State from protobuf object *otg.Neighborsv4State
	// and doesn't set defaults
	setMsg(*otg.Neighborsv4State) Neighborsv4State
	// provides marshal interface
	Marshal() marshalNeighborsv4State
	// provides unmarshal interface
	Unmarshal() unMarshalNeighborsv4State
	// validate validates Neighborsv4State
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Neighborsv4State, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EthernetName returns string, set in Neighborsv4State.
	EthernetName() string
	// SetEthernetName assigns string provided by user to Neighborsv4State
	SetEthernetName(value string) Neighborsv4State
	// Ipv4Address returns string, set in Neighborsv4State.
	Ipv4Address() string
	// SetIpv4Address assigns string provided by user to Neighborsv4State
	SetIpv4Address(value string) Neighborsv4State
	// LinkLayerAddress returns string, set in Neighborsv4State.
	LinkLayerAddress() string
	// SetLinkLayerAddress assigns string provided by user to Neighborsv4State
	SetLinkLayerAddress(value string) Neighborsv4State
	// HasLinkLayerAddress checks if LinkLayerAddress has been set in Neighborsv4State
	HasLinkLayerAddress() bool
}

// The name of the Ethernet interface associated with the Neighbor state (ARP cache entry).
// EthernetName returns a string
func (obj *neighborsv4State) EthernetName() string {

	return *obj.obj.EthernetName

}

// The name of the Ethernet interface associated with the Neighbor state (ARP cache entry).
// SetEthernetName sets the string value in the Neighborsv4State object
func (obj *neighborsv4State) SetEthernetName(value string) Neighborsv4State {

	obj.obj.EthernetName = &value
	return obj
}

// The IPv4 address of the neighbor.
// Ipv4Address returns a string
func (obj *neighborsv4State) Ipv4Address() string {

	return *obj.obj.Ipv4Address

}

// The IPv4 address of the neighbor.
// SetIpv4Address sets the string value in the Neighborsv4State object
func (obj *neighborsv4State) SetIpv4Address(value string) Neighborsv4State {

	obj.obj.Ipv4Address = &value
	return obj
}

// The link-layer address (MAC) of the neighbor.
// LinkLayerAddress returns a string
func (obj *neighborsv4State) LinkLayerAddress() string {

	return *obj.obj.LinkLayerAddress

}

// The link-layer address (MAC) of the neighbor.
// LinkLayerAddress returns a string
func (obj *neighborsv4State) HasLinkLayerAddress() bool {
	return obj.obj.LinkLayerAddress != nil
}

// The link-layer address (MAC) of the neighbor.
// SetLinkLayerAddress sets the string value in the Neighborsv4State object
func (obj *neighborsv4State) SetLinkLayerAddress(value string) Neighborsv4State {

	obj.obj.LinkLayerAddress = &value
	return obj
}

func (obj *neighborsv4State) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// EthernetName is required
	if obj.obj.EthernetName == nil {
		vObj.validationErrors = append(vObj.validationErrors, "EthernetName is required field on interface Neighborsv4State")
	}

	// Ipv4Address is required
	if obj.obj.Ipv4Address == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv4Address is required field on interface Neighborsv4State")
	}
	if obj.obj.Ipv4Address != nil {

		err := obj.validateIpv4(obj.Ipv4Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Neighborsv4State.Ipv4Address"))
		}

	}

	if obj.obj.LinkLayerAddress != nil {

		err := obj.validateMac(obj.LinkLayerAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Neighborsv4State.LinkLayerAddress"))
		}

	}

}

func (obj *neighborsv4State) setDefault() {

}

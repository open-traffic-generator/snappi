package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Neighborsv6State *****
type neighborsv6State struct {
	validation
	obj          *otg.Neighborsv6State
	marshaller   marshalNeighborsv6State
	unMarshaller unMarshalNeighborsv6State
}

func NewNeighborsv6State() Neighborsv6State {
	obj := neighborsv6State{obj: &otg.Neighborsv6State{}}
	obj.setDefault()
	return &obj
}

func (obj *neighborsv6State) msg() *otg.Neighborsv6State {
	return obj.obj
}

func (obj *neighborsv6State) setMsg(msg *otg.Neighborsv6State) Neighborsv6State {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalneighborsv6State struct {
	obj *neighborsv6State
}

type marshalNeighborsv6State interface {
	// ToProto marshals Neighborsv6State to protobuf object *otg.Neighborsv6State
	ToProto() (*otg.Neighborsv6State, error)
	// ToPbText marshals Neighborsv6State to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Neighborsv6State to YAML text
	ToYaml() (string, error)
	// ToJson marshals Neighborsv6State to JSON text
	ToJson() (string, error)
}

type unMarshalneighborsv6State struct {
	obj *neighborsv6State
}

type unMarshalNeighborsv6State interface {
	// FromProto unmarshals Neighborsv6State from protobuf object *otg.Neighborsv6State
	FromProto(msg *otg.Neighborsv6State) (Neighborsv6State, error)
	// FromPbText unmarshals Neighborsv6State from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Neighborsv6State from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Neighborsv6State from JSON text
	FromJson(value string) error
}

func (obj *neighborsv6State) Marshal() marshalNeighborsv6State {
	if obj.marshaller == nil {
		obj.marshaller = &marshalneighborsv6State{obj: obj}
	}
	return obj.marshaller
}

func (obj *neighborsv6State) Unmarshal() unMarshalNeighborsv6State {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalneighborsv6State{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalneighborsv6State) ToProto() (*otg.Neighborsv6State, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalneighborsv6State) FromProto(msg *otg.Neighborsv6State) (Neighborsv6State, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalneighborsv6State) ToPbText() (string, error) {
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

func (m *unMarshalneighborsv6State) FromPbText(value string) error {
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

func (m *marshalneighborsv6State) ToYaml() (string, error) {
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

func (m *unMarshalneighborsv6State) FromYaml(value string) error {
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

func (m *marshalneighborsv6State) ToJson() (string, error) {
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

func (m *unMarshalneighborsv6State) FromJson(value string) error {
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

func (obj *neighborsv6State) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *neighborsv6State) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *neighborsv6State) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *neighborsv6State) Clone() (Neighborsv6State, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewNeighborsv6State()
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

// Neighborsv6State is iPv6 Neighbor state (NDISC cache entry).
type Neighborsv6State interface {
	Validation
	// msg marshals Neighborsv6State to protobuf object *otg.Neighborsv6State
	// and doesn't set defaults
	msg() *otg.Neighborsv6State
	// setMsg unmarshals Neighborsv6State from protobuf object *otg.Neighborsv6State
	// and doesn't set defaults
	setMsg(*otg.Neighborsv6State) Neighborsv6State
	// provides marshal interface
	Marshal() marshalNeighborsv6State
	// provides unmarshal interface
	Unmarshal() unMarshalNeighborsv6State
	// validate validates Neighborsv6State
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Neighborsv6State, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EthernetName returns string, set in Neighborsv6State.
	EthernetName() string
	// SetEthernetName assigns string provided by user to Neighborsv6State
	SetEthernetName(value string) Neighborsv6State
	// Ipv6Address returns string, set in Neighborsv6State.
	Ipv6Address() string
	// SetIpv6Address assigns string provided by user to Neighborsv6State
	SetIpv6Address(value string) Neighborsv6State
	// LinkLayerAddress returns string, set in Neighborsv6State.
	LinkLayerAddress() string
	// SetLinkLayerAddress assigns string provided by user to Neighborsv6State
	SetLinkLayerAddress(value string) Neighborsv6State
	// HasLinkLayerAddress checks if LinkLayerAddress has been set in Neighborsv6State
	HasLinkLayerAddress() bool
}

// The name of the Ethernet interface associated with the Neighbor state (NDISC cache entry).
// EthernetName returns a string
func (obj *neighborsv6State) EthernetName() string {

	return *obj.obj.EthernetName

}

// The name of the Ethernet interface associated with the Neighbor state (NDISC cache entry).
// SetEthernetName sets the string value in the Neighborsv6State object
func (obj *neighborsv6State) SetEthernetName(value string) Neighborsv6State {

	obj.obj.EthernetName = &value
	return obj
}

// The IPv6 address of the neighbor.
// Ipv6Address returns a string
func (obj *neighborsv6State) Ipv6Address() string {

	return *obj.obj.Ipv6Address

}

// The IPv6 address of the neighbor.
// SetIpv6Address sets the string value in the Neighborsv6State object
func (obj *neighborsv6State) SetIpv6Address(value string) Neighborsv6State {

	obj.obj.Ipv6Address = &value
	return obj
}

// The link-layer address (MAC) of the neighbor.
// LinkLayerAddress returns a string
func (obj *neighborsv6State) LinkLayerAddress() string {

	return *obj.obj.LinkLayerAddress

}

// The link-layer address (MAC) of the neighbor.
// LinkLayerAddress returns a string
func (obj *neighborsv6State) HasLinkLayerAddress() bool {
	return obj.obj.LinkLayerAddress != nil
}

// The link-layer address (MAC) of the neighbor.
// SetLinkLayerAddress sets the string value in the Neighborsv6State object
func (obj *neighborsv6State) SetLinkLayerAddress(value string) Neighborsv6State {

	obj.obj.LinkLayerAddress = &value
	return obj
}

func (obj *neighborsv6State) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// EthernetName is required
	if obj.obj.EthernetName == nil {
		vObj.validationErrors = append(vObj.validationErrors, "EthernetName is required field on interface Neighborsv6State")
	}

	// Ipv6Address is required
	if obj.obj.Ipv6Address == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv6Address is required field on interface Neighborsv6State")
	}
	if obj.obj.Ipv6Address != nil {

		err := obj.validateIpv6(obj.Ipv6Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Neighborsv6State.Ipv6Address"))
		}

	}

	if obj.obj.LinkLayerAddress != nil {

		err := obj.validateMac(obj.LinkLayerAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Neighborsv6State.LinkLayerAddress"))
		}

	}

}

func (obj *neighborsv6State) setDefault() {

}

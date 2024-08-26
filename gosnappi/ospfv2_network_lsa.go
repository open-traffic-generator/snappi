package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2NetworkLsa *****
type ospfv2NetworkLsa struct {
	validation
	obj               *otg.Ospfv2NetworkLsa
	marshaller        marshalOspfv2NetworkLsa
	unMarshaller      unMarshalOspfv2NetworkLsa
	commonAttrsHolder Ospfv2CommonAttrs
}

func NewOspfv2NetworkLsa() Ospfv2NetworkLsa {
	obj := ospfv2NetworkLsa{obj: &otg.Ospfv2NetworkLsa{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2NetworkLsa) msg() *otg.Ospfv2NetworkLsa {
	return obj.obj
}

func (obj *ospfv2NetworkLsa) setMsg(msg *otg.Ospfv2NetworkLsa) Ospfv2NetworkLsa {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2NetworkLsa struct {
	obj *ospfv2NetworkLsa
}

type marshalOspfv2NetworkLsa interface {
	// ToProto marshals Ospfv2NetworkLsa to protobuf object *otg.Ospfv2NetworkLsa
	ToProto() (*otg.Ospfv2NetworkLsa, error)
	// ToPbText marshals Ospfv2NetworkLsa to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2NetworkLsa to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2NetworkLsa to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2NetworkLsa struct {
	obj *ospfv2NetworkLsa
}

type unMarshalOspfv2NetworkLsa interface {
	// FromProto unmarshals Ospfv2NetworkLsa from protobuf object *otg.Ospfv2NetworkLsa
	FromProto(msg *otg.Ospfv2NetworkLsa) (Ospfv2NetworkLsa, error)
	// FromPbText unmarshals Ospfv2NetworkLsa from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2NetworkLsa from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2NetworkLsa from JSON text
	FromJson(value string) error
}

func (obj *ospfv2NetworkLsa) Marshal() marshalOspfv2NetworkLsa {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2NetworkLsa{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2NetworkLsa) Unmarshal() unMarshalOspfv2NetworkLsa {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2NetworkLsa{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2NetworkLsa) ToProto() (*otg.Ospfv2NetworkLsa, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2NetworkLsa) FromProto(msg *otg.Ospfv2NetworkLsa) (Ospfv2NetworkLsa, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2NetworkLsa) ToPbText() (string, error) {
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

func (m *unMarshalospfv2NetworkLsa) FromPbText(value string) error {
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

func (m *marshalospfv2NetworkLsa) ToYaml() (string, error) {
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

func (m *unMarshalospfv2NetworkLsa) FromYaml(value string) error {
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

func (m *marshalospfv2NetworkLsa) ToJson() (string, error) {
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

func (m *unMarshalospfv2NetworkLsa) FromJson(value string) error {
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

func (obj *ospfv2NetworkLsa) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2NetworkLsa) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2NetworkLsa) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2NetworkLsa) Clone() (Ospfv2NetworkLsa, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2NetworkLsa()
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

func (obj *ospfv2NetworkLsa) setNil() {
	obj.commonAttrsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2NetworkLsa is contents of the Network LSA
type Ospfv2NetworkLsa interface {
	Validation
	// msg marshals Ospfv2NetworkLsa to protobuf object *otg.Ospfv2NetworkLsa
	// and doesn't set defaults
	msg() *otg.Ospfv2NetworkLsa
	// setMsg unmarshals Ospfv2NetworkLsa from protobuf object *otg.Ospfv2NetworkLsa
	// and doesn't set defaults
	setMsg(*otg.Ospfv2NetworkLsa) Ospfv2NetworkLsa
	// provides marshal interface
	Marshal() marshalOspfv2NetworkLsa
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2NetworkLsa
	// validate validates Ospfv2NetworkLsa
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2NetworkLsa, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CommonAttrs returns Ospfv2CommonAttrs, set in Ospfv2NetworkLsa.
	// Ospfv2CommonAttrs is tBD
	CommonAttrs() Ospfv2CommonAttrs
	// SetCommonAttrs assigns Ospfv2CommonAttrs provided by user to Ospfv2NetworkLsa.
	// Ospfv2CommonAttrs is tBD
	SetCommonAttrs(value Ospfv2CommonAttrs) Ospfv2NetworkLsa
	// NetworkMask returns string, set in Ospfv2NetworkLsa.
	NetworkMask() string
	// SetNetworkMask assigns string provided by user to Ospfv2NetworkLsa
	SetNetworkMask(value string) Ospfv2NetworkLsa
	// HasNetworkMask checks if NetworkMask has been set in Ospfv2NetworkLsa
	HasNetworkMask() bool
	// NeighborRouterId returns string, set in Ospfv2NetworkLsa.
	NeighborRouterId() string
	// SetNeighborRouterId assigns string provided by user to Ospfv2NetworkLsa
	SetNeighborRouterId(value string) Ospfv2NetworkLsa
	// HasNeighborRouterId checks if NeighborRouterId has been set in Ospfv2NetworkLsa
	HasNeighborRouterId() bool
	setNil()
}

// Contents of the router LSA.
// CommonAttrs returns a Ospfv2CommonAttrs
func (obj *ospfv2NetworkLsa) CommonAttrs() Ospfv2CommonAttrs {
	if obj.obj.CommonAttrs == nil {
		obj.obj.CommonAttrs = NewOspfv2CommonAttrs().msg()
	}
	if obj.commonAttrsHolder == nil {
		obj.commonAttrsHolder = &ospfv2CommonAttrs{obj: obj.obj.CommonAttrs}
	}
	return obj.commonAttrsHolder
}

// Contents of the router LSA.
// SetCommonAttrs sets the Ospfv2CommonAttrs value in the Ospfv2NetworkLsa object
func (obj *ospfv2NetworkLsa) SetCommonAttrs(value Ospfv2CommonAttrs) Ospfv2NetworkLsa {

	obj.commonAttrsHolder = nil
	obj.obj.CommonAttrs = value.msg()

	return obj
}

// The IPv4 address mask for the network.
// NetworkMask returns a string
func (obj *ospfv2NetworkLsa) NetworkMask() string {

	return *obj.obj.NetworkMask

}

// The IPv4 address mask for the network.
// NetworkMask returns a string
func (obj *ospfv2NetworkLsa) HasNetworkMask() bool {
	return obj.obj.NetworkMask != nil
}

// The IPv4 address mask for the network.
// SetNetworkMask sets the string value in the Ospfv2NetworkLsa object
func (obj *ospfv2NetworkLsa) SetNetworkMask(value string) Ospfv2NetworkLsa {

	obj.obj.NetworkMask = &value
	return obj
}

// Neighbor's Router ID for the link.
// NeighborRouterId returns a string
func (obj *ospfv2NetworkLsa) NeighborRouterId() string {

	return *obj.obj.NeighborRouterId

}

// Neighbor's Router ID for the link.
// NeighborRouterId returns a string
func (obj *ospfv2NetworkLsa) HasNeighborRouterId() bool {
	return obj.obj.NeighborRouterId != nil
}

// Neighbor's Router ID for the link.
// SetNeighborRouterId sets the string value in the Ospfv2NetworkLsa object
func (obj *ospfv2NetworkLsa) SetNeighborRouterId(value string) Ospfv2NetworkLsa {

	obj.obj.NeighborRouterId = &value
	return obj
}

func (obj *ospfv2NetworkLsa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// CommonAttrs is required
	if obj.obj.CommonAttrs == nil {
		vObj.validationErrors = append(vObj.validationErrors, "CommonAttrs is required field on interface Ospfv2NetworkLsa")
	}

	if obj.obj.CommonAttrs != nil {

		obj.CommonAttrs().validateObj(vObj, set_default)
	}

	if obj.obj.NetworkMask != nil {

		err := obj.validateIpv4(obj.NetworkMask())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2NetworkLsa.NetworkMask"))
		}

	}

	if obj.obj.NeighborRouterId != nil {

		err := obj.validateIpv4(obj.NeighborRouterId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2NetworkLsa.NeighborRouterId"))
		}

	}

}

func (obj *ospfv2NetworkLsa) setDefault() {

}

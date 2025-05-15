package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3V6RRCapabilities *****
type ospfv3V6RRCapabilities struct {
	validation
	obj                     *otg.Ospfv3V6RRCapabilities
	marshaller              marshalOspfv3V6RRCapabilities
	unMarshaller            unMarshalOspfv3V6RRCapabilities
	forwardingAddressHolder Ospfv3V6RRForwardingAddress
}

func NewOspfv3V6RRCapabilities() Ospfv3V6RRCapabilities {
	obj := ospfv3V6RRCapabilities{obj: &otg.Ospfv3V6RRCapabilities{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3V6RRCapabilities) msg() *otg.Ospfv3V6RRCapabilities {
	return obj.obj
}

func (obj *ospfv3V6RRCapabilities) setMsg(msg *otg.Ospfv3V6RRCapabilities) Ospfv3V6RRCapabilities {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3V6RRCapabilities struct {
	obj *ospfv3V6RRCapabilities
}

type marshalOspfv3V6RRCapabilities interface {
	// ToProto marshals Ospfv3V6RRCapabilities to protobuf object *otg.Ospfv3V6RRCapabilities
	ToProto() (*otg.Ospfv3V6RRCapabilities, error)
	// ToPbText marshals Ospfv3V6RRCapabilities to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3V6RRCapabilities to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3V6RRCapabilities to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3V6RRCapabilities struct {
	obj *ospfv3V6RRCapabilities
}

type unMarshalOspfv3V6RRCapabilities interface {
	// FromProto unmarshals Ospfv3V6RRCapabilities from protobuf object *otg.Ospfv3V6RRCapabilities
	FromProto(msg *otg.Ospfv3V6RRCapabilities) (Ospfv3V6RRCapabilities, error)
	// FromPbText unmarshals Ospfv3V6RRCapabilities from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3V6RRCapabilities from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3V6RRCapabilities from JSON text
	FromJson(value string) error
}

func (obj *ospfv3V6RRCapabilities) Marshal() marshalOspfv3V6RRCapabilities {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3V6RRCapabilities{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3V6RRCapabilities) Unmarshal() unMarshalOspfv3V6RRCapabilities {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3V6RRCapabilities{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3V6RRCapabilities) ToProto() (*otg.Ospfv3V6RRCapabilities, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3V6RRCapabilities) FromProto(msg *otg.Ospfv3V6RRCapabilities) (Ospfv3V6RRCapabilities, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3V6RRCapabilities) ToPbText() (string, error) {
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

func (m *unMarshalospfv3V6RRCapabilities) FromPbText(value string) error {
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

func (m *marshalospfv3V6RRCapabilities) ToYaml() (string, error) {
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

func (m *unMarshalospfv3V6RRCapabilities) FromYaml(value string) error {
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

func (m *marshalospfv3V6RRCapabilities) ToJson() (string, error) {
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

func (m *unMarshalospfv3V6RRCapabilities) FromJson(value string) error {
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

func (obj *ospfv3V6RRCapabilities) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3V6RRCapabilities) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3V6RRCapabilities) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3V6RRCapabilities) Clone() (Ospfv3V6RRCapabilities, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3V6RRCapabilities()
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

func (obj *ospfv3V6RRCapabilities) setNil() {
	obj.forwardingAddressHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv3V6RRCapabilities is container for the capabilities associated with route origin.
type Ospfv3V6RRCapabilities interface {
	Validation
	// msg marshals Ospfv3V6RRCapabilities to protobuf object *otg.Ospfv3V6RRCapabilities
	// and doesn't set defaults
	msg() *otg.Ospfv3V6RRCapabilities
	// setMsg unmarshals Ospfv3V6RRCapabilities from protobuf object *otg.Ospfv3V6RRCapabilities
	// and doesn't set defaults
	setMsg(*otg.Ospfv3V6RRCapabilities) Ospfv3V6RRCapabilities
	// provides marshal interface
	Marshal() marshalOspfv3V6RRCapabilities
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3V6RRCapabilities
	// validate validates Ospfv3V6RRCapabilities
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3V6RRCapabilities, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Propagation returns bool, set in Ospfv3V6RRCapabilities.
	Propagation() bool
	// SetPropagation assigns bool provided by user to Ospfv3V6RRCapabilities
	SetPropagation(value bool) Ospfv3V6RRCapabilities
	// HasPropagation checks if Propagation has been set in Ospfv3V6RRCapabilities
	HasPropagation() bool
	// ForwardingAddress returns Ospfv3V6RRForwardingAddress, set in Ospfv3V6RRCapabilities.
	// Ospfv3V6RRForwardingAddress is container for the forwarding address of NSSA External route origin configuration.
	ForwardingAddress() Ospfv3V6RRForwardingAddress
	// SetForwardingAddress assigns Ospfv3V6RRForwardingAddress provided by user to Ospfv3V6RRCapabilities.
	// Ospfv3V6RRForwardingAddress is container for the forwarding address of NSSA External route origin configuration.
	SetForwardingAddress(value Ospfv3V6RRForwardingAddress) Ospfv3V6RRCapabilities
	// HasForwardingAddress checks if ForwardingAddress has been set in Ospfv3V6RRCapabilities
	HasForwardingAddress() bool
	setNil()
}

// If set, LSAs will be propagated between Areas.
// Propagation returns a bool
func (obj *ospfv3V6RRCapabilities) Propagation() bool {

	return *obj.obj.Propagation

}

// If set, LSAs will be propagated between Areas.
// Propagation returns a bool
func (obj *ospfv3V6RRCapabilities) HasPropagation() bool {
	return obj.obj.Propagation != nil
}

// If set, LSAs will be propagated between Areas.
// SetPropagation sets the bool value in the Ospfv3V6RRCapabilities object
func (obj *ospfv3V6RRCapabilities) SetPropagation(value bool) Ospfv3V6RRCapabilities {

	obj.obj.Propagation = &value
	return obj
}

// Configuration for forwarding address of NSSA External route origin.
// ForwardingAddress returns a Ospfv3V6RRForwardingAddress
func (obj *ospfv3V6RRCapabilities) ForwardingAddress() Ospfv3V6RRForwardingAddress {
	if obj.obj.ForwardingAddress == nil {
		obj.obj.ForwardingAddress = NewOspfv3V6RRForwardingAddress().msg()
	}
	if obj.forwardingAddressHolder == nil {
		obj.forwardingAddressHolder = &ospfv3V6RRForwardingAddress{obj: obj.obj.ForwardingAddress}
	}
	return obj.forwardingAddressHolder
}

// Configuration for forwarding address of NSSA External route origin.
// ForwardingAddress returns a Ospfv3V6RRForwardingAddress
func (obj *ospfv3V6RRCapabilities) HasForwardingAddress() bool {
	return obj.obj.ForwardingAddress != nil
}

// Configuration for forwarding address of NSSA External route origin.
// SetForwardingAddress sets the Ospfv3V6RRForwardingAddress value in the Ospfv3V6RRCapabilities object
func (obj *ospfv3V6RRCapabilities) SetForwardingAddress(value Ospfv3V6RRForwardingAddress) Ospfv3V6RRCapabilities {

	obj.forwardingAddressHolder = nil
	obj.obj.ForwardingAddress = value.msg()

	return obj
}

func (obj *ospfv3V6RRCapabilities) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ForwardingAddress != nil {

		obj.ForwardingAddress().validateObj(vObj, set_default)
	}

}

func (obj *ospfv3V6RRCapabilities) setDefault() {
	if obj.obj.Propagation == nil {
		obj.SetPropagation(false)
	}

}

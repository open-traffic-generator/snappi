package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2InterfaceNeighbor *****
type ospfv2InterfaceNeighbor struct {
	validation
	obj          *otg.Ospfv2InterfaceNeighbor
	marshaller   marshalOspfv2InterfaceNeighbor
	unMarshaller unMarshalOspfv2InterfaceNeighbor
}

func NewOspfv2InterfaceNeighbor() Ospfv2InterfaceNeighbor {
	obj := ospfv2InterfaceNeighbor{obj: &otg.Ospfv2InterfaceNeighbor{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2InterfaceNeighbor) msg() *otg.Ospfv2InterfaceNeighbor {
	return obj.obj
}

func (obj *ospfv2InterfaceNeighbor) setMsg(msg *otg.Ospfv2InterfaceNeighbor) Ospfv2InterfaceNeighbor {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2InterfaceNeighbor struct {
	obj *ospfv2InterfaceNeighbor
}

type marshalOspfv2InterfaceNeighbor interface {
	// ToProto marshals Ospfv2InterfaceNeighbor to protobuf object *otg.Ospfv2InterfaceNeighbor
	ToProto() (*otg.Ospfv2InterfaceNeighbor, error)
	// ToPbText marshals Ospfv2InterfaceNeighbor to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2InterfaceNeighbor to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2InterfaceNeighbor to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2InterfaceNeighbor struct {
	obj *ospfv2InterfaceNeighbor
}

type unMarshalOspfv2InterfaceNeighbor interface {
	// FromProto unmarshals Ospfv2InterfaceNeighbor from protobuf object *otg.Ospfv2InterfaceNeighbor
	FromProto(msg *otg.Ospfv2InterfaceNeighbor) (Ospfv2InterfaceNeighbor, error)
	// FromPbText unmarshals Ospfv2InterfaceNeighbor from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2InterfaceNeighbor from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2InterfaceNeighbor from JSON text
	FromJson(value string) error
}

func (obj *ospfv2InterfaceNeighbor) Marshal() marshalOspfv2InterfaceNeighbor {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2InterfaceNeighbor{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2InterfaceNeighbor) Unmarshal() unMarshalOspfv2InterfaceNeighbor {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2InterfaceNeighbor{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2InterfaceNeighbor) ToProto() (*otg.Ospfv2InterfaceNeighbor, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2InterfaceNeighbor) FromProto(msg *otg.Ospfv2InterfaceNeighbor) (Ospfv2InterfaceNeighbor, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2InterfaceNeighbor) ToPbText() (string, error) {
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

func (m *unMarshalospfv2InterfaceNeighbor) FromPbText(value string) error {
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

func (m *marshalospfv2InterfaceNeighbor) ToYaml() (string, error) {
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

func (m *unMarshalospfv2InterfaceNeighbor) FromYaml(value string) error {
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

func (m *marshalospfv2InterfaceNeighbor) ToJson() (string, error) {
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

func (m *unMarshalospfv2InterfaceNeighbor) FromJson(value string) error {
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

func (obj *ospfv2InterfaceNeighbor) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2InterfaceNeighbor) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2InterfaceNeighbor) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2InterfaceNeighbor) Clone() (Ospfv2InterfaceNeighbor, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2InterfaceNeighbor()
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

// Ospfv2InterfaceNeighbor is configuration of a neighbor.
type Ospfv2InterfaceNeighbor interface {
	Validation
	// msg marshals Ospfv2InterfaceNeighbor to protobuf object *otg.Ospfv2InterfaceNeighbor
	// and doesn't set defaults
	msg() *otg.Ospfv2InterfaceNeighbor
	// setMsg unmarshals Ospfv2InterfaceNeighbor from protobuf object *otg.Ospfv2InterfaceNeighbor
	// and doesn't set defaults
	setMsg(*otg.Ospfv2InterfaceNeighbor) Ospfv2InterfaceNeighbor
	// provides marshal interface
	Marshal() marshalOspfv2InterfaceNeighbor
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2InterfaceNeighbor
	// validate validates Ospfv2InterfaceNeighbor
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2InterfaceNeighbor, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// NeighborIp returns string, set in Ospfv2InterfaceNeighbor.
	NeighborIp() string
	// SetNeighborIp assigns string provided by user to Ospfv2InterfaceNeighbor
	SetNeighborIp(value string) Ospfv2InterfaceNeighbor
	// HasNeighborIp checks if NeighborIp has been set in Ospfv2InterfaceNeighbor
	HasNeighborIp() bool
}

// description is TBD
// NeighborIp returns a string
func (obj *ospfv2InterfaceNeighbor) NeighborIp() string {

	return *obj.obj.NeighborIp

}

// description is TBD
// NeighborIp returns a string
func (obj *ospfv2InterfaceNeighbor) HasNeighborIp() bool {
	return obj.obj.NeighborIp != nil
}

// description is TBD
// SetNeighborIp sets the string value in the Ospfv2InterfaceNeighbor object
func (obj *ospfv2InterfaceNeighbor) SetNeighborIp(value string) Ospfv2InterfaceNeighbor {

	obj.obj.NeighborIp = &value
	return obj
}

func (obj *ospfv2InterfaceNeighbor) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.NeighborIp != nil {

		err := obj.validateIpv4(obj.NeighborIp())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2InterfaceNeighbor.NeighborIp"))
		}

	}

}

func (obj *ospfv2InterfaceNeighbor) setDefault() {

}

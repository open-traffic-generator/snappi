package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2LsaNeighbor *****
type ospfv2LsaNeighbor struct {
	validation
	obj          *otg.Ospfv2LsaNeighbor
	marshaller   marshalOspfv2LsaNeighbor
	unMarshaller unMarshalOspfv2LsaNeighbor
}

func NewOspfv2LsaNeighbor() Ospfv2LsaNeighbor {
	obj := ospfv2LsaNeighbor{obj: &otg.Ospfv2LsaNeighbor{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2LsaNeighbor) msg() *otg.Ospfv2LsaNeighbor {
	return obj.obj
}

func (obj *ospfv2LsaNeighbor) setMsg(msg *otg.Ospfv2LsaNeighbor) Ospfv2LsaNeighbor {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2LsaNeighbor struct {
	obj *ospfv2LsaNeighbor
}

type marshalOspfv2LsaNeighbor interface {
	// ToProto marshals Ospfv2LsaNeighbor to protobuf object *otg.Ospfv2LsaNeighbor
	ToProto() (*otg.Ospfv2LsaNeighbor, error)
	// ToPbText marshals Ospfv2LsaNeighbor to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2LsaNeighbor to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2LsaNeighbor to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2LsaNeighbor struct {
	obj *ospfv2LsaNeighbor
}

type unMarshalOspfv2LsaNeighbor interface {
	// FromProto unmarshals Ospfv2LsaNeighbor from protobuf object *otg.Ospfv2LsaNeighbor
	FromProto(msg *otg.Ospfv2LsaNeighbor) (Ospfv2LsaNeighbor, error)
	// FromPbText unmarshals Ospfv2LsaNeighbor from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2LsaNeighbor from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2LsaNeighbor from JSON text
	FromJson(value string) error
}

func (obj *ospfv2LsaNeighbor) Marshal() marshalOspfv2LsaNeighbor {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2LsaNeighbor{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2LsaNeighbor) Unmarshal() unMarshalOspfv2LsaNeighbor {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2LsaNeighbor{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2LsaNeighbor) ToProto() (*otg.Ospfv2LsaNeighbor, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2LsaNeighbor) FromProto(msg *otg.Ospfv2LsaNeighbor) (Ospfv2LsaNeighbor, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2LsaNeighbor) ToPbText() (string, error) {
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

func (m *unMarshalospfv2LsaNeighbor) FromPbText(value string) error {
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

func (m *marshalospfv2LsaNeighbor) ToYaml() (string, error) {
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

func (m *unMarshalospfv2LsaNeighbor) FromYaml(value string) error {
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

func (m *marshalospfv2LsaNeighbor) ToJson() (string, error) {
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

func (m *unMarshalospfv2LsaNeighbor) FromJson(value string) error {
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

func (obj *ospfv2LsaNeighbor) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2LsaNeighbor) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2LsaNeighbor) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2LsaNeighbor) Clone() (Ospfv2LsaNeighbor, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2LsaNeighbor()
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

// Ospfv2LsaNeighbor is generic attributes used to identify OSPFv2 Network LSA neighbor(s).
type Ospfv2LsaNeighbor interface {
	Validation
	// msg marshals Ospfv2LsaNeighbor to protobuf object *otg.Ospfv2LsaNeighbor
	// and doesn't set defaults
	msg() *otg.Ospfv2LsaNeighbor
	// setMsg unmarshals Ospfv2LsaNeighbor from protobuf object *otg.Ospfv2LsaNeighbor
	// and doesn't set defaults
	setMsg(*otg.Ospfv2LsaNeighbor) Ospfv2LsaNeighbor
	// provides marshal interface
	Marshal() marshalOspfv2LsaNeighbor
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2LsaNeighbor
	// validate validates Ospfv2LsaNeighbor
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2LsaNeighbor, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterId returns string, set in Ospfv2LsaNeighbor.
	RouterId() string
	// SetRouterId assigns string provided by user to Ospfv2LsaNeighbor
	SetRouterId(value string) Ospfv2LsaNeighbor
	// HasRouterId checks if RouterId has been set in Ospfv2LsaNeighbor
	HasRouterId() bool
}

// Neighbor's Router ID for the link.
// RouterId returns a string
func (obj *ospfv2LsaNeighbor) RouterId() string {

	return *obj.obj.RouterId

}

// Neighbor's Router ID for the link.
// RouterId returns a string
func (obj *ospfv2LsaNeighbor) HasRouterId() bool {
	return obj.obj.RouterId != nil
}

// Neighbor's Router ID for the link.
// SetRouterId sets the string value in the Ospfv2LsaNeighbor object
func (obj *ospfv2LsaNeighbor) SetRouterId(value string) Ospfv2LsaNeighbor {

	obj.obj.RouterId = &value
	return obj
}

func (obj *ospfv2LsaNeighbor) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RouterId != nil {

		err := obj.validateIpv4(obj.RouterId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2LsaNeighbor.RouterId"))
		}

	}

}

func (obj *ospfv2LsaNeighbor) setDefault() {

}

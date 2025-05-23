package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3InterfaceBroadcast *****
type ospfv3InterfaceBroadcast struct {
	validation
	obj          *otg.Ospfv3InterfaceBroadcast
	marshaller   marshalOspfv3InterfaceBroadcast
	unMarshaller unMarshalOspfv3InterfaceBroadcast
}

func NewOspfv3InterfaceBroadcast() Ospfv3InterfaceBroadcast {
	obj := ospfv3InterfaceBroadcast{obj: &otg.Ospfv3InterfaceBroadcast{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3InterfaceBroadcast) msg() *otg.Ospfv3InterfaceBroadcast {
	return obj.obj
}

func (obj *ospfv3InterfaceBroadcast) setMsg(msg *otg.Ospfv3InterfaceBroadcast) Ospfv3InterfaceBroadcast {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3InterfaceBroadcast struct {
	obj *ospfv3InterfaceBroadcast
}

type marshalOspfv3InterfaceBroadcast interface {
	// ToProto marshals Ospfv3InterfaceBroadcast to protobuf object *otg.Ospfv3InterfaceBroadcast
	ToProto() (*otg.Ospfv3InterfaceBroadcast, error)
	// ToPbText marshals Ospfv3InterfaceBroadcast to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3InterfaceBroadcast to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3InterfaceBroadcast to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv3InterfaceBroadcast to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv3InterfaceBroadcast struct {
	obj *ospfv3InterfaceBroadcast
}

type unMarshalOspfv3InterfaceBroadcast interface {
	// FromProto unmarshals Ospfv3InterfaceBroadcast from protobuf object *otg.Ospfv3InterfaceBroadcast
	FromProto(msg *otg.Ospfv3InterfaceBroadcast) (Ospfv3InterfaceBroadcast, error)
	// FromPbText unmarshals Ospfv3InterfaceBroadcast from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3InterfaceBroadcast from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3InterfaceBroadcast from JSON text
	FromJson(value string) error
}

func (obj *ospfv3InterfaceBroadcast) Marshal() marshalOspfv3InterfaceBroadcast {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3InterfaceBroadcast{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3InterfaceBroadcast) Unmarshal() unMarshalOspfv3InterfaceBroadcast {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3InterfaceBroadcast{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3InterfaceBroadcast) ToProto() (*otg.Ospfv3InterfaceBroadcast, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3InterfaceBroadcast) FromProto(msg *otg.Ospfv3InterfaceBroadcast) (Ospfv3InterfaceBroadcast, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3InterfaceBroadcast) ToPbText() (string, error) {
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

func (m *unMarshalospfv3InterfaceBroadcast) FromPbText(value string) error {
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

func (m *marshalospfv3InterfaceBroadcast) ToYaml() (string, error) {
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

func (m *unMarshalospfv3InterfaceBroadcast) FromYaml(value string) error {
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

func (m *marshalospfv3InterfaceBroadcast) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalospfv3InterfaceBroadcast) ToJson() (string, error) {
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

func (m *unMarshalospfv3InterfaceBroadcast) FromJson(value string) error {
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

func (obj *ospfv3InterfaceBroadcast) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3InterfaceBroadcast) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3InterfaceBroadcast) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3InterfaceBroadcast) Clone() (Ospfv3InterfaceBroadcast, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3InterfaceBroadcast()
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

// Ospfv3InterfaceBroadcast is container for capabilities associated with network type broadcast.
type Ospfv3InterfaceBroadcast interface {
	Validation
	// msg marshals Ospfv3InterfaceBroadcast to protobuf object *otg.Ospfv3InterfaceBroadcast
	// and doesn't set defaults
	msg() *otg.Ospfv3InterfaceBroadcast
	// setMsg unmarshals Ospfv3InterfaceBroadcast from protobuf object *otg.Ospfv3InterfaceBroadcast
	// and doesn't set defaults
	setMsg(*otg.Ospfv3InterfaceBroadcast) Ospfv3InterfaceBroadcast
	// provides marshal interface
	Marshal() marshalOspfv3InterfaceBroadcast
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3InterfaceBroadcast
	// validate validates Ospfv3InterfaceBroadcast
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3InterfaceBroadcast, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Priority returns uint32, set in Ospfv3InterfaceBroadcast.
	Priority() uint32
	// SetPriority assigns uint32 provided by user to Ospfv3InterfaceBroadcast
	SetPriority(value uint32) Ospfv3InterfaceBroadcast
	// HasPriority checks if Priority has been set in Ospfv3InterfaceBroadcast
	HasPriority() bool
}

// The Priority for (Backup) Designated Router election.
// This value is used in Hello packets for the Designated Router (DR) election process.
// The default is 0, which indicates that the router will not participate in the DR election process.
// Priority returns a uint32
func (obj *ospfv3InterfaceBroadcast) Priority() uint32 {

	return *obj.obj.Priority

}

// The Priority for (Backup) Designated Router election.
// This value is used in Hello packets for the Designated Router (DR) election process.
// The default is 0, which indicates that the router will not participate in the DR election process.
// Priority returns a uint32
func (obj *ospfv3InterfaceBroadcast) HasPriority() bool {
	return obj.obj.Priority != nil
}

// The Priority for (Backup) Designated Router election.
// This value is used in Hello packets for the Designated Router (DR) election process.
// The default is 0, which indicates that the router will not participate in the DR election process.
// SetPriority sets the uint32 value in the Ospfv3InterfaceBroadcast object
func (obj *ospfv3InterfaceBroadcast) SetPriority(value uint32) Ospfv3InterfaceBroadcast {

	obj.obj.Priority = &value
	return obj
}

func (obj *ospfv3InterfaceBroadcast) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Priority != nil {

		if *obj.obj.Priority > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv3InterfaceBroadcast.Priority <= 255 but Got %d", *obj.obj.Priority))
		}

	}

}

func (obj *ospfv3InterfaceBroadcast) setDefault() {
	if obj.obj.Priority == nil {
		obj.SetPriority(0)
	}

}

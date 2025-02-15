package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagPort *****
type lagPort struct {
	validation
	obj            *otg.LagPort
	marshaller     marshalLagPort
	unMarshaller   unMarshalLagPort
	lacpHolder     LagPortLacp
	ethernetHolder DeviceEthernetBase
}

func NewLagPort() LagPort {
	obj := lagPort{obj: &otg.LagPort{}}
	obj.setDefault()
	return &obj
}

func (obj *lagPort) msg() *otg.LagPort {
	return obj.obj
}

func (obj *lagPort) setMsg(msg *otg.LagPort) LagPort {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagPort struct {
	obj *lagPort
}

type marshalLagPort interface {
	// ToProto marshals LagPort to protobuf object *otg.LagPort
	ToProto() (*otg.LagPort, error)
	// ToPbText marshals LagPort to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagPort to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagPort to JSON text
	ToJson() (string, error)
}

type unMarshallagPort struct {
	obj *lagPort
}

type unMarshalLagPort interface {
	// FromProto unmarshals LagPort from protobuf object *otg.LagPort
	FromProto(msg *otg.LagPort) (LagPort, error)
	// FromPbText unmarshals LagPort from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagPort from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagPort from JSON text
	FromJson(value string) error
}

func (obj *lagPort) Marshal() marshalLagPort {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagPort{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagPort) Unmarshal() unMarshalLagPort {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagPort{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagPort) ToProto() (*otg.LagPort, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagPort) FromProto(msg *otg.LagPort) (LagPort, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagPort) ToPbText() (string, error) {
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

func (m *unMarshallagPort) FromPbText(value string) error {
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

func (m *marshallagPort) ToYaml() (string, error) {
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

func (m *unMarshallagPort) FromYaml(value string) error {
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

func (m *marshallagPort) ToJson() (string, error) {
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

func (m *unMarshallagPort) FromJson(value string) error {
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

func (obj *lagPort) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagPort) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagPort) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagPort) Clone() (LagPort, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagPort()
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

func (obj *lagPort) setNil() {
	obj.lacpHolder = nil
	obj.ethernetHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LagPort is the container for a port's ethernet interface and LAG protocol settings
type LagPort interface {
	Validation
	// msg marshals LagPort to protobuf object *otg.LagPort
	// and doesn't set defaults
	msg() *otg.LagPort
	// setMsg unmarshals LagPort from protobuf object *otg.LagPort
	// and doesn't set defaults
	setMsg(*otg.LagPort) LagPort
	// provides marshal interface
	Marshal() marshalLagPort
	// provides unmarshal interface
	Unmarshal() unMarshalLagPort
	// validate validates LagPort
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagPort, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PortName returns string, set in LagPort.
	PortName() string
	// SetPortName assigns string provided by user to LagPort
	SetPortName(value string) LagPort
	// Lacp returns LagPortLacp, set in LagPort.
	// LagPortLacp is the container for link aggregation control protocol settings of a LAG member (port).
	Lacp() LagPortLacp
	// SetLacp assigns LagPortLacp provided by user to LagPort.
	// LagPortLacp is the container for link aggregation control protocol settings of a LAG member (port).
	SetLacp(value LagPortLacp) LagPort
	// HasLacp checks if Lacp has been set in LagPort
	HasLacp() bool
	// Ethernet returns DeviceEthernetBase, set in LagPort.
	// DeviceEthernetBase is base Ethernet interface.
	Ethernet() DeviceEthernetBase
	// SetEthernet assigns DeviceEthernetBase provided by user to LagPort.
	// DeviceEthernetBase is base Ethernet interface.
	SetEthernet(value DeviceEthernetBase) LagPort
	setNil()
}

// The name of a port object that will be part of the LAG.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// PortName returns a string
func (obj *lagPort) PortName() string {

	return *obj.obj.PortName

}

// The name of a port object that will be part of the LAG.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// SetPortName sets the string value in the LagPort object
func (obj *lagPort) SetPortName(value string) LagPort {

	obj.obj.PortName = &value
	return obj
}

// description is TBD
// Lacp returns a LagPortLacp
func (obj *lagPort) Lacp() LagPortLacp {
	if obj.obj.Lacp == nil {
		obj.obj.Lacp = NewLagPortLacp().msg()
	}
	if obj.lacpHolder == nil {
		obj.lacpHolder = &lagPortLacp{obj: obj.obj.Lacp}
	}
	return obj.lacpHolder
}

// description is TBD
// Lacp returns a LagPortLacp
func (obj *lagPort) HasLacp() bool {
	return obj.obj.Lacp != nil
}

// description is TBD
// SetLacp sets the LagPortLacp value in the LagPort object
func (obj *lagPort) SetLacp(value LagPortLacp) LagPort {

	obj.lacpHolder = nil
	obj.obj.Lacp = value.msg()

	return obj
}

// description is TBD
// Ethernet returns a DeviceEthernetBase
func (obj *lagPort) Ethernet() DeviceEthernetBase {
	if obj.obj.Ethernet == nil {
		obj.obj.Ethernet = NewDeviceEthernetBase().msg()
	}
	if obj.ethernetHolder == nil {
		obj.ethernetHolder = &deviceEthernetBase{obj: obj.obj.Ethernet}
	}
	return obj.ethernetHolder
}

// description is TBD
// SetEthernet sets the DeviceEthernetBase value in the LagPort object
func (obj *lagPort) SetEthernet(value DeviceEthernetBase) LagPort {

	obj.ethernetHolder = nil
	obj.obj.Ethernet = value.msg()

	return obj
}

func (obj *lagPort) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// PortName is required
	if obj.obj.PortName == nil {
		vObj.validationErrors = append(vObj.validationErrors, "PortName is required field on interface LagPort")
	}

	if obj.obj.Lacp != nil {

		obj.Lacp().validateObj(vObj, set_default)
	}

	// Ethernet is required
	if obj.obj.Ethernet == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ethernet is required field on interface LagPort")
	}

	if obj.obj.Ethernet != nil {

		obj.Ethernet().validateObj(vObj, set_default)
	}

}

func (obj *lagPort) setDefault() {

}

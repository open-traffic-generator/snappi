package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceMkaEthernetInterface *****
type deviceMkaEthernetInterface struct {
	validation
	obj          *otg.DeviceMkaEthernetInterface
	marshaller   marshalDeviceMkaEthernetInterface
	unMarshaller unMarshalDeviceMkaEthernetInterface
	kayHolder    Mka
}

func NewDeviceMkaEthernetInterface() DeviceMkaEthernetInterface {
	obj := deviceMkaEthernetInterface{obj: &otg.DeviceMkaEthernetInterface{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceMkaEthernetInterface) msg() *otg.DeviceMkaEthernetInterface {
	return obj.obj
}

func (obj *deviceMkaEthernetInterface) setMsg(msg *otg.DeviceMkaEthernetInterface) DeviceMkaEthernetInterface {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceMkaEthernetInterface struct {
	obj *deviceMkaEthernetInterface
}

type marshalDeviceMkaEthernetInterface interface {
	// ToProto marshals DeviceMkaEthernetInterface to protobuf object *otg.DeviceMkaEthernetInterface
	ToProto() (*otg.DeviceMkaEthernetInterface, error)
	// ToPbText marshals DeviceMkaEthernetInterface to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceMkaEthernetInterface to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceMkaEthernetInterface to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceMkaEthernetInterface struct {
	obj *deviceMkaEthernetInterface
}

type unMarshalDeviceMkaEthernetInterface interface {
	// FromProto unmarshals DeviceMkaEthernetInterface from protobuf object *otg.DeviceMkaEthernetInterface
	FromProto(msg *otg.DeviceMkaEthernetInterface) (DeviceMkaEthernetInterface, error)
	// FromPbText unmarshals DeviceMkaEthernetInterface from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceMkaEthernetInterface from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceMkaEthernetInterface from JSON text
	FromJson(value string) error
}

func (obj *deviceMkaEthernetInterface) Marshal() marshalDeviceMkaEthernetInterface {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceMkaEthernetInterface{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceMkaEthernetInterface) Unmarshal() unMarshalDeviceMkaEthernetInterface {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceMkaEthernetInterface{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceMkaEthernetInterface) ToProto() (*otg.DeviceMkaEthernetInterface, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceMkaEthernetInterface) FromProto(msg *otg.DeviceMkaEthernetInterface) (DeviceMkaEthernetInterface, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceMkaEthernetInterface) ToPbText() (string, error) {
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

func (m *unMarshaldeviceMkaEthernetInterface) FromPbText(value string) error {
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

func (m *marshaldeviceMkaEthernetInterface) ToYaml() (string, error) {
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

func (m *unMarshaldeviceMkaEthernetInterface) FromYaml(value string) error {
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

func (m *marshaldeviceMkaEthernetInterface) ToJson() (string, error) {
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

func (m *unMarshaldeviceMkaEthernetInterface) FromJson(value string) error {
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

func (obj *deviceMkaEthernetInterface) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceMkaEthernetInterface) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceMkaEthernetInterface) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceMkaEthernetInterface) Clone() (DeviceMkaEthernetInterface, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceMkaEthernetInterface()
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

func (obj *deviceMkaEthernetInterface) setNil() {
	obj.kayHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceMkaEthernetInterface is configuration for single MKA interface.
type DeviceMkaEthernetInterface interface {
	Validation
	// msg marshals DeviceMkaEthernetInterface to protobuf object *otg.DeviceMkaEthernetInterface
	// and doesn't set defaults
	msg() *otg.DeviceMkaEthernetInterface
	// setMsg unmarshals DeviceMkaEthernetInterface from protobuf object *otg.DeviceMkaEthernetInterface
	// and doesn't set defaults
	setMsg(*otg.DeviceMkaEthernetInterface) DeviceMkaEthernetInterface
	// provides marshal interface
	Marshal() marshalDeviceMkaEthernetInterface
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceMkaEthernetInterface
	// validate validates DeviceMkaEthernetInterface
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceMkaEthernetInterface, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EthName returns string, set in DeviceMkaEthernetInterface.
	EthName() string
	// SetEthName assigns string provided by user to DeviceMkaEthernetInterface
	SetEthName(value string) DeviceMkaEthernetInterface
	// Kay returns Mka, set in DeviceMkaEthernetInterface.
	// Mka is configuration of a Key Agreement Entity (KaY).
	Kay() Mka
	// SetKay assigns Mka provided by user to DeviceMkaEthernetInterface.
	// Mka is configuration of a Key Agreement Entity (KaY).
	SetKay(value Mka) DeviceMkaEthernetInterface
	setNil()
}

// The unique name of the Ethernet interface on which MKA is enabled.
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// EthName returns a string
func (obj *deviceMkaEthernetInterface) EthName() string {

	return *obj.obj.EthName

}

// The unique name of the Ethernet interface on which MKA is enabled.
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// SetEthName sets the string value in the DeviceMkaEthernetInterface object
func (obj *deviceMkaEthernetInterface) SetEthName(value string) DeviceMkaEthernetInterface {

	obj.obj.EthName = &value
	return obj
}

// This contains the properties of Key Agreement Entity (KaY).
// Kay returns a Mka
func (obj *deviceMkaEthernetInterface) Kay() Mka {
	if obj.obj.Kay == nil {
		obj.obj.Kay = NewMka().msg()
	}
	if obj.kayHolder == nil {
		obj.kayHolder = &mka{obj: obj.obj.Kay}
	}
	return obj.kayHolder
}

// This contains the properties of Key Agreement Entity (KaY).
// SetKay sets the Mka value in the DeviceMkaEthernetInterface object
func (obj *deviceMkaEthernetInterface) SetKay(value Mka) DeviceMkaEthernetInterface {

	obj.kayHolder = nil
	obj.obj.Kay = value.msg()

	return obj
}

func (obj *deviceMkaEthernetInterface) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// EthName is required
	if obj.obj.EthName == nil {
		vObj.validationErrors = append(vObj.validationErrors, "EthName is required field on interface DeviceMkaEthernetInterface")
	}

	// Kay is required
	if obj.obj.Kay == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Kay is required field on interface DeviceMkaEthernetInterface")
	}

	if obj.obj.Kay != nil {

		obj.Kay().validateObj(vObj, set_default)
	}

}

func (obj *deviceMkaEthernetInterface) setDefault() {

}

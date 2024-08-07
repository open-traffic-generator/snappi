package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceDhcpv6ClientNoDuid *****
type deviceDhcpv6ClientNoDuid struct {
	validation
	obj          *otg.DeviceDhcpv6ClientNoDuid
	marshaller   marshalDeviceDhcpv6ClientNoDuid
	unMarshaller unMarshalDeviceDhcpv6ClientNoDuid
}

func NewDeviceDhcpv6ClientNoDuid() DeviceDhcpv6ClientNoDuid {
	obj := deviceDhcpv6ClientNoDuid{obj: &otg.DeviceDhcpv6ClientNoDuid{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceDhcpv6ClientNoDuid) msg() *otg.DeviceDhcpv6ClientNoDuid {
	return obj.obj
}

func (obj *deviceDhcpv6ClientNoDuid) setMsg(msg *otg.DeviceDhcpv6ClientNoDuid) DeviceDhcpv6ClientNoDuid {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceDhcpv6ClientNoDuid struct {
	obj *deviceDhcpv6ClientNoDuid
}

type marshalDeviceDhcpv6ClientNoDuid interface {
	// ToProto marshals DeviceDhcpv6ClientNoDuid to protobuf object *otg.DeviceDhcpv6ClientNoDuid
	ToProto() (*otg.DeviceDhcpv6ClientNoDuid, error)
	// ToPbText marshals DeviceDhcpv6ClientNoDuid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceDhcpv6ClientNoDuid to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceDhcpv6ClientNoDuid to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceDhcpv6ClientNoDuid struct {
	obj *deviceDhcpv6ClientNoDuid
}

type unMarshalDeviceDhcpv6ClientNoDuid interface {
	// FromProto unmarshals DeviceDhcpv6ClientNoDuid from protobuf object *otg.DeviceDhcpv6ClientNoDuid
	FromProto(msg *otg.DeviceDhcpv6ClientNoDuid) (DeviceDhcpv6ClientNoDuid, error)
	// FromPbText unmarshals DeviceDhcpv6ClientNoDuid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceDhcpv6ClientNoDuid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceDhcpv6ClientNoDuid from JSON text
	FromJson(value string) error
}

func (obj *deviceDhcpv6ClientNoDuid) Marshal() marshalDeviceDhcpv6ClientNoDuid {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceDhcpv6ClientNoDuid{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceDhcpv6ClientNoDuid) Unmarshal() unMarshalDeviceDhcpv6ClientNoDuid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceDhcpv6ClientNoDuid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceDhcpv6ClientNoDuid) ToProto() (*otg.DeviceDhcpv6ClientNoDuid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceDhcpv6ClientNoDuid) FromProto(msg *otg.DeviceDhcpv6ClientNoDuid) (DeviceDhcpv6ClientNoDuid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceDhcpv6ClientNoDuid) ToPbText() (string, error) {
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

func (m *unMarshaldeviceDhcpv6ClientNoDuid) FromPbText(value string) error {
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

func (m *marshaldeviceDhcpv6ClientNoDuid) ToYaml() (string, error) {
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

func (m *unMarshaldeviceDhcpv6ClientNoDuid) FromYaml(value string) error {
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

func (m *marshaldeviceDhcpv6ClientNoDuid) ToJson() (string, error) {
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

func (m *unMarshaldeviceDhcpv6ClientNoDuid) FromJson(value string) error {
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

func (obj *deviceDhcpv6ClientNoDuid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceDhcpv6ClientNoDuid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceDhcpv6ClientNoDuid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceDhcpv6ClientNoDuid) Clone() (DeviceDhcpv6ClientNoDuid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceDhcpv6ClientNoDuid()
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

// DeviceDhcpv6ClientNoDuid is the container for DUID-LL and DUID-LLT.
type DeviceDhcpv6ClientNoDuid interface {
	Validation
	// msg marshals DeviceDhcpv6ClientNoDuid to protobuf object *otg.DeviceDhcpv6ClientNoDuid
	// and doesn't set defaults
	msg() *otg.DeviceDhcpv6ClientNoDuid
	// setMsg unmarshals DeviceDhcpv6ClientNoDuid from protobuf object *otg.DeviceDhcpv6ClientNoDuid
	// and doesn't set defaults
	setMsg(*otg.DeviceDhcpv6ClientNoDuid) DeviceDhcpv6ClientNoDuid
	// provides marshal interface
	Marshal() marshalDeviceDhcpv6ClientNoDuid
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceDhcpv6ClientNoDuid
	// validate validates DeviceDhcpv6ClientNoDuid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceDhcpv6ClientNoDuid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
}

func (obj *deviceDhcpv6ClientNoDuid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *deviceDhcpv6ClientNoDuid) setDefault() {

}

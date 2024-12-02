package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRIP *****
type isisSRIP struct {
	validation
	obj          *otg.IsisSRIP
	marshaller   marshalIsisSRIP
	unMarshaller unMarshalIsisSRIP
}

func NewIsisSRIP() IsisSRIP {
	obj := isisSRIP{obj: &otg.IsisSRIP{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRIP) msg() *otg.IsisSRIP {
	return obj.obj
}

func (obj *isisSRIP) setMsg(msg *otg.IsisSRIP) IsisSRIP {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRIP struct {
	obj *isisSRIP
}

type marshalIsisSRIP interface {
	// ToProto marshals IsisSRIP to protobuf object *otg.IsisSRIP
	ToProto() (*otg.IsisSRIP, error)
	// ToPbText marshals IsisSRIP to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRIP to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRIP to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRIP struct {
	obj *isisSRIP
}

type unMarshalIsisSRIP interface {
	// FromProto unmarshals IsisSRIP from protobuf object *otg.IsisSRIP
	FromProto(msg *otg.IsisSRIP) (IsisSRIP, error)
	// FromPbText unmarshals IsisSRIP from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRIP from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRIP from JSON text
	FromJson(value string) error
}

func (obj *isisSRIP) Marshal() marshalIsisSRIP {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRIP{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRIP) Unmarshal() unMarshalIsisSRIP {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRIP{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRIP) ToProto() (*otg.IsisSRIP, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRIP) FromProto(msg *otg.IsisSRIP) (IsisSRIP, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRIP) ToPbText() (string, error) {
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

func (m *unMarshalisisSRIP) FromPbText(value string) error {
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

func (m *marshalisisSRIP) ToYaml() (string, error) {
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

func (m *unMarshalisisSRIP) FromYaml(value string) error {
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

func (m *marshalisisSRIP) ToJson() (string, error) {
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

func (m *unMarshalisisSRIP) FromJson(value string) error {
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

func (obj *isisSRIP) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRIP) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRIP) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRIP) Clone() (IsisSRIP, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRIP()
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

// IsisSRIP is tBD
type IsisSRIP interface {
	Validation
	// msg marshals IsisSRIP to protobuf object *otg.IsisSRIP
	// and doesn't set defaults
	msg() *otg.IsisSRIP
	// setMsg unmarshals IsisSRIP from protobuf object *otg.IsisSRIP
	// and doesn't set defaults
	setMsg(*otg.IsisSRIP) IsisSRIP
	// provides marshal interface
	Marshal() marshalIsisSRIP
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRIP
	// validate validates IsisSRIP
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRIP, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Address returns string, set in IsisSRIP.
	Address() string
	// SetAddress assigns string provided by user to IsisSRIP
	SetAddress(value string) IsisSRIP
	// HasAddress checks if Address has been set in IsisSRIP
	HasAddress() bool
	// Prefix returns uint32, set in IsisSRIP.
	Prefix() uint32
	// SetPrefix assigns uint32 provided by user to IsisSRIP
	SetPrefix(value uint32) IsisSRIP
	// HasPrefix checks if Prefix has been set in IsisSRIP
	HasPrefix() bool
}

// The starting address of the network.
// Address returns a string
func (obj *isisSRIP) Address() string {

	return *obj.obj.Address

}

// The starting address of the network.
// Address returns a string
func (obj *isisSRIP) HasAddress() bool {
	return obj.obj.Address != nil
}

// The starting address of the network.
// SetAddress sets the string value in the IsisSRIP object
func (obj *isisSRIP) SetAddress(value string) IsisSRIP {

	obj.obj.Address = &value
	return obj
}

// The IPv4 network prefix length to be applied to the address.
// Prefix returns a uint32
func (obj *isisSRIP) Prefix() uint32 {

	return *obj.obj.Prefix

}

// The IPv4 network prefix length to be applied to the address.
// Prefix returns a uint32
func (obj *isisSRIP) HasPrefix() bool {
	return obj.obj.Prefix != nil
}

// The IPv4 network prefix length to be applied to the address.
// SetPrefix sets the uint32 value in the IsisSRIP object
func (obj *isisSRIP) SetPrefix(value uint32) IsisSRIP {

	obj.obj.Prefix = &value
	return obj
}

func (obj *isisSRIP) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Address != nil {

		err := obj.validateIpv4(obj.Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on IsisSRIP.Address"))
		}

	}

	if obj.obj.Prefix != nil {

		if *obj.obj.Prefix > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRIP.Prefix <= 32 but Got %d", *obj.obj.Prefix))
		}

	}

}

func (obj *isisSRIP) setDefault() {
	if obj.obj.Prefix == nil {
		obj.SetPrefix(24)
	}

}

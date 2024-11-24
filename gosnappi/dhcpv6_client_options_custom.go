package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ClientOptionsCustom *****
type dhcpv6ClientOptionsCustom struct {
	validation
	obj          *otg.Dhcpv6ClientOptionsCustom
	marshaller   marshalDhcpv6ClientOptionsCustom
	unMarshaller unMarshalDhcpv6ClientOptionsCustom
}

func NewDhcpv6ClientOptionsCustom() Dhcpv6ClientOptionsCustom {
	obj := dhcpv6ClientOptionsCustom{obj: &otg.Dhcpv6ClientOptionsCustom{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ClientOptionsCustom) msg() *otg.Dhcpv6ClientOptionsCustom {
	return obj.obj
}

func (obj *dhcpv6ClientOptionsCustom) setMsg(msg *otg.Dhcpv6ClientOptionsCustom) Dhcpv6ClientOptionsCustom {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ClientOptionsCustom struct {
	obj *dhcpv6ClientOptionsCustom
}

type marshalDhcpv6ClientOptionsCustom interface {
	// ToProto marshals Dhcpv6ClientOptionsCustom to protobuf object *otg.Dhcpv6ClientOptionsCustom
	ToProto() (*otg.Dhcpv6ClientOptionsCustom, error)
	// ToPbText marshals Dhcpv6ClientOptionsCustom to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ClientOptionsCustom to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ClientOptionsCustom to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv6ClientOptionsCustom to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv6ClientOptionsCustom struct {
	obj *dhcpv6ClientOptionsCustom
}

type unMarshalDhcpv6ClientOptionsCustom interface {
	// FromProto unmarshals Dhcpv6ClientOptionsCustom from protobuf object *otg.Dhcpv6ClientOptionsCustom
	FromProto(msg *otg.Dhcpv6ClientOptionsCustom) (Dhcpv6ClientOptionsCustom, error)
	// FromPbText unmarshals Dhcpv6ClientOptionsCustom from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ClientOptionsCustom from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ClientOptionsCustom from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ClientOptionsCustom) Marshal() marshalDhcpv6ClientOptionsCustom {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ClientOptionsCustom{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ClientOptionsCustom) Unmarshal() unMarshalDhcpv6ClientOptionsCustom {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ClientOptionsCustom{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ClientOptionsCustom) ToProto() (*otg.Dhcpv6ClientOptionsCustom, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ClientOptionsCustom) FromProto(msg *otg.Dhcpv6ClientOptionsCustom) (Dhcpv6ClientOptionsCustom, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ClientOptionsCustom) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsCustom) FromPbText(value string) error {
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

func (m *marshaldhcpv6ClientOptionsCustom) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsCustom) FromYaml(value string) error {
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

func (m *marshaldhcpv6ClientOptionsCustom) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv6ClientOptionsCustom) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsCustom) FromJson(value string) error {
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

func (obj *dhcpv6ClientOptionsCustom) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsCustom) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsCustom) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ClientOptionsCustom) Clone() (Dhcpv6ClientOptionsCustom, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ClientOptionsCustom()
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

// Dhcpv6ClientOptionsCustom is the Custom option is used to provide a not so well known option in the message between a client and a server.
type Dhcpv6ClientOptionsCustom interface {
	Validation
	// msg marshals Dhcpv6ClientOptionsCustom to protobuf object *otg.Dhcpv6ClientOptionsCustom
	// and doesn't set defaults
	msg() *otg.Dhcpv6ClientOptionsCustom
	// setMsg unmarshals Dhcpv6ClientOptionsCustom from protobuf object *otg.Dhcpv6ClientOptionsCustom
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ClientOptionsCustom) Dhcpv6ClientOptionsCustom
	// provides marshal interface
	Marshal() marshalDhcpv6ClientOptionsCustom
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ClientOptionsCustom
	// validate validates Dhcpv6ClientOptionsCustom
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ClientOptionsCustom, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns uint32, set in Dhcpv6ClientOptionsCustom.
	Type() uint32
	// SetType assigns uint32 provided by user to Dhcpv6ClientOptionsCustom
	SetType(value uint32) Dhcpv6ClientOptionsCustom
}

// The type of the Custom option TLV.
// Type returns a uint32
func (obj *dhcpv6ClientOptionsCustom) Type() uint32 {

	return *obj.obj.Type

}

// The type of the Custom option TLV.
// SetType sets the uint32 value in the Dhcpv6ClientOptionsCustom object
func (obj *dhcpv6ClientOptionsCustom) SetType(value uint32) Dhcpv6ClientOptionsCustom {

	obj.obj.Type = &value
	return obj
}

func (obj *dhcpv6ClientOptionsCustom) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Type is required
	if obj.obj.Type == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Type is required field on interface Dhcpv6ClientOptionsCustom")
	}
	if obj.obj.Type != nil {

		if *obj.obj.Type > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Dhcpv6ClientOptionsCustom.Type <= 65535 but Got %d", *obj.obj.Type))
		}

	}

}

func (obj *dhcpv6ClientOptionsCustom) setDefault() {

}

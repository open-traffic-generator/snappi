package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ProtocolOptions *****
type protocolOptions struct {
	validation
	obj          *otg.ProtocolOptions
	marshaller   marshalProtocolOptions
	unMarshaller unMarshalProtocolOptions
}

func NewProtocolOptions() ProtocolOptions {
	obj := protocolOptions{obj: &otg.ProtocolOptions{}}
	obj.setDefault()
	return &obj
}

func (obj *protocolOptions) msg() *otg.ProtocolOptions {
	return obj.obj
}

func (obj *protocolOptions) setMsg(msg *otg.ProtocolOptions) ProtocolOptions {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalprotocolOptions struct {
	obj *protocolOptions
}

type marshalProtocolOptions interface {
	// ToProto marshals ProtocolOptions to protobuf object *otg.ProtocolOptions
	ToProto() (*otg.ProtocolOptions, error)
	// ToPbText marshals ProtocolOptions to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ProtocolOptions to YAML text
	ToYaml() (string, error)
	// ToJson marshals ProtocolOptions to JSON text
	ToJson() (string, error)
}

type unMarshalprotocolOptions struct {
	obj *protocolOptions
}

type unMarshalProtocolOptions interface {
	// FromProto unmarshals ProtocolOptions from protobuf object *otg.ProtocolOptions
	FromProto(msg *otg.ProtocolOptions) (ProtocolOptions, error)
	// FromPbText unmarshals ProtocolOptions from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ProtocolOptions from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ProtocolOptions from JSON text
	FromJson(value string) error
}

func (obj *protocolOptions) Marshal() marshalProtocolOptions {
	if obj.marshaller == nil {
		obj.marshaller = &marshalprotocolOptions{obj: obj}
	}
	return obj.marshaller
}

func (obj *protocolOptions) Unmarshal() unMarshalProtocolOptions {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalprotocolOptions{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalprotocolOptions) ToProto() (*otg.ProtocolOptions, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalprotocolOptions) FromProto(msg *otg.ProtocolOptions) (ProtocolOptions, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalprotocolOptions) ToPbText() (string, error) {
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

func (m *unMarshalprotocolOptions) FromPbText(value string) error {
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

func (m *marshalprotocolOptions) ToYaml() (string, error) {
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

func (m *unMarshalprotocolOptions) FromYaml(value string) error {
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

func (m *marshalprotocolOptions) ToJson() (string, error) {
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

func (m *unMarshalprotocolOptions) FromJson(value string) error {
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

func (obj *protocolOptions) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *protocolOptions) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *protocolOptions) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *protocolOptions) Clone() (ProtocolOptions, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewProtocolOptions()
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

// ProtocolOptions is common options that apply to all configured protocols and interfaces.
type ProtocolOptions interface {
	Validation
	// msg marshals ProtocolOptions to protobuf object *otg.ProtocolOptions
	// and doesn't set defaults
	msg() *otg.ProtocolOptions
	// setMsg unmarshals ProtocolOptions from protobuf object *otg.ProtocolOptions
	// and doesn't set defaults
	setMsg(*otg.ProtocolOptions) ProtocolOptions
	// provides marshal interface
	Marshal() marshalProtocolOptions
	// provides unmarshal interface
	Unmarshal() unMarshalProtocolOptions
	// validate validates ProtocolOptions
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ProtocolOptions, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// AutoStartAll returns bool, set in ProtocolOptions.
	AutoStartAll() bool
	// SetAutoStartAll assigns bool provided by user to ProtocolOptions
	SetAutoStartAll(value bool) ProtocolOptions
	// HasAutoStartAll checks if AutoStartAll has been set in ProtocolOptions
	HasAutoStartAll() bool
}

// When set to true, all underlying resources for configured protocols   and interfaces shall be created and corresponding protocol session negotiation shall be initiated. Otherwise, when set to false,  corresponding protocol session negotiation will need to be initiated  using a separate set_protocol_state API call.
// AutoStartAll returns a bool
func (obj *protocolOptions) AutoStartAll() bool {

	return *obj.obj.AutoStartAll

}

// When set to true, all underlying resources for configured protocols   and interfaces shall be created and corresponding protocol session negotiation shall be initiated. Otherwise, when set to false,  corresponding protocol session negotiation will need to be initiated  using a separate set_protocol_state API call.
// AutoStartAll returns a bool
func (obj *protocolOptions) HasAutoStartAll() bool {
	return obj.obj.AutoStartAll != nil
}

// When set to true, all underlying resources for configured protocols   and interfaces shall be created and corresponding protocol session negotiation shall be initiated. Otherwise, when set to false,  corresponding protocol session negotiation will need to be initiated  using a separate set_protocol_state API call.
// SetAutoStartAll sets the bool value in the ProtocolOptions object
func (obj *protocolOptions) SetAutoStartAll(value bool) ProtocolOptions {

	obj.obj.AutoStartAll = &value
	return obj
}

func (obj *protocolOptions) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *protocolOptions) setDefault() {
	if obj.obj.AutoStartAll == nil {
		obj.SetAutoStartAll(true)
	}

}

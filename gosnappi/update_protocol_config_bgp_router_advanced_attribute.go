package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateProtocolConfigBgpRouterAdvancedAttribute *****
type updateProtocolConfigBgpRouterAdvancedAttribute struct {
	validation
	obj          *otg.UpdateProtocolConfigBgpRouterAdvancedAttribute
	marshaller   marshalUpdateProtocolConfigBgpRouterAdvancedAttribute
	unMarshaller unMarshalUpdateProtocolConfigBgpRouterAdvancedAttribute
}

func NewUpdateProtocolConfigBgpRouterAdvancedAttribute() UpdateProtocolConfigBgpRouterAdvancedAttribute {
	obj := updateProtocolConfigBgpRouterAdvancedAttribute{obj: &otg.UpdateProtocolConfigBgpRouterAdvancedAttribute{}}
	obj.setDefault()
	return &obj
}

func (obj *updateProtocolConfigBgpRouterAdvancedAttribute) msg() *otg.UpdateProtocolConfigBgpRouterAdvancedAttribute {
	return obj.obj
}

func (obj *updateProtocolConfigBgpRouterAdvancedAttribute) setMsg(msg *otg.UpdateProtocolConfigBgpRouterAdvancedAttribute) UpdateProtocolConfigBgpRouterAdvancedAttribute {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateProtocolConfigBgpRouterAdvancedAttribute struct {
	obj *updateProtocolConfigBgpRouterAdvancedAttribute
}

type marshalUpdateProtocolConfigBgpRouterAdvancedAttribute interface {
	// ToProto marshals UpdateProtocolConfigBgpRouterAdvancedAttribute to protobuf object *otg.UpdateProtocolConfigBgpRouterAdvancedAttribute
	ToProto() (*otg.UpdateProtocolConfigBgpRouterAdvancedAttribute, error)
	// ToPbText marshals UpdateProtocolConfigBgpRouterAdvancedAttribute to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateProtocolConfigBgpRouterAdvancedAttribute to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateProtocolConfigBgpRouterAdvancedAttribute to JSON text
	ToJson() (string, error)
}

type unMarshalupdateProtocolConfigBgpRouterAdvancedAttribute struct {
	obj *updateProtocolConfigBgpRouterAdvancedAttribute
}

type unMarshalUpdateProtocolConfigBgpRouterAdvancedAttribute interface {
	// FromProto unmarshals UpdateProtocolConfigBgpRouterAdvancedAttribute from protobuf object *otg.UpdateProtocolConfigBgpRouterAdvancedAttribute
	FromProto(msg *otg.UpdateProtocolConfigBgpRouterAdvancedAttribute) (UpdateProtocolConfigBgpRouterAdvancedAttribute, error)
	// FromPbText unmarshals UpdateProtocolConfigBgpRouterAdvancedAttribute from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateProtocolConfigBgpRouterAdvancedAttribute from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateProtocolConfigBgpRouterAdvancedAttribute from JSON text
	FromJson(value string) error
}

func (obj *updateProtocolConfigBgpRouterAdvancedAttribute) Marshal() marshalUpdateProtocolConfigBgpRouterAdvancedAttribute {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateProtocolConfigBgpRouterAdvancedAttribute{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateProtocolConfigBgpRouterAdvancedAttribute) Unmarshal() unMarshalUpdateProtocolConfigBgpRouterAdvancedAttribute {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateProtocolConfigBgpRouterAdvancedAttribute{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateProtocolConfigBgpRouterAdvancedAttribute) ToProto() (*otg.UpdateProtocolConfigBgpRouterAdvancedAttribute, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateProtocolConfigBgpRouterAdvancedAttribute) FromProto(msg *otg.UpdateProtocolConfigBgpRouterAdvancedAttribute) (UpdateProtocolConfigBgpRouterAdvancedAttribute, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateProtocolConfigBgpRouterAdvancedAttribute) ToPbText() (string, error) {
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

func (m *unMarshalupdateProtocolConfigBgpRouterAdvancedAttribute) FromPbText(value string) error {
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

func (m *marshalupdateProtocolConfigBgpRouterAdvancedAttribute) ToYaml() (string, error) {
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

func (m *unMarshalupdateProtocolConfigBgpRouterAdvancedAttribute) FromYaml(value string) error {
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

func (m *marshalupdateProtocolConfigBgpRouterAdvancedAttribute) ToJson() (string, error) {
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

func (m *unMarshalupdateProtocolConfigBgpRouterAdvancedAttribute) FromJson(value string) error {
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

func (obj *updateProtocolConfigBgpRouterAdvancedAttribute) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateProtocolConfigBgpRouterAdvancedAttribute) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateProtocolConfigBgpRouterAdvancedAttribute) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateProtocolConfigBgpRouterAdvancedAttribute) Clone() (UpdateProtocolConfigBgpRouterAdvancedAttribute, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateProtocolConfigBgpRouterAdvancedAttribute()
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

// UpdateProtocolConfigBgpRouterAdvancedAttribute is the advanced properties of a BGP router to be updated.
type UpdateProtocolConfigBgpRouterAdvancedAttribute interface {
	Validation
	// msg marshals UpdateProtocolConfigBgpRouterAdvancedAttribute to protobuf object *otg.UpdateProtocolConfigBgpRouterAdvancedAttribute
	// and doesn't set defaults
	msg() *otg.UpdateProtocolConfigBgpRouterAdvancedAttribute
	// setMsg unmarshals UpdateProtocolConfigBgpRouterAdvancedAttribute from protobuf object *otg.UpdateProtocolConfigBgpRouterAdvancedAttribute
	// and doesn't set defaults
	setMsg(*otg.UpdateProtocolConfigBgpRouterAdvancedAttribute) UpdateProtocolConfigBgpRouterAdvancedAttribute
	// provides marshal interface
	Marshal() marshalUpdateProtocolConfigBgpRouterAdvancedAttribute
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateProtocolConfigBgpRouterAdvancedAttribute
	// validate validates UpdateProtocolConfigBgpRouterAdvancedAttribute
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateProtocolConfigBgpRouterAdvancedAttribute, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Md5Key returns string, set in UpdateProtocolConfigBgpRouterAdvancedAttribute.
	Md5Key() string
	// SetMd5Key assigns string provided by user to UpdateProtocolConfigBgpRouterAdvancedAttribute
	SetMd5Key(value string) UpdateProtocolConfigBgpRouterAdvancedAttribute
	// HasMd5Key checks if Md5Key has been set in UpdateProtocolConfigBgpRouterAdvancedAttribute
	HasMd5Key() bool
	// PassiveMode returns bool, set in UpdateProtocolConfigBgpRouterAdvancedAttribute.
	PassiveMode() bool
	// SetPassiveMode assigns bool provided by user to UpdateProtocolConfigBgpRouterAdvancedAttribute
	SetPassiveMode(value bool) UpdateProtocolConfigBgpRouterAdvancedAttribute
	// HasPassiveMode checks if PassiveMode has been set in UpdateProtocolConfigBgpRouterAdvancedAttribute
	HasPassiveMode() bool
}

// The MD5 authentication key to be updated on the BGP router. If not configured, MD5 authentication will not be enabled.
// Md5Key returns a string
func (obj *updateProtocolConfigBgpRouterAdvancedAttribute) Md5Key() string {

	return *obj.obj.Md5Key

}

// The MD5 authentication key to be updated on the BGP router. If not configured, MD5 authentication will not be enabled.
// Md5Key returns a string
func (obj *updateProtocolConfigBgpRouterAdvancedAttribute) HasMd5Key() bool {
	return obj.obj.Md5Key != nil
}

// The MD5 authentication key to be updated on the BGP router. If not configured, MD5 authentication will not be enabled.
// SetMd5Key sets the string value in the UpdateProtocolConfigBgpRouterAdvancedAttribute object
func (obj *updateProtocolConfigBgpRouterAdvancedAttribute) SetMd5Key(value string) UpdateProtocolConfigBgpRouterAdvancedAttribute {

	obj.obj.Md5Key = &value
	return obj
}

// If set to true, the local BGP peer will wait for the remote peer to initiate the BGP session by establishing the TCP connection, rather than initiating sessions from the local peer.
// PassiveMode returns a bool
func (obj *updateProtocolConfigBgpRouterAdvancedAttribute) PassiveMode() bool {

	return *obj.obj.PassiveMode

}

// If set to true, the local BGP peer will wait for the remote peer to initiate the BGP session by establishing the TCP connection, rather than initiating sessions from the local peer.
// PassiveMode returns a bool
func (obj *updateProtocolConfigBgpRouterAdvancedAttribute) HasPassiveMode() bool {
	return obj.obj.PassiveMode != nil
}

// If set to true, the local BGP peer will wait for the remote peer to initiate the BGP session by establishing the TCP connection, rather than initiating sessions from the local peer.
// SetPassiveMode sets the bool value in the UpdateProtocolConfigBgpRouterAdvancedAttribute object
func (obj *updateProtocolConfigBgpRouterAdvancedAttribute) SetPassiveMode(value bool) UpdateProtocolConfigBgpRouterAdvancedAttribute {

	obj.obj.PassiveMode = &value
	return obj
}

func (obj *updateProtocolConfigBgpRouterAdvancedAttribute) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *updateProtocolConfigBgpRouterAdvancedAttribute) setDefault() {
	if obj.obj.PassiveMode == nil {
		obj.SetPassiveMode(false)
	}

}

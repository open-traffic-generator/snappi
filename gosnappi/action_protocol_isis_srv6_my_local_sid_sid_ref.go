package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolIsisSrv6MyLocalSidSidRef *****
type actionProtocolIsisSrv6MyLocalSidSidRef struct {
	validation
	obj          *otg.ActionProtocolIsisSrv6MyLocalSidSidRef
	marshaller   marshalActionProtocolIsisSrv6MyLocalSidSidRef
	unMarshaller unMarshalActionProtocolIsisSrv6MyLocalSidSidRef
}

func NewActionProtocolIsisSrv6MyLocalSidSidRef() ActionProtocolIsisSrv6MyLocalSidSidRef {
	obj := actionProtocolIsisSrv6MyLocalSidSidRef{obj: &otg.ActionProtocolIsisSrv6MyLocalSidSidRef{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolIsisSrv6MyLocalSidSidRef) msg() *otg.ActionProtocolIsisSrv6MyLocalSidSidRef {
	return obj.obj
}

func (obj *actionProtocolIsisSrv6MyLocalSidSidRef) setMsg(msg *otg.ActionProtocolIsisSrv6MyLocalSidSidRef) ActionProtocolIsisSrv6MyLocalSidSidRef {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolIsisSrv6MyLocalSidSidRef struct {
	obj *actionProtocolIsisSrv6MyLocalSidSidRef
}

type marshalActionProtocolIsisSrv6MyLocalSidSidRef interface {
	// ToProto marshals ActionProtocolIsisSrv6MyLocalSidSidRef to protobuf object *otg.ActionProtocolIsisSrv6MyLocalSidSidRef
	ToProto() (*otg.ActionProtocolIsisSrv6MyLocalSidSidRef, error)
	// ToPbText marshals ActionProtocolIsisSrv6MyLocalSidSidRef to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolIsisSrv6MyLocalSidSidRef to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolIsisSrv6MyLocalSidSidRef to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolIsisSrv6MyLocalSidSidRef struct {
	obj *actionProtocolIsisSrv6MyLocalSidSidRef
}

type unMarshalActionProtocolIsisSrv6MyLocalSidSidRef interface {
	// FromProto unmarshals ActionProtocolIsisSrv6MyLocalSidSidRef from protobuf object *otg.ActionProtocolIsisSrv6MyLocalSidSidRef
	FromProto(msg *otg.ActionProtocolIsisSrv6MyLocalSidSidRef) (ActionProtocolIsisSrv6MyLocalSidSidRef, error)
	// FromPbText unmarshals ActionProtocolIsisSrv6MyLocalSidSidRef from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolIsisSrv6MyLocalSidSidRef from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolIsisSrv6MyLocalSidSidRef from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolIsisSrv6MyLocalSidSidRef) Marshal() marshalActionProtocolIsisSrv6MyLocalSidSidRef {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolIsisSrv6MyLocalSidSidRef{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolIsisSrv6MyLocalSidSidRef) Unmarshal() unMarshalActionProtocolIsisSrv6MyLocalSidSidRef {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolIsisSrv6MyLocalSidSidRef{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolIsisSrv6MyLocalSidSidRef) ToProto() (*otg.ActionProtocolIsisSrv6MyLocalSidSidRef, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolIsisSrv6MyLocalSidSidRef) FromProto(msg *otg.ActionProtocolIsisSrv6MyLocalSidSidRef) (ActionProtocolIsisSrv6MyLocalSidSidRef, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolIsisSrv6MyLocalSidSidRef) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolIsisSrv6MyLocalSidSidRef) FromPbText(value string) error {
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

func (m *marshalactionProtocolIsisSrv6MyLocalSidSidRef) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolIsisSrv6MyLocalSidSidRef) FromYaml(value string) error {
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

func (m *marshalactionProtocolIsisSrv6MyLocalSidSidRef) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolIsisSrv6MyLocalSidSidRef) FromJson(value string) error {
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

func (obj *actionProtocolIsisSrv6MyLocalSidSidRef) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolIsisSrv6MyLocalSidSidRef) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolIsisSrv6MyLocalSidSidRef) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolIsisSrv6MyLocalSidSidRef) Clone() (ActionProtocolIsisSrv6MyLocalSidSidRef, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolIsisSrv6MyLocalSidSidRef()
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

// ActionProtocolIsisSrv6MyLocalSidSidRef is identifies a My Local SID entry by its prefix and length, used in delete operations.
type ActionProtocolIsisSrv6MyLocalSidSidRef interface {
	Validation
	// msg marshals ActionProtocolIsisSrv6MyLocalSidSidRef to protobuf object *otg.ActionProtocolIsisSrv6MyLocalSidSidRef
	// and doesn't set defaults
	msg() *otg.ActionProtocolIsisSrv6MyLocalSidSidRef
	// setMsg unmarshals ActionProtocolIsisSrv6MyLocalSidSidRef from protobuf object *otg.ActionProtocolIsisSrv6MyLocalSidSidRef
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolIsisSrv6MyLocalSidSidRef) ActionProtocolIsisSrv6MyLocalSidSidRef
	// provides marshal interface
	Marshal() marshalActionProtocolIsisSrv6MyLocalSidSidRef
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolIsisSrv6MyLocalSidSidRef
	// validate validates ActionProtocolIsisSrv6MyLocalSidSidRef
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolIsisSrv6MyLocalSidSidRef, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SidPrefix returns string, set in ActionProtocolIsisSrv6MyLocalSidSidRef.
	SidPrefix() string
	// SetSidPrefix assigns string provided by user to ActionProtocolIsisSrv6MyLocalSidSidRef
	SetSidPrefix(value string) ActionProtocolIsisSrv6MyLocalSidSidRef
	// PrefixLength returns uint32, set in ActionProtocolIsisSrv6MyLocalSidSidRef.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to ActionProtocolIsisSrv6MyLocalSidSidRef
	SetPrefixLength(value uint32) ActionProtocolIsisSrv6MyLocalSidSidRef
}

// IPv6 address of the My Local SID entry to delete.
// SidPrefix returns a string
func (obj *actionProtocolIsisSrv6MyLocalSidSidRef) SidPrefix() string {

	return *obj.obj.SidPrefix

}

// IPv6 address of the My Local SID entry to delete.
// SetSidPrefix sets the string value in the ActionProtocolIsisSrv6MyLocalSidSidRef object
func (obj *actionProtocolIsisSrv6MyLocalSidSidRef) SetSidPrefix(value string) ActionProtocolIsisSrv6MyLocalSidSidRef {

	obj.obj.SidPrefix = &value
	return obj
}

// Prefix length of the My Local SID to delete, in bits.
// PrefixLength returns a uint32
func (obj *actionProtocolIsisSrv6MyLocalSidSidRef) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// Prefix length of the My Local SID to delete, in bits.
// SetPrefixLength sets the uint32 value in the ActionProtocolIsisSrv6MyLocalSidSidRef object
func (obj *actionProtocolIsisSrv6MyLocalSidSidRef) SetPrefixLength(value uint32) ActionProtocolIsisSrv6MyLocalSidSidRef {

	obj.obj.PrefixLength = &value
	return obj
}

func (obj *actionProtocolIsisSrv6MyLocalSidSidRef) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// SidPrefix is required
	if obj.obj.SidPrefix == nil {
		vObj.validationErrors = append(vObj.validationErrors, "SidPrefix is required field on interface ActionProtocolIsisSrv6MyLocalSidSidRef")
	}
	if obj.obj.SidPrefix != nil {

		err := obj.validateIpv6(obj.SidPrefix())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on ActionProtocolIsisSrv6MyLocalSidSidRef.SidPrefix"))
		}

	}

	// PrefixLength is required
	if obj.obj.PrefixLength == nil {
		vObj.validationErrors = append(vObj.validationErrors, "PrefixLength is required field on interface ActionProtocolIsisSrv6MyLocalSidSidRef")
	}
	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength < 1 || *obj.obj.PrefixLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= ActionProtocolIsisSrv6MyLocalSidSidRef.PrefixLength <= 128 but Got %d", *obj.obj.PrefixLength))
		}

	}

}

func (obj *actionProtocolIsisSrv6MyLocalSidSidRef) setDefault() {

}

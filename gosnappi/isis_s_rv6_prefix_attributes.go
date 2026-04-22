package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRv6PrefixAttributes *****
type isisSRv6PrefixAttributes struct {
	validation
	obj          *otg.IsisSRv6PrefixAttributes
	marshaller   marshalIsisSRv6PrefixAttributes
	unMarshaller unMarshalIsisSRv6PrefixAttributes
}

func NewIsisSRv6PrefixAttributes() IsisSRv6PrefixAttributes {
	obj := isisSRv6PrefixAttributes{obj: &otg.IsisSRv6PrefixAttributes{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRv6PrefixAttributes) msg() *otg.IsisSRv6PrefixAttributes {
	return obj.obj
}

func (obj *isisSRv6PrefixAttributes) setMsg(msg *otg.IsisSRv6PrefixAttributes) IsisSRv6PrefixAttributes {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRv6PrefixAttributes struct {
	obj *isisSRv6PrefixAttributes
}

type marshalIsisSRv6PrefixAttributes interface {
	// ToProto marshals IsisSRv6PrefixAttributes to protobuf object *otg.IsisSRv6PrefixAttributes
	ToProto() (*otg.IsisSRv6PrefixAttributes, error)
	// ToPbText marshals IsisSRv6PrefixAttributes to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRv6PrefixAttributes to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRv6PrefixAttributes to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRv6PrefixAttributes struct {
	obj *isisSRv6PrefixAttributes
}

type unMarshalIsisSRv6PrefixAttributes interface {
	// FromProto unmarshals IsisSRv6PrefixAttributes from protobuf object *otg.IsisSRv6PrefixAttributes
	FromProto(msg *otg.IsisSRv6PrefixAttributes) (IsisSRv6PrefixAttributes, error)
	// FromPbText unmarshals IsisSRv6PrefixAttributes from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRv6PrefixAttributes from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRv6PrefixAttributes from JSON text
	FromJson(value string) error
}

func (obj *isisSRv6PrefixAttributes) Marshal() marshalIsisSRv6PrefixAttributes {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRv6PrefixAttributes{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRv6PrefixAttributes) Unmarshal() unMarshalIsisSRv6PrefixAttributes {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRv6PrefixAttributes{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRv6PrefixAttributes) ToProto() (*otg.IsisSRv6PrefixAttributes, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRv6PrefixAttributes) FromProto(msg *otg.IsisSRv6PrefixAttributes) (IsisSRv6PrefixAttributes, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRv6PrefixAttributes) ToPbText() (string, error) {
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

func (m *unMarshalisisSRv6PrefixAttributes) FromPbText(value string) error {
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

func (m *marshalisisSRv6PrefixAttributes) ToYaml() (string, error) {
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

func (m *unMarshalisisSRv6PrefixAttributes) FromYaml(value string) error {
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

func (m *marshalisisSRv6PrefixAttributes) ToJson() (string, error) {
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

func (m *unMarshalisisSRv6PrefixAttributes) FromJson(value string) error {
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

func (obj *isisSRv6PrefixAttributes) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRv6PrefixAttributes) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRv6PrefixAttributes) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRv6PrefixAttributes) Clone() (IsisSRv6PrefixAttributes, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRv6PrefixAttributes()
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

// IsisSRv6PrefixAttributes is prefix Attribute Flags Sub-TLV (sub-TLV type 4, RFC 7794) for an SRv6 locator prefix advertisement. Carried within the IS-IS IPv6 Reachability TLV (TLV 236/237) when the locator is also advertised as a prefix. Presence of this object causes the sub-TLV to be included; absence suppresses it. Reference: RFC 7794.
type IsisSRv6PrefixAttributes interface {
	Validation
	// msg marshals IsisSRv6PrefixAttributes to protobuf object *otg.IsisSRv6PrefixAttributes
	// and doesn't set defaults
	msg() *otg.IsisSRv6PrefixAttributes
	// setMsg unmarshals IsisSRv6PrefixAttributes from protobuf object *otg.IsisSRv6PrefixAttributes
	// and doesn't set defaults
	setMsg(*otg.IsisSRv6PrefixAttributes) IsisSRv6PrefixAttributes
	// provides marshal interface
	Marshal() marshalIsisSRv6PrefixAttributes
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRv6PrefixAttributes
	// validate validates IsisSRv6PrefixAttributes
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRv6PrefixAttributes, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// XFlag returns bool, set in IsisSRv6PrefixAttributes.
	XFlag() bool
	// SetXFlag assigns bool provided by user to IsisSRv6PrefixAttributes
	SetXFlag(value bool) IsisSRv6PrefixAttributes
	// HasXFlag checks if XFlag has been set in IsisSRv6PrefixAttributes
	HasXFlag() bool
	// RFlag returns bool, set in IsisSRv6PrefixAttributes.
	RFlag() bool
	// SetRFlag assigns bool provided by user to IsisSRv6PrefixAttributes
	SetRFlag(value bool) IsisSRv6PrefixAttributes
	// HasRFlag checks if RFlag has been set in IsisSRv6PrefixAttributes
	HasRFlag() bool
	// NFlag returns bool, set in IsisSRv6PrefixAttributes.
	NFlag() bool
	// SetNFlag assigns bool provided by user to IsisSRv6PrefixAttributes
	SetNFlag(value bool) IsisSRv6PrefixAttributes
	// HasNFlag checks if NFlag has been set in IsisSRv6PrefixAttributes
	HasNFlag() bool
	// AFlag returns bool, set in IsisSRv6PrefixAttributes.
	AFlag() bool
	// SetAFlag assigns bool provided by user to IsisSRv6PrefixAttributes
	SetAFlag(value bool) IsisSRv6PrefixAttributes
	// HasAFlag checks if AFlag has been set in IsisSRv6PrefixAttributes
	HasAFlag() bool
}

// External prefix flag (bit 0, RFC 7794). When set, indicates this locator prefix was redistributed from another protocol.
// XFlag returns a bool
func (obj *isisSRv6PrefixAttributes) XFlag() bool {

	return *obj.obj.XFlag

}

// External prefix flag (bit 0, RFC 7794). When set, indicates this locator prefix was redistributed from another protocol.
// XFlag returns a bool
func (obj *isisSRv6PrefixAttributes) HasXFlag() bool {
	return obj.obj.XFlag != nil
}

// External prefix flag (bit 0, RFC 7794). When set, indicates this locator prefix was redistributed from another protocol.
// SetXFlag sets the bool value in the IsisSRv6PrefixAttributes object
func (obj *isisSRv6PrefixAttributes) SetXFlag(value bool) IsisSRv6PrefixAttributes {

	obj.obj.XFlag = &value
	return obj
}

// Re-advertisement flag (bit 1, RFC 7794). When set, indicates this locator prefix has been leaked between IS-IS levels.
// RFlag returns a bool
func (obj *isisSRv6PrefixAttributes) RFlag() bool {

	return *obj.obj.RFlag

}

// Re-advertisement flag (bit 1, RFC 7794). When set, indicates this locator prefix has been leaked between IS-IS levels.
// RFlag returns a bool
func (obj *isisSRv6PrefixAttributes) HasRFlag() bool {
	return obj.obj.RFlag != nil
}

// Re-advertisement flag (bit 1, RFC 7794). When set, indicates this locator prefix has been leaked between IS-IS levels.
// SetRFlag sets the bool value in the IsisSRv6PrefixAttributes object
func (obj *isisSRv6PrefixAttributes) SetRFlag(value bool) IsisSRv6PrefixAttributes {

	obj.obj.RFlag = &value
	return obj
}

// Node flag (bit 2, RFC 7794). When set, indicates this prefix identifies the advertising router (e.g., a loopback or router-ID prefix).
// NFlag returns a bool
func (obj *isisSRv6PrefixAttributes) NFlag() bool {

	return *obj.obj.NFlag

}

// Node flag (bit 2, RFC 7794). When set, indicates this prefix identifies the advertising router (e.g., a loopback or router-ID prefix).
// NFlag returns a bool
func (obj *isisSRv6PrefixAttributes) HasNFlag() bool {
	return obj.obj.NFlag != nil
}

// Node flag (bit 2, RFC 7794). When set, indicates this prefix identifies the advertising router (e.g., a loopback or router-ID prefix).
// SetNFlag sets the bool value in the IsisSRv6PrefixAttributes object
func (obj *isisSRv6PrefixAttributes) SetNFlag(value bool) IsisSRv6PrefixAttributes {

	obj.obj.NFlag = &value
	return obj
}

// Anycast flag (bit 5, RFC 7794). When set, indicates this prefix is an anycast prefix shared across multiple routers.
// AFlag returns a bool
func (obj *isisSRv6PrefixAttributes) AFlag() bool {

	return *obj.obj.AFlag

}

// Anycast flag (bit 5, RFC 7794). When set, indicates this prefix is an anycast prefix shared across multiple routers.
// AFlag returns a bool
func (obj *isisSRv6PrefixAttributes) HasAFlag() bool {
	return obj.obj.AFlag != nil
}

// Anycast flag (bit 5, RFC 7794). When set, indicates this prefix is an anycast prefix shared across multiple routers.
// SetAFlag sets the bool value in the IsisSRv6PrefixAttributes object
func (obj *isisSRv6PrefixAttributes) SetAFlag(value bool) IsisSRv6PrefixAttributes {

	obj.obj.AFlag = &value
	return obj
}

func (obj *isisSRv6PrefixAttributes) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisSRv6PrefixAttributes) setDefault() {
	if obj.obj.XFlag == nil {
		obj.SetXFlag(false)
	}
	if obj.obj.RFlag == nil {
		obj.SetRFlag(false)
	}
	if obj.obj.NFlag == nil {
		obj.SetNFlag(false)
	}
	if obj.obj.AFlag == nil {
		obj.SetAFlag(false)
	}

}

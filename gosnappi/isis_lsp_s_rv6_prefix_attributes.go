package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspSRv6PrefixAttributes *****
type isisLspSRv6PrefixAttributes struct {
	validation
	obj          *otg.IsisLspSRv6PrefixAttributes
	marshaller   marshalIsisLspSRv6PrefixAttributes
	unMarshaller unMarshalIsisLspSRv6PrefixAttributes
}

func NewIsisLspSRv6PrefixAttributes() IsisLspSRv6PrefixAttributes {
	obj := isisLspSRv6PrefixAttributes{obj: &otg.IsisLspSRv6PrefixAttributes{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspSRv6PrefixAttributes) msg() *otg.IsisLspSRv6PrefixAttributes {
	return obj.obj
}

func (obj *isisLspSRv6PrefixAttributes) setMsg(msg *otg.IsisLspSRv6PrefixAttributes) IsisLspSRv6PrefixAttributes {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspSRv6PrefixAttributes struct {
	obj *isisLspSRv6PrefixAttributes
}

type marshalIsisLspSRv6PrefixAttributes interface {
	// ToProto marshals IsisLspSRv6PrefixAttributes to protobuf object *otg.IsisLspSRv6PrefixAttributes
	ToProto() (*otg.IsisLspSRv6PrefixAttributes, error)
	// ToPbText marshals IsisLspSRv6PrefixAttributes to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspSRv6PrefixAttributes to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspSRv6PrefixAttributes to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspSRv6PrefixAttributes struct {
	obj *isisLspSRv6PrefixAttributes
}

type unMarshalIsisLspSRv6PrefixAttributes interface {
	// FromProto unmarshals IsisLspSRv6PrefixAttributes from protobuf object *otg.IsisLspSRv6PrefixAttributes
	FromProto(msg *otg.IsisLspSRv6PrefixAttributes) (IsisLspSRv6PrefixAttributes, error)
	// FromPbText unmarshals IsisLspSRv6PrefixAttributes from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspSRv6PrefixAttributes from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspSRv6PrefixAttributes from JSON text
	FromJson(value string) error
}

func (obj *isisLspSRv6PrefixAttributes) Marshal() marshalIsisLspSRv6PrefixAttributes {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspSRv6PrefixAttributes{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspSRv6PrefixAttributes) Unmarshal() unMarshalIsisLspSRv6PrefixAttributes {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspSRv6PrefixAttributes{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspSRv6PrefixAttributes) ToProto() (*otg.IsisLspSRv6PrefixAttributes, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspSRv6PrefixAttributes) FromProto(msg *otg.IsisLspSRv6PrefixAttributes) (IsisLspSRv6PrefixAttributes, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspSRv6PrefixAttributes) ToPbText() (string, error) {
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

func (m *unMarshalisisLspSRv6PrefixAttributes) FromPbText(value string) error {
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

func (m *marshalisisLspSRv6PrefixAttributes) ToYaml() (string, error) {
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

func (m *unMarshalisisLspSRv6PrefixAttributes) FromYaml(value string) error {
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

func (m *marshalisisLspSRv6PrefixAttributes) ToJson() (string, error) {
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

func (m *unMarshalisisLspSRv6PrefixAttributes) FromJson(value string) error {
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

func (obj *isisLspSRv6PrefixAttributes) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspSRv6PrefixAttributes) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspSRv6PrefixAttributes) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspSRv6PrefixAttributes) Clone() (IsisLspSRv6PrefixAttributes, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspSRv6PrefixAttributes()
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

// IsisLspSRv6PrefixAttributes is prefix Attribute Flags Sub-TLV (sub-TLV type 4, RFC 7794) learned from the IPv6 Reachability advertisement of an SRv6 locator prefix. Reference: RFC 7794.
type IsisLspSRv6PrefixAttributes interface {
	Validation
	// msg marshals IsisLspSRv6PrefixAttributes to protobuf object *otg.IsisLspSRv6PrefixAttributes
	// and doesn't set defaults
	msg() *otg.IsisLspSRv6PrefixAttributes
	// setMsg unmarshals IsisLspSRv6PrefixAttributes from protobuf object *otg.IsisLspSRv6PrefixAttributes
	// and doesn't set defaults
	setMsg(*otg.IsisLspSRv6PrefixAttributes) IsisLspSRv6PrefixAttributes
	// provides marshal interface
	Marshal() marshalIsisLspSRv6PrefixAttributes
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspSRv6PrefixAttributes
	// validate validates IsisLspSRv6PrefixAttributes
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspSRv6PrefixAttributes, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// XFlag returns bool, set in IsisLspSRv6PrefixAttributes.
	XFlag() bool
	// SetXFlag assigns bool provided by user to IsisLspSRv6PrefixAttributes
	SetXFlag(value bool) IsisLspSRv6PrefixAttributes
	// HasXFlag checks if XFlag has been set in IsisLspSRv6PrefixAttributes
	HasXFlag() bool
	// RFlag returns bool, set in IsisLspSRv6PrefixAttributes.
	RFlag() bool
	// SetRFlag assigns bool provided by user to IsisLspSRv6PrefixAttributes
	SetRFlag(value bool) IsisLspSRv6PrefixAttributes
	// HasRFlag checks if RFlag has been set in IsisLspSRv6PrefixAttributes
	HasRFlag() bool
	// NFlag returns bool, set in IsisLspSRv6PrefixAttributes.
	NFlag() bool
	// SetNFlag assigns bool provided by user to IsisLspSRv6PrefixAttributes
	SetNFlag(value bool) IsisLspSRv6PrefixAttributes
	// HasNFlag checks if NFlag has been set in IsisLspSRv6PrefixAttributes
	HasNFlag() bool
	// AFlag returns bool, set in IsisLspSRv6PrefixAttributes.
	AFlag() bool
	// SetAFlag assigns bool provided by user to IsisLspSRv6PrefixAttributes
	SetAFlag(value bool) IsisLspSRv6PrefixAttributes
	// HasAFlag checks if AFlag has been set in IsisLspSRv6PrefixAttributes
	HasAFlag() bool
}

// External prefix flag (bit 0, RFC 7794). Set if the locator prefix was redistributed from another protocol.
// XFlag returns a bool
func (obj *isisLspSRv6PrefixAttributes) XFlag() bool {

	return *obj.obj.XFlag

}

// External prefix flag (bit 0, RFC 7794). Set if the locator prefix was redistributed from another protocol.
// XFlag returns a bool
func (obj *isisLspSRv6PrefixAttributes) HasXFlag() bool {
	return obj.obj.XFlag != nil
}

// External prefix flag (bit 0, RFC 7794). Set if the locator prefix was redistributed from another protocol.
// SetXFlag sets the bool value in the IsisLspSRv6PrefixAttributes object
func (obj *isisLspSRv6PrefixAttributes) SetXFlag(value bool) IsisLspSRv6PrefixAttributes {

	obj.obj.XFlag = &value
	return obj
}

// Re-advertisement flag (bit 1, RFC 7794). Set if the locator prefix was leaked between IS-IS levels.
// RFlag returns a bool
func (obj *isisLspSRv6PrefixAttributes) RFlag() bool {

	return *obj.obj.RFlag

}

// Re-advertisement flag (bit 1, RFC 7794). Set if the locator prefix was leaked between IS-IS levels.
// RFlag returns a bool
func (obj *isisLspSRv6PrefixAttributes) HasRFlag() bool {
	return obj.obj.RFlag != nil
}

// Re-advertisement flag (bit 1, RFC 7794). Set if the locator prefix was leaked between IS-IS levels.
// SetRFlag sets the bool value in the IsisLspSRv6PrefixAttributes object
func (obj *isisLspSRv6PrefixAttributes) SetRFlag(value bool) IsisLspSRv6PrefixAttributes {

	obj.obj.RFlag = &value
	return obj
}

// Node flag (bit 2, RFC 7794). Set if this prefix identifies the advertising router.
// NFlag returns a bool
func (obj *isisLspSRv6PrefixAttributes) NFlag() bool {

	return *obj.obj.NFlag

}

// Node flag (bit 2, RFC 7794). Set if this prefix identifies the advertising router.
// NFlag returns a bool
func (obj *isisLspSRv6PrefixAttributes) HasNFlag() bool {
	return obj.obj.NFlag != nil
}

// Node flag (bit 2, RFC 7794). Set if this prefix identifies the advertising router.
// SetNFlag sets the bool value in the IsisLspSRv6PrefixAttributes object
func (obj *isisLspSRv6PrefixAttributes) SetNFlag(value bool) IsisLspSRv6PrefixAttributes {

	obj.obj.NFlag = &value
	return obj
}

// Anycast flag (bit 5, RFC 7794). Set if this locator is an anycast prefix shared across multiple routers.
// AFlag returns a bool
func (obj *isisLspSRv6PrefixAttributes) AFlag() bool {

	return *obj.obj.AFlag

}

// Anycast flag (bit 5, RFC 7794). Set if this locator is an anycast prefix shared across multiple routers.
// AFlag returns a bool
func (obj *isisLspSRv6PrefixAttributes) HasAFlag() bool {
	return obj.obj.AFlag != nil
}

// Anycast flag (bit 5, RFC 7794). Set if this locator is an anycast prefix shared across multiple routers.
// SetAFlag sets the bool value in the IsisLspSRv6PrefixAttributes object
func (obj *isisLspSRv6PrefixAttributes) SetAFlag(value bool) IsisLspSRv6PrefixAttributes {

	obj.obj.AFlag = &value
	return obj
}

func (obj *isisLspSRv6PrefixAttributes) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisLspSRv6PrefixAttributes) setDefault() {

}

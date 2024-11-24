package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspPrefixAttributes *****
type isisLspPrefixAttributes struct {
	validation
	obj          *otg.IsisLspPrefixAttributes
	marshaller   marshalIsisLspPrefixAttributes
	unMarshaller unMarshalIsisLspPrefixAttributes
}

func NewIsisLspPrefixAttributes() IsisLspPrefixAttributes {
	obj := isisLspPrefixAttributes{obj: &otg.IsisLspPrefixAttributes{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspPrefixAttributes) msg() *otg.IsisLspPrefixAttributes {
	return obj.obj
}

func (obj *isisLspPrefixAttributes) setMsg(msg *otg.IsisLspPrefixAttributes) IsisLspPrefixAttributes {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspPrefixAttributes struct {
	obj *isisLspPrefixAttributes
}

type marshalIsisLspPrefixAttributes interface {
	// ToProto marshals IsisLspPrefixAttributes to protobuf object *otg.IsisLspPrefixAttributes
	ToProto() (*otg.IsisLspPrefixAttributes, error)
	// ToPbText marshals IsisLspPrefixAttributes to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspPrefixAttributes to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspPrefixAttributes to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisLspPrefixAttributes to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisLspPrefixAttributes struct {
	obj *isisLspPrefixAttributes
}

type unMarshalIsisLspPrefixAttributes interface {
	// FromProto unmarshals IsisLspPrefixAttributes from protobuf object *otg.IsisLspPrefixAttributes
	FromProto(msg *otg.IsisLspPrefixAttributes) (IsisLspPrefixAttributes, error)
	// FromPbText unmarshals IsisLspPrefixAttributes from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspPrefixAttributes from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspPrefixAttributes from JSON text
	FromJson(value string) error
}

func (obj *isisLspPrefixAttributes) Marshal() marshalIsisLspPrefixAttributes {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspPrefixAttributes{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspPrefixAttributes) Unmarshal() unMarshalIsisLspPrefixAttributes {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspPrefixAttributes{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspPrefixAttributes) ToProto() (*otg.IsisLspPrefixAttributes, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspPrefixAttributes) FromProto(msg *otg.IsisLspPrefixAttributes) (IsisLspPrefixAttributes, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspPrefixAttributes) ToPbText() (string, error) {
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

func (m *unMarshalisisLspPrefixAttributes) FromPbText(value string) error {
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

func (m *marshalisisLspPrefixAttributes) ToYaml() (string, error) {
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

func (m *unMarshalisisLspPrefixAttributes) FromYaml(value string) error {
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

func (m *marshalisisLspPrefixAttributes) ToJsonRaw() (string, error) {
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

func (m *marshalisisLspPrefixAttributes) ToJson() (string, error) {
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

func (m *unMarshalisisLspPrefixAttributes) FromJson(value string) error {
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

func (obj *isisLspPrefixAttributes) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspPrefixAttributes) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspPrefixAttributes) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspPrefixAttributes) Clone() (IsisLspPrefixAttributes, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspPrefixAttributes()
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

// IsisLspPrefixAttributes is this contains the properties of ISIS Prefix attributes for  the extended IPv4 and IPv6 reachability. https://www.rfc-editor.org/rfc/rfc7794.html
type IsisLspPrefixAttributes interface {
	Validation
	// msg marshals IsisLspPrefixAttributes to protobuf object *otg.IsisLspPrefixAttributes
	// and doesn't set defaults
	msg() *otg.IsisLspPrefixAttributes
	// setMsg unmarshals IsisLspPrefixAttributes from protobuf object *otg.IsisLspPrefixAttributes
	// and doesn't set defaults
	setMsg(*otg.IsisLspPrefixAttributes) IsisLspPrefixAttributes
	// provides marshal interface
	Marshal() marshalIsisLspPrefixAttributes
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspPrefixAttributes
	// validate validates IsisLspPrefixAttributes
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspPrefixAttributes, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// XFlag returns bool, set in IsisLspPrefixAttributes.
	XFlag() bool
	// SetXFlag assigns bool provided by user to IsisLspPrefixAttributes
	SetXFlag(value bool) IsisLspPrefixAttributes
	// HasXFlag checks if XFlag has been set in IsisLspPrefixAttributes
	HasXFlag() bool
	// RFlag returns bool, set in IsisLspPrefixAttributes.
	RFlag() bool
	// SetRFlag assigns bool provided by user to IsisLspPrefixAttributes
	SetRFlag(value bool) IsisLspPrefixAttributes
	// HasRFlag checks if RFlag has been set in IsisLspPrefixAttributes
	HasRFlag() bool
	// NFlag returns bool, set in IsisLspPrefixAttributes.
	NFlag() bool
	// SetNFlag assigns bool provided by user to IsisLspPrefixAttributes
	SetNFlag(value bool) IsisLspPrefixAttributes
	// HasNFlag checks if NFlag has been set in IsisLspPrefixAttributes
	HasNFlag() bool
}

// External Prefix Flag (Bit 0)
// XFlag returns a bool
func (obj *isisLspPrefixAttributes) XFlag() bool {

	return *obj.obj.XFlag

}

// External Prefix Flag (Bit 0)
// XFlag returns a bool
func (obj *isisLspPrefixAttributes) HasXFlag() bool {
	return obj.obj.XFlag != nil
}

// External Prefix Flag (Bit 0)
// SetXFlag sets the bool value in the IsisLspPrefixAttributes object
func (obj *isisLspPrefixAttributes) SetXFlag(value bool) IsisLspPrefixAttributes {

	obj.obj.XFlag = &value
	return obj
}

// Re-advertisement Flag (Bit 1)
// RFlag returns a bool
func (obj *isisLspPrefixAttributes) RFlag() bool {

	return *obj.obj.RFlag

}

// Re-advertisement Flag (Bit 1)
// RFlag returns a bool
func (obj *isisLspPrefixAttributes) HasRFlag() bool {
	return obj.obj.RFlag != nil
}

// Re-advertisement Flag (Bit 1)
// SetRFlag sets the bool value in the IsisLspPrefixAttributes object
func (obj *isisLspPrefixAttributes) SetRFlag(value bool) IsisLspPrefixAttributes {

	obj.obj.RFlag = &value
	return obj
}

// Node Flag (Bit 2)
// NFlag returns a bool
func (obj *isisLspPrefixAttributes) NFlag() bool {

	return *obj.obj.NFlag

}

// Node Flag (Bit 2)
// NFlag returns a bool
func (obj *isisLspPrefixAttributes) HasNFlag() bool {
	return obj.obj.NFlag != nil
}

// Node Flag (Bit 2)
// SetNFlag sets the bool value in the IsisLspPrefixAttributes object
func (obj *isisLspPrefixAttributes) SetNFlag(value bool) IsisLspPrefixAttributes {

	obj.obj.NFlag = &value
	return obj
}

func (obj *isisLspPrefixAttributes) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisLspPrefixAttributes) setDefault() {

}

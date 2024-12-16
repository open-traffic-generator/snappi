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

// IsisLspPrefixAttributes is one bit value of ISIS Prefix attributes for the extended IPv4 and IPv6 reachability. https://www.rfc-editor.org/rfc/rfc7794.html.
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
	// External returns bool, set in IsisLspPrefixAttributes.
	External() bool
	// SetExternal assigns bool provided by user to IsisLspPrefixAttributes
	SetExternal(value bool) IsisLspPrefixAttributes
	// HasExternal checks if External has been set in IsisLspPrefixAttributes
	HasExternal() bool
	// Redistribution returns bool, set in IsisLspPrefixAttributes.
	Redistribution() bool
	// SetRedistribution assigns bool provided by user to IsisLspPrefixAttributes
	SetRedistribution(value bool) IsisLspPrefixAttributes
	// HasRedistribution checks if Redistribution has been set in IsisLspPrefixAttributes
	HasRedistribution() bool
	// Node returns bool, set in IsisLspPrefixAttributes.
	Node() bool
	// SetNode assigns bool provided by user to IsisLspPrefixAttributes
	SetNode(value bool) IsisLspPrefixAttributes
	// HasNode checks if Node has been set in IsisLspPrefixAttributes
	HasNode() bool
}

// External prefix flag (Bit 0). Set if the prefix has been
// redistributed from another protocol. This includes
// the case where multiple virtual routers are
// supported and the source of the redistributed prefix
// is another IS-IS instance.
// External returns a bool
func (obj *isisLspPrefixAttributes) External() bool {

	return *obj.obj.External

}

// External prefix flag (Bit 0). Set if the prefix has been
// redistributed from another protocol. This includes
// the case where multiple virtual routers are
// supported and the source of the redistributed prefix
// is another IS-IS instance.
// External returns a bool
func (obj *isisLspPrefixAttributes) HasExternal() bool {
	return obj.obj.External != nil
}

// External prefix flag (Bit 0). Set if the prefix has been
// redistributed from another protocol. This includes
// the case where multiple virtual routers are
// supported and the source of the redistributed prefix
// is another IS-IS instance.
// SetExternal sets the bool value in the IsisLspPrefixAttributes object
func (obj *isisLspPrefixAttributes) SetExternal(value bool) IsisLspPrefixAttributes {

	obj.obj.External = &value
	return obj
}

// Readvertisement flag (Bit 1). Set when the prefix has been
// leaked from one level to another (upwards or
// downwards).
// Redistribution returns a bool
func (obj *isisLspPrefixAttributes) Redistribution() bool {

	return *obj.obj.Redistribution

}

// Readvertisement flag (Bit 1). Set when the prefix has been
// leaked from one level to another (upwards or
// downwards).
// Redistribution returns a bool
func (obj *isisLspPrefixAttributes) HasRedistribution() bool {
	return obj.obj.Redistribution != nil
}

// Readvertisement flag (Bit 1). Set when the prefix has been
// leaked from one level to another (upwards or
// downwards).
// SetRedistribution sets the bool value in the IsisLspPrefixAttributes object
func (obj *isisLspPrefixAttributes) SetRedistribution(value bool) IsisLspPrefixAttributes {

	obj.obj.Redistribution = &value
	return obj
}

// Node Flag (Bit 2).
// Set when the prefix identifies the
// advertising router, i.e., the prefix is a host
// prefix advertising  a globally reachable address
// typically associated with a loopback address.
// Node returns a bool
func (obj *isisLspPrefixAttributes) Node() bool {

	return *obj.obj.Node

}

// Node Flag (Bit 2).
// Set when the prefix identifies the
// advertising router, i.e., the prefix is a host
// prefix advertising  a globally reachable address
// typically associated with a loopback address.
// Node returns a bool
func (obj *isisLspPrefixAttributes) HasNode() bool {
	return obj.obj.Node != nil
}

// Node Flag (Bit 2).
// Set when the prefix identifies the
// advertising router, i.e., the prefix is a host
// prefix advertising  a globally reachable address
// typically associated with a loopback address.
// SetNode sets the bool value in the IsisLspPrefixAttributes object
func (obj *isisLspPrefixAttributes) SetNode(value bool) IsisLspPrefixAttributes {

	obj.obj.Node = &value
	return obj
}

func (obj *isisLspPrefixAttributes) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisLspPrefixAttributes) setDefault() {

}

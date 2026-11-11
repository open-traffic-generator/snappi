package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIcmpv6NdpLinkLayerAddressOption *****
type flowIcmpv6NdpLinkLayerAddressOption struct {
	validation
	obj                    *otg.FlowIcmpv6NdpLinkLayerAddressOption
	marshaller             marshalFlowIcmpv6NdpLinkLayerAddressOption
	unMarshaller           unMarshalFlowIcmpv6NdpLinkLayerAddressOption
	lengthHolder           PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	linkLayerAddressHolder PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
}

func NewFlowIcmpv6NdpLinkLayerAddressOption() FlowIcmpv6NdpLinkLayerAddressOption {
	obj := flowIcmpv6NdpLinkLayerAddressOption{obj: &otg.FlowIcmpv6NdpLinkLayerAddressOption{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIcmpv6NdpLinkLayerAddressOption) msg() *otg.FlowIcmpv6NdpLinkLayerAddressOption {
	return obj.obj
}

func (obj *flowIcmpv6NdpLinkLayerAddressOption) setMsg(msg *otg.FlowIcmpv6NdpLinkLayerAddressOption) FlowIcmpv6NdpLinkLayerAddressOption {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIcmpv6NdpLinkLayerAddressOption struct {
	obj *flowIcmpv6NdpLinkLayerAddressOption
}

type marshalFlowIcmpv6NdpLinkLayerAddressOption interface {
	// ToProto marshals FlowIcmpv6NdpLinkLayerAddressOption to protobuf object *otg.FlowIcmpv6NdpLinkLayerAddressOption
	ToProto() (*otg.FlowIcmpv6NdpLinkLayerAddressOption, error)
	// ToPbText marshals FlowIcmpv6NdpLinkLayerAddressOption to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIcmpv6NdpLinkLayerAddressOption to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIcmpv6NdpLinkLayerAddressOption to JSON text
	ToJson() (string, error)
}

type unMarshalflowIcmpv6NdpLinkLayerAddressOption struct {
	obj *flowIcmpv6NdpLinkLayerAddressOption
}

type unMarshalFlowIcmpv6NdpLinkLayerAddressOption interface {
	// FromProto unmarshals FlowIcmpv6NdpLinkLayerAddressOption from protobuf object *otg.FlowIcmpv6NdpLinkLayerAddressOption
	FromProto(msg *otg.FlowIcmpv6NdpLinkLayerAddressOption) (FlowIcmpv6NdpLinkLayerAddressOption, error)
	// FromPbText unmarshals FlowIcmpv6NdpLinkLayerAddressOption from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIcmpv6NdpLinkLayerAddressOption from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIcmpv6NdpLinkLayerAddressOption from JSON text
	FromJson(value string) error
}

func (obj *flowIcmpv6NdpLinkLayerAddressOption) Marshal() marshalFlowIcmpv6NdpLinkLayerAddressOption {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIcmpv6NdpLinkLayerAddressOption{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIcmpv6NdpLinkLayerAddressOption) Unmarshal() unMarshalFlowIcmpv6NdpLinkLayerAddressOption {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIcmpv6NdpLinkLayerAddressOption{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIcmpv6NdpLinkLayerAddressOption) ToProto() (*otg.FlowIcmpv6NdpLinkLayerAddressOption, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIcmpv6NdpLinkLayerAddressOption) FromProto(msg *otg.FlowIcmpv6NdpLinkLayerAddressOption) (FlowIcmpv6NdpLinkLayerAddressOption, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIcmpv6NdpLinkLayerAddressOption) ToPbText() (string, error) {
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

func (m *unMarshalflowIcmpv6NdpLinkLayerAddressOption) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalflowIcmpv6NdpLinkLayerAddressOption) ToYaml() (string, error) {
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

func (m *unMarshalflowIcmpv6NdpLinkLayerAddressOption) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalflowIcmpv6NdpLinkLayerAddressOption) ToJson() (string, error) {
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

func (m *unMarshalflowIcmpv6NdpLinkLayerAddressOption) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *flowIcmpv6NdpLinkLayerAddressOption) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIcmpv6NdpLinkLayerAddressOption) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIcmpv6NdpLinkLayerAddressOption) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIcmpv6NdpLinkLayerAddressOption) Clone() (FlowIcmpv6NdpLinkLayerAddressOption, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIcmpv6NdpLinkLayerAddressOption()
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

func (obj *flowIcmpv6NdpLinkLayerAddressOption) setNil() {
	obj.lengthHolder = nil
	obj.linkLayerAddressHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIcmpv6NdpLinkLayerAddressOption is nDP Source Link-Layer Address (Type 1) or Target Link-Layer Address (Type 2) option (RFC 4861 Section 4.6.1). The option Type byte is set from the parent choice.
type FlowIcmpv6NdpLinkLayerAddressOption interface {
	Validation
	// msg marshals FlowIcmpv6NdpLinkLayerAddressOption to protobuf object *otg.FlowIcmpv6NdpLinkLayerAddressOption
	// and doesn't set defaults
	msg() *otg.FlowIcmpv6NdpLinkLayerAddressOption
	// setMsg unmarshals FlowIcmpv6NdpLinkLayerAddressOption from protobuf object *otg.FlowIcmpv6NdpLinkLayerAddressOption
	// and doesn't set defaults
	setMsg(*otg.FlowIcmpv6NdpLinkLayerAddressOption) FlowIcmpv6NdpLinkLayerAddressOption
	// provides marshal interface
	Marshal() marshalFlowIcmpv6NdpLinkLayerAddressOption
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIcmpv6NdpLinkLayerAddressOption
	// validate validates FlowIcmpv6NdpLinkLayerAddressOption
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIcmpv6NdpLinkLayerAddressOption, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Length returns PatternFlowIcmpv6NdpLinkLayerAddressOptionLength, set in FlowIcmpv6NdpLinkLayerAddressOption.
	// PatternFlowIcmpv6NdpLinkLayerAddressOptionLength is option length in units of 8 octets, including the type and length bytes. For a 6-octet Ethernet address this is 1.
	Length() PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	// SetLength assigns PatternFlowIcmpv6NdpLinkLayerAddressOptionLength provided by user to FlowIcmpv6NdpLinkLayerAddressOption.
	// PatternFlowIcmpv6NdpLinkLayerAddressOptionLength is option length in units of 8 octets, including the type and length bytes. For a 6-octet Ethernet address this is 1.
	SetLength(value PatternFlowIcmpv6NdpLinkLayerAddressOptionLength) FlowIcmpv6NdpLinkLayerAddressOption
	// HasLength checks if Length has been set in FlowIcmpv6NdpLinkLayerAddressOption
	HasLength() bool
	// LinkLayerAddress returns PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress, set in FlowIcmpv6NdpLinkLayerAddressOption.
	// PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress is link-layer address of the sender (source option) or target (target option).
	LinkLayerAddress() PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	// SetLinkLayerAddress assigns PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress provided by user to FlowIcmpv6NdpLinkLayerAddressOption.
	// PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress is link-layer address of the sender (source option) or target (target option).
	SetLinkLayerAddress(value PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) FlowIcmpv6NdpLinkLayerAddressOption
	// HasLinkLayerAddress checks if LinkLayerAddress has been set in FlowIcmpv6NdpLinkLayerAddressOption
	HasLinkLayerAddress() bool
	setNil()
}

// description is TBD
// Length returns a PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
func (obj *flowIcmpv6NdpLinkLayerAddressOption) Length() PatternFlowIcmpv6NdpLinkLayerAddressOptionLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewPatternFlowIcmpv6NdpLinkLayerAddressOptionLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &patternFlowIcmpv6NdpLinkLayerAddressOptionLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
func (obj *flowIcmpv6NdpLinkLayerAddressOption) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the PatternFlowIcmpv6NdpLinkLayerAddressOptionLength value in the FlowIcmpv6NdpLinkLayerAddressOption object
func (obj *flowIcmpv6NdpLinkLayerAddressOption) SetLength(value PatternFlowIcmpv6NdpLinkLayerAddressOptionLength) FlowIcmpv6NdpLinkLayerAddressOption {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// LinkLayerAddress returns a PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
func (obj *flowIcmpv6NdpLinkLayerAddressOption) LinkLayerAddress() PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress {
	if obj.obj.LinkLayerAddress == nil {
		obj.obj.LinkLayerAddress = NewPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress().msg()
	}
	if obj.linkLayerAddressHolder == nil {
		obj.linkLayerAddressHolder = &patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress{obj: obj.obj.LinkLayerAddress}
	}
	return obj.linkLayerAddressHolder
}

// description is TBD
// LinkLayerAddress returns a PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
func (obj *flowIcmpv6NdpLinkLayerAddressOption) HasLinkLayerAddress() bool {
	return obj.obj.LinkLayerAddress != nil
}

// description is TBD
// SetLinkLayerAddress sets the PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress value in the FlowIcmpv6NdpLinkLayerAddressOption object
func (obj *flowIcmpv6NdpLinkLayerAddressOption) SetLinkLayerAddress(value PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) FlowIcmpv6NdpLinkLayerAddressOption {

	obj.linkLayerAddressHolder = nil
	obj.obj.LinkLayerAddress = value.msg()

	return obj
}

func (obj *flowIcmpv6NdpLinkLayerAddressOption) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.LinkLayerAddress != nil {

		obj.LinkLayerAddress().validateObj(vObj, set_default)
	}

}

func (obj *flowIcmpv6NdpLinkLayerAddressOption) setDefault() {

}

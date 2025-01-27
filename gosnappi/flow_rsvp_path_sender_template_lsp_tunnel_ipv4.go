package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathSenderTemplateLspTunnelIpv4 *****
type flowRSVPPathSenderTemplateLspTunnelIpv4 struct {
	validation
	obj                           *otg.FlowRSVPPathSenderTemplateLspTunnelIpv4
	marshaller                    marshalFlowRSVPPathSenderTemplateLspTunnelIpv4
	unMarshaller                  unMarshalFlowRSVPPathSenderTemplateLspTunnelIpv4
	ipv4TunnelSenderAddressHolder PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	reservedHolder                PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	lspIdHolder                   PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
}

func NewFlowRSVPPathSenderTemplateLspTunnelIpv4() FlowRSVPPathSenderTemplateLspTunnelIpv4 {
	obj := flowRSVPPathSenderTemplateLspTunnelIpv4{obj: &otg.FlowRSVPPathSenderTemplateLspTunnelIpv4{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathSenderTemplateLspTunnelIpv4) msg() *otg.FlowRSVPPathSenderTemplateLspTunnelIpv4 {
	return obj.obj
}

func (obj *flowRSVPPathSenderTemplateLspTunnelIpv4) setMsg(msg *otg.FlowRSVPPathSenderTemplateLspTunnelIpv4) FlowRSVPPathSenderTemplateLspTunnelIpv4 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathSenderTemplateLspTunnelIpv4 struct {
	obj *flowRSVPPathSenderTemplateLspTunnelIpv4
}

type marshalFlowRSVPPathSenderTemplateLspTunnelIpv4 interface {
	// ToProto marshals FlowRSVPPathSenderTemplateLspTunnelIpv4 to protobuf object *otg.FlowRSVPPathSenderTemplateLspTunnelIpv4
	ToProto() (*otg.FlowRSVPPathSenderTemplateLspTunnelIpv4, error)
	// ToPbText marshals FlowRSVPPathSenderTemplateLspTunnelIpv4 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathSenderTemplateLspTunnelIpv4 to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathSenderTemplateLspTunnelIpv4 to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowRSVPPathSenderTemplateLspTunnelIpv4 to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowRSVPPathSenderTemplateLspTunnelIpv4 struct {
	obj *flowRSVPPathSenderTemplateLspTunnelIpv4
}

type unMarshalFlowRSVPPathSenderTemplateLspTunnelIpv4 interface {
	// FromProto unmarshals FlowRSVPPathSenderTemplateLspTunnelIpv4 from protobuf object *otg.FlowRSVPPathSenderTemplateLspTunnelIpv4
	FromProto(msg *otg.FlowRSVPPathSenderTemplateLspTunnelIpv4) (FlowRSVPPathSenderTemplateLspTunnelIpv4, error)
	// FromPbText unmarshals FlowRSVPPathSenderTemplateLspTunnelIpv4 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathSenderTemplateLspTunnelIpv4 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathSenderTemplateLspTunnelIpv4 from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathSenderTemplateLspTunnelIpv4) Marshal() marshalFlowRSVPPathSenderTemplateLspTunnelIpv4 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathSenderTemplateLspTunnelIpv4{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathSenderTemplateLspTunnelIpv4) Unmarshal() unMarshalFlowRSVPPathSenderTemplateLspTunnelIpv4 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathSenderTemplateLspTunnelIpv4{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathSenderTemplateLspTunnelIpv4) ToProto() (*otg.FlowRSVPPathSenderTemplateLspTunnelIpv4, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathSenderTemplateLspTunnelIpv4) FromProto(msg *otg.FlowRSVPPathSenderTemplateLspTunnelIpv4) (FlowRSVPPathSenderTemplateLspTunnelIpv4, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathSenderTemplateLspTunnelIpv4) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathSenderTemplateLspTunnelIpv4) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathSenderTemplateLspTunnelIpv4) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathSenderTemplateLspTunnelIpv4) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathSenderTemplateLspTunnelIpv4) ToJsonRaw() (string, error) {
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

func (m *marshalflowRSVPPathSenderTemplateLspTunnelIpv4) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathSenderTemplateLspTunnelIpv4) FromJson(value string) error {
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

func (obj *flowRSVPPathSenderTemplateLspTunnelIpv4) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathSenderTemplateLspTunnelIpv4) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathSenderTemplateLspTunnelIpv4) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathSenderTemplateLspTunnelIpv4) Clone() (FlowRSVPPathSenderTemplateLspTunnelIpv4, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathSenderTemplateLspTunnelIpv4()
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

func (obj *flowRSVPPathSenderTemplateLspTunnelIpv4) setNil() {
	obj.ipv4TunnelSenderAddressHolder = nil
	obj.reservedHolder = nil
	obj.lspIdHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathSenderTemplateLspTunnelIpv4 is class = SENDER_TEMPLATE, LSP_TUNNEL_IPv4 C-Type = 7
type FlowRSVPPathSenderTemplateLspTunnelIpv4 interface {
	Validation
	// msg marshals FlowRSVPPathSenderTemplateLspTunnelIpv4 to protobuf object *otg.FlowRSVPPathSenderTemplateLspTunnelIpv4
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathSenderTemplateLspTunnelIpv4
	// setMsg unmarshals FlowRSVPPathSenderTemplateLspTunnelIpv4 from protobuf object *otg.FlowRSVPPathSenderTemplateLspTunnelIpv4
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathSenderTemplateLspTunnelIpv4) FlowRSVPPathSenderTemplateLspTunnelIpv4
	// provides marshal interface
	Marshal() marshalFlowRSVPPathSenderTemplateLspTunnelIpv4
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathSenderTemplateLspTunnelIpv4
	// validate validates FlowRSVPPathSenderTemplateLspTunnelIpv4
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathSenderTemplateLspTunnelIpv4, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4TunnelSenderAddress returns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress, set in FlowRSVPPathSenderTemplateLspTunnelIpv4.
	// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress is iPv4 address for a sender node.
	Ipv4TunnelSenderAddress() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	// SetIpv4TunnelSenderAddress assigns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress provided by user to FlowRSVPPathSenderTemplateLspTunnelIpv4.
	// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress is iPv4 address for a sender node.
	SetIpv4TunnelSenderAddress(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) FlowRSVPPathSenderTemplateLspTunnelIpv4
	// HasIpv4TunnelSenderAddress checks if Ipv4TunnelSenderAddress has been set in FlowRSVPPathSenderTemplateLspTunnelIpv4
	HasIpv4TunnelSenderAddress() bool
	// Reserved returns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved, set in FlowRSVPPathSenderTemplateLspTunnelIpv4.
	// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved is reserved field, MUST be zero.
	Reserved() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	// SetReserved assigns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved provided by user to FlowRSVPPathSenderTemplateLspTunnelIpv4.
	// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved is reserved field, MUST be zero.
	SetReserved(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) FlowRSVPPathSenderTemplateLspTunnelIpv4
	// HasReserved checks if Reserved has been set in FlowRSVPPathSenderTemplateLspTunnelIpv4
	HasReserved() bool
	// LspId returns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId, set in FlowRSVPPathSenderTemplateLspTunnelIpv4.
	// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId is a 16-bit identifier used in the SENDER_TEMPLATE that can be changed to allow a sender to share resources with itself.
	LspId() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	// SetLspId assigns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId provided by user to FlowRSVPPathSenderTemplateLspTunnelIpv4.
	// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId is a 16-bit identifier used in the SENDER_TEMPLATE that can be changed to allow a sender to share resources with itself.
	SetLspId(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) FlowRSVPPathSenderTemplateLspTunnelIpv4
	// HasLspId checks if LspId has been set in FlowRSVPPathSenderTemplateLspTunnelIpv4
	HasLspId() bool
	setNil()
}

// description is TBD
// Ipv4TunnelSenderAddress returns a PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
func (obj *flowRSVPPathSenderTemplateLspTunnelIpv4) Ipv4TunnelSenderAddress() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress {
	if obj.obj.Ipv4TunnelSenderAddress == nil {
		obj.obj.Ipv4TunnelSenderAddress = NewPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress().msg()
	}
	if obj.ipv4TunnelSenderAddressHolder == nil {
		obj.ipv4TunnelSenderAddressHolder = &patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress{obj: obj.obj.Ipv4TunnelSenderAddress}
	}
	return obj.ipv4TunnelSenderAddressHolder
}

// description is TBD
// Ipv4TunnelSenderAddress returns a PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
func (obj *flowRSVPPathSenderTemplateLspTunnelIpv4) HasIpv4TunnelSenderAddress() bool {
	return obj.obj.Ipv4TunnelSenderAddress != nil
}

// description is TBD
// SetIpv4TunnelSenderAddress sets the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress value in the FlowRSVPPathSenderTemplateLspTunnelIpv4 object
func (obj *flowRSVPPathSenderTemplateLspTunnelIpv4) SetIpv4TunnelSenderAddress(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) FlowRSVPPathSenderTemplateLspTunnelIpv4 {

	obj.ipv4TunnelSenderAddressHolder = nil
	obj.obj.Ipv4TunnelSenderAddress = value.msg()

	return obj
}

// description is TBD
// Reserved returns a PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
func (obj *flowRSVPPathSenderTemplateLspTunnelIpv4) Reserved() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = NewPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved().msg()
	}
	if obj.reservedHolder == nil {
		obj.reservedHolder = &patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved{obj: obj.obj.Reserved}
	}
	return obj.reservedHolder
}

// description is TBD
// Reserved returns a PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
func (obj *flowRSVPPathSenderTemplateLspTunnelIpv4) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// description is TBD
// SetReserved sets the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved value in the FlowRSVPPathSenderTemplateLspTunnelIpv4 object
func (obj *flowRSVPPathSenderTemplateLspTunnelIpv4) SetReserved(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) FlowRSVPPathSenderTemplateLspTunnelIpv4 {

	obj.reservedHolder = nil
	obj.obj.Reserved = value.msg()

	return obj
}

// description is TBD
// LspId returns a PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
func (obj *flowRSVPPathSenderTemplateLspTunnelIpv4) LspId() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId {
	if obj.obj.LspId == nil {
		obj.obj.LspId = NewPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId().msg()
	}
	if obj.lspIdHolder == nil {
		obj.lspIdHolder = &patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId{obj: obj.obj.LspId}
	}
	return obj.lspIdHolder
}

// description is TBD
// LspId returns a PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
func (obj *flowRSVPPathSenderTemplateLspTunnelIpv4) HasLspId() bool {
	return obj.obj.LspId != nil
}

// description is TBD
// SetLspId sets the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId value in the FlowRSVPPathSenderTemplateLspTunnelIpv4 object
func (obj *flowRSVPPathSenderTemplateLspTunnelIpv4) SetLspId(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) FlowRSVPPathSenderTemplateLspTunnelIpv4 {

	obj.lspIdHolder = nil
	obj.obj.LspId = value.msg()

	return obj
}

func (obj *flowRSVPPathSenderTemplateLspTunnelIpv4) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv4TunnelSenderAddress != nil {

		obj.Ipv4TunnelSenderAddress().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved != nil {

		obj.Reserved().validateObj(vObj, set_default)
	}

	if obj.obj.LspId != nil {

		obj.LspId().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathSenderTemplateLspTunnelIpv4) setDefault() {

}

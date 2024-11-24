package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathSessionLspTunnelIpv4 *****
type flowRSVPPathSessionLspTunnelIpv4 struct {
	validation
	obj                             *otg.FlowRSVPPathSessionLspTunnelIpv4
	marshaller                      marshalFlowRSVPPathSessionLspTunnelIpv4
	unMarshaller                    unMarshalFlowRSVPPathSessionLspTunnelIpv4
	ipv4TunnelEndPointAddressHolder PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	reservedHolder                  PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	tunnelIdHolder                  PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	extendedTunnelIdHolder          FlowRSVPPathSessionExtTunnelId
}

func NewFlowRSVPPathSessionLspTunnelIpv4() FlowRSVPPathSessionLspTunnelIpv4 {
	obj := flowRSVPPathSessionLspTunnelIpv4{obj: &otg.FlowRSVPPathSessionLspTunnelIpv4{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathSessionLspTunnelIpv4) msg() *otg.FlowRSVPPathSessionLspTunnelIpv4 {
	return obj.obj
}

func (obj *flowRSVPPathSessionLspTunnelIpv4) setMsg(msg *otg.FlowRSVPPathSessionLspTunnelIpv4) FlowRSVPPathSessionLspTunnelIpv4 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathSessionLspTunnelIpv4 struct {
	obj *flowRSVPPathSessionLspTunnelIpv4
}

type marshalFlowRSVPPathSessionLspTunnelIpv4 interface {
	// ToProto marshals FlowRSVPPathSessionLspTunnelIpv4 to protobuf object *otg.FlowRSVPPathSessionLspTunnelIpv4
	ToProto() (*otg.FlowRSVPPathSessionLspTunnelIpv4, error)
	// ToPbText marshals FlowRSVPPathSessionLspTunnelIpv4 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathSessionLspTunnelIpv4 to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathSessionLspTunnelIpv4 to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowRSVPPathSessionLspTunnelIpv4 to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowRSVPPathSessionLspTunnelIpv4 struct {
	obj *flowRSVPPathSessionLspTunnelIpv4
}

type unMarshalFlowRSVPPathSessionLspTunnelIpv4 interface {
	// FromProto unmarshals FlowRSVPPathSessionLspTunnelIpv4 from protobuf object *otg.FlowRSVPPathSessionLspTunnelIpv4
	FromProto(msg *otg.FlowRSVPPathSessionLspTunnelIpv4) (FlowRSVPPathSessionLspTunnelIpv4, error)
	// FromPbText unmarshals FlowRSVPPathSessionLspTunnelIpv4 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathSessionLspTunnelIpv4 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathSessionLspTunnelIpv4 from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathSessionLspTunnelIpv4) Marshal() marshalFlowRSVPPathSessionLspTunnelIpv4 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathSessionLspTunnelIpv4{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathSessionLspTunnelIpv4) Unmarshal() unMarshalFlowRSVPPathSessionLspTunnelIpv4 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathSessionLspTunnelIpv4{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathSessionLspTunnelIpv4) ToProto() (*otg.FlowRSVPPathSessionLspTunnelIpv4, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathSessionLspTunnelIpv4) FromProto(msg *otg.FlowRSVPPathSessionLspTunnelIpv4) (FlowRSVPPathSessionLspTunnelIpv4, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathSessionLspTunnelIpv4) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathSessionLspTunnelIpv4) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathSessionLspTunnelIpv4) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathSessionLspTunnelIpv4) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathSessionLspTunnelIpv4) ToJsonRaw() (string, error) {
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

func (m *marshalflowRSVPPathSessionLspTunnelIpv4) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathSessionLspTunnelIpv4) FromJson(value string) error {
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

func (obj *flowRSVPPathSessionLspTunnelIpv4) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathSessionLspTunnelIpv4) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathSessionLspTunnelIpv4) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathSessionLspTunnelIpv4) Clone() (FlowRSVPPathSessionLspTunnelIpv4, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathSessionLspTunnelIpv4()
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

func (obj *flowRSVPPathSessionLspTunnelIpv4) setNil() {
	obj.ipv4TunnelEndPointAddressHolder = nil
	obj.reservedHolder = nil
	obj.tunnelIdHolder = nil
	obj.extendedTunnelIdHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathSessionLspTunnelIpv4 is class = SESSION, LSP_TUNNEL_IPv4 C-Type = 7.
type FlowRSVPPathSessionLspTunnelIpv4 interface {
	Validation
	// msg marshals FlowRSVPPathSessionLspTunnelIpv4 to protobuf object *otg.FlowRSVPPathSessionLspTunnelIpv4
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathSessionLspTunnelIpv4
	// setMsg unmarshals FlowRSVPPathSessionLspTunnelIpv4 from protobuf object *otg.FlowRSVPPathSessionLspTunnelIpv4
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathSessionLspTunnelIpv4) FlowRSVPPathSessionLspTunnelIpv4
	// provides marshal interface
	Marshal() marshalFlowRSVPPathSessionLspTunnelIpv4
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathSessionLspTunnelIpv4
	// validate validates FlowRSVPPathSessionLspTunnelIpv4
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathSessionLspTunnelIpv4, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4TunnelEndPointAddress returns PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress, set in FlowRSVPPathSessionLspTunnelIpv4.
	// PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress is iPv4 address of the egress node for the tunnel.
	Ipv4TunnelEndPointAddress() PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	// SetIpv4TunnelEndPointAddress assigns PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress provided by user to FlowRSVPPathSessionLspTunnelIpv4.
	// PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress is iPv4 address of the egress node for the tunnel.
	SetIpv4TunnelEndPointAddress(value PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) FlowRSVPPathSessionLspTunnelIpv4
	// HasIpv4TunnelEndPointAddress checks if Ipv4TunnelEndPointAddress has been set in FlowRSVPPathSessionLspTunnelIpv4
	HasIpv4TunnelEndPointAddress() bool
	// Reserved returns PatternFlowRSVPPathSessionLspTunnelIpv4Reserved, set in FlowRSVPPathSessionLspTunnelIpv4.
	// PatternFlowRSVPPathSessionLspTunnelIpv4Reserved is reserved field, MUST be zero.
	Reserved() PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	// SetReserved assigns PatternFlowRSVPPathSessionLspTunnelIpv4Reserved provided by user to FlowRSVPPathSessionLspTunnelIpv4.
	// PatternFlowRSVPPathSessionLspTunnelIpv4Reserved is reserved field, MUST be zero.
	SetReserved(value PatternFlowRSVPPathSessionLspTunnelIpv4Reserved) FlowRSVPPathSessionLspTunnelIpv4
	// HasReserved checks if Reserved has been set in FlowRSVPPathSessionLspTunnelIpv4
	HasReserved() bool
	// TunnelId returns PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId, set in FlowRSVPPathSessionLspTunnelIpv4.
	// PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId is a 16-bit identifier used in the SESSION that remains constant over the life of the tunnel.
	TunnelId() PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	// SetTunnelId assigns PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId provided by user to FlowRSVPPathSessionLspTunnelIpv4.
	// PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId is a 16-bit identifier used in the SESSION that remains constant over the life of the tunnel.
	SetTunnelId(value PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId) FlowRSVPPathSessionLspTunnelIpv4
	// HasTunnelId checks if TunnelId has been set in FlowRSVPPathSessionLspTunnelIpv4
	HasTunnelId() bool
	// ExtendedTunnelId returns FlowRSVPPathSessionExtTunnelId, set in FlowRSVPPathSessionLspTunnelIpv4.
	// FlowRSVPPathSessionExtTunnelId is description is TBD
	ExtendedTunnelId() FlowRSVPPathSessionExtTunnelId
	// SetExtendedTunnelId assigns FlowRSVPPathSessionExtTunnelId provided by user to FlowRSVPPathSessionLspTunnelIpv4.
	// FlowRSVPPathSessionExtTunnelId is description is TBD
	SetExtendedTunnelId(value FlowRSVPPathSessionExtTunnelId) FlowRSVPPathSessionLspTunnelIpv4
	// HasExtendedTunnelId checks if ExtendedTunnelId has been set in FlowRSVPPathSessionLspTunnelIpv4
	HasExtendedTunnelId() bool
	setNil()
}

// description is TBD
// Ipv4TunnelEndPointAddress returns a PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
func (obj *flowRSVPPathSessionLspTunnelIpv4) Ipv4TunnelEndPointAddress() PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress {
	if obj.obj.Ipv4TunnelEndPointAddress == nil {
		obj.obj.Ipv4TunnelEndPointAddress = NewPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress().msg()
	}
	if obj.ipv4TunnelEndPointAddressHolder == nil {
		obj.ipv4TunnelEndPointAddressHolder = &patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress{obj: obj.obj.Ipv4TunnelEndPointAddress}
	}
	return obj.ipv4TunnelEndPointAddressHolder
}

// description is TBD
// Ipv4TunnelEndPointAddress returns a PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
func (obj *flowRSVPPathSessionLspTunnelIpv4) HasIpv4TunnelEndPointAddress() bool {
	return obj.obj.Ipv4TunnelEndPointAddress != nil
}

// description is TBD
// SetIpv4TunnelEndPointAddress sets the PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress value in the FlowRSVPPathSessionLspTunnelIpv4 object
func (obj *flowRSVPPathSessionLspTunnelIpv4) SetIpv4TunnelEndPointAddress(value PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) FlowRSVPPathSessionLspTunnelIpv4 {

	obj.ipv4TunnelEndPointAddressHolder = nil
	obj.obj.Ipv4TunnelEndPointAddress = value.msg()

	return obj
}

// description is TBD
// Reserved returns a PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
func (obj *flowRSVPPathSessionLspTunnelIpv4) Reserved() PatternFlowRSVPPathSessionLspTunnelIpv4Reserved {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = NewPatternFlowRSVPPathSessionLspTunnelIpv4Reserved().msg()
	}
	if obj.reservedHolder == nil {
		obj.reservedHolder = &patternFlowRSVPPathSessionLspTunnelIpv4Reserved{obj: obj.obj.Reserved}
	}
	return obj.reservedHolder
}

// description is TBD
// Reserved returns a PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
func (obj *flowRSVPPathSessionLspTunnelIpv4) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// description is TBD
// SetReserved sets the PatternFlowRSVPPathSessionLspTunnelIpv4Reserved value in the FlowRSVPPathSessionLspTunnelIpv4 object
func (obj *flowRSVPPathSessionLspTunnelIpv4) SetReserved(value PatternFlowRSVPPathSessionLspTunnelIpv4Reserved) FlowRSVPPathSessionLspTunnelIpv4 {

	obj.reservedHolder = nil
	obj.obj.Reserved = value.msg()

	return obj
}

// description is TBD
// TunnelId returns a PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
func (obj *flowRSVPPathSessionLspTunnelIpv4) TunnelId() PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId {
	if obj.obj.TunnelId == nil {
		obj.obj.TunnelId = NewPatternFlowRSVPPathSessionLspTunnelIpv4TunnelId().msg()
	}
	if obj.tunnelIdHolder == nil {
		obj.tunnelIdHolder = &patternFlowRSVPPathSessionLspTunnelIpv4TunnelId{obj: obj.obj.TunnelId}
	}
	return obj.tunnelIdHolder
}

// description is TBD
// TunnelId returns a PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
func (obj *flowRSVPPathSessionLspTunnelIpv4) HasTunnelId() bool {
	return obj.obj.TunnelId != nil
}

// description is TBD
// SetTunnelId sets the PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId value in the FlowRSVPPathSessionLspTunnelIpv4 object
func (obj *flowRSVPPathSessionLspTunnelIpv4) SetTunnelId(value PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId) FlowRSVPPathSessionLspTunnelIpv4 {

	obj.tunnelIdHolder = nil
	obj.obj.TunnelId = value.msg()

	return obj
}

// A 32-bit identifier used in the SESSION that remains constant over the life of the tunnel. Normally set to all zeros. Ingress nodes that wish to narrow the scope of a SESSION to the ingress-egress pair may place their IPv4 address here as a globally unique identifier.
// ExtendedTunnelId returns a FlowRSVPPathSessionExtTunnelId
func (obj *flowRSVPPathSessionLspTunnelIpv4) ExtendedTunnelId() FlowRSVPPathSessionExtTunnelId {
	if obj.obj.ExtendedTunnelId == nil {
		obj.obj.ExtendedTunnelId = NewFlowRSVPPathSessionExtTunnelId().msg()
	}
	if obj.extendedTunnelIdHolder == nil {
		obj.extendedTunnelIdHolder = &flowRSVPPathSessionExtTunnelId{obj: obj.obj.ExtendedTunnelId}
	}
	return obj.extendedTunnelIdHolder
}

// A 32-bit identifier used in the SESSION that remains constant over the life of the tunnel. Normally set to all zeros. Ingress nodes that wish to narrow the scope of a SESSION to the ingress-egress pair may place their IPv4 address here as a globally unique identifier.
// ExtendedTunnelId returns a FlowRSVPPathSessionExtTunnelId
func (obj *flowRSVPPathSessionLspTunnelIpv4) HasExtendedTunnelId() bool {
	return obj.obj.ExtendedTunnelId != nil
}

// A 32-bit identifier used in the SESSION that remains constant over the life of the tunnel. Normally set to all zeros. Ingress nodes that wish to narrow the scope of a SESSION to the ingress-egress pair may place their IPv4 address here as a globally unique identifier.
// SetExtendedTunnelId sets the FlowRSVPPathSessionExtTunnelId value in the FlowRSVPPathSessionLspTunnelIpv4 object
func (obj *flowRSVPPathSessionLspTunnelIpv4) SetExtendedTunnelId(value FlowRSVPPathSessionExtTunnelId) FlowRSVPPathSessionLspTunnelIpv4 {

	obj.extendedTunnelIdHolder = nil
	obj.obj.ExtendedTunnelId = value.msg()

	return obj
}

func (obj *flowRSVPPathSessionLspTunnelIpv4) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv4TunnelEndPointAddress != nil {

		obj.Ipv4TunnelEndPointAddress().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved != nil {

		obj.Reserved().validateObj(vObj, set_default)
	}

	if obj.obj.TunnelId != nil {

		obj.TunnelId().validateObj(vObj, set_default)
	}

	if obj.obj.ExtendedTunnelId != nil {

		obj.ExtendedTunnelId().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathSessionLspTunnelIpv4) setDefault() {

}

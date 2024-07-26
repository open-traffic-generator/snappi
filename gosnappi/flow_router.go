package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRouter *****
type flowRouter struct {
	validation
	obj          *otg.FlowRouter
	marshaller   marshalFlowRouter
	unMarshaller unMarshalFlowRouter
}

func NewFlowRouter() FlowRouter {
	obj := flowRouter{obj: &otg.FlowRouter{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRouter) msg() *otg.FlowRouter {
	return obj.obj
}

func (obj *flowRouter) setMsg(msg *otg.FlowRouter) FlowRouter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRouter struct {
	obj *flowRouter
}

type marshalFlowRouter interface {
	// ToProto marshals FlowRouter to protobuf object *otg.FlowRouter
	ToProto() (*otg.FlowRouter, error)
	// ToPbText marshals FlowRouter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRouter to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRouter to JSON text
	ToJson() (string, error)
}

type unMarshalflowRouter struct {
	obj *flowRouter
}

type unMarshalFlowRouter interface {
	// FromProto unmarshals FlowRouter from protobuf object *otg.FlowRouter
	FromProto(msg *otg.FlowRouter) (FlowRouter, error)
	// FromPbText unmarshals FlowRouter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRouter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRouter from JSON text
	FromJson(value string) error
}

func (obj *flowRouter) Marshal() marshalFlowRouter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRouter{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRouter) Unmarshal() unMarshalFlowRouter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRouter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRouter) ToProto() (*otg.FlowRouter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRouter) FromProto(msg *otg.FlowRouter) (FlowRouter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRouter) ToPbText() (string, error) {
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

func (m *unMarshalflowRouter) FromPbText(value string) error {
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

func (m *marshalflowRouter) ToYaml() (string, error) {
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

func (m *unMarshalflowRouter) FromYaml(value string) error {
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

func (m *marshalflowRouter) ToJson() (string, error) {
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

func (m *unMarshalflowRouter) FromJson(value string) error {
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

func (obj *flowRouter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRouter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRouter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRouter) Clone() (FlowRouter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRouter()
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

// FlowRouter is a container for declaring a map of 1..n transmit devices to 1..n receive devices. This allows for a single flow to have  different tx to rx device flows such as a single one to one map or a  many to many map.
type FlowRouter interface {
	Validation
	// msg marshals FlowRouter to protobuf object *otg.FlowRouter
	// and doesn't set defaults
	msg() *otg.FlowRouter
	// setMsg unmarshals FlowRouter from protobuf object *otg.FlowRouter
	// and doesn't set defaults
	setMsg(*otg.FlowRouter) FlowRouter
	// provides marshal interface
	Marshal() marshalFlowRouter
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRouter
	// validate validates FlowRouter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRouter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Mode returns FlowRouterModeEnum, set in FlowRouter
	Mode() FlowRouterModeEnum
	// SetMode assigns FlowRouterModeEnum provided by user to FlowRouter
	SetMode(value FlowRouterModeEnum) FlowRouter
	// HasMode checks if Mode has been set in FlowRouter
	HasMode() bool
	// TxNames returns []string, set in FlowRouter.
	TxNames() []string
	// SetTxNames assigns []string provided by user to FlowRouter
	SetTxNames(value []string) FlowRouter
	// RxNames returns []string, set in FlowRouter.
	RxNames() []string
	// SetRxNames assigns []string provided by user to FlowRouter
	SetRxNames(value []string) FlowRouter
}

type FlowRouterModeEnum string

// Enum of Mode on FlowRouter
var FlowRouterMode = struct {
	MESH       FlowRouterModeEnum
	ONE_TO_ONE FlowRouterModeEnum
}{
	MESH:       FlowRouterModeEnum("mesh"),
	ONE_TO_ONE: FlowRouterModeEnum("one_to_one"),
}

func (obj *flowRouter) Mode() FlowRouterModeEnum {
	return FlowRouterModeEnum(obj.obj.Mode.Enum().String())
}

// Determines the intent of creating traffic sub-flow(s) between the device
// endpoints, from the entities of <b>tx_names</b> to the entities of <b>rx_names</b>
// to derive how <b>auto</b> packet fields can be populated with
// the actual value(s) by the implementation.
//
// The <b>one_to_one</b> mode creates traffic sub-flow(s) between each device endpoint pair in
// tx_names to rx_names by index.
// The length of tx_names and rx_names MUST be the same.
// The same device name can be repeated multiple times in tx_names or rx_names, in any order to create desired meshing between device(s).
// For 2 values in tx_names and 2 values in rx_names, 2 device endpoint pairs would be generated (each pair representing a traffic sub-flow).
//
// The <b>mesh</b> mode creates traffic sub-flow(s) between each value in tx_names to
// every value in rx_names, forming the device endpoint pair(s).
// For 2 values in tx_names and 3 values in rx_names, generated device endpoint pairs would be 2x3=6.
//
// A generated device endpoint pair with same device endpoint name for both transmit & receive device endpoint MUST raise an error.
//
// Packet fields of type <b>auto</b> would be populated with one value for each device endpoint pair (representing the traffic sub-flow).
// The value would be determined considering transmit & receive device of the sub-flow. And the sequence of the populated value(s)
// would be in the order of generated device endpoint pair(s).
// If 2 device endpoint pairs are generated (based on mode, tx_names and rx_names), say (d1 to d3) and (d2 to d3), and ethernet.dst is set as <b>auto</b>, then
// the auto field would be <b>replaced</b> by the implementation with a sequence of 2 values, [v1,v2] where
// v1 is determined using context (d1,d3) and v2 using context (d2,d3).
// The final outcome is that packets generated on the wire will contain the values v1,v2,v1,... for ethernet.dst field. Any non-auto packet fields
// should be configured accordingly. For example, non-auto packet field ethernet.src can be configured with values [u1, u2], where
// u1 & u2 are source MAC of the connected interface of device d1 and d2 respectively. Then packets on the wire will contain correct value pairs
// (u1,v1),(u2,v2),(u1,v1),... for (ethernet.src,ethernet.dst) fields.
// Mode returns a string
func (obj *flowRouter) HasMode() bool {
	return obj.obj.Mode != nil
}

func (obj *flowRouter) SetMode(value FlowRouterModeEnum) FlowRouter {
	intValue, ok := otg.FlowRouter_Mode_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRouterModeEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRouter_Mode_Enum(intValue)
	obj.obj.Mode = &enumValue

	return obj
}

// TBD
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv6/properties/name
// - /components/schemas/Bgp.V4RouteRange/properties/name
// - /components/schemas/Bgp.V6RouteRange/properties/name
// - /components/schemas/Bgp.CMacIpRange/properties/name
// - /components/schemas/Rsvp.LspIpv4Interface.P2PIngressIpv4Lsp/properties/name
// - /components/schemas/Isis.V4RouteRange/properties/name
// - /components/schemas/Isis.V6RouteRange/properties/name
// - /components/schemas/Device.Dhcpv4client/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv6/properties/name
// - /components/schemas/Bgp.V4RouteRange/properties/name
// - /components/schemas/Bgp.V6RouteRange/properties/name
// - /components/schemas/Bgp.CMacIpRange/properties/name
// - /components/schemas/Rsvp.LspIpv4Interface.P2PIngressIpv4Lsp/properties/name
// - /components/schemas/Isis.V4RouteRange/properties/name
// - /components/schemas/Isis.V6RouteRange/properties/name
// - /components/schemas/Device.Dhcpv4client/properties/name
//
// TxNames returns a []string
func (obj *flowRouter) TxNames() []string {
	if obj.obj.TxNames == nil {
		obj.obj.TxNames = make([]string, 0)
	}
	return obj.obj.TxNames
}

// TBD
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv6/properties/name
// - /components/schemas/Bgp.V4RouteRange/properties/name
// - /components/schemas/Bgp.V6RouteRange/properties/name
// - /components/schemas/Bgp.CMacIpRange/properties/name
// - /components/schemas/Rsvp.LspIpv4Interface.P2PIngressIpv4Lsp/properties/name
// - /components/schemas/Isis.V4RouteRange/properties/name
// - /components/schemas/Isis.V6RouteRange/properties/name
// - /components/schemas/Device.Dhcpv4client/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv6/properties/name
// - /components/schemas/Bgp.V4RouteRange/properties/name
// - /components/schemas/Bgp.V6RouteRange/properties/name
// - /components/schemas/Bgp.CMacIpRange/properties/name
// - /components/schemas/Rsvp.LspIpv4Interface.P2PIngressIpv4Lsp/properties/name
// - /components/schemas/Isis.V4RouteRange/properties/name
// - /components/schemas/Isis.V6RouteRange/properties/name
// - /components/schemas/Device.Dhcpv4client/properties/name
//
// SetTxNames sets the []string value in the FlowRouter object
func (obj *flowRouter) SetTxNames(value []string) FlowRouter {

	if obj.obj.TxNames == nil {
		obj.obj.TxNames = make([]string, 0)
	}
	obj.obj.TxNames = value

	return obj
}

// TBD
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv6/properties/name
// - /components/schemas/Bgp.V4RouteRange/properties/name
// - /components/schemas/Bgp.V6RouteRange/properties/name
// - /components/schemas/Bgp.CMacIpRange/properties/name
// - /components/schemas/Rsvp.LspIpv4Interface.P2PEgressIpv4Lsp/properties/name
// - /components/schemas/Isis.V4RouteRange/properties/name
// - /components/schemas/Isis.V6RouteRange/properties/name
// - /components/schemas/Device.Dhcpv4client/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv6/properties/name
// - /components/schemas/Bgp.V4RouteRange/properties/name
// - /components/schemas/Bgp.V6RouteRange/properties/name
// - /components/schemas/Bgp.CMacIpRange/properties/name
// - /components/schemas/Rsvp.LspIpv4Interface.P2PEgressIpv4Lsp/properties/name
// - /components/schemas/Isis.V4RouteRange/properties/name
// - /components/schemas/Isis.V6RouteRange/properties/name
// - /components/schemas/Device.Dhcpv4client/properties/name
//
// RxNames returns a []string
func (obj *flowRouter) RxNames() []string {
	if obj.obj.RxNames == nil {
		obj.obj.RxNames = make([]string, 0)
	}
	return obj.obj.RxNames
}

// TBD
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv6/properties/name
// - /components/schemas/Bgp.V4RouteRange/properties/name
// - /components/schemas/Bgp.V6RouteRange/properties/name
// - /components/schemas/Bgp.CMacIpRange/properties/name
// - /components/schemas/Rsvp.LspIpv4Interface.P2PEgressIpv4Lsp/properties/name
// - /components/schemas/Isis.V4RouteRange/properties/name
// - /components/schemas/Isis.V6RouteRange/properties/name
// - /components/schemas/Device.Dhcpv4client/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv6/properties/name
// - /components/schemas/Bgp.V4RouteRange/properties/name
// - /components/schemas/Bgp.V6RouteRange/properties/name
// - /components/schemas/Bgp.CMacIpRange/properties/name
// - /components/schemas/Rsvp.LspIpv4Interface.P2PEgressIpv4Lsp/properties/name
// - /components/schemas/Isis.V4RouteRange/properties/name
// - /components/schemas/Isis.V6RouteRange/properties/name
// - /components/schemas/Device.Dhcpv4client/properties/name
//
// SetRxNames sets the []string value in the FlowRouter object
func (obj *flowRouter) SetRxNames(value []string) FlowRouter {

	if obj.obj.RxNames == nil {
		obj.obj.RxNames = make([]string, 0)
	}
	obj.obj.RxNames = value

	return obj
}

func (obj *flowRouter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *flowRouter) setDefault() {
	if obj.obj.Mode == nil {
		obj.SetMode(FlowRouterMode.MESH)

	}

}

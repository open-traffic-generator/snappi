package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowHeader *****
type flowHeader struct {
	validation
	obj                 *otg.FlowHeader
	marshaller          marshalFlowHeader
	unMarshaller        unMarshalFlowHeader
	customHolder        FlowCustom
	ethernetHolder      FlowEthernet
	vlanHolder          FlowVlan
	vxlanHolder         FlowVxlan
	ipv4Holder          FlowIpv4
	ipv6Holder          FlowIpv6
	pfcpauseHolder      FlowPfcPause
	ethernetpauseHolder FlowEthernetPause
	tcpHolder           FlowTcp
	udpHolder           FlowUdp
	greHolder           FlowGre
	gtpv1Holder         FlowGtpv1
	gtpv2Holder         FlowGtpv2
	arpHolder           FlowArp
	icmpHolder          FlowIcmp
	icmpv6Holder        FlowIcmpv6
	pppHolder           FlowPpp
	igmpv1Holder        FlowIgmpv1
	mplsHolder          FlowMpls
	snmpv2CHolder       FlowSnmpv2C
}

func NewFlowHeader() FlowHeader {
	obj := flowHeader{obj: &otg.FlowHeader{}}
	obj.setDefault()
	return &obj
}

func (obj *flowHeader) msg() *otg.FlowHeader {
	return obj.obj
}

func (obj *flowHeader) setMsg(msg *otg.FlowHeader) FlowHeader {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowHeader struct {
	obj *flowHeader
}

type marshalFlowHeader interface {
	// ToProto marshals FlowHeader to protobuf object *otg.FlowHeader
	ToProto() (*otg.FlowHeader, error)
	// ToPbText marshals FlowHeader to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowHeader to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowHeader to JSON text
	ToJson() (string, error)
}

type unMarshalflowHeader struct {
	obj *flowHeader
}

type unMarshalFlowHeader interface {
	// FromProto unmarshals FlowHeader from protobuf object *otg.FlowHeader
	FromProto(msg *otg.FlowHeader) (FlowHeader, error)
	// FromPbText unmarshals FlowHeader from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowHeader from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowHeader from JSON text
	FromJson(value string) error
}

func (obj *flowHeader) Marshal() marshalFlowHeader {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowHeader{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowHeader) Unmarshal() unMarshalFlowHeader {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowHeader{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowHeader) ToProto() (*otg.FlowHeader, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowHeader) FromProto(msg *otg.FlowHeader) (FlowHeader, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowHeader) ToPbText() (string, error) {
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

func (m *unMarshalflowHeader) FromPbText(value string) error {
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

func (m *marshalflowHeader) ToYaml() (string, error) {
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

func (m *unMarshalflowHeader) FromYaml(value string) error {
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

func (m *marshalflowHeader) ToJson() (string, error) {
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

func (m *unMarshalflowHeader) FromJson(value string) error {
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

func (obj *flowHeader) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowHeader) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowHeader) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowHeader) Clone() (FlowHeader, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowHeader()
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

func (obj *flowHeader) setNil() {
	obj.customHolder = nil
	obj.ethernetHolder = nil
	obj.vlanHolder = nil
	obj.vxlanHolder = nil
	obj.ipv4Holder = nil
	obj.ipv6Holder = nil
	obj.pfcpauseHolder = nil
	obj.ethernetpauseHolder = nil
	obj.tcpHolder = nil
	obj.udpHolder = nil
	obj.greHolder = nil
	obj.gtpv1Holder = nil
	obj.gtpv2Holder = nil
	obj.arpHolder = nil
	obj.icmpHolder = nil
	obj.icmpv6Holder = nil
	obj.pppHolder = nil
	obj.igmpv1Holder = nil
	obj.mplsHolder = nil
	obj.snmpv2CHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowHeader is configuration for all traffic packet headers
type FlowHeader interface {
	Validation
	// msg marshals FlowHeader to protobuf object *otg.FlowHeader
	// and doesn't set defaults
	msg() *otg.FlowHeader
	// setMsg unmarshals FlowHeader from protobuf object *otg.FlowHeader
	// and doesn't set defaults
	setMsg(*otg.FlowHeader) FlowHeader
	// provides marshal interface
	Marshal() marshalFlowHeader
	// provides unmarshal interface
	Unmarshal() unMarshalFlowHeader
	// validate validates FlowHeader
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowHeader, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowHeaderChoiceEnum, set in FlowHeader
	Choice() FlowHeaderChoiceEnum
	// setChoice assigns FlowHeaderChoiceEnum provided by user to FlowHeader
	setChoice(value FlowHeaderChoiceEnum) FlowHeader
	// HasChoice checks if Choice has been set in FlowHeader
	HasChoice() bool
	// Custom returns FlowCustom, set in FlowHeader.
	// FlowCustom is custom packet header
	Custom() FlowCustom
	// SetCustom assigns FlowCustom provided by user to FlowHeader.
	// FlowCustom is custom packet header
	SetCustom(value FlowCustom) FlowHeader
	// HasCustom checks if Custom has been set in FlowHeader
	HasCustom() bool
	// Ethernet returns FlowEthernet, set in FlowHeader.
	// FlowEthernet is ethernet packet header
	Ethernet() FlowEthernet
	// SetEthernet assigns FlowEthernet provided by user to FlowHeader.
	// FlowEthernet is ethernet packet header
	SetEthernet(value FlowEthernet) FlowHeader
	// HasEthernet checks if Ethernet has been set in FlowHeader
	HasEthernet() bool
	// Vlan returns FlowVlan, set in FlowHeader.
	// FlowVlan is vLAN packet header
	Vlan() FlowVlan
	// SetVlan assigns FlowVlan provided by user to FlowHeader.
	// FlowVlan is vLAN packet header
	SetVlan(value FlowVlan) FlowHeader
	// HasVlan checks if Vlan has been set in FlowHeader
	HasVlan() bool
	// Vxlan returns FlowVxlan, set in FlowHeader.
	// FlowVxlan is vXLAN packet header
	Vxlan() FlowVxlan
	// SetVxlan assigns FlowVxlan provided by user to FlowHeader.
	// FlowVxlan is vXLAN packet header
	SetVxlan(value FlowVxlan) FlowHeader
	// HasVxlan checks if Vxlan has been set in FlowHeader
	HasVxlan() bool
	// Ipv4 returns FlowIpv4, set in FlowHeader.
	// FlowIpv4 is iPv4 packet header
	Ipv4() FlowIpv4
	// SetIpv4 assigns FlowIpv4 provided by user to FlowHeader.
	// FlowIpv4 is iPv4 packet header
	SetIpv4(value FlowIpv4) FlowHeader
	// HasIpv4 checks if Ipv4 has been set in FlowHeader
	HasIpv4() bool
	// Ipv6 returns FlowIpv6, set in FlowHeader.
	// FlowIpv6 is iPv6 packet header
	Ipv6() FlowIpv6
	// SetIpv6 assigns FlowIpv6 provided by user to FlowHeader.
	// FlowIpv6 is iPv6 packet header
	SetIpv6(value FlowIpv6) FlowHeader
	// HasIpv6 checks if Ipv6 has been set in FlowHeader
	HasIpv6() bool
	// Pfcpause returns FlowPfcPause, set in FlowHeader.
	// FlowPfcPause is iEEE 802.1Qbb PFC Pause packet header.
	Pfcpause() FlowPfcPause
	// SetPfcpause assigns FlowPfcPause provided by user to FlowHeader.
	// FlowPfcPause is iEEE 802.1Qbb PFC Pause packet header.
	SetPfcpause(value FlowPfcPause) FlowHeader
	// HasPfcpause checks if Pfcpause has been set in FlowHeader
	HasPfcpause() bool
	// Ethernetpause returns FlowEthernetPause, set in FlowHeader.
	// FlowEthernetPause is iEEE 802.3x global ethernet pause packet header
	Ethernetpause() FlowEthernetPause
	// SetEthernetpause assigns FlowEthernetPause provided by user to FlowHeader.
	// FlowEthernetPause is iEEE 802.3x global ethernet pause packet header
	SetEthernetpause(value FlowEthernetPause) FlowHeader
	// HasEthernetpause checks if Ethernetpause has been set in FlowHeader
	HasEthernetpause() bool
	// Tcp returns FlowTcp, set in FlowHeader.
	// FlowTcp is tCP packet header
	Tcp() FlowTcp
	// SetTcp assigns FlowTcp provided by user to FlowHeader.
	// FlowTcp is tCP packet header
	SetTcp(value FlowTcp) FlowHeader
	// HasTcp checks if Tcp has been set in FlowHeader
	HasTcp() bool
	// Udp returns FlowUdp, set in FlowHeader.
	// FlowUdp is uDP packet header
	Udp() FlowUdp
	// SetUdp assigns FlowUdp provided by user to FlowHeader.
	// FlowUdp is uDP packet header
	SetUdp(value FlowUdp) FlowHeader
	// HasUdp checks if Udp has been set in FlowHeader
	HasUdp() bool
	// Gre returns FlowGre, set in FlowHeader.
	// FlowGre is standard GRE packet header (RFC2784)
	Gre() FlowGre
	// SetGre assigns FlowGre provided by user to FlowHeader.
	// FlowGre is standard GRE packet header (RFC2784)
	SetGre(value FlowGre) FlowHeader
	// HasGre checks if Gre has been set in FlowHeader
	HasGre() bool
	// Gtpv1 returns FlowGtpv1, set in FlowHeader.
	// FlowGtpv1 is gTPv1 packet header
	Gtpv1() FlowGtpv1
	// SetGtpv1 assigns FlowGtpv1 provided by user to FlowHeader.
	// FlowGtpv1 is gTPv1 packet header
	SetGtpv1(value FlowGtpv1) FlowHeader
	// HasGtpv1 checks if Gtpv1 has been set in FlowHeader
	HasGtpv1() bool
	// Gtpv2 returns FlowGtpv2, set in FlowHeader.
	// FlowGtpv2 is gTPv2 packet header
	Gtpv2() FlowGtpv2
	// SetGtpv2 assigns FlowGtpv2 provided by user to FlowHeader.
	// FlowGtpv2 is gTPv2 packet header
	SetGtpv2(value FlowGtpv2) FlowHeader
	// HasGtpv2 checks if Gtpv2 has been set in FlowHeader
	HasGtpv2() bool
	// Arp returns FlowArp, set in FlowHeader.
	// FlowArp is aRP packet header
	Arp() FlowArp
	// SetArp assigns FlowArp provided by user to FlowHeader.
	// FlowArp is aRP packet header
	SetArp(value FlowArp) FlowHeader
	// HasArp checks if Arp has been set in FlowHeader
	HasArp() bool
	// Icmp returns FlowIcmp, set in FlowHeader.
	// FlowIcmp is iCMP packet header
	Icmp() FlowIcmp
	// SetIcmp assigns FlowIcmp provided by user to FlowHeader.
	// FlowIcmp is iCMP packet header
	SetIcmp(value FlowIcmp) FlowHeader
	// HasIcmp checks if Icmp has been set in FlowHeader
	HasIcmp() bool
	// Icmpv6 returns FlowIcmpv6, set in FlowHeader.
	// FlowIcmpv6 is iCMPv6 packet header
	Icmpv6() FlowIcmpv6
	// SetIcmpv6 assigns FlowIcmpv6 provided by user to FlowHeader.
	// FlowIcmpv6 is iCMPv6 packet header
	SetIcmpv6(value FlowIcmpv6) FlowHeader
	// HasIcmpv6 checks if Icmpv6 has been set in FlowHeader
	HasIcmpv6() bool
	// Ppp returns FlowPpp, set in FlowHeader.
	// FlowPpp is pPP packet header
	Ppp() FlowPpp
	// SetPpp assigns FlowPpp provided by user to FlowHeader.
	// FlowPpp is pPP packet header
	SetPpp(value FlowPpp) FlowHeader
	// HasPpp checks if Ppp has been set in FlowHeader
	HasPpp() bool
	// Igmpv1 returns FlowIgmpv1, set in FlowHeader.
	// FlowIgmpv1 is iGMPv1 packet header
	Igmpv1() FlowIgmpv1
	// SetIgmpv1 assigns FlowIgmpv1 provided by user to FlowHeader.
	// FlowIgmpv1 is iGMPv1 packet header
	SetIgmpv1(value FlowIgmpv1) FlowHeader
	// HasIgmpv1 checks if Igmpv1 has been set in FlowHeader
	HasIgmpv1() bool
	// Mpls returns FlowMpls, set in FlowHeader.
	// FlowMpls is mPLS packet header; When configuring multiple such headers, the count shall not exceed 20.
	Mpls() FlowMpls
	// SetMpls assigns FlowMpls provided by user to FlowHeader.
	// FlowMpls is mPLS packet header; When configuring multiple such headers, the count shall not exceed 20.
	SetMpls(value FlowMpls) FlowHeader
	// HasMpls checks if Mpls has been set in FlowHeader
	HasMpls() bool
	// Snmpv2C returns FlowSnmpv2C, set in FlowHeader.
	Snmpv2C() FlowSnmpv2C
	// SetSnmpv2C assigns FlowSnmpv2C provided by user to FlowHeader.
	SetSnmpv2C(value FlowSnmpv2C) FlowHeader
	// HasSnmpv2C checks if Snmpv2C has been set in FlowHeader
	HasSnmpv2C() bool
	setNil()
}

type FlowHeaderChoiceEnum string

// Enum of Choice on FlowHeader
var FlowHeaderChoice = struct {
	CUSTOM        FlowHeaderChoiceEnum
	ETHERNET      FlowHeaderChoiceEnum
	VLAN          FlowHeaderChoiceEnum
	VXLAN         FlowHeaderChoiceEnum
	IPV4          FlowHeaderChoiceEnum
	IPV6          FlowHeaderChoiceEnum
	PFCPAUSE      FlowHeaderChoiceEnum
	ETHERNETPAUSE FlowHeaderChoiceEnum
	TCP           FlowHeaderChoiceEnum
	UDP           FlowHeaderChoiceEnum
	GRE           FlowHeaderChoiceEnum
	GTPV1         FlowHeaderChoiceEnum
	GTPV2         FlowHeaderChoiceEnum
	ARP           FlowHeaderChoiceEnum
	ICMP          FlowHeaderChoiceEnum
	ICMPV6        FlowHeaderChoiceEnum
	PPP           FlowHeaderChoiceEnum
	IGMPV1        FlowHeaderChoiceEnum
	MPLS          FlowHeaderChoiceEnum
	SNMPV2C       FlowHeaderChoiceEnum
}{
	CUSTOM:        FlowHeaderChoiceEnum("custom"),
	ETHERNET:      FlowHeaderChoiceEnum("ethernet"),
	VLAN:          FlowHeaderChoiceEnum("vlan"),
	VXLAN:         FlowHeaderChoiceEnum("vxlan"),
	IPV4:          FlowHeaderChoiceEnum("ipv4"),
	IPV6:          FlowHeaderChoiceEnum("ipv6"),
	PFCPAUSE:      FlowHeaderChoiceEnum("pfcpause"),
	ETHERNETPAUSE: FlowHeaderChoiceEnum("ethernetpause"),
	TCP:           FlowHeaderChoiceEnum("tcp"),
	UDP:           FlowHeaderChoiceEnum("udp"),
	GRE:           FlowHeaderChoiceEnum("gre"),
	GTPV1:         FlowHeaderChoiceEnum("gtpv1"),
	GTPV2:         FlowHeaderChoiceEnum("gtpv2"),
	ARP:           FlowHeaderChoiceEnum("arp"),
	ICMP:          FlowHeaderChoiceEnum("icmp"),
	ICMPV6:        FlowHeaderChoiceEnum("icmpv6"),
	PPP:           FlowHeaderChoiceEnum("ppp"),
	IGMPV1:        FlowHeaderChoiceEnum("igmpv1"),
	MPLS:          FlowHeaderChoiceEnum("mpls"),
	SNMPV2C:       FlowHeaderChoiceEnum("snmpv2c"),
}

func (obj *flowHeader) Choice() FlowHeaderChoiceEnum {
	return FlowHeaderChoiceEnum(obj.obj.Choice.Enum().String())
}

// The available types of flow headers. If one is not provided the
// default ethernet packet header MUST be provided.
// Choice returns a string
func (obj *flowHeader) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowHeader) setChoice(value FlowHeaderChoiceEnum) FlowHeader {
	intValue, ok := otg.FlowHeader_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowHeaderChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowHeader_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Snmpv2C = nil
	obj.snmpv2CHolder = nil
	obj.obj.Mpls = nil
	obj.mplsHolder = nil
	obj.obj.Igmpv1 = nil
	obj.igmpv1Holder = nil
	obj.obj.Ppp = nil
	obj.pppHolder = nil
	obj.obj.Icmpv6 = nil
	obj.icmpv6Holder = nil
	obj.obj.Icmp = nil
	obj.icmpHolder = nil
	obj.obj.Arp = nil
	obj.arpHolder = nil
	obj.obj.Gtpv2 = nil
	obj.gtpv2Holder = nil
	obj.obj.Gtpv1 = nil
	obj.gtpv1Holder = nil
	obj.obj.Gre = nil
	obj.greHolder = nil
	obj.obj.Udp = nil
	obj.udpHolder = nil
	obj.obj.Tcp = nil
	obj.tcpHolder = nil
	obj.obj.Ethernetpause = nil
	obj.ethernetpauseHolder = nil
	obj.obj.Pfcpause = nil
	obj.pfcpauseHolder = nil
	obj.obj.Ipv6 = nil
	obj.ipv6Holder = nil
	obj.obj.Ipv4 = nil
	obj.ipv4Holder = nil
	obj.obj.Vxlan = nil
	obj.vxlanHolder = nil
	obj.obj.Vlan = nil
	obj.vlanHolder = nil
	obj.obj.Ethernet = nil
	obj.ethernetHolder = nil
	obj.obj.Custom = nil
	obj.customHolder = nil

	if value == FlowHeaderChoice.CUSTOM {
		obj.obj.Custom = NewFlowCustom().msg()
	}

	if value == FlowHeaderChoice.ETHERNET {
		obj.obj.Ethernet = NewFlowEthernet().msg()
	}

	if value == FlowHeaderChoice.VLAN {
		obj.obj.Vlan = NewFlowVlan().msg()
	}

	if value == FlowHeaderChoice.VXLAN {
		obj.obj.Vxlan = NewFlowVxlan().msg()
	}

	if value == FlowHeaderChoice.IPV4 {
		obj.obj.Ipv4 = NewFlowIpv4().msg()
	}

	if value == FlowHeaderChoice.IPV6 {
		obj.obj.Ipv6 = NewFlowIpv6().msg()
	}

	if value == FlowHeaderChoice.PFCPAUSE {
		obj.obj.Pfcpause = NewFlowPfcPause().msg()
	}

	if value == FlowHeaderChoice.ETHERNETPAUSE {
		obj.obj.Ethernetpause = NewFlowEthernetPause().msg()
	}

	if value == FlowHeaderChoice.TCP {
		obj.obj.Tcp = NewFlowTcp().msg()
	}

	if value == FlowHeaderChoice.UDP {
		obj.obj.Udp = NewFlowUdp().msg()
	}

	if value == FlowHeaderChoice.GRE {
		obj.obj.Gre = NewFlowGre().msg()
	}

	if value == FlowHeaderChoice.GTPV1 {
		obj.obj.Gtpv1 = NewFlowGtpv1().msg()
	}

	if value == FlowHeaderChoice.GTPV2 {
		obj.obj.Gtpv2 = NewFlowGtpv2().msg()
	}

	if value == FlowHeaderChoice.ARP {
		obj.obj.Arp = NewFlowArp().msg()
	}

	if value == FlowHeaderChoice.ICMP {
		obj.obj.Icmp = NewFlowIcmp().msg()
	}

	if value == FlowHeaderChoice.ICMPV6 {
		obj.obj.Icmpv6 = NewFlowIcmpv6().msg()
	}

	if value == FlowHeaderChoice.PPP {
		obj.obj.Ppp = NewFlowPpp().msg()
	}

	if value == FlowHeaderChoice.IGMPV1 {
		obj.obj.Igmpv1 = NewFlowIgmpv1().msg()
	}

	if value == FlowHeaderChoice.MPLS {
		obj.obj.Mpls = NewFlowMpls().msg()
	}

	if value == FlowHeaderChoice.SNMPV2C {
		obj.obj.Snmpv2C = NewFlowSnmpv2C().msg()
	}

	return obj
}

// description is TBD
// Custom returns a FlowCustom
func (obj *flowHeader) Custom() FlowCustom {
	if obj.obj.Custom == nil {
		obj.setChoice(FlowHeaderChoice.CUSTOM)
	}
	if obj.customHolder == nil {
		obj.customHolder = &flowCustom{obj: obj.obj.Custom}
	}
	return obj.customHolder
}

// description is TBD
// Custom returns a FlowCustom
func (obj *flowHeader) HasCustom() bool {
	return obj.obj.Custom != nil
}

// description is TBD
// SetCustom sets the FlowCustom value in the FlowHeader object
func (obj *flowHeader) SetCustom(value FlowCustom) FlowHeader {
	obj.setChoice(FlowHeaderChoice.CUSTOM)
	obj.customHolder = nil
	obj.obj.Custom = value.msg()

	return obj
}

// description is TBD
// Ethernet returns a FlowEthernet
func (obj *flowHeader) Ethernet() FlowEthernet {
	if obj.obj.Ethernet == nil {
		obj.setChoice(FlowHeaderChoice.ETHERNET)
	}
	if obj.ethernetHolder == nil {
		obj.ethernetHolder = &flowEthernet{obj: obj.obj.Ethernet}
	}
	return obj.ethernetHolder
}

// description is TBD
// Ethernet returns a FlowEthernet
func (obj *flowHeader) HasEthernet() bool {
	return obj.obj.Ethernet != nil
}

// description is TBD
// SetEthernet sets the FlowEthernet value in the FlowHeader object
func (obj *flowHeader) SetEthernet(value FlowEthernet) FlowHeader {
	obj.setChoice(FlowHeaderChoice.ETHERNET)
	obj.ethernetHolder = nil
	obj.obj.Ethernet = value.msg()

	return obj
}

// description is TBD
// Vlan returns a FlowVlan
func (obj *flowHeader) Vlan() FlowVlan {
	if obj.obj.Vlan == nil {
		obj.setChoice(FlowHeaderChoice.VLAN)
	}
	if obj.vlanHolder == nil {
		obj.vlanHolder = &flowVlan{obj: obj.obj.Vlan}
	}
	return obj.vlanHolder
}

// description is TBD
// Vlan returns a FlowVlan
func (obj *flowHeader) HasVlan() bool {
	return obj.obj.Vlan != nil
}

// description is TBD
// SetVlan sets the FlowVlan value in the FlowHeader object
func (obj *flowHeader) SetVlan(value FlowVlan) FlowHeader {
	obj.setChoice(FlowHeaderChoice.VLAN)
	obj.vlanHolder = nil
	obj.obj.Vlan = value.msg()

	return obj
}

// description is TBD
// Vxlan returns a FlowVxlan
func (obj *flowHeader) Vxlan() FlowVxlan {
	if obj.obj.Vxlan == nil {
		obj.setChoice(FlowHeaderChoice.VXLAN)
	}
	if obj.vxlanHolder == nil {
		obj.vxlanHolder = &flowVxlan{obj: obj.obj.Vxlan}
	}
	return obj.vxlanHolder
}

// description is TBD
// Vxlan returns a FlowVxlan
func (obj *flowHeader) HasVxlan() bool {
	return obj.obj.Vxlan != nil
}

// description is TBD
// SetVxlan sets the FlowVxlan value in the FlowHeader object
func (obj *flowHeader) SetVxlan(value FlowVxlan) FlowHeader {
	obj.setChoice(FlowHeaderChoice.VXLAN)
	obj.vxlanHolder = nil
	obj.obj.Vxlan = value.msg()

	return obj
}

// description is TBD
// Ipv4 returns a FlowIpv4
func (obj *flowHeader) Ipv4() FlowIpv4 {
	if obj.obj.Ipv4 == nil {
		obj.setChoice(FlowHeaderChoice.IPV4)
	}
	if obj.ipv4Holder == nil {
		obj.ipv4Holder = &flowIpv4{obj: obj.obj.Ipv4}
	}
	return obj.ipv4Holder
}

// description is TBD
// Ipv4 returns a FlowIpv4
func (obj *flowHeader) HasIpv4() bool {
	return obj.obj.Ipv4 != nil
}

// description is TBD
// SetIpv4 sets the FlowIpv4 value in the FlowHeader object
func (obj *flowHeader) SetIpv4(value FlowIpv4) FlowHeader {
	obj.setChoice(FlowHeaderChoice.IPV4)
	obj.ipv4Holder = nil
	obj.obj.Ipv4 = value.msg()

	return obj
}

// description is TBD
// Ipv6 returns a FlowIpv6
func (obj *flowHeader) Ipv6() FlowIpv6 {
	if obj.obj.Ipv6 == nil {
		obj.setChoice(FlowHeaderChoice.IPV6)
	}
	if obj.ipv6Holder == nil {
		obj.ipv6Holder = &flowIpv6{obj: obj.obj.Ipv6}
	}
	return obj.ipv6Holder
}

// description is TBD
// Ipv6 returns a FlowIpv6
func (obj *flowHeader) HasIpv6() bool {
	return obj.obj.Ipv6 != nil
}

// description is TBD
// SetIpv6 sets the FlowIpv6 value in the FlowHeader object
func (obj *flowHeader) SetIpv6(value FlowIpv6) FlowHeader {
	obj.setChoice(FlowHeaderChoice.IPV6)
	obj.ipv6Holder = nil
	obj.obj.Ipv6 = value.msg()

	return obj
}

// description is TBD
// Pfcpause returns a FlowPfcPause
func (obj *flowHeader) Pfcpause() FlowPfcPause {
	if obj.obj.Pfcpause == nil {
		obj.setChoice(FlowHeaderChoice.PFCPAUSE)
	}
	if obj.pfcpauseHolder == nil {
		obj.pfcpauseHolder = &flowPfcPause{obj: obj.obj.Pfcpause}
	}
	return obj.pfcpauseHolder
}

// description is TBD
// Pfcpause returns a FlowPfcPause
func (obj *flowHeader) HasPfcpause() bool {
	return obj.obj.Pfcpause != nil
}

// description is TBD
// SetPfcpause sets the FlowPfcPause value in the FlowHeader object
func (obj *flowHeader) SetPfcpause(value FlowPfcPause) FlowHeader {
	obj.setChoice(FlowHeaderChoice.PFCPAUSE)
	obj.pfcpauseHolder = nil
	obj.obj.Pfcpause = value.msg()

	return obj
}

// description is TBD
// Ethernetpause returns a FlowEthernetPause
func (obj *flowHeader) Ethernetpause() FlowEthernetPause {
	if obj.obj.Ethernetpause == nil {
		obj.setChoice(FlowHeaderChoice.ETHERNETPAUSE)
	}
	if obj.ethernetpauseHolder == nil {
		obj.ethernetpauseHolder = &flowEthernetPause{obj: obj.obj.Ethernetpause}
	}
	return obj.ethernetpauseHolder
}

// description is TBD
// Ethernetpause returns a FlowEthernetPause
func (obj *flowHeader) HasEthernetpause() bool {
	return obj.obj.Ethernetpause != nil
}

// description is TBD
// SetEthernetpause sets the FlowEthernetPause value in the FlowHeader object
func (obj *flowHeader) SetEthernetpause(value FlowEthernetPause) FlowHeader {
	obj.setChoice(FlowHeaderChoice.ETHERNETPAUSE)
	obj.ethernetpauseHolder = nil
	obj.obj.Ethernetpause = value.msg()

	return obj
}

// description is TBD
// Tcp returns a FlowTcp
func (obj *flowHeader) Tcp() FlowTcp {
	if obj.obj.Tcp == nil {
		obj.setChoice(FlowHeaderChoice.TCP)
	}
	if obj.tcpHolder == nil {
		obj.tcpHolder = &flowTcp{obj: obj.obj.Tcp}
	}
	return obj.tcpHolder
}

// description is TBD
// Tcp returns a FlowTcp
func (obj *flowHeader) HasTcp() bool {
	return obj.obj.Tcp != nil
}

// description is TBD
// SetTcp sets the FlowTcp value in the FlowHeader object
func (obj *flowHeader) SetTcp(value FlowTcp) FlowHeader {
	obj.setChoice(FlowHeaderChoice.TCP)
	obj.tcpHolder = nil
	obj.obj.Tcp = value.msg()

	return obj
}

// description is TBD
// Udp returns a FlowUdp
func (obj *flowHeader) Udp() FlowUdp {
	if obj.obj.Udp == nil {
		obj.setChoice(FlowHeaderChoice.UDP)
	}
	if obj.udpHolder == nil {
		obj.udpHolder = &flowUdp{obj: obj.obj.Udp}
	}
	return obj.udpHolder
}

// description is TBD
// Udp returns a FlowUdp
func (obj *flowHeader) HasUdp() bool {
	return obj.obj.Udp != nil
}

// description is TBD
// SetUdp sets the FlowUdp value in the FlowHeader object
func (obj *flowHeader) SetUdp(value FlowUdp) FlowHeader {
	obj.setChoice(FlowHeaderChoice.UDP)
	obj.udpHolder = nil
	obj.obj.Udp = value.msg()

	return obj
}

// description is TBD
// Gre returns a FlowGre
func (obj *flowHeader) Gre() FlowGre {
	if obj.obj.Gre == nil {
		obj.setChoice(FlowHeaderChoice.GRE)
	}
	if obj.greHolder == nil {
		obj.greHolder = &flowGre{obj: obj.obj.Gre}
	}
	return obj.greHolder
}

// description is TBD
// Gre returns a FlowGre
func (obj *flowHeader) HasGre() bool {
	return obj.obj.Gre != nil
}

// description is TBD
// SetGre sets the FlowGre value in the FlowHeader object
func (obj *flowHeader) SetGre(value FlowGre) FlowHeader {
	obj.setChoice(FlowHeaderChoice.GRE)
	obj.greHolder = nil
	obj.obj.Gre = value.msg()

	return obj
}

// description is TBD
// Gtpv1 returns a FlowGtpv1
func (obj *flowHeader) Gtpv1() FlowGtpv1 {
	if obj.obj.Gtpv1 == nil {
		obj.setChoice(FlowHeaderChoice.GTPV1)
	}
	if obj.gtpv1Holder == nil {
		obj.gtpv1Holder = &flowGtpv1{obj: obj.obj.Gtpv1}
	}
	return obj.gtpv1Holder
}

// description is TBD
// Gtpv1 returns a FlowGtpv1
func (obj *flowHeader) HasGtpv1() bool {
	return obj.obj.Gtpv1 != nil
}

// description is TBD
// SetGtpv1 sets the FlowGtpv1 value in the FlowHeader object
func (obj *flowHeader) SetGtpv1(value FlowGtpv1) FlowHeader {
	obj.setChoice(FlowHeaderChoice.GTPV1)
	obj.gtpv1Holder = nil
	obj.obj.Gtpv1 = value.msg()

	return obj
}

// description is TBD
// Gtpv2 returns a FlowGtpv2
func (obj *flowHeader) Gtpv2() FlowGtpv2 {
	if obj.obj.Gtpv2 == nil {
		obj.setChoice(FlowHeaderChoice.GTPV2)
	}
	if obj.gtpv2Holder == nil {
		obj.gtpv2Holder = &flowGtpv2{obj: obj.obj.Gtpv2}
	}
	return obj.gtpv2Holder
}

// description is TBD
// Gtpv2 returns a FlowGtpv2
func (obj *flowHeader) HasGtpv2() bool {
	return obj.obj.Gtpv2 != nil
}

// description is TBD
// SetGtpv2 sets the FlowGtpv2 value in the FlowHeader object
func (obj *flowHeader) SetGtpv2(value FlowGtpv2) FlowHeader {
	obj.setChoice(FlowHeaderChoice.GTPV2)
	obj.gtpv2Holder = nil
	obj.obj.Gtpv2 = value.msg()

	return obj
}

// description is TBD
// Arp returns a FlowArp
func (obj *flowHeader) Arp() FlowArp {
	if obj.obj.Arp == nil {
		obj.setChoice(FlowHeaderChoice.ARP)
	}
	if obj.arpHolder == nil {
		obj.arpHolder = &flowArp{obj: obj.obj.Arp}
	}
	return obj.arpHolder
}

// description is TBD
// Arp returns a FlowArp
func (obj *flowHeader) HasArp() bool {
	return obj.obj.Arp != nil
}

// description is TBD
// SetArp sets the FlowArp value in the FlowHeader object
func (obj *flowHeader) SetArp(value FlowArp) FlowHeader {
	obj.setChoice(FlowHeaderChoice.ARP)
	obj.arpHolder = nil
	obj.obj.Arp = value.msg()

	return obj
}

// description is TBD
// Icmp returns a FlowIcmp
func (obj *flowHeader) Icmp() FlowIcmp {
	if obj.obj.Icmp == nil {
		obj.setChoice(FlowHeaderChoice.ICMP)
	}
	if obj.icmpHolder == nil {
		obj.icmpHolder = &flowIcmp{obj: obj.obj.Icmp}
	}
	return obj.icmpHolder
}

// description is TBD
// Icmp returns a FlowIcmp
func (obj *flowHeader) HasIcmp() bool {
	return obj.obj.Icmp != nil
}

// description is TBD
// SetIcmp sets the FlowIcmp value in the FlowHeader object
func (obj *flowHeader) SetIcmp(value FlowIcmp) FlowHeader {
	obj.setChoice(FlowHeaderChoice.ICMP)
	obj.icmpHolder = nil
	obj.obj.Icmp = value.msg()

	return obj
}

// description is TBD
// Icmpv6 returns a FlowIcmpv6
func (obj *flowHeader) Icmpv6() FlowIcmpv6 {
	if obj.obj.Icmpv6 == nil {
		obj.setChoice(FlowHeaderChoice.ICMPV6)
	}
	if obj.icmpv6Holder == nil {
		obj.icmpv6Holder = &flowIcmpv6{obj: obj.obj.Icmpv6}
	}
	return obj.icmpv6Holder
}

// description is TBD
// Icmpv6 returns a FlowIcmpv6
func (obj *flowHeader) HasIcmpv6() bool {
	return obj.obj.Icmpv6 != nil
}

// description is TBD
// SetIcmpv6 sets the FlowIcmpv6 value in the FlowHeader object
func (obj *flowHeader) SetIcmpv6(value FlowIcmpv6) FlowHeader {
	obj.setChoice(FlowHeaderChoice.ICMPV6)
	obj.icmpv6Holder = nil
	obj.obj.Icmpv6 = value.msg()

	return obj
}

// description is TBD
// Ppp returns a FlowPpp
func (obj *flowHeader) Ppp() FlowPpp {
	if obj.obj.Ppp == nil {
		obj.setChoice(FlowHeaderChoice.PPP)
	}
	if obj.pppHolder == nil {
		obj.pppHolder = &flowPpp{obj: obj.obj.Ppp}
	}
	return obj.pppHolder
}

// description is TBD
// Ppp returns a FlowPpp
func (obj *flowHeader) HasPpp() bool {
	return obj.obj.Ppp != nil
}

// description is TBD
// SetPpp sets the FlowPpp value in the FlowHeader object
func (obj *flowHeader) SetPpp(value FlowPpp) FlowHeader {
	obj.setChoice(FlowHeaderChoice.PPP)
	obj.pppHolder = nil
	obj.obj.Ppp = value.msg()

	return obj
}

// description is TBD
// Igmpv1 returns a FlowIgmpv1
func (obj *flowHeader) Igmpv1() FlowIgmpv1 {
	if obj.obj.Igmpv1 == nil {
		obj.setChoice(FlowHeaderChoice.IGMPV1)
	}
	if obj.igmpv1Holder == nil {
		obj.igmpv1Holder = &flowIgmpv1{obj: obj.obj.Igmpv1}
	}
	return obj.igmpv1Holder
}

// description is TBD
// Igmpv1 returns a FlowIgmpv1
func (obj *flowHeader) HasIgmpv1() bool {
	return obj.obj.Igmpv1 != nil
}

// description is TBD
// SetIgmpv1 sets the FlowIgmpv1 value in the FlowHeader object
func (obj *flowHeader) SetIgmpv1(value FlowIgmpv1) FlowHeader {
	obj.setChoice(FlowHeaderChoice.IGMPV1)
	obj.igmpv1Holder = nil
	obj.obj.Igmpv1 = value.msg()

	return obj
}

// description is TBD
// Mpls returns a FlowMpls
func (obj *flowHeader) Mpls() FlowMpls {
	if obj.obj.Mpls == nil {
		obj.setChoice(FlowHeaderChoice.MPLS)
	}
	if obj.mplsHolder == nil {
		obj.mplsHolder = &flowMpls{obj: obj.obj.Mpls}
	}
	return obj.mplsHolder
}

// description is TBD
// Mpls returns a FlowMpls
func (obj *flowHeader) HasMpls() bool {
	return obj.obj.Mpls != nil
}

// description is TBD
// SetMpls sets the FlowMpls value in the FlowHeader object
func (obj *flowHeader) SetMpls(value FlowMpls) FlowHeader {
	obj.setChoice(FlowHeaderChoice.MPLS)
	obj.mplsHolder = nil
	obj.obj.Mpls = value.msg()

	return obj
}

// description is TBD
// Snmpv2C returns a FlowSnmpv2C
func (obj *flowHeader) Snmpv2C() FlowSnmpv2C {
	if obj.obj.Snmpv2C == nil {
		obj.setChoice(FlowHeaderChoice.SNMPV2C)
	}
	if obj.snmpv2CHolder == nil {
		obj.snmpv2CHolder = &flowSnmpv2C{obj: obj.obj.Snmpv2C}
	}
	return obj.snmpv2CHolder
}

// description is TBD
// Snmpv2C returns a FlowSnmpv2C
func (obj *flowHeader) HasSnmpv2C() bool {
	return obj.obj.Snmpv2C != nil
}

// description is TBD
// SetSnmpv2C sets the FlowSnmpv2C value in the FlowHeader object
func (obj *flowHeader) SetSnmpv2C(value FlowSnmpv2C) FlowHeader {
	obj.setChoice(FlowHeaderChoice.SNMPV2C)
	obj.snmpv2CHolder = nil
	obj.obj.Snmpv2C = value.msg()

	return obj
}

func (obj *flowHeader) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		obj.Custom().validateObj(vObj, set_default)
	}

	if obj.obj.Ethernet != nil {

		obj.Ethernet().validateObj(vObj, set_default)
	}

	if obj.obj.Vlan != nil {

		obj.Vlan().validateObj(vObj, set_default)
	}

	if obj.obj.Vxlan != nil {

		obj.Vxlan().validateObj(vObj, set_default)
	}

	if obj.obj.Ipv4 != nil {

		obj.Ipv4().validateObj(vObj, set_default)
	}

	if obj.obj.Ipv6 != nil {

		obj.Ipv6().validateObj(vObj, set_default)
	}

	if obj.obj.Pfcpause != nil {

		obj.Pfcpause().validateObj(vObj, set_default)
	}

	if obj.obj.Ethernetpause != nil {

		obj.Ethernetpause().validateObj(vObj, set_default)
	}

	if obj.obj.Tcp != nil {

		obj.Tcp().validateObj(vObj, set_default)
	}

	if obj.obj.Udp != nil {

		obj.Udp().validateObj(vObj, set_default)
	}

	if obj.obj.Gre != nil {

		obj.Gre().validateObj(vObj, set_default)
	}

	if obj.obj.Gtpv1 != nil {

		obj.Gtpv1().validateObj(vObj, set_default)
	}

	if obj.obj.Gtpv2 != nil {

		obj.Gtpv2().validateObj(vObj, set_default)
	}

	if obj.obj.Arp != nil {

		obj.Arp().validateObj(vObj, set_default)
	}

	if obj.obj.Icmp != nil {

		obj.Icmp().validateObj(vObj, set_default)
	}

	if obj.obj.Icmpv6 != nil {

		obj.Icmpv6().validateObj(vObj, set_default)
	}

	if obj.obj.Ppp != nil {

		obj.Ppp().validateObj(vObj, set_default)
	}

	if obj.obj.Igmpv1 != nil {

		obj.Igmpv1().validateObj(vObj, set_default)
	}

	if obj.obj.Mpls != nil {

		obj.Mpls().validateObj(vObj, set_default)
	}

	if obj.obj.Snmpv2C != nil {

		obj.Snmpv2C().validateObj(vObj, set_default)
	}

}

func (obj *flowHeader) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(FlowHeaderChoice.ETHERNET)

	}

}

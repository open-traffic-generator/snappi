package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSegmentRoutingPolicyTypeG *****
type bgpAttributesSegmentRoutingPolicyTypeG struct {
	validation
	obj             *otg.BgpAttributesSegmentRoutingPolicyTypeG
	marshaller      marshalBgpAttributesSegmentRoutingPolicyTypeG
	unMarshaller    unMarshalBgpAttributesSegmentRoutingPolicyTypeG
	flagsHolder     BgpAttributesSegmentRoutingPolicyTypeFlags
	srMplsSidHolder BgpAttributesSidMpls
}

func NewBgpAttributesSegmentRoutingPolicyTypeG() BgpAttributesSegmentRoutingPolicyTypeG {
	obj := bgpAttributesSegmentRoutingPolicyTypeG{obj: &otg.BgpAttributesSegmentRoutingPolicyTypeG{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeG) msg() *otg.BgpAttributesSegmentRoutingPolicyTypeG {
	return obj.obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeG) setMsg(msg *otg.BgpAttributesSegmentRoutingPolicyTypeG) BgpAttributesSegmentRoutingPolicyTypeG {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSegmentRoutingPolicyTypeG struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeG
}

type marshalBgpAttributesSegmentRoutingPolicyTypeG interface {
	// ToProto marshals BgpAttributesSegmentRoutingPolicyTypeG to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeG
	ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeG, error)
	// ToPbText marshals BgpAttributesSegmentRoutingPolicyTypeG to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSegmentRoutingPolicyTypeG to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSegmentRoutingPolicyTypeG to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesSegmentRoutingPolicyTypeG to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesSegmentRoutingPolicyTypeG struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeG
}

type unMarshalBgpAttributesSegmentRoutingPolicyTypeG interface {
	// FromProto unmarshals BgpAttributesSegmentRoutingPolicyTypeG from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeG
	FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeG) (BgpAttributesSegmentRoutingPolicyTypeG, error)
	// FromPbText unmarshals BgpAttributesSegmentRoutingPolicyTypeG from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSegmentRoutingPolicyTypeG from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSegmentRoutingPolicyTypeG from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeG) Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeG {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSegmentRoutingPolicyTypeG{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeG) Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeG {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSegmentRoutingPolicyTypeG{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeG) ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeG, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeG) FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeG) (BgpAttributesSegmentRoutingPolicyTypeG, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeG) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeG) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeG) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeG) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeG) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeG) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeG) FromJson(value string) error {
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeG) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeG) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeG) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeG) Clone() (BgpAttributesSegmentRoutingPolicyTypeG, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSegmentRoutingPolicyTypeG()
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeG) setNil() {
	obj.flagsHolder = nil
	obj.srMplsSidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesSegmentRoutingPolicyTypeG is type G: IPv6 Address, Interface ID for local and remote pair with optional SID for SR MPLS.
// It is encoded as a Segment of Type 7 in the SEGMENT_LIST sub-tlv.
type BgpAttributesSegmentRoutingPolicyTypeG interface {
	Validation
	// msg marshals BgpAttributesSegmentRoutingPolicyTypeG to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeG
	// and doesn't set defaults
	msg() *otg.BgpAttributesSegmentRoutingPolicyTypeG
	// setMsg unmarshals BgpAttributesSegmentRoutingPolicyTypeG from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeG
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSegmentRoutingPolicyTypeG) BgpAttributesSegmentRoutingPolicyTypeG
	// provides marshal interface
	Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeG
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeG
	// validate validates BgpAttributesSegmentRoutingPolicyTypeG
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSegmentRoutingPolicyTypeG, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns BgpAttributesSegmentRoutingPolicyTypeFlags, set in BgpAttributesSegmentRoutingPolicyTypeG.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	Flags() BgpAttributesSegmentRoutingPolicyTypeFlags
	// SetFlags assigns BgpAttributesSegmentRoutingPolicyTypeFlags provided by user to BgpAttributesSegmentRoutingPolicyTypeG.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeG
	// HasFlags checks if Flags has been set in BgpAttributesSegmentRoutingPolicyTypeG
	HasFlags() bool
	// LocalInterfaceId returns uint32, set in BgpAttributesSegmentRoutingPolicyTypeG.
	LocalInterfaceId() uint32
	// SetLocalInterfaceId assigns uint32 provided by user to BgpAttributesSegmentRoutingPolicyTypeG
	SetLocalInterfaceId(value uint32) BgpAttributesSegmentRoutingPolicyTypeG
	// HasLocalInterfaceId checks if LocalInterfaceId has been set in BgpAttributesSegmentRoutingPolicyTypeG
	HasLocalInterfaceId() bool
	// LocalIpv6NodeAddress returns string, set in BgpAttributesSegmentRoutingPolicyTypeG.
	LocalIpv6NodeAddress() string
	// SetLocalIpv6NodeAddress assigns string provided by user to BgpAttributesSegmentRoutingPolicyTypeG
	SetLocalIpv6NodeAddress(value string) BgpAttributesSegmentRoutingPolicyTypeG
	// HasLocalIpv6NodeAddress checks if LocalIpv6NodeAddress has been set in BgpAttributesSegmentRoutingPolicyTypeG
	HasLocalIpv6NodeAddress() bool
	// RemoteInterfaceId returns uint32, set in BgpAttributesSegmentRoutingPolicyTypeG.
	RemoteInterfaceId() uint32
	// SetRemoteInterfaceId assigns uint32 provided by user to BgpAttributesSegmentRoutingPolicyTypeG
	SetRemoteInterfaceId(value uint32) BgpAttributesSegmentRoutingPolicyTypeG
	// HasRemoteInterfaceId checks if RemoteInterfaceId has been set in BgpAttributesSegmentRoutingPolicyTypeG
	HasRemoteInterfaceId() bool
	// RemoteIpv6NodeAddress returns string, set in BgpAttributesSegmentRoutingPolicyTypeG.
	RemoteIpv6NodeAddress() string
	// SetRemoteIpv6NodeAddress assigns string provided by user to BgpAttributesSegmentRoutingPolicyTypeG
	SetRemoteIpv6NodeAddress(value string) BgpAttributesSegmentRoutingPolicyTypeG
	// HasRemoteIpv6NodeAddress checks if RemoteIpv6NodeAddress has been set in BgpAttributesSegmentRoutingPolicyTypeG
	HasRemoteIpv6NodeAddress() bool
	// SrMplsSid returns BgpAttributesSidMpls, set in BgpAttributesSegmentRoutingPolicyTypeG.
	// BgpAttributesSidMpls is this carries a 20 bit Multi Protocol Label Switching alongwith 3 bits traffic class, 1 bit indicating presence
	// or absence of Bottom-Of-Stack and 8 bits carrying the Time to Live value.
	SrMplsSid() BgpAttributesSidMpls
	// SetSrMplsSid assigns BgpAttributesSidMpls provided by user to BgpAttributesSegmentRoutingPolicyTypeG.
	// BgpAttributesSidMpls is this carries a 20 bit Multi Protocol Label Switching alongwith 3 bits traffic class, 1 bit indicating presence
	// or absence of Bottom-Of-Stack and 8 bits carrying the Time to Live value.
	SetSrMplsSid(value BgpAttributesSidMpls) BgpAttributesSegmentRoutingPolicyTypeG
	// HasSrMplsSid checks if SrMplsSid has been set in BgpAttributesSegmentRoutingPolicyTypeG
	HasSrMplsSid() bool
	setNil()
}

// description is TBD
// Flags returns a BgpAttributesSegmentRoutingPolicyTypeFlags
func (obj *bgpAttributesSegmentRoutingPolicyTypeG) Flags() BgpAttributesSegmentRoutingPolicyTypeFlags {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewBgpAttributesSegmentRoutingPolicyTypeFlags().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &bgpAttributesSegmentRoutingPolicyTypeFlags{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// description is TBD
// Flags returns a BgpAttributesSegmentRoutingPolicyTypeFlags
func (obj *bgpAttributesSegmentRoutingPolicyTypeG) HasFlags() bool {
	return obj.obj.Flags != nil
}

// description is TBD
// SetFlags sets the BgpAttributesSegmentRoutingPolicyTypeFlags value in the BgpAttributesSegmentRoutingPolicyTypeG object
func (obj *bgpAttributesSegmentRoutingPolicyTypeG) SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeG {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// The local Interface Index as defined in [RFC8664].
// LocalInterfaceId returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicyTypeG) LocalInterfaceId() uint32 {

	return *obj.obj.LocalInterfaceId

}

// The local Interface Index as defined in [RFC8664].
// LocalInterfaceId returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicyTypeG) HasLocalInterfaceId() bool {
	return obj.obj.LocalInterfaceId != nil
}

// The local Interface Index as defined in [RFC8664].
// SetLocalInterfaceId sets the uint32 value in the BgpAttributesSegmentRoutingPolicyTypeG object
func (obj *bgpAttributesSegmentRoutingPolicyTypeG) SetLocalInterfaceId(value uint32) BgpAttributesSegmentRoutingPolicyTypeG {

	obj.obj.LocalInterfaceId = &value
	return obj
}

// The IPv6 address representing the local node.
// LocalIpv6NodeAddress returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeG) LocalIpv6NodeAddress() string {

	return *obj.obj.LocalIpv6NodeAddress

}

// The IPv6 address representing the local node.
// LocalIpv6NodeAddress returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeG) HasLocalIpv6NodeAddress() bool {
	return obj.obj.LocalIpv6NodeAddress != nil
}

// The IPv6 address representing the local node.
// SetLocalIpv6NodeAddress sets the string value in the BgpAttributesSegmentRoutingPolicyTypeG object
func (obj *bgpAttributesSegmentRoutingPolicyTypeG) SetLocalIpv6NodeAddress(value string) BgpAttributesSegmentRoutingPolicyTypeG {

	obj.obj.LocalIpv6NodeAddress = &value
	return obj
}

// The remote Interface Index as defined in [RFC8664]. The value MAY be set to zero when the local node address and interface identifiers are sufficient to describe the link.
// RemoteInterfaceId returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicyTypeG) RemoteInterfaceId() uint32 {

	return *obj.obj.RemoteInterfaceId

}

// The remote Interface Index as defined in [RFC8664]. The value MAY be set to zero when the local node address and interface identifiers are sufficient to describe the link.
// RemoteInterfaceId returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicyTypeG) HasRemoteInterfaceId() bool {
	return obj.obj.RemoteInterfaceId != nil
}

// The remote Interface Index as defined in [RFC8664]. The value MAY be set to zero when the local node address and interface identifiers are sufficient to describe the link.
// SetRemoteInterfaceId sets the uint32 value in the BgpAttributesSegmentRoutingPolicyTypeG object
func (obj *bgpAttributesSegmentRoutingPolicyTypeG) SetRemoteInterfaceId(value uint32) BgpAttributesSegmentRoutingPolicyTypeG {

	obj.obj.RemoteInterfaceId = &value
	return obj
}

// IPv6 address representing the remote node. The value MAY be set to zero when the local node address and interface identifiers are sufficient to describe the link.
// RemoteIpv6NodeAddress returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeG) RemoteIpv6NodeAddress() string {

	return *obj.obj.RemoteIpv6NodeAddress

}

// IPv6 address representing the remote node. The value MAY be set to zero when the local node address and interface identifiers are sufficient to describe the link.
// RemoteIpv6NodeAddress returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeG) HasRemoteIpv6NodeAddress() bool {
	return obj.obj.RemoteIpv6NodeAddress != nil
}

// IPv6 address representing the remote node. The value MAY be set to zero when the local node address and interface identifiers are sufficient to describe the link.
// SetRemoteIpv6NodeAddress sets the string value in the BgpAttributesSegmentRoutingPolicyTypeG object
func (obj *bgpAttributesSegmentRoutingPolicyTypeG) SetRemoteIpv6NodeAddress(value string) BgpAttributesSegmentRoutingPolicyTypeG {

	obj.obj.RemoteIpv6NodeAddress = &value
	return obj
}

// Optional SR-MPLS SID.
// SrMplsSid returns a BgpAttributesSidMpls
func (obj *bgpAttributesSegmentRoutingPolicyTypeG) SrMplsSid() BgpAttributesSidMpls {
	if obj.obj.SrMplsSid == nil {
		obj.obj.SrMplsSid = NewBgpAttributesSidMpls().msg()
	}
	if obj.srMplsSidHolder == nil {
		obj.srMplsSidHolder = &bgpAttributesSidMpls{obj: obj.obj.SrMplsSid}
	}
	return obj.srMplsSidHolder
}

// Optional SR-MPLS SID.
// SrMplsSid returns a BgpAttributesSidMpls
func (obj *bgpAttributesSegmentRoutingPolicyTypeG) HasSrMplsSid() bool {
	return obj.obj.SrMplsSid != nil
}

// Optional SR-MPLS SID.
// SetSrMplsSid sets the BgpAttributesSidMpls value in the BgpAttributesSegmentRoutingPolicyTypeG object
func (obj *bgpAttributesSegmentRoutingPolicyTypeG) SetSrMplsSid(value BgpAttributesSidMpls) BgpAttributesSegmentRoutingPolicyTypeG {

	obj.srMplsSidHolder = nil
	obj.obj.SrMplsSid = value.msg()

	return obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeG) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

	if obj.obj.LocalIpv6NodeAddress != nil {

		err := obj.validateIpv6(obj.LocalIpv6NodeAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesSegmentRoutingPolicyTypeG.LocalIpv6NodeAddress"))
		}

	}

	if obj.obj.RemoteIpv6NodeAddress != nil {

		err := obj.validateIpv6(obj.RemoteIpv6NodeAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesSegmentRoutingPolicyTypeG.RemoteIpv6NodeAddress"))
		}

	}

	if obj.obj.SrMplsSid != nil {

		obj.SrMplsSid().validateObj(vObj, set_default)
	}

}

func (obj *bgpAttributesSegmentRoutingPolicyTypeG) setDefault() {
	if obj.obj.LocalInterfaceId == nil {
		obj.SetLocalInterfaceId(0)
	}
	if obj.obj.LocalIpv6NodeAddress == nil {
		obj.SetLocalIpv6NodeAddress("0::0")
	}
	if obj.obj.RemoteInterfaceId == nil {
		obj.SetRemoteInterfaceId(0)
	}
	if obj.obj.RemoteIpv6NodeAddress == nil {
		obj.SetRemoteIpv6NodeAddress("0::0")
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSegmentRoutingPolicyTypeH *****
type bgpAttributesSegmentRoutingPolicyTypeH struct {
	validation
	obj             *otg.BgpAttributesSegmentRoutingPolicyTypeH
	marshaller      marshalBgpAttributesSegmentRoutingPolicyTypeH
	unMarshaller    unMarshalBgpAttributesSegmentRoutingPolicyTypeH
	flagsHolder     BgpAttributesSegmentRoutingPolicyTypeFlags
	srMplsSidHolder BgpAttributesSidMpls
}

func NewBgpAttributesSegmentRoutingPolicyTypeH() BgpAttributesSegmentRoutingPolicyTypeH {
	obj := bgpAttributesSegmentRoutingPolicyTypeH{obj: &otg.BgpAttributesSegmentRoutingPolicyTypeH{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeH) msg() *otg.BgpAttributesSegmentRoutingPolicyTypeH {
	return obj.obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeH) setMsg(msg *otg.BgpAttributesSegmentRoutingPolicyTypeH) BgpAttributesSegmentRoutingPolicyTypeH {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSegmentRoutingPolicyTypeH struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeH
}

type marshalBgpAttributesSegmentRoutingPolicyTypeH interface {
	// ToProto marshals BgpAttributesSegmentRoutingPolicyTypeH to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeH
	ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeH, error)
	// ToPbText marshals BgpAttributesSegmentRoutingPolicyTypeH to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSegmentRoutingPolicyTypeH to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSegmentRoutingPolicyTypeH to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesSegmentRoutingPolicyTypeH to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesSegmentRoutingPolicyTypeH struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeH
}

type unMarshalBgpAttributesSegmentRoutingPolicyTypeH interface {
	// FromProto unmarshals BgpAttributesSegmentRoutingPolicyTypeH from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeH
	FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeH) (BgpAttributesSegmentRoutingPolicyTypeH, error)
	// FromPbText unmarshals BgpAttributesSegmentRoutingPolicyTypeH from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSegmentRoutingPolicyTypeH from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSegmentRoutingPolicyTypeH from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeH) Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeH {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSegmentRoutingPolicyTypeH{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeH) Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeH {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSegmentRoutingPolicyTypeH{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeH) ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeH, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeH) FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeH) (BgpAttributesSegmentRoutingPolicyTypeH, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeH) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeH) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeH) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeH) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeH) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeH) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeH) FromJson(value string) error {
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeH) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeH) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeH) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeH) Clone() (BgpAttributesSegmentRoutingPolicyTypeH, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSegmentRoutingPolicyTypeH()
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeH) setNil() {
	obj.flagsHolder = nil
	obj.srMplsSidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesSegmentRoutingPolicyTypeH is type H: IPv6 Local and Remote addresses with optional SID for SR MPLS.
// It is encoded as a Segment of Type 8 in the SEGMENT_LIST sub-tlv.
type BgpAttributesSegmentRoutingPolicyTypeH interface {
	Validation
	// msg marshals BgpAttributesSegmentRoutingPolicyTypeH to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeH
	// and doesn't set defaults
	msg() *otg.BgpAttributesSegmentRoutingPolicyTypeH
	// setMsg unmarshals BgpAttributesSegmentRoutingPolicyTypeH from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeH
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSegmentRoutingPolicyTypeH) BgpAttributesSegmentRoutingPolicyTypeH
	// provides marshal interface
	Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeH
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeH
	// validate validates BgpAttributesSegmentRoutingPolicyTypeH
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSegmentRoutingPolicyTypeH, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns BgpAttributesSegmentRoutingPolicyTypeFlags, set in BgpAttributesSegmentRoutingPolicyTypeH.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	Flags() BgpAttributesSegmentRoutingPolicyTypeFlags
	// SetFlags assigns BgpAttributesSegmentRoutingPolicyTypeFlags provided by user to BgpAttributesSegmentRoutingPolicyTypeH.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeH
	// HasFlags checks if Flags has been set in BgpAttributesSegmentRoutingPolicyTypeH
	HasFlags() bool
	// LocalIpv6Address returns string, set in BgpAttributesSegmentRoutingPolicyTypeH.
	LocalIpv6Address() string
	// SetLocalIpv6Address assigns string provided by user to BgpAttributesSegmentRoutingPolicyTypeH
	SetLocalIpv6Address(value string) BgpAttributesSegmentRoutingPolicyTypeH
	// HasLocalIpv6Address checks if LocalIpv6Address has been set in BgpAttributesSegmentRoutingPolicyTypeH
	HasLocalIpv6Address() bool
	// RemoteIpv6Address returns string, set in BgpAttributesSegmentRoutingPolicyTypeH.
	RemoteIpv6Address() string
	// SetRemoteIpv6Address assigns string provided by user to BgpAttributesSegmentRoutingPolicyTypeH
	SetRemoteIpv6Address(value string) BgpAttributesSegmentRoutingPolicyTypeH
	// HasRemoteIpv6Address checks if RemoteIpv6Address has been set in BgpAttributesSegmentRoutingPolicyTypeH
	HasRemoteIpv6Address() bool
	// SrMplsSid returns BgpAttributesSidMpls, set in BgpAttributesSegmentRoutingPolicyTypeH.
	// BgpAttributesSidMpls is this carries a 20 bit Multi Protocol Label Switching alongwith 3 bits traffic class, 1 bit indicating presence
	// or absence of Bottom-Of-Stack and 8 bits carrying the Time to Live value.
	SrMplsSid() BgpAttributesSidMpls
	// SetSrMplsSid assigns BgpAttributesSidMpls provided by user to BgpAttributesSegmentRoutingPolicyTypeH.
	// BgpAttributesSidMpls is this carries a 20 bit Multi Protocol Label Switching alongwith 3 bits traffic class, 1 bit indicating presence
	// or absence of Bottom-Of-Stack and 8 bits carrying the Time to Live value.
	SetSrMplsSid(value BgpAttributesSidMpls) BgpAttributesSegmentRoutingPolicyTypeH
	// HasSrMplsSid checks if SrMplsSid has been set in BgpAttributesSegmentRoutingPolicyTypeH
	HasSrMplsSid() bool
	setNil()
}

// description is TBD
// Flags returns a BgpAttributesSegmentRoutingPolicyTypeFlags
func (obj *bgpAttributesSegmentRoutingPolicyTypeH) Flags() BgpAttributesSegmentRoutingPolicyTypeFlags {
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
func (obj *bgpAttributesSegmentRoutingPolicyTypeH) HasFlags() bool {
	return obj.obj.Flags != nil
}

// description is TBD
// SetFlags sets the BgpAttributesSegmentRoutingPolicyTypeFlags value in the BgpAttributesSegmentRoutingPolicyTypeH object
func (obj *bgpAttributesSegmentRoutingPolicyTypeH) SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeH {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// Local IPv6 Address.
// LocalIpv6Address returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeH) LocalIpv6Address() string {

	return *obj.obj.LocalIpv6Address

}

// Local IPv6 Address.
// LocalIpv6Address returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeH) HasLocalIpv6Address() bool {
	return obj.obj.LocalIpv6Address != nil
}

// Local IPv6 Address.
// SetLocalIpv6Address sets the string value in the BgpAttributesSegmentRoutingPolicyTypeH object
func (obj *bgpAttributesSegmentRoutingPolicyTypeH) SetLocalIpv6Address(value string) BgpAttributesSegmentRoutingPolicyTypeH {

	obj.obj.LocalIpv6Address = &value
	return obj
}

// Remote IPv6 Address.
// RemoteIpv6Address returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeH) RemoteIpv6Address() string {

	return *obj.obj.RemoteIpv6Address

}

// Remote IPv6 Address.
// RemoteIpv6Address returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeH) HasRemoteIpv6Address() bool {
	return obj.obj.RemoteIpv6Address != nil
}

// Remote IPv6 Address.
// SetRemoteIpv6Address sets the string value in the BgpAttributesSegmentRoutingPolicyTypeH object
func (obj *bgpAttributesSegmentRoutingPolicyTypeH) SetRemoteIpv6Address(value string) BgpAttributesSegmentRoutingPolicyTypeH {

	obj.obj.RemoteIpv6Address = &value
	return obj
}

// Optional SR-MPLS SID.
// SrMplsSid returns a BgpAttributesSidMpls
func (obj *bgpAttributesSegmentRoutingPolicyTypeH) SrMplsSid() BgpAttributesSidMpls {
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
func (obj *bgpAttributesSegmentRoutingPolicyTypeH) HasSrMplsSid() bool {
	return obj.obj.SrMplsSid != nil
}

// Optional SR-MPLS SID.
// SetSrMplsSid sets the BgpAttributesSidMpls value in the BgpAttributesSegmentRoutingPolicyTypeH object
func (obj *bgpAttributesSegmentRoutingPolicyTypeH) SetSrMplsSid(value BgpAttributesSidMpls) BgpAttributesSegmentRoutingPolicyTypeH {

	obj.srMplsSidHolder = nil
	obj.obj.SrMplsSid = value.msg()

	return obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeH) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

	if obj.obj.LocalIpv6Address != nil {

		err := obj.validateIpv6(obj.LocalIpv6Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesSegmentRoutingPolicyTypeH.LocalIpv6Address"))
		}

	}

	if obj.obj.RemoteIpv6Address != nil {

		err := obj.validateIpv6(obj.RemoteIpv6Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesSegmentRoutingPolicyTypeH.RemoteIpv6Address"))
		}

	}

	if obj.obj.SrMplsSid != nil {

		obj.SrMplsSid().validateObj(vObj, set_default)
	}

}

func (obj *bgpAttributesSegmentRoutingPolicyTypeH) setDefault() {
	if obj.obj.LocalIpv6Address == nil {
		obj.SetLocalIpv6Address("0::0")
	}
	if obj.obj.RemoteIpv6Address == nil {
		obj.SetRemoteIpv6Address("0::0")
	}

}

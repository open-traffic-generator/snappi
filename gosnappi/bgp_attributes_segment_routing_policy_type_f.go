package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSegmentRoutingPolicyTypeF *****
type bgpAttributesSegmentRoutingPolicyTypeF struct {
	validation
	obj             *otg.BgpAttributesSegmentRoutingPolicyTypeF
	marshaller      marshalBgpAttributesSegmentRoutingPolicyTypeF
	unMarshaller    unMarshalBgpAttributesSegmentRoutingPolicyTypeF
	flagsHolder     BgpAttributesSegmentRoutingPolicyTypeFlags
	srMplsSidHolder BgpAttributesSidMpls
}

func NewBgpAttributesSegmentRoutingPolicyTypeF() BgpAttributesSegmentRoutingPolicyTypeF {
	obj := bgpAttributesSegmentRoutingPolicyTypeF{obj: &otg.BgpAttributesSegmentRoutingPolicyTypeF{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeF) msg() *otg.BgpAttributesSegmentRoutingPolicyTypeF {
	return obj.obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeF) setMsg(msg *otg.BgpAttributesSegmentRoutingPolicyTypeF) BgpAttributesSegmentRoutingPolicyTypeF {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSegmentRoutingPolicyTypeF struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeF
}

type marshalBgpAttributesSegmentRoutingPolicyTypeF interface {
	// ToProto marshals BgpAttributesSegmentRoutingPolicyTypeF to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeF
	ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeF, error)
	// ToPbText marshals BgpAttributesSegmentRoutingPolicyTypeF to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSegmentRoutingPolicyTypeF to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSegmentRoutingPolicyTypeF to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesSegmentRoutingPolicyTypeF to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesSegmentRoutingPolicyTypeF struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeF
}

type unMarshalBgpAttributesSegmentRoutingPolicyTypeF interface {
	// FromProto unmarshals BgpAttributesSegmentRoutingPolicyTypeF from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeF
	FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeF) (BgpAttributesSegmentRoutingPolicyTypeF, error)
	// FromPbText unmarshals BgpAttributesSegmentRoutingPolicyTypeF from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSegmentRoutingPolicyTypeF from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSegmentRoutingPolicyTypeF from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeF) Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeF {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSegmentRoutingPolicyTypeF{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeF) Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeF {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSegmentRoutingPolicyTypeF{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeF) ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeF, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeF) FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeF) (BgpAttributesSegmentRoutingPolicyTypeF, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeF) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeF) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeF) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeF) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeF) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeF) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeF) FromJson(value string) error {
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeF) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeF) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeF) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeF) Clone() (BgpAttributesSegmentRoutingPolicyTypeF, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSegmentRoutingPolicyTypeF()
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeF) setNil() {
	obj.flagsHolder = nil
	obj.srMplsSidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesSegmentRoutingPolicyTypeF is type F: IPv4 Local and Remote addresses with optional SR-MPLS SID.
// It is encoded as a Segment of Type 6 in the SEGMENT_LIST sub-tlv.
type BgpAttributesSegmentRoutingPolicyTypeF interface {
	Validation
	// msg marshals BgpAttributesSegmentRoutingPolicyTypeF to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeF
	// and doesn't set defaults
	msg() *otg.BgpAttributesSegmentRoutingPolicyTypeF
	// setMsg unmarshals BgpAttributesSegmentRoutingPolicyTypeF from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeF
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSegmentRoutingPolicyTypeF) BgpAttributesSegmentRoutingPolicyTypeF
	// provides marshal interface
	Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeF
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeF
	// validate validates BgpAttributesSegmentRoutingPolicyTypeF
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSegmentRoutingPolicyTypeF, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns BgpAttributesSegmentRoutingPolicyTypeFlags, set in BgpAttributesSegmentRoutingPolicyTypeF.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	Flags() BgpAttributesSegmentRoutingPolicyTypeFlags
	// SetFlags assigns BgpAttributesSegmentRoutingPolicyTypeFlags provided by user to BgpAttributesSegmentRoutingPolicyTypeF.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeF
	// HasFlags checks if Flags has been set in BgpAttributesSegmentRoutingPolicyTypeF
	HasFlags() bool
	// LocalIpv4Address returns string, set in BgpAttributesSegmentRoutingPolicyTypeF.
	LocalIpv4Address() string
	// SetLocalIpv4Address assigns string provided by user to BgpAttributesSegmentRoutingPolicyTypeF
	SetLocalIpv4Address(value string) BgpAttributesSegmentRoutingPolicyTypeF
	// HasLocalIpv4Address checks if LocalIpv4Address has been set in BgpAttributesSegmentRoutingPolicyTypeF
	HasLocalIpv4Address() bool
	// RemoteIpv4Address returns string, set in BgpAttributesSegmentRoutingPolicyTypeF.
	RemoteIpv4Address() string
	// SetRemoteIpv4Address assigns string provided by user to BgpAttributesSegmentRoutingPolicyTypeF
	SetRemoteIpv4Address(value string) BgpAttributesSegmentRoutingPolicyTypeF
	// HasRemoteIpv4Address checks if RemoteIpv4Address has been set in BgpAttributesSegmentRoutingPolicyTypeF
	HasRemoteIpv4Address() bool
	// SrMplsSid returns BgpAttributesSidMpls, set in BgpAttributesSegmentRoutingPolicyTypeF.
	// BgpAttributesSidMpls is this carries a 20 bit Multi Protocol Label Switching alongwith 3 bits traffic class, 1 bit indicating presence
	// or absence of Bottom-Of-Stack and 8 bits carrying the Time to Live value.
	SrMplsSid() BgpAttributesSidMpls
	// SetSrMplsSid assigns BgpAttributesSidMpls provided by user to BgpAttributesSegmentRoutingPolicyTypeF.
	// BgpAttributesSidMpls is this carries a 20 bit Multi Protocol Label Switching alongwith 3 bits traffic class, 1 bit indicating presence
	// or absence of Bottom-Of-Stack and 8 bits carrying the Time to Live value.
	SetSrMplsSid(value BgpAttributesSidMpls) BgpAttributesSegmentRoutingPolicyTypeF
	// HasSrMplsSid checks if SrMplsSid has been set in BgpAttributesSegmentRoutingPolicyTypeF
	HasSrMplsSid() bool
	setNil()
}

// description is TBD
// Flags returns a BgpAttributesSegmentRoutingPolicyTypeFlags
func (obj *bgpAttributesSegmentRoutingPolicyTypeF) Flags() BgpAttributesSegmentRoutingPolicyTypeFlags {
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
func (obj *bgpAttributesSegmentRoutingPolicyTypeF) HasFlags() bool {
	return obj.obj.Flags != nil
}

// description is TBD
// SetFlags sets the BgpAttributesSegmentRoutingPolicyTypeFlags value in the BgpAttributesSegmentRoutingPolicyTypeF object
func (obj *bgpAttributesSegmentRoutingPolicyTypeF) SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeF {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// Local IPv4 Address.
// LocalIpv4Address returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeF) LocalIpv4Address() string {

	return *obj.obj.LocalIpv4Address

}

// Local IPv4 Address.
// LocalIpv4Address returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeF) HasLocalIpv4Address() bool {
	return obj.obj.LocalIpv4Address != nil
}

// Local IPv4 Address.
// SetLocalIpv4Address sets the string value in the BgpAttributesSegmentRoutingPolicyTypeF object
func (obj *bgpAttributesSegmentRoutingPolicyTypeF) SetLocalIpv4Address(value string) BgpAttributesSegmentRoutingPolicyTypeF {

	obj.obj.LocalIpv4Address = &value
	return obj
}

// Remote IPv4 Address.
// RemoteIpv4Address returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeF) RemoteIpv4Address() string {

	return *obj.obj.RemoteIpv4Address

}

// Remote IPv4 Address.
// RemoteIpv4Address returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeF) HasRemoteIpv4Address() bool {
	return obj.obj.RemoteIpv4Address != nil
}

// Remote IPv4 Address.
// SetRemoteIpv4Address sets the string value in the BgpAttributesSegmentRoutingPolicyTypeF object
func (obj *bgpAttributesSegmentRoutingPolicyTypeF) SetRemoteIpv4Address(value string) BgpAttributesSegmentRoutingPolicyTypeF {

	obj.obj.RemoteIpv4Address = &value
	return obj
}

// Optional SR-MPLS SID.
// SrMplsSid returns a BgpAttributesSidMpls
func (obj *bgpAttributesSegmentRoutingPolicyTypeF) SrMplsSid() BgpAttributesSidMpls {
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
func (obj *bgpAttributesSegmentRoutingPolicyTypeF) HasSrMplsSid() bool {
	return obj.obj.SrMplsSid != nil
}

// Optional SR-MPLS SID.
// SetSrMplsSid sets the BgpAttributesSidMpls value in the BgpAttributesSegmentRoutingPolicyTypeF object
func (obj *bgpAttributesSegmentRoutingPolicyTypeF) SetSrMplsSid(value BgpAttributesSidMpls) BgpAttributesSegmentRoutingPolicyTypeF {

	obj.srMplsSidHolder = nil
	obj.obj.SrMplsSid = value.msg()

	return obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeF) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

	if obj.obj.LocalIpv4Address != nil {

		err := obj.validateIpv4(obj.LocalIpv4Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesSegmentRoutingPolicyTypeF.LocalIpv4Address"))
		}

	}

	if obj.obj.RemoteIpv4Address != nil {

		err := obj.validateIpv4(obj.RemoteIpv4Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesSegmentRoutingPolicyTypeF.RemoteIpv4Address"))
		}

	}

	if obj.obj.SrMplsSid != nil {

		obj.SrMplsSid().validateObj(vObj, set_default)
	}

}

func (obj *bgpAttributesSegmentRoutingPolicyTypeF) setDefault() {
	if obj.obj.LocalIpv4Address == nil {
		obj.SetLocalIpv4Address("0.0.0.0")
	}
	if obj.obj.RemoteIpv4Address == nil {
		obj.SetRemoteIpv4Address("0.0.0.0")
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSegmentRoutingPolicyTypeA *****
type bgpAttributesSegmentRoutingPolicyTypeA struct {
	validation
	obj           *otg.BgpAttributesSegmentRoutingPolicyTypeA
	marshaller    marshalBgpAttributesSegmentRoutingPolicyTypeA
	unMarshaller  unMarshalBgpAttributesSegmentRoutingPolicyTypeA
	flagsHolder   BgpAttributesSegmentRoutingPolicyTypeFlags
	mplsSidHolder BgpAttributesSidMpls
}

func NewBgpAttributesSegmentRoutingPolicyTypeA() BgpAttributesSegmentRoutingPolicyTypeA {
	obj := bgpAttributesSegmentRoutingPolicyTypeA{obj: &otg.BgpAttributesSegmentRoutingPolicyTypeA{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeA) msg() *otg.BgpAttributesSegmentRoutingPolicyTypeA {
	return obj.obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeA) setMsg(msg *otg.BgpAttributesSegmentRoutingPolicyTypeA) BgpAttributesSegmentRoutingPolicyTypeA {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSegmentRoutingPolicyTypeA struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeA
}

type marshalBgpAttributesSegmentRoutingPolicyTypeA interface {
	// ToProto marshals BgpAttributesSegmentRoutingPolicyTypeA to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeA
	ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeA, error)
	// ToPbText marshals BgpAttributesSegmentRoutingPolicyTypeA to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSegmentRoutingPolicyTypeA to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSegmentRoutingPolicyTypeA to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributesSegmentRoutingPolicyTypeA struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeA
}

type unMarshalBgpAttributesSegmentRoutingPolicyTypeA interface {
	// FromProto unmarshals BgpAttributesSegmentRoutingPolicyTypeA from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeA
	FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeA) (BgpAttributesSegmentRoutingPolicyTypeA, error)
	// FromPbText unmarshals BgpAttributesSegmentRoutingPolicyTypeA from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSegmentRoutingPolicyTypeA from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSegmentRoutingPolicyTypeA from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeA) Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeA {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSegmentRoutingPolicyTypeA{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeA) Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeA {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSegmentRoutingPolicyTypeA{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeA) ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeA, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeA) FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeA) (BgpAttributesSegmentRoutingPolicyTypeA, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeA) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeA) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeA) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeA) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeA) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeA) FromJson(value string) error {
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeA) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeA) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeA) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeA) Clone() (BgpAttributesSegmentRoutingPolicyTypeA, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSegmentRoutingPolicyTypeA()
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeA) setNil() {
	obj.flagsHolder = nil
	obj.mplsSidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesSegmentRoutingPolicyTypeA is type A: SID only, in the form of MPLS Label.
// It is encoded as a Segment of Type 1 in the SEGMENT_LIST sub-tlv.
type BgpAttributesSegmentRoutingPolicyTypeA interface {
	Validation
	// msg marshals BgpAttributesSegmentRoutingPolicyTypeA to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeA
	// and doesn't set defaults
	msg() *otg.BgpAttributesSegmentRoutingPolicyTypeA
	// setMsg unmarshals BgpAttributesSegmentRoutingPolicyTypeA from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeA
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSegmentRoutingPolicyTypeA) BgpAttributesSegmentRoutingPolicyTypeA
	// provides marshal interface
	Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeA
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeA
	// validate validates BgpAttributesSegmentRoutingPolicyTypeA
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSegmentRoutingPolicyTypeA, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns BgpAttributesSegmentRoutingPolicyTypeFlags, set in BgpAttributesSegmentRoutingPolicyTypeA.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	Flags() BgpAttributesSegmentRoutingPolicyTypeFlags
	// SetFlags assigns BgpAttributesSegmentRoutingPolicyTypeFlags provided by user to BgpAttributesSegmentRoutingPolicyTypeA.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeA
	// HasFlags checks if Flags has been set in BgpAttributesSegmentRoutingPolicyTypeA
	HasFlags() bool
	// MplsSid returns BgpAttributesSidMpls, set in BgpAttributesSegmentRoutingPolicyTypeA.
	// BgpAttributesSidMpls is this carries a 20 bit Multi Protocol Label Switching alongwith 3 bits traffic class, 1 bit indicating presence
	// or absence of Bottom-Of-Stack and 8 bits carrying the Time to Live value.
	MplsSid() BgpAttributesSidMpls
	// SetMplsSid assigns BgpAttributesSidMpls provided by user to BgpAttributesSegmentRoutingPolicyTypeA.
	// BgpAttributesSidMpls is this carries a 20 bit Multi Protocol Label Switching alongwith 3 bits traffic class, 1 bit indicating presence
	// or absence of Bottom-Of-Stack and 8 bits carrying the Time to Live value.
	SetMplsSid(value BgpAttributesSidMpls) BgpAttributesSegmentRoutingPolicyTypeA
	// HasMplsSid checks if MplsSid has been set in BgpAttributesSegmentRoutingPolicyTypeA
	HasMplsSid() bool
	setNil()
}

// description is TBD
// Flags returns a BgpAttributesSegmentRoutingPolicyTypeFlags
func (obj *bgpAttributesSegmentRoutingPolicyTypeA) Flags() BgpAttributesSegmentRoutingPolicyTypeFlags {
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
func (obj *bgpAttributesSegmentRoutingPolicyTypeA) HasFlags() bool {
	return obj.obj.Flags != nil
}

// description is TBD
// SetFlags sets the BgpAttributesSegmentRoutingPolicyTypeFlags value in the BgpAttributesSegmentRoutingPolicyTypeA object
func (obj *bgpAttributesSegmentRoutingPolicyTypeA) SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeA {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// description is TBD
// MplsSid returns a BgpAttributesSidMpls
func (obj *bgpAttributesSegmentRoutingPolicyTypeA) MplsSid() BgpAttributesSidMpls {
	if obj.obj.MplsSid == nil {
		obj.obj.MplsSid = NewBgpAttributesSidMpls().msg()
	}
	if obj.mplsSidHolder == nil {
		obj.mplsSidHolder = &bgpAttributesSidMpls{obj: obj.obj.MplsSid}
	}
	return obj.mplsSidHolder
}

// description is TBD
// MplsSid returns a BgpAttributesSidMpls
func (obj *bgpAttributesSegmentRoutingPolicyTypeA) HasMplsSid() bool {
	return obj.obj.MplsSid != nil
}

// description is TBD
// SetMplsSid sets the BgpAttributesSidMpls value in the BgpAttributesSegmentRoutingPolicyTypeA object
func (obj *bgpAttributesSegmentRoutingPolicyTypeA) SetMplsSid(value BgpAttributesSidMpls) BgpAttributesSegmentRoutingPolicyTypeA {

	obj.mplsSidHolder = nil
	obj.obj.MplsSid = value.msg()

	return obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeA) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

	if obj.obj.MplsSid != nil {

		obj.MplsSid().validateObj(vObj, set_default)
	}

}

func (obj *bgpAttributesSegmentRoutingPolicyTypeA) setDefault() {

}

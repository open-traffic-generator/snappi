package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSegmentRoutingPolicyTypeFlags *****
type bgpAttributesSegmentRoutingPolicyTypeFlags struct {
	validation
	obj          *otg.BgpAttributesSegmentRoutingPolicyTypeFlags
	marshaller   marshalBgpAttributesSegmentRoutingPolicyTypeFlags
	unMarshaller unMarshalBgpAttributesSegmentRoutingPolicyTypeFlags
}

func NewBgpAttributesSegmentRoutingPolicyTypeFlags() BgpAttributesSegmentRoutingPolicyTypeFlags {
	obj := bgpAttributesSegmentRoutingPolicyTypeFlags{obj: &otg.BgpAttributesSegmentRoutingPolicyTypeFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) msg() *otg.BgpAttributesSegmentRoutingPolicyTypeFlags {
	return obj.obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) setMsg(msg *otg.BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeFlags {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSegmentRoutingPolicyTypeFlags struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeFlags
}

type marshalBgpAttributesSegmentRoutingPolicyTypeFlags interface {
	// ToProto marshals BgpAttributesSegmentRoutingPolicyTypeFlags to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeFlags
	ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeFlags, error)
	// ToPbText marshals BgpAttributesSegmentRoutingPolicyTypeFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSegmentRoutingPolicyTypeFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSegmentRoutingPolicyTypeFlags to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesSegmentRoutingPolicyTypeFlags to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesSegmentRoutingPolicyTypeFlags struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeFlags
}

type unMarshalBgpAttributesSegmentRoutingPolicyTypeFlags interface {
	// FromProto unmarshals BgpAttributesSegmentRoutingPolicyTypeFlags from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeFlags
	FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeFlags) (BgpAttributesSegmentRoutingPolicyTypeFlags, error)
	// FromPbText unmarshals BgpAttributesSegmentRoutingPolicyTypeFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSegmentRoutingPolicyTypeFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSegmentRoutingPolicyTypeFlags from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSegmentRoutingPolicyTypeFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSegmentRoutingPolicyTypeFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeFlags) ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeFlags) FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeFlags) (BgpAttributesSegmentRoutingPolicyTypeFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeFlags) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeFlags) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeFlags) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeFlags) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeFlags) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeFlags) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeFlags) FromJson(value string) error {
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) Clone() (BgpAttributesSegmentRoutingPolicyTypeFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSegmentRoutingPolicyTypeFlags()
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

// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
type BgpAttributesSegmentRoutingPolicyTypeFlags interface {
	Validation
	// msg marshals BgpAttributesSegmentRoutingPolicyTypeFlags to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeFlags
	// and doesn't set defaults
	msg() *otg.BgpAttributesSegmentRoutingPolicyTypeFlags
	// setMsg unmarshals BgpAttributesSegmentRoutingPolicyTypeFlags from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeFlags
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeFlags
	// provides marshal interface
	Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeFlags
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeFlags
	// validate validates BgpAttributesSegmentRoutingPolicyTypeFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSegmentRoutingPolicyTypeFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// VFlag returns bool, set in BgpAttributesSegmentRoutingPolicyTypeFlags.
	VFlag() bool
	// SetVFlag assigns bool provided by user to BgpAttributesSegmentRoutingPolicyTypeFlags
	SetVFlag(value bool) BgpAttributesSegmentRoutingPolicyTypeFlags
	// HasVFlag checks if VFlag has been set in BgpAttributesSegmentRoutingPolicyTypeFlags
	HasVFlag() bool
	// AFlag returns bool, set in BgpAttributesSegmentRoutingPolicyTypeFlags.
	AFlag() bool
	// SetAFlag assigns bool provided by user to BgpAttributesSegmentRoutingPolicyTypeFlags
	SetAFlag(value bool) BgpAttributesSegmentRoutingPolicyTypeFlags
	// HasAFlag checks if AFlag has been set in BgpAttributesSegmentRoutingPolicyTypeFlags
	HasAFlag() bool
	// SFlag returns bool, set in BgpAttributesSegmentRoutingPolicyTypeFlags.
	SFlag() bool
	// SetSFlag assigns bool provided by user to BgpAttributesSegmentRoutingPolicyTypeFlags
	SetSFlag(value bool) BgpAttributesSegmentRoutingPolicyTypeFlags
	// HasSFlag checks if SFlag has been set in BgpAttributesSegmentRoutingPolicyTypeFlags
	HasSFlag() bool
	// BFlag returns bool, set in BgpAttributesSegmentRoutingPolicyTypeFlags.
	BFlag() bool
	// SetBFlag assigns bool provided by user to BgpAttributesSegmentRoutingPolicyTypeFlags
	SetBFlag(value bool) BgpAttributesSegmentRoutingPolicyTypeFlags
	// HasBFlag checks if BFlag has been set in BgpAttributesSegmentRoutingPolicyTypeFlags
	HasBFlag() bool
}

// Indicates verification of segment data in is enabled.
// VFlag returns a bool
func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) VFlag() bool {

	return *obj.obj.VFlag

}

// Indicates verification of segment data in is enabled.
// VFlag returns a bool
func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) HasVFlag() bool {
	return obj.obj.VFlag != nil
}

// Indicates verification of segment data in is enabled.
// SetVFlag sets the bool value in the BgpAttributesSegmentRoutingPolicyTypeFlags object
func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) SetVFlag(value bool) BgpAttributesSegmentRoutingPolicyTypeFlags {

	obj.obj.VFlag = &value
	return obj
}

// Indicates presence of SR Algorithm field applicable to Segment Types 3, 4, and 9.
// AFlag returns a bool
func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) AFlag() bool {

	return *obj.obj.AFlag

}

// Indicates presence of SR Algorithm field applicable to Segment Types 3, 4, and 9.
// AFlag returns a bool
func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) HasAFlag() bool {
	return obj.obj.AFlag != nil
}

// Indicates presence of SR Algorithm field applicable to Segment Types 3, 4, and 9.
// SetAFlag sets the bool value in the BgpAttributesSegmentRoutingPolicyTypeFlags object
func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) SetAFlag(value bool) BgpAttributesSegmentRoutingPolicyTypeFlags {

	obj.obj.AFlag = &value
	return obj
}

// This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type.
// SFlag returns a bool
func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) SFlag() bool {

	return *obj.obj.SFlag

}

// This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type.
// SFlag returns a bool
func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) HasSFlag() bool {
	return obj.obj.SFlag != nil
}

// This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type.
// SetSFlag sets the bool value in the BgpAttributesSegmentRoutingPolicyTypeFlags object
func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) SetSFlag(value bool) BgpAttributesSegmentRoutingPolicyTypeFlags {

	obj.obj.SFlag = &value
	return obj
}

// Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding specified in Section 2.4.4.2.4
// of draft-ietf-idr-sr-policy-safi-02.
// BFlag returns a bool
func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) BFlag() bool {

	return *obj.obj.BFlag

}

// Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding specified in Section 2.4.4.2.4
// of draft-ietf-idr-sr-policy-safi-02.
// BFlag returns a bool
func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) HasBFlag() bool {
	return obj.obj.BFlag != nil
}

// Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding specified in Section 2.4.4.2.4
// of draft-ietf-idr-sr-policy-safi-02.
// SetBFlag sets the bool value in the BgpAttributesSegmentRoutingPolicyTypeFlags object
func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) SetBFlag(value bool) BgpAttributesSegmentRoutingPolicyTypeFlags {

	obj.obj.BFlag = &value
	return obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpAttributesSegmentRoutingPolicyTypeFlags) setDefault() {
	if obj.obj.VFlag == nil {
		obj.SetVFlag(false)
	}
	if obj.obj.AFlag == nil {
		obj.SetAFlag(false)
	}
	if obj.obj.SFlag == nil {
		obj.SetSFlag(false)
	}
	if obj.obj.BFlag == nil {
		obj.SetBFlag(false)
	}

}

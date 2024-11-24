package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSrv6Bsid *****
type bgpAttributesSrv6Bsid struct {
	validation
	obj                        *otg.BgpAttributesSrv6Bsid
	marshaller                 marshalBgpAttributesSrv6Bsid
	unMarshaller               unMarshalBgpAttributesSrv6Bsid
	srv6EndpointBehaviorHolder BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
}

func NewBgpAttributesSrv6Bsid() BgpAttributesSrv6Bsid {
	obj := bgpAttributesSrv6Bsid{obj: &otg.BgpAttributesSrv6Bsid{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSrv6Bsid) msg() *otg.BgpAttributesSrv6Bsid {
	return obj.obj
}

func (obj *bgpAttributesSrv6Bsid) setMsg(msg *otg.BgpAttributesSrv6Bsid) BgpAttributesSrv6Bsid {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSrv6Bsid struct {
	obj *bgpAttributesSrv6Bsid
}

type marshalBgpAttributesSrv6Bsid interface {
	// ToProto marshals BgpAttributesSrv6Bsid to protobuf object *otg.BgpAttributesSrv6Bsid
	ToProto() (*otg.BgpAttributesSrv6Bsid, error)
	// ToPbText marshals BgpAttributesSrv6Bsid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSrv6Bsid to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSrv6Bsid to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesSrv6Bsid to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesSrv6Bsid struct {
	obj *bgpAttributesSrv6Bsid
}

type unMarshalBgpAttributesSrv6Bsid interface {
	// FromProto unmarshals BgpAttributesSrv6Bsid from protobuf object *otg.BgpAttributesSrv6Bsid
	FromProto(msg *otg.BgpAttributesSrv6Bsid) (BgpAttributesSrv6Bsid, error)
	// FromPbText unmarshals BgpAttributesSrv6Bsid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSrv6Bsid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSrv6Bsid from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSrv6Bsid) Marshal() marshalBgpAttributesSrv6Bsid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSrv6Bsid{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSrv6Bsid) Unmarshal() unMarshalBgpAttributesSrv6Bsid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSrv6Bsid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSrv6Bsid) ToProto() (*otg.BgpAttributesSrv6Bsid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSrv6Bsid) FromProto(msg *otg.BgpAttributesSrv6Bsid) (BgpAttributesSrv6Bsid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSrv6Bsid) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSrv6Bsid) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSrv6Bsid) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSrv6Bsid) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSrv6Bsid) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesSrv6Bsid) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSrv6Bsid) FromJson(value string) error {
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

func (obj *bgpAttributesSrv6Bsid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSrv6Bsid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSrv6Bsid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSrv6Bsid) Clone() (BgpAttributesSrv6Bsid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSrv6Bsid()
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

func (obj *bgpAttributesSrv6Bsid) setNil() {
	obj.srv6EndpointBehaviorHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesSrv6Bsid is the SRv6 Binding SID sub-TLV is an optional sub-TLV of type 20 that is used to signal the SRv6 Binding SID
// related information of an SR Policy candidate path.
// - More than one SRv6 Binding SID sub-TLVs MAY be signaled in the same SR Policy encoding to indicate one or
// more SRv6 SIDs, each with potentially different SRv6 Endpoint Behaviors to be instantiated.
// - The format of the sub-TLV is defined in draft-ietf-idr-sr-policy-safi-02 Section 2.4.3 .
type BgpAttributesSrv6Bsid interface {
	Validation
	// msg marshals BgpAttributesSrv6Bsid to protobuf object *otg.BgpAttributesSrv6Bsid
	// and doesn't set defaults
	msg() *otg.BgpAttributesSrv6Bsid
	// setMsg unmarshals BgpAttributesSrv6Bsid from protobuf object *otg.BgpAttributesSrv6Bsid
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSrv6Bsid) BgpAttributesSrv6Bsid
	// provides marshal interface
	Marshal() marshalBgpAttributesSrv6Bsid
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSrv6Bsid
	// validate validates BgpAttributesSrv6Bsid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSrv6Bsid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// FlagSpecifiedBsidOnly returns bool, set in BgpAttributesSrv6Bsid.
	FlagSpecifiedBsidOnly() bool
	// SetFlagSpecifiedBsidOnly assigns bool provided by user to BgpAttributesSrv6Bsid
	SetFlagSpecifiedBsidOnly(value bool) BgpAttributesSrv6Bsid
	// HasFlagSpecifiedBsidOnly checks if FlagSpecifiedBsidOnly has been set in BgpAttributesSrv6Bsid
	HasFlagSpecifiedBsidOnly() bool
	// FlagDropUponInvalid returns bool, set in BgpAttributesSrv6Bsid.
	FlagDropUponInvalid() bool
	// SetFlagDropUponInvalid assigns bool provided by user to BgpAttributesSrv6Bsid
	SetFlagDropUponInvalid(value bool) BgpAttributesSrv6Bsid
	// HasFlagDropUponInvalid checks if FlagDropUponInvalid has been set in BgpAttributesSrv6Bsid
	HasFlagDropUponInvalid() bool
	// FlagSrv6EndpointBehavior returns bool, set in BgpAttributesSrv6Bsid.
	FlagSrv6EndpointBehavior() bool
	// SetFlagSrv6EndpointBehavior assigns bool provided by user to BgpAttributesSrv6Bsid
	SetFlagSrv6EndpointBehavior(value bool) BgpAttributesSrv6Bsid
	// HasFlagSrv6EndpointBehavior checks if FlagSrv6EndpointBehavior has been set in BgpAttributesSrv6Bsid
	HasFlagSrv6EndpointBehavior() bool
	// Ipv6Addr returns string, set in BgpAttributesSrv6Bsid.
	Ipv6Addr() string
	// SetIpv6Addr assigns string provided by user to BgpAttributesSrv6Bsid
	SetIpv6Addr(value string) BgpAttributesSrv6Bsid
	// HasIpv6Addr checks if Ipv6Addr has been set in BgpAttributesSrv6Bsid
	HasIpv6Addr() bool
	// Srv6EndpointBehavior returns BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure, set in BgpAttributesSrv6Bsid.
	// BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure is configuration for optional SRv6 Endpoint Behavior and SID Structure. Summation of lengths for Locator Block, Locator Node,  Function, and Argument MUST be less than or equal to 128. - This is specified in draft-ietf-idr-sr-policy-safi-02 Section 2.4.4.2.4
	Srv6EndpointBehavior() BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	// SetSrv6EndpointBehavior assigns BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure provided by user to BgpAttributesSrv6Bsid.
	// BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure is configuration for optional SRv6 Endpoint Behavior and SID Structure. Summation of lengths for Locator Block, Locator Node,  Function, and Argument MUST be less than or equal to 128. - This is specified in draft-ietf-idr-sr-policy-safi-02 Section 2.4.4.2.4
	SetSrv6EndpointBehavior(value BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) BgpAttributesSrv6Bsid
	// HasSrv6EndpointBehavior checks if Srv6EndpointBehavior has been set in BgpAttributesSrv6Bsid
	HasSrv6EndpointBehavior() bool
	setNil()
}

// S-Flag: This flag encodes the "Specified-BSID-only" behavior. It's usage is
// described in section 6.2.3 in [RFC9256].
// FlagSpecifiedBsidOnly returns a bool
func (obj *bgpAttributesSrv6Bsid) FlagSpecifiedBsidOnly() bool {

	return *obj.obj.FlagSpecifiedBsidOnly

}

// S-Flag: This flag encodes the "Specified-BSID-only" behavior. It's usage is
// described in section 6.2.3 in [RFC9256].
// FlagSpecifiedBsidOnly returns a bool
func (obj *bgpAttributesSrv6Bsid) HasFlagSpecifiedBsidOnly() bool {
	return obj.obj.FlagSpecifiedBsidOnly != nil
}

// S-Flag: This flag encodes the "Specified-BSID-only" behavior. It's usage is
// described in section 6.2.3 in [RFC9256].
// SetFlagSpecifiedBsidOnly sets the bool value in the BgpAttributesSrv6Bsid object
func (obj *bgpAttributesSrv6Bsid) SetFlagSpecifiedBsidOnly(value bool) BgpAttributesSrv6Bsid {

	obj.obj.FlagSpecifiedBsidOnly = &value
	return obj
}

// I-Flag: This flag encodes the "Drop Upon Invalid" behavior.
// It's usage is described in section 8.2 in [RFC9256].
// FlagDropUponInvalid returns a bool
func (obj *bgpAttributesSrv6Bsid) FlagDropUponInvalid() bool {

	return *obj.obj.FlagDropUponInvalid

}

// I-Flag: This flag encodes the "Drop Upon Invalid" behavior.
// It's usage is described in section 8.2 in [RFC9256].
// FlagDropUponInvalid returns a bool
func (obj *bgpAttributesSrv6Bsid) HasFlagDropUponInvalid() bool {
	return obj.obj.FlagDropUponInvalid != nil
}

// I-Flag: This flag encodes the "Drop Upon Invalid" behavior.
// It's usage is described in section 8.2 in [RFC9256].
// SetFlagDropUponInvalid sets the bool value in the BgpAttributesSrv6Bsid object
func (obj *bgpAttributesSrv6Bsid) SetFlagDropUponInvalid(value bool) BgpAttributesSrv6Bsid {

	obj.obj.FlagDropUponInvalid = &value
	return obj
}

// B-Flag: This flag, when set, indicates the presence of the SRv6 Endpoint Behavior
// and SID Structure encoding specified in Section 2.4.4.2.4 of draft-ietf-idr-sr-policy-safi-02.
// FlagSrv6EndpointBehavior returns a bool
func (obj *bgpAttributesSrv6Bsid) FlagSrv6EndpointBehavior() bool {

	return *obj.obj.FlagSrv6EndpointBehavior

}

// B-Flag: This flag, when set, indicates the presence of the SRv6 Endpoint Behavior
// and SID Structure encoding specified in Section 2.4.4.2.4 of draft-ietf-idr-sr-policy-safi-02.
// FlagSrv6EndpointBehavior returns a bool
func (obj *bgpAttributesSrv6Bsid) HasFlagSrv6EndpointBehavior() bool {
	return obj.obj.FlagSrv6EndpointBehavior != nil
}

// B-Flag: This flag, when set, indicates the presence of the SRv6 Endpoint Behavior
// and SID Structure encoding specified in Section 2.4.4.2.4 of draft-ietf-idr-sr-policy-safi-02.
// SetFlagSrv6EndpointBehavior sets the bool value in the BgpAttributesSrv6Bsid object
func (obj *bgpAttributesSrv6Bsid) SetFlagSrv6EndpointBehavior(value bool) BgpAttributesSrv6Bsid {

	obj.obj.FlagSrv6EndpointBehavior = &value
	return obj
}

// IPv6 address denoting the SRv6 SID.
// Ipv6Addr returns a string
func (obj *bgpAttributesSrv6Bsid) Ipv6Addr() string {

	return *obj.obj.Ipv6Addr

}

// IPv6 address denoting the SRv6 SID.
// Ipv6Addr returns a string
func (obj *bgpAttributesSrv6Bsid) HasIpv6Addr() bool {
	return obj.obj.Ipv6Addr != nil
}

// IPv6 address denoting the SRv6 SID.
// SetIpv6Addr sets the string value in the BgpAttributesSrv6Bsid object
func (obj *bgpAttributesSrv6Bsid) SetIpv6Addr(value string) BgpAttributesSrv6Bsid {

	obj.obj.Ipv6Addr = &value
	return obj
}

// description is TBD
// Srv6EndpointBehavior returns a BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
func (obj *bgpAttributesSrv6Bsid) Srv6EndpointBehavior() BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure {
	if obj.obj.Srv6EndpointBehavior == nil {
		obj.obj.Srv6EndpointBehavior = NewBgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure().msg()
	}
	if obj.srv6EndpointBehaviorHolder == nil {
		obj.srv6EndpointBehaviorHolder = &bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure{obj: obj.obj.Srv6EndpointBehavior}
	}
	return obj.srv6EndpointBehaviorHolder
}

// description is TBD
// Srv6EndpointBehavior returns a BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
func (obj *bgpAttributesSrv6Bsid) HasSrv6EndpointBehavior() bool {
	return obj.obj.Srv6EndpointBehavior != nil
}

// description is TBD
// SetSrv6EndpointBehavior sets the BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure value in the BgpAttributesSrv6Bsid object
func (obj *bgpAttributesSrv6Bsid) SetSrv6EndpointBehavior(value BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) BgpAttributesSrv6Bsid {

	obj.srv6EndpointBehaviorHolder = nil
	obj.obj.Srv6EndpointBehavior = value.msg()

	return obj
}

func (obj *bgpAttributesSrv6Bsid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv6Addr != nil {

		err := obj.validateIpv6(obj.Ipv6Addr())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesSrv6Bsid.Ipv6Addr"))
		}

	}

	if obj.obj.Srv6EndpointBehavior != nil {

		obj.Srv6EndpointBehavior().validateObj(vObj, set_default)
	}

}

func (obj *bgpAttributesSrv6Bsid) setDefault() {
	if obj.obj.FlagSpecifiedBsidOnly == nil {
		obj.SetFlagSpecifiedBsidOnly(false)
	}
	if obj.obj.FlagDropUponInvalid == nil {
		obj.SetFlagDropUponInvalid(false)
	}
	if obj.obj.FlagSrv6EndpointBehavior == nil {
		obj.SetFlagSrv6EndpointBehavior(false)
	}
	if obj.obj.Ipv6Addr == nil {
		obj.SetIpv6Addr("0::0")
	}

}

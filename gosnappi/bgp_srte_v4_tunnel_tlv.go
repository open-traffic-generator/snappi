package gosnappi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrteV4TunnelTlv *****
type bgpSrteV4TunnelTlv struct {
	validation
	obj                                 *otg.BgpSrteV4TunnelTlv
	marshaller                          marshalBgpSrteV4TunnelTlv
	unMarshaller                        unMarshalBgpSrteV4TunnelTlv
	remoteEndpointSubTlvHolder          BgpSrteRemoteEndpointSubTlv
	colorSubTlvHolder                   BgpSrteColorSubTlv
	bindingSubTlvHolder                 BgpSrteBindingSubTlv
	preferenceSubTlvHolder              BgpSrtePreferenceSubTlv
	policyPrioritySubTlvHolder          BgpSrtePolicyPrioritySubTlv
	policyNameSubTlvHolder              BgpSrtePolicyNameSubTlv
	explicitNullLabelPolicySubTlvHolder BgpSrteExplicitNullLabelPolicySubTlv
	segmentListsHolder                  BgpSrteV4TunnelTlvBgpSrteSegmentListIter
}

func NewBgpSrteV4TunnelTlv() BgpSrteV4TunnelTlv {
	obj := bgpSrteV4TunnelTlv{obj: &otg.BgpSrteV4TunnelTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteV4TunnelTlv) msg() *otg.BgpSrteV4TunnelTlv {
	return obj.obj
}

func (obj *bgpSrteV4TunnelTlv) setMsg(msg *otg.BgpSrteV4TunnelTlv) BgpSrteV4TunnelTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteV4TunnelTlv struct {
	obj *bgpSrteV4TunnelTlv
}

type marshalBgpSrteV4TunnelTlv interface {
	// ToProto marshals BgpSrteV4TunnelTlv to protobuf object *otg.BgpSrteV4TunnelTlv
	ToProto() (*otg.BgpSrteV4TunnelTlv, error)
	// ToPbText marshals BgpSrteV4TunnelTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteV4TunnelTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteV4TunnelTlv to JSON text
	ToJson() (string, error)
}

type unMarshalbgpSrteV4TunnelTlv struct {
	obj *bgpSrteV4TunnelTlv
}

type unMarshalBgpSrteV4TunnelTlv interface {
	// FromProto unmarshals BgpSrteV4TunnelTlv from protobuf object *otg.BgpSrteV4TunnelTlv
	FromProto(msg *otg.BgpSrteV4TunnelTlv) (BgpSrteV4TunnelTlv, error)
	// FromPbText unmarshals BgpSrteV4TunnelTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteV4TunnelTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteV4TunnelTlv from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteV4TunnelTlv) Marshal() marshalBgpSrteV4TunnelTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteV4TunnelTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteV4TunnelTlv) Unmarshal() unMarshalBgpSrteV4TunnelTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteV4TunnelTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteV4TunnelTlv) ToProto() (*otg.BgpSrteV4TunnelTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteV4TunnelTlv) FromProto(msg *otg.BgpSrteV4TunnelTlv) (BgpSrteV4TunnelTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteV4TunnelTlv) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteV4TunnelTlv) FromPbText(value string) error {
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

func (m *marshalbgpSrteV4TunnelTlv) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteV4TunnelTlv) FromYaml(value string) error {
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

func (m *marshalbgpSrteV4TunnelTlv) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteV4TunnelTlv) FromJson(value string) error {
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

func (obj *bgpSrteV4TunnelTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteV4TunnelTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteV4TunnelTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteV4TunnelTlv) Clone() (BgpSrteV4TunnelTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteV4TunnelTlv()
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

func (obj *bgpSrteV4TunnelTlv) setNil() {
	obj.remoteEndpointSubTlvHolder = nil
	obj.colorSubTlvHolder = nil
	obj.bindingSubTlvHolder = nil
	obj.preferenceSubTlvHolder = nil
	obj.policyPrioritySubTlvHolder = nil
	obj.policyNameSubTlvHolder = nil
	obj.explicitNullLabelPolicySubTlvHolder = nil
	obj.segmentListsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpSrteV4TunnelTlv is configuration for BGP SRTE Tunnel TLV.
type BgpSrteV4TunnelTlv interface {
	Validation
	// msg marshals BgpSrteV4TunnelTlv to protobuf object *otg.BgpSrteV4TunnelTlv
	// and doesn't set defaults
	msg() *otg.BgpSrteV4TunnelTlv
	// setMsg unmarshals BgpSrteV4TunnelTlv from protobuf object *otg.BgpSrteV4TunnelTlv
	// and doesn't set defaults
	setMsg(*otg.BgpSrteV4TunnelTlv) BgpSrteV4TunnelTlv
	// provides marshal interface
	Marshal() marshalBgpSrteV4TunnelTlv
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteV4TunnelTlv
	// validate validates BgpSrteV4TunnelTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteV4TunnelTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RemoteEndpointSubTlv returns BgpSrteRemoteEndpointSubTlv, set in BgpSrteV4TunnelTlv.
	// BgpSrteRemoteEndpointSubTlv is configuration for the BGP remote endpoint sub TLV.
	RemoteEndpointSubTlv() BgpSrteRemoteEndpointSubTlv
	// SetRemoteEndpointSubTlv assigns BgpSrteRemoteEndpointSubTlv provided by user to BgpSrteV4TunnelTlv.
	// BgpSrteRemoteEndpointSubTlv is configuration for the BGP remote endpoint sub TLV.
	SetRemoteEndpointSubTlv(value BgpSrteRemoteEndpointSubTlv) BgpSrteV4TunnelTlv
	// HasRemoteEndpointSubTlv checks if RemoteEndpointSubTlv has been set in BgpSrteV4TunnelTlv
	HasRemoteEndpointSubTlv() bool
	// ColorSubTlv returns BgpSrteColorSubTlv, set in BgpSrteV4TunnelTlv.
	// BgpSrteColorSubTlv is configuration for the Policy Color attribute sub-TLV. The Color sub-TLV MAY be used as a way to "color" the corresponding Tunnel TLV. The Value field of the sub-TLV is eight octets long and consists of a Color Extended Community. First two octets of its Value field are 0x030b as type and subtype of extended  community. Remaining six octets are are exposed to configure.
	ColorSubTlv() BgpSrteColorSubTlv
	// SetColorSubTlv assigns BgpSrteColorSubTlv provided by user to BgpSrteV4TunnelTlv.
	// BgpSrteColorSubTlv is configuration for the Policy Color attribute sub-TLV. The Color sub-TLV MAY be used as a way to "color" the corresponding Tunnel TLV. The Value field of the sub-TLV is eight octets long and consists of a Color Extended Community. First two octets of its Value field are 0x030b as type and subtype of extended  community. Remaining six octets are are exposed to configure.
	SetColorSubTlv(value BgpSrteColorSubTlv) BgpSrteV4TunnelTlv
	// HasColorSubTlv checks if ColorSubTlv has been set in BgpSrteV4TunnelTlv
	HasColorSubTlv() bool
	// BindingSubTlv returns BgpSrteBindingSubTlv, set in BgpSrteV4TunnelTlv.
	// BgpSrteBindingSubTlv is configuration for the binding SID sub-TLV.  This is used to signal the binding SID related information  of the SR Policy candidate path.
	BindingSubTlv() BgpSrteBindingSubTlv
	// SetBindingSubTlv assigns BgpSrteBindingSubTlv provided by user to BgpSrteV4TunnelTlv.
	// BgpSrteBindingSubTlv is configuration for the binding SID sub-TLV.  This is used to signal the binding SID related information  of the SR Policy candidate path.
	SetBindingSubTlv(value BgpSrteBindingSubTlv) BgpSrteV4TunnelTlv
	// HasBindingSubTlv checks if BindingSubTlv has been set in BgpSrteV4TunnelTlv
	HasBindingSubTlv() bool
	// PreferenceSubTlv returns BgpSrtePreferenceSubTlv, set in BgpSrteV4TunnelTlv.
	// BgpSrtePreferenceSubTlv is configuration for BGP preference sub TLV of the SR Policy candidate path.
	PreferenceSubTlv() BgpSrtePreferenceSubTlv
	// SetPreferenceSubTlv assigns BgpSrtePreferenceSubTlv provided by user to BgpSrteV4TunnelTlv.
	// BgpSrtePreferenceSubTlv is configuration for BGP preference sub TLV of the SR Policy candidate path.
	SetPreferenceSubTlv(value BgpSrtePreferenceSubTlv) BgpSrteV4TunnelTlv
	// HasPreferenceSubTlv checks if PreferenceSubTlv has been set in BgpSrteV4TunnelTlv
	HasPreferenceSubTlv() bool
	// PolicyPrioritySubTlv returns BgpSrtePolicyPrioritySubTlv, set in BgpSrteV4TunnelTlv.
	// BgpSrtePolicyPrioritySubTlv is configuration for the Policy Priority sub-TLV. The Policy Priority to indicate the order in which the SR policies  are re-computed upon topological change.
	PolicyPrioritySubTlv() BgpSrtePolicyPrioritySubTlv
	// SetPolicyPrioritySubTlv assigns BgpSrtePolicyPrioritySubTlv provided by user to BgpSrteV4TunnelTlv.
	// BgpSrtePolicyPrioritySubTlv is configuration for the Policy Priority sub-TLV. The Policy Priority to indicate the order in which the SR policies  are re-computed upon topological change.
	SetPolicyPrioritySubTlv(value BgpSrtePolicyPrioritySubTlv) BgpSrteV4TunnelTlv
	// HasPolicyPrioritySubTlv checks if PolicyPrioritySubTlv has been set in BgpSrteV4TunnelTlv
	HasPolicyPrioritySubTlv() bool
	// PolicyNameSubTlv returns BgpSrtePolicyNameSubTlv, set in BgpSrteV4TunnelTlv.
	// BgpSrtePolicyNameSubTlv is configuration for the Policy Name sub-TLV. The Policy Name sub-TLV is used to attach a symbolic name to the SR Policy candidate path.
	PolicyNameSubTlv() BgpSrtePolicyNameSubTlv
	// SetPolicyNameSubTlv assigns BgpSrtePolicyNameSubTlv provided by user to BgpSrteV4TunnelTlv.
	// BgpSrtePolicyNameSubTlv is configuration for the Policy Name sub-TLV. The Policy Name sub-TLV is used to attach a symbolic name to the SR Policy candidate path.
	SetPolicyNameSubTlv(value BgpSrtePolicyNameSubTlv) BgpSrteV4TunnelTlv
	// HasPolicyNameSubTlv checks if PolicyNameSubTlv has been set in BgpSrteV4TunnelTlv
	HasPolicyNameSubTlv() bool
	// ExplicitNullLabelPolicySubTlv returns BgpSrteExplicitNullLabelPolicySubTlv, set in BgpSrteV4TunnelTlv.
	// BgpSrteExplicitNullLabelPolicySubTlv is configuration for BGP explicit null label policy sub TLV settings.
	ExplicitNullLabelPolicySubTlv() BgpSrteExplicitNullLabelPolicySubTlv
	// SetExplicitNullLabelPolicySubTlv assigns BgpSrteExplicitNullLabelPolicySubTlv provided by user to BgpSrteV4TunnelTlv.
	// BgpSrteExplicitNullLabelPolicySubTlv is configuration for BGP explicit null label policy sub TLV settings.
	SetExplicitNullLabelPolicySubTlv(value BgpSrteExplicitNullLabelPolicySubTlv) BgpSrteV4TunnelTlv
	// HasExplicitNullLabelPolicySubTlv checks if ExplicitNullLabelPolicySubTlv has been set in BgpSrteV4TunnelTlv
	HasExplicitNullLabelPolicySubTlv() bool
	// SegmentLists returns BgpSrteV4TunnelTlvBgpSrteSegmentListIterIter, set in BgpSrteV4TunnelTlv
	SegmentLists() BgpSrteV4TunnelTlvBgpSrteSegmentListIter
	// Name returns string, set in BgpSrteV4TunnelTlv.
	Name() string
	// SetName assigns string provided by user to BgpSrteV4TunnelTlv
	SetName(value string) BgpSrteV4TunnelTlv
	// Active returns bool, set in BgpSrteV4TunnelTlv.
	Active() bool
	// SetActive assigns bool provided by user to BgpSrteV4TunnelTlv
	SetActive(value bool) BgpSrteV4TunnelTlv
	// HasActive checks if Active has been set in BgpSrteV4TunnelTlv
	HasActive() bool
	setNil()
}

// description is TBD
// RemoteEndpointSubTlv returns a BgpSrteRemoteEndpointSubTlv
func (obj *bgpSrteV4TunnelTlv) RemoteEndpointSubTlv() BgpSrteRemoteEndpointSubTlv {
	if obj.obj.RemoteEndpointSubTlv == nil {
		obj.obj.RemoteEndpointSubTlv = NewBgpSrteRemoteEndpointSubTlv().msg()
	}
	if obj.remoteEndpointSubTlvHolder == nil {
		obj.remoteEndpointSubTlvHolder = &bgpSrteRemoteEndpointSubTlv{obj: obj.obj.RemoteEndpointSubTlv}
	}
	return obj.remoteEndpointSubTlvHolder
}

// description is TBD
// RemoteEndpointSubTlv returns a BgpSrteRemoteEndpointSubTlv
func (obj *bgpSrteV4TunnelTlv) HasRemoteEndpointSubTlv() bool {
	return obj.obj.RemoteEndpointSubTlv != nil
}

// description is TBD
// SetRemoteEndpointSubTlv sets the BgpSrteRemoteEndpointSubTlv value in the BgpSrteV4TunnelTlv object
func (obj *bgpSrteV4TunnelTlv) SetRemoteEndpointSubTlv(value BgpSrteRemoteEndpointSubTlv) BgpSrteV4TunnelTlv {

	obj.remoteEndpointSubTlvHolder = nil
	obj.obj.RemoteEndpointSubTlv = value.msg()

	return obj
}

// description is TBD
// ColorSubTlv returns a BgpSrteColorSubTlv
func (obj *bgpSrteV4TunnelTlv) ColorSubTlv() BgpSrteColorSubTlv {
	if obj.obj.ColorSubTlv == nil {
		obj.obj.ColorSubTlv = NewBgpSrteColorSubTlv().msg()
	}
	if obj.colorSubTlvHolder == nil {
		obj.colorSubTlvHolder = &bgpSrteColorSubTlv{obj: obj.obj.ColorSubTlv}
	}
	return obj.colorSubTlvHolder
}

// description is TBD
// ColorSubTlv returns a BgpSrteColorSubTlv
func (obj *bgpSrteV4TunnelTlv) HasColorSubTlv() bool {
	return obj.obj.ColorSubTlv != nil
}

// description is TBD
// SetColorSubTlv sets the BgpSrteColorSubTlv value in the BgpSrteV4TunnelTlv object
func (obj *bgpSrteV4TunnelTlv) SetColorSubTlv(value BgpSrteColorSubTlv) BgpSrteV4TunnelTlv {

	obj.colorSubTlvHolder = nil
	obj.obj.ColorSubTlv = value.msg()

	return obj
}

// description is TBD
// BindingSubTlv returns a BgpSrteBindingSubTlv
func (obj *bgpSrteV4TunnelTlv) BindingSubTlv() BgpSrteBindingSubTlv {
	if obj.obj.BindingSubTlv == nil {
		obj.obj.BindingSubTlv = NewBgpSrteBindingSubTlv().msg()
	}
	if obj.bindingSubTlvHolder == nil {
		obj.bindingSubTlvHolder = &bgpSrteBindingSubTlv{obj: obj.obj.BindingSubTlv}
	}
	return obj.bindingSubTlvHolder
}

// description is TBD
// BindingSubTlv returns a BgpSrteBindingSubTlv
func (obj *bgpSrteV4TunnelTlv) HasBindingSubTlv() bool {
	return obj.obj.BindingSubTlv != nil
}

// description is TBD
// SetBindingSubTlv sets the BgpSrteBindingSubTlv value in the BgpSrteV4TunnelTlv object
func (obj *bgpSrteV4TunnelTlv) SetBindingSubTlv(value BgpSrteBindingSubTlv) BgpSrteV4TunnelTlv {

	obj.bindingSubTlvHolder = nil
	obj.obj.BindingSubTlv = value.msg()

	return obj
}

// description is TBD
// PreferenceSubTlv returns a BgpSrtePreferenceSubTlv
func (obj *bgpSrteV4TunnelTlv) PreferenceSubTlv() BgpSrtePreferenceSubTlv {
	if obj.obj.PreferenceSubTlv == nil {
		obj.obj.PreferenceSubTlv = NewBgpSrtePreferenceSubTlv().msg()
	}
	if obj.preferenceSubTlvHolder == nil {
		obj.preferenceSubTlvHolder = &bgpSrtePreferenceSubTlv{obj: obj.obj.PreferenceSubTlv}
	}
	return obj.preferenceSubTlvHolder
}

// description is TBD
// PreferenceSubTlv returns a BgpSrtePreferenceSubTlv
func (obj *bgpSrteV4TunnelTlv) HasPreferenceSubTlv() bool {
	return obj.obj.PreferenceSubTlv != nil
}

// description is TBD
// SetPreferenceSubTlv sets the BgpSrtePreferenceSubTlv value in the BgpSrteV4TunnelTlv object
func (obj *bgpSrteV4TunnelTlv) SetPreferenceSubTlv(value BgpSrtePreferenceSubTlv) BgpSrteV4TunnelTlv {

	obj.preferenceSubTlvHolder = nil
	obj.obj.PreferenceSubTlv = value.msg()

	return obj
}

// description is TBD
// PolicyPrioritySubTlv returns a BgpSrtePolicyPrioritySubTlv
func (obj *bgpSrteV4TunnelTlv) PolicyPrioritySubTlv() BgpSrtePolicyPrioritySubTlv {
	if obj.obj.PolicyPrioritySubTlv == nil {
		obj.obj.PolicyPrioritySubTlv = NewBgpSrtePolicyPrioritySubTlv().msg()
	}
	if obj.policyPrioritySubTlvHolder == nil {
		obj.policyPrioritySubTlvHolder = &bgpSrtePolicyPrioritySubTlv{obj: obj.obj.PolicyPrioritySubTlv}
	}
	return obj.policyPrioritySubTlvHolder
}

// description is TBD
// PolicyPrioritySubTlv returns a BgpSrtePolicyPrioritySubTlv
func (obj *bgpSrteV4TunnelTlv) HasPolicyPrioritySubTlv() bool {
	return obj.obj.PolicyPrioritySubTlv != nil
}

// description is TBD
// SetPolicyPrioritySubTlv sets the BgpSrtePolicyPrioritySubTlv value in the BgpSrteV4TunnelTlv object
func (obj *bgpSrteV4TunnelTlv) SetPolicyPrioritySubTlv(value BgpSrtePolicyPrioritySubTlv) BgpSrteV4TunnelTlv {

	obj.policyPrioritySubTlvHolder = nil
	obj.obj.PolicyPrioritySubTlv = value.msg()

	return obj
}

// description is TBD
// PolicyNameSubTlv returns a BgpSrtePolicyNameSubTlv
func (obj *bgpSrteV4TunnelTlv) PolicyNameSubTlv() BgpSrtePolicyNameSubTlv {
	if obj.obj.PolicyNameSubTlv == nil {
		obj.obj.PolicyNameSubTlv = NewBgpSrtePolicyNameSubTlv().msg()
	}
	if obj.policyNameSubTlvHolder == nil {
		obj.policyNameSubTlvHolder = &bgpSrtePolicyNameSubTlv{obj: obj.obj.PolicyNameSubTlv}
	}
	return obj.policyNameSubTlvHolder
}

// description is TBD
// PolicyNameSubTlv returns a BgpSrtePolicyNameSubTlv
func (obj *bgpSrteV4TunnelTlv) HasPolicyNameSubTlv() bool {
	return obj.obj.PolicyNameSubTlv != nil
}

// description is TBD
// SetPolicyNameSubTlv sets the BgpSrtePolicyNameSubTlv value in the BgpSrteV4TunnelTlv object
func (obj *bgpSrteV4TunnelTlv) SetPolicyNameSubTlv(value BgpSrtePolicyNameSubTlv) BgpSrteV4TunnelTlv {

	obj.policyNameSubTlvHolder = nil
	obj.obj.PolicyNameSubTlv = value.msg()

	return obj
}

// description is TBD
// ExplicitNullLabelPolicySubTlv returns a BgpSrteExplicitNullLabelPolicySubTlv
func (obj *bgpSrteV4TunnelTlv) ExplicitNullLabelPolicySubTlv() BgpSrteExplicitNullLabelPolicySubTlv {
	if obj.obj.ExplicitNullLabelPolicySubTlv == nil {
		obj.obj.ExplicitNullLabelPolicySubTlv = NewBgpSrteExplicitNullLabelPolicySubTlv().msg()
	}
	if obj.explicitNullLabelPolicySubTlvHolder == nil {
		obj.explicitNullLabelPolicySubTlvHolder = &bgpSrteExplicitNullLabelPolicySubTlv{obj: obj.obj.ExplicitNullLabelPolicySubTlv}
	}
	return obj.explicitNullLabelPolicySubTlvHolder
}

// description is TBD
// ExplicitNullLabelPolicySubTlv returns a BgpSrteExplicitNullLabelPolicySubTlv
func (obj *bgpSrteV4TunnelTlv) HasExplicitNullLabelPolicySubTlv() bool {
	return obj.obj.ExplicitNullLabelPolicySubTlv != nil
}

// description is TBD
// SetExplicitNullLabelPolicySubTlv sets the BgpSrteExplicitNullLabelPolicySubTlv value in the BgpSrteV4TunnelTlv object
func (obj *bgpSrteV4TunnelTlv) SetExplicitNullLabelPolicySubTlv(value BgpSrteExplicitNullLabelPolicySubTlv) BgpSrteV4TunnelTlv {

	obj.explicitNullLabelPolicySubTlvHolder = nil
	obj.obj.ExplicitNullLabelPolicySubTlv = value.msg()

	return obj
}

// description is TBD
// SegmentLists returns a []BgpSrteSegmentList
func (obj *bgpSrteV4TunnelTlv) SegmentLists() BgpSrteV4TunnelTlvBgpSrteSegmentListIter {
	if len(obj.obj.SegmentLists) == 0 {
		obj.obj.SegmentLists = []*otg.BgpSrteSegmentList{}
	}
	if obj.segmentListsHolder == nil {
		obj.segmentListsHolder = newBgpSrteV4TunnelTlvBgpSrteSegmentListIter(&obj.obj.SegmentLists).setMsg(obj)
	}
	return obj.segmentListsHolder
}

type bgpSrteV4TunnelTlvBgpSrteSegmentListIter struct {
	obj                     *bgpSrteV4TunnelTlv
	bgpSrteSegmentListSlice []BgpSrteSegmentList
	fieldPtr                *[]*otg.BgpSrteSegmentList
}

func newBgpSrteV4TunnelTlvBgpSrteSegmentListIter(ptr *[]*otg.BgpSrteSegmentList) BgpSrteV4TunnelTlvBgpSrteSegmentListIter {
	return &bgpSrteV4TunnelTlvBgpSrteSegmentListIter{fieldPtr: ptr}
}

type BgpSrteV4TunnelTlvBgpSrteSegmentListIter interface {
	setMsg(*bgpSrteV4TunnelTlv) BgpSrteV4TunnelTlvBgpSrteSegmentListIter
	Items() []BgpSrteSegmentList
	Add() BgpSrteSegmentList
	Append(items ...BgpSrteSegmentList) BgpSrteV4TunnelTlvBgpSrteSegmentListIter
	Set(index int, newObj BgpSrteSegmentList) BgpSrteV4TunnelTlvBgpSrteSegmentListIter
	Clear() BgpSrteV4TunnelTlvBgpSrteSegmentListIter
	clearHolderSlice() BgpSrteV4TunnelTlvBgpSrteSegmentListIter
	appendHolderSlice(item BgpSrteSegmentList) BgpSrteV4TunnelTlvBgpSrteSegmentListIter
}

func (obj *bgpSrteV4TunnelTlvBgpSrteSegmentListIter) setMsg(msg *bgpSrteV4TunnelTlv) BgpSrteV4TunnelTlvBgpSrteSegmentListIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpSrteSegmentList{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpSrteV4TunnelTlvBgpSrteSegmentListIter) Items() []BgpSrteSegmentList {
	return obj.bgpSrteSegmentListSlice
}

func (obj *bgpSrteV4TunnelTlvBgpSrteSegmentListIter) Add() BgpSrteSegmentList {
	newObj := &otg.BgpSrteSegmentList{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpSrteSegmentList{obj: newObj}
	newLibObj.setDefault()
	obj.bgpSrteSegmentListSlice = append(obj.bgpSrteSegmentListSlice, newLibObj)
	return newLibObj
}

func (obj *bgpSrteV4TunnelTlvBgpSrteSegmentListIter) Append(items ...BgpSrteSegmentList) BgpSrteV4TunnelTlvBgpSrteSegmentListIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpSrteSegmentListSlice = append(obj.bgpSrteSegmentListSlice, item)
	}
	return obj
}

func (obj *bgpSrteV4TunnelTlvBgpSrteSegmentListIter) Set(index int, newObj BgpSrteSegmentList) BgpSrteV4TunnelTlvBgpSrteSegmentListIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpSrteSegmentListSlice[index] = newObj
	return obj
}
func (obj *bgpSrteV4TunnelTlvBgpSrteSegmentListIter) Clear() BgpSrteV4TunnelTlvBgpSrteSegmentListIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpSrteSegmentList{}
		obj.bgpSrteSegmentListSlice = []BgpSrteSegmentList{}
	}
	return obj
}
func (obj *bgpSrteV4TunnelTlvBgpSrteSegmentListIter) clearHolderSlice() BgpSrteV4TunnelTlvBgpSrteSegmentListIter {
	if len(obj.bgpSrteSegmentListSlice) > 0 {
		obj.bgpSrteSegmentListSlice = []BgpSrteSegmentList{}
	}
	return obj
}
func (obj *bgpSrteV4TunnelTlvBgpSrteSegmentListIter) appendHolderSlice(item BgpSrteSegmentList) BgpSrteV4TunnelTlvBgpSrteSegmentListIter {
	obj.bgpSrteSegmentListSlice = append(obj.bgpSrteSegmentListSlice, item)
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *bgpSrteV4TunnelTlv) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the BgpSrteV4TunnelTlv object
func (obj *bgpSrteV4TunnelTlv) SetName(value string) BgpSrteV4TunnelTlv {

	obj.obj.Name = &value
	return obj
}

// If enabled means that this part of the configuration including any active 'children' nodes will be advertised to peer.  If disabled, this means that though config is present, it is not taking any part of the test but can be activated at run-time to advertise just this part of the configuration to the peer.
// Active returns a bool
func (obj *bgpSrteV4TunnelTlv) Active() bool {

	return *obj.obj.Active

}

// If enabled means that this part of the configuration including any active 'children' nodes will be advertised to peer.  If disabled, this means that though config is present, it is not taking any part of the test but can be activated at run-time to advertise just this part of the configuration to the peer.
// Active returns a bool
func (obj *bgpSrteV4TunnelTlv) HasActive() bool {
	return obj.obj.Active != nil
}

// If enabled means that this part of the configuration including any active 'children' nodes will be advertised to peer.  If disabled, this means that though config is present, it is not taking any part of the test but can be activated at run-time to advertise just this part of the configuration to the peer.
// SetActive sets the bool value in the BgpSrteV4TunnelTlv object
func (obj *bgpSrteV4TunnelTlv) SetActive(value bool) BgpSrteV4TunnelTlv {

	obj.obj.Active = &value
	return obj
}

func (obj *bgpSrteV4TunnelTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RemoteEndpointSubTlv != nil {

		obj.RemoteEndpointSubTlv().validateObj(vObj, set_default)
	}

	if obj.obj.ColorSubTlv != nil {

		obj.ColorSubTlv().validateObj(vObj, set_default)
	}

	if obj.obj.BindingSubTlv != nil {

		obj.BindingSubTlv().validateObj(vObj, set_default)
	}

	if obj.obj.PreferenceSubTlv != nil {

		obj.PreferenceSubTlv().validateObj(vObj, set_default)
	}

	if obj.obj.PolicyPrioritySubTlv != nil {

		obj.PolicyPrioritySubTlv().validateObj(vObj, set_default)
	}

	if obj.obj.PolicyNameSubTlv != nil {

		obj.PolicyNameSubTlv().validateObj(vObj, set_default)
	}

	if obj.obj.ExplicitNullLabelPolicySubTlv != nil {

		obj.ExplicitNullLabelPolicySubTlv().validateObj(vObj, set_default)
	}

	if len(obj.obj.SegmentLists) != 0 {

		if set_default {
			obj.SegmentLists().clearHolderSlice()
			for _, item := range obj.obj.SegmentLists {
				obj.SegmentLists().appendHolderSlice(&bgpSrteSegmentList{obj: item})
			}
		}
		for _, item := range obj.SegmentLists().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface BgpSrteV4TunnelTlv")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"BgpSrteV4TunnelTlv.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

}

func (obj *bgpSrteV4TunnelTlv) setDefault() {
	if obj.obj.Active == nil {
		obj.SetActive(true)
	}

}

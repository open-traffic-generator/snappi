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

// ***** BgpSrteV6TunnelTlv *****
type bgpSrteV6TunnelTlv struct {
	validation
	obj                                 *otg.BgpSrteV6TunnelTlv
	marshaller                          marshalBgpSrteV6TunnelTlv
	unMarshaller                        unMarshalBgpSrteV6TunnelTlv
	remoteEndpointSubTlvHolder          BgpSrteRemoteEndpointSubTlv
	colorSubTlvHolder                   BgpSrteColorSubTlv
	bindingSubTlvHolder                 BgpSrteBindingSubTlv
	preferenceSubTlvHolder              BgpSrtePreferenceSubTlv
	policyPrioritySubTlvHolder          BgpSrtePolicyPrioritySubTlv
	policyNameSubTlvHolder              BgpSrtePolicyNameSubTlv
	explicitNullLabelPolicySubTlvHolder BgpSrteExplicitNullLabelPolicySubTlv
	segmentListsHolder                  BgpSrteV6TunnelTlvBgpSrteSegmentListIter
}

func NewBgpSrteV6TunnelTlv() BgpSrteV6TunnelTlv {
	obj := bgpSrteV6TunnelTlv{obj: &otg.BgpSrteV6TunnelTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteV6TunnelTlv) msg() *otg.BgpSrteV6TunnelTlv {
	return obj.obj
}

func (obj *bgpSrteV6TunnelTlv) setMsg(msg *otg.BgpSrteV6TunnelTlv) BgpSrteV6TunnelTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteV6TunnelTlv struct {
	obj *bgpSrteV6TunnelTlv
}

type marshalBgpSrteV6TunnelTlv interface {
	// ToProto marshals BgpSrteV6TunnelTlv to protobuf object *otg.BgpSrteV6TunnelTlv
	ToProto() (*otg.BgpSrteV6TunnelTlv, error)
	// ToPbText marshals BgpSrteV6TunnelTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteV6TunnelTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteV6TunnelTlv to JSON text
	ToJson() (string, error)
}

type unMarshalbgpSrteV6TunnelTlv struct {
	obj *bgpSrteV6TunnelTlv
}

type unMarshalBgpSrteV6TunnelTlv interface {
	// FromProto unmarshals BgpSrteV6TunnelTlv from protobuf object *otg.BgpSrteV6TunnelTlv
	FromProto(msg *otg.BgpSrteV6TunnelTlv) (BgpSrteV6TunnelTlv, error)
	// FromPbText unmarshals BgpSrteV6TunnelTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteV6TunnelTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteV6TunnelTlv from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteV6TunnelTlv) Marshal() marshalBgpSrteV6TunnelTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteV6TunnelTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteV6TunnelTlv) Unmarshal() unMarshalBgpSrteV6TunnelTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteV6TunnelTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteV6TunnelTlv) ToProto() (*otg.BgpSrteV6TunnelTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteV6TunnelTlv) FromProto(msg *otg.BgpSrteV6TunnelTlv) (BgpSrteV6TunnelTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteV6TunnelTlv) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteV6TunnelTlv) FromPbText(value string) error {
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

func (m *marshalbgpSrteV6TunnelTlv) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteV6TunnelTlv) FromYaml(value string) error {
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

func (m *marshalbgpSrteV6TunnelTlv) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteV6TunnelTlv) FromJson(value string) error {
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

func (obj *bgpSrteV6TunnelTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteV6TunnelTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteV6TunnelTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteV6TunnelTlv) Clone() (BgpSrteV6TunnelTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteV6TunnelTlv()
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

func (obj *bgpSrteV6TunnelTlv) setNil() {
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

// BgpSrteV6TunnelTlv is configuration for BGP SRTE Tunnel TLV.
type BgpSrteV6TunnelTlv interface {
	Validation
	// msg marshals BgpSrteV6TunnelTlv to protobuf object *otg.BgpSrteV6TunnelTlv
	// and doesn't set defaults
	msg() *otg.BgpSrteV6TunnelTlv
	// setMsg unmarshals BgpSrteV6TunnelTlv from protobuf object *otg.BgpSrteV6TunnelTlv
	// and doesn't set defaults
	setMsg(*otg.BgpSrteV6TunnelTlv) BgpSrteV6TunnelTlv
	// provides marshal interface
	Marshal() marshalBgpSrteV6TunnelTlv
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteV6TunnelTlv
	// validate validates BgpSrteV6TunnelTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteV6TunnelTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RemoteEndpointSubTlv returns BgpSrteRemoteEndpointSubTlv, set in BgpSrteV6TunnelTlv.
	// BgpSrteRemoteEndpointSubTlv is configuration for the BGP remote endpoint sub TLV.
	RemoteEndpointSubTlv() BgpSrteRemoteEndpointSubTlv
	// SetRemoteEndpointSubTlv assigns BgpSrteRemoteEndpointSubTlv provided by user to BgpSrteV6TunnelTlv.
	// BgpSrteRemoteEndpointSubTlv is configuration for the BGP remote endpoint sub TLV.
	SetRemoteEndpointSubTlv(value BgpSrteRemoteEndpointSubTlv) BgpSrteV6TunnelTlv
	// HasRemoteEndpointSubTlv checks if RemoteEndpointSubTlv has been set in BgpSrteV6TunnelTlv
	HasRemoteEndpointSubTlv() bool
	// ColorSubTlv returns BgpSrteColorSubTlv, set in BgpSrteV6TunnelTlv.
	// BgpSrteColorSubTlv is configuration for the Policy Color attribute sub-TLV. The Color sub-TLV MAY be used as a way to "color" the corresponding Tunnel TLV. The Value field of the sub-TLV is eight octets long and consists of a Color Extended Community. First two octets of its Value field are 0x030b as type and subtype of extended  community. Remaining six octets are are exposed to configure.
	ColorSubTlv() BgpSrteColorSubTlv
	// SetColorSubTlv assigns BgpSrteColorSubTlv provided by user to BgpSrteV6TunnelTlv.
	// BgpSrteColorSubTlv is configuration for the Policy Color attribute sub-TLV. The Color sub-TLV MAY be used as a way to "color" the corresponding Tunnel TLV. The Value field of the sub-TLV is eight octets long and consists of a Color Extended Community. First two octets of its Value field are 0x030b as type and subtype of extended  community. Remaining six octets are are exposed to configure.
	SetColorSubTlv(value BgpSrteColorSubTlv) BgpSrteV6TunnelTlv
	// HasColorSubTlv checks if ColorSubTlv has been set in BgpSrteV6TunnelTlv
	HasColorSubTlv() bool
	// BindingSubTlv returns BgpSrteBindingSubTlv, set in BgpSrteV6TunnelTlv.
	// BgpSrteBindingSubTlv is configuration for the binding SID sub-TLV.  This is used to signal the binding SID related information  of the SR Policy candidate path.
	BindingSubTlv() BgpSrteBindingSubTlv
	// SetBindingSubTlv assigns BgpSrteBindingSubTlv provided by user to BgpSrteV6TunnelTlv.
	// BgpSrteBindingSubTlv is configuration for the binding SID sub-TLV.  This is used to signal the binding SID related information  of the SR Policy candidate path.
	SetBindingSubTlv(value BgpSrteBindingSubTlv) BgpSrteV6TunnelTlv
	// HasBindingSubTlv checks if BindingSubTlv has been set in BgpSrteV6TunnelTlv
	HasBindingSubTlv() bool
	// PreferenceSubTlv returns BgpSrtePreferenceSubTlv, set in BgpSrteV6TunnelTlv.
	// BgpSrtePreferenceSubTlv is configuration for BGP preference sub TLV of the SR Policy candidate path.
	PreferenceSubTlv() BgpSrtePreferenceSubTlv
	// SetPreferenceSubTlv assigns BgpSrtePreferenceSubTlv provided by user to BgpSrteV6TunnelTlv.
	// BgpSrtePreferenceSubTlv is configuration for BGP preference sub TLV of the SR Policy candidate path.
	SetPreferenceSubTlv(value BgpSrtePreferenceSubTlv) BgpSrteV6TunnelTlv
	// HasPreferenceSubTlv checks if PreferenceSubTlv has been set in BgpSrteV6TunnelTlv
	HasPreferenceSubTlv() bool
	// PolicyPrioritySubTlv returns BgpSrtePolicyPrioritySubTlv, set in BgpSrteV6TunnelTlv.
	// BgpSrtePolicyPrioritySubTlv is configuration for the Policy Priority sub-TLV. The Policy Priority to indicate the order in which the SR policies  are re-computed upon topological change.
	PolicyPrioritySubTlv() BgpSrtePolicyPrioritySubTlv
	// SetPolicyPrioritySubTlv assigns BgpSrtePolicyPrioritySubTlv provided by user to BgpSrteV6TunnelTlv.
	// BgpSrtePolicyPrioritySubTlv is configuration for the Policy Priority sub-TLV. The Policy Priority to indicate the order in which the SR policies  are re-computed upon topological change.
	SetPolicyPrioritySubTlv(value BgpSrtePolicyPrioritySubTlv) BgpSrteV6TunnelTlv
	// HasPolicyPrioritySubTlv checks if PolicyPrioritySubTlv has been set in BgpSrteV6TunnelTlv
	HasPolicyPrioritySubTlv() bool
	// PolicyNameSubTlv returns BgpSrtePolicyNameSubTlv, set in BgpSrteV6TunnelTlv.
	// BgpSrtePolicyNameSubTlv is configuration for the Policy Name sub-TLV. The Policy Name sub-TLV is used to attach a symbolic name to the SR Policy candidate path.
	PolicyNameSubTlv() BgpSrtePolicyNameSubTlv
	// SetPolicyNameSubTlv assigns BgpSrtePolicyNameSubTlv provided by user to BgpSrteV6TunnelTlv.
	// BgpSrtePolicyNameSubTlv is configuration for the Policy Name sub-TLV. The Policy Name sub-TLV is used to attach a symbolic name to the SR Policy candidate path.
	SetPolicyNameSubTlv(value BgpSrtePolicyNameSubTlv) BgpSrteV6TunnelTlv
	// HasPolicyNameSubTlv checks if PolicyNameSubTlv has been set in BgpSrteV6TunnelTlv
	HasPolicyNameSubTlv() bool
	// ExplicitNullLabelPolicySubTlv returns BgpSrteExplicitNullLabelPolicySubTlv, set in BgpSrteV6TunnelTlv.
	// BgpSrteExplicitNullLabelPolicySubTlv is configuration for BGP explicit null label policy sub TLV settings.
	ExplicitNullLabelPolicySubTlv() BgpSrteExplicitNullLabelPolicySubTlv
	// SetExplicitNullLabelPolicySubTlv assigns BgpSrteExplicitNullLabelPolicySubTlv provided by user to BgpSrteV6TunnelTlv.
	// BgpSrteExplicitNullLabelPolicySubTlv is configuration for BGP explicit null label policy sub TLV settings.
	SetExplicitNullLabelPolicySubTlv(value BgpSrteExplicitNullLabelPolicySubTlv) BgpSrteV6TunnelTlv
	// HasExplicitNullLabelPolicySubTlv checks if ExplicitNullLabelPolicySubTlv has been set in BgpSrteV6TunnelTlv
	HasExplicitNullLabelPolicySubTlv() bool
	// SegmentLists returns BgpSrteV6TunnelTlvBgpSrteSegmentListIterIter, set in BgpSrteV6TunnelTlv
	SegmentLists() BgpSrteV6TunnelTlvBgpSrteSegmentListIter
	// Name returns string, set in BgpSrteV6TunnelTlv.
	Name() string
	// SetName assigns string provided by user to BgpSrteV6TunnelTlv
	SetName(value string) BgpSrteV6TunnelTlv
	// Active returns bool, set in BgpSrteV6TunnelTlv.
	Active() bool
	// SetActive assigns bool provided by user to BgpSrteV6TunnelTlv
	SetActive(value bool) BgpSrteV6TunnelTlv
	// HasActive checks if Active has been set in BgpSrteV6TunnelTlv
	HasActive() bool
	setNil()
}

// description is TBD
// RemoteEndpointSubTlv returns a BgpSrteRemoteEndpointSubTlv
func (obj *bgpSrteV6TunnelTlv) RemoteEndpointSubTlv() BgpSrteRemoteEndpointSubTlv {
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
func (obj *bgpSrteV6TunnelTlv) HasRemoteEndpointSubTlv() bool {
	return obj.obj.RemoteEndpointSubTlv != nil
}

// description is TBD
// SetRemoteEndpointSubTlv sets the BgpSrteRemoteEndpointSubTlv value in the BgpSrteV6TunnelTlv object
func (obj *bgpSrteV6TunnelTlv) SetRemoteEndpointSubTlv(value BgpSrteRemoteEndpointSubTlv) BgpSrteV6TunnelTlv {

	obj.remoteEndpointSubTlvHolder = nil
	obj.obj.RemoteEndpointSubTlv = value.msg()

	return obj
}

// description is TBD
// ColorSubTlv returns a BgpSrteColorSubTlv
func (obj *bgpSrteV6TunnelTlv) ColorSubTlv() BgpSrteColorSubTlv {
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
func (obj *bgpSrteV6TunnelTlv) HasColorSubTlv() bool {
	return obj.obj.ColorSubTlv != nil
}

// description is TBD
// SetColorSubTlv sets the BgpSrteColorSubTlv value in the BgpSrteV6TunnelTlv object
func (obj *bgpSrteV6TunnelTlv) SetColorSubTlv(value BgpSrteColorSubTlv) BgpSrteV6TunnelTlv {

	obj.colorSubTlvHolder = nil
	obj.obj.ColorSubTlv = value.msg()

	return obj
}

// description is TBD
// BindingSubTlv returns a BgpSrteBindingSubTlv
func (obj *bgpSrteV6TunnelTlv) BindingSubTlv() BgpSrteBindingSubTlv {
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
func (obj *bgpSrteV6TunnelTlv) HasBindingSubTlv() bool {
	return obj.obj.BindingSubTlv != nil
}

// description is TBD
// SetBindingSubTlv sets the BgpSrteBindingSubTlv value in the BgpSrteV6TunnelTlv object
func (obj *bgpSrteV6TunnelTlv) SetBindingSubTlv(value BgpSrteBindingSubTlv) BgpSrteV6TunnelTlv {

	obj.bindingSubTlvHolder = nil
	obj.obj.BindingSubTlv = value.msg()

	return obj
}

// description is TBD
// PreferenceSubTlv returns a BgpSrtePreferenceSubTlv
func (obj *bgpSrteV6TunnelTlv) PreferenceSubTlv() BgpSrtePreferenceSubTlv {
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
func (obj *bgpSrteV6TunnelTlv) HasPreferenceSubTlv() bool {
	return obj.obj.PreferenceSubTlv != nil
}

// description is TBD
// SetPreferenceSubTlv sets the BgpSrtePreferenceSubTlv value in the BgpSrteV6TunnelTlv object
func (obj *bgpSrteV6TunnelTlv) SetPreferenceSubTlv(value BgpSrtePreferenceSubTlv) BgpSrteV6TunnelTlv {

	obj.preferenceSubTlvHolder = nil
	obj.obj.PreferenceSubTlv = value.msg()

	return obj
}

// description is TBD
// PolicyPrioritySubTlv returns a BgpSrtePolicyPrioritySubTlv
func (obj *bgpSrteV6TunnelTlv) PolicyPrioritySubTlv() BgpSrtePolicyPrioritySubTlv {
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
func (obj *bgpSrteV6TunnelTlv) HasPolicyPrioritySubTlv() bool {
	return obj.obj.PolicyPrioritySubTlv != nil
}

// description is TBD
// SetPolicyPrioritySubTlv sets the BgpSrtePolicyPrioritySubTlv value in the BgpSrteV6TunnelTlv object
func (obj *bgpSrteV6TunnelTlv) SetPolicyPrioritySubTlv(value BgpSrtePolicyPrioritySubTlv) BgpSrteV6TunnelTlv {

	obj.policyPrioritySubTlvHolder = nil
	obj.obj.PolicyPrioritySubTlv = value.msg()

	return obj
}

// description is TBD
// PolicyNameSubTlv returns a BgpSrtePolicyNameSubTlv
func (obj *bgpSrteV6TunnelTlv) PolicyNameSubTlv() BgpSrtePolicyNameSubTlv {
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
func (obj *bgpSrteV6TunnelTlv) HasPolicyNameSubTlv() bool {
	return obj.obj.PolicyNameSubTlv != nil
}

// description is TBD
// SetPolicyNameSubTlv sets the BgpSrtePolicyNameSubTlv value in the BgpSrteV6TunnelTlv object
func (obj *bgpSrteV6TunnelTlv) SetPolicyNameSubTlv(value BgpSrtePolicyNameSubTlv) BgpSrteV6TunnelTlv {

	obj.policyNameSubTlvHolder = nil
	obj.obj.PolicyNameSubTlv = value.msg()

	return obj
}

// description is TBD
// ExplicitNullLabelPolicySubTlv returns a BgpSrteExplicitNullLabelPolicySubTlv
func (obj *bgpSrteV6TunnelTlv) ExplicitNullLabelPolicySubTlv() BgpSrteExplicitNullLabelPolicySubTlv {
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
func (obj *bgpSrteV6TunnelTlv) HasExplicitNullLabelPolicySubTlv() bool {
	return obj.obj.ExplicitNullLabelPolicySubTlv != nil
}

// description is TBD
// SetExplicitNullLabelPolicySubTlv sets the BgpSrteExplicitNullLabelPolicySubTlv value in the BgpSrteV6TunnelTlv object
func (obj *bgpSrteV6TunnelTlv) SetExplicitNullLabelPolicySubTlv(value BgpSrteExplicitNullLabelPolicySubTlv) BgpSrteV6TunnelTlv {

	obj.explicitNullLabelPolicySubTlvHolder = nil
	obj.obj.ExplicitNullLabelPolicySubTlv = value.msg()

	return obj
}

// description is TBD
// SegmentLists returns a []BgpSrteSegmentList
func (obj *bgpSrteV6TunnelTlv) SegmentLists() BgpSrteV6TunnelTlvBgpSrteSegmentListIter {
	if len(obj.obj.SegmentLists) == 0 {
		obj.obj.SegmentLists = []*otg.BgpSrteSegmentList{}
	}
	if obj.segmentListsHolder == nil {
		obj.segmentListsHolder = newBgpSrteV6TunnelTlvBgpSrteSegmentListIter(&obj.obj.SegmentLists).setMsg(obj)
	}
	return obj.segmentListsHolder
}

type bgpSrteV6TunnelTlvBgpSrteSegmentListIter struct {
	obj                     *bgpSrteV6TunnelTlv
	bgpSrteSegmentListSlice []BgpSrteSegmentList
	fieldPtr                *[]*otg.BgpSrteSegmentList
}

func newBgpSrteV6TunnelTlvBgpSrteSegmentListIter(ptr *[]*otg.BgpSrteSegmentList) BgpSrteV6TunnelTlvBgpSrteSegmentListIter {
	return &bgpSrteV6TunnelTlvBgpSrteSegmentListIter{fieldPtr: ptr}
}

type BgpSrteV6TunnelTlvBgpSrteSegmentListIter interface {
	setMsg(*bgpSrteV6TunnelTlv) BgpSrteV6TunnelTlvBgpSrteSegmentListIter
	Items() []BgpSrteSegmentList
	Add() BgpSrteSegmentList
	Append(items ...BgpSrteSegmentList) BgpSrteV6TunnelTlvBgpSrteSegmentListIter
	Set(index int, newObj BgpSrteSegmentList) BgpSrteV6TunnelTlvBgpSrteSegmentListIter
	Clear() BgpSrteV6TunnelTlvBgpSrteSegmentListIter
	clearHolderSlice() BgpSrteV6TunnelTlvBgpSrteSegmentListIter
	appendHolderSlice(item BgpSrteSegmentList) BgpSrteV6TunnelTlvBgpSrteSegmentListIter
}

func (obj *bgpSrteV6TunnelTlvBgpSrteSegmentListIter) setMsg(msg *bgpSrteV6TunnelTlv) BgpSrteV6TunnelTlvBgpSrteSegmentListIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpSrteSegmentList{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpSrteV6TunnelTlvBgpSrteSegmentListIter) Items() []BgpSrteSegmentList {
	return obj.bgpSrteSegmentListSlice
}

func (obj *bgpSrteV6TunnelTlvBgpSrteSegmentListIter) Add() BgpSrteSegmentList {
	newObj := &otg.BgpSrteSegmentList{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpSrteSegmentList{obj: newObj}
	newLibObj.setDefault()
	obj.bgpSrteSegmentListSlice = append(obj.bgpSrteSegmentListSlice, newLibObj)
	return newLibObj
}

func (obj *bgpSrteV6TunnelTlvBgpSrteSegmentListIter) Append(items ...BgpSrteSegmentList) BgpSrteV6TunnelTlvBgpSrteSegmentListIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpSrteSegmentListSlice = append(obj.bgpSrteSegmentListSlice, item)
	}
	return obj
}

func (obj *bgpSrteV6TunnelTlvBgpSrteSegmentListIter) Set(index int, newObj BgpSrteSegmentList) BgpSrteV6TunnelTlvBgpSrteSegmentListIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpSrteSegmentListSlice[index] = newObj
	return obj
}
func (obj *bgpSrteV6TunnelTlvBgpSrteSegmentListIter) Clear() BgpSrteV6TunnelTlvBgpSrteSegmentListIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpSrteSegmentList{}
		obj.bgpSrteSegmentListSlice = []BgpSrteSegmentList{}
	}
	return obj
}
func (obj *bgpSrteV6TunnelTlvBgpSrteSegmentListIter) clearHolderSlice() BgpSrteV6TunnelTlvBgpSrteSegmentListIter {
	if len(obj.bgpSrteSegmentListSlice) > 0 {
		obj.bgpSrteSegmentListSlice = []BgpSrteSegmentList{}
	}
	return obj
}
func (obj *bgpSrteV6TunnelTlvBgpSrteSegmentListIter) appendHolderSlice(item BgpSrteSegmentList) BgpSrteV6TunnelTlvBgpSrteSegmentListIter {
	obj.bgpSrteSegmentListSlice = append(obj.bgpSrteSegmentListSlice, item)
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *bgpSrteV6TunnelTlv) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the BgpSrteV6TunnelTlv object
func (obj *bgpSrteV6TunnelTlv) SetName(value string) BgpSrteV6TunnelTlv {

	obj.obj.Name = &value
	return obj
}

// If enabled means that this part of the configuration including any active 'children' nodes will be advertised to peer.  If disabled, this means that though config is present, it is not taking any part of the test but can be activated at run-time to advertise just this part of the configuration to the peer.
// Active returns a bool
func (obj *bgpSrteV6TunnelTlv) Active() bool {

	return *obj.obj.Active

}

// If enabled means that this part of the configuration including any active 'children' nodes will be advertised to peer.  If disabled, this means that though config is present, it is not taking any part of the test but can be activated at run-time to advertise just this part of the configuration to the peer.
// Active returns a bool
func (obj *bgpSrteV6TunnelTlv) HasActive() bool {
	return obj.obj.Active != nil
}

// If enabled means that this part of the configuration including any active 'children' nodes will be advertised to peer.  If disabled, this means that though config is present, it is not taking any part of the test but can be activated at run-time to advertise just this part of the configuration to the peer.
// SetActive sets the bool value in the BgpSrteV6TunnelTlv object
func (obj *bgpSrteV6TunnelTlv) SetActive(value bool) BgpSrteV6TunnelTlv {

	obj.obj.Active = &value
	return obj
}

func (obj *bgpSrteV6TunnelTlv) validateObj(vObj *validation, set_default bool) {
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
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface BgpSrteV6TunnelTlv")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"BgpSrteV6TunnelTlv.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

}

func (obj *bgpSrteV6TunnelTlv) setDefault() {
	if obj.obj.Active == nil {
		obj.SetActive(true)
	}

}

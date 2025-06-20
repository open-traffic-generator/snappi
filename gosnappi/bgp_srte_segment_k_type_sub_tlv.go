package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrteSegmentKTypeSubTlv *****
type bgpSrteSegmentKTypeSubTlv struct {
	validation
	obj                           *otg.BgpSrteSegmentKTypeSubTlv
	marshaller                    marshalBgpSrteSegmentKTypeSubTlv
	unMarshaller                  unMarshalBgpSrteSegmentKTypeSubTlv
	srv6SidEndpointBehaviorHolder BgpSrteSRv6SIDEndpointBehaviorAndStructure
}

func NewBgpSrteSegmentKTypeSubTlv() BgpSrteSegmentKTypeSubTlv {
	obj := bgpSrteSegmentKTypeSubTlv{obj: &otg.BgpSrteSegmentKTypeSubTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteSegmentKTypeSubTlv) msg() *otg.BgpSrteSegmentKTypeSubTlv {
	return obj.obj
}

func (obj *bgpSrteSegmentKTypeSubTlv) setMsg(msg *otg.BgpSrteSegmentKTypeSubTlv) BgpSrteSegmentKTypeSubTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteSegmentKTypeSubTlv struct {
	obj *bgpSrteSegmentKTypeSubTlv
}

type marshalBgpSrteSegmentKTypeSubTlv interface {
	// ToProto marshals BgpSrteSegmentKTypeSubTlv to protobuf object *otg.BgpSrteSegmentKTypeSubTlv
	ToProto() (*otg.BgpSrteSegmentKTypeSubTlv, error)
	// ToPbText marshals BgpSrteSegmentKTypeSubTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteSegmentKTypeSubTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteSegmentKTypeSubTlv to JSON text
	ToJson() (string, error)
}

type unMarshalbgpSrteSegmentKTypeSubTlv struct {
	obj *bgpSrteSegmentKTypeSubTlv
}

type unMarshalBgpSrteSegmentKTypeSubTlv interface {
	// FromProto unmarshals BgpSrteSegmentKTypeSubTlv from protobuf object *otg.BgpSrteSegmentKTypeSubTlv
	FromProto(msg *otg.BgpSrteSegmentKTypeSubTlv) (BgpSrteSegmentKTypeSubTlv, error)
	// FromPbText unmarshals BgpSrteSegmentKTypeSubTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteSegmentKTypeSubTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteSegmentKTypeSubTlv from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteSegmentKTypeSubTlv) Marshal() marshalBgpSrteSegmentKTypeSubTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteSegmentKTypeSubTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteSegmentKTypeSubTlv) Unmarshal() unMarshalBgpSrteSegmentKTypeSubTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteSegmentKTypeSubTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteSegmentKTypeSubTlv) ToProto() (*otg.BgpSrteSegmentKTypeSubTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteSegmentKTypeSubTlv) FromProto(msg *otg.BgpSrteSegmentKTypeSubTlv) (BgpSrteSegmentKTypeSubTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteSegmentKTypeSubTlv) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteSegmentKTypeSubTlv) FromPbText(value string) error {
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

func (m *marshalbgpSrteSegmentKTypeSubTlv) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteSegmentKTypeSubTlv) FromYaml(value string) error {
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

func (m *marshalbgpSrteSegmentKTypeSubTlv) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteSegmentKTypeSubTlv) FromJson(value string) error {
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

func (obj *bgpSrteSegmentKTypeSubTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentKTypeSubTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentKTypeSubTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteSegmentKTypeSubTlv) Clone() (BgpSrteSegmentKTypeSubTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteSegmentKTypeSubTlv()
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

func (obj *bgpSrteSegmentKTypeSubTlv) setNil() {
	obj.srv6SidEndpointBehaviorHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpSrteSegmentKTypeSubTlv is type K: IPv6 Local and Remote addresses for SRv6 with optional SID.
type BgpSrteSegmentKTypeSubTlv interface {
	Validation
	// msg marshals BgpSrteSegmentKTypeSubTlv to protobuf object *otg.BgpSrteSegmentKTypeSubTlv
	// and doesn't set defaults
	msg() *otg.BgpSrteSegmentKTypeSubTlv
	// setMsg unmarshals BgpSrteSegmentKTypeSubTlv from protobuf object *otg.BgpSrteSegmentKTypeSubTlv
	// and doesn't set defaults
	setMsg(*otg.BgpSrteSegmentKTypeSubTlv) BgpSrteSegmentKTypeSubTlv
	// provides marshal interface
	Marshal() marshalBgpSrteSegmentKTypeSubTlv
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteSegmentKTypeSubTlv
	// validate validates BgpSrteSegmentKTypeSubTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteSegmentKTypeSubTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns string, set in BgpSrteSegmentKTypeSubTlv.
	Flags() string
	// SetFlags assigns string provided by user to BgpSrteSegmentKTypeSubTlv
	SetFlags(value string) BgpSrteSegmentKTypeSubTlv
	// HasFlags checks if Flags has been set in BgpSrteSegmentKTypeSubTlv
	HasFlags() bool
	// SrAlgorithm returns uint32, set in BgpSrteSegmentKTypeSubTlv.
	SrAlgorithm() uint32
	// SetSrAlgorithm assigns uint32 provided by user to BgpSrteSegmentKTypeSubTlv
	SetSrAlgorithm(value uint32) BgpSrteSegmentKTypeSubTlv
	// HasSrAlgorithm checks if SrAlgorithm has been set in BgpSrteSegmentKTypeSubTlv
	HasSrAlgorithm() bool
	// LocalIpv6Address returns string, set in BgpSrteSegmentKTypeSubTlv.
	LocalIpv6Address() string
	// SetLocalIpv6Address assigns string provided by user to BgpSrteSegmentKTypeSubTlv
	SetLocalIpv6Address(value string) BgpSrteSegmentKTypeSubTlv
	// RemoteIpv6Address returns string, set in BgpSrteSegmentKTypeSubTlv.
	RemoteIpv6Address() string
	// SetRemoteIpv6Address assigns string provided by user to BgpSrteSegmentKTypeSubTlv
	SetRemoteIpv6Address(value string) BgpSrteSegmentKTypeSubTlv
	// Srv6Sid returns string, set in BgpSrteSegmentKTypeSubTlv.
	Srv6Sid() string
	// SetSrv6Sid assigns string provided by user to BgpSrteSegmentKTypeSubTlv
	SetSrv6Sid(value string) BgpSrteSegmentKTypeSubTlv
	// HasSrv6Sid checks if Srv6Sid has been set in BgpSrteSegmentKTypeSubTlv
	HasSrv6Sid() bool
	// Srv6SidEndpointBehavior returns BgpSrteSRv6SIDEndpointBehaviorAndStructure, set in BgpSrteSegmentKTypeSubTlv.
	// BgpSrteSRv6SIDEndpointBehaviorAndStructure is configuration for SRv6 Endpoint Behavior and SID Structure.  Its optional. Summation of lengths for Locator Block, Locator Node,  Function, and Argument MUST be less than or equal to 128.
	Srv6SidEndpointBehavior() BgpSrteSRv6SIDEndpointBehaviorAndStructure
	// SetSrv6SidEndpointBehavior assigns BgpSrteSRv6SIDEndpointBehaviorAndStructure provided by user to BgpSrteSegmentKTypeSubTlv.
	// BgpSrteSRv6SIDEndpointBehaviorAndStructure is configuration for SRv6 Endpoint Behavior and SID Structure.  Its optional. Summation of lengths for Locator Block, Locator Node,  Function, and Argument MUST be less than or equal to 128.
	SetSrv6SidEndpointBehavior(value BgpSrteSRv6SIDEndpointBehaviorAndStructure) BgpSrteSegmentKTypeSubTlv
	// HasSrv6SidEndpointBehavior checks if Srv6SidEndpointBehavior has been set in BgpSrteSegmentKTypeSubTlv
	HasSrv6SidEndpointBehavior() bool
	setNil()
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentKTypeSubTlv) Flags() string {

	return *obj.obj.Flags

}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentKTypeSubTlv) HasFlags() bool {
	return obj.obj.Flags != nil
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// SetFlags sets the string value in the BgpSrteSegmentKTypeSubTlv object
func (obj *bgpSrteSegmentKTypeSubTlv) SetFlags(value string) BgpSrteSegmentKTypeSubTlv {

	obj.obj.Flags = &value
	return obj
}

// SR Algorithm identifier when A-Flag in on.
// SrAlgorithm returns a uint32
func (obj *bgpSrteSegmentKTypeSubTlv) SrAlgorithm() uint32 {

	return *obj.obj.SrAlgorithm

}

// SR Algorithm identifier when A-Flag in on.
// SrAlgorithm returns a uint32
func (obj *bgpSrteSegmentKTypeSubTlv) HasSrAlgorithm() bool {
	return obj.obj.SrAlgorithm != nil
}

// SR Algorithm identifier when A-Flag in on.
// SetSrAlgorithm sets the uint32 value in the BgpSrteSegmentKTypeSubTlv object
func (obj *bgpSrteSegmentKTypeSubTlv) SetSrAlgorithm(value uint32) BgpSrteSegmentKTypeSubTlv {

	obj.obj.SrAlgorithm = &value
	return obj
}

// IPv6 address representing a node.
// LocalIpv6Address returns a string
func (obj *bgpSrteSegmentKTypeSubTlv) LocalIpv6Address() string {

	return *obj.obj.LocalIpv6Address

}

// IPv6 address representing a node.
// SetLocalIpv6Address sets the string value in the BgpSrteSegmentKTypeSubTlv object
func (obj *bgpSrteSegmentKTypeSubTlv) SetLocalIpv6Address(value string) BgpSrteSegmentKTypeSubTlv {

	obj.obj.LocalIpv6Address = &value
	return obj
}

// IPv6 address representing a node.
// RemoteIpv6Address returns a string
func (obj *bgpSrteSegmentKTypeSubTlv) RemoteIpv6Address() string {

	return *obj.obj.RemoteIpv6Address

}

// IPv6 address representing a node.
// SetRemoteIpv6Address sets the string value in the BgpSrteSegmentKTypeSubTlv object
func (obj *bgpSrteSegmentKTypeSubTlv) SetRemoteIpv6Address(value string) BgpSrteSegmentKTypeSubTlv {

	obj.obj.RemoteIpv6Address = &value
	return obj
}

// Optional SRv6 SID.
// Srv6Sid returns a string
func (obj *bgpSrteSegmentKTypeSubTlv) Srv6Sid() string {

	return *obj.obj.Srv6Sid

}

// Optional SRv6 SID.
// Srv6Sid returns a string
func (obj *bgpSrteSegmentKTypeSubTlv) HasSrv6Sid() bool {
	return obj.obj.Srv6Sid != nil
}

// Optional SRv6 SID.
// SetSrv6Sid sets the string value in the BgpSrteSegmentKTypeSubTlv object
func (obj *bgpSrteSegmentKTypeSubTlv) SetSrv6Sid(value string) BgpSrteSegmentKTypeSubTlv {

	obj.obj.Srv6Sid = &value
	return obj
}

// Optional SRv6 Endpoint Behavior and SID Structure.
// Srv6SidEndpointBehavior returns a BgpSrteSRv6SIDEndpointBehaviorAndStructure
func (obj *bgpSrteSegmentKTypeSubTlv) Srv6SidEndpointBehavior() BgpSrteSRv6SIDEndpointBehaviorAndStructure {
	if obj.obj.Srv6SidEndpointBehavior == nil {
		obj.obj.Srv6SidEndpointBehavior = NewBgpSrteSRv6SIDEndpointBehaviorAndStructure().msg()
	}
	if obj.srv6SidEndpointBehaviorHolder == nil {
		obj.srv6SidEndpointBehaviorHolder = &bgpSrteSRv6SIDEndpointBehaviorAndStructure{obj: obj.obj.Srv6SidEndpointBehavior}
	}
	return obj.srv6SidEndpointBehaviorHolder
}

// Optional SRv6 Endpoint Behavior and SID Structure.
// Srv6SidEndpointBehavior returns a BgpSrteSRv6SIDEndpointBehaviorAndStructure
func (obj *bgpSrteSegmentKTypeSubTlv) HasSrv6SidEndpointBehavior() bool {
	return obj.obj.Srv6SidEndpointBehavior != nil
}

// Optional SRv6 Endpoint Behavior and SID Structure.
// SetSrv6SidEndpointBehavior sets the BgpSrteSRv6SIDEndpointBehaviorAndStructure value in the BgpSrteSegmentKTypeSubTlv object
func (obj *bgpSrteSegmentKTypeSubTlv) SetSrv6SidEndpointBehavior(value BgpSrteSRv6SIDEndpointBehaviorAndStructure) BgpSrteSegmentKTypeSubTlv {

	obj.srv6SidEndpointBehaviorHolder = nil
	obj.obj.Srv6SidEndpointBehavior = value.msg()

	return obj
}

func (obj *bgpSrteSegmentKTypeSubTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentKTypeSubTlv.Flags"))
		}

	}

	// LocalIpv6Address is required
	if obj.obj.LocalIpv6Address == nil {
		vObj.validationErrors = append(vObj.validationErrors, "LocalIpv6Address is required field on interface BgpSrteSegmentKTypeSubTlv")
	}
	if obj.obj.LocalIpv6Address != nil {

		err := obj.validateIpv6(obj.LocalIpv6Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentKTypeSubTlv.LocalIpv6Address"))
		}

	}

	// RemoteIpv6Address is required
	if obj.obj.RemoteIpv6Address == nil {
		vObj.validationErrors = append(vObj.validationErrors, "RemoteIpv6Address is required field on interface BgpSrteSegmentKTypeSubTlv")
	}
	if obj.obj.RemoteIpv6Address != nil {

		err := obj.validateIpv6(obj.RemoteIpv6Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentKTypeSubTlv.RemoteIpv6Address"))
		}

	}

	if obj.obj.Srv6Sid != nil {

		err := obj.validateIpv6(obj.Srv6Sid())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentKTypeSubTlv.Srv6Sid"))
		}

	}

	if obj.obj.Srv6SidEndpointBehavior != nil {

		obj.Srv6SidEndpointBehavior().validateObj(vObj, set_default)
	}

}

func (obj *bgpSrteSegmentKTypeSubTlv) setDefault() {
	if obj.obj.SrAlgorithm == nil {
		obj.SetSrAlgorithm(0)
	}

}

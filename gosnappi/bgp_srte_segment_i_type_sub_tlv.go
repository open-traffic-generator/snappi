package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrteSegmentITypeSubTlv *****
type bgpSrteSegmentITypeSubTlv struct {
	validation
	obj                           *otg.BgpSrteSegmentITypeSubTlv
	marshaller                    marshalBgpSrteSegmentITypeSubTlv
	unMarshaller                  unMarshalBgpSrteSegmentITypeSubTlv
	srv6SidEndpointBehaviorHolder BgpSrteSRv6SIDEndpointBehaviorAndStructure
}

func NewBgpSrteSegmentITypeSubTlv() BgpSrteSegmentITypeSubTlv {
	obj := bgpSrteSegmentITypeSubTlv{obj: &otg.BgpSrteSegmentITypeSubTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteSegmentITypeSubTlv) msg() *otg.BgpSrteSegmentITypeSubTlv {
	return obj.obj
}

func (obj *bgpSrteSegmentITypeSubTlv) setMsg(msg *otg.BgpSrteSegmentITypeSubTlv) BgpSrteSegmentITypeSubTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteSegmentITypeSubTlv struct {
	obj *bgpSrteSegmentITypeSubTlv
}

type marshalBgpSrteSegmentITypeSubTlv interface {
	// ToProto marshals BgpSrteSegmentITypeSubTlv to protobuf object *otg.BgpSrteSegmentITypeSubTlv
	ToProto() (*otg.BgpSrteSegmentITypeSubTlv, error)
	// ToPbText marshals BgpSrteSegmentITypeSubTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteSegmentITypeSubTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteSegmentITypeSubTlv to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpSrteSegmentITypeSubTlv to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpSrteSegmentITypeSubTlv struct {
	obj *bgpSrteSegmentITypeSubTlv
}

type unMarshalBgpSrteSegmentITypeSubTlv interface {
	// FromProto unmarshals BgpSrteSegmentITypeSubTlv from protobuf object *otg.BgpSrteSegmentITypeSubTlv
	FromProto(msg *otg.BgpSrteSegmentITypeSubTlv) (BgpSrteSegmentITypeSubTlv, error)
	// FromPbText unmarshals BgpSrteSegmentITypeSubTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteSegmentITypeSubTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteSegmentITypeSubTlv from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteSegmentITypeSubTlv) Marshal() marshalBgpSrteSegmentITypeSubTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteSegmentITypeSubTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteSegmentITypeSubTlv) Unmarshal() unMarshalBgpSrteSegmentITypeSubTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteSegmentITypeSubTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteSegmentITypeSubTlv) ToProto() (*otg.BgpSrteSegmentITypeSubTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteSegmentITypeSubTlv) FromProto(msg *otg.BgpSrteSegmentITypeSubTlv) (BgpSrteSegmentITypeSubTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteSegmentITypeSubTlv) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteSegmentITypeSubTlv) FromPbText(value string) error {
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

func (m *marshalbgpSrteSegmentITypeSubTlv) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteSegmentITypeSubTlv) FromYaml(value string) error {
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

func (m *marshalbgpSrteSegmentITypeSubTlv) ToJsonRaw() (string, error) {
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

func (m *marshalbgpSrteSegmentITypeSubTlv) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteSegmentITypeSubTlv) FromJson(value string) error {
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

func (obj *bgpSrteSegmentITypeSubTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentITypeSubTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentITypeSubTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteSegmentITypeSubTlv) Clone() (BgpSrteSegmentITypeSubTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteSegmentITypeSubTlv()
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

func (obj *bgpSrteSegmentITypeSubTlv) setNil() {
	obj.srv6SidEndpointBehaviorHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpSrteSegmentITypeSubTlv is type I: IPv6 Node Address with optional SRv6 SID.
type BgpSrteSegmentITypeSubTlv interface {
	Validation
	// msg marshals BgpSrteSegmentITypeSubTlv to protobuf object *otg.BgpSrteSegmentITypeSubTlv
	// and doesn't set defaults
	msg() *otg.BgpSrteSegmentITypeSubTlv
	// setMsg unmarshals BgpSrteSegmentITypeSubTlv from protobuf object *otg.BgpSrteSegmentITypeSubTlv
	// and doesn't set defaults
	setMsg(*otg.BgpSrteSegmentITypeSubTlv) BgpSrteSegmentITypeSubTlv
	// provides marshal interface
	Marshal() marshalBgpSrteSegmentITypeSubTlv
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteSegmentITypeSubTlv
	// validate validates BgpSrteSegmentITypeSubTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteSegmentITypeSubTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns string, set in BgpSrteSegmentITypeSubTlv.
	Flags() string
	// SetFlags assigns string provided by user to BgpSrteSegmentITypeSubTlv
	SetFlags(value string) BgpSrteSegmentITypeSubTlv
	// HasFlags checks if Flags has been set in BgpSrteSegmentITypeSubTlv
	HasFlags() bool
	// Ipv6NodeAddress returns string, set in BgpSrteSegmentITypeSubTlv.
	Ipv6NodeAddress() string
	// SetIpv6NodeAddress assigns string provided by user to BgpSrteSegmentITypeSubTlv
	SetIpv6NodeAddress(value string) BgpSrteSegmentITypeSubTlv
	// Srv6Sid returns string, set in BgpSrteSegmentITypeSubTlv.
	Srv6Sid() string
	// SetSrv6Sid assigns string provided by user to BgpSrteSegmentITypeSubTlv
	SetSrv6Sid(value string) BgpSrteSegmentITypeSubTlv
	// HasSrv6Sid checks if Srv6Sid has been set in BgpSrteSegmentITypeSubTlv
	HasSrv6Sid() bool
	// Srv6SidEndpointBehavior returns BgpSrteSRv6SIDEndpointBehaviorAndStructure, set in BgpSrteSegmentITypeSubTlv.
	// BgpSrteSRv6SIDEndpointBehaviorAndStructure is configuration for SRv6 Endpoint Behavior and SID Structure.  Its optional. Summation of lengths for Locator Block, Locator Node,  Function, and Argument MUST be less than or equal to 128.
	Srv6SidEndpointBehavior() BgpSrteSRv6SIDEndpointBehaviorAndStructure
	// SetSrv6SidEndpointBehavior assigns BgpSrteSRv6SIDEndpointBehaviorAndStructure provided by user to BgpSrteSegmentITypeSubTlv.
	// BgpSrteSRv6SIDEndpointBehaviorAndStructure is configuration for SRv6 Endpoint Behavior and SID Structure.  Its optional. Summation of lengths for Locator Block, Locator Node,  Function, and Argument MUST be less than or equal to 128.
	SetSrv6SidEndpointBehavior(value BgpSrteSRv6SIDEndpointBehaviorAndStructure) BgpSrteSegmentITypeSubTlv
	// HasSrv6SidEndpointBehavior checks if Srv6SidEndpointBehavior has been set in BgpSrteSegmentITypeSubTlv
	HasSrv6SidEndpointBehavior() bool
	setNil()
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentITypeSubTlv) Flags() string {

	return *obj.obj.Flags

}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentITypeSubTlv) HasFlags() bool {
	return obj.obj.Flags != nil
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// SetFlags sets the string value in the BgpSrteSegmentITypeSubTlv object
func (obj *bgpSrteSegmentITypeSubTlv) SetFlags(value string) BgpSrteSegmentITypeSubTlv {

	obj.obj.Flags = &value
	return obj
}

// IPv6 address representing a node.
// Ipv6NodeAddress returns a string
func (obj *bgpSrteSegmentITypeSubTlv) Ipv6NodeAddress() string {

	return *obj.obj.Ipv6NodeAddress

}

// IPv6 address representing a node.
// SetIpv6NodeAddress sets the string value in the BgpSrteSegmentITypeSubTlv object
func (obj *bgpSrteSegmentITypeSubTlv) SetIpv6NodeAddress(value string) BgpSrteSegmentITypeSubTlv {

	obj.obj.Ipv6NodeAddress = &value
	return obj
}

// Optional SRv6 SID.
// Srv6Sid returns a string
func (obj *bgpSrteSegmentITypeSubTlv) Srv6Sid() string {

	return *obj.obj.Srv6Sid

}

// Optional SRv6 SID.
// Srv6Sid returns a string
func (obj *bgpSrteSegmentITypeSubTlv) HasSrv6Sid() bool {
	return obj.obj.Srv6Sid != nil
}

// Optional SRv6 SID.
// SetSrv6Sid sets the string value in the BgpSrteSegmentITypeSubTlv object
func (obj *bgpSrteSegmentITypeSubTlv) SetSrv6Sid(value string) BgpSrteSegmentITypeSubTlv {

	obj.obj.Srv6Sid = &value
	return obj
}

// Optional SRv6 Endpoint Behavior and SID Structure.
// Srv6SidEndpointBehavior returns a BgpSrteSRv6SIDEndpointBehaviorAndStructure
func (obj *bgpSrteSegmentITypeSubTlv) Srv6SidEndpointBehavior() BgpSrteSRv6SIDEndpointBehaviorAndStructure {
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
func (obj *bgpSrteSegmentITypeSubTlv) HasSrv6SidEndpointBehavior() bool {
	return obj.obj.Srv6SidEndpointBehavior != nil
}

// Optional SRv6 Endpoint Behavior and SID Structure.
// SetSrv6SidEndpointBehavior sets the BgpSrteSRv6SIDEndpointBehaviorAndStructure value in the BgpSrteSegmentITypeSubTlv object
func (obj *bgpSrteSegmentITypeSubTlv) SetSrv6SidEndpointBehavior(value BgpSrteSRv6SIDEndpointBehaviorAndStructure) BgpSrteSegmentITypeSubTlv {

	obj.srv6SidEndpointBehaviorHolder = nil
	obj.obj.Srv6SidEndpointBehavior = value.msg()

	return obj
}

func (obj *bgpSrteSegmentITypeSubTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentITypeSubTlv.Flags"))
		}

	}

	// Ipv6NodeAddress is required
	if obj.obj.Ipv6NodeAddress == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv6NodeAddress is required field on interface BgpSrteSegmentITypeSubTlv")
	}
	if obj.obj.Ipv6NodeAddress != nil {

		err := obj.validateIpv6(obj.Ipv6NodeAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentITypeSubTlv.Ipv6NodeAddress"))
		}

	}

	if obj.obj.Srv6Sid != nil {

		err := obj.validateIpv6(obj.Srv6Sid())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentITypeSubTlv.Srv6Sid"))
		}

	}

	if obj.obj.Srv6SidEndpointBehavior != nil {

		obj.Srv6SidEndpointBehavior().validateObj(vObj, set_default)
	}

}

func (obj *bgpSrteSegmentITypeSubTlv) setDefault() {

}

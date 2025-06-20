package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrteSegmentJTypeSubTlv *****
type bgpSrteSegmentJTypeSubTlv struct {
	validation
	obj                           *otg.BgpSrteSegmentJTypeSubTlv
	marshaller                    marshalBgpSrteSegmentJTypeSubTlv
	unMarshaller                  unMarshalBgpSrteSegmentJTypeSubTlv
	srv6SidEndpointBehaviorHolder BgpSrteSRv6SIDEndpointBehaviorAndStructure
}

func NewBgpSrteSegmentJTypeSubTlv() BgpSrteSegmentJTypeSubTlv {
	obj := bgpSrteSegmentJTypeSubTlv{obj: &otg.BgpSrteSegmentJTypeSubTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteSegmentJTypeSubTlv) msg() *otg.BgpSrteSegmentJTypeSubTlv {
	return obj.obj
}

func (obj *bgpSrteSegmentJTypeSubTlv) setMsg(msg *otg.BgpSrteSegmentJTypeSubTlv) BgpSrteSegmentJTypeSubTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteSegmentJTypeSubTlv struct {
	obj *bgpSrteSegmentJTypeSubTlv
}

type marshalBgpSrteSegmentJTypeSubTlv interface {
	// ToProto marshals BgpSrteSegmentJTypeSubTlv to protobuf object *otg.BgpSrteSegmentJTypeSubTlv
	ToProto() (*otg.BgpSrteSegmentJTypeSubTlv, error)
	// ToPbText marshals BgpSrteSegmentJTypeSubTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteSegmentJTypeSubTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteSegmentJTypeSubTlv to JSON text
	ToJson() (string, error)
}

type unMarshalbgpSrteSegmentJTypeSubTlv struct {
	obj *bgpSrteSegmentJTypeSubTlv
}

type unMarshalBgpSrteSegmentJTypeSubTlv interface {
	// FromProto unmarshals BgpSrteSegmentJTypeSubTlv from protobuf object *otg.BgpSrteSegmentJTypeSubTlv
	FromProto(msg *otg.BgpSrteSegmentJTypeSubTlv) (BgpSrteSegmentJTypeSubTlv, error)
	// FromPbText unmarshals BgpSrteSegmentJTypeSubTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteSegmentJTypeSubTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteSegmentJTypeSubTlv from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteSegmentJTypeSubTlv) Marshal() marshalBgpSrteSegmentJTypeSubTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteSegmentJTypeSubTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteSegmentJTypeSubTlv) Unmarshal() unMarshalBgpSrteSegmentJTypeSubTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteSegmentJTypeSubTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteSegmentJTypeSubTlv) ToProto() (*otg.BgpSrteSegmentJTypeSubTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteSegmentJTypeSubTlv) FromProto(msg *otg.BgpSrteSegmentJTypeSubTlv) (BgpSrteSegmentJTypeSubTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteSegmentJTypeSubTlv) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteSegmentJTypeSubTlv) FromPbText(value string) error {
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

func (m *marshalbgpSrteSegmentJTypeSubTlv) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteSegmentJTypeSubTlv) FromYaml(value string) error {
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

func (m *marshalbgpSrteSegmentJTypeSubTlv) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteSegmentJTypeSubTlv) FromJson(value string) error {
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

func (obj *bgpSrteSegmentJTypeSubTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentJTypeSubTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentJTypeSubTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteSegmentJTypeSubTlv) Clone() (BgpSrteSegmentJTypeSubTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteSegmentJTypeSubTlv()
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

func (obj *bgpSrteSegmentJTypeSubTlv) setNil() {
	obj.srv6SidEndpointBehaviorHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpSrteSegmentJTypeSubTlv is type J: IPv6 Address, Interface ID for local and remote pair for SRv6 with optional SID.
type BgpSrteSegmentJTypeSubTlv interface {
	Validation
	// msg marshals BgpSrteSegmentJTypeSubTlv to protobuf object *otg.BgpSrteSegmentJTypeSubTlv
	// and doesn't set defaults
	msg() *otg.BgpSrteSegmentJTypeSubTlv
	// setMsg unmarshals BgpSrteSegmentJTypeSubTlv from protobuf object *otg.BgpSrteSegmentJTypeSubTlv
	// and doesn't set defaults
	setMsg(*otg.BgpSrteSegmentJTypeSubTlv) BgpSrteSegmentJTypeSubTlv
	// provides marshal interface
	Marshal() marshalBgpSrteSegmentJTypeSubTlv
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteSegmentJTypeSubTlv
	// validate validates BgpSrteSegmentJTypeSubTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteSegmentJTypeSubTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns string, set in BgpSrteSegmentJTypeSubTlv.
	Flags() string
	// SetFlags assigns string provided by user to BgpSrteSegmentJTypeSubTlv
	SetFlags(value string) BgpSrteSegmentJTypeSubTlv
	// HasFlags checks if Flags has been set in BgpSrteSegmentJTypeSubTlv
	HasFlags() bool
	// SrAlgorithm returns uint32, set in BgpSrteSegmentJTypeSubTlv.
	SrAlgorithm() uint32
	// SetSrAlgorithm assigns uint32 provided by user to BgpSrteSegmentJTypeSubTlv
	SetSrAlgorithm(value uint32) BgpSrteSegmentJTypeSubTlv
	// HasSrAlgorithm checks if SrAlgorithm has been set in BgpSrteSegmentJTypeSubTlv
	HasSrAlgorithm() bool
	// LocalInterfaceId returns uint32, set in BgpSrteSegmentJTypeSubTlv.
	LocalInterfaceId() uint32
	// SetLocalInterfaceId assigns uint32 provided by user to BgpSrteSegmentJTypeSubTlv
	SetLocalInterfaceId(value uint32) BgpSrteSegmentJTypeSubTlv
	// HasLocalInterfaceId checks if LocalInterfaceId has been set in BgpSrteSegmentJTypeSubTlv
	HasLocalInterfaceId() bool
	// LocalIpv6NodeAddress returns string, set in BgpSrteSegmentJTypeSubTlv.
	LocalIpv6NodeAddress() string
	// SetLocalIpv6NodeAddress assigns string provided by user to BgpSrteSegmentJTypeSubTlv
	SetLocalIpv6NodeAddress(value string) BgpSrteSegmentJTypeSubTlv
	// RemoteInterfaceId returns uint32, set in BgpSrteSegmentJTypeSubTlv.
	RemoteInterfaceId() uint32
	// SetRemoteInterfaceId assigns uint32 provided by user to BgpSrteSegmentJTypeSubTlv
	SetRemoteInterfaceId(value uint32) BgpSrteSegmentJTypeSubTlv
	// HasRemoteInterfaceId checks if RemoteInterfaceId has been set in BgpSrteSegmentJTypeSubTlv
	HasRemoteInterfaceId() bool
	// RemoteIpv6NodeAddress returns string, set in BgpSrteSegmentJTypeSubTlv.
	RemoteIpv6NodeAddress() string
	// SetRemoteIpv6NodeAddress assigns string provided by user to BgpSrteSegmentJTypeSubTlv
	SetRemoteIpv6NodeAddress(value string) BgpSrteSegmentJTypeSubTlv
	// Srv6Sid returns string, set in BgpSrteSegmentJTypeSubTlv.
	Srv6Sid() string
	// SetSrv6Sid assigns string provided by user to BgpSrteSegmentJTypeSubTlv
	SetSrv6Sid(value string) BgpSrteSegmentJTypeSubTlv
	// HasSrv6Sid checks if Srv6Sid has been set in BgpSrteSegmentJTypeSubTlv
	HasSrv6Sid() bool
	// Srv6SidEndpointBehavior returns BgpSrteSRv6SIDEndpointBehaviorAndStructure, set in BgpSrteSegmentJTypeSubTlv.
	// BgpSrteSRv6SIDEndpointBehaviorAndStructure is configuration for SRv6 Endpoint Behavior and SID Structure.  Its optional. Summation of lengths for Locator Block, Locator Node,  Function, and Argument MUST be less than or equal to 128.
	Srv6SidEndpointBehavior() BgpSrteSRv6SIDEndpointBehaviorAndStructure
	// SetSrv6SidEndpointBehavior assigns BgpSrteSRv6SIDEndpointBehaviorAndStructure provided by user to BgpSrteSegmentJTypeSubTlv.
	// BgpSrteSRv6SIDEndpointBehaviorAndStructure is configuration for SRv6 Endpoint Behavior and SID Structure.  Its optional. Summation of lengths for Locator Block, Locator Node,  Function, and Argument MUST be less than or equal to 128.
	SetSrv6SidEndpointBehavior(value BgpSrteSRv6SIDEndpointBehaviorAndStructure) BgpSrteSegmentJTypeSubTlv
	// HasSrv6SidEndpointBehavior checks if Srv6SidEndpointBehavior has been set in BgpSrteSegmentJTypeSubTlv
	HasSrv6SidEndpointBehavior() bool
	setNil()
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentJTypeSubTlv) Flags() string {

	return *obj.obj.Flags

}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentJTypeSubTlv) HasFlags() bool {
	return obj.obj.Flags != nil
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// SetFlags sets the string value in the BgpSrteSegmentJTypeSubTlv object
func (obj *bgpSrteSegmentJTypeSubTlv) SetFlags(value string) BgpSrteSegmentJTypeSubTlv {

	obj.obj.Flags = &value
	return obj
}

// SR Algorithm identifier when A-Flag in on.
// SrAlgorithm returns a uint32
func (obj *bgpSrteSegmentJTypeSubTlv) SrAlgorithm() uint32 {

	return *obj.obj.SrAlgorithm

}

// SR Algorithm identifier when A-Flag in on.
// SrAlgorithm returns a uint32
func (obj *bgpSrteSegmentJTypeSubTlv) HasSrAlgorithm() bool {
	return obj.obj.SrAlgorithm != nil
}

// SR Algorithm identifier when A-Flag in on.
// SetSrAlgorithm sets the uint32 value in the BgpSrteSegmentJTypeSubTlv object
func (obj *bgpSrteSegmentJTypeSubTlv) SetSrAlgorithm(value uint32) BgpSrteSegmentJTypeSubTlv {

	obj.obj.SrAlgorithm = &value
	return obj
}

// Local Interface ID: The Interface Index as defined in [RFC8664].
// LocalInterfaceId returns a uint32
func (obj *bgpSrteSegmentJTypeSubTlv) LocalInterfaceId() uint32 {

	return *obj.obj.LocalInterfaceId

}

// Local Interface ID: The Interface Index as defined in [RFC8664].
// LocalInterfaceId returns a uint32
func (obj *bgpSrteSegmentJTypeSubTlv) HasLocalInterfaceId() bool {
	return obj.obj.LocalInterfaceId != nil
}

// Local Interface ID: The Interface Index as defined in [RFC8664].
// SetLocalInterfaceId sets the uint32 value in the BgpSrteSegmentJTypeSubTlv object
func (obj *bgpSrteSegmentJTypeSubTlv) SetLocalInterfaceId(value uint32) BgpSrteSegmentJTypeSubTlv {

	obj.obj.LocalInterfaceId = &value
	return obj
}

// IPv6 address representing a node.
// LocalIpv6NodeAddress returns a string
func (obj *bgpSrteSegmentJTypeSubTlv) LocalIpv6NodeAddress() string {

	return *obj.obj.LocalIpv6NodeAddress

}

// IPv6 address representing a node.
// SetLocalIpv6NodeAddress sets the string value in the BgpSrteSegmentJTypeSubTlv object
func (obj *bgpSrteSegmentJTypeSubTlv) SetLocalIpv6NodeAddress(value string) BgpSrteSegmentJTypeSubTlv {

	obj.obj.LocalIpv6NodeAddress = &value
	return obj
}

// Local Interface ID: The Interface Index as defined in [RFC8664].
// RemoteInterfaceId returns a uint32
func (obj *bgpSrteSegmentJTypeSubTlv) RemoteInterfaceId() uint32 {

	return *obj.obj.RemoteInterfaceId

}

// Local Interface ID: The Interface Index as defined in [RFC8664].
// RemoteInterfaceId returns a uint32
func (obj *bgpSrteSegmentJTypeSubTlv) HasRemoteInterfaceId() bool {
	return obj.obj.RemoteInterfaceId != nil
}

// Local Interface ID: The Interface Index as defined in [RFC8664].
// SetRemoteInterfaceId sets the uint32 value in the BgpSrteSegmentJTypeSubTlv object
func (obj *bgpSrteSegmentJTypeSubTlv) SetRemoteInterfaceId(value uint32) BgpSrteSegmentJTypeSubTlv {

	obj.obj.RemoteInterfaceId = &value
	return obj
}

// IPv6 address representing a node.
// RemoteIpv6NodeAddress returns a string
func (obj *bgpSrteSegmentJTypeSubTlv) RemoteIpv6NodeAddress() string {

	return *obj.obj.RemoteIpv6NodeAddress

}

// IPv6 address representing a node.
// SetRemoteIpv6NodeAddress sets the string value in the BgpSrteSegmentJTypeSubTlv object
func (obj *bgpSrteSegmentJTypeSubTlv) SetRemoteIpv6NodeAddress(value string) BgpSrteSegmentJTypeSubTlv {

	obj.obj.RemoteIpv6NodeAddress = &value
	return obj
}

// Optional SRv6 SID.
// Srv6Sid returns a string
func (obj *bgpSrteSegmentJTypeSubTlv) Srv6Sid() string {

	return *obj.obj.Srv6Sid

}

// Optional SRv6 SID.
// Srv6Sid returns a string
func (obj *bgpSrteSegmentJTypeSubTlv) HasSrv6Sid() bool {
	return obj.obj.Srv6Sid != nil
}

// Optional SRv6 SID.
// SetSrv6Sid sets the string value in the BgpSrteSegmentJTypeSubTlv object
func (obj *bgpSrteSegmentJTypeSubTlv) SetSrv6Sid(value string) BgpSrteSegmentJTypeSubTlv {

	obj.obj.Srv6Sid = &value
	return obj
}

// Optional SRv6 Endpoint Behavior and SID Structure.
// Srv6SidEndpointBehavior returns a BgpSrteSRv6SIDEndpointBehaviorAndStructure
func (obj *bgpSrteSegmentJTypeSubTlv) Srv6SidEndpointBehavior() BgpSrteSRv6SIDEndpointBehaviorAndStructure {
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
func (obj *bgpSrteSegmentJTypeSubTlv) HasSrv6SidEndpointBehavior() bool {
	return obj.obj.Srv6SidEndpointBehavior != nil
}

// Optional SRv6 Endpoint Behavior and SID Structure.
// SetSrv6SidEndpointBehavior sets the BgpSrteSRv6SIDEndpointBehaviorAndStructure value in the BgpSrteSegmentJTypeSubTlv object
func (obj *bgpSrteSegmentJTypeSubTlv) SetSrv6SidEndpointBehavior(value BgpSrteSRv6SIDEndpointBehaviorAndStructure) BgpSrteSegmentJTypeSubTlv {

	obj.srv6SidEndpointBehaviorHolder = nil
	obj.obj.Srv6SidEndpointBehavior = value.msg()

	return obj
}

func (obj *bgpSrteSegmentJTypeSubTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentJTypeSubTlv.Flags"))
		}

	}

	// LocalIpv6NodeAddress is required
	if obj.obj.LocalIpv6NodeAddress == nil {
		vObj.validationErrors = append(vObj.validationErrors, "LocalIpv6NodeAddress is required field on interface BgpSrteSegmentJTypeSubTlv")
	}
	if obj.obj.LocalIpv6NodeAddress != nil {

		err := obj.validateIpv6(obj.LocalIpv6NodeAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentJTypeSubTlv.LocalIpv6NodeAddress"))
		}

	}

	// RemoteIpv6NodeAddress is required
	if obj.obj.RemoteIpv6NodeAddress == nil {
		vObj.validationErrors = append(vObj.validationErrors, "RemoteIpv6NodeAddress is required field on interface BgpSrteSegmentJTypeSubTlv")
	}
	if obj.obj.RemoteIpv6NodeAddress != nil {

		err := obj.validateIpv6(obj.RemoteIpv6NodeAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentJTypeSubTlv.RemoteIpv6NodeAddress"))
		}

	}

	if obj.obj.Srv6Sid != nil {

		err := obj.validateIpv6(obj.Srv6Sid())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentJTypeSubTlv.Srv6Sid"))
		}

	}

	if obj.obj.Srv6SidEndpointBehavior != nil {

		obj.Srv6SidEndpointBehavior().validateObj(vObj, set_default)
	}

}

func (obj *bgpSrteSegmentJTypeSubTlv) setDefault() {
	if obj.obj.SrAlgorithm == nil {
		obj.SetSrAlgorithm(0)
	}
	if obj.obj.LocalInterfaceId == nil {
		obj.SetLocalInterfaceId(0)
	}
	if obj.obj.RemoteInterfaceId == nil {
		obj.SetRemoteInterfaceId(0)
	}

}

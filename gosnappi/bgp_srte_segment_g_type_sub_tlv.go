package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrteSegmentGTypeSubTlv *****
type bgpSrteSegmentGTypeSubTlv struct {
	validation
	obj             *otg.BgpSrteSegmentGTypeSubTlv
	marshaller      marshalBgpSrteSegmentGTypeSubTlv
	unMarshaller    unMarshalBgpSrteSegmentGTypeSubTlv
	srMplsSidHolder BgpSrteSrMplsSid
}

func NewBgpSrteSegmentGTypeSubTlv() BgpSrteSegmentGTypeSubTlv {
	obj := bgpSrteSegmentGTypeSubTlv{obj: &otg.BgpSrteSegmentGTypeSubTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteSegmentGTypeSubTlv) msg() *otg.BgpSrteSegmentGTypeSubTlv {
	return obj.obj
}

func (obj *bgpSrteSegmentGTypeSubTlv) setMsg(msg *otg.BgpSrteSegmentGTypeSubTlv) BgpSrteSegmentGTypeSubTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteSegmentGTypeSubTlv struct {
	obj *bgpSrteSegmentGTypeSubTlv
}

type marshalBgpSrteSegmentGTypeSubTlv interface {
	// ToProto marshals BgpSrteSegmentGTypeSubTlv to protobuf object *otg.BgpSrteSegmentGTypeSubTlv
	ToProto() (*otg.BgpSrteSegmentGTypeSubTlv, error)
	// ToPbText marshals BgpSrteSegmentGTypeSubTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteSegmentGTypeSubTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteSegmentGTypeSubTlv to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpSrteSegmentGTypeSubTlv to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpSrteSegmentGTypeSubTlv struct {
	obj *bgpSrteSegmentGTypeSubTlv
}

type unMarshalBgpSrteSegmentGTypeSubTlv interface {
	// FromProto unmarshals BgpSrteSegmentGTypeSubTlv from protobuf object *otg.BgpSrteSegmentGTypeSubTlv
	FromProto(msg *otg.BgpSrteSegmentGTypeSubTlv) (BgpSrteSegmentGTypeSubTlv, error)
	// FromPbText unmarshals BgpSrteSegmentGTypeSubTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteSegmentGTypeSubTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteSegmentGTypeSubTlv from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteSegmentGTypeSubTlv) Marshal() marshalBgpSrteSegmentGTypeSubTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteSegmentGTypeSubTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteSegmentGTypeSubTlv) Unmarshal() unMarshalBgpSrteSegmentGTypeSubTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteSegmentGTypeSubTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteSegmentGTypeSubTlv) ToProto() (*otg.BgpSrteSegmentGTypeSubTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteSegmentGTypeSubTlv) FromProto(msg *otg.BgpSrteSegmentGTypeSubTlv) (BgpSrteSegmentGTypeSubTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteSegmentGTypeSubTlv) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteSegmentGTypeSubTlv) FromPbText(value string) error {
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

func (m *marshalbgpSrteSegmentGTypeSubTlv) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteSegmentGTypeSubTlv) FromYaml(value string) error {
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

func (m *marshalbgpSrteSegmentGTypeSubTlv) ToJsonRaw() (string, error) {
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

func (m *marshalbgpSrteSegmentGTypeSubTlv) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteSegmentGTypeSubTlv) FromJson(value string) error {
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

func (obj *bgpSrteSegmentGTypeSubTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentGTypeSubTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentGTypeSubTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteSegmentGTypeSubTlv) Clone() (BgpSrteSegmentGTypeSubTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteSegmentGTypeSubTlv()
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

func (obj *bgpSrteSegmentGTypeSubTlv) setNil() {
	obj.srMplsSidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpSrteSegmentGTypeSubTlv is type G: IPv6 Address, Interface ID for local and remote pair with optional SID for SR MPLS.
type BgpSrteSegmentGTypeSubTlv interface {
	Validation
	// msg marshals BgpSrteSegmentGTypeSubTlv to protobuf object *otg.BgpSrteSegmentGTypeSubTlv
	// and doesn't set defaults
	msg() *otg.BgpSrteSegmentGTypeSubTlv
	// setMsg unmarshals BgpSrteSegmentGTypeSubTlv from protobuf object *otg.BgpSrteSegmentGTypeSubTlv
	// and doesn't set defaults
	setMsg(*otg.BgpSrteSegmentGTypeSubTlv) BgpSrteSegmentGTypeSubTlv
	// provides marshal interface
	Marshal() marshalBgpSrteSegmentGTypeSubTlv
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteSegmentGTypeSubTlv
	// validate validates BgpSrteSegmentGTypeSubTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteSegmentGTypeSubTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns string, set in BgpSrteSegmentGTypeSubTlv.
	Flags() string
	// SetFlags assigns string provided by user to BgpSrteSegmentGTypeSubTlv
	SetFlags(value string) BgpSrteSegmentGTypeSubTlv
	// HasFlags checks if Flags has been set in BgpSrteSegmentGTypeSubTlv
	HasFlags() bool
	// LocalInterfaceId returns uint32, set in BgpSrteSegmentGTypeSubTlv.
	LocalInterfaceId() uint32
	// SetLocalInterfaceId assigns uint32 provided by user to BgpSrteSegmentGTypeSubTlv
	SetLocalInterfaceId(value uint32) BgpSrteSegmentGTypeSubTlv
	// HasLocalInterfaceId checks if LocalInterfaceId has been set in BgpSrteSegmentGTypeSubTlv
	HasLocalInterfaceId() bool
	// LocalIpv6NodeAddress returns string, set in BgpSrteSegmentGTypeSubTlv.
	LocalIpv6NodeAddress() string
	// SetLocalIpv6NodeAddress assigns string provided by user to BgpSrteSegmentGTypeSubTlv
	SetLocalIpv6NodeAddress(value string) BgpSrteSegmentGTypeSubTlv
	// RemoteInterfaceId returns uint32, set in BgpSrteSegmentGTypeSubTlv.
	RemoteInterfaceId() uint32
	// SetRemoteInterfaceId assigns uint32 provided by user to BgpSrteSegmentGTypeSubTlv
	SetRemoteInterfaceId(value uint32) BgpSrteSegmentGTypeSubTlv
	// HasRemoteInterfaceId checks if RemoteInterfaceId has been set in BgpSrteSegmentGTypeSubTlv
	HasRemoteInterfaceId() bool
	// RemoteIpv6NodeAddress returns string, set in BgpSrteSegmentGTypeSubTlv.
	RemoteIpv6NodeAddress() string
	// SetRemoteIpv6NodeAddress assigns string provided by user to BgpSrteSegmentGTypeSubTlv
	SetRemoteIpv6NodeAddress(value string) BgpSrteSegmentGTypeSubTlv
	// SrMplsSid returns BgpSrteSrMplsSid, set in BgpSrteSegmentGTypeSubTlv.
	// BgpSrteSrMplsSid is configuration for SR-MPLS with Label, TC, Bottom-of-Stack and TTL.
	SrMplsSid() BgpSrteSrMplsSid
	// SetSrMplsSid assigns BgpSrteSrMplsSid provided by user to BgpSrteSegmentGTypeSubTlv.
	// BgpSrteSrMplsSid is configuration for SR-MPLS with Label, TC, Bottom-of-Stack and TTL.
	SetSrMplsSid(value BgpSrteSrMplsSid) BgpSrteSegmentGTypeSubTlv
	// HasSrMplsSid checks if SrMplsSid has been set in BgpSrteSegmentGTypeSubTlv
	HasSrMplsSid() bool
	setNil()
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentGTypeSubTlv) Flags() string {

	return *obj.obj.Flags

}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentGTypeSubTlv) HasFlags() bool {
	return obj.obj.Flags != nil
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// SetFlags sets the string value in the BgpSrteSegmentGTypeSubTlv object
func (obj *bgpSrteSegmentGTypeSubTlv) SetFlags(value string) BgpSrteSegmentGTypeSubTlv {

	obj.obj.Flags = &value
	return obj
}

// Local Interface ID: The Interface Index as defined in [RFC8664].
// LocalInterfaceId returns a uint32
func (obj *bgpSrteSegmentGTypeSubTlv) LocalInterfaceId() uint32 {

	return *obj.obj.LocalInterfaceId

}

// Local Interface ID: The Interface Index as defined in [RFC8664].
// LocalInterfaceId returns a uint32
func (obj *bgpSrteSegmentGTypeSubTlv) HasLocalInterfaceId() bool {
	return obj.obj.LocalInterfaceId != nil
}

// Local Interface ID: The Interface Index as defined in [RFC8664].
// SetLocalInterfaceId sets the uint32 value in the BgpSrteSegmentGTypeSubTlv object
func (obj *bgpSrteSegmentGTypeSubTlv) SetLocalInterfaceId(value uint32) BgpSrteSegmentGTypeSubTlv {

	obj.obj.LocalInterfaceId = &value
	return obj
}

// IPv6 address representing a node.
// LocalIpv6NodeAddress returns a string
func (obj *bgpSrteSegmentGTypeSubTlv) LocalIpv6NodeAddress() string {

	return *obj.obj.LocalIpv6NodeAddress

}

// IPv6 address representing a node.
// SetLocalIpv6NodeAddress sets the string value in the BgpSrteSegmentGTypeSubTlv object
func (obj *bgpSrteSegmentGTypeSubTlv) SetLocalIpv6NodeAddress(value string) BgpSrteSegmentGTypeSubTlv {

	obj.obj.LocalIpv6NodeAddress = &value
	return obj
}

// Local Interface ID: The Interface Index as defined in [RFC8664].
// RemoteInterfaceId returns a uint32
func (obj *bgpSrteSegmentGTypeSubTlv) RemoteInterfaceId() uint32 {

	return *obj.obj.RemoteInterfaceId

}

// Local Interface ID: The Interface Index as defined in [RFC8664].
// RemoteInterfaceId returns a uint32
func (obj *bgpSrteSegmentGTypeSubTlv) HasRemoteInterfaceId() bool {
	return obj.obj.RemoteInterfaceId != nil
}

// Local Interface ID: The Interface Index as defined in [RFC8664].
// SetRemoteInterfaceId sets the uint32 value in the BgpSrteSegmentGTypeSubTlv object
func (obj *bgpSrteSegmentGTypeSubTlv) SetRemoteInterfaceId(value uint32) BgpSrteSegmentGTypeSubTlv {

	obj.obj.RemoteInterfaceId = &value
	return obj
}

// IPv6 address representing a node.
// RemoteIpv6NodeAddress returns a string
func (obj *bgpSrteSegmentGTypeSubTlv) RemoteIpv6NodeAddress() string {

	return *obj.obj.RemoteIpv6NodeAddress

}

// IPv6 address representing a node.
// SetRemoteIpv6NodeAddress sets the string value in the BgpSrteSegmentGTypeSubTlv object
func (obj *bgpSrteSegmentGTypeSubTlv) SetRemoteIpv6NodeAddress(value string) BgpSrteSegmentGTypeSubTlv {

	obj.obj.RemoteIpv6NodeAddress = &value
	return obj
}

// Optional SR-MPLS SID.
// SrMplsSid returns a BgpSrteSrMplsSid
func (obj *bgpSrteSegmentGTypeSubTlv) SrMplsSid() BgpSrteSrMplsSid {
	if obj.obj.SrMplsSid == nil {
		obj.obj.SrMplsSid = NewBgpSrteSrMplsSid().msg()
	}
	if obj.srMplsSidHolder == nil {
		obj.srMplsSidHolder = &bgpSrteSrMplsSid{obj: obj.obj.SrMplsSid}
	}
	return obj.srMplsSidHolder
}

// Optional SR-MPLS SID.
// SrMplsSid returns a BgpSrteSrMplsSid
func (obj *bgpSrteSegmentGTypeSubTlv) HasSrMplsSid() bool {
	return obj.obj.SrMplsSid != nil
}

// Optional SR-MPLS SID.
// SetSrMplsSid sets the BgpSrteSrMplsSid value in the BgpSrteSegmentGTypeSubTlv object
func (obj *bgpSrteSegmentGTypeSubTlv) SetSrMplsSid(value BgpSrteSrMplsSid) BgpSrteSegmentGTypeSubTlv {

	obj.srMplsSidHolder = nil
	obj.obj.SrMplsSid = value.msg()

	return obj
}

func (obj *bgpSrteSegmentGTypeSubTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentGTypeSubTlv.Flags"))
		}

	}

	// LocalIpv6NodeAddress is required
	if obj.obj.LocalIpv6NodeAddress == nil {
		vObj.validationErrors = append(vObj.validationErrors, "LocalIpv6NodeAddress is required field on interface BgpSrteSegmentGTypeSubTlv")
	}
	if obj.obj.LocalIpv6NodeAddress != nil {

		err := obj.validateIpv6(obj.LocalIpv6NodeAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentGTypeSubTlv.LocalIpv6NodeAddress"))
		}

	}

	// RemoteIpv6NodeAddress is required
	if obj.obj.RemoteIpv6NodeAddress == nil {
		vObj.validationErrors = append(vObj.validationErrors, "RemoteIpv6NodeAddress is required field on interface BgpSrteSegmentGTypeSubTlv")
	}
	if obj.obj.RemoteIpv6NodeAddress != nil {

		err := obj.validateIpv6(obj.RemoteIpv6NodeAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentGTypeSubTlv.RemoteIpv6NodeAddress"))
		}

	}

	if obj.obj.SrMplsSid != nil {

		obj.SrMplsSid().validateObj(vObj, set_default)
	}

}

func (obj *bgpSrteSegmentGTypeSubTlv) setDefault() {
	if obj.obj.LocalInterfaceId == nil {
		obj.SetLocalInterfaceId(0)
	}
	if obj.obj.RemoteInterfaceId == nil {
		obj.SetRemoteInterfaceId(0)
	}

}

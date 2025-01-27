package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrteSegmentETypeSubTlv *****
type bgpSrteSegmentETypeSubTlv struct {
	validation
	obj             *otg.BgpSrteSegmentETypeSubTlv
	marshaller      marshalBgpSrteSegmentETypeSubTlv
	unMarshaller    unMarshalBgpSrteSegmentETypeSubTlv
	srMplsSidHolder BgpSrteSrMplsSid
}

func NewBgpSrteSegmentETypeSubTlv() BgpSrteSegmentETypeSubTlv {
	obj := bgpSrteSegmentETypeSubTlv{obj: &otg.BgpSrteSegmentETypeSubTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteSegmentETypeSubTlv) msg() *otg.BgpSrteSegmentETypeSubTlv {
	return obj.obj
}

func (obj *bgpSrteSegmentETypeSubTlv) setMsg(msg *otg.BgpSrteSegmentETypeSubTlv) BgpSrteSegmentETypeSubTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteSegmentETypeSubTlv struct {
	obj *bgpSrteSegmentETypeSubTlv
}

type marshalBgpSrteSegmentETypeSubTlv interface {
	// ToProto marshals BgpSrteSegmentETypeSubTlv to protobuf object *otg.BgpSrteSegmentETypeSubTlv
	ToProto() (*otg.BgpSrteSegmentETypeSubTlv, error)
	// ToPbText marshals BgpSrteSegmentETypeSubTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteSegmentETypeSubTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteSegmentETypeSubTlv to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpSrteSegmentETypeSubTlv to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpSrteSegmentETypeSubTlv struct {
	obj *bgpSrteSegmentETypeSubTlv
}

type unMarshalBgpSrteSegmentETypeSubTlv interface {
	// FromProto unmarshals BgpSrteSegmentETypeSubTlv from protobuf object *otg.BgpSrteSegmentETypeSubTlv
	FromProto(msg *otg.BgpSrteSegmentETypeSubTlv) (BgpSrteSegmentETypeSubTlv, error)
	// FromPbText unmarshals BgpSrteSegmentETypeSubTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteSegmentETypeSubTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteSegmentETypeSubTlv from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteSegmentETypeSubTlv) Marshal() marshalBgpSrteSegmentETypeSubTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteSegmentETypeSubTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteSegmentETypeSubTlv) Unmarshal() unMarshalBgpSrteSegmentETypeSubTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteSegmentETypeSubTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteSegmentETypeSubTlv) ToProto() (*otg.BgpSrteSegmentETypeSubTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteSegmentETypeSubTlv) FromProto(msg *otg.BgpSrteSegmentETypeSubTlv) (BgpSrteSegmentETypeSubTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteSegmentETypeSubTlv) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteSegmentETypeSubTlv) FromPbText(value string) error {
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

func (m *marshalbgpSrteSegmentETypeSubTlv) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteSegmentETypeSubTlv) FromYaml(value string) error {
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

func (m *marshalbgpSrteSegmentETypeSubTlv) ToJsonRaw() (string, error) {
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

func (m *marshalbgpSrteSegmentETypeSubTlv) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteSegmentETypeSubTlv) FromJson(value string) error {
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

func (obj *bgpSrteSegmentETypeSubTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentETypeSubTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentETypeSubTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteSegmentETypeSubTlv) Clone() (BgpSrteSegmentETypeSubTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteSegmentETypeSubTlv()
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

func (obj *bgpSrteSegmentETypeSubTlv) setNil() {
	obj.srMplsSidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpSrteSegmentETypeSubTlv is type E: IPv4 Address and Local Interface ID with optional SID
type BgpSrteSegmentETypeSubTlv interface {
	Validation
	// msg marshals BgpSrteSegmentETypeSubTlv to protobuf object *otg.BgpSrteSegmentETypeSubTlv
	// and doesn't set defaults
	msg() *otg.BgpSrteSegmentETypeSubTlv
	// setMsg unmarshals BgpSrteSegmentETypeSubTlv from protobuf object *otg.BgpSrteSegmentETypeSubTlv
	// and doesn't set defaults
	setMsg(*otg.BgpSrteSegmentETypeSubTlv) BgpSrteSegmentETypeSubTlv
	// provides marshal interface
	Marshal() marshalBgpSrteSegmentETypeSubTlv
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteSegmentETypeSubTlv
	// validate validates BgpSrteSegmentETypeSubTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteSegmentETypeSubTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns string, set in BgpSrteSegmentETypeSubTlv.
	Flags() string
	// SetFlags assigns string provided by user to BgpSrteSegmentETypeSubTlv
	SetFlags(value string) BgpSrteSegmentETypeSubTlv
	// HasFlags checks if Flags has been set in BgpSrteSegmentETypeSubTlv
	HasFlags() bool
	// LocalInterfaceId returns uint32, set in BgpSrteSegmentETypeSubTlv.
	LocalInterfaceId() uint32
	// SetLocalInterfaceId assigns uint32 provided by user to BgpSrteSegmentETypeSubTlv
	SetLocalInterfaceId(value uint32) BgpSrteSegmentETypeSubTlv
	// HasLocalInterfaceId checks if LocalInterfaceId has been set in BgpSrteSegmentETypeSubTlv
	HasLocalInterfaceId() bool
	// Ipv4NodeAddress returns string, set in BgpSrteSegmentETypeSubTlv.
	Ipv4NodeAddress() string
	// SetIpv4NodeAddress assigns string provided by user to BgpSrteSegmentETypeSubTlv
	SetIpv4NodeAddress(value string) BgpSrteSegmentETypeSubTlv
	// SrMplsSid returns BgpSrteSrMplsSid, set in BgpSrteSegmentETypeSubTlv.
	// BgpSrteSrMplsSid is configuration for SR-MPLS with Label, TC, Bottom-of-Stack and TTL.
	SrMplsSid() BgpSrteSrMplsSid
	// SetSrMplsSid assigns BgpSrteSrMplsSid provided by user to BgpSrteSegmentETypeSubTlv.
	// BgpSrteSrMplsSid is configuration for SR-MPLS with Label, TC, Bottom-of-Stack and TTL.
	SetSrMplsSid(value BgpSrteSrMplsSid) BgpSrteSegmentETypeSubTlv
	// HasSrMplsSid checks if SrMplsSid has been set in BgpSrteSegmentETypeSubTlv
	HasSrMplsSid() bool
	setNil()
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentETypeSubTlv) Flags() string {

	return *obj.obj.Flags

}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentETypeSubTlv) HasFlags() bool {
	return obj.obj.Flags != nil
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// SetFlags sets the string value in the BgpSrteSegmentETypeSubTlv object
func (obj *bgpSrteSegmentETypeSubTlv) SetFlags(value string) BgpSrteSegmentETypeSubTlv {

	obj.obj.Flags = &value
	return obj
}

// Local Interface ID: The Interface Index as defined in [RFC8664].
// LocalInterfaceId returns a uint32
func (obj *bgpSrteSegmentETypeSubTlv) LocalInterfaceId() uint32 {

	return *obj.obj.LocalInterfaceId

}

// Local Interface ID: The Interface Index as defined in [RFC8664].
// LocalInterfaceId returns a uint32
func (obj *bgpSrteSegmentETypeSubTlv) HasLocalInterfaceId() bool {
	return obj.obj.LocalInterfaceId != nil
}

// Local Interface ID: The Interface Index as defined in [RFC8664].
// SetLocalInterfaceId sets the uint32 value in the BgpSrteSegmentETypeSubTlv object
func (obj *bgpSrteSegmentETypeSubTlv) SetLocalInterfaceId(value uint32) BgpSrteSegmentETypeSubTlv {

	obj.obj.LocalInterfaceId = &value
	return obj
}

// IPv4 address representing a node.
// Ipv4NodeAddress returns a string
func (obj *bgpSrteSegmentETypeSubTlv) Ipv4NodeAddress() string {

	return *obj.obj.Ipv4NodeAddress

}

// IPv4 address representing a node.
// SetIpv4NodeAddress sets the string value in the BgpSrteSegmentETypeSubTlv object
func (obj *bgpSrteSegmentETypeSubTlv) SetIpv4NodeAddress(value string) BgpSrteSegmentETypeSubTlv {

	obj.obj.Ipv4NodeAddress = &value
	return obj
}

// Optional SR-MPLS SID.
// SrMplsSid returns a BgpSrteSrMplsSid
func (obj *bgpSrteSegmentETypeSubTlv) SrMplsSid() BgpSrteSrMplsSid {
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
func (obj *bgpSrteSegmentETypeSubTlv) HasSrMplsSid() bool {
	return obj.obj.SrMplsSid != nil
}

// Optional SR-MPLS SID.
// SetSrMplsSid sets the BgpSrteSrMplsSid value in the BgpSrteSegmentETypeSubTlv object
func (obj *bgpSrteSegmentETypeSubTlv) SetSrMplsSid(value BgpSrteSrMplsSid) BgpSrteSegmentETypeSubTlv {

	obj.srMplsSidHolder = nil
	obj.obj.SrMplsSid = value.msg()

	return obj
}

func (obj *bgpSrteSegmentETypeSubTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentETypeSubTlv.Flags"))
		}

	}

	// Ipv4NodeAddress is required
	if obj.obj.Ipv4NodeAddress == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv4NodeAddress is required field on interface BgpSrteSegmentETypeSubTlv")
	}
	if obj.obj.Ipv4NodeAddress != nil {

		err := obj.validateIpv4(obj.Ipv4NodeAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentETypeSubTlv.Ipv4NodeAddress"))
		}

	}

	if obj.obj.SrMplsSid != nil {

		obj.SrMplsSid().validateObj(vObj, set_default)
	}

}

func (obj *bgpSrteSegmentETypeSubTlv) setDefault() {
	if obj.obj.LocalInterfaceId == nil {
		obj.SetLocalInterfaceId(0)
	}

}

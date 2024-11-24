package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrteSegmentDTypeSubTlv *****
type bgpSrteSegmentDTypeSubTlv struct {
	validation
	obj             *otg.BgpSrteSegmentDTypeSubTlv
	marshaller      marshalBgpSrteSegmentDTypeSubTlv
	unMarshaller    unMarshalBgpSrteSegmentDTypeSubTlv
	srMplsSidHolder BgpSrteSrMplsSid
}

func NewBgpSrteSegmentDTypeSubTlv() BgpSrteSegmentDTypeSubTlv {
	obj := bgpSrteSegmentDTypeSubTlv{obj: &otg.BgpSrteSegmentDTypeSubTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteSegmentDTypeSubTlv) msg() *otg.BgpSrteSegmentDTypeSubTlv {
	return obj.obj
}

func (obj *bgpSrteSegmentDTypeSubTlv) setMsg(msg *otg.BgpSrteSegmentDTypeSubTlv) BgpSrteSegmentDTypeSubTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteSegmentDTypeSubTlv struct {
	obj *bgpSrteSegmentDTypeSubTlv
}

type marshalBgpSrteSegmentDTypeSubTlv interface {
	// ToProto marshals BgpSrteSegmentDTypeSubTlv to protobuf object *otg.BgpSrteSegmentDTypeSubTlv
	ToProto() (*otg.BgpSrteSegmentDTypeSubTlv, error)
	// ToPbText marshals BgpSrteSegmentDTypeSubTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteSegmentDTypeSubTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteSegmentDTypeSubTlv to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpSrteSegmentDTypeSubTlv to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpSrteSegmentDTypeSubTlv struct {
	obj *bgpSrteSegmentDTypeSubTlv
}

type unMarshalBgpSrteSegmentDTypeSubTlv interface {
	// FromProto unmarshals BgpSrteSegmentDTypeSubTlv from protobuf object *otg.BgpSrteSegmentDTypeSubTlv
	FromProto(msg *otg.BgpSrteSegmentDTypeSubTlv) (BgpSrteSegmentDTypeSubTlv, error)
	// FromPbText unmarshals BgpSrteSegmentDTypeSubTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteSegmentDTypeSubTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteSegmentDTypeSubTlv from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteSegmentDTypeSubTlv) Marshal() marshalBgpSrteSegmentDTypeSubTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteSegmentDTypeSubTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteSegmentDTypeSubTlv) Unmarshal() unMarshalBgpSrteSegmentDTypeSubTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteSegmentDTypeSubTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteSegmentDTypeSubTlv) ToProto() (*otg.BgpSrteSegmentDTypeSubTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteSegmentDTypeSubTlv) FromProto(msg *otg.BgpSrteSegmentDTypeSubTlv) (BgpSrteSegmentDTypeSubTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteSegmentDTypeSubTlv) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteSegmentDTypeSubTlv) FromPbText(value string) error {
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

func (m *marshalbgpSrteSegmentDTypeSubTlv) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteSegmentDTypeSubTlv) FromYaml(value string) error {
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

func (m *marshalbgpSrteSegmentDTypeSubTlv) ToJsonRaw() (string, error) {
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

func (m *marshalbgpSrteSegmentDTypeSubTlv) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteSegmentDTypeSubTlv) FromJson(value string) error {
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

func (obj *bgpSrteSegmentDTypeSubTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentDTypeSubTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentDTypeSubTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteSegmentDTypeSubTlv) Clone() (BgpSrteSegmentDTypeSubTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteSegmentDTypeSubTlv()
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

func (obj *bgpSrteSegmentDTypeSubTlv) setNil() {
	obj.srMplsSidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpSrteSegmentDTypeSubTlv is type D: IPv6 Node Address with optional SID for SR MPLS.
type BgpSrteSegmentDTypeSubTlv interface {
	Validation
	// msg marshals BgpSrteSegmentDTypeSubTlv to protobuf object *otg.BgpSrteSegmentDTypeSubTlv
	// and doesn't set defaults
	msg() *otg.BgpSrteSegmentDTypeSubTlv
	// setMsg unmarshals BgpSrteSegmentDTypeSubTlv from protobuf object *otg.BgpSrteSegmentDTypeSubTlv
	// and doesn't set defaults
	setMsg(*otg.BgpSrteSegmentDTypeSubTlv) BgpSrteSegmentDTypeSubTlv
	// provides marshal interface
	Marshal() marshalBgpSrteSegmentDTypeSubTlv
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteSegmentDTypeSubTlv
	// validate validates BgpSrteSegmentDTypeSubTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteSegmentDTypeSubTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns string, set in BgpSrteSegmentDTypeSubTlv.
	Flags() string
	// SetFlags assigns string provided by user to BgpSrteSegmentDTypeSubTlv
	SetFlags(value string) BgpSrteSegmentDTypeSubTlv
	// HasFlags checks if Flags has been set in BgpSrteSegmentDTypeSubTlv
	HasFlags() bool
	// SrAlgorithm returns uint32, set in BgpSrteSegmentDTypeSubTlv.
	SrAlgorithm() uint32
	// SetSrAlgorithm assigns uint32 provided by user to BgpSrteSegmentDTypeSubTlv
	SetSrAlgorithm(value uint32) BgpSrteSegmentDTypeSubTlv
	// HasSrAlgorithm checks if SrAlgorithm has been set in BgpSrteSegmentDTypeSubTlv
	HasSrAlgorithm() bool
	// Ipv6NodeAddress returns string, set in BgpSrteSegmentDTypeSubTlv.
	Ipv6NodeAddress() string
	// SetIpv6NodeAddress assigns string provided by user to BgpSrteSegmentDTypeSubTlv
	SetIpv6NodeAddress(value string) BgpSrteSegmentDTypeSubTlv
	// SrMplsSid returns BgpSrteSrMplsSid, set in BgpSrteSegmentDTypeSubTlv.
	// BgpSrteSrMplsSid is configuration for SR-MPLS with Label, TC, Bottom-of-Stack and TTL.
	SrMplsSid() BgpSrteSrMplsSid
	// SetSrMplsSid assigns BgpSrteSrMplsSid provided by user to BgpSrteSegmentDTypeSubTlv.
	// BgpSrteSrMplsSid is configuration for SR-MPLS with Label, TC, Bottom-of-Stack and TTL.
	SetSrMplsSid(value BgpSrteSrMplsSid) BgpSrteSegmentDTypeSubTlv
	// HasSrMplsSid checks if SrMplsSid has been set in BgpSrteSegmentDTypeSubTlv
	HasSrMplsSid() bool
	setNil()
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentDTypeSubTlv) Flags() string {

	return *obj.obj.Flags

}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentDTypeSubTlv) HasFlags() bool {
	return obj.obj.Flags != nil
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// SetFlags sets the string value in the BgpSrteSegmentDTypeSubTlv object
func (obj *bgpSrteSegmentDTypeSubTlv) SetFlags(value string) BgpSrteSegmentDTypeSubTlv {

	obj.obj.Flags = &value
	return obj
}

// specifying SR Algorithm when when A-Flag as defined in above flags.
// SrAlgorithm returns a uint32
func (obj *bgpSrteSegmentDTypeSubTlv) SrAlgorithm() uint32 {

	return *obj.obj.SrAlgorithm

}

// specifying SR Algorithm when when A-Flag as defined in above flags.
// SrAlgorithm returns a uint32
func (obj *bgpSrteSegmentDTypeSubTlv) HasSrAlgorithm() bool {
	return obj.obj.SrAlgorithm != nil
}

// specifying SR Algorithm when when A-Flag as defined in above flags.
// SetSrAlgorithm sets the uint32 value in the BgpSrteSegmentDTypeSubTlv object
func (obj *bgpSrteSegmentDTypeSubTlv) SetSrAlgorithm(value uint32) BgpSrteSegmentDTypeSubTlv {

	obj.obj.SrAlgorithm = &value
	return obj
}

// IPv6 address representing a node.
// Ipv6NodeAddress returns a string
func (obj *bgpSrteSegmentDTypeSubTlv) Ipv6NodeAddress() string {

	return *obj.obj.Ipv6NodeAddress

}

// IPv6 address representing a node.
// SetIpv6NodeAddress sets the string value in the BgpSrteSegmentDTypeSubTlv object
func (obj *bgpSrteSegmentDTypeSubTlv) SetIpv6NodeAddress(value string) BgpSrteSegmentDTypeSubTlv {

	obj.obj.Ipv6NodeAddress = &value
	return obj
}

// Optional SR-MPLS SID.
// SrMplsSid returns a BgpSrteSrMplsSid
func (obj *bgpSrteSegmentDTypeSubTlv) SrMplsSid() BgpSrteSrMplsSid {
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
func (obj *bgpSrteSegmentDTypeSubTlv) HasSrMplsSid() bool {
	return obj.obj.SrMplsSid != nil
}

// Optional SR-MPLS SID.
// SetSrMplsSid sets the BgpSrteSrMplsSid value in the BgpSrteSegmentDTypeSubTlv object
func (obj *bgpSrteSegmentDTypeSubTlv) SetSrMplsSid(value BgpSrteSrMplsSid) BgpSrteSegmentDTypeSubTlv {

	obj.srMplsSidHolder = nil
	obj.obj.SrMplsSid = value.msg()

	return obj
}

func (obj *bgpSrteSegmentDTypeSubTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentDTypeSubTlv.Flags"))
		}

	}

	if obj.obj.SrAlgorithm != nil {

		if *obj.obj.SrAlgorithm > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpSrteSegmentDTypeSubTlv.SrAlgorithm <= 255 but Got %d", *obj.obj.SrAlgorithm))
		}

	}

	// Ipv6NodeAddress is required
	if obj.obj.Ipv6NodeAddress == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv6NodeAddress is required field on interface BgpSrteSegmentDTypeSubTlv")
	}
	if obj.obj.Ipv6NodeAddress != nil {

		err := obj.validateIpv6(obj.Ipv6NodeAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentDTypeSubTlv.Ipv6NodeAddress"))
		}

	}

	if obj.obj.SrMplsSid != nil {

		obj.SrMplsSid().validateObj(vObj, set_default)
	}

}

func (obj *bgpSrteSegmentDTypeSubTlv) setDefault() {
	if obj.obj.SrAlgorithm == nil {
		obj.SetSrAlgorithm(0)
	}

}

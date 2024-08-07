package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrteSegmentCTypeSubTlv *****
type bgpSrteSegmentCTypeSubTlv struct {
	validation
	obj             *otg.BgpSrteSegmentCTypeSubTlv
	marshaller      marshalBgpSrteSegmentCTypeSubTlv
	unMarshaller    unMarshalBgpSrteSegmentCTypeSubTlv
	srMplsSidHolder BgpSrteSrMplsSid
}

func NewBgpSrteSegmentCTypeSubTlv() BgpSrteSegmentCTypeSubTlv {
	obj := bgpSrteSegmentCTypeSubTlv{obj: &otg.BgpSrteSegmentCTypeSubTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteSegmentCTypeSubTlv) msg() *otg.BgpSrteSegmentCTypeSubTlv {
	return obj.obj
}

func (obj *bgpSrteSegmentCTypeSubTlv) setMsg(msg *otg.BgpSrteSegmentCTypeSubTlv) BgpSrteSegmentCTypeSubTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteSegmentCTypeSubTlv struct {
	obj *bgpSrteSegmentCTypeSubTlv
}

type marshalBgpSrteSegmentCTypeSubTlv interface {
	// ToProto marshals BgpSrteSegmentCTypeSubTlv to protobuf object *otg.BgpSrteSegmentCTypeSubTlv
	ToProto() (*otg.BgpSrteSegmentCTypeSubTlv, error)
	// ToPbText marshals BgpSrteSegmentCTypeSubTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteSegmentCTypeSubTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteSegmentCTypeSubTlv to JSON text
	ToJson() (string, error)
}

type unMarshalbgpSrteSegmentCTypeSubTlv struct {
	obj *bgpSrteSegmentCTypeSubTlv
}

type unMarshalBgpSrteSegmentCTypeSubTlv interface {
	// FromProto unmarshals BgpSrteSegmentCTypeSubTlv from protobuf object *otg.BgpSrteSegmentCTypeSubTlv
	FromProto(msg *otg.BgpSrteSegmentCTypeSubTlv) (BgpSrteSegmentCTypeSubTlv, error)
	// FromPbText unmarshals BgpSrteSegmentCTypeSubTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteSegmentCTypeSubTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteSegmentCTypeSubTlv from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteSegmentCTypeSubTlv) Marshal() marshalBgpSrteSegmentCTypeSubTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteSegmentCTypeSubTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteSegmentCTypeSubTlv) Unmarshal() unMarshalBgpSrteSegmentCTypeSubTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteSegmentCTypeSubTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteSegmentCTypeSubTlv) ToProto() (*otg.BgpSrteSegmentCTypeSubTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteSegmentCTypeSubTlv) FromProto(msg *otg.BgpSrteSegmentCTypeSubTlv) (BgpSrteSegmentCTypeSubTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteSegmentCTypeSubTlv) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteSegmentCTypeSubTlv) FromPbText(value string) error {
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

func (m *marshalbgpSrteSegmentCTypeSubTlv) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteSegmentCTypeSubTlv) FromYaml(value string) error {
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

func (m *marshalbgpSrteSegmentCTypeSubTlv) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteSegmentCTypeSubTlv) FromJson(value string) error {
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

func (obj *bgpSrteSegmentCTypeSubTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentCTypeSubTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentCTypeSubTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteSegmentCTypeSubTlv) Clone() (BgpSrteSegmentCTypeSubTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteSegmentCTypeSubTlv()
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

func (obj *bgpSrteSegmentCTypeSubTlv) setNil() {
	obj.srMplsSidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpSrteSegmentCTypeSubTlv is type C: IPv4 Node Address with optional SID.
type BgpSrteSegmentCTypeSubTlv interface {
	Validation
	// msg marshals BgpSrteSegmentCTypeSubTlv to protobuf object *otg.BgpSrteSegmentCTypeSubTlv
	// and doesn't set defaults
	msg() *otg.BgpSrteSegmentCTypeSubTlv
	// setMsg unmarshals BgpSrteSegmentCTypeSubTlv from protobuf object *otg.BgpSrteSegmentCTypeSubTlv
	// and doesn't set defaults
	setMsg(*otg.BgpSrteSegmentCTypeSubTlv) BgpSrteSegmentCTypeSubTlv
	// provides marshal interface
	Marshal() marshalBgpSrteSegmentCTypeSubTlv
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteSegmentCTypeSubTlv
	// validate validates BgpSrteSegmentCTypeSubTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteSegmentCTypeSubTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns string, set in BgpSrteSegmentCTypeSubTlv.
	Flags() string
	// SetFlags assigns string provided by user to BgpSrteSegmentCTypeSubTlv
	SetFlags(value string) BgpSrteSegmentCTypeSubTlv
	// HasFlags checks if Flags has been set in BgpSrteSegmentCTypeSubTlv
	HasFlags() bool
	// SrAlgorithm returns uint32, set in BgpSrteSegmentCTypeSubTlv.
	SrAlgorithm() uint32
	// SetSrAlgorithm assigns uint32 provided by user to BgpSrteSegmentCTypeSubTlv
	SetSrAlgorithm(value uint32) BgpSrteSegmentCTypeSubTlv
	// HasSrAlgorithm checks if SrAlgorithm has been set in BgpSrteSegmentCTypeSubTlv
	HasSrAlgorithm() bool
	// Ipv4NodeAddress returns string, set in BgpSrteSegmentCTypeSubTlv.
	Ipv4NodeAddress() string
	// SetIpv4NodeAddress assigns string provided by user to BgpSrteSegmentCTypeSubTlv
	SetIpv4NodeAddress(value string) BgpSrteSegmentCTypeSubTlv
	// SrMplsSid returns BgpSrteSrMplsSid, set in BgpSrteSegmentCTypeSubTlv.
	// BgpSrteSrMplsSid is configuration for SR-MPLS with Label, TC, Bottom-of-Stack and TTL.
	SrMplsSid() BgpSrteSrMplsSid
	// SetSrMplsSid assigns BgpSrteSrMplsSid provided by user to BgpSrteSegmentCTypeSubTlv.
	// BgpSrteSrMplsSid is configuration for SR-MPLS with Label, TC, Bottom-of-Stack and TTL.
	SetSrMplsSid(value BgpSrteSrMplsSid) BgpSrteSegmentCTypeSubTlv
	// HasSrMplsSid checks if SrMplsSid has been set in BgpSrteSegmentCTypeSubTlv
	HasSrMplsSid() bool
	setNil()
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentCTypeSubTlv) Flags() string {

	return *obj.obj.Flags

}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentCTypeSubTlv) HasFlags() bool {
	return obj.obj.Flags != nil
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// SetFlags sets the string value in the BgpSrteSegmentCTypeSubTlv object
func (obj *bgpSrteSegmentCTypeSubTlv) SetFlags(value string) BgpSrteSegmentCTypeSubTlv {

	obj.obj.Flags = &value
	return obj
}

// SR Algorithm identifier when A-Flag in on.
// SrAlgorithm returns a uint32
func (obj *bgpSrteSegmentCTypeSubTlv) SrAlgorithm() uint32 {

	return *obj.obj.SrAlgorithm

}

// SR Algorithm identifier when A-Flag in on.
// SrAlgorithm returns a uint32
func (obj *bgpSrteSegmentCTypeSubTlv) HasSrAlgorithm() bool {
	return obj.obj.SrAlgorithm != nil
}

// SR Algorithm identifier when A-Flag in on.
// SetSrAlgorithm sets the uint32 value in the BgpSrteSegmentCTypeSubTlv object
func (obj *bgpSrteSegmentCTypeSubTlv) SetSrAlgorithm(value uint32) BgpSrteSegmentCTypeSubTlv {

	obj.obj.SrAlgorithm = &value
	return obj
}

// IPv4 address representing a node.
// Ipv4NodeAddress returns a string
func (obj *bgpSrteSegmentCTypeSubTlv) Ipv4NodeAddress() string {

	return *obj.obj.Ipv4NodeAddress

}

// IPv4 address representing a node.
// SetIpv4NodeAddress sets the string value in the BgpSrteSegmentCTypeSubTlv object
func (obj *bgpSrteSegmentCTypeSubTlv) SetIpv4NodeAddress(value string) BgpSrteSegmentCTypeSubTlv {

	obj.obj.Ipv4NodeAddress = &value
	return obj
}

// Optional SR-MPLS SID.
// SrMplsSid returns a BgpSrteSrMplsSid
func (obj *bgpSrteSegmentCTypeSubTlv) SrMplsSid() BgpSrteSrMplsSid {
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
func (obj *bgpSrteSegmentCTypeSubTlv) HasSrMplsSid() bool {
	return obj.obj.SrMplsSid != nil
}

// Optional SR-MPLS SID.
// SetSrMplsSid sets the BgpSrteSrMplsSid value in the BgpSrteSegmentCTypeSubTlv object
func (obj *bgpSrteSegmentCTypeSubTlv) SetSrMplsSid(value BgpSrteSrMplsSid) BgpSrteSegmentCTypeSubTlv {

	obj.srMplsSidHolder = nil
	obj.obj.SrMplsSid = value.msg()

	return obj
}

func (obj *bgpSrteSegmentCTypeSubTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentCTypeSubTlv.Flags"))
		}

	}

	if obj.obj.SrAlgorithm != nil {

		if *obj.obj.SrAlgorithm > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpSrteSegmentCTypeSubTlv.SrAlgorithm <= 255 but Got %d", *obj.obj.SrAlgorithm))
		}

	}

	// Ipv4NodeAddress is required
	if obj.obj.Ipv4NodeAddress == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv4NodeAddress is required field on interface BgpSrteSegmentCTypeSubTlv")
	}
	if obj.obj.Ipv4NodeAddress != nil {

		err := obj.validateIpv4(obj.Ipv4NodeAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentCTypeSubTlv.Ipv4NodeAddress"))
		}

	}

	if obj.obj.SrMplsSid != nil {

		obj.SrMplsSid().validateObj(vObj, set_default)
	}

}

func (obj *bgpSrteSegmentCTypeSubTlv) setDefault() {
	if obj.obj.SrAlgorithm == nil {
		obj.SetSrAlgorithm(0)
	}

}

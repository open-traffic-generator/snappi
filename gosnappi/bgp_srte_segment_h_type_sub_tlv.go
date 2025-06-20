package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrteSegmentHTypeSubTlv *****
type bgpSrteSegmentHTypeSubTlv struct {
	validation
	obj             *otg.BgpSrteSegmentHTypeSubTlv
	marshaller      marshalBgpSrteSegmentHTypeSubTlv
	unMarshaller    unMarshalBgpSrteSegmentHTypeSubTlv
	srMplsSidHolder BgpSrteSrMplsSid
}

func NewBgpSrteSegmentHTypeSubTlv() BgpSrteSegmentHTypeSubTlv {
	obj := bgpSrteSegmentHTypeSubTlv{obj: &otg.BgpSrteSegmentHTypeSubTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteSegmentHTypeSubTlv) msg() *otg.BgpSrteSegmentHTypeSubTlv {
	return obj.obj
}

func (obj *bgpSrteSegmentHTypeSubTlv) setMsg(msg *otg.BgpSrteSegmentHTypeSubTlv) BgpSrteSegmentHTypeSubTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteSegmentHTypeSubTlv struct {
	obj *bgpSrteSegmentHTypeSubTlv
}

type marshalBgpSrteSegmentHTypeSubTlv interface {
	// ToProto marshals BgpSrteSegmentHTypeSubTlv to protobuf object *otg.BgpSrteSegmentHTypeSubTlv
	ToProto() (*otg.BgpSrteSegmentHTypeSubTlv, error)
	// ToPbText marshals BgpSrteSegmentHTypeSubTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteSegmentHTypeSubTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteSegmentHTypeSubTlv to JSON text
	ToJson() (string, error)
}

type unMarshalbgpSrteSegmentHTypeSubTlv struct {
	obj *bgpSrteSegmentHTypeSubTlv
}

type unMarshalBgpSrteSegmentHTypeSubTlv interface {
	// FromProto unmarshals BgpSrteSegmentHTypeSubTlv from protobuf object *otg.BgpSrteSegmentHTypeSubTlv
	FromProto(msg *otg.BgpSrteSegmentHTypeSubTlv) (BgpSrteSegmentHTypeSubTlv, error)
	// FromPbText unmarshals BgpSrteSegmentHTypeSubTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteSegmentHTypeSubTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteSegmentHTypeSubTlv from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteSegmentHTypeSubTlv) Marshal() marshalBgpSrteSegmentHTypeSubTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteSegmentHTypeSubTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteSegmentHTypeSubTlv) Unmarshal() unMarshalBgpSrteSegmentHTypeSubTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteSegmentHTypeSubTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteSegmentHTypeSubTlv) ToProto() (*otg.BgpSrteSegmentHTypeSubTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteSegmentHTypeSubTlv) FromProto(msg *otg.BgpSrteSegmentHTypeSubTlv) (BgpSrteSegmentHTypeSubTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteSegmentHTypeSubTlv) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteSegmentHTypeSubTlv) FromPbText(value string) error {
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

func (m *marshalbgpSrteSegmentHTypeSubTlv) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteSegmentHTypeSubTlv) FromYaml(value string) error {
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

func (m *marshalbgpSrteSegmentHTypeSubTlv) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteSegmentHTypeSubTlv) FromJson(value string) error {
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

func (obj *bgpSrteSegmentHTypeSubTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentHTypeSubTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentHTypeSubTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteSegmentHTypeSubTlv) Clone() (BgpSrteSegmentHTypeSubTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteSegmentHTypeSubTlv()
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

func (obj *bgpSrteSegmentHTypeSubTlv) setNil() {
	obj.srMplsSidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpSrteSegmentHTypeSubTlv is type H: IPv6 Local and Remote addresses with optional SID for SR MPLS.
type BgpSrteSegmentHTypeSubTlv interface {
	Validation
	// msg marshals BgpSrteSegmentHTypeSubTlv to protobuf object *otg.BgpSrteSegmentHTypeSubTlv
	// and doesn't set defaults
	msg() *otg.BgpSrteSegmentHTypeSubTlv
	// setMsg unmarshals BgpSrteSegmentHTypeSubTlv from protobuf object *otg.BgpSrteSegmentHTypeSubTlv
	// and doesn't set defaults
	setMsg(*otg.BgpSrteSegmentHTypeSubTlv) BgpSrteSegmentHTypeSubTlv
	// provides marshal interface
	Marshal() marshalBgpSrteSegmentHTypeSubTlv
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteSegmentHTypeSubTlv
	// validate validates BgpSrteSegmentHTypeSubTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteSegmentHTypeSubTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns string, set in BgpSrteSegmentHTypeSubTlv.
	Flags() string
	// SetFlags assigns string provided by user to BgpSrteSegmentHTypeSubTlv
	SetFlags(value string) BgpSrteSegmentHTypeSubTlv
	// HasFlags checks if Flags has been set in BgpSrteSegmentHTypeSubTlv
	HasFlags() bool
	// LocalIpv6Address returns string, set in BgpSrteSegmentHTypeSubTlv.
	LocalIpv6Address() string
	// SetLocalIpv6Address assigns string provided by user to BgpSrteSegmentHTypeSubTlv
	SetLocalIpv6Address(value string) BgpSrteSegmentHTypeSubTlv
	// RemoteIpv6Address returns string, set in BgpSrteSegmentHTypeSubTlv.
	RemoteIpv6Address() string
	// SetRemoteIpv6Address assigns string provided by user to BgpSrteSegmentHTypeSubTlv
	SetRemoteIpv6Address(value string) BgpSrteSegmentHTypeSubTlv
	// SrMplsSid returns BgpSrteSrMplsSid, set in BgpSrteSegmentHTypeSubTlv.
	// BgpSrteSrMplsSid is configuration for SR-MPLS with Label, TC, Bottom-of-Stack and TTL.
	SrMplsSid() BgpSrteSrMplsSid
	// SetSrMplsSid assigns BgpSrteSrMplsSid provided by user to BgpSrteSegmentHTypeSubTlv.
	// BgpSrteSrMplsSid is configuration for SR-MPLS with Label, TC, Bottom-of-Stack and TTL.
	SetSrMplsSid(value BgpSrteSrMplsSid) BgpSrteSegmentHTypeSubTlv
	// HasSrMplsSid checks if SrMplsSid has been set in BgpSrteSegmentHTypeSubTlv
	HasSrMplsSid() bool
	setNil()
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentHTypeSubTlv) Flags() string {

	return *obj.obj.Flags

}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentHTypeSubTlv) HasFlags() bool {
	return obj.obj.Flags != nil
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// SetFlags sets the string value in the BgpSrteSegmentHTypeSubTlv object
func (obj *bgpSrteSegmentHTypeSubTlv) SetFlags(value string) BgpSrteSegmentHTypeSubTlv {

	obj.obj.Flags = &value
	return obj
}

// Local IPv6 Address.
// LocalIpv6Address returns a string
func (obj *bgpSrteSegmentHTypeSubTlv) LocalIpv6Address() string {

	return *obj.obj.LocalIpv6Address

}

// Local IPv6 Address.
// SetLocalIpv6Address sets the string value in the BgpSrteSegmentHTypeSubTlv object
func (obj *bgpSrteSegmentHTypeSubTlv) SetLocalIpv6Address(value string) BgpSrteSegmentHTypeSubTlv {

	obj.obj.LocalIpv6Address = &value
	return obj
}

// Remote IPv6 Address.
// RemoteIpv6Address returns a string
func (obj *bgpSrteSegmentHTypeSubTlv) RemoteIpv6Address() string {

	return *obj.obj.RemoteIpv6Address

}

// Remote IPv6 Address.
// SetRemoteIpv6Address sets the string value in the BgpSrteSegmentHTypeSubTlv object
func (obj *bgpSrteSegmentHTypeSubTlv) SetRemoteIpv6Address(value string) BgpSrteSegmentHTypeSubTlv {

	obj.obj.RemoteIpv6Address = &value
	return obj
}

// Optional SR-MPLS SID.
// SrMplsSid returns a BgpSrteSrMplsSid
func (obj *bgpSrteSegmentHTypeSubTlv) SrMplsSid() BgpSrteSrMplsSid {
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
func (obj *bgpSrteSegmentHTypeSubTlv) HasSrMplsSid() bool {
	return obj.obj.SrMplsSid != nil
}

// Optional SR-MPLS SID.
// SetSrMplsSid sets the BgpSrteSrMplsSid value in the BgpSrteSegmentHTypeSubTlv object
func (obj *bgpSrteSegmentHTypeSubTlv) SetSrMplsSid(value BgpSrteSrMplsSid) BgpSrteSegmentHTypeSubTlv {

	obj.srMplsSidHolder = nil
	obj.obj.SrMplsSid = value.msg()

	return obj
}

func (obj *bgpSrteSegmentHTypeSubTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentHTypeSubTlv.Flags"))
		}

	}

	// LocalIpv6Address is required
	if obj.obj.LocalIpv6Address == nil {
		vObj.validationErrors = append(vObj.validationErrors, "LocalIpv6Address is required field on interface BgpSrteSegmentHTypeSubTlv")
	}
	if obj.obj.LocalIpv6Address != nil {

		err := obj.validateIpv6(obj.LocalIpv6Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentHTypeSubTlv.LocalIpv6Address"))
		}

	}

	// RemoteIpv6Address is required
	if obj.obj.RemoteIpv6Address == nil {
		vObj.validationErrors = append(vObj.validationErrors, "RemoteIpv6Address is required field on interface BgpSrteSegmentHTypeSubTlv")
	}
	if obj.obj.RemoteIpv6Address != nil {

		err := obj.validateIpv6(obj.RemoteIpv6Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentHTypeSubTlv.RemoteIpv6Address"))
		}

	}

	if obj.obj.SrMplsSid != nil {

		obj.SrMplsSid().validateObj(vObj, set_default)
	}

}

func (obj *bgpSrteSegmentHTypeSubTlv) setDefault() {

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrteSegmentFTypeSubTlv *****
type bgpSrteSegmentFTypeSubTlv struct {
	validation
	obj             *otg.BgpSrteSegmentFTypeSubTlv
	marshaller      marshalBgpSrteSegmentFTypeSubTlv
	unMarshaller    unMarshalBgpSrteSegmentFTypeSubTlv
	srMplsSidHolder BgpSrteSrMplsSid
}

func NewBgpSrteSegmentFTypeSubTlv() BgpSrteSegmentFTypeSubTlv {
	obj := bgpSrteSegmentFTypeSubTlv{obj: &otg.BgpSrteSegmentFTypeSubTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteSegmentFTypeSubTlv) msg() *otg.BgpSrteSegmentFTypeSubTlv {
	return obj.obj
}

func (obj *bgpSrteSegmentFTypeSubTlv) setMsg(msg *otg.BgpSrteSegmentFTypeSubTlv) BgpSrteSegmentFTypeSubTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteSegmentFTypeSubTlv struct {
	obj *bgpSrteSegmentFTypeSubTlv
}

type marshalBgpSrteSegmentFTypeSubTlv interface {
	// ToProto marshals BgpSrteSegmentFTypeSubTlv to protobuf object *otg.BgpSrteSegmentFTypeSubTlv
	ToProto() (*otg.BgpSrteSegmentFTypeSubTlv, error)
	// ToPbText marshals BgpSrteSegmentFTypeSubTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteSegmentFTypeSubTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteSegmentFTypeSubTlv to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpSrteSegmentFTypeSubTlv to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpSrteSegmentFTypeSubTlv struct {
	obj *bgpSrteSegmentFTypeSubTlv
}

type unMarshalBgpSrteSegmentFTypeSubTlv interface {
	// FromProto unmarshals BgpSrteSegmentFTypeSubTlv from protobuf object *otg.BgpSrteSegmentFTypeSubTlv
	FromProto(msg *otg.BgpSrteSegmentFTypeSubTlv) (BgpSrteSegmentFTypeSubTlv, error)
	// FromPbText unmarshals BgpSrteSegmentFTypeSubTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteSegmentFTypeSubTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteSegmentFTypeSubTlv from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteSegmentFTypeSubTlv) Marshal() marshalBgpSrteSegmentFTypeSubTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteSegmentFTypeSubTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteSegmentFTypeSubTlv) Unmarshal() unMarshalBgpSrteSegmentFTypeSubTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteSegmentFTypeSubTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteSegmentFTypeSubTlv) ToProto() (*otg.BgpSrteSegmentFTypeSubTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteSegmentFTypeSubTlv) FromProto(msg *otg.BgpSrteSegmentFTypeSubTlv) (BgpSrteSegmentFTypeSubTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteSegmentFTypeSubTlv) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteSegmentFTypeSubTlv) FromPbText(value string) error {
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

func (m *marshalbgpSrteSegmentFTypeSubTlv) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteSegmentFTypeSubTlv) FromYaml(value string) error {
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

func (m *marshalbgpSrteSegmentFTypeSubTlv) ToJsonRaw() (string, error) {
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

func (m *marshalbgpSrteSegmentFTypeSubTlv) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteSegmentFTypeSubTlv) FromJson(value string) error {
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

func (obj *bgpSrteSegmentFTypeSubTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentFTypeSubTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentFTypeSubTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteSegmentFTypeSubTlv) Clone() (BgpSrteSegmentFTypeSubTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteSegmentFTypeSubTlv()
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

func (obj *bgpSrteSegmentFTypeSubTlv) setNil() {
	obj.srMplsSidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpSrteSegmentFTypeSubTlv is type F: IPv4 Local and Remote addresses with optional SID.
type BgpSrteSegmentFTypeSubTlv interface {
	Validation
	// msg marshals BgpSrteSegmentFTypeSubTlv to protobuf object *otg.BgpSrteSegmentFTypeSubTlv
	// and doesn't set defaults
	msg() *otg.BgpSrteSegmentFTypeSubTlv
	// setMsg unmarshals BgpSrteSegmentFTypeSubTlv from protobuf object *otg.BgpSrteSegmentFTypeSubTlv
	// and doesn't set defaults
	setMsg(*otg.BgpSrteSegmentFTypeSubTlv) BgpSrteSegmentFTypeSubTlv
	// provides marshal interface
	Marshal() marshalBgpSrteSegmentFTypeSubTlv
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteSegmentFTypeSubTlv
	// validate validates BgpSrteSegmentFTypeSubTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteSegmentFTypeSubTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns string, set in BgpSrteSegmentFTypeSubTlv.
	Flags() string
	// SetFlags assigns string provided by user to BgpSrteSegmentFTypeSubTlv
	SetFlags(value string) BgpSrteSegmentFTypeSubTlv
	// HasFlags checks if Flags has been set in BgpSrteSegmentFTypeSubTlv
	HasFlags() bool
	// LocalIpv4Address returns string, set in BgpSrteSegmentFTypeSubTlv.
	LocalIpv4Address() string
	// SetLocalIpv4Address assigns string provided by user to BgpSrteSegmentFTypeSubTlv
	SetLocalIpv4Address(value string) BgpSrteSegmentFTypeSubTlv
	// RemoteIpv4Address returns string, set in BgpSrteSegmentFTypeSubTlv.
	RemoteIpv4Address() string
	// SetRemoteIpv4Address assigns string provided by user to BgpSrteSegmentFTypeSubTlv
	SetRemoteIpv4Address(value string) BgpSrteSegmentFTypeSubTlv
	// SrMplsSid returns BgpSrteSrMplsSid, set in BgpSrteSegmentFTypeSubTlv.
	// BgpSrteSrMplsSid is configuration for SR-MPLS with Label, TC, Bottom-of-Stack and TTL.
	SrMplsSid() BgpSrteSrMplsSid
	// SetSrMplsSid assigns BgpSrteSrMplsSid provided by user to BgpSrteSegmentFTypeSubTlv.
	// BgpSrteSrMplsSid is configuration for SR-MPLS with Label, TC, Bottom-of-Stack and TTL.
	SetSrMplsSid(value BgpSrteSrMplsSid) BgpSrteSegmentFTypeSubTlv
	// HasSrMplsSid checks if SrMplsSid has been set in BgpSrteSegmentFTypeSubTlv
	HasSrMplsSid() bool
	setNil()
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentFTypeSubTlv) Flags() string {

	return *obj.obj.Flags

}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentFTypeSubTlv) HasFlags() bool {
	return obj.obj.Flags != nil
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// SetFlags sets the string value in the BgpSrteSegmentFTypeSubTlv object
func (obj *bgpSrteSegmentFTypeSubTlv) SetFlags(value string) BgpSrteSegmentFTypeSubTlv {

	obj.obj.Flags = &value
	return obj
}

// Local IPv4 Address.
// LocalIpv4Address returns a string
func (obj *bgpSrteSegmentFTypeSubTlv) LocalIpv4Address() string {

	return *obj.obj.LocalIpv4Address

}

// Local IPv4 Address.
// SetLocalIpv4Address sets the string value in the BgpSrteSegmentFTypeSubTlv object
func (obj *bgpSrteSegmentFTypeSubTlv) SetLocalIpv4Address(value string) BgpSrteSegmentFTypeSubTlv {

	obj.obj.LocalIpv4Address = &value
	return obj
}

// Remote IPv4 Address.
// RemoteIpv4Address returns a string
func (obj *bgpSrteSegmentFTypeSubTlv) RemoteIpv4Address() string {

	return *obj.obj.RemoteIpv4Address

}

// Remote IPv4 Address.
// SetRemoteIpv4Address sets the string value in the BgpSrteSegmentFTypeSubTlv object
func (obj *bgpSrteSegmentFTypeSubTlv) SetRemoteIpv4Address(value string) BgpSrteSegmentFTypeSubTlv {

	obj.obj.RemoteIpv4Address = &value
	return obj
}

// Optional SR-MPLS SID.
// SrMplsSid returns a BgpSrteSrMplsSid
func (obj *bgpSrteSegmentFTypeSubTlv) SrMplsSid() BgpSrteSrMplsSid {
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
func (obj *bgpSrteSegmentFTypeSubTlv) HasSrMplsSid() bool {
	return obj.obj.SrMplsSid != nil
}

// Optional SR-MPLS SID.
// SetSrMplsSid sets the BgpSrteSrMplsSid value in the BgpSrteSegmentFTypeSubTlv object
func (obj *bgpSrteSegmentFTypeSubTlv) SetSrMplsSid(value BgpSrteSrMplsSid) BgpSrteSegmentFTypeSubTlv {

	obj.srMplsSidHolder = nil
	obj.obj.SrMplsSid = value.msg()

	return obj
}

func (obj *bgpSrteSegmentFTypeSubTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentFTypeSubTlv.Flags"))
		}

	}

	// LocalIpv4Address is required
	if obj.obj.LocalIpv4Address == nil {
		vObj.validationErrors = append(vObj.validationErrors, "LocalIpv4Address is required field on interface BgpSrteSegmentFTypeSubTlv")
	}
	if obj.obj.LocalIpv4Address != nil {

		err := obj.validateIpv4(obj.LocalIpv4Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentFTypeSubTlv.LocalIpv4Address"))
		}

	}

	// RemoteIpv4Address is required
	if obj.obj.RemoteIpv4Address == nil {
		vObj.validationErrors = append(vObj.validationErrors, "RemoteIpv4Address is required field on interface BgpSrteSegmentFTypeSubTlv")
	}
	if obj.obj.RemoteIpv4Address != nil {

		err := obj.validateIpv4(obj.RemoteIpv4Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentFTypeSubTlv.RemoteIpv4Address"))
		}

	}

	if obj.obj.SrMplsSid != nil {

		obj.SrMplsSid().validateObj(vObj, set_default)
	}

}

func (obj *bgpSrteSegmentFTypeSubTlv) setDefault() {

}

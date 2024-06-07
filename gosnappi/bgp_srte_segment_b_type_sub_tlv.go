package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrteSegmentBTypeSubTlv *****
type bgpSrteSegmentBTypeSubTlv struct {
	validation
	obj                           *otg.BgpSrteSegmentBTypeSubTlv
	marshaller                    marshalBgpSrteSegmentBTypeSubTlv
	unMarshaller                  unMarshalBgpSrteSegmentBTypeSubTlv
	srv6SidEndpointBehaviorHolder BgpSrteSRv6SIDEndpointBehaviorAndStructure
}

func NewBgpSrteSegmentBTypeSubTlv() BgpSrteSegmentBTypeSubTlv {
	obj := bgpSrteSegmentBTypeSubTlv{obj: &otg.BgpSrteSegmentBTypeSubTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteSegmentBTypeSubTlv) msg() *otg.BgpSrteSegmentBTypeSubTlv {
	return obj.obj
}

func (obj *bgpSrteSegmentBTypeSubTlv) setMsg(msg *otg.BgpSrteSegmentBTypeSubTlv) BgpSrteSegmentBTypeSubTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteSegmentBTypeSubTlv struct {
	obj *bgpSrteSegmentBTypeSubTlv
}

type marshalBgpSrteSegmentBTypeSubTlv interface {
	// ToProto marshals BgpSrteSegmentBTypeSubTlv to protobuf object *otg.BgpSrteSegmentBTypeSubTlv
	ToProto() (*otg.BgpSrteSegmentBTypeSubTlv, error)
	// ToPbText marshals BgpSrteSegmentBTypeSubTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteSegmentBTypeSubTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteSegmentBTypeSubTlv to JSON text
	ToJson() (string, error)
}

type unMarshalbgpSrteSegmentBTypeSubTlv struct {
	obj *bgpSrteSegmentBTypeSubTlv
}

type unMarshalBgpSrteSegmentBTypeSubTlv interface {
	// FromProto unmarshals BgpSrteSegmentBTypeSubTlv from protobuf object *otg.BgpSrteSegmentBTypeSubTlv
	FromProto(msg *otg.BgpSrteSegmentBTypeSubTlv) (BgpSrteSegmentBTypeSubTlv, error)
	// FromPbText unmarshals BgpSrteSegmentBTypeSubTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteSegmentBTypeSubTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteSegmentBTypeSubTlv from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteSegmentBTypeSubTlv) Marshal() marshalBgpSrteSegmentBTypeSubTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteSegmentBTypeSubTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteSegmentBTypeSubTlv) Unmarshal() unMarshalBgpSrteSegmentBTypeSubTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteSegmentBTypeSubTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteSegmentBTypeSubTlv) ToProto() (*otg.BgpSrteSegmentBTypeSubTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteSegmentBTypeSubTlv) FromProto(msg *otg.BgpSrteSegmentBTypeSubTlv) (BgpSrteSegmentBTypeSubTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteSegmentBTypeSubTlv) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteSegmentBTypeSubTlv) FromPbText(value string) error {
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

func (m *marshalbgpSrteSegmentBTypeSubTlv) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteSegmentBTypeSubTlv) FromYaml(value string) error {
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

func (m *marshalbgpSrteSegmentBTypeSubTlv) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteSegmentBTypeSubTlv) FromJson(value string) error {
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

func (obj *bgpSrteSegmentBTypeSubTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentBTypeSubTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentBTypeSubTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteSegmentBTypeSubTlv) Clone() (BgpSrteSegmentBTypeSubTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteSegmentBTypeSubTlv()
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

func (obj *bgpSrteSegmentBTypeSubTlv) setNil() {
	obj.srv6SidEndpointBehaviorHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpSrteSegmentBTypeSubTlv is type  B: SID only, in the form of IPv6 address.
type BgpSrteSegmentBTypeSubTlv interface {
	Validation
	// msg marshals BgpSrteSegmentBTypeSubTlv to protobuf object *otg.BgpSrteSegmentBTypeSubTlv
	// and doesn't set defaults
	msg() *otg.BgpSrteSegmentBTypeSubTlv
	// setMsg unmarshals BgpSrteSegmentBTypeSubTlv from protobuf object *otg.BgpSrteSegmentBTypeSubTlv
	// and doesn't set defaults
	setMsg(*otg.BgpSrteSegmentBTypeSubTlv) BgpSrteSegmentBTypeSubTlv
	// provides marshal interface
	Marshal() marshalBgpSrteSegmentBTypeSubTlv
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteSegmentBTypeSubTlv
	// validate validates BgpSrteSegmentBTypeSubTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteSegmentBTypeSubTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns string, set in BgpSrteSegmentBTypeSubTlv.
	Flags() string
	// SetFlags assigns string provided by user to BgpSrteSegmentBTypeSubTlv
	SetFlags(value string) BgpSrteSegmentBTypeSubTlv
	// HasFlags checks if Flags has been set in BgpSrteSegmentBTypeSubTlv
	HasFlags() bool
	// Srv6Sid returns string, set in BgpSrteSegmentBTypeSubTlv.
	Srv6Sid() string
	// SetSrv6Sid assigns string provided by user to BgpSrteSegmentBTypeSubTlv
	SetSrv6Sid(value string) BgpSrteSegmentBTypeSubTlv
	// Srv6SidEndpointBehavior returns BgpSrteSRv6SIDEndpointBehaviorAndStructure, set in BgpSrteSegmentBTypeSubTlv.
	// BgpSrteSRv6SIDEndpointBehaviorAndStructure is configuration for SRv6 Endpoint Behavior and SID Structure.  Its optional. Summation of lengths for Locator Block, Locator Node,  Function, and Argument MUST be less than or equal to 128.
	Srv6SidEndpointBehavior() BgpSrteSRv6SIDEndpointBehaviorAndStructure
	// SetSrv6SidEndpointBehavior assigns BgpSrteSRv6SIDEndpointBehaviorAndStructure provided by user to BgpSrteSegmentBTypeSubTlv.
	// BgpSrteSRv6SIDEndpointBehaviorAndStructure is configuration for SRv6 Endpoint Behavior and SID Structure.  Its optional. Summation of lengths for Locator Block, Locator Node,  Function, and Argument MUST be less than or equal to 128.
	SetSrv6SidEndpointBehavior(value BgpSrteSRv6SIDEndpointBehaviorAndStructure) BgpSrteSegmentBTypeSubTlv
	// HasSrv6SidEndpointBehavior checks if Srv6SidEndpointBehavior has been set in BgpSrteSegmentBTypeSubTlv
	HasSrv6SidEndpointBehavior() bool
	setNil()
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentBTypeSubTlv) Flags() string {

	return *obj.obj.Flags

}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentBTypeSubTlv) HasFlags() bool {
	return obj.obj.Flags != nil
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// SetFlags sets the string value in the BgpSrteSegmentBTypeSubTlv object
func (obj *bgpSrteSegmentBTypeSubTlv) SetFlags(value string) BgpSrteSegmentBTypeSubTlv {

	obj.obj.Flags = &value
	return obj
}

// SRv6 SID.
// Srv6Sid returns a string
func (obj *bgpSrteSegmentBTypeSubTlv) Srv6Sid() string {

	return *obj.obj.Srv6Sid

}

// SRv6 SID.
// SetSrv6Sid sets the string value in the BgpSrteSegmentBTypeSubTlv object
func (obj *bgpSrteSegmentBTypeSubTlv) SetSrv6Sid(value string) BgpSrteSegmentBTypeSubTlv {

	obj.obj.Srv6Sid = &value
	return obj
}

// Optional SRv6 Endpoint Behavior and SID Structure.
// Srv6SidEndpointBehavior returns a BgpSrteSRv6SIDEndpointBehaviorAndStructure
func (obj *bgpSrteSegmentBTypeSubTlv) Srv6SidEndpointBehavior() BgpSrteSRv6SIDEndpointBehaviorAndStructure {
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
func (obj *bgpSrteSegmentBTypeSubTlv) HasSrv6SidEndpointBehavior() bool {
	return obj.obj.Srv6SidEndpointBehavior != nil
}

// Optional SRv6 Endpoint Behavior and SID Structure.
// SetSrv6SidEndpointBehavior sets the BgpSrteSRv6SIDEndpointBehaviorAndStructure value in the BgpSrteSegmentBTypeSubTlv object
func (obj *bgpSrteSegmentBTypeSubTlv) SetSrv6SidEndpointBehavior(value BgpSrteSRv6SIDEndpointBehaviorAndStructure) BgpSrteSegmentBTypeSubTlv {

	obj.srv6SidEndpointBehaviorHolder = nil
	obj.obj.Srv6SidEndpointBehavior = value.msg()

	return obj
}

func (obj *bgpSrteSegmentBTypeSubTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentBTypeSubTlv.Flags"))
		}

	}

	// Srv6Sid is required
	if obj.obj.Srv6Sid == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Srv6Sid is required field on interface BgpSrteSegmentBTypeSubTlv")
	}
	if obj.obj.Srv6Sid != nil {

		err := obj.validateIpv6(obj.Srv6Sid())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentBTypeSubTlv.Srv6Sid"))
		}

	}

	if obj.obj.Srv6SidEndpointBehavior != nil {

		obj.Srv6SidEndpointBehavior().validateObj(vObj, set_default)
	}

}

func (obj *bgpSrteSegmentBTypeSubTlv) setDefault() {

}

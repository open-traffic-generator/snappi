package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesBsidSrv6 *****
type bgpAttributesBsidSrv6 struct {
	validation
	obj          *otg.BgpAttributesBsidSrv6
	marshaller   marshalBgpAttributesBsidSrv6
	unMarshaller unMarshalBgpAttributesBsidSrv6
}

func NewBgpAttributesBsidSrv6() BgpAttributesBsidSrv6 {
	obj := bgpAttributesBsidSrv6{obj: &otg.BgpAttributesBsidSrv6{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesBsidSrv6) msg() *otg.BgpAttributesBsidSrv6 {
	return obj.obj
}

func (obj *bgpAttributesBsidSrv6) setMsg(msg *otg.BgpAttributesBsidSrv6) BgpAttributesBsidSrv6 {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesBsidSrv6 struct {
	obj *bgpAttributesBsidSrv6
}

type marshalBgpAttributesBsidSrv6 interface {
	// ToProto marshals BgpAttributesBsidSrv6 to protobuf object *otg.BgpAttributesBsidSrv6
	ToProto() (*otg.BgpAttributesBsidSrv6, error)
	// ToPbText marshals BgpAttributesBsidSrv6 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesBsidSrv6 to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesBsidSrv6 to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesBsidSrv6 to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesBsidSrv6 struct {
	obj *bgpAttributesBsidSrv6
}

type unMarshalBgpAttributesBsidSrv6 interface {
	// FromProto unmarshals BgpAttributesBsidSrv6 from protobuf object *otg.BgpAttributesBsidSrv6
	FromProto(msg *otg.BgpAttributesBsidSrv6) (BgpAttributesBsidSrv6, error)
	// FromPbText unmarshals BgpAttributesBsidSrv6 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesBsidSrv6 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesBsidSrv6 from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesBsidSrv6) Marshal() marshalBgpAttributesBsidSrv6 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesBsidSrv6{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesBsidSrv6) Unmarshal() unMarshalBgpAttributesBsidSrv6 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesBsidSrv6{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesBsidSrv6) ToProto() (*otg.BgpAttributesBsidSrv6, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesBsidSrv6) FromProto(msg *otg.BgpAttributesBsidSrv6) (BgpAttributesBsidSrv6, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesBsidSrv6) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesBsidSrv6) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalbgpAttributesBsidSrv6) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesBsidSrv6) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalbgpAttributesBsidSrv6) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesBsidSrv6) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesBsidSrv6) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *bgpAttributesBsidSrv6) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesBsidSrv6) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesBsidSrv6) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesBsidSrv6) Clone() (BgpAttributesBsidSrv6, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesBsidSrv6()
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

// BgpAttributesBsidSrv6 is when the active candidate path has a specified Binding Segment Identifier, the SR Policy uses that BSID defined
// as an IPv6 Address.The format of the sub-TLV is defined in draft-ietf-idr-sr-policy-safi-02 Section 2.4.2 .
type BgpAttributesBsidSrv6 interface {
	Validation
	// msg marshals BgpAttributesBsidSrv6 to protobuf object *otg.BgpAttributesBsidSrv6
	// and doesn't set defaults
	msg() *otg.BgpAttributesBsidSrv6
	// setMsg unmarshals BgpAttributesBsidSrv6 from protobuf object *otg.BgpAttributesBsidSrv6
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesBsidSrv6) BgpAttributesBsidSrv6
	// provides marshal interface
	Marshal() marshalBgpAttributesBsidSrv6
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesBsidSrv6
	// validate validates BgpAttributesBsidSrv6
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesBsidSrv6, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// FlagSpecifiedBsidOnly returns bool, set in BgpAttributesBsidSrv6.
	FlagSpecifiedBsidOnly() bool
	// SetFlagSpecifiedBsidOnly assigns bool provided by user to BgpAttributesBsidSrv6
	SetFlagSpecifiedBsidOnly(value bool) BgpAttributesBsidSrv6
	// HasFlagSpecifiedBsidOnly checks if FlagSpecifiedBsidOnly has been set in BgpAttributesBsidSrv6
	HasFlagSpecifiedBsidOnly() bool
	// FlagDropUponInvalid returns bool, set in BgpAttributesBsidSrv6.
	FlagDropUponInvalid() bool
	// SetFlagDropUponInvalid assigns bool provided by user to BgpAttributesBsidSrv6
	SetFlagDropUponInvalid(value bool) BgpAttributesBsidSrv6
	// HasFlagDropUponInvalid checks if FlagDropUponInvalid has been set in BgpAttributesBsidSrv6
	HasFlagDropUponInvalid() bool
	// Ipv6Addr returns string, set in BgpAttributesBsidSrv6.
	Ipv6Addr() string
	// SetIpv6Addr assigns string provided by user to BgpAttributesBsidSrv6
	SetIpv6Addr(value string) BgpAttributesBsidSrv6
	// HasIpv6Addr checks if Ipv6Addr has been set in BgpAttributesBsidSrv6
	HasIpv6Addr() bool
}

// S-Flag: This flag encodes the "Specified-BSID-only" behavior. It's usage is
// described in section 6.2.3 in [RFC9256].
// FlagSpecifiedBsidOnly returns a bool
func (obj *bgpAttributesBsidSrv6) FlagSpecifiedBsidOnly() bool {

	return *obj.obj.FlagSpecifiedBsidOnly

}

// S-Flag: This flag encodes the "Specified-BSID-only" behavior. It's usage is
// described in section 6.2.3 in [RFC9256].
// FlagSpecifiedBsidOnly returns a bool
func (obj *bgpAttributesBsidSrv6) HasFlagSpecifiedBsidOnly() bool {
	return obj.obj.FlagSpecifiedBsidOnly != nil
}

// S-Flag: This flag encodes the "Specified-BSID-only" behavior. It's usage is
// described in section 6.2.3 in [RFC9256].
// SetFlagSpecifiedBsidOnly sets the bool value in the BgpAttributesBsidSrv6 object
func (obj *bgpAttributesBsidSrv6) SetFlagSpecifiedBsidOnly(value bool) BgpAttributesBsidSrv6 {

	obj.obj.FlagSpecifiedBsidOnly = &value
	return obj
}

// I-Flag: This flag encodes the "Drop Upon Invalid" behavior.
// It's usage is described in section 8.2 in [RFC9256].
// FlagDropUponInvalid returns a bool
func (obj *bgpAttributesBsidSrv6) FlagDropUponInvalid() bool {

	return *obj.obj.FlagDropUponInvalid

}

// I-Flag: This flag encodes the "Drop Upon Invalid" behavior.
// It's usage is described in section 8.2 in [RFC9256].
// FlagDropUponInvalid returns a bool
func (obj *bgpAttributesBsidSrv6) HasFlagDropUponInvalid() bool {
	return obj.obj.FlagDropUponInvalid != nil
}

// I-Flag: This flag encodes the "Drop Upon Invalid" behavior.
// It's usage is described in section 8.2 in [RFC9256].
// SetFlagDropUponInvalid sets the bool value in the BgpAttributesBsidSrv6 object
func (obj *bgpAttributesBsidSrv6) SetFlagDropUponInvalid(value bool) BgpAttributesBsidSrv6 {

	obj.obj.FlagDropUponInvalid = &value
	return obj
}

// IPv6 address denoting the SRv6 SID.
// Ipv6Addr returns a string
func (obj *bgpAttributesBsidSrv6) Ipv6Addr() string {

	return *obj.obj.Ipv6Addr

}

// IPv6 address denoting the SRv6 SID.
// Ipv6Addr returns a string
func (obj *bgpAttributesBsidSrv6) HasIpv6Addr() bool {
	return obj.obj.Ipv6Addr != nil
}

// IPv6 address denoting the SRv6 SID.
// SetIpv6Addr sets the string value in the BgpAttributesBsidSrv6 object
func (obj *bgpAttributesBsidSrv6) SetIpv6Addr(value string) BgpAttributesBsidSrv6 {

	obj.obj.Ipv6Addr = &value
	return obj
}

func (obj *bgpAttributesBsidSrv6) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv6Addr != nil {

		err := obj.validateIpv6(obj.Ipv6Addr())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesBsidSrv6.Ipv6Addr"))
		}

	}

}

func (obj *bgpAttributesBsidSrv6) setDefault() {
	if obj.obj.FlagSpecifiedBsidOnly == nil {
		obj.SetFlagSpecifiedBsidOnly(false)
	}
	if obj.obj.FlagDropUponInvalid == nil {
		obj.SetFlagDropUponInvalid(false)
	}
	if obj.obj.Ipv6Addr == nil {
		obj.SetIpv6Addr("0::0")
	}

}

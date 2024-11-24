package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget *****
type bgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget struct {
	validation
	obj          *otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	marshaller   marshalBgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	unMarshaller unMarshalBgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
}

func NewBgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget() BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget {
	obj := bgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget{obj: &otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) msg() *otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget {
	return obj.obj
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) setMsg(msg *otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget struct {
	obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
}

type marshalBgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget interface {
	// ToProto marshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget to protobuf object *otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	ToProto() (*otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget, error)
	// ToPbText marshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget struct {
	obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
}

type unMarshalBgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget interface {
	// FromProto unmarshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget from protobuf object *otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	FromProto(msg *otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) (BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget, error)
	// FromPbText unmarshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget from JSON text
	FromJson(value string) error
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) Marshal() marshalBgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) Unmarshal() unMarshalBgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) ToProto() (*otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) FromProto(msg *otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) (BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) ToPbText() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) FromPbText(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) ToYaml() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) FromYaml(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) ToJsonRaw() (string, error) {
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

func (m *marshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) ToJson() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) FromJson(value string) error {
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

func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) Clone() (BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget()
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

// BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget is the Route Target Community identifies one or more routers that may receive a set of routes (that carry this Community) carried by BGP.  It is sent with sub-type as 0x02.
type BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget interface {
	Validation
	// msg marshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget to protobuf object *otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	// and doesn't set defaults
	msg() *otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	// setMsg unmarshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget from protobuf object *otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	// and doesn't set defaults
	setMsg(*otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	// provides marshal interface
	Marshal() marshalBgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	// provides unmarshal interface
	Unmarshal() unMarshalBgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	// validate validates BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// GlobalIpv4Admin returns string, set in BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget.
	GlobalIpv4Admin() string
	// SetGlobalIpv4Admin assigns string provided by user to BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	SetGlobalIpv4Admin(value string) BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	// HasGlobalIpv4Admin checks if GlobalIpv4Admin has been set in BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	HasGlobalIpv4Admin() bool
	// Local2ByteAdmin returns uint32, set in BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget.
	Local2ByteAdmin() uint32
	// SetLocal2ByteAdmin assigns uint32 provided by user to BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	SetLocal2ByteAdmin(value uint32) BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	// HasLocal2ByteAdmin checks if Local2ByteAdmin has been set in BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	HasLocal2ByteAdmin() bool
}

// An IPv4 unicast address assigned by one of the Internet registries.
// GlobalIpv4Admin returns a string
func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) GlobalIpv4Admin() string {

	return *obj.obj.GlobalIpv4Admin

}

// An IPv4 unicast address assigned by one of the Internet registries.
// GlobalIpv4Admin returns a string
func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) HasGlobalIpv4Admin() bool {
	return obj.obj.GlobalIpv4Admin != nil
}

// An IPv4 unicast address assigned by one of the Internet registries.
// SetGlobalIpv4Admin sets the string value in the BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget object
func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) SetGlobalIpv4Admin(value string) BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget {

	obj.obj.GlobalIpv4Admin = &value
	return obj
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which  the IP address carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local2ByteAdmin returns a uint32
func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) Local2ByteAdmin() uint32 {

	return *obj.obj.Local_2ByteAdmin

}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which  the IP address carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local2ByteAdmin returns a uint32
func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) HasLocal2ByteAdmin() bool {
	return obj.obj.Local_2ByteAdmin != nil
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which  the IP address carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// SetLocal2ByteAdmin sets the uint32 value in the BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget object
func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) SetLocal2ByteAdmin(value uint32) BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget {

	obj.obj.Local_2ByteAdmin = &value
	return obj
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.GlobalIpv4Admin != nil {

		err := obj.validateIpv4(obj.GlobalIpv4Admin())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget.GlobalIpv4Admin"))
		}

	}

	if obj.obj.Local_2ByteAdmin != nil {

		if *obj.obj.Local_2ByteAdmin > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget.Local_2ByteAdmin <= 65535 but Got %d", *obj.obj.Local_2ByteAdmin))
		}

	}

}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) setDefault() {
	if obj.obj.GlobalIpv4Admin == nil {
		obj.SetGlobalIpv4Admin("0.0.0.0")
	}
	if obj.obj.Local_2ByteAdmin == nil {
		obj.SetLocal2ByteAdmin(1)
	}

}

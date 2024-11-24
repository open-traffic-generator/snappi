package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin *****
type bgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin struct {
	validation
	obj          *otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	marshaller   marshalBgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	unMarshaller unMarshalBgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
}

func NewBgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin() BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin {
	obj := bgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin{obj: &otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) msg() *otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin {
	return obj.obj
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) setMsg(msg *otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin struct {
	obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
}

type marshalBgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin interface {
	// ToProto marshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin to protobuf object *otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	ToProto() (*otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin, error)
	// ToPbText marshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin struct {
	obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
}

type unMarshalBgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin interface {
	// FromProto unmarshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin from protobuf object *otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	FromProto(msg *otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) (BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin, error)
	// FromPbText unmarshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin from JSON text
	FromJson(value string) error
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) Marshal() marshalBgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) Unmarshal() unMarshalBgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) ToProto() (*otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) FromProto(msg *otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) (BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) ToPbText() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) FromPbText(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) ToYaml() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) FromYaml(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) ToJsonRaw() (string, error) {
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

func (m *marshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) ToJson() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) FromJson(value string) error {
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

func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) Clone() (BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin()
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

// BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin is the Route Origin Community identifies one or more routers that inject a set of routes (that carry this Community) into BGP It is sent with sub-type as 0x03.
type BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin interface {
	Validation
	// msg marshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin to protobuf object *otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	// and doesn't set defaults
	msg() *otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	// setMsg unmarshals BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin from protobuf object *otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	// and doesn't set defaults
	setMsg(*otg.BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	// provides marshal interface
	Marshal() marshalBgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	// provides unmarshal interface
	Unmarshal() unMarshalBgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	// validate validates BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// GlobalIpv4Admin returns string, set in BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin.
	GlobalIpv4Admin() string
	// SetGlobalIpv4Admin assigns string provided by user to BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	SetGlobalIpv4Admin(value string) BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	// HasGlobalIpv4Admin checks if GlobalIpv4Admin has been set in BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	HasGlobalIpv4Admin() bool
	// Local2ByteAdmin returns uint32, set in BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin.
	Local2ByteAdmin() uint32
	// SetLocal2ByteAdmin assigns uint32 provided by user to BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	SetLocal2ByteAdmin(value uint32) BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	// HasLocal2ByteAdmin checks if Local2ByteAdmin has been set in BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	HasLocal2ByteAdmin() bool
}

// An IPv4 unicast address assigned by one of the Internet registries.
// GlobalIpv4Admin returns a string
func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) GlobalIpv4Admin() string {

	return *obj.obj.GlobalIpv4Admin

}

// An IPv4 unicast address assigned by one of the Internet registries.
// GlobalIpv4Admin returns a string
func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) HasGlobalIpv4Admin() bool {
	return obj.obj.GlobalIpv4Admin != nil
}

// An IPv4 unicast address assigned by one of the Internet registries.
// SetGlobalIpv4Admin sets the string value in the BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin object
func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) SetGlobalIpv4Admin(value string) BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin {

	obj.obj.GlobalIpv4Admin = &value
	return obj
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which  the IP address carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local2ByteAdmin returns a uint32
func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) Local2ByteAdmin() uint32 {

	return *obj.obj.Local_2ByteAdmin

}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which  the IP address carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local2ByteAdmin returns a uint32
func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) HasLocal2ByteAdmin() bool {
	return obj.obj.Local_2ByteAdmin != nil
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which  the IP address carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// SetLocal2ByteAdmin sets the uint32 value in the BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin object
func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) SetLocal2ByteAdmin(value uint32) BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin {

	obj.obj.Local_2ByteAdmin = &value
	return obj
}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.GlobalIpv4Admin != nil {

		err := obj.validateIpv4(obj.GlobalIpv4Admin())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin.GlobalIpv4Admin"))
		}

	}

	if obj.obj.Local_2ByteAdmin != nil {

		if *obj.obj.Local_2ByteAdmin > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin.Local_2ByteAdmin <= 65535 but Got %d", *obj.obj.Local_2ByteAdmin))
		}

	}

}

func (obj *bgpExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) setDefault() {
	if obj.obj.GlobalIpv4Admin == nil {
		obj.SetGlobalIpv4Admin("0.0.0.0")
	}
	if obj.obj.Local_2ByteAdmin == nil {
		obj.SetLocal2ByteAdmin(1)
	}

}

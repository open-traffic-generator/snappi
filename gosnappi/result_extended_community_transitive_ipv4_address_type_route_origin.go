package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin *****
type resultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin struct {
	validation
	obj          *otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	marshaller   marshalResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	unMarshaller unMarshalResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
}

func NewResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin() ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin {
	obj := resultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin{obj: &otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin{}}
	obj.setDefault()
	return &obj
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) msg() *otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin {
	return obj.obj
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) setMsg(msg *otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin struct {
	obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
}

type marshalResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin interface {
	// ToProto marshals ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin to protobuf object *otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	ToProto() (*otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin, error)
	// ToPbText marshals ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin to YAML text
	ToYaml() (string, error)
	// ToJson marshals ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin to JSON text
	ToJson() (string, error)
}

type unMarshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin struct {
	obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
}

type unMarshalResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin interface {
	// FromProto unmarshals ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin from protobuf object *otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	FromProto(msg *otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) (ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin, error)
	// FromPbText unmarshals ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin from JSON text
	FromJson(value string) error
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) Marshal() marshalResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin {
	if obj.marshaller == nil {
		obj.marshaller = &marshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin{obj: obj}
	}
	return obj.marshaller
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) Unmarshal() unMarshalResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) ToProto() (*otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) FromProto(msg *otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) (ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) ToPbText() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) FromPbText(value string) error {
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

func (m *marshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) ToYaml() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) FromYaml(value string) error {
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

func (m *marshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) ToJson() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) FromJson(value string) error {
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

func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) Clone() (ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin()
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

// ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin is the Route Origin Community identifies one or more routers that inject a set of routes (that carry this Community) into BGP It is sent with sub-type as 0x03.
type ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin interface {
	Validation
	// msg marshals ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin to protobuf object *otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	// and doesn't set defaults
	msg() *otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	// setMsg unmarshals ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin from protobuf object *otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	// and doesn't set defaults
	setMsg(*otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	// provides marshal interface
	Marshal() marshalResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	// provides unmarshal interface
	Unmarshal() unMarshalResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	// validate validates ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// GlobalIpv4Admin returns string, set in ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin.
	GlobalIpv4Admin() string
	// SetGlobalIpv4Admin assigns string provided by user to ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	SetGlobalIpv4Admin(value string) ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	// HasGlobalIpv4Admin checks if GlobalIpv4Admin has been set in ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	HasGlobalIpv4Admin() bool
	// Local2ByteAdmin returns uint32, set in ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin.
	Local2ByteAdmin() uint32
	// SetLocal2ByteAdmin assigns uint32 provided by user to ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	SetLocal2ByteAdmin(value uint32) ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	// HasLocal2ByteAdmin checks if Local2ByteAdmin has been set in ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin
	HasLocal2ByteAdmin() bool
}

// An IPv4 unicast address assigned by one of the Internet registries.
// GlobalIpv4Admin returns a string
func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) GlobalIpv4Admin() string {

	return *obj.obj.GlobalIpv4Admin

}

// An IPv4 unicast address assigned by one of the Internet registries.
// GlobalIpv4Admin returns a string
func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) HasGlobalIpv4Admin() bool {
	return obj.obj.GlobalIpv4Admin != nil
}

// An IPv4 unicast address assigned by one of the Internet registries.
// SetGlobalIpv4Admin sets the string value in the ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin object
func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) SetGlobalIpv4Admin(value string) ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin {

	obj.obj.GlobalIpv4Admin = &value
	return obj
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which  the IP address carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local2ByteAdmin returns a uint32
func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) Local2ByteAdmin() uint32 {

	return *obj.obj.Local_2ByteAdmin

}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which  the IP address carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local2ByteAdmin returns a uint32
func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) HasLocal2ByteAdmin() bool {
	return obj.obj.Local_2ByteAdmin != nil
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which  the IP address carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// SetLocal2ByteAdmin sets the uint32 value in the ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin object
func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) SetLocal2ByteAdmin(value uint32) ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin {

	obj.obj.Local_2ByteAdmin = &value
	return obj
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.GlobalIpv4Admin != nil {

		err := obj.validateIpv4(obj.GlobalIpv4Admin())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin.GlobalIpv4Admin"))
		}

	}

	if obj.obj.Local_2ByteAdmin != nil {

		if *obj.obj.Local_2ByteAdmin > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= ResultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin.Local_2ByteAdmin <= 65535 but Got %d", *obj.obj.Local_2ByteAdmin))
		}

	}

}

func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteOrigin) setDefault() {

}

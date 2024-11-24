package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget *****
type resultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget struct {
	validation
	obj          *otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	marshaller   marshalResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	unMarshaller unMarshalResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
}

func NewResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget() ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget {
	obj := resultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget{obj: &otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget{}}
	obj.setDefault()
	return &obj
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) msg() *otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget {
	return obj.obj
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) setMsg(msg *otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget struct {
	obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
}

type marshalResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget interface {
	// ToProto marshals ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget to protobuf object *otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	ToProto() (*otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget, error)
	// ToPbText marshals ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget to YAML text
	ToYaml() (string, error)
	// ToJson marshals ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget struct {
	obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
}

type unMarshalResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget interface {
	// FromProto unmarshals ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget from protobuf object *otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	FromProto(msg *otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) (ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget, error)
	// FromPbText unmarshals ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget from JSON text
	FromJson(value string) error
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) Marshal() marshalResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget {
	if obj.marshaller == nil {
		obj.marshaller = &marshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget{obj: obj}
	}
	return obj.marshaller
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) Unmarshal() unMarshalResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) ToProto() (*otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) FromProto(msg *otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) (ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) ToPbText() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) FromPbText(value string) error {
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

func (m *marshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) ToYaml() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) FromYaml(value string) error {
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

func (m *marshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) ToJsonRaw() (string, error) {
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

func (m *marshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) ToJson() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) FromJson(value string) error {
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

func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) Clone() (ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget()
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

// ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget is the Route Target Community identifies one or more routers that may receive a set of routes (that carry this Community) carried by BGP.  It is sent with sub-type as 0x02.
type ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget interface {
	Validation
	// msg marshals ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget to protobuf object *otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	// and doesn't set defaults
	msg() *otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	// setMsg unmarshals ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget from protobuf object *otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	// and doesn't set defaults
	setMsg(*otg.ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	// provides marshal interface
	Marshal() marshalResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	// provides unmarshal interface
	Unmarshal() unMarshalResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	// validate validates ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// GlobalIpv4Admin returns string, set in ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget.
	GlobalIpv4Admin() string
	// SetGlobalIpv4Admin assigns string provided by user to ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	SetGlobalIpv4Admin(value string) ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	// HasGlobalIpv4Admin checks if GlobalIpv4Admin has been set in ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	HasGlobalIpv4Admin() bool
	// Local2ByteAdmin returns uint32, set in ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget.
	Local2ByteAdmin() uint32
	// SetLocal2ByteAdmin assigns uint32 provided by user to ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	SetLocal2ByteAdmin(value uint32) ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	// HasLocal2ByteAdmin checks if Local2ByteAdmin has been set in ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget
	HasLocal2ByteAdmin() bool
}

// An IPv4 unicast address assigned by one of the Internet registries.
// GlobalIpv4Admin returns a string
func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) GlobalIpv4Admin() string {

	return *obj.obj.GlobalIpv4Admin

}

// An IPv4 unicast address assigned by one of the Internet registries.
// GlobalIpv4Admin returns a string
func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) HasGlobalIpv4Admin() bool {
	return obj.obj.GlobalIpv4Admin != nil
}

// An IPv4 unicast address assigned by one of the Internet registries.
// SetGlobalIpv4Admin sets the string value in the ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget object
func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) SetGlobalIpv4Admin(value string) ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget {

	obj.obj.GlobalIpv4Admin = &value
	return obj
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which  the IP address carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local2ByteAdmin returns a uint32
func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) Local2ByteAdmin() uint32 {

	return *obj.obj.Local_2ByteAdmin

}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which  the IP address carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local2ByteAdmin returns a uint32
func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) HasLocal2ByteAdmin() bool {
	return obj.obj.Local_2ByteAdmin != nil
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which  the IP address carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// SetLocal2ByteAdmin sets the uint32 value in the ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget object
func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) SetLocal2ByteAdmin(value uint32) ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget {

	obj.obj.Local_2ByteAdmin = &value
	return obj
}

func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.GlobalIpv4Admin != nil {

		err := obj.validateIpv4(obj.GlobalIpv4Admin())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget.GlobalIpv4Admin"))
		}

	}

	if obj.obj.Local_2ByteAdmin != nil {

		if *obj.obj.Local_2ByteAdmin > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= ResultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget.Local_2ByteAdmin <= 65535 but Got %d", *obj.obj.Local_2ByteAdmin))
		}

	}

}

func (obj *resultExtendedCommunityTransitiveIpv4AddressTypeRouteTarget) setDefault() {

}

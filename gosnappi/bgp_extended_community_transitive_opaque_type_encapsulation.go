package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation *****
type bgpExtendedCommunityTransitiveOpaqueTypeEncapsulation struct {
	validation
	obj          *otg.BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
	marshaller   marshalBgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
	unMarshaller unMarshalBgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
}

func NewBgpExtendedCommunityTransitiveOpaqueTypeEncapsulation() BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation {
	obj := bgpExtendedCommunityTransitiveOpaqueTypeEncapsulation{obj: &otg.BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) msg() *otg.BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation {
	return obj.obj
}

func (obj *bgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) setMsg(msg *otg.BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpExtendedCommunityTransitiveOpaqueTypeEncapsulation struct {
	obj *bgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
}

type marshalBgpExtendedCommunityTransitiveOpaqueTypeEncapsulation interface {
	// ToProto marshals BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation to protobuf object *otg.BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
	ToProto() (*otg.BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation, error)
	// ToPbText marshals BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpExtendedCommunityTransitiveOpaqueTypeEncapsulation struct {
	obj *bgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
}

type unMarshalBgpExtendedCommunityTransitiveOpaqueTypeEncapsulation interface {
	// FromProto unmarshals BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation from protobuf object *otg.BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
	FromProto(msg *otg.BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) (BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation, error)
	// FromPbText unmarshals BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation from JSON text
	FromJson(value string) error
}

func (obj *bgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) Marshal() marshalBgpExtendedCommunityTransitiveOpaqueTypeEncapsulation {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpExtendedCommunityTransitiveOpaqueTypeEncapsulation{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) Unmarshal() unMarshalBgpExtendedCommunityTransitiveOpaqueTypeEncapsulation {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpExtendedCommunityTransitiveOpaqueTypeEncapsulation{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) ToProto() (*otg.BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) FromProto(msg *otg.BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) (BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) ToPbText() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) FromPbText(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) ToYaml() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) FromYaml(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) ToJsonRaw() (string, error) {
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

func (m *marshalbgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) ToJson() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) FromJson(value string) error {
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

func (obj *bgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) Clone() (BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpExtendedCommunityTransitiveOpaqueTypeEncapsulation()
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

// BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation is this identifies the type of tunneling technology being signalled. It is defined in RFC9012 and sent with sub-type as 0x0c.
type BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation interface {
	Validation
	// msg marshals BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation to protobuf object *otg.BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
	// and doesn't set defaults
	msg() *otg.BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
	// setMsg unmarshals BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation from protobuf object *otg.BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
	// and doesn't set defaults
	setMsg(*otg.BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
	// provides marshal interface
	Marshal() marshalBgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
	// provides unmarshal interface
	Unmarshal() unMarshalBgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
	// validate validates BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Reserved returns uint32, set in BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation.
	Reserved() uint32
	// SetReserved assigns uint32 provided by user to BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
	SetReserved(value uint32) BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
	// HasReserved checks if Reserved has been set in BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
	HasReserved() bool
	// TunnelType returns uint32, set in BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation.
	TunnelType() uint32
	// SetTunnelType assigns uint32 provided by user to BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
	SetTunnelType(value uint32) BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
	// HasTunnelType checks if TunnelType has been set in BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation
	HasTunnelType() bool
}

// Four bytes of reserved values. Normally set to 0 on transmit and ignored on receive.
// Reserved returns a uint32
func (obj *bgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) Reserved() uint32 {

	return *obj.obj.Reserved

}

// Four bytes of reserved values. Normally set to 0 on transmit and ignored on receive.
// Reserved returns a uint32
func (obj *bgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// Four bytes of reserved values. Normally set to 0 on transmit and ignored on receive.
// SetReserved sets the uint32 value in the BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation object
func (obj *bgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) SetReserved(value uint32) BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation {

	obj.obj.Reserved = &value
	return obj
}

// Identifies the type of tunneling technology being signalled. Initially defined in RFC5512 and extended in RFC9012. Some of the important tunnel types include  1 L2TPv3 over IP	[RFC9012],
// 2	GRE	[RFC9012]
// 7	IP in IP	[RFC9012]
// 8	VXLAN Encapsulation	[RFC8365]
// 9	NVGRE Encapsulation	[RFC8365]
// 10	MPLS Encapsulation	[RFC8365]
// 15	SR TE Policy Type	[draft-ietf-idr-segment-routing-te-policy]
// 19	Geneve Encapsulation	[RFC8926]
// TunnelType returns a uint32
func (obj *bgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) TunnelType() uint32 {

	return *obj.obj.TunnelType

}

// Identifies the type of tunneling technology being signalled. Initially defined in RFC5512 and extended in RFC9012. Some of the important tunnel types include  1 L2TPv3 over IP	[RFC9012],
// 2	GRE	[RFC9012]
// 7	IP in IP	[RFC9012]
// 8	VXLAN Encapsulation	[RFC8365]
// 9	NVGRE Encapsulation	[RFC8365]
// 10	MPLS Encapsulation	[RFC8365]
// 15	SR TE Policy Type	[draft-ietf-idr-segment-routing-te-policy]
// 19	Geneve Encapsulation	[RFC8926]
// TunnelType returns a uint32
func (obj *bgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) HasTunnelType() bool {
	return obj.obj.TunnelType != nil
}

// Identifies the type of tunneling technology being signalled. Initially defined in RFC5512 and extended in RFC9012. Some of the important tunnel types include  1 L2TPv3 over IP	[RFC9012],
// 2	GRE	[RFC9012]
// 7	IP in IP	[RFC9012]
// 8	VXLAN Encapsulation	[RFC8365]
// 9	NVGRE Encapsulation	[RFC8365]
// 10	MPLS Encapsulation	[RFC8365]
// 15	SR TE Policy Type	[draft-ietf-idr-segment-routing-te-policy]
// 19	Geneve Encapsulation	[RFC8926]
// SetTunnelType sets the uint32 value in the BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation object
func (obj *bgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) SetTunnelType(value uint32) BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation {

	obj.obj.TunnelType = &value
	return obj
}

func (obj *bgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TunnelType != nil {

		if *obj.obj.TunnelType > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpExtendedCommunityTransitiveOpaqueTypeEncapsulation.TunnelType <= 65535 but Got %d", *obj.obj.TunnelType))
		}

	}

}

func (obj *bgpExtendedCommunityTransitiveOpaqueTypeEncapsulation) setDefault() {
	if obj.obj.Reserved == nil {
		obj.SetReserved(0)
	}
	if obj.obj.TunnelType == nil {
		obj.SetTunnelType(1)
	}

}

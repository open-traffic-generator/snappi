package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation *****
type resultExtendedCommunityTransitiveOpaqueTypeEncapsulation struct {
	validation
	obj          *otg.ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
	marshaller   marshalResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
	unMarshaller unMarshalResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
}

func NewResultExtendedCommunityTransitiveOpaqueTypeEncapsulation() ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation {
	obj := resultExtendedCommunityTransitiveOpaqueTypeEncapsulation{obj: &otg.ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation{}}
	obj.setDefault()
	return &obj
}

func (obj *resultExtendedCommunityTransitiveOpaqueTypeEncapsulation) msg() *otg.ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation {
	return obj.obj
}

func (obj *resultExtendedCommunityTransitiveOpaqueTypeEncapsulation) setMsg(msg *otg.ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation) ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalresultExtendedCommunityTransitiveOpaqueTypeEncapsulation struct {
	obj *resultExtendedCommunityTransitiveOpaqueTypeEncapsulation
}

type marshalResultExtendedCommunityTransitiveOpaqueTypeEncapsulation interface {
	// ToProto marshals ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation to protobuf object *otg.ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
	ToProto() (*otg.ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation, error)
	// ToPbText marshals ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation to YAML text
	ToYaml() (string, error)
	// ToJson marshals ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation to JSON text
	ToJson() (string, error)
}

type unMarshalresultExtendedCommunityTransitiveOpaqueTypeEncapsulation struct {
	obj *resultExtendedCommunityTransitiveOpaqueTypeEncapsulation
}

type unMarshalResultExtendedCommunityTransitiveOpaqueTypeEncapsulation interface {
	// FromProto unmarshals ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation from protobuf object *otg.ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
	FromProto(msg *otg.ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation) (ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation, error)
	// FromPbText unmarshals ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation from JSON text
	FromJson(value string) error
}

func (obj *resultExtendedCommunityTransitiveOpaqueTypeEncapsulation) Marshal() marshalResultExtendedCommunityTransitiveOpaqueTypeEncapsulation {
	if obj.marshaller == nil {
		obj.marshaller = &marshalresultExtendedCommunityTransitiveOpaqueTypeEncapsulation{obj: obj}
	}
	return obj.marshaller
}

func (obj *resultExtendedCommunityTransitiveOpaqueTypeEncapsulation) Unmarshal() unMarshalResultExtendedCommunityTransitiveOpaqueTypeEncapsulation {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalresultExtendedCommunityTransitiveOpaqueTypeEncapsulation{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalresultExtendedCommunityTransitiveOpaqueTypeEncapsulation) ToProto() (*otg.ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalresultExtendedCommunityTransitiveOpaqueTypeEncapsulation) FromProto(msg *otg.ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation) (ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalresultExtendedCommunityTransitiveOpaqueTypeEncapsulation) ToPbText() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitiveOpaqueTypeEncapsulation) FromPbText(value string) error {
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

func (m *marshalresultExtendedCommunityTransitiveOpaqueTypeEncapsulation) ToYaml() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitiveOpaqueTypeEncapsulation) FromYaml(value string) error {
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

func (m *marshalresultExtendedCommunityTransitiveOpaqueTypeEncapsulation) ToJson() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitiveOpaqueTypeEncapsulation) FromJson(value string) error {
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

func (obj *resultExtendedCommunityTransitiveOpaqueTypeEncapsulation) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitiveOpaqueTypeEncapsulation) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitiveOpaqueTypeEncapsulation) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *resultExtendedCommunityTransitiveOpaqueTypeEncapsulation) Clone() (ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewResultExtendedCommunityTransitiveOpaqueTypeEncapsulation()
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

// ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation is this identifies the type of tunneling technology being signalled. It is defined in RFC9012 and sent with sub-type as 0x0c.
type ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation interface {
	Validation
	// msg marshals ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation to protobuf object *otg.ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
	// and doesn't set defaults
	msg() *otg.ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
	// setMsg unmarshals ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation from protobuf object *otg.ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
	// and doesn't set defaults
	setMsg(*otg.ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation) ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
	// provides marshal interface
	Marshal() marshalResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
	// provides unmarshal interface
	Unmarshal() unMarshalResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
	// validate validates ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Reserved returns uint32, set in ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation.
	Reserved() uint32
	// SetReserved assigns uint32 provided by user to ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
	SetReserved(value uint32) ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
	// HasReserved checks if Reserved has been set in ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
	HasReserved() bool
	// TunnelType returns uint32, set in ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation.
	TunnelType() uint32
	// SetTunnelType assigns uint32 provided by user to ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
	SetTunnelType(value uint32) ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
	// HasTunnelType checks if TunnelType has been set in ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
	HasTunnelType() bool
}

// Four bytes of reserved values. Normally set to 0 on transmit and ignored on receive.
// Reserved returns a uint32
func (obj *resultExtendedCommunityTransitiveOpaqueTypeEncapsulation) Reserved() uint32 {

	return *obj.obj.Reserved

}

// Four bytes of reserved values. Normally set to 0 on transmit and ignored on receive.
// Reserved returns a uint32
func (obj *resultExtendedCommunityTransitiveOpaqueTypeEncapsulation) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// Four bytes of reserved values. Normally set to 0 on transmit and ignored on receive.
// SetReserved sets the uint32 value in the ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation object
func (obj *resultExtendedCommunityTransitiveOpaqueTypeEncapsulation) SetReserved(value uint32) ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation {

	obj.obj.Reserved = &value
	return obj
}

// Identifies the type of tunneling technology being signalled. Initially defined in RFC5512 and extended in RFC9012.
// Some of the important tunnel types include
// - 1 L2TPv3 over IP	[RFC9012],
// - 2	GRE	[RFC9012],
// - 7	IP in IP	[RFC9012],
// - 8	VXLAN Encapsulation	[RFC8365],
// - 9	NVGRE Encapsulation	[RFC8365],
// - 10	MPLS Encapsulation	[RFC8365],
// - 15	SR TE Policy Type	[draft-ietf-idr-segment-routing-te-policy],
// - 19	Geneve Encapsulation	[RFC8926]
// TunnelType returns a uint32
func (obj *resultExtendedCommunityTransitiveOpaqueTypeEncapsulation) TunnelType() uint32 {

	return *obj.obj.TunnelType

}

// Identifies the type of tunneling technology being signalled. Initially defined in RFC5512 and extended in RFC9012.
// Some of the important tunnel types include
// - 1 L2TPv3 over IP	[RFC9012],
// - 2	GRE	[RFC9012],
// - 7	IP in IP	[RFC9012],
// - 8	VXLAN Encapsulation	[RFC8365],
// - 9	NVGRE Encapsulation	[RFC8365],
// - 10	MPLS Encapsulation	[RFC8365],
// - 15	SR TE Policy Type	[draft-ietf-idr-segment-routing-te-policy],
// - 19	Geneve Encapsulation	[RFC8926]
// TunnelType returns a uint32
func (obj *resultExtendedCommunityTransitiveOpaqueTypeEncapsulation) HasTunnelType() bool {
	return obj.obj.TunnelType != nil
}

// Identifies the type of tunneling technology being signalled. Initially defined in RFC5512 and extended in RFC9012.
// Some of the important tunnel types include
// - 1 L2TPv3 over IP	[RFC9012],
// - 2	GRE	[RFC9012],
// - 7	IP in IP	[RFC9012],
// - 8	VXLAN Encapsulation	[RFC8365],
// - 9	NVGRE Encapsulation	[RFC8365],
// - 10	MPLS Encapsulation	[RFC8365],
// - 15	SR TE Policy Type	[draft-ietf-idr-segment-routing-te-policy],
// - 19	Geneve Encapsulation	[RFC8926]
// SetTunnelType sets the uint32 value in the ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation object
func (obj *resultExtendedCommunityTransitiveOpaqueTypeEncapsulation) SetTunnelType(value uint32) ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation {

	obj.obj.TunnelType = &value
	return obj
}

func (obj *resultExtendedCommunityTransitiveOpaqueTypeEncapsulation) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TunnelType != nil {

		if *obj.obj.TunnelType > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation.TunnelType <= 65535 but Got %d", *obj.obj.TunnelType))
		}

	}

}

func (obj *resultExtendedCommunityTransitiveOpaqueTypeEncapsulation) setDefault() {

}

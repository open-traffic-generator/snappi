package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth *****
type bgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth struct {
	validation
	obj          *otg.BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	marshaller   marshalBgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	unMarshaller unMarshalBgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
}

func NewBgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth() BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth {
	obj := bgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth{obj: &otg.BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) msg() *otg.BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth {
	return obj.obj
}

func (obj *bgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) setMsg(msg *otg.BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth struct {
	obj *bgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
}

type marshalBgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth interface {
	// ToProto marshals BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth to protobuf object *otg.BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	ToProto() (*otg.BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth, error)
	// ToPbText marshals BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth to JSON text
	ToJson() (string, error)
}

type unMarshalbgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth struct {
	obj *bgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
}

type unMarshalBgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth interface {
	// FromProto unmarshals BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth from protobuf object *otg.BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	FromProto(msg *otg.BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) (BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth, error)
	// FromPbText unmarshals BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth from JSON text
	FromJson(value string) error
}

func (obj *bgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) Marshal() marshalBgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) Unmarshal() unMarshalBgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) ToProto() (*otg.BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) FromProto(msg *otg.BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) (BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) ToPbText() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) FromPbText(value string) error {
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

func (m *marshalbgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) ToYaml() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) FromYaml(value string) error {
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

func (m *marshalbgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) ToJson() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) FromJson(value string) error {
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

func (obj *bgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) Clone() (BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth()
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

// BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth is the Link Bandwidth Extended Community attribute is defined in draft-ietf-idr-link-bandwidth. It is sent with sub-type as 0x04.
type BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth interface {
	Validation
	// msg marshals BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth to protobuf object *otg.BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	// and doesn't set defaults
	msg() *otg.BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	// setMsg unmarshals BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth from protobuf object *otg.BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	// and doesn't set defaults
	setMsg(*otg.BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	// provides marshal interface
	Marshal() marshalBgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	// provides unmarshal interface
	Unmarshal() unMarshalBgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	// validate validates BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Global2ByteAs returns uint32, set in BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth.
	Global2ByteAs() uint32
	// SetGlobal2ByteAs assigns uint32 provided by user to BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	SetGlobal2ByteAs(value uint32) BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	// HasGlobal2ByteAs checks if Global2ByteAs has been set in BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	HasGlobal2ByteAs() bool
	// Bandwidth returns float32, set in BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth.
	Bandwidth() float32
	// SetBandwidth assigns float32 provided by user to BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	SetBandwidth(value float32) BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	// HasBandwidth checks if Bandwidth has been set in BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	HasBandwidth() bool
}

// The value of the Global Administrator subfield should represent the Autonomous System of the router that attaches the Link Bandwidth Community. If four octet AS numbering scheme is used, AS_TRANS (23456) should be used.
// Global2ByteAs returns a uint32
func (obj *bgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) Global2ByteAs() uint32 {

	return *obj.obj.Global_2ByteAs

}

// The value of the Global Administrator subfield should represent the Autonomous System of the router that attaches the Link Bandwidth Community. If four octet AS numbering scheme is used, AS_TRANS (23456) should be used.
// Global2ByteAs returns a uint32
func (obj *bgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) HasGlobal2ByteAs() bool {
	return obj.obj.Global_2ByteAs != nil
}

// The value of the Global Administrator subfield should represent the Autonomous System of the router that attaches the Link Bandwidth Community. If four octet AS numbering scheme is used, AS_TRANS (23456) should be used.
// SetGlobal2ByteAs sets the uint32 value in the BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth object
func (obj *bgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) SetGlobal2ByteAs(value uint32) BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth {

	obj.obj.Global_2ByteAs = &value
	return obj
}

// Bandwidth of the link in bytes per second. ( 1 Kbps is 1000 bytes per second and 1 Mbps is 1000 Kbps per second )
// Bandwidth returns a float32
func (obj *bgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) Bandwidth() float32 {

	return *obj.obj.Bandwidth

}

// Bandwidth of the link in bytes per second. ( 1 Kbps is 1000 bytes per second and 1 Mbps is 1000 Kbps per second )
// Bandwidth returns a float32
func (obj *bgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) HasBandwidth() bool {
	return obj.obj.Bandwidth != nil
}

// Bandwidth of the link in bytes per second. ( 1 Kbps is 1000 bytes per second and 1 Mbps is 1000 Kbps per second )
// SetBandwidth sets the float32 value in the BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth object
func (obj *bgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) SetBandwidth(value float32) BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth {

	obj.obj.Bandwidth = &value
	return obj
}

func (obj *bgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Global_2ByteAs != nil {

		if *obj.obj.Global_2ByteAs > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth.Global_2ByteAs <= 65535 but Got %d", *obj.obj.Global_2ByteAs))
		}

	}

}

func (obj *bgpExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) setDefault() {
	if obj.obj.Global_2ByteAs == nil {
		obj.SetGlobal2ByteAs(100)
	}
	if obj.obj.Bandwidth == nil {
		obj.SetBandwidth(0)
	}

}

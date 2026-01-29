package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth *****
type bgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth struct {
	validation
	obj          *otg.BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	marshaller   marshalBgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	unMarshaller unMarshalBgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
}

func NewBgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth() BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth {
	obj := bgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth{obj: &otg.BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) msg() *otg.BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth {
	return obj.obj
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) setMsg(msg *otg.BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth struct {
	obj *bgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
}

type marshalBgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth interface {
	// ToProto marshals BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth to protobuf object *otg.BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	ToProto() (*otg.BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth, error)
	// ToPbText marshals BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth to JSON text
	ToJson() (string, error)
}

type unMarshalbgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth struct {
	obj *bgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
}

type unMarshalBgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth interface {
	// FromProto unmarshals BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth from protobuf object *otg.BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	FromProto(msg *otg.BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) (BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth, error)
	// FromPbText unmarshals BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth from JSON text
	FromJson(value string) error
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) Marshal() marshalBgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) Unmarshal() unMarshalBgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) ToProto() (*otg.BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) FromProto(msg *otg.BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) (BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) ToPbText() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) FromPbText(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) ToYaml() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) FromYaml(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) ToJson() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) FromJson(value string) error {
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

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) Clone() (BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth()
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

// BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth is the Link Bandwidth Extended Community attribute is defined in draft-ietf-idr-link-bandwidth. (https://datatracker.ietf.org/doc/draft-ietf-idr-link-bandwidth) It is sent with sub-type as 0x04.
type BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth interface {
	Validation
	// msg marshals BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth to protobuf object *otg.BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	// and doesn't set defaults
	msg() *otg.BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	// setMsg unmarshals BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth from protobuf object *otg.BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	// and doesn't set defaults
	setMsg(*otg.BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	// provides marshal interface
	Marshal() marshalBgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	// provides unmarshal interface
	Unmarshal() unMarshalBgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	// validate validates BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Global2ByteAs returns uint32, set in BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth.
	Global2ByteAs() uint32
	// SetGlobal2ByteAs assigns uint32 provided by user to BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	SetGlobal2ByteAs(value uint32) BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	// HasGlobal2ByteAs checks if Global2ByteAs has been set in BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	HasGlobal2ByteAs() bool
	// Bandwidth returns float32, set in BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth.
	Bandwidth() float32
	// SetBandwidth assigns float32 provided by user to BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	SetBandwidth(value float32) BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	// HasBandwidth checks if Bandwidth has been set in BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	HasBandwidth() bool
}

// The value of the Global Administrator subfield should represent the Autonomous System of the router that attaches the Link Bandwidth Community. If four octet AS numbering scheme is used, AS_TRANS (23456) should be used.
// Global2ByteAs returns a uint32
func (obj *bgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) Global2ByteAs() uint32 {

	return *obj.obj.Global_2ByteAs

}

// The value of the Global Administrator subfield should represent the Autonomous System of the router that attaches the Link Bandwidth Community. If four octet AS numbering scheme is used, AS_TRANS (23456) should be used.
// Global2ByteAs returns a uint32
func (obj *bgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) HasGlobal2ByteAs() bool {
	return obj.obj.Global_2ByteAs != nil
}

// The value of the Global Administrator subfield should represent the Autonomous System of the router that attaches the Link Bandwidth Community. If four octet AS numbering scheme is used, AS_TRANS (23456) should be used.
// SetGlobal2ByteAs sets the uint32 value in the BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth object
func (obj *bgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) SetGlobal2ByteAs(value uint32) BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth {

	obj.obj.Global_2ByteAs = &value
	return obj
}

// Bandwidth of the link in bytes per second. ( 1 Kbps is 1000 bytes per second and 1 Mbps is 1000 Kbps per second )
// Bandwidth returns a float32
func (obj *bgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) Bandwidth() float32 {

	return *obj.obj.Bandwidth

}

// Bandwidth of the link in bytes per second. ( 1 Kbps is 1000 bytes per second and 1 Mbps is 1000 Kbps per second )
// Bandwidth returns a float32
func (obj *bgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) HasBandwidth() bool {
	return obj.obj.Bandwidth != nil
}

// Bandwidth of the link in bytes per second. ( 1 Kbps is 1000 bytes per second and 1 Mbps is 1000 Kbps per second )
// SetBandwidth sets the float32 value in the BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth object
func (obj *bgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) SetBandwidth(value float32) BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth {

	obj.obj.Bandwidth = &value
	return obj
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Global_2ByteAs != nil {

		if *obj.obj.Global_2ByteAs > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth.Global_2ByteAs <= 65535 but Got %d", *obj.obj.Global_2ByteAs))
		}

	}

}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) setDefault() {
	if obj.obj.Global_2ByteAs == nil {
		obj.SetGlobal2ByteAs(100)
	}
	if obj.obj.Bandwidth == nil {
		obj.SetBandwidth(0)
	}

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth *****
type resultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth struct {
	validation
	obj          *otg.ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	marshaller   marshalResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	unMarshaller unMarshalResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
}

func NewResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth() ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth {
	obj := resultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth{obj: &otg.ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth{}}
	obj.setDefault()
	return &obj
}

func (obj *resultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) msg() *otg.ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth {
	return obj.obj
}

func (obj *resultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) setMsg(msg *otg.ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalresultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth struct {
	obj *resultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
}

type marshalResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth interface {
	// ToProto marshals ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth to protobuf object *otg.ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	ToProto() (*otg.ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth, error)
	// ToPbText marshals ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth to YAML text
	ToYaml() (string, error)
	// ToJson marshals ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth to JSON text
	ToJson() (string, error)
}

type unMarshalresultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth struct {
	obj *resultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
}

type unMarshalResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth interface {
	// FromProto unmarshals ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth from protobuf object *otg.ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	FromProto(msg *otg.ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) (ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth, error)
	// FromPbText unmarshals ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth from JSON text
	FromJson(value string) error
}

func (obj *resultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) Marshal() marshalResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth {
	if obj.marshaller == nil {
		obj.marshaller = &marshalresultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth{obj: obj}
	}
	return obj.marshaller
}

func (obj *resultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) Unmarshal() unMarshalResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalresultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalresultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) ToProto() (*otg.ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalresultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) FromProto(msg *otg.ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) (ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalresultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) ToPbText() (string, error) {
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

func (m *unMarshalresultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) FromPbText(value string) error {
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

func (m *marshalresultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) ToYaml() (string, error) {
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

func (m *unMarshalresultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) FromYaml(value string) error {
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

func (m *marshalresultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) ToJson() (string, error) {
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

func (m *unMarshalresultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) FromJson(value string) error {
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

func (obj *resultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *resultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) Clone() (ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth()
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

// ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth is the Link Bandwidth Extended Community attribute is defined in draft-ietf-idr-link-bandwidth. It is sent with sub-type as 0x04.
type ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth interface {
	Validation
	// msg marshals ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth to protobuf object *otg.ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	// and doesn't set defaults
	msg() *otg.ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	// setMsg unmarshals ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth from protobuf object *otg.ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	// and doesn't set defaults
	setMsg(*otg.ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	// provides marshal interface
	Marshal() marshalResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	// provides unmarshal interface
	Unmarshal() unMarshalResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	// validate validates ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Global2ByteAs returns uint32, set in ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth.
	Global2ByteAs() uint32
	// SetGlobal2ByteAs assigns uint32 provided by user to ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	SetGlobal2ByteAs(value uint32) ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	// HasGlobal2ByteAs checks if Global2ByteAs has been set in ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	HasGlobal2ByteAs() bool
	// Bandwidth returns float32, set in ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth.
	Bandwidth() float32
	// SetBandwidth assigns float32 provided by user to ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	SetBandwidth(value float32) ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	// HasBandwidth checks if Bandwidth has been set in ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	HasBandwidth() bool
}

// The value of the Global Administrator subfield should represent the Autonomous System of the router that attaches the Link Bandwidth Community. If four octet AS numbering scheme is used, AS_TRANS (23456) should be used.
// Global2ByteAs returns a uint32
func (obj *resultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) Global2ByteAs() uint32 {

	return *obj.obj.Global_2ByteAs

}

// The value of the Global Administrator subfield should represent the Autonomous System of the router that attaches the Link Bandwidth Community. If four octet AS numbering scheme is used, AS_TRANS (23456) should be used.
// Global2ByteAs returns a uint32
func (obj *resultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) HasGlobal2ByteAs() bool {
	return obj.obj.Global_2ByteAs != nil
}

// The value of the Global Administrator subfield should represent the Autonomous System of the router that attaches the Link Bandwidth Community. If four octet AS numbering scheme is used, AS_TRANS (23456) should be used.
// SetGlobal2ByteAs sets the uint32 value in the ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth object
func (obj *resultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) SetGlobal2ByteAs(value uint32) ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth {

	obj.obj.Global_2ByteAs = &value
	return obj
}

// Bandwidth of the link in bytes per second. ( 1 Kbps is 1000 bytes per second and 1 Mbps is 1000 Kbps per second )
// Bandwidth returns a float32
func (obj *resultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) Bandwidth() float32 {

	return *obj.obj.Bandwidth

}

// Bandwidth of the link in bytes per second. ( 1 Kbps is 1000 bytes per second and 1 Mbps is 1000 Kbps per second )
// Bandwidth returns a float32
func (obj *resultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) HasBandwidth() bool {
	return obj.obj.Bandwidth != nil
}

// Bandwidth of the link in bytes per second. ( 1 Kbps is 1000 bytes per second and 1 Mbps is 1000 Kbps per second )
// SetBandwidth sets the float32 value in the ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth object
func (obj *resultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) SetBandwidth(value float32) ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth {

	obj.obj.Bandwidth = &value
	return obj
}

func (obj *resultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Global_2ByteAs != nil {

		if *obj.obj.Global_2ByteAs > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth.Global_2ByteAs <= 65535 but Got %d", *obj.obj.Global_2ByteAs))
		}

	}

}

func (obj *resultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) setDefault() {

}

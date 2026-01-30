package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth *****
type resultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth struct {
	validation
	obj          *otg.ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	marshaller   marshalResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	unMarshaller unMarshalResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
}

func NewResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth() ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth {
	obj := resultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth{obj: &otg.ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth{}}
	obj.setDefault()
	return &obj
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) msg() *otg.ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth {
	return obj.obj
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) setMsg(msg *otg.ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalresultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth struct {
	obj *resultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
}

type marshalResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth interface {
	// ToProto marshals ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth to protobuf object *otg.ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	ToProto() (*otg.ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth, error)
	// ToPbText marshals ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth to YAML text
	ToYaml() (string, error)
	// ToJson marshals ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth to JSON text
	ToJson() (string, error)
}

type unMarshalresultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth struct {
	obj *resultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
}

type unMarshalResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth interface {
	// FromProto unmarshals ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth from protobuf object *otg.ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	FromProto(msg *otg.ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) (ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth, error)
	// FromPbText unmarshals ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth from JSON text
	FromJson(value string) error
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) Marshal() marshalResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth {
	if obj.marshaller == nil {
		obj.marshaller = &marshalresultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth{obj: obj}
	}
	return obj.marshaller
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) Unmarshal() unMarshalResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalresultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalresultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) ToProto() (*otg.ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalresultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) FromProto(msg *otg.ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) (ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalresultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) ToPbText() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) FromPbText(value string) error {
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

func (m *marshalresultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) ToYaml() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) FromYaml(value string) error {
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

func (m *marshalresultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) ToJson() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) FromJson(value string) error {
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

func (obj *resultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) Clone() (ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth()
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

// ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth is the Link Bandwidth Extended Community attribute is defined in draft-ietf-idr-link-bandwidth. (https://datatracker.ietf.org/doc/draft-ietf-idr-link-bandwidth) It is sent with sub-type as 0x04.
type ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth interface {
	Validation
	// msg marshals ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth to protobuf object *otg.ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	// and doesn't set defaults
	msg() *otg.ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	// setMsg unmarshals ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth from protobuf object *otg.ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	// and doesn't set defaults
	setMsg(*otg.ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	// provides marshal interface
	Marshal() marshalResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	// provides unmarshal interface
	Unmarshal() unMarshalResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	// validate validates ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Global2ByteAs returns uint32, set in ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth.
	Global2ByteAs() uint32
	// SetGlobal2ByteAs assigns uint32 provided by user to ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	SetGlobal2ByteAs(value uint32) ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	// HasGlobal2ByteAs checks if Global2ByteAs has been set in ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	HasGlobal2ByteAs() bool
	// Bandwidth returns float32, set in ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth.
	Bandwidth() float32
	// SetBandwidth assigns float32 provided by user to ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	SetBandwidth(value float32) ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	// HasBandwidth checks if Bandwidth has been set in ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth
	HasBandwidth() bool
}

// The value of the Global Administrator subfield should represent the Autonomous System of the router that attaches the Link Bandwidth Community. If four octet AS numbering scheme is used, AS_TRANS (23456) should be used.
// Global2ByteAs returns a uint32
func (obj *resultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) Global2ByteAs() uint32 {

	return *obj.obj.Global_2ByteAs

}

// The value of the Global Administrator subfield should represent the Autonomous System of the router that attaches the Link Bandwidth Community. If four octet AS numbering scheme is used, AS_TRANS (23456) should be used.
// Global2ByteAs returns a uint32
func (obj *resultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) HasGlobal2ByteAs() bool {
	return obj.obj.Global_2ByteAs != nil
}

// The value of the Global Administrator subfield should represent the Autonomous System of the router that attaches the Link Bandwidth Community. If four octet AS numbering scheme is used, AS_TRANS (23456) should be used.
// SetGlobal2ByteAs sets the uint32 value in the ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth object
func (obj *resultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) SetGlobal2ByteAs(value uint32) ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth {

	obj.obj.Global_2ByteAs = &value
	return obj
}

// Bandwidth of the link in bytes per second. ( 1 Kbps is 1000 bytes per second and 1 Mbps is 1000 Kbps per second )
// Bandwidth returns a float32
func (obj *resultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) Bandwidth() float32 {

	return *obj.obj.Bandwidth

}

// Bandwidth of the link in bytes per second. ( 1 Kbps is 1000 bytes per second and 1 Mbps is 1000 Kbps per second )
// Bandwidth returns a float32
func (obj *resultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) HasBandwidth() bool {
	return obj.obj.Bandwidth != nil
}

// Bandwidth of the link in bytes per second. ( 1 Kbps is 1000 bytes per second and 1 Mbps is 1000 Kbps per second )
// SetBandwidth sets the float32 value in the ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth object
func (obj *resultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) SetBandwidth(value float32) ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth {

	obj.obj.Bandwidth = &value
	return obj
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Global_2ByteAs != nil {

		if *obj.obj.Global_2ByteAs > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= ResultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth.Global_2ByteAs <= 65535 but Got %d", *obj.obj.Global_2ByteAs))
		}

	}

}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeLinkBandwidth) setDefault() {

}

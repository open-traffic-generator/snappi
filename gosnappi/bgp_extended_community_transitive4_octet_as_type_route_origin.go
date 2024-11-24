package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin *****
type bgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin struct {
	validation
	obj          *otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	marshaller   marshalBgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	unMarshaller unMarshalBgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
}

func NewBgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin() BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin {
	obj := bgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin{obj: &otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) msg() *otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin {
	return obj.obj
}

func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) setMsg(msg *otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin struct {
	obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
}

type marshalBgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin interface {
	// ToProto marshals BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin to protobuf object *otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	ToProto() (*otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin, error)
	// ToPbText marshals BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin struct {
	obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
}

type unMarshalBgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin interface {
	// FromProto unmarshals BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin from protobuf object *otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	FromProto(msg *otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) (BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin, error)
	// FromPbText unmarshals BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin from JSON text
	FromJson(value string) error
}

func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) Marshal() marshalBgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) Unmarshal() unMarshalBgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) ToProto() (*otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) FromProto(msg *otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) (BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) ToPbText() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) FromPbText(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) ToYaml() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) FromYaml(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) ToJsonRaw() (string, error) {
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

func (m *marshalbgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) ToJson() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) FromJson(value string) error {
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

func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) Clone() (BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin()
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

// BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin is the Route Origin Community identifies one or more routers that inject a set of routes (that carry this Community) into BGP. It is sent with sub-type as 0x03.
type BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin interface {
	Validation
	// msg marshals BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin to protobuf object *otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	// and doesn't set defaults
	msg() *otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	// setMsg unmarshals BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin from protobuf object *otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	// and doesn't set defaults
	setMsg(*otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	// provides marshal interface
	Marshal() marshalBgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	// provides unmarshal interface
	Unmarshal() unMarshalBgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	// validate validates BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Global4ByteAs returns uint32, set in BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin.
	Global4ByteAs() uint32
	// SetGlobal4ByteAs assigns uint32 provided by user to BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	SetGlobal4ByteAs(value uint32) BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	// HasGlobal4ByteAs checks if Global4ByteAs has been set in BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	HasGlobal4ByteAs() bool
	// Local2ByteAdmin returns uint32, set in BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin.
	Local2ByteAdmin() uint32
	// SetLocal2ByteAdmin assigns uint32 provided by user to BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	SetLocal2ByteAdmin(value uint32) BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	// HasLocal2ByteAdmin checks if Local2ByteAdmin has been set in BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin
	HasLocal2ByteAdmin() bool
}

// The four octet IANA assigned AS value assigned to the Autonomous System.
// Global4ByteAs returns a uint32
func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) Global4ByteAs() uint32 {

	return *obj.obj.Global_4ByteAs

}

// The four octet IANA assigned AS value assigned to the Autonomous System.
// Global4ByteAs returns a uint32
func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) HasGlobal4ByteAs() bool {
	return obj.obj.Global_4ByteAs != nil
}

// The four octet IANA assigned AS value assigned to the Autonomous System.
// SetGlobal4ByteAs sets the uint32 value in the BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin object
func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) SetGlobal4ByteAs(value uint32) BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin {

	obj.obj.Global_4ByteAs = &value
	return obj
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local2ByteAdmin returns a uint32
func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) Local2ByteAdmin() uint32 {

	return *obj.obj.Local_2ByteAdmin

}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local2ByteAdmin returns a uint32
func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) HasLocal2ByteAdmin() bool {
	return obj.obj.Local_2ByteAdmin != nil
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// SetLocal2ByteAdmin sets the uint32 value in the BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin object
func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) SetLocal2ByteAdmin(value uint32) BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin {

	obj.obj.Local_2ByteAdmin = &value
	return obj
}

func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Local_2ByteAdmin != nil {

		if *obj.obj.Local_2ByteAdmin > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin.Local_2ByteAdmin <= 65535 but Got %d", *obj.obj.Local_2ByteAdmin))
		}

	}

}

func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteOrigin) setDefault() {
	if obj.obj.Global_4ByteAs == nil {
		obj.SetGlobal4ByteAs(100)
	}
	if obj.obj.Local_2ByteAdmin == nil {
		obj.SetLocal2ByteAdmin(1)
	}

}

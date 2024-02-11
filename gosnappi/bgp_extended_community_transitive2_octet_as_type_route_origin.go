package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin *****
type bgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin struct {
	validation
	obj          *otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	marshaller   marshalBgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	unMarshaller unMarshalBgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
}

func NewBgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin() BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin {
	obj := bgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin{obj: &otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) msg() *otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin {
	return obj.obj
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) setMsg(msg *otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin struct {
	obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
}

type marshalBgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin interface {
	// ToProto marshals BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin to protobuf object *otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	ToProto() (*otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin, error)
	// ToPbText marshals BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin to JSON text
	ToJson() (string, error)
}

type unMarshalbgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin struct {
	obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
}

type unMarshalBgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin interface {
	// FromProto unmarshals BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin from protobuf object *otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	FromProto(msg *otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) (BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin, error)
	// FromPbText unmarshals BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin from JSON text
	FromJson(value string) error
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) Marshal() marshalBgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) Unmarshal() unMarshalBgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) ToProto() (*otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) FromProto(msg *otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) (BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) ToPbText() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) FromPbText(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) ToYaml() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) FromYaml(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) ToJson() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) FromJson(value string) error {
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

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) Clone() (BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin()
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

// BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin is the Route Origin Community identifies one or more routers that inject a set of routes (that carry this Community) into BGP. It is sent with sub-type as 0x03 .
type BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin interface {
	Validation
	// msg marshals BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin to protobuf object *otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	// and doesn't set defaults
	msg() *otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	// setMsg unmarshals BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin from protobuf object *otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	// and doesn't set defaults
	setMsg(*otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	// provides marshal interface
	Marshal() marshalBgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	// provides unmarshal interface
	Unmarshal() unMarshalBgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	// validate validates BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Global2ByteAs returns uint32, set in BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin.
	Global2ByteAs() uint32
	// SetGlobal2ByteAs assigns uint32 provided by user to BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	SetGlobal2ByteAs(value uint32) BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	// HasGlobal2ByteAs checks if Global2ByteAs has been set in BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	HasGlobal2ByteAs() bool
	// Local4ByteAdmin returns uint32, set in BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin.
	Local4ByteAdmin() uint32
	// SetLocal4ByteAdmin assigns uint32 provided by user to BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	SetLocal4ByteAdmin(value uint32) BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	// HasLocal4ByteAdmin checks if Local4ByteAdmin has been set in BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	HasLocal4ByteAdmin() bool
}

// The two octet IANA assigned AS value assigned to the Autonomous System.
// Global2ByteAs returns a uint32
func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) Global2ByteAs() uint32 {

	return *obj.obj.Global_2ByteAs

}

// The two octet IANA assigned AS value assigned to the Autonomous System.
// Global2ByteAs returns a uint32
func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) HasGlobal2ByteAs() bool {
	return obj.obj.Global_2ByteAs != nil
}

// The two octet IANA assigned AS value assigned to the Autonomous System.
// SetGlobal2ByteAs sets the uint32 value in the BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin object
func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) SetGlobal2ByteAs(value uint32) BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin {

	obj.obj.Global_2ByteAs = &value
	return obj
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local4ByteAdmin returns a uint32
func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) Local4ByteAdmin() uint32 {

	return *obj.obj.Local_4ByteAdmin

}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local4ByteAdmin returns a uint32
func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) HasLocal4ByteAdmin() bool {
	return obj.obj.Local_4ByteAdmin != nil
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// SetLocal4ByteAdmin sets the uint32 value in the BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin object
func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) SetLocal4ByteAdmin(value uint32) BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin {

	obj.obj.Local_4ByteAdmin = &value
	return obj
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Global_2ByteAs != nil {

		if *obj.obj.Global_2ByteAs > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin.Global_2ByteAs <= 65535 but Got %d", *obj.obj.Global_2ByteAs))
		}

	}

}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteOrigin) setDefault() {
	if obj.obj.Global_2ByteAs == nil {
		obj.SetGlobal2ByteAs(100)
	}
	if obj.obj.Local_4ByteAdmin == nil {
		obj.SetLocal4ByteAdmin(1)
	}

}

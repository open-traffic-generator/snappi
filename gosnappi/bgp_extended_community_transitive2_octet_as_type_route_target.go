package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget *****
type bgpExtendedCommunityTransitive2OctetAsTypeRouteTarget struct {
	validation
	obj          *otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
	marshaller   marshalBgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
	unMarshaller unMarshalBgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
}

func NewBgpExtendedCommunityTransitive2OctetAsTypeRouteTarget() BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget {
	obj := bgpExtendedCommunityTransitive2OctetAsTypeRouteTarget{obj: &otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) msg() *otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget {
	return obj.obj
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) setMsg(msg *otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpExtendedCommunityTransitive2OctetAsTypeRouteTarget struct {
	obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
}

type marshalBgpExtendedCommunityTransitive2OctetAsTypeRouteTarget interface {
	// ToProto marshals BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget to protobuf object *otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
	ToProto() (*otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget, error)
	// ToPbText marshals BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget to JSON text
	ToJson() (string, error)
}

type unMarshalbgpExtendedCommunityTransitive2OctetAsTypeRouteTarget struct {
	obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
}

type unMarshalBgpExtendedCommunityTransitive2OctetAsTypeRouteTarget interface {
	// FromProto unmarshals BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget from protobuf object *otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
	FromProto(msg *otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) (BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget, error)
	// FromPbText unmarshals BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget from JSON text
	FromJson(value string) error
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) Marshal() marshalBgpExtendedCommunityTransitive2OctetAsTypeRouteTarget {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpExtendedCommunityTransitive2OctetAsTypeRouteTarget{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) Unmarshal() unMarshalBgpExtendedCommunityTransitive2OctetAsTypeRouteTarget {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpExtendedCommunityTransitive2OctetAsTypeRouteTarget{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) ToProto() (*otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) FromProto(msg *otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) (BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) ToPbText() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) FromPbText(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) ToYaml() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) FromYaml(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) ToJson() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) FromJson(value string) error {
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

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) Clone() (BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpExtendedCommunityTransitive2OctetAsTypeRouteTarget()
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

// BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget is the Route Target Community identifies one or more routers that may receive a set of routes (that carry this Community) carried by BGP.  It is sent with sub-type as 0x02.
type BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget interface {
	Validation
	// msg marshals BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget to protobuf object *otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
	// and doesn't set defaults
	msg() *otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
	// setMsg unmarshals BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget from protobuf object *otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
	// and doesn't set defaults
	setMsg(*otg.BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
	// provides marshal interface
	Marshal() marshalBgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
	// provides unmarshal interface
	Unmarshal() unMarshalBgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
	// validate validates BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Global2ByteAs returns uint32, set in BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget.
	Global2ByteAs() uint32
	// SetGlobal2ByteAs assigns uint32 provided by user to BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
	SetGlobal2ByteAs(value uint32) BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
	// HasGlobal2ByteAs checks if Global2ByteAs has been set in BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
	HasGlobal2ByteAs() bool
	// Local4ByteAdmin returns uint32, set in BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget.
	Local4ByteAdmin() uint32
	// SetLocal4ByteAdmin assigns uint32 provided by user to BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
	SetLocal4ByteAdmin(value uint32) BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
	// HasLocal4ByteAdmin checks if Local4ByteAdmin has been set in BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget
	HasLocal4ByteAdmin() bool
}

// The two octet IANA assigned AS value assigned to the Autonomous System.
// Global2ByteAs returns a uint32
func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) Global2ByteAs() uint32 {

	return *obj.obj.Global_2ByteAs

}

// The two octet IANA assigned AS value assigned to the Autonomous System.
// Global2ByteAs returns a uint32
func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) HasGlobal2ByteAs() bool {
	return obj.obj.Global_2ByteAs != nil
}

// The two octet IANA assigned AS value assigned to the Autonomous System.
// SetGlobal2ByteAs sets the uint32 value in the BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget object
func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) SetGlobal2ByteAs(value uint32) BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget {

	obj.obj.Global_2ByteAs = &value
	return obj
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local4ByteAdmin returns a uint32
func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) Local4ByteAdmin() uint32 {

	return *obj.obj.Local_4ByteAdmin

}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local4ByteAdmin returns a uint32
func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) HasLocal4ByteAdmin() bool {
	return obj.obj.Local_4ByteAdmin != nil
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// SetLocal4ByteAdmin sets the uint32 value in the BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget object
func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) SetLocal4ByteAdmin(value uint32) BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget {

	obj.obj.Local_4ByteAdmin = &value
	return obj
}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Global_2ByteAs != nil {

		if *obj.obj.Global_2ByteAs > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpExtendedCommunityTransitive2OctetAsTypeRouteTarget.Global_2ByteAs <= 65535 but Got %d", *obj.obj.Global_2ByteAs))
		}

	}

}

func (obj *bgpExtendedCommunityTransitive2OctetAsTypeRouteTarget) setDefault() {
	if obj.obj.Global_2ByteAs == nil {
		obj.SetGlobal2ByteAs(100)
	}
	if obj.obj.Local_4ByteAdmin == nil {
		obj.SetLocal4ByteAdmin(1)
	}

}

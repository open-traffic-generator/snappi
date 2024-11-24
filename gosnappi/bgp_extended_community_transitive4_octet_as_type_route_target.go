package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget *****
type bgpExtendedCommunityTransitive4OctetAsTypeRouteTarget struct {
	validation
	obj          *otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
	marshaller   marshalBgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
	unMarshaller unMarshalBgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
}

func NewBgpExtendedCommunityTransitive4OctetAsTypeRouteTarget() BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget {
	obj := bgpExtendedCommunityTransitive4OctetAsTypeRouteTarget{obj: &otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) msg() *otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget {
	return obj.obj
}

func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) setMsg(msg *otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpExtendedCommunityTransitive4OctetAsTypeRouteTarget struct {
	obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
}

type marshalBgpExtendedCommunityTransitive4OctetAsTypeRouteTarget interface {
	// ToProto marshals BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget to protobuf object *otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
	ToProto() (*otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget, error)
	// ToPbText marshals BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpExtendedCommunityTransitive4OctetAsTypeRouteTarget struct {
	obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
}

type unMarshalBgpExtendedCommunityTransitive4OctetAsTypeRouteTarget interface {
	// FromProto unmarshals BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget from protobuf object *otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
	FromProto(msg *otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) (BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget, error)
	// FromPbText unmarshals BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget from JSON text
	FromJson(value string) error
}

func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) Marshal() marshalBgpExtendedCommunityTransitive4OctetAsTypeRouteTarget {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpExtendedCommunityTransitive4OctetAsTypeRouteTarget{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) Unmarshal() unMarshalBgpExtendedCommunityTransitive4OctetAsTypeRouteTarget {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpExtendedCommunityTransitive4OctetAsTypeRouteTarget{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) ToProto() (*otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) FromProto(msg *otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) (BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) ToPbText() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) FromPbText(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) ToYaml() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) FromYaml(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) ToJsonRaw() (string, error) {
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

func (m *marshalbgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) ToJson() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) FromJson(value string) error {
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

func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) Clone() (BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpExtendedCommunityTransitive4OctetAsTypeRouteTarget()
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

// BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget is the Route Target Community identifies one or more routers that may receive a set of routes (that carry this Community) carried by BGP.  It is sent with sub-type as 0x02
type BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget interface {
	Validation
	// msg marshals BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget to protobuf object *otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
	// and doesn't set defaults
	msg() *otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
	// setMsg unmarshals BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget from protobuf object *otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
	// and doesn't set defaults
	setMsg(*otg.BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
	// provides marshal interface
	Marshal() marshalBgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
	// provides unmarshal interface
	Unmarshal() unMarshalBgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
	// validate validates BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Global4ByteAs returns uint32, set in BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget.
	Global4ByteAs() uint32
	// SetGlobal4ByteAs assigns uint32 provided by user to BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
	SetGlobal4ByteAs(value uint32) BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
	// HasGlobal4ByteAs checks if Global4ByteAs has been set in BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
	HasGlobal4ByteAs() bool
	// Local2ByteAdmin returns uint32, set in BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget.
	Local2ByteAdmin() uint32
	// SetLocal2ByteAdmin assigns uint32 provided by user to BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
	SetLocal2ByteAdmin(value uint32) BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
	// HasLocal2ByteAdmin checks if Local2ByteAdmin has been set in BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget
	HasLocal2ByteAdmin() bool
}

// The four octet IANA assigned AS value assigned to the Autonomous System.
// Global4ByteAs returns a uint32
func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) Global4ByteAs() uint32 {

	return *obj.obj.Global_4ByteAs

}

// The four octet IANA assigned AS value assigned to the Autonomous System.
// Global4ByteAs returns a uint32
func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) HasGlobal4ByteAs() bool {
	return obj.obj.Global_4ByteAs != nil
}

// The four octet IANA assigned AS value assigned to the Autonomous System.
// SetGlobal4ByteAs sets the uint32 value in the BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget object
func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) SetGlobal4ByteAs(value uint32) BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget {

	obj.obj.Global_4ByteAs = &value
	return obj
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local2ByteAdmin returns a uint32
func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) Local2ByteAdmin() uint32 {

	return *obj.obj.Local_2ByteAdmin

}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local2ByteAdmin returns a uint32
func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) HasLocal2ByteAdmin() bool {
	return obj.obj.Local_2ByteAdmin != nil
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// SetLocal2ByteAdmin sets the uint32 value in the BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget object
func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) SetLocal2ByteAdmin(value uint32) BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget {

	obj.obj.Local_2ByteAdmin = &value
	return obj
}

func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Local_2ByteAdmin != nil {

		if *obj.obj.Local_2ByteAdmin > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpExtendedCommunityTransitive4OctetAsTypeRouteTarget.Local_2ByteAdmin <= 65535 but Got %d", *obj.obj.Local_2ByteAdmin))
		}

	}

}

func (obj *bgpExtendedCommunityTransitive4OctetAsTypeRouteTarget) setDefault() {
	if obj.obj.Global_4ByteAs == nil {
		obj.SetGlobal4ByteAs(100)
	}
	if obj.obj.Local_2ByteAdmin == nil {
		obj.SetLocal2ByteAdmin(1)
	}

}

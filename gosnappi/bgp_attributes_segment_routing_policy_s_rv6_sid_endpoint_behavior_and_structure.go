package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure *****
type bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure struct {
	validation
	obj          *otg.BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	marshaller   marshalBgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	unMarshaller unMarshalBgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
}

func NewBgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure() BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure {
	obj := bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure{obj: &otg.BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) msg() *otg.BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure {
	return obj.obj
}

func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) setMsg(msg *otg.BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure struct {
	obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
}

type marshalBgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure interface {
	// ToProto marshals BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure to protobuf object *otg.BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	ToProto() (*otg.BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure, error)
	// ToPbText marshals BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure struct {
	obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
}

type unMarshalBgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure interface {
	// FromProto unmarshals BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure from protobuf object *otg.BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	FromProto(msg *otg.BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) (BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure, error)
	// FromPbText unmarshals BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) Marshal() marshalBgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) ToProto() (*otg.BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) FromProto(msg *otg.BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) (BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) FromJson(value string) error {
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

func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) Clone() (BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure()
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

// BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure is configuration for optional SRv6 Endpoint Behavior and SID Structure. Summation of lengths for Locator Block, Locator Node,  Function, and Argument MUST be less than or equal to 128. - This is specified in draft-ietf-idr-sr-policy-safi-02 Section 2.4.4.2.4
type BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure interface {
	Validation
	// msg marshals BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure to protobuf object *otg.BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	// and doesn't set defaults
	msg() *otg.BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	// setMsg unmarshals BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure from protobuf object *otg.BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	// provides marshal interface
	Marshal() marshalBgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	// validate validates BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EndpointBehaviour returns string, set in BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure.
	EndpointBehaviour() string
	// SetEndpointBehaviour assigns string provided by user to BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	SetEndpointBehaviour(value string) BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	// HasEndpointBehaviour checks if EndpointBehaviour has been set in BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	HasEndpointBehaviour() bool
	// LbLength returns uint32, set in BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure.
	LbLength() uint32
	// SetLbLength assigns uint32 provided by user to BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	SetLbLength(value uint32) BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	// HasLbLength checks if LbLength has been set in BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	HasLbLength() bool
	// LnLength returns uint32, set in BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure.
	LnLength() uint32
	// SetLnLength assigns uint32 provided by user to BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	SetLnLength(value uint32) BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	// HasLnLength checks if LnLength has been set in BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	HasLnLength() bool
	// FuncLength returns uint32, set in BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure.
	FuncLength() uint32
	// SetFuncLength assigns uint32 provided by user to BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	SetFuncLength(value uint32) BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	// HasFuncLength checks if FuncLength has been set in BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	HasFuncLength() bool
	// ArgLength returns uint32, set in BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure.
	ArgLength() uint32
	// SetArgLength assigns uint32 provided by user to BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	SetArgLength(value uint32) BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	// HasArgLength checks if ArgLength has been set in BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	HasArgLength() bool
}

// This is a 2-octet field that is used to specify the SRv6 Endpoint Behavior code point for the SRv6 SID as defined
// in section 9.2 of [RFC8986]. When set with the value 0xFFFF (i.e., Opaque), the choice of SRv6 Endpoint Behavior is
// left to the headend. Well known 16-bit values for this field are available at
// https://www.iana.org/assignments/segment-routing/segment-routing.xhtml .
// EndpointBehaviour returns a string
func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) EndpointBehaviour() string {

	return *obj.obj.EndpointBehaviour

}

// This is a 2-octet field that is used to specify the SRv6 Endpoint Behavior code point for the SRv6 SID as defined
// in section 9.2 of [RFC8986]. When set with the value 0xFFFF (i.e., Opaque), the choice of SRv6 Endpoint Behavior is
// left to the headend. Well known 16-bit values for this field are available at
// https://www.iana.org/assignments/segment-routing/segment-routing.xhtml .
// EndpointBehaviour returns a string
func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) HasEndpointBehaviour() bool {
	return obj.obj.EndpointBehaviour != nil
}

// This is a 2-octet field that is used to specify the SRv6 Endpoint Behavior code point for the SRv6 SID as defined
// in section 9.2 of [RFC8986]. When set with the value 0xFFFF (i.e., Opaque), the choice of SRv6 Endpoint Behavior is
// left to the headend. Well known 16-bit values for this field are available at
// https://www.iana.org/assignments/segment-routing/segment-routing.xhtml .
// SetEndpointBehaviour sets the string value in the BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure object
func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) SetEndpointBehaviour(value string) BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure {

	obj.obj.EndpointBehaviour = &value
	return obj
}

// SRv6 SID Locator Block length in bits.
// LbLength returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) LbLength() uint32 {

	return *obj.obj.LbLength

}

// SRv6 SID Locator Block length in bits.
// LbLength returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) HasLbLength() bool {
	return obj.obj.LbLength != nil
}

// SRv6 SID Locator Block length in bits.
// SetLbLength sets the uint32 value in the BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure object
func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) SetLbLength(value uint32) BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure {

	obj.obj.LbLength = &value
	return obj
}

// SRv6 SID Locator Node length in bits.
// LnLength returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) LnLength() uint32 {

	return *obj.obj.LnLength

}

// SRv6 SID Locator Node length in bits.
// LnLength returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) HasLnLength() bool {
	return obj.obj.LnLength != nil
}

// SRv6 SID Locator Node length in bits.
// SetLnLength sets the uint32 value in the BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure object
func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) SetLnLength(value uint32) BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure {

	obj.obj.LnLength = &value
	return obj
}

// SRv6 SID Function length in bits.
// FuncLength returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) FuncLength() uint32 {

	return *obj.obj.FuncLength

}

// SRv6 SID Function length in bits.
// FuncLength returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) HasFuncLength() bool {
	return obj.obj.FuncLength != nil
}

// SRv6 SID Function length in bits.
// SetFuncLength sets the uint32 value in the BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure object
func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) SetFuncLength(value uint32) BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure {

	obj.obj.FuncLength = &value
	return obj
}

// SRv6 SID Arguments length in bits.
// ArgLength returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) ArgLength() uint32 {

	return *obj.obj.ArgLength

}

// SRv6 SID Arguments length in bits.
// ArgLength returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) HasArgLength() bool {
	return obj.obj.ArgLength != nil
}

// SRv6 SID Arguments length in bits.
// SetArgLength sets the uint32 value in the BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure object
func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) SetArgLength(value uint32) BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure {

	obj.obj.ArgLength = &value
	return obj
}

func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.EndpointBehaviour != nil {

		if len(*obj.obj.EndpointBehaviour) > 4 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"None <= length of BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure.EndpointBehaviour <= 4 but Got %d",
					len(*obj.obj.EndpointBehaviour)))
		}

	}

	if obj.obj.LbLength != nil {

		if *obj.obj.LbLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure.LbLength <= 128 but Got %d", *obj.obj.LbLength))
		}

	}

	if obj.obj.LnLength != nil {

		if *obj.obj.LnLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure.LnLength <= 128 but Got %d", *obj.obj.LnLength))
		}

	}

	if obj.obj.FuncLength != nil {

		if *obj.obj.FuncLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure.FuncLength <= 128 but Got %d", *obj.obj.FuncLength))
		}

	}

	if obj.obj.ArgLength != nil {

		if *obj.obj.ArgLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure.ArgLength <= 128 but Got %d", *obj.obj.ArgLength))
		}

	}

}

func (obj *bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) setDefault() {
	if obj.obj.EndpointBehaviour == nil {
		obj.SetEndpointBehaviour("ffff")
	}
	if obj.obj.LbLength == nil {
		obj.SetLbLength(0)
	}
	if obj.obj.LnLength == nil {
		obj.SetLnLength(0)
	}
	if obj.obj.FuncLength == nil {
		obj.SetFuncLength(0)
	}
	if obj.obj.ArgLength == nil {
		obj.SetArgLength(0)
	}

}

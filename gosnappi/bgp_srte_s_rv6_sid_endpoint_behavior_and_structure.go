package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrteSRv6SIDEndpointBehaviorAndStructure *****
type bgpSrteSRv6SIDEndpointBehaviorAndStructure struct {
	validation
	obj          *otg.BgpSrteSRv6SIDEndpointBehaviorAndStructure
	marshaller   marshalBgpSrteSRv6SIDEndpointBehaviorAndStructure
	unMarshaller unMarshalBgpSrteSRv6SIDEndpointBehaviorAndStructure
}

func NewBgpSrteSRv6SIDEndpointBehaviorAndStructure() BgpSrteSRv6SIDEndpointBehaviorAndStructure {
	obj := bgpSrteSRv6SIDEndpointBehaviorAndStructure{obj: &otg.BgpSrteSRv6SIDEndpointBehaviorAndStructure{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) msg() *otg.BgpSrteSRv6SIDEndpointBehaviorAndStructure {
	return obj.obj
}

func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) setMsg(msg *otg.BgpSrteSRv6SIDEndpointBehaviorAndStructure) BgpSrteSRv6SIDEndpointBehaviorAndStructure {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteSRv6SIDEndpointBehaviorAndStructure struct {
	obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure
}

type marshalBgpSrteSRv6SIDEndpointBehaviorAndStructure interface {
	// ToProto marshals BgpSrteSRv6SIDEndpointBehaviorAndStructure to protobuf object *otg.BgpSrteSRv6SIDEndpointBehaviorAndStructure
	ToProto() (*otg.BgpSrteSRv6SIDEndpointBehaviorAndStructure, error)
	// ToPbText marshals BgpSrteSRv6SIDEndpointBehaviorAndStructure to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteSRv6SIDEndpointBehaviorAndStructure to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteSRv6SIDEndpointBehaviorAndStructure to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpSrteSRv6SIDEndpointBehaviorAndStructure to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpSrteSRv6SIDEndpointBehaviorAndStructure struct {
	obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure
}

type unMarshalBgpSrteSRv6SIDEndpointBehaviorAndStructure interface {
	// FromProto unmarshals BgpSrteSRv6SIDEndpointBehaviorAndStructure from protobuf object *otg.BgpSrteSRv6SIDEndpointBehaviorAndStructure
	FromProto(msg *otg.BgpSrteSRv6SIDEndpointBehaviorAndStructure) (BgpSrteSRv6SIDEndpointBehaviorAndStructure, error)
	// FromPbText unmarshals BgpSrteSRv6SIDEndpointBehaviorAndStructure from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteSRv6SIDEndpointBehaviorAndStructure from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteSRv6SIDEndpointBehaviorAndStructure from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) Marshal() marshalBgpSrteSRv6SIDEndpointBehaviorAndStructure {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteSRv6SIDEndpointBehaviorAndStructure{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) Unmarshal() unMarshalBgpSrteSRv6SIDEndpointBehaviorAndStructure {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteSRv6SIDEndpointBehaviorAndStructure{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteSRv6SIDEndpointBehaviorAndStructure) ToProto() (*otg.BgpSrteSRv6SIDEndpointBehaviorAndStructure, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteSRv6SIDEndpointBehaviorAndStructure) FromProto(msg *otg.BgpSrteSRv6SIDEndpointBehaviorAndStructure) (BgpSrteSRv6SIDEndpointBehaviorAndStructure, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteSRv6SIDEndpointBehaviorAndStructure) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteSRv6SIDEndpointBehaviorAndStructure) FromPbText(value string) error {
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

func (m *marshalbgpSrteSRv6SIDEndpointBehaviorAndStructure) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteSRv6SIDEndpointBehaviorAndStructure) FromYaml(value string) error {
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

func (m *marshalbgpSrteSRv6SIDEndpointBehaviorAndStructure) ToJsonRaw() (string, error) {
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

func (m *marshalbgpSrteSRv6SIDEndpointBehaviorAndStructure) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteSRv6SIDEndpointBehaviorAndStructure) FromJson(value string) error {
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

func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) Clone() (BgpSrteSRv6SIDEndpointBehaviorAndStructure, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteSRv6SIDEndpointBehaviorAndStructure()
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

// BgpSrteSRv6SIDEndpointBehaviorAndStructure is configuration for SRv6 Endpoint Behavior and SID Structure.  Its optional. Summation of lengths for Locator Block, Locator Node,  Function, and Argument MUST be less than or equal to 128.
type BgpSrteSRv6SIDEndpointBehaviorAndStructure interface {
	Validation
	// msg marshals BgpSrteSRv6SIDEndpointBehaviorAndStructure to protobuf object *otg.BgpSrteSRv6SIDEndpointBehaviorAndStructure
	// and doesn't set defaults
	msg() *otg.BgpSrteSRv6SIDEndpointBehaviorAndStructure
	// setMsg unmarshals BgpSrteSRv6SIDEndpointBehaviorAndStructure from protobuf object *otg.BgpSrteSRv6SIDEndpointBehaviorAndStructure
	// and doesn't set defaults
	setMsg(*otg.BgpSrteSRv6SIDEndpointBehaviorAndStructure) BgpSrteSRv6SIDEndpointBehaviorAndStructure
	// provides marshal interface
	Marshal() marshalBgpSrteSRv6SIDEndpointBehaviorAndStructure
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteSRv6SIDEndpointBehaviorAndStructure
	// validate validates BgpSrteSRv6SIDEndpointBehaviorAndStructure
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteSRv6SIDEndpointBehaviorAndStructure, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LbLength returns uint32, set in BgpSrteSRv6SIDEndpointBehaviorAndStructure.
	LbLength() uint32
	// SetLbLength assigns uint32 provided by user to BgpSrteSRv6SIDEndpointBehaviorAndStructure
	SetLbLength(value uint32) BgpSrteSRv6SIDEndpointBehaviorAndStructure
	// HasLbLength checks if LbLength has been set in BgpSrteSRv6SIDEndpointBehaviorAndStructure
	HasLbLength() bool
	// LnLength returns uint32, set in BgpSrteSRv6SIDEndpointBehaviorAndStructure.
	LnLength() uint32
	// SetLnLength assigns uint32 provided by user to BgpSrteSRv6SIDEndpointBehaviorAndStructure
	SetLnLength(value uint32) BgpSrteSRv6SIDEndpointBehaviorAndStructure
	// HasLnLength checks if LnLength has been set in BgpSrteSRv6SIDEndpointBehaviorAndStructure
	HasLnLength() bool
	// FuncLength returns uint32, set in BgpSrteSRv6SIDEndpointBehaviorAndStructure.
	FuncLength() uint32
	// SetFuncLength assigns uint32 provided by user to BgpSrteSRv6SIDEndpointBehaviorAndStructure
	SetFuncLength(value uint32) BgpSrteSRv6SIDEndpointBehaviorAndStructure
	// HasFuncLength checks if FuncLength has been set in BgpSrteSRv6SIDEndpointBehaviorAndStructure
	HasFuncLength() bool
	// ArgLength returns uint32, set in BgpSrteSRv6SIDEndpointBehaviorAndStructure.
	ArgLength() uint32
	// SetArgLength assigns uint32 provided by user to BgpSrteSRv6SIDEndpointBehaviorAndStructure
	SetArgLength(value uint32) BgpSrteSRv6SIDEndpointBehaviorAndStructure
	// HasArgLength checks if ArgLength has been set in BgpSrteSRv6SIDEndpointBehaviorAndStructure
	HasArgLength() bool
}

// SRv6 SID Locator Block length in bits.
// LbLength returns a uint32
func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) LbLength() uint32 {

	return *obj.obj.LbLength

}

// SRv6 SID Locator Block length in bits.
// LbLength returns a uint32
func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) HasLbLength() bool {
	return obj.obj.LbLength != nil
}

// SRv6 SID Locator Block length in bits.
// SetLbLength sets the uint32 value in the BgpSrteSRv6SIDEndpointBehaviorAndStructure object
func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) SetLbLength(value uint32) BgpSrteSRv6SIDEndpointBehaviorAndStructure {

	obj.obj.LbLength = &value
	return obj
}

// SRv6 SID Locator Node length in bits.
// LnLength returns a uint32
func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) LnLength() uint32 {

	return *obj.obj.LnLength

}

// SRv6 SID Locator Node length in bits.
// LnLength returns a uint32
func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) HasLnLength() bool {
	return obj.obj.LnLength != nil
}

// SRv6 SID Locator Node length in bits.
// SetLnLength sets the uint32 value in the BgpSrteSRv6SIDEndpointBehaviorAndStructure object
func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) SetLnLength(value uint32) BgpSrteSRv6SIDEndpointBehaviorAndStructure {

	obj.obj.LnLength = &value
	return obj
}

// SRv6 SID Function length in bits.
// FuncLength returns a uint32
func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) FuncLength() uint32 {

	return *obj.obj.FuncLength

}

// SRv6 SID Function length in bits.
// FuncLength returns a uint32
func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) HasFuncLength() bool {
	return obj.obj.FuncLength != nil
}

// SRv6 SID Function length in bits.
// SetFuncLength sets the uint32 value in the BgpSrteSRv6SIDEndpointBehaviorAndStructure object
func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) SetFuncLength(value uint32) BgpSrteSRv6SIDEndpointBehaviorAndStructure {

	obj.obj.FuncLength = &value
	return obj
}

// SRv6 SID Arguments length in bits.
// ArgLength returns a uint32
func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) ArgLength() uint32 {

	return *obj.obj.ArgLength

}

// SRv6 SID Arguments length in bits.
// ArgLength returns a uint32
func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) HasArgLength() bool {
	return obj.obj.ArgLength != nil
}

// SRv6 SID Arguments length in bits.
// SetArgLength sets the uint32 value in the BgpSrteSRv6SIDEndpointBehaviorAndStructure object
func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) SetArgLength(value uint32) BgpSrteSRv6SIDEndpointBehaviorAndStructure {

	obj.obj.ArgLength = &value
	return obj
}

func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LbLength != nil {

		if *obj.obj.LbLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpSrteSRv6SIDEndpointBehaviorAndStructure.LbLength <= 128 but Got %d", *obj.obj.LbLength))
		}

	}

	if obj.obj.LnLength != nil {

		if *obj.obj.LnLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpSrteSRv6SIDEndpointBehaviorAndStructure.LnLength <= 128 but Got %d", *obj.obj.LnLength))
		}

	}

	if obj.obj.FuncLength != nil {

		if *obj.obj.FuncLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpSrteSRv6SIDEndpointBehaviorAndStructure.FuncLength <= 128 but Got %d", *obj.obj.FuncLength))
		}

	}

	if obj.obj.ArgLength != nil {

		if *obj.obj.ArgLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpSrteSRv6SIDEndpointBehaviorAndStructure.ArgLength <= 128 but Got %d", *obj.obj.ArgLength))
		}

	}

}

func (obj *bgpSrteSRv6SIDEndpointBehaviorAndStructure) setDefault() {
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

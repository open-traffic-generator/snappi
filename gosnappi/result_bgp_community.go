package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ResultBgpCommunity *****
type resultBgpCommunity struct {
	validation
	obj          *otg.ResultBgpCommunity
	marshaller   marshalResultBgpCommunity
	unMarshaller unMarshalResultBgpCommunity
}

func NewResultBgpCommunity() ResultBgpCommunity {
	obj := resultBgpCommunity{obj: &otg.ResultBgpCommunity{}}
	obj.setDefault()
	return &obj
}

func (obj *resultBgpCommunity) msg() *otg.ResultBgpCommunity {
	return obj.obj
}

func (obj *resultBgpCommunity) setMsg(msg *otg.ResultBgpCommunity) ResultBgpCommunity {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalresultBgpCommunity struct {
	obj *resultBgpCommunity
}

type marshalResultBgpCommunity interface {
	// ToProto marshals ResultBgpCommunity to protobuf object *otg.ResultBgpCommunity
	ToProto() (*otg.ResultBgpCommunity, error)
	// ToPbText marshals ResultBgpCommunity to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ResultBgpCommunity to YAML text
	ToYaml() (string, error)
	// ToJson marshals ResultBgpCommunity to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals ResultBgpCommunity to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalresultBgpCommunity struct {
	obj *resultBgpCommunity
}

type unMarshalResultBgpCommunity interface {
	// FromProto unmarshals ResultBgpCommunity from protobuf object *otg.ResultBgpCommunity
	FromProto(msg *otg.ResultBgpCommunity) (ResultBgpCommunity, error)
	// FromPbText unmarshals ResultBgpCommunity from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ResultBgpCommunity from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ResultBgpCommunity from JSON text
	FromJson(value string) error
}

func (obj *resultBgpCommunity) Marshal() marshalResultBgpCommunity {
	if obj.marshaller == nil {
		obj.marshaller = &marshalresultBgpCommunity{obj: obj}
	}
	return obj.marshaller
}

func (obj *resultBgpCommunity) Unmarshal() unMarshalResultBgpCommunity {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalresultBgpCommunity{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalresultBgpCommunity) ToProto() (*otg.ResultBgpCommunity, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalresultBgpCommunity) FromProto(msg *otg.ResultBgpCommunity) (ResultBgpCommunity, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalresultBgpCommunity) ToPbText() (string, error) {
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

func (m *unMarshalresultBgpCommunity) FromPbText(value string) error {
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

func (m *marshalresultBgpCommunity) ToYaml() (string, error) {
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

func (m *unMarshalresultBgpCommunity) FromYaml(value string) error {
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

func (m *marshalresultBgpCommunity) ToJsonRaw() (string, error) {
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

func (m *marshalresultBgpCommunity) ToJson() (string, error) {
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

func (m *unMarshalresultBgpCommunity) FromJson(value string) error {
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

func (obj *resultBgpCommunity) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *resultBgpCommunity) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *resultBgpCommunity) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *resultBgpCommunity) Clone() (ResultBgpCommunity, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewResultBgpCommunity()
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

// ResultBgpCommunity is bGP communities provide additional capability for tagging routes and  for modifying BGP routing policy on upstream and downstream routers. BGP community is a 32-bit number which is broken into 16-bit AS number and  a 16-bit custom value.
type ResultBgpCommunity interface {
	Validation
	// msg marshals ResultBgpCommunity to protobuf object *otg.ResultBgpCommunity
	// and doesn't set defaults
	msg() *otg.ResultBgpCommunity
	// setMsg unmarshals ResultBgpCommunity from protobuf object *otg.ResultBgpCommunity
	// and doesn't set defaults
	setMsg(*otg.ResultBgpCommunity) ResultBgpCommunity
	// provides marshal interface
	Marshal() marshalResultBgpCommunity
	// provides unmarshal interface
	Unmarshal() unMarshalResultBgpCommunity
	// validate validates ResultBgpCommunity
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ResultBgpCommunity, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns ResultBgpCommunityTypeEnum, set in ResultBgpCommunity
	Type() ResultBgpCommunityTypeEnum
	// SetType assigns ResultBgpCommunityTypeEnum provided by user to ResultBgpCommunity
	SetType(value ResultBgpCommunityTypeEnum) ResultBgpCommunity
	// HasType checks if Type has been set in ResultBgpCommunity
	HasType() bool
	// AsNumber returns uint32, set in ResultBgpCommunity.
	AsNumber() uint32
	// SetAsNumber assigns uint32 provided by user to ResultBgpCommunity
	SetAsNumber(value uint32) ResultBgpCommunity
	// HasAsNumber checks if AsNumber has been set in ResultBgpCommunity
	HasAsNumber() bool
	// AsCustom returns uint32, set in ResultBgpCommunity.
	AsCustom() uint32
	// SetAsCustom assigns uint32 provided by user to ResultBgpCommunity
	SetAsCustom(value uint32) ResultBgpCommunity
	// HasAsCustom checks if AsCustom has been set in ResultBgpCommunity
	HasAsCustom() bool
}

type ResultBgpCommunityTypeEnum string

// Enum of Type on ResultBgpCommunity
var ResultBgpCommunityType = struct {
	MANUAL_AS_NUMBER    ResultBgpCommunityTypeEnum
	NO_EXPORT           ResultBgpCommunityTypeEnum
	NO_ADVERTISED       ResultBgpCommunityTypeEnum
	NO_EXPORT_SUBCONFED ResultBgpCommunityTypeEnum
	LLGR_STALE          ResultBgpCommunityTypeEnum
	NO_LLGR             ResultBgpCommunityTypeEnum
}{
	MANUAL_AS_NUMBER:    ResultBgpCommunityTypeEnum("manual_as_number"),
	NO_EXPORT:           ResultBgpCommunityTypeEnum("no_export"),
	NO_ADVERTISED:       ResultBgpCommunityTypeEnum("no_advertised"),
	NO_EXPORT_SUBCONFED: ResultBgpCommunityTypeEnum("no_export_subconfed"),
	LLGR_STALE:          ResultBgpCommunityTypeEnum("llgr_stale"),
	NO_LLGR:             ResultBgpCommunityTypeEnum("no_llgr"),
}

func (obj *resultBgpCommunity) Type() ResultBgpCommunityTypeEnum {
	return ResultBgpCommunityTypeEnum(obj.obj.Type.Enum().String())
}

// The type of community AS number. If community type is manual_as_number then as_number and as_custom will be available.
// Type returns a string
func (obj *resultBgpCommunity) HasType() bool {
	return obj.obj.Type != nil
}

func (obj *resultBgpCommunity) SetType(value ResultBgpCommunityTypeEnum) ResultBgpCommunity {
	intValue, ok := otg.ResultBgpCommunity_Type_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ResultBgpCommunityTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.ResultBgpCommunity_Type_Enum(intValue)
	obj.obj.Type = &enumValue

	return obj
}

// First two octets of 32 bit community AS number.
// AsNumber returns a uint32
func (obj *resultBgpCommunity) AsNumber() uint32 {

	return *obj.obj.AsNumber

}

// First two octets of 32 bit community AS number.
// AsNumber returns a uint32
func (obj *resultBgpCommunity) HasAsNumber() bool {
	return obj.obj.AsNumber != nil
}

// First two octets of 32 bit community AS number.
// SetAsNumber sets the uint32 value in the ResultBgpCommunity object
func (obj *resultBgpCommunity) SetAsNumber(value uint32) ResultBgpCommunity {

	obj.obj.AsNumber = &value
	return obj
}

// Last two octets of the community value.
// AsCustom returns a uint32
func (obj *resultBgpCommunity) AsCustom() uint32 {

	return *obj.obj.AsCustom

}

// Last two octets of the community value.
// AsCustom returns a uint32
func (obj *resultBgpCommunity) HasAsCustom() bool {
	return obj.obj.AsCustom != nil
}

// Last two octets of the community value.
// SetAsCustom sets the uint32 value in the ResultBgpCommunity object
func (obj *resultBgpCommunity) SetAsCustom(value uint32) ResultBgpCommunity {

	obj.obj.AsCustom = &value
	return obj
}

func (obj *resultBgpCommunity) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.AsNumber != nil {

		if *obj.obj.AsNumber > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= ResultBgpCommunity.AsNumber <= 65535 but Got %d", *obj.obj.AsNumber))
		}

	}

	if obj.obj.AsCustom != nil {

		if *obj.obj.AsCustom > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= ResultBgpCommunity.AsCustom <= 65535 but Got %d", *obj.obj.AsCustom))
		}

	}

}

func (obj *resultBgpCommunity) setDefault() {

}

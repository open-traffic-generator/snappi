package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesFourByteAsPathSegment *****
type bgpAttributesFourByteAsPathSegment struct {
	validation
	obj          *otg.BgpAttributesFourByteAsPathSegment
	marshaller   marshalBgpAttributesFourByteAsPathSegment
	unMarshaller unMarshalBgpAttributesFourByteAsPathSegment
}

func NewBgpAttributesFourByteAsPathSegment() BgpAttributesFourByteAsPathSegment {
	obj := bgpAttributesFourByteAsPathSegment{obj: &otg.BgpAttributesFourByteAsPathSegment{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesFourByteAsPathSegment) msg() *otg.BgpAttributesFourByteAsPathSegment {
	return obj.obj
}

func (obj *bgpAttributesFourByteAsPathSegment) setMsg(msg *otg.BgpAttributesFourByteAsPathSegment) BgpAttributesFourByteAsPathSegment {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesFourByteAsPathSegment struct {
	obj *bgpAttributesFourByteAsPathSegment
}

type marshalBgpAttributesFourByteAsPathSegment interface {
	// ToProto marshals BgpAttributesFourByteAsPathSegment to protobuf object *otg.BgpAttributesFourByteAsPathSegment
	ToProto() (*otg.BgpAttributesFourByteAsPathSegment, error)
	// ToPbText marshals BgpAttributesFourByteAsPathSegment to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesFourByteAsPathSegment to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesFourByteAsPathSegment to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributesFourByteAsPathSegment struct {
	obj *bgpAttributesFourByteAsPathSegment
}

type unMarshalBgpAttributesFourByteAsPathSegment interface {
	// FromProto unmarshals BgpAttributesFourByteAsPathSegment from protobuf object *otg.BgpAttributesFourByteAsPathSegment
	FromProto(msg *otg.BgpAttributesFourByteAsPathSegment) (BgpAttributesFourByteAsPathSegment, error)
	// FromPbText unmarshals BgpAttributesFourByteAsPathSegment from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesFourByteAsPathSegment from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesFourByteAsPathSegment from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesFourByteAsPathSegment) Marshal() marshalBgpAttributesFourByteAsPathSegment {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesFourByteAsPathSegment{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesFourByteAsPathSegment) Unmarshal() unMarshalBgpAttributesFourByteAsPathSegment {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesFourByteAsPathSegment{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesFourByteAsPathSegment) ToProto() (*otg.BgpAttributesFourByteAsPathSegment, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesFourByteAsPathSegment) FromProto(msg *otg.BgpAttributesFourByteAsPathSegment) (BgpAttributesFourByteAsPathSegment, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesFourByteAsPathSegment) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesFourByteAsPathSegment) FromPbText(value string) error {
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

func (m *marshalbgpAttributesFourByteAsPathSegment) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesFourByteAsPathSegment) FromYaml(value string) error {
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

func (m *marshalbgpAttributesFourByteAsPathSegment) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesFourByteAsPathSegment) FromJson(value string) error {
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

func (obj *bgpAttributesFourByteAsPathSegment) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesFourByteAsPathSegment) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesFourByteAsPathSegment) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesFourByteAsPathSegment) Clone() (BgpAttributesFourByteAsPathSegment, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesFourByteAsPathSegment()
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

// BgpAttributesFourByteAsPathSegment is configuration for a single BGP AS path segment containing 4 byte AS numbers.
type BgpAttributesFourByteAsPathSegment interface {
	Validation
	// msg marshals BgpAttributesFourByteAsPathSegment to protobuf object *otg.BgpAttributesFourByteAsPathSegment
	// and doesn't set defaults
	msg() *otg.BgpAttributesFourByteAsPathSegment
	// setMsg unmarshals BgpAttributesFourByteAsPathSegment from protobuf object *otg.BgpAttributesFourByteAsPathSegment
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesFourByteAsPathSegment) BgpAttributesFourByteAsPathSegment
	// provides marshal interface
	Marshal() marshalBgpAttributesFourByteAsPathSegment
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesFourByteAsPathSegment
	// validate validates BgpAttributesFourByteAsPathSegment
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesFourByteAsPathSegment, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns BgpAttributesFourByteAsPathSegmentTypeEnum, set in BgpAttributesFourByteAsPathSegment
	Type() BgpAttributesFourByteAsPathSegmentTypeEnum
	// SetType assigns BgpAttributesFourByteAsPathSegmentTypeEnum provided by user to BgpAttributesFourByteAsPathSegment
	SetType(value BgpAttributesFourByteAsPathSegmentTypeEnum) BgpAttributesFourByteAsPathSegment
	// HasType checks if Type has been set in BgpAttributesFourByteAsPathSegment
	HasType() bool
	// AsNumbers returns []uint32, set in BgpAttributesFourByteAsPathSegment.
	AsNumbers() []uint32
	// SetAsNumbers assigns []uint32 provided by user to BgpAttributesFourByteAsPathSegment
	SetAsNumbers(value []uint32) BgpAttributesFourByteAsPathSegment
}

type BgpAttributesFourByteAsPathSegmentTypeEnum string

// Enum of Type on BgpAttributesFourByteAsPathSegment
var BgpAttributesFourByteAsPathSegmentType = struct {
	AS_SEQ        BgpAttributesFourByteAsPathSegmentTypeEnum
	AS_SET        BgpAttributesFourByteAsPathSegmentTypeEnum
	AS_CONFED_SEQ BgpAttributesFourByteAsPathSegmentTypeEnum
	AS_CONFED_SET BgpAttributesFourByteAsPathSegmentTypeEnum
}{
	AS_SEQ:        BgpAttributesFourByteAsPathSegmentTypeEnum("as_seq"),
	AS_SET:        BgpAttributesFourByteAsPathSegmentTypeEnum("as_set"),
	AS_CONFED_SEQ: BgpAttributesFourByteAsPathSegmentTypeEnum("as_confed_seq"),
	AS_CONFED_SET: BgpAttributesFourByteAsPathSegmentTypeEnum("as_confed_set"),
}

func (obj *bgpAttributesFourByteAsPathSegment) Type() BgpAttributesFourByteAsPathSegmentTypeEnum {
	return BgpAttributesFourByteAsPathSegmentTypeEnum(obj.obj.Type.Enum().String())
}

// AS sequence is the most common type of AS_PATH, it contains the  list
// of ASNs starting with the most recent ASN being added read  from left
// to right.
// The other three AS_PATH types are used for Confederations
// - AS_SET is the type of AS_PATH attribute that summarizes routes using
// using the aggregate-address command, allowing AS_PATHs to be  summarized
// in the update as well.
// - AS_CONFED_SEQ gives the list of ASNs in the path starting with the  most
// recent ASN to be added reading left to right
// - AS_CONFED_SET will allow summarization of multiple AS PATHs to be  sent
// in BGP Updates.
// Type returns a string
func (obj *bgpAttributesFourByteAsPathSegment) HasType() bool {
	return obj.obj.Type != nil
}

func (obj *bgpAttributesFourByteAsPathSegment) SetType(value BgpAttributesFourByteAsPathSegmentTypeEnum) BgpAttributesFourByteAsPathSegment {
	intValue, ok := otg.BgpAttributesFourByteAsPathSegment_Type_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpAttributesFourByteAsPathSegmentTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpAttributesFourByteAsPathSegment_Type_Enum(intValue)
	obj.obj.Type = &enumValue

	return obj
}

// The 4 byte AS numbers in this AS path segment.
// AsNumbers returns a []uint32
func (obj *bgpAttributesFourByteAsPathSegment) AsNumbers() []uint32 {
	if obj.obj.AsNumbers == nil {
		obj.obj.AsNumbers = make([]uint32, 0)
	}
	return obj.obj.AsNumbers
}

// The 4 byte AS numbers in this AS path segment.
// SetAsNumbers sets the []uint32 value in the BgpAttributesFourByteAsPathSegment object
func (obj *bgpAttributesFourByteAsPathSegment) SetAsNumbers(value []uint32) BgpAttributesFourByteAsPathSegment {

	if obj.obj.AsNumbers == nil {
		obj.obj.AsNumbers = make([]uint32, 0)
	}
	obj.obj.AsNumbers = value

	return obj
}

func (obj *bgpAttributesFourByteAsPathSegment) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpAttributesFourByteAsPathSegment) setDefault() {
	if obj.obj.Type == nil {
		obj.SetType(BgpAttributesFourByteAsPathSegmentType.AS_SEQ)

	}

}

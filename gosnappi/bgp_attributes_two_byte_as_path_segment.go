package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesTwoByteAsPathSegment *****
type bgpAttributesTwoByteAsPathSegment struct {
	validation
	obj          *otg.BgpAttributesTwoByteAsPathSegment
	marshaller   marshalBgpAttributesTwoByteAsPathSegment
	unMarshaller unMarshalBgpAttributesTwoByteAsPathSegment
}

func NewBgpAttributesTwoByteAsPathSegment() BgpAttributesTwoByteAsPathSegment {
	obj := bgpAttributesTwoByteAsPathSegment{obj: &otg.BgpAttributesTwoByteAsPathSegment{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesTwoByteAsPathSegment) msg() *otg.BgpAttributesTwoByteAsPathSegment {
	return obj.obj
}

func (obj *bgpAttributesTwoByteAsPathSegment) setMsg(msg *otg.BgpAttributesTwoByteAsPathSegment) BgpAttributesTwoByteAsPathSegment {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesTwoByteAsPathSegment struct {
	obj *bgpAttributesTwoByteAsPathSegment
}

type marshalBgpAttributesTwoByteAsPathSegment interface {
	// ToProto marshals BgpAttributesTwoByteAsPathSegment to protobuf object *otg.BgpAttributesTwoByteAsPathSegment
	ToProto() (*otg.BgpAttributesTwoByteAsPathSegment, error)
	// ToPbText marshals BgpAttributesTwoByteAsPathSegment to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesTwoByteAsPathSegment to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesTwoByteAsPathSegment to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributesTwoByteAsPathSegment struct {
	obj *bgpAttributesTwoByteAsPathSegment
}

type unMarshalBgpAttributesTwoByteAsPathSegment interface {
	// FromProto unmarshals BgpAttributesTwoByteAsPathSegment from protobuf object *otg.BgpAttributesTwoByteAsPathSegment
	FromProto(msg *otg.BgpAttributesTwoByteAsPathSegment) (BgpAttributesTwoByteAsPathSegment, error)
	// FromPbText unmarshals BgpAttributesTwoByteAsPathSegment from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesTwoByteAsPathSegment from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesTwoByteAsPathSegment from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesTwoByteAsPathSegment) Marshal() marshalBgpAttributesTwoByteAsPathSegment {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesTwoByteAsPathSegment{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesTwoByteAsPathSegment) Unmarshal() unMarshalBgpAttributesTwoByteAsPathSegment {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesTwoByteAsPathSegment{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesTwoByteAsPathSegment) ToProto() (*otg.BgpAttributesTwoByteAsPathSegment, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesTwoByteAsPathSegment) FromProto(msg *otg.BgpAttributesTwoByteAsPathSegment) (BgpAttributesTwoByteAsPathSegment, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesTwoByteAsPathSegment) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesTwoByteAsPathSegment) FromPbText(value string) error {
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

func (m *marshalbgpAttributesTwoByteAsPathSegment) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesTwoByteAsPathSegment) FromYaml(value string) error {
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

func (m *marshalbgpAttributesTwoByteAsPathSegment) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesTwoByteAsPathSegment) FromJson(value string) error {
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

func (obj *bgpAttributesTwoByteAsPathSegment) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesTwoByteAsPathSegment) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesTwoByteAsPathSegment) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesTwoByteAsPathSegment) Clone() (BgpAttributesTwoByteAsPathSegment, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesTwoByteAsPathSegment()
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

// BgpAttributesTwoByteAsPathSegment is configuration for a single BGP AS path segment containing 2 byte AS numbers.
type BgpAttributesTwoByteAsPathSegment interface {
	Validation
	// msg marshals BgpAttributesTwoByteAsPathSegment to protobuf object *otg.BgpAttributesTwoByteAsPathSegment
	// and doesn't set defaults
	msg() *otg.BgpAttributesTwoByteAsPathSegment
	// setMsg unmarshals BgpAttributesTwoByteAsPathSegment from protobuf object *otg.BgpAttributesTwoByteAsPathSegment
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesTwoByteAsPathSegment) BgpAttributesTwoByteAsPathSegment
	// provides marshal interface
	Marshal() marshalBgpAttributesTwoByteAsPathSegment
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesTwoByteAsPathSegment
	// validate validates BgpAttributesTwoByteAsPathSegment
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesTwoByteAsPathSegment, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns BgpAttributesTwoByteAsPathSegmentTypeEnum, set in BgpAttributesTwoByteAsPathSegment
	Type() BgpAttributesTwoByteAsPathSegmentTypeEnum
	// SetType assigns BgpAttributesTwoByteAsPathSegmentTypeEnum provided by user to BgpAttributesTwoByteAsPathSegment
	SetType(value BgpAttributesTwoByteAsPathSegmentTypeEnum) BgpAttributesTwoByteAsPathSegment
	// HasType checks if Type has been set in BgpAttributesTwoByteAsPathSegment
	HasType() bool
	// AsNumbers returns []uint32, set in BgpAttributesTwoByteAsPathSegment.
	AsNumbers() []uint32
	// SetAsNumbers assigns []uint32 provided by user to BgpAttributesTwoByteAsPathSegment
	SetAsNumbers(value []uint32) BgpAttributesTwoByteAsPathSegment
}

type BgpAttributesTwoByteAsPathSegmentTypeEnum string

// Enum of Type on BgpAttributesTwoByteAsPathSegment
var BgpAttributesTwoByteAsPathSegmentType = struct {
	AS_SEQ        BgpAttributesTwoByteAsPathSegmentTypeEnum
	AS_SET        BgpAttributesTwoByteAsPathSegmentTypeEnum
	AS_CONFED_SEQ BgpAttributesTwoByteAsPathSegmentTypeEnum
	AS_CONFED_SET BgpAttributesTwoByteAsPathSegmentTypeEnum
}{
	AS_SEQ:        BgpAttributesTwoByteAsPathSegmentTypeEnum("as_seq"),
	AS_SET:        BgpAttributesTwoByteAsPathSegmentTypeEnum("as_set"),
	AS_CONFED_SEQ: BgpAttributesTwoByteAsPathSegmentTypeEnum("as_confed_seq"),
	AS_CONFED_SET: BgpAttributesTwoByteAsPathSegmentTypeEnum("as_confed_set"),
}

func (obj *bgpAttributesTwoByteAsPathSegment) Type() BgpAttributesTwoByteAsPathSegmentTypeEnum {
	return BgpAttributesTwoByteAsPathSegmentTypeEnum(obj.obj.Type.Enum().String())
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
func (obj *bgpAttributesTwoByteAsPathSegment) HasType() bool {
	return obj.obj.Type != nil
}

func (obj *bgpAttributesTwoByteAsPathSegment) SetType(value BgpAttributesTwoByteAsPathSegmentTypeEnum) BgpAttributesTwoByteAsPathSegment {
	intValue, ok := otg.BgpAttributesTwoByteAsPathSegment_Type_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpAttributesTwoByteAsPathSegmentTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpAttributesTwoByteAsPathSegment_Type_Enum(intValue)
	obj.obj.Type = &enumValue

	return obj
}

// The 2 byte AS numbers in this AS path segment.
// AsNumbers returns a []uint32
func (obj *bgpAttributesTwoByteAsPathSegment) AsNumbers() []uint32 {
	if obj.obj.AsNumbers == nil {
		obj.obj.AsNumbers = make([]uint32, 0)
	}
	return obj.obj.AsNumbers
}

// The 2 byte AS numbers in this AS path segment.
// SetAsNumbers sets the []uint32 value in the BgpAttributesTwoByteAsPathSegment object
func (obj *bgpAttributesTwoByteAsPathSegment) SetAsNumbers(value []uint32) BgpAttributesTwoByteAsPathSegment {

	if obj.obj.AsNumbers == nil {
		obj.obj.AsNumbers = make([]uint32, 0)
	}
	obj.obj.AsNumbers = value

	return obj
}

func (obj *bgpAttributesTwoByteAsPathSegment) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.AsNumbers != nil {

		for _, item := range obj.obj.AsNumbers {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= BgpAttributesTwoByteAsPathSegment.AsNumbers <= 65535 but Got %d", item))
			}

		}

	}

}

func (obj *bgpAttributesTwoByteAsPathSegment) setDefault() {
	if obj.obj.Type == nil {
		obj.SetType(BgpAttributesTwoByteAsPathSegmentType.AS_SEQ)

	}

}

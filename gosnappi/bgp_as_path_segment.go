package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAsPathSegment *****
type bgpAsPathSegment struct {
	validation
	obj          *otg.BgpAsPathSegment
	marshaller   marshalBgpAsPathSegment
	unMarshaller unMarshalBgpAsPathSegment
}

func NewBgpAsPathSegment() BgpAsPathSegment {
	obj := bgpAsPathSegment{obj: &otg.BgpAsPathSegment{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAsPathSegment) msg() *otg.BgpAsPathSegment {
	return obj.obj
}

func (obj *bgpAsPathSegment) setMsg(msg *otg.BgpAsPathSegment) BgpAsPathSegment {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAsPathSegment struct {
	obj *bgpAsPathSegment
}

type marshalBgpAsPathSegment interface {
	// ToProto marshals BgpAsPathSegment to protobuf object *otg.BgpAsPathSegment
	ToProto() (*otg.BgpAsPathSegment, error)
	// ToPbText marshals BgpAsPathSegment to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAsPathSegment to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAsPathSegment to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAsPathSegment to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAsPathSegment struct {
	obj *bgpAsPathSegment
}

type unMarshalBgpAsPathSegment interface {
	// FromProto unmarshals BgpAsPathSegment from protobuf object *otg.BgpAsPathSegment
	FromProto(msg *otg.BgpAsPathSegment) (BgpAsPathSegment, error)
	// FromPbText unmarshals BgpAsPathSegment from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAsPathSegment from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAsPathSegment from JSON text
	FromJson(value string) error
}

func (obj *bgpAsPathSegment) Marshal() marshalBgpAsPathSegment {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAsPathSegment{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAsPathSegment) Unmarshal() unMarshalBgpAsPathSegment {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAsPathSegment{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAsPathSegment) ToProto() (*otg.BgpAsPathSegment, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAsPathSegment) FromProto(msg *otg.BgpAsPathSegment) (BgpAsPathSegment, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAsPathSegment) ToPbText() (string, error) {
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

func (m *unMarshalbgpAsPathSegment) FromPbText(value string) error {
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

func (m *marshalbgpAsPathSegment) ToYaml() (string, error) {
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

func (m *unMarshalbgpAsPathSegment) FromYaml(value string) error {
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

func (m *marshalbgpAsPathSegment) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAsPathSegment) ToJson() (string, error) {
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

func (m *unMarshalbgpAsPathSegment) FromJson(value string) error {
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

func (obj *bgpAsPathSegment) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAsPathSegment) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAsPathSegment) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAsPathSegment) Clone() (BgpAsPathSegment, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAsPathSegment()
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

// BgpAsPathSegment is configuration for a single BGP AS path segment
type BgpAsPathSegment interface {
	Validation
	// msg marshals BgpAsPathSegment to protobuf object *otg.BgpAsPathSegment
	// and doesn't set defaults
	msg() *otg.BgpAsPathSegment
	// setMsg unmarshals BgpAsPathSegment from protobuf object *otg.BgpAsPathSegment
	// and doesn't set defaults
	setMsg(*otg.BgpAsPathSegment) BgpAsPathSegment
	// provides marshal interface
	Marshal() marshalBgpAsPathSegment
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAsPathSegment
	// validate validates BgpAsPathSegment
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAsPathSegment, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns BgpAsPathSegmentTypeEnum, set in BgpAsPathSegment
	Type() BgpAsPathSegmentTypeEnum
	// SetType assigns BgpAsPathSegmentTypeEnum provided by user to BgpAsPathSegment
	SetType(value BgpAsPathSegmentTypeEnum) BgpAsPathSegment
	// HasType checks if Type has been set in BgpAsPathSegment
	HasType() bool
	// AsNumbers returns []uint32, set in BgpAsPathSegment.
	AsNumbers() []uint32
	// SetAsNumbers assigns []uint32 provided by user to BgpAsPathSegment
	SetAsNumbers(value []uint32) BgpAsPathSegment
}

type BgpAsPathSegmentTypeEnum string

// Enum of Type on BgpAsPathSegment
var BgpAsPathSegmentType = struct {
	AS_SEQ        BgpAsPathSegmentTypeEnum
	AS_SET        BgpAsPathSegmentTypeEnum
	AS_CONFED_SEQ BgpAsPathSegmentTypeEnum
	AS_CONFED_SET BgpAsPathSegmentTypeEnum
}{
	AS_SEQ:        BgpAsPathSegmentTypeEnum("as_seq"),
	AS_SET:        BgpAsPathSegmentTypeEnum("as_set"),
	AS_CONFED_SEQ: BgpAsPathSegmentTypeEnum("as_confed_seq"),
	AS_CONFED_SET: BgpAsPathSegmentTypeEnum("as_confed_set"),
}

func (obj *bgpAsPathSegment) Type() BgpAsPathSegmentTypeEnum {
	return BgpAsPathSegmentTypeEnum(obj.obj.Type.Enum().String())
}

// AS sequence is the most common type of AS_PATH, it contains the  list of ASNs starting with the most recent ASN being added read  from left to right.
// The other three AS_PATH types are used for Confederations - AS_SET is the type of AS_PATH attribute that summarizes routes using using the aggregate-address command, allowing AS_PATHs to be  summarized in the update as well. - AS_CONFED_SEQ gives the list of ASNs in the path starting with the  most recent ASN to be added reading left to right - AS_CONFED_SET will allow summarization of multiple AS PATHs to be  sent in BGP Updates.
// Type returns a string
func (obj *bgpAsPathSegment) HasType() bool {
	return obj.obj.Type != nil
}

func (obj *bgpAsPathSegment) SetType(value BgpAsPathSegmentTypeEnum) BgpAsPathSegment {
	intValue, ok := otg.BgpAsPathSegment_Type_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpAsPathSegmentTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpAsPathSegment_Type_Enum(intValue)
	obj.obj.Type = &enumValue

	return obj
}

// The AS numbers in this AS path segment.
// AsNumbers returns a []uint32
func (obj *bgpAsPathSegment) AsNumbers() []uint32 {
	if obj.obj.AsNumbers == nil {
		obj.obj.AsNumbers = make([]uint32, 0)
	}
	return obj.obj.AsNumbers
}

// The AS numbers in this AS path segment.
// SetAsNumbers sets the []uint32 value in the BgpAsPathSegment object
func (obj *bgpAsPathSegment) SetAsNumbers(value []uint32) BgpAsPathSegment {

	if obj.obj.AsNumbers == nil {
		obj.obj.AsNumbers = make([]uint32, 0)
	}
	obj.obj.AsNumbers = value

	return obj
}

func (obj *bgpAsPathSegment) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpAsPathSegment) setDefault() {
	if obj.obj.Type == nil {
		obj.SetType(BgpAsPathSegmentType.AS_SEQ)

	}

}

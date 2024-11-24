package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ResultBgpAsPathSegment *****
type resultBgpAsPathSegment struct {
	validation
	obj          *otg.ResultBgpAsPathSegment
	marshaller   marshalResultBgpAsPathSegment
	unMarshaller unMarshalResultBgpAsPathSegment
}

func NewResultBgpAsPathSegment() ResultBgpAsPathSegment {
	obj := resultBgpAsPathSegment{obj: &otg.ResultBgpAsPathSegment{}}
	obj.setDefault()
	return &obj
}

func (obj *resultBgpAsPathSegment) msg() *otg.ResultBgpAsPathSegment {
	return obj.obj
}

func (obj *resultBgpAsPathSegment) setMsg(msg *otg.ResultBgpAsPathSegment) ResultBgpAsPathSegment {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalresultBgpAsPathSegment struct {
	obj *resultBgpAsPathSegment
}

type marshalResultBgpAsPathSegment interface {
	// ToProto marshals ResultBgpAsPathSegment to protobuf object *otg.ResultBgpAsPathSegment
	ToProto() (*otg.ResultBgpAsPathSegment, error)
	// ToPbText marshals ResultBgpAsPathSegment to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ResultBgpAsPathSegment to YAML text
	ToYaml() (string, error)
	// ToJson marshals ResultBgpAsPathSegment to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals ResultBgpAsPathSegment to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalresultBgpAsPathSegment struct {
	obj *resultBgpAsPathSegment
}

type unMarshalResultBgpAsPathSegment interface {
	// FromProto unmarshals ResultBgpAsPathSegment from protobuf object *otg.ResultBgpAsPathSegment
	FromProto(msg *otg.ResultBgpAsPathSegment) (ResultBgpAsPathSegment, error)
	// FromPbText unmarshals ResultBgpAsPathSegment from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ResultBgpAsPathSegment from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ResultBgpAsPathSegment from JSON text
	FromJson(value string) error
}

func (obj *resultBgpAsPathSegment) Marshal() marshalResultBgpAsPathSegment {
	if obj.marshaller == nil {
		obj.marshaller = &marshalresultBgpAsPathSegment{obj: obj}
	}
	return obj.marshaller
}

func (obj *resultBgpAsPathSegment) Unmarshal() unMarshalResultBgpAsPathSegment {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalresultBgpAsPathSegment{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalresultBgpAsPathSegment) ToProto() (*otg.ResultBgpAsPathSegment, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalresultBgpAsPathSegment) FromProto(msg *otg.ResultBgpAsPathSegment) (ResultBgpAsPathSegment, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalresultBgpAsPathSegment) ToPbText() (string, error) {
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

func (m *unMarshalresultBgpAsPathSegment) FromPbText(value string) error {
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

func (m *marshalresultBgpAsPathSegment) ToYaml() (string, error) {
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

func (m *unMarshalresultBgpAsPathSegment) FromYaml(value string) error {
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

func (m *marshalresultBgpAsPathSegment) ToJsonRaw() (string, error) {
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

func (m *marshalresultBgpAsPathSegment) ToJson() (string, error) {
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

func (m *unMarshalresultBgpAsPathSegment) FromJson(value string) error {
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

func (obj *resultBgpAsPathSegment) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *resultBgpAsPathSegment) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *resultBgpAsPathSegment) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *resultBgpAsPathSegment) Clone() (ResultBgpAsPathSegment, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewResultBgpAsPathSegment()
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

// ResultBgpAsPathSegment is configuration for a single BGP AS path segment
type ResultBgpAsPathSegment interface {
	Validation
	// msg marshals ResultBgpAsPathSegment to protobuf object *otg.ResultBgpAsPathSegment
	// and doesn't set defaults
	msg() *otg.ResultBgpAsPathSegment
	// setMsg unmarshals ResultBgpAsPathSegment from protobuf object *otg.ResultBgpAsPathSegment
	// and doesn't set defaults
	setMsg(*otg.ResultBgpAsPathSegment) ResultBgpAsPathSegment
	// provides marshal interface
	Marshal() marshalResultBgpAsPathSegment
	// provides unmarshal interface
	Unmarshal() unMarshalResultBgpAsPathSegment
	// validate validates ResultBgpAsPathSegment
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ResultBgpAsPathSegment, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns ResultBgpAsPathSegmentTypeEnum, set in ResultBgpAsPathSegment
	Type() ResultBgpAsPathSegmentTypeEnum
	// SetType assigns ResultBgpAsPathSegmentTypeEnum provided by user to ResultBgpAsPathSegment
	SetType(value ResultBgpAsPathSegmentTypeEnum) ResultBgpAsPathSegment
	// HasType checks if Type has been set in ResultBgpAsPathSegment
	HasType() bool
	// AsNumbers returns []uint32, set in ResultBgpAsPathSegment.
	AsNumbers() []uint32
	// SetAsNumbers assigns []uint32 provided by user to ResultBgpAsPathSegment
	SetAsNumbers(value []uint32) ResultBgpAsPathSegment
}

type ResultBgpAsPathSegmentTypeEnum string

// Enum of Type on ResultBgpAsPathSegment
var ResultBgpAsPathSegmentType = struct {
	AS_SEQ        ResultBgpAsPathSegmentTypeEnum
	AS_SET        ResultBgpAsPathSegmentTypeEnum
	AS_CONFED_SEQ ResultBgpAsPathSegmentTypeEnum
	AS_CONFED_SET ResultBgpAsPathSegmentTypeEnum
}{
	AS_SEQ:        ResultBgpAsPathSegmentTypeEnum("as_seq"),
	AS_SET:        ResultBgpAsPathSegmentTypeEnum("as_set"),
	AS_CONFED_SEQ: ResultBgpAsPathSegmentTypeEnum("as_confed_seq"),
	AS_CONFED_SET: ResultBgpAsPathSegmentTypeEnum("as_confed_set"),
}

func (obj *resultBgpAsPathSegment) Type() ResultBgpAsPathSegmentTypeEnum {
	return ResultBgpAsPathSegmentTypeEnum(obj.obj.Type.Enum().String())
}

// AS sequence is the most common type of AS_PATH, it contains the  list of ASNs starting with the most recent ASN being added read  from left to right.
// The other three AS_PATH types are used for Confederations - AS_SET is the type of AS_PATH attribute that summarizes routes using using the aggregate-address command, allowing AS_PATHs to be  summarized in the update as well. - AS_CONFED_SEQ gives the list of ASNs in the path starting with the  most recent ASN to be added reading left to right - AS_CONFED_SET will allow summarization of multiple AS PATHs to be  sent in BGP Updates.
// Type returns a string
func (obj *resultBgpAsPathSegment) HasType() bool {
	return obj.obj.Type != nil
}

func (obj *resultBgpAsPathSegment) SetType(value ResultBgpAsPathSegmentTypeEnum) ResultBgpAsPathSegment {
	intValue, ok := otg.ResultBgpAsPathSegment_Type_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ResultBgpAsPathSegmentTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.ResultBgpAsPathSegment_Type_Enum(intValue)
	obj.obj.Type = &enumValue

	return obj
}

// The AS numbers in this AS path segment.
// AsNumbers returns a []uint32
func (obj *resultBgpAsPathSegment) AsNumbers() []uint32 {
	if obj.obj.AsNumbers == nil {
		obj.obj.AsNumbers = make([]uint32, 0)
	}
	return obj.obj.AsNumbers
}

// The AS numbers in this AS path segment.
// SetAsNumbers sets the []uint32 value in the ResultBgpAsPathSegment object
func (obj *resultBgpAsPathSegment) SetAsNumbers(value []uint32) ResultBgpAsPathSegment {

	if obj.obj.AsNumbers == nil {
		obj.obj.AsNumbers = make([]uint32, 0)
	}
	obj.obj.AsNumbers = value

	return obj
}

func (obj *resultBgpAsPathSegment) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *resultBgpAsPathSegment) setDefault() {

}

package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ResultExtendedCommunity *****
type resultExtendedCommunity struct {
	validation
	obj              *otg.ResultExtendedCommunity
	marshaller       marshalResultExtendedCommunity
	unMarshaller     unMarshalResultExtendedCommunity
	structuredHolder ResultExtendedCommunityStructured
}

func NewResultExtendedCommunity() ResultExtendedCommunity {
	obj := resultExtendedCommunity{obj: &otg.ResultExtendedCommunity{}}
	obj.setDefault()
	return &obj
}

func (obj *resultExtendedCommunity) msg() *otg.ResultExtendedCommunity {
	return obj.obj
}

func (obj *resultExtendedCommunity) setMsg(msg *otg.ResultExtendedCommunity) ResultExtendedCommunity {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalresultExtendedCommunity struct {
	obj *resultExtendedCommunity
}

type marshalResultExtendedCommunity interface {
	// ToProto marshals ResultExtendedCommunity to protobuf object *otg.ResultExtendedCommunity
	ToProto() (*otg.ResultExtendedCommunity, error)
	// ToPbText marshals ResultExtendedCommunity to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ResultExtendedCommunity to YAML text
	ToYaml() (string, error)
	// ToJson marshals ResultExtendedCommunity to JSON text
	ToJson() (string, error)
}

type unMarshalresultExtendedCommunity struct {
	obj *resultExtendedCommunity
}

type unMarshalResultExtendedCommunity interface {
	// FromProto unmarshals ResultExtendedCommunity from protobuf object *otg.ResultExtendedCommunity
	FromProto(msg *otg.ResultExtendedCommunity) (ResultExtendedCommunity, error)
	// FromPbText unmarshals ResultExtendedCommunity from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ResultExtendedCommunity from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ResultExtendedCommunity from JSON text
	FromJson(value string) error
}

func (obj *resultExtendedCommunity) Marshal() marshalResultExtendedCommunity {
	if obj.marshaller == nil {
		obj.marshaller = &marshalresultExtendedCommunity{obj: obj}
	}
	return obj.marshaller
}

func (obj *resultExtendedCommunity) Unmarshal() unMarshalResultExtendedCommunity {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalresultExtendedCommunity{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalresultExtendedCommunity) ToProto() (*otg.ResultExtendedCommunity, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalresultExtendedCommunity) FromProto(msg *otg.ResultExtendedCommunity) (ResultExtendedCommunity, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalresultExtendedCommunity) ToPbText() (string, error) {
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

func (m *unMarshalresultExtendedCommunity) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalresultExtendedCommunity) ToYaml() (string, error) {
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

func (m *unMarshalresultExtendedCommunity) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalresultExtendedCommunity) ToJson() (string, error) {
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

func (m *unMarshalresultExtendedCommunity) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *resultExtendedCommunity) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *resultExtendedCommunity) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *resultExtendedCommunity) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *resultExtendedCommunity) Clone() (ResultExtendedCommunity, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewResultExtendedCommunity()
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

func (obj *resultExtendedCommunity) setNil() {
	obj.structuredHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ResultExtendedCommunity is each received Extended Community attribute is available for retrieval in two forms. Support of the 'raw' format in which all 8 bytes (16 hex characters) is always present and available for use. In addition, if supported by the implementation, the Extended Community attribute may also be retrieved in the 'structured' format which is an optional field.
type ResultExtendedCommunity interface {
	Validation
	// msg marshals ResultExtendedCommunity to protobuf object *otg.ResultExtendedCommunity
	// and doesn't set defaults
	msg() *otg.ResultExtendedCommunity
	// setMsg unmarshals ResultExtendedCommunity from protobuf object *otg.ResultExtendedCommunity
	// and doesn't set defaults
	setMsg(*otg.ResultExtendedCommunity) ResultExtendedCommunity
	// provides marshal interface
	Marshal() marshalResultExtendedCommunity
	// provides unmarshal interface
	Unmarshal() unMarshalResultExtendedCommunity
	// validate validates ResultExtendedCommunity
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ResultExtendedCommunity, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Raw returns string, set in ResultExtendedCommunity.
	Raw() string
	// SetRaw assigns string provided by user to ResultExtendedCommunity
	SetRaw(value string) ResultExtendedCommunity
	// HasRaw checks if Raw has been set in ResultExtendedCommunity
	HasRaw() bool
	// Structured returns ResultExtendedCommunityStructured, set in ResultExtendedCommunity.
	// ResultExtendedCommunityStructured is the Extended Communities Attribute is a optional BGP attribute,defined in RFC4360 with the Type Code 16.
	// Community and Extended Communities  attributes are utilized to trigger routing decisions, such as acceptance, rejection,  preference, or redistribution.
	// An extended community is an 8-bytes value. It is divided into two main parts. The first 2 bytes of the community  encode a type and optonal sub-type field.
	// The last 6 bytes (or 7 bytes for types without a sub-type) carry a unique set of data in a format defined by the type and optional sub-type field.
	// Extended communities provide a larger  range for grouping or categorizing communities.
	Structured() ResultExtendedCommunityStructured
	// SetStructured assigns ResultExtendedCommunityStructured provided by user to ResultExtendedCommunity.
	// ResultExtendedCommunityStructured is the Extended Communities Attribute is a optional BGP attribute,defined in RFC4360 with the Type Code 16.
	// Community and Extended Communities  attributes are utilized to trigger routing decisions, such as acceptance, rejection,  preference, or redistribution.
	// An extended community is an 8-bytes value. It is divided into two main parts. The first 2 bytes of the community  encode a type and optonal sub-type field.
	// The last 6 bytes (or 7 bytes for types without a sub-type) carry a unique set of data in a format defined by the type and optional sub-type field.
	// Extended communities provide a larger  range for grouping or categorizing communities.
	SetStructured(value ResultExtendedCommunityStructured) ResultExtendedCommunity
	// HasStructured checks if Structured has been set in ResultExtendedCommunity
	HasStructured() bool
	setNil()
}

// The raw byte contents of the 8 bytes received in the Extended Community as 16 hex characters.
// Raw returns a string
func (obj *resultExtendedCommunity) Raw() string {

	return *obj.obj.Raw

}

// The raw byte contents of the 8 bytes received in the Extended Community as 16 hex characters.
// Raw returns a string
func (obj *resultExtendedCommunity) HasRaw() bool {
	return obj.obj.Raw != nil
}

// The raw byte contents of the 8 bytes received in the Extended Community as 16 hex characters.
// SetRaw sets the string value in the ResultExtendedCommunity object
func (obj *resultExtendedCommunity) SetRaw(value string) ResultExtendedCommunity {

	obj.obj.Raw = &value
	return obj
}

// description is TBD
// Structured returns a ResultExtendedCommunityStructured
func (obj *resultExtendedCommunity) Structured() ResultExtendedCommunityStructured {
	if obj.obj.Structured == nil {
		obj.obj.Structured = NewResultExtendedCommunityStructured().msg()
	}
	if obj.structuredHolder == nil {
		obj.structuredHolder = &resultExtendedCommunityStructured{obj: obj.obj.Structured}
	}
	return obj.structuredHolder
}

// description is TBD
// Structured returns a ResultExtendedCommunityStructured
func (obj *resultExtendedCommunity) HasStructured() bool {
	return obj.obj.Structured != nil
}

// description is TBD
// SetStructured sets the ResultExtendedCommunityStructured value in the ResultExtendedCommunity object
func (obj *resultExtendedCommunity) SetStructured(value ResultExtendedCommunityStructured) ResultExtendedCommunity {

	obj.structuredHolder = nil
	obj.obj.Structured = value.msg()

	return obj
}

func (obj *resultExtendedCommunity) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Raw != nil {

		if len(*obj.obj.Raw) > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"None <= length of ResultExtendedCommunity.Raw <= 16 but Got %d",
					len(*obj.obj.Raw)))
		}

	}

	if obj.obj.Structured != nil {

		obj.Structured().validateObj(vObj, set_default)
	}

}

func (obj *resultExtendedCommunity) setDefault() {

}

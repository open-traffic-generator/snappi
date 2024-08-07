package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Error *****
type _error struct {
	validation
	obj          *otg.Error
	marshaller   marshalError
	unMarshaller unMarshalError
}

func NewError() Error {
	obj := _error{obj: &otg.Error{}}
	obj.setDefault()
	return &obj
}

func (obj *_error) msg() *otg.Error {
	return obj.obj
}

func (obj *_error) setMsg(msg *otg.Error) Error {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshal_error struct {
	obj *_error
}

type marshalError interface {
	// ToProto marshals Error to protobuf object *otg.Error
	ToProto() (*otg.Error, error)
	// ToPbText marshals Error to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Error to YAML text
	ToYaml() (string, error)
	// ToJson marshals Error to JSON text
	ToJson() (string, error)
}

type unMarshal_error struct {
	obj *_error
}

type unMarshalError interface {
	// FromProto unmarshals Error from protobuf object *otg.Error
	FromProto(msg *otg.Error) (Error, error)
	// FromPbText unmarshals Error from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Error from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Error from JSON text
	FromJson(value string) error
}

func (obj *_error) Marshal() marshalError {
	if obj.marshaller == nil {
		obj.marshaller = &marshal_error{obj: obj}
	}
	return obj.marshaller
}

func (obj *_error) Unmarshal() unMarshalError {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshal_error{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshal_error) ToProto() (*otg.Error, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshal_error) FromProto(msg *otg.Error) (Error, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshal_error) ToPbText() (string, error) {
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

func (m *unMarshal_error) FromPbText(value string) error {
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

func (m *marshal_error) ToYaml() (string, error) {
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

func (m *unMarshal_error) FromYaml(value string) error {
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

func (m *marshal_error) ToJson() (string, error) {
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

func (m *unMarshal_error) FromJson(value string) error {
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

func (obj *_error) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *_error) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *_error) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *_error) Clone() (Error, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewError()
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

// Error is error response generated while serving API request.
type Error interface {
	Validation
	// msg marshals Error to protobuf object *otg.Error
	// and doesn't set defaults
	msg() *otg.Error
	// setMsg unmarshals Error from protobuf object *otg.Error
	// and doesn't set defaults
	setMsg(*otg.Error) Error
	// provides marshal interface
	Marshal() marshalError
	// provides unmarshal interface
	Unmarshal() unMarshalError
	// validate validates Error
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Error, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Code returns int32, set in Error.
	Code() int32
	// SetCode assigns int32 provided by user to Error
	SetCode(value int32) Error
	// Kind returns ErrorKindEnum, set in Error
	Kind() ErrorKindEnum
	// SetKind assigns ErrorKindEnum provided by user to Error
	SetKind(value ErrorKindEnum) Error
	// HasKind checks if Kind has been set in Error
	HasKind() bool
	// Errors returns []string, set in Error.
	Errors() []string
	// SetErrors assigns []string provided by user to Error
	SetErrors(value []string) Error
	// implement Error function for implementingnative Error Interface.
	Error() string
}

func (obj *_error) Error() string {
	json, err := obj.Marshal().ToJson()
	if err != nil {
		return fmt.Sprintf("could not convert Error to JSON: %v", err)
	}
	return json
}

// Numeric status code based on the underlying transport being used.
// The API server MUST set this code explicitly based on following references:
// - HTTP 4xx errors: https://datatracker.ietf.org/doc/html/rfc9110#section-15.5
// - HTTP 5xx errors: https://datatracker.ietf.org/doc/html/rfc9110#section-15.6
// - gRPC errors: https://grpc.github.io/grpc/core/md_doc_statuscodes.html
// Code returns a int32
func (obj *_error) Code() int32 {

	return *obj.obj.Code

}

// Numeric status code based on the underlying transport being used.
// The API server MUST set this code explicitly based on following references:
// - HTTP 4xx errors: https://datatracker.ietf.org/doc/html/rfc9110#section-15.5
// - HTTP 5xx errors: https://datatracker.ietf.org/doc/html/rfc9110#section-15.6
// - gRPC errors: https://grpc.github.io/grpc/core/md_doc_statuscodes.html
// SetCode sets the int32 value in the Error object
func (obj *_error) SetCode(value int32) Error {

	obj.obj.Code = &value
	return obj
}

type ErrorKindEnum string

// Enum of Kind on Error
var ErrorKind = struct {
	VALIDATION ErrorKindEnum
	INTERNAL   ErrorKindEnum
}{
	VALIDATION: ErrorKindEnum("validation"),
	INTERNAL:   ErrorKindEnum("internal"),
}

func (obj *_error) Kind() ErrorKindEnum {
	return ErrorKindEnum(obj.obj.Kind.Enum().String())
}

// Classification of error originating from within API server that may not be mapped to the value in `code`.
// Absence of this field may indicate that the error did not originate from within API server.
// Kind returns a string
func (obj *_error) HasKind() bool {
	return obj.obj.Kind != nil
}

func (obj *_error) SetKind(value ErrorKindEnum) Error {
	intValue, ok := otg.Error_Kind_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ErrorKindEnum", string(value)))
		return obj
	}
	enumValue := otg.Error_Kind_Enum(intValue)
	obj.obj.Kind = &enumValue

	return obj
}

// List of error messages generated while executing the request.
// Errors returns a []string
func (obj *_error) Errors() []string {
	if obj.obj.Errors == nil {
		obj.obj.Errors = make([]string, 0)
	}
	return obj.obj.Errors
}

// List of error messages generated while executing the request.
// SetErrors sets the []string value in the Error object
func (obj *_error) SetErrors(value []string) Error {

	if obj.obj.Errors == nil {
		obj.obj.Errors = make([]string, 0)
	}
	obj.obj.Errors = value

	return obj
}

func (obj *_error) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Code is required
	if obj.obj.Code == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Code is required field on interface Error")
	}
}

func (obj *_error) setDefault() {

}

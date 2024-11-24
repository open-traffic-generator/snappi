package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisInterfaceAuthentication *****
type isisInterfaceAuthentication struct {
	validation
	obj          *otg.IsisInterfaceAuthentication
	marshaller   marshalIsisInterfaceAuthentication
	unMarshaller unMarshalIsisInterfaceAuthentication
}

func NewIsisInterfaceAuthentication() IsisInterfaceAuthentication {
	obj := isisInterfaceAuthentication{obj: &otg.IsisInterfaceAuthentication{}}
	obj.setDefault()
	return &obj
}

func (obj *isisInterfaceAuthentication) msg() *otg.IsisInterfaceAuthentication {
	return obj.obj
}

func (obj *isisInterfaceAuthentication) setMsg(msg *otg.IsisInterfaceAuthentication) IsisInterfaceAuthentication {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisInterfaceAuthentication struct {
	obj *isisInterfaceAuthentication
}

type marshalIsisInterfaceAuthentication interface {
	// ToProto marshals IsisInterfaceAuthentication to protobuf object *otg.IsisInterfaceAuthentication
	ToProto() (*otg.IsisInterfaceAuthentication, error)
	// ToPbText marshals IsisInterfaceAuthentication to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisInterfaceAuthentication to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisInterfaceAuthentication to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisInterfaceAuthentication to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisInterfaceAuthentication struct {
	obj *isisInterfaceAuthentication
}

type unMarshalIsisInterfaceAuthentication interface {
	// FromProto unmarshals IsisInterfaceAuthentication from protobuf object *otg.IsisInterfaceAuthentication
	FromProto(msg *otg.IsisInterfaceAuthentication) (IsisInterfaceAuthentication, error)
	// FromPbText unmarshals IsisInterfaceAuthentication from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisInterfaceAuthentication from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisInterfaceAuthentication from JSON text
	FromJson(value string) error
}

func (obj *isisInterfaceAuthentication) Marshal() marshalIsisInterfaceAuthentication {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisInterfaceAuthentication{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisInterfaceAuthentication) Unmarshal() unMarshalIsisInterfaceAuthentication {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisInterfaceAuthentication{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisInterfaceAuthentication) ToProto() (*otg.IsisInterfaceAuthentication, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisInterfaceAuthentication) FromProto(msg *otg.IsisInterfaceAuthentication) (IsisInterfaceAuthentication, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisInterfaceAuthentication) ToPbText() (string, error) {
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

func (m *unMarshalisisInterfaceAuthentication) FromPbText(value string) error {
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

func (m *marshalisisInterfaceAuthentication) ToYaml() (string, error) {
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

func (m *unMarshalisisInterfaceAuthentication) FromYaml(value string) error {
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

func (m *marshalisisInterfaceAuthentication) ToJsonRaw() (string, error) {
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

func (m *marshalisisInterfaceAuthentication) ToJson() (string, error) {
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

func (m *unMarshalisisInterfaceAuthentication) FromJson(value string) error {
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

func (obj *isisInterfaceAuthentication) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisInterfaceAuthentication) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisInterfaceAuthentication) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisInterfaceAuthentication) Clone() (IsisInterfaceAuthentication, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisInterfaceAuthentication()
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

// IsisInterfaceAuthentication is optional container for circuit authentication properties.
type IsisInterfaceAuthentication interface {
	Validation
	// msg marshals IsisInterfaceAuthentication to protobuf object *otg.IsisInterfaceAuthentication
	// and doesn't set defaults
	msg() *otg.IsisInterfaceAuthentication
	// setMsg unmarshals IsisInterfaceAuthentication from protobuf object *otg.IsisInterfaceAuthentication
	// and doesn't set defaults
	setMsg(*otg.IsisInterfaceAuthentication) IsisInterfaceAuthentication
	// provides marshal interface
	Marshal() marshalIsisInterfaceAuthentication
	// provides unmarshal interface
	Unmarshal() unMarshalIsisInterfaceAuthentication
	// validate validates IsisInterfaceAuthentication
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisInterfaceAuthentication, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// AuthType returns IsisInterfaceAuthenticationAuthTypeEnum, set in IsisInterfaceAuthentication
	AuthType() IsisInterfaceAuthenticationAuthTypeEnum
	// SetAuthType assigns IsisInterfaceAuthenticationAuthTypeEnum provided by user to IsisInterfaceAuthentication
	SetAuthType(value IsisInterfaceAuthenticationAuthTypeEnum) IsisInterfaceAuthentication
	// Md5 returns string, set in IsisInterfaceAuthentication.
	Md5() string
	// SetMd5 assigns string provided by user to IsisInterfaceAuthentication
	SetMd5(value string) IsisInterfaceAuthentication
	// HasMd5 checks if Md5 has been set in IsisInterfaceAuthentication
	HasMd5() bool
	// Password returns string, set in IsisInterfaceAuthentication.
	Password() string
	// SetPassword assigns string provided by user to IsisInterfaceAuthentication
	SetPassword(value string) IsisInterfaceAuthentication
	// HasPassword checks if Password has been set in IsisInterfaceAuthentication
	HasPassword() bool
}

type IsisInterfaceAuthenticationAuthTypeEnum string

// Enum of AuthType on IsisInterfaceAuthentication
var IsisInterfaceAuthenticationAuthType = struct {
	MD5      IsisInterfaceAuthenticationAuthTypeEnum
	PASSWORD IsisInterfaceAuthenticationAuthTypeEnum
}{
	MD5:      IsisInterfaceAuthenticationAuthTypeEnum("md5"),
	PASSWORD: IsisInterfaceAuthenticationAuthTypeEnum("password"),
}

func (obj *isisInterfaceAuthentication) AuthType() IsisInterfaceAuthenticationAuthTypeEnum {
	return IsisInterfaceAuthenticationAuthTypeEnum(obj.obj.AuthType.Enum().String())
}

func (obj *isisInterfaceAuthentication) SetAuthType(value IsisInterfaceAuthenticationAuthTypeEnum) IsisInterfaceAuthentication {
	intValue, ok := otg.IsisInterfaceAuthentication_AuthType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisInterfaceAuthenticationAuthTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisInterfaceAuthentication_AuthType_Enum(intValue)
	obj.obj.AuthType = &enumValue

	return obj
}

// MD5 key to be used for authentication.
// Md5 returns a string
func (obj *isisInterfaceAuthentication) Md5() string {

	return *obj.obj.Md5

}

// MD5 key to be used for authentication.
// Md5 returns a string
func (obj *isisInterfaceAuthentication) HasMd5() bool {
	return obj.obj.Md5 != nil
}

// MD5 key to be used for authentication.
// SetMd5 sets the string value in the IsisInterfaceAuthentication object
func (obj *isisInterfaceAuthentication) SetMd5(value string) IsisInterfaceAuthentication {

	obj.obj.Md5 = &value
	return obj
}

// The password, in clear text, to be used for Authentication.
// Password returns a string
func (obj *isisInterfaceAuthentication) Password() string {

	return *obj.obj.Password

}

// The password, in clear text, to be used for Authentication.
// Password returns a string
func (obj *isisInterfaceAuthentication) HasPassword() bool {
	return obj.obj.Password != nil
}

// The password, in clear text, to be used for Authentication.
// SetPassword sets the string value in the IsisInterfaceAuthentication object
func (obj *isisInterfaceAuthentication) SetPassword(value string) IsisInterfaceAuthentication {

	obj.obj.Password = &value
	return obj
}

func (obj *isisInterfaceAuthentication) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// AuthType is required
	if obj.obj.AuthType == nil {
		vObj.validationErrors = append(vObj.validationErrors, "AuthType is required field on interface IsisInterfaceAuthentication")
	}

	if obj.obj.Md5 != nil {

		if len(*obj.obj.Md5) < 0 || len(*obj.obj.Md5) > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"0 <= length of IsisInterfaceAuthentication.Md5 <= 255 but Got %d",
					len(*obj.obj.Md5)))
		}

	}

	if obj.obj.Password != nil {

		if len(*obj.obj.Password) < 0 || len(*obj.obj.Password) > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"0 <= length of IsisInterfaceAuthentication.Password <= 255 but Got %d",
					len(*obj.obj.Password)))
		}

	}

}

func (obj *isisInterfaceAuthentication) setDefault() {

}

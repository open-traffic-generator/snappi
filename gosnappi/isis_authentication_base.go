package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisAuthenticationBase *****
type isisAuthenticationBase struct {
	validation
	obj          *otg.IsisAuthenticationBase
	marshaller   marshalIsisAuthenticationBase
	unMarshaller unMarshalIsisAuthenticationBase
}

func NewIsisAuthenticationBase() IsisAuthenticationBase {
	obj := isisAuthenticationBase{obj: &otg.IsisAuthenticationBase{}}
	obj.setDefault()
	return &obj
}

func (obj *isisAuthenticationBase) msg() *otg.IsisAuthenticationBase {
	return obj.obj
}

func (obj *isisAuthenticationBase) setMsg(msg *otg.IsisAuthenticationBase) IsisAuthenticationBase {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisAuthenticationBase struct {
	obj *isisAuthenticationBase
}

type marshalIsisAuthenticationBase interface {
	// ToProto marshals IsisAuthenticationBase to protobuf object *otg.IsisAuthenticationBase
	ToProto() (*otg.IsisAuthenticationBase, error)
	// ToPbText marshals IsisAuthenticationBase to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisAuthenticationBase to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisAuthenticationBase to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisAuthenticationBase to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisAuthenticationBase struct {
	obj *isisAuthenticationBase
}

type unMarshalIsisAuthenticationBase interface {
	// FromProto unmarshals IsisAuthenticationBase from protobuf object *otg.IsisAuthenticationBase
	FromProto(msg *otg.IsisAuthenticationBase) (IsisAuthenticationBase, error)
	// FromPbText unmarshals IsisAuthenticationBase from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisAuthenticationBase from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisAuthenticationBase from JSON text
	FromJson(value string) error
}

func (obj *isisAuthenticationBase) Marshal() marshalIsisAuthenticationBase {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisAuthenticationBase{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisAuthenticationBase) Unmarshal() unMarshalIsisAuthenticationBase {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisAuthenticationBase{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisAuthenticationBase) ToProto() (*otg.IsisAuthenticationBase, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisAuthenticationBase) FromProto(msg *otg.IsisAuthenticationBase) (IsisAuthenticationBase, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisAuthenticationBase) ToPbText() (string, error) {
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

func (m *unMarshalisisAuthenticationBase) FromPbText(value string) error {
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

func (m *marshalisisAuthenticationBase) ToYaml() (string, error) {
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

func (m *unMarshalisisAuthenticationBase) FromYaml(value string) error {
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

func (m *marshalisisAuthenticationBase) ToJsonRaw() (string, error) {
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

func (m *marshalisisAuthenticationBase) ToJson() (string, error) {
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

func (m *unMarshalisisAuthenticationBase) FromJson(value string) error {
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

func (obj *isisAuthenticationBase) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisAuthenticationBase) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisAuthenticationBase) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisAuthenticationBase) Clone() (IsisAuthenticationBase, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisAuthenticationBase()
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

// IsisAuthenticationBase is optional container for ISIS authentication properties.
type IsisAuthenticationBase interface {
	Validation
	// msg marshals IsisAuthenticationBase to protobuf object *otg.IsisAuthenticationBase
	// and doesn't set defaults
	msg() *otg.IsisAuthenticationBase
	// setMsg unmarshals IsisAuthenticationBase from protobuf object *otg.IsisAuthenticationBase
	// and doesn't set defaults
	setMsg(*otg.IsisAuthenticationBase) IsisAuthenticationBase
	// provides marshal interface
	Marshal() marshalIsisAuthenticationBase
	// provides unmarshal interface
	Unmarshal() unMarshalIsisAuthenticationBase
	// validate validates IsisAuthenticationBase
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisAuthenticationBase, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// AuthType returns IsisAuthenticationBaseAuthTypeEnum, set in IsisAuthenticationBase
	AuthType() IsisAuthenticationBaseAuthTypeEnum
	// SetAuthType assigns IsisAuthenticationBaseAuthTypeEnum provided by user to IsisAuthenticationBase
	SetAuthType(value IsisAuthenticationBaseAuthTypeEnum) IsisAuthenticationBase
	// Md5 returns string, set in IsisAuthenticationBase.
	Md5() string
	// SetMd5 assigns string provided by user to IsisAuthenticationBase
	SetMd5(value string) IsisAuthenticationBase
	// HasMd5 checks if Md5 has been set in IsisAuthenticationBase
	HasMd5() bool
	// Password returns string, set in IsisAuthenticationBase.
	Password() string
	// SetPassword assigns string provided by user to IsisAuthenticationBase
	SetPassword(value string) IsisAuthenticationBase
	// HasPassword checks if Password has been set in IsisAuthenticationBase
	HasPassword() bool
}

type IsisAuthenticationBaseAuthTypeEnum string

// Enum of AuthType on IsisAuthenticationBase
var IsisAuthenticationBaseAuthType = struct {
	MD5      IsisAuthenticationBaseAuthTypeEnum
	PASSWORD IsisAuthenticationBaseAuthTypeEnum
}{
	MD5:      IsisAuthenticationBaseAuthTypeEnum("md5"),
	PASSWORD: IsisAuthenticationBaseAuthTypeEnum("password"),
}

func (obj *isisAuthenticationBase) AuthType() IsisAuthenticationBaseAuthTypeEnum {
	return IsisAuthenticationBaseAuthTypeEnum(obj.obj.AuthType.Enum().String())
}

func (obj *isisAuthenticationBase) SetAuthType(value IsisAuthenticationBaseAuthTypeEnum) IsisAuthenticationBase {
	intValue, ok := otg.IsisAuthenticationBase_AuthType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisAuthenticationBaseAuthTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisAuthenticationBase_AuthType_Enum(intValue)
	obj.obj.AuthType = &enumValue

	return obj
}

// Authentication as an MD5 key.
// Md5 returns a string
func (obj *isisAuthenticationBase) Md5() string {

	return *obj.obj.Md5

}

// Authentication as an MD5 key.
// Md5 returns a string
func (obj *isisAuthenticationBase) HasMd5() bool {
	return obj.obj.Md5 != nil
}

// Authentication as an MD5 key.
// SetMd5 sets the string value in the IsisAuthenticationBase object
func (obj *isisAuthenticationBase) SetMd5(value string) IsisAuthenticationBase {

	obj.obj.Md5 = &value
	return obj
}

// Authentication as a clear text password.
// Password returns a string
func (obj *isisAuthenticationBase) Password() string {

	return *obj.obj.Password

}

// Authentication as a clear text password.
// Password returns a string
func (obj *isisAuthenticationBase) HasPassword() bool {
	return obj.obj.Password != nil
}

// Authentication as a clear text password.
// SetPassword sets the string value in the IsisAuthenticationBase object
func (obj *isisAuthenticationBase) SetPassword(value string) IsisAuthenticationBase {

	obj.obj.Password = &value
	return obj
}

func (obj *isisAuthenticationBase) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// AuthType is required
	if obj.obj.AuthType == nil {
		vObj.validationErrors = append(vObj.validationErrors, "AuthType is required field on interface IsisAuthenticationBase")
	}

	if obj.obj.Md5 != nil {

		if len(*obj.obj.Md5) < 0 || len(*obj.obj.Md5) > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"0 <= length of IsisAuthenticationBase.Md5 <= 255 but Got %d",
					len(*obj.obj.Md5)))
		}

	}

	if obj.obj.Password != nil {

		if len(*obj.obj.Password) < 0 || len(*obj.obj.Password) > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"0 <= length of IsisAuthenticationBase.Password <= 255 but Got %d",
					len(*obj.obj.Password)))
		}

	}

}

func (obj *isisAuthenticationBase) setDefault() {

}

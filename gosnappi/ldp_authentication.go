package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LdpAuthentication *****
type ldpAuthentication struct {
	validation
	obj          *otg.LdpAuthentication
	marshaller   marshalLdpAuthentication
	unMarshaller unMarshalLdpAuthentication
}

func NewLdpAuthentication() LdpAuthentication {
	obj := ldpAuthentication{obj: &otg.LdpAuthentication{}}
	obj.setDefault()
	return &obj
}

func (obj *ldpAuthentication) msg() *otg.LdpAuthentication {
	return obj.obj
}

func (obj *ldpAuthentication) setMsg(msg *otg.LdpAuthentication) LdpAuthentication {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalldpAuthentication struct {
	obj *ldpAuthentication
}

type marshalLdpAuthentication interface {
	// ToProto marshals LdpAuthentication to protobuf object *otg.LdpAuthentication
	ToProto() (*otg.LdpAuthentication, error)
	// ToPbText marshals LdpAuthentication to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LdpAuthentication to YAML text
	ToYaml() (string, error)
	// ToJson marshals LdpAuthentication to JSON text
	ToJson() (string, error)
}

type unMarshalldpAuthentication struct {
	obj *ldpAuthentication
}

type unMarshalLdpAuthentication interface {
	// FromProto unmarshals LdpAuthentication from protobuf object *otg.LdpAuthentication
	FromProto(msg *otg.LdpAuthentication) (LdpAuthentication, error)
	// FromPbText unmarshals LdpAuthentication from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LdpAuthentication from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LdpAuthentication from JSON text
	FromJson(value string) error
}

func (obj *ldpAuthentication) Marshal() marshalLdpAuthentication {
	if obj.marshaller == nil {
		obj.marshaller = &marshalldpAuthentication{obj: obj}
	}
	return obj.marshaller
}

func (obj *ldpAuthentication) Unmarshal() unMarshalLdpAuthentication {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalldpAuthentication{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalldpAuthentication) ToProto() (*otg.LdpAuthentication, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalldpAuthentication) FromProto(msg *otg.LdpAuthentication) (LdpAuthentication, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalldpAuthentication) ToPbText() (string, error) {
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

func (m *unMarshalldpAuthentication) FromPbText(value string) error {
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

func (m *marshalldpAuthentication) ToYaml() (string, error) {
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

func (m *unMarshalldpAuthentication) FromYaml(value string) error {
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

func (m *marshalldpAuthentication) ToJson() (string, error) {
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

func (m *unMarshalldpAuthentication) FromJson(value string) error {
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

func (obj *ldpAuthentication) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ldpAuthentication) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ldpAuthentication) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ldpAuthentication) Clone() (LdpAuthentication, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLdpAuthentication()
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

// LdpAuthentication is lDP session authentication. When this object is present, the TCP MD5 Signature Option (RFC 2385) is used on the LDP session TCP connection (RFC 5036 Section 2.9). Omit the object to disable authentication.
type LdpAuthentication interface {
	Validation
	// msg marshals LdpAuthentication to protobuf object *otg.LdpAuthentication
	// and doesn't set defaults
	msg() *otg.LdpAuthentication
	// setMsg unmarshals LdpAuthentication from protobuf object *otg.LdpAuthentication
	// and doesn't set defaults
	setMsg(*otg.LdpAuthentication) LdpAuthentication
	// provides marshal interface
	Marshal() marshalLdpAuthentication
	// provides unmarshal interface
	Unmarshal() unMarshalLdpAuthentication
	// validate validates LdpAuthentication
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LdpAuthentication, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns LdpAuthenticationChoiceEnum, set in LdpAuthentication
	Choice() LdpAuthenticationChoiceEnum
	// setChoice assigns LdpAuthenticationChoiceEnum provided by user to LdpAuthentication
	setChoice(value LdpAuthenticationChoiceEnum) LdpAuthentication
	// HasChoice checks if Choice has been set in LdpAuthentication
	HasChoice() bool
	// Md5 returns string, set in LdpAuthentication.
	Md5() string
	// SetMd5 assigns string provided by user to LdpAuthentication
	SetMd5(value string) LdpAuthentication
	// HasMd5 checks if Md5 has been set in LdpAuthentication
	HasMd5() bool
}

type LdpAuthenticationChoiceEnum string

// Enum of Choice on LdpAuthentication
var LdpAuthenticationChoice = struct {
	MD5 LdpAuthenticationChoiceEnum
}{
	MD5: LdpAuthenticationChoiceEnum("md5"),
}

func (obj *ldpAuthentication) Choice() LdpAuthenticationChoiceEnum {
	return LdpAuthenticationChoiceEnum(obj.obj.Choice.Enum().String())
}

// The authentication mechanism.
// Choice returns a string
func (obj *ldpAuthentication) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ldpAuthentication) setChoice(value LdpAuthenticationChoiceEnum) LdpAuthentication {
	intValue, ok := otg.LdpAuthentication_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LdpAuthenticationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.LdpAuthentication_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Md5 = nil
	return obj
}

// The shared key used to compute the TCP MD5 signature (RFC 2385). A key must be set when choice is md5.
// Md5 returns a string
func (obj *ldpAuthentication) Md5() string {

	if obj.obj.Md5 == nil {
		obj.setChoice(LdpAuthenticationChoice.MD5)
	}

	return *obj.obj.Md5

}

// The shared key used to compute the TCP MD5 signature (RFC 2385). A key must be set when choice is md5.
// Md5 returns a string
func (obj *ldpAuthentication) HasMd5() bool {
	return obj.obj.Md5 != nil
}

// The shared key used to compute the TCP MD5 signature (RFC 2385). A key must be set when choice is md5.
// SetMd5 sets the string value in the LdpAuthentication object
func (obj *ldpAuthentication) SetMd5(value string) LdpAuthentication {
	obj.setChoice(LdpAuthenticationChoice.MD5)
	obj.obj.Md5 = &value
	return obj
}

func (obj *ldpAuthentication) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Md5 != nil {

		if len(*obj.obj.Md5) < 1 || len(*obj.obj.Md5) > 80 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of LdpAuthentication.Md5 <= 80 but Got %d",
					len(*obj.obj.Md5)))
		}

	}

}

func (obj *ldpAuthentication) setDefault() {
	var choices_set int = 0
	var choice LdpAuthenticationChoiceEnum

	if obj.obj.Md5 != nil {
		choices_set += 1
		choice = LdpAuthenticationChoice.MD5
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(LdpAuthenticationChoice.MD5)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in LdpAuthentication")
			}
		} else {
			intVal := otg.LdpAuthentication_Choice_Enum_value[string(choice)]
			enumValue := otg.LdpAuthentication_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}

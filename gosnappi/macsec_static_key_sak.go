package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecStaticKeySak *****
type macsecStaticKeySak struct {
	validation
	obj          *otg.MacsecStaticKeySak
	marshaller   marshalMacsecStaticKeySak
	unMarshaller unMarshalMacsecStaticKeySak
}

func NewMacsecStaticKeySak() MacsecStaticKeySak {
	obj := macsecStaticKeySak{obj: &otg.MacsecStaticKeySak{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecStaticKeySak) msg() *otg.MacsecStaticKeySak {
	return obj.obj
}

func (obj *macsecStaticKeySak) setMsg(msg *otg.MacsecStaticKeySak) MacsecStaticKeySak {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecStaticKeySak struct {
	obj *macsecStaticKeySak
}

type marshalMacsecStaticKeySak interface {
	// ToProto marshals MacsecStaticKeySak to protobuf object *otg.MacsecStaticKeySak
	ToProto() (*otg.MacsecStaticKeySak, error)
	// ToPbText marshals MacsecStaticKeySak to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecStaticKeySak to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecStaticKeySak to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecStaticKeySak struct {
	obj *macsecStaticKeySak
}

type unMarshalMacsecStaticKeySak interface {
	// FromProto unmarshals MacsecStaticKeySak from protobuf object *otg.MacsecStaticKeySak
	FromProto(msg *otg.MacsecStaticKeySak) (MacsecStaticKeySak, error)
	// FromPbText unmarshals MacsecStaticKeySak from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecStaticKeySak from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecStaticKeySak from JSON text
	FromJson(value string) error
}

func (obj *macsecStaticKeySak) Marshal() marshalMacsecStaticKeySak {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecStaticKeySak{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecStaticKeySak) Unmarshal() unMarshalMacsecStaticKeySak {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecStaticKeySak{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecStaticKeySak) ToProto() (*otg.MacsecStaticKeySak, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecStaticKeySak) FromProto(msg *otg.MacsecStaticKeySak) (MacsecStaticKeySak, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecStaticKeySak) ToPbText() (string, error) {
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

func (m *unMarshalmacsecStaticKeySak) FromPbText(value string) error {
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

func (m *marshalmacsecStaticKeySak) ToYaml() (string, error) {
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

func (m *unMarshalmacsecStaticKeySak) FromYaml(value string) error {
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

func (m *marshalmacsecStaticKeySak) ToJson() (string, error) {
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

func (m *unMarshalmacsecStaticKeySak) FromJson(value string) error {
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

func (obj *macsecStaticKeySak) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecStaticKeySak) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecStaticKeySak) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecStaticKeySak) Clone() (MacsecStaticKeySak, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecStaticKeySak()
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

// MacsecStaticKeySak is the container for SAK.
type MacsecStaticKeySak interface {
	Validation
	// msg marshals MacsecStaticKeySak to protobuf object *otg.MacsecStaticKeySak
	// and doesn't set defaults
	msg() *otg.MacsecStaticKeySak
	// setMsg unmarshals MacsecStaticKeySak from protobuf object *otg.MacsecStaticKeySak
	// and doesn't set defaults
	setMsg(*otg.MacsecStaticKeySak) MacsecStaticKeySak
	// provides marshal interface
	Marshal() marshalMacsecStaticKeySak
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecStaticKeySak
	// validate validates MacsecStaticKeySak
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecStaticKeySak, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Sak returns string, set in MacsecStaticKeySak.
	Sak() string
	// SetSak assigns string provided by user to MacsecStaticKeySak
	SetSak(value string) MacsecStaticKeySak
	// HasSak checks if Sak has been set in MacsecStaticKeySak
	HasSak() bool
	// Ssci returns string, set in MacsecStaticKeySak.
	Ssci() string
	// SetSsci assigns string provided by user to MacsecStaticKeySak
	SetSsci(value string) MacsecStaticKeySak
	// HasSsci checks if Ssci has been set in MacsecStaticKeySak
	HasSsci() bool
	// Salt returns string, set in MacsecStaticKeySak.
	Salt() string
	// SetSalt assigns string provided by user to MacsecStaticKeySak
	SetSalt(value string) MacsecStaticKeySak
	// HasSalt checks if Salt has been set in MacsecStaticKeySak
	HasSalt() bool
}

// Secure association key(SAK) bits as hex string. Either 128 bits or 256 bits depending on the citpher suite chosen.
// Sak returns a string
func (obj *macsecStaticKeySak) Sak() string {

	return *obj.obj.Sak

}

// Secure association key(SAK) bits as hex string. Either 128 bits or 256 bits depending on the citpher suite chosen.
// Sak returns a string
func (obj *macsecStaticKeySak) HasSak() bool {
	return obj.obj.Sak != nil
}

// Secure association key(SAK) bits as hex string. Either 128 bits or 256 bits depending on the citpher suite chosen.
// SetSak sets the string value in the MacsecStaticKeySak object
func (obj *macsecStaticKeySak) SetSak(value string) MacsecStaticKeySak {

	obj.obj.Sak = &value
	return obj
}

// 4 bytes short SCI(SSCI) used in case of XPN cipher suites.
// Ssci returns a string
func (obj *macsecStaticKeySak) Ssci() string {

	return *obj.obj.Ssci

}

// 4 bytes short SCI(SSCI) used in case of XPN cipher suites.
// Ssci returns a string
func (obj *macsecStaticKeySak) HasSsci() bool {
	return obj.obj.Ssci != nil
}

// 4 bytes short SCI(SSCI) used in case of XPN cipher suites.
// SetSsci sets the string value in the MacsecStaticKeySak object
func (obj *macsecStaticKeySak) SetSsci(value string) MacsecStaticKeySak {

	obj.obj.Ssci = &value
	return obj
}

// 12 bytes salt used in case of XPN cipher suites.
// Salt returns a string
func (obj *macsecStaticKeySak) Salt() string {

	return *obj.obj.Salt

}

// 12 bytes salt used in case of XPN cipher suites.
// Salt returns a string
func (obj *macsecStaticKeySak) HasSalt() bool {
	return obj.obj.Salt != nil
}

// 12 bytes salt used in case of XPN cipher suites.
// SetSalt sets the string value in the MacsecStaticKeySak object
func (obj *macsecStaticKeySak) SetSalt(value string) MacsecStaticKeySak {

	obj.obj.Salt = &value
	return obj
}

func (obj *macsecStaticKeySak) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Sak != nil {

		if len(*obj.obj.Sak) < 16 || len(*obj.obj.Sak) > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"16 <= length of MacsecStaticKeySak.Sak <= 32 but Got %d",
					len(*obj.obj.Sak)))
		}

	}

	if obj.obj.Ssci != nil {

		if len(*obj.obj.Ssci) < 4 || len(*obj.obj.Ssci) > 4 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"4 <= length of MacsecStaticKeySak.Ssci <= 4 but Got %d",
					len(*obj.obj.Ssci)))
		}

	}

	if obj.obj.Salt != nil {

		if len(*obj.obj.Salt) < 12 || len(*obj.obj.Salt) > 12 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"12 <= length of MacsecStaticKeySak.Salt <= 12 but Got %d",
					len(*obj.obj.Salt)))
		}

	}

}

func (obj *macsecStaticKeySak) setDefault() {
	if obj.obj.Sak == nil {
		obj.SetSak("F123456789ABCDEF0123456789ABCDEF")
	}
	if obj.obj.Ssci == nil {
		obj.SetSsci("00000001")
	}
	if obj.obj.Salt == nil {
		obj.SetSalt("000000000000000000000001")
	}

}

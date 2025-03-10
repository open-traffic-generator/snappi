package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2AuthenticationMd5 *****
type ospfv2AuthenticationMd5 struct {
	validation
	obj          *otg.Ospfv2AuthenticationMd5
	marshaller   marshalOspfv2AuthenticationMd5
	unMarshaller unMarshalOspfv2AuthenticationMd5
}

func NewOspfv2AuthenticationMd5() Ospfv2AuthenticationMd5 {
	obj := ospfv2AuthenticationMd5{obj: &otg.Ospfv2AuthenticationMd5{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2AuthenticationMd5) msg() *otg.Ospfv2AuthenticationMd5 {
	return obj.obj
}

func (obj *ospfv2AuthenticationMd5) setMsg(msg *otg.Ospfv2AuthenticationMd5) Ospfv2AuthenticationMd5 {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2AuthenticationMd5 struct {
	obj *ospfv2AuthenticationMd5
}

type marshalOspfv2AuthenticationMd5 interface {
	// ToProto marshals Ospfv2AuthenticationMd5 to protobuf object *otg.Ospfv2AuthenticationMd5
	ToProto() (*otg.Ospfv2AuthenticationMd5, error)
	// ToPbText marshals Ospfv2AuthenticationMd5 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2AuthenticationMd5 to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2AuthenticationMd5 to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv2AuthenticationMd5 to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv2AuthenticationMd5 struct {
	obj *ospfv2AuthenticationMd5
}

type unMarshalOspfv2AuthenticationMd5 interface {
	// FromProto unmarshals Ospfv2AuthenticationMd5 from protobuf object *otg.Ospfv2AuthenticationMd5
	FromProto(msg *otg.Ospfv2AuthenticationMd5) (Ospfv2AuthenticationMd5, error)
	// FromPbText unmarshals Ospfv2AuthenticationMd5 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2AuthenticationMd5 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2AuthenticationMd5 from JSON text
	FromJson(value string) error
}

func (obj *ospfv2AuthenticationMd5) Marshal() marshalOspfv2AuthenticationMd5 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2AuthenticationMd5{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2AuthenticationMd5) Unmarshal() unMarshalOspfv2AuthenticationMd5 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2AuthenticationMd5{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2AuthenticationMd5) ToProto() (*otg.Ospfv2AuthenticationMd5, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2AuthenticationMd5) FromProto(msg *otg.Ospfv2AuthenticationMd5) (Ospfv2AuthenticationMd5, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2AuthenticationMd5) ToPbText() (string, error) {
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

func (m *unMarshalospfv2AuthenticationMd5) FromPbText(value string) error {
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

func (m *marshalospfv2AuthenticationMd5) ToYaml() (string, error) {
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

func (m *unMarshalospfv2AuthenticationMd5) FromYaml(value string) error {
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

func (m *marshalospfv2AuthenticationMd5) ToJsonRaw() (string, error) {
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

func (m *marshalospfv2AuthenticationMd5) ToJson() (string, error) {
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

func (m *unMarshalospfv2AuthenticationMd5) FromJson(value string) error {
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

func (obj *ospfv2AuthenticationMd5) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2AuthenticationMd5) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2AuthenticationMd5) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2AuthenticationMd5) Clone() (Ospfv2AuthenticationMd5, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2AuthenticationMd5()
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

// Ospfv2AuthenticationMd5 is container of Cryptographic authentication.
// If the authentication type is of 'md5' then 'md5_key_id' and 'md5_key'
// both are to be configured. A shared secret key is configured in all routers attached to a common network/subnet.
// For each OSPF protocol packet, the key is used to generate/verify a "message digest" that is appended to the end
// of the OSPF packet.
type Ospfv2AuthenticationMd5 interface {
	Validation
	// msg marshals Ospfv2AuthenticationMd5 to protobuf object *otg.Ospfv2AuthenticationMd5
	// and doesn't set defaults
	msg() *otg.Ospfv2AuthenticationMd5
	// setMsg unmarshals Ospfv2AuthenticationMd5 from protobuf object *otg.Ospfv2AuthenticationMd5
	// and doesn't set defaults
	setMsg(*otg.Ospfv2AuthenticationMd5) Ospfv2AuthenticationMd5
	// provides marshal interface
	Marshal() marshalOspfv2AuthenticationMd5
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2AuthenticationMd5
	// validate validates Ospfv2AuthenticationMd5
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2AuthenticationMd5, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// KeyId returns uint32, set in Ospfv2AuthenticationMd5.
	KeyId() uint32
	// SetKeyId assigns uint32 provided by user to Ospfv2AuthenticationMd5
	SetKeyId(value uint32) Ospfv2AuthenticationMd5
	// HasKeyId checks if KeyId has been set in Ospfv2AuthenticationMd5
	HasKeyId() bool
	// Key returns string, set in Ospfv2AuthenticationMd5.
	Key() string
	// SetKey assigns string provided by user to Ospfv2AuthenticationMd5
	SetKey(value string) Ospfv2AuthenticationMd5
	// HasKey checks if Key has been set in Ospfv2AuthenticationMd5
	HasKey() bool
}

// The unique MD5 Key Identifier per-interface.
// KeyId returns a uint32
func (obj *ospfv2AuthenticationMd5) KeyId() uint32 {

	return *obj.obj.KeyId

}

// The unique MD5 Key Identifier per-interface.
// KeyId returns a uint32
func (obj *ospfv2AuthenticationMd5) HasKeyId() bool {
	return obj.obj.KeyId != nil
}

// The unique MD5 Key Identifier per-interface.
// SetKeyId sets the uint32 value in the Ospfv2AuthenticationMd5 object
func (obj *ospfv2AuthenticationMd5) SetKeyId(value uint32) Ospfv2AuthenticationMd5 {

	obj.obj.KeyId = &value
	return obj
}

// An alphanumeric secret used to generate the 16 byte MD5 hash value added
// to the OSPFv2 PDU in the Authentication TLV.
// Key returns a string
func (obj *ospfv2AuthenticationMd5) Key() string {

	return *obj.obj.Key

}

// An alphanumeric secret used to generate the 16 byte MD5 hash value added
// to the OSPFv2 PDU in the Authentication TLV.
// Key returns a string
func (obj *ospfv2AuthenticationMd5) HasKey() bool {
	return obj.obj.Key != nil
}

// An alphanumeric secret used to generate the 16 byte MD5 hash value added
// to the OSPFv2 PDU in the Authentication TLV.
// SetKey sets the string value in the Ospfv2AuthenticationMd5 object
func (obj *ospfv2AuthenticationMd5) SetKey(value string) Ospfv2AuthenticationMd5 {

	obj.obj.Key = &value
	return obj
}

func (obj *ospfv2AuthenticationMd5) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.KeyId != nil {

		if *obj.obj.KeyId < 1 || *obj.obj.KeyId > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= Ospfv2AuthenticationMd5.KeyId <= 255 but Got %d", *obj.obj.KeyId))
		}

	}

	if obj.obj.Key != nil {

		if len(*obj.obj.Key) < 1 || len(*obj.obj.Key) > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of Ospfv2AuthenticationMd5.Key <= 16 but Got %d",
					len(*obj.obj.Key)))
		}

	}

}

func (obj *ospfv2AuthenticationMd5) setDefault() {

}

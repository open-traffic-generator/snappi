package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic *****
type macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic struct {
	validation
	obj          *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic
	marshaller   marshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic
	unMarshaller unMarshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic
}

func NewMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic() MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic {
	obj := macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic{obj: &otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) msg() *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic {
	return obj.obj
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) setMsg(msg *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic struct {
	obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic
}

type marshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic interface {
	// ToProto marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic to protobuf object *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic
	ToProto() (*otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic, error)
	// ToPbText marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic struct {
	obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic
}

type unMarshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic interface {
	// FromProto unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic from protobuf object *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic
	FromProto(msg *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) (MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic, error)
	// FromPbText unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) Marshal() marshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) Unmarshal() unMarshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) ToProto() (*otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) FromProto(msg *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) (MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) FromPbText(value string) error {
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

func (m *marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) FromYaml(value string) error {
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

func (m *marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) FromJson(value string) error {
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

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) Clone() (MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic()
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

// MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic is encryption only traffic settings.
type MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic interface {
	Validation
	// msg marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic to protobuf object *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic
	// and doesn't set defaults
	msg() *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic
	// setMsg unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic from protobuf object *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic
	// provides marshal interface
	Marshal() marshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic
	// validate validates MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SendGratarp returns bool, set in MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic.
	SendGratarp() bool
	// SetSendGratarp assigns bool provided by user to MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic
	SetSendGratarp(value bool) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic
	// HasSendGratarp checks if SendGratarp has been set in MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic
	HasSendGratarp() bool
}

// Send gratuitous ARP or not.
// SendGratarp returns a bool
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) SendGratarp() bool {

	return *obj.obj.SendGratarp

}

// Send gratuitous ARP or not.
// SendGratarp returns a bool
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) HasSendGratarp() bool {
	return obj.obj.SendGratarp != nil
}

// Send gratuitous ARP or not.
// SetSendGratarp sets the bool value in the MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic object
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) SetSendGratarp(value bool) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic {

	obj.obj.SendGratarp = &value
	return obj
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) setDefault() {
	if obj.obj.SendGratarp == nil {
		obj.SetSendGratarp(true)
	}

}

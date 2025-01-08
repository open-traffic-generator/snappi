package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic *****
type macsecCryptoEngineTypeStatelessEncryptionOnlyTraffic struct {
	validation
	obj          *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic
	marshaller   marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic
	unMarshaller unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic
}

func NewMacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic() MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic {
	obj := macsecCryptoEngineTypeStatelessEncryptionOnlyTraffic{obj: &otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) msg() *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic {
	return obj.obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) setMsg(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic struct {
	obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTraffic
}

type marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic interface {
	// ToProto marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic to protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic
	ToProto() (*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic, error)
	// ToPbText marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic struct {
	obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTraffic
}

type unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic interface {
	// FromProto unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic from protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic
	FromProto(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) (MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic, error)
	// FromPbText unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) Marshal() marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) Unmarshal() unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) ToProto() (*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) FromProto(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) (MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) FromJson(value string) error {
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

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) Clone() (MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic()
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

// MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic is encryption only traffic settings.
type MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic interface {
	Validation
	// msg marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic to protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic
	// setMsg unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic from protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic
	// validate validates MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SendGratarp returns bool, set in MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic.
	SendGratarp() bool
	// SetSendGratarp assigns bool provided by user to MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic
	SetSendGratarp(value bool) MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic
	// HasSendGratarp checks if SendGratarp has been set in MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic
	HasSendGratarp() bool
}

// Send gratuitous ARP or not.
// SendGratarp returns a bool
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) SendGratarp() bool {

	return *obj.obj.SendGratarp

}

// Send gratuitous ARP or not.
// SendGratarp returns a bool
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) HasSendGratarp() bool {
	return obj.obj.SendGratarp != nil
}

// Send gratuitous ARP or not.
// SetSendGratarp sets the bool value in the MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic object
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) SetSendGratarp(value bool) MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic {

	obj.obj.SendGratarp = &value
	return obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) setDefault() {
	if obj.obj.SendGratarp == nil {
		obj.SetSendGratarp(true)
	}

}

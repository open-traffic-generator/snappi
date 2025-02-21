package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineStatelessEncryptionOnlyTraffic *****
type macsecCryptoEngineStatelessEncryptionOnlyTraffic struct {
	validation
	obj          *otg.MacsecCryptoEngineStatelessEncryptionOnlyTraffic
	marshaller   marshalMacsecCryptoEngineStatelessEncryptionOnlyTraffic
	unMarshaller unMarshalMacsecCryptoEngineStatelessEncryptionOnlyTraffic
}

func NewMacsecCryptoEngineStatelessEncryptionOnlyTraffic() MacsecCryptoEngineStatelessEncryptionOnlyTraffic {
	obj := macsecCryptoEngineStatelessEncryptionOnlyTraffic{obj: &otg.MacsecCryptoEngineStatelessEncryptionOnlyTraffic{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTraffic) msg() *otg.MacsecCryptoEngineStatelessEncryptionOnlyTraffic {
	return obj.obj
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTraffic) setMsg(msg *otg.MacsecCryptoEngineStatelessEncryptionOnlyTraffic) MacsecCryptoEngineStatelessEncryptionOnlyTraffic {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineStatelessEncryptionOnlyTraffic struct {
	obj *macsecCryptoEngineStatelessEncryptionOnlyTraffic
}

type marshalMacsecCryptoEngineStatelessEncryptionOnlyTraffic interface {
	// ToProto marshals MacsecCryptoEngineStatelessEncryptionOnlyTraffic to protobuf object *otg.MacsecCryptoEngineStatelessEncryptionOnlyTraffic
	ToProto() (*otg.MacsecCryptoEngineStatelessEncryptionOnlyTraffic, error)
	// ToPbText marshals MacsecCryptoEngineStatelessEncryptionOnlyTraffic to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineStatelessEncryptionOnlyTraffic to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineStatelessEncryptionOnlyTraffic to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineStatelessEncryptionOnlyTraffic struct {
	obj *macsecCryptoEngineStatelessEncryptionOnlyTraffic
}

type unMarshalMacsecCryptoEngineStatelessEncryptionOnlyTraffic interface {
	// FromProto unmarshals MacsecCryptoEngineStatelessEncryptionOnlyTraffic from protobuf object *otg.MacsecCryptoEngineStatelessEncryptionOnlyTraffic
	FromProto(msg *otg.MacsecCryptoEngineStatelessEncryptionOnlyTraffic) (MacsecCryptoEngineStatelessEncryptionOnlyTraffic, error)
	// FromPbText unmarshals MacsecCryptoEngineStatelessEncryptionOnlyTraffic from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineStatelessEncryptionOnlyTraffic from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineStatelessEncryptionOnlyTraffic from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTraffic) Marshal() marshalMacsecCryptoEngineStatelessEncryptionOnlyTraffic {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineStatelessEncryptionOnlyTraffic{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTraffic) Unmarshal() unMarshalMacsecCryptoEngineStatelessEncryptionOnlyTraffic {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineStatelessEncryptionOnlyTraffic{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineStatelessEncryptionOnlyTraffic) ToProto() (*otg.MacsecCryptoEngineStatelessEncryptionOnlyTraffic, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineStatelessEncryptionOnlyTraffic) FromProto(msg *otg.MacsecCryptoEngineStatelessEncryptionOnlyTraffic) (MacsecCryptoEngineStatelessEncryptionOnlyTraffic, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineStatelessEncryptionOnlyTraffic) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatelessEncryptionOnlyTraffic) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineStatelessEncryptionOnlyTraffic) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatelessEncryptionOnlyTraffic) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineStatelessEncryptionOnlyTraffic) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatelessEncryptionOnlyTraffic) FromJson(value string) error {
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

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTraffic) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTraffic) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTraffic) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTraffic) Clone() (MacsecCryptoEngineStatelessEncryptionOnlyTraffic, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineStatelessEncryptionOnlyTraffic()
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

// MacsecCryptoEngineStatelessEncryptionOnlyTraffic is encryption only traffic configuration.
type MacsecCryptoEngineStatelessEncryptionOnlyTraffic interface {
	Validation
	// msg marshals MacsecCryptoEngineStatelessEncryptionOnlyTraffic to protobuf object *otg.MacsecCryptoEngineStatelessEncryptionOnlyTraffic
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineStatelessEncryptionOnlyTraffic
	// setMsg unmarshals MacsecCryptoEngineStatelessEncryptionOnlyTraffic from protobuf object *otg.MacsecCryptoEngineStatelessEncryptionOnlyTraffic
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineStatelessEncryptionOnlyTraffic) MacsecCryptoEngineStatelessEncryptionOnlyTraffic
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineStatelessEncryptionOnlyTraffic
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineStatelessEncryptionOnlyTraffic
	// validate validates MacsecCryptoEngineStatelessEncryptionOnlyTraffic
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineStatelessEncryptionOnlyTraffic, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SendGratarp returns bool, set in MacsecCryptoEngineStatelessEncryptionOnlyTraffic.
	SendGratarp() bool
	// SetSendGratarp assigns bool provided by user to MacsecCryptoEngineStatelessEncryptionOnlyTraffic
	SetSendGratarp(value bool) MacsecCryptoEngineStatelessEncryptionOnlyTraffic
	// HasSendGratarp checks if SendGratarp has been set in MacsecCryptoEngineStatelessEncryptionOnlyTraffic
	HasSendGratarp() bool
}

// Send gratuitous ARP or not.
// SendGratarp returns a bool
func (obj *macsecCryptoEngineStatelessEncryptionOnlyTraffic) SendGratarp() bool {

	return *obj.obj.SendGratarp

}

// Send gratuitous ARP or not.
// SendGratarp returns a bool
func (obj *macsecCryptoEngineStatelessEncryptionOnlyTraffic) HasSendGratarp() bool {
	return obj.obj.SendGratarp != nil
}

// Send gratuitous ARP or not.
// SetSendGratarp sets the bool value in the MacsecCryptoEngineStatelessEncryptionOnlyTraffic object
func (obj *macsecCryptoEngineStatelessEncryptionOnlyTraffic) SetSendGratarp(value bool) MacsecCryptoEngineStatelessEncryptionOnlyTraffic {

	obj.obj.SendGratarp = &value
	return obj
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTraffic) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyTraffic) setDefault() {
	if obj.obj.SendGratarp == nil {
		obj.SetSendGratarp(true)
	}

}

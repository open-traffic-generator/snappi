package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MkaBasicSupportedCipherSuites *****
type mkaBasicSupportedCipherSuites struct {
	validation
	obj          *otg.MkaBasicSupportedCipherSuites
	marshaller   marshalMkaBasicSupportedCipherSuites
	unMarshaller unMarshalMkaBasicSupportedCipherSuites
}

func NewMkaBasicSupportedCipherSuites() MkaBasicSupportedCipherSuites {
	obj := mkaBasicSupportedCipherSuites{obj: &otg.MkaBasicSupportedCipherSuites{}}
	obj.setDefault()
	return &obj
}

func (obj *mkaBasicSupportedCipherSuites) msg() *otg.MkaBasicSupportedCipherSuites {
	return obj.obj
}

func (obj *mkaBasicSupportedCipherSuites) setMsg(msg *otg.MkaBasicSupportedCipherSuites) MkaBasicSupportedCipherSuites {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmkaBasicSupportedCipherSuites struct {
	obj *mkaBasicSupportedCipherSuites
}

type marshalMkaBasicSupportedCipherSuites interface {
	// ToProto marshals MkaBasicSupportedCipherSuites to protobuf object *otg.MkaBasicSupportedCipherSuites
	ToProto() (*otg.MkaBasicSupportedCipherSuites, error)
	// ToPbText marshals MkaBasicSupportedCipherSuites to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MkaBasicSupportedCipherSuites to YAML text
	ToYaml() (string, error)
	// ToJson marshals MkaBasicSupportedCipherSuites to JSON text
	ToJson() (string, error)
}

type unMarshalmkaBasicSupportedCipherSuites struct {
	obj *mkaBasicSupportedCipherSuites
}

type unMarshalMkaBasicSupportedCipherSuites interface {
	// FromProto unmarshals MkaBasicSupportedCipherSuites from protobuf object *otg.MkaBasicSupportedCipherSuites
	FromProto(msg *otg.MkaBasicSupportedCipherSuites) (MkaBasicSupportedCipherSuites, error)
	// FromPbText unmarshals MkaBasicSupportedCipherSuites from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MkaBasicSupportedCipherSuites from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MkaBasicSupportedCipherSuites from JSON text
	FromJson(value string) error
}

func (obj *mkaBasicSupportedCipherSuites) Marshal() marshalMkaBasicSupportedCipherSuites {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmkaBasicSupportedCipherSuites{obj: obj}
	}
	return obj.marshaller
}

func (obj *mkaBasicSupportedCipherSuites) Unmarshal() unMarshalMkaBasicSupportedCipherSuites {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmkaBasicSupportedCipherSuites{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmkaBasicSupportedCipherSuites) ToProto() (*otg.MkaBasicSupportedCipherSuites, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmkaBasicSupportedCipherSuites) FromProto(msg *otg.MkaBasicSupportedCipherSuites) (MkaBasicSupportedCipherSuites, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmkaBasicSupportedCipherSuites) ToPbText() (string, error) {
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

func (m *unMarshalmkaBasicSupportedCipherSuites) FromPbText(value string) error {
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

func (m *marshalmkaBasicSupportedCipherSuites) ToYaml() (string, error) {
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

func (m *unMarshalmkaBasicSupportedCipherSuites) FromYaml(value string) error {
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

func (m *marshalmkaBasicSupportedCipherSuites) ToJson() (string, error) {
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

func (m *unMarshalmkaBasicSupportedCipherSuites) FromJson(value string) error {
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

func (obj *mkaBasicSupportedCipherSuites) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mkaBasicSupportedCipherSuites) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mkaBasicSupportedCipherSuites) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mkaBasicSupportedCipherSuites) Clone() (MkaBasicSupportedCipherSuites, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMkaBasicSupportedCipherSuites()
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

// MkaBasicSupportedCipherSuites is the container for supported cipher suites.
type MkaBasicSupportedCipherSuites interface {
	Validation
	// msg marshals MkaBasicSupportedCipherSuites to protobuf object *otg.MkaBasicSupportedCipherSuites
	// and doesn't set defaults
	msg() *otg.MkaBasicSupportedCipherSuites
	// setMsg unmarshals MkaBasicSupportedCipherSuites from protobuf object *otg.MkaBasicSupportedCipherSuites
	// and doesn't set defaults
	setMsg(*otg.MkaBasicSupportedCipherSuites) MkaBasicSupportedCipherSuites
	// provides marshal interface
	Marshal() marshalMkaBasicSupportedCipherSuites
	// provides unmarshal interface
	Unmarshal() unMarshalMkaBasicSupportedCipherSuites
	// validate validates MkaBasicSupportedCipherSuites
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MkaBasicSupportedCipherSuites, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// GcmAes128 returns bool, set in MkaBasicSupportedCipherSuites.
	GcmAes128() bool
	// SetGcmAes128 assigns bool provided by user to MkaBasicSupportedCipherSuites
	SetGcmAes128(value bool) MkaBasicSupportedCipherSuites
	// HasGcmAes128 checks if GcmAes128 has been set in MkaBasicSupportedCipherSuites
	HasGcmAes128() bool
	// GcmAes256 returns bool, set in MkaBasicSupportedCipherSuites.
	GcmAes256() bool
	// SetGcmAes256 assigns bool provided by user to MkaBasicSupportedCipherSuites
	SetGcmAes256(value bool) MkaBasicSupportedCipherSuites
	// HasGcmAes256 checks if GcmAes256 has been set in MkaBasicSupportedCipherSuites
	HasGcmAes256() bool
	// GcmAesXpn128 returns bool, set in MkaBasicSupportedCipherSuites.
	GcmAesXpn128() bool
	// SetGcmAesXpn128 assigns bool provided by user to MkaBasicSupportedCipherSuites
	SetGcmAesXpn128(value bool) MkaBasicSupportedCipherSuites
	// HasGcmAesXpn128 checks if GcmAesXpn128 has been set in MkaBasicSupportedCipherSuites
	HasGcmAesXpn128() bool
	// GcmAesXpn256 returns bool, set in MkaBasicSupportedCipherSuites.
	GcmAesXpn256() bool
	// SetGcmAesXpn256 assigns bool provided by user to MkaBasicSupportedCipherSuites
	SetGcmAesXpn256(value bool) MkaBasicSupportedCipherSuites
	// HasGcmAesXpn256 checks if GcmAesXpn256 has been set in MkaBasicSupportedCipherSuites
	HasGcmAesXpn256() bool
}

// GCM-AES-128.
// GcmAes128 returns a bool
func (obj *mkaBasicSupportedCipherSuites) GcmAes128() bool {

	return *obj.obj.GcmAes_128

}

// GCM-AES-128.
// GcmAes128 returns a bool
func (obj *mkaBasicSupportedCipherSuites) HasGcmAes128() bool {
	return obj.obj.GcmAes_128 != nil
}

// GCM-AES-128.
// SetGcmAes128 sets the bool value in the MkaBasicSupportedCipherSuites object
func (obj *mkaBasicSupportedCipherSuites) SetGcmAes128(value bool) MkaBasicSupportedCipherSuites {

	obj.obj.GcmAes_128 = &value
	return obj
}

// GCM-AES-256.
// GcmAes256 returns a bool
func (obj *mkaBasicSupportedCipherSuites) GcmAes256() bool {

	return *obj.obj.GcmAes_256

}

// GCM-AES-256.
// GcmAes256 returns a bool
func (obj *mkaBasicSupportedCipherSuites) HasGcmAes256() bool {
	return obj.obj.GcmAes_256 != nil
}

// GCM-AES-256.
// SetGcmAes256 sets the bool value in the MkaBasicSupportedCipherSuites object
func (obj *mkaBasicSupportedCipherSuites) SetGcmAes256(value bool) MkaBasicSupportedCipherSuites {

	obj.obj.GcmAes_256 = &value
	return obj
}

// GCM-AES-XPN-128.
// GcmAesXpn128 returns a bool
func (obj *mkaBasicSupportedCipherSuites) GcmAesXpn128() bool {

	return *obj.obj.GcmAesXpn_128

}

// GCM-AES-XPN-128.
// GcmAesXpn128 returns a bool
func (obj *mkaBasicSupportedCipherSuites) HasGcmAesXpn128() bool {
	return obj.obj.GcmAesXpn_128 != nil
}

// GCM-AES-XPN-128.
// SetGcmAesXpn128 sets the bool value in the MkaBasicSupportedCipherSuites object
func (obj *mkaBasicSupportedCipherSuites) SetGcmAesXpn128(value bool) MkaBasicSupportedCipherSuites {

	obj.obj.GcmAesXpn_128 = &value
	return obj
}

// GCM-AES-XPN-256.
// GcmAesXpn256 returns a bool
func (obj *mkaBasicSupportedCipherSuites) GcmAesXpn256() bool {

	return *obj.obj.GcmAesXpn_256

}

// GCM-AES-XPN-256.
// GcmAesXpn256 returns a bool
func (obj *mkaBasicSupportedCipherSuites) HasGcmAesXpn256() bool {
	return obj.obj.GcmAesXpn_256 != nil
}

// GCM-AES-XPN-256.
// SetGcmAesXpn256 sets the bool value in the MkaBasicSupportedCipherSuites object
func (obj *mkaBasicSupportedCipherSuites) SetGcmAesXpn256(value bool) MkaBasicSupportedCipherSuites {

	obj.obj.GcmAesXpn_256 = &value
	return obj
}

func (obj *mkaBasicSupportedCipherSuites) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *mkaBasicSupportedCipherSuites) setDefault() {
	if obj.obj.GcmAes_128 == nil {
		obj.SetGcmAes128(true)
	}
	if obj.obj.GcmAes_256 == nil {
		obj.SetGcmAes256(true)
	}
	if obj.obj.GcmAesXpn_128 == nil {
		obj.SetGcmAesXpn128(true)
	}
	if obj.obj.GcmAesXpn_256 == nil {
		obj.SetGcmAesXpn256(true)
	}

}

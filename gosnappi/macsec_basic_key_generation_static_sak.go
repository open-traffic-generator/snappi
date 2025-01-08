package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecBasicKeyGenerationStaticSak *****
type macsecBasicKeyGenerationStaticSak struct {
	validation
	obj          *otg.MacsecBasicKeyGenerationStaticSak
	marshaller   marshalMacsecBasicKeyGenerationStaticSak
	unMarshaller unMarshalMacsecBasicKeyGenerationStaticSak
}

func NewMacsecBasicKeyGenerationStaticSak() MacsecBasicKeyGenerationStaticSak {
	obj := macsecBasicKeyGenerationStaticSak{obj: &otg.MacsecBasicKeyGenerationStaticSak{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecBasicKeyGenerationStaticSak) msg() *otg.MacsecBasicKeyGenerationStaticSak {
	return obj.obj
}

func (obj *macsecBasicKeyGenerationStaticSak) setMsg(msg *otg.MacsecBasicKeyGenerationStaticSak) MacsecBasicKeyGenerationStaticSak {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecBasicKeyGenerationStaticSak struct {
	obj *macsecBasicKeyGenerationStaticSak
}

type marshalMacsecBasicKeyGenerationStaticSak interface {
	// ToProto marshals MacsecBasicKeyGenerationStaticSak to protobuf object *otg.MacsecBasicKeyGenerationStaticSak
	ToProto() (*otg.MacsecBasicKeyGenerationStaticSak, error)
	// ToPbText marshals MacsecBasicKeyGenerationStaticSak to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecBasicKeyGenerationStaticSak to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecBasicKeyGenerationStaticSak to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecBasicKeyGenerationStaticSak struct {
	obj *macsecBasicKeyGenerationStaticSak
}

type unMarshalMacsecBasicKeyGenerationStaticSak interface {
	// FromProto unmarshals MacsecBasicKeyGenerationStaticSak from protobuf object *otg.MacsecBasicKeyGenerationStaticSak
	FromProto(msg *otg.MacsecBasicKeyGenerationStaticSak) (MacsecBasicKeyGenerationStaticSak, error)
	// FromPbText unmarshals MacsecBasicKeyGenerationStaticSak from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecBasicKeyGenerationStaticSak from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecBasicKeyGenerationStaticSak from JSON text
	FromJson(value string) error
}

func (obj *macsecBasicKeyGenerationStaticSak) Marshal() marshalMacsecBasicKeyGenerationStaticSak {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecBasicKeyGenerationStaticSak{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecBasicKeyGenerationStaticSak) Unmarshal() unMarshalMacsecBasicKeyGenerationStaticSak {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecBasicKeyGenerationStaticSak{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecBasicKeyGenerationStaticSak) ToProto() (*otg.MacsecBasicKeyGenerationStaticSak, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecBasicKeyGenerationStaticSak) FromProto(msg *otg.MacsecBasicKeyGenerationStaticSak) (MacsecBasicKeyGenerationStaticSak, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecBasicKeyGenerationStaticSak) ToPbText() (string, error) {
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

func (m *unMarshalmacsecBasicKeyGenerationStaticSak) FromPbText(value string) error {
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

func (m *marshalmacsecBasicKeyGenerationStaticSak) ToYaml() (string, error) {
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

func (m *unMarshalmacsecBasicKeyGenerationStaticSak) FromYaml(value string) error {
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

func (m *marshalmacsecBasicKeyGenerationStaticSak) ToJson() (string, error) {
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

func (m *unMarshalmacsecBasicKeyGenerationStaticSak) FromJson(value string) error {
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

func (obj *macsecBasicKeyGenerationStaticSak) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecBasicKeyGenerationStaticSak) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecBasicKeyGenerationStaticSak) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecBasicKeyGenerationStaticSak) Clone() (MacsecBasicKeyGenerationStaticSak, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecBasicKeyGenerationStaticSak()
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

// MacsecBasicKeyGenerationStaticSak is the container for SAK.
type MacsecBasicKeyGenerationStaticSak interface {
	Validation
	// msg marshals MacsecBasicKeyGenerationStaticSak to protobuf object *otg.MacsecBasicKeyGenerationStaticSak
	// and doesn't set defaults
	msg() *otg.MacsecBasicKeyGenerationStaticSak
	// setMsg unmarshals MacsecBasicKeyGenerationStaticSak from protobuf object *otg.MacsecBasicKeyGenerationStaticSak
	// and doesn't set defaults
	setMsg(*otg.MacsecBasicKeyGenerationStaticSak) MacsecBasicKeyGenerationStaticSak
	// provides marshal interface
	Marshal() marshalMacsecBasicKeyGenerationStaticSak
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecBasicKeyGenerationStaticSak
	// validate validates MacsecBasicKeyGenerationStaticSak
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecBasicKeyGenerationStaticSak, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in MacsecBasicKeyGenerationStaticSak.
	Name() string
	// SetName assigns string provided by user to MacsecBasicKeyGenerationStaticSak
	SetName(value string) MacsecBasicKeyGenerationStaticSak
	// HasName checks if Name has been set in MacsecBasicKeyGenerationStaticSak
	HasName() bool
	// Sak returns string, set in MacsecBasicKeyGenerationStaticSak.
	Sak() string
	// SetSak assigns string provided by user to MacsecBasicKeyGenerationStaticSak
	SetSak(value string) MacsecBasicKeyGenerationStaticSak
	// HasSak checks if Sak has been set in MacsecBasicKeyGenerationStaticSak
	HasSak() bool
}

// SAK name.
// Name returns a string
func (obj *macsecBasicKeyGenerationStaticSak) Name() string {

	return *obj.obj.Name

}

// SAK name.
// Name returns a string
func (obj *macsecBasicKeyGenerationStaticSak) HasName() bool {
	return obj.obj.Name != nil
}

// SAK name.
// SetName sets the string value in the MacsecBasicKeyGenerationStaticSak object
func (obj *macsecBasicKeyGenerationStaticSak) SetName(value string) MacsecBasicKeyGenerationStaticSak {

	obj.obj.Name = &value
	return obj
}

// SAK bits as hex string. Either 128 bits or 256 bits depending on the citpher suite chosen.
// Sak returns a string
func (obj *macsecBasicKeyGenerationStaticSak) Sak() string {

	return *obj.obj.Sak

}

// SAK bits as hex string. Either 128 bits or 256 bits depending on the citpher suite chosen.
// Sak returns a string
func (obj *macsecBasicKeyGenerationStaticSak) HasSak() bool {
	return obj.obj.Sak != nil
}

// SAK bits as hex string. Either 128 bits or 256 bits depending on the citpher suite chosen.
// SetSak sets the string value in the MacsecBasicKeyGenerationStaticSak object
func (obj *macsecBasicKeyGenerationStaticSak) SetSak(value string) MacsecBasicKeyGenerationStaticSak {

	obj.obj.Sak = &value
	return obj
}

func (obj *macsecBasicKeyGenerationStaticSak) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecBasicKeyGenerationStaticSak) setDefault() {

}

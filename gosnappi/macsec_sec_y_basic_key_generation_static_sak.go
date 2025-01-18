package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYBasicKeyGenerationStaticSak *****
type macsecSecYBasicKeyGenerationStaticSak struct {
	validation
	obj          *otg.MacsecSecYBasicKeyGenerationStaticSak
	marshaller   marshalMacsecSecYBasicKeyGenerationStaticSak
	unMarshaller unMarshalMacsecSecYBasicKeyGenerationStaticSak
}

func NewMacsecSecYBasicKeyGenerationStaticSak() MacsecSecYBasicKeyGenerationStaticSak {
	obj := macsecSecYBasicKeyGenerationStaticSak{obj: &otg.MacsecSecYBasicKeyGenerationStaticSak{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYBasicKeyGenerationStaticSak) msg() *otg.MacsecSecYBasicKeyGenerationStaticSak {
	return obj.obj
}

func (obj *macsecSecYBasicKeyGenerationStaticSak) setMsg(msg *otg.MacsecSecYBasicKeyGenerationStaticSak) MacsecSecYBasicKeyGenerationStaticSak {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYBasicKeyGenerationStaticSak struct {
	obj *macsecSecYBasicKeyGenerationStaticSak
}

type marshalMacsecSecYBasicKeyGenerationStaticSak interface {
	// ToProto marshals MacsecSecYBasicKeyGenerationStaticSak to protobuf object *otg.MacsecSecYBasicKeyGenerationStaticSak
	ToProto() (*otg.MacsecSecYBasicKeyGenerationStaticSak, error)
	// ToPbText marshals MacsecSecYBasicKeyGenerationStaticSak to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYBasicKeyGenerationStaticSak to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYBasicKeyGenerationStaticSak to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYBasicKeyGenerationStaticSak struct {
	obj *macsecSecYBasicKeyGenerationStaticSak
}

type unMarshalMacsecSecYBasicKeyGenerationStaticSak interface {
	// FromProto unmarshals MacsecSecYBasicKeyGenerationStaticSak from protobuf object *otg.MacsecSecYBasicKeyGenerationStaticSak
	FromProto(msg *otg.MacsecSecYBasicKeyGenerationStaticSak) (MacsecSecYBasicKeyGenerationStaticSak, error)
	// FromPbText unmarshals MacsecSecYBasicKeyGenerationStaticSak from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYBasicKeyGenerationStaticSak from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYBasicKeyGenerationStaticSak from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYBasicKeyGenerationStaticSak) Marshal() marshalMacsecSecYBasicKeyGenerationStaticSak {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYBasicKeyGenerationStaticSak{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYBasicKeyGenerationStaticSak) Unmarshal() unMarshalMacsecSecYBasicKeyGenerationStaticSak {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYBasicKeyGenerationStaticSak{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYBasicKeyGenerationStaticSak) ToProto() (*otg.MacsecSecYBasicKeyGenerationStaticSak, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYBasicKeyGenerationStaticSak) FromProto(msg *otg.MacsecSecYBasicKeyGenerationStaticSak) (MacsecSecYBasicKeyGenerationStaticSak, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYBasicKeyGenerationStaticSak) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYBasicKeyGenerationStaticSak) FromPbText(value string) error {
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

func (m *marshalmacsecSecYBasicKeyGenerationStaticSak) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYBasicKeyGenerationStaticSak) FromYaml(value string) error {
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

func (m *marshalmacsecSecYBasicKeyGenerationStaticSak) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYBasicKeyGenerationStaticSak) FromJson(value string) error {
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

func (obj *macsecSecYBasicKeyGenerationStaticSak) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYBasicKeyGenerationStaticSak) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYBasicKeyGenerationStaticSak) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYBasicKeyGenerationStaticSak) Clone() (MacsecSecYBasicKeyGenerationStaticSak, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYBasicKeyGenerationStaticSak()
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

// MacsecSecYBasicKeyGenerationStaticSak is the container for SAK.
type MacsecSecYBasicKeyGenerationStaticSak interface {
	Validation
	// msg marshals MacsecSecYBasicKeyGenerationStaticSak to protobuf object *otg.MacsecSecYBasicKeyGenerationStaticSak
	// and doesn't set defaults
	msg() *otg.MacsecSecYBasicKeyGenerationStaticSak
	// setMsg unmarshals MacsecSecYBasicKeyGenerationStaticSak from protobuf object *otg.MacsecSecYBasicKeyGenerationStaticSak
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYBasicKeyGenerationStaticSak) MacsecSecYBasicKeyGenerationStaticSak
	// provides marshal interface
	Marshal() marshalMacsecSecYBasicKeyGenerationStaticSak
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYBasicKeyGenerationStaticSak
	// validate validates MacsecSecYBasicKeyGenerationStaticSak
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYBasicKeyGenerationStaticSak, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in MacsecSecYBasicKeyGenerationStaticSak.
	Name() string
	// SetName assigns string provided by user to MacsecSecYBasicKeyGenerationStaticSak
	SetName(value string) MacsecSecYBasicKeyGenerationStaticSak
	// HasName checks if Name has been set in MacsecSecYBasicKeyGenerationStaticSak
	HasName() bool
	// Sak returns string, set in MacsecSecYBasicKeyGenerationStaticSak.
	Sak() string
	// SetSak assigns string provided by user to MacsecSecYBasicKeyGenerationStaticSak
	SetSak(value string) MacsecSecYBasicKeyGenerationStaticSak
	// HasSak checks if Sak has been set in MacsecSecYBasicKeyGenerationStaticSak
	HasSak() bool
}

// SAK name.
// Name returns a string
func (obj *macsecSecYBasicKeyGenerationStaticSak) Name() string {

	return *obj.obj.Name

}

// SAK name.
// Name returns a string
func (obj *macsecSecYBasicKeyGenerationStaticSak) HasName() bool {
	return obj.obj.Name != nil
}

// SAK name.
// SetName sets the string value in the MacsecSecYBasicKeyGenerationStaticSak object
func (obj *macsecSecYBasicKeyGenerationStaticSak) SetName(value string) MacsecSecYBasicKeyGenerationStaticSak {

	obj.obj.Name = &value
	return obj
}

// SAK bits as hex string. Either 128 bits or 256 bits depending on the citpher suite chosen.
// Sak returns a string
func (obj *macsecSecYBasicKeyGenerationStaticSak) Sak() string {

	return *obj.obj.Sak

}

// SAK bits as hex string. Either 128 bits or 256 bits depending on the citpher suite chosen.
// Sak returns a string
func (obj *macsecSecYBasicKeyGenerationStaticSak) HasSak() bool {
	return obj.obj.Sak != nil
}

// SAK bits as hex string. Either 128 bits or 256 bits depending on the citpher suite chosen.
// SetSak sets the string value in the MacsecSecYBasicKeyGenerationStaticSak object
func (obj *macsecSecYBasicKeyGenerationStaticSak) SetSak(value string) MacsecSecYBasicKeyGenerationStaticSak {

	obj.obj.Sak = &value
	return obj
}

func (obj *macsecSecYBasicKeyGenerationStaticSak) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecSecYBasicKeyGenerationStaticSak) setDefault() {

}

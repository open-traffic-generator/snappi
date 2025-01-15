package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecAdvance *****
type macsecAdvance struct {
	validation
	obj             *otg.MacsecAdvance
	marshaller      marshalMacsecAdvance
	unMarshaller    unMarshalMacsecAdvance
	staticKeyHolder MacsecAdvanceStaticKey
}

func NewMacsecAdvance() MacsecAdvance {
	obj := macsecAdvance{obj: &otg.MacsecAdvance{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecAdvance) msg() *otg.MacsecAdvance {
	return obj.obj
}

func (obj *macsecAdvance) setMsg(msg *otg.MacsecAdvance) MacsecAdvance {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecAdvance struct {
	obj *macsecAdvance
}

type marshalMacsecAdvance interface {
	// ToProto marshals MacsecAdvance to protobuf object *otg.MacsecAdvance
	ToProto() (*otg.MacsecAdvance, error)
	// ToPbText marshals MacsecAdvance to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecAdvance to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecAdvance to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecAdvance struct {
	obj *macsecAdvance
}

type unMarshalMacsecAdvance interface {
	// FromProto unmarshals MacsecAdvance from protobuf object *otg.MacsecAdvance
	FromProto(msg *otg.MacsecAdvance) (MacsecAdvance, error)
	// FromPbText unmarshals MacsecAdvance from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecAdvance from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecAdvance from JSON text
	FromJson(value string) error
}

func (obj *macsecAdvance) Marshal() marshalMacsecAdvance {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecAdvance{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecAdvance) Unmarshal() unMarshalMacsecAdvance {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecAdvance{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecAdvance) ToProto() (*otg.MacsecAdvance, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecAdvance) FromProto(msg *otg.MacsecAdvance) (MacsecAdvance, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecAdvance) ToPbText() (string, error) {
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

func (m *unMarshalmacsecAdvance) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalmacsecAdvance) ToYaml() (string, error) {
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

func (m *unMarshalmacsecAdvance) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalmacsecAdvance) ToJson() (string, error) {
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

func (m *unMarshalmacsecAdvance) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *macsecAdvance) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecAdvance) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecAdvance) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecAdvance) Clone() (MacsecAdvance, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecAdvance()
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

func (obj *macsecAdvance) setNil() {
	obj.staticKeyHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecAdvance is a container of advance properties for a MACsec interface.
type MacsecAdvance interface {
	Validation
	// msg marshals MacsecAdvance to protobuf object *otg.MacsecAdvance
	// and doesn't set defaults
	msg() *otg.MacsecAdvance
	// setMsg unmarshals MacsecAdvance from protobuf object *otg.MacsecAdvance
	// and doesn't set defaults
	setMsg(*otg.MacsecAdvance) MacsecAdvance
	// provides marshal interface
	Marshal() marshalMacsecAdvance
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecAdvance
	// validate validates MacsecAdvance
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecAdvance, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StaticKey returns MacsecAdvanceStaticKey, set in MacsecAdvance.
	// MacsecAdvanceStaticKey is static key advance settings.
	StaticKey() MacsecAdvanceStaticKey
	// SetStaticKey assigns MacsecAdvanceStaticKey provided by user to MacsecAdvance.
	// MacsecAdvanceStaticKey is static key advance settings.
	SetStaticKey(value MacsecAdvanceStaticKey) MacsecAdvance
	// HasStaticKey checks if StaticKey has been set in MacsecAdvance
	HasStaticKey() bool
	setNil()
}

// description is TBD
// StaticKey returns a MacsecAdvanceStaticKey
func (obj *macsecAdvance) StaticKey() MacsecAdvanceStaticKey {
	if obj.obj.StaticKey == nil {
		obj.obj.StaticKey = NewMacsecAdvanceStaticKey().msg()
	}
	if obj.staticKeyHolder == nil {
		obj.staticKeyHolder = &macsecAdvanceStaticKey{obj: obj.obj.StaticKey}
	}
	return obj.staticKeyHolder
}

// description is TBD
// StaticKey returns a MacsecAdvanceStaticKey
func (obj *macsecAdvance) HasStaticKey() bool {
	return obj.obj.StaticKey != nil
}

// description is TBD
// SetStaticKey sets the MacsecAdvanceStaticKey value in the MacsecAdvance object
func (obj *macsecAdvance) SetStaticKey(value MacsecAdvanceStaticKey) MacsecAdvance {

	obj.staticKeyHolder = nil
	obj.obj.StaticKey = value.msg()

	return obj
}

func (obj *macsecAdvance) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.StaticKey != nil {

		obj.StaticKey().validateObj(vObj, set_default)
	}

}

func (obj *macsecAdvance) setDefault() {

}

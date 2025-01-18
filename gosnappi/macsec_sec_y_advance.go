package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYAdvance *****
type macsecSecYAdvance struct {
	validation
	obj             *otg.MacsecSecYAdvance
	marshaller      marshalMacsecSecYAdvance
	unMarshaller    unMarshalMacsecSecYAdvance
	staticKeyHolder MacsecSecYAdvanceStaticKey
}

func NewMacsecSecYAdvance() MacsecSecYAdvance {
	obj := macsecSecYAdvance{obj: &otg.MacsecSecYAdvance{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYAdvance) msg() *otg.MacsecSecYAdvance {
	return obj.obj
}

func (obj *macsecSecYAdvance) setMsg(msg *otg.MacsecSecYAdvance) MacsecSecYAdvance {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYAdvance struct {
	obj *macsecSecYAdvance
}

type marshalMacsecSecYAdvance interface {
	// ToProto marshals MacsecSecYAdvance to protobuf object *otg.MacsecSecYAdvance
	ToProto() (*otg.MacsecSecYAdvance, error)
	// ToPbText marshals MacsecSecYAdvance to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYAdvance to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYAdvance to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYAdvance struct {
	obj *macsecSecYAdvance
}

type unMarshalMacsecSecYAdvance interface {
	// FromProto unmarshals MacsecSecYAdvance from protobuf object *otg.MacsecSecYAdvance
	FromProto(msg *otg.MacsecSecYAdvance) (MacsecSecYAdvance, error)
	// FromPbText unmarshals MacsecSecYAdvance from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYAdvance from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYAdvance from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYAdvance) Marshal() marshalMacsecSecYAdvance {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYAdvance{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYAdvance) Unmarshal() unMarshalMacsecSecYAdvance {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYAdvance{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYAdvance) ToProto() (*otg.MacsecSecYAdvance, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYAdvance) FromProto(msg *otg.MacsecSecYAdvance) (MacsecSecYAdvance, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYAdvance) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYAdvance) FromPbText(value string) error {
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

func (m *marshalmacsecSecYAdvance) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYAdvance) FromYaml(value string) error {
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

func (m *marshalmacsecSecYAdvance) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYAdvance) FromJson(value string) error {
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

func (obj *macsecSecYAdvance) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYAdvance) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYAdvance) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYAdvance) Clone() (MacsecSecYAdvance, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYAdvance()
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

func (obj *macsecSecYAdvance) setNil() {
	obj.staticKeyHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecSecYAdvance is a container of advance properties SecY.
type MacsecSecYAdvance interface {
	Validation
	// msg marshals MacsecSecYAdvance to protobuf object *otg.MacsecSecYAdvance
	// and doesn't set defaults
	msg() *otg.MacsecSecYAdvance
	// setMsg unmarshals MacsecSecYAdvance from protobuf object *otg.MacsecSecYAdvance
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYAdvance) MacsecSecYAdvance
	// provides marshal interface
	Marshal() marshalMacsecSecYAdvance
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYAdvance
	// validate validates MacsecSecYAdvance
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYAdvance, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StaticKey returns MacsecSecYAdvanceStaticKey, set in MacsecSecYAdvance.
	// MacsecSecYAdvanceStaticKey is static key advance settings.
	StaticKey() MacsecSecYAdvanceStaticKey
	// SetStaticKey assigns MacsecSecYAdvanceStaticKey provided by user to MacsecSecYAdvance.
	// MacsecSecYAdvanceStaticKey is static key advance settings.
	SetStaticKey(value MacsecSecYAdvanceStaticKey) MacsecSecYAdvance
	// HasStaticKey checks if StaticKey has been set in MacsecSecYAdvance
	HasStaticKey() bool
	setNil()
}

// description is TBD
// StaticKey returns a MacsecSecYAdvanceStaticKey
func (obj *macsecSecYAdvance) StaticKey() MacsecSecYAdvanceStaticKey {
	if obj.obj.StaticKey == nil {
		obj.obj.StaticKey = NewMacsecSecYAdvanceStaticKey().msg()
	}
	if obj.staticKeyHolder == nil {
		obj.staticKeyHolder = &macsecSecYAdvanceStaticKey{obj: obj.obj.StaticKey}
	}
	return obj.staticKeyHolder
}

// description is TBD
// StaticKey returns a MacsecSecYAdvanceStaticKey
func (obj *macsecSecYAdvance) HasStaticKey() bool {
	return obj.obj.StaticKey != nil
}

// description is TBD
// SetStaticKey sets the MacsecSecYAdvanceStaticKey value in the MacsecSecYAdvance object
func (obj *macsecSecYAdvance) SetStaticKey(value MacsecSecYAdvanceStaticKey) MacsecSecYAdvance {

	obj.staticKeyHolder = nil
	obj.obj.StaticKey = value.msg()

	return obj
}

func (obj *macsecSecYAdvance) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.StaticKey != nil {

		obj.StaticKey().validateObj(vObj, set_default)
	}

}

func (obj *macsecSecYAdvance) setDefault() {

}

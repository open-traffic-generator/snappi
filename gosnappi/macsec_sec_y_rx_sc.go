package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYRxSc *****
type macsecSecYRxSc struct {
	validation
	obj             *otg.MacsecSecYRxSc
	marshaller      marshalMacsecSecYRxSc
	unMarshaller    unMarshalMacsecSecYRxSc
	staticKeyHolder MacsecSecYRxScStaticKey
}

func NewMacsecSecYRxSc() MacsecSecYRxSc {
	obj := macsecSecYRxSc{obj: &otg.MacsecSecYRxSc{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYRxSc) msg() *otg.MacsecSecYRxSc {
	return obj.obj
}

func (obj *macsecSecYRxSc) setMsg(msg *otg.MacsecSecYRxSc) MacsecSecYRxSc {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYRxSc struct {
	obj *macsecSecYRxSc
}

type marshalMacsecSecYRxSc interface {
	// ToProto marshals MacsecSecYRxSc to protobuf object *otg.MacsecSecYRxSc
	ToProto() (*otg.MacsecSecYRxSc, error)
	// ToPbText marshals MacsecSecYRxSc to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYRxSc to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYRxSc to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYRxSc struct {
	obj *macsecSecYRxSc
}

type unMarshalMacsecSecYRxSc interface {
	// FromProto unmarshals MacsecSecYRxSc from protobuf object *otg.MacsecSecYRxSc
	FromProto(msg *otg.MacsecSecYRxSc) (MacsecSecYRxSc, error)
	// FromPbText unmarshals MacsecSecYRxSc from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYRxSc from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYRxSc from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYRxSc) Marshal() marshalMacsecSecYRxSc {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYRxSc{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYRxSc) Unmarshal() unMarshalMacsecSecYRxSc {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYRxSc{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYRxSc) ToProto() (*otg.MacsecSecYRxSc, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYRxSc) FromProto(msg *otg.MacsecSecYRxSc) (MacsecSecYRxSc, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYRxSc) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYRxSc) FromPbText(value string) error {
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

func (m *marshalmacsecSecYRxSc) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYRxSc) FromYaml(value string) error {
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

func (m *marshalmacsecSecYRxSc) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYRxSc) FromJson(value string) error {
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

func (obj *macsecSecYRxSc) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYRxSc) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYRxSc) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYRxSc) Clone() (MacsecSecYRxSc, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYRxSc()
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

func (obj *macsecSecYRxSc) setNil() {
	obj.staticKeyHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecSecYRxSc is the container for Rx SC settings.
type MacsecSecYRxSc interface {
	Validation
	// msg marshals MacsecSecYRxSc to protobuf object *otg.MacsecSecYRxSc
	// and doesn't set defaults
	msg() *otg.MacsecSecYRxSc
	// setMsg unmarshals MacsecSecYRxSc from protobuf object *otg.MacsecSecYRxSc
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYRxSc) MacsecSecYRxSc
	// provides marshal interface
	Marshal() marshalMacsecSecYRxSc
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYRxSc
	// validate validates MacsecSecYRxSc
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYRxSc, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StaticKey returns MacsecSecYRxScStaticKey, set in MacsecSecYRxSc.
	// MacsecSecYRxScStaticKey is rx SC settings for static key.
	StaticKey() MacsecSecYRxScStaticKey
	// SetStaticKey assigns MacsecSecYRxScStaticKey provided by user to MacsecSecYRxSc.
	// MacsecSecYRxScStaticKey is rx SC settings for static key.
	SetStaticKey(value MacsecSecYRxScStaticKey) MacsecSecYRxSc
	// HasStaticKey checks if StaticKey has been set in MacsecSecYRxSc
	HasStaticKey() bool
	setNil()
}

// description is TBD
// StaticKey returns a MacsecSecYRxScStaticKey
func (obj *macsecSecYRxSc) StaticKey() MacsecSecYRxScStaticKey {
	if obj.obj.StaticKey == nil {
		obj.obj.StaticKey = NewMacsecSecYRxScStaticKey().msg()
	}
	if obj.staticKeyHolder == nil {
		obj.staticKeyHolder = &macsecSecYRxScStaticKey{obj: obj.obj.StaticKey}
	}
	return obj.staticKeyHolder
}

// description is TBD
// StaticKey returns a MacsecSecYRxScStaticKey
func (obj *macsecSecYRxSc) HasStaticKey() bool {
	return obj.obj.StaticKey != nil
}

// description is TBD
// SetStaticKey sets the MacsecSecYRxScStaticKey value in the MacsecSecYRxSc object
func (obj *macsecSecYRxSc) SetStaticKey(value MacsecSecYRxScStaticKey) MacsecSecYRxSc {

	obj.staticKeyHolder = nil
	obj.obj.StaticKey = value.msg()

	return obj
}

func (obj *macsecSecYRxSc) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.StaticKey != nil {

		obj.StaticKey().validateObj(vObj, set_default)
	}

}

func (obj *macsecSecYRxSc) setDefault() {

}

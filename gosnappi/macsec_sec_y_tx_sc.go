package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYTxSc *****
type macsecSecYTxSc struct {
	validation
	obj             *otg.MacsecSecYTxSc
	marshaller      marshalMacsecSecYTxSc
	unMarshaller    unMarshalMacsecSecYTxSc
	staticKeyHolder MacsecSecYTxScStaticKey
}

func NewMacsecSecYTxSc() MacsecSecYTxSc {
	obj := macsecSecYTxSc{obj: &otg.MacsecSecYTxSc{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYTxSc) msg() *otg.MacsecSecYTxSc {
	return obj.obj
}

func (obj *macsecSecYTxSc) setMsg(msg *otg.MacsecSecYTxSc) MacsecSecYTxSc {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYTxSc struct {
	obj *macsecSecYTxSc
}

type marshalMacsecSecYTxSc interface {
	// ToProto marshals MacsecSecYTxSc to protobuf object *otg.MacsecSecYTxSc
	ToProto() (*otg.MacsecSecYTxSc, error)
	// ToPbText marshals MacsecSecYTxSc to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYTxSc to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYTxSc to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYTxSc struct {
	obj *macsecSecYTxSc
}

type unMarshalMacsecSecYTxSc interface {
	// FromProto unmarshals MacsecSecYTxSc from protobuf object *otg.MacsecSecYTxSc
	FromProto(msg *otg.MacsecSecYTxSc) (MacsecSecYTxSc, error)
	// FromPbText unmarshals MacsecSecYTxSc from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYTxSc from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYTxSc from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYTxSc) Marshal() marshalMacsecSecYTxSc {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYTxSc{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYTxSc) Unmarshal() unMarshalMacsecSecYTxSc {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYTxSc{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYTxSc) ToProto() (*otg.MacsecSecYTxSc, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYTxSc) FromProto(msg *otg.MacsecSecYTxSc) (MacsecSecYTxSc, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYTxSc) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYTxSc) FromPbText(value string) error {
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

func (m *marshalmacsecSecYTxSc) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYTxSc) FromYaml(value string) error {
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

func (m *marshalmacsecSecYTxSc) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYTxSc) FromJson(value string) error {
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

func (obj *macsecSecYTxSc) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYTxSc) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYTxSc) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYTxSc) Clone() (MacsecSecYTxSc, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYTxSc()
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

func (obj *macsecSecYTxSc) setNil() {
	obj.staticKeyHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecSecYTxSc is the container for Tx SC settings.
type MacsecSecYTxSc interface {
	Validation
	// msg marshals MacsecSecYTxSc to protobuf object *otg.MacsecSecYTxSc
	// and doesn't set defaults
	msg() *otg.MacsecSecYTxSc
	// setMsg unmarshals MacsecSecYTxSc from protobuf object *otg.MacsecSecYTxSc
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYTxSc) MacsecSecYTxSc
	// provides marshal interface
	Marshal() marshalMacsecSecYTxSc
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYTxSc
	// validate validates MacsecSecYTxSc
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYTxSc, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StaticKey returns MacsecSecYTxScStaticKey, set in MacsecSecYTxSc.
	// MacsecSecYTxScStaticKey is tx SC setting for static key.
	StaticKey() MacsecSecYTxScStaticKey
	// SetStaticKey assigns MacsecSecYTxScStaticKey provided by user to MacsecSecYTxSc.
	// MacsecSecYTxScStaticKey is tx SC setting for static key.
	SetStaticKey(value MacsecSecYTxScStaticKey) MacsecSecYTxSc
	// HasStaticKey checks if StaticKey has been set in MacsecSecYTxSc
	HasStaticKey() bool
	setNil()
}

// description is TBD
// StaticKey returns a MacsecSecYTxScStaticKey
func (obj *macsecSecYTxSc) StaticKey() MacsecSecYTxScStaticKey {
	if obj.obj.StaticKey == nil {
		obj.obj.StaticKey = NewMacsecSecYTxScStaticKey().msg()
	}
	if obj.staticKeyHolder == nil {
		obj.staticKeyHolder = &macsecSecYTxScStaticKey{obj: obj.obj.StaticKey}
	}
	return obj.staticKeyHolder
}

// description is TBD
// StaticKey returns a MacsecSecYTxScStaticKey
func (obj *macsecSecYTxSc) HasStaticKey() bool {
	return obj.obj.StaticKey != nil
}

// description is TBD
// SetStaticKey sets the MacsecSecYTxScStaticKey value in the MacsecSecYTxSc object
func (obj *macsecSecYTxSc) SetStaticKey(value MacsecSecYTxScStaticKey) MacsecSecYTxSc {

	obj.staticKeyHolder = nil
	obj.obj.StaticKey = value.msg()

	return obj
}

func (obj *macsecSecYTxSc) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.StaticKey != nil {

		obj.StaticKey().validateObj(vObj, set_default)
	}

}

func (obj *macsecSecYTxSc) setDefault() {

}

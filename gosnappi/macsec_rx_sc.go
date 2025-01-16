package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecRxSc *****
type macsecRxSc struct {
	validation
	obj             *otg.MacsecRxSc
	marshaller      marshalMacsecRxSc
	unMarshaller    unMarshalMacsecRxSc
	staticKeyHolder MacsecRxScStaticKey
}

func NewMacsecRxSc() MacsecRxSc {
	obj := macsecRxSc{obj: &otg.MacsecRxSc{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecRxSc) msg() *otg.MacsecRxSc {
	return obj.obj
}

func (obj *macsecRxSc) setMsg(msg *otg.MacsecRxSc) MacsecRxSc {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecRxSc struct {
	obj *macsecRxSc
}

type marshalMacsecRxSc interface {
	// ToProto marshals MacsecRxSc to protobuf object *otg.MacsecRxSc
	ToProto() (*otg.MacsecRxSc, error)
	// ToPbText marshals MacsecRxSc to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecRxSc to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecRxSc to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecRxSc struct {
	obj *macsecRxSc
}

type unMarshalMacsecRxSc interface {
	// FromProto unmarshals MacsecRxSc from protobuf object *otg.MacsecRxSc
	FromProto(msg *otg.MacsecRxSc) (MacsecRxSc, error)
	// FromPbText unmarshals MacsecRxSc from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecRxSc from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecRxSc from JSON text
	FromJson(value string) error
}

func (obj *macsecRxSc) Marshal() marshalMacsecRxSc {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecRxSc{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecRxSc) Unmarshal() unMarshalMacsecRxSc {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecRxSc{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecRxSc) ToProto() (*otg.MacsecRxSc, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecRxSc) FromProto(msg *otg.MacsecRxSc) (MacsecRxSc, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecRxSc) ToPbText() (string, error) {
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

func (m *unMarshalmacsecRxSc) FromPbText(value string) error {
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

func (m *marshalmacsecRxSc) ToYaml() (string, error) {
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

func (m *unMarshalmacsecRxSc) FromYaml(value string) error {
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

func (m *marshalmacsecRxSc) ToJson() (string, error) {
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

func (m *unMarshalmacsecRxSc) FromJson(value string) error {
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

func (obj *macsecRxSc) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecRxSc) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecRxSc) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecRxSc) Clone() (MacsecRxSc, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecRxSc()
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

func (obj *macsecRxSc) setNil() {
	obj.staticKeyHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecRxSc is the container for Rx SC settings.
type MacsecRxSc interface {
	Validation
	// msg marshals MacsecRxSc to protobuf object *otg.MacsecRxSc
	// and doesn't set defaults
	msg() *otg.MacsecRxSc
	// setMsg unmarshals MacsecRxSc from protobuf object *otg.MacsecRxSc
	// and doesn't set defaults
	setMsg(*otg.MacsecRxSc) MacsecRxSc
	// provides marshal interface
	Marshal() marshalMacsecRxSc
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecRxSc
	// validate validates MacsecRxSc
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecRxSc, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StaticKey returns MacsecRxScStaticKey, set in MacsecRxSc.
	// MacsecRxScStaticKey is rx SC settings for static key.
	StaticKey() MacsecRxScStaticKey
	// SetStaticKey assigns MacsecRxScStaticKey provided by user to MacsecRxSc.
	// MacsecRxScStaticKey is rx SC settings for static key.
	SetStaticKey(value MacsecRxScStaticKey) MacsecRxSc
	// HasStaticKey checks if StaticKey has been set in MacsecRxSc
	HasStaticKey() bool
	setNil()
}

// description is TBD
// StaticKey returns a MacsecRxScStaticKey
func (obj *macsecRxSc) StaticKey() MacsecRxScStaticKey {
	if obj.obj.StaticKey == nil {
		obj.obj.StaticKey = NewMacsecRxScStaticKey().msg()
	}
	if obj.staticKeyHolder == nil {
		obj.staticKeyHolder = &macsecRxScStaticKey{obj: obj.obj.StaticKey}
	}
	return obj.staticKeyHolder
}

// description is TBD
// StaticKey returns a MacsecRxScStaticKey
func (obj *macsecRxSc) HasStaticKey() bool {
	return obj.obj.StaticKey != nil
}

// description is TBD
// SetStaticKey sets the MacsecRxScStaticKey value in the MacsecRxSc object
func (obj *macsecRxSc) SetStaticKey(value MacsecRxScStaticKey) MacsecRxSc {

	obj.staticKeyHolder = nil
	obj.obj.StaticKey = value.msg()

	return obj
}

func (obj *macsecRxSc) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.StaticKey != nil {

		obj.StaticKey().validateObj(vObj, set_default)
	}

}

func (obj *macsecRxSc) setDefault() {

}

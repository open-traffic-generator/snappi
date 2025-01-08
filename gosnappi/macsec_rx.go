package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecRx *****
type macsecRx struct {
	validation
	obj             *otg.MacsecRx
	marshaller      marshalMacsecRx
	unMarshaller    unMarshalMacsecRx
	staticKeyHolder MacsecRxStaticKey
}

func NewMacsecRx() MacsecRx {
	obj := macsecRx{obj: &otg.MacsecRx{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecRx) msg() *otg.MacsecRx {
	return obj.obj
}

func (obj *macsecRx) setMsg(msg *otg.MacsecRx) MacsecRx {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecRx struct {
	obj *macsecRx
}

type marshalMacsecRx interface {
	// ToProto marshals MacsecRx to protobuf object *otg.MacsecRx
	ToProto() (*otg.MacsecRx, error)
	// ToPbText marshals MacsecRx to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecRx to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecRx to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecRx struct {
	obj *macsecRx
}

type unMarshalMacsecRx interface {
	// FromProto unmarshals MacsecRx from protobuf object *otg.MacsecRx
	FromProto(msg *otg.MacsecRx) (MacsecRx, error)
	// FromPbText unmarshals MacsecRx from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecRx from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecRx from JSON text
	FromJson(value string) error
}

func (obj *macsecRx) Marshal() marshalMacsecRx {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecRx{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecRx) Unmarshal() unMarshalMacsecRx {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecRx{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecRx) ToProto() (*otg.MacsecRx, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecRx) FromProto(msg *otg.MacsecRx) (MacsecRx, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecRx) ToPbText() (string, error) {
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

func (m *unMarshalmacsecRx) FromPbText(value string) error {
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

func (m *marshalmacsecRx) ToYaml() (string, error) {
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

func (m *unMarshalmacsecRx) FromYaml(value string) error {
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

func (m *marshalmacsecRx) ToJson() (string, error) {
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

func (m *unMarshalmacsecRx) FromJson(value string) error {
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

func (obj *macsecRx) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecRx) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecRx) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecRx) Clone() (MacsecRx, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecRx()
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

func (obj *macsecRx) setNil() {
	obj.staticKeyHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecRx is the container for Rx settings.
type MacsecRx interface {
	Validation
	// msg marshals MacsecRx to protobuf object *otg.MacsecRx
	// and doesn't set defaults
	msg() *otg.MacsecRx
	// setMsg unmarshals MacsecRx from protobuf object *otg.MacsecRx
	// and doesn't set defaults
	setMsg(*otg.MacsecRx) MacsecRx
	// provides marshal interface
	Marshal() marshalMacsecRx
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecRx
	// validate validates MacsecRx
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecRx, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StaticKey returns MacsecRxStaticKey, set in MacsecRx.
	// MacsecRxStaticKey is static key Rx settings.
	StaticKey() MacsecRxStaticKey
	// SetStaticKey assigns MacsecRxStaticKey provided by user to MacsecRx.
	// MacsecRxStaticKey is static key Rx settings.
	SetStaticKey(value MacsecRxStaticKey) MacsecRx
	// HasStaticKey checks if StaticKey has been set in MacsecRx
	HasStaticKey() bool
	setNil()
}

// description is TBD
// StaticKey returns a MacsecRxStaticKey
func (obj *macsecRx) StaticKey() MacsecRxStaticKey {
	if obj.obj.StaticKey == nil {
		obj.obj.StaticKey = NewMacsecRxStaticKey().msg()
	}
	if obj.staticKeyHolder == nil {
		obj.staticKeyHolder = &macsecRxStaticKey{obj: obj.obj.StaticKey}
	}
	return obj.staticKeyHolder
}

// description is TBD
// StaticKey returns a MacsecRxStaticKey
func (obj *macsecRx) HasStaticKey() bool {
	return obj.obj.StaticKey != nil
}

// description is TBD
// SetStaticKey sets the MacsecRxStaticKey value in the MacsecRx object
func (obj *macsecRx) SetStaticKey(value MacsecRxStaticKey) MacsecRx {

	obj.staticKeyHolder = nil
	obj.obj.StaticKey = value.msg()

	return obj
}

func (obj *macsecRx) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.StaticKey != nil {

		obj.StaticKey().validateObj(vObj, set_default)
	}

}

func (obj *macsecRx) setDefault() {

}

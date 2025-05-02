package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisIIHTlvs *****
type isisIIHTlvs struct {
	validation
	obj              *otg.IsisIIHTlvs
	marshaller       marshalIsisIIHTlvs
	unMarshaller     unMarshalIsisIIHTlvs
	restartTlvHolder IsisIIHRestartTlv
}

func NewIsisIIHTlvs() IsisIIHTlvs {
	obj := isisIIHTlvs{obj: &otg.IsisIIHTlvs{}}
	obj.setDefault()
	return &obj
}

func (obj *isisIIHTlvs) msg() *otg.IsisIIHTlvs {
	return obj.obj
}

func (obj *isisIIHTlvs) setMsg(msg *otg.IsisIIHTlvs) IsisIIHTlvs {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisIIHTlvs struct {
	obj *isisIIHTlvs
}

type marshalIsisIIHTlvs interface {
	// ToProto marshals IsisIIHTlvs to protobuf object *otg.IsisIIHTlvs
	ToProto() (*otg.IsisIIHTlvs, error)
	// ToPbText marshals IsisIIHTlvs to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisIIHTlvs to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisIIHTlvs to JSON text
	ToJson() (string, error)
}

type unMarshalisisIIHTlvs struct {
	obj *isisIIHTlvs
}

type unMarshalIsisIIHTlvs interface {
	// FromProto unmarshals IsisIIHTlvs from protobuf object *otg.IsisIIHTlvs
	FromProto(msg *otg.IsisIIHTlvs) (IsisIIHTlvs, error)
	// FromPbText unmarshals IsisIIHTlvs from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisIIHTlvs from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisIIHTlvs from JSON text
	FromJson(value string) error
}

func (obj *isisIIHTlvs) Marshal() marshalIsisIIHTlvs {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisIIHTlvs{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisIIHTlvs) Unmarshal() unMarshalIsisIIHTlvs {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisIIHTlvs{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisIIHTlvs) ToProto() (*otg.IsisIIHTlvs, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisIIHTlvs) FromProto(msg *otg.IsisIIHTlvs) (IsisIIHTlvs, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisIIHTlvs) ToPbText() (string, error) {
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

func (m *unMarshalisisIIHTlvs) FromPbText(value string) error {
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

func (m *marshalisisIIHTlvs) ToYaml() (string, error) {
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

func (m *unMarshalisisIIHTlvs) FromYaml(value string) error {
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

func (m *marshalisisIIHTlvs) ToJson() (string, error) {
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

func (m *unMarshalisisIIHTlvs) FromJson(value string) error {
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

func (obj *isisIIHTlvs) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisIIHTlvs) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisIIHTlvs) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisIIHTlvs) Clone() (IsisIIHTlvs, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisIIHTlvs()
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

func (obj *isisIIHTlvs) setNil() {
	obj.restartTlvHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisIIHTlvs is this contains the list of TLVs present in a IIH packet.
type IsisIIHTlvs interface {
	Validation
	// msg marshals IsisIIHTlvs to protobuf object *otg.IsisIIHTlvs
	// and doesn't set defaults
	msg() *otg.IsisIIHTlvs
	// setMsg unmarshals IsisIIHTlvs from protobuf object *otg.IsisIIHTlvs
	// and doesn't set defaults
	setMsg(*otg.IsisIIHTlvs) IsisIIHTlvs
	// provides marshal interface
	Marshal() marshalIsisIIHTlvs
	// provides unmarshal interface
	Unmarshal() unMarshalIsisIIHTlvs
	// validate validates IsisIIHTlvs
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisIIHTlvs, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RestartTlv returns IsisIIHRestartTlv, set in IsisIIHTlvs.
	// IsisIIHRestartTlv is container of Restart Tlv.
	RestartTlv() IsisIIHRestartTlv
	// SetRestartTlv assigns IsisIIHRestartTlv provided by user to IsisIIHTlvs.
	// IsisIIHRestartTlv is container of Restart Tlv.
	SetRestartTlv(value IsisIIHRestartTlv) IsisIIHTlvs
	// HasRestartTlv checks if RestartTlv has been set in IsisIIHTlvs
	HasRestartTlv() bool
	setNil()
}

// Restart Tlv.
// RestartTlv returns a IsisIIHRestartTlv
func (obj *isisIIHTlvs) RestartTlv() IsisIIHRestartTlv {
	if obj.obj.RestartTlv == nil {
		obj.obj.RestartTlv = NewIsisIIHRestartTlv().msg()
	}
	if obj.restartTlvHolder == nil {
		obj.restartTlvHolder = &isisIIHRestartTlv{obj: obj.obj.RestartTlv}
	}
	return obj.restartTlvHolder
}

// Restart Tlv.
// RestartTlv returns a IsisIIHRestartTlv
func (obj *isisIIHTlvs) HasRestartTlv() bool {
	return obj.obj.RestartTlv != nil
}

// Restart Tlv.
// SetRestartTlv sets the IsisIIHRestartTlv value in the IsisIIHTlvs object
func (obj *isisIIHTlvs) SetRestartTlv(value IsisIIHRestartTlv) IsisIIHTlvs {

	obj.restartTlvHolder = nil
	obj.obj.RestartTlv = value.msg()

	return obj
}

func (obj *isisIIHTlvs) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RestartTlv != nil {

		obj.RestartTlv().validateObj(vObj, set_default)
	}

}

func (obj *isisIIHTlvs) setDefault() {

}

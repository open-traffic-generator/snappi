package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisIIHNeighborTlvs *****
type isisIIHNeighborTlvs struct {
	validation
	obj              *otg.IsisIIHNeighborTlvs
	marshaller       marshalIsisIIHNeighborTlvs
	unMarshaller     unMarshalIsisIIHNeighborTlvs
	restartTlvHolder IsisIIHRestartTlv
}

func NewIsisIIHNeighborTlvs() IsisIIHNeighborTlvs {
	obj := isisIIHNeighborTlvs{obj: &otg.IsisIIHNeighborTlvs{}}
	obj.setDefault()
	return &obj
}

func (obj *isisIIHNeighborTlvs) msg() *otg.IsisIIHNeighborTlvs {
	return obj.obj
}

func (obj *isisIIHNeighborTlvs) setMsg(msg *otg.IsisIIHNeighborTlvs) IsisIIHNeighborTlvs {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisIIHNeighborTlvs struct {
	obj *isisIIHNeighborTlvs
}

type marshalIsisIIHNeighborTlvs interface {
	// ToProto marshals IsisIIHNeighborTlvs to protobuf object *otg.IsisIIHNeighborTlvs
	ToProto() (*otg.IsisIIHNeighborTlvs, error)
	// ToPbText marshals IsisIIHNeighborTlvs to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisIIHNeighborTlvs to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisIIHNeighborTlvs to JSON text
	ToJson() (string, error)
}

type unMarshalisisIIHNeighborTlvs struct {
	obj *isisIIHNeighborTlvs
}

type unMarshalIsisIIHNeighborTlvs interface {
	// FromProto unmarshals IsisIIHNeighborTlvs from protobuf object *otg.IsisIIHNeighborTlvs
	FromProto(msg *otg.IsisIIHNeighborTlvs) (IsisIIHNeighborTlvs, error)
	// FromPbText unmarshals IsisIIHNeighborTlvs from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisIIHNeighborTlvs from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisIIHNeighborTlvs from JSON text
	FromJson(value string) error
}

func (obj *isisIIHNeighborTlvs) Marshal() marshalIsisIIHNeighborTlvs {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisIIHNeighborTlvs{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisIIHNeighborTlvs) Unmarshal() unMarshalIsisIIHNeighborTlvs {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisIIHNeighborTlvs{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisIIHNeighborTlvs) ToProto() (*otg.IsisIIHNeighborTlvs, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisIIHNeighborTlvs) FromProto(msg *otg.IsisIIHNeighborTlvs) (IsisIIHNeighborTlvs, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisIIHNeighborTlvs) ToPbText() (string, error) {
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

func (m *unMarshalisisIIHNeighborTlvs) FromPbText(value string) error {
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

func (m *marshalisisIIHNeighborTlvs) ToYaml() (string, error) {
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

func (m *unMarshalisisIIHNeighborTlvs) FromYaml(value string) error {
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

func (m *marshalisisIIHNeighborTlvs) ToJson() (string, error) {
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

func (m *unMarshalisisIIHNeighborTlvs) FromJson(value string) error {
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

func (obj *isisIIHNeighborTlvs) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisIIHNeighborTlvs) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisIIHNeighborTlvs) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisIIHNeighborTlvs) Clone() (IsisIIHNeighborTlvs, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisIIHNeighborTlvs()
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

func (obj *isisIIHNeighborTlvs) setNil() {
	obj.restartTlvHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisIIHNeighborTlvs is this contains the list of TLVs present in a IIH PDU received from a neighbor IS-IS router.
type IsisIIHNeighborTlvs interface {
	Validation
	// msg marshals IsisIIHNeighborTlvs to protobuf object *otg.IsisIIHNeighborTlvs
	// and doesn't set defaults
	msg() *otg.IsisIIHNeighborTlvs
	// setMsg unmarshals IsisIIHNeighborTlvs from protobuf object *otg.IsisIIHNeighborTlvs
	// and doesn't set defaults
	setMsg(*otg.IsisIIHNeighborTlvs) IsisIIHNeighborTlvs
	// provides marshal interface
	Marshal() marshalIsisIIHNeighborTlvs
	// provides unmarshal interface
	Unmarshal() unMarshalIsisIIHNeighborTlvs
	// validate validates IsisIIHNeighborTlvs
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisIIHNeighborTlvs, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RestartTlv returns IsisIIHRestartTlv, set in IsisIIHNeighborTlvs.
	// IsisIIHRestartTlv is container of Restart TLV in IIH PDU. Reference: https://datatracker.ietf.org/doc/html/rfc8706#name-restart-tlv
	RestartTlv() IsisIIHRestartTlv
	// SetRestartTlv assigns IsisIIHRestartTlv provided by user to IsisIIHNeighborTlvs.
	// IsisIIHRestartTlv is container of Restart TLV in IIH PDU. Reference: https://datatracker.ietf.org/doc/html/rfc8706#name-restart-tlv
	SetRestartTlv(value IsisIIHRestartTlv) IsisIIHNeighborTlvs
	// HasRestartTlv checks if RestartTlv has been set in IsisIIHNeighborTlvs
	HasRestartTlv() bool
	setNil()
}

// Restart Tlv.
// RestartTlv returns a IsisIIHRestartTlv
func (obj *isisIIHNeighborTlvs) RestartTlv() IsisIIHRestartTlv {
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
func (obj *isisIIHNeighborTlvs) HasRestartTlv() bool {
	return obj.obj.RestartTlv != nil
}

// Restart Tlv.
// SetRestartTlv sets the IsisIIHRestartTlv value in the IsisIIHNeighborTlvs object
func (obj *isisIIHNeighborTlvs) SetRestartTlv(value IsisIIHRestartTlv) IsisIIHNeighborTlvs {

	obj.restartTlvHolder = nil
	obj.obj.RestartTlv = value.msg()

	return obj
}

func (obj *isisIIHNeighborTlvs) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RestartTlv != nil {

		obj.RestartTlv().validateObj(vObj, set_default)
	}

}

func (obj *isisIIHNeighborTlvs) setDefault() {

}
